// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousContainerDatabaseAddStandbyResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		Create: createDatabaseAutonomousContainerDatabaseAddStandby,
		Read:   readDatabaseAutonomousContainerDatabaseAddStandby,
		Delete: deleteDatabaseAutonomousContainerDatabaseAddStandby,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"fast_start_fail_over_lag_limit_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_automatic_failover_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
									"backup_retention_policy_on_terminate": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"dbrs_policy_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"id": {
										Type:     schema.TypeString,
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
			"peer_autonomous_vm_cluster_id": {
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
			"standby_maintenance_buffer_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"autonomous_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"available_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"backup_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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
									"is_retention_lock_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"internet_proxy": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vpc_password": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vpc_user": {
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
					},
				},
			},
			"cloud_autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compute_model": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dataguard": {
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
						"automatic_failover_target": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"autonomous_container_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fast_start_fail_over_lag_limit_in_seconds": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_automatic_failover_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protection_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"redo_transport_mode": {
							Type:     schema.TypeString,
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
						"time_lag_refreshed_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_role_changed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_synced": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"transport_lag": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dataguard_group_members": {
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
						"automatic_failover_target": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"autonomous_container_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fast_start_fail_over_lag_limit_in_seconds": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_automatic_failover_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protection_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"redo_transport_mode": {
							Type:     schema.TypeString,
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
						"time_lag_refreshed_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_role_changed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_synced": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"transport_lag": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"db_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_split_threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"db_unique_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"distribution_affinity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dst_file_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"infrastructure_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_data_guard_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_dst_file_update_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_multiple_standby": {
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
			"kms_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kms_key_version_id": {
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
			"list_one_off_patches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
						"skip_ru": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
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
			"memory_per_compute_unit_in_gbs": {
				Type:     schema.TypeFloat,
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
			"next_maintenance_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"okv_end_point_group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_model": {
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
			"service_level_agreement_type": {
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
			"vault_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version_preference": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_failover_reservation": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createDatabaseAutonomousContainerDatabaseAddStandby(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousContainerDatabaseAddStandby(d *schema.ResourceData, m interface{}) error {
	//sync := &DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud{}
	//sync.D = d
	//sync.Client = m.(*client.OracleClients).DatabaseClient()
	//sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	//
	//return tfresource.ReadResource(sync)
	return nil
}

func deleteDatabaseAutonomousContainerDatabaseAddStandby(d *schema.ResourceData, m interface{}) error {
	//sync := &DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud{}
	//sync.D = d
	//sync.Client = m.(*client.OracleClients).DatabaseClient()
	//sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	//sync.DisableNotFoundRetries = true
	//
	//return tfresource.DeleteResource(d, sync)
	return nil
}

type DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousContainerDatabase
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) Get() error {
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

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateProvisioning),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateRestoring),
	}
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminating),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateUnavailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) Delete() error {
	request := oci_database.TerminateAutonomousContainerDatabaseRequest{}
	log.Println("adbcc: delete called")
	var id string
	if dgMembers, ok := s.D.GetOkExists("dataguard_group_members"); ok {
		interfaces := dgMembers.([]interface{})
		for i := range interfaces {
			converted := interfaces[i].(map[string]interface{})
			if converted["role"] == "STANDBY" {
				id = converted["autonomous_container_database_id"].(string)
				request.AutonomousContainerDatabaseId = &id
				log.Println("adbcc: ", id)
			}
		}
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
	response, err := s.Client.TerminateAutonomousContainerDatabase(context.Background(), request)
	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return err
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) Create() error {
	request := oci_database.AddStandbyAutonomousContainerDatabaseRequest{}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	if fastStartFailOverLagLimitInSeconds, ok := s.D.GetOkExists("fast_start_fail_over_lag_limit_in_seconds"); ok {
		tmp := fastStartFailOverLagLimitInSeconds.(int)
		request.FastStartFailOverLagLimitInSeconds = &tmp
	}

	if isAutomaticFailoverEnabled, ok := s.D.GetOkExists("is_automatic_failover_enabled"); ok {
		tmp := isAutomaticFailoverEnabled.(bool)
		request.IsAutomaticFailoverEnabled = &tmp
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

	if peerAutonomousVmClusterId, ok := s.D.GetOkExists("peer_autonomous_vm_cluster_id"); ok {
		tmp := peerAutonomousVmClusterId.(string)
		request.PeerAutonomousVmClusterId = &tmp
	}

	if peerCloudAutonomousVmClusterId, ok := s.D.GetOkExists("peer_cloud_autonomous_vm_cluster_id"); ok {
		tmp := peerCloudAutonomousVmClusterId.(string)
		request.PeerCloudAutonomousVmClusterId = &tmp
	}

	if peerDbUniqueName, ok := s.D.GetOkExists("peer_db_unique_name"); ok {
		tmp := peerDbUniqueName.(string)
		request.PeerDbUniqueName = &tmp
	}

	if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
		request.ProtectionMode = oci_database.AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum(protectionMode.(string))
	}

	if standbyMaintenanceBufferInDays, ok := s.D.GetOkExists("standby_maintenance_buffer_in_days"); ok {
		tmp := standbyMaintenanceBufferInDays.(int)
		request.StandbyMaintenanceBufferInDays = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.AddStandbyAutonomousContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.AutonomousContainerDatabase

	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomouscontainerdatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) SetData() error {
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
		s.D.Set("backup_config", []interface{}{AutonomousContainerDatabaseAddStandbyBackupConfigToMap(s.Res.BackupConfig)})
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

	if s.Res.Dataguard != nil {
		s.D.Set("dataguard", []interface{}{AutonomousContainerDatabaseDataguardToMap(s.Res.Dataguard)})
		s.D.Set("protection_mode", s.Res.Dataguard.ProtectionMode)
		s.D.Set("fast_start_fail_over_lag_limit_in_seconds", s.Res.Dataguard.FastStartFailOverLagLimitInSeconds)
		s.D.Set("is_automatic_failover_enabled", s.Res.Dataguard.IsAutomaticFailoverEnabled)
	} else {
		s.D.Set("dataguard", nil)
	}

	dataguardGroupMembers := []interface{}{}
	for _, item := range s.Res.DataguardGroupMembers {
		dataguardGroupMembers = append(dataguardGroupMembers, AutonomousContainerDatabaseDataguardToMap(&item))
	}
	s.D.Set("dataguard_group_members", dataguardGroupMembers)

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbSplitThreshold != nil {
		s.D.Set("db_split_threshold", *s.Res.DbSplitThreshold)
	}

	if s.Res.DbUniqueName != nil {
		s.D.Set("db_unique_name", *s.Res.DbUniqueName)
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

	s.D.Set("distribution_affinity", s.Res.DistributionAffinity)

	if s.Res.DstFileVersion != nil {
		s.D.Set("dst_file_version", *s.Res.DstFileVersion)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("infrastructure_type", s.Res.InfrastructureType)

	if s.Res.IsDataGuardEnabled != nil {
		s.D.Set("is_data_guard_enabled", *s.Res.IsDataGuardEnabled)
	}

	if s.Res.IsDstFileUpdateEnabled != nil {
		s.D.Set("is_dst_file_update_enabled", *s.Res.IsDstFileUpdateEnabled)
	}

	if s.Res.IsMultipleStandby != nil {
		s.D.Set("is_multiple_standby", *s.Res.IsMultipleStandby)
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

	s.D.Set("list_one_off_patches", s.Res.ListOneOffPatches)

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MemoryPerComputeUnitInGBs != nil {
		s.D.Set("memory_per_compute_unit_in_gbs", *s.Res.MemoryPerComputeUnitInGBs)
	}

	if s.Res.MemoryPerOracleComputeUnitInGBs != nil {
		s.D.Set("memory_per_oracle_compute_unit_in_gbs", *s.Res.MemoryPerOracleComputeUnitInGBs)
	}

	s.D.Set("net_services_architecture", s.Res.NetServicesArchitecture)

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	if s.Res.OkvEndPointGroupName != nil {
		s.D.Set("okv_end_point_group_name", *s.Res.OkvEndPointGroupName)
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

	if s.Res.VmFailoverReservation != nil {
		s.D.Set("vm_failover_reservation", *s.Res.VmFailoverReservation)
	}

	return nil
}

func AutonomousContainerDatabaseAddStandbyBackupConfigToMap(obj *oci_database.AutonomousContainerDatabaseBackupConfig) map[string]interface{} {
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

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) AutonomousContainerDatabaseDataguardToMap(obj oci_database.AutonomousContainerDatabaseDataguard) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplyLag != nil {
		result["apply_lag"] = string(*obj.ApplyLag)
	}

	if obj.ApplyRate != nil {
		result["apply_rate"] = string(*obj.ApplyRate)
	}

	if obj.AutomaticFailoverTarget != nil {
		result["automatic_failover_target"] = string(*obj.AutomaticFailoverTarget)
	}

	if obj.AutonomousContainerDatabaseId != nil {
		result["autonomous_container_database_id"] = string(*obj.AutonomousContainerDatabaseId)
	}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.FastStartFailOverLagLimitInSeconds != nil {
		result["fast_start_fail_over_lag_limit_in_seconds"] = int(*obj.FastStartFailOverLagLimitInSeconds)
	}

	if obj.IsAutomaticFailoverEnabled != nil {
		result["is_automatic_failover_enabled"] = bool(*obj.IsAutomaticFailoverEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["protection_mode"] = string(obj.ProtectionMode)

	if obj.RedoTransportMode != nil {
		result["redo_transport_mode"] = string(*obj.RedoTransportMode)
	}

	result["role"] = string(obj.Role)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLagRefreshedOn != nil {
		result["time_lag_refreshed_on"] = obj.TimeLagRefreshedOn.String()
	}

	if obj.TimeLastRoleChanged != nil {
		result["time_last_role_changed"] = obj.TimeLastRoleChanged.String()
	}

	if obj.TimeLastSynced != nil {
		result["time_last_synced"] = obj.TimeLastSynced.String()
	}

	if obj.TransportLag != nil {
		result["transport_lag"] = string(*obj.TransportLag)
	}

	return result
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) AutonomousDatabaseKeyHistoryEntryToMap(obj oci_database.AutonomousDatabaseKeyHistoryEntry) map[string]interface{} {
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

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) mapToBackupDestinationDetails(fieldKeyFormat string) (oci_database.BackupDestinationDetails, error) {
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

	if isRetentionLockEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_retention_lock_enabled")); ok {
		tmp := isRetentionLockEnabled.(bool)
		result.IsRetentionLockEnabled = &tmp
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

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) BackupDestinationDetailsToMap(obj oci_database.BackupDestinationDetails) map[string]interface{} {
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

	result["type"] = string(obj.Type)

	if obj.VpcPassword != nil {
		result["vpc_password"] = string(*obj.VpcPassword)
	}

	if obj.IsRetentionLockEnabled != nil {
		result["is_retention_lock_enabled"] = bool(*obj.IsRetentionLockEnabled)
	}

	if obj.VpcUser != nil {
		result["vpc_user"] = string(*obj.VpcUser)
	}

	return result
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) DayOfWeekToMap(obj oci_database.DayOfWeek) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) MaintenanceWindowToMap(obj *oci_database.MaintenanceWindow) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomActionTimeoutInMins != nil {
		result["custom_action_timeout_in_mins"] = int(*obj.CustomActionTimeoutInMins)
	}

	daysOfWeek := []interface{}{}
	for _, item := range obj.DaysOfWeek {
		daysOfWeek = append(daysOfWeek, DayOfWeekToMap(item))
	}
	result["days_of_week"] = daysOfWeek

	result["hours_of_day"] = obj.HoursOfDay

	if obj.IsCustomActionTimeoutEnabled != nil {
		result["is_custom_action_timeout_enabled"] = bool(*obj.IsCustomActionTimeoutEnabled)
	}

	if obj.IsMonthlyPatchingEnabled != nil {
		result["is_monthly_patching_enabled"] = bool(*obj.IsMonthlyPatchingEnabled)
	}

	if obj.LeadTimeInWeeks != nil {
		result["lead_time_in_weeks"] = int(*obj.LeadTimeInWeeks)
	}

	months := []interface{}{}
	for _, item := range obj.Months {
		months = append(months, MonthToMap(item))
	}
	result["months"] = months

	result["patching_mode"] = string(obj.PatchingMode)

	result["preference"] = string(obj.Preference)

	result["skip_ru"] = obj.SkipRu

	result["weeks_of_month"] = obj.WeeksOfMonth

	return result
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) MonthToMap(obj oci_database.Month) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseAutonomousContainerDatabaseAddStandbyResourceCrud) mapToPeerAutonomousContainerDatabaseBackupConfig(fieldKeyFormat string) (oci_database.PeerAutonomousContainerDatabaseBackupConfig, error) {
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
