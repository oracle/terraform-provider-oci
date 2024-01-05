// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"fmt"
	"strings"
	"time"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"
)

func IdentityImportStandardTagsManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityImportStandardTagsManagement,
		Read:     readIdentityImportStandardTagsManagement,
		Delete:   deleteIdentityImportStandardTagsManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"standard_tag_namespace_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityImportStandardTagsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityImportStandardTagsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*tf_client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityImportStandardTagsManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteIdentityImportStandardTagsManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityImportStandardTagsManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.ImportStandardTagsResponse
	DisableNotFoundRetries bool
}

func (s *IdentityImportStandardTagsManagementResourceCrud) ID() string {
	return *s.Res.OpcWorkRequestId
}

func (s *IdentityImportStandardTagsManagementResourceCrud) Create() error {
	request := oci_identity.ImportStandardTagsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if standardTagNamespaceName, ok := s.D.GetOkExists("standard_tag_namespace_name"); ok {
		tmp := standardTagNamespaceName.(string)
		request.StandardTagNamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ImportStandardTags(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.D.SetId(*workId)

	s.Res = &response

	return s.getImportStandardTagsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity"), oci_identity.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IdentityImportStandardTagsManagementResourceCrud) getImportStandardTagsManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_identity.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := importStandardTagsManagementWaitForWorkRequest(workId, "identity",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	s.D.SetId(*workId)

	return nil
}

func importStandardTagsManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "identity", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_identity.GetTaggingWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func importStandardTagsManagementWaitForWorkRequest(wId *string, entityType string, action oci_identity.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_identity.IdentityClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "identity")
	retryPolicy.ShouldRetryOperation = importStandardTagsManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_identity.GetTaggingWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_identity.TaggingWorkRequestStatusInProgress),
			string(oci_identity.TaggingWorkRequestStatusAccepted),
			string(oci_identity.TaggingWorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_identity.TaggingWorkRequestStatusSucceeded),
			string(oci_identity.TaggingWorkRequestStatusFailed),
			string(oci_identity.TaggingWorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetTaggingWorkRequest(context.Background(),
				oci_identity.GetTaggingWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.TaggingWorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	// The work request response contains an array of resources, but that array will always be empty
	// Hence, no handling of this array is necessary (unlike in other Terraform managed resources)

	// The workrequest may have failed - so check for failed or canceled work requests
	if response.Status == oci_identity.TaggingWorkRequestStatusFailed || response.Status == oci_identity.TaggingWorkRequestStatusCanceled {
		return nil, getErrorFromIdentityImportStandardTagsManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return wId, nil

}

func getErrorFromIdentityImportStandardTagsManagementWorkRequest(client *oci_identity.IdentityClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_identity.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListTaggingWorkRequestErrors(context.Background(),
		oci_identity.ListTaggingWorkRequestErrorsRequest{
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

func (s *IdentityImportStandardTagsManagementResourceCrud) SetData() error {
	if s.Res.OpcWorkRequestId != nil {
		s.D.Set("work_request_id", *s.Res.OpcWorkRequestId)
	}

	return nil
}
