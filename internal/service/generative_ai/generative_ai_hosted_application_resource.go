// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"
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

func GenerativeAiHostedApplicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createGenerativeAiHostedApplicationWithContext,
		ReadContext:   readGenerativeAiHostedApplicationWithContext,
		UpdateContext: updateGenerativeAiHostedApplicationWithContext,
		DeleteContext: deleteGenerativeAiHostedApplicationWithContext,
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
			"inbound_auth_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"inbound_auth_config_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"idcs_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"domain_url": {
										Type:     schema.TypeString,
										Required: true,
									},
									"scope": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"audience": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
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

func createGenerativeAiHostedApplicationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readGenerativeAiHostedApplicationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateGenerativeAiHostedApplicationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteGenerativeAiHostedApplicationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type GenerativeAiHostedApplicationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai.GenerativeAiClient
	Res                    *oci_generative_ai.HostedApplication
	DisableNotFoundRetries bool
}

func (s *GenerativeAiHostedApplicationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiHostedApplicationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai.HostedApplicationLifecycleStateCreating),
	}
}

func (s *GenerativeAiHostedApplicationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai.HostedApplicationLifecycleStateActive),
	}
}

func (s *GenerativeAiHostedApplicationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai.HostedApplicationLifecycleStateDeleting),
	}
}

func (s *GenerativeAiHostedApplicationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai.HostedApplicationLifecycleStateDeleted),
	}
}

func (s *GenerativeAiHostedApplicationResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_generative_ai.CreateHostedApplicationRequest{}

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

	if inboundAuthConfig, ok := s.D.GetOkExists("inbound_auth_config"); ok {
		if tmpList := inboundAuthConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "inbound_auth_config", 0)
			tmp, err := s.mapToInboundAuthConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InboundAuthConfig = &tmp
		}
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

	response, err := s.Client.CreateHostedApplication(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getHostedApplicationFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiHostedApplicationResourceCrud) getHostedApplicationFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	hostedApplicationId, err := hostedApplicationWaitForWorkRequest(ctx, workId, "hostedapplication",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*hostedApplicationId)

	return s.GetWithContext(ctx)
}

func hostedApplicationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func hostedApplicationWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_generative_ai.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai.GenerativeAiClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai")
	retryPolicy.ShouldRetryOperation = hostedApplicationWorkRequestShouldRetryFunc(timeout)

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
	if identifier == nil || response.Status == oci_generative_ai.OperationStatusFailed || response.Status == oci_generative_ai.OperationStatusCanceled {
		return nil, getErrorFromGenerativeAiHostedApplicationWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiHostedApplicationWorkRequest(ctx context.Context, client *oci_generative_ai.GenerativeAiClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai.ActionTypeEnum) error {
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

func (s *GenerativeAiHostedApplicationResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.GetHostedApplicationRequest{}

	tmp := s.D.Id()
	request.HostedApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.GetHostedApplication(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.HostedApplication
	return nil
}

func (s *GenerativeAiHostedApplicationResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_generative_ai.UpdateHostedApplicationRequest{}

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
	request.HostedApplicationId = &tmp

	if inboundAuthConfig, ok := s.D.GetOkExists("inbound_auth_config"); ok {
		if tmpList := inboundAuthConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "inbound_auth_config", 0)
			tmp, err := s.mapToInboundAuthConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InboundAuthConfig = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.UpdateHostedApplication(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getHostedApplicationFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GenerativeAiHostedApplicationResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_generative_ai.DeleteHostedApplicationRequest{}

	tmp := s.D.Id()
	request.HostedApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.DeleteHostedApplication(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := hostedApplicationWaitForWorkRequest(ctx, workId, "hostedapplication",
		oci_generative_ai.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiHostedApplicationResourceCrud) SetData() error {
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

	if s.Res.InboundAuthConfig != nil {
		s.D.Set("inbound_auth_config", []interface{}{InboundAuthConfigToMap(s.Res.InboundAuthConfig)})
	} else {
		s.D.Set("inbound_auth_config", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NetworkingConfig != nil {
		s.D.Set("networking_config", []interface{}{NetworkingConfigToMap(s.Res.NetworkingConfig)})
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

func (s *GenerativeAiHostedApplicationResourceCrud) mapToEnvironmentVariable(fieldKeyFormat string) (oci_generative_ai.EnvironmentVariable, error) {
	result := oci_generative_ai.EnvironmentVariable{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_generative_ai.EnvironmentVariableTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := interface{}(value.(string))
		result.Value = &tmp
	}

	return result, nil
}

func EnvironmentVariableToMap(obj oci_generative_ai.EnvironmentVariable) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		if value, ok := (*obj.Value).(string); ok {
			result["value"] = value
		} else {
			result["value"] = *obj.Value
		}
	}

	return result
}

func HostedApplicationSummaryToMap(obj oci_generative_ai.HostedApplicationSummary) map[string]interface{} {
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

func (s *GenerativeAiHostedApplicationResourceCrud) mapToIdcsAuthConfig(fieldKeyFormat string) (oci_generative_ai.IdcsAuthConfig, error) {
	result := oci_generative_ai.IdcsAuthConfig{}

	if audience, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "audience")); ok {
		tmp := audience.(string)
		result.Audience = &tmp
	}

	if domainUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_url")); ok {
		tmp := domainUrl.(string)
		result.DomainUrl = &tmp
	}

	if scope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope")); ok {
		tmp := scope.(string)
		result.Scope = &tmp
	}

	return result, nil
}

func IdcsAuthConfigToMap(obj *oci_generative_ai.IdcsAuthConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Audience != nil {
		result["audience"] = string(*obj.Audience)
	}

	if obj.DomainUrl != nil {
		result["domain_url"] = string(*obj.DomainUrl)
	}

	if obj.Scope != nil {
		result["scope"] = string(*obj.Scope)
	}

	return result
}

func (s *GenerativeAiHostedApplicationResourceCrud) mapToInboundAuthConfig(fieldKeyFormat string) (oci_generative_ai.InboundAuthConfig, error) {
	result := oci_generative_ai.InboundAuthConfig{}

	if idcsConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idcs_config")); ok {
		if tmpList := idcsConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "idcs_config"), 0)
			tmp, err := s.mapToIdcsAuthConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert idcs_config, encountered error: %v", err)
			}
			result.IdcsConfig = &tmp
		}
	}

	if inboundAuthConfigType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "inbound_auth_config_type")); ok {
		result.InboundAuthConfigType = oci_generative_ai.InboundAuthConfigInboundAuthConfigTypeEnum(inboundAuthConfigType.(string))
	}

	return result, nil
}

func InboundAuthConfigToMap(obj *oci_generative_ai.InboundAuthConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IdcsConfig != nil {
		result["idcs_config"] = []interface{}{IdcsAuthConfigToMap(obj.IdcsConfig)}
	}

	result["inbound_auth_config_type"] = string(obj.InboundAuthConfigType)

	return result
}

func (s *GenerativeAiHostedApplicationResourceCrud) mapToInboundNetworkingConfig(fieldKeyFormat string) (oci_generative_ai.InboundNetworkingConfig, error) {
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

func InboundNetworkingConfigToMap(obj *oci_generative_ai.InboundNetworkingConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["endpoint_mode"] = string(obj.EndpointMode)

	if obj.PrivateEndpointId != nil {
		result["private_endpoint_id"] = string(*obj.PrivateEndpointId)
	}

	return result
}

func (s *GenerativeAiHostedApplicationResourceCrud) mapToNetworkingConfig(fieldKeyFormat string) (oci_generative_ai.NetworkingConfig, error) {
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

func NetworkingConfigToMap(obj *oci_generative_ai.NetworkingConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InboundNetworkingConfig != nil {
		result["inbound_networking_config"] = []interface{}{InboundNetworkingConfigToMap(obj.InboundNetworkingConfig)}
	}

	if obj.OutboundNetworkingConfig != nil {
		result["outbound_networking_config"] = []interface{}{OutboundNetworkingConfigToMap(obj.OutboundNetworkingConfig, false)}
	}

	return result
}

func (s *GenerativeAiHostedApplicationResourceCrud) mapToOutboundNetworkingConfig(fieldKeyFormat string) (oci_generative_ai.OutboundNetworkingConfig, error) {
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

func OutboundNetworkingConfigToMap(obj *oci_generative_ai.OutboundNetworkingConfig, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomSubnetId != nil {
		result["custom_subnet_id"] = string(*obj.CustomSubnetId)
	}

	result["network_mode"] = string(obj.NetworkMode)

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	return result
}

func (s *GenerativeAiHostedApplicationResourceCrud) mapToScalingConfig(fieldKeyFormat string) (oci_generative_ai.ScalingConfig, error) {
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
		result.TargetConcurrencyThreshold = &tmp
	}

	if targetCpuThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_cpu_threshold")); ok {
		tmp := targetCpuThreshold.(int)
		result.TargetCpuThreshold = &tmp
	}

	if targetMemoryThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_memory_threshold")); ok {
		tmp := targetMemoryThreshold.(int)
		result.TargetMemoryThreshold = &tmp
	}

	if targetRpsThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_rps_threshold")); ok {
		tmp := targetRpsThreshold.(int)
		result.TargetRpsThreshold = &tmp
	}

	return result, nil
}

func ScalingConfigToMap(obj *oci_generative_ai.ScalingConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxReplica != nil {
		result["max_replica"] = int(*obj.MaxReplica)
	}

	if obj.MinReplica != nil {
		result["min_replica"] = int(*obj.MinReplica)
	}

	result["scaling_type"] = string(obj.ScalingType)

	if obj.TargetConcurrencyThreshold != nil {
		result["target_concurrency_threshold"] = int(*obj.TargetConcurrencyThreshold)
	}

	if obj.TargetCpuThreshold != nil {
		result["target_cpu_threshold"] = int(*obj.TargetCpuThreshold)
	}

	if obj.TargetMemoryThreshold != nil {
		result["target_memory_threshold"] = int(*obj.TargetMemoryThreshold)
	}

	if obj.TargetRpsThreshold != nil {
		result["target_rps_threshold"] = int(*obj.TargetRpsThreshold)
	}

	return result
}

func (s *GenerativeAiHostedApplicationResourceCrud) mapToStorageConfig(fieldKeyFormat string) (oci_generative_ai.StorageConfig, error) {
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

func StorageConfigToMap(obj oci_generative_ai.StorageConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EnvironmentVariableKey != nil {
		result["environment_variable_key"] = string(*obj.EnvironmentVariableKey)
	}

	if obj.StorageId != nil {
		result["storage_id"] = string(*obj.StorageId)
	}

	return result
}

func (s *GenerativeAiHostedApplicationResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_generative_ai.ChangeHostedApplicationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.HostedApplicationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	_, err := s.Client.ChangeHostedApplicationCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
