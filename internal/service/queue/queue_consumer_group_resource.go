// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package queue

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_queue "github.com/oracle/oci-go-sdk/v65/queue"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func QueueConsumerGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createQueueConsumerGroup,
		Read:     readQueueConsumerGroup,
		Update:   updateQueueConsumerGroup,
		Delete:   deleteQueueConsumerGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"queue_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"consumer_group_filter": {
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
			"is_enabled": {
				Type:     schema.TypeBool,
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

func createQueueConsumerGroup(d *schema.ResourceData, m interface{}) error {
	sync := &QueueConsumerGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueueAdminClient()

	return tfresource.CreateResource(d, sync)
}

func readQueueConsumerGroup(d *schema.ResourceData, m interface{}) error {
	sync := &QueueConsumerGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueueAdminClient()

	return tfresource.ReadResource(sync)
}

func updateQueueConsumerGroup(d *schema.ResourceData, m interface{}) error {
	sync := &QueueConsumerGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueueAdminClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteQueueConsumerGroup(d *schema.ResourceData, m interface{}) error {
	sync := &QueueConsumerGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueueAdminClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type QueueConsumerGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_queue.QueueAdminClient
	Res                    *oci_queue.ConsumerGroup
	DisableNotFoundRetries bool
}

func (s *QueueConsumerGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *QueueConsumerGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_queue.ConsumerGroupLifecycleStateCreating),
	}
}

func (s *QueueConsumerGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_queue.ConsumerGroupLifecycleStateActive),
	}
}

func (s *QueueConsumerGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_queue.ConsumerGroupLifecycleStateDeleting),
	}
}

func (s *QueueConsumerGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_queue.ConsumerGroupLifecycleStateDeleted),
	}
}

func (s *QueueConsumerGroupResourceCrud) Create() error {
	request := oci_queue.CreateConsumerGroupRequest{}

	if consumerGroupFilter, ok := s.D.GetOkExists("consumer_group_filter"); ok {
		tmp := consumerGroupFilter.(string)
		request.Filter = &tmp
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

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if queueId, ok := s.D.GetOkExists("queue_id"); ok {
		tmp := queueId.(string)
		request.QueueId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue")

	response, err := s.Client.CreateConsumerGroup(context.Background(), request)
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "consumergroup") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getConsumerGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue"), oci_queue.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *QueueConsumerGroupResourceCrud) getConsumerGroupFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_queue.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	consumerGroupId, err := consumerGroupWaitForWorkRequest(workId, "consumergroup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*consumerGroupId)

	return s.Get()
}

func consumerGroupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func consumerGroupWaitForWorkRequest(wId *string, entityType string, action oci_queue.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_queue.QueueAdminClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "queue")
	retryPolicy.ShouldRetryOperation = consumerGroupWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromQueueConsumerGroupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromQueueConsumerGroupWorkRequest(client *oci_queue.QueueAdminClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_queue.ActionTypeEnum) error {
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

func (s *QueueConsumerGroupResourceCrud) Get() error {
	request := oci_queue.GetConsumerGroupRequest{}

	tmp := s.D.Id()
	request.ConsumerGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue")

	response, err := s.Client.GetConsumerGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsumerGroup
	return nil
}

func (s *QueueConsumerGroupResourceCrud) Update() error {
	request := oci_queue.UpdateConsumerGroupRequest{}

	tmp := s.D.Id()
	request.ConsumerGroupId = &tmp

	if consumerGroupFilter, ok := s.D.GetOkExists("consumer_group_filter"); ok {
		tmp := consumerGroupFilter.(string)
		request.Filter = &tmp
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

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue")

	response, err := s.Client.UpdateConsumerGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getConsumerGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue"), oci_queue.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *QueueConsumerGroupResourceCrud) Delete() error {
	request := oci_queue.DeleteConsumerGroupRequest{}

	tmp := s.D.Id()
	request.ConsumerGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "queue")

	response, err := s.Client.DeleteConsumerGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := consumerGroupWaitForWorkRequest(workId, "consumergroup",
		oci_queue.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *QueueConsumerGroupResourceCrud) SetData() error {
	if s.Res.Filter != nil {
		s.D.Set("consumer_group_filter", *s.Res.Filter)
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

	if s.Res.QueueId != nil {
		s.D.Set("queue_id", *s.Res.QueueId)
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

	return nil
}

func ConsumerGroupSummaryToMap(obj oci_queue.ConsumerGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Filter != nil {
		result["consumer_group_filter"] = string(*obj.Filter)
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

	if obj.QueueId != nil {
		result["queue_id"] = string(*obj.QueueId)
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
