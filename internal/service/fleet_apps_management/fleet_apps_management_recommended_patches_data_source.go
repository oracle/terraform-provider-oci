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

func FleetAppsManagementRecommendedPatchesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readFleetAppsManagementRecommendedPatchesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patch_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"patch_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"patch_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recommended_patch_collection": {
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
					},
				},
			},
		},
	}
}

func readFleetAppsManagementRecommendedPatchesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetAppsManagementRecommendedPatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FleetAppsManagementRecommendedPatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.ListRecommendedPatchesResponse
}

func (s *FleetAppsManagementRecommendedPatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementRecommendedPatchesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_fleet_apps_management.ListRecommendedPatchesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if patchId, ok := s.D.GetOkExists("patch_id"); ok {
		tmp := patchId.(string)
		request.PatchId = &tmp
	}

	if patchLevel, ok := s.D.GetOkExists("patch_level"); ok {
		request.PatchLevel = oci_fleet_apps_management.ListRecommendedPatchesPatchLevelEnum(patchLevel.(string))
	}

	if patchType, ok := s.D.GetOkExists("patch_type"); ok {
		tmp := patchType.(string)
		request.PatchType = &tmp
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_fleet_apps_management.ListRecommendedPatchesSeverityEnum(severity.(string))
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if targetName, ok := s.D.GetOkExists("target_name"); ok {
		tmp := targetName.(string)
		request.TargetName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListRecommendedPatches(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRecommendedPatches(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementRecommendedPatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementRecommendedPatchesDataSource-", FleetAppsManagementRecommendedPatchesDataSource(), s.D))
	resources := []map[string]interface{}{}
	recommendedPatch := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RecommendedPatchSummaryToMap(item))
	}
	recommendedPatch["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementRecommendedPatchesDataSource().Schema["recommended_patch_collection"].Elem.(*schema.Resource).Schema)
		recommendedPatch["items"] = items
	}

	resources = append(resources, recommendedPatch)
	if err := s.D.Set("recommended_patch_collection", resources); err != nil {
		return err
	}

	return nil
}

func RecommendedPatchSummaryToMap(obj oci_fleet_apps_management.RecommendedPatchSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PatchDescription != nil {
		result["patch_description"] = string(*obj.PatchDescription)
	}

	if obj.PatchId != nil {
		result["patch_id"] = string(*obj.PatchId)
	}

	result["patch_level"] = string(obj.PatchLevel)

	if obj.PatchName != nil {
		result["patch_name"] = string(*obj.PatchName)
	}

	if obj.PatchType != nil {
		result["patch_type"] = string(*obj.PatchType)
	}

	result["severity"] = string(obj.Severity)

	if obj.TimeReleased != nil {
		result["time_released"] = obj.TimeReleased.String()
	}

	return result
}
