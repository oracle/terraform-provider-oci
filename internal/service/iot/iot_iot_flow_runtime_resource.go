// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotIotFlowRuntimeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotIotFlowRuntimeWithContext,
		ReadContext:   readIotIotFlowRuntimeWithContext,
		UpdateContext: updateIotIotFlowRuntimeWithContext,
		DeleteContext: deleteIotIotFlowRuntimeWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"iot_domain_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"display_name": {
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
			"log_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"log_group_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"log_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"network_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"file_storage_mounts": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"export_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"mount_path": {
										Type:     schema.TypeString,
										Required: true,
									},
									"mount_target_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"network_security_group_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"scale": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"flow_runtime_host": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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

func createIotIotFlowRuntimeWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotFlowRuntimeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotIotFlowRuntimeWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotFlowRuntimeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateIotIotFlowRuntimeWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotFlowRuntimeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteIotIotFlowRuntimeWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotFlowRuntimeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type IotIotFlowRuntimeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    *oci_iot.IotFlowRuntime
	DisableNotFoundRetries bool
}

func (s *IotIotFlowRuntimeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IotIotFlowRuntimeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_iot.IotFlowRuntimeLifecycleStateCreating),
	}
}

func (s *IotIotFlowRuntimeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_iot.IotFlowRuntimeLifecycleStateActive),
	}
}

func (s *IotIotFlowRuntimeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_iot.IotFlowRuntimeLifecycleStateDeleting),
	}
}

func (s *IotIotFlowRuntimeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_iot.IotFlowRuntimeLifecycleStateDeleted),
	}
}

func (s *IotIotFlowRuntimeResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.CreateIotFlowRuntimeRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if iotDomainId, ok := s.D.GetOkExists("iot_domain_id"); ok {
		tmp := iotDomainId.(string)
		request.IotDomainId = &tmp
	}

	if logConfig, ok := s.D.GetOkExists("log_config"); ok {
		if tmpList := logConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "log_config", 0)
			tmp, err := s.mapToLogConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LogConfig = &tmp
		}
	}

	if networkConfig, ok := s.D.GetOkExists("network_config"); ok {
		if tmpList := networkConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_config", 0)
			tmp, err := s.mapToNetworkConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkConfig = &tmp
		}
	}

	if scale, ok := s.D.GetOkExists("scale"); ok {
		request.Scale = oci_iot.CreateIotFlowRuntimeDetailsScaleEnum(scale.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.CreateIotFlowRuntime(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getIotFlowRuntimeFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot"), oci_iot.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IotIotFlowRuntimeResourceCrud) getIotFlowRuntimeFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_iot.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	iotFlowRuntimeId, err := iotFlowRuntimeWaitForWorkRequest(ctx, workId, "iotflowruntime",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*iotFlowRuntimeId)

	return s.GetWithContext(ctx)
}

func iotFlowRuntimeWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "iot", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_iot.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func iotFlowRuntimeWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_iot.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_iot.IotClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "iot")
	retryPolicy.ShouldRetryOperation = iotFlowRuntimeWorkRequestShouldRetryFunc(timeout)

	response := oci_iot.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_iot.OperationStatusAccepted),
			string(oci_iot.OperationStatusInProgress),
			string(oci_iot.OperationStatusWaiting),
			string(oci_iot.OperationStatusNeedsAttention),
		},
		Target: []string{
			string(oci_iot.OperationStatusFailed),
			string(oci_iot.OperationStatusSucceeded),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_iot.GetWorkRequestRequest{
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
	if _, e := stateConf.WaitForStateContext(ctx); e != nil {
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
	if identifier == nil || response.Status == oci_iot.OperationStatusFailed {
		return nil, getErrorFromIotIotFlowRuntimeWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromIotIotFlowRuntimeWorkRequest(ctx context.Context, client *oci_iot.IotClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_iot.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_iot.ListWorkRequestErrorsRequest{
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

func (s *IotIotFlowRuntimeResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetIotFlowRuntimeRequest{}

	tmp := s.D.Id()
	request.IotFlowRuntimeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.GetIotFlowRuntime(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.IotFlowRuntime
	return nil
}

func (s *IotIotFlowRuntimeResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_iot.UpdateIotFlowRuntimeRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.IotFlowRuntimeId = &tmp

	if logConfig, ok := s.D.GetOkExists("log_config"); ok {
		if tmpList := logConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "log_config", 0)
			tmp, err := s.mapToLogConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LogConfig = &tmp
		}
	}

	if networkConfig, ok := s.D.GetOkExists("network_config"); ok {
		if tmpList := networkConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_config", 0)
			tmp, err := s.mapToNetworkConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkConfig = &tmp
		}
	}

	if scale, ok := s.D.GetOkExists("scale"); ok {
		request.Scale = oci_iot.UpdateIotFlowRuntimeDetailsScaleEnum(scale.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.UpdateIotFlowRuntime(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIotFlowRuntimeFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot"), oci_iot.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *IotIotFlowRuntimeResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_iot.DeleteIotFlowRuntimeRequest{}

	tmp := s.D.Id()
	request.IotFlowRuntimeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.DeleteIotFlowRuntime(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := iotFlowRuntimeWaitForWorkRequest(ctx, workId, "iotflowruntime",
		oci_iot.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *IotIotFlowRuntimeResourceCrud) SetData() error {
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

	if s.Res.FlowRuntimeHost != nil {
		s.D.Set("flow_runtime_host", *s.Res.FlowRuntimeHost)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IotDomainId != nil {
		s.D.Set("iot_domain_id", *s.Res.IotDomainId)
	}

	if s.Res.LogConfig != nil {
		s.D.Set("log_config", []interface{}{LogConfigDetailsToMap(s.Res.LogConfig)})
	} else {
		s.D.Set("log_config", nil)
	}

	if s.Res.NetworkConfig != nil {
		s.D.Set("network_config", []interface{}{NetworkConfigDetailsToMap(s.Res.NetworkConfig, false)})
	} else {
		s.D.Set("network_config", nil)
	}

	s.D.Set("scale", s.Res.Scale)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *IotIotFlowRuntimeResourceCrud) mapToFileStorageMountDetails(fieldKeyFormat string) (oci_iot.FileStorageMountDetails, error) {
	result := oci_iot.FileStorageMountDetails{}

	if exportId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_id")); ok {
		tmp := exportId.(string)
		result.ExportId = &tmp
	}

	if mountPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_path")); ok {
		tmp := mountPath.(string)
		result.MountPath = &tmp
	}

	if mountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_target_id")); ok {
		tmp := mountTargetId.(string)
		result.MountTargetId = &tmp
	}

	return result, nil
}

func FileStorageMountDetailsToMap(obj oci_iot.FileStorageMountDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExportId != nil {
		result["export_id"] = string(*obj.ExportId)
	}

	if obj.MountPath != nil {
		result["mount_path"] = string(*obj.MountPath)
	}

	if obj.MountTargetId != nil {
		result["mount_target_id"] = string(*obj.MountTargetId)
	}

	return result
}

func IotFlowRuntimeSummaryToMap(obj oci_iot.IotFlowRuntimeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IotDomainId != nil {
		result["iot_domain_id"] = string(*obj.IotDomainId)
	}

	result["scale"] = string(obj.Scale)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *IotIotFlowRuntimeResourceCrud) mapToLogConfigDetails(fieldKeyFormat string) (oci_iot.LogConfigDetails, error) {
	result := oci_iot.LogConfigDetails{}

	if logGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_group_id")); ok {
		tmp := logGroupId.(string)
		result.LogGroupId = &tmp
	}

	if logId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_id")); ok {
		tmp := logId.(string)
		result.LogId = &tmp
	}

	return result, nil
}

func LogConfigDetailsToMap(obj *oci_iot.LogConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}

func (s *IotIotFlowRuntimeResourceCrud) mapToNetworkConfigDetails(fieldKeyFormat string) (oci_iot.NetworkConfigDetails, error) {
	result := oci_iot.NetworkConfigDetails{}

	if fileStorageMounts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_storage_mounts")); ok {
		interfaces := fileStorageMounts.([]interface{})
		tmp := make([]oci_iot.FileStorageMountDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_storage_mounts"), stateDataIndex)
			converted, err := s.mapToFileStorageMountDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "file_storage_mounts")) {
			result.FileStorageMounts = tmp
		}
	}

	if networkSecurityGroupIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_security_group_ids")); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "network_security_group_ids")) {
			result.NetworkSecurityGroupIds = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func NetworkConfigDetailsToMap(obj *oci_iot.NetworkConfigDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	fileStorageMounts := []interface{}{}
	for _, item := range obj.FileStorageMounts {
		fileStorageMounts = append(fileStorageMounts, FileStorageMountDetailsToMap(item))
	}
	result["file_storage_mounts"] = fileStorageMounts

	networkSecurityGroupIds := []interface{}{}
	for _, item := range obj.NetworkSecurityGroupIds {
		networkSecurityGroupIds = append(networkSecurityGroupIds, item)
	}
	if datasource {
		result["network_security_group_ids"] = networkSecurityGroupIds
	} else {
		result["network_security_group_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, networkSecurityGroupIds)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *IotIotFlowRuntimeResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_iot.ChangeIotFlowRuntimeCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.IotFlowRuntimeId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.ChangeIotFlowRuntimeCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIotFlowRuntimeFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot"), oci_iot.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
