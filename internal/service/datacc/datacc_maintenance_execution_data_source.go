// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datacc "github.com/oracle/oci-go-sdk/v65/datacc"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataccMaintenanceExecutionDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularDataccMaintenanceExecutionWithContext,
		Schema: map[string]*schema.Schema{
			"maintenance_execution_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_action_timeout_in_mins": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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
			"infrastructure_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_custom_action_timeout_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_subtype": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patching_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_version": {
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
			"target_resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_ended": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_time_taken_in_mins": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"workflow_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDataccMaintenanceExecutionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccMaintenanceExecutionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataccMaintenanceExecutionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacc.BaseinfraClient
	Res    *oci_datacc.GetMaintenanceExecutionResponse
}

func (s *DataccMaintenanceExecutionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataccMaintenanceExecutionDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datacc.GetMaintenanceExecutionRequest{}

	if maintenanceExecutionId, ok := s.D.GetOkExists("maintenance_execution_id"); ok {
		tmp := maintenanceExecutionId.(string)
		request.MaintenanceExecutionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacc")

	response, err := s.Client.GetMaintenanceExecution(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataccMaintenanceExecutionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CustomActionTimeoutInMins != nil {
		s.D.Set("custom_action_timeout_in_mins", *s.Res.CustomActionTimeoutInMins)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InfrastructureId != nil {
		s.D.Set("infrastructure_id", *s.Res.InfrastructureId)
	}

	if s.Res.IsCustomActionTimeoutEnabled != nil {
		s.D.Set("is_custom_action_timeout_enabled", *s.Res.IsCustomActionTimeoutEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaintenanceRunId != nil {
		s.D.Set("maintenance_run_id", *s.Res.MaintenanceRunId)
	}

	s.D.Set("maintenance_subtype", s.Res.MaintenanceSubtype)

	s.D.Set("maintenance_type", s.Res.MaintenanceType)

	s.D.Set("patching_mode", s.Res.PatchingMode)

	if s.Res.SourceVersion != nil {
		s.D.Set("source_version", *s.Res.SourceVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("target_resource_type", s.Res.TargetResourceType)

	if s.Res.TargetVersion != nil {
		s.D.Set("target_version", *s.Res.TargetVersion)
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TotalTimeTakenInMins != nil {
		s.D.Set("total_time_taken_in_mins", *s.Res.TotalTimeTakenInMins)
	}

	if s.Res.WorkflowId != nil {
		s.D.Set("workflow_id", *s.Res.WorkflowId)
	}

	return nil
}
