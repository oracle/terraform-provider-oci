// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
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
	DataSafeMaskingPolicyRequiredOnlyResource = DataSafeMaskingPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Required, acctest.Create, maskingPolicyRepresentation)

	DataSafeMaskingPolicyResourceConfig = DataSafeMaskingPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Optional, acctest.Update, maskingPolicyRepresentation)

	DataSafemaskingPolicySingularDataSourceRepresentation = map[string]interface{}{
		"masking_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_masking_policy.test_masking_policy.id}`},
	}

	DataSafemaskingPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: maskingPolicyDataSourceFilterRepresentation}}
	maskingPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_masking_policy.test_masking_policy.id}`}},
	}

	maskingPolicyRepresentation = map[string]interface{}{
		"column_source":  acctest.RepresentationGroup{RepType: acctest.Required, Group: maskingPolicyColumnSourceRepresentation},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		//"is_drop_temp_tables_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"is_redo_logging_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_refresh_stats_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"parallel_degree":          acctest.Representation{RepType: acctest.Optional, Create: `parallelDegree`, Update: `parallelDegree2`},
		"post_masking_script":      acctest.Representation{RepType: acctest.Optional, Create: `postMaskingScript`, Update: `postMaskingScript2`},
		"pre_masking_script":       acctest.Representation{RepType: acctest.Optional, Create: `preMaskingScript`, Update: `preMaskingScript2`},
		"recompile":                acctest.Representation{RepType: acctest.Optional, Create: `SERIAL`, Update: `SERIAL`},
	}
	maskingPolicyColumnSourceRepresentation = map[string]interface{}{
		"column_source":           acctest.Representation{RepType: acctest.Required, Create: `SENSITIVE_DATA_MODEL`},
		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model1.id}`, Update: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model2.id}`},
	}

	DataSafeMaskingPolicyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model1", acctest.Required, acctest.Create, sensitiveDataModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model2", acctest.Required, acctest.Create, sensitiveDataModelRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeMaskingPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeMaskingPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_masking_policy.test_masking_policy"
	datasourceName := "data.oci_data_safe_masking_policies.test_masking_policies"
	singularDatasourceName := "data.oci_data_safe_masking_policy.test_masking_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeMaskingPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Optional, acctest.Create, maskingPolicyRepresentation), "datasafe", "maskingPolicy", t)

	acctest.ResourceTest(t, testAccCheckDataSafeMaskingPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeMaskingPolicyResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Required, acctest.Create, maskingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_source.0.column_source", "SENSITIVE_DATA_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "column_source.0.sensitive_data_model_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeMaskingPolicyResourceDependencies + targetIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeMaskingPolicyResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Optional, acctest.Create, maskingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_source.0.column_source", "SENSITIVE_DATA_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "column_source.0.sensitive_data_model_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_redo_logging_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_refresh_stats_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "parallel_degree", "parallelDegree"),
				resource.TestCheckResourceAttr(resourceName, "post_masking_script", "postMaskingScript"),
				resource.TestCheckResourceAttr(resourceName, "pre_masking_script", "preMaskingScript"),
				resource.TestCheckResourceAttr(resourceName, "recompile", "SERIAL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeMaskingPolicyResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(maskingPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_source.0.column_source", "SENSITIVE_DATA_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "column_source.0.sensitive_data_model_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_redo_logging_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_refresh_stats_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "parallel_degree", "parallelDegree"),
				resource.TestCheckResourceAttr(resourceName, "post_masking_script", "postMaskingScript"),
				resource.TestCheckResourceAttr(resourceName, "pre_masking_script", "preMaskingScript"),
				resource.TestCheckResourceAttr(resourceName, "recompile", "SERIAL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + DataSafeMaskingPolicyResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Optional, acctest.Update, maskingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_source.0.column_source", "SENSITIVE_DATA_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "column_source.0.sensitive_data_model_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_redo_logging_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_refresh_stats_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "parallel_degree", "parallelDegree2"),
				resource.TestCheckResourceAttr(resourceName, "post_masking_script", "postMaskingScript2"),
				resource.TestCheckResourceAttr(resourceName, "pre_masking_script", "preMaskingScript2"),
				resource.TestCheckResourceAttr(resourceName, "recompile", "SERIAL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_policies", "test_masking_policies", acctest.Optional, acctest.Update, DataSafemaskingPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeMaskingPolicyResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Optional, acctest.Update, maskingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "masking_policy_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "masking_policy_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Required, acctest.Create, DataSafemaskingPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeMaskingPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "masking_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "column_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_source.0.column_source", "SENSITIVE_DATA_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "is_drop_temp_tables_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_redo_logging_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_refresh_stats_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parallel_degree", "parallelDegree2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "post_masking_script", "postMaskingScript2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pre_masking_script", "preMaskingScript2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recompile", "SERIAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DataSafeMaskingPolicyResourceConfig + targetIdVariableStr,
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

func testAccCheckDataSafeMaskingPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_masking_policy" {
			noResourceFound = false
			request := oci_data_safe.GetMaskingPolicyRequest{}

			tmp := rs.Primary.ID
			request.MaskingPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetMaskingPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.MaskingLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeMaskingPolicy") {
		resource.AddTestSweepers("DataSafeMaskingPolicy", &resource.Sweeper{
			Name:         "DataSafeMaskingPolicy",
			Dependencies: acctest.DependencyGraph["maskingPolicy"],
			F:            sweepDataSafeMaskingPolicyResource,
		})
	}
}

func sweepDataSafeMaskingPolicyResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	maskingPolicyIds, err := getDataSafeMaskingPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, maskingPolicyId := range maskingPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[maskingPolicyId]; !ok {
			deleteMaskingPolicyRequest := oci_data_safe.DeleteMaskingPolicyRequest{}

			deleteMaskingPolicyRequest.MaskingPolicyId = &maskingPolicyId

			deleteMaskingPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteMaskingPolicy(context.Background(), deleteMaskingPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting MaskingPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", maskingPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &maskingPolicyId, DataSafemaskingPoliciesSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafemaskingPoliciesSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeMaskingPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MaskingPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listMaskingPoliciesRequest := oci_data_safe.ListMaskingPoliciesRequest{}
	listMaskingPoliciesRequest.CompartmentId = &compartmentId
	listMaskingPoliciesRequest.LifecycleState = oci_data_safe.ListMaskingPoliciesLifecycleStateActive
	listMaskingPoliciesResponse, err := dataSafeClient.ListMaskingPolicies(context.Background(), listMaskingPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MaskingPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, maskingPolicy := range listMaskingPoliciesResponse.Items {
		id := *maskingPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MaskingPolicyId", id)
	}
	return resourceIds, nil
}

func DataSafemaskingPoliciesSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if maskingPolicyResponse, ok := response.Response.(oci_data_safe.GetMaskingPolicyResponse); ok {
		return maskingPolicyResponse.LifecycleState != oci_data_safe.MaskingLifecycleStateDeleted
	}
	return false
}

func DataSafemaskingPoliciesSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetMaskingPolicy(context.Background(), oci_data_safe.GetMaskingPolicyRequest{
		MaskingPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
