// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreBootVolumeTestSuite struct {
	suite.Suite
	Providers    map[string]*schema.Provider
	Config       string
	ResourceName string
}

func (s *ResourceCoreBootVolumeTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "test_availability_domains" {
		compartment_id = "${var.tenancy_ocid}"
	}

	resource "oci_core_dhcp_options" "test_dhcp_options" {
		compartment_id = "${var.compartment_id}"
		options {
			server_type = "VcnLocalPlusInternet"
			type = "DomainNameServer"
		}
		options {
			search_domain_names = ["test.com"]
			type = "SearchDomain"
		}
		vcn_id = "${oci_core_vcn.test_vcn.id}"
	}
	
	resource "oci_core_vcn" "test_vcn" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-vcn"
		dns_label = "dnslabel"
		freeform_tags = {
			"Department" = "Finance"
		}
	}

    resource "oci_core_drg" "test_drg" {
        #Required
        compartment_id = "${var.compartment_id}"
		display_name = "-tf-drg"
    }
    
    data "oci_core_services" "test_services" {}
	
	resource "oci_core_internet_gateway" "test_network_entity" {
    	compartment_id = "${var.compartment_id}"
    	vcn_id = "${oci_core_vcn.test_vcn.id}"
    	display_name = "-tf-internet-gateway"
	}

	resource "oci_core_route_table" "test_route_table" {
		compartment_id = "${var.compartment_id}"
		route_rules {
			cidr_block = "0.0.0.0/0"
			network_entity_id = "${oci_core_internet_gateway.test_network_entity.id}"
		}
		vcn_id = "${oci_core_vcn.test_vcn.id}"
	}

	resource "oci_core_subnet" "test_subnet" {
		availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		dhcp_options_id = "${oci_core_dhcp_options.test_dhcp_options.id}"
		display_name = "-tf-subnet"
		dns_label = "dnslabel"
		freeform_tags = {
			"Department" = "Accounting"
		}
		prohibit_public_ip_on_vnic = "false"
		route_table_id = "${oci_core_route_table.test_route_table.id}"
		security_list_ids = ["${oci_core_vcn.test_vcn.default_security_list_id}"]
		vcn_id = "${oci_core_vcn.test_vcn.id}"
	}

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
	
	resource "oci_core_instance" "test_instance" {
		availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
		compartment_id = "${var.compartment_id}"
		image = "${var.InstanceImageOCID[var.region]}"
		shape = "VM.Standard2.1"
		subnet_id = "${oci_core_subnet.test_subnet.id}"
		display_name = "-tf-instance"
	}
	`

	s.ResourceName = "oci_core_boot_volume.test_boot_volume"
}

func (s *ResourceCoreBootVolumeTestSuite) TestResourceCoreBootVolume_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: s.Config +
					`
						resource "oci_core_boot_volume" "test_boot_volume" {
							availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
							compartment_id = "${var.compartment_id}"
							source_details {
								id = "${oci_core_instance.test_instance.boot_volume_id}"
								type = "bootVolume"
							}
							display_name = "-tf-bootVolume-clone"
							size_in_gbs = "51"
						}

						resource "oci_core_instance" "test_instance_new" {
							availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
							compartment_id = "${var.compartment_id}"
							shape = "VM.Standard2.1"
							subnet_id = "${oci_core_subnet.test_subnet.id}"
							source_details {
								source_id = "${oci_core_boot_volume.test_boot_volume.id}"
								source_type = "bootVolume"
							}
							display_name = "-tf-instance-2"
						}
					`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-bootVolume-clone"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					func(ts *terraform.State) (err error) {
						_, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			{
				Config: s.Config +
					`
						resource "oci_core_boot_volume" "test_boot_volume" {
							availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
							compartment_id = "${var.compartment_id}"
							source_details {
								id = "${oci_core_instance.test_instance.boot_volume_id}"
								type = "bootVolume"
							}
							display_name = "-tf-bootVolume-2"
							size_in_gbs = "51"		
						}

						resource "oci_core_instance" "test_instance_new" {
							availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
							compartment_id = "${var.compartment_id}"
							shape = "VM.Standard2.1"
							subnet_id = "${oci_core_subnet.test_subnet.id}"
							source_details {
								source_id = "${oci_core_boot_volume.test_boot_volume.id}"
								source_type = "bootVolume"
							}
							display_name = "-tf-instance-2"
						}
					`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-bootVolume-2"),
					func(ts *terraform.State) (err error) {
						_, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
		},
	})
}

// issue-routing-tag: core/blockStorage
func TestResourceCoreBootVolumeTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreBootVolumeTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreBootVolumeTestSuite))
}
