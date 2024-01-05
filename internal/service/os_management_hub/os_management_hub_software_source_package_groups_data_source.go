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

func OsManagementHubSoftwareSourcePackageGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubSoftwareSourcePackageGroups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_source_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"package_group_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
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
									"display_order": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"group_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_default": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_user_visible": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"packages": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"repositories": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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

func readOsManagementHubSoftwareSourcePackageGroups(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourcePackageGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwareSourcePackageGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.ListPackageGroupsResponse
}

func (s *OsManagementHubSoftwareSourcePackageGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwareSourcePackageGroupsDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListPackageGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if groupType, ok := s.D.GetOkExists("group_type"); ok {
		interfaces := groupType.([]interface{})
		tmp := make([]oci_os_management_hub.PackageGroupGroupTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.PackageGroupGroupTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_type") {
			request.GroupType = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListPackageGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPackageGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubSoftwareSourcePackageGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubSoftwareSourcePackageGroupsDataSource-", OsManagementHubSoftwareSourcePackageGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	softwareSourcePackageGroup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PackageGroupSummaryToMap(item))
	}
	softwareSourcePackageGroup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubSoftwareSourcePackageGroupsDataSource().Schema["package_group_collection"].Elem.(*schema.Resource).Schema)
		softwareSourcePackageGroup["items"] = items
	}

	resources = append(resources, softwareSourcePackageGroup)
	if err := s.D.Set("package_group_collection", resources); err != nil {
		return err
	}

	return nil
}

func PackageGroupSummaryToMap(obj oci_os_management_hub.PackageGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayOrder != nil {
		result["display_order"] = int(*obj.DisplayOrder)
	}

	result["group_type"] = string(obj.GroupType)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsDefault != nil {
		result["is_default"] = bool(*obj.IsDefault)
	}

	if obj.IsUserVisible != nil {
		result["is_user_visible"] = bool(*obj.IsUserVisible)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["repositories"] = obj.Repositories

	return result
}
