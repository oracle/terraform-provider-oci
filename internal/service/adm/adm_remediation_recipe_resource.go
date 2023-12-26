// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package adm

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AdmRemediationRecipeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAdmRemediationRecipe,
		Read:     readAdmRemediationRecipe,
		Update:   updateAdmRemediationRecipe,
		Delete:   deleteAdmRemediationRecipe,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"detect_configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"exclusions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"max_permissible_cvss_v2score": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"max_permissible_cvss_v3score": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"max_permissible_severity": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"upgrade_policy": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"is_run_triggered_on_kb_change": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"knowledge_base_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_configuration": {
				Type:     schema.TypeList,
				Required: true,
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
						"nsg_ids": {
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
			"scm_configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"branch": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_automerge_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"scm_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"EXTERNAL_SCM",
								"OCI_CODE_REPOSITORY",
							}, true),
						},

						// Optional
						"build_file_location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_scm_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oci_code_repository_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pat_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"repository_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"verify_configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"build_service_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"GITHUB_ACTIONS",
								"GITLAB_PIPELINE",
								"JENKINS_PIPELINE",
								"NONE",
								"OCI_DEVOPS_BUILD",
							}, true),
						},

						// Optional
						"additional_parameters": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"jenkins_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"job_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pat_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pipeline_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"repository_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trigger_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"workflow_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

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
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_adm.RemediationRecipeLifecycleStateInactive),
					string(oci_adm.RemediationRecipeLifecycleStateActive),
				}, true),
			},

			// Computed
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

func createAdmRemediationRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_adm.RemediationRecipeLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_adm.RemediationRecipeLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopRemediationRecipe(); err != nil {
			return err
		}
		sync.D.Set("state", oci_adm.RemediationRecipeLifecycleStateInactive)
	}
	return nil

}

func readAdmRemediationRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.ReadResource(sync)
}

func updateAdmRemediationRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_adm.RemediationRecipeLifecycleStateActive == oci_adm.RemediationRecipeLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_adm.RemediationRecipeLifecycleStateInactive == oci_adm.RemediationRecipeLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartRemediationRecipe(); err != nil {
			return err
		}
		sync.D.Set("state", oci_adm.RemediationRecipeLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopRemediationRecipe(); err != nil {
			return err
		}
		sync.D.Set("state", oci_adm.RemediationRecipeLifecycleStateInactive)
	}

	return nil
}

func deleteAdmRemediationRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AdmRemediationRecipeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_adm.ApplicationDependencyManagementClient
	Res                    *oci_adm.RemediationRecipe
	DisableNotFoundRetries bool
}

func (s *AdmRemediationRecipeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AdmRemediationRecipeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_adm.RemediationRecipeLifecycleStateCreating),
	}
}

func (s *AdmRemediationRecipeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_adm.RemediationRecipeLifecycleStateActive),
		string(oci_adm.RemediationRecipeLifecycleStateNeedsAttention),
	}
}

func (s *AdmRemediationRecipeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_adm.RemediationRecipeLifecycleStateDeleting),
	}
}

func (s *AdmRemediationRecipeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_adm.RemediationRecipeLifecycleStateDeleted),
	}
}

func (s *AdmRemediationRecipeResourceCrud) Create() error {
	request := oci_adm.CreateRemediationRecipeRequest{}

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

	if detectConfiguration, ok := s.D.GetOkExists("detect_configuration"); ok {
		if tmpList := detectConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "detect_configuration", 0)
			tmp, err := s.mapToDetectConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DetectConfiguration = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isRunTriggeredOnKbChange, ok := s.D.GetOkExists("is_run_triggered_on_kb_change"); ok {
		tmp := isRunTriggeredOnKbChange.(bool)
		request.IsRunTriggeredOnKbChange = &tmp
	}

	if knowledgeBaseId, ok := s.D.GetOkExists("knowledge_base_id"); ok {
		tmp := knowledgeBaseId.(string)
		request.KnowledgeBaseId = &tmp
	}

	if networkConfiguration, ok := s.D.GetOkExists("network_configuration"); ok {
		if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_configuration", 0)
			tmp, err := s.mapToNetworkConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkConfiguration = &tmp
		}
	}

	if scmConfiguration, ok := s.D.GetOkExists("scm_configuration"); ok {
		if tmpList := scmConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scm_configuration", 0)
			tmp, err := s.mapToScmConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ScmConfiguration = tmp
		}
	}

	if verifyConfiguration, ok := s.D.GetOkExists("verify_configuration"); ok {
		if tmpList := verifyConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "verify_configuration", 0)
			tmp, err := s.mapToVerifyConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.VerifyConfiguration = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.CreateRemediationRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_adm.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_adm.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "remediationrecipe") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getRemediationRecipeFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm"), oci_adm.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AdmRemediationRecipeResourceCrud) getRemediationRecipeFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_adm.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	remediationRecipeId, err := remediationRecipeWaitForWorkRequest(workId, "remediationrecipe",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, remediationRecipeId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_adm.CancelWorkRequestRequest{
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
	s.D.SetId(*remediationRecipeId)

	return s.Get()
}

func remediationRecipeWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "adm", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_adm.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func remediationRecipeWaitForWorkRequest(wId *string, entityType string, action oci_adm.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_adm.ApplicationDependencyManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "adm")
	retryPolicy.ShouldRetryOperation = remediationRecipeWorkRequestShouldRetryFunc(timeout)

	response := oci_adm.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_adm.OperationStatusInProgress),
			string(oci_adm.OperationStatusAccepted),
			string(oci_adm.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_adm.OperationStatusSucceeded),
			string(oci_adm.OperationStatusFailed),
			string(oci_adm.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_adm.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_adm.OperationStatusFailed || response.Status == oci_adm.OperationStatusCanceled {
		return nil, getErrorFromAdmRemediationRecipeWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAdmRemediationRecipeWorkRequest(client *oci_adm.ApplicationDependencyManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_adm.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_adm.ListWorkRequestErrorsRequest{
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

func (s *AdmRemediationRecipeResourceCrud) Get() error {
	request := oci_adm.GetRemediationRecipeRequest{}

	tmp := s.D.Id()
	request.RemediationRecipeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.GetRemediationRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RemediationRecipe
	return nil
}

func (s *AdmRemediationRecipeResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_adm.UpdateRemediationRecipeRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if detectConfiguration, ok := s.D.GetOkExists("detect_configuration"); ok {
		if tmpList := detectConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "detect_configuration", 0)
			tmp, err := s.mapToDetectConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DetectConfiguration = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isRunTriggeredOnKbChange, ok := s.D.GetOkExists("is_run_triggered_on_kb_change"); ok {
		tmp := isRunTriggeredOnKbChange.(bool)
		request.IsRunTriggeredOnKbChange = &tmp
	}

	if knowledgeBaseId, ok := s.D.GetOkExists("knowledge_base_id"); ok {
		tmp := knowledgeBaseId.(string)
		request.KnowledgeBaseId = &tmp
	}

	if networkConfiguration, ok := s.D.GetOkExists("network_configuration"); ok {
		if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_configuration", 0)
			tmp, err := s.mapToNetworkConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkConfiguration = &tmp
		}
	}

	tmp := s.D.Id()
	request.RemediationRecipeId = &tmp

	if scmConfiguration, ok := s.D.GetOkExists("scm_configuration"); ok {
		if tmpList := scmConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scm_configuration", 0)
			tmp, err := s.mapToScmConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ScmConfiguration = tmp
		}
	}

	if verifyConfiguration, ok := s.D.GetOkExists("verify_configuration"); ok {
		if tmpList := verifyConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "verify_configuration", 0)
			tmp, err := s.mapToVerifyConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.VerifyConfiguration = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.UpdateRemediationRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRemediationRecipeFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm"), oci_adm.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AdmRemediationRecipeResourceCrud) Delete() error {
	request := oci_adm.DeleteRemediationRecipeRequest{}

	tmp := s.D.Id()
	request.RemediationRecipeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.DeleteRemediationRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := remediationRecipeWaitForWorkRequest(workId, "remediationrecipe",
		oci_adm.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AdmRemediationRecipeResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DetectConfiguration != nil {
		s.D.Set("detect_configuration", []interface{}{DetectConfigurationToMap(s.Res.DetectConfiguration)})
	} else {
		s.D.Set("detect_configuration", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsRunTriggeredOnKbChange != nil {
		s.D.Set("is_run_triggered_on_kb_change", *s.Res.IsRunTriggeredOnKbChange)
	}

	if s.Res.KnowledgeBaseId != nil {
		s.D.Set("knowledge_base_id", *s.Res.KnowledgeBaseId)
	}

	if s.Res.NetworkConfiguration != nil {
		s.D.Set("network_configuration", []interface{}{NetworkConfigurationToMap(s.Res.NetworkConfiguration, false)})
	} else {
		s.D.Set("network_configuration", nil)
	}

	if s.Res.ScmConfiguration != nil {
		scmConfigurationArray := []interface{}{}
		if scmConfigurationMap := ScmConfigurationToMap(&s.Res.ScmConfiguration); scmConfigurationMap != nil {
			scmConfigurationArray = append(scmConfigurationArray, scmConfigurationMap)
		}
		s.D.Set("scm_configuration", scmConfigurationArray)
	} else {
		s.D.Set("scm_configuration", nil)
	}

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

	if s.Res.VerifyConfiguration != nil {
		verifyConfigurationArray := []interface{}{}
		if verifyConfigurationMap := VerifyConfigurationToMap(&s.Res.VerifyConfiguration); verifyConfigurationMap != nil {
			verifyConfigurationArray = append(verifyConfigurationArray, verifyConfigurationMap)
		}
		s.D.Set("verify_configuration", verifyConfigurationArray)
	} else {
		s.D.Set("verify_configuration", nil)
	}

	return nil
}

func (s *AdmRemediationRecipeResourceCrud) StartRemediationRecipe() error {
	request := oci_adm.ActivateRemediationRecipeRequest{}

	idTmp := s.D.Id()
	request.RemediationRecipeId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	_, err := s.Client.ActivateRemediationRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_adm.RemediationRecipeLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AdmRemediationRecipeResourceCrud) StopRemediationRecipe() error {
	request := oci_adm.DeactivateRemediationRecipeRequest{}

	idTmp := s.D.Id()
	request.RemediationRecipeId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	_, err := s.Client.DeactivateRemediationRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_adm.RemediationRecipeLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AdmRemediationRecipeResourceCrud) mapToDetectConfiguration(fieldKeyFormat string) (oci_adm.DetectConfiguration, error) {
	result := oci_adm.DetectConfiguration{}

	if exclusions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclusions")); ok {
		interfaces := exclusions.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclusions")) {
			result.Exclusions = tmp
		}
	}

	if maxPermissibleCvssV2Score, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_permissible_cvss_v2score")); ok {
		tmp := float32(maxPermissibleCvssV2Score.(float64))
		result.MaxPermissibleCvssV2Score = &tmp
	}

	if maxPermissibleCvssV3Score, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_permissible_cvss_v3score")); ok {
		tmp := float32(maxPermissibleCvssV3Score.(float64))
		result.MaxPermissibleCvssV3Score = &tmp
	}

	if maxPermissibleSeverity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_permissible_severity")); ok {
		result.MaxPermissibleSeverity = oci_adm.ConfigSeverityEnum(maxPermissibleSeverity.(string))
	}

	if upgradePolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "upgrade_policy")); ok {
		result.UpgradePolicy = oci_adm.DetectConfigurationUpgradePolicyEnum(upgradePolicy.(string))
	}

	return result, nil
}

func DetectConfigurationToMap(obj *oci_adm.DetectConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	result["exclusions"] = obj.Exclusions

	if obj.MaxPermissibleCvssV2Score != nil {
		result["max_permissible_cvss_v2score"] = float32(*obj.MaxPermissibleCvssV2Score)
	}

	if obj.MaxPermissibleCvssV3Score != nil {
		result["max_permissible_cvss_v3score"] = float32(*obj.MaxPermissibleCvssV3Score)
	}

	result["max_permissible_severity"] = string(obj.MaxPermissibleSeverity)

	result["upgrade_policy"] = string(obj.UpgradePolicy)

	return result
}

func (s *AdmRemediationRecipeResourceCrud) mapToNetworkConfiguration(fieldKeyFormat string) (oci_adm.NetworkConfiguration, error) {
	result := oci_adm.NetworkConfiguration{}

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

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func NetworkConfigurationToMap(obj *oci_adm.NetworkConfiguration, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func RemediationRecipeSummaryToMap(obj oci_adm.RemediationRecipeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsRunTriggeredOnKbChange != nil {
		result["is_run_triggered_on_kb_change"] = bool(*obj.IsRunTriggeredOnKbChange)
	}

	if obj.KnowledgeBaseId != nil {
		result["knowledge_base_id"] = string(*obj.KnowledgeBaseId)
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

func (s *AdmRemediationRecipeResourceCrud) mapToScmConfiguration(fieldKeyFormat string) (oci_adm.ScmConfiguration, error) {
	var baseObject oci_adm.ScmConfiguration
	//discriminator
	scmTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scm_type"))
	var scmType string
	if ok {
		scmType = scmTypeRaw.(string)
	} else {
		scmType = "" // default value
	}
	switch strings.ToLower(scmType) {
	case strings.ToLower("EXTERNAL_SCM"):
		details := oci_adm.ExternalScmConfiguration{}
		if externalScmType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "external_scm_type")); ok {
			details.ExternalScmType = oci_adm.ExternalScmConfigurationExternalScmTypeEnum(externalScmType.(string))
		}
		if patSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pat_secret_id")); ok {
			tmp := patSecretId.(string)
			details.PatSecretId = &tmp
		}
		if repositoryUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_url")); ok {
			tmp := repositoryUrl.(string)
			details.RepositoryUrl = &tmp
		}
		if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if branch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "branch")); ok {
			tmp := branch.(string)
			details.Branch = &tmp
		}
		if buildFileLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "build_file_location")); ok {
			tmp := buildFileLocation.(string)
			details.BuildFileLocation = &tmp
		}
		if isAutomergeEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_automerge_enabled")); ok {
			tmp := isAutomergeEnabled.(bool)
			details.IsAutomergeEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_CODE_REPOSITORY"):
		details := oci_adm.OciCodeRepositoryConfiguration{}
		if ociCodeRepositoryId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oci_code_repository_id")); ok {
			tmp := ociCodeRepositoryId.(string)
			details.OciCodeRepositoryId = &tmp
		}
		if branch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "branch")); ok {
			tmp := branch.(string)
			details.Branch = &tmp
		}
		if buildFileLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "build_file_location")); ok {
			tmp := buildFileLocation.(string)
			details.BuildFileLocation = &tmp
		}
		if isAutomergeEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_automerge_enabled")); ok {
			tmp := isAutomergeEnabled.(bool)
			details.IsAutomergeEnabled = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown scm_type '%v' was specified", scmType)
	}
	return baseObject, nil
}

func ScmConfigurationToMap(obj *oci_adm.ScmConfiguration) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_adm.ExternalScmConfiguration:
		result["scm_type"] = "EXTERNAL_SCM"

		result["external_scm_type"] = string(v.ExternalScmType)

		if v.PatSecretId != nil {
			result["pat_secret_id"] = string(*v.PatSecretId)
		}

		if v.RepositoryUrl != nil {
			result["repository_url"] = string(*v.RepositoryUrl)
		}

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}

		if v.Branch != nil {
			result["branch"] = string(*v.Branch)
		}

		if v.BuildFileLocation != nil {
			result["build_file_location"] = string(*v.BuildFileLocation)
		}

		if v.IsAutomergeEnabled != nil {
			result["is_automerge_enabled"] = bool(*v.IsAutomergeEnabled)
		}
	case oci_adm.OciCodeRepositoryConfiguration:
		result["scm_type"] = "OCI_CODE_REPOSITORY"

		if v.OciCodeRepositoryId != nil {
			result["oci_code_repository_id"] = string(*v.OciCodeRepositoryId)
		}

		if v.Branch != nil {
			result["branch"] = string(*v.Branch)
		}

		if v.BuildFileLocation != nil {
			result["build_file_location"] = string(*v.BuildFileLocation)
		}

		if v.IsAutomergeEnabled != nil {
			result["is_automerge_enabled"] = bool(*v.IsAutomergeEnabled)
		}
	default:
		log.Printf("[WARN] Received 'scm_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AdmRemediationRecipeResourceCrud) mapToVerifyConfiguration(fieldKeyFormat string) (oci_adm.VerifyConfiguration, error) {
	var baseObject oci_adm.VerifyConfiguration
	//discriminator
	buildServiceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "build_service_type"))
	var buildServiceType string
	if ok {
		buildServiceType = buildServiceTypeRaw.(string)
	} else {
		buildServiceType = "" // default value
	}
	switch strings.ToLower(buildServiceType) {
	case strings.ToLower("GITHUB_ACTIONS"):
		details := oci_adm.GitHubActionsConfiguration{}
		if additionalParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_parameters")); ok {
			details.AdditionalParameters = tfresource.ObjectMapToStringMap(additionalParameters.(map[string]interface{}))
		}
		if patSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pat_secret_id")); ok {
			tmp := patSecretId.(string)
			details.PatSecretId = &tmp
		}
		if repositoryUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_url")); ok {
			tmp := repositoryUrl.(string)
			details.RepositoryUrl = &tmp
		}
		if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if workflowName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "workflow_name")); ok {
			tmp := workflowName.(string)
			details.WorkflowName = &tmp
		}
		baseObject = details
	case strings.ToLower("GITLAB_PIPELINE"):
		details := oci_adm.GitLabPipelineConfiguration{}
		if additionalParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_parameters")); ok {
			details.AdditionalParameters = tfresource.ObjectMapToStringMap(additionalParameters.(map[string]interface{}))
		}
		if patSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pat_secret_id")); ok {
			tmp := patSecretId.(string)
			details.PatSecretId = &tmp
		}
		if repositoryUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_url")); ok {
			tmp := repositoryUrl.(string)
			details.RepositoryUrl = &tmp
		}
		if triggerSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trigger_secret_id")); ok {
			tmp := triggerSecretId.(string)
			details.TriggerSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		baseObject = details
	case strings.ToLower("JENKINS_PIPELINE"):
		details := oci_adm.JenkinsPipelineConfiguration{}
		if additionalParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_parameters")); ok {
			details.AdditionalParameters = tfresource.ObjectMapToStringMap(additionalParameters.(map[string]interface{}))
		}
		if jenkinsUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "jenkins_url")); ok {
			tmp := jenkinsUrl.(string)
			details.JenkinsUrl = &tmp
		}
		if jobName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_name")); ok {
			tmp := jobName.(string)
			details.JobName = &tmp
		}
		if patSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pat_secret_id")); ok {
			tmp := patSecretId.(string)
			details.PatSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_adm.NoneVerifyConfiguration{}
		baseObject = details
	case strings.ToLower("OCI_DEVOPS_BUILD"):
		details := oci_adm.OciDevOpsBuildConfiguration{}
		if additionalParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_parameters")); ok {
			details.AdditionalParameters = tfresource.ObjectMapToStringMap(additionalParameters.(map[string]interface{}))
		}
		if pipelineId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pipeline_id")); ok {
			tmp := pipelineId.(string)
			details.PipelineId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown build_service_type '%v' was specified", buildServiceType)
	}
	return baseObject, nil
}

func VerifyConfigurationToMap(obj *oci_adm.VerifyConfiguration) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_adm.GitHubActionsConfiguration:
		result["build_service_type"] = "GITHUB_ACTIONS"

		result["additional_parameters"] = v.AdditionalParameters

		if v.PatSecretId != nil {
			result["pat_secret_id"] = string(*v.PatSecretId)
		}

		if v.RepositoryUrl != nil {
			result["repository_url"] = string(*v.RepositoryUrl)
		}

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}

		if v.WorkflowName != nil {
			result["workflow_name"] = string(*v.WorkflowName)
		}
	case oci_adm.GitLabPipelineConfiguration:
		result["build_service_type"] = "GITLAB_PIPELINE"

		result["additional_parameters"] = v.AdditionalParameters

		if v.PatSecretId != nil {
			result["pat_secret_id"] = string(*v.PatSecretId)
		}

		if v.RepositoryUrl != nil {
			result["repository_url"] = string(*v.RepositoryUrl)
		}

		if v.TriggerSecretId != nil {
			result["trigger_secret_id"] = string(*v.TriggerSecretId)
		}

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_adm.JenkinsPipelineConfiguration:
		result["build_service_type"] = "JENKINS_PIPELINE"

		result["additional_parameters"] = v.AdditionalParameters

		if v.JenkinsUrl != nil {
			result["jenkins_url"] = string(*v.JenkinsUrl)
		}

		if v.JobName != nil {
			result["job_name"] = string(*v.JobName)
		}

		if v.PatSecretId != nil {
			result["pat_secret_id"] = string(*v.PatSecretId)
		}

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_adm.NoneVerifyConfiguration:
		result["build_service_type"] = "NONE"
	case oci_adm.OciDevOpsBuildConfiguration:
		result["build_service_type"] = "OCI_DEVOPS_BUILD"

		result["additional_parameters"] = v.AdditionalParameters

		if v.PipelineId != nil {
			result["pipeline_id"] = string(*v.PipelineId)
		}
	default:
		log.Printf("[WARN] Received 'build_service_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AdmRemediationRecipeResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_adm.ChangeRemediationRecipeCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.RemediationRecipeId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.ChangeRemediationRecipeCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRemediationRecipeFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm"), oci_adm.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
