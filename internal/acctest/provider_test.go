// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package acctest

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"

	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/provider"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

// This test runs the Provider sanity checks.
// issue-routing-tag: terraform/default
func TestUnitProvider(t *testing.T) {
	// Real client for the sanity check. Makes this more of an acceptance test.
	client := &tf_client.OracleClients{}
	testProvider := &schema.Provider{
		DataSourcesMap: provider.DataSourcesMap(),
		Schema:         provider.SchemaMap(),
		ResourcesMap:   provider.ResourcesMap(),
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		},
	}
	if err := testProvider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

// issue-routing-tag: terraform/default
func TestUnitProviderConfig(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip ProviderConfigTest in HttpReplay mode.")
	}
	if os.Getenv("TF_HOME_OVERRIDE") == "" {
		t.Skip("This run requires you to set TF_HOME_OVERRIDE")
	}
	ProviderConfigTest(t, true, true, globalvar.AuthAPIKeySetting, "", nil)              // ApiKey with required fields + disable auto-retries
	ProviderConfigTest(t, false, true, globalvar.AuthAPIKeySetting, "", nil)             // ApiKey without required fields
	ProviderConfigTest(t, false, false, globalvar.AuthInstancePrincipalSetting, "", nil) // InstancePrincipal
	ProviderConfigTest(t, true, false, "invalid-auth-setting", "", nil)                  // Invalid auth + disable auto-retries
	configFile, keyFile, err := WriteConfigFile()
	assert.Nil(t, err)
	ProviderConfigTest(t, true, true, globalvar.AuthAPIKeySetting, "DEFAULT", nil)              // ApiKey with required fields + disable auto-retries
	ProviderConfigTest(t, false, true, globalvar.AuthAPIKeySetting, "DEFAULT", nil)             // ApiKey without required fields
	ProviderConfigTest(t, false, false, globalvar.AuthInstancePrincipalSetting, "DEFAULT", nil) // InstancePrincipal
	ProviderConfigTest(t, true, false, "invalid-auth-setting", "DEFAULT", nil)                  // Invalid auth + disable auto-retries
	ProviderConfigTest(t, false, false, globalvar.AuthAPIKeySetting, "PROFILE1", nil)           // correct profileName
	ProviderConfigTest(t, false, false, globalvar.AuthAPIKeySetting, "wrongProfile", nil)       // Invalid profileName
	//acctest.ProviderConfigTest(t, false, false, globalvar.AuthAPIKeySetting, "PROFILE2", nil)           // correct profileName with mix and match
	ProviderConfigTest(t, false, false, globalvar.AuthAPIKeySetting, "PROFILE3", nil) // correct profileName with mix and match & env
	defer func() {
		_ = utils.RemoveFile(configFile)
	}()
	defer func() {
		_ = utils.RemoveFile(keyFile)
	}()
	defer func() {
		_ = os.RemoveAll(path.Join(utils.GetHomeFolder(), globalvar.DefaultConfigDirName))
	}()
}

// ensure the http client is configured with the expected settings
// issue-routing-tag: terraform/default
func TestUnitBuildHttpClient(t *testing.T) {
	client := provider.BuildHttpClient()
	assert.Equal(t, time.Duration(globalvar.DefaultRequestTimeout), client.Timeout)

	tr := client.Transport.(*http.Transport)
	assert.NotNil(t, tr.TLSClientConfig)
	assert.Equal(t, uint16(tls.VersionTLS12), tr.TLSClientConfig.MinVersion, "expected min tls 1.2")
	assert.NotNil(t, tr.Proxy, "expected http.ProxyFromEnvironment fn")
}

// ensure desired http client settings are not removed when sdk clients are configured
// issue-routing-tag: terraform/default
func TestUnitBuildClientConfigureFn(t *testing.T) {
	configProvider := oci_common.DefaultConfigProvider()
	httpClient := provider.BuildHttpClient()
	configureClientFn, err := provider.BuildConfigureClientFn(configProvider, httpClient)
	assert.NoError(t, err)

	baseClient := &oci_common.BaseClient{}
	err = configureClientFn(baseClient)
	assert.NoError(t, err)

	tr := httpClient.Transport.(*http.Transport)
	assert.NotNil(t, tr.TLSClientConfig)
	assert.Equal(t, uint16(tls.VersionTLS12), tr.TLSClientConfig.MinVersion, "expected min tls 1.2")
	assert.NotNil(t, tr.Proxy, "expected http.ProxyFromEnvironment fn")
}

// ensure custom certs can be added to the cert pool and expected http client settings are preserved
// issue-routing-tag: terraform/default
func TestUnitBuildClientConfigureFn_withCustomCert(t *testing.T) {
	ca := "-----BEGIN CERTIFICATE-----\nMIIC9jCCAd4CCQD2rPUVJETHGzANBgkqhkiG9w0BAQsFADA9MQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTAeFw0xOTAxMTcyMjU4MDVaFw0yMTAxMTYyMjU4MDVaMD0xCzAJBgNVBAYTAlVT\nMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGT3JhY2xl\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA30+wt7OlUB/YpmWbTRkx\nnLG0lKWiV+oupNKj8luXmC5jvOFTUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU\n+DWVV2So2B/obYxpiiyWF2tcF/cYi1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oC\nMQ2985/MTdCXONgnbmePU64GrJwfvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOO\njLKRM68KXC5us4879IrSA77NQr1KwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6\nytM66P/1CTpk1YpbI4gqiG0HBbuXG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc\n2wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAuI53m8Va6EafDi6GQdQrzNNQFCAVQ\nxIABAB0uaSYCs3H+pqTktHzOrOluSUEogXRl0UU5/OuvxAz4idA4cfBdId4i7AcY\nqZsBjA/xqH/rxR3pcgfaGyxQzrUsJFf0ZwnzqYJs7fUvuatHJYi/cRBxrKR2+4Oj\nlUbb9TSmezlzHK5CaD5XzN+lZqbsSvN3OQbOryJCbtjZVQFGZ1SmL6OLrwpbBKuP\nn2ob+gaP57YSzO3zk1NDXMlQPHRsdSOqocyKx8y+7J0g6MqPvBzIe+wI3QW85MQY\nj1/IHmj84LNGp7pHCyiYx/oI+00gRch04H2pJv0TP3sAQ37gplBwDrUo\n-----END CERTIFICATE-----"
	tempCert, err := ioutil.TempFile("", "caCert*.pem")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tempCert.Name())

	if _, err := tempCert.Write([]byte(ca)); err != nil {
		t.Error(err)
	}
	if err := tempCert.Close(); err != nil {
		t.Error(err)
	}

	prevEnvVar, hadPreviousEnvVar := os.LookupEnv(globalvar.CustomCertLocationEnv)
	if hadPreviousEnvVar {
		defer os.Setenv(globalvar.CustomCertLocationEnv, prevEnvVar)
	} else {
		defer os.Unsetenv(globalvar.CustomCertLocationEnv)
	}

	os.Setenv(globalvar.CustomCertLocationEnv, tempCert.Name())
	assert.Equal(t, tempCert.Name(), utils.GetEnvSettingWithBlankDefault(globalvar.CustomCertLocationEnv))
	configProvider := oci_common.DefaultConfigProvider()
	httpClient := provider.BuildHttpClient()
	configureClientFn, err := provider.BuildConfigureClientFn(configProvider, httpClient)
	assert.NoError(t, err)

	baseClient := &oci_common.BaseClient{}
	err = configureClientFn(baseClient)
	assert.NoError(t, err)

	tr := httpClient.Transport.(*http.Transport)
	assert.NotNil(t, tr.TLSClientConfig)
	assert.Equal(t, uint16(tls.VersionTLS12), tr.TLSClientConfig.MinVersion, "expected min tls 1.2")
	assert.NotNil(t, tr.Proxy, "expected http.ProxyFromEnvironment fn")
	assert.NotNil(t, tr.TLSClientConfig.RootCAs)
}

// ensure local certs can be admitted
// issue-routing-tag: terraform/default
func TestUnitBuildClientConfigureFn_acceptLocalCerts(t *testing.T) {
	prevEnvVar, hadPreviousEnvVar := os.LookupEnv(globalvar.AcceptLocalCerts)
	if hadPreviousEnvVar {
		defer os.Setenv(globalvar.AcceptLocalCerts, prevEnvVar)
	} else {
		defer os.Unsetenv(globalvar.AcceptLocalCerts)
	}

	// ensure disabled by default - no env var
	os.Unsetenv(globalvar.AcceptLocalCerts)
	assert.Empty(t, utils.GetEnvSettingWithBlankDefault(globalvar.AcceptLocalCerts))
	configProvider := oci_common.DefaultConfigProvider()
	httpClient := provider.BuildHttpClient()
	configureClientFn, _ := provider.BuildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr := httpClient.Transport.(*http.Transport)
	assert.NotNil(t, tr.Proxy, "expected http.ProxyFromEnvironment fn")
	assert.Equal(t, uint16(tls.VersionTLS12), tr.TLSClientConfig.MinVersion, "expected min tls 1.2")
	assert.False(t, tr.TLSClientConfig.InsecureSkipVerify)

	// ensure disabled by default - env var with empty string
	os.Setenv(globalvar.AcceptLocalCerts, "")
	configProvider = oci_common.DefaultConfigProvider()
	httpClient = provider.BuildHttpClient()
	configureClientFn, _ = provider.BuildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr = httpClient.Transport.(*http.Transport)
	assert.False(t, tr.TLSClientConfig.InsecureSkipVerify)

	// ensure disabled by default - not parsable boolean string
	os.Setenv(globalvar.AcceptLocalCerts, "ftarlusee")
	configProvider = oci_common.DefaultConfigProvider()
	httpClient = provider.BuildHttpClient()
	configureClientFn, _ = provider.BuildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr = httpClient.Transport.(*http.Transport)
	assert.False(t, tr.TLSClientConfig.InsecureSkipVerify)

	// ensure explicitly disabled - env var set to false
	os.Setenv(globalvar.AcceptLocalCerts, "false")
	configProvider = oci_common.DefaultConfigProvider()
	httpClient = provider.BuildHttpClient()
	configureClientFn, _ = provider.BuildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr = httpClient.Transport.(*http.Transport)
	assert.False(t, tr.TLSClientConfig.InsecureSkipVerify)

	// ensure explicitly enabled - env var set to true
	os.Setenv(globalvar.AcceptLocalCerts, "true")
	configProvider = oci_common.DefaultConfigProvider()
	httpClient = provider.BuildHttpClient()
	configureClientFn, _ = provider.BuildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr = httpClient.Transport.(*http.Transport)
	assert.True(t, tr.TLSClientConfig.InsecureSkipVerify)

	// verify assumption that "1" will also coerce to true
	os.Setenv(globalvar.AcceptLocalCerts, "1")
	configProvider = oci_common.DefaultConfigProvider()
	httpClient = provider.BuildHttpClient()
	configureClientFn, _ = provider.BuildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr = httpClient.Transport.(*http.Transport)
	assert.True(t, tr.TLSClientConfig.InsecureSkipVerify)
}

// ensure a custom domain can be targeted and expected http client settings are preserved
// issue-routing-tag: terraform/default
func TestUnitBuildClientConfigureFn_withDomainNameOverride(t *testing.T) {

	prevEnvVar, hadPreviousEnvVar := os.LookupEnv(globalvar.DomainNameOverrideEnv)
	if hadPreviousEnvVar {
		defer os.Setenv(globalvar.DomainNameOverrideEnv, prevEnvVar)
	} else {
		defer os.Unsetenv(globalvar.DomainNameOverrideEnv)
	}

	if hadPreviousEnvVar {
		defer os.Setenv(globalvar.HasCorrectDomainNameEnv, prevEnvVar)
	} else {
		defer os.Unsetenv(globalvar.HasCorrectDomainNameEnv)
	}

	os.Setenv(globalvar.DomainNameOverrideEnv, "0r4-c10ud.com")
	assert.Equal(t, "0r4-c10ud.com", utils.GetEnvSettingWithBlankDefault(globalvar.DomainNameOverrideEnv))
	configProvider := oci_common.DefaultConfigProvider()
	httpClient := provider.BuildHttpClient()
	configureClientFn, err := provider.BuildConfigureClientFn(configProvider, httpClient)
	assert.NoError(t, err)

	baseClient := &oci_common.BaseClient{}
	baseClient.Host = "https://svc.region.oraclecloud.com"
	err = configureClientFn(baseClient)
	assert.NoError(t, err)

	// verify transport settings are unchanged
	tr := httpClient.Transport.(*http.Transport)
	assert.NotNil(t, tr.TLSClientConfig)
	assert.Equal(t, uint16(tls.VersionTLS12), tr.TLSClientConfig.MinVersion, "expected min tls 1.2")
	assert.NotNil(t, tr.Proxy, "expected http.ProxyFromEnvironment fn")
	assert.Nil(t, tr.TLSClientConfig.RootCAs)

	// verify url has expected domain
	assert.Equal(t, `https://svc.region.0r4-c10ud.com`, baseClient.Host)

	// verify subdomains are preserved
	baseClient = &oci_common.BaseClient{}
	baseClient.Host = "avnzdivwaadfa-management.kms.us-phoenix-1.oraclecloud.com"
	err = configureClientFn(baseClient)
	assert.NoError(t, err)
	assert.Equal(t, `avnzdivwaadfa-management.kms.us-phoenix-1.0r4-c10ud.com`, baseClient.Host)

	// verify non-match preserves original url
	baseClient = &oci_common.BaseClient{}
	baseClient.Host = "DUMMY_ENDPOINT"
	err = configureClientFn(baseClient)
	assert.NoError(t, err)
	assert.Equal(t, `DUMMY_ENDPOINT`, baseClient.Host)
}

// ensure a custom domain that has already override with more than 2 dots can be targeted and expected http client settings are preserved
// issue-routing-tag: terraform/default
func TestUnitBuildClientConfigureFn_withDomainNameOverrideAndCorrectDomainName(t *testing.T) {

	prevEnvVar, hadPreviousEnvVar := os.LookupEnv(globalvar.DomainNameOverrideEnv)
	if hadPreviousEnvVar {
		defer os.Setenv(globalvar.DomainNameOverrideEnv, prevEnvVar)
	} else {
		defer os.Unsetenv(globalvar.DomainNameOverrideEnv)
	}

	if hadPreviousEnvVar {
		defer os.Setenv(globalvar.HasCorrectDomainNameEnv, prevEnvVar)
	} else {
		defer os.Unsetenv(globalvar.HasCorrectDomainNameEnv)
	}

	os.Setenv(globalvar.DomainNameOverrideEnv, "oc.0r4-c10ud.com")
	os.Setenv(globalvar.HasCorrectDomainNameEnv, "oc.0r4-c10ud.com")
	assert.Equal(t, "oc.0r4-c10ud.com", utils.GetEnvSettingWithBlankDefault(globalvar.DomainNameOverrideEnv))
	assert.Equal(t, "oc.0r4-c10ud.com", utils.GetEnvSettingWithBlankDefault(globalvar.HasCorrectDomainNameEnv))
	configProvider := oci_common.DefaultConfigProvider()
	httpClient := provider.BuildHttpClient()
	configureClientFn, err := provider.BuildConfigureClientFn(configProvider, httpClient)
	assert.NoError(t, err)

	baseClient := &oci_common.BaseClient{}
	baseClient.Host = "https://svc.region.oc.0r4-c10ud.com"
	err = configureClientFn(baseClient)
	assert.NoError(t, err)

	// verify transport settings are unchanged
	tr := httpClient.Transport.(*http.Transport)
	assert.NotNil(t, tr.TLSClientConfig)
	assert.Equal(t, uint16(tls.VersionTLS12), tr.TLSClientConfig.MinVersion, "expected min tls 1.2")
	assert.NotNil(t, tr.Proxy, "expected http.ProxyFromEnvironment fn")
	assert.Nil(t, tr.TLSClientConfig.RootCAs)

	// verify url has expected domain
	assert.Equal(t, `https://svc.region.oc.0r4-c10ud.com`, baseClient.Host)

	// verify subdomains are preserved
	baseClient = &oci_common.BaseClient{}
	baseClient.Host = "avnzdivwaadfa-management.kms.us-phoenix-1.oraclecloud.com"
	err = configureClientFn(baseClient)
	assert.NoError(t, err)
	assert.Equal(t, `avnzdivwaadfa-management.kms.us-phoenix-1.oc.0r4-c10ud.com`, baseClient.Host)

	// verify non-match preserves original url
	baseClient = &oci_common.BaseClient{}
	baseClient.Host = "DUMMY_ENDPOINT"
	err = configureClientFn(baseClient)
	assert.NoError(t, err)
	assert.Equal(t, `DUMMY_ENDPOINT`, baseClient.Host)
}

// ensure use_obo_token env var results in `opc-obo-token` http header injection
// issue-routing-tag: terraform/default
func TestUnitBuildClientConfigureFn_interceptor(t *testing.T) {

	prevEnvVar, hadPreviousEnvVar := os.LookupEnv("use_obo_token")
	if hadPreviousEnvVar {
		defer os.Setenv("use_obo_token", prevEnvVar)
	} else {
		defer os.Unsetenv("use_obo_token")
	}

	os.Setenv("use_obo_token", "true")
	os.Setenv(globalvar.OboTokenAttrName, "fake-token")
	defer os.Unsetenv(globalvar.OboTokenAttrName)
	assert.Equal(t, "true", utils.GetEnvSettingWithBlankDefault("use_obo_token"))
	configProvider := oci_common.DefaultConfigProvider()
	httpClient := provider.BuildHttpClient()
	configureClientFn, err := provider.BuildConfigureClientFn(configProvider, httpClient)
	assert.NoError(t, err)

	baseClient := &oci_common.BaseClient{}
	err = configureClientFn(baseClient)
	assert.NoError(t, err)

	assert.NotNil(t, baseClient.Interceptor)
	r, _ := http.NewRequest("GET", "cloud.com", nil)
	baseClient.Interceptor(r)
	assert.Equal(t, "fake-token", r.Header.Get(globalvar.RequestHeaderOpcOboToken))

	// Update obo token and check
	os.Setenv(globalvar.OboTokenAttrName, "another-token")
	baseClient.Interceptor(r)
	assert.NotEqual(t, "fake-token", r.Header.Get(globalvar.RequestHeaderOpcOboToken))
	assert.Equal(t, "another-token", r.Header.Get(globalvar.RequestHeaderOpcOboToken))
}

/*
// issue-routing-tag: terraform/default
func TestUnitSupportChangeOboToken(t *testing.T) {
	t.Skip("Run manual with a valid obo token")

	for _, apiKeyConfigAttribute := range provider.ApiKeyConfigAttributes {
		apiKeyConfigAttributeEnvValue := utils.GetEnvSettingWithBlankDefault(apiKeyConfigAttribute)
		if apiKeyConfigAttributeEnvValue != "" {
			unsetAtr := "TF_VAR_" + apiKeyConfigAttribute
			os.Unsetenv(unsetAtr)
			defer os.Setenv(unsetAtr, apiKeyConfigAttributeEnvValue)
		}
	}

	os.Setenv("use_obo_token", "true")
	os.Setenv(globalvar.OboTokenAttrName, "fake-token")
	defer os.Unsetenv(globalvar.OboTokenAttrName)
	assert.Equal(t, "true", utils.GetEnvSettingWithBlankDefault("use_obo_token"))
	r := &schema.Resource{
		Schema: provider.SchemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", "InstancePrincipal")
	d.Set("region", "us-phoenix-1")

	client := GetTestClients(d).budgetClient()
	assert.NotEmpty(t, client.Host)

	request := oci_budget.ListBudgetsRequest{}
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id")
	request.CompartmentId = &compartmentId
	fmt.Println("======= First List call with token fake-token ======")

	// manual verify request that contains "Opc-Obo-Token: fake-token"
	client.ListBudgets(context.Background(), request)

	fmt.Println("======= Second List call with token another-token ======")
	os.Setenv(globalvar.OboTokenAttrName, "another-token")
	// manual verify request that contains "Opc-Obo-Token: another-token"
	client.ListBudgets(context.Background(), request)
}

// issue-routing-tag: terraform/default
func TestUnitReadOboTokenFromFile(t *testing.T) {
	t.Skip("Run manual with a valid obo token")

	for _, apiKeyConfigAttribute := range provider.ApiKeyConfigAttributes {
		apiKeyConfigAttributeEnvValue := utils.GetEnvSettingWithBlankDefault(apiKeyConfigAttribute)
		if apiKeyConfigAttributeEnvValue != "" {
			unsetAtr := "TF_VAR_" + apiKeyConfigAttribute
			os.Unsetenv(unsetAtr)
			defer os.Setenv(unsetAtr, apiKeyConfigAttributeEnvValue)
		}
	}

	os.Setenv("use_obo_token", "true")

	tokenFile := "token_file"

	var file *os.File
	_, err := os.Stat(tokenFile)
	if os.IsNotExist(err) {
		file, _ = os.Create(tokenFile)
		file.WriteString("fake-token")
		defer os.Remove(tokenFile)
	}

	os.Setenv(globalvar.OboTokenPath, tokenFile)

	assert.Equal(t, "true", utils.GetEnvSettingWithBlankDefault("use_obo_token"))

	r := &schema.Resource{
		Schema: provider.SchemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", "InstancePrincipal")
	d.Set("region", "us-phoenix-1")

	client := GetTestClients(d).budgetClient()
	assert.NotEmpty(t, client.Host)

	request := oci_budget.ListBudgetsRequest{}
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id")
	request.CompartmentId = &compartmentId
	fmt.Println("======= First List call with token fake-token ======")

	// manual verify request that contains "Opc-Obo-Token: fake-token"
	client.ListBudgets(context.Background(), request)

	fmt.Println("======= Second List call with token another-token ======")

	// overwrite the token file
	file.WriteAt([]byte("another-token"), 0)
	// manual verify request that contains "Opc-Obo-Token: another-token"
	client.ListBudgets(context.Background(), request)
}

func TestUnitOboTokenAndApiKey(t *testing.T) {
	t.Skip("Run manual with a valid obo token")

	os.Setenv("use_obo_token", "true")
	os.Setenv(globalvar.OboTokenAttrName, "fake-token")
	defer os.Unsetenv(globalvar.OboTokenAttrName)
	assert.Equal(t, "true", utils.GetEnvSettingWithBlankDefault("use_obo_token"))
	r := &schema.Resource{
		Schema: provider.SchemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", "InstancePrincipal")
	d.Set("region", "us-phoenix-1")

	// Set API key with auth=InstancePrincipal, the API should be unset
	d.Set("user_ocid", utils.GetEnvSettingWithBlankDefault("user_ocid"))
	d.Set("fingerprint", utils.GetEnvSettingWithBlankDefault("fingerprint"))
	d.Set("private_key_path", utils.GetEnvSettingWithBlankDefault("private_key_path"))
	d.Set("private_key_password", utils.GetEnvSettingWithBlankDefault("private_key_password"))
	d.Set("private_key", utils.GetEnvSettingWithBlankDefault("private_key"))

	client := GetTestClients(d).budgetClient()
	assert.NotEmpty(t, client.Host)

	request := oci_budget.ListBudgetsRequest{}
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id")
	request.CompartmentId = &compartmentId
	fmt.Println("======= First List call with token fake-token ======")

	// manual verify request that contains "Opc-Obo-Token: fake-token"
	client.ListBudgets(context.Background(), request)

	fmt.Println("======= Second List call with token another-token ======")
	os.Setenv(oboTokenAttrName, "another-token")
	// manual verify request that contains "Opc-Obo-Token: another-token"
	client.ListBudgets(context.Background(), request)
}
*/
// issue-routing-tag: terraform/default
func TestUnitVerifyConfigForAPIKeyAuthIsNotSet_basic(t *testing.T) {
	httpreplay.SetScenario("TestVerifyConfigForAPIKeyAuthIsNotSet_basic")
	defer httpreplay.SaveScenario()
	for _, apiKeyConfigAttribute := range provider.ApiKeyConfigAttributes {
		apiKeyConfigAttributeEnvValue := utils.GetEnvSettingWithBlankDefault(apiKeyConfigAttribute)
		if apiKeyConfigAttributeEnvValue != "" {
			t.Skip("apiKeyConfigAttributes are set through environment variables, skip the test")
		}
	}

	r := &schema.Resource{
		Schema: provider.SchemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", "InstancePrincipal")
	d.Set("region", "us-phoenix-1")

	apiKeyConfigVariablesToUnset, ok := utils.CheckIncompatibleAttrsForApiKeyAuth(d, provider.ApiKeyConfigAttributes)
	assert.True(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 0, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("tenancy_ocid", testTenancyOCID)
	apiKeyConfigVariablesToUnset, ok = utils.CheckIncompatibleAttrsForApiKeyAuth(d, provider.ApiKeyConfigAttributes)
	assert.True(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 0, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("user_ocid", testUserOCID)
	apiKeyConfigVariablesToUnset, ok = utils.CheckIncompatibleAttrsForApiKeyAuth(d, provider.ApiKeyConfigAttributes)
	assert.False(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 1, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("fingerprint", testKeyFingerPrint)
	apiKeyConfigVariablesToUnset, ok = utils.CheckIncompatibleAttrsForApiKeyAuth(d, provider.ApiKeyConfigAttributes)
	assert.False(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 2, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("private_key", testPrivateKey)
	apiKeyConfigVariablesToUnset, ok = utils.CheckIncompatibleAttrsForApiKeyAuth(d, provider.ApiKeyConfigAttributes)
	assert.False(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 3, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("private_key_path", "path")
	apiKeyConfigVariablesToUnset, ok = utils.CheckIncompatibleAttrsForApiKeyAuth(d, provider.ApiKeyConfigAttributes)
	assert.False(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 4, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("private_key_password", "password")
	apiKeyConfigVariablesToUnset, ok = utils.CheckIncompatibleAttrsForApiKeyAuth(d, provider.ApiKeyConfigAttributes)
	assert.False(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 5, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)
}

/*
// This test verifies that user can specify private key paths with "~/" and they should resolve to the home directory
// issue-routing-tag: terraform/default
func TestUnitHomeDirectoryPrivateKeyPath_basic(t *testing.T) {
	privateKeyName := "TestUnitHomeDirectoryPrivateKeyPath_basic.pem"
	privateKeyPath := path.Join(utils.GetHomeFolder(), privateKeyName)
	err := utils.WriteTempFile(testPrivateKey, privateKeyPath)
	if err != nil {
		t.Fatalf("unable to write test private key into directory %s. Error: %v", privateKeyPath, err)
	}

	defer utils.RemoveFile(privateKeyPath)

	r := &schema.Resource{
		Schema: provider.SchemaMap(),
	}
	d := r.Data(nil)
	d.Set(globalvar.PrivateKeyPathAttrName, path.Join("~", privateKeyName))

	d.Set(globalvar.TenancyOcidAttrName, testTenancyOCID)
	d.Set(globalvar.AuthAttrName, globalvar.AuthAPIKeySetting)
	d.Set(globalvar.UserOcidAttrName, testUserOCID)
	d.Set(globalvar.FingerprintAttrName, testKeyFingerPrint)
	d.Set(globalvar.RegionAttrName, "us-phoenix-1")

	clients := &tf_client.OracleClients{
		SdkClientMap:  make(map[string]interface{}, len(oracleClientRegistrations.registeredClients)),
		Configuration: make(map[string]string),
	}
	sdkConfigProvider, err := provider.GetSdkConfigProvider(d, clients)
	assert.NoError(t, err)

	privateRsaKey, err := sdkConfigProvider.PrivateRSAKey()
	assert.NoError(t, err)
	assert.True(t, privateRsaKey != nil)
}

// issue-routing-tag: terraform/default
func TestUnitSecurityToken_basic(t *testing.T) {
	t.Skip("Run manual with a valid security token")

	r := &schema.Resource{
		Schema: provider.SchemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", globalvar.AuthSecurityToken)
	d.Set(globalvar.ConfigFileProfileAttrName, "DEFAULT")

	// Set API key, should be removed by auth=SecurityToken
	d.Set("user_ocid", utils.GetEnvSettingWithBlankDefault("user_ocid"))
	d.Set("fingerprint", utils.GetEnvSettingWithBlankDefault("fingerprint"))
	d.Set("private_key_path", utils.GetEnvSettingWithBlankDefault("private_key_path"))
	d.Set("private_key_password", utils.GetEnvSettingWithBlankDefault("private_key_password"))
	d.Set("private_key", utils.GetEnvSettingWithBlankDefault("private_key"))
	// Run CLI command "oci session authenticate" to get token and profile
	clients := &tf_client.OracleClients{
		SdkClientMap:  make(map[string]interface{}, len(oracleClientRegistrations.registeredClients)),
		Configuration: make(map[string]string),
	}
	sdkConfigProvider, err := provider.GetSdkConfigProvider(d, clients)
	_, empty := utils.CheckIncompatibleAttrsForApiKeyAuth(d)
	// API key should be removed
	assert.True(t, true, empty)
	assert.NoError(t, err)
	finger, _ := sdkConfigProvider.KeyFingerprint()
	assert.NotNil(t, finger)
	keyId, _ := sdkConfigProvider.KeyID()
	assert.NotNil(t, keyId)
	// Token format start with ST$
	assert.True(t, strings.HasPrefix(keyId, "ST$"))
	region, _ := sdkConfigProvider.Region()
	assert.NotNil(t, region)
	privateKey, _ := sdkConfigProvider.PrivateRSAKey()
	assert.NotNil(t, privateKey)
	client, err := oci_identity.NewIdentityClientWithConfigurationProvider(sdkConfigProvider)
	assert.NoError(t, err)
	assert.NotEmpty(t, client.Host)

	_, err = client.ListRegions(context.Background())
	assert.NoError(t, err)
}
*/
