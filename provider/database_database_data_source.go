// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DatabaseDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabase,
		Schema: map[string]*schema.Schema{
			// @codegen: support legacy property
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Required

			// Optional

			// Computed
			"character_set": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_backup_config": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"auto_backup_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"db_home_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_unique_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_workload": {
				Type:     schema.TypeString,
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
			"ncharacter_set": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pdb_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type DatabaseDataSourceCrud struct {
	crud.BaseCrud
	Client *oci_database.DatabaseClient
	Res    *oci_database.Database
}

func (s *DatabaseDataSourceCrud) Get() (e error) {
	request := oci_database.GetDatabaseRequest{}

	tmp := s.D.Get("database_id").(string)
	request.DatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Database
	return nil
}

func (s *DatabaseDataSourceCrud) SetData() {
	s.D.SetId(*s.Res.Id)

	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbBackupConfig != nil {
		s.D.Set("db_backup_config", []interface{}{dbBackupConfigToMap(s.Res.DbBackupConfig)})
	}

	if s.Res.DbHomeId != nil {
		s.D.Set("db_home_id", *s.Res.DbHomeId)
	}

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbUniqueName != nil {
		s.D.Set("db_unique_name", *s.Res.DbUniqueName)
	}

	if s.Res.DbWorkload != nil {
		s.D.Set("db_workload", *s.Res.DbWorkload)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

	if s.Res.PdbName != nil {
		s.D.Set("pdb_name", *s.Res.PdbName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func dbBackupConfigToMap(obj *oci_database.DbBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AutoBackupEnabled != nil {
		result["auto_backup_enabled"] = bool(*obj.AutoBackupEnabled)
	}

	return result
}

func mapToDbBackupConfig(raw map[string]interface{}) oci_database.DbBackupConfig {
	result := oci_database.DbBackupConfig{}

	if autoBackupEnabled, ok := raw["auto_backup_enabled"]; ok {
		tmp := autoBackupEnabled.(bool)
		result.AutoBackupEnabled = &tmp
	}

	return result
}
