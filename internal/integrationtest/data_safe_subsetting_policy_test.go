// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

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
	DataSafeSubsettingPolicyRequiredOnlyResource = DataSafeSubsettingPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, DataSafeSubsettingPolicyRepresentation)

	DataSafeSubsettingPolicyResourceConfig = DataSafeSubsettingPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Optional, acctest.Update, DataSafeSubsettingPolicyRepresentation)

	DataSafeSubsettingPolicySingularDataSourceRepresentation = map[string]interface{}{
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
	}

	DataSafeSubsettingPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                          acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"masking_policy_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${var.masking_policy_id}`},
		"state":                                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"subsetting_policy_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"target_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSubsettingPolicyDataSourceFilterRepresentation}}
	DataSafeSubsettingPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`}},
	}

	DataSafeSubsettingPolicyRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"schema_source":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSubsettingPolicySchemaSourceRepresentation},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"is_redo_logging_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_refresh_stats_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"masking_policy_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.masking_policy_id}`},
		"parallel_degree":          acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `2`},
		"post_subsetting_script":   acctest.Representation{RepType: acctest.Optional, Create: `postSubsettingScript`, Update: `postSubsettingScript2`},
		"pre_subsetting_script":    acctest.Representation{RepType: acctest.Optional, Create: `preSubsettingScript`, Update: `preSubsettingScript2`},
		"recompile":                acctest.Representation{RepType: acctest.Optional, Create: `SERIAL`, Update: `PARALLEL`},
		"unrelated_tables_action":  acctest.Representation{RepType: acctest.Optional, Create: `TRUNCATE`, Update: `KEEP`},
	}
	DataSafeSubsettingPolicySchemaSourceRepresentation = map[string]interface{}{
		"schema_source":          acctest.Representation{RepType: acctest.Required, Create: `TARGET`, Update: `TARGET`},
		"schemas_for_subsetting": acctest.Representation{RepType: acctest.Required, Create: []string{`HR_TEST`}, Update: []string{`HR_TEST`}},
		"target_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
	}

	DataSafeSubsettingPolicyResourceDependencies = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	maskingPolicyId := utils.GetEnvSettingWithBlankDefault("data_safe_masking_policy_id")
	maskingPolicyIdVariableStr := fmt.Sprintf("variable \"masking_policy_id\" { default = \"%s\" }\n", maskingPolicyId)

	resourceName := "oci_data_safe_subsetting_policy.test_subsetting_policy"
	datasourceName := "data.oci_data_safe_subsetting_policies.test_subsetting_policies"
	singularDatasourceName := "data.oci_data_safe_subsetting_policy.test_subsetting_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+targetIdVariableStr+maskingPolicyIdVariableStr+DataSafeSubsettingPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Optional, acctest.Create, DataSafeSubsettingPolicyRepresentation), "datasafe", "subsettingPolicy", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSubsettingPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + maskingPolicyIdVariableStr + DataSafeSubsettingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, DataSafeSubsettingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "schema_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schema_source.0.schema_source", "TARGET"),
				resource.TestCheckResourceAttrSet(resourceName, "schema_source.0.target_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + maskingPolicyIdVariableStr + DataSafeSubsettingPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + maskingPolicyIdVariableStr + DataSafeSubsettingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Optional, acctest.Create, DataSafeSubsettingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_redo_logging_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_refresh_stats_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "masking_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "parallel_degree", "2"),
				resource.TestCheckResourceAttr(resourceName, "post_subsetting_script", "postSubsettingScript"),
				resource.TestCheckResourceAttr(resourceName, "pre_subsetting_script", "preSubsettingScript"),
				resource.TestCheckResourceAttr(resourceName, "recompile", "SERIAL"),
				resource.TestCheckResourceAttr(resourceName, "schema_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schema_source.0.schema_source", "TARGET"),
				resource.TestCheckResourceAttr(resourceName, "schema_source.0.schemas_for_subsetting.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "schema_source.0.target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "unrelated_tables_action", "TRUNCATE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						resourceId := "oci_data_safe_subsetting_policy:" + resId
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resourceId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + targetIdVariableStr + maskingPolicyIdVariableStr + DataSafeSubsettingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeSubsettingPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_redo_logging_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_refresh_stats_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "masking_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "parallel_degree", "2"),
				resource.TestCheckResourceAttr(resourceName, "post_subsetting_script", "postSubsettingScript"),
				resource.TestCheckResourceAttr(resourceName, "pre_subsetting_script", "preSubsettingScript"),
				resource.TestCheckResourceAttr(resourceName, "recompile", "SERIAL"),
				resource.TestCheckResourceAttr(resourceName, "schema_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schema_source.0.schema_source", "TARGET"),
				resource.TestCheckResourceAttr(resourceName, "schema_source.0.schemas_for_subsetting.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "schema_source.0.target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "unrelated_tables_action", "TRUNCATE"),

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
			Config: config + compartmentIdVariableStr + targetIdVariableStr + maskingPolicyIdVariableStr + DataSafeSubsettingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Optional, acctest.Update, DataSafeSubsettingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_redo_logging_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_refresh_stats_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "masking_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "parallel_degree", "2"),
				resource.TestCheckResourceAttr(resourceName, "post_subsetting_script", "postSubsettingScript2"),
				resource.TestCheckResourceAttr(resourceName, "pre_subsetting_script", "preSubsettingScript2"),
				resource.TestCheckResourceAttr(resourceName, "recompile", "PARALLEL"),
				resource.TestCheckResourceAttr(resourceName, "schema_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schema_source.0.schema_source", "TARGET"),
				resource.TestCheckResourceAttr(resourceName, "schema_source.0.schemas_for_subsetting.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "schema_source.0.target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "unrelated_tables_action", "KEEP"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policies", "test_subsetting_policies", acctest.Optional, acctest.Update, DataSafeSubsettingPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + maskingPolicyIdVariableStr + DataSafeSubsettingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Optional, acctest.Update, DataSafeSubsettingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "masking_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "subsetting_policy_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subsetting_policy_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, DataSafeSubsettingPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + maskingPolicyIdVariableStr + DataSafeSubsettingPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subsetting_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_redo_logging_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_refresh_stats_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parallel_degree", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "post_subsetting_script", "postSubsettingScript2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pre_subsetting_script", "preSubsettingScript2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recompile", "PARALLEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema_source.0.schema_source", "TARGET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema_source.0.schemas_for_subsetting.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schema_source.0.target_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "unrelated_tables_action", "KEEP"),
			),
		},
		// verify resource import
		{
			Config:                  config + targetIdVariableStr + maskingPolicyIdVariableStr + DataSafeSubsettingPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSubsettingPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_subsetting_policy" {
			noResourceFound = false
			request := oci_data_safe.GetSubsettingPolicyRequest{}

			tmp := rs.Primary.ID
			request.SubsettingPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSubsettingPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.SubsettingPolicyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeSubsettingPolicy") {
		resource.AddTestSweepers("DataSafeSubsettingPolicy", &resource.Sweeper{
			Name:         "DataSafeSubsettingPolicy",
			Dependencies: acctest.DependencyGraph["subsettingPolicy"],
			F:            sweepDataSafeSubsettingPolicyResource,
		})
	}
}

func sweepDataSafeSubsettingPolicyResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	subsettingPolicyIds, err := getDataSafeSubsettingPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, subsettingPolicyId := range subsettingPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[subsettingPolicyId]; !ok {
			deleteSubsettingPolicyRequest := oci_data_safe.DeleteSubsettingPolicyRequest{}

			deleteSubsettingPolicyRequest.SubsettingPolicyId = &subsettingPolicyId

			deleteSubsettingPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSubsettingPolicy(context.Background(), deleteSubsettingPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting SubsettingPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", subsettingPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &subsettingPolicyId, DataSafeSubsettingPolicySweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSubsettingPolicySweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSubsettingPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubsettingPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSubsettingPoliciesRequest := oci_data_safe.ListSubsettingPoliciesRequest{}
	listSubsettingPoliciesRequest.CompartmentId = &compartmentId
	listSubsettingPoliciesRequest.LifecycleState = oci_data_safe.ListSubsettingPoliciesLifecycleStateActive
	listSubsettingPoliciesResponse, err := dataSafeClient.ListSubsettingPolicies(context.Background(), listSubsettingPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SubsettingPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, subsettingPolicy := range listSubsettingPoliciesResponse.Items {
		id := *subsettingPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubsettingPolicyId", id)
	}
	return resourceIds, nil
}

func DataSafeSubsettingPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if subsettingPolicyResponse, ok := response.Response.(oci_data_safe.GetSubsettingPolicyResponse); ok {
		return subsettingPolicyResponse.LifecycleState != oci_data_safe.SubsettingPolicyLifecycleStateDeleted
	}
	return false
}

func DataSafeSubsettingPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSubsettingPolicy(context.Background(), oci_data_safe.GetSubsettingPolicyRequest{
		SubsettingPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
