// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	steeringPolicyFailOverRepresentation = map[string]interface{}{
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":            Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"template":                Representation{repType: Required, create: `FAILOVER`},
		"answers":                 RepresentationGroup{Optional, steeringPolicyFailOverAnswersRepresentation},
		"defined_tags":            Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":           Representation{repType: Optional, create: map[string]string{"freeformTags": "freeformTags"}, update: map[string]string{"freeformTags2": "freeformTags2"}},
		"health_check_monitor_id": Representation{repType: Optional, create: `${oci_health_checks_http_monitor.test_http_monitor.id}`},
		"rules": []RepresentationGroup{
			{Optional, steeringPolicyFailOverRulesFilterRuleTypeRepresentation},
			{Optional, steeringPolicyFailOverRulesHealthRuleTypeRepresentation},
			{Optional, steeringPolicyFailOverRulesPriorityRuleTypeRepresentation},
			{Optional, steeringPolicyFailOverRulesLimitRuleTypeRepresentation},
		},
		"ttl": Representation{repType: Optional, create: `10`, update: `11`},
	}

	steeringPolicyFailOverAnswersRepresentation = map[string]interface{}{
		"name":        Representation{repType: Required, create: `name`},
		"rdata":       Representation{repType: Required, create: `192.0.2.1`},
		"rtype":       Representation{repType: Required, create: `A`},
		"is_disabled": Representation{repType: Optional, create: `false`},
		"pool":        Representation{repType: Optional, create: `primary`},
	}

	steeringPolicyFailOverRulesFilterRuleTypeRepresentation = map[string]interface{}{
		"rule_type":           Representation{repType: Required, create: `FILTER`},
		"default_answer_data": RepresentationGroup{Optional, steeringPolicyFailOverRulesDefaultAnswerDataFilterRuleTypeRepresentation},
		"description":         Representation{repType: Optional, create: `filter description`},
	}

	steeringPolicyFailOverRulesDefaultAnswerDataFilterRuleTypeRepresentation = map[string]interface{}{
		"answer_condition": Representation{repType: Optional, create: `answer.isDisabled != true`},
		"should_keep":      Representation{repType: Optional, create: `true`},
	}

	steeringPolicyFailOverRulesHealthRuleTypeRepresentation = map[string]interface{}{
		"rule_type":   Representation{repType: Required, create: `HEALTH`},
		"description": Representation{repType: Optional, create: `health description`},
	}

	steeringPolicyFailOverRulesPriorityRuleTypeRepresentation = map[string]interface{}{
		"rule_type":           Representation{repType: Required, create: `PRIORITY`},
		"default_answer_data": RepresentationGroup{Optional, steeringPolicyFailOverRulesDefaultAnswerDataPriorityRuleTypeRepresentation},
		"description":         Representation{repType: Optional, create: `priority description`},
	}

	steeringPolicyFailOverRulesDefaultAnswerDataPriorityRuleTypeRepresentation = map[string]interface{}{
		"answer_condition": Representation{repType: Optional, create: `answer.pool == 'primary'`},
		"value":            Representation{repType: Optional, create: `1`},
	}

	steeringPolicyFailOverRulesLimitRuleTypeRepresentation = map[string]interface{}{
		"rule_type":     Representation{repType: Required, create: `LIMIT`},
		"default_count": Representation{repType: Optional, create: `1`},
		"description":   Representation{repType: Optional, create: `limit description`},
	}

	SteeringPolicyFailOverResourceDependencies = HttpMonitorRequiredOnlyResource
)

func TestResourceDnsSteeringPolicyFailOver(t *testing.T) {
	httpreplay.SetScenario("TestResourceDnsSteeringPolicyFailOver")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_steering_policy.test_steering_policy"

	content := config + compartmentIdVariableStr + SteeringPolicyFailOverResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", Optional, Create, steeringPolicyFailOverRepresentation)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDnsSteeringPolicyDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: content,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "answers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "answers.0.is_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "answers.0.name", "name"),
					resource.TestCheckResourceAttr(resourceName, "answers.0.pool", "primary"),
					resource.TestCheckResourceAttr(resourceName, "answers.0.rtype", "A"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "health_check_monitor_id"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.0.answer_condition", "answer.isDisabled != true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.0.should_keep", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "filter description"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.rule_type", "FILTER"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "health description"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.rule_type", "HEALTH"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.default_answer_data.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.default_answer_data.0.answer_condition", "answer.pool == 'primary'"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.default_answer_data.0.value", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.description", "priority description"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.rule_type", "PRIORITY"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.default_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.description", "limit description"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.rule_type", "LIMIT"),
					resource.TestCheckResourceAttr(resourceName, "template", "FAILOVER"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
		},
	})
}
