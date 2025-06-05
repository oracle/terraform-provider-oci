// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package auth

import (
	"bytes"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/utils"
)

const (
	rpstValidForRatio float64 = 0.5
)

// Workload RPST Issuance Service (WRIS)
// x509FederationClientForOkeWorkloadIdentity retrieves a security token from Auth service.
type x509FederationClientForOkeWorkloadIdentity struct {
	tenancyID                    string
	sessionKeySupplier           sessionKeySupplier
	securityToken                securityToken
	authClient                   *common.BaseClient
	mux                          sync.Mutex
	proxymuxEndpoint             string
	saTokenProvider              ServiceAccountTokenProvider
	kubernetesServiceAccountCert *x509.CertPool
}

func newX509FederationClientForOkeWorkloadIdentity(endpoint string, saTokenProvider ServiceAccountTokenProvider,
	kubernetesServiceAccountCert *x509.CertPool) (federationClient, error) {
	client := &x509FederationClientForOkeWorkloadIdentity{
		proxymuxEndpoint:             endpoint,
		saTokenProvider:              saTokenProvider,
		kubernetesServiceAccountCert: kubernetesServiceAccountCert,
	}

	client.sessionKeySupplier = newSessionKeySupplier()

	return client, nil
}

func (c *x509FederationClientForOkeWorkloadIdentity) renewSecurityToken() (err error) {
	if err = c.sessionKeySupplier.Refresh(); err != nil {
		return fmt.Errorf("failed to refresh session key: %s", err.Error())
	}

	common.Logf("Renewing security token at: %v\n", time.Now().Format("15:04:05.000"))
	if c.securityToken, err = c.getSecurityToken(); err != nil {
		return fmt.Errorf("failed to get security token: %s", err.Error())
	}
	common.Logf("Security token renewed at: %v\n", time.Now().Format("15:04:05.000"))

	return nil
}

type workloadIdentityRequestPayload struct {
	Podkey string `json:"podKey"`
}
type token struct {
	Token string
}

// getSecurityToken get security token from Proxymux
func (c *x509FederationClientForOkeWorkloadIdentity) getSecurityToken() (securityToken, error) {
	client := http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: c.kubernetesServiceAccountCert,
			},
		},
	}

	publicKey := string(c.sessionKeySupplier.PublicKeyPemRaw())
	common.Logf("Public Key for OKE Workload Identity is:", publicKey)
	rawPayload := workloadIdentityRequestPayload{Podkey: publicKey}
	payload, err := json.Marshal(rawPayload)
	if err != nil {
		return nil, fmt.Errorf("error getting security token%s", err)
	}

	common.Logf("Payload for OKE Workload Identity is:", string(payload))
	request, err := http.NewRequest(http.MethodPost, c.proxymuxEndpoint, bytes.NewBuffer(payload))

	if err != nil {
		common.Logf("error %s", err)
		return nil, fmt.Errorf("error getting security token %s", err)
	}

	kubernetesServiceAccountToken, err := c.saTokenProvider.ServiceAccountToken()
	if err != nil {
		common.Logf("error %s", err)
		return nil, fmt.Errorf("error getting service account token %s", err)
	}

	common.Logf("Service Account Token for OKE Workload Identity is: ", kubernetesServiceAccountToken)
	request.Header.Add("Authorization", "Bearer "+kubernetesServiceAccountToken)
	request.Header.Set("Content-Type", "application/json")
	opcRequestID := utils.GenerateOpcRequestID()
	request.Header.Set("opc-request-id", opcRequestID)

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error %s", err)
	}

	var body bytes.Buffer
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			common.Logf("error %s", err)
		}
	}(response.Body)

	statusCode := response.StatusCode
	if statusCode != http.StatusOK {
		if statusCode == http.StatusForbidden {
			return nil, fmt.Errorf("please ensure the cluster type is enhanced: Status: %s, Message: %s",
				response.Status, body.String())
		} else {
			return nil, fmt.Errorf("failed to get a RPST token from Proxymux: URL: %s, Status: %s, Message: %s",
				c.proxymuxEndpoint, response.Status, body.String())
		}

	}

	if _, err = body.ReadFrom(response.Body); err != nil {
		return nil, fmt.Errorf("error reading body from Proxymux response: %s", err)
	}

	rawBody := body.String()
	rawBody = rawBody[1 : len(rawBody)-1]
	decodedBodyStr, err := base64.StdEncoding.DecodeString(rawBody)
	if err != nil {
		return nil, fmt.Errorf("error decoding Proxymux response using base64 scheme: %s", err)
	}

	var parsedBody token
	err = json.Unmarshal(decodedBodyStr, &parsedBody)
	if err != nil {
		return nil, fmt.Errorf("error parsing Proxymux response body: %s", err)
	}

	token := parsedBody.Token
	if len(token) == 0 {
		return nil, fmt.Errorf("invalid (empty) token received from Proxymux")
	}
	if len(token) < 3 {
		return nil, fmt.Errorf("invalid token received from Proxymux")
	}

	logTokenInfo(token)
	return newPrincipalToken(token[3:])
}

func logTokenInfo(token string) {
	if strings.TrimSpace(token) == "" {
		common.Logf("Token is null or empty")
		return
	}

	parts := strings.Split(token, ".")
	if len(parts) < 3 {
		common.Logf("Invalid JWT token")
		return
	}

	// Decode the payload
	decodedPayload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		common.Logf("Failed to decode payload: %v\n", err)
		return
	}

	// Parse JSON payload into a map
	var payload map[string]interface{}
	if err := json.Unmarshal(decodedPayload, &payload); err != nil {
		common.Logf("Failed to parse payload JSON: %v\n", err)
		return
	}

	logData := map[string]interface{}{
		"sub":                 payload["sub"],
		"res_id":              payload["res_id"],
		"res_type":            payload["res_type"],
		"ttype":               payload["ttype"],
		"var_service_account": payload["var_service_account"],
		"var_namespace":       payload["var_namespace"],
		"iat":                 payload["iat"],
		"exp":                 payload["exp"],
	}

	logJson, err := json.MarshalIndent(logData, "", "  ")
	if err != nil {
		common.Logf("Failed to serialize log data: %v\n", err)
		return
	}

	common.Logf("RPST token details:\n%s\n", string(logJson))
}

func (c *x509FederationClientForOkeWorkloadIdentity) PrivateKey() (*rsa.PrivateKey, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if err := c.renewSecurityTokenIfNotValid(); err != nil {
		return nil, err
	}
	return c.sessionKeySupplier.PrivateKey(), nil
}

func (c *x509FederationClientForOkeWorkloadIdentity) SecurityToken() (token string, err error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if err = c.renewSecurityTokenIfNotValid(); err != nil {
		return "", err
	}
	return c.securityToken.String(), nil
}

func (c *x509FederationClientForOkeWorkloadIdentity) renewSecurityTokenIfNotValid() (err error) {
	if c.securityToken == nil || !c.securityToken.Valid() {
		if err = c.renewSecurityToken(); err != nil {
			return fmt.Errorf("failed to renew security token: %s", err.Error())
		}
	}
	return nil
}

type workloadIdentityPrincipalToken struct {
	principalToken
}

func (t *workloadIdentityPrincipalToken) Valid() bool {
	// TODO: read rpstValidForRatio from rpst token
	issuedAt := int64(t.jwtToken.payload["iat"].(float64))
	expiredAt := int64(t.jwtToken.payload["exp"].(float64))
	softExpiredAt := issuedAt + int64(float64(expiredAt-issuedAt)*rpstValidForRatio)
	softExpiredAtTime := time.Unix(softExpiredAt, 0)
	now := time.Now().Unix() + int64(bufferTimeBeforeTokenExpiration.Seconds())
	expired := softExpiredAt <= now
	if expired {
		common.Debugf("Token expired at: %v", softExpiredAtTime.Format("15:04:05.000"))
	}
	return !expired
}

func (c *x509FederationClientForOkeWorkloadIdentity) GetClaim(key string) (interface{}, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if err := c.renewSecurityTokenIfNotValid(); err != nil {
		return nil, err
	}
	return c.securityToken.GetClaim(key)
}
