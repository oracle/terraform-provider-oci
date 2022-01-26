// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v56/osmanagement"
)

func OsmanagementManagedInstanceGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsmanagementManagedInstanceGroups,
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
			"os_family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(OsmanagementManagedInstanceGroupResource()),
			},
		},
	}
}

func readOsmanagementManagedInstanceGroups(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

type OsmanagementManagedInstanceGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.ListManagedInstanceGroupsResponse
}

func (s *OsmanagementManagedInstanceGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementManagedInstanceGroupsDataSourceCrud) Get() error {
	request := oci_osmanagement.ListManagedInstanceGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		request.OsFamily = oci_osmanagement.ListManagedInstanceGroupsOsFamilyEnum(osFamily.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_osmanagement.ListManagedInstanceGroupsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osmanagement")

	response, err := s.Client.ListManagedInstanceGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstanceGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsmanagementManagedInstanceGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsmanagementManagedInstanceGroupsDataSource-", OsmanagementManagedInstanceGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		managedInstanceGroup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			managedInstanceGroup["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			managedInstanceGroup["description"] = *r.Description
		}

		if r.DisplayName != nil {
			managedInstanceGroup["display_name"] = *r.DisplayName
		}

		managedInstanceGroup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			managedInstanceGroup["id"] = *r.Id
		}

		if r.ManagedInstanceCount != nil {
			managedInstanceGroup["managed_instance_count"] = *r.ManagedInstanceCount
		}

		managedInstanceGroup["os_family"] = r.OsFamily

		managedInstanceGroup["state"] = r.LifecycleState
		managedInstanceGroup["managed_instance_count"] = r.ManagedInstanceCount

		resources = append(resources, managedInstanceGroup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsmanagementManagedInstanceGroupsDataSource().Schema["managed_instance_groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("managed_instance_groups", resources); err != nil {
		return err
	}

	return nil
}
