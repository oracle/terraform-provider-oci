// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v56/workrequests"
)

func DatabaseAutonomousDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		Create: createDatabaseAutonomousDatabase,
		Read:   readDatabaseAutonomousDatabase,
		Update: updateDatabaseAutonomousDatabase,
		Delete: deleteDatabaseAutonomousDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"admin_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"are_primary_whitelisted_ips_used": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"autonomous_database_backup_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"autonomous_maintenance_schedule_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clone_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"customer_contacts": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"email": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"data_safe_status": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_database.AutonomousDatabaseDataSafeStatusRegistered),
					string(oci_database.AutonomousDatabaseSummaryDataSafeStatusNotRegistered),
				}, true),
			},
			"data_storage_size_in_gb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_version": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DbVersionDiffSuppress,
			},
			"db_workload": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_access_control_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_auto_scaling_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_data_guard_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_dedicated": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_free_tier": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_mtls_connection_required": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_preview_version_with_service_terms_accepted": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_refreshable_clone": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"license_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      utils.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ocpu_count": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"private_endpoint_label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refreshable_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_database.CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeManual),
					string(oci_database.CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeAutomatic),
					string(oci_database.UpdateAutonomousDatabaseDetailsRefreshableModeAutomatic),
					string(oci_database.UpdateAutonomousDatabaseDetailsRefreshableModeManual),
				}, false),
			},
			"scheduled_operations": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"day_of_week": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional
						"scheduled_start_time": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scheduled_stop_time": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"source": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"BACKUP_FROM_ID",
					"BACKUP_FROM_TIMESTAMP",
					"CLONE_TO_REFRESHABLE",
					"DATABASE",
					"NONE",
				}, true),
			},
			"source_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"standby_whitelisted_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timestamp": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: utils.TimeDiffSuppressFunction,
			},
			"vault_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"whitelisted_ips": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      utils.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"switchover_to": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"PRIMARY",
					"STANDBY",
				}, true),
			},
			"rotate_key_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// Computed
			"apex_details": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"apex_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ords_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"available_upgrade_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"backup_config": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"manual_backup_bucket_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"manual_backup_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"connection_strings": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
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
						"dedicated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"high": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"low": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"medium": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"profiles": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"consumer_group": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"host_format": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"session_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"syntax_format": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tls_authentication": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"connection_urls": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"apex_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"graph_studio_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"machine_learning_user_management_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sql_dev_web_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"database_management_status": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"failed_data_recovery_in_seconds": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"infrastructure_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_preview": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_reconnect_clone_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"key_history_entry": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kms_key_version_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_activated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vault_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"key_store_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_store_wallet_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kms_key_lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kms_key_version_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"open_mode": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"operations_insights_status": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"permission_level": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"private_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_endpoint_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"refreshable_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_console_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"standby_db": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"lag_time_in_seconds": {
							Type:     schema.TypeInt,
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
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_database.AutonomousDatabaseLifecycleStateStopped),
					string(oci_database.AutonomousDatabaseLifecycleStateAvailable),
				}, true),
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
			"time_deletion_of_free_autonomous_database": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_maintenance_begin": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_maintenance_end": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_last_failover": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_last_refresh": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_last_refresh_point": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_last_switchover": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_next_refresh": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_reclamation_of_free_autonomous_database": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_until_reconnect_clone_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"used_data_storage_size_in_tbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	configDataSafeStatus := oci_database.AutonomousDatabaseDataSafeStatusNotRegistered
	if dataSafeStatus, ok := sync.D.GetOkExists("data_safe_status"); ok {
		configDataSafeStatus = oci_database.AutonomousDatabaseDataSafeStatusEnum(strings.ToUpper(dataSafeStatus.(string)))
	}

	configOpenMode := oci_database.UpdateAutonomousDatabaseDetailsOpenModeWrite
	if openMode, ok := sync.D.GetOkExists("open_mode"); ok {
		configOpenMode = oci_database.UpdateAutonomousDatabaseDetailsOpenModeEnum(openMode.(string))
	}

	configPermissionLevel := oci_database.UpdateAutonomousDatabaseDetailsPermissionLevelUnrestricted
	if permissionLevel, ok := sync.D.GetOkExists("permission_level"); ok {
		configPermissionLevel = oci_database.UpdateAutonomousDatabaseDetailsPermissionLevelEnum(permissionLevel.(string))
	}

	var isInactiveRequest = false
	if configState, ok := sync.D.GetOkExists("state"); ok {
		wantedState := oci_database.AutonomousDatabaseLifecycleStateEnum(strings.ToUpper(configState.(string)))
		if wantedState == oci_database.AutonomousDatabaseLifecycleStateStopped {
			isInactiveRequest = true
		}
	}

	configOperationsInsightsStatus := oci_database.AutonomousDatabaseOperationsInsightsStatusNotEnabled
	if operationsInsightsStatus, ok := sync.D.GetOkExists("operations_insights_status"); ok {
		configOperationsInsightsStatus = oci_database.AutonomousDatabaseOperationsInsightsStatusEnum((operationsInsightsStatus.(string)))
	}

	configDatabaseManagementStatus := oci_database.AutonomousDatabaseDatabaseManagementStatusNotEnabled
	if databaseManagementStatus, ok := sync.D.GetOkExists("database_management_status"); ok {
		configDatabaseManagementStatus = oci_database.AutonomousDatabaseDatabaseManagementStatusEnum((databaseManagementStatus.(string)))
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("rotate_key_trigger"); ok {
		err := sync.RotateAutonomousDatabaseEncryptionKey()
		if err != nil {
			return err
		}
	}

	if isInactiveRequest {
		return inactiveAutonomousDatabaseIfNeeded(d, sync)
	}

	if configDataSafeStatus == oci_database.AutonomousDatabaseDataSafeStatusRegistered {
		err := sync.updateDataSafeStatus(sync.D.Id(), oci_database.AutonomousDatabaseDataSafeStatusRegistered)
		if err != nil {
			return err
		}
		return tfresource.ReadResource(sync)
	}

	if configOpenMode == oci_database.UpdateAutonomousDatabaseDetailsOpenModeOnly || configPermissionLevel == oci_database.UpdateAutonomousDatabaseDetailsPermissionLevelRestricted {
		if configOpenMode == oci_database.UpdateAutonomousDatabaseDetailsOpenModeOnly {
			sync.D.Set("open_mode", configOpenMode)
		}
		if configPermissionLevel == oci_database.UpdateAutonomousDatabaseDetailsPermissionLevelRestricted {
			sync.D.Set("permission_level", configPermissionLevel)
		}
		err := sync.updateOpenModeAndPermission(sync.D.Id(), configOpenMode, configPermissionLevel)
		if err != nil {
			return err
		}
		return tfresource.ReadResource(sync)
	}

	if configDatabaseManagementStatus == oci_database.AutonomousDatabaseDatabaseManagementStatusEnabled {
		err := sync.updateAutonomousDatabaseManagementStatus(sync.D.Id(), oci_database.AutonomousDatabaseDatabaseManagementStatusEnabled)
		if err != nil {
			return err
		}
		return tfresource.ReadResource(sync)
	}

	if configOperationsInsightsStatus == oci_database.AutonomousDatabaseOperationsInsightsStatusEnabled {
		err := sync.updateOperationsInsightsStatus(sync.D.Id(), oci_database.AutonomousDatabaseOperationsInsightsStatusEnabled)
		if err != nil {
			return err
		}
		return tfresource.ReadResource(sync)
	}

	return nil
}

func readDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	err := sync.validateSwitchoverDatabase()
	if err != nil {
		return err
	}

	stateActive, stateInactive := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_database.AutonomousDatabaseLifecycleStateAvailable == oci_database.AutonomousDatabaseLifecycleStateEnum(wantedState) {
			stateActive = true
			stateInactive = false
		} else if oci_database.AutonomousDatabaseLifecycleStateStopped == oci_database.AutonomousDatabaseLifecycleStateEnum(wantedState) {
			stateInactive = true
			stateActive = false
		}
	}

	if stateActive {
		if err := sync.StartAutonomousDatabase(oci_database.AutonomousDatabaseLifecycleStateAvailable); err != nil {
			return err
		}
		if err := sync.D.Set("state", oci_database.AutonomousDatabaseLifecycleStateAvailable); err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("rotate_key_trigger"); ok && sync.D.HasChange("rotate_key_trigger") {
		err := sync.RotateAutonomousDatabaseEncryptionKey()
		if err != nil {
			return err
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if stateInactive {
		if err := sync.StopAutonomousDatabase(oci_database.AutonomousDatabaseLifecycleStateStopped); err != nil {
			return err
		}
		if err := sync.D.Set("state", oci_database.AutonomousDatabaseLifecycleStateStopped); err != nil {
			return err
		}
	}

	return nil
}

func deleteDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseAutonomousDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousDatabase
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseAutonomousDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateProvisioning),
		string(oci_database.AutonomousDatabaseLifecycleStateStarting),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateTerminating),
		string(oci_database.AutonomousDatabaseLifecycleStateUnavailable),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateStarting),
		string(oci_database.AutonomousDatabaseLifecycleStateProvisioning),
		string(oci_database.AutonomousDatabaseLifecycleStateUnavailable),
		string(oci_database.AutonomousDatabaseLifecycleStateScaleInProgress),
		string(oci_database.AutonomousDatabaseLifecycleStateUpdating),
		string(oci_database.AutonomousDatabaseLifecycleStateMaintenanceInProgress),
		string(oci_database.AutonomousDatabaseLifecycleStateUpgrading),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) Create() error {
	request := oci_database.CreateAutonomousDatabaseRequest{}
	err := s.populateTopLevelPolymorphicCreateAutonomousDatabaseRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousDatabase(context.Background(), request)
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

	s.Res = &response.AutonomousDatabase
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabase
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	if dataSafeStatus, ok := s.D.GetOkExists("data_safe_status"); ok && s.D.HasChange("data_safe_status") {
		oldRaw, newRaw := s.D.GetChange("data_safe_status")
		if newRaw != "" && oldRaw != "" {
			configDataSafeStatus := oci_database.AutonomousDatabaseDataSafeStatusEnum(strings.ToUpper(dataSafeStatus.(string)))
			err := s.updateDataSafeStatus(s.D.Id(), configDataSafeStatus)
			if err != nil {
				return err
			}
		}
	}

	updateFlag := false
	var openModeConfig oci_database.UpdateAutonomousDatabaseDetailsOpenModeEnum
	var permissionLevelConfig oci_database.UpdateAutonomousDatabaseDetailsPermissionLevelEnum
	if openMode, ok := s.D.GetOkExists("open_mode"); ok && s.D.HasChange("open_mode") {
		updateFlag = true
		openModeConfig = oci_database.UpdateAutonomousDatabaseDetailsOpenModeEnum(openMode.(string))
	}
	if permissionLevel, ok := s.D.GetOkExists("permission_level"); ok && s.D.HasChange("permission_level") {
		updateFlag = true
		permissionLevelConfig = oci_database.UpdateAutonomousDatabaseDetailsPermissionLevelEnum(permissionLevel.(string))
	}
	if updateFlag == true {
		err := s.updateOpenModeAndPermission(s.D.Id(), openModeConfig, permissionLevelConfig)
		if err != nil {
			return err
		}
	}

	if operationsInsightsStatus, ok := s.D.GetOkExists("operations_insights_status"); ok && s.D.HasChange("operations_insights_status") {
		oldRaw, newRaw := s.D.GetChange("operations_insights_status")
		if newRaw != "" && oldRaw != "" {
			configOperationsInsightsStatus := oci_database.AutonomousDatabaseOperationsInsightsStatusEnum((operationsInsightsStatus.(string)))
			err := s.updateOperationsInsightsStatus(s.D.Id(), configOperationsInsightsStatus)
			if err != nil {
				return err
			}
		}
	}

	if databaseManagementStatus, ok := s.D.GetOkExists("database_management_status"); ok && s.D.HasChange("database_management_status") {
		_, newRaw := s.D.GetChange("database_management_status")
		if newRaw != "" {
			configDatabaseManagementStatus := oci_database.AutonomousDatabaseDatabaseManagementStatusEnum((databaseManagementStatus.(string)))
			err := s.updateAutonomousDatabaseManagementStatus(s.D.Id(), configDatabaseManagementStatus)
			if err != nil {
				return err
			}
		}
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok && s.D.HasChange("kms_key_id") {
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok && s.D.HasChange("vault_id") {
			oldRaw1, newRaw1 := s.D.GetChange("kms_key_id")
			oldRaw2, newRaw2 := s.D.GetChange("vault_id")
			if newRaw1 != "" && oldRaw1 != "" && newRaw2 != "" && oldRaw2 != "" {
				err := s.ConfigureAutonomousDatabaseVaultKey(s.D.Id(), kmsKeyId.(string), vaultId.(string))
				if err != nil {
					return err
				}
			}
		}
	}

	request := oci_database.UpdateAutonomousDatabaseRequest{}

	if adminPassword, ok := s.D.GetOkExists("admin_password"); ok && s.D.HasChange("admin_password") {
		tmp := adminPassword.(string)
		request.AdminPassword = &tmp
	}

	if arePrimaryWhitelistedIpsUsed, ok := s.D.GetOkExists("are_primary_whitelisted_ips_used"); ok {
		tmp := arePrimaryWhitelistedIpsUsed.(bool)
		request.ArePrimaryWhitelistedIpsUsed = &tmp
	}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok && s.D.HasChange("cpu_core_count") {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}

	if customerContacts, ok := s.D.GetOkExists("customer_contacts"); ok && s.D.HasChange("customer_contacts") {
		interfaces := customerContacts.([]interface{})
		tmp := make([]oci_database.CustomerContact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "customer_contacts", stateDataIndex)
			converted, err := s.mapToCustomerContact(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("customer_contacts") {
			request.CustomerContacts = tmp
		}
	}

	if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
		tmp := dataStorageSizeInGB.(int)
		request.DataStorageSizeInGBs = &tmp
	}

	if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok && s.D.HasChange("data_storage_size_in_tbs") {
		tmp := dataStorageSizeInTBs.(int)
		request.DataStorageSizeInTBs = &tmp
	}

	if dbVersion, ok := s.D.GetOkExists("db_version"); ok && s.D.HasChange("db_version") {
		err := s.updateDbVersion(dbVersion.(string))
		if err != nil {
			return err
		}
	}

	if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok && s.D.HasChange("db_workload") {
		request.DbWorkload = oci_database.UpdateAutonomousDatabaseDetailsDbWorkloadEnum(dbWorkload.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAccessControlEnabled, ok := s.D.GetOkExists("is_access_control_enabled"); ok {
		tmp := isAccessControlEnabled.(bool)
		request.IsAccessControlEnabled = &tmp
	}

	if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok && s.D.HasChange("is_auto_scaling_enabled") {
		tmp := isAutoScalingEnabled.(bool)
		request.IsAutoScalingEnabled = &tmp
	}

	if isDataGuardEnabled, ok := s.D.GetOkExists("is_data_guard_enabled"); ok && s.D.HasChange("is_data_guard_enabled") {
		tmp := isDataGuardEnabled.(bool)
		request.IsDataGuardEnabled = &tmp
	}

	if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok && s.D.HasChange("is_free_tier") {
		tmp := isFreeTier.(bool)
		request.IsFreeTier = &tmp
	}

	if isMtlsConnectionRequired, ok := s.D.GetOkExists("is_mtls_connection_required"); ok && s.D.HasChange("is_mtls_connection_required") {
		tmp := isMtlsConnectionRequired.(bool)
		request.IsMtlsConnectionRequired = &tmp
	}

	if isRefreshableClone, ok := s.D.GetOkExists("is_refreshable_clone"); ok && s.D.HasChange("is_refreshable_clone") {
		tmp := isRefreshableClone.(bool)
		request.IsRefreshableClone = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok && s.D.HasChange("license_model") {
		request.LicenseModel = oci_database.UpdateAutonomousDatabaseDetailsLicenseModelEnum(licenseModel.(string))
	}

	var updateNewtworkAccessFlag = false
	if _, ok := s.D.GetOkExists("nsg_ids"); ok && s.D.HasChange("nsg_ids") {
		updateNewtworkAccessFlag = true
	}
	if _, ok := s.D.GetOkExists("private_endpoint_label"); ok && s.D.HasChange("private_endpoint_label") {
		updateNewtworkAccessFlag = true
	}
	if updateNewtworkAccessFlag == true {
		var nsgIds, _ = s.D.GetOkExists("nsg_ids")
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") || updateNewtworkAccessFlag {
			err := s.updateNsgIds(tmp)
			if err != nil {
				return err
			}
		}
	}

	if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok {
		tmp := float32(ocpuCount.(float64))
		request.OcpuCount = &tmp
	}

	if refreshableMode, ok := s.D.GetOkExists("refreshable_mode"); ok && s.D.HasChange("refreshable_mode") {
		request.RefreshableMode = oci_database.UpdateAutonomousDatabaseDetailsRefreshableModeEnum(refreshableMode.(string))
	}

	if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok && s.D.HasChange("scheduled_operations") {
		interfaces := scheduledOperations.([]interface{})
		tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scheduled_operations", stateDataIndex)
			converted, err := s.mapToScheduledOperationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("scheduled_operations") {
			request.ScheduledOperations = tmp
		}
	}

	if standbyWhitelistedIps, ok := s.D.GetOkExists("standby_whitelisted_ips"); ok {
		interfaces := standbyWhitelistedIps.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("standby_whitelisted_ips") {
			request.StandbyWhitelistedIps = tmp
		}
	}

	if whitelistedIps, ok := s.D.GetOkExists("whitelisted_ips"); ok && s.D.HasChange("whitelisted_ips") {
		set := whitelistedIps.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 {
			request.WhitelistedIps = tmp
		} else if s.D.HasChange("whitelisted_ips") {
			request.WhitelistedIps = []string{""}
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousDatabase(context.Background(), request)
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

	s.Res = &response.AutonomousDatabase
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) Delete() error {
	request := oci_database.DeleteAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteAutonomousDatabase(context.Background(), request)
	return err
}

func (s *DatabaseAutonomousDatabaseResourceCrud) SetData() error {
	if s.Res.ArePrimaryWhitelistedIpsUsed != nil {
		s.D.Set("are_primary_whitelisted_ips_used", *s.Res.ArePrimaryWhitelistedIpsUsed)
	}

	if s.Res.ApexDetails != nil {
		s.D.Set("apex_details", []interface{}{AutonomousDatabaseApexToMap(s.Res.ApexDetails)})
	} else {
		s.D.Set("apex_details", nil)
	}

	if s.Res.AutonomousContainerDatabaseId != nil {
		s.D.Set("autonomous_container_database_id", *s.Res.AutonomousContainerDatabaseId)
	}

	s.D.Set("autonomous_maintenance_schedule_type", s.Res.AutonomousMaintenanceScheduleType)

	s.D.Set("available_upgrade_versions", s.Res.AvailableUpgradeVersions)

	if s.Res.BackupConfig != nil {
		s.D.Set("backup_config", []interface{}{AutonomousDatabaseBackupConfigToMap(s.Res.BackupConfig)})
	} else {
		s.D.Set("backup_config", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{AutonomousDatabaseConnectionStringsToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.ConnectionUrls != nil {
		s.D.Set("connection_urls", []interface{}{AutonomousDatabaseConnectionUrlsToMap(s.Res.ConnectionUrls)})
	} else {
		s.D.Set("connection_urls", nil)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	customerContacts := []interface{}{}
	for _, item := range s.Res.CustomerContacts {
		customerContacts = append(customerContacts, CustomerContactToMap(item))
	}
	s.D.Set("customer_contacts", customerContacts)

	s.D.Set("data_safe_status", s.Res.DataSafeStatus)

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	s.D.Set("database_management_status", s.Res.DatabaseManagementStatus)

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	s.D.Set("db_workload", s.Res.DbWorkload)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FailedDataRecoveryInSeconds != nil {
		s.D.Set("failed_data_recovery_in_seconds", *s.Res.FailedDataRecoveryInSeconds)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("infrastructure_type", s.Res.InfrastructureType)

	if s.Res.IsAccessControlEnabled != nil {
		s.D.Set("is_access_control_enabled", *s.Res.IsAccessControlEnabled)
	}

	if s.Res.IsAutoScalingEnabled != nil {
		s.D.Set("is_auto_scaling_enabled", *s.Res.IsAutoScalingEnabled)
	}

	if s.Res.IsDataGuardEnabled != nil {
		s.D.Set("is_data_guard_enabled", *s.Res.IsDataGuardEnabled)
	}

	if s.Res.IsDedicated != nil {
		s.D.Set("is_dedicated", *s.Res.IsDedicated)
	}

	if s.Res.IsFreeTier != nil {
		s.D.Set("is_free_tier", *s.Res.IsFreeTier)
	}

	if s.Res.IsMtlsConnectionRequired != nil {
		s.D.Set("is_mtls_connection_required", *s.Res.IsMtlsConnectionRequired)
	}

	if s.Res.IsPreview != nil {
		s.D.Set("is_preview", *s.Res.IsPreview)
	}

	if s.Res.IsReconnectCloneEnabled != nil {
		s.D.Set("is_reconnect_clone_enabled", *s.Res.IsReconnectCloneEnabled)
	}

	if s.Res.IsRefreshableClone != nil {
		s.D.Set("is_refreshable_clone", *s.Res.IsRefreshableClone)
	}

	keyHistoryEntry := []interface{}{}
	for _, item := range s.Res.KeyHistoryEntry {
		keyHistoryEntry = append(keyHistoryEntry, AutonomousDatabaseKeyHistoryEntryToMap(item))
	}
	s.D.Set("key_history_entry", keyHistoryEntry)

	if s.Res.KeyStoreId != nil {
		s.D.Set("key_store_id", *s.Res.KeyStoreId)
	}

	if s.Res.KeyStoreWalletName != nil {
		s.D.Set("key_store_wallet_name", *s.Res.KeyStoreWalletName)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.KmsKeyLifecycleDetails != nil {
		s.D.Set("kms_key_lifecycle_details", *s.Res.KmsKeyLifecycleDetails)
	}

	if s.Res.KmsKeyVersionId != nil {
		s.D.Set("kms_key_version_id", *s.Res.KmsKeyVersionId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(utils.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.OcpuCount != nil {
		s.D.Set("ocpu_count", *s.Res.OcpuCount)
	}

	s.D.Set("open_mode", s.Res.OpenMode)

	s.D.Set("operations_insights_status", s.Res.OperationsInsightsStatus)

	s.D.Set("permission_level", s.Res.PermissionLevel)

	if s.Res.PrivateEndpoint != nil {
		s.D.Set("private_endpoint", *s.Res.PrivateEndpoint)
	}

	if s.Res.PrivateEndpointIp != nil {
		s.D.Set("private_endpoint_ip", *s.Res.PrivateEndpointIp)
	}

	if s.Res.PrivateEndpointLabel != nil {
		s.D.Set("private_endpoint_label", *s.Res.PrivateEndpointLabel)
	}

	if s.Res.RefreshableMode != "" {
		s.D.Set("refreshable_mode", s.Res.RefreshableMode)
	}

	s.D.Set("refreshable_status", s.Res.RefreshableStatus)

	s.D.Set("role", s.Res.Role)

	scheduledOperations := []interface{}{}
	for _, item := range s.Res.ScheduledOperations {
		scheduledOperations = append(scheduledOperations, ScheduledOperationDetailsToMap(item))
	}
	s.D.Set("scheduled_operations", scheduledOperations)

	if s.Res.ServiceConsoleUrl != nil {
		s.D.Set("service_console_url", *s.Res.ServiceConsoleUrl)
	}

	if s.Res.SourceId != nil {
		s.D.Set("source_id", *s.Res.SourceId)
	}

	if s.Res.StandbyDb != nil {
		s.D.Set("standby_db", []interface{}{AutonomousDatabaseStandbySummaryToMap(s.Res.StandbyDb)})
	} else {
		s.D.Set("standby_db", nil)
	}

	s.D.Set("standby_whitelisted_ips", s.Res.StandbyWhitelistedIps)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDeletionOfFreeAutonomousDatabase != nil {
		s.D.Set("time_deletion_of_free_autonomous_database", s.Res.TimeDeletionOfFreeAutonomousDatabase.String())
	}

	if s.Res.TimeMaintenanceBegin != nil {
		s.D.Set("time_maintenance_begin", s.Res.TimeMaintenanceBegin.String())
	}

	if s.Res.TimeMaintenanceEnd != nil {
		s.D.Set("time_maintenance_end", s.Res.TimeMaintenanceEnd.String())
	}

	if s.Res.TimeOfLastFailover != nil {
		s.D.Set("time_of_last_failover", s.Res.TimeOfLastFailover.String())
	}

	if s.Res.TimeOfLastRefresh != nil {
		s.D.Set("time_of_last_refresh", s.Res.TimeOfLastRefresh.String())
	}

	if s.Res.TimeOfLastRefreshPoint != nil {
		s.D.Set("time_of_last_refresh_point", s.Res.TimeOfLastRefreshPoint.String())
	}

	if s.Res.TimeOfLastSwitchover != nil {
		s.D.Set("time_of_last_switchover", s.Res.TimeOfLastSwitchover.String())
	}

	if s.Res.TimeOfNextRefresh != nil {
		s.D.Set("time_of_next_refresh", s.Res.TimeOfNextRefresh.String())
	}

	if s.Res.TimeReclamationOfFreeAutonomousDatabase != nil {
		s.D.Set("time_reclamation_of_free_autonomous_database", s.Res.TimeReclamationOfFreeAutonomousDatabase.String())
	}

	if s.Res.TimeUntilReconnectCloneEnabled != nil {
		s.D.Set("time_until_reconnect_clone_enabled", s.Res.TimeUntilReconnectCloneEnabled.String())
	}

	if s.Res.UsedDataStorageSizeInTBs != nil {
		s.D.Set("used_data_storage_size_in_tbs", *s.Res.UsedDataStorageSizeInTBs)
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	whitelistedIps := []interface{}{}
	for _, item := range s.Res.WhitelistedIps {
		whitelistedIps = append(whitelistedIps, item)
	}
	s.D.Set("whitelisted_ips", schema.NewSet(utils.LiteralTypeHashCodeForSets, whitelistedIps))

	return nil
}

func AutonomousDatabaseApexToMap(obj *oci_database.AutonomousDatabaseApex) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApexVersion != nil {
		result["apex_version"] = string(*obj.ApexVersion)
	}

	if obj.OrdsVersion != nil {
		result["ords_version"] = string(*obj.OrdsVersion)
	}

	return result
}

func AutonomousDatabaseBackupConfigToMap(obj *oci_database.AutonomousDatabaseBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ManualBackupBucketName != nil {
		result["manual_backup_bucket_name"] = string(*obj.ManualBackupBucketName)
	}

	result["manual_backup_type"] = string(obj.ManualBackupType)

	return result
}

func AutonomousDatabaseConnectionStringsToMap(obj *oci_database.AutonomousDatabaseConnectionStrings) map[string]interface{} {
	result := map[string]interface{}{}

	result["all_connection_strings"] = obj.AllConnectionStrings

	if obj.Dedicated != nil {
		result["dedicated"] = string(*obj.Dedicated)
	}

	if obj.High != nil {
		result["high"] = string(*obj.High)
	}

	if obj.Low != nil {
		result["low"] = string(*obj.Low)
	}

	if obj.Medium != nil {
		result["medium"] = string(*obj.Medium)
	}

	profiles := []interface{}{}
	for _, item := range obj.Profiles {
		profiles = append(profiles, DatabaseConnectionStringProfileToMap(item))
	}
	result["profiles"] = profiles

	return result
}

func AutonomousDatabaseConnectionUrlsToMap(obj *oci_database.AutonomousDatabaseConnectionUrls) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApexUrl != nil {
		result["apex_url"] = string(*obj.ApexUrl)
	}

	if obj.GraphStudioUrl != nil {
		result["graph_studio_url"] = string(*obj.GraphStudioUrl)
	}

	if obj.MachineLearningUserManagementUrl != nil {
		result["machine_learning_user_management_url"] = string(*obj.MachineLearningUserManagementUrl)
	}

	if obj.SqlDevWebUrl != nil {
		result["sql_dev_web_url"] = string(*obj.SqlDevWebUrl)
	}

	return result
}

func AutonomousDatabaseKeyHistoryEntryToMap(obj oci_database.AutonomousDatabaseKeyHistoryEntry) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.KmsKeyVersionId != nil {
		result["kms_key_version_id"] = string(*obj.KmsKeyVersionId)
	}

	if obj.TimeActivated != nil {
		result["time_activated"] = obj.TimeActivated.String()
	}

	if obj.VaultId != nil {
		result["vault_id"] = string(*obj.VaultId)
	}

	return result
}
func AutonomousDatabaseStandbySummaryToMap(obj *oci_database.AutonomousDatabaseStandbySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LagTimeInSeconds != nil {
		result["lag_time_in_seconds"] = int(*obj.LagTimeInSeconds)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}

func (s *DatabaseAutonomousDatabaseResourceCrud) mapToCustomerContact(fieldKeyFormat string) (oci_database.CustomerContact, error) {
	result := oci_database.CustomerContact{}

	if email, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email")); ok {
		tmp := email.(string)
		result.Email = &tmp
	}

	return result, nil
}

func CustomerContactToMap(obj oci_database.CustomerContact) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Email != nil {
		result["email"] = string(*obj.Email)
	}

	return result
}

func DatabaseConnectionStringProfileToMap(obj oci_database.DatabaseConnectionStringProfile) map[string]interface{} {
	result := map[string]interface{}{}

	result["consumer_group"] = string(obj.ConsumerGroup)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["host_format"] = string(obj.HostFormat)

	result["protocol"] = string(obj.Protocol)

	result["session_mode"] = string(obj.SessionMode)

	result["syntax_format"] = string(obj.SyntaxFormat)

	result["tls_authentication"] = string(obj.TlsAuthentication)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DatabaseAutonomousDatabaseResourceCrud) adbMapToDayOfWeek(fieldKeyFormat string) (oci_database.DayOfWeek, error) {
	result := oci_database.DayOfWeek{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.DayOfWeekNameEnum(name.(string))
	}

	return result, nil
}

func AdbDayOfWeekToMap(obj *oci_database.DayOfWeek) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseAutonomousDatabaseResourceCrud) mapToScheduledOperationDetails(fieldKeyFormat string) (oci_database.ScheduledOperationDetails, error) {
	result := oci_database.ScheduledOperationDetails{}

	if dayOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "day_of_week")); ok {
		if tmpList := dayOfWeek.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "day_of_week"), 0)
			tmp, err := s.adbMapToDayOfWeek(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert day_of_week, encountered error: %v", err)
			}
			result.DayOfWeek = &tmp
		}
	}

	if scheduledStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scheduled_start_time")); ok {
		tmp := scheduledStartTime.(string)
		result.ScheduledStartTime = &tmp
	}

	if scheduledStopTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scheduled_stop_time")); ok {
		tmp := scheduledStopTime.(string)
		result.ScheduledStopTime = &tmp
	}

	return result, nil
}

func ScheduledOperationDetailsToMap(obj oci_database.ScheduledOperationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DayOfWeek != nil {
		result["day_of_week"] = []interface{}{AdbDayOfWeekToMap(obj.DayOfWeek)}
	}

	if obj.ScheduledStartTime != nil {
		result["scheduled_start_time"] = string(*obj.ScheduledStartTime)
	}

	if obj.ScheduledStopTime != nil {
		result["scheduled_stop_time"] = string(*obj.ScheduledStopTime)
	}

	return result
}

func (s *DatabaseAutonomousDatabaseResourceCrud) populateTopLevelPolymorphicCreateAutonomousDatabaseRequest(request *oci_database.CreateAutonomousDatabaseRequest) error {
	//discriminator
	sourceRaw, ok := s.D.GetOkExists("source")
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "NONE" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("BACKUP_FROM_ID"):
		details := oci_database.CreateAutonomousDatabaseFromBackupDetails{}
		if autonomousDatabaseBackupId, ok := s.D.GetOkExists("autonomous_database_backup_id"); ok {
			tmp := autonomousDatabaseBackupId.(string)
			details.AutonomousDatabaseBackupId = &tmp
		}
		if cloneType, ok := s.D.GetOkExists("clone_type"); ok {
			details.CloneType = oci_database.CreateAutonomousDatabaseFromBackupDetailsCloneTypeEnum(cloneType.(string))
		}
		if adminPassword, ok := s.D.GetOkExists("admin_password"); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
		}
		if arePrimaryWhitelistedIpsUsed, ok := s.D.GetOkExists("are_primary_whitelisted_ips_used"); ok {
			tmp := arePrimaryWhitelistedIpsUsed.(bool)
			details.ArePrimaryWhitelistedIpsUsed = &tmp
		}
		if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
			tmp := autonomousContainerDatabaseId.(string)
			details.AutonomousContainerDatabaseId = &tmp
		}
		if autonomousMaintenanceScheduleType, ok := s.D.GetOkExists("autonomous_maintenance_schedule_type"); ok {
			details.AutonomousMaintenanceScheduleType = oci_database.CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum(autonomousMaintenanceScheduleType.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
		}
		if customerContacts, ok := s.D.GetOkExists("customer_contacts"); ok {
			interfaces := customerContacts.([]interface{})
			tmp := make([]oci_database.CustomerContact, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "customer_contacts", stateDataIndex)
				converted, err := s.mapToCustomerContact(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("customer_contacts") {
				details.CustomerContacts = tmp
			}
		}
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			details.DataStorageSizeInGBs = &tmp
		}
		if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
			tmp := dataStorageSizeInTBs.(int)
			details.DataStorageSizeInTBs = &tmp
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
			details.DbWorkload = oci_database.CreateAutonomousDatabaseBaseDbWorkloadEnum(dbWorkload.(string))
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
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isDataGuardEnabled, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			tmp := isDataGuardEnabled.(bool)
			details.IsDataGuardEnabled = &tmp
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if isMtlsConnectionRequired, ok := s.D.GetOkExists("is_mtls_connection_required"); ok {
			tmp := isMtlsConnectionRequired.(bool)
			details.IsMtlsConnectionRequired = &tmp
		}
		if isPreviewVersionWithServiceTermsAccepted, ok := s.D.GetOkExists("is_preview_version_with_service_terms_accepted"); ok {
			tmp := isPreviewVersionWithServiceTermsAccepted.(bool)
			details.IsPreviewVersionWithServiceTermsAccepted = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.CreateAutonomousDatabaseBaseLicenseModelEnum(licenseModel.(string))
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok {
			tmp := float32(ocpuCount.(float64))
			details.OcpuCount = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			interfaces := scheduledOperations.([]interface{})
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scheduled_operations", stateDataIndex)
				converted, err := s.mapToScheduledOperationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("scheduled_operations") {
				details.ScheduledOperations = tmp
			}
		}
		if standbyWhitelistedIps, ok := s.D.GetOkExists("standby_whitelisted_ips"); ok {
			interfaces := standbyWhitelistedIps.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("standby_whitelisted_ips") {
				details.StandbyWhitelistedIps = tmp
			}
		}

		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if whitelistedIps, ok := s.D.GetOkExists("whitelisted_ips"); ok {
			set := whitelistedIps.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("whitelisted_ips") {
				details.WhitelistedIps = tmp
			}
		}
		request.CreateAutonomousDatabaseDetails = details
	case strings.ToLower("BACKUP_FROM_TIMESTAMP"):
		details := oci_database.CreateAutonomousDatabaseFromBackupTimestampDetails{}
		if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
			tmp := autonomousDatabaseId.(string)
			details.AutonomousDatabaseId = &tmp
		}
		if cloneType, ok := s.D.GetOkExists("clone_type"); ok {
			details.CloneType = oci_database.CreateAutonomousDatabaseFromBackupTimestampDetailsCloneTypeEnum(cloneType.(string))
		}
		if timestamp, ok := s.D.GetOkExists("timestamp"); ok {
			tmp, err := time.Parse(time.RFC3339, timestamp.(string))
			if err != nil {
				return err
			}
			details.Timestamp = &oci_common.SDKTime{Time: tmp}
		}
		if adminPassword, ok := s.D.GetOkExists("admin_password"); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
		}
		if arePrimaryWhitelistedIpsUsed, ok := s.D.GetOkExists("are_primary_whitelisted_ips_used"); ok {
			tmp := arePrimaryWhitelistedIpsUsed.(bool)
			details.ArePrimaryWhitelistedIpsUsed = &tmp
		}
		if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
			tmp := autonomousContainerDatabaseId.(string)
			details.AutonomousContainerDatabaseId = &tmp
		}
		if autonomousMaintenanceScheduleType, ok := s.D.GetOkExists("autonomous_maintenance_schedule_type"); ok {
			details.AutonomousMaintenanceScheduleType = oci_database.CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum(autonomousMaintenanceScheduleType.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
		}
		if customerContacts, ok := s.D.GetOkExists("customer_contacts"); ok {
			interfaces := customerContacts.([]interface{})
			tmp := make([]oci_database.CustomerContact, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "customer_contacts", stateDataIndex)
				converted, err := s.mapToCustomerContact(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("customer_contacts") {
				details.CustomerContacts = tmp
			}
		}
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			details.DataStorageSizeInGBs = &tmp
		}
		if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
			tmp := dataStorageSizeInTBs.(int)
			details.DataStorageSizeInTBs = &tmp
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
			details.DbWorkload = oci_database.CreateAutonomousDatabaseBaseDbWorkloadEnum(dbWorkload.(string))
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
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isDataGuardEnabled, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			tmp := isDataGuardEnabled.(bool)
			details.IsDataGuardEnabled = &tmp
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if isMtlsConnectionRequired, ok := s.D.GetOkExists("is_mtls_connection_required"); ok {
			tmp := isMtlsConnectionRequired.(bool)
			details.IsMtlsConnectionRequired = &tmp
		}
		if isPreviewVersionWithServiceTermsAccepted, ok := s.D.GetOkExists("is_preview_version_with_service_terms_accepted"); ok {
			tmp := isPreviewVersionWithServiceTermsAccepted.(bool)
			details.IsPreviewVersionWithServiceTermsAccepted = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.CreateAutonomousDatabaseBaseLicenseModelEnum(licenseModel.(string))
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok {
			tmp := float32(ocpuCount.(float64))
			details.OcpuCount = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			interfaces := scheduledOperations.([]interface{})
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scheduled_operations", stateDataIndex)
				converted, err := s.mapToScheduledOperationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("scheduled_operations") {
				details.ScheduledOperations = tmp
			}
		}
		if standbyWhitelistedIps, ok := s.D.GetOkExists("standby_whitelisted_ips"); ok {
			interfaces := standbyWhitelistedIps.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("standby_whitelisted_ips") {
				details.StandbyWhitelistedIps = tmp
			}
		}

		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if whitelistedIps, ok := s.D.GetOkExists("whitelisted_ips"); ok {
			set := whitelistedIps.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("whitelisted_ips") {
				details.WhitelistedIps = tmp
			}
		}
		request.CreateAutonomousDatabaseDetails = details
	case strings.ToLower("CLONE_TO_REFRESHABLE"):
		details := oci_database.CreateRefreshableAutonomousDatabaseCloneDetails{}
		if refreshableMode, ok := s.D.GetOkExists("refreshable_mode"); ok {
			details.RefreshableMode = oci_database.CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum(refreshableMode.(string))
		}
		if sourceId, ok := s.D.GetOkExists("source_id"); ok {
			tmp := sourceId.(string)
			details.SourceId = &tmp
		}
		if adminPassword, ok := s.D.GetOkExists("admin_password"); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
		}
		if arePrimaryWhitelistedIpsUsed, ok := s.D.GetOkExists("are_primary_whitelisted_ips_used"); ok {
			tmp := arePrimaryWhitelistedIpsUsed.(bool)
			details.ArePrimaryWhitelistedIpsUsed = &tmp
		}
		if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
			tmp := autonomousContainerDatabaseId.(string)
			details.AutonomousContainerDatabaseId = &tmp
		}
		if autonomousMaintenanceScheduleType, ok := s.D.GetOkExists("autonomous_maintenance_schedule_type"); ok {
			details.AutonomousMaintenanceScheduleType = oci_database.CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum(autonomousMaintenanceScheduleType.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
		}
		if customerContacts, ok := s.D.GetOkExists("customer_contacts"); ok {
			interfaces := customerContacts.([]interface{})
			tmp := make([]oci_database.CustomerContact, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "customer_contacts", stateDataIndex)
				converted, err := s.mapToCustomerContact(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("customer_contacts") {
				details.CustomerContacts = tmp
			}
		}
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			details.DataStorageSizeInGBs = &tmp
		}
		if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
			tmp := dataStorageSizeInTBs.(int)
			details.DataStorageSizeInTBs = &tmp
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
			details.DbWorkload = oci_database.CreateAutonomousDatabaseBaseDbWorkloadEnum(dbWorkload.(string))
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
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isDataGuardEnabled, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			tmp := isDataGuardEnabled.(bool)
			details.IsDataGuardEnabled = &tmp
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if isMtlsConnectionRequired, ok := s.D.GetOkExists("is_mtls_connection_required"); ok {
			tmp := isMtlsConnectionRequired.(bool)
			details.IsMtlsConnectionRequired = &tmp
		}
		if isPreviewVersionWithServiceTermsAccepted, ok := s.D.GetOkExists("is_preview_version_with_service_terms_accepted"); ok {
			tmp := isPreviewVersionWithServiceTermsAccepted.(bool)
			details.IsPreviewVersionWithServiceTermsAccepted = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.CreateAutonomousDatabaseBaseLicenseModelEnum(licenseModel.(string))
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok {
			tmp := float32(ocpuCount.(float64))
			details.OcpuCount = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if refreshableMode, ok := s.D.GetOkExists("refreshable_mode"); ok {
			details.RefreshableMode = oci_database.CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum(refreshableMode.(string))
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			interfaces := scheduledOperations.([]interface{})
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scheduled_operations", stateDataIndex)
				converted, err := s.mapToScheduledOperationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("scheduled_operations") {
				details.ScheduledOperations = tmp
			}
		}
		if standbyWhitelistedIps, ok := s.D.GetOkExists("standby_whitelisted_ips"); ok {
			interfaces := standbyWhitelistedIps.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("standby_whitelisted_ips") {
				details.StandbyWhitelistedIps = tmp
			}
		}

		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if whitelistedIps, ok := s.D.GetOkExists("whitelisted_ips"); ok {
			set := whitelistedIps.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("whitelisted_ips") {
				details.WhitelistedIps = tmp
			}
		}
		request.CreateAutonomousDatabaseDetails = details
	case strings.ToLower("DATABASE"):
		details := oci_database.CreateAutonomousDatabaseCloneDetails{}
		if cloneType, ok := s.D.GetOkExists("clone_type"); ok {
			details.CloneType = oci_database.CreateAutonomousDatabaseCloneDetailsCloneTypeEnum(cloneType.(string))
		}
		if sourceId, ok := s.D.GetOkExists("source_id"); ok {
			tmp := sourceId.(string)
			details.SourceId = &tmp
		}
		if adminPassword, ok := s.D.GetOkExists("admin_password"); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
		}
		if arePrimaryWhitelistedIpsUsed, ok := s.D.GetOkExists("are_primary_whitelisted_ips_used"); ok {
			tmp := arePrimaryWhitelistedIpsUsed.(bool)
			details.ArePrimaryWhitelistedIpsUsed = &tmp
		}
		if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
			tmp := autonomousContainerDatabaseId.(string)
			details.AutonomousContainerDatabaseId = &tmp
		}
		if autonomousMaintenanceScheduleType, ok := s.D.GetOkExists("autonomous_maintenance_schedule_type"); ok {
			details.AutonomousMaintenanceScheduleType = oci_database.CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum(autonomousMaintenanceScheduleType.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
		}
		if customerContacts, ok := s.D.GetOkExists("customer_contacts"); ok {
			interfaces := customerContacts.([]interface{})
			tmp := make([]oci_database.CustomerContact, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "customer_contacts", stateDataIndex)
				converted, err := s.mapToCustomerContact(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("customer_contacts") {
				details.CustomerContacts = tmp
			}
		}
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			details.DataStorageSizeInGBs = &tmp
		}
		if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
			tmp := dataStorageSizeInTBs.(int)
			details.DataStorageSizeInTBs = &tmp
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
			details.DbWorkload = oci_database.CreateAutonomousDatabaseBaseDbWorkloadEnum(dbWorkload.(string))
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
		if isAccessControlEnabled, ok := s.D.GetOkExists("is_access_control_enabled"); ok {
			tmp := isAccessControlEnabled.(bool)
			details.IsAccessControlEnabled = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isDataGuardEnabled, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			tmp := isDataGuardEnabled.(bool)
			details.IsDataGuardEnabled = &tmp
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if isMtlsConnectionRequired, ok := s.D.GetOkExists("is_mtls_connection_required"); ok {
			tmp := isMtlsConnectionRequired.(bool)
			details.IsMtlsConnectionRequired = &tmp
		}
		if isPreviewVersionWithServiceTermsAccepted, ok := s.D.GetOkExists("is_preview_version_with_service_terms_accepted"); ok {
			tmp := isPreviewVersionWithServiceTermsAccepted.(bool)
			details.IsPreviewVersionWithServiceTermsAccepted = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.CreateAutonomousDatabaseBaseLicenseModelEnum(licenseModel.(string))
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok {
			tmp := float32(ocpuCount.(float64))
			details.OcpuCount = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			interfaces := scheduledOperations.([]interface{})
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scheduled_operations", stateDataIndex)
				converted, err := s.mapToScheduledOperationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("scheduled_operations") {
				details.ScheduledOperations = tmp
			}
		}
		if standbyWhitelistedIps, ok := s.D.GetOkExists("standby_whitelisted_ips"); ok {
			interfaces := standbyWhitelistedIps.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("standby_whitelisted_ips") {
				details.StandbyWhitelistedIps = tmp
			}
		}

		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if whitelistedIps, ok := s.D.GetOkExists("whitelisted_ips"); ok {
			set := whitelistedIps.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.WhitelistedIps = tmp
		}
		request.CreateAutonomousDatabaseDetails = details
	case strings.ToLower("NONE"):
		details := oci_database.CreateAutonomousDatabaseDetails{}
		if adminPassword, ok := s.D.GetOkExists("admin_password"); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
		}
		if arePrimaryWhitelistedIpsUsed, ok := s.D.GetOkExists("are_primary_whitelisted_ips_used"); ok {
			tmp := arePrimaryWhitelistedIpsUsed.(bool)
			details.ArePrimaryWhitelistedIpsUsed = &tmp
		}
		if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
			tmp := autonomousContainerDatabaseId.(string)
			details.AutonomousContainerDatabaseId = &tmp
		}
		if autonomousMaintenanceScheduleType, ok := s.D.GetOkExists("autonomous_maintenance_schedule_type"); ok {
			details.AutonomousMaintenanceScheduleType = oci_database.CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum(autonomousMaintenanceScheduleType.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
		}
		if customerContacts, ok := s.D.GetOkExists("customer_contacts"); ok {
			interfaces := customerContacts.([]interface{})
			tmp := make([]oci_database.CustomerContact, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "customer_contacts", stateDataIndex)
				converted, err := s.mapToCustomerContact(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("customer_contacts") {
				details.CustomerContacts = tmp
			}
		}
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			details.DataStorageSizeInGBs = &tmp
		}
		if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
			tmp := dataStorageSizeInTBs.(int)
			details.DataStorageSizeInTBs = &tmp
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
			details.DbWorkload = oci_database.CreateAutonomousDatabaseBaseDbWorkloadEnum(dbWorkload.(string))
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
		if isAccessControlEnabled, ok := s.D.GetOkExists("is_access_control_enabled"); ok {
			tmp := isAccessControlEnabled.(bool)
			details.IsAccessControlEnabled = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isDataGuardEnabled, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			tmp := isDataGuardEnabled.(bool)
			details.IsDataGuardEnabled = &tmp
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if isMtlsConnectionRequired, ok := s.D.GetOkExists("is_mtls_connection_required"); ok {
			tmp := isMtlsConnectionRequired.(bool)
			details.IsMtlsConnectionRequired = &tmp
		}
		if isPreviewVersionWithServiceTermsAccepted, ok := s.D.GetOkExists("is_preview_version_with_service_terms_accepted"); ok {
			tmp := isPreviewVersionWithServiceTermsAccepted.(bool)
			details.IsPreviewVersionWithServiceTermsAccepted = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.CreateAutonomousDatabaseBaseLicenseModelEnum(licenseModel.(string))
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok {
			tmp := float32(ocpuCount.(float64))
			details.OcpuCount = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			interfaces := scheduledOperations.([]interface{})
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scheduled_operations", stateDataIndex)
				converted, err := s.mapToScheduledOperationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("scheduled_operations") {
				details.ScheduledOperations = tmp
			}
		}
		if standbyWhitelistedIps, ok := s.D.GetOkExists("standby_whitelisted_ips"); ok {
			interfaces := standbyWhitelistedIps.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("standby_whitelisted_ips") {
				details.StandbyWhitelistedIps = tmp
			}
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if whitelistedIps, ok := s.D.GetOkExists("whitelisted_ips"); ok {
			set := whitelistedIps.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.WhitelistedIps = tmp
		}
		request.CreateAutonomousDatabaseDetails = details
	default:
		return fmt.Errorf("unknown source '%v' was specified", source)
	}
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeAutonomousDatabaseCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AutonomousDatabaseId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeAutonomousDatabaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) updateDataSafeStatus(autonomousDatabaseId string, dataSafeStatus oci_database.AutonomousDatabaseDataSafeStatusEnum) error {
	switch dataSafeStatus {
	case oci_database.AutonomousDatabaseDataSafeStatusRegistered:
		request := oci_database.RegisterAutonomousDatabaseDataSafeRequest{}
		if adminPassword, ok := s.D.GetOkExists("admin_password"); ok {
			tmp := adminPassword.(string)
			request.PdbAdminPassword = &tmp
		}
		request.AutonomousDatabaseId = &autonomousDatabaseId
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.RegisterAutonomousDatabaseDataSafe(context.Background(), request)

		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}

		return nil
	case oci_database.AutonomousDatabaseDataSafeStatusNotRegistered:
		request := oci_database.DeregisterAutonomousDatabaseDataSafeRequest{}
		if adminPassword, ok := s.D.GetOkExists("admin_password"); ok {
			tmp := adminPassword.(string)
			request.PdbAdminPassword = &tmp
		}
		request.AutonomousDatabaseId = &autonomousDatabaseId
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.DeregisterAutonomousDatabaseDataSafe(context.Background(), request)

		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("received unknown 'data_safe_status' %s", dataSafeStatus)
	}

}

func (s *DatabaseAutonomousDatabaseResourceCrud) updateDbVersion(dbVersion string) error {
	changeDbVersionRequest := oci_database.UpdateAutonomousDatabaseRequest{}
	changeDbVersionRequest.DbVersion = &dbVersion

	tmp := s.D.Id()
	changeDbVersionRequest.AutonomousDatabaseId = &tmp

	changeDbVersionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousDatabase(context.Background(), changeDbVersionRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) updateNsgIds(nsgIds []string) error {
	changeNsgIdsRequest := oci_database.UpdateAutonomousDatabaseRequest{}
	changeNsgIdsRequest.NsgIds = nsgIds

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok && s.D.HasChange("subnet_id") {
		tmp := subnetId.(string)
		if tmp == "null" {
			changeNsgIdsRequest.SubnetId = &(*new(string))
		} else {
			changeNsgIdsRequest.SubnetId = &tmp
		}
	}

	if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok && s.D.HasChange("private_endpoint_label") {
		tmp := privateEndpointLabel.(string)
		if tmp == "null" {
			changeNsgIdsRequest.PrivateEndpointLabel = &(*new(string))
		} else {
			changeNsgIdsRequest.PrivateEndpointLabel = &tmp
		}
	}
	tmp := s.D.Id()
	changeNsgIdsRequest.AutonomousDatabaseId = &tmp

	changeNsgIdsRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousDatabase(context.Background(), changeNsgIdsRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}
	s.Res = &response.AutonomousDatabase
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) validateSwitchoverDatabase() error {
	err := s.Get()
	if err != nil {
		return err
	}

	if switchoverTo, ok := s.D.GetOkExists("switchover_to"); ok && s.D.HasChange("switchover_to") {
		oldRaw, newRaw := s.D.GetChange("switchover_to")
		oldRaw = strings.ToUpper(oldRaw.(string))
		newRaw = strings.ToUpper(newRaw.(string))
		switchoverTo = strings.ToUpper(switchoverTo.(string))
		if isDataGuardEnabled, ok := s.D.GetOkExists("is_data_guard_enabled"); ok && isDataGuardEnabled.(bool) == true {
			if newRaw.(string) != "PRIMARY" || oldRaw.(string) != "" {
				if stateStatus, ok := s.D.GetOkExists("state"); ok {
					wantedStateStatus := oci_database.AutonomousDatabaseLifecycleStateEnum(strings.ToUpper(stateStatus.(string)))
					if wantedStateStatus == oci_database.AutonomousDatabaseLifecycleStateAvailable {
						if _, ok := s.D.GetOkExists("standby_db"); ok {
							fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "standby_db", 0)
							if standbyState, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "state")); ok {
								wantedStandByState := oci_database.AutonomousDatabaseStandbySummaryLifecycleStateEnum(strings.ToUpper(standbyState.(string)))
								if wantedStandByState == oci_database.AutonomousDatabaseStandbySummaryLifecycleStateAvailable {
									if err := s.switchoverDatabase(); err != nil {
										s.D.Set("switchover_to", oldRaw.(string))
										return err
									}

									s.D.Set("switchover_to", switchoverTo.(string))
								} else {
									s.D.Set("switchover_to", oldRaw.(string))

									return fmt.Errorf("Autonomous standby state: %s is not present in [AVAILABLE] states", wantedStandByState)
								}
							}
						}
					} else {
						s.D.Set("switchover_to", oldRaw.(string))

						return fmt.Errorf("Autonomous database state: %s is not present in [AVAILABLE] states", wantedStateStatus)
					}
				}
			} else {
				s.D.Set("switchover_to", "PRIMARY")
			}
		} else {
			if oldRaw.(string) == "" {
				s.D.Set("switchover_to", "PRIMARY")
			} else {
				s.D.Set("switchover_to", oldRaw.(string))
			}

			return fmt.Errorf("Autonomous Data Guard not found in enabled state")
		}
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) switchoverDatabase() error {
	request := oci_database.SwitchoverAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.SwitchoverAutonomousDatabase(context.Background(), request)
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

func (s *DatabaseAutonomousDatabaseResourceCrud) updateOperationsInsightsStatus(autonomousDatabaseId string, operationsInsightsStatus oci_database.AutonomousDatabaseOperationsInsightsStatusEnum) error {
	switch operationsInsightsStatus {
	case oci_database.AutonomousDatabaseOperationsInsightsStatusEnabled:
		request := oci_database.EnableAutonomousDatabaseOperationsInsightsRequest{}

		request.AutonomousDatabaseId = &autonomousDatabaseId
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
		response, err := s.Client.EnableAutonomousDatabaseOperationsInsights(context.Background(), request)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}

		return nil

	case oci_database.AutonomousDatabaseOperationsInsightsStatusNotEnabled:
		request := oci_database.DisableAutonomousDatabaseOperationsInsightsRequest{}

		request.AutonomousDatabaseId = &autonomousDatabaseId
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
		response, err := s.Client.DisableAutonomousDatabaseOperationsInsights(context.Background(), request)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("received unknown 'operations_insights_status' %s", operationsInsightsStatus)
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) updateAutonomousDatabaseManagementStatus(autonomousDatabaseId string, autonomousDatabaseManagement oci_database.AutonomousDatabaseDatabaseManagementStatusEnum) error {
	switch autonomousDatabaseManagement {
	case oci_database.AutonomousDatabaseDatabaseManagementStatusEnabled:
		request := oci_database.EnableAutonomousDatabaseManagementRequest{}

		request.AutonomousDatabaseId = &autonomousDatabaseId
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
		response, err := s.Client.EnableAutonomousDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}

		return nil

	case oci_database.AutonomousDatabaseDatabaseManagementStatusNotEnabled:
		request := oci_database.DisableAutonomousDatabaseManagementRequest{}
		request.AutonomousDatabaseId = &autonomousDatabaseId
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
		response, err := s.Client.DisableAutonomousDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("received unknown 'database_management_status' %s", autonomousDatabaseManagement)
	}
}

func inactiveAutonomousDatabaseIfNeeded(d *schema.ResourceData, sync *DatabaseAutonomousDatabaseResourceCrud) error {
	if err := sync.StopAutonomousDatabase(oci_database.AutonomousDatabaseLifecycleStateStopped); err != nil {
		return err
	}
	return tfresource.ReadResource(sync)
}

func (s *DatabaseAutonomousDatabaseResourceCrud) StartAutonomousDatabase(state oci_database.AutonomousDatabaseLifecycleStateEnum) error {
	request := oci_database.StartAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	if _, err := s.Client.StartAutonomousDatabase(context.Background(), request); err != nil {
		return err
	}
	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == state }

	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseAutonomousDatabaseResourceCrud) StopAutonomousDatabase(state oci_database.AutonomousDatabaseLifecycleStateEnum) error {
	request := oci_database.StopAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	if _, err := s.Client.StopAutonomousDatabase(context.Background(), request); err != nil {
		return err
	}
	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == state }

	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseAutonomousDatabaseResourceCrud) updateOpenModeAndPermission(autonomousDatabaseId string, openMode oci_database.UpdateAutonomousDatabaseDetailsOpenModeEnum, permissionLevel oci_database.UpdateAutonomousDatabaseDetailsPermissionLevelEnum) error {
	updateRequest := oci_database.UpdateAutonomousDatabaseRequest{}
	updateRequest.AutonomousDatabaseId = &autonomousDatabaseId

	if openMode, ok := s.D.GetOkExists("open_mode"); ok {
		oldVal, newVal := s.D.GetChange("open_mode")
		if oldVal == "" {
			newValFormatted := fmt.Sprintf("%v", oci_database.UpdateAutonomousDatabaseDetailsOpenModeOnly)
			if oldVal != newVal && newVal == newValFormatted {
				updateRequest.OpenMode = oci_database.UpdateAutonomousDatabaseDetailsOpenModeEnum(openMode.(string))
			}
		} else if s.D.HasChange("open_mode") {
			updateRequest.OpenMode = oci_database.UpdateAutonomousDatabaseDetailsOpenModeEnum(openMode.(string))
		}
	}
	if permissionLevel, ok := s.D.GetOkExists("permission_level"); ok {
		oldVal, newVal := s.D.GetChange("permission_level")
		if oldVal == "" {
			newValFormatted := fmt.Sprintf("%v", oci_database.UpdateAutonomousDatabaseDetailsPermissionLevelRestricted)
			if oldVal != newVal && newVal == newValFormatted {
				updateRequest.PermissionLevel = oci_database.UpdateAutonomousDatabaseDetailsPermissionLevelEnum(permissionLevel.(string))
			}
		} else if s.D.HasChange("permission_level") {
			updateRequest.PermissionLevel = oci_database.UpdateAutonomousDatabaseDetailsPermissionLevelEnum(permissionLevel.(string))
		}
	}
	updateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	updateResponse, err := s.Client.UpdateAutonomousDatabase(context.Background(), updateRequest)
	if err != nil {
		return err
	}

	workId := updateResponse.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) RotateAutonomousDatabaseEncryptionKey() error {
	request := oci_database.RotateAutonomousDatabaseEncryptionKeyRequest{}

	if isDedicated, ok := s.D.GetOkExists("is_dedicated"); !ok || isDedicated.(bool) == false {
		return fmt.Errorf("Autonomous database is not dedicated")
	}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RotateAutonomousDatabaseEncryptionKey(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}

	val := s.D.Get("rotate_key_trigger")
	s.D.Set("rotate_key_trigger", val)

	s.Res = &response.AutonomousDatabase
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) ConfigureAutonomousDatabaseVaultKey(autonomousDatabaseId string, kmsKeyId string, vautlId string) error {
	request := oci_database.ConfigureAutonomousDatabaseVaultKeyRequest{}

	request.AutonomousDatabaseId = &autonomousDatabaseId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	request.KmsKeyId = &kmsKeyId

	request.VaultId = &vautlId

	response, err := s.Client.ConfigureAutonomousDatabaseVaultKey(context.Background(), request)
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
