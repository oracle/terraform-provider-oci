// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func DevopsTriggerDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["trigger_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DevopsTriggerResource(), fieldMap, readSingularDevopsTrigger)
}

func readSingularDevopsTrigger(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsTriggerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsTriggerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetTriggerResponse
}

func (s *DevopsTriggerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsTriggerDataSourceCrud) Get() error {
	request := oci_devops.GetTriggerRequest{}

	if triggerId, ok := s.D.GetOkExists("trigger_id"); ok {
		tmp := triggerId.(string)
		request.TriggerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetTrigger(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsTriggerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())

	actions := []interface{}{}
	for _, item := range s.Res.GetActions() {
		actions = append(actions, TriggerActionToMap(item))
	}
	s.D.Set("actions", actions)

	if s.Res.GetCompartmentId() != nil {
		s.D.Set("compartment_id", *s.Res.GetCompartmentId())
	}

	if s.Res.GetDefinedTags() != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.GetDefinedTags()))
	}

	if s.Res.GetDescription() != nil {
		s.D.Set("description", *s.Res.GetDescription())
	}

	if s.Res.GetDisplayName() != nil {
		s.D.Set("display_name", *s.Res.GetDisplayName())
	}

	s.D.Set("freeform_tags", s.Res.GetFreeformTags())

	if s.Res.GetLifecycleDetails() != nil {
		s.D.Set("lifecycle_details", *s.Res.GetLifecycleDetails())
	}

	if s.Res.GetProjectId() != nil {
		s.D.Set("project_id", *s.Res.GetProjectId())
	}

	s.D.Set("state", s.Res.GetLifecycleState())

	if s.Res.GetSystemTags() != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.GetSystemTags()))
	}

	if s.Res.GetTimeCreated() != nil {
		s.D.Set("time_created", s.Res.GetTimeCreated().String())
	}

	if s.Res.GetTimeUpdated() != nil {
		s.D.Set("time_updated", s.Res.GetTimeUpdated().String())
	}

	switch v := (s.Res.Trigger).(type) {
	case oci_devops.GithubTrigger:
		s.D.Set("trigger_source", "GITHUB")

		if v.TriggerUrl != nil {
			s.D.Set("trigger_url", v.TriggerUrl)
		}
	case oci_devops.GitlabTrigger:
		s.D.Set("trigger_source", "GITLAB")

		if v.TriggerUrl != nil {
			s.D.Set("trigger_url", v.TriggerUrl)
		}
	case oci_devops.DevopsCodeRepositoryTrigger:
		s.D.Set("trigger_source", "DEVOPS_CODE_REPOSITORY")

	}

	return nil
}
