// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package gdp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_gdp "github.com/oracle/oci-go-sdk/v65/gdp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var gdpPipelineEntityTypes = []string{"gdppipeline", "cdspipeline"}
var commercialSubdomain = "prod.cp.cdsaas"
var gdpUSGovCode = "USGOV"
var gdpCommercialCode = "COMMERCIAL"

func GdpGdpPipelineResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createGdpGdpPipelineWithContext,
		ReadContext:   readGdpGdpPipelineWithContext,
		UpdateContext: updateGdpGdpPipelineWithContext,
		DeleteContext: deleteGdpGdpPipelineWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"bucket_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"bucket_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"peering_region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pipeline_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"approval_key_vault_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization_details": {
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
			"env": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  gdpCommercialCode,
			},
			"file_types": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_approval_needed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_chunking_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_file_override_in_destination_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_scanning_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"service_log_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_gdp.GdpPipelineLifecycleStateInactive),
					string(oci_gdp.GdpPipelineLifecycleStateActive),
				}, true),
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peered_gdp_pipeline_id": {
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

func createGdpGdpPipelineWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GdpGdpPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GuardedDataPipelineClient()

	if env, ok := sync.D.GetOk("env"); !ok || env.(string) != gdpUSGovCode {
		currentHost := sync.Client.Host
		newHost := strings.Replace(currentHost, "gdp", commercialSubdomain, 1)
		sync.Client.Host = newHost
	}

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	return nil

}

func readGdpGdpPipelineWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GdpGdpPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GuardedDataPipelineClient()

	if env, ok := sync.D.GetOk("env"); !ok || env.(string) != gdpUSGovCode {
		currentHost := sync.Client.Host
		newHost := strings.Replace(currentHost, "gdp", commercialSubdomain, 1)
		sync.Client.Host = newHost
	}

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateGdpGdpPipelineWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GdpGdpPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GuardedDataPipelineClient()

	if env, ok := sync.D.GetOk("env"); !ok || env.(string) != gdpUSGovCode {
		currentHost := sync.Client.Host
		newHost := strings.Replace(currentHost, "gdp", commercialSubdomain, 1)
		sync.Client.Host = newHost
	}

	//powerOn, powerOff := false, false

	// We dont provide state change from terraform
	//if sync.D.HasChange("state") {
	//	wantedState := strings.ToUpper(sync.D.Get("state").(string))
	//	if oci_gdp.GdpPipelineLifecycleStateActive == oci_gdp.GdpPipelineLifecycleStateEnum(wantedState) {
	//		powerOn = true
	//	} else if oci_gdp.GdpPipelineLifecycleStateInactive == oci_gdp.GdpPipelineLifecycleStateEnum(wantedState) {
	//		powerOff = true
	//	}
	//}
	//
	//if powerOn {
	//	if err := sync.StartGdpPipeline(ctx); err != nil {
	//		return tfresource.HandleDiagError(m, err)
	//	}
	//	sync.D.Set("state", oci_gdp.GdpPipelineLifecycleStateActive)
	//}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	//if powerOff {
	//	if err := sync.StopGdpPipeline(ctx); err != nil {
	//		return tfresource.HandleDiagError(m, err)
	//	}
	//	sync.D.Set("state", oci_gdp.GdpPipelineLifecycleStateInactive)
	//}

	return nil
}

func deleteGdpGdpPipelineWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GdpGdpPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GuardedDataPipelineClient()
	if env, ok := sync.D.GetOk("env"); !ok || env.(string) != gdpUSGovCode {
		currentHost := sync.Client.Host
		newHost := strings.Replace(currentHost, "gdp", commercialSubdomain, 1)
		sync.Client.Host = newHost
	}

	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type GdpGdpPipelineResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_gdp.GuardedDataPipelineClient
	Res                    *oci_gdp.GdpPipeline
	DisableNotFoundRetries bool
}

func (s *GdpGdpPipelineResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GdpGdpPipelineResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_gdp.GdpPipelineLifecycleStateCreating),
	}
}

func (s *GdpGdpPipelineResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_gdp.GdpPipelineLifecycleStateInactive),
		string(oci_gdp.GdpPipelineLifecycleStateNeedsAttention),
	}
}

func (s *GdpGdpPipelineResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_gdp.GdpPipelineLifecycleStateDeleting),
	}
}

func (s *GdpGdpPipelineResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_gdp.GdpPipelineLifecycleStateDeleted),
	}
}

func (s *GdpGdpPipelineResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_gdp.CreateGdpPipelineRequest{}

	if approvalKeyVaultId, ok := s.D.GetOkExists("approval_key_vault_id"); ok {
		tmp := approvalKeyVaultId.(string)
		request.ApprovalKeyVaultId = &tmp
	}

	if authorizationDetails, ok := s.D.GetOkExists("authorization_details"); ok {
		tmp := authorizationDetails.(string)
		request.AuthorizationDetails = &tmp
	}

	if bucketDetails, ok := s.D.GetOkExists("bucket_details"); ok {
		interfaces := bucketDetails.([]interface{})
		tmp := make([]oci_gdp.BucketDetailsDefinition, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "bucket_details", stateDataIndex)
			converted, err := s.mapToBucketDetailsDefinition(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("bucket_details") {
			request.BucketDetails = tmp
		}
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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fileTypes, ok := s.D.GetOkExists("file_types"); ok {
		interfaces := fileTypes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("file_types") {
			request.FileTypes = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isApprovalNeeded, ok := s.D.GetOkExists("is_approval_needed"); ok {
		tmp := isApprovalNeeded.(bool)
		request.IsApprovalNeeded = &tmp
	}

	if isChunkingEnabled, ok := s.D.GetOkExists("is_chunking_enabled"); ok {
		tmp := isChunkingEnabled.(bool)
		request.IsChunkingEnabled = &tmp
	}

	if isFileOverrideInDestinationEnabled, ok := s.D.GetOkExists("is_file_override_in_destination_enabled"); ok {
		tmp := isFileOverrideInDestinationEnabled.(bool)
		request.IsFileOverrideInDestinationEnabled = &tmp
	}

	if isScanningEnabled, ok := s.D.GetOkExists("is_scanning_enabled"); ok {
		tmp := isScanningEnabled.(bool)
		request.IsScanningEnabled = &tmp
	}

	if peeringRegion, ok := s.D.GetOkExists("peering_region"); ok {
		tmp := peeringRegion.(string)
		request.PeeringRegion = &tmp
	}

	if pipelineType, ok := s.D.GetOkExists("pipeline_type"); ok {
		request.PipelineType = oci_gdp.GdpPipelinePipelineTypeEnum(pipelineType.(string))
	}

	if serviceLogGroupId, ok := s.D.GetOkExists("service_log_group_id"); ok {
		tmp := serviceLogGroupId.(string)
		request.ServiceLogGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp")

	response, err := s.Client.CreateGdpPipeline(ctx, request)
	if err != nil {
		return err
	}

	for key, value := range response.RawResponse.Header {
		if strings.ToLower(key) == "peer-validation-nonce" {
			fmt.Printf("%s for GDP pipeline is: %s\n", key, value)
		}
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_gdp.GetGdpWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetGdpWorkRequest(ctx,
		oci_gdp.GetGdpWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && contains(gdpPipelineEntityTypes, *res.EntityType) && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getGdpPipelineFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp"), oci_gdp.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GdpGdpPipelineResourceCrud) getGdpPipelineFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_gdp.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	gdpPipelineId, err := gdpPipelineWaitForWorkRequest(ctx, workId, actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*gdpPipelineId)

	return s.GetWithContext(ctx)
}

func gdpPipelineWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "gdp", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_gdp.GetGdpWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func gdpPipelineWaitForWorkRequest(ctx context.Context, wId *string, action oci_gdp.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_gdp.GuardedDataPipelineClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "gdp")
	retryPolicy.ShouldRetryOperation = gdpPipelineWorkRequestShouldRetryFunc(timeout)

	response := oci_gdp.GetGdpWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_gdp.OperationStatusInProgress),
			string(oci_gdp.OperationStatusAccepted),
			string(oci_gdp.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_gdp.OperationStatusSucceeded),
			string(oci_gdp.OperationStatusFailed),
			string(oci_gdp.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetGdpWorkRequest(ctx,
				oci_gdp.GetGdpWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.GdpWorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	entityType := ""

	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if contains(gdpPipelineEntityTypes, *res.EntityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				entityType = *res.EntityType
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_gdp.OperationStatusFailed || response.Status == oci_gdp.OperationStatusCanceled {
		return nil, getErrorFromGdpGdpPipelineWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGdpGdpPipelineWorkRequest(ctx context.Context, client *oci_gdp.GuardedDataPipelineClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_gdp.ActionTypeEnum) error {
	response, err := client.ListGdpWorkRequestErrors(ctx,
		oci_gdp.ListGdpWorkRequestErrorsRequest{
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

func (s *GdpGdpPipelineResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_gdp.GetGdpPipelineRequest{}

	tmp := s.D.Id()
	request.GdpPipelineId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp")

	response, err := s.Client.GetGdpPipeline(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.GdpPipeline
	return nil
}

func (s *GdpGdpPipelineResourceCrud) UpdateWithContext(ctx context.Context) error {

	// Not support peering via terraform
	//if _, ok := s.D.GetOkExists("peerValidationNonce"); ok && s.D.HasChange("peerValidationNonce") {
	//	err := s.PeerGdpPipeline()
	//	if err != nil {
	//		return err
	//	}
	//}

	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_gdp.UpdateGdpPipelineRequest{}

	if approvalKeyVaultId, ok := s.D.GetOkExists("approval_key_vault_id"); ok {
		tmp := approvalKeyVaultId.(string)
		request.ApprovalKeyVaultId = &tmp
	}

	if authorizationDetails, ok := s.D.GetOkExists("authorization_details"); ok {
		tmp := authorizationDetails.(string)
		request.AuthorizationDetails = &tmp
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

	if fileTypes, ok := s.D.GetOkExists("file_types"); ok {
		interfaces := fileTypes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("file_types") {
			request.FileTypes = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.GdpPipelineId = &tmp

	if isApprovalNeeded, ok := s.D.GetOkExists("is_approval_needed"); ok {
		tmp := isApprovalNeeded.(bool)
		request.IsApprovalNeeded = &tmp
	}

	if isChunkingEnabled, ok := s.D.GetOkExists("is_chunking_enabled"); ok {
		tmp := isChunkingEnabled.(bool)
		request.IsChunkingEnabled = &tmp
	}

	if isFileOverrideInDestinationEnabled, ok := s.D.GetOkExists("is_file_override_in_destination_enabled"); ok {
		tmp := isFileOverrideInDestinationEnabled.(bool)
		request.IsFileOverrideInDestinationEnabled = &tmp
	}

	if isScanningEnabled, ok := s.D.GetOkExists("is_scanning_enabled"); ok {
		tmp := isScanningEnabled.(bool)
		request.IsScanningEnabled = &tmp
	}

	if serviceLogGroupId, ok := s.D.GetOkExists("service_log_group_id"); ok {
		tmp := serviceLogGroupId.(string)
		request.ServiceLogGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp")

	response, err := s.Client.UpdateGdpPipeline(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getGdpPipelineFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp"), oci_gdp.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GdpGdpPipelineResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_gdp.DeleteGdpPipelineRequest{}

	tmp := s.D.Id()
	request.GdpPipelineId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp")

	response, err := s.Client.DeleteGdpPipeline(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := gdpPipelineWaitForWorkRequest(ctx, workId, oci_gdp.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GdpGdpPipelineResourceCrud) SetData() error {
	if s.Res.ApprovalKeyVaultId != nil {
		s.D.Set("approval_key_vault_id", *s.Res.ApprovalKeyVaultId)
	}

	if s.Res.AuthorizationDetails != nil {
		s.D.Set("authorization_details", *s.Res.AuthorizationDetails)
	}

	bucketDetails := []interface{}{}
	for _, item := range s.Res.BucketDetails {
		bucketDetails = append(bucketDetails, BucketDetailsDefinitionToMap(item))
	}
	s.D.Set("bucket_details", bucketDetails)

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

	s.D.Set("file_types", s.Res.FileTypes)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsApprovalNeeded != nil {
		s.D.Set("is_approval_needed", *s.Res.IsApprovalNeeded)
	}

	if s.Res.IsChunkingEnabled != nil {
		s.D.Set("is_chunking_enabled", *s.Res.IsChunkingEnabled)
	}

	if s.Res.IsFileOverrideInDestinationEnabled != nil {
		s.D.Set("is_file_override_in_destination_enabled", *s.Res.IsFileOverrideInDestinationEnabled)
	}

	if s.Res.IsScanningEnabled != nil {
		s.D.Set("is_scanning_enabled", *s.Res.IsScanningEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PeeredGdpPipelineId != nil {
		s.D.Set("peered_gdp_pipeline_id", *s.Res.PeeredGdpPipelineId)
	}

	if s.Res.PeeringRegion != nil {
		s.D.Set("peering_region", *s.Res.PeeringRegion)
	}

	s.D.Set("pipeline_type", s.Res.PipelineType)

	if s.Res.ServiceLogGroupId != nil {
		s.D.Set("service_log_group_id", *s.Res.ServiceLogGroupId)
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

func (s *GdpGdpPipelineResourceCrud) StartGdpPipeline(ctx context.Context) error {
	request := oci_gdp.StartGdpPipelineRequest{}

	idTmp := s.D.Id()
	request.GdpPipelineId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp")

	_, err := s.Client.StartGdpPipeline(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_gdp.GdpPipelineLifecycleStateActive }
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GdpGdpPipelineResourceCrud) StopGdpPipeline(ctx context.Context) error {
	request := oci_gdp.StopGdpPipelineRequest{}

	idTmp := s.D.Id()
	request.GdpPipelineId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp")

	_, err := s.Client.StopGdpPipeline(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_gdp.GdpPipelineLifecycleStateInactive }
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

//func (s *GdpGdpPipelineResourceCrud) PeerGdpPipeline() error {
//	request := oci_gdp.PeerGdpPipelineRequest{}
//
//	idTmp := s.D.Id()
//	request.GdpPipelineId = &idTmp
//
//	if peerValidationNonce, ok := s.D.GetOkExists("peer_validation_nonce"); ok {
//		tmp := peerValidationNonce.(string)
//		request.PeerValidationNonce = &tmp
//	}
//
//	if peeringGdpPipelineId, ok := s.D.GetOkExists("peering_gdp_pipeline_id"); ok {
//		tmp := peeringGdpPipelineId.(string)
//		request.PeeringGdpPipelineId = &tmp
//	}
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp")
//
//	_, err := s.Client.PeerGdpPipeline(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	if waitErr := tfresource.WaitForUpdatedStateWithContext(s.D, s); waitErr != nil {
//		return waitErr
//	}
//
//	return nil
//}

func (s *GdpGdpPipelineResourceCrud) mapToBucketDetailsDefinition(fieldKeyFormat string) (oci_gdp.BucketDetailsDefinition, error) {
	result := oci_gdp.BucketDetailsDefinition{}

	if bucketType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket_type")); ok {
		result.BucketType = oci_gdp.BucketDetailsDefinitionBucketTypeEnum(bucketType.(string))
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	return result, nil
}

func BucketDetailsDefinitionToMap(obj oci_gdp.BucketDetailsDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	result["bucket_type"] = string(obj.BucketType)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	return result
}

func GdpPipelineSummaryToMap(obj oci_gdp.GdpPipelineSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuthorizationDetails != nil {
		result["authorization_details"] = string(*obj.AuthorizationDetails)
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

	if obj.IsApprovalNeeded != nil {
		result["is_approval_needed"] = bool(*obj.IsApprovalNeeded)
	}

	if obj.IsChunkingEnabled != nil {
		result["is_chunking_enabled"] = bool(*obj.IsChunkingEnabled)
	}

	if obj.IsFileOverrideInDestinationEnabled != nil {
		result["is_file_override_in_destination_enabled"] = bool(*obj.IsFileOverrideInDestinationEnabled)
	}

	if obj.IsScanningEnabled != nil {
		result["is_scanning_enabled"] = bool(*obj.IsScanningEnabled)
	}

	if obj.PeeringRegion != nil {
		result["peering_region"] = string(*obj.PeeringRegion)
	}

	result["pipeline_type"] = string(obj.PipelineType)

	if obj.ServiceLogGroupId != nil {
		result["service_log_group_id"] = string(*obj.ServiceLogGroupId)
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

func (s *GdpGdpPipelineResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_gdp.ChangeGdpPipelineCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.GdpPipelineId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp")

	response, err := s.Client.ChangeGdpPipelineCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getGdpPipelineFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "gdp"), oci_gdp.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func contains(s []string, str string) bool {
	fmt.Printf("EntityType is %s\n", str)
	for _, v := range s {
		if v == strings.ToLower(str) {
			return true
		}
	}
	return false
}
