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

func CoreNatGatewaysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreNatGateways,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
				Elem:     tfresource.GetDataSourceItemSchema(CoreNatGatewayResource()),
			},
		},
	}
}

func readCoreNatGateways(d *schema.ResourceData, m interface{}) error {
	sync := &CoreNatGatewaysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreNatGatewaysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListNatGatewaysResponse
}

func (s *CoreNatGatewaysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreNatGatewaysDataSourceCrud) Get() error {
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
		request.LifecycleState = oci_core.NatGatewayLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

func (s *CoreNatGatewaysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreNatGatewaysDataSource-", CoreNatGatewaysDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		natGateway := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.BlockTraffic != nil {
			natGateway["block_traffic"] = *r.BlockTraffic
		}

		if r.DefinedTags != nil {
			natGateway["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			natGateway["display_name"] = *r.DisplayName
		}

		natGateway["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			natGateway["id"] = *r.Id
		}

		if r.NatIp != nil {
			natGateway["nat_ip"] = *r.NatIp
		}

		if r.PublicIpId != nil {
			natGateway["public_ip_id"] = *r.PublicIpId
		}

		if r.RouteTableId != nil {
			natGateway["route_table_id"] = *r.RouteTableId
		}

		natGateway["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			natGateway["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			natGateway["vcn_id"] = *r.VcnId
		}

		resources = append(resources, natGateway)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreNatGatewaysDataSource().Schema["nat_gateways"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("nat_gateways", resources); err != nil {
		return err
	}

	return nil
}
