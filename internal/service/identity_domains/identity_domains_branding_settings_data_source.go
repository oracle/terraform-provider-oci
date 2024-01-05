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

func IdentityDomainsBrandingSettingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDomainsBrandingSettings,
		Schema: map[string]*schema.Schema{
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"branding_settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityDomainsBrandingSettingDataSource()),
			},
			"items_per_page": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"start_index": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_results": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readIdentityDomainsBrandingSettings(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsBrandingSettingsDataSourceCrud{}
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

type IdentityDomainsBrandingSettingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.ListBrandingSettingsResponse
}

func (s *IdentityDomainsBrandingSettingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsBrandingSettingsDataSourceCrud) Get() error {
	request := oci_identity_domains.ListBrandingSettingsRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.ListBrandingSettings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsBrandingSettingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDomainsBrandingSettingsDataSource-", IdentityDomainsBrandingSettingsDataSource(), s.D))

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, BrandingSettingToMap(item))
	}
	s.D.Set("branding_settings", resources)

	if s.Res.ItemsPerPage != nil {
		s.D.Set("items_per_page", *s.Res.ItemsPerPage)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.StartIndex != nil {
		s.D.Set("start_index", *s.Res.StartIndex)
	}

	if s.Res.TotalResults != nil {
		s.D.Set("total_results", *s.Res.TotalResults)
	}

	return nil
}

func BrandingSettingToMap(obj oci_identity_domains.BrandingSetting) map[string]interface{} {
	result := map[string]interface{}{}

	companyNames := []interface{}{}
	for _, item := range obj.CompanyNames {
		companyNames = append(companyNames, BrandingSettingsCompanyNamesToMap(item))
	}
	result["company_names"] = companyNames

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.CustomBranding != nil {
		result["custom_branding"] = bool(*obj.CustomBranding)
	}

	if obj.CustomCssLocation != nil {
		result["custom_css_location"] = string(*obj.CustomCssLocation)
	}

	if obj.CustomHtmlLocation != nil {
		result["custom_html_location"] = string(*obj.CustomHtmlLocation)
	}

	if obj.CustomTranslation != nil {
		result["custom_translation"] = string(*obj.CustomTranslation)
	}

	defaultCompanyNames := []interface{}{}
	for _, item := range obj.DefaultCompanyNames {
		defaultCompanyNames = append(defaultCompanyNames, BrandingSettingsDefaultCompanyNamesToMap(item))
	}
	result["default_company_names"] = defaultCompanyNames

	defaultImages := []interface{}{}
	for _, item := range obj.DefaultImages {
		defaultImages = append(defaultImages, BrandingSettingsDefaultImagesToMap(item))
	}
	result["default_images"] = defaultImages

	defaultLoginTexts := []interface{}{}
	for _, item := range obj.DefaultLoginTexts {
		defaultLoginTexts = append(defaultLoginTexts, BrandingSettingsDefaultLoginTextsToMap(item))
	}
	result["default_login_texts"] = defaultLoginTexts

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.EnableTermsOfUse != nil {
		result["enable_terms_of_use"] = bool(*obj.EnableTermsOfUse)
	}

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IdcsCreatedBy != nil {
		result["idcs_created_by"] = []interface{}{idcsCreatedByToMap(obj.IdcsCreatedBy)}
	}

	if obj.IdcsLastModifiedBy != nil {
		result["idcs_last_modified_by"] = []interface{}{idcsLastModifiedByToMap(obj.IdcsLastModifiedBy)}
	}

	if obj.IdcsLastUpgradedInRelease != nil {
		result["idcs_last_upgraded_in_release"] = string(*obj.IdcsLastUpgradedInRelease)
	}

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	images := []interface{}{}
	for _, item := range obj.Images {
		images = append(images, BrandingSettingsImagesToMap(item))
	}
	result["images"] = images

	if obj.IsHostedPage != nil {
		result["is_hosted_page"] = bool(*obj.IsHostedPage)
	}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	loginTexts := []interface{}{}
	for _, item := range obj.LoginTexts {
		loginTexts = append(loginTexts, BrandingSettingsLoginTextsToMap(item))
	}
	result["login_texts"] = loginTexts

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.PreferredLanguage != nil {
		result["preferred_language"] = string(*obj.PreferredLanguage)
	}

	if obj.PrivacyPolicyUrl != nil {
		result["privacy_policy_url"] = string(*obj.PrivacyPolicyUrl)
	}

	result["schemas"] = obj.Schemas

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.TermsOfUseUrl != nil {
		result["terms_of_use_url"] = string(*obj.TermsOfUseUrl)
	}

	if obj.Timezone != nil {
		result["timezone"] = string(*obj.Timezone)
	}

	return result
}

func BrandingSettingsCompanyNamesToMap(obj oci_identity_domains.BrandingSettingsCompanyNames) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func BrandingSettingsDefaultCompanyNamesToMap(obj oci_identity_domains.BrandingSettingsDefaultCompanyNames) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func BrandingSettingsDefaultImagesToMap(obj oci_identity_domains.BrandingSettingsDefaultImages) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func BrandingSettingsDefaultLoginTextsToMap(obj oci_identity_domains.BrandingSettingsDefaultLoginTexts) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func BrandingSettingsImagesToMap(obj oci_identity_domains.BrandingSettingsImages) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func BrandingSettingsLoginTextsToMap(obj oci_identity_domains.BrandingSettingsLoginTexts) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
