// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstancePatchHistoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstancePatchHistories,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patch_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"patch_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"patch_histories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Computed
						"patch_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readBdsBdsInstancePatchHistories(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstancePatchHistoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstancePatchHistoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListPatchHistoriesResponse
}

func (s *BdsBdsInstancePatchHistoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstancePatchHistoriesDataSourceCrud) Get() error {
	request := oci_bds.ListPatchHistoriesRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if patchType, ok := s.D.GetOkExists("patch_type"); ok {
		request.PatchType = oci_bds.PatchHistorySummaryPatchTypeEnum(patchType.(string))
	}

	if patchVersion, ok := s.D.GetOkExists("patch_version"); ok {
		tmp := patchVersion.(string)
		request.PatchVersion = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.PatchHistorySummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListPatchHistories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPatchHistories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstancePatchHistoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstancePatchHistoriesDataSource-", BdsBdsInstancePatchHistoriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstancePatchHistory := map[string]interface{}{}

		bdsInstancePatchHistory["patch_type"] = r.PatchType

		bdsInstancePatchHistory["state"] = r.LifecycleState

		if r.TimeUpdated != nil {
			bdsInstancePatchHistory["time_updated"] = r.TimeUpdated.String()
		}

		if r.Version != nil {
			bdsInstancePatchHistory["version"] = *r.Version
		}

		resources = append(resources, bdsInstancePatchHistory)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsInstancePatchHistoriesDataSource().Schema["patch_histories"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("patch_histories", resources); err != nil {
		return err
	}

	return nil
}
