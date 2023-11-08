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

func IdentityDomainsUserDataSource() *schema.Resource {
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
	fieldMap["user_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsUserResource(), fieldMap, readSingularIdentityDomainsUser)
}

func readSingularIdentityDomainsUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsUserDataSourceCrud{}
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

type IdentityDomainsUserDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetUserResponse
}

func (s *IdentityDomainsUserDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsUserDataSourceCrud) Get() error {
	request := oci_identity_domains.GetUserRequest{}

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

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsUserDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Active != nil {
		s.D.Set("active", *s.Res.Active)
	}

	addresses := []interface{}{}
	for _, item := range s.Res.Addresses {
		addresses = append(addresses, addressesToMap(item))
	}
	s.D.Set("addresses", addresses)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	emails := []interface{}{}
	for _, item := range s.Res.Emails {
		emails = append(emails, UserEmailsToMap(item))
	}
	s.D.Set("emails", emails)

	entitlements := []interface{}{}
	for _, item := range s.Res.Entitlements {
		entitlements = append(entitlements, UserEntitlementsToMap(item))
	}
	s.D.Set("entitlements", entitlements)

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	groups := []interface{}{}
	for _, item := range s.Res.Groups {
		groups = append(groups, UserGroupsToMap(item))
	}
	s.D.Set("groups", groups)

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

	ims := []interface{}{}
	for _, item := range s.Res.Ims {
		ims = append(ims, UserImsToMap(item))
	}
	s.D.Set("ims", ims)

	if s.Res.Locale != nil {
		s.D.Set("locale", *s.Res.Locale)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", []interface{}{UserNameToMap(s.Res.Name)})
	} else {
		s.D.Set("name", nil)
	}

	if s.Res.NickName != nil {
		s.D.Set("nick_name", *s.Res.NickName)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.Password != nil {
		s.D.Set("password", *s.Res.Password)
	}

	phoneNumbers := []interface{}{}
	for _, item := range s.Res.PhoneNumbers {
		phoneNumbers = append(phoneNumbers, UserPhoneNumbersToMap(item))
	}
	s.D.Set("phone_numbers", phoneNumbers)

	photos := []interface{}{}
	for _, item := range s.Res.Photos {
		photos = append(photos, UserPhotosToMap(item))
	}
	s.D.Set("photos", photos)

	if s.Res.PreferredLanguage != nil {
		s.D.Set("preferred_language", *s.Res.PreferredLanguage)
	}

	if s.Res.ProfileUrl != nil {
		s.D.Set("profile_url", *s.Res.ProfileUrl)
	}

	roles := []interface{}{}
	for _, item := range s.Res.Roles {
		roles = append(roles, UserRolesToMap(item))
	}
	s.D.Set("roles", roles)

	s.D.Set("schemas", s.Res.Schemas)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.Timezone != nil {
		s.D.Set("timezone", *s.Res.Timezone)
	}

	if s.Res.Title != nil {
		s.D.Set("title", *s.Res.Title)
	}

	if s.Res.UrnIetfParamsScimSchemasExtensionEnterprise2_0User != nil {
		s.D.Set("urnietfparamsscimschemasextensionenterprise20user", []interface{}{ExtensionEnterprise20UserToMap(s.Res.UrnIetfParamsScimSchemasExtensionEnterprise2_0User)})
	} else {
		s.D.Set("urnietfparamsscimschemasextensionenterprise20user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", []interface{}{ExtensionOCITagsToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionAdaptiveUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionadaptive_user", []interface{}{ExtensionAdaptiveUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionAdaptiveUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionadaptive_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionCapabilitiesUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensioncapabilities_user", []interface{}{ExtensionCapabilitiesUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionCapabilitiesUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensioncapabilities_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbCredentialsUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondb_credentials_user", []interface{}{ExtensionDbCredentialsUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbCredentialsUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondb_credentials_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbUserUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondb_user_user", []interface{}{ExtensionDbUserUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbUserUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondb_user_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosUserUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionkerberos_user_user", []interface{}{ExtensionKerberosUserUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosUserUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionkerberos_user_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionMfaUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmfa_user", []interface{}{ExtensionMfaUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionMfaUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmfa_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordStateUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionpassword_state_user", []interface{}{ExtensionPasswordStateUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordStateUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionpassword_state_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordlessUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionpasswordless_user", []interface{}{ExtensionPasswordlessUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordlessUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionpasswordless_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionposix_user", []interface{}{ExtensionPosixUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionposix_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSecurityQuestionsUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user", []interface{}{ExtensionSecurityQuestionsUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSecurityQuestionsUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionself_change_user", []interface{}{ExtensionSelfChangeUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionself_change_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfRegistrationUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionself_registration_user", []interface{}{ExtensionSelfRegistrationUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfRegistrationUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionself_registration_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSffUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsff_user", []interface{}{ExtensionSffUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSffUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsff_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialAccountUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsocial_account_user", []interface{}{ExtensionSocialAccountUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialAccountUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsocial_account_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionTermsOfUseUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionterms_of_use_user", []interface{}{ExtensionTermsOfUseUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionTermsOfUseUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionterms_of_use_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserCredentialsUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_credentials_user", []interface{}{ExtensionUserCredentialsUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserCredentialsUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_credentials_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserStateUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_state_user", []interface{}{ExtensionUserStateUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserStateUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_state_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_user", []interface{}{ExtensionUserUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_user", nil)
	}

	if s.Res.UserName != nil {
		s.D.Set("user_name", *s.Res.UserName)
	}

	s.D.Set("user_type", s.Res.UserType)

	x509Certificates := []interface{}{}
	for _, item := range s.Res.X509Certificates {
		x509Certificates = append(x509Certificates, UserX509CertificatesToMap(item))
	}
	s.D.Set("x509certificates", x509Certificates)

	return nil
}
