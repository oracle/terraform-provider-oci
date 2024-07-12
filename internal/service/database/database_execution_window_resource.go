// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseExecutionWindowResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExecutionWindow,
		Read:     readDatabaseExecutionWindow,
		Update:   updateDatabaseExecutionWindow,
		Delete:   deleteDatabaseExecutionWindow,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"execution_resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_scheduled": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"window_duration_in_mins": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_enforced_duration": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"estimated_time_in_mins": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_substate": {
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
			"time_ended": {
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
			"total_time_taken_in_mins": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"window_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseExecutionWindow(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseExecutionWindow(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseExecutionWindow(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseExecutionWindow(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExecutionWindowResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ExecutionWindow
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExecutionWindowResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseExecutionWindowResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExecutionWindowLifecycleStateScheduled),
		string(oci_database.ExecutionWindowLifecycleStateInProgress),
		string(oci_database.ExecutionWindowLifecycleStateCreating),
	}
}

func (s *DatabaseExecutionWindowResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExecutionWindowLifecycleStateCreated),
		string(oci_database.ExecutionWindowLifecycleStateSucceeded),
	}
}

func (s *DatabaseExecutionWindowResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExecutionWindowLifecycleStateDeleting),
	}
}

func (s *DatabaseExecutionWindowResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExecutionWindowLifecycleStateDeleted),
	}
}

func (s *DatabaseExecutionWindowResourceCrud) Create() error {
	request := oci_database.CreateExecutionWindowRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if executionResourceId, ok := s.D.GetOkExists("execution_resource_id"); ok {
		tmp := executionResourceId.(string)
		request.ExecutionResourceId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnforcedDuration, ok := s.D.GetOkExists("is_enforced_duration"); ok {
		tmp := isEnforcedDuration.(bool)
		request.IsEnforcedDuration = &tmp
	}

	if timeScheduled, ok := s.D.GetOkExists("time_scheduled"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduled.(string))
		if err != nil {
			return err
		}
		request.TimeScheduled = &oci_common.SDKTime{Time: tmp}
	}

	if windowDurationInMins, ok := s.D.GetOkExists("window_duration_in_mins"); ok {
		tmp := windowDurationInMins.(int)
		request.WindowDurationInMins = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateExecutionWindow(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.ExecutionWindow

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "executionwindow", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseExecutionWindowResourceCrud) Get() error {
	request := oci_database.GetExecutionWindowRequest{}

	tmp := s.D.Id()
	request.ExecutionWindowId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExecutionWindow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExecutionWindow
	return nil
}

func (s *DatabaseExecutionWindowResourceCrud) Update() error {
	request := oci_database.UpdateExecutionWindowRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.ExecutionWindowId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnforcedDuration, ok := s.D.GetOkExists("is_enforced_duration"); ok {
		tmp := isEnforcedDuration.(bool)
		request.IsEnforcedDuration = &tmp
	}

	if timeScheduled, ok := s.D.GetOkExists("time_scheduled"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduled.(string))
		if err != nil {
			return err
		}
		request.TimeScheduled = &oci_common.SDKTime{Time: tmp}
	}

	if windowDurationInMins, ok := s.D.GetOkExists("window_duration_in_mins"); ok {
		tmp := windowDurationInMins.(int)
		request.WindowDurationInMins = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateExecutionWindow(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "executionwindow", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseExecutionWindowResourceCrud) Delete() error {
	request := oci_database.DeleteExecutionWindowRequest{}

	tmp := s.D.Id()
	request.ExecutionWindowId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteExecutionWindow(context.Background(), request)
	return err
}

func (s *DatabaseExecutionWindowResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.EstimatedTimeInMins != nil {
		s.D.Set("estimated_time_in_mins", *s.Res.EstimatedTimeInMins)
	}

	if s.Res.ExecutionResourceId != nil {
		s.D.Set("execution_resource_id", *s.Res.ExecutionResourceId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnforcedDuration != nil {
		s.D.Set("is_enforced_duration", *s.Res.IsEnforcedDuration)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_substate", s.Res.LifecycleSubstate)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeScheduled != nil {
		s.D.Set("time_scheduled", s.Res.TimeScheduled.Format(time.RFC3339Nano))
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalTimeTakenInMins != nil {
		s.D.Set("total_time_taken_in_mins", *s.Res.TotalTimeTakenInMins)
	}

	if s.Res.WindowDurationInMins != nil {
		s.D.Set("window_duration_in_mins", *s.Res.WindowDurationInMins)
	}

	s.D.Set("window_type", s.Res.WindowType)

	return nil
}
