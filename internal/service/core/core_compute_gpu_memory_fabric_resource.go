// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreComputeGpuMemoryFabricResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreComputeGpuMemoryFabric,
		Read:     readCoreComputeGpuMemoryFabric,
		Update:   updateCoreComputeGpuMemoryFabric,
		Delete:   deleteCoreComputeGpuMemoryFabric,
		Schema: map[string]*schema.Schema{
			// Required
			"compute_gpu_memory_fabric_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

			// Computed
			"additional_data": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"available_host_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compute_hpc_island_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compute_local_block_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compute_network_block_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fabric_health": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"healthy_host_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"total_host_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreComputeGpuMemoryFabric(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeGpuMemoryFabricResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	compartment, ok := sync.D.GetOkExists("compartment_id")

	err := tfresource.CreateResource(d, sync)
	if err != nil {
		return err
	}

	if ok && compartment != *sync.Res.CompartmentId {
		err = sync.updateCompartment(compartment)
		if err != nil {
			return err
		}
		tmp := compartment.(string)
		sync.Res.CompartmentId = &tmp
		err := sync.Get()
		if err != nil {
			log.Printf("error doing a Get() after compartment update: %v", err)
		}
		err = sync.SetData()
		if err != nil {
			log.Printf("error doing a SetData() after compartment update: %v", err)
		}
	}
	return nil
}

func readCoreComputeGpuMemoryFabric(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeGpuMemoryFabricResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreComputeGpuMemoryFabric(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeGpuMemoryFabricResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreComputeGpuMemoryFabric(d *schema.ResourceData, m interface{}) error {
	return nil
}

type CoreComputeGpuMemoryFabricResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.ComputeGpuMemoryFabric
	DisableNotFoundRetries bool
}

func (s *CoreComputeGpuMemoryFabricResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreComputeGpuMemoryFabricResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ComputeGpuMemoryFabricLifecycleStateProvisioning),
	}
}

func (s *CoreComputeGpuMemoryFabricResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ComputeGpuMemoryFabricLifecycleStateAvailable),
	}
}

func (s *CoreComputeGpuMemoryFabricResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.ComputeGpuMemoryFabricLifecycleStateUnavailable),
	}
}

func (s *CoreComputeGpuMemoryFabricResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *CoreComputeGpuMemoryFabricResourceCrud) Create() error {
	request := oci_core.UpdateComputeGpuMemoryFabricRequest{}

	if computeGpuMemoryFabricId, ok := s.D.GetOkExists("compute_gpu_memory_fabric_id"); ok {
		tmp := computeGpuMemoryFabricId.(string)
		request.ComputeGpuMemoryFabricId = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateComputeGpuMemoryFabric(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeGpuMemoryFabric
	return nil
}

func (s *CoreComputeGpuMemoryFabricResourceCrud) Get() error {
	request := oci_core.GetComputeGpuMemoryFabricRequest{}

	tmp := s.D.Id()
	request.ComputeGpuMemoryFabricId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetComputeGpuMemoryFabric(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeGpuMemoryFabric
	return nil
}

func (s *CoreComputeGpuMemoryFabricResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateComputeGpuMemoryFabricRequest{}

	tmp := s.D.Id()
	request.ComputeGpuMemoryFabricId = &tmp

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

	response, err := s.Client.UpdateComputeGpuMemoryFabric(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeGpuMemoryFabric
	return nil
}

func (s *CoreComputeGpuMemoryFabricResourceCrud) SetData() error {
	s.D.Set("additional_data", flattenAdditionalData(s.Res.AdditionalData))

	if s.Res.AvailableHostCount != nil {
		s.D.Set("available_host_count", strconv.FormatInt(*s.Res.AvailableHostCount, 10))
	}

	s.D.Set("compute_gpu_memory_fabric_id", *s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeHpcIslandId != nil {
		s.D.Set("compute_hpc_island_id", *s.Res.ComputeHpcIslandId)
	}

	if s.Res.ComputeLocalBlockId != nil {
		s.D.Set("compute_local_block_id", *s.Res.ComputeLocalBlockId)
	}

	if s.Res.ComputeNetworkBlockId != nil {
		s.D.Set("compute_network_block_id", *s.Res.ComputeNetworkBlockId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("fabric_health", s.Res.FabricHealth)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HealthyHostCount != nil {
		s.D.Set("healthy_host_count", strconv.FormatInt(*s.Res.HealthyHostCount, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TotalHostCount != nil {
		s.D.Set("total_host_count", strconv.FormatInt(*s.Res.TotalHostCount, 10))
	}

	return nil
}

func ComputeGpuMemoryFabricSummaryToMap(obj oci_core.ComputeGpuMemoryFabricSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailableHostCount != nil {
		result["available_host_count"] = strconv.FormatInt(*obj.AvailableHostCount, 10)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComputeHpcIslandId != nil {
		result["compute_hpc_island_id"] = string(*obj.ComputeHpcIslandId)
	}

	if obj.ComputeLocalBlockId != nil {
		result["compute_local_block_id"] = string(*obj.ComputeLocalBlockId)
	}

	if obj.ComputeNetworkBlockId != nil {
		result["compute_network_block_id"] = string(*obj.ComputeNetworkBlockId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["fabric_health"] = string(obj.FabricHealth)

	result["freeform_tags"] = obj.FreeformTags

	if obj.HealthyHostCount != nil {
		result["healthy_host_count"] = strconv.FormatInt(*obj.HealthyHostCount, 10)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Id != nil {
		result["compute_gpu_memory_fabric_id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TotalHostCount != nil {
		result["total_host_count"] = strconv.FormatInt(*obj.TotalHostCount, 10)
	}

	return result
}

func (s *CoreComputeGpuMemoryFabricResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeComputeGpuMemoryFabricCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ComputeGpuMemoryFabricId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeComputeGpuMemoryFabricCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
