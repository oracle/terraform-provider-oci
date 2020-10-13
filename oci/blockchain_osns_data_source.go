// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_blockchain "github.com/oracle/oci-go-sdk/v27/blockchain"
)

func init() {
	RegisterDatasource("oci_blockchain_osns", BlockchainOsnsDataSource())
}

func BlockchainOsnsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBlockchainOsns,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"blockchain_platform_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"osn_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     BlockchainOsnResource(),
						},
					},
				},
			},
		},
	}
}

func readBlockchainOsns(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainOsnsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockchainPlatformClient()

	return ReadResource(sync)
}

type BlockchainOsnsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_blockchain.BlockchainPlatformClient
	Res    *oci_blockchain.ListOsnsResponse
}

func (s *BlockchainOsnsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BlockchainOsnsDataSourceCrud) Get() error {
	request := oci_blockchain.ListOsnsRequest{}

	if blockchainPlatformId, ok := s.D.GetOkExists("blockchain_platform_id"); ok {
		tmp := blockchainPlatformId.(string)
		request.BlockchainPlatformId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "blockchain")

	response, err := s.Client.ListOsns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOsns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BlockchainOsnsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}
	osn := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OsnSummaryToMap(item))
	}
	osn["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, BlockchainOsnsDataSource().Schema["osn_collection"].Elem.(*schema.Resource).Schema)
		osn["items"] = items
	}

	resources = append(resources, osn)
	if err := s.D.Set("osn_collection", resources); err != nil {
		return err
	}

	return nil
}
