// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreDedicatedVmHostInstanceShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDedicatedVmHostInstanceShapes,
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
			"dedicated_vm_host_shape": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dedicated_vm_host_instance_shapes": {
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
						"instance_shape_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreDedicatedVmHostInstanceShapes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostInstanceShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreDedicatedVmHostInstanceShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListDedicatedVmHostInstanceShapesResponse
}

func (s *CoreDedicatedVmHostInstanceShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDedicatedVmHostInstanceShapesDataSourceCrud) Get() error {
	request := oci_core.ListDedicatedVmHostInstanceShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dedicatedVmHostShape, ok := s.D.GetOkExists("dedicated_vm_host_shape"); ok {
		tmp := dedicatedVmHostShape.(string)
		request.DedicatedVmHostShape = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDedicatedVmHostInstanceShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDedicatedVmHostInstanceShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDedicatedVmHostInstanceShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDedicatedVmHostInstanceShapesDataSource-", CoreDedicatedVmHostInstanceShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dedicatedVmHostInstanceShape := map[string]interface{}{}

		if r.AvailabilityDomain != nil {
			dedicatedVmHostInstanceShape["availability_domain"] = *r.AvailabilityDomain
		}

		if r.InstanceShapeName != nil {
			dedicatedVmHostInstanceShape["instance_shape_name"] = *r.InstanceShapeName
		}

		resources = append(resources, dedicatedVmHostInstanceShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDedicatedVmHostInstanceShapesDataSource().Schema["dedicated_vm_host_instance_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("dedicated_vm_host_instance_shapes", resources); err != nil {
		return err
	}

	return nil
}
