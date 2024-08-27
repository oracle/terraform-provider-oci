// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package delegate_access_control

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_delegate_access_control "github.com/oracle/oci-go-sdk/v65/delegateaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DelegateAccessControlDelegatedResourceAccessRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDelegateAccessControlDelegatedResourceAccessRequests,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"delegation_control_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"delegated_resource_access_request_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"approval_info": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"approval_action": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"approval_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"approver_additional_message": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"approver_comment": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"approver_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_approved_for_access": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"audit_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"closure_comment": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_name_list": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"delegation_control_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"delegation_subscription_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"duration_in_hours": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"extend_duration_in_hours": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_auto_approved": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_pending_more_info": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"lifecycle_state_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"num_extension_approvals": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"num_initial_approvals": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"provided_service_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"reason_for_request": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"request_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"requested_action_names": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"requester_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"ticket_numbers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"time_access_requested": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
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
					},
				},
			},
		},
	}
}

func readDelegateAccessControlDelegatedResourceAccessRequests(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegatedResourceAccessRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

type DelegateAccessControlDelegatedResourceAccessRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_delegate_access_control.DelegateAccessControlClient
	Res    *oci_delegate_access_control.ListDelegatedResourceAccessRequestsResponse
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestsDataSourceCrud) Get() error {
	request := oci_delegate_access_control.ListDelegatedResourceAccessRequestsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if delegationControlId, ok := s.D.GetOkExists("delegation_control_id"); ok {
		tmp := delegationControlId.(string)
		request.DelegationControlId = &tmp
	}

	if requestStatus, ok := s.D.GetOkExists("request_status"); ok {
		request.RequestStatus = oci_delegate_access_control.ListDelegatedResourceAccessRequestsRequestStatusEnum(requestStatus.(string))
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_delegate_access_control.DelegatedResourceAccessRequestLifecycleStateEnum(state.(string))
	}

	if timeEnd, ok := s.D.GetOkExists("time_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
		if err != nil {
			return err
		}
		request.TimeEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeStart, ok := s.D.GetOkExists("time_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return err
		}
		request.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "delegate_access_control")

	response, err := s.Client.ListDelegatedResourceAccessRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDelegatedResourceAccessRequests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DelegateAccessControlDelegatedResourceAccessRequestsDataSource-", DelegateAccessControlDelegatedResourceAccessRequestsDataSource(), s.D))
	resources := []map[string]interface{}{}
	delegatedResourceAccessRequest := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DelegatedResourceAccessRequestSummaryToMap(item))
	}
	delegatedResourceAccessRequest["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DelegateAccessControlDelegatedResourceAccessRequestsDataSource().Schema["delegated_resource_access_request_summary_collection"].Elem.(*schema.Resource).Schema)
		delegatedResourceAccessRequest["items"] = items
	}

	resources = append(resources, delegatedResourceAccessRequest)
	if err := s.D.Set("delegated_resource_access_request_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func DelegatedResourceAccessRequestApprovalDetailsToMap(obj oci_delegate_access_control.DelegatedResourceAccessRequestApprovalDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["approval_action"] = string(obj.ApprovalAction)

	result["approval_type"] = string(obj.ApprovalType)

	if obj.ApproverAdditionalMessage != nil {
		result["approver_additional_message"] = string(*obj.ApproverAdditionalMessage)
	}

	if obj.ApproverComment != nil {
		result["approver_comment"] = string(*obj.ApproverComment)
	}

	if obj.ApproverId != nil {
		result["approver_id"] = string(*obj.ApproverId)
	}

	if obj.TimeApprovedForAccess != nil {
		result["time_approved_for_access"] = obj.TimeApprovedForAccess.String()
	}

	return result
}

func DelegatedResourceAccessRequestSummaryToMap(obj oci_delegate_access_control.DelegatedResourceAccessRequestSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DelegationControlId != nil {
		result["delegation_control_id"] = string(*obj.DelegationControlId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DurationInHours != nil {
		result["duration_in_hours"] = int(*obj.DurationInHours)
	}

	if obj.ExtendDurationInHours != nil {
		result["extend_duration_in_hours"] = int(*obj.ExtendDurationInHours)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsAutoApproved != nil {
		result["is_auto_approved"] = bool(*obj.IsAutoApproved)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.ReasonForRequest != nil {
		result["reason_for_request"] = string(*obj.ReasonForRequest)
	}

	result["request_status"] = string(obj.RequestStatus)

	result["requested_action_names"] = obj.RequestedActionNames

	result["requester_type"] = string(obj.RequesterType)

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	result["resource_type"] = string(obj.ResourceType)

	result["severity"] = string(obj.Severity)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	result["ticket_numbers"] = obj.TicketNumbers

	if obj.TimeAccessRequested != nil {
		result["time_access_requested"] = obj.TimeAccessRequested.String()
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
