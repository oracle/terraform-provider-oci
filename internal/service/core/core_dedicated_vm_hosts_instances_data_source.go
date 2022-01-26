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

func CoreDedicatedVmHostsInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDedicatedVmHostsInstances,
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
			"dedicated_vm_host_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dedicated_vm_host_instances": {
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
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreDedicatedVmHostsInstances(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostsInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreDedicatedVmHostsInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListDedicatedVmHostInstancesResponse
}

func (s *CoreDedicatedVmHostsInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDedicatedVmHostsInstancesDataSourceCrud) Get() error {
	request := oci_core.ListDedicatedVmHostInstancesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dedicatedVmHostId, ok := s.D.GetOkExists("dedicated_vm_host_id"); ok {
		tmp := dedicatedVmHostId.(string)
		request.DedicatedVmHostId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDedicatedVmHostInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDedicatedVmHostInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDedicatedVmHostsInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDedicatedVmHostsInstancesDataSource-", CoreDedicatedVmHostsInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dedicatedVmHostsInstance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			dedicatedVmHostsInstance["availability_domain"] = *r.AvailabilityDomain
		}

		if r.InstanceId != nil {
			dedicatedVmHostsInstance["instance_id"] = *r.InstanceId
		}

		if r.Shape != nil {
			dedicatedVmHostsInstance["shape"] = *r.Shape
		}

		if r.TimeCreated != nil {
			dedicatedVmHostsInstance["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, dedicatedVmHostsInstance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDedicatedVmHostsInstancesDataSource().Schema["dedicated_vm_host_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("dedicated_vm_host_instances", resources); err != nil {
		return err
	}

	return nil
}
