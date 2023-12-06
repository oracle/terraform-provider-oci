// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
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

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
			"is_paid_usage_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"change_retention_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"offline_months": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"online_months": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_override_global_retention_setting": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"audit_collected_volume": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
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

	err := sync.GetAuditProfileWorkReq()
	err1 := sync.Get()
	if err != nil {
		return err
	}
	if err1 != nil {
		return err1
	}
	return updateDataSafeAuditProfileManagement(d, m)
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

	if _, ok := sync.D.GetOkExists("change_retention_trigger"); ok {
		err := sync.ChangeAuditRetention()

		if err != nil {
			return err
		}
	}

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeAuditProfileManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeAuditProfileManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.AuditProfile
	DisableNotFoundRetries bool
}

func (s *DataSafeAuditProfileManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeAuditProfileManagementResourceCrud) ChangeAuditRetention() error {
	request := oci_data_safe.ChangeRetentionRequest{}

	idTmp := s.D.Id()
	request.AuditProfileId = &idTmp

	if isOverrideGlobalRetentionSetting, ok := s.D.GetOkExists("is_override_global_retention_setting"); ok {
		tmp := isOverrideGlobalRetentionSetting.(bool)
		request.IsOverrideGlobalRetentionSetting = &tmp
	}

	if offlineMonths, ok := s.D.GetOkExists("offline_months"); ok {
		tmp := offlineMonths.(int)
		request.OfflineMonths = &tmp
	}

	if onlineMonths, ok := s.D.GetOkExists("online_months"); ok {
		tmp := onlineMonths.(int)
		request.OnlineMonths = &tmp
	}

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
	if listWorkRequestsResponse.Items != nil && len(listWorkRequestsResponse.Items) > 0 {
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

func (s *DataSafeAuditProfileManagementResourceCrud) SetData() error {
	if s.Res.AuditCollectedVolume != nil {
		s.D.Set("audit_collected_volume", strconv.FormatInt(*s.Res.AuditCollectedVolume, 10))
	}

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

	if s.Res.OnlineMonths != nil {
		s.D.Set("online_months", *s.Res.OnlineMonths)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
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
