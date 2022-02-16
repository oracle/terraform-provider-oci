// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/core"
	"github.com/stretchr/testify/suite"
)

var (
	vnicAttachmentRepresentationVlan = map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: vnicAttachmentCreateVnicDetailsVlanRepresentation},
		"instance_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"nic_index":           acctest.Representation{RepType: acctest.Optional, Create: `0`},
	}

	vnicAttachmentCreateVnicDetailsVlanRepresentation = map[string]interface{}{
		"assign_public_ip": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Accounting"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"vlan_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_vlan.id}`},
	}

	VnicAttachmentResourceDependenciesVlan = acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Required, acctest.Create,
			acctest.GetUpdatedRepresentationCopy("cidr_block", acctest.Representation{RepType: acctest.Required, Create: `10.0.1.0/30`}, vlanRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`}})) +
		AvailabilityDomainConfig
)

type ResourceCoreVnicAttachmentTestSuite struct {
	suite.Suite
	Providers        map[string]terraform.ResourceProvider
	Config           string
	ResourceName     string
	VnicResourceName string
	VlanResourceName string
	VlanDataSource   string
}

func (s *ResourceCoreVnicAttachmentTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + instanceDnsConfig
	s.ResourceName = "oci_core_vnic_attachment.va"
	s.VnicResourceName = "data.oci_core_vnic.v"
	s.VlanResourceName = "oci_core_vnic_attachment.test_vnic_attachment"
}

func (s *ResourceCoreVnicAttachmentTestSuite) TestAccResourceCoreVnicAttachment_basic() {

	var vaId string

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + `
					resource "oci_core_vnic_attachment" "va" {
						instance_id = "${oci_core_instance.t.id}"
						display_name = "-tf-va1"
						create_vnic_details {
							subnet_id = "${oci_core_subnet.t.id}"
							assign_public_ip = false
							defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
							freeform_tags = { "Department" = "Accounting" }
							nsg_ids = ["${oci_core_network_security_group.test_network_security_group1.id}"]
						}
					}
					data "oci_core_vnic" "v" {
						vnic_id = "${oci_core_vnic_attachment.va.vnic_id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-va1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VnicAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vlan_tag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.nsg_ids.#", "1"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "private_ip_address"),
					// @SDK 1/2018: Since we don't assign a public IP to this vnic, we will get a response from server
					// without a public_ip_address. Old SDK would have set it to empty, but new SDK will set it to nil.
					// Commenting out until we have a better way of handling this.
					//resource.TestCheckResourceAttr(s.VnicResourceName, "public_ip_address", ""),
					resource.TestCheckNoResourceAttr(s.VnicResourceName, "public_ip_address"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "skip_source_dest_check", "false"),
					func(ts *terraform.State) (err error) {
						vaId, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			{
				// Update the VNIC
				Config: s.Config + `
					resource "oci_core_vnic_attachment" "va" {
						instance_id = "${oci_core_instance.t.id}"
						display_name = "-tf-va1"
						create_vnic_details {
							subnet_id = "${oci_core_subnet.t.id}"
							display_name = "-tf-vnic-2"
							assign_public_ip = false
							hostname_label = "myvnichostname"
							skip_source_dest_check = true
							defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}"
							freeform_tags = { "Department" = "Finance" }
							nsg_ids = ["${oci_core_network_security_group.test_network_security_group1.id}", "${oci_core_network_security_group.test_network_security_group2.id}"]
						}
					}
					data "oci_core_vnic" "v" {
						vnic_id = "${oci_core_vnic_attachment.va.vnic_id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-va1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VnicAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vlan_tag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.display_name", "-tf-vnic-2"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.nsg_ids.#", "2"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "private_ip_address"),
					// @SDK 1/2018: Since we don't assign a public IP to this vnic, we will get a response from server
					// without a public_ip_address. Old SDK would have set it to empty, but new SDK will set it to nil.
					// Commenting out until we have a better way of handling this.
					//resource.TestCheckResourceAttr(s.VnicResourceName, "public_ip_address", ""),
					resource.TestCheckNoResourceAttr(s.VnicResourceName, "public_ip_address"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "true"),
					func(ts *terraform.State) (err error) {
						newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
						if newId != vaId {
							return fmt.Errorf("Expected same ocid, got different.")
						}
						return err
					},
				),
			},
			{
				// Create a new VNIC and VNIC Attachment with different options.
				Config: s.Config + `
						resource "oci_core_vnic_attachment" "va" {
							instance_id = "${oci_core_instance.t.id}"
							display_name = "-tf-va1"
							create_vnic_details {
								subnet_id = "${oci_core_subnet.t.id}"
								display_name = "-tf-vnic"
								assign_public_ip = true
								private_ip = "10.0.1.20"
								hostname_label = "myvnichostname"
								skip_source_dest_check = true
							}
						}
						data "oci_core_vnic" "v" {
						  vnic_id = "${oci_core_vnic_attachment.va.vnic_id}"
						}
					`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VnicAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "id"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "private_ip_address", "10.0.1.20"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "public_ip_address"),
					resource.TestMatchResourceAttr(s.VnicResourceName, "public_ip_address", regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+\.[0-9]`)),
					resource.TestCheckResourceAttr(s.VnicResourceName, "hostname_label", "myvnichostname"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "skip_source_dest_check", "true"),
					func(ts *terraform.State) (err error) {
						newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
						if newId == vaId {
							return fmt.Errorf("Expected new ocid, got the same.")
						}
						vaId = newId
						return err
					},
				),
			},
			{
				// Switching skip_source_dest_check and assign_public_ip from true to "true" will destroy and recreate, but should result in a
				// VNIC with the same value.
				Config: s.Config + `
						resource "oci_core_vnic_attachment" "va" {
							instance_id = "${oci_core_instance.t.id}"
							display_name = "-tf-va1"
							create_vnic_details {
								subnet_id = "${oci_core_subnet.t.id}"
								display_name = "-tf-vnic"
								assign_public_ip = "true"
								private_ip = "10.0.1.20"
								hostname_label = "myvnichostname"
								skip_source_dest_check = "true"
							}
						}
						data "oci_core_vnic" "v" {
						  vnic_id = "${oci_core_vnic_attachment.va.vnic_id}"
						}
					`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VnicAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttr(s.VnicResourceName, "private_ip_address", "10.0.1.20"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "skip_source_dest_check", "true"),
					func(ts *terraform.State) (err error) {
						newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
						if newId != vaId {
							return fmt.Errorf("Expected same ocid, got different.")
						}
						return err
					},
				),
			},
			{
				Config: s.Config +
					acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", acctest.Required, acctest.Create, vnicAttachmentRepresentationVlan) +
					`data "oci_core_vnic" "v" {
						vnic_id = "${oci_core_vnic_attachment.test_vnic_attachment.vnic_id}"
					}` + VnicAttachmentResourceDependenciesVlan,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "instance_id"),
					resource.TestCheckResourceAttr(s.VlanResourceName, "state", string(core.VnicAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "vlan_id"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "vlan_tag"),
					resource.TestCheckResourceAttr(s.VlanResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "skip_source_dest_check", "true"),
					func(ts *terraform.State) (err error) {
						vaId, err = acctest.FromInstanceState(ts, s.VlanResourceName, "id")
						return err
					},
				),
			},

			{
				Config: s.Config + VnicAttachmentResourceDependenciesVlan,
			},

			{
				Config: s.Config +
					acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", acctest.Optional, acctest.Create, vnicAttachmentRepresentationVlan) +
					`data "oci_core_vnic" "v" {
						vnic_id = "${oci_core_vnic_attachment.test_vnic_attachment.vnic_id}"
					}` + VnicAttachmentResourceDependenciesVlan,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.VlanResourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "instance_id"),
					resource.TestCheckResourceAttr(s.VlanResourceName, "state", string(core.VnicAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "vlan_id"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.VlanResourceName, "vlan_tag"),
					resource.TestCheckResourceAttr(s.VlanResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(s.VlanResourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "skip_source_dest_check", "true"),
					func(ts *terraform.State) (err error) {
						vaId, err = acctest.FromInstanceState(ts, s.VlanResourceName, "id")
						return err
					},
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestResourceCoreVnicAttachmentTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreVnicAttachmentTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreVnicAttachmentTestSuite))
}
