// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v56/operatoraccesscontrol"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OperatorAccessControlAccessRequestDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOperatorAccessControlAccessRequest,
		Schema: map[string]*schema.Schema{
			"access_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"access_reason_summary": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"action_requests_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"approver_comment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"audit_type": {
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
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"extend_duration": {
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
			"opctl_additional_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"opctl_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"opctl_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operator_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"request_id": {
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
			"system_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_creation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_modification": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_user_creation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"workflow_id": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularOperatorAccessControlAccessRequest(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlAccessRequestDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AccessRequestsClient()

	return tfresource.ReadResource(sync)
}

type OperatorAccessControlAccessRequestDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_operator_access_control.AccessRequestsClient
	Res    *oci_operator_access_control.GetAccessRequestResponse
}

func (s *OperatorAccessControlAccessRequestDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OperatorAccessControlAccessRequestDataSourceCrud) Get() error {
	request := oci_operator_access_control.GetAccessRequestRequest{}

	if accessRequestId, ok := s.D.GetOkExists("access_request_id"); ok {
		tmp := accessRequestId.(string)
		request.AccessRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "operator_access_control")

	response, err := s.Client.GetAccessRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OperatorAccessControlAccessRequestDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AccessReasonSummary != nil {
		s.D.Set("access_reason_summary", *s.Res.AccessReasonSummary)
	}

	s.D.Set("action_requests_list", s.Res.ActionRequestsList)

	if s.Res.ApproverComment != nil {
		s.D.Set("approver_comment", *s.Res.ApproverComment)
	}

	s.D.Set("audit_type", s.Res.AuditType)

	if s.Res.ClosureComment != nil {
		s.D.Set("closure_comment", *s.Res.ClosureComment)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Duration != nil {
		s.D.Set("duration", *s.Res.Duration)
	}

	if s.Res.ExtendDuration != nil {
		s.D.Set("extend_duration", *s.Res.ExtendDuration)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoApproved != nil {
		s.D.Set("is_auto_approved", *s.Res.IsAutoApproved)
	}

	if s.Res.OpctlAdditionalMessage != nil {
		s.D.Set("opctl_additional_message", *s.Res.OpctlAdditionalMessage)
	}

	if s.Res.OpctlId != nil {
		s.D.Set("opctl_id", *s.Res.OpctlId)
	}

	if s.Res.OpctlName != nil {
		s.D.Set("opctl_name", *s.Res.OpctlName)
	}

	if s.Res.OperatorId != nil {
		s.D.Set("operator_id", *s.Res.OperatorId)
	}

	if s.Res.Reason != nil {
		s.D.Set("reason", *s.Res.Reason)
	}

	if s.Res.RequestId != nil {
		s.D.Set("request_id", *s.Res.RequestId)
	}

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	s.D.Set("resource_type", s.Res.ResourceType)

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemMessage != nil {
		s.D.Set("system_message", *s.Res.SystemMessage)
	}

	if s.Res.TimeOfCreation != nil {
		s.D.Set("time_of_creation", s.Res.TimeOfCreation.String())
	}

	if s.Res.TimeOfModification != nil {
		s.D.Set("time_of_modification", s.Res.TimeOfModification.String())
	}

	if s.Res.TimeOfUserCreation != nil {
		s.D.Set("time_of_user_creation", s.Res.TimeOfUserCreation.String())
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	s.D.Set("workflow_id", s.Res.WorkflowId)

	return nil
}

func AccessRequestSummaryToMap(obj oci_operator_access_control.AccessRequestSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccessReasonSummary != nil {
		result["access_reason_summary"] = string(*obj.AccessReasonSummary)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Duration != nil {
		result["duration"] = int(*obj.Duration)
	}

	if obj.ExtendDuration != nil {
		result["extend_duration"] = int(*obj.ExtendDuration)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsAutoApproved != nil {
		result["is_auto_approved"] = bool(*obj.IsAutoApproved)
	}

	if obj.RequestId != nil {
		result["request_id"] = string(*obj.RequestId)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	result["resource_type"] = string(obj.ResourceType)

	result["severity"] = string(obj.Severity)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeOfCreation != nil {
		result["time_of_creation"] = obj.TimeOfCreation.String()
	}

	if obj.TimeOfModification != nil {
		result["time_of_modification"] = obj.TimeOfModification.String()
	}

	if obj.TimeOfUserCreation != nil {
		result["time_of_user_creation"] = obj.TimeOfUserCreation.String()
	}

	return result
}
