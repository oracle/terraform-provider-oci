// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func AutonomousDatabaseBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createAutonomousDatabaseBackup,
		Read:     readAutonomousDatabaseBackup,
		Delete:   deleteAutonomousDatabaseBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_automatic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return CreateResource(d, sync)
}

func readAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

func deleteAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	return nil
}

type AutonomousDatabaseBackupResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousDatabaseBackup
	DisableNotFoundRetries bool
}

func (s *AutonomousDatabaseBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AutonomousDatabaseBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateCreating),
	}
}

func (s *AutonomousDatabaseBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateActive),
	}
}

func (s *AutonomousDatabaseBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateDeleting),
	}
}

func (s *AutonomousDatabaseBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateDeleted),
	}
}

func (s *AutonomousDatabaseBackupResourceCrud) Create() error {
	request := oci_database.CreateAutonomousDatabaseBackupRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousDatabaseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseBackup
	return nil
}

func (s *AutonomousDatabaseBackupResourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseBackupRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousDatabaseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseBackup
	return nil
}

func (s *AutonomousDatabaseBackupResourceCrud) SetData() error {
	if s.Res.AutonomousDatabaseId != nil {
		s.D.Set("autonomous_database_id", *s.Res.AutonomousDatabaseId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IsAutomatic != nil {
		s.D.Set("is_automatic", *s.Res.IsAutomatic)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
