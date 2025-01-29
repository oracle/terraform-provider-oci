// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneLinksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readTenantmanagercontrolplaneLinks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"child_tenancy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_tenancy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"link_collection": {
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
									"child_tenancy_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"parent_tenancy_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_terminated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

func readTenantmanagercontrolplaneLinks(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneLinksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LinkClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneLinksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.LinkClient
	Res    *oci_tenantmanagercontrolplane.ListLinksResponse
}

func (s *TenantmanagercontrolplaneLinksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneLinksDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.ListLinksRequest{}

	if childTenancyId, ok := s.D.GetOkExists("child_tenancy_id"); ok {
		tmp := childTenancyId.(string)
		request.ChildTenancyId = &tmp
	}

	if parentTenancyId, ok := s.D.GetOkExists("parent_tenancy_id"); ok {
		tmp := parentTenancyId.(string)
		request.ParentTenancyId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_tenantmanagercontrolplane.ListLinksLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.ListLinks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLinks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *TenantmanagercontrolplaneLinksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("TenantmanagercontrolplaneLinksDataSource-", TenantmanagercontrolplaneLinksDataSource(), s.D))
	resources := []map[string]interface{}{}
	link := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LinkSummaryToMap(item))
	}
	link["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, TenantmanagercontrolplaneLinksDataSource().Schema["link_collection"].Elem.(*schema.Resource).Schema)
		link["items"] = items
	}

	resources = append(resources, link)
	if err := s.D.Set("link_collection", resources); err != nil {
		return err
	}

	return nil
}

func LinkSummaryToMap(obj oci_tenantmanagercontrolplane.LinkSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ChildTenancyId != nil {
		result["child_tenancy_id"] = string(*obj.ChildTenancyId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ParentTenancyId != nil {
		result["parent_tenancy_id"] = string(*obj.ParentTenancyId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeTerminated != nil {
		result["time_terminated"] = obj.TimeTerminated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
