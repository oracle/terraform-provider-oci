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

func OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubManagedInstanceGroupManageModuleStreamsManagement,
		Read:     readOsManagementHubManagedInstanceGroupManageModuleStreamsManagement,
		Delete:   deleteOsManagementHubManagedInstanceGroupManageModuleStreamsManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_instance_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"disable": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"module_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"stream_name": {
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

						// Computed
					},
				},
			},
			"enable": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"module_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"stream_name": {
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

						// Computed
					},
				},
			},
			"install": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"module_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"profile_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"stream_name": {
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

						// Computed
					},
				},
			},
			"is_dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"remove": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"module_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"profile_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"stream_name": {
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

						// Computed
					},
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

func createOsManagementHubManagedInstanceGroupManageModuleStreamsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceGroupClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubManagedInstanceGroupManageModuleStreamsManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubManagedInstanceGroupManageModuleStreamsManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.ManagedInstanceGroupClient
	Res                    *oci_os_management_hub.GetManagedInstanceGroupResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_os_management_hub.WorkRequestClient
}

func (s *OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceCrud) Get() error {
	request := oci_os_management_hub.GetManagedInstanceGroupRequest{}

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceCrud) Create() error {
	request := oci_os_management_hub.ManageModuleStreamsOnManagedInstanceGroupRequest{}

	if disable, ok := s.D.GetOkExists("disable"); ok {
		interfaces := disable.([]interface{})
		tmp := make([]oci_os_management_hub.ModuleStreamDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "disable", stateDataIndex)
			converted, err := s.mapToModuleStreamDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("disable") {
			request.Disable = tmp
		}
	}

	if enable, ok := s.D.GetOkExists("enable"); ok {
		interfaces := enable.([]interface{})
		tmp := make([]oci_os_management_hub.ModuleStreamDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "enable", stateDataIndex)
			converted, err := s.mapToModuleStreamDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("enable") {
			request.Enable = tmp
		}
	}

	if install, ok := s.D.GetOkExists("install"); ok {
		interfaces := install.([]interface{})
		tmp := make([]oci_os_management_hub.ModuleStreamProfileDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "install", stateDataIndex)
			converted, err := s.mapToModuleStreamProfileDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("install") {
			request.Install = tmp
		}
	}

	if isDryRun, ok := s.D.GetOkExists("is_dry_run"); ok {
		tmp := isDryRun.(bool)
		request.IsDryRun = &tmp
	}

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	if remove, ok := s.D.GetOkExists("remove"); ok {
		interfaces := remove.([]interface{})
		tmp := make([]oci_os_management_hub.ModuleStreamProfileDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "remove", stateDataIndex)
			converted, err := s.mapToModuleStreamProfileDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("remove") {
			request.Remove = tmp
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

	response, err := s.Client.ManageModuleStreamsOnManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getManagedInstanceGroupManageModuleStreamsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub"), oci_os_management_hub.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceCrud) getManagedInstanceGroupManageModuleStreamsManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_os_management_hub.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	managedInstanceGroupManageModuleStreamsManagementId, err := managedInstanceGroupManageModuleStreamsManagementWaitForWorkRequest(workId, "group",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*managedInstanceGroupManageModuleStreamsManagementId)

	return s.Get()
}

func managedInstanceGroupManageModuleStreamsManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func managedInstanceGroupManageModuleStreamsManagementWaitForWorkRequest(wId *string, entityType string, action oci_os_management_hub.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_os_management_hub.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "os_management_hub")
	retryPolicy.ShouldRetryOperation = managedInstanceGroupManageModuleStreamsManagementWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromOsManagementHubManagedInstanceGroupManageModuleStreamsManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOsManagementHubManagedInstanceGroupManageModuleStreamsManagementWorkRequest(client *oci_os_management_hub.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_os_management_hub.ActionTypeEnum) error {
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

func (s *OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceCrud) SetData() error {
	return nil
}

func (s *OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceCrud) mapToModuleStreamDetails(fieldKeyFormat string) (oci_os_management_hub.ModuleStreamDetails, error) {
	result := oci_os_management_hub.ModuleStreamDetails{}

	if moduleName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "module_name")); ok {
		tmp := moduleName.(string)
		result.ModuleName = &tmp
	}

	if softwareSourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "software_source_id")); ok {
		tmp := softwareSourceId.(string)
		result.SoftwareSourceId = &tmp
	}

	if streamName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_name")); ok {
		tmp := streamName.(string)
		result.StreamName = &tmp
	}

	return result, nil
}

func ModuleStreamDetailsToMap(obj oci_os_management_hub.ModuleStreamDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ModuleName != nil {
		result["module_name"] = string(*obj.ModuleName)
	}

	if obj.SoftwareSourceId != nil {
		result["software_source_id"] = string(*obj.SoftwareSourceId)
	}

	if obj.StreamName != nil {
		result["stream_name"] = string(*obj.StreamName)
	}

	return result
}

func (s *OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceCrud) mapToModuleStreamProfileDetails(fieldKeyFormat string) (oci_os_management_hub.ModuleStreamProfileDetails, error) {
	result := oci_os_management_hub.ModuleStreamProfileDetails{}

	if moduleName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "module_name")); ok {
		tmp := moduleName.(string)
		result.ModuleName = &tmp
	}

	if profileName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "profile_name")); ok {
		tmp := profileName.(string)
		result.ProfileName = &tmp
	}

	if softwareSourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "software_source_id")); ok {
		tmp := softwareSourceId.(string)
		result.SoftwareSourceId = &tmp
	}

	if streamName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_name")); ok {
		tmp := streamName.(string)
		result.StreamName = &tmp
	}

	return result, nil
}

func ModuleStreamProfileDetailsToMap(obj oci_os_management_hub.ModuleStreamProfileDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ModuleName != nil {
		result["module_name"] = string(*obj.ModuleName)
	}

	if obj.ProfileName != nil {
		result["profile_name"] = string(*obj.ProfileName)
	}

	if obj.SoftwareSourceId != nil {
		result["software_source_id"] = string(*obj.SoftwareSourceId)
	}

	if obj.StreamName != nil {
		result["stream_name"] = string(*obj.StreamName)
	}

	return result
}

func (s *OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceCrud) mapToWorkRequestDetails(fieldKeyFormat string) (oci_os_management_hub.WorkRequestDetails, error) {
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
