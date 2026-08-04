// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"crypto/rsa"
	"errors"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_common_auth "github.com/oracle/oci-go-sdk/v65/common/auth"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	okeWorkloadRegion = "us-phoenix-1"
	okeTargetRegion   = "ap-mumbai-1"
)

type fakeOKEWorkloadIdentityConfigurationProvider struct {
	privateKey  *rsa.PrivateKey
	authTypeErr error
}

func (p *fakeOKEWorkloadIdentityConfigurationProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	return p.privateKey, nil
}

func (p *fakeOKEWorkloadIdentityConfigurationProvider) KeyID() (string, error) {
	return "test-key-id", nil
}

func (p *fakeOKEWorkloadIdentityConfigurationProvider) TenancyOCID() (string, error) {
	return "test-tenancy", nil
}

func (p *fakeOKEWorkloadIdentityConfigurationProvider) UserOCID() (string, error) {
	return "test-user", nil
}

func (p *fakeOKEWorkloadIdentityConfigurationProvider) KeyFingerprint() (string, error) {
	return "test-fingerprint", nil
}

func (p *fakeOKEWorkloadIdentityConfigurationProvider) Region() (string, error) {
	return okeWorkloadRegion, nil
}

func (p *fakeOKEWorkloadIdentityConfigurationProvider) AuthType() (oci_common.AuthConfig, error) {
	return oci_common.AuthConfig{AuthType: oci_common.UnknownAuthenticationType}, p.authTypeErr
}

func (p *fakeOKEWorkloadIdentityConfigurationProvider) GetClaim(claim string) (interface{}, error) {
	return "claim:" + claim, nil
}

func (p *fakeOKEWorkloadIdentityConfigurationProvider) Refreshable() bool {
	return true
}

var _ oci_common_auth.ConfigurationProviderWithClaimAccess = (*fakeOKEWorkloadIdentityConfigurationProvider)(nil)
var _ oci_common.RefreshableConfigurationProvider = (*fakeOKEWorkloadIdentityConfigurationProvider)(nil)

func stubOKEWorkloadIdentityConfigurationProvider(t *testing.T, configProvider oci_common_auth.ConfigurationProviderWithClaimAccess, err error) {
	t.Helper()
	original := okeWorkloadIdentityConfigurationProviderFn
	okeWorkloadIdentityConfigurationProviderFn = func() (oci_common_auth.ConfigurationProviderWithClaimAccess, error) {
		return configProvider, err
	}
	t.Cleanup(func() {
		okeWorkloadIdentityConfigurationProviderFn = original
	})
}

func newTestOracleClients() *tf_client.OracleClients {
	return &tf_client.OracleClients{
		SdkClientMap:  make(map[string]interface{}, len(tf_client.OracleClientRegistrationsVar.RegisteredClients)),
		Configuration: make(map[string]string),
	}
}

func TestUnitOKEWorkloadIdentityRegionOverrideDelegatesAuthentication(t *testing.T) {
	authTypeErr := errors.New("test auth type error")
	privateKey := &rsa.PrivateKey{}
	baseProvider := &fakeOKEWorkloadIdentityConfigurationProvider{
		privateKey:  privateKey,
		authTypeErr: authTypeErr,
	}
	stubOKEWorkloadIdentityConfigurationProvider(t, baseProvider, nil)

	configProvider, err := newOKEWorkloadIdentityConfigurationProvider(okeTargetRegion)
	require.NoError(t, err)

	region, err := configProvider.Region()
	require.NoError(t, err)
	assert.Equal(t, okeTargetRegion, region)

	keyID, err := configProvider.KeyID()
	require.NoError(t, err)
	assert.Equal(t, "test-key-id", keyID)

	tenancyOCID, err := configProvider.TenancyOCID()
	require.NoError(t, err)
	assert.Equal(t, "test-tenancy", tenancyOCID)

	userOCID, err := configProvider.UserOCID()
	require.NoError(t, err)
	assert.Equal(t, "test-user", userOCID)

	fingerprint, err := configProvider.KeyFingerprint()
	require.NoError(t, err)
	assert.Equal(t, "test-fingerprint", fingerprint)

	actualPrivateKey, err := configProvider.PrivateRSAKey()
	require.NoError(t, err)
	assert.Same(t, privateKey, actualPrivateKey)

	claim, err := configProvider.GetClaim("res_tenant")
	require.NoError(t, err)
	assert.Equal(t, "claim:res_tenant", claim)

	_, err = configProvider.AuthType()
	assert.ErrorIs(t, err, authTypeErr)

	refreshableProvider, ok := configProvider.(oci_common.RefreshableConfigurationProvider)
	require.True(t, ok)
	assert.True(t, refreshableProvider.Refreshable())
}

func TestUnitOKEWorkloadIdentityNoRegionPreservesSDKProvider(t *testing.T) {
	baseProvider := &fakeOKEWorkloadIdentityConfigurationProvider{}
	stubOKEWorkloadIdentityConfigurationProvider(t, baseProvider, nil)

	configProvider, err := newOKEWorkloadIdentityConfigurationProvider("")
	require.NoError(t, err)
	assert.Same(t, baseProvider, configProvider)

	region, err := configProvider.Region()
	require.NoError(t, err)
	assert.Equal(t, okeWorkloadRegion, region)
}

func TestUnitOKEWorkloadIdentityConstructorError(t *testing.T) {
	constructorErr := errors.New("test constructor error")
	stubOKEWorkloadIdentityConfigurationProvider(t, nil, constructorErr)

	configProvider, err := newOKEWorkloadIdentityConfigurationProvider(okeTargetRegion)
	assert.Nil(t, configProvider)
	assert.ErrorIs(t, err, constructorErr)
}

func TestUnitOKEWorkloadIdentityProviderConstructorError(t *testing.T) {
	constructorErr := errors.New("test constructor error")

	t.Run("SDKv2", func(t *testing.T) {
		stubOKEWorkloadIdentityConfigurationProvider(t, nil, constructorErr)

		resource := &schema.Resource{Schema: SchemaMap()}
		data := resource.Data(nil)
		require.NoError(t, data.Set(globalvar.AuthAttrName, globalvar.AuthOKEWorkloadIdentity))
		require.NoError(t, data.Set(globalvar.RegionAttrName, okeTargetRegion))

		configProvider, err := GetSdkConfigProvider(data, newTestOracleClients())
		assert.Nil(t, configProvider)
		assert.ErrorIs(t, err, constructorErr)
	})

	t.Run("framework", func(t *testing.T) {
		stubOKEWorkloadIdentityConfigurationProvider(t, nil, constructorErr)

		provider := &ociPluginProvider{
			auth:   globalvar.AuthOKEWorkloadIdentity,
			region: okeTargetRegion,
		}
		configProvider, err := provider._GetSdkConfigProvider(newTestOracleClients())
		assert.Nil(t, configProvider)
		assert.ErrorIs(t, err, constructorErr)
	})
}

func TestUnitOKEWorkloadIdentitySDKv2RegionSelection(t *testing.T) {
	tests := []struct {
		name           string
		targetRegion   string
		expectedRegion string
	}{
		{name: "configured region overrides workload region", targetRegion: okeTargetRegion, expectedRegion: okeTargetRegion},
		{name: "workload region remains the fallback", expectedRegion: okeWorkloadRegion},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stubOKEWorkloadIdentityConfigurationProvider(t, &fakeOKEWorkloadIdentityConfigurationProvider{}, nil)

			resource := &schema.Resource{Schema: SchemaMap()}
			data := resource.Data(nil)
			require.NoError(t, data.Set(globalvar.AuthAttrName, globalvar.AuthOKEWorkloadIdentity))
			if test.targetRegion != "" {
				require.NoError(t, data.Set(globalvar.RegionAttrName, test.targetRegion))
			}

			configProvider, err := GetSdkConfigProvider(data, newTestOracleClients())
			require.NoError(t, err)
			region, err := configProvider.Region()
			require.NoError(t, err)
			assert.Equal(t, test.expectedRegion, region)
		})
	}
}

func TestUnitOKEWorkloadIdentityFrameworkRegionSelection(t *testing.T) {
	tests := []struct {
		name           string
		targetRegion   string
		expectedRegion string
	}{
		{name: "configured region overrides workload region", targetRegion: okeTargetRegion, expectedRegion: okeTargetRegion},
		{name: "workload region remains the fallback", expectedRegion: okeWorkloadRegion},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stubOKEWorkloadIdentityConfigurationProvider(t, &fakeOKEWorkloadIdentityConfigurationProvider{}, nil)

			provider := &ociPluginProvider{
				auth:   globalvar.AuthOKEWorkloadIdentity,
				region: test.targetRegion,
			}
			configProvider, err := provider._GetSdkConfigProvider(newTestOracleClients())
			require.NoError(t, err)
			region, err := configProvider.Region()
			require.NoError(t, err)
			assert.Equal(t, test.expectedRegion, region)
		})
	}
}
