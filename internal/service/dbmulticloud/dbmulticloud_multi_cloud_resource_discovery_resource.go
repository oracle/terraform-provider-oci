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

func DbmulticloudMultiCloudResourceDiscoveryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDbmulticloudMultiCloudResourceDiscovery,
		Read:     readDbmulticloudMultiCloudResourceDiscovery,
		Update:   updateDbmulticloudMultiCloudResourceDiscovery,
		Delete:   deleteDbmulticloudMultiCloudResourceDiscovery,
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
			"oracle_db_connector_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			// "oracle_db_azure_vault_id": {
			// 	Type:     schema.TypeString,
			// 	Required: true,
			// },

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
			"resources_filter": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
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
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"properties": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createDbmulticloudMultiCloudResourceDiscovery(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudMultiCloudResourceDiscoveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MultiCloudResourceDiscoveryClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readDbmulticloudMultiCloudResourceDiscovery(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudMultiCloudResourceDiscoveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MultiCloudResourceDiscoveryClient()

	return tfresource.ReadResource(sync)
}

func updateDbmulticloudMultiCloudResourceDiscovery(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudMultiCloudResourceDiscoveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MultiCloudResourceDiscoveryClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDbmulticloudMultiCloudResourceDiscovery(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudMultiCloudResourceDiscoveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MultiCloudResourceDiscoveryClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type DbmulticloudMultiCloudResourceDiscoveryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dbmulticloud.MultiCloudResourceDiscoveryClient
	Res                    *oci_dbmulticloud.MultiCloudResourceDiscovery
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_dbmulticloud.WorkRequestClient
}

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateInProgress),
	}
}

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateSucceeded),
		string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateNeedsAttention),
		string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateFailed),
		string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateCanceled),
	}
}

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateNeedsAttention),
		string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateFailed),
		string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateCanceled),
	}
}

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) Create() error {
	request := oci_dbmulticloud.CreateMultiCloudResourceDiscoveryRequest{}

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

	if oracleDbConnectorId, ok := s.D.GetOkExists("oracle_db_connector_id"); ok {
		tmp := oracleDbConnectorId.(string)
		request.OracleDbConnectorId = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		request.ResourceType = oci_dbmulticloud.MultiCloudResourceDiscoveryResourceTypeEnum(resourceType.(string))
	}

	if resourcesFilter, ok := s.D.GetOkExists("resources_filter"); ok {
		request.ResourcesFilter = tfresource.ObjectMapToStringMap(resourcesFilter.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.CreateMultiCloudResourceDiscovery(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getMultiCloudResourceDiscoveryFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) getMultiCloudResourceDiscoveryFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_dbmulticloud.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	multiCloudResourceDiscoveryId, err := multiCloudResourceDiscoveryWaitForWorkRequest(workId, "multicloudresourcediscovery",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, multiCloudResourceDiscoveryId)
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
	s.D.SetId(*multiCloudResourceDiscoveryId)

	return s.Get()
}

func multiCloudResourceDiscoveryWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func multiCloudResourceDiscoveryWaitForWorkRequest(wId *string, entityType string, action oci_dbmulticloud.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_dbmulticloud.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "dbmulticloud")
	retryPolicy.ShouldRetryOperation = multiCloudResourceDiscoveryWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDbmulticloudMultiCloudResourceDiscoveryWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDbmulticloudMultiCloudResourceDiscoveryWorkRequest(client *oci_dbmulticloud.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_dbmulticloud.ActionTypeEnum) error {
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

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) Get() error {
	request := oci_dbmulticloud.GetMultiCloudResourceDiscoveryRequest{}

	tmp := s.D.Id()
	request.MultiCloudResourceDiscoveryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.GetMultiCloudResourceDiscovery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MultiCloudResourceDiscovery
	return nil
}

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dbmulticloud.UpdateMultiCloudResourceDiscoveryRequest{}

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
	request.MultiCloudResourceDiscoveryId = &tmp

	if oracleDbConnectorId, ok := s.D.GetOkExists("oracle_db_connector_id"); ok {
		tmp := oracleDbConnectorId.(string)
		request.OracleDbConnectorId = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		request.ResourceType = oci_dbmulticloud.MultiCloudResourceDiscoveryResourceTypeEnum(resourceType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.UpdateMultiCloudResourceDiscovery(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMultiCloudResourceDiscoveryFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) Delete() error {
	request := oci_dbmulticloud.DeleteMultiCloudResourceDiscoveryRequest{}

	tmp := s.D.Id()
	request.MultiCloudResourceDiscoveryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.DeleteMultiCloudResourceDiscovery(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := multiCloudResourceDiscoveryWaitForWorkRequest(workId, "multicloudresourcediscovery",
		oci_dbmulticloud.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) SetData() error {
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

	if s.Res.OracleDbConnectorId != nil {
		s.D.Set("oracle_db_connector_id", *s.Res.OracleDbConnectorId)
	}

	s.D.Set("resource_type", s.Res.ResourceType)

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, ResourcesToMap(item))
	}
	s.D.Set("resources", resources)

	s.D.Set("resources_filter", s.Res.ResourcesFilter)

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

func MultiCloudResourceDiscoverySummaryToMap(obj oci_dbmulticloud.MultiCloudResourceDiscoverySummary) map[string]interface{} {
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

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LastModification != nil {
		result["last_modification"] = string(*obj.LastModification)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.OracleDbConnectorId != nil {
		result["oracle_db_connector_id"] = string(*obj.OracleDbConnectorId)
	}

	result["resource_type"] = string(obj.ResourceType)

	resources := []interface{}{}
	for _, item := range obj.Resources {
		resources = append(resources, ResourcesToMap(item))
	}
	result["resources"] = resources

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

func ResourcesToMap(obj oci_dbmulticloud.Resources) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Location != nil {
		result["location"] = string(*obj.Location)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["properties"] = obj.Properties

	if obj.ResourceGroup != nil {
		result["resource_group"] = string(*obj.ResourceGroup)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func (s *DbmulticloudMultiCloudResourceDiscoveryResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dbmulticloud.ChangeMultiCloudResourceDiscoveryCompartmentRequest{}

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
	changeCompartmentRequest.MultiCloudResourceDiscoveryId = &idTmp

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		changeCompartmentRequest.SystemTags = convertedSystemTags
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.ChangeMultiCloudResourceDiscoveryCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMultiCloudResourceDiscoveryFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
