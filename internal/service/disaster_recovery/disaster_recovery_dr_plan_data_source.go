// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package disaster_recovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DisasterRecoveryDrPlanDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["dr_plan_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DisasterRecoveryDrPlanResource(), fieldMap, readSingularDisasterRecoveryDrPlan)
}

func readSingularDisasterRecoveryDrPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrPlanDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.ReadResource(sync)
}

type DisasterRecoveryDrPlanDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_disaster_recovery.DisasterRecoveryClient
	Res    *oci_disaster_recovery.GetDrPlanResponse
}

func (s *DisasterRecoveryDrPlanDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DisasterRecoveryDrPlanDataSourceCrud) Get() error {
	request := oci_disaster_recovery.GetDrPlanRequest{}

	if drPlanId, ok := s.D.GetOkExists("dr_plan_id"); ok {
		tmp := drPlanId.(string)
		request.DrPlanId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "disaster_recovery")

	response, err := s.Client.GetDrPlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DisasterRecoveryDrPlanDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DrProtectionGroupId != nil {
		s.D.Set("dr_protection_group_id", *s.Res.DrProtectionGroupId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifeCycleDetails != nil {
		s.D.Set("life_cycle_details", *s.Res.LifeCycleDetails)
	}

	if s.Res.PeerDrProtectionGroupId != nil {
		s.D.Set("peer_dr_protection_group_id", *s.Res.PeerDrProtectionGroupId)
	}

	if s.Res.PeerRegion != nil {
		s.D.Set("peer_region", *s.Res.PeerRegion)
	}

	planGroups := []interface{}{}
	for _, item := range s.Res.PlanGroups {
		planGroups = append(planGroups, DrPlanGroupToMap(item))
	}
	s.D.Set("plan_groups", planGroups)

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

	s.D.Set("type", s.Res.Type)

	return nil
}
