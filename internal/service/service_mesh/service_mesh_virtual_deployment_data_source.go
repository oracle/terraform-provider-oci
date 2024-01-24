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

func ServiceMeshVirtualDeploymentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["virtual_deployment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ServiceMeshVirtualDeploymentResource(), fieldMap, readSingularServiceMeshVirtualDeployment)
}

func readSingularServiceMeshVirtualDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualDeploymentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

type ServiceMeshVirtualDeploymentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_mesh.ServiceMeshClient
	Res    *oci_service_mesh.GetVirtualDeploymentResponse
}

func (s *ServiceMeshVirtualDeploymentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceMeshVirtualDeploymentDataSourceCrud) Get() error {
	request := oci_service_mesh.GetVirtualDeploymentRequest{}

	if virtualDeploymentId, ok := s.D.GetOkExists("virtual_deployment_id"); ok {
		tmp := virtualDeploymentId.(string)
		request.VirtualDeploymentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_mesh")

	response, err := s.Client.GetVirtualDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ServiceMeshVirtualDeploymentDataSourceCrud) SetData() error {
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	listeners := []interface{}{}
	for _, item := range s.Res.Listeners {
		listeners = append(listeners, VirtualDeploymentListenerToMap(item))
	}
	s.D.Set("listeners", listeners)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ServiceDiscovery != nil {
		serviceDiscoveryArray := []interface{}{}
		if serviceDiscoveryMap := ServiceDiscoveryConfigurationToMap(&s.Res.ServiceDiscovery); serviceDiscoveryMap != nil {
			serviceDiscoveryArray = append(serviceDiscoveryArray, serviceDiscoveryMap)
		}
		s.D.Set("service_discovery", serviceDiscoveryArray)
	} else {
		s.D.Set("service_discovery", nil)
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

	if s.Res.VirtualServiceId != nil {
		s.D.Set("virtual_service_id", *s.Res.VirtualServiceId)
	}

	return nil
}
