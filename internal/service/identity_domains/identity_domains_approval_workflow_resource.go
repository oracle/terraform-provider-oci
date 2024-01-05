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

func IdentityDomainsApprovalWorkflowResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsApprovalWorkflow,
		Read:     readIdentityDomainsApprovalWorkflow,
		Update:   updateIdentityDomainsApprovalWorkflow,
		Delete:   deleteIdentityDomainsApprovalWorkflow,
		Schema: map[string]*schema.Schema{
			// Required
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"max_duration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"unit": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"name": {
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
			"approval_workflow_steps": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"order": {
							Type:     schema.TypeInt,
							Computed: true,
						},
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
			"description": {
				Type:     schema.TypeString,
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

func createIdentityDomainsApprovalWorkflow(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsApprovalWorkflowResourceCrud{}
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

func readIdentityDomainsApprovalWorkflow(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsApprovalWorkflowResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "approvalWorkflows")
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

func updateIdentityDomainsApprovalWorkflow(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsApprovalWorkflowResourceCrud{}
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

func deleteIdentityDomainsApprovalWorkflow(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsApprovalWorkflowResourceCrud{}
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

type IdentityDomainsApprovalWorkflowResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.ApprovalWorkflow
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsApprovalWorkflowResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsApprovalWorkflowResourceCrud) Create() error {
	request := oci_identity_domains.CreateApprovalWorkflowRequest{}

	if approvalWorkflowSteps, ok := s.D.GetOkExists("approval_workflow_steps"); ok {
		interfaces := approvalWorkflowSteps.([]interface{})
		tmp := make([]oci_identity_domains.ApprovalWorkflowApprovalWorkflowSteps, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "approval_workflow_steps", stateDataIndex)
			converted, err := s.mapToApprovalWorkflowApprovalWorkflowSteps(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("approval_workflow_steps") {
			request.ApprovalWorkflowSteps = tmp
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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if maxDuration, ok := s.D.GetOkExists("max_duration"); ok {
		if tmpList := maxDuration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "max_duration", 0)
			tmp, err := s.mapToApprovalWorkflowMaxDuration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaxDuration = &tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
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

	response, err := s.Client.CreateApprovalWorkflow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApprovalWorkflow
	return nil
}

func (s *IdentityDomainsApprovalWorkflowResourceCrud) Get() error {
	request := oci_identity_domains.GetApprovalWorkflowRequest{}

	tmp := s.D.Id()
	request.ApprovalWorkflowId = &tmp

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

	approvalWorkflowId, err := parseApprovalWorkflowCompositeId(s.D.Id())
	if err == nil {
		request.ApprovalWorkflowId = &approvalWorkflowId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetApprovalWorkflow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApprovalWorkflow
	return nil
}

func (s *IdentityDomainsApprovalWorkflowResourceCrud) Update() error {
	request := oci_identity_domains.PutApprovalWorkflowRequest{}

	tmp := s.D.Id()
	request.ApprovalWorkflowId = &tmp

	if approvalWorkflowSteps, ok := s.D.GetOkExists("approval_workflow_steps"); ok {
		interfaces := approvalWorkflowSteps.([]interface{})
		tmp := make([]oci_identity_domains.ApprovalWorkflowApprovalWorkflowSteps, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "approval_workflow_steps", stateDataIndex)
			converted, err := s.mapToApprovalWorkflowApprovalWorkflowSteps(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("approval_workflow_steps") {
			request.ApprovalWorkflowSteps = tmp
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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp = s.D.Id()
	request.Id = &tmp

	if maxDuration, ok := s.D.GetOkExists("max_duration"); ok {
		if tmpList := maxDuration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "max_duration", 0)
			tmp, err := s.mapToApprovalWorkflowMaxDuration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaxDuration = &tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
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

	response, err := s.Client.PutApprovalWorkflow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApprovalWorkflow
	return nil
}

func (s *IdentityDomainsApprovalWorkflowResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteApprovalWorkflowRequest{}

	tmp := s.D.Id()
	request.ApprovalWorkflowId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteApprovalWorkflow(context.Background(), request)
	return err
}

func (s *IdentityDomainsApprovalWorkflowResourceCrud) SetData() error {

	approvalWorkflowId, err := parseApprovalWorkflowCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(approvalWorkflowId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	approvalWorkflowSteps := []interface{}{}
	for _, item := range s.Res.ApprovalWorkflowSteps {
		approvalWorkflowSteps = append(approvalWorkflowSteps, ApprovalWorkflowApprovalWorkflowStepsToMap(item))
	}
	s.D.Set("approval_workflow_steps", approvalWorkflowSteps)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
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

	if s.Res.MaxDuration != nil {
		s.D.Set("max_duration", []interface{}{ApprovalWorkflowMaxDurationToMap(s.Res.MaxDuration)})
	} else {
		s.D.Set("max_duration", nil)
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

func parseApprovalWorkflowCompositeId(compositeId string) (approvalWorkflowId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/approvalWorkflows/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	approvalWorkflowId, _ = url.PathUnescape(parts[3])

	return
}

func ApprovalWorkflowToMap(obj oci_identity_domains.ApprovalWorkflow) map[string]interface{} {
	result := map[string]interface{}{}

	approvalWorkflowSteps := []interface{}{}
	for _, item := range obj.ApprovalWorkflowSteps {
		approvalWorkflowSteps = append(approvalWorkflowSteps, ApprovalWorkflowApprovalWorkflowStepsToMap(item))
	}
	result["approval_workflow_steps"] = approvalWorkflowSteps

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

	if obj.MaxDuration != nil {
		result["max_duration"] = []interface{}{ApprovalWorkflowMaxDurationToMap(obj.MaxDuration)}
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

func (s *IdentityDomainsApprovalWorkflowResourceCrud) mapToApprovalWorkflowApprovalWorkflowSteps(fieldKeyFormat string) (oci_identity_domains.ApprovalWorkflowApprovalWorkflowSteps, error) {
	result := oci_identity_domains.ApprovalWorkflowApprovalWorkflowSteps{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if ocid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocid")); ok {
		tmp := ocid.(string)
		result.Ocid = &tmp
	}

	if order, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "order")); ok {
		tmp := order.(int)
		result.Order = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func ApprovalWorkflowApprovalWorkflowStepsToMap(obj oci_identity_domains.ApprovalWorkflowApprovalWorkflowSteps) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Order != nil {
		result["order"] = int(*obj.Order)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsApprovalWorkflowResourceCrud) mapToApprovalWorkflowMaxDuration(fieldKeyFormat string) (oci_identity_domains.ApprovalWorkflowMaxDuration, error) {
	result := oci_identity_domains.ApprovalWorkflowMaxDuration{}

	if unit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit")); ok {
		result.Unit = oci_identity_domains.ApprovalWorkflowMaxDurationUnitEnum(unit.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(int)
		result.Value = &tmp
	}

	return result, nil
}

func ApprovalWorkflowMaxDurationToMap(obj *oci_identity_domains.ApprovalWorkflowMaxDuration) map[string]interface{} {
	result := map[string]interface{}{}

	result["unit"] = string(obj.Unit)

	if obj.Value != nil {
		result["value"] = int(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsApprovalWorkflowResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
