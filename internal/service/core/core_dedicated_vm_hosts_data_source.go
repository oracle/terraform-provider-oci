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

func CoreDedicatedVmHostsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDedicatedVmHosts,
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
			"instance_shape_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remaining_memory_in_gbs_greater_than_or_equal_to": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"remaining_ocpus_greater_than_or_equal_to": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dedicated_vm_hosts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreDedicatedVmHostResource()),
			},
		},
	}
}

func readCoreDedicatedVmHosts(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreDedicatedVmHostsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListDedicatedVmHostsResponse
}

func (s *CoreDedicatedVmHostsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDedicatedVmHostsDataSourceCrud) Get() error {
	request := oci_core.ListDedicatedVmHostsRequest{}

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

	if instanceShapeName, ok := s.D.GetOkExists("instance_shape_name"); ok {
		tmp := instanceShapeName.(string)
		request.InstanceShapeName = &tmp
	}

	if remainingMemoryInGBsGreaterThanOrEqualTo, ok := s.D.GetOkExists("remaining_memory_in_gbs_greater_than_or_equal_to"); ok {
		tmp := float32(remainingMemoryInGBsGreaterThanOrEqualTo.(float64))
		request.RemainingMemoryInGBsGreaterThanOrEqualTo = &tmp
	}

	if remainingOcpusGreaterThanOrEqualTo, ok := s.D.GetOkExists("remaining_ocpus_greater_than_or_equal_to"); ok {
		tmp := float32(remainingOcpusGreaterThanOrEqualTo.(float64))
		request.RemainingOcpusGreaterThanOrEqualTo = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.ListDedicatedVmHostsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDedicatedVmHosts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDedicatedVmHosts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDedicatedVmHostsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDedicatedVmHostsDataSource-", CoreDedicatedVmHostsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dedicatedVmHost := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			dedicatedVmHost["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DedicatedVmHostShape != nil {
			dedicatedVmHost["dedicated_vm_host_shape"] = *r.DedicatedVmHostShape
		}

		if r.DisplayName != nil {
			dedicatedVmHost["display_name"] = *r.DisplayName
		}

		if r.FaultDomain != nil {
			dedicatedVmHost["fault_domain"] = *r.FaultDomain
		}

		if r.Id != nil {
			dedicatedVmHost["id"] = *r.Id
		}

		if r.RemainingMemoryInGBs != nil {
			dedicatedVmHost["remaining_memory_in_gbs"] = *r.RemainingMemoryInGBs
		}

		if r.RemainingOcpus != nil {
			dedicatedVmHost["remaining_ocpus"] = *r.RemainingOcpus
		}

		dedicatedVmHost["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			dedicatedVmHost["time_created"] = r.TimeCreated.String()
		}

		if r.TotalMemoryInGBs != nil {
			dedicatedVmHost["total_memory_in_gbs"] = *r.TotalMemoryInGBs
		}

		if r.TotalOcpus != nil {
			dedicatedVmHost["total_ocpus"] = *r.TotalOcpus
		}

		resources = append(resources, dedicatedVmHost)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDedicatedVmHostsDataSource().Schema["dedicated_vm_hosts"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("dedicated_vm_hosts", resources); err != nil {
		return err
	}

	return nil
}
