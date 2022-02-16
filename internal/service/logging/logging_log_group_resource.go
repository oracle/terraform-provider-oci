// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package logging

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_logging "github.com/oracle/oci-go-sdk/v58/logging"
)

func LoggingLogGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLoggingLogGroup,
		Read:     readLoggingLogGroup,
		Update:   updateLoggingLogGroup,
		Delete:   deleteLoggingLogGroup,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoggingLogGroup(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readLoggingLogGroup(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.ReadResource(sync)
}

func updateLoggingLogGroup(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLoggingLogGroup(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LoggingLogGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_logging.LoggingManagementClient
	Res                    *oci_logging.LogGroup
	DisableNotFoundRetries bool
}

func (s *LoggingLogGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *LoggingLogGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_logging.LogGroupLifecycleStateCreating),
	}
}

func (s *LoggingLogGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_logging.LogGroupLifecycleStateActive),
	}
}

func (s *LoggingLogGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_logging.LogGroupLifecycleStateDeleting),
	}
}

func (s *LoggingLogGroupResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *LoggingLogGroupResourceCrud) Create() error {
	request := oci_logging.CreateLogGroupRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.CreateLogGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getLogGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesCreated, 5*time.Minute)

}

func (s *LoggingLogGroupResourceCrud) getLogGroupFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_logging.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	logGroupId, err := logGroupWaitForWorkRequest(workId, "loggroup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, logGroupId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_logging.DeleteWorkRequestRequest{
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
	s.D.SetId(*logGroupId)

	return s.Get()
}

func logGroupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "logging", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		/*if workRequestResponse, ok := response.Response.(oci_logging.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}*/
		return false
	}
}

func logGroupWaitForWorkRequest(wId *string, entityType string, action oci_logging.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_logging.LoggingManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "logging")
	retryPolicy.ShouldRetryOperation = logGroupWorkRequestShouldRetryFunc(timeout)

	response := oci_logging.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_logging.OperationStatusInProgress),
			string(oci_logging.OperationStatusAccepted),
			string(oci_logging.OperationStatusCancelling),
		},
		Target: []string{
			string(oci_logging.OperationStatusSucceeded),
			string(oci_logging.OperationStatusFailed),
			string(oci_logging.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_logging.GetWorkRequestRequest{
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
			if res.ActionType == action || res.ActionType == oci_logging.ActionTypesInProgress {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_logging.OperationStatusFailed || response.Status == oci_logging.OperationStatusCanceled {
		return nil, getErrorFromLoggingLogGroupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromLoggingLogGroupWorkRequest(client *oci_logging.LoggingManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_logging.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_logging.ListWorkRequestErrorsRequest{
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

func (s *LoggingLogGroupResourceCrud) Get() error {
	request := oci_logging.GetLogGroupRequest{}

	tmp := s.D.Id()
	request.LogGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.GetLogGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogGroup
	return nil
}

func (s *LoggingLogGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_logging.UpdateLogGroupRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.LogGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.UpdateLogGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getLogGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *LoggingLogGroupResourceCrud) Delete() error {
	request := oci_logging.DeleteLogGroupRequest{}

	tmp := s.D.Id()
	request.LogGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	_, err := s.Client.DeleteLogGroup(context.Background(), request)
	return err
}

func (s *LoggingLogGroupResourceCrud) SetData() error {
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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}

func (s *LoggingLogGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_logging.ChangeLogGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.LogGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.ChangeLogGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes

	return s.getLogGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesRelated, s.D.Timeout(schema.TimeoutUpdate))
}
