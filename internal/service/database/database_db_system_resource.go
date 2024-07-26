// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseDbSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("2h"),
			Update: tfresource.GetTimeoutDuration("2h"),
			Delete: tfresource.GetTimeoutDuration("2h"),
		},
		Create: createDatabaseDbSystem,
		Read:   readDatabaseDbSystem,
		Update: updateDatabaseDbSystem,
		Delete: deleteDatabaseDbSystem,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_home": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
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
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													DiffSuppressFunc: disableAutoBackupDbSystemSuppressfunc,
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
													DiffSuppressFunc: disableAutoBackupDbSystemSuppressfunc,
												},
												"backup_deletion_policy": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"backup_destination_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
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
															"type": {
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
												"run_immediate_full_backup": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"db_domain": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
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
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"vault_id": {
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
										Optional: true,
										Computed: true,
										ForceNew: true,
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
						"create_async": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"database_software_image_id": {
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

						// Computed
						"id": {
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
				},
			},
			"hostname": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DbSystemHostnameDiffSuppress,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ssh_public_keys": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"backup_network_nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"backup_subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cluster_name": {
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
			"data_collection_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_diagnostics_events_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_health_monitoring_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_incident_logs_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"data_storage_percentage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"data_storage_size_in_gb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"database_edition": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"db_system_options": {
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
						"storage_management": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"disk_redundancy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"fault_domains": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kms_key_version_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_database.DbSystemLicenseModelLicenseIncluded),
					string(oci_database.DbSystemLicenseModelBringYourOwnLicense)}, false),
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

									// Optional
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

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

									// Optional
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

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
						"skip_ru": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
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
			"node_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"private_ip": {
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
					"DB_SYSTEM",
					"NONE",
				}, true),
			},
			"source_db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sparse_diskgroup": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"storage_volume_performance_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"iorm_config_cache": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"db_system_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						// Optional

						// Computed
						"db_plans": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"db_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"flash_cache_limit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"share": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"objective": {
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
			"last_maintenance_run_id": {
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
			"listener_port": {
				Type:     schema.TypeInt,
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
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"next_maintenance_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"point_in_time_data_disk_clone_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reco_storage_size_in_gb": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"scan_dns_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scan_dns_record_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scan_ip_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vip_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return createDBSystemResource(d, sync)
}

func readDatabaseDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func updateDatabaseDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseDbSystemResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.DbSystem
	DbHome                 *oci_database.DbHome
	Database               *oci_database.Database
	DisableNotFoundRetries bool
}

func (s *DatabaseDbSystemResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDbSystemResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateProvisioning),
	}
}

func (s *DatabaseDbSystemResourceCrud) CreatedTarget() []string {
	if createAsyn, ok := s.D.GetOk("create_async"); ok {
		tmp := createAsyn.(bool)
		if tmp {
			return []string{
				string(oci_database.DbSystemLifecycleStateAvailable),
				string(oci_database.DbSystemLifecycleStateProvisioning),
			}
		}
	}
	return []string{
		string(oci_database.DbSystemLifecycleStateAvailable),
	}
}

func (s *DatabaseDbSystemResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateTerminating),
	}
}

func (s *DatabaseDbSystemResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateTerminated),
		string(oci_database.DbSystemLifecycleStateMigrated),
	}
}

func (s *DatabaseDbSystemResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateUpdating),
		string(oci_database.DbSystemLifecycleStateUpgrading),
	}
}

func (s *DatabaseDbSystemResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateAvailable),
	}
}

func (s *DatabaseDbSystemResourceCrud) Create() error {
	request := oci_database.LaunchDbSystemRequest{}
	err := s.populateTopLevelPolymorphicLaunchDbSystemRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.LaunchDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem

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

	err = s.getDbHomeInfo()
	if err != nil {
		log.Printf("[WARN] Could not get info about the first DbHome in the dbSystem: %v", err)
	}

	return nil
}

func (s *DatabaseDbSystemResourceCrud) Get() error {
	request := oci_database.GetDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem

	err = s.getDbHomeInfo()
	if err != nil {
		log.Printf("[WARN] Could not get info about the first DbHome in the dbSystem: %v", err)
	}

	return nil
}

func (s *DatabaseDbSystemResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateDbSystemRequest{}

	if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
		set := backupNetworkNsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("backup_network_nsg_ids") {
			request.BackupNetworkNsgIds = tmp
		}
	}

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok && s.D.HasChange("cpu_core_count") {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}

	if dataCollectionOptions, ok := s.D.GetOkExists("data_collection_options"); ok {
		if tmpList := dataCollectionOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_collection_options", 0)
			tmp, err := s.mapToDataCollectionOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataCollectionOptions = &tmp
		}
	}

	if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok && s.D.HasChange("data_storage_size_in_gb") {
		tmp := dataStorageSizeInGB.(int)
		request.DataStorageSizeInGBs = &tmp
	}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok && s.D.HasChange("license_model") {
		if err := s.sendUpdateForLicenseModel(s.D.Id(), licenseModel); err != nil {
			return err
		}
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
			request.NsgIds = tmp
		}
	}

	if recoStorageSizeInGBs, ok := s.D.GetOkExists("reco_storage_size_in_gb"); ok && s.D.HasChange("reco_storage_size_in_gb") {
		tmp := recoStorageSizeInGBs.(int)
		request.RecoStorageSizeInGBs = &tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok && s.D.HasChange("shape") {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
		set := sshPublicKeys.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ssh_public_keys") {
			request.SshPublicKeys = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateDbSystem(context.Background(), request)
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

	s.Res = &response.DbSystem

	// Check lifecycle state of db system
	getDbSystemResponse, err := waitForDbSystemIfItIsUpdating(s.Res.Id, s.Client, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		err = s.SetData()
		if err != nil {
			log.Printf("[ERROR] error setting data after polling error on the dbSystem: %v", err)
		}
		return fmt.Errorf("[ERROR] unable to get dbSystem after the Update: %v", err)
	}

	s.Res = &getDbSystemResponse.DbSystem

	err = s.SetData()
	if err != nil {
		return fmt.Errorf("[ERROR] error setting data after dbsystem update but before database Update: %v", err)
	}

	return s.UpdateDatabaseOperation()
}

func (s *DatabaseDbSystemResourceCrud) Delete() error {
	request := oci_database.TerminateDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.TerminateDbSystem(context.Background(), request)
	return err
}

func (s *DatabaseDbSystemResourceCrud) SetData() error {

	if s.DbHome != nil {
		s.D.Set("db_home", []interface{}{s.DbHomeToMap(s.DbHome)})
	}

	if source, ok := s.D.GetOkExists("source"); !ok || source.(string) == "" {
		s.D.Set("source", "NONE")
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.BackupNetworkNsgIds != nil {
		backupNetworkNsgIds := []interface{}{}
		for _, item := range s.Res.BackupNetworkNsgIds {
			backupNetworkNsgIds = append(backupNetworkNsgIds, item)
		}
		s.D.Set("backup_network_nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, backupNetworkNsgIds))
	}

	if s.Res.BackupSubnetId != nil {
		s.D.Set("backup_subnet_id", *s.Res.BackupSubnetId)
	}

	if s.Res.ClusterName != nil {
		s.D.Set("cluster_name", *s.Res.ClusterName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DataCollectionOptions != nil {
		s.D.Set("data_collection_options", []interface{}{DataCollectionOptionsToMap(s.Res.DataCollectionOptions)})
	} else {
		s.D.Set("data_collection_options", nil)
	}

	if s.Res.DataStoragePercentage != nil {
		s.D.Set("data_storage_percentage", *s.Res.DataStoragePercentage)
	}

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	s.D.Set("database_edition", s.Res.DatabaseEdition)

	if s.Res.DbSystemOptions != nil {
		s.D.Set("db_system_options", []interface{}{DbSystemOptionsToMap(s.Res.DbSystemOptions)})
	} else {
		s.D.Set("db_system_options", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("disk_redundancy", s.Res.DiskRedundancy)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	s.D.Set("fault_domains", s.Res.FaultDomains)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.IormConfigCache != nil {
		s.D.Set("iorm_config_cache", []interface{}{ExadataIormConfigToMap(s.Res.IormConfigCache)})
	} else {
		s.D.Set("iorm_config_cache", nil)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.LastMaintenanceRunId != nil {
		s.D.Set("last_maintenance_run_id", *s.Res.LastMaintenanceRunId)
	}

	if s.Res.LastPatchHistoryEntryId != nil {
		s.D.Set("last_patch_history_entry_id", *s.Res.LastPatchHistoryEntryId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ListenerPort != nil {
		s.D.Set("listener_port", *s.Res.ListenerPort)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	if s.Res.NsgIds != nil {
		nsgIds := []interface{}{}
		for _, item := range s.Res.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))
	}

	if s.Res.OsVersion != nil {
		s.D.Set("os_version", *s.Res.OsVersion)
	}

	if s.Res.PointInTimeDataDiskCloneTimestamp != nil {
		s.D.Set("point_in_time_data_disk_clone_timestamp", s.Res.PointInTimeDataDiskCloneTimestamp.String())
	}

	if s.Res.RecoStorageSizeInGB != nil {
		s.D.Set("reco_storage_size_in_gb", *s.Res.RecoStorageSizeInGB)
	}

	if s.Res.ScanDnsName != nil {
		s.D.Set("scan_dns_name", *s.Res.ScanDnsName)
	}

	if s.Res.ScanDnsRecordId != nil {
		s.D.Set("scan_dns_record_id", *s.Res.ScanDnsRecordId)
	}

	s.D.Set("scan_ip_ids", s.Res.ScanIpIds)

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	if s.Res.SourceDbSystemId != nil {
		s.D.Set("source_db_system_id", *s.Res.SourceDbSystemId)
	}

	if s.Res.SparseDiskgroup != nil {
		s.D.Set("sparse_diskgroup", *s.Res.SparseDiskgroup)
	}

	sshPublicKeys := []interface{}{}
	for _, item := range s.Res.SshPublicKeys {
		sshPublicKeys = append(sshPublicKeys, item)
	}
	s.D.Set("ssh_public_keys", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, sshPublicKeys))

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("storage_volume_performance_mode", s.Res.StorageVolumePerformanceMode)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	s.D.Set("vip_ids", s.Res.VipIds)

	if s.Res.ZoneId != nil {
		s.D.Set("zone_id", *s.Res.ZoneId)
	}

	return nil
}

func (s *DatabaseDbSystemResourceCrud) mapToBackupDestinationDetails(fieldKeyFormat string) (oci_database.BackupDestinationDetails, error) {
	result := oci_database.BackupDestinationDetails{}

	if dbrsPolicyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dbrs_policy_id")); ok {
		tmp := dbrsPolicyId.(string)
		result.DbrsPolicyId = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_database.BackupDestinationDetailsTypeEnum(type_.(string))
	}

	return result, nil
}

func (s *DatabaseDbSystemResourceCrud) mapToCreateDatabaseDetails(fieldKeyFormat string) (oci_database.CreateDatabaseDetails, error) {
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

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

func CreateDatabaseDetailsToMap(obj *oci_database.CreateDatabaseDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminPassword != nil {
		result["admin_password"] = string(*obj.AdminPassword)
	}

	if obj.CharacterSet != nil {
		result["character_set"] = string(*obj.CharacterSet)
	}

	if obj.DatabaseSoftwareImageId != nil {
		result["database_software_image_id"] = string(*obj.DatabaseSoftwareImageId)
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

	result["db_workload"] = string(obj.DbWorkload)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}

	if obj.KmsKeyVersionId != nil {
		result["kms_key_version_id"] = string(*obj.KmsKeyVersionId)
	}

	if obj.NcharacterSet != nil {
		result["ncharacter_set"] = string(*obj.NcharacterSet)
	}

	if obj.PdbName != nil {
		result["pdb_name"] = string(*obj.PdbName)
	}

	if obj.TdeWalletPassword != nil {
		result["tde_wallet_password"] = string(*obj.TdeWalletPassword)
	}

	if obj.VaultId != nil {
		result["vault_id"] = string(*obj.VaultId)
	}

	return result
}

func (s *DatabaseDbSystemResourceCrud) mapToCreateDatabaseFromAnotherDatabaseDetails(fieldKeyFormat string) (oci_database.CreateDatabaseFromAnotherDatabaseDetails, error) {
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

	if timeStampForPointInTimeRecovery, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_stamp_for_point_in_time_recovery")); ok {
		tmp, err := time.Parse(time.RFC3339, timeStampForPointInTimeRecovery.(string))
		if err != nil {
			return result, err
		}
		result.TimeStampForPointInTimeRecovery = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func CreateDatabaseFromAnotherDatabaseDetailsToMap(obj *oci_database.CreateDatabaseFromAnotherDatabaseDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminPassword != nil {
		result["admin_password"] = string(*obj.AdminPassword)
	}

	if obj.BackupTDEPassword != nil {
		result["backup_tde_password"] = string(*obj.BackupTDEPassword)
	}

	if obj.DatabaseId != nil {
		result["database_id"] = string(*obj.DatabaseId)
	}

	if obj.DbName != nil {
		result["db_name"] = string(*obj.DbName)
	}

	if obj.DbUniqueName != nil {
		result["db_unique_name"] = string(*obj.DbUniqueName)
	}

	result["pluggable_databases"] = obj.PluggableDatabases

	if obj.TimeStampForPointInTimeRecovery != nil {
		result["time_stamp_for_point_in_time_recovery"] = obj.TimeStampForPointInTimeRecovery.Format(time.RFC3339Nano)
	}

	return result
}

func (s *DatabaseDbSystemResourceCrud) mapToCreateDatabaseFromBackupDetails(fieldKeyFormat string) (oci_database.CreateDatabaseFromBackupDetails, error) {
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

	return result, nil
}

func CreateDatabaseFromBackupDetailsToMap(obj *oci_database.CreateDatabaseFromBackupDetails) map[string]interface{} {
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

	if obj.DbUniqueName != nil {
		result["db_unique_name"] = string(*obj.DbUniqueName)
	}

	result["pluggable_databases"] = obj.PluggableDatabases

	return result
}

func (s *DatabaseDbSystemResourceCrud) mapToCreateDatabaseFromDbSystemDetails(fieldKeyFormat string) (oci_database.CreateDatabaseFromDbSystemDetails, error) {
	result := oci_database.CreateDatabaseFromDbSystemDetails{}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
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

	if dbDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_domain")); ok {
		tmp := dbDomain.(string)
		result.DbDomain = &tmp
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

	return result, nil
}

func CreateDatabaseFromDbSystemDetailsToMap(obj *oci_database.CreateDatabaseFromDbSystemDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminPassword != nil {
		result["admin_password"] = string(*obj.AdminPassword)
	}

	if obj.DbBackupConfig != nil {
		result["db_backup_config"] = []interface{}{DbBackupConfigToMap(obj.DbBackupConfig)}
	}

	if obj.DbDomain != nil {
		result["db_domain"] = string(*obj.DbDomain)
	}

	if obj.DbName != nil {
		result["db_name"] = string(*obj.DbName)
	}

	if obj.DbUniqueName != nil {
		result["db_unique_name"] = string(*obj.DbUniqueName)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	return result
}

func (s *DatabaseDbSystemResourceCrud) mapToCreateDbHomeDetails(fieldKeyFormat string) (oci_database.CreateDbHomeDetails, error) {
	result := oci_database.CreateDbHomeDetails{}

	if database, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database")); ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database"), 0)
			tmp, err := s.mapToCreateDatabaseDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert database, encountered error: %v", err)
			}
			result.Database = &tmp
		}
	}

	if databaseSoftwareImageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_software_image_id")); ok {
		tmp := databaseSoftwareImageId.(string)
		result.DatabaseSoftwareImageId = &tmp
	}

	if dbVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_version")); ok {
		tmp := dbVersion.(string)
		result.DbVersion = &tmp
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	return result, nil
}

func CreateDbHomeDetailsToMap(obj *oci_database.CreateDbHomeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Database != nil {
		result["database"] = []interface{}{CreateDatabaseDetailsToMap(obj.Database)}
	}

	if obj.DatabaseSoftwareImageId != nil {
		result["database_software_image_id"] = string(*obj.DatabaseSoftwareImageId)
	}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	return result
}

func (s *DatabaseDbSystemResourceCrud) mapToCreateDbHomeFromBackupDetails(fieldKeyFormat string) (oci_database.CreateDbHomeFromBackupDetails, error) {
	result := oci_database.CreateDbHomeFromBackupDetails{}

	if database, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database")); ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database"), 0)
			tmp, err := s.mapToCreateDatabaseFromBackupDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert database, encountered error: %v", err)
			}
			result.Database = &tmp
		}
	}

	if databaseSoftwareImageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_software_image_id")); ok {
		tmp := databaseSoftwareImageId.(string)
		result.DatabaseSoftwareImageId = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result, nil
}

func CreateDbHomeFromBackupDetailsToMap(obj *oci_database.CreateDbHomeFromBackupDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Database != nil {
		result["database"] = []interface{}{CreateDatabaseFromBackupDetailsToMap(obj.Database)}
	}

	if obj.DatabaseSoftwareImageId != nil {
		result["database_software_image_id"] = string(*obj.DatabaseSoftwareImageId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	return result
}

func (s *DatabaseDbSystemResourceCrud) mapToCreateDbHomeFromDbSystemDetails(fieldKeyFormat string) (oci_database.CreateDbHomeFromDbSystemDetails, error) {
	result := oci_database.CreateDbHomeFromDbSystemDetails{}

	if database, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database")); ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database"), 0)
			tmp, err := s.mapToCreateDatabaseFromDbSystemDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert database, encountered error: %v", err)
			}
			result.Database = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	return result, nil
}

func CreateDbHomeFromDbSystemDetailsToMap(obj *oci_database.CreateDbHomeFromDbSystemDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Database != nil {
		result["database"] = []interface{}{CreateDatabaseFromDbSystemDetailsToMap(obj.Database)}
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	return result
}

func (s *DatabaseDbSystemResourceCrud) mapToCreateDbHomeFromDatabaseDetails(fieldKeyFormat string) (oci_database.CreateDbHomeFromDatabaseDetails, error) {
	result := oci_database.CreateDbHomeFromDatabaseDetails{}

	if database, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database")); ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database"), 0)
			tmp, err := s.mapToCreateDatabaseFromAnotherDatabaseDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert database, encountered error: %v", err)
			}
			result.Database = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result, nil
}

func CreateDbHomeFromDatabaseDetailsToMap(obj *oci_database.CreateDbHomeFromDatabaseDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Database != nil {
		result["database"] = []interface{}{CreateDatabaseFromAnotherDatabaseDetailsToMap(obj.Database)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	return result
}

func (s *DatabaseDbSystemResourceCrud) mapToDataCollectionOptions(fieldKeyFormat string) (oci_database.DataCollectionOptions, error) {
	result := oci_database.DataCollectionOptions{}

	if isDiagnosticsEventsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_diagnostics_events_enabled")); ok {
		tmp := isDiagnosticsEventsEnabled.(bool)
		result.IsDiagnosticsEventsEnabled = &tmp
	}

	if isHealthMonitoringEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_health_monitoring_enabled")); ok {
		tmp := isHealthMonitoringEnabled.(bool)
		result.IsHealthMonitoringEnabled = &tmp
	}

	if isIncidentLogsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_incident_logs_enabled")); ok {
		tmp := isIncidentLogsEnabled.(bool)
		result.IsIncidentLogsEnabled = &tmp
	}

	return result, nil
}

func (s *DatabaseDbSystemResourceCrud) mapToDayOfWeek(fieldKeyFormat string) (oci_database.DayOfWeek, error) {
	result := oci_database.DayOfWeek{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.DayOfWeekNameEnum(name.(string))
	}

	return result, nil
}

func (s *DatabaseDbSystemResourceCrud) mapToDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
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

func DbBackupConfigToMap(obj *oci_database.DbBackupConfig) map[string]interface{} {
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
		backupDestinationDetails = append(backupDestinationDetails, BackupDestinationDetailsToMap(item))
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

func (s *DatabaseDbSystemResourceCrud) mapToDbSystemOptions(fieldKeyFormat string) (oci_database.DbSystemOptions, error) {
	result := oci_database.DbSystemOptions{}

	if storageManagement, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_management")); ok {
		result.StorageManagement = oci_database.DbSystemOptionsStorageManagementEnum(storageManagement.(string))
	}

	return result, nil
}

func DbSystemOptionsToMap(obj *oci_database.DbSystemOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["storage_management"] = string(obj.StorageManagement)

	return result
}

func ExadataIormConfigToMap(obj *oci_database.ExadataIormConfig) map[string]interface{} {
	result := map[string]interface{}{}

	dbPlans := []interface{}{}
	for _, item := range obj.DbPlans {
		dbPlans = append(dbPlans, dbIormConfigToMap(item))
	}
	result["db_plans"] = dbPlans

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["objective"] = string(obj.Objective)

	result["state"] = string(obj.LifecycleState)

	return result
}

func (s *DatabaseDbSystemResourceCrud) mapToMaintenanceWindow(fieldKeyFormat string) (oci_database.MaintenanceWindow, error) {
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
		result.LeadTimeInWeeks = &tmp
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

	if skipRu, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "skip_ru")); ok {
		interfaces := skipRu.([]interface{})
		tmp := make([]bool, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(bool)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "skip_ru")) {
			result.SkipRu = tmp
		}
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

func (s *DatabaseDbSystemResourceCrud) mapToMonth(fieldKeyFormat string) (oci_database.Month, error) {
	result := oci_database.Month{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.MonthNameEnum(name.(string))
	}

	return result, nil
}

func (s *DatabaseDbSystemResourceCrud) populateTopLevelPolymorphicLaunchDbSystemRequest(request *oci_database.LaunchDbSystemRequest) error {
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
		details := oci_database.LaunchDbSystemFromDatabaseDetails{}
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.LaunchDbSystemFromDatabaseDetailsDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbHome, ok := s.D.GetOkExists("db_home"); ok {
			if tmpList := dbHome.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_home", 0)
				tmp, err := s.mapToCreateDbHomeFromDatabaseDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DbHome = &tmp
			}
		}
		if diskRedundancy, ok := s.D.GetOkExists("disk_redundancy"); ok {
			details.DiskRedundancy = oci_database.LaunchDbSystemFromDatabaseDetailsDiskRedundancyEnum(diskRedundancy.(string))
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.LaunchDbSystemFromDatabaseDetailsLicenseModelEnum(licenseModel.(string))
		}
		if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
			set := backupNetworkNsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("backup_network_nsg_ids") {
				details.BackupNetworkNsgIds = tmp
			}
		}
		if backupSubnetId, ok := s.D.GetOkExists("backup_subnet_id"); ok {
			tmp := backupSubnetId.(string)
			details.BackupSubnetId = &tmp
		}
		if clusterName, ok := s.D.GetOkExists("cluster_name"); ok {
			tmp := clusterName.(string)
			details.ClusterName = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
		}
		if dataCollectionOptions, ok := s.D.GetOkExists("data_collection_options"); ok {
			if tmpList := dataCollectionOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_collection_options", 0)
				tmp, err := s.mapToDataCollectionOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataCollectionOptions = &tmp
			}
		}
		if dataStoragePercentage, ok := s.D.GetOkExists("data_storage_percentage"); ok {
			tmp := dataStoragePercentage.(int)
			details.DataStoragePercentage = &tmp
		}
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			details.InitialDataStorageSizeInGB = &tmp
		}
		if dbSystemOptions, ok := s.D.GetOkExists("db_system_options"); ok {
			if tmpList := dbSystemOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_system_options", 0)
				tmp, err := s.mapToDbSystemOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DbSystemOptions = &tmp
			}
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
		if domain, ok := s.D.GetOkExists("domain"); ok {
			tmp := domain.(string)
			details.Domain = &tmp
		}
		if faultDomains, ok := s.D.GetOkExists("fault_domains"); ok {
			interfaces := faultDomains.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("fault_domains") {
				details.FaultDomains = tmp
			}
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if hostname, ok := s.D.GetOkExists("hostname"); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.LaunchDbSystemFromDatabaseDetailsLicenseModelEnum(licenseModel.(string))
		}
		if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
			tmp := nodeCount.(int)
			details.NodeCount = &tmp
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
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if shape, ok := s.D.GetOkExists("shape"); ok {
			tmp := shape.(string)
			details.Shape = &tmp
		}
		if sparseDiskgroup, ok := s.D.GetOkExists("sparse_diskgroup"); ok {
			tmp := sparseDiskgroup.(bool)
			details.SparseDiskgroup = &tmp
		}
		if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
			set := sshPublicKeys.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("ssh_public_keys") {
				details.SshPublicKeys = tmp
			}
		}
		if storageVolumePerformanceMode, ok := s.D.GetOkExists("storage_volume_performance_mode"); ok {
			details.StorageVolumePerformanceMode = oci_database.LaunchDbSystemBaseStorageVolumePerformanceModeEnum(storageVolumePerformanceMode.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
			tmp := timeZone.(string)
			details.TimeZone = &tmp
		}
		request.LaunchDbSystemDetails = details
	case strings.ToLower("DB_BACKUP"):
		details := oci_database.LaunchDbSystemFromBackupDetails{}
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.LaunchDbSystemFromBackupDetailsDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbHome, ok := s.D.GetOkExists("db_home"); ok {
			if tmpList := dbHome.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_home", 0)
				tmp, err := s.mapToCreateDbHomeFromBackupDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DbHome = &tmp
			}
		}
		if diskRedundancy, ok := s.D.GetOkExists("disk_redundancy"); ok {
			details.DiskRedundancy = oci_database.LaunchDbSystemFromBackupDetailsDiskRedundancyEnum(diskRedundancy.(string))
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.LaunchDbSystemFromBackupDetailsLicenseModelEnum(licenseModel.(string))
		}
		if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
			set := backupNetworkNsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("backup_network_nsg_ids") {
				details.BackupNetworkNsgIds = tmp
			}
		}
		if backupSubnetId, ok := s.D.GetOkExists("backup_subnet_id"); ok {
			tmp := backupSubnetId.(string)
			details.BackupSubnetId = &tmp
		}
		if clusterName, ok := s.D.GetOkExists("cluster_name"); ok {
			tmp := clusterName.(string)
			details.ClusterName = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
		}
		if dataCollectionOptions, ok := s.D.GetOkExists("data_collection_options"); ok {
			if tmpList := dataCollectionOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_collection_options", 0)
				tmp, err := s.mapToDataCollectionOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataCollectionOptions = &tmp
			}
		}
		if dataStoragePercentage, ok := s.D.GetOkExists("data_storage_percentage"); ok {
			tmp := dataStoragePercentage.(int)
			details.DataStoragePercentage = &tmp
		}
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			details.InitialDataStorageSizeInGB = &tmp
		}
		if dbSystemOptions, ok := s.D.GetOkExists("db_system_options"); ok {
			if tmpList := dbSystemOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_system_options", 0)
				tmp, err := s.mapToDbSystemOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DbSystemOptions = &tmp
			}
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
		if domain, ok := s.D.GetOkExists("domain"); ok {
			tmp := domain.(string)
			details.Domain = &tmp
		}
		if faultDomains, ok := s.D.GetOkExists("fault_domains"); ok {
			interfaces := faultDomains.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("fault_domains") {
				details.FaultDomains = tmp
			}
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if hostname, ok := s.D.GetOkExists("hostname"); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.LaunchDbSystemFromBackupDetailsLicenseModelEnum(licenseModel.(string))
		}
		if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
			tmp := nodeCount.(int)
			details.NodeCount = &tmp
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
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if shape, ok := s.D.GetOkExists("shape"); ok {
			tmp := shape.(string)
			details.Shape = &tmp
		}
		if sparseDiskgroup, ok := s.D.GetOkExists("sparse_diskgroup"); ok {
			tmp := sparseDiskgroup.(bool)
			details.SparseDiskgroup = &tmp
		}
		if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
			set := sshPublicKeys.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("ssh_public_keys") {
				details.SshPublicKeys = tmp
			}
		}
		if storageVolumePerformanceMode, ok := s.D.GetOkExists("storage_volume_performance_mode"); ok {
			details.StorageVolumePerformanceMode = oci_database.LaunchDbSystemBaseStorageVolumePerformanceModeEnum(storageVolumePerformanceMode.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
			tmp := timeZone.(string)
			details.TimeZone = &tmp
		}
		request.LaunchDbSystemDetails = details
	case strings.ToLower("DB_SYSTEM"):
		details := oci_database.LaunchDbSystemFromDbSystemDetails{}
		if dbHome, ok := s.D.GetOkExists("db_home"); ok {
			if tmpList := dbHome.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_home", 0)
				tmp, err := s.mapToCreateDbHomeFromDbSystemDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DbHome = &tmp
			}
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.LaunchDbSystemFromDbSystemDetailsLicenseModelEnum(licenseModel.(string))
		}
		if sourceDbSystemId, ok := s.D.GetOkExists("source_db_system_id"); ok {
			tmp := sourceDbSystemId.(string)
			details.SourceDbSystemId = &tmp
		}
		if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
			set := backupNetworkNsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("backup_network_nsg_ids") {
				details.BackupNetworkNsgIds = tmp
			}
		}
		if backupSubnetId, ok := s.D.GetOkExists("backup_subnet_id"); ok {
			tmp := backupSubnetId.(string)
			details.BackupSubnetId = &tmp
		}
		if clusterName, ok := s.D.GetOkExists("cluster_name"); ok {
			tmp := clusterName.(string)
			details.ClusterName = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
		}
		if dataCollectionOptions, ok := s.D.GetOkExists("data_collection_options"); ok {
			if tmpList := dataCollectionOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_collection_options", 0)
				tmp, err := s.mapToDataCollectionOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataCollectionOptions = &tmp
			}
		}
		if dataStoragePercentage, ok := s.D.GetOkExists("data_storage_percentage"); ok {
			tmp := dataStoragePercentage.(int)
			details.DataStoragePercentage = &tmp
		}
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			details.InitialDataStorageSizeInGB = &tmp
		}
		if dbSystemOptions, ok := s.D.GetOkExists("db_system_options"); ok {
			if tmpList := dbSystemOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_system_options", 0)
				tmp, err := s.mapToDbSystemOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DbSystemOptions = &tmp
			}
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
		if domain, ok := s.D.GetOkExists("domain"); ok {
			tmp := domain.(string)
			details.Domain = &tmp
		}
		if faultDomains, ok := s.D.GetOkExists("fault_domains"); ok {
			interfaces := faultDomains.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("fault_domains") {
				details.FaultDomains = tmp
			}
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if hostname, ok := s.D.GetOkExists("hostname"); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.LaunchDbSystemFromDbSystemDetailsLicenseModelEnum(licenseModel.(string))
		}
		if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
			tmp := nodeCount.(int)
			details.NodeCount = &tmp
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
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if shape, ok := s.D.GetOkExists("shape"); ok {
			tmp := shape.(string)
			details.Shape = &tmp
		}
		if sparseDiskgroup, ok := s.D.GetOkExists("sparse_diskgroup"); ok {
			tmp := sparseDiskgroup.(bool)
			details.SparseDiskgroup = &tmp
		}
		if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
			set := sshPublicKeys.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("ssh_public_keys") {
				details.SshPublicKeys = tmp
			}
		}
		if storageVolumePerformanceMode, ok := s.D.GetOkExists("storage_volume_performance_mode"); ok {
			details.StorageVolumePerformanceMode = oci_database.LaunchDbSystemBaseStorageVolumePerformanceModeEnum(storageVolumePerformanceMode.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
			tmp := timeZone.(string)
			details.TimeZone = &tmp
		}
		request.LaunchDbSystemDetails = details
	case strings.ToLower("NONE"):
		details := oci_database.LaunchDbSystemDetails{}
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.LaunchDbSystemDetailsDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbHome, ok := s.D.GetOkExists("db_home"); ok {
			if tmpList := dbHome.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_home", 0)
				tmp, err := s.mapToCreateDbHomeDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DbHome = &tmp
			}
		}
		if diskRedundancy, ok := s.D.GetOkExists("disk_redundancy"); ok {
			details.DiskRedundancy = oci_database.LaunchDbSystemDetailsDiskRedundancyEnum(diskRedundancy.(string))
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.LaunchDbSystemDetailsLicenseModelEnum(licenseModel.(string))
		}
		if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
			set := backupNetworkNsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("backup_network_nsg_ids") {
				details.BackupNetworkNsgIds = tmp
			}
		}
		if backupSubnetId, ok := s.D.GetOkExists("backup_subnet_id"); ok {
			tmp := backupSubnetId.(string)
			details.BackupSubnetId = &tmp
		}
		if clusterName, ok := s.D.GetOkExists("cluster_name"); ok {
			tmp := clusterName.(string)
			details.ClusterName = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
		}
		if dataCollectionOptions, ok := s.D.GetOkExists("data_collection_options"); ok {
			if tmpList := dataCollectionOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_collection_options", 0)
				tmp, err := s.mapToDataCollectionOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataCollectionOptions = &tmp
			}
		}
		if dataStoragePercentage, ok := s.D.GetOkExists("data_storage_percentage"); ok {
			tmp := dataStoragePercentage.(int)
			details.DataStoragePercentage = &tmp
		}
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			details.InitialDataStorageSizeInGB = &tmp
		}
		if dbSystemOptions, ok := s.D.GetOkExists("db_system_options"); ok {
			if tmpList := dbSystemOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_system_options", 0)
				tmp, err := s.mapToDbSystemOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DbSystemOptions = &tmp
			}
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
		if domain, ok := s.D.GetOkExists("domain"); ok {
			tmp := domain.(string)
			details.Domain = &tmp
		}
		if faultDomains, ok := s.D.GetOkExists("fault_domains"); ok {
			interfaces := faultDomains.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("fault_domains") {
				details.FaultDomains = tmp
			}
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if hostname, ok := s.D.GetOkExists("hostname"); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.LaunchDbSystemDetailsLicenseModelEnum(licenseModel.(string))
		}
		if maintenanceWindowDetails, ok := s.D.GetOkExists("maintenance_window_details"); ok {
			if tmpList := maintenanceWindowDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window_details", 0)
				tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.MaintenanceWindowDetails = &tmp
			}
		}
		if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
			tmp := nodeCount.(int)
			details.NodeCount = &tmp
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
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if shape, ok := s.D.GetOkExists("shape"); ok {
			tmp := shape.(string)
			details.Shape = &tmp
		}
		if sparseDiskgroup, ok := s.D.GetOkExists("sparse_diskgroup"); ok {
			tmp := sparseDiskgroup.(bool)
			details.SparseDiskgroup = &tmp
		}
		if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
			set := sshPublicKeys.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("ssh_public_keys") {
				details.SshPublicKeys = tmp
			}
		}
		if storageVolumePerformanceMode, ok := s.D.GetOkExists("storage_volume_performance_mode"); ok {
			details.StorageVolumePerformanceMode = oci_database.LaunchDbSystemBaseStorageVolumePerformanceModeEnum(storageVolumePerformanceMode.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
			tmp := timeZone.(string)
			details.TimeZone = &tmp
		}
		request.LaunchDbSystemDetails = details
	default:
		return fmt.Errorf("unknown source '%v' was specified", source)
	}
	return nil
}

func (s *DatabaseDbSystemResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeDbSystemCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DbSystemId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeDbSystemCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func waitForDatabaseUpdateRetryPolicy(timeout time.Duration) *oci_common.RetryPolicy {
	startTime := time.Now()
	return &oci_common.RetryPolicy{
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			if tfresource.ShouldRetry(response, false, "database", startTime) {
				return true
			}
			if getDatabaseResponse, ok := response.Response.(oci_database.GetDatabaseResponse); ok {
				if getDatabaseResponse.LifecycleState == oci_database.DatabaseLifecycleStateUpdating {
					timeWaited := tfresource.GetElapsedRetryDuration(startTime)
					return timeWaited < timeout
				}
			}
			return false
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return tfresource.GetRetryBackoffDuration(response, false, "database", startTime)
		},
		MaximumNumberAttempts: 0,
	}
}

func (s *DatabaseDbSystemResourceCrud) mapToUpdateDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
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

	if runImmediateFullBackup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_immediate_full_backup")); ok {
		tmp := runImmediateFullBackup.(bool)
		result.RunImmediateFullBackup = &tmp
	}

	return result, nil
}

func (s *DatabaseDbSystemResourceCrud) getDbHomeInfo() error {
	if s.DbHome == nil {
		s.DbHome = &oci_database.DbHome{}
	}

	if s.Database == nil {
		s.Database = &oci_database.Database{}
	}

	var dbHomeId *string
	if s.DbHome.Id != nil {
		dbHomeId = s.DbHome.Id
	}
	if dbHomeId == nil || *dbHomeId == "" {
		dbHomeIdStr, ok := s.D.GetOkExists("db_home.0.id")
		if !ok || dbHomeIdStr == "" {
			listDbHomeRequest := oci_database.ListDbHomesRequest{}

			listDbHomeRequest.CompartmentId = s.Res.CompartmentId
			listDbHomeRequest.DbSystemId = s.Res.Id
			listDbHomeRequest.SortBy = oci_database.ListDbHomesSortByTimecreated
			listDbHomeRequest.SortOrder = oci_database.ListDbHomesSortOrderAsc
			listDbHomeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")
			listDbHomeResponse, err := s.Client.ListDbHomes(context.Background(), listDbHomeRequest)
			if err != nil {
				return err
			}
			if len(listDbHomeResponse.Items) == 0 {
				return fmt.Errorf("could not get details of the database for the dbHome")
			}
			if listDbHomeResponse.Items[0].TimeCreated.Sub(s.Res.TimeCreated.Time) > time.Hour*24 {
				return fmt.Errorf("The first database of the dbSystem has since been terminated. The details of the db_home will not be populated")
			}

			dbHomeId = listDbHomeResponse.Items[0].Id
		} else {
			tmp := dbHomeIdStr.(string)
			dbHomeId = &tmp
		}
	}
	getDbHomeRequest := oci_database.GetDbHomeRequest{}
	getDbHomeRequest.DbHomeId = dbHomeId
	getDbHomeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")
	getDbHomeResponse, err := s.Client.GetDbHome(context.Background(), getDbHomeRequest)
	if err != nil {
		return err
	}
	if getDbHomeResponse.DbHome.LifecycleState == oci_database.DbHomeLifecycleStateTerminated {
		return fmt.Errorf("the associated dbHome %s is in a TERMINATED state", *dbHomeId)
	}

	var databaseId *string
	if s.Database.Id != nil {
		databaseId = s.Database.Id
	}
	if databaseId == nil || *databaseId == "" {
		databaseIdStr, ok := s.D.GetOkExists("db_home.0.database.0.id")
		if !ok || databaseIdStr == "" {
			listDatabasesRequest := oci_database.ListDatabasesRequest{}

			listDatabasesRequest.CompartmentId = s.Res.CompartmentId
			listDatabasesRequest.DbHomeId = getDbHomeResponse.DbHome.Id
			listDatabasesRequest.SortBy = oci_database.ListDatabasesSortByTimecreated
			listDatabasesRequest.SortOrder = oci_database.ListDatabasesSortOrderAsc
			listDatabasesRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")
			listDatabasesResponse, err := s.Client.ListDatabases(context.Background(), listDatabasesRequest)
			if err != nil {
				return err
			}
			if len(listDatabasesResponse.Items) == 0 {
				return fmt.Errorf("could not get details of the database for the dbHome")
			}

			databaseId = listDatabasesResponse.Items[0].Id
		} else {
			tmp := databaseIdStr.(string)
			databaseId = &tmp
		}
	}

	getDatabaseRequest := oci_database.GetDatabaseRequest{}
	getDatabaseRequest.DatabaseId = databaseId
	getDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")
	getDatabaseResponse, err := s.Client.GetDatabase(context.Background(), getDatabaseRequest)
	if err != nil {
		return err
	}
	if getDatabaseResponse.Database.LifecycleState == oci_database.DatabaseLifecycleStateTerminated {
		return fmt.Errorf("the associated database is in a TERMINATED state")
	}
	if dbName, ok := s.D.GetOkExists("db_home.0.database.0.db_name"); ok {
		if getDatabaseResponse.Database.DbName != nil && dbName != *getDatabaseResponse.Database.DbName {
			return fmt.Errorf("the database name from the earliest database '%s' did not match the one on the config '%s'", *getDatabaseResponse.Database.DbName, dbName)
		}
	}

	s.DbHome = &getDbHomeResponse.DbHome
	s.Database = &getDatabaseResponse.Database

	return nil
}

func waitForDbSystemIfItIsUpdating(dbSystemID *string, client *oci_database.DatabaseClient, timeout time.Duration) (*oci_database.GetDbSystemResponse, error) {
	getDbSystemRequest := oci_database.GetDbSystemRequest{}

	getDbSystemRequest.DbSystemId = dbSystemID

	dbSystemUpdating := func(response oci_common.OCIOperationResponse) bool {
		if getDbSystemResponse, ok := response.Response.(oci_database.GetDbSystemResponse); ok {
			if getDbSystemResponse.LifecycleState == oci_database.DbSystemLifecycleStateUpdating {
				return true
			}
		}
		return false
	}

	getDbSystemRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicyWithAdditionalRetryCondition(timeout, dbSystemUpdating, "database")
	getDbSystemResponse, err := client.GetDbSystem(context.Background(), getDbSystemRequest)
	return &getDbSystemResponse, err
}

func (s *DatabaseDbSystemResourceCrud) UpdateDatabaseOperation() error {
	err := s.getDbHomeInfo()
	if err != nil {
		return err
	}

	updateDatabaseRequest := oci_database.UpdateDatabaseRequest{}

	updateDatabaseRequest.DatabaseId = s.Database.Id

	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_home.0.database", 0)
	updateDatabaseRequest.UpdateDatabaseDetails, err = s.mapToUpdateDatabaseDetails(fieldKeyFormat)
	if err != nil {
		return err
	}

	updateDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
	updateDatabaseResponse, err := s.Client.UpdateDatabase(context.Background(), updateDatabaseRequest)
	if err != nil {
		return err
	}

	workId := updateDatabaseResponse.OpcWorkRequestId
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

	errKms := s.setDbKeyVersion(s.Database.Id)
	if errKms != nil {
		return errKms
	}

	s.Database = &getDatabaseResponse.Database
	return nil
}

func (s *DatabaseDbSystemResourceCrud) setDbKeyVersion(databaseId *string) error {
	setDbKeyVersionRequest := oci_database.SetDbKeyVersionRequest{}
	setDbKeyVersionRequest.DatabaseId = databaseId
	setDbKeyVersionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
	details := oci_database.OciProviderSetKeyVersionDetails{}

	if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok && s.D.HasChange("kms_key_version_id") {
		oldRaw, newRaw := s.D.GetChange("kms_key_version_id")
		if oldRaw == "" && newRaw != "" {
			temp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &temp
			setDbKeyVersionRequest.SetKeyVersionDetails = details
		} else {
			return nil
		}
	}

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
	return nil
}

func (s *DatabaseDbSystemResourceCrud) sendUpdateForLicenseModel(dbSystemId string, licenseModel interface{}) error {
	request := oci_database.UpdateDbSystemRequest{}
	request.LicenseModel = oci_database.UpdateDbSystemDetailsLicenseModelEnum(licenseModel.(string))

	request.DbSystemId = &dbSystemId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateDbSystem(context.Background(), request)
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

	s.Res = &response.DbSystem

	getDbSystemResponse, err := waitForDbSystemIfItIsUpdating(s.Res.Id, s.Client, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		err = s.SetData()
		if err != nil {
			log.Printf("[ERROR] error setting data after polling error on the dbSystem: %v", err)
		}
		return fmt.Errorf("[ERROR] unable to get dbSystem after the Update: %v", err)
	}

	s.Res = &getDbSystemResponse.DbSystem

	return nil
}

func (s *DatabaseDbSystemResourceCrud) mapToUpdateDatabaseDetails(fieldKeyFormat string) (oci_database.UpdateDatabaseDetails, error) {
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
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

func (s *DatabaseDbSystemResourceCrud) DbHomeToMap(obj *oci_database.DbHome) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LastPatchHistoryEntryId != nil {
		result["last_patch_history_entry_id"] = string(*obj.LastPatchHistoryEntryId)
	}

	if obj.DbHomeLocation != nil {
		result["db_home_location"] = string(*obj.DbHomeLocation)
	}

	if obj.DatabaseSoftwareImageId != nil {
		result["database_software_image_id"] = string(*obj.DatabaseSoftwareImageId)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = obj.LifecycleState

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if s.Database != nil {
		result["database"] = []interface{}{s.DatabaseToMap(s.Database)}
	}

	return result
}

func (s *DatabaseDbSystemResourceCrud) DatabaseToMap(obj *oci_database.Database) map[string]interface{} {
	result := map[string]interface{}{}

	if adminPassword, ok := s.D.GetOkExists("db_home.0.database.0.admin_password"); ok && adminPassword != nil {
		result["admin_password"] = adminPassword.(string)
	}

	if kmsKeyId, ok := s.D.GetOkExists("db_home.0.database.0.kms_key_id"); ok && kmsKeyId != nil {
		result["kms_key_id"] = kmsKeyId.(string)
	}

	if kmsKeyVersionId, ok := s.D.GetOkExists("db_home.0.database.0.kms_key_version_id"); ok && kmsKeyVersionId != nil {
		result["kms_key_version_id"] = kmsKeyVersionId.(string)
	}

	if vaultId, ok := s.D.GetOkExists("db_home.0.database.0.vault_id"); ok && vaultId != nil {
		result["vault_id"] = vaultId.(string)
	}

	if tdeWalletPassword, ok := s.D.GetOkExists("db_home.0.database.0.tde_wallet_password"); ok && tdeWalletPassword != nil {
		result["tde_wallet_password"] = tdeWalletPassword.(string)
	}

	if backupId, ok := s.D.GetOkExists("db_home.0.database.0.backup_id"); ok && backupId != nil {
		result["backup_id"] = backupId.(string)
	}

	if backupTDEPassword, ok := s.D.GetOkExists("db_home.0.database.0.backup_tde_password"); ok && backupTDEPassword != nil {
		result["backup_tde_password"] = backupTDEPassword.(string)
	}

	if databaseId, ok := s.D.GetOkExists("db_home.0.database.0.database_id"); ok && databaseId != nil {
		result["database_id"] = databaseId.(string)
	}

	if obj.CharacterSet != nil {
		result["character_set"] = string(*obj.CharacterSet)
	}

	if obj.ConnectionStrings != nil {
		result["connection_strings"] = []interface{}{DatabaseConnectionStringsToMap(obj.ConnectionStrings)}
	}

	if obj.DatabaseSoftwareImageId != nil {
		result["database_software_image_id"] = string(*obj.DatabaseSoftwareImageId)
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

	if timeStampForPointInTimeRecovery, ok := s.D.GetOkExists(fmt.Sprintf("db_home.0.database.0.time_stamp_for_point_in_time_recovery")); ok {
		result["time_stamp_for_point_in_time_recovery"] = timeStampForPointInTimeRecovery
	}

	return result
}

func createDBSystemResource(d *schema.ResourceData, sync tfresource.ResourceCreator) error {
	if e := sync.Create(); e != nil {
		return tfresource.HandleError(sync, e)
	}

	// ID is required for state refresh
	d.SetId(sync.ID())

	var timeout time.Duration
	shape := d.Get("shape")
	timeout = d.Timeout(schema.TimeoutCreate)
	if timeout == 0 {
		if strings.HasPrefix(shape.(string), "Exadata") {
			timeout = tfresource.TwelveHours
		} else {
			timeout = tfresource.TwoHours
		}
	}
	if stateful, ok := sync.(tfresource.StatefullyCreatedResource); ok {
		if e := tfresource.WaitForStateRefresh(stateful, timeout, "creation", stateful.CreatedPending(), stateful.CreatedTarget()); e != nil {
			//We need to SetData() here because if there is an error or timeout in the wait for state after the Create() was successful we want to store the resource in the statefile to avoid dangling resources
			if setDataErr := sync.SetData(); setDataErr != nil {
				log.Printf("[ERROR] error setting data after waitForStateRefresh() error: %v", setDataErr)
			}
			return e
		}
	}

	d.SetId(sync.ID())
	if e := sync.SetData(); e != nil {
		return e
	}

	if ew, waitOK := sync.(tfresource.ExtraWaitPostCreateDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostCreateDelete())
	}

	return nil
}

func disableAutoBackupDbSystemSuppressfunc(k string, old, new string, d *schema.ResourceData) bool {
	// if autoBackupEnabled is false then ignore any field in the state and config backupWindow
	if autoBackupEnabled, ok := d.GetOkExists("db_home.0.database.0.db_backup_config.0.auto_backup_enabled"); ok {
		if !autoBackupEnabled.(bool) {
			return true
		}
	}
	return false
}
