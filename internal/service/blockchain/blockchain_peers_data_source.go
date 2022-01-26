// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package blockchain

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_blockchain "github.com/oracle/oci-go-sdk/v56/blockchain"
)

func BlockchainPeersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBlockchainPeers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"blockchain_platform_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"peer_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     BlockchainPeerResource(),
						},
					},
				},
			},
		},
	}
}

func readBlockchainPeers(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainPeersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()

	return tfresource.ReadResource(sync)
}

type BlockchainPeersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_blockchain.BlockchainPlatformClient
	Res    *oci_blockchain.ListPeersResponse
}

func (s *BlockchainPeersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BlockchainPeersDataSourceCrud) Get() error {
	request := oci_blockchain.ListPeersRequest{}

	if blockchainPlatformId, ok := s.D.GetOkExists("blockchain_platform_id"); ok {
		tmp := blockchainPlatformId.(string)
		request.BlockchainPlatformId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "blockchain")

	response, err := s.Client.ListPeers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPeers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BlockchainPeersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BlockchainPeersDataSource-", BlockchainPeersDataSource(), s.D))
	resources := []map[string]interface{}{}
	peer := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PeerSummaryToMap(item))
	}
	peer["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, BlockchainPeersDataSource().Schema["peer_collection"].Elem.(*schema.Resource).Schema)
		peer["items"] = items
	}

	resources = append(resources, peer)
	if err := s.D.Set("peer_collection", resources); err != nil {
		return err
	}

	return nil
}
