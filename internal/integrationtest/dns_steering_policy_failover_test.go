// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	steeringPolicyFailOverRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"template":                acctest.Representation{RepType: acctest.Required, Create: `FAILOVER`},
		"answers":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: steeringPolicyFailOverAnswersRepresentation},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"health_check_monitor_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_health_checks_http_monitor.test_http_monitor.id}`},
		"rules": []acctest.RepresentationGroup{
			{RepType: acctest.Optional, Group: steeringPolicyFailOverRulesFilterRuleTypeRepresentation},
			{RepType: acctest.Optional, Group: steeringPolicyFailOverRulesHealthRuleTypeRepresentation},
			{RepType: acctest.Optional, Group: steeringPolicyFailOverRulesPriorityRuleTypeRepresentation},
			{RepType: acctest.Optional, Group: steeringPolicyFailOverRulesLimitRuleTypeRepresentation},
		},
		"ttl": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	steeringPolicyFailOverAnswersRepresentation = map[string]interface{}{
		"name":        acctest.Representation{RepType: acctest.Required, Create: `name`},
		"rdata":       acctest.Representation{RepType: acctest.Required, Create: `192.0.2.1`},
		"rtype":       acctest.Representation{RepType: acctest.Required, Create: `A`},
		"is_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"pool":        acctest.Representation{RepType: acctest.Optional, Create: `primary`},
	}

	steeringPolicyFailOverRulesFilterRuleTypeRepresentation = map[string]interface{}{
		"rule_type":           acctest.Representation{RepType: acctest.Required, Create: `FILTER`},
		"default_answer_data": acctest.RepresentationGroup{RepType: acctest.Optional, Group: steeringPolicyFailOverRulesDefaultAnswerDataFilterRuleTypeRepresentation},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `filter description`},
	}

	steeringPolicyFailOverRulesDefaultAnswerDataFilterRuleTypeRepresentation = map[string]interface{}{
		"answer_condition": acctest.Representation{RepType: acctest.Optional, Create: `answer.isDisabled != true`},
		"should_keep":      acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	steeringPolicyFailOverRulesHealthRuleTypeRepresentation = map[string]interface{}{
		"rule_type":   acctest.Representation{RepType: acctest.Required, Create: `HEALTH`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `health description`},
	}

	steeringPolicyFailOverRulesPriorityRuleTypeRepresentation = map[string]interface{}{
		"rule_type":           acctest.Representation{RepType: acctest.Required, Create: `PRIORITY`},
		"default_answer_data": acctest.RepresentationGroup{RepType: acctest.Optional, Group: steeringPolicyFailOverRulesDefaultAnswerDataPriorityRuleTypeRepresentation},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `priority description`},
	}

	steeringPolicyFailOverRulesDefaultAnswerDataPriorityRuleTypeRepresentation = map[string]interface{}{
		"answer_condition": acctest.Representation{RepType: acctest.Optional, Create: `answer.pool == 'primary'`},
		"value":            acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	steeringPolicyFailOverRulesLimitRuleTypeRepresentation = map[string]interface{}{
		"rule_type":     acctest.Representation{RepType: acctest.Required, Create: `LIMIT`},
		"default_count": acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `limit description`},
	}

	SteeringPolicyFailOverResourceDependencies = HttpMonitorRequiredOnlyResource
)

// issue-routing-tag: dns/default
func TestResourceDnsSteeringPolicyFailOver(t *testing.T) {
	httpreplay.SetScenario("TestResourceDnsSteeringPolicyFailOver")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_steering_policy.test_steering_policy"

	content := config + compartmentIdVariableStr + SteeringPolicyFailOverResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", acctest.Optional, acctest.Create, steeringPolicyFailOverRepresentation)

	acctest.ResourceTest(t, testAccCheckDnsSteeringPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: content,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "answers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.is_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.pool", "primary"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.rtype", "A"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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
	})
}
