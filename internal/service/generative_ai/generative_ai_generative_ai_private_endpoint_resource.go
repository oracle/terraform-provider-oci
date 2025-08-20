// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiGenerativeAiPrivateEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createGenerativeAiGenerativeAiPrivateEndpoint,
		Read:     readGenerativeAiGenerativeAiPrivateEndpoint,
		Update:   updateGenerativeAiGenerativeAiPrivateEndpoint,
		Delete:   deleteGenerativeAiGenerativeAiPrivateEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dns_prefix": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
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
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"previous_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_endpoint_ip": {
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

func createGenerativeAiGenerativeAiPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiGenerativeAiPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.CreateResource(d, sync)
}

func readGenerativeAiGenerativeAiPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiGenerativeAiPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.ReadResource(sync)
}

func updateGenerativeAiGenerativeAiPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiGenerativeAiPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteGenerativeAiGenerativeAiPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiGenerativeAiPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GenerativeAiGenerativeAiPrivateEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai.GenerativeAiClient
	Res                    *oci_generative_ai.GenerativeAiPrivateEndpoint
	DisableNotFoundRetries bool
}

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai.GenerativeAiPrivateEndpointLifecycleStateCreating),
	}
}

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai.GenerativeAiPrivateEndpointLifecycleStateActive),
	}
}

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai.GenerativeAiPrivateEndpointLifecycleStateDeleting),
	}
}

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai.GenerativeAiPrivateEndpointLifecycleStateDeleted),
	}
}

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) Create() error {
	request := oci_generative_ai.CreateGenerativeAiPrivateEndpointRequest{}

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

	if dnsPrefix, ok := s.D.GetOkExists("dns_prefix"); ok {
		tmp := dnsPrefix.(string)
		request.DnsPrefix = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.CreateGenerativeAiPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getGenerativeAiPrivateEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) getGenerativeAiPrivateEndpointFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	generativeAiPrivateEndpointId, err := generativeAiPrivateEndpointWaitForWorkRequest(workId, "generativeaiprivateendpoint",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*generativeAiPrivateEndpointId)

	return s.Get()
}

func generativeAiPrivateEndpointWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "generative_ai", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_generative_ai.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func generativeAiPrivateEndpointWaitForWorkRequest(wId *string, entityType string, action oci_generative_ai.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai.GenerativeAiClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai")
	retryPolicy.ShouldRetryOperation = generativeAiPrivateEndpointWorkRequestShouldRetryFunc(timeout)

	response := oci_generative_ai.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_generative_ai.OperationStatusInProgress),
			string(oci_generative_ai.OperationStatusAccepted),
			string(oci_generative_ai.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_generative_ai.OperationStatusSucceeded),
			string(oci_generative_ai.OperationStatusFailed),
			string(oci_generative_ai.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_generative_ai.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_generative_ai.OperationStatusFailed || response.Status == oci_generative_ai.OperationStatusCanceled {
		return nil, getErrorFromGenerativeAiGenerativeAiPrivateEndpointWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiGenerativeAiPrivateEndpointWorkRequest(client *oci_generative_ai.GenerativeAiClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_generative_ai.ListWorkRequestErrorsRequest{
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

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) Get() error {
	request := oci_generative_ai.GetGenerativeAiPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.GenerativeAiPrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.GetGenerativeAiPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.GenerativeAiPrivateEndpoint
	return nil
}

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_generative_ai.UpdateGenerativeAiPrivateEndpointRequest{}

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

	if dnsPrefix, ok := s.D.GetOkExists("dns_prefix"); ok {
		tmp := dnsPrefix.(string)
		request.DnsPrefix = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.GenerativeAiPrivateEndpointId = &tmp

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.UpdateGenerativeAiPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getGenerativeAiPrivateEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) Delete() error {
	request := oci_generative_ai.DeleteGenerativeAiPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.GenerativeAiPrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.DeleteGenerativeAiPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := generativeAiPrivateEndpointWaitForWorkRequest(workId, "generativeaiprivateendpoint",
		oci_generative_ai.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.Fqdn != nil {
		s.D.Set("fqdn", *s.Res.Fqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.PreviousState != nil {
		s.D.Set("previous_state", *s.Res.PreviousState)
	}

	if s.Res.PrivateEndpointIp != nil {
		s.D.Set("private_endpoint_ip", *s.Res.PrivateEndpointIp)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

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

func GenerativeAiPrivateEndpointSummaryToMap(obj oci_generative_ai.GenerativeAiPrivateEndpointSummary) map[string]interface{} {
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

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Fqdn != nil {
		result["fqdn"] = string(*obj.Fqdn)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.PrivateEndpointIp != nil {
		result["private_endpoint_ip"] = string(*obj.PrivateEndpointIp)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

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

func (s *GenerativeAiGenerativeAiPrivateEndpointResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_generative_ai.ChangeGenerativeAiPrivateEndpointCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.GenerativeAiPrivateEndpointId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.ChangeGenerativeAiPrivateEndpointCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getGenerativeAiPrivateEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
