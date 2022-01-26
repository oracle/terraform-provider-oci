// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"
	"fmt"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v56/loganalytics"
)

func LogAnalyticsNamespaceScheduledTaskDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["namespace"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["scheduled_task_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LogAnalyticsNamespaceScheduledTaskResource(), fieldMap, readSingularLogAnalyticsNamespaceScheduledTask)
}

func readSingularLogAnalyticsNamespaceScheduledTask(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceScheduledTaskDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceScheduledTaskDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetScheduledTaskResponse
}

func (s *LogAnalyticsNamespaceScheduledTaskDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceScheduledTaskDataSourceCrud) Get() error {
	request := oci_log_analytics.GetScheduledTaskRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	} else {
		return fmt.Errorf("namespace is required to get a scheduled task")
	}

	if scheduledTaskId, ok := s.D.GetOkExists("scheduled_task_id"); ok {
		tmp := scheduledTaskId.(string)
		request.ScheduledTaskId = &tmp
	} else {
		return fmt.Errorf("scheduled_task_id is required to get a scheduled task")
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetScheduledTask(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespaceScheduledTaskDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	var namespace, scheduledTaskId string
	if tmp, ok := s.D.GetOkExists("namespace"); ok {
		namespace = tmp.(string)
	} else {
		return fmt.Errorf("namespace is required to set data for singular data source of scheduled task")
	}

	if tmp, ok := s.D.GetOkExists("scheduled_task_id"); ok {
		scheduledTaskId = tmp.(string)
	} else {
		scheduledTaskId = *((*(s.Res)).GetId())
	}
	s.D.Set("scheduled_task_id", scheduledTaskId)

	s.D.SetId(GetNamespaceScheduledTaskCompositeId(namespace, scheduledTaskId))

	if s.Res.GetAction() != nil {
		actionArray := []interface{}{}
		if actionMap := LAActionToMap(s.Res.GetAction()); actionMap != nil {
			actionArray = append(actionArray, actionMap)
		}
		s.D.Set("action", actionArray)
	} else {
		s.D.Set("action", nil)
	}

	if s.Res.GetCompartmentId() != nil {
		s.D.Set("compartment_id", *s.Res.GetCompartmentId())
	}

	if s.Res.GetDefinedTags() != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.GetDefinedTags()))
	}

	if s.Res.GetDisplayName() != nil {
		s.D.Set("display_name", *s.Res.GetDisplayName())
	}

	s.D.Set("freeform_tags", s.Res.GetFreeformTags())

	if s.Res.GetNumOccurrences() != nil {
		s.D.Set("num_occurrences", strconv.FormatInt(*s.Res.GetNumOccurrences(), 10))
	}

	if s.Res.GetSchedules() != nil {
		s.D.Set("schedules", []interface{}{ScheduleListToMap(s.Res.GetSchedules(), true)})
	} else {
		s.D.Set("schedules", nil)
	}

	s.D.Set("state", s.Res.GetLifecycleState())

	s.D.Set("task_status", s.Res.GetTaskStatus())

	s.D.Set("task_type", s.Res.GetTaskType())

	if s.Res.GetTimeCreated() != nil {
		s.D.Set("time_created", s.Res.GetTimeCreated().String())
	}

	if s.Res.GetTimeUpdated() != nil {
		s.D.Set("time_updated", s.Res.GetTimeUpdated().String())
	}

	if s.Res.GetWorkRequestId() != nil {
		s.D.Set("work_request_id", *s.Res.GetWorkRequestId())
	}

	return nil
}
