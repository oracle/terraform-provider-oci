// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseDbSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			// ZeroTime is a marker so a user supplied default is not overwritten. See CreateDBSystemResource
			Create: &ZeroTime,
			Delete: &TwoHours,
			Update: &TwoHours,
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
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// @CODEGEN cpu_core_count was made optional because the service ignores it when one provides a VM shape. This causes diffs after an apply
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"database_edition": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
									// server side defaults to AL32UTF8, but returns as "" when not supplied
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
									// this supports OLTP or DSS, returns "" if not supplied
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
									// serverside defaults to AL16UTF16, but returns as "" if not supplied
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
						"db_version": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: NewIsPrefixOfOldDiffSuppress,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"id": {
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
				},
			},
			"hostname": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: dbSystemHostnameDiffSuppress,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssh_public_keys": {
				Type:     schema.TypeList,
				Required: true,
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
				Set:      literalTypeHashCodeForSets,
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
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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
					DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"license_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_database.DbSystemLicenseModelLicenseIncluded),
					string(oci_database.DbSystemLicenseModelBringYourOwnLicense)}, false),
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
				Set:      literalTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
				}, true),
			},
			"sparse_diskgroup": {
				Type:     schema.TypeBool,
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
				MaxItems: 1,
				MinItems: 1,
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
			"reco_storage_size_in_gb": {
				Type:     schema.TypeInt,
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
		},
	}
}

func createDatabaseDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return CreateDBSystemResource(d, sync)
}

func readDatabaseDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

func updateDatabaseDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return UpdateResource(d, sync)
}

func deleteDatabaseDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseDbSystemResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
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
	return []string{
		string(oci_database.DbSystemLifecycleStateAvailable),
	}
}

func (s *DatabaseDbSystemResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateUpdating),
	}
}

func (s *DatabaseDbSystemResourceCrud) UpdatedTarget() []string {
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
	}
}

func (s *DatabaseDbSystemResourceCrud) Create() error {
	request := oci_database.LaunchDbSystemRequest{}
	err := s.populateTopLevelPolymorphicLaunchDbSystemRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.LaunchDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem

	err = s.getDbHomeInfo()
	if err != nil {
		log.Printf("[ERROR] Could not get info about the first DbHome in the dbSystem: %v", err)
	}
	return nil
}

func (s *DatabaseDbSystemResourceCrud) Get() error {
	request := oci_database.GetDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem

	err = s.getDbHomeInfo()
	if err != nil {
		log.Printf("[ERROR] Could not get info about the first DbHome in the dbSystem: %v", err)
	}

	return nil
}

func (s *DatabaseDbSystemResourceCrud) getDbHomeInfo() error {
	var dbHomeId *string
	if s.DbHome != nil && s.DbHome.Id != nil {
		dbHomeId = s.DbHome.Id
	}
	if dbHomeId == nil || *dbHomeId == "" {
		dbHomeIdStr, ok := s.D.GetOkExists("db_home.0.id")
		// if we don't have the DbHome Id in the config we get the earliest dbHome in the dbSystem from the service
		if !ok || dbHomeIdStr == "" {
			listDbHomeRequest := oci_database.ListDbHomesRequest{}

			listDbHomeRequest.CompartmentId = s.Res.CompartmentId
			listDbHomeRequest.DbSystemId = s.Res.Id
			listDbHomeRequest.SortBy = oci_database.ListDbHomesSortByTimecreated
			listDbHomeRequest.SortOrder = oci_database.ListDbHomesSortOrderAsc
			listDbHomeRequest.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")
			listDbHomeResponse, err := s.Client.ListDbHomes(context.Background(), listDbHomeRequest)
			if err != nil {
				return err
			}
			// if the dbSystem has no dbHomes return an error
			if len(listDbHomeResponse.Items) == 0 {
				return fmt.Errorf("could not get details of the database for the dbHome")
			}
			// If the time between the TimeCreated of the dbSystem is more than a day apart from the TimeCreated of the dbHome then we are not able to get the earliest dbHome.
			// DbHomes in a TERMINATED state are still returned as part of the list result for a few days
			if listDbHomeResponse.Items[0].TimeCreated.Sub(s.Res.TimeCreated.Time) > time.Hour*24 {
				return fmt.Errorf("The first database of the dbSystem has since been terminated. The details of the db_home will not be populated")
			}

			dbHomeId = listDbHomeResponse.Items[0].Id
		} else {
			tmp := dbHomeIdStr.(string)
			dbHomeId = &tmp
		}
	}
	// We do a get even if we have already done a list because the dbHomeSummary in the list response is a subset of the dbHome in the get response
	getDbHomeRequest := oci_database.GetDbHomeRequest{}
	getDbHomeRequest.DbHomeId = dbHomeId
	getDbHomeRequest.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")
	getDbHomeResponse, err := s.Client.GetDbHome(context.Background(), getDbHomeRequest)
	if err != nil {
		return err
	}
	if getDbHomeResponse.DbHome.LifecycleState == oci_database.DbHomeLifecycleStateTerminated {
		return fmt.Errorf("the associated dbHome %s is in a TERMINATED state", *dbHomeId)
	}

	var databaseId *string
	if s.Database != nil && s.Database.Id != nil {
		databaseId = s.Database.Id
	}
	if databaseId == nil || *databaseId == "" {
		databaseIdStr, ok := s.D.GetOkExists("db_home.0.database.0.id")
		// if we don't have the Database Id in the config we get the earliest database in the dbHome from the service
		if !ok || databaseIdStr == "" {
			listDatabasesRequest := oci_database.ListDatabasesRequest{}

			listDatabasesRequest.CompartmentId = s.Res.CompartmentId
			listDatabasesRequest.DbHomeId = getDbHomeResponse.DbHome.Id
			listDatabasesRequest.SortBy = oci_database.ListDatabasesSortByTimecreated
			listDatabasesRequest.SortOrder = oci_database.ListDatabasesSortOrderAsc
			listDatabasesRequest.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")
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

	// We do a get even if we have already done a list because the databaseSummary in the list response is a subset of the Database in the get response
	getDatabaseRequest := oci_database.GetDatabaseRequest{}
	getDatabaseRequest.DatabaseId = databaseId
	getDatabaseRequest.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")
	getDatabaseResponse, err := s.Client.GetDatabase(context.Background(), getDatabaseRequest)
	if err != nil {
		return err
	}
	if getDatabaseResponse.Database.LifecycleState == oci_database.DatabaseLifecycleStateTerminated {
		return fmt.Errorf("the associated database is in a TERMINATED state")
	}
	//if the dbName is not the same as what the user has in the config then we don't have the database that the user is trying to update
	if dbName, ok := s.D.GetOkExists("db_home.0.database.0.db_name"); ok {
		if getDatabaseResponse.Database.DbName != nil && dbName != *getDatabaseResponse.Database.DbName {
			return fmt.Errorf("the database name from the earliest database '%s' did not match the one on the config '%s'", *getDatabaseResponse.Database.DbName, dbName)
		}
	}

	s.DbHome = &getDbHomeResponse.DbHome
	s.Database = &getDatabaseResponse.Database

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

	request.BackupNetworkNsgIds = []string{}
	if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
		set := backupNetworkNsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.BackupNetworkNsgIds = tmp
	}

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}
	if s.D.HasChange("data_storage_size_in_gb") {
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			request.DataStorageSizeInGBs = &tmp
		}
	}
	tmp := s.D.Id()
	request.DbSystemId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.NsgIds = []string{}
	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.NsgIds = tmp
	}

	request.SshPublicKeys = []string{}
	if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
		interfaces := sshPublicKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.SshPublicKeys = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem

	// Wait for dbSystem to not be in updating state after the update. UpdateDatabase returns 409 if the dbSystem is in Updating state
	// We cannot use the usual waitForState logic here because a Get() before the SetData() would interfere with the subsequent Database Update
	getDbSystemResponse, err := waitForDbSystemIfItIsUpdating(s.Res.Id, s.Client, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		// Do SetData here in case the service returns updated values immediately on the Update request that don't need to wait for the waitForState
		err = s.SetData()
		if err != nil {
			log.Printf("[ERROR] error setting data after polling error on the dbSystem: %v", err)
		}
		return fmt.Errorf("[ERROR] unable to get dbSystem after the update: %v", err)
	}

	s.Res = &getDbSystemResponse.DbSystem

	err = s.SetData()
	if err != nil {
		return fmt.Errorf("[ERROR] error setting data after dbsystem update but before database update: %v", err)
	}

	return s.UpdateDatabaseOperation()
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

	getDbSystemRequest.RequestMetadata.RetryPolicy = getRetryPolicyWithAdditionalRetryCondition(timeout, dbSystemUpdating, "database")
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

	updateDatabaseRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")
	updateDatabaseResponse, err := s.Client.UpdateDatabase(context.Background(), updateDatabaseRequest)
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

	return nil
}

func (s *DatabaseDbSystemResourceCrud) Delete() error {
	request := oci_database.TerminateDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.TerminateDbSystem(context.Background(), request)
	return err
}

func (s *DatabaseDbSystemResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	backupNetworkNsgIds := []interface{}{}
	for _, item := range s.Res.BackupNetworkNsgIds {
		backupNetworkNsgIds = append(backupNetworkNsgIds, item)
	}
	s.D.Set("backup_network_nsg_ids", schema.NewSet(literalTypeHashCodeForSets, backupNetworkNsgIds))

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

	if s.Res.DataStoragePercentage != nil {
		s.D.Set("data_storage_percentage", *s.Res.DataStoragePercentage)
	}

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	s.D.Set("database_edition", s.Res.DatabaseEdition)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(literalTypeHashCodeForSets, nsgIds))

	if s.Res.RecoStorageSizeInGB != nil {
		s.D.Set("reco_storage_size_in_gb", *s.Res.RecoStorageSizeInGB)
	}

	if s.Res.ScanDnsRecordId != nil {
		s.D.Set("scan_dns_record_id", *s.Res.ScanDnsRecordId)
	}

	s.D.Set("scan_ip_ids", s.Res.ScanIpIds)

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	if s.Res.SparseDiskgroup != nil {
		s.D.Set("sparse_diskgroup", *s.Res.SparseDiskgroup)
	}

	s.D.Set("ssh_public_keys", s.Res.SshPublicKeys)

	s.D.Set("state", s.Res.LifecycleState)

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

	if s.Res.IormConfigCache != nil {
		s.D.Set("iorm_config_cache", []interface{}{IormConfigCacheToMap(s.Res.IormConfigCache)})
	} else {
		s.D.Set("iorm_config_cache", []interface{}{})
	}

	if s.DbHome != nil {
		s.D.Set("db_home", []interface{}{s.DbHomeToMap(s.DbHome)})
	}

	if source, ok := s.D.GetOkExists("source"); !ok || source.(string) == "" {
		s.D.Set("source", "NONE")
	}

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

func (s *DatabaseDbSystemResourceCrud) DbHomeToMap(obj *oci_database.DbHome) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LastPatchHistoryEntryId != nil {
		result["last_patch_history_entry_id"] = string(*obj.LastPatchHistoryEntryId)
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

	//Create parameters that are not returned by the service
	if adminPassword, ok := s.D.GetOkExists("db_home.0.database.0.admin_password"); ok && adminPassword != nil {
		result["admin_password"] = adminPassword.(string)
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

func IormConfigCacheToMap(obj *oci_database.ExadataIormConfig) map[string]interface{} {
	result := map[string]interface{}{}

	dbPlans := []interface{}{}
	for _, item := range obj.DbPlans {
		if configMap := dbIormConfigToMap(item); configMap != nil {
			dbPlans = append(dbPlans, configMap)
		}
	}
	result["db_plans"] = dbPlans

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = *obj.LifecycleDetails
	}

	result["objective"] = obj.Objective

	result["state"] = obj.LifecycleState

	return result
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

func CreateDatabaseDetailsToMap(obj *oci_database.CreateDatabaseDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminPassword != nil {
		result["admin_password"] = string(*obj.AdminPassword)
	}

	if obj.CharacterSet != nil {
		result["character_set"] = string(*obj.CharacterSet)
	}

	if obj.DbBackupConfig != nil {
		result["db_backup_config"] = []interface{}{DbBackupConfigToMap(obj.DbBackupConfig)}
	}

	if obj.DbName != nil {
		result["db_name"] = string(*obj.DbName)
	}

	result["db_workload"] = string(obj.DbWorkload)

	if obj.NcharacterSet != nil {
		result["ncharacter_set"] = string(*obj.NcharacterSet)
	}

	if obj.PdbName != nil {
		result["pdb_name"] = string(*obj.PdbName)
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

	if dbVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_version")); ok {
		tmp := dbVersion.(string)
		result.DbVersion = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result, nil
}

func CreateDbHomeDetailsToMap(obj *oci_database.CreateDbHomeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Database != nil {
		result["database"] = []interface{}{CreateDatabaseDetailsToMap(obj.Database)}
	}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

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

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	return result
}

// We cannot use the same function we use in create because the HasChanged check needed for the update to succeed interferes with the Create functionality
func (s *DatabaseDbSystemResourceCrud) mapToUpdateDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
	result := oci_database.DbBackupConfig{}

	// Service does not allow to update auto_backup_enabled and recovery_window_in_days at the same time so we must have the HasChanged check
	if autoBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")) {
		tmp := autoBackupEnabled.(bool)
		result.AutoBackupEnabled = &tmp
	}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")) {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	return result, nil
}

func (s *DatabaseDbSystemResourceCrud) mapToDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
	result := oci_database.DbBackupConfig{}

	if autoBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")); ok {
		tmp := autoBackupEnabled.(bool)
		result.AutoBackupEnabled = &tmp
	}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	return result, nil
}

func DbBackupConfigToMap(obj *oci_database.DbBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AutoBackupEnabled != nil {
		result["auto_backup_enabled"] = bool(*obj.AutoBackupEnabled)
	}

	if obj.RecoveryWindowInDays != nil {
		result["recovery_window_in_days"] = int(*obj.RecoveryWindowInDays)
	}

	return result
}

// @CODEGEN 08/2018: mapToPatchDetails and PatchDetailsToMap are not yet supported

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
		details.BackupNetworkNsgIds = []string{}
		if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
			set := backupNetworkNsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.BackupNetworkNsgIds = tmp
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
		if dataStoragePercentage, ok := s.D.GetOkExists("data_storage_percentage"); ok {
			tmp := dataStoragePercentage.(int)
			details.DataStoragePercentage = &tmp
		}
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			details.InitialDataStorageSizeInGB = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		details.FaultDomains = []string{}
		if faultDomains, ok := s.D.GetOkExists("fault_domains"); ok {
			interfaces := faultDomains.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.FaultDomains = tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if hostname, ok := s.D.GetOkExists("hostname"); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
			tmp := nodeCount.(int)
			details.NodeCount = &tmp
		}
		details.NsgIds = []string{}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.NsgIds = tmp
		}
		if shape, ok := s.D.GetOkExists("shape"); ok {
			tmp := shape.(string)
			details.Shape = &tmp
		}
		if sparseDiskgroup, ok := s.D.GetOkExists("sparse_diskgroup"); ok {
			tmp := sparseDiskgroup.(bool)
			details.SparseDiskgroup = &tmp
		}
		details.SshPublicKeys = []string{}
		if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
			interfaces := sshPublicKeys.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.SshPublicKeys = tmp
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
		details.BackupNetworkNsgIds = []string{}
		if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
			set := backupNetworkNsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.BackupNetworkNsgIds = tmp
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
		if dataStoragePercentage, ok := s.D.GetOkExists("data_storage_percentage"); ok {
			tmp := dataStoragePercentage.(int)
			details.DataStoragePercentage = &tmp
		}
		if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
			tmp := dataStorageSizeInGB.(int)
			details.InitialDataStorageSizeInGB = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		details.FaultDomains = []string{}
		if faultDomains, ok := s.D.GetOkExists("fault_domains"); ok {
			interfaces := faultDomains.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.FaultDomains = tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if hostname, ok := s.D.GetOkExists("hostname"); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
			tmp := nodeCount.(int)
			details.NodeCount = &tmp
		}
		details.NsgIds = []string{}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.NsgIds = tmp
		}
		if shape, ok := s.D.GetOkExists("shape"); ok {
			tmp := shape.(string)
			details.Shape = &tmp
		}
		if sparseDiskgroup, ok := s.D.GetOkExists("sparse_diskgroup"); ok {
			tmp := sparseDiskgroup.(bool)
			details.SparseDiskgroup = &tmp
		}
		details.SshPublicKeys = []string{}
		if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
			interfaces := sshPublicKeys.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.SshPublicKeys = tmp
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

func waitForDatabaseUpdateRetryPolicy(timeout time.Duration) *oci_common.RetryPolicy {
	startTime := time.Now()
	// wait for status of the database to not be UPDATING
	return &oci_common.RetryPolicy{
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			if shouldRetry(response, false, "database", startTime) {
				return true
			}
			if getDatabaseResponse, ok := response.Response.(oci_database.GetDatabaseResponse); ok {
				if getDatabaseResponse.LifecycleState == oci_database.DatabaseLifecycleStateUpdating {
					timeWaited := getElapsedRetryDuration(startTime)
					return timeWaited < timeout
				}
			}
			return false
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return getRetryBackoffDuration(response, false, "database", startTime)
		},
		MaximumNumberAttempts: 0,
	}
}

func (s *DatabaseDbSystemResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeDbSystemCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DbSystemId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeDbSystemCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
