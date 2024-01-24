// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v65/operatoraccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OperatorAccessControlOperatorControlDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["operator_control_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OperatorAccessControlOperatorControlResource(), fieldMap, readSingularOperatorAccessControlOperatorControl)
}

func readSingularOperatorAccessControlOperatorControl(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlClient()

	return tfresource.ReadResource(sync)
}

type OperatorAccessControlOperatorControlDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_operator_access_control.OperatorControlClient
	Res    *oci_operator_access_control.GetOperatorControlResponse
}

func (s *OperatorAccessControlOperatorControlDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OperatorAccessControlOperatorControlDataSourceCrud) Get() error {
	request := oci_operator_access_control.GetOperatorControlRequest{}

	if operatorControlId, ok := s.D.GetOkExists("operator_control_id"); ok {
		tmp := operatorControlId.(string)
		request.OperatorControlId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "operator_access_control")

	response, err := s.Client.GetOperatorControl(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OperatorAccessControlOperatorControlDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("approval_required_op_action_list", s.Res.ApprovalRequiredOpActionList)

	s.D.Set("approver_groups_list", s.Res.ApproverGroupsList)

	s.D.Set("approvers_list", s.Res.ApproversList)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("email_id_list", s.Res.EmailIdList)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsDefaultOperatorControl != nil {
		s.D.Set("is_default_operator_control", *s.Res.IsDefaultOperatorControl)
	}

	if s.Res.IsFullyPreApproved != nil {
		s.D.Set("is_fully_pre_approved", *s.Res.IsFullyPreApproved)
	}

	if s.Res.LastModifiedInfo != nil {
		s.D.Set("last_modified_info", *s.Res.LastModifiedInfo)
	}

	if s.Res.OperatorControlName != nil {
		s.D.Set("operator_control_name", *s.Res.OperatorControlName)
	}

	s.D.Set("pre_approved_op_action_list", s.Res.PreApprovedOpActionList)

	s.D.Set("resource_type", s.Res.ResourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemMessage != nil {
		s.D.Set("system_message", *s.Res.SystemMessage)
	}

	if s.Res.TimeOfCreation != nil {
		s.D.Set("time_of_creation", s.Res.TimeOfCreation.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.TimeOfModification != nil {
		s.D.Set("time_of_modification", s.Res.TimeOfModification.String())
	}

	return nil
}
