// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsqlInsightCapabilitiesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readPsqlInsightCapabilitiesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"insight_capability_collection": {
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
									"data_type_capabilities": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"data_contract": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"kind": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"unit": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"date_time_range_support": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"is_required": {
																Type:     schema.TypeBool,
																Computed: true,
															},
														},
													},
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"filters": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"can_use_partial_match": {
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
															"values": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"granularity": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"max_seconds": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"min_seconds": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"type": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"insight_data_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"limits": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"max_rows": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"max_time_range_days": {
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"pagination": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"default_limit": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"is_supported": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"max_limit": {
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"sorting": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"default_sort": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"field": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"order": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"fields": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"is_supported": {
																Type:     schema.TypeBool,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"insight_type": {
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

func readPsqlInsightCapabilitiesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &PsqlInsightCapabilitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type PsqlInsightCapabilitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psql.PostgresqlClient
	Res    *oci_psql.ListInsightCapabilitiesResponse
}

func (s *PsqlInsightCapabilitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsqlInsightCapabilitiesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_psql.ListInsightCapabilitiesRequest{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psql")

	response, err := s.Client.ListInsightCapabilities(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInsightCapabilities(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *PsqlInsightCapabilitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("PsqlInsightCapabilitiesDataSource-", PsqlInsightCapabilitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	insightCapability := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InsightCapabilitySummaryToMap(item))
	}
	insightCapability["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, PsqlInsightCapabilitiesDataSource().Schema["insight_capability_collection"].Elem.(*schema.Resource).Schema)
		insightCapability["items"] = items
	}

	resources = append(resources, insightCapability)
	if err := s.D.Set("insight_capability_collection", resources); err != nil {
		return err
	}

	return nil
}

func DateTimeRangeCapabilityToMap(obj *oci_psql.DateTimeRangeCapability) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsRequired != nil {
		result["is_required"] = bool(*obj.IsRequired)
	}

	return result
}

func GranularityCapabilityToMap(obj *oci_psql.GranularityCapability) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxSeconds != nil {
		result["max_seconds"] = int(*obj.MaxSeconds)
	}

	if obj.MinSeconds != nil {
		result["min_seconds"] = int(*obj.MinSeconds)
	}

	result["type"] = string(obj.Type)

	return result
}

func InsightCapabilitySummaryToMap(obj oci_psql.InsightCapabilitySummary) map[string]interface{} {
	result := map[string]interface{}{}

	dataTypeCapabilities := []interface{}{}
	for _, item := range obj.DataTypeCapabilities {
		dataTypeCapabilities = append(dataTypeCapabilities, InsightDataTypeCapabilityToMap(item))
	}
	result["data_type_capabilities"] = dataTypeCapabilities

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["insight_type"] = string(obj.InsightType)

	return result
}

func InsightDataContractToMap(obj *oci_psql.InsightDataContract) map[string]interface{} {
	result := map[string]interface{}{}

	result["kind"] = string(obj.Kind)

	if obj.Unit != nil {
		result["unit"] = string(*obj.Unit)
	}

	return result
}

func InsightDataTypeCapabilityToMap(obj oci_psql.InsightDataTypeCapability) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DataContract != nil {
		result["data_contract"] = []interface{}{InsightDataContractToMap(obj.DataContract)}
	}

	if obj.DateTimeRangeSupport != nil {
		result["date_time_range_support"] = []interface{}{DateTimeRangeCapabilityToMap(obj.DateTimeRangeSupport)}
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	filters := []interface{}{}
	for _, item := range obj.Filters {
		filters = append(filters, InsightFilterCapabilityToMap(item))
	}
	result["filters"] = filters

	if obj.Granularity != nil {
		result["granularity"] = []interface{}{GranularityCapabilityToMap(obj.Granularity)}
	}

	result["insight_data_type"] = string(obj.InsightDataType)

	if obj.Limits != nil {
		result["limits"] = []interface{}{InsightLimitsToMap(obj.Limits)}
	}

	if obj.Pagination != nil {
		result["pagination"] = []interface{}{PaginationCapabilityToMap(obj.Pagination)}
	}

	if obj.Sorting != nil {
		result["sorting"] = []interface{}{SortingCapabilityToMap(obj.Sorting)}
	}

	return result
}

func InsightFilterCapabilityToMap(obj oci_psql.InsightFilterCapability) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CanUsePartialMatch != nil {
		result["can_use_partial_match"] = bool(*obj.CanUsePartialMatch)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	result["values"] = obj.Values

	return result
}

func InsightLimitsToMap(obj *oci_psql.InsightLimits) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxRows != nil {
		result["max_rows"] = int(*obj.MaxRows)
	}

	if obj.MaxTimeRangeDays != nil {
		result["max_time_range_days"] = int(*obj.MaxTimeRangeDays)
	}

	return result
}

func PaginationCapabilityToMap(obj *oci_psql.PaginationCapability) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultLimit != nil {
		result["default_limit"] = int(*obj.DefaultLimit)
	}

	if obj.IsSupported != nil {
		result["is_supported"] = bool(*obj.IsSupported)
	}

	if obj.MaxLimit != nil {
		result["max_limit"] = int(*obj.MaxLimit)
	}

	return result
}

func SortingCapabilityToMap(obj *oci_psql.SortingCapability) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultSort != nil {
		result["default_sort"] = []interface{}{SortingDefaultToMap(obj.DefaultSort)}
	}

	result["fields"] = obj.Fields

	if obj.IsSupported != nil {
		result["is_supported"] = bool(*obj.IsSupported)
	}

	return result
}

func SortingDefaultToMap(obj *oci_psql.SortingDefault) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Field != nil {
		result["field"] = string(*obj.Field)
	}

	result["order"] = string(obj.Order)

	return result
}
