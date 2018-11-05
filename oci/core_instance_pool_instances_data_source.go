// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func InstancePoolInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readInstancePoolInstances,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
						"display_name": {
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
						"instance_configuration_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
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

func readInstancePoolInstances(d *schema.ResourceData, m interface{}) error {
	sync := &InstancePoolInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient

	return ReadResource(sync)
}

type InstancePoolInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.ListInstancePoolInstancesResponse
}

func (s *InstancePoolInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *InstancePoolInstancesDataSourceCrud) Get() error {
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

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

func (s *InstancePoolInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
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
		resources = ApplyFilters(f.(*schema.Set), resources, InstancePoolInstancesDataSource().Schema["instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instances", resources); err != nil {
		return err
	}

	return nil
}
