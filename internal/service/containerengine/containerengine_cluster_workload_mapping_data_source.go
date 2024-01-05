// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineClusterWorkloadMappingDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["workload_mapping_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ContainerengineClusterWorkloadMappingResource(), fieldMap, readSingularContainerengineClusterWorkloadMapping)
}

func readSingularContainerengineClusterWorkloadMapping(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterWorkloadMappingDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineClusterWorkloadMappingDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetWorkloadMappingResponse
}

func (s *ContainerengineClusterWorkloadMappingDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterWorkloadMappingDataSourceCrud) Get() error {
	request := oci_containerengine.GetWorkloadMappingRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if workloadMappingId, ok := s.D.GetOkExists("workload_mapping_id"); ok {
		tmp := workloadMappingId.(string)
		request.WorkloadMappingId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.GetWorkloadMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineClusterWorkloadMappingDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MappedCompartmentId != nil {
		s.D.Set("mapped_compartment_id", *s.Res.MappedCompartmentId)
	}

	if s.Res.MappedTenancyId != nil {
		s.D.Set("mapped_tenancy_id", *s.Res.MappedTenancyId)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
