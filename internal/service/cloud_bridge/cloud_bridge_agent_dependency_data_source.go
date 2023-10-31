// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeAgentDependencyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["agent_dependency_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudBridgeAgentDependencyResource(), fieldMap, readSingularCloudBridgeAgentDependency)
}

func readSingularCloudBridgeAgentDependency(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentDependencyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeAgentDependencyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.OcbAgentSvcClient
	Res    *oci_cloud_bridge.GetAgentDependencyResponse
}

func (s *CloudBridgeAgentDependencyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeAgentDependencyDataSourceCrud) Get() error {
	request := oci_cloud_bridge.GetAgentDependencyRequest{}

	if agentDependencyId, ok := s.D.GetOkExists("agent_dependency_id"); ok {
		tmp := agentDependencyId.(string)
		request.AgentDependencyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.GetAgentDependency(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudBridgeAgentDependencyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.AgentDependency.Id)

	if s.Res.Bucket != nil {
		s.D.Set("bucket", *s.Res.Bucket)
	}

	if s.Res.Checksum != nil {
		s.D.Set("checksum", *s.Res.Checksum)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DependencyName != nil {
		s.D.Set("dependency_name", *s.Res.DependencyName)
	}

	if s.Res.DependencyVersion != nil {
		s.D.Set("dependency_version", *s.Res.DependencyVersion)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ETag != nil {
		s.D.Set("e_tag", *s.Res.ETag)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
