// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_database "github.com/oracle/oci-go-sdk/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/workrequests"
)

func init() {
	RegisterResource("oci_database_autonomous_database", DatabaseAutonomousDatabaseResource())
}

func DatabaseAutonomousDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: getTimeoutDuration("12h"),
			Update: getTimeoutDuration("12h"),
			Delete: getTimeoutDuration("12h"),
		},
		Create: createDatabaseAutonomousDatabase,
		Read:   readDatabaseAutonomousDatabase,
		Update: updateDatabaseAutonomousDatabase,
		Delete: deleteDatabaseAutonomousDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
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
			"clone_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"data_safe_status": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_database.AutonomousDatabaseDataSafeStatusRegistered),
					string(oci_database.AutonomousDatabaseSummaryDataSafeStatusNotRegistered),
				}, true),
			},
			"db_version": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: dbVersionDiffSuppress,
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
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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
			"is_auto_scaling_enabled": {
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
			"is_preview_version_with_service_terms_accepted": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
				Set:      literalTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"private_endpoint_label": {
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
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"BACKUP_FROM_ID",
					"BACKUP_FROM_TIMESTAMP",
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
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"timestamp": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: timeDiffSuppressFunction,
			},
			"whitelisted_ips": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      literalTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"available_upgrade_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"is_preview": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_console_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
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
			"time_reclamation_of_free_autonomous_database": {
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
	sync.Client = m.(*OracleClients).databaseClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient

	configDataSafeStatus := oci_database.AutonomousDatabaseDataSafeStatusNotRegistered
	if dataSafeStatus, ok := sync.D.GetOkExists("data_safe_status"); ok {
		configDataSafeStatus = oci_database.AutonomousDatabaseDataSafeStatusEnum(strings.ToUpper(dataSafeStatus.(string)))
	}

	if e := CreateResource(d, sync); e != nil {
		return e
	}

	if configDataSafeStatus == oci_database.AutonomousDatabaseDataSafeStatusRegistered {
		err := sync.updateDataSafeStatus(sync.D.Id(), oci_database.AutonomousDatabaseDataSafeStatusRegistered)
		if err != nil {
			return err
		}
		return ReadResource(sync)
	}

	return nil
}

func readDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

func updateDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient

	return UpdateResource(d, sync)
}

func deleteDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseAutonomousDatabaseResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	workRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.AutonomousDatabase
	DisableNotFoundRetries bool
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabase
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

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

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok && s.D.HasChange("nsg_ids") {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			nsgUpdateRequest := oci_database.UpdateAutonomousDatabaseRequest{}

			autonomousDatabaseId := s.D.Id()
			nsgUpdateRequest.AutonomousDatabaseId = &autonomousDatabaseId

			nsgUpdateRequest.NsgIds = tmp

			nsgUpdateRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

			nsgUpdateResponse, err := s.Client.UpdateAutonomousDatabase(context.Background(), nsgUpdateRequest)
			if err != nil {
				return err
			}

			s.Res = &nsgUpdateResponse.AutonomousDatabase
		}
	}

	request := oci_database.UpdateAutonomousDatabaseRequest{}

	if adminPassword, ok := s.D.GetOkExists("admin_password"); ok && s.D.HasChange("admin_password") {
		tmp := adminPassword.(string)
		request.AdminPassword = &tmp
	}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok && s.D.HasChange("cpu_core_count") {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
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

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok && s.D.HasChange("is_auto_scaling_enabled") {
		tmp := isAutoScalingEnabled.(bool)
		request.IsAutoScalingEnabled = &tmp
	}

	if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok && s.D.HasChange("is_free_tier") {
		tmp := isFreeTier.(bool)
		request.IsFreeTier = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok && s.D.HasChange("license_model") {
		request.LicenseModel = oci_database.UpdateAutonomousDatabaseDetailsLicenseModelEnum(licenseModel.(string))
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabase
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) Delete() error {
	request := oci_database.DeleteAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteAutonomousDatabase(context.Background(), request)
	return err
}

func (s *DatabaseAutonomousDatabaseResourceCrud) SetData() error {
	if s.Res.AutonomousContainerDatabaseId != nil {
		s.D.Set("autonomous_container_database_id", *s.Res.AutonomousContainerDatabaseId)
	}

	s.D.Set("available_upgrade_versions", s.Res.AvailableUpgradeVersions)

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

	s.D.Set("data_safe_status", s.Res.DataSafeStatus)

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	s.D.Set("db_workload", s.Res.DbWorkload)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoScalingEnabled != nil {
		s.D.Set("is_auto_scaling_enabled", *s.Res.IsAutoScalingEnabled)
	}

	if s.Res.IsDedicated != nil {
		s.D.Set("is_dedicated", *s.Res.IsDedicated)
	}

	if s.Res.IsFreeTier != nil {
		s.D.Set("is_free_tier", *s.Res.IsFreeTier)
	}

	if s.Res.IsPreview != nil {
		s.D.Set("is_preview", *s.Res.IsPreview)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(literalTypeHashCodeForSets, nsgIds))

	if s.Res.PrivateEndpoint != nil {
		s.D.Set("private_endpoint", *s.Res.PrivateEndpoint)
	}

	if s.Res.PrivateEndpointLabel != nil {
		s.D.Set("private_endpoint_label", *s.Res.PrivateEndpointLabel)
	}

	if s.Res.ServiceConsoleUrl != nil {
		s.D.Set("service_console_url", *s.Res.ServiceConsoleUrl)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
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

	if s.Res.TimeReclamationOfFreeAutonomousDatabase != nil {
		s.D.Set("time_reclamation_of_free_autonomous_database", s.Res.TimeReclamationOfFreeAutonomousDatabase.String())
	}

	if s.Res.UsedDataStorageSizeInTBs != nil {
		s.D.Set("used_data_storage_size_in_tbs", *s.Res.UsedDataStorageSizeInTBs)
	}

	whitelistedIps := []interface{}{}
	for _, item := range s.Res.WhitelistedIps {
		whitelistedIps = append(whitelistedIps, item)
	}
	s.D.Set("whitelisted_ips", schema.NewSet(literalTypeHashCodeForSets, whitelistedIps))

	return nil
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

	return result
}

func AutonomousDatabaseConnectionUrlsToMap(obj *oci_database.AutonomousDatabaseConnectionUrls) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApexUrl != nil {
		result["apex_url"] = string(*obj.ApexUrl)
	}

	if obj.MachineLearningUserManagementUrl != nil {
		result["machine_learning_user_management_url"] = string(*obj.MachineLearningUserManagementUrl)
	}

	if obj.SqlDevWebUrl != nil {
		result["sql_dev_web_url"] = string(*obj.SqlDevWebUrl)
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
		if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
			tmp := autonomousContainerDatabaseId.(string)
			details.AutonomousContainerDatabaseId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
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
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if isPreviewVersionWithServiceTermsAccepted, ok := s.D.GetOkExists("is_preview_version_with_service_terms_accepted"); ok {
			tmp := isPreviewVersionWithServiceTermsAccepted.(bool)
			details.IsPreviewVersionWithServiceTermsAccepted = &tmp
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
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
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
		if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
			tmp := autonomousContainerDatabaseId.(string)
			details.AutonomousContainerDatabaseId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
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
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if isPreviewVersionWithServiceTermsAccepted, ok := s.D.GetOkExists("is_preview_version_with_service_terms_accepted"); ok {
			tmp := isPreviewVersionWithServiceTermsAccepted.(bool)
			details.IsPreviewVersionWithServiceTermsAccepted = &tmp
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
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
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
		if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
			tmp := autonomousContainerDatabaseId.(string)
			details.AutonomousContainerDatabaseId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
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
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if isPreviewVersionWithServiceTermsAccepted, ok := s.D.GetOkExists("is_preview_version_with_service_terms_accepted"); ok {
			tmp := isPreviewVersionWithServiceTermsAccepted.(bool)
			details.IsPreviewVersionWithServiceTermsAccepted = &tmp
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
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
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
		if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
			tmp := autonomousContainerDatabaseId.(string)
			details.AutonomousContainerDatabaseId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
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
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists("is_auto_scaling_enabled"); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
			tmp := isDedicated.(bool)
			details.IsDedicated = &tmp
		}
		if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
			tmp := isFreeTier.(bool)
			details.IsFreeTier = &tmp
		}
		if isPreviewVersionWithServiceTermsAccepted, ok := s.D.GetOkExists("is_preview_version_with_service_terms_accepted"); ok {
			tmp := isPreviewVersionWithServiceTermsAccepted.(bool)
			details.IsPreviewVersionWithServiceTermsAccepted = &tmp
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
		if privateEndpointLabel, ok := s.D.GetOkExists("private_endpoint_label"); ok {
			tmp := privateEndpointLabel.(string)
			details.PrivateEndpointLabel = &tmp
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
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

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeAutonomousDatabaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) updateDataSafeStatus(autonomousDatabaseId string, dataSafeStatus oci_database.AutonomousDatabaseDataSafeStatusEnum) error {
	switch dataSafeStatus {
	case oci_database.AutonomousDatabaseDataSafeStatusRegistered:
		request := oci_database.RegisterAutonomousDatabaseDataSafeRequest{}
		request.AutonomousDatabaseId = &autonomousDatabaseId
		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.RegisterAutonomousDatabaseDataSafe(context.Background(), request)

		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}

		return nil
	case oci_database.AutonomousDatabaseDataSafeStatusNotRegistered:
		request := oci_database.DeregisterAutonomousDatabaseDataSafeRequest{}
		request.AutonomousDatabaseId = &autonomousDatabaseId
		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.DeregisterAutonomousDatabaseDataSafe(context.Background(), request)

		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
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

	changeDbVersionRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousDatabase(context.Background(), changeDbVersionRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}

	return nil
}
