// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v31/common"
	oci_core "github.com/oracle/oci-go-sdk/v31/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	RouteTableRequiredOnlyResource = RouteTableResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation)

	RouteTableResource = RouteTableResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Optional, Create, routeTableRepresentation)

	routeTableDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `MyRouteTable`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"vcn_id":         Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, routeTableDataSourceFilterRepresentation}}
	routeTableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_route_table.test_route_table.id}`}},
	}

	routeTableRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"vcn_id":         Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `MyRouteTable`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"route_rules":    RepresentationGroup{Optional, routeTableRouteRulesRepresentation},
	}
	routeTableRouteRulesRepresentation = map[string]interface{}{
		"network_entity_id": Representation{repType: Required, create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
		"description":       Representation{repType: Optional, create: `description`, update: `description2`},
		"destination":       Representation{repType: Optional, create: `0.0.0.0/0`, update: `10.0.0.0/8`},
		"destination_type":  Representation{repType: Optional, create: `CIDR_BLOCK`},
	}

	RouteTableResourceDependencies = generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

func TestCoreRouteTableResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreRouteTableResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_route_table.test_route_table"
	datasourceName := "data.oci_core_route_tables.test_route_tables"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreRouteTableDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + RouteTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + RouteTableResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + RouteTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Optional, Create, routeTableRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyRouteTable"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"description":      "description",
						"destination":      "0.0.0.0/0",
						"destination_type": "CIDR_BLOCK",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RouteTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Optional, Create,
						representationCopyWithNewProperties(routeTableRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyRouteTable"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"description":      "description",
						"destination":      "0.0.0.0/0",
						"destination_type": "CIDR_BLOCK",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + RouteTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Optional, Update, routeTableRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"description":      "description2",
						"destination":      "10.0.0.0/8",
						"destination_type": "CIDR_BLOCK",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_route_tables", "test_route_tables", Optional, Update, routeTableDataSourceRepresentation) +
					compartmentIdVariableStr + RouteTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Optional, Update, routeTableRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "route_tables.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "route_tables.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(datasourceName, "route_tables.0.route_rules", map[string]string{
						"description":      "description2",
						"destination":      "10.0.0.0/8",
						"destination_type": "CIDR_BLOCK",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(datasourceName, "route_tables.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "route_tables.0.vcn_id"),
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
		},
	})
}

func testAccCheckCoreRouteTableDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_route_table" {
			noResourceFound = false
			request := oci_core.GetRouteTableRequest{}

			tmp := rs.Primary.ID
			request.RtId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			response, err := client.GetRouteTable(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.RouteTableLifecycleStateTerminated): true,
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
	if !inSweeperExcludeList("CoreRouteTable") {
		resource.AddTestSweepers("CoreRouteTable", &resource.Sweeper{
			Name:         "CoreRouteTable",
			Dependencies: DependencyGraph["routeTable"],
			F:            sweepCoreRouteTableResource,
		})
	}
}

func sweepCoreRouteTableResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	routeTableIds, err := getRouteTableIds(compartment)
	if err != nil {
		return err
	}
	for _, routeTableId := range routeTableIds {
		if ok := SweeperDefaultResourceId[routeTableId]; !ok {
			deleteRouteTableRequest := oci_core.DeleteRouteTableRequest{}

			deleteRouteTableRequest.RtId = &routeTableId

			deleteRouteTableRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteRouteTable(context.Background(), deleteRouteTableRequest)
			if error != nil {
				fmt.Printf("Error deleting RouteTable %s %s, It is possible that the resource is already deleted. Please verify manually \n", routeTableId, error)
				continue
			}
			waitTillCondition(testAccProvider, &routeTableId, routeTableSweepWaitCondition, time.Duration(3*time.Minute),
				routeTableSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getRouteTableIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "RouteTableId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listRouteTablesRequest := oci_core.ListRouteTablesRequest{}
	listRouteTablesRequest.CompartmentId = &compartmentId
	listRouteTablesRequest.LifecycleState = oci_core.RouteTableLifecycleStateAvailable
	listRouteTablesResponse, err := virtualNetworkClient.ListRouteTables(context.Background(), listRouteTablesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting RouteTable list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, routeTable := range listRouteTablesResponse.Items {
		id := *routeTable.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "RouteTableId", id)
	}
	return resourceIds, nil
}

func routeTableSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if routeTableResponse, ok := response.Response.(oci_core.GetRouteTableResponse); ok {
		return routeTableResponse.LifecycleState != oci_core.RouteTableLifecycleStateTerminated
	}
	return false
}

func routeTableSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetRouteTable(context.Background(), oci_core.GetRouteTableRequest{
		RtId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
