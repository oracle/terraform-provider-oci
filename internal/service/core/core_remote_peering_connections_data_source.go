// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreRemotePeeringConnectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreRemotePeeringConnections,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
				Elem:     tfresource.GetDataSourceItemSchema(CoreRemotePeeringConnectionResource()),
			},
		},
	}
}

func readCoreRemotePeeringConnections(d *schema.ResourceData, m interface{}) error {
	sync := &CoreRemotePeeringConnectionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreRemotePeeringConnectionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListRemotePeeringConnectionsResponse
}

func (s *CoreRemotePeeringConnectionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreRemotePeeringConnectionsDataSourceCrud) Get() error {
	request := oci_core.ListRemotePeeringConnectionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

func (s *CoreRemotePeeringConnectionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreRemotePeeringConnectionsDataSource-", CoreRemotePeeringConnectionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		remotePeeringConnection := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			remotePeeringConnection["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			remotePeeringConnection["display_name"] = *r.DisplayName
		}

		if r.DrgId != nil {
			remotePeeringConnection["drg_id"] = *r.DrgId
		}

		remotePeeringConnection["freeform_tags"] = r.FreeformTags

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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreRemotePeeringConnectionsDataSource().Schema["remote_peering_connections"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("remote_peering_connections", resources); err != nil {
		return err
	}

	return nil
}
