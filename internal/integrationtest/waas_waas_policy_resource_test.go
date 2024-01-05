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
	waasPolicyScenarioRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"domain":         acctest.Representation{RepType: acctest.Required, Create: waasPolicyDomain},
		"origins":        []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: waasOriginRepresentationMap1}, {RepType: acctest.Optional, Group: waasOriginRepresentationMap2}},
		"waf_config":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: waasPolicyWafConfigScenarioRepresentation},
		"timeouts":       acctest.RepresentationGroup{RepType: acctest.Required, Group: waasPolicyTimeoutsRepresentation},
	}

	waasPolicyWafConfigScenarioRepresentation = map[string]interface{}{
		"caching_rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: waasPolicyWafConfigCachingRulesScenarioRepresentation},
		"origin":        acctest.Representation{RepType: acctest.Optional, Create: `primary`, Update: `primary2`},
	}

	waasPolicyWafConfigCachingRulesScenarioRepresentation = map[string]interface{}{
		"action":                    acctest.Representation{RepType: acctest.Required, Create: `BYPASS_CACHE`},
		"criteria":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: WaasWaasPolicyWafConfigCachingRulesCriteriaRepresentation},
		"name":                      acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"is_client_caching_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key":                       acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
	}

	waasPolicyWafConfigScenarioRepresentation2 = map[string]interface{}{
		"caching_rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: waasPolicyWafConfigCachingRulesScenarioRepresentation2},
		"origin":        acctest.Representation{RepType: acctest.Optional, Create: `primary`, Update: `primary2`},
	}

	waasPolicyWafConfigCachingRulesScenarioRepresentation2 = map[string]interface{}{
		"action":                    acctest.Representation{RepType: acctest.Required, Create: `CACHE`},
		"criteria":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: WaasWaasPolicyWafConfigCachingRulesCriteriaRepresentation},
		"caching_duration":          acctest.Representation{RepType: acctest.Optional, Create: `PT1S`, Update: `PT2S`},
		"client_caching_duration":   acctest.Representation{RepType: acctest.Optional, Create: `PT1S`, Update: `PT2S`},
		"name":                      acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"is_client_caching_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key":                       acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
	}

	WaasPolicyResourceCachingOnlyConfig = acctest.GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_scenario_waas_policy", acctest.Optional, acctest.Update,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(waasPolicyScenarioRepresentation, []string{"waf_config"}),
			map[string]interface{}{"waf_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: waasPolicyWafConfigScenarioRepresentation2}}))
)

// issue-routing-tag: waas/default
func TestResourceWaasWaasPolicyResource_basic(t *testing.T) {

	httpreplay.SetScenario("TestResourceWaasWaasPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_waas_waas_policy.test_scenario_waas_policy"

	var resId, resId2 string
	acctest.ResourceTest(t, testAccCheckWaasWaasPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_scenario_waas_policy", acctest.Optional, acctest.Create, waasPolicyScenarioRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},

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
		// verify Update
		{
			Config: config + compartmentIdVariableStr + WaasPolicyResourceCachingOnlyConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
	})
}
