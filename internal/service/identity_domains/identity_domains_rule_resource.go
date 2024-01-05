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

func IdentityDomainsRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsRule,
		Read:     readIdentityDomainsRule,
		Update:   updateIdentityDomainsRule,
		Delete:   deleteIdentityDomainsRule,
		Schema: map[string]*schema.Schema{
			// Required
			"condition": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"policy_type": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
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
			"return": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"return_groovy": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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
			"condition_group": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"locked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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
			"rule_groovy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func createIdentityDomainsRule(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsRuleResourceCrud{}
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

func readIdentityDomainsRule(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsRuleResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "rules")
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

func updateIdentityDomainsRule(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsRuleResourceCrud{}
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

func deleteIdentityDomainsRule(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsRuleResourceCrud{}
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

type IdentityDomainsRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.Rule
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsRuleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsRuleResourceCrud) Create() error {
	request := oci_identity_domains.CreateRuleRequest{}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
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

	if condition, ok := s.D.GetOkExists("condition"); ok {
		tmp := condition.(string)
		request.Condition = &tmp
	}

	if conditionGroup, ok := s.D.GetOkExists("condition_group"); ok {
		if tmpList := conditionGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "condition_group", 0)
			tmp, err := s.mapToRuleConditionGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConditionGroup = &tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if locked, ok := s.D.GetOkExists("locked"); ok {
		tmp := locked.(bool)
		request.Locked = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if policyType, ok := s.D.GetOkExists("policy_type"); ok {
		if tmpList := policyType.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "policy_type", 0)
			tmp, err := s.mapToRulePolicyType(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PolicyType = &tmp
		}
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if return_, ok := s.D.GetOkExists("return"); ok {
		interfaces := return_.([]interface{})
		tmp := make([]oci_identity_domains.RuleReturn, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "return", stateDataIndex)
			converted, err := s.mapToRuleReturn(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("return") {
			request.Return = tmp
		}
	}

	if ruleGroovy, ok := s.D.GetOkExists("rule_groovy"); ok {
		tmp := ruleGroovy.(string)
		request.RuleGroovy = &tmp
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

	response, err := s.Client.CreateRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Rule
	return nil
}

func (s *IdentityDomainsRuleResourceCrud) Get() error {
	request := oci_identity_domains.GetRuleRequest{}

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

	tmp := s.D.Id()
	request.RuleId = &tmp

	ruleId, err := parseRuleCompositeId(s.D.Id())
	if err == nil {
		request.RuleId = &ruleId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Rule
	return nil
}

func (s *IdentityDomainsRuleResourceCrud) Update() error {
	request := oci_identity_domains.PutRuleRequest{}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
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

	if condition, ok := s.D.GetOkExists("condition"); ok {
		tmp := condition.(string)
		request.Condition = &tmp
	}

	if conditionGroup, ok := s.D.GetOkExists("condition_group"); ok {
		if tmpList := conditionGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "condition_group", 0)
			tmp, err := s.mapToRuleConditionGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConditionGroup = &tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if locked, ok := s.D.GetOkExists("locked"); ok {
		tmp := locked.(bool)
		request.Locked = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if policyType, ok := s.D.GetOkExists("policy_type"); ok {
		if tmpList := policyType.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "policy_type", 0)
			tmp, err := s.mapToRulePolicyType(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PolicyType = &tmp
		}
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if return_, ok := s.D.GetOkExists("return"); ok {
		interfaces := return_.([]interface{})
		tmp := make([]oci_identity_domains.RuleReturn, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "return", stateDataIndex)
			converted, err := s.mapToRuleReturn(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("return") {
			request.Return = tmp
		}
	}

	if ruleGroovy, ok := s.D.GetOkExists("rule_groovy"); ok {
		tmp := ruleGroovy.(string)
		request.RuleGroovy = &tmp
	}

	tmp := s.D.Id()
	request.RuleId = &tmp

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

	response, err := s.Client.PutRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Rule
	return nil
}

func (s *IdentityDomainsRuleResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteRuleRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	tmp := s.D.Id()
	request.RuleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteRule(context.Background(), request)
	return err
}

func (s *IdentityDomainsRuleResourceCrud) SetData() error {

	ruleId, err := parseRuleCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(ruleId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Active != nil {
		s.D.Set("active", *s.Res.Active)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.Condition != nil {
		s.D.Set("condition", *s.Res.Condition)
	}

	if s.Res.ConditionGroup != nil {
		s.D.Set("condition_group", []interface{}{RuleConditionGroupToMap(s.Res.ConditionGroup)})
	} else {
		s.D.Set("condition_group", nil)
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

	if s.Res.Locked != nil {
		s.D.Set("locked", *s.Res.Locked)
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

	if s.Res.PolicyType != nil {
		s.D.Set("policy_type", []interface{}{RulePolicyTypeToMap(s.Res.PolicyType)})
	} else {
		s.D.Set("policy_type", nil)
	}

	return_ := []interface{}{}
	for _, item := range s.Res.Return {
		return_ = append(return_, RuleReturnToMap(item))
	}
	s.D.Set("return", return_)

	if s.Res.RuleGroovy != nil {
		s.D.Set("rule_groovy", *s.Res.RuleGroovy)
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

func parseRuleCompositeId(compositeId string) (ruleId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/rules/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	ruleId, _ = url.PathUnescape(parts[3])

	return
}

func RuleToMap(obj oci_identity_domains.Rule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Active != nil {
		result["active"] = bool(*obj.Active)
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.Condition != nil {
		result["condition"] = string(*obj.Condition)
	}

	if obj.ConditionGroup != nil {
		result["condition_group"] = []interface{}{RuleConditionGroupToMap(obj.ConditionGroup)}
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

	if obj.Locked != nil {
		result["locked"] = bool(*obj.Locked)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.PolicyType != nil {
		result["policy_type"] = []interface{}{RulePolicyTypeToMap(obj.PolicyType)}
	}

	return_ := []interface{}{}
	for _, item := range obj.Return {
		return_ = append(return_, RuleReturnToMap(item))
	}
	result["return"] = return_

	if obj.RuleGroovy != nil {
		result["rule_groovy"] = string(*obj.RuleGroovy)
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

func (s *IdentityDomainsRuleResourceCrud) mapToRuleConditionGroup(fieldKeyFormat string) (oci_identity_domains.RuleConditionGroup, error) {
	result := oci_identity_domains.RuleConditionGroup{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.RuleConditionGroupTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func RuleConditionGroupToMap(obj *oci_identity_domains.RuleConditionGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsRuleResourceCrud) mapToRulePolicyType(fieldKeyFormat string) (oci_identity_domains.RulePolicyType, error) {
	result := oci_identity_domains.RulePolicyType{}

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

func RulePolicyTypeToMap(obj *oci_identity_domains.RulePolicyType) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsRuleResourceCrud) mapToRuleReturn(fieldKeyFormat string) (oci_identity_domains.RuleReturn, error) {
	result := oci_identity_domains.RuleReturn{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if returnGroovy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "return_groovy")); ok {
		tmp := returnGroovy.(string)
		result.ReturnGroovy = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func RuleReturnToMap(obj oci_identity_domains.RuleReturn) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ReturnGroovy != nil {
		result["return_groovy"] = string(*obj.ReturnGroovy)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsRuleResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
