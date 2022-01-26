// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_waf "github.com/oracle/oci-go-sdk/v56/waf"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	WebAppFirewallPolicyRequiredOnlyResource = WebAppFirewallPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", acctest.Required, acctest.Create, webAppFirewallPolicyRepresentation)

	WebAppFirewallPolicyResourceConfig = WebAppFirewallPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", acctest.Optional, acctest.Update, webAppFirewallPolicyRepresentation)

	webAppFirewallPolicySingularDataSourceRepresentation = map[string]interface{}{
		"web_app_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waf_web_app_firewall_policy.test_web_app_firewall_policy.id}`},
	}

	webAppFirewallPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_waf_web_app_firewall_policy.test_web_app_firewall_policy.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: webAppFirewallPolicyDataSourceFilterRepresentation}}
	webAppFirewallPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_waf_web_app_firewall_policy.test_web_app_firewall_policy.id}`}},
	}

	webAppFirewallPolicyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"actions":        []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: webAppFirewallPolicyActionsRepresentation1}, {RepType: acctest.Optional, Group: webAppFirewallPolicyActionsRepresentation2}, {RepType: acctest.Optional, Group: webAppFirewallPolicyActionsRepresentation3}},
		//"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"request_access_control":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestAccessControlRepresentation},
		"request_protection":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestProtectionRepresentation},
		"request_rate_limiting":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestRateLimitingRepresentation},
		"response_access_control": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyResponseAccessControlRepresentation},
		//"response_protection":     acctest.RepresentationGroup{RepType: acctest.Optional,Group: webAppFirewallPolicyResponseProtectionRepresentation}, // can not be created at this point
	}
	webAppFirewallPolicyActionsRepresentation1 = map[string]interface{}{
		"name":    acctest.Representation{RepType: acctest.Required, Create: `actionName`, Update: `actionName2`},
		"type":    acctest.Representation{RepType: acctest.Required, Create: `RETURN_HTTP_RESPONSE`, Update: `RETURN_HTTP_RESPONSE`},
		"body":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyActionsBodyRepresentation},
		"code":    acctest.Representation{RepType: acctest.Optional, Create: `400`, Update: `500`},
		"headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyActionsHeadersRepresentation},
	}

	webAppFirewallPolicyActionsRepresentation2 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `checkAction`},
		"type": acctest.Representation{RepType: acctest.Required, Create: `CHECK`},
	}

	webAppFirewallPolicyActionsRepresentation3 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `allowAction`},
		"type": acctest.Representation{RepType: acctest.Required, Create: `ALLOW`},
	}

	webAppFirewallPolicyRequestAccessControlRepresentation = map[string]interface{}{
		"default_action_name": acctest.Representation{RepType: acctest.Required, Create: `allowAction`},
		"rules":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestAccessControlRulesRepresentation},
	}
	webAppFirewallPolicyRequestProtectionRepresentation = map[string]interface{}{
		"rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestProtectionRulesRepresentation},
	}
	webAppFirewallPolicyRequestRateLimitingRepresentation = map[string]interface{}{
		"rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestRateLimitingRulesRepresentation},
	}
	webAppFirewallPolicyResponseAccessControlRepresentation = map[string]interface{}{
		"rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyResponseAccessControlRulesRepresentation},
	}
	//webAppFirewallPolicyResponseProtectionRepresentation = map[string]interface{}{
	//	"rules": acctest.RepresentationGroup{RepType: acctest.Optional,Group: webAppFirewallPolicyResponseProtectionRulesRepresentation},
	//}
	webAppFirewallPolicyActionsBodyRepresentation = map[string]interface{}{
		"text": acctest.Representation{RepType: acctest.Required, Create: `text`, Update: `text2`},
		"type": acctest.Representation{RepType: acctest.Required, Create: `STATIC_TEXT`},
	}
	webAppFirewallPolicyActionsHeadersRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	webAppFirewallPolicyRequestAccessControlRulesRepresentation = map[string]interface{}{
		"action_name":        acctest.Representation{RepType: acctest.Required, Create: `actionName`, Update: `actionName2`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"type":               acctest.Representation{RepType: acctest.Required, Create: `ACCESS_CONTROL`},
		"condition":          acctest.Representation{RepType: acctest.Optional, Create: `i_contains(keys(http.request.headers), 'header1')`, Update: `i_contains(keys(http.request.headers), 'header2')`},
		"condition_language": acctest.Representation{RepType: acctest.Optional, Create: `JMESPATH`},
	}
	webAppFirewallPolicyRequestProtectionRulesRepresentation = map[string]interface{}{
		"action_name":                    acctest.Representation{RepType: acctest.Required, Create: `actionName`, Update: `actionName2`},
		"name":                           acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"protection_capabilities":        acctest.RepresentationGroup{RepType: acctest.Required, Group: webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesRepresentation},
		"type":                           acctest.Representation{RepType: acctest.Required, Create: `PROTECTION`},
		"condition":                      acctest.Representation{RepType: acctest.Optional, Create: `i_contains(keys(http.request.headers), 'header1')`, Update: `i_contains(keys(http.request.headers), 'header2')`},
		"condition_language":             acctest.Representation{RepType: acctest.Optional, Create: `JMESPATH`},
		"protection_capability_settings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitySettingsRepresentation},
	}
	webAppFirewallPolicyRequestRateLimitingRulesRepresentation = map[string]interface{}{
		"action_name":        acctest.Representation{RepType: acctest.Required, Create: `actionName`, Update: `actionName2`},
		"configurations":     acctest.RepresentationGroup{RepType: acctest.Required, Group: webAppFirewallPolicyRequestRateLimitingRulesConfigurationsRepresentation},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"type":               acctest.Representation{RepType: acctest.Required, Create: `REQUEST_RATE_LIMITING`},
		"condition":          acctest.Representation{RepType: acctest.Optional, Create: `i_contains(keys(http.request.headers), 'header1')`, Update: `i_contains(keys(http.request.headers), 'header2')`},
		"condition_language": acctest.Representation{RepType: acctest.Optional, Create: `JMESPATH`},
	}
	webAppFirewallPolicyResponseAccessControlRulesRepresentation = map[string]interface{}{
		"action_name":        acctest.Representation{RepType: acctest.Required, Create: `actionName`, Update: `actionName2`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"type":               acctest.Representation{RepType: acctest.Required, Create: `ACCESS_CONTROL`},
		"condition":          acctest.Representation{RepType: acctest.Optional, Create: `i_contains(keys(http.response.headers), 'header1')`, Update: `i_contains(keys(http.response.headers), 'header2')`},
		"condition_language": acctest.Representation{RepType: acctest.Optional, Create: `JMESPATH`},
	}
	//webAppFirewallPolicyResponseProtectionRulesRepresentation = map[string]interface{}{
	//	"action_name":                    acctest.Representation{RepType: acctest.Required, Create: `actionName`, Update: `actionName2`},
	//	"name":                           acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
	//	"protection_capabilities":        acctest.RepresentationGroup{RepType: acctest.Required,Group: webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesRepresentation},
	//	"type":                           acctest.Representation{RepType: acctest.Required, Create: `ACCESS_CONTROL`, Update: `PROTECTION`},
	//	"condition":                      acctest.Representation{RepType: acctest.Optional, Create: `condition`, Update: `condition2`},
	//	"condition_language":             acctest.Representation{RepType: acctest.Optional, Create: `JMESPATH`},
	//	"protection_capability_settings": acctest.RepresentationGroup{RepType: acctest.Optional,Group: webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitySettingsRepresentation},
	//}
	webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesRepresentation = map[string]interface{}{
		"key":                            acctest.Representation{RepType: acctest.Required, Create: `920360`, Update: `920350`},
		"version":                        acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"collaborative_action_threshold": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"action_name":                    acctest.Representation{RepType: acctest.Optional, Create: `checkAction`},
		//"collaborative_weights":          acctest.RepresentationGroup{RepType: acctest.Optional,Group: webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesCollaborativeWeightsRepresentation},
		"exclusions": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesExclusionsRepresentation},
	}
	webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitySettingsRepresentation = map[string]interface{}{
		"allowed_http_methods":           acctest.Representation{RepType: acctest.Optional, Create: []string{`GET`}, Update: []string{`POST`}},
		"max_http_request_header_length": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_http_request_headers":       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_number_of_arguments":        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_single_argument_length":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_total_argument_length":      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	webAppFirewallPolicyRequestRateLimitingRulesConfigurationsRepresentation = map[string]interface{}{
		"period_in_seconds":          acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"requests_limit":             acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"action_duration_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	//webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesRepresentation = map[string]interface{}{
	//	"key":                            acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
	//	"version":                        acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	//	"action_name":                    acctest.Representation{RepType: acctest.Optional, Create: `actionName`, Update: `actionName2`},
	//	"collaborative_action_threshold": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	//	"collaborative_weights":          acctest.RepresentationGroup{RepType: acctest.Optional,Group: webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesCollaborativeWeightsRepresentation},
	//	"exclusions":                     acctest.RepresentationGroup{RepType: acctest.Optional,Group: webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesExclusionsRepresentation},
	//}
	//webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitySettingsRepresentation = map[string]interface{}{
	//	"allowed_http_methods":           acctest.Representation{RepType: acctest.Optional, Create: []string{`GET`}, Update: []string{`POST`}},
	//	"max_http_request_header_length": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	//	"max_http_request_headers":       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	//	"max_number_of_arguments":        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	//	"max_single_argument_length":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	//	"max_total_argument_length":      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	//}
	//webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesCollaborativeWeightsRepresentation = map[string]interface{}{
	//	"key":    acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
	//	"weight": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	//}
	webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesExclusionsRepresentation = map[string]interface{}{
		"args":            acctest.Representation{RepType: acctest.Optional, Create: []string{`args`}, Update: []string{`args2`}},
		"request_cookies": acctest.Representation{RepType: acctest.Optional, Create: []string{`requestCookies`}, Update: []string{`requestCookies2`}},
	}
	//webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesCollaborativeWeightsRepresentation = map[string]interface{}{
	//	"key":    acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
	//	"weight": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	//}
	//webAppFirewallPolicyResponseProtectionRulesProtectionCapabilitiesExclusionsRepresentation = map[string]interface{}{
	//	"args":            acctest.Representation{RepType: acctest.Optional, Create: []string{`args`}, Update: []string{`args2`}},
	//	"request_cookies": acctest.Representation{RepType: acctest.Optional, Create: []string{`requestCookies`}, Update: []string{`requestCookies2`}},
	//}

	WebAppFirewallPolicyResourceDependencies = ""
	//DefinedTagsDependencies
)

// issue-routing-tag: waf/default
func TestWafWebAppFirewallPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWafWebAppFirewallPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waf_web_app_firewall_policy.test_web_app_firewall_policy"
	datasourceName := "data.oci_waf_web_app_firewall_policies.test_web_app_firewall_policies"
	singularDatasourceName := "data.oci_waf_web_app_firewall_policy.test_web_app_firewall_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+WebAppFirewallPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", acctest.Optional, acctest.Create, webAppFirewallPolicyRepresentation), "waf", "webAppFirewallPolicy", t)

	acctest.ResourceTest(t, testAccCheckWafWebAppFirewallPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + WebAppFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", acctest.Required, acctest.Create, webAppFirewallPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", acctest.Optional, acctest.Create, webAppFirewallPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + WebAppFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(webAppFirewallPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", acctest.Optional, acctest.Update, webAppFirewallPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waf_web_app_firewall_policies", "test_web_app_firewall_policies", acctest.Optional, acctest.Update, webAppFirewallPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + WebAppFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", acctest.Optional, acctest.Update, webAppFirewallPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", acctest.Required, acctest.Create, webAppFirewallPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + WebAppFirewallPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).WafClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waf_web_app_firewall_policy" {
			noResourceFound = false
			request := oci_waf.GetWebAppFirewallPolicyRequest{}

			tmp := rs.Primary.ID
			request.WebAppFirewallPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waf")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("WafWebAppFirewallPolicy") {
		resource.AddTestSweepers("WafWebAppFirewallPolicy", &resource.Sweeper{
			Name:         "WafWebAppFirewallPolicy",
			Dependencies: acctest.DependencyGraph["webAppFirewallPolicy"],
			F:            sweepWafWebAppFirewallPolicyResource,
		})
	}
}

func sweepWafWebAppFirewallPolicyResource(compartment string) error {
	wafClient := acctest.GetTestClients(&schema.ResourceData{}).WafClient()
	webAppFirewallPolicyIds, err := getWebAppFirewallPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, webAppFirewallPolicyId := range webAppFirewallPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[webAppFirewallPolicyId]; !ok {
			deleteWebAppFirewallPolicyRequest := oci_waf.DeleteWebAppFirewallPolicyRequest{}

			deleteWebAppFirewallPolicyRequest.WebAppFirewallPolicyId = &webAppFirewallPolicyId

			deleteWebAppFirewallPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waf")
			_, error := wafClient.DeleteWebAppFirewallPolicy(context.Background(), deleteWebAppFirewallPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting WebAppFirewallPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", webAppFirewallPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &webAppFirewallPolicyId, webAppFirewallPolicySweepWaitCondition, time.Duration(3*time.Minute),
				webAppFirewallPolicySweepResponseFetchOperation, "waf", true)
		}
	}
	return nil
}

func getWebAppFirewallPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WebAppFirewallPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	wafClient := acctest.GetTestClients(&schema.ResourceData{}).WafClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WebAppFirewallPolicyId", id)
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

func webAppFirewallPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.WafClient().GetWebAppFirewallPolicy(context.Background(), oci_waf.GetWebAppFirewallPolicyRequest{
		WebAppFirewallPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
