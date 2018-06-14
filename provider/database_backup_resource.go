// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func BackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createBackup,
		Read:     readBackup,
		Delete:   deleteBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"database_id": {
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
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_edition": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_data_size_in_mbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
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

func createBackup(d *schema.ResourceData, m interface{}) error {
	sync := &BackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.CreateResource(d, sync)
}

func readBackup(d *schema.ResourceData, m interface{}) error {
	sync := &BackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

func deleteBackup(d *schema.ResourceData, m interface{}) error {
	sync := &BackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type BackupResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.Backup
	DisableNotFoundRetries bool
}

func (s *BackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.BackupLifecycleStateCreating),
		string(oci_database.BackupLifecycleStateRestoring),
	}
}

func (s *BackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.BackupLifecycleStateActive),
	}
}

func (s *BackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.BackupLifecycleStateDeleting),
	}
}

func (s *BackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.BackupLifecycleStateDeleted),
	}
}

func (s *BackupResourceCrud) Create() error {
	request := oci_database.CreateBackupRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup
	return nil
}

func (s *BackupResourceCrud) Get() error {
	request := oci_database.GetBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup
	return nil
}

func (s *BackupResourceCrud) Delete() error {
	request := oci_database.DeleteBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteBackup(context.Background(), request)
	return err
}

func (s *BackupResourceCrud) SetData() {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseEdition != nil {
		s.D.Set("database_edition", *s.Res.DatabaseEdition)
	}

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
	}

	if s.Res.DbDataSizeInMBs != nil {
		s.D.Set("db_data_size_in_mbs", *s.Res.DbDataSizeInMBs)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
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

}
