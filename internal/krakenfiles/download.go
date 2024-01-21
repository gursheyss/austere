package krakenfiles

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ApiResponse struct {
    Status string `json:"status"`
    Url    string `json:"url"`
}

func Convert(urlStr string) (string, error) {
    formAction, token, err := extractFormData(urlStr)
    if err != nil {
        return "", err
    }

    formData := url.Values{}
    formData.Set("token", token)

    req, err := http.NewRequest("POST", formAction, strings.NewReader(formData.Encode()))
    if err != nil {
        return "", err
    }
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var apiResponse ApiResponse
    if err := json.Unmarshal(body, &apiResponse); err != nil {
        return "", err
    }

    return apiResponse.Url, nil
}

func extractFormData(urlStr string) (string, string, error) {
    resp, err := http.Get(urlStr)
    if err != nil {
        return "", "", err
    }
    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return "", "", err
    }

    formAction, exists := doc.Find("form#dl-form").Attr("action")
    if !exists {
        return "", "", errors.New("form action not found")
    }
    if strings.HasPrefix(formAction, "//") {
        formAction = "https:" + formAction
    }

    token, exists := doc.Find("input#dl-token").Attr("value")
    if !exists {
        return "", "", errors.New("token value not found")
    }

    return formAction, token, nil
}