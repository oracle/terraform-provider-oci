// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSensitiveColumnAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSensitiveColumnAnalytics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"column_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_by": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"object": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"schema_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_type_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_type_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_database_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_column_analytics_collection": {
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
									"dimensions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"column_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"object": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"schema_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sensitive_data_model_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sensitive_type_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"target_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"sensitive_column_analytic_count": {
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

func readDataSafeSensitiveColumnAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveColumnAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveColumnAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSensitiveColumnAnalyticsResponse
}

func (s *DataSafeSensitiveColumnAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveColumnAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListSensitiveColumnAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSensitiveColumnAnalyticsAccessLevelEnum(accessLevel.(string))
	}

	if columnName, ok := s.D.GetOkExists("column_name"); ok {
		interfaces := columnName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("column_name") {
			request.ColumnName = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		interfaces := groupBy.([]interface{})
		tmp := make([]oci_data_safe.ListSensitiveColumnAnalyticsGroupByEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.ListSensitiveColumnAnalyticsGroupByEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_by") {
			request.GroupBy = tmp
		}
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		interfaces := object.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("object") {
			request.ObjectName = tmp
		}
	}

	if schemaName, ok := s.D.GetOkExists("schema_name"); ok {
		interfaces := schemaName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schema_name") {
			request.SchemaName = tmp
		}
	}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	if sensitiveTypeGroupId, ok := s.D.GetOkExists("sensitive_type_group_id"); ok {
		tmp := sensitiveTypeGroupId.(string)
		request.SensitiveTypeGroupId = &tmp
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		interfaces := sensitiveTypeId.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("sensitive_type_id") {
			request.SensitiveTypeId = tmp
		}
	}

	if targetDatabaseGroupId, ok := s.D.GetOkExists("target_database_group_id"); ok {
		tmp := targetDatabaseGroupId.(string)
		request.TargetDatabaseGroupId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSensitiveColumnAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSensitiveColumnAnalytics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSensitiveColumnAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSensitiveColumnAnalyticsDataSource-", DataSafeSensitiveColumnAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	sensitiveColumnAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SensitiveColumnAnalyticsSummaryToMap(item))
	}
	sensitiveColumnAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSensitiveColumnAnalyticsDataSource().Schema["sensitive_column_analytics_collection"].Elem.(*schema.Resource).Schema)
		sensitiveColumnAnalytic["items"] = items
	}

	resources = append(resources, sensitiveColumnAnalytic)
	if err := s.D.Set("sensitive_column_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func SensitiveColumnAnalyticsDimensionsToMap(obj *oci_data_safe.SensitiveColumnAnalyticsDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ColumnName != nil {
		result["column_name"] = string(*obj.ColumnName)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	if obj.SensitiveDataModelId != nil {
		result["sensitive_data_model_id"] = string(*obj.SensitiveDataModelId)
	}

	if obj.SensitiveTypeId != nil {
		result["sensitive_type_id"] = string(*obj.SensitiveTypeId)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	return result
}

func SensitiveColumnAnalyticsSummaryToMap(obj oci_data_safe.SensitiveColumnAnalyticsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{SensitiveColumnAnalyticsDimensionsToMap(obj.Dimensions)}
	}

	if obj.Count != nil {
		result["sensitive_column_analytic_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	return result
}
