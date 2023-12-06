// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package queue

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_queue "github.com/oracle/oci-go-sdk/v65/queue"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func QueueQueueResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createQueueQueue,
		Read:     readQueueQueue,
		Update:   updateQueueQueue,
		Delete:   deleteQueueQueue,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"channel_consumption_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"custom_encryption_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dead_letter_queue_delivery_count": {
				Type:     schema.TypeInt,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"retention_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"visibility_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"purge_queue": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"purge_type": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"messages_endpoint": {
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

func createQueueQueue(d *schema.ResourceData, m interface{}) error {
	sync := &QueueQueueResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueueAdminClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if value, ok := sync.D.GetOkExists("purge_queue"); ok && value.(bool) {
		err := sync.PurgeQueue()
		if err != nil {
			return err
		}
	}
	return nil

}

func readQueueQueue(d *schema.ResourceData, m interface{}) error {
	sync := &QueueQueueResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueueAdminClient()

	return tfresource.ReadResource(sync)
}

func updateQueueQueue(d *schema.ResourceData, m interface{}) error {
	sync := &QueueQueueResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueueAdminClient()

	// Checking only if the purge queue key exists in configuration and is set to true.
	if value, ok := sync.D.GetOkExists("purge_queue"); ok && value.(bool) {
		err := sync.PurgeQueue()
		if err != nil {
			return err
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteQueueQueue(d *schema.ResourceData, m interface{}) error {
	sync := &QueueQueueResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueueAdminClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type QueueQueueResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_queue.QueueAdminClient
	Res                    *oci_queue.Queue
	DisableNotFoundRetries bool
}

func (s *QueueQueueResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *QueueQueueResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_queue.QueueLifecycleStateCreating),
	}
}

func (s *QueueQueueResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_queue.QueueLifecycleStateActive),
	}
}

func (s *QueueQueueResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_queue.QueueLifecycleStateDeleting),
	}
}

func (s *QueueQueueResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_queue.QueueLifecycleStateDeleted),
	}
}

func (s *QueueQueueResourceCrud) Create() error {
	request := oci_queue.CreateQueueRequest{}

	if channelConsumptionLimit, ok := s.D.GetOkExists("channel_consumption_limit"); ok {
		tmp := channelConsumptionLimit.(int)
		request.ChannelConsumptionLimit = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if customEncryptionKeyId, ok := s.D.GetOkExists("custom_encryption_key_id"); ok {
		tmp := customEncryptionKeyId.(string)
		request.CustomEncryptionKeyId = &tmp
	}

	if deadLetterQueueDeliveryCount, ok := s.D.GetOkExists("dead_letter_queue_delivery_count"); ok {
		tmp := deadLetterQueueDeliveryCount.(int)
		request.DeadLetterQueueDeliveryCount = &tmp
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if retentionInSeconds, ok := s.D.GetOkExists("retention_in_seconds"); ok {
		tmp := retentionInSeconds.(int)
		request.RetentionInSeconds = &tmp
	}

	if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
		tmp := timeoutInSeconds.(int)
		request.TimeoutInSeconds = &tmp
	}

	if visibilityInSeconds, ok := s.D.GetOkExists("visibility_in_seconds"); ok {
		tmp := visibilityInSeconds.(int)
		request.VisibilityInSeconds = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue")

	response, err := s.Client.CreateQueue(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_queue.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_queue.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "queue") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getQueueFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue"), oci_queue.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *QueueQueueResourceCrud) getQueueFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_queue.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	queueId, err := queueWaitForWorkRequest(workId, "queue",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*queueId)

	return s.Get()
}

func queueWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "queue", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_queue.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func queueWaitForWorkRequest(wId *string, entityType string, action oci_queue.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_queue.QueueAdminClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "queue")
	retryPolicy.ShouldRetryOperation = queueWorkRequestShouldRetryFunc(timeout)

	response := oci_queue.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_queue.OperationStatusInProgress),
			string(oci_queue.OperationStatusAccepted),
			string(oci_queue.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_queue.OperationStatusSucceeded),
			string(oci_queue.OperationStatusFailed),
			string(oci_queue.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_queue.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_queue.OperationStatusFailed || response.Status == oci_queue.OperationStatusCanceled {
		return nil, getErrorFromQueueQueueWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromQueueQueueWorkRequest(client *oci_queue.QueueAdminClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_queue.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_queue.ListWorkRequestErrorsRequest{
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

func (s *QueueQueueResourceCrud) Get() error {
	request := oci_queue.GetQueueRequest{}

	tmp := s.D.Id()
	request.QueueId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue")

	response, err := s.Client.GetQueue(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Queue
	return nil
}

func (s *QueueQueueResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_queue.UpdateQueueRequest{}

	if channelConsumptionLimit, ok := s.D.GetOkExists("channel_consumption_limit"); ok {
		tmp := channelConsumptionLimit.(int)
		request.ChannelConsumptionLimit = &tmp
	}

	if customEncryptionKeyId, ok := s.D.GetOkExists("custom_encryption_key_id"); ok {
		tmp := customEncryptionKeyId.(string)
		request.CustomEncryptionKeyId = &tmp
	}

	if deadLetterQueueDeliveryCount, ok := s.D.GetOkExists("dead_letter_queue_delivery_count"); ok {
		tmp := deadLetterQueueDeliveryCount.(int)
		request.DeadLetterQueueDeliveryCount = &tmp
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.QueueId = &tmp

	if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
		tmp := timeoutInSeconds.(int)
		request.TimeoutInSeconds = &tmp
	}

	if visibilityInSeconds, ok := s.D.GetOkExists("visibility_in_seconds"); ok {
		tmp := visibilityInSeconds.(int)
		request.VisibilityInSeconds = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue")

	response, err := s.Client.UpdateQueue(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getQueueFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue"), oci_queue.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *QueueQueueResourceCrud) Delete() error {
	request := oci_queue.DeleteQueueRequest{}

	tmp := s.D.Id()
	request.QueueId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue")

	response, err := s.Client.DeleteQueue(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := queueWaitForWorkRequest(workId, "queue",
		oci_queue.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *QueueQueueResourceCrud) SetData() error {
	if s.Res.ChannelConsumptionLimit != nil {
		s.D.Set("channel_consumption_limit", *s.Res.ChannelConsumptionLimit)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CustomEncryptionKeyId != nil {
		s.D.Set("custom_encryption_key_id", *s.Res.CustomEncryptionKeyId)
	}

	if s.Res.DeadLetterQueueDeliveryCount != nil {
		s.D.Set("dead_letter_queue_delivery_count", *s.Res.DeadLetterQueueDeliveryCount)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MessagesEndpoint != nil {
		s.D.Set("messages_endpoint", *s.Res.MessagesEndpoint)
	}

	if s.Res.RetentionInSeconds != nil {
		s.D.Set("retention_in_seconds", *s.Res.RetentionInSeconds)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeoutInSeconds != nil {
		s.D.Set("timeout_in_seconds", *s.Res.TimeoutInSeconds)
	}

	if s.Res.VisibilityInSeconds != nil {
		s.D.Set("visibility_in_seconds", *s.Res.VisibilityInSeconds)
	}

	return nil
}

func (s *QueueQueueResourceCrud) PurgeQueue() error {
	request := oci_queue.PurgeQueueRequest{}

	//This is an auto generated code for channelIds, may not be used since we wont be passing channel IDs in tests
	if channelIds, ok := s.D.GetOkExists("channel_ids"); ok {
		interfaces := channelIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("channel_ids") {
			request.ChannelIds = tmp
		}
	}

	if purgeType, ok := s.D.GetOkExists("purge_type"); ok {
		request.PurgeType, _ = oci_queue.GetMappingPurgeQueueDetailsPurgeTypeEnum(purgeType.(string))
	}

	idTmp := s.D.Id()
	request.QueueId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue")

	response, err := s.Client.PurgeQueue(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("purge_queue")
	s.D.Set("purge_queue", val)

	workId := response.OpcWorkRequestId
	return s.getQueueFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue"), oci_queue.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func QueueSummaryToMap(obj oci_queue.QueueSummary) map[string]interface{} {
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

	if obj.MessagesEndpoint != nil {
		result["messages_endpoint"] = string(*obj.MessagesEndpoint)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *QueueQueueResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_queue.ChangeQueueCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.QueueId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue")

	response, err := s.Client.ChangeQueueCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getQueueFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue"), oci_queue.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
