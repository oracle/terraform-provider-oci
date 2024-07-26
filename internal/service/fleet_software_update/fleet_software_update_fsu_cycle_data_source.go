// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_software_update

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetSoftwareUpdateFsuCycleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fsu_cycle_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetSoftwareUpdateFsuCycleResource(), fieldMap, readSingularFleetSoftwareUpdateFsuCycle)
}

func readSingularFleetSoftwareUpdateFsuCycle(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCycleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.ReadResource(sync)
}

type FleetSoftwareUpdateFsuCycleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_software_update.FleetSoftwareUpdateClient
	Res    *oci_fleet_software_update.GetFsuCycleResponse
}

func (s *FleetSoftwareUpdateFsuCycleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetSoftwareUpdateFsuCycleDataSourceCrud) Get() error {
	request := oci_fleet_software_update.GetFsuCycleRequest{}

	if fsuCycleId, ok := s.D.GetOkExists("fsu_cycle_id"); ok {
		tmp := fsuCycleId.(string)
		request.FsuCycleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_software_update")

	response, err := s.Client.GetFsuCycle(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetSoftwareUpdateFsuCycleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.FsuCycle).(type) {
	case oci_fleet_software_update.PatchFsuCycle:
		s.D.Set("type", "PATCH")

		s.D.Set("is_ignore_missing_patches", v.IsIgnoreMissingPatches)

		if v.IsIgnorePatches != nil {
			s.D.Set("is_ignore_patches", *v.IsIgnorePatches)
		}

		if v.IsKeepPlacement != nil {
			s.D.Set("is_keep_placement", *v.IsKeepPlacement)
		}

		if v.MaxDrainTimeoutInSeconds != nil {
			s.D.Set("max_drain_timeout_in_seconds", *v.MaxDrainTimeoutInSeconds)
		}

		if v.ApplyActionSchedule != nil {
			applyActionScheduleArray := []interface{}{}
			if applyActionScheduleMap := ScheduleDetailsToMap(&v.ApplyActionSchedule); applyActionScheduleMap != nil {
				applyActionScheduleArray = append(applyActionScheduleArray, applyActionScheduleMap)
			}
			s.D.Set("apply_action_schedule", applyActionScheduleArray)
		} else {
			s.D.Set("apply_action_schedule", nil)
		}

		if v.BatchingStrategy != nil {
			batchingStrategyArray := []interface{}{}
			if batchingStrategyMap := BatchingStrategyDetailsToMap(&v.BatchingStrategy); batchingStrategyMap != nil {
				batchingStrategyArray = append(batchingStrategyArray, batchingStrategyMap)
			}
			s.D.Set("batching_strategy", batchingStrategyArray)
		} else {
			s.D.Set("batching_strategy", nil)
		}

		s.D.Set("collection_type", v.CollectionType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DiagnosticsCollection != nil {
			s.D.Set("diagnostics_collection", []interface{}{DiagnosticsCollectionDetailsToMap(v.DiagnosticsCollection)})
		} else {
			s.D.Set("diagnostics_collection", nil)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.ExecutingFsuActionId != nil {
			s.D.Set("executing_fsu_action_id", *v.ExecutingFsuActionId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.FsuCollectionId != nil {
			s.D.Set("fsu_collection_id", *v.FsuCollectionId)
		}

		if v.GoalVersionDetails != nil {
			goalVersionDetailsArray := []interface{}{}
			if goalVersionDetailsMap := FsuGoalVersionDetailsToMap(&v.GoalVersionDetails); goalVersionDetailsMap != nil {
				goalVersionDetailsArray = append(goalVersionDetailsArray, goalVersionDetailsMap)
			}
			s.D.Set("goal_version_details", goalVersionDetailsArray)
		} else {
			s.D.Set("goal_version_details", nil)
		}

		s.D.Set("last_completed_action", v.LastCompletedAction)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		nextActionToExecute := []interface{}{}
		for _, item := range v.NextActionToExecute {
			nextActionToExecute = append(nextActionToExecute, NextActionToExecuteDetailsToMap(item))
		}
		s.D.Set("next_action_to_execute", nextActionToExecute)

		if v.StageActionSchedule != nil {
			stageActionScheduleArray := []interface{}{}
			if stageActionScheduleMap := ScheduleDetailsToMap(&v.StageActionSchedule); stageActionScheduleMap != nil {
				stageActionScheduleArray = append(stageActionScheduleArray, stageActionScheduleMap)
			}
			s.D.Set("stage_action_schedule", stageActionScheduleArray)
		} else {
			s.D.Set("stage_action_schedule", nil)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeFinished != nil {
			s.D.Set("time_finished", v.TimeFinished.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.FsuCycle)
		return nil
	}

	return nil
}
