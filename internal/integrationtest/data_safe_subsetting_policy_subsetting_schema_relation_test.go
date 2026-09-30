// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
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
	DataSafeSubsettingPolicySubsettingSchemaRelationRequiredOnlyResource = DataSafeSubsettingPolicySubsettingSchemaRelationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_schema_relation", "test_subsetting_policy_subsetting_schema_relation", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingSchemaRelationRepresentation)

	DataSafeSubsettingPolicySubsettingSchemaRelationResourceConfig = DataSafeSubsettingPolicySubsettingSchemaRelationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_schema_relation", "test_subsetting_policy_subsetting_schema_relation", acctest.Optional, acctest.Update, DataSafeSubsettingPolicySubsettingSchemaRelationRepresentation)

	DataSafeSubsettingPolicySubsettingSchemaRelationSingularDataSourceRepresentation = map[string]interface{}{
		"subsetting_policy_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"subsetting_schema_relation_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy_subsetting_schema_relation.test_subsetting_policy_subsetting_schema_relation.key}`},
	}

	DataSafeSubsettingPolicySubsettingSchemaRelationDataSourceRepresentation = map[string]interface{}{
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"object":               acctest.Representation{RepType: acctest.Optional, Create: []string{`PROJECTS`}},
		"relation_type":        acctest.Representation{RepType: acctest.Optional, Create: `APP_DEFINED`},
		"schema_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`HR_TEST`}},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSubsettingPolicySubsettingSchemaRelationDataSourceFilterRepresentation}}
	DataSafeSubsettingPolicySubsettingSchemaRelationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_subsetting_policy_subsetting_schema_relation.test_subsetting_policy_subsetting_schema_relation.key}`}},
	}

	DataSafeSubsettingPolicySubsettingSchemaRelationRepresentation = map[string]interface{}{
		"child_columns":        acctest.Representation{RepType: acctest.Required, Create: []string{`START_DATE`}},
		"child_object_name":    acctest.Representation{RepType: acctest.Required, Create: `PROJECTS`},
		"child_schema_name":    acctest.Representation{RepType: acctest.Required, Create: `HR_TEST`},
		"parent_columns":       acctest.Representation{RepType: acctest.Required, Create: []string{`HIRE_DATE`}},
		"parent_object_name":   acctest.Representation{RepType: acctest.Required, Create: `EMPLOYEES`},
		"parent_schema_name":   acctest.Representation{RepType: acctest.Required, Create: `HR_TEST`},
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
	}

	DataSafeSubsettingPolicySubsettingSchemaRelationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, DataSafeSubsettingPolicyRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingPolicySubsettingSchemaRelationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingPolicySubsettingSchemaRelationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_subsetting_policy_subsetting_schema_relation.test_subsetting_policy_subsetting_schema_relation"
	datasourceName := "data.oci_data_safe_subsetting_policy_subsetting_schema_relations.test_subsetting_policy_subsetting_schema_relations"
	singularDatasourceName := "data.oci_data_safe_subsetting_policy_subsetting_schema_relation.test_subsetting_policy_subsetting_schema_relation"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+targetIdVariableStr+DataSafeSubsettingPolicySubsettingSchemaRelationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_schema_relation", "test_subsetting_policy_subsetting_schema_relation", acctest.Optional, acctest.Create, DataSafeSubsettingPolicySubsettingSchemaRelationRepresentation), "datasafe", "subsettingPolicySubsettingSchemaRelation", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSubsettingPolicySubsettingSchemaRelationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingSchemaRelationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_schema_relation", "test_subsetting_policy_subsetting_schema_relation", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingSchemaRelationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "child_columns.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "child_object_name"),
				resource.TestCheckResourceAttrSet(resourceName, "child_object_key"),
				resource.TestCheckResourceAttr(resourceName, "child_schema_name", "HR_TEST"),
				resource.TestCheckResourceAttr(resourceName, "parent_columns.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_object_name"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_object_key"),
				resource.TestCheckResourceAttr(resourceName, "parent_schema_name", "HR_TEST"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttrSet(resourceName, "relation_type"),
				resource.TestCheckResourceAttrSet(resourceName, "subsetting_policy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					fullPath := "oci_data_safe_subsetting_policy_subsetting_schema_relation:" + resId
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&fullPath, &compartmentId, resourceName); errExport != nil {
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_schema_relations", "test_subsetting_policy_subsetting_schema_relations", acctest.Optional, acctest.Update, DataSafeSubsettingPolicySubsettingSchemaRelationDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingSchemaRelationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_schema_relation", "test_subsetting_policy_subsetting_schema_relation", acctest.Optional, acctest.Update, DataSafeSubsettingPolicySubsettingSchemaRelationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "object.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "relation_type", "APP_DEFINED"),
				resource.TestCheckResourceAttr(datasourceName, "schema_name.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "subsetting_schema_relation_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subsetting_schema_relation_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_schema_relation", "test_subsetting_policy_subsetting_schema_relation", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingSchemaRelationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingSchemaRelationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subsetting_policy_id"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "subsetting_schema_relation_key", resourceName, "key"),

				resource.TestCheckResourceAttr(singularDatasourceName, "child_columns.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "child_object_key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "child_schema_name", "HR_TEST"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parent_columns.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_object_key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parent_schema_name", "HR_TEST"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "relation_type"),
				// TODO: Re-enable after the service returns timeCreated and timeUpdated for schema relations.
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeSubsettingPolicySubsettingSchemaRelationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSubsettingPolicySubsettingSchemaRelationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_subsetting_policy_subsetting_schema_relation" {
			noResourceFound = false
			request := oci_data_safe.GetSubsettingSchemaRelationRequest{}

			if value, ok := rs.Primary.Attributes["subsetting_policy_id"]; ok {
				request.SubsettingPolicyId = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.SubsettingSchemaRelationKey = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetSubsettingSchemaRelation(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataSafeSubsettingPolicySubsettingSchemaRelation") {
		resource.AddTestSweepers("DataSafeSubsettingPolicySubsettingSchemaRelation", &resource.Sweeper{
			Name:         "DataSafeSubsettingPolicySubsettingSchemaRelation",
			Dependencies: acctest.DependencyGraph["subsettingPolicySubsettingSchemaRelation"],
			F:            sweepDataSafeSubsettingPolicySubsettingSchemaRelationResource,
		})
	}
}

func sweepDataSafeSubsettingPolicySubsettingSchemaRelationResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	subsettingPolicySubsettingSchemaRelationIds, err := getDataSafeSubsettingPolicySubsettingSchemaRelationIds(compartment)
	if err != nil {
		return err
	}
	for _, subsettingPolicySubsettingSchemaRelationId := range subsettingPolicySubsettingSchemaRelationIds {
		if ok := acctest.SweeperDefaultResourceId[subsettingPolicySubsettingSchemaRelationId]; !ok {
			parts := strings.Split(subsettingPolicySubsettingSchemaRelationId, "/")
			if len(parts) != 4 {
				fmt.Printf("Skipping SubsettingPolicySubsettingSchemaRelation with invalid ID %s\n", subsettingPolicySubsettingSchemaRelationId)
				continue
			}
			deleteSubsettingSchemaRelationRequest := oci_data_safe.DeleteSubsettingSchemaRelationRequest{}
			deleteSubsettingSchemaRelationRequest.SubsettingPolicyId = &parts[1]
			deleteSubsettingSchemaRelationRequest.SubsettingSchemaRelationKey = &parts[3]

			deleteSubsettingSchemaRelationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSubsettingSchemaRelation(context.Background(), deleteSubsettingSchemaRelationRequest)
			if error != nil {
				fmt.Printf("Error deleting SubsettingPolicySubsettingSchemaRelation %s %s, It is possible that the resource is already deleted. Please verify manually \n", subsettingPolicySubsettingSchemaRelationId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeSubsettingPolicySubsettingSchemaRelationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubsettingPolicySubsettingSchemaRelationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSubsettingSchemaRelationsRequest := oci_data_safe.ListSubsettingSchemaRelationsRequest{}

	subsettingPolicyIds, error := getDataSafeSubsettingPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting subsettingPolicyId required for SubsettingPolicySubsettingSchemaRelation resource requests \n")
	}
	for _, subsettingPolicyId := range subsettingPolicyIds {
		listSubsettingSchemaRelationsRequest.SubsettingPolicyId = &subsettingPolicyId

		listSubsettingSchemaRelationsResponse, err := dataSafeClient.ListSubsettingSchemaRelations(context.Background(), listSubsettingSchemaRelationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting SubsettingPolicySubsettingSchemaRelation list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, subsettingPolicySubsettingSchemaRelation := range listSubsettingSchemaRelationsResponse.Items {
			id := fmt.Sprintf("subsettingPolicies/%s/subsettingSchemaRelations/%s", subsettingPolicyId, *subsettingPolicySubsettingSchemaRelation.Key)
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubsettingPolicySubsettingSchemaRelationId", id)
		}

	}
	return resourceIds, nil
}
