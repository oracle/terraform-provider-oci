// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreInstanceMaintenanceEventsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreInstanceMaintenanceEvents,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"correlation_token": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_window_start_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_window_start_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_maintenance_events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreInstanceMaintenanceEventResource()),
			},
		},
	}
}

func readCoreInstanceMaintenanceEvents(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceMaintenanceEventsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreInstanceMaintenanceEventsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListInstanceMaintenanceEventsResponse
}

func (s *CoreInstanceMaintenanceEventsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstanceMaintenanceEventsDataSourceCrud) Get() error {
	request := oci_core.ListInstanceMaintenanceEventsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if correlationToken, ok := s.D.GetOkExists("correlation_token"); ok {
		tmp := correlationToken.(string)
		request.CorrelationToken = &tmp
	}

	if instanceAction, ok := s.D.GetOkExists("instance_action"); ok {
		tmp := instanceAction.(string)
		request.InstanceAction = &tmp
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.InstanceMaintenanceEventLifecycleStateEnum(state.(string))
	}

	if timeWindowStartGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_window_start_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeWindowStartGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeWindowStartGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeWindowStartLessThanOrEqualTo, ok := s.D.GetOkExists("time_window_start_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeWindowStartLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeWindowStartLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListInstanceMaintenanceEvents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstanceMaintenanceEvents(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreInstanceMaintenanceEventsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreInstanceMaintenanceEventsDataSource-", CoreInstanceMaintenanceEventsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instanceMaintenanceEvent := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		instanceMaintenanceEvent["alternative_resolution_actions"] = r.AlternativeResolutionActions
		instanceMaintenanceEvent["alternative_resolution_actions"] = r.AlternativeResolutionActions

		if r.CanReschedule != nil {
			instanceMaintenanceEvent["can_reschedule"] = *r.CanReschedule
		}

		if r.CorrelationToken != nil {
			instanceMaintenanceEvent["correlation_token"] = *r.CorrelationToken
		}

		instanceMaintenanceEvent["created_by"] = r.CreatedBy

		if r.DefinedTags != nil {
			instanceMaintenanceEvent["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			instanceMaintenanceEvent["description"] = *r.Description
		}

		if r.DisplayName != nil {
			instanceMaintenanceEvent["display_name"] = *r.DisplayName
		}

		if r.EstimatedDuration != nil {
			instanceMaintenanceEvent["estimated_duration"] = *r.EstimatedDuration
		}

		instanceMaintenanceEvent["freeform_tags"] = r.FreeformTags
		instanceMaintenanceEvent["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			instanceMaintenanceEvent["id"] = *r.Id
		}

		instanceMaintenanceEvent["instance_action"] = r.InstanceAction

		if r.InstanceId != nil {
			instanceMaintenanceEvent["instance_id"] = *r.InstanceId
		}

		instanceMaintenanceEvent["maintenance_category"] = r.MaintenanceCategory

		instanceMaintenanceEvent["maintenance_reason"] = r.MaintenanceReason

		if r.StartWindowDuration != nil {
			instanceMaintenanceEvent["start_window_duration"] = *r.StartWindowDuration
		}

		instanceMaintenanceEvent["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			instanceMaintenanceEvent["time_created"] = r.TimeCreated.String()
		}

		if r.TimeFinished != nil {
			instanceMaintenanceEvent["time_finished"] = r.TimeFinished.String()
		}

		if r.TimeHardDueDate != nil {
			instanceMaintenanceEvent["time_hard_due_date"] = r.TimeHardDueDate.String()
		}

		if r.TimeStarted != nil {
			instanceMaintenanceEvent["time_started"] = r.TimeStarted.String()
		}

		if r.TimeWindowStart != nil {
			instanceMaintenanceEvent["time_window_start"] = r.TimeWindowStart.Format(time.RFC3339Nano)
		}

		resources = append(resources, instanceMaintenanceEvent)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreInstanceMaintenanceEventsDataSource().Schema["instance_maintenance_events"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instance_maintenance_events", resources); err != nil {
		return err
	}

	return nil
}
