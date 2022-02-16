// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v58/loganalytics"
)

func LogAnalyticsLogAnalyticsCategoriesListDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsLogAnalyticsCategoriesList,
		Schema: map[string]*schema.Schema{
			"category_display_text": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"category_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
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
						"is_system": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularLogAnalyticsLogAnalyticsCategoriesList(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsCategoriesListDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsLogAnalyticsCategoriesListDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListCategoriesResponse
}

func (s *LogAnalyticsLogAnalyticsCategoriesListDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsCategoriesListDataSourceCrud) Get() error {
	request := oci_log_analytics.ListCategoriesRequest{}

	if categoryDisplayText, ok := s.D.GetOkExists("category_display_text"); ok {
		tmp := categoryDisplayText.(string)
		request.CategoryDisplayText = &tmp
	}

	if categoryType, ok := s.D.GetOkExists("category_type"); ok {
		tmp := categoryType.(string)
		request.CategoryType = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListCategories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsLogAnalyticsCategoriesListDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsLogAnalyticsCategoriesListDataSource-", LogAnalyticsLogAnalyticsCategoriesListDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LogAnalyticsCategoryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func LogAnalyticsCategoryToMap(obj oci_log_analytics.LogAnalyticsCategory) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.IsSystem != nil {
		result["is_system"] = bool(*obj.IsSystem)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}
