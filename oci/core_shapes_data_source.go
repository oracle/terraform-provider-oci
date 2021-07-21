// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v45/core"
)

func init() {
	RegisterDatasource("oci_core_shapes", CoreShapesDataSource())
}

func CoreShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreShapes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"baseline_ocpu_utilizations": {
							Type:     schema.TypeList,
							Computed: true,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"gpu_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gpus": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_live_migration_supported": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"local_disk_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"local_disks": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"local_disks_total_size_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"max_vnic_attachment_options": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"default_per_ocpu": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"max": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"min": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"max_vnic_attachments": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"memory_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"memory_options": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
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
						"min_total_baseline_ocpus_required": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"networking_bandwidth_in_gbps": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"networking_bandwidth_options": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
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
							MaxItems: 1,
							MinItems: 1,
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
						"ocpus": {
							Type:     schema.TypeFloat,
							Computed: true,
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

func readCoreShapes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return ReadResource(sync)
}

type CoreShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListShapesResponse
}

func (s *CoreShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreShapesDataSourceCrud) Get() error {
	request := oci_core.ListShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreShapesDataSource-", CoreShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		shape := map[string]interface{}{}

		bous := []string{}
		for _, bou := range r.BaselineOcpuUtilizations {
			bous = append(bous, string(bou))
		}

		shape["baseline_ocpu_utilizations"] = bous

		if r.GpuDescription != nil {
			shape["gpu_description"] = *r.GpuDescription
		}

		if r.Gpus != nil {
			shape["gpus"] = *r.Gpus
		}

		if r.IsLiveMigrationSupported != nil {
			shape["is_live_migration_supported"] = *r.IsLiveMigrationSupported
		}

		if r.LocalDiskDescription != nil {
			shape["local_disk_description"] = *r.LocalDiskDescription
		}

		if r.LocalDisks != nil {
			shape["local_disks"] = *r.LocalDisks
		}

		if r.LocalDisksTotalSizeInGBs != nil {
			shape["local_disks_total_size_in_gbs"] = *r.LocalDisksTotalSizeInGBs
		}

		if r.MaxVnicAttachmentOptions != nil {
			shape["max_vnic_attachment_options"] = []interface{}{ShapeMaxVnicAttachmentOptionsToMap(r.MaxVnicAttachmentOptions)}
		} else {
			shape["max_vnic_attachment_options"] = nil
		}

		if r.MaxVnicAttachments != nil {
			shape["max_vnic_attachments"] = *r.MaxVnicAttachments
		}

		if r.MemoryInGBs != nil {
			shape["memory_in_gbs"] = *r.MemoryInGBs
		}

		if r.MemoryOptions != nil {
			shape["memory_options"] = []interface{}{ShapeMemoryOptionsToMap(r.MemoryOptions)}
		} else {
			shape["memory_options"] = nil
		}

		if r.MinTotalBaselineOcpusRequired != nil {
			shape["min_total_baseline_ocpus_required"] = *r.MinTotalBaselineOcpusRequired
		}

		if r.Shape != nil {
			shape["name"] = *r.Shape
		}

		if r.NetworkingBandwidthInGbps != nil {
			shape["networking_bandwidth_in_gbps"] = *r.NetworkingBandwidthInGbps
		}

		if r.NetworkingBandwidthOptions != nil {
			shape["networking_bandwidth_options"] = []interface{}{ShapeNetworkingBandwidthOptionsToMap(r.NetworkingBandwidthOptions)}
		} else {
			shape["networking_bandwidth_options"] = nil
		}

		if r.OcpuOptions != nil {
			shape["ocpu_options"] = []interface{}{ShapeOcpuOptionsToMap(r.OcpuOptions)}
		} else {
			shape["ocpu_options"] = nil
		}

		if r.Ocpus != nil {
			shape["ocpus"] = *r.Ocpus
		}

		if r.ProcessorDescription != nil {
			shape["processor_description"] = *r.ProcessorDescription
		}

		resources = append(resources, shape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreShapesDataSource().Schema["shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("shapes", resources); err != nil {
		return err
	}

	return nil
}

func ShapeMaxVnicAttachmentOptionsToMap(obj *oci_core.ShapeMaxVnicAttachmentOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultPerOcpu != nil {
		result["default_per_ocpu"] = float32(*obj.DefaultPerOcpu)
	}

	if obj.Max != nil {
		result["max"] = float32(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = int(*obj.Min)
	}

	return result
}

func ShapeMemoryOptionsToMap(obj *oci_core.ShapeMemoryOptions) map[string]interface{} {
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

func ShapeNetworkingBandwidthOptionsToMap(obj *oci_core.ShapeNetworkingBandwidthOptions) map[string]interface{} {
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

func ShapeOcpuOptionsToMap(obj *oci_core.ShapeOcpuOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Max != nil {
		result["max"] = float32(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = float32(*obj.Min)
	}

	return result
}
