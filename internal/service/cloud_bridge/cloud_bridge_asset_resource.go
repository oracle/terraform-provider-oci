// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeAssetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudBridgeAsset,
		Read:     readCloudBridgeAsset,
		Update:   updateCloudBridgeAsset,
		Delete:   deleteCloudBridgeAsset,
		Schema: map[string]*schema.Schema{
			// Required
			"asset_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"VM",
					"VMWARE_VM",
				}, true),
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_asset_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"inventory_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"asset_source_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compute": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"connected_networks": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"cores_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"cpu_model": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"disks": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"boot_order": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"location": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"persistent_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"size_in_mbs": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ValidateFunc:     tfresource.ValidateInt64TypeString,
										DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
									},
									"uuid": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"uuid_lun": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"disks_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dns_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"firmware": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gpu_devices": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"cores_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"manufacturer": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"memory_in_mbs": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ValidateFunc:     tfresource.ValidateInt64TypeString,
										DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"gpu_devices_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"guest_state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hardware_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_pmem_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_tpm_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"latency_sensitivity": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"memory_in_mbs": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"nics": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"ip_addresses": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"label": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac_address": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac_address_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"network_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"switch_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"nics_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"nvdimm_controller": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"bus_number": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"label": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"nvdimms": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"controller_key": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"label": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"unit_number": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"operating_system": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"operating_system_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pmem_in_mbs": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"power_state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"primary_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scsi_controller": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"label": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"shared_bus": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"unit_number": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"storage_provisioned_in_mbs": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"threads_per_core_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"vm": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"hypervisor_host": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hypervisor_vendor": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hypervisor_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"vmware_vcenter": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"data_center": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vcenter_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vcenter_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"vmware_vm": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"cluster": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"customer_fields": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"customer_tags": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"fault_tolerance_bandwidth": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fault_tolerance_secondary_latency": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fault_tolerance_state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"instance_uuid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_disks_cbt_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_disks_uuid_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vmware_tools_status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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
	}
}

func createCloudBridgeAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InventoryClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudBridgeAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InventoryClient()

	return tfresource.ReadResource(sync)
}

func updateCloudBridgeAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InventoryClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudBridgeAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InventoryClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudBridgeAssetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_bridge.InventoryClient
	Res                    *oci_cloud_bridge.Asset
	DisableNotFoundRetries bool
}

func (s *CloudBridgeAssetResourceCrud) ID() string {
	asset := *s.Res
	return *asset.GetId()
}

func (s *CloudBridgeAssetResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *CloudBridgeAssetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_bridge.AssetLifecycleStateActive),
	}
}

func (s *CloudBridgeAssetResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *CloudBridgeAssetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_bridge.AssetLifecycleStateDeleted),
	}
}

func (s *CloudBridgeAssetResourceCrud) Create() error {
	request := oci_cloud_bridge.CreateAssetRequest{}
	err := s.populateTopLevelPolymorphicCreateAssetRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.CreateAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Asset
	return nil
}

func (s *CloudBridgeAssetResourceCrud) Get() error {
	request := oci_cloud_bridge.GetAssetRequest{}

	tmp := s.D.Id()
	request.AssetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.GetAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Asset
	return nil
}

func (s *CloudBridgeAssetResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_bridge.UpdateAssetRequest{}
	err := s.populateTopLevelPolymorphicUpdateAssetRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	_, err = s.Client.UpdateAsset(context.Background(), request)
	if err != nil {
		return err
	}

	changeAssetTagsRequest := oci_cloud_bridge.ChangeAssetTagsRequest{}
	tmp := s.D.Id()
	changeAssetTagsRequest.AssetId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		changeAssetTagsRequest.ChangeAssetTagsDetails.DefinedTags = convertedDefinedTags
	}
	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		changeAssetTagsRequest.ChangeAssetTagsDetails.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}
	changeAssetTagsResponse, err := s.Client.ChangeAssetTags(context.Background(), changeAssetTagsRequest)
	if err != nil {
		return err
	}
	s.Res = &changeAssetTagsResponse.Asset

	return nil
}

func (s *CloudBridgeAssetResourceCrud) Delete() error {
	request := oci_cloud_bridge.DeleteAssetRequest{}

	tmp := s.D.Id()
	request.AssetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	_, err := s.Client.DeleteAsset(context.Background(), request)
	return err
}

func (s *CloudBridgeAssetResourceCrud) SetData() error {

	switch v := (*s.Res).(type) {
	case oci_cloud_bridge.VmAsset:
		s.D.Set("asset_type", "VM")

		if v.Compute != nil {
			s.D.Set("compute", []interface{}{ComputePropertiesToMap(v.Compute)})
		} else {
			s.D.Set("compute", nil)
		}

		if v.Vm != nil {
			s.D.Set("vm", []interface{}{VmPropertiesToMap(v.Vm)})
		} else {
			s.D.Set("vm", nil)
		}

		s.D.Set("asset_source_ids", v.AssetSourceIds)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.ExternalAssetKey != nil {
			s.D.Set("external_asset_key", *v.ExternalAssetKey)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		if v.InventoryId != nil {
			s.D.Set("inventory_id", *v.InventoryId)
		}

		if v.SourceKey != nil {
			s.D.Set("source_key", *v.SourceKey)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_cloud_bridge.VmwareVmAsset:
		s.D.Set("asset_type", "VMWARE_VM")

		if v.Compute != nil {
			s.D.Set("compute", []interface{}{ComputePropertiesToMap(v.Compute)})
		} else {
			s.D.Set("compute", nil)
		}

		if v.Vm != nil {
			s.D.Set("vm", []interface{}{VmPropertiesToMap(v.Vm)})
		} else {
			s.D.Set("vm", nil)
		}

		if v.VmwareVCenter != nil {
			s.D.Set("vmware_vcenter", []interface{}{VmwareVCenterPropertiesToMap(v.VmwareVCenter)})
		} else {
			s.D.Set("vmware_vcenter", nil)
		}

		if v.VmwareVm != nil {
			s.D.Set("vmware_vm", []interface{}{VmwareVmPropertiesToMap(v.VmwareVm)})
		} else {
			s.D.Set("vmware_vm", nil)
		}

		s.D.Set("asset_source_ids", v.AssetSourceIds)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.ExternalAssetKey != nil {
			s.D.Set("external_asset_key", *v.ExternalAssetKey)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		if v.InventoryId != nil {
			s.D.Set("inventory_id", *v.InventoryId)
		}

		if v.SourceKey != nil {
			s.D.Set("source_key", *v.SourceKey)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'asset_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func AssetSummaryToMap(obj oci_cloud_bridge.AssetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["asset_source_ids"] = obj.AssetSourceIds

	result["asset_type"] = string(obj.AssetType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExternalAssetKey != nil {
		result["external_asset_key"] = string(*obj.ExternalAssetKey)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InventoryId != nil {
		result["inventory_id"] = string(*obj.InventoryId)
	}

	if obj.SourceKey != nil {
		result["source_key"] = string(*obj.SourceKey)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) mapToComputeProperties(fieldKeyFormat string) (oci_cloud_bridge.ComputeProperties, error) {
	result := oci_cloud_bridge.ComputeProperties{}

	if connectedNetworks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connected_networks")); ok {
		tmp := connectedNetworks.(int)
		result.ConnectedNetworks = &tmp
	}

	if coresCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cores_count")); ok {
		tmp := coresCount.(int)
		result.CoresCount = &tmp
	}

	if cpuModel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cpu_model")); ok {
		tmp := cpuModel.(string)
		result.CpuModel = &tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if disks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "disks")); ok {
		interfaces := disks.([]interface{})
		tmp := make([]oci_cloud_bridge.Disk, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "disks"), stateDataIndex)
			converted, err := s.mapToDisk(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "disks")) {
			result.Disks = tmp
		}
	}

	if disksCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "disks_count")); ok {
		tmp := disksCount.(int)
		result.DisksCount = &tmp
	}

	if dnsName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dns_name")); ok {
		tmp := dnsName.(string)
		result.DnsName = &tmp
	}

	if firmware, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "firmware")); ok {
		tmp := firmware.(string)
		result.Firmware = &tmp
	}

	if gpuDevices, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "gpu_devices")); ok {
		interfaces := gpuDevices.([]interface{})
		tmp := make([]oci_cloud_bridge.GpuDevice, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "gpu_devices"), stateDataIndex)
			converted, err := s.mapToGpuDevice(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "gpu_devices")) {
			result.GpuDevices = tmp
		}
	}

	if gpuDevicesCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "gpu_devices_count")); ok {
		tmp := gpuDevicesCount.(int)
		result.GpuDevicesCount = &tmp
	}

	if guestState, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "guest_state")); ok {
		tmp := guestState.(string)
		result.GuestState = &tmp
	}

	if hardwareVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hardware_version")); ok {
		tmp := hardwareVersion.(string)
		result.HardwareVersion = &tmp
	}

	if hostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name")); ok {
		tmp := hostName.(string)
		result.HostName = &tmp
	}

	if isPmemEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pmem_enabled")); ok {
		tmp := isPmemEnabled.(bool)
		result.IsPmemEnabled = &tmp
	}

	if isTpmEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_tpm_enabled")); ok {
		tmp := isTpmEnabled.(bool)
		result.IsTpmEnabled = &tmp
	}

	if latencySensitivity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "latency_sensitivity")); ok {
		tmp := latencySensitivity.(string)
		result.LatencySensitivity = &tmp
	}

	if memoryInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_mbs")); ok {
		tmp := memoryInMBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert memoryInMBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.MemoryInMBs = &tmpInt64
	}

	if nics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nics")); ok {
		interfaces := nics.([]interface{})
		tmp := make([]oci_cloud_bridge.Nic, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "nics"), stateDataIndex)
			converted, err := s.mapToNic(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nics")) {
			result.Nics = tmp
		}
	}

	if nicsCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nics_count")); ok {
		tmp := nicsCount.(int)
		result.NicsCount = &tmp
	}

	if nvdimmController, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nvdimm_controller")); ok {
		if tmpList := nvdimmController.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "nvdimm_controller"), 0)
			tmp, err := s.mapToNvdimmController(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert nvdimm_controller, encountered error: %v", err)
			}
			result.NvdimmController = &tmp
		}
	}

	if nvdimms, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nvdimms")); ok {
		interfaces := nvdimms.([]interface{})
		tmp := make([]oci_cloud_bridge.Nvdimm, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "nvdimms"), stateDataIndex)
			converted, err := s.mapToNvdimm(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nvdimms")) {
			result.Nvdimms = tmp
		}
	}

	if operatingSystem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operating_system")); ok {
		tmp := operatingSystem.(string)
		result.OperatingSystem = &tmp
	}

	if operatingSystemVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operating_system_version")); ok {
		tmp := operatingSystemVersion.(string)
		result.OperatingSystemVersion = &tmp
	}

	if pmemInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pmem_in_mbs")); ok {
		tmp := pmemInMBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert pmemInMBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.PmemInMBs = &tmpInt64
	}

	if powerState, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "power_state")); ok {
		tmp := powerState.(string)
		result.PowerState = &tmp
	}

	if primaryIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_ip")); ok {
		tmp := primaryIp.(string)
		result.PrimaryIp = &tmp
	}

	if scsiController, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scsi_controller")); ok {
		if tmpList := scsiController.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scsi_controller"), 0)
			tmp, err := s.mapToScsiController(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert scsi_controller, encountered error: %v", err)
			}
			result.ScsiController = &tmp
		}
	}

	if storageProvisionedInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_provisioned_in_mbs")); ok {
		tmp := storageProvisionedInMBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert storageProvisionedInMBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.StorageProvisionedInMBs = &tmpInt64
	}

	if threadsPerCoreCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "threads_per_core_count")); ok {
		tmp := threadsPerCoreCount.(int)
		result.ThreadsPerCoreCount = &tmp
	}

	return result, nil
}

func ComputePropertiesToMap(obj *oci_cloud_bridge.ComputeProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectedNetworks != nil {
		result["connected_networks"] = int(*obj.ConnectedNetworks)
	}

	if obj.CoresCount != nil {
		result["cores_count"] = int(*obj.CoresCount)
	}

	if obj.CpuModel != nil {
		result["cpu_model"] = string(*obj.CpuModel)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	disks := []interface{}{}
	for _, item := range obj.Disks {
		disks = append(disks, DiskToMap(item))
	}
	result["disks"] = disks

	if obj.DisksCount != nil {
		result["disks_count"] = int(*obj.DisksCount)
	}

	if obj.DnsName != nil {
		result["dns_name"] = string(*obj.DnsName)
	}

	if obj.Firmware != nil {
		result["firmware"] = string(*obj.Firmware)
	}

	gpuDevices := []interface{}{}
	for _, item := range obj.GpuDevices {
		gpuDevices = append(gpuDevices, GpuDeviceToMap(item))
	}
	result["gpu_devices"] = gpuDevices

	if obj.GpuDevicesCount != nil {
		result["gpu_devices_count"] = int(*obj.GpuDevicesCount)
	}

	if obj.GuestState != nil {
		result["guest_state"] = string(*obj.GuestState)
	}

	if obj.HardwareVersion != nil {
		result["hardware_version"] = string(*obj.HardwareVersion)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.IsPmemEnabled != nil {
		result["is_pmem_enabled"] = bool(*obj.IsPmemEnabled)
	}

	if obj.IsTpmEnabled != nil {
		result["is_tpm_enabled"] = bool(*obj.IsTpmEnabled)
	}

	if obj.LatencySensitivity != nil {
		result["latency_sensitivity"] = string(*obj.LatencySensitivity)
	}

	if obj.MemoryInMBs != nil {
		result["memory_in_mbs"] = strconv.FormatInt(*obj.MemoryInMBs, 10)
	}

	nics := []interface{}{}
	for _, item := range obj.Nics {
		nics = append(nics, NicToMap(item))
	}
	result["nics"] = nics

	if obj.NicsCount != nil {
		result["nics_count"] = int(*obj.NicsCount)
	}

	if obj.NvdimmController != nil {
		result["nvdimm_controller"] = []interface{}{NvdimmControllerToMap(obj.NvdimmController)}
	}

	nvdimms := []interface{}{}
	for _, item := range obj.Nvdimms {
		nvdimms = append(nvdimms, NvdimmToMap(item))
	}
	result["nvdimms"] = nvdimms

	if obj.OperatingSystem != nil {
		result["operating_system"] = string(*obj.OperatingSystem)
	}

	if obj.OperatingSystemVersion != nil {
		result["operating_system_version"] = string(*obj.OperatingSystemVersion)
	}

	if obj.PmemInMBs != nil {
		result["pmem_in_mbs"] = strconv.FormatInt(*obj.PmemInMBs, 10)
	}

	if obj.PowerState != nil {
		result["power_state"] = string(*obj.PowerState)
	}

	if obj.PrimaryIp != nil {
		result["primary_ip"] = string(*obj.PrimaryIp)
	}

	if obj.ScsiController != nil {
		result["scsi_controller"] = []interface{}{ScsiControllerToMap(obj.ScsiController)}
	}

	if obj.StorageProvisionedInMBs != nil {
		result["storage_provisioned_in_mbs"] = strconv.FormatInt(*obj.StorageProvisionedInMBs, 10)
	}

	if obj.ThreadsPerCoreCount != nil {
		result["threads_per_core_count"] = int(*obj.ThreadsPerCoreCount)
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) mapToCustomerTag(fieldKeyFormat string) (oci_cloud_bridge.CustomerTag, error) {
	result := oci_cloud_bridge.CustomerTag{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func CustomerTagToMap(obj oci_cloud_bridge.CustomerTag) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) mapToDisk(fieldKeyFormat string) (oci_cloud_bridge.Disk, error) {
	result := oci_cloud_bridge.Disk{}

	if bootOrder, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_order")); ok {
		tmp := bootOrder.(int)
		result.BootOrder = &tmp
	}

	if location, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location")); ok {
		tmp := location.(string)
		result.Location = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if persistentMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "persistent_mode")); ok {
		tmp := persistentMode.(string)
		result.PersistentMode = &tmp
	}

	if sizeInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "size_in_mbs")); ok {
		tmp := sizeInMBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert sizeInMBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.SizeInMBs = &tmpInt64
	}

	if uuid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "uuid")); ok {
		tmp := uuid.(string)
		result.Uuid = &tmp
	}

	if uuidLun, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "uuid_lun")); ok {
		tmp := uuidLun.(string)
		result.UuidLun = &tmp
	}

	return result, nil
}

func DiskToMap(obj oci_cloud_bridge.Disk) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BootOrder != nil {
		result["boot_order"] = int(*obj.BootOrder)
	}

	if obj.Location != nil {
		result["location"] = string(*obj.Location)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PersistentMode != nil {
		result["persistent_mode"] = string(*obj.PersistentMode)
	}

	if obj.SizeInMBs != nil {
		result["size_in_mbs"] = strconv.FormatInt(*obj.SizeInMBs, 10)
	}

	if obj.Uuid != nil {
		result["uuid"] = string(*obj.Uuid)
	}

	if obj.UuidLun != nil {
		result["uuid_lun"] = string(*obj.UuidLun)
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) mapToGpuDevice(fieldKeyFormat string) (oci_cloud_bridge.GpuDevice, error) {
	result := oci_cloud_bridge.GpuDevice{}

	if coresCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cores_count")); ok {
		tmp := coresCount.(int)
		result.CoresCount = &tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if manufacturer, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "manufacturer")); ok {
		tmp := manufacturer.(string)
		result.Manufacturer = &tmp
	}

	if memoryInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_mbs")); ok {
		tmp := memoryInMBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert memoryInMBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.MemoryInMBs = &tmpInt64
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func GpuDeviceToMap(obj oci_cloud_bridge.GpuDevice) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CoresCount != nil {
		result["cores_count"] = int(*obj.CoresCount)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Manufacturer != nil {
		result["manufacturer"] = string(*obj.Manufacturer)
	}

	if obj.MemoryInMBs != nil {
		result["memory_in_mbs"] = strconv.FormatInt(*obj.MemoryInMBs, 10)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) mapToNic(fieldKeyFormat string) (oci_cloud_bridge.Nic, error) {
	result := oci_cloud_bridge.Nic{}

	if ipAddresses, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ip_addresses")); ok {
		interfaces := ipAddresses.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ip_addresses")) {
			result.IpAddresses = tmp
		}
	}

	if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
		tmp := label.(string)
		result.Label = &tmp
	}

	if macAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mac_address")); ok {
		tmp := macAddress.(string)
		result.MacAddress = &tmp
	}

	if macAddressType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mac_address_type")); ok {
		tmp := macAddressType.(string)
		result.MacAddressType = &tmp
	}

	if networkName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_name")); ok {
		tmp := networkName.(string)
		result.NetworkName = &tmp
	}

	if switchName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "switch_name")); ok {
		tmp := switchName.(string)
		result.SwitchName = &tmp
	}

	return result, nil
}

func NicToMap(obj oci_cloud_bridge.Nic) map[string]interface{} {
	result := map[string]interface{}{}

	result["ip_addresses"] = obj.IpAddresses

	if obj.Label != nil {
		result["label"] = string(*obj.Label)
	}

	if obj.MacAddress != nil {
		result["mac_address"] = string(*obj.MacAddress)
	}

	if obj.MacAddressType != nil {
		result["mac_address_type"] = string(*obj.MacAddressType)
	}

	if obj.NetworkName != nil {
		result["network_name"] = string(*obj.NetworkName)
	}

	if obj.SwitchName != nil {
		result["switch_name"] = string(*obj.SwitchName)
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) mapToNvdimm(fieldKeyFormat string) (oci_cloud_bridge.Nvdimm, error) {
	result := oci_cloud_bridge.Nvdimm{}

	if controllerKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "controller_key")); ok {
		tmp := controllerKey.(int)
		result.ControllerKey = &tmp
	}

	if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
		tmp := label.(string)
		result.Label = &tmp
	}

	if unitNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit_number")); ok {
		tmp := unitNumber.(int)
		result.UnitNumber = &tmp
	}

	return result, nil
}

func NvdimmToMap(obj oci_cloud_bridge.Nvdimm) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ControllerKey != nil {
		result["controller_key"] = int(*obj.ControllerKey)
	}

	if obj.Label != nil {
		result["label"] = string(*obj.Label)
	}

	if obj.UnitNumber != nil {
		result["unit_number"] = int(*obj.UnitNumber)
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) mapToNvdimmController(fieldKeyFormat string) (oci_cloud_bridge.NvdimmController, error) {
	result := oci_cloud_bridge.NvdimmController{}

	if busNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bus_number")); ok {
		tmp := busNumber.(int)
		result.BusNumber = &tmp
	}

	if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
		tmp := label.(string)
		result.Label = &tmp
	}

	return result, nil
}

func NvdimmControllerToMap(obj *oci_cloud_bridge.NvdimmController) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BusNumber != nil {
		result["bus_number"] = int(*obj.BusNumber)
	}

	if obj.Label != nil {
		result["label"] = string(*obj.Label)
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) mapToScsiController(fieldKeyFormat string) (oci_cloud_bridge.ScsiController, error) {
	result := oci_cloud_bridge.ScsiController{}

	if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
		tmp := label.(string)
		result.Label = &tmp
	}

	if sharedBus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shared_bus")); ok {
		tmp := sharedBus.(string)
		result.SharedBus = &tmp
	}

	if unitNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit_number")); ok {
		tmp := unitNumber.(int)
		result.UnitNumber = &tmp
	}

	return result, nil
}

func ScsiControllerToMap(obj *oci_cloud_bridge.ScsiController) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Label != nil {
		result["label"] = string(*obj.Label)
	}

	if obj.SharedBus != nil {
		result["shared_bus"] = string(*obj.SharedBus)
	}

	if obj.UnitNumber != nil {
		result["unit_number"] = int(*obj.UnitNumber)
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) mapToVmProperties(fieldKeyFormat string) (oci_cloud_bridge.VmProperties, error) {
	result := oci_cloud_bridge.VmProperties{}

	if hypervisorHost, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hypervisor_host")); ok {
		tmp := hypervisorHost.(string)
		result.HypervisorHost = &tmp
	}

	if hypervisorVendor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hypervisor_vendor")); ok {
		tmp := hypervisorVendor.(string)
		result.HypervisorVendor = &tmp
	}

	if hypervisorVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hypervisor_version")); ok {
		tmp := hypervisorVersion.(string)
		result.HypervisorVersion = &tmp
	}

	return result, nil
}

func VmPropertiesToMap(obj *oci_cloud_bridge.VmProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HypervisorHost != nil {
		result["hypervisor_host"] = string(*obj.HypervisorHost)
	}

	if obj.HypervisorVendor != nil {
		result["hypervisor_vendor"] = string(*obj.HypervisorVendor)
	}

	if obj.HypervisorVersion != nil {
		result["hypervisor_version"] = string(*obj.HypervisorVersion)
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) mapToVmwareVCenterProperties(fieldKeyFormat string) (oci_cloud_bridge.VmwareVCenterProperties, error) {
	result := oci_cloud_bridge.VmwareVCenterProperties{}

	if dataCenter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_center")); ok {
		tmp := dataCenter.(string)
		result.DataCenter = &tmp
	}

	if vcenterKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcenter_key")); ok {
		tmp := vcenterKey.(string)
		result.VcenterKey = &tmp
	}

	if vcenterVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcenter_version")); ok {
		tmp := vcenterVersion.(string)
		result.VcenterVersion = &tmp
	}

	return result, nil
}

func VmwareVCenterPropertiesToMap(obj *oci_cloud_bridge.VmwareVCenterProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DataCenter != nil {
		result["data_center"] = string(*obj.DataCenter)
	}

	if obj.VcenterKey != nil {
		result["vcenter_key"] = string(*obj.VcenterKey)
	}

	if obj.VcenterVersion != nil {
		result["vcenter_version"] = string(*obj.VcenterVersion)
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) mapToVmwareVmProperties(fieldKeyFormat string) (oci_cloud_bridge.VmwareVmProperties, error) {
	result := oci_cloud_bridge.VmwareVmProperties{}

	if cluster, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cluster")); ok {
		tmp := cluster.(string)
		result.Cluster = &tmp
	}

	if customerFields, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_fields")); ok {
		interfaces := customerFields.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "customer_fields")) {
			result.CustomerFields = tmp
		}
	}

	if customerTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_tags")); ok {
		interfaces := customerTags.([]interface{})
		tmp := make([]oci_cloud_bridge.CustomerTag, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "customer_tags"), stateDataIndex)
			converted, err := s.mapToCustomerTag(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "customer_tags")) {
			result.CustomerTags = tmp
		}
	}

	if faultToleranceBandwidth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_tolerance_bandwidth")); ok {
		tmp := faultToleranceBandwidth.(int)
		result.FaultToleranceBandwidth = &tmp
	}

	if faultToleranceSecondaryLatency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_tolerance_secondary_latency")); ok {
		tmp := faultToleranceSecondaryLatency.(int)
		result.FaultToleranceSecondaryLatency = &tmp
	}

	if faultToleranceState, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_tolerance_state")); ok {
		tmp := faultToleranceState.(string)
		result.FaultToleranceState = &tmp
	}

	if instanceUuid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_uuid")); ok {
		tmp := instanceUuid.(string)
		result.InstanceUuid = &tmp
	}

	if isDisksCbtEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_disks_cbt_enabled")); ok {
		tmp := isDisksCbtEnabled.(bool)
		result.IsDisksCbtEnabled = &tmp
	}

	if isDisksUuidEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_disks_uuid_enabled")); ok {
		tmp := isDisksUuidEnabled.(bool)
		result.IsDisksUuidEnabled = &tmp
	}

	if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
		tmp := path.(string)
		result.Path = &tmp
	}

	if vmwareToolsStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vmware_tools_status")); ok {
		tmp := vmwareToolsStatus.(string)
		result.VmwareToolsStatus = &tmp
	}

	return result, nil
}

func VmwareVmPropertiesToMap(obj *oci_cloud_bridge.VmwareVmProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Cluster != nil {
		result["cluster"] = string(*obj.Cluster)
	}

	result["customer_fields"] = obj.CustomerFields

	customerTags := []interface{}{}
	for _, item := range obj.CustomerTags {
		customerTags = append(customerTags, CustomerTagToMap(item))
	}
	result["customer_tags"] = customerTags

	if obj.FaultToleranceBandwidth != nil {
		result["fault_tolerance_bandwidth"] = int(*obj.FaultToleranceBandwidth)
	}

	if obj.FaultToleranceSecondaryLatency != nil {
		result["fault_tolerance_secondary_latency"] = int(*obj.FaultToleranceSecondaryLatency)
	}

	if obj.FaultToleranceState != nil {
		result["fault_tolerance_state"] = string(*obj.FaultToleranceState)
	}

	if obj.InstanceUuid != nil {
		result["instance_uuid"] = string(*obj.InstanceUuid)
	}

	if obj.IsDisksCbtEnabled != nil {
		result["is_disks_cbt_enabled"] = bool(*obj.IsDisksCbtEnabled)
	}

	if obj.IsDisksUuidEnabled != nil {
		result["is_disks_uuid_enabled"] = bool(*obj.IsDisksUuidEnabled)
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	if obj.VmwareToolsStatus != nil {
		result["vmware_tools_status"] = string(*obj.VmwareToolsStatus)
	}

	return result
}

func (s *CloudBridgeAssetResourceCrud) populateTopLevelPolymorphicCreateAssetRequest(request *oci_cloud_bridge.CreateAssetRequest) error {
	//discriminator
	assetTypeRaw, ok := s.D.GetOkExists("asset_type")
	var assetType string
	if ok {
		assetType = assetTypeRaw.(string)
	} else {
		assetType = "" // default value
	}
	switch strings.ToLower(assetType) {
	case strings.ToLower("VMWARE_VM"):
		details := oci_cloud_bridge.CreateVmwareVmAssetDetails{}
		if compute, ok := s.D.GetOkExists("compute"); ok {
			if tmpList := compute.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "compute", 0)
				tmp, err := s.mapToComputeProperties(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Compute = &tmp
			}
		}
		if vm, ok := s.D.GetOkExists("vm"); ok {
			if tmpList := vm.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vm", 0)
				tmp, err := s.mapToVmProperties(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Vm = &tmp
			}
		}
		if vmwareVCenter, ok := s.D.GetOkExists("vmware_vcenter"); ok {
			if tmpList := vmwareVCenter.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vmware_vcenter", 0)
				tmp, err := s.mapToVmwareVCenterProperties(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.VmwareVCenter = &tmp
			}
		}
		if vmwareVm, ok := s.D.GetOkExists("vmware_vm"); ok {
			if tmpList := vmwareVm.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vmware_vm", 0)
				tmp, err := s.mapToVmwareVmProperties(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.VmwareVm = &tmp
			}
		}
		if assetSourceIds, ok := s.D.GetOkExists("asset_source_ids"); ok {
			interfaces := assetSourceIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("asset_source_ids") {
				details.AssetSourceIds = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if externalAssetKey, ok := s.D.GetOkExists("external_asset_key"); ok {
			tmp := externalAssetKey.(string)
			details.ExternalAssetKey = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if inventoryId, ok := s.D.GetOkExists("inventory_id"); ok {
			tmp := inventoryId.(string)
			details.InventoryId = &tmp
		}
		if sourceKey, ok := s.D.GetOkExists("source_key"); ok {
			tmp := sourceKey.(string)
			details.SourceKey = &tmp
		}
		request.CreateAssetDetails = details
	default:
		return fmt.Errorf("unknown asset_type '%v' was specified", assetType)
	}
	return nil
}

func (s *CloudBridgeAssetResourceCrud) populateTopLevelPolymorphicUpdateAssetRequest(request *oci_cloud_bridge.UpdateAssetRequest) error {
	//discriminator
	assetTypeRaw, ok := s.D.GetOkExists("asset_type")
	var assetType string
	if ok {
		assetType = assetTypeRaw.(string)
	} else {
		assetType = "" // default value
	}
	switch strings.ToLower(assetType) {
	case strings.ToLower("VM"):
		details := oci_cloud_bridge.UpdateVmAssetDetails{}
		tmp := s.D.Id()
		request.AssetId = &tmp
		if assetSourceIds, ok := s.D.GetOkExists("asset_source_ids"); ok {
			interfaces := assetSourceIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("asset_source_ids") {
				details.AssetSourceIds = tmp
			}
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		request.UpdateAssetDetails = details
	case strings.ToLower("VMWARE_VM"):
		details := oci_cloud_bridge.UpdateVmwareVmAssetDetails{}
		if compute, ok := s.D.GetOkExists("compute"); ok {
			if tmpList := compute.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "compute", 0)
				tmp, err := s.mapToComputeProperties(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Compute = &tmp
			}
		}
		if vm, ok := s.D.GetOkExists("vm"); ok {
			if tmpList := vm.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vm", 0)
				tmp, err := s.mapToVmProperties(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Vm = &tmp
			}
		}
		if vmwareVCenter, ok := s.D.GetOkExists("vmware_vcenter"); ok {
			if tmpList := vmwareVCenter.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vmware_vcenter", 0)
				tmp, err := s.mapToVmwareVCenterProperties(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.VmwareVCenter = &tmp
			}
		}
		if vmwareVm, ok := s.D.GetOkExists("vmware_vm"); ok {
			if tmpList := vmwareVm.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vmware_vm", 0)
				tmp, err := s.mapToVmwareVmProperties(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.VmwareVm = &tmp
			}
		}
		tmp := s.D.Id()
		request.AssetId = &tmp
		if assetSourceIds, ok := s.D.GetOkExists("asset_source_ids"); ok {
			interfaces := assetSourceIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("asset_source_ids") {
				details.AssetSourceIds = tmp
			}
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		request.UpdateAssetDetails = details
	default:
		return fmt.Errorf("unknown asset_type '%v' was specified", assetType)
	}
	return nil
}

func (s *CloudBridgeAssetResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_bridge.ChangeAssetCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AssetId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	_, err := s.Client.ChangeAssetCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
