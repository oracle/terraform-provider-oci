// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v29/database"
)

func init() {
	RegisterDatasource("oci_database_autonomous_database_instance_wallet_management", DatabaseAutonomousDatabaseInstanceWalletManagementDataSource())
}

func DatabaseAutonomousDatabaseInstanceWalletManagementDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["autonomous_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DatabaseAutonomousDatabaseInstanceWalletManagementResource(), fieldMap, readSingularDatabaseAutonomousDatabaseInstanceWalletManagement)
}

func readSingularDatabaseAutonomousDatabaseInstanceWalletManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseInstanceWalletManagementDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseAutonomousDatabaseInstanceWalletManagementDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.AutonomousDatabaseWallet
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseWalletRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousDatabaseWallet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseWallet
	return nil
}

func (s *DatabaseAutonomousDatabaseInstanceWalletManagementDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseAutonomousDatabaseInstanceWalletManagementDataSource-", DatabaseAutonomousDatabaseInstanceWalletManagementDataSource(), s.D))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeRotated != nil {
		s.D.Set("time_rotated", s.Res.TimeRotated.String())
	}

	return nil
}
