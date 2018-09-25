// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/terraform-providers/terraform-provider-oci/crud"
)

func NatGatewaysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNatGateways,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat_gateways": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     NatGatewayResource(),
			},
		},
	}
}

func readNatGateways(d *schema.ResourceData, m interface{}) error {
	sync := &NatGatewaysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type NatGatewaysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListNatGatewaysResponse
}

func (s *NatGatewaysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NatGatewaysDataSourceCrud) Get() error {
	request := oci_core.ListNatGatewaysRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.ListNatGatewaysLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListNatGateways(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNatGateways(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *NatGatewaysDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		natGateway := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.BlockTraffic != nil {
			natGateway["block_traffic"] = *r.BlockTraffic
		}

		if r.DisplayName != nil {
			natGateway["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			natGateway["id"] = *r.Id
		}

		if r.NatIp != nil {
			natGateway["nat_ip"] = *r.NatIp
		}

		natGateway["state"] = r.LifecycleState

		natGateway["time_created"] = r.TimeCreated.String()

		if r.VcnId != nil {
			natGateway["vcn_id"] = *r.VcnId
		}

		resources = append(resources, natGateway)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("nat_gateways", resources); err != nil {
		panic(err)
	}

	return
}
