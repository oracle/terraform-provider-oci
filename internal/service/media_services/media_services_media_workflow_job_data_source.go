// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesMediaWorkflowJobDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["media_workflow_job_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MediaServicesMediaWorkflowJobResource(), fieldMap, readSingularMediaServicesMediaWorkflowJob)
}

func readSingularMediaServicesMediaWorkflowJob(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowJobDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesMediaWorkflowJobDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.GetMediaWorkflowJobResponse
}

func (s *MediaServicesMediaWorkflowJobDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesMediaWorkflowJobDataSourceCrud) Get() error {
	request := oci_media_services.GetMediaWorkflowJobRequest{}

	if mediaWorkflowJobId, ok := s.D.GetOkExists("media_workflow_job_id"); ok {
		tmp := mediaWorkflowJobId.(string)
		request.MediaWorkflowJobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.GetMediaWorkflowJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MediaServicesMediaWorkflowJobDataSourceCrud) SetData() error {
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

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("media_workflow_configuration_ids", s.Res.MediaWorkflowConfigurationIds)

	if s.Res.MediaWorkflowId != nil {
		s.D.Set("media_workflow_id", *s.Res.MediaWorkflowId)
	}

	outputs := []interface{}{}
	for _, item := range s.Res.Outputs {
		outputs = append(outputs, JobOutputToMap(item))
	}
	s.D.Set("outputs", outputs)

	if s.Res.Parameters != nil {
		jsonStr, err := json.Marshal(s.Res.Parameters)
		if err == nil {
			s.D.Set("parameters", string(jsonStr))
		}
	}

	if s.Res.Runnable != nil {
		jsonStr, err := json.Marshal(s.Res.Runnable)
		if err == nil {
			s.D.Set("runnable", string(jsonStr))
		}
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	taskLifecycleState := []interface{}{}
	for _, item := range s.Res.TaskLifecycleState {
		taskLifecycleState = append(taskLifecycleState, MediaWorkflowTaskStateToMap(item))
	}
	s.D.Set("task_lifecycle_state", taskLifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
