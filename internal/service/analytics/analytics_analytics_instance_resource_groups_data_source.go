// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_analytics "github.com/oracle/oci-go-sdk/v65/analytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AnalyticsAnalyticsInstanceResourceGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readAnalyticsAnalyticsInstanceResourceGroupsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"analytics_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_resource_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(AnalyticsAnalyticsInstanceResourceGroupResource()),
			},
		},
	}
}

func readAnalyticsAnalyticsInstanceResourceGroupsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AnalyticsAnalyticsInstanceResourceGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type AnalyticsAnalyticsInstanceResourceGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_analytics.AnalyticsClient
	Res    *oci_analytics.ListResourceGroupsResponse
}

func (s *AnalyticsAnalyticsInstanceResourceGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AnalyticsAnalyticsInstanceResourceGroupsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_analytics.ListResourceGroupsRequest{}

	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "analytics")

	response, err := s.Client.ListResourceGroups(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResourceGroups(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AnalyticsAnalyticsInstanceResourceGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AnalyticsAnalyticsInstanceResourceGroupsDataSource-", AnalyticsAnalyticsInstanceResourceGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	analyticsInstanceId := s.D.Get("analytics_instance_id").(string)

	for _, r := range s.Res.Items {
		analyticsInstanceResourceGroup := InstanceResourceGroupSummaryToMap(r, analyticsInstanceId)

		resources = append(resources, analyticsInstanceResourceGroup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, AnalyticsAnalyticsInstanceResourceGroupsDataSource().Schema["instance_resource_groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instance_resource_groups", resources); err != nil {
		return err
	}

	return nil
}

func InstanceResourceGroupSummaryToMap(obj oci_analytics.InstanceResourceGroupSummary, analyticsInstanceId string) map[string]interface{} {
	result := map[string]interface{}{
		"analytics_instance_id": analyticsInstanceId,
	}

	if obj.Id != nil {
		result["id"] = GetAnalyticsInstanceResourceGroupCompositeId(analyticsInstanceId, *obj.Id)
	}

	if obj.Capacity != nil {
		result["capacity"] = *obj.Capacity
	}

	if obj.Description != nil {
		result["description"] = *obj.Description
	}

	if obj.DisplayName != nil {
		result["display_name"] = *obj.DisplayName
	}

	if obj.ResourceName != nil {
		result["resource_name"] = *obj.ResourceName
	}

	return result
}
