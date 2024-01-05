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

func DatabaseDbSystemStoragePerformancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbSystemStoragePerformances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"shape_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_management": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_storage_performances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"data_storage_performance_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"balanced_disk_performance": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"disk_iops": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"disk_throughput_in_mbps": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"high_disk_performance": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"disk_iops": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"disk_throughput_in_mbps": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"size_in_gbs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"reco_storage_performance_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"balanced_disk_performance": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"disk_iops": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"disk_throughput_in_mbps": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"high_disk_performance": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"disk_iops": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"disk_throughput_in_mbps": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"size_in_gbs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"shape_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseDbSystemStoragePerformances(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemStoragePerformancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbSystemStoragePerformancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbSystemStoragePerformancesResponse
}

func (s *DatabaseDbSystemStoragePerformancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbSystemStoragePerformancesDataSourceCrud) Get() error {
	request := oci_database.ListDbSystemStoragePerformancesRequest{}

	if shapeType, ok := s.D.GetOkExists("shape_type"); ok {
		tmp := oci_database.DbSystemStoragePerformanceSummaryShapeTypeEnum(shapeType.(string))
		request.ShapeType = (*string)(&tmp)
	}

	if storageManagement, ok := s.D.GetOkExists("storage_management"); ok {
		request.StorageManagement = oci_database.DbSystemOptionsStorageManagementEnum(storageManagement.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDbSystemStoragePerformances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDbSystemStoragePerformancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbSystemStoragePerformancesDataSource-", DatabaseDbSystemStoragePerformancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbSystemStoragePerformance := map[string]interface{}{}

		dataStoragePerformanceList := []interface{}{}
		for _, item := range r.DataStoragePerformanceList {
			dataStoragePerformanceList = append(dataStoragePerformanceList, StoragePerformanceDetailsToMap(item))
		}
		dbSystemStoragePerformance["data_storage_performance_list"] = dataStoragePerformanceList

		recoStoragePerformanceList := []interface{}{}
		for _, item := range r.RecoStoragePerformanceList {
			recoStoragePerformanceList = append(recoStoragePerformanceList, StoragePerformanceDetailsToMap(item))
		}
		dbSystemStoragePerformance["reco_storage_performance_list"] = recoStoragePerformanceList

		dbSystemStoragePerformance["shape_type"] = r.ShapeType

		resources = append(resources, dbSystemStoragePerformance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDbSystemStoragePerformancesDataSource().Schema["db_system_storage_performances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_system_storage_performances", resources); err != nil {
		return err
	}

	return nil
}

func DiskPerformanceDetailsToMap(obj *oci_database.DiskPerformanceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DiskIops != nil {
		result["disk_iops"] = float32(*obj.DiskIops)
	}

	if obj.DiskThroughputInMbps != nil {
		result["disk_throughput_in_mbps"] = float32(*obj.DiskThroughputInMbps)
	}

	return result
}

func StoragePerformanceDetailsToMap(obj oci_database.StoragePerformanceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BalancedDiskPerformance != nil {
		result["balanced_disk_performance"] = []interface{}{DiskPerformanceDetailsToMap(obj.BalancedDiskPerformance)}
	}

	if obj.HighDiskPerformance != nil {
		result["high_disk_performance"] = []interface{}{DiskPerformanceDetailsToMap(obj.HighDiskPerformance)}
	}

	if obj.SizeInGBs != nil {
		result["size_in_gbs"] = int(*obj.SizeInGBs)
	}

	return result
}
