// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_work_requests "github.com/oracle/oci-go-sdk/v56/workrequests"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/oracle/oci-go-sdk/v56/common"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseDbHomeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("2h"),
			Delete: tfresource.GetTimeoutDuration("2h"),
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
						"database_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"database_software_image_id": {
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
												// Required

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

												// Computed
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
							DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
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
						"tde_wallet_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							Sensitive: true,
						},
						"time_stamp_for_point_in_time_recovery": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: utils.TimeDiffSuppressFunction,
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
						"one_off_patches": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sid_prefix": {
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
			"database_software_image_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
				DiffSuppressFunc: tfresource.DbVersionDiffSuppress,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_desupported_version": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"kms_key_version_id": {
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
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"DATABASE",
					"DB_BACKUP",
					"NONE",
					"VM_CLUSTER_BACKUP",
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
			"db_home_location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_patch_history_entry_id": {
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
	}
}

func createDatabaseDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func deleteDatabaseDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseDbHomeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
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
	return []string{
		string(oci_database.DbHomeLifecycleStateAvailable),
	}
}

func (s *DatabaseDbHomeResourceCrud) Create() error {
	request := oci_database.CreateDbHomeRequest{}
	err := s.populateTopLevelPolymorphicCreateDbHomeRequest(&request)
	if err != nil {
		return err
	}

	// Special override to ensure that CreateDbHome retries for the duration of the Terraform configured Create timeout
	// The underlying db system or vm cluster may be in an updating state. So keep retrying the CreateDbHome.
	createDbHomeRetryDurationFn := tfresource.GetDbHomeRetryDurationFunction(s.D.Timeout(schema.TimeoutCreate))

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database", createDbHomeRetryDurationFn)

	response, err := s.Client.CreateDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.DbHome

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

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

func (s *DatabaseDbHomeResourceCrud) Update() error {
	updateDbHomeRequest := oci_database.UpdateDbHomeRequest{}

	tmp := s.D.Id()
	updateDbHomeRequest.DbHomeId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		updateDbHomeRequest.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		updateDbHomeRequest.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}
	updateDbHomeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	if oneOffPatches, ok := s.D.GetOkExists("one_off_patches"); ok {
		interfaces := oneOffPatches.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
	}
	response, err := s.Client.UpdateDbHome(context.Background(), updateDbHomeRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.DbHome
	err = s.Get()
	if err != nil {
		log.Printf("[ERROR] error refreshing the dbHome information before an upate: %v", err)
	}
	if s.Database == nil || s.Database.Id == nil {
		err := s.getDatabaseInfo()
		if err != nil {
			return fmt.Errorf("could not perform an Update as we could not get the databaseId in the dbHome: %v", err)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
	updateDatabaseResponse, err := s.Client.UpdateDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId = response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	getDatabaseRequest := oci_database.GetDatabaseRequest{}

	getDatabaseRequest.DatabaseId = s.Database.Id

	getDatabaseRequest.RequestMetadata.RetryPolicy = waitForDatabaseUpdateRetryPolicy(s.D.Timeout(schema.TimeoutUpdate))
	getDatabaseResponse, err := s.Client.GetDatabase(context.Background(), getDatabaseRequest)
	if err != nil {
		s.Database = &updateDatabaseResponse.Database
		err = s.SetData()
		if err != nil {
			log.Printf("[ERROR] error setting data after polling error on database: %v", err)
		}
		return fmt.Errorf("[ERROR] unable to get database after the Update: %v", err)
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

	// Special override to ensure that DeleteDbHome retries for the duration of the Terraform configured Create timeout
	// The underlying db system or vm cluster may be in an updating state. So keep retrying it.
	deleteDbHomeRetryDurationFn := tfresource.GetDbHomeRetryDurationFunction(s.D.Timeout(schema.TimeoutDelete))
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database", deleteDbHomeRetryDurationFn)

	dbErr := s.deleteNestedDB()
	if dbErr != nil {
		log.Printf("[WARN] Could not delete nested database in DbHome. Will proceed to delete dbHome: %v", dbErr)
	}

	_, err := s.Client.DeleteDbHome(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatabaseDbHomeResourceCrud) SetData() error {

	if s.Res.Id != nil {
		s.D.SetId(*s.Res.Id)
	}

	if s.Database != nil {
		s.D.Set("database", []interface{}{s.DatabaseToMap(s.Database)})
	}

	if source, ok := s.D.GetOkExists("source"); !ok || source.(string) == "" {
		s.D.Set("source", "NONE")
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseSoftwareImageId != nil {
		s.D.Set("database_software_image_id", *s.Res.DatabaseSoftwareImageId)
	}

	if s.Res.DbHomeLocation != nil {
		s.D.Set("db_home_location", *s.Res.DbHomeLocation)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.LastPatchHistoryEntryId != nil {
		s.D.Set("last_patch_history_entry_id", *s.Res.LastPatchHistoryEntryId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VmClusterId != nil {
		s.D.Set("vm_cluster_id", *s.Res.VmClusterId)
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

	if databaseSoftwareImageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_software_image_id")); ok {
		tmp := databaseSoftwareImageId.(string)
		result.DatabaseSoftwareImageId = &tmp
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
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ncharacterSet, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ncharacter_set")); ok {
		tmp := ncharacterSet.(string)
		result.NcharacterSet = &tmp
	}

	if pdbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pdb_name")); ok {
		tmp := pdbName.(string)
		result.PdbName = &tmp
	}

	if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
		tmp := tdeWalletPassword.(string)
		result.TdeWalletPassword = &tmp
	}

	return result, nil
}

func (s *DatabaseDbHomeResourceCrud) mapToCreateDatabaseFromAnotherDatabaseDetails(fieldKeyFormat string) (oci_database.CreateDatabaseFromAnotherDatabaseDetails, error) {
	result := oci_database.CreateDatabaseFromAnotherDatabaseDetails{}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
	}

	if backupTDEPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_tde_password")); ok {
		tmp := backupTDEPassword.(string)
		result.BackupTDEPassword = &tmp
	}

	if databaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_id")); ok {
		tmp := databaseId.(string)
		result.DatabaseId = &tmp
	}

	if dbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_name")); ok {
		tmp := dbName.(string)
		result.DbName = &tmp
	}

	if dbUniqueName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_unique_name")); ok {
		tmp := dbUniqueName.(string)
		result.DbUniqueName = &tmp
	}

	if timeStampForPointInTimeRecovery, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_stamp_for_point_in_time_recovery")); ok {
		tmp, err := time.Parse(time.RFC3339, timeStampForPointInTimeRecovery.(string))
		if err != nil {
			return result, err
		}
		result.TimeStampForPointInTimeRecovery = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
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
	case strings.ToLower("DATABASE"):
		details := oci_database.CreateDbHomeWithDbSystemIdFromDatabaseDetails{}
		if database, ok := s.D.GetOkExists("database"); ok {
			if tmpList := database.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
				tmp, err := s.mapToCreateDatabaseFromAnotherDatabaseDetails(fieldKeyFormat)
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
		if databaseSoftwareImageId, ok := s.D.GetOkExists("database_software_image_id"); ok {
			tmp := databaseSoftwareImageId.(string)
			details.DatabaseSoftwareImageId = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if isDesupportedVersion, ok := s.D.GetOkExists("is_desupported_version"); ok {
			tmp := isDesupportedVersion.(bool)
			details.IsDesupportedVersion = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		request.CreateDbHomeWithDbSystemIdDetails = details
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
		if databaseSoftwareImageId, ok := s.D.GetOkExists("database_software_image_id"); ok {
			tmp := databaseSoftwareImageId.(string)
			details.DatabaseSoftwareImageId = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if isDesupportedVersion, ok := s.D.GetOkExists("is_desupported_version"); ok {
			tmp := isDesupportedVersion.(bool)
			details.IsDesupportedVersion = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
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
		if databaseSoftwareImageId, ok := s.D.GetOkExists("database_software_image_id"); ok {
			tmp := databaseSoftwareImageId.(string)
			details.DatabaseSoftwareImageId = &tmp
		}
		if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isDesupportedVersion, ok := s.D.GetOkExists("is_desupported_version"); ok {
			tmp := isDesupportedVersion.(bool)
			details.IsDesupportedVersion = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		request.CreateDbHomeWithDbSystemIdDetails = details
	case strings.ToLower("VM_CLUSTER_BACKUP"):
		details := oci_database.CreateDbHomeWithVmClusterIdFromBackupDetails{}
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
		if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
			tmp := vmClusterId.(string)
			details.VmClusterId = &tmp
		}
		if databaseSoftwareImageId, ok := s.D.GetOkExists("database_software_image_id"); ok {
			tmp := databaseSoftwareImageId.(string)
			details.DatabaseSoftwareImageId = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if isDesupportedVersion, ok := s.D.GetOkExists("is_desupported_version"); ok {
			tmp := isDesupportedVersion.(bool)
			details.IsDesupportedVersion = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
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
		if databaseSoftwareImageId, ok := s.D.GetOkExists("database_software_image_id"); ok {
			tmp := databaseSoftwareImageId.(string)
			details.DatabaseSoftwareImageId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isDesupportedVersion, ok := s.D.GetOkExists("is_desupported_version"); ok {
			tmp := isDesupportedVersion.(bool)
			details.IsDesupportedVersion = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
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

func updateDatabaseDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbHomeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func (s *DatabaseDbHomeResourceCrud) deleteNestedDB() error {

	request := oci_database.DeleteDatabaseRequest{}

	dbCompartment, ok := s.D.GetOkExists("compartment_id")
	if !ok {
		return fmt.Errorf("no compartment information to delete nested database")
	}

	dbHomeIdStr := s.D.Id()
	dbCompartmentStr := dbCompartment.(string)

	listDBRequest := oci_database.ListDatabasesRequest{}
	listDBRequest.CompartmentId = &dbCompartmentStr
	listDBRequest.DbHomeId = &dbHomeIdStr
	listDBRequest.SortBy = oci_database.ListDatabasesSortByTimecreated
	listDBRequest.SortOrder = oci_database.ListDatabasesSortOrderAsc
	listDBRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")
	listDatabasesResponse, err := s.Client.ListDatabases(context.Background(), listDBRequest)
	if err != nil {
		return err
	}

	if len(listDatabasesResponse.Items) == 0 {
		return fmt.Errorf("failed to get details of the nested database in dbHome")
	}

	dbHomeTimeCreated, ok := s.D.GetOkExists("time_created")
	if !ok {
		tmp, err := time.Parse(time.RFC3339, dbHomeTimeCreated.(string))
		if err != nil {
			return err
		}
		if listDatabasesResponse.Items[0].TimeCreated.Sub(common.SDKTime{Time: tmp}.Time) > time.Hour*24 {
			return fmt.Errorf("the first database of the dbHome has since been terminated. Will not try to delete dbHome's database")
		}
	}

	request.DatabaseId = listDatabasesResponse.Items[0].Id

	if performFinalBackup, ok := s.D.GetOkExists("perform_final_backup"); ok {
		tmp := performFinalBackup.(bool)
		request.PerformFinalBackup = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteDatabase(context.Background(), request)

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseDbHomeResourceCrud) getDatabaseInfo() error {
	listDatabasesRequest := oci_database.ListDatabasesRequest{}

	listDatabasesRequest.CompartmentId = s.Res.CompartmentId
	listDatabasesRequest.DbHomeId = s.Res.Id
	listDatabasesRequest.SortBy = oci_database.ListDatabasesSortByTimecreated
	listDatabasesRequest.SortOrder = oci_database.ListDatabasesSortOrderAsc
	listDatabasesRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")
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
	getDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")
	getDatabaseResponse, err := s.Client.GetDatabase(context.Background(), getDatabaseRequest)
	if err != nil {
		return err
	}

	s.Database = &getDatabaseResponse.Database

	return nil
}

func (s *DatabaseDbHomeResourceCrud) DatabaseToMap(obj *oci_database.Database) map[string]interface{} {
	result := map[string]interface{}{}

	if adminPassword, ok := s.D.GetOkExists("database.0.admin_password"); ok && adminPassword != nil {
		result["admin_password"] = adminPassword.(string)
	}

	if tdeWalletPassword, ok := s.D.GetOkExists("database.0.tde_wallet_password"); ok && tdeWalletPassword != nil {
		result["tde_wallet_password"] = tdeWalletPassword.(string)
	}

	if backupId, ok := s.D.GetOkExists("database.0.backup_id"); ok && backupId != nil {
		result["backup_id"] = backupId.(string)
	}

	if databaseSoftwareImageId, ok := s.D.GetOkExists("database.0.database_software_image_id"); ok && databaseSoftwareImageId != nil {
		result["database_software_image_id"] = databaseSoftwareImageId.(string)
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

	if databaseId, ok := s.D.GetOkExists("database.0.database_id"); ok && databaseId != nil {
		result["database_id"] = databaseId.(string)
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

	if s.Res.OneOffPatches != nil {
		result["one_off_patches"] = s.Res.OneOffPatches
	}

	if obj.DbWorkload != nil {
		result["db_workload"] = string(*obj.DbWorkload)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
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

	if timeStampForPointInTimeRecovery, ok := s.D.GetOkExists(fmt.Sprintf("database.0.time_stamp_for_point_in_time_recovery")); ok {
		result["time_stamp_for_point_in_time_recovery"] = timeStampForPointInTimeRecovery
	}

	if obj.SidPrefix != nil {
		result["sid_prefix"] = string(*obj.SidPrefix)
	}

	return result
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
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "admin_password")) {
		tmp := adminPassword.(string)
		result.NewAdminPassword = &tmp
	}

	if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")) {
		tmp := tdeWalletPassword.(string)
		result.NewTdeWalletPassword = &tmp
		oldTdePassword, _ := s.D.GetChange(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password"))
		tmp1 := oldTdePassword.(string)
		result.OldTdeWalletPassword = &tmp1
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

func (s *DatabaseDbHomeResourceCrud) mapToUpdateDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
	result := oci_database.DbBackupConfig{}

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
