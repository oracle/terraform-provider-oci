// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
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
				//ForceNew: true, /* Expectation: This should be false. This is coming from Spec - needs to be checked.*/
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
			"auto_refresh_frequency_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_refresh_point_lag_in_seconds": {
				Type:     schema.TypeInt,
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
			},
			"backup_retention_period_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"byol_compute_count_limit": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"character_set": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clone_table_space_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"clone_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"compute_count": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"compute_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"database_edition": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"db_tools_details": {
				Type:     schema.TypeSet,
				Set:      dbToolsForSets,
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
						"compute_count": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"max_idle_time_in_minutes": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"disaster_recovery_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"arn_role": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"autonomous_database_provider": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AWS",
								"AZURE",
								"OCI",
								"OKV",
								"ORACLE_MANAGED",
							}, true),
						},
						"certificate_directory_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"certificate_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"directory_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_arn": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"kms_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"okv_kms_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"okv_uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_endpoint_uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vault_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vault_uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"in_memory_percentage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"is_auto_scaling_for_storage_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_backup_retention_locked": {
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
			"is_dev_tier": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_free_tier": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_local_data_guard_enabled": {
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
			"max_cpu_core_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ncharacter_set": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ocpu_count": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"private_endpoint_ip": {
				Type:     schema.TypeString,
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
					"",
				}, false),
			},
			"remote_disaster_recovery_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_replicate_automatic_backups": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"resource_pool_leader_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_pool_summary": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"pool_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"scheduled_operations": {
				Type:     schema.TypeSet,
				Set:      scheduledOperationsForSets,
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
						},
						"scheduled_stop_time": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"secret_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secret_version_number": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"security_attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
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
					"CROSS_REGION_DATAGUARD",
					"CROSS_REGION_DISASTER_RECOVERY",
					"CROSS_TENANCY_DISASTER_RECOVERY",
					"DATABASE",
					"NONE",
					"UNDELETE_ADB",
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
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time_of_auto_refresh_start": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"timestamp": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"use_latest_available_backup_time_stamp": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vault_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"whitelisted_ips": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
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
			"switchover_to_remote_peer_id": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"rotate_key_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_shrink_only": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"shrink_adb_trigger"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("is_shrink_only", "shrink_adb_trigger"),
			},
			"shrink_adb_trigger": {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"is_shrink_only"},
			},
			"is_disconnect_peer": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// Computed
			"actual_used_data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"allocated_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"apex_details": {
				Type:     schema.TypeList,
				Computed: true,
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
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_version_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			"cluster_placement_group_id": {
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
									"is_regional": {
										Type:     schema.TypeBool,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"apex_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_transforms_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"graph_studio_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"machine_learning_notebook_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"machine_learning_user_management_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mongo_db_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ords_url": {
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
			"dataguard_region_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disaster_recovery_region_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"encryption_key_history_entry": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"encryption_key": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"arn_role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"autonomous_database_provider": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"certificate_directory_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"certificate_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"directory_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_arn": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"kms_key_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"okv_kms_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"okv_uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_endpoint_uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vault_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vault_uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_activated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"failed_data_recovery_in_seconds": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"in_memory_area_in_gbs": {
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
			"is_remote_data_guard_enabled": {
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
			"local_adg_auto_failover_max_data_loss_limit": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"local_disaster_recovery_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_standby_db": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lag_time_in_seconds": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"maintenance_target_component": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_data_guard_role_changed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_disaster_recovery_role_changed": {
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
					},
				},
			},
			"long_term_backup_schedule": {
				Type:             schema.TypeList,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: longTermBackupSupressDiff,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_disabled": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},
						"repeat_cadence": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"retention_period_in_days": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},
						"time_of_backup": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
					},
				},
			},
			"maintenance_target_component": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_per_oracle_compute_unit_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"net_services_architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"next_long_term_backup_time_stamp": {
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
			"peer_db_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"peer_db_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"provisionable_cpus": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
			"public_connection_urls": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"apex_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_transforms_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"graph_studio_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"machine_learning_notebook_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"machine_learning_user_management_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mongo_db_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ords_url": {
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
			"public_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"refreshable_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_disaster_recovery_configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"disaster_recovery_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_replicate_automatic_backups": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_snapshot_standby": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"time_snapshot_standby_enabled_till": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lag_time_in_seconds": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"maintenance_target_component": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_data_guard_role_changed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_disaster_recovery_role_changed": {
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
			"supported_regions_to_clone_to": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"time_data_guard_role_changed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_deletion_of_free_autonomous_database": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_disaster_recovery_role_changed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_local_data_guard_enabled": {
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
			"time_of_joining_resource_pool": {
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
			"time_undeleted": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_until_reconnect_clone_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_backup_storage_size_in_gbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"used_data_storage_size_in_gbs": {
				Type:     schema.TypeInt,
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
		configOperationsInsightsStatus = oci_database.AutonomousDatabaseOperationsInsightsStatusEnum(operationsInsightsStatus.(string))
	}

	configDatabaseManagementStatus := oci_database.AutonomousDatabaseDatabaseManagementStatusNotEnabled
	if databaseManagementStatus, ok := sync.D.GetOkExists("database_management_status"); ok {
		configDatabaseManagementStatus = oci_database.AutonomousDatabaseDatabaseManagementStatusEnum(databaseManagementStatus.(string))
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

	if _, ok := sync.D.GetOkExists("is_shrink_only"); ok {
		raw := sync.D.Get("is_shrink_only")
		if raw.(bool) {
			err := sync.ShrinkAutonomousDatabase(oci_database.AutonomousDatabaseLifecycleStateAvailable)
			if err != nil {
				return err
			}
		}
	}

	if _, ok := sync.D.GetOkExists("shrink_adb_trigger"); ok {
		err := sync.ShrinkAutonomousDatabase(oci_database.AutonomousDatabaseLifecycleStateAvailable)
		if err != nil {
			return err
		}
	}

	if isInactiveRequest {
		return inactiveAutonomousDatabaseIfNeeded(d, sync)
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
	}

	if configDataSafeStatus == oci_database.AutonomousDatabaseDataSafeStatusRegistered {
		err := sync.updateDataSafeStatus(sync.D.Id(), oci_database.AutonomousDatabaseDataSafeStatusRegistered)
		if err != nil {
			return err
		}
		if e := tfresource.ReadResource(sync); e != nil {
			return e
		}
	}

	if configDatabaseManagementStatus == oci_database.AutonomousDatabaseDatabaseManagementStatusEnabled {
		err := sync.updateAutonomousDatabaseManagementStatus(sync.D.Id(), oci_database.AutonomousDatabaseDatabaseManagementStatusEnabled)
		if err != nil {
			return err
		}
		if e := tfresource.ReadResource(sync); e != nil {
			return e
		}
	}

	if configOperationsInsightsStatus == oci_database.AutonomousDatabaseOperationsInsightsStatusEnabled {
		err := sync.updateOperationsInsightsStatus(sync.D.Id(), oci_database.AutonomousDatabaseOperationsInsightsStatusEnabled)
		if err != nil {
			return err
		}
		if e := tfresource.ReadResource(sync); e != nil {
			return e
		}
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

	if _, ok := sync.D.GetOkExists("is_shrink_only"); ok && sync.D.HasChange("is_shrink_only") {
		oldRaw, newRaw := sync.D.GetChange("is_shrink_only")
		if !oldRaw.(bool) && newRaw.(bool) {
			err = sync.ShrinkAutonomousDatabase(oci_database.AutonomousDatabaseLifecycleStateAvailable)
			if err != nil {
				return err
			}
		}
	}

	if _, ok := sync.D.GetOkExists("shrink_adb_trigger"); ok && sync.D.HasChange("shrink_adb_trigger") {
		oldRaw, newRaw := sync.D.GetChange("shrink_adb_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ShrinkAutonomousDatabase(oci_database.AutonomousDatabaseLifecycleStateAvailable)

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("shrink_adb_trigger", oldRaw)
			return fmt.Errorf("new value of shrink_adb_trigger should be greater than the old value")
		}
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

	if _, ok := sync.D.GetOkExists("key_version_id"); ok && sync.D.HasChange("key_version_id") {
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
		string(oci_database.AutonomousDatabaseLifecycleStateStandby),
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
		string(oci_database.AutonomousDatabaseLifecycleStateRestarting),
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
	localAdg, localAdgDataLossLimit := validateLocalAdgCreate(s)
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

	s.Res = &response.AutonomousDatabase
	log.Printf("local adg value is, %v, data loss limit is %v", localAdg, localAdgDataLossLimit)
	if localAdg == true {
		err := s.UpdateLocalAdg(localAdg, localAdgDataLossLimit)
		if err != nil {
			return fmt.Errorf("resource created but standby could not be enabled because %v", err)
		}
	}
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

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok && s.D.HasChange("subscription_id") {
		oldRaw, newRaw := s.D.GetChange("subscription_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateSubscription(subscriptionId.(string))
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
			configOperationsInsightsStatus := oci_database.AutonomousDatabaseOperationsInsightsStatusEnum(operationsInsightsStatus.(string))
			err := s.updateOperationsInsightsStatus(s.D.Id(), configOperationsInsightsStatus)
			if err != nil {
				return err
			}
		}
	}

	if databaseManagementStatus, ok := s.D.GetOkExists("database_management_status"); ok && s.D.HasChange("database_management_status") {
		_, newRaw := s.D.GetChange("database_management_status")
		if newRaw != "" {
			configDatabaseManagementStatus := oci_database.AutonomousDatabaseDatabaseManagementStatusEnum(databaseManagementStatus.(string))
			err := s.updateAutonomousDatabaseManagementStatus(s.D.Id(), configDatabaseManagementStatus)
			if err != nil {
				return err
			}
		}
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			oldRaw1, newRaw1 := s.D.GetChange("kms_key_id")
			oldRaw2, newRaw2 := s.D.GetChange("vault_id")
			if newRaw1 != "" && oldRaw1 != "" && newRaw2 != "" && oldRaw2 != "" && (s.D.HasChange("kms_key_id") || s.D.HasChange("vault_id")) {
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

	if autoRefreshFrequencyInSeconds, ok := s.D.GetOkExists("auto_refresh_frequency_in_seconds"); ok && s.D.HasChange("auto_refresh_frequency_in_seconds") {
		tmp := autoRefreshFrequencyInSeconds.(int)
		request.AutoRefreshFrequencyInSeconds = &tmp
	}

	if autoRefreshPointLagInSeconds, ok := s.D.GetOkExists("auto_refresh_point_lag_in_seconds"); ok && s.D.HasChange("auto_refresh_point_lag_in_seconds") {
		tmp := autoRefreshPointLagInSeconds.(int)
		request.AutoRefreshPointLagInSeconds = &tmp
	}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	if byolComputeCountLimit, ok := s.D.GetOkExists("byol_compute_count_limit"); ok && s.D.HasChange("byol_compute_count_limit") {
		tmp := float32(byolComputeCountLimit.(float64))
		request.ByolComputeCountLimit = &tmp
	}

	if autonomousMaintenanceScheduleType, ok := s.D.GetOkExists("autonomous_maintenance_schedule_type"); ok && s.D.HasChange("autonomous_maintenance_schedule_type") {
		request.AutonomousMaintenanceScheduleType = oci_database.UpdateAutonomousDatabaseDetailsAutonomousMaintenanceScheduleTypeEnum(autonomousMaintenanceScheduleType.(string))
	}

	if computeCount, ok := s.D.GetOkExists("compute_count"); ok && s.D.HasChange("compute_count") {
		tmp := float32(computeCount.(float64))
		request.ComputeCount = &tmp
	}

	if computeModel, ok := s.D.GetOkExists("compute_model"); ok && s.D.HasChange("compute_model") {
		request.ComputeModel = oci_database.UpdateAutonomousDatabaseDetailsComputeModelEnum(computeModel.(string))
	}

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

	if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok && s.D.HasChange("data_storage_size_in_gb") {
		tmp := dataStorageSizeInGB.(int)
		request.DataStorageSizeInGBs = &tmp
	}

	if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok && s.D.HasChange("data_storage_size_in_tbs") {
		tmp := dataStorageSizeInTBs.(int)
		request.DataStorageSizeInTBs = &tmp
	}

	if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok && s.D.HasChange("database_edition") {
		request.DatabaseEdition = oci_database.AutonomousDatabaseSummaryDatabaseEditionEnum(databaseEdition.(string))
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

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok && s.D.HasChange("display_name") {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if encryptionKey, ok := s.D.GetOkExists("encryption_key"); ok && s.D.HasChange("encryption_key") {
		if tmpList := encryptionKey.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "encryption_key", 0)
			tmp, err := s.mapToAutonomousDatabaseEncryptionKeyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.EncryptionKey = tmp
		}
	}

	if dbName, ok := s.D.GetOkExists("db_name"); ok && s.D.HasChange("db_name") {
		tmp := dbName.(string)
		request.DbName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if inMemoryPercentage, ok := s.D.GetOkExists("in_memory_percentage"); ok {
		tmp := inMemoryPercentage.(int)
		request.InMemoryPercentage = &tmp
	}

	if isAccessControlEnabled, ok := s.D.GetOkExists("is_access_control_enabled"); ok && s.D.HasChange("is_access_control_enabled") {
		tmp := isAccessControlEnabled.(bool)
		request.IsAccessControlEnabled = &tmp
	}

	if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok && s.D.HasChange("is_auto_scaling_enabled") {
		tmp := isAutoScalingEnabled.(bool)
		request.IsAutoScalingEnabled = &tmp
	}

	if isAutoScalingForStorageEnabled, ok := s.D.GetOkExists("is_auto_scaling_for_storage_enabled"); ok && s.D.HasChange("is_auto_scaling_for_storage_enabled") {
		tmp := isAutoScalingForStorageEnabled.(bool)
		request.IsAutoScalingForStorageEnabled = &tmp
	}

	if isBackupRetentionLocked, ok := s.D.GetOkExists("is_backup_retention_locked"); ok && s.D.HasChange("is_backup_retention_locked") {
		tmp := isBackupRetentionLocked.(bool)
		request.IsBackupRetentionLocked = &tmp
	}

	if isDataGuardEnabled, ok := s.D.GetOkExists("is_data_guard_enabled"); ok && s.D.HasChange("is_data_guard_enabled") {
		tmp := isDataGuardEnabled.(bool)
		request.IsDataGuardEnabled = &tmp
	}

	if isDevTier, ok := s.D.GetOkExists("is_dev_tier"); ok && s.D.HasChange("is_dev_tier") {
		tmp := isDevTier.(bool)
		request.IsDevTier = &tmp
	}

	if isDisconnectPeer, ok := s.D.GetOkExists("is_disconnect_peer"); ok && s.D.HasChange("is_disconnect_peer") {
		tmp := isDisconnectPeer.(bool)
		request.IsDisconnectPeer = &tmp
	}

	if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok && s.D.HasChange("is_free_tier") {
		tmp := isFreeTier.(bool)
		request.IsFreeTier = &tmp
	}

	if isLocalDataGuardEnabled, ok := s.D.GetOkExists("is_local_data_guard_enabled"); ok && s.D.HasChange("is_local_data_guard_enabled") {
		tmp := isLocalDataGuardEnabled.(bool)
		request.IsLocalDataGuardEnabled = &tmp
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

	if localAdgAutoFailoverMaxDataLossLimit, ok := s.D.GetOkExists("local_adg_auto_failover_max_data_loss_limit"); ok && s.D.HasChange("local_adg_auto_failover_max_data_loss_limit") {
		tmp := localAdgAutoFailoverMaxDataLossLimit.(int)
		request.LocalAdgAutoFailoverMaxDataLossLimit = &tmp
	}

	if longTermBackupSchedule, ok := s.D.GetOkExists("long_term_backup_schedule"); ok && s.D.HasChange("long_term_backup_schedule") {
		if tmpList := longTermBackupSchedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "long_term_backup_schedule", 0)
			tmp, err := s.mapToLongTermBackUpScheduleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LongTermBackupSchedule = &tmp
		}
	}

	//if maxCpuCoreCount, ok := s.D.GetOkExists("max_cpu_core_count"); ok && s.D.HasChange("max_cpu_core_count") {
	//	tmp := maxCpuCoreCount.(int)
	//	request.MaxCpuCoreCount = &tmp
	//}

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

	if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok && s.D.HasChange("ocpu_count") {
		tmp := float32(ocpuCount.(float64))
		request.OcpuCount = &tmp
	}

	if peerDbId, ok := s.D.GetOkExists("peer_db_id"); ok && s.D.HasChange("peer_db_id") {
		tmp := peerDbId.(string)
		request.PeerDbId = &tmp
	}

	if privateEndpointIp, ok := s.D.GetOkExists("private_endpoint_ip"); ok && s.D.HasChange("private_endpoint_ip") {
		tmp := privateEndpointIp.(string)
		request.PrivateEndpointIp = &tmp
	}

	if refreshableMode, ok := s.D.GetOkExists("refreshable_mode"); ok && s.D.HasChange("refreshable_mode") {
		request.RefreshableMode = oci_database.UpdateAutonomousDatabaseDetailsRefreshableModeEnum(refreshableMode.(string))
	}

	if resourcePoolLeaderId, ok := s.D.GetOk("resource_pool_leader_id"); ok && s.D.HasChange("resource_pool_leader_id") {
		tmp := resourcePoolLeaderId.(string)
		request.ResourcePoolLeaderId = &tmp
	}

	if resourcePoolSummary, ok := s.D.GetOkExists("resource_pool_summary"); ok && s.D.HasChange("resource_pool_summary") {
		if tmpList := resourcePoolSummary.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_pool_summary", 0)
			tmp, err := s.mapToResourcePoolSummary(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResourcePoolSummary = &tmp
		}
	}

	if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok && s.D.HasChange("scheduled_operations") {
		set := scheduledOperations.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := scheduledOperationsForSets(interfaces[i])
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

	if dbToolsDetails, ok := s.D.GetOkExists("db_tools_details"); ok && s.D.HasChange("db_tools_details") {
		set := dbToolsDetails.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database.DatabaseTool, len(interfaces))
		for i := range interfaces {
			stateDataIndex := dbToolsForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_tools_details", stateDataIndex)
			converted, err := s.mapToDatabaseTool(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}

		if len(tmp) != 0 || s.D.HasChange("db_tools_details") {
			request.DbToolsDetails = tmp
		}
		if _, ok := s.D.GetOkExists("freeform_tags"); ok && !s.D.HasChange("freeform_tags") {
			request.FreeformTags = nil
		}
		if _, ok := s.D.GetOkExists("defined_tags"); ok && !s.D.HasChange("defined_tags") {
			request.DefinedTags = nil
		}
	}

	if secretId, ok := s.D.GetOkExists("secret_id"); ok && s.D.HasChange("secret_version_number") {
		tmp := secretId.(string)
		request.SecretId = &tmp
		if _, ok := s.D.GetOkExists("freeform_tags"); ok && !s.D.HasChange("freeform_tags") {
			request.FreeformTags = nil
		}
		if _, ok := s.D.GetOkExists("defined_tags"); ok && !s.D.HasChange("defined_tags") {
			request.DefinedTags = nil
		}
	}

	if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok && s.D.HasChange("secret_version_number") {
		tmp := secretVersionNumber.(int)
		request.SecretVersionNumber = &tmp
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok && s.D.HasChange("security_attributes") {
		//request.SecurityAttributes = tfresource.ObjectMapToStringMap(securityAttributes.(map[string]interface{}))
		request.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
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

	if timeOfAutoRefreshStart, ok := s.D.GetOkExists("time_of_auto_refresh_start"); ok && s.D.HasChange("time_of_auto_refresh_start") {
		tmp, err := time.Parse(time.RFC3339, timeOfAutoRefreshStart.(string))
		if err != nil {
			return err
		}
		request.TimeOfAutoRefreshStart = &oci_common.SDKTime{Time: tmp}
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

	if backupRetentionPeriodInDays, ok := s.D.GetOkExists("backup_retention_period_in_days"); ok && s.D.HasChange("backup_retention_period_in_days") {
		tmp := backupRetentionPeriodInDays.(int)
		request.BackupRetentionPeriodInDays = &tmp
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

func longTermBackupSupressDiff(key string, old string, new string, d *schema.ResourceData) bool {
	if longTermBackupSchedule, ok := d.GetOkExists("long_term_backup_schedule"); ok {
		if tmpList := longTermBackupSchedule.([]interface{}); len(tmpList) > 0 {
			isDisabled, _ := d.GetOkExists(fmt.Sprintf("long_term_backup_schedule.%d.is_disabled", 0))
			if isDisabled.(bool) == true {
				return true
			}
			return false
		}
		return false
	}
	return false
}

func (s *DatabaseAutonomousDatabaseResourceCrud) SetData() error {
	if s.Res.ActualUsedDataStorageSizeInTBs != nil {
		s.D.Set("actual_used_data_storage_size_in_tbs", *s.Res.ActualUsedDataStorageSizeInTBs)
	}

	if s.Res.AllocatedStorageSizeInTBs != nil {
		s.D.Set("allocated_storage_size_in_tbs", *s.Res.AllocatedStorageSizeInTBs)
	}

	if s.Res.ArePrimaryWhitelistedIpsUsed != nil {
		s.D.Set("are_primary_whitelisted_ips_used", *s.Res.ArePrimaryWhitelistedIpsUsed)
	}

	if s.Res.ApexDetails != nil {
		s.D.Set("apex_details", []interface{}{AutonomousDatabaseApexToMap(s.Res.ApexDetails)})
	} else {
		s.D.Set("apex_details", nil)
	}

	if s.Res.AutoRefreshFrequencyInSeconds != nil {
		s.D.Set("auto_refresh_frequency_in_seconds", *s.Res.AutoRefreshFrequencyInSeconds)
	}

	if s.Res.AutoRefreshPointLagInSeconds != nil {
		s.D.Set("auto_refresh_point_lag_in_seconds", *s.Res.AutoRefreshPointLagInSeconds)
	}

	if s.Res.AutonomousContainerDatabaseId != nil {
		s.D.Set("autonomous_container_database_id", *s.Res.AutonomousContainerDatabaseId)
	}

	s.D.Set("autonomous_maintenance_schedule_type", s.Res.AutonomousMaintenanceScheduleType)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	s.D.Set("available_upgrade_versions", s.Res.AvailableUpgradeVersions)

	if s.Res.BackupConfig != nil {
		s.D.Set("backup_config", []interface{}{AutonomousDatabaseBackupConfigToMap(s.Res.BackupConfig)})
	} else {
		s.D.Set("backup_config", nil)
	}

	if s.Res.BackupRetentionPeriodInDays != nil {
		s.D.Set("backup_retention_period_in_days", *s.Res.BackupRetentionPeriodInDays)
	}

	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	s.D.Set("clone_table_space_list", s.Res.CloneTableSpaceList)

	if s.Res.ClusterPlacementGroupId != nil {
		s.D.Set("cluster_placement_group_id", *s.Res.ClusterPlacementGroupId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeCount != nil {
		s.D.Set("compute_count", *s.Res.ComputeCount)
	}

	s.D.Set("compute_model", s.Res.ComputeModel)

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

	s.D.Set("database_edition", s.Res.DatabaseEdition)

	s.D.Set("database_management_status", s.Res.DatabaseManagementStatus)

	s.D.Set("dataguard_region_type", s.Res.DataguardRegionType)

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	dbToolsDetails := []interface{}{}
	for _, item := range s.Res.DbToolsDetails {
		dbToolsDetails = append(dbToolsDetails, DatabaseToolToMap(item))
	}
	s.D.Set("db_tools_details", schema.NewSet(dbToolsForSets, dbToolsDetails))

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	s.D.Set("db_workload", s.Res.DbWorkload)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("disaster_recovery_region_type", s.Res.DisasterRecoveryRegionType)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EncryptionKey != nil {
		encryptionKeyArray := []interface{}{}
		if encryptionKeyMap := AutonomousDatabaseEncryptionKeyDetailsToMap(&s.Res.EncryptionKey); encryptionKeyMap != nil {
			encryptionKeyArray = append(encryptionKeyArray, encryptionKeyMap)
		}
		s.D.Set("encryption_key", encryptionKeyArray)
	} else {
		s.D.Set("encryption_key", nil)
	}

	encryptionKeyHistoryEntry := []interface{}{}
	for _, item := range s.Res.EncryptionKeyHistoryEntry {
		encryptionKeyHistoryEntry = append(encryptionKeyHistoryEntry, AutonomousDatabaseEncryptionKeyHistoryEntryToMap(item))
	}
	s.D.Set("encryption_key_history_entry", encryptionKeyHistoryEntry)

	if s.Res.FailedDataRecoveryInSeconds != nil {
		s.D.Set("failed_data_recovery_in_seconds", *s.Res.FailedDataRecoveryInSeconds)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InMemoryAreaInGBs != nil {
		s.D.Set("in_memory_area_in_gbs", *s.Res.InMemoryAreaInGBs)
	}

	if s.Res.InMemoryPercentage != nil {
		s.D.Set("in_memory_percentage", *s.Res.InMemoryPercentage)
	}

	if s.Res.InMemoryAreaInGBs != nil {
		s.D.Set("in_memory_area_in_gbs", *s.Res.InMemoryAreaInGBs)
	}

	if s.Res.InMemoryPercentage != nil {
		s.D.Set("in_memory_percentage", *s.Res.InMemoryPercentage)
	}

	s.D.Set("infrastructure_type", s.Res.InfrastructureType)

	if s.Res.IsAccessControlEnabled != nil {
		s.D.Set("is_access_control_enabled", *s.Res.IsAccessControlEnabled)
	}

	if s.Res.IsAutoScalingEnabled != nil {
		s.D.Set("is_auto_scaling_enabled", *s.Res.IsAutoScalingEnabled)
	}

	if s.Res.IsAutoScalingForStorageEnabled != nil {
		s.D.Set("is_auto_scaling_for_storage_enabled", *s.Res.IsAutoScalingForStorageEnabled)
	}

	if s.Res.IsBackupRetentionLocked != nil {
		s.D.Set("is_backup_retention_locked", *s.Res.IsBackupRetentionLocked)
	}

	if s.Res.IsDataGuardEnabled != nil {
		s.D.Set("is_data_guard_enabled", *s.Res.IsDataGuardEnabled)
	}

	if s.Res.IsDedicated != nil {
		s.D.Set("is_dedicated", *s.Res.IsDedicated)
	}

	if s.Res.IsDevTier != nil {
		s.D.Set("is_dev_tier", *s.Res.IsDevTier)
	}

	if s.Res.IsFreeTier != nil {
		s.D.Set("is_free_tier", *s.Res.IsFreeTier)
	}

	if s.Res.IsLocalDataGuardEnabled != nil {
		s.D.Set("is_local_data_guard_enabled", *s.Res.IsLocalDataGuardEnabled)
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

	if s.Res.IsRemoteDataGuardEnabled != nil {
		s.D.Set("is_remote_data_guard_enabled", *s.Res.IsRemoteDataGuardEnabled)
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

	if s.Res.ByolComputeCountLimit != nil {
		s.D.Set("byol_compute_count_limit", s.Res.ByolComputeCountLimit)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LocalAdgAutoFailoverMaxDataLossLimit != nil {
		s.D.Set("local_adg_auto_failover_max_data_loss_limit", *s.Res.LocalAdgAutoFailoverMaxDataLossLimit)
	}

	s.D.Set("local_disaster_recovery_type", s.Res.LocalDisasterRecoveryType)

	if s.Res.LocalStandbyDb != nil {
		s.D.Set("local_standby_db", []interface{}{AutonomousDatabaseStandbySummaryToMap(s.Res.LocalStandbyDb)})
	} else {
		s.D.Set("local_standby_db", nil)
	}

	if s.Res.LongTermBackupSchedule != nil {
		s.D.Set("long_term_backup_schedule", []interface{}{LongTermBackUpScheduleDetailsToMap(s.Res.LongTermBackupSchedule)})
	} else {
		s.D.Set("long_term_backup_schedule", nil)
	}

	if s.Res.MaintenanceTargetComponent != nil {
		s.D.Set("maintenance_target_component", *s.Res.MaintenanceTargetComponent)
	}
	//if s.Res.MaxCpuCoreCount != nil {
	//	s.D.Set("max_cpu_core_count", *s.Res.MaxCpuCoreCount)
	//}

	if s.Res.MemoryPerOracleComputeUnitInGBs != nil {
		s.D.Set("memory_per_oracle_compute_unit_in_gbs", *s.Res.MemoryPerOracleComputeUnitInGBs)
	}

	//if s.Res.MaxCpuCoreCount != nil {
	//	s.D.Set("max_cpu_core_count", *s.Res.MaxCpuCoreCount)
	//}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

	s.D.Set("net_services_architecture", s.Res.NetServicesArchitecture)

	if s.Res.NextLongTermBackupTimeStamp != nil {
		s.D.Set("next_long_term_backup_time_stamp", s.Res.NextLongTermBackupTimeStamp.String())
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.OcpuCount != nil {
		s.D.Set("ocpu_count", *s.Res.OcpuCount)
	}

	s.D.Set("open_mode", s.Res.OpenMode)

	s.D.Set("operations_insights_status", s.Res.OperationsInsightsStatus)

	s.D.Set("peer_db_ids", s.Res.PeerDbIds)

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

	s.D.Set("provisionable_cpus", s.Res.ProvisionableCpus)

	if s.Res.PublicConnectionUrls != nil {
		s.D.Set("public_connection_urls", []interface{}{AutonomousDatabaseConnectionUrlsToMap(s.Res.PublicConnectionUrls)})
	} else {
		s.D.Set("public_connection_urls", nil)
	}

	if s.Res.PublicEndpoint != nil {
		s.D.Set("public_endpoint", *s.Res.PublicEndpoint)
	}

	if s.Res.RefreshableMode != "" {
		s.D.Set("refreshable_mode", s.Res.RefreshableMode)
	}

	s.D.Set("refreshable_status", s.Res.RefreshableStatus)

	if s.Res.RemoteDisasterRecoveryConfiguration != nil {
		s.D.Set("remote_disaster_recovery_configuration", []interface{}{DisasterRecoveryConfigurationToMap(s.Res.RemoteDisasterRecoveryConfiguration)})
	} else {
		s.D.Set("remote_disaster_recovery_configuration", nil)
	}

	if s.Res.ResourcePoolLeaderId != nil {
		s.D.Set("resource_pool_leader_id", *s.Res.ResourcePoolLeaderId)
	}

	if s.Res.ResourcePoolSummary != nil {
		s.D.Set("resource_pool_summary", []interface{}{ResourcePoolSummaryToMap(s.Res.ResourcePoolSummary)})
	} else {
		s.D.Set("resource_pool_summary", nil)
	}

	s.D.Set("role", s.Res.Role)

	scheduledOperations := []interface{}{}
	for _, item := range s.Res.ScheduledOperations {
		scheduledOperations = append(scheduledOperations, ScheduledOperationDetailsToMap(item))
	}
	s.D.Set("scheduled_operations", schema.NewSet(scheduledOperationsForSets, scheduledOperations))

	s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))

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

	if s.Res.SubscriptionId != nil {
		s.D.Set("subscription_id", *s.Res.SubscriptionId)
	}

	s.D.Set("supported_regions_to_clone_to", s.Res.SupportedRegionsToCloneTo)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDataGuardRoleChanged != nil {
		s.D.Set("time_data_guard_role_changed", s.Res.TimeDataGuardRoleChanged.String())
	}

	if s.Res.TimeDeletionOfFreeAutonomousDatabase != nil {
		s.D.Set("time_deletion_of_free_autonomous_database", s.Res.TimeDeletionOfFreeAutonomousDatabase.String())
	}

	if s.Res.TimeDisasterRecoveryRoleChanged != nil {
		s.D.Set("time_disaster_recovery_role_changed", s.Res.TimeDisasterRecoveryRoleChanged.String())
	}

	if s.Res.TimeLocalDataGuardEnabled != nil {
		s.D.Set("time_local_data_guard_enabled", s.Res.TimeLocalDataGuardEnabled.String())
	}

	if s.Res.TimeMaintenanceBegin != nil {
		s.D.Set("time_maintenance_begin", s.Res.TimeMaintenanceBegin.String())
	}

	if s.Res.TimeMaintenanceEnd != nil {
		s.D.Set("time_maintenance_end", s.Res.TimeMaintenanceEnd.String())
	}

	if s.Res.TimeOfAutoRefreshStart != nil {
		s.D.Set("time_of_auto_refresh_start", s.Res.TimeOfAutoRefreshStart.Format(time.RFC3339Nano))
	}

	if s.Res.TimeOfJoiningResourcePool != nil {
		s.D.Set("time_of_joining_resource_pool", s.Res.TimeOfJoiningResourcePool.String())
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

	if s.Res.TimeUndeleted != nil {
		s.D.Set("time_undeleted", s.Res.TimeUndeleted.String())
	}

	if s.Res.TimeUntilReconnectCloneEnabled != nil {
		s.D.Set("time_until_reconnect_clone_enabled", s.Res.TimeUntilReconnectCloneEnabled.String())
	}

	if s.Res.TotalBackupStorageSizeInGBs != nil {
		s.D.Set("total_backup_storage_size_in_gbs", *s.Res.TotalBackupStorageSizeInGBs)
	}

	if s.Res.UsedDataStorageSizeInGBs != nil {
		s.D.Set("used_data_storage_size_in_gbs", *s.Res.UsedDataStorageSizeInGBs)
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
	s.D.Set("whitelisted_ips", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, whitelistedIps))

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

	if obj.DatabaseTransformsUrl != nil {
		result["database_transforms_url"] = string(*obj.DatabaseTransformsUrl)
	}

	if obj.GraphStudioUrl != nil {
		result["graph_studio_url"] = string(*obj.GraphStudioUrl)
	}

	if obj.MachineLearningNotebookUrl != nil {
		result["machine_learning_notebook_url"] = string(*obj.MachineLearningNotebookUrl)
	}

	if obj.MachineLearningUserManagementUrl != nil {
		result["machine_learning_user_management_url"] = string(*obj.MachineLearningUserManagementUrl)
	}

	if obj.MongoDbUrl != nil {
		result["mongo_db_url"] = string(*obj.MongoDbUrl)
	}

	if obj.OrdsUrl != nil {
		result["ords_url"] = string(*obj.OrdsUrl)
	}

	if obj.SqlDevWebUrl != nil {
		result["sql_dev_web_url"] = string(*obj.SqlDevWebUrl)
	}

	return result
}

func (s *DatabaseAutonomousDatabaseResourceCrud) mapToAutonomousDatabaseEncryptionKeyDetails(fieldKeyFormat string) (oci_database.AutonomousDatabaseEncryptionKeyDetails, error) {
	var baseObject oci_database.AutonomousDatabaseEncryptionKeyDetails
	//discriminator
	providerRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "autonomous_database_provider"))
	var provider string
	if ok {
		provider = providerRaw.(string)
	} else {
		provider = "" // default value
	}
	switch strings.ToLower(provider) {
	case strings.ToLower("AWS"):
		details := oci_database.AwsKeyDetails{}
		if arnRole, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "arn_role")); ok {
			tmp := arnRole.(string)
			details.ArnRole = &tmp
		}
		if externalId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "external_id")); ok {
			tmp := externalId.(string)
			details.ExternalId = &tmp
		}
		if keyArn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_arn")); ok {
			tmp := keyArn.(string)
			details.KeyArn = &tmp
		}
		if serviceEndpointUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_endpoint_uri")); ok {
			tmp := serviceEndpointUri.(string)
			details.ServiceEndpointUri = &tmp
		}
		baseObject = details
	case strings.ToLower("AZURE"):
		details := oci_database.AzureKeyDetails{}
		if keyName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_name")); ok {
			tmp := keyName.(string)
			details.KeyName = &tmp
		}
		if vaultUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_uri")); ok {
			tmp := vaultUri.(string)
			details.VaultUri = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI"):
		details := oci_database.OciKeyDetails{}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		baseObject = details
	case strings.ToLower("OKV"):
		details := oci_database.OkvKeyDetails{}
		if certificateDirectoryName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_directory_name")); ok {
			tmp := certificateDirectoryName.(string)
			details.CertificateDirectoryName = &tmp
		}
		if certificateId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_id")); ok {
			tmp := certificateId.(string)
			details.CertificateId = &tmp
		}
		if directoryName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "directory_name")); ok {
			tmp := directoryName.(string)
			details.DirectoryName = &tmp
		}
		if okvKmsKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "okv_kms_key")); ok {
			tmp := okvKmsKey.(string)
			details.OkvKmsKey = &tmp
		}
		if okvUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "okv_uri")); ok {
			tmp := okvUri.(string)
			details.OkvUri = &tmp
		}
		baseObject = details
	case strings.ToLower("ORACLE_MANAGED"):
		details := oci_database.OracleManagedKeyDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown provider '%v' was specified", provider)
	}
	return baseObject, nil
}

func AutonomousDatabaseEncryptionKeyDetailsToMap(obj *oci_database.AutonomousDatabaseEncryptionKeyDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database.AwsKeyDetails:
		result["autonomous_database_provider"] = "AWS"

		if v.ArnRole != nil {
			result["arn_role"] = string(*v.ArnRole)
		}

		if v.ExternalId != nil {
			result["external_id"] = string(*v.ExternalId)
		}

		if v.KeyArn != nil {
			result["key_arn"] = string(*v.KeyArn)
		}

		if v.ServiceEndpointUri != nil {
			result["service_endpoint_uri"] = string(*v.ServiceEndpointUri)
		}
	case oci_database.AzureKeyDetails:
		result["autonomous_database_provider"] = "AZURE"

		if v.KeyName != nil {
			result["key_name"] = string(*v.KeyName)
		}

		if v.VaultUri != nil {
			result["vault_uri"] = string(*v.VaultUri)
		}
	case oci_database.OciKeyDetails:
		result["autonomous_database_provider"] = "OCI"

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}
	case oci_database.OkvKeyDetails:
		result["autonomous_database_provider"] = "OKV"

		if v.CertificateDirectoryName != nil {
			result["certificate_directory_name"] = string(*v.CertificateDirectoryName)
		}

		if v.CertificateId != nil {
			result["certificate_id"] = string(*v.CertificateId)
		}

		if v.DirectoryName != nil {
			result["directory_name"] = string(*v.DirectoryName)
		}

		if v.OkvKmsKey != nil {
			result["okv_kms_key"] = string(*v.OkvKmsKey)
		}

		if v.OkvUri != nil {
			result["okv_uri"] = string(*v.OkvUri)
		}
	case oci_database.OracleManagedKeyDetails:
		result["autonomous_database_provider"] = "ORACLE_MANAGED"
	default:
		log.Printf("[WARN] Received 'provider' of unknown type %v", *obj)
		return nil
	}

	return result
}

func AutonomousDatabaseEncryptionKeyHistoryEntryToMap(obj oci_database.AutonomousDatabaseEncryptionKeyHistoryEntry) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EncryptionKey != nil {
		encryptionKeyArray := []interface{}{}
		if encryptionKeyMap := AutonomousDatabaseEncryptionKeyDetailsToMap(&obj.EncryptionKey); encryptionKeyMap != nil {
			encryptionKeyArray = append(encryptionKeyArray, encryptionKeyMap)
		}
		result["encryption_key"] = encryptionKeyArray
	}

	if obj.TimeActivated != nil {
		result["time_activated"] = obj.TimeActivated.String()
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

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.LagTimeInSeconds != nil {
		result["lag_time_in_seconds"] = int(*obj.LagTimeInSeconds)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MaintenanceTargetComponent != nil {
		result["maintenance_target_component"] = string(*obj.MaintenanceTargetComponent)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeDataGuardRoleChanged != nil {
		result["time_data_guard_role_changed"] = obj.TimeDataGuardRoleChanged.String()
	}

	if obj.TimeDisasterRecoveryRoleChanged != nil {
		result["time_disaster_recovery_role_changed"] = obj.TimeDisasterRecoveryRoleChanged.String()
	}

	if obj.TimeMaintenanceBegin != nil {
		result["time_maintenance_begin"] = obj.TimeMaintenanceBegin.String()
	}

	if obj.TimeMaintenanceEnd != nil {
		result["time_maintenance_end"] = obj.TimeMaintenanceEnd.String()
	}

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

	if obj.IsRegional != nil {
		result["is_regional"] = bool(*obj.IsRegional)
	}

	result["protocol"] = string(obj.Protocol)

	result["session_mode"] = string(obj.SessionMode)

	result["syntax_format"] = string(obj.SyntaxFormat)

	result["tls_authentication"] = string(obj.TlsAuthentication)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DatabaseAutonomousDatabaseResourceCrud) mapToDatabaseTool(fieldKeyFormat string) (oci_database.DatabaseTool, error) {
	result := oci_database.DatabaseTool{}

	if computeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_count")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "compute_count")) {
		tmp := float32(computeCount.(float64))
		result.ComputeCount = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if maxIdleTimeInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_idle_time_in_minutes")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "max_idle_time_in_minutes")) {
		tmp := maxIdleTimeInMinutes.(int)
		result.MaxIdleTimeInMinutes = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.DatabaseToolNameEnum(name.(string))
	}

	return result, nil
}

func DatabaseToolToMap(obj oci_database.DatabaseTool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComputeCount != nil {
		result["compute_count"] = float32(*obj.ComputeCount)
	}
	if obj.ComputeCount == nil {
		result["compute_count"] = nil
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.MaxIdleTimeInMinutes != nil {
		result["max_idle_time_in_minutes"] = int(*obj.MaxIdleTimeInMinutes)
	} else {
		result["max_idle_time_in_minutes"] = nil
	}

	result["name"] = string(obj.Name)

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

func DisasterRecoveryConfigurationToMap(obj *oci_database.DisasterRecoveryConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	result["disaster_recovery_type"] = string(obj.DisasterRecoveryType)

	if obj.IsReplicateAutomaticBackups != nil {
		result["is_replicate_automatic_backups"] = bool(*obj.IsReplicateAutomaticBackups)
	}

	if obj.IsSnapshotStandby != nil {
		result["is_snapshot_standby"] = bool(*obj.IsSnapshotStandby)
	}

	if obj.TimeSnapshotStandbyEnabledTill != nil {
		result["time_snapshot_standby_enabled_till"] = obj.TimeSnapshotStandbyEnabledTill.String()
	}

	return result
}

func (s *DatabaseAutonomousDatabaseResourceCrud) mapToLongTermBackUpScheduleDetails(fieldKeyFormat string) (oci_database.LongTermBackUpScheduleDetails, error) {
	result := oci_database.LongTermBackUpScheduleDetails{}

	if isDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_disabled")); ok {
		tmp := isDisabled.(bool)
		result.IsDisabled = &tmp
	}

	if repeatCadence, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repeat_cadence")); ok {
		result.RepeatCadence = oci_database.LongTermBackUpScheduleDetailsRepeatCadenceEnum(repeatCadence.(string))
	}

	if retentionPeriodInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_period_in_days")); ok {
		tmp := retentionPeriodInDays.(int)
		result.RetentionPeriodInDays = &tmp
	}

	if timeOfBackup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_of_backup")); ok {
		tmp, err := time.Parse(time.RFC3339, timeOfBackup.(string))
		if err != nil {
			return result, err
		}
		result.TimeOfBackup = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) mapToResourcePoolSummary(fieldKeyFormat string) (oci_database.ResourcePoolSummary, error) {
	result := oci_database.ResourcePoolSummary{}

	if isDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_disabled")); ok {
		tmp := isDisabled.(bool)
		result.IsDisabled = &tmp
	}

	if poolSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pool_size")); ok {
		tmp := poolSize.(int)
		result.PoolSize = &tmp
	}

	return result, nil
}

func ResourcePoolSummaryToMap(obj *oci_database.ResourcePoolSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDisabled != nil {
		result["is_disabled"] = bool(*obj.IsDisabled)
	}

	if obj.PoolSize != nil {
		result["pool_size"] = int(*obj.PoolSize)
	}

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
		if cloneTableSpaceList, ok := s.D.GetOkExists("clone_table_space_list"); ok {
			interfaces := cloneTableSpaceList.([]interface{})
			tmp := make([]int, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(int)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("clone_table_space_list") {
				details.CloneTableSpaceList = tmp
			}
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
		if backupRetentionPeriodInDays, ok := s.D.GetOkExists("backup_retention_period_in_days"); ok {
			tmp := backupRetentionPeriodInDays.(int)
			details.BackupRetentionPeriodInDays = &tmp
		}
		if byolComputeCountLimit, ok := s.D.GetOkExists("byol_compute_count_limit"); ok {
			tmp := float32(byolComputeCountLimit.(float64))
			details.ByolComputeCountLimit = &tmp
		}
		if characterSet, ok := s.D.GetOkExists("character_set"); ok {
			tmp := characterSet.(string)
			details.CharacterSet = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if computeCount, ok := s.D.GetOkExists("compute_count"); ok {
			tmp := float32(computeCount.(float64))
			details.ComputeCount = &tmp
		}
		if computeModel, ok := s.D.GetOkExists("compute_model"); ok {
			details.ComputeModel = oci_database.CreateAutonomousDatabaseBaseComputeModelEnum(computeModel.(string))
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
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok && s.D.HasChange("data_storage_size_in_gb") {
			tmp := dataStorageSizeInGB.(int)
			details.DataStorageSizeInGBs = &tmp
		}
		if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
			tmp := dataStorageSizeInTBs.(int)
			details.DataStorageSizeInTBs = &tmp
		}
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.AutonomousDatabaseSummaryDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbToolsDetails, ok := s.D.GetOkExists("db_tools_details"); ok {
			set := dbToolsDetails.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.DatabaseTool, len(interfaces))
			for i := range interfaces {
				stateDataIndex := dbToolsForSets(interfaces[i])
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_tools_details", stateDataIndex)
				converted, err := s.mapToDatabaseTool(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("db_tools_details") {
				details.DbToolsDetails = tmp
			}
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
		if encryptionKey, ok := s.D.GetOkExists("encryption_key"); ok {
			if tmpList := encryptionKey.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "encryption_key", 0)
				tmp, err := s.mapToAutonomousDatabaseEncryptionKeyDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.EncryptionKey = tmp
			}
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if inMemoryPercentage, ok := s.D.GetOkExists("in_memory_percentage"); ok {
			tmp := inMemoryPercentage.(int)
			details.InMemoryPercentage = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isAutoScalingForStorageEnabled, ok := s.D.GetOkExists("is_auto_scaling_for_storage_enabled"); ok {
			tmp := isAutoScalingForStorageEnabled.(bool)
			details.IsAutoScalingForStorageEnabled = &tmp
		}
		if isBackupRetentionLocked, ok := s.D.GetOkExists("is_backup_retention_locked"); ok {
			tmp := isBackupRetentionLocked.(bool)
			details.IsBackupRetentionLocked = &tmp
		}
		if _, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			details.IsDataGuardEnabled = nil
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isDevTier, ok := s.D.GetOkExists("is_dev_tier"); ok {
			tmp := isDevTier.(bool)
			details.IsDevTier = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if _, ok := s.D.GetOkExists("is_local_data_guard_enabled"); ok {
			details.IsLocalDataGuardEnabled = nil
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
		//if maxCpuCoreCount, ok := s.D.GetOkExists("max_cpu_core_count"); ok {
		//	tmp := maxCpuCoreCount.(int)
		//	details.MaxCpuCoreCount = &tmp
		//}
		if ncharacterSet, ok := s.D.GetOkExists("ncharacter_set"); ok {
			tmp := ncharacterSet.(string)
			details.NcharacterSet = &tmp
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
		if privateEndpointIp, ok := s.D.GetOkExists("private_endpoint_ip"); ok {
			tmp := privateEndpointIp.(string)
			details.PrivateEndpointIp = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}

		if resourcePoolLeaderId, ok := s.D.GetOkExists("resource_pool_leader_id"); ok {
			tmp := resourcePoolLeaderId.(string)
			details.ResourcePoolLeaderId = &tmp
		}

		if resourcePoolSummary, ok := s.D.GetOk("resource_pool_summary"); ok {
			t := fmt.Sprintf("%s rp create", resourcePoolSummary)
			_, _ = io.WriteString(os.Stdout, t)
			if tmpList := resourcePoolSummary.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_pool_summary", 0)
				tmp, err := s.mapToResourcePoolSummary(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ResourcePoolSummary = &tmp
			}
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			set := scheduledOperations.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := scheduledOperationsForSets(interfaces[i])
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
		if secretId, ok := s.D.GetOkExists("secret_id"); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok {
			tmp := secretVersionNumber.(int)
			details.SecretVersionNumber = &tmp
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
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
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
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
		if cloneTableSpaceList, ok := s.D.GetOkExists("clone_table_space_list"); ok {
			interfaces := cloneTableSpaceList.([]interface{})
			tmp := make([]int, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(int)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("clone_table_space_list") {
				details.CloneTableSpaceList = tmp
			}
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
		if useLatestAvailableBackupTimeStamp, ok := s.D.GetOkExists("use_latest_available_backup_time_stamp"); ok {
			tmp := useLatestAvailableBackupTimeStamp.(bool)
			details.UseLatestAvailableBackupTimeStamp = &tmp
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
		if backupRetentionPeriodInDays, ok := s.D.GetOkExists("backup_retention_period_in_days"); ok {
			tmp := backupRetentionPeriodInDays.(int)
			details.BackupRetentionPeriodInDays = &tmp
		}
		if byolComputeCountLimit, ok := s.D.GetOkExists("byol_compute_count_limit"); ok {
			tmp := float32(byolComputeCountLimit.(float64))
			details.ByolComputeCountLimit = &tmp
		}
		if characterSet, ok := s.D.GetOkExists("character_set"); ok {
			tmp := characterSet.(string)
			details.CharacterSet = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if computeCount, ok := s.D.GetOkExists("compute_count"); ok {
			tmp := float32(computeCount.(float64))
			details.ComputeCount = &tmp
		}
		if computeModel, ok := s.D.GetOkExists("compute_model"); ok {
			details.ComputeModel = oci_database.CreateAutonomousDatabaseBaseComputeModelEnum(computeModel.(string))
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
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.AutonomousDatabaseSummaryDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbToolsDetails, ok := s.D.GetOkExists("db_tools_details"); ok {
			set := dbToolsDetails.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.DatabaseTool, len(interfaces))
			for i := range interfaces {
				stateDataIndex := dbToolsForSets(interfaces[i])
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_tools_details", stateDataIndex)
				converted, err := s.mapToDatabaseTool(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("db_tools_details") {
				details.DbToolsDetails = tmp
			}
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
		if encryptionKey, ok := s.D.GetOkExists("encryption_key"); ok {
			if tmpList := encryptionKey.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "encryption_key", 0)
				tmp, err := s.mapToAutonomousDatabaseEncryptionKeyDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.EncryptionKey = tmp
			}
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if inMemoryPercentage, ok := s.D.GetOkExists("in_memory_percentage"); ok {
			tmp := inMemoryPercentage.(int)
			details.InMemoryPercentage = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isAutoScalingForStorageEnabled, ok := s.D.GetOkExists("is_auto_scaling_for_storage_enabled"); ok {
			tmp := isAutoScalingForStorageEnabled.(bool)
			details.IsAutoScalingForStorageEnabled = &tmp
		}
		if isBackupRetentionLocked, ok := s.D.GetOkExists("is_backup_retention_locked"); ok {
			tmp := isBackupRetentionLocked.(bool)
			details.IsBackupRetentionLocked = &tmp
		}
		if _, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			details.IsDataGuardEnabled = nil
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isDevTier, ok := s.D.GetOkExists("is_dev_tier"); ok {
			tmp := isDevTier.(bool)
			details.IsDevTier = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if _, ok := s.D.GetOkExists("is_local_data_guard_enabled"); ok {
			details.IsLocalDataGuardEnabled = nil
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
		//if maxCpuCoreCount, ok := s.D.GetOkExists("max_cpu_core_count"); ok {
		//	tmp := maxCpuCoreCount.(int)
		//	details.MaxCpuCoreCount = &tmp
		//}
		if ncharacterSet, ok := s.D.GetOkExists("ncharacter_set"); ok {
			tmp := ncharacterSet.(string)
			details.NcharacterSet = &tmp
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
		if privateEndpointIp, ok := s.D.GetOkExists("private_endpoint_ip"); ok {
			tmp := privateEndpointIp.(string)
			details.PrivateEndpointIp = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if resourcePoolLeaderId, ok := s.D.GetOkExists("resource_pool_leader_id"); ok {
			tmp := resourcePoolLeaderId.(string)
			details.ResourcePoolLeaderId = &tmp
		}
		if resourcePoolSummary, ok := s.D.GetOkExists("resource_pool_summary"); ok {
			t := fmt.Sprintf("%s rp create backup", resourcePoolSummary)
			_, _ = io.WriteString(os.Stdout, t)
			if tmpList := resourcePoolSummary.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_pool_summary", 0)
				tmp, err := s.mapToResourcePoolSummary(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ResourcePoolSummary = &tmp
			}
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			set := scheduledOperations.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := scheduledOperationsForSets(interfaces[i])
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
		if secretId, ok := s.D.GetOkExists("secret_id"); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok {
			tmp := secretVersionNumber.(int)
			details.SecretVersionNumber = &tmp
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
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
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
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
		if autoRefreshFrequencyInSeconds, ok := s.D.GetOkExists("auto_refresh_frequency_in_seconds"); ok {
			tmp := autoRefreshFrequencyInSeconds.(int)
			details.AutoRefreshFrequencyInSeconds = &tmp
		}
		if autoRefreshPointLagInSeconds, ok := s.D.GetOkExists("auto_refresh_point_lag_in_seconds"); ok {
			tmp := autoRefreshPointLagInSeconds.(int)
			details.AutoRefreshPointLagInSeconds = &tmp
		}
		if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
			tmp := autonomousContainerDatabaseId.(string)
			details.AutonomousContainerDatabaseId = &tmp
		}
		if autonomousMaintenanceScheduleType, ok := s.D.GetOkExists("autonomous_maintenance_schedule_type"); ok {
			details.AutonomousMaintenanceScheduleType = oci_database.CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum(autonomousMaintenanceScheduleType.(string))
		}
		if backupRetentionPeriodInDays, ok := s.D.GetOkExists("backup_retention_period_in_days"); ok {
			tmp := backupRetentionPeriodInDays.(int)
			details.BackupRetentionPeriodInDays = &tmp
		}
		if byolComputeCountLimit, ok := s.D.GetOkExists("byol_compute_count_limit"); ok {
			tmp := float32(byolComputeCountLimit.(float64))
			details.ByolComputeCountLimit = &tmp
		}
		if characterSet, ok := s.D.GetOkExists("character_set"); ok {
			tmp := characterSet.(string)
			details.CharacterSet = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if computeCount, ok := s.D.GetOkExists("compute_count"); ok {
			tmp := float32(computeCount.(float64))
			details.ComputeCount = &tmp
		}
		if computeModel, ok := s.D.GetOkExists("compute_model"); ok {
			details.ComputeModel = oci_database.CreateAutonomousDatabaseBaseComputeModelEnum(computeModel.(string))
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
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.AutonomousDatabaseSummaryDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbToolsDetails, ok := s.D.GetOkExists("db_tools_details"); ok {
			set := dbToolsDetails.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.DatabaseTool, len(interfaces))
			for i := range interfaces {
				stateDataIndex := dbToolsForSets(interfaces[i])
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_tools_details", stateDataIndex)
				converted, err := s.mapToDatabaseTool(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("db_tools_details") {
				details.DbToolsDetails = tmp
			}
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
		if encryptionKey, ok := s.D.GetOkExists("encryption_key"); ok {
			if tmpList := encryptionKey.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "encryption_key", 0)
				tmp, err := s.mapToAutonomousDatabaseEncryptionKeyDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.EncryptionKey = tmp
			}
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if inMemoryPercentage, ok := s.D.GetOkExists("in_memory_percentage"); ok {
			tmp := inMemoryPercentage.(int)
			details.InMemoryPercentage = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isAutoScalingForStorageEnabled, ok := s.D.GetOkExists("is_auto_scaling_for_storage_enabled"); ok {
			tmp := isAutoScalingForStorageEnabled.(bool)
			details.IsAutoScalingForStorageEnabled = &tmp
		}
		if isBackupRetentionLocked, ok := s.D.GetOkExists("is_backup_retention_locked"); ok {
			tmp := isBackupRetentionLocked.(bool)
			details.IsBackupRetentionLocked = &tmp
		}
		if _, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			details.IsDataGuardEnabled = nil
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isDevTier, ok := s.D.GetOkExists("is_dev_tier"); ok {
			tmp := isDevTier.(bool)
			details.IsDevTier = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if _, ok := s.D.GetOkExists("is_local_data_guard_enabled"); ok {
			details.IsLocalDataGuardEnabled = nil
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
		//if maxCpuCoreCount, ok := s.D.GetOkExists("max_cpu_core_count"); ok {
		//	tmp := maxCpuCoreCount.(int)
		//	details.MaxCpuCoreCount = &tmp
		//}
		if ncharacterSet, ok := s.D.GetOkExists("ncharacter_set"); ok {
			tmp := ncharacterSet.(string)
			details.NcharacterSet = &tmp
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
		if privateEndpointIp, ok := s.D.GetOkExists("private_endpoint_ip"); ok {
			tmp := privateEndpointIp.(string)
			details.PrivateEndpointIp = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if refreshableMode, ok := s.D.GetOkExists("refreshable_mode"); ok {
			details.RefreshableMode = oci_database.CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum(refreshableMode.(string))
		}
		if resourcePoolLeaderId, ok := s.D.GetOkExists("resource_pool_leader_id"); ok {
			tmp := resourcePoolLeaderId.(string)
			details.ResourcePoolLeaderId = &tmp
		}
		if resourcePoolSummary, ok := s.D.GetOkExists("resource_pool_summary"); ok {
			t := fmt.Sprintf("%s rp create clone to refreshable", resourcePoolSummary)
			_, _ = io.WriteString(os.Stdout, t)
			if tmpList := resourcePoolSummary.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_pool_summary", 0)
				tmp, err := s.mapToResourcePoolSummary(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ResourcePoolSummary = &tmp
			}
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			set := scheduledOperations.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := scheduledOperationsForSets(interfaces[i])
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
		if secretId, ok := s.D.GetOkExists("secret_id"); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok {
			tmp := secretVersionNumber.(int)
			details.SecretVersionNumber = &tmp
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
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
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if timeOfAutoRefreshStart, ok := s.D.GetOkExists("time_of_auto_refresh_start"); ok {
			tmp, err := time.Parse(time.RFC3339, timeOfAutoRefreshStart.(string))
			if err != nil {
				return err
			}
			details.TimeOfAutoRefreshStart = &oci_common.SDKTime{Time: tmp}
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
	case strings.ToLower("CROSS_REGION_DATAGUARD"):
		details := oci_database.CreateCrossRegionAutonomousDatabaseDataGuardDetails{}
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
		if backupRetentionPeriodInDays, ok := s.D.GetOkExists("backup_retention_period_in_days"); ok {
			tmp := backupRetentionPeriodInDays.(int)
			details.BackupRetentionPeriodInDays = &tmp
		}
		if byolComputeCountLimit, ok := s.D.GetOkExists("byol_compute_count_limit"); ok {
			tmp := float32(byolComputeCountLimit.(float64))
			details.ByolComputeCountLimit = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if computeCount, ok := s.D.GetOkExists("compute_count"); ok {
			tmp := float32(computeCount.(float64))
			details.ComputeCount = &tmp
		}
		if computeModel, ok := s.D.GetOkExists("compute_model"); ok {
			details.ComputeModel = oci_database.CreateAutonomousDatabaseBaseComputeModelEnum(computeModel.(string))
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
		if dbToolsDetails, ok := s.D.GetOkExists("db_tools_details"); ok {
			set := dbToolsDetails.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.DatabaseTool, len(interfaces))
			for i := range interfaces {
				stateDataIndex := dbToolsForSets(interfaces[i])
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_tools_details", stateDataIndex)
				converted, err := s.mapToDatabaseTool(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("db_tools_details") {
				details.DbToolsDetails = tmp
			}
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
		if encryptionKey, ok := s.D.GetOkExists("encryption_key"); ok {
			if tmpList := encryptionKey.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "encryption_key", 0)
				tmp, err := s.mapToAutonomousDatabaseEncryptionKeyDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.EncryptionKey = tmp
			}
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if inMemoryPercentage, ok := s.D.GetOkExists("in_memory_percentage"); ok {
			tmp := inMemoryPercentage.(int)
			details.InMemoryPercentage = &tmp
		}
		if isAccessControlEnabled, ok := s.D.GetOkExists("is_access_control_enabled"); ok {
			tmp := isAccessControlEnabled.(bool)
			details.IsAccessControlEnabled = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isBackupRetentionLocked, ok := s.D.GetOkExists("is_backup_retention_locked"); ok {
			tmp := isBackupRetentionLocked.(bool)
			details.IsBackupRetentionLocked = &tmp
		}
		if _, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			details.IsDataGuardEnabled = nil
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isDevTier, ok := s.D.GetOkExists("is_dev_tier"); ok {
			tmp := isDevTier.(bool)
			details.IsDevTier = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if _, ok := s.D.GetOkExists("is_local_data_guard_enabled"); ok {
			details.IsLocalDataGuardEnabled = nil
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
			tmp := ocpuCount.(float32)
			details.OcpuCount = &tmp
		}
		if privateEndpointIp, ok := s.D.GetOkExists("private_endpoint_ip"); ok {
			tmp := privateEndpointIp.(string)
			details.PrivateEndpointIp = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}

		if resourcePoolLeaderId, ok := s.D.GetOkExists("resource_pool_leader_id"); ok {
			tmp := resourcePoolLeaderId.(string)
			details.ResourcePoolLeaderId = &tmp
		}
		if resourcePoolSummary, ok := s.D.GetOkExists("resource_pool_summary"); ok {
			t := fmt.Sprintf("%s rp create cross region DG", resourcePoolSummary)
			_, _ = io.WriteString(os.Stdout, t)
			if tmpList := resourcePoolSummary.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_pool_summary", 0)
				tmp, err := s.mapToResourcePoolSummary(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ResourcePoolSummary = &tmp
			}
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			set := scheduledOperations.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := scheduledOperationsForSets(interfaces[i])
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
		if secretId, ok := s.D.GetOkExists("secret_id"); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok {
			tmp := secretVersionNumber.(int)
			details.SecretVersionNumber = &tmp
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
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
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
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
	case strings.ToLower("CROSS_REGION_DISASTER_RECOVERY"):
		details := oci_database.CreateCrossRegionDisasterRecoveryDetails{}
		if isReplicateAutomaticBackups, ok := s.D.GetOkExists("is_replicate_automatic_backups"); ok {
			tmp := isReplicateAutomaticBackups.(bool)
			details.IsReplicateAutomaticBackups = &tmp
		}
		if remoteDisasterRecoveryType, ok := s.D.GetOkExists("remote_disaster_recovery_type"); ok {
			details.RemoteDisasterRecoveryType = oci_database.DisasterRecoveryConfigurationDisasterRecoveryTypeEnum(remoteDisasterRecoveryType.(string))
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
		if byolComputeCountLimit, ok := s.D.GetOkExists("byol_compute_count_limit"); ok {
			tmp := float32(byolComputeCountLimit.(float64))
			details.ByolComputeCountLimit = &tmp
		}
		if characterSet, ok := s.D.GetOkExists("character_set"); ok {
			tmp := characterSet.(string)
			details.CharacterSet = &tmp
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
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.AutonomousDatabaseSummaryDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbToolsDetails, ok := s.D.GetOkExists("db_tools_details"); ok {
			set := dbToolsDetails.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.DatabaseTool, len(interfaces))
			for i := range interfaces {
				stateDataIndex := dbToolsForSets(interfaces[i])
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_tools_details", stateDataIndex)
				converted, err := s.mapToDatabaseTool(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("db_tools_details") {
				details.DbToolsDetails = tmp
			}
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
		if encryptionKey, ok := s.D.GetOkExists("encryption_key"); ok {
			if tmpList := encryptionKey.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "encryption_key", 0)
				tmp, err := s.mapToAutonomousDatabaseEncryptionKeyDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.EncryptionKey = tmp
			}
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if inMemoryPercentage, ok := s.D.GetOkExists("in_memory_percentage"); ok {
			tmp := inMemoryPercentage.(int)
			details.InMemoryPercentage = &tmp
		}
		if isAccessControlEnabled, ok := s.D.GetOkExists("is_access_control_enabled"); ok {
			tmp := isAccessControlEnabled.(bool)
			details.IsAccessControlEnabled = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isAutoScalingForStorageEnabled, ok := s.D.GetOkExists("is_auto_scaling_for_storage_enabled"); ok {
			tmp := isAutoScalingForStorageEnabled.(bool)
			details.IsAutoScalingForStorageEnabled = &tmp
		}
		if isBackupRetentionLocked, ok := s.D.GetOkExists("is_backup_retention_locked"); ok {
			tmp := isBackupRetentionLocked.(bool)
			details.IsBackupRetentionLocked = &tmp
		}
		if isDataGuardEnabled, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			tmp := isDataGuardEnabled.(bool)
			details.IsDataGuardEnabled = &tmp
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isDevTier, ok := s.D.GetOkExists("is_dev_tier"); ok {
			tmp := isDevTier.(bool)
			details.IsDevTier = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if isLocalDataGuardEnabled, ok := s.D.GetOkExists("is_local_data_guard_enabled"); ok {
			tmp := isLocalDataGuardEnabled.(bool)
			details.IsLocalDataGuardEnabled = &tmp
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
		if ncharacterSet, ok := s.D.GetOkExists("ncharacter_set"); ok {
			tmp := ncharacterSet.(string)
			details.NcharacterSet = &tmp
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
			tmp := ocpuCount.(float32)
			details.OcpuCount = &tmp
		}
		if privateEndpointIp, ok := s.D.GetOkExists("private_endpoint_ip"); ok {
			tmp := privateEndpointIp.(string)
			details.PrivateEndpointIp = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			set := scheduledOperations.(*schema.Set)
			interfaces := set.List()
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
		if secretId, ok := s.D.GetOkExists("secret_id"); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok {
			tmp := secretVersionNumber.(int)
			details.SecretVersionNumber = &tmp
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
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
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
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
	case strings.ToLower("CROSS_TENANCY_DISASTER_RECOVERY"):
		details := oci_database.CreateCrossTenancyDisasterRecoveryDetails{}
		if disasterRecoveryType, ok := s.D.GetOkExists("disaster_recovery_type"); ok {
			details.DisasterRecoveryType = oci_database.DisasterRecoveryConfigurationDisasterRecoveryTypeEnum(disasterRecoveryType.(string))
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
		if backupRetentionPeriodInDays, ok := s.D.GetOkExists("backup_retention_period_in_days"); ok {
			tmp := backupRetentionPeriodInDays.(int)
			details.BackupRetentionPeriodInDays = &tmp
		}
		if byolComputeCountLimit, ok := s.D.GetOkExists("byol_compute_count_limit"); ok {
			tmp := float32(byolComputeCountLimit.(float64))
			details.ByolComputeCountLimit = &tmp
		}
		if characterSet, ok := s.D.GetOkExists("character_set"); ok {
			tmp := characterSet.(string)
			details.CharacterSet = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if computeCount, ok := s.D.GetOkExists("compute_count"); ok {
			tmp := computeCount.(float32)
			details.ComputeCount = &tmp
		}
		if computeModel, ok := s.D.GetOkExists("compute_model"); ok {
			details.ComputeModel = oci_database.CreateAutonomousDatabaseBaseComputeModelEnum(computeModel.(string))
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
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.AutonomousDatabaseSummaryDatabaseEditionEnum(databaseEdition.(string))
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
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if inMemoryPercentage, ok := s.D.GetOkExists("in_memory_percentage"); ok {
			tmp := inMemoryPercentage.(int)
			details.InMemoryPercentage = &tmp
		}
		if isAccessControlEnabled, ok := s.D.GetOkExists("is_access_control_enabled"); ok {
			tmp := isAccessControlEnabled.(bool)
			details.IsAccessControlEnabled = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isAutoScalingForStorageEnabled, ok := s.D.GetOkExists("is_auto_scaling_for_storage_enabled"); ok {
			tmp := isAutoScalingForStorageEnabled.(bool)
			details.IsAutoScalingForStorageEnabled = &tmp
		}
		if isBackupRetentionLocked, ok := s.D.GetOkExists("is_backup_retention_locked"); ok {
			tmp := isBackupRetentionLocked.(bool)
			details.IsBackupRetentionLocked = &tmp
		}
		if _, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			details.IsDataGuardEnabled = nil
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isDevTier, ok := s.D.GetOkExists("is_dev_tier"); ok {
			tmp := isDevTier.(bool)
			details.IsDevTier = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if _, ok := s.D.GetOkExists("is_local_data_guard_enabled"); ok {
			details.IsLocalDataGuardEnabled = nil
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
		//if maxCpuCoreCount, ok := s.D.GetOkExists("max_cpu_core_count"); ok {
		//	tmp := maxCpuCoreCount.(int)
		//	details.MaxCpuCoreCount = &tmp
		//}
		if ncharacterSet, ok := s.D.GetOkExists("ncharacter_set"); ok {
			tmp := ncharacterSet.(string)
			details.NcharacterSet = &tmp
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
			tmp := ocpuCount.(float32)
			details.OcpuCount = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}

		if resourcePoolLeaderId, ok := s.D.GetOkExists("resource_pool_leader_id"); ok {
			tmp := resourcePoolLeaderId.(string)
			details.ResourcePoolLeaderId = &tmp
		}
		if resourcePoolSummary, ok := s.D.GetOkExists("resource_pool_summary"); ok {
			t := fmt.Sprintf("%s rp create cross region disaster recovery", resourcePoolSummary)
			_, _ = io.WriteString(os.Stdout, t)
			if tmpList := resourcePoolSummary.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_pool_summary", 0)
				tmp, err := s.mapToResourcePoolSummary(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ResourcePoolSummary = &tmp
			}
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			set := scheduledOperations.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := scheduledOperationsForSets(interfaces[i])
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
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
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
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
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
		if backupRetentionPeriodInDays, ok := s.D.GetOkExists("backup_retention_period_in_days"); ok {
			tmp := backupRetentionPeriodInDays.(int)
			details.BackupRetentionPeriodInDays = &tmp
		}
		if byolComputeCountLimit, ok := s.D.GetOkExists("byol_compute_count_limit"); ok {
			tmp := float32(byolComputeCountLimit.(float64))
			details.ByolComputeCountLimit = &tmp
		}
		if characterSet, ok := s.D.GetOkExists("character_set"); ok {
			tmp := characterSet.(string)
			details.CharacterSet = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if computeCount, ok := s.D.GetOkExists("compute_count"); ok {
			tmp := float32(computeCount.(float64))
			details.ComputeCount = &tmp
		}
		if computeModel, ok := s.D.GetOkExists("compute_model"); ok {
			details.ComputeModel = oci_database.CreateAutonomousDatabaseBaseComputeModelEnum(computeModel.(string))
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
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.AutonomousDatabaseSummaryDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbToolsDetails, ok := s.D.GetOkExists("db_tools_details"); ok {
			set := dbToolsDetails.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.DatabaseTool, len(interfaces))
			for i := range interfaces {
				stateDataIndex := dbToolsForSets(interfaces[i])
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_tools_details", stateDataIndex)
				converted, err := s.mapToDatabaseTool(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("db_tools_details") {
				details.DbToolsDetails = tmp
			}
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
		if encryptionKey, ok := s.D.GetOkExists("encryption_key"); ok {
			if tmpList := encryptionKey.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "encryption_key", 0)
				tmp, err := s.mapToAutonomousDatabaseEncryptionKeyDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.EncryptionKey = tmp
			}
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if inMemoryPercentage, ok := s.D.GetOkExists("in_memory_percentage"); ok {
			tmp := inMemoryPercentage.(int)
			details.InMemoryPercentage = &tmp
		}
		if isAccessControlEnabled, ok := s.D.GetOkExists("is_access_control_enabled"); ok {
			tmp := isAccessControlEnabled.(bool)
			details.IsAccessControlEnabled = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isAutoScalingForStorageEnabled, ok := s.D.GetOkExists("is_auto_scaling_for_storage_enabled"); ok {
			tmp := isAutoScalingForStorageEnabled.(bool)
			details.IsAutoScalingForStorageEnabled = &tmp
		}
		if isBackupRetentionLocked, ok := s.D.GetOkExists("is_backup_retention_locked"); ok {
			tmp := isBackupRetentionLocked.(bool)
			details.IsBackupRetentionLocked = &tmp
		}
		if _, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			details.IsDataGuardEnabled = nil
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isDevTier, ok := s.D.GetOkExists("is_dev_tier"); ok {
			tmp := isDevTier.(bool)
			details.IsDevTier = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if _, ok := s.D.GetOkExists("is_local_data_guard_enabled"); ok {
			details.IsLocalDataGuardEnabled = nil
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
		//if maxCpuCoreCount, ok := s.D.GetOkExists("max_cpu_core_count"); ok {
		//	tmp := maxCpuCoreCount.(int)
		//	details.MaxCpuCoreCount = &tmp
		//}
		if ncharacterSet, ok := s.D.GetOkExists("ncharacter_set"); ok {
			tmp := ncharacterSet.(string)
			details.NcharacterSet = &tmp
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
		if privateEndpointIp, ok := s.D.GetOkExists("private_endpoint_ip"); ok {
			tmp := privateEndpointIp.(string)
			details.PrivateEndpointIp = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if resourcePoolLeaderId, ok := s.D.GetOkExists("resource_pool_leader_id"); ok {
			tmp := resourcePoolLeaderId.(string)
			details.ResourcePoolLeaderId = &tmp
		}
		if resourcePoolSummary, ok := s.D.GetOkExists("resource_pool_summary"); ok {
			t := fmt.Sprintf("%s rp create database", resourcePoolSummary)
			_, _ = io.WriteString(os.Stdout, t)
			if tmpList := resourcePoolSummary.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_pool_summary", 0)
				tmp, err := s.mapToResourcePoolSummary(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ResourcePoolSummary = &tmp
			}
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			set := scheduledOperations.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := scheduledOperationsForSets(interfaces[i])
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
		if secretId, ok := s.D.GetOkExists("secret_id"); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok {
			tmp := secretVersionNumber.(int)
			details.SecretVersionNumber = &tmp
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
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
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
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
		if backupRetentionPeriodInDays, ok := s.D.GetOkExists("backup_retention_period_in_days"); ok {
			tmp := backupRetentionPeriodInDays.(int)
			details.BackupRetentionPeriodInDays = &tmp
		}
		if byolComputeCountLimit, ok := s.D.GetOkExists("byol_compute_count_limit"); ok {
			tmp := float32(byolComputeCountLimit.(float64))
			details.ByolComputeCountLimit = &tmp
		}
		if characterSet, ok := s.D.GetOkExists("character_set"); ok {
			tmp := characterSet.(string)
			details.CharacterSet = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if computeCount, ok := s.D.GetOkExists("compute_count"); ok {
			tmp := float32(computeCount.(float64))
			details.ComputeCount = &tmp
		}
		if computeModel, ok := s.D.GetOkExists("compute_model"); ok {
			details.ComputeModel = oci_database.CreateAutonomousDatabaseBaseComputeModelEnum(computeModel.(string))
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
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.AutonomousDatabaseSummaryDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbToolsDetails, ok := s.D.GetOkExists("db_tools_details"); ok {
			set := dbToolsDetails.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.DatabaseTool, len(interfaces))
			for i := range interfaces {
				stateDataIndex := dbToolsForSets(interfaces[i])
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_tools_details", stateDataIndex)
				converted, err := s.mapToDatabaseTool(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("db_tools_details") {
				details.DbToolsDetails = tmp
			}
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
		if encryptionKey, ok := s.D.GetOkExists("encryption_key"); ok {
			if tmpList := encryptionKey.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "encryption_key", 0)
				tmp, err := s.mapToAutonomousDatabaseEncryptionKeyDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.EncryptionKey = tmp
			}
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if inMemoryPercentage, ok := s.D.GetOkExists("in_memory_percentage"); ok {
			tmp := inMemoryPercentage.(int)
			details.InMemoryPercentage = &tmp
		}
		if isAccessControlEnabled, ok := s.D.GetOkExists("is_access_control_enabled"); ok {
			tmp := isAccessControlEnabled.(bool)
			details.IsAccessControlEnabled = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isAutoScalingForStorageEnabled, ok := s.D.GetOkExists("is_auto_scaling_for_storage_enabled"); ok {
			tmp := isAutoScalingForStorageEnabled.(bool)
			details.IsAutoScalingForStorageEnabled = &tmp
		}
		if isBackupRetentionLocked, ok := s.D.GetOkExists("is_backup_retention_locked"); ok {
			tmp := isBackupRetentionLocked.(bool)
			details.IsBackupRetentionLocked = &tmp
		}
		if _, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
			details.IsDataGuardEnabled = nil
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isDevTier, ok := s.D.GetOkExists("is_dev_tier"); ok {
			tmp := isDevTier.(bool)
			details.IsDevTier = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if _, ok := s.D.GetOkExists("is_local_data_guard_enabled"); ok {
			details.IsLocalDataGuardEnabled = nil
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
		//if maxCpuCoreCount, ok := s.D.GetOkExists("max_cpu_core_count"); ok {
		//	tmp := maxCpuCoreCount.(int)
		//	details.MaxCpuCoreCount = &tmp
		//}
		if ncharacterSet, ok := s.D.GetOkExists("ncharacter_set"); ok {
			tmp := ncharacterSet.(string)
			details.NcharacterSet = &tmp
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
		if privateEndpointIp, ok := s.D.GetOkExists("private_endpoint_ip"); ok {
			tmp := privateEndpointIp.(string)
			details.PrivateEndpointIp = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if resourcePoolLeaderId, ok := s.D.GetOkExists("resource_pool_leader_id"); ok {
			tmp := resourcePoolLeaderId.(string)
			details.ResourcePoolLeaderId = &tmp
		}
		if resourcePoolSummary, ok := s.D.GetOkExists("resource_pool_summary"); ok {
			t := fmt.Sprintf("%s rp create none", resourcePoolSummary)
			_, _ = io.WriteString(os.Stdout, t)
			if tmpList := resourcePoolSummary.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_pool_summary", 0)
				tmp, err := s.mapToResourcePoolSummary(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ResourcePoolSummary = &tmp
			}
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			set := scheduledOperations.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database.ScheduledOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := scheduledOperationsForSets(interfaces[i])
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
		if secretId, ok := s.D.GetOkExists("secret_id"); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok {
			tmp := secretVersionNumber.(int)
			details.SecretVersionNumber = &tmp
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
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
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
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
	case strings.ToLower("UNDELETE_ADB"):
		details := oci_database.UndeleteAutonomousDatabaseDetails{}
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
		if backupRetentionPeriodInDays, ok := s.D.GetOkExists("backup_retention_period_in_days"); ok {
			tmp := backupRetentionPeriodInDays.(int)
			details.BackupRetentionPeriodInDays = &tmp
		}
		if characterSet, ok := s.D.GetOkExists("character_set"); ok {
			tmp := characterSet.(string)
			details.CharacterSet = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if computeCount, ok := s.D.GetOkExists("compute_count"); ok {
			tmp := float32(computeCount.(float64))
			details.ComputeCount = &tmp
		}
		if computeModel, ok := s.D.GetOkExists("compute_model"); ok {
			details.ComputeModel = oci_database.CreateAutonomousDatabaseBaseComputeModelEnum(computeModel.(string))
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
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.AutonomousDatabaseSummaryDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbName, ok := s.D.GetOkExists("db_name"); ok {
			tmp := dbName.(string)
			details.DbName = &tmp
		}
		if dbToolsDetails, ok := s.D.GetOkExists("db_tools_details"); ok {
			interfaces := dbToolsDetails.([]interface{})
			tmp := make([]oci_database.DatabaseTool, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_tools_details", stateDataIndex)
				converted, err := s.mapToDatabaseTool(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("db_tools_details") {
				details.DbToolsDetails = tmp
			}
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
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if inMemoryPercentage, ok := s.D.GetOkExists("in_memory_percentage"); ok {
			tmp := inMemoryPercentage.(int)
			details.InMemoryPercentage = &tmp
		}
		if isAccessControlEnabled, ok := s.D.GetOkExists("is_access_control_enabled"); ok {
			tmp := isAccessControlEnabled.(bool)
			details.IsAccessControlEnabled = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isAutoScalingForStorageEnabled, ok := s.D.GetOkExists("is_auto_scaling_for_storage_enabled"); ok {
			tmp := isAutoScalingForStorageEnabled.(bool)
			details.IsAutoScalingForStorageEnabled = &tmp
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
		if isLocalDataGuardEnabled, ok := s.D.GetOkExists("is_local_data_guard_enabled"); ok {
			tmp := isLocalDataGuardEnabled.(bool)
			details.IsLocalDataGuardEnabled = &tmp
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
		if ncharacterSet, ok := s.D.GetOkExists("ncharacter_set"); ok {
			tmp := ncharacterSet.(string)
			details.NcharacterSet = &tmp
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
			tmp := ocpuCount.(float32)
			details.OcpuCount = &tmp
		}
		if privateEndpointIp, ok := s.D.GetOkExists("private_endpoint_ip"); ok {
			tmp := privateEndpointIp.(string)
			details.PrivateEndpointIp = &tmp
		}
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if resourcePoolLeaderId, ok := s.D.GetOkExists("resource_pool_leader_id"); ok {
			tmp := resourcePoolLeaderId.(string)
			details.ResourcePoolLeaderId = &tmp
		}
		if resourcePoolSummary, ok := s.D.GetOkExists("resource_pool_summary"); ok {
			if tmpList := resourcePoolSummary.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_pool_summary", 0)
				tmp, err := s.mapToResourcePoolSummary(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ResourcePoolSummary = &tmp
			}
		}
		if scheduledOperations, ok := s.D.GetOkExists("scheduled_operations"); ok {
			set := scheduledOperations.(*schema.Set)
			interfaces := set.List()
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
		if secretId, ok := s.D.GetOkExists("secret_id"); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok {
			tmp := secretVersionNumber.(int)
			details.SecretVersionNumber = &tmp
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
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
func (s *DatabaseAutonomousDatabaseResourceCrud) updateSubscription(subscriptionId string) error {
	changeSubscriptionRequest := oci_database.ChangeAutonomousDatabaseSubscriptionRequest{}

	idTmp := s.D.Id()
	changeSubscriptionRequest.AutonomousDatabaseId = &idTmp

	subscriptionTmp := subscriptionId
	changeSubscriptionRequest.SubscriptionId = &subscriptionTmp

	changeSubscriptionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeAutonomousDatabaseSubscription(context.Background(), changeSubscriptionRequest)
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

	//	Local peer
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
								if (wantedStandByState == oci_database.AutonomousDatabaseStandbySummaryLifecycleStateAvailable) || (wantedStandByState == oci_database.AutonomousDatabaseStandbySummaryLifecycleStateStandby) {
									if err := s.switchoverDatabase(""); err != nil {
										s.D.Set("switchover_to", oldRaw.(string))
										return err
									}

									s.D.Set("switchover_to", switchoverTo.(string))
								} else {
									s.D.Set("switchover_to", oldRaw.(string))

									return fmt.Errorf("Autonomous standby state: %s is not present in [AVAILABLE, STANDBY] states", wantedStandByState)
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

	//	Remote peer
	if _, ok := s.D.GetOkExists("switchover_to_remote_peer_id"); ok && s.D.HasChange("switchover_to_remote_peer_id") {
		oldIdRaw, newIdRaw := s.D.GetChange("switchover_to_remote_peer_id")
		oldId := strings.ToLower(strings.TrimSpace(oldIdRaw.(string)))
		newId := strings.ToLower(strings.TrimSpace(newIdRaw.(string)))

		if newId != "" {
			_, dgRegionTypeExists := s.D.GetOkExists("dataguard_region_type")
			_, dgRoleExists := s.D.GetOkExists("role")
			if !dgRegionTypeExists || !dgRoleExists {
				return fmt.Errorf("Autonomous Data Guard not found in enabled state")
			}

			if err := s.switchoverDatabase(newId); err != nil {
				s.D.Set("switchover_to_remote_peer_id", oldId)
				return err
			}
		}
		s.D.Set("switchover_to_remote_peer_id", newId)
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) switchoverDatabase(peerDbId string) error {
	request := oci_database.SwitchoverAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp
	if peerDbId != "" {
		request.PeerDbId = &peerDbId
	}

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

func (s *DatabaseAutonomousDatabaseResourceCrud) ShrinkAutonomousDatabase(state oci_database.AutonomousDatabaseLifecycleStateEnum) error {
	request := oci_database.ShrinkAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	if _, err := s.Client.ShrinkAutonomousDatabase(context.Background(), request); err != nil {
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

	if keyVersionId, ok := s.D.GetOkExists("key_version_id"); ok {
		tmp := keyVersionId.(string)
		request.KeyVersionId = &tmp
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

func validateLocalAdgCreate(s *DatabaseAutonomousDatabaseResourceCrud) (bool, int) {
	localAdg := false
	var localAdgDataLossLimit int
	if isLocalDataGuardEnabled, ok := s.D.GetOkExists("is_local_data_guard_enabled"); ok && isLocalDataGuardEnabled.(bool) == true {
		localAdg = true
		if localAdgAutoFailoverMaxDataLossLimit, ok := s.D.GetOkExists("local_adg_auto_failover_max_data_loss_limit"); ok {
			localAdgDataLossLimit = localAdgAutoFailoverMaxDataLossLimit.(int)
		}
	}
	return localAdg, localAdgDataLossLimit
}

func (s *DatabaseAutonomousDatabaseResourceCrud) UpdateLocalAdg(adg bool, limit int) error {
	updateLocalAdgRequest := oci_database.UpdateAutonomousDatabaseRequest{}
	updateLocalAdgRequest.IsLocalDataGuardEnabled = &adg
	updateLocalAdgRequest.LocalAdgAutoFailoverMaxDataLossLimit = &limit

	tmp := s.D.Id()
	updateLocalAdgRequest.AutonomousDatabaseId = &tmp

	updateLocalAdgRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousDatabase(context.Background(), updateLocalAdgRequest)
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

func scheduledOperationsForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if dayOfWeek, ok := m["day_of_week"]; ok && dayOfWeek != "" {
		if tmpList, ok := dayOfWeek.([]interface{}); ok && len(tmpList) > 0 && tmpList[0] != "" {
			buf.WriteString("day_of_week-")
			for _, dayOfWeekTemp := range tmpList {
				buf.WriteString(fmt.Sprintf("%v-", dayOfWeekTemp))
			}
		}
	}
	if startTime, ok := m["scheduled_start_time"]; ok && startTime != "" {
		buf.WriteString(fmt.Sprintf("%v-", startTime))
	}
	if stopTime, ok := m["scheduled_stop_time"]; ok && stopTime != "" {
		buf.WriteString(fmt.Sprintf("%v-", strings.ToLower(stopTime.(string))))
	}
	return utils.GetStringHashcode(buf.String())
}

func dbToolsForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if name, ok := m["name"]; ok && name != "" {
		buf.WriteString(fmt.Sprintf("%v-", name))
	}
	if computeCount, ok := m["compute_count"]; ok && computeCount != "" {
		buf.WriteString(fmt.Sprintf("%v-", computeCount))
	}
	if isEnabled, ok := m["is_enabled"]; ok && isEnabled != "" {
		buf.WriteString(fmt.Sprintf("%v-", isEnabled))
	}
	if maxIdleTimeInMinutes, ok := m["max_idle_time_in_minutes"]; ok && maxIdleTimeInMinutes != "" {
		buf.WriteString(fmt.Sprintf("%v-", maxIdleTimeInMinutes))
	}
	return utils.GetStringHashcode(buf.String())
}
