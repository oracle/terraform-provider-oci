// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementReportMetadataDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readFleetAppsManagementReportMetadataWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"report_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"report_metadata_collection": {
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
									"column_metadata": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"default_order_clause": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"sort_by": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sort_order": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"filters": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value_source": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"metric": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
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

func readFleetAppsManagementReportMetadataWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetAppsManagementReportMetadataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FleetAppsManagementReportMetadataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.ListReportMetadataResponse
}

func (s *FleetAppsManagementReportMetadataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementReportMetadataDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_fleet_apps_management.ListReportMetadataRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if reportName, ok := s.D.GetOkExists("report_name"); ok {
		tmp := reportName.(string)
		request.ReportName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListReportMetadata(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListReportMetadata(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementReportMetadataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementReportMetadataDataSource-", FleetAppsManagementReportMetadataDataSource(), s.D))
	resources := []map[string]interface{}{}
	reportMetadata := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ReportMetadataSummaryToMap(item))
	}
	reportMetadata["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementReportMetadataDataSource().Schema["report_metadata_collection"].Elem.(*schema.Resource).Schema)
		reportMetadata["items"] = items
	}

	resources = append(resources, reportMetadata)
	if err := s.D.Set("report_metadata_collection", resources); err != nil {
		return err
	}

	return nil
}

func ColumnMetadataToMap(obj oci_fleet_apps_management.ColumnMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["type"] = string(obj.Type)

	return result
}

func OrderClauseToMap(obj oci_fleet_apps_management.OrderClause) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SortBy != nil {
		result["sort_by"] = string(*obj.SortBy)
	}

	result["sort_order"] = string(obj.SortOrder)

	return result
}

func ReportFilterToMap(obj oci_fleet_apps_management.ReportFilter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ValueSource != nil {
		result["value_source"] = string(*obj.ValueSource)
	}

	return result
}

func ReportMetadataSummaryToMap(obj oci_fleet_apps_management.ReportMetadataSummary) map[string]interface{} {
	result := map[string]interface{}{}

	columnMetadata := []interface{}{}
	for _, item := range obj.ColumnMetadata {
		columnMetadata = append(columnMetadata, ColumnMetadataToMap(item))
	}
	result["column_metadata"] = columnMetadata

	defaultOrderClause := []interface{}{}
	for _, item := range obj.DefaultOrderClause {
		defaultOrderClause = append(defaultOrderClause, OrderClauseToMap(item))
	}
	result["default_order_clause"] = defaultOrderClause

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	filters := []interface{}{}
	for _, item := range obj.Filters {
		filters = append(filters, ReportFilterToMap(item))
	}
	result["filters"] = filters

	if obj.Metric != nil {
		result["metric"] = string(*obj.Metric)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
