// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"errors"
	"testing"

	"os"
	"strconv"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"
	"github.com/stretchr/testify/mock"
)

func testProviderConfig() string {
	return `
	provider "baremetal" {
		tenancy_ocid = "ocid.tenancy.aaaa"
		user_ocid = "ocid.user.bbbbb"
		fingerprint = "xxxxxxxxxx"
		private_key_path = "/home/foo/private_key.pem"
		private_key_password = "password"
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
data "baremetal_identity_availability_domains" "ADs" {
  compartment_id = "${var.compartment_id}"
}

resource "baremetal_core_virtual_network" "t" {
	cidr_block = "10.0.0.0/16"
	compartment_id = "${var.compartment_id}"
	display_name = "network_name"
}

resource "baremetal_core_internet_gateway" "CompleteIG" {
    compartment_id = "${var.compartment_id}"
    display_name = "CompleteIG"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
}

resource "baremetal_core_route_table" "RouteForComplete" {
    compartment_id = "${var.compartment_id}"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
    display_name = "RouteTableForComplete"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${baremetal_core_internet_gateway.CompleteIG.id}"
    }
}

resource "baremetal_core_security_list" "WebSubnet" {
    compartment_id = "${var.compartment_id}"
    display_name = "Public"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
    egress_security_rules = [{
        destination = "0.0.0.0/0"
        protocol = "6"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 80
            "min" = 80
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
	{
	protocol = "6"
	source = "10.0.0.0/16"
    }]
}

resource "baremetal_core_subnet" "WebSubnetAD1" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block = "10.0.1.0/24"
  display_name = "WebSubnetAD1"
  compartment_id = "${var.compartment_id}"
  vcn_id = "${baremetal_core_virtual_network.t.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.WebSubnet.id}"]
}

`

var instanceConfig = subnetConfig + `
data "baremetal_core_images" "t" {
	compartment_id = "${var.compartment_id}"
	limit = 1
}

data "baremetal_core_shape" "shapes" {
	compartment_id = "${var.compartment_id}"
	availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
	image_id = "${data.baremetal_core_images.t.images.0.id}"
}

resource "baremetal_core_instance" "t" {
	availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	display_name = "instance_name"
      	image = "${data.baremetal_core_images.t.images.0.id}"
      	shape = "${data.baremetal_core_shape.shapes.shapes.0.name}"
      	subnet_id = "${baremetal_core_subnet.WebSubnetAD1.id}"
      	metadata {
        	ssh_authorized_keys = "${var.ssh_public_key}"
      	}

      	timeouts {
      		create = "60m"
      	}
}
`

var certificateConfig = `
resource "baremetal_load_balancer_certificate" "t" {
  load_balancer_id   = "${baremetal_load_balancer.t.id}"
  ca_certificate     = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"
  certificate_name   = "stub_certificate_name"
  private_key        = "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----"
  public_certificate = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"
}
`

var loadbalancerConfig = subnetConfig + `

resource "baremetal_core_subnet" "WebSubnetAD2" {
  availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.1.name}"
  cidr_block = "10.0.2.0/24"
  display_name = "WebSubnetAD2"
  compartment_id = "${var.compartment_id}"
  vcn_id = "${baremetal_core_virtual_network.t.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.WebSubnet.id}"]
}

data "baremetal_load_balancer_shapes" "t" {
  compartment_id = "${var.compartment_id}"
}
resource "baremetal_load_balancer" "t" {
  shape          = "${data.baremetal_load_balancer_shapes.t.shapes.0.name}"
  compartment_id = "${var.compartment_id}"
  display_name   = "lb_display_name"
  subnet_ids     = ["${baremetal_core_subnet.WebSubnetAD1.id}", "${baremetal_core_subnet.WebSubnetAD2.id}"]
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

	resource "baremetal_database_db_system" "t" {
	  availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
	  compartment_id = "${var.compartment_id}"
	  cpu_core_count = "${var.CPUCoreCount}"
	  database_edition = "${var.DBEdition}"
	  db_home {
	    database {
	      "admin_password" = "${var.DBAdminPassword}"
	      "db_name" = "${var.DBName}"
	    }
	    db_version = "${var.DBVersion}"
	    display_name = "MyTFDB"
	  }
	  disk_redundancy = "NORMAL"
	  shape = "${var.DBNodeShape}"
	  subnet_id = "${baremetal_core_subnet.WebSubnetAD1.id}"
	  ssh_public_keys = ["${var.ssh_public_key}"]
	  display_name = "MyTFDatabaseNode0"
	  domain = "${var.DBNodeDomainName}"
	  hostname = "${var.DBNodeHostName}"
	}
	`

// This is a dummy object allowing coexistance between mocked API calls and real API calls in acceptance tests
// Acceptance tests will use this object that "mocks" the mocks
type mockableClient interface {
	client.BareMetalClient
	On(methodName string, arguments ...interface{}) *mock.Call
	AssertCalled(t mock.TestingT, methodName string, arguments ...interface{}) bool
}

type testClient struct {
	client.BareMetalClient
}

func (r *testClient) On(methodName string, arguments ...interface{}) *mock.Call {
	// Do Nothing. Return this object so mocks continue to work
	return &mock.Call{Parent: &mock.Mock{}}
}
func (r *testClient) AssertCalled(t mock.TestingT, methodName string, arguments ...interface{}) bool {
	// Do Nothing. Just return true and assume errors are caught elsewhere
	return true
}

func IsAccTest() bool {
	val := os.Getenv(resource.TestEnvVar)
	if val == "" {
		return false
	}
	acc, err := strconv.ParseBool(val)
	if err != nil {
		panic("Err testing TF_ACC env var. It should be blank or a boolean value.")
	}
	return acc
}

func GetTestProvider() mockableClient {
	if IsAccTest() {
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

		client, err := providerConfig(d)
		if err != nil {
			panic(err)
		}
		return &testClient{client.(*baremetal.Client)}
	}
	return &mocks.BareMetalClient{}
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
var testPublicKey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtBLQAGmKJ7tpfzYJyqLG
ZDwHL51+d6T8Z00BnP9CFfzxZZZ48PcYSUHuTyCM8mR5JqYLyH6C8tZ/DKqwxUnc
ONgBytG3MM42bgxfHIhsZRj5rCz1oqWlSLuXvgww1kuqWnt6r+NtnXog439YsGTH
RotrTLTdEgOxH0EFP5uHUc9w/Uix7rWU7GB2ra060oeTB/hKpts5U70eI2EI6ec9
1sJdUIj7xNfBJeQQrz4CFUrkyzL06211CFvhmxH2hA9gBKOqC3rGL8XraHZBhGWn
mXlrQB7nNKsJrrv5fHwaPDrAY4iNP2W0q3LRpyNigJ6cgRuGJhHa82iHPmxgIx8m
fwIDAQAB
-----END PUBLIC KEY-----`

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
