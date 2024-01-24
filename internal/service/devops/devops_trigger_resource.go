// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsTriggerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsTrigger,
		Read:     readDevopsTrigger,
		Update:   updateDevopsTrigger,
		Delete:   deleteDevopsTrigger,
		Schema: map[string]*schema.Schema{
			// Required
			"actions": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"build_pipeline_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"TRIGGER_BUILD_PIPELINE",
							}, true),
						},

						// Optional
						"filter": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"trigger_source": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"BITBUCKET_CLOUD",
											"BITBUCKET_SERVER",
											"DEVOPS_CODE_REPOSITORY",
											"GITHUB",
											"GITLAB",
											"GITLAB_SERVER",
											"VBS",
										}, true),
									},

									// Optional
									"events": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"exclude": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"file_filter": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"file_paths": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
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
									"include": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"base_ref": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"file_filter": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"file_paths": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															// Computed
														},
													},
												},
												"head_ref": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"repository_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

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
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"trigger_source": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"BITBUCKET_CLOUD",
					"BITBUCKET_SERVER",
					"DEVOPS_CODE_REPOSITORY",
					"GITHUB",
					"GITLAB",
					"GITLAB_SERVER",
					"VBS",
				}, true),
			},

			// Optional
			"connection_id": {
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
			"repository_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"trigger_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDevopsTrigger(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsTriggerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsTrigger(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsTriggerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsTrigger(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsTriggerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsTrigger(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsTriggerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DevopsTriggerResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.Trigger
	DisableNotFoundRetries bool
}

func (s *DevopsTriggerResourceCrud) ID() string {
	trigger := *s.Res
	return *trigger.GetId()
}

func (s *DevopsTriggerResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DevopsTriggerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_devops.TriggerLifecycleStateActive),
	}
}

func (s *DevopsTriggerResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DevopsTriggerResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DevopsTriggerResourceCrud) Create() error {
	request := oci_devops.CreateTriggerRequest{}
	err := s.populateTopLevelPolymorphicCreateTriggerRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.CreateTrigger(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getTriggerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DevopsTriggerResourceCrud) getTriggerFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_devops.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	triggerId, err := triggerWaitForWorkRequest(workId, "trigger",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*triggerId)

	return s.Get()
}

func triggerWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func triggerWaitForWorkRequest(wId *string, entityType string, action oci_devops.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_devops.DevopsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "devops")
	retryPolicy.ShouldRetryOperation = triggerWorkRequestShouldRetryFunc(timeout)

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
	if identifier == nil || response.Status == oci_devops.OperationStatusFailed || response.Status == oci_devops.OperationStatusCanceled {
		return nil, getErrorFromDevopsTriggerWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDevopsTriggerWorkRequest(client *oci_devops.DevopsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_devops.ActionTypeEnum) error {
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

func (s *DevopsTriggerResourceCrud) Get() error {
	request := oci_devops.GetTriggerRequest{}

	tmp := s.D.Id()
	request.TriggerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetTrigger(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Trigger
	return nil
}

func (s *DevopsTriggerResourceCrud) Update() error {
	request := oci_devops.UpdateTriggerRequest{}
	err := s.populateTopLevelPolymorphicUpdateTriggerRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateTrigger(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getTriggerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DevopsTriggerResourceCrud) Delete() error {
	request := oci_devops.DeleteTriggerRequest{}

	tmp := s.D.Id()
	request.TriggerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.DeleteTrigger(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := triggerWaitForWorkRequest(workId, "trigger",
		oci_devops.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DevopsTriggerResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_devops.BitbucketCloudTrigger:
		s.D.Set("trigger_source", "BITBUCKET_CLOUD")

		if v.ConnectionId != nil {
			s.D.Set("connection_id", *v.ConnectionId)
		}

		if v.TriggerUrl != nil {
			s.D.Set("trigger_url", *v.TriggerUrl)
		}

		actions := []interface{}{}
		for _, item := range v.Actions {
			actions = append(actions, TriggerActionToMap(item))
		}
		s.D.Set("actions", actions)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.BitbucketServerTrigger:
		s.D.Set("trigger_source", "BITBUCKET_SERVER")

		if v.TriggerUrl != nil {
			s.D.Set("trigger_url", *v.TriggerUrl)
		}

		actions := []interface{}{}
		for _, item := range v.Actions {
			actions = append(actions, TriggerActionToMap(item))
		}
		s.D.Set("actions", actions)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.DevopsCodeRepositoryTrigger:
		s.D.Set("trigger_source", "DEVOPS_CODE_REPOSITORY")

		if v.RepositoryId != nil {
			s.D.Set("repository_id", *v.RepositoryId)
		}

		actions := []interface{}{}
		for _, item := range v.Actions {
			actions = append(actions, TriggerActionToMap(item))
		}
		s.D.Set("actions", actions)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.GithubTrigger:
		s.D.Set("trigger_source", "GITHUB")

		if v.ConnectionId != nil {
			s.D.Set("connection_id", *v.ConnectionId)
		}

		if v.TriggerUrl != nil {
			s.D.Set("trigger_url", *v.TriggerUrl)
		}

		actions := []interface{}{}
		for _, item := range v.Actions {
			actions = append(actions, TriggerActionToMap(item))
		}
		s.D.Set("actions", actions)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.GitlabTrigger:
		s.D.Set("trigger_source", "GITLAB")

		if v.ConnectionId != nil {
			s.D.Set("connection_id", *v.ConnectionId)
		}

		if v.TriggerUrl != nil {
			s.D.Set("trigger_url", *v.TriggerUrl)
		}

		actions := []interface{}{}
		for _, item := range v.Actions {
			actions = append(actions, TriggerActionToMap(item))
		}
		s.D.Set("actions", actions)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.GitlabServerTrigger:
		s.D.Set("trigger_source", "GITLAB_SERVER")

		if v.TriggerUrl != nil {
			s.D.Set("trigger_url", *v.TriggerUrl)
		}

		actions := []interface{}{}
		for _, item := range v.Actions {
			actions = append(actions, TriggerActionToMap(item))
		}
		s.D.Set("actions", actions)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.VbsTrigger:
		s.D.Set("trigger_source", "VBS")

		if v.ConnectionId != nil {
			s.D.Set("connection_id", *v.ConnectionId)
		}

		if v.TriggerUrl != nil {
			s.D.Set("trigger_url", *v.TriggerUrl)
		}

		actions := []interface{}{}
		for _, item := range v.Actions {
			actions = append(actions, TriggerActionToMap(item))
		}
		s.D.Set("actions", actions)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'trigger_source' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *DevopsTriggerResourceCrud) mapToBitbucketCloudFilterAttributes(fieldKeyFormat string) (oci_devops.BitbucketCloudFilterAttributes, error) {
	result := oci_devops.BitbucketCloudFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func TriggerBitbucketCloudFilterAttributesToMap(obj *oci_devops.BitbucketCloudFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseRef != nil {
		result["base_ref"] = string(*obj.BaseRef)
	}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{TriggerFileFilterToMap(obj.FileFilter)}
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	return result
}

func (s *DevopsTriggerResourceCrud) mapToBitbucketCloudFilterExclusionAttributes(fieldKeyFormat string) (oci_devops.BitbucketCloudFilterExclusionAttributes, error) {
	result := oci_devops.BitbucketCloudFilterExclusionAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	return result, nil
}

func TriggerBitbucketCloudFilterExclusionAttributesToMap(obj *oci_devops.BitbucketCloudFilterExclusionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{TriggerFileFilterToMap(obj.FileFilter)}
	}

	return result
}

func (s *DevopsTriggerResourceCrud) mapToBitbucketServerFilterAttributes(fieldKeyFormat string) (oci_devops.BitbucketServerFilterAttributes, error) {
	result := oci_devops.BitbucketServerFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func (s *DevopsTriggerResourceCrud) mapToDevopsCodeRepositoryFilterAttributes(fieldKeyFormat string) (oci_devops.DevopsCodeRepositoryFilterAttributes, error) {
	result := oci_devops.DevopsCodeRepositoryFilterAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func TriggerDevopsCodeRepositoryFilterAttributesToMap(obj *oci_devops.DevopsCodeRepositoryFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{TriggerFileFilterToMap(obj.FileFilter)}
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	return result
}

func (s *DevopsTriggerResourceCrud) mapToDevopsCodeRepositoryFilterExclusionAttributes(fieldKeyFormat string) (oci_devops.DevopsCodeRepositoryFilterExclusionAttributes, error) {
	result := oci_devops.DevopsCodeRepositoryFilterExclusionAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	return result, nil
}

func TriggerDevopsCodeRepositoryFilterExclusionAttributesToMap(obj *oci_devops.DevopsCodeRepositoryFilterExclusionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{TriggerFileFilterToMap(obj.FileFilter)}
	}

	return result
}

func (s *DevopsTriggerResourceCrud) mapToFileFilter(fieldKeyFormat string) (oci_devops.FileFilter, error) {
	result := oci_devops.FileFilter{}

	if filePaths, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_paths")); ok {
		interfaces := filePaths.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "file_paths")) {
			result.FilePaths = tmp
		}
	}

	return result, nil
}

func TriggerFileFilterToMap(obj *oci_devops.FileFilter) map[string]interface{} {
	result := map[string]interface{}{}

	result["file_paths"] = obj.FilePaths

	return result
}

func (s *DevopsTriggerResourceCrud) mapToFilter(fieldKeyFormat string) (oci_devops.Filter, error) {
	var baseObject oci_devops.Filter
	//discriminator
	triggerSourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trigger_source"))
	var triggerSource string
	if ok {
		triggerSource = triggerSourceRaw.(string)
	} else {
		triggerSource = "" // default value
	}
	switch strings.ToLower(triggerSource) {
	case strings.ToLower("BITBUCKET_CLOUD"):
		details := oci_devops.BitbucketCloudFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.BitbucketCloudFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_devops.BitbucketCloudFilterEventsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if exclude, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude")); ok {
			if tmpList := exclude.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude"), 0)
				tmp, err := s.mapToBitbucketCloudFilterExclusionAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert exclude, encountered error: %v", err)
				}
				details.Exclude = &tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToBitbucketCloudFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("BITBUCKET_SERVER"):
		details := oci_devops.BitbucketServerFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.BitbucketServerFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_devops.BitbucketServerFilterEventsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToBitbucketServerFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("DEVOPS_CODE_REPOSITORY"):
		details := oci_devops.DevopsCodeRepositoryFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.DevopsCodeRepositoryFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_devops.DevopsCodeRepositoryFilterEventsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if exclude, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude")); ok {
			if tmpList := exclude.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude"), 0)
				tmp, err := s.mapToDevopsCodeRepositoryFilterExclusionAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert exclude, encountered error: %v", err)
				}
				details.Exclude = &tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToDevopsCodeRepositoryFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("GITHUB"):
		details := oci_devops.GithubFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.GithubFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_devops.GithubFilterEventsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if exclude, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude")); ok {
			if tmpList := exclude.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude"), 0)
				tmp, err := s.mapToGithubFilterExclusionAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert exclude, encountered error: %v", err)
				}
				details.Exclude = &tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToGithubFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("GITLAB"):
		details := oci_devops.GitlabFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.GitlabFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_devops.GitlabFilterEventsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if exclude, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude")); ok {
			if tmpList := exclude.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude"), 0)
				tmp, err := s.mapToGitlabFilterExclusionAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert exclude, encountered error: %v", err)
				}
				details.Exclude = &tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToGitlabFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("GITLAB_SERVER"):
		details := oci_devops.GitlabServerFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.GitlabServerFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_devops.GitlabServerFilterEventsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToGitlabServerFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("VBS"):
		details := oci_devops.VbsFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.VbsFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_devops.VbsFilterEventsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if exclude, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude")); ok {
			if tmpList := exclude.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude"), 0)
				tmp, err := s.mapToVbsFilterExclusionAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert exclude, encountered error: %v", err)
				}
				details.Exclude = &tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToVbsFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown trigger_source '%v' was specified", triggerSource)
	}
	return baseObject, nil
}

func TriggerFilterToMap(obj *oci_devops.Filter) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.BitbucketCloudFilter:
		result["trigger_source"] = "BITBUCKET_CLOUD"

		result["events"] = v.Events

		if v.Exclude != nil {
			result["exclude"] = []interface{}{TriggerBitbucketCloudFilterExclusionAttributesToMap(v.Exclude)}
		}

		if v.Include != nil {
			result["include"] = []interface{}{TriggerBitbucketCloudFilterAttributesToMap(v.Include)}
		}
	case oci_devops.DevopsCodeRepositoryFilter:
		result["trigger_source"] = "DEVOPS_CODE_REPOSITORY"

		result["events"] = v.Events

		if v.Exclude != nil {
			result["exclude"] = []interface{}{TriggerDevopsCodeRepositoryFilterExclusionAttributesToMap(v.Exclude)}
		}

		if v.Include != nil {
			result["include"] = []interface{}{TriggerDevopsCodeRepositoryFilterAttributesToMap(v.Include)}
		}
	case oci_devops.GithubFilter:
		result["trigger_source"] = "GITHUB"

		result["events"] = v.Events

		if v.Exclude != nil {
			result["exclude"] = []interface{}{TriggerGithubFilterExclusionAttributesToMap(v.Exclude)}
		}

		if v.Include != nil {
			result["include"] = []interface{}{TriggerGithubFilterAttributesToMap(v.Include)}
		}
	case oci_devops.GitlabFilter:
		result["trigger_source"] = "GITLAB"

		result["events"] = v.Events

		if v.Exclude != nil {
			result["exclude"] = []interface{}{TriggerGitlabFilterExclusionAttributesToMap(v.Exclude)}
		}

		if v.Include != nil {
			result["include"] = []interface{}{TriggerGitlabFilterAttributesToMap(v.Include)}
		}
	case oci_devops.VbsFilter:
		result["trigger_source"] = "VBS"

		result["events"] = v.Events
		result["events"] = v.Events

		if v.Exclude != nil {
			result["exclude"] = []interface{}{TriggerVbsFilterExclusionAttributesToMap(v.Exclude)}
		}

		if v.Include != nil {
			result["include"] = []interface{}{TriggerVbsFilterAttributesToMap(v.Include)}
		}
	default:
		log.Printf("[WARN] Received 'trigger_source' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DevopsTriggerResourceCrud) mapToGithubFilterAttributes(fieldKeyFormat string) (oci_devops.GithubFilterAttributes, error) {
	result := oci_devops.GithubFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func TriggerGithubFilterAttributesToMap(obj *oci_devops.GithubFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseRef != nil {
		result["base_ref"] = string(*obj.BaseRef)
	}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{TriggerFileFilterToMap(obj.FileFilter)}
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	return result
}

func (s *DevopsTriggerResourceCrud) mapToGithubFilterExclusionAttributes(fieldKeyFormat string) (oci_devops.GithubFilterExclusionAttributes, error) {
	result := oci_devops.GithubFilterExclusionAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	return result, nil
}

func TriggerGithubFilterExclusionAttributesToMap(obj *oci_devops.GithubFilterExclusionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{TriggerFileFilterToMap(obj.FileFilter)}
	}

	return result
}

func (s *DevopsTriggerResourceCrud) mapToGitlabFilterAttributes(fieldKeyFormat string) (oci_devops.GitlabFilterAttributes, error) {
	result := oci_devops.GitlabFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func TriggerGitlabFilterAttributesToMap(obj *oci_devops.GitlabFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseRef != nil {
		result["base_ref"] = string(*obj.BaseRef)
	}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{TriggerFileFilterToMap(obj.FileFilter)}
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	return result
}

func (s *DevopsTriggerResourceCrud) mapToGitlabFilterExclusionAttributes(fieldKeyFormat string) (oci_devops.GitlabFilterExclusionAttributes, error) {
	result := oci_devops.GitlabFilterExclusionAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	return result, nil
}

func TriggerGitlabFilterExclusionAttributesToMap(obj *oci_devops.GitlabFilterExclusionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{TriggerFileFilterToMap(obj.FileFilter)}
	}

	return result
}

func (s *DevopsTriggerResourceCrud) mapToGitlabServerFilterAttributes(fieldKeyFormat string) (oci_devops.GitlabServerFilterAttributes, error) {
	result := oci_devops.GitlabServerFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func (s *DevopsTriggerResourceCrud) mapToTriggerAction(fieldKeyFormat string) (oci_devops.TriggerAction, error) {
	var baseObject oci_devops.TriggerAction
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("TRIGGER_BUILD_PIPELINE"):
		details := oci_devops.TriggerBuildPipelineAction{}
		if buildPipelineId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "build_pipeline_id")); ok {
			tmp := buildPipelineId.(string)
			details.BuildPipelineId = &tmp
		}
		if filter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filter")); ok {
			if tmpList := filter.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "filter"), 0)
				tmp, err := s.mapToFilter(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert filter, encountered error: %v", err)
				}
				details.Filter = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func TriggerActionToMap(obj oci_devops.TriggerAction) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_devops.TriggerBuildPipelineAction:
		result["type"] = "TRIGGER_BUILD_PIPELINE"

		if v.BuildPipelineId != nil {
			result["build_pipeline_id"] = string(*v.BuildPipelineId)
		}

		if v.Filter != nil {
			filterArray := []interface{}{}
			if filterMap := FilterToMap(&v.Filter); filterMap != nil {
				filterArray = append(filterArray, filterMap)
			}
			result["filter"] = filterArray
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func TriggerSummaryToMap(obj oci_devops.TriggerSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GetId() != nil {
		result["id"] = string(*obj.GetId())
	}

	if obj.GetCompartmentId() != nil {
		result["compartment_id"] = string(*obj.GetCompartmentId())
	}

	if obj.GetProjectId() != nil {
		result["project_id"] = string(*obj.GetProjectId())
	}

	if obj.GetDisplayName() != nil {
		result["display_name"] = string(*obj.GetDisplayName())
	}

	if obj.GetDescription() != nil {
		result["description"] = string(*obj.GetDescription())
	}

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}

	if obj.GetTimeUpdated() != nil {
		result["time_updated"] = obj.GetTimeUpdated().String()
	}

	result["state"] = obj.GetLifecycleState()

	if obj.GetLifecycleDetails() != nil {
		result["lifecycle_details"] = string(*obj.GetLifecycleDetails())
	}

	if obj.GetFreeformTags() != nil {
		result["freeform_tags"] = obj.GetFreeformTags()
	}

	if obj.GetSystemTags() != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.GetSystemTags())
	}

	if obj.GetDefinedTags() != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
	}
	switch v := (obj).(type) {
	case oci_devops.BitbucketCloudTriggerSummary:
		result["trigger_source"] = "BITBUCKET_CLOUD"

		if v.ConnectionId != nil {
			result["connection_id"] = string(*v.ConnectionId)
		}
	case oci_devops.BitbucketServerTriggerSummary:
		result["trigger_source"] = "BITBUCKET_SERVER"
	case oci_devops.DevopsCodeRepositoryTriggerSummary:
		result["trigger_source"] = "DEVOPS_CODE_REPOSITORY"

		if v.RepositoryId != nil {
			result["repository_id"] = string(*v.RepositoryId)
		}
	case oci_devops.GithubTriggerSummary:
		result["trigger_source"] = "GITHUB"

		if v.ConnectionId != nil {
			result["connection_id"] = string(*v.ConnectionId)
		}
	case oci_devops.GitlabTriggerSummary:
		result["trigger_source"] = "GITLAB"

		if v.ConnectionId != nil {
			result["connection_id"] = string(*v.ConnectionId)
		}
	case oci_devops.GitlabServerTriggerSummary:
		result["trigger_source"] = "GITLAB_SERVER"
	case oci_devops.VbsTriggerSummary:
		result["trigger_source"] = "VBS"

		if v.ConnectionId != nil {
			result["connection_id"] = string(*v.ConnectionId)
		}
	default:
		log.Printf("[WARN] Received 'trigger_source' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DevopsTriggerResourceCrud) mapToVbsFilterAttributes(fieldKeyFormat string) (oci_devops.VbsFilterAttributes, error) {
	result := oci_devops.VbsFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	if repositoryName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_name")); ok {
		tmp := repositoryName.(string)
		result.RepositoryName = &tmp
	}

	return result, nil
}

func TriggerVbsFilterAttributesToMap(obj *oci_devops.VbsFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseRef != nil {
		result["base_ref"] = string(*obj.BaseRef)
	}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	if obj.RepositoryName != nil {
		result["repository_name"] = string(*obj.RepositoryName)
	}

	return result
}

func (s *DevopsTriggerResourceCrud) mapToVbsFilterExclusionAttributes(fieldKeyFormat string) (oci_devops.VbsFilterExclusionAttributes, error) {
	result := oci_devops.VbsFilterExclusionAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	return result, nil
}

func TriggerVbsFilterExclusionAttributesToMap(obj *oci_devops.VbsFilterExclusionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	return result
}

func (s *DevopsTriggerResourceCrud) populateTopLevelPolymorphicCreateTriggerRequest(request *oci_devops.CreateTriggerRequest) error {
	//discriminator
	triggerSourceRaw, ok := s.D.GetOkExists("trigger_source")
	var triggerSource string
	if ok {
		triggerSource = triggerSourceRaw.(string)
	} else {
		triggerSource = "" // default value
	}
	switch strings.ToLower(triggerSource) {
	case strings.ToLower("BITBUCKET_CLOUD"):
		details := oci_devops.CreateBitbucketCloudTriggerDetails{}
		if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
			tmp := connectionId.(string)
			details.ConnectionId = &tmp
		}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		if projectId, ok := s.D.GetOkExists("project_id"); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		request.CreateTriggerDetails = details
	case strings.ToLower("BITBUCKET_SERVER"):
		details := oci_devops.CreateBitbucketServerTriggerDetails{}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		if projectId, ok := s.D.GetOkExists("project_id"); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		request.CreateTriggerDetails = details
	case strings.ToLower("DEVOPS_CODE_REPOSITORY"):
		details := oci_devops.CreateDevopsCodeRepositoryTriggerDetails{}
		if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
			tmp := repositoryId.(string)
			details.RepositoryId = &tmp
		}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		if projectId, ok := s.D.GetOkExists("project_id"); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		request.CreateTriggerDetails = details
	case strings.ToLower("GITHUB"):
		details := oci_devops.CreateGithubTriggerDetails{}
		if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
			tmp := connectionId.(string)
			details.ConnectionId = &tmp
		}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		if projectId, ok := s.D.GetOkExists("project_id"); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		request.CreateTriggerDetails = details
	case strings.ToLower("GITLAB"):
		details := oci_devops.CreateGitlabTriggerDetails{}
		if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
			tmp := connectionId.(string)
			details.ConnectionId = &tmp
		}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		if projectId, ok := s.D.GetOkExists("project_id"); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		request.CreateTriggerDetails = details
	case strings.ToLower("GITLAB_SERVER"):
		details := oci_devops.CreateGitlabServerTriggerDetails{}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		if projectId, ok := s.D.GetOkExists("project_id"); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		request.CreateTriggerDetails = details
	case strings.ToLower("VBS"):
		details := oci_devops.CreateVbsTriggerDetails{}
		if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
			tmp := connectionId.(string)
			details.ConnectionId = &tmp
		}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		if projectId, ok := s.D.GetOkExists("project_id"); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		request.CreateTriggerDetails = details
	default:
		return fmt.Errorf("unknown trigger_source '%v' was specified", triggerSource)
	}
	return nil
}

func (s *DevopsTriggerResourceCrud) populateTopLevelPolymorphicUpdateTriggerRequest(request *oci_devops.UpdateTriggerRequest) error {
	//discriminator
	triggerSourceRaw, ok := s.D.GetOkExists("trigger_source")
	var triggerSource string
	if ok {
		triggerSource = triggerSourceRaw.(string)
	} else {
		triggerSource = "" // default value
	}
	switch strings.ToLower(triggerSource) {
	case strings.ToLower("BITBUCKET_CLOUD"):
		details := oci_devops.UpdateBitbucketCloudTriggerDetails{}
		if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
			tmp := connectionId.(string)
			details.ConnectionId = &tmp
		}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		tmp := s.D.Id()
		request.TriggerId = &tmp
		request.UpdateTriggerDetails = details
	case strings.ToLower("BITBUCKET_SERVER"):
		details := oci_devops.UpdateBitbucketServerTriggerDetails{}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		tmp := s.D.Id()
		request.TriggerId = &tmp
		request.UpdateTriggerDetails = details
	case strings.ToLower("DEVOPS_CODE_REPOSITORY"):
		details := oci_devops.UpdateDevopsCodeRepositoryTriggerDetails{}
		if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
			tmp := repositoryId.(string)
			details.RepositoryId = &tmp
		}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		tmp := s.D.Id()
		request.TriggerId = &tmp
		request.UpdateTriggerDetails = details
	case strings.ToLower("GITHUB"):
		details := oci_devops.UpdateGithubTriggerDetails{}
		if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
			tmp := connectionId.(string)
			details.ConnectionId = &tmp
		}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		tmp := s.D.Id()
		request.TriggerId = &tmp
		request.UpdateTriggerDetails = details
	case strings.ToLower("GITLAB"):
		details := oci_devops.UpdateGitlabTriggerDetails{}
		if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
			tmp := connectionId.(string)
			details.ConnectionId = &tmp
		}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		tmp := s.D.Id()
		request.TriggerId = &tmp
		request.UpdateTriggerDetails = details
	case strings.ToLower("GITLAB_SERVER"):
		details := oci_devops.UpdateGitlabServerTriggerDetails{}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		tmp := s.D.Id()
		request.TriggerId = &tmp
		request.UpdateTriggerDetails = details
	case strings.ToLower("VBS"):
		details := oci_devops.UpdateVbsTriggerDetails{}
		if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
			tmp := connectionId.(string)
			details.ConnectionId = &tmp
		}
		if actions, ok := s.D.GetOkExists("actions"); ok {
			interfaces := actions.([]interface{})
			tmp := make([]oci_devops.TriggerAction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
				converted, err := s.mapToTriggerAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("actions") {
				details.Actions = tmp
			}
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
		tmp := s.D.Id()
		request.TriggerId = &tmp
		request.UpdateTriggerDetails = details
	default:
		return fmt.Errorf("unknown trigger_source '%v' was specified", triggerSource)
	}
	return nil
}
