// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiDedicatedAiClusterAllocatableGpusDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readGenerativeAiDedicatedAiClusterAllocatableGpusWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"capacity_reservation_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_shape": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"allocatable_gpu_collection": {
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
									"compute_shape": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"gpu_allocations": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"available_gpus": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"used_gpus": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"max_allocatable_gpu_units": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"total_available_gpu_units": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"total_used_gpu_units": {
										Type:     schema.TypeInt,
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

func readGenerativeAiDedicatedAiClusterAllocatableGpusWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiDedicatedAiClusterAllocatableGpusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GenerativeAiDedicatedAiClusterAllocatableGpusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.ListDedicatedAiClustersAllocatableGpusResponse
}

func (s *GenerativeAiDedicatedAiClusterAllocatableGpusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiDedicatedAiClusterAllocatableGpusDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.ListDedicatedAiClustersAllocatableGpusRequest{}

	if capacityReservationId, ok := s.D.GetOkExists("capacity_reservation_id"); ok {
		tmp := capacityReservationId.(string)
		request.CapacityReservationId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeShape, ok := s.D.GetOkExists("compute_shape"); ok {
		tmp := computeShape.(string)
		request.ComputeShape = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.ListDedicatedAiClustersAllocatableGpus(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDedicatedAiClustersAllocatableGpus(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GenerativeAiDedicatedAiClusterAllocatableGpusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GenerativeAiDedicatedAiClusterAllocatableGpusDataSource-", GenerativeAiDedicatedAiClusterAllocatableGpusDataSource(), s.D))
	resources := []map[string]interface{}{}
	dedicatedAiClusterAllocatableGpu := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AllocatableGpuSummaryToMap(item))
	}
	dedicatedAiClusterAllocatableGpu["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GenerativeAiDedicatedAiClusterAllocatableGpusDataSource().Schema["allocatable_gpu_collection"].Elem.(*schema.Resource).Schema)
		dedicatedAiClusterAllocatableGpu["items"] = items
	}

	resources = append(resources, dedicatedAiClusterAllocatableGpu)
	if err := s.D.Set("allocatable_gpu_collection", resources); err != nil {
		return err
	}

	return nil
}

func AllocatableGpuSummaryToMap(obj oci_generative_ai.AllocatableGpuSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComputeShape != nil {
		result["compute_shape"] = string(*obj.ComputeShape)
	}

	gpuAllocations := []interface{}{}
	for _, item := range obj.GpuAllocations {
		gpuAllocations = append(gpuAllocations, GpuAllocationToMap(item))
	}
	result["gpu_allocations"] = gpuAllocations

	if obj.MaxAllocatableGpuUnits != nil {
		result["max_allocatable_gpu_units"] = int(*obj.MaxAllocatableGpuUnits)
	}

	if obj.TotalAvailableGpuUnits != nil {
		result["total_available_gpu_units"] = int(*obj.TotalAvailableGpuUnits)
	}

	if obj.TotalUsedGpuUnits != nil {
		result["total_used_gpu_units"] = int(*obj.TotalUsedGpuUnits)
	}

	return result
}

func GpuAllocationToMap(obj oci_generative_ai.GpuAllocation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailableGpus != nil {
		result["available_gpus"] = int(*obj.AvailableGpus)
	}

	if obj.UsedGpus != nil {
		result["used_gpus"] = int(*obj.UsedGpus)
	}

	return result
}
