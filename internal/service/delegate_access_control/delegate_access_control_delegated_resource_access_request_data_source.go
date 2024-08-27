// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package delegate_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_delegate_access_control "github.com/oracle/oci-go-sdk/v65/delegateaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DelegateAccessControlDelegatedResourceAccessRequestDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDelegateAccessControlDelegatedResourceAccessRequest,
		Schema: map[string]*schema.Schema{
			"delegated_resource_access_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularDelegateAccessControlDelegatedResourceAccessRequest(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegatedResourceAccessRequestDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

type DelegateAccessControlDelegatedResourceAccessRequestDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_delegate_access_control.DelegateAccessControlClient
	Res    *oci_delegate_access_control.GetDelegatedResourceAccessRequestResponse
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestDataSourceCrud) Get() error {
	request := oci_delegate_access_control.GetDelegatedResourceAccessRequestRequest{}

	if delegatedResourceAccessRequestId, ok := s.D.GetOkExists("delegated_resource_access_request_id"); ok {
		tmp := delegatedResourceAccessRequestId.(string)
		request.DelegatedResourceAccessRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "delegate_access_control")

	response, err := s.Client.GetDelegatedResourceAccessRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	approvalInfo := []interface{}{}
	for _, item := range s.Res.ApprovalInfo {
		approvalInfo = append(approvalInfo, DelegatedResourceAccessRequestApprovalDetailsToMap(item))
	}
	s.D.Set("approval_info", approvalInfo)

	s.D.Set("audit_types", s.Res.AuditTypes)

	if s.Res.ClosureComment != nil {
		s.D.Set("closure_comment", *s.Res.ClosureComment)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_name_list", s.Res.DatabaseNameList)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DelegationControlId != nil {
		s.D.Set("delegation_control_id", *s.Res.DelegationControlId)
	}

	s.D.Set("delegation_subscription_ids", s.Res.DelegationSubscriptionIds)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DurationInHours != nil {
		s.D.Set("duration_in_hours", *s.Res.DurationInHours)
	}

	if s.Res.ExtendDurationInHours != nil {
		s.D.Set("extend_duration_in_hours", *s.Res.ExtendDurationInHours)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoApproved != nil {
		s.D.Set("is_auto_approved", *s.Res.IsAutoApproved)
	}

	if s.Res.IsPendingMoreInfo != nil {
		s.D.Set("is_pending_more_info", *s.Res.IsPendingMoreInfo)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.NumExtensionApprovals != nil {
		s.D.Set("num_extension_approvals", *s.Res.NumExtensionApprovals)
	}

	if s.Res.NumInitialApprovals != nil {
		s.D.Set("num_initial_approvals", *s.Res.NumInitialApprovals)
	}

	s.D.Set("provided_service_types", s.Res.ProvidedServiceTypes)

	if s.Res.ReasonForRequest != nil {
		s.D.Set("reason_for_request", *s.Res.ReasonForRequest)
	}

	s.D.Set("request_status", s.Res.RequestStatus)

	s.D.Set("requested_action_names", s.Res.RequestedActionNames)

	s.D.Set("requester_type", s.Res.RequesterType)

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	s.D.Set("resource_type", s.Res.ResourceType)

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("ticket_numbers", s.Res.TicketNumbers)

	if s.Res.TimeAccessRequested != nil {
		s.D.Set("time_access_requested", s.Res.TimeAccessRequested.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
