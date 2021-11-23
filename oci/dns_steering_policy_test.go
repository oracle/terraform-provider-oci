// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v53/common"
	oci_dns "github.com/oracle/oci-go-sdk/v53/dns"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SteeringPolicyRequiredOnlyResource = SteeringPolicyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", Required, Create, steeringPolicyRepresentation)

	SteeringPolicyResourceConfig = SteeringPolicyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", Optional, Update, steeringPolicyRepresentation)

	steeringPolicySingularDataSourceRepresentation = map[string]interface{}{
		"steering_policy_id": Representation{RepType: Required, Create: `${oci_dns_steering_policy.test_steering_policy.id}`},
	}

	steeringPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        Representation{RepType: Required, Create: `${var.compartment_id}`},
		"health_check_monitor_id":               Representation{RepType: Optional, Create: `${oci_health_checks_http_monitor.test_http_monitor.id}`},
		"id":                                    Representation{RepType: Optional, Create: `${oci_dns_steering_policy.test_steering_policy.id}`},
		"state":                                 Representation{RepType: Optional, Create: `ACTIVE`},
		"template":                              Representation{RepType: Optional, Create: `CUSTOM`},
		"time_created_greater_than_or_equal_to": Representation{RepType: Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                Representation{RepType: Optional, Create: `2038-01-01T00:00:00.000Z`},
		"filter":                                RepresentationGroup{Required, steeringPolicyDataSourceFilterRepresentation}}
	steeringPolicyDataSourceRepresentationWithDisplayNameFilter = RepresentationCopyWithNewProperties(steeringPolicyDataSourceRepresentation, map[string]interface{}{
		"display_name": Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
	})
	steeringPolicyDataSourceRepresentationWithDisplayNameContainsFilter = RepresentationCopyWithNewProperties(steeringPolicyDataSourceRepresentation, map[string]interface{}{
		"display_name_contains": Representation{RepType: Optional, Create: `displayName`},
	})
	steeringPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_dns_steering_policy.test_steering_policy.id}`}},
	}

	steeringPolicyRepresentation = map[string]interface{}{
		"compartment_id":          Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":            Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"template":                Representation{RepType: Required, Create: `CUSTOM`},
		"answers":                 RepresentationGroup{Optional, steeringPolicyAnswersRepresentation},
		"defined_tags":            Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":           Representation{RepType: Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"health_check_monitor_id": Representation{RepType: Optional, Create: `${oci_health_checks_http_monitor.test_http_monitor.id}`},
		"rules": []RepresentationGroup{
			{Optional, steeringPolicyRulesFilterRuleTypeRepresentation},
			{Optional, steeringPolicyRulesHealthRuleTypeRepresentation},
			{Optional, steeringPolicyRulesLimitRuleTypeRepresentation},
			{Optional, steeringPolicyRulesPriorityRuleTypeRepresentation},
			{Optional, steeringPolicyRulesWeightedRuleTypeRepresentation},
		},
		"ttl": Representation{RepType: Optional, Create: `10`, Update: `11`},
	}
	steeringPolicyAnswersRepresentation = map[string]interface{}{
		"name":        Representation{RepType: Required, Create: `name`},
		"rdata":       Representation{RepType: Required, Create: `192.0.2.1`},
		"rtype":       Representation{RepType: Required, Create: `A`},
		"is_disabled": Representation{RepType: Optional, Create: `false`},
		"pool":        Representation{RepType: Optional, Create: `pool`},
	}
	steeringPolicyRulesFilterRuleTypeRepresentation = map[string]interface{}{
		"rule_type":           Representation{RepType: Required, Create: `FILTER`},
		"cases":               RepresentationGroup{Optional, steeringPolicyRulesCasesFilterRuleTypeRepresentation},
		"default_answer_data": RepresentationGroup{Optional, steeringPolicyRulesDefaultAnswerDataFilterRuleTypeRepresentation},
		"description":         Representation{RepType: Optional, Create: `filter description`},
	}
	steeringPolicyRulesCasesFilterRuleTypeRepresentation = map[string]interface{}{
		"answer_data":    RepresentationGroup{Optional, steeringPolicyRulesCasesAnswerDataFilterRuleTypeRepresentation},
		"case_condition": Representation{RepType: Optional, Create: `query.client.address in (subnet '198.51.100.0/24')`},
	}
	steeringPolicyRulesDefaultAnswerDataFilterRuleTypeRepresentation = map[string]interface{}{
		"answer_condition": Representation{RepType: Optional, Create: `answer.name == 'sampler'`},
		"should_keep":      Representation{RepType: Optional, Create: `false`},
	}
	steeringPolicyRulesCasesAnswerDataFilterRuleTypeRepresentation = map[string]interface{}{
		"answer_condition": Representation{RepType: Optional, Create: `answer.name == 'sampler'`},
		"should_keep":      Representation{RepType: Optional, Create: `false`},
	}
	steeringPolicyRulesHealthRuleTypeRepresentation = map[string]interface{}{
		"rule_type":   Representation{RepType: Required, Create: `HEALTH`},
		"cases":       RepresentationGroup{Optional, steeringPolicyRulesCasesHealthRuleTypeRepresentation},
		"description": Representation{RepType: Optional, Create: `health description`},
	}
	steeringPolicyRulesCasesHealthRuleTypeRepresentation = map[string]interface{}{
		"case_condition": Representation{RepType: Optional, Create: `query.client.address in (subnet '198.51.100.0/24')`},
	}
	steeringPolicyRulesLimitRuleTypeRepresentation = map[string]interface{}{
		"rule_type":     Representation{RepType: Required, Create: `LIMIT`},
		"cases":         RepresentationGroup{Optional, steeringPolicyRulesCasesLimitRuleTypeRepresentation},
		"default_count": Representation{RepType: Optional, Create: `10`},
		"description":   Representation{RepType: Optional, Create: `limit description`},
	}
	steeringPolicyRulesCasesLimitRuleTypeRepresentation = map[string]interface{}{
		"case_condition": Representation{RepType: Optional, Create: `query.client.address in (subnet '198.51.100.0/24')`},
		"count":          Representation{RepType: Optional, Create: `10`},
	}
	steeringPolicyRulesPriorityRuleTypeRepresentation = map[string]interface{}{
		"rule_type":           Representation{RepType: Required, Create: `PRIORITY`},
		"cases":               RepresentationGroup{Optional, steeringPolicyRulesCasesPriorityRuleTypeRepresentation},
		"default_answer_data": RepresentationGroup{Optional, steeringPolicyRulesDefaultAnswerDataPriorityRuleTypeRepresentation},
		"description":         Representation{RepType: Optional, Create: `priority description`},
	}
	steeringPolicyRulesCasesPriorityRuleTypeRepresentation = map[string]interface{}{
		"answer_data":    RepresentationGroup{Optional, steeringPolicyRulesCasesAnswerDataPriorityRuleTypeRepresentation},
		"case_condition": Representation{RepType: Optional, Create: `query.client.address in (subnet '198.51.100.0/24')`},
	}
	steeringPolicyRulesDefaultAnswerDataPriorityRuleTypeRepresentation = map[string]interface{}{
		"answer_condition": Representation{RepType: Optional, Create: `answer.name == 'sampler'`},
		"value":            Representation{RepType: Optional, Create: `10`},
	}
	steeringPolicyRulesCasesAnswerDataPriorityRuleTypeRepresentation = map[string]interface{}{
		"answer_condition": Representation{RepType: Optional, Create: `answer.name == 'sampler'`},
		"value":            Representation{RepType: Optional, Create: `10`},
	}
	steeringPolicyRulesWeightedRuleTypeRepresentation = map[string]interface{}{
		"rule_type":           Representation{RepType: Required, Create: `WEIGHTED`},
		"cases":               RepresentationGroup{Optional, steeringPolicyRulesCasesWeightedRuleTypeRepresentation},
		"default_answer_data": RepresentationGroup{Optional, steeringPolicyRulesDefaultAnswerDataWeightedRuleTypeRepresentation},
		"description":         Representation{RepType: Optional, Create: `weighted description`},
	}
	steeringPolicyRulesCasesWeightedRuleTypeRepresentation = map[string]interface{}{
		"answer_data":    RepresentationGroup{Optional, steeringPolicyRulesCasesAnswerDataWeightedRuleTypeRepresentation},
		"case_condition": Representation{RepType: Optional, Create: `query.client.address in (subnet '198.51.100.0/24')`},
	}
	steeringPolicyRulesDefaultAnswerDataWeightedRuleTypeRepresentation = map[string]interface{}{
		"answer_condition": Representation{RepType: Optional, Create: `answer.name == 'sampler'`},
		"value":            Representation{RepType: Optional, Create: `10`},
	}
	steeringPolicyRulesCasesAnswerDataWeightedRuleTypeRepresentation = map[string]interface{}{
		"answer_condition": Representation{RepType: Optional, Create: `answer.name == 'sampler'`},
		"value":            Representation{RepType: Optional, Create: `10`},
	}

	SteeringPolicyResourceDependencies = GenerateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", Required, Create, httpMonitorRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: dns/default
func TestDnsSteeringPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsSteeringPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dns_steering_policy.test_steering_policy"

	datasourceName := "data.oci_dns_steering_policies.test_steering_policies"
	singularDatasourceName := "data.oci_dns_steering_policy.test_steering_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+SteeringPolicyResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", Optional, Create, steeringPolicyRepresentation), "dns", "steeringPolicy", t)

	ResourceTest(t, testAccCheckDnsSteeringPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SteeringPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", Required, Create, steeringPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "template", "CUSTOM"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SteeringPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SteeringPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", Optional, Create, steeringPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "answers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.is_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.pool", "pool"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.rdata", "192.0.2.1"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.rtype", "A"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "health_check_monitor_id"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "5"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.answer_data.0.should_keep", "false"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.0.should_keep", "false"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.description", "filter description"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.rule_type", "FILTER"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.description", "health description"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.rule_type", "HEALTH"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.cases.0.count", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.default_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.description", "limit description"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.rule_type", "LIMIT"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.default_answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.description", "priority description"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.rule_type", "PRIORITY"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.default_answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.description", "weighted description"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.rule_type", "WEIGHTED"),
				resource.TestCheckResourceAttr(resourceName, "template", "CUSTOM"),
				resource.TestCheckResourceAttr(resourceName, "ttl", "10"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SteeringPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", Optional, Create,
					RepresentationCopyWithNewProperties(steeringPolicyRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "answers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.is_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.pool", "pool"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.rdata", "192.0.2.1"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.rtype", "A"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "health_check_monitor_id"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "5"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.answer_data.0.should_keep", "false"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.0.should_keep", "false"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.description", "filter description"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.rule_type", "FILTER"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.description", "health description"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.rule_type", "HEALTH"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.cases.0.count", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.default_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.description", "limit description"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.rule_type", "LIMIT"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.default_answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.description", "priority description"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.rule_type", "PRIORITY"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.default_answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.description", "weighted description"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.rule_type", "WEIGHTED"),
				resource.TestCheckResourceAttr(resourceName, "template", "CUSTOM"),
				resource.TestCheckResourceAttr(resourceName, "ttl", "10"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + SteeringPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", Optional, Update, steeringPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "answers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.is_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.pool", "pool"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.rdata", "192.0.2.1"),
				resource.TestCheckResourceAttr(resourceName, "answers.0.rtype", "A"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "health_check_monitor_id"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "5"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.answer_data.0.should_keep", "false"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.default_answer_data.0.should_keep", "false"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.description", "filter description"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.rule_type", "FILTER"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.description", "health description"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.rule_type", "HEALTH"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.cases.0.count", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.default_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.description", "limit description"),
				resource.TestCheckResourceAttr(resourceName, "rules.2.rule_type", "LIMIT"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.default_answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.description", "priority description"),
				resource.TestCheckResourceAttr(resourceName, "rules.3.rule_type", "PRIORITY"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.default_answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.description", "weighted description"),
				resource.TestCheckResourceAttr(resourceName, "rules.4.rule_type", "WEIGHTED"),
				resource.TestCheckResourceAttr(resourceName, "template", "CUSTOM"),
				resource.TestCheckResourceAttr(resourceName, "ttl", "11"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_dns_steering_policies", "test_steering_policies", Optional, Update, steeringPolicyDataSourceRepresentationWithDisplayNameFilter) +
				GenerateDataSourceFromRepresentationMap("oci_dns_steering_policies", "test_steering_policies2", Optional, Update, steeringPolicyDataSourceRepresentationWithDisplayNameContainsFilter) +
				compartmentIdVariableStr + SteeringPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", Optional, Update, steeringPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "health_check_monitor_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "template", "CUSTOM"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "steering_policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "steering_policies.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "steering_policies.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "steering_policies.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policies.0.health_check_monitor_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policies.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policies.0.self"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policies.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "steering_policies.0.template", "CUSTOM"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policies.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "steering_policies.0.ttl", "11"),

				resource.TestCheckResourceAttr(datasourceName+"2", "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName+"2", "display_name_contains", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName+"2", "health_check_monitor_id"),
				resource.TestCheckResourceAttr(datasourceName+"2", "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName+"2", "template", "CUSTOM"),
				resource.TestCheckResourceAttr(datasourceName+"2", "time_created_greater_than_or_equal_to", "2018-01-01T00:00:00.000Z"),
				resource.TestCheckResourceAttr(datasourceName+"2", "time_created_less_than", "2038-01-01T00:00:00.000Z"),

				resource.TestCheckResourceAttr(datasourceName+"2", "steering_policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName+"2", "steering_policies.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName+"2", "steering_policies.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName+"2", "steering_policies.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName+"2", "steering_policies.0.health_check_monitor_id"),
				resource.TestCheckResourceAttr(datasourceName+"2", "steering_policies.0.template", "CUSTOM"),
				resource.TestCheckResourceAttr(datasourceName+"2", "steering_policies.0.ttl", "11"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", Required, Create, steeringPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + SteeringPolicyResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "health_check_monitor_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "steering_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "answers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "answers.0.is_disabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "answers.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "answers.0.pool", "pool"),
				resource.TestCheckResourceAttr(singularDatasourceName, "answers.0.rdata", "192.0.2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "answers.0.rtype", "A"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "health_check_monitor_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.cases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.cases.0.answer_data.0.should_keep", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.default_answer_data.0.should_keep", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.description", "filter description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.rule_type", "FILTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.1.cases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.1.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.1.description", "health description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.1.rule_type", "HEALTH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.2.cases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.2.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.2.cases.0.count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.2.default_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.2.description", "limit description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.2.rule_type", "LIMIT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.3.cases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.3.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.3.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.3.cases.0.answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.3.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.3.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.3.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.3.default_answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.3.description", "priority description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.3.rule_type", "PRIORITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.4.cases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.4.cases.0.answer_data.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.4.cases.0.answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.4.cases.0.answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.4.cases.0.case_condition", "query.client.address in (subnet '198.51.100.0/24')"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.4.default_answer_data.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.4.default_answer_data.0.answer_condition", "answer.name == 'sampler'"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.4.default_answer_data.0.value", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.4.description", "weighted description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.4.rule_type", "WEIGHTED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "self"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "template", "CUSTOM"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ttl", "11"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + SteeringPolicyResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDnsSteeringPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dnsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_steering_policy" {
			noResourceFound = false
			request := oci_dns.GetSteeringPolicyRequest{}

			tmp := rs.Primary.ID
			request.SteeringPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "dns")

			response, err := client.GetSteeringPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dns.SteeringPolicyLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("DnsSteeringPolicy") {
		resource.AddTestSweepers("DnsSteeringPolicy", &resource.Sweeper{
			Name:         "DnsSteeringPolicy",
			Dependencies: DependencyGraph["steeringPolicy"],
			F:            sweepDnsSteeringPolicyResource,
		})
	}
}

func sweepDnsSteeringPolicyResource(compartment string) error {
	dnsClient := GetTestClients(&schema.ResourceData{}).dnsClient()
	steeringPolicyIds, err := getSteeringPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, steeringPolicyId := range steeringPolicyIds {
		if ok := SweeperDefaultResourceId[steeringPolicyId]; !ok {
			deleteSteeringPolicyRequest := oci_dns.DeleteSteeringPolicyRequest{}

			deleteSteeringPolicyRequest.SteeringPolicyId = &steeringPolicyId

			deleteSteeringPolicyRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "dns")
			_, error := dnsClient.DeleteSteeringPolicy(context.Background(), deleteSteeringPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting SteeringPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", steeringPolicyId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &steeringPolicyId, steeringPolicySweepWaitCondition, time.Duration(3*time.Minute),
				steeringPolicySweepResponseFetchOperation, "dns", true)
		}
	}
	return nil
}

func getSteeringPolicyIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "SteeringPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dnsClient := GetTestClients(&schema.ResourceData{}).dnsClient()

	listSteeringPoliciesRequest := oci_dns.ListSteeringPoliciesRequest{}
	listSteeringPoliciesRequest.CompartmentId = &compartmentId
	listSteeringPoliciesRequest.LifecycleState = oci_dns.SteeringPolicySummaryLifecycleStateActive
	listSteeringPoliciesResponse, err := dnsClient.ListSteeringPolicies(context.Background(), listSteeringPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SteeringPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, steeringPolicy := range listSteeringPoliciesResponse.Items {
		id := *steeringPolicy.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "SteeringPolicyId", id)
	}
	return resourceIds, nil
}

func steeringPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if steeringPolicyResponse, ok := response.Response.(oci_dns.GetSteeringPolicyResponse); ok {
		return steeringPolicyResponse.LifecycleState != oci_dns.SteeringPolicyLifecycleStateDeleted
	}
	return false
}

func steeringPolicySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dnsClient().GetSteeringPolicy(context.Background(), oci_dns.GetSteeringPolicyRequest{
		SteeringPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
