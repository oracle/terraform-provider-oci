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

func FleetAppsManagementCompliancePolicyRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementCompliancePolicyRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"patch_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_policy_rule_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FleetAppsManagementCompliancePolicyRuleResource()),
						},
					},
				},
			},
		},
	}
}

func readFleetAppsManagementCompliancePolicyRules(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCompliancePolicyRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementCompliancePolicyRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementAdminClient
	Res    *oci_fleet_apps_management.ListCompliancePolicyRulesResponse
}

func (s *FleetAppsManagementCompliancePolicyRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementCompliancePolicyRulesDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListCompliancePolicyRulesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compliancePolicyId, ok := s.D.GetOkExists("compliance_policy_id"); ok {
		tmp := compliancePolicyId.(string)
		request.CompliancePolicyId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if patchName, ok := s.D.GetOkExists("patch_name"); ok {
		tmp := patchName.(string)
		request.PatchName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_apps_management.CompliancePolicyRuleLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListCompliancePolicyRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCompliancePolicyRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementCompliancePolicyRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementCompliancePolicyRulesDataSource-", FleetAppsManagementCompliancePolicyRulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	compliancePolicyRule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CompliancePolicyRuleSummaryToMap(item))
	}
	compliancePolicyRule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementCompliancePolicyRulesDataSource().Schema["compliance_policy_rule_collection"].Elem.(*schema.Resource).Schema)
		compliancePolicyRule["items"] = items
	}

	resources = append(resources, compliancePolicyRule)
	if err := s.D.Set("compliance_policy_rule_collection", resources); err != nil {
		return err
	}

	return nil
}
