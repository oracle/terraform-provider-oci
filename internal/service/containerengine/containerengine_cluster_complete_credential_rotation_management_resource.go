// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineClusterCompleteCredentialRotationManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createContainerengineClusterCompleteCredentialRotationManagement,
		Read:   readContainerengineClusterCompleteCredentialRotationManagement,
		Delete: deleteContainerengineClusterCompleteCredentialRotationManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createContainerengineClusterCompleteCredentialRotationManagement(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterCompleteCredentialRotationManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.CreateResource(d, sync)
}

func readContainerengineClusterCompleteCredentialRotationManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteContainerengineClusterCompleteCredentialRotationManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type ContainerengineClusterCompleteCredentialRotationManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_containerengine.ContainerEngineClient
	Res                    *oci_containerengine.Cluster
	DisableNotFoundRetries bool
}

func (s *ContainerengineClusterCompleteCredentialRotationManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ContainerengineClusterCompleteCredentialRotationManagementResourceCrud) Create() error {
	request := oci_containerengine.CompleteCredentialRotationRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.CompleteCredentialRotation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	// Wait until it finishes
	clusterId, waitErr := clusterCompleteCredentialRotationManagementWaitForWorkRequest(workId, "cluster",
		oci_containerengine.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
	if waitErr != nil {
		return waitErr
	}

	requestGet := oci_containerengine.GetClusterRequest{}
	requestGet.ClusterId = clusterId
	requestGet.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")
	responseGet, getClusterErr := s.Client.GetCluster(context.Background(), requestGet)
	if getClusterErr != nil {
		return getClusterErr
	}
	s.Res = &responseGet.Cluster

	return nil
}

func clusterCompleteCredentialRotationManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "containerengine", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_containerengine.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func clusterCompleteCredentialRotationManagementWaitForWorkRequest(wId *string, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_containerengine.ContainerEngineClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "containerengine")
	retryPolicy.ShouldRetryOperation = clusterCompleteCredentialRotationManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_containerengine.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_containerengine.WorkRequestStatusInProgress),
			string(oci_containerengine.WorkRequestStatusAccepted),
			string(oci_containerengine.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_containerengine.WorkRequestStatusSucceeded),
			string(oci_containerengine.WorkRequestStatusFailed),
			string(oci_containerengine.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_containerengine.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_containerengine.WorkRequestStatusFailed || response.Status == oci_containerengine.WorkRequestStatusCanceled {
		return nil, getErrorFromContainerengineClusterCompleteCredentialRotationManagementWorkRequest(client, wId, response.CompartmentId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromContainerengineClusterCompleteCredentialRotationManagementWorkRequest(client *oci_containerengine.ContainerEngineClient, workId *string, compartmentId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_containerengine.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			CompartmentId: compartmentId,
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

func (s *ContainerengineClusterCompleteCredentialRotationManagementResourceCrud) SetData() error {
	return nil
}
