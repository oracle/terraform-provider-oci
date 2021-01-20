// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v33/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v33/workrequests"
)

func init() {
	RegisterResource("oci_database_vm_cluster", DatabaseVmClusterResource())
}

func DatabaseVmClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatabaseVmCluster,
		Read:     readDatabaseVmCluster,
		Update:   updateDatabaseVmCluster,
		Delete:   deleteDatabaseVmCluster,
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssh_public_keys": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      literalTypeHashCodeForSets,
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
			"time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
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
			"shape": {
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
	}
}

func createDatabaseVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return CreateResource(d, sync)
}

func readDatabaseVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

func updateDatabaseVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient

	return UpdateResource(d, sync)
}

func deleteDatabaseVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseVmClusterResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	workRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.VmCluster
	DisableNotFoundRetries bool
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

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
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

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
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

	if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
		tmp := timeZone.(string)
		request.TimeZone = &tmp
	}

	if vmClusterNetworkId, ok := s.D.GetOkExists("vm_cluster_network_id"); ok {
		tmp := vmClusterNetworkId.(string)
		request.VmClusterNetworkId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

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

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok && s.D.HasChange("cpu_core_count") {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
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
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok && s.D.HasChange("license_model") {
		request.LicenseModel = oci_database.UpdateVmClusterDetailsLicenseModelEnum(licenseModel.(string))
	}

	if memorySizeInGBs, ok := s.D.GetOkExists("memory_size_in_gbs"); ok {
		tmp := memorySizeInGBs.(int)
		request.MemorySizeInGBs = &tmp
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteVmCluster(context.Background(), request)
	return err
}

func (s *DatabaseVmClusterResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpusEnabled != nil {
		s.D.Set("cpus_enabled", *s.Res.CpusEnabled)
		s.D.Set("cpu_core_count", *s.Res.CpusEnabled)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExadataInfrastructureId != nil {
		s.D.Set("exadata_infrastructure_id", *s.Res.ExadataInfrastructureId)
	}

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

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	sshPublicKeys := []interface{}{}
	for _, item := range s.Res.SshPublicKeys {
		sshPublicKeys = append(sshPublicKeys, item)
	}
	s.D.Set("ssh_public_keys", schema.NewSet(literalTypeHashCodeForSets, sshPublicKeys))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	if s.Res.VmClusterNetworkId != nil {
		s.D.Set("vm_cluster_network_id", *s.Res.VmClusterNetworkId)
	}

	return nil
}

func (s *DatabaseVmClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeVmClusterCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VmClusterId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeVmClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "vmCluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}

	return nil
}
