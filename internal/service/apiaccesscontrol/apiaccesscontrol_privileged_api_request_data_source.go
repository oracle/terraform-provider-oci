// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apiaccesscontrol

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apiaccesscontrol "github.com/oracle/oci-go-sdk/v65/apiaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApiaccesscontrolPrivilegedApiRequestDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["privileged_api_request_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApiaccesscontrolPrivilegedApiRequestResource(), fieldMap, readSingularApiaccesscontrolPrivilegedApiRequest)
}

func readSingularApiaccesscontrolPrivilegedApiRequest(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolPrivilegedApiRequestDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivilegedApiRequestsClient()

	return tfresource.ReadResource(sync)
}

type ApiaccesscontrolPrivilegedApiRequestDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apiaccesscontrol.PrivilegedApiRequestsClient
	Res    *oci_apiaccesscontrol.GetPrivilegedApiRequestResponse
}

func (s *ApiaccesscontrolPrivilegedApiRequestDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApiaccesscontrolPrivilegedApiRequestDataSourceCrud) Get() error {
	request := oci_apiaccesscontrol.GetPrivilegedApiRequestRequest{}

	if privilegedApiRequestId, ok := s.D.GetOkExists("privileged_api_request_id"); ok {
		tmp := privilegedApiRequestId.(string)
		request.PrivilegedApiRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apiaccesscontrol")

	response, err := s.Client.GetPrivilegedApiRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApiaccesscontrolPrivilegedApiRequestDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	approverDetails := []interface{}{}
	for _, item := range s.Res.ApproverDetails {
		approverDetails = append(approverDetails, ApproverDetailToMap(item))
	}
	s.D.Set("approver_details", approverDetails)

	if s.Res.ClosureComment != nil {
		s.D.Set("closure_comment", *s.Res.ClosureComment)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DurationInHrs != nil {
		s.D.Set("duration_in_hrs", *s.Res.DurationInHrs)
	}

	if s.Res.EntityType != nil {
		s.D.Set("entity_type", *s.Res.EntityType)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NotificationTopicId != nil {
		s.D.Set("notification_topic_id", *s.Res.NotificationTopicId)
	}

	if s.Res.NumberOfApproversRequired != nil {
		s.D.Set("number_of_approvers_required", *s.Res.NumberOfApproversRequired)
	}

	if s.Res.PrivilegedApiControlId != nil {
		s.D.Set("privileged_api_control_id", *s.Res.PrivilegedApiControlId)
	}

	if s.Res.PrivilegedApiControlName != nil {
		s.D.Set("privileged_api_control_name", *s.Res.PrivilegedApiControlName)
	}

	privilegedOperationList := []interface{}{}
	for _, item := range s.Res.PrivilegedOperationList {
		privilegedOperationList = append(privilegedOperationList, PrivilegedApiRequestOperationDetailsToMap(item))
	}
	s.D.Set("privileged_operation_list", privilegedOperationList)

	if s.Res.ReasonDetail != nil {
		s.D.Set("reason_detail", *s.Res.ReasonDetail)
	}

	if s.Res.ReasonSummary != nil {
		s.D.Set("reason_summary", *s.Res.ReasonSummary)
	}

	if s.Res.RequestId != nil {
		s.D.Set("request_id", *s.Res.RequestId)
	}

	s.D.Set("requested_by", s.Res.RequestedBy)

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", *s.Res.ResourceType)
	}

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.State)

	if s.Res.StateDetails != nil {
		s.D.Set("state_details", *s.Res.StateDetails)
	}

	s.D.Set("sub_resource_name_list", s.Res.SubResourceNameList)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("ticket_numbers", s.Res.TicketNumbers)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeRequestedForFutureAccess != nil {
		s.D.Set("time_requested_for_future_access", s.Res.TimeRequestedForFutureAccess.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
