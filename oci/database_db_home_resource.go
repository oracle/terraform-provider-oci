// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseDbHomeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &TwelveHours,
			Delete: &TwoHours,
			Update: &TwoHours,
		},
		Create: createDatabaseDbHome,
		Read:   readDatabaseDbHome,
		Update: updateDatabaseDbHome,
		Delete: deleteDatabaseDbHome,
		Schema: map[string]*schema.Schema{
			// Required
			"database": {
				Type:     schema.TypeList,
				Required: true,
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

						// Optional
						"backup_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"backup_tde_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},
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
									},
									"auto_backup_window": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"backup_destination_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												// Optional
												"id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
											},
										},
									},
									"recovery_window_in_days": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"db_name": {
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
						"defined_tags": {
							Type:             schema.TypeMap,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: definedTagsDiffSuppressFunction,
							Elem:             schema.TypeString,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
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
						"connection_strings": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"all_connection_strings": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"cdb_default": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cdb_ip_default": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"db_unique_name": {
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
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Optional
			"db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"db_version": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: dbVersionDiffSuppress,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"DB_BACKUP",
					"NONE",
					"VM_CLUSTER_NEW",
				}, true),
			},
			"vm_cluster_id": {
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

func createDatabaseDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return CreateResource(d, sync)
}

func readDatabaseDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

func updateDatabaseDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return UpdateResource(d, sync)
}

func deleteDatabaseDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseDbHomeResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.DbHome
	Database               *oci_database.Database
	DisableNotFoundRetries bool
}

func (s *DatabaseDbHomeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDbHomeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DbHomeLifecycleStateProvisioning),
	}
}

func (s *DatabaseDbHomeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.DbHomeLifecycleStateAvailable),
	}
}

func (s *DatabaseDbHomeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DbHomeLifecycleStateTerminating),
	}
}

func (s *DatabaseDbHomeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DbHomeLifecycleStateTerminated),
	}
}

func (s *DatabaseDbHomeResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.DbHomeLifecycleStateUpdating),
	}
}

func (s *DatabaseDbHomeResourceCrud) UpdatedTarget() []string {
	return s.CreatedTarget()
}

func (s *DatabaseDbHomeResourceCrud) Create() error {
	request := oci_database.CreateDbHomeRequest{}
	err := s.populateTopLevelPolymorphicCreateDbHomeRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	// Only one dbHome can be created at a time on service side as the dbSystem will go into UPDATING state.
	// A 409 is returned if you try to create a dbHome when the state of the dbSystem is UPDATING
	// For the case of multiple dbHomes we want to wait for the dbSystem to not be UPDATING so that we use the Create Timeout instead of the retry timeout
	dbSystemMayBeUpdating := true
	startTime := time.Now()
	endTime := startTime.Add(s.D.Timeout(schema.TimeoutCreate))
	for dbSystemMayBeUpdating {
		timeout := endTime.Sub(time.Now())

		switch v := (request.CreateDbHomeWithDbSystemIdDetails).(type) {
		case oci_database.CreateDbHomeWithDbSystemIdDetails:
		case oci_database.CreateDbHomeWithDbSystemIdFromBackupDetails:
			_, err = waitForDbSystemIfItIsUpdating(v.DbSystemId, s.Client, timeout)
			if err != nil {
				return err
			}
		case oci_database.CreateDbHomeWithVmClusterIdDetails:
			_, err = waitForVmClusterIfItIsUpdating(v.VmClusterId, s.Client, timeout)
			if err != nil {
				return err
			}
		default:
			log.Printf("[WARN] Received 'CreateDbHomeWithDbSystemIdDetails' of unknown type %v", request.CreateDbHomeWithDbSystemIdDetails)
		}

		response, err := s.Client.CreateDbHome(context.Background(), request)
		// because we cannot control what sets the DbSystem into UPDATING state we have to check again to account for race conditions
		if err != nil {
			if strings.Contains(err.Error(), "has a conflicting state of UPDATING") && timeout > 0 {
				continue
			}
			return err
		}
		dbSystemMayBeUpdating = false
		s.Res = &response.DbHome
	}
	err = s.getDatabaseInfo()
	if err != nil {
		log.Printf("[ERROR] Could not get Database info for the dbHome: %v", err)
	}
	return nil
}

func (s *DatabaseDbHomeResourceCrud) Get() error {
	request := oci_database.GetDbHomeRequest{}

	tmp := s.D.Id()
	request.DbHomeId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbHome

	err = s.getDatabaseInfo()
	if err != nil {
		log.Printf("[ERROR] Could not get Database info for the dbHome: %v", err)
	}

	return nil
}

func (s *DatabaseDbHomeResourceCrud) getDatabaseInfo() error {
	listDatabasesRequest := oci_database.ListDatabasesRequest{}

	listDatabasesRequest.CompartmentId = s.Res.CompartmentId
	listDatabasesRequest.DbHomeId = s.Res.Id
	listDatabasesRequest.SortBy = oci_database.ListDatabasesSortByTimecreated
	listDatabasesRequest.SortOrder = oci_database.ListDatabasesSortOrderAsc
	listDatabasesRequest.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")
	listDatabasesResponse, err := s.Client.ListDatabases(context.Background(), listDatabasesRequest)
	if err != nil {
		return err
	}
	if len(listDatabasesResponse.Items) <= 0 {
		return fmt.Errorf("could not get details of the database for the dbHome")
	}

	databaseId := listDatabasesResponse.Items[0].Id

	getDatabaseRequest := oci_database.GetDatabaseRequest{}
	getDatabaseRequest.DatabaseId = databaseId
	getDatabaseRequest.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")
	getDatabaseResponse, err := s.Client.GetDatabase(context.Background(), getDatabaseRequest)
	if err != nil {
		return err
	}

	s.Database = &getDatabaseResponse.Database

	return nil
}

func (s *DatabaseDbHomeResourceCrud) Update() error {
	err := s.Get()
	if err != nil {
		log.Printf("[ERROR] error refreshing the dbHome information before an upate: %v", err)
	}
	if s.Database == nil || s.Database.Id == nil {
		err := s.getDatabaseInfo()
		if err != nil {
			return fmt.Errorf("could not perform an update as we could not get the databaseId in the dbHome: %v", err)
		}
	}

	request := oci_database.UpdateDatabaseRequest{}

	request.DatabaseId = s.Database.Id

	if database, ok := s.D.GetOkExists("database"); ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
			tmp, err := s.mapToUpdateDatabaseDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UpdateDatabaseDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")
	updateDatabaseResponse, err := s.Client.UpdateDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	//wait for database to not be in updating state after the update
	getDatabaseRequest := oci_database.GetDatabaseRequest{}

	getDatabaseRequest.DatabaseId = s.Database.Id

	getDatabaseRequest.RequestMetadata.RetryPolicy = waitForDatabaseUpdateRetryPolicy(s.D.Timeout(schema.TimeoutUpdate))
	getDatabaseResponse, err := s.Client.GetDatabase(context.Background(), getDatabaseRequest)
	if err != nil {
		// In UpdateDatabase some properties are updated right away like tags but others like auto_backup_enabled are only updated after lifecycleState is not Updating so we update the state here as well in the case of an error in the polling
		s.Database = &updateDatabaseResponse.Database
		err = s.SetData()
		if err != nil {
			log.Printf("[ERROR] error setting data after polling error on database: %v", err)
		}
		return fmt.Errorf("[ERROR] unable to get database after the update: %v", err)
	}

	s.Database = &getDatabaseResponse.Database

	return err
}

func (s *DatabaseDbHomeResourceCrud) Delete() error {
	request := oci_database.DeleteDbHomeRequest{}

	tmp := s.D.Id()
	request.DbHomeId = &tmp

	if performFinalBackup, ok := s.D.GetOkExists("perform_final_backup"); ok {
		tmp := performFinalBackup.(bool)
		request.PerformFinalBackup = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	// Only one dbHome can be deleted at a time on service side as the dbSystem will go into UPDATING state.
	// A 409 is returned if you try to delete a dbHome when the state of the dbSystem is UPDATING
	// For the case of multiple dbHomes we want to wait for the dbSystem to not be UPDATING so that we use the Delete Timeout instead of the retry timeout
	dbSystemMayBeUpdating := true
	startTime := time.Now()
	endTime := startTime.Add(s.D.Timeout(schema.TimeoutDelete))
	for dbSystemMayBeUpdating {
		timeout := endTime.Sub(time.Now())
		dbSystemId, ok := s.D.GetOkExists("db_system_id")
		if ok && dbSystemId != nil {
			dbSystemIdStr := dbSystemId.(string)
			_, err := waitForDbSystemIfItIsUpdating(&dbSystemIdStr, s.Client, timeout)
			if err != nil {
				return err
			}
		} else {
			dbSystemMayBeUpdating = false
		}

		_, err := s.Client.DeleteDbHome(context.Background(), request)
		// because we cannot control what sets the DbSystem into UPDATING state we have to check again to account for race conditions
		if err != nil {
			if strings.Contains(err.Error(), "has a conflicting state of UPDATING") && timeout > 0 {
				continue
			}
			return err
		}
		dbSystemMayBeUpdating = false
	}
	return nil
}

func (s *DatabaseDbHomeResourceCrud) SetData() error {
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
		s.D.SetId(*s.Res.Id)
	}

	if s.Res.LastPatchHistoryEntryId != nil {
		s.D.Set("last_patch_history_entry_id", *s.Res.LastPatchHistoryEntryId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VmClusterId != nil {
		s.D.Set("vm_cluster_id", *s.Res.VmClusterId)
	}

	if s.Database != nil {
		s.D.Set("database", []interface{}{s.DatabaseToMap(s.Database)})
	}

	if source, ok := s.D.GetOkExists("source"); !ok || source.(string) == "" {
		s.D.Set("source", "NONE")
	}
	return nil
}

func (s *DatabaseDbHomeResourceCrud) mapToBackupDestinationDetails(fieldKeyFormat string) (oci_database.BackupDestinationDetails, error) {
	result := oci_database.BackupDestinationDetails{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_database.BackupDestinationDetailsTypeEnum(type_.(string))
	}

	return result, nil
}

func (s *DatabaseDbHomeResourceCrud) mapToCreateDatabaseDetails(fieldKeyFormat string) (oci_database.CreateDatabaseDetails, error) {
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
				return result, fmt.Errorf("unable to convert db_backup_config, encountered error: %v", err)
			}
			result.DbBackupConfig = &tmp
		}
	}

	if dbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_name")); ok {
		tmp := dbName.(string)
		result.DbName = &tmp
	}

	if dbWorkload, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_workload")); ok {
		result.DbWorkload = oci_database.CreateDatabaseDetailsDbWorkloadEnum(dbWorkload.(string))
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
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

func (s *DatabaseDbHomeResourceCrud) DatabaseToMap(obj *oci_database.Database) map[string]interface{} {
	result := map[string]interface{}{}

	//Create parameters that are not returned by the service
	if adminPassword, ok := s.D.GetOkExists("database.0.admin_password"); ok && adminPassword != nil {
		result["admin_password"] = adminPassword.(string)
	}

	if backupId, ok := s.D.GetOkExists("database.0.backup_id"); ok && backupId != nil {
		result["backup_id"] = backupId.(string)
	}

	if backupTDEPassword, ok := s.D.GetOkExists("database.0.backup_tde_password"); ok && backupTDEPassword != nil {
		result["backup_tde_password"] = backupTDEPassword.(string)
	}

	if obj.CharacterSet != nil {
		result["character_set"] = string(*obj.CharacterSet)
	}

	if obj.ConnectionStrings != nil {
		result["connection_strings"] = []interface{}{DatabaseConnectionStringsToMap(obj.ConnectionStrings)}
	}

	if obj.DbBackupConfig != nil {
		result["db_backup_config"] = []interface{}{DbBackupConfigToMap(obj.DbBackupConfig)}
	}

	if obj.DbName != nil {
		result["db_name"] = string(*obj.DbName)
	}

	if obj.DbUniqueName != nil {
		result["db_unique_name"] = string(*obj.DbUniqueName)
	}

	if obj.DbWorkload != nil {
		result["db_workload"] = string(*obj.DbWorkload)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.NcharacterSet != nil {
		result["ncharacter_set"] = string(*obj.NcharacterSet)
	}

	if obj.PdbName != nil {
		result["pdb_name"] = string(*obj.PdbName)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *DatabaseDbHomeResourceCrud) mapToCreateDatabaseFromBackupDetails(fieldKeyFormat string) (oci_database.CreateDatabaseFromBackupDetails, error) {
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

	if dbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_name")); ok {
		tmp := dbName.(string)
		result.DbName = &tmp
	}

	return result, nil
}

func (s *DatabaseDbHomeResourceCrud) mapToUpdateDatabaseDetails(fieldKeyFormat string) (oci_database.UpdateDatabaseDetails, error) {
	result := oci_database.UpdateDatabaseDetails{}

	if dbBackupConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_backup_config")); ok {
		if tmpList := dbBackupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "db_backup_config"), 0)
			tmp, err := s.mapToUpdateDbBackupConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			result.DbBackupConfig = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	return result, nil
}

func (s *DatabaseDbHomeResourceCrud) CreateDatabaseFromBackupDetailsToMap(obj *oci_database.CreateDatabaseFromBackupDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminPassword != nil {
		result["admin_password"] = string(*obj.AdminPassword)
	}

	if obj.BackupId != nil {
		result["backup_id"] = string(*obj.BackupId)
	}

	if obj.BackupTDEPassword != nil {
		result["backup_tde_password"] = string(*obj.BackupTDEPassword)
	}

	if obj.DbName != nil {
		result["db_name"] = string(*obj.DbName)
	}

	return result
}

// We cannot use the same function we use in create because the HasChanged check needed for the update to succeed interferes with the Create functionality
func (s *DatabaseDbHomeResourceCrud) mapToUpdateDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
	result := oci_database.DbBackupConfig{}

	// Service does not allow to update auto_backup_enabled and recovery_window_in_days at the same time so we must have the HasChanged check
	if autoBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")) {
		tmp := autoBackupEnabled.(bool)
		result.AutoBackupEnabled = &tmp
	}

	if autoBackupWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_window")); ok {
		if result.AutoBackupEnabled != nil && *result.AutoBackupEnabled == true {
			result.AutoBackupWindow = oci_database.DbBackupConfigAutoBackupWindowEnum(autoBackupWindow.(string))
		}
	}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")) {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	return result, nil
}

func (s *DatabaseDbHomeResourceCrud) mapToDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
	result := oci_database.DbBackupConfig{}

	if autoBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")); ok {
		tmp := autoBackupEnabled.(bool)
		result.AutoBackupEnabled = &tmp
	}

	if autoBackupWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_window")); ok {
		if result.AutoBackupEnabled != nil && *result.AutoBackupEnabled == true {
			result.AutoBackupWindow = oci_database.DbBackupConfigAutoBackupWindowEnum(autoBackupWindow.(string))
		}
	}

	if backupDestinationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")); ok {
		result.BackupDestinationDetails = []oci_database.BackupDestinationDetails{}
		interfaces := backupDestinationDetails.([]interface{})
		tmp := make([]oci_database.BackupDestinationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backup_destination_details"), stateDataIndex)
			converted, err := s.mapToBackupDestinationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")) {
			result.BackupDestinationDetails = tmp
		}
	}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	return result, nil
}

func (s *DatabaseDbHomeResourceCrud) DbBackupConfigToMap(obj *oci_database.DbBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AutoBackupEnabled != nil {
		result["auto_backup_enabled"] = bool(*obj.AutoBackupEnabled)
	}

	result["auto_backup_window"] = string(obj.AutoBackupWindow)

	backupDestinationDetails := []interface{}{}
	for _, item := range obj.BackupDestinationDetails {
		backupDestinationDetails = append(backupDestinationDetails, BackupDestinationDetailsToMap(item))
	}
	result["backup_destination_details"] = backupDestinationDetails

	if obj.RecoveryWindowInDays != nil {
		result["recovery_window_in_days"] = int(*obj.RecoveryWindowInDays)
	}

	return result
}

func (s *DatabaseDbHomeResourceCrud) populateTopLevelPolymorphicCreateDbHomeRequest(request *oci_database.CreateDbHomeRequest) error {
	//discriminator
	sourceRaw, ok := s.D.GetOkExists("source")
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "NONE" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("DB_BACKUP"):
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
		request.CreateDbHomeWithDbSystemIdDetails = details
	case strings.ToLower("NONE"):
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
		if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		request.CreateDbHomeWithDbSystemIdDetails = details
	case strings.ToLower("VM_CLUSTER_NEW"):
		details := oci_database.CreateDbHomeWithVmClusterIdDetails{}
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
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
			tmp := vmClusterId.(string)
			details.VmClusterId = &tmp
		}
		request.CreateDbHomeWithDbSystemIdDetails = details
	default:
		return fmt.Errorf("unknown source '%v' was specified", source)
	}
	return nil
}
