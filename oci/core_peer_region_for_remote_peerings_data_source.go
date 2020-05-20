// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func init() {
	RegisterDatasource("oci_core_peer_region_for_remote_peerings", CorePeerRegionForRemotePeeringsDataSource())
}

func CorePeerRegionForRemotePeeringsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCorePeerRegionForRemotePeerings,
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

func readCorePeerRegionForRemotePeerings(d *schema.ResourceData, m interface{}) error {
	sync := &CorePeerRegionForRemotePeeringsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CorePeerRegionForRemotePeeringsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListAllowedPeerRegionsForRemotePeeringResponse
}

func (s *CorePeerRegionForRemotePeeringsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CorePeerRegionForRemotePeeringsDataSourceCrud) Get() error {
	request := oci_core.ListAllowedPeerRegionsForRemotePeeringRequest{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListAllowedPeerRegionsForRemotePeering(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CorePeerRegionForRemotePeeringsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		peerRegionForRemotePeering := map[string]interface{}{}

		if r.Name != nil {
			peerRegionForRemotePeering["name"] = *r.Name
		}

		resources = append(resources, peerRegionForRemotePeering)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CorePeerRegionForRemotePeeringsDataSource().Schema["peer_region_for_remote_peerings"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("peer_region_for_remote_peerings", resources); err != nil {
		return err
	}

	return nil
}
