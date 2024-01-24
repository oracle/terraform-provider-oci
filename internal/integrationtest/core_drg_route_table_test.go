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
	CoreDrgRouteTableRequiredOnlyResource = CoreDrgRouteTableResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Required, acctest.Create, CoreDrgRouteTableRepresentation)

	CoreDrgRouteTableResourceConfig = CoreDrgRouteTableResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Optional, acctest.Update, CoreDrgRouteTableRepresentation)

	CoreCoreDrgRouteTableSingularDataSourceRepresentation = map[string]interface{}{
		"drg_route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
	}

	CoreCoreDrgRouteTableDataSourceRepresentation = map[string]interface{}{
		"drg_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"import_drg_route_distribution_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"state":                            acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgRouteTableDataSourceFilterRepresentation}}
	CoreDrgRouteTableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_drg_route_table.test_drg_route_table.id}`}},
	}

	CoreDrgRouteTableRepresentation = map[string]interface{}{
		"drg_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"import_drg_route_distribution_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"is_ecmp_enabled":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	CoreDrgRouteTableTriggerRepresentation = map[string]interface{}{
		"drg_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName3`, Update: `displayName4`},
		"is_ecmp_enabled":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"remove_import_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	CoreDrgRouteTableResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", acctest.Required, acctest.Create, CoreDrgRouteDistributionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, CoreDrgRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/pnp
func TestCoreDrgRouteTableResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgRouteTableResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_route_table.test_drg_route_table"
	datasourceName := "data.oci_core_drg_route_tables.test_drg_route_tables"
	singularDatasourceName := "data.oci_core_drg_route_table.test_drg_route_table"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreDrgRouteTableResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Optional, acctest.Create, CoreDrgRouteTableRepresentation), "core", "drgRouteTable", t)

	acctest.ResourceTest(t, testAccCheckCoreDrgRouteTableDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Required, acctest.Create, CoreDrgRouteTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Optional, acctest.Create, CoreDrgRouteTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "import_drg_route_distribution_id"),
				resource.TestCheckResourceAttr(resourceName, "is_ecmp_enabled", "false"),
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
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Optional, acctest.Update, CoreDrgRouteTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "import_drg_route_distribution_id", "0"),
				resource.TestCheckResourceAttr(resourceName, "is_ecmp_enabled", "true"),
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
		// verify remove import trigger
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Required, acctest.Create, CoreDrgRouteTableTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(resourceName, "import_drg_route_distribution_id"),
			),
		},
		// verify updates with import trigger
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Optional, acctest.Update, CoreDrgRouteTableTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "import_drg_route_distribution_id", ""),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgRouteTableResourceDependencies,
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_drg_route_tables", "test_drg_route_tables", acctest.Optional, acctest.Update, CoreCoreDrgRouteTableDataSourceRepresentation) +
				compartmentIdVariableStr + CoreDrgRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Optional, acctest.Update, CoreDrgRouteTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "import_drg_route_distribution_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "drg_route_tables.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_tables.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_tables.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_tables.0.drg_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_tables.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_tables.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_tables.0.import_drg_route_distribution_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_route_tables.0.is_ecmp_enabled", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_tables.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_tables.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Required, acctest.Create, CoreCoreDrgRouteTableSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreDrgRouteTableResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "drg_route_table_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_ecmp_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreDrgRouteTableRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreDrgRouteTableDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_drg_route_table" {
			noResourceFound = false
			request := oci_core.GetDrgRouteTableRequest{}

			tmp := rs.Primary.ID
			request.DrgRouteTableId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetDrgRouteTable(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.DrgRouteTableLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreDrgRouteTable") {
		resource.AddTestSweepers("CoreDrgRouteTable", &resource.Sweeper{
			Name:         "CoreDrgRouteTable",
			Dependencies: acctest.DependencyGraph["drgRouteTable"],
			F:            sweepCoreDrgRouteTableResource,
		})
	}
}

func sweepCoreDrgRouteTableResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	drgRouteTableIds, err := getCoreDrgRouteTableIds(compartment)
	if err != nil {
		return err
	}
	for _, drgRouteTableId := range drgRouteTableIds {
		if ok := acctest.SweeperDefaultResourceId[drgRouteTableId]; !ok {
			deleteDrgRouteTableRequest := oci_core.DeleteDrgRouteTableRequest{}

			deleteDrgRouteTableRequest.DrgRouteTableId = &drgRouteTableId

			deleteDrgRouteTableRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteDrgRouteTable(context.Background(), deleteDrgRouteTableRequest)
			if error != nil {
				fmt.Printf("Error deleting DrgRouteTable %s %s, It is possible that the resource is already deleted. Please verify manually \n", drgRouteTableId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &drgRouteTableId, CoreDrgRouteTableSweepWaitCondition, time.Duration(3*time.Minute),
				CoreDrgRouteTableSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreDrgRouteTableIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DrgRouteTableId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listDrgRouteTablesRequest := oci_core.ListDrgRouteTablesRequest{}
	// listDrgRouteTablesRequest.CompartmentId = &compartmentId

	drgIds, error := getCoreDrgIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting drgId required for DrgRouteTable resource requests \n")
	}
	for _, drgId := range drgIds {
		listDrgRouteTablesRequest.DrgId = &drgId

		listDrgRouteTablesRequest.LifecycleState = oci_core.DrgRouteTableLifecycleStateAvailable
		listDrgRouteTablesResponse, err := virtualNetworkClient.ListDrgRouteTables(context.Background(), listDrgRouteTablesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DrgRouteTable list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, drgRouteTable := range listDrgRouteTablesResponse.Items {
			id := *drgRouteTable.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DrgRouteTableId", id)
		}

	}
	return resourceIds, nil
}

func CoreDrgRouteTableSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if drgRouteTableResponse, ok := response.Response.(oci_core.GetDrgRouteTableResponse); ok {
		return drgRouteTableResponse.LifecycleState != oci_core.DrgRouteTableLifecycleStateTerminated
	}
	return false
}

func CoreDrgRouteTableSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetDrgRouteTable(context.Background(), oci_core.GetDrgRouteTableRequest{
		DrgRouteTableId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
