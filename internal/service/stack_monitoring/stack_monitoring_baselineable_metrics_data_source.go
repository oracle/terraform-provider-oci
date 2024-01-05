// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringBaselineableMetricsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStackMonitoringBaselineableMetrics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"baselineable_metric_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"baselineable_metric_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(StackMonitoringBaselineableMetricResource()),
						},
					},
				},
			},
		},
	}
}

func readStackMonitoringBaselineableMetrics(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringBaselineableMetricsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringBaselineableMetricsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.ListBaselineableMetricsResponse
}

func (s *StackMonitoringBaselineableMetricsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringBaselineableMetricsDataSourceCrud) Get() error {
	request := oci_stack_monitoring.ListBaselineableMetricsRequest{}

	if baselineableMetricId, ok := s.D.GetOkExists("id"); ok {
		tmp := baselineableMetricId.(string)
		request.BaselineableMetricId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if metricNamespace, ok := s.D.GetOkExists("metric_namespace"); ok {
		tmp := metricNamespace.(string)
		request.MetricNamespace = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if resourceGroup, ok := s.D.GetOkExists("resource_group"); ok {
		tmp := resourceGroup.(string)
		request.ResourceGroup = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.ListBaselineableMetrics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBaselineableMetrics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StackMonitoringBaselineableMetricsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StackMonitoringBaselineableMetricsDataSource-", StackMonitoringBaselineableMetricsDataSource(), s.D))
	resources := []map[string]interface{}{}
	baselineableMetric := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BaselineableMetricSummaryToMap(item))
	}
	baselineableMetric["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, StackMonitoringBaselineableMetricsDataSource().Schema["baselineable_metric_summary_collection"].Elem.(*schema.Resource).Schema)
		baselineableMetric["items"] = items
	}

	resources = append(resources, baselineableMetric)
	if err := s.D.Set("baselineable_metric_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
