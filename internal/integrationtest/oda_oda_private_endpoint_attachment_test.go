// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OdaOdaPrivateEndpointAttachmentRequiredOnlyResource = OdaOdaPrivateEndpointAttachmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint_attachment", "test_oda_private_endpoint_attachment", acctest.Required, acctest.Create, OdaOdaPrivateEndpointAttachmentRepresentation)

	OdaOdaPrivateEndpointAttachmentResourceConfig = OdaOdaPrivateEndpointAttachmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint_attachment", "test_oda_private_endpoint_attachment", acctest.Optional, acctest.Update, OdaOdaPrivateEndpointAttachmentRepresentation)

	OdaOdaOdaPrivateEndpointAttachmentSingularDataSourceRepresentation = map[string]interface{}{
		"oda_private_endpoint_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_oda_oda_private_endpoint_attachment.test_oda_private_endpoint_attachment.id}`},
	}

	OdaOdaOdaPrivateEndpointAttachmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"oda_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_oda_oda_private_endpoint.test_oda_private_endpoint.id}`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `CREATING`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: OdaOdaPrivateEndpointAttachmentDataSourceFilterRepresentation}}

	OdaOdaPrivateEndpointAttachmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_oda_oda_private_endpoint_attachment.test_oda_private_endpoint_attachment.id}`}},
	}

	OdaOdaPrivateEndpointAttachmentRepresentation = map[string]interface{}{
		"oda_instance_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_oda_oda_instance.test_oda_instance.id}`},
		"oda_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_oda_oda_private_endpoint.test_oda_private_endpoint.id}`},
	}

	OdaOdaPrivateEndpointAttachmentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", acctest.Required, acctest.Create, OdaOdaInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint", "test_oda_private_endpoint", acctest.Required, acctest.Create, OdaOdaPrivateEndpointRepresentation) +
		OdaOdaPrivateEndpointResourceDependencies
)

// issue-routing-tag: oda/default
func TestOdaOdaPrivateEndpointAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOdaOdaPrivateEndpointAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_oda_oda_private_endpoint_attachment.test_oda_private_endpoint_attachment"
	datasourceName := "data.oci_oda_oda_private_endpoint_attachments.test_oda_private_endpoint_attachments"
	singularDatasourceName := "data.oci_oda_oda_private_endpoint_attachment.test_oda_private_endpoint_attachment"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OdaOdaPrivateEndpointAttachmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint_attachment", "test_oda_private_endpoint_attachment", acctest.Required, acctest.Create, OdaOdaPrivateEndpointAttachmentRepresentation), "oda", "odaPrivateEndpointAttachment", t)

	acctest.ResourceTest(t, testAccCheckOdaOdaPrivateEndpointAttachmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OdaOdaPrivateEndpointAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint_attachment", "test_oda_private_endpoint_attachment", acctest.Required, acctest.Create, OdaOdaPrivateEndpointAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "oda_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oda_private_endpoint_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_oda_oda_private_endpoint_attachments", "test_oda_private_endpoint_attachments", acctest.Optional, acctest.Update, OdaOdaOdaPrivateEndpointAttachmentDataSourceRepresentation) +
				compartmentIdVariableStr + OdaOdaPrivateEndpointAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint_attachment", "test_oda_private_endpoint_attachment", acctest.Optional, acctest.Update, OdaOdaPrivateEndpointAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "oda_private_endpoint_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "CREATING"),
				resource.TestCheckResourceAttr(datasourceName, "oda_private_endpoint_attachment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oda_private_endpoint_attachment_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_oda_oda_private_endpoint_attachment", "test_oda_private_endpoint_attachment", acctest.Required, acctest.Create, OdaOdaOdaPrivateEndpointAttachmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OdaOdaPrivateEndpointAttachmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oda_private_endpoint_attachment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OdaOdaPrivateEndpointAttachmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOdaOdaPrivateEndpointAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_oda_oda_private_endpoint_attachment" {
			noResourceFound = false
			request := oci_oda.GetOdaPrivateEndpointAttachmentRequest{}

			tmp := rs.Primary.ID
			request.OdaPrivateEndpointAttachmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oda")

			response, err := client.GetOdaPrivateEndpointAttachment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_oda.OdaPrivateEndpointAttachmentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OdaOdaPrivateEndpointAttachment") {
		resource.AddTestSweepers("OdaOdaPrivateEndpointAttachment", &resource.Sweeper{
			Name:         "OdaOdaPrivateEndpointAttachment",
			Dependencies: acctest.DependencyGraph["odaPrivateEndpointAttachment"],
			F:            sweepOdaOdaPrivateEndpointAttachmentResource,
		})
	}
}

func sweepOdaOdaPrivateEndpointAttachmentResource(compartment string) error {
	managementClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementClient()
	odaPrivateEndpointAttachmentIds, err := getOdaOdaPrivateEndpointAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, odaPrivateEndpointAttachmentId := range odaPrivateEndpointAttachmentIds {
		if ok := acctest.SweeperDefaultResourceId[odaPrivateEndpointAttachmentId]; !ok {
			deleteOdaPrivateEndpointAttachmentRequest := oci_oda.DeleteOdaPrivateEndpointAttachmentRequest{}

			deleteOdaPrivateEndpointAttachmentRequest.OdaPrivateEndpointAttachmentId = &odaPrivateEndpointAttachmentId

			deleteOdaPrivateEndpointAttachmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oda")
			_, error := managementClient.DeleteOdaPrivateEndpointAttachment(context.Background(), deleteOdaPrivateEndpointAttachmentRequest)
			if error != nil {
				fmt.Printf("Error deleting OdaPrivateEndpointAttachment %s %s, It is possible that the resource is already deleted. Please verify manually \n", odaPrivateEndpointAttachmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &odaPrivateEndpointAttachmentId, OdaOdaPrivateEndpointAttachmentSweepWaitCondition, time.Duration(3*time.Minute),
				OdaOdaPrivateEndpointAttachmentSweepResponseFetchOperation, "oda", true)
		}
	}
	return nil
}

func getOdaOdaPrivateEndpointAttachmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OdaPrivateEndpointAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managementClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementClient()

	listOdaPrivateEndpointAttachmentsRequest := oci_oda.ListOdaPrivateEndpointAttachmentsRequest{}
	listOdaPrivateEndpointAttachmentsRequest.CompartmentId = &compartmentId

	odaPrivateEndpointIds, error := getOdaOdaPrivateEndpointIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting odaPrivateEndpointId required for OdaPrivateEndpointAttachment resource requests \n")
	}
	for _, odaPrivateEndpointId := range odaPrivateEndpointIds {
		listOdaPrivateEndpointAttachmentsRequest.OdaPrivateEndpointId = &odaPrivateEndpointId

		listOdaPrivateEndpointAttachmentsRequest.LifecycleState = oci_oda.OdaPrivateEndpointAttachmentLifecycleStateActive
		listOdaPrivateEndpointAttachmentsResponse, err := managementClient.ListOdaPrivateEndpointAttachments(context.Background(), listOdaPrivateEndpointAttachmentsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting OdaPrivateEndpointAttachment list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, odaPrivateEndpointAttachment := range listOdaPrivateEndpointAttachmentsResponse.Items {
			id := *odaPrivateEndpointAttachment.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OdaPrivateEndpointAttachmentId", id)
		}

	}
	return resourceIds, nil
}

func OdaOdaPrivateEndpointAttachmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if odaPrivateEndpointAttachmentResponse, ok := response.Response.(oci_oda.GetOdaPrivateEndpointAttachmentResponse); ok {
		return odaPrivateEndpointAttachmentResponse.LifecycleState != oci_oda.OdaPrivateEndpointAttachmentLifecycleStateDeleted
	}
	return false
}

func OdaOdaPrivateEndpointAttachmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ManagementClient().GetOdaPrivateEndpointAttachment(context.Background(), oci_oda.GetOdaPrivateEndpointAttachmentRequest{
		OdaPrivateEndpointAttachmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
