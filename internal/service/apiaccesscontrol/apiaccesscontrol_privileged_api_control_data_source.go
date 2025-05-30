// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apiaccesscontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apiaccesscontrol "github.com/oracle/oci-go-sdk/v65/apiaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApiaccesscontrolPrivilegedApiControlDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["privileged_api_control_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApiaccesscontrolPrivilegedApiControlResource(), fieldMap, readSingularApiaccesscontrolPrivilegedApiControl)
}

func readSingularApiaccesscontrolPrivilegedApiControl(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolPrivilegedApiControlDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivilegedApiControlClient()

	return tfresource.ReadResource(sync)
}

type ApiaccesscontrolPrivilegedApiControlDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apiaccesscontrol.PrivilegedApiControlClient
	Res    *oci_apiaccesscontrol.GetPrivilegedApiControlResponse
}

func (s *ApiaccesscontrolPrivilegedApiControlDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApiaccesscontrolPrivilegedApiControlDataSourceCrud) Get() error {
	request := oci_apiaccesscontrol.GetPrivilegedApiControlRequest{}

	if privilegedApiControlId, ok := s.D.GetOkExists("privileged_api_control_id"); ok {
		tmp := privilegedApiControlId.(string)
		request.PrivilegedApiControlId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apiaccesscontrol")

	response, err := s.Client.GetPrivilegedApiControl(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApiaccesscontrolPrivilegedApiControlDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("approver_group_id_list", s.Res.ApproverGroupIdList)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NotificationTopicId != nil {
		s.D.Set("notification_topic_id", *s.Res.NotificationTopicId)
	}

	if s.Res.NumberOfApprovers != nil {
		s.D.Set("number_of_approvers", *s.Res.NumberOfApprovers)
	}

	privilegedOperationList := []interface{}{}
	for _, item := range s.Res.PrivilegedOperationList {
		privilegedOperationList = append(privilegedOperationList, PrivilegedApiDetailsToMap(item))
	}
	s.D.Set("privileged_operation_list", privilegedOperationList)

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", *s.Res.ResourceType)
	}

	s.D.Set("resources", s.Res.Resources)

	if s.Res.State != nil {
		s.D.Set("state", *s.Res.State)
	}

	if s.Res.StateDetails != nil {
		s.D.Set("state_details", *s.Res.StateDetails)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDeleted != nil {
		s.D.Set("time_deleted", s.Res.TimeDeleted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
