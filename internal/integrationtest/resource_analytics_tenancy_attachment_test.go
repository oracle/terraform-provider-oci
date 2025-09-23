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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_resource_analytics "github.com/oracle/oci-go-sdk/v65/resourceanalytics"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ResourceAnalyticsTenancyAttachmentRequiredOnlyResource = ResourceAnalyticsTenancyAttachmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_tenancy_attachment", "test_tenancy_attachment", acctest.Required, acctest.Create, ResourceAnalyticsTenancyAttachmentRepresentation)

	ResourceAnalyticsTenancyAttachmentResourceConfig = ResourceAnalyticsTenancyAttachmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_tenancy_attachment", "test_tenancy_attachment", acctest.Optional, acctest.Update, ResourceAnalyticsTenancyAttachmentRepresentation)

	ResourceAnalyticsTenancyAttachmentSingularDataSourceRepresentation = map[string]interface{}{
		"tenancy_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resource_analytics_tenancy_attachment.test_tenancy_attachment.id}`},
	}

	ResourceAnalyticsTenancyAttachmentDataSourceRepresentation = map[string]interface{}{
		"id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_resource_analytics_tenancy_attachment.test_tenancy_attachment.id}`},
		"resource_analytics_instance_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id}`},
		"state":                          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: ResourceAnalyticsTenancyAttachmentDataSourceFilterRepresentation}}
	ResourceAnalyticsTenancyAttachmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_resource_analytics_tenancy_attachment.test_tenancy_attachment.id}`}},
	}

	ResourceAnalyticsTenancyAttachmentRepresentation = map[string]interface{}{
		"resource_analytics_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id}`},
		"tenancy_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ta2_ocid}`},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
	}

	ResourceAnalyticsTenancyAttachmentResourceDependencies = ResourceAnalyticsResourceAnalyticsInstanceRequiredOnlyResource
)

// issue-routing-tag: resource_analytics/default
func TestResourceAnalyticsTenancyAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceAnalyticsTenancyAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	tenancy2 := utils.GetEnvSettingWithBlankDefault("tenancy_ta2_ocid")
	defaultVarsStr := subnetIdVariableStr + compartmentIdStr + fmt.Sprintf("variable \"tenancy_ta2_ocid\" { default = \"%s\" }\n", tenancy2)

	resourceName := "oci_resource_analytics_tenancy_attachment.test_tenancy_attachment"
	datasourceName := "data.oci_resource_analytics_tenancy_attachments.test_tenancy_attachments"
	singularDatasourceName := "data.oci_resource_analytics_tenancy_attachment.test_tenancy_attachment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+defaultVarsStr+ResourceAnalyticsTenancyAttachmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_tenancy_attachment", "test_tenancy_attachment", acctest.Optional, acctest.Create, ResourceAnalyticsTenancyAttachmentRepresentation), "resourceanalytics", "tenancyAttachment", t)

	acctest.ResourceTest(t, testAccCheckResourceAnalyticsTenancyAttachmentDestroy, []resource.TestStep{
		// STEP 0 - verify Create
		{
			ExpectNonEmptyPlan: true,
			Config: config + defaultVarsStr + ResourceAnalyticsTenancyAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_tenancy_attachment", "test_tenancy_attachment", acctest.Required, acctest.Create, ResourceAnalyticsTenancyAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "resource_analytics_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// STEP 1 - delete before next Create
		{
			Config: config + defaultVarsStr + ResourceAnalyticsTenancyAttachmentResourceDependencies,
		},
		// STEP 2 - verify Create with optionals
		{
			ExpectNonEmptyPlan: true,
			Config: config + defaultVarsStr + ResourceAnalyticsTenancyAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_tenancy_attachment", "test_tenancy_attachment", acctest.Optional, acctest.Create, ResourceAnalyticsTenancyAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_reporting_tenancy"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_analytics_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// STEP 3 - verify updates to updatable parameters
		{
			ExpectNonEmptyPlan: true,
			Config: config + defaultVarsStr + ResourceAnalyticsTenancyAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_tenancy_attachment", "test_tenancy_attachment", acctest.Optional, acctest.Update, ResourceAnalyticsTenancyAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_reporting_tenancy"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_analytics_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// STEP 4 - verify datasource
		{
			ExpectNonEmptyPlan: true,
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_analytics_tenancy_attachments", "test_tenancy_attachments", acctest.Optional, acctest.Update, ResourceAnalyticsTenancyAttachmentDataSourceRepresentation) +
				defaultVarsStr + ResourceAnalyticsTenancyAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_tenancy_attachment", "test_tenancy_attachment", acctest.Optional, acctest.Update, ResourceAnalyticsTenancyAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "resource_analytics_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "tenancy_attachment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "tenancy_attachment_collection.0.items.#", "1"),
			),
		},
		// STEP 5 - verify singular datasource
		{
			ExpectNonEmptyPlan: true,
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_analytics_tenancy_attachment", "test_tenancy_attachment", acctest.Required, acctest.Create, ResourceAnalyticsTenancyAttachmentSingularDataSourceRepresentation) +
				defaultVarsStr + ResourceAnalyticsTenancyAttachmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_attachment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_reporting_tenancy"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// STEP 6 - verify resource import
		{
			Config:                  config + ResourceAnalyticsTenancyAttachmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckResourceAnalyticsTenancyAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).TenancyAttachmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_resource_analytics_tenancy_attachment" {
			noResourceFound = false
			request := oci_resource_analytics.GetTenancyAttachmentRequest{}

			tmp := rs.Primary.ID
			request.TenancyAttachmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resource_analytics")

			response, err := client.GetTenancyAttachment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_resource_analytics.TenancyAttachmentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ResourceAnalyticsTenancyAttachment") {
		resource.AddTestSweepers("ResourceAnalyticsTenancyAttachment", &resource.Sweeper{
			Name:         "ResourceAnalyticsTenancyAttachment",
			Dependencies: acctest.DependencyGraph["tenancyAttachment"],
			F:            sweepResourceAnalyticsTenancyAttachmentResource,
		})
	}
}

func sweepResourceAnalyticsTenancyAttachmentResource(compartment string) error {
	tenancyAttachmentClient := acctest.GetTestClients(&schema.ResourceData{}).TenancyAttachmentClient()
	tenancyAttachmentIds, err := getResourceAnalyticsTenancyAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, tenancyAttachmentId := range tenancyAttachmentIds {
		if ok := acctest.SweeperDefaultResourceId[tenancyAttachmentId]; !ok {
			deleteTenancyAttachmentRequest := oci_resource_analytics.DeleteTenancyAttachmentRequest{}

			deleteTenancyAttachmentRequest.TenancyAttachmentId = &tenancyAttachmentId

			deleteTenancyAttachmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resource_analytics")
			_, error := tenancyAttachmentClient.DeleteTenancyAttachment(context.Background(), deleteTenancyAttachmentRequest)
			if error != nil {
				fmt.Printf("Error deleting TenancyAttachment %s %s, It is possible that the resource is already deleted. Please verify manually \n", tenancyAttachmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &tenancyAttachmentId, ResourceAnalyticsTenancyAttachmentSweepWaitCondition, time.Duration(3*time.Minute),
				ResourceAnalyticsTenancyAttachmentSweepResponseFetchOperation, "resource_analytics", true)
		}
	}
	return nil
}

func getResourceAnalyticsTenancyAttachmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TenancyAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	tenancyAttachmentClient := acctest.GetTestClients(&schema.ResourceData{}).TenancyAttachmentClient()

	listTenancyAttachmentsRequest := oci_resource_analytics.ListTenancyAttachmentsRequest{}
	//listTenancyAttachmentsRequest.CompartmentId = &compartmentId
	listTenancyAttachmentsRequest.LifecycleState = oci_resource_analytics.TenancyAttachmentLifecycleStateNeedsAttention
	listTenancyAttachmentsResponse, err := tenancyAttachmentClient.ListTenancyAttachments(context.Background(), listTenancyAttachmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting TenancyAttachment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, tenancyAttachment := range listTenancyAttachmentsResponse.Items {
		id := *tenancyAttachment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TenancyAttachmentId", id)
	}
	return resourceIds, nil
}

func ResourceAnalyticsTenancyAttachmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if tenancyAttachmentResponse, ok := response.Response.(oci_resource_analytics.GetTenancyAttachmentResponse); ok {
		return tenancyAttachmentResponse.LifecycleState != oci_resource_analytics.TenancyAttachmentLifecycleStateDeleted
	}
	return false
}

func ResourceAnalyticsTenancyAttachmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.TenancyAttachmentClient().GetTenancyAttachment(context.Background(), oci_resource_analytics.GetTenancyAttachmentRequest{
		TenancyAttachmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
