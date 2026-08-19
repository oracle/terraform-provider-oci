// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiHostedDeploymentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createGenerativeAiHostedDeploymentWithContext,
		ReadContext:   readGenerativeAiHostedDeploymentWithContext,
		UpdateContext: updateGenerativeAiHostedDeploymentWithContext,
		DeleteContext: deleteGenerativeAiHostedDeploymentWithContext,
		Schema: map[string]*schema.Schema{
			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Required
			"active_artifact": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"artifact_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"SIMPLE_DOCKER_ARTIFACT",
							}, true),
						},
						"container_uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hosted_deployment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_vulnerability_scan_required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_created": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
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
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"hosted_application_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Computed
			"artifacts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"artifact_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"container_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hosted_deployment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_vulnerability_scan_required": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
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

func createGenerativeAiHostedDeploymentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readGenerativeAiHostedDeploymentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateGenerativeAiHostedDeploymentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteGenerativeAiHostedDeploymentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type GenerativeAiHostedDeploymentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai.GenerativeAiClient
	Res                    *oci_generative_ai.HostedDeployment
	DisableNotFoundRetries bool
}

func (s *GenerativeAiHostedDeploymentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiHostedDeploymentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai.HostedDeploymentLifecycleStateCreating),
	}
}

func (s *GenerativeAiHostedDeploymentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai.HostedDeploymentLifecycleStateNeedsAttention),
		string(oci_generative_ai.HostedDeploymentLifecycleStateActive),
	}
}

func (s *GenerativeAiHostedDeploymentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai.HostedDeploymentLifecycleStateDeleting),
	}
}

func (s *GenerativeAiHostedDeploymentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai.HostedDeploymentLifecycleStateDeleted),
	}
}

func (s *GenerativeAiHostedDeploymentResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_generative_ai.CreateHostedDeploymentRequest{}

	if activeArtifact, ok := s.D.GetOkExists("active_artifact"); ok {
		if tmpList := activeArtifact.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "active_artifact", 0)
			tmp, err := s.mapToSingleDockerArtifact(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ActiveArtifact = tmp
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostedApplicationId, ok := s.D.GetOkExists("hosted_application_id"); ok {
		tmp := hostedApplicationId.(string)
		request.HostedApplicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.CreateHostedDeployment(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getHostedDeploymentFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiHostedDeploymentResourceCrud) getHostedDeploymentFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai.ActionTypeEnum, timeout time.Duration) error {
	if workId == nil || *workId == "" {
		// Some create and update operations complete synchronously without an
		// opc-work-request-id. Refresh the deployment already stored in state
		// instead of issuing an invalid GetWorkRequest call with a nil ID.
		return s.GetWithContext(ctx)
	}

	// Wait until it finishes
	hostedDeploymentId, err := hostedDeploymentWaitForWorkRequest(ctx, workId, "hosteddeployment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*hostedDeploymentId)

	return s.GetWithContext(ctx)
}

func hostedDeploymentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func hostedDeploymentWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_generative_ai.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai.GenerativeAiClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai")
	retryPolicy.ShouldRetryOperation = hostedDeploymentWorkRequestShouldRetryFunc(timeout)

	response := oci_generative_ai.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
			response, err = client.GetWorkRequest(ctx,
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
	if _, e := stateConf.WaitForStateContext(ctx); e != nil {
		return nil, e
	}

	var identifier *string
	normalizedExpectedEntityType := strings.ReplaceAll(strings.ToLower(entityType), "_", "")
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if res.EntityType == nil {
			continue
		}
		normalizedEntityType := strings.ReplaceAll(strings.ToLower(*res.EntityType), "_", "")
		if strings.Contains(normalizedEntityType, normalizedExpectedEntityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_generative_ai.OperationStatusFailed || response.Status == oci_generative_ai.OperationStatusCanceled {
		return nil, getErrorFromGenerativeAiHostedDeploymentWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiHostedDeploymentWorkRequest(ctx context.Context, client *oci_generative_ai.GenerativeAiClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
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

func (s *GenerativeAiHostedDeploymentResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.GetHostedDeploymentRequest{}

	tmp := s.D.Id()
	request.HostedDeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.GetHostedDeployment(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.HostedDeployment
	return nil
}

func (s *GenerativeAiHostedDeploymentResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_generative_ai.UpdateHostedDeploymentRequest{}

	if activeArtifact, ok := s.D.GetOkExists("active_artifact"); ok {
		if tmpList := activeArtifact.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "active_artifact", 0)
			tmp, err := s.mapToSingleDockerArtifact(fieldKeyFormat)
			if err != nil {
				return err
			}
			if s.D.HasChange("active_artifact") {
				artifact, err := s.mapToCreateArtifactDetails(fieldKeyFormat)
				if err != nil {
					return err
				}

				idTmp := s.D.Id()
				addArtifactRequest := oci_generative_ai.AddArtifactRequest{}
				addArtifactRequest.HostedDeploymentId = &idTmp
				addArtifactRequest.Artifact = artifact
				addArtifactRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

				addArtifactResponse, err := s.Client.AddArtifact(ctx, addArtifactRequest)
				if err != nil {
					return err
				}

				// AddArtifact is asynchronous. Wait until the new artifact is available before
				// selecting it as active with UpdateHostedDeployment.
				if err := s.getHostedDeploymentFromWorkRequest(
					ctx,
					addArtifactResponse.OpcWorkRequestId,
					tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"),
					oci_generative_ai.ActionTypeCreated,
					s.D.Timeout(schema.TimeoutUpdate),
				); err != nil {
					return err
				}

				// These fields identify the previously active artifact in Terraform state. The
				// update API resolves the newly added artifact by container URI and tag.
				tmp.Id = nil
				tmp.HostedDeploymentId = nil
				tmp.Status = ""
			}
			request.ActiveArtifact = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.HostedDeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.UpdateHostedDeployment(ctx, request)
	if err != nil {
		if response.RawResponse != nil && response.RawResponse.StatusCode >= 200 && response.RawResponse.StatusCode < 300 && err.Error() == "unexpected end of JSON input" {
			// The service currently completes this update successfully with an empty
			// response body, while the preview SDK expects a HostedDeployment body.
			// Preserve asynchronous polling when the response includes a work request.
			if headerWorkId := response.RawResponse.Header.Get("opc-work-request-id"); headerWorkId != "" {
				return s.getHostedDeploymentFromWorkRequest(
					ctx,
					&headerWorkId,
					tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"),
					oci_generative_ai.ActionTypeUpdated,
					s.D.Timeout(schema.TimeoutUpdate),
				)
			}
			return s.GetWithContext(ctx)
		}
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getHostedDeploymentFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GenerativeAiHostedDeploymentResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_generative_ai.DeleteHostedDeploymentRequest{}

	tmp := s.D.Id()
	request.HostedDeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.DeleteHostedDeployment(ctx, request)
	if err != nil {
		if s.isActiveIamDeploymentDeleteError(err) {
			return s.deleteIamApplicationAndDeployment(ctx)
		}
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := hostedDeploymentWaitForWorkRequest(ctx, workId, "hosteddeployment",
		oci_generative_ai.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiHostedDeploymentResourceCrud) isActiveIamDeploymentDeleteError(err error) bool {
	hostedApplicationId, ok := s.D.GetOkExists("hosted_application_id")
	if !ok || !strings.HasPrefix(hostedApplicationId.(string), "ocid1.generativeaihostedapplicationiam.") {
		return false
	}

	serviceErr, ok := oci_common.IsServiceError(err)
	return ok && serviceErr.GetHTTPStatusCode() == 403 && serviceErr.GetCode() == "NotAllowed" &&
		strings.Contains(serviceErr.GetMessage(), "activeDeployment")
}

func (s *GenerativeAiHostedDeploymentResourceCrud) deleteIamApplicationAndDeployment(ctx context.Context) error {
	hostedApplicationId := s.D.Get("hosted_application_id").(string)
	request := oci_generative_ai.DeleteHostedApplicationIamRequest{
		HostedApplicationIamId: &hostedApplicationId,
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.DeleteHostedApplicationIam(ctx, request)
	if err != nil {
		if serviceErr, ok := oci_common.IsServiceError(err); ok && serviceErr.GetHTTPStatusCode() == 404 {
			return nil
		}
		return err
	}
	if response.OpcWorkRequestId == nil || *response.OpcWorkRequestId == "" {
		return fmt.Errorf("delete IAM application response did not include opc-work-request-id")
	}

	_, err = hostedApplicationIamWaitForWorkRequest(
		ctx,
		response.OpcWorkRequestId,
		"hostedapplication",
		oci_generative_ai.ActionTypeDeleted,
		s.D.Timeout(schema.TimeoutDelete),
		s.DisableNotFoundRetries,
		s.Client,
	)
	return err
}

func (s *GenerativeAiHostedDeploymentResourceCrud) SetData() error {
	if s.Res.ActiveArtifact != nil {
		activeArtifactArray := []interface{}{}
		if activeArtifactMap := ArtifactToMap(s.Res.ActiveArtifact); activeArtifactMap != nil {
			activeArtifactArray = append(activeArtifactArray, activeArtifactMap)
		}
		s.D.Set("active_artifact", activeArtifactArray)
	} else {
		s.D.Set("active_artifact", nil)
	}

	artifacts := []interface{}{}
	for _, item := range s.Res.Artifacts {
		artifacts = append(artifacts, ArtifactToMap(item))
	}
	s.D.Set("artifacts", artifacts)

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

	if s.Res.HostedApplicationId != nil {
		s.D.Set("hosted_application_id", *s.Res.HostedApplicationId)
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

func (s *GenerativeAiHostedDeploymentResourceCrud) mapToSingleDockerArtifact(fieldKeyFormat string) (*oci_generative_ai.SingleDockerArtifact, error) {
	result := &oci_generative_ai.SingleDockerArtifact{}

	if containerUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_uri")); ok {
		tmp := containerUri.(string)
		result.ContainerUri = &tmp
	}

	if hostedDeploymentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hosted_deployment_id")); ok {
		tmp := hostedDeploymentId.(string)
		result.HostedDeploymentId = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if isVulnerabilityScanRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_vulnerability_scan_required")); ok {
		tmp := isVulnerabilityScanRequired.(bool)
		result.IsVulnerabilityScanRequired = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_generative_ai.ArtifactStatusEnum(status.(string))
	}

	if tag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag")); ok {
		tmp := tag.(string)
		result.Tag = &tmp
	}

	return result, nil
}

func (s *GenerativeAiHostedDeploymentResourceCrud) mapToCreateArtifactDetails(fieldKeyFormat string) (oci_generative_ai.CreateArtifactDetails, error) {
	result := oci_generative_ai.CreateSingleDockerArtifactDetails{}

	if containerUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_uri")); ok {
		tmp := containerUri.(string)
		result.ContainerUri = &tmp
	}

	if isVulnerabilityScanRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_vulnerability_scan_required")); ok {
		tmp := isVulnerabilityScanRequired.(bool)
		result.IsVulnerabilityScanRequired = &tmp
	}

	if tag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag")); ok {
		tmp := tag.(string)
		result.Tag = &tmp
	}

	return result, nil
}

func ArtifactToMap(obj oci_generative_ai.Artifact) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_generative_ai.SingleDockerArtifact:
		result["artifact_type"] = "SIMPLE_DOCKER_ARTIFACT"

		if v.ContainerUri != nil {
			result["container_uri"] = string(*v.ContainerUri)
		}

		if v.Tag != nil {
			result["tag"] = string(*v.Tag)
		}

		if v.HostedDeploymentId != nil {
			result["hosted_deployment_id"] = string(*v.HostedDeploymentId)
		}

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.IsVulnerabilityScanRequired != nil {
			result["is_vulnerability_scan_required"] = bool(*v.IsVulnerabilityScanRequired)
		}

		result["status"] = string(v.Status)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}
	default:
		log.Printf("[WARN] Received 'artifact_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func HostedDeploymentSummaryToMap(obj oci_generative_ai.HostedDeploymentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActiveArtifact != nil {
		activeArtifactArray := []interface{}{}
		if activeArtifactMap := ArtifactToMap(obj.ActiveArtifact); activeArtifactMap != nil {
			activeArtifactArray = append(activeArtifactArray, activeArtifactMap)
		}
		result["active_artifact"] = activeArtifactArray
	}

	artifacts := []interface{}{}
	for _, item := range obj.Artifacts {
		artifacts = append(artifacts, ArtifactToMap(item))
	}
	result["artifacts"] = artifacts

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostedApplicationId != nil {
		result["hosted_application_id"] = string(*obj.HostedApplicationId)
	}

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

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
