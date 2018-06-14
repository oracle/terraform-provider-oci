// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeAttachmentTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreVolumeAttachmentTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
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
		display_name = "-tf-instance"
		image = "${var.InstanceImageOCID[var.region]}"
		shape = "VM.Standard1.1"
		subnet_id = "${oci_core_subnet.t.id}"
		metadata {
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
		shape = "VM.Standard1.1"
		subnet_id = "${oci_core_subnet.t.id}"
		metadata {
			ssh_authorized_keys = "${var.ssh_public_key}"
		}
		timeouts {
			create = "15m"
		}
	}

	resource "oci_core_volume" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
	}

	resource "oci_core_volume" "t2" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
	}`

	s.ResourceName = "oci_core_volume_attachment.t"
}

func (s *ResourceCoreVolumeAttachmentTestSuite) TestResourceCoreVolumeAttachment_basic() {
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_volume_attachment" "t" {
					attachment_type = "iSCSI"	# case-insensitive
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t.id}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_read_only", "false"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "chap_secret"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "chap_username"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "port"),
					resource.TestCheckResourceAttr(s.ResourceName, "attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					func(ts *terraform.State) (err error) {
						resId, err = fromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// ensure that changing the case for attachment_type (polymorphic discriminator) is a no-op.
			{
				Config: s.Config + `
				resource "oci_core_volume_attachment" "t" {
					attachment_type = "IscSi"	# case-insensitive
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t.id}"
				}`,
				PlanOnly: true,
			},
			// verify display_name, is_read_only, and use_chap update forces creation of a new resource
			{
				Config: s.Config + `
				resource "oci_core_volume_attachment" "t" {
					attachment_type = "IscSi"	# case-insensitive
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t.id}"
					display_name = "tf-vol-attach-upd"
					is_read_only = true
					use_chap = true
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "tf-vol-attach-upd"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_read_only", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "use_chap", "true"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "chap_secret"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "chap_username"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "port"),
					resource.TestCheckResourceAttr(s.ResourceName, "attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					func(ts *terraform.State) (err error) {
						resId2, err = fromInstanceState(ts, s.ResourceName, "id")
						if resId2 == resId {
							return fmt.Errorf("resource not recreated when expected to be when updating display name")
						}
						resId = resId2
						return err
					},
				),
			},
			// verify instance id update forces new resource creation
			{
				Config: s.Config + `
				resource "oci_core_volume_attachment" "t" {
					attachment_type = "IscSi"	# case-insensitive
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t2.id}"
					volume_id = "${oci_core_volume.t.id}"
					display_name = "tf-vol-attach-upd"
					is_read_only = true
					use_chap = true
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "tf-vol-attach-upd"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_read_only", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "use_chap", "true"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "chap_secret"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "chap_username"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "port"),
					resource.TestCheckResourceAttr(s.ResourceName, "attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					func(ts *terraform.State) (err error) {
						resId2, err = fromInstanceState(ts, s.ResourceName, "id")
						if resId2 == resId {
							return fmt.Errorf("resource not recreated when expected to be when updating instance id")
						}
						resId = resId2
						return err
					},
				),
			},
			// verify volume id update forces new resource creation
			{
				Config: s.Config + `
				resource "oci_core_volume_attachment" "t" {
					attachment_type = "IscSi"	# case-insensitive
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t2.id}"
					volume_id = "${oci_core_volume.t2.id}"
					display_name = "tf-vol-attach-upd"
					is_read_only = true
					use_chap = true
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "tf-vol-attach-upd"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_read_only", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "use_chap", "true"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "chap_secret"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "chap_username"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "port"),
					resource.TestCheckResourceAttr(s.ResourceName, "attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					func(ts *terraform.State) (err error) {
						resId2, err = fromInstanceState(ts, s.ResourceName, "id")
						if resId2 == resId {
							return fmt.Errorf("resource not recreated when expected to be when updating volume id")
						}
						resId = resId2
						return err
					},
				),
			},
			// verify attachment type update to paravirtualized forces creation of a new resource
			{
				Config: s.Config + `
				resource "oci_core_volume_attachment" "t" {
					attachment_type = "paRAviRTualized"	# case-insensitive
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t2.id}"
					volume_id = "${oci_core_volume.t2.id}"
					display_name = "tf-vol-attach-upd"
					is_read_only = true
					use_chap = true # This should be ignored for paravirtualized attachments
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "tf-vol-attach-upd"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_read_only", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "chap_secret"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "chap_username"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "ipv4"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "iqn"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "port"),
					resource.TestCheckResourceAttr(s.ResourceName, "attachment_type", "paravirtualized"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					func(ts *terraform.State) (err error) {
						resId2, err = fromInstanceState(ts, s.ResourceName, "id")
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

func TestResourceCoreVolumeAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeAttachmentTestSuite))
}
