// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v45/common"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v45/usageapi"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CustomTableResourceConfig = CustomTableResourceDependencies +
		generateResourceFromRepresentationMap("oci_metering_computation_custom_table", "test_custom_table", Optional, Update, customTableRepresentation)

	customTableSingularDataSourceRepresentation = map[string]interface{}{
		"custom_table_id": Representation{repType: Required, create: `${oci_metering_computation_custom_table.test_custom_table.id}`},
	}

	customTableDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  Representation{repType: Required, create: `${var.compartment_id}`},
		"saved_report_id": Representation{repType: Required, create: `savedReportId`},
		"filter":          RepresentationGroup{Required, customTableDataSourceFilterRepresentation}}
	customTableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_metering_computation_custom_table.test_custom_table.id}`}},
	}

	customTableRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"saved_custom_table": RepresentationGroup{Required, customTableSavedCustomTableRepresentation},
		"saved_report_id":    Representation{repType: Required, create: `savedReportId`},
	}
	customTableSavedCustomTableRepresentation = map[string]interface{}{
		"display_name":      Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"column_group_by":   Representation{repType: Required, create: []string{`columnGroupBy`}, update: []string{`columnGroupBy2`}},
		"compartment_depth": Representation{repType: Required, create: `1.0`, update: `2.0`},
		"group_by_tag":      RepresentationGroup{Optional, customTableSavedCustomTableGroupByTagRepresentation},
		"row_group_by":      Representation{repType: Required, create: []string{`rowGroupBy`}, update: []string{}},
		"version":           Representation{repType: Required, create: `1.0`, update: `1.0`},
	}
	customTableSavedCustomTableGroupByTagRepresentation = map[string]interface{}{
		"key":       Representation{repType: Optional, create: `key`, update: `key2`},
		"namespace": Representation{repType: Optional, create: `namespace`, update: `namespace2`},
		"value":     Representation{repType: Optional, create: `value`, update: `value2`},
	}

	CustomTableResourceDependencies = ""
)

func TestMeteringComputationCustomTableResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationCustomTableResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_metering_computation_custom_table.test_custom_table"
	datasourceName := "data.oci_metering_computation_custom_tables.test_custom_tables"
	singularDatasourceName := "data.oci_metering_computation_custom_table.test_custom_table"

	var resId, resId2 string
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+CustomTableResourceDependencies+
		generateResourceFromRepresentationMap("oci_metering_computation_custom_table", "test_custom_table", Required, Create, customTableRepresentation), "usageapi", "customTable", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckMeteringComputationCustomTableDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + CustomTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_metering_computation_custom_table", "test_custom_table", Required, Create, customTableRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "saved_custom_table.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "saved_custom_table.0.display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "saved_report_id"),

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
				Config: config + compartmentIdVariableStr + CustomTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_metering_computation_custom_table", "test_custom_table", Optional, Update, customTableRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "saved_custom_table.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "saved_custom_table.0.column_group_by.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "saved_custom_table.0.compartment_depth", "2"),
					resource.TestCheckResourceAttr(resourceName, "saved_custom_table.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "saved_custom_table.0.group_by_tag.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "saved_custom_table.0.group_by_tag.0.key", "key2"),
					resource.TestCheckResourceAttr(resourceName, "saved_custom_table.0.version", "1"),
					resource.TestCheckResourceAttr(resourceName, "saved_custom_table.0.row_group_by.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "saved_report_id"),

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
					generateDataSourceFromRepresentationMap("oci_metering_computation_custom_tables", "test_custom_tables", Optional, Update, customTableDataSourceRepresentation) +
					compartmentIdVariableStr + CustomTableResourceDependencies +
					generateResourceFromRepresentationMap("oci_metering_computation_custom_table", "test_custom_table", Optional, Update, customTableRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "saved_report_id"),

					resource.TestCheckResourceAttr(datasourceName, "custom_table_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "custom_table_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_metering_computation_custom_table", "test_custom_table", Required, Create, customTableSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CustomTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_table_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "saved_custom_table.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "saved_custom_table.0.column_group_by.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "saved_custom_table.0.compartment_depth", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "saved_custom_table.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "saved_custom_table.0.group_by_tag.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "saved_custom_table.0.group_by_tag.0.key", "key2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "saved_custom_table.0.group_by_tag.0.namespace", "namespace2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "saved_custom_table.0.group_by_tag.0.value", "value2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "saved_custom_table.0.version", "1"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + CustomTableResourceConfig,
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

func testAccCheckMeteringComputationCustomTableDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).usageapiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_metering_computation_custom_table" {
			noResourceFound = false
			request := oci_metering_computation.GetCustomTableRequest{}

			tmp := rs.Primary.ID
			request.CustomTableId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "metering_computation")

			_, err := client.GetCustomTable(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !inSweeperExcludeList("MeteringComputationCustomTable") {
		resource.AddTestSweepers("MeteringComputationCustomTable", &resource.Sweeper{
			Name:         "MeteringComputationCustomTable",
			Dependencies: DependencyGraph["customTable"],
			F:            sweepMeteringComputationCustomTableResource,
		})
	}
}

func sweepMeteringComputationCustomTableResource(compartment string) error {
	usageapiClient := GetTestClients(&schema.ResourceData{}).usageapiClient()
	customTableIds, err := getCustomTableIds(compartment)
	if err != nil {
		return err
	}
	for _, customTableId := range customTableIds {
		if ok := SweeperDefaultResourceId[customTableId]; !ok {
			deleteCustomTableRequest := oci_metering_computation.DeleteCustomTableRequest{}

			deleteCustomTableRequest.CustomTableId = &customTableId

			deleteCustomTableRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "metering_computation")
			_, error := usageapiClient.DeleteCustomTable(context.Background(), deleteCustomTableRequest)
			if error != nil {
				fmt.Printf("Error deleting CustomTable %s %s, It is possible that the resource is already deleted. Please verify manually \n", customTableId, error)
				continue
			}
		}
	}
	return nil
}

func getCustomTableIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "CustomTableId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	usageapiClient := GetTestClients(&schema.ResourceData{}).usageapiClient()

	listCustomTablesRequest := oci_metering_computation.ListCustomTablesRequest{}
	listCustomTablesRequest.CompartmentId = &compartmentId

	savedReportIds, error := getQueryIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting savedReportId required for CustomTable resource requests \n")
	}
	for _, savedReportId := range savedReportIds {
		listCustomTablesRequest.SavedReportId = &savedReportId

		listCustomTablesResponse, err := usageapiClient.ListCustomTables(context.Background(), listCustomTablesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting CustomTable list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, customTable := range listCustomTablesResponse.Items {
			id := *customTable.Id
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "CustomTableId", id)
		}

	}
	return resourceIds, nil
}
