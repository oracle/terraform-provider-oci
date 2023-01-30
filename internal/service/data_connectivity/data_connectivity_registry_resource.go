// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v65/dataconnectivity"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataConnectivityRegistryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataConnectivityRegistry,
		Read:     readDataConnectivityRegistry,
		Update:   updateDataConnectivityRegistry,
		Delete:   deleteDataConnectivityRegistry,
		Schema: map[string]*schema.Schema{
			// Required
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
			"updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataConnectivityRegistry(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDataConnectivityRegistry(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDataConnectivityRegistry(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataConnectivityRegistry(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataConnectivityRegistryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_connectivity.DataConnectivityManagementClient
	Res                    *oci_data_connectivity.Registry
	DisableNotFoundRetries bool
}

func (s *DataConnectivityRegistryResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataConnectivityRegistryResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_connectivity.RegistryLifecycleStateCreating),
	}
}

func (s *DataConnectivityRegistryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_connectivity.RegistryLifecycleStateActive),
	}
}

func (s *DataConnectivityRegistryResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_connectivity.RegistryLifecycleStateDeleting),
	}
}

func (s *DataConnectivityRegistryResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_connectivity.RegistryLifecycleStateDeleted),
	}
}

func (s *DataConnectivityRegistryResourceCrud) Create() error {
	request := oci_data_connectivity.CreateRegistryRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.CreateRegistry(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRegistryFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity"), oci_data_connectivity.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataConnectivityRegistryResourceCrud) getRegistryFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_connectivity.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	registryId, err := registryWaitForWorkRequest(workId, "dcmsregistry",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*registryId)

	return s.Get()
}

func registryWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_connectivity", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_connectivity.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func registryWaitForWorkRequest(wId *string, entityType string, action oci_data_connectivity.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_connectivity.DataConnectivityManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_connectivity")
	retryPolicy.ShouldRetryOperation = registryWorkRequestShouldRetryFunc(timeout)

	response := oci_data_connectivity.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_data_connectivity.WorkRequestStatusInProgress),
			string(oci_data_connectivity.WorkRequestStatusAccepted),
			string(oci_data_connectivity.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_data_connectivity.WorkRequestStatusSucceeded),
			string(oci_data_connectivity.WorkRequestStatusFailed),
			string(oci_data_connectivity.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_data_connectivity.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_data_connectivity.WorkRequestStatusFailed || response.Status == oci_data_connectivity.WorkRequestStatusCanceled {
		return nil, getErrorFromDataConnectivityRegistryWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataConnectivityRegistryWorkRequest(client *oci_data_connectivity.DataConnectivityManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_connectivity.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_data_connectivity.ListWorkRequestErrorsRequest{
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

func (s *DataConnectivityRegistryResourceCrud) Get() error {
	request := oci_data_connectivity.GetRegistryRequest{}

	tmp := s.D.Id()
	request.RegistryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.GetRegistry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Registry
	return nil
}

func (s *DataConnectivityRegistryResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_connectivity.UpdateRegistryRequest{}

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

	tmp := s.D.Id()
	request.RegistryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.UpdateRegistry(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRegistryFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity"), oci_data_connectivity.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataConnectivityRegistryResourceCrud) Delete() error {
	request := oci_data_connectivity.DeleteRegistryRequest{}

	if isForceOperation, ok := s.D.GetOkExists("is_force_operation"); ok {
		tmp := isForceOperation.(bool)
		request.IsForceOperation = &tmp
	}

	tmp := s.D.Id()
	request.RegistryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.DeleteRegistry(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := registryWaitForWorkRequest(workId, "dcmsregistry",
		oci_data_connectivity.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataConnectivityRegistryResourceCrud) SetData() error {
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

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UpdatedBy != nil {
		s.D.Set("updated_by", *s.Res.UpdatedBy)
	}

	return nil
}

func RegistrySummaryToMap(obj oci_data_connectivity.RegistrySummary) map[string]interface{} {
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

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.StateMessage != nil {
		result["state_message"] = string(*obj.StateMessage)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.UpdatedBy != nil {
		result["updated_by"] = string(*obj.UpdatedBy)
	}

	return result
}

func (s *DataConnectivityRegistryResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_connectivity.ChangeRegistryCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.RegistryId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.ChangeRegistryCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRegistryFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity"), oci_data_connectivity.WorkRequestResourceActionTypeMoved, s.D.Timeout(schema.TimeoutUpdate))
}
