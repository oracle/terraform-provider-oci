// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseDbSystemComputePerformancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbSystemComputePerformances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"db_system_shape": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_system_compute_performances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compute_performance_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"cpu_core_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"network_bandwidth_in_gbps": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"network_iops": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"network_throughput_in_mbps": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseDbSystemComputePerformances(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemComputePerformancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbSystemComputePerformancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbSystemComputePerformancesResponse
}

func (s *DatabaseDbSystemComputePerformancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbSystemComputePerformancesDataSourceCrud) Get() error {
	request := oci_database.ListDbSystemComputePerformancesRequest{}

	if dbSystemShape, ok := s.D.GetOkExists("db_system_shape"); ok {
		tmp := dbSystemShape.(string)
		request.DbSystemShape = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDbSystemComputePerformances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDbSystemComputePerformancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbSystemComputePerformancesDataSource-", DatabaseDbSystemComputePerformancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbSystemComputePerformance := map[string]interface{}{}

		computePerformanceList := []interface{}{}
		for _, item := range r.ComputePerformanceList {
			computePerformanceList = append(computePerformanceList, ComputePerformanceSummaryToMap(item))
		}
		dbSystemComputePerformance["compute_performance_list"] = computePerformanceList

		if r.Shape != nil {
			dbSystemComputePerformance["shape"] = *r.Shape
		}

		resources = append(resources, dbSystemComputePerformance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDbSystemComputePerformancesDataSource().Schema["db_system_compute_performances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_system_compute_performances", resources); err != nil {
		return err
	}

	return nil
}

func ComputePerformanceSummaryToMap(obj oci_database.ComputePerformanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CpuCoreCount != nil {
		result["cpu_core_count"] = int(*obj.CpuCoreCount)
	}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float64(*obj.MemoryInGBs)
	}

	if obj.NetworkBandwidthInGbps != nil {
		result["network_bandwidth_in_gbps"] = float32(*obj.NetworkBandwidthInGbps)
	}

	if obj.NetworkIops != nil {
		result["network_iops"] = float32(*obj.NetworkIops)
	}

	if obj.NetworkThroughputInMbps != nil {
		result["network_throughput_in_mbps"] = float32(*obj.NetworkThroughputInMbps)
	}

	return result
}
