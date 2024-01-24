// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreServiceGatewaysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreServiceGateways,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_gateways": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreServiceGatewayResource()),
			},
		},
	}
}

func readCoreServiceGateways(d *schema.ResourceData, m interface{}) error {
	sync := &CoreServiceGatewaysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreServiceGatewaysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListServiceGatewaysResponse
}

func (s *CoreServiceGatewaysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreServiceGatewaysDataSourceCrud) Get() error {
	request := oci_core.ListServiceGatewaysRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.ServiceGatewayLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListServiceGateways(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListServiceGateways(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreServiceGatewaysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreServiceGatewaysDataSource-", CoreServiceGatewaysDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		serviceGateway := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.BlockTraffic != nil {
			serviceGateway["block_traffic"] = *r.BlockTraffic
		}

		if r.DefinedTags != nil {
			serviceGateway["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			serviceGateway["display_name"] = *r.DisplayName
		}

		serviceGateway["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			serviceGateway["id"] = *r.Id
		}

		if r.RouteTableId != nil {
			serviceGateway["route_table_id"] = *r.RouteTableId
		}

		services := []interface{}{}
		for _, item := range r.Services {
			services = append(services, ServiceIdResponseDetailsToMap(item))
		}
		serviceGateway["services"] = services

		serviceGateway["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			serviceGateway["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			serviceGateway["vcn_id"] = *r.VcnId
		}

		resources = append(resources, serviceGateway)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreServiceGatewaysDataSource().Schema["service_gateways"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("service_gateways", resources); err != nil {
		return err
	}

	return nil
}
