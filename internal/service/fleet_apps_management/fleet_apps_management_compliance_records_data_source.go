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

func FleetAppsManagementComplianceRecordsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementComplianceRecords,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_stack": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_record_collection": {
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
									"entity_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"patch": {
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
												"patch_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"patch_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"product": {
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
															"product_stack": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"product_version": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"severity": {
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
									"policy": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"compliance_policy_display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"compliance_policy_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"compliance_policy_rule_display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"compliance_policy_rule_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"grace_period": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"patch_selection": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"days_since_release": {
																Type:     schema.TypeInt,
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
															"selection_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"resource": {
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
												"compartment_id": {
													Type:     schema.TypeString,
													Computed: true,
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
											},
										},
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
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
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

func readFleetAppsManagementComplianceRecords(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementComplianceRecordsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementComplianceRecordsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.ListComplianceRecordsResponse
}

func (s *FleetAppsManagementComplianceRecordsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementComplianceRecordsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListComplianceRecordsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if complianceState, ok := s.D.GetOkExists("compliance_state"); ok {
		tmp := complianceState.(string)
		request.ComplianceState = &tmp
	}

	if entityId, ok := s.D.GetOkExists("entity_id"); ok {
		tmp := entityId.(string)
		request.EntityId = &tmp
	}

	if productName, ok := s.D.GetOkExists("product_name"); ok {
		tmp := productName.(string)
		request.ProductName = &tmp
	}

	if productStack, ok := s.D.GetOkExists("product_stack"); ok {
		tmp := productStack.(string)
		request.ProductStack = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if targetName, ok := s.D.GetOkExists("target_name"); ok {
		tmp := targetName.(string)
		request.TargetName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListComplianceRecords(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComplianceRecords(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementComplianceRecordsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementComplianceRecordsDataSource-", FleetAppsManagementComplianceRecordsDataSource(), s.D))
	resources := []map[string]interface{}{}
	complianceRecord := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ComplianceRecordSummaryToMap(item))
	}
	complianceRecord["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementComplianceRecordsDataSource().Schema["compliance_record_collection"].Elem.(*schema.Resource).Schema)
		complianceRecord["items"] = items
	}

	resources = append(resources, complianceRecord)
	if err := s.D.Set("compliance_record_collection", resources); err != nil {
		return err
	}

	return nil
}

func ComplianceDetailPolicyToMap(obj *oci_fleet_apps_management.ComplianceDetailPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompliancePolicyDisplayName != nil {
		result["compliance_policy_display_name"] = string(*obj.CompliancePolicyDisplayName)
	}

	if obj.CompliancePolicyId != nil {
		result["compliance_policy_id"] = string(*obj.CompliancePolicyId)
	}

	if obj.CompliancePolicyRuleDisplayName != nil {
		result["compliance_policy_rule_display_name"] = string(*obj.CompliancePolicyRuleDisplayName)
	}

	if obj.CompliancePolicyRuleId != nil {
		result["compliance_policy_rule_id"] = string(*obj.CompliancePolicyRuleId)
	}

	if obj.GracePeriod != nil {
		result["grace_period"] = string(*obj.GracePeriod)
	}

	if obj.PatchSelection != nil {
		patchSelectionArray := []interface{}{}
		if patchSelectionMap := PatchSelectionDetailsToMap(&obj.PatchSelection); patchSelectionMap != nil {
			patchSelectionArray = append(patchSelectionArray, patchSelectionMap)
		}
		result["patch_selection"] = patchSelectionArray
	}

	return result
}

func ComplianceDetailProductToMap(obj *oci_fleet_apps_management.ComplianceDetailProduct) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ProductName != nil {
		result["product_name"] = string(*obj.ProductName)
	}

	if obj.ProductStack != nil {
		result["product_stack"] = string(*obj.ProductStack)
	}

	if obj.ProductVersion != nil {
		result["product_version"] = string(*obj.ProductVersion)
	}

	return result
}

func ComplianceDetailResourceToMap(obj *oci_fleet_apps_management.ComplianceDetailResource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Compartment != nil {
		result["compartment"] = string(*obj.Compartment)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
	}

	return result
}

func ComplianceDetailTargetToMap(obj *oci_fleet_apps_management.ComplianceDetailTarget) map[string]interface{} {
	result := map[string]interface{}{}

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

func CompliancePatchDetailToMap(obj *oci_fleet_apps_management.CompliancePatchDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PatchDescription != nil {
		result["patch_description"] = string(*obj.PatchDescription)
	}

	if obj.PatchId != nil {
		result["patch_id"] = string(*obj.PatchId)
	}

	if obj.PatchName != nil {
		result["patch_name"] = string(*obj.PatchName)
	}

	if obj.PatchType != nil {
		result["patch_type"] = string(*obj.PatchType)
	}

	if obj.Product != nil {
		result["product"] = []interface{}{ComplianceDetailProductToMap(obj.Product)}
	}

	result["severity"] = string(obj.Severity)

	if obj.TimeReleased != nil {
		result["time_released"] = obj.TimeReleased.String()
	}

	return result
}

func ComplianceRecordSummaryToMap(obj oci_fleet_apps_management.ComplianceRecordSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["compliance_state"] = string(obj.ComplianceState)

	if obj.EntityDisplayName != nil {
		result["entity_display_name"] = string(*obj.EntityDisplayName)
	}

	if obj.EntityId != nil {
		result["entity_id"] = string(*obj.EntityId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Patch != nil {
		result["patch"] = []interface{}{CompliancePatchDetailToMap(obj.Patch)}
	}

	if obj.Policy != nil {
		result["policy"] = []interface{}{ComplianceDetailPolicyToMap(obj.Policy)}
	}

	if obj.Resource != nil {
		result["resource"] = []interface{}{ComplianceDetailResourceToMap(obj.Resource)}
	}

	result["state"] = string(obj.LifecycleState)

	if obj.Target != nil {
		result["target"] = []interface{}{ComplianceDetailTargetToMap(obj.Target)}
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
