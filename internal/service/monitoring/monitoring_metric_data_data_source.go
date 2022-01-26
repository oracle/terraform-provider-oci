// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package monitoring

import (
	"context"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_monitoring "github.com/oracle/oci-go-sdk/v56/monitoring"
)

func MonitoringMetricDataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMonitoringMetricData,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"query": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resolution": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
						},
						"query": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"compartment_id_in_subtree": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"end_time": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: utils.TimeDiffSuppressFunction,
						},
						"resolution": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"start_time": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: utils.TimeDiffSuppressFunction,
						},

						// Computed
						"aggregated_datapoints": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
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
						"dimensions": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"metadata": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readMonitoringMetricData(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringMetricDataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MonitoringClient()

	return tfresource.ReadResource(sync)
}

type MonitoringMetricDataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_monitoring.MonitoringClient
	Res    *oci_monitoring.SummarizeMetricsDataResponse
}

func (s *MonitoringMetricDataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MonitoringMetricDataDataSourceCrud) Get() error {
	request := oci_monitoring.SummarizeMetricsDataRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if endTime, ok := s.D.GetOkExists("end_time"); ok {
		tmp, err := time.Parse(time.RFC3339, endTime.(string))
		if err != nil {
			return err
		}
		request.EndTime = &oci_common.SDKTime{Time: tmp}
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.Namespace = &tmp
	}

	if query, ok := s.D.GetOkExists("query"); ok {
		tmp := query.(string)
		request.Query = &tmp
	}

	if resolution, ok := s.D.GetOkExists("resolution"); ok {
		tmp := resolution.(string)
		request.Resolution = &tmp
	}

	if resourceGroup, ok := s.D.GetOkExists("resource_group"); ok {
		tmp := resourceGroup.(string)
		request.ResourceGroup = &tmp
	}

	if startTime, ok := s.D.GetOkExists("start_time"); ok {
		tmp, err := time.Parse(time.RFC3339, startTime.(string))
		if err != nil {
			return err
		}
		request.StartTime = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "monitoring")

	response, err := s.Client.SummarizeMetricsData(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MonitoringMetricDataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MonitoringMetricDataDataSource-", MonitoringMetricDataDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		metricData := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"namespace":      *r.Namespace,
		}

		aggregatedDatapoints := []interface{}{}
		for _, item := range r.AggregatedDatapoints {
			aggregatedDatapoints = append(aggregatedDatapoints, AggregatedDatapointToMap(item))
		}
		metricData["aggregated_datapoints"] = aggregatedDatapoints

		metricData["dimensions"] = r.Dimensions

		metricData["metadata"] = r.Metadata

		if r.Name != nil {
			metricData["name"] = *r.Name
		}

		if r.Resolution != nil {
			metricData["resolution"] = *r.Resolution
		}

		if r.ResourceGroup != nil {
			metricData["resource_group"] = *r.ResourceGroup
		}

		resources = append(resources, metricData)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MonitoringMetricDataDataSource().Schema["metric_data"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("metric_data", resources); err != nil {
		return err
	}

	return nil
}

func AggregatedDatapointToMap(obj oci_monitoring.AggregatedDatapoint) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Timestamp != nil {
		result["timestamp"] = obj.Timestamp.String()
	}

	if obj.Value != nil {
		result["value"] = float64(*obj.Value)
	}

	return result
}
