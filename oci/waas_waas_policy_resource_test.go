// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"domain":         Representation{repType: Required, create: waasPolicyDomain},
		"origins":        []RepresentationGroup{{Optional, waasOriginRepresentationMap1}, {Optional, waasOriginRepresentationMap2}},
		"waf_config":     RepresentationGroup{Optional, waasPolicyWafConfigScenarioRepresentation},
		"timeouts":       RepresentationGroup{Required, waasPolicyTimeoutsRepresentation},
	}

	waasPolicyWafConfigScenarioRepresentation = map[string]interface{}{
		"caching_rules": RepresentationGroup{Optional, waasPolicyWafConfigCachingRulesScenarioRepresentation},
		"origin":        Representation{repType: Optional, create: `primary`, update: `primary2`},
	}

	waasPolicyWafConfigCachingRulesScenarioRepresentation = map[string]interface{}{
		"action":                    Representation{repType: Required, create: `BYPASS_CACHE`},
		"criteria":                  RepresentationGroup{Required, waasPolicyWafConfigCachingRulesCriteriaRepresentation},
		"name":                      Representation{repType: Required, create: `name`, update: `name2`},
		"is_client_caching_enabled": Representation{repType: Optional, create: `false`, update: `true`},
		"key":                       Representation{repType: Optional, create: `key`, update: `key2`},
	}

	waasPolicyWafConfigScenarioRepresentation2 = map[string]interface{}{
		"caching_rules": RepresentationGroup{Optional, waasPolicyWafConfigCachingRulesScenarioRepresentation2},
		"origin":        Representation{repType: Optional, create: `primary`, update: `primary2`},
	}

	waasPolicyWafConfigCachingRulesScenarioRepresentation2 = map[string]interface{}{
		"action":                    Representation{repType: Required, create: `CACHE`},
		"criteria":                  RepresentationGroup{Required, waasPolicyWafConfigCachingRulesCriteriaRepresentation},
		"caching_duration":          Representation{repType: Optional, create: `PT1S`, update: `PT2S`},
		"client_caching_duration":   Representation{repType: Optional, create: `PT1S`, update: `PT2S`},
		"name":                      Representation{repType: Required, create: `name`, update: `name2`},
		"is_client_caching_enabled": Representation{repType: Optional, create: `false`, update: `true`},
		"key":                       Representation{repType: Optional, create: `key`, update: `key2`},
	}

	WaasPolicyResourceCachingOnlyConfig = generateResourceFromRepresentationMap("oci_waas_waas_policy", "test_scenario_waas_policy", Optional, Update,
		representationCopyWithNewProperties(representationCopyWithRemovedProperties(waasPolicyScenarioRepresentation, []string{"waf_config"}),
			map[string]interface{}{"waf_config": RepresentationGroup{Optional, waasPolicyWafConfigScenarioRepresentation2}}))
)

func TestResourceWaasWaasPolicyResource_basic(t *testing.T) {

	httpreplay.SetScenario("TestResourceWaasWaasPolicyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_waas_waas_policy.test_scenario_waas_policy"

	var resId, resId2 string
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckWaasWaasPolicyDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr +
					generateResourceFromRepresentationMap("oci_waas_waas_policy", "test_scenario_waas_policy", Optional, Create, waasPolicyScenarioRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},

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
			// verify update
			{
				Config: config + compartmentIdVariableStr + WaasPolicyResourceCachingOnlyConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
		},
	})
}
