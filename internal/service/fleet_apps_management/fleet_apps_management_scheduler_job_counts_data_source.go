// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementSchedulerJobCountsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementSchedulerJobCounts,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduler_job_aggregation_collection": {
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
									"dimensions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"lifecycle_details": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"scheduler_job_count_count": {
										Type:     schema.TypeInt,
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

func readFleetAppsManagementSchedulerJobCounts(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementSchedulerJobCountsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementSchedulerJobCountsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.SummarizeSchedulerJobCountsResponse
}

func (s *FleetAppsManagementSchedulerJobCountsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementSchedulerJobCountsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.SummarizeSchedulerJobCountsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.SummarizeSchedulerJobCounts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.SummarizeSchedulerJobCounts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementSchedulerJobCountsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementSchedulerJobCountsDataSource-", FleetAppsManagementSchedulerJobCountsDataSource(), s.D))
	resources := []map[string]interface{}{}
	schedulerJobCount := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SchedulerJobAggregationToMap(item))
	}
	schedulerJobCount["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementSchedulerJobCountsDataSource().Schema["scheduler_job_aggregation_collection"].Elem.(*schema.Resource).Schema)
		schedulerJobCount["items"] = items
	}

	resources = append(resources, schedulerJobCount)
	if err := s.D.Set("scheduler_job_aggregation_collection", resources); err != nil {
		return err
	}

	return nil
}

func SchedulerJobAggregationToMap(obj oci_fleet_apps_management.SchedulerJobAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{SchedulerJobDimensionToMap(obj.Dimensions)}
	}

	if obj.Count != nil {
		result["scheduler_job_count_count"] = int(*obj.Count)
	}

	return result
}

func SchedulerJobDimensionToMap(obj *oci_fleet_apps_management.SchedulerJobDimension) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	return result
}
