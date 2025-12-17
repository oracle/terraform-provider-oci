// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsMappedAttributeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsMappedAttribute,
		Read:     readIdentityDomainsMappedAttribute,
		Update:   updateIdentityDomainsMappedAttribute,
		Delete:   deleteIdentityDomainsMappedAttribute,
		Schema: map[string]*schema.Schema{
			// Required
			"direction": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"idcs_resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mapped_attribute_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ref_resource_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ref_resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"attribute_mappings": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      attributeMappingsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"idcs_attribute_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"managed_object_attribute_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"applies_to_actions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"saml_format": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
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
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
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
			"meta": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"tenancy_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityDomainsMappedAttribute(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsMappedAttributeResourceCrud{}
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

	return tfresource.CreateResource(d, sync)
}

func readIdentityDomainsMappedAttribute(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsMappedAttributeResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "mappedAttributes")
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

func updateIdentityDomainsMappedAttribute(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsMappedAttributeResourceCrud{}
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

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityDomainsMappedAttribute(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityDomainsMappedAttributeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.MappedAttribute
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsMappedAttributeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsMappedAttributeResourceCrud) Create() error {
	request := oci_identity_domains.PutMappedAttributeRequest{}

	if attributeMappings, ok := s.D.GetOkExists("attribute_mappings"); ok {
		set := attributeMappings.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.MappedAttributeAttributeMappings, len(interfaces))
		for i := range interfaces {
			stateDataIndex := attributeMappingsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "attribute_mappings", stateDataIndex)
			converted, err := s.mapToMappedAttributeAttributeMappings(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_mappings") {
			request.AttributeMappings = tmp
		}
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

	if direction, ok := s.D.GetOkExists("direction"); ok {
		request.Direction = oci_identity_domains.MappedAttributeDirectionEnum(direction.(string))
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if idcsResourceType, ok := s.D.GetOkExists("idcs_resource_type"); ok {
		request.IdcsResourceType = oci_identity_domains.MappedAttributeIdcsResourceTypeEnum(idcsResourceType.(string))
	}

	if mappedAttributeId, ok := s.D.GetOkExists("mapped_attribute_id"); ok {
		tmp := mappedAttributeId.(string)
		request.MappedAttributeId = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if refResourceID, ok := s.D.GetOkExists("ref_resource_id"); ok {
		tmp := refResourceID.(string)
		request.RefResourceID = &tmp
	}

	if refResourceType, ok := s.D.GetOkExists("ref_resource_type"); ok {
		request.RefResourceType = oci_identity_domains.MappedAttributeRefResourceTypeEnum(refResourceType.(string))
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutMappedAttribute(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MappedAttribute
	return nil
}

func (s *IdentityDomainsMappedAttributeResourceCrud) Get() error {
	request := oci_identity_domains.GetMappedAttributeRequest{}

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

	tmp := s.D.Id()
	request.MappedAttributeId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	mappedAttributeId, err := parseMappedAttributeCompositeId(s.D.Id())
	if err == nil {
		request.MappedAttributeId = &mappedAttributeId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetMappedAttribute(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MappedAttribute
	return nil
}

func (s *IdentityDomainsMappedAttributeResourceCrud) Update() error {
	request := oci_identity_domains.PutMappedAttributeRequest{}

	if attributeMappings, ok := s.D.GetOkExists("attribute_mappings"); ok {
		set := attributeMappings.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.MappedAttributeAttributeMappings, len(interfaces))
		for i := range interfaces {
			stateDataIndex := attributeMappingsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "attribute_mappings", stateDataIndex)
			converted, err := s.mapToMappedAttributeAttributeMappings(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_mappings") {
			request.AttributeMappings = tmp
		}
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

	if direction, ok := s.D.GetOkExists("direction"); ok {
		request.Direction = oci_identity_domains.MappedAttributeDirectionEnum(direction.(string))
	}

	tmp := s.D.Id()
	request.Id = &tmp

	if idcsResourceType, ok := s.D.GetOkExists("idcs_resource_type"); ok {
		request.IdcsResourceType = oci_identity_domains.MappedAttributeIdcsResourceTypeEnum(idcsResourceType.(string))
	}

	if mappedAttributeId, ok := s.D.GetOkExists("mapped_attribute_id"); ok {
		tmp := mappedAttributeId.(string)
		request.MappedAttributeId = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if refResourceID, ok := s.D.GetOkExists("ref_resource_id"); ok {
		tmp := refResourceID.(string)
		request.RefResourceID = &tmp
	}

	if refResourceType, ok := s.D.GetOkExists("ref_resource_type"); ok {
		request.RefResourceType = oci_identity_domains.MappedAttributeRefResourceTypeEnum(refResourceType.(string))
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutMappedAttribute(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MappedAttribute
	return nil
}

func (s *IdentityDomainsMappedAttributeResourceCrud) SetData() error {

	mappedAttributeId, err := parseMappedAttributeCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(mappedAttributeId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	attributeMappings := []interface{}{}
	for _, item := range s.Res.AttributeMappings {
		attributeMappings = append(attributeMappings, MappedAttributeAttributeMappingsToMap(item))
	}
	s.D.Set("attribute_mappings", schema.NewSet(attributeMappingsHashCodeForSets, attributeMappings))

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	s.D.Set("direction", s.Res.Direction)

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
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

	s.D.Set("idcs_resource_type", s.Res.IdcsResourceType)

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.RefResourceID != nil {
		s.D.Set("ref_resource_id", *s.Res.RefResourceID)
	}

	s.D.Set("ref_resource_type", s.Res.RefResourceType)

	s.D.Set("schemas", s.Res.Schemas)

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

func parseMappedAttributeCompositeId(compositeId string) (mappedAttributeId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/mappedAttributes/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	mappedAttributeId, _ = url.PathUnescape(parts[3])

	return
}

func MappedAttributeToMap(obj oci_identity_domains.MappedAttribute, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	attributeMappings := []interface{}{}
	for _, item := range obj.AttributeMappings {
		attributeMappings = append(attributeMappings, MappedAttributeAttributeMappingsToMap(item))
	}
	if datasource {
		result["attribute_mappings"] = attributeMappings
	} else {
		result["attribute_mappings"] = schema.NewSet(attributeMappingsHashCodeForSets, attributeMappings)
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	result["direction"] = string(obj.Direction)

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

	result["idcs_resource_type"] = string(obj.IdcsResourceType)

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.RefResourceID != nil {
		result["ref_resource_id"] = string(*obj.RefResourceID)
	}

	result["ref_resource_type"] = string(obj.RefResourceType)

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

func (s *IdentityDomainsMappedAttributeResourceCrud) mapToMappedAttributeAttributeMappings(fieldKeyFormat string) (oci_identity_domains.MappedAttributeAttributeMappings, error) {
	result := oci_identity_domains.MappedAttributeAttributeMappings{}

	if appliesToActions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "applies_to_actions")); ok {
		interfaces := appliesToActions.([]interface{})
		tmp := make([]oci_identity_domains.MappedAttributeAttributeMappingsAppliesToActionsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.MappedAttributeAttributeMappingsAppliesToActionsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "applies_to_actions")) {
			result.AppliesToActions = tmp
		}
	}

	if idcsAttributeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idcs_attribute_name")); ok {
		tmp := idcsAttributeName.(string)
		result.IdcsAttributeName = &tmp
	}

	if managedObjectAttributeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_object_attribute_name")); ok {
		tmp := managedObjectAttributeName.(string)
		result.ManagedObjectAttributeName = &tmp
	}

	if required, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "required")); ok {
		tmp := required.(bool)
		result.Required = &tmp
	}

	if samlFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "saml_format")); ok {
		tmp := samlFormat.(string)
		result.SamlFormat = &tmp
	}

	return result, nil
}

func MappedAttributeAttributeMappingsToMap(obj oci_identity_domains.MappedAttributeAttributeMappings) map[string]interface{} {
	result := map[string]interface{}{}

	result["applies_to_actions"] = obj.AppliesToActions

	if obj.IdcsAttributeName != nil {
		result["idcs_attribute_name"] = string(*obj.IdcsAttributeName)
	}

	if obj.ManagedObjectAttributeName != nil {
		result["managed_object_attribute_name"] = string(*obj.ManagedObjectAttributeName)
	}

	if obj.Required != nil {
		result["required"] = bool(*obj.Required)
	}

	if obj.SamlFormat != nil {
		result["saml_format"] = string(*obj.SamlFormat)
	}

	return result
}

func (s *IdentityDomainsMappedAttributeResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
	result := oci_identity_domains.Tags{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func attributeMappingsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if appliesToActions, ok := m["applies_to_actions"]; ok && appliesToActions != "" {
	}
	if idcsAttributeName, ok := m["idcs_attribute_name"]; ok && idcsAttributeName != "" {
		buf.WriteString(fmt.Sprintf("%v-", idcsAttributeName))
	}
	if managedObjectAttributeName, ok := m["managed_object_attribute_name"]; ok && managedObjectAttributeName != "" {
		buf.WriteString(fmt.Sprintf("%v-", managedObjectAttributeName))
	}
	if required, ok := m["required"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", required))
	}
	if samlFormat, ok := m["saml_format"]; ok && samlFormat != "" {
		buf.WriteString(fmt.Sprintf("%v-", samlFormat))
	}
	return utils.GetStringHashcode(buf.String())
}
