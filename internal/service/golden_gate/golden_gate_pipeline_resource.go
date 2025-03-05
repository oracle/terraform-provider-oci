// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGatePipelineResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createGoldenGatePipeline,
		Read:     readGoldenGatePipeline,
		Update:   updateGoldenGatePipeline,
		Delete:   deleteGoldenGatePipeline,
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
			"license_model": {
				Type:     schema.TypeString,
				Required: true,
			},
			"recipe_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"ZERO_ETL",
				}, true),
			},
			"source_connection_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"connection_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"target_connection_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"connection_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
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
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"message": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"process_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"initial_data_load": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"is_initial_load": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"action_on_existing_table": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"replicate_schema_change": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"can_replicate_schema_change": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"action_on_ddl_error": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action_on_dml_error": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"should_restart_on_failure": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"start_using_default_mapping": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_auto_scaling_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_sub_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mapping_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"mapping_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"pipeline_diagnostic_data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"bucket": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"diagnostic_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_collected": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"time_last_recorded": {
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

func createGoldenGatePipeline(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGatePipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.CreateResource(d, sync)
}

func readGoldenGatePipeline(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGatePipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

func updateGoldenGatePipeline(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGatePipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteGoldenGatePipeline(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGatePipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GoldenGatePipelineResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_golden_gate.GoldenGateClient
	Res                    *oci_golden_gate.Pipeline
	DisableNotFoundRetries bool
}

func (s *GoldenGatePipelineResourceCrud) ID() string {
	pipeline := *s.Res
	return *pipeline.GetId()
}

func (s *GoldenGatePipelineResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_golden_gate.PipelineLifecycleStateCreating),
	}
}

func (s *GoldenGatePipelineResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_golden_gate.PipelineLifecycleStateActive),
		string(oci_golden_gate.PipelineLifecycleStateNeedsAttention),
	}
}

func (s *GoldenGatePipelineResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_golden_gate.PipelineLifecycleStateDeleting),
	}
}

func (s *GoldenGatePipelineResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_golden_gate.PipelineLifecycleStateDeleted),
	}
}

func (s *GoldenGatePipelineResourceCrud) Create() error {
	request := oci_golden_gate.CreatePipelineRequest{}
	err := s.populateTopLevelPolymorphicCreatePipelineRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.CreatePipeline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getPipelineFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GoldenGatePipelineResourceCrud) getPipelineFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_golden_gate.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	pipelineId, err := pipelineWaitForWorkRequest(workId, "pipeline",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*pipelineId)

	return s.Get()
}

func pipelineWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func pipelineWaitForWorkRequest(wId *string, entityType string, action oci_golden_gate.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_golden_gate.GoldenGateClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "golden_gate")
	retryPolicy.ShouldRetryOperation = pipelineWorkRequestShouldRetryFunc(timeout)

	response := oci_golden_gate.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_golden_gate.OperationStatusInProgress),
			string(oci_golden_gate.OperationStatusAccepted),
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
	if identifier == nil || response.Status == oci_golden_gate.OperationStatusFailed || response.Status == oci_golden_gate.OperationStatusCanceled {
		return nil, getErrorFromGoldenGatePipelineWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGoldenGatePipelineWorkRequest(client *oci_golden_gate.GoldenGateClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_golden_gate.ActionTypeEnum) error {
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

func (s *GoldenGatePipelineResourceCrud) Get() error {
	request := oci_golden_gate.GetPipelineRequest{}

	tmp := s.D.Id()
	request.PipelineId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.GetPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Pipeline
	return nil
}

func (s *GoldenGatePipelineResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_golden_gate.UpdatePipelineRequest{}
	err := s.populateTopLevelPolymorphicUpdatePipelineRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.UpdatePipeline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPipelineFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GoldenGatePipelineResourceCrud) Delete() error {
	request := oci_golden_gate.DeletePipelineRequest{}

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	tmp := s.D.Id()
	request.PipelineId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.DeletePipeline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := pipelineWaitForWorkRequest(workId, "pipeline",
		oci_golden_gate.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GoldenGatePipelineResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_golden_gate.ZeroEtlPipeline:
		s.D.Set("recipe_type", "ZERO_ETL")

		mappingRules := []interface{}{}
		for _, item := range v.MappingRules {
			mappingRules = append(mappingRules, MappingRuleToMap(item))
		}
		s.D.Set("mapping_rules", mappingRules)

		if v.ProcessOptions != nil {
			s.D.Set("process_options", []interface{}{ProcessOptionsToMap(v.ProcessOptions)})
		} else {
			s.D.Set("process_options", nil)
		}

		if v.TimeLastRecorded != nil {
			s.D.Set("time_last_recorded", v.TimeLastRecorded.Format(time.RFC3339Nano))
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.CpuCoreCount != nil {
			s.D.Set("cpu_core_count", *v.CpuCoreCount)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		if v.IsAutoScalingEnabled != nil {
			s.D.Set("is_auto_scaling_enabled", *v.IsAutoScalingEnabled)
		}

		s.D.Set("license_model", v.LicenseModel)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("lifecycle_sub_state", v.LifecycleSubState)

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		if v.PipelineDiagnosticData != nil {
			s.D.Set("pipeline_diagnostic_data", []interface{}{PipelineDiagnosticDataToMap(v.PipelineDiagnosticData)})
		} else {
			s.D.Set("pipeline_diagnostic_data", nil)
		}

		if v.SourceConnectionDetails != nil {
			s.D.Set("source_connection_details", []interface{}{SourcePipelineConnectionDetailsToMap(v.SourceConnectionDetails)})
		} else {
			s.D.Set("source_connection_details", nil)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetConnectionDetails != nil {
			s.D.Set("target_connection_details", []interface{}{TargetPipelineConnectionDetailsToMap(v.TargetConnectionDetails)})
		} else {
			s.D.Set("target_connection_details", nil)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'recipe_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *GoldenGatePipelineResourceCrud) mapToInitialDataLoad(fieldKeyFormat string) (oci_golden_gate.InitialDataLoad, error) {
	result := oci_golden_gate.InitialDataLoad{}

	if actionOnExistingTable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_on_existing_table")); ok {
		result.ActionOnExistingTable = oci_golden_gate.InitialLoadActionEnum(actionOnExistingTable.(string))
	}

	if isInitialLoad, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_initial_load")); ok {
		result.IsInitialLoad = oci_golden_gate.InitialDataLoadIsInitialLoadEnum(isInitialLoad.(string))
	}

	return result, nil
}

func InitialDataLoadToMap(obj *oci_golden_gate.InitialDataLoad) map[string]interface{} {
	result := map[string]interface{}{}

	result["action_on_existing_table"] = string(obj.ActionOnExistingTable)

	result["is_initial_load"] = string(obj.IsInitialLoad)

	return result
}

func (s *GoldenGatePipelineResourceCrud) mapToMappingRule(fieldKeyFormat string) (oci_golden_gate.MappingRule, error) {
	result := oci_golden_gate.MappingRule{}

	if mappingType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mapping_type")); ok {
		result.MappingType = oci_golden_gate.MappingTypeEnum(mappingType.(string))
	}

	if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
		tmp := source.(string)
		result.Source = &tmp
	}

	if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
		tmp := target.(string)
		result.Target = &tmp
	}

	return result, nil
}

func MappingRuleToMap(obj oci_golden_gate.MappingRule) map[string]interface{} {
	result := map[string]interface{}{}

	result["mapping_type"] = string(obj.MappingType)

	if obj.Source != nil {
		result["source"] = string(*obj.Source)
	}

	if obj.Target != nil {
		result["target"] = string(*obj.Target)
	}

	return result
}

func PipelineDiagnosticDataToMap(obj *oci_golden_gate.PipelineDiagnosticData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	result["diagnostic_state"] = string(obj.DiagnosticState)

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	if obj.TimeLastCollected != nil {
		result["time_last_collected"] = obj.TimeLastCollected.String()
	}

	return result
}

func PipelineSummaryToMap(obj oci_golden_gate.PipelineSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_golden_gate.ZeroEtlPipelineSummary:
		result["recipe_type"] = "ZERO_ETL"

		if v.ProcessOptions != nil {
			result["process_options"] = []interface{}{ProcessOptionsToMap(v.ProcessOptions)}
		}

		if v.TimeLastRecorded != nil {
			result["time_last_recorded"] = v.TimeLastRecorded.Format(time.RFC3339Nano)
		}
	default:
		log.Printf("[WARN] Received 'recipe_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *GoldenGatePipelineResourceCrud) mapToProcessOptions(fieldKeyFormat string) (oci_golden_gate.ProcessOptions, error) {
	result := oci_golden_gate.ProcessOptions{}

	if initialDataLoad, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "initial_data_load")); ok {
		if tmpList := initialDataLoad.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "initial_data_load"), 0)
			tmp, err := s.mapToInitialDataLoad(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert initial_data_load, encountered error: %v", err)
			}
			result.InitialDataLoad = &tmp
		}
	}

	if replicateSchemaChange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replicate_schema_change")); ok {
		if tmpList := replicateSchemaChange.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "replicate_schema_change"), 0)
			tmp, err := s.mapToReplicateSchemaChange(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert replicate_schema_change, encountered error: %v", err)
			}
			result.ReplicateSchemaChange = &tmp
		}
	}

	if shouldRestartOnFailure, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_restart_on_failure")); ok {
		result.ShouldRestartOnFailure = oci_golden_gate.ProcessOptionsShouldRestartOnFailureEnum(shouldRestartOnFailure.(string))
	}

	if startUsingDefaultMapping, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_using_default_mapping")); ok {
		result.StartUsingDefaultMapping = oci_golden_gate.ProcessOptionsStartUsingDefaultMappingEnum(startUsingDefaultMapping.(string))
	}

	return result, nil
}

func ProcessOptionsToMap(obj *oci_golden_gate.ProcessOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InitialDataLoad != nil {
		result["initial_data_load"] = []interface{}{InitialDataLoadToMap(obj.InitialDataLoad)}
	}

	if obj.ReplicateSchemaChange != nil {
		result["replicate_schema_change"] = []interface{}{ReplicateSchemaChangeToMap(obj.ReplicateSchemaChange)}
	}

	result["should_restart_on_failure"] = string(obj.ShouldRestartOnFailure)

	result["start_using_default_mapping"] = string(obj.StartUsingDefaultMapping)

	return result
}

func (s *GoldenGatePipelineResourceCrud) mapToReplicateSchemaChange(fieldKeyFormat string) (oci_golden_gate.ReplicateSchemaChange, error) {
	result := oci_golden_gate.ReplicateSchemaChange{}

	if actionOnDdlError, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_on_ddl_error")); ok {
		result.ActionOnDdlError = oci_golden_gate.ReplicateDdlErrorActionEnum(actionOnDdlError.(string))
	}

	if actionOnDmlError, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_on_dml_error")); ok {
		result.ActionOnDmlError = oci_golden_gate.ReplicateDmlErrorActionEnum(actionOnDmlError.(string))
	}

	if canReplicateSchemaChange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "can_replicate_schema_change")); ok {
		result.CanReplicateSchemaChange = oci_golden_gate.ReplicateSchemaChangeCanReplicateSchemaChangeEnum(canReplicateSchemaChange.(string))
	}

	return result, nil
}

func ReplicateSchemaChangeToMap(obj *oci_golden_gate.ReplicateSchemaChange) map[string]interface{} {
	result := map[string]interface{}{}

	result["action_on_ddl_error"] = string(obj.ActionOnDdlError)

	result["action_on_dml_error"] = string(obj.ActionOnDmlError)

	result["can_replicate_schema_change"] = string(obj.CanReplicateSchemaChange)

	return result
}

func (s *GoldenGatePipelineResourceCrud) mapToResourceLock(fieldKeyFormat string) (oci_golden_gate.ResourceLock, error) {
	result := oci_golden_gate.ResourceLock{}

	if message, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message")); ok {
		tmp := message.(string)
		result.Message = &tmp
	}

	if relatedResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "related_resource_id")); ok {
		tmp := relatedResourceId.(string)
		result.RelatedResourceId = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_golden_gate.ResourceLockTypeEnum(type_.(string))
	}

	return result, nil
}

func (s *GoldenGatePipelineResourceCrud) mapToSourcePipelineConnectionDetails(fieldKeyFormat string) (oci_golden_gate.SourcePipelineConnectionDetails, error) {
	result := oci_golden_gate.SourcePipelineConnectionDetails{}

	if connectionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_id")); ok {
		tmp := connectionId.(string)
		result.ConnectionId = &tmp
	}

	return result, nil
}

func SourcePipelineConnectionDetailsToMap(obj *oci_golden_gate.SourcePipelineConnectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectionId != nil {
		result["connection_id"] = string(*obj.ConnectionId)
	}

	return result
}

func (s *GoldenGatePipelineResourceCrud) mapToTargetPipelineConnectionDetails(fieldKeyFormat string) (oci_golden_gate.TargetPipelineConnectionDetails, error) {
	result := oci_golden_gate.TargetPipelineConnectionDetails{}

	if connectionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_id")); ok {
		tmp := connectionId.(string)
		result.ConnectionId = &tmp
	}

	return result, nil
}

func TargetPipelineConnectionDetailsToMap(obj *oci_golden_gate.TargetPipelineConnectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectionId != nil {
		result["connection_id"] = string(*obj.ConnectionId)
	}

	return result
}

func (s *GoldenGatePipelineResourceCrud) populateTopLevelPolymorphicCreatePipelineRequest(request *oci_golden_gate.CreatePipelineRequest) error {
	//discriminator
	recipeTypeRaw, ok := s.D.GetOkExists("recipe_type")
	var recipeType string
	if ok {
		recipeType = recipeTypeRaw.(string)
	} else {
		recipeType = "" // default value
	}
	switch strings.ToLower(recipeType) {
	case strings.ToLower("ZERO_ETL"):
		details := oci_golden_gate.CreateZeroEtlPipelineDetails{}
		if processOptions, ok := s.D.GetOkExists("process_options"); ok {
			if tmpList := processOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "process_options", 0)
				tmp, err := s.mapToProcessOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ProcessOptions = &tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_golden_gate.LicenseModelEnum(licenseModel.(string))
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.ResourceLock, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToResourceLock(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if sourceConnectionDetails, ok := s.D.GetOkExists("source_connection_details"); ok {
			if tmpList := sourceConnectionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_connection_details", 0)
				tmp, err := s.mapToSourcePipelineConnectionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.SourceConnectionDetails = &tmp
			}
		}
		if targetConnectionDetails, ok := s.D.GetOkExists("target_connection_details"); ok {
			if tmpList := targetConnectionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_connection_details", 0)
				tmp, err := s.mapToTargetPipelineConnectionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TargetConnectionDetails = &tmp
			}
		}
		request.CreatePipelineDetails = details
	default:
		return fmt.Errorf("unknown recipe_type '%v' was specified", recipeType)
	}
	return nil
}

func (s *GoldenGatePipelineResourceCrud) populateTopLevelPolymorphicUpdatePipelineRequest(request *oci_golden_gate.UpdatePipelineRequest) error {
	//discriminator
	recipeTypeRaw, ok := s.D.GetOkExists("recipe_type")
	var recipeType string
	if ok {
		recipeType = recipeTypeRaw.(string)
	} else {
		recipeType = "" // default value
	}
	switch strings.ToLower(recipeType) {
	case strings.ToLower("ZERO_ETL"):
		details := oci_golden_gate.UpdateZeroEtlPipelineDetails{}
		if mappingRules, ok := s.D.GetOkExists("mapping_rules"); ok {
			interfaces := mappingRules.([]interface{})
			tmp := make([]oci_golden_gate.MappingRule, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mapping_rules", stateDataIndex)
				converted, err := s.mapToMappingRule(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("mapping_rules") {
				details.MappingRules = tmp
			}
		}
		if processOptions, ok := s.D.GetOkExists("process_options"); ok {
			if tmpList := processOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "process_options", 0)
				tmp, err := s.mapToProcessOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ProcessOptions = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_golden_gate.LicenseModelEnum(licenseModel.(string))
		}
		tmp := s.D.Id()
		request.PipelineId = &tmp
		request.UpdatePipelineDetails = details
	default:
		return fmt.Errorf("unknown recipe_type '%v' was specified", recipeType)
	}
	return nil
}

func (s *GoldenGatePipelineResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_golden_gate.ChangePipelineCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	idTmp := s.D.Id()
	changeCompartmentRequest.PipelineId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.ChangePipelineCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPipelineFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
