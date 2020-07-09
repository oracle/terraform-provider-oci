// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v25/common"
	oci_integration "github.com/oracle/oci-go-sdk/v25/integration"
)

func init() {
	RegisterResource("oci_integration_integration_instance", IntegrationIntegrationInstanceResource())
}

func IntegrationIntegrationInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createIntegrationIntegrationInstance,
		Read:     readIntegrationIntegrationInstance,
		Update:   updateIntegrationIntegrationInstance,
		Delete:   deleteIntegrationIntegrationInstance,
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
			"integration_instance_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_byol": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"message_packs": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Optional
			"consumption_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"idcs_at": {
				Type:      schema.TypeString,
				Optional:  true,
				StateFunc: getMd5Hash,
				Sensitive: true,
			},
			"is_file_server_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"instance_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_integration.IntegrationInstanceLifecycleStateActive),
					string(oci_integration.IntegrationInstanceLifecycleStateInactive),
				}, true),
			},
			"state_message": {
				Type:     schema.TypeString,
				Computed: true,
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

func createIntegrationIntegrationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationIntegrationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).integrationInstanceClient()

	var powerOff = false
	if configState, ok := sync.D.GetOkExists("state"); ok {
		wantedState := oci_integration.IntegrationInstanceLifecycleStateEnum(strings.ToUpper(configState.(string)))
		if wantedState == oci_integration.IntegrationInstanceLifecycleStateInactive {
			powerOff = true
		}
	}

	if error := CreateResource(d, sync); error != nil {
		return error
	}

	if powerOff {
		return powerOffIntegrationInstance(d, sync)
	}

	return nil
}

func powerOffIntegrationInstance(d *schema.ResourceData, sync *IntegrationIntegrationInstanceResourceCrud) error {
	if err := sync.StopIntegerationInstance(); err != nil {
		return err
	}
	return ReadResource(sync)
}

func readIntegrationIntegrationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationIntegrationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).integrationInstanceClient()

	return ReadResource(sync)
}

func updateIntegrationIntegrationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationIntegrationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).integrationInstanceClient()

	// Start/Stop Integration instance
	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_integration.IntegrationInstanceLifecycleStateActive == oci_integration.IntegrationInstanceLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_integration.IntegrationInstanceLifecycleStateInactive == oci_integration.IntegrationInstanceLifecycleStateEnum(wantedState) {
			powerOff = true
		} else {
			return fmt.Errorf("[ERROR] Invalid state input for update %v", wantedState)
		}
	}

	if powerOn {
		if err := sync.StartIntegerationInstance(); err != nil {
			return err
		}
		if err := sync.D.Set("state", oci_integration.IntegrationInstanceLifecycleStateActive); err != nil {
			return err
		}
	}

	if err := UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopIntegerationInstance(); err != nil {
			return err
		}
		if err := sync.D.Set("state", oci_integration.IntegrationInstanceLifecycleStateInactive); err != nil {
			return err
		}
	}

	return nil
}

func deleteIntegrationIntegrationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationIntegrationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).integrationInstanceClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type IntegrationIntegrationInstanceResourceCrud struct {
	BaseCrud
	Client                 *oci_integration.IntegrationInstanceClient
	Res                    *oci_integration.IntegrationInstance
	DisableNotFoundRetries bool
}

func (s *IntegrationIntegrationInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IntegrationIntegrationInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateCreating),
	}
}

func (s *IntegrationIntegrationInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateActive),
	}
}

func (s *IntegrationIntegrationInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateDeleting),
	}
}

func (s *IntegrationIntegrationInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateDeleted),
	}
}

func (s *IntegrationIntegrationInstanceResourceCrud) Create() error {
	request := oci_integration.CreateIntegrationInstanceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if consumptionModel, ok := s.D.GetOkExists("consumption_model"); ok {
		request.ConsumptionModel = oci_integration.CreateIntegrationInstanceDetailsConsumptionModelEnum(consumptionModel.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if idcsAt, ok := s.D.GetOkExists("idcs_at"); ok {
		tmp := idcsAt.(string)
		request.IdcsAt = &tmp
	}

	if integrationInstanceType, ok := s.D.GetOkExists("integration_instance_type"); ok {
		request.IntegrationInstanceType = oci_integration.CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum(integrationInstanceType.(string))
	}

	if isByol, ok := s.D.GetOkExists("is_byol"); ok {
		tmp := isByol.(bool)
		request.IsByol = &tmp
	}

	if isFileServerEnabled, ok := s.D.GetOkExists("is_file_server_enabled"); ok {
		tmp := isFileServerEnabled.(bool)
		request.IsFileServerEnabled = &tmp
	}

	if messagePacks, ok := s.D.GetOkExists("message_packs"); ok {
		tmp := messagePacks.(int)
		request.MessagePacks = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "integration")

	response, err := s.Client.CreateIntegrationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIntegrationInstanceFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "integration"), oci_integration.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IntegrationIntegrationInstanceResourceCrud) getIntegrationInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_integration.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	integrationInstanceId, err := integrationInstanceWaitForWorkRequest(workId, "integration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*integrationInstanceId)

	return s.Get()
}

func integrationInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "integration", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_integration.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func integrationInstanceWaitForWorkRequest(wId *string, entityType string, action oci_integration.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_integration.IntegrationInstanceClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "integration")
	retryPolicy.ShouldRetryOperation = integrationInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_integration.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_integration.WorkRequestStatusInProgress),
			string(oci_integration.WorkRequestStatusAccepted),
			string(oci_integration.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_integration.WorkRequestStatusSucceeded),
			string(oci_integration.WorkRequestStatusFailed),
			string(oci_integration.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_integration.GetWorkRequestRequest{
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
	// The workrequest may have failed, check for errors if identifier is not found
	if identifier == nil {
		return nil, getErrorFromIntegrationInstanceWorkRequest(client, wId, response.CompartmentId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromIntegrationInstanceWorkRequest(client *oci_integration.IntegrationInstanceClient, workRequestId *string, compartmentId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_integration.WorkRequestResourceActionTypeEnum) error {

	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_integration.ListWorkRequestErrorsRequest{
			CompartmentId: compartmentId,
			WorkRequestId: workRequestId,
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

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workRequestId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *IntegrationIntegrationInstanceResourceCrud) Get() error {
	request := oci_integration.GetIntegrationInstanceRequest{}

	tmp := s.D.Id()
	request.IntegrationInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "integration")

	response, err := s.Client.GetIntegrationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IntegrationInstance
	return nil
}

func (s *IntegrationIntegrationInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_integration.UpdateIntegrationInstanceRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.IntegrationInstanceId = &tmp

	if integrationInstanceType, ok := s.D.GetOkExists("integration_instance_type"); ok {
		request.IntegrationInstanceType = oci_integration.UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum(integrationInstanceType.(string))
	}

	if isByol, ok := s.D.GetOkExists("is_byol"); ok {
		tmp := isByol.(bool)
		request.IsByol = &tmp
	}

	if isFileServerEnabled, ok := s.D.GetOkExists("is_file_server_enabled"); ok {
		tmp := isFileServerEnabled.(bool)
		request.IsFileServerEnabled = &tmp
	}

	if messagePacks, ok := s.D.GetOkExists("message_packs"); ok {
		tmp := messagePacks.(int)
		request.MessagePacks = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "integration")

	response, err := s.Client.UpdateIntegrationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIntegrationInstanceFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "integration"), oci_integration.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *IntegrationIntegrationInstanceResourceCrud) Delete() error {
	request := oci_integration.DeleteIntegrationInstanceRequest{}

	tmp := s.D.Id()
	request.IntegrationInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "integration")

	response, err := s.Client.DeleteIntegrationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := integrationInstanceWaitForWorkRequest(workId, "integration",
		oci_integration.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *IntegrationIntegrationInstanceResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("consumption_model", s.Res.ConsumptionModel)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceUrl != nil {
		s.D.Set("instance_url", *s.Res.InstanceUrl)
	}

	s.D.Set("integration_instance_type", s.Res.IntegrationInstanceType)

	if s.Res.IsByol != nil {
		s.D.Set("is_byol", *s.Res.IsByol)
	}

	if s.Res.IsFileServerEnabled != nil {
		s.D.Set("is_file_server_enabled", *s.Res.IsFileServerEnabled)
	}

	if s.Res.MessagePacks != nil {
		s.D.Set("message_packs", *s.Res.MessagePacks)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *IntegrationIntegrationInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_integration.ChangeIntegrationInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.IntegrationInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "integration")

	_, err := s.Client.ChangeIntegrationInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *IntegrationIntegrationInstanceResourceCrud) StartIntegerationInstance() error {
	state := oci_integration.IntegrationInstanceLifecycleStateActive
	if err := s.Get(); err != nil {
		return err
	}
	if s.Res.LifecycleState == state {
		fmt.Printf("[WARN] The Integration instance already in the wanted state: %v", state)
		return nil
	}
	request := oci_integration.StartIntegrationInstanceRequest{}

	tmp := s.D.Id()
	request.IntegrationInstanceId = &tmp
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "integration")

	if _, err := s.Client.StartIntegrationInstance(context.Background(), request); err != nil {
		return err
	}
	resourceChangedFunc := func() bool { return s.Res.LifecycleState == state }

	return WaitForResourceCondition(s, resourceChangedFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *IntegrationIntegrationInstanceResourceCrud) StopIntegerationInstance() error {
	state := oci_integration.IntegrationInstanceLifecycleStateInactive
	if err := s.Get(); err != nil {
		return err
	}
	if s.Res.LifecycleState == state {
		fmt.Printf("[WARN] The Integration instance already in the wanted state: %v", state)
		return nil
	}
	request := oci_integration.StopIntegrationInstanceRequest{}

	tmp := s.D.Id()
	request.IntegrationInstanceId = &tmp
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "integration")

	if _, err := s.Client.StopIntegrationInstance(context.Background(), request); err != nil {
		return err
	}
	resourceChangedFunc := func() bool { return s.Res.LifecycleState == state }

	return WaitForResourceCondition(s, resourceChangedFunc, s.D.Timeout(schema.TimeoutUpdate))
}
