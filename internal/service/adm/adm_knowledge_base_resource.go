// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AdmKnowledgeBaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAdmKnowledgeBase,
		Read:     readAdmKnowledgeBase,
		Update:   updateAdmKnowledgeBase,
		Delete:   deleteAdmKnowledgeBase,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
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

			// Computed
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

func createAdmKnowledgeBase(d *schema.ResourceData, m interface{}) error {
	sync := &AdmKnowledgeBaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readAdmKnowledgeBase(d *schema.ResourceData, m interface{}) error {
	sync := &AdmKnowledgeBaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.ReadResource(sync)
}

func updateAdmKnowledgeBase(d *schema.ResourceData, m interface{}) error {
	sync := &AdmKnowledgeBaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAdmKnowledgeBase(d *schema.ResourceData, m interface{}) error {
	sync := &AdmKnowledgeBaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AdmKnowledgeBaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_adm.ApplicationDependencyManagementClient
	Res                    *oci_adm.KnowledgeBase
	DisableNotFoundRetries bool
}

func (s *AdmKnowledgeBaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AdmKnowledgeBaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_adm.KnowledgeBaseLifecycleStateCreating),
	}
}

func (s *AdmKnowledgeBaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_adm.KnowledgeBaseLifecycleStateActive),
	}
}

func (s *AdmKnowledgeBaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_adm.KnowledgeBaseLifecycleStateDeleting),
	}
}

func (s *AdmKnowledgeBaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_adm.KnowledgeBaseLifecycleStateDeleted),
	}
}

func (s *AdmKnowledgeBaseResourceCrud) Create() error {
	request := oci_adm.CreateKnowledgeBaseRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.CreateKnowledgeBase(context.Background(), request)
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "knowledgebase") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getKnowledgeBaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm"), oci_adm.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AdmKnowledgeBaseResourceCrud) getKnowledgeBaseFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_adm.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	knowledgeBaseId, err := knowledgeBaseWaitForWorkRequest(workId, "knowledgebase",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, knowledgeBaseId)
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
	s.D.SetId(*knowledgeBaseId)

	return s.Get()
}

func knowledgeBaseWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func knowledgeBaseWaitForWorkRequest(wId *string, entityType string, action oci_adm.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_adm.ApplicationDependencyManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "adm")
	retryPolicy.ShouldRetryOperation = knowledgeBaseWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromAdmKnowledgeBaseWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAdmKnowledgeBaseWorkRequest(client *oci_adm.ApplicationDependencyManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_adm.ActionTypeEnum) error {
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

func (s *AdmKnowledgeBaseResourceCrud) Get() error {
	request := oci_adm.GetKnowledgeBaseRequest{}

	tmp := s.D.Id()
	request.KnowledgeBaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.GetKnowledgeBase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KnowledgeBase
	return nil
}

func (s *AdmKnowledgeBaseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_adm.UpdateKnowledgeBaseRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.KnowledgeBaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.UpdateKnowledgeBase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getKnowledgeBaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm"), oci_adm.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AdmKnowledgeBaseResourceCrud) Delete() error {
	request := oci_adm.DeleteKnowledgeBaseRequest{}

	tmp := s.D.Id()
	request.KnowledgeBaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.DeleteKnowledgeBase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := knowledgeBaseWaitForWorkRequest(workId, "knowledgebase",
		oci_adm.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AdmKnowledgeBaseResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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

func KnowledgeBaseSummaryToMap(obj oci_adm.KnowledgeBaseSummary) map[string]interface{} {
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

func (s *AdmKnowledgeBaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_adm.ChangeKnowledgeBaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.KnowledgeBaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm")

	response, err := s.Client.ChangeKnowledgeBaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getKnowledgeBaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "adm"), oci_adm.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
