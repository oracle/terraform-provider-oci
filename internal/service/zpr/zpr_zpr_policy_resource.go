// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package zpr

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_zpr "github.com/oracle/oci-go-sdk/v65/zpr"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ZprZprPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createZprZprPolicy,
		Read:     readZprZprPolicy,
		Update:   updateZprZprPolicy,
		Delete:   deleteZprZprPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"statements": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
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

func createZprZprPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ZprZprPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ZprClient()

	return tfresource.CreateResource(d, sync)
}

func readZprZprPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ZprZprPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ZprClient()

	return tfresource.ReadResource(sync)
}

func updateZprZprPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ZprZprPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ZprClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteZprZprPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ZprZprPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ZprClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ZprZprPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_zpr.ZprClient
	Res                    *oci_zpr.ZprPolicy
	DisableNotFoundRetries bool
}

func (s *ZprZprPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ZprZprPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_zpr.ZprPolicyLifecycleStateCreating),
	}
}

func (s *ZprZprPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_zpr.ZprPolicyLifecycleStateActive),
	}
}

func (s *ZprZprPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_zpr.ZprPolicyLifecycleStateDeleting),
	}
}

func (s *ZprZprPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_zpr.ZprPolicyLifecycleStateDeleted),
	}
}

func (s *ZprZprPolicyResourceCrud) Create() error {
	request := oci_zpr.CreateZprPolicyRequest{}

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

	if statements, ok := s.D.GetOkExists("statements"); ok {
		interfaces := statements.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("statements") {
			request.Statements = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "zpr")

	response, err := s.Client.CreateZprPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getZprPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "zpr"), oci_zpr.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ZprZprPolicyResourceCrud) getZprPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_zpr.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	zprPolicyId, err := zprPolicyWaitForWorkRequest(workId, "zprpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*zprPolicyId)

	return s.Get()
}

func zprPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "zpr", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_zpr.GetZprPolicyWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func zprPolicyWaitForWorkRequest(wId *string, entityType string, action oci_zpr.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_zpr.ZprClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "zpr")
	retryPolicy.ShouldRetryOperation = zprPolicyWorkRequestShouldRetryFunc(timeout)

	response := oci_zpr.GetZprPolicyWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_zpr.WorkRequestStatusInProgress),
			string(oci_zpr.WorkRequestStatusAccepted),
			string(oci_zpr.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_zpr.WorkRequestStatusSucceeded),
			string(oci_zpr.WorkRequestStatusFailed),
			string(oci_zpr.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetZprPolicyWorkRequest(context.Background(),
				oci_zpr.GetZprPolicyWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_zpr.WorkRequestStatusFailed || response.Status == oci_zpr.WorkRequestStatusCanceled {
		return nil, getErrorFromZprZprPolicyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromZprZprPolicyWorkRequest(client *oci_zpr.ZprClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_zpr.ActionTypeEnum) error {
	response, err := client.ListZprPolicyWorkRequestErrors(context.Background(),
		oci_zpr.ListZprPolicyWorkRequestErrorsRequest{
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

func (s *ZprZprPolicyResourceCrud) Get() error {
	request := oci_zpr.GetZprPolicyRequest{}

	tmp := s.D.Id()
	request.ZprPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "zpr")

	response, err := s.Client.GetZprPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ZprPolicy
	return nil
}

func (s *ZprZprPolicyResourceCrud) Update() error {
	request := oci_zpr.UpdateZprPolicyRequest{}

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

	if statements, ok := s.D.GetOkExists("statements"); ok {
		interfaces := statements.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("statements") {
			request.Statements = tmp
		}
	}

	tmp := s.D.Id()
	request.ZprPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "zpr")

	response, err := s.Client.UpdateZprPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getZprPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "zpr"), oci_zpr.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ZprZprPolicyResourceCrud) Delete() error {
	request := oci_zpr.DeleteZprPolicyRequest{}

	tmp := s.D.Id()
	request.ZprPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "zpr")

	response, err := s.Client.DeleteZprPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := zprPolicyWaitForWorkRequest(workId, "zprpolicy",
		oci_zpr.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ZprZprPolicyResourceCrud) SetData() error {
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("statements", s.Res.Statements)

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

func ZprPolicySummaryToMap(obj oci_zpr.ZprPolicySummary) map[string]interface{} {
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	result["statements"] = obj.Statements

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
