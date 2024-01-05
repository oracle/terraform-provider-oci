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

func OsManagementHubManagedInstanceGroupAvailableSoftwareSourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubManagedInstanceGroupAvailableSoftwareSources,
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
			"managed_instance_group_id": {
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

func readOsManagementHubManagedInstanceGroupAvailableSoftwareSources(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceGroupAvailableSoftwareSourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceGroupClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagedInstanceGroupAvailableSoftwareSourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceGroupClient
	Res    *oci_os_management_hub.ListManagedInstanceGroupAvailableSoftwareSourcesResponse
}

func (s *OsManagementHubManagedInstanceGroupAvailableSoftwareSourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagedInstanceGroupAvailableSoftwareSourcesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListManagedInstanceGroupAvailableSoftwareSourcesRequest{}

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

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListManagedInstanceGroupAvailableSoftwareSources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstanceGroupAvailableSoftwareSources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubManagedInstanceGroupAvailableSoftwareSourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubManagedInstanceGroupAvailableSoftwareSourcesDataSource-", OsManagementHubManagedInstanceGroupAvailableSoftwareSourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedInstanceGroupAvailableSoftwareSource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AvailableSoftwareSourceSummaryToMap(item))
	}
	managedInstanceGroupAvailableSoftwareSource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubManagedInstanceGroupAvailableSoftwareSourcesDataSource().Schema["available_software_source_collection"].Elem.(*schema.Resource).Schema)
		managedInstanceGroupAvailableSoftwareSource["items"] = items
	}

	resources = append(resources, managedInstanceGroupAvailableSoftwareSource)
	if err := s.D.Set("available_software_source_collection", resources); err != nil {
		return err
	}

	return nil
}

func AvailableSoftwareSourceSummaryToMap(obj oci_os_management_hub.AvailableSoftwareSourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}
