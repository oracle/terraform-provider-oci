// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreComputeCapacityTopologyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compute_capacity_topology_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreComputeCapacityTopologyResource(), fieldMap, readSingularCoreComputeCapacityTopology)
}

func readSingularCoreComputeCapacityTopology(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityTopologyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeCapacityTopologyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetComputeCapacityTopologyResponse
}

func (s *CoreComputeCapacityTopologyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeCapacityTopologyDataSourceCrud) Get() error {
	request := oci_core.GetComputeCapacityTopologyRequest{}

	if computeCapacityTopologyId, ok := s.D.GetOkExists("compute_capacity_topology_id"); ok {
		tmp := computeCapacityTopologyId.(string)
		request.ComputeCapacityTopologyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetComputeCapacityTopology(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreComputeCapacityTopologyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CapacitySource != nil {
		capacitySourceArray := []interface{}{}
		if capacitySourceMap := CapacitySourceToMap(&s.Res.CapacitySource); capacitySourceMap != nil {
			capacitySourceArray = append(capacitySourceArray, capacitySourceMap)
		}
		s.D.Set("capacity_source", capacitySourceArray)
	} else {
		s.D.Set("capacity_source", nil)
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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
