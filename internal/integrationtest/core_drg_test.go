// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	DrgRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, drgRepresentation)

	drgDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: drgDataSourceFilterRepresentation}}
	drgDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_drg.test_drg.id}`}},
	}

	drgRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MyDrg`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}

	DrgResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreDrgResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_drg.test_drg"
	datasourceName := "data.oci_core_drgs.test_drgs"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DrgResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Optional, acctest.Create, drgRepresentation), "core", "drg", t)

	acctest.ResourceTest(t, testAccCheckCoreDrgDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DrgResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, drgRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DrgResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DrgResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Optional, acctest.Create, drgRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyDrg"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "redundancy_status"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DrgResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(drgRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyDrg"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "redundancy_status"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DrgResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Optional, acctest.Update, drgRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "redundancy_status"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_drgs", "test_drgs", acctest.Optional, acctest.Update, drgDataSourceRepresentation) +
				compartmentIdVariableStr + DrgResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Optional, acctest.Update, drgRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "redundancy_status"),
				resource.TestCheckResourceAttr(datasourceName, "drgs.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "drgs.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "drgs.0.default_drg_route_tables.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "drgs.0.default_export_drg_route_distribution_id"),
				resource.TestCheckResourceAttr(datasourceName, "drgs.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "drgs.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "drgs.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drgs.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "drgs.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreDrgDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_drg" {
			noResourceFound = false
			request := oci_core.GetDrgRequest{}

			tmp := rs.Primary.ID
			request.DrgId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetDrg(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.DrgLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreDrg") {
		resource.AddTestSweepers("CoreDrg", &resource.Sweeper{
			Name:         "CoreDrg",
			Dependencies: acctest.DependencyGraph["drg"],
			F:            sweepCoreDrgResource,
		})
	}
}

func sweepCoreDrgResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	drgIds, err := getDrgIds(compartment)
	if err != nil {
		return err
	}
	for _, drgId := range drgIds {
		if ok := acctest.SweeperDefaultResourceId[drgId]; !ok {
			deleteDrgRequest := oci_core.DeleteDrgRequest{}

			deleteDrgRequest.DrgId = &drgId

			deleteDrgRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteDrg(context.Background(), deleteDrgRequest)
			if error != nil {
				fmt.Printf("Error deleting Drg %s %s, It is possible that the resource is already deleted. Please verify manually \n", drgId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &drgId, drgSweepWaitCondition, time.Duration(3*time.Minute),
				drgSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getDrgIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DrgId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listDrgsRequest := oci_core.ListDrgsRequest{}
	listDrgsRequest.CompartmentId = &compartmentId
	listDrgsResponse, err := virtualNetworkClient.ListDrgs(context.Background(), listDrgsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Drg list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, drg := range listDrgsResponse.Items {
		id := *drg.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DrgId", id)
		acctest.SweeperDefaultResourceId[*drg.DefaultExportDrgRouteDistributionId] = true

	}
	return resourceIds, nil
}

func drgSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if drgResponse, ok := response.Response.(oci_core.GetDrgResponse); ok {
		return drgResponse.LifecycleState != oci_core.DrgLifecycleStateTerminated
	}
	return false
}

func drgSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetDrg(context.Background(), oci_core.GetDrgRequest{
		DrgId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
