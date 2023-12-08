// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationUsageCarbonEmissionsQueryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMeteringComputationUsageCarbonEmissionsQuery,
		Read:     readMeteringComputationUsageCarbonEmissionsQuery,
		Update:   updateMeteringComputationUsageCarbonEmissionsQuery,
		Delete:   deleteMeteringComputationUsageCarbonEmissionsQuery,
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
									"tenant_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"compartment_depth": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"date_range_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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
									"time_usage_ended": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"time_usage_started": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"usage_carbon_emissions_query_filter": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"version": {
							Type:     schema.TypeInt,
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

func createMeteringComputationUsageCarbonEmissionsQuery(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageCarbonEmissionsQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.CreateResource(d, sync)
}

func readMeteringComputationUsageCarbonEmissionsQuery(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageCarbonEmissionsQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

func updateMeteringComputationUsageCarbonEmissionsQuery(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageCarbonEmissionsQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMeteringComputationUsageCarbonEmissionsQuery(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageCarbonEmissionsQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MeteringComputationUsageCarbonEmissionsQueryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_metering_computation.UsageapiClient
	Res                    *oci_metering_computation.UsageCarbonEmissionsQuery
	DisableNotFoundRetries bool
}

func (s *MeteringComputationUsageCarbonEmissionsQueryResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MeteringComputationUsageCarbonEmissionsQueryResourceCrud) Create() error {
	request := oci_metering_computation.CreateUsageCarbonEmissionsQueryRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if queryDefinition, ok := s.D.GetOkExists("query_definition"); ok {
		if tmpList := queryDefinition.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "query_definition", 0)
			tmp, err := s.mapToUsageCarbonEmissionsQueryDefinition(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.QueryDefinition = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.CreateUsageCarbonEmissionsQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UsageCarbonEmissionsQuery
	return nil
}

func (s *MeteringComputationUsageCarbonEmissionsQueryResourceCrud) Get() error {
	request := oci_metering_computation.GetUsageCarbonEmissionsQueryRequest{}

	tmp := s.D.Id()
	request.UsageCarbonEmissionsQueryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.GetUsageCarbonEmissionsQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UsageCarbonEmissionsQuery
	return nil
}

func (s *MeteringComputationUsageCarbonEmissionsQueryResourceCrud) Update() error {
	request := oci_metering_computation.UpdateUsageCarbonEmissionsQueryRequest{}

	if queryDefinition, ok := s.D.GetOkExists("query_definition"); ok {
		if tmpList := queryDefinition.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "query_definition", 0)
			tmp, err := s.mapToUsageCarbonEmissionsQueryDefinition(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.QueryDefinition = &tmp
		}
	}

	tmp := s.D.Id()
	request.UsageCarbonEmissionsQueryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.UpdateUsageCarbonEmissionsQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UsageCarbonEmissionsQuery
	return nil
}

func (s *MeteringComputationUsageCarbonEmissionsQueryResourceCrud) Delete() error {
	request := oci_metering_computation.DeleteUsageCarbonEmissionsQueryRequest{}

	tmp := s.D.Id()
	request.UsageCarbonEmissionsQueryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	_, err := s.Client.DeleteUsageCarbonEmissionsQuery(context.Background(), request)
	return err
}

func (s *MeteringComputationUsageCarbonEmissionsQueryResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.QueryDefinition != nil {
		s.D.Set("query_definition", []interface{}{UsageCarbonEmissionsQueryDefinitionToMap(s.Res.QueryDefinition)})
	} else {
		s.D.Set("query_definition", nil)
	}

	return nil
}

func (s *MeteringComputationUsageCarbonEmissionsQueryResourceCrud) mapToCostAnalysisUI(fieldKeyFormat string) (oci_metering_computation.CostAnalysisUi, error) {
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

func (s *MeteringComputationUsageCarbonEmissionsQueryResourceCrud) mapToTag(fieldKeyFormat string) (oci_metering_computation.Tag, error) {
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

func (s *MeteringComputationUsageCarbonEmissionsQueryResourceCrud) mapToUsageCarbonEmissionsQueryDefinition(fieldKeyFormat string) (oci_metering_computation.UsageCarbonEmissionsQueryDefinition, error) {
	result := oci_metering_computation.UsageCarbonEmissionsQueryDefinition{}

	if costAnalysisUI, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cost_analysis_ui")); ok {
		if tmpList := costAnalysisUI.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "cost_analysis_ui"), 0)
			tmp, err := s.mapToCostAnalysisUI(fieldKeyFormatNextLevel)
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
			tmp, err := s.mapToUsageCarbonEmissionsReportQuery(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert report_query, encountered error: %v", err)
			}
			result.ReportQuery = &tmp
		}
	}

	if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
		tmp := version.(int)
		result.Version = &tmp
	}

	return result, nil
}

func UsageCarbonEmissionsQueryDefinitionToMap(obj *oci_metering_computation.UsageCarbonEmissionsQueryDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CostAnalysisUI != nil {
		result["cost_analysis_ui"] = []interface{}{CostAnalysisUIToMap(obj.CostAnalysisUI)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ReportQuery != nil {
		result["report_query"] = []interface{}{UsageCarbonEmissionsReportQueryToMap(obj.ReportQuery)}
	}

	if obj.Version != nil {
		result["version"] = int(*obj.Version)
	}

	return result
}

func UsageCarbonEmissionsQuerySummaryToMap(obj oci_metering_computation.UsageCarbonEmissionsQuerySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.QueryDefinition != nil {
		result["query_definition"] = []interface{}{UsageCarbonEmissionsQueryDefinitionToMap(obj.QueryDefinition)}
	}

	return result
}

func (s *MeteringComputationUsageCarbonEmissionsQueryResourceCrud) mapToUsageCarbonEmissionsReportQuery(fieldKeyFormat string) (oci_metering_computation.UsageCarbonEmissionsReportQuery, error) {
	result := oci_metering_computation.UsageCarbonEmissionsReportQuery{}

	if compartmentDepth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_depth")); ok {
		tmp := compartmentDepth.(int)
		result.CompartmentDepth = &tmp
	}

	if dateRangeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "date_range_name")); ok {
		result.DateRangeName = oci_metering_computation.UsageCarbonEmissionsReportQueryDateRangeNameEnum(dateRangeName.(string))
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
			converted, err := s.mapToTag(fieldKeyFormatNextLevel)
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

	if usageCarbonEmissionsQueryFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "usage_carbon_emissions_query_filter")); ok {
		tmp := usageCarbonEmissionsQueryFilter.(string)
		var usageCarbonEmissionsQuery_filterObj oci_metering_computation.Filter
		err := json.Unmarshal([]byte(tmp), &usageCarbonEmissionsQuery_filterObj)
		if err != nil {
			return result, err
		}
		result.Filter = &usageCarbonEmissionsQuery_filterObj
	}

	return result, nil
}

func UsageCarbonEmissionsReportQueryToMap(obj *oci_metering_computation.UsageCarbonEmissionsReportQuery) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentDepth != nil {
		result["compartment_depth"] = int(*obj.CompartmentDepth)
	}

	result["date_range_name"] = string(obj.DateRangeName)

	result["group_by"] = obj.GroupBy

	groupByTag := []interface{}{}
	for _, item := range obj.GroupByTag {
		groupByTag = append(groupByTag, TagToMap(item))
	}
	result["group_by_tag"] = groupByTag

	if obj.IsAggregateByTime != nil {
		result["is_aggregate_by_time"] = bool(*obj.IsAggregateByTime)
	}

	if obj.TenantId != nil {
		result["tenant_id"] = string(*obj.TenantId)
	}

	if obj.TimeUsageEnded != nil {
		result["time_usage_ended"] = obj.TimeUsageEnded.Format(time.RFC3339Nano)
	}

	if obj.TimeUsageStarted != nil {
		result["time_usage_started"] = obj.TimeUsageStarted.Format(time.RFC3339Nano)
	}

	if obj.Filter != nil {
		tmp, _ := json.Marshal(obj.Filter)
		result["usage_carbon_emissions_query_filter"] = string(tmp)
	}

	return result
}
