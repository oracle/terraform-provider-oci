// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeAttachmentTestSuite struct {
	suite.Suite
	Providers    map[string]*schema.Provider
	Config       string
	ResourceName [2]string
}

func (s *ResourceCoreVolumeAttachmentTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-vcn"
	}
	
	resource "oci_core_subnet" "t" {
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
	}

	resource "oci_core_instance" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-instance"
		image = "${var.InstanceImageOCID[var.region]}"
		shape = "VM.Standard2.1"
		subnet_id = "${oci_core_subnet.t.id}"
		is_pv_encryption_in_transit_enabled = "true"
		metadata = {
			ssh_authorized_keys = "${var.ssh_public_key}"
		}
		timeouts {
			create = "15m"
		}
	}
	
	resource "oci_core_instance" "t2" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-instance"
		image = "${var.InstanceImageOCID[var.region]}"
		shape = "VM.Standard2.1"
		subnet_id = "${oci_core_subnet.t.id}"
		is_pv_encryption_in_transit_enabled = "true"
		metadata = {
			ssh_authorized_keys = "${var.ssh_public_key}"
		}
		timeouts {
			create = "15m"
		}
	}
	
	resource "oci_core_volume" "t2" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
	}
` + utils.OciImageIdsVariable

	s.ResourceName[0] = "oci_core_volume_attachment.t"
	s.ResourceName[1] = "oci_core_volume_attachment.t2"
}

func (s *ResourceCoreVolumeAttachmentTestSuite) TestResourceCoreVolumeAttachment_basic() {
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: s.Config + `
				
				resource "oci_core_volume" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					display_name = "display_name"
				}

				resource "oci_core_volume_attachment" "t" {
					attachment_type = "iSCSI"	# case-insensitive
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t.id}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_read_only", "false"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "chap_secret"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "chap_username"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "port"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "time_created"),
					func(ts *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(ts, s.ResourceName[0], "id")
						return err
					},
				),
			},
			// ensure that changing the case for attachment_type (polymorphic discriminator) is a no-op.
			{
				Config: s.Config + `
				
				resource "oci_core_volume" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					display_name = "display_name"
				}

				resource "oci_core_volume_attachment" "t" {
					attachment_type = "IscSi"	# case-insensitive
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t.id}"
				}`,
				PlanOnly: true,
			},
			// verify display_name, is_read_only, and use_chap Update forces creation of a new resource
			// verify when display name of attached volume is updates the operation should not error
			{
				Config: s.Config + `
				resource "oci_core_volume" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					display_name = "updated_display_name"
				}

				resource "oci_core_volume_attachment" "t" {
					attachment_type = "IscSi"	# case-insensitive
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t.id}"
					display_name = "tf-vol-attach-upd"
					is_read_only = true
					use_chap = true
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "display_name", "tf-vol-attach-upd"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_read_only", "true"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "use_chap", "true"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "chap_secret"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "chap_username"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "port"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "time_created"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName[0], "id")
						if resId2 == resId {
							return fmt.Errorf("resource not recreated when expected to be when updating display name")
						}
						resId = resId2
						return err
					},
				),
			},
			// verify shared block volume attachment with iscsi attachment type
			{
				Config: s.Config + `
				resource "oci_core_volume_attachment" "t" {
					attachment_type = "IscSi"	# case-insensitive
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t2.id}"
					display_name = "tf-vol-attach-01"
					is_shareable = true
				}
				resource "oci_core_volume_attachment" "t2" {
					attachment_type = "IscSi"	# case-insensitive
					instance_id = "${oci_core_instance.t2.id}"
					volume_id = "${oci_core_volume.t2.id}"
					display_name = "tf-vol-attach-02"
					is_shareable = true
				}	
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_pv_encryption_in_transit_enabled", "false"),
					//resource.TestCheckResourceAttr(s.ResourceName[0], "is_read_only", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "chap_secret"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "chap_username"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "port"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_shareable", "true"),

					resource.TestCheckResourceAttrSet(s.ResourceName[1], "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "is_pv_encryption_in_transit_enabled", "false"),
					//resource.TestCheckResourceAttr(s.ResourceName[1], "is_read_only", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "chap_secret"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "chap_username"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "port"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "is_shareable", "true"),
					func(ts *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(ts, s.ResourceName[1], "id")
						return err
					},
				),
			},
			// verify shared block volume attachment with iscsi attachment type as read-only
			{
				Config: s.Config + `
				resource "oci_core_volume_attachment" "t" {
					attachment_type = "IscSi"	# case-insensitive
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t2.id}"
					display_name = "tf-vol-attach-01"
					is_read_only = true
					is_shareable = true
				}
				resource "oci_core_volume_attachment" "t2" {
					attachment_type = "IscSi"	# case-insensitive
					instance_id = "${oci_core_instance.t2.id}"
					volume_id = "${oci_core_volume.t2.id}"
					display_name = "tf-vol-attach-02"
					is_read_only = true
					is_shareable = true
				}	
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_read_only", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "chap_secret"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "chap_username"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "port"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_shareable", "true"),

					resource.TestCheckResourceAttrSet(s.ResourceName[1], "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "is_read_only", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "chap_secret"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "chap_username"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "port"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "is_shareable", "true"),
					func(ts *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(ts, s.ResourceName[1], "id")
						return err
					},
				),
			},
			// verify shared block volume attachment with paravirtualized attachment type
			{
				Config: s.Config + `
				resource "oci_core_volume_attachment" "t" {
					attachment_type = "paRAviRTualized"	# case-insensitive
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t2.id}"
					display_name = "tf-vol-attach-upd"
					use_chap = true # This should be ignored for paravirtualized attachments
					is_pv_encryption_in_transit_enabled = true
					is_shareable = true
				}
				resource "oci_core_volume_attachment" "t2" {
					attachment_type = "paRAviRTualized"	# case-insensitive
					instance_id = "${oci_core_instance.t2.id}"
					volume_id = "${oci_core_volume.t2.id}"
					display_name = "tf-vol-attach-upd"
					use_chap = true # This should be ignored for paravirtualized attachments
					is_pv_encryption_in_transit_enabled = true
					is_shareable = true
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "display_name", "tf-vol-attach-upd"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_pv_encryption_in_transit_enabled", "true"),
					//resource.TestCheckResourceAttr(s.ResourceName[0], "is_read_only", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "chap_secret"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "chap_username"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "ipv4"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "iqn"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "port"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "attachment_type", "paravirtualized"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_shareable", "true"),

					resource.TestCheckResourceAttrSet(s.ResourceName[1], "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "display_name", "tf-vol-attach-upd"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "is_pv_encryption_in_transit_enabled", "true"),
					//resource.TestCheckResourceAttr(s.ResourceName[1], "is_read_only", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "chap_secret"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "chap_username"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "ipv4"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "iqn"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "port"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "attachment_type", "paravirtualized"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "is_shareable", "true"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName[0], "id")
						if resId2 == resId {
							return fmt.Errorf("resource not recreated when expected to be when updating attachment type")
						}
						resId = resId2
						return err
					},
				),
			},
			// verify shared block volume attachment with paravirtualized attachment type as read-only
			{
				Config: s.Config + `
				resource "oci_core_volume_attachment" "t" {
					attachment_type = "paRAviRTualized"	# case-insensitive
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t2.id}"
					display_name = "tf-vol-attach-upd"
					use_chap = true # This should be ignored for paravirtualized attachments
					is_pv_encryption_in_transit_enabled = true
					is_read_only = true
					is_shareable = true
				}
				resource "oci_core_volume_attachment" "t2" {
					attachment_type = "paRAviRTualized"	# case-insensitive
					instance_id = "${oci_core_instance.t2.id}"
					volume_id = "${oci_core_volume.t2.id}"
					display_name = "tf-vol-attach-upd"
					use_chap = true # This should be ignored for paravirtualized attachments
					is_pv_encryption_in_transit_enabled = true
					is_read_only = true
					is_shareable = true
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "display_name", "tf-vol-attach-upd"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_pv_encryption_in_transit_enabled", "true"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_read_only", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "chap_secret"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "chap_username"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "ipv4"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "iqn"),
					resource.TestCheckNoResourceAttr(s.ResourceName[0], "port"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "attachment_type", "paravirtualized"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName[0], "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName[0], "is_shareable", "true"),

					resource.TestCheckResourceAttrSet(s.ResourceName[1], "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "display_name", "tf-vol-attach-upd"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "is_pv_encryption_in_transit_enabled", "true"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "is_read_only", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "chap_secret"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "chap_username"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "ipv4"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "iqn"),
					resource.TestCheckNoResourceAttr(s.ResourceName[1], "port"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "attachment_type", "paravirtualized"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName[1], "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName[1], "is_shareable", "true"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName[0], "id")
						if resId2 == resId {
							return fmt.Errorf("resource not recreated when expected to be when updating attachment type")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestResourceCoreVolumeAttachmentTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreVolumeAttachmentTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreVolumeAttachmentTestSuite))
}
