// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudOracleDbAwsIdentityConnectorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDbmulticloudOracleDbAwsIdentityConnectorWithContext,
		ReadContext:   readDbmulticloudOracleDbAwsIdentityConnectorWithContext,
		UpdateContext: updateDbmulticloudOracleDbAwsIdentityConnectorWithContext,
		DeleteContext: deleteDbmulticloudOracleDbAwsIdentityConnectorWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"aws_location": {
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
			"issuer_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"oidc_scope": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_role_details": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"role_arn": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service_private_endpoint": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"assume_role_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_nodes": {
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
					},
				},
			},

			// Optional
			"aws_account_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aws_sts_private_endpoint": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
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

func createDbmulticloudOracleDbAwsIdentityConnectorWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsIdentityConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDbmulticloudOracleDbAwsIdentityConnectorWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsIdentityConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDbmulticloudOracleDbAwsIdentityConnectorWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsIdentityConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDbmulticloudOracleDbAwsIdentityConnectorWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsIdentityConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DbmulticloudOracleDbAwsIdentityConnectorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dbmulticloud.DbMulticloudAwsProviderClient
	Res                    *oci_dbmulticloud.OracleDbAwsIdentityConnector
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_dbmulticloud.WorkRequestClient
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsIdentityConnectorLifecycleStateCreating),
	}
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsIdentityConnectorLifecycleStateActive),
	}
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsIdentityConnectorLifecycleStateDeleting),
	}
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsIdentityConnectorLifecycleStateDeleted),
	}
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.CreateOracleDbAwsIdentityConnectorRequest{}

	if awsAccountId, ok := s.D.GetOkExists("aws_account_id"); ok {
		tmp := awsAccountId.(string)
		request.AwsAccountId = &tmp
	}

	if awsLocation, ok := s.D.GetOkExists("aws_location"); ok {
		tmp := awsLocation.(string)
		request.AwsLocation = &tmp
	}

	if awsStsPrivateEndpoint, ok := s.D.GetOkExists("aws_sts_private_endpoint"); ok {
		tmp := awsStsPrivateEndpoint.(string)
		request.AwsStsPrivateEndpoint = &tmp
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

	if issuerUrl, ok := s.D.GetOkExists("issuer_url"); ok {
		tmp := issuerUrl.(string)
		request.IssuerUrl = &tmp
	}

	if oidcScope, ok := s.D.GetOkExists("oidc_scope"); ok {
		tmp := oidcScope.(string)
		request.OidcScope = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if serviceRoleDetails, ok := s.D.GetOkExists("service_role_details"); ok {
		interfaces := serviceRoleDetails.([]interface{})
		tmp := make([]oci_dbmulticloud.ServiceRoleDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "service_role_details", stateDataIndex)
			converted, err := s.mapToServiceRoleDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("service_role_details") {
			request.ServiceRoleDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.CreateOracleDbAwsIdentityConnector(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getOracleDbAwsIdentityConnectorFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) getOracleDbAwsIdentityConnectorFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_dbmulticloud.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	oracleDbAwsIdentityConnectorId, err := oracleDbAwsIdentityConnectorWaitForWorkRequest(ctx, workId, "oracledbawsconnector",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, oracleDbAwsIdentityConnectorId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(ctx,
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
	s.D.SetId(*oracleDbAwsIdentityConnectorId)

	return s.GetWithContext(ctx)
}

func oracleDbAwsIdentityConnectorWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func oracleDbAwsIdentityConnectorWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_dbmulticloud.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_dbmulticloud.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "dbmulticloud")
	retryPolicy.ShouldRetryOperation = oracleDbAwsIdentityConnectorWorkRequestShouldRetryFunc(timeout)

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
			response, err = client.GetWorkRequest(ctx,
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
		return nil, getErrorFromDbmulticloudOracleDbAwsIdentityConnectorWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDbmulticloudOracleDbAwsIdentityConnectorWorkRequest(ctx context.Context, client *oci_dbmulticloud.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_dbmulticloud.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
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

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.GetOracleDbAwsIdentityConnectorRequest{}

	tmp := s.D.Id()
	request.OracleDbAwsIdentityConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.GetOracleDbAwsIdentityConnector(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.OracleDbAwsIdentityConnector
	return nil
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dbmulticloud.UpdateOracleDbAwsIdentityConnectorRequest{}

	if awsAccountId, ok := s.D.GetOkExists("aws_account_id"); ok {
		tmp := awsAccountId.(string)
		request.AwsAccountId = &tmp
	}

	if awsLocation, ok := s.D.GetOkExists("aws_location"); ok {
		tmp := awsLocation.(string)
		request.AwsLocation = &tmp
	}

	if awsStsPrivateEndpoint, ok := s.D.GetOkExists("aws_sts_private_endpoint"); ok {
		tmp := awsStsPrivateEndpoint.(string)
		request.AwsStsPrivateEndpoint = &tmp
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

	if issuerUrl, ok := s.D.GetOkExists("issuer_url"); ok {
		tmp := issuerUrl.(string)
		request.IssuerUrl = &tmp
	}

	if oidcScope, ok := s.D.GetOkExists("oidc_scope"); ok {
		tmp := oidcScope.(string)
		request.OidcScope = &tmp
	}

	tmp := s.D.Id()
	request.OracleDbAwsIdentityConnectorId = &tmp

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if serviceRoleDetails, ok := s.D.GetOkExists("service_role_details"); ok {
		interfaces := serviceRoleDetails.([]interface{})
		tmp := make([]oci_dbmulticloud.ServiceRoleDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "service_role_details", stateDataIndex)
			converted, err := s.mapToServiceRoleDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("service_role_details") {
			request.ServiceRoleDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.UpdateOracleDbAwsIdentityConnector(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOracleDbAwsIdentityConnectorFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.DeleteOracleDbAwsIdentityConnectorRequest{}

	tmp := s.D.Id()
	request.OracleDbAwsIdentityConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.DeleteOracleDbAwsIdentityConnector(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := oracleDbAwsIdentityConnectorWaitForWorkRequest(ctx, workId, "oracledbawsconnector",
		oci_dbmulticloud.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) SetData() error {
	if s.Res.AwsAccountId != nil {
		s.D.Set("aws_account_id", *s.Res.AwsAccountId)
	}

	if s.Res.AwsLocation != nil {
		s.D.Set("aws_location", *s.Res.AwsLocation)
	}

	if s.Res.AwsStsPrivateEndpoint != nil {
		s.D.Set("aws_sts_private_endpoint", *s.Res.AwsStsPrivateEndpoint)
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

	if s.Res.IssuerUrl != nil {
		s.D.Set("issuer_url", *s.Res.IssuerUrl)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.OidcScope != nil {
		s.D.Set("oidc_scope", *s.Res.OidcScope)
	}

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	serviceRoleDetails := []interface{}{}
	for _, item := range s.Res.ServiceRoleDetails {
		serviceRoleDetails = append(serviceRoleDetails, ServiceRoleDetailsToMap(item))
	}
	s.D.Set("service_role_details", serviceRoleDetails)

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

func AwsNodesToMap(obj oci_dbmulticloud.AwsNodes) map[string]interface{} {
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

func OracleDbAwsIdentityConnectorSummaryToMap(obj oci_dbmulticloud.OracleDbAwsIdentityConnectorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AwsAccountId != nil {
		result["aws_account_id"] = string(*obj.AwsAccountId)
	}

	if obj.AwsLocation != nil {
		result["aws_location"] = string(*obj.AwsLocation)
	}

	if obj.AwsStsPrivateEndpoint != nil {
		result["aws_sts_private_endpoint"] = string(*obj.AwsStsPrivateEndpoint)
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

	if obj.IssuerUrl != nil {
		result["issuer_url"] = string(*obj.IssuerUrl)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.OidcScope != nil {
		result["oidc_scope"] = string(*obj.OidcScope)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	serviceRoleDetails := []interface{}{}
	for _, item := range obj.ServiceRoleDetails {
		serviceRoleDetails = append(serviceRoleDetails, ServiceRoleDetailsToMap(item))
	}
	result["service_role_details"] = serviceRoleDetails

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

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) mapToServiceRoleDetail(fieldKeyFormat string) (oci_dbmulticloud.ServiceRoleDetail, error) {
	result := oci_dbmulticloud.ServiceRoleDetail{}

	if roleArn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role_arn")); ok {
		tmp := roleArn.(string)
		result.RoleArn = &tmp
	}

	if servicePrivateEndpoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_private_endpoint")); ok {
		tmp := servicePrivateEndpoint.(string)
		result.ServicePrivateEndpoint = &tmp
	}

	if serviceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_type")); ok {
		result.ServiceType = oci_dbmulticloud.ServiceRoleDetailServiceTypeEnum(serviceType.(string))
	}

	return result, nil
}

func ServiceRoleDetailsToMap(obj oci_dbmulticloud.ServiceRoleDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["assume_role_status"] = string(obj.AssumeRoleStatus)

	awsNodes := []interface{}{}
	for _, item := range obj.AwsNodes {
		awsNodes = append(awsNodes, AwsNodesToMap(item))
	}
	result["aws_nodes"] = awsNodes

	if obj.RoleArn != nil {
		result["role_arn"] = string(*obj.RoleArn)
	}

	if obj.ServicePrivateEndpoint != nil {
		result["service_private_endpoint"] = string(*obj.ServicePrivateEndpoint)
	}

	result["service_type"] = string(obj.ServiceType)

	return result
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_dbmulticloud.ChangeOracleDbAwsIdentityConnectorCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OracleDbAwsIdentityConnectorId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.ChangeOracleDbAwsIdentityConnectorCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOracleDbAwsIdentityConnectorFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
