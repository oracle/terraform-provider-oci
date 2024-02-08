// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiAwrHubSourceAwrhubsourcesManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiAwrHubSourceAwrhubsourcesManagement,
		Read:     readOpsiAwrHubSourceAwrhubsourcesManagement,
		Update:   updateOpsiAwrHubSourceAwrhubsourcesManagement,
		Delete:   deleteOpsiAwrHubSourceAwrhubsourcesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"awr_hub_source_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_awrhubsource": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional

			// Computed
		},
	}
}

func createOpsiAwrHubSourceAwrhubsourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubSourceAwrhubsourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.Res = &OpsiAwrHubSourceAwrhubsourcesManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readOpsiAwrHubSourceAwrhubsourcesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateOpsiAwrHubSourceAwrhubsourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubSourceAwrhubsourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.Res = &OpsiAwrHubSourceAwrhubsourcesManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteOpsiAwrHubSourceAwrhubsourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubSourceAwrhubsourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.Res = &OpsiAwrHubSourceAwrhubsourcesManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpsiAwrHubSourceAwrhubsourcesManagementResponse struct {
	enableResponse  *oci_opsi.EnableAwrHubSourceResponse
	disableResponse *oci_opsi.DisableAwrHubSourceResponse
}

type OpsiAwrHubSourceAwrhubsourcesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *OpsiAwrHubSourceAwrhubsourcesManagementResponse
	DisableNotFoundRetries bool
}

func (s *OpsiAwrHubSourceAwrhubsourcesManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("OpsiAwrHubSourceAwrhubsourcesManagementResource-", OpsiAwrHubSourceAwrhubsourcesManagementResource(), s.D)
}

func (s *OpsiAwrHubSourceAwrhubsourcesManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_awrhubsource"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_opsi.EnableAwrHubSourceRequest{}

		if awrHubSourceId, ok := s.D.GetOkExists("awr_hub_source_id"); ok {
			tmp := awrHubSourceId.(string)
			request.AwrHubSourceId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

		response, err := s.Client.EnableAwrHubSource(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getAwrHubSourceAwrhubsourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}

		s.Res.enableResponse = &response
		return nil
	}

	request := oci_opsi.DisableAwrHubSourceRequest{}

	if awrHubSourceId, ok := s.D.GetOkExists("awr_hub_source_id"); ok {
		tmp := awrHubSourceId.(string)
		request.AwrHubSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")
	response, err := s.Client.DisableAwrHubSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getAwrHubSourceAwrhubsourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *OpsiAwrHubSourceAwrhubsourcesManagementResourceCrud) getAwrHubSourceAwrhubsourcesManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := awrHubSourceAwrhubsourcesManagementWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func awrHubSourceAwrhubsourcesManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "opsi", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_opsi.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func awrHubSourceAwrhubsourcesManagementWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = awrHubSourceAwrhubsourcesManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_opsi.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_opsi.OperationStatusInProgress),
			string(oci_opsi.OperationStatusAccepted),
			string(oci_opsi.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_opsi.OperationStatusSucceeded),
			string(oci_opsi.OperationStatusFailed),
			string(oci_opsi.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_opsi.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_opsi.OperationStatusFailed || response.Status == oci_opsi.OperationStatusCanceled {
		return nil, getErrorFromOpsiAwrHubSourceAwrhubsourcesManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiAwrHubSourceAwrhubsourcesManagementWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_opsi.ListWorkRequestErrorsRequest{
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

func (s *OpsiAwrHubSourceAwrhubsourcesManagementResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_awrhubsource"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_opsi.EnableAwrHubSourceRequest{}

		if awrHubSourceId, ok := s.D.GetOkExists("awr_hub_source_id"); ok {
			tmp := awrHubSourceId.(string)
			request.AwrHubSourceId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

		response, err := s.Client.EnableAwrHubSource(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getAwrHubSourceAwrhubsourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}

		s.Res.enableResponse = &response
		return nil
	}

	request := oci_opsi.DisableAwrHubSourceRequest{}

	if awrHubSourceId, ok := s.D.GetOkExists("awr_hub_source_id"); ok {
		tmp := awrHubSourceId.(string)
		request.AwrHubSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DisableAwrHubSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getAwrHubSourceAwrhubsourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *OpsiAwrHubSourceAwrhubsourcesManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_awrhubsource"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_opsi.DisableAwrHubSourceRequest{}

	if awrHubSourceId, ok := s.D.GetOkExists("awr_hub_source_id"); ok {
		tmp := awrHubSourceId.(string)
		request.AwrHubSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DisableAwrHubSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getAwrHubSourceAwrhubsourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *OpsiAwrHubSourceAwrhubsourcesManagementResourceCrud) SetData() error {
	return nil
}
