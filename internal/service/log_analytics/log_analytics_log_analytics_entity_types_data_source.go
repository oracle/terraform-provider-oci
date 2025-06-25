// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsLogAnalyticsEntityTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsLogAnalyticsEntityTypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cloud_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_analytics_entity_type_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     LogAnalyticsLogAnalyticsEntityTypeResource(),
						},
					},
				},
			},
		},
	}
}

func readLogAnalyticsLogAnalyticsEntityTypes(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntityTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsLogAnalyticsEntityTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListLogAnalyticsEntityTypesResponse
}

func (s *LogAnalyticsLogAnalyticsEntityTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsEntityTypesDataSourceCrud) Get() error {
	request := oci_log_analytics.ListLogAnalyticsEntityTypesRequest{}

	if cloudType, ok := s.D.GetOkExists("cloud_type"); ok {
		request.CloudType = oci_log_analytics.ListLogAnalyticsEntityTypesCloudTypeEnum(cloudType.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_log_analytics.ListLogAnalyticsEntityTypesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListLogAnalyticsEntityTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLogAnalyticsEntityTypes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsEntityTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsLogAnalyticsEntityTypesDataSource-", LogAnalyticsLogAnalyticsEntityTypesDataSource(), s.D))
	resources := []map[string]interface{}{}
	logAnalyticsEntityType := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LogAnalyticsEntityTypeSummaryToMap(item))
	}
	logAnalyticsEntityType["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsLogAnalyticsEntityTypesDataSource().Schema["log_analytics_entity_type_collection"].Elem.(*schema.Resource).Schema)
		logAnalyticsEntityType["items"] = items
	}

	resources = append(resources, logAnalyticsEntityType)
	if err := s.D.Set("log_analytics_entity_type_collection", resources); err != nil {
		return err
	}

	return nil
}
