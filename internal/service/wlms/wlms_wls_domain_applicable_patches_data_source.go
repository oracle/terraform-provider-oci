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

func WlmsWlsDomainApplicablePatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWlmsWlsDomainApplicablePatches,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"wls_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"applicable_patch_collection": {
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
									"middleware_type": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"os_arch": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"weblogic_version": {
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

func readWlmsWlsDomainApplicablePatches(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsWlsDomainApplicablePatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsWlsDomainApplicablePatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.ListApplicablePatchesResponse
}

func (s *WlmsWlsDomainApplicablePatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsWlsDomainApplicablePatchesDataSourceCrud) Get() error {
	request := oci_wlms.ListApplicablePatchesRequest{}

	if wlsDomainId, ok := s.D.GetOkExists("wls_domain_id"); ok {
		tmp := wlsDomainId.(string)
		request.WlsDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.ListApplicablePatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListApplicablePatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WlmsWlsDomainApplicablePatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WlmsWlsDomainApplicablePatchesDataSource-", WlmsWlsDomainApplicablePatchesDataSource(), s.D))
	resources := []map[string]interface{}{}
	wlsDomainApplicablePatch := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ApplicablePatchSummaryToMap(item))
	}
	wlsDomainApplicablePatch["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WlmsWlsDomainApplicablePatchesDataSource().Schema["applicable_patch_collection"].Elem.(*schema.Resource).Schema)
		wlsDomainApplicablePatch["items"] = items
	}

	resources = append(resources, wlsDomainApplicablePatch)
	if err := s.D.Set("applicable_patch_collection", resources); err != nil {
		return err
	}

	return nil
}

func ApplicablePatchSummaryToMap(obj oci_wlms.ApplicablePatchSummary) map[string]interface{} {
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

	result["middleware_type"] = obj.MiddlewareType

	if obj.OsArch != nil {
		result["os_arch"] = string(*obj.OsArch)
	}

	if obj.WeblogicVersion != nil {
		result["weblogic_version"] = string(*obj.WeblogicVersion)
	}

	return result
}
