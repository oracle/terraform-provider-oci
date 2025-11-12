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

func FleetAppsManagementTargetPropertiesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readFleetAppsManagementTargetPropertiesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"target_property_collection": {
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
		},
	}
}

func readFleetAppsManagementTargetPropertiesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetAppsManagementTargetPropertiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FleetAppsManagementTargetPropertiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.ListTargetPropertiesResponse
}

func (s *FleetAppsManagementTargetPropertiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementTargetPropertiesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_fleet_apps_management.ListTargetPropertiesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_fleet_apps_management.ListTargetPropertiesSeverityEnum(severity.(string))
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

	response, err := s.Client.ListTargetProperties(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTargetProperties(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementTargetPropertiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementTargetPropertiesDataSource-", FleetAppsManagementTargetPropertiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	targetProperty := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TargetPropertySummaryToMap(item))
	}
	targetProperty["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementTargetPropertiesDataSource().Schema["target_property_collection"].Elem.(*schema.Resource).Schema)
		targetProperty["items"] = items
	}

	resources = append(resources, targetProperty)
	if err := s.D.Set("target_property_collection", resources); err != nil {
		return err
	}

	return nil
}

// func TargetPropertySummaryToMap(obj oci_fleet_apps_management.TargetPropertySummary) map[string]interface{} {
// 	result := map[string]interface{}{}

// 	if obj.Name != nil {
// 		result["name"] = string(*obj.Name)
// 	}

// 	if obj.Value != nil {
// 		result["value"] = string(*obj.Value)
// 	}

// 	return result
// }
