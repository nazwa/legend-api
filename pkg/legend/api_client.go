package legend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
)

// LEGEND_DATE_FORMAT has the time format used by most Legend dates
const LEGEND_DATE_FORMAT = "2006-01-02"

type ApiClient struct {
	client    *retryablehttp.Client
	header    http.Header
	endpoint  string
	useMocks  bool
	saveMocks bool
}

// ConnectionConfig stores configuration data for Legend ApiClient
// Credentials MUST be provided at least
type ConnectionConfig struct {
	UserAgent string
	Username  string
	Password  string
	Guid      string
	Endpoint  string

	// How many times should 5xx requests be retried. Defaults to 0 (don't retry)
	MaxRetries int
	// Request timeout (defaults to 120s)
	HttpTimeout time.Duration
	// ProxyUrl allows you to connect through proxies (usefull for local debugging)
	ProxyUrl string
	// Use locally-stored dumps for core-data instead of making GET requests
	// Price calculations & POST requests are still sent externally
	UseMocks  bool
	// Persist core-data dumps after making GET requests
	SaveMocks bool
}

// NewApiClient creates a basic Legend API client with a retry-mechanic
// It will automatically retry on 5xx errors up to the specified retryMax
func NewApiClient(config ConnectionConfig) *ApiClient {
	client := retryablehttp.NewClient()

	if config.HttpTimeout == 0 {
		config.HttpTimeout = 120 * time.Second
	}

	client.RetryWaitMin = 300 * time.Millisecond
	client.RetryWaitMax = 1 * time.Second
	client.RetryMax = config.MaxRetries
	client.HTTPClient.Timeout = config.HttpTimeout

	client.CheckRetry = retryPolicy

	if config.ProxyUrl != "" {
		proxyUrl, _ := url.Parse(config.ProxyUrl)
		httpClient := &http.Client{
			Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
			Timeout:   config.HttpTimeout,
		}
		client.HTTPClient = httpClient
	}

	headers := make(http.Header)
	headers.Add("User-Agent", config.UserAgent)
	headers.Add("Username", config.Username)
	headers.Add("Password", config.Password)
	headers.Add("Customerguid", config.Guid)
	headers.Set("Content-Type", "application/json")

	return &ApiClient{
		client:    client,
		header:    headers,
		endpoint:  config.Endpoint,
		useMocks:  config.UseMocks,
		saveMocks: config.SaveMocks,
	}
}

// GetApiRequest issues a GET requests and unmarshals JSON response into the provided struct
func (this *ApiClient) GetApiRequest(path string, params url.Values, responseTarget interface{}) error {
	finalUrl := fmt.Sprintf("%s%s?%s", this.endpoint, path, params.Encode())

	req, err := retryablehttp.NewRequest("GET", finalUrl, nil)
	if err != nil {
		return errors.Wrap(err, "NewRequest")
	}

	req.Header = this.header

	resp, err := this.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "Do")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.Wrap(errors.New(resp.Status), "Response")
	}

	err = json.NewDecoder(resp.Body).Decode(responseTarget)
	if err != nil {
		return errors.Wrap(err, "Decoder")
	}
	return nil
}

// PostApiRequest issues a POST requests and unmarshals JSON response into the provided struct
func (this *ApiClient) PostApiRequest(path string, params url.Values, body, responseTarget interface{}) error {
	finalUrl := fmt.Sprintf("%s%s?%s", this.endpoint, path, params.Encode())

	encodedJson, err := json.Marshal(body)
	if err != nil {
		return errors.Wrap(err, "Encoder")
	}

	req, err := retryablehttp.NewRequest("POST", finalUrl, bytes.NewBuffer(encodedJson))
	if err != nil {
		return errors.Wrap(err, "NewRequest")
	}

	req.Header = this.header

	resp, err := this.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "Do")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(responseTarget)
	if err != nil {
		return errors.Wrap(err, "Decoder")
	}

	return nil
}

// PostApiRequestRawResponse issues a POST requests and returns the raw, unprocessed response as a string
// Usefull for requests which don't respond with a valid JSON (recurring fee for example)
func (this *ApiClient) PostApiRequestRawResponse(path string, params url.Values, body interface{}) (string, error) {
	finalUrl := fmt.Sprintf("%s%s?%s", this.endpoint, path, params.Encode())

	encodedJson, err := json.Marshal(body)
	if err != nil {
		return "", errors.Wrap(err, "Encoder")
	}

	req, err := retryablehttp.NewRequest("POST", finalUrl, bytes.NewBuffer(encodedJson))
	if err != nil {
		return "", errors.Wrap(err, "NewRequest")
	}

	req.Header = this.header

	resp, err := this.client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "Do")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", errors.Wrap(err, "RespReader")
	}

	return string(respBody), nil
}

func retryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	shouldRetry, _ := retryablehttp.DefaultRetryPolicy(ctx, resp, err)

	if shouldRetry {
		return true, nil
	}

	return resp.StatusCode == http.StatusForbidden, nil
}

func (this *ApiClient) getMockData(path string, target interface{}) error {
	mockFile, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "MockOpening")
	}

	defer mockFile.Close()

	return json.NewDecoder(mockFile).Decode(target)
}

func (this *ApiClient) saveMockData(path string, source interface{}) error {
	mockFile, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "MockSaveOpening")
	}

	defer mockFile.Close()

	return json.NewEncoder(mockFile).Encode(source)
}
