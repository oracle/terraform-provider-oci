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

func CoreInstancePoolsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreInstancePools,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_pools": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreInstancePoolResource()),
			},
		},
	}
}

func readCoreInstancePools(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.ReadResource(sync)
}

type CoreInstancePoolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.ListInstancePoolsResponse
}

func (s *CoreInstancePoolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstancePoolsDataSourceCrud) Get() error {
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

func (s *CoreInstancePoolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreInstancePoolsDataSource-", CoreInstancePoolsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instancePool := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			instancePool["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			instancePool["display_name"] = *r.DisplayName
		}

		instancePool["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			instancePool["id"] = *r.Id
		}

		if r.InstanceConfigurationId != nil {
			instancePool["instance_configuration_id"] = *r.InstanceConfigurationId
		}

		if r.Size != nil {
			instancePool["size"] = *r.Size
			instancePool["actual_size"] = *r.Size
		}

		instancePool["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			instancePool["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, instancePool)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreInstancePoolsDataSource().Schema["instance_pools"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instance_pools", resources); err != nil {
		return err
	}

	return nil
}
