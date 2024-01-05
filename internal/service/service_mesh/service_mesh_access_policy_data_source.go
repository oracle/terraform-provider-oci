// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_mesh

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ServiceMeshAccessPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["access_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ServiceMeshAccessPolicyResource(), fieldMap, readSingularServiceMeshAccessPolicy)
}

func readSingularServiceMeshAccessPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshAccessPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

type ServiceMeshAccessPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_mesh.ServiceMeshClient
	Res    *oci_service_mesh.GetAccessPolicyResponse
}

func (s *ServiceMeshAccessPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceMeshAccessPolicyDataSourceCrud) Get() error {
	request := oci_service_mesh.GetAccessPolicyRequest{}

	if accessPolicyId, ok := s.D.GetOkExists("access_policy_id"); ok {
		tmp := accessPolicyId.(string)
		request.AccessPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_mesh")

	response, err := s.Client.GetAccessPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ServiceMeshAccessPolicyDataSourceCrud) SetData() error {
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

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MeshId != nil {
		s.D.Set("mesh_id", *s.Res.MeshId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	rules := []interface{}{}
	for _, item := range s.Res.Rules {
		rules = append(rules, AccessPolicyRuleToMap(item))
	}
	s.D.Set("rules", rules)

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
