// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreComputeGpuMemoryFabricDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compute_gpu_memory_fabric_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreComputeGpuMemoryFabricResource(), fieldMap, readSingularCoreComputeGpuMemoryFabric)
}

func readSingularCoreComputeGpuMemoryFabric(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeGpuMemoryFabricDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeGpuMemoryFabricDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetComputeGpuMemoryFabricResponse
}

func (s *CoreComputeGpuMemoryFabricDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeGpuMemoryFabricDataSourceCrud) Get() error {
	request := oci_core.GetComputeGpuMemoryFabricRequest{}

	if computeGpuMemoryFabricId, ok := s.D.GetOkExists("compute_gpu_memory_fabric_id"); ok {
		tmp := computeGpuMemoryFabricId.(string)
		request.ComputeGpuMemoryFabricId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetComputeGpuMemoryFabric(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreComputeGpuMemoryFabricDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

func flattenAdditionalData(input map[string]interface{}) map[string]interface{} {
	flatMap := make(map[string]interface{})

	for k, v := range input {
		if strVal, ok := v.(string); ok {
			flatMap[k] = strVal
		} else {
			jsonStrVal, err := json.Marshal(v)
			if err != nil {
				log.Printf("[ERROR] Failed to marshal additional_data[%q]: %v", k, err)
				flatMap[k] = "" // Optional fallback
			} else {
				flatMap[k] = string(jsonStrVal)
			}
		}
	}

	return flatMap
}
