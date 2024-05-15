// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func EmailEmailReturnPathResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createEmailEmailReturnPath,
		Read:     readEmailEmailReturnPath,
		Update:   updateEmailEmailReturnPath,
		Delete:   deleteEmailEmailReturnPath,
		Schema: map[string]*schema.Schema{
			// Required
			"parent_resource_id": {
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
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"cname_record_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_subdomain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
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

func createEmailEmailReturnPath(d *schema.ResourceData, m interface{}) error {
	sync := &EmailEmailReturnPathResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.CreateResource(d, sync)
}

func readEmailEmailReturnPath(d *schema.ResourceData, m interface{}) error {
	sync := &EmailEmailReturnPathResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

func updateEmailEmailReturnPath(d *schema.ResourceData, m interface{}) error {
	sync := &EmailEmailReturnPathResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteEmailEmailReturnPath(d *schema.ResourceData, m interface{}) error {
	sync := &EmailEmailReturnPathResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type EmailEmailReturnPathResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_email.EmailClient
	Res                    *oci_email.EmailReturnPath
	DisableNotFoundRetries bool
}

func (s *EmailEmailReturnPathResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *EmailEmailReturnPathResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_email.EmailReturnPathLifecycleStateCreating),
	}
}

func (s *EmailEmailReturnPathResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_email.EmailReturnPathLifecycleStateActive),
		string(oci_email.EmailReturnPathLifecycleStateNeedsAttention),
	}
}

func (s *EmailEmailReturnPathResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_email.EmailReturnPathLifecycleStateDeleting),
	}
}

func (s *EmailEmailReturnPathResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_email.EmailReturnPathLifecycleStateDeleted),
	}
}

func (s *EmailEmailReturnPathResourceCrud) Create() error {
	request := oci_email.CreateEmailReturnPathRequest{}

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

	if parentResourceId, ok := s.D.GetOkExists("parent_resource_id"); ok {
		tmp := parentResourceId.(string)
		request.ParentResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.CreateEmailReturnPath(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getEmailReturnPathFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email"), oci_email.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *EmailEmailReturnPathResourceCrud) getEmailReturnPathFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_email.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	emailReturnPathId, err := emailReturnPathWaitForWorkRequest(workId, "emailreturnpath",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*emailReturnPathId)

	return s.Get()
}

func emailReturnPathWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func emailReturnPathWaitForWorkRequest(wId *string, entityType string, action oci_email.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_email.EmailClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "email")
	retryPolicy.ShouldRetryOperation = emailReturnPathWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromEmailEmailReturnPathWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromEmailEmailReturnPathWorkRequest(client *oci_email.EmailClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_email.ActionTypeEnum) error {
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

func (s *EmailEmailReturnPathResourceCrud) Get() error {
	request := oci_email.GetEmailReturnPathRequest{}

	tmp := s.D.Id()
	request.EmailReturnPathId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.GetEmailReturnPath(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EmailReturnPath
	return nil
}

func (s *EmailEmailReturnPathResourceCrud) Update() error {
	request := oci_email.UpdateEmailReturnPathRequest{}

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
	request.EmailReturnPathId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.UpdateEmailReturnPath(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEmailReturnPathFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email"), oci_email.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *EmailEmailReturnPathResourceCrud) Delete() error {
	request := oci_email.DeleteEmailReturnPathRequest{}

	tmp := s.D.Id()
	request.EmailReturnPathId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.DeleteEmailReturnPath(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := emailReturnPathWaitForWorkRequest(workId, "emailreturnpath",
		oci_email.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *EmailEmailReturnPathResourceCrud) SetData() error {
	if s.Res.CnameRecordValue != nil {
		s.D.Set("cname_record_value", *s.Res.CnameRecordValue)
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

	if s.Res.DnsSubdomainName != nil {
		s.D.Set("dns_subdomain_name", *s.Res.DnsSubdomainName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
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

	return nil
}

func EmailReturnPathSummaryToMap(obj oci_email.EmailReturnPathSummary) map[string]interface{} {
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

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentResourceId != nil {
		result["parent_resource_id"] = string(*obj.ParentResourceId)
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
