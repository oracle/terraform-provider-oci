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

func CoreInstancePoolInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreInstancePoolInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_pool_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreInstancePoolInstanceResource()),
			},
		},
	}
}

func readCoreInstancePoolInstances(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.ReadResource(sync)
}

type CoreInstancePoolInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.ListInstancePoolInstancesResponse
}

func (s *CoreInstancePoolInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstancePoolInstancesDataSourceCrud) Get() error {
	request := oci_core.ListInstancePoolInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if instancePoolId, ok := s.D.GetOkExists("instance_pool_id"); ok {
		tmp := instancePoolId.(string)
		request.InstancePoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListInstancePoolInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstancePoolInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreInstancePoolInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreInstancePoolInstancesDataSource-", CoreInstancePoolInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instancePoolInstance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			instancePoolInstance["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DisplayName != nil {
			instancePoolInstance["display_name"] = *r.DisplayName
		}

		if r.FaultDomain != nil {
			instancePoolInstance["fault_domain"] = *r.FaultDomain
		}

		if r.Id != nil {
			instancePoolInstance["id"] = *r.Id
		}

		if r.InstanceConfigurationId != nil {
			instancePoolInstance["instance_configuration_id"] = *r.InstanceConfigurationId
		}

		loadBalancerBackends := []interface{}{}
		for _, item := range r.LoadBalancerBackends {
			loadBalancerBackends = append(loadBalancerBackends, InstancePoolInstanceLoadBalancerBackendToMap(item))
		}
		instancePoolInstance["load_balancer_backends"] = loadBalancerBackends

		if r.Region != nil {
			instancePoolInstance["region"] = *r.Region
		}

		if r.Shape != nil {
			instancePoolInstance["shape"] = *r.Shape
		}

		if r.State != nil {
			instancePoolInstance["state"] = *r.State
		}

		if r.TimeCreated != nil {
			instancePoolInstance["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, instancePoolInstance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreInstancePoolInstancesDataSource().Schema["instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instances", resources); err != nil {
		return err
	}

	return nil
}
