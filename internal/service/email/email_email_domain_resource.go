// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"
)

func EmailEmailDomainResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createEmailEmailDomain,
		Read:     readEmailEmailDomain,
		Update:   updateEmailEmailDomain,
		Delete:   deleteEmailEmailDomain,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"active_dkim_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_spf": {
				Type:     schema.TypeBool,
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
		},
	}
}

func createEmailEmailDomain(d *schema.ResourceData, m interface{}) error {
	sync := &EmailEmailDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.CreateResource(d, sync)
}

func readEmailEmailDomain(d *schema.ResourceData, m interface{}) error {
	sync := &EmailEmailDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

func updateEmailEmailDomain(d *schema.ResourceData, m interface{}) error {
	sync := &EmailEmailDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteEmailEmailDomain(d *schema.ResourceData, m interface{}) error {
	sync := &EmailEmailDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type EmailEmailDomainResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_email.EmailClient
	Res                    *oci_email.EmailDomain
	DisableNotFoundRetries bool
}

func (s *EmailEmailDomainResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *EmailEmailDomainResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_email.EmailDomainLifecycleStateCreating),
	}
}

func (s *EmailEmailDomainResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_email.EmailDomainLifecycleStateActive),
	}
}

func (s *EmailEmailDomainResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_email.EmailDomainLifecycleStateDeleting),
	}
}

func (s *EmailEmailDomainResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_email.EmailDomainLifecycleStateDeleted),
	}
}

func (s *EmailEmailDomainResourceCrud) Create() error {
	request := oci_email.CreateEmailDomainRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.CreateEmailDomain(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getEmailDomainFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email"), oci_email.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *EmailEmailDomainResourceCrud) getEmailDomainFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_email.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	emailDomainId, err := emailDomainWaitForWorkRequest(workId, "email",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest")
		return err
	}
	s.D.SetId(*emailDomainId)

	return s.Get()
}

func emailDomainWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "email", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_email.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func emailDomainWaitForWorkRequest(wId *string, entityType string, action oci_email.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_email.EmailClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "email")
	retryPolicy.ShouldRetryOperation = emailDomainWorkRequestShouldRetryFunc(timeout)

	response := oci_email.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_email.OperationStatusInProgress),
			string(oci_email.OperationStatusAccepted),
			string(oci_email.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_email.OperationStatusSucceeded),
			string(oci_email.OperationStatusFailed),
			string(oci_email.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_email.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_email.OperationStatusFailed || response.Status == oci_email.OperationStatusCanceled {
		return nil, getErrorFromEmailEmailDomainWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromEmailEmailDomainWorkRequest(client *oci_email.EmailClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_email.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_email.ListWorkRequestErrorsRequest{
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

func (s *EmailEmailDomainResourceCrud) Get() error {
	request := oci_email.GetEmailDomainRequest{}

	tmp := s.D.Id()
	request.EmailDomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.GetEmailDomain(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EmailDomain
	return nil
}

func (s *EmailEmailDomainResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_email.UpdateEmailDomainRequest{}

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

	tmp := s.D.Id()
	request.EmailDomainId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.UpdateEmailDomain(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEmailDomainFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email"), oci_email.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *EmailEmailDomainResourceCrud) Delete() error {
	request := oci_email.DeleteEmailDomainRequest{}

	tmp := s.D.Id()
	request.EmailDomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.DeleteEmailDomain(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := emailDomainWaitForWorkRequest(workId, "email",
		oci_email.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *EmailEmailDomainResourceCrud) SetData() error {
	if s.Res.ActiveDkimId != nil {
		s.D.Set("active_dkim_id", *s.Res.ActiveDkimId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSpf != nil {
		s.D.Set("is_spf", *s.Res.IsSpf)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func EmailDomainSummaryToMap(obj oci_email.EmailDomainSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActiveDkimId != nil {
		result["active_dkim_id"] = string(*obj.ActiveDkimId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *EmailEmailDomainResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_email.ChangeEmailDomainCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.EmailDomainId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.ChangeEmailDomainCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEmailDomainFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email"), oci_email.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
