// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v58/operatoraccesscontrol"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OperatorAccessControlOperatorControlAssignmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["operator_control_assignment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OperatorAccessControlOperatorControlAssignmentResource(), fieldMap, readSingularOperatorAccessControlOperatorControlAssignment)
}

func readSingularOperatorAccessControlOperatorControlAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlAssignmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlAssignmentClient()

	return tfresource.ReadResource(sync)
}

type OperatorAccessControlOperatorControlAssignmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_operator_access_control.OperatorControlAssignmentClient
	Res    *oci_operator_access_control.GetOperatorControlAssignmentResponse
}

func (s *OperatorAccessControlOperatorControlAssignmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OperatorAccessControlOperatorControlAssignmentDataSourceCrud) Get() error {
	request := oci_operator_access_control.GetOperatorControlAssignmentRequest{}

	if operatorControlAssignmentId, ok := s.D.GetOkExists("operator_control_assignment_id"); ok {
		tmp := operatorControlAssignmentId.(string)
		request.OperatorControlAssignmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "operator_access_control")

	response, err := s.Client.GetOperatorControlAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OperatorAccessControlOperatorControlAssignmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AssignerId != nil {
		s.D.Set("assigner_id", *s.Res.AssignerId)
	}

	if s.Res.Comment != nil {
		s.D.Set("comment", *s.Res.Comment)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DetachmentDescription != nil {
		s.D.Set("detachment_description", *s.Res.DetachmentDescription)
	}

	if s.Res.ErrorCode != nil {
		s.D.Set("error_code", *s.Res.ErrorCode)
	}

	if s.Res.ErrorMessage != nil {
		s.D.Set("error_message", *s.Res.ErrorMessage)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoApproveDuringMaintenance != nil {
		s.D.Set("is_auto_approve_during_maintenance", *s.Res.IsAutoApproveDuringMaintenance)
	}

	if s.Res.IsEnforcedAlways != nil {
		s.D.Set("is_enforced_always", *s.Res.IsEnforcedAlways)
	}

	if s.Res.IsLogForwarded != nil {
		s.D.Set("is_log_forwarded", *s.Res.IsLogForwarded)
	}

	if s.Res.OperatorControlId != nil {
		s.D.Set("operator_control_id", *s.Res.OperatorControlId)
	}

	if s.Res.RemoteSyslogServerAddress != nil {
		s.D.Set("remote_syslog_server_address", *s.Res.RemoteSyslogServerAddress)
	}

	if s.Res.RemoteSyslogServerCACert != nil {
		s.D.Set("remote_syslog_server_ca_cert", *s.Res.RemoteSyslogServerCACert)
	}

	if s.Res.RemoteSyslogServerPort != nil {
		s.D.Set("remote_syslog_server_port", *s.Res.RemoteSyslogServerPort)
	}

	if s.Res.ResourceCompartmentId != nil {
		s.D.Set("resource_compartment_id", *s.Res.ResourceCompartmentId)
	}

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	s.D.Set("resource_type", s.Res.ResourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeAssignmentFrom != nil {
		s.D.Set("time_assignment_from", s.Res.TimeAssignmentFrom.Format(time.RFC3339Nano))
	}

	if s.Res.TimeAssignmentTo != nil {
		s.D.Set("time_assignment_to", s.Res.TimeAssignmentTo.Format(time.RFC3339Nano))
	}

	if s.Res.TimeOfAssignment != nil {
		s.D.Set("time_of_assignment", s.Res.TimeOfAssignment.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.UnassignerId != nil {
		s.D.Set("unassigner_id", *s.Res.UnassignerId)
	}

	return nil
}
