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
	waasPolicyScenarioRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"domain":         Representation{RepType: Required, Create: waasPolicyDomain},
		"origins":        []RepresentationGroup{{Optional, waasOriginRepresentationMap1}, {Optional, waasOriginRepresentationMap2}},
		"waf_config":     RepresentationGroup{Optional, waasPolicyWafConfigScenarioRepresentation},
		"timeouts":       RepresentationGroup{Required, waasPolicyTimeoutsRepresentation},
	}

	waasPolicyWafConfigScenarioRepresentation = map[string]interface{}{
		"caching_rules": RepresentationGroup{Optional, waasPolicyWafConfigCachingRulesScenarioRepresentation},
		"origin":        Representation{RepType: Optional, Create: `primary`, Update: `primary2`},
	}

	waasPolicyWafConfigCachingRulesScenarioRepresentation = map[string]interface{}{
		"action":                    Representation{RepType: Required, Create: `BYPASS_CACHE`},
		"criteria":                  RepresentationGroup{Required, waasPolicyWafConfigCachingRulesCriteriaRepresentation},
		"name":                      Representation{RepType: Required, Create: `name`, Update: `name2`},
		"is_client_caching_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"key":                       Representation{RepType: Optional, Create: `key`, Update: `key2`},
	}

	waasPolicyWafConfigScenarioRepresentation2 = map[string]interface{}{
		"caching_rules": RepresentationGroup{Optional, waasPolicyWafConfigCachingRulesScenarioRepresentation2},
		"origin":        Representation{RepType: Optional, Create: `primary`, Update: `primary2`},
	}

	waasPolicyWafConfigCachingRulesScenarioRepresentation2 = map[string]interface{}{
		"action":                    Representation{RepType: Required, Create: `CACHE`},
		"criteria":                  RepresentationGroup{Required, waasPolicyWafConfigCachingRulesCriteriaRepresentation},
		"caching_duration":          Representation{RepType: Optional, Create: `PT1S`, Update: `PT2S`},
		"client_caching_duration":   Representation{RepType: Optional, Create: `PT1S`, Update: `PT2S`},
		"name":                      Representation{RepType: Required, Create: `name`, Update: `name2`},
		"is_client_caching_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"key":                       Representation{RepType: Optional, Create: `key`, Update: `key2`},
	}

	WaasPolicyResourceCachingOnlyConfig = GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_scenario_waas_policy", Optional, Update,
		RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(waasPolicyScenarioRepresentation, []string{"waf_config"}),
			map[string]interface{}{"waf_config": RepresentationGroup{Optional, waasPolicyWafConfigScenarioRepresentation2}}))
)

// issue-routing-tag: waas/default
func TestResourceWaasWaasPolicyResource_basic(t *testing.T) {

	httpreplay.SetScenario("TestResourceWaasWaasPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_waas_waas_policy.test_scenario_waas_policy"

	var resId, resId2 string
	ResourceTest(t, testAccCheckWaasWaasPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_scenario_waas_policy", Optional, Create, waasPolicyScenarioRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", waasPolicyDomain),
				resource.TestCheckResourceAttr(resourceName, "waf_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.action", "BYPASS_CACHE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.0.condition", "URL_IS"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.0.value", "/public"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.is_client_caching_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.name", "name"),
				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify Update
		{
			Config: config + compartmentIdVariableStr + WaasPolicyResourceCachingOnlyConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", waasPolicyDomain),
				resource.TestCheckResourceAttr(resourceName, "waf_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.action", "CACHE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.caching_duration", "PT2S"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.client_caching_duration", "PT2S"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.0.condition", "URL_STARTS_WITH"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.0.value", "/publ"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.is_client_caching_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.name", "name2"),
				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
	})
}
