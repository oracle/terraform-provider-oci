// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	frameworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	frameworkschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	sdkschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitWorkloadIdentityFederationLegacySchema(t *testing.T) {
	schemaMap := SchemaMap()

	warnings, errors := schemaMap[globalvar.AuthAttrName].ValidateFunc(globalvar.AuthWorkloadIdentityFederation, globalvar.AuthAttrName)
	assert.Empty(t, warnings)
	assert.Empty(t, errors)

	warnings, errors = schemaMap[globalvar.TokenExchangeAuthAttrName].ValidateFunc(globalvar.WorkloadIdentityTokenExchangeAuthOAuthClient, globalvar.TokenExchangeAuthAttrName)
	assert.Empty(t, warnings)
	assert.Empty(t, errors)

	warnings, errors = schemaMap[globalvar.TokenExchangeAuthAttrName].ValidateFunc(globalvar.WorkloadIdentityTokenExchangeAuthInstancePrincipal, globalvar.TokenExchangeAuthAttrName)
	assert.Empty(t, warnings)
	assert.Empty(t, errors)

	defaultValue, err := schemaMap[globalvar.TokenExchangeAuthAttrName].DefaultFunc()
	require.NoError(t, err)
	assert.Equal(t, globalvar.WorkloadIdentityTokenExchangeAuthOAuthClient, defaultValue)

	for _, attrName := range workloadIdentityFederationAttrNames() {
		require.Contains(t, schemaMap, attrName)
	}

	assert.True(t, schemaMap[globalvar.TokenExchangeClientSecretAttrName].Sensitive)
}

func TestUnitWorkloadIdentityFederationFrameworkSchema(t *testing.T) {
	var resp frameworkprovider.SchemaResponse
	New().Schema(context.Background(), frameworkprovider.SchemaRequest{}, &resp)

	for _, attrName := range workloadIdentityFederationAttrNames() {
		require.Contains(t, resp.Schema.Attributes, attrName)
	}

	secretAttr, ok := resp.Schema.Attributes[globalvar.TokenExchangeClientSecretAttrName].(frameworkschema.StringAttribute)
	require.True(t, ok)
	assert.True(t, secretAttr.Sensitive)
}

func TestUnitFileBackedTokenIssuerReadsLatestTrimmedToken(t *testing.T) {
	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte(" first-token\n"), 0600))

	issuer := fileBackedTokenIssuer{tokenPath: tokenPath}
	token, err := issuer.GetToken()
	require.NoError(t, err)
	assert.Equal(t, "first-token", token)

	require.NoError(t, os.WriteFile(tokenPath, []byte("\nsecond-token\t"), 0600))
	token, err = issuer.GetToken()
	require.NoError(t, err)
	assert.Equal(t, "second-token", token)

	require.NoError(t, os.WriteFile(tokenPath, []byte(" \n\t"), 0600))
	_, err = issuer.GetToken()
	require.ErrorContains(t, err, "workload identity token file is empty")
}

func TestUnitWorkloadIdentityFederationValidationRequiresFields(t *testing.T) {
	_, err := newWorkloadIdentityFederationConfigurationProvider(workloadIdentityFederationConfig{})
	require.Error(t, err)

	for _, attrName := range []string{
		globalvar.RegionAttrName,
		globalvar.WorkloadIdentityTokenPathAttrName,
		globalvar.TokenExchangeDomainUrlAttrName,
		globalvar.TokenExchangeClientIdAttrName,
		globalvar.TokenExchangeClientSecretAttrName,
		globalvar.TokenExchangeRequestedTokenTypeAttrName,
	} {
		assert.Contains(t, err.Error(), attrName)
	}
}

func TestUnitWorkloadIdentityFederationValidationRejectsInvalidTokenExchangeAuth(t *testing.T) {
	config := minimalWorkloadIdentityFederationConfig()
	config.TokenExchangeAuth = "invalid"

	_, err := newWorkloadIdentityFederationConfigurationProvider(config)
	require.ErrorContains(t, err, globalvar.TokenExchangeAuthAttrName)
	require.ErrorContains(t, err, globalvar.WorkloadIdentityTokenExchangeAuthOAuthClient)
	require.ErrorContains(t, err, globalvar.WorkloadIdentityTokenExchangeAuthInstancePrincipal)
}

func TestUnitWorkloadIdentityFederationValidationRejectsClientCredentialsForInstancePrincipal(t *testing.T) {
	testCases := []struct {
		name         string
		clientID     string
		clientSecret string
	}{
		{
			name:     "client ID set",
			clientID: "client-id",
		},
		{
			name:         "client secret set",
			clientSecret: "client-secret",
		},
		{
			name:         "client ID and secret set",
			clientID:     "client-id",
			clientSecret: "client-secret",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			config := minimalWorkloadIdentityFederationConfig()
			config.TokenExchangeAuth = globalvar.WorkloadIdentityTokenExchangeAuthInstancePrincipal
			config.TokenExchangeClientID = testCase.clientID
			config.TokenExchangeClientSecret = testCase.clientSecret

			_, err := newWorkloadIdentityFederationConfigurationProvider(config)
			require.ErrorContains(t, err, globalvar.TokenExchangeClientIdAttrName)
			require.ErrorContains(t, err, globalvar.TokenExchangeClientSecretAttrName)
			require.ErrorContains(t, err, globalvar.TokenExchangeAuthAttrName)
			require.ErrorContains(t, err, globalvar.WorkloadIdentityTokenExchangeAuthInstancePrincipal)
		})
	}
}

func TestUnitWorkloadIdentityFederationLegacyConfigProvider(t *testing.T) {
	r := &sdkschema.Resource{Schema: SchemaMap()}
	d := r.Data(nil)
	setWorkloadIdentityFederationResourceData(t, d, minimalWorkloadIdentityFederationConfig())

	configProviders, err := getConfigProviders(d, strings.ToLower(globalvar.AuthWorkloadIdentityFederation))
	require.NoError(t, err)
	require.Len(t, configProviders, 1)
	assertWorkloadIdentityFederationConfigProvider(t, configProviders[0], "us-ashburn-1")
}

func TestUnitWorkloadIdentityFederationInstancePrincipalTokenExchangeAuth(t *testing.T) {
	originalFactory := instancePrincipalConfigurationProviderForRegion
	t.Cleanup(func() {
		instancePrincipalConfigurationProviderForRegion = originalFactory
	})

	factoryCalled := false
	instancePrincipalConfigurationProviderForRegion = func(region string) (oci_common.ConfigurationProvider, error) {
		factoryCalled = true
		assert.Equal(t, "us-ashburn-1", region)
		return oci_common.NewRawConfigurationProvider(testTenancyOCID, testUserOCID, "us-ashburn-1", testKeyFingerPrint, testPrivateKey, nil), nil
	}

	config := minimalWorkloadIdentityFederationConfig()
	config.TokenExchangeAuth = globalvar.WorkloadIdentityTokenExchangeAuthInstancePrincipal
	config.TokenExchangeClientID = ""
	config.TokenExchangeClientSecret = ""

	configProvider, err := newWorkloadIdentityFederationConfigurationProvider(config)
	require.NoError(t, err)
	assert.True(t, factoryCalled)
	assertWorkloadIdentityFederationConfigProvider(t, configProvider, "us-ashburn-1")
}

func TestUnitWorkloadIdentityFederationFrameworkConfigProvider(t *testing.T) {
	config := minimalWorkloadIdentityFederationConfig()
	p := &ociPluginProvider{
		auth:                            globalvar.AuthWorkloadIdentityFederation,
		region:                          config.Region,
		workloadIdentityTokenPath:       config.WorkloadIdentityTokenPath,
		tokenExchangeDomainURL:          config.TokenExchangeDomainURL,
		tokenExchangeAuth:               config.TokenExchangeAuth,
		tokenExchangeClientID:           config.TokenExchangeClientID,
		tokenExchangeClientSecret:       config.TokenExchangeClientSecret,
		tokenExchangeRequestedTokenType: config.RequestedTokenType,
		tokenExchangeSubjectTokenType:   config.SubjectTokenType,
		tokenExchangeResourceType:       config.TokenExchangeResourceType,
		tokenExchangeRpstExp:            config.TokenExchangeRpstExp,
		tokenExchangePublicKey:          config.TokenExchangePublicKey,
	}

	configProviders, err := p._getConfigProviders()
	require.NoError(t, err)
	require.Len(t, configProviders, 1)
	assertWorkloadIdentityFederationConfigProvider(t, configProviders[0], "us-ashburn-1")
}

func TestUnitWorkloadIdentityFederationFrameworkSetDefaults(t *testing.T) {
	t.Setenv(tfVarName(globalvar.WorkloadIdentityTokenPathAttrName), "/env/token")
	t.Setenv(tfVarName(globalvar.TokenExchangeDomainUrlAttrName), "https://identity.example.com")
	t.Setenv(tfVarName(globalvar.TokenExchangeAuthAttrName), globalvar.WorkloadIdentityTokenExchangeAuthInstancePrincipal)
	t.Setenv(tfVarName(globalvar.TokenExchangeClientIdAttrName), "env-client-id")
	t.Setenv(tfVarName(globalvar.TokenExchangeClientSecretAttrName), "env-client-secret")
	t.Setenv(tfVarName(globalvar.TokenExchangeRequestedTokenTypeAttrName), "urn:oci:token-type:oci-rpst")
	t.Setenv(tfVarName(globalvar.TokenExchangeSubjectTokenTypeAttrName), "urn:ietf:params:oauth:token-type:jwt")
	t.Setenv(tfVarName(globalvar.TokenExchangeResourceTypeAttrName), "env-resource")

	p := &ociPluginProvider{}
	p.SetDefaults(&ociProviderModel{})

	assert.Equal(t, "/env/token", p.workloadIdentityTokenPath)
	assert.Equal(t, "https://identity.example.com", p.tokenExchangeDomainURL)
	assert.Equal(t, globalvar.WorkloadIdentityTokenExchangeAuthInstancePrincipal, p.tokenExchangeAuth)
	assert.Equal(t, "env-client-id", p.tokenExchangeClientID)
	assert.Equal(t, "env-client-secret", p.tokenExchangeClientSecret)
	assert.Equal(t, "urn:oci:token-type:oci-rpst", p.tokenExchangeRequestedTokenType)
	assert.Equal(t, "urn:ietf:params:oauth:token-type:jwt", p.tokenExchangeSubjectTokenType)
	assert.Equal(t, "env-resource", p.tokenExchangeResourceType)
}

func TestUnitWorkloadIdentityFederationFrameworkSetDefaultsTokenExchangeAuth(t *testing.T) {
	p := &ociPluginProvider{}
	p.SetDefaults(&ociProviderModel{})

	assert.Equal(t, globalvar.WorkloadIdentityTokenExchangeAuthOAuthClient, p.tokenExchangeAuth)
}

func minimalWorkloadIdentityFederationConfig() workloadIdentityFederationConfig {
	return workloadIdentityFederationConfig{
		Region:                    "us-ashburn-1",
		WorkloadIdentityTokenPath: "/var/run/secrets/tokens/oci",
		TokenExchangeDomainURL:    "https://identity.example.com",
		TokenExchangeAuth:         globalvar.WorkloadIdentityTokenExchangeAuthOAuthClient,
		TokenExchangeClientID:     "client-id",
		TokenExchangeClientSecret: "client-secret",
		RequestedTokenType:        "urn:oci:token-type:oci-rpst",
		SubjectTokenType:          "urn:ietf:params:oauth:token-type:jwt",
		TokenExchangeResourceType: "resource-type",
		TokenExchangeRpstExp:      "3600",
		TokenExchangePublicKey:    "public-key",
	}
}

func setWorkloadIdentityFederationResourceData(t *testing.T, d *sdkschema.ResourceData, config workloadIdentityFederationConfig) {
	t.Helper()

	require.NoError(t, d.Set(globalvar.AuthAttrName, globalvar.AuthWorkloadIdentityFederation))
	require.NoError(t, d.Set(globalvar.RegionAttrName, config.Region))
	require.NoError(t, d.Set(globalvar.WorkloadIdentityTokenPathAttrName, config.WorkloadIdentityTokenPath))
	require.NoError(t, d.Set(globalvar.TokenExchangeDomainUrlAttrName, config.TokenExchangeDomainURL))
	require.NoError(t, d.Set(globalvar.TokenExchangeAuthAttrName, config.TokenExchangeAuth))
	require.NoError(t, d.Set(globalvar.TokenExchangeClientIdAttrName, config.TokenExchangeClientID))
	require.NoError(t, d.Set(globalvar.TokenExchangeClientSecretAttrName, config.TokenExchangeClientSecret))
	require.NoError(t, d.Set(globalvar.TokenExchangeRequestedTokenTypeAttrName, config.RequestedTokenType))
	require.NoError(t, d.Set(globalvar.TokenExchangeSubjectTokenTypeAttrName, config.SubjectTokenType))
	require.NoError(t, d.Set(globalvar.TokenExchangeResourceTypeAttrName, config.TokenExchangeResourceType))
	require.NoError(t, d.Set(globalvar.TokenExchangeRpstExpAttrName, config.TokenExchangeRpstExp))
	require.NoError(t, d.Set(globalvar.TokenExchangePublicKeyAttrName, config.TokenExchangePublicKey))
}

func assertWorkloadIdentityFederationConfigProvider(t *testing.T, configProvider oci_common.ConfigurationProvider, expectedRegion string) {
	t.Helper()

	authConfig, err := configProvider.AuthType()
	require.NoError(t, err)
	assert.Equal(t, oci_common.WorkloadIdentityFederation, authConfig.AuthType)

	region, err := configProvider.Region()
	require.NoError(t, err)
	assert.Equal(t, expectedRegion, region)
}

func workloadIdentityFederationAttrNames() []string {
	return []string{
		globalvar.WorkloadIdentityTokenPathAttrName,
		globalvar.TokenExchangeDomainUrlAttrName,
		globalvar.TokenExchangeAuthAttrName,
		globalvar.TokenExchangeClientIdAttrName,
		globalvar.TokenExchangeClientSecretAttrName,
		globalvar.TokenExchangeRequestedTokenTypeAttrName,
		globalvar.TokenExchangeSubjectTokenTypeAttrName,
		globalvar.TokenExchangeResourceTypeAttrName,
		globalvar.TokenExchangeRpstExpAttrName,
		globalvar.TokenExchangePublicKeyAttrName,
	}
}
