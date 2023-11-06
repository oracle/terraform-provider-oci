// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsSettingDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["attribute_sets"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	fieldMap["attributes"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["authorization"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["idcs_endpoint"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resource_type_schema_version"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["setting_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsSettingResource(), fieldMap, readSingularIdentityDomainsSetting)
}

func readSingularIdentityDomainsSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSettingDataSourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

type IdentityDomainsSettingDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetSettingResponse
}

func (s *IdentityDomainsSettingDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsSettingDataSourceCrud) Get() error {
	request := oci_identity_domains.GetSettingRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if settingId, ok := s.D.GetOkExists("setting_id"); ok {
		tmp := settingId.(string)
		request.SettingId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsSettingDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AccountAlwaysTrustScope != nil {
		s.D.Set("account_always_trust_scope", *s.Res.AccountAlwaysTrustScope)
	}

	s.D.Set("allowed_domains", s.Res.AllowedDomains)

	s.D.Set("allowed_forgot_password_flow_return_urls", s.Res.AllowedForgotPasswordFlowReturnUrls)

	s.D.Set("allowed_notification_redirect_urls", s.Res.AllowedNotificationRedirectUrls)

	s.D.Set("audit_event_retention_period", s.Res.AuditEventRetentionPeriod)

	if s.Res.CertificateValidation != nil {
		s.D.Set("certificate_validation", []interface{}{SettingsCertificateValidationToMap(s.Res.CertificateValidation)})
	} else {
		s.D.Set("certificate_validation", nil)
	}

	if s.Res.CloudAccountName != nil {
		s.D.Set("cloud_account_name", *s.Res.CloudAccountName)
	}

	if s.Res.CloudGateCorsSettings != nil {
		s.D.Set("cloud_gate_cors_settings", []interface{}{SettingsCloudGateCorsSettingsToMap(s.Res.CloudGateCorsSettings)})
	} else {
		s.D.Set("cloud_gate_cors_settings", nil)
	}

	if s.Res.CloudMigrationCustomUrl != nil {
		s.D.Set("cloud_migration_custom_url", *s.Res.CloudMigrationCustomUrl)
	}

	if s.Res.CloudMigrationUrlEnabled != nil {
		s.D.Set("cloud_migration_url_enabled", *s.Res.CloudMigrationUrlEnabled)
	}

	companyNames := []interface{}{}
	for _, item := range s.Res.CompanyNames {
		companyNames = append(companyNames, SettingsCompanyNamesToMap(item))
	}
	s.D.Set("company_names", companyNames)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	s.D.Set("contact_emails", s.Res.ContactEmails)

	s.D.Set("csr_access", s.Res.CsrAccess)

	if s.Res.CustomBranding != nil {
		s.D.Set("custom_branding", *s.Res.CustomBranding)
	}

	if s.Res.CustomCssLocation != nil {
		s.D.Set("custom_css_location", *s.Res.CustomCssLocation)
	}

	if s.Res.CustomHtmlLocation != nil {
		s.D.Set("custom_html_location", *s.Res.CustomHtmlLocation)
	}

	if s.Res.CustomTranslation != nil {
		s.D.Set("custom_translation", *s.Res.CustomTranslation)
	}

	defaultCompanyNames := []interface{}{}
	for _, item := range s.Res.DefaultCompanyNames {
		defaultCompanyNames = append(defaultCompanyNames, SettingsDefaultCompanyNamesToMap(item))
	}
	s.D.Set("default_company_names", defaultCompanyNames)

	defaultImages := []interface{}{}
	for _, item := range s.Res.DefaultImages {
		defaultImages = append(defaultImages, SettingsDefaultImagesToMap(item))
	}
	s.D.Set("default_images", defaultImages)

	defaultLoginTexts := []interface{}{}
	for _, item := range s.Res.DefaultLoginTexts {
		defaultLoginTexts = append(defaultLoginTexts, SettingsDefaultLoginTextsToMap(item))
	}
	s.D.Set("default_login_texts", defaultLoginTexts)

	s.D.Set("default_trust_scope", s.Res.DefaultTrustScope)

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DiagnosticLevel != nil {
		s.D.Set("diagnostic_level", *s.Res.DiagnosticLevel)
	}

	if s.Res.DiagnosticRecordForSearchIdentifiesReturnedResources != nil {
		s.D.Set("diagnostic_record_for_search_identifies_returned_resources", *s.Res.DiagnosticRecordForSearchIdentifiesReturnedResources)
	}

	if s.Res.DiagnosticTracingUpto != nil {
		s.D.Set("diagnostic_tracing_upto", *s.Res.DiagnosticTracingUpto)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.EnableTermsOfUse != nil {
		s.D.Set("enable_terms_of_use", *s.Res.EnableTermsOfUse)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.IamUpstSessionExpiry != nil {
		s.D.Set("iam_upst_session_expiry", *s.Res.IamUpstSessionExpiry)
	}

	if s.Res.IdcsCreatedBy != nil {
		s.D.Set("idcs_created_by", []interface{}{idcsCreatedByToMap(s.Res.IdcsCreatedBy)})
	} else {
		s.D.Set("idcs_created_by", nil)
	}

	if s.Res.IdcsLastModifiedBy != nil {
		s.D.Set("idcs_last_modified_by", []interface{}{idcsLastModifiedByToMap(s.Res.IdcsLastModifiedBy)})
	} else {
		s.D.Set("idcs_last_modified_by", nil)
	}

	if s.Res.IdcsLastUpgradedInRelease != nil {
		s.D.Set("idcs_last_upgraded_in_release", *s.Res.IdcsLastUpgradedInRelease)
	}

	s.D.Set("idcs_prevented_operations", s.Res.IdcsPreventedOperations)

	images := []interface{}{}
	for _, item := range s.Res.Images {
		images = append(images, SettingsImagesToMap(item))
	}
	s.D.Set("images", images)

	if s.Res.IsHostedPage != nil {
		s.D.Set("is_hosted_page", *s.Res.IsHostedPage)
	}

	if s.Res.Issuer != nil {
		s.D.Set("issuer", *s.Res.Issuer)
	}

	if s.Res.Locale != nil {
		s.D.Set("locale", *s.Res.Locale)
	}

	loginTexts := []interface{}{}
	for _, item := range s.Res.LoginTexts {
		loginTexts = append(loginTexts, SettingsLoginTextsToMap(item))
	}
	s.D.Set("login_texts", loginTexts)

	if s.Res.MaxNoOfAppCMVAToReturn != nil {
		s.D.Set("max_no_of_app_cmva_to_return", *s.Res.MaxNoOfAppCMVAToReturn)
	}

	if s.Res.MaxNoOfAppRoleMembersToReturn != nil {
		s.D.Set("max_no_of_app_role_members_to_return", *s.Res.MaxNoOfAppRoleMembersToReturn)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.MigrationStatus != nil {
		s.D.Set("migration_status", *s.Res.MigrationStatus)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.OnPremisesProvisioning != nil {
		s.D.Set("on_premises_provisioning", *s.Res.OnPremisesProvisioning)
	}

	if s.Res.PreferredLanguage != nil {
		s.D.Set("preferred_language", *s.Res.PreferredLanguage)
	}

	if s.Res.PrevIssuer != nil {
		s.D.Set("prev_issuer", *s.Res.PrevIssuer)
	}

	if s.Res.PrivacyPolicyUrl != nil {
		s.D.Set("privacy_policy_url", *s.Res.PrivacyPolicyUrl)
	}

	purgeConfigs := []interface{}{}
	for _, item := range s.Res.PurgeConfigs {
		purgeConfigs = append(purgeConfigs, SettingsPurgeConfigsToMap(item))
	}
	s.D.Set("purge_configs", purgeConfigs)

	s.D.Set("re_auth_factor", s.Res.ReAuthFactor)

	if s.Res.ReAuthWhenChangingMyAuthenticationFactors != nil {
		s.D.Set("re_auth_when_changing_my_authentication_factors", *s.Res.ReAuthWhenChangingMyAuthenticationFactors)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.ServiceAdminCannotListOtherUsers != nil {
		s.D.Set("service_admin_cannot_list_other_users", *s.Res.ServiceAdminCannotListOtherUsers)
	}

	if s.Res.SigningCertPublicAccess != nil {
		s.D.Set("signing_cert_public_access", *s.Res.SigningCertPublicAccess)
	}

	if s.Res.SubMappingAttr != nil {
		s.D.Set("sub_mapping_attr", *s.Res.SubMappingAttr)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	tenantCustomClaims := []interface{}{}
	for _, item := range s.Res.TenantCustomClaims {
		tenantCustomClaims = append(tenantCustomClaims, SettingsTenantCustomClaimsToMap(item))
	}
	s.D.Set("tenant_custom_claims", tenantCustomClaims)

	if s.Res.TermsOfUseUrl != nil {
		s.D.Set("terms_of_use_url", *s.Res.TermsOfUseUrl)
	}

	if s.Res.Timezone != nil {
		s.D.Set("timezone", *s.Res.Timezone)
	}

	return nil
}
