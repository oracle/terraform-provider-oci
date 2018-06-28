// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DbHomeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createDbHome,
		Read:     readDbHome,
		Delete:   deleteDbHome,
		Schema: map[string]*schema.Schema{
			// Required
			"database": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"admin_password": {
							Type:      schema.TypeString,
							Required:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"backup_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"backup_tde_password": {
							Type:      schema.TypeString,
							Optional:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"db_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"character_set": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"db_backup_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"auto_backup_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"db_workload": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"defined_tags": {
							Type:             schema.TypeMap,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: definedTagsDiffSuppressFunction,
							Elem:             schema.TypeString,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem:     schema.TypeString,
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

						// Computed
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
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
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

func createDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.CreateResource(d, sync)
}

func readDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

func deleteDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type DbHomeResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.DbHome
	DisableNotFoundRetries bool
}

func (s *DbHomeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DbHomeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DbHomeLifecycleStateProvisioning),
	}
}

func (s *DbHomeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.DbHomeLifecycleStateAvailable),
	}
}

func (s *DbHomeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DbHomeLifecycleStateTerminating),
	}
}

func (s *DbHomeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DbHomeLifecycleStateTerminated),
	}
}

func (s *DbHomeResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.DbHomeLifecycleStateUpdating),
	}
}

func (s *DbHomeResourceCrud) UpdatedTarget() []string {
	return s.CreatedTarget()
}

func (s *DbHomeResourceCrud) Create() error {
	request := oci_database.CreateDbHomeRequest{}
	err := s.populateTopLevelPolymorphicCreateDbHomeRequest(&request)
	if err != nil {
		return err
	}

	handleDbSimulationFlag(s.Client)

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbHome
	return nil
}

func (s *DbHomeResourceCrud) Get() error {
	request := oci_database.GetDbHomeRequest{}

	tmp := s.D.Id()
	request.DbHomeId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbHome
	return nil
}

func (s *DbHomeResourceCrud) Delete() error {
	request := oci_database.DeleteDbHomeRequest{}

	tmp := s.D.Id()
	request.DbHomeId = &tmp

	if performFinalBackup, ok := s.D.GetOkExists("perform_final_backup"); ok {
		tmp := performFinalBackup.(bool)
		request.PerformFinalBackup = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteDbHome(context.Background(), request)
	return err
}

func (s *DbHomeResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.LastPatchHistoryEntryId != nil {
		s.D.Set("last_patch_history_entry_id", *s.Res.LastPatchHistoryEntryId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}

func (s *DbHomeResourceCrud) populateTopLevelPolymorphicCreateDbHomeRequest(request *oci_database.CreateDbHomeRequest) error {
	//discriminator
	sourceRaw, ok := s.D.GetOkExists("source")
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "NONE" // default value
	}

	switch source {
	case "DB_BACKUP":
		details := oci_database.CreateDbHomeWithDbSystemIdFromBackupDetails{}
		if database, ok := s.D.GetOkExists("database"); ok {
			if tmpList := database.([]interface{}); len(tmpList) > 0 {
				tmp := mapToCreateDatabaseFromBackupDetails(tmpList[0].(map[string]interface{}))
				details.Database = &tmp
			}
		}
		if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		request.CreateDbHomeWithDbSystemIdBase = details

	case "NONE":
		details := oci_database.CreateDbHomeWithDbSystemIdDetails{}
		if database, ok := s.D.GetOkExists("database"); ok {
			if tmpList := database.([]interface{}); len(tmpList) > 0 {
				tmp := mapToCreateDatabaseDetails(tmpList[0].(map[string]interface{}))
				details.Database = &tmp
			}
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		request.CreateDbHomeWithDbSystemIdBase = details
	default:
		return fmt.Errorf("Unknown source '%v' was specified", source)
	}
	return nil
}
