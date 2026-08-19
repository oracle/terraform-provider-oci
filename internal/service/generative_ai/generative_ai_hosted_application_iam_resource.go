// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiHostedApplicationIamResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createGenerativeAiHostedApplicationIamWithContext,
		ReadContext:   readGenerativeAiHostedApplicationIamWithContext,
		UpdateContext: updateGenerativeAiHostedApplicationIamWithContext,
		DeleteContext: deleteGenerativeAiHostedApplicationIamWithContext,
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
			"environment_variables": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     validation.StringIsJSON,
							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
						},

						// Optional

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"networking_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"inbound_networking_config": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"endpoint_mode": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"private_endpoint_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"outbound_networking_config": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"network_mode": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"custom_subnet_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"nsg_ids": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Set:      tfresource.LiteralTypeHashCodeForSets,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"scaling_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"scaling_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"max_replica": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"min_replica": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"target_concurrency_threshold": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"target_cpu_threshold": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"target_memory_threshold": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"target_rps_threshold": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"storage_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"environment_variable_key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"storage_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"lifecycle_details": {
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

func createGenerativeAiHostedApplicationIamWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedApplicationIamResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readGenerativeAiHostedApplicationIamWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedApplicationIamResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateGenerativeAiHostedApplicationIamWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedApplicationIamResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteGenerativeAiHostedApplicationIamWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedApplicationIamResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type GenerativeAiHostedApplicationIamResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai.GenerativeAiClient
	Res                    *oci_generative_ai.HostedApplicationIam
	DisableNotFoundRetries bool
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai.HostedApplicationIamLifecycleStateCreating),
	}
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai.HostedApplicationIamLifecycleStateActive),
	}
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai.HostedApplicationIamLifecycleStateDeleting),
	}
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai.HostedApplicationIamLifecycleStateDeleted),
	}
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_generative_ai.CreateHostedApplicationIamRequest{}

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

	if environmentVariables, ok := s.D.GetOkExists("environment_variables"); ok {
		interfaces := environmentVariables.([]interface{})
		tmp := make([]oci_generative_ai.EnvironmentVariable, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "environment_variables", stateDataIndex)
			converted, err := s.mapToEnvironmentVariable(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("environment_variables") {
			request.EnvironmentVariables = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if networkingConfig, ok := s.D.GetOkExists("networking_config"); ok {
		if tmpList := networkingConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "networking_config", 0)
			tmp, err := s.mapToNetworkingConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkingConfig = &tmp
		}
	}

	if scalingConfig, ok := s.D.GetOkExists("scaling_config"); ok {
		if tmpList := scalingConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scaling_config", 0)
			tmp, err := s.mapToScalingConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ScalingConfig = &tmp
		}
	}

	if storageConfigs, ok := s.D.GetOkExists("storage_configs"); ok {
		interfaces := storageConfigs.([]interface{})
		tmp := make([]oci_generative_ai.StorageConfig, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "storage_configs", stateDataIndex)
			converted, err := s.mapToStorageConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("storage_configs") {
			request.StorageConfigs = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.CreateHostedApplicationIam(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getHostedApplicationIamFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) getHostedApplicationIamFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	hostedApplicationIamId, err := hostedApplicationIamWaitForWorkRequest(ctx, workId, "hostedapplication",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*hostedApplicationIamId)

	return s.GetWithContext(ctx)
}

func hostedApplicationIamWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "generative_ai", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_generative_ai.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func hostedApplicationIamWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_generative_ai.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai.GenerativeAiClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai")
	retryPolicy.ShouldRetryOperation = hostedApplicationIamWorkRequestShouldRetryFunc(timeout)

	response := oci_generative_ai.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_generative_ai.OperationStatusInProgress),
			string(oci_generative_ai.OperationStatusAccepted),
			string(oci_generative_ai.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_generative_ai.OperationStatusSucceeded),
			string(oci_generative_ai.OperationStatusFailed),
			string(oci_generative_ai.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_generative_ai.GetWorkRequestRequest{
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
	normalizedExpectedEntityType := strings.ReplaceAll(strings.ToLower(entityType), "_", "")
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if res.EntityType == nil {
			continue
		}
		normalizedEntityType := strings.ReplaceAll(strings.ToLower(*res.EntityType), "_", "")
		if strings.Contains(normalizedEntityType, normalizedExpectedEntityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_generative_ai.OperationStatusFailed || response.Status == oci_generative_ai.OperationStatusCanceled {
		return nil, getErrorFromGenerativeAiHostedApplicationIamWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiHostedApplicationIamWorkRequest(ctx context.Context, client *oci_generative_ai.GenerativeAiClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_generative_ai.ListWorkRequestErrorsRequest{
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

func (s *GenerativeAiHostedApplicationIamResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.GetHostedApplicationIamRequest{}

	tmp := s.D.Id()
	request.HostedApplicationIamId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.GetHostedApplicationIam(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.HostedApplicationIam
	return nil
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_generative_ai.UpdateHostedApplicationIamRequest{}

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

	if environmentVariables, ok := s.D.GetOkExists("environment_variables"); ok {
		interfaces := environmentVariables.([]interface{})
		tmp := make([]oci_generative_ai.EnvironmentVariable, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "environment_variables", stateDataIndex)
			converted, err := s.mapToEnvironmentVariable(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("environment_variables") {
			request.EnvironmentVariables = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.HostedApplicationIamId = &tmp

	if scalingConfig, ok := s.D.GetOkExists("scaling_config"); ok {
		if tmpList := scalingConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scaling_config", 0)
			tmp, err := s.mapToScalingConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ScalingConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.UpdateHostedApplicationIam(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getHostedApplicationIamFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_generative_ai.DeleteHostedApplicationIamRequest{}

	tmp := s.D.Id()
	request.HostedApplicationIamId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.DeleteHostedApplicationIam(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := hostedApplicationIamWaitForWorkRequest(ctx, workId, "hostedapplication",
		oci_generative_ai.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) SetData() error {
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

	environmentVariables := []interface{}{}
	for _, item := range s.Res.EnvironmentVariables {
		environmentVariables = append(environmentVariables, EnvironmentVariableToMap(item))
	}
	s.D.Set("environment_variables", environmentVariables)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NetworkingConfig != nil {
		s.D.Set("networking_config", []interface{}{NetworkingConfigToMap(s.Res.NetworkingConfig, false)})
	} else {
		s.D.Set("networking_config", nil)
	}

	if s.Res.ScalingConfig != nil {
		s.D.Set("scaling_config", []interface{}{ScalingConfigToMap(s.Res.ScalingConfig)})
	} else {
		s.D.Set("scaling_config", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	storageConfigs := []interface{}{}
	for _, item := range s.Res.StorageConfigs {
		storageConfigs = append(storageConfigs, StorageConfigToMap(item))
	}
	s.D.Set("storage_configs", storageConfigs)

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

func (s *GenerativeAiHostedApplicationIamResourceCrud) mapToEnvironmentVariable(fieldKeyFormat string) (oci_generative_ai.EnvironmentVariable, error) {
	result := oci_generative_ai.EnvironmentVariable{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_generative_ai.EnvironmentVariableTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		var tmp interface{}
		if err := json.Unmarshal([]byte(value.(string)), &tmp); err != nil {
			return result, err
		}
		result.Value = &tmp
	}

	return result, nil
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) mapToInboundNetworkingConfig(fieldKeyFormat string) (oci_generative_ai.InboundNetworkingConfig, error) {
	result := oci_generative_ai.InboundNetworkingConfig{}

	if endpointMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "endpoint_mode")); ok {
		result.EndpointMode = oci_generative_ai.InboundNetworkingConfigEndpointModeEnum(endpointMode.(string))
	}

	if privateEndpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_endpoint_id")); ok {
		tmp := privateEndpointId.(string)
		result.PrivateEndpointId = &tmp
	}

	return result, nil
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) mapToNetworkingConfig(fieldKeyFormat string) (oci_generative_ai.NetworkingConfig, error) {
	result := oci_generative_ai.NetworkingConfig{}

	if inboundNetworkingConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "inbound_networking_config")); ok {
		if tmpList := inboundNetworkingConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "inbound_networking_config"), 0)
			tmp, err := s.mapToInboundNetworkingConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert inbound_networking_config, encountered error: %v", err)
			}
			result.InboundNetworkingConfig = &tmp
		}
	}

	if outboundNetworkingConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "outbound_networking_config")); ok {
		if tmpList := outboundNetworkingConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "outbound_networking_config"), 0)
			tmp, err := s.mapToOutboundNetworkingConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert outbound_networking_config, encountered error: %v", err)
			}
			result.OutboundNetworkingConfig = &tmp
		}
	}

	return result, nil
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) mapToOutboundNetworkingConfig(fieldKeyFormat string) (oci_generative_ai.OutboundNetworkingConfig, error) {
	result := oci_generative_ai.OutboundNetworkingConfig{}

	if customSubnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_subnet_id")); ok {
		tmp := customSubnetId.(string)
		result.CustomSubnetId = &tmp
	}

	if networkMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_mode")); ok {
		result.NetworkMode = oci_generative_ai.OutboundNetworkingConfigNetworkModeEnum(networkMode.(string))
	}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
			result.NsgIds = tmp
		}
	}

	return result, nil
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) mapToScalingConfig(fieldKeyFormat string) (oci_generative_ai.ScalingConfig, error) {
	result := oci_generative_ai.ScalingConfig{}

	if maxReplica, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_replica")); ok {
		tmp := maxReplica.(int)
		result.MaxReplica = &tmp
	}

	if minReplica, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_replica")); ok {
		tmp := minReplica.(int)
		result.MinReplica = &tmp
	}

	if scalingType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scaling_type")); ok {
		result.ScalingType = oci_generative_ai.ScalingConfigScalingTypeEnum(scalingType.(string))
	}

	if targetConcurrencyThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_concurrency_threshold")); ok {
		tmp := targetConcurrencyThreshold.(int)
		if tmp > 0 {
			result.TargetConcurrencyThreshold = &tmp
		}
	}

	if targetCpuThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_cpu_threshold")); ok {
		tmp := targetCpuThreshold.(int)
		if tmp > 0 {
			result.TargetCpuThreshold = &tmp
		}
	}

	if targetMemoryThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_memory_threshold")); ok {
		tmp := targetMemoryThreshold.(int)
		if tmp > 0 {
			result.TargetMemoryThreshold = &tmp
		}
	}

	if targetRpsThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_rps_threshold")); ok {
		tmp := targetRpsThreshold.(int)
		if tmp > 0 {
			result.TargetRpsThreshold = &tmp
		}
	}

	return result, nil
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) mapToStorageConfig(fieldKeyFormat string) (oci_generative_ai.StorageConfig, error) {
	result := oci_generative_ai.StorageConfig{}

	if environmentVariableKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variable_key")); ok {
		tmp := environmentVariableKey.(string)
		result.EnvironmentVariableKey = &tmp
	}

	if storageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_id")); ok {
		tmp := storageId.(string)
		result.StorageId = &tmp
	}

	return result, nil
}

func (s *GenerativeAiHostedApplicationIamResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_generative_ai.ChangeHostedApplicationIamCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.HostedApplicationIamId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	_, err := s.Client.ChangeHostedApplicationIamCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
