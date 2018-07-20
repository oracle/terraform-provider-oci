// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ServiceGatewaysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readServiceGateways,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
				Elem:     crud.GetDataSourceItemSchema(ServiceGatewayResource()),
			},
		},
	}
}

func readServiceGateways(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceGatewaysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type ServiceGatewaysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListServiceGatewaysResponse
}

func (s *ServiceGatewaysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceGatewaysDataSourceCrud) Get() error {
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

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

func (s *ServiceGatewaysDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		serviceGateway := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.BlockTraffic != nil {
			serviceGateway["block_traffic"] = *r.BlockTraffic
		}

		if r.DefinedTags != nil {
			serviceGateway["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			serviceGateway["display_name"] = *r.DisplayName
		}

		serviceGateway["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			serviceGateway["id"] = *r.Id
		}

		services := []interface{}{}
		for _, item := range r.Services {
			services = append(services, ServiceIdResponseDetailsToMap(item))
		}
		serviceGateway["services"] = schema.NewSet(servicesHashCodeForSets, services)

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
		resources = ApplyFilters(f.(*schema.Set), resources, ServiceGatewaysDataSource().Schema["service_gateways"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("service_gateways", resources); err != nil {
		panic(err)
	}

	return
}
