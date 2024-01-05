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

func LogAnalyticsNamespaceEffectivePropertiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsNamespaceEffectiveProperties,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"agent_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_include_patterns": {
				Type:     schema.TypeBool,
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
			"pattern_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"effective_property_collection": {
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
									"effective_level": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"patterns": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"effective_level": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"value": {
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

func readLogAnalyticsNamespaceEffectiveProperties(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceEffectivePropertiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceEffectivePropertiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListEffectivePropertiesResponse
}

func (s *LogAnalyticsNamespaceEffectivePropertiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceEffectivePropertiesDataSourceCrud) Get() error {
	request := oci_log_analytics.ListEffectivePropertiesRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if entityId, ok := s.D.GetOkExists("entity_id"); ok {
		tmp := entityId.(string)
		request.EntityId = &tmp
	}

	if isIncludePatterns, ok := s.D.GetOkExists("is_include_patterns"); ok {
		tmp := isIncludePatterns.(bool)
		request.IsIncludePatterns = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if patternId, ok := s.D.GetOkExists("pattern_id"); ok {
		tmp := patternId.(int)
		request.PatternId = &tmp
	}

	if sourceName, ok := s.D.GetOkExists("source_name"); ok {
		tmp := sourceName.(string)
		request.SourceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListEffectiveProperties(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEffectiveProperties(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LogAnalyticsNamespaceEffectivePropertiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceEffectivePropertiesDataSource-", LogAnalyticsNamespaceEffectivePropertiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	namespaceEffectiveProperty := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EffectivePropertySummaryToMap(item))
	}
	namespaceEffectiveProperty["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsNamespaceEffectivePropertiesDataSource().Schema["effective_property_collection"].Elem.(*schema.Resource).Schema)
		namespaceEffectiveProperty["items"] = items
	}

	resources = append(resources, namespaceEffectiveProperty)
	if err := s.D.Set("effective_property_collection", resources); err != nil {
		return err
	}

	return nil
}

func EffectivePropertySummaryToMap(obj oci_log_analytics.EffectivePropertySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EffectiveLevel != nil {
		result["effective_level"] = string(*obj.EffectiveLevel)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	patterns := []interface{}{}
	for _, item := range obj.Patterns {
		patterns = append(patterns, PatternOverrideToMap(item))
	}
	result["patterns"] = patterns

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func PatternOverrideToMap(obj oci_log_analytics.PatternOverride) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EffectiveLevel != nil {
		result["effective_level"] = string(*obj.EffectiveLevel)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
