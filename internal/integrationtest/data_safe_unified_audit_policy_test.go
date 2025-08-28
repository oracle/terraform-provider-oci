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
	DataSafeUnifiedAuditPolicyRequiredOnlyResource = DataSafeUnifiedAuditPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy", "test_unified_audit_policy", acctest.Required, acctest.Create, DataSafeUnifiedAuditPolicyRepresentation)

	DataSafeUnifiedAuditPolicyResourceConfig = DataSafeUnifiedAuditPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy", "test_unified_audit_policy", acctest.Optional, acctest.Update, DataSafeUnifiedAuditPolicyRepresentation)

	DataSafeUnifiedAuditPolicySingularDataSourceRepresentation = map[string]interface{}{
		"unified_audit_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_unified_audit_policy.test_unified_audit_policy.id}`},
	}

	DataSafeUnifiedAuditPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                          acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree":             acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"is_seeded":                             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"security_policy_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.security_policy_id}`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2025-06-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"unified_audit_policy_definition_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.unified_audit_policy_definition_id}`},
		"unified_audit_policy_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_unified_audit_policy.test_unified_audit_policy.id}`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeUnifiedAuditPolicyDataSourceFilterRepresentation}}
	DataSafeUnifiedAuditPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_unified_audit_policy.test_unified_audit_policy.id}`}},
	}

	DataSafeUnifiedAuditPolicyRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"conditions":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeUnifiedAuditPolicyConditionsRepresentation},
		"security_policy_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_id}`},
		"status":                             acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `DISABLED`},
		"unified_audit_policy_definition_id": acctest.Representation{RepType: acctest.Required, Create: `${var.unified_audit_policy_definition_id}`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreUnifiedAuditPolicySystemTagsChangesRep},
	}

	DataSafeUnifiedAuditPolicyConditionsRepresentation = map[string]interface{}{
		"entity_selection": acctest.Representation{RepType: acctest.Required, Create: `INCLUDE`, Update: `EXCLUDE`},
		"entity_type":      acctest.Representation{RepType: acctest.Required, Create: `USER`, Update: `USER`},
		"operation_status": acctest.Representation{RepType: acctest.Required, Create: `SUCCESS`, Update: `FAILURE`},
		"user_names":       acctest.Representation{RepType: acctest.Required, Create: []string{`userNames`}, Update: []string{`userNames2`}},
	}

	DataSafeUnifiedAuditPolicyResourceDependencies = DefinedTagsDependencies

	ignoreUnifiedAuditPolicySystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeUnifiedAuditPolicyResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the security policy and unified audit policy definition are hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeUnifiedAuditPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	securityPolicyId := utils.GetEnvSettingWithBlankDefault("security_policy_id")
	securityPolicyIdVariableStr := fmt.Sprintf("variable \"security_policy_id\" { default = \"%s\" }\n", securityPolicyId)

	unifiedAuditPolicyDefinitionId := utils.GetEnvSettingWithBlankDefault("unified_audit_policy_definition_id")
	unifiedAuditPolicyDefinitionIdVariableStr := fmt.Sprintf("variable \"unified_audit_policy_definition_id\" { default = \"%s\" }\n", unifiedAuditPolicyDefinitionId)

	resourceName := "oci_data_safe_unified_audit_policy.test_unified_audit_policy"
	datasourceName := "data.oci_data_safe_unified_audit_policies.test_unified_audit_policies"
	singularDatasourceName := "data.oci_data_safe_unified_audit_policy.test_unified_audit_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeUnifiedAuditPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy", "test_unified_audit_policy", acctest.Optional, acctest.Create, DataSafeUnifiedAuditPolicyRepresentation), "datasafe", "unifiedAuditPolicy", t)

	acctest.ResourceTest(t, testAccCheckDataSafeUnifiedAuditPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeUnifiedAuditPolicyResourceDependencies + securityPolicyIdVariableStr + unifiedAuditPolicyDefinitionIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy", "test_unified_audit_policy", acctest.Required, acctest.Create, DataSafeUnifiedAuditPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.entity_selection", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.entity_type", "USER"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.operation_status", "SUCCESS"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.user_names.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "unified_audit_policy_definition_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeUnifiedAuditPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeUnifiedAuditPolicyResourceDependencies + securityPolicyIdVariableStr + unifiedAuditPolicyDefinitionIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy", "test_unified_audit_policy", acctest.Optional, acctest.Create, DataSafeUnifiedAuditPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.entity_selection", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.entity_type", "USER"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.operation_status", "SUCCESS"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.user_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "unified_audit_policy_definition_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + securityPolicyIdVariableStr + unifiedAuditPolicyDefinitionIdVariableStr + DataSafeUnifiedAuditPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy", "test_unified_audit_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeUnifiedAuditPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.entity_selection", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.entity_type", "USER"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.operation_status", "SUCCESS"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.user_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "unified_audit_policy_definition_id"),

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
			Config: config + compartmentIdVariableStr + securityPolicyIdVariableStr + unifiedAuditPolicyDefinitionIdVariableStr + DataSafeUnifiedAuditPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy", "test_unified_audit_policy", acctest.Optional, acctest.Update, DataSafeUnifiedAuditPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.entity_selection", "EXCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.entity_type", "USER"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.operation_status", "FAILURE"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.user_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "unified_audit_policy_definition_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_unified_audit_policies", "test_unified_audit_policies", acctest.Optional, acctest.Update, DataSafeUnifiedAuditPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyIdVariableStr + unifiedAuditPolicyDefinitionIdVariableStr + DataSafeUnifiedAuditPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unified_audit_policy", "test_unified_audit_policy", acctest.Optional, acctest.Update, DataSafeUnifiedAuditPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "is_seeded", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "unified_audit_policy_definition_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "unified_audit_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "unified_audit_policy_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "unified_audit_policy_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_unified_audit_policy", "test_unified_audit_policy", acctest.Required, acctest.Create, DataSafeUnifiedAuditPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyIdVariableStr + unifiedAuditPolicyDefinitionIdVariableStr + DataSafeUnifiedAuditPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "unified_audit_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.entity_selection", "EXCLUDE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.entity_type", "USER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.operation_status", "FAILURE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.user_names.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "enabled_entities"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_seeded"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeUnifiedAuditPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeUnifiedAuditPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_unified_audit_policy" {
			noResourceFound = false
			request := oci_data_safe.GetUnifiedAuditPolicyRequest{}

			tmp := rs.Primary.ID
			request.UnifiedAuditPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetUnifiedAuditPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.UnifiedAuditPolicyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeUnifiedAuditPolicy") {
		resource.AddTestSweepers("DataSafeUnifiedAuditPolicy", &resource.Sweeper{
			Name:         "DataSafeUnifiedAuditPolicy",
			Dependencies: acctest.DependencyGraph["unifiedAuditPolicy"],
			F:            sweepDataSafeUnifiedAuditPolicyResource,
		})
	}
}

func sweepDataSafeUnifiedAuditPolicyResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	unifiedAuditPolicyIds, err := getDataSafeUnifiedAuditPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, unifiedAuditPolicyId := range unifiedAuditPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[unifiedAuditPolicyId]; !ok {
			deleteUnifiedAuditPolicyRequest := oci_data_safe.DeleteUnifiedAuditPolicyRequest{}

			deleteUnifiedAuditPolicyRequest.UnifiedAuditPolicyId = &unifiedAuditPolicyId

			deleteUnifiedAuditPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteUnifiedAuditPolicy(context.Background(), deleteUnifiedAuditPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting UnifiedAuditPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", unifiedAuditPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &unifiedAuditPolicyId, DataSafeUnifiedAuditPolicySweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeUnifiedAuditPolicySweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeUnifiedAuditPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "UnifiedAuditPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listUnifiedAuditPoliciesRequest := oci_data_safe.ListUnifiedAuditPoliciesRequest{}
	listUnifiedAuditPoliciesRequest.CompartmentId = &compartmentId
	listUnifiedAuditPoliciesRequest.LifecycleState = oci_data_safe.ListUnifiedAuditPoliciesLifecycleStateActive
	listUnifiedAuditPoliciesResponse, err := dataSafeClient.ListUnifiedAuditPolicies(context.Background(), listUnifiedAuditPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting UnifiedAuditPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, unifiedAuditPolicy := range listUnifiedAuditPoliciesResponse.Items {
		id := *unifiedAuditPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "UnifiedAuditPolicyId", id)
	}
	return resourceIds, nil
}

func DataSafeUnifiedAuditPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if unifiedAuditPolicyResponse, ok := response.Response.(oci_data_safe.GetUnifiedAuditPolicyResponse); ok {
		return unifiedAuditPolicyResponse.LifecycleState != oci_data_safe.UnifiedAuditPolicyLifecycleStateDeleted
	}
	return false
}

func DataSafeUnifiedAuditPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetUnifiedAuditPolicy(context.Background(), oci_data_safe.GetUnifiedAuditPolicyRequest{
		UnifiedAuditPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
