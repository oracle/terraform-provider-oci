// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v58/opsi"
)

func OpsiOperationsInsightsWarehouseRotateWarehouseWalletResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiOperationsInsightsWarehouseRotateWarehouseWallet,
		Read:     readOpsiOperationsInsightsWarehouseRotateWarehouseWallet,
		Delete:   deleteOpsiOperationsInsightsWarehouseRotateWarehouseWallet,
		Schema: map[string]*schema.Schema{
			// Required
			"operations_insights_warehouse_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createOpsiOperationsInsightsWarehouseRotateWarehouseWallet(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseRotateWarehouseWalletResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.CreateResource(d, sync)
}

func readOpsiOperationsInsightsWarehouseRotateWarehouseWallet(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOpsiOperationsInsightsWarehouseRotateWarehouseWallet(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OpsiOperationsInsightsWarehouseRotateWarehouseWalletResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *string
	DisableNotFoundRetries bool
}

func (s *OpsiOperationsInsightsWarehouseRotateWarehouseWalletResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("OpsiOperationsInsightsWarehouseRotateWarehouseWalletResource-", OpsiOperationsInsightsWarehouseRotateWarehouseWalletResource(), s.D)
}

func (s *OpsiOperationsInsightsWarehouseRotateWarehouseWalletResourceCrud) Create() error {
	request := oci_opsi.RotateOperationsInsightsWarehouseWalletRequest{}

	if operationsInsightsWarehouseId, ok := s.D.GetOkExists("operations_insights_warehouse_id"); ok {
		tmp := operationsInsightsWarehouseId.(string)
		request.OperationsInsightsWarehouseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.RotateOperationsInsightsWarehouseWallet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getOperationsInsightsWarehouseRotateWarehouseWalletFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeRelated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	if operationsInsightsWarehouseId, ok := s.D.GetOkExists("operations_insights_warehouse_id"); ok {
		tmp := operationsInsightsWarehouseId.(string)
		s.Res = &tmp
	}
	return nil
}

func (s *OpsiOperationsInsightsWarehouseRotateWarehouseWalletResourceCrud) getOperationsInsightsWarehouseRotateWarehouseWalletFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	operationsInsightsWarehouseRotateWarehouseWalletId, err := operationsInsightsWarehouseRotateWarehouseWalletWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*operationsInsightsWarehouseRotateWarehouseWalletId)
	return nil
	//	return tfresource.GenerateDataSourceHashID("OpsiOperationsInsightsWarehouseRotateWarehouseWalletResource", OpsiOperationsInsightsWarehouseRotateWarehouseWalletResource(), s.D)
}

func operationsInsightsWarehouseRotateWarehouseWalletWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func operationsInsightsWarehouseRotateWarehouseWalletWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = operationsInsightsWarehouseRotateWarehouseWalletWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromOpsiOperationsInsightsWarehouseRotateWarehouseWalletWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiOperationsInsightsWarehouseRotateWarehouseWalletWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
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

func (s *OpsiOperationsInsightsWarehouseRotateWarehouseWalletResourceCrud) SetData() error {
	return nil
}
