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

func OsmanagementManagedInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsmanagementManagedInstances,
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
			"managed_instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(OsmanagementManagedInstanceResource()),
			},
		},
	}
}

func readOsmanagementManagedInstances(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

type OsmanagementManagedInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.ListManagedInstancesResponse
}

func (s *OsmanagementManagedInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementManagedInstancesDataSourceCrud) Get() error {
	request := oci_osmanagement.ListManagedInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		request.OsFamily = oci_osmanagement.ListManagedInstancesOsFamilyEnum(osFamily.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osmanagement")

	response, err := s.Client.ListManagedInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsmanagementManagedInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsmanagementManagedInstancesDataSource-", OsmanagementManagedInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		managedInstance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Description != nil {
			managedInstance["description"] = *r.Description
		}

		if r.DisplayName != nil {
			managedInstance["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			managedInstance["id"] = *r.Id
		}

		if r.IsRebootRequired != nil {
			managedInstance["is_reboot_required"] = *r.IsRebootRequired
		}

		if r.LastBoot != nil {
			managedInstance["last_boot"] = *r.LastBoot
		}

		if r.LastCheckin != nil {
			managedInstance["last_checkin"] = *r.LastCheckin
		}

		managedInstance["os_family"] = r.OsFamily

		managedInstance["status"] = r.Status

		if r.UpdatesAvailable != nil {
			managedInstance["updates_available"] = *r.UpdatesAvailable
		}

		resources = append(resources, managedInstance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsmanagementManagedInstancesDataSource().Schema["managed_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("managed_instances", resources); err != nil {
		return err
	}

	return nil
}
