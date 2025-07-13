// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceModelGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatascienceModelGroup,
		Read:     readDatascienceModelGroup,
		Update:   updateDatascienceModelGroup,
		Delete:   deleteDatascienceModelGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"create_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"CLONE",
					"CREATE",
				}, true),
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
			"member_model_entries": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"member_model_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"inference_key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"model_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"model_group_clone_source_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"model_group_clone_source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"MODEL_GROUP",
								"MODEL_GROUP_VERSION_HISTORY",
							}, true),
						},
						"source_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"modify_model_group_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"defined_tags": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
										Elem:             schema.TypeString,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
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
										ForceNew: true,
										Elem:     schema.TypeString,
									},
									"model_group_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"type": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"HETEROGENEOUS",
														"HOMOGENEOUS",
														"STACKED",
													}, true),
												},

												// Optional
												"base_model_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"custom_metadata_list": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"category": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"description": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},
									"model_group_version_history_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"version_label": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"patch_model_group_member_model_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"items": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"operation": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"INSERT",
														"REMOVE",
													}, true),
												},
												"values": {
													Type:     schema.TypeList,
													Required: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"model_id": {
																Type:     schema.TypeString,
																Required: true,
																ForceNew: true,
															},

															// Optional
															"inference_key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},

												// Optional

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"model_group_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"HETEROGENEOUS",
								"HOMOGENEOUS",
								"STACKED",
							}, true),
						},

						// Optional
						"base_model_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"custom_metadata_list": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"category": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"model_group_version_history_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version_label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"model_group_version_history_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_model_group_id": {
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
			"version_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatascienceModelGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.CreateResource(d, sync)
}

func readDatascienceModelGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

func updateDatascienceModelGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatascienceModelGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatascienceModelGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.ModelGroup
	DisableNotFoundRetries bool
}

func (s *DatascienceModelGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceModelGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datascience.ModelGroupLifecycleStateCreating),
	}
}

func (s *DatascienceModelGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.ModelGroupLifecycleStateActive),
	}
}

func (s *DatascienceModelGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datascience.ModelGroupLifecycleStateDeleting),
	}
}

func (s *DatascienceModelGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.ModelGroupLifecycleStateDeleted),
	}
}

func (s *DatascienceModelGroupResourceCrud) Create() error {
	request := oci_datascience.CreateModelGroupRequest{}
	err := s.populateTopLevelPolymorphicCreateModelGroupRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateModelGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getModelGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatascienceModelGroupResourceCrud) getModelGroupFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datascience.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	modelGroupId, err := modelGroupWaitForWorkRequest(workId, "model-group",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, modelGroupId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_datascience.CancelWorkRequestRequest{
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
	s.D.SetId(*modelGroupId)

	return s.Get()
}

func modelGroupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "datascience", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_datascience.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func modelGroupWaitForWorkRequest(wId *string, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datascience.DataScienceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datascience")
	retryPolicy.ShouldRetryOperation = modelGroupWorkRequestShouldRetryFunc(timeout)

	response := oci_datascience.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_datascience.WorkRequestStatusInProgress),
			string(oci_datascience.WorkRequestStatusAccepted),
			string(oci_datascience.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_datascience.WorkRequestStatusSucceeded),
			string(oci_datascience.WorkRequestStatusFailed),
			string(oci_datascience.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_datascience.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_datascience.WorkRequestStatusFailed || response.Status == oci_datascience.WorkRequestStatusCanceled {
		return nil, getErrorFromDatascienceModelGroupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatascienceModelGroupWorkRequest(client *oci_datascience.DataScienceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_datascience.ListWorkRequestErrorsRequest{
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

func (s *DatascienceModelGroupResourceCrud) Get() error {
	request := oci_datascience.GetModelGroupRequest{}

	tmp := s.D.Id()
	request.ModelGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetModelGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ModelGroup
	return nil
}

func (s *DatascienceModelGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateModelGroupRequest{}

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
	request.ModelGroupId = &tmp

	if modelGroupVersionHistoryId, ok := s.D.GetOkExists("model_group_version_history_id"); ok {
		tmp := modelGroupVersionHistoryId.(string)
		request.ModelGroupVersionHistoryId = &tmp
	}

	if versionLabel, ok := s.D.GetOkExists("version_label"); ok {
		tmp := versionLabel.(string)
		request.VersionLabel = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateModelGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ModelGroup
	return nil
}

func (s *DatascienceModelGroupResourceCrud) Delete() error {
	request := oci_datascience.DeleteModelGroupRequest{}

	tmp := s.D.Id()
	request.ModelGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.DeleteModelGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := modelGroupWaitForWorkRequest(workId, "model-group",
		oci_datascience.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatascienceModelGroupResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreateType != nil {
		s.D.Set("create_type", *s.Res.CreateType)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MemberModelEntries != nil {
		s.D.Set("member_model_entries", []interface{}{MemberModelEntriesToMap(s.Res.MemberModelEntries)})
	} else {
		s.D.Set("member_model_entries", nil)
	}

	if s.Res.ModelGroupDetails != nil {
		modelGroupDetailsArray := []interface{}{}
		if modelGroupDetailsMap := ModelGroupDetailsToMap(&s.Res.ModelGroupDetails); modelGroupDetailsMap != nil {
			modelGroupDetailsArray = append(modelGroupDetailsArray, modelGroupDetailsMap)
		}
		s.D.Set("model_group_details", modelGroupDetailsArray)
	} else {
		s.D.Set("model_group_details", nil)
	}

	if s.Res.ModelGroupVersionHistoryId != nil {
		s.D.Set("model_group_version_history_id", *s.Res.ModelGroupVersionHistoryId)
	}

	if s.Res.ModelGroupVersionHistoryName != nil {
		s.D.Set("model_group_version_history_name", *s.Res.ModelGroupVersionHistoryName)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	if s.Res.SourceModelGroupId != nil {
		s.D.Set("source_model_group_id", *s.Res.SourceModelGroupId)
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

	if s.Res.VersionId != nil {
		s.D.Set("version_id", strconv.FormatInt(*s.Res.VersionId, 10))
	}

	if s.Res.VersionLabel != nil {
		s.D.Set("version_label", *s.Res.VersionLabel)
	}

	return nil
}

func (s *DatascienceModelGroupResourceCrud) mapToCustomMetadata(fieldKeyFormat string) (oci_datascience.CustomMetadata, error) {
	result := oci_datascience.CustomMetadata{}

	if category, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "category")); ok {
		tmp := category.(string)
		result.Category = &tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func CustomMetadataToMap(obj oci_datascience.CustomMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Category != nil {
		result["category"] = string(*obj.Category)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DatascienceModelGroupResourceCrud) mapToMemberModelDetails(fieldKeyFormat string) (oci_datascience.MemberModelDetails, error) {
	result := oci_datascience.MemberModelDetails{}

	if inferenceKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "inference_key")); ok {
		tmp := inferenceKey.(string)
		result.InferenceKey = &tmp
	}

	if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
		tmp := modelId.(string)
		result.ModelId = &tmp
	}

	return result, nil
}

func MemberModelDetailsToMap(obj oci_datascience.MemberModelDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InferenceKey != nil {
		result["inference_key"] = string(*obj.InferenceKey)
	}

	if obj.ModelId != nil {
		result["model_id"] = string(*obj.ModelId)
	}

	return result
}

func (s *DatascienceModelGroupResourceCrud) mapToMemberModelEntries(fieldKeyFormat string) (oci_datascience.MemberModelEntries, error) {
	result := oci_datascience.MemberModelEntries{}

	if memberModelDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_model_details")); ok {
		interfaces := memberModelDetails.([]interface{})
		tmp := make([]oci_datascience.MemberModelDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "member_model_details"), stateDataIndex)
			converted, err := s.mapToMemberModelDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "member_model_details")) {
			result.MemberModelDetails = tmp
		}
	}

	return result, nil
}

func MemberModelEntriesToMap(obj *oci_datascience.MemberModelEntries) map[string]interface{} {
	result := map[string]interface{}{}

	memberModelDetails := []interface{}{}
	for _, item := range obj.MemberModelDetails {
		memberModelDetails = append(memberModelDetails, MemberModelDetailsToMap(item))
	}
	result["member_model_details"] = memberModelDetails

	return result
}

func (s *DatascienceModelGroupResourceCrud) mapToModelGroupCloneSourceDetails(fieldKeyFormat string) (oci_datascience.ModelGroupCloneSourceDetails, error) {
	var baseObject oci_datascience.ModelGroupCloneSourceDetails
	//discriminator
	modelGroupCloneSourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_group_clone_source_type"))
	var modelGroupCloneSourceType string
	if ok {
		modelGroupCloneSourceType = modelGroupCloneSourceTypeRaw.(string)
	} else {
		modelGroupCloneSourceType = "" // default value
	}
	switch strings.ToLower(modelGroupCloneSourceType) {
	case strings.ToLower("MODEL_GROUP"):
		details := oci_datascience.CloneCreateFromModelGroupDetails{}
		if modifyModelGroupDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "modify_model_group_details")); ok {
			if tmpList := modifyModelGroupDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "modify_model_group_details"), 0)
				tmp, err := s.mapToModifyModelGroupDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert modify_model_group_details, encountered error: %v", err)
				}
				details.ModifyModelGroupDetails = &tmp
			}
		}
		if patchModelGroupMemberModelDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patch_model_group_member_model_details")); ok {
			if tmpList := patchModelGroupMemberModelDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "patch_model_group_member_model_details"), 0)
				tmp, err := s.mapToPatchModelGroupMemberModelDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert patch_model_group_member_model_details, encountered error: %v", err)
				}
				details.PatchModelGroupMemberModelDetails = &tmp
			}
		}
		if sourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_id")); ok {
			tmp := sourceId.(string)
			details.SourceId = &tmp
		}
		baseObject = details
	case strings.ToLower("MODEL_GROUP_VERSION_HISTORY"):
		details := oci_datascience.CloneCreateFromModelGroupVersionHistoryDetails{}
		if modifyModelGroupDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "modify_model_group_details")); ok {
			if tmpList := modifyModelGroupDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "modify_model_group_details"), 0)
				tmp, err := s.mapToModifyModelGroupDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert modify_model_group_details, encountered error: %v", err)
				}
				details.ModifyModelGroupDetails = &tmp
			}
		}
		if patchModelGroupMemberModelDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patch_model_group_member_model_details")); ok {
			if tmpList := patchModelGroupMemberModelDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "patch_model_group_member_model_details"), 0)
				tmp, err := s.mapToPatchModelGroupMemberModelDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert patch_model_group_member_model_details, encountered error: %v", err)
				}
				details.PatchModelGroupMemberModelDetails = &tmp
			}
		}
		if sourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_id")); ok {
			tmp := sourceId.(string)
			details.SourceId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown model_group_clone_source_type '%v' was specified", modelGroupCloneSourceType)
	}
	return baseObject, nil
}

func ModelGroupCloneSourceDetailsToMap(obj *oci_datascience.ModelGroupCloneSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.CloneCreateFromModelGroupDetails:
		result["model_group_clone_source_type"] = "MODEL_GROUP"

		if v.ModifyModelGroupDetails != nil {
			result["modify_model_group_details"] = []interface{}{ModifyModelGroupDetailsToMap(v.ModifyModelGroupDetails)}
		}

		if v.PatchModelGroupMemberModelDetails != nil {
			result["patch_model_group_member_model_details"] = []interface{}{PatchModelGroupMemberModelDetailsToMap(v.PatchModelGroupMemberModelDetails)}
		}

		if v.SourceId != nil {
			result["source_id"] = string(*v.SourceId)
		}
	case oci_datascience.CloneCreateFromModelGroupVersionHistoryDetails:
		result["model_group_clone_source_type"] = "MODEL_GROUP_VERSION_HISTORY"

		if v.ModifyModelGroupDetails != nil {
			result["modify_model_group_details"] = []interface{}{ModifyModelGroupDetailsToMap(v.ModifyModelGroupDetails)}
		}

		if v.PatchModelGroupMemberModelDetails != nil {
			result["patch_model_group_member_model_details"] = []interface{}{PatchModelGroupMemberModelDetailsToMap(v.PatchModelGroupMemberModelDetails)}
		}

		if v.SourceId != nil {
			result["source_id"] = string(*v.SourceId)
		}
	default:
		log.Printf("[WARN] Received 'model_group_clone_source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceModelGroupResourceCrud) mapToModelGroupDetails(fieldKeyFormat string) (oci_datascience.ModelGroupDetails, error) {
	var baseObject oci_datascience.ModelGroupDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("HETEROGENEOUS"):
		details := oci_datascience.HeterogeneousModelGroupDetails{}
		if customMetadataList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_metadata_list")); ok {
			interfaces := customMetadataList.([]interface{})
			tmp := make([]oci_datascience.CustomMetadata, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "custom_metadata_list"), stateDataIndex)
				converted, err := s.mapToCustomMetadata(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "custom_metadata_list")) {
				details.CustomMetadataList = tmp
			}
		}
		baseObject = details
	case strings.ToLower("HOMOGENEOUS"):
		details := oci_datascience.HomogeneousModelGroupDetails{}
		if customMetadataList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_metadata_list")); ok {
			interfaces := customMetadataList.([]interface{})
			tmp := make([]oci_datascience.CustomMetadata, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "custom_metadata_list"), stateDataIndex)
				converted, err := s.mapToCustomMetadata(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "custom_metadata_list")) {
				details.CustomMetadataList = tmp
			}
		}
		baseObject = details
	case strings.ToLower("STACKED"):
		details := oci_datascience.StackedModelGroupDetails{}
		if baseModelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_model_id")); ok {
			tmp := baseModelId.(string)
			details.BaseModelId = &tmp
		}
		if customMetadataList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_metadata_list")); ok {
			interfaces := customMetadataList.([]interface{})
			tmp := make([]oci_datascience.CustomMetadata, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "custom_metadata_list"), stateDataIndex)
				converted, err := s.mapToCustomMetadata(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "custom_metadata_list")) {
				details.CustomMetadataList = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ModelGroupDetailsToMap(obj *oci_datascience.ModelGroupDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.HeterogeneousModelGroupDetails:
		result["type"] = "HETEROGENEOUS"
	case oci_datascience.HomogeneousModelGroupDetails:
		result["type"] = "HOMOGENEOUS"
	case oci_datascience.StackedModelGroupDetails:
		result["type"] = "STACKED"

		if v.BaseModelId != nil {
			result["base_model_id"] = string(*v.BaseModelId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceModelGroupResourceCrud) mapToModifyModelGroupDetails(fieldKeyFormat string) (oci_datascience.ModifyModelGroupDetails, error) {
	result := oci_datascience.ModifyModelGroupDetails{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if modelGroupDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_group_details")); ok {
		if tmpList := modelGroupDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "model_group_details"), 0)
			tmp, err := s.mapToModelGroupDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert model_group_details, encountered error: %v", err)
			}
			result.ModelGroupDetails = tmp
		}
	}

	if modelGroupVersionHistoryId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_group_version_history_id")); ok {
		tmp := modelGroupVersionHistoryId.(string)
		result.ModelGroupVersionHistoryId = &tmp
	}

	if versionLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_label")); ok {
		tmp := versionLabel.(string)
		result.VersionLabel = &tmp
	}

	return result, nil
}

func ModifyModelGroupDetailsToMap(obj *oci_datascience.ModifyModelGroupDetails) map[string]interface{} {
	result := map[string]interface{}{}

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

	if obj.ModelGroupDetails != nil {
		modelGroupDetailsArray := []interface{}{}
		if modelGroupDetailsMap := ModelGroupDetailsToMap(&obj.ModelGroupDetails); modelGroupDetailsMap != nil {
			modelGroupDetailsArray = append(modelGroupDetailsArray, modelGroupDetailsMap)
		}
		result["model_group_details"] = modelGroupDetailsArray
	}

	if obj.ModelGroupVersionHistoryId != nil {
		result["model_group_version_history_id"] = string(*obj.ModelGroupVersionHistoryId)
	}

	if obj.VersionLabel != nil {
		result["version_label"] = string(*obj.VersionLabel)
	}

	return result
}

func (s *DatascienceModelGroupResourceCrud) mapToPatchInstruction(fieldKeyFormat string) (oci_datascience.PatchInstruction, error) {
	var baseObject oci_datascience.PatchInstruction
	//discriminator
	operationRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation"))
	var operation string
	if ok {
		operation = operationRaw.(string)
	} else {
		operation = "" // default value
	}
	switch strings.ToLower(operation) {
	case strings.ToLower("INSERT"):
		details := oci_datascience.PatchInsertNewMemberModels{}
		if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
			interfaces := values.([]interface{})
			tmp := make([]oci_datascience.MemberModelDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "values"), stateDataIndex)
				converted, err := s.mapToMemberModelDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
				details.Values = tmp
			}
		}
		baseObject = details
	case strings.ToLower("REMOVE"):
		details := oci_datascience.PatchRemoveMemberModels{}
		if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
			interfaces := values.([]interface{})
			tmp := make([]oci_datascience.MemberModelDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "values"), stateDataIndex)
				converted, err := s.mapToMemberModelDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
				details.Values = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown operation '%v' was specified", operation)
	}
	return baseObject, nil
}

func PatchInstructionToMap(obj oci_datascience.PatchInstruction) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_datascience.PatchInsertNewMemberModels:
		result["operation"] = "INSERT"

		values := []interface{}{}
		for _, item := range v.Values {
			values = append(values, MemberModelDetailsToMap(item))
		}
		result["values"] = values
	case oci_datascience.PatchRemoveMemberModels:
		result["operation"] = "REMOVE"

		values := []interface{}{}
		for _, item := range v.Values {
			values = append(values, MemberModelDetailsToMap(item))
		}
		result["values"] = values
	default:
		log.Printf("[WARN] Received 'operation' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatascienceModelGroupResourceCrud) mapToPatchModelGroupMemberModelDetails(fieldKeyFormat string) (oci_datascience.PatchModelGroupMemberModelDetails, error) {
	result := oci_datascience.PatchModelGroupMemberModelDetails{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_datascience.PatchInstruction, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToPatchInstruction(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func PatchModelGroupMemberModelDetailsToMap(obj *oci_datascience.PatchModelGroupMemberModelDetails) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, PatchInstructionToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DatascienceModelGroupResourceCrud) populateTopLevelPolymorphicCreateModelGroupRequest(request *oci_datascience.CreateModelGroupRequest) error {
	//discriminator
	createTypeRaw, ok := s.D.GetOkExists("create_type")
	var createType string
	if ok {
		createType = createTypeRaw.(string)
	} else {
		createType = "" // default value
	}
	switch strings.ToLower(createType) {
	case strings.ToLower("CLONE"):
		details := oci_datascience.CloneModelGroupDetails{}
		if modelGroupCloneSourceDetails, ok := s.D.GetOkExists("model_group_clone_source_details"); ok {
			if tmpList := modelGroupCloneSourceDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "model_group_clone_source_details", 0)
				tmp, err := s.mapToModelGroupCloneSourceDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ModelGroupCloneSourceDetails = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if projectId, ok := s.D.GetOkExists("project_id"); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		request.CreateBaseModelGroupDetails = details
	case strings.ToLower("CREATE"):
		details := oci_datascience.CreateModelGroupDetails{}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if memberModelEntries, ok := s.D.GetOkExists("member_model_entries"); ok {
			if tmpList := memberModelEntries.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "member_model_entries", 0)
				tmp, err := s.mapToMemberModelEntries(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.MemberModelEntries = &tmp
			}
		}
		if modelGroupDetails, ok := s.D.GetOkExists("model_group_details"); ok {
			if tmpList := modelGroupDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "model_group_details", 0)
				tmp, err := s.mapToModelGroupDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ModelGroupDetails = tmp
			}
		}
		if modelGroupVersionHistoryId, ok := s.D.GetOkExists("model_group_version_history_id"); ok {
			tmp := modelGroupVersionHistoryId.(string)
			details.ModelGroupVersionHistoryId = &tmp
		}
		if versionLabel, ok := s.D.GetOkExists("version_label"); ok {
			tmp := versionLabel.(string)
			details.VersionLabel = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if modelGroupVersionHistoryId, ok := s.D.GetOkExists("model_group_version_history_id"); ok {
			tmp := modelGroupVersionHistoryId.(string)
			details.ModelGroupVersionHistoryId = &tmp
		}
		if projectId, ok := s.D.GetOkExists("project_id"); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		if versionLabel, ok := s.D.GetOkExists("version_label"); ok {
			tmp := versionLabel.(string)
			details.VersionLabel = &tmp
		}
		request.CreateBaseModelGroupDetails = details
	default:
		return fmt.Errorf("unknown create_type '%v' was specified", createType)
	}
	return nil
}

func (s *DatascienceModelGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeModelGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ModelGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ChangeModelGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
