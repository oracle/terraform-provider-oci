// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerenginePodShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerenginePodShapes,
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
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pod_shapes": {
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
						"network_bandwidth_options": {
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
									"max": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"min": {
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
	}
}

func readContainerenginePodShapes(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerenginePodShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerenginePodShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListPodShapesResponse
}

func (s *ContainerenginePodShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerenginePodShapesDataSourceCrud) Get() error {
	request := oci_containerengine.ListPodShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.ListPodShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPodShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ContainerenginePodShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerenginePodShapesDataSource-", ContainerenginePodShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		podShape := map[string]interface{}{}

		memoryOptions := []interface{}{}
		for _, item := range r.MemoryOptions {
			memoryOptions = append(memoryOptions, ShapeMemoryOptionsToMap(item))
		}
		podShape["memory_options"] = memoryOptions

		if r.Name != nil {
			podShape["name"] = *r.Name
		}

		networkBandwidthOptions := []interface{}{}
		for _, item := range r.NetworkBandwidthOptions {
			networkBandwidthOptions = append(networkBandwidthOptions, ShapeNetworkBandwidthOptionsToMap(item))
		}
		podShape["network_bandwidth_options"] = networkBandwidthOptions

		ocpuOptions := []interface{}{}
		for _, item := range r.OcpuOptions {
			ocpuOptions = append(ocpuOptions, ShapeOcpuOptionsToMap(item))
		}
		podShape["ocpu_options"] = ocpuOptions

		if r.ProcessorDescription != nil {
			podShape["processor_description"] = *r.ProcessorDescription
		}

		resources = append(resources, podShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ContainerenginePodShapesDataSource().Schema["pod_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("pod_shapes", resources); err != nil {
		return err
	}

	return nil
}

func ShapeMemoryOptionsToMap(obj oci_containerengine.ShapeMemoryOptions) map[string]interface{} {
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

func ShapeNetworkBandwidthOptionsToMap(obj oci_containerengine.ShapeNetworkBandwidthOptions) map[string]interface{} {
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

func ShapeOcpuOptionsToMap(obj oci_containerengine.ShapeOcpuOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Max != nil {
		result["max"] = float32(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = float32(*obj.Min)
	}

	return result
}
