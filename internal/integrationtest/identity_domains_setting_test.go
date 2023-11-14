// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsSettingRequiredOnlyResource = IdentityDomainsSettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_setting", "test_setting", acctest.Required, acctest.Create, IdentityDomainsSettingRepresentation)

	IdentityDomainsSettingResourceConfig = IdentityDomainsSettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_setting", "test_setting", acctest.Optional, acctest.Update, IdentityDomainsSettingRepresentation)

	IdentityDomainsSettingSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"setting_id":     acctest.Representation{RepType: acctest.Required, Create: `Settings`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsSettingDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsSettingRepresentation = map[string]interface{}{
		"csr_access":                 acctest.Representation{RepType: acctest.Required, Create: `readOnly`, Update: `none`},
		"idcs_endpoint":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":                    acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:Settings`}},
		"setting_id":                 acctest.Representation{RepType: acctest.Required, Create: `Settings`},
		"account_always_trust_scope": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"allowed_domains":            acctest.Representation{RepType: acctest.Optional, Create: []string{`test.com`}},
		"allowed_forgot_password_flow_return_urls": acctest.Representation{RepType: acctest.Optional, Create: []string{`allowedForgotPasswordFlowReturnUrls`}, Update: []string{`allowedForgotPasswordFlowReturnUrls2`}},
		"allowed_notification_redirect_urls":       acctest.Representation{RepType: acctest.Optional, Create: []string{`allowedNotificationRedirectUrls`}, Update: []string{`allowedNotificationRedirectUrls2`}},
		"attribute_sets":                           acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"audit_event_retention_period":             acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `60`},
		"certificate_validation":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSettingCertificateValidationRepresentation},
		"cloud_gate_cors_settings":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSettingCloudGateCorsSettingsRepresentation},
		"cloud_migration_custom_url":               acctest.Representation{RepType: acctest.Optional, Create: `cloudMigrationCustomUrl`, Update: `cloudMigrationCustomUrl2`},
		"cloud_migration_url_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"company_names":                            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSettingCompanyNamesRepresentation},
		"contact_emails":                           acctest.Representation{RepType: acctest.Optional, Create: []string{`contactEmails@test.com`}, Update: []string{`contactEmails2@test.com`}},
		"custom_branding":                          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"custom_css_location":                      acctest.Representation{RepType: acctest.Optional, Create: `customCssLocation`, Update: `customCssLocation2`},
		"custom_html_location":                     acctest.Representation{RepType: acctest.Optional, Create: `customHtmlLocation`, Update: `customHtmlLocation2`},
		"custom_translation":                       acctest.Representation{RepType: acctest.Optional, Create: `customTranslation`, Update: `customTranslation2`},
		"default_trust_scope":                      acctest.Representation{RepType: acctest.Optional, Create: `Explicit`, Update: `Account`},
		"diagnostic_level":                         acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"diagnostic_record_for_search_identifies_returned_resources": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"enable_terms_of_use":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"external_id":                          acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"iam_upst_session_expiry":              acctest.Representation{RepType: acctest.Optional, Create: `0`},
		"images":                               acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSettingImagesRepresentation},
		"is_hosted_page":                       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"issuer":                               acctest.Representation{RepType: acctest.Optional, Create: `issuer`, Update: `issuer2`},
		"locale":                               acctest.Representation{RepType: acctest.Optional, Create: `en`, Update: `fr`},
		"login_texts":                          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSettingLoginTextsRepresentation},
		"max_no_of_app_cmva_to_return":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_no_of_app_role_members_to_return": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"preferred_language":                   acctest.Representation{RepType: acctest.Optional, Create: `en`, Update: `fr`},
		"privacy_policy_url":                   acctest.Representation{RepType: acctest.Optional, Create: `privacyPolicyUrl`, Update: `privacyPolicyUrl2`},
		"purge_configs":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSettingPurgeConfigsRepresentation},
		"re_auth_factor":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`password`}},
		"re_auth_when_changing_my_authentication_factors": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"service_admin_cannot_list_other_users":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"signing_cert_public_access":                      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"sub_mapping_attr":                                acctest.Representation{RepType: acctest.Optional, Create: `userName`},
		"tags":                                            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSettingTagsRepresentation},
		"tenant_custom_claims":                            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSettingTenantCustomClaimsRepresentation},
		"terms_of_use_url":                                acctest.Representation{RepType: acctest.Optional, Create: `termsOfUseUrl`, Update: `termsOfUseUrl2`},
		"timezone":                                        acctest.Representation{RepType: acctest.Optional, Create: `America/Los_Angeles`, Update: `America/Vancouver`},
	}
	IdentityDomainsSettingCertificateValidationRepresentation = map[string]interface{}{
		"crl_check_on_ocsp_failure_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"crl_enabled":                           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"crl_location":                          acctest.Representation{RepType: acctest.Optional, Create: `crlLocation`, Update: `crlLocation2`},
		"crl_refresh_interval":                  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"ocsp_enabled":                          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"ocsp_responder_url":                    acctest.Representation{RepType: acctest.Optional, Create: `ocspResponderURL`, Update: `ocspResponderURL2`},
		"ocsp_settings_responder_url_preferred": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"ocsp_signing_certificate_alias":        acctest.Representation{RepType: acctest.Optional, Create: `ocspSigningCertificateAlias`, Update: `ocspSigningCertificateAlias2`},
		"ocsp_timeout_duration":                 acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `9`},
		"ocsp_unknown_response_status_allowed":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsSettingCloudGateCorsSettingsRepresentation = map[string]interface{}{
		"cloud_gate_cors_allow_null_origin": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"cloud_gate_cors_allowed_origins":   acctest.Representation{RepType: acctest.Optional, Create: []string{`https://test.com`}, Update: []string{`https://test2.com`}},
		"cloud_gate_cors_enabled":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"cloud_gate_cors_exposed_headers":   acctest.Representation{RepType: acctest.Optional, Create: []string{`cloudGateCorsExposedHeaders`}, Update: []string{`cloudGateCorsExposedHeaders2`}},
		"cloud_gate_cors_max_age":           acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	IdentityDomainsSettingCompanyNamesRepresentation = map[string]interface{}{
		"locale": acctest.Representation{RepType: acctest.Required, Create: `en`, Update: `fr`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsSettingImagesRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `desktop logo`, Update: `mobile logo`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `https://idcs-guid.identity.oraclecloud.com/oracle-desktop-logo.gif`, Update: `https://idcs-guid.identity.oraclecloud.com/oracle-mobile-logo.gif`},
		"display": acctest.Representation{RepType: acctest.Optional, Create: `display`, Update: `display2`},
	}
	IdentityDomainsSettingLoginTextsRepresentation = map[string]interface{}{
		"locale": acctest.Representation{RepType: acctest.Required, Create: `en`, Update: `fr`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsSettingPurgeConfigsRepresentation = map[string]interface{}{
		"resource_name":    acctest.Representation{RepType: acctest.Required, Create: `resourceName`},
		"retention_period": acctest.Representation{RepType: acctest.Required, Create: `30`, Update: `60`},
	}
	IdentityDomainsSettingTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsSettingTenantCustomClaimsRepresentation = map[string]interface{}{
		"all_scopes": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"expression": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"mode":       acctest.Representation{RepType: acctest.Required, Create: `always`, Update: `request`},
		"name":       acctest.Representation{RepType: acctest.Required, Create: `customClaimName`, Update: `customClaimName2`},
		"token_type": acctest.Representation{RepType: acctest.Required, Create: `AT`, Update: `IT`},
		"value":      acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `$user.test$`},
		// when "all_scopes" is true, "scopes" cannot be set.
		"scopes": acctest.Representation{RepType: acctest.Optional, Create: []string{`scopes`}, Update: []string{}},
	}

	IdentityDomainsSettingResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsSettingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsSettingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_setting.test_setting"
	datasourceName := "data.oci_identity_domains_settings.test_settings"
	singularDatasourceName := "data.oci_identity_domains_setting.test_setting"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsSettingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_setting", "test_setting", acctest.Optional, acctest.Create, IdentityDomainsSettingRepresentation), "identitydomains", "setting", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_setting", "test_setting", acctest.Required, acctest.Create, IdentityDomainsSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "csr_access", "readOnly"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "setting_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSettingResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_setting", "test_setting", acctest.Optional, acctest.Create, IdentityDomainsSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "account_always_trust_scope", "false"),
				resource.TestCheckResourceAttr(resourceName, "allowed_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_forgot_password_flow_return_urls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_notification_redirect_urls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "audit_event_retention_period", "30"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.crl_check_on_ocsp_failure_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.crl_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.crl_location", "crlLocation"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.crl_refresh_interval", "10"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_responder_url", "ocspResponderURL"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_settings_responder_url_preferred", "false"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_signing_certificate_alias", "ocspSigningCertificateAlias"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_timeout_duration", "10"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_unknown_response_status_allowed", "false"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_allow_null_origin", "false"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_exposed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_max_age", "10"),
				resource.TestCheckResourceAttr(resourceName, "cloud_migration_custom_url", "cloudMigrationCustomUrl"),
				resource.TestCheckResourceAttr(resourceName, "cloud_migration_url_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "company_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "company_names.0.locale", "en"),
				resource.TestCheckResourceAttr(resourceName, "company_names.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "contact_emails.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "csr_access", "readOnly"),
				resource.TestCheckResourceAttr(resourceName, "custom_branding", "false"),
				resource.TestCheckResourceAttr(resourceName, "custom_css_location", "customCssLocation"),
				resource.TestCheckResourceAttr(resourceName, "custom_html_location", "customHtmlLocation"),
				resource.TestCheckResourceAttr(resourceName, "custom_translation", "customTranslation"),
				resource.TestCheckResourceAttr(resourceName, "default_trust_scope", "Explicit"),
				resource.TestCheckResourceAttr(resourceName, "diagnostic_level", "0"),
				resource.TestCheckResourceAttr(resourceName, "diagnostic_record_for_search_identifies_returned_resources", "false"),
				resource.TestCheckResourceAttr(resourceName, "enable_terms_of_use", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "iam_upst_session_expiry", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "images.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "images.0.display", "display"),
				resource.TestCheckResourceAttr(resourceName, "images.0.type", "desktop logo"),
				resource.TestCheckResourceAttr(resourceName, "images.0.value", "https://idcs-guid.identity.oraclecloud.com/oracle-desktop-logo.gif"),
				resource.TestCheckResourceAttr(resourceName, "is_hosted_page", "false"),
				resource.TestCheckResourceAttr(resourceName, "issuer", "issuer"),
				resource.TestCheckResourceAttr(resourceName, "locale", "en"),
				resource.TestCheckResourceAttr(resourceName, "login_texts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "login_texts.0.locale", "en"),
				resource.TestCheckResourceAttr(resourceName, "login_texts.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "max_no_of_app_cmva_to_return", "10"),
				resource.TestCheckResourceAttr(resourceName, "max_no_of_app_role_members_to_return", "10"),
				resource.TestCheckResourceAttr(resourceName, "preferred_language", "en"),
				resource.TestCheckResourceAttr(resourceName, "privacy_policy_url", "privacyPolicyUrl"),
				resource.TestCheckResourceAttr(resourceName, "purge_configs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "purge_configs.0.resource_name"),
				resource.TestCheckResourceAttr(resourceName, "purge_configs.0.retention_period", "30"),
				resource.TestCheckResourceAttr(resourceName, "re_auth_factor.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "re_auth_when_changing_my_authentication_factors", "false"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_admin_cannot_list_other_users", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "setting_id"),
				resource.TestCheckResourceAttr(resourceName, "signing_cert_public_access", "false"),
				resource.TestCheckResourceAttr(resourceName, "sub_mapping_attr", "userName"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.all_scopes", "false"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.expression", "false"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.mode", "always"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.name", "customClaimName"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.scopes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.token_type", "AT"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "terms_of_use_url", "termsOfUseUrl"),
				resource.TestCheckResourceAttr(resourceName, "timezone", "America/Los_Angeles"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "settings", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_setting", "test_setting", acctest.Optional, acctest.Update, IdentityDomainsSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "account_always_trust_scope", "true"),
				resource.TestCheckResourceAttr(resourceName, "allowed_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_forgot_password_flow_return_urls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_notification_redirect_urls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "audit_event_retention_period", "60"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.crl_check_on_ocsp_failure_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.crl_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.crl_location", "crlLocation2"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.crl_refresh_interval", "11"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_responder_url", "ocspResponderURL2"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_settings_responder_url_preferred", "true"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_signing_certificate_alias", "ocspSigningCertificateAlias2"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_timeout_duration", "9"),
				resource.TestCheckResourceAttr(resourceName, "certificate_validation.0.ocsp_unknown_response_status_allowed", "true"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_allow_null_origin", "true"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_exposed_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_max_age", "11"),
				resource.TestCheckResourceAttr(resourceName, "cloud_migration_custom_url", "cloudMigrationCustomUrl2"),
				resource.TestCheckResourceAttr(resourceName, "cloud_migration_url_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "company_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "company_names.0.locale", "fr"),
				resource.TestCheckResourceAttr(resourceName, "company_names.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "contact_emails.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "csr_access", "none"),
				resource.TestCheckResourceAttr(resourceName, "custom_branding", "true"),
				resource.TestCheckResourceAttr(resourceName, "custom_css_location", "customCssLocation2"),
				resource.TestCheckResourceAttr(resourceName, "custom_html_location", "customHtmlLocation2"),
				resource.TestCheckResourceAttr(resourceName, "custom_translation", "customTranslation2"),
				resource.TestCheckResourceAttr(resourceName, "default_trust_scope", "Account"),
				resource.TestCheckResourceAttr(resourceName, "diagnostic_level", "1"),
				resource.TestCheckResourceAttr(resourceName, "diagnostic_record_for_search_identifies_returned_resources", "true"),
				resource.TestCheckResourceAttr(resourceName, "enable_terms_of_use", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "iam_upst_session_expiry", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "images.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "images.0.display", "display2"),
				resource.TestCheckResourceAttr(resourceName, "images.0.type", "mobile logo"),
				resource.TestCheckResourceAttr(resourceName, "images.0.value", "https://idcs-guid.identity.oraclecloud.com/oracle-mobile-logo.gif"),
				resource.TestCheckResourceAttr(resourceName, "is_hosted_page", "true"),
				resource.TestCheckResourceAttr(resourceName, "issuer", "issuer2"),
				resource.TestCheckResourceAttr(resourceName, "locale", "fr"),
				resource.TestCheckResourceAttr(resourceName, "login_texts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "login_texts.0.locale", "fr"),
				resource.TestCheckResourceAttr(resourceName, "login_texts.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "max_no_of_app_cmva_to_return", "11"),
				resource.TestCheckResourceAttr(resourceName, "max_no_of_app_role_members_to_return", "11"),
				resource.TestCheckResourceAttr(resourceName, "preferred_language", "fr"),
				resource.TestCheckResourceAttr(resourceName, "privacy_policy_url", "privacyPolicyUrl2"),
				resource.TestCheckResourceAttr(resourceName, "purge_configs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "purge_configs.0.resource_name"),
				resource.TestCheckResourceAttr(resourceName, "purge_configs.0.retention_period", "60"),
				resource.TestCheckResourceAttr(resourceName, "re_auth_factor.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "re_auth_when_changing_my_authentication_factors", "true"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_admin_cannot_list_other_users", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "setting_id"),
				resource.TestCheckResourceAttr(resourceName, "signing_cert_public_access", "true"),
				resource.TestCheckResourceAttr(resourceName, "sub_mapping_attr", "userName"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.all_scopes", "true"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.expression", "true"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.mode", "request"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.name", "customClaimName2"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.scopes.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.token_type", "IT"),
				resource.TestCheckResourceAttr(resourceName, "tenant_custom_claims.0.value", "$user.test$"),
				resource.TestCheckResourceAttr(resourceName, "terms_of_use_url", "termsOfUseUrl2"),
				resource.TestCheckResourceAttr(resourceName, "timezone", "America/Vancouver"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_settings", "test_settings", acctest.Optional, acctest.Update, IdentityDomainsSettingDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_setting", "test_setting", acctest.Optional, acctest.Update, IdentityDomainsSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "settings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "settings.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_setting", "test_setting", acctest.Required, acctest.Create, IdentityDomainsSettingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsSettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "setting_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "account_always_trust_scope", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_domains.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_forgot_password_flow_return_urls.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_notification_redirect_urls.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "audit_event_retention_period", "60"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_validation.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_validation.0.crl_check_on_ocsp_failure_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_validation.0.crl_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_validation.0.crl_location", "crlLocation2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_validation.0.crl_refresh_interval", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_validation.0.ocsp_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_validation.0.ocsp_responder_url", "ocspResponderURL2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_validation.0.ocsp_settings_responder_url_preferred", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_validation.0.ocsp_signing_certificate_alias", "ocspSigningCertificateAlias2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_validation.0.ocsp_timeout_duration", "9"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_validation.0.ocsp_unknown_response_status_allowed", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_gate_cors_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_allow_null_origin", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_allowed_origins.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_exposed_headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_gate_cors_settings.0.cloud_gate_cors_max_age", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_migration_custom_url", "cloudMigrationCustomUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_migration_url_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "company_names.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "company_names.0.locale", "fr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "company_names.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "contact_emails.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "csr_access", "none"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_branding", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_css_location", "customCssLocation2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_html_location", "customHtmlLocation2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_translation", "customTranslation2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "default_trust_scope", "Account"),
				resource.TestCheckResourceAttr(singularDatasourceName, "diagnostic_level", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "diagnostic_record_for_search_identifies_returned_resources", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "enable_terms_of_use", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "images.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "images.0.display", "display2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "images.0.type", "mobile logo"),
				resource.TestCheckResourceAttr(singularDatasourceName, "images.0.value", "https://idcs-guid.identity.oraclecloud.com/oracle-mobile-logo.gif"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_hosted_page", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "issuer", "issuer2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locale", "fr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "login_texts.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "login_texts.0.locale", "fr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "login_texts.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_no_of_app_cmva_to_return", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_no_of_app_role_members_to_return", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "preferred_language", "fr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "privacy_policy_url", "privacyPolicyUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "purge_configs.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "purge_configs.0.retention_period", "60"),
				resource.TestCheckResourceAttr(singularDatasourceName, "re_auth_factor.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "re_auth_when_changing_my_authentication_factors", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_admin_cannot_list_other_users", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "signing_cert_public_access", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sub_mapping_attr", "userName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tenant_custom_claims.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tenant_custom_claims.0.all_scopes", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tenant_custom_claims.0.expression", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tenant_custom_claims.0.mode", "request"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tenant_custom_claims.0.name", "customClaimName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tenant_custom_claims.0.scopes.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tenant_custom_claims.0.token_type", "IT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tenant_custom_claims.0.value", "$user.test$"),
				resource.TestCheckResourceAttr(singularDatasourceName, "terms_of_use_url", "termsOfUseUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timezone", "America/Vancouver"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsSettingRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_setting", "settings"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"idcs_last_upgraded_in_release",
				"tags",
				"setting_id",
				"prev_issuer", // excluded from test because it is not meant to be modified by user
			},
			ResourceName: resourceName,
		},
	})
}
