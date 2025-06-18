// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package wlms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_wlms "github.com/oracle/oci-go-sdk/v65/wlms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WlmsManagedInstanceServerInstalledPatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWlmsManagedInstanceServerInstalledPatches,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"installed_patch_collection": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readWlmsManagedInstanceServerInstalledPatches(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsManagedInstanceServerInstalledPatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsManagedInstanceServerInstalledPatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.ListManagedInstanceServerInstalledPatchesResponse
}

func (s *WlmsManagedInstanceServerInstalledPatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsManagedInstanceServerInstalledPatchesDataSourceCrud) Get() error {
	request := oci_wlms.ListManagedInstanceServerInstalledPatchesRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if serverId, ok := s.D.GetOkExists("server_id"); ok {
		tmp := serverId.(string)
		request.ServerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.ListManagedInstanceServerInstalledPatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstanceServerInstalledPatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WlmsManagedInstanceServerInstalledPatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WlmsManagedInstanceServerInstalledPatchesDataSource-", WlmsManagedInstanceServerInstalledPatchesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedInstanceServerInstalledPatch := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InstalledPatchSummaryToMap(item))
	}
	managedInstanceServerInstalledPatch["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WlmsManagedInstanceServerInstalledPatchesDataSource().Schema["installed_patch_collection"].Elem.(*schema.Resource).Schema)
		managedInstanceServerInstalledPatch["items"] = items
	}

	resources = append(resources, managedInstanceServerInstalledPatch)
	if err := s.D.Set("installed_patch_collection", resources); err != nil {
		return err
	}

	return nil
}

func InstalledPatchSummaryToMap(obj oci_wlms.InstalledPatchSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}
