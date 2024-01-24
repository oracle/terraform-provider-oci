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
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreDrgRouteDistributionRequiredOnlyResource = CoreDrgRouteDistributionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", acctest.Required, acctest.Create, CoreDrgRouteDistributionRepresentation)

	CoreDrgRouteDistributionResourceConfig = CoreDrgRouteDistributionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", acctest.Optional, acctest.Update, CoreDrgRouteDistributionRepresentation)

	CoreCoreDrgRouteDistributionSingularDataSourceRepresentation = map[string]interface{}{
		"drg_route_distribution_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
	}

	CoreCoreDrgRouteDistributionDataSourceRepresentation = map[string]interface{}{
		"drg_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":        acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":       acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgRouteDistributionDataSourceFilterRepresentation}}
	CoreDrgRouteDistributionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_drg_route_distribution.test_drg_route_distribution.id}`}},
	}

	CoreDrgRouteDistributionRepresentation = map[string]interface{}{
		"distribution_type": acctest.Representation{RepType: acctest.Required, Create: `IMPORT`},
		"drg_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}

	CoreDrgRouteDistributionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, CoreDrgRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/pnp
func TestCoreDrgRouteDistributionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgRouteDistributionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_route_distribution.test_drg_route_distribution"
	datasourceName := "data.oci_core_drg_route_distributions.test_drg_route_distributions"
	singularDatasourceName := "data.oci_core_drg_route_distribution.test_drg_route_distribution"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreDrgRouteDistributionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", acctest.Optional, acctest.Create, CoreDrgRouteDistributionRepresentation), "core", "drgRouteDistribution", t)

	acctest.ResourceTest(t, testAccCheckCoreDrgRouteDistributionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteDistributionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", acctest.Required, acctest.Create, CoreDrgRouteDistributionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "distribution_type", "IMPORT"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteDistributionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteDistributionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", acctest.Optional, acctest.Create, CoreDrgRouteDistributionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "distribution_type", "IMPORT"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteDistributionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", acctest.Optional, acctest.Update, CoreDrgRouteDistributionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "distribution_type", "IMPORT"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_drg_route_distributions", "test_drg_route_distributions", acctest.Optional, acctest.Update, CoreCoreDrgRouteDistributionDataSourceRepresentation) +
				compartmentIdVariableStr + CoreDrgRouteDistributionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", acctest.Optional, acctest.Update, CoreDrgRouteDistributionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "drg_route_distributions.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_distributions.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distributions.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distributions.0.distribution_type", "IMPORT"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_distributions.0.drg_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_distributions.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_distributions.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_distributions.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_distributions.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", acctest.Required, acctest.Create, CoreCoreDrgRouteDistributionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreDrgRouteDistributionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "drg_route_distribution_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "distribution_type", "IMPORT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreDrgRouteDistributionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreDrgRouteDistributionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_drg_route_distribution" {
			noResourceFound = false
			request := oci_core.GetDrgRouteDistributionRequest{}

			tmp := rs.Primary.ID
			request.DrgRouteDistributionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetDrgRouteDistribution(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.DrgRouteDistributionLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreDrgRouteDistribution") {
		resource.AddTestSweepers("CoreDrgRouteDistribution", &resource.Sweeper{
			Name:         "CoreDrgRouteDistribution",
			Dependencies: acctest.DependencyGraph["drgRouteDistribution"],
			F:            sweepCoreDrgRouteDistributionResource,
		})
	}
}

func sweepCoreDrgRouteDistributionResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	drgRouteDistributionIds, err := getCoreDrgRouteDistributionIds(compartment)
	if err != nil {
		return err
	}
	for _, drgRouteDistributionId := range drgRouteDistributionIds {
		if ok := acctest.SweeperDefaultResourceId[drgRouteDistributionId]; !ok {
			deleteDrgRouteDistributionRequest := oci_core.DeleteDrgRouteDistributionRequest{}

			deleteDrgRouteDistributionRequest.DrgRouteDistributionId = &drgRouteDistributionId

			deleteDrgRouteDistributionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteDrgRouteDistribution(context.Background(), deleteDrgRouteDistributionRequest)
			if error != nil {
				fmt.Printf("Error deleting DrgRouteDistribution %s %s, It is possible that the resource is already deleted. Please verify manually \n", drgRouteDistributionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &drgRouteDistributionId, CoreDrgRouteDistributionSweepWaitCondition, time.Duration(3*time.Minute),
				CoreDrgRouteDistributionSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreDrgRouteDistributionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DrgRouteDistributionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listDrgRouteDistributionsRequest := oci_core.ListDrgRouteDistributionsRequest{}
	// listDrgRouteDistributionsRequest.CompartmentId = &compartmentId

	drgIds, error := getCoreDrgIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting drgId required for DrgRouteDistribution resource requests \n")
	}
	for _, drgId := range drgIds {
		listDrgRouteDistributionsRequest.DrgId = &drgId

		listDrgRouteDistributionsRequest.LifecycleState = oci_core.DrgRouteDistributionLifecycleStateAvailable
		listDrgRouteDistributionsResponse, err := virtualNetworkClient.ListDrgRouteDistributions(context.Background(), listDrgRouteDistributionsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DrgRouteDistribution list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, drgRouteDistribution := range listDrgRouteDistributionsResponse.Items {
			id := *drgRouteDistribution.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DrgRouteDistributionId", id)
		}

	}
	return resourceIds, nil
}

func CoreDrgRouteDistributionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if drgRouteDistributionResponse, ok := response.Response.(oci_core.GetDrgRouteDistributionResponse); ok {
		return drgRouteDistributionResponse.LifecycleState != oci_core.DrgRouteDistributionLifecycleStateTerminated
	}
	return false
}

func CoreDrgRouteDistributionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetDrgRouteDistribution(context.Background(), oci_core.GetDrgRouteDistributionRequest{
		DrgRouteDistributionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
