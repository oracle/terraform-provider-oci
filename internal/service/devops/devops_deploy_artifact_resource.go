// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsDeployArtifactResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsDeployArtifact,
		Read:     readDevopsDeployArtifact,
		Update:   updateDevopsDeployArtifact,
		Delete:   deleteDevopsDeployArtifact,
		Schema: map[string]*schema.Schema{
			// Required
			"argument_substitution_mode": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deploy_artifact_source": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"deploy_artifact_source_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"GENERIC_ARTIFACT",
								"HELM_CHART",
								"HELM_COMMAND_SPEC",
								"INLINE",
								"OCIR",
							}, true),
						},

						// Optional
						"base64encoded_content": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"chart_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"deploy_artifact_path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"deploy_artifact_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"helm_artifact_source_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"helm_verification_key_source": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"verification_key_source_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"INLINE_PUBLIC_KEY",
											"NONE",
											"VAULT_SECRET",
										}, true),
									},

									// Optional
									"current_public_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"previous_public_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"vault_secret_id": {
										Type:     schema.TypeString,
										Optional: true,
									},

									// Computed
								},
							},
						},
						"image_digest": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"image_uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"repository_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"deploy_artifact_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
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
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
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

func createDevopsDeployArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsDeployArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsDeployArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsDeployArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DevopsDeployArtifactResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.DeployArtifact
	DisableNotFoundRetries bool
}

func (s *DevopsDeployArtifactResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DevopsDeployArtifactResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_devops.DeployArtifactLifecycleStateCreating),
	}
}

func (s *DevopsDeployArtifactResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_devops.DeployArtifactLifecycleStateActive),
	}
}

func (s *DevopsDeployArtifactResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_devops.DeployArtifactLifecycleStateDeleting),
	}
}

func (s *DevopsDeployArtifactResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_devops.DeployArtifactLifecycleStateDeleted),
	}
}

func (s *DevopsDeployArtifactResourceCrud) Create() error {
	request := oci_devops.CreateDeployArtifactRequest{}

	if argumentSubstitutionMode, ok := s.D.GetOkExists("argument_substitution_mode"); ok {
		request.ArgumentSubstitutionMode = oci_devops.DeployArtifactArgumentSubstitutionModeEnum(argumentSubstitutionMode.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if deployArtifactSource, ok := s.D.GetOkExists("deploy_artifact_source"); ok {
		if tmpList := deployArtifactSource.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_artifact_source", 0)
			tmp, err := s.mapToDeployArtifactSource(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DeployArtifactSource = tmp
		}
	}

	if deployArtifactType, ok := s.D.GetOkExists("deploy_artifact_type"); ok {
		request.DeployArtifactType = oci_devops.DeployArtifactDeployArtifactTypeEnum(deployArtifactType.(string))
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

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.CreateDeployArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDeployArtifactFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DevopsDeployArtifactResourceCrud) getDeployArtifactFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_devops.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	deployArtifactId, err := deployArtifactWaitForWorkRequest(workId, "artifact",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*deployArtifactId)

	return s.Get()
}

func deployArtifactWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "devops", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_devops.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func deployArtifactWaitForWorkRequest(wId *string, entityType string, action oci_devops.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_devops.DevopsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "devops")
	retryPolicy.ShouldRetryOperation = deployArtifactWorkRequestShouldRetryFunc(timeout)

	response := oci_devops.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_devops.OperationStatusInProgress),
			string(oci_devops.OperationStatusAccepted),
			string(oci_devops.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_devops.OperationStatusSucceeded),
			string(oci_devops.OperationStatusFailed),
			string(oci_devops.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_devops.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_devops.OperationStatusFailed {
		return nil, getErrorFromDevopsDeployArtifactWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDevopsDeployArtifactWorkRequest(client *oci_devops.DevopsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_devops.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_devops.ListWorkRequestErrorsRequest{
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

func (s *DevopsDeployArtifactResourceCrud) Get() error {
	request := oci_devops.GetDeployArtifactRequest{}

	tmp := s.D.Id()
	request.DeployArtifactId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetDeployArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DeployArtifact
	return nil
}

func (s *DevopsDeployArtifactResourceCrud) Update() error {
	request := oci_devops.UpdateDeployArtifactRequest{}

	if argumentSubstitutionMode, ok := s.D.GetOkExists("argument_substitution_mode"); ok {
		request.ArgumentSubstitutionMode = oci_devops.DeployArtifactArgumentSubstitutionModeEnum(argumentSubstitutionMode.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.DeployArtifactId = &tmp

	if deployArtifactSource, ok := s.D.GetOkExists("deploy_artifact_source"); ok {
		if tmpList := deployArtifactSource.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_artifact_source", 0)
			tmp, err := s.mapToDeployArtifactSource(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DeployArtifactSource = tmp
		}
	}

	if deployArtifactType, ok := s.D.GetOkExists("deploy_artifact_type"); ok {
		request.DeployArtifactType = oci_devops.DeployArtifactDeployArtifactTypeEnum(deployArtifactType.(string))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateDeployArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeployArtifactFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DevopsDeployArtifactResourceCrud) Delete() error {
	request := oci_devops.DeleteDeployArtifactRequest{}

	tmp := s.D.Id()
	request.DeployArtifactId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.DeleteDeployArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := deployArtifactWaitForWorkRequest(workId, "artifact",
		oci_devops.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DevopsDeployArtifactResourceCrud) SetData() error {
	s.D.Set("argument_substitution_mode", s.Res.ArgumentSubstitutionMode)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeployArtifactSource != nil {
		deployArtifactSourceArray := []interface{}{}
		if deployArtifactSourceMap := DeployArtifactSourceToMap(&s.Res.DeployArtifactSource); deployArtifactSourceMap != nil {
			deployArtifactSourceArray = append(deployArtifactSourceArray, deployArtifactSourceMap)
		}
		s.D.Set("deploy_artifact_source", deployArtifactSourceArray)
	} else {
		s.D.Set("deploy_artifact_source", nil)
	}

	s.D.Set("deploy_artifact_type", s.Res.DeployArtifactType)

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

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
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

func (s *DevopsDeployArtifactResourceCrud) mapToDeployArtifactSource(fieldKeyFormat string) (oci_devops.DeployArtifactSource, error) {

	if deployArtifactSourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deploy_artifact_source_type")); ok {
		tmp := deployArtifactSourceType.(string)
		switch tmp {
		case "GENERIC_ARTIFACT":
			result := oci_devops.GenericDeployArtifactSource{}
			if repositoryId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_id")); ok {
				tmp := repositoryId.(string)
				result.RepositoryId = &tmp
			}
			if deployArtifactPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deploy_artifact_path")); ok {
				tmp := deployArtifactPath.(string)
				result.DeployArtifactPath = &tmp
			}
			if deployArtifactVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deploy_artifact_version")); ok {
				tmp := deployArtifactVersion.(string)
				result.DeployArtifactVersion = &tmp
			}
			return result, nil
		case "HELM_CHART":
			result := oci_devops.HelmRepositoryDeployArtifactSource{}
			if chartUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "chart_url")); ok {
				tmp := chartUrl.(string)
				result.ChartUrl = &tmp
			}
			if deployArtifactVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deploy_artifact_version")); ok {
				tmp := deployArtifactVersion.(string)
				result.DeployArtifactVersion = &tmp
			}
			if helmVerificationKeySource, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "helm_verification_key_source")); ok {
				if tmpList := helmVerificationKeySource.([]interface{}); len(tmpList) > 0 {
					helmKeyFieldKeyFormat := fmt.Sprintf("%s.%d.%s.%d.%%s", "deploy_artifact_source", 0, "helm_verification_key_source", 0)
					tmp, err := s.mapToVerificationKeySource(helmKeyFieldKeyFormat)
					if err != nil {
						return nil, err
					}
					result.HelmVerificationKeySource = tmp
				}
			}
			return result, nil
		case "INLINE":
			result := oci_devops.InlineDeployArtifactSource{}
			if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base64encoded_content")); ok {
				tmp := []byte(content.(string))
				result.Base64EncodedContent = tmp
			}
			return result, nil
		case "OCIR":
			result := oci_devops.OcirDeployArtifactSource{}
			if imageUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_uri")); ok {
				tmp := imageUri.(string)
				result.ImageUri = &tmp
			}
			if imageDigest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_digest")); ok {
				tmp := imageDigest.(string)
				result.ImageDigest = &tmp
			}
			return result, nil
		case "HELM_COMMAND_SPEC":
			result := oci_devops.HelmCommandSpecArtifactSource{}
			if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base64encoded_content")); ok {
				tmp := content.(string)
				result.Base64EncodedContent = &tmp
			}
			if helmArtifactSourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "helm_artifact_source_type")); ok {
				result.HelmArtifactSourceType = oci_devops.HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum(helmArtifactSourceType.(string))
			}
			return result, nil
		default:
			return nil, fmt.Errorf("[ERROR] Received 'deploy_artifact_source_type' of unknown type %v", tmp)
		}
	}
	return nil, fmt.Errorf("[ERROR] Unable to mapToDeployArtifactSource")
}

func DeployArtifactSourceToMap(obj *oci_devops.DeployArtifactSource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.GenericDeployArtifactSource:
		result["deploy_artifact_source_type"] = "GENERIC_ARTIFACT"

		if v.DeployArtifactPath != nil {
			result["deploy_artifact_path"] = string(*v.DeployArtifactPath)
		}

		if v.DeployArtifactVersion != nil {
			result["deploy_artifact_version"] = string(*v.DeployArtifactVersion)
		}

		if v.RepositoryId != nil {
			result["repository_id"] = string(*v.RepositoryId)
		}
	case oci_devops.HelmRepositoryDeployArtifactSource:

		result["deploy_artifact_source_type"] = "HELM_CHART"

		if v.ChartUrl != nil {
			result["chart_url"] = string(*v.ChartUrl)
		}

		if v.DeployArtifactVersion != nil {
			result["deploy_artifact_version"] = string(*v.DeployArtifactVersion)
		}

		if v.HelmVerificationKeySource != nil {
			helmVerificationKeySourceArray := []interface{}{}
			if helmVerificationKeySourceMap := VerificationKeySourceToMap(&v.HelmVerificationKeySource); helmVerificationKeySourceMap != nil {
				helmVerificationKeySourceArray = append(helmVerificationKeySourceArray, helmVerificationKeySourceMap)
			}
			result["helm_verification_key_source"] = helmVerificationKeySourceArray
		}
	case oci_devops.HelmCommandSpecArtifactSource:
		result["deploy_artifact_source_type"] = "HELM_COMMAND_SPEC"

		if v.Base64EncodedContent != nil {
			result["base64encoded_content"] = string(*v.Base64EncodedContent)
		}

		result["helm_artifact_source_type"] = string(v.HelmArtifactSourceType)
	case oci_devops.InlineDeployArtifactSource:
		result["deploy_artifact_source_type"] = "INLINE"

		if v.Base64EncodedContent != nil {
			contentReader := v.Base64EncodedContent
			result["base64encoded_content"] = DevopsDeployArtifactBase64Decode(contentReader)
		}
	case oci_devops.OcirDeployArtifactSource:
		result["deploy_artifact_source_type"] = "OCIR"

		if v.ImageDigest != nil {
			result["image_digest"] = string(*v.ImageDigest)
		}

		if v.ImageUri != nil {
			result["image_uri"] = string(*v.ImageUri)
		}
	default:
		log.Printf("[WARN] Received 'deploy_artifact_source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DeployArtifactSummaryToMap(obj oci_devops.DeployArtifactSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["argument_substitution_mode"] = string(obj.ArgumentSubstitutionMode)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DeployArtifactSource != nil {
		deployArtifactSourceArray := []interface{}{}
		if deployArtifactSourceMap := DeployArtifactSourceToMap(&obj.DeployArtifactSource); deployArtifactSourceMap != nil {
			deployArtifactSourceArray = append(deployArtifactSourceArray, deployArtifactSourceMap)
		}
		result["deploy_artifact_source"] = deployArtifactSourceArray
	}

	result["deploy_artifact_type"] = string(obj.DeployArtifactType)

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

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
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

func (s *DevopsDeployArtifactResourceCrud) mapToVerificationKeySource(fieldKeyFormat string) (oci_devops.VerificationKeySource, error) {
	var baseObject oci_devops.VerificationKeySource
	//discriminator
	verificationKeySourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verification_key_source_type"))
	var verificationKeySourceType string
	if ok {
		verificationKeySourceType = verificationKeySourceTypeRaw.(string)
	} else {
		verificationKeySourceType = "" // default value
	}
	switch strings.ToLower(verificationKeySourceType) {
	case strings.ToLower("INLINE_PUBLIC_KEY"):
		details := oci_devops.InlinePublicKeyVerificationKeySource{}
		if currentPublicKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "current_public_key")); ok {
			tmp := currentPublicKey.(string)
			details.CurrentPublicKey = &tmp
		}
		if previousPublicKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "previous_public_key")); ok {
			tmp := previousPublicKey.(string)
			details.PreviousPublicKey = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_devops.NoneVerificationKeySource{}
		baseObject = details
	case strings.ToLower("VAULT_SECRET"):
		details := oci_devops.VaultSecretVerificationKeySource{}
		if vaultSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_secret_id")); ok {
			tmp := vaultSecretId.(string)
			details.VaultSecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown verification_key_source_type '%v' was specified", verificationKeySourceType)
	}
	return baseObject, nil
}

func VerificationKeySourceToMap(obj *oci_devops.VerificationKeySource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.InlinePublicKeyVerificationKeySource:
		result["verification_key_source_type"] = "INLINE_PUBLIC_KEY"

		if v.CurrentPublicKey != nil {
			result["current_public_key"] = string(*v.CurrentPublicKey)
		}

		if v.PreviousPublicKey != nil {
			result["previous_public_key"] = string(*v.PreviousPublicKey)
		}
	case oci_devops.NoneVerificationKeySource:
		result["verification_key_source_type"] = "NONE"
	case oci_devops.VaultSecretVerificationKeySource:
		result["verification_key_source_type"] = "VAULT_SECRET"

		if v.VaultSecretId != nil {
			result["vault_secret_id"] = string(*v.VaultSecretId)
		}
	default:
		log.Printf("[WARN] Received 'verification_key_source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DevopsDeployArtifactBase64Decode(content []byte) string {
	text := b64.StdEncoding.EncodeToString(content)
	return text
}
