// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudOracleDbGcpIdentityConnectorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDbmulticloudOracleDbGcpIdentityConnector,
		Read:     readDbmulticloudOracleDbGcpIdentityConnector,
		Update:   updateDbmulticloudOracleDbGcpIdentityConnector,
		Delete:   deleteDbmulticloudOracleDbGcpIdentityConnector,
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
			"gcp_location": {
				Type:     schema.TypeString,
				Required: true,
			},
			"gcp_resource_service_agent_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"gcp_workload_identity_pool_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"gcp_workload_identity_provider_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"issuer_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_id": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"gcp_identity_connectivity_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gcp_nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"host_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_checked": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_state_details": {
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

func createDbmulticloudOracleDbGcpIdentityConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbGcpIdentityConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudGCPProviderClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readDbmulticloudOracleDbGcpIdentityConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbGcpIdentityConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudGCPProviderClient()

	return tfresource.ReadResource(sync)
}

func updateDbmulticloudOracleDbGcpIdentityConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbGcpIdentityConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudGCPProviderClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDbmulticloudOracleDbGcpIdentityConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbGcpIdentityConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudGCPProviderClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type DbmulticloudOracleDbGcpIdentityConnectorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dbmulticloud.DbMulticloudGCPProviderClient
	Res                    *oci_dbmulticloud.OracleDbGcpIdentityConnector
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_dbmulticloud.WorkRequestClient
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbGcpIdentityConnectorLifecycleStateCreating),
	}
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbGcpIdentityConnectorLifecycleStateActive),
	}
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbGcpIdentityConnectorLifecycleStateDeleting),
	}
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbGcpIdentityConnectorLifecycleStateDeleted),
	}
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) Create() error {
	request := oci_dbmulticloud.CreateOracleDbGcpIdentityConnectorRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gcpLocation, ok := s.D.GetOkExists("gcp_location"); ok {
		tmp := gcpLocation.(string)
		request.GcpLocation = &tmp
	}

	if gcpResourceServiceAgentId, ok := s.D.GetOkExists("gcp_resource_service_agent_id"); ok {
		tmp := gcpResourceServiceAgentId.(string)
		request.GcpResourceServiceAgentId = &tmp
	}

	if gcpWorkloadIdentityPoolId, ok := s.D.GetOkExists("gcp_workload_identity_pool_id"); ok {
		tmp := gcpWorkloadIdentityPoolId.(string)
		request.GcpWorkloadIdentityPoolId = &tmp
	}

	if gcpWorkloadIdentityProviderId, ok := s.D.GetOkExists("gcp_workload_identity_provider_id"); ok {
		tmp := gcpWorkloadIdentityProviderId.(string)
		request.GcpWorkloadIdentityProviderId = &tmp
	}

	if issuerUrl, ok := s.D.GetOkExists("issuer_url"); ok {
		tmp := issuerUrl.(string)
		request.IssuerUrl = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.CreateOracleDbGcpIdentityConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getOracleDbGcpIdentityConnectorFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) getOracleDbGcpIdentityConnectorFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_dbmulticloud.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	oracleDbGcpIdentityConnectorId, err := oracleDbGcpIdentityConnectorWaitForWorkRequest(workId, "oracledbgcpconnector",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, oracleDbGcpIdentityConnectorId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(context.Background(),
			oci_dbmulticloud.CancelWorkRequestRequest{
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
	s.D.SetId(*oracleDbGcpIdentityConnectorId)

	return s.Get()
}

func oracleDbGcpIdentityConnectorWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "dbmulticloud", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_dbmulticloud.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func oracleDbGcpIdentityConnectorWaitForWorkRequest(wId *string, entityType string, action oci_dbmulticloud.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_dbmulticloud.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "dbmulticloud")
	retryPolicy.ShouldRetryOperation = oracleDbGcpIdentityConnectorWorkRequestShouldRetryFunc(timeout)

	response := oci_dbmulticloud.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_dbmulticloud.OperationStatusInProgress),
			string(oci_dbmulticloud.OperationStatusAccepted),
			string(oci_dbmulticloud.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_dbmulticloud.OperationStatusSucceeded),
			string(oci_dbmulticloud.OperationStatusFailed),
			string(oci_dbmulticloud.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_dbmulticloud.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_dbmulticloud.OperationStatusFailed || response.Status == oci_dbmulticloud.OperationStatusCanceled {
		return nil, getErrorFromDbmulticloudOracleDbGcpIdentityConnectorWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDbmulticloudOracleDbGcpIdentityConnectorWorkRequest(client *oci_dbmulticloud.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_dbmulticloud.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_dbmulticloud.ListWorkRequestErrorsRequest{
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

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) Get() error {
	request := oci_dbmulticloud.GetOracleDbGcpIdentityConnectorRequest{}

	tmp := s.D.Id()
	request.OracleDbGcpIdentityConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.GetOracleDbGcpIdentityConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OracleDbGcpIdentityConnector
	return nil
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dbmulticloud.UpdateOracleDbGcpIdentityConnectorRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gcpLocation, ok := s.D.GetOkExists("gcp_location"); ok {
		tmp := gcpLocation.(string)
		request.GcpLocation = &tmp
	}

	if gcpResourceServiceAgentId, ok := s.D.GetOkExists("gcp_resource_service_agent_id"); ok {
		tmp := gcpResourceServiceAgentId.(string)
		request.GcpResourceServiceAgentId = &tmp
	}

	if gcpWorkloadIdentityPoolId, ok := s.D.GetOkExists("gcp_workload_identity_pool_id"); ok {
		tmp := gcpWorkloadIdentityPoolId.(string)
		request.GcpWorkloadIdentityPoolId = &tmp
	}

	if gcpWorkloadIdentityProviderId, ok := s.D.GetOkExists("gcp_workload_identity_provider_id"); ok {
		tmp := gcpWorkloadIdentityProviderId.(string)
		request.GcpWorkloadIdentityProviderId = &tmp
	}

	if issuerUrl, ok := s.D.GetOkExists("issuer_url"); ok {
		tmp := issuerUrl.(string)
		request.IssuerUrl = &tmp
	}

	tmp := s.D.Id()
	request.OracleDbGcpIdentityConnectorId = &tmp

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.UpdateOracleDbGcpIdentityConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOracleDbGcpIdentityConnectorFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) Delete() error {
	request := oci_dbmulticloud.DeleteOracleDbGcpIdentityConnectorRequest{}

	tmp := s.D.Id()
	request.OracleDbGcpIdentityConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.DeleteOracleDbGcpIdentityConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := oracleDbGcpIdentityConnectorWaitForWorkRequest(workId, "oracledbgcpconnector",
		oci_dbmulticloud.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("gcp_identity_connectivity_status", s.Res.GcpIdentityConnectivityStatus)

	if s.Res.GcpLocation != nil {
		s.D.Set("gcp_location", *s.Res.GcpLocation)
	}

	gcpNodes := []interface{}{}
	for _, item := range s.Res.GcpNodes {
		gcpNodes = append(gcpNodes, GcpNodesToMap(item))
	}
	s.D.Set("gcp_nodes", gcpNodes)

	if s.Res.GcpResourceServiceAgentId != nil {
		s.D.Set("gcp_resource_service_agent_id", *s.Res.GcpResourceServiceAgentId)
	}

	if s.Res.GcpWorkloadIdentityPoolId != nil {
		s.D.Set("gcp_workload_identity_pool_id", *s.Res.GcpWorkloadIdentityPoolId)
	}

	if s.Res.GcpWorkloadIdentityProviderId != nil {
		s.D.Set("gcp_workload_identity_provider_id", *s.Res.GcpWorkloadIdentityProviderId)
	}

	if s.Res.IssuerUrl != nil {
		s.D.Set("issuer_url", *s.Res.IssuerUrl)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
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

func GcpNodesToMap(obj oci_dbmulticloud.GcpNodes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HostId != nil {
		result["host_id"] = string(*obj.HostId)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	result["status"] = string(obj.Status)

	if obj.TimeLastChecked != nil {
		result["time_last_checked"] = obj.TimeLastChecked.String()
	}

	return result
}

func OracleDbGcpIdentityConnectorSummaryToMap(obj oci_dbmulticloud.OracleDbGcpIdentityConnectorSummary) map[string]interface{} {
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

	result["freeform_tags"] = obj.FreeformTags

	result["gcp_identity_connectivity_status"] = string(obj.GcpIdentityConnectivityStatus)

	if obj.GcpLocation != nil {
		result["gcp_location"] = string(*obj.GcpLocation)
	}

	gcpNodes := []interface{}{}
	for _, item := range obj.GcpNodes {
		gcpNodes = append(gcpNodes, GcpNodesToMap(item))
	}
	result["gcp_nodes"] = gcpNodes

	if obj.GcpResourceServiceAgentId != nil {
		result["gcp_resource_service_agent_id"] = string(*obj.GcpResourceServiceAgentId)
	}

	if obj.GcpWorkloadIdentityPoolId != nil {
		result["gcp_workload_identity_pool_id"] = string(*obj.GcpWorkloadIdentityPoolId)
	}

	if obj.GcpWorkloadIdentityProviderId != nil {
		result["gcp_workload_identity_provider_id"] = string(*obj.GcpWorkloadIdentityProviderId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IssuerUrl != nil {
		result["issuer_url"] = string(*obj.IssuerUrl)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
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

func (s *DbmulticloudOracleDbGcpIdentityConnectorResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dbmulticloud.ChangeOracleDbGcpIdentityConnectorCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OracleDbGcpIdentityConnectorId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.ChangeOracleDbGcpIdentityConnectorCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOracleDbGcpIdentityConnectorFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
