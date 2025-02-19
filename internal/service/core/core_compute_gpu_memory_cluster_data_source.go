// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreComputeGpuMemoryClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compute_gpu_memory_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreComputeGpuMemoryClusterResource(), fieldMap, readSingularCoreComputeGpuMemoryCluster)
}

func readSingularCoreComputeGpuMemoryCluster(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeGpuMemoryClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeGpuMemoryClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetComputeGpuMemoryClusterResponse
}

func (s *CoreComputeGpuMemoryClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeGpuMemoryClusterDataSourceCrud) Get() error {
	request := oci_core.GetComputeGpuMemoryClusterRequest{}

	if computeGpuMemoryClusterId, ok := s.D.GetOkExists("compute_gpu_memory_cluster_id"); ok {
		tmp := computeGpuMemoryClusterId.(string)
		request.ComputeGpuMemoryClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetComputeGpuMemoryCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreComputeGpuMemoryClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeClusterId != nil {
		s.D.Set("compute_cluster_id", *s.Res.ComputeClusterId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GpuMemoryFabricId != nil {
		s.D.Set("gpu_memory_fabric_id", *s.Res.GpuMemoryFabricId)
	}

	if s.Res.InstanceConfigurationId != nil {
		s.D.Set("instance_configuration_id", *s.Res.InstanceConfigurationId)
	}

	if s.Res.Size != nil {
		s.D.Set("size", strconv.FormatInt(*s.Res.Size, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
