// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsAuthenticationFactorSettingRequiredOnlyResource = IdentityDomainsAuthenticationFactorSettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_authentication_factor_setting", "test_authentication_factor_setting", acctest.Required, acctest.Create, IdentityDomainsAuthenticationFactorSettingRepresentation)

	IdentityDomainsAuthenticationFactorSettingResourceConfig = IdentityDomainsAuthenticationFactorSettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_authentication_factor_setting", "test_authentication_factor_setting", acctest.Optional, acctest.Update, IdentityDomainsAuthenticationFactorSettingRepresentation)

	IdentityDomainsIdentityDomainsAuthenticationFactorSettingSingularDataSourceRepresentation = map[string]interface{}{
		"authentication_factor_setting_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_authentication_factor_setting.test_authentication_factor_setting.id}`},
		"idcs_endpoint":                    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsAuthenticationFactorSettingDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsAuthenticationFactorSettingRepresentation = map[string]interface{}{
		"authentication_factor_setting_id": acctest.Representation{RepType: acctest.Required, Create: `AuthenticationFactorSettings`},
		"bypass_code_enabled":              acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"bypass_code_settings":             acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingBypassCodeSettingsRepresentation},
		"client_app_settings":              acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingClientAppSettingsRepresentation},
		"compliance_policy": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation},
			{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation1},
			{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation2},
			{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation3},
			{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation4},
			{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation5},
			{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation6},
			{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation7},
			{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation8},
			{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation9},
		},
		"endpoint_restrictions":             acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingEndpointRestrictionsRepresentation},
		"idcs_endpoint":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"mfa_enrollment_type":               acctest.Representation{RepType: acctest.Required, Create: `Optional`},
		"notification_settings":             acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingNotificationSettingsRepresentation},
		"push_enabled":                      acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"schemas":                           acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:AuthenticationFactorSettings`}},
		"security_questions_enabled":        acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"sms_enabled":                       acctest.Representation{RepType: acctest.Required, Create: `true`},
		"totp_enabled":                      acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"totp_settings":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsAuthenticationFactorSettingTotpSettingsRepresentation},
		"attribute_sets":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"auto_enroll_email_factor_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"email_enabled":                     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"email_settings":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAuthenticationFactorSettingEmailSettingsRepresentation},
		"fido_authenticator_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"hide_backup_factor_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"identity_store_settings":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAuthenticationFactorSettingIdentityStoreSettingsRepresentation},
		"phone_call_enabled":                acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"tags":                              acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAuthenticationFactorSettingTagsRepresentation},
		"third_party_factor":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAuthenticationFactorSettingThirdPartyFactorRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAuthenticationFactorSettingUrnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettingsRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAuthenticationFactorSettingUrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettingsRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangeForIdentityDomainsAuthenticationFactorSetting},
	}

	ignoreChangeForIdentityDomainsAuthenticationFactorSetting = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings[0].duo_security_settings[0].attestation_key`,
			`schemas`,
		}},
	}
	IdentityDomainsAuthenticationFactorSettingBypassCodeSettingsRepresentation = map[string]interface{}{
		"help_desk_code_expiry_in_mins":   acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"help_desk_generation_enabled":    acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"help_desk_max_usage":             acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"length":                          acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"max_active":                      acctest.Representation{RepType: acctest.Required, Create: `5`, Update: `6`},
		"self_service_generation_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}
	IdentityDomainsAuthenticationFactorSettingClientAppSettingsRepresentation = map[string]interface{}{
		"device_protection_policy":            acctest.Representation{RepType: acctest.Required, Create: `APP_PIN`, Update: `NONE`},
		"initial_lockout_period_in_secs":      acctest.Representation{RepType: acctest.Required, Create: `60`, Update: `30`},
		"key_pair_length":                     acctest.Representation{RepType: acctest.Required, Create: `32`, Update: `2048`},
		"lockout_escalation_pattern":          acctest.Representation{RepType: acctest.Required, Create: `Linear`, Update: `Constant`},
		"max_failures_before_lockout":         acctest.Representation{RepType: acctest.Required, Create: `5`, Update: `10`},
		"max_failures_before_warning":         acctest.Representation{RepType: acctest.Required, Create: `0`, Update: `5`},
		"max_lockout_interval_in_secs":        acctest.Representation{RepType: acctest.Required, Create: `90`, Update: `86400`},
		"min_pin_length":                      acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `6`},
		"policy_update_freq_in_days":          acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"request_signing_algo":                acctest.Representation{RepType: acctest.Required, Create: `SHA256withRSA`, Update: `SHA384withRSA`},
		"shared_secret_encoding":              acctest.Representation{RepType: acctest.Required, Create: `Base32`, Update: `Base64`},
		"unlock_app_for_each_request_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"unlock_app_interval_in_secs":         acctest.Representation{RepType: acctest.Required, Create: `0`, Update: `300`},
		"unlock_on_app_foreground_enabled":    acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"unlock_on_app_start_enabled":         acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `Block`, Update: `Allow`},
		"name":   acctest.Representation{RepType: acctest.Required, Create: `lockScreenRequired`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation1 = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `Block`, Update: `Allow`},
		"name":   acctest.Representation{RepType: acctest.Required, Create: `lockScreenRequiredUnknown`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation2 = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `Block`, Update: `Allow`},
		"name":   acctest.Representation{RepType: acctest.Required, Create: `jailBrokenDevice`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation3 = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `Block`, Update: `Allow`},
		"name":   acctest.Representation{RepType: acctest.Required, Create: `jailBrokenDeviceUnknown`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation4 = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `Block`, Update: `Allow`},
		"name":   acctest.Representation{RepType: acctest.Required, Create: `minWindowsVersion`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `8.1`},
	}
	IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation5 = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `Block`, Update: `Allow`},
		"name":   acctest.Representation{RepType: acctest.Required, Create: `minIosVersion`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `7.1`},
	}
	IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation6 = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `Block`, Update: `Allow`},
		"name":   acctest.Representation{RepType: acctest.Required, Create: `minAndroidVersion`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `4.1`},
	}
	IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation7 = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `Block`, Update: `Allow`},
		"name":   acctest.Representation{RepType: acctest.Required, Create: `minIosAppVersion`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `4.0`},
	}
	IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation8 = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `Block`, Update: `Allow`},
		"name":   acctest.Representation{RepType: acctest.Required, Create: `minAndroidAppVersion`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `8.0`},
	}
	IdentityDomainsAuthenticationFactorSettingCompliancePolicyRepresentation9 = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `Block`, Update: `Allow`},
		"name":   acctest.Representation{RepType: acctest.Required, Create: `minWindowsAppVersion`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
	}
	IdentityDomainsAuthenticationFactorSettingEndpointRestrictionsRepresentation = map[string]interface{}{
		"max_endpoint_trust_duration_in_days": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `180`},
		"max_enrolled_devices":                acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"max_incorrect_attempts":              acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `20`},
		"max_trusted_endpoints":               acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `20`},
		"trusted_endpoints_enabled":           acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}
	IdentityDomainsAuthenticationFactorSettingNotificationSettingsRepresentation = map[string]interface{}{
		"pull_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}
	IdentityDomainsAuthenticationFactorSettingTotpSettingsRepresentation = map[string]interface{}{
		"email_otp_validity_duration_in_mins": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"email_passcode_length":               acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `6`},
		"hashing_algorithm":                   acctest.Representation{RepType: acctest.Required, Create: `SHA1`, Update: `SHA256`},
		"jwt_validity_duration_in_secs":       acctest.Representation{RepType: acctest.Required, Create: `30`, Update: `300`},
		"key_refresh_interval_in_days":        acctest.Representation{RepType: acctest.Required, Create: `30`, Update: `60`},
		"passcode_length":                     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `6`},
		"sms_otp_validity_duration_in_mins":   acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `6`},
		"sms_passcode_length":                 acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `6`},
		"time_step_in_secs":                   acctest.Representation{RepType: acctest.Required, Create: `300`, Update: `30`},
		"time_step_tolerance":                 acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
	}
	IdentityDomainsAuthenticationFactorSettingEmailSettingsRepresentation = map[string]interface{}{
		"email_link_enabled":    acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"email_link_custom_url": acctest.Representation{RepType: acctest.Optional, Create: `emailLinkCustomUrl`, Update: `emailLinkCustomUrl2`},
	}
	IdentityDomainsAuthenticationFactorSettingIdentityStoreSettingsRepresentation = map[string]interface{}{
		"mobile_number_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"mobile_number_update_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsAuthenticationFactorSettingTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsAuthenticationFactorSettingThirdPartyFactorRepresentation = map[string]interface{}{
		"duo_security": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}
	IdentityDomainsAuthenticationFactorSettingUrnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettingsRepresentation = map[string]interface{}{
		"attestation":                                  acctest.Representation{RepType: acctest.Required, Create: `DIRECT`, Update: `NONE`},
		"authenticator_selection_attachment":           acctest.Representation{RepType: acctest.Required, Create: `PLATFORM`, Update: `BOTH`},
		"authenticator_selection_require_resident_key": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"authenticator_selection_resident_key":         acctest.Representation{RepType: acctest.Required, Create: `REQUIRED`, Update: `NONE`},
		"authenticator_selection_user_verification":    acctest.Representation{RepType: acctest.Required, Create: `REQUIRED`, Update: `PREFERRED`},
		"exclude_credentials":                          acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"public_key_types":                             acctest.Representation{RepType: acctest.Required, Create: []string{`RS1`}, Update: []string{`ES256`}},
		"timeout":                                      acctest.Representation{RepType: acctest.Required, Create: `10000`, Update: `60000`},
		"domain_validation_level":                      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}
	IdentityDomainsAuthenticationFactorSettingUrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettingsRepresentation = map[string]interface{}{
		"duo_security_settings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAuthenticationFactorSettingUrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettingsDuoSecuritySettingsRepresentation},
	}
	IdentityDomainsAuthenticationFactorSettingUrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettingsDuoSecuritySettingsRepresentation = map[string]interface{}{
		"api_hostname":           acctest.Representation{RepType: acctest.Required, Create: `apiHostname`, Update: `apiHostname2`},
		"integration_key":        acctest.Representation{RepType: acctest.Required, Create: `integrationKey`, Update: `integrationKey2`},
		"secret_key":             acctest.Representation{RepType: acctest.Required, Create: `secretKey`, Update: `secretKey2`},
		"user_mapping_attribute": acctest.Representation{RepType: acctest.Required, Create: `primaryEmail`, Update: `userName`},
		"attestation_key":        acctest.Representation{RepType: acctest.Optional, Create: `attestationKey`, Update: `attestationKey2`},
	}

	IdentityDomainsAuthenticationFactorSettingResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsAuthenticationFactorSettingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsAuthenticationFactorSettingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_authentication_factor_setting.test_authentication_factor_setting"
	datasourceName := "data.oci_identity_domains_authentication_factor_settings.test_authentication_factor_settings"
	singularDatasourceName := "data.oci_identity_domains_authentication_factor_setting.test_authentication_factor_setting"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsAuthenticationFactorSettingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_authentication_factor_setting", "test_authentication_factor_setting", acctest.Optional, acctest.Create, IdentityDomainsAuthenticationFactorSettingRepresentation), "identitydomains", "authenticationFactorSetting", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create (Create with PUT)
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAuthenticationFactorSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_authentication_factor_setting", "test_authentication_factor_setting", acctest.Required, acctest.Create, IdentityDomainsAuthenticationFactorSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "authentication_factor_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.help_desk_code_expiry_in_mins", "10"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.help_desk_generation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.help_desk_max_usage", "10"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.length", "10"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.max_active", "5"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.self_service_generation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.device_protection_policy", "APP_PIN"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.initial_lockout_period_in_secs", "60"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.key_pair_length", "32"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.lockout_escalation_pattern", "Linear"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.max_failures_before_lockout", "5"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.max_failures_before_warning", "0"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.max_lockout_interval_in_secs", "90"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.min_pin_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.policy_update_freq_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.request_signing_algo", "SHA256withRSA"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.shared_secret_encoding", "Base32"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_app_for_each_request_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_app_interval_in_secs", "0"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_on_app_foreground_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_on_app_start_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.#", "10"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.0.action", "Block"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.0.name", "lockScreenRequired"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.0.value", "true"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_endpoint_trust_duration_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_enrolled_devices", "10"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_incorrect_attempts", "10"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_trusted_endpoints", "10"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.trusted_endpoints_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "mfa_enrollment_type", "Optional"),
				resource.TestCheckResourceAttr(resourceName, "notification_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_settings.0.pull_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "push_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_questions_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "sms_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "totp_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.email_otp_validity_duration_in_mins", "10"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.email_passcode_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.hashing_algorithm", "SHA1"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.jwt_validity_duration_in_secs", "30"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.key_refresh_interval_in_days", "30"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.passcode_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.sms_otp_validity_duration_in_mins", "10"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.sms_passcode_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.time_step_in_secs", "300"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.time_step_tolerance", "2"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					print(resId)
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAuthenticationFactorSettingResourceDependencies,
		},
		// verify Create with optionals (Create with PUT)
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAuthenticationFactorSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_authentication_factor_setting", "test_authentication_factor_setting", acctest.Optional, acctest.Create, IdentityDomainsAuthenticationFactorSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "authentication_factor_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "auto_enroll_email_factor_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.help_desk_code_expiry_in_mins", "10"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.help_desk_generation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.help_desk_max_usage", "10"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.length", "10"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.max_active", "5"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.self_service_generation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.device_protection_policy", "APP_PIN"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.initial_lockout_period_in_secs", "60"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.key_pair_length", "32"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.lockout_escalation_pattern", "Linear"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.max_failures_before_lockout", "5"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.max_failures_before_warning", "0"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.max_lockout_interval_in_secs", "90"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.min_pin_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.policy_update_freq_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.request_signing_algo", "SHA256withRSA"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.shared_secret_encoding", "Base32"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_app_for_each_request_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_app_interval_in_secs", "0"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_on_app_foreground_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_on_app_start_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.#", "10"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.0.action", "Block"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.0.name", "lockScreenRequired"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.0.value", "true"),
				resource.TestCheckResourceAttr(resourceName, "email_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "email_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "email_settings.0.email_link_custom_url", "emailLinkCustomUrl"),
				resource.TestCheckResourceAttr(resourceName, "email_settings.0.email_link_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_endpoint_trust_duration_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_enrolled_devices", "10"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_incorrect_attempts", "10"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_trusted_endpoints", "10"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.trusted_endpoints_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "fido_authenticator_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "hide_backup_factor_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "id", "AuthenticationFactorSettings"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "identity_store_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "identity_store_settings.0.mobile_number_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "identity_store_settings.0.mobile_number_update_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "mfa_enrollment_type", "Optional"),
				resource.TestCheckResourceAttr(resourceName, "notification_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_settings.0.pull_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "phone_call_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "push_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_questions_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "sms_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "third_party_factor.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "third_party_factor.0.duo_security", "false"),
				resource.TestCheckResourceAttr(resourceName, "totp_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.email_otp_validity_duration_in_mins", "10"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.email_passcode_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.hashing_algorithm", "SHA1"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.jwt_validity_duration_in_secs", "30"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.key_refresh_interval_in_days", "30"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.passcode_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.sms_otp_validity_duration_in_mins", "10"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.sms_passcode_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.time_step_in_secs", "300"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.time_step_tolerance", "2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.attestation", "DIRECT"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_attachment", "PLATFORM"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_require_resident_key", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_resident_key", "REQUIRED"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_user_verification", "REQUIRED"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.domain_validation_level", "0"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.exclude_credentials", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.public_key_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.timeout", "10000"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.api_hostname", "apiHostname"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.integration_key", "integrationKey"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.secret_key", "secretKey"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.user_mapping_attribute", "primaryEmail"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "authenticationFactorSettings", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAuthenticationFactorSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_authentication_factor_setting", "test_authentication_factor_setting", acctest.Optional, acctest.Update, IdentityDomainsAuthenticationFactorSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "authentication_factor_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "auto_enroll_email_factor_disabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.help_desk_code_expiry_in_mins", "11"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.help_desk_generation_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.help_desk_max_usage", "11"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.length", "11"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.max_active", "6"),
				resource.TestCheckResourceAttr(resourceName, "bypass_code_settings.0.self_service_generation_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.device_protection_policy", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.initial_lockout_period_in_secs", "30"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.key_pair_length", "2048"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.lockout_escalation_pattern", "Constant"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.max_failures_before_lockout", "10"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.max_failures_before_warning", "5"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.max_lockout_interval_in_secs", "86400"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.min_pin_length", "6"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.policy_update_freq_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.request_signing_algo", "SHA384withRSA"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.shared_secret_encoding", "Base64"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_app_for_each_request_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_app_interval_in_secs", "300"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_on_app_foreground_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "client_app_settings.0.unlock_on_app_start_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.#", "10"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.0.action", "Allow"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.0.name", "lockScreenRequired"),
				resource.TestCheckResourceAttr(resourceName, "compliance_policy.0.value", "false"),
				resource.TestCheckResourceAttr(resourceName, "email_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "email_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "email_settings.0.email_link_custom_url", "emailLinkCustomUrl2"),
				resource.TestCheckResourceAttr(resourceName, "email_settings.0.email_link_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_endpoint_trust_duration_in_days", "180"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_enrolled_devices", "11"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_incorrect_attempts", "20"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.max_trusted_endpoints", "20"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_restrictions.0.trusted_endpoints_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "fido_authenticator_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "hide_backup_factor_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "id", "AuthenticationFactorSettings"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "identity_store_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "identity_store_settings.0.mobile_number_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "identity_store_settings.0.mobile_number_update_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "mfa_enrollment_type", "Optional"),
				resource.TestCheckResourceAttr(resourceName, "notification_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_settings.0.pull_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "phone_call_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "push_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_questions_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "sms_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "third_party_factor.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "third_party_factor.0.duo_security", "true"),
				resource.TestCheckResourceAttr(resourceName, "totp_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.email_otp_validity_duration_in_mins", "11"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.email_passcode_length", "6"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.hashing_algorithm", "SHA256"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.jwt_validity_duration_in_secs", "300"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.key_refresh_interval_in_days", "60"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.passcode_length", "6"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.sms_otp_validity_duration_in_mins", "6"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.sms_passcode_length", "6"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.time_step_in_secs", "30"),
				resource.TestCheckResourceAttr(resourceName, "totp_settings.0.time_step_tolerance", "3"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.attestation", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_attachment", "BOTH"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_require_resident_key", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_resident_key", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_user_verification", "PREFERRED"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.domain_validation_level", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.exclude_credentials", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.public_key_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.timeout", "60000"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.api_hostname", "apiHostname2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.integration_key", "integrationKey2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.secret_key", "secretKey2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.user_mapping_attribute", "userName"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_authentication_factor_settings", "test_authentication_factor_settings", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsAuthenticationFactorSettingDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAuthenticationFactorSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_authentication_factor_setting", "test_authentication_factor_setting", acctest.Optional, acctest.Update, IdentityDomainsAuthenticationFactorSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "authentication_factor_settings.#", "1"),
				resource.TestMatchResourceAttr(datasourceName, "authentication_factor_settings.0.schemas.#", regexp.MustCompile("[1-9]+")),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_authentication_factor_setting", "test_authentication_factor_setting", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsAuthenticationFactorSettingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAuthenticationFactorSettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "authentication_factor_setting_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "auto_enroll_email_factor_disabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bypass_code_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bypass_code_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bypass_code_settings.0.help_desk_code_expiry_in_mins", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bypass_code_settings.0.help_desk_generation_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bypass_code_settings.0.help_desk_max_usage", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bypass_code_settings.0.length", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bypass_code_settings.0.max_active", "6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bypass_code_settings.0.self_service_generation_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.device_protection_policy", "NONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.initial_lockout_period_in_secs", "30"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.key_pair_length", "2048"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.lockout_escalation_pattern", "Constant"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.max_failures_before_lockout", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.max_failures_before_warning", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.max_lockout_interval_in_secs", "86400"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.min_pin_length", "6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.policy_update_freq_in_days", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.request_signing_algo", "SHA384withRSA"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.shared_secret_encoding", "Base64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.unlock_app_for_each_request_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.unlock_app_interval_in_secs", "300"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.unlock_on_app_foreground_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_app_settings.0.unlock_on_app_start_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compliance_policy.#", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compliance_policy.0.action", "Allow"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compliance_policy.0.name", "lockScreenRequired"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compliance_policy.0.value", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "email_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "email_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "email_settings.0.email_link_custom_url", "emailLinkCustomUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "email_settings.0.email_link_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint_restrictions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint_restrictions.0.max_endpoint_trust_duration_in_days", "180"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint_restrictions.0.max_enrolled_devices", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint_restrictions.0.max_incorrect_attempts", "20"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint_restrictions.0.max_trusted_endpoints", "20"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint_restrictions.0.trusted_endpoints_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fido_authenticator_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hide_backup_factor_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "id", "AuthenticationFactorSettings"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identity_store_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identity_store_settings.0.mobile_number_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identity_store_settings.0.mobile_number_update_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mfa_enrollment_type", "Optional"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_settings.0.pull_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phone_call_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "push_enabled", "true"),
				resource.TestMatchResourceAttr(singularDatasourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_questions_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sms_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "third_party_factor.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "third_party_factor.0.duo_security", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_settings.0.email_otp_validity_duration_in_mins", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_settings.0.email_passcode_length", "6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_settings.0.hashing_algorithm", "SHA256"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_settings.0.jwt_validity_duration_in_secs", "300"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_settings.0.key_refresh_interval_in_days", "60"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_settings.0.passcode_length", "6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_settings.0.sms_otp_validity_duration_in_mins", "6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_settings.0.sms_passcode_length", "6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_settings.0.time_step_in_secs", "30"),
				resource.TestCheckResourceAttr(singularDatasourceName, "totp_settings.0.time_step_tolerance", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.attestation", "NONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_attachment", "BOTH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_require_resident_key", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_resident_key", "NONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.authenticator_selection_user_verification", "PREFERRED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.domain_validation_level", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.exclude_credentials", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.public_key_types.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings.0.timeout", "60000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.api_hostname", "apiHostname2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.integration_key", "integrationKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.secret_key", "secretKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings.0.duo_security_settings.0.user_mapping_attribute", "userName"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsAuthenticationFactorSettingRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_authentication_factor_setting", "authenticationFactorSettings"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"idcs_last_upgraded_in_release",
				"tags",
				"authentication_factor_setting_id",
			},
			ResourceName: resourceName,
		},
	})
}
