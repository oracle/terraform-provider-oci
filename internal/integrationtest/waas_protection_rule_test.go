// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	WaasProtectionRuleRepresentation = map[string]interface{}{
		"waas_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waas_waas_policy.test_waas_policy.id}`},
		"key":            acctest.Representation{RepType: acctest.Required, Create: `933161`, Update: `933111`},
		"action":         acctest.Representation{RepType: acctest.Required, Create: `BLOCK`, Update: `DETECT`},
		"exclusions":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: WaasProtectionRuleExclusionsRepresentation},
	}

	WaasProtectionRuleExclusionsRepresentation = map[string]interface{}{
		"exclusions": acctest.Representation{RepType: acctest.Optional, Create: []string{`example.com`}, Update: []string{`OAMAuthnCookie`}},
		"target":     acctest.Representation{RepType: acctest.Optional, Create: `REQUEST_COOKIES`, Update: `REQUEST_COOKIE_NAMES`},
	}

	WaasProtectionRuleSingularDataSourceRepresentation = map[string]interface{}{
		"protection_rule_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_waas_protection_rule.test_protection_rule.key}`},
		"waas_policy_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_waas_waas_policy.test_waas_policy.id}`},
	}

	WaasWaasProtectionRuleDataSourceRepresentation = map[string]interface{}{
		"waas_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waas_waas_policy.test_waas_policy.id}`},
		"action":         acctest.Representation{RepType: acctest.Optional, Create: []string{`DETECT`}},
	}

	ProtectionRuleResourceConfig = WaasWaasPolicyResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", acctest.Optional, acctest.Create, WaasWaasPolicyRepresentation)
)

// issue-routing-tag: waas/default
func TestWaasProtectionRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasProtectionRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_waas_protection_rule.test_protection_rule"
	datasourceName := "data.oci_waas_protection_rules.test_protection_rules"
	singularDatasourceName := "data.oci_waas_protection_rule.test_protection_rule"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ProtectionRuleResourceConfig+
		acctest.GenerateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", acctest.Required, acctest.Create, WaasProtectionRuleRepresentation), "waas", "protectionRule", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ProtectionRuleResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", acctest.Required, acctest.Create, WaasProtectionRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "key", "933161"),
				resource.TestCheckResourceAttr(resourceName, "action", "BLOCK"),
				resource.TestCheckResourceAttrSet(resourceName, "waas_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "labels.#"),
				resource.TestCheckResourceAttrSet(resourceName, "mod_security_rule_ids.#"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),

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
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ProtectionRuleResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", acctest.Optional, acctest.Update, WaasProtectionRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waas_protection_rules", "test_protection_rules", acctest.Optional, acctest.Update, WaasWaasProtectionRuleDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", acctest.Optional, acctest.Update, WaasProtectionRuleRepresentation) +
				compartmentIdVariableStr + ProtectionRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", acctest.Required, acctest.Create, WaasProtectionRuleSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", acctest.Optional, acctest.Update, WaasProtectionRuleRepresentation) +
				compartmentIdVariableStr + ProtectionRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
