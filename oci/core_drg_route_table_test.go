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
	"github.com/oracle/oci-go-sdk/v46/common"
	oci_core "github.com/oracle/oci-go-sdk/v46/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DrgRouteTableRequiredOnlyResource = DrgRouteTableResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Required, Create, drgRouteTableRepresentation)

	DrgRouteTableResourceConfig = DrgRouteTableResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Optional, Update, drgRouteTableRepresentation)

	drgRouteTableSingularDataSourceRepresentation = map[string]interface{}{
		"drg_route_table_id": Representation{repType: Required, create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
	}

	drgRouteTableDataSourceRepresentation = map[string]interface{}{
		"drg_id":                           Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
		"display_name":                     Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"import_drg_route_distribution_id": Representation{repType: Optional, create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"state":                            Representation{repType: Optional, create: `AVAILABLE`},
		"filter":                           RepresentationGroup{Required, drgRouteTableDataSourceFilterRepresentation}}
	drgRouteTableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_drg_route_table.test_drg_route_table.id}`}},
	}

	drgRouteTableRepresentation = map[string]interface{}{
		"drg_id":                           Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
		"defined_tags":                     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                     Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":                    Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"import_drg_route_distribution_id": Representation{repType: Optional, create: `${oci_core_drg_route_distribution.test_drg_route_distribution.id}`},
		"is_ecmp_enabled":                  Representation{repType: Optional, create: `false`, update: `true`},
	}

	drgRouteTableTriggerRepresentation = map[string]interface{}{
		"drg_id":                Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
		"display_name":          Representation{repType: Optional, create: `displayName3`, update: `displayName4`},
		"is_ecmp_enabled":       Representation{repType: Optional, create: `false`, update: `true`},
		"remove_import_trigger": Representation{repType: Optional, create: `false`, update: `true`},
	}

	DrgRouteTableResourceDependencies = generateResourceFromRepresentationMap("oci_core_drg_route_distribution", "test_drg_route_distribution", Required, Create, drgRouteDistributionRepresentation) +
		generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/pnp
func TestCoreDrgRouteTableResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgRouteTableResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_route_table.test_drg_route_table"
	datasourceName := "data.oci_core_drg_route_tables.test_drg_route_tables"
	singularDatasourceName := "data.oci_core_drg_route_table.test_drg_route_table"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DrgRouteTableResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Optional, Create, drgRouteTableRepresentation), "core", "drgRouteTable", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreDrgRouteTableDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DrgRouteTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Required, Create, drgRouteTableRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DrgRouteTableResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DrgRouteTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Optional, Create, drgRouteTableRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "import_drg_route_distribution_id"),
					resource.TestCheckResourceAttr(resourceName, "is_ecmp_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DrgRouteTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Optional, Update, drgRouteTableRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					//resource.TestCheckResourceAttr(resourceName, "import_drg_route_distribution_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "is_ecmp_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify remove import trigger
			{
				Config: config + compartmentIdVariableStr + DrgRouteTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Required, Create, drgRouteTableTriggerRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "import_drg_route_distribution_id"),
				),
			},
			// verify updates with import trigger
			{
				Config: config + compartmentIdVariableStr + DrgRouteTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Optional, Update, drgRouteTableTriggerRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "import_drg_route_distribution_id", ""),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DrgRouteTableResourceDependencies,
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_drg_route_tables", "test_drg_route_tables", Optional, Update, drgRouteTableDataSourceRepresentation) +
					compartmentIdVariableStr + DrgRouteTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Optional, Update, drgRouteTableRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "import_drg_route_distribution_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "drg_route_tables.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_route_tables.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "drg_route_tables.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Required, Create, drgRouteTableSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DrgRouteTableResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "drg_route_table_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_ecmp_enabled", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + DrgRouteTableResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreDrgRouteTableDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_drg_route_table" {
			noResourceFound = false
			request := oci_core.GetDrgRouteTableRequest{}

			tmp := rs.Primary.ID
			request.DrgRouteTableId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreDrgRouteTable") {
		resource.AddTestSweepers("CoreDrgRouteTable", &resource.Sweeper{
			Name:         "CoreDrgRouteTable",
			Dependencies: DependencyGraph["drgRouteTable"],
			F:            sweepCoreDrgRouteTableResource,
		})
	}
}

func sweepCoreDrgRouteTableResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	drgRouteTableIds, err := getDrgRouteTableIds(compartment)
	if err != nil {
		return err
	}
	for _, drgRouteTableId := range drgRouteTableIds {
		if ok := SweeperDefaultResourceId[drgRouteTableId]; !ok {
			deleteDrgRouteTableRequest := oci_core.DeleteDrgRouteTableRequest{}

			deleteDrgRouteTableRequest.DrgRouteTableId = &drgRouteTableId

			deleteDrgRouteTableRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteDrgRouteTable(context.Background(), deleteDrgRouteTableRequest)
			if error != nil {
				fmt.Printf("Error deleting DrgRouteTable %s %s, It is possible that the resource is already deleted. Please verify manually \n", drgRouteTableId, error)
				continue
			}
			waitTillCondition(testAccProvider, &drgRouteTableId, drgRouteTableSweepWaitCondition, time.Duration(3*time.Minute),
				drgRouteTableSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getDrgRouteTableIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DrgRouteTableId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listDrgRouteTablesRequest := oci_core.ListDrgRouteTablesRequest{}
	// listDrgRouteTablesRequest.CompartmentId = &compartmentId

	drgIds, error := getDrgIds(compartment)
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
			addResourceIdToSweeperResourceIdMap(compartmentId, "DrgRouteTableId", id)
		}

	}
	return resourceIds, nil
}

func drgRouteTableSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if drgRouteTableResponse, ok := response.Response.(oci_core.GetDrgRouteTableResponse); ok {
		return drgRouteTableResponse.LifecycleState != oci_core.DrgRouteTableLifecycleStateTerminated
	}
	return false
}

func drgRouteTableSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetDrgRouteTable(context.Background(), oci_core.GetDrgRouteTableRequest{
		DrgRouteTableId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
