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

func DbmulticloudOracleDbAzureBlobContainerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDbmulticloudOracleDbAzureBlobContainer,
		Read:     readDbmulticloudOracleDbAzureBlobContainer,
		Update:   updateDbmulticloudOracleDbAzureBlobContainer,
		Delete:   deleteDbmulticloudOracleDbAzureBlobContainer,
		Schema: map[string]*schema.Schema{
			// Required
			"azure_storage_account_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"azure_storage_container_name": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"private_endpoint_dns_alias": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_endpoint_ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"last_modification": {
				Type:     schema.TypeString,
				Computed: true,
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

func createDbmulticloudOracleDbAzureBlobContainer(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureBlobContainerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureBlobContainerClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readDbmulticloudOracleDbAzureBlobContainer(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureBlobContainerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureBlobContainerClient()

	return tfresource.ReadResource(sync)
}

func updateDbmulticloudOracleDbAzureBlobContainer(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureBlobContainerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureBlobContainerClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDbmulticloudOracleDbAzureBlobContainer(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureBlobContainerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureBlobContainerClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type DbmulticloudOracleDbAzureBlobContainerResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dbmulticloud.OracleDBAzureBlobContainerClient
	Res                    *oci_dbmulticloud.OracleDbAzureBlobContainer
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_dbmulticloud.WorkRequestClient
}

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAzureBlobContainerLifecycleStateCreating),
	}
}

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAzureBlobContainerLifecycleStateActive),
	}
}

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAzureBlobContainerLifecycleStateDeleting),
	}
}

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAzureBlobContainerLifecycleStateDeleted),
	}
}

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) Create() error {
	request := oci_dbmulticloud.CreateOracleDbAzureBlobContainerRequest{}

	if azureStorageAccountName, ok := s.D.GetOkExists("azure_storage_account_name"); ok {
		tmp := azureStorageAccountName.(string)
		request.AzureStorageAccountName = &tmp
	}

	if azureStorageContainerName, ok := s.D.GetOkExists("azure_storage_container_name"); ok {
		tmp := azureStorageContainerName.(string)
		request.AzureStorageContainerName = &tmp
	}

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

	if privateEndpointDnsAlias, ok := s.D.GetOkExists("private_endpoint_dns_alias"); ok {
		tmp := privateEndpointDnsAlias.(string)
		request.PrivateEndpointDnsAlias = &tmp
	}

	if privateEndpointIpAddress, ok := s.D.GetOkExists("private_endpoint_ip_address"); ok {
		tmp := privateEndpointIpAddress.(string)
		request.PrivateEndpointIpAddress = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.CreateOracleDbAzureBlobContainer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getOracleDbAzureBlobContainerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) getOracleDbAzureBlobContainerFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_dbmulticloud.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	oracleDbAzureBlobContainerId, err := oracleDbAzureBlobContainerWaitForWorkRequest(workId, "oracledbazureblobcontainer",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, oracleDbAzureBlobContainerId)
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
	s.D.SetId(*oracleDbAzureBlobContainerId)

	return s.Get()
}

func oracleDbAzureBlobContainerWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func oracleDbAzureBlobContainerWaitForWorkRequest(wId *string, entityType string, action oci_dbmulticloud.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_dbmulticloud.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "dbmulticloud")
	retryPolicy.ShouldRetryOperation = oracleDbAzureBlobContainerWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDbmulticloudOracleDbAzureBlobContainerWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDbmulticloudOracleDbAzureBlobContainerWorkRequest(client *oci_dbmulticloud.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_dbmulticloud.ActionTypeEnum) error {
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

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) Get() error {
	request := oci_dbmulticloud.GetOracleDbAzureBlobContainerRequest{}

	tmp := s.D.Id()
	request.OracleDbAzureBlobContainerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.GetOracleDbAzureBlobContainer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OracleDbAzureBlobContainer
	return nil
}

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dbmulticloud.UpdateOracleDbAzureBlobContainerRequest{}

	if azureStorageAccountName, ok := s.D.GetOkExists("azure_storage_account_name"); ok {
		tmp := azureStorageAccountName.(string)
		request.AzureStorageAccountName = &tmp
	}

	if azureStorageContainerName, ok := s.D.GetOkExists("azure_storage_container_name"); ok {
		tmp := azureStorageContainerName.(string)
		request.AzureStorageContainerName = &tmp
	}

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

	tmp := s.D.Id()
	request.OracleDbAzureBlobContainerId = &tmp

	if privateEndpointDnsAlias, ok := s.D.GetOkExists("private_endpoint_dns_alias"); ok {
		tmp := privateEndpointDnsAlias.(string)
		request.PrivateEndpointDnsAlias = &tmp
	}

	if privateEndpointIpAddress, ok := s.D.GetOkExists("private_endpoint_ip_address"); ok {
		tmp := privateEndpointIpAddress.(string)
		request.PrivateEndpointIpAddress = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.UpdateOracleDbAzureBlobContainer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOracleDbAzureBlobContainerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) Delete() error {
	request := oci_dbmulticloud.DeleteOracleDbAzureBlobContainerRequest{}

	tmp := s.D.Id()
	request.OracleDbAzureBlobContainerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.DeleteOracleDbAzureBlobContainer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := oracleDbAzureBlobContainerWaitForWorkRequest(workId, "oracledbazureblobcontainer",
		oci_dbmulticloud.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) SetData() error {
	if s.Res.AzureStorageAccountName != nil {
		s.D.Set("azure_storage_account_name", *s.Res.AzureStorageAccountName)
	}

	if s.Res.AzureStorageContainerName != nil {
		s.D.Set("azure_storage_container_name", *s.Res.AzureStorageContainerName)
	}

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

	if s.Res.LastModification != nil {
		s.D.Set("last_modification", *s.Res.LastModification)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.PrivateEndpointDnsAlias != nil {
		s.D.Set("private_endpoint_dns_alias", *s.Res.PrivateEndpointDnsAlias)
	}

	if s.Res.PrivateEndpointIpAddress != nil {
		s.D.Set("private_endpoint_ip_address", *s.Res.PrivateEndpointIpAddress)
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

func OracleDbAzureBlobContainerSummaryToMap(obj oci_dbmulticloud.OracleDbAzureBlobContainerSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AzureStorageAccountName != nil {
		result["azure_storage_account_name"] = string(*obj.AzureStorageAccountName)
	}

	if obj.AzureStorageContainerName != nil {
		result["azure_storage_container_name"] = string(*obj.AzureStorageContainerName)
	}

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

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LastModification != nil {
		result["last_modification"] = string(*obj.LastModification)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.PrivateEndpointDnsAlias != nil {
		result["private_endpoint_dns_alias"] = string(*obj.PrivateEndpointDnsAlias)
	}

	if obj.PrivateEndpointIpAddress != nil {
		result["private_endpoint_ip_address"] = string(*obj.PrivateEndpointIpAddress)
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

func (s *DbmulticloudOracleDbAzureBlobContainerResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dbmulticloud.ChangeOracleDbAzureBlobContainerCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		changeCompartmentRequest.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		changeCompartmentRequest.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	idTmp := s.D.Id()
	changeCompartmentRequest.OracleDbAzureBlobContainerId = &idTmp

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		changeCompartmentRequest.SystemTags = convertedSystemTags
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.ChangeOracleDbAzureBlobContainerCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOracleDbAzureBlobContainerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
