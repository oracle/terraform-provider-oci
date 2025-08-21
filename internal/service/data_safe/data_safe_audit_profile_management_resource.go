// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAuditProfileManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeAuditProfileManagement,
		Read:     readDataSafeAuditProfileManagement,
		Update:   updateDataSafeAuditProfileManagement,
		Delete:   deleteDataSafeAuditProfileManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"is_override_global_paid_usage": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_paid_usage_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"offline_months": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"online_months": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"change_retention_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_override_global_retention_setting": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},

			// Computed
			"audit_collected_volume": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"audit_trails": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						// Computed
						"audit_collection_start_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"audit_profile_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"can_update_last_archive_time_on_target": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_auto_purge_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"peer_target_database_key": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"purge_job_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"purge_job_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"purge_job_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_collected": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trail_location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trail_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"work_request_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"offline_months_source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"online_months_source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"paid_usage_source": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeAuditProfileManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}
	return nil
}

func readDataSafeAuditProfileManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeAuditProfileManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	offlineMonths, _ := sync.D.GetOk("offline_months")
	onlineMonths, _ := sync.D.GetOk("online_months")
	isOverrideGlobalRetentionSetting, _ := sync.D.GetOk("is_override_global_retention_setting")
	if sync.D.Get("target_type") == string(oci_data_safe.AuditProfileTargetTypeTargetDatabase) && sync.D.Id() == "" {
		id, err := sync.fetchAuditProfileIdByTargetId()
		if err != nil {
			return fmt.Errorf("failed to fetch audit profile ID: %v", err)
		}
		sync.D.SetId(id)
	}
	if _, ok := sync.D.GetOkExists("change_retention_trigger"); ok && sync.D.HasChange("change_retention_trigger") {
		oldRaw, newRaw := sync.D.GetChange("change_retention_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ChangeRetention(offlineMonths.(int), onlineMonths.(int), isOverrideGlobalRetentionSetting.(bool))

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("change_retention_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteDataSafeAuditProfileManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true
	if sync.D.Get("target_type") == string(oci_data_safe.AuditProfileTargetTypeTargetDatabase) {
		return nil
	}
	return tfresource.DeleteResource(d, sync)
}

type DataSafeAuditProfileManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.AuditProfile
	DisableNotFoundRetries bool
}

func (s *DataSafeAuditProfileManagementResourceCrud) ID() string {
	// return *s.Res.Id
	if s.Res != nil && s.Res.Id != nil && *s.Res.Id != "" {
		return *s.Res.Id
	}
	return s.D.Id()
}

func (s *DataSafeAuditProfileManagementResourceCrud) CreatedPending() []string {
	switch s.D.Get("target_type") {
	case string(oci_data_safe.AuditProfileTargetTypeTargetDatabase):
		return []string{string(oci_data_safe.AuditProfileLifecycleStateActive), string(oci_data_safe.AuditProfileLifecycleStateNeedsAttention)}
	default:
		return []string{string(oci_data_safe.AuditProfileLifecycleStateCreating)}
	}
}

func (s *DataSafeAuditProfileManagementResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.AuditProfileLifecycleStateActive),
		string(oci_data_safe.AuditProfileLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeAuditProfileManagementResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.AuditProfileLifecycleStateDeleting),
	}
}

func (s *DataSafeAuditProfileManagementResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.AuditProfileLifecycleStateDeleted),
	}
}

func (s *DataSafeAuditProfileManagementResourceCrud) Create() error {
	if s.D.Get("target_type") == string(oci_data_safe.AuditProfileTargetTypeTargetDatabase) {

		err := s.GetAuditProfileWorkReq()
		if err != nil {
			return fmt.Errorf("failed to fetch audit profile ID: %v", err)
		}
		return nil
	}
	request := oci_data_safe.CreateAuditProfileRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isOverrideGlobalPaidUsage, ok := s.D.GetOkExists("is_override_global_paid_usage"); ok {
		tmp := isOverrideGlobalPaidUsage.(bool)
		request.IsOverrideGlobalPaidUsage = &tmp
	}

	if isPaidUsageEnabled, ok := s.D.GetOkExists("is_paid_usage_enabled"); ok {
		tmp := isPaidUsageEnabled.(bool)
		request.IsPaidUsageEnabled = &tmp
	}

	if offlineMonths, ok := s.D.GetOkExists("offline_months"); ok {
		tmp := offlineMonths.(int)
		request.OfflineMonths = &tmp
	}

	if onlineMonths, ok := s.D.GetOkExists("online_months"); ok {
		tmp := onlineMonths.(int)
		request.OnlineMonths = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if targetType, ok := s.D.GetOkExists("target_type"); ok {
		request.TargetType = oci_data_safe.AuditProfileTargetTypeEnum(targetType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateAuditProfile(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getAuditProfileFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeAuditProfileManagementResourceCrud) GetAuditProfileWorkReq() error {
	listWorkRequestsRequest := oci_data_safe.ListWorkRequestsRequest{SortBy: oci_data_safe.ListWorkRequestsSortByEnum("ACCEPTEDTIME"), SortOrder: oci_data_safe.ListWorkRequestsSortOrderEnum("DESC")}
	var workId *string
	tmp := "CREATE_AUDIT_PROFILE"
	listWorkRequestsRequest.OperationType = &tmp

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		listWorkRequestsRequest.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		listWorkRequestsRequest.TargetDatabaseId = &tmp
	}

	listWorkRequestsResponse, err := s.Client.ListWorkRequests(context.Background(), listWorkRequestsRequest)
	if len(listWorkRequestsResponse.Items) > 0 {
		var tmp1 = &listWorkRequestsResponse.Items[0]
		workId = tmp1.Id
	}

	if err != nil {
		return err
	}

	if workId != nil {
		return s.getAuditProfileFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate))
	} else {
		return s.GetAuditProfileList()
	}
}

func (s *DataSafeAuditProfileManagementResourceCrud) GetAuditProfileList() error {
	request := oci_data_safe.ListAuditProfilesRequest{}
	var auditProfile = new(oci_data_safe.AuditProfile)
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListAuditProfiles(context.Background(), request)
	if err != nil {
		return err
	}
	if response.AuditProfileCollection.Items != nil && len(response.AuditProfileCollection.Items) > 0 {
		tmp1 := &response.AuditProfileCollection.Items[0]
		auditProfile.Id = tmp1.Id
		s.D.Set("is_override_global_retention_setting", tmp1.IsOverrideGlobalRetentionSetting)
	}

	if auditProfile.Id == nil {
		return nil
	}

	s.D.SetId(*auditProfile.Id)
	return nil
}

func (s *DataSafeAuditProfileManagementResourceCrud) Get() error {
	request := oci_data_safe.GetAuditProfileRequest{}

	tmp := s.D.Id()
	request.AuditProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetAuditProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuditProfile
	return nil
}

func (s *DataSafeAuditProfileManagementResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateAuditProfileRequest{}

	tmp := s.D.Id()
	request.AuditProfileId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isPaidUsageEnabled, ok := s.D.GetOkExists("is_paid_usage_enabled"); ok {
		tmp := isPaidUsageEnabled.(bool)
		request.IsPaidUsageEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAuditProfile(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditProfileFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditProfileManagementResourceCrud) Delete() error {
	if s.D.Get("target_type").(string) == string(oci_data_safe.AuditProfileTargetTypeTargetDatabase) {
		return fmt.Errorf("audit profile for a target database cannot be deleted")
	}
	request := oci_data_safe.DeleteAuditProfileRequest{}

	tmp := s.D.Id()
	request.AuditProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteAuditProfile(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := auditProfileWaitForWorkRequest(workId, "auditprofile",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeAuditProfileManagementResourceCrud) SetData() error {
	if s.Res.AuditCollectedVolume != nil {
		s.D.Set("audit_collected_volume", strconv.FormatInt(*s.Res.AuditCollectedVolume, 10))
	}

	auditTrails := []interface{}{}
	for _, item := range s.Res.AuditTrails {
		auditTrails = append(auditTrails, AuditTrailToMap(item))
	}
	s.D.Set("audit_trails", auditTrails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsOverrideGlobalRetentionSetting != nil {
		s.D.Set("is_override_global_retention_setting", *s.Res.IsOverrideGlobalRetentionSetting)
	}

	if s.Res.IsPaidUsageEnabled != nil {
		s.D.Set("is_paid_usage_enabled", *s.Res.IsPaidUsageEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OfflineMonths != nil {
		s.D.Set("offline_months", *s.Res.OfflineMonths)
	}

	if s.Res.OfflineMonthsSource != nil {
		s.D.Set("offline_months_source", *s.Res.OfflineMonthsSource)
	}

	if s.Res.OnlineMonths != nil {
		s.D.Set("online_months", *s.Res.OnlineMonths)
	}

	if s.Res.OnlineMonthsSource != nil {
		s.D.Set("online_months_source", *s.Res.OnlineMonthsSource)
	}

	if s.Res.PaidUsageSource != nil {
		s.D.Set("paid_usage_source", *s.Res.PaidUsageSource)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	s.D.Set("target_type", s.Res.TargetType)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.IsOverrideGlobalPaidUsage != nil {
		s.D.Set("is_override_global_paid_usage", *s.Res.IsOverrideGlobalPaidUsage)
	}

	return nil
}

func (s *DataSafeAuditProfileManagementResourceCrud) getAuditProfileFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	auditProfileId, err := auditProfileWaitForWorkRequest(workId, "auditprofile",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*auditProfileId)
	return s.Get()
}

func (s *DataSafeAuditProfileManagementResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeAuditProfileCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AuditProfileId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeAuditProfileCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditProfileFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))

}

func (s *DataSafeAuditProfileManagementResourceCrud) fetchAuditProfileIdByTargetId() (string, error) {
	request := oci_data_safe.ListAuditProfilesRequest{}
	var auditProfile = new(oci_data_safe.AuditProfile)
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListAuditProfiles(context.Background(), request)
	if err != nil {
		return "", err
	}
	if len(response.AuditProfileCollection.Items) > 0 {
		tmp1 := &response.AuditProfileCollection.Items[0]
		auditProfile.Id = tmp1.Id
		s.D.Set("is_override_global_retention_setting", tmp1.IsOverrideGlobalRetentionSetting)
	}

	if auditProfile.Id == nil {
		return "", fmt.Errorf("no audit profile found matching the criteria")
	}
	s.D.SetId(*auditProfile.Id)
	return *auditProfile.Id, nil
}

func (s *DataSafeAuditProfileManagementResourceCrud) ChangeRetention(offlineMonths int, onlineMonths int, isOverrideGlobalRetentionSetting bool) error {
	request := oci_data_safe.ChangeRetentionRequest{}

	idTmp := s.D.Id()
	request.AuditProfileId = &idTmp
	request.OfflineMonths = &offlineMonths
	request.OnlineMonths = &onlineMonths
	request.IsOverrideGlobalRetentionSetting = &isOverrideGlobalRetentionSetting
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")
	response, err := s.Client.ChangeRetention(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("change_retention_trigger")
	s.D.Set("change_retention_trigger", val)

	workId := response.OpcWorkRequestId
	return s.getAuditProfileFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))

}
