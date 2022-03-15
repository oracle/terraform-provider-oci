// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v62/database"
)

func DatabaseAutonomousVmClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseAutonomousVmCluster,
		Read:     readDatabaseAutonomousVmCluster,
		Update:   updateDatabaseAutonomousVmCluster,
		Delete:   deleteDatabaseAutonomousVmCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
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
			"vm_cluster_network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"autonomous_data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cpu_core_count_per_node": {
				Type:     schema.TypeInt,
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
			"is_local_backup_enabled": {
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
			"maintenance_window_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"preference": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"days_of_week": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"hours_of_day": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
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
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"weeks_of_month": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},

						// Computed
					},
				},
			},
			"memory_per_oracle_compute_unit_in_gbs": {
				Type:     schema.TypeInt,
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
			"total_container_databases": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"available_cpus": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"available_data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"cpus_enabled": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"data_storage_size_in_gb": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"db_node_storage_size_in_gbs": {
				Type:     schema.TypeInt,
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
			"maintenance_window": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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
						"preference": {
							Type:     schema.TypeString,
							Computed: true,
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
			"ocpus_enabled": {
				Type:     schema.TypeFloat,
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

func createDatabaseAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseAutonomousVmClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousVmCluster
	DisableNotFoundRetries bool
}

func (s *DatabaseAutonomousVmClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousVmClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousVmClusterLifecycleStateProvisioning),
	}
}

func (s *DatabaseAutonomousVmClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousVmClusterLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousVmClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousVmClusterLifecycleStateTerminating),
	}
}

func (s *DatabaseAutonomousVmClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousVmClusterLifecycleStateTerminated),
	}
}

func (s *DatabaseAutonomousVmClusterResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.AutonomousVmClusterLifecycleStateProvisioning),
		string(oci_database.AutonomousVmClusterLifecycleStateUpdating),
		string(oci_database.AutonomousVmClusterLifecycleStateMaintenanceInProgress),
	}
}

func (s *DatabaseAutonomousVmClusterResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.AutonomousVmClusterLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousVmClusterResourceCrud) Create() error {
	request := oci_database.CreateAutonomousVmClusterRequest{}

	if autonomousDataStorageSizeInTBs, ok := s.D.GetOkExists("autonomous_data_storage_size_in_tbs"); ok {
		tmp := autonomousDataStorageSizeInTBs.(float64)
		request.AutonomousDataStorageSizeInTBs = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpuCoreCountPerNode, ok := s.D.GetOkExists("cpu_core_count_per_node"); ok {
		tmp := cpuCoreCountPerNode.(int)
		request.CpuCoreCountPerNode = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isLocalBackupEnabled, ok := s.D.GetOkExists("is_local_backup_enabled"); ok {
		tmp := isLocalBackupEnabled.(bool)
		request.IsLocalBackupEnabled = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database.CreateAutonomousVmClusterDetailsLicenseModelEnum(licenseModel.(string))
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

	if memoryPerOracleComputeUnitInGBs, ok := s.D.GetOkExists("memory_per_oracle_compute_unit_in_gbs"); ok {
		tmp := memoryPerOracleComputeUnitInGBs.(int)
		request.MemoryPerOracleComputeUnitInGBs = &tmp
	}

	if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
		tmp := timeZone.(string)
		request.TimeZone = &tmp
	}

	if totalContainerDatabases, ok := s.D.GetOkExists("total_container_databases"); ok {
		tmp := totalContainerDatabases.(int)
		request.TotalContainerDatabases = &tmp
	}

	if vmClusterNetworkId, ok := s.D.GetOkExists("vm_cluster_network_id"); ok {
		tmp := vmClusterNetworkId.(string)
		request.VmClusterNetworkId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousVmCluster
	return nil
}

func (s *DatabaseAutonomousVmClusterResourceCrud) Get() error {
	request := oci_database.GetAutonomousVmClusterRequest{}

	tmp := s.D.Id()
	request.AutonomousVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousVmCluster
	return nil
}

func (s *DatabaseAutonomousVmClusterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateAutonomousVmClusterRequest{}

	tmp := s.D.Id()
	request.AutonomousVmClusterId = &tmp

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
		request.LicenseModel = oci_database.UpdateAutonomousVmClusterDetailsLicenseModelEnum(licenseModel.(string))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousVmCluster
	return nil
}

func (s *DatabaseAutonomousVmClusterResourceCrud) Delete() error {
	request := oci_database.DeleteAutonomousVmClusterRequest{}

	tmp := s.D.Id()
	request.AutonomousVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteAutonomousVmCluster(context.Background(), request)
	return err
}

func (s *DatabaseAutonomousVmClusterResourceCrud) SetData() error {
	if s.Res.AutonomousDataStorageSizeInTBs != nil {
		s.D.Set("autonomous_data_storage_size_in_tbs", *s.Res.AutonomousDataStorageSizeInTBs)
	}

	if s.Res.AvailableCpus != nil {
		s.D.Set("available_cpus", *s.Res.AvailableCpus)
	}

	if s.Res.AvailableDataStorageSizeInTBs != nil {
		s.D.Set("available_data_storage_size_in_tbs", *s.Res.AvailableDataStorageSizeInTBs)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpuCoreCountPerNode != nil {
		s.D.Set("cpu_core_count_per_node", *s.Res.CpuCoreCountPerNode)
	}

	if s.Res.CpusEnabled != nil {
		s.D.Set("cpus_enabled", *s.Res.CpusEnabled)
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

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExadataInfrastructureId != nil {
		s.D.Set("exadata_infrastructure_id", *s.Res.ExadataInfrastructureId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsLocalBackupEnabled != nil {
		s.D.Set("is_local_backup_enabled", *s.Res.IsLocalBackupEnabled)
	}

	if s.Res.LastMaintenanceRunId != nil {
		s.D.Set("last_maintenance_run_id", *s.Res.LastMaintenanceRunId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{AvmMaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MemoryPerOracleComputeUnitInGBs != nil {
		s.D.Set("memory_per_oracle_compute_unit_in_gbs", *s.Res.MemoryPerOracleComputeUnitInGBs)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	if s.Res.OcpusEnabled != nil {
		s.D.Set("ocpus_enabled", *s.Res.OcpusEnabled)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	if s.Res.TotalContainerDatabases != nil {
		s.D.Set("total_container_databases", *s.Res.TotalContainerDatabases)
	}

	if s.Res.VmClusterNetworkId != nil {
		s.D.Set("vm_cluster_network_id", *s.Res.VmClusterNetworkId)
	}

	return nil
}

func (s *DatabaseAutonomousVmClusterResourceCrud) mapToDayOfWeek(fieldKeyFormat string) (oci_database.DayOfWeek, error) {
	result := oci_database.DayOfWeek{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.DayOfWeekNameEnum(name.(string))
	}

	return result, nil
}

func AvmDayOfWeekToMap(obj oci_database.DayOfWeek) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseAutonomousVmClusterResourceCrud) mapToMaintenanceWindow(fieldKeyFormat string) (oci_database.MaintenanceWindow, error) {
	result := oci_database.MaintenanceWindow{}

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

	if preference, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preference")); ok {
		result.Preference = oci_database.MaintenanceWindowPreferenceEnum(preference.(string))
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

func AvmMaintenanceWindowToMap(obj *oci_database.MaintenanceWindow) map[string]interface{} {
	result := map[string]interface{}{}

	daysOfWeek := []interface{}{}
	for _, item := range obj.DaysOfWeek {
		daysOfWeek = append(daysOfWeek, AvmDayOfWeekToMap(item))
	}
	result["days_of_week"] = daysOfWeek

	result["hours_of_day"] = obj.HoursOfDay

	if obj.LeadTimeInWeeks != nil {
		result["lead_time_in_weeks"] = int(*obj.LeadTimeInWeeks)
	}

	months := []interface{}{}
	for _, item := range obj.Months {
		months = append(months, AvmMonthToMap(item))
	}
	result["months"] = months

	result["preference"] = string(obj.Preference)

	result["weeks_of_month"] = obj.WeeksOfMonth

	return result
}

func (s *DatabaseAutonomousVmClusterResourceCrud) mapToMonth(fieldKeyFormat string) (oci_database.Month, error) {
	result := oci_database.Month{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.MonthNameEnum(name.(string))
	}

	return result, nil
}

func AvmMonthToMap(obj oci_database.Month) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseAutonomousVmClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeAutonomousVmClusterCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AutonomousVmClusterId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeAutonomousVmClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
