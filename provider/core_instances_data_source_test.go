// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreInstanceTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *DatasourceCoreInstanceTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + s.TokenFn(`
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-vcn"
		cidr_block = "10.0.0.0/16"
	}

	resource "oci_core_subnet" "t" {
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		display_name        = "-tf-subnet"
		cidr_block          = "10.0.1.0/24"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
	}

	variable "InstanceImageOCID" {
	  type = "map"
	  default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
	  }
	}

	resource "oci_core_instance" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "{{.token}}"
		subnet_id = "${oci_core_subnet.t.id}"
		image = "${var.InstanceImageOCID[var.region]}"
		shape = "VM.Standard1.1"
		metadata {
			ssh_authorized_keys = "${var.ssh_public_key}"
		}
		timeouts {
			create = "15m"
		}
	}`, nil)

	s.ResourceName = "data.oci_core_instances.t"
}

func (s *DatasourceCoreInstanceTestSuite) TestAccDatasourceCoreInstance_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + s.TokenFn(`
				data "oci_core_instances" "t" {
					compartment_id = "${var.compartment_id}"
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					display_name = "{{.token}}"
					filter {
						name = "id"
						values = ["${oci_core_instance.t.id}"]
					}
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.display_name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.state", string(core.InstanceLifecycleStateRunning)),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.shape", "VM.Standard1.1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.region"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.image"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.metadata.%"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.metadata.ssh_authorized_keys"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.ipxe_script", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.extended_metadata.%", "0"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "instances.0.create_vnic_details"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.subnet_id", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.private_ip", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.public_ip", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.hostname_label", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.launch_mode", "NATIVE"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.launch_options.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.launch_options.0.boot_volume_type", "ISCSI"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.launch_options.0.network_type", "VFIO"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
				),
			},
			// Check that the optional "state" field can be queried on
			{
				Config: s.Config + s.TokenFn(`
					data "oci_core_instances" "t" {
						compartment_id = "${var.compartment_id}"
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						display_name = "{{.token}}"
						state = "{{.lifecycleState1}}"
						filter {
							name = "id"
							values = ["${oci_core_instance.t.id}"]
						}
					}

					data "oci_core_instances" "t2" {
						compartment_id = "${var.compartment_id}"
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						display_name = "{{.token}}"
						state = "{{.lifecycleState2}}"
						filter {
							name = "id"
							values = ["${oci_core_instance.t.id}"]
						}
					}`,
					map[string]string{
						"lifecycleState1": string(core.InstanceLifecycleStateRunning),
						"lifecycleState2": string(core.InstanceLifecycleStateTerminated),
					},
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.state", string(core.InstanceLifecycleStateRunning)),
					resource.TestCheckResourceAttr("data.oci_core_instances.t2", "instances.#", "0"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreInstanceTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreInstanceTestSuite))
}
