// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagedInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubManagedInstance,
		Read:     readOsManagementHubManagedInstance,
		Update:   updateOsManagementHubManagedInstance,
		Delete:   deleteOsManagementHubManagedInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"autonomous_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_data_collection_authorized": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
						"scheduled_job_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"notification_topic_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary_management_station_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_management_station_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bug_updates_available": {
				Type:     schema.TypeInt,
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
			"enhancement_updates_available": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"installed_packages": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"installed_windows_updates": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_managed_by_autonomous_linux": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_management_station": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_reboot_required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ksplice_effective_kernel_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_environment": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_stage": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_instance_group": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"os_family": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_kernel_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"other_updates_available": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"profile": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_job_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"security_updates_available": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"software_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_mandatory_for_autonomous_linux": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"software_source_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_boot": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_checkin": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updates_available": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"work_request_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createOsManagementHubManagedInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubManagedInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()

	return tfresource.ReadResource(sync)
}

func updateOsManagementHubManagedInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOsManagementHubManagedInstance(d *schema.ResourceData, m interface{}) error {
	//sync := &OsManagementHubManagedInstanceResourceCrud{}
	//sync.D = d
	//sync.Client = m.(*client.OracleClients).ManagedInstanceClient()
	//sync.DisableNotFoundRetries = true
	//sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()
	//
	//return tfresource.DeleteResource(d, sync)
	return nil
}

type OsManagementHubManagedInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.ManagedInstanceClient
	Res                    *oci_os_management_hub.ManagedInstance
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_os_management_hub.WorkRequestClient
}

func (s *OsManagementHubManagedInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubManagedInstanceResourceCrud) Create() error {
	request := oci_os_management_hub.UpdateManagedInstanceRequest{}

	if autonomousSettings, ok := s.D.GetOkExists("autonomous_settings"); ok {
		if tmpList := autonomousSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "autonomous_settings", 0)
			tmp, err := s.mapToUpdatableAutonomousSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AutonomousSettings = &tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if notificationTopicId, ok := s.D.GetOkExists("notification_topic_id"); ok {
		tmp := notificationTopicId.(string)
		request.NotificationTopicId = &tmp
	}

	if primaryManagementStationId, ok := s.D.GetOkExists("primary_management_station_id"); ok {
		tmp := primaryManagementStationId.(string)
		request.PrimaryManagementStationId = &tmp
	}

	if secondaryManagementStationId, ok := s.D.GetOkExists("secondary_management_station_id"); ok {
		tmp := secondaryManagementStationId.(string)
		request.SecondaryManagementStationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.UpdateManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstance
	return nil
}

func (s *OsManagementHubManagedInstanceResourceCrud) getManagedInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_os_management_hub.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	managedInstanceId, err := managedInstanceWaitForWorkRequest(workId, "managedinstance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*managedInstanceId)

	return s.Get()
}

func managedInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "os_management_hub", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_os_management_hub.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func managedInstanceWaitForWorkRequest(wId *string, entityType string, action oci_os_management_hub.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_os_management_hub.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "os_management_hub")
	retryPolicy.ShouldRetryOperation = managedInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_os_management_hub.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_os_management_hub.OperationStatusInProgress),
			string(oci_os_management_hub.OperationStatusAccepted),
			string(oci_os_management_hub.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_os_management_hub.OperationStatusSucceeded),
			string(oci_os_management_hub.OperationStatusFailed),
			string(oci_os_management_hub.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_os_management_hub.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(string(res.EntityType)), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_os_management_hub.OperationStatusFailed || response.Status == oci_os_management_hub.OperationStatusCanceled {
		return nil, getErrorFromOsManagementHubManagedInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOsManagementHubManagedInstanceWorkRequest(client *oci_os_management_hub.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_os_management_hub.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_os_management_hub.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *OsManagementHubManagedInstanceResourceCrud) Get() error {
	request := oci_os_management_hub.GetManagedInstanceRequest{}

	tmp := s.D.Id()
	request.ManagedInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.GetManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstance
	return nil
}

func (s *OsManagementHubManagedInstanceResourceCrud) Update() error {
	request := oci_os_management_hub.UpdateManagedInstanceRequest{}

	if autonomousSettings, ok := s.D.GetOkExists("autonomous_settings"); ok {
		if tmpList := autonomousSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "autonomous_settings", 0)
			tmp, err := s.mapToUpdatableAutonomousSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AutonomousSettings = &tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.ManagedInstanceId = &tmp

	if notificationTopicId, ok := s.D.GetOkExists("notification_topic_id"); ok {
		tmp := notificationTopicId.(string)
		request.NotificationTopicId = &tmp
	}

	if primaryManagementStationId, ok := s.D.GetOkExists("primary_management_station_id"); ok {
		tmp := primaryManagementStationId.(string)
		request.PrimaryManagementStationId = &tmp
	}

	if secondaryManagementStationId, ok := s.D.GetOkExists("secondary_management_station_id"); ok {
		tmp := secondaryManagementStationId.(string)
		request.SecondaryManagementStationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.UpdateManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstance
	return nil
}

func (s *OsManagementHubManagedInstanceResourceCrud) Delete() error {
	//request := oci_os_management_hub.DeleteManagedInstanceRequest{}
	//
	//tmp := s.D.Id()
	//request.ManagedInstanceId = &tmp
	//
	//request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")
	//
	//response, err := s.Client.DeleteManagedInstance(context.Background(), request)
	//if err != nil {
	//	return err
	//}
	//
	//workId := response.OpcWorkRequestId
	//// Wait until it finishes
	//_, delWorkRequestErr := managedInstanceWaitForWorkRequest(workId, "managedinstance",
	//	oci_os_management_hub.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	//return delWorkRequestErr
	return nil
}

func (s *OsManagementHubManagedInstanceResourceCrud) SetData() error {
	s.D.Set("architecture", s.Res.Architecture)

	if s.Res.AutonomousSettings != nil {
		s.D.Set("autonomous_settings", []interface{}{AutonomousSettingsToMap(s.Res.AutonomousSettings)})
	} else {
		s.D.Set("autonomous_settings", nil)
	}

	if s.Res.BugUpdatesAvailable != nil {
		s.D.Set("bug_updates_available", *s.Res.BugUpdatesAvailable)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EnhancementUpdatesAvailable != nil {
		s.D.Set("enhancement_updates_available", *s.Res.EnhancementUpdatesAvailable)
	}

	if s.Res.InstalledPackages != nil {
		s.D.Set("installed_packages", *s.Res.InstalledPackages)
	}

	if s.Res.InstalledWindowsUpdates != nil {
		s.D.Set("installed_windows_updates", *s.Res.InstalledWindowsUpdates)
	}

	if s.Res.IsManagedByAutonomousLinux != nil {
		s.D.Set("is_managed_by_autonomous_linux", *s.Res.IsManagedByAutonomousLinux)
	}

	if s.Res.IsManagementStation != nil {
		s.D.Set("is_management_station", *s.Res.IsManagementStation)
	}

	if s.Res.IsRebootRequired != nil {
		s.D.Set("is_reboot_required", *s.Res.IsRebootRequired)
	}

	if s.Res.KspliceEffectiveKernelVersion != nil {
		s.D.Set("ksplice_effective_kernel_version", *s.Res.KspliceEffectiveKernelVersion)
	}

	if s.Res.LifecycleEnvironment != nil {
		s.D.Set("lifecycle_environment", []interface{}{IdToMap(s.Res.LifecycleEnvironment)})
	} else {
		s.D.Set("lifecycle_environment", nil)
	}

	if s.Res.LifecycleStage != nil {
		s.D.Set("lifecycle_stage", []interface{}{IdToMap(s.Res.LifecycleStage)})
	} else {
		s.D.Set("lifecycle_stage", nil)
	}

	s.D.Set("location", s.Res.Location)

	if s.Res.ManagedInstanceGroup != nil {
		s.D.Set("managed_instance_group", []interface{}{IdToMap(s.Res.ManagedInstanceGroup)})
	} else {
		s.D.Set("managed_instance_group", nil)
	}

	if s.Res.NotificationTopicId != nil {
		s.D.Set("notification_topic_id", *s.Res.NotificationTopicId)
	}

	s.D.Set("os_family", s.Res.OsFamily)

	if s.Res.OsKernelVersion != nil {
		s.D.Set("os_kernel_version", *s.Res.OsKernelVersion)
	}

	if s.Res.OsName != nil {
		s.D.Set("os_name", *s.Res.OsName)
	}

	if s.Res.OsVersion != nil {
		s.D.Set("os_version", *s.Res.OsVersion)
	}

	if s.Res.OtherUpdatesAvailable != nil {
		s.D.Set("other_updates_available", *s.Res.OtherUpdatesAvailable)
	}

	if s.Res.PrimaryManagementStationId != nil {
		s.D.Set("primary_management_station_id", *s.Res.PrimaryManagementStationId)
	}

	if s.Res.Profile != nil {
		s.D.Set("profile", *s.Res.Profile)
	}

	if s.Res.ScheduledJobCount != nil {
		s.D.Set("scheduled_job_count", *s.Res.ScheduledJobCount)
	}

	if s.Res.SecondaryManagementStationId != nil {
		s.D.Set("secondary_management_station_id", *s.Res.SecondaryManagementStationId)
	}

	if s.Res.SecurityUpdatesAvailable != nil {
		s.D.Set("security_updates_available", *s.Res.SecurityUpdatesAvailable)
	}

	softwareSources := []interface{}{}
	for _, item := range s.Res.SoftwareSources {
		softwareSources = append(softwareSources, SoftwareSourceDetailsToMap(item))
	}
	s.D.Set("software_sources", softwareSources)

	s.D.Set("status", s.Res.Status)

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastBoot != nil {
		s.D.Set("time_last_boot", s.Res.TimeLastBoot.String())
	}

	if s.Res.TimeLastCheckin != nil {
		s.D.Set("time_last_checkin", s.Res.TimeLastCheckin.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UpdatesAvailable != nil {
		s.D.Set("updates_available", *s.Res.UpdatesAvailable)
	}

	if s.Res.WorkRequestCount != nil {
		s.D.Set("work_request_count", *s.Res.WorkRequestCount)
	}

	return nil
}

func AutonomousSettingsToMap(obj *oci_os_management_hub.AutonomousSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDataCollectionAuthorized != nil {
		result["is_data_collection_authorized"] = bool(*obj.IsDataCollectionAuthorized)
	}

	if obj.ScheduledJobId != nil {
		result["scheduled_job_id"] = string(*obj.ScheduledJobId)
	}

	return result
}

func IdToMap(obj *oci_os_management_hub.Id) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func ManagedInstanceSummaryToMap(obj oci_os_management_hub.ManagedInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["architecture"] = string(obj.Architecture)

	if obj.AutonomousSettings != nil {
		result["autonomous_settings"] = []interface{}{AutonomousSettingsToMap(obj.AutonomousSettings)}
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsManagedByAutonomousLinux != nil {
		result["is_managed_by_autonomous_linux"] = bool(*obj.IsManagedByAutonomousLinux)
	}

	if obj.IsManagementStation != nil {
		result["is_management_station"] = bool(*obj.IsManagementStation)
	}

	if obj.IsRebootRequired != nil {
		result["is_reboot_required"] = bool(*obj.IsRebootRequired)
	}

	if obj.LifecycleEnvironment != nil {
		result["lifecycle_environment"] = []interface{}{IdToMap(obj.LifecycleEnvironment)}
	}

	if obj.LifecycleStage != nil {
		result["lifecycle_stage"] = []interface{}{IdToMap(obj.LifecycleStage)}
	}

	result["location"] = string(obj.Location)

	if obj.ManagedInstanceGroup != nil {
		result["managed_instance_group"] = []interface{}{IdToMap(obj.ManagedInstanceGroup)}
	}

	if obj.NotificationTopicId != nil {
		result["notification_topic_id"] = string(*obj.NotificationTopicId)
	}

	result["os_family"] = string(obj.OsFamily)

	result["status"] = string(obj.Status)

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	if obj.UpdatesAvailable != nil {
		result["updates_available"] = int(*obj.UpdatesAvailable)
	}

	return result
}

func (s *OsManagementHubManagedInstanceResourceCrud) mapToUpdatableAutonomousSettings(fieldKeyFormat string) (oci_os_management_hub.UpdatableAutonomousSettings, error) {
	result := oci_os_management_hub.UpdatableAutonomousSettings{}

	if isDataCollectionAuthorized, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_data_collection_authorized")); ok {
		tmp := isDataCollectionAuthorized.(bool)
		result.IsDataCollectionAuthorized = &tmp
	}

	return result, nil
}
