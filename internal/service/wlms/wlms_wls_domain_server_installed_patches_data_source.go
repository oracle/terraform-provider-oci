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

func WlmsWlsDomainServerInstalledPatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWlmsWlsDomainServerInstalledPatches,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"wls_domain_id": {
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

func readWlmsWlsDomainServerInstalledPatches(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsWlsDomainServerInstalledPatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsWlsDomainServerInstalledPatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.ListWlsDomainServerInstalledPatchesResponse
}

func (s *WlmsWlsDomainServerInstalledPatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsWlsDomainServerInstalledPatchesDataSourceCrud) Get() error {
	request := oci_wlms.ListWlsDomainServerInstalledPatchesRequest{}

	if serverId, ok := s.D.GetOkExists("server_id"); ok {
		tmp := serverId.(string)
		request.ServerId = &tmp
	}

	if wlsDomainId, ok := s.D.GetOkExists("wls_domain_id"); ok {
		tmp := wlsDomainId.(string)
		request.WlsDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.ListWlsDomainServerInstalledPatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWlsDomainServerInstalledPatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WlmsWlsDomainServerInstalledPatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WlmsWlsDomainServerInstalledPatchesDataSource-", WlmsWlsDomainServerInstalledPatchesDataSource(), s.D))
	resources := []map[string]interface{}{}
	wlsDomainServerInstalledPatch := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, WlsDomainInstalledPatchSummaryToMap(item))
	}
	wlsDomainServerInstalledPatch["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WlmsWlsDomainServerInstalledPatchesDataSource().Schema["installed_patch_collection"].Elem.(*schema.Resource).Schema)
		wlsDomainServerInstalledPatch["items"] = items
	}

	resources = append(resources, wlsDomainServerInstalledPatch)
	if err := s.D.Set("installed_patch_collection", resources); err != nil {
		return err
	}

	return nil
}

func WlsDomainInstalledPatchSummaryToMap(obj oci_wlms.InstalledPatchSummary) map[string]interface{} {
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
