// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package disaster_recovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DisasterRecoveryAutomaticDrConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["automatic_dr_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DisasterRecoveryAutomaticDrConfigurationResource(), fieldMap, readSingularDisasterRecoveryAutomaticDrConfigurationWithContext)
}

func readSingularDisasterRecoveryAutomaticDrConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DisasterRecoveryAutomaticDrConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DisasterRecoveryAutomaticDrConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_disaster_recovery.DisasterRecoveryClient
	Res    *oci_disaster_recovery.GetAutomaticDrConfigurationResponse
}

func (s *DisasterRecoveryAutomaticDrConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DisasterRecoveryAutomaticDrConfigurationDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_disaster_recovery.GetAutomaticDrConfigurationRequest{}

	if automaticDrConfigurationId, ok := s.D.GetOkExists("automatic_dr_configuration_id"); ok {
		tmp := automaticDrConfigurationId.(string)
		request.AutomaticDrConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "disaster_recovery")

	response, err := s.Client.GetAutomaticDrConfiguration(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DisasterRecoveryAutomaticDrConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefaultFailoverDrPlanId != nil {
		s.D.Set("default_failover_dr_plan_id", *s.Res.DefaultFailoverDrPlanId)
	}

	if s.Res.DefaultSwitchoverDrPlanId != nil {
		s.D.Set("default_switchover_dr_plan_id", *s.Res.DefaultSwitchoverDrPlanId)
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

	if s.Res.LastAutomaticDrExecutionSubmitDetails != nil {
		s.D.Set("last_automatic_dr_execution_submit_details", *s.Res.LastAutomaticDrExecutionSubmitDetails)
	}

	s.D.Set("last_automatic_dr_execution_submit_status", s.Res.LastAutomaticDrExecutionSubmitStatus)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	members := []interface{}{}
	for _, item := range s.Res.Members {
		members = append(members, AutomaticDrConfigurationMemberToMap(item))
	}
	s.D.Set("members", members)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastAutomaticDrExecutionSubmitAttempt != nil {
		s.D.Set("time_last_automatic_dr_execution_submit_attempt", s.Res.TimeLastAutomaticDrExecutionSubmitAttempt.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
