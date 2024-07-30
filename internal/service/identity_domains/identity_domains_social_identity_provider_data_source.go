// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsSocialIdentityProviderDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
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
	fieldMap["social_identity_provider_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsSocialIdentityProviderResource(), fieldMap, readSingularIdentityDomainsSocialIdentityProvider)
}

func readSingularIdentityDomainsSocialIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSocialIdentityProviderDataSourceCrud{}
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

type IdentityDomainsSocialIdentityProviderDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetSocialIdentityProviderResponse
}

func (s *IdentityDomainsSocialIdentityProviderDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsSocialIdentityProviderDataSourceCrud) Get() error {
	request := oci_identity_domains.GetSocialIdentityProviderRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if socialIdentityProviderId, ok := s.D.GetOkExists("social_identity_provider_id"); ok {
		tmp := socialIdentityProviderId.(string)
		request.SocialIdentityProviderId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetSocialIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsSocialIdentityProviderDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AccessTokenUrl != nil {
		s.D.Set("access_token_url", *s.Res.AccessTokenUrl)
	}

	if s.Res.AccountLinkingEnabled != nil {
		s.D.Set("account_linking_enabled", *s.Res.AccountLinkingEnabled)
	}

	s.D.Set("admin_scope", s.Res.AdminScope)

	if s.Res.AuthzUrl != nil {
		s.D.Set("authz_url", *s.Res.AuthzUrl)
	}

	if s.Res.AutoRedirectEnabled != nil {
		s.D.Set("auto_redirect_enabled", *s.Res.AutoRedirectEnabled)
	}

	if s.Res.ClientCredentialInPayload != nil {
		s.D.Set("client_credential_in_payload", *s.Res.ClientCredentialInPayload)
	}

	if s.Res.ClockSkewInSeconds != nil {
		s.D.Set("clock_skew_in_seconds", *s.Res.ClockSkewInSeconds)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.ConsumerKey != nil {
		s.D.Set("consumer_key", *s.Res.ConsumerKey)
	}

	if s.Res.ConsumerSecret != nil {
		s.D.Set("consumer_secret", *s.Res.ConsumerSecret)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DiscoveryUrl != nil {
		s.D.Set("discovery_url", *s.Res.DiscoveryUrl)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.Enabled != nil {
		s.D.Set("enabled", *s.Res.Enabled)
	}

	if s.Res.IconUrl != nil {
		s.D.Set("icon_url", *s.Res.IconUrl)
	}

	if s.Res.IdAttribute != nil {
		s.D.Set("id_attribute", *s.Res.IdAttribute)
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

	jitProvAssignedGroups := []interface{}{}
	for _, item := range s.Res.JitProvAssignedGroups {
		jitProvAssignedGroups = append(jitProvAssignedGroups, SocialIdentityProviderJitProvAssignedGroupsToMap(item))
	}
	s.D.Set("jit_prov_assigned_groups", jitProvAssignedGroups)

	if s.Res.JitProvGroupStaticListEnabled != nil {
		s.D.Set("jit_prov_group_static_list_enabled", *s.Res.JitProvGroupStaticListEnabled)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.ProfileUrl != nil {
		s.D.Set("profile_url", *s.Res.ProfileUrl)
	}

	if s.Res.RedirectUrl != nil {
		s.D.Set("redirect_url", *s.Res.RedirectUrl)
	}

	if s.Res.RefreshTokenUrl != nil {
		s.D.Set("refresh_token_url", *s.Res.RefreshTokenUrl)
	}

	if s.Res.RegistrationEnabled != nil {
		s.D.Set("registration_enabled", *s.Res.RegistrationEnabled)
	}

	relayIdpParamMappings := []interface{}{}
	for _, item := range s.Res.RelayIdpParamMappings {
		relayIdpParamMappings = append(relayIdpParamMappings, SocialIdentityProviderRelayIdpParamMappingsToMap(item))
	}
	s.D.Set("relay_idp_param_mappings", relayIdpParamMappings)

	s.D.Set("schemas", s.Res.Schemas)

	s.D.Set("scope", s.Res.Scope)

	if s.Res.ServiceProviderName != nil {
		s.D.Set("service_provider_name", *s.Res.ServiceProviderName)
	}

	if s.Res.ShowOnLogin != nil {
		s.D.Set("show_on_login", *s.Res.ShowOnLogin)
	}

	if s.Res.SocialJitProvisioningEnabled != nil {
		s.D.Set("social_jit_provisioning_enabled", *s.Res.SocialJitProvisioningEnabled)
	}

	s.D.Set("status", s.Res.Status)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	return nil
}
