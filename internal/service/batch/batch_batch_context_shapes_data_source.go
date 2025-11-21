// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package batch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BatchBatchContextShapesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readBatchBatchContextShapesWithContext,
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
			"batch_context_shape_collection": {
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
									"memory_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"default_per_ocpu_in_gbs": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"max_in_gbs": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"max_per_ocpu_in_gbs": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"min_in_gbs": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"min_per_ocpu_in_gbs": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"networking_bandwidth_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"default_per_ocpu_in_gbps": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"max_in_gbps": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"min_in_gbps": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"ocpu_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"max_ocpus": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"min_ocpus": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"processor_description": {
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

func readBatchBatchContextShapesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchContextShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BatchBatchContextShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_batch.BatchComputingClient
	Res    *oci_batch.ListBatchContextShapesResponse
}

func (s *BatchBatchContextShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BatchBatchContextShapesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_batch.ListBatchContextShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "batch")

	response, err := s.Client.ListBatchContextShapes(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBatchContextShapes(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BatchBatchContextShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BatchBatchContextShapesDataSource-", BatchBatchContextShapesDataSource(), s.D))
	resources := []map[string]interface{}{}
	batchContextShape := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BatchContextShapeSummaryToMap(item))
	}
	batchContextShape["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, BatchBatchContextShapesDataSource().Schema["batch_context_shape_collection"].Elem.(*schema.Resource).Schema)
		batchContextShape["items"] = items
	}

	resources = append(resources, batchContextShape)
	if err := s.D.Set("batch_context_shape_collection", resources); err != nil {
		return err
	}

	return nil
}

func BatchContextShapeSummaryToMap(obj oci_batch.BatchContextShapeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryOptions != nil {
		result["memory_options"] = []interface{}{ShapeMemoryOptionsToMap(obj.MemoryOptions)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NetworkingBandwidthOptions != nil {
		result["networking_bandwidth_options"] = []interface{}{ShapeNetworkingBandwidthOptionsToMap(obj.NetworkingBandwidthOptions)}
	}

	if obj.OcpuOptions != nil {
		result["ocpu_options"] = []interface{}{ShapeOcpuOptionsToMap(obj.OcpuOptions)}
	}

	if obj.ProcessorDescription != nil {
		result["processor_description"] = string(*obj.ProcessorDescription)
	}

	return result
}

func ShapeMemoryOptionsToMap(obj *oci_batch.ShapeMemoryOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultPerOcpuInGBs != nil {
		result["default_per_ocpu_in_gbs"] = float32(*obj.DefaultPerOcpuInGBs)
	}

	if obj.MaxInGBs != nil {
		result["max_in_gbs"] = float32(*obj.MaxInGBs)
	}

	if obj.MaxPerOcpuInGBs != nil {
		result["max_per_ocpu_in_gbs"] = float32(*obj.MaxPerOcpuInGBs)
	}

	if obj.MinInGBs != nil {
		result["min_in_gbs"] = float32(*obj.MinInGBs)
	}

	if obj.MinPerOcpuInGBs != nil {
		result["min_per_ocpu_in_gbs"] = float32(*obj.MinPerOcpuInGBs)
	}

	return result
}

func ShapeNetworkingBandwidthOptionsToMap(obj *oci_batch.ShapeNetworkingBandwidthOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultPerOcpuInGbps != nil {
		result["default_per_ocpu_in_gbps"] = float32(*obj.DefaultPerOcpuInGbps)
	}

	if obj.MaxInGbps != nil {
		result["max_in_gbps"] = float32(*obj.MaxInGbps)
	}

	if obj.MinInGbps != nil {
		result["min_in_gbps"] = float32(*obj.MinInGbps)
	}

	return result
}

func ShapeOcpuOptionsToMap(obj *oci_batch.ShapeOcpuOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxOcpus != nil {
		result["max_ocpus"] = float32(*obj.MaxOcpus)
	}

	if obj.MinOcpus != nil {
		result["min_ocpus"] = float32(*obj.MinOcpus)
	}

	return result
}
