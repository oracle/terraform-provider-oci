// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubLifecycleStagePromoteSoftwareSourceManagement,
		Read:     readOsManagementHubLifecycleStagePromoteSoftwareSourceManagement,
		Delete:   deleteOsManagementHubLifecycleStagePromoteSoftwareSourceManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"lifecycle_stage_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"software_source_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"work_request_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Computed
		},
	}
}

func createOsManagementHubLifecycleStagePromoteSoftwareSourceManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LifecycleEnvironmentClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubLifecycleStagePromoteSoftwareSourceManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubLifecycleStagePromoteSoftwareSourceManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.LifecycleEnvironmentClient
	Res                    *oci_os_management_hub.GetLifecycleStageResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_os_management_hub.WorkRequestClient
}

func (s *OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceCrud) Get() error {
	request := oci_os_management_hub.GetLifecycleStageRequest{}

	if lifecycleStageId, ok := s.D.GetOkExists("lifecycle_stage_id"); ok {
		tmp := lifecycleStageId.(string)
		request.LifecycleStageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetLifecycleStage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceCrud) Create() error {
	request := oci_os_management_hub.PromoteSoftwareSourceToLifecycleStageRequest{}

	if lifecycleStageId, ok := s.D.GetOkExists("lifecycle_stage_id"); ok {
		tmp := lifecycleStageId.(string)
		request.LifecycleStageId = &tmp
	}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	if workRequestDetails, ok := s.D.GetOkExists("work_request_details"); ok {
		if tmpList := workRequestDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "work_request_details", 0)
			tmp, err := s.mapToWorkRequestDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.WorkRequestDetails = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.PromoteSoftwareSourceToLifecycleStage(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getLifecycleStagePromoteSoftwareSourceManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub"), oci_os_management_hub.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceCrud) getLifecycleStagePromoteSoftwareSourceManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_os_management_hub.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	lifecycleStagePromoteSoftwareSourceManagementId, err := lifecycleStagePromoteSoftwareSourceManagementWaitForWorkRequest(workId, "lifecycle_environment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*lifecycleStagePromoteSoftwareSourceManagementId)

	return s.Get()
}

func lifecycleStagePromoteSoftwareSourceManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "os_management_hub", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_os_management_hub.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func lifecycleStagePromoteSoftwareSourceManagementWaitForWorkRequest(wId *string, entityType string, action oci_os_management_hub.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_os_management_hub.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "os_management_hub")
	retryPolicy.ShouldRetryOperation = lifecycleStagePromoteSoftwareSourceManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_os_management_hub.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_os_management_hub.OperationStatusInProgress),
			string(oci_os_management_hub.OperationStatusAccepted),
			string(oci_os_management_hub.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_os_management_hub.OperationStatusSucceeded),
			string(oci_os_management_hub.OperationStatusFailed),
			string(oci_os_management_hub.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_os_management_hub.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(string(res.EntityType)), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_os_management_hub.OperationStatusFailed || response.Status == oci_os_management_hub.OperationStatusCanceled {
		return nil, getErrorFromOsManagementHubLifecycleStagePromoteSoftwareSourceManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOsManagementHubLifecycleStagePromoteSoftwareSourceManagementWorkRequest(client *oci_os_management_hub.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_os_management_hub.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_os_management_hub.ListWorkRequestErrorsRequest{
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

func (s *OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceCrud) SetData() error {
	return nil
}

func (s *OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceCrud) mapToWorkRequestDetails(fieldKeyFormat string) (oci_os_management_hub.WorkRequestDetails, error) {
	result := oci_os_management_hub.WorkRequestDetails{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result, nil
}
