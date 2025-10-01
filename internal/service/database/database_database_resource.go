// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

type expectedRetryDurationFn func(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, optionals ...interface{}) time.Duration

func DatabaseDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("2h"),
			Delete: tfresource.GetTimeoutDuration("2h"),
		},
		Create: createDatabaseDatabase,
		Read:   readDatabaseDatabase,
		Update: updateDatabaseDatabase,
		Delete: deleteDatabaseDatabase,
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

						// Optional
						"admin_password": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							//remove ForceNew field - TERSI-3772
							//ForceNew:  true,
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
						"character_set": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"database_admin_password": {
							Type:             schema.TypeString,
							Optional:         true,
							Sensitive:        true,
							Computed:         true,
							DiffSuppressFunc: dbAdminPasswordDiffSuppress,
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
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: disableAutoBackupSuppressfunc,
									},
									"auto_full_backup_day": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "SUNDAY",
									},
									"auto_full_backup_window": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: disableAutoBackupSuppressfunc,
									},
									"backup_deletion_policy": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"backup_destination_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"backup_retention_policy_on_terminate": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"dbrs_policy_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"is_remote": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_retention_lock_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"remote_region": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"type": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														string(oci_database.BackupBackupDestinationTypeAwsS3),
														string(oci_database.BackupBackupDestinationTypeDbrs),
														string(oci_database.BackupBackupDestinationTypeObjectStore),
														string(oci_database.BackupDestinationDetailsTypeNfs),
														string(oci_database.BackupDestinationDetailsTypeRecoveryAppliance),
														string(oci_database.BackupDestinationDetailsTypeLocal),
													}, true),
												},
												"vpc_user": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"vpc_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
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
									"run_immediate_full_backup": {
										Type:     schema.TypeBool,
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
						"db_unique_name": {
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
						"encryption_key_location_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"provider_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"AZURE",
											"EXTERNAL",
										}, true),
									},

									// Optional
									"azure_encryption_key_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"hsm_password": {
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},

									// Computed
								},
							},
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"is_active_data_guard_enabled": {
							Type:             schema.TypeBool,
							Optional:         true,
							DiffSuppressFunc: adgDiffSuppress,
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
						"pluggable_databases": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"protection_mode": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: protectionModeDiffSuppress,
						},
						"sid_prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"source_encryption_key_location_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"hsm_password": {
										Type:      schema.TypeString,
										Required:  true,
										Sensitive: true,
									},
									"provider_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"AZURE",
											"EXTERNAL",
										}, true),
									},

									// Optional

									// Computed
								},
							},
						},
						"source_database_id": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: sourceDbIdDiffSuppress,
						},
						"source_tde_wallet_password": {
							Type:             schema.TypeString,
							Optional:         true,
							Sensitive:        true,
							DiffSuppressFunc: sourceDbDetailsDiffSuppress,
						},
						"storage_size_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"data_storage_size_in_gb": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"reco_storage_size_in_gbs": {
										Type:     schema.TypeInt,
										Required: true,
									},

									// Optional

									// Computed
									"redo_log_storage_size_in_gbs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"tde_wallet_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							Sensitive: true,
						},
						"transport_type": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: transportTypeDiffSuppress,
						},
						"vault_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"db_home_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"DATAGUARD",
					"DB_BACKUP",
					"NONE",
				}, true),
			},

			// Optional
			"db_version": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DbVersionDiffSuppress,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kms_key_version_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kms_key_rotation": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"kms_key_migration": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"data_guard_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"character_set": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"data_guard_group": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"members": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"apply_lag": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"apply_rate": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"db_system_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_active_data_guard_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"transport_lag": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"transport_lag_refresh": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"transport_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"protection_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"database_management_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"management_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"database_software_image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_backup_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"auto_backup_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"auto_backup_window": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_full_backup_day": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_full_backup_window": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_deletion_policy": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_destination_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"backup_retention_policy_on_terminate": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"dbrs_policy_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_remote": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_retention_lock_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"remote_region": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vpc_user": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vpc_password": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"recovery_window_in_days": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"run_immediate_full_backup": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"db_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_system_id": {
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
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_cdb": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_backup_duration_in_seconds": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"key_store_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_store_wallet_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_backup_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_failed_backup_timestamp": {
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
			"sid_prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_database_point_in_time_recovery_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_size_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"data_storage_size_in_gb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"reco_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"redo_log_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vault_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func deleteDatabaseDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.Database
	DisableNotFoundRetries bool
}

func (s *DatabaseDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateProvisioning),
	}
}

func (s *DatabaseDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateAvailable),
		string(oci_database.DatabaseLifecycleStateBackupInProgress),
	}
}

func (s *DatabaseDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateTerminating),
	}
}

func (s *DatabaseDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseDatabaseResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateProvisioning),
		string(oci_database.DatabaseLifecycleStateUpdating),
		string(oci_database.DatabaseLifecycleStateUpgrading),
	}
}

func (s *DatabaseDatabaseResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseDatabaseResourceCrud) Create() error {
	request := oci_database.CreateDatabaseRequest{}
	err := s.populateTopLevelPolymorphicCreateDatabaseRequest(&request)
	if err != nil {
		return err
	}

	createDatabaseRetryDurationFn := getdatabaseRetryDurationFunction(s.D.Timeout(schema.TimeoutCreate))
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database", createDatabaseRetryDurationFn)

	response, err := s.Client.CreateDatabase(context.Background(), request)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	s.Res = &response.Database

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseDatabaseResourceCrud) Get() error {
	request := oci_database.GetDatabaseRequest{}

	tmp := s.D.Id()
	request.DatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Database
	return nil
}

func (s *DatabaseDatabaseResourceCrud) Delete() error {
	request := oci_database.DeleteDatabaseRequest{}

	tmp := s.D.Id()
	request.DatabaseId = &tmp

	if performFinalBackup, ok := s.D.GetOkExists("perform_final_backup"); ok {
		tmp := performFinalBackup.(bool)
		request.PerformFinalBackup = &tmp
	}

	deleteDatabaseRetryDurationFn := getdatabaseRetryDurationFunction(s.D.Timeout(schema.TimeoutDelete))
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database", deleteDatabaseRetryDurationFn)

	_, err := s.Client.DeleteDatabase(context.Background(), request)
	return err
}

func (s *DatabaseDatabaseResourceCrud) SetData() error {

	s.D.Set("database", []interface{}{s.DatabaseToMap(s.Res)})

	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{DatabaseConnectionStringsToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.DataGuardGroup != nil {
		s.D.Set("data_guard_group", []interface{}{DataGuardGroupToMap(s.Res.DataGuardGroup)})
	} else {
		s.D.Set("data_guard_group", nil)
	}

	if s.Res.DatabaseManagementConfig != nil {
		s.D.Set("database_management_config", []interface{}{CloudDatabaseManagementConfigToMap(s.Res.DatabaseManagementConfig)})
	} else {
		s.D.Set("database_management_config", nil)
	}

	if s.Res.DatabaseSoftwareImageId != nil {
		s.D.Set("database_software_image_id", *s.Res.DatabaseSoftwareImageId)
	}

	if s.Res.DbBackupConfig != nil {
		s.D.Set("db_backup_config", []interface{}{DbBackupConfigToMap(s.Res.DbBackupConfig)})
	} else {
		s.D.Set("db_backup_config", nil)
	}

	if s.Res.DbHomeId != nil {
		s.D.Set("db_home_id", *s.Res.DbHomeId)
	}

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.DbUniqueName != nil {
		s.D.Set("db_unique_name", *s.Res.DbUniqueName)
	}

	if s.Res.DbWorkload != nil {
		s.D.Set("db_workload", *s.Res.DbWorkload)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCdb != nil {
		s.D.Set("is_cdb", *s.Res.IsCdb)
	}

	if s.Res.KeyStoreId != nil {
		s.D.Set("key_store_id", *s.Res.KeyStoreId)
	}

	if s.Res.KeyStoreWalletName != nil {
		s.D.Set("key_store_wallet_name", *s.Res.KeyStoreWalletName)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.KmsKeyVersionId != nil {
		s.D.Set("kms_key_version_id", *s.Res.KmsKeyVersionId)
	}

	if s.Res.LastBackupDurationInSeconds != nil {
		s.D.Set("last_backup_duration_in_seconds", *s.Res.LastBackupDurationInSeconds)
	}

	if s.Res.LastBackupTimestamp != nil {
		s.D.Set("last_backup_timestamp", s.Res.LastBackupTimestamp.String())
	}

	if s.Res.LastFailedBackupTimestamp != nil {
		s.D.Set("last_failed_backup_timestamp", s.Res.LastFailedBackupTimestamp.String())
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

	if s.Res.SidPrefix != nil {
		s.D.Set("sid_prefix", *s.Res.SidPrefix)
	}

	if s.Res.SourceDatabasePointInTimeRecoveryTimestamp != nil {
		s.D.Set("source_database_point_in_time_recovery_timestamp", s.Res.SourceDatabasePointInTimeRecoveryTimestamp.String())
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageSizeDetails != nil {
		s.D.Set("storage_size_details", []interface{}{DatabaseStorageSizeResponseDetailsToMap(s.Res.StorageSizeDetails)})
	} else {
		s.D.Set("storage_size_details", nil)
	}

	s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	if s.Res.VmClusterId != nil {
		s.D.Set("vm_cluster_id", *s.Res.VmClusterId)
	}

	return nil
}

func (s *DatabaseDatabaseResourceCrud) ChangeEncryptionKeyLocation(fieldKeyFormat string) error {
	request := oci_database.ChangeEncryptionKeyLocationRequest{}

	idTmp := s.D.Id()
	request.DatabaseId = &idTmp

	if encryptionKeyLocationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_key_location_details")); ok {
		if tmpList := encryptionKeyLocationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "encryption_key_location_details"), 0)
			tmp, err := s.mapToEncryptionKeyLocationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return fmt.Errorf("unable to convert encryption_key_location_details, encountered error: %v", err)
			}
			request.EncryptionKeyLocationDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeEncryptionKeyLocation(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatabaseDatabaseResourceCrud) ChangeKeyStoreType() error {
	if _, ok := s.D.GetOkExists("key_store_id"); ok && s.D.HasChange("key_store_id") {
		request := oci_database.ChangeKeyStoreTypeRequest{}

		idTmp := s.D.Id()
		request.DatabaseId = &idTmp

		if keyStoreId, ok := s.D.GetOkExists("key_store_id"); ok {
			tmp := keyStoreId.(string)
			request.KeyStoreId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		_, err := s.Client.ChangeKeyStoreType(context.Background(), request)
		if err != nil {
			return err
		}

		if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
			return waitErr
		}

		return nil
	}
	return nil
}

func (s *DatabaseDatabaseResourceCrud) mapToBackupDestinationDetails(fieldKeyFormat string) (oci_database.BackupDestinationDetails, error) {
	result := oci_database.BackupDestinationDetails{}

	if backupRetentionPolicyOnTerminate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_retention_policy_on_terminate")); ok {
		result.BackupRetentionPolicyOnTerminate = oci_database.BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum(backupRetentionPolicyOnTerminate.(string))
	}

	if dbrsPolicyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dbrs_policy_id")); ok {
		tmp := dbrsPolicyId.(string)
		result.DbrsPolicyId = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if internetProxy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "internet_proxy")); ok {
		tmp := internetProxy.(string)
		result.InternetProxy = &tmp
	}

	if isRemote, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_remote")); ok {
		tmp := isRemote.(bool)
		result.IsRemote = &tmp
	}

	if isRetentionLockEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_retention_lock_enabled")); ok {
		tmp := isRetentionLockEnabled.(bool)
		result.IsRetentionLockEnabled = &tmp
	}

	if remoteRegion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remote_region")); ok {
		tmp := remoteRegion.(string)
		result.RemoteRegion = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_database.BackupDestinationDetailsTypeEnum(type_.(string))
	}

	if vpcPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpc_password")); ok {
		tmp := vpcPassword.(string)
		result.VpcPassword = &tmp
	}

	if vpcUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpc_user")); ok {
		tmp := vpcUser.(string)
		result.VpcUser = &tmp
	}

	return result, nil
}

func (s *DatabaseDatabaseResourceCrud) mapToUpdateBackupDestinationDetails(fieldKeyFormat string) (oci_database.BackupDestinationDetails, error) {
	result := oci_database.BackupDestinationDetails{}
	fields := map[string]func(string){
		"dbrs_policy_id": func(val string) { tmp := val; result.DbrsPolicyId = &tmp },
		"id":             func(val string) { tmp := val; result.Id = &tmp },
		"type":           func(val string) { result.Type = oci_database.BackupDestinationDetailsTypeEnum(val) },
		"vpc_password":   func(val string) { tmp := val; result.VpcPassword = &tmp },
		"vpc_user":       func(val string) { tmp := val; result.VpcUser = &tmp },
	}

	for fieldName, setter := range fields {
		key := fmt.Sprintf(fieldKeyFormat, fieldName)
		if val, ok := s.D.GetOkExists(key); ok {
			if s.D.HasChange(key) {
				setter(val.(string))
			} else {
				_, oldVal := s.D.GetChange(key)
				oldValStr := oldVal.(string)
				if oldValStr != "" {
					setter(oldValStr)
				}
			}
		}
	}

	// handle bool field
	if isRetentionLockEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_retention_lock_enabled")); ok {
		tmp := isRetentionLockEnabled.(bool)
		result.IsRetentionLockEnabled = &tmp
	}

	// handle enum field
	if backupRetentionPolicyOnTerminate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_retention_policy_on_terminate")); ok {
		result.BackupRetentionPolicyOnTerminate = oci_database.BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum(backupRetentionPolicyOnTerminate.(string))
	}

	return result, nil
}

func (s *DatabaseDatabaseResourceCrud) DatabaseBackupConfigToMap(obj *oci_database.DbBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AutoBackupEnabled != nil {
		result["auto_backup_enabled"] = bool(*obj.AutoBackupEnabled)
	}

	if obj.AutoBackupWindow != "" {
		result["auto_backup_window"] = string(obj.AutoBackupWindow)
	}

	result["auto_full_backup_day"] = string(obj.AutoFullBackupDay)

	if obj.AutoFullBackupWindow != "" {
		result["auto_full_backup_window"] = string(obj.AutoFullBackupWindow)
	}

	result["backup_deletion_policy"] = string(obj.BackupDeletionPolicy)

	backupDestinationDetails := []interface{}{}
	for _, item := range obj.BackupDestinationDetails {
		backupDestinationDetails = append(backupDestinationDetails, s.BackupDestinationDetailsToMap(item))
	}
	result["backup_destination_details"] = backupDestinationDetails

	if obj.RecoveryWindowInDays != nil {
		result["recovery_window_in_days"] = int(*obj.RecoveryWindowInDays)
	}

	if obj.RunImmediateFullBackup != nil {
		result["run_immediate_full_backup"] = bool(*obj.RunImmediateFullBackup)
	}

	return result
}

func (s *DatabaseDatabaseResourceCrud) BackupDestinationDetailsToMap(obj oci_database.BackupDestinationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["backup_retention_policy_on_terminate"] = string(obj.BackupRetentionPolicyOnTerminate)

	if obj.DbrsPolicyId != nil {
		result["dbrs_policy_id"] = string(*obj.DbrsPolicyId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InternetProxy != nil {
		result["internet_proxy"] = string(*obj.InternetProxy)
	}

	if obj.IsRemote != nil {
		result["is_remote"] = bool(*obj.IsRemote)
	}

	if obj.IsRetentionLockEnabled != nil {
		result["is_retention_lock_enabled"] = bool(*obj.IsRetentionLockEnabled)
	}

	if obj.RemoteRegion != nil {
		result["remote_region"] = string(*obj.RemoteRegion)
	}

	result["type"] = string(obj.Type)

	if vpcPassword, ok := s.D.GetOkExists("database.0.db_backup_config.0.backup_destination_details.0.vpc_password"); ok && vpcPassword != nil {
		result["vpc_password"] = vpcPassword.(string)
	}

	if obj.VpcUser != nil {
		result["vpc_user"] = string(*obj.VpcUser)
	}

	return result
}

func BackupDestinationDetailsToMap(obj oci_database.BackupDestinationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["backup_retention_policy_on_terminate"] = string(obj.BackupRetentionPolicyOnTerminate)

	if obj.DbrsPolicyId != nil {
		result["dbrs_policy_id"] = string(*obj.DbrsPolicyId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InternetProxy != nil {
		result["internet_proxy"] = string(*obj.InternetProxy)
	}

	if obj.IsRemote != nil {
		result["is_remote"] = bool(*obj.IsRemote)
	}

	if obj.IsRetentionLockEnabled != nil {
		result["is_retention_lock_enabled"] = bool(*obj.IsRetentionLockEnabled)
	}

	if obj.RemoteRegion != nil {
		result["remote_region"] = string(*obj.RemoteRegion)
	}

	result["type"] = string(obj.Type)

	if obj.VpcPassword != nil {
		result["vpc_password"] = string(*obj.VpcPassword)
	}

	if obj.VpcUser != nil {
		result["vpc_user"] = string(*obj.VpcUser)
	}

	return result
}

func CloudDatabaseManagementConfigToMap(obj *oci_database.CloudDatabaseManagementConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["management_status"] = string(obj.ManagementStatus)

	result["management_type"] = string(obj.ManagementType)

	return result
}

func (s *DatabaseDatabaseResourceCrud) mapToCreateDatabaseDetails(fieldKeyFormat string) (oci_database.CreateDatabaseDetails, error) {
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

	if dbUniqueName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_unique_name")); ok {
		tmp := dbUniqueName.(string)
		result.DbUniqueName = &tmp
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

	if encryptionKeyLocationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_key_location_details")); ok {
		if tmpList := encryptionKeyLocationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "encryption_key_location_details"), 0)
			tmp, err := s.mapToEncryptionKeyLocationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert encryption_key_location_details, encountered error: %v", err)
			}
			result.EncryptionKeyLocationDetails = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if keyStoreId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_id")); ok {
		tmp := keyStoreId.(string)
		result.KeyStoreId = &tmp
	}

	if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
		tmp := kmsKeyId.(string)
		result.KmsKeyId = &tmp
	}

	if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
		tmp := kmsKeyVersionId.(string)
		result.KmsKeyVersionId = &tmp
	}

	if ncharacterSet, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ncharacter_set")); ok {
		tmp := ncharacterSet.(string)
		result.NcharacterSet = &tmp
	}

	if pdbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pdb_name")); ok {
		tmp := pdbName.(string)
		result.PdbName = &tmp
	}

	if sidPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sid_prefix")); ok {
		tmp := sidPrefix.(string)
		result.SidPrefix = &tmp
	}

	if storageSizeDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_size_details")); ok {
		if tmpList := storageSizeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "storage_size_details"), 0)
			tmp, err := s.mapToDatabaseStorageSizeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert storage_size_details, encountered error: %v", err)
			}
			result.StorageSizeDetails = &tmp
		}
	}

	if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
		tmp := tdeWalletPassword.(string)
		result.TdeWalletPassword = &tmp
	}

	if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
		tmp := vaultId.(string)
		result.VaultId = &tmp
	}

	return result, nil
}

func (s *DatabaseDatabaseResourceCrud) mapToCreateDatabaseFromBackupDetails(fieldKeyFormat string) (oci_database.CreateDatabaseFromBackupDetails, error) {
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

	if dbUniqueName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_unique_name")); ok {
		tmp := dbUniqueName.(string)
		result.DbUniqueName = &tmp
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if pluggableDatabases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pluggable_databases")); ok {
		interfaces := pluggableDatabases.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "pluggable_databases")) {
			result.PluggableDatabases = tmp
		}
	}

	if sidPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sid_prefix")); ok {
		tmp := sidPrefix.(string)
		result.SidPrefix = &tmp
	}

	if sourceEncryptionKeyLocationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_encryption_key_location_details")); ok {
		if tmpList := sourceEncryptionKeyLocationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_encryption_key_location_details"), 0)
			tmp, err := s.mapToEncryptionKeyLocationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_encryption_key_location_details, encountered error: %v", err)
			}
			result.SourceEncryptionKeyLocationDetails = tmp
		}
	}

	if storageSizeDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_size_details")); ok {
		if tmpList := storageSizeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "storage_size_details"), 0)
			tmp, err := s.mapToDatabaseStorageSizeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert storage_size_details, encountered error: %v", err)
			}
			result.StorageSizeDetails = &tmp
		}
	}

	return result, nil
}

func (s *DatabaseDatabaseResourceCrud) mapToCreateStandbyDetails(fieldKeyFormat string) (oci_database.CreateStandbyDetails, error) {
	result := oci_database.CreateStandbyDetails{}

	if databaseAdminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_admin_password")); ok {
		tmp := databaseAdminPassword.(string)
		result.DatabaseAdminPassword = &tmp
	}

	if dbUniqueName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_unique_name")); ok {
		tmp := dbUniqueName.(string)
		result.DbUniqueName = &tmp
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isActiveDataGuardEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_active_data_guard_enabled")); ok {
		tmp := isActiveDataGuardEnabled.(bool)
		result.IsActiveDataGuardEnabled = &tmp
	}

	if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
		result.ProtectionMode = oci_database.CreateStandbyDetailsProtectionModeEnum(protectionMode.(string))
	}

	if sidPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sid_prefix")); ok {
		tmp := sidPrefix.(string)
		result.SidPrefix = &tmp
	}

	if sourceDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_database_id")); ok {
		tmp := sourceDatabaseId.(string)
		result.SourceDatabaseId = &tmp
	}

	if sourceTdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_tde_wallet_password")); ok {
		tmp := sourceTdeWalletPassword.(string)
		result.SourceTdeWalletPassword = &tmp
	}

	if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
		result.TransportType = oci_database.CreateStandbyDetailsTransportTypeEnum(transportType.(string))
	}

	if sourceEncryptionKeyLocationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_encryption_key_location_details")); ok {
		if tmpList := sourceEncryptionKeyLocationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_encryption_key_location_details"), 0)
			tmp, err := s.mapToEncryptionKeyLocationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_encryption_key_location_details, encountered error: %v", err)
			}
			result.SourceEncryptionKeyLocationDetails = tmp
		}
	}

	if storageSizeDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_size_details")); ok {
		if tmpList := storageSizeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "storage_size_details"), 0)
			tmp, err := s.mapToDatabaseStorageSizeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert storage_size_details, encountered error: %v", err)
			}
			result.StorageSizeDetails = &tmp
		}
	}

	return result, nil
}

func DataGuardGroupToMap(obj *oci_database.DataGuardGroup) map[string]interface{} {
	result := map[string]interface{}{}

	members := []interface{}{}
	for _, item := range obj.Members {
		members = append(members, DataGuardGroupMemberToMap(item))
	}
	result["members"] = members

	result["protection_mode"] = string(obj.ProtectionMode)

	return result
}

func DataGuardGroupMemberToMap(obj oci_database.DataGuardGroupMember) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplyLag != nil {
		result["apply_lag"] = string(*obj.ApplyLag)
	}

	if obj.ApplyRate != nil {
		result["apply_rate"] = string(*obj.ApplyRate)
	}

	if obj.DatabaseId != nil {
		result["database_id"] = string(*obj.DatabaseId)
	}

	if obj.DbSystemId != nil {
		result["db_system_id"] = string(*obj.DbSystemId)
	}

	if obj.IsActiveDataGuardEnabled != nil {
		result["is_active_data_guard_enabled"] = bool(*obj.IsActiveDataGuardEnabled)
	}

	result["role"] = string(obj.Role)

	if obj.TransportLag != nil {
		result["transport_lag"] = string(*obj.TransportLag)
	}

	if obj.TransportLagRefresh != nil {
		result["transport_lag_refresh"] = string(*obj.TransportLagRefresh)
	}

	result["transport_type"] = string(obj.TransportType)

	return result
}

func DatabaseConnectionStringsToMap(obj *oci_database.DatabaseConnectionStrings) map[string]interface{} {
	result := map[string]interface{}{}

	result["all_connection_strings"] = obj.AllConnectionStrings

	if obj.CdbDefault != nil {
		result["cdb_default"] = string(*obj.CdbDefault)
	}

	if obj.CdbIpDefault != nil {
		result["cdb_ip_default"] = string(*obj.CdbIpDefault)
	}

	return result
}

func (s *DatabaseDatabaseResourceCrud) mapToDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
	result := oci_database.DbBackupConfig{}

	if autoBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")); ok {
		tmp := autoBackupEnabled.(bool)
		result.AutoBackupEnabled = &tmp
	}

	if autoBackupWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_window")); ok {
		result.AutoBackupWindow = oci_database.DbBackupConfigAutoBackupWindowEnum(autoBackupWindow.(string))
	}

	if autoFullBackupDay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_full_backup_day")); ok {
		result.AutoFullBackupDay = oci_database.DbBackupConfigAutoFullBackupDayEnum(autoFullBackupDay.(string))
	}

	if autoFullBackupWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_full_backup_window")); ok {
		result.AutoFullBackupWindow = oci_database.DbBackupConfigAutoFullBackupWindowEnum(autoFullBackupWindow.(string))
	}

	if backupDeletionPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_deletion_policy")); ok {
		result.BackupDeletionPolicy = oci_database.DbBackupConfigBackupDeletionPolicyEnum(backupDeletionPolicy.(string))
	}

	if backupDestinationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")); ok {
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

	if runImmediateFullBackup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_immediate_full_backup")); ok {
		tmp := runImmediateFullBackup.(bool)
		result.RunImmediateFullBackup = &tmp
	}

	return result, nil
}

func (s *DatabaseDatabaseResourceCrud) mapToEncryptionKeyLocationDetails(fieldKeyFormat string) (oci_database.EncryptionKeyLocationDetails, error) {
	var baseObject oci_database.EncryptionKeyLocationDetails
	//discriminator
	providerTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "provider_type"))
	var providerType string
	if ok {
		providerType = providerTypeRaw.(string)
	} else {
		providerType = "" // default value
	}
	switch strings.ToLower(providerType) {
	case strings.ToLower("AZURE"):
		details := oci_database.AzureEncryptionKeyDetails{}
		if azureEncryptionKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "azure_encryption_key_id")); ok {
			tmp := azureEncryptionKeyId.(string)
			details.AzureEncryptionKeyId = &tmp
		}
		baseObject = details
	case strings.ToLower("EXTERNAL"):
		details := oci_database.ExternalHsmEncryptionDetails{}
		if hsmPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hsm_password")); ok {
			tmp := hsmPassword.(string)
			details.HsmPassword = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown provider_type '%v' was specified", providerType)
	}
	return baseObject, nil
}

func EncryptionKeyLocationDetailsToMap(obj *oci_database.EncryptionKeyLocationDetails, hsmPassword string) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database.AzureEncryptionKeyDetails:
		result["provider_type"] = "AZURE"

		if v.AzureEncryptionKeyId != nil {
			result["azure_encryption_key_id"] = string(*v.AzureEncryptionKeyId)
		}
	case oci_database.ExternalHsmEncryptionDetails:
		result["provider_type"] = "EXTERNAL"
		result["hsm_password"] = hsmPassword

	default:
		log.Printf("[WARN] Received 'provider_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseDatabaseResourceCrud) populateTopLevelPolymorphicCreateDatabaseRequest(request *oci_database.CreateDatabaseRequest) error {
	//discriminator
	sourceRaw, ok := s.D.GetOkExists("source")
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "NONE" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("DATAGUARD"):
		details := oci_database.CreateStandByDatabaseDetails{}
		if database, ok := s.D.GetOkExists("database"); ok {
			if tmpList := database.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
				tmp, err := s.mapToCreateStandbyDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Database = &tmp
			}
		}
		if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok {
			tmp := dbHomeId.(string)
			details.DbHomeId = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		request.CreateNewDatabaseDetails = details
	case strings.ToLower("DB_BACKUP"):
		details := oci_database.CreateDatabaseFromBackup{}
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
		if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok {
			tmp := dbHomeId.(string)
			details.DbHomeId = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		request.CreateNewDatabaseDetails = details
	case strings.ToLower("NONE"):
		details := oci_database.CreateNewDatabaseDetails{}
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
		if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok {
			tmp := dbHomeId.(string)
			details.DbHomeId = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		request.CreateNewDatabaseDetails = details
	default:
		return fmt.Errorf("unknown source '%v' was specified", source)
	}
	return nil
}

func updateDatabaseDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func (s *DatabaseDatabaseResourceCrud) Update() error {
	request := oci_database.UpdateDatabaseRequest{}

	tmp := s.D.Id()
	request.DatabaseId = &tmp

	//fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)

	if _, ok := s.D.GetOkExists("action_trigger"); ok && s.D.HasChange("action_trigger") {
		oldRaw, newRaw := s.D.GetChange("action_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			if dataguardActionType, ok := s.D.GetOkExists("data_guard_action"); ok {
				action := dataguardActionType.(string)
				if strings.EqualFold(action, "switchover") {
					err := s.switchoverAction(tmp)
					if err != nil {
						return err
					}
				}
				if strings.EqualFold(action, "failover") {
					err := s.failoverAction(tmp)
					if err != nil {
						return err
					}
				}
				if strings.EqualFold(action, "reinstate") {
					err := s.reintateAction(tmp)
					if err != nil {
						return err
					}
				}
				if strings.EqualFold(action, "convertToStandalone") {
					err := s.convertToStandaloneAction(tmp)
					if err != nil {
						return err
					}
				}
				if strings.EqualFold(action, "dgConfig") {
					err := s.dataGuardConfigUpdate(tmp)
					if err != nil {
						return err
					}
				}

			}
		} else {
			s.D.Set("action_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}

	}

	error := s.kmsRotation(tmp)
	if error != nil {
		return error
	}

	err := s.kmsMigration(tmp)
	if err != nil {
		return err
	}

	if s.D.HasChange("key_store_id") {
		oldVal, confVal := s.D.GetChange("key_store_id")
		if oldVal.(string) != "" && confVal.(string) != "" {
			return fmt.Errorf("[ERROR] no support for oldVal = '%s', confVal = '%s' now", oldVal.(string), confVal.(string))
		}
		if oldVal.(string) == "" {
			errExaCC := s.ChangeKeyStoreType()
			if errExaCC != nil {
				return errExaCC
			}
		}
		if confVal.(string) == "" && oldVal.(string) != "" {
			return fmt.Errorf("[ERROR] no support for migrate to Oracle now")
		}
	}
	errKms := s.setDbKeyVersion(tmp)
	if errKms != nil {
		return errKms
	}

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

	if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok && s.D.HasChange("db_home_id") {
		tmp := dbHomeId.(string)
		request.DbHomeId = &tmp
	}

	updateDatabaseRetryDurationFn := getdatabaseRetryDurationFunction(s.D.Timeout(schema.TimeoutUpdate))
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database", updateDatabaseRetryDurationFn)

	response, err := s.Client.UpdateDatabase(context.Background(), request)
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

	s.Res = &response.Database
	return nil
}

func (s *DatabaseDatabaseResourceCrud) mapToUpdateDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
	result := oci_database.DbBackupConfig{}

	if autoBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")); ok {
		tmp := autoBackupEnabled.(bool)
		result.AutoBackupEnabled = &tmp
	}

	if autoBackupWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_window")); ok {
		result.AutoBackupWindow = oci_database.DbBackupConfigAutoBackupWindowEnum(autoBackupWindow.(string))
	}

	if autoFullBackupDay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_full_backup_day")); ok {
		result.AutoFullBackupDay = oci_database.DbBackupConfigAutoFullBackupDayEnum(autoFullBackupDay.(string))
	}

	if autoFullBackupWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_full_backup_window")); ok {
		result.AutoFullBackupWindow = oci_database.DbBackupConfigAutoFullBackupWindowEnum(autoFullBackupWindow.(string))
	}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")) {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	if runImmediateFullBackup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_immediate_full_backup")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "run_immediate_full_backup")) {
		tmp := runImmediateFullBackup.(bool)
		result.RunImmediateFullBackup = &tmp
	}

	if backup_deletion_policy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_deletion_policy")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backup_deletion_policy")) {
		result.BackupDeletionPolicy = oci_database.DbBackupConfigBackupDeletionPolicyEnum(backup_deletion_policy.(string))
	}

	if backupDestinationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")); ok {
		interfaces := backupDestinationDetails.([]interface{})
		tmp := make([]oci_database.BackupDestinationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backup_destination_details"), stateDataIndex)
			converted, err := s.mapToUpdateBackupDestinationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")) {
			result.BackupDestinationDetails = tmp
		}
	}

	return result, nil
}

func (s *DatabaseDatabaseResourceCrud) mapToUpdateDatabaseDetails(fieldKeyFormat string) (oci_database.UpdateDatabaseDetails, error) {
	result := oci_database.UpdateDatabaseDetails{}

	if _, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_key_location_details")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "encryption_key_location_details")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf(fieldKeyFormat, "encryption_key_location_details"))
		oldList := oldRaw.([]interface{})
		newList := newRaw.([]interface{})

		if len(oldList) > 0 && len(newList) > 0 {
			return result, fmt.Errorf("[ERROR] no support for updating External HSM now")
		}
		if len(oldList) == 0 {
			err := s.ChangeEncryptionKeyLocation(fieldKeyFormat)
			if err != nil {
				return result, err
			}
		}
		if len(newList) == 0 && len(oldList) > 0 {
			return result, fmt.Errorf("[ERROR] no support for migrate from External HSM now")
		}
	}

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
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "admin_password")) {
		tmp := adminPassword.(string)
		result.NewAdminPassword = &tmp
	}

	if storageSizeDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_size_details")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "storage_size_details")) {
		if tmpList := storageSizeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "storage_size_details"), 0)
			tmp, err := s.mapToDatabaseStorageSizeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			result.StorageSizeDetails = &tmp
		}
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

func (s *DatabaseDatabaseResourceCrud) DatabaseToMap(obj *oci_database.Database) map[string]interface{} {
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

	if backupTDEPassword, ok := s.D.GetOkExists("database.0.backup_tde_password"); ok && backupTDEPassword != nil {
		result["backup_tde_password"] = backupTDEPassword.(string)
	}

	if obj.CharacterSet != nil {
		result["character_set"] = string(*obj.CharacterSet)
	}

	if obj.DatabaseManagementConfig != nil {
		result["database_management_config"] = []interface{}{CloudDatabaseManagementConfigToMap(obj.DatabaseManagementConfig)}
	}

	if obj.DbBackupConfig != nil {
		result["db_backup_config"] = []interface{}{s.DatabaseBackupConfigToMap(obj.DbBackupConfig)}
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
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.NcharacterSet != nil {
		result["ncharacter_set"] = string(*obj.NcharacterSet)
	}

	if obj.PdbName != nil {
		result["pdb_name"] = string(*obj.PdbName)
	}

	if obj.SidPrefix != nil {
		result["sid_prefix"] = string(*obj.SidPrefix)
	}

	if obj.StorageSizeDetails != nil {
		result["storage_size_details"] = []interface{}{DatabaseStorageSizeResponseDetailsToMap(obj.StorageSizeDetails)}
	} else {
		result["storage_size_details"] = nil
	}

	if hsmPassword, ok := s.D.GetOkExists("database.0.encryption_key_location_details.0.hsm_password"); ok && hsmPassword != nil {
		if s.Res.EncryptionKeyLocationDetails != nil {
			result["encryption_key_location_details"] = []interface{}{EncryptionKeyLocationDetailsToMap(&s.Res.EncryptionKeyLocationDetails, hsmPassword.(string))}
		}
	}

	if sourceHsmPassword, ok := s.D.GetOkExists("database.0.source_encryption_key_location_details.0.hsm_password"); ok && sourceHsmPassword != nil {
		if s.Res.EncryptionKeyLocationDetails != nil {
			result["source_encryption_key_location_details"] = []interface{}{EncryptionKeyLocationDetailsToMap(&s.Res.EncryptionKeyLocationDetails, sourceHsmPassword.(string))}
		}
	}

	return result
}

func (s *DatabaseDatabaseResourceCrud) setDbKeyVersion(databaseId string) error {
	if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok && s.D.HasChange("kms_key_version_id") {
		oldRaw, newRaw := s.D.GetChange("kms_key_version_id")
		if oldRaw == "" && newRaw != "" {
			setDbKeyVersionRequest := oci_database.SetDbKeyVersionRequest{}
			setDbKeyVersionRequest.DatabaseId = &databaseId
			setDbKeyVersionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
			details := oci_database.OciProviderSetKeyVersionDetails{}

			temp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &temp
			setDbKeyVersionRequest.SetKeyVersionDetails = details

			response, err := s.Client.SetDbKeyVersion(context.Background(), setDbKeyVersionRequest)
			if err != nil {
				return err
			}
			workId := response.OpcWorkRequestId
			if workId != nil {
				_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
				if err != nil {
				}
			}
		}
	}
	return nil
}

func (s *DatabaseDatabaseResourceCrud) switchoverAction(databaseId string) error {
	switchoverDataGuardRequest := oci_database.SwitchOverDataGuardRequest{}
	switchoverDetails := oci_database.SwitchOverDataGuardDetails{}
	switchoverDataGuardRequest.DatabaseId = &databaseId
	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_admin_password")); ok {
		tmp := adminPassword.(string)
		switchoverDetails.DatabaseAdminPassword = &tmp
	}
	switchoverDataGuardRequest.SwitchOverDataGuardDetails = switchoverDetails
	response, err := s.Client.SwitchOverDataGuard(context.Background(), switchoverDataGuardRequest)
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
	val := s.D.Get("action_trigger")
	s.D.Set("action_trigger", val)
	val2 := s.D.Get("data_guard_action")
	s.D.Set("data_guard_action", val2)
	return nil
}

func (s *DatabaseDatabaseResourceCrud) failoverAction(databaseId string) error {
	failoverDataGuardRequest := oci_database.FailoverDataGuardRequest{}
	failoverDataGuardRequest.DatabaseId = &databaseId
	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_admin_password")); ok {
		tmp := adminPassword.(string)
		failoverDataGuardRequest.FailoverDataGuardDetails.DatabaseAdminPassword = &tmp
	}
	response, err := s.Client.FailoverDataGuard(context.Background(), failoverDataGuardRequest)
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
	val := s.D.Get("action_trigger")
	s.D.Set("action_trigger", val)
	val2 := s.D.Get("data_guard_action")
	s.D.Set("data_guard_action", val2)
	return nil
}

func (s *DatabaseDatabaseResourceCrud) reintateAction(databaseId string) error {
	reinstateDataGuardRequest := oci_database.ReinstateDataGuardRequest{}
	reinstateDataGuardRequest.DatabaseId = &databaseId
	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_admin_password")); ok {
		tmp := adminPassword.(string)
		reinstateDataGuardRequest.ReinstateDataGuardDetails.DatabaseAdminPassword = &tmp
	}
	if sourceDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_database_id")); ok {
		tmp := sourceDatabaseId.(string)
		reinstateDataGuardRequest.ReinstateDataGuardDetails.SourceDatabaseId = &tmp
	}
	response, err := s.Client.ReinstateDataGuard(context.Background(), reinstateDataGuardRequest)
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
	val := s.D.Get("action_trigger")
	s.D.Set("action_trigger", val)
	val2 := s.D.Get("data_guard_action")
	s.D.Set("data_guard_action", val2)
	return nil
}

func (s *DatabaseDatabaseResourceCrud) convertToStandaloneAction(databaseId string) error {
	convertToStandaloneRequest := oci_database.ConvertToStandaloneRequest{}
	convertToStandaloneRequest.DatabaseId = &databaseId
	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_admin_password")); ok {
		tmp := adminPassword.(string)
		convertToStandaloneRequest.ConvertToStandaloneDetails.DatabaseAdminPassword = &tmp
	}
	response, err := s.Client.ConvertToStandalone(context.Background(), convertToStandaloneRequest)
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
	val := s.D.Get("action_trigger")
	s.D.Set("action_trigger", val)
	val2 := s.D.Get("data_guard_action")
	s.D.Set("data_guard_action", val2)
	return nil
}

func (s *DatabaseDatabaseResourceCrud) dataGuardConfigUpdate(databaseId string) error {
	updateDataGuardRequest := oci_database.UpdateDataGuardRequest{}
	updateDataGuardRequest.DatabaseId = &databaseId
	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_admin_password")); ok {
		tmp := adminPassword.(string)
		updateDataGuardRequest.UpdateDataGuardDetails.DatabaseAdminPassword = &tmp
	}
	if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
		updateDataGuardRequest.UpdateDataGuardDetails.TransportType = oci_database.UpdateDataGuardDetailsTransportTypeEnum(transportType.(string))
	}
	if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
		updateDataGuardRequest.UpdateDataGuardDetails.ProtectionMode = oci_database.UpdateDataGuardDetailsProtectionModeEnum(protectionMode.(string))
	}
	if activeDgEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_active_data_guard_enabled")); ok {
		tmp := activeDgEnabled.(bool)
		updateDataGuardRequest.UpdateDataGuardDetails.IsActiveDataGuardEnabled = &tmp
	}
	response, err := s.Client.UpdateDataGuard(context.Background(), updateDataGuardRequest)
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
	val := s.D.Get("action_trigger")
	s.D.Set("action_trigger", val)
	val2 := s.D.Get("data_guard_action")
	s.D.Set("data_guard_action", val2)
	return nil
}

func (s *DatabaseDatabaseResourceCrud) kmsRotation(databaseId string) error {
	if _, ok := s.D.GetOkExists("kms_key_rotation"); ok && s.D.HasChange("kms_key_rotation") {
		rotateKeyRequest := oci_database.RotateVaultKeyRequest{}
		rotateKeyRequest.DatabaseId = &databaseId
		rotateKeyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
		response, err := s.Client.RotateVaultKey(context.Background(), rotateKeyRequest)
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
		return nil
	}
	return nil
}

func (s *DatabaseDatabaseResourceCrud) kmsMigration(databaseId string) error {
	migrateOperation := false
	if _, ok := s.D.GetOkExists("kms_key_migration"); ok && s.D.HasChange("kms_key_migration") && s.D.Get("kms_key_migration").(bool) {
		migrationKeyRequest := oci_database.MigrateVaultKeyRequest{}
		migrationKeyRequest.DatabaseId = &databaseId
		migrationKeyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok && s.D.HasChange("kms_key_id") {
			oldRaw, newRaw := s.D.GetChange("kms_key_id")
			if oldRaw == "" && newRaw != "" {
				temp := kmsKeyId.(string)
				migrationKeyRequest.KmsKeyId = &temp
			}
		}

		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok && s.D.HasChange("kms_key_version_id") {
			oldRaw, newRaw := s.D.GetChange("kms_key_version_id")
			if oldRaw == "" && newRaw != "" {
				temp := kmsKeyVersionId.(string)
				migrationKeyRequest.KmsKeyVersionId = &temp
			}
		}
		response, err := s.Client.MigrateVaultKey(context.Background(), migrationKeyRequest)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
			if err != nil {
			}
		}
		migrateOperation = true
	}

	if _, ok := s.D.GetOkExists("kms_key_id"); ok && s.D.HasChange("kms_key_id") && !migrateOperation {
		return errors.New("kms_key_id can not be updated, please use migration or rotation")
	}

	return nil
}

func getdatabaseRetryDurationFunction(retryTimeout time.Duration) expectedRetryDurationFn {
	return func(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, optionals ...interface{}) time.Duration {
		defaultRetryTime := tfresource.GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)
		if response.Response == nil || response.Response.HTTPResponse() == nil {
			return defaultRetryTime
		}
		e := response.Error
		switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
		case 409:
			if isDisable409Retry, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("disable_409_retry", "false")); isDisable409Retry {
				log.Printf("[ERROR] Resource is in conflict state due to multiple update request: %v", e.Error())
				return 0
			}
			if e := response.Error; e != nil {
				if strings.Contains(e.Error(), "IncorrectState") {
					defaultRetryTime = retryTimeout
				} else if strings.Contains(e.Error(), "InvalidatedRetryToken") {
					defaultRetryTime = 0
				} else {
					defaultRetryTime = tfresource.LongRetryTime
				}
			}
		}
		return defaultRetryTime
	}
}

func protectionModeDiffSuppress(k string, old, new string, d *schema.ResourceData) bool {
	if v, ok := d.GetOk("data_guard_group"); ok {
		if list, ok := v.([]interface{}); ok && len(list) > 0 {
			if firstMap, ok := list[0].(map[string]interface{}); ok {
				if fieldValue, exists := firstMap["protection_mode"]; exists {
					return fieldValue.(string) == new
				}
			}
		}
	}
	return false
}

func dbAdminPasswordDiffSuppress(k string, old, new string, d *schema.ResourceData) bool {
	if _, ok := d.GetOkExists("action_trigger"); ok && d.HasChange("action_trigger") {
		return false
	}

	if actionType, ok := d.GetOkExists("data_guard_action"); ok && strings.EqualFold(actionType.(string), "convertToStandalone") {
		return true
	}

	dgGroup, ok := d.GetOk("data_guard_group")
	if !ok {
		return false
	}
	list, ok := dgGroup.([]interface{})
	if !ok && !(len(list) > 0) {
		return false
	}
	firstMap, ok := list[0].(map[string]interface{})
	if !ok {
		return false
	}
	members, ok := firstMap["members"].([]interface{})
	if !ok && !(len(members) > 0) {
		return false
	}
	for _, member := range members {
		item, ok := member.(map[string]interface{})
		if ok {
			if databaseId, exists := item["database_id"]; exists && databaseId.(string) == d.Id() {
				return true
			}
		}
	}
	return false
}

func transportTypeDiffSuppress(k string, old, new string, d *schema.ResourceData) bool {
	dgGroup, ok := d.GetOk("data_guard_group")
	if !ok {
		return false
	}
	list, ok := dgGroup.([]interface{})
	if !ok && !(len(list) > 0) {
		return false
	}
	firstMap, ok := list[0].(map[string]interface{})
	if !ok {
		return false
	}
	members, ok := firstMap["members"].([]interface{})
	if !ok && !(len(members) > 0) {
		return false
	}
	for _, member := range members {
		item, ok := member.(map[string]interface{})
		if ok {
			if databaseId, exists := item["database_id"]; exists && databaseId.(string) == d.Id() {
				if fieldValue, exists := item["transport_type"]; exists {
					return fieldValue.(string) == new
				}
			}
		}
	}
	return false
}

func sourceDbIdDiffSuppress(k string, old, new string, d *schema.ResourceData) bool {
	if _, ok := d.GetOkExists("action_trigger"); ok && d.HasChange("action_trigger") {
		if dataguardActionType, ok := d.GetOkExists("data_guard_action"); ok {
			action := dataguardActionType.(string)
			if strings.EqualFold(action, "reinstate") {
				return false
			}
		}
	}
	return sourceDbDetailsDiffSuppress(k, old, new, d)
}

func sourceDbDetailsDiffSuppress(k string, old, new string, d *schema.ResourceData) bool {
	dgGroup, ok := d.GetOk("data_guard_group")
	if !ok {
		return false
	}
	list, ok := dgGroup.([]interface{})
	if !ok && !(len(list) > 0) {
		return false
	}
	firstMap, ok := list[0].(map[string]interface{})
	if !ok {
		return false
	}
	members, ok := firstMap["members"].([]interface{})
	if !ok && !(len(members) > 0) {
		return false
	}
	for _, member := range members {
		item, ok := member.(map[string]interface{})
		if ok {
			if databaseId, exists := item["database_id"]; exists && databaseId.(string) == d.Id() {
				return true
			}
		}
	}
	return false
}

func adgDiffSuppress(k string, old, new string, d *schema.ResourceData) bool {
	dgGroup, ok := d.GetOk("data_guard_group")
	if !ok {
		return false
	}
	list, ok := dgGroup.([]interface{})
	if !ok && !(len(list) > 0) {
		return false
	}
	firstMap, ok := list[0].(map[string]interface{})
	if !ok {
		return false
	}
	members, ok := firstMap["members"].([]interface{})
	if !ok && !(len(members) > 0) {
		return false
	}
	for _, member := range members {
		item, ok := member.(map[string]interface{})
		if ok {
			if databaseId, exists := item["database_id"]; exists && databaseId.(string) == d.Id() {
				// To suppress diff on the primary database for the field "is_active_data_guard_enabled" because during
				//role change with the standby (failover/switchover), there is a diff in the primary db if the standby db is created with ADG value as true.
				if dgRoleValue, exists := item["role"]; exists && dgRoleValue.(string) == "PRIMARY" {
					return true
				}
				if fieldValue, exists := item["is_active_data_guard_enabled"]; exists {
					return compareBooleanField(fieldValue, new)
				}
			}
		}
	}
	return false
}

func compareBooleanField(fieldValue interface{}, new string) bool {
	boolValue, _ := fieldValue.(bool)

	newBool, err := strconv.ParseBool(new)
	if err != nil {
		if new == "" {
			// Empty string case
			return !boolValue
		}
		return false
	}
	return boolValue == newBool
}

func (s *DatabaseDatabaseResourceCrud) mapToDatabaseStorageSizeDetails(fieldKeyFormat string) (oci_database.DatabaseStorageSizeDetails, error) {
	result := oci_database.DatabaseStorageSizeDetails{}

	if dataStorageSizeInGB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_gb")); ok {
		tmp := dataStorageSizeInGB.(int)
		result.DataStorageSizeInGBs = &tmp
	}

	if recoStorageSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reco_storage_size_in_gbs")); ok {
		tmp := recoStorageSizeInGBs.(int)
		result.RecoStorageSizeInGBs = &tmp
	}

	return result, nil
}

func DatabaseStorageSizeResponseDetailsToMap(obj *oci_database.DatabaseStorageSizeResponseDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DataStorageSizeInGBs != nil {
		result["data_storage_size_in_gb"] = int(*obj.DataStorageSizeInGBs)
	}

	if obj.RecoStorageSizeInGBs != nil {
		result["reco_storage_size_in_gbs"] = int(*obj.RecoStorageSizeInGBs)
	}

	if obj.RedoLogStorageSizeInGBs != nil {
		result["redo_log_storage_size_in_gbs"] = int(*obj.RedoLogStorageSizeInGBs)
	}

	return result
}
