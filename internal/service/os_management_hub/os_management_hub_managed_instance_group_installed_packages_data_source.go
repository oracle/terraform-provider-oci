// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagedInstanceGroupInstalledPackagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubManagedInstanceGroupInstalledPackages,
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
			"time_install_date_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_install_date_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_group_installed_package_collection": {
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
									"architecture": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
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

func readOsManagementHubManagedInstanceGroupInstalledPackages(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceGroupInstalledPackagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceGroupClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagedInstanceGroupInstalledPackagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceGroupClient
	Res    *oci_os_management_hub.ListManagedInstanceGroupInstalledPackagesResponse
}

func (s *OsManagementHubManagedInstanceGroupInstalledPackagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagedInstanceGroupInstalledPackagesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListManagedInstanceGroupInstalledPackagesRequest{}

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

	if timeInstallDateEnd, ok := s.D.GetOkExists("time_install_date_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeInstallDateEnd.(string))
		if err != nil {
			return err
		}
		request.TimeInstallDateEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeInstallDateStart, ok := s.D.GetOkExists("time_install_date_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeInstallDateStart.(string))
		if err != nil {
			return err
		}
		request.TimeInstallDateStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListManagedInstanceGroupInstalledPackages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstanceGroupInstalledPackages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubManagedInstanceGroupInstalledPackagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubManagedInstanceGroupInstalledPackagesDataSource-", OsManagementHubManagedInstanceGroupInstalledPackagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedInstanceGroupInstalledPackage := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedInstanceGroupInstalledPackageSummaryToMap(item))
	}
	managedInstanceGroupInstalledPackage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubManagedInstanceGroupInstalledPackagesDataSource().Schema["managed_instance_group_installed_package_collection"].Elem.(*schema.Resource).Schema)
		managedInstanceGroupInstalledPackage["items"] = items
	}

	resources = append(resources, managedInstanceGroupInstalledPackage)
	if err := s.D.Set("managed_instance_group_installed_package_collection", resources); err != nil {
		return err
	}

	return nil
}

func ManagedInstanceGroupInstalledPackageSummaryToMap(obj oci_os_management_hub.ManagedInstanceGroupInstalledPackageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Architecture != nil {
		result["architecture"] = string(*obj.Architecture)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
