// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v49/common"
	oci_core "github.com/oracle/oci-go-sdk/v49/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VolumeAttachmentRequiredOnlyResource = VolumeAttachmentResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_volume_attachment", "test_volume_attachment", Required, Create, volumeAttachmentRepresentation)

	volumeAttachmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"availability_domain": Representation{RepType: Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"instance_id":         Representation{RepType: Optional, Create: `${oci_core_instance.test_instance.id}`},
		"volume_id":           Representation{RepType: Optional, Create: `${oci_core_volume.test_volume.id}`},
		"filter":              RepresentationGroup{Required, volumeAttachmentDataSourceFilterRepresentation}}
	volumeAttachmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_volume_attachment.test_volume_attachment.id}`}},
	}

	volumeAttachmentRepresentation = map[string]interface{}{
		"attachment_type":                     Representation{RepType: Required, Create: `iscsi`},
		"instance_id":                         Representation{RepType: Required, Create: `${oci_core_instance.test_instance.id}`},
		"volume_id":                           Representation{RepType: Required, Create: `${oci_core_volume.test_volume.id}`},
		"device":                              Representation{RepType: Optional, Create: `/dev/oracleoci/oraclevdb`},
		"display_name":                        Representation{RepType: Optional, Create: `displayName`},
		"encryption_in_transit_type":          Representation{RepType: Optional, Create: `NONE`},
		"is_pv_encryption_in_transit_enabled": Representation{RepType: Optional, Create: `false`},
		"is_read_only":                        Representation{RepType: Optional, Create: `false`},
		"is_shareable":                        Representation{RepType: Optional, Create: `false`},
	}

	VolumeAttachmentResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, SubnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		OciImageIdsVariable +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", Required, Create, volumeRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreVolumeAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_attachment.test_volume_attachment"
	datasourceName := "data.oci_core_volume_attachments.test_volume_attachments"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+VolumeAttachmentResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_volume_attachment", "test_volume_attachment", Optional, Create, volumeAttachmentRepresentation), "core", "volumeAttachment", t)

	ResourceTest(t, testAccCheckCoreVolumeAttachmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VolumeAttachmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_volume_attachment", "test_volume_attachment", Required, Create, volumeAttachmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attachment_type", "iscsi"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VolumeAttachmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VolumeAttachmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_volume_attachment", "test_volume_attachment", Optional, Create, volumeAttachmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attachment_type", "iscsi"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "device", "/dev/oracleoci/oraclevdb"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "encryption_in_transit_type", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_shareable", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_volume_attachments", "test_volume_attachments", Optional, Update, volumeAttachmentDataSourceRepresentation) +
				compartmentIdVariableStr + VolumeAttachmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_volume_attachment", "test_volume_attachment", Optional, Update, volumeAttachmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_id"),

				resource.TestCheckResourceAttr(datasourceName, "volume_attachments.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.attachment_type", "iscsi"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.device", "/dev/oracleoci/oraclevdb"),
				resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.encryption_in_transit_type", "NONE"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.ipv4"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.iqn"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.is_multipath"),
				resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.is_read_only", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.port"),
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
	})
}

func testAccCheckCoreVolumeAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := TestAccProvider.Meta().(*OracleClients).computeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_attachment" {
			noResourceFound = false
			request := oci_core.GetVolumeAttachmentRequest{}

			tmp := rs.Primary.ID
			request.VolumeAttachmentId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")

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

func init() {
	if DependencyGraph == nil {
		InitDependencyGraph()
	}
	if !InSweeperExcludeList("CoreVolumeAttachment") {
		resource.AddTestSweepers("CoreVolumeAttachment", &resource.Sweeper{
			Name:         "CoreVolumeAttachment",
			Dependencies: DependencyGraph["volumeAttachment"],
			F:            sweepCoreVolumeAttachmentResource,
		})
	}
}

func sweepCoreVolumeAttachmentResource(compartment string) error {
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()
	volumeAttachmentIds, err := getVolumeAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, volumeAttachmentId := range volumeAttachmentIds {
		if ok := SweeperDefaultResourceId[volumeAttachmentId]; !ok {
			detachVolumeRequest := oci_core.DetachVolumeRequest{}

			detachVolumeRequest.VolumeAttachmentId = &volumeAttachmentId

			detachVolumeRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")
			_, error := computeClient.DetachVolume(context.Background(), detachVolumeRequest)
			if error != nil {
				fmt.Printf("Error deleting VolumeAttachment %s %s, It is possible that the resource is already deleted. Please verify manually \n", volumeAttachmentId, error)
				continue
			}
			WaitTillCondition(TestAccProvider, &volumeAttachmentId, volumeAttachmentSweepWaitCondition, time.Duration(3*time.Minute),
				volumeAttachmentSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVolumeAttachmentIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "VolumeAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()

	listVolumeAttachmentsRequest := oci_core.ListVolumeAttachmentsRequest{}
	listVolumeAttachmentsRequest.CompartmentId = &compartmentId
	listVolumeAttachmentsResponse, err := computeClient.ListVolumeAttachments(context.Background(), listVolumeAttachmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VolumeAttachment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, volumeAttachment := range listVolumeAttachmentsResponse.Items {
		id := *volumeAttachment.GetId()
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "VolumeAttachmentId", id)
	}
	return resourceIds, nil
}

func volumeAttachmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if volumeAttachmentResponse, ok := response.Response.(oci_core.GetVolumeAttachmentResponse); ok {
		return volumeAttachmentResponse.GetLifecycleState() != oci_core.VolumeAttachmentLifecycleStateDetached
	}
	return false
}

func volumeAttachmentSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.computeClient().GetVolumeAttachment(context.Background(), oci_core.GetVolumeAttachmentRequest{
		VolumeAttachmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
