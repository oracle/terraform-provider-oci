// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreDrgNatPolicyDrgNatRuleResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy", "test_drg_nat_policy", acctest.Required, acctest.Create, CoreDrgNatPolicyRepresentation)

	CoreDrgNatPolicyDrgNatRuleDataSourceRepresentation = map[string]interface{}{
		"drg_nat_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_nat_policy.test_drg_nat_policy.id}`},
	}

	CoreDrgNatPolicyDrgNatRuleRequiredOnlyResource = CoreDrgNatPolicyDrgNatRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy_drg_nat_rule", "test_drg_nat_policy_drg_nat_rule", acctest.Required, acctest.Create, CoreDrgNatPolicyDrgNatRuleRepresentation)

	CoreDrgNatPolicyDrgNatRuleRepresentation = map[string]interface{}{
		"drg_nat_policy_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_nat_policy.test_drg_nat_policy.id}`},
		"drg_nat_rule_priority":  acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `10`},
		"original_source":        acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`, Update: `192.0.0.0/24`},
		"translated_source":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`, Update: `192.0.0.0/24`},
		"original_destination":   acctest.Representation{RepType: acctest.Required, Create: `192.0.0.0/24`, Update: `0.0.0.0/24`},
		"translated_destination": acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/24`},
	}

	CoreDrgNatPolicyDrgNatRuleRepresentation2 = map[string]interface{}{
		"drg_nat_policy_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_nat_policy.test_drg_nat_policy.id}`},
		"drg_nat_rule_priority":  acctest.Representation{RepType: acctest.Required, Create: `2`},
		"original_destination":   acctest.Representation{RepType: acctest.Required, Create: `192.0.0.0/24`},
		"translated_destination": acctest.Representation{RepType: acctest.Required, Create: `192.0.0.0/24`},
	}

	CoreDrgNatPolicyDrgNatRuleRepresentation3 = map[string]interface{}{
		"drg_nat_policy_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_nat_policy.test_drg_nat_policy.id}`},
		"drg_nat_rule_priority":  acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `4`},
		"original_source":        acctest.Representation{RepType: acctest.Required, Update: `192.0.0.0/24`},
		"translated_source":      acctest.Representation{RepType: acctest.Required, Update: `192.0.0.0/24`},
		"original_destination":   acctest.Representation{RepType: acctest.Required, Create: `192.0.0.0/24`, Update: `192.0.1.0/24`},
		"translated_destination": acctest.Representation{RepType: acctest.Required, Create: `192.0.1.0/24`},
	}
)

// issue-routing-tag: core/pnp
func TestCoreDrgNatPolicyDrgNatRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgNatPolicyDrgNatRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_nat_policy_drg_nat_rule.test_drg_nat_policy_drg_nat_rule"
	datasourceName := "data.oci_core_drg_nat_policy_drg_nat_rules.test_drg_nat_policy_drg_nat_rules"

	resourceName2 := "oci_core_drg_nat_policy_drg_nat_rule.test_drg_nat_policy_drg_nat_rule2"
	resourceName3 := "oci_core_drg_nat_policy_drg_nat_rule.test_drg_nat_policy_drg_nat_rule3"

	//var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreDrgNatPolicyDrgNatRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy_drg_nat_rule", "test_drg_nat_policy_drg_nat_rule", acctest.Optional, acctest.Create, CoreDrgNatPolicyDrgNatRuleRepresentation), "core", "drgNatPolicyDrgNatRule", t)

	//acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify create with multiple rules
		{
			Config: config + compartmentIdVariableStr + CoreDrgNatPolicyDrgNatRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy_drg_nat_rule", "test_drg_nat_policy_drg_nat_rule", acctest.Required, acctest.Create, CoreDrgNatPolicyDrgNatRuleRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy_drg_nat_rule", "test_drg_nat_policy_drg_nat_rule2", acctest.Required, acctest.Create, CoreDrgNatPolicyDrgNatRuleRepresentation2) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy_drg_nat_rule", "test_drg_nat_policy_drg_nat_rule3", acctest.Required, acctest.Create, CoreDrgNatPolicyDrgNatRuleRepresentation3),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_nat_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_nat_rule_priority"),
				resource.TestCheckResourceAttrSet(resourceName, "original_source"),
				resource.TestCheckResourceAttrSet(resourceName, "translated_source"),
				resource.TestCheckResourceAttrSet(resourceName, "original_destination"),
				resource.TestCheckResourceAttrSet(resourceName, "translated_destination"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				resource.TestCheckResourceAttrSet(resourceName2, "drg_nat_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName2, "drg_nat_rule_priority"),
				resource.TestCheckResourceAttrSet(resourceName2, "original_destination"),
				resource.TestCheckResourceAttrSet(resourceName2, "translated_destination"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				resource.TestCheckResourceAttrSet(resourceName3, "drg_nat_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName3, "drg_nat_rule_priority"),
				resource.TestCheckResourceAttrSet(resourceName3, "original_destination"),
				resource.TestCheckResourceAttrSet(resourceName3, "translated_destination"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
			),
		},

		// verify update rules
		{
			Config: config + compartmentIdVariableStr + CoreDrgNatPolicyDrgNatRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy_drg_nat_rule", "test_drg_nat_policy_drg_nat_rule", acctest.Required, acctest.Update, CoreDrgNatPolicyDrgNatRuleRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy_drg_nat_rule", "test_drg_nat_policy_drg_nat_rule2", acctest.Required, acctest.Update, CoreDrgNatPolicyDrgNatRuleRepresentation2) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy_drg_nat_rule", "test_drg_nat_policy_drg_nat_rule3", acctest.Required, acctest.Update, CoreDrgNatPolicyDrgNatRuleRepresentation3),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "drg_nat_rule_priority", `10`),
				resource.TestCheckResourceAttr(resourceName, "original_source", `192.0.0.0/24`),
				resource.TestCheckResourceAttr(resourceName, "translated_source", `192.0.0.0/24`),
				resource.TestCheckResourceAttr(resourceName, "original_destination", `0.0.0.0/24`),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_nat_policy_id"),

				resource.TestCheckResourceAttrSet(resourceName2, "id"),
				resource.TestCheckResourceAttrSet(resourceName2, "drg_nat_policy_id"),

				resource.TestCheckResourceAttr(resourceName3, "original_source", `192.0.0.0/24`),
				resource.TestCheckResourceAttr(resourceName3, "original_destination", `192.0.1.0/24`),
				resource.TestCheckResourceAttr(resourceName3, "drg_nat_rule_priority", `4`),
				resource.TestCheckResourceAttrSet(resourceName3, "id"),
				resource.TestCheckResourceAttrSet(resourceName3, "drg_nat_policy_id"),

				/*func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updatedr")
					}
					return err
				},*/
			),
		},

		// verify datasource
		{
			Config: config + CoreDrgNatPolicyDrgNatRuleResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_core_drg_nat_policy_drg_nat_rules",
					"test_drg_nat_policy_drg_nat_rules",
					acctest.Required, acctest.Create, CoreDrgNatPolicyDrgNatRuleDataSourceRepresentation,
				) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy_drg_nat_rule", "test_drg_nat_policy_drg_nat_rule", acctest.Required, acctest.Update, CoreDrgNatPolicyDrgNatRuleRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy_drg_nat_rule", "test_drg_nat_policy_drg_nat_rule2", acctest.Required, acctest.Update, CoreDrgNatPolicyDrgNatRuleRepresentation2) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy_drg_nat_rule", "test_drg_nat_policy_drg_nat_rule3", acctest.Required, acctest.Update, CoreDrgNatPolicyDrgNatRuleRepresentation3),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// Check datasource's own id/attributes set
				resource.TestCheckResourceAttrSet(datasourceName, "drg_nat_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.#", "3"),

				// Rule 1
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.0.drg_nat_rule_priority", "2"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.0.original_source", "0.0.0.0/0"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.0.translated_source", "0.0.0.0/0"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.0.original_destination", "192.0.0.0/24"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.0.translated_destination", "192.0.0.0/24"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_nat_rules.0.id"),

				// Rule 2
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.1.drg_nat_rule_priority", "4"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.1.original_source", "192.0.0.0/24"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.1.translated_source", "192.0.0.0/24"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.1.original_destination", "192.0.1.0/24"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.1.translated_destination", "192.0.1.0/24"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_nat_rules.1.id"),

				// Rule 3
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.2.drg_nat_rule_priority", "10"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.2.original_source", "192.0.0.0/24"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.2.translated_source", "192.0.0.0/24"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.2.original_destination", "0.0.0.0/24"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_rules.2.translated_destination", "0.0.0.0/24"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_nat_rules.2.id"),
			),
		},

		// verify resource import
		{
			Config:                  config + CoreDrgNatPolicyDrgNatRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},

		// delete - Terraform’s behavior (when you run terraform apply) is to destroy all resources that exist in the state but are now missing from the config.
		{
			Config: config + compartmentIdVariableStr + CoreDrgNatPolicyDrgNatRuleResourceDependencies,
		},
	})
}
