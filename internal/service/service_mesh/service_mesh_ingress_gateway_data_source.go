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

func ServiceMeshIngressGatewayDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["ingress_gateway_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ServiceMeshIngressGatewayResource(), fieldMap, readSingularServiceMeshIngressGateway)
}

func readSingularServiceMeshIngressGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshIngressGatewayDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

type ServiceMeshIngressGatewayDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_mesh.ServiceMeshClient
	Res    *oci_service_mesh.GetIngressGatewayResponse
}

func (s *ServiceMeshIngressGatewayDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceMeshIngressGatewayDataSourceCrud) Get() error {
	request := oci_service_mesh.GetIngressGatewayRequest{}

	if ingressGatewayId, ok := s.D.GetOkExists("ingress_gateway_id"); ok {
		tmp := ingressGatewayId.(string)
		request.IngressGatewayId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_mesh")

	response, err := s.Client.GetIngressGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ServiceMeshIngressGatewayDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AccessLogging != nil {
		s.D.Set("access_logging", []interface{}{AccessLoggingConfigurationToMap(s.Res.AccessLogging)})
	} else {
		s.D.Set("access_logging", nil)
	}

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

	hosts := []interface{}{}
	for _, item := range s.Res.Hosts {
		hosts = append(hosts, IngressGatewayHostToMap(item))
	}
	s.D.Set("hosts", hosts)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MeshId != nil {
		s.D.Set("mesh_id", *s.Res.MeshId)
	}

	if s.Res.Mtls != nil {
		s.D.Set("mtls", []interface{}{IngressGatewayMutualTransportLayerSecurityToMap(s.Res.Mtls)})
	} else {
		s.D.Set("mtls", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
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
