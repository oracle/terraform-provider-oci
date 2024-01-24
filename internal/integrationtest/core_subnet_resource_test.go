// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/core"
)

// issue-routing-tag: core/virtualNetwork
func TestAccResourceCoreSubnetCreate_basic(t *testing.T) {
	httpreplay.SetScenario("TestAccResourceCoreSubnetCreate_basic")
	defer httpreplay.SaveScenario()
	acctest.PreCheck(t)
	config := acctest.LegacyTestProviderConfig() + `
		data "oci_identity_availability_domains" "ADs" {
			compartment_id = "${var.compartment_id}"
		}

		resource "oci_core_virtual_network" "t" {
			cidr_block     = "10.0.0.0/16"
			compartment_id = "${var.compartment_id}"
			display_name   = "network_name"
			dns_label      = "myvcn"
		}

		resource "oci_core_security_list" "seclist1" {
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-security_list1"
			vcn_id = "${oci_core_virtual_network.t.id}"
		}

		resource "oci_core_security_list" "seclist2" {
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-security_list2"
			vcn_id = "${oci_core_virtual_network.t.id}"
		}`

	commonSubnetParams := `
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					route_table_id = "${oci_core_virtual_network.t.default_route_table_id}"
					dhcp_options_id = "${oci_core_virtual_network.t.default_dhcp_options_id}"
					cidr_block = "10.0.2.0/24"`

	singleSecurityListId := `
					security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]`

	extraSecurityListIds := `
					security_list_ids = [
						"${oci_core_virtual_network.t.default_security_list_id}",
						"${oci_core_security_list.seclist1.id}",
						"${oci_core_security_list.seclist2.id}"
					]`

	reorderedSecurityListIds := `
					security_list_ids = [
						"${oci_core_virtual_network.t.default_security_list_id}",
						"${oci_core_security_list.seclist2.id}",
						"${oci_core_security_list.seclist1.id}",
					]`

	resourceName := "oci_core_subnet.s"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + `
				resource "oci_core_subnet" "s" {` + commonSubnetParams + extraSecurityListIds + `
				}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestMatchResourceAttr(resourceName, "vcn_id", regexp.MustCompile("ocid1\\.vcn\\.oc1\\..*")),
				resource.TestMatchResourceAttr(resourceName, "dhcp_options_id", regexp.MustCompile("ocid1\\.dhcpoptions\\.oc1\\..*")),
				resource.TestMatchResourceAttr(resourceName, "route_table_id", regexp.MustCompile("ocid1\\.routetable\\.oc1\\..*")),
				resource.TestCheckResourceAttr(resourceName, "security_list_ids.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.2.0/24"),
				resource.TestCheckNoResourceAttr(resourceName, "dns_label"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestMatchResourceAttr(resourceName, "id", regexp.MustCompile("ocid1\\.subnet\\.oc1\\..*")),
				resource.TestCheckResourceAttr(resourceName, "state", string(core.SubnetLifecycleStateAvailable)),
				// TODO: Add test for scenario where subnet_domain_name is set?
				resource.TestCheckNoResourceAttr(resourceName, "subnet_domain_name"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update
		{
			Config: config + `
				resource "oci_core_subnet" "s" {
					` + commonSubnetParams + extraSecurityListIds + `
					display_name = "-tf-subnet"
				}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "-tf-subnet"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Expected same subnet ocid, got the different.")
					}
					return err
				},
			),
		},
		// verify no diffs when reordering security list IDs
		{
			Config: config + `
				resource "oci_core_subnet" "s" {
					` + reorderedSecurityListIds + commonSubnetParams + `
					display_name = "-tf-subnet"
				}`,
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},
		// test a destructive Update results in a new resource
		{
			Config: config + `
				resource "oci_core_subnet" "s" {
					` + commonSubnetParams + singleSecurityListId + `
					display_name = "-tf-subnet"
					prohibit_public_ip_on_vnic = "true"
					dns_label = "MyTestLabel"
				}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "true"),
				resource.TestCheckResourceAttr(resourceName, "dns_label", "mytestlabel"),
				func(s *terraform.State) (err error) {
					resId3, err := acctest.FromInstanceState(s, resourceName, "id")
					if resId2 == resId3 {
						return fmt.Errorf("Expected new subnet ocid, got the same.")
					}
					return err
				},
			),
		},
		// DNS capitalization changes should be ignored.
		{
			Config: config + `
				resource "oci_core_subnet" "s" {
					` + commonSubnetParams + singleSecurityListId + `
					display_name = "-tf-subnet"
					prohibit_public_ip_on_vnic = "true"
					dns_label = "mytestlabel"
				}`,
			ExpectNonEmptyPlan: false,
			PlanOnly:           true,
		},
		// DNS label change should cause a change
		{
			Config: config + `
				resource "oci_core_subnet" "s" {
					` + commonSubnetParams + singleSecurityListId + `
					display_name = "-tf-subnet"
					prohibit_public_ip_on_vnic = "true"
					dns_label = "NewLabel"
				}`,
			ExpectNonEmptyPlan: true,
			PlanOnly:           true,
		},
	})
}
