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

	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreComputeClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreComputeCluster,
		Read:     readCoreComputeCluster,
		Update:   updateCoreComputeCluster,
		Delete:   deleteCoreComputeCluster,
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
			"placement_constraint_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"COMPUTE_CLUSTER",
							}, true),
						},

						// Optional
						"hpc_island_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"logical_placement_constraint": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"target_memory_fabric_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"target_network_block_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			// Computed
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
		},
	}
}

func createCoreComputeCluster(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreComputeCluster(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreComputeCluster(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreComputeCluster(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreComputeClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.ComputeCluster
	DisableNotFoundRetries bool
}

func (s *CoreComputeClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreComputeClusterResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *CoreComputeClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ComputeClusterLifecycleStateActive),
	}
}

func (s *CoreComputeClusterResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *CoreComputeClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.ComputeClusterLifecycleStateDeleted),
	}
}

func (s *CoreComputeClusterResourceCrud) Create() error {
	request := oci_core.CreateComputeClusterRequest{}

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

	response, err := s.Client.CreateComputeCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeCluster
	return nil
}

func (s *CoreComputeClusterResourceCrud) Get() error {
	request := oci_core.GetComputeClusterRequest{}

	tmp := s.D.Id()
	request.ComputeClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetComputeCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeCluster
	return nil
}

func (s *CoreComputeClusterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateComputeClusterRequest{}

	tmp := s.D.Id()
	request.ComputeClusterId = &tmp

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

	if placementConstraintDetails, ok := s.D.GetOkExists("placement_constraint_details"); ok && s.D.HasChange("placement_constraint_details") {
		if tmpList := placementConstraintDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "placement_constraint_details", 0)
			tmp, err := s.mapToPlacementConstraintDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			if computeClusterDetails, ok := tmp.(oci_core.ComputeClusterPlacementConstraintDetails); ok {
				computeClusterDetails.HpcIslandId = nil
				tmp = computeClusterDetails
			}
			request.PlacementConstraintDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateComputeCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeCluster
	return nil
}

func (s *CoreComputeClusterResourceCrud) Delete() error {
	request := oci_core.DeleteComputeClusterRequest{}

	tmp := s.D.Id()
	request.ComputeClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteComputeCluster(context.Background(), request)
	return err
}

func (s *CoreComputeClusterResourceCrud) SetData() error {
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

	if s.Res.PlacementConstraintDetails != nil {
		placementConstraintDetailsArray := []interface{}{}
		if placementConstraintDetailsMap := ComputeClusterPlacementConstraintDetailsToMap(&s.Res.PlacementConstraintDetails); placementConstraintDetailsMap != nil {
			placementConstraintDetailsArray = append(placementConstraintDetailsArray, placementConstraintDetailsMap)
		}
		s.D.Set("placement_constraint_details", placementConstraintDetailsArray)
	} else {
		s.D.Set("placement_constraint_details", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ComputeClusterSummaryToMap(obj oci_core.ComputeClusterSummary) map[string]interface{} {
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

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *CoreComputeClusterResourceCrud) mapToPlacementConstraintDetails(fieldKeyFormat string) (oci_core.PlacementConstraintDetails, error) {
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
	case strings.ToLower("COMPUTE_CLUSTER"):
		details := oci_core.ComputeClusterPlacementConstraintDetails{}
		if hpcIslandId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hpc_island_id")); ok {
			tmp := hpcIslandId.(string)
			details.HpcIslandId = &tmp
		}
		if logicalPlacementConstraint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logical_placement_constraint")); ok {
			details.LogicalPlacementConstraint = oci_core.ComputeClusterLogicalPlacementConstraintEnum(logicalPlacementConstraint.(string))
		}
		if targetMemoryFabricIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_memory_fabric_ids")); ok {
			interfaces := targetMemoryFabricIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "target_memory_fabric_ids")) {
				details.TargetMemoryFabricIds = tmp
			}
		}
		if targetNetworkBlockIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_network_block_ids")); ok {
			interfaces := targetNetworkBlockIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "target_network_block_ids")) {
				details.TargetNetworkBlockIds = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ComputeClusterPlacementConstraintDetailsToMap(obj *oci_core.PlacementConstraintDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.ComputeClusterPlacementConstraintDetails:
		result["type"] = "COMPUTE_CLUSTER"

		if v.HpcIslandId != nil {
			result["hpc_island_id"] = string(*v.HpcIslandId)
		}

		result["logical_placement_constraint"] = string(v.LogicalPlacementConstraint)

		result["target_memory_fabric_ids"] = v.TargetMemoryFabricIds

		result["target_network_block_ids"] = v.TargetNetworkBlockIds
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreComputeClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeComputeClusterCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ComputeClusterId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeComputeClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
