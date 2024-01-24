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

func BdsBdsInstanceListOsPatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstanceListOsPatches,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_patches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"bds_instance_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"os_patch_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"release_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readBdsBdsInstanceListOsPatches(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceListOsPatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceListOsPatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListOsPatchesResponse
}

func (s *BdsBdsInstanceListOsPatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceListOsPatchesDataSourceCrud) Get() error {
	request := oci_bds.ListOsPatchesRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListOsPatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOsPatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstanceListOsPatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceListOsPatchesDataSource-", BdsBdsInstanceListOsPatchesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstanceListOsPatch := map[string]interface{}{}

		if r.OsPatchVersion != nil {
			bdsInstanceListOsPatch["os_patch_version"] = *r.OsPatchVersion
		}

		if r.ReleaseDate != nil {
			bdsInstanceListOsPatch["release_date"] = r.ReleaseDate.String()
		}

		resources = append(resources, bdsInstanceListOsPatch)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsInstanceListOsPatchesDataSource().Schema["os_patches"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("os_patches", resources); err != nil {
		return err
	}

	return nil
}
