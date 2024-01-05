// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseAutonomousContainerDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		Create: createDatabaseAutonomousContainerDatabase,
		Read:   readDatabaseAutonomousContainerDatabase,
		Update: updateDatabaseAutonomousContainerDatabase,
		Delete: deleteDatabaseAutonomousContainerDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patch_model": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"autonomous_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"backup_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"backup_destination_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"internet_proxy": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vpc_password": {
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
									"vpc_user": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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
			"cloud_autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"db_version": {
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
			"fast_start_fail_over_lag_limit_in_seconds": {
				Type:     schema.TypeInt,
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
			"is_automatic_failover_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_dst_file_update_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"key_store_id": {
				Type:     schema.TypeString,
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
			"maintenance_window_details": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"custom_action_timeout_in_mins": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"days_of_week": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
						"hours_of_day": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 20,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"is_custom_action_timeout_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_monthly_patching_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"lead_time_in_weeks": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"months": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
						"patching_mode": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preference": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weeks_of_month": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 4,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},

						// Computed
					},
				},
			},
			"peer_autonomous_container_database_backup_config": {
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
						"backup_destination_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"internet_proxy": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"vpc_password": {
										Type:      schema.TypeString,
										Optional:  true,
										Computed:  true,
										ForceNew:  true,
										Sensitive: true,
									},
									"vpc_user": {
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
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"peer_autonomous_container_database_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"peer_autonomous_container_database_display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"peer_autonomous_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"peer_cloud_autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"peer_autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"peer_db_unique_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"protection_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"service_level_agreement_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"standby_maintenance_buffer_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rotate_key_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"vault_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"version_preference": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"available_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"compute_model": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dst_file_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"infrastructure_type": {
				Type:     schema.TypeString,
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
			"key_store_wallet_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"largest_provisionable_autonomous_database_in_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"last_maintenance_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_window": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"custom_action_timeout_in_mins": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"days_of_week": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"hours_of_day": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"is_custom_action_timeout_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_monthly_patching_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lead_time_in_weeks": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"months": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"patching_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"preference": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"weeks_of_month": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
					},
				},
			},
			"memory_per_oracle_compute_unit_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"next_maintenance_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provisionable_cpus": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
			"provisioned_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"reclaimable_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"reserved_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"role": {
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
			"time_of_last_backup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_snapshot_standby_revert": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_cpus": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createDatabaseAutonomousContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("rotate_key_trigger"); ok {
		err := sync.RotateContainerDatabaseEncryptionKey()
		if err != nil {
			return err
		}
	}

	return nil
}

func readDatabaseAutonomousContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func updateDatabaseAutonomousContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if _, ok := sync.D.GetOkExists("rotate_key_trigger"); ok && sync.D.HasChange("rotate_key_trigger") {
		err := sync.RotateContainerDatabaseEncryptionKey()
		if err != nil {
			return err
		}
	}

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseAutonomousContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseAutonomousContainerDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.AutonomousContainerDatabase
	DisableNotFoundRetries bool
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateProvisioning),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateBackupInProgress),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateRestoring),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminating),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateUnavailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateProvisioning),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateUpdating),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateRestarting),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateMaintenanceInProgress),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateUnavailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) Create() error {
	request := oci_database.CreateAutonomousContainerDatabaseRequest{}

	if autonomousExadataInfrastructureId, ok := s.D.GetOkExists("autonomous_exadata_infrastructure_id"); ok {
		tmp := autonomousExadataInfrastructureId.(string)
		request.AutonomousExadataInfrastructureId = &tmp
	}

	if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
		tmp := autonomousVmClusterId.(string)
		request.AutonomousVmClusterId = &tmp
	}

	if backupConfig, ok := s.D.GetOkExists("backup_config"); ok {
		if tmpList := backupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_config", 0)
			tmp, err := s.mapToAutonomousContainerDatabaseBackupConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupConfig = &tmp
		}
	}

	if cloudAutonomousVmClusterId, ok := s.D.GetOkExists("cloud_autonomous_vm_cluster_id"); ok {
		tmp := cloudAutonomousVmClusterId.(string)
		request.CloudAutonomousVmClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbName, ok := s.D.GetOkExists("db_name"); ok {
		tmp := dbName.(string)
		request.DbName = &tmp
	}

	if dbUniqueName, ok := s.D.GetOkExists("db_unique_name"); ok {
		tmp := dbUniqueName.(string)
		request.DbUniqueName = &tmp
	}

	if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
		tmp := dbVersion.(string)
		request.DbVersion = &tmp
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

	if fastStartFailOverLagLimitInSeconds, ok := s.D.GetOkExists("fast_start_fail_over_lag_limit_in_seconds"); ok && s.D.HasChange("fast_start_fail_over_lag_limit_in_seconds") {
		tmp := fastStartFailOverLagLimitInSeconds.(int)
		request.FastStartFailOverLagLimitInSeconds = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAutomaticFailoverEnabled, ok := s.D.GetOkExists("is_automatic_failover_enabled"); ok {
		tmp := isAutomaticFailoverEnabled.(bool)
		request.IsAutomaticFailoverEnabled = &tmp
	}

	if isDstFileUpdateEnabled, ok := s.D.GetOkExists("is_dst_file_update_enabled"); ok {
		tmp := isDstFileUpdateEnabled.(bool)
		request.IsDstFileUpdateEnabled = &tmp
	}

	if keyStoreId, ok := s.D.GetOkExists("key_store_id"); ok {
		tmp := keyStoreId.(string)
		request.KeyStoreId = &tmp
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if maintenanceWindowDetails, ok := s.D.GetOkExists("maintenance_window_details"); ok {
		if tmpList := maintenanceWindowDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window_details", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindowDetails = &tmp
		}
	}

	if patchModel, ok := s.D.GetOkExists("patch_model"); ok {
		request.PatchModel = oci_database.CreateAutonomousContainerDatabaseDetailsPatchModelEnum(patchModel.(string))
	}

	if peerAutonomousContainerDatabaseBackupConfig, ok := s.D.GetOkExists("peer_autonomous_container_database_backup_config"); ok {
		if tmpList := peerAutonomousContainerDatabaseBackupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "peer_autonomous_container_database_backup_config", 0)
			tmp, err := s.mapToPeerAutonomousContainerDatabaseBackupConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PeerAutonomousContainerDatabaseBackupConfig = &tmp
		}
	}

	if peerAutonomousContainerDatabaseCompartmentId, ok := s.D.GetOkExists("peer_autonomous_container_database_compartment_id"); ok {
		tmp := peerAutonomousContainerDatabaseCompartmentId.(string)
		request.PeerAutonomousContainerDatabaseCompartmentId = &tmp
	}

	if peerAutonomousContainerDatabaseDisplayName, ok := s.D.GetOkExists("peer_autonomous_container_database_display_name"); ok {
		tmp := peerAutonomousContainerDatabaseDisplayName.(string)
		request.PeerAutonomousContainerDatabaseDisplayName = &tmp
	}

	if peerAutonomousExadataInfrastructureId, ok := s.D.GetOkExists("peer_autonomous_exadata_infrastructure_id"); ok {
		tmp := peerAutonomousExadataInfrastructureId.(string)
		request.PeerAutonomousExadataInfrastructureId = &tmp
	}

	if peerCloudAutonomousVmClusterId, ok := s.D.GetOkExists("peer_cloud_autonomous_vm_cluster_id"); ok {
		tmp := peerCloudAutonomousVmClusterId.(string)
		request.PeerCloudAutonomousVmClusterId = &tmp
	}

	if peerAutonomousVmClusterId, ok := s.D.GetOkExists("peer_autonomous_vm_cluster_id"); ok {
		tmp := peerAutonomousVmClusterId.(string)
		request.PeerAutonomousVmClusterId = &tmp
	}

	if peerDbUniqueName, ok := s.D.GetOkExists("peer_db_unique_name"); ok {
		tmp := peerDbUniqueName.(string)
		request.PeerDbUniqueName = &tmp
	}

	if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
		request.ProtectionMode = oci_database.CreateAutonomousContainerDatabaseDetailsProtectionModeEnum(protectionMode.(string))
	}

	if serviceLevelAgreementType, ok := s.D.GetOkExists("service_level_agreement_type"); ok {
		request.ServiceLevelAgreementType = oci_database.CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum(serviceLevelAgreementType.(string))
	}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	if versionPreference, ok := s.D.GetOkExists("version_preference"); ok {
		request.VersionPreference = oci_database.CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum(versionPreference.(string))
	}

	if standbyMaintenanceBufferInDays, ok := s.D.GetOkExists("standby_maintenance_buffer_in_days"); ok {
		tmp := standbyMaintenanceBufferInDays.(int)
		request.StandbyMaintenanceBufferInDays = &tmp

	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.AutonomousContainerDatabase
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) Get() error {
	request := oci_database.GetAutonomousContainerDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousContainerDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousContainerDatabase
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateAutonomousContainerDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousContainerDatabaseId = &tmp

	if backupConfig, ok := s.D.GetOkExists("backup_config"); ok {
		if tmpList := backupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_config", 0)
			tmp, err := s.mapToAutonomousContainerDatabaseBackupConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupConfig = &tmp
		}
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isDstFileUpdateEnabled, ok := s.D.GetOkExists("is_dst_file_update_enabled"); ok {
		tmp := isDstFileUpdateEnabled.(bool)
		request.IsDstFileUpdateEnabled = &tmp
	}

	if maintenanceWindowDetails, ok := s.D.GetOkExists("maintenance_window_details"); ok {
		if tmpList := maintenanceWindowDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window_details", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindowDetails = &tmp
		}
	}

	if patchModel, ok := s.D.GetOkExists("patch_model"); ok {
		request.PatchModel = oci_database.UpdateAutonomousContainerDatabaseDetailsPatchModelEnum(patchModel.(string))
	}

	if standbyMaintenanceBufferInDays, ok := s.D.GetOkExists("standby_maintenance_buffer_in_days"); ok {
		tmp := standbyMaintenanceBufferInDays.(int)
		request.StandbyMaintenanceBufferInDays = &tmp
	}

	if versionPreference, ok := s.D.GetOkExists("version_preference"); ok {
		request.VersionPreference = oci_database.UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum(versionPreference.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousContainerDatabase(context.Background(), request)
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

	s.Res = &response.AutonomousContainerDatabase
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) Delete() error {
	request := oci_database.TerminateAutonomousContainerDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousContainerDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.TerminateAutonomousContainerDatabase(context.Background(), request)
	return err
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) SetData() error {
	if s.Res.AutonomousExadataInfrastructureId != nil {
		s.D.Set("autonomous_exadata_infrastructure_id", *s.Res.AutonomousExadataInfrastructureId)
	}

	if s.Res.AutonomousVmClusterId != nil {
		s.D.Set("autonomous_vm_cluster_id", *s.Res.AutonomousVmClusterId)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.AvailableCpus != nil {
		s.D.Set("available_cpus", *s.Res.AvailableCpus)
	}

	if s.Res.BackupConfig != nil {
		s.D.Set("backup_config", []interface{}{AutonomousContainerDatabaseBackupConfigToMap(s.Res.BackupConfig, s, false)})
	} else {
		s.D.Set("backup_config", nil)
	}

	if s.Res.CloudAutonomousVmClusterId != nil {
		s.D.Set("cloud_autonomous_vm_cluster_id", *s.Res.CloudAutonomousVmClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("compute_model", s.Res.ComputeModel)

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	if s.Res.DbUniqueName != nil {
		s.D.Set("db_unique_name", *s.Res.DbUniqueName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DstFileVersion != nil {
		s.D.Set("dst_file_version", *s.Res.DstFileVersion)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("infrastructure_type", s.Res.InfrastructureType)

	if s.Res.IsDstFileUpdateEnabled != nil {
		s.D.Set("is_dst_file_update_enabled", *s.Res.IsDstFileUpdateEnabled)
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

	if s.Res.KmsKeyVersionId != nil {
		s.D.Set("kms_key_version_id", *s.Res.KmsKeyVersionId)
	}

	if s.Res.LargestProvisionableAutonomousDatabaseInCpus != nil {
		s.D.Set("largest_provisionable_autonomous_database_in_cpus", *s.Res.LargestProvisionableAutonomousDatabaseInCpus)
	}

	if s.Res.LastMaintenanceRunId != nil {
		s.D.Set("last_maintenance_run_id", *s.Res.LastMaintenanceRunId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MemoryPerOracleComputeUnitInGBs != nil {
		s.D.Set("memory_per_oracle_compute_unit_in_gbs", *s.Res.MemoryPerOracleComputeUnitInGBs)
	}

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	if s.Res.PatchId != nil {
		s.D.Set("patch_id", *s.Res.PatchId)
	}

	s.D.Set("patch_model", s.Res.PatchModel)

	s.D.Set("provisionable_cpus", s.Res.ProvisionableCpus)

	if s.Res.ProvisionedCpus != nil {
		s.D.Set("provisioned_cpus", *s.Res.ProvisionedCpus)
	}

	if s.Res.ReclaimableCpus != nil {
		s.D.Set("reclaimable_cpus", *s.Res.ReclaimableCpus)
	}

	if s.Res.ReservedCpus != nil {
		s.D.Set("reserved_cpus", *s.Res.ReservedCpus)
	}

	s.D.Set("role", s.Res.Role)

	s.D.Set("service_level_agreement_type", s.Res.ServiceLevelAgreementType)

	if s.Res.StandbyMaintenanceBufferInDays != nil {
		s.D.Set("standby_maintenance_buffer_in_days", *s.Res.StandbyMaintenanceBufferInDays)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfLastBackup != nil {
		s.D.Set("time_of_last_backup", s.Res.TimeOfLastBackup.String())
	}

	if s.Res.TimeSnapshotStandbyRevert != nil {
		s.D.Set("time_snapshot_standby_revert", s.Res.TimeSnapshotStandbyRevert.String())
	}

	if s.Res.TotalCpus != nil {
		s.D.Set("total_cpus", *s.Res.TotalCpus)
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	s.D.Set("version_preference", s.Res.VersionPreference)

	return nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) mapToAutonomousContainerDatabaseBackupConfig(fieldKeyFormat string) (oci_database.AutonomousContainerDatabaseBackupConfig, error) {
	result := oci_database.AutonomousContainerDatabaseBackupConfig{}

	if backupDestinationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")); ok {
		interfaces := backupDestinationDetails.([]interface{})
		tmp := make([]oci_database.BackupDestinationDetails, len(interfaces))
		if len(interfaces) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backup_destination_details"), 0)
			converted, err := s.mapToBackupDestinationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[0] = converted
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

// service currently supports only one backupDestination
func AutonomousContainerDatabaseBackupConfigToMap(obj *oci_database.AutonomousContainerDatabaseBackupConfig, s *DatabaseAutonomousContainerDatabaseResourceCrud, dataSource bool) map[string]interface{} {
	result := map[string]interface{}{}

	backupDestinationDetails := []interface{}{}
	// s will be nil for datasource
	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_config.0.backup_destination_details", 0)
	if len(obj.BackupDestinationDetails) > 0 {
		backupDestinationDetails = append(backupDestinationDetails, AutonomousContainerDatabaseBackupDestinationDetailsToMap(obj.BackupDestinationDetails[0], s, dataSource, fieldKeyFormat))
		result["backup_destination_details"] = backupDestinationDetails
	}

	if obj.RecoveryWindowInDays != nil {
		result["recovery_window_in_days"] = int(*obj.RecoveryWindowInDays)
	}

	return result
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) mapToBackupDestinationDetails(fieldKeyFormat string) (oci_database.BackupDestinationDetails, error) {
	result := oci_database.BackupDestinationDetails{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if internetProxy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "internet_proxy")); ok {
		tmp := internetProxy.(string)
		result.InternetProxy = &tmp
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

func AutonomousContainerDatabaseBackupDestinationDetailsToMap(obj oci_database.BackupDestinationDetails, s *DatabaseAutonomousContainerDatabaseResourceCrud, dataSource bool, fieldKeyFormat string) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InternetProxy != nil {
		result["internet_proxy"] = string(*obj.InternetProxy)
	}

	result["type"] = string(obj.Type)

	if obj.VpcUser != nil {
		result["vpc_user"] = string(*obj.VpcUser)
	}

	if dataSource {
		return result
	}

	if vpcPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpc_password")); ok {
		tmp := vpcPassword.(string)
		result["vpc_password"] = &tmp
	}

	return result
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) mapToDayOfWeek(fieldKeyFormat string) (oci_database.DayOfWeek, error) {
	result := oci_database.DayOfWeek{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.DayOfWeekNameEnum(name.(string))
	}

	return result, nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) mapToMaintenanceWindow(fieldKeyFormat string) (oci_database.MaintenanceWindow, error) {
	result := oci_database.MaintenanceWindow{}

	if preference, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preference")); ok {
		result.Preference = oci_database.MaintenanceWindowPreferenceEnum(preference.(string))

		if result.Preference == oci_database.MaintenanceWindowPreferenceNoPreference {
			return result, nil
		}
	}

	if customActionTimeoutInMins, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_action_timeout_in_mins")); ok {
		tmp := customActionTimeoutInMins.(int)
		result.CustomActionTimeoutInMins = &tmp
	}

	if daysOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days_of_week")); ok {
		interfaces := daysOfWeek.([]interface{})
		tmp := make([]oci_database.DayOfWeek, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "days_of_week"), stateDataIndex)
			converted, err := s.mapToDayOfWeek(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "days_of_week")) {
			result.DaysOfWeek = tmp
		}
	}

	if hoursOfDay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hours_of_day")); ok {
		interfaces := hoursOfDay.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hours_of_day")) {
			result.HoursOfDay = tmp
		}
	}

	if isCustomActionTimeoutEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_custom_action_timeout_enabled")); ok {
		tmp := isCustomActionTimeoutEnabled.(bool)
		result.IsCustomActionTimeoutEnabled = &tmp
	}

	if isMonthlyPatchingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_monthly_patching_enabled")); ok {
		tmp := isMonthlyPatchingEnabled.(bool)
		result.IsMonthlyPatchingEnabled = &tmp
	}

	if leadTimeInWeeks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lead_time_in_weeks")); ok {
		tmp := leadTimeInWeeks.(int)
		if tmp > 0 {
			result.LeadTimeInWeeks = &tmp
		}
	}

	if months, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "months")); ok {
		interfaces := months.([]interface{})
		tmp := make([]oci_database.Month, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "months"), stateDataIndex)
			converted, err := s.mapToMonth(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "months")) {
			result.Months = tmp
		}
	}

	if patchingMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patching_mode")); ok {
		result.PatchingMode = oci_database.MaintenanceWindowPatchingModeEnum(patchingMode.(string))
	}

	if preference, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preference")); ok {
		result.Preference = oci_database.MaintenanceWindowPreferenceEnum(preference.(string))
	}

	if weeksOfMonth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weeks_of_month")); ok {
		interfaces := weeksOfMonth.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "weeks_of_month")) {
			result.WeeksOfMonth = tmp
		}
	}

	return result, nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) mapToMonth(fieldKeyFormat string) (oci_database.Month, error) {
	result := oci_database.Month{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.MonthNameEnum(name.(string))
	}

	return result, nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) mapToPeerAutonomousContainerDatabaseBackupConfig(fieldKeyFormat string) (oci_database.PeerAutonomousContainerDatabaseBackupConfig, error) {
	result := oci_database.PeerAutonomousContainerDatabaseBackupConfig{}

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

	return result, nil
}

func PeerAutonomousContainerDatabaseBackupConfigToMap(obj *oci_database.PeerAutonomousContainerDatabaseBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

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

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeAutonomousContainerDatabaseCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AutonomousContainerDatabaseId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeAutonomousContainerDatabaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) RotateContainerDatabaseEncryptionKey() error {
	request := oci_database.RotateAutonomousContainerDatabaseEncryptionKeyRequest{}

	if _, isDedicated := s.D.GetOkExists("cloud_autonomous_vm_cluster_id"); !isDedicated {
		return fmt.Errorf("Container database is not dedicated")
	}

	tmp := s.D.Id()
	request.AutonomousContainerDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RotateAutonomousContainerDatabaseEncryptionKey(context.Background(), request)
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

	val := s.D.Get("rotate_key_trigger")
	s.D.Set("rotate_key_trigger", val)

	s.Res = &response.AutonomousContainerDatabase
	return nil
}
