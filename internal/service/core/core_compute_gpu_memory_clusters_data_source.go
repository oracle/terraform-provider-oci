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

func CoreComputeGpuMemoryClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeGpuMemoryClusters,
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
			"compute_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_gpu_memory_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_gpu_memory_cluster_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CoreComputeGpuMemoryClusterResource()),
						},
					},
				},
			},
		},
	}
}

func readCoreComputeGpuMemoryClusters(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeGpuMemoryClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeGpuMemoryClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeGpuMemoryClustersResponse
}

func (s *CoreComputeGpuMemoryClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeGpuMemoryClustersDataSourceCrud) Get() error {
	request := oci_core.ListComputeGpuMemoryClustersRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeClusterId, ok := s.D.GetOkExists("compute_cluster_id"); ok {
		tmp := computeClusterId.(string)
		request.ComputeClusterId = &tmp
	}

	if computeGpuMemoryClusterId, ok := s.D.GetOkExists("id"); ok {
		tmp := computeGpuMemoryClusterId.(string)
		request.ComputeGpuMemoryClusterId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeGpuMemoryClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeGpuMemoryClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeGpuMemoryClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreComputeGpuMemoryClustersDataSource-", CoreComputeGpuMemoryClustersDataSource(), s.D))
	resources := []map[string]interface{}{}
	computeGpuMemoryCluster := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ComputeGpuMemoryClusterSummaryToMap(item))
	}
	computeGpuMemoryCluster["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CoreComputeGpuMemoryClustersDataSource().Schema["compute_gpu_memory_cluster_collection"].Elem.(*schema.Resource).Schema)
		computeGpuMemoryCluster["items"] = items
	}

	resources = append(resources, computeGpuMemoryCluster)
	if err := s.D.Set("compute_gpu_memory_cluster_collection", resources); err != nil {
		return err
	}

	return nil
}
