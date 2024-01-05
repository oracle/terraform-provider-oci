// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreComputeCapacityReservationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["capacity_reservation_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreComputeCapacityReservationResource(), fieldMap, readSingularCoreComputeCapacityReservation)
}

func readSingularCoreComputeCapacityReservation(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityReservationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeCapacityReservationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetComputeCapacityReservationResponse
}

func (s *CoreComputeCapacityReservationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeCapacityReservationDataSourceCrud) Get() error {
	request := oci_core.GetComputeCapacityReservationRequest{}

	if capacityReservationId, ok := s.D.GetOkExists("capacity_reservation_id"); ok {
		tmp := capacityReservationId.(string)
		request.CapacityReservationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetComputeCapacityReservation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreComputeCapacityReservationDataSourceCrud) SetData() error {
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
