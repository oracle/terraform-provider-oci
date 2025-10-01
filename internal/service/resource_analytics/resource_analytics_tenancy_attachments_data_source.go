// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_resource_analytics "github.com/oracle/oci-go-sdk/v65/resourceanalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourceAnalyticsTenancyAttachmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readResourceAnalyticsTenancyAttachments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_analytics_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenancy_attachment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ResourceAnalyticsTenancyAttachmentResource()),
						},
					},
				},
			},
		},
	}
}

func readResourceAnalyticsTenancyAttachments(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsTenancyAttachmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).TenancyAttachmentClient()

	return tfresource.ReadResource(sync)
}

type ResourceAnalyticsTenancyAttachmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resource_analytics.TenancyAttachmentClient
	Res    *oci_resource_analytics.ListTenancyAttachmentsResponse
}

func (s *ResourceAnalyticsTenancyAttachmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourceAnalyticsTenancyAttachmentsDataSourceCrud) Get() error {
	request := oci_resource_analytics.ListTenancyAttachmentsRequest{}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
		tmp := resourceAnalyticsInstanceId.(string)
		request.ResourceAnalyticsInstanceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_resource_analytics.TenancyAttachmentLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resource_analytics")

	response, err := s.Client.ListTenancyAttachments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTenancyAttachments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ResourceAnalyticsTenancyAttachmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ResourceAnalyticsTenancyAttachmentsDataSource-", ResourceAnalyticsTenancyAttachmentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	tenancyAttachment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TenancyAttachmentSummaryToMap(item))
	}
	tenancyAttachment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ResourceAnalyticsTenancyAttachmentsDataSource().Schema["tenancy_attachment_collection"].Elem.(*schema.Resource).Schema)
		tenancyAttachment["items"] = items
	}

	resources = append(resources, tenancyAttachment)
	if err := s.D.Set("tenancy_attachment_collection", resources); err != nil {
		return err
	}

	return nil
}
