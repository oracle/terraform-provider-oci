// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apiaccesscontrol

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_apiaccesscontrol "github.com/oracle/oci-go-sdk/v65/apiaccesscontrol"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApiaccesscontrolPrivilegedApiRequestResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApiaccesscontrolPrivilegedApiRequest,
		Read:     readApiaccesscontrolPrivilegedApiRequest,
		Delete:   deleteApiaccesscontrolPrivilegedApiRequest,
		Schema: map[string]*schema.Schema{
			// Required
			"privileged_operation_list": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"api_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"attribute_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"reason_summary": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"duration_in_hrs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"notification_topic_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"reason_detail": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sub_resource_name_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ticket_numbers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_requested_for_future_access": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},

			// Computed
			"approver_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"approval_action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"approval_comment": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"approver_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_approved_for_access": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_authorization": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"closure_comment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"number_of_approvers_required": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"privileged_api_control_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"privileged_api_control_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"requested_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resource_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_details": {
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

func createApiaccesscontrolPrivilegedApiRequest(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolPrivilegedApiRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivilegedApiRequestsClient()
	sync.WorkRequestClient = m.(*client.OracleClients).ApiaccesscontrolPrivilegedApiWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readApiaccesscontrolPrivilegedApiRequest(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolPrivilegedApiRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivilegedApiRequestsClient()

	return tfresource.ReadResource(sync)
}

func deleteApiaccesscontrolPrivilegedApiRequest(d *schema.ResourceData, m interface{}) error {
	return nil
}

type ApiaccesscontrolPrivilegedApiRequestResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apiaccesscontrol.PrivilegedApiRequestsClient
	Res                    *oci_apiaccesscontrol.PrivilegedApiRequest
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_apiaccesscontrol.PrivilegedApiWorkRequestClient
}

func (s *ApiaccesscontrolPrivilegedApiRequestResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ApiaccesscontrolPrivilegedApiRequestResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *ApiaccesscontrolPrivilegedApiRequestResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_apiaccesscontrol.PrivilegedApiRequestLifecycleStateInProgress),
		string(oci_apiaccesscontrol.PrivilegedApiRequestLifecycleStateAccepted),
		string(oci_apiaccesscontrol.PrivilegedApiRequestLifecycleStateWaiting),
		string(oci_apiaccesscontrol.PrivilegedApiRequestLifecycleStateSucceeded),
	}
}

func (s *ApiaccesscontrolPrivilegedApiRequestResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *ApiaccesscontrolPrivilegedApiRequestResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *ApiaccesscontrolPrivilegedApiRequestResourceCrud) Create() error {
	request := oci_apiaccesscontrol.CreatePrivilegedApiRequestRequest{}

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

	if durationInHrs, ok := s.D.GetOkExists("duration_in_hrs"); ok {
		tmp := durationInHrs.(int)
		request.DurationInHrs = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if notificationTopicId, ok := s.D.GetOkExists("notification_topic_id"); ok {
		tmp := notificationTopicId.(string)
		request.NotificationTopicId = &tmp
	}

	if privilegedOperationList, ok := s.D.GetOkExists("privileged_operation_list"); ok {
		interfaces := privilegedOperationList.([]interface{})
		tmp := make([]oci_apiaccesscontrol.PrivilegedApiRequestOperationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "privileged_operation_list", stateDataIndex)
			converted, err := s.mapToPrivilegedApiRequestOperationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("privileged_operation_list") {
			request.PrivilegedOperationList = tmp
		}
	}

	if reasonDetail, ok := s.D.GetOkExists("reason_detail"); ok {
		tmp := reasonDetail.(string)
		request.ReasonDetail = &tmp
	}

	if reasonSummary, ok := s.D.GetOkExists("reason_summary"); ok {
		tmp := reasonSummary.(string)
		request.ReasonSummary = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_apiaccesscontrol.PrivilegedApiRequestSeverityEnum(severity.(string))
	}

	if subResourceNameList, ok := s.D.GetOkExists("sub_resource_name_list"); ok {
		interfaces := subResourceNameList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("sub_resource_name_list") {
			request.SubResourceNameList = tmp
		}
	}

	if ticketNumbers, ok := s.D.GetOkExists("ticket_numbers"); ok {
		interfaces := ticketNumbers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ticket_numbers") {
			request.TicketNumbers = tmp
		}
	}

	if timeRequestedForFutureAccess, ok := s.D.GetOkExists("time_requested_for_future_access"); ok {
		tmp, err := time.Parse(time.RFC3339, timeRequestedForFutureAccess.(string))
		if err != nil {
			return err
		}
		request.TimeRequestedForFutureAccess = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apiaccesscontrol")

	response, err := s.Client.CreatePrivilegedApiRequest(context.Background(), request)
	if err != nil {
		return err
	}

	//workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.Get()
	//return s.getPrivilegedApiRequestFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apiaccesscontrol"), oci_apiaccesscontrol.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ApiaccesscontrolPrivilegedApiRequestResourceCrud) getPrivilegedApiRequestFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_apiaccesscontrol.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	privilegedApiRequestId, err := privilegedApiRequestWaitForWorkRequest(workId, "privilegedapirequest",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, privilegedApiRequestId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(context.Background(),
			oci_apiaccesscontrol.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*privilegedApiRequestId)

	return s.Get()
}

func privilegedApiRequestWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "apiaccesscontrol", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_apiaccesscontrol.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func privilegedApiRequestWaitForWorkRequest(wId *string, entityType string, action oci_apiaccesscontrol.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_apiaccesscontrol.PrivilegedApiWorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "apiaccesscontrol")
	retryPolicy.ShouldRetryOperation = privilegedApiRequestWorkRequestShouldRetryFunc(timeout)

	response := oci_apiaccesscontrol.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_apiaccesscontrol.OperationStatusInProgress),
			string(oci_apiaccesscontrol.OperationStatusAccepted),
			string(oci_apiaccesscontrol.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_apiaccesscontrol.OperationStatusSucceeded),
			string(oci_apiaccesscontrol.OperationStatusFailed),
			string(oci_apiaccesscontrol.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_apiaccesscontrol.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_apiaccesscontrol.OperationStatusFailed || response.Status == oci_apiaccesscontrol.OperationStatusCanceled {
		return nil, getErrorFromApiaccesscontrolPrivilegedApiRequestWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromApiaccesscontrolPrivilegedApiRequestWorkRequest(client *oci_apiaccesscontrol.PrivilegedApiWorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_apiaccesscontrol.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_apiaccesscontrol.ListWorkRequestErrorsRequest{
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

func (s *ApiaccesscontrolPrivilegedApiRequestResourceCrud) Get() error {
	request := oci_apiaccesscontrol.GetPrivilegedApiRequestRequest{}

	tmp := s.D.Id()
	request.PrivilegedApiRequestId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apiaccesscontrol")

	response, err := s.Client.GetPrivilegedApiRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivilegedApiRequest
	return nil
}

func (s *ApiaccesscontrolPrivilegedApiRequestResourceCrud) SetData() error {
	approverDetails := []interface{}{}
	for _, item := range s.Res.ApproverDetails {
		approverDetails = append(approverDetails, ApproverDetailToMap(item))
	}
	s.D.Set("approver_details", approverDetails)

	if s.Res.ClosureComment != nil {
		s.D.Set("closure_comment", *s.Res.ClosureComment)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DurationInHrs != nil {
		s.D.Set("duration_in_hrs", *s.Res.DurationInHrs)
	}

	if s.Res.EntityType != nil {
		s.D.Set("entity_type", *s.Res.EntityType)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NotificationTopicId != nil {
		s.D.Set("notification_topic_id", *s.Res.NotificationTopicId)
	}

	if s.Res.NumberOfApproversRequired != nil {
		s.D.Set("number_of_approvers_required", *s.Res.NumberOfApproversRequired)
	}

	if s.Res.PrivilegedApiControlId != nil {
		s.D.Set("privileged_api_control_id", *s.Res.PrivilegedApiControlId)
	}

	if s.Res.PrivilegedApiControlName != nil {
		s.D.Set("privileged_api_control_name", *s.Res.PrivilegedApiControlName)
	}

	privilegedOperationList := []interface{}{}
	for _, item := range s.Res.PrivilegedOperationList {
		privilegedOperationList = append(privilegedOperationList, PrivilegedApiRequestOperationDetailsToMap(item))
	}
	s.D.Set("privileged_operation_list", privilegedOperationList)

	if s.Res.ReasonDetail != nil {
		s.D.Set("reason_detail", *s.Res.ReasonDetail)
	}

	if s.Res.ReasonSummary != nil {
		s.D.Set("reason_summary", *s.Res.ReasonSummary)
	}

	if s.Res.RequestId != nil {
		s.D.Set("request_id", *s.Res.RequestId)
	}

	s.D.Set("requested_by", s.Res.RequestedBy)

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", *s.Res.ResourceType)
	}

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.State)

	if s.Res.StateDetails != nil {
		s.D.Set("state_details", *s.Res.StateDetails)
	}

	s.D.Set("sub_resource_name_list", s.Res.SubResourceNameList)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("ticket_numbers", s.Res.TicketNumbers)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeRequestedForFutureAccess != nil {
		s.D.Set("time_requested_for_future_access", s.Res.TimeRequestedForFutureAccess.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ApproverDetailToMap(obj oci_apiaccesscontrol.ApproverDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApprovalAction != nil {
		result["approval_action"] = string(*obj.ApprovalAction)
	}

	if obj.ApprovalComment != nil {
		result["approval_comment"] = string(*obj.ApprovalComment)
	}

	if obj.ApproverId != nil {
		result["approver_id"] = string(*obj.ApproverId)
	}

	if obj.TimeApprovedForAccess != nil {
		result["time_approved_for_access"] = obj.TimeApprovedForAccess.String()
	}

	if obj.TimeOfAuthorization != nil {
		result["time_of_authorization"] = obj.TimeOfAuthorization.String()
	}

	return result
}

func (s *ApiaccesscontrolPrivilegedApiRequestResourceCrud) mapToPrivilegedApiRequestOperationDetails(fieldKeyFormat string) (oci_apiaccesscontrol.PrivilegedApiRequestOperationDetails, error) {
	result := oci_apiaccesscontrol.PrivilegedApiRequestOperationDetails{}

	if apiName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "api_name")); ok {
		tmp := apiName.(string)
		result.ApiName = &tmp
	}

	if attributeNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_names")); ok {
		interfaces := attributeNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "attribute_names")) {
			result.AttributeNames = tmp
		}
	}

	return result, nil
}

func PrivilegedApiRequestOperationDetailsToMap(obj oci_apiaccesscontrol.PrivilegedApiRequestOperationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApiName != nil {
		result["api_name"] = string(*obj.ApiName)
	}

	result["attribute_names"] = obj.AttributeNames

	return result
}

func PrivilegedApiRequestSummaryToMap(obj oci_apiaccesscontrol.PrivilegedApiRequestSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DurationInHrs != nil {
		result["duration_in_hrs"] = int(*obj.DurationInHrs)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	privilegedOperationList := []interface{}{}
	for _, item := range obj.PrivilegedOperationList {
		privilegedOperationList = append(privilegedOperationList, PrivilegedApiRequestOperationDetailsToMap(item))
	}
	result["privileged_operation_list"] = privilegedOperationList

	if obj.ReasonSummary != nil {
		result["reason_summary"] = string(*obj.ReasonSummary)
	}

	if obj.RequestId != nil {
		result["request_id"] = string(*obj.RequestId)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["severity"] = string(obj.Severity)

	result["state"] = string(obj.State)

	result["sub_resource_name_list"] = obj.SubResourceNameList

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeRequestedForFutureAccess != nil {
		result["time_requested_for_future_access"] = obj.TimeRequestedForFutureAccess.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
