// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreComputeHostResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreComputeHost,
		Read:     readCoreComputeHost,
		Update:   updateCoreComputeHost,
		Delete:   deleteCoreComputeHost,
		Schema: map[string]*schema.Schema{
			// Required
			"compute_host_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compute_host_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"additional_data": {
				Type:     schema.TypeString,
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
			"configuration_data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"check_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"configuration_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"firmware_bundle_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recycle_level": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_last_apply": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configuration_state": {
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
			"firmware_bundle_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"gpu_memory_fabric_id": {
				Type:     schema.TypeString,
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
			"impacted_component_details": {
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
			"network_block_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recycle_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compute_host_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"recycle_level": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configuration_action_type": {
				Type:     schema.TypeString,
				Computed: false,
				Optional: true,
			},
			"time_configuration_check": {
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
	}
}

func createCoreComputeHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readCoreComputeHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreComputeHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreComputeHost(d *schema.ResourceData, m interface{}) error {
	return nil
}

type CoreComputeHostResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.ComputeHost
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *CoreComputeHostResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreComputeHostResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ComputeHostLifecycleStateProvisioning),
	}
}

func (s *CoreComputeHostResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ComputeHostLifecycleStateAvailable),
	}
}

func (s *CoreComputeHostResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.ComputeHostLifecycleStateUnavailable),
	}
}

func (s *CoreComputeHostResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *CoreComputeHostResourceCrud) Create() error {
	computeHostId := s.D.Get("compute_host_id").(string)
	computeHost, err := s.getComputeHost(computeHostId)
	if err != nil {
		return err
	}

	s.Res = computeHost

	return s.attachDetachHostGroupIfNecessary(computeHost)
}

func (s *CoreComputeHostResourceCrud) Get() error {
	computeHost, err := s.getComputeHost(s.D.Id())
	if err != nil {
		return err
	}

	s.Res = computeHost
	return nil
}

func (s *CoreComputeHostResourceCrud) Update() error {
	if _, ok := s.D.GetOk("configuration_action_type"); ok && s.D.HasChange("configuration_action_type") {
		if s.D.Get("configuration_action_type").(string) == "check" {
			err := s.CheckHostConfiguration()
			if err != nil {
				return err
			}
		}
		if s.D.Get("configuration_action_type").(string) == "apply" {
			err := s.ApplyHostConfiguration()
			if err != nil {
				return err
			}
		}
	}

	computeHost, err := s.getComputeHost(s.D.Id())
	if err != nil {
		return err
	}

	return s.attachDetachHostGroupIfNecessary(computeHost)
}

func (s *CoreComputeHostResourceCrud) SetData() error {
	s.D.Set("additional_data", readData(s.Res.AdditionalData))

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CapacityReservationId != nil {
		s.D.Set("capacity_reservation_id", *s.Res.CapacityReservationId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeHostGroupId != nil {
		s.D.Set("compute_host_group_id", *s.Res.ComputeHostGroupId)
	}

	if s.Res.ConfigurationData != nil {
		s.D.Set("configuration_data", []interface{}{ComputeHostConfigurationDataToMap(s.Res.ConfigurationData)})
	} else {
		s.D.Set("configuration_data", nil)
	}

	s.D.Set("configuration_state", s.Res.ConfigurationState)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	if s.Res.FirmwareBundleId != nil {
		s.D.Set("firmware_bundle_id", *s.Res.FirmwareBundleId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GpuMemoryFabricId != nil {
		s.D.Set("gpu_memory_fabric_id", *s.Res.GpuMemoryFabricId)
	}

	s.D.Set("health", s.Res.Health)

	if s.Res.HpcIslandId != nil {
		s.D.Set("hpc_island_id", *s.Res.HpcIslandId)
	}

	if s.Res.Id != nil {
		s.D.Set("compute_host_id", *s.Res.Id)
	}

	s.D.Set("impacted_component_details", readData(s.Res.ImpactedComponentDetails))

	if s.Res.InstanceId != nil {
		s.D.Set("instance_id", *s.Res.InstanceId)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.LocalBlockId != nil {
		s.D.Set("local_block_id", *s.Res.LocalBlockId)
	}

	if s.Res.NetworkBlockId != nil {
		s.D.Set("network_block_id", *s.Res.NetworkBlockId)
	}

	if s.Res.Platform != nil {
		s.D.Set("platform", *s.Res.Platform)
	}

	if s.Res.RecycleDetails != nil {
		s.D.Set("recycle_details", []interface{}{RecycleDetailsToMap(s.Res.RecycleDetails)})
	} else {
		s.D.Set("recycle_details", nil)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeConfigurationCheck != nil {
		s.D.Set("time_configuration_check", s.Res.TimeConfigurationCheck.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *CoreComputeHostResourceCrud) ApplyHostConfiguration() error {
	request := oci_core.ApplyHostConfigurationRequest{}

	idTmp := s.D.Id()
	request.ComputeHostId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ApplyHostConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *CoreComputeHostResourceCrud) CheckHostConfiguration() error {
	request := oci_core.CheckHostConfigurationRequest{}

	idTmp := s.D.Id()
	request.ComputeHostId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.CheckHostConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func ComputeHostConfigurationCheckDetailsToMap(obj *oci_core.ComputeHostConfigurationCheckDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["configuration_state"] = string(obj.ConfigurationState)

	if obj.FirmwareBundleId != nil {
		result["firmware_bundle_id"] = string(*obj.FirmwareBundleId)
	}

	result["recycle_level"] = string(obj.RecycleLevel)

	result["type"] = string(obj.Type)

	return result
}

func ComputeHostConfigurationDataToMap(obj *oci_core.ComputeHostConfigurationData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CheckDetails != nil {
		result["check_details"] = []interface{}{ComputeHostConfigurationCheckDetailsToMap(obj.CheckDetails)}
	}

	if obj.TimeLastApply != nil {
		result["time_last_apply"] = obj.TimeLastApply.String()
	}

	return result
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

	if obj.ComputeHostGroupId != nil {
		result["compute_host_group_id"] = string(*obj.ComputeHostGroupId)
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

	if obj.GpuMemoryFabricId != nil {
		result["gpu_memory_fabric_id"] = string(*obj.GpuMemoryFabricId)
	}

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

	if obj.NetworkBlockId != nil {
		result["network_block_id"] = string(*obj.NetworkBlockId)
	}

	if obj.Platform != nil {
		result["platform"] = string(*obj.Platform)
	}

	if obj.RecycleDetails != nil {
		result["recycle_details"] = []interface{}{RecycleDetailsToMap(obj.RecycleDetails)}
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

func (s *CoreComputeHostResourceCrud) waitForUpdateWorkRequest(opcWorkRequestId string) error {
	_, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, &opcWorkRequestId, "computebaremetalhost", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	return err
}

func (s *CoreComputeHostResourceCrud) getComputeHost(computeHostId string) (*oci_core.ComputeHost, error) {
	request := oci_core.GetComputeHostRequest{}

	request.ComputeHostId = &computeHostId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetComputeHost(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return &response.ComputeHost, nil
}

func (s *CoreComputeHostResourceCrud) attachDetachHostGroupIfNecessary(computeHost *oci_core.ComputeHost) error {
	currentComputeHostGroupId := ""
	if computeHost.ComputeHostGroupId != nil {
		currentComputeHostGroupId = *computeHost.ComputeHostGroupId
	}

	newComputeHostGroupId := ""
	if computeHostGroupId, ok := s.D.GetOk("compute_host_group_id"); ok {
		newComputeHostGroupId = computeHostGroupId.(string)
	}

	if newComputeHostGroupId != currentComputeHostGroupId {
		if currentComputeHostGroupId != "" {
			request := oci_core.DetachComputeHostGroupHostRequest{}
			request.ComputeHostId = computeHost.Id
			request.ComputeHostGroupId = &currentComputeHostGroupId

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
			response, err := s.Client.DetachComputeHostGroupHost(context.Background(), request)
			if err != nil {
				return err
			}

			if response.OpcWorkRequestId != nil {
				if err := s.waitForUpdateWorkRequest(*response.OpcWorkRequestId); err != nil {
					return err
				}
			}
		}
		if newComputeHostGroupId != "" {
			request := oci_core.AttachComputeHostGroupHostRequest{}
			request.ComputeHostId = computeHost.Id
			request.ComputeHostGroupId = &newComputeHostGroupId
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

			response, err := s.Client.AttachComputeHostGroupHost(context.Background(), request)
			if err != nil {
				return err
			}

			if response.OpcWorkRequestId != nil {
				if err := s.waitForUpdateWorkRequest(*response.OpcWorkRequestId); err != nil {
					return err
				}
			}
		}
	}

	computeHost, err := s.getComputeHost(*computeHost.Id)
	if err != nil {
		return err
	}

	s.Res = computeHost

	return nil
}
