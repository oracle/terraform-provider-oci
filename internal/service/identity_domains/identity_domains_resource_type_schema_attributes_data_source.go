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

func IdentityDomainsResourceTypeSchemaAttributesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDomainsResourceTypeSchemaAttributes,
		Schema: map[string]*schema.Schema{
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type_schema_attribute_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"resource_type_schema_attribute_filter": {
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
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_order": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Computed
			"resource_type_schema_attributes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"canonical_values": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"case_exact": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"compartment_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"delete_in_progress": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_user_mutability": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_user_mutability_allowed_values": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_added_since_release_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_added_since_version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"idcs_attribute_cacheable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_attribute_mappable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_auditable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_auto_increment_seq_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_canonical_value_source_filter": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_canonical_value_source_resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_composite_key": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"idcs_created_by": {
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
									"ocid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ref": {
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
						"idcs_csv_column_header_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_custom_attribute": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_deprecated_since_release_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_deprecated_since_version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"idcs_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_display_name_message_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_fetch_complex_attribute_values": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_from_target_mapper": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_fully_qualified_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_generated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_icf_attribute_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_icf_bundle_attribute_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_icf_required": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_indirect_ref_resource_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"idcs_internal": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_last_modified_by": {
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
									"ocid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ref": {
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
						"idcs_max_length": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"idcs_max_value": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"idcs_min_length": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"idcs_min_value": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"idcs_multi_language": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_prevented_operations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"idcs_ref_resource_attribute": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_ref_resource_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"idcs_schema_urn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_scim_compliant": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_searchable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_sensitive": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_target_attribute_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_target_attribute_name_to_migrate_from": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_target_norm_attribute_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_target_unique_constraint_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_to_target_mapper": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_trim_string_value": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_validate_reference": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"idcs_value_persisted": {
							Type:     schema.TypeBool,
							Computed: true,
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
						"multi_valued": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mutability": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reference_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"required": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"returned": {
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
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uniqueness": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"total_results": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readIdentityDomainsResourceTypeSchemaAttributes(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsResourceTypeSchemaAttributesDataSourceCrud{}
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

type IdentityDomainsResourceTypeSchemaAttributesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.ListResourceTypeSchemaAttributesResponse
}

func (s *IdentityDomainsResourceTypeSchemaAttributesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsResourceTypeSchemaAttributesDataSourceCrud) Get() error {
	request := oci_identity_domains.ListResourceTypeSchemaAttributesRequest{}

	if resourceTypeSchemaAttributeCount, ok := s.D.GetOkExists("resource_type_schema_attribute_count"); ok {
		tmp := resourceTypeSchemaAttributeCount.(int)
		request.Count = &tmp
	}

	if resourceTypeSchemaAttributeFilter, ok := s.D.GetOkExists("resource_type_schema_attribute_filter"); ok {
		tmp := resourceTypeSchemaAttributeFilter.(string)
		request.Filter = &tmp
	}

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

	if startIndex, ok := s.D.GetOkExists("start_index"); ok {
		tmp := startIndex.(int)
		request.StartIndex = &tmp
	}

	if sortOrder, ok := s.D.GetOkExists("sort_order"); ok {
		tmp := oci_identity_domains.ListResourceTypeSchemaAttributesSortOrderEnum(sortOrder.(string))
		request.SortOrder = tmp
	}

	if sortBy, ok := s.D.GetOkExists("sort_by"); ok {
		tmp := sortBy.(string)
		request.SortBy = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.ListResourceTypeSchemaAttributes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	// IDCS pagination
	startIndex := *response.StartIndex
	for startIndex+*response.ItemsPerPage <= *response.TotalResults {
		startIndex += *response.ItemsPerPage
		request.StartIndex = &startIndex
		listResponse, err := s.Client.ListResourceTypeSchemaAttributes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Resources = append(s.Res.Resources, listResponse.Resources...)
	}

	return nil
}

func (s *IdentityDomainsResourceTypeSchemaAttributesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDomainsResourceTypeSchemaAttributesDataSource-", IdentityDomainsResourceTypeSchemaAttributesDataSource(), s.D))

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, ResourceTypeSchemaAttributeToMap(item))
	}
	s.D.Set("resource_type_schema_attributes", resources)

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

func ResourceTypeSchemaAttributeToMap(obj oci_identity_domains.ResourceTypeSchemaAttribute) map[string]interface{} {
	result := map[string]interface{}{}

	result["canonical_values"] = obj.CanonicalValues

	if obj.CaseExact != nil {
		result["case_exact"] = bool(*obj.CaseExact)
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	result["end_user_mutability"] = string(obj.EndUserMutability)

	result["end_user_mutability_allowed_values"] = obj.EndUserMutabilityAllowedValues

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IdcsAddedSinceReleaseNumber != nil {
		result["idcs_added_since_release_number"] = string(*obj.IdcsAddedSinceReleaseNumber)
	}

	if obj.IdcsAddedSinceVersion != nil {
		result["idcs_added_since_version"] = int(*obj.IdcsAddedSinceVersion)
	}

	if obj.IdcsAttributeCacheable != nil {
		result["idcs_attribute_cacheable"] = bool(*obj.IdcsAttributeCacheable)
	}

	if obj.IdcsAttributeMappable != nil {
		result["idcs_attribute_mappable"] = bool(*obj.IdcsAttributeMappable)
	}

	if obj.IdcsAuditable != nil {
		result["idcs_auditable"] = bool(*obj.IdcsAuditable)
	}

	if obj.IdcsAutoIncrementSeqName != nil {
		result["idcs_auto_increment_seq_name"] = string(*obj.IdcsAutoIncrementSeqName)
	}

	if obj.IdcsCanonicalValueSourceFilter != nil {
		result["idcs_canonical_value_source_filter"] = string(*obj.IdcsCanonicalValueSourceFilter)
	}

	if obj.IdcsCanonicalValueSourceResourceType != nil {
		result["idcs_canonical_value_source_resource_type"] = string(*obj.IdcsCanonicalValueSourceResourceType)
	}

	result["idcs_composite_key"] = obj.IdcsCompositeKey

	if obj.IdcsCreatedBy != nil {
		result["idcs_created_by"] = []interface{}{idcsCreatedByToMap(obj.IdcsCreatedBy)}
	}

	if obj.IdcsCsvColumnHeaderName != nil {
		result["idcs_csv_column_header_name"] = string(*obj.IdcsCsvColumnHeaderName)
	}

	if obj.IdcsCustomAttribute != nil {
		result["idcs_custom_attribute"] = bool(*obj.IdcsCustomAttribute)
	}

	if obj.IdcsDeprecatedSinceReleaseNumber != nil {
		result["idcs_deprecated_since_release_number"] = string(*obj.IdcsDeprecatedSinceReleaseNumber)
	}

	if obj.IdcsDeprecatedSinceVersion != nil {
		result["idcs_deprecated_since_version"] = int(*obj.IdcsDeprecatedSinceVersion)
	}

	if obj.IdcsDisplayName != nil {
		result["idcs_display_name"] = string(*obj.IdcsDisplayName)
	}

	if obj.IdcsDisplayNameMessageId != nil {
		result["idcs_display_name_message_id"] = string(*obj.IdcsDisplayNameMessageId)
	}

	if obj.IdcsFetchComplexAttributeValues != nil {
		result["idcs_fetch_complex_attribute_values"] = bool(*obj.IdcsFetchComplexAttributeValues)
	}

	if obj.IdcsFromTargetMapper != nil {
		result["idcs_from_target_mapper"] = string(*obj.IdcsFromTargetMapper)
	}

	if obj.IdcsFullyQualifiedName != nil {
		result["idcs_fully_qualified_name"] = string(*obj.IdcsFullyQualifiedName)
	}

	if obj.IdcsGenerated != nil {
		result["idcs_generated"] = bool(*obj.IdcsGenerated)
	}

	result["idcs_icf_attribute_type"] = string(obj.IdcsICFAttributeType)

	if obj.IdcsICFBundleAttributeName != nil {
		result["idcs_icf_bundle_attribute_name"] = string(*obj.IdcsICFBundleAttributeName)
	}

	if obj.IdcsICFRequired != nil {
		result["idcs_icf_required"] = bool(*obj.IdcsICFRequired)
	}

	result["idcs_indirect_ref_resource_attributes"] = obj.IdcsIndirectRefResourceAttributes

	if obj.IdcsInternal != nil {
		result["idcs_internal"] = bool(*obj.IdcsInternal)
	}

	if obj.IdcsLastModifiedBy != nil {
		result["idcs_last_modified_by"] = []interface{}{idcsLastModifiedByToMap(obj.IdcsLastModifiedBy)}
	}

	if obj.IdcsLastUpgradedInRelease != nil {
		result["idcs_last_upgraded_in_release"] = string(*obj.IdcsLastUpgradedInRelease)
	}

	if obj.IdcsMaxLength != nil {
		result["idcs_max_length"] = int(*obj.IdcsMaxLength)
	}

	if obj.IdcsMaxValue != nil {
		result["idcs_max_value"] = int(*obj.IdcsMaxValue)
	}

	if obj.IdcsMinLength != nil {
		result["idcs_min_length"] = int(*obj.IdcsMinLength)
	}

	if obj.IdcsMinValue != nil {
		result["idcs_min_value"] = int(*obj.IdcsMinValue)
	}

	if obj.IdcsMultiLanguage != nil {
		result["idcs_multi_language"] = bool(*obj.IdcsMultiLanguage)
	}

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	if obj.IdcsRefResourceAttribute != nil {
		result["idcs_ref_resource_attribute"] = string(*obj.IdcsRefResourceAttribute)
	}

	result["idcs_ref_resource_attributes"] = obj.IdcsRefResourceAttributes

	if obj.IdcsSchemaUrn != nil {
		result["idcs_schema_urn"] = string(*obj.IdcsSchemaUrn)
	}

	if obj.IdcsScimCompliant != nil {
		result["idcs_scim_compliant"] = bool(*obj.IdcsScimCompliant)
	}

	if obj.IdcsSearchable != nil {
		result["idcs_searchable"] = bool(*obj.IdcsSearchable)
	}

	result["idcs_sensitive"] = string(obj.IdcsSensitive)

	if obj.IdcsTargetAttributeName != nil {
		result["idcs_target_attribute_name"] = string(*obj.IdcsTargetAttributeName)
	}

	if obj.IdcsTargetAttributeNameToMigrateFrom != nil {
		result["idcs_target_attribute_name_to_migrate_from"] = string(*obj.IdcsTargetAttributeNameToMigrateFrom)
	}

	if obj.IdcsTargetNormAttributeName != nil {
		result["idcs_target_norm_attribute_name"] = string(*obj.IdcsTargetNormAttributeName)
	}

	if obj.IdcsTargetUniqueConstraintName != nil {
		result["idcs_target_unique_constraint_name"] = string(*obj.IdcsTargetUniqueConstraintName)
	}

	if obj.IdcsToTargetMapper != nil {
		result["idcs_to_target_mapper"] = string(*obj.IdcsToTargetMapper)
	}

	if obj.IdcsTrimStringValue != nil {
		result["idcs_trim_string_value"] = bool(*obj.IdcsTrimStringValue)
	}

	if obj.IdcsValidateReference != nil {
		result["idcs_validate_reference"] = bool(*obj.IdcsValidateReference)
	}

	if obj.IdcsValuePersisted != nil {
		result["idcs_value_persisted"] = bool(*obj.IdcsValuePersisted)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.MultiValued != nil {
		result["multi_valued"] = bool(*obj.MultiValued)
	}

	result["mutability"] = string(obj.Mutability)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["reference_types"] = obj.ReferenceTypes

	if obj.Required != nil {
		result["required"] = bool(*obj.Required)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["returned"] = string(obj.Returned)

	result["schemas"] = obj.Schemas

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	result["type"] = string(obj.Type)

	result["uniqueness"] = string(obj.Uniqueness)

	return result
}
