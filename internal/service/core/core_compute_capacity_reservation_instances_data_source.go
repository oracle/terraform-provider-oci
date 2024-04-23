// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreComputeCapacityReservationInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeCapacityReservationInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"capacity_reservation_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"capacity_reservation_instances": {
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
						"cluster_placement_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fault_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"ocpus": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readCoreComputeCapacityReservationInstances(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityReservationInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeCapacityReservationInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeCapacityReservationInstancesResponse
}

func (s *CoreComputeCapacityReservationInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeCapacityReservationInstancesDataSourceCrud) Get() error {
	request := oci_core.ListComputeCapacityReservationInstancesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if capacityReservationId, ok := s.D.GetOkExists("capacity_reservation_id"); ok {
		tmp := capacityReservationId.(string)
		request.CapacityReservationId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeCapacityReservationInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeCapacityReservationInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeCapacityReservationInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreComputeCapacityReservationInstancesDataSource-", CoreComputeCapacityReservationInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		computeCapacityReservationInstance := map[string]interface{}{}

		if r.AvailabilityDomain != nil {
			computeCapacityReservationInstance["availability_domain"] = *r.AvailabilityDomain
		}

		if r.ClusterPlacementGroupId != nil {
			computeCapacityReservationInstance["cluster_placement_group_id"] = *r.ClusterPlacementGroupId
		}

		if r.CompartmentId != nil {
			computeCapacityReservationInstance["compartment_id"] = *r.CompartmentId
		}

		if r.FaultDomain != nil {
			computeCapacityReservationInstance["fault_domain"] = *r.FaultDomain
		}

		if r.Id != nil {
			computeCapacityReservationInstance["id"] = *r.Id
		}

		if r.Shape != nil {
			computeCapacityReservationInstance["shape"] = *r.Shape
		}

		if r.ShapeConfig != nil {
			computeCapacityReservationInstance["shape_config"] = []interface{}{InstanceReservationShapeConfigDetailsToMap(r.ShapeConfig)}
		} else {
			computeCapacityReservationInstance["shape_config"] = nil
		}

		resources = append(resources, computeCapacityReservationInstance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreComputeCapacityReservationInstancesDataSource().Schema["capacity_reservation_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("capacity_reservation_instances", resources); err != nil {
		return err
	}

	return nil
}
