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

func ResourceAnalyticsMonitoredRegionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readResourceAnalyticsMonitoredRegions,
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
			"monitored_region_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ResourceAnalyticsMonitoredRegionResource()),
						},
					},
				},
			},
		},
	}
}

func readResourceAnalyticsMonitoredRegions(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsMonitoredRegionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MonitoredRegionClient()

	return tfresource.ReadResource(sync)
}

type ResourceAnalyticsMonitoredRegionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resource_analytics.MonitoredRegionClient
	Res    *oci_resource_analytics.ListMonitoredRegionsResponse
}

func (s *ResourceAnalyticsMonitoredRegionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourceAnalyticsMonitoredRegionsDataSourceCrud) Get() error {
	request := oci_resource_analytics.ListMonitoredRegionsRequest{}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
		tmp := resourceAnalyticsInstanceId.(string)
		request.ResourceAnalyticsInstanceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_resource_analytics.MonitoredRegionLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resource_analytics")

	response, err := s.Client.ListMonitoredRegions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMonitoredRegions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ResourceAnalyticsMonitoredRegionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ResourceAnalyticsMonitoredRegionsDataSource-", ResourceAnalyticsMonitoredRegionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	monitoredRegion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MonitoredRegionSummaryToMap(item))
	}
	monitoredRegion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ResourceAnalyticsMonitoredRegionsDataSource().Schema["monitored_region_collection"].Elem.(*schema.Resource).Schema)
		monitoredRegion["items"] = items
	}

	resources = append(resources, monitoredRegion)
	if err := s.D.Set("monitored_region_collection", resources); err != nil {
		return err
	}

	return nil
}
