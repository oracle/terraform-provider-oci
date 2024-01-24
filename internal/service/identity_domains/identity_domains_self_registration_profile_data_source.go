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

func IdentityDomainsSelfRegistrationProfileDataSource() *schema.Resource {
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
	fieldMap["self_registration_profile_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsSelfRegistrationProfileResource(), fieldMap, readSingularIdentityDomainsSelfRegistrationProfile)
}

func readSingularIdentityDomainsSelfRegistrationProfile(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSelfRegistrationProfileDataSourceCrud{}
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

type IdentityDomainsSelfRegistrationProfileDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetSelfRegistrationProfileResponse
}

func (s *IdentityDomainsSelfRegistrationProfileDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsSelfRegistrationProfileDataSourceCrud) Get() error {
	request := oci_identity_domains.GetSelfRegistrationProfileRequest{}

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

	if selfRegistrationProfileId, ok := s.D.GetOkExists("self_registration_profile_id"); ok {
		tmp := selfRegistrationProfileId.(string)
		request.SelfRegistrationProfileId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetSelfRegistrationProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsSelfRegistrationProfileDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ActivationEmailRequired != nil {
		s.D.Set("activation_email_required", *s.Res.ActivationEmailRequired)
	}

	if s.Res.Active != nil {
		s.D.Set("active", *s.Res.Active)
	}

	afterSubmitText := []interface{}{}
	for _, item := range s.Res.AfterSubmitText {
		afterSubmitText = append(afterSubmitText, SelfRegistrationProfileAfterSubmitTextToMap(item))
	}
	s.D.Set("after_submit_text", afterSubmitText)

	s.D.Set("allowed_email_domains", s.Res.AllowedEmailDomains)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	consentText := []interface{}{}
	for _, item := range s.Res.ConsentText {
		consentText = append(consentText, SelfRegistrationProfileConsentTextToMap(item))
	}
	s.D.Set("consent_text", consentText)

	if s.Res.ConsentTextPresent != nil {
		s.D.Set("consent_text_present", *s.Res.ConsentTextPresent)
	}

	defaultGroups := []interface{}{}
	for _, item := range s.Res.DefaultGroups {
		defaultGroups = append(defaultGroups, SelfRegistrationProfileDefaultGroupsToMap(item))
	}
	s.D.Set("default_groups", defaultGroups)

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	s.D.Set("disallowed_email_domains", s.Res.DisallowedEmailDomains)

	displayName := []interface{}{}
	for _, item := range s.Res.DisplayName {
		displayName = append(displayName, SelfRegistrationProfileDisplayNameToMap(item))
	}
	s.D.Set("display_name", displayName)

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.EmailTemplate != nil {
		s.D.Set("email_template", []interface{}{SelfRegistrationProfileEmailTemplateToMap(s.Res.EmailTemplate)})
	} else {
		s.D.Set("email_template", nil)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.FooterLogo != nil {
		s.D.Set("footer_logo", *s.Res.FooterLogo)
	}

	footerText := []interface{}{}
	for _, item := range s.Res.FooterText {
		footerText = append(footerText, SelfRegistrationProfileFooterTextToMap(item))
	}
	s.D.Set("footer_text", footerText)

	if s.Res.HeaderLogo != nil {
		s.D.Set("header_logo", *s.Res.HeaderLogo)
	}

	headerText := []interface{}{}
	for _, item := range s.Res.HeaderText {
		headerText = append(headerText, SelfRegistrationProfileHeaderTextToMap(item))
	}
	s.D.Set("header_text", headerText)

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

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.NumberOfDaysRedirectUrlIsValid != nil {
		s.D.Set("number_of_days_redirect_url_is_valid", *s.Res.NumberOfDaysRedirectUrlIsValid)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.RedirectUrl != nil {
		s.D.Set("redirect_url", *s.Res.RedirectUrl)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.ShowOnLoginPage != nil {
		s.D.Set("show_on_login_page", *s.Res.ShowOnLoginPage)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	userAttributes := []interface{}{}
	for _, item := range s.Res.UserAttributes {
		userAttributes = append(userAttributes, SelfRegistrationProfileUserAttributesToMap(item))
	}
	s.D.Set("user_attributes", userAttributes)

	return nil
}
