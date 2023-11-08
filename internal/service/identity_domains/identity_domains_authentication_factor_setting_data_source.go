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

func IdentityDomainsAuthenticationFactorSettingDataSource() *schema.Resource {
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
	fieldMap["authentication_factor_setting_id"] = &schema.Schema{
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
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsAuthenticationFactorSettingResource(), fieldMap, readSingularIdentityDomainsAuthenticationFactorSetting)
}

func readSingularIdentityDomainsAuthenticationFactorSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAuthenticationFactorSettingDataSourceCrud{}
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

type IdentityDomainsAuthenticationFactorSettingDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetAuthenticationFactorSettingResponse
}

func (s *IdentityDomainsAuthenticationFactorSettingDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsAuthenticationFactorSettingDataSourceCrud) Get() error {
	request := oci_identity_domains.GetAuthenticationFactorSettingRequest{}

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

	if authenticationFactorSettingId, ok := s.D.GetOkExists("authentication_factor_setting_id"); ok {
		tmp := authenticationFactorSettingId.(string)
		request.AuthenticationFactorSettingId = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetAuthenticationFactorSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsAuthenticationFactorSettingDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AutoEnrollEmailFactorDisabled != nil {
		s.D.Set("auto_enroll_email_factor_disabled", *s.Res.AutoEnrollEmailFactorDisabled)
	}

	if s.Res.BypassCodeEnabled != nil {
		s.D.Set("bypass_code_enabled", *s.Res.BypassCodeEnabled)
	}

	if s.Res.BypassCodeSettings != nil {
		s.D.Set("bypass_code_settings", []interface{}{AuthenticationFactorSettingsBypassCodeSettingsToMap(s.Res.BypassCodeSettings)})
	} else {
		s.D.Set("bypass_code_settings", nil)
	}

	if s.Res.ClientAppSettings != nil {
		s.D.Set("client_app_settings", []interface{}{AuthenticationFactorSettingsClientAppSettingsToMap(s.Res.ClientAppSettings)})
	} else {
		s.D.Set("client_app_settings", nil)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	compliancePolicy := []interface{}{}
	for _, item := range s.Res.CompliancePolicy {
		compliancePolicy = append(compliancePolicy, AuthenticationFactorSettingsCompliancePolicyToMap(item))
	}
	s.D.Set("compliance_policy", compliancePolicy)

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.EmailEnabled != nil {
		s.D.Set("email_enabled", *s.Res.EmailEnabled)
	}

	if s.Res.EmailSettings != nil {
		s.D.Set("email_settings", []interface{}{AuthenticationFactorSettingsEmailSettingsToMap(s.Res.EmailSettings)})
	} else {
		s.D.Set("email_settings", nil)
	}

	if s.Res.EndpointRestrictions != nil {
		s.D.Set("endpoint_restrictions", []interface{}{AuthenticationFactorSettingsEndpointRestrictionsToMap(s.Res.EndpointRestrictions)})
	} else {
		s.D.Set("endpoint_restrictions", nil)
	}

	if s.Res.FidoAuthenticatorEnabled != nil {
		s.D.Set("fido_authenticator_enabled", *s.Res.FidoAuthenticatorEnabled)
	}

	if s.Res.HideBackupFactorEnabled != nil {
		s.D.Set("hide_backup_factor_enabled", *s.Res.HideBackupFactorEnabled)
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

	if s.Res.IdentityStoreSettings != nil {
		s.D.Set("identity_store_settings", []interface{}{AuthenticationFactorSettingsIdentityStoreSettingsToMap(s.Res.IdentityStoreSettings)})
	} else {
		s.D.Set("identity_store_settings", nil)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.MfaEnabledCategory != nil {
		s.D.Set("mfa_enabled_category", *s.Res.MfaEnabledCategory)
	}

	if s.Res.MfaEnrollmentType != nil {
		s.D.Set("mfa_enrollment_type", *s.Res.MfaEnrollmentType)
	}

	if s.Res.NotificationSettings != nil {
		s.D.Set("notification_settings", []interface{}{AuthenticationFactorSettingsNotificationSettingsToMap(s.Res.NotificationSettings)})
	} else {
		s.D.Set("notification_settings", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PhoneCallEnabled != nil {
		s.D.Set("phone_call_enabled", *s.Res.PhoneCallEnabled)
	}

	if s.Res.PushEnabled != nil {
		s.D.Set("push_enabled", *s.Res.PushEnabled)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.SecurityQuestionsEnabled != nil {
		s.D.Set("security_questions_enabled", *s.Res.SecurityQuestionsEnabled)
	}

	if s.Res.SmsEnabled != nil {
		s.D.Set("sms_enabled", *s.Res.SmsEnabled)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.ThirdPartyFactor != nil {
		s.D.Set("third_party_factor", []interface{}{AuthenticationFactorSettingsThirdPartyFactorToMap(s.Res.ThirdPartyFactor)})
	} else {
		s.D.Set("third_party_factor", nil)
	}

	if s.Res.TotpEnabled != nil {
		s.D.Set("totp_enabled", *s.Res.TotpEnabled)
	}

	if s.Res.TotpSettings != nil {
		s.D.Set("totp_settings", []interface{}{AuthenticationFactorSettingsTotpSettingsToMap(s.Res.TotpSettings)})
	} else {
		s.D.Set("totp_settings", nil)
	}

	if s.Res.UrnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings", []interface{}{ExtensionFidoAuthenticationFactorSettingsToMap(s.Res.UrnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings", nil)
	}

	if s.Res.UrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings", []interface{}{ExtensionThirdPartyAuthenticationFactorSettingsToMap(s.Res.UrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings", nil)
	}

	s.D.Set("user_enrollment_disabled_factors", s.Res.UserEnrollmentDisabledFactors)

	if s.Res.YubicoOtpEnabled != nil {
		s.D.Set("yubico_otp_enabled", *s.Res.YubicoOtpEnabled)
	}

	return nil
}
