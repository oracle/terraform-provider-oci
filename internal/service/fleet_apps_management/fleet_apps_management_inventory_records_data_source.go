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

func FleetAppsManagementInventoryRecordsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementInventoryRecords,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_details_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inventory_record_collection": {
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
									"architecture": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"components": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"component_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"component_path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"component_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"properties": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"value": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"installed_patches": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"patch_description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"patch_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"patch_level": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"patch_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"patch_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_applied": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_released": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"os_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"properties": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
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
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_product_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_product_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_resource_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
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

func readFleetAppsManagementInventoryRecords(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementInventoryRecordsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementInventoryRecordsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.ListInventoryRecordsResponse
}

func (s *FleetAppsManagementInventoryRecordsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementInventoryRecordsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListInventoryRecordsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if isDetailsRequired, ok := s.D.GetOkExists("is_details_required"); ok {
		tmp := isDetailsRequired.(bool)
		request.IsDetailsRequired = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListInventoryRecords(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInventoryRecords(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementInventoryRecordsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementInventoryRecordsDataSource-", FleetAppsManagementInventoryRecordsDataSource(), s.D))
	resources := []map[string]interface{}{}
	inventoryRecord := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InventoryRecordSummaryToMap(item))
	}
	inventoryRecord["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementInventoryRecordsDataSource().Schema["inventory_record_collection"].Elem.(*schema.Resource).Schema)
		inventoryRecord["items"] = items
	}

	resources = append(resources, inventoryRecord)
	if err := s.D.Set("inventory_record_collection", resources); err != nil {
		return err
	}

	return nil
}

func InventoryRecordComponentToMap(obj oci_fleet_apps_management.InventoryRecordComponent) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComponentName != nil {
		result["component_name"] = string(*obj.ComponentName)
	}

	if obj.ComponentPath != nil {
		result["component_path"] = string(*obj.ComponentPath)
	}

	if obj.ComponentVersion != nil {
		result["component_version"] = string(*obj.ComponentVersion)
	}

	properties := []interface{}{}
	for _, item := range obj.Properties {
		properties = append(properties, InventoryRecordPropertyToMap(item))
	}
	result["properties"] = properties

	return result
}

func InventoryRecordPatchDetailsToMap(obj oci_fleet_apps_management.InventoryRecordPatchDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PatchDescription != nil {
		result["patch_description"] = string(*obj.PatchDescription)
	}

	if obj.PatchId != nil {
		result["patch_id"] = string(*obj.PatchId)
	}

	if obj.PatchLevel != nil {
		result["patch_level"] = string(*obj.PatchLevel)
	}

	if obj.PatchName != nil {
		result["patch_name"] = string(*obj.PatchName)
	}

	if obj.PatchType != nil {
		result["patch_type"] = string(*obj.PatchType)
	}

	if obj.TimeApplied != nil {
		result["time_applied"] = obj.TimeApplied.String()
	}

	if obj.TimeReleased != nil {
		result["time_released"] = obj.TimeReleased.String()
	}

	return result
}

func InventoryRecordPropertyToMap(obj oci_fleet_apps_management.InventoryRecordProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func InventoryRecordSummaryToMap(obj oci_fleet_apps_management.InventoryRecordSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Architecture != nil {
		result["architecture"] = string(*obj.Architecture)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	components := []interface{}{}
	for _, item := range obj.Components {
		components = append(components, InventoryRecordComponentToMap(item))
	}
	result["components"] = components

	installedPatches := []interface{}{}
	for _, item := range obj.InstalledPatches {
		installedPatches = append(installedPatches, InventoryRecordPatchDetailsToMap(item))
	}
	result["installed_patches"] = installedPatches

	if obj.OsType != nil {
		result["os_type"] = string(*obj.OsType)
	}

	properties := []interface{}{}
	for _, item := range obj.Properties {
		properties = append(properties, InventoryRecordPropertyToMap(item))
	}
	result["properties"] = properties

	result["state"] = string(obj.LifecycleState)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TargetName != nil {
		result["target_name"] = string(*obj.TargetName)
	}

	if obj.TargetProductId != nil {
		result["target_product_id"] = string(*obj.TargetProductId)
	}

	if obj.TargetProductName != nil {
		result["target_product_name"] = string(*obj.TargetProductName)
	}

	if obj.TargetResourceId != nil {
		result["target_resource_id"] = string(*obj.TargetResourceId)
	}

	if obj.TargetResourceName != nil {
		result["target_resource_name"] = string(*obj.TargetResourceName)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
