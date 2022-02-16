// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v58/aianomalydetection"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func AiAnomalyDetectionAiPrivateEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiAnomalyDetectionAiPrivateEndpoint,
		Read:     readAiAnomalyDetectionAiPrivateEndpoint,
		Update:   updateAiAnomalyDetectionAiPrivateEndpoint,
		Delete:   deleteAiAnomalyDetectionAiPrivateEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dns_zones": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"attached_data_assets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createAiAnomalyDetectionAiPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionAiPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.CreateResource(d, sync)
}

func readAiAnomalyDetectionAiPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionAiPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

func updateAiAnomalyDetectionAiPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionAiPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiAnomalyDetectionAiPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionAiPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiAnomalyDetectionAiPrivateEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res                    *oci_ai_anomaly_detection.AiPrivateEndpoint
	DisableNotFoundRetries bool
}

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_anomaly_detection.AiPrivateEndpointLifecycleStateCreating),
	}
}

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_anomaly_detection.AiPrivateEndpointLifecycleStateActive),
	}
}

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_anomaly_detection.AiPrivateEndpointLifecycleStateDeleting),
	}
}

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_anomaly_detection.AiPrivateEndpointLifecycleStateDeleted),
	}
}

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) Create() error {
	request := oci_ai_anomaly_detection.CreateAiPrivateEndpointRequest{}

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

	if dnsZones, ok := s.D.GetOkExists("dns_zones"); ok {
		interfaces := dnsZones.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns_zones") {
			request.DnsZones = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.CreateAiPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAiPrivateEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection"), oci_ai_anomaly_detection.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) getAiPrivateEndpointFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_anomaly_detection.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	aiPrivateEndpointId, err := aiPrivateEndpointWaitForWorkRequest(workId, "aiprivateendpoint",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, aiPrivateEndpointId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_ai_anomaly_detection.CancelWorkRequestRequest{
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
	s.D.SetId(*aiPrivateEndpointId)

	return s.Get()
}

func aiPrivateEndpointWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ai_anomaly_detection", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ai_anomaly_detection.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func aiPrivateEndpointWaitForWorkRequest(wId *string, entityType string, action oci_ai_anomaly_detection.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_anomaly_detection.AnomalyDetectionClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_anomaly_detection")
	retryPolicy.ShouldRetryOperation = aiPrivateEndpointWorkRequestShouldRetryFunc(timeout)

	response := oci_ai_anomaly_detection.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_ai_anomaly_detection.OperationStatusInProgress),
			string(oci_ai_anomaly_detection.OperationStatusAccepted),
			string(oci_ai_anomaly_detection.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ai_anomaly_detection.OperationStatusSucceeded),
			string(oci_ai_anomaly_detection.OperationStatusFailed),
			string(oci_ai_anomaly_detection.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ai_anomaly_detection.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ai_anomaly_detection.OperationStatusFailed || response.Status == oci_ai_anomaly_detection.OperationStatusCanceled {
		return nil, getErrorFromAiAnomalyDetectionAiPrivateEndpointWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAiAnomalyDetectionAiPrivateEndpointWorkRequest(client *oci_ai_anomaly_detection.AnomalyDetectionClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_anomaly_detection.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ai_anomaly_detection.ListWorkRequestErrorsRequest{
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

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) Get() error {
	request := oci_ai_anomaly_detection.GetAiPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.AiPrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.GetAiPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AiPrivateEndpoint
	return nil
}

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_anomaly_detection.UpdateAiPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.AiPrivateEndpointId = &tmp

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

	if dnsZones, ok := s.D.GetOkExists("dns_zones"); ok {
		interfaces := dnsZones.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns_zones") {
			request.DnsZones = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.UpdateAiPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAiPrivateEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection"), oci_ai_anomaly_detection.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) Delete() error {
	request := oci_ai_anomaly_detection.DeleteAiPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.AiPrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.DeleteAiPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := aiPrivateEndpointWaitForWorkRequest(workId, "aiprivateendpoint",
		oci_ai_anomaly_detection.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) SetData() error {
	s.D.Set("attached_data_assets", s.Res.AttachedDataAssets)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("dns_zones", s.Res.DnsZones)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
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

func AiPrivateEndpointSummaryToMap(obj oci_ai_anomaly_detection.AiPrivateEndpointSummary) map[string]interface{} {
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

	result["dns_zones"] = obj.DnsZones

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

func (s *AiAnomalyDetectionAiPrivateEndpointResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_anomaly_detection.ChangeAiPrivateEndpointCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AiPrivateEndpointId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.ChangeAiPrivateEndpointCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAiPrivateEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection"), oci_ai_anomaly_detection.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
