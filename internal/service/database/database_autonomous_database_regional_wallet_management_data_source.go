// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseAutonomousDatabaseRegionalWalletManagementDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	return tfresource.GetSingularDataSourceItemSchema(DatabaseAutonomousDatabaseRegionalWalletManagementResource(), fieldMap, readSingularDatabaseAutonomousDatabaseRegionalWalletManagement)
}

func readSingularDatabaseAutonomousDatabaseRegionalWalletManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseRegionalWalletManagementDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabaseRegionalWalletManagementDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.AutonomousDatabaseWallet
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseRegionalWalletRequest{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousDatabaseRegionalWallet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseWallet
	return nil
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseRegionalWalletManagementDataSource-", DatabaseAutonomousDatabaseRegionalWalletManagementDataSource(), s.D))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeRotated != nil {
		s.D.Set("time_rotated", s.Res.TimeRotated.String())
	}

	return nil
}
