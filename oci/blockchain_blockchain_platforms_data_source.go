// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_blockchain "github.com/oracle/oci-go-sdk/v25/blockchain"
)

func init() {
	RegisterDatasource("oci_blockchain_blockchain_platforms", BlockchainBlockchainPlatformsDataSource())
}

func BlockchainBlockchainPlatformsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBlockchainBlockchainPlatforms,
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
			"blockchain_platform_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(BlockchainBlockchainPlatformResource()),
						},
					},
				},
			},
		},
	}
}

func readBlockchainBlockchainPlatforms(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainBlockchainPlatformsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockchainPlatformClient()

	return ReadResource(sync)
}

type BlockchainBlockchainPlatformsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_blockchain.BlockchainPlatformClient
	Res    *oci_blockchain.ListBlockchainPlatformsResponse
}

func (s *BlockchainBlockchainPlatformsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BlockchainBlockchainPlatformsDataSourceCrud) Get() error {
	request := oci_blockchain.ListBlockchainPlatformsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_blockchain.BlockchainPlatformLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "blockchain")

	response, err := s.Client.ListBlockchainPlatforms(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBlockchainPlatforms(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BlockchainBlockchainPlatformsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}
	blockchainPlatform := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BlockchainPlatformSummaryToMap(item))
	}
	blockchainPlatform["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, BlockchainBlockchainPlatformsDataSource().Schema["blockchain_platform_collection"].Elem.(*schema.Resource).Schema)
		blockchainPlatform["items"] = items
	}

	resources = append(resources, blockchainPlatform)
	if err := s.D.Set("blockchain_platform_collection", resources); err != nil {
		return err
	}

	return nil
}
