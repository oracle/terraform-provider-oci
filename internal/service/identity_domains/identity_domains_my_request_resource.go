// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsMyRequestResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsMyRequest,
		Read:     readIdentityDomainsMyRequest,
		Delete:   deleteIdentityDomainsMyRequest,
		Schema: map[string]*schema.Schema{
			// Required
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"justification": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"requesting": {
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
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
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
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			/*
				approval_details is read-only, but not always has value in response.
				Keep it optional:true to allow it to be empty.
			*/
			"approval_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"approval_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"approver_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"approver_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"justification": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"order": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
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
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"requestor": {
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
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
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
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires": {
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
							ForceNew: true,
						},

						// Optional
						"display": {
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
						"ocid": {
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
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"ocid": {
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
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityDomainsMyRequest(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsMyRequestResourceCrud{}
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

func readIdentityDomainsMyRequest(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteIdentityDomainsMyRequest(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityDomainsMyRequestResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.MyRequest
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsMyRequestResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsMyRequestResourceCrud) Create() error {
	request := oci_identity_domains.CreateMyRequestRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		request.Action = oci_identity_domains.MyRequestActionEnum(action.(string))
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

	if justification, ok := s.D.GetOkExists("justification"); ok {
		tmp := justification.(string)
		request.Justification = &tmp
	}

	if requesting, ok := s.D.GetOkExists("requesting"); ok {
		if tmpList := requesting.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "requesting", 0)
			tmp, err := s.mapToMyRequestRequesting(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Requesting = &tmp
		}
	}

	if requestor, ok := s.D.GetOkExists("requestor"); ok {
		if tmpList := requestor.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "requestor", 0)
			tmp, err := s.mapToMyRequestRequestor(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Requestor = &tmp
		}
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

	response, err := s.Client.CreateMyRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MyRequest
	return nil
}

func (s *IdentityDomainsMyRequestResourceCrud) SetData() error {

	s.D.Set("action", s.Res.Action)

	approvalDetails := []interface{}{}
	for _, item := range s.Res.ApprovalDetails {
		approvalDetails = append(approvalDetails, MyRequestApprovalDetailsToMap(item))
	}
	s.D.Set("approval_details", approvalDetails)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.Expires != nil {
		s.D.Set("expires", *s.Res.Expires)
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

	if s.Res.Justification != nil {
		s.D.Set("justification", *s.Res.Justification)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.Requesting != nil {
		s.D.Set("requesting", []interface{}{MyRequestRequestingToMap(s.Res.Requesting)})
	} else {
		s.D.Set("requesting", nil)
	}

	if s.Res.Requestor != nil {
		s.D.Set("requestor", []interface{}{MyRequestRequestorToMap(s.Res.Requestor)})
	} else {
		s.D.Set("requestor", nil)
	}

	s.D.Set("schemas", s.Res.Schemas)

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

func MyRequestToMap(obj oci_identity_domains.MyRequest) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	approvalDetails := []interface{}{}
	for _, item := range obj.ApprovalDetails {
		approvalDetails = append(approvalDetails, MyRequestApprovalDetailsToMap(item))
	}
	result["approval_details"] = approvalDetails

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.Expires != nil {
		result["expires"] = string(*obj.Expires)
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

	if obj.Justification != nil {
		result["justification"] = string(*obj.Justification)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Requesting != nil {
		result["requesting"] = []interface{}{MyRequestRequestingToMap(obj.Requesting)}
	}

	if obj.Requestor != nil {
		result["requestor"] = []interface{}{MyRequestRequestorToMap(obj.Requestor)}
	}

	result["schemas"] = obj.Schemas

	result["status"] = string(obj.Status)

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

func (s *IdentityDomainsMyRequestResourceCrud) mapToMyRequestApprovalDetails(fieldKeyFormat string) (oci_identity_domains.MyRequestApprovalDetails, error) {
	result := oci_identity_domains.MyRequestApprovalDetails{}

	if approvalType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "approval_type")); ok {
		tmp := approvalType.(string)
		result.ApprovalType = &tmp
	}

	if approverDisplayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "approver_display_name")); ok {
		tmp := approverDisplayName.(string)
		result.ApproverDisplayName = &tmp
	}

	if approverId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "approver_id")); ok {
		tmp := approverId.(string)
		result.ApproverId = &tmp
	}

	if justification, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "justification")); ok {
		tmp := justification.(string)
		result.Justification = &tmp
	}

	if order, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "order")); ok {
		tmp := order.(int)
		result.Order = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		tmp := status.(string)
		result.Status = &tmp
	}

	if timeUpdated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_updated")); ok {
		tmp := timeUpdated.(string)
		result.TimeUpdated = &tmp
	}

	return result, nil
}

func MyRequestApprovalDetailsToMap(obj oci_identity_domains.MyRequestApprovalDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApprovalType != nil {
		result["approval_type"] = string(*obj.ApprovalType)
	}

	if obj.ApproverDisplayName != nil {
		result["approver_display_name"] = string(*obj.ApproverDisplayName)
	}

	if obj.ApproverId != nil {
		result["approver_id"] = string(*obj.ApproverId)
	}

	if obj.Justification != nil {
		result["justification"] = string(*obj.Justification)
	}

	if obj.Order != nil {
		result["order"] = int(*obj.Order)
	}

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = string(*obj.TimeUpdated)
	}

	return result
}

func (s *IdentityDomainsMyRequestResourceCrud) mapToMyRequestRequesting(fieldKeyFormat string) (oci_identity_domains.MyRequestRequesting, error) {
	result := oci_identity_domains.MyRequestRequesting{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.MyRequestRequestingTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func MyRequestRequestingToMap(obj *oci_identity_domains.MyRequestRequesting) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

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

func (s *IdentityDomainsMyRequestResourceCrud) mapToMyRequestRequestor(fieldKeyFormat string) (oci_identity_domains.MyRequestRequestor, error) {
	result := oci_identity_domains.MyRequestRequestor{}

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

func MyRequestRequestorToMap(obj *oci_identity_domains.MyRequestRequestor) map[string]interface{} {
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

func (s *IdentityDomainsMyRequestResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
