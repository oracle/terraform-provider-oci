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

func FleetAppsManagementComplianceRecordCountsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementComplianceRecordCounts,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_record_aggregation_collection": {
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
									"compliance_record_count_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"dimensions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"compliance_level": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"compliance_state": {
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
				},
			},
		},
	}
}

func readFleetAppsManagementComplianceRecordCounts(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementComplianceRecordCountsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementComplianceRecordCountsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.SummarizeComplianceRecordCountsResponse
}

func (s *FleetAppsManagementComplianceRecordCountsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementComplianceRecordCountsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.SummarizeComplianceRecordCountsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.SummarizeComplianceRecordCounts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.SummarizeComplianceRecordCounts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementComplianceRecordCountsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementComplianceRecordCountsDataSource-", FleetAppsManagementComplianceRecordCountsDataSource(), s.D))
	resources := []map[string]interface{}{}
	complianceRecordCount := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ComplianceRecordAggregationToMap(item))
	}
	complianceRecordCount["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementComplianceRecordCountsDataSource().Schema["compliance_record_aggregation_collection"].Elem.(*schema.Resource).Schema)
		complianceRecordCount["items"] = items
	}

	resources = append(resources, complianceRecordCount)
	if err := s.D.Set("compliance_record_aggregation_collection", resources); err != nil {
		return err
	}

	return nil
}

func ComplianceRecordAggregationToMap(obj oci_fleet_apps_management.ComplianceRecordAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["compliance_record_count_count"] = int(*obj.Count)
	}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{ComplianceRecordDimensionToMap(obj.Dimensions)}
	}

	return result
}

func ComplianceRecordDimensionToMap(obj *oci_fleet_apps_management.ComplianceRecordDimension) map[string]interface{} {
	result := map[string]interface{}{}

	result["compliance_level"] = string(obj.ComplianceLevel)

	result["compliance_state"] = string(obj.ComplianceState)

	return result
}
