// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/oracle/oci-go-sdk/v56/common"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/oracle/oci-go-sdk/v56/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVolumeAttachmentTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreVolumeAttachmentTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + instanceConfig + `
	resource "oci_core_volume" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-volume"
	}
	resource "oci_core_volume_attachment" "t" {
		instance_id = "${oci_core_instance.t.id}"
		volume_id = "${oci_core_volume.t.id}"
		attachment_type = "iscsi"
		use_chap = true
	}`
	s.ResourceName = "data.oci_core_volume_attachments.t"
}

func (s *DatasourceCoreVolumeAttachmentTestSuite) TestAccDatasourceCoreVolumeAttachment_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + `
				data "oci_core_volume_attachments" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t.id}"
					filter {
						name = "id"
						values = ["${oci_core_volume_attachment.t.id}"]
					}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.0.attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.0.state", string(core.VolumeAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.0.is_read_only", "false"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.port"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.chap_secret"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.chap_username"),
				),
			},
		},
	},
	)
}

type customVolumeAttachment struct {
	ad            string
	compartmentId string
	id            string
	instanceId    string
	isReadOnly    bool
	volumeId      string
	displayName   string
	timeCreated   common.SDKTime
	state         oci_core.VolumeAttachmentLifecycleStateEnum
}

//GetAvailabilityDomain returns AvailabilityDomain
func (m customVolumeAttachment) GetAvailabilityDomain() *string {
	return &m.ad
}

//GetCompartmentId returns CompartmentId
func (m customVolumeAttachment) GetCompartmentId() *string {
	return &m.compartmentId
}

//GetId returns Id
func (m customVolumeAttachment) GetId() *string {
	return &m.id
}

//GetInstanceId returns InstanceId
func (m customVolumeAttachment) GetInstanceId() *string {
	return &m.instanceId
}

//GetIsReadOnly returns IsReadOnly
func (m customVolumeAttachment) GetIsReadOnly() *bool {
	return &m.isReadOnly
}

//GetLifecycleState returns LifecycleState
func (m customVolumeAttachment) GetLifecycleState() oci_core.VolumeAttachmentLifecycleStateEnum {
	return m.state
}

//GetTimeCreated returns TimeCreated
func (m customVolumeAttachment) GetTimeCreated() *common.SDKTime {
	return &m.timeCreated
}

//GetVolumeId returns VolumeId
func (m customVolumeAttachment) GetVolumeId() *string {
	return &m.volumeId
}

//GetDisplayName returns DisplayName
func (m customVolumeAttachment) GetDisplayName() *string {
	return &m.displayName
}

func checkExpectedValue(mapped map[string]interface{}, key string, expected string, t *testing.T) {
	if value := mapped[key].(string); value != expected {
		t.Errorf("Expected attachment to have type %s, but got %s", expected, value)
	}
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestDatasourceCoreVolumeAttachmentTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreVolumeAttachmentTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreVolumeAttachmentTestSuite))
}
