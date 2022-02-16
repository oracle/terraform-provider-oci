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

func DatabaseAutonomousDatabaseBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDatabaseBackups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_database_backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseAutonomousDatabaseBackupResource()),
			},
		},
	}
}

func readDatabaseAutonomousDatabaseBackups(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabaseBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDatabaseBackupsResponse
}

func (s *DatabaseAutonomousDatabaseBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseBackupsDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDatabaseBackupsRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousDatabaseBackupSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDatabaseBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDatabaseBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseBackupsDataSource-", DatabaseAutonomousDatabaseBackupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDatabaseBackup := map[string]interface{}{}

		if r.AutonomousDatabaseId != nil {
			autonomousDatabaseBackup["autonomous_database_id"] = *r.AutonomousDatabaseId
		}

		if r.CompartmentId != nil {
			autonomousDatabaseBackup["compartment_id"] = *r.CompartmentId
		}

		if r.DatabaseSizeInTBs != nil {
			autonomousDatabaseBackup["database_size_in_tbs"] = *r.DatabaseSizeInTBs
		}

		if r.DisplayName != nil {
			autonomousDatabaseBackup["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			autonomousDatabaseBackup["id"] = *r.Id
		}

		if r.IsAutomatic != nil {
			autonomousDatabaseBackup["is_automatic"] = *r.IsAutomatic
		}

		if r.IsRestorable != nil {
			autonomousDatabaseBackup["is_restorable"] = *r.IsRestorable
		}

		if r.KeyStoreId != nil {
			autonomousDatabaseBackup["key_store_id"] = *r.KeyStoreId
		}

		if r.KeyStoreWalletName != nil {
			autonomousDatabaseBackup["key_store_wallet_name"] = *r.KeyStoreWalletName
		}

		if r.KmsKeyId != nil {
			autonomousDatabaseBackup["kms_key_id"] = *r.KmsKeyId
		}

		if r.KmsKeyVersionId != nil {
			autonomousDatabaseBackup["kms_key_version_id"] = *r.KmsKeyVersionId
		}

		if r.LifecycleDetails != nil {
			autonomousDatabaseBackup["lifecycle_details"] = *r.LifecycleDetails
		}

		autonomousDatabaseBackup["state"] = r.LifecycleState

		if r.TimeEnded != nil {
			autonomousDatabaseBackup["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeStarted != nil {
			autonomousDatabaseBackup["time_started"] = r.TimeStarted.String()
		}

		autonomousDatabaseBackup["type"] = r.Type

		if r.VaultId != nil {
			autonomousDatabaseBackup["vault_id"] = *r.VaultId
		}

		resources = append(resources, autonomousDatabaseBackup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousDatabaseBackupsDataSource().Schema["autonomous_database_backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_database_backups", resources); err != nil {
		return err
	}

	return nil
}
