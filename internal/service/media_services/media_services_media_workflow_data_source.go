// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesMediaWorkflowDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["media_workflow_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MediaServicesMediaWorkflowResource(), fieldMap, readSingularMediaServicesMediaWorkflow)
}

func readSingularMediaServicesMediaWorkflow(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesMediaWorkflowDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.GetMediaWorkflowResponse
}

func (s *MediaServicesMediaWorkflowDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesMediaWorkflowDataSourceCrud) Get() error {
	request := oci_media_services.GetMediaWorkflowRequest{}

	if mediaWorkflowId, ok := s.D.GetOkExists("media_workflow_id"); ok {
		tmp := mediaWorkflowId.(string)
		request.MediaWorkflowId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.GetMediaWorkflow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MediaServicesMediaWorkflowDataSourceCrud) SetData() error {
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

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, ResourceLockToMap(item))
	}
	s.D.Set("locks", locks)

	s.D.Set("media_workflow_configuration_ids", s.Res.MediaWorkflowConfigurationIds)

	if s.Res.Parameters != nil {
		jsonStr, err := json.Marshal(s.Res.Parameters)
		if err == nil {
			s.D.Set("parameters", string(jsonStr))
		}
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	tasks := []interface{}{}
	for _, item := range s.Res.Tasks {
		tasks = append(tasks, MediaWorkflowTaskToMap(item))
	}
	s.D.Set("tasks", tasks)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", strconv.FormatInt(*s.Res.Version, 10))
	}

	return nil
}
