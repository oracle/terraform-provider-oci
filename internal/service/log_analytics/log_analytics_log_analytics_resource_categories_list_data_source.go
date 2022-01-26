// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v56/loganalytics"
)

func LogAnalyticsLogAnalyticsResourceCategoriesListDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsLogAnalyticsResourceCategoriesList,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_ids": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_types": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_categories": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"categories": {
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
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"category_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_system": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularLogAnalyticsLogAnalyticsResourceCategoriesList(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsResourceCategoriesListDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsLogAnalyticsResourceCategoriesListDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListResourceCategoriesResponse
}

func (s *LogAnalyticsLogAnalyticsResourceCategoriesListDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsResourceCategoriesListDataSourceCrud) Get() error {
	request := oci_log_analytics.ListResourceCategoriesRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if resourceIds, ok := s.D.GetOkExists("resource_ids"); ok {
		tmp := resourceIds.(string)
		request.ResourceIds = &tmp
	}

	if resourceTypes, ok := s.D.GetOkExists("resource_types"); ok {
		tmp := resourceTypes.(string)
		request.ResourceTypes = &tmp
	}

	if resourceCategories, ok := s.D.GetOkExists("resource_categories"); ok {
		tmp := resourceCategories.(string)
		request.Categories = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListResourceCategories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsLogAnalyticsResourceCategoriesListDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsLogAnalyticsResourceCategoriesListDataSource-", LogAnalyticsLogAnalyticsResourceCategoriesListDataSource(), s.D))

	categories := []interface{}{}
	for _, item := range s.Res.Categories {
		categories = append(categories, LogAnalyticsCategoryToMap(item))
	}
	s.D.Set("categories", categories)

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LogAnalyticsResourceCategoryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func LogAnalyticsResourceCategoryToMap(obj oci_log_analytics.LogAnalyticsResourceCategory) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CategoryName != nil {
		result["category_name"] = string(*obj.CategoryName)
	}

	if obj.IsSystem != nil {
		result["is_system"] = bool(*obj.IsSystem)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	return result
}
