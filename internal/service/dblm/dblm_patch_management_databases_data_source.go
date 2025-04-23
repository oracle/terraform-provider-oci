// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dblm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dblm "github.com/oracle/oci-go-sdk/v65/dblm"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DblmPatchManagementDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDblmPatchManagementDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_release": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drifter_patch_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"image_compliance": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"patch_databases_collection": {
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
									"additional_patches": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"category": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"patch_id": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"patch_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"current_patch_watermark": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"host_or_cluster": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"image_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"created_by": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"current_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"image_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"image_owner": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"image_status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"image_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"subscribed_image": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_image_creation": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"up_to_date_image_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"oracle_home_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"patch_activity_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"deploy_operation_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"deploy_status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"deploy_task_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"migrate_listener_operation_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"migrate_listener_status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"migrate_listener_task_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"update_operation_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"update_status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"update_task_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"patch_compliance_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"patch_compliance_status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"patch_compliance_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"patch_user": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"release": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"release_full_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sudo_file_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"vulnerabilities_summary": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Computed
												"total": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"critical": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"high": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"medium": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"info": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"low": {
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
				},
			},
		},
	}
}

func readDblmPatchManagementDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DblmPatchManagementDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbLifeCycleManagementClient()

	return tfresource.ReadResource(sync)
}

type DblmPatchManagementDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dblm.DbLifeCycleManagementClient
	Res    *oci_dblm.ListDatabasesResponse
}

func (s *DblmPatchManagementDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DblmPatchManagementDatabasesDataSourceCrud) Get() error {
	request := oci_dblm.ListDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if databaseRelease, ok := s.D.GetOkExists("database_release"); ok {
		tmp := databaseRelease.(string)
		request.DatabaseRelease = &tmp
	}

	if databaseType, ok := s.D.GetOkExists("database_type"); ok {
		request.DatabaseType = oci_dblm.ListDatabasesDatabaseTypeEnum(databaseType.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drifterPatchId, ok := s.D.GetOkExists("drifter_patch_id"); ok {
		tmp := drifterPatchId.(int)
		request.DrifterPatchId = &tmp
	}

	if imageCompliance, ok := s.D.GetOkExists("image_compliance"); ok {
		request.ImageCompliance = oci_dblm.ListDatabasesImageComplianceEnum(imageCompliance.(string))
	}

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	if severityType, ok := s.D.GetOkExists("severity_type"); ok {
		interfaces := severityType.([]interface{})
		tmp := make([]oci_dblm.ResourcesSeveritiesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_dblm.ResourcesSeveritiesEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("severity_type") {
			request.SeverityType = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dblm.DblmVulnerabilityLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dblm")

	response, err := s.Client.ListDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DblmPatchManagementDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DblmPatchManagementDatabasesDataSource-", DblmPatchManagementDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	dblmPatchManagementDatabase := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabasesSummaryToMap(item))
	}
	dblmPatchManagementDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DblmPatchManagementDatabasesDataSource().Schema["patch_databases_collection"].Elem.(*schema.Resource).Schema)
		dblmPatchManagementDatabase["items"] = items
	}

	resources = append(resources, dblmPatchManagementDatabase)
	if err := s.D.Set("patch_databases_collection", resources); err != nil {
		return err
	}

	return nil
}

func AdditionalPatchesToMap(obj oci_dblm.AdditionalPatches) map[string]interface{} {
	result := map[string]interface{}{}

	result["category"] = string(obj.Category)

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.PatchId != nil {
		result["patch_id"] = int(*obj.PatchId)
	}

	if obj.PatchName != nil {
		result["patch_name"] = string(*obj.PatchName)
	}

	return result
}

func DatabasesSummaryToMap(obj oci_dblm.DatabasesSummary) map[string]interface{} {
	result := map[string]interface{}{}

	additionalPatches := []interface{}{}
	for _, item := range obj.AdditionalPatches {
		additionalPatches = append(additionalPatches, AdditionalPatchesToMap(item))
	}
	result["additional_patches"] = additionalPatches

	if obj.CurrentPatchWatermark != nil {
		result["current_patch_watermark"] = string(*obj.CurrentPatchWatermark)
	}

	if obj.DatabaseId != nil {
		result["database_id"] = string(*obj.DatabaseId)
	}

	if obj.DatabaseName != nil {
		result["database_name"] = string(*obj.DatabaseName)
	}

	result["database_type"] = string(obj.DatabaseType)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostOrCluster != nil {
		result["host_or_cluster"] = string(*obj.HostOrCluster)
	}

	if obj.ImageDetails != nil {
		result["image_details"] = []interface{}{ImageDetailsToMap(obj.ImageDetails)}
	}

	if obj.OracleHomePath != nil {
		result["oracle_home_path"] = string(*obj.OracleHomePath)
	}

	if obj.PatchActivityDetails != nil {
		result["patch_activity_details"] = []interface{}{PatchActivityDetailsToMap(obj.PatchActivityDetails)}
	}

	if obj.PatchComplianceDetails != nil {
		result["patch_compliance_details"] = []interface{}{PatchComplianceDetailsToMap(obj.PatchComplianceDetails)}
	}

	if obj.PatchUser != nil {
		result["patch_user"] = string(*obj.PatchUser)
	}

	if obj.Release != nil {
		result["release"] = string(*obj.Release)
	}

	if obj.ReleaseFullVersion != nil {
		result["release_full_version"] = string(*obj.ReleaseFullVersion)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SudoFilePath != nil {
		result["sudo_file_path"] = string(*obj.SudoFilePath)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.VulnerabilitiesSummary != nil {
		result["vulnerabilities_summary"] = []interface{}{objectToMap(obj.VulnerabilitiesSummary)}
	}

	return result
}

func ImageDetailsToMap(obj *oci_dblm.ImageDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.CurrentVersion != nil {
		result["current_version"] = string(*obj.CurrentVersion)
	}

	if obj.ImageId != nil {
		result["image_id"] = string(*obj.ImageId)
	}

	if obj.ImageOwner != nil {
		result["image_owner"] = string(*obj.ImageOwner)
	}

	result["image_status"] = string(obj.ImageStatus)

	if obj.ImageVersion != nil {
		result["image_version"] = string(*obj.ImageVersion)
	}

	if obj.SubscribedImage != nil {
		result["subscribed_image"] = string(*obj.SubscribedImage)
	}

	if obj.TimeImageCreation != nil {
		result["time_image_creation"] = obj.TimeImageCreation.String()
	}

	if obj.UpToDateImageVersion != nil {
		result["up_to_date_image_version"] = string(*obj.UpToDateImageVersion)
	}

	return result
}

func PatchActivityDetailsToMap(obj *oci_dblm.PatchActivityDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DeployOperationId != nil {
		result["deploy_operation_id"] = string(*obj.DeployOperationId)
	}

	result["deploy_status"] = string(obj.DeployStatus)

	if obj.DeployTaskId != nil {
		result["deploy_task_id"] = string(*obj.DeployTaskId)
	}

	if obj.MigrateListenerOperationId != nil {
		result["migrate_listener_operation_id"] = string(*obj.MigrateListenerOperationId)
	}

	result["migrate_listener_status"] = string(obj.MigrateListenerStatus)

	if obj.MigrateListenerTaskId != nil {
		result["migrate_listener_task_id"] = string(*obj.MigrateListenerTaskId)
	}

	if obj.UpdateOperationId != nil {
		result["update_operation_id"] = string(*obj.UpdateOperationId)
	}

	result["update_status"] = string(obj.UpdateStatus)

	if obj.UpdateTaskId != nil {
		result["update_task_id"] = string(*obj.UpdateTaskId)
	}

	return result
}

func PatchComplianceDetailsToMap(obj *oci_dblm.PatchComplianceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["patch_compliance_status"] = string(obj.PatchComplianceStatus)

	if obj.PatchComplianceVersion != nil {
		result["patch_compliance_version"] = string(*obj.PatchComplianceVersion)
	}

	return result
}
