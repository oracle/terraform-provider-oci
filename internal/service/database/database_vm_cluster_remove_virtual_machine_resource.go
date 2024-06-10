// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseVmClusterRemoveVirtualMachineResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseVmClusterRemoveVirtualMachine,
		Read:     readDatabaseVmClusterRemoveVirtualMachine,
		Delete:   deleteDatabaseVmClusterRemoveVirtualMachine,
		Schema: map[string]*schema.Schema{
			// Required
			"db_servers": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"db_server_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_automation_update_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"apply_update_time_preference": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"apply_update_preferred_end_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"apply_update_preferred_start_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"freeze_period": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"freeze_period_end_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeze_period_start_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"is_early_adoption_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_freeze_period_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpus_enabled": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"data_collection_options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_diagnostics_events_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_health_monitoring_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_incident_logs_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"db_node_storage_size_in_gbs": {
				Type:     schema.TypeInt,
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
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_system_configuration_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"file_system_size_gb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mount_point": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"gi_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_local_backup_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_sparse_diskgroup_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_patch_history_entry_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_model": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_public_keys": {
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
			"system_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_cluster_network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseVmClusterRemoveVirtualMachine(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterRemoveVirtualMachineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseVmClusterRemoveVirtualMachine(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseVmClusterRemoveVirtualMachine(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseVmClusterRemoveVirtualMachineResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.VmCluster
	DisableNotFoundRetries bool
}

func (s *DatabaseVmClusterRemoveVirtualMachineResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseVmClusterRemoveVirtualMachineResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.VmClusterLifecycleStateProvisioning),
	}
}

func (s *DatabaseVmClusterRemoveVirtualMachineResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.VmClusterLifecycleStateAvailable),
	}
}

func (s *DatabaseVmClusterRemoveVirtualMachineResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.VmClusterLifecycleStateTerminating),
	}
}

func (s *DatabaseVmClusterRemoveVirtualMachineResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.VmClusterLifecycleStateTerminated),
	}
}

func (s *DatabaseVmClusterRemoveVirtualMachineResourceCrud) Create() error {
	request := oci_database.RemoveVirtualMachineFromVmClusterRequest{}

	if dbServers, ok := s.D.GetOkExists("db_servers"); ok {
		interfaces := dbServers.([]interface{})
		tmp := make([]oci_database.DbServerDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_servers", stateDataIndex)
			converted, err := s.mapToDbServerDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("db_servers") {
			request.DbServers = tmp
		}
	}

	if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
		tmp := vmClusterId.(string)
		request.VmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RemoveVirtualMachineFromVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VmCluster
	return nil
}

func (s *DatabaseVmClusterRemoveVirtualMachineResourceCrud) SetData() error {
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

	if s.Res.CpusEnabled != nil {
		s.D.Set("cpus_enabled", *s.Res.CpusEnabled)
	}

	if s.Res.DataCollectionOptions != nil {
		s.D.Set("data_collection_options", []interface{}{DataCollectionOptionsToMap(s.Res.DataCollectionOptions)})
	} else {
		s.D.Set("data_collection_options", nil)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExadataInfrastructureId != nil {
		s.D.Set("exadata_infrastructure_id", *s.Res.ExadataInfrastructureId)
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

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("ssh_public_keys", s.Res.SshPublicKeys)

	s.D.Set("state", s.Res.LifecycleState)

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

	return nil
}

func (s *DatabaseVmClusterRemoveVirtualMachineResourceCrud) CloudAutomationApplyUpdateTimePreferenceToMap(obj *oci_database.CloudAutomationApplyUpdateTimePreference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplyUpdatePreferredEndTime != nil {
		result["apply_update_preferred_end_time"] = string(*obj.ApplyUpdatePreferredEndTime)
	}

	if obj.ApplyUpdatePreferredStartTime != nil {
		result["apply_update_preferred_start_time"] = string(*obj.ApplyUpdatePreferredStartTime)
	}

	return result
}

func (s *DatabaseVmClusterRemoveVirtualMachineResourceCrud) CloudAutomationFreezePeriodToMap(obj *oci_database.CloudAutomationFreezePeriod) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FreezePeriodEndTime != nil {
		result["freeze_period_end_time"] = string(*obj.FreezePeriodEndTime)
	}

	if obj.FreezePeriodStartTime != nil {
		result["freeze_period_start_time"] = string(*obj.FreezePeriodStartTime)
	}

	return result
}

func (s *DatabaseVmClusterRemoveVirtualMachineResourceCrud) CloudAutomationUpdateDetailsToMap(obj *oci_database.CloudAutomationUpdateDetails) map[string]interface{} {
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

func (s *DatabaseVmClusterRemoveVirtualMachineResourceCrud) mapToDbServerDetails(fieldKeyFormat string) (oci_database.DbServerDetails, error) {
	result := oci_database.DbServerDetails{}

	if dbServerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_server_id")); ok {
		tmp := dbServerId.(string)
		result.DbServerId = &tmp
	}

	return result, nil
}
