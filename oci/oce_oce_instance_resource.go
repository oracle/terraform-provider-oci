// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_oce "github.com/oracle/oci-go-sdk/oce"
)

func OceOceInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &TwentyMinutes,
			Update: &TwentyMinutes,
			Delete: &TwentyMinutes,
		},
		Create: createOceOceInstance,
		Read:   readOceOceInstance,
		Update: updateOceOceInstance,
		Delete: deleteOceOceInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"admin_email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idcs_access_token": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				StateFunc: getMd5Hash,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"object_storage_namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenancy_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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

			// Computed
			"guid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_tenancy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
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

func createOceOceInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).oceInstanceClient

	return CreateResource(d, sync)
}

func readOceOceInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).oceInstanceClient

	return ReadResource(sync)
}

func updateOceOceInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).oceInstanceClient

	return UpdateResource(d, sync)
}

func deleteOceOceInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).oceInstanceClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type OceOceInstanceResourceCrud struct {
	BaseCrud
	Client                 *oci_oce.OceInstanceClient
	Res                    *oci_oce.OceInstance
	DisableNotFoundRetries bool
}

func (s *OceOceInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OceOceInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_oce.OceInstanceLifecycleStateCreating),
	}
}

func (s *OceOceInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_oce.OceInstanceLifecycleStateActive),
	}
}

func (s *OceOceInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_oce.OceInstanceLifecycleStateDeleting),
	}
}

func (s *OceOceInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_oce.OceInstanceLifecycleStateDeleted),
	}
}

func (s *OceOceInstanceResourceCrud) Create() error {
	request := oci_oce.CreateOceInstanceRequest{}

	if adminEmail, ok := s.D.GetOkExists("admin_email"); ok {
		tmp := adminEmail.(string)
		request.AdminEmail = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if idcsAccessToken, ok := s.D.GetOkExists("idcs_access_token"); ok {
		tmp := idcsAccessToken.(string)
		request.IdcsAccessToken = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if objectStorageNamespace, ok := s.D.GetOkExists("object_storage_namespace"); ok {
		tmp := objectStorageNamespace.(string)
		request.ObjectStorageNamespace = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	if tenancyName, ok := s.D.GetOkExists("tenancy_name"); ok {
		tmp := tenancyName.(string)
		request.TenancyName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "oce")

	response, err := s.Client.CreateOceInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOceInstanceFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "oce"), oci_oce.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OceOceInstanceResourceCrud) getOceInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_oce.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	oceInstanceId, err := oceInstanceWaitForWorkRequest(workId, "oce",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		log.Printf("[DEBUG] operation failed: %v for identifier: %v\n", workId, oceInstanceId)
		return err
	}

	if oceInstanceId == nil {
		return fmt.Errorf("operation failed: %v for identifier: %v\n", workId, oceInstanceId)
	}

	s.D.SetId(*oceInstanceId)

	return s.Get()
}

func oceInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "oce", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_oce.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func oceInstanceWaitForWorkRequest(wId *string, entityType string, action oci_oce.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_oce.OceInstanceClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "oce")
	retryPolicy.ShouldRetryOperation = oceInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_oce.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_oce.WorkRequestStatusInProgress),
			string(oci_oce.WorkRequestStatusAccepted),
			string(oci_oce.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_oce.WorkRequestStatusSucceeded),
			string(oci_oce.WorkRequestStatusFailed),
			string(oci_oce.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_oce.GetWorkRequestRequest{
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

	// The OCE workrequest may have failed, check for errors if identifier is not found
	if identifier == nil {
		return nil, getErrorFromOceInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOceInstanceWorkRequest(client *oci_oce.OceInstanceClient, wId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_oce.WorkRequestResourceActionTypeEnum) error {

	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_oce.ListWorkRequestErrorsRequest{
			WorkRequestId: wId,
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

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *OceOceInstanceResourceCrud) Get() error {
	request := oci_oce.GetOceInstanceRequest{}

	tmp := s.D.Id()
	request.OceInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "oce")

	response, err := s.Client.GetOceInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OceInstance
	return nil
}

func (s *OceOceInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	request := oci_oce.UpdateOceInstanceRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.OceInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "oce")

	response, err := s.Client.UpdateOceInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOceInstanceFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "oce"), oci_oce.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OceOceInstanceResourceCrud) Delete() error {
	request := oci_oce.DeleteOceInstanceRequest{}

	tmp := s.D.Id()
	request.OceInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "oce")

	response, err := s.Client.DeleteOceInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := oceInstanceWaitForWorkRequest(workId, "oce",
		oci_oce.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OceOceInstanceResourceCrud) SetData() error {
	if s.Res.AdminEmail != nil {
		s.D.Set("admin_email", *s.Res.AdminEmail)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Guid != nil {
		s.D.Set("guid", *s.Res.Guid)
	}

	if s.Res.IdcsTenancy != nil {
		s.D.Set("idcs_tenancy", *s.Res.IdcsTenancy)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectStorageNamespace != nil {
		s.D.Set("object_storage_namespace", *s.Res.ObjectStorageNamespace)
	}

	s.D.Set("service", genericMapToJsonMap(s.Res.Service))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TenancyName != nil {
		s.D.Set("tenancy_name", *s.Res.TenancyName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *OceOceInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_oce.ChangeOceInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OceInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "oce")

	_, err := s.Client.ChangeOceInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
