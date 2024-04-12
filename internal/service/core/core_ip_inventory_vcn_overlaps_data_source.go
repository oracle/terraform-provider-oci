// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreIpInventoryVcnOverlapsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreIpInventoryVcnOverlaps,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_list": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"region_list": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"ip_inventory_vcn_overlap_summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cidr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"overlapping_cidr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"overlapping_vcn_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"overlapping_vcn_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"last_updated_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"overlap_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readCoreIpInventoryVcnOverlaps(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpInventoryVcnOverlapsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreIpInventoryVcnOverlapsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetVcnOverlapResponse
}

func (s *CoreIpInventoryVcnOverlapsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreIpInventoryVcnOverlapsDataSourceCrud) Get() error {
	request := oci_core.GetVcnOverlapRequest{}

	if compartmentList, ok := s.D.GetOkExists("compartment_list"); ok {
		interfaces := compartmentList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("compartment_list") {
			request.GetVcnOverlapDetails.CompartmentList = tmp
		}
	}

	if regionList, ok := s.D.GetOkExists("region_list"); ok {
		interfaces := regionList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("region_list") {
			request.GetVcnOverlapDetails.RegionList = tmp
		}
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetVcnOverlap(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreIpInventoryVcnOverlapsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreIpInventoryVcnOverlapsDataSource-", CoreIpInventoryVcnOverlapsDataSource(), s.D))

	ipInventoryVcnOverlapSummary := []interface{}{}
	for _, item := range s.Res.IpInventoryVcnOverlapSummary {
		ipInventoryVcnOverlapSummary = append(ipInventoryVcnOverlapSummary, IpInventoryVcnOverlapSummaryToMap(item))
	}
	s.D.Set("ip_inventory_vcn_overlap_summary", ipInventoryVcnOverlapSummary)

	if s.Res.LastUpdatedTimestamp != nil {
		s.D.Set("last_updated_timestamp", s.Res.LastUpdatedTimestamp.String())
	}

	if s.Res.Message != nil {
		s.D.Set("message", *s.Res.Message)
	}

	if s.Res.OverlapCount != nil {
		s.D.Set("overlap_count", *s.Res.OverlapCount)
	}

	return nil
}

func IpInventoryVcnOverlapSummaryToMap(obj oci_core.IpInventoryVcnOverlapSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Cidr != nil {
		result["cidr"] = string(*obj.Cidr)
	}

	if obj.OverlappingCidr != nil {
		result["overlapping_cidr"] = string(*obj.OverlappingCidr)
	}

	if obj.OverlappingVcnId != nil {
		result["overlapping_vcn_id"] = string(*obj.OverlappingVcnId)
	}

	if obj.OverlappingVcnName != nil {
		result["overlapping_vcn_name"] = string(*obj.OverlappingVcnName)
	}

	return result
}
