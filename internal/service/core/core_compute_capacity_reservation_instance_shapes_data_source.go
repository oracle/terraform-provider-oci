// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"
)

func CoreComputeCapacityReservationInstanceShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeCapacityReservationInstanceShapes,
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
			"compute_capacity_reservation_instance_shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreComputeCapacityReservationInstanceShapes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityReservationInstanceShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeCapacityReservationInstanceShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeCapacityReservationInstanceShapesResponse
}

func (s *CoreComputeCapacityReservationInstanceShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeCapacityReservationInstanceShapesDataSourceCrud) Get() error {
	request := oci_core.ListComputeCapacityReservationInstanceShapesRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeCapacityReservationInstanceShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeCapacityReservationInstanceShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeCapacityReservationInstanceShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreComputeCapacityReservationInstanceShapesDataSource-", CoreComputeCapacityReservationInstanceShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		computeCapacityReservationInstanceShape := map[string]interface{}{}

		if r.AvailabilityDomain != nil {
			computeCapacityReservationInstanceShape["availability_domain"] = *r.AvailabilityDomain
		}

		if r.InstanceShape != nil {
			computeCapacityReservationInstanceShape["instance_shape"] = *r.InstanceShape
		}

		resources = append(resources, computeCapacityReservationInstanceShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreComputeCapacityReservationInstanceShapesDataSource().Schema["compute_capacity_reservation_instance_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("compute_capacity_reservation_instance_shapes", resources); err != nil {
		return err
	}

	return nil
}
