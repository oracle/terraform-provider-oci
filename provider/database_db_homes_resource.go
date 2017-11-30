// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DBHomeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createDBHome,
		Read:     readDBHome,
		Update:   updateDBHome,
		Delete:   deleteDBHome,
		Schema: map[string]*schema.Schema{
			//Required
			"database": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						//Required
						"admin_password": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"db_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						//Optional
						"character_set": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"db_workload": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"ncharacter_set": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"pdb_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						//Computed
					},
				},
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Required: true,
			},

			//Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			//Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_patch_history_entry_id": {
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

func createDBHome(d *schema.ResourceData, m interface{}) (e error) {
	sync := &DBHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.CreateResource(d, sync)
}

func readDBHome(d *schema.ResourceData, m interface{}) (e error) {
	sync := &DBHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.ReadResource(sync)
}

func updateDBHome(d *schema.ResourceData, m interface{}) (e error) {
	sync := &DBHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.UpdateResource(d, sync)
}

func deleteDBHome(d *schema.ResourceData, m interface{}) (e error) {
	sync := &DBHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).clientWithoutNotFoundRetries
	return crud.DeleteResource(d, sync)
}

type DBHomeResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.DBHome
}

func (s *DBHomeResourceCrud) ID() string {
	return s.Res.ID
}

func (s *DBHomeResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *DBHomeResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *DBHomeResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *DBHomeResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *DBHomeResourceCrud) Create() (e error) {
	opts := &baremetal.CreateDBHomeOptions{}
	databaseRaw := s.D.Get("database").([]interface{})
	database := &baremetal.CreateDatabaseDetails{}
	if len(databaseRaw) > 0 {
		databaseMap := databaseRaw[0].(map[string]interface{})
		adminPassword := databaseMap["admin_password"].(string)
		database.AdminPassword = adminPassword
		dbName := databaseMap["db_name"].(string)
		database.DBName = dbName
		characterSet, ok := databaseMap["character_set"]
		if ok && characterSet != nil {
			database.CharacterSet = characterSet.(string)
		}
		dbWorkload, ok := databaseMap["db_workload"]
		if ok && dbWorkload != nil {
			database.DBWorkload = dbWorkload.(string)
		}
		ncharacterSet, ok := databaseMap["ncharacter_set"]
		if ok && ncharacterSet != nil {
			database.NcharacterSet = ncharacterSet.(string)
		}
		pdbName, ok := databaseMap["pdb_name"]
		if ok && pdbName != nil {
			database.PDBName = pdbName.(string)
		}

	}
	dbSystemID := s.D.Get("db_system_id").(string)
	dbVersion := s.D.Get("db_version").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok && displayName != nil {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.CreateDBHome(database, dbSystemID, dbVersion, opts)
	return
}

func (s *DBHomeResourceCrud) Get() (e error) {
	dbHomeID := s.D.Get("id").(string)

	s.Res, e = s.Client.GetDBHome(dbHomeID)
	return
}

func (s *DBHomeResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateDBHomeOptions{}
	dbHomeID := s.D.Get("id").(string)
	dbVersion, ok := s.D.GetOk("db_version")
	if ok && dbVersion != nil {
		opts.DBVersion = dbVersion.(string)
	}

	s.Res, e = s.Client.UpdateDBHome(dbHomeID, opts)
	return
}

func (s *DBHomeResourceCrud) Delete() (e error) {
	dbHomeID := s.D.Get("id").(string)

	e = s.Client.DeleteDBHome(dbHomeID, nil)
	return
}

func (s *DBHomeResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("db_system_id", s.Res.DBSystemID)
	s.D.Set("db_version", s.Res.DBVersion)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("id", s.Res.ID)
	s.D.Set("last_patch_history_entry_id", s.Res.LastPatchHistoryEntryID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}
