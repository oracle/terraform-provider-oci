// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreComputeHostGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreComputeHostGroup,
		Read:     readCoreComputeHostGroup,
		Update:   updateCoreComputeHostGroup,
		Delete:   deleteCoreComputeHostGroup,
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
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_targeted_placement_required": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"configurations": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"firmware_bundle_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"recycle_level": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"target": {
							Type:     schema.TypeString,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
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

func createCoreComputeHostGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeHostGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readCoreComputeHostGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeHostGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreComputeHostGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeHostGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreComputeHostGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeHostGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type CoreComputeHostGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.ComputeHostGroup
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *CoreComputeHostGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreComputeHostGroupResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *CoreComputeHostGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ComputeHostGroupLifecycleStateActive),
	}
}

func (s *CoreComputeHostGroupResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *CoreComputeHostGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.ComputeHostGroupLifecycleStateDeleted),
	}
}

func (s *CoreComputeHostGroupResourceCrud) Create() error {
	request := oci_core.CreateComputeHostGroupRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurations, ok := s.D.GetOkExists("configurations"); ok {
		interfaces := configurations.([]interface{})
		tmp := make([]oci_core.HostGroupConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configurations", stateDataIndex)
			converted, err := s.mapToHostGroupConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configurations") {
			request.Configurations = tmp
		}
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

	if isTargetedPlacementRequired, ok := s.D.GetOkExists("is_targeted_placement_required"); ok {
		tmp := isTargetedPlacementRequired.(bool)
		request.IsTargetedPlacementRequired = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateComputeHostGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeHostGroup
	s.D.SetId(*response.Id)

	return s.Get()
}

func (s *CoreComputeHostGroupResourceCrud) Get() error {
	request := oci_core.GetComputeHostGroupRequest{}

	tmp := s.D.Id()
	request.ComputeHostGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetComputeHostGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeHostGroup
	return nil
}

func (s *CoreComputeHostGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateComputeHostGroupRequest{}

	tmp := s.D.Id()
	request.ComputeHostGroupId = &tmp

	if configurations, ok := s.D.GetOkExists("configurations"); ok {
		interfaces := configurations.([]interface{})
		tmp := make([]oci_core.HostGroupConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configurations", stateDataIndex)
			converted, err := s.mapToHostGroupConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configurations") {
			request.Configurations = tmp
		}
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

	if isTargetedPlacementRequired, ok := s.D.GetOkExists("is_targeted_placement_required"); ok {
		tmp := isTargetedPlacementRequired.(bool)
		request.IsTargetedPlacementRequired = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateComputeHostGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "computehostgroup", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *CoreComputeHostGroupResourceCrud) Delete() error {
	request := oci_core.DeleteComputeHostGroupRequest{}

	tmp := s.D.Id()
	request.ComputeHostGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteComputeHostGroup(context.Background(), request)
	return err
}

func (s *CoreComputeHostGroupResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	configurations := []interface{}{}
	for _, item := range s.Res.Configurations {
		configurations = append(configurations, HostGroupConfigurationToMap(item))
	}
	s.D.Set("configurations", configurations)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsTargetedPlacementRequired != nil {
		s.D.Set("is_targeted_placement_required", *s.Res.IsTargetedPlacementRequired)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ComputeHostGroupSummaryToMap(obj oci_core.ComputeHostGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsTargetedPlacementRequired != nil {
		result["is_targeted_placement_required"] = bool(*obj.IsTargetedPlacementRequired)
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

func (s *CoreComputeHostGroupResourceCrud) mapToHostGroupConfiguration(fieldKeyFormat string) (oci_core.HostGroupConfiguration, error) {
	result := oci_core.HostGroupConfiguration{}

	if firmwareBundleId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "firmware_bundle_id")); ok {
		tmp := firmwareBundleId.(string)
		result.FirmwareBundleId = &tmp
	}

	if recycleLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recycle_level")); ok {
		result.RecycleLevel = oci_core.HostGroupConfigurationRecycleLevelEnum(recycleLevel.(string))
	}

	if state, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "state")); ok {
		result.State = oci_core.HostGroupConfigurationStateEnum(state.(string))
	}

	if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
		tmp := target.(string)
		result.Target = &tmp
	}

	return result, nil
}

func HostGroupConfigurationToMap(obj oci_core.HostGroupConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FirmwareBundleId != nil {
		result["firmware_bundle_id"] = string(*obj.FirmwareBundleId)
	}

	result["recycle_level"] = string(obj.RecycleLevel)

	result["state"] = string(obj.State)

	if obj.Target != nil {
		result["target"] = string(*obj.Target)
	}

	return result
}

func (s *CoreComputeHostGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeComputeHostGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ComputeHostGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.ChangeComputeHostGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "computehostgroup", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
