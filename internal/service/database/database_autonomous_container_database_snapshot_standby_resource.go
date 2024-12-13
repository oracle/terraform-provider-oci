// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousContainerDatabaseSnapshotStandbyResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("2h"),
			Update: tfresource.GetTimeoutDuration("2h"),
			Delete: tfresource.GetTimeoutDuration("2h"),
		},
		Create: createDatabaseAutonomousContainerDatabaseSnapshotStandby,
		Read:   readDatabaseAutonomousContainerDatabaseSnapshotStandby,
		Delete: deleteDatabaseAutonomousContainerDatabaseSnapshotStandby,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"connection_strings_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createDatabaseAutonomousContainerDatabaseSnapshotStandby(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousContainerDatabaseSnapshotStandby(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseAutonomousContainerDatabaseSnapshotStandby(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousContainerDatabase
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateProvisioning),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateRestoring),
	}
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminating),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateUnavailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) Create() error {
	request := oci_database.ConvertStandbyAutonomousContainerDatabaseRequest{}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	if connectionStringsType, ok := s.D.GetOkExists("connection_strings_type"); ok {
		request.ConnectionStringsType = oci_database.ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum(connectionStringsType.(string))
	}

	if role, ok := s.D.GetOkExists("role"); ok {
		request.Role = oci_database.ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum(role.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ConvertStandbyAutonomousContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.AutonomousContainerDatabase

	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomouscontainerdatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) SetData() error {
	return nil
}
