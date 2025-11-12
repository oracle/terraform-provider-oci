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

func FleetAppsManagementRunbookExportStatusesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readFleetAppsManagementRunbookExportStatusesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"runbook_export_status_collection": {
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
									"runbook_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"runbook_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"runbook_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tracking_id": {
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

func readFleetAppsManagementRunbookExportStatusesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetAppsManagementRunbookExportStatusesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FleetAppsManagementRunbookExportStatusesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementRunbooksClient
	Res    *oci_fleet_apps_management.ListRunbookExportStatusesResponse
}

func (s *FleetAppsManagementRunbookExportStatusesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementRunbookExportStatusesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_fleet_apps_management.ListRunbookExportStatusesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListRunbookExportStatuses(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRunbookExportStatuses(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementRunbookExportStatusesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementRunbookExportStatusesDataSource-", FleetAppsManagementRunbookExportStatusesDataSource(), s.D))
	resources := []map[string]interface{}{}
	runbookExportStatus := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RunbookExportStatusSummaryToMap(item))
	}
	runbookExportStatus["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementRunbookExportStatusesDataSource().Schema["runbook_export_status_collection"].Elem.(*schema.Resource).Schema)
		runbookExportStatus["items"] = items
	}

	resources = append(resources, runbookExportStatus)
	if err := s.D.Set("runbook_export_status_collection", resources); err != nil {
		return err
	}

	return nil
}

func RunbookExportStatusSummaryToMap(obj oci_fleet_apps_management.RunbookExportStatusSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RunbookId != nil {
		result["runbook_id"] = string(*obj.RunbookId)
	}

	if obj.RunbookName != nil {
		result["runbook_name"] = string(*obj.RunbookName)
	}

	if obj.RunbookVersion != nil {
		result["runbook_version"] = string(*obj.RunbookVersion)
	}

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.TrackingId != nil {
		result["tracking_id"] = string(*obj.TrackingId)
	}

	return result
}
