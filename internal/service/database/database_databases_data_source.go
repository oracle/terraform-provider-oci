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

func DatabaseDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_home_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseDatabaseResource()),
			},
		},
	}
}

func readDatabaseDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDatabasesResponse
}

func (s *DatabaseDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDatabasesDataSourceCrud) Get() error {
	request := oci_database.ListDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok {
		tmp := dbHomeId.(string)
		request.DbHomeId = &tmp
	}

	if dbName, ok := s.D.GetOkExists("db_name"); ok {
		tmp := dbName.(string)
		request.DbName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.DatabaseSummaryLifecycleStateEnum(state.(string))
	}

	if systemId, ok := s.D.GetOkExists("system_id"); ok {
		tmp := systemId.(string)
		request.SystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDatabasesDataSource-", DatabaseDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		database := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CharacterSet != nil {
			database["character_set"] = *r.CharacterSet
		}

		if r.ConnectionStrings != nil {
			database["connection_strings"] = []interface{}{DatabaseConnectionStringsToMap(r.ConnectionStrings)}
		} else {
			database["connection_strings"] = nil
		}

		if r.DatabaseManagementConfig != nil {
			database["database_management_config"] = []interface{}{CloudDatabaseManagementConfigToMap(r.DatabaseManagementConfig)}
		} else {
			database["database_management_config"] = nil
		}

		if r.DatabaseSoftwareImageId != nil {
			database["database_software_image_id"] = *r.DatabaseSoftwareImageId
		}

		if r.DbBackupConfig != nil {
			database["db_backup_config"] = []interface{}{DbBackupConfigToMap(r.DbBackupConfig)}
		} else {
			database["db_backup_config"] = nil
		}

		if r.DbHomeId != nil {
			database["db_home_id"] = *r.DbHomeId
		}

		if r.DbName != nil {
			database["db_name"] = *r.DbName
		}

		if r.DbSystemId != nil {
			database["db_system_id"] = *r.DbSystemId
		}

		if r.DbUniqueName != nil {
			database["db_unique_name"] = *r.DbUniqueName
		}

		if r.DbWorkload != nil {
			database["db_workload"] = *r.DbWorkload
		}

		if r.DefinedTags != nil {
			database["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		database["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			database["id"] = *r.Id
		}

		if r.IsCdb != nil {
			database["is_cdb"] = *r.IsCdb
		}

		if r.KmsKeyId != nil {
			database["kms_key_id"] = *r.KmsKeyId
		}

		if r.LastBackupTimestamp != nil {
			database["last_backup_timestamp"] = r.LastBackupTimestamp.String()
		}

		if r.LifecycleDetails != nil {
			database["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.NcharacterSet != nil {
			database["ncharacter_set"] = *r.NcharacterSet
		}

		if r.PdbName != nil {
			database["pdb_name"] = *r.PdbName
		}

		if r.SidPrefix != nil {
			database["sid_prefix"] = *r.SidPrefix
		}

		if r.SourceDatabasePointInTimeRecoveryTimestamp != nil {
			database["source_database_point_in_time_recovery_timestamp"] = r.SourceDatabasePointInTimeRecoveryTimestamp.String()
		}

		database["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			database["time_created"] = r.TimeCreated.String()
		}

		if r.VmClusterId != nil {
			database["vm_cluster_id"] = *r.VmClusterId
		}

		resources = append(resources, database)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDatabasesDataSource().Schema["databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("databases", resources); err != nil {
		return err
	}

	return nil
}
