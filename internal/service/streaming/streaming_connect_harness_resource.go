// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package streaming

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_streaming "github.com/oracle/oci-go-sdk/v65/streaming"
)

func StreamingConnectHarnessResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStreamingConnectHarness,
		Read:     readStreamingConnectHarness,
		Update:   updateStreamingConnectHarness,
		Delete:   deleteStreamingConnectHarness,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"lifecycle_state_details": {
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
		},
	}
}

func createStreamingConnectHarness(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingConnectHarnessResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.CreateResource(d, sync)
}

func readStreamingConnectHarness(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingConnectHarnessResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.ReadResource(sync)
}

func updateStreamingConnectHarness(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingConnectHarnessResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStreamingConnectHarness(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingConnectHarnessResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StreamingConnectHarnessResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_streaming.StreamAdminClient
	Res                    *oci_streaming.ConnectHarness
	DisableNotFoundRetries bool
}

func (s *StreamingConnectHarnessResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StreamingConnectHarnessResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateCreating),
	}
}

func (s *StreamingConnectHarnessResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateActive),
	}
}

func (s *StreamingConnectHarnessResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateDeleting),
	}
}

func (s *StreamingConnectHarnessResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateDeleted),
	}
}

func (s *StreamingConnectHarnessResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateUpdating),
	}
}

func (s *StreamingConnectHarnessResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateActive),
	}
}

func (s *StreamingConnectHarnessResourceCrud) Create() error {
	request := oci_streaming.CreateConnectHarnessRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.CreateConnectHarness(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getConnectHarnessFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming"), oci_streaming.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *StreamingConnectHarnessResourceCrud) getConnectHarnessFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_streaming.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	connectHarnessId, err := connectHarnessWaitForWorkRequest(workId, "connect_harness",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*connectHarnessId)

	return s.Get()
}

func connectHarnessWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "streaming", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_streaming.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func connectHarnessWaitForWorkRequest(wId *string, entityType string, action oci_streaming.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_streaming.StreamAdminClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "streaming")
	retryPolicy.ShouldRetryOperation = connectHarnessWorkRequestShouldRetryFunc(timeout)

	response := oci_streaming.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_streaming.OperationStatusInProgress),
			string(oci_streaming.OperationStatusAccepted),
			string(oci_streaming.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_streaming.OperationStatusSucceeded),
			string(oci_streaming.OperationStatusFailed),
			string(oci_streaming.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_streaming.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_streaming.OperationStatusFailed || response.Status == oci_streaming.OperationStatusCanceled {
		return nil, getErrorFromStreamingConnectHarnessWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromStreamingConnectHarnessWorkRequest(client *oci_streaming.StreamAdminClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_streaming.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_streaming.ListWorkRequestErrorsRequest{
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

func (s *StreamingConnectHarnessResourceCrud) Get() error {
	request := oci_streaming.GetConnectHarnessRequest{}

	tmp := s.D.Id()
	request.ConnectHarnessId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.GetConnectHarness(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConnectHarness
	return nil
}

func (s *StreamingConnectHarnessResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_streaming.UpdateConnectHarnessRequest{}

	tmp := s.D.Id()
	request.ConnectHarnessId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.UpdateConnectHarness(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getConnectHarnessFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming"), oci_streaming.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *StreamingConnectHarnessResourceCrud) Delete() error {
	request := oci_streaming.DeleteConnectHarnessRequest{}

	tmp := s.D.Id()
	request.ConnectHarnessId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.DeleteConnectHarness(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := connectHarnessWaitForWorkRequest(workId, "connect_harness",
		oci_streaming.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *StreamingConnectHarnessResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *StreamingConnectHarnessResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_streaming.ChangeConnectHarnessCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ConnectHarnessId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.ChangeConnectHarnessCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getConnectHarnessFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming"), oci_streaming.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
