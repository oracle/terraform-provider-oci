// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagedInstanceAttachSoftwareSourcesManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubManagedInstanceAttachSoftwareSourcesManagement,
		Read:     readOsManagementHubManagedInstanceAttachSoftwareSourcesManagement,
		Delete:   deleteOsManagementHubManagedInstanceAttachSoftwareSourcesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"software_sources": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
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

func createOsManagementHubManagedInstanceAttachSoftwareSourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceAttachSoftwareSourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubManagedInstanceAttachSoftwareSourcesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubManagedInstanceAttachSoftwareSourcesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubManagedInstanceAttachSoftwareSourcesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.ManagedInstanceClient
	Res                    *oci_os_management_hub.GetManagedInstanceResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_os_management_hub.WorkRequestClient
}

func (s *OsManagementHubManagedInstanceAttachSoftwareSourcesManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubManagedInstanceAttachSoftwareSourcesManagementResourceCrud) Get() error {
	request := oci_os_management_hub.GetManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubManagedInstanceAttachSoftwareSourcesManagementResourceCrud) Create() error {
	request := oci_os_management_hub.AttachSoftwareSourcesToManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if softwareSources, ok := s.D.GetOkExists("software_sources"); ok {
		interfaces := softwareSources.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("software_sources") {
			request.SoftwareSources = tmp
		}
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

	response, err := s.Client.AttachSoftwareSourcesToManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getManagedInstanceAttachSoftwareSourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub"), oci_os_management_hub.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OsManagementHubManagedInstanceAttachSoftwareSourcesManagementResourceCrud) getManagedInstanceAttachSoftwareSourcesManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_os_management_hub.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	managedInstanceAttachSoftwareSourcesManagementId, err := managedInstanceAttachSoftwareSourcesManagementWaitForWorkRequest(workId, "instance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*managedInstanceAttachSoftwareSourcesManagementId)

	return s.Get()
}

func managedInstanceAttachSoftwareSourcesManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func managedInstanceAttachSoftwareSourcesManagementWaitForWorkRequest(wId *string, entityType string, action oci_os_management_hub.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_os_management_hub.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "os_management_hub")
	retryPolicy.ShouldRetryOperation = managedInstanceAttachSoftwareSourcesManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_os_management_hub.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
		return nil, getErrorFromOsManagementHubManagedInstanceAttachSoftwareSourcesManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOsManagementHubManagedInstanceAttachSoftwareSourcesManagementWorkRequest(client *oci_os_management_hub.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_os_management_hub.ActionTypeEnum) error {
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

func (s *OsManagementHubManagedInstanceAttachSoftwareSourcesManagementResourceCrud) SetData() error {
	return nil
}

func (s *OsManagementHubManagedInstanceAttachSoftwareSourcesManagementResourceCrud) mapToWorkRequestDetails(fieldKeyFormat string) (oci_os_management_hub.WorkRequestDetails, error) {
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
