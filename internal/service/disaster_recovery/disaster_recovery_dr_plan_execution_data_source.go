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

func DisasterRecoveryDrPlanExecutionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["dr_plan_execution_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DisasterRecoveryDrPlanExecutionResource(), fieldMap, readSingularDisasterRecoveryDrPlanExecution)
}

func readSingularDisasterRecoveryDrPlanExecution(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrPlanExecutionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.ReadResource(sync)
}

type DisasterRecoveryDrPlanExecutionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_disaster_recovery.DisasterRecoveryClient
	Res    *oci_disaster_recovery.GetDrPlanExecutionResponse
}

func (s *DisasterRecoveryDrPlanExecutionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DisasterRecoveryDrPlanExecutionDataSourceCrud) Get() error {
	request := oci_disaster_recovery.GetDrPlanExecutionRequest{}

	if drPlanExecutionId, ok := s.D.GetOkExists("dr_plan_execution_id"); ok {
		tmp := drPlanExecutionId.(string)
		request.DrPlanExecutionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "disaster_recovery")

	response, err := s.Client.GetDrPlanExecution(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DisasterRecoveryDrPlanExecutionDataSourceCrud) SetData() error {
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

	if s.Res.ExecutionDurationInSec != nil {
		s.D.Set("execution_duration_in_sec", *s.Res.ExecutionDurationInSec)
	}

	if s.Res.ExecutionOptions != nil {
		executionOptionsArray := []interface{}{}
		if executionOptionsMap := DrPlanExecutionOptionsToMap(&s.Res.ExecutionOptions); executionOptionsMap != nil {
			executionOptionsArray = append(executionOptionsArray, executionOptionsMap)
		}
		s.D.Set("execution_options", executionOptionsArray)
	} else {
		s.D.Set("execution_options", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	groupExecutions := []interface{}{}
	for _, item := range s.Res.GroupExecutions {
		groupExecutions = append(groupExecutions, DrPlanGroupExecutionToMap(item))
	}
	s.D.Set("group_executions", groupExecutions)

	if s.Res.LifeCycleDetails != nil {
		s.D.Set("life_cycle_details", *s.Res.LifeCycleDetails)
	}

	if s.Res.LogLocation != nil {
		s.D.Set("log_location", []interface{}{ObjectStorageLogLocationToMap(s.Res.LogLocation)})
	} else {
		s.D.Set("log_location", nil)
	}

	if s.Res.PeerDrProtectionGroupId != nil {
		s.D.Set("peer_dr_protection_group_id", *s.Res.PeerDrProtectionGroupId)
	}

	if s.Res.PeerRegion != nil {
		s.D.Set("peer_region", *s.Res.PeerRegion)
	}

	s.D.Set("plan_execution_type", s.Res.PlanExecutionType)

	if s.Res.PlanId != nil {
		s.D.Set("plan_id", *s.Res.PlanId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
