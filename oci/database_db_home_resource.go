// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DbHomeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
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

	return CreateResource(d, sync)
}

func readDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

func deleteDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DbHomeResourceCrud struct {
	BaseCrud
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

func (s *DbHomeResourceCrud) SetData() error {
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

	if s.Res.LastPatchHistoryEntryId != nil {
		s.D.Set("last_patch_history_entry_id", *s.Res.LastPatchHistoryEntryId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
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
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
				tmp, err := s.mapToCreateDatabaseFromBackupDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
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
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
				tmp, err := s.mapToCreateDatabaseDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
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
		return fmt.Errorf("unknown source '%v' was specified", source)
	}
	return nil
}

func (s *DbHomeResourceCrud) mapToCreateDatabaseFromBackupDetails(fieldKeyFormat string) (oci_database.CreateDatabaseFromBackupDetails, error) {
	result := oci_database.CreateDatabaseFromBackupDetails{}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
	}

	if backupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_id")); ok {
		tmp := backupId.(string)
		result.BackupId = &tmp
	}

	if backupTDEPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_tde_password")); ok {
		tmp := backupTDEPassword.(string)
		result.BackupTDEPassword = &tmp
	}

	return result, nil
}

func (s *DbHomeResourceCrud) mapToCreateDatabaseDetails(fieldKeyFormat string) (oci_database.CreateDatabaseDetails, error) {
	result := oci_database.CreateDatabaseDetails{}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
	}

	if characterSet, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "character_set")); ok {
		tmp := characterSet.(string)
		result.CharacterSet = &tmp
	}

	if dbBackupConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_backup_config")); ok {
		if tmpList := dbBackupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "db_backup_config"), 0)
			tmp, err := s.mapToDbBackupConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			result.DbBackupConfig = &tmp
		}
	}

	if dbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_name")); ok {
		tmp := dbName.(string)
		result.DbName = &tmp
	}

	if dbWorkload, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_workload")); ok {
		tmp := oci_database.CreateDatabaseDetailsDbWorkloadEnum(dbWorkload.(string))
		result.DbWorkload = tmp
	}

	if ncharacterSet, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ncharacter_set")); ok {
		tmp := ncharacterSet.(string)
		result.NcharacterSet = &tmp
	}

	if pdbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pdb_name")); ok {
		tmp := pdbName.(string)
		result.PdbName = &tmp
	}

	return result, nil
}

func (s *DbHomeResourceCrud) mapToDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
	result := oci_database.DbBackupConfig{}

	if autoBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")); ok {
		tmp := autoBackupEnabled.(bool)
		result.AutoBackupEnabled = &tmp
	}

	return result, nil
}
