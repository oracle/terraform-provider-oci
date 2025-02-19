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

func CoreComputeGpuMemoryFabricsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeGpuMemoryFabrics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_gpu_memory_fabric_health": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_gpu_memory_fabric_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_gpu_memory_fabric_lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_hpc_island_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_network_block_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_gpu_memory_fabric_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CoreComputeGpuMemoryFabricResource()),
						},
					},
				},
			},
		},
	}
}

func readCoreComputeGpuMemoryFabrics(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeGpuMemoryFabricsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeGpuMemoryFabricsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeGpuMemoryFabricsResponse
}

func (s *CoreComputeGpuMemoryFabricsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeGpuMemoryFabricsDataSourceCrud) Get() error {
	request := oci_core.ListComputeGpuMemoryFabricsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeGpuMemoryFabricHealth, ok := s.D.GetOkExists("compute_gpu_memory_fabric_health"); ok {
		request.ComputeGpuMemoryFabricHealth = oci_core.ComputeGpuMemoryFabricFabricHealthEnum(computeGpuMemoryFabricHealth.(string))
	}

	if computeGpuMemoryFabricId, ok := s.D.GetOkExists("compute_gpu_memory_fabric_id"); ok {
		tmp := computeGpuMemoryFabricId.(string)
		request.ComputeGpuMemoryFabricId = &tmp
	}

	if computeGpuMemoryFabricLifecycleState, ok := s.D.GetOkExists("compute_gpu_memory_fabric_lifecycle_state"); ok {
		request.ComputeGpuMemoryFabricLifecycleState = oci_core.ComputeGpuMemoryFabricLifecycleStateEnum(computeGpuMemoryFabricLifecycleState.(string))
	}

	if computeHpcIslandId, ok := s.D.GetOkExists("compute_hpc_island_id"); ok {
		tmp := computeHpcIslandId.(string)
		request.ComputeHpcIslandId = &tmp
	}

	if computeNetworkBlockId, ok := s.D.GetOkExists("compute_network_block_id"); ok {
		tmp := computeNetworkBlockId.(string)
		request.ComputeNetworkBlockId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeGpuMemoryFabrics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeGpuMemoryFabrics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeGpuMemoryFabricsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreComputeGpuMemoryFabricsDataSource-", CoreComputeGpuMemoryFabricsDataSource(), s.D))
	resources := []map[string]interface{}{}
	computeGpuMemoryFabric := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ComputeGpuMemoryFabricSummaryToMap(item))
	}
	computeGpuMemoryFabric["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CoreComputeGpuMemoryFabricsDataSource().Schema["compute_gpu_memory_fabric_collection"].Elem.(*schema.Resource).Schema)
		computeGpuMemoryFabric["items"] = items
	}

	resources = append(resources, computeGpuMemoryFabric)
	if err := s.D.Set("compute_gpu_memory_fabric_collection", resources); err != nil {
		return err
	}

	return nil
}
