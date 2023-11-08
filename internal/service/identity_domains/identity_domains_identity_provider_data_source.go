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

func IdentityDomainsIdentityProviderDataSource() *schema.Resource {
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
	fieldMap["identity_provider_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resource_type_schema_version"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsIdentityProviderResource(), fieldMap, readSingularIdentityDomainsIdentityProvider)
}

func readSingularIdentityDomainsIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProviderDataSourceCrud{}
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

type IdentityDomainsIdentityProviderDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetIdentityProviderResponse
}

func (s *IdentityDomainsIdentityProviderDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsIdentityProviderDataSourceCrud) Get() error {
	request := oci_identity_domains.GetIdentityProviderRequest{}

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

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsIdentityProviderDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AssertionAttribute != nil {
		s.D.Set("assertion_attribute", *s.Res.AssertionAttribute)
	}

	s.D.Set("authn_request_binding", s.Res.AuthnRequestBinding)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.CorrelationPolicy != nil {
		s.D.Set("correlation_policy", []interface{}{IdentityProviderCorrelationPolicyToMap(s.Res.CorrelationPolicy)})
	} else {
		s.D.Set("correlation_policy", nil)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.Enabled != nil {
		s.D.Set("enabled", *s.Res.Enabled)
	}

	if s.Res.EncryptionCertificate != nil {
		s.D.Set("encryption_certificate", *s.Res.EncryptionCertificate)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.IconUrl != nil {
		s.D.Set("icon_url", *s.Res.IconUrl)
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

	if s.Res.IdpSsoUrl != nil {
		s.D.Set("idp_sso_url", *s.Res.IdpSsoUrl)
	}

	if s.Res.IncludeSigningCertInSignature != nil {
		s.D.Set("include_signing_cert_in_signature", *s.Res.IncludeSigningCertInSignature)
	}

	jitUserProvAssignedGroups := []interface{}{}
	for _, item := range s.Res.JitUserProvAssignedGroups {
		jitUserProvAssignedGroups = append(jitUserProvAssignedGroups, IdentityProviderJitUserProvAssignedGroupsToMap(item))
	}
	s.D.Set("jit_user_prov_assigned_groups", jitUserProvAssignedGroups)

	if s.Res.JitUserProvAttributeUpdateEnabled != nil {
		s.D.Set("jit_user_prov_attribute_update_enabled", *s.Res.JitUserProvAttributeUpdateEnabled)
	}

	if s.Res.JitUserProvAttributes != nil {
		s.D.Set("jit_user_prov_attributes", []interface{}{IdentityProviderJitUserProvAttributesToMap(s.Res.JitUserProvAttributes)})
	} else {
		s.D.Set("jit_user_prov_attributes", nil)
	}

	if s.Res.JitUserProvCreateUserEnabled != nil {
		s.D.Set("jit_user_prov_create_user_enabled", *s.Res.JitUserProvCreateUserEnabled)
	}

	if s.Res.JitUserProvEnabled != nil {
		s.D.Set("jit_user_prov_enabled", *s.Res.JitUserProvEnabled)
	}

	if s.Res.JitUserProvGroupAssertionAttributeEnabled != nil {
		s.D.Set("jit_user_prov_group_assertion_attribute_enabled", *s.Res.JitUserProvGroupAssertionAttributeEnabled)
	}

	s.D.Set("jit_user_prov_group_assignment_method", s.Res.JitUserProvGroupAssignmentMethod)

	s.D.Set("jit_user_prov_group_mapping_mode", s.Res.JitUserProvGroupMappingMode)

	jitUserProvGroupMappings := []interface{}{}
	for _, item := range s.Res.JitUserProvGroupMappings {
		jitUserProvGroupMappings = append(jitUserProvGroupMappings, IdentityProviderJitUserProvGroupMappingsToMap(item))
	}
	s.D.Set("jit_user_prov_group_mappings", jitUserProvGroupMappings)

	if s.Res.JitUserProvGroupSAMLAttributeName != nil {
		s.D.Set("jit_user_prov_group_saml_attribute_name", *s.Res.JitUserProvGroupSAMLAttributeName)
	}

	if s.Res.JitUserProvGroupStaticListEnabled != nil {
		s.D.Set("jit_user_prov_group_static_list_enabled", *s.Res.JitUserProvGroupStaticListEnabled)
	}

	if s.Res.JitUserProvIgnoreErrorOnAbsentGroups != nil {
		s.D.Set("jit_user_prov_ignore_error_on_absent_groups", *s.Res.JitUserProvIgnoreErrorOnAbsentGroups)
	}

	if s.Res.LastNotificationSentTime != nil {
		s.D.Set("last_notification_sent_time", *s.Res.LastNotificationSentTime)
	}

	s.D.Set("logout_binding", s.Res.LogoutBinding)

	if s.Res.LogoutEnabled != nil {
		s.D.Set("logout_enabled", *s.Res.LogoutEnabled)
	}

	if s.Res.LogoutRequestUrl != nil {
		s.D.Set("logout_request_url", *s.Res.LogoutRequestUrl)
	}

	if s.Res.LogoutResponseUrl != nil {
		s.D.Set("logout_response_url", *s.Res.LogoutResponseUrl)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", *s.Res.Metadata)
	}

	if s.Res.NameIdFormat != nil {
		s.D.Set("name_id_format", *s.Res.NameIdFormat)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PartnerName != nil {
		s.D.Set("partner_name", *s.Res.PartnerName)
	}

	if s.Res.PartnerProviderId != nil {
		s.D.Set("partner_provider_id", *s.Res.PartnerProviderId)
	}

	s.D.Set("requested_authentication_context", s.Res.RequestedAuthenticationContext)

	if s.Res.RequireForceAuthn != nil {
		s.D.Set("require_force_authn", *s.Res.RequireForceAuthn)
	}

	if s.Res.RequiresEncryptedAssertion != nil {
		s.D.Set("requires_encrypted_assertion", *s.Res.RequiresEncryptedAssertion)
	}

	if s.Res.SamlHoKRequired != nil {
		s.D.Set("saml_ho_krequired", *s.Res.SamlHoKRequired)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.ServiceInstanceIdentifier != nil {
		s.D.Set("service_instance_identifier", *s.Res.ServiceInstanceIdentifier)
	}

	if s.Res.ShownOnLoginPage != nil {
		s.D.Set("shown_on_login_page", *s.Res.ShownOnLoginPage)
	}

	s.D.Set("signature_hash_algorithm", s.Res.SignatureHashAlgorithm)

	if s.Res.SigningCertificate != nil {
		s.D.Set("signing_certificate", *s.Res.SigningCertificate)
	}

	if s.Res.SuccinctId != nil {
		s.D.Set("succinct_id", *s.Res.SuccinctId)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.TenantProviderId != nil {
		s.D.Set("tenant_provider_id", *s.Res.TenantProviderId)
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialIdentityProvider != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider", []interface{}{ExtensionSocialIdentityProviderToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialIdentityProvider)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionX509IdentityProvider != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionx509identity_provider", []interface{}{ExtensionX509IdentityProviderToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionX509IdentityProvider)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionx509identity_provider", nil)
	}

	s.D.Set("user_mapping_method", s.Res.UserMappingMethod)

	if s.Res.UserMappingStoreAttribute != nil {
		s.D.Set("user_mapping_store_attribute", *s.Res.UserMappingStoreAttribute)
	}

	return nil
}
