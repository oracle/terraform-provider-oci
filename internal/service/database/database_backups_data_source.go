// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseBackups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"backup_destination_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"shape_family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_expiry_scheduled_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_expiry_scheduled_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseBackupResource()),
			},
		},
	}
}

func readDatabaseBackups(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListBackupsResponse
}

func (s *DatabaseBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseBackupsDataSourceCrud) Get() error {
	request := oci_database.ListBackupsRequest{}

	if backupDestinationType, ok := s.D.GetOkExists("backup_destination_type"); ok {
		tmp := backupDestinationType.(string)
		request.BackupDestinationType = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if shapeFamily, ok := s.D.GetOkExists("shape_family"); ok {
		request.ShapeFamily = oci_database.ListBackupsShapeFamilyEnum(shapeFamily.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.BackupSummaryLifecycleStateEnum(state.(string))
	}

	if timeExpiryScheduledGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_expiry_scheduled_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeExpiryScheduledGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeExpiryScheduledGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeExpiryScheduledLessThan, ok := s.D.GetOkExists("time_expiry_scheduled_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeExpiryScheduledLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeExpiryScheduledLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		tmp := type_.(string)
		request.Type = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseBackupsDataSource-", DatabaseBackupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		backup := map[string]interface{}{}

		if r.AvailabilityDomain != nil {
			backup["availability_domain"] = *r.AvailabilityDomain
		}

		backup["backup_destination_type"] = r.BackupDestinationType

		if r.CompartmentId != nil {
			backup["compartment_id"] = *r.CompartmentId
		}

		backup["database_edition"] = r.DatabaseEdition

		if r.DatabaseId != nil {
			backup["database_id"] = *r.DatabaseId
		}

		if r.DatabaseSizeInGBs != nil {
			backup["database_size_in_gbs"] = *r.DatabaseSizeInGBs
		}

		if r.DisplayName != nil {
			backup["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			backup["id"] = *r.Id
		}

		if r.IsUsingOracleManagedKeys != nil {
			backup["is_using_oracle_managed_keys"] = *r.IsUsingOracleManagedKeys
		}

		if r.KeyStoreId != nil {
			backup["key_store_id"] = *r.KeyStoreId
		}

		if r.KeyStoreWalletName != nil {
			backup["key_store_wallet_name"] = *r.KeyStoreWalletName
		}

		if r.KmsKeyId != nil {
			backup["kms_key_id"] = *r.KmsKeyId
		}

		if r.KmsKeyVersionId != nil {
			backup["kms_key_version_id"] = *r.KmsKeyVersionId
		}

		if r.LifecycleDetails != nil {
			backup["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.RetentionPeriodInDays != nil {
			backup["retention_period_in_days"] = *r.RetentionPeriodInDays
		}

		if r.RetentionPeriodInYears != nil {
			backup["retention_period_in_years"] = *r.RetentionPeriodInYears
		}

		backup["secondary_kms_key_ids"] = r.SecondaryKmsKeyIds

		if r.Shape != nil {
			backup["shape"] = *r.Shape
		}

		backup["state"] = r.LifecycleState

		if r.TimeEnded != nil {
			backup["time_ended"] = r.TimeEnded.Format(time.RFC3339Nano)
		}

		if r.TimeExpiryScheduled != nil {
			backup["time_expiry_scheduled"] = r.TimeExpiryScheduled.String()
		}

		if r.TimeStarted != nil {
			backup["time_started"] = r.TimeStarted.Format(time.RFC3339Nano)
		}

		backup["type"] = r.Type

		if r.VaultId != nil {
			backup["vault_id"] = *r.VaultId
		}

		if r.Version != nil {
			backup["version"] = *r.Version
		}

		resources = append(resources, backup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseBackupsDataSource().Schema["backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("backups", resources); err != nil {
		return err
	}

	return nil
}
