// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementPatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementPatches,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"patch_type_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"should_compliance_policy_rules_be_applied": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_released_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_released_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"patch_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FleetAppsManagementPatchResource()),
						},
					},
				},
			},
		},
	}
}

func readFleetAppsManagementPatches(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementPatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.ListPatchesResponse
}

func (s *FleetAppsManagementPatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementPatchesDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListPatchesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if patchTypeId, ok := s.D.GetOkExists("patch_type_id"); ok {
		tmp := patchTypeId.(string)
		request.PatchTypeId = &tmp
	}

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		request.ProductId = &tmp
	}

	if shouldCompliancePolicyRulesBeApplied, ok := s.D.GetOkExists("should_compliance_policy_rules_be_applied"); ok {
		tmp := shouldCompliancePolicyRulesBeApplied.(bool)
		request.ShouldCompliancePolicyRulesBeApplied = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_apps_management.PatchLifecycleStateEnum(state.(string))
	}

	if timeReleasedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_released_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeReleasedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeReleasedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeReleasedLessThan, ok := s.D.GetOkExists("time_released_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeReleasedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeReleasedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_fleet_apps_management.PatchTypeEnum(type_.(string))
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListPatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementPatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementPatchesDataSource-", FleetAppsManagementPatchesDataSource(), s.D))
	resources := []map[string]interface{}{}
	patch := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PatchSummaryToMap(item))
	}
	patch["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementPatchesDataSource().Schema["patch_collection"].Elem.(*schema.Resource).Schema)
		patch["items"] = items
	}

	resources = append(resources, patch)
	if err := s.D.Set("patch_collection", resources); err != nil {
		return err
	}

	return nil
}
