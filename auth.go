package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var tokenCache sync.Map

var clientID string
var clientSecret string
var introspectURL string
var issuerPrefix string

type TokenData struct {
	Sub      string `json:"sub,omitempty"`
	Email    string `json:"email,omitempty"`
	Scope    string `json:"scope,omitempty"`
	Exp 		 int    `json:"exp,omitempty"`
	Iss			 string `json:"iss,omitempty"`
	Active   bool   `json:"active"`
}

func (t *TokenData) GetIssuer() string {
	return strings.TrimSuffix(strings.TrimPrefix(t.Iss, issuerPrefix), "/")
}

func SetupAuth() {
	clientID = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	introspectURL = os.Getenv("INTROSPECT_URL")
	issuerPrefix = os.Getenv("ISSUER_PREFIX")

	if clientID == "" || clientSecret == "" || introspectURL == "" || issuerPrefix == "" {
		panic("missing environment variables")
	}
}

func IntrospectToken(token string) (TokenData, error) {
	data := url.Values{}
	data.Set("token", token)

	req, err := http.NewRequest("POST", introspectURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return TokenData{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return TokenData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return TokenData{}, errors.New("invalid token")
	}

	// Read the JSON response body into a TokenData struct
	var tokenData TokenData
	err = json.NewDecoder(resp.Body).Decode(&tokenData)
	if err != nil {
		return TokenData{}, err
	}

	if !tokenData.Active {
		return TokenData{}, errors.New("token is not active")
	}

	return tokenData, nil
}

func ValidateToken(token string) (TokenData, error) {
	if token == "" {
		return TokenData{}, errors.New("missing token")
	}

	if data, ok := tokenCache.Load(token); ok {
		if tokenData, ok := data.(TokenData); ok {
			if tokenData.Exp < int(time.Now().Unix()) {
				return TokenData{}, errors.New("token expired")
			}
			return tokenData, nil
		} else {
			return TokenData{}, errors.New("invalid token")
		}
	}

	tokenData, err := IntrospectToken(token)
	if err != nil {
		return TokenData{}, err
	}

	if !strings.HasPrefix(tokenData.Iss, issuerPrefix) {
		return TokenData{}, errors.New("invalid issuer")
	}

	tokenCache.Store(token, tokenData)
	return tokenData, nil
}

func ValidateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		tokenData, err := ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("tokenData", tokenData)
		c.Next()
	}
}
