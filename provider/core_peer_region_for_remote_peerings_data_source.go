// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func PeerRegionForRemotePeeringsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readPeerRegionForRemotePeerings,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"peer_region_for_remote_peerings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readPeerRegionForRemotePeerings(d *schema.ResourceData, m interface{}) error {
	sync := &PeerRegionForRemotePeeringsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type PeerRegionForRemotePeeringsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListAllowedPeerRegionsForRemotePeeringResponse
}

func (s *PeerRegionForRemotePeeringsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PeerRegionForRemotePeeringsDataSourceCrud) Get() error {
	request := oci_core.ListAllowedPeerRegionsForRemotePeeringRequest{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListAllowedPeerRegionsForRemotePeering(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PeerRegionForRemotePeeringsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		peerRegionForRemotePeering := map[string]interface{}{}

		if r.Name != nil {
			peerRegionForRemotePeering["name"] = *r.Name
		}

		resources = append(resources, peerRegionForRemotePeering)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, PeerRegionForRemotePeeringsDataSource().Schema["peer_region_for_remote_peerings"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("peer_region_for_remote_peerings", resources); err != nil {
		panic(err)
	}

	return
}
