// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoredResourceTaskDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["monitored_resource_task_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StackMonitoringMonitoredResourceTaskResource(), fieldMap, readSingularStackMonitoringMonitoredResourceTask)
}

func readSingularStackMonitoringMonitoredResourceTask(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceTaskDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMonitoredResourceTaskDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.GetMonitoredResourceTaskResponse
}

func (s *StackMonitoringMonitoredResourceTaskDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMonitoredResourceTaskDataSourceCrud) Get() error {
	request := oci_stack_monitoring.GetMonitoredResourceTaskRequest{}

	if monitoredResourceTaskId, ok := s.D.GetOkExists("monitored_resource_task_id"); ok {
		tmp := monitoredResourceTaskId.(string)
		request.MonitoredResourceTaskId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.GetMonitoredResourceTask(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StackMonitoringMonitoredResourceTaskDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TaskDetails != nil {
		taskDetailsArray := []interface{}{}
		if taskDetailsMap := MonitoredResourceTaskDetailsToMap(&s.Res.TaskDetails); taskDetailsMap != nil {
			taskDetailsArray = append(taskDetailsArray, taskDetailsMap)
		}
		s.D.Set("task_details", taskDetailsArray)
	} else {
		s.D.Set("task_details", nil)
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	s.D.Set("work_request_ids", s.Res.WorkRequestIds)

	return nil
}
