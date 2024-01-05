// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FusionAppsFusionEnvironmentServiceAttachmentRequiredOnlyResource = acctest.
										GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_service_attachment", "test_fusion_environment_service_attachment", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentServiceAttachmentRepresentation)

	FusionAppsFusionAppsFusionEnvironmentServiceAttachmentSingularDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.fusion_pod_id}`},
		"service_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment_service_attachment.test_fusion_environment_service_attachment.id}`},
	}

	FusionAppsFusionAppsFusionEnvironmentServiceAttachmentDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.fusion_pod_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `devcs-ppdibaejb-test`},
		"service_instance_type": acctest.Representation{RepType: acctest.Optional, Create: `VISUAL_BUILDER_STUDIO`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	FusionAppsFusionEnvironmentServiceAttachmentRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.fusion_pod_id}`},
		"service_instance_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.service_instance_id}`},
		"service_instance_type": acctest.Representation{RepType: acctest.Required, Create: `VISUAL_BUILDER_STUDIO`},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: definedTagsIgnoreRepresentation_fusionapps},
	}

	FusionAppsFusionEnvironmentServiceAttachmentResourceConfig = acctest.
									GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_service_attachment", "test_fusion_environment_service_attachment", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentServiceAttachmentRepresentation)

	FusionAppsFusionEnvironmentServiceAttachmentResourceDependencies = `
		data "oci_fusion_apps_fusion_environment" "test_fusion_environment" {
		  fusion_environment_id = "${var.fusion_pod_id}"
	}`
)

// issue-routing-tag: fusion_apps/default
func TestFusionAppsFusionEnvironmentServiceAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFusionAppsFusionEnvironmentServiceAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	fusionPodId := utils.GetEnvSettingWithBlankDefault("fusion_pod_id")
	fusionPodIdStr := fmt.Sprintf("variable \"fusion_pod_id\" { default = \"%s\" }\n", fusionPodId)

	serviceInstanceId := utils.GetEnvSettingWithBlankDefault("service_instance_id")
	serviceInstanceIdStr := fmt.Sprintf("variable \"service_instance_id\" { default = \"%s\" }\n", serviceInstanceId)
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fusion_apps_fusion_environment_service_attachment.test_fusion_environment_service_attachment"
	datasourceName := "data.oci_fusion_apps_fusion_environment_service_attachments.test_fusion_environment_service_attachments"
	singularDatasourceName := "data.oci_fusion_apps_fusion_environment_service_attachment.test_fusion_environment_service_attachment"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+fusionPodIdStr+serviceInstanceIdStr+
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_service_attachment", "test_fusion_environment_service_attachment", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentServiceAttachmentRepresentation), "fusionapps", "fusionEnvironmentServiceAttachment", t)

	acctest.ResourceTest(t, testAccCheckFusionAppsFusionEnvironmentServiceAttachmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + fusionPodIdStr + serviceInstanceIdStr +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_service_attachment", "test_fusion_environment_service_attachment", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentServiceAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "service_instance_type", "VISUAL_BUILDER_STUDIO"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault(
						"enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config + compartmentIdVariableStr + fusionPodIdStr + serviceInstanceIdStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_service_attachments", "test_fusion_environment_service_attachments", acctest.Optional, acctest.Update, FusionAppsFusionAppsFusionEnvironmentServiceAttachmentDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_service_attachment", "test_fusion_environment_service_attachment", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentServiceAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "devcs-ppdibaejb-test"),
				resource.TestCheckResourceAttrSet(datasourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttr(datasourceName, "service_instance_type", "VISUAL_BUILDER_STUDIO"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "service_attachment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "service_attachment_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + fusionPodIdStr + serviceInstanceIdStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_service_attachment", "test_fusion_environment_service_attachment", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentServiceAttachmentSingularDataSourceRepresentation) +
				FusionAppsFusionEnvironmentServiceAttachmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_attachment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_sku_based"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_instance_type", "VISUAL_BUILDER_STUDIO"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + FusionAppsFusionEnvironmentServiceAttachmentRequiredOnlyResource + fusionPodIdStr,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFusionAppsFusionEnvironmentServiceAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FusionApplicationsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fusion_apps_fusion_environment_service_attachment" {
			noResourceFound = false
			request := oci_fusion_apps.GetServiceAttachmentRequest{}
			if value, ok := rs.Primary.Attributes["fusion_environment_id"]; ok {
				request.FusionEnvironmentId = &value
			}

			request.ServiceAttachmentId = &rs.Primary.ID
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fusion_apps")

			response, err := client.GetServiceAttachment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fusion_apps.ServiceAttachmentLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("FusionAppsFusionEnvironmentServiceAttachment") {
		resource.AddTestSweepers("FusionAppsFusionEnvironmentServiceAttachment", &resource.Sweeper{
			Name:         "FusionAppsFusionEnvironmentServiceAttachment",
			Dependencies: acctest.DependencyGraph["fusionEnvironmentServiceAttachment"],
			F:            sweepFusionAppsFusionEnvironmentServiceAttachmentResource,
		})
	}
}

func sweepFusionAppsFusionEnvironmentServiceAttachmentResource(compartment string) error {
	fusionApplicationsClient := acctest.GetTestClients(&schema.ResourceData{}).FusionApplicationsClient()
	fusionEnvironmentServiceAttachmentIds, err := getFusionAppsFusionEnvironmentServiceAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, fusionEnvironmentServiceAttachmentId := range fusionEnvironmentServiceAttachmentIds {
		if ok := acctest.SweeperDefaultResourceId[fusionEnvironmentServiceAttachmentId]; !ok {
			deleteServiceAttachmentRequest := oci_fusion_apps.DeleteServiceAttachmentRequest{}

			deleteServiceAttachmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fusion_apps")
			_, error := fusionApplicationsClient.DeleteServiceAttachment(context.Background(), deleteServiceAttachmentRequest)
			if error != nil {
				fmt.Printf("Error deleting FusionEnvironmentServiceAttachment %s %s, It is possible that the resource is already deleted. Please verify manually \n", fusionEnvironmentServiceAttachmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fusionEnvironmentServiceAttachmentId, FusionAppsFusionEnvironmentServiceAttachmentSweepWaitCondition, time.Duration(3*time.Minute),
				FusionAppsFusionEnvironmentServiceAttachmentSweepResponseFetchOperation, "fusion_apps", true)
		}
	}
	return nil
}

func getFusionAppsFusionEnvironmentServiceAttachmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FusionEnvironmentServiceAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fusionApplicationsClient := acctest.GetTestClients(&schema.ResourceData{}).FusionApplicationsClient()

	listServiceAttachmentsRequest := oci_fusion_apps.ListServiceAttachmentsRequest{}

	fusionEnvironmentIds, error := getFusionAppsFusionEnvironmentIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting fusionEnvironmentId required for FusionEnvironmentServiceAttachment resource requests \n")
	}
	for _, fusionEnvironmentId := range fusionEnvironmentIds {
		listServiceAttachmentsRequest.FusionEnvironmentId = &fusionEnvironmentId

		listServiceAttachmentsRequest.LifecycleState = oci_fusion_apps.ServiceAttachmentLifecycleStateActive
		listServiceAttachmentsResponse, err := fusionApplicationsClient.ListServiceAttachments(context.Background(), listServiceAttachmentsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting FusionEnvironmentServiceAttachment list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, fusionEnvironmentServiceAttachment := range listServiceAttachmentsResponse.Items {
			id := *fusionEnvironmentServiceAttachment.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FusionEnvironmentServiceAttachmentId", id)
		}

	}
	return resourceIds, nil
}

func FusionAppsFusionEnvironmentServiceAttachmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fusionEnvironmentServiceAttachmentResponse, ok := response.Response.(oci_fusion_apps.GetServiceAttachmentResponse); ok {
		return fusionEnvironmentServiceAttachmentResponse.LifecycleState != oci_fusion_apps.ServiceAttachmentLifecycleStateDeleted
	}
	return false
}

func FusionAppsFusionEnvironmentServiceAttachmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FusionApplicationsClient().GetServiceAttachment(context.Background(), oci_fusion_apps.GetServiceAttachmentRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
