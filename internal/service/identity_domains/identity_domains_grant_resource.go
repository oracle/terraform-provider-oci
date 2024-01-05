// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsGrantResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsGrant,
		Read:     readIdentityDomainsGrant,
		Delete:   deleteIdentityDomainsGrant,
		Schema: map[string]*schema.Schema{
			// Required
			"grant_mechanism": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"grantee": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"app_entitlement_collection": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"entitlement": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"attribute_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"attribute_value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"granted_attribute_values_json": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
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
			"composite_key": {
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
			"grantor": {
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
							ForceNew: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
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
							ForceNew: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
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
			"is_fulfilled": {
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
						"created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
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

func createIdentityDomainsGrant(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsGrantResourceCrud{}
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

func readIdentityDomainsGrant(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsGrantResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "grants")
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

func deleteIdentityDomainsGrant(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsGrantResourceCrud{}
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
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityDomainsGrantResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.Grant
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsGrantResourceCrud) ID() string {
	return *s.Res.Id
	//return GetGrantCompositeId(s.D.Get("id").(string))
}

func (s *IdentityDomainsGrantResourceCrud) Create() error {
	request := oci_identity_domains.CreateGrantRequest{}

	if app, ok := s.D.GetOkExists("app"); ok {
		if tmpList := app.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "app", 0)
			tmp, err := s.mapToGrantApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.App = &tmp
		}
	}

	if appEntitlementCollection, ok := s.D.GetOkExists("app_entitlement_collection"); ok {
		if tmpList := appEntitlementCollection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "app_entitlement_collection", 0)
			tmp, err := s.mapToGrantAppEntitlementCollection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AppEntitlementCollection = &tmp
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

	if entitlement, ok := s.D.GetOkExists("entitlement"); ok {
		if tmpList := entitlement.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "entitlement", 0)
			tmp, err := s.mapToGrantEntitlement(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Entitlement = &tmp
		}
	}

	if grantMechanism, ok := s.D.GetOkExists("grant_mechanism"); ok {
		request.GrantMechanism = oci_identity_domains.GrantGrantMechanismEnum(grantMechanism.(string))
	}

	if grantedAttributeValuesJson, ok := s.D.GetOkExists("granted_attribute_values_json"); ok {
		tmp := grantedAttributeValuesJson.(string)
		request.GrantedAttributeValuesJson = &tmp
	}

	if grantee, ok := s.D.GetOkExists("grantee"); ok {
		if tmpList := grantee.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "grantee", 0)
			tmp, err := s.mapToGrantGrantee(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Grantee = &tmp
		}
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
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
			converted, err := s.mapToTags(fieldKeyFormat)
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

	response, err := s.Client.CreateGrant(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Grant
	return nil
}

func (s *IdentityDomainsGrantResourceCrud) Get() error {
	request := oci_identity_domains.GetGrantRequest{}

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
	request.GrantId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	grantId, err := parseGrantCompositeId(s.D.Id())
	if err == nil {
		request.GrantId = &grantId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetGrant(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Grant
	return nil
}

func (s *IdentityDomainsGrantResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteGrantRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	tmp := s.D.Id()
	request.GrantId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteGrant(context.Background(), request)
	return err
}

func (s *IdentityDomainsGrantResourceCrud) SetData() error {

	grantId, err := parseGrantCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(grantId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.App != nil {
		s.D.Set("app", []interface{}{GrantAppToMap(s.Res.App)})
	} else {
		s.D.Set("app", nil)
	}

	if s.Res.AppEntitlementCollection != nil {
		s.D.Set("app_entitlement_collection", []interface{}{GrantAppEntitlementCollectionToMap(s.Res.AppEntitlementCollection)})
	} else {
		s.D.Set("app_entitlement_collection", nil)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.CompositeKey != nil {
		s.D.Set("composite_key", *s.Res.CompositeKey)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.Entitlement != nil {
		s.D.Set("entitlement", []interface{}{GrantEntitlementToMap(s.Res.Entitlement)})
	} else {
		s.D.Set("entitlement", nil)
	}

	s.D.Set("grant_mechanism", s.Res.GrantMechanism)

	if s.Res.GrantedAttributeValuesJson != nil {
		s.D.Set("granted_attribute_values_json", *s.Res.GrantedAttributeValuesJson)
	}

	if s.Res.Grantee != nil {
		s.D.Set("grantee", []interface{}{GrantGranteeToMap(s.Res.Grantee)})
	} else {
		s.D.Set("grantee", nil)
	}

	if s.Res.Grantor != nil {
		s.D.Set("grantor", []interface{}{GrantGrantorToMap(s.Res.Grantor)})
	} else {
		s.D.Set("grantor", nil)
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

	if s.Res.IsFulfilled != nil {
		s.D.Set("is_fulfilled", *s.Res.IsFulfilled)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
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

	return nil
}

//func GetGrantCompositeId(grantId string) string {
//	grantId = url.PathEscape(grantId)
//	id = url.PathEscape(id)
//	idcsEndpoint = url.PathEscape(idcsEndpoint)
//	compositeId := "idcsEndpoint/" + idcsEndpoint + "/grants/" + grantId
//	return compositeId
//}

func parseGrantCompositeId(compositeId string) (grantId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/grants/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	grantId, _ = url.PathUnescape(parts[3])

	return
}

func GrantToMap(obj oci_identity_domains.Grant) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.App != nil {
		result["app"] = []interface{}{GrantAppToMap(obj.App)}
	}

	if obj.AppEntitlementCollection != nil {
		result["app_entitlement_collection"] = []interface{}{GrantAppEntitlementCollectionToMap(obj.AppEntitlementCollection)}
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.CompositeKey != nil {
		result["composite_key"] = string(*obj.CompositeKey)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.Entitlement != nil {
		result["entitlement"] = []interface{}{GrantEntitlementToMap(obj.Entitlement)}
	}

	result["grant_mechanism"] = string(obj.GrantMechanism)

	if obj.GrantedAttributeValuesJson != nil {
		result["granted_attribute_values_json"] = string(*obj.GrantedAttributeValuesJson)
	}

	if obj.Grantee != nil {
		result["grantee"] = []interface{}{GrantGranteeToMap(obj.Grantee)}
	}

	if obj.Grantor != nil {
		result["grantor"] = []interface{}{GrantGrantorToMap(obj.Grantor)}
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

	if obj.IsFulfilled != nil {
		result["is_fulfilled"] = bool(*obj.IsFulfilled)
	}

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

func (s *IdentityDomainsGrantResourceCrud) mapToGrantApp(fieldKeyFormat string) (oci_identity_domains.GrantApp, error) {
	result := oci_identity_domains.GrantApp{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func GrantAppToMap(obj *oci_identity_domains.GrantApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsGrantResourceCrud) mapToGrantAppEntitlementCollection(fieldKeyFormat string) (oci_identity_domains.GrantAppEntitlementCollection, error) {
	result := oci_identity_domains.GrantAppEntitlementCollection{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func GrantAppEntitlementCollectionToMap(obj *oci_identity_domains.GrantAppEntitlementCollection) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsGrantResourceCrud) mapToGrantEntitlement(fieldKeyFormat string) (oci_identity_domains.GrantEntitlement, error) {
	result := oci_identity_domains.GrantEntitlement{}

	if attributeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_name")); ok {
		tmp := attributeName.(string)
		result.AttributeName = &tmp
	}

	if attributeValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_value")); ok {
		tmp := attributeValue.(string)
		result.AttributeValue = &tmp
	}

	return result, nil
}

func GrantEntitlementToMap(obj *oci_identity_domains.GrantEntitlement) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AttributeName != nil {
		result["attribute_name"] = string(*obj.AttributeName)
	}

	if obj.AttributeValue != nil {
		result["attribute_value"] = string(*obj.AttributeValue)
	}

	return result
}

func (s *IdentityDomainsGrantResourceCrud) mapToGrantGrantee(fieldKeyFormat string) (oci_identity_domains.GrantGrantee, error) {
	result := oci_identity_domains.GrantGrantee{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.GrantGranteeTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func GrantGranteeToMap(obj *oci_identity_domains.GrantGrantee) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func GrantGrantorToMap(obj *oci_identity_domains.GrantGrantor) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsGrantResourceCrud) mapToTags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
