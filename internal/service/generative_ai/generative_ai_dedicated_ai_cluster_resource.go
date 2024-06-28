// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiDedicatedAiClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("80m"),
			Update: tfresource.GetTimeoutDuration("80m"),
			Delete: tfresource.GetTimeoutDuration("20m"),
		},
		Create: createGenerativeAiDedicatedAiCluster,
		Read:   readGenerativeAiDedicatedAiCluster,
		Update: updateGenerativeAiDedicatedAiCluster,
		Delete: deleteGenerativeAiDedicatedAiCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"unit_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"unit_shape": {
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

			// Computed
			"capacity": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"capacity_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_endpoint_capacity": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used_endpoint_capacity": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
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

func createGenerativeAiDedicatedAiCluster(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiDedicatedAiClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.CreateResource(d, sync)
}

func readGenerativeAiDedicatedAiCluster(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiDedicatedAiClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.ReadResource(sync)
}

func updateGenerativeAiDedicatedAiCluster(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiDedicatedAiClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteGenerativeAiDedicatedAiCluster(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiDedicatedAiClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GenerativeAiDedicatedAiClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai.GenerativeAiClient
	Res                    *oci_generative_ai.DedicatedAiCluster
	DisableNotFoundRetries bool
}

func (s *GenerativeAiDedicatedAiClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiDedicatedAiClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai.DedicatedAiClusterLifecycleStateCreating),
	}
}

func (s *GenerativeAiDedicatedAiClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai.DedicatedAiClusterLifecycleStateActive),
		string(oci_generative_ai.DedicatedAiClusterLifecycleStateNeedsAttention),
	}
}

func (s *GenerativeAiDedicatedAiClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai.DedicatedAiClusterLifecycleStateDeleting),
	}
}

func (s *GenerativeAiDedicatedAiClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai.DedicatedAiClusterLifecycleStateDeleted),
	}
}

func (s *GenerativeAiDedicatedAiClusterResourceCrud) Create() error {
	request := oci_generative_ai.CreateDedicatedAiClusterRequest{}

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

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_generative_ai.DedicatedAiClusterTypeEnum(type_.(string))
	}

	if unitCount, ok := s.D.GetOkExists("unit_count"); ok {
		tmp := unitCount.(int)
		request.UnitCount = &tmp
	}

	if unitShape, ok := s.D.GetOkExists("unit_shape"); ok {
		request.UnitShape = oci_generative_ai.DedicatedAiClusterUnitShapeEnum(unitShape.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.CreateDedicatedAiCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDedicatedAiClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiDedicatedAiClusterResourceCrud) getDedicatedAiClusterFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	dedicatedAiClusterId, err := dedicatedAiClusterWaitForWorkRequest(workId, "dedicatedaicluster",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*dedicatedAiClusterId)

	return s.Get()
}

func dedicatedAiClusterWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func dedicatedAiClusterWaitForWorkRequest(wId *string, entityType string, action oci_generative_ai.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai.GenerativeAiClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai")
	retryPolicy.ShouldRetryOperation = dedicatedAiClusterWorkRequestShouldRetryFunc(timeout)

	response := oci_generative_ai.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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
		return nil, getErrorFromGenerativeAiDedicatedAiClusterWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiDedicatedAiClusterWorkRequest(client *oci_generative_ai.GenerativeAiClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai.ActionTypeEnum) error {
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

func (s *GenerativeAiDedicatedAiClusterResourceCrud) Get() error {
	request := oci_generative_ai.GetDedicatedAiClusterRequest{}

	tmp := s.D.Id()
	request.DedicatedAiClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.GetDedicatedAiCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DedicatedAiCluster
	return nil
}

func (s *GenerativeAiDedicatedAiClusterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_generative_ai.UpdateDedicatedAiClusterRequest{}

	tmp := s.D.Id()
	request.DedicatedAiClusterId = &tmp

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

	if unitCount, ok := s.D.GetOkExists("unit_count"); ok {
		tmp := unitCount.(int)
		request.UnitCount = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.UpdateDedicatedAiCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDedicatedAiClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GenerativeAiDedicatedAiClusterResourceCrud) Delete() error {
	request := oci_generative_ai.DeleteDedicatedAiClusterRequest{}

	tmp := s.D.Id()
	request.DedicatedAiClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.DeleteDedicatedAiCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := dedicatedAiClusterWaitForWorkRequest(workId, "dedicatedaicluster",
		oci_generative_ai.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiDedicatedAiClusterResourceCrud) SetData() error {
	if s.Res.Capacity != nil {
		capacityArray := []interface{}{}
		if capacityMap := DedicatedAiClusterCapacityToMap(&s.Res.Capacity); capacityMap != nil {
			capacityArray = append(capacityArray, capacityMap)
		}
		s.D.Set("capacity", capacityArray)
	} else {
		s.D.Set("capacity", nil)
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

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
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

	s.D.Set("type", s.Res.Type)

	if s.Res.UnitCount != nil {
		s.D.Set("unit_count", *s.Res.UnitCount)
	}

	s.D.Set("unit_shape", s.Res.UnitShape)

	return nil
}

func DedicatedAiClusterCapacityToMap(obj *oci_generative_ai.DedicatedAiClusterCapacity) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai.DedicatedAiClusterHostingCapacity:
		result["capacity_type"] = "HOSTING_CAPACITY"

		if v.TotalEndpointCapacity != nil {
			result["total_endpoint_capacity"] = int(*v.TotalEndpointCapacity)
		}

		if v.UsedEndpointCapacity != nil {
			result["used_endpoint_capacity"] = int(*v.UsedEndpointCapacity)
		}
	default:
		log.Printf("[WARN] Received 'capacity_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DedicatedAiClusterSummaryToMap(obj oci_generative_ai.DedicatedAiClusterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Capacity != nil {
		capacityArray := []interface{}{}
		if capacityMap := DedicatedAiClusterCapacityToMap(&obj.Capacity); capacityMap != nil {
			capacityArray = append(capacityArray, capacityMap)
		}
		result["capacity"] = capacityArray
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

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

	result["type"] = string(obj.Type)

	if obj.UnitCount != nil {
		result["unit_count"] = int(*obj.UnitCount)
	}

	result["unit_shape"] = string(obj.UnitShape)

	return result
}

func (s *GenerativeAiDedicatedAiClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_generative_ai.ChangeDedicatedAiClusterCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DedicatedAiClusterId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	_, err := s.Client.ChangeDedicatedAiClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
