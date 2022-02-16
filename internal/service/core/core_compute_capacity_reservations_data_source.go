// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreComputeCapacityReservationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeCapacityReservations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_capacity_reservations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreComputeCapacityReservationResource()),
			},
		},
	}
}

func readCoreComputeCapacityReservations(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityReservationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeCapacityReservationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeCapacityReservationsResponse
}

func (s *CoreComputeCapacityReservationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeCapacityReservationsDataSourceCrud) Get() error {
	request := oci_core.ListComputeCapacityReservationsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.ComputeCapacityReservationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeCapacityReservations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeCapacityReservations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeCapacityReservationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		computeCapacityReservation := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			computeCapacityReservation["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DefinedTags != nil {
			computeCapacityReservation["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			computeCapacityReservation["display_name"] = *r.DisplayName
		}

		computeCapacityReservation["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			computeCapacityReservation["id"] = *r.Id
		}

		if r.IsDefaultReservation != nil {
			computeCapacityReservation["is_default_reservation"] = *r.IsDefaultReservation
		}

		if r.ReservedInstanceCount != nil {
			computeCapacityReservation["reserved_instance_count"] = strconv.FormatInt(*r.ReservedInstanceCount, 10)
		}

		computeCapacityReservation["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			computeCapacityReservation["time_created"] = r.TimeCreated.String()
		}

		if r.UsedInstanceCount != nil {
			computeCapacityReservation["used_instance_count"] = strconv.FormatInt(*r.UsedInstanceCount, 10)
		}

		resources = append(resources, computeCapacityReservation)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreComputeCapacityReservationsDataSource().Schema["compute_capacity_reservations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("compute_capacity_reservations", resources); err != nil {
		return err
	}

	return nil
}
