// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func InstancePoolsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readInstancePools,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_pools": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(InstancePoolResource()),
			},
		},
	}
}

func readInstancePools(d *schema.ResourceData, m interface{}) error {
	sync := &InstancePoolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient

	return ReadResource(sync)
}

type InstancePoolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.ListInstancePoolsResponse
}

func (s *InstancePoolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *InstancePoolsDataSourceCrud) Get() error {
	request := oci_core.ListInstancePoolsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.InstancePoolSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListInstancePools(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstancePools(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *InstancePoolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instancePool := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DisplayName != nil {
			instancePool["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			instancePool["id"] = *r.Id
		}

		if r.InstanceConfigurationId != nil {
			instancePool["instance_configuration_id"] = *r.InstanceConfigurationId
		}

		if r.Size != nil {
			instancePool["size"] = *r.Size
		}

		instancePool["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			instancePool["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, instancePool)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, InstancePoolsDataSource().Schema["instance_pools"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instance_pools", resources); err != nil {
		return err
	}

	return nil
}
