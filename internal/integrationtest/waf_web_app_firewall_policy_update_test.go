// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_waf "github.com/oracle/oci-go-sdk/v65/waf"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

/*
This test checks if update works when some of the optional parameters are absent. It only creates resource and then updates display name.
Persisting absent primitive parameters resulted having golang zero values in terrafom state which caused failures on subsequent updates.
More here: https://jira.oci.oraclecorp.com/browse/TERSI-1623
*/
// issue-routing-tag: waf/default
func TestWafWebAppFirewallPolicyResourceOptionalsUpdate_basic(t *testing.T) {
	var (
		WebAppFirewallPolicyResourceDependencies = ""

		webAppFirewallPolicyStaticTextActionsBodyRepresentation = map[string]interface{}{
			"text": acctest.Representation{RepType: acctest.Required, Create: `text`},
			"type": acctest.Representation{RepType: acctest.Required, Create: `STATIC_TEXT`},
		}

		webAppFirewallPolicyDynamicActionsBodyRepresentation = map[string]interface{}{
			"template": acctest.Representation{RepType: acctest.Required, Create: `template`},
			"type":     acctest.Representation{RepType: acctest.Required, Create: `DYNAMIC`},
		}

		webAppFirewallPolicyActionsHeadersRepresentation = map[string]interface{}{
			"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`},
			"value": acctest.Representation{RepType: acctest.Optional, Create: `value`},
		}

		webAppFirewallPolicyStaticTextActionsRepresentation1 = map[string]interface{}{
			"name":    acctest.Representation{RepType: acctest.Required, Create: `actionName`},
			"type":    acctest.Representation{RepType: acctest.Required, Create: `RETURN_HTTP_RESPONSE`},
			"body":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyStaticTextActionsBodyRepresentation},
			"code":    acctest.Representation{RepType: acctest.Optional, Create: `400`},
			"headers": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyActionsHeadersRepresentation},
		}

		webAppFirewallPolicyDynamicActionsRepresentation1 = map[string]interface{}{
			"name":    acctest.Representation{RepType: acctest.Required, Create: `dynamicActionName`},
			"type":    acctest.Representation{RepType: acctest.Required, Create: `RETURN_HTTP_RESPONSE`},
			"body":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyDynamicActionsBodyRepresentation},
			"code":    acctest.Representation{RepType: acctest.Optional, Create: `401`},
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

		webAppFirewallPolicyRequestAccessControlRulesRepresentation = map[string]interface{}{
			"action_name":        acctest.Representation{RepType: acctest.Required, Create: `actionName`},
			"name":               acctest.Representation{RepType: acctest.Required, Create: `name`},
			"type":               acctest.Representation{RepType: acctest.Required, Create: `ACCESS_CONTROL`},
			"condition":          acctest.Representation{RepType: acctest.Optional, Create: `i_contains(keys(http.request.headers), 'header1')`},
			"condition_language": acctest.Representation{RepType: acctest.Optional, Create: `JMESPATH`},
		}

		webAppFirewallPolicyRequestAccessControlRepresentation = map[string]interface{}{
			"default_action_name": acctest.Representation{RepType: acctest.Required, Create: `allowAction`},
			"rules":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestAccessControlRulesRepresentation},
		}

		webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesExclusionsRepresentation = map[string]interface{}{
			"args":            acctest.Representation{RepType: acctest.Optional, Create: []string{`args`}},
			"request_cookies": acctest.Representation{RepType: acctest.Optional, Create: []string{`requestCookies`}},
		}

		webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesRepresentation = map[string]interface{}{
			"key":        acctest.Representation{RepType: acctest.Required, Create: `920360`},
			"version":    acctest.Representation{RepType: acctest.Required, Create: `1`},
			"exclusions": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesExclusionsRepresentation},
		}

		webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitySettingsRepresentation = map[string]interface{}{
			"allowed_http_methods": acctest.Representation{RepType: acctest.Optional, Create: []string{`GET`}},
		}

		webAppFirewallPolicyRequestProtectionRulesRepresentation = map[string]interface{}{
			"action_name":                    acctest.Representation{RepType: acctest.Required, Create: `actionName`},
			"name":                           acctest.Representation{RepType: acctest.Required, Create: `name`},
			"protection_capabilities":        acctest.RepresentationGroup{RepType: acctest.Required, Group: webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitiesRepresentation},
			"type":                           acctest.Representation{RepType: acctest.Required, Create: `PROTECTION`},
			"condition_language":             acctest.Representation{RepType: acctest.Optional, Create: `JMESPATH`},
			"protection_capability_settings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestProtectionRulesProtectionCapabilitySettingsRepresentation},
		}

		webAppFirewallPolicyRequestProtectionRepresentation = map[string]interface{}{
			"rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestProtectionRulesRepresentation},
		}

		webAppFirewallPolicyRequestRateLimitingRulesConfigurationsRepresentation = map[string]interface{}{
			"period_in_seconds":          acctest.Representation{RepType: acctest.Required, Create: `10`},
			"requests_limit":             acctest.Representation{RepType: acctest.Required, Create: `10`},
			"action_duration_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		}

		webAppFirewallPolicyRequestRateLimitingRulesRepresentation = map[string]interface{}{
			"action_name":        acctest.Representation{RepType: acctest.Required, Create: `actionName`},
			"configurations":     acctest.RepresentationGroup{RepType: acctest.Required, Group: webAppFirewallPolicyRequestRateLimitingRulesConfigurationsRepresentation},
			"name":               acctest.Representation{RepType: acctest.Required, Create: `name`},
			"type":               acctest.Representation{RepType: acctest.Required, Create: `REQUEST_RATE_LIMITING`},
			"condition_language": acctest.Representation{RepType: acctest.Optional, Create: `JMESPATH`},
		}

		webAppFirewallPolicyRequestRateLimitingRepresentation = map[string]interface{}{
			"rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestRateLimitingRulesRepresentation},
		}

		webAppFirewallPolicyResponseAccessControlRulesRepresentation = map[string]interface{}{
			"action_name":        acctest.Representation{RepType: acctest.Required, Create: `actionName`},
			"name":               acctest.Representation{RepType: acctest.Required, Create: `name`},
			"type":               acctest.Representation{RepType: acctest.Required, Create: `ACCESS_CONTROL`},
			"condition":          acctest.Representation{RepType: acctest.Optional, Create: `i_contains(keys(http.response.headers), 'header1')`},
			"condition_language": acctest.Representation{RepType: acctest.Optional, Create: `JMESPATH`},
		}

		webAppFirewallPolicyResponseAccessControlRepresentation = map[string]interface{}{
			"rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyResponseAccessControlRulesRepresentation},
		}

		webAppFirewallPolicyRepresentation = map[string]interface{}{
			"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"actions":                 []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: webAppFirewallPolicyStaticTextActionsRepresentation1}, {RepType: acctest.Optional, Group: webAppFirewallPolicyActionsRepresentation2}, {RepType: acctest.Optional, Group: webAppFirewallPolicyActionsRepresentation3}, {RepType: acctest.Optional, Group: webAppFirewallPolicyDynamicActionsRepresentation1}},
			"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
			"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
			"request_access_control":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestAccessControlRepresentation},
			"request_protection":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestProtectionRepresentation},
			"request_rate_limiting":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyRequestRateLimitingRepresentation},
			"response_access_control": acctest.RepresentationGroup{RepType: acctest.Optional, Group: webAppFirewallPolicyResponseAccessControlRepresentation},
		}
	)

	httpreplay.SetScenario("TestWafWebAppFirewallPolicyResourceOptionalsUpdate_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_waf_web_app_firewall_policy.test_web_app_firewall_policy_update_test"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+WebAppFirewallPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy_update_test", acctest.Optional, acctest.Create, webAppFirewallPolicyRepresentation), "waf", "webAppFirewallPolicy", t)

	acctest.ResourceTest(t, testAccCheckWafWebAppFirewallPolicyDestroyForUpdateTest, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + WebAppFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy_update_test", acctest.Optional, acctest.Create, webAppFirewallPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + WebAppFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy_update_test", acctest.Optional, acctest.Update, webAppFirewallPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

func testAccCheckWafWebAppFirewallPolicyDestroyForUpdateTest(s *terraform.State) error {
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
