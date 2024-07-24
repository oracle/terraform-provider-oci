// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAuditTrailManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeAuditTrailManagement,
		Read:     readDataSafeAuditTrailManagement,
		Update:   updateDataSafeAuditTrailManagement,
		Delete:   deleteDataSafeAuditTrailManagement,
		Schema: map[string]*schema.Schema{

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
			"is_auto_purge_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_data_safe.AuditTrailLifecycleStateInactive),
					string(oci_data_safe.AuditTrailLifecycleStateActive),
				}, true),
			},
			"resume_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"stop_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"start_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"audit_collection_start_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trail_location": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed

			"audit_profile_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
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
			"work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeAuditTrailManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditTrailManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	err := sync.GetAuditTrailWorkReq()
	err1 := sync.Get()
	if err != nil {
		return err
	}
	if err1 != nil {
		return err1
	}
	return updateDataSafeAuditTrailManagement(d, m)
}

func readDataSafeAuditTrailManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditTrailManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeAuditTrailManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditTrailManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	if startTrigger, ok := sync.D.GetOkExists("start_trigger"); ok {
		if startTrigger == true {
			if err := sync.StartAuditTrail(); err != nil {
				return err
			}
			sync.D.Set("state", oci_data_safe.AuditTrailLifecycleStateActive)
		}
	}
	if resumeTrigger, ok := sync.D.GetOkExists("resume_trigger"); ok {
		if resumeTrigger == true {
			err := sync.ResumeAuditTrail()
			if err != nil {
				return err
			}
		}
	}
	if stopTrigger, ok := sync.D.GetOkExists("stop_trigger"); ok {
		if stopTrigger == true {
			if err := sync.StopAuditTrail(); err != nil {
				return err
			}
			sync.D.Set("state", oci_data_safe.AuditTrailLifecycleStateInactive)
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteDataSafeAuditTrailManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditTrailManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeAuditTrailManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.AuditTrail
	DisableNotFoundRetries bool
}

func (s *DataSafeAuditTrailManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeAuditTrailManagementResourceCrud) StartAuditTrail() error {
	request := oci_data_safe.StartAuditTrailRequest{}
	idTmp := s.D.Id()
	request.AuditTrailId = &idTmp

	if auditCollectionStartTime, ok := s.D.GetOkExists("audit_collection_start_time"); ok {
		tmp, err := time.Parse(time.RFC3339, auditCollectionStartTime.(string))
		if err != nil {
			return err
		}
		request.AuditCollectionStartTime = &oci_common.SDKTime{Time: tmp}
	}

	if isAutoPurgeEnabled, ok := s.D.GetOkExists("is_auto_purge_enabled"); ok {
		tmp := isAutoPurgeEnabled.(bool)
		request.IsAutoPurgeEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.StartAuditTrail(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_data_safe.AuditTrailLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditTrailManagementResourceCrud) StopAuditTrail() error {
	request := oci_data_safe.StopAuditTrailRequest{}

	idTmp := s.D.Id()
	request.AuditTrailId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.StopAuditTrail(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_data_safe.AuditTrailLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditTrailManagementResourceCrud) ResumeAuditTrail() error {
	request := oci_data_safe.ResumeAuditTrailRequest{}

	idTmp := s.D.Id()
	request.AuditTrailId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ResumeAuditTrail(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("resume_trigger")
	s.D.Set("resume_trigger", val)

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_data_safe.AuditTrailLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditTrailManagementResourceCrud) GetAuditTrailWorkReq() error {
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
		_ = s.getAuditProfileFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate))
		return s.GetAuditTrailList()
	} else {
		return s.GetAuditTrailList()
	}
}

func (s *DataSafeAuditTrailManagementResourceCrud) GetAuditTrailList() error {
	request := oci_data_safe.ListAuditTrailsRequest{}
	var auditTrail = new(oci_data_safe.AuditTrail)
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}
	var trailLocation, _ = s.D.GetOkExists("trail_location")

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListAuditTrails(context.Background(), request)
	if err != nil {
		return err
	}
	if response.AuditTrailCollection.Items != nil && len(response.AuditTrailCollection.Items) > 0 {
		for i := 0; i < len(response.AuditTrailCollection.Items); i++ {
			tmp1 := &response.AuditTrailCollection.Items[i]
			if *tmp1.TrailLocation == trailLocation {
				auditTrail.Id = tmp1.Id
			}
		}
	}

	if auditTrail.Id == nil {
		return nil
	}

	s.D.SetId(*auditTrail.Id)
	return nil
}

func (s *DataSafeAuditTrailManagementResourceCrud) Get() error {
	request := oci_data_safe.GetAuditTrailRequest{}

	tmp := s.D.Id()
	request.AuditTrailId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetAuditTrail(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuditTrail
	return nil
}

func (s *DataSafeAuditTrailManagementResourceCrud) Update() error {
	request := oci_data_safe.UpdateAuditTrailRequest{}

	tmp := s.D.Id()
	request.AuditTrailId = &tmp

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

	if isAutoPurgeEnabled, ok := s.D.GetOkExists("is_auto_purge_enabled"); ok {
		tmp := isAutoPurgeEnabled.(bool)
		request.IsAutoPurgeEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAuditTrail(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditTrailFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditTrailManagementResourceCrud) Delete() error {
	request := oci_data_safe.DeleteAuditTrailRequest{}

	tmp := s.D.Id()
	request.AuditTrailId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteAuditTrail(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := auditTrailWaitForWorkRequest(workId, "audittrail",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeAuditTrailManagementResourceCrud) SetData() error {
	if s.Res.AuditCollectionStartTime != nil {
		s.D.Set("audit_collection_start_time", s.Res.AuditCollectionStartTime.String())
	}

	if s.Res.AuditProfileId != nil {
		s.D.Set("audit_profile_id", *s.Res.AuditProfileId)
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

	if s.Res.IsAutoPurgeEnabled != nil {
		s.D.Set("is_auto_purge_enabled", *s.Res.IsAutoPurgeEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastCollected != nil {
		s.D.Set("time_last_collected", s.Res.TimeLastCollected.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TrailLocation != nil {
		s.D.Set("trail_location", *s.Res.TrailLocation)
	}

	if s.Res.WorkRequestId != nil {
		s.D.Set("work_request_id", *s.Res.WorkRequestId)
	}

	return nil
}

func (s *DataSafeAuditTrailManagementResourceCrud) getAuditTrailFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	auditTrailId, err := auditTrailWaitForWorkRequest(workId, "audittrail",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*auditTrailId)

	return s.Get()
}

func (s *DataSafeAuditTrailManagementResourceCrud) getAuditProfileFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := auditProfileWaitForWorkRequest(workId, "auditprofile",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}
