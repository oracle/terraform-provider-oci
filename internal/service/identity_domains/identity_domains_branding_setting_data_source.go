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

func IdentityDomainsBrandingSettingDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularIdentityDomainsBrandingSetting,
		Schema: map[string]*schema.Schema{
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
			"branding_setting_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"company_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"locale": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_branding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"custom_css_location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_html_location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_translation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_company_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"locale": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"default_images": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"default_login_texts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"locale": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_terms_of_use": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_modified_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_upgraded_in_release": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_prevented_operations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"images": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_hosted_page": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"locale": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login_texts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"locale": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"meta": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"preferred_language": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"privacy_policy_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tenancy_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"terms_of_use_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularIdentityDomainsBrandingSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsBrandingSettingDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityDomainsClient()
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

type IdentityDomainsBrandingSettingDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetBrandingSettingResponse
}

func (s *IdentityDomainsBrandingSettingDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsBrandingSettingDataSourceCrud) Get() error {
	request := oci_identity_domains.GetBrandingSettingRequest{}

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

	if brandingSettingId, ok := s.D.GetOkExists("branding_setting_id"); ok {
		tmp := brandingSettingId.(string)
		request.BrandingSettingId = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetBrandingSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsBrandingSettingDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	companyNames := []interface{}{}
	for _, item := range s.Res.CompanyNames {
		companyNames = append(companyNames, BrandingSettingsCompanyNamesToMap(item))
	}
	s.D.Set("company_names", companyNames)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

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
		defaultCompanyNames = append(defaultCompanyNames, BrandingSettingsDefaultCompanyNamesToMap(item))
	}
	s.D.Set("default_company_names", defaultCompanyNames)

	defaultImages := []interface{}{}
	for _, item := range s.Res.DefaultImages {
		defaultImages = append(defaultImages, BrandingSettingsDefaultImagesToMap(item))
	}
	s.D.Set("default_images", defaultImages)

	defaultLoginTexts := []interface{}{}
	for _, item := range s.Res.DefaultLoginTexts {
		defaultLoginTexts = append(defaultLoginTexts, BrandingSettingsDefaultLoginTextsToMap(item))
	}
	s.D.Set("default_login_texts", defaultLoginTexts)

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
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
		images = append(images, BrandingSettingsImagesToMap(item))
	}
	s.D.Set("images", images)

	if s.Res.IsHostedPage != nil {
		s.D.Set("is_hosted_page", *s.Res.IsHostedPage)
	}

	if s.Res.Locale != nil {
		s.D.Set("locale", *s.Res.Locale)
	}

	loginTexts := []interface{}{}
	for _, item := range s.Res.LoginTexts {
		loginTexts = append(loginTexts, BrandingSettingsLoginTextsToMap(item))
	}
	s.D.Set("login_texts", loginTexts)

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PreferredLanguage != nil {
		s.D.Set("preferred_language", *s.Res.PreferredLanguage)
	}

	if s.Res.PrivacyPolicyUrl != nil {
		s.D.Set("privacy_policy_url", *s.Res.PrivacyPolicyUrl)
	}

	s.D.Set("schemas", s.Res.Schemas)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.TermsOfUseUrl != nil {
		s.D.Set("terms_of_use_url", *s.Res.TermsOfUseUrl)
	}

	if s.Res.Timezone != nil {
		s.D.Set("timezone", *s.Res.Timezone)
	}

	return nil
}
