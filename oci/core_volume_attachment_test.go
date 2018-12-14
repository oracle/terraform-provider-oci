// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

var (
	VolumeAttachmentRequiredOnlyResource = VolumeAttachmentResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_volume_attachment", "test_volume_attachment", Required, Create, volumeAttachmentRepresentation)

	volumeAttachmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"instance_id":         Representation{repType: Optional, create: `${oci_core_instance.test_instance.id}`},
		"volume_id":           Representation{repType: Optional, create: `${oci_core_volume.test_volume.id}`},
		"filter":              RepresentationGroup{Required, volumeAttachmentDataSourceFilterRepresentation}}
	volumeAttachmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_volume_attachment.test_volume_attachment.id}`}},
	}

	volumeAttachmentRepresentation = map[string]interface{}{
		"attachment_type":                     Representation{repType: Required, create: `iscsi`},
		"instance_id":                         Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"volume_id":                           Representation{repType: Required, create: `${oci_core_volume.test_volume.id}`},
		"display_name":                        Representation{repType: Optional, create: `displayName`},
		"is_pv_encryption_in_transit_enabled": Representation{repType: Optional, create: `false`},
		"is_read_only":                        Representation{repType: Optional, create: `false`},
	}

	VolumeAttachmentResourceDependencies = InstanceRequiredOnlyResource + VolumeRequiredOnlyResource
)

func TestCoreVolumeAttachmentResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_attachment.test_volume_attachment"
	datasourceName := "data.oci_core_volume_attachments.test_volume_attachments"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeAttachmentDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VolumeAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_attachment", "test_volume_attachment", Required, Create, volumeAttachmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "iscsi"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeAttachmentResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + VolumeAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_attachment", "test_volume_attachment", Optional, Create, volumeAttachmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "iscsi"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_volume_attachments", "test_volume_attachments", Optional, Update, volumeAttachmentDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_attachment", "test_volume_attachment", Optional, Update, volumeAttachmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_id"),

					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.attachment_type", "iscsi"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.display_name", "displayName"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.is_read_only", "false"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.volume_id"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"use_chap",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckCoreVolumeAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_attachment" {
			noResourceFound = false
			request := oci_core.GetVolumeAttachmentRequest{}

			tmp := rs.Primary.ID
			request.VolumeAttachmentId = &tmp

			response, err := client.GetVolumeAttachment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VolumeAttachmentLifecycleStateDetached): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
