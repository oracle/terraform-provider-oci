// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"
)

func MysqlMysqlDbSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createMysqlMysqlDbSystem,
		Read:   readMysqlMysqlDbSystem,
		Update: updateMysqlMysqlDbSystem,
		Delete: deleteMysqlMysqlDbSystem,
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
				ForceNew: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"admin_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"admin_username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"backup_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
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
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"pitr_policy": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"retention_in_days": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"window_start_time": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"configuration_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crash_recovery": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_storage": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_auto_expand_storage_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"max_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
						"allocated_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"data_storage_size_in_gb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"data_storage_size_limit_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"customer_contacts": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"email": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"data_storage_size_in_gb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"database_management": {
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
			"deletion_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"automatic_backup_retention": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"final_backup": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_delete_protected": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fault_domain": {
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
			"hostname_label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_highly_available": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"maintenance": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"window_start_time": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"mysql_version": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.MySqlVersionDiffSuppress,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port_x": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"secure_connections": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"certificate_generation_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"certificate_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"source": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"BACKUP",
								"IMPORTURL",
								"NONE",
								"PITR",
							}, true),
						},

						// Optional
						"backup_id": {
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
						"recovery_point": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"source_url": {
							Type:      schema.TypeString,
							Optional:  true,
							ForceNew:  true,
							Sensitive: true,
						},

						// Computed
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_mysql.DbSystemLifecycleStateInactive),
					string(oci_mysql.DbSystemLifecycleStateActive),
				}, true),
			},
			"shutdown_type": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_mysql.InnoDbShutdownModeFast),
					string(oci_mysql.InnoDbShutdownModeImmediate),
					string(oci_mysql.InnoDbShutdownModeSlow),
				}, true),
			},

			// Computed
			"channels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
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
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"anonymous_transactions_handling": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"last_configured_log_filename": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"last_configured_log_offset": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"policy": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"uuid": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"hostname": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"source_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssl_ca_certificate": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"certificate_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"contents": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"ssl_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"username": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"applier_username": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"channel_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"db_system_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"delay_in_seconds": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"filters": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"type": {
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
									"tables_without_primary_key_handling": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"current_placement": {
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
						"fault_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"modes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"port_x": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"heat_wave_cluster": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cluster_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_lakehouse_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"shape_name": {
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
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_heat_wave_cluster_attached": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"point_in_time_recovery_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"time_earliest_recovery_point": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_latest_recovery_point": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMysqlMysqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_mysql.DbSystemLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_mysql.DbSystemLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopMysqlDbInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.DbSystemLifecycleStateInactive)
	}

	return nil
}

func readMysqlMysqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	return tfresource.ReadResource(sync)
}

func updateMysqlMysqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	// switch to power on
	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_mysql.DbSystemLifecycleStateActive == oci_mysql.DbSystemLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_mysql.DbSystemLifecycleStateInactive == oci_mysql.DbSystemLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartMysqlDbInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.DbSystemLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	// switch to power off
	if powerOff {
		if err := sync.StopMysqlDbInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.DbSystemLifecycleStateInactive)
	}

	return nil
}

func deleteMysqlMysqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MysqlMysqlDbSystemResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_mysql.DbSystemClient
	Res                    *oci_mysql.DbSystem
	DisableNotFoundRetries bool
}

func (s *MysqlMysqlDbSystemResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MysqlMysqlDbSystemResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateCreating),
		string(oci_mysql.DbSystemLifecycleStateUpdating),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateActive),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateDeleting),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateDeleted),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateUpdating),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateActive),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) Create() error {
	request := oci_mysql.CreateDbSystemRequest{}

	if adminPassword, ok := s.D.GetOkExists("admin_password"); ok {
		tmp := adminPassword.(string)
		request.AdminPassword = &tmp
	}

	if adminUsername, ok := s.D.GetOkExists("admin_username"); ok {
		tmp := adminUsername.(string)
		request.AdminUsername = &tmp
	}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if backupPolicy, ok := s.D.GetOkExists("backup_policy"); ok {
		if tmpList := backupPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_policy", 0)
			tmp, err := s.mapToCreateBackupPolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupPolicy = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurationId, ok := s.D.GetOkExists("configuration_id"); ok {
		tmp := configurationId.(string)
		request.ConfigurationId = &tmp
	}

	if crashRecovery, ok := s.D.GetOkExists("crash_recovery"); ok {
		request.CrashRecovery = oci_mysql.CrashRecoveryStatusEnum(crashRecovery.(string))
	}

	if dataStorage, ok := s.D.GetOkExists("data_storage"); ok {
		if tmpList := dataStorage.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_storage", 0)
			tmp, err := s.mapToDataStorageDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataStorage = &tmp
		}
	}

	if customerContacts, ok := s.D.GetOkExists("customer_contacts"); ok {
		interfaces := customerContacts.([]interface{})
		tmp := make([]oci_mysql.CustomerContact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "customer_contacts", stateDataIndex)
			converted, err := s.mapToCustomerContact(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		// customer_contacts should only be set if:
		//   1) some customer contacts are provided (not `nil`)
		// and
		//   2) the customer contacts have changed
		if tmp != nil && s.D.HasChange("customer_contacts") {
			request.CustomerContacts = tmp
		}
	}

	if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
		tmp := dataStorageSizeInGB.(int)
		request.DataStorageSizeInGBs = &tmp
	}

	if databaseManagement, ok := s.D.GetOkExists("database_management"); ok {
		request.DatabaseManagement = oci_mysql.DatabaseManagementStatusEnum(databaseManagement.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if deletionPolicy, ok := s.D.GetOkExists("deletion_policy"); ok {
		if tmpList := deletionPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deletion_policy", 0)
			tmp, err := s.mapToCreateDeletionPolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DeletionPolicy = &tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if faultDomain, ok := s.D.GetOkExists("fault_domain"); ok {
		tmp := faultDomain.(string)
		request.FaultDomain = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists("hostname_label"); ok {
		tmp := hostnameLabel.(string)
		request.HostnameLabel = &tmp
	}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if isHighlyAvailable, ok := s.D.GetOkExists("is_highly_available"); ok {
		tmp := isHighlyAvailable.(bool)
		request.IsHighlyAvailable = &tmp
	}

	if maintenance, ok := s.D.GetOkExists("maintenance"); ok {
		if tmpList := maintenance.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance", 0)
			tmp, err := s.mapToCreateMaintenanceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Maintenance = &tmp
		}
	}

	if mysqlVersion, ok := s.D.GetOkExists("mysql_version"); ok {
		tmp := mysqlVersion.(string)
		request.MysqlVersion = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if portX, ok := s.D.GetOkExists("port_x"); ok {
		tmp := portX.(int)
		request.PortX = &tmp
	}

	if secureConnections, ok := s.D.GetOkExists("secure_connections"); ok {
		if tmpList := secureConnections.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "secure_connections", 0)
			tmp, err := s.mapToSecureConnectionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SecureConnections = &tmp
		}
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	if source, ok := s.D.GetOkExists("source"); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source", 0)
			tmp, err := s.mapToCreateDbSystemSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Source = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.CreateDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem
	return nil
}

func (s *MysqlMysqlDbSystemResourceCrud) Get() error {
	request := oci_mysql.GetDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem
	return nil
}

func (s *MysqlMysqlDbSystemResourceCrud) Update() error {
	request := oci_mysql.UpdateDbSystemRequest{}

	if backupPolicy, ok := s.D.GetOkExists("backup_policy"); ok && s.D.HasChange("backup_policy") {
		if tmpList := backupPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_policy", 0)
			tmp, err := s.mapToUpdateBackupPolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupPolicy = &tmp
		}
	}

	if crashRecovery, ok := s.D.GetOkExists("crash_recovery"); ok && s.D.HasChange("crash_recovery") {
		request.CrashRecovery = oci_mysql.CrashRecoveryStatusEnum(crashRecovery.(string))
	}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	if configurationId, ok := s.D.GetOkExists("configuration_id"); ok && s.D.HasChange("configuration_id") {
		tmp := configurationId.(string)
		request.ConfigurationId = &tmp
	}

	if dataStorage, ok := s.D.GetOkExists("data_storage"); ok && s.D.HasChange("data_storage") {
		if tmpList := dataStorage.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_storage", 0)
			tmp, err := s.mapToDataStorageDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataStorage = &tmp
		}
	}

	if customerContacts, ok := s.D.GetOkExists("customer_contacts"); ok {
		interfaces := customerContacts.([]interface{})
		tmp := make([]oci_mysql.CustomerContact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "customer_contacts", stateDataIndex)
			converted, err := s.mapToCustomerContact(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		// customer_contacts should only be set if:
		//   1) some customer contacts are provided (not `nil`)
		// and
		//   2) the customer contacts have changed
		if tmp != nil && s.D.HasChange("customer_contacts") {
			request.CustomerContacts = tmp
		}
	}

	if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok && s.D.HasChange("data_storage_size_in_gb") {
		tmp := dataStorageSizeInGB.(int)
		request.DataStorageSizeInGBs = &tmp
	}

	if databaseManagement, ok := s.D.GetOkExists("database_management"); ok && s.D.HasChange("database_management") {
		request.DatabaseManagement = oci_mysql.DatabaseManagementStatusEnum(databaseManagement.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if deletionPolicy, ok := s.D.GetOkExists("deletion_policy"); ok {
		if tmpList := deletionPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deletion_policy", 0)
			tmp, err := s.mapToUpdateDeletionPolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DeletionPolicy = &tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok && s.D.HasChange("description") {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok && s.D.HasChange("display_name") {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isHighlyAvailable, ok := s.D.GetOkExists("is_highly_available"); ok && s.D.HasChange("is_highly_available") {
		tmp := isHighlyAvailable.(bool)
		request.IsHighlyAvailable = &tmp
	}

	if maintenance, ok := s.D.GetOkExists("maintenance"); ok && s.D.HasChange("maintenance") {
		if tmpList := maintenance.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance", 0)
			tmp, err := s.mapToUpdateMaintenanceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Maintenance = &tmp
		}
	}

	if secureConnections, ok := s.D.GetOkExists("secure_connections"); ok && s.D.HasChange("secure_connections") {
		if tmpList := secureConnections.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "secure_connections", 0)
			tmp, err := s.mapToSecureConnectionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SecureConnections = &tmp
		}
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok && s.D.HasChange("shape_name") {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.UpdateDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *MysqlMysqlDbSystemResourceCrud) Delete() error {
	request := oci_mysql.DeleteDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.DeleteDbSystem(context.Background(), request)
	return err
}

func (s *MysqlMysqlDbSystemResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.BackupPolicy != nil {
		s.D.Set("backup_policy", []interface{}{BackupPolicyToMap(s.Res.BackupPolicy)})
	} else {
		s.D.Set("backup_policy", nil)
	}

	channels := []interface{}{}
	for _, item := range s.Res.Channels {
		channels = append(channels, ChannelSummaryToMap(item))
	}
	s.D.Set("channels", channels)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigurationId != nil {
		s.D.Set("configuration_id", *s.Res.ConfigurationId)
	}

	s.D.Set("crash_recovery", s.Res.CrashRecovery)

	if s.Res.CurrentPlacement != nil {
		s.D.Set("current_placement", []interface{}{DbSystemPlacementToMap(s.Res.CurrentPlacement)})
	} else {
		s.D.Set("current_placement", nil)
	}

	if s.Res.DataStorage != nil {
		s.D.Set("data_storage", []interface{}{DataStorageToMap(s.Res.DataStorage)})
	} else {
		s.D.Set("data_storage", nil)
	}

	customerContacts := []interface{}{}
	for _, item := range s.Res.CustomerContacts {
		customerContacts = append(customerContacts, CustomerContactToMap(item))
	}
	s.D.Set("customer_contacts", customerContacts)

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	s.D.Set("database_management", s.Res.DatabaseManagement)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeletionPolicy != nil {
		s.D.Set("deletion_policy", []interface{}{DeletionPolicyDetailsToMap(s.Res.DeletionPolicy)})
	} else {
		s.D.Set("deletion_policy", nil)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	endpoints := []interface{}{}
	for _, item := range s.Res.Endpoints {
		endpoints = append(endpoints, DbSystemEndpointToMap(item))
	}
	s.D.Set("endpoints", endpoints)

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HeatWaveCluster != nil {
		s.D.Set("heat_wave_cluster", []interface{}{HeatWaveClusterSummaryToMap(s.Res.HeatWaveCluster)})
	} else {
		s.D.Set("heat_wave_cluster", nil)
	}

	if s.Res.HostnameLabel != nil {
		s.D.Set("hostname_label", *s.Res.HostnameLabel)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.IsHeatWaveClusterAttached != nil {
		s.D.Set("is_heat_wave_cluster_attached", *s.Res.IsHeatWaveClusterAttached)
	}

	if s.Res.IsHighlyAvailable != nil {
		s.D.Set("is_highly_available", *s.Res.IsHighlyAvailable)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Maintenance != nil {
		s.D.Set("maintenance", []interface{}{MaintenanceDetailsToMap(s.Res.Maintenance)})
	} else {
		s.D.Set("maintenance", nil)
	}

	if s.Res.MysqlVersion != nil {
		s.D.Set("mysql_version", *s.Res.MysqlVersion)
	}

	if s.Res.PointInTimeRecoveryDetails != nil {
		s.D.Set("point_in_time_recovery_details", []interface{}{PointInTimeRecoveryDetailsToMap(s.Res.PointInTimeRecoveryDetails)})
	} else {
		s.D.Set("point_in_time_recovery_details", nil)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.PortX != nil {
		s.D.Set("port_x", *s.Res.PortX)
	}

	if s.Res.SecureConnections != nil {
		s.D.Set("secure_connections", []interface{}{SecureConnectionDetailsToMap(s.Res.SecureConnections)})
	} else {
		s.D.Set("secure_connections", nil)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	if s.Res.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := DbSystemSourceToMap(&s.Res.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		s.D.Set("source", sourceArray)
	} else {
		result := map[string]interface{}{}
		result["source_type"] = "NONE"
		sourceArray := []interface{}{}
		sourceArray = append(sourceArray, result)
		s.D.Set("source", sourceArray)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToAnonymousTransactionsHandling(fieldKeyFormat string) (oci_mysql.AnonymousTransactionsHandling, error) {
	var baseObject oci_mysql.AnonymousTransactionsHandling
	//discriminator
	policyRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "policy"))
	var policy string
	if ok {
		policy = policyRaw.(string)
	} else {
		policy = "" // default value
	}
	switch strings.ToLower(policy) {
	case strings.ToLower("ASSIGN_MANUAL_UUID"):
		details := oci_mysql.AssignManualUuidHandling{}
		if lastConfiguredLogFilename, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_configured_log_filename")); ok {
			tmp := lastConfiguredLogFilename.(string)
			details.LastConfiguredLogFilename = &tmp
		}
		if lastConfiguredLogOffset, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_configured_log_offset")); ok {
			tmp := lastConfiguredLogOffset.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert lastConfiguredLogOffset string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.LastConfiguredLogOffset = &tmpInt64
		}
		if uuid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "uuid")); ok {
			tmp := uuid.(string)
			details.Uuid = &tmp
		}
		baseObject = details
	case strings.ToLower("ASSIGN_TARGET_UUID"):
		details := oci_mysql.AssignTargetUuidHandling{}
		if lastConfiguredLogFilename, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_configured_log_filename")); ok {
			tmp := lastConfiguredLogFilename.(string)
			details.LastConfiguredLogFilename = &tmp
		}
		if lastConfiguredLogOffset, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_configured_log_offset")); ok {
			tmp := lastConfiguredLogOffset.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert lastConfiguredLogOffset string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.LastConfiguredLogOffset = &tmpInt64
		}
		baseObject = details
	case strings.ToLower("ERROR_ON_ANONYMOUS"):
		details := oci_mysql.ErrorOnAnonymousHandling{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown policy '%v' was specified", policy)
	}
	return baseObject, nil
}

func AnonymousTransactionsHandlingToMap(obj *oci_mysql.AnonymousTransactionsHandling) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_mysql.AssignManualUuidHandling:
		result["policy"] = "ASSIGN_MANUAL_UUID"

		if v.LastConfiguredLogFilename != nil {
			result["last_configured_log_filename"] = string(*v.LastConfiguredLogFilename)
		}

		if v.LastConfiguredLogOffset != nil {
			result["last_configured_log_offset"] = strconv.FormatInt(*v.LastConfiguredLogOffset, 10)
		}

		if v.Uuid != nil {
			result["uuid"] = string(*v.Uuid)
		}
	case oci_mysql.AssignTargetUuidHandling:
		result["policy"] = "ASSIGN_TARGET_UUID"

		if v.LastConfiguredLogFilename != nil {
			result["last_configured_log_filename"] = string(*v.LastConfiguredLogFilename)
		}

		if v.LastConfiguredLogOffset != nil {
			result["last_configured_log_offset"] = strconv.FormatInt(*v.LastConfiguredLogOffset, 10)
		}
	case oci_mysql.ErrorOnAnonymousHandling:
		result["policy"] = "ERROR_ON_ANONYMOUS"
	default:
		log.Printf("[WARN] Received 'policy' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToCaCertificate(fieldKeyFormat string) (oci_mysql.CaCertificate, error) {
	var baseObject oci_mysql.CaCertificate
	//discriminator
	certificateTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_type"))
	var certificateType string
	if ok {
		certificateType = certificateTypeRaw.(string)
	} else {
		certificateType = "" // default value
	}
	switch strings.ToLower(certificateType) {
	case strings.ToLower("PEM"):
		details := oci_mysql.PemCaCertificate{}
		if contents, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "contents")); ok {
			tmp := contents.(string)
			details.Contents = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown certificate_type '%v' was specified", certificateType)
	}
	return baseObject, nil
}

func CaCertificateToMap(obj *oci_mysql.CaCertificate) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_mysql.PemCaCertificate:
		result["certificate_type"] = "PEM"

		if v.Contents != nil {
			result["contents"] = string(*v.Contents)
		}
	default:
		log.Printf("[WARN] Received 'certificate_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToChannelFilter(fieldKeyFormat string) (oci_mysql.ChannelFilter, error) {
	result := oci_mysql.ChannelFilter{}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_mysql.ChannelFilterTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func ChannelFilterToMap(obj oci_mysql.ChannelFilter) map[string]interface{} {
	result := map[string]interface{}{}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func ChannelSourceToMap(obj *oci_mysql.ChannelSource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_mysql.ChannelSourceMysql:
		result["source_type"] = "MYSQL"

		if v.AnonymousTransactionsHandling != nil {
			anonymousTransactionsHandlingArray := []interface{}{}
			if anonymousTransactionsHandlingMap := AnonymousTransactionsHandlingToMap(&v.AnonymousTransactionsHandling); anonymousTransactionsHandlingMap != nil {
				anonymousTransactionsHandlingArray = append(anonymousTransactionsHandlingArray, anonymousTransactionsHandlingMap)
			}
			result["anonymous_transactions_handling"] = anonymousTransactionsHandlingArray
		}

		if v.Hostname != nil {
			result["hostname"] = string(*v.Hostname)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		if v.SslCaCertificate != nil {
			sslCaCertificateArray := []interface{}{}
			if sslCaCertificateMap := CaCertificateToMap(&v.SslCaCertificate); sslCaCertificateMap != nil {
				sslCaCertificateArray = append(sslCaCertificateArray, sslCaCertificateMap)
			}
			result["ssl_ca_certificate"] = sslCaCertificateArray
		}

		result["ssl_mode"] = string(v.SslMode)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func ChannelSummaryToMap(obj oci_mysql.ChannelSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := ChannelSourceToMap(&obj.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		result["source"] = sourceArray
	}

	result["state"] = string(obj.LifecycleState)

	if obj.Target != nil {
		targetArray := []interface{}{}
		if targetMap := ChannelTargetToMap(&obj.Target); targetMap != nil {
			targetArray = append(targetArray, targetMap)
		}
		result["target"] = targetArray
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func ChannelTargetToMap(obj *oci_mysql.ChannelTarget) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_mysql.ChannelTargetDbSystem:
		result["target_type"] = "DBSYSTEM"

		if v.ApplierUsername != nil {
			result["applier_username"] = string(*v.ApplierUsername)
		}

		if v.ChannelName != nil {
			result["channel_name"] = string(*v.ChannelName)
		}

		if v.DbSystemId != nil {
			result["db_system_id"] = string(*v.DbSystemId)
		}

		if v.DelayInSeconds != nil {
			result["delay_in_seconds"] = int(*v.DelayInSeconds)
		}

		filters := []interface{}{}
		for _, item := range v.Filters {
			filters = append(filters, ChannelFilterToMap(item))
		}
		result["filters"] = filters

		result["tables_without_primary_key_handling"] = string(v.TablesWithoutPrimaryKeyHandling)
	default:
		log.Printf("[WARN] Received 'target_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToCreateBackupPolicyDetails(fieldKeyFormat string) (oci_mysql.CreateBackupPolicyDetails, error) {
	result := oci_mysql.CreateBackupPolicyDetails{}

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

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if pitrPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pitr_policy")); ok {
		if tmpList := pitrPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "pitr_policy"), 0)
			tmp, err := s.mapToPitrPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert pitr_policy, encountered error: %v", err)
			}
			result.PitrPolicy = &tmp
		}
	}

	if retentionInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_in_days")); ok {
		tmp := retentionInDays.(int)
		result.RetentionInDays = &tmp
	}

	if windowStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "window_start_time")); ok {
		tmp := windowStartTime.(string)
		result.WindowStartTime = &tmp
	}

	return result, nil
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToCreateDbSystemSourceDetails(fieldKeyFormat string) (oci_mysql.CreateDbSystemSourceDetails, error) {
	var baseObject oci_mysql.CreateDbSystemSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "NONE" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("BACKUP"):
		details := oci_mysql.CreateDbSystemSourceFromBackupDetails{}
		if backupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_id")); ok {
			tmp := backupId.(string)
			details.BackupId = &tmp
		}
		baseObject = details
	case strings.ToLower("IMPORTURL"):
		details := oci_mysql.CreateDbSystemSourceImportFromUrlDetails{}
		if sourceUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_url")); ok {
			tmp := sourceUrl.(string)
			details.SourceUrl = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_mysql.CreateDbSystemSourceFromNoneDetails{}
		baseObject = details
	case strings.ToLower("PITR"):
		details := oci_mysql.CreateDbSystemSourceFromPitrDetails{}
		if dbSystemId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_system_id")); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if recoveryPoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_point")); ok {
			tmp, err := time.Parse(time.RFC3339, recoveryPoint.(string))
			if err != nil {
				return details, err
			}
			details.RecoveryPoint = &oci_common.SDKTime{Time: tmp}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func DbSystemSourceToMap(obj *oci_mysql.DbSystemSource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_mysql.DbSystemSourceFromBackup:
		result["source_type"] = "BACKUP"

		if v.BackupId != nil {
			result["backup_id"] = string(*v.BackupId)
		}
	case oci_mysql.DbSystemSourceImportFromUrl:
		result["source_type"] = "IMPORTURL"
	case oci_mysql.DbSystemSourceFromNone:
		result["source_type"] = "NONE"
	case oci_mysql.DbSystemSourceFromPitr:
		result["source_type"] = "PITR"

		if v.DbSystemId != nil {
			result["db_system_id"] = string(*v.DbSystemId)
		}

		if v.RecoveryPoint != nil {
			result["recovery_point"] = v.RecoveryPoint.Format(time.RFC3339Nano)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToCreateDeletionPolicyDetails(fieldKeyFormat string) (oci_mysql.CreateDeletionPolicyDetails, error) {
	result := oci_mysql.CreateDeletionPolicyDetails{}

	if automaticBackupRetention, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "automatic_backup_retention")); ok {
		result.AutomaticBackupRetention = oci_mysql.CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum(automaticBackupRetention.(string))
	}

	if finalBackup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "final_backup")); ok {
		result.FinalBackup = oci_mysql.CreateDeletionPolicyDetailsFinalBackupEnum(finalBackup.(string))
	}

	if isDeleteProtected, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_delete_protected")); ok {
		tmp := isDeleteProtected.(bool)
		result.IsDeleteProtected = &tmp
	}

	return result, nil
}

func DeletionPolicyDetailsToMap(obj *oci_mysql.DeletionPolicyDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["automatic_backup_retention"] = string(obj.AutomaticBackupRetention)

	result["final_backup"] = string(obj.FinalBackup)

	if obj.IsDeleteProtected != nil {
		result["is_delete_protected"] = bool(*obj.IsDeleteProtected)
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToUpdateDeletionPolicyDetails(fieldKeyFormat string) (oci_mysql.UpdateDeletionPolicyDetails, error) {
	result := oci_mysql.UpdateDeletionPolicyDetails{}

	if automaticBackupRetention, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "automatic_backup_retention")); ok {
		result.AutomaticBackupRetention = oci_mysql.UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum(automaticBackupRetention.(string))
	}

	if finalBackup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "final_backup")); ok {
		result.FinalBackup = oci_mysql.UpdateDeletionPolicyDetailsFinalBackupEnum(finalBackup.(string))
	}

	if isDeleteProtected, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_delete_protected")); ok {
		tmp := isDeleteProtected.(bool)
		result.IsDeleteProtected = &tmp
	}

	return result, nil
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToCreateMaintenanceDetails(fieldKeyFormat string) (oci_mysql.CreateMaintenanceDetails, error) {
	result := oci_mysql.CreateMaintenanceDetails{}

	if windowStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "window_start_time")); ok {
		tmp := windowStartTime.(string)
		result.WindowStartTime = &tmp
	}

	return result, nil
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToDataStorageDetails(fieldKeyFormat string) (oci_mysql.DataStorageDetails, error) {
	result := oci_mysql.DataStorageDetails{}

	if isAutoExpandStorageEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_expand_storage_enabled")); ok {
		tmp := isAutoExpandStorageEnabled.(bool)
		result.IsAutoExpandStorageEnabled = &tmp
	}

	if maxStorageSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_storage_size_in_gbs")); ok {
		tmp := maxStorageSizeInGBs.(int)
		result.MaxStorageSizeInGBs = &tmp
	}

	return result, nil
}

func DataStorageToMap(obj *oci_mysql.DataStorage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllocatedStorageSizeInGBs != nil {
		result["allocated_storage_size_in_gbs"] = int(*obj.AllocatedStorageSizeInGBs)
	}

	if obj.DataStorageSizeInGBs != nil {
		result["data_storage_size_in_gb"] = int(*obj.DataStorageSizeInGBs)
	}

	if obj.DataStorageSizeLimitInGBs != nil {
		result["data_storage_size_limit_in_gbs"] = int(*obj.DataStorageSizeLimitInGBs)
	}

	if obj.IsAutoExpandStorageEnabled != nil {
		result["is_auto_expand_storage_enabled"] = bool(*obj.IsAutoExpandStorageEnabled)
	}

	if obj.MaxStorageSizeInGBs != nil {
		result["max_storage_size_in_gbs"] = int(*obj.MaxStorageSizeInGBs)
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToCustomerContact(fieldKeyFormat string) (oci_mysql.CustomerContact, error) {
	result := oci_mysql.CustomerContact{}

	if email, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email")); ok {
		tmp := email.(string)
		result.Email = &tmp
	}

	return result, nil
}

func CustomerContactToMap(obj oci_mysql.CustomerContact) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Email != nil {
		result["email"] = string(*obj.Email)
	}

	return result
}

func DbSystemPlacementToMap(obj *oci_mysql.DbSystemPlacement) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	return result
}

func HeatWaveClusterSummaryToMap(obj *oci_mysql.HeatWaveClusterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClusterSize != nil {
		result["cluster_size"] = int(*obj.ClusterSize)
	}

	if obj.IsLakehouseEnabled != nil {
		result["is_lakehouse_enabled"] = bool(*obj.IsLakehouseEnabled)
	}

	if obj.ShapeName != nil {
		result["shape_name"] = string(*obj.ShapeName)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToPitrPolicy(fieldKeyFormat string) (oci_mysql.PitrPolicy, error) {
	result := oci_mysql.PitrPolicy{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func PitrPolicyToMap(obj *oci_mysql.PitrPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func PointInTimeRecoveryDetailsToMap(obj *oci_mysql.PointInTimeRecoveryDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TimeEarliestRecoveryPoint != nil {
		result["time_earliest_recovery_point"] = obj.TimeEarliestRecoveryPoint.String()
	}

	if obj.TimeLatestRecoveryPoint != nil {
		result["time_latest_recovery_point"] = obj.TimeLatestRecoveryPoint.String()
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToSecureConnectionDetails(fieldKeyFormat string) (oci_mysql.SecureConnectionDetails, error) {
	result := oci_mysql.SecureConnectionDetails{}

	if certificateGenerationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_generation_type")); ok {
		result.CertificateGenerationType = oci_mysql.CertificateGenerationTypeEnum(certificateGenerationType.(string))
	}

	if certificateId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_id")); ok {
		tmp := certificateId.(string)
		result.CertificateId = &tmp
	}

	return result, nil
}

func SecureConnectionDetailsToMap(obj *oci_mysql.SecureConnectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["certificate_generation_type"] = string(obj.CertificateGenerationType)

	if obj.CertificateId != nil {
		result["certificate_id"] = string(*obj.CertificateId)
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToUpdateBackupPolicyDetails(fieldKeyFormat string) (oci_mysql.UpdateBackupPolicyDetails, error) {
	result := oci_mysql.UpdateBackupPolicyDetails{}

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

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if retentionInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_in_days")); ok {
		tmp := retentionInDays.(int)
		result.RetentionInDays = &tmp
	}

	if windowStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "window_start_time")); ok {
		tmp := windowStartTime.(string)
		result.WindowStartTime = &tmp
	}

	if pitrPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pitr_policy")); ok {
		if tmpList := pitrPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "pitr_policy"), 0)
			tmp, err := s.mapToPitrPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert pitr_policy, encountered error: %v", err)
			}
			result.PitrPolicy = &tmp
		}
	}

	return result, nil
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToUpdateMaintenanceDetails(fieldKeyFormat string) (oci_mysql.UpdateMaintenanceDetails, error) {
	result := oci_mysql.UpdateMaintenanceDetails{}

	if windowStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "window_start_time")); ok {
		tmp := windowStartTime.(string)
		result.WindowStartTime = &tmp
	}

	return result, nil
}

func (s *MysqlMysqlDbSystemResourceCrud) StartMysqlDbInstance() error {
	request := oci_mysql.StartDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StartDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.DbSystemLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *MysqlMysqlDbSystemResourceCrud) StopMysqlDbInstance() error {
	request := oci_mysql.StopDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	shutdownTypeRaw, ok := s.D.GetOkExists("shutdown_type")

	var shutdown_type string
	if ok {
		shutdown_type = shutdownTypeRaw.(string)
	} else {
		shutdown_type = "FAST"
	}

	switch strings.ToLower(shutdown_type) {
	case strings.ToLower("SLOW"):
		request.ShutdownType = oci_mysql.InnoDbShutdownModeSlow
	case strings.ToLower("FAST"):
		request.ShutdownType = oci_mysql.InnoDbShutdownModeFast
	case strings.ToLower("IMMEDIATE"):
		request.ShutdownType = oci_mysql.InnoDbShutdownModeImmediate
	default:
		return fmt.Errorf("unknown shutdown_type '%v' was specified", shutdown_type)
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StopDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.DbSystemLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}
