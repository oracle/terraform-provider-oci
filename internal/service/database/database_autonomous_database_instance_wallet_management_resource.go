// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseAutonomousDatabaseInstanceWalletManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("20m"),
			Update: tfresource.GetTimeoutDuration("20m"),
			Delete: tfresource.GetTimeoutDuration("20m"),
		},
		Create: createDatabaseAutonomousDatabaseInstanceWalletManagement,
		Read:   readDatabaseAutonomousDatabaseInstanceWalletManagement,
		Update: updateDatabaseAutonomousDatabaseInstanceWalletManagement,
		Delete: deleteDatabaseAutonomousDatabaseInstanceWalletManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"should_rotate": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_rotated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseAutonomousDatabaseInstanceWalletManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousDatabaseInstanceWalletManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseAutonomousDatabaseInstanceWalletManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseAutonomousDatabaseInstanceWalletManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousDatabaseWallet
	DisableNotFoundRetries bool
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud) ID() string {
	return s.D.Get("autonomous_database_id").(string)
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseWalletLifecycleStateUpdating),
	}
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseWalletLifecycleStateActive),
	}
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseWalletLifecycleStateUpdating),
	}
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseWalletLifecycleStateActive),
	}
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud) Create() error {
	if shouldRotate, ok := s.D.GetOkExists("should_rotate"); ok {
		if tmp := shouldRotate.(bool); tmp {
			return s.Update()
		}
	}
	return s.Get()
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseWalletRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousDatabaseWallet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseWallet
	return nil
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud) Update() error {
	request := oci_database.UpdateAutonomousDatabaseWalletRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if shouldRotate, ok := s.D.GetOkExists("should_rotate"); ok {
		tmp := shouldRotate.(bool)
		request.ShouldRotate = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.UpdateAutonomousDatabaseWallet(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}
	return s.Get()
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementResourceCrud) SetData() error {
	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeRotated != nil {
		s.D.Set("time_rotated", s.Res.TimeRotated.String())
	}

	return nil
}
