// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v58/devops"
)

func DevopsDeployEnvironmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["deploy_environment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DevopsDeployEnvironmentResource(), fieldMap, readSingularDevopsDeployEnvironment)
}

func readSingularDevopsDeployEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployEnvironmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsDeployEnvironmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetDeployEnvironmentResponse
}

func (s *DevopsDeployEnvironmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsDeployEnvironmentDataSourceCrud) Get() error {
	request := oci_devops.GetDeployEnvironmentRequest{}

	if deployEnvironmentId, ok := s.D.GetOkExists("deploy_environment_id"); ok {
		tmp := deployEnvironmentId.(string)
		request.DeployEnvironmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetDeployEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsDeployEnvironmentDataSourceCrud) SetData() error {
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

	switch (s.Res.DeployEnvironment).(type) {
	case oci_devops.ComputeInstanceGroupDeployEnvironment:
		s.D.Set("deploy_environment_type", "COMPUTE_INSTANCE_GROUP")
	case oci_devops.FunctionDeployEnvironment:
		s.D.Set("deploy_environment_type", "FUNCTION")
	case oci_devops.OkeClusterDeployEnvironment:
		s.D.Set("deploy_environment_type", "OKE_CLUSTER")
	default:
		log.Printf("[WARN] Received 'deploy_environment_type' of unknown type %v", s.Res.DeployEnvironment)
	}

	return nil
}
