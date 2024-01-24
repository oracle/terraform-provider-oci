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

func ServiceMeshMeshDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["mesh_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ServiceMeshMeshResource(), fieldMap, readSingularServiceMeshMesh)
}

func readSingularServiceMeshMesh(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshMeshDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

type ServiceMeshMeshDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_mesh.ServiceMeshClient
	Res    *oci_service_mesh.GetMeshResponse
}

func (s *ServiceMeshMeshDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceMeshMeshDataSourceCrud) Get() error {
	request := oci_service_mesh.GetMeshRequest{}

	if meshId, ok := s.D.GetOkExists("mesh_id"); ok {
		tmp := meshId.(string)
		request.MeshId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_mesh")

	response, err := s.Client.GetMesh(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ServiceMeshMeshDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	certificateAuthorities := []interface{}{}
	for _, item := range s.Res.CertificateAuthorities {
		certificateAuthorities = append(certificateAuthorities, CertificateAuthorityToMap(item))
	}
	s.D.Set("certificate_authorities", certificateAuthorities)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

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

	if s.Res.Mtls != nil {
		s.D.Set("mtls", []interface{}{MeshMutualTransportLayerSecurityToMap(s.Res.Mtls)})
	} else {
		s.D.Set("mtls", nil)
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
