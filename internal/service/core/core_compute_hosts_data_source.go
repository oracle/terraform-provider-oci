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

func CoreComputeHostsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeHosts,
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
			"compute_host_health": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_host_lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_host_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"compartment_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"availability_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"capacity_reservation_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fault_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									// Optional
									"has_impacted_components": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"health": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"hpc_island_id": {
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
									"lifecycle_details": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"local_block_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"gpu_memory_fabric_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_block_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"shape": {
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

func readCoreComputeHosts(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeHostsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeHostsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeHostsResponse
}

func (s *CoreComputeHostsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeHostsDataSourceCrud) Get() error {
	request := oci_core.ListComputeHostsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeHostHealth, ok := s.D.GetOkExists("compute_host_health"); ok {
		tmp := computeHostHealth.(string)
		request.ComputeHostHealth = &tmp
	}

	if computeHostLifecycleState, ok := s.D.GetOkExists("compute_host_lifecycle_state"); ok {
		tmp := computeHostLifecycleState.(string)
		request.ComputeHostLifecycleState = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if networkResourceId, ok := s.D.GetOkExists("network_resource_id"); ok {
		tmp := networkResourceId.(string)
		request.NetworkResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeHosts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeHosts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeHostsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreComputeHostsDataSource-", CoreComputeHostsDataSource(), s.D))
	resources := []map[string]interface{}{}
	computeHost := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ComputeHostSummaryToMap(item))
	}
	computeHost["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CoreComputeHostsDataSource().Schema["compute_host_collection"].Elem.(*schema.Resource).Schema)
		computeHost["items"] = items
	}

	resources = append(resources, computeHost)
	if err := s.D.Set("compute_host_collection", resources); err != nil {
		return err
	}

	return nil
}

func ComputeHostSummaryToMap(obj oci_core.ComputeHostSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CapacityReservationId != nil {
		result["capacity_reservation_id"] = string(*obj.CapacityReservationId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HasImpactedComponents != nil {
		result["has_impacted_components"] = bool(*obj.HasImpactedComponents)
	}

	result["health"] = string(obj.Health)

	if obj.HpcIslandId != nil {
		result["hpc_island_id"] = string(*obj.HpcIslandId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InstanceId != nil {
		result["instance_id"] = string(*obj.InstanceId)
	}

	if obj.LocalBlockId != nil {
		result["local_block_id"] = string(*obj.LocalBlockId)
	}

	if obj.GpuMemoryFabricId != nil {
		result["gpu_memory_fabric_id"] = string(*obj.GpuMemoryFabricId)
	}

	if obj.NetworkBlockId != nil {
		result["network_block_id"] = string(*obj.NetworkBlockId)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
