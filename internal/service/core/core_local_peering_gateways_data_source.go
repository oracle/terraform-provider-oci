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

func CoreLocalPeeringGatewaysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreLocalPeeringGateways,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_peering_gateways": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreLocalPeeringGatewayResource()),
			},
		},
	}
}

func readCoreLocalPeeringGateways(d *schema.ResourceData, m interface{}) error {
	sync := &CoreLocalPeeringGatewaysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreLocalPeeringGatewaysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListLocalPeeringGatewaysResponse
}

func (s *CoreLocalPeeringGatewaysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreLocalPeeringGatewaysDataSourceCrud) Get() error {
	request := oci_core.ListLocalPeeringGatewaysRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListLocalPeeringGateways(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLocalPeeringGateways(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreLocalPeeringGatewaysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreLocalPeeringGatewaysDataSource-", CoreLocalPeeringGatewaysDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		localPeeringGateway := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			localPeeringGateway["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			localPeeringGateway["display_name"] = *r.DisplayName
		}

		localPeeringGateway["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			localPeeringGateway["id"] = *r.Id
		}

		if r.IsCrossTenancyPeering != nil {
			localPeeringGateway["is_cross_tenancy_peering"] = *r.IsCrossTenancyPeering
		}

		if r.PeerAdvertisedCidr != nil {
			localPeeringGateway["peer_advertised_cidr"] = *r.PeerAdvertisedCidr
		}

		localPeeringGateway["peer_advertised_cidr_details"] = r.PeerAdvertisedCidrDetails

		if r.PeerId != nil {
			localPeeringGateway["peer_id"] = *r.PeerId
		}

		localPeeringGateway["peering_status"] = r.PeeringStatus

		if r.PeeringStatusDetails != nil {
			localPeeringGateway["peering_status_details"] = *r.PeeringStatusDetails
		}

		if r.RouteTableId != nil {
			localPeeringGateway["route_table_id"] = *r.RouteTableId
		}

		localPeeringGateway["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			localPeeringGateway["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			localPeeringGateway["vcn_id"] = *r.VcnId
		}

		resources = append(resources, localPeeringGateway)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreLocalPeeringGatewaysDataSource().Schema["local_peering_gateways"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("local_peering_gateways", resources); err != nil {
		return err
	}

	return nil
}
