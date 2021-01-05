// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"runtime"
	"sort"
	"strings"
	"testing"

	oci_budget "github.com/oracle/oci-go-sdk/v31/budget"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"crypto/tls"
	"io/ioutil"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	sdkMeta "github.com/hashicorp/terraform-plugin-sdk/meta"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v31/common"
	"github.com/stretchr/testify/assert"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider
var requiredTestEnvVars = []string{"compartment_ocid", "compartment_id_for_create", "compartment_id_for_update", "tags_import_if_exists"}
var requiredKeyAuthEnvVars = []string{"tenancy_ocid", "user_ocid", "fingerprint"}
var requiredOboTokenAuthEnvVars = []string{"tenancy_ocid", "obo_token"}

type ConfigFunc func(d *schema.ResourceData) (interface{}, error)

func init() {
	testAccProvider = testProvider(func(d *schema.ResourceData) (interface{}, error) {
		return GetTestClients(d), nil
	}).(*schema.Provider)

	testAccProviders = map[string]terraform.ResourceProvider{
		"oci": testAccProvider,
	}
}

// Provider is the adapter for terraform, that gives access to all the resources
func testProvider(configfn schema.ConfigureFunc) terraform.ResourceProvider {
	result := &schema.Provider{
		DataSourcesMap: DataSourcesMap(),
		Schema:         schemaMap(),
		ResourcesMap:   ResourcesMap(),
		ConfigureFunc:  configfn,
	}

	// Additions for test parameters
	result.Schema["test_time_maintenance_reboot_due"] = &schema.Schema{Type: schema.TypeString, Optional: true}

	return result
}

func commonTestVariables() string {
	return `
	variable "tenancy_ocid" {
		default = "` + getEnvSettingWithBlankDefault("tenancy_ocid") + `"
	}

	variable "ssh_public_key" {
		default = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
	}

	variable "region" {
		default = "` + getEnvSettingWithBlankDefault("region") + `"
	}

	`
}

func testProviderConfig() string {
	return `
	# Need to have this block even though it's empty; for import testing
	provider "oci" {
	}
	` + commonTestVariables()
}

func testAccPreCheck(t *testing.T) {
	envVarChecklist := []string{}
	copy(envVarChecklist, requiredTestEnvVars)
	if getEnvSettingWithDefault("use_obo_token", "false") != "false" {
		envVarChecklist = append(envVarChecklist, requiredOboTokenAuthEnvVars...)
	} else {
		envVarChecklist = append(envVarChecklist, requiredKeyAuthEnvVars...)
	}

	for _, envVar := range envVarChecklist {
		assertEnvAvailable(envVar, t)
	}

}

func assertEnvAvailable(envVar string, t *testing.T) {
	if v := getEnvSettingWithBlankDefault(envVar); v == "" {
		t.Fatal("TF_VAR_" + envVar + " must be set for acceptance tests")
	}
}

func getCompartmentIDForLegacyTests() string {
	var compartmentId string
	if compartmentId = getEnvSettingWithDefault("compartment_ocid", "compartment_ocid"); compartmentId == "compartment_ocid" {
		compartmentId = getRequiredEnvSetting("compartment_id_for_create")
	}
	return compartmentId
}

func legacyTestProviderConfig() string {
	// Use the same config as the generated tests.
	config := testProviderConfig()

	// Add the 'compartment_id' used by the legacy tests.
	return config + `variable "compartment_id" {
		default = "` + getCompartmentIDForLegacyTests() + `"
	}
	`
}

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
}` + DefinedTagsDependencies
)

const (
	requestQueryOpcTimeMaintenanceRebootDue = "opc-time-maintenance-reboot-due"
)

func writeTempFile(data string, originFileName string) (err error) {
	f, err := os.OpenFile(originFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		f.WriteString(data)
	}
	return err
}

func GetTestClients(data *schema.ResourceData) *OracleClients {
	r := &schema.Resource{
		Schema: schemaMap(),
	}
	d := r.Data(nil)
	d.SetId(getEnvSettingWithBlankDefault("tenancy_ocid"))
	d.Set("tenancy_ocid", getEnvSettingWithBlankDefault("tenancy_ocid"))
	d.Set("region", getEnvSettingWithDefault("region", "us-phoenix-1"))

	if auth := getEnvSettingWithDefault("auth", authAPIKeySetting); auth == authAPIKeySetting {
		d.Set("auth", getEnvSettingWithDefault("auth", authAPIKeySetting))
		d.Set("user_ocid", getEnvSettingWithBlankDefault("user_ocid"))
		d.Set("fingerprint", getEnvSettingWithBlankDefault("fingerprint"))
		d.Set("private_key_path", getEnvSettingWithBlankDefault("private_key_path"))
		d.Set("private_key_password", getEnvSettingWithBlankDefault("private_key_password"))
		d.Set("private_key", getEnvSettingWithBlankDefault("private_key"))
	} else {
		d.Set("auth", getEnvSettingWithDefault("auth", auth))
	}

	terraformCLIVersion = testTerraformCLIVersion
	client, err := ProviderConfig(d)
	if err != nil {
		panic(err)
	}

	// This is a test hook to support creating instances that have a maintenance reboot time set
	// The test hook allows 'time_maintenance_reboot_due' field to be tested for instance datasources/resources
	// This is controlled by a provider option rather than environment variable: so that the tests can run in parallel
	// without affecting one another and also allow individual test steps to alter this
	//
	// If we have additional test hooks that need to be supported in this manner, then the following logic should be
	// compartmentalized and registered with the test provider in a scalable manner.
	maintenanceRebootTime, ok := data.GetOkExists("test_time_maintenance_reboot_due")
	if ok {
		computeClient := client.(*OracleClients).computeClient()
		baseInterceptor := computeClient.Interceptor
		computeClient.Interceptor = func(r *http.Request) error {
			if err := baseInterceptor(r); err != nil {
				return err
			}

			if r.Method == http.MethodPost && (strings.Contains(r.URL.Path, "/instances")) {
				query := r.URL.Query()
				query.Set(requestQueryOpcTimeMaintenanceRebootDue, maintenanceRebootTime.(string))
				r.URL.RawQuery = query.Encode()
			}
			return nil
		}
	}

	return client.(*OracleClients)
}

// This test runs the Provider sanity checks.
func TestUnitProvider(t *testing.T) {
	// Real client for the sanity check. Makes this more of an acceptance test.
	client := &OracleClients{}
	testProvider := &schema.Provider{
		DataSourcesMap: DataSourcesMap(),
		Schema:         schemaMap(),
		ResourcesMap:   ResourcesMap(),
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		},
	}
	if err := testProvider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

// Don't worry, this key is NOT a valid API key
var testPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-EDE3-CBC,9F4D00DEF02B2B75

IbSQEhNjPeRt49jUhZbhAEaAIG4L9IokDksw/P/QdCPXzZT008xzYK/zmxkz7so1
ZwvIYHn07E0Ul6fIHR6kjw/+MD7AWluCN1FLHs3PHc4XF4THUCKFCC90FvGJ2PEs
kEh7oJ4azZA/PH51g4rSgWpYtH5B/S6ioE2eZ9jJ/prH+34pCuOpX4AvXEFl5zue
pjFm5FhsReAhZ/9eCvjgjIWDHKc7PRfinwSydVHQSzgDnuq+GTMzQh6eztS+EuAp
MLg7w0mazTqmPOuMT+mw9SHGaIePGzA9TcwB1y3QgkYsg3Ch20uN/sUymgQ4PEKI
njXLldWDYvFvv1Tv3/8IOjCEodQ4P/5oWz7msrLh3QF+EhF7lQPYO7132e9Hvz3C
hTmcygmVGrPCtOY1jzuqy+/Kmt4Gv8FQpSnO7i8wFvt5v0N26av18RO10CzYY1ut
EV6WvynimFUtg1Lo03cadh7bspNohSXfFLpbNTji5NwHrIa+UQqTw3h4/zSPZHJl
NwHwM2I8N5lcCsqmSbM01+uTRG3QZ5i1BS8fsArHaAcvPyLvOy4mZGKkpuNlLDXo
qrCCsb+0m9jHR2bzx5AGp4impdHm2Qi3vTV3dMe277wqKkU5qfd5yDbL2eTqAYzQ
hXpPmTjquOTNYdbvoNsOg4TCHZv7WCsGY0nNMPrRO7zXCDApA6cKDJzagbqhW5Zu
/yz7sDT2D3wzE2WXUbtIBLevXyF0OS3AL7AgfbcyAviByOfmEb7WCP9jmdCFaLwY
SgNh9AjeOgkEEr/cRg1kBAXt0kuE7By0w+/ODJHZYelG0wg5nxhseA9Kc596XIJl
NyjbL87CXGfXmMoSYYTA4rzbtCDMmee7xHtbWiYKF1VGxNaGkQ5nnZSJLhCaI6rH
AD0XYwxv92j4fIjHqonbY/dlIKPot1t3VRcdnebbZMjAcNZ63n+I/iVla3DJpWLO
1gT50A4H2uEAve+WWFWmDQe2rfg5wwUtVVkot+Tn3McB6RzNqgcs0c+7uNDnDcOB
WtQ1OfniE1TdoFCPfYcDw8ngimw7uMYwp4mZIYtwlk7Z5GFl4YpNQeLOgh368ao4
8HL7EnTZmiU5cMbuaA8cZmUbgBqiQY0DtLF22VquThi0QOeUMJxJ6N1QUPckD3AU
dikEn0gilOsDQ51fnOsgk9J2uCz8rd5bnyUXlIguj5pyz6S7agyYFhRrXessVzHd
3889QM9V82+px5mv4qCvMn6ReYOvC+KSY1hn4ljXsndOM+6hQzD5CZKeL948pXRn
G7nqbG9D44wLklOz6mkIvqLn3qxEFWapl9UK7yfzjoezGoqeNFweadZ10Kp2+Umu
Sa759/2YDCZLDzaVVoLDTHLzi9ejpAkUIXgEFaPNGzQ8DYiL8N2klRozLSlnDEMr
xTHuOMkklNO7SiTluAUBvXrjxfGqe/gwJOHxXQGHC8W6vyhR2BdVx9PKFVebWjlr
gzRMpGgWnjsaz0ldu3uO7ozRxZg8FgdToIzAIaTytpHKI8HvONvPJlYywOMC1gRi
KwX6p26xaVtCV8PbDpF3RHuEJV1NU6PDIhaIHhdL374BiX/KmcJ6yv7tbkczpK+V
-----END RSA PRIVATE KEY-----`

var testKeyFingerPrint = "b4:8a:7d:54:e6:81:04:b2:fa:ce:ba:55:34:dd:00:00"
var testTenancyOCID = "ocid1.tenancy.oc1..faketenancy"
var testUserOCID = "ocid1.user.oc1..fakeuser"

func providerConfigTest(t *testing.T, disableRetries bool, skipRequiredField bool, auth string, configFileProfile string, configFunc ConfigFunc) {
	r := &schema.Resource{
		Schema: schemaMap(),
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
	userAgent := fmt.Sprintf(exportUserAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, Version)

	// If no ConfigFunc use ProviderConfig
	if configureProviderFn == nil {
		configureProviderFn = ProviderConfig
		userAgent = fmt.Sprintf(userAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, sdkMeta.SDKVersionString(), terraformCLIVersion, defaultUserAgentProviderName, Version)

	}
	client, err := configureProviderFn(d)

	if configFileProfile == "wrongProfile" {
		assert.Equal(t, "configuration file did not contain profile: wrongProfile", err.Error())
		return
	}
	if configFileProfile == "PROFILE2" {
		assert.Equal(t, "can not create client, bad configuration: did not find a proper configuration for private key", err.Error())
		return
	}
	switch auth {
	case authAPIKeySetting, "":
		if skipRequiredField {
			assert.Equal(t, err, nil)
			return
		}
	case authInstancePrincipalSetting:
		apiKeyConfigVariablesToUnset, ok := checkIncompatibleAttrsForApiKeyAuth(d)
		assert.False(t, ok)
		assert.Equal(t, fmt.Sprintf("user credentials %v should be removed from the configuration", strings.Join(apiKeyConfigVariablesToUnset, ", ")), err.Error())
		return
	case authInstancePrincipalWithCertsSetting:
		apiKeyConfigVariablesToUnset, ok := checkIncompatibleAttrsForApiKeyAuth(d)
		assert.False(t, ok)
		assert.Equal(t, fmt.Sprintf("user credentials %v should be removed from the configuration", strings.Join(apiKeyConfigVariablesToUnset, ", ")), err.Error())
		return
	case authSecurityToken:
		apiKeyConfigVariablesToUnset, ok := checkIncompatibleAttrsForApiKeyAuth(d)
		assert.False(t, ok)
		assert.Equal(t, fmt.Sprintf("user credentials %v should be removed from the configuration", strings.Join(apiKeyConfigVariablesToUnset, ", ")), err.Error())
		return
	default:
		assert.Error(t, err, fmt.Sprintf("auth must be one of '%s' or '%s' or '%s'", authAPIKeySetting, authInstancePrincipalSetting, authInstancePrincipalWithCertsSetting))
		return
	}
	assert.Nil(t, err)
	assert.NotNil(t, client)

	oracleClient, ok := client.(*OracleClients)
	assert.True(t, ok)

	testClient := func(c *oci_common.BaseClient) {
		assert.NotNil(t, c)
		assert.NotNil(t, c.HTTPClient)
		assert.Exactly(t, c.UserAgent, userAgent)
		assert.NotNil(t, c.Interceptor)
		assert.NotNil(t, c.Signer)
	}

	testClient(&oracleClient.blockstorageClient().BaseClient)
	testClient(&oracleClient.computeClient().BaseClient)
	testClient(&oracleClient.databaseClient().BaseClient)
	testClient(&oracleClient.identityClient().BaseClient)
	testClient(&oracleClient.virtualNetworkClient().BaseClient)
	testClient(&oracleClient.objectStorageClient().BaseClient)
	testClient(&oracleClient.loadBalancerClient().BaseClient)
}

func writeConfigFile() (string, string, error) {
	dataTpl := `[DEFAULT]
user=%s
fingerprint=%s
tenancy=%s
region=%s
[PROFILE1]
user=%s
fingerprint=%s
key_file=%s
passphrase=%s
[PROFILE2]
user=%s
[PROFILE3]
user=%s
key_file=%s
passphrase=%s
[PROFILE4]
key_file=%s
fingerprint=%s
tenancy=%s
region=%s
security_token_file=%s
`
	keyPath := path.Join(getHomeFolder(), defaultConfigDirName, "oci_api_key.pem")
	configPath := path.Join(getHomeFolder(), defaultConfigDirName, defaultConfigFileName)
	os.MkdirAll(path.Join(getHomeFolder(), defaultConfigDirName), 0700)
	err := writeTempFile(testPrivateKey, keyPath)
	if err != nil {
		return "", "", err
	}
	data := fmt.Sprintf(dataTpl, "invalid user", "invalid fingerprint", testTenancyOCID, "us-phoenix-1", testUserOCID, testKeyFingerPrint, keyPath, "password", "invalid user2",
		testUserOCID, keyPath, "password")
	err = writeTempFile(data, configPath)
	return keyPath, configPath, err
}

func removeFile(file string) {
	os.Remove(file)
}

func TestUnitProviderConfig(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip TestProviderConfig in HttpReplay mode.")
	}
	if os.Getenv("TF_HOME_OVERRIDE") == "" {
		t.Skip("This run requires you to set TF_HOME_OVERRIDE")
	}
	providerConfigTest(t, true, true, authAPIKeySetting, "", nil)              // ApiKey with required fields + disable auto-retries
	providerConfigTest(t, false, true, authAPIKeySetting, "", nil)             // ApiKey without required fields
	providerConfigTest(t, false, false, authInstancePrincipalSetting, "", nil) // InstancePrincipal
	providerConfigTest(t, true, false, "invalid-auth-setting", "", nil)        // Invalid auth + disable auto-retries
	configFile, keyFile, err := writeConfigFile()
	assert.Nil(t, err)
	providerConfigTest(t, true, true, authAPIKeySetting, "DEFAULT", nil)              // ApiKey with required fields + disable auto-retries
	providerConfigTest(t, false, true, authAPIKeySetting, "DEFAULT", nil)             // ApiKey without required fields
	providerConfigTest(t, false, false, authInstancePrincipalSetting, "DEFAULT", nil) // InstancePrincipal
	providerConfigTest(t, true, false, "invalid-auth-setting", "DEFAULT", nil)        // Invalid auth + disable auto-retries
	providerConfigTest(t, false, false, authAPIKeySetting, "PROFILE1", nil)           // correct profileName
	providerConfigTest(t, false, false, authAPIKeySetting, "wrongProfile", nil)       // Invalid profileName
	//providerConfigTest(t, false, false, authAPIKeySetting, "PROFILE2", nil)           // correct profileName with mix and match
	providerConfigTest(t, false, false, authAPIKeySetting, "PROFILE3", nil) // correct profileName with mix and match & env
	defer removeFile(configFile)
	defer removeFile(keyFile)
	os.RemoveAll(path.Join(getHomeFolder(), defaultConfigDirName))
}

// ensure the http client is configured with the expected settings
func TestUnitBuildHttpClient(t *testing.T) {
	client := buildHttpClient()
	assert.Equal(t, time.Duration(defaultRequestTimeout), client.Timeout)

	tr := client.Transport.(*http.Transport)
	assert.NotNil(t, tr.TLSClientConfig)
	assert.Equal(t, uint16(tls.VersionTLS12), tr.TLSClientConfig.MinVersion, "expected min tls 1.2")
	assert.NotNil(t, tr.Proxy, "expected http.ProxyFromEnvironment fn")
}

// ensure desired http client settings are not removed when sdk clients are configured
func TestUnitBuildClientConfigureFn(t *testing.T) {
	configProvider := oci_common.DefaultConfigProvider()
	httpClient := buildHttpClient()
	configureClientFn, err := buildConfigureClientFn(configProvider, httpClient)
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

	prevEnvVar, hadPreviousEnvVar := os.LookupEnv(customCertLocationEnv)
	if hadPreviousEnvVar {
		defer os.Setenv(customCertLocationEnv, prevEnvVar)
	} else {
		defer os.Unsetenv(customCertLocationEnv)
	}

	os.Setenv(customCertLocationEnv, tempCert.Name())
	assert.Equal(t, tempCert.Name(), getEnvSettingWithBlankDefault(customCertLocationEnv))
	configProvider := oci_common.DefaultConfigProvider()
	httpClient := buildHttpClient()
	configureClientFn, err := buildConfigureClientFn(configProvider, httpClient)
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
func TestUnitBuildClientConfigureFn_acceptLocalCerts(t *testing.T) {
	prevEnvVar, hadPreviousEnvVar := os.LookupEnv(acceptLocalCerts)
	if hadPreviousEnvVar {
		defer os.Setenv(acceptLocalCerts, prevEnvVar)
	} else {
		defer os.Unsetenv(acceptLocalCerts)
	}

	// ensure disabled by default - no env var
	os.Unsetenv(acceptLocalCerts)
	assert.Empty(t, getEnvSettingWithBlankDefault(acceptLocalCerts))
	configProvider := oci_common.DefaultConfigProvider()
	httpClient := buildHttpClient()
	configureClientFn, _ := buildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr := httpClient.Transport.(*http.Transport)
	assert.NotNil(t, tr.Proxy, "expected http.ProxyFromEnvironment fn")
	assert.Equal(t, uint16(tls.VersionTLS12), tr.TLSClientConfig.MinVersion, "expected min tls 1.2")
	assert.False(t, tr.TLSClientConfig.InsecureSkipVerify)

	// ensure disabled by default - env var with empty string
	os.Setenv(acceptLocalCerts, "")
	configProvider = oci_common.DefaultConfigProvider()
	httpClient = buildHttpClient()
	configureClientFn, _ = buildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr = httpClient.Transport.(*http.Transport)
	assert.False(t, tr.TLSClientConfig.InsecureSkipVerify)

	// ensure disabled by default - not parsable boolean string
	os.Setenv(acceptLocalCerts, "ftarlusee")
	configProvider = oci_common.DefaultConfigProvider()
	httpClient = buildHttpClient()
	configureClientFn, _ = buildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr = httpClient.Transport.(*http.Transport)
	assert.False(t, tr.TLSClientConfig.InsecureSkipVerify)

	// ensure explicitly disabled - env var set to false
	os.Setenv(acceptLocalCerts, "false")
	configProvider = oci_common.DefaultConfigProvider()
	httpClient = buildHttpClient()
	configureClientFn, _ = buildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr = httpClient.Transport.(*http.Transport)
	assert.False(t, tr.TLSClientConfig.InsecureSkipVerify)

	// ensure explicitly enabled - env var set to true
	os.Setenv(acceptLocalCerts, "true")
	configProvider = oci_common.DefaultConfigProvider()
	httpClient = buildHttpClient()
	configureClientFn, _ = buildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr = httpClient.Transport.(*http.Transport)
	assert.True(t, tr.TLSClientConfig.InsecureSkipVerify)

	// verify assumption that "1" will also coerce to true
	os.Setenv(acceptLocalCerts, "1")
	configProvider = oci_common.DefaultConfigProvider()
	httpClient = buildHttpClient()
	configureClientFn, _ = buildConfigureClientFn(configProvider, httpClient)
	configureClientFn(&oci_common.BaseClient{})

	tr = httpClient.Transport.(*http.Transport)
	assert.True(t, tr.TLSClientConfig.InsecureSkipVerify)
}

// ensure a custom domain can be targeted and expected http client settings are preserved
func TestUnitBuildClientConfigureFn_withDomainNameOverride(t *testing.T) {

	prevEnvVar, hadPreviousEnvVar := os.LookupEnv(domainNameOverrideEnv)
	if hadPreviousEnvVar {
		defer os.Setenv(domainNameOverrideEnv, prevEnvVar)
	} else {
		defer os.Unsetenv(domainNameOverrideEnv)
	}

	os.Setenv(domainNameOverrideEnv, "0r4-c10ud.com")
	assert.Equal(t, "0r4-c10ud.com", getEnvSettingWithBlankDefault(domainNameOverrideEnv))
	configProvider := oci_common.DefaultConfigProvider()
	httpClient := buildHttpClient()
	configureClientFn, err := buildConfigureClientFn(configProvider, httpClient)
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

// ensure use_obo_token env var results in `opc-obo-token` http header injection
func TestUnitBuildClientConfigureFn_interceptor(t *testing.T) {

	prevEnvVar, hadPreviousEnvVar := os.LookupEnv("use_obo_token")
	if hadPreviousEnvVar {
		defer os.Setenv("use_obo_token", prevEnvVar)
	} else {
		defer os.Unsetenv("use_obo_token")
	}

	os.Setenv("use_obo_token", "true")
	os.Setenv(oboTokenAttrName, "fake-token")
	defer os.Unsetenv(oboTokenAttrName)
	assert.Equal(t, "true", getEnvSettingWithBlankDefault("use_obo_token"))
	configProvider := oci_common.DefaultConfigProvider()
	httpClient := buildHttpClient()
	configureClientFn, err := buildConfigureClientFn(configProvider, httpClient)
	assert.NoError(t, err)

	baseClient := &oci_common.BaseClient{}
	err = configureClientFn(baseClient)
	assert.NoError(t, err)

	assert.NotNil(t, baseClient.Interceptor)
	r, _ := http.NewRequest("GET", "cloud.com", nil)
	baseClient.Interceptor(r)
	assert.Equal(t, "fake-token", r.Header.Get(requestHeaderOpcOboToken))

	// Update obo token and check
	os.Setenv(oboTokenAttrName, "another-token")
	baseClient.Interceptor(r)
	assert.NotEqual(t, "fake-token", r.Header.Get(requestHeaderOpcOboToken))
	assert.Equal(t, "another-token", r.Header.Get(requestHeaderOpcOboToken))
}

func TestUnitSupportChangeOboToken(t *testing.T) {
	t.Skip("Run manual with a valid obo token")

	for _, apiKeyConfigAttribute := range apiKeyConfigAttributes {
		apiKeyConfigAttributeEnvValue := getEnvSettingWithBlankDefault(apiKeyConfigAttribute)
		if apiKeyConfigAttributeEnvValue != "" {
			unsetAtr := "TF_VAR_" + apiKeyConfigAttribute
			os.Unsetenv(unsetAtr)
			defer os.Setenv(unsetAtr, apiKeyConfigAttributeEnvValue)
		}
	}

	os.Setenv("use_obo_token", "true")
	os.Setenv(oboTokenAttrName, "fake-token")
	defer os.Unsetenv(oboTokenAttrName)
	assert.Equal(t, "true", getEnvSettingWithBlankDefault("use_obo_token"))
	r := &schema.Resource{
		Schema: schemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", "InstancePrincipal")
	d.Set("region", "us-phoenix-1")

	client := GetTestClients(d).budgetClient()
	assert.NotEmpty(t, client.Host)

	request := oci_budget.ListBudgetsRequest{}
	compartmentId := getEnvSettingWithBlankDefault("compartment_id")
	request.CompartmentId = &compartmentId
	fmt.Println("======= First List call with token fake-token ======")

	// manual verify request that contains "Opc-Obo-Token: fake-token"
	client.ListBudgets(context.Background(), request)

	fmt.Println("======= Second List call with token another-token ======")
	os.Setenv(oboTokenAttrName, "another-token")
	// manual verify request that contains "Opc-Obo-Token: another-token"
	client.ListBudgets(context.Background(), request)
}

func TestUnitReadOboTokenFromFile(t *testing.T) {
	t.Skip("Run manual with a valid obo token")

	for _, apiKeyConfigAttribute := range apiKeyConfigAttributes {
		apiKeyConfigAttributeEnvValue := getEnvSettingWithBlankDefault(apiKeyConfigAttribute)
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

	os.Setenv(oboTokenPath, tokenFile)

	assert.Equal(t, "true", getEnvSettingWithBlankDefault("use_obo_token"))

	r := &schema.Resource{
		Schema: schemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", "InstancePrincipal")
	d.Set("region", "us-phoenix-1")

	client := GetTestClients(d).budgetClient()
	assert.NotEmpty(t, client.Host)

	request := oci_budget.ListBudgetsRequest{}
	compartmentId := getEnvSettingWithBlankDefault("compartment_id")
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

func TestUnitVerifyConfigForAPIKeyAuthIsNotSet_basic(t *testing.T) {
	httpreplay.SetScenario("TestVerifyConfigForAPIKeyAuthIsNotSet_basic")
	defer httpreplay.SaveScenario()
	for _, apiKeyConfigAttribute := range apiKeyConfigAttributes {
		apiKeyConfigAttributeEnvValue := getEnvSettingWithBlankDefault(apiKeyConfigAttribute)
		if apiKeyConfigAttributeEnvValue != "" {
			t.Skip("apiKeyConfigAttributes are set through environment variables, skip the test")
		}
	}

	r := &schema.Resource{
		Schema: schemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", "InstancePrincipal")
	d.Set("region", "us-phoenix-1")

	apiKeyConfigVariablesToUnset, ok := checkIncompatibleAttrsForApiKeyAuth(d)
	assert.True(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 0, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("tenancy_ocid", testTenancyOCID)
	apiKeyConfigVariablesToUnset, ok = checkIncompatibleAttrsForApiKeyAuth(d)
	assert.True(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 0, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("user_ocid", testUserOCID)
	apiKeyConfigVariablesToUnset, ok = checkIncompatibleAttrsForApiKeyAuth(d)
	assert.False(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 1, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("fingerprint", testKeyFingerPrint)
	apiKeyConfigVariablesToUnset, ok = checkIncompatibleAttrsForApiKeyAuth(d)
	assert.False(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 2, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("private_key", testPrivateKey)
	apiKeyConfigVariablesToUnset, ok = checkIncompatibleAttrsForApiKeyAuth(d)
	assert.False(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 3, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("private_key_path", "path")
	apiKeyConfigVariablesToUnset, ok = checkIncompatibleAttrsForApiKeyAuth(d)
	assert.False(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 4, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)

	d.Set("private_key_password", "password")
	apiKeyConfigVariablesToUnset, ok = checkIncompatibleAttrsForApiKeyAuth(d)
	assert.False(t, ok)
	assert.True(t, len(apiKeyConfigVariablesToUnset) == 5, "apiKey config variables to unset: %v", apiKeyConfigVariablesToUnset)
}

// This test verifies that user can specify private key paths with "~/" and they should resolve to the home directory
func TestUnitHomeDirectoryPrivateKeyPath_basic(t *testing.T) {
	privateKeyName := "TestUnitHomeDirectoryPrivateKeyPath_basic.pem"
	privateKeyPath := path.Join(getHomeFolder(), privateKeyName)
	err := writeTempFile(testPrivateKey, privateKeyPath)
	if err != nil {
		t.Fatalf("unable to write test private key into directory %s. Error: %v", privateKeyPath, err)
	}

	defer removeFile(privateKeyPath)

	r := &schema.Resource{
		Schema: schemaMap(),
	}
	d := r.Data(nil)
	d.Set(privateKeyPathAttrName, path.Join("~", privateKeyName))

	d.Set(tenancyOcidAttrName, testTenancyOCID)
	d.Set(authAttrName, authAPIKeySetting)
	d.Set(userOcidAttrName, testUserOCID)
	d.Set(fingerprintAttrName, testKeyFingerPrint)
	d.Set(regionAttrName, "us-phoenix-1")

	clients := &OracleClients{
		sdkClientMap:  make(map[string]interface{}, len(oracleClientRegistrations.registeredClients)),
		configuration: make(map[string]string),
	}
	sdkConfigProvider, err := getSdkConfigProvider(d, clients)
	assert.NoError(t, err)

	privateRsaKey, err := sdkConfigProvider.PrivateRSAKey()
	assert.NoError(t, err)
	assert.True(t, privateRsaKey != nil)
}

func TestUnitSecurityToken_basic(t *testing.T) {
	t.Skip("Run manual with a valid security token")
	for _, apiKeyConfigAttribute := range apiKeyConfigAttributes {
		apiKeyConfigAttributeEnvValue := getEnvSettingWithBlankDefault(apiKeyConfigAttribute)
		if apiKeyConfigAttributeEnvValue != "" {
			t.Skip("apiKeyConfigAttributes are set through environment variables, skip the test")
		}
	}
	r := &schema.Resource{
		Schema: schemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", authSecurityToken)
	d.Set(configFileProfileAttrName, "PROFILE4") // Run CLI command "oci session authenticate" to get token and profile
	clients := &OracleClients{
		sdkClientMap:  make(map[string]interface{}, len(oracleClientRegistrations.registeredClients)),
		configuration: make(map[string]string),
	}
	sdkConfigProvider, err := getSdkConfigProvider(d, clients)
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
	client, err := oci_budget.NewBudgetClientWithConfigurationProvider(sdkConfigProvider)
	assert.NoError(t, err)
	assert.NotEmpty(t, client.Host)

	request := oci_budget.ListBudgetsRequest{}
	compartmentId, ok := d.GetOk("compartment_id")
	assert.True(t, ok)
	compartmentIdString := compartmentId.(string)

	request.CompartmentId = &compartmentIdString

	_, err = client.ListBudgets(context.Background(), request)
	assert.NoError(t, err)
}

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
