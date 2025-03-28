// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubProfileAvailableSoftwareSourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubProfileAvailableSoftwareSources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"available_software_source_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubProfileAvailableSoftwareSources(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfileAvailableSoftwareSourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubProfileAvailableSoftwareSourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.OnboardingClient
	Res    *oci_os_management_hub.ListProfileAvailableSoftwareSourcesResponse
}

func (s *OsManagementHubProfileAvailableSoftwareSourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubProfileAvailableSoftwareSourcesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListProfileAvailableSoftwareSourcesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		interfaces := displayName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("display_name") {
			request.DisplayName = tmp
		}
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if profileId, ok := s.D.GetOkExists("profile_id"); ok {
		tmp := profileId.(string)
		request.ProfileId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListProfileAvailableSoftwareSources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProfileAvailableSoftwareSources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubProfileAvailableSoftwareSourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubProfileAvailableSoftwareSourcesDataSource-", OsManagementHubProfileAvailableSoftwareSourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	profileAvailableSoftwareSource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AvailableSoftwareSourceSummaryToMap(item))
	}
	profileAvailableSoftwareSource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubProfileAvailableSoftwareSourcesDataSource().Schema["available_software_source_collection"].Elem.(*schema.Resource).Schema)
		profileAvailableSoftwareSource["items"] = items
	}

	resources = append(resources, profileAvailableSoftwareSource)
	if err := s.D.Set("available_software_source_collection", resources); err != nil {
		return err
	}

	return nil
}
