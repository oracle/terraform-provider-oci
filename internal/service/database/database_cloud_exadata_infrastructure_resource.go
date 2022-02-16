// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseCloudExadataInfrastructureResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		Create: createDatabaseCloudExadataInfrastructure,
		Read:   readDatabaseCloudExadataInfrastructure,
		Update: updateDatabaseCloudExadataInfrastructure,
		Delete: deleteDatabaseCloudExadataInfrastructure,
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
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compute_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"customer_contacts": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"email": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"maintenance_window": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
			"storage_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"available_storage_size_in_gbs": {
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
			"total_storage_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createDatabaseCloudExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseCloudExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseCloudExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseCloudExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseCloudExadataInfrastructureResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.CloudExadataInfrastructure
	DisableNotFoundRetries bool
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateProvisioning),
	}
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateAvailable),
	}
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateTerminating),
	}
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateTerminated),
	}
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateProvisioning),
		string(oci_database.CloudExadataInfrastructureLifecycleStateUpdating),
	}
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateAvailable),
	}
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) Create() error {
	request := oci_database.CreateCloudExadataInfrastructureRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeCount, ok := s.D.GetOkExists("compute_count"); ok {
		tmp := computeCount.(int)
		request.ComputeCount = &tmp
	}

	if customerContacts, ok := s.D.GetOkExists("customer_contacts"); ok {
		interfaces := customerContacts.([]interface{})
		tmp := make([]oci_database.CustomerContact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "customer_contacts", stateDataIndex)
			converted, err := s.mapToCustomerContact(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("customer_contacts") {
			request.CustomerContacts = tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if maintenanceWindow, ok := s.D.GetOkExists("maintenance_window"); ok {
		if tmpList := maintenanceWindow.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindow = &tmp
		}
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if storageCount, ok := s.D.GetOkExists("storage_count"); ok {
		tmp := storageCount.(int)
		request.StorageCount = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateCloudExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudExadataInfrastructure
	return nil
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) Get() error {
	request := oci_database.GetCloudExadataInfrastructureRequest{}

	tmp := s.D.Id()
	request.CloudExadataInfrastructureId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetCloudExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudExadataInfrastructure
	return nil
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateCloudExadataInfrastructureRequest{}

	tmp := s.D.Id()
	request.CloudExadataInfrastructureId = &tmp

	if computeCount, ok := s.D.GetOkExists("compute_count"); ok {
		tmp := computeCount.(int)
		request.ComputeCount = &tmp
	}

	if customerContacts, ok := s.D.GetOkExists("customer_contacts"); ok {
		interfaces := customerContacts.([]interface{})
		tmp := make([]oci_database.CustomerContact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "customer_contacts", stateDataIndex)
			converted, err := s.mapToCustomerContact(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("customer_contacts") {
			request.CustomerContacts = tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if maintenanceWindow, ok := s.D.GetOkExists("maintenance_window"); ok {
		if tmpList := maintenanceWindow.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindow = &tmp
		}
	}

	if storageCount, ok := s.D.GetOkExists("storage_count"); ok {
		tmp := storageCount.(int)
		request.StorageCount = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateCloudExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudExadataInfrastructure
	return nil
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) Delete() error {
	request := oci_database.DeleteCloudExadataInfrastructureRequest{}

	tmp := s.D.Id()
	request.CloudExadataInfrastructureId = &tmp

	if isDeleteVmClusters, ok := s.D.GetOkExists("is_delete_vm_clusters"); ok {
		tmp := isDeleteVmClusters.(bool)
		request.IsDeleteVmClusters = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteCloudExadataInfrastructure(context.Background(), request)
	return err
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.AvailableStorageSizeInGBs != nil {
		s.D.Set("available_storage_size_in_gbs", *s.Res.AvailableStorageSizeInGBs)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeCount != nil {
		s.D.Set("compute_count", *s.Res.ComputeCount)
	}

	customerContacts := []interface{}{}
	for _, item := range s.Res.CustomerContacts {
		customerContacts = append(customerContacts, CustomerContactToMap(item))
	}
	s.D.Set("customer_contacts", customerContacts)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageCount != nil {
		s.D.Set("storage_count", *s.Res.StorageCount)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TotalStorageSizeInGBs != nil {
		s.D.Set("total_storage_size_in_gbs", *s.Res.TotalStorageSizeInGBs)
	}

	return nil
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) mapToCustomerContact(fieldKeyFormat string) (oci_database.CustomerContact, error) {
	result := oci_database.CustomerContact{}

	if email, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email")); ok {
		tmp := email.(string)
		result.Email = &tmp
	}

	return result, nil
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) mapToDayOfWeek(fieldKeyFormat string) (oci_database.DayOfWeek, error) {
	result := oci_database.DayOfWeek{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.DayOfWeekNameEnum(name.(string))
	}

	return result, nil
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) mapToMaintenanceWindow(fieldKeyFormat string) (oci_database.MaintenanceWindow, error) {
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

func (s *DatabaseCloudExadataInfrastructureResourceCrud) mapToMonth(fieldKeyFormat string) (oci_database.Month, error) {
	result := oci_database.Month{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.MonthNameEnum(name.(string))
	}

	return result, nil
}

func (s *DatabaseCloudExadataInfrastructureResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeCloudExadataInfrastructureCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CloudExadataInfrastructureId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeCloudExadataInfrastructureCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
