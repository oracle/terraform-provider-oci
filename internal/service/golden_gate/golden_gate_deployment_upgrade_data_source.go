// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"
)

func GoldenGateDeploymentUpgradeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularGoldenGateDeploymentUpgrade,
		Schema: map[string]*schema.Schema{
			"deployment_upgrade_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"deployment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deployment_upgrade_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_cancel_allowed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_reschedule_allowed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_rollback_allowed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_security_fix": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_snoozed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_sub_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ogg_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"previous_ogg_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"release_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_ogg_version_supported_until": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_released": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_schedule": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_schedule_max": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_snoozed_until": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularGoldenGateDeploymentUpgrade(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentUpgradeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateDeploymentUpgradeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.GetDeploymentUpgradeResponse
}

func (s *GoldenGateDeploymentUpgradeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDeploymentUpgradeDataSourceCrud) Get() error {
	request := oci_golden_gate.GetDeploymentUpgradeRequest{}

	if deploymentUpgradeId, ok := s.D.GetOkExists("deployment_upgrade_id"); ok {
		tmp := deploymentUpgradeId.(string)
		request.DeploymentUpgradeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.GetDeploymentUpgrade(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGateDeploymentUpgradeDataSourceCrud) SetData() error {
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

	if s.Res.DeploymentId != nil {
		s.D.Set("deployment_id", *s.Res.DeploymentId)
	}

	s.D.Set("deployment_upgrade_type", s.Res.DeploymentUpgradeType)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCancelAllowed != nil {
		s.D.Set("is_cancel_allowed", *s.Res.IsCancelAllowed)
	}

	if s.Res.IsRescheduleAllowed != nil {
		s.D.Set("is_reschedule_allowed", *s.Res.IsRescheduleAllowed)
	}

	if s.Res.IsRollbackAllowed != nil {
		s.D.Set("is_rollback_allowed", *s.Res.IsRollbackAllowed)
	}

	if s.Res.IsSecurityFix != nil {
		s.D.Set("is_security_fix", *s.Res.IsSecurityFix)
	}

	if s.Res.IsSnoozed != nil {
		s.D.Set("is_snoozed", *s.Res.IsSnoozed)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_sub_state", s.Res.LifecycleSubState)

	if s.Res.OggVersion != nil {
		s.D.Set("ogg_version", *s.Res.OggVersion)
	}

	if s.Res.PreviousOggVersion != nil {
		s.D.Set("previous_ogg_version", *s.Res.PreviousOggVersion)
	}

	s.D.Set("release_type", s.Res.ReleaseType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeOggVersionSupportedUntil != nil {
		s.D.Set("time_ogg_version_supported_until", s.Res.TimeOggVersionSupportedUntil.String())
	}

	if s.Res.TimeReleased != nil {
		s.D.Set("time_released", s.Res.TimeReleased.String())
	}

	if s.Res.TimeSchedule != nil {
		s.D.Set("time_schedule", s.Res.TimeSchedule.String())
	}

	if s.Res.TimeScheduleMax != nil {
		s.D.Set("time_schedule_max", s.Res.TimeScheduleMax.String())
	}

	if s.Res.TimeSnoozedUntil != nil {
		s.D.Set("time_snoozed_until", s.Res.TimeSnoozedUntil.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
