// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreInstanceMaintenanceEventDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["instance_maintenance_event_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreInstanceMaintenanceEventResource(), fieldMap, readSingularCoreInstanceMaintenanceEvent)
}

func readSingularCoreInstanceMaintenanceEvent(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceMaintenanceEventDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreInstanceMaintenanceEventDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetInstanceMaintenanceEventResponse
}

func (s *CoreInstanceMaintenanceEventDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstanceMaintenanceEventDataSourceCrud) Get() error {
	request := oci_core.GetInstanceMaintenanceEventRequest{}

	if instanceMaintenanceEventId, ok := s.D.GetOkExists("instance_maintenance_event_id"); ok {
		tmp := instanceMaintenanceEventId.(string)
		request.InstanceMaintenanceEventId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetInstanceMaintenanceEvent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreInstanceMaintenanceEventDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_details", s.Res.AdditionalDetails)
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	s.D.Set("alternative_resolution_actions", s.Res.AlternativeResolutionActions)
	s.D.Set("alternative_resolution_actions", s.Res.AlternativeResolutionActions)

	if s.Res.CanDeleteLocalStorage != nil {
		s.D.Set("can_delete_local_storage", *s.Res.CanDeleteLocalStorage)
	}

	if s.Res.CanReschedule != nil {
		s.D.Set("can_reschedule", *s.Res.CanReschedule)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CorrelationToken != nil {
		s.D.Set("correlation_token", *s.Res.CorrelationToken)
	}

	s.D.Set("created_by", s.Res.CreatedBy)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EstimatedDuration != nil {
		s.D.Set("estimated_duration", *s.Res.EstimatedDuration)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("instance_action", s.Res.InstanceAction)

	if s.Res.InstanceId != nil {
		s.D.Set("instance_id", *s.Res.InstanceId)
	}

	s.D.Set("maintenance_category", s.Res.MaintenanceCategory)

	s.D.Set("maintenance_reason", s.Res.MaintenanceReason)

	if s.Res.StartWindowDuration != nil {
		s.D.Set("start_window_duration", *s.Res.StartWindowDuration)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeHardDueDate != nil {
		s.D.Set("time_hard_due_date", s.Res.TimeHardDueDate.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeWindowStart != nil {
		s.D.Set("time_window_start", s.Res.TimeWindowStart.Format(time.RFC3339Nano))
		s.D.Set("time_window_start", s.Res.TimeWindowStart.String())
	}

	return nil
}
