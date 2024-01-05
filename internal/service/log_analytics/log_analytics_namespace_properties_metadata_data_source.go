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

func LogAnalyticsNamespacePropertiesMetadataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsNamespacePropertiesMetadata,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"constraints": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_text": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"level": {
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
			"property_metadata_summary_collection": {
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
									"default_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"levels": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"constraints": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
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

func readLogAnalyticsNamespacePropertiesMetadata(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespacePropertiesMetadataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespacePropertiesMetadataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListPropertiesMetadataResponse
}

func (s *LogAnalyticsNamespacePropertiesMetadataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespacePropertiesMetadataDataSourceCrud) Get() error {
	request := oci_log_analytics.ListPropertiesMetadataRequest{}

	if constraints, ok := s.D.GetOkExists("constraints"); ok {
		tmp := constraints.(string)
		request.Constraints = &tmp
	}

	if displayText, ok := s.D.GetOkExists("display_text"); ok {
		tmp := displayText.(string)
		request.DisplayText = &tmp
	}

	if level, ok := s.D.GetOkExists("level"); ok {
		tmp := level.(string)
		request.Level = &tmp
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

	response, err := s.Client.ListPropertiesMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPropertiesMetadata(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LogAnalyticsNamespacePropertiesMetadataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespacePropertiesMetadataDataSource-", LogAnalyticsNamespacePropertiesMetadataDataSource(), s.D))
	resources := []map[string]interface{}{}
	namespacePropertiesMetadata := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PropertyMetadataSummaryToMap(item))
	}
	namespacePropertiesMetadata["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsNamespacePropertiesMetadataDataSource().Schema["property_metadata_summary_collection"].Elem.(*schema.Resource).Schema)
		namespacePropertiesMetadata["items"] = items
	}

	resources = append(resources, namespacePropertiesMetadata)
	if err := s.D.Set("property_metadata_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func LevelToMap(obj oci_log_analytics.Level) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Constraints != nil {
		result["constraints"] = string(*obj.Constraints)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func PropertyMetadataSummaryToMap(obj oci_log_analytics.PropertyMetadataSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	levels := []interface{}{}
	for _, item := range obj.Levels {
		levels = append(levels, LevelToMap(item))
	}
	result["levels"] = levels

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
