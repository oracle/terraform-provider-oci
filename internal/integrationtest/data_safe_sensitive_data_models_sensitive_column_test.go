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
	DataSafeSensitiveDataModelsSensitiveColumnRequiredOnlyResource = DataSafeSensitiveDataModelsSensitiveColumnResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_models_sensitive_column", "test_sensitive_data_models_sensitive_column", acctest.Required, acctest.Create, sensitiveDataModelsSensitiveColumnRepresentation)

	DataSafeSensitiveDataModelsSensitiveColumnResourceConfig = DataSafeSensitiveDataModelsSensitiveColumnResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_models_sensitive_column", "test_sensitive_data_models_sensitive_column", acctest.Optional, acctest.Update, sensitiveDataModelsSensitiveColumnRepresentation)

	DataSafesensitiveDataModelsSensitiveColumnSingularDataSourceRepresentation = map[string]interface{}{
		"sensitive_column_key":    acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_data_models_sensitive_column.test_sensitive_data_models_sensitive_column.key}`},
		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
	}

	DataSafesensitiveDataModelsSensitiveColumnDataSourceRepresentation = map[string]interface{}{

		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
		"column_name":             acctest.Representation{RepType: acctest.Optional, Create: []string{`FIRST_NAME`}},
		"object":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`EMPLOYEES`}},
		"schema_name":             acctest.Representation{RepType: acctest.Optional, Create: []string{`ADMIN`}},
		"time_created_less_than":  acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: sensitiveDataModelsSensitiveColumnDataSourceFilterRepresentation}}
	sensitiveDataModelsSensitiveColumnDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_sensitive_data_models_sensitive_column.test_sensitive_data_models_sensitive_column.key}`}},
	}

	sensitiveDataModelsSensitiveColumnRepresentation = map[string]interface{}{
		"column_name":             acctest.Representation{RepType: acctest.Required, Create: `FIRST_NAME`},
		"object":                  acctest.Representation{RepType: acctest.Required, Create: `EMPLOYEES`},
		"schema_name":             acctest.Representation{RepType: acctest.Required, Create: `ADMIN`},
		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
		"app_name":                acctest.Representation{RepType: acctest.Optional, Create: `ADMIN`},
		"data_type":               acctest.Representation{RepType: acctest.Optional, Create: `VARCHAR2`, Update: `VARCHAR2`},
		"object_type":             acctest.Representation{RepType: acctest.Optional, Create: `TABLE`},
		"relation_type":           acctest.Representation{RepType: acctest.Optional, Create: `NONE`, Update: `NONE`},
		"status":                  acctest.Representation{RepType: acctest.Optional, Create: `VALID`, Update: `INVALID`},
	}

	DataSafeSensitiveDataModelsSensitiveColumnResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Required, acctest.Create, sensitiveDataModelRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveDataModelsSensitiveColumnResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSensitiveDataModelsSensitiveColumnResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_sensitive_data_models_sensitive_column.test_sensitive_data_models_sensitive_column"
	datasourceName := "data.oci_data_safe_sensitive_data_models_sensitive_columns.test_sensitive_data_models_sensitive_columns"
	singularDatasourceName := "data.oci_data_safe_sensitive_data_models_sensitive_column.test_sensitive_data_models_sensitive_column"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+targetIdVariableStr+DataSafeSensitiveDataModelsSensitiveColumnResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_models_sensitive_column", "test_sensitive_data_models_sensitive_column", acctest.Optional, acctest.Create, sensitiveDataModelsSensitiveColumnRepresentation), "datasafe", "sensitiveDataModelsSensitiveColumn", t)
	acctest.ResourceTest(t, testAccCheckDataSafeSensitiveDataModelsSensitiveColumnDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSensitiveDataModelsSensitiveColumnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_models_sensitive_column", "test_sensitive_data_models_sensitive_column", acctest.Required, acctest.Create, sensitiveDataModelsSensitiveColumnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_name", "FIRST_NAME"),
				resource.TestCheckResourceAttr(resourceName, "object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "schema_name", "ADMIN"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_data_model_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "key")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + targetIdVariableStr + compartmentIdVariableStr + DataSafeSensitiveDataModelsSensitiveColumnResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSensitiveDataModelsSensitiveColumnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_models_sensitive_column", "test_sensitive_data_models_sensitive_column", acctest.Optional, acctest.Create, sensitiveDataModelsSensitiveColumnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "app_name", "ADMIN"),
				resource.TestCheckResourceAttr(resourceName, "column_name", "FIRST_NAME"),
				resource.TestCheckResourceAttr(resourceName, "data_type", "VARCHAR2"),
				resource.TestCheckResourceAttrSet(resourceName, "estimated_data_value_count"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "object_type", "TABLE"),
				resource.TestCheckResourceAttr(resourceName, "schema_name", "ADMIN"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source"),
				resource.TestCheckResourceAttr(resourceName, "status", "VALID"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "key")
					var compositeId string
					compositeId, err = acctest.FromInstanceState(s, resourceName, "id")
					prefix := "oci_data_safe_sensitive_data_models_sensitive_column:"
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
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSensitiveDataModelsSensitiveColumnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_models_sensitive_column", "test_sensitive_data_models_sensitive_column", acctest.Optional, acctest.Update, sensitiveDataModelsSensitiveColumnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "app_name", "ADMIN"),
				resource.TestCheckResourceAttr(resourceName, "column_name", "FIRST_NAME"),
				resource.TestCheckResourceAttrSet(resourceName, "estimated_data_value_count"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "object_type", "TABLE"),
				resource.TestCheckResourceAttr(resourceName, "relation_type", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "schema_name", "ADMIN"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source"),
				resource.TestCheckResourceAttr(resourceName, "status", "INVALID"),
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
			Config: config + targetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_data_models_sensitive_columns", "test_sensitive_data_models_sensitive_columns", acctest.Optional, acctest.Update, DataSafesensitiveDataModelsSensitiveColumnDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSensitiveDataModelsSensitiveColumnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_models_sensitive_column", "test_sensitive_data_models_sensitive_column", acctest.Optional, acctest.Update, sensitiveDataModelsSensitiveColumnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "column_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "object.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "schema_name.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "sensitive_column_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sensitive_column_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_data_models_sensitive_column", "test_sensitive_data_models_sensitive_column", acctest.Required, acctest.Create, DataSafesensitiveDataModelsSensitiveColumnSingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSensitiveDataModelsSensitiveColumnResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_name", "FIRST_NAME"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_defined_child_column_keys.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "estimated_data_value_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_type", "TABLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parent_column_keys.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema_name", "ADMIN"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "INVALID"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + targetIdVariableStr + DataSafeSensitiveDataModelsSensitiveColumnRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSensitiveDataModelsSensitiveColumnDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_sensitive_data_models_sensitive_column" {
			noResourceFound = false
			request := oci_data_safe.GetSensitiveColumnRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.SensitiveColumnKey = &value
			}

			if value, ok := rs.Primary.Attributes["sensitive_data_model_id"]; ok {
				request.SensitiveDataModelId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetSensitiveColumn(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataSafeSensitiveDataModelsSensitiveColumn") {
		resource.AddTestSweepers("DataSafeSensitiveDataModelsSensitiveColumn", &resource.Sweeper{
			Name:         "DataSafeSensitiveDataModelsSensitiveColumn",
			Dependencies: acctest.DependencyGraph["sensitiveDataModelsSensitiveColumn"],
			F:            sweepDataSafeSensitiveDataModelsSensitiveColumnResource,
		})
	}
}

func sweepDataSafeSensitiveDataModelsSensitiveColumnResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	sensitiveDataModelsSensitiveColumnIds, err := getDataSafeSensitiveDataModelsSensitiveColumnIds(compartment)
	if err != nil {
		return err
	}
	for _, sensitiveDataModelsSensitiveColumnId := range sensitiveDataModelsSensitiveColumnIds {
		if ok := acctest.SweeperDefaultResourceId[sensitiveDataModelsSensitiveColumnId]; !ok {
			deleteSensitiveColumnRequest := oci_data_safe.DeleteSensitiveColumnRequest{}

			deleteSensitiveColumnRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSensitiveColumn(context.Background(), deleteSensitiveColumnRequest)
			if error != nil {
				fmt.Printf("Error deleting SensitiveDataModelsSensitiveColumn %s %s, It is possible that the resource is already deleted. Please verify manually \n", sensitiveDataModelsSensitiveColumnId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeSensitiveDataModelsSensitiveColumnIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SensitiveDataModelsSensitiveColumnId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSensitiveColumnsRequest := oci_data_safe.ListSensitiveColumnsRequest{}

	sensitiveDataModelIds, error := getDataSafeSensitiveDataModelIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting sensitiveDataModelId required for SensitiveDataModelsSensitiveColumn resource requests \n")
	}
	for _, sensitiveDataModelId := range sensitiveDataModelIds {
		listSensitiveColumnsRequest.SensitiveDataModelId = &sensitiveDataModelId

		listSensitiveColumnsResponse, err := dataSafeClient.ListSensitiveColumns(context.Background(), listSensitiveColumnsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting SensitiveDataModelsSensitiveColumn list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, sensitiveDataModelsSensitiveColumn := range listSensitiveColumnsResponse.Items {
			id := *sensitiveDataModelsSensitiveColumn.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SensitiveDataModelsSensitiveColumnId", id)
		}

	}
	return resourceIds, nil
}
