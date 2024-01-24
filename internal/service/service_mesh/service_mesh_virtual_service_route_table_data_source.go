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

func ServiceMeshVirtualServiceRouteTableDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["virtual_service_route_table_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ServiceMeshVirtualServiceRouteTableResource(), fieldMap, readSingularServiceMeshVirtualServiceRouteTable)
}

func readSingularServiceMeshVirtualServiceRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualServiceRouteTableDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

type ServiceMeshVirtualServiceRouteTableDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_mesh.ServiceMeshClient
	Res    *oci_service_mesh.GetVirtualServiceRouteTableResponse
}

func (s *ServiceMeshVirtualServiceRouteTableDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceMeshVirtualServiceRouteTableDataSourceCrud) Get() error {
	request := oci_service_mesh.GetVirtualServiceRouteTableRequest{}

	if virtualServiceRouteTableId, ok := s.D.GetOkExists("virtual_service_route_table_id"); ok {
		tmp := virtualServiceRouteTableId.(string)
		request.VirtualServiceRouteTableId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_mesh")

	response, err := s.Client.GetVirtualServiceRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ServiceMeshVirtualServiceRouteTableDataSourceCrud) SetData() error {
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

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Priority != nil {
		s.D.Set("priority", *s.Res.Priority)
	}

	routeRules := []interface{}{}
	for _, item := range s.Res.RouteRules {
		routeRules = append(routeRules, VirtualServiceTrafficRouteRuleToMap(item))
	}
	s.D.Set("route_rules", routeRules)

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

	if s.Res.VirtualServiceId != nil {
		s.D.Set("virtual_service_id", *s.Res.VirtualServiceId)
	}

	return nil
}
