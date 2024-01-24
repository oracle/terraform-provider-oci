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

func BdsBdsInstancePatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstancePatches,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Computed
						"time_released": {
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

func readBdsBdsInstancePatches(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstancePatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstancePatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListPatchesResponse
}

func (s *BdsBdsInstancePatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstancePatchesDataSourceCrud) Get() error {
	request := oci_bds.ListPatchesRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListPatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstancePatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstancePatchesDataSource-", BdsBdsInstancePatchesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstancePatch := map[string]interface{}{}

		if r.TimeReleased != nil {
			bdsInstancePatch["time_released"] = r.TimeReleased.String()
		}

		if r.Version != nil {
			bdsInstancePatch["version"] = *r.Version
		}

		resources = append(resources, bdsInstancePatch)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsInstancePatchesDataSource().Schema["patches"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("patches", resources); err != nil {
		return err
	}

	return nil
}
