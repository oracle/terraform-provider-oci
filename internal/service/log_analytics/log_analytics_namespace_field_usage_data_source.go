// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceFieldUsageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsNamespaceFieldUsage,
		Schema: map[string]*schema.Schema{
			"field_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"dependent_parsers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"dependencies": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"reference_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reference_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reference_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reference_type": {
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
						"is_system": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"parser_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parser_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parser_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parser_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dependent_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"dependencies": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"reference_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reference_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reference_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reference_type": {
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
						"entity_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"entity_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_type_category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_type_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"is_auto_association_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_system": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"source_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularLogAnalyticsNamespaceFieldUsage(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceFieldUsageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceFieldUsageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetFieldUsagesResponse
}

func (s *LogAnalyticsNamespaceFieldUsageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceFieldUsageDataSourceCrud) Get() error {
	request := oci_log_analytics.GetFieldUsagesRequest{}

	if fieldName, ok := s.D.GetOkExists("field_name"); ok {
		tmp := fieldName.(string)
		request.FieldName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetFieldUsages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespaceFieldUsageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceFieldUsageDataSource-", LogAnalyticsNamespaceFieldUsageDataSource(), s.D))

	dependentParsers := []interface{}{}
	for _, item := range s.Res.DependentParsers {
		dependentParsers = append(dependentParsers, DependentParserToMap(item))
	}
	s.D.Set("dependent_parsers", dependentParsers)

	dependentSources := []interface{}{}
	for _, item := range s.Res.DependentSources {
		dependentSources = append(dependentSources, DependentSourceToMap(item))
	}
	s.D.Set("dependent_sources", dependentSources)

	return nil
}

func DependencyToMap(obj oci_log_analytics.Dependency) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ReferenceDisplayName != nil {
		result["reference_display_name"] = string(*obj.ReferenceDisplayName)
	}

	if obj.ReferenceId != nil {
		result["reference_id"] = strconv.FormatInt(*obj.ReferenceId, 10)
	}

	if obj.ReferenceName != nil {
		result["reference_name"] = string(*obj.ReferenceName)
	}

	if obj.ReferenceType != nil {
		result["reference_type"] = string(*obj.ReferenceType)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func DependentParserToMap(obj oci_log_analytics.DependentParser) map[string]interface{} {
	result := map[string]interface{}{}

	dependencies := []interface{}{}
	for _, item := range obj.Dependencies {
		dependencies = append(dependencies, DependencyToMap(item))
	}
	result["dependencies"] = dependencies

	if obj.IsSystem != nil {
		result["is_system"] = bool(*obj.IsSystem)
	}

	if obj.ParserDisplayName != nil {
		result["parser_display_name"] = string(*obj.ParserDisplayName)
	}

	if obj.ParserId != nil {
		result["parser_id"] = strconv.FormatInt(*obj.ParserId, 10)
	}

	if obj.ParserName != nil {
		result["parser_name"] = string(*obj.ParserName)
	}

	result["parser_type"] = string(obj.ParserType)

	return result
}

func DependentSourceToMap(obj oci_log_analytics.DependentSource) map[string]interface{} {
	result := map[string]interface{}{}

	dependencies := []interface{}{}
	for _, item := range obj.Dependencies {
		dependencies = append(dependencies, DependencyToMap(item))
	}
	result["dependencies"] = dependencies

	entityTypes := []interface{}{}
	for _, item := range obj.EntityTypes {
		entityTypes = append(entityTypes, LogAnalyticsSourceEntityTypeToMap(item))
	}
	result["entity_types"] = entityTypes

	if obj.IsAutoAssociationEnabled != nil {
		result["is_auto_association_enabled"] = bool(*obj.IsAutoAssociationEnabled)
	}

	if obj.IsSystem != nil {
		result["is_system"] = bool(*obj.IsSystem)
	}

	if obj.SourceDisplayName != nil {
		result["source_display_name"] = string(*obj.SourceDisplayName)
	}

	if obj.SourceId != nil {
		result["source_id"] = strconv.FormatInt(*obj.SourceId, 10)
	}

	if obj.SourceName != nil {
		result["source_name"] = string(*obj.SourceName)
	}

	if obj.SourceType != nil {
		result["source_type"] = string(*obj.SourceType)
	}

	return result
}

func LogAnalyticsSourceEntityTypeToMap(obj oci_log_analytics.LogAnalyticsSourceEntityType) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EntityType != nil {
		result["entity_type"] = string(*obj.EntityType)
	}

	if obj.EntityTypeCategory != nil {
		result["entity_type_category"] = string(*obj.EntityTypeCategory)
	}

	if obj.EntityTypeDisplayName != nil {
		result["entity_type_display_name"] = string(*obj.EntityTypeDisplayName)
	}

	if obj.SourceId != nil {
		result["source_id"] = strconv.FormatInt(*obj.SourceId, 10)
	}

	return result
}
