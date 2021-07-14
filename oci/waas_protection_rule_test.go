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
		"waas_policy_id": Representation{repType: Required, create: `${oci_waas_waas_policy.test_waas_policy.id}`},
		"key":            Representation{repType: Required, create: `933161`, update: `933111`},
		"action":         Representation{repType: Required, create: `BLOCK`, update: `DETECT`},
		"exclusions":     RepresentationGroup{Optional, protectionRuleExclusionsRepresentation},
	}

	protectionRuleExclusionsRepresentation = map[string]interface{}{
		"exclusions": Representation{repType: Optional, create: []string{`example.com`}, update: []string{`OAMAuthnCookie`}},
		"target":     Representation{repType: Optional, create: `REQUEST_COOKIES`, update: `REQUEST_COOKIE_NAMES`},
	}

	protectionRuleSingularDataSourceRepresentation = map[string]interface{}{
		"protection_rule_key": Representation{repType: Required, create: `${oci_waas_protection_rule.test_protection_rule.key}`},
		"waas_policy_id":      Representation{repType: Required, create: `${oci_waas_waas_policy.test_waas_policy.id}`},
	}

	protectionRuleDataSourceRepresentation = map[string]interface{}{
		"waas_policy_id": Representation{repType: Required, create: `${oci_waas_waas_policy.test_waas_policy.id}`},
		"action":         Representation{repType: Optional, create: []string{`DETECT`}},
	}

	ProtectionRuleResourceConfig = WaasPolicyResourceDependencies + generateResourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", Optional, Create, waasPolicyRepresentation)
)

func TestWaasProtectionRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasProtectionRuleResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_waas_protection_rule.test_protection_rule"
	datasourceName := "data.oci_waas_protection_rules.test_protection_rules"
	singularDatasourceName := "data.oci_waas_protection_rule.test_protection_rule"

	var resId string
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ProtectionRuleResourceConfig+
		generateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Required, Create, protectionRuleRepresentation), "waas", "protectionRule", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ProtectionRuleResourceConfig +
					generateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Required, Create, protectionRuleRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "key", "933161"),
					resource.TestCheckResourceAttr(resourceName, "action", "BLOCK"),
					resource.TestCheckResourceAttrSet(resourceName, "waas_policy_id"),
					resource.TestCheckResourceAttrSet(resourceName, "labels.#"),
					resource.TestCheckResourceAttrSet(resourceName, "mod_security_rule_ids.#"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Optional, Update, protectionRuleRepresentation),
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
					generateDataSourceFromRepresentationMap("oci_waas_protection_rules", "test_protection_rules", Optional, Update, protectionRuleDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Optional, Update, protectionRuleRepresentation) +
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
					generateDataSourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Required, Create, protectionRuleSingularDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_waas_protection_rule", "test_protection_rule", Optional, Update, protectionRuleRepresentation) +
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
		},
	})
}
