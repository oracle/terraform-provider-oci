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
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_waf "github.com/oracle/oci-go-sdk/v54/waf"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	WebAppFirewallPolicyRequiredOnlyResource = WebAppFirewallPolicyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", Required, Create, webAppFirewallPolicyRepresentation)

	WebAppFirewallPolicyResourceConfig = WebAppFirewallPolicyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", Optional, Update, webAppFirewallPolicyRepresentation)

	webAppFirewallPolicySingularDataSourceRepresentation = map[string]interface{}{
		"web_app_firewall_policy_id": Representation{RepType: Required, Create: `${oci_waf_web_app_firewall_policy.test_web_app_firewall_policy.id}`},
	}

	webAppFirewallPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":             Representation{RepType: Optional, Create: `${oci_waf_web_app_firewall_policy.test_web_app_firewall_policy.id}`},
		"state":          Representation{RepType: Optional, Create: []string{`ACTIVE`}},
		"filter":         RepresentationGroup{Required, webAppFirewallPolicyDataSourceFilterRepresentation}}
	webAppFirewallPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_waf_web_app_firewall_policy.test_web_app_firewall_policy.id}`}},
	}

	webAppFirewallPolicyRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"actions":        []RepresentationGroup{{Optional, webAppFirewallPolicyActionsRepresentation1}, {Optional, webAppFirewallPolicyActionsRepresentation2}, {Optional, webAppFirewallPolicyActionsRepresentation3}},
		//"defined_tags":            Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"request_access_control":  RepresentationGroup{Optional, webAppFirewallPolicyRequestAccessControlRepresentation},
		"request_protection":      RepresentationGroup{Optional, webAppFirewallPolicyRequestProtectionRepresentation},
		"request_rate_limiting":   RepresentationGroup{Optional, webAppFirewallPolicyRequestRateLimitingRepresentation},
		"response_access_control": RepresentationGroup{Optional, webAppFirewallPolicyResponseAccessControlRepresentation},
		//"response_protection":     RepresentationGroup{Optional, webAppFirewallPolicyResponseProtectionRepresentation}, // can not be created at this point
	}
	webAppFirewallPolicyActionsRepresentation1 = map[string]interface{}{
		"name":    Representation{RepType: Required, Create: `actionName`, Update: `actionName2`},
		"type":    Representation{RepType: Required, Create: `RETURN_HTTP_RESPONSE`, Update: `RETURN_HTTP_RESPONSE`},
		"body":    RepresentationGroup{Optional, webAppFirewallPolicyActionsBodyRepresentation},
		"code":    Representation{RepType: Optional, Create: `400`, Update: `500`},
		"headers": RepresentationGroup{Optional, webAppFirewallPolicyActionsHeadersRepresentation},
	}

	webAppFirewallPolicyActionsRepresentation2 = map[string]interface{}{
		"name": Representation{RepType: Required, Create: `checkAction`},
		"type": Representation{RepType: Required, Create: `CHECK`},
	}

	webAppFirewallPolicyActionsRepresentation3 = map[string]interface{}{
		"name": Representation{RepType: Required, Create: `allowAction`},
		"type": Representation{RepType: Required, Create: `ALLOW`},
	}

	webAppFirewallPolicyRequestAccessControlRepresentation = map[string]interface{}{
		"default_action_name": Representation{RepType: Required, Create: `allowAction`},
		"rules":               RepresentationGroup{Optional, webAppFirewallPolicyRequestAccessControlRulesRepresentation},
	}
	webAppFirewallPolicyRequestProtectionRepresentation = map[string]interface{}{
		"rules": RepresentationGroup{Optional, webAppFirewallPolicyRequestProtectionRulesRepresentation},
	}
	webAppFirewallPolicyRequestRateLimitingRepresentation = map[string]interface{}{
		"rules": RepresentationGroup{Optional, webAppFirewallPolicyRequestRateLimitingRulesRepresentation},
	}
	webAppFirewallPolicyResponseAccessControlRepresentation = map[string]interface{}{
		"rules": RepresentationGroup{Optional, webAppFirewallPolicyResponseAccessControlRulesRepresentation},
	}
	//webAppFirewallPolicyResponseProtectionRepresentation = map[string]interface{}{
	//	"rules": RepresentationGroup{Optional, webAppFirewallPolicyResponseProtectionRulesRepresentation},
	//}
	webAppFirewallPolicyActionsBodyRepresentation = map[string]interface{}{
		"text": Representation{RepType: Required, Create: `text`, Update: `text2`},
		"type": Representation{RepType: Required, Create: `STATIC_TEXT`},
	}
	webAppFirewallPolicyActionsHeadersRepresentation = map[string]interface{}{
		"name":  Representation{RepType: Optional, Create: `name`, Update: `name2`},
		"value": Representation{RepType: Optional, Create: `value`, Update: `value2`},
	}
	webAppFirewallPolicyRequestAccessControlRulesRepresentation = map[string]interface{}{
		"action_name":        Representation{RepType: Required, Create: `actionName`, Update: `actionName2`},
		"name":               Representation{RepType: Required, Create: `name`, Update: `name2`},
		"type":               Representation{RepType: Required, Create: `ACCESS_CONTROL`},
		"condition":          Representation{RepType: Optional, Create: `i_contains(keys(http.request.headers), 'header1')`, Update: `i_contains(keys(http.request.headers), 'header2')`},
		"condition_language": Representation{RepType: Optional, Create: `JMESPATH`},
	}
	webAppFirewallPolicyRequestProtectionRulesRepresentation = map[string]interface{}{
		"action_name":                    Representation{RepType: Required, Create: `actionName`, Update: `actionName2`},
		"name":                           Representation{RepType: Required, Create: `name`, Update: `name2`},
		"protection_capabilities":        RepresentationGroup{Required, webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesRepresentation},
		"type":                           Representation{RepType: Required, Create: `PROTECTION`},
		"condition":                      Representation{RepType: Optional, Create: `i_contains(keys(http.request.headers), 'header1')`, Update: `i_contains(keys(http.request.headers), 'header2')`},
		"condition_language":             Representation{RepType: Optional, Create: `JMESPATH`},
		"protection_capability_settings": RepresentationGroup{Optional, webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitySettingsRepresentation},
	}
	webAppFirewallPolicyRequestRateLimitingRulesRepresentation = map[string]interface{}{
		"action_name":        Representation{RepType: Required, Create: `actionName`, Update: `actionName2`},
		"configurations":     RepresentationGroup{Required, webAppFirewallPolicyRequestRateLimitingRulesConfigurationsRepresentation},
		"name":               Representation{RepType: Required, Create: `name`, Update: `name2`},
		"type":               Representation{RepType: Required, Create: `REQUEST_RATE_LIMITING`},
		"condition":          Representation{RepType: Optional, Create: `i_contains(keys(http.request.headers), 'header1')`, Update: `i_contains(keys(http.request.headers), 'header2')`},
		"condition_language": Representation{RepType: Optional, Create: `JMESPATH`},
	}
	webAppFirewallPolicyResponseAccessControlRulesRepresentation = map[string]interface{}{
		"action_name":        Representation{RepType: Required, Create: `actionName`, Update: `actionName2`},
		"name":               Representation{RepType: Required, Create: `name`, Update: `name2`},
		"type":               Representation{RepType: Required, Create: `ACCESS_CONTROL`},
		"condition":          Representation{RepType: Optional, Create: `i_contains(keys(http.response.headers), 'header1')`, Update: `i_contains(keys(http.response.headers), 'header2')`},
		"condition_language": Representation{RepType: Optional, Create: `JMESPATH`},
	}
	//webAppFirewallPolicyResponseProtectionRulesRepresentation = map[string]interface{}{
	//	"action_name":                    Representation{RepType: Required, Create: `actionName`, Update: `actionName2`},
	//	"name":                           Representation{RepType: Required, Create: `name`, Update: `name2`},
	//	"protection_capabilities":        RepresentationGroup{Required, webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesRepresentation},
	//	"type":                           Representation{RepType: Required, Create: `ACCESS_CONTROL`, Update: `PROTECTION`},
	//	"condition":                      Representation{RepType: Optional, Create: `condition`, Update: `condition2`},
	//	"condition_language":             Representation{RepType: Optional, Create: `JMESPATH`},
	//	"protection_capability_settings": RepresentationGroup{Optional, webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitySettingsRepresentation},
	//}
	webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesRepresentation = map[string]interface{}{
		"key":                            Representation{RepType: Required, Create: `920360`, Update: `920350`},
		"version":                        Representation{RepType: Required, Create: `1`, Update: `2`},
		"collaborative_action_threshold": Representation{RepType: Optional, Create: `10`, Update: `11`},
		"action_name":                    Representation{RepType: Optional, Create: `checkAction`},
		//"collaborative_weights":          RepresentationGroup{Optional, webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesCollaborativeWeightsRepresentation},
		"exclusions": RepresentationGroup{Optional, webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesExclusionsRepresentation},
	}
	webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitySettingsRepresentation = map[string]interface{}{
		"allowed_http_methods":           Representation{RepType: Optional, Create: []string{`GET`}, Update: []string{`POST`}},
		"max_http_request_header_length": Representation{RepType: Optional, Create: `10`, Update: `11`},
		"max_http_request_headers":       Representation{RepType: Optional, Create: `10`, Update: `11`},
		"max_number_of_arguments":        Representation{RepType: Optional, Create: `10`, Update: `11`},
		"max_single_argument_length":     Representation{RepType: Optional, Create: `10`, Update: `11`},
		"max_total_argument_length":      Representation{RepType: Optional, Create: `10`, Update: `11`},
	}
	webAppFirewallPolicyRequestRateLimitingRulesConfigurationsRepresentation = map[string]interface{}{
		"period_in_seconds":          Representation{RepType: Required, Create: `10`, Update: `11`},
		"requests_limit":             Representation{RepType: Required, Create: `10`, Update: `11`},
		"action_duration_in_seconds": Representation{RepType: Optional, Create: `10`, Update: `11`},
	}
	//webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesRepresentation = map[string]interface{}{
	//	"key":                            Representation{RepType: Required, Create: `key`, Update: `key2`},
	//	"version":                        Representation{RepType: Required, Create: `10`, Update: `11`},
	//	"action_name":                    Representation{RepType: Optional, Create: `actionName`, Update: `actionName2`},
	//	"collaborative_action_threshold": Representation{RepType: Optional, Create: `10`, Update: `11`},
	//	"collaborative_weights":          RepresentationGroup{Optional, webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesCollaborativeWeightsRepresentation},
	//	"exclusions":                     RepresentationGroup{Optional, webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesExclusionsRepresentation},
	//}
	//webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitySettingsRepresentation = map[string]interface{}{
	//	"allowed_http_methods":           Representation{RepType: Optional, Create: []string{`GET`}, Update: []string{`POST`}},
	//	"max_http_request_header_length": Representation{RepType: Optional, Create: `10`, Update: `11`},
	//	"max_http_request_headers":       Representation{RepType: Optional, Create: `10`, Update: `11`},
	//	"max_number_of_arguments":        Representation{RepType: Optional, Create: `10`, Update: `11`},
	//	"max_single_argument_length":     Representation{RepType: Optional, Create: `10`, Update: `11`},
	//	"max_total_argument_length":      Representation{RepType: Optional, Create: `10`, Update: `11`},
	//}
	//webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesCollaborativeWeightsRepresentation = map[string]interface{}{
	//	"key":    Representation{RepType: Required, Create: `key`, Update: `key2`},
	//	"weight": Representation{RepType: Required, Create: `10`, Update: `11`},
	//}
	webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesExclusionsRepresentation = map[string]interface{}{
		"args":            Representation{RepType: Optional, Create: []string{`args`}, Update: []string{`args2`}},
		"request_cookies": Representation{RepType: Optional, Create: []string{`requestCookies`}, Update: []string{`requestCookies2`}},
	}
	//webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesCollaborativeWeightsRepresentation = map[string]interface{}{
	//	"key":    Representation{RepType: Required, Create: `key`, Update: `key2`},
	//	"weight": Representation{RepType: Required, Create: `10`, Update: `11`},
	//}
	//webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesExclusionsRepresentation = map[string]interface{}{
	//	"args":            Representation{RepType: Optional, Create: []string{`args`}, Update: []string{`args2`}},
	//	"request_cookies": Representation{RepType: Optional, Create: []string{`requestCookies`}, Update: []string{`requestCookies2`}},
	//}

	WebAppFirewallPolicyResourceDependencies = ""
	//DefinedTagsDependencies
)

// issue-routing-tag: waf/default
func TestWafWebAppFirewallPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWafWebAppFirewallPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waf_web_app_firewall_policy.test_web_app_firewall_policy"
	datasourceName := "data.oci_waf_web_app_firewall_policies.test_web_app_firewall_policies"
	singularDatasourceName := "data.oci_waf_web_app_firewall_policy.test_web_app_firewall_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+WebAppFirewallPolicyResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", Optional, Create, webAppFirewallPolicyRepresentation), "waf", "webAppFirewallPolicy", t)

	ResourceTest(t, testAccCheckWafWebAppFirewallPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + WebAppFirewallPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", Required, Create, webAppFirewallPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + WebAppFirewallPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + WebAppFirewallPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", Optional, Create, webAppFirewallPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.body.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.body.0.text", "text"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.body.0.type", "STATIC_TEXT"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.code", "400"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.headers.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.headers.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.name", "actionName"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "RETURN_HTTP_RESPONSE"),

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.default_action_name", "allowAction"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.action_name", "actionName"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header1')"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.type", "ACCESS_CONTROL"),

				resource.TestCheckResourceAttr(resourceName, "request_protection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.action_name", "actionName"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header1')"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.name", "name"),

				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.0.args.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.0.request_cookies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.key", "920360"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.version", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.allowed_http_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_http_request_header_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_http_request_headers", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_number_of_arguments", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_single_argument_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_total_argument_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.type", "PROTECTION"),

				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.action_name", "actionName"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header1')"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.0.action_duration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.0.period_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.0.requests_limit", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.type", "REQUEST_RATE_LIMITING"),

				resource.TestCheckResourceAttr(resourceName, "response_access_control.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.action_name", "actionName"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.condition", "i_contains(keys(http.response.headers), 'header1')"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.type", "ACCESS_CONTROL"),
				// Can not be created at this point in time since we are missing response protection capabilities
				//resource.TestCheckResourceAttr(resourceName, "response_protection.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.action_name", "actionName"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.condition", "condition"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.condition_language", "JMESPATH"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.name", "name"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.action_name", "actionName"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_action_threshold", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.0.key", "key"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.0.weight", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.0.args.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.0.request_cookies.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.key", "key"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.version", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.allowed_http_methods.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_http_request_header_length", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_http_request_headers", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_number_of_arguments", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_single_argument_length", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_total_argument_length", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.type", "ACCESS_CONTROL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + WebAppFirewallPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", Optional, Create,
					RepresentationCopyWithNewProperties(webAppFirewallPolicyRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.body.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.body.0.text", "text"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.body.0.type", "STATIC_TEXT"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.code", "400"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.headers.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.headers.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.name", "actionName"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "RETURN_HTTP_RESPONSE"),

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				resource.TestCheckResourceAttr(resourceName, "request_access_control.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.default_action_name", "allowAction"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.action_name", "actionName"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header1')"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.type", "ACCESS_CONTROL"),

				resource.TestCheckResourceAttr(resourceName, "request_protection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.action_name", "actionName"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header1')"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.0.args.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.0.request_cookies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.key", "920360"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.version", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.allowed_http_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_http_request_header_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_http_request_headers", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_number_of_arguments", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_single_argument_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_total_argument_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.type", "PROTECTION"),

				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.action_name", "actionName"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header1')"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.0.action_duration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.0.period_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.0.requests_limit", "10"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.type", "REQUEST_RATE_LIMITING"),

				resource.TestCheckResourceAttr(resourceName, "response_access_control.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.action_name", "actionName"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.condition", "i_contains(keys(http.response.headers), 'header1')"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.type", "ACCESS_CONTROL"),

				//resource.TestCheckResourceAttr(resourceName, "response_protection.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.action_name", "actionName"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.condition", "condition"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.condition_language", "JMESPATH"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.name", "name"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.action_name", "actionName"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_action_threshold", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.0.key", "key"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.0.weight", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.0.args.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.0.request_cookies.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.key", "key"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.version", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.allowed_http_methods.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_http_request_header_length", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_http_request_headers", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_number_of_arguments", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_single_argument_length", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_total_argument_length", "10"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.type", "ACCESS_CONTROL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + WebAppFirewallPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", Optional, Update, webAppFirewallPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.body.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.body.0.text", "text2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.body.0.type", "STATIC_TEXT"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.code", "500"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.headers.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.headers.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.name", "actionName2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "RETURN_HTTP_RESPONSE"),

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				resource.TestCheckResourceAttr(resourceName, "request_access_control.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.default_action_name", "allowAction"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.action_name", "actionName2"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header2')"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "request_access_control.0.rules.0.type", "ACCESS_CONTROL"),

				resource.TestCheckResourceAttr(resourceName, "request_protection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.action_name", "actionName2"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header2')"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.action_name", "checkAction"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.0.args.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.0.request_cookies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.key", "920350"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capabilities.0.version", "2"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.allowed_http_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_http_request_header_length", "11"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_http_request_headers", "11"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_number_of_arguments", "11"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_single_argument_length", "11"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_total_argument_length", "11"),
				resource.TestCheckResourceAttr(resourceName, "request_protection.0.rules.0.type", "PROTECTION"),

				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.action_name", "actionName2"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header2')"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.0.action_duration_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.0.period_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.configurations.0.requests_limit", "11"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "request_rate_limiting.0.rules.0.type", "REQUEST_RATE_LIMITING"),

				resource.TestCheckResourceAttr(resourceName, "response_access_control.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.action_name", "actionName2"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.condition", "i_contains(keys(http.response.headers), 'header2')"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "response_access_control.0.rules.0.type", "ACCESS_CONTROL"),

				//resource.TestCheckResourceAttr(resourceName, "response_protection.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.action_name", "actionName2"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.condition", "condition2"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.condition_language", "JMESPATH"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.name", "name2"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.action_name", "actionName2"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_action_threshold", "11"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.0.key", "key2"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.0.weight", "11"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.0.args.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.0.request_cookies.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.key", "key2"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capabilities.0.version", "11"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.allowed_http_methods.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_http_request_header_length", "11"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_http_request_headers", "11"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_number_of_arguments", "11"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_single_argument_length", "11"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_total_argument_length", "11"),
				//resource.TestCheckResourceAttr(resourceName, "response_protection.0.rules.0.type", "PROTECTION"),

				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				GenerateDataSourceFromRepresentationMap("oci_waf_web_app_firewall_policies", "test_web_app_firewall_policies", Optional, Update, webAppFirewallPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + WebAppFirewallPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", Optional, Update, webAppFirewallPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "web_app_firewall_policy_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "web_app_firewall_policy_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", Required, Create, webAppFirewallPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + WebAppFirewallPolicyResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "web_app_firewall_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "actions.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.body.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.body.0.text", "text2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.body.0.type", "STATIC_TEXT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.code", "500"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.headers.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.headers.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.name", "actionName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.type", "RETURN_HTTP_RESPONSE"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "request_access_control.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_access_control.0.default_action_name", "allowAction"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_access_control.0.rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_access_control.0.rules.0.action_name", "actionName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_access_control.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header2')"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_access_control.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_access_control.0.rules.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_access_control.0.rules.0.type", "ACCESS_CONTROL"),

				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.action_name", "actionName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header2')"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capabilities.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capabilities.0.action_name", "checkAction"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.0.args.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capabilities.0.exclusions.0.request_cookies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capabilities.0.key", "920350"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capabilities.0.version", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capability_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capability_settings.0.allowed_http_methods.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_http_request_header_length", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_http_request_headers", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_number_of_arguments", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_single_argument_length", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.protection_capability_settings.0.max_total_argument_length", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_protection.0.rules.0.type", "PROTECTION"),

				resource.TestCheckResourceAttr(singularDatasourceName, "request_rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_rate_limiting.0.rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_rate_limiting.0.rules.0.action_name", "actionName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_rate_limiting.0.rules.0.condition", "i_contains(keys(http.request.headers), 'header2')"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_rate_limiting.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_rate_limiting.0.rules.0.configurations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_rate_limiting.0.rules.0.configurations.0.action_duration_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_rate_limiting.0.rules.0.configurations.0.period_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_rate_limiting.0.rules.0.configurations.0.requests_limit", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_rate_limiting.0.rules.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_rate_limiting.0.rules.0.type", "REQUEST_RATE_LIMITING"),

				resource.TestCheckResourceAttr(singularDatasourceName, "response_access_control.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_access_control.0.rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_access_control.0.rules.0.action_name", "actionName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_access_control.0.rules.0.condition", "i_contains(keys(http.response.headers), 'header2')"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_access_control.0.rules.0.condition_language", "JMESPATH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_access_control.0.rules.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_access_control.0.rules.0.type", "ACCESS_CONTROL"),

				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.action_name", "actionName2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.condition", "condition2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.condition_language", "JMESPATH"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.name", "name2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capabilities.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capabilities.0.action_name", "actionName2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_action_threshold", "11"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.0.key", "key2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capabilities.0.collaborative_weights.0.weight", "11"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.0.args.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capabilities.0.exclusions.0.request_cookies.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capabilities.0.key", "key2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capabilities.0.version", "11"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capability_settings.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capability_settings.0.allowed_http_methods.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_http_request_header_length", "11"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_http_request_headers", "11"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_number_of_arguments", "11"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_single_argument_length", "11"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.protection_capability_settings.0.max_total_argument_length", "11"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "response_protection.0.rules.0.type", "PROTECTION"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + WebAppFirewallPolicyResourceConfig,
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

func testAccCheckWafWebAppFirewallPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).wafClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waf_web_app_firewall_policy" {
			noResourceFound = false
			request := oci_waf.GetWebAppFirewallPolicyRequest{}

			tmp := rs.Primary.ID
			request.WebAppFirewallPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "waf")

			response, err := client.GetWebAppFirewallPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_waf.WebAppFirewallPolicyLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("WafWebAppFirewallPolicy") {
		resource.AddTestSweepers("WafWebAppFirewallPolicy", &resource.Sweeper{
			Name:         "WafWebAppFirewallPolicy",
			Dependencies: DependencyGraph["webAppFirewallPolicy"],
			F:            sweepWafWebAppFirewallPolicyResource,
		})
	}
}

func sweepWafWebAppFirewallPolicyResource(compartment string) error {
	wafClient := GetTestClients(&schema.ResourceData{}).wafClient()
	webAppFirewallPolicyIds, err := getWebAppFirewallPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, webAppFirewallPolicyId := range webAppFirewallPolicyIds {
		if ok := SweeperDefaultResourceId[webAppFirewallPolicyId]; !ok {
			deleteWebAppFirewallPolicyRequest := oci_waf.DeleteWebAppFirewallPolicyRequest{}

			deleteWebAppFirewallPolicyRequest.WebAppFirewallPolicyId = &webAppFirewallPolicyId

			deleteWebAppFirewallPolicyRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "waf")
			_, error := wafClient.DeleteWebAppFirewallPolicy(context.Background(), deleteWebAppFirewallPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting WebAppFirewallPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", webAppFirewallPolicyId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &webAppFirewallPolicyId, webAppFirewallPolicySweepWaitCondition, time.Duration(3*time.Minute),
				webAppFirewallPolicySweepResponseFetchOperation, "waf", true)
		}
	}
	return nil
}

func getWebAppFirewallPolicyIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "WebAppFirewallPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	wafClient := GetTestClients(&schema.ResourceData{}).wafClient()

	listWebAppFirewallPoliciesRequest := oci_waf.ListWebAppFirewallPoliciesRequest{}
	listWebAppFirewallPoliciesRequest.CompartmentId = &compartmentId
	listWebAppFirewallPoliciesRequest.LifecycleState = []oci_waf.WebAppFirewallPolicyLifecycleStateEnum{oci_waf.WebAppFirewallPolicyLifecycleStateActive}
	listWebAppFirewallPoliciesResponse, err := wafClient.ListWebAppFirewallPolicies(context.Background(), listWebAppFirewallPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting WebAppFirewallPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, webAppFirewallPolicy := range listWebAppFirewallPoliciesResponse.Items {
		id := *webAppFirewallPolicy.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "WebAppFirewallPolicyId", id)
	}
	return resourceIds, nil
}

func webAppFirewallPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if webAppFirewallPolicyResponse, ok := response.Response.(oci_waf.GetWebAppFirewallPolicyResponse); ok {
		return webAppFirewallPolicyResponse.LifecycleState != oci_waf.WebAppFirewallPolicyLifecycleStateDeleted
	}
	return false
}

func webAppFirewallPolicySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.wafClient().GetWebAppFirewallPolicy(context.Background(), oci_waf.GetWebAppFirewallPolicyRequest{
		WebAppFirewallPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
