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

func CoreShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreShapes,
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
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"billing_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gpu_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gpus": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_billed_for_stopped_instance": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_flexible": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_live_migration_supported": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_subcore": {
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
									"max_per_numa_node_in_gbs": {
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
						"network_ports": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"networking_bandwidth_in_gbps": {
							Type:     schema.TypeFloat,
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
									"max_per_numa_node": {
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
						"platform_config_options": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"access_control_service_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeBool,
													},
												},
												"is_default_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"input_output_memory_management_unit_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeBool,
													},
												},
												"is_default_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"measured_boot_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeBool,
													},
												},
												"is_default_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"memory_encryption_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeBool,
													},
												},
												"is_default_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"numa_nodes_per_socket_platform_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"default_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"percentage_of_cores_enabled_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"default_value": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"max": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"min": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"secure_boot_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeBool,
													},
												},
												"is_default_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"symmetric_multi_threading_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeBool,
													},
												},
												"is_default_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"trusted_platform_module_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeBool,
													},
												},
												"is_default_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"virtual_instructions_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeBool,
													},
												},
												"is_default_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"processor_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quota_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"rdma_bandwidth_in_gbps": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rdma_ports": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"recommended_alternatives": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"shape_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"resize_compatible_shapes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreShapesDataSource-", CoreShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		shape := map[string]interface{}{}

		bous := []string{}
		for _, bou := range r.BaselineOcpuUtilizations {
			bous = append(bous, string(bou))
		}

		shape["baseline_ocpu_utilizations"] = bous

		shape["billing_type"] = r.BillingType

		if r.GpuDescription != nil {
			shape["gpu_description"] = *r.GpuDescription
		}

		if r.Gpus != nil {
			shape["gpus"] = *r.Gpus
		}

		if r.IsBilledForStoppedInstance != nil {
			shape["is_billed_for_stopped_instance"] = *r.IsBilledForStoppedInstance
		}

		if r.IsFlexible != nil {
			shape["is_flexible"] = *r.IsFlexible
		}

		if r.IsLiveMigrationSupported != nil {
			shape["is_live_migration_supported"] = *r.IsLiveMigrationSupported
		}

		if r.IsSubcore != nil {
			shape["is_subcore"] = *r.IsSubcore
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

		if r.NetworkPorts != nil {
			shape["network_ports"] = *r.NetworkPorts
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

		if r.PlatformConfigOptions != nil {
			shape["platform_config_options"] = []interface{}{ShapePlatformConfigOptionsToMap(r.PlatformConfigOptions)}
		} else {
			shape["platform_config_options"] = nil
		}

		if r.ProcessorDescription != nil {
			shape["processor_description"] = *r.ProcessorDescription
		}

		shape["quota_names"] = r.QuotaNames

		if r.RdmaBandwidthInGbps != nil {
			shape["rdma_bandwidth_in_gbps"] = *r.RdmaBandwidthInGbps
		}

		if r.RdmaPorts != nil {
			shape["rdma_ports"] = *r.RdmaPorts
		}

		recommendedAlternatives := []interface{}{}
		for _, item := range r.RecommendedAlternatives {
			recommendedAlternatives = append(recommendedAlternatives, ShapeAlternativeObjectToMap(item))
		}
		shape["recommended_alternatives"] = recommendedAlternatives

		shape["resize_compatible_shapes"] = r.ResizeCompatibleShapes

		resources = append(resources, shape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreShapesDataSource().Schema["shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("shapes", resources); err != nil {
		return err
	}

	return nil
}

func PercentageOfCoresEnabledOptionsToMap(obj *oci_core.PercentageOfCoresEnabledOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultValue != nil {
		result["default_value"] = int(*obj.DefaultValue)
	}

	if obj.Max != nil {
		result["max"] = int(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = int(*obj.Min)
	}

	return result
}

func ShapeAccessControlServiceEnabledPlatformOptionsToMap(obj *oci_core.ShapeAccessControlServiceEnabledPlatformOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_values"] = obj.AllowedValues

	if obj.IsDefaultEnabled != nil {
		result["is_default_enabled"] = bool(*obj.IsDefaultEnabled)
	}

	return result
}

func ShapeAlternativeObjectToMap(obj oci_core.ShapeAlternativeObject) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ShapeName != nil {
		result["shape_name"] = string(*obj.ShapeName)
	}

	return result
}

func ShapeInputOutputMemoryManagementUnitEnabledPlatformOptionsToMap(obj *oci_core.ShapeInputOutputMemoryManagementUnitEnabledPlatformOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_values"] = obj.AllowedValues

	if obj.IsDefaultEnabled != nil {
		result["is_default_enabled"] = bool(*obj.IsDefaultEnabled)
	}

	return result
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

func ShapeMeasuredBootOptionsToMap(obj *oci_core.ShapeMeasuredBootOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_values"] = obj.AllowedValues

	if obj.IsDefaultEnabled != nil {
		result["is_default_enabled"] = bool(*obj.IsDefaultEnabled)
	}

	return result
}

func ShapeMemoryEncryptionOptionsToMap(obj *oci_core.ShapeMemoryEncryptionOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_values"] = obj.AllowedValues
	result["allowed_values"] = obj.AllowedValues

	if obj.IsDefaultEnabled != nil {
		result["is_default_enabled"] = bool(*obj.IsDefaultEnabled)
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

	if obj.MaxPerNumaNodeInGBs != nil {
		result["max_per_numa_node_in_gbs"] = float32(*obj.MaxPerNumaNodeInGBs)
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

func ShapeNumaNodesPerSocketPlatformOptionsToMap(obj *oci_core.ShapeNumaNodesPerSocketPlatformOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_values"] = obj.AllowedValues

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	return result
}

func ShapeOcpuOptionsToMap(obj *oci_core.ShapeOcpuOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Max != nil {
		result["max"] = float32(*obj.Max)
	}

	if obj.MaxPerNumaNode != nil {
		result["max_per_numa_node"] = float32(*obj.MaxPerNumaNode)
	}

	if obj.Min != nil {
		result["min"] = float32(*obj.Min)
	}

	return result
}

func ShapePlatformConfigOptionsToMap(obj *oci_core.ShapePlatformConfigOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccessControlServiceOptions != nil {
		result["access_control_service_options"] = []interface{}{ShapeAccessControlServiceEnabledPlatformOptionsToMap(obj.AccessControlServiceOptions)}
	}

	if obj.InputOutputMemoryManagementUnitOptions != nil {
		result["input_output_memory_management_unit_options"] = []interface{}{ShapeInputOutputMemoryManagementUnitEnabledPlatformOptionsToMap(obj.InputOutputMemoryManagementUnitOptions)}
	}

	if obj.MeasuredBootOptions != nil {
		result["measured_boot_options"] = []interface{}{ShapeMeasuredBootOptionsToMap(obj.MeasuredBootOptions)}
	}

	if obj.MemoryEncryptionOptions != nil {
		result["memory_encryption_options"] = []interface{}{ShapeMemoryEncryptionOptionsToMap(obj.MemoryEncryptionOptions)}
	}

	if obj.NumaNodesPerSocketPlatformOptions != nil {
		result["numa_nodes_per_socket_platform_options"] = []interface{}{ShapeNumaNodesPerSocketPlatformOptionsToMap(obj.NumaNodesPerSocketPlatformOptions)}
	}

	if obj.PercentageOfCoresEnabledOptions != nil {
		result["percentage_of_cores_enabled_options"] = []interface{}{PercentageOfCoresEnabledOptionsToMap(obj.PercentageOfCoresEnabledOptions)}
	}

	if obj.SecureBootOptions != nil {
		result["secure_boot_options"] = []interface{}{ShapeSecureBootOptionsToMap(obj.SecureBootOptions)}
	}

	if obj.SymmetricMultiThreadingOptions != nil {
		result["symmetric_multi_threading_options"] = []interface{}{ShapeSymmetricMultiThreadingEnabledPlatformOptionsToMap(obj.SymmetricMultiThreadingOptions)}
	}

	if obj.TrustedPlatformModuleOptions != nil {
		result["trusted_platform_module_options"] = []interface{}{ShapeTrustedPlatformModuleOptionsToMap(obj.TrustedPlatformModuleOptions)}
	}

	result["type"] = string(obj.Type)

	if obj.VirtualInstructionsOptions != nil {
		result["virtual_instructions_options"] = []interface{}{ShapeVirtualInstructionsEnabledPlatformOptionsToMap(obj.VirtualInstructionsOptions)}
	}

	return result
}

func ShapeSecureBootOptionsToMap(obj *oci_core.ShapeSecureBootOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_values"] = obj.AllowedValues

	if obj.IsDefaultEnabled != nil {
		result["is_default_enabled"] = bool(*obj.IsDefaultEnabled)
	}

	return result
}

func ShapeSymmetricMultiThreadingEnabledPlatformOptionsToMap(obj *oci_core.ShapeSymmetricMultiThreadingEnabledPlatformOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_values"] = obj.AllowedValues

	if obj.IsDefaultEnabled != nil {
		result["is_default_enabled"] = bool(*obj.IsDefaultEnabled)
	}

	return result
}

func ShapeTrustedPlatformModuleOptionsToMap(obj *oci_core.ShapeTrustedPlatformModuleOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_values"] = obj.AllowedValues

	if obj.IsDefaultEnabled != nil {
		result["is_default_enabled"] = bool(*obj.IsDefaultEnabled)
	}

	return result
}

func ShapeVirtualInstructionsEnabledPlatformOptionsToMap(obj *oci_core.ShapeVirtualInstructionsEnabledPlatformOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_values"] = obj.AllowedValues

	if obj.IsDefaultEnabled != nil {
		result["is_default_enabled"] = bool(*obj.IsDefaultEnabled)
	}

	return result
}
