// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreComputeCapacityTopologyComputeNetworkBlocksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeCapacityTopologyComputeNetworkBlocks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_capacity_topology_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_hpc_island_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_network_block_collection": {
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
									"compute_capacity_topology_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compute_hpc_island_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_compute_bare_metal_host_count": {
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

func readCoreComputeCapacityTopologyComputeNetworkBlocks(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityTopologyComputeNetworkBlocksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeCapacityTopologyComputeNetworkBlocksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeCapacityTopologyComputeNetworkBlocksResponse
}

func (s *CoreComputeCapacityTopologyComputeNetworkBlocksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeCapacityTopologyComputeNetworkBlocksDataSourceCrud) Get() error {
	request := oci_core.ListComputeCapacityTopologyComputeNetworkBlocksRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeCapacityTopologyId, ok := s.D.GetOkExists("compute_capacity_topology_id"); ok {
		tmp := computeCapacityTopologyId.(string)
		request.ComputeCapacityTopologyId = &tmp
	}

	if computeHpcIslandId, ok := s.D.GetOkExists("compute_hpc_island_id"); ok {
		tmp := computeHpcIslandId.(string)
		request.ComputeHpcIslandId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeCapacityTopologyComputeNetworkBlocks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeCapacityTopologyComputeNetworkBlocks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeCapacityTopologyComputeNetworkBlocksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreComputeCapacityTopologyComputeNetworkBlocksDataSource-", CoreComputeCapacityTopologyComputeNetworkBlocksDataSource(), s.D))
	resources := []map[string]interface{}{}
	computeCapacityTopologyComputeNetworkBlock := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ComputeNetworkBlockSummaryToMap(item))
	}
	computeCapacityTopologyComputeNetworkBlock["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CoreComputeCapacityTopologyComputeNetworkBlocksDataSource().Schema["compute_network_block_collection"].Elem.(*schema.Resource).Schema)
		computeCapacityTopologyComputeNetworkBlock["items"] = items
	}

	resources = append(resources, computeCapacityTopologyComputeNetworkBlock)
	if err := s.D.Set("compute_network_block_collection", resources); err != nil {
		return err
	}

	return nil
}

func ComputeNetworkBlockSummaryToMap(obj oci_core.ComputeNetworkBlockSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComputeCapacityTopologyId != nil {
		result["compute_capacity_topology_id"] = string(*obj.ComputeCapacityTopologyId)
	}

	if obj.ComputeHpcIslandId != nil {
		result["compute_hpc_island_id"] = string(*obj.ComputeHpcIslandId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TotalComputeBareMetalHostCount != nil {
		result["total_compute_bare_metal_host_count"] = strconv.FormatInt(*obj.TotalComputeBareMetalHostCount, 10)
	}

	return result
}
