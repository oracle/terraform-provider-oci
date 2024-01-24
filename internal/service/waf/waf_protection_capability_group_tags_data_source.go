// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_waf "github.com/oracle/oci-go-sdk/v65/waf"
)

func WafProtectionCapabilityGroupTagsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWafProtectionCapabilityGroupTags,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protection_capability_group_tag_collection": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readWafProtectionCapabilityGroupTags(d *schema.ResourceData, m interface{}) error {
	sync := &WafProtectionCapabilityGroupTagsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.ReadResource(sync)
}

type WafProtectionCapabilityGroupTagsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waf.WafClient
	Res    *oci_waf.ListProtectionCapabilityGroupTagsResponse
}

func (s *WafProtectionCapabilityGroupTagsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WafProtectionCapabilityGroupTagsDataSourceCrud) Get() error {
	request := oci_waf.ListProtectionCapabilityGroupTagsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_waf.ProtectionCapabilitySummaryTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waf")

	response, err := s.Client.ListProtectionCapabilityGroupTags(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProtectionCapabilityGroupTags(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WafProtectionCapabilityGroupTagsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WafProtectionCapabilityGroupTagsDataSource-", WafProtectionCapabilityGroupTagsDataSource(), s.D))
	resources := []map[string]interface{}{}
	protectionCapabilityGroupTag := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProtectionCapabilityGroupTagSummaryToMap(item))
	}
	protectionCapabilityGroupTag["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WafProtectionCapabilityGroupTagsDataSource().Schema["protection_capability_group_tag_collection"].Elem.(*schema.Resource).Schema)
		protectionCapabilityGroupTag["items"] = items
	}

	resources = append(resources, protectionCapabilityGroupTag)
	if err := s.D.Set("protection_capability_group_tag_collection", resources); err != nil {
		return err
	}

	return nil
}

func ProtectionCapabilityGroupTagSummaryToMap(obj oci_waf.ProtectionCapabilityGroupTagSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
