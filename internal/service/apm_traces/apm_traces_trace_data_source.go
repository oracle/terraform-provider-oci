// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_traces

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_traces "github.com/oracle/oci-go-sdk/v65/apmtraces"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmTracesTraceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularApmTracesTrace,
		Schema: map[string]*schema.Schema{
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"trace_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"error_span_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_fault": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"root_span_duration_in_ms": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"root_span_operation_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"root_span_service_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_summaries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"error_spans": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"span_service_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_spans": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"span_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"span_summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"error_span_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_fault": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"root_span_duration_in_ms": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"root_span_operation_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"root_span_service_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_summaries": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"error_spans": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"span_service_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_spans": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"span_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"time_earliest_span_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_latest_span_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_root_span_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_root_span_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trace_duration_in_ms": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"trace_error_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trace_error_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trace_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"spans": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"duration_in_ms": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_error": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logs": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"span_logs": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"log_key": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"log_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"operation_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_span_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"tag_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tag_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trace_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_earliest_span_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_latest_span_ended": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_root_span_ended": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_root_span_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trace_duration_in_ms": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trace_error_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trace_error_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trace_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularApmTracesTrace(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesTraceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).TraceClient()

	return tfresource.ReadResource(sync)
}

type ApmTracesTraceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_traces.TraceClient
	Res    *oci_apm_traces.GetTraceResponse
}

func (s *ApmTracesTraceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmTracesTraceDataSourceCrud) Get() error {
	request := oci_apm_traces.GetTraceRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if traceKey, ok := s.D.GetOkExists("trace_key"); ok {
		tmp := traceKey.(string)
		request.TraceKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_traces")

	response, err := s.Client.GetTrace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmTracesTraceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmTracesTraceDataSource-", ApmTracesTraceDataSource(), s.D))

	if s.Res.ErrorSpanCount != nil {
		s.D.Set("error_span_count", *s.Res.ErrorSpanCount)
	}

	if s.Res.IsFault != nil {
		s.D.Set("is_fault", *s.Res.IsFault)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.RootSpanDurationInMs != nil {
		s.D.Set("root_span_duration_in_ms", *s.Res.RootSpanDurationInMs)
	}

	if s.Res.RootSpanOperationName != nil {
		s.D.Set("root_span_operation_name", *s.Res.RootSpanOperationName)
	}

	if s.Res.RootSpanServiceName != nil {
		s.D.Set("root_span_service_name", *s.Res.RootSpanServiceName)
	}

	serviceSummaries := []interface{}{}
	for _, item := range s.Res.ServiceSummaries {
		serviceSummaries = append(serviceSummaries, TraceServiceSummaryToMap(item))
	}
	s.D.Set("service_summaries", serviceSummaries)

	if s.Res.SpanCount != nil {
		s.D.Set("span_count", *s.Res.SpanCount)
	}

	if s.Res.SpanSummary != nil {
		s.D.Set("span_summary", []interface{}{TraceSpanSummaryToMap(s.Res.SpanSummary)})
	} else {
		s.D.Set("span_summary", nil)
	}

	spans := []interface{}{}
	for _, item := range s.Res.Spans {
		spans = append(spans, SpanToMap(item))
	}
	s.D.Set("spans", spans)

	if s.Res.TimeEarliestSpanStarted != nil {
		s.D.Set("time_earliest_span_started", s.Res.TimeEarliestSpanStarted.String())
	}

	if s.Res.TimeLatestSpanEnded != nil {
		s.D.Set("time_latest_span_ended", s.Res.TimeLatestSpanEnded.String())
	}

	if s.Res.TimeRootSpanEnded != nil {
		s.D.Set("time_root_span_ended", s.Res.TimeRootSpanEnded.String())
	}

	if s.Res.TimeRootSpanStarted != nil {
		s.D.Set("time_root_span_started", s.Res.TimeRootSpanStarted.String())
	}

	if s.Res.TraceDurationInMs != nil {
		s.D.Set("trace_duration_in_ms", *s.Res.TraceDurationInMs)
	}

	if s.Res.TraceErrorCode != nil {
		s.D.Set("trace_error_code", *s.Res.TraceErrorCode)
	}

	if s.Res.TraceErrorType != nil {
		s.D.Set("trace_error_type", *s.Res.TraceErrorType)
	}

	if s.Res.TraceStatus != nil {
		s.D.Set("trace_status", *s.Res.TraceStatus)
	}

	return nil
}

func SpanToMap(obj oci_apm_traces.Span) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DurationInMs != nil {
		result["duration_in_ms"] = strconv.FormatInt(*obj.DurationInMs, 10)
	}

	if obj.IsError != nil {
		result["is_error"] = bool(*obj.IsError)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Kind != nil {
		result["kind"] = string(*obj.Kind)
	}

	logs := []interface{}{}
	for _, item := range obj.Logs {
		logs = append(logs, SpanLogCollectionToMap(item))
	}
	result["logs"] = logs

	if obj.OperationName != nil {
		result["operation_name"] = string(*obj.OperationName)
	}

	if obj.ParentSpanKey != nil {
		result["parent_span_key"] = string(*obj.ParentSpanKey)
	}

	if obj.ServiceName != nil {
		result["service_name"] = string(*obj.ServiceName)
	}

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, TagToMap(item))
	}
	result["tags"] = tags

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.TraceKey != nil {
		result["trace_key"] = string(*obj.TraceKey)
	}

	return result
}

func SpanLogToMap(obj oci_apm_traces.SpanLog) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogKey != nil {
		result["log_key"] = string(*obj.LogKey)
	}

	if obj.LogValue != nil {
		result["log_value"] = string(*obj.LogValue)
	}

	return result
}

func SpanLogCollectionToMap(obj oci_apm_traces.SpanLogCollection) map[string]interface{} {
	result := map[string]interface{}{}

	spanLogs := []interface{}{}
	for _, item := range obj.SpanLogs {
		spanLogs = append(spanLogs, SpanLogToMap(item))
	}
	result["span_logs"] = spanLogs

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func TagToMap(obj oci_apm_traces.Tag) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TagName != nil {
		result["tag_name"] = string(*obj.TagName)
	}

	if obj.TagValue != nil {
		result["tag_value"] = string(*obj.TagValue)
	}

	return result
}

func TraceServiceSummaryToMap(obj oci_apm_traces.TraceServiceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ErrorSpans != nil {
		result["error_spans"] = strconv.FormatInt(*obj.ErrorSpans, 10)
	}

	if obj.SpanServiceName != nil {
		result["span_service_name"] = string(*obj.SpanServiceName)
	}

	if obj.TotalSpans != nil {
		result["total_spans"] = strconv.FormatInt(*obj.TotalSpans, 10)
	}

	return result
}

func TraceSpanSummaryToMap(obj *oci_apm_traces.TraceSpanSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ErrorSpanCount != nil {
		result["error_span_count"] = int(*obj.ErrorSpanCount)
	}

	if obj.IsFault != nil {
		result["is_fault"] = bool(*obj.IsFault)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.RootSpanDurationInMs != nil {
		result["root_span_duration_in_ms"] = int(*obj.RootSpanDurationInMs)
	}

	if obj.RootSpanOperationName != nil {
		result["root_span_operation_name"] = string(*obj.RootSpanOperationName)
	}

	if obj.RootSpanServiceName != nil {
		result["root_span_service_name"] = string(*obj.RootSpanServiceName)
	}

	serviceSummaries := []interface{}{}
	for _, item := range obj.ServiceSummaries {
		serviceSummaries = append(serviceSummaries, TraceServiceSummaryToMap(item))
	}
	result["service_summaries"] = serviceSummaries

	if obj.SpanCount != nil {
		result["span_count"] = int(*obj.SpanCount)
	}

	if obj.TimeEarliestSpanStarted != nil {
		result["time_earliest_span_started"] = obj.TimeEarliestSpanStarted.String()
	}

	if obj.TimeLatestSpanEnded != nil {
		result["time_latest_span_ended"] = obj.TimeLatestSpanEnded.String()
	}

	if obj.TimeRootSpanEnded != nil {
		result["time_root_span_ended"] = obj.TimeRootSpanEnded.String()
	}

	if obj.TimeRootSpanStarted != nil {
		result["time_root_span_started"] = obj.TimeRootSpanStarted.String()
	}

	if obj.TraceDurationInMs != nil {
		result["trace_duration_in_ms"] = int(*obj.TraceDurationInMs)
	}

	if obj.TraceErrorCode != nil {
		result["trace_error_code"] = string(*obj.TraceErrorCode)
	}

	if obj.TraceErrorType != nil {
		result["trace_error_type"] = string(*obj.TraceErrorType)
	}

	if obj.TraceStatus != nil {
		result["trace_status"] = string(*obj.TraceStatus)
	}

	return result
}
