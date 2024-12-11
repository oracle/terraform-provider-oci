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

func FleetAppsManagementCompliancePolicyRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compliance_policy_rule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetAppsManagementCompliancePolicyRuleResource(), fieldMap, readSingularFleetAppsManagementCompliancePolicyRule)
}

func readSingularFleetAppsManagementCompliancePolicyRule(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCompliancePolicyRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementCompliancePolicyRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementAdminClient
	Res    *oci_fleet_apps_management.GetCompliancePolicyRuleResponse
}

func (s *FleetAppsManagementCompliancePolicyRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementCompliancePolicyRuleDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetCompliancePolicyRuleRequest{}

	if compliancePolicyRuleId, ok := s.D.GetOkExists("compliance_policy_rule_id"); ok {
		tmp := compliancePolicyRuleId.(string)
		request.CompliancePolicyRuleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetCompliancePolicyRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementCompliancePolicyRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CompliancePolicyId != nil {
		s.D.Set("compliance_policy_id", *s.Res.CompliancePolicyId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GracePeriod != nil {
		s.D.Set("grace_period", *s.Res.GracePeriod)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PatchSelection != nil {
		patchSelectionArray := []interface{}{}
		if patchSelectionMap := PatchSelectionDetailsToMap(&s.Res.PatchSelection); patchSelectionMap != nil {
			patchSelectionArray = append(patchSelectionArray, patchSelectionMap)
		}
		s.D.Set("patch_selection", patchSelectionArray)
	} else {
		s.D.Set("patch_selection", nil)
	}

	s.D.Set("patch_type", s.Res.PatchType)

	if s.Res.ProductVersion != nil {
		s.D.Set("product_version", []interface{}{ProductVersionDetailsToMap(s.Res.ProductVersion)})
	} else {
		s.D.Set("product_version", nil)
	}

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
