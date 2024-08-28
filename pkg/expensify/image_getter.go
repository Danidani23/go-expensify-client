package expensify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"time"
)

// Image struct holds the file name, file extension, and data
type Image struct {
	FileName string
	FileExt  string
	Data     []byte
}

// GetImageFromReceipt get image downloads the image from a URL, it requires a set of active session cookies
func GetImageFromReceipt(url string, cookies []*http.Cookie) (*Image, error) {
	myImage := Image{}

	// Create a new HTTP client with a cookie jar
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Add cookies to the request
	for _, cookie := range cookies {
		// Add the cookie to the cookie jar for the specific domain
		jar.SetCookies(req.URL, []*http.Cookie{cookie})
	}

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	// Read the image data
	myImage.Data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading image data: %w", err)
	}

	// Parse the URL to get the file name and extension
	myImage.FileName = filepath.Base(url)
	myImage.FileExt = filepath.Ext(url)

	return &myImage, nil
}

func (i *Image) WriteToDisk(path string) error {
	fullPath := filepath.Join(path, i.FileName)
	return os.WriteFile(fullPath, i.Data, 0644)
}

// cookie structure for reading JSON cookies
type cookie struct {
	Domain         string  `json:"domain"`
	ExpirationDate float64 `json:"expirationDate"`
	HostOnly       bool    `json:"hostOnly"`
	HttpOnly       bool    `json:"httpOnly"`
	Name           string  `json:"name"`
	Path           string  `json:"path"`
	SameSite       string  `json:"sameSite"`
	Secure         bool    `json:"secure"`
	Session        bool    `json:"session"`
	StoreID        string  `json:"storeId"`
	Value          string  `json:"value"`
	ID             int     `json:"id"`
}

// LoadCookiesFromJSON  reads cookies from a JSON file and returns them
func LoadCookiesFromJSON(filePath string) ([]*http.Cookie, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open cookie file: %w", err)
	}
	defer file.Close()

	var jsonCookies []cookie
	if err := json.NewDecoder(file).Decode(&jsonCookies); err != nil {
		return nil, fmt.Errorf("failed to decode cookies: %w", err)
	}

	var cookies []*http.Cookie
	for _, c := range jsonCookies {
		cookies = append(cookies, &http.Cookie{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  time.Unix(int64(c.ExpirationDate), 0),
			Secure:   c.Secure,
			HttpOnly: c.HttpOnly,
		})
	}

	return cookies, nil
}

// LoadCookiesFromString reads cookies from a JSON string and returns them
func LoadCookiesFromString(cookiesJsonString string) ([]*http.Cookie, error) {
	var jsonCookies []cookie

	// Parse the JSON string into the jsonCookies slice
	if err := json.Unmarshal([]byte(cookiesJsonString), &jsonCookies); err != nil {
		return nil, fmt.Errorf("failed to decode cookies: %w", err)
	}

	var cookies []*http.Cookie
	for _, c := range jsonCookies {
		cookies = append(cookies, &http.Cookie{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  time.Unix(int64(c.ExpirationDate), 0),
			Secure:   c.Secure,
			HttpOnly: c.HttpOnly,
		})
	}

	return cookies, nil
}
