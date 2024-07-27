// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreInstanceMaintenanceEventResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreInstanceMaintenanceEvent,
		Read:     readCoreInstanceMaintenanceEvent,
		Update:   updateCoreInstanceMaintenanceEvent,
		Delete:   deleteCoreInstanceMaintenanceEvent,
		Schema: map[string]*schema.Schema{
			// Required
			"instance_maintenance_event_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"alternative_resolution_action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"can_delete_local_storage": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_window_start": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},

			// Computed
			"additional_details": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"alternative_resolution_actions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"can_reschedule": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"correlation_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"estimated_duration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_window_duration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_hard_due_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreInstanceMaintenanceEvent(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceMaintenanceEventResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readCoreInstanceMaintenanceEvent(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceMaintenanceEventResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreInstanceMaintenanceEvent(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceMaintenanceEventResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreInstanceMaintenanceEvent(d *schema.ResourceData, m interface{}) error {
	return nil
}

type CoreInstanceMaintenanceEventResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.InstanceMaintenanceEvent
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *CoreInstanceMaintenanceEventResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreInstanceMaintenanceEventResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.InstanceMaintenanceEventLifecycleStateScheduled),
	}
}

func (s *CoreInstanceMaintenanceEventResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.InstanceMaintenanceEventLifecycleStateSucceeded),
	}
}

func (s *CoreInstanceMaintenanceEventResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *CoreInstanceMaintenanceEventResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *CoreInstanceMaintenanceEventResourceCrud) Create() error {
	request := oci_core.UpdateInstanceMaintenanceEventRequest{}

	if alternativeResolutionAction, ok := s.D.GetOkExists("alternative_resolution_action"); ok {
		request.AlternativeResolutionAction = oci_core.InstanceMaintenanceAlternativeResolutionActionsEnum(alternativeResolutionAction.(string))
	}

	if canDeleteLocalStorage, ok := s.D.GetOkExists("can_delete_local_storage"); ok {
		tmp := canDeleteLocalStorage.(bool)
		request.CanDeleteLocalStorage = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceMaintenanceEventId, ok := s.D.GetOkExists("instance_maintenance_event_id"); ok {
		tmp := instanceMaintenanceEventId.(string)
		request.InstanceMaintenanceEventId = &tmp
	}

	if timeWindowStart, ok := s.D.GetOkExists("time_window_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeWindowStart.(string))
		if err != nil {
			return err
		}
		request.TimeWindowStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInstanceMaintenanceEvent(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	if workId != nil {
		var identifier *string
		var err error
		identifier = tfresource.GetResourceIDFromWorkRequest(s.WorkRequestClient, workId, "instancemaintenanceevent", s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "instancemaintenanceevent", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *CoreInstanceMaintenanceEventResourceCrud) Get() error {
	request := oci_core.GetInstanceMaintenanceEventRequest{}

	tmp := s.D.Id()
	request.InstanceMaintenanceEventId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInstanceMaintenanceEvent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceMaintenanceEvent
	return nil
}

func (s *CoreInstanceMaintenanceEventResourceCrud) Update() error {
	request := oci_core.UpdateInstanceMaintenanceEventRequest{}

	if alternativeResolutionAction, ok := s.D.GetOkExists("alternative_resolution_action"); ok {
		request.AlternativeResolutionAction = oci_core.InstanceMaintenanceAlternativeResolutionActionsEnum(alternativeResolutionAction.(string))
	}

	if canDeleteLocalStorage, ok := s.D.GetOkExists("can_delete_local_storage"); ok {
		tmp := canDeleteLocalStorage.(bool)
		request.CanDeleteLocalStorage = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.InstanceMaintenanceEventId = &tmp

	if timeWindowStart, ok := s.D.GetOkExists("time_window_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeWindowStart.(string))
		if err != nil {
			return err
		}
		request.TimeWindowStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInstanceMaintenanceEvent(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "instanceMaintenanceEvent", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *CoreInstanceMaintenanceEventResourceCrud) SetData() error {
	s.D.Set("additional_details", s.Res.AdditionalDetails)
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	s.D.Set("alternative_resolution_actions", s.Res.AlternativeResolutionActions)
	s.D.Set("alternative_resolution_actions", s.Res.AlternativeResolutionActions)

	if s.Res.CanDeleteLocalStorage != nil {
		s.D.Set("can_delete_local_storage", *s.Res.CanDeleteLocalStorage)
	}

	if s.Res.CanReschedule != nil {
		s.D.Set("can_reschedule", *s.Res.CanReschedule)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CorrelationToken != nil {
		s.D.Set("correlation_token", *s.Res.CorrelationToken)
	}

	s.D.Set("created_by", s.Res.CreatedBy)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EstimatedDuration != nil {
		s.D.Set("estimated_duration", *s.Res.EstimatedDuration)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("instance_action", s.Res.InstanceAction)

	if s.Res.InstanceId != nil {
		s.D.Set("instance_id", *s.Res.InstanceId)
	}

	s.D.Set("maintenance_category", s.Res.MaintenanceCategory)

	s.D.Set("maintenance_reason", s.Res.MaintenanceReason)

	if s.Res.StartWindowDuration != nil {
		s.D.Set("start_window_duration", *s.Res.StartWindowDuration)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeHardDueDate != nil {
		s.D.Set("time_hard_due_date", s.Res.TimeHardDueDate.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeWindowStart != nil {
		s.D.Set("time_window_start", s.Res.TimeWindowStart.Format(time.RFC3339Nano))
	}

	return nil
}
