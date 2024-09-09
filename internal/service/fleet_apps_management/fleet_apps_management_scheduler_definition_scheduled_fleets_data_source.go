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

func FleetAppsManagementSchedulerDefinitionScheduledFleetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementSchedulerDefinitionScheduledFleets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduler_definition_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scheduled_fleet_collection": {
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
									"action_group_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"application_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"count_of_affected_resources": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"count_of_affected_targets": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"tenancy_id": {
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

func readFleetAppsManagementSchedulerDefinitionScheduledFleets(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementSchedulerDefinitionScheduledFleetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementSchedulerDefinitionScheduledFleetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.ListScheduledFleetsResponse
}

func (s *FleetAppsManagementSchedulerDefinitionScheduledFleetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementSchedulerDefinitionScheduledFleetsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListScheduledFleetsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if schedulerDefinitionId, ok := s.D.GetOkExists("scheduler_definition_id"); ok {
		tmp := schedulerDefinitionId.(string)
		request.SchedulerDefinitionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListScheduledFleets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListScheduledFleets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementSchedulerDefinitionScheduledFleetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementSchedulerDefinitionScheduledFleetsDataSource-", FleetAppsManagementSchedulerDefinitionScheduledFleetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	schedulerDefinitionScheduledFleet := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScheduledFleetSummaryToMap(item))
	}
	schedulerDefinitionScheduledFleet["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementSchedulerDefinitionScheduledFleetsDataSource().Schema["scheduled_fleet_collection"].Elem.(*schema.Resource).Schema)
		schedulerDefinitionScheduledFleet["items"] = items
	}

	resources = append(resources, schedulerDefinitionScheduledFleet)
	if err := s.D.Set("scheduled_fleet_collection", resources); err != nil {
		return err
	}

	return nil
}

func ScheduledFleetSummaryToMap(obj oci_fleet_apps_management.ScheduledFleetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["action_group_types"] = obj.ActionGroupTypes

	result["application_types"] = obj.ApplicationTypes

	if obj.CountOfAffectedResources != nil {
		result["count_of_affected_resources"] = int(*obj.CountOfAffectedResources)
	}

	if obj.CountOfAffectedTargets != nil {
		result["count_of_affected_targets"] = int(*obj.CountOfAffectedTargets)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	return result
}
