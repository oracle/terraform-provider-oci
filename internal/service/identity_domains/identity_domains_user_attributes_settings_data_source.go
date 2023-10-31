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

func IdentityDomainsUserAttributesSettingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDomainsUserAttributesSettings,
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
			"user_attributes_settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityDomainsUserAttributesSettingDataSource()),
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

func readIdentityDomainsUserAttributesSettings(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsUserAttributesSettingsDataSourceCrud{}
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

type IdentityDomainsUserAttributesSettingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.ListUserAttributesSettingsResponse
}

func (s *IdentityDomainsUserAttributesSettingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsUserAttributesSettingsDataSourceCrud) Get() error {
	request := oci_identity_domains.ListUserAttributesSettingsRequest{}

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

	response, err := s.Client.ListUserAttributesSettings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsUserAttributesSettingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDomainsUserAttributesSettingsDataSource-", IdentityDomainsUserAttributesSettingsDataSource(), s.D))

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, UserAttributesSettingToMap(item))
	}
	s.D.Set("user_attributes_settings", resources)

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

func UserAttributesSettingToMap(obj oci_identity_domains.UserAttributesSetting) map[string]interface{} {
	result := map[string]interface{}{}

	attributeSettings := []interface{}{}
	for _, item := range obj.AttributeSettings {
		attributeSettings = append(attributeSettings, UserAttributesSettingsAttributeSettingsToMap(item))
	}
	result["attribute_settings"] = attributeSettings

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
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

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
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

	return result
}

func UserAttributesSettingsAttributeSettingsToMap(obj oci_identity_domains.UserAttributesSettingsAttributeSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EndUserMutability != nil {
		result["end_user_mutability"] = string(*obj.EndUserMutability)
	}

	result["end_user_mutability_canonical_values"] = obj.EndUserMutabilityCanonicalValues

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
