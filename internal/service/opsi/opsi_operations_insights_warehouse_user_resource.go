// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

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

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v58/opsi"
)

func OpsiOperationsInsightsWarehouseUserResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiOperationsInsightsWarehouseUser,
		Read:     readOpsiOperationsInsightsWarehouseUser,
		Update:   updateOpsiOperationsInsightsWarehouseUser,
		Delete:   deleteOpsiOperationsInsightsWarehouseUser,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"connection_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"is_awr_data_access": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"operations_insights_warehouse_id": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_em_data_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_opsi_data_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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

func createOpsiOperationsInsightsWarehouseUser(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.CreateResource(d, sync)
}

func readOpsiOperationsInsightsWarehouseUser(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

func updateOpsiOperationsInsightsWarehouseUser(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOpsiOperationsInsightsWarehouseUser(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpsiOperationsInsightsWarehouseUserResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *oci_opsi.OperationsInsightsWarehouseUser
	DisableNotFoundRetries bool
}

func (s *OpsiOperationsInsightsWarehouseUserResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OpsiOperationsInsightsWarehouseUserResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opsi.OperationsInsightsWarehouseUserLifecycleStateCreating),
	}
}

func (s *OpsiOperationsInsightsWarehouseUserResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opsi.OperationsInsightsWarehouseUserLifecycleStateActive),
	}
}

func (s *OpsiOperationsInsightsWarehouseUserResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opsi.OperationsInsightsWarehouseUserLifecycleStateDeleting),
	}
}

func (s *OpsiOperationsInsightsWarehouseUserResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opsi.OperationsInsightsWarehouseUserLifecycleStateDeleted),
	}
}

func (s *OpsiOperationsInsightsWarehouseUserResourceCrud) Create() error {
	request := oci_opsi.CreateOperationsInsightsWarehouseUserRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if connectionPassword, ok := s.D.GetOkExists("connection_password"); ok {
		tmp := connectionPassword.(string)
		request.ConnectionPassword = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAwrDataAccess, ok := s.D.GetOkExists("is_awr_data_access"); ok {
		tmp := isAwrDataAccess.(bool)
		request.IsAwrDataAccess = &tmp
	}

	if isEmDataAccess, ok := s.D.GetOkExists("is_em_data_access"); ok {
		tmp := isEmDataAccess.(bool)
		request.IsEmDataAccess = &tmp
	}

	if isOpsiDataAccess, ok := s.D.GetOkExists("is_opsi_data_access"); ok {
		tmp := isOpsiDataAccess.(bool)
		request.IsOpsiDataAccess = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if operationsInsightsWarehouseId, ok := s.D.GetOkExists("operations_insights_warehouse_id"); ok {
		tmp := operationsInsightsWarehouseId.(string)
		request.OperationsInsightsWarehouseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.CreateOperationsInsightsWarehouseUser(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOperationsInsightsWarehouseUserFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OpsiOperationsInsightsWarehouseUserResourceCrud) getOperationsInsightsWarehouseUserFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	operationsInsightsWarehouseUserId, err := operationsInsightsWarehouseUserWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*operationsInsightsWarehouseUserId)

	return s.Get()
}

func operationsInsightsWarehouseUserWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "opsi", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_opsi.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func operationsInsightsWarehouseUserWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = operationsInsightsWarehouseUserWorkRequestShouldRetryFunc(timeout)

	response := oci_opsi.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_opsi.OperationStatusInProgress),
			string(oci_opsi.OperationStatusAccepted),
			string(oci_opsi.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_opsi.OperationStatusSucceeded),
			string(oci_opsi.OperationStatusFailed),
			string(oci_opsi.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_opsi.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_opsi.OperationStatusFailed || response.Status == oci_opsi.OperationStatusCanceled {
		return nil, getErrorFromOpsiOperationsInsightsWarehouseUserWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiOperationsInsightsWarehouseUserWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_opsi.ListWorkRequestErrorsRequest{
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

func (s *OpsiOperationsInsightsWarehouseUserResourceCrud) Get() error {
	request := oci_opsi.GetOperationsInsightsWarehouseUserRequest{}

	tmp := s.D.Id()
	request.OperationsInsightsWarehouseUserId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.GetOperationsInsightsWarehouseUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OperationsInsightsWarehouseUser
	return nil
}

func (s *OpsiOperationsInsightsWarehouseUserResourceCrud) Update() error {
	request := oci_opsi.UpdateOperationsInsightsWarehouseUserRequest{}

	if connectionPassword, ok := s.D.GetOkExists("connection_password"); ok {
		tmp := connectionPassword.(string)
		request.ConnectionPassword = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAwrDataAccess, ok := s.D.GetOkExists("is_awr_data_access"); ok {
		tmp := isAwrDataAccess.(bool)
		request.IsAwrDataAccess = &tmp
	}

	if isEmDataAccess, ok := s.D.GetOkExists("is_em_data_access"); ok {
		tmp := isEmDataAccess.(bool)
		request.IsEmDataAccess = &tmp
	}

	if isOpsiDataAccess, ok := s.D.GetOkExists("is_opsi_data_access"); ok {
		tmp := isOpsiDataAccess.(bool)
		request.IsOpsiDataAccess = &tmp
	}

	tmp := s.D.Id()
	request.OperationsInsightsWarehouseUserId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.UpdateOperationsInsightsWarehouseUser(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOperationsInsightsWarehouseUserFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpsiOperationsInsightsWarehouseUserResourceCrud) Delete() error {
	request := oci_opsi.DeleteOperationsInsightsWarehouseUserRequest{}

	tmp := s.D.Id()
	request.OperationsInsightsWarehouseUserId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DeleteOperationsInsightsWarehouseUser(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := operationsInsightsWarehouseUserWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpsiOperationsInsightsWarehouseUserResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionPassword != nil {
		s.D.Set("connection_password", *s.Res.ConnectionPassword)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAwrDataAccess != nil {
		s.D.Set("is_awr_data_access", *s.Res.IsAwrDataAccess)
	}

	if s.Res.IsEmDataAccess != nil {
		s.D.Set("is_em_data_access", *s.Res.IsEmDataAccess)
	}

	if s.Res.IsOpsiDataAccess != nil {
		s.D.Set("is_opsi_data_access", *s.Res.IsOpsiDataAccess)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.OperationsInsightsWarehouseId != nil {
		s.D.Set("operations_insights_warehouse_id", *s.Res.OperationsInsightsWarehouseId)
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

func OperationsInsightsWarehouseUserSummaryToMap(obj oci_opsi.OperationsInsightsWarehouseUserSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConnectionPassword != nil {
		result["connection_password"] = string(*obj.ConnectionPassword)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsAwrDataAccess != nil {
		result["is_awr_data_access"] = bool(*obj.IsAwrDataAccess)
	}

	if obj.IsEmDataAccess != nil {
		result["is_em_data_access"] = bool(*obj.IsEmDataAccess)
	}

	if obj.IsOpsiDataAccess != nil {
		result["is_opsi_data_access"] = bool(*obj.IsOpsiDataAccess)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.OperationsInsightsWarehouseId != nil {
		result["operations_insights_warehouse_id"] = string(*obj.OperationsInsightsWarehouseId)
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
