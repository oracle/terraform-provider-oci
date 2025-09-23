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
	ResourceAnalyticsMonitoredRegionResourceConfig = ResourceAnalyticsMonitoredRegionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_monitored_region", "test_monitored_region", acctest.Optional, acctest.Update, ResourceAnalyticsMonitoredRegionRepresentation)

	ResourceAnalyticsMonitoredRegionRepresentation = map[string]interface{}{
		"region_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.monitor_region}`},
		"resource_analytics_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id}`},
	}

	ResourceAnalyticsMonitoredRegionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance", "test_resource_analytics_instance", acctest.Required, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceRepresentation)
)

// issue-routing-tag: resource_analytics/default
func TestResourceAnalyticsMonitoredRegionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceAnalyticsMonitoredRegionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	monitorRegion := utils.GetEnvSettingWithBlankDefault("monitor_region")
	monitorRegionStr := fmt.Sprintf("variable \"monitor_region\" { default = \"%s\" }\n", monitorRegion)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	defaultVariablesStr := monitorRegionStr + subnetIdVariableStr + fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_resource_analytics_monitored_region.test_monitored_region"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+defaultVariablesStr+ResourceAnalyticsMonitoredRegionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_monitored_region", "test_monitored_region", acctest.Required, acctest.Create, ResourceAnalyticsMonitoredRegionRepresentation), "resourceanalytics", "monitoredRegion", t)

	acctest.ResourceTest(t, testAccCheckResourceAnalyticsMonitoredRegionDestroy, []resource.TestStep{
		// STEP 0 - verify Create
		{
			ExpectNonEmptyPlan: true,
			Config: config + defaultVariablesStr + ResourceAnalyticsMonitoredRegionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_monitored_region", "test_monitored_region", acctest.Required, acctest.Create, ResourceAnalyticsMonitoredRegionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "region_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_analytics_instance_id"),

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
		// STEP 1 - verify resource import
		{
			Config:                  config + ResourceAnalyticsMonitoredRegionResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckResourceAnalyticsMonitoredRegionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MonitoredRegionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_resource_analytics_monitored_region" {
			noResourceFound = false
			request := oci_resource_analytics.GetMonitoredRegionRequest{}

			tmp := rs.Primary.ID
			request.MonitoredRegionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resource_analytics")

			response, err := client.GetMonitoredRegion(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_resource_analytics.MonitoredRegionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ResourceAnalyticsMonitoredRegion") {
		resource.AddTestSweepers("ResourceAnalyticsMonitoredRegion", &resource.Sweeper{
			Name:         "ResourceAnalyticsMonitoredRegion",
			Dependencies: acctest.DependencyGraph["monitoredRegion"],
			F:            sweepResourceAnalyticsMonitoredRegionResource,
		})
	}
}

func sweepResourceAnalyticsMonitoredRegionResource(compartment string) error {
	monitoredRegionClient := acctest.GetTestClients(&schema.ResourceData{}).MonitoredRegionClient()
	monitoredRegionIds, err := getResourceAnalyticsMonitoredRegionIds(compartment)
	if err != nil {
		return err
	}
	for _, monitoredRegionId := range monitoredRegionIds {
		if ok := acctest.SweeperDefaultResourceId[monitoredRegionId]; !ok {
			deleteMonitoredRegionRequest := oci_resource_analytics.DeleteMonitoredRegionRequest{}

			deleteMonitoredRegionRequest.MonitoredRegionId = &monitoredRegionId

			deleteMonitoredRegionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resource_analytics")
			_, error := monitoredRegionClient.DeleteMonitoredRegion(context.Background(), deleteMonitoredRegionRequest)
			if error != nil {
				fmt.Printf("Error deleting MonitoredRegion %s %s, It is possible that the resource is already deleted. Please verify manually \n", monitoredRegionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &monitoredRegionId, ResourceAnalyticsMonitoredRegionSweepWaitCondition, time.Duration(3*time.Minute),
				ResourceAnalyticsMonitoredRegionSweepResponseFetchOperation, "resource_analytics", true)
		}
	}
	return nil
}

func getResourceAnalyticsMonitoredRegionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MonitoredRegionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	monitoredRegionClient := acctest.GetTestClients(&schema.ResourceData{}).MonitoredRegionClient()

	listMonitoredRegionsRequest := oci_resource_analytics.ListMonitoredRegionsRequest{}
	//listMonitoredRegionsRequest.CompartmentId = &compartmentId
	listMonitoredRegionsRequest.LifecycleState = oci_resource_analytics.MonitoredRegionLifecycleStateActive
	listMonitoredRegionsResponse, err := monitoredRegionClient.ListMonitoredRegions(context.Background(), listMonitoredRegionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MonitoredRegion list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, monitoredRegion := range listMonitoredRegionsResponse.Items {
		id := *monitoredRegion.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MonitoredRegionId", id)
	}
	return resourceIds, nil
}

func ResourceAnalyticsMonitoredRegionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if monitoredRegionResponse, ok := response.Response.(oci_resource_analytics.GetMonitoredRegionResponse); ok {
		return monitoredRegionResponse.LifecycleState != oci_resource_analytics.MonitoredRegionLifecycleStateDeleted
	}
	return false
}

func ResourceAnalyticsMonitoredRegionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MonitoredRegionClient().GetMonitoredRegion(context.Background(), oci_resource_analytics.GetMonitoredRegionRequest{
		MonitoredRegionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
