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
	DataSafeSubsettingPolicySubsettingRuleRequiredOnlyResource = DataSafeSubsettingPolicySubsettingRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingRuleRepresentation)

	DataSafeSubsettingPolicySubsettingRuleSingularDataSourceRepresentation = map[string]interface{}{
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"subsetting_rule_key":  acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule.key}`},
	}

	DataSafeSubsettingPolicySubsettingRuleDataSourceRepresentation = map[string]interface{}{
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"schema_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`HR_TEST`}},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSubsettingPolicySubsettingRuleDataSourceFilterRepresentation}}
	DataSafeSubsettingPolicySubsettingRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule.key}`}},
	}

	DataSafeSubsettingPolicySubsettingRuleRepresentation = map[string]interface{}{
		"scope":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSubsettingPolicySubsettingRuleScopeRepresentation},
		"subset_rule_entry":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSubsettingPolicySubsettingRuleSubsetRuleEntryRepresentation},
		"subsetting_policy_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName`},
		"peer_tables_action":         acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"related_tables_propagation": acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"rule_combination_mode":      acctest.Representation{RepType: acctest.Optional, Create: `UNION`, Update: `UNION`},
	}
	DataSafeSubsettingPolicySubsettingRuleScopeRepresentation = map[string]interface{}{
		"schema_name": acctest.Representation{RepType: acctest.Required, Create: `HR_TEST`, Update: `HR_TEST`},
		"scope_type":  acctest.Representation{RepType: acctest.Required, Create: `SPECIFIC`, Update: `SPECIFIC`},
		"object":      acctest.Representation{RepType: acctest.Required, Create: `EMPLOYEES`, Update: `EMPLOYEES`},
	}
	DataSafeSubsettingPolicySubsettingRuleSubsetRuleEntryRepresentation = map[string]interface{}{
		"rule_type":           acctest.Representation{RepType: acctest.Required, Create: `PERCENT`, Update: `PERCENT`},
		"condition":           acctest.Representation{RepType: acctest.Optional, Create: `condition`, Update: nil},
		"partitions_list":     acctest.Representation{RepType: acctest.Optional, Create: []string{`partitionsList`}, Update: nil},
		"percent":             acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"sub_partitions_list": acctest.Representation{RepType: acctest.Optional, Create: []string{`subPartitionsList`}, Update: nil},
	}
	DataSafeSubsettingPolicySubsettingConditionRuleRepresentation = map[string]interface{}{
		"scope": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"schema_name": acctest.Representation{RepType: acctest.Required, Create: `HR_TEST`},
			"scope_type":  acctest.Representation{RepType: acctest.Required, Create: `SPECIFIC`},
			"object":      acctest.Representation{RepType: acctest.Required, Create: `EMPLOYEES`},
		}},
		"subset_rule_entry": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"rule_type": acctest.Representation{RepType: acctest.Required, Create: `CONDITION`},
			"condition": acctest.Representation{RepType: acctest.Required, Create: `JOB_ID = 'IT_PROG'`},
		}},
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
	}
	DataSafeSubsettingPolicySubsettingPartitionRuleRepresentation = map[string]interface{}{
		"scope": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"schema_name": acctest.Representation{RepType: acctest.Required, Create: `HR_TEST`},
			"scope_type":  acctest.Representation{RepType: acctest.Required, Create: `SPECIFIC`},
			"object":      acctest.Representation{RepType: acctest.Required, Create: `PARTITION_TEST`},
		}},
		"subset_rule_entry": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"rule_type":           acctest.Representation{RepType: acctest.Required, Create: `PARTITION`},
			"sub_partitions_list": acctest.Representation{RepType: acctest.Required, Create: []string{`P2025.P2025_HR`, `P2026.P2026_HR`, `P2026.P2026_IT`}},
		}},
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
	}

	DataSafeSubsettingPolicySubsettingRuleResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, DataSafeSubsettingPolicyRepresentation)
	DataSafeSubsettingPolicySubsettingConditionRuleResource    = strings.TrimSuffix(acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_condition_subsetting_rule", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingConditionRuleRepresentation), "}\n") +
		"depends_on = [\"oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule\"]\n}\n"
	DataSafeSubsettingPolicySubsettingPartitionRuleResource = strings.TrimSuffix(acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_partition_subsetting_rule", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingPartitionRuleRepresentation), "}\n") +
		"depends_on = [\"oci_data_safe_subsetting_policy_subsetting_rule.test_condition_subsetting_rule\"]\n}\n"
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingPolicySubsettingRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingPolicySubsettingRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule"
	datasourceName := "data.oci_data_safe_subsetting_policy_subsetting_rules.test_subsetting_policy_subsetting_rules"
	singularDatasourceName := "data.oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule"

	var resId, resId2 string
	ruleCreateRepresentation := acctest.RepresentationCopyWithRemovedNestedProperties("subset_rule_entry.condition", DataSafeSubsettingPolicySubsettingRuleRepresentation)
	ruleCreateRepresentation = acctest.RepresentationCopyWithRemovedNestedProperties("subset_rule_entry.partitions_list", ruleCreateRepresentation)
	ruleCreateRepresentation = acctest.RepresentationCopyWithRemovedNestedProperties("subset_rule_entry.sub_partitions_list", ruleCreateRepresentation)
	ruleUpdateRepresentation := acctest.RepresentationCopyWithRemovedNestedProperties("subset_rule_entry.condition", DataSafeSubsettingPolicySubsettingRuleRepresentation)
	ruleUpdateRepresentation = acctest.RepresentationCopyWithRemovedNestedProperties("subset_rule_entry.partitions_list", ruleUpdateRepresentation)
	ruleUpdateRepresentation = acctest.RepresentationCopyWithRemovedNestedProperties("subset_rule_entry.sub_partitions_list", ruleUpdateRepresentation)
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+targetIdVariableStr+DataSafeSubsettingPolicySubsettingRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Optional, acctest.Create, ruleCreateRepresentation)+
		DataSafeSubsettingPolicySubsettingConditionRuleResource+
		DataSafeSubsettingPolicySubsettingPartitionRuleResource, "datasafe", "subsettingPolicySubsettingRule", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSubsettingPolicySubsettingRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "scope.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scope.0.schema_name", "HR_TEST"),
				resource.TestCheckResourceAttr(resourceName, "scope.0.scope_type", "SPECIFIC"),
				resource.TestCheckResourceAttr(resourceName, "scope.0.object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "subset_rule_entry.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subset_rule_entry.0.percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "subset_rule_entry.0.rule_type", "PERCENT"),
				resource.TestCheckResourceAttrSet(resourceName, "subsetting_policy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Optional, acctest.Create, ruleCreateRepresentation) +
				DataSafeSubsettingPolicySubsettingConditionRuleResource +
				DataSafeSubsettingPolicySubsettingPartitionRuleResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "rule_combination_mode", "UNION"),
				resource.TestCheckResourceAttr(resourceName, "scope.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scope.0.schema_name", "HR_TEST"),
				resource.TestCheckResourceAttr(resourceName, "scope.0.scope_type", "SPECIFIC"),
				resource.TestCheckResourceAttr(resourceName, "scope.0.object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "subset_rule_entry.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subset_rule_entry.0.percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "subset_rule_entry.0.rule_type", "PERCENT"),
				resource.TestCheckResourceAttrSet(resourceName, "subsetting_policy_id"),
				resource.TestCheckResourceAttr("oci_data_safe_subsetting_policy_subsetting_rule.test_condition_subsetting_rule", "scope.0.schema_name", "HR_TEST"),
				resource.TestCheckResourceAttr("oci_data_safe_subsetting_policy_subsetting_rule.test_condition_subsetting_rule", "scope.0.object", "EMPLOYEES"),
				resource.TestCheckResourceAttr("oci_data_safe_subsetting_policy_subsetting_rule.test_condition_subsetting_rule", "subset_rule_entry.0.rule_type", "CONDITION"),
				resource.TestCheckResourceAttr("oci_data_safe_subsetting_policy_subsetting_rule.test_condition_subsetting_rule", "subset_rule_entry.0.condition", "JOB_ID = 'IT_PROG'"),
				resource.TestCheckResourceAttr("oci_data_safe_subsetting_policy_subsetting_rule.test_partition_subsetting_rule", "scope.0.schema_name", "HR_TEST"),
				resource.TestCheckResourceAttr("oci_data_safe_subsetting_policy_subsetting_rule.test_partition_subsetting_rule", "scope.0.object", "PARTITION_TEST"),
				resource.TestCheckResourceAttr("oci_data_safe_subsetting_policy_subsetting_rule.test_partition_subsetting_rule", "subset_rule_entry.0.rule_type", "PARTITION"),
				resource.TestCheckResourceAttr("oci_data_safe_subsetting_policy_subsetting_rule.test_partition_subsetting_rule", "subset_rule_entry.0.sub_partitions_list.#", "3"),
				resource.TestCheckResourceAttr("oci_data_safe_subsetting_policy_subsetting_rule.test_partition_subsetting_rule", "subset_rule_entry.0.sub_partitions_list.0", "P2025.P2025_HR"),
				resource.TestCheckResourceAttr("oci_data_safe_subsetting_policy_subsetting_rule.test_partition_subsetting_rule", "subset_rule_entry.0.sub_partitions_list.1", "P2026.P2026_HR"),
				resource.TestCheckResourceAttr("oci_data_safe_subsetting_policy_subsetting_rule.test_partition_subsetting_rule", "subset_rule_entry.0.sub_partitions_list.2", "P2026.P2026_IT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					fullPath := "oci_data_safe_subsetting_policy_subsetting_rule:" + resId
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
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Optional, acctest.Update, ruleUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "rule_combination_mode", "UNION"),
				resource.TestCheckResourceAttr(resourceName, "scope.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scope.0.schema_name", "HR_TEST"),
				resource.TestCheckResourceAttr(resourceName, "scope.0.scope_type", "SPECIFIC"),
				resource.TestCheckResourceAttr(resourceName, "scope.0.object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "subset_rule_entry.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subset_rule_entry.0.percent", "11"),
				resource.TestCheckResourceAttr(resourceName, "subset_rule_entry.0.rule_type", "PERCENT"),
				resource.TestCheckResourceAttrSet(resourceName, "subsetting_policy_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rules", "test_subsetting_policy_subsetting_rules", acctest.Optional, acctest.Update, DataSafeSubsettingPolicySubsettingRuleDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Optional, acctest.Update, ruleUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "schema_name.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "subsetting_rule_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subsetting_rule_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Optional, acctest.Update, ruleUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subsetting_policy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subsetting_rule_key"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rule_combination_mode", "UNION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scope.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scope.0.schema_name", "HR_TEST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scope.0.scope_type", "SPECIFIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scope.0.object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subset_rule_entry.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subset_rule_entry.0.percent", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subset_rule_entry.0.rule_type", "PERCENT"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeSubsettingPolicySubsettingRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSubsettingPolicySubsettingRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_subsetting_policy_subsetting_rule" {
			noResourceFound = false
			request := oci_data_safe.GetSubsettingRuleRequest{}

			if value, ok := rs.Primary.Attributes["subsetting_policy_id"]; ok {
				request.SubsettingPolicyId = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.SubsettingRuleKey = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetSubsettingRule(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataSafeSubsettingPolicySubsettingRule") {
		resource.AddTestSweepers("DataSafeSubsettingPolicySubsettingRule", &resource.Sweeper{
			Name:         "DataSafeSubsettingPolicySubsettingRule",
			Dependencies: acctest.DependencyGraph["subsettingPolicySubsettingRule"],
			F:            sweepDataSafeSubsettingPolicySubsettingRuleResource,
		})
	}
}

func sweepDataSafeSubsettingPolicySubsettingRuleResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	subsettingPolicySubsettingRuleIds, err := getDataSafeSubsettingPolicySubsettingRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, subsettingPolicySubsettingRuleId := range subsettingPolicySubsettingRuleIds {
		if ok := acctest.SweeperDefaultResourceId[subsettingPolicySubsettingRuleId]; !ok {
			parts := strings.Split(subsettingPolicySubsettingRuleId, "/")
			if len(parts) != 4 {
				fmt.Printf("Skipping SubsettingPolicySubsettingRule with invalid ID %s\n", subsettingPolicySubsettingRuleId)
				continue
			}
			deleteSubsettingRuleRequest := oci_data_safe.DeleteSubsettingRuleRequest{}
			deleteSubsettingRuleRequest.SubsettingPolicyId = &parts[1]
			deleteSubsettingRuleRequest.SubsettingRuleKey = &parts[3]

			deleteSubsettingRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSubsettingRule(context.Background(), deleteSubsettingRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting SubsettingPolicySubsettingRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", subsettingPolicySubsettingRuleId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeSubsettingPolicySubsettingRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubsettingPolicySubsettingRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSubsettingRulesRequest := oci_data_safe.ListSubsettingRulesRequest{}

	subsettingPolicyIds, error := getDataSafeSubsettingPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting subsettingPolicyId required for SubsettingPolicySubsettingRule resource requests \n")
	}
	for _, subsettingPolicyId := range subsettingPolicyIds {
		listSubsettingRulesRequest.SubsettingPolicyId = &subsettingPolicyId

		listSubsettingRulesResponse, err := dataSafeClient.ListSubsettingRules(context.Background(), listSubsettingRulesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting SubsettingPolicySubsettingRule list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, subsettingPolicySubsettingRule := range listSubsettingRulesResponse.Items {
			id := fmt.Sprintf("subsettingPolicies/%s/subsettingRules/%s", subsettingPolicyId, *subsettingPolicySubsettingRule.Key)
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubsettingPolicySubsettingRuleId", id)
		}

	}
	return resourceIds, nil
}
