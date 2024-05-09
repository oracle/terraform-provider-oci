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

func OsManagementHubManagedInstanceGroupAvailablePackagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubManagedInstanceGroupAvailablePackages,
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
			"is_latest": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"managed_instance_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_instance_group_available_package_collection": {
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
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_latest": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"software_sources": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"description": {
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
												"is_mandatory_for_autonomous_linux": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"software_source_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
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

func readOsManagementHubManagedInstanceGroupAvailablePackages(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceGroupAvailablePackagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceGroupClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagedInstanceGroupAvailablePackagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceGroupClient
	Res    *oci_os_management_hub.ListManagedInstanceGroupAvailablePackagesResponse
}

func (s *OsManagementHubManagedInstanceGroupAvailablePackagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagedInstanceGroupAvailablePackagesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListManagedInstanceGroupAvailablePackagesRequest{}

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

	if isLatest, ok := s.D.GetOkExists("is_latest"); ok {
		tmp := isLatest.(bool)
		request.IsLatest = &tmp
	}

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListManagedInstanceGroupAvailablePackages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstanceGroupAvailablePackages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubManagedInstanceGroupAvailablePackagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubManagedInstanceGroupAvailablePackagesDataSource-", OsManagementHubManagedInstanceGroupAvailablePackagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedInstanceGroupAvailablePackage := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedInstanceGroupAvailablePackageSummaryToMap(item))
	}
	managedInstanceGroupAvailablePackage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubManagedInstanceGroupAvailablePackagesDataSource().Schema["managed_instance_group_available_package_collection"].Elem.(*schema.Resource).Schema)
		managedInstanceGroupAvailablePackage["items"] = items
	}

	resources = append(resources, managedInstanceGroupAvailablePackage)
	if err := s.D.Set("managed_instance_group_available_package_collection", resources); err != nil {
		return err
	}

	return nil
}

func ManagedInstanceGroupAvailablePackageSummaryToMap(obj oci_os_management_hub.ManagedInstanceGroupAvailablePackageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["architecture"] = string(obj.Architecture)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.IsLatest != nil {
		result["is_latest"] = bool(*obj.IsLatest)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	softwareSources := []interface{}{}
	for _, item := range obj.SoftwareSources {
		softwareSources = append(softwareSources, SoftwareSourceDetailsToMap(item))
	}
	result["software_sources"] = softwareSources

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
