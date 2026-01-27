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

func OsManagementHubDynamicSetInstallPackagesManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubDynamicSetInstallPackagesManagement,
		Read:     readOsManagementHubDynamicSetInstallPackagesManagement,
		Delete:   deleteOsManagementHubDynamicSetInstallPackagesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"dynamic_set_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"package_names": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"managed_instances": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createOsManagementHubDynamicSetInstallPackagesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubDynamicSetInstallPackagesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DynamicSetClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubDynamicSetInstallPackagesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubDynamicSetInstallPackagesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubDynamicSetInstallPackagesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.DynamicSetClient
	Res                    *oci_os_management_hub.GetDynamicSetResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_os_management_hub.WorkRequestClient
}

func (s *OsManagementHubDynamicSetInstallPackagesManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubDynamicSetInstallPackagesManagementResourceCrud) Get() error {
	request := oci_os_management_hub.GetDynamicSetRequest{}

	if dynamicSetId, ok := s.D.GetOkExists("dynamic_set_id"); ok {
		tmp := dynamicSetId.(string)
		request.DynamicSetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetDynamicSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubDynamicSetInstallPackagesManagementResourceCrud) Create() error {
	request := oci_os_management_hub.InstallPackagesOnDynamicSetRequest{}

	if dynamicSetId, ok := s.D.GetOkExists("dynamic_set_id"); ok {
		tmp := dynamicSetId.(string)
		request.DynamicSetId = &tmp
	}

	if managedInstances, ok := s.D.GetOkExists("managed_instances"); ok {
		interfaces := managedInstances.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("managed_instances") {
			request.ManagedInstances = tmp
		}
	}

	if packageNames, ok := s.D.GetOkExists("package_names"); ok {
		interfaces := packageNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("package_names") {
			request.PackageNames = tmp
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

	response, err := s.Client.InstallPackagesOnDynamicSet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDynamicSetInstallPackagesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub"), oci_os_management_hub.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OsManagementHubDynamicSetInstallPackagesManagementResourceCrud) getDynamicSetInstallPackagesManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_os_management_hub.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	dynamicSetInstallPackagesManagementId, err := dynamicSetInstallPackagesManagementWaitForWorkRequest(workId, "dynamic_set",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*dynamicSetInstallPackagesManagementId)

	return s.Get()
}

func dynamicSetInstallPackagesManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func dynamicSetInstallPackagesManagementWaitForWorkRequest(wId *string, entityType string, action oci_os_management_hub.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_os_management_hub.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "os_management_hub")
	retryPolicy.ShouldRetryOperation = dynamicSetInstallPackagesManagementWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromOsManagementHubDynamicSetInstallPackagesManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOsManagementHubDynamicSetInstallPackagesManagementWorkRequest(client *oci_os_management_hub.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_os_management_hub.ActionTypeEnum) error {
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

func (s *OsManagementHubDynamicSetInstallPackagesManagementResourceCrud) SetData() error {
	return nil
}

func (s *OsManagementHubDynamicSetInstallPackagesManagementResourceCrud) mapToWorkRequestDetails(fieldKeyFormat string) (oci_os_management_hub.WorkRequestDetails, error) {
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
