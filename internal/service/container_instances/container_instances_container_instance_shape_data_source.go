// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package container_instances

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_container_instances "github.com/oracle/oci-go-sdk/v65/containerinstances"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerInstancesContainerInstanceShapeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularContainerInstancesContainerInstanceShape,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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

func readSingularContainerInstancesContainerInstanceShape(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerInstancesContainerInstanceShapeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerInstanceClient()

	return tfresource.ReadResource(sync)
}

type ContainerInstancesContainerInstanceShapeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_container_instances.ContainerInstanceClient
	Res    *oci_container_instances.ListContainerInstanceShapesResponse
}

func (s *ContainerInstancesContainerInstanceShapeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerInstancesContainerInstanceShapeDataSourceCrud) Get() error {
	request := oci_container_instances.ListContainerInstanceShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerinstance")

	response, err := s.Client.ListContainerInstanceShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerInstancesContainerInstanceShapeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerInstancesContainerInstanceShapeDataSource-", ContainerInstancesContainerInstanceShapeDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ContainerInstanceShapeSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func ContainerInstanceShapeSummaryToMap(obj oci_container_instances.ContainerInstanceShapeSummary) map[string]interface{} {
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

func ShapeMemoryOptionsToMap(obj *oci_container_instances.ShapeMemoryOptions) map[string]interface{} {
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

func ShapeNetworkingBandwidthOptionsToMap(obj *oci_container_instances.ShapeNetworkingBandwidthOptions) map[string]interface{} {
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

func ShapeOcpuOptionsToMap(obj *oci_container_instances.ShapeOcpuOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Max != nil {
		result["max"] = float32(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = float32(*obj.Min)
	}

	return result
}
