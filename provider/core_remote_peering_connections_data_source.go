// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func RemotePeeringConnectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readRemotePeeringConnections,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_peering_connections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     RemotePeeringConnectionResource(),
			},
		},
	}
}

func readRemotePeeringConnections(d *schema.ResourceData, m interface{}) error {
	sync := &RemotePeeringConnectionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type RemotePeeringConnectionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListRemotePeeringConnectionsResponse
}

func (s *RemotePeeringConnectionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RemotePeeringConnectionsDataSourceCrud) Get() error {
	request := oci_core.ListRemotePeeringConnectionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListRemotePeeringConnections(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRemotePeeringConnections(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *RemotePeeringConnectionsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		remotePeeringConnection := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DisplayName != nil {
			remotePeeringConnection["display_name"] = *r.DisplayName
		}

		if r.DrgId != nil {
			remotePeeringConnection["drg_id"] = *r.DrgId
		}

		if r.Id != nil {
			remotePeeringConnection["id"] = *r.Id
		}

		if r.IsCrossTenancyPeering != nil {
			remotePeeringConnection["is_cross_tenancy_peering"] = *r.IsCrossTenancyPeering
		}

		if r.PeerId != nil {
			remotePeeringConnection["peer_id"] = *r.PeerId
		}

		if r.PeerRegionName != nil {
			remotePeeringConnection["peer_region_name"] = *r.PeerRegionName
		}

		if r.PeerTenancyId != nil {
			remotePeeringConnection["peer_tenancy_id"] = *r.PeerTenancyId
		}

		remotePeeringConnection["peering_status"] = r.PeeringStatus

		remotePeeringConnection["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			remotePeeringConnection["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, remotePeeringConnection)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, RemotePeeringConnectionsDataSource().Schema["remote_peering_connections"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("remote_peering_connections", resources); err != nil {
		panic(err)
	}

	return
}
