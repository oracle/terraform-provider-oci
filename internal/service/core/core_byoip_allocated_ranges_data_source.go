// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreByoipAllocatedRangesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreByoipAllocatedRanges,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"byoip_range_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"byoip_allocated_range_collection": {
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
									"cidr_block": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"public_ip_pool_id": {
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
	}
}

func readCoreByoipAllocatedRanges(d *schema.ResourceData, m interface{}) error {
	sync := &CoreByoipAllocatedRangesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreByoipAllocatedRangesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListByoipAllocatedRangesResponse
}

func (s *CoreByoipAllocatedRangesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreByoipAllocatedRangesDataSourceCrud) Get() error {
	request := oci_core.ListByoipAllocatedRangesRequest{}

	if byoipRangeId, ok := s.D.GetOkExists("byoip_range_id"); ok {
		tmp := byoipRangeId.(string)
		request.ByoipRangeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListByoipAllocatedRanges(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListByoipAllocatedRanges(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreByoipAllocatedRangesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreByoipAllocatedRangesDataSource-", CoreByoipAllocatedRangesDataSource(), s.D))
	resources := []map[string]interface{}{}
	byoipAllocatedRange := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ByoipAllocatedRangeSummaryToMap(item))
	}
	byoipAllocatedRange["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CoreByoipAllocatedRangesDataSource().Schema["byoip_allocated_range_collection"].Elem.(*schema.Resource).Schema)
		byoipAllocatedRange["items"] = items
	}

	resources = append(resources, byoipAllocatedRange)
	if err := s.D.Set("byoip_allocated_range_collection", resources); err != nil {
		return err
	}

	return nil
}

func ByoipAllocatedRangeSummaryToMap(obj oci_core.ByoipAllocatedRangeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CidrBlock != nil {
		result["cidr_block"] = string(*obj.CidrBlock)
	}

	if obj.PublicIpPoolId != nil {
		result["public_ip_pool_id"] = string(*obj.PublicIpPoolId)
	}

	return result
}
