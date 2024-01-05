// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseExadataInfrastructureComputeManagedResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   updateDatabaseExadataInfrastructureComputeManagedResource,
		Read:     readDatabaseExadataInfrastructureComputeManagedResource,
		Update:   updateDatabaseExadataInfrastructureComputeManagedResource,
		Delete:   deleteDatabaseExadataInfrastructureComputeManagedResource,
		Schema: map[string]*schema.Schema{
			//Required
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			//Optional
			"additional_compute_count_compute_managed_resource": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"additional_compute_system_model_compute_managed_resource": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"activation_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			//Computed
			"additional_compute_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"additional_compute_system_model": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_network_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_control_plane_server1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_control_plane_server2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_server": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"infini_band_network_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_server": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_async": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"compute_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"contacts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Computed
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_primary": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_contact_mos_validated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"phone_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"corporate_proxy": {
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
			"is_cps_offline_report_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"maintenance_window": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"preference": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"custom_action_timeout_in_mins": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"days_of_week": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
						"is_custom_action_timeout_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_monthly_patching_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
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
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"patching_mode": {
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
			"storage_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"activated_storage_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"additional_storage_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cpus_enabled": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"csi_number": {
				Type:     schema.TypeString,
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
			"db_server_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_slo_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_cpu_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_data_storage_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"max_db_node_storage_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_memory_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"monthly_db_server_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_server_version": {
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

func updateDatabaseExadataInfrastructureComputeManagedResource(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureComputeManagedResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func readDatabaseExadataInfrastructureComputeManagedResource(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureComputeManagedResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func deleteDatabaseExadataInfrastructureComputeManagedResource(d *schema.ResourceData, m interface{}) error {
	return nil
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateActivating),
		string(oci_database.ExadataInfrastructureLifecycleStateUpdating),
		string(oci_database.ExadataInfrastructureLifecycleStateMaintenanceInProgress),
	}
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateRequiresActivation),
		string(oci_database.ExadataInfrastructureLifecycleStateActive),
		string(oci_database.ExadataInfrastructureLifecycleStateActivationFailed),
		string(oci_database.ExadataInfrastructureLifecycleStateDisconnected),
	}
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateDeleting),
		"TERMINATING",
	}
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateDeleted),
		"TERMINATED",
	}
}

type DatabaseExadataInfrastructureComputeManagedResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.ExadataInfrastructure
	DisableNotFoundRetries bool
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) Update() error {
	request := oci_database.UpdateExadataInfrastructureRequest{}

	if additionalComputeCount, ok := s.D.GetOkExists("additional_compute_count_compute_managed_resource"); ok {
		tmp := additionalComputeCount.(int)
		request.AdditionalComputeCount = &tmp
	}

	if additionalComputeSystemModel, ok := s.D.GetOkExists("additional_compute_system_model_compute_managed_resource"); ok {
		request.AdditionalComputeSystemModel = oci_database.UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum(additionalComputeSystemModel.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		s.D.Set("exadata_infrastructure_id", &tmp)
		request.ExadataInfrastructureId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	if s.D.Get("state").(string) == string(oci_database.ExadataInfrastructureLifecycleStateRequiresActivation) ||
		s.D.Get("additional_compute_count").(int) > 0 {
		if activationFile, ok := s.D.GetOkExists("activation_file"); ok &&
			s.D.Get("activation_file").(string) != "" {
			response, err := s.activateExadataInfrastructureComputeManagedResource(activationFile.(string), s.D.Id())
			if err != nil {
				s.D.Set("activation_file", "")
				return err
			}
			s.Res = &response.ExadataInfrastructure
			return nil
		}
	}

	response, err := s.Client.UpdateExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataInfrastructure

	s.D.SetId(*s.Res.Id)

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exadataInfrastructure", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) SetData() error {

	s.D.SetId(*s.Res.Id)

	if s.Res.AdditionalComputeCount != nil {
		s.D.Set("additional_compute_count", *s.Res.AdditionalComputeCount)
	}

	s.D.Set("additional_compute_system_model", s.Res.AdditionalComputeSystemModel)

	if s.Res.ActivatedStorageCount != nil {
		s.D.Set("activated_storage_count", *s.Res.ActivatedStorageCount)
	}

	if s.Res.AdditionalStorageCount != nil {
		s.D.Set("additional_storage_count", *s.Res.AdditionalStorageCount)
	}

	if s.Res.AdminNetworkCIDR != nil {
		s.D.Set("admin_network_cidr", *s.Res.AdminNetworkCIDR)
	}

	if s.Res.CloudControlPlaneServer1 != nil {
		s.D.Set("cloud_control_plane_server1", *s.Res.CloudControlPlaneServer1)
	}

	if s.Res.CloudControlPlaneServer2 != nil {
		s.D.Set("cloud_control_plane_server2", *s.Res.CloudControlPlaneServer2)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeCount != nil {
		s.D.Set("compute_count", *s.Res.ComputeCount)
	}

	contacts := []interface{}{}
	for _, item := range s.Res.Contacts {
		contacts = append(contacts, ExadataInfrastructureContactToMap(item))
	}
	s.D.Set("contacts", contacts)

	if s.Res.CorporateProxy != nil {
		s.D.Set("corporate_proxy", *s.Res.CorporateProxy)
	}

	if s.Res.CpusEnabled != nil {
		s.D.Set("cpus_enabled", *s.Res.CpusEnabled)
	}

	if s.Res.CsiNumber != nil {
		s.D.Set("csi_number", *s.Res.CsiNumber)
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

	s.D.Set("dns_server", s.Res.DnsServer)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Gateway != nil {
		s.D.Set("gateway", *s.Res.Gateway)
	}

	if s.Res.InfiniBandNetworkCIDR != nil {
		s.D.Set("infini_band_network_cidr", *s.Res.InfiniBandNetworkCIDR)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("maintenance_slo_status", s.Res.MaintenanceSLOStatus)

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{ExadataInfrastructureMaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MaxCpuCount != nil {
		s.D.Set("max_cpu_count", *s.Res.MaxCpuCount)
	}

	if s.Res.MaxDataStorageInTBs != nil {
		s.D.Set("max_data_storage_in_tbs", *s.Res.MaxDataStorageInTBs)
	}

	if s.Res.MaxDbNodeStorageInGBs != nil {
		s.D.Set("max_db_node_storage_in_gbs", *s.Res.MaxDbNodeStorageInGBs)
	}

	if s.Res.MaxMemoryInGBs != nil {
		s.D.Set("max_memory_in_gbs", *s.Res.MaxMemoryInGBs)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.Netmask != nil {
		s.D.Set("netmask", *s.Res.Netmask)
	}

	s.D.Set("ntp_server", s.Res.NtpServer)

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

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	return nil
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) Get() error {
	request := oci_database.GetExadataInfrastructureRequest{}

	tmp := s.D.Id()
	request.ExadataInfrastructureId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataInfrastructure
	return nil
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) activateExadataInfrastructureComputeManagedResource(activationFile string, exadataInfrastructureId string) (*oci_database.ActivateExadataInfrastructureResponse, error) {
	request := oci_database.ActivateExadataInfrastructureRequest{}

	activationKeyFile, err := ioutil.ReadFile(activationFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open activation key file: %s", err)
	}

	actionKeyFileBase64Encoded := []byte(base64.StdEncoding.EncodeToString(activationKeyFile))
	request.ActivationFile = actionKeyFileBase64Encoded

	request.ExadataInfrastructureId = &exadataInfrastructureId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ActivateExadataInfrastructure(context.Background(), request)
	if err != nil {
		return nil, err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exadataInfrastructure", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return nil, err
		}
	}

	return &response, nil
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) mapToExadataInfrastructureContact(fieldKeyFormat string) (oci_database.ExadataInfrastructureContact, error) {
	result := oci_database.ExadataInfrastructureContact{}

	if email, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email")); ok {
		tmp := email.(string)
		result.Email = &tmp
	}

	if isContactMosValidated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_contact_mos_validated")); ok {
		tmp := isContactMosValidated.(bool)
		result.IsContactMosValidated = &tmp
	}

	if isPrimary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_primary")); ok {
		tmp := isPrimary.(bool)
		result.IsPrimary = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if phoneNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "phone_number")); ok {
		tmp := phoneNumber.(string)
		result.PhoneNumber = &tmp
	}

	return result, nil
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) mapToMaintenanceWindow(fieldKeyFormat string) (oci_database.MaintenanceWindow, error) {
	result := oci_database.MaintenanceWindow{}

	if customActionTimeoutInMins, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_action_timeout_in_mins")); ok {
		tmp := customActionTimeoutInMins.(int)
		result.CustomActionTimeoutInMins = &tmp
	}
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

	if isCustomActionTimeoutEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_custom_action_timeout_enabled")); ok {
		tmp := isCustomActionTimeoutEnabled.(bool)
		result.IsCustomActionTimeoutEnabled = &tmp
	}

	if isMonthlyPatchingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_monthly_patching_enabled")); ok {
		tmp := isMonthlyPatchingEnabled.(bool)
		result.IsMonthlyPatchingEnabled = &tmp
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

	if patchingMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patching_mode")); ok {
		result.PatchingMode = oci_database.MaintenanceWindowPatchingModeEnum(patchingMode.(string))
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

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) mapToDayOfWeek(fieldKeyFormat string) (oci_database.DayOfWeek, error) {
	result := oci_database.DayOfWeek{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.DayOfWeekNameEnum(name.(string))
	}

	return result, nil
}

func (s *DatabaseExadataInfrastructureComputeManagedResourceCrud) mapToMonth(fieldKeyFormat string) (oci_database.Month, error) {
	result := oci_database.Month{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.MonthNameEnum(name.(string))
	}

	return result, nil
}
