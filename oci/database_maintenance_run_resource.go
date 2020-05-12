// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"time"

	"github.com/hashicorp/terraform/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func init() {
	RegisterResource("oci_database_maintenance_run", DatabaseMaintenanceRunResource())
}

func DatabaseMaintenanceRunResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatabaseMaintenanceRun,
		Read:     readDatabaseMaintenanceRun,
		Update:   updateDatabaseMaintenanceRun,
		Delete:   deleteDatabaseMaintenanceRun,
		Schema: map[string]*schema.Schema{
			// Required
			"maintenance_run_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"time_scheduled": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: timeDiffSuppressFunction,
			},

			// Computed
			"compartment_id": {
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
			"lifecycle_details": {
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
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_resource_type": {
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
		},
	}
}

func createDatabaseMaintenanceRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMaintenanceRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return CreateResource(d, sync)
}

func readDatabaseMaintenanceRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMaintenanceRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

func updateDatabaseMaintenanceRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMaintenanceRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return UpdateResource(d, sync)
}

func deleteDatabaseMaintenanceRun(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseMaintenanceRunResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.MaintenanceRun
	DisableNotFoundRetries bool
}

func (s *DatabaseMaintenanceRunResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseMaintenanceRunResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.MaintenanceRunLifecycleStateScheduled),
		string(oci_database.MaintenanceRunLifecycleStateInProgress),
	}
}

func (s *DatabaseMaintenanceRunResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.MaintenanceRunLifecycleStateSucceeded),
		string(oci_database.MaintenanceRunLifecycleStateSkipped),
	}
}

func (s *DatabaseMaintenanceRunResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.MaintenanceRunLifecycleStateDeleting),
	}
}

func (s *DatabaseMaintenanceRunResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.MaintenanceRunLifecycleStateDeleted),
	}
}

func (s *DatabaseMaintenanceRunResourceCrud) Create() error {
	request := oci_database.UpdateMaintenanceRunRequest{}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if maintenanceRunId, ok := s.D.GetOkExists("maintenance_run_id"); ok {
		tmp := maintenanceRunId.(string)
		request.MaintenanceRunId = &tmp
	}

	if timeScheduled, ok := s.D.GetOkExists("time_scheduled"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduled.(string))
		if err != nil {
			return err
		}
		request.TimeScheduled = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateMaintenanceRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaintenanceRun
	return nil
}

func (s *DatabaseMaintenanceRunResourceCrud) Get() error {
	request := oci_database.GetMaintenanceRunRequest{}

	tmp := s.D.Id()
	request.MaintenanceRunId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetMaintenanceRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaintenanceRun
	return nil
}

func (s *DatabaseMaintenanceRunResourceCrud) Update() error {
	request := oci_database.UpdateMaintenanceRunRequest{}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	tmp := s.D.Id()
	request.MaintenanceRunId = &tmp

	if timeScheduled, ok := s.D.GetOkExists("time_scheduled"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduled.(string))
		if err != nil {
			return err
		}
		request.TimeScheduled = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateMaintenanceRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaintenanceRun
	return nil
}

func (s *DatabaseMaintenanceRunResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("maintenance_subtype", s.Res.MaintenanceSubtype)

	s.D.Set("maintenance_type", s.Res.MaintenanceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetResourceId != nil {
		s.D.Set("target_resource_id", *s.Res.TargetResourceId)
	}

	s.D.Set("target_resource_type", s.Res.TargetResourceType)

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeScheduled != nil {
		s.D.Set("time_scheduled", s.Res.TimeScheduled.Format(time.RFC3339Nano))
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}
