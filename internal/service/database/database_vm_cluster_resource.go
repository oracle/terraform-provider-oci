// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabaseVmClusterResource() *schema.Resource {
	return &schema.Resource{
		CustomizeDiff: fileSystemConfigDiffFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("30m"),
			Update: tfresource.GetTimeoutDuration("30m"),
			Delete: tfresource.GetTimeoutDuration("30m"),
		},
		Create: createDatabaseVmCluster,
		Read:   readDatabaseVmCluster,
		Update: updateDatabaseVmCluster,
		Delete: deleteDatabaseVmCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"gi_version": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.GiVersionDiffSuppress,
			},
			"ssh_public_keys": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vm_cluster_network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"cloud_automation_update_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"apply_update_time_preference": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"apply_update_preferred_end_time": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"apply_update_preferred_start_time": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"freeze_period": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"freeze_period_end_time": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"freeze_period_start_time": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"is_early_adoption_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_freeze_period_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"data_storage_size_in_gb": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"db_node_storage_size_in_gbs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"exascale_db_storage_vault_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"file_system_configuration_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"file_system_size_gb": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mount_point": {
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
			"is_local_backup_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_sparse_diskgroup_enabled": {
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
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ocpu_count": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"system_version": {
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
			"vm_cluster_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compute_model": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpus_enabled": {
				Type:     schema.TypeInt,
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
			"ocpus_enabled": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_management_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseVmClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.VmCluster
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseVmClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseVmClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.VmClusterLifecycleStateProvisioning),
	}
}

func (s *DatabaseVmClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.VmClusterLifecycleStateAvailable),
	}
}

func (s *DatabaseVmClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.VmClusterLifecycleStateTerminating),
	}
}

func (s *DatabaseVmClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.VmClusterLifecycleStateTerminated),
	}
}

func (s *DatabaseVmClusterResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.VmClusterLifecycleStateUpdating),
		string(oci_database.VmClusterLifecycleStateMaintenanceInProgress),
	}
}

func (s *DatabaseVmClusterResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.VmClusterLifecycleStateAvailable),
	}
}

func (s *DatabaseVmClusterResourceCrud) Create() error {
	request := oci_database.CreateVmClusterRequest{}

	if cloudAutomationUpdateDetails, ok := s.D.GetOkExists("cloud_automation_update_details"); ok {
		if tmpList := cloudAutomationUpdateDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cloud_automation_update_details", 0)
			tmp, err := s.mapToCloudAutomationUpdateDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CloudAutomationUpdateDetails = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
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

	if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
		tmp := dataStorageSizeInGB.(float64)
		request.DataStorageSizeInGBs = &tmp
	}

	if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
		tmp := dataStorageSizeInTBs.(float64)
		request.DataStorageSizeInTBs = &tmp
	}

	if dbNodeStorageSizeInGBs, ok := s.D.GetOkExists("db_node_storage_size_in_gbs"); ok {
		tmp := dbNodeStorageSizeInGBs.(int)
		request.DbNodeStorageSizeInGBs = &tmp
	}

	if dbServers, ok := s.D.GetOkExists("db_servers"); ok {
		interfaces := dbServers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("db_servers") {
			request.DbServers = tmp
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

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	if exascaleDbStorageVaultId, ok := s.D.GetOkExists("exascale_db_storage_vault_id"); ok {
		tmp := exascaleDbStorageVaultId.(string)
		request.ExascaleDbStorageVaultId = &tmp
	}

	if fileSystemConfigurationDetails, ok := s.D.GetOkExists("file_system_configuration_details"); ok {
		interfaces := fileSystemConfigurationDetails.([]interface{})
		tmp := make([]oci_database.FileSystemConfigurationDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "file_system_configuration_details", stateDataIndex)
			converted, err := s.mapToFileSystemConfigurationDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("file_system_configuration_details") {
			request.FileSystemConfigurationDetails = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if giVersion, ok := s.D.GetOkExists("gi_version"); ok {
		tmp := giVersion.(string)
		request.GiVersion = &tmp
	}

	if isLocalBackupEnabled, ok := s.D.GetOkExists("is_local_backup_enabled"); ok {
		tmp := isLocalBackupEnabled.(bool)
		request.IsLocalBackupEnabled = &tmp
	}

	if isSparseDiskgroupEnabled, ok := s.D.GetOkExists("is_sparse_diskgroup_enabled"); ok {
		tmp := isSparseDiskgroupEnabled.(bool)
		request.IsSparseDiskgroupEnabled = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database.CreateVmClusterDetailsLicenseModelEnum(licenseModel.(string))
	}

	if memorySizeInGBs, ok := s.D.GetOkExists("memory_size_in_gbs"); ok {
		tmp := memorySizeInGBs.(int)
		request.MemorySizeInGBs = &tmp
	}

	if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok {
		tmp := float32(ocpuCount.(float64))
		request.OcpuCount = &tmp
	}

	if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
		request.SshPublicKeys = []string{}
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

	if systemVersion, ok := s.D.GetOkExists("system_version"); ok {
		tmp := systemVersion.(string)
		request.SystemVersion = &tmp
	}

	if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
		tmp := timeZone.(string)
		request.TimeZone = &tmp
	}

	if vmClusterNetworkId, ok := s.D.GetOkExists("vm_cluster_network_id"); ok {
		tmp := vmClusterNetworkId.(string)
		request.VmClusterNetworkId = &tmp
	}

	if vmClusterType, ok := s.D.GetOkExists("vm_cluster_type"); ok {
		request.VmClusterType = oci_database.CreateVmClusterDetailsVmClusterTypeEnum(vmClusterType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VmCluster
	return nil
}

func (s *DatabaseVmClusterResourceCrud) Get() error {
	request := oci_database.GetVmClusterRequest{}

	tmp := s.D.Id()
	request.VmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VmCluster
	return nil
}

func (s *DatabaseVmClusterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateVmClusterRequest{}

	if cloudAutomationUpdateDetails, ok := s.D.GetOkExists("cloud_automation_update_details"); ok && s.D.HasChange("cloud_automation_update_details") { //
		if tmpList := cloudAutomationUpdateDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cloud_automation_update_details", 0)
			tmp, err := s.mapToCloudAutomationUpdateDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CloudAutomationUpdateDetails = &tmp
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

	if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
		tmp := dataStorageSizeInGB.(float64)
		request.DataStorageSizeInGBs = &tmp
	}

	if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
		tmp := dataStorageSizeInTBs.(float64)
		request.DataStorageSizeInTBs = &tmp
	}

	if dbNodeStorageSizeInGBs, ok := s.D.GetOkExists("db_node_storage_size_in_gbs"); ok {
		tmp := dbNodeStorageSizeInGBs.(int)
		request.DbNodeStorageSizeInGBs = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if fileSystemConfigurationDetails, ok := s.D.GetOkExists("file_system_configuration_details"); ok {
		interfaces := fileSystemConfigurationDetails.([]interface{})
		tmp := make([]oci_database.FileSystemConfigurationDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "file_system_configuration_details", stateDataIndex)
			converted, err := s.mapToFileSystemConfigurationDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("file_system_configuration_details") {
			request.FileSystemConfigurationDetails = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok && s.D.HasChange("license_model") {
		request.LicenseModel = oci_database.UpdateVmClusterDetailsLicenseModelEnum(licenseModel.(string))
	}

	if memorySizeInGBs, ok := s.D.GetOkExists("memory_size_in_gbs"); ok {
		tmp := memorySizeInGBs.(int)
		request.MemorySizeInGBs = &tmp
	}

	if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok {
		tmp := float32(ocpuCount.(float64))
		request.OcpuCount = &tmp
	}

	if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok && s.D.HasChange("ssh_public_keys") {
		request.SshPublicKeys = []string{}
		set := sshPublicKeys.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.SshPublicKeys = tmp
	}

	tmp := s.D.Id()
	request.VmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VmCluster
	return nil
}

func (s *DatabaseVmClusterResourceCrud) Delete() error {
	request := oci_database.DeleteVmClusterRequest{}

	tmp := s.D.Id()
	request.VmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteVmCluster(context.Background(), request)
	return err
}

func (s *DatabaseVmClusterResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CloudAutomationUpdateDetails != nil {
		s.D.Set("cloud_automation_update_details", []interface{}{CloudAutomationUpdateDetailsToMap(s.Res.CloudAutomationUpdateDetails)})
	} else {
		s.D.Set("cloud_automation_update_details", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("compute_model", s.Res.ComputeModel)

	if s.Res.CpusEnabled != nil {
		s.D.Set("cpus_enabled", *s.Res.CpusEnabled)
		s.D.Set("cpu_core_count", *s.Res.CpusEnabled)
	}

	if s.Res.DataCollectionOptions != nil {
		s.D.Set("data_collection_options", []interface{}{DataCollectionOptionsToMap(s.Res.DataCollectionOptions)})
	} else {
		s.D.Set("data_collection_options", nil)
	}

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	s.D.Set("db_servers", s.Res.DbServers)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExadataInfrastructureId != nil {
		s.D.Set("exadata_infrastructure_id", *s.Res.ExadataInfrastructureId)
	}

	if s.Res.ExascaleDbStorageVaultId != nil {
		s.D.Set("exascale_db_storage_vault_id", *s.Res.ExascaleDbStorageVaultId)
	}

	fileSystemConfigurationDetails := []interface{}{}
	for _, item := range s.Res.FileSystemConfigurationDetails {
		fileSystemConfigurationDetails = append(fileSystemConfigurationDetails, FileSystemConfigurationDetailToMap(item))
	}
	s.D.Set("file_system_configuration_details", fileSystemConfigurationDetails)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GiVersion != nil {
		s.D.Set("gi_version", *s.Res.GiVersion)
	}

	if s.Res.IsLocalBackupEnabled != nil {
		s.D.Set("is_local_backup_enabled", *s.Res.IsLocalBackupEnabled)
	}

	if s.Res.IsSparseDiskgroupEnabled != nil {
		s.D.Set("is_sparse_diskgroup_enabled", *s.Res.IsSparseDiskgroupEnabled)
	}

	if s.Res.LastPatchHistoryEntryId != nil {
		s.D.Set("last_patch_history_entry_id", *s.Res.LastPatchHistoryEntryId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.OcpusEnabled != nil {
		s.D.Set("ocpus_enabled", *s.Res.OcpusEnabled)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	sshPublicKeys := []interface{}{}
	for _, item := range s.Res.SshPublicKeys {
		sshPublicKeys = append(sshPublicKeys, item)
	}
	s.D.Set("ssh_public_keys", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, sshPublicKeys))

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("storage_management_type", s.Res.StorageManagementType)

	if s.Res.SystemVersion != nil {
		s.D.Set("system_version", *s.Res.SystemVersion)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	if s.Res.VmClusterNetworkId != nil {
		s.D.Set("vm_cluster_network_id", *s.Res.VmClusterNetworkId)
	}

	s.D.Set("vm_cluster_type", s.Res.VmClusterType)

	return nil
}

func (s *DatabaseVmClusterResourceCrud) mapToCloudAutomationApplyUpdateTimePreference(fieldKeyFormat string) (oci_database.CloudAutomationApplyUpdateTimePreference, error) {
	result := oci_database.CloudAutomationApplyUpdateTimePreference{}

	if applyUpdatePreferredEndTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "apply_update_preferred_end_time")); ok {
		tmp := applyUpdatePreferredEndTime.(string)
		result.ApplyUpdatePreferredEndTime = &tmp
	}

	if applyUpdatePreferredStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "apply_update_preferred_start_time")); ok {
		tmp := applyUpdatePreferredStartTime.(string)
		result.ApplyUpdatePreferredStartTime = &tmp
	}

	return result, nil
}

func (s *DatabaseVmClusterResourceCrud) CloudAutomationApplyUpdateTimePreferenceToMap(obj *oci_database.CloudAutomationApplyUpdateTimePreference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplyUpdatePreferredEndTime != nil {
		result["apply_update_preferred_end_time"] = string(*obj.ApplyUpdatePreferredEndTime)
	}

	if obj.ApplyUpdatePreferredStartTime != nil {
		result["apply_update_preferred_start_time"] = string(*obj.ApplyUpdatePreferredStartTime)
	}

	return result
}

func (s *DatabaseVmClusterResourceCrud) mapToCloudAutomationFreezePeriod(fieldKeyFormat string) (oci_database.CloudAutomationFreezePeriod, error) {
	result := oci_database.CloudAutomationFreezePeriod{}

	if freezePeriodEndTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeze_period_end_time")); ok {
		tmp := freezePeriodEndTime.(string)
		result.FreezePeriodEndTime = &tmp
	}

	if freezePeriodStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeze_period_start_time")); ok {
		tmp := freezePeriodStartTime.(string)
		result.FreezePeriodStartTime = &tmp
	}

	return result, nil
}

func (s *DatabaseVmClusterResourceCrud) CloudAutomationFreezePeriodToMap(obj *oci_database.CloudAutomationFreezePeriod) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FreezePeriodEndTime != nil {
		result["freeze_period_end_time"] = string(*obj.FreezePeriodEndTime)
	}

	if obj.FreezePeriodStartTime != nil {
		result["freeze_period_start_time"] = string(*obj.FreezePeriodStartTime)
	}

	return result
}

func (s *DatabaseVmClusterResourceCrud) mapToCloudAutomationUpdateDetails(fieldKeyFormat string) (oci_database.CloudAutomationUpdateDetails, error) {
	result := oci_database.CloudAutomationUpdateDetails{}

	if applyUpdateTimePreference, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "apply_update_time_preference")); ok {
		if tmpList := applyUpdateTimePreference.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "apply_update_time_preference"), 0)
			tmp, err := s.mapToCloudAutomationApplyUpdateTimePreference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert apply_update_time_preference, encountered error: %v", err)
			}
			result.ApplyUpdateTimePreference = &tmp
		}
	}

	if freezePeriod, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeze_period")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "freeze_period")) {
		if tmpList := freezePeriod.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "freeze_period"), 0)
			tmp, err := s.mapToCloudAutomationFreezePeriod(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert freeze_period, encountered error: %v", err)
			}
			result.FreezePeriod = &tmp
		}
	}

	if isEarlyAdoptionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_early_adoption_enabled")); ok {
		tmp := isEarlyAdoptionEnabled.(bool)
		result.IsEarlyAdoptionEnabled = &tmp
	}

	if isFreezePeriodEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_freeze_period_enabled")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "is_freeze_period_enabled")) {
		tmp := isFreezePeriodEnabled.(bool)
		result.IsFreezePeriodEnabled = &tmp
	}

	return result, nil
}

func (s *DatabaseVmClusterResourceCrud) CloudAutomationUpdateDetailsToMap(obj *oci_database.CloudAutomationUpdateDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplyUpdateTimePreference != nil {
		result["apply_update_time_preference"] = []interface{}{CloudAutomationApplyUpdateTimePreferenceToMap(obj.ApplyUpdateTimePreference)}
	}

	if obj.FreezePeriod != nil {
		result["freeze_period"] = []interface{}{CloudAutomationFreezePeriodToMap(obj.FreezePeriod)}
	}

	if obj.IsEarlyAdoptionEnabled != nil {
		result["is_early_adoption_enabled"] = bool(*obj.IsEarlyAdoptionEnabled)
	}

	if obj.IsFreezePeriodEnabled != nil {
		result["is_freeze_period_enabled"] = bool(*obj.IsFreezePeriodEnabled)
	}

	return result
}

func (s *DatabaseVmClusterResourceCrud) mapToDataCollectionOptions(fieldKeyFormat string) (oci_database.DataCollectionOptions, error) {
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

func DataCollectionOptionsToMap(obj *oci_database.DataCollectionOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDiagnosticsEventsEnabled != nil {
		result["is_diagnostics_events_enabled"] = bool(*obj.IsDiagnosticsEventsEnabled)
	}

	if obj.IsHealthMonitoringEnabled != nil {
		result["is_health_monitoring_enabled"] = bool(*obj.IsHealthMonitoringEnabled)
	}

	if obj.IsIncidentLogsEnabled != nil {
		result["is_incident_logs_enabled"] = bool(*obj.IsIncidentLogsEnabled)
	}

	return result
}

func (s *DatabaseVmClusterResourceCrud) mapToFileSystemConfigurationDetail(fieldKeyFormat string) (oci_database.FileSystemConfigurationDetail, error) {
	result := oci_database.FileSystemConfigurationDetail{}

	if fileSystemSizeGb, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_system_size_gb")); ok {
		tmp := fileSystemSizeGb.(int)
		result.FileSystemSizeGb = &tmp
	}

	if mountPoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_point")); ok {
		tmp := mountPoint.(string)
		result.MountPoint = &tmp
	}

	return result, nil
}

func (s *DatabaseVmClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeVmClusterCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VmClusterId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeVmClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "vmCluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}
