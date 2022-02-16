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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/v58/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v58/workrequests"
)

func DatabaseDatabaseUpgradeResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseDatabaseUpgrade,
		Read:     readDatabaseDatabaseUpgrade,
		Delete:   deleteDatabaseDatabaseUpgrade,
		Schema: map[string]*schema.Schema{
			// Required
			"action": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"database_upgrade_source_details": {
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
						"database_software_image_id": {
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
						"options": {
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
								"DB_SOFTWARE_IMAGE",
								"DB_VERSION",
							}, true),
						},

						// Computed
					},
				},
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
			"database_software_image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_backup_config": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
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
						"backup_destination_details": {
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
			"db_home_id": {
				Type:     schema.TypeString,
				Computed: true,
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

			"last_backup_timestamp": {
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseDatabaseUpgrade(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseUpgradeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseDatabaseUpgrade(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseUpgradeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func deleteDatabaseDatabaseUpgrade(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseDatabaseUpgradeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.Database
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseDatabaseUpgradeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDatabaseUpgradeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateUpgrading),
	}
}

func (s *DatabaseDatabaseUpgradeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseDatabaseUpgradeResourceCrud) Create() error {
	request := oci_database.UpgradeDatabaseRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		request.Action = oci_database.UpgradeDatabaseDetailsActionEnum(action.(string))
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if databaseUpgradeSourceDetails, ok := s.D.GetOkExists("database_upgrade_source_details"); ok {
		if tmpList := databaseUpgradeSourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database_upgrade_source_details", 0)
			tmp, err := s.mapToDatabaseUpgradeSourceBase(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatabaseUpgradeSourceDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpgradeDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Database

	workId := response.OpcWorkRequestId

	return s.getDatabaseUpgradeFromWorkRequest(workId, oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseDatabaseUpgradeResourceCrud) getDatabaseUpgradeFromWorkRequest(workId *string, actionTypeEnum oci_work_requests.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {
	databaseUpgradeId, err := tfresource.WaitForWorkRequest(s.WorkRequestClient, workId, "database", actionTypeEnum, timeout, s.DisableNotFoundRetries, true)
	log.Printf("[DEBUG] WaitForWorkRequest finished. databaseUpgradeId: %v err: %v for workId: %v, actionTypeEnum: %v\n", *databaseUpgradeId, err, *workId, actionTypeEnum)
	if err != nil {
		log.Printf("[ERROR] Database upgrade operation failed, attempting to cancel the workrequest: %v for identifier: %v\n", *workId, databaseUpgradeId)
		return err
	}

	s.D.SetId(*databaseUpgradeId)

	return s.Get()
}

func (s *DatabaseDatabaseUpgradeResourceCrud) SetData() error {
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

	if s.Res.LastBackupTimestamp != nil {
		s.D.Set("last_backup_timestamp", s.Res.LastBackupTimestamp.String())
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

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VmClusterId != nil {
		s.D.Set("vm_cluster_id", *s.Res.VmClusterId)
	}

	return nil
}

func (s *DatabaseDatabaseUpgradeResourceCrud) Get() error {
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

func (s *DatabaseDatabaseUpgradeResourceCrud) mapToDatabaseUpgradeSourceBase(fieldKeyFormat string) (oci_database.DatabaseUpgradeSourceBase, error) {
	var baseObject oci_database.DatabaseUpgradeSourceBase
	//discriminator
	sourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source"))
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("DB_SOFTWARE_IMAGE"):
		details := oci_database.DatabaseUpgradeWithDatabaseSoftwareImageDetails{}
		if databaseSoftwareImageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_software_image_id")); ok {
			tmp := databaseSoftwareImageId.(string)
			details.DatabaseSoftwareImageId = &tmp
		}
		if options, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "options")); ok {
			tmp := options.(string)
			details.Options = &tmp
		}
		baseObject = details
	case strings.ToLower("DB_VERSION"):
		details := oci_database.DatabaseUpgradeWithDbVersionDetails{}
		if dbVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_version")); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if options, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "options")); ok {
			tmp := options.(string)
			details.Options = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source '%v' was specified", source)
	}
	return baseObject, nil
}

func DatabaseUpgradeSourceBaseToMap(obj *oci_database.DatabaseUpgradeSourceBase) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database.DatabaseUpgradeWithDatabaseSoftwareImageDetails:
		result["source"] = "DB_SOFTWARE_IMAGE"

		if v.DatabaseSoftwareImageId != nil {
			result["database_software_image_id"] = string(*v.DatabaseSoftwareImageId)
		}
	case oci_database.DatabaseUpgradeWithDbVersionDetails:
		result["source"] = "DB_VERSION"

		if v.DbVersion != nil {
			result["db_version"] = string(*v.DbVersion)
		}
	default:
		log.Printf("[WARN] Received 'source' of unknown type %v", *obj)
		return nil
	}

	return result
}
