// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oda

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OdaOdaPrivateEndpointAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOdaOdaPrivateEndpointAttachment,
		Read:     readOdaOdaPrivateEndpointAttachment,
		Delete:   deleteOdaOdaPrivateEndpointAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"oda_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"oda_private_endpoint_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"compartment_id": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOdaOdaPrivateEndpointAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()
	sync.OdaClient = m.(*client.OracleClients).OdaClient()

	return tfresource.CreateResource(d, sync)
}

func readOdaOdaPrivateEndpointAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()
	sync.OdaClient = m.(*client.OracleClients).OdaClient()

	return tfresource.ReadResource(sync)
}

func deleteOdaOdaPrivateEndpointAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()
	sync.OdaClient = m.(*client.OracleClients).OdaClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OdaOdaPrivateEndpointAttachmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_oda.ManagementClient
	OdaClient              *oci_oda.OdaClient
	Res                    *oci_oda.OdaPrivateEndpointAttachment
	DisableNotFoundRetries bool
}

func (s *OdaOdaPrivateEndpointAttachmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OdaOdaPrivateEndpointAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_oda.OdaPrivateEndpointAttachmentLifecycleStateCreating),
	}
}

func (s *OdaOdaPrivateEndpointAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_oda.OdaPrivateEndpointAttachmentLifecycleStateActive),
	}
}

func (s *OdaOdaPrivateEndpointAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_oda.OdaPrivateEndpointAttachmentLifecycleStateDeleting),
	}
}

func (s *OdaOdaPrivateEndpointAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_oda.OdaPrivateEndpointAttachmentLifecycleStateDeleted),
	}
}

func (s *OdaOdaPrivateEndpointAttachmentResourceCrud) Create() error {
	request := oci_oda.CreateOdaPrivateEndpointAttachmentRequest{}

	if odaInstanceId, ok := s.D.GetOkExists("oda_instance_id"); ok {
		tmp := odaInstanceId.(string)
		request.OdaInstanceId = &tmp
	}

	if odaPrivateEndpointId, ok := s.D.GetOkExists("oda_private_endpoint_id"); ok {
		tmp := odaPrivateEndpointId.(string)
		request.OdaPrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.CreateOdaPrivateEndpointAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getOdaPrivateEndpointAttachmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda"), oci_oda.WorkRequestResourceResourceActionCreate, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OdaOdaPrivateEndpointAttachmentResourceCrud) getOdaPrivateEndpointAttachmentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_oda.WorkRequestResourceResourceActionEnum, timeout time.Duration) error {

	// Wait until it finishes
	odaPrivateEndpointAttachmentId, err := odaPrivateEndpointAttachmentWaitForWorkRequest(workId, "oda",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.OdaClient)

	if err != nil {
		return err
	}
	s.D.SetId(*odaPrivateEndpointAttachmentId)

	return s.Get()
}

func odaPrivateEndpointAttachmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "oda", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_oda.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func odaPrivateEndpointAttachmentWaitForWorkRequest(wId *string, entityType string, action oci_oda.WorkRequestResourceResourceActionEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_oda.OdaClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "oda")
	retryPolicy.ShouldRetryOperation = odaPrivateEndpointAttachmentWorkRequestShouldRetryFunc(timeout)

	response := oci_oda.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_oda.WorkRequestStatusInProgress),
			string(oci_oda.WorkRequestStatusAccepted),
			string(oci_oda.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_oda.WorkRequestStatusSucceeded),
			string(oci_oda.WorkRequestStatusFailed),
			string(oci_oda.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_oda.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.ResourceType), entityType) {
			if res.ResourceAction == action {
				identifier = res.ResourceId
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_oda.WorkRequestStatusFailed || response.Status == oci_oda.WorkRequestStatusCanceled {
		return nil, getErrorFromOdaOdaPrivateEndpointAttachmentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOdaOdaPrivateEndpointAttachmentWorkRequest(client *oci_oda.OdaClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_oda.WorkRequestResourceResourceActionEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_oda.ListWorkRequestErrorsRequest{
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

func (s *OdaOdaPrivateEndpointAttachmentResourceCrud) Get() error {
	request := oci_oda.GetOdaPrivateEndpointAttachmentRequest{}

	tmp := s.D.Id()
	request.OdaPrivateEndpointAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.GetOdaPrivateEndpointAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OdaPrivateEndpointAttachment
	return nil
}

func (s *OdaOdaPrivateEndpointAttachmentResourceCrud) Delete() error {
	request := oci_oda.DeleteOdaPrivateEndpointAttachmentRequest{}

	tmp := s.D.Id()
	request.OdaPrivateEndpointAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.DeleteOdaPrivateEndpointAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := odaPrivateEndpointAttachmentWaitForWorkRequest(workId, "oda",
		oci_oda.WorkRequestResourceResourceActionDelete, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.OdaClient)
	return delWorkRequestErr
}

func (s *OdaOdaPrivateEndpointAttachmentResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.OdaInstanceId != nil {
		s.D.Set("oda_instance_id", *s.Res.OdaInstanceId)
	}

	if s.Res.OdaPrivateEndpointId != nil {
		s.D.Set("oda_private_endpoint_id", *s.Res.OdaPrivateEndpointId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func OdaPrivateEndpointAttachmentSummaryToMap(obj oci_oda.OdaPrivateEndpointAttachmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.OdaInstanceId != nil {
		result["oda_instance_id"] = string(*obj.OdaInstanceId)
	}

	if obj.OdaPrivateEndpointId != nil {
		result["oda_private_endpoint_id"] = string(*obj.OdaPrivateEndpointId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
