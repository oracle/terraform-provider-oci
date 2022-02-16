// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_jms "github.com/oracle/oci-go-sdk/v58/jms"
)

func JmsFleetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createJmsFleet,
		Read:     readJmsFleet,
		Update:   updateJmsFleet,
		Delete:   deleteJmsFleet,
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

			// Computed
			"approximate_application_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"approximate_installation_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"approximate_jre_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"approximate_managed_instance_count": {
				Type:     schema.TypeInt,
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

func createJmsFleet(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.CreateResource(d, sync)
}

func readJmsFleet(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

func updateJmsFleet(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteJmsFleet(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type JmsFleetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_jms.JavaManagementServiceClient
	Res                    *oci_jms.Fleet
	DisableNotFoundRetries bool
}

func (s *JmsFleetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *JmsFleetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_jms.LifecycleStateCreating),
	}
}

func (s *JmsFleetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_jms.LifecycleStateActive),
	}
}

func (s *JmsFleetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_jms.LifecycleStateDeleting),
	}
}

func (s *JmsFleetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_jms.LifecycleStateDeleted),
	}
}

func (s *JmsFleetResourceCrud) Create() error {
	request := oci_jms.CreateFleetRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.CreateFleet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFleetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms"), oci_jms.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *JmsFleetResourceCrud) getFleetFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_jms.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fleetId, err := fleetWaitForWorkRequest(workId, "fleet",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		log.Printf("[DEBUG] creation failed for the workrequest: %v for identifier: %v\n", workId, fleetId)
		return err
	}
	s.D.SetId(*fleetId)

	return s.Get()
}

func fleetWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "jms", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_jms.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func fleetWaitForWorkRequest(wId *string, entityType string, action oci_jms.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_jms.JavaManagementServiceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "jms")
	retryPolicy.ShouldRetryOperation = fleetWorkRequestShouldRetryFunc(timeout)

	response := oci_jms.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_jms.OperationStatusInProgress),
			string(oci_jms.OperationStatusAccepted),
			string(oci_jms.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_jms.OperationStatusSucceeded),
			string(oci_jms.OperationStatusFailed),
			string(oci_jms.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_jms.GetWorkRequestRequest{
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

	// The workrequest didn't do all its intended tasks, if the errors is set; so we should check for it
	var workRequestErr error

	return identifier, workRequestErr
}

func (s *JmsFleetResourceCrud) Get() error {
	request := oci_jms.GetFleetRequest{}

	tmp := s.D.Id()
	request.FleetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.GetFleet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Fleet
	return nil
}

func (s *JmsFleetResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_jms.UpdateFleetRequest{}

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

	tmp := s.D.Id()
	request.FleetId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.UpdateFleet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFleetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms"), oci_jms.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *JmsFleetResourceCrud) Delete() error {
	request := oci_jms.DeleteFleetRequest{}

	tmp := s.D.Id()
	request.FleetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	_, err := s.Client.DeleteFleet(context.Background(), request)
	return err
}

func (s *JmsFleetResourceCrud) SetData() error {
	if s.Res.ApproximateApplicationCount != nil {
		s.D.Set("approximate_application_count", *s.Res.ApproximateApplicationCount)
	}

	if s.Res.ApproximateInstallationCount != nil {
		s.D.Set("approximate_installation_count", *s.Res.ApproximateInstallationCount)
	}

	if s.Res.ApproximateJreCount != nil {
		s.D.Set("approximate_jre_count", *s.Res.ApproximateJreCount)
	}

	if s.Res.ApproximateManagedInstanceCount != nil {
		s.D.Set("approximate_managed_instance_count", *s.Res.ApproximateManagedInstanceCount)
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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func FleetSummaryToMap(obj oci_jms.FleetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApproximateApplicationCount != nil {
		result["approximate_application_count"] = int(*obj.ApproximateApplicationCount)
	}

	if obj.ApproximateInstallationCount != nil {
		result["approximate_installation_count"] = int(*obj.ApproximateInstallationCount)
	}

	if obj.ApproximateJreCount != nil {
		result["approximate_jre_count"] = int(*obj.ApproximateJreCount)
	}

	if obj.ApproximateManagedInstanceCount != nil {
		result["approximate_managed_instance_count"] = int(*obj.ApproximateManagedInstanceCount)
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

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *JmsFleetResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_jms.ChangeFleetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.FleetId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	_, err := s.Client.ChangeFleetCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
