// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	protectionRuleRepresentation = map[string]interface{}{
		"waas_policy_id": Representation{RepType: Required, Create: `${oci_waas_waas_policy.test_waas_policy.id}`},
		"key":            Representation{RepType: Required, Create: `933161`, Update: `933111`},
		"action":         Representation{RepType: Required, Create: `BLOCK`, Update: `DETECT`},
		"exclusions":     RepresentationGroup{Optional, protectionRuleExclusionsRepresentation},
	}

	protectionRuleExclusionsRepresentation = map[string]interface{}{
		"exclusions": Representation{RepType: Optional, Create: []string{`example.com`}, Update: []string{`OAMAuthnCookie`}},
		"target":     Representation{RepType: Optional, Create: `REQUEST_COOKIES`, Update: `REQUEST_COOKIE_NAMES`},
	}

	protectionRuleSingularDataSourceRepresentation = map[string]interface{}{
		"protection_rule_key": Representation{RepType: Required, Create: `${oci_waas_protection_rule.test_protection_rule.key}`},
		"waas_policy_id":      Representation{RepType: Required, Create: `${oci_waas_waas_policy.test_waas_policy.id}`},
	}

	protectionRuleDataSourceRepresentation = map[string]interface{}{
		"waas_policy_id": Representation{RepType: Required, Create: `${oci_waas_waas_policy.test_waas_policy.id}`},
		"action":         Representation{RepType: Optional, Create: []string{`DETECT`}},
	}

	ProtectionRuleResourceConfig = WaasPolicyResourceDependencies + GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", Optional, Create, waasPolicyRepresentation)
)

// issue-routing-tag: waas/default
func TestWaasProtectionRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasProtectionRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_waas_protection_rule.test_protection_rule"
	datasourceName := "data.oci_waas_protection_rules.test_protection_rules"
	singularDatasourceName := "data.oci_waas_protection_rule.test_protection_rule"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ProtectionRuleResourceConfig+
		GenerateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Required, Create, protectionRuleRepresentation), "waas", "protectionRule", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ProtectionRuleResourceConfig +
				GenerateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Required, Create, protectionRuleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "key", "933161"),
				resource.TestCheckResourceAttr(resourceName, "action", "BLOCK"),
				resource.TestCheckResourceAttrSet(resourceName, "waas_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "labels.#"),
				resource.TestCheckResourceAttrSet(resourceName, "mod_security_rule_ids.#"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ProtectionRuleResourceConfig +
				GenerateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Optional, Update, protectionRuleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "key", "933111"),
				resource.TestCheckResourceAttr(resourceName, "action", "DETECT"),
				resource.TestCheckResourceAttrSet(resourceName, "waas_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "labels.#"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
			),
		},

		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_waas_protection_rules", "test_protection_rules", Optional, Update, protectionRuleDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Optional, Update, protectionRuleRepresentation) +
				compartmentIdVariableStr + ProtectionRuleResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "action.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "waas_policy_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "protection_rules.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "protection_rules.0.action"),
				resource.TestCheckResourceAttrSet(datasourceName, "protection_rules.0.description"),
				resource.TestCheckResourceAttr(datasourceName, "protection_rules.0.exclusions.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "protection_rules.0.key"),
				resource.TestCheckResourceAttrSet(datasourceName, "protection_rules.0.name"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Required, Create, protectionRuleSingularDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Optional, Update, protectionRuleRepresentation) +
				compartmentIdVariableStr + ProtectionRuleResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "protection_rule_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "waas_policy_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "action"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "exclusions.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "labels.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
			),
		},
	})
}
