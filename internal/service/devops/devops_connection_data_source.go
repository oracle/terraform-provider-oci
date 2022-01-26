// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"
)

func DevopsConnectionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["connection_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DevopsConnectionResource(), fieldMap, readSingularDevopsConnection)
}

func readSingularDevopsConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsConnectionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsConnectionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetConnectionResponse
}

func (s *DevopsConnectionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsConnectionDataSourceCrud) Get() error {
	request := oci_devops.GetConnectionRequest{}

	if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
		tmp := connectionId.(string)
		request.ConnectionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsConnectionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())

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

	switch v := (s.Res.Connection).(type) {
	case oci_devops.GithubAccessTokenConnection:
		s.D.Set("connection_type", "GITHUB_ACCESS_TOKEN")
	case oci_devops.GitlabAccessTokenConnection:
		s.D.Set("connection_type", "GITLAB_ACCESS_TOKEN")
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", v)
		return nil
	}

	return nil
}
