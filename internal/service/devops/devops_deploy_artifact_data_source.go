// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsDeployArtifactDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["deploy_artifact_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DevopsDeployArtifactResource(), fieldMap, readSingularDevopsDeployArtifact)
}

func readSingularDevopsDeployArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployArtifactDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsDeployArtifactDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetDeployArtifactResponse
}

func (s *DevopsDeployArtifactDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsDeployArtifactDataSourceCrud) Get() error {
	request := oci_devops.GetDeployArtifactRequest{}

	if deployArtifactId, ok := s.D.GetOkExists("deploy_artifact_id"); ok {
		tmp := deployArtifactId.(string)
		request.DeployArtifactId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetDeployArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsDeployArtifactDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("argument_substitution_mode", s.Res.ArgumentSubstitutionMode)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeployArtifactSource != nil {
		deployArtifactSourceArray := []interface{}{}
		if deployArtifactSourceMap := DeployArtifactSourceToMap(&s.Res.DeployArtifactSource); deployArtifactSourceMap != nil {
			deployArtifactSourceArray = append(deployArtifactSourceArray, deployArtifactSourceMap)
		}
		s.D.Set("deploy_artifact_source", deployArtifactSourceArray)
	} else {
		s.D.Set("deploy_artifact_source", nil)
	}

	s.D.Set("deploy_artifact_type", s.Res.DeployArtifactType)

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

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
