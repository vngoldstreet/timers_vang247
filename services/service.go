package service

import (
	"net/http"
	"net/url"
	"strings"
)

// PerformRequest thực hiện yêu cầu HTTP
func PerformRequest(client *http.Client, method string, url string, token string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return client.Do(req)
}

// GetRequest gửi yêu cầu GET và đợi phản hồi
func GetRequest(url string, token string) (*http.Response, error) {
	client := &http.Client{
		// Timeout: time.Second * 30,
	}
	resp, err := PerformRequest(client, "GET", url, token, nil)
	if err != nil {
		return nil, err
	}
	// Không đóng resp.Body ở đây vì chúng ta sẽ đọc từ nó trong hàm khác
	return resp, nil
}

func PutRequest(url string, token string) (*http.Response, error) {
	client := &http.Client{
		// Timeout: time.Second * 30,
	}
	resp, err := PerformRequest(client, "PUT", url, token, nil)
	if err != nil {
		return nil, err
	}
	// Không đóng resp.Body ở đây vì chúng ta sẽ đọc từ nó trong hàm khác
	return resp, nil
}

func PostRequest(url string, token string, body []byte) (*http.Response, error) {
	client := &http.Client{
		// Timeout: time.Second * 30,
	}
	resp, err := PerformRequest(client, "POST", url, token, body)
	if err != nil {
		return nil, err
	}
	// Không đóng resp.Body ở đây vì chúng ta sẽ đọc từ nó trong hàm khác
	return resp, nil
}

// Dien vao url, token & form
func PostFormRequest(url string, token string, formData url.Values) (*http.Response, error) {
	client := &http.Client{
		// Timeout: time.Second * 30,
	}
	body := strings.NewReader(formData.Encode())
	resp, err := PerformRequestWithForm(client, "POST", url, token, body, "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}
	// Không đóng resp.Body ở đây vì chúng ta sẽ đọc từ nó trong hàm khác
	return resp, nil
}

func PutFormRequest(url string, token string, formData url.Values) (*http.Response, error) {
	client := &http.Client{
		// Timeout: time.Second * 30,
	}
	body := strings.NewReader(formData.Encode())
	resp, err := PerformRequestWithForm(client, "PUT", url, token, body, "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}
	// Không đóng resp.Body ở đây vì chúng ta sẽ đọc từ nó trong hàm khác
	return resp, nil
}

func PerformRequestWithForm(client *http.Client, method string, url string, token string, body *strings.Reader, contentType string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	return client.Do(req)
}
