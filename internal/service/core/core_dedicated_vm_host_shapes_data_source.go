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

func CoreDedicatedVmHostShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDedicatedVmHostShapes,
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
			"instance_shape_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dedicated_vm_host_shapes": {
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
						"dedicated_vm_host_shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreDedicatedVmHostShapes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreDedicatedVmHostShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListDedicatedVmHostShapesResponse
}

func (s *CoreDedicatedVmHostShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDedicatedVmHostShapesDataSourceCrud) Get() error {
	request := oci_core.ListDedicatedVmHostShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if instanceShapeName, ok := s.D.GetOkExists("instance_shape_name"); ok {
		tmp := instanceShapeName.(string)
		request.InstanceShapeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDedicatedVmHostShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDedicatedVmHostShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDedicatedVmHostShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDedicatedVmHostShapesDataSource-", CoreDedicatedVmHostShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dedicatedVmHostShape := map[string]interface{}{}

		if r.AvailabilityDomain != nil {
			dedicatedVmHostShape["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DedicatedVmHostShape != nil {
			dedicatedVmHostShape["dedicated_vm_host_shape"] = *r.DedicatedVmHostShape
		}

		resources = append(resources, dedicatedVmHostShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDedicatedVmHostShapesDataSource().Schema["dedicated_vm_host_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("dedicated_vm_host_shapes", resources); err != nil {
		return err
	}

	return nil
}
