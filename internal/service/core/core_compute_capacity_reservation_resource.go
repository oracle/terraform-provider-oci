// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func CoreComputeCapacityReservationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreComputeCapacityReservation,
		Read:     readCoreComputeCapacityReservation,
		Update:   updateCoreComputeCapacityReservation,
		Delete:   deleteCoreComputeCapacityReservation,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
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
			"instance_reservation_configs": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      instanceReservationConfigsHashCodeForSets,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"instance_shape": {
							Type:     schema.TypeString,
							Required: true,
						},
						"reserved_count": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						// Optional
						"cluster_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"hpc_island_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"network_block_ids": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},
						"cluster_placement_group_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"fault_domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"ocpus": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
						"used_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_default_reservation": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"reserved_instance_count": {
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
			"used_instance_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreComputeCapacityReservation(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityReservationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readCoreComputeCapacityReservation(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityReservationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreComputeCapacityReservation(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityReservationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreComputeCapacityReservation(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityReservationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type CoreComputeCapacityReservationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	workRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_core.ComputeCapacityReservation
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *CoreComputeCapacityReservationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreComputeCapacityReservationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ComputeCapacityReservationLifecycleStateCreating),
	}
}

func (s *CoreComputeCapacityReservationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ComputeCapacityReservationLifecycleStateActive),
	}
}

func (s *CoreComputeCapacityReservationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.ComputeCapacityReservationLifecycleStateDeleting),
	}
}

func (s *CoreComputeCapacityReservationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.ComputeCapacityReservationLifecycleStateDeleted),
	}
}

func (s *CoreComputeCapacityReservationResourceCrud) Create() error {
	request := oci_core.CreateComputeCapacityReservationRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceReservationConfigs, ok := s.D.GetOkExists("instance_reservation_configs"); ok {

		set := instanceReservationConfigs.(*schema.Set)
		interfaces := set.List()

		tmp := make([]oci_core.InstanceReservationConfigDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := instanceReservationConfigsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "instance_reservation_configs", stateDataIndex)
			converted, err := s.mapToInstanceReservationConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("instance_reservation_configs") {
			request.InstanceReservationConfigs = tmp
		}
	}

	if isDefaultReservation, ok := s.D.GetOkExists("is_default_reservation"); ok {
		tmp := isDefaultReservation.(bool)
		request.IsDefaultReservation = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateComputeCapacityReservation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId == nil {
		return fmt.Errorf("CreateComputeCapacityReservation response.OpcWorkRequestId was nil")
	}
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "capacityreservation", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *CoreComputeCapacityReservationResourceCrud) Get() error {
	request := oci_core.GetComputeCapacityReservationRequest{}

	tmp := s.D.Id()
	request.CapacityReservationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetComputeCapacityReservation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeCapacityReservation
	return nil
}

func (s *CoreComputeCapacityReservationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateComputeCapacityReservationRequest{}

	tmp := s.D.Id()
	request.CapacityReservationId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceReservationConfigs, ok := s.D.GetOkExists("instance_reservation_configs"); ok {
		set := instanceReservationConfigs.(*schema.Set)
		interfaces := set.List()

		tmp := make([]oci_core.InstanceReservationConfigDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := instanceReservationConfigsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "instance_reservation_configs", stateDataIndex)
			converted, err := s.mapToInstanceReservationConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("instance_reservation_configs") {
			request.InstanceReservationConfigs = tmp
		}
	}

	if isDefaultReservation, ok := s.D.GetOkExists("is_default_reservation"); ok {
		tmp := isDefaultReservation.(bool)
		request.IsDefaultReservation = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateComputeCapacityReservation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "capacityreservation", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *CoreComputeCapacityReservationResourceCrud) Delete() error {
	request := oci_core.DeleteComputeCapacityReservationRequest{}

	tmp := s.D.Id()
	request.CapacityReservationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.DeleteComputeCapacityReservation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "capacityreservation", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *CoreComputeCapacityReservationResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	instanceReservationConfigs := []interface{}{}
	for _, item := range s.Res.InstanceReservationConfigs {
		instanceReservationConfigs = append(instanceReservationConfigs, InstanceReservationConfigToMap(item))
	}
	s.D.Set("instance_reservation_configs", instanceReservationConfigs)

	if s.Res.IsDefaultReservation != nil {
		s.D.Set("is_default_reservation", *s.Res.IsDefaultReservation)
	}

	if s.Res.ReservedInstanceCount != nil {
		s.D.Set("reserved_instance_count", strconv.FormatInt(*s.Res.ReservedInstanceCount, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UsedInstanceCount != nil {
		s.D.Set("used_instance_count", strconv.FormatInt(*s.Res.UsedInstanceCount, 10))
	}

	return nil
}

func (s *CoreComputeCapacityReservationResourceCrud) mapToClusterConfigDetails(fieldKeyFormat string) (oci_core.ClusterConfigDetails, error) {
	result := oci_core.ClusterConfigDetails{}

	if hpcIslandId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hpc_island_id")); ok {
		tmp := hpcIslandId.(string)
		result.HpcIslandId = &tmp
	}

	if networkBlockIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_block_ids")); ok {
		interfaces := networkBlockIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "network_block_ids")) {
			result.NetworkBlockIds = tmp
		}
	}

	return result, nil
}

func ClusterConfigDetailsToMap(obj *oci_core.ClusterConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HpcIslandId != nil {
		result["hpc_island_id"] = string(*obj.HpcIslandId)
	}

	result["network_block_ids"] = obj.NetworkBlockIds
	result["network_block_ids"] = obj.NetworkBlockIds

	return result
}

func (s *CoreComputeCapacityReservationResourceCrud) mapToInstanceReservationConfigDetails(fieldKeyFormat string) (oci_core.InstanceReservationConfigDetails, error) {
	result := oci_core.InstanceReservationConfigDetails{}

	if clusterConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cluster_config")); ok {
		if tmpList := clusterConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "cluster_config"), 0)
			tmp, err := s.mapToClusterConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert cluster_config, encountered error: %v", err)
			}
			result.ClusterConfig = &tmp
		}
	}

	if clusterPlacementGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cluster_placement_group_id")); ok {
		tmp := clusterPlacementGroupId.(string)
		if tmp != "" {
			result.ClusterPlacementGroupId = &tmp
		}
	}

	if faultDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_domain")); ok {
		tmp := faultDomain.(string)
		if tmp != "" {
			result.FaultDomain = &tmp
		}
	}

	if instanceShape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_shape")); ok {
		tmp := instanceShape.(string)
		result.InstanceShape = &tmp
	}

	if instanceShapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_shape_config")); ok {
		if tmpList := instanceShapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "instance_shape_config"), 0)
			tmp, err := s.mapToInstanceReservationShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert instance_shape_config, encountered error: %v", err)
			}
			result.InstanceShapeConfig = &tmp
		}
	}

	if reservedCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reserved_count")); ok {
		tmp := reservedCount.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert reservedCount string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.ReservedCount = &tmpInt64
	}

	return result, nil
}

func InstanceReservationConfigToMap(obj oci_core.InstanceReservationConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClusterConfig != nil {
		result["cluster_config"] = []interface{}{ClusterConfigDetailsToMap(obj.ClusterConfig)}
	}

	if obj.ClusterPlacementGroupId != nil {
		result["cluster_placement_group_id"] = string(*obj.ClusterPlacementGroupId)
	}

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	if obj.InstanceShape != nil {
		result["instance_shape"] = string(*obj.InstanceShape)
	}

	if obj.InstanceShapeConfig != nil {
		result["instance_shape_config"] = []interface{}{InstanceReservationShapeConfigDetailsToMap(obj.InstanceShapeConfig)}
	}

	if obj.ReservedCount != nil {
		result["reserved_count"] = strconv.FormatInt(*obj.ReservedCount, 10)
	}

	if obj.UsedCount != nil {
		result["used_count"] = strconv.FormatInt(*obj.UsedCount, 10)
	}

	return result
}

func (s *CoreComputeCapacityReservationResourceCrud) mapToInstanceReservationShapeConfigDetails(fieldKeyFormat string) (oci_core.InstanceReservationShapeConfigDetails, error) {
	result := oci_core.InstanceReservationShapeConfigDetails{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	return result, nil
}

func InstanceReservationShapeConfigDetailsToMap(obj *oci_core.InstanceReservationShapeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	return result
}

func (s *CoreComputeCapacityReservationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeComputeCapacityReservationCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CapacityReservationId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.ChangeComputeCapacityReservationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "capacityreservation", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}

func instanceReservationConfigsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	if instanceShape, ok := m["instance_shape"]; ok && instanceShape != "" {
		buf.WriteString(fmt.Sprintf("%v-", instanceShape))
	}

	if reservedCount, ok := m["reserved_count"]; ok && reservedCount != "" {
		buf.WriteString(fmt.Sprintf("%v-", reservedCount))
	}

	if clusterConfig, ok := m["cluster_config"]; ok && clusterConfig != "" {
		if clusterConfigList, ok := clusterConfig.([]interface{}); ok && len(clusterConfigList) > 0 {
			buf.WriteString("cluster_config-")
			for _, clusterConfigMapInterface := range clusterConfigList {
				clusterConfigMap := clusterConfigMapInterface.(map[string]interface{})
				if hpcIslandIdVal, ok := clusterConfigMap["hpc_island_id"].(string); ok && hpcIslandIdVal != "" {
					buf.WriteString(fmt.Sprintf("%v-", hpcIslandIdVal))
				}
				if networkBlockIdsVal, ok := clusterConfigMap["network_block_ids"]; ok && networkBlockIdsVal != "" {
					if networkBlockIdsList, ok := networkBlockIdsVal.([]interface{}); ok && len(networkBlockIdsList) > 0 {
						buf.WriteString("network_block_ids-")
						for _, networkBlockIdInterface := range networkBlockIdsList {
							networkBlockId := networkBlockIdInterface.(string)
							buf.WriteString(fmt.Sprintf("%s-", networkBlockId))
						}
					}
				}
			}
		}
	}

	if clusterPlacementGroupId, ok := m["cluster_placement_group_id"]; ok && clusterPlacementGroupId != "" {
		buf.WriteString(fmt.Sprintf("%v-", clusterPlacementGroupId))
	}

	if faultDomain, ok := m["fault_domain"]; ok && faultDomain != "" {
		buf.WriteString(fmt.Sprintf("%v-", faultDomain))
	}

	if usedCountInterface, ok := m["used_count"]; ok && usedCountInterface != "" {
		if usedCountInt, ok := usedCountInterface.(int); ok && usedCountInt > 0 {
			buf.WriteString(fmt.Sprintf("%d-", usedCountInt))
		}
	}

	return utils.GetStringHashcode(buf.String())
}
