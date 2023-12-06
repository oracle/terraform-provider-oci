// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringBaselineableMetricsEvaluateDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularStackMonitoringBaselineableMetricsEvaluate,
		Schema: map[string]*schema.Schema{
			"baselineable_metric_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"items": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"evaluation_data_points": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"timestamp": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"value": {
										Type:     schema.TypeFloat,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"training_data_points": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"timestamp": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"value": {
										Type:     schema.TypeFloat,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional
						"dimensions": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},

						// Computed
						"data_points": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"anomaly": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"high": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"low": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"timestamp": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"data_points": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"anomaly": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"high": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"low": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"timestamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularStackMonitoringBaselineableMetricsEvaluate(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringBaselineableMetricsEvaluateDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringBaselineableMetricsEvaluateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.EvaluateBaselineableMetricResponse
}

func (s *StackMonitoringBaselineableMetricsEvaluateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringBaselineableMetricsEvaluateDataSourceCrud) Get() error {
	request := oci_stack_monitoring.EvaluateBaselineableMetricRequest{}

	if baselineableMetricId, ok := s.D.GetOkExists("baselineable_metric_id"); ok {
		tmp := baselineableMetricId.(string)
		request.BaselineableMetricId = &tmp
	}

	if items, ok := s.D.GetOkExists("items"); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_stack_monitoring.MetricData, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "items", stateDataIndex)
			converted, err := s.mapToMetricData(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("items") {
			request.Items = tmp
		}
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.EvaluateBaselineableMetric(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StackMonitoringBaselineableMetricsEvaluateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StackMonitoringBaselineableMetricsEvaluateDataSource-", StackMonitoringBaselineableMetricsEvaluateDataSource(), s.D))

	return nil
}

func AnomalyDataPointToMap(obj oci_stack_monitoring.AnomalyDataPoint) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Anomaly != nil {
		result["anomaly"] = float64(*obj.Anomaly)
	}

	if obj.High != nil {
		result["high"] = float64(*obj.High)
	}

	if obj.Low != nil {
		result["low"] = float64(*obj.Low)
	}

	if obj.Timestamp != nil {
		result["timestamp"] = obj.Timestamp.String()
	}

	if obj.Value != nil {
		result["value"] = float64(*obj.Value)
	}

	return result
}

func (s *StackMonitoringBaselineableMetricsEvaluateDataSourceCrud) mapToDataPoint(fieldKeyFormat string) (oci_stack_monitoring.DataPoint, error) {
	result := oci_stack_monitoring.DataPoint{}

	if timestamp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timestamp")); ok {
		tmp, err := time.Parse(time.RFC3339, timestamp.(string))
		if err != nil {
			return result, err
		}
		result.Timestamp = &oci_common.SDKTime{Time: tmp}
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(float64)
		result.Value = &tmp
	}

	return result, nil
}

func DataPointToMap(obj oci_stack_monitoring.DataPoint) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Timestamp != nil {
		result["timestamp"] = obj.Timestamp.Format(time.RFC3339Nano)
	}

	if obj.Value != nil {
		result["value"] = float64(*obj.Value)
	}

	return result
}

func (s *StackMonitoringBaselineableMetricsEvaluateDataSourceCrud) mapToMetricData(fieldKeyFormat string) (oci_stack_monitoring.MetricData, error) {
	result := oci_stack_monitoring.MetricData{}

	if dimensions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dimensions")); ok {
		result.Dimensions = tfresource.ObjectMapToStringMap(dimensions.(map[string]interface{}))
	}

	if evaluationDataPoints, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "evaluation_data_points")); ok {
		interfaces := evaluationDataPoints.([]interface{})
		tmp := make([]oci_stack_monitoring.DataPoint, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "evaluation_data_points"), stateDataIndex)
			converted, err := s.mapToDataPoint(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "evaluation_data_points")) {
			result.EvaluationDataPoints = tmp
		}
	}

	if trainingDataPoints, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "training_data_points")); ok {
		interfaces := trainingDataPoints.([]interface{})
		tmp := make([]oci_stack_monitoring.DataPoint, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "training_data_points"), stateDataIndex)
			converted, err := s.mapToDataPoint(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "training_data_points")) {
			result.TrainingDataPoints = tmp
		}
	}

	return result, nil
}

func MetricDataToMap(obj oci_stack_monitoring.MetricData) map[string]interface{} {
	result := map[string]interface{}{}

	result["dimensions"] = obj.Dimensions

	evaluationDataPoints := []interface{}{}
	for _, item := range obj.EvaluationDataPoints {
		evaluationDataPoints = append(evaluationDataPoints, DataPointToMap(item))
	}
	result["evaluation_data_points"] = evaluationDataPoints

	trainingDataPoints := []interface{}{}
	for _, item := range obj.TrainingDataPoints {
		trainingDataPoints = append(trainingDataPoints, DataPointToMap(item))
	}
	result["training_data_points"] = trainingDataPoints

	return result
}
