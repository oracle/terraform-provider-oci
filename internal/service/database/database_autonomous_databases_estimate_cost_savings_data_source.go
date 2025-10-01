// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousDatabasesEstimateCostSavingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDatabasesEstimateCostSavings,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_cpu_autoscale": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"estimate_cost_savings_summary_collection": {
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
									"cost_savings_with_elastic_pool": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"estimated_usage_without_elastic_pool": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_cpu_autoscale": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"time_ended": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"usage_with_elastic_pool": {
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

func readDatabaseAutonomousDatabasesEstimateCostSavings(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabasesEstimateCostSavingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabasesEstimateCostSavingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListEstimateCostSavingsResponse
}

func (s *DatabaseAutonomousDatabasesEstimateCostSavingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabasesEstimateCostSavingsDataSourceCrud) Get() error {
	request := oci_database.ListEstimateCostSavingsRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if isCpuAutoscale, ok := s.D.GetOkExists("is_cpu_autoscale"); ok {
		tmp := isCpuAutoscale.(bool)
		request.IsCpuAutoscale = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListEstimateCostSavings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEstimateCostSavings(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDatabasesEstimateCostSavingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabasesEstimateCostSavingsDataSource-", DatabaseAutonomousDatabasesEstimateCostSavingsDataSource(), s.D))
	resources := []map[string]interface{}{}
	autonomousDatabasesEstimateCostSaving := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EstimateCostSavingSummaryToMap(item))
	}
	autonomousDatabasesEstimateCostSaving["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseAutonomousDatabasesEstimateCostSavingsDataSource().Schema["estimate_cost_savings_summary_collection"].Elem.(*schema.Resource).Schema)
		autonomousDatabasesEstimateCostSaving["items"] = items
	}

	resources = append(resources, autonomousDatabasesEstimateCostSaving)
	if err := s.D.Set("estimate_cost_savings_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func EstimateCostSavingSummaryToMap(obj oci_database.EstimateCostSavingSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CostSavingsWithElasticPool != nil {
		result["cost_savings_with_elastic_pool"] = float64(*obj.CostSavingsWithElasticPool)
	}

	if obj.EstimatedUsageWithoutElasticPool != nil {
		result["estimated_usage_without_elastic_pool"] = strconv.FormatInt(*obj.EstimatedUsageWithoutElasticPool, 10)
	}

	if obj.IsCpuAutoscale != nil {
		result["is_cpu_autoscale"] = bool(*obj.IsCpuAutoscale)
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.UsageWithElasticPool != nil {
		result["usage_with_elastic_pool"] = strconv.FormatInt(*obj.UsageWithElasticPool, 10)
	}

	return result
}
