// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_common_auth "github.com/oracle/oci-go-sdk/v65/common/auth"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

var supportedAuthSettings = []string{
	globalvar.AuthAPIKeySetting,
	globalvar.AuthInstancePrincipalSetting,
	globalvar.AuthInstancePrincipalWithCertsSetting,
	globalvar.AuthSecurityToken,
	globalvar.ResourcePrincipal,
	globalvar.AuthOKEWorkloadIdentity,
	globalvar.AuthWorkloadIdentityFederation,
}

var supportedTokenExchangeAuthSettings = []string{
	globalvar.WorkloadIdentityTokenExchangeAuthOAuthClient,
	globalvar.WorkloadIdentityTokenExchangeAuthInstancePrincipal,
}

type workloadIdentityFederationConfig struct {
	Region                    string
	WorkloadIdentityTokenPath string
	TokenExchangeDomainURL    string
	TokenExchangeAuth         string
	TokenExchangeClientID     string
	TokenExchangeClientSecret string
	RequestedTokenType        string
	SubjectTokenType          string
	TokenExchangeResourceType string
	TokenExchangeRpstExp      string
	TokenExchangePublicKey    string
}

type fileBackedTokenIssuer struct {
	tokenPath string
}

func (i fileBackedTokenIssuer) GetToken() (string, error) {
	tokenBytes, err := os.ReadFile(i.tokenPath)
	if err != nil {
		return "", err
	}

	token := strings.TrimSpace(string(tokenBytes))
	if token == "" {
		return "", fmt.Errorf("workload identity token file is empty")
	}

	return token, nil
}

func supportedAuthSettingsDescription() string {
	return "'" + strings.Join(supportedAuthSettings, "', '") + "'"
}

func tokenExchangeAuthSettingsDescription() string {
	return "'" + strings.Join(supportedTokenExchangeAuthSettings, "', '") + "'"
}

func newWorkloadIdentityFederationConfigurationProvider(config workloadIdentityFederationConfig) (oci_common.ConfigurationProvider, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	builder := oci_common_auth.TokenExchangeBuilder{
		DomainUrl:          config.TokenExchangeDomainURL,
		Region:             config.Region,
		RequestedTokenType: config.RequestedTokenType,
		SubjectTokenType:   config.SubjectTokenType,
		ResType:            config.TokenExchangeResourceType,
		RpstExp:            config.TokenExchangeRpstExp,
		PublicKey:          config.TokenExchangePublicKey,
	}

	switch strings.ToLower(config.normalizedTokenExchangeAuth()) {
	case strings.ToLower(globalvar.WorkloadIdentityTokenExchangeAuthOAuthClient):
		builder.ClientId = config.TokenExchangeClientID
		builder.ClientSecret = config.TokenExchangeClientSecret
	case strings.ToLower(globalvar.WorkloadIdentityTokenExchangeAuthInstancePrincipal):
		instancePrincipalProvider, err := instancePrincipalConfigurationProviderForRegion(config.Region)
		if err != nil {
			return nil, fmt.Errorf("can not get workload identity federation instance principal token exchange auth provider %v", err)
		}
		builder.InstancePrincipalProvider = instancePrincipalProvider
	default:
		return nil, fmt.Errorf("%s must be one of %s", globalvar.TokenExchangeAuthAttrName, tokenExchangeAuthSettingsDescription())
	}

	configProvider, err := oci_common_auth.TokenExchangeConfigurationProviderFromIssuer(
		fileBackedTokenIssuer{tokenPath: config.WorkloadIdentityTokenPath},
		builder,
	)
	if err != nil {
		return nil, fmt.Errorf("can not get workload identity federation based auth config provider %v", err)
	}

	return configProvider, nil
}

func (config workloadIdentityFederationConfig) validate() error {
	missingFields := make([]string, 0)
	requiredValues := []struct {
		field string
		value string
	}{
		{field: globalvar.RegionAttrName, value: config.Region},
		{field: globalvar.WorkloadIdentityTokenPathAttrName, value: config.WorkloadIdentityTokenPath},
		{field: globalvar.TokenExchangeDomainUrlAttrName, value: config.TokenExchangeDomainURL},
		{field: globalvar.TokenExchangeAuthAttrName, value: config.normalizedTokenExchangeAuth()},
		{field: globalvar.TokenExchangeRequestedTokenTypeAttrName, value: config.RequestedTokenType},
	}

	for _, requiredValue := range requiredValues {
		if strings.TrimSpace(requiredValue.value) == "" {
			missingFields = append(missingFields, requiredValue.field)
		}
	}

	switch strings.ToLower(config.normalizedTokenExchangeAuth()) {
	case strings.ToLower(globalvar.WorkloadIdentityTokenExchangeAuthOAuthClient):
		if strings.TrimSpace(config.TokenExchangeClientID) == "" {
			missingFields = append(missingFields, globalvar.TokenExchangeClientIdAttrName)
		}
		if strings.TrimSpace(config.TokenExchangeClientSecret) == "" {
			missingFields = append(missingFields, globalvar.TokenExchangeClientSecretAttrName)
		}
	case strings.ToLower(globalvar.WorkloadIdentityTokenExchangeAuthInstancePrincipal):
		if strings.TrimSpace(config.TokenExchangeClientID) != "" || strings.TrimSpace(config.TokenExchangeClientSecret) != "" {
			return fmt.Errorf("%s and %s can not be set when %s is '%s'", globalvar.TokenExchangeClientIdAttrName, globalvar.TokenExchangeClientSecretAttrName, globalvar.TokenExchangeAuthAttrName, globalvar.WorkloadIdentityTokenExchangeAuthInstancePrincipal)
		}
	default:
		return fmt.Errorf("%s must be one of %s", globalvar.TokenExchangeAuthAttrName, tokenExchangeAuthSettingsDescription())
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("missing required %s configuration field(s): %s", globalvar.AuthWorkloadIdentityFederation, strings.Join(missingFields, ", "))
	}

	if strings.TrimSpace(config.RequestedTokenType) == "urn:oci:token-type:oci-rpst" && strings.TrimSpace(config.TokenExchangeResourceType) == "" {
		return fmt.Errorf("%s must be set if %s is set to 'urn:oci:token-type:oci-rpst' for %s auth type", globalvar.TokenExchangeResourceTypeAttrName, globalvar.TokenExchangeRequestedTokenTypeAttrName, globalvar.AuthWorkloadIdentityFederation)
	}
	return nil
}

func (config workloadIdentityFederationConfig) normalizedTokenExchangeAuth() string {
	if strings.TrimSpace(config.TokenExchangeAuth) == "" {
		return globalvar.WorkloadIdentityTokenExchangeAuthOAuthClient
	}
	return config.TokenExchangeAuth
}

func workloadIdentityFederationConfigFromResourceData(d *schema.ResourceData) workloadIdentityFederationConfig {
	return workloadIdentityFederationConfig{
		Region:                    getStringFromResourceData(d, globalvar.RegionAttrName),
		WorkloadIdentityTokenPath: getStringFromResourceData(d, globalvar.WorkloadIdentityTokenPathAttrName),
		TokenExchangeDomainURL:    getStringFromResourceData(d, globalvar.TokenExchangeDomainUrlAttrName),
		TokenExchangeAuth:         getStringFromResourceData(d, globalvar.TokenExchangeAuthAttrName),
		TokenExchangeClientID:     getStringFromResourceData(d, globalvar.TokenExchangeClientIdAttrName),
		TokenExchangeClientSecret: getStringFromResourceData(d, globalvar.TokenExchangeClientSecretAttrName),
		RequestedTokenType:        getStringFromResourceData(d, globalvar.TokenExchangeRequestedTokenTypeAttrName),
		SubjectTokenType:          getStringFromResourceData(d, globalvar.TokenExchangeSubjectTokenTypeAttrName),
		TokenExchangeResourceType: getStringFromResourceData(d, globalvar.TokenExchangeResourceTypeAttrName),
		TokenExchangeRpstExp:      getStringFromResourceData(d, globalvar.TokenExchangeRpstExpAttrName),
		TokenExchangePublicKey:    getStringFromResourceData(d, globalvar.TokenExchangePublicKeyAttrName),
	}
}

func getStringFromResourceData(d *schema.ResourceData, attrName string) string {
	if value, ok := d.GetOkExists(attrName); ok {
		if stringValue, ok := value.(string); ok {
			return stringValue
		}
	}
	return ""
}

func workloadIdentityFederationConfigFromPluginProvider(p *ociPluginProvider) workloadIdentityFederationConfig {
	return workloadIdentityFederationConfig{
		Region:                    p.region,
		WorkloadIdentityTokenPath: p.workloadIdentityTokenPath,
		TokenExchangeDomainURL:    p.tokenExchangeDomainURL,
		TokenExchangeAuth:         p.tokenExchangeAuth,
		TokenExchangeClientID:     p.tokenExchangeClientID,
		TokenExchangeClientSecret: p.tokenExchangeClientSecret,
		RequestedTokenType:        p.tokenExchangeRequestedTokenType,
		SubjectTokenType:          p.tokenExchangeSubjectTokenType,
		TokenExchangeResourceType: p.tokenExchangeResourceType,
		TokenExchangeRpstExp:      p.tokenExchangeRpstExp,
		TokenExchangePublicKey:    p.tokenExchangePublicKey,
	}
}
