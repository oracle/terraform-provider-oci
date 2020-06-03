// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func init() {
	RegisterResource("oci_database_autonomous_container_database", DatabaseAutonomousContainerDatabaseResource())
}

func DatabaseAutonomousContainerDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: getTimeoutDuration("12h"),
			Update: getTimeoutDuration("12h"),
			Delete: getTimeoutDuration("12h"),
		},
		Create: createDatabaseAutonomousContainerDatabase,
		Read:   readDatabaseAutonomousContainerDatabase,
		Update: updateDatabaseAutonomousContainerDatabase,
		Delete: deleteDatabaseAutonomousContainerDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patch_model": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"backup_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"recovery_window_in_days": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
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
			"maintenance_window_details": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
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
							MaxItems: 20,
							MinItems: 0,
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
			"service_level_agreement_type": {
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
				MaxItems: 1,
				MinItems: 1,
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
							MaxItems: 20,
							MinItems: 0,
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
							MaxItems: 4,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
					},
				},
			},
			"next_maintenance_run_id": {
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

func createDatabaseAutonomousContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return CreateResource(d, sync)
}

func readDatabaseAutonomousContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

func updateDatabaseAutonomousContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return UpdateResource(d, sync)
}

func deleteDatabaseAutonomousContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseAutonomousContainerDatabaseResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousContainerDatabase
	DisableNotFoundRetries bool
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateProvisioning),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateBackupInProgress),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateRestoring),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminating),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateProvisioning),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateUpdating),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateRestarting),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateMaintenanceInProgress),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) Create() error {
	request := oci_database.CreateAutonomousContainerDatabaseRequest{}

	if autonomousExadataInfrastructureId, ok := s.D.GetOkExists("autonomous_exadata_infrastructure_id"); ok {
		tmp := autonomousExadataInfrastructureId.(string)
		request.AutonomousExadataInfrastructureId = &tmp
	}

	if backupConfig, ok := s.D.GetOkExists("backup_config"); ok {
		if tmpList := backupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_config", 0)
			tmp, err := s.mapToAutonomousContainerDatabaseBackupConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupConfig = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	if patchModel, ok := s.D.GetOkExists("patch_model"); ok {
		request.PatchModel = oci_database.CreateAutonomousContainerDatabaseDetailsPatchModelEnum(patchModel.(string))
	}

	if serviceLevelAgreementType, ok := s.D.GetOkExists("service_level_agreement_type"); ok {
		request.ServiceLevelAgreementType = oci_database.CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum(serviceLevelAgreementType.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousContainerDatabase
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) Get() error {
	request := oci_database.GetAutonomousContainerDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousContainerDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousContainerDatabase
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateAutonomousContainerDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousContainerDatabaseId = &tmp

	if backupConfig, ok := s.D.GetOkExists("backup_config"); ok {
		if tmpList := backupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_config", 0)
			tmp, err := s.mapToAutonomousContainerDatabaseBackupConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupConfig = &tmp
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

	if patchModel, ok := s.D.GetOkExists("patch_model"); ok {
		request.PatchModel = oci_database.UpdateAutonomousContainerDatabaseDetailsPatchModelEnum(patchModel.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousContainerDatabase
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) Delete() error {
	request := oci_database.TerminateAutonomousContainerDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousContainerDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.TerminateAutonomousContainerDatabase(context.Background(), request)
	return err
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) SetData() error {
	if s.Res.AutonomousExadataInfrastructureId != nil {
		s.D.Set("autonomous_exadata_infrastructure_id", *s.Res.AutonomousExadataInfrastructureId)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.BackupConfig != nil {
		s.D.Set("backup_config", []interface{}{AutonomousContainerDatabaseBackupConfigToMap(s.Res.BackupConfig)})
	} else {
		s.D.Set("backup_config", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LastMaintenanceRunId != nil {
		s.D.Set("last_maintenance_run_id", *s.Res.LastMaintenanceRunId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	s.D.Set("patch_model", s.Res.PatchModel)

	s.D.Set("service_level_agreement_type", s.Res.ServiceLevelAgreementType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) mapToAutonomousContainerDatabaseBackupConfig(fieldKeyFormat string) (oci_database.AutonomousContainerDatabaseBackupConfig, error) {
	result := oci_database.AutonomousContainerDatabaseBackupConfig{}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	return result, nil
}

func AutonomousContainerDatabaseBackupConfigToMap(obj *oci_database.AutonomousContainerDatabaseBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RecoveryWindowInDays != nil {
		result["recovery_window_in_days"] = int(*obj.RecoveryWindowInDays)
	}

	return result
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) mapToDayOfWeek(fieldKeyFormat string) (oci_database.DayOfWeek, error) {
	result := oci_database.DayOfWeek{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.DayOfWeekNameEnum(name.(string))
	}

	return result, nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) mapToMaintenanceWindow(fieldKeyFormat string) (oci_database.MaintenanceWindow, error) {
	result := oci_database.MaintenanceWindow{}

	if preference, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preference")); ok {
		result.Preference = oci_database.MaintenanceWindowPreferenceEnum(preference.(string))

		if result.Preference == oci_database.MaintenanceWindowPreferenceNoPreference {
			return result, nil
		}
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

	if leadTimeInWeeks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lead_time_in_weeks")); ok {
		tmp := leadTimeInWeeks.(int)
		if tmp > 0 {
			result.LeadTimeInWeeks = &tmp
		}
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

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) mapToMonth(fieldKeyFormat string) (oci_database.Month, error) {
	result := oci_database.Month{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.MonthNameEnum(name.(string))
	}

	return result, nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeAutonomousContainerDatabaseCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AutonomousContainerDatabaseId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeAutonomousContainerDatabaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
