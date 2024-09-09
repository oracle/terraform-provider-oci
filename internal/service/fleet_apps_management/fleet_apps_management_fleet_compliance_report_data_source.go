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

func FleetAppsManagementFleetComplianceReportDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularFleetAppsManagementFleetComplianceReport,
		Schema: map[string]*schema.Schema{
			"compliance_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compliance_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compliance_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"products": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"product_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"targets": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"compliance_state": {
													Type:     schema.TypeString,
													Computed: true,
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
												"recommended_patches": {
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
												"target_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"target_name": {
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
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenancy_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenancy_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularFleetAppsManagementFleetComplianceReport(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetComplianceReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementFleetComplianceReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementClient
	Res    *oci_fleet_apps_management.GetComplianceReportResponse
}

func (s *FleetAppsManagementFleetComplianceReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementFleetComplianceReportDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetComplianceReportRequest{}

	if complianceReportId, ok := s.D.GetOkExists("compliance_report_id"); ok {
		tmp := complianceReportId.(string)
		request.ComplianceReportId = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetComplianceReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementFleetComplianceReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("compliance_state", s.Res.ComplianceState)

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, ComplianceReportResourceToMap(item))
	}
	s.D.Set("resources", resources)

	return nil
}

func ComplianceReportPatchDetailToMap(obj oci_fleet_apps_management.ComplianceReportPatchDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PatchDescription != nil {
		result["patch_description"] = string(*obj.PatchDescription)
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

func ComplianceReportProductToMap(obj oci_fleet_apps_management.ComplianceReportProduct) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ProductName != nil {
		result["product_name"] = string(*obj.ProductName)
	}

	targets := []interface{}{}
	for _, item := range obj.Targets {
		targets = append(targets, ComplianceReportTargetToMap(item))
	}
	result["targets"] = targets

	return result
}

func ComplianceReportResourceToMap(obj oci_fleet_apps_management.ComplianceReportResource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Compartment != nil {
		result["compartment"] = string(*obj.Compartment)
	}

	result["compliance_state"] = string(obj.ComplianceState)

	products := []interface{}{}
	for _, item := range obj.Products {
		products = append(products, ComplianceReportProductToMap(item))
	}
	result["products"] = products

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	if obj.TenancyName != nil {
		result["tenancy_name"] = string(*obj.TenancyName)
	}

	return result
}

func ComplianceReportTargetToMap(obj oci_fleet_apps_management.ComplianceReportTarget) map[string]interface{} {
	result := map[string]interface{}{}

	result["compliance_state"] = string(obj.ComplianceState)

	installedPatches := []interface{}{}
	for _, item := range obj.InstalledPatches {
		installedPatches = append(installedPatches, ComplianceReportPatchDetailToMap(item))
	}
	result["installed_patches"] = installedPatches

	recommendedPatches := []interface{}{}
	for _, item := range obj.RecommendedPatches {
		recommendedPatches = append(recommendedPatches, ComplianceReportPatchDetailToMap(item))
	}
	result["recommended_patches"] = recommendedPatches

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TargetName != nil {
		result["target_name"] = string(*obj.TargetName)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
