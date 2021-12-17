// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourcePrivateIPTestSuite struct {
	suite.Suite
	Providers        map[string]terraform.ResourceProvider
	Config           string
	ResourceName     string
	VlanConfig       string
	VlanResourceName string
}

func (s *ResourcePrivateIPTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + tfresource.TestADs() + tfresource.TestVCN1() + tfresource.TestSubnet1() + tfresource.TestImage1() + tfresource.TestInstance1() + DefinedTagsDependencies + `
	data "oci_core_vnic_attachments" "t" {
		availability_domain = "${data.oci_identity_availability_domains.t.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		instance_id = "${oci_core_instance.t.id}"
	}`
	s.VlanConfig = acctest.LegacyTestProviderConfig() +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Required, acctest.Create, vlanRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		AvailabilityDomainConfig + DefinedTagsDependencies

	s.ResourceName = "oci_core_private_ip.t"
	s.VlanResourceName = "oci_core_private_ip.tpvlan"
}

func (s *ResourcePrivateIPTestSuite) TestAccCoreResourcePrivateIP_basic() {
	var resId string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test Create
			{
				Config: s.Config + `
				resource "oci_core_private_ip" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}"
					display_name = "-private-ip"
					defined_tags = "${map(
									"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value"
									)}"
                    freeform_tags = { "Department" = "Finance"}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					// hostname_label should not be set unless explicitly set
					resource.TestCheckNoResourceAttr(s.ResourceName, "hostname_label"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-private-ip"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_core_private_ip.t", "id")
						return
					},
					resource.TestCheckResourceAttr(s.ResourceName, "freeform_tags.%", "1"),
				),
			},
			// test Update
			{
				Config: s.Config + `
				resource "oci_core_private_ip" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}"
					display_name = "-private-ip2"
					defined_tags = "${map(
									"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue"
									)}"
                    freeform_tags = { "Department" = "Accounting"}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-private-ip2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					// hostname_label should not be set unless explicitly set
					resource.TestCheckNoResourceAttr(s.ResourceName, "hostname_label"),
					resource.TestCheckResourceAttr(s.ResourceName, "freeform_tags.%", "1"),
				),
			},
			// test add host name label
			{
				Config: s.Config + `
				resource "oci_core_private_ip" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}"
					display_name = "-private-ip2"
					hostname_label = "ahostname"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-private-ip2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "ahostname"),
				),
			},
			// test destructive ip address change
			{
				Config: s.Config + `
				resource "oci_core_private_ip" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}"
					display_name = "-private-ip2"	
					hostname_label = "ahostname"
					ip_address = "10.0.1.22"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-private-ip2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "ip_address", "10.0.1.22"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "ahostname"),
					func(s *terraform.State) (err error) {
						resId2, err := acctest.FromInstanceState(s, "oci_core_private_ip.t", "id")
						if resId == resId2 {
							return fmt.Errorf("Expected new private_ip ocid, got the same")
						}
						return err
					},
				),
			},
		},
	})
}

func (s *ResourcePrivateIPTestSuite) TestAccCoreResourcePrivateIPVlan_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers:    s.Providers,
		CheckDestroy: testAccCheckCorePrivateIpDestroy,
		Steps: []resource.TestStep{
			// test Create
			{
				Config: s.VlanConfig + `
				resource "oci_core_private_ip" "tpvlan" {
					vlan_id		 = "${oci_core_vlan.test_vlan.id}"
					ip_address	 = "10.0.0.5"
					display_name = "-private-ip"
					defined_tags = "${map(
									"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value"
									)}"
                    freeform_tags = { "Department" = "Finance"}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "id"),
					resource.TestCheckResourceAttr(s.VlanResourceName, "ip_address", "10.0.0.5"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "is_primary"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "vlan_id"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.VlanResourceName, "display_name", "-private-ip"),
					resource.TestCheckResourceAttr(s.VlanResourceName, "freeform_tags.%", "1"),
				),
			},
			// test Update
			{
				Config: s.VlanConfig + `
				resource "oci_core_private_ip" "tpvlan" {
					vlan_id		 = "${oci_core_vlan.test_vlan.id}"
					ip_address	 = "10.0.0.10"
					display_name = "-private-ip2"
					defined_tags = "${map(
									"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue"
									)}"
                    freeform_tags = { "Department" = "Accounting"}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.VlanResourceName, "display_name", "-private-ip2"),
					resource.TestCheckResourceAttr(s.VlanResourceName, "ip_address", "10.0.0.10"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "is_primary"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "vlan_id"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.VlanResourceName, "freeform_tags.%", "1"),
				),
			},
		},
	})
}

// issue-routing-tag: core/virtualNetwork
func TestResourceCorePrivateIPTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCorePrivateIPTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourcePrivateIPTestSuite))
}
