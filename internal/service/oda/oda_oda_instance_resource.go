// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oda

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OdaOdaInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Update: schema.DefaultTimeout(40 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},
		Create: createOdaOdaInstance,
		Read:   readOdaOdaInstance,
		Update: updateOdaOdaInstance,
		Delete: deleteOdaOdaInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shape_name": {
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
			"identity_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_role_based_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"attachment_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attachment_types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connector_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity_app_console_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity_app_guid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"imported_package_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"imported_package_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"lifecycle_sub_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"restricted_operations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"operation_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"restricting_service": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_oda.OdaInstanceLifecycleStateActive),
					string(oci_oda.OdaInstanceLifecycleStateInactive),
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
			"web_app_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOdaOdaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OdaClient()

	var isInactiveRequest = false
	if configState, ok := sync.D.GetOkExists("state"); ok {
		wantedState := oci_oda.OdaInstanceLifecycleStateEnum(strings.ToUpper(configState.(string)))
		if wantedState == oci_oda.OdaInstanceLifecycleStateInactive {
			isInactiveRequest = true
		}
	}

	if error := tfresource.CreateResource(d, sync); error != nil {
		return error
	}

	if isInactiveRequest {
		return inactiveOdaIfNeeded(d, sync)
	}

	return nil
}

func inactiveOdaIfNeeded(d *schema.ResourceData, sync *OdaOdaInstanceResourceCrud) error {
	if err := sync.StopOdaInstance(); err != nil {
		return err
	}
	return tfresource.ReadResource(sync)
}

func readOdaOdaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OdaClient()

	return tfresource.ReadResource(sync)
}

func updateOdaOdaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OdaClient()

	// Start/Stop ODA instance
	stateActive, stateInactive := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_oda.OdaInstanceLifecycleStateActive == oci_oda.OdaInstanceLifecycleStateEnum(wantedState) {
			stateActive = true
			stateInactive = false
		} else if oci_oda.OdaInstanceLifecycleStateInactive == oci_oda.OdaInstanceLifecycleStateEnum(wantedState) {
			stateInactive = true
			stateActive = false
		} else {
			return fmt.Errorf("[ERROR] Invalid state input for Update %v", wantedState)
		}
	}

	if stateActive {
		if err := sync.StartOdaInstance(); err != nil {
			return err
		}
		if err := sync.D.Set("state", oci_oda.OdaInstanceLifecycleStateActive); err != nil {
			return err
		}
	}

	// when state is inactive, it is invalid to Update resource
	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if stateInactive {
		if err := sync.StopOdaInstance(); err != nil {
			return err
		}
		if err := sync.D.Set("state", oci_oda.OdaInstanceLifecycleStateInactive); err != nil {
			return err
		}
	}

	return nil
}

func deleteOdaOdaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OdaClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OdaOdaInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_oda.OdaClient
	Res                    *oci_oda.OdaInstance
	DisableNotFoundRetries bool
}

func (s *OdaOdaInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OdaOdaInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateCreating),
	}
}

func (s *OdaOdaInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateActive),
	}
}

func (s *OdaOdaInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateDeleting),
	}
}

func (s *OdaOdaInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateDeleted),
	}
}

func (s *OdaOdaInstanceResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateUpdating),
	}
}

func (s *OdaOdaInstanceResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateActive),
	}
}

func (s *OdaOdaInstanceResourceCrud) Create() error {
	request := oci_oda.CreateOdaInstanceRequest{}

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

	if identityDomain, ok := s.D.GetOkExists("identity_domain"); ok {
		tmp := identityDomain.(string)
		request.IdentityDomain = &tmp
	}

	if isRoleBasedAccess, ok := s.D.GetOkExists("is_role_based_access"); ok {
		tmp := isRoleBasedAccess.(bool)
		request.IsRoleBasedAccess = &tmp
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		request.ShapeName = oci_oda.CreateOdaInstanceDetailsShapeNameEnum(shapeName.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.CreateOdaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getOdaInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda"), oci_oda.WorkRequestResourceResourceActionCreate, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OdaOdaInstanceResourceCrud) getOdaInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_oda.WorkRequestResourceResourceActionEnum, timeout time.Duration) error {

	// Wait until it finishes
	odaInstanceId, err := odaInstanceWaitForWorkRequest(workId, "oda",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*odaInstanceId)

	return s.Get()
}

func odaInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "oda", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_oda.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func odaInstanceWaitForWorkRequest(wId *string, entityType string, action oci_oda.WorkRequestResourceResourceActionEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_oda.OdaClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "oda")
	retryPolicy.ShouldRetryOperation = odaInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_oda.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_oda.WorkRequestStatusInProgress),
			string(oci_oda.WorkRequestStatusAccepted),
			string(oci_oda.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_oda.WorkRequestStatusSucceeded),
			string(oci_oda.WorkRequestStatusFailed),
			string(oci_oda.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_oda.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.ResourceType), entityType) {
			if res.ResourceAction == action {
				identifier = res.ResourceId
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_oda.WorkRequestStatusFailed || response.Status == oci_oda.WorkRequestStatusCanceled {
		return nil, getErrorFromOdaOdaInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOdaOdaInstanceWorkRequest(client *oci_oda.OdaClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_oda.WorkRequestResourceResourceActionEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_oda.ListWorkRequestErrorsRequest{
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

func (s *OdaOdaInstanceResourceCrud) Get() error {
	request := oci_oda.GetOdaInstanceRequest{}

	tmp := s.D.Id()
	request.OdaInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.GetOdaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OdaInstance
	return nil
}

func (s *OdaOdaInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_oda.UpdateOdaInstanceRequest{}

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
	request.OdaInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.UpdateOdaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OdaInstance
	return nil
}

func (s *OdaOdaInstanceResourceCrud) Delete() error {
	request := oci_oda.DeleteOdaInstanceRequest{}

	tmp := s.D.Id()
	request.OdaInstanceId = &tmp

	if retentionTime, ok := s.D.GetOkExists("retention_time"); ok {
		tmp := retentionTime.(int)
		request.RetentionTime = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.DeleteOdaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := odaInstanceWaitForWorkRequest(workId, "oda",
		oci_oda.WorkRequestResourceResourceActionDelete, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OdaOdaInstanceResourceCrud) SetData() error {
	s.D.Set("attachment_ids", s.Res.AttachmentIds)

	s.D.Set("attachment_types", s.Res.AttachmentTypes)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectorUrl != nil {
		s.D.Set("connector_url", *s.Res.ConnectorUrl)
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdentityAppConsoleUrl != nil {
		s.D.Set("identity_app_console_url", *s.Res.IdentityAppConsoleUrl)
	}

	if s.Res.IdentityAppGuid != nil {
		s.D.Set("identity_app_guid", *s.Res.IdentityAppGuid)
	}

	if s.Res.IdentityDomain != nil {
		s.D.Set("identity_domain", *s.Res.IdentityDomain)
	}

	s.D.Set("imported_package_ids", s.Res.ImportedPackageIds)

	s.D.Set("imported_package_names", s.Res.ImportedPackageNames)

	if s.Res.IsRoleBasedAccess != nil {
		s.D.Set("is_role_based_access", *s.Res.IsRoleBasedAccess)
	}

	s.D.Set("lifecycle_sub_state", s.Res.LifecycleSubState)

	restrictedOperations := []interface{}{}
	for _, item := range s.Res.RestrictedOperations {
		restrictedOperations = append(restrictedOperations, RestrictedOperationToMap(item))
	}
	s.D.Set("restricted_operations", restrictedOperations)

	s.D.Set("shape_name", s.Res.ShapeName)

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

	if s.Res.WebAppUrl != nil {
		s.D.Set("web_app_url", *s.Res.WebAppUrl)
	}

	return nil
}

func RestrictedOperationToMap(obj oci_oda.RestrictedOperation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.OperationName != nil {
		result["operation_name"] = string(*obj.OperationName)
	}

	if obj.RestrictingService != nil {
		result["restricting_service"] = string(*obj.RestrictingService)
	}

	return result
}

func (s *OdaOdaInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_oda.ChangeOdaInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OdaInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.ChangeOdaInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOdaInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda"), oci_oda.WorkRequestResourceResourceActionChangeCompartment, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OdaOdaInstanceResourceCrud) StartOdaInstance() error {
	state := oci_oda.OdaInstanceLifecycleStateActive
	if err := s.Get(); err != nil {
		return err
	}
	if s.Res.LifecycleState == state {
		fmt.Printf("[WARN] The ODA instance already in the wanted state: %v", state)
		return nil
	}
	request := oci_oda.StartOdaInstanceRequest{}

	tmp := s.D.Id()
	request.OdaInstanceId = &tmp

	if _, err := s.Client.StartOdaInstance(context.Background(), request); err != nil {
		return err
	}
	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == state }

	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OdaOdaInstanceResourceCrud) StopOdaInstance() error {
	state := oci_oda.OdaInstanceLifecycleStateInactive
	if err := s.Get(); err != nil {
		return err
	}
	if s.Res.LifecycleState == state {
		fmt.Printf("[WARN] The ODA instance already in the wanted state: %v", state)
		return nil
	}
	request := oci_oda.StopOdaInstanceRequest{}

	tmp := s.D.Id()
	request.OdaInstanceId = &tmp

	if _, err := s.Client.StopOdaInstance(context.Background(), request); err != nil {
		return err
	}
	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == state }

	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}
