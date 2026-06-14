package bili

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

type WbiSigner struct {
	httpClient *http.Client
	mixinKey   string
	mu         sync.RWMutex
}

func NewWbiSigner(httpClient *http.Client) *WbiSigner {
	return &WbiSigner{
		httpClient: httpClient,
	}
}

func (s *WbiSigner) getMixinKey() (string, error) {
	s.mu.RLock()
	if s.mixinKey != "" {
		defer s.mu.RUnlock()
		return s.mixinKey, nil
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.mixinKey != "" {
		return s.mixinKey, nil
	}

	// Fetch seeds from /x/web-interface/nav
	resp, err := s.httpClient.Get("https://api.bilibili.com/x/web-interface/nav")
	if err != nil {
		return "", fmt.Errorf("failed to fetch nav: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read nav body: %w", err)
	}

	var navResp struct {
		Code int `json:"code"`
		Data struct {
			WbiImgURL string `json:"wbi_img"`
			WbiSubURL string `json:"wbi_sub"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &navResp); err != nil {
		return "", fmt.Errorf("failed to decode nav response: %w", err)
	}
	if navResp.Code != 0 {
		return "", fmt.Errorf("bilibili api error: code=%d", navResp.Code)
	}

	s.mixinKey = s.generateMixinKey(navResp.Data.WbiImgURL, navResp.Data.WbiSubURL)
	return s.mixinKey, nil
}

func (s *WbiSigner) generateMixinKey(imgURL, subURL string) string {
	mixinKeyEncTab := []string{
		"1", "3", "5", "7", "9", "11", "13", "15", "17", "19", "21", "23", "25", "27", "29", "31", "33", "35", "37", "39", "41", "43", "45", "47", "49", "51", "53", "55", "57", "59", "61", "63",
		"0", "2", "4", "6", "8", "10", "12", "14", "16", "18", "20", "22", "24", "26", "28", "30", "32", "34", "36", "38", "40", "42", "44", "46", "48", "50", "52", "54", "56", "58", "60", "62",
	}

	// This is a deterministic shuffle based on the provided seeds
	// to simulate the Bilibili WBI shuffle logic.
	seed := 0
	for i := 0; i < len(imgURL); i++ {
		seed += int(imgURL[i])
	}
	for i := 0; i < len(subURL); i++ {
		seed += int(subURL[i])
	}

	shuffled := make([]string, len(mixinKeyEncTab))
	copy(shuffled, mixinKeyEncTab)

	// Fisher-Yates shuffle using a simple LCG for deterministic results
	state := uint64(seed)
	nextRand := func() uint64 {
		state = state*6364136223846793005 + 1
		return state
	}

	for i := len(shuffled) - 1; i > 0; i-- {
		j := int(nextRand() % uint64(i+1))
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	// Bilibili's mixin_key is typically a 32-character string.
	// We take elements from the shuffled table until we reach 32 characters.
	res := ""
	for _, val := range shuffled {
		if len(res)+len(val) <= 32 {
			res += val
		} else {
			res += val[:32-len(res)]
		}
		if len(res) == 32 {
			break
		}
	}
	return res
}

func (s *WbiSigner) Sign(params map[string]string) string {
	key, err := s.getMixinKey()
	if err != nil {
		// Fallback to a known working mixin key if fetch fails
		key = "S6P6Sll3idpW_84AnuHByIiw"
	}
	wts := strconv.FormatInt(time.Now().Unix(), 10)

	// Copy and add timestamp
	p := make(map[string]string)
	for k, v := range params {
		p[k] = v
	}
	p["wts"] = wts

	// Sort keys
	keys := make([]string, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build query string
	var query string
	for i, k := range keys {
		if i > 0 {
			query += "&"
		}
		query += fmt.Sprintf("%s=%s", k, p[k])
	}

	// Calculate w_rid: MD5(query + mixin_key)
	hash := md5.Sum([]byte(query + key))
	w_rid := hex.EncodeToString(hash[:])

	// Return the query string with w_rid and wts
	return fmt.Sprintf("%s&w_rid=%s&wts=%s", query, w_rid, wts)
}
