// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateConnectionAssignmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createGoldenGateConnectionAssignment,
		Read:     readGoldenGateConnectionAssignment,
		Delete:   deleteGoldenGateConnectionAssignment,
		Schema: map[string]*schema.Schema{
			// Required
			"connection_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"deployment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"alias_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
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

func createGoldenGateConnectionAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateConnectionAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.CreateResource(d, sync)
}

func readGoldenGateConnectionAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateConnectionAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

func deleteGoldenGateConnectionAssignment(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateConnectionAssignmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GoldenGateConnectionAssignmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_golden_gate.GoldenGateClient
	Res                    *oci_golden_gate.ConnectionAssignment
	DisableNotFoundRetries bool
}

func (s *GoldenGateConnectionAssignmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GoldenGateConnectionAssignmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_golden_gate.ConnectionAssignmentLifecycleStateDeleting),
	}
}

func (s *GoldenGateConnectionAssignmentResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *GoldenGateConnectionAssignmentResourceCrud) Create() error {
	request := oci_golden_gate.CreateConnectionAssignmentRequest{}

	if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
		tmp := connectionId.(string)
		request.ConnectionId = &tmp
	}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")
	response, err := s.Client.CreateConnectionAssignment(context.Background(), request)

	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	// Wait until it finishes

	_, err = connectionAssignmentWaitForWorkRequest(workId, "goldengateconnectionassignment",
		oci_golden_gate.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*response.ConnectionAssignment.Id)

	return s.Get()
}

func connectionAssignmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "golden_gate", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_golden_gate.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func connectionAssignmentWaitForWorkRequest(wId *string, entityType string, action oci_golden_gate.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_golden_gate.GoldenGateClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "golden_gate")
	retryPolicy.ShouldRetryOperation = connectionAssignmentWorkRequestShouldRetryFunc(timeout)

	response := oci_golden_gate.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_golden_gate.OperationStatusInProgress),
			string(oci_golden_gate.OperationStatusAccepted),
			string(oci_golden_gate.OperationStatusCanceled),
		},
		Target: []string{
			string(oci_golden_gate.OperationStatusSucceeded),
			string(oci_golden_gate.OperationStatusFailed),
			string(oci_golden_gate.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_golden_gate.GetWorkRequestRequest{
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

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if response.Status == oci_golden_gate.OperationStatusFailed || response.Status == oci_golden_gate.OperationStatusCanceled {
		return nil, getErrorFromGoldenGateConnectionAssignmentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return nil, nil
}

func getErrorFromGoldenGateConnectionAssignmentWorkRequest(client *oci_golden_gate.GoldenGateClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_golden_gate.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_golden_gate.ListWorkRequestErrorsRequest{
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

func (s *GoldenGateConnectionAssignmentResourceCrud) Get() error {
	request := oci_golden_gate.GetConnectionAssignmentRequest{}

	tmp := s.D.Id()
	request.ConnectionAssignmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.GetConnectionAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConnectionAssignment
	return nil
}

func (s *GoldenGateConnectionAssignmentResourceCrud) Delete() error {
	request := oci_golden_gate.DeleteConnectionAssignmentRequest{}

	tmp := s.D.Id()
	request.ConnectionAssignmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.DeleteConnectionAssignment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := connectionAssignmentWaitForWorkRequest(workId, "goldengateconnectionassignment",
		oci_golden_gate.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GoldenGateConnectionAssignmentResourceCrud) SetData() error {
	if s.Res.Id != nil {
		s.D.SetId(*s.Res.Id)
	}

	if s.Res.AliasName != nil {
		s.D.Set("alias_name", *s.Res.AliasName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionId != nil {
		s.D.Set("connection_id", *s.Res.ConnectionId)
	}

	if s.Res.DeploymentId != nil {
		s.D.Set("deployment_id", *s.Res.DeploymentId)
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

func ConnectionAssignmentSummaryToMap(obj oci_golden_gate.ConnectionAssignmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AliasName != nil {
		result["alias_name"] = string(*obj.AliasName)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConnectionId != nil {
		result["connection_id"] = string(*obj.ConnectionId)
	}

	if obj.DeploymentId != nil {
		result["deployment_id"] = string(*obj.DeploymentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
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
