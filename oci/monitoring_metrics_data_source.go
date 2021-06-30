// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_monitoring "github.com/oracle/oci-go-sdk/v43/monitoring"
)

func init() {
	RegisterDatasource("oci_monitoring_metrics", MonitoringMetricsDataSource())
}

func MonitoringMetricsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMonitoringMetrics,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dimension_filters": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"group_by": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metrics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"compartment_id_in_subtree": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"dimension_filters": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"group_by": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"dimensions": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func readMonitoringMetrics(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringMetricsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).monitoringClient()

	return ReadResource(sync)
}

type MonitoringMetricsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_monitoring.MonitoringClient
	Res    *oci_monitoring.ListMetricsResponse
}

func (s *MonitoringMetricsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MonitoringMetricsDataSourceCrud) Get() error {
	request := oci_monitoring.ListMetricsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if dimensionFilters, ok := s.D.GetOkExists("dimension_filters"); ok {
		request.DimensionFilters = objectMapToStringMap(dimensionFilters.(map[string]interface{}))
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		interfaces := groupBy.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_by") {
			request.GroupBy = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.Namespace = &tmp
	}

	if resourceGroup, ok := s.D.GetOkExists("resource_group"); ok {
		tmp := resourceGroup.(string)
		request.ResourceGroup = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "monitoring")

	response, err := s.Client.ListMetrics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMetrics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MonitoringMetricsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("MonitoringMetricsDataSource-", MonitoringMetricsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		metric := map[string]interface{}{}

		if r.CompartmentId != nil {
			metric["compartment_id"] = *r.CompartmentId
		}

		metric["dimensions"] = r.Dimensions

		if r.Name != nil {
			metric["name"] = *r.Name
		}

		if r.Namespace != nil {
			metric["namespace"] = *r.Namespace
		}

		if r.ResourceGroup != nil {
			metric["resource_group"] = *r.ResourceGroup
		}

		resources = append(resources, metric)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, MonitoringMetricsDataSource().Schema["metrics"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("metrics", resources); err != nil {
		return err
	}

	return nil
}
