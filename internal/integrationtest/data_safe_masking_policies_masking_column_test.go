// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
		"schema_name":        acctest.Representation{RepType: acctest.Optional, Create: []string{`ADMIN`}},
		"object":             acctest.Representation{RepType: acctest.Optional, Create: []string{`LOCATIONS`}},
		"column_name":        acctest.Representation{RepType: acctest.Optional, Create: []string{`STREET_ADDRESS`}},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: maskingPoliciesMaskingColumnDataSourceFilterRepresentation}}
	maskingPoliciesMaskingColumnDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_masking_policies_masking_column.test_masking_policies_masking_column.key}`}},
	}

	maskingPoliciesMaskingColumnRepresentation = map[string]interface{}{
		"column_name":       acctest.Representation{RepType: acctest.Required, Create: `STREET_ADDRESS`},
		"masking_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_masking_policy.test_masking_policy.id}`},
		"object":            acctest.Representation{RepType: acctest.Required, Create: `LOCATIONS`},
		"schema_name":       acctest.Representation{RepType: acctest.Required, Create: `ADMIN`},
		"masking_formats":   acctest.RepresentationGroup{RepType: acctest.Required, Group: maskingPoliciesMaskingColumnMaskingFormatsRepresentation},
	}
	maskingPoliciesMaskingColumnMaskingFormatsRepresentation = map[string]interface{}{
		"format_entries": acctest.RepresentationGroup{RepType: acctest.Required, Group: maskingPoliciesMaskingColumnMaskingFormatsFormatEntriesRepresentation},
	}
	maskingPoliciesMaskingColumnMaskingFormatsFormatEntriesRepresentation = map[string]interface{}{
		"type":         acctest.Representation{RepType: acctest.Required, Create: `RANDOM_STRING`, Update: `RANDOM_STRING`},
		"end_length":   acctest.Representation{RepType: acctest.Required, Create: `50`, Update: `60`},
		"start_length": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `20`},
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
				resource.TestCheckResourceAttr(resourceName, "column_name", "STREET_ADDRESS"),
				resource.TestCheckResourceAttrSet(resourceName, "masking_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "object", "LOCATIONS"),
				resource.TestCheckResourceAttr(resourceName, "schema_name", "ADMIN"),

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
				resource.TestCheckResourceAttr(resourceName, "column_name", "STREET_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "is_masking_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.condition", "1=1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.0.end_length", "50"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.0.start_length", "1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.0.type", "RANDOM_STRING"),
				resource.TestCheckResourceAttr(resourceName, "object", "LOCATIONS"),
				resource.TestCheckResourceAttr(resourceName, "object_type", "TABLE"),
				resource.TestCheckResourceAttr(resourceName, "schema_name", "ADMIN"),
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
				resource.TestCheckResourceAttr(resourceName, "column_name", "STREET_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "is_masking_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.condition", "1=1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.0.end_length", "60"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.0.grouping_columns.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.0.start_length", "20"),
				resource.TestCheckResourceAttr(resourceName, "masking_formats.0.format_entries.0.type", "RANDOM_STRING"),
				resource.TestCheckResourceAttrSet(resourceName, "masking_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "object", "LOCATIONS"),
				resource.TestCheckResourceAttr(resourceName, "object_type", "TABLE"),
				resource.TestCheckResourceAttr(resourceName, "schema_name", "ADMIN"),
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

				resource.TestCheckResourceAttr(singularDatasourceName, "column_name", "STREET_ADDRESS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_masking_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.0.condition", "1=1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.0.format_entries.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.0.format_entries.0.end_length", "60"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.0.format_entries.0.grouping_columns.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.0.format_entries.0.length", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.0.format_entries.0.start_length", "20"),
				resource.TestCheckResourceAttr(singularDatasourceName, "masking_formats.0.format_entries.0.type", "RANDOM_STRING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object", "LOCATIONS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_type", "TABLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema_name", "ADMIN"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DataSafeMaskingPoliciesMaskingColumnResourceConfig + targetIdVariableStr,
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
