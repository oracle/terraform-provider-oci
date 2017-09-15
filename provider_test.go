// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"errors"
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/assert"
	"fmt"
)

var testAccClient *baremetal.Client
var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccClient = GetTestProvider()
	
	testAccProvider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return testAccClient, nil
	}).(*schema.Provider)
	
	testAccProviders = map[string]terraform.ResourceProvider{
		"oci": testAccProvider,
	}
}

func testProviderConfig() string {
	return `
	provider "oci" {
		tenancy_ocid = "ocid.tenancy.aaaa"
		user_ocid = "ocid.user.bbbbb"
		fingerprint = "xxxxxxxxxx"
		private_key_path = "/home/foo/private_key.pem"
		private_key_password = "password"
		region = "us-phoenix-1"
	}

	variable "compartment_id" {
		default = "` + getEnvSetting("compartment_id", "compartment_id") + `"
	}

	variable "tenancy_ocid" {
		default = "` + getRequiredEnvSetting("tenancy_ocid") + `"
	}

	variable "namespace" {
		default = "` + getEnvSetting("namespace", "mustwin") + `"
	}

	variable "ssh_public_key" {
		default = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
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

var imageConf = `
data "oci_core_images" "t" {
	compartment_id = "${var.compartment_id}"
  	operating_system = "Oracle Linux"
  	operating_system_version = "7.3"
  	limit = 1
}`

var instanceConfig = subnetConfig + `
data "oci_core_images" "t" {
	compartment_id = "${var.compartment_id}"
  	operating_system = "Oracle Linux"
  	operating_system_version = "7.3"
  	limit = 1
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
	image_id = "${data.oci_core_images.t.images.0.id}"
}

resource "oci_core_instance" "t" {
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	display_name = "-tf-instance"
	image = "${data.oci_core_images.t.images.0.id}"
	shape = "VM.Standard1.1"
	subnet_id = "${oci_core_subnet.WebSubnetAD1.id}"
	metadata {
		ssh_authorized_keys = "${var.ssh_public_key}"
	}

	timeouts {
		create = "15m"
	}
}
`
var vnicConfig = instanceConfig + `
data "oci_core_vnic_attachments" "vnics" {
    compartment_id = "${var.compartment_id}"
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
    instance_id = "${oci_core_instance.t.id}"
}
`

var instanceDnsConfig = `
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

data "oci_core_images" "t" {
	compartment_id = "${var.compartment_id}"
  	operating_system = "Oracle Linux"
  	operating_system_version = "7.3"
  	limit = 1
}

resource "oci_core_instance" "t" {
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	display_name = "-tf-instance"
	image = "${data.oci_core_images.t.images.0.id}"
	shape = "VM.Standard1.8"
	create_vnic_details {
        subnet_id = "${oci_core_subnet.t.id}"
        hostname_label = "testinstance"
        display_name = "-tf-instance-vnic"
  	}
	metadata {
		ssh_authorized_keys = "${var.ssh_public_key}"
	}
	timeouts {
		create = "15m"
	}
}
`

var certificateConfig = `
resource "oci_load_balancer_certificate" "t" {
  load_balancer_id   = "${oci_load_balancer.t.id}"
  ca_certificate     = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"
  certificate_name   = "stub_certificate_name"
  private_key        = "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----"
  public_certificate = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"
}
`

var loadbalancerConfig = subnetConfig + `

resource "oci_core_subnet" "WebSubnetAD2" {
  availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.1.name}"
  cidr_block = "10.0.2.0/24"
  display_name = "WebSubnetAD2"
  compartment_id = "${var.compartment_id}"
  vcn_id = "${oci_core_virtual_network.t.id}"
  route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
  security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
  dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
}

data "oci_load_balancer_shapes" "t" {
  compartment_id = "${var.compartment_id}"
}
resource "oci_load_balancer" "t" {
  shape          = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
  compartment_id = "${var.compartment_id}"
  display_name   = "lb_display_name"
  subnet_ids     = ["${oci_core_subnet.WebSubnetAD1.id}", "${oci_core_subnet.WebSubnetAD2.id}"]
}
`

var databaseConfig = subnetConfig + `
variable "DBNodeShape" {
    default = "BM.DenseIO1.36"
}

variable "CPUCoreCount" {
    default = "2"
}

variable "DBEdition" {
    default = "ENTERPRISE_EDITION"
}

variable "DBAdminPassword" {
    default = "BEstrO0ng_#11"
}

variable "DBName" {
    default = "aTFdb"
}

variable "DBVersion" {
    default = "12.1.0.2"
}


variable "DBDiskRedundancy" {
    default = "HIGH"
}

variable "DBNodeDisplayName" {
    default = "MyTFDatabaseNode0"
}

variable "DBNodeDomainName" {
    default = "mycompany.com"
}

variable "DBNodeHostName" {
    default = "myOracleDB"
}

	resource "oci_database_db_system" "t" {
	  availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	  compartment_id = "${var.compartment_id}"
	  cpu_core_count = "${var.CPUCoreCount}"
	  database_edition = "${var.DBEdition}"
	  db_home {
	    database {
	      "admin_password" = "${var.DBAdminPassword}"
	      "db_name" = "${var.DBName}"
	      character_set = "AL32UTF8"
	      ncharacter_set = "AL16UTF16"
	    }
	    db_version = "${var.DBVersion}"
	    display_name = "MyTFDB"
	  }
	  disk_redundancy = "NORMAL"
	  shape = "${var.DBNodeShape}"
	  subnet_id = "${oci_core_subnet.WebSubnetAD1.id}"
	  ssh_public_keys = ["${var.ssh_public_key}"]
	  display_name = "MyTFDatabaseNode0"
	  domain = "${var.DBNodeDomainName}"
	  hostname = "${var.DBNodeHostName}"
	}
	`

func GetTestProvider() *baremetal.Client {
	r := &schema.Resource{
		Schema: schemaMap(),
	}
	d := r.Data(nil)
	d.SetId(getRequiredEnvSetting("tenancy_ocid"))

	d.Set("tenancy_ocid", getRequiredEnvSetting("tenancy_ocid"))
	d.Set("user_ocid", getRequiredEnvSetting("user_ocid"))
	d.Set("fingerprint", getRequiredEnvSetting("fingerprint"))
	d.Set("private_key_path", getRequiredEnvSetting("private_key_path"))
	d.Set("private_key_password", getEnvSetting("private_key_password", ""))
	d.Set("private_key", getEnvSetting("private_key", ""))
	d.Set("disable_auto_retries", true)

	client, err := providerConfig(d)
	if err != nil {
		panic(err)
	}
	return client.(*baremetal.Client)
}

// This test runs the Provider sanity checks.
func TestProvider(t *testing.T) {

	// Real client for the sanity check. Makes this more of an acceptance test.
	client := &baremetal.Client{}
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

func TestProviderConfig(t *testing.T) {
	r := &schema.Resource{
		Schema: schemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")

	d.Set("tenancy_ocid", testTenancyOCID)
	d.Set("user_ocid", testUserOCID)
	d.Set("fingerprint", testKeyFingerPrint)
	d.Set("private_key", testPrivateKey)
	//d.Set("private_key_path", "")
	d.Set("private_key_password", "password")

	client, err := providerConfig(d)
	assert.Nil(t, err)
	assert.NotNil(t, client)
	_, ok := client.(*baremetal.Client)
	assert.True(t, ok)
}

// TestNoInstanceState determines if there is any state for a given name.
func testNoInstanceState(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ms := s.RootModule()
		rs, ok := ms.Resources[name]
		if !ok {
			return nil
		}

		is := rs.Primary
		if is == nil {
			return nil
		}

		return errors.New("State exists for primary resource " + name)
	}
}

// custom TestCheckFunc helper, returns a value associated with a key from an instance in the current state
func fromInstanceState(s *terraform.State, name, key string) (string, error) {
	ms := s.RootModule()
	rs, ok := ms.Resources[name]
	if !ok {
		return "", fmt.Errorf("Not found: %s", name)
	}

	is := rs.Primary
	if is == nil {
		return "", fmt.Errorf("No primary instance: %s", name)
	}

	v, ok := is.Attributes[key];

	if ok {
		return v, nil
	} else {
		return "", fmt.Errorf("%s: Attribute '%s' not found", name, key)
	}
}
