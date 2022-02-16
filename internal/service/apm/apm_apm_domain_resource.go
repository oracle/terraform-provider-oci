// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_apm "github.com/oracle/oci-go-sdk/v58/apmcontrolplane"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func ApmApmDomainResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApmApmDomain,
		Read:     readApmApmDomain,
		Update:   updateApmApmDomain,
		Delete:   deleteApmApmDomain,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_free_tier": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"data_upload_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
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

func createApmApmDomain(d *schema.ResourceData, m interface{}) error {
	sync := &ApmApmDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmDomainClient()

	return tfresource.CreateResource(d, sync)
}

func readApmApmDomain(d *schema.ResourceData, m interface{}) error {
	sync := &ApmApmDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmDomainClient()

	return tfresource.ReadResource(sync)
}

func updateApmApmDomain(d *schema.ResourceData, m interface{}) error {
	sync := &ApmApmDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmDomainClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApmApmDomain(d *schema.ResourceData, m interface{}) error {
	sync := &ApmApmDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmDomainClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ApmApmDomainResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apm.ApmDomainClient
	Res                    *oci_apm.ApmDomain
	DisableNotFoundRetries bool
}

func (s *ApmApmDomainResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ApmApmDomainResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_apm.LifecycleStatesCreating),
	}
}

func (s *ApmApmDomainResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_apm.LifecycleStatesActive),
	}
}

func (s *ApmApmDomainResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_apm.LifecycleStatesDeleting),
	}
}

func (s *ApmApmDomainResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_apm.LifecycleStatesDeleted),
	}
}

func (s *ApmApmDomainResourceCrud) Create() error {
	request := oci_apm.CreateApmDomainRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
		tmp := isFreeTier.(bool)
		request.IsFreeTier = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm")

	response, err := s.Client.CreateApmDomain(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getApmDomainFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm"), oci_apm.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ApmApmDomainResourceCrud) getApmDomainFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_apm.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	apmDomainId, err := apmDomainWaitForWorkRequest(workId, "apmDomain",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*apmDomainId)

	return s.Get()
}

func apmDomainWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "apm", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_apm.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func apmDomainWaitForWorkRequest(wId *string, entityType string, action oci_apm.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_apm.ApmDomainClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "apm")
	retryPolicy.ShouldRetryOperation = apmDomainWorkRequestShouldRetryFunc(timeout)

	response := oci_apm.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_apm.OperationStatusInProgress),
			string(oci_apm.OperationStatusAccepted),
			string(oci_apm.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_apm.OperationStatusSucceeded),
			string(oci_apm.OperationStatusFailed),
			string(oci_apm.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_apm.GetWorkRequestRequest{
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
	for _, res := range response.WorkRequest.Resources {
		if *res.EntityType == entityType {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_apm.OperationStatusFailed || response.Status == oci_apm.OperationStatusCanceled {
		return nil, getErrorFromApmControlPlaneWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromApmControlPlaneWorkRequest(client *oci_apm.ApmDomainClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_apm.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_apm.ListWorkRequestErrorsRequest{
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

func (s *ApmApmDomainResourceCrud) Get() error {
	request := oci_apm.GetApmDomainRequest{}

	tmp := s.D.Id()
	request.ApmDomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm")

	response, err := s.Client.GetApmDomain(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApmDomain
	return nil
}

func (s *ApmApmDomainResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_apm.UpdateApmDomainRequest{}

	tmp := s.D.Id()
	request.ApmDomainId = &tmp

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm")

	response, err := s.Client.UpdateApmDomain(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getApmDomainFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm"), oci_apm.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ApmApmDomainResourceCrud) Delete() error {
	request := oci_apm.DeleteApmDomainRequest{}

	tmp := s.D.Id()
	request.ApmDomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm")

	response, err := s.Client.DeleteApmDomain(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := apmDomainWaitForWorkRequest(workId, "apmDomain",
		oci_apm.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ApmApmDomainResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataUploadEndpoint != nil {
		s.D.Set("data_upload_endpoint", *s.Res.DataUploadEndpoint)
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

	if s.Res.IsFreeTier != nil {
		s.D.Set("is_free_tier", *s.Res.IsFreeTier)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *ApmApmDomainResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_apm.ChangeApmDomainCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ApmDomainId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm")

	response, err := s.Client.ChangeApmDomainCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getApmDomainFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm"), oci_apm.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
