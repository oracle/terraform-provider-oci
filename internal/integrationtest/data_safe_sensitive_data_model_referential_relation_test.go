// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSensitiveDataModelReferentialRelationResourceConfig = DataSafeSensitiveDataModelReferentialRelationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model_referential_relation", "test_sensitive_data_model_referential_relation", acctest.Optional, acctest.Update, DataSafeSensitiveDataModelReferentialRelationRepresentation)

	DataSafeSensitiveDataModelReferentialRelationSingularDataSourceRepresentation = map[string]interface{}{
		"key":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_data_model_referential_relation.test_sensitive_data_model_referential_relation.key}`},
		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sensitive_data_model_id}`},
	}

	DataSafeSensitiveDataModelReferentialRelationDataSourceRepresentation = map[string]interface{}{
		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sensitive_data_model_id}`},
		"object":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`EMPLOYEES`}},
		"relation_type":           acctest.Representation{RepType: acctest.Optional, Create: []string{`APP_DEFINED`}},
		"schema_name":             acctest.Representation{RepType: acctest.Optional, Create: []string{`HR`}},
		"column_name":             acctest.Representation{RepType: acctest.Optional, Create: []string{`EMAIL`}},
		"is_sensitive":            acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSensitiveDataModelReferentialRelationDataSourceFilterRepresentation}}
	DataSafeSensitiveDataModelReferentialRelationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_sensitive_data_model_referential_relation.test_sensitive_data_model_referential_relation.key}`}},
	}

	DataSafeSensitiveDataModelReferentialRelationRepresentation = map[string]interface{}{
		"child":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSensitiveDataModelReferentialRelationChildRepresentation},
		"parent":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSensitiveDataModelReferentialRelationParentRepresentation},
		"relation_type":           acctest.Representation{RepType: acctest.Required, Create: `APP_DEFINED`},
		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sensitive_data_model_id}`},
		"is_sensitive":            acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	DataSafeSensitiveDataModelReferentialRelationChildRepresentation = map[string]interface{}{
		"app_name":     acctest.Representation{RepType: acctest.Required, Create: `HR`},
		"column_group": acctest.Representation{RepType: acctest.Required, Create: []string{`FIRST_NAME`}},
		"object":       acctest.Representation{RepType: acctest.Required, Create: `EMPLOYEES`},
		"object_type":  acctest.Representation{RepType: acctest.Required, Create: `TABLE`},
		"schema_name":  acctest.Representation{RepType: acctest.Required, Create: `HR`},
	}
	DataSafeSensitiveDataModelReferentialRelationParentRepresentation = map[string]interface{}{
		"app_name":     acctest.Representation{RepType: acctest.Required, Create: `HR`},
		"column_group": acctest.Representation{RepType: acctest.Required, Create: []string{`EMAIL`}},
		"object":       acctest.Representation{RepType: acctest.Required, Create: `EMPLOYEES`},
		"object_type":  acctest.Representation{RepType: acctest.Required, Create: `TABLE`},
		"schema_name":  acctest.Representation{RepType: acctest.Required, Create: `HR`},
	}

	DataSafeSensitiveDataModelReferentialRelationResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveDataModelReferentialRelationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSensitiveDataModelReferentialRelationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sdmId := utils.GetEnvSettingWithBlankDefault("sensitive_data_model_id")
	sdmIdVariableStr := fmt.Sprintf("variable \"sensitive_data_model_id\" { default = \"%s\" }\n", sdmId)

	resourceName := "oci_data_safe_sensitive_data_model_referential_relation.test_sensitive_data_model_referential_relation"
	datasourceName := "data.oci_data_safe_sensitive_data_model_referential_relations.test_sensitive_data_model_referential_relations"
	singularDatasourceName := "data.oci_data_safe_sensitive_data_model_referential_relation.test_sensitive_data_model_referential_relation"

	var resId string

	// var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeSensitiveDataModelReferentialRelationResourceDependencies+sdmIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model_referential_relation", "test_sensitive_data_model_referential_relation", acctest.Required, acctest.Create, DataSafeSensitiveDataModelReferentialRelationRepresentation), "datasafe", "sensitiveDataModelReferentialRelation", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSensitiveDataModelReferentialRelationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveDataModelReferentialRelationResourceDependencies + sdmIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model_referential_relation", "test_sensitive_data_model_referential_relation", acctest.Optional, acctest.Create, DataSafeSensitiveDataModelReferentialRelationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "child.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "child.0.app_name", "HR"),
				resource.TestCheckResourceAttr(resourceName, "child.0.column_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "child.0.object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "child.0.object_type", "TABLE"),
				resource.TestCheckResourceAttr(resourceName, "child.0.schema_name", "HR"),
				resource.TestCheckResourceAttr(resourceName, "parent.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parent.0.app_name", "HR"),
				resource.TestCheckResourceAttr(resourceName, "parent.0.column_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parent.0.object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "parent.0.object_type", "TABLE"),
				resource.TestCheckResourceAttr(resourceName, "parent.0.schema_name", "HR"),
				resource.TestCheckResourceAttr(resourceName, "relation_type", "APP_DEFINED"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_data_model_id"),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_data_model_referential_relations", "test_sensitive_data_model_referential_relations", acctest.Optional, acctest.Update, DataSafeSensitiveDataModelReferentialRelationDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSensitiveDataModelReferentialRelationResourceDependencies + sdmIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model_referential_relation", "test_sensitive_data_model_referential_relation", acctest.Optional, acctest.Update, DataSafeSensitiveDataModelReferentialRelationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "object.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "relation_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "schema_name.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttr(datasourceName, "is_sensitive", "true"),

				resource.TestCheckResourceAttr(datasourceName, "referential_relation_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "referential_relation_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_data_model_referential_relation", "test_sensitive_data_model_referential_relation", acctest.Required, acctest.Create, DataSafeSensitiveDataModelReferentialRelationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + sdmIdVariableStr + DataSafeSensitiveDataModelReferentialRelationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "referential_relation_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_data_model_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "child.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "child.0.app_name", "HR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "child.0.column_group.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "child.0.object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "child.0.object_type", "TABLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "child.0.schema_name", "HR"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parent.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parent.0.app_name", "HR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parent.0.column_group.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parent.0.object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parent.0.object_type", "TABLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parent.0.schema_name", "HR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "relation_type", "APP_DEFINED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeSensitiveDataModelReferentialRelationResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSensitiveDataModelReferentialRelationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_sensitive_data_model_referential_relation" {
			noResourceFound = false
			request := oci_data_safe.GetReferentialRelationRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.ReferentialRelationKey = &value
			}

			if value, ok := rs.Primary.Attributes["sensitive_data_model_id"]; ok {
				request.SensitiveDataModelId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetReferentialRelation(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataSafeSensitiveDataModelReferentialRelation") {
		resource.AddTestSweepers("DataSafeSensitiveDataModelReferentialRelation", &resource.Sweeper{
			Name:         "DataSafeSensitiveDataModelReferentialRelation",
			Dependencies: acctest.DependencyGraph["sensitiveDataModelReferentialRelation"],
			F:            sweepDataSafeSensitiveDataModelReferentialRelationResource,
		})
	}
}

func sweepDataSafeSensitiveDataModelReferentialRelationResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	sensitiveDataModelReferentialRelationIds, err := getDataSafeSensitiveDataModelReferentialRelationIds(compartment)
	if err != nil {
		return err
	}
	for _, sensitiveDataModelReferentialRelationId := range sensitiveDataModelReferentialRelationIds {
		if ok := acctest.SweeperDefaultResourceId[sensitiveDataModelReferentialRelationId]; !ok {
			deleteReferentialRelationRequest := oci_data_safe.DeleteReferentialRelationRequest{}

			deleteReferentialRelationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteReferentialRelation(context.Background(), deleteReferentialRelationRequest)
			if error != nil {
				fmt.Printf("Error deleting SensitiveDataModelReferentialRelation %s %s, It is possible that the resource is already deleted. Please verify manually \n", sensitiveDataModelReferentialRelationId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeSensitiveDataModelReferentialRelationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SensitiveDataModelReferentialRelationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listReferentialRelationsRequest := oci_data_safe.ListReferentialRelationsRequest{}

	sensitiveDataModelIds, error := getDataSafeSensitiveDataModelIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting sensitiveDataModelId required for SensitiveDataModelReferentialRelation resource requests \n")
	}
	for _, sensitiveDataModelId := range sensitiveDataModelIds {
		listReferentialRelationsRequest.SensitiveDataModelId = &sensitiveDataModelId

		listReferentialRelationsResponse, err := dataSafeClient.ListReferentialRelations(context.Background(), listReferentialRelationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting SensitiveDataModelReferentialRelation list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, sensitiveDataModelReferentialRelation := range listReferentialRelationsResponse.Items {
			id := *sensitiveDataModelReferentialRelation.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SensitiveDataModelReferentialRelationId", id)
		}

	}
	return resourceIds, nil
}
