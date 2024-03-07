// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceApplicationTaskScheduleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["application_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["task_schedule_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["workspace_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataintegrationWorkspaceApplicationTaskScheduleResource(), fieldMap, readSingularDataintegrationWorkspaceApplicationTaskSchedule)
}

func readSingularDataintegrationWorkspaceApplicationTaskSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationTaskScheduleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

type DataintegrationWorkspaceApplicationTaskScheduleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataintegration.DataIntegrationClient
	Res    *oci_dataintegration.GetTaskScheduleResponse
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleDataSourceCrud) Get() error {
	request := oci_dataintegration.GetTaskScheduleRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if taskScheduleKey, ok := s.D.GetOkExists("task_schedule_key"); ok {
		tmp := taskScheduleKey.(string)
		request.TaskScheduleKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataintegration")

	response, err := s.Client.GetTaskSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataintegrationWorkspaceApplicationTaskScheduleDataSource-", DataintegrationWorkspaceApplicationTaskScheduleDataSource(), s.D))

	s.D.Set("auth_mode", s.Res.AuthMode)

	if s.Res.ConfigProviderDelegate != nil {
		tempDelegate, err := json.Marshal(*s.Res.ConfigProviderDelegate)
		if err == nil {
			jsonString := string(tempDelegate)
			s.D.Set("config_provider_delegate", jsonString)
		} else {
			log.Printf("error in parsing delegate object: %v", err)
		}
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.EndTimeMillis != nil {
		s.D.Set("end_time_millis", strconv.FormatInt(*s.Res.EndTimeMillis, 10))
	}

	if s.Res.ExpectedDuration != nil {
		s.D.Set("expected_duration", *s.Res.ExpectedDuration)
	}

	s.D.Set("expected_duration_unit", s.Res.ExpectedDurationUnit)

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.IsBackfillEnabled != nil {
		s.D.Set("is_backfill_enabled", *s.Res.IsBackfillEnabled)
	}

	if s.Res.IsConcurrentAllowed != nil {
		s.D.Set("is_concurrent_allowed", *s.Res.IsConcurrentAllowed)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LastRunDetails != nil {
		s.D.Set("last_run_details", []interface{}{LastRunDetailsToMap(s.Res.LastRunDetails)})
	} else {
		s.D.Set("last_run_details", nil)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{ObjectMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.ModelType != nil {
		s.D.Set("model_type", *s.Res.ModelType)
	}

	if s.Res.ModelVersion != nil {
		s.D.Set("model_version", *s.Res.ModelVersion)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectStatus != nil {
		s.D.Set("object_status", *s.Res.ObjectStatus)
	}

	if s.Res.ObjectVersion != nil {
		s.D.Set("object_version", *s.Res.ObjectVersion)
	}

	if s.Res.ParentRef != nil {
		s.D.Set("parent_ref", []interface{}{ParentReferenceToMap(s.Res.ParentRef)})
	} else {
		s.D.Set("parent_ref", nil)
	}

	if s.Res.RetryAttempts != nil {
		s.D.Set("retry_attempts", *s.Res.RetryAttempts)
	}

	if s.Res.RetryDelay != nil {
		s.D.Set("retry_delay", *s.Res.RetryDelay)
	}

	s.D.Set("retry_delay_unit", s.Res.RetryDelayUnit)

	if s.Res.ScheduleRef != nil {
		s.D.Set("schedule_ref", []interface{}{ScheduleToMap(s.Res.ScheduleRef)})
	} else {
		s.D.Set("schedule_ref", nil)
	}

	if s.Res.StartTimeMillis != nil {
		s.D.Set("start_time_millis", strconv.FormatInt(*s.Res.StartTimeMillis, 10))
	}

	return nil
}
