// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabaseTableStatisticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseTableStatistics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"table_statistics_collection": {
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
									"count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"percentage": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"type": {
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

func readDatabaseManagementManagedDatabaseTableStatistics(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseTableStatisticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseTableStatisticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListTableStatisticsResponse
}

func (s *DatabaseManagementManagedDatabaseTableStatisticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseTableStatisticsDataSourceCrud) Get() error {
	request := oci_database_management.ListTableStatisticsRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListTableStatistics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseTableStatisticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseTableStatisticsDataSource-", DatabaseManagementManagedDatabaseTableStatisticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseTableStatistic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TableStatisticSummaryToMap(item))
	}
	managedDatabaseTableStatistic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseTableStatisticsDataSource().Schema["table_statistics_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseTableStatistic["items"] = items
	}

	resources = append(resources, managedDatabaseTableStatistic)
	if err := s.D.Set("table_statistics_collection", resources); err != nil {
		return err
	}

	return nil
}

func TableStatisticSummaryToMap(obj oci_database_management.TableStatisticSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["count"] = int(*obj.Count)
	}

	if obj.Percentage != nil {
		result["percentage"] = float64(*obj.Percentage)
	}

	result["type"] = string(obj.Type)

	return result
}
