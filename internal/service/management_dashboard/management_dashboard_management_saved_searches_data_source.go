// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_dashboard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_management_dashboard "github.com/oracle/oci-go-sdk/v65/managementdashboard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementDashboardManagementSavedSearchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readManagementDashboardManagementSavedSearches,
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
			"management_saved_search_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ManagementDashboardManagementSavedSearchResource()),
						},
					},
				},
			},
		},
	}
}

func readManagementDashboardManagementSavedSearches(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementDashboardManagementSavedSearchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DashxApisClient()

	return tfresource.ReadResource(sync)
}

type ManagementDashboardManagementSavedSearchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_dashboard.DashxApisClient
	Res    *oci_management_dashboard.ListManagementSavedSearchesResponse
}

func (s *ManagementDashboardManagementSavedSearchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementDashboardManagementSavedSearchesDataSourceCrud) Get() error {
	request := oci_management_dashboard.ListManagementSavedSearchesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_dashboard")

	response, err := s.Client.ListManagementSavedSearches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagementSavedSearches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagementDashboardManagementSavedSearchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementDashboardManagementSavedSearchesDataSource-", ManagementDashboardManagementSavedSearchesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managementSavedSearch := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagementSavedSearchSummaryToMap(item))
	}
	managementSavedSearch["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ManagementDashboardManagementSavedSearchesDataSource().Schema["management_saved_search_collection"].Elem.(*schema.Resource).Schema)
		managementSavedSearch["items"] = items
	}

	resources = append(resources, managementSavedSearch)
	if err := s.D.Set("management_saved_search_collection", resources); err != nil {
		return err
	}

	return nil
}
