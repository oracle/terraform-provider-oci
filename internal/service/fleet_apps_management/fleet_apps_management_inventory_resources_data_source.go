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

func FleetAppsManagementInventoryResourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementInventoryResources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"defined_tag_equals": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"freeform_tag_equals": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"inventory_properties": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"matching_criteria": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inventory_resource_collection": {
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
									"availability_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_region": {
										Type:     schema.TypeString,
										Computed: true,
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

func readFleetAppsManagementInventoryResources(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementInventoryResourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementInventoryResourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementClient
	Res    *oci_fleet_apps_management.ListInventoryResourcesResponse
}

func (s *FleetAppsManagementInventoryResourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementInventoryResourcesDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListInventoryResourcesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTagEquals, ok := s.D.GetOkExists("defined_tag_equals"); ok {
		interfaces := definedTagEquals.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("defined_tag_equals") {
			request.DefinedTagEquals = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTagEquals, ok := s.D.GetOkExists("freeform_tag_equals"); ok {
		interfaces := freeformTagEquals.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("freeform_tag_equals") {
			request.FreeformTagEquals = tmp
		}
	}

	if inventoryProperties, ok := s.D.GetOkExists("inventory_properties"); ok {
		interfaces := inventoryProperties.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("inventory_properties") {
			request.InventoryProperties = tmp
		}
	}

	if matchingCriteria, ok := s.D.GetOkExists("matching_criteria"); ok {
		tmp := matchingCriteria.(string)
		request.MatchingCriteria = &tmp
	}

	if resourceCompartmentId, ok := s.D.GetOkExists("resource_compartment_id"); ok {
		tmp := resourceCompartmentId.(string)
		request.ResourceCompartmentId = &tmp
	}

	if resourceRegion, ok := s.D.GetOkExists("resource_region"); ok {
		tmp := resourceRegion.(string)
		request.ResourceRegion = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		tmp := state.(string)
		request.LifecycleState = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListInventoryResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInventoryResources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementInventoryResourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementInventoryResourcesDataSource-", FleetAppsManagementInventoryResourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	inventoryResource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InventoryResourceSummaryToMap(item))
	}
	inventoryResource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementInventoryResourcesDataSource().Schema["inventory_resource_collection"].Elem.(*schema.Resource).Schema)
		inventoryResource["items"] = items
	}

	resources = append(resources, inventoryResource)
	if err := s.D.Set("inventory_resource_collection", resources); err != nil {
		return err
	}

	return nil
}

func InventoryResourceSummaryToMap(obj oci_fleet_apps_management.InventoryResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ResourceCompartmentId != nil {
		result["resource_compartment_id"] = string(*obj.ResourceCompartmentId)
	}

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
	}

	if obj.LifecycleState != nil {
		result["state"] = string(*obj.LifecycleState)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}
