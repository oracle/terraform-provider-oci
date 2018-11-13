// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"net/http"
	"runtime"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	oci_common "github.com/oracle/oci-go-sdk/common"
	"github.com/stretchr/testify/assert"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider
var requiredTestEnvVars = []string{"compartment_ocid", "compartment_id_for_create", "compartment_id_for_update", "tags_import_if_exists"}
var requiredKeyAuthEnvVars = []string{"tenancy_ocid", "user_ocid", "fingerprint"}
var requiredOboTokenAuthEnvVars = []string{"tenancy_ocid", "obo_token"}

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
		DataSourcesMap: dataSourcesMap(),
		Schema:         schemaMap(),
		ResourcesMap:   resourcesMap(),
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
	}`
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
	us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"
	us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
	eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
	uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
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
	shape = "VM.Standard1.8"
	subnet_id = "${oci_core_subnet.WebSubnetAD1.id}"
	metadata {
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
	shape = "VM.Standard1.8"
	create_vnic_details {
        subnet_id = "${oci_core_subnet.t.id}"
        hostname_label = "testinstance"
        display_name = "-tf-instance-vnic"
		defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
		freeform_tags = { "Department" = "Accounting" }
  	}
	metadata {
		ssh_authorized_keys = "${var.ssh_public_key}"
	}
	timeouts {
		create = "15m"
	}
}
` + DefinedTagsDependencies
)

const (
	requestQueryOpcTimeMaintenanceRebootDue = "opc-time-maintenance-reboot-due"
)

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
		computeClient := client.(*OracleClients).computeClient
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
func TestProvider(t *testing.T) {
	// Real client for the sanity check. Makes this more of an acceptance test.
	client := &OracleClients{}
	if err := Provider(func(d *schema.ResourceData) (interface{}, error) {
		return client, nil
	}).(*schema.Provider).InternalValidate(); err != nil {
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

func providerConfigTest(t *testing.T, disableRetries bool, skipRequiredField bool, auth string) {
	r := &schema.Resource{
		Schema: schemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", auth)
	if !skipRequiredField {
		d.Set("tenancy_ocid", testTenancyOCID)
	}
	d.Set("user_ocid", testUserOCID)
	d.Set("fingerprint", testKeyFingerPrint)
	d.Set("private_key", testPrivateKey)
	//d.Set("private_key_path", "")
	d.Set("private_key_password", "password")
	d.Set("region", "us-phoenix-1")

	if disableRetries {
		d.Set("disable_auto_retries", disableRetries)
	}

	client, err := ProviderConfig(d)

	switch auth {
	case authAPIKeySetting, "":
		if skipRequiredField {
			assert.Error(t, err, fmt.Sprintf("when auth is set to '%s', tenancy_ocid, user_ocid, and fingerprint are required", authAPIKeySetting))
			return
		}
	case authInstancePrincipalSetting:
		assert.Regexp(t, "failed to create a new key provider for instance principal.*", err.Error())
		return
	case authInstancePrincipalWithCertsSetting:
		assert.Regexp(t, "failed to create a new key provider for instance principal.*", err.Error())
		return
	default:
		assert.Error(t, err, fmt.Sprintf("auth must be one of '%s' or '%s' or '%s'", authAPIKeySetting, authInstancePrincipalSetting, authInstancePrincipalWithCertsSetting))
		return
	}
	assert.Nil(t, err)
	assert.NotNil(t, client)

	oracleClient, ok := client.(*OracleClients)
	assert.True(t, ok)

	userAgent := fmt.Sprintf(userAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, terraform.VersionString(), defaultUserAgentProviderName, Version)
	testClient := func(c *oci_common.BaseClient) {
		assert.NotNil(t, c)
		assert.NotNil(t, c.HTTPClient)
		assert.Exactly(t, c.UserAgent, userAgent)
		assert.NotNil(t, c.Interceptor)
		assert.NotNil(t, c.Signer)
	}

	testClient(&oracleClient.blockstorageClient.BaseClient)
	testClient(&oracleClient.computeClient.BaseClient)
	testClient(&oracleClient.databaseClient.BaseClient)
	testClient(&oracleClient.identityClient.BaseClient)
	testClient(&oracleClient.virtualNetworkClient.BaseClient)
	testClient(&oracleClient.objectStorageClient.BaseClient)
	testClient(&oracleClient.loadBalancerClient.BaseClient)
}

func TestProviderConfig(t *testing.T) {
	providerConfigTest(t, true, true, authAPIKeySetting)              // ApiKey with required fields + disable auto-retries
	providerConfigTest(t, false, true, authAPIKeySetting)             // ApiKey without required fields
	providerConfigTest(t, false, false, authInstancePrincipalSetting) // InstancePrincipal
	providerConfigTest(t, true, false, "invalid-auth-setting")        // Invalid auth + disable auto-retries
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
