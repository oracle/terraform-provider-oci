// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreDedicatedVmHostDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["dedicated_vm_host_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreDedicatedVmHostResource(), fieldMap, readSingularCoreDedicatedVmHost)
}

func readSingularCoreDedicatedVmHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreDedicatedVmHostDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetDedicatedVmHostResponse
}

func (s *CoreDedicatedVmHostDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDedicatedVmHostDataSourceCrud) Get() error {
	request := oci_core.GetDedicatedVmHostRequest{}

	if dedicatedVmHostId, ok := s.D.GetOkExists("dedicated_vm_host_id"); ok {
		tmp := dedicatedVmHostId.(string)
		request.DedicatedVmHostId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetDedicatedVmHost(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreDedicatedVmHostDataSourceCrud) SetData() error {
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
