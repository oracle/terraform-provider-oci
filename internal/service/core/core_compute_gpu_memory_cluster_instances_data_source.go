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

func CoreComputeGpuMemoryClusterInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeGpuMemoryClusterInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compute_gpu_memory_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_gpu_memory_cluster_instance_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"availability_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fault_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_configuration_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_shape": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"region": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readCoreComputeGpuMemoryClusterInstances(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeGpuMemoryClusterInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeGpuMemoryClusterInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeGpuMemoryClusterInstancesResponse
}

func (s *CoreComputeGpuMemoryClusterInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeGpuMemoryClusterInstancesDataSourceCrud) Get() error {
	request := oci_core.ListComputeGpuMemoryClusterInstancesRequest{}

	if computeGpuMemoryClusterId, ok := s.D.GetOkExists("compute_gpu_memory_cluster_id"); ok {
		tmp := computeGpuMemoryClusterId.(string)
		request.ComputeGpuMemoryClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeGpuMemoryClusterInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeGpuMemoryClusterInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeGpuMemoryClusterInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreComputeGpuMemoryClusterInstancesDataSource-", CoreComputeGpuMemoryClusterInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	computeGpuMemoryClusterInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ComputeGpuMemoryClusterInstanceSummaryToMap(item))
	}
	computeGpuMemoryClusterInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CoreComputeGpuMemoryClusterInstancesDataSource().Schema["compute_gpu_memory_cluster_instance_collection"].Elem.(*schema.Resource).Schema)
		computeGpuMemoryClusterInstance["items"] = items
	}

	resources = append(resources, computeGpuMemoryClusterInstance)
	if err := s.D.Set("compute_gpu_memory_cluster_instance_collection", resources); err != nil {
		return err
	}

	return nil
}

func ComputeGpuMemoryClusterInstanceSummaryToMap(obj oci_core.ComputeGpuMemoryClusterInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InstanceConfigurationId != nil {
		result["instance_configuration_id"] = string(*obj.InstanceConfigurationId)
	}

	if obj.InstanceShape != nil {
		result["instance_shape"] = string(*obj.InstanceShape)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
