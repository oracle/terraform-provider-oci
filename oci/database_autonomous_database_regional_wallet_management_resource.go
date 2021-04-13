// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v39/database"
)

func init() {
	RegisterResource("oci_database_autonomous_database_regional_wallet_management", DatabaseAutonomousDatabaseRegionalWalletManagementResource())
}

func DatabaseAutonomousDatabaseRegionalWalletManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: getTimeoutDuration("2h"),
			Update: getTimeoutDuration("2h"),
			Delete: getTimeoutDuration("2h"),
		},
		Create: createDatabaseAutonomousDatabaseRegionalWalletManagement,
		Read:   readDatabaseAutonomousDatabaseRegionalWalletManagement,
		Update: updateDatabaseAutonomousDatabaseRegionalWalletManagement,
		Delete: deleteDatabaseAutonomousDatabaseRegionalWalletManagement,
		Schema: map[string]*schema.Schema{
			// Required

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

func createDatabaseAutonomousDatabaseRegionalWalletManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.Configuration = m.(*OracleClients).configuration

	return CreateResource(d, sync)
}

func readDatabaseAutonomousDatabaseRegionalWalletManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.Configuration = m.(*OracleClients).configuration

	return ReadResource(sync)
}

func updateDatabaseAutonomousDatabaseRegionalWalletManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.Configuration = m.(*OracleClients).configuration

	return UpdateResource(d, sync)
}

func deleteDatabaseAutonomousDatabaseRegionalWalletManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousDatabaseWallet
	DisableNotFoundRetries bool
	Configuration          map[string]string
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud) ID() string {
	return "/autonomousDatabaseRegionalWallet/" + s.Configuration["region"]
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseWalletLifecycleStateActive),
	}
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseWalletLifecycleStateUpdating),
	}
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseWalletLifecycleStateActive),
	}
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud) Create() error {
	if shouldRotate, ok := s.D.GetOkExists("should_rotate"); ok {
		if tmp := shouldRotate.(bool); tmp {
			return s.Update()
		}
	}
	return s.Get()
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseRegionalWalletRequest{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousDatabaseRegionalWallet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseWallet
	return nil
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud) Update() error {
	request := oci_database.UpdateAutonomousDatabaseRegionalWalletRequest{}

	if shouldRotate, ok := s.D.GetOkExists("should_rotate"); ok {
		tmp := shouldRotate.(bool)
		request.ShouldRotate = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.UpdateAutonomousDatabaseRegionalWallet(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := waitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}
	return s.Get()
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementResourceCrud) SetData() error {

	if s.Res == nil {
		return nil
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeRotated != nil {
		s.D.Set("time_rotated", s.Res.TimeRotated.String())
	}

	return nil
}
