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

func ResourceAnalyticsMonitoredRegionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["monitored_region_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ResourceAnalyticsMonitoredRegionResource(), fieldMap, readSingularResourceAnalyticsMonitoredRegion)
}

func readSingularResourceAnalyticsMonitoredRegion(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsMonitoredRegionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MonitoredRegionClient()

	return tfresource.ReadResource(sync)
}

type ResourceAnalyticsMonitoredRegionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resource_analytics.MonitoredRegionClient
	Res    *oci_resource_analytics.GetMonitoredRegionResponse
}

func (s *ResourceAnalyticsMonitoredRegionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourceAnalyticsMonitoredRegionDataSourceCrud) Get() error {
	request := oci_resource_analytics.GetMonitoredRegionRequest{}

	if monitoredRegionId, ok := s.D.GetOkExists("monitored_region_id"); ok {
		tmp := monitoredRegionId.(string)
		request.MonitoredRegionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resource_analytics")

	response, err := s.Client.GetMonitoredRegion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ResourceAnalyticsMonitoredRegionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RegionId != nil {
		s.D.Set("region_id", *s.Res.RegionId)
	}

	if s.Res.ResourceAnalyticsInstanceId != nil {
		s.D.Set("resource_analytics_instance_id", *s.Res.ResourceAnalyticsInstanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
