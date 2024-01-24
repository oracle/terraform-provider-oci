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

func CoreComputeCapacityTopologyComputeBareMetalHostsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeCapacityTopologyComputeBareMetalHosts,
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
			"compute_local_block_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_network_block_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_bare_metal_host_collection": {
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
									"compute_local_block_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compute_network_block_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_shape": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readCoreComputeCapacityTopologyComputeBareMetalHosts(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityTopologyComputeBareMetalHostsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeCapacityTopologyComputeBareMetalHostsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeCapacityTopologyComputeBareMetalHostsResponse
}

func (s *CoreComputeCapacityTopologyComputeBareMetalHostsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeCapacityTopologyComputeBareMetalHostsDataSourceCrud) Get() error {
	request := oci_core.ListComputeCapacityTopologyComputeBareMetalHostsRequest{}

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

	if computeLocalBlockId, ok := s.D.GetOkExists("compute_local_block_id"); ok {
		tmp := computeLocalBlockId.(string)
		request.ComputeLocalBlockId = &tmp
	}

	if computeNetworkBlockId, ok := s.D.GetOkExists("compute_network_block_id"); ok {
		tmp := computeNetworkBlockId.(string)
		request.ComputeNetworkBlockId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeCapacityTopologyComputeBareMetalHosts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeCapacityTopologyComputeBareMetalHosts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeCapacityTopologyComputeBareMetalHostsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreComputeCapacityTopologyComputeBareMetalHostsDataSource-", CoreComputeCapacityTopologyComputeBareMetalHostsDataSource(), s.D))
	resources := []map[string]interface{}{}
	computeCapacityTopologyComputeBareMetalHost := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ComputeBareMetalHostSummaryToMap(item))
	}
	computeCapacityTopologyComputeBareMetalHost["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CoreComputeCapacityTopologyComputeBareMetalHostsDataSource().Schema["compute_bare_metal_host_collection"].Elem.(*schema.Resource).Schema)
		computeCapacityTopologyComputeBareMetalHost["items"] = items
	}

	resources = append(resources, computeCapacityTopologyComputeBareMetalHost)
	if err := s.D.Set("compute_bare_metal_host_collection", resources); err != nil {
		return err
	}

	return nil
}

func ComputeBareMetalHostSummaryToMap(obj oci_core.ComputeBareMetalHostSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComputeCapacityTopologyId != nil {
		result["compute_capacity_topology_id"] = string(*obj.ComputeCapacityTopologyId)
	}

	if obj.ComputeHpcIslandId != nil {
		result["compute_hpc_island_id"] = string(*obj.ComputeHpcIslandId)
	}

	if obj.ComputeLocalBlockId != nil {
		result["compute_local_block_id"] = string(*obj.ComputeLocalBlockId)
	}

	if obj.ComputeNetworkBlockId != nil {
		result["compute_network_block_id"] = string(*obj.ComputeNetworkBlockId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InstanceId != nil {
		result["instance_id"] = string(*obj.InstanceId)
	}

	if obj.InstanceShape != nil {
		result["instance_shape"] = string(*obj.InstanceShape)
	}

	result["lifecycle_details"] = string(obj.LifecycleDetails)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
