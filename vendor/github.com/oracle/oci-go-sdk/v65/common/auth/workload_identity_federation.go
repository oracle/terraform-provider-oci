// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package auth

import (
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/oracle/oci-go-sdk/v65/common"
)

// TokenIssuer defines a type capable of retrieving tokens for the issuing
// authorization server.
type TokenIssuer interface {
	GetToken() (string, error)
}

// StaticTokenIssuer is a defined TokenIssuer that holds a static token. Not suitable
// for use longer than the validity period of the token.
type StaticTokenIssuer struct {
	token string
}

// GetToken satisfies the TokenIssuer interface for StaticTokenIssuer by returning
// the token held by StaticTokenIssuer.
func (s StaticTokenIssuer) GetToken() (string, error) {
	return s.token, nil
}

// TokenExchangeConfigurationProvider provides OCI configuration via token exchange,
// exposing claims and supporting a custom HTTP client.
type TokenExchangeConfigurationProvider struct {
	federationClient federationClient
	region           common.Region
}

// TokenExchangeConfigurationProviderFromIssuer creates a Configuration Provider from a
// function provided to retrieve a token from an identity provider.
func TokenExchangeConfigurationProviderFromIssuer(tokenIssuer TokenIssuer,
	domainUrl string, clientId string, clientSecret string,
	region string, requestedTokenType string,
	resType string) (common.ConfigurationProvider, error) {

	if tokenIssuer == nil {
		return nil, fmt.Errorf("invalid TokenIssuer")
	}

	authCode := base64.StdEncoding.EncodeToString([]byte(
		clientId + ":" + clientSecret))

	requestData := map[string][]string{
		"requested_token_type": {requestedTokenType},
		"grant_type":           {"urn:ietf:params:oauth:grant-type:token-exchange"},
		"subject_token_type":   {"jwt"},
	}

	if requestedTokenType == "urn:oci:token-type:oci-rpst" && resType != "" {
		requestData["res_type"] = []string{resType}
	}

	fc := newTokenExchangeFederationClient(tokenIssuer, domainUrl, authCode, requestData)

	return TokenExchangeConfigurationProvider{
		federationClient: fc,
		region:           common.StringToRegion(region),
	}, nil
}

// TokenExchangeConfigurationProviderFromToken returns a new configuration provider
// from a static token.
func TokenExchangeConfigurationProviderFromToken(token string, domainEndpoint string,
	clientId string, clientSecret string, region string, requestedTokenType string,
	resType string) (common.ConfigurationProvider, error) {

	issuer := StaticTokenIssuer{token: token}

	return TokenExchangeConfigurationProviderFromIssuer(issuer, domainEndpoint, clientId,
		clientSecret, region, requestedTokenType, resType)
}

func (c TokenExchangeConfigurationProvider) GetClaim(key string) (interface{}, error) {
	return c.federationClient.GetClaim(key)
}

func (c TokenExchangeConfigurationProvider) KeyID() (string, error) {
	return c.federationClient.SecurityToken()
}

func (c TokenExchangeConfigurationProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	return c.federationClient.PrivateKey()
}

// TenancyOCID provides the required receiver for the ConfigurationProvider interface
func (c TokenExchangeConfigurationProvider) TenancyOCID() (string, error) {
	claim, err := c.federationClient.GetClaim("tenant")
	if err != nil {
		return "", err
	}

	ocid, ok := claim.(string)
	if !ok {
		return "", ErrNonStringClaim
	}

	return ocid, nil
}

// UserOCID provides the required receiver for the ConfigurationProvider interface.
func (c TokenExchangeConfigurationProvider) UserOCID() (string, error) {
	claim, err := c.federationClient.GetClaim("sub")
	if err != nil {
		return "", err
	}

	ocid, ok := claim.(string)
	if !ok {
		return "", ErrNonStringClaim
	}

	return ocid, nil
}

// KeyFingerprint provides the required receiver for the ConfigurationProvider
// interface.
func (c TokenExchangeConfigurationProvider) KeyFingerprint() (string, error) {
	privateKey, err := c.PrivateRSAKey()
	if err != nil {
		return "", err
	}
	der, err := x509.MarshalPKIXPublicKey(privateKey.Public())
	if err != nil {
		return "", err
	}

	sum := md5.Sum(der)
	hexStr := hex.EncodeToString(sum[:]) // 32 hex chars

	var sb strings.Builder
	for i := 0; i < len(hexStr); i += 2 {
		if i > 0 {
			sb.WriteByte(':')
		}
		sb.WriteString(hexStr[i : i+2])
	}
	return sb.String(), nil

}

// Region provides the required receiver for the ConfigurationProvider interface.
func (c TokenExchangeConfigurationProvider) Region() (string, error) {
	r := string(c.region)
	if r == "" {
		return "", fmt.Errorf("no region assigned")
	}

	return r, nil
}

// AuthType provides the required receiver for the ConfigurationProvider interface.
func (c TokenExchangeConfigurationProvider) AuthType() (common.AuthConfig, error) {

	return common.AuthConfig{
		AuthType:         common.WorkloadIdentityFederation,
		IsFromConfigFile: false,
	}, nil
}

// tokenExchangeToken contains token and any related fields.
type tokenExchangeToken struct {
	token jwtToken
}

// String implements fmt.Stringer.
func (t tokenExchangeToken) String() string {
	return t.token.raw
}

// Valid implements the securityToken interface.
func (t tokenExchangeToken) Valid() bool {
	return !t.token.expired()
}

// GetClaim implements the ClaimHolder interface.
func (t tokenExchangeToken) GetClaim(key string) (interface{}, error) {

	// Per RFC7519 parsers should return only the lexically last member in the case
	// of duplicate claim names. We check payload first and return if claim found
	// and check header only if claim is not found in payload.
	if claim, ok := t.token.payload[key]; ok {
		return claim, nil
	}

	if claim, ok := t.token.header[key]; ok {
		return claim, nil
	}

	return nil, ErrNoSuchClaim
}
