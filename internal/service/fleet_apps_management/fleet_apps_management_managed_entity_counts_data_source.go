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

func FleetAppsManagementManagedEntityCountsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementManagedEntityCounts,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_entity_aggregation_collection": {
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
												"entity": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"managed_entity_count_count": {
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

func readFleetAppsManagementManagedEntityCounts(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementManagedEntityCountsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementManagedEntityCountsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.SummarizeManagedEntityCountsResponse
}

func (s *FleetAppsManagementManagedEntityCountsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementManagedEntityCountsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.SummarizeManagedEntityCountsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.SummarizeManagedEntityCounts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.SummarizeManagedEntityCounts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementManagedEntityCountsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementManagedEntityCountsDataSource-", FleetAppsManagementManagedEntityCountsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedEntityCount := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedEntityAggregationToMap(item))
	}
	managedEntityCount["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementManagedEntityCountsDataSource().Schema["managed_entity_aggregation_collection"].Elem.(*schema.Resource).Schema)
		managedEntityCount["items"] = items
	}

	resources = append(resources, managedEntityCount)
	if err := s.D.Set("managed_entity_aggregation_collection", resources); err != nil {
		return err
	}

	return nil
}

func ManagedEntityAggregationToMap(obj oci_fleet_apps_management.ManagedEntityAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{ManagedEntityDimensionToMap(obj.Dimensions)}
	}

	if obj.Count != nil {
		result["managed_entity_count_count"] = int(*obj.Count)
	}

	return result
}

func ManagedEntityDimensionToMap(obj *oci_fleet_apps_management.ManagedEntityDimension) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity"] = string(obj.Entity)

	return result
}
