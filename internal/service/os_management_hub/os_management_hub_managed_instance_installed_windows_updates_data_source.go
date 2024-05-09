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

func OsManagementHubManagedInstanceInstalledWindowsUpdatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubManagedInstanceInstalledWindowsUpdates,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"installed_windows_update_collection": {
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
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"update_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"update_type": {
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

func readOsManagementHubManagedInstanceInstalledWindowsUpdates(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceInstalledWindowsUpdatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagedInstanceInstalledWindowsUpdatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceClient
	Res    *oci_os_management_hub.ListManagedInstanceInstalledWindowsUpdatesResponse
}

func (s *OsManagementHubManagedInstanceInstalledWindowsUpdatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagedInstanceInstalledWindowsUpdatesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListManagedInstanceInstalledWindowsUpdatesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		interfaces := name.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("name") {
			request.Name = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListManagedInstanceInstalledWindowsUpdates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstanceInstalledWindowsUpdates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubManagedInstanceInstalledWindowsUpdatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubManagedInstanceInstalledWindowsUpdatesDataSource-", OsManagementHubManagedInstanceInstalledWindowsUpdatesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedInstanceInstalledWindowsUpdate := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InstalledWindowsUpdateSummaryToMap(item))
	}
	managedInstanceInstalledWindowsUpdate["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubManagedInstanceInstalledWindowsUpdatesDataSource().Schema["installed_windows_update_collection"].Elem.(*schema.Resource).Schema)
		managedInstanceInstalledWindowsUpdate["items"] = items
	}

	resources = append(resources, managedInstanceInstalledWindowsUpdate)
	if err := s.D.Set("installed_windows_update_collection", resources); err != nil {
		return err
	}

	return nil
}

func InstalledWindowsUpdateSummaryToMap(obj oci_os_management_hub.InstalledWindowsUpdateSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.UpdateId != nil {
		result["update_id"] = string(*obj.UpdateId)
	}

	result["update_type"] = string(obj.UpdateType)

	return result
}
