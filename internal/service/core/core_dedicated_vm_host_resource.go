// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func CoreDedicatedVmHostResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreDedicatedVmHost,
		Read:     readCoreDedicatedVmHost,
		Update:   updateCoreDedicatedVmHost,
		Delete:   deleteCoreDedicatedVmHost,
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
			"dedicated_vm_host_shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"capacity_config": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"fault_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_memory_encryption_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"placement_constraint_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"COMPUTE_BARE_METAL_HOST",
							}, true),
						},
						// Optional
						"compute_bare_metal_host_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						// Computed
					},
				},
			},

			// Computed
			"capacity_bins": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"capacity_index": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"remaining_memory_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"remaining_ocpus": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"supported_shapes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"total_memory_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"total_ocpus": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
			"compute_bare_metal_host_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remaining_memory_in_gbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"remaining_ocpus": {
				Type:     schema.TypeFloat,
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
			"total_memory_in_gbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"total_ocpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func createCoreDedicatedVmHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readCoreDedicatedVmHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreDedicatedVmHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreDedicatedVmHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type CoreDedicatedVmHostResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	workRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_core.DedicatedVmHost
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *CoreDedicatedVmHostResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreDedicatedVmHostResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DedicatedVmHostLifecycleStateCreating),
	}
}

func (s *CoreDedicatedVmHostResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DedicatedVmHostLifecycleStateActive),
	}
}

func (s *CoreDedicatedVmHostResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DedicatedVmHostLifecycleStateDeleting),
	}
}

func (s *CoreDedicatedVmHostResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DedicatedVmHostLifecycleStateDeleted),
	}
}

func (s *CoreDedicatedVmHostResourceCrud) Create() error {
	request := oci_core.CreateDedicatedVmHostRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if capacityConfig, ok := s.D.GetOkExists("capacity_config"); ok {
		tmp := capacityConfig.(string)
		request.CapacityConfig = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dedicatedVmHostShape, ok := s.D.GetOkExists("dedicated_vm_host_shape"); ok {
		tmp := dedicatedVmHostShape.(string)
		request.DedicatedVmHostShape = &tmp
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

	if faultDomain, ok := s.D.GetOkExists("fault_domain"); ok {
		tmp := faultDomain.(string)
		request.FaultDomain = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isMemoryEncryptionEnabled, ok := s.D.GetOkExists("is_memory_encryption_enabled"); ok {
		tmp := isMemoryEncryptionEnabled.(bool)
		request.IsMemoryEncryptionEnabled = &tmp
	}

	if placementConstraintDetails, ok := s.D.GetOkExists("placement_constraint_details"); ok {
		if tmpList := placementConstraintDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "placement_constraint_details", 0)
			tmp, err := s.mapToPlacementConstraintDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PlacementConstraintDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDedicatedVmHost(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.DedicatedVmHost

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "dedicatedvmhost", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *CoreDedicatedVmHostResourceCrud) Get() error {
	request := oci_core.GetDedicatedVmHostRequest{}

	tmp := s.D.Id()
	request.DedicatedVmHostId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDedicatedVmHost(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DedicatedVmHost
	return nil
}

func (s *CoreDedicatedVmHostResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateDedicatedVmHostRequest{}

	tmp := s.D.Id()
	request.DedicatedVmHostId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDedicatedVmHost(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DedicatedVmHost
	return nil
}

func (s *CoreDedicatedVmHostResourceCrud) Delete() error {
	request := oci_core.DeleteDedicatedVmHostRequest{}

	tmp := s.D.Id()
	request.DedicatedVmHostId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.DeleteDedicatedVmHost(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "dedicatedvmhost", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *CoreDedicatedVmHostResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	capacityBins := []interface{}{}
	for _, item := range s.Res.CapacityBins {
		capacityBins = append(capacityBins, CapacityBinToMap(item))
	}
	s.D.Set("capacity_bins", capacityBins)

	if s.Res.CapacityConfig != nil {
		s.D.Set("capacity_config", *s.Res.CapacityConfig)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeBareMetalHostId != nil {
		s.D.Set("compute_bare_metal_host_id", *s.Res.ComputeBareMetalHostId)
	}

	if s.Res.DedicatedVmHostShape != nil {
		s.D.Set("dedicated_vm_host_shape", *s.Res.DedicatedVmHostShape)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsMemoryEncryptionEnabled != nil {
		s.D.Set("is_memory_encryption_enabled", *s.Res.IsMemoryEncryptionEnabled)
	}

	if s.Res.PlacementConstraintDetails != nil {
		placementConstraintDetailsArray := []interface{}{}
		if placementConstraintDetailsMap := PlacementConstraintDetailsToMap(&s.Res.PlacementConstraintDetails); placementConstraintDetailsMap != nil {
			placementConstraintDetailsArray = append(placementConstraintDetailsArray, placementConstraintDetailsMap)
		}
		s.D.Set("placement_constraint_details", placementConstraintDetailsArray)
	} else {
		s.D.Set("placement_constraint_details", nil)
	}

	if s.Res.RemainingMemoryInGBs != nil {
		s.D.Set("remaining_memory_in_gbs", *s.Res.RemainingMemoryInGBs)
	}

	if s.Res.RemainingOcpus != nil {
		s.D.Set("remaining_ocpus", *s.Res.RemainingOcpus)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TotalMemoryInGBs != nil {
		s.D.Set("total_memory_in_gbs", *s.Res.TotalMemoryInGBs)
	}

	if s.Res.TotalOcpus != nil {
		s.D.Set("total_ocpus", *s.Res.TotalOcpus)
	}

	return nil
}

func (s *CoreDedicatedVmHostResourceCrud) mapToPlacementConstraintDetails(fieldKeyFormat string) (oci_core.PlacementConstraintDetails, error) {
	var baseObject oci_core.PlacementConstraintDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("COMPUTE_BARE_METAL_HOST"):
		details := oci_core.ComputeBareMetalHostPlacementConstraintDetails{}
		if computeBareMetalHostId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_bare_metal_host_id")); ok {
			tmp := computeBareMetalHostId.(string)
			details.ComputeBareMetalHostId = &tmp
		}
		baseObject = details
	case strings.ToLower("HOST_GROUP"):
		details := oci_core.HostGroupPlacementConstraintDetails{}
		if computeHostGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_host_group_id")); ok {
			tmp := computeHostGroupId.(string)
			details.ComputeHostGroupId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func PlacementConstraintDetailsToMap(obj *oci_core.PlacementConstraintDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.ComputeBareMetalHostPlacementConstraintDetails:
		result["type"] = "COMPUTE_BARE_METAL_HOST"

		if v.ComputeBareMetalHostId != nil {
			result["compute_bare_metal_host_id"] = string(*v.ComputeBareMetalHostId)
		}
	case oci_core.HostGroupPlacementConstraintDetails:
		result["type"] = "HOST_GROUP"

		if v.ComputeHostGroupId != nil {
			result["compute_host_group_id"] = string(*v.ComputeHostGroupId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func CapacityBinToMap(obj oci_core.CapacityBin) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CapacityIndex != nil {
		result["capacity_index"] = int(*obj.CapacityIndex)
	}

	if obj.RemainingMemoryInGBs != nil {
		result["remaining_memory_in_gbs"] = float32(*obj.RemainingMemoryInGBs)
	}

	if obj.RemainingOcpus != nil {
		result["remaining_ocpus"] = float32(*obj.RemainingOcpus)
	}

	result["supported_shapes"] = obj.SupportedShapes

	if obj.TotalMemoryInGBs != nil {
		result["total_memory_in_gbs"] = float32(*obj.TotalMemoryInGBs)
	}

	if obj.TotalOcpus != nil {
		result["total_ocpus"] = float32(*obj.TotalOcpus)
	}

	return result
}

func (s *CoreDedicatedVmHostResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeDedicatedVmHostCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DedicatedVmHostId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.ChangeDedicatedVmHostCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "dedicatedvmhost", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
