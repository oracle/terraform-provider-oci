// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v58/usageapi"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func MeteringComputationQueryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMeteringComputationQuery,
		Read:     readMeteringComputationQuery,
		Update:   updateMeteringComputationQuery,
		Delete:   deleteMeteringComputationQuery,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"query_definition": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"cost_analysis_ui": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"graph": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"is_cumulative_graph": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"display_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"report_query": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"granularity": {
										Type:     schema.TypeString,
										Required: true,
									},
									"tenant_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"compartment_depth": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"date_range_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"filter": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"forecast": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"time_forecast_ended": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: utils.TimeDiffSuppressFunction,
												},

												// Optional
												"forecast_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"time_forecast_started": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													DiffSuppressFunc: utils.TimeDiffSuppressFunction,
												},

												// Computed
											},
										},
									},
									"group_by": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"group_by_tag": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"namespace": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"is_aggregate_by_time": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"query_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"time_usage_ended": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: utils.TimeDiffSuppressFunction,
									},
									"time_usage_started": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: utils.TimeDiffSuppressFunction,
									},

									// Computed
								},
							},
						},
						"version": {
							Type:     schema.TypeFloat,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional

			// Computed
		},
	}
}

func createMeteringComputationQuery(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.CreateResource(d, sync)
}

func readMeteringComputationQuery(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

func updateMeteringComputationQuery(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMeteringComputationQuery(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MeteringComputationQueryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_metering_computation.UsageapiClient
	Res                    *oci_metering_computation.Query
	DisableNotFoundRetries bool
}

func (s *MeteringComputationQueryResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MeteringComputationQueryResourceCrud) Create() error {
	request := oci_metering_computation.CreateQueryRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if queryDefinition, ok := s.D.GetOkExists("query_definition"); ok {
		if tmpList := queryDefinition.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "query_definition", 0)
			tmp, err := s.mapToQueryDefinition(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.QueryDefinition = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.CreateQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Query
	return nil
}

func (s *MeteringComputationQueryResourceCrud) Get() error {
	request := oci_metering_computation.GetQueryRequest{}

	tmp := s.D.Id()
	request.QueryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.GetQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Query
	return nil
}

func (s *MeteringComputationQueryResourceCrud) Update() error {
	request := oci_metering_computation.UpdateQueryRequest{}

	if queryDefinition, ok := s.D.GetOkExists("query_definition"); ok {
		if tmpList := queryDefinition.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "query_definition", 0)
			tmp, err := s.mapToQueryDefinition(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.QueryDefinition = &tmp
		}
	}

	tmp := s.D.Id()
	request.QueryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.UpdateQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Query
	return nil
}

func (s *MeteringComputationQueryResourceCrud) Delete() error {
	request := oci_metering_computation.DeleteQueryRequest{}

	tmp := s.D.Id()
	request.QueryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	_, err := s.Client.DeleteQuery(context.Background(), request)
	return err
}

func (s *MeteringComputationQueryResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.QueryDefinition != nil {
		s.D.Set("query_definition", []interface{}{QueryDefinitionToMap(s.Res.QueryDefinition)})
	} else {
		s.D.Set("query_definition", nil)
	}

	return nil
}

func (s *MeteringComputationQueryResourceCrud) mapToCostAnalysisUi(fieldKeyFormat string) (oci_metering_computation.CostAnalysisUi, error) {
	result := oci_metering_computation.CostAnalysisUi{}

	if graph, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "graph")); ok {
		result.Graph = oci_metering_computation.CostAnalysisUiGraphEnum(graph.(string))
	}

	if isCumulativeGraph, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_cumulative_graph")); ok {
		tmp := isCumulativeGraph.(bool)
		result.IsCumulativeGraph = &tmp
	}

	return result, nil
}

func CostAnalysisUIToMap(obj *oci_metering_computation.CostAnalysisUi) map[string]interface{} {
	result := map[string]interface{}{}

	result["graph"] = string(obj.Graph)

	if obj.IsCumulativeGraph != nil {
		result["is_cumulative_graph"] = bool(*obj.IsCumulativeGraph)
	}

	return result
}

func (s *MeteringComputationQueryResourceCrud) mapToForecast(fieldKeyFormat string) (oci_metering_computation.Forecast, error) {
	result := oci_metering_computation.Forecast{}

	if forecastType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "forecast_type")); ok {
		result.ForecastType = oci_metering_computation.ForecastForecastTypeEnum(forecastType.(string))
	}

	if timeForecastEnded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_forecast_ended")); ok {
		tmp, err := time.Parse(time.RFC3339, timeForecastEnded.(string))
		if err != nil {
			return result, err
		}
		result.TimeForecastEnded = &oci_common.SDKTime{Time: tmp}
	}

	if timeForecastStarted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_forecast_started")); ok {
		tmp, err := time.Parse(time.RFC3339, timeForecastStarted.(string))
		if err != nil {
			return result, err
		}
		result.TimeForecastStarted = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func ForecastToMap(obj *oci_metering_computation.Forecast) map[string]interface{} {
	result := map[string]interface{}{}

	result["forecast_type"] = string(obj.ForecastType)

	if obj.TimeForecastEnded != nil {
		result["time_forecast_ended"] = obj.TimeForecastEnded.Format(time.RFC3339Nano)
	}

	if obj.TimeForecastStarted != nil {
		result["time_forecast_started"] = obj.TimeForecastStarted.Format(time.RFC3339Nano)
	}

	return result
}

func (s *MeteringComputationQueryResourceCrud) mapToQueryDefinition(fieldKeyFormat string) (oci_metering_computation.QueryDefinition, error) {
	result := oci_metering_computation.QueryDefinition{}

	if costAnalysisUI, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cost_analysis_ui")); ok {
		if tmpList := costAnalysisUI.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "cost_analysis_ui"), 0)
			tmp, err := s.mapToCostAnalysisUi(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert cost_analysis_ui, encountered error: %v", err)
			}
			result.CostAnalysisUI = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if reportQuery, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "report_query")); ok {
		if tmpList := reportQuery.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "report_query"), 0)
			tmp, err := s.mapToReportQuery(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert report_query, encountered error: %v", err)
			}
			result.ReportQuery = &tmp
		}
	}

	if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
		tmp := float32(version.(float64))
		result.Version = &tmp
	}

	return result, nil
}

func QueryDefinitionToMap(obj *oci_metering_computation.QueryDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CostAnalysisUI != nil {
		result["cost_analysis_ui"] = []interface{}{CostAnalysisUIToMap(obj.CostAnalysisUI)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ReportQuery != nil {
		result["report_query"] = []interface{}{ReportQueryToMap(obj.ReportQuery)}
	}

	if obj.Version != nil {
		result["version"] = float32(*obj.Version)
	}

	return result
}

func QuerySummaryToMap(obj oci_metering_computation.QuerySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.QueryDefinition != nil {
		result["query_definition"] = []interface{}{QueryDefinitionToMap(obj.QueryDefinition)}
	}

	return result
}

func (s *MeteringComputationQueryResourceCrud) mapToReportQuery(fieldKeyFormat string) (oci_metering_computation.ReportQuery, error) {
	result := oci_metering_computation.ReportQuery{}

	if compartmentDepth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_depth")); ok {
		tmp := float32(compartmentDepth.(float64))
		result.CompartmentDepth = &tmp
	}

	if dateRangeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "date_range_name")); ok {
		result.DateRangeName = oci_metering_computation.ReportQueryDateRangeNameEnum(dateRangeName.(string))
	}

	if filter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filter")); ok {
		tmp := filter.(string)
		//follow the filter in usage_resource.go
		var filterObj oci_metering_computation.Filter
		err := json.Unmarshal([]byte(tmp), &filterObj)
		if err != nil {
			return result, err
		}
		result.Filter = &filterObj
	}

	if forecast, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "forecast")); ok {
		if tmpList := forecast.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "forecast"), 0)
			tmp, err := s.mapToForecast(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert forecast, encountered error: %v", err)
			}
			result.Forecast = &tmp
		}
	}

	if granularity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "granularity")); ok {
		result.Granularity = oci_metering_computation.ReportQueryGranularityEnum(granularity.(string))
	}

	if groupBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_by")); ok {
		interfaces := groupBy.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "group_by")) {
			result.GroupBy = tmp
		}
	}

	if groupByTag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_by_tag")); ok {
		interfaces := groupByTag.([]interface{})
		tmp := make([]oci_metering_computation.Tag, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "group_by_tag"), stateDataIndex)
			converted, err := s.mapToTagInQuery(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "group_by_tag")) {
			result.GroupByTag = tmp
		}
	}

	if isAggregateByTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_aggregate_by_time")); ok {
		tmp := isAggregateByTime.(bool)
		result.IsAggregateByTime = &tmp
	}

	if queryType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query_type")); ok {
		result.QueryType = oci_metering_computation.ReportQueryQueryTypeEnum(queryType.(string))
	}

	if tenantId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tenant_id")); ok {
		tmp := tenantId.(string)
		result.TenantId = &tmp
	}

	if timeUsageEnded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_usage_ended")); ok {
		tmp, err := time.Parse(time.RFC3339, timeUsageEnded.(string))
		if err != nil {
			return result, err
		}
		result.TimeUsageEnded = &oci_common.SDKTime{Time: tmp}
	}

	if timeUsageStarted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_usage_started")); ok {
		tmp, err := time.Parse(time.RFC3339, timeUsageStarted.(string))
		if err != nil {
			return result, err
		}
		result.TimeUsageStarted = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func ReportQueryToMap(obj *oci_metering_computation.ReportQuery) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentDepth != nil {
		result["compartment_depth"] = float32(*obj.CompartmentDepth)
	}

	result["date_range_name"] = string(obj.DateRangeName)

	if obj.Filter != nil {
		tmp, _ := json.Marshal(obj.Filter)
		result["filter"] = string(tmp)
	}

	if obj.Forecast != nil {
		result["forecast"] = []interface{}{ForecastToMap(obj.Forecast)}
	}

	result["granularity"] = string(obj.Granularity)

	result["group_by"] = obj.GroupBy

	groupByTag := []interface{}{}
	for _, item := range obj.GroupByTag {
		groupByTag = append(groupByTag, TagToMapInQuery(item))
	}
	result["group_by_tag"] = groupByTag

	if obj.IsAggregateByTime != nil {
		result["is_aggregate_by_time"] = bool(*obj.IsAggregateByTime)
	}

	result["query_type"] = string(obj.QueryType)

	if obj.TenantId != nil {
		result["tenant_id"] = string(*obj.TenantId)
	}

	if obj.TimeUsageEnded != nil {
		result["time_usage_ended"] = obj.TimeUsageEnded.Format(time.RFC3339Nano)
	}

	if obj.TimeUsageStarted != nil {
		result["time_usage_started"] = obj.TimeUsageStarted.Format(time.RFC3339Nano)
	}

	return result
}

func (s *MeteringComputationQueryResourceCrud) mapToTagInQuery(fieldKeyFormat string) (oci_metering_computation.Tag, error) {
	result := oci_metering_computation.Tag{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func TagToMapInQuery(obj oci_metering_computation.Tag) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
