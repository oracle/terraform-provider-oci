// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package acctest

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/oci/globalvar"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"

	oci_common "github.com/oracle/oci-go-sdk/v49/common"

	tf_client "github.com/terraform-providers/terraform-provider-oci/oci/client"
	"github.com/terraform-providers/terraform-provider-oci/oci/provider"
	"github.com/terraform-providers/terraform-provider-oci/oci/utils"
)

/*
var subnetConfig = `
data "oci_identity_availability_domains" "ADs" {
	compartment_id = "${var.compartment_id}"
}

resource "oci_core_virtual_network" "t" {
	cidr_block = "10.0.0.0/16"
	compartment_id = "${var.compartment_id}"
	display_name = "network_name"
}

resource "oci_core_subnet" "WebSubnetAD1" {
	availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
	cidr_block          = "10.0.1.0/24"
	display_name        = "WebSubnetAD1"
	compartment_id      = "${var.compartment_id}"
	vcn_id              = "${oci_core_virtual_network.t.id}"
	route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
	security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
	dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
}`

var instanceConfig = subnetConfig + `
variable "InstanceImageOCID" {
  type = "map"
  default = {
	// See https://docs.us-phoenix-1.oraclecloud.com/images/
	// Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
	us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaadjnj3da72bztpxinmqpih62c2woscbp6l3wjn36by2cvmdhjub6a"
	us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaawufnve5jxze4xf7orejupw5iq3pms6cuadzjc7klojix6vmk42va"
	eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaagbrvhganmn7awcr7plaaf5vhabmzhx763z5afiitswjwmzh7upna"
	uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaajwtut4l7fo3cvyraate6erdkyf2wdk5vpk6fp6ycng3dv2y3ymvq"
  }
}

data "oci_identity_policies" "policies" {
	compartment_id = "${var.compartment_id}"
}

data "oci_load_balancer_protocols" "protocols" {
	compartment_id = "${var.compartment_id}"
}

data "oci_core_shape" "shapes" {
	compartment_id = "${var.compartment_id}"
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	image_id =  "${var.InstanceImageOCID[var.region]}"
}

resource "oci_core_instance" "t" {
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	display_name = "-tf-instance"
	image = "${var.InstanceImageOCID[var.region]}"
	shape = "VM.Standard2.1"
	subnet_id = "${oci_core_subnet.WebSubnetAD1.id}"
	metadata = {
		ssh_authorized_keys = "${var.ssh_public_key}"
	}

	timeouts {
		create = "15m"
	}
}
`

const (
	instanceDnsConfig = `
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.compartment_id}"
}

resource "oci_core_virtual_network" "t" {
	cidr_block      = "10.0.0.0/16"
	compartment_id  = "${var.compartment_id}"
	display_name    = "-tf-vcn"
	dns_label		= "testvcn"
}

resource "oci_core_subnet" "t" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block          = "10.0.1.0/24"
  display_name        = "-tf-subnet"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${oci_core_virtual_network.t.id}"
  route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
  security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
  dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
  dns_label			  = "testsubnet"
}

variable "InstanceImageOCID" {
  type = "map"
  default = {
    // Oracle-provided image "Oracle-Linux-7.4-2017.12.18-0"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaasc56hnpnx7swoyd2fw5gyvbn3kcdmqc2guiiuvnztl2erth62xnq"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaxrqeombwty6jyqgk3fraczdd63bv66xgfsqka4ktr7c57awr3p5a"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaayxmzu6n5hsntq4wlffpb4h6qh6z3uskpbm5v3v4egqlqvwicfbyq"
  }
}

resource "oci_core_instance" "t" {
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	display_name = "-tf-instance"
	image = "${var.InstanceImageOCID[var.region]}"
	shape = "VM.Standard2.1"
	create_vnic_details {
        subnet_id = "${oci_core_subnet.t.id}"
        hostname_label = "testinstance"
        display_name = "-tf-instance-vnic"
		defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
		freeform_tags = { "Department" = "Accounting" }
  	}
	metadata = {
		ssh_authorized_keys = "${var.ssh_public_key}"
	}
	timeouts {
		create = "15m"
	}
}
resource "oci_core_network_security_group" "test_network_security_group1" {
	compartment_id = "${var.compartment_id}"
	vcn_id         = "${oci_core_virtual_network.t.id}"
	display_name = "testNetworkSecurityGroup1"
}
resource "oci_core_network_security_group" "test_network_security_group2" {
	compartment_id = "${var.compartment_id}"
	vcn_id         = "${oci_core_virtual_network.t.id}"
	display_name = "testNetworkSecurityGroup2"
}` + tf_common.DefinedTagsDependencies
)
*/

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
	configFile, keyFile, err := writeConfigFile()
	assert.Nil(t, err)
	ProviderConfigTest(t, true, true, globalvar.AuthAPIKeySetting, "DEFAULT", nil)              // ApiKey with required fields + disable auto-retries
	ProviderConfigTest(t, false, true, globalvar.AuthAPIKeySetting, "DEFAULT", nil)             // ApiKey without required fields
	ProviderConfigTest(t, false, false, globalvar.AuthInstancePrincipalSetting, "DEFAULT", nil) // InstancePrincipal
	ProviderConfigTest(t, true, false, "invalid-auth-setting", "DEFAULT", nil)                  // Invalid auth + disable auto-retries
	ProviderConfigTest(t, false, false, globalvar.AuthAPIKeySetting, "PROFILE1", nil)           // correct profileName
	ProviderConfigTest(t, false, false, globalvar.AuthAPIKeySetting, "wrongProfile", nil)       // Invalid profileName
	//providerConfigTest(t, false, false, authAPIKeySetting, "PROFILE2", nil)           // correct profileName with mix and match
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
/* This function is used in the test asserts to verify that an element in a set contains certain properties
 * properties is a map of nameOfProperty -> expectedValueOfProperty
 * presentProperties is an array of property names that are expected to be set in the set element but we don't care about matching the value
 * will return nil (the positive response) if there is an element in the set that matches all properties in properties and presentProperties
 */
func CheckResourceSetContainsElementWithProperties(name, setKey string, properties map[string]string, presentProperties []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rm := s.RootModule()
		rs, ok := rm.Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		is := rs.Primary
		if is == nil {
			return fmt.Errorf("No primary instance: %s", name)
		}

		orderedKeys := []string{}
		for key, _ := range is.Attributes {
			orderedKeys = append(orderedKeys, key)
		}
		sort.Strings(orderedKeys)
		var currSetElementId string
		currMatchedAttributes := []string{}
		currMatchedPresentProperties := []string{}
		setElementMatch := func() bool {
			return len(currMatchedAttributes) == len(properties) && (presentProperties == nil || len(currMatchedPresentProperties) == len(presentProperties))
		}
		for _, key := range orderedKeys {
			prefix := fmt.Sprintf("%s.", setKey)
			if !strings.HasPrefix(key, prefix) {
				continue
			}
			attrWithSetIdRaw := strings.TrimPrefix(key, prefix)

			attrWithSetIdRawArr := strings.Split(attrWithSetIdRaw, ".")
			if len(attrWithSetIdRawArr) < 2 {
				continue
			}
			if currSetElementId == "" {
				currSetElementId = attrWithSetIdRawArr[0]
			}
			if attrWithSetIdRawArr[0] != currSetElementId {
				if setElementMatch() {
					return nil
				}
				currMatchedPresentProperties = []string{}
				currMatchedAttributes = []string{}
				currSetElementId = attrWithSetIdRawArr[0]
			}
			attributeName := strings.Join(attrWithSetIdRawArr[1:], ".")
			for propName, value := range properties {
				if propName == attributeName && value == is.Attributes[key] {
					currMatchedAttributes = append(currMatchedAttributes, propName)
				}
			}
			if presentProperties != nil {
				for _, propName := range presentProperties {
					if propName == attributeName {
						currMatchedPresentProperties = append(currMatchedPresentProperties, propName)
					}
				}
			}
		}
		if setElementMatch() {
			return nil
		}

		return fmt.Errorf("%s: Set Attribute '%s' does not contain an element with attributes %v %v\nAttributesInStatefile: %v", name, setKey, properties, presentProperties, is.Attributes)
	}
}

func CheckResourceSetContainsElementWithPropertiesContainingNestedSets(name, setKey string, properties map[string]interface{}, presentProperties []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rm := s.RootModule()
		rs, ok := rm.Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		is := rs.Primary
		if is == nil {
			return fmt.Errorf("No primary instance: %s", name)
		}

		orderedKeys := []string{}
		for key, _ := range is.Attributes {
			orderedKeys = append(orderedKeys, key)
		}
		sort.Strings(orderedKeys)
		var currSetElementId string
		currMatchedAttributes := []string{}
		currMatchedPresentProperties := []string{}
		leafProperties := 0
		for _, value := range properties {
			if _, ok := value.(string); ok {
				leafProperties++
			}
		}
		setElementMatch := func() bool {
			return len(currMatchedAttributes) == leafProperties && (presentProperties == nil || len(currMatchedPresentProperties) == len(presentProperties))
		}
		for _, key := range orderedKeys {
			prefix := fmt.Sprintf("%s.", setKey)
			if !strings.HasPrefix(key, prefix) {
				continue
			}
			attrWithSetIdRaw := strings.TrimPrefix(key, prefix)

			attrWithSetIdRawArr := strings.Split(attrWithSetIdRaw, ".")
			if len(attrWithSetIdRawArr) < 2 {
				continue
			}
			if attrWithSetIdRawArr[0] != currSetElementId {
				if setElementMatch() {
					return nil
				}
				currMatchedPresentProperties = []string{}
				currMatchedAttributes = []string{}
				currSetElementId = attrWithSetIdRawArr[0]

				//check nested set properties, we do it in this if statement to avoid repeating the same checks for each key in the loop. We only need to check once per set element id
				for propName, value := range properties {
					if valueSet, ok := value.([]map[string]interface{}); ok {
						for _, nestedSetElement := range valueSet {
							nestedSetCheck := CheckResourceSetContainsElementWithPropertiesContainingNestedSets(name, fmt.Sprintf("%s.%s.%s", setKey, currSetElementId, propName), nestedSetElement, nil)
							if err := nestedSetCheck(s); err != nil {
								return err
							}
						}
					}
				}
			}
			attributeName := strings.Join(attrWithSetIdRawArr[1:], ".")
			for propName, value := range properties {
				if valueStr, ok := value.(string); ok {
					if propName == attributeName && valueStr == is.Attributes[key] {
						currMatchedAttributes = append(currMatchedAttributes, propName)
					}
				}
			}
			if presentProperties != nil {
				for _, propName := range presentProperties {
					if propName == attributeName {
						currMatchedPresentProperties = append(currMatchedPresentProperties, propName)
					}
				}
			}
		}
		if setElementMatch() {
			return nil
		}

		return fmt.Errorf("%s: Set Attribute '%s' does not contain an element with attributes %v %v\nAttributesInStatefile: %v", name, setKey, properties, presentProperties, is.Attributes)
	}

}

func ProviderConfigTest(t *testing.T, disableRetries bool, skipRequiredField bool, auth string, configFileProfile string, configFunc ConfigFunc) {
	r := &schema.Resource{
		Schema: provider.SchemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", auth)
	if !skipRequiredField {
		d.Set("tenancy_ocid", testTenancyOCID)
	}
	if configFileProfile == "" || configFileProfile == "DEFAULT" {
		d.Set("user_ocid", testUserOCID)
		d.Set("fingerprint", testKeyFingerPrint)
		d.Set("private_key", testPrivateKey)
		//d.Set("private_key_path", "")
		d.Set("region", "us-phoenix-1")
		d.Set("private_key_password", "password")
	}
	if configFileProfile == "PROFILE3" {
		d.Set("fingerprint", testKeyFingerPrint)
	}
	if disableRetries {
		d.Set("disable_auto_retries", disableRetries)
	}
	if configFileProfile != "" {
		d.Set("config_file_profile", configFileProfile)
	}

	// Use config func for export (resource discovery)
	configureProviderFn := configFunc
	//userAgent := fmt.Sprintf(globalvar.ExportUserAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, globalvar.Version)

	// If no ConfigFunc use ProviderConfig
	if configureProviderFn == nil {
		configureProviderFn = provider.ProviderConfig
		//userAgent = fmt.Sprintf(globalvar.UserAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, sdkMeta.SDKVersionString(), provider.TerraformCLIVersion, globalvar.DefaultUserAgentProviderName, globalvar.Version)

	}
	client, err := configureProviderFn(d)

	if configFileProfile == "wrongProfile" {
		assert.Equal(t, "configuration file did not contain profile: wrongProfile", err.Error())
		return
	}
	if configFileProfile == "PROFILE2" {
		assert.Equal(t, "can not Create client, bad configuration: did not find a proper configuration for private key", err.Error())
		return
	}
	switch auth {
	case globalvar.AuthAPIKeySetting, "":
		if skipRequiredField {
			assert.Equal(t, err, nil)
			return
		}
	default:
		assert.Error(t, err, fmt.Sprintf("auth must be one of '%s' or '%s' or '%s'", globalvar.AuthAPIKeySetting, globalvar.AuthInstancePrincipalSetting, globalvar.AuthInstancePrincipalWithCertsSetting))
		return
	}
	assert.Nil(t, err)
	assert.NotNil(t, client)

	oracleClient, ok := client.(*tf_client.OracleClients)
	assert.True(t, ok)

	testClient := func(c *oci_common.BaseClient) {
		assert.NotNil(t, c)
		assert.NotNil(t, c.HTTPClient)
		assert.Exactly(t, c.UserAgent, globalvar.UserAgentFormatter)
		assert.NotNil(t, c.Interceptor)
		assert.NotNil(t, c.Signer)
	}

	testClient(&oracleClient.IdentityClient().BaseClient)
}
