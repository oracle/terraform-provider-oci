// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectRequiredOnlyResource = DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object", "test_subsetting_policy_subsetting_rule_processing_chain_object", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectRepresentation)

	DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectDataSourceRepresentation = map[string]interface{}{
		"subsetting_policy_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"subsetting_rule_key":       acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule.key}`},
		"is_enabled_for_processing": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectDataSourceFilterRepresentation},
	}
	DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `propagation_impact`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`PARENT_CHILD_SUBSET`}},
	}

	DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectRepresentation = map[string]interface{}{
		"processing_chain_object_key": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_objects.test_subsetting_policy_subsetting_rule_processing_chain_objects.subsetting_rule_processing_chain_objects_collection.0.items.0.key}`},
		"subsetting_policy_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"subsetting_rule_key":         acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule.key}`},
		"is_enabled_for_processing":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}

	DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectRuleRepresentation = map[string]interface{}{
		"scope":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectScopeRepresentation},
		"subset_rule_entry":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSubsettingPolicySubsettingRuleSubsetRuleEntryRepresentation},
		"subsetting_policy_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"peer_tables_action":         acctest.Representation{RepType: acctest.Required, Create: `MINIMUM_ROWS`},
		"related_tables_propagation": acctest.Representation{RepType: acctest.Required, Create: `DESCENDANTS`},
	}
	DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectScopeRepresentation = map[string]interface{}{
		"schema_name": acctest.Representation{RepType: acctest.Required, Create: `HR_TEST`},
		"scope_type":  acctest.Representation{RepType: acctest.Required, Create: `SPECIFIC`},
		"object":      acctest.Representation{RepType: acctest.Required, Create: `EMPLOYEES`},
	}

	DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, DataSafeSubsettingPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectRuleRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_objects", "test_subsetting_policy_subsetting_rule_processing_chain_objects", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object.test_subsetting_policy_subsetting_rule_processing_chain_object"
	datasourceName := "data.oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_objects.test_subsetting_policy_subsetting_rule_processing_chain_objects"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+targetIdVariableStr+DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object", "test_subsetting_policy_subsetting_rule_processing_chain_object", acctest.Optional, acctest.Create, DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectRepresentation), "datasafe", "subsettingPolicySubsettingRuleProcessingChainObject", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object", "test_subsetting_policy_subsetting_rule_processing_chain_object", acctest.Optional, acctest.Create, DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_enabled_for_processing", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "processing_chain_object_key"),
				resource.TestCheckResourceAttrSet(resourceName, "subsetting_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subsetting_rule_key"),

				func(s *terraform.State) (err error) {
					resId, stateErr := acctest.FromInstanceState(s, resourceName, "id")
					if stateErr != nil {
						return stateErr
					}
					fullPath := "oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object:" + resId
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
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object", "test_subsetting_policy_subsetting_rule_processing_chain_object", acctest.Optional, acctest.Update, DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_enabled_for_processing", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "processing_chain_object_key"),
				resource.TestCheckResourceAttrSet(resourceName, "subsetting_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subsetting_rule_key"),
			),
		},
		// verify datasource
		{
			Config: config +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object", "test_subsetting_policy_subsetting_rule_processing_chain_object", acctest.Optional, acctest.Update, DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_rule_key"),

				resource.TestCheckResourceAttr(datasourceName, "subsetting_rule_processing_chain_objects_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subsetting_rule_processing_chain_objects_collection.0.items.0.is_enabled_for_processing", "true"),
			),
		},
		// verify resource import
		{
			Config:            config + DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_enabled_for_processing",
			},
			ResourceName: resourceName,
		},
	})
}
