// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafeMaskingPoliciesMaskingColumnRequiredOnlyResource = DataSafeMaskingPoliciesMaskingColumnResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policies_masking_column", "test_masking_policies_masking_column", acctest.Required, acctest.Create, maskingPoliciesMaskingColumnRepresentation)

	DataSafeMaskingPoliciesMaskingColumnResourceConfig = DataSafeMaskingPoliciesMaskingColumnResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policies_masking_column", "test_masking_policies_masking_column", acctest.Optional, acctest.Update, maskingPoliciesMaskingColumnRepresentation)

	DataSafemaskingPoliciesMaskingColumnSingularDataSourceRepresentation = map[string]interface{}{
		"masking_column_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_masking_policies_masking_column.test_masking_policies_masking_column.key}`},
		"masking_policy_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_masking_policy.test_masking_policy.id}`},
	}

	DataSafemaskingPoliciesMaskingColumnDataSourceRepresentation = map[string]interface{}{
		"masking_policy_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_masking_policy.test_masking_policy.id}`},
		"is_masking_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"schema_name":        acctest.Representation{RepType: acctest.Optional, Create: []string{`HCM`}},
		"object":             acctest.Representation{RepType: acctest.Optional, Create: []string{`EMPLOYEES`}},
		"column_name":        acctest.Representation{RepType: acctest.Optional, Create: []string{`FIRST_NAME`}},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: maskingPoliciesMaskingColumnDataSourceFilterRepresentation}}
	maskingPoliciesMaskingColumnDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_masking_policies_masking_column.test_masking_policies_masking_column.key}`}},
	}

	maskingPoliciesMaskingColumnRepresentation = map[string]interface{}{
		"column_name":       acctest.Representation{RepType: acctest.Required, Create: `FIRST_NAME`},
		"masking_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_masking_policy.test_masking_policy.id}`},
		"object":            acctest.Representation{RepType: acctest.Required, Create: `EMPLOYEES`},
		"schema_name":       acctest.Representation{RepType: acctest.Required, Create: `HCM`},
		"masking_formats":   acctest.RepresentationGroup{RepType: acctest.Required, Group: maskingPoliciesMaskingColumnMaskingFormatsRepresentation},
	}

	maskingPoliciesMaskingColumnMaskingFormatsRepresentation = map[string]interface{}{
		"format_entries": acctest.RepresentationGroup{RepType: acctest.Required, Group: maskingPoliciesMaskingColumnMaskingFormatsFormatEntriesRepresentation},
	}

	maskingPoliciesMaskingColumnMaskingFormatsFormatEntriesRepresentation = map[string]interface{}{
		"type":         acctest.Representation{RepType: acctest.Required, Create: `FIXED_STRING`, Update: `FIXED_STRING`},
		"fixed_string": acctest.Representation{RepType: acctest.Required, Create: `FixedStringName`, Update: `FixedStringUpdate`},
	}

	DataSafeMaskingPoliciesMaskingColumnMaskingFormatsFormatEntriesRepresentation = map[string]interface{}{
		"type":                      acctest.Representation{RepType: acctest.Required, Create: `DELETE_ROWS`, Update: `DETERMINISTIC_SUBSTITUTION`},
		"column_name":               acctest.Representation{RepType: acctest.Optional, Create: `columnName`, Update: `columnName2`},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"end_date":                  acctest.Representation{RepType: acctest.Optional, Create: `endDate`, Update: `endDate2`},
		"end_length":                acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"end_value":                 acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `1.1`},
		"fixed_number":              acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `1.1`},
		"fixed_string":              acctest.Representation{RepType: acctest.Optional, Create: `fixedString`, Update: `fixedString2`},
		"grouping_columns":          acctest.Representation{RepType: acctest.Optional, Create: []string{`groupingColumns`}, Update: []string{`groupingColumns2`}},
		"length":                    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"library_masking_format_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_library_masking_format.test_library_masking_format.id}`},
		"pattern":                   acctest.Representation{RepType: acctest.Optional, Create: `pattern`, Update: `pattern2`},
		"post_processing_function":  acctest.Representation{RepType: acctest.Optional, Create: `postProcessingFunction`, Update: `postProcessingFunction2`},
		"random_list":               acctest.Representation{RepType: acctest.Optional, Create: []string{`randomList`}, Update: []string{`randomList2`}},
		"regular_expression":        acctest.Representation{RepType: acctest.Optional, Create: `regularExpression`, Update: `regularExpression2`},
		"replace_with":              acctest.Representation{RepType: acctest.Optional, Create: `replaceWith`, Update: `replaceWith2`},
		"schema_name":               acctest.Representation{RepType: acctest.Optional, Create: `schemaName`, Update: `schemaName2`},
		"sql_expression":            acctest.Representation{RepType: acctest.Optional, Create: `sqlExpression`, Update: `sqlExpression2`},
		"start_date":                acctest.Representation{RepType: acctest.Optional, Create: `startDate`, Update: `startDate2`},
		"start_length":              acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"start_position":            acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"start_value":               acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `1.1`},
		"table_name":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_nosql_table.test_table.name}`},
		"user_defined_function":     acctest.Representation{RepType: acctest.Optional, Create: `userDefinedFunction`, Update: `userDefinedFunction2`},
	}

	DataSafeMaskingPoliciesMaskingColumnResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Required, acctest.Create, maskingPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model1", acctest.Required, acctest.Create, sensitiveDataModelRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeMaskingPoliciesMaskingColumnResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeMaskingPoliciesMaskingColumnResource_basic	")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_masking_policies_masking_column.test_masking_policies_masking_column"
	datasourceName := "data.oci_data_safe_masking_policies_masking_columns.test_masking_policies_masking_columns"
	singularDatasourceName := "data.oci_data_safe_masking_policies_masking_column.test_masking_policies_masking_column"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeMaskingPoliciesMaskingColumnResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policies_masking_column", "test_masking_policies_masking_column", acctest.Optional, acctest.Create, maskingPoliciesMaskingColumnRepresentation), "datasafe", "maskingPoliciesMaskingColumn", t)

	acctest.ResourceTest(t, testAccCheckDataSafeMaskingPoliciesMaskingColumnDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeMaskingPoliciesMaskingColumnResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policies_masking_column", "test_masking_policies_masking_column", acctest.Required, acctest.Create, maskingPoliciesMaskingColumnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_name", "FIRST_NAME"),
				resource.TestCheckResourceAttrSet(resourceName, "masking_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "schema_name", "HCM"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "key")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeMaskingPoliciesMaskingColumnResourceDependencies + targetIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeMaskingPoliciesMaskingColumnResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policies_masking_column", "test_masking_policies_masking_column", acctest.Optional, acctest.Create, maskingPoliciesMaskingColumnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_name", "FIRST_NAME"),
				resource.TestCheckResourceAttr(resourceName, "is_masking_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.condition", "1=1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.0.fixed_string", "FixedStringName"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.0.type", "FIXED_STRING"),
				resource.TestCheckResourceAttr(resourceName, "object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "object_type", "TABLE"),
				resource.TestCheckResourceAttr(resourceName, "schema_name", "HCM"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "key")
					var compositeId string
					compositeId, err = acctest.FromInstanceState(s, resourceName, "id")
					prefix := "oci_data_safe_masking_policies_masking_column:"
					fullPath := prefix + compositeId
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&fullPath, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DataSafeMaskingPoliciesMaskingColumnResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policies_masking_column", "test_masking_policies_masking_column", acctest.Optional, acctest.Update, maskingPoliciesMaskingColumnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_name", "FIRST_NAME"),
				resource.TestCheckResourceAttr(resourceName, "is_masking_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.condition", "1=1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.0.fixed_string", "FixedStringUpdate"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.0.type", "FIXED_STRING"),
				resource.TestCheckResourceAttrSet(resourceName, "masking_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "object_type", "TABLE"),
				resource.TestCheckResourceAttr(resourceName, "schema_name", "HCM"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "key")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_policies_masking_columns", "test_masking_policies_masking_columns", acctest.Optional, acctest.Update, DataSafemaskingPoliciesMaskingColumnDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeMaskingPoliciesMaskingColumnResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policies_masking_column", "test_masking_policies_masking_column", acctest.Optional, acctest.Update, maskingPoliciesMaskingColumnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "is_masking_enabled", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "masking_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "schema_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "object.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "column_name.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "masking_column_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "masking_column_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_policies_masking_column", "test_masking_policies_masking_column", acctest.Required, acctest.Create, DataSafemaskingPoliciesMaskingColumnSingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeMaskingPoliciesMaskingColumnResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "masking_column_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "masking_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "column_name", "FIRST_NAME"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_masking_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.0.condition", "1=1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.0.format_entries.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.0.format_entries.0.fixed_string", "FixedStringUpdate"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.0.format_entries.0.type", "FIXED_STRING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_type", "TABLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema_name", "HCM"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeMaskingPoliciesMaskingColumnResourceConfig + targetIdVariableStr,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeMaskingPoliciesMaskingColumnDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_masking_policies_masking_column" {
			noResourceFound = false
			request := oci_data_safe.GetMaskingColumnRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.MaskingColumnKey = &value
			}

			if value, ok := rs.Primary.Attributes["masking_policy_id"]; ok {
				request.MaskingPolicyId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetMaskingColumn(context.Background(), request)

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DataSafeMaskingPoliciesMaskingColumn") {
		resource.AddTestSweepers("DataSafeMaskingPoliciesMaskingColumn", &resource.Sweeper{
			Name:         "DataSafeMaskingPoliciesMaskingColumn",
			Dependencies: acctest.DependencyGraph["maskingPoliciesMaskingColumn"],
			F:            sweepDataSafeMaskingPoliciesMaskingColumnResource,
		})
	}
}

func sweepDataSafeMaskingPoliciesMaskingColumnResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	maskingPoliciesMaskingColumnIds, err := getDataSafeMaskingPoliciesMaskingColumnIds(compartment)
	if err != nil {
		return err
	}
	for _, maskingPoliciesMaskingColumnId := range maskingPoliciesMaskingColumnIds {
		if ok := acctest.SweeperDefaultResourceId[maskingPoliciesMaskingColumnId]; !ok {
			deleteMaskingColumnRequest := oci_data_safe.DeleteMaskingColumnRequest{}

			deleteMaskingColumnRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteMaskingColumn(context.Background(), deleteMaskingColumnRequest)
			if error != nil {
				fmt.Printf("Error deleting MaskingPoliciesMaskingColumn %s %s, It is possible that the resource is already deleted. Please verify manually \n", maskingPoliciesMaskingColumnId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeMaskingPoliciesMaskingColumnIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MaskingPoliciesMaskingColumnId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listMaskingColumnsRequest := oci_data_safe.ListMaskingColumnsRequest{}
	// listMaskingColumnsRequest.CompartmentId = &compartmentId

	maskingPolicyIds, error := getDataSafeMaskingPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting maskingPolicyId required for MaskingPoliciesMaskingColumn resource requests \n")
	}
	for _, maskingPolicyId := range maskingPolicyIds {
		listMaskingColumnsRequest.MaskingPolicyId = &maskingPolicyId

		listMaskingColumnsResponse, err := dataSafeClient.ListMaskingColumns(context.Background(), listMaskingColumnsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting MaskingPoliciesMaskingColumn list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, maskingPoliciesMaskingColumn := range listMaskingColumnsResponse.Items {
			id := *maskingPoliciesMaskingColumn.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MaskingPoliciesMaskingColumnId", id)
		}

	}
	return resourceIds, nil
}
