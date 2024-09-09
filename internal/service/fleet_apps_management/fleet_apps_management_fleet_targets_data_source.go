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

func FleetAppsManagementFleetTargetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementFleetTargets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"product": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_target_collection": {
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
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compliance_state": {
										Type:     schema.TypeString,
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
									"product": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"resource_display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"resource_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"version": {
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

func readFleetAppsManagementFleetTargets(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetTargetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementFleetTargetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementClient
	Res    *oci_fleet_apps_management.ListFleetTargetsResponse
}

func (s *FleetAppsManagementFleetTargetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementFleetTargetsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListFleetTargetsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if product, ok := s.D.GetOkExists("product"); ok {
		tmp := product.(string)
		request.Product = &tmp
	}

	if resourceDisplayName, ok := s.D.GetOkExists("resource_display_name"); ok {
		tmp := resourceDisplayName.(string)
		request.ResourceDisplayName = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListFleetTargets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFleetTargets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementFleetTargetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementFleetTargetsDataSource-", FleetAppsManagementFleetTargetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetTarget := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FleetTargetSummaryToMap(item))
	}
	fleetTarget["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementFleetTargetsDataSource().Schema["fleet_target_collection"].Elem.(*schema.Resource).Schema)
		fleetTarget["items"] = items
	}

	resources = append(resources, fleetTarget)
	if err := s.D.Set("fleet_target_collection", resources); err != nil {
		return err
	}

	return nil
}

func FleetTargetSummaryToMap(obj oci_fleet_apps_management.FleetTargetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["compliance_state"] = string(obj.ComplianceState)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Product != nil {
		result["product"] = string(*obj.Product)
	}

	if obj.Resource != nil {
		result["resource"] = []interface{}{TargetResourceToMap(obj.Resource)}
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func TargetResourceToMap(obj *oci_fleet_apps_management.TargetResource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ResourceDisplayName != nil {
		result["resource_display_name"] = string(*obj.ResourceDisplayName)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	return result
}
