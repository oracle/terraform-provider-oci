// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_database "github.com/oracle/oci-go-sdk/database"

	"fmt"
)

func DbSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			// crud.ZeroTime is a marker so a user supplied default is not overwritten. See crud.CreateDBSystemResource
			Create: &crud.ZeroTime,
			Delete: &crud.TwoHours,
			Update: &crud.TwoHours,
		},
		Create: createDbSystem,
		Read:   readDbSystem,
		Update: updateDbSystem,
		Delete: deleteDbSystem,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
									"db_name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
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

												// Computed
											},
										},
									},
									// this supports OLTP or DSS, returns "" if not supplied
									"db_workload": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
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
						"db_version": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
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
				DiffSuppressFunc: crud.EqualIgnoreCaseSuppressDiff,
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
			"source": {
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

func createDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	err := crud.CreateDBSystemResource(d, sync)
	if err != nil {
		return err
	}

	//Continue with the creation of subresources. The previous operation's result
	//should been persisted by now
	err = sync.createSubResources()
	if err != nil {
		return err
	}
	sync.SetData()

	return nil
}

func readDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

func updateDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.UpdateResource(d, sync)
}

func deleteDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type DbSystemResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.DbSystem
	DbHome                 *oci_database.DbHome
	Database               *oci_database.Database
	DisableNotFoundRetries bool
}

func (s *DbSystemResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DbSystemResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateProvisioning),
	}
}

func (s *DbSystemResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateAvailable),
	}
}

func (s *DbSystemResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateUpdating),
	}
}

func (s *DbSystemResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateAvailable),
	}
}

func (s *DbSystemResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateTerminating),
	}
}

func (s *DbSystemResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DbSystemLifecycleStateTerminated),
	}
}

func (s *DbSystemResourceCrud) Create() error {
	request := oci_database.LaunchDbSystemRequest{}
	s.populateTopLevelPolymorphicLaunchDbSystemRequest(&request)

	handleDbSimulationFlag(s.Client)

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.LaunchDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Client.Interceptor = nil

	s.Res = &response.DbSystem

	return nil
}

func (s *DbSystemResourceCrud) Get() error {
	tmp := s.D.Id()

	dbHomeID, _, _ := readDBHomeFromState(s.D)

	dbID, _, _ := readDatabaseFromState(s.D)

	request := oci_database.GetDbSystemRequest{}
	request.DbSystemId = &tmp
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")
	response, err := s.Client.GetDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem

	if dbHomeID != "" {
		dbHomeGet := oci_database.GetDbHomeRequest{DbHomeId: &dbHomeID}
		dbHomeGet.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")
		dbHomeRes, err := s.Client.GetDbHome(context.Background(), dbHomeGet)
		if err != nil {
			err = fmt.Errorf("when reading db homes in db system, db home was not avaialable due to: %s", err.Error())
			return err
		}
		s.DbHome = &dbHomeRes.DbHome
	}

	if dbID != "" {
		dbGet := oci_database.GetDatabaseRequest{DatabaseId: &dbID}
		dbGet.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")
		dbRes, err := s.Client.GetDatabase(context.Background(), dbGet)
		if err != nil {
			err = fmt.Errorf("when reading databases in db system, database was not avaialable due to: %s", err.Error())
			return err
		}
		s.Database = &dbRes.Database
	}

	return nil
}

func (s *DbSystemResourceCrud) updateDatabase() (database *oci_database.Database, err error) {
	dbID, databaseState, err := readDatabaseFromState(s.D)
	if err != nil {
		return
	}

	var backupConfig oci_database.DbBackupConfig
	if obj, ok := databaseState["db_backup_config"].([]interface{}); !ok {
		err = fmt.Errorf("db backup config state is not available. Can not update database in db system id: %s", s.D.Id())
		return
	} else {
		backupConfig = mapToDbBackupConfig(obj[0].(map[string]interface{}))
	}

	dbReq := oci_database.UpdateDatabaseRequest{}
	dbReq.DatabaseId = &dbID
	dbReq.UpdateDatabaseDetails.DbBackupConfig = &backupConfig
	dbReq.RequestMetadata.RetryPolicy = getDatabaseUpdateRetryPolicy(s.DisableNotFoundRetries)

	databaseUpdatesRes, err := s.Client.UpdateDatabase(context.Background(), dbReq)
	if err != nil {
		return
	}

	database = &databaseUpdatesRes.Database
	return
}

func (s *DbSystemResourceCrud) createSubResources() error {
	dbHomes, err := getDBHomesByDBSystem(s.Client, s.Res.Id, s.Res.CompartmentId, s.DisableNotFoundRetries)
	if err != nil || len(dbHomes) != 1 {
		if len(dbHomes) != 1 {
			err = fmt.Errorf("no db homes found in db system: %s", *s.Res.Id)
		}
		return err
	}

	// On creation there is only one db home
	s.DbHome = &dbHomes[0]

	//Get all databases for db homes
	s.Database, err = getDatabasesByDBHome(s.Client, s.DbHome.Id, s.Res.CompartmentId, s.DisableNotFoundRetries)
	return err
}

func (s *DbSystemResourceCrud) Update() error {
	request := oci_database.UpdateDbSystemRequest{}

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}

	if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
		tmp := dataStorageSizeInGB.(int)
		request.DataStorageSizeInGBs = &tmp
	}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.SshPublicKeys = []string{}
	if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
		interfaces := sshPublicKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.SshPublicKeys = tmp
	}

	//Update database first. Skip db home
	var err error
	s.Database, err = s.updateDatabase()
	if err != nil {
		return err
	}

	//save intermediate state
	s.SetData()

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem
	return nil
}

func (s *DbSystemResourceCrud) Delete() error {
	request := oci_database.TerminateDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.TerminateDbSystem(context.Background(), request)
	return err
}

func (s *DbSystemResourceCrud) SetData() {

	if s.DbHome != nil {
		_, dbHome, _ := readDBHomeFromState(s.D)
		_, database, _ := readDatabaseFromState(s.D)
		s.D.Set("db_home", []map[string]interface{}{dbHomeToMap(s.DbHome, s.Database, dbHome, database)})
	}

	if s.Res == nil {
		return
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
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

	if s.Res.DataStoragePercentage != nil {
		s.D.Set("data_storage_percentage", *s.Res.DataStoragePercentage)
	}

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	s.D.Set("database_edition", s.Res.DatabaseEdition)

	s.D.Set("disk_redundancy", s.Res.DiskRedundancy)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	// @codegen: Do not set hostname. Refreshing hostname causes undesirable diffs because the service may add a suffix
	// as in the case of Exadatas. Possible implication when importing the resource.

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
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

	s.D.Set("ssh_public_keys", s.Res.SshPublicKeys)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	s.D.Set("vip_ids", s.Res.VipIds)

}

func dbHomeToMap(obj *oci_database.DbHome, database *oci_database.Database, currentDBHome, currentDatabase map[string]interface{}) (res map[string]interface{}) {
	res = currentDBHome

	if obj.DbVersion != nil {
		res["db_version"] = *obj.DbVersion
	}

	if obj.DisplayName != nil {
		res["display_name"] = *obj.DisplayName
	}

	if obj.Id != nil {
		res["id"] = *obj.Id
	}

	if obj.LastPatchHistoryEntryId != nil {
		res["last_patch_history_entry_id"] = *obj.LastPatchHistoryEntryId
	}

	if string(obj.LifecycleState) != "" {
		res["state"] = string(obj.LifecycleState)
	}

	if obj.TimeCreated != nil {
		res["time_created"] = obj.TimeCreated.String()
	}

	if database != nil {
		res["database"] = []map[string]interface{}{databaseToMap(database, currentDatabase)}
	}
	return
}

func databaseToMap(database *oci_database.Database, currentDatabase map[string]interface{}) map[string]interface{} {
	result := currentDatabase

	if database.Id != nil {
		result["id"] = *database.Id
	}

	if database.LifecycleState != "" {
		result["state"] = string(database.LifecycleState)
	}

	if database.LifecycleDetails != nil {
		result["lifecycle_details"] = *database.LifecycleDetails
	}

	if database.CharacterSet != nil {
		result["character_set"] = string(*database.CharacterSet)
	}

	if database.DbBackupConfig != nil {
		objDbConfig := result["db_backup_config"].([]interface{})
		result["db_backup_config"] = []interface{}{DbBackupConfigToMap(database.DbBackupConfig, objDbConfig[0].(map[string]interface{}))}
	}

	if database.DbName != nil {
		result["db_name"] = *database.DbName
	}

	if database.DbUniqueName != nil {
		result["db_unique_name"] = *database.DbName
	}

	if database.DbWorkload != nil {
		result["db_workload"] = *database.DbWorkload
	}

	if database.NcharacterSet != nil {
		result["ncharacter_set"] = *database.NcharacterSet
	}

	if database.PdbName != nil {
		result["pdb_name"] = *database.PdbName
	}

	if database.TimeCreated != nil {
		result["time_created"] = database.TimeCreated.String()
	}

	return result

}

func mapToCreateDatabaseDetails(raw map[string]interface{}) oci_database.CreateDatabaseDetails {
	result := oci_database.CreateDatabaseDetails{}

	if adminPassword, ok := raw["admin_password"]; ok && adminPassword != "" {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
	}

	if characterSet, ok := raw["character_set"]; ok && characterSet != "" {
		tmp := characterSet.(string)
		result.CharacterSet = &tmp
	}

	if dbBackupConfig, ok := raw["db_backup_config"]; ok {
		if tmpList := dbBackupConfig.([]interface{}); len(tmpList) > 0 {
			tmp := mapToDbBackupConfig(tmpList[0].(map[string]interface{}))
			result.DbBackupConfig = &tmp
		}
	}

	if dbName, ok := raw["db_name"]; ok && dbName != "" {
		tmp := dbName.(string)
		result.DbName = &tmp
	}

	if dbWorkload, ok := raw["db_workload"]; ok && dbWorkload != "" {
		tmp := oci_database.CreateDatabaseDetailsDbWorkloadEnum(dbWorkload.(string))
		result.DbWorkload = tmp
	}

	if ncharacterSet, ok := raw["ncharacter_set"]; ok && ncharacterSet != "" {
		tmp := ncharacterSet.(string)
		result.NcharacterSet = &tmp
	}

	if pdbName, ok := raw["pdb_name"]; ok && pdbName != "" {
		tmp := pdbName.(string)
		result.PdbName = &tmp
	}

	return result
}

func mapToCreateDatabaseFromBackupDetails(raw map[string]interface{}) oci_database.CreateDatabaseFromBackupDetails {
	result := oci_database.CreateDatabaseFromBackupDetails{}

	if adminPassword, ok := raw["admin_password"]; ok && adminPassword != "" {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
	}

	if backupId, ok := raw["backup_id"]; ok && backupId != "" {
		tmp := backupId.(string)
		result.BackupId = &tmp
	}

	if backupTDEPassword, ok := raw["backup_tde_password"]; ok && backupTDEPassword != "" {
		tmp := backupTDEPassword.(string)
		result.BackupTDEPassword = &tmp
	}

	return result
}

func mapToCreateDbHomeDetails(raw map[string]interface{}) oci_database.CreateDbHomeDetails {
	result := oci_database.CreateDbHomeDetails{}

	if database, ok := raw["database"]; ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			tmp := mapToCreateDatabaseDetails(tmpList[0].(map[string]interface{}))
			result.Database = &tmp
		}
	}

	if dbVersion, ok := raw["db_version"]; ok && dbVersion != "" {
		tmp := dbVersion.(string)
		result.DbVersion = &tmp
	}

	if displayName, ok := raw["display_name"]; ok && displayName != "" {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result
}

func mapToCreateDbHomeFromBackupDetails(raw map[string]interface{}) oci_database.CreateDbHomeFromBackupDetails {
	result := oci_database.CreateDbHomeFromBackupDetails{}

	if database, ok := raw["database"]; ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			tmp := mapToCreateDatabaseFromBackupDetails(tmpList[0].(map[string]interface{}))
			result.Database = &tmp
		}
	}

	if displayName, ok := raw["display_name"]; ok && displayName != "" {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result
}

func DbBackupConfigToMap(obj *oci_database.DbBackupConfig, currDbConfig map[string]interface{}) map[string]interface{} {
	result := currDbConfig

	if obj.AutoBackupEnabled != nil {
		result["auto_backup_enabled"] = bool(*obj.AutoBackupEnabled)
	}

	return result
}

func (s *DbSystemResourceCrud) populateTopLevelPolymorphicLaunchDbSystemRequest(request *oci_database.LaunchDbSystemRequest) {
	//discriminator
	sourceRaw, ok := s.D.GetOkExists("source")
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "NONE" // default value
	}

	switch source {
	case "DB_BACKUP":
		details := oci_database.LaunchDbSystemFromBackupDetails{}
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.LaunchDbSystemFromBackupDetailsDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbHome, ok := s.D.GetOkExists("db_home"); ok {
			if tmpList := dbHome.([]interface{}); len(tmpList) > 0 {
				tmp := mapToCreateDbHomeFromBackupDetails(tmpList[0].(map[string]interface{}))
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
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if domain, ok := s.D.GetOkExists("domain"); ok {
			tmp := domain.(string)
			details.Domain = &tmp
		}
		if hostname, ok := s.D.GetOkExists("hostname"); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
			tmp := nodeCount.(int)
			details.NodeCount = &tmp
		}
		if shape, ok := s.D.GetOkExists("shape"); ok {
			tmp := shape.(string)
			details.Shape = &tmp
		}
		if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
			interfaces := sshPublicKeys.([]interface{})
			tmp := make([]string, len(interfaces))
			for i, toBeConverted := range interfaces {
				tmp[i] = toBeConverted.(string)
			}
			details.SshPublicKeys = tmp
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		request.LaunchDbSystemDetails = details

	case "NONE":
		details := oci_database.LaunchDbSystemDetails{}
		if databaseEdition, ok := s.D.GetOkExists("database_edition"); ok {
			details.DatabaseEdition = oci_database.LaunchDbSystemDetailsDatabaseEditionEnum(databaseEdition.(string))
		}
		if dbHome, ok := s.D.GetOkExists("db_home"); ok {
			if tmpList := dbHome.([]interface{}); len(tmpList) > 0 {
				tmp := mapToCreateDbHomeDetails(tmpList[0].(map[string]interface{}))
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
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if domain, ok := s.D.GetOkExists("domain"); ok {
			tmp := domain.(string)
			details.Domain = &tmp
		}
		if hostname, ok := s.D.GetOkExists("hostname"); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
			tmp := nodeCount.(int)
			details.NodeCount = &tmp
		}
		if shape, ok := s.D.GetOkExists("shape"); ok {
			tmp := shape.(string)
			details.Shape = &tmp
		}
		if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
			interfaces := sshPublicKeys.([]interface{})
			tmp := make([]string, len(interfaces))
			for i, toBeConverted := range interfaces {
				tmp[i] = toBeConverted.(string)
			}
			details.SshPublicKeys = tmp
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		request.LaunchDbSystemDetails = details
	default:
		log.Printf("[WARN] Unknown source '%v' was specified", source)
	}
}
