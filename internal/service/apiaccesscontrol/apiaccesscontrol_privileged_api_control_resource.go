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

func ApiaccesscontrolPrivilegedApiControlResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApiaccesscontrolPrivilegedApiControl,
		Read:     readApiaccesscontrolPrivilegedApiControl,
		Update:   updateApiaccesscontrolPrivilegedApiControl,
		Delete:   deleteApiaccesscontrolPrivilegedApiControl,
		Schema: map[string]*schema.Schema{
			// Required
			"approver_group_id_list": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notification_topic_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"privileged_operation_list": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"api_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"attribute_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"entity_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resources": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"number_of_approvers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_details": {
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
			"time_deleted": {
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

func createApiaccesscontrolPrivilegedApiControl(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolPrivilegedApiControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivilegedApiControlClient()
	sync.WorkRequestClient = m.(*client.OracleClients).ApiaccesscontrolPrivilegedApiWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readApiaccesscontrolPrivilegedApiControl(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolPrivilegedApiControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivilegedApiControlClient()

	return tfresource.ReadResource(sync)
}

func updateApiaccesscontrolPrivilegedApiControl(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolPrivilegedApiControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivilegedApiControlClient()
	sync.WorkRequestClient = m.(*client.OracleClients).ApiaccesscontrolPrivilegedApiWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApiaccesscontrolPrivilegedApiControl(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolPrivilegedApiControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivilegedApiControlClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).ApiaccesscontrolPrivilegedApiWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type ApiaccesscontrolPrivilegedApiControlResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apiaccesscontrol.PrivilegedApiControlClient
	Res                    *oci_apiaccesscontrol.PrivilegedApiControl
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_apiaccesscontrol.PrivilegedApiWorkRequestClient
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_apiaccesscontrol.PrivilegedApiControlLifecycleStateCreating),
	}
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_apiaccesscontrol.PrivilegedApiControlLifecycleStateActive),
		string(oci_apiaccesscontrol.PrivilegedApiControlLifecycleStateNeedsAttention),
	}
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_apiaccesscontrol.PrivilegedApiControlLifecycleStateDeleting),
	}
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_apiaccesscontrol.PrivilegedApiControlLifecycleStateDeleted),
	}
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) Create() error {
	request := oci_apiaccesscontrol.CreatePrivilegedApiControlRequest{}

	if approverGroupIdList, ok := s.D.GetOkExists("approver_group_id_list"); ok {
		interfaces := approverGroupIdList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("approver_group_id_list") {
			request.ApproverGroupIdList = tmp
		}
	}

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

	if notificationTopicId, ok := s.D.GetOkExists("notification_topic_id"); ok {
		tmp := notificationTopicId.(string)
		request.NotificationTopicId = &tmp
	}

	if numberOfApprovers, ok := s.D.GetOkExists("number_of_approvers"); ok {
		tmp := numberOfApprovers.(int)
		request.NumberOfApprovers = &tmp
	}

	if privilegedOperationList, ok := s.D.GetOkExists("privileged_operation_list"); ok {
		interfaces := privilegedOperationList.([]interface{})
		tmp := make([]oci_apiaccesscontrol.PrivilegedApiDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "privileged_operation_list", stateDataIndex)
			converted, err := s.mapToPrivilegedApiDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("privileged_operation_list") {
			request.PrivilegedOperationList = tmp
		}
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	if resources, ok := s.D.GetOkExists("resources"); ok {
		interfaces := resources.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("resources") {
			request.Resources = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apiaccesscontrol")

	response, err := s.Client.CreatePrivilegedApiControl(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getPrivilegedApiControlFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apiaccesscontrol"), oci_apiaccesscontrol.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) getPrivilegedApiControlFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_apiaccesscontrol.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	privilegedApiControlId, err := privilegedApiControlWaitForWorkRequest(workId, "privilegedapicontrol",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, privilegedApiControlId)
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
	s.D.SetId(*privilegedApiControlId)

	return s.Get()
}

func privilegedApiControlWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func privilegedApiControlWaitForWorkRequest(wId *string, entityType string, action oci_apiaccesscontrol.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_apiaccesscontrol.PrivilegedApiWorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "apiaccesscontrol")
	retryPolicy.ShouldRetryOperation = privilegedApiControlWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromApiaccesscontrolPrivilegedApiControlWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromApiaccesscontrolPrivilegedApiControlWorkRequest(client *oci_apiaccesscontrol.PrivilegedApiWorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_apiaccesscontrol.ActionTypeEnum) error {
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

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) Get() error {
	request := oci_apiaccesscontrol.GetPrivilegedApiControlRequest{}

	tmp := s.D.Id()
	request.PrivilegedApiControlId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apiaccesscontrol")

	response, err := s.Client.GetPrivilegedApiControl(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivilegedApiControl
	return nil
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_apiaccesscontrol.UpdatePrivilegedApiControlRequest{}

	if approverGroupIdList, ok := s.D.GetOkExists("approver_group_id_list"); ok {
		interfaces := approverGroupIdList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("approver_group_id_list") {
			request.ApproverGroupIdList = tmp
		}
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

	if notificationTopicId, ok := s.D.GetOkExists("notification_topic_id"); ok {
		tmp := notificationTopicId.(string)
		request.NotificationTopicId = &tmp
	}

	if numberOfApprovers, ok := s.D.GetOkExists("number_of_approvers"); ok {
		tmp := numberOfApprovers.(int)
		request.NumberOfApprovers = &tmp
	}

	tmp := s.D.Id()
	request.PrivilegedApiControlId = &tmp

	if privilegedOperationList, ok := s.D.GetOkExists("privileged_operation_list"); ok {
		interfaces := privilegedOperationList.([]interface{})
		tmp := make([]oci_apiaccesscontrol.PrivilegedApiDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "privileged_operation_list", stateDataIndex)
			converted, err := s.mapToPrivilegedApiDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("privileged_operation_list") {
			request.PrivilegedOperationList = tmp
		}
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	if resources, ok := s.D.GetOkExists("resources"); ok {
		interfaces := resources.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("resources") {
			request.Resources = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apiaccesscontrol")

	response, err := s.Client.UpdatePrivilegedApiControl(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPrivilegedApiControlFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apiaccesscontrol"), oci_apiaccesscontrol.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) Delete() error {
	request := oci_apiaccesscontrol.DeletePrivilegedApiControlRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.PrivilegedApiControlId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apiaccesscontrol")

	response, err := s.Client.DeletePrivilegedApiControl(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := privilegedApiControlWaitForWorkRequest(workId, "privilegedapicontrol",
		oci_apiaccesscontrol.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) SetData() error {
	s.D.Set("approver_group_id_list", s.Res.ApproverGroupIdList)

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NotificationTopicId != nil {
		s.D.Set("notification_topic_id", *s.Res.NotificationTopicId)
	}

	if s.Res.NumberOfApprovers != nil {
		s.D.Set("number_of_approvers", *s.Res.NumberOfApprovers)
	}

	privilegedOperationList := []interface{}{}
	for _, item := range s.Res.PrivilegedOperationList {
		privilegedOperationList = append(privilegedOperationList, PrivilegedApiDetailsToMap(item))
	}
	s.D.Set("privileged_operation_list", privilegedOperationList)

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", *s.Res.ResourceType)
	}

	s.D.Set("resources", s.Res.Resources)

	if s.Res.State != nil {
		s.D.Set("state", *s.Res.State)
	}

	if s.Res.StateDetails != nil {
		s.D.Set("state_details", *s.Res.StateDetails)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDeleted != nil {
		s.D.Set("time_deleted", s.Res.TimeDeleted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func PrivilegedApiControlSummaryToMap(obj oci_apiaccesscontrol.PrivilegedApiControlSummary) map[string]interface{} {
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

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.NumberOfApprovers != nil {
		result["number_of_approvers"] = int(*obj.NumberOfApprovers)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeDeleted != nil {
		result["time_deleted"] = obj.TimeDeleted.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) mapToPrivilegedApiDetails(fieldKeyFormat string) (oci_apiaccesscontrol.PrivilegedApiDetails, error) {
	result := oci_apiaccesscontrol.PrivilegedApiDetails{}

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

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		tmp := entityType.(string)
		result.EntityType = &tmp
	}

	return result, nil
}

func PrivilegedApiDetailsToMap(obj oci_apiaccesscontrol.PrivilegedApiDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApiName != nil {
		result["api_name"] = string(*obj.ApiName)
	}

	result["attribute_names"] = obj.AttributeNames

	if obj.EntityType != nil {
		result["entity_type"] = string(*obj.EntityType)
	}

	return result
}

func (s *ApiaccesscontrolPrivilegedApiControlResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_apiaccesscontrol.ChangePrivilegedApiControlCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.PrivilegedApiControlId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apiaccesscontrol")

	response, err := s.Client.ChangePrivilegedApiControlCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPrivilegedApiControlFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apiaccesscontrol"), oci_apiaccesscontrol.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
