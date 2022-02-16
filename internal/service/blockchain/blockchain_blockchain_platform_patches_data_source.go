// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package blockchain

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_blockchain "github.com/oracle/oci-go-sdk/v58/blockchain"
)

func BlockchainBlockchainPlatformPatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBlockchainBlockchainPlatformPatches,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"blockchain_platform_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"blockchain_platform_patch_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"patch_info_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"service_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_patch_due": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readBlockchainBlockchainPlatformPatches(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainBlockchainPlatformPatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()

	return tfresource.ReadResource(sync)
}

type BlockchainBlockchainPlatformPatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_blockchain.BlockchainPlatformClient
	Res    *oci_blockchain.ListBlockchainPlatformPatchesResponse
}

func (s *BlockchainBlockchainPlatformPatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BlockchainBlockchainPlatformPatchesDataSourceCrud) Get() error {
	request := oci_blockchain.ListBlockchainPlatformPatchesRequest{}

	if blockchainPlatformId, ok := s.D.GetOkExists("blockchain_platform_id"); ok {
		tmp := blockchainPlatformId.(string)
		request.BlockchainPlatformId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "blockchain")

	response, err := s.Client.ListBlockchainPlatformPatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBlockchainPlatformPatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BlockchainBlockchainPlatformPatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BlockchainBlockchainPlatformPatchesDataSource-", BlockchainBlockchainPlatformPatchesDataSource(), s.D))
	resources := []map[string]interface{}{}
	blockchainPlatformPatch := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BlockchainPlatformPatchSummaryToMap(item))
	}
	blockchainPlatformPatch["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, BlockchainBlockchainPlatformPatchesDataSource().Schema["blockchain_platform_patch_collection"].Elem.(*schema.Resource).Schema)
		blockchainPlatformPatch["items"] = items
	}

	resources = append(resources, blockchainPlatformPatch)
	if err := s.D.Set("blockchain_platform_patch_collection", resources); err != nil {
		return err
	}

	return nil
}

func BlockchainPlatformPatchSummaryToMap(obj oci_blockchain.BlockchainPlatformPatchSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.PatchInfoUrl != nil {
		result["patch_info_url"] = string(*obj.PatchInfoUrl)
	}

	if obj.ServiceVersion != nil {
		result["service_version"] = string(*obj.ServiceVersion)
	}

	if obj.TimePatchDue != nil {
		result["time_patch_due"] = obj.TimePatchDue.String()
	}

	return result
}
