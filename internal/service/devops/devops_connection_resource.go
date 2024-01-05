// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
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

func DevopsConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsConnection,
		Read:     readDevopsConnection,
		Update:   updateDevopsConnection,
		Delete:   deleteDevopsConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"connection_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"BITBUCKET_CLOUD_APP_PASSWORD",
					"BITBUCKET_SERVER_ACCESS_TOKEN",
					"GITHUB_ACCESS_TOKEN",
					"GITLAB_ACCESS_TOKEN",
					"GITLAB_SERVER_ACCESS_TOKEN",
					"VBS_ACCESS_TOKEN",
				}, true),
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"access_token": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"base_url": {
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
			"tls_verify_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"ca_certificate_bundle_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tls_verify_mode": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CA_CERTIFICATE_VERIFY",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_connection_validation_result": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"result": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_validated": {
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

func createDevopsConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DevopsConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.Connection
	DisableNotFoundRetries bool
}

func (s *DevopsConnectionResourceCrud) ID() string {
	connection := *s.Res
	return *connection.GetId()
}

func (s *DevopsConnectionResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DevopsConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_devops.ConnectionLifecycleStateActive),
	}
}

func (s *DevopsConnectionResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DevopsConnectionResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DevopsConnectionResourceCrud) Create() error {
	request := oci_devops.CreateConnectionRequest{}
	err := s.populateTopLevelPolymorphicCreateConnectionRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.CreateConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DevopsConnectionResourceCrud) getConnectionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_devops.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	connectionId, err := devopsConnectionWaitForWorkRequest(workId, "connection",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*connectionId)

	return s.Get()
}

func devopsConnectionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func devopsConnectionWaitForWorkRequest(wId *string, entityType string, action oci_devops.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_devops.DevopsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "devops")
	retryPolicy.ShouldRetryOperation = devopsConnectionWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDevopsConnectionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDevopsConnectionWorkRequest(client *oci_devops.DevopsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_devops.ActionTypeEnum) error {
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

func (s *DevopsConnectionResourceCrud) Get() error {
	request := oci_devops.GetConnectionRequest{}

	tmp := s.D.Id()
	request.ConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Connection
	return nil
}

func (s *DevopsConnectionResourceCrud) Update() error {
	request := oci_devops.UpdateConnectionRequest{}
	err := s.populateTopLevelPolymorphicUpdateConnectionRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DevopsConnectionResourceCrud) Delete() error {
	request := oci_devops.DeleteConnectionRequest{}

	tmp := s.D.Id()
	request.ConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.DeleteConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := devopsConnectionWaitForWorkRequest(workId, "connection",
		oci_devops.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DevopsConnectionResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_devops.BitbucketCloudAppPasswordConnection:
		s.D.Set("connection_type", "BITBUCKET_CLOUD_APP_PASSWORD")

		if v.AppPassword != nil {
			s.D.Set("app_password", *v.AppPassword)
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		if v.LastConnectionValidationResult != nil {
			s.D.Set("last_connection_validation_result", []interface{}{ConnectionValidationResultToMap(v.LastConnectionValidationResult)})
		} else {
			s.D.Set("last_connection_validation_result", nil)
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
	case oci_devops.BitbucketServerAccessTokenConnection:
		s.D.Set("connection_type", "BITBUCKET_SERVER_ACCESS_TOKEN")

		if v.AccessToken != nil {
			s.D.Set("access_token", *v.AccessToken)
		}

		if v.BaseUrl != nil {
			s.D.Set("base_url", *v.BaseUrl)
		}

		if v.TlsVerifyConfig != nil {
			tlsVerifyConfigArray := []interface{}{}
			if tlsVerifyConfigMap := TlsVerifyConfigToMap(&v.TlsVerifyConfig); tlsVerifyConfigMap != nil {
				tlsVerifyConfigArray = append(tlsVerifyConfigArray, tlsVerifyConfigMap)
			}
			s.D.Set("tls_verify_config", tlsVerifyConfigArray)
		} else {
			s.D.Set("tls_verify_config", nil)
		}

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

		if v.LastConnectionValidationResult != nil {
			s.D.Set("last_connection_validation_result", []interface{}{ConnectionValidationResultToMap(v.LastConnectionValidationResult)})
		} else {
			s.D.Set("last_connection_validation_result", nil)
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
	case oci_devops.GithubAccessTokenConnection:
		s.D.Set("connection_type", "GITHUB_ACCESS_TOKEN")

		if v.AccessToken != nil {
			s.D.Set("access_token", *v.AccessToken)
		}

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

		if v.LastConnectionValidationResult != nil {
			s.D.Set("last_connection_validation_result", []interface{}{ConnectionValidationResultToMap(v.LastConnectionValidationResult)})
		} else {
			s.D.Set("last_connection_validation_result", nil)
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
	case oci_devops.GitlabAccessTokenConnection:
		s.D.Set("connection_type", "GITLAB_ACCESS_TOKEN")

		if v.AccessToken != nil {
			s.D.Set("access_token", *v.AccessToken)
		}

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

		if v.LastConnectionValidationResult != nil {
			s.D.Set("last_connection_validation_result", []interface{}{ConnectionValidationResultToMap(v.LastConnectionValidationResult)})
		} else {
			s.D.Set("last_connection_validation_result", nil)
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
	case oci_devops.GitlabServerAccessTokenConnection:
		s.D.Set("connection_type", "GITLAB_SERVER_ACCESS_TOKEN")

		if v.AccessToken != nil {
			s.D.Set("access_token", *v.AccessToken)
		}

		if v.BaseUrl != nil {
			s.D.Set("base_url", *v.BaseUrl)
		}

		if v.TlsVerifyConfig != nil {
			tlsVerifyConfigArray := []interface{}{}
			if tlsVerifyConfigMap := TlsVerifyConfigToMap(&v.TlsVerifyConfig); tlsVerifyConfigMap != nil {
				tlsVerifyConfigArray = append(tlsVerifyConfigArray, tlsVerifyConfigMap)
			}
			s.D.Set("tls_verify_config", tlsVerifyConfigArray)
		} else {
			s.D.Set("tls_verify_config", nil)
		}

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

		if v.LastConnectionValidationResult != nil {
			s.D.Set("last_connection_validation_result", []interface{}{ConnectionValidationResultToMap(v.LastConnectionValidationResult)})
		} else {
			s.D.Set("last_connection_validation_result", nil)
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
	case oci_devops.VbsAccessTokenConnection:
		s.D.Set("connection_type", "VBS_ACCESS_TOKEN")

		if v.AccessToken != nil {
			s.D.Set("access_token", *v.AccessToken)
		}

		if v.BaseUrl != nil {
			s.D.Set("base_url", *v.BaseUrl)
		}

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

		if v.LastConnectionValidationResult != nil {
			s.D.Set("last_connection_validation_result", []interface{}{ConnectionValidationResultToMap(v.LastConnectionValidationResult)})
		} else {
			s.D.Set("last_connection_validation_result", nil)
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
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func devopsConnectionSummaryToMap(obj oci_devops.ConnectionSummary) map[string]interface{} {
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
	case oci_devops.BitbucketCloudAppPasswordConnectionSummary:
		result["connection_type"] = "BITBUCKET_CLOUD_APP_PASSWORD"

		if v.AppPassword != nil {
			result["app_password"] = string(*v.AppPassword)
		}

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_devops.BitbucketServerTokenConnectionSummary:
		result["connection_type"] = "BITBUCKET_SERVER_ACCESS_TOKEN"

		if v.AccessToken != nil {
			result["access_token"] = string(*v.AccessToken)
		}

		if v.BaseUrl != nil {
			result["base_url"] = string(*v.BaseUrl)
		}

		if v.TlsVerifyConfig != nil {
			tlsVerifyConfigArray := []interface{}{}
			if tlsVerifyConfigMap := TlsVerifyConfigToMap(&v.TlsVerifyConfig); tlsVerifyConfigMap != nil {
				tlsVerifyConfigArray = append(tlsVerifyConfigArray, tlsVerifyConfigMap)
			}
			result["tls_verify_config"] = tlsVerifyConfigArray
		}
	case oci_devops.GithubAccessTokenConnectionSummary:
		result["connection_type"] = "GITHUB_ACCESS_TOKEN"

		if v.AccessToken != nil {
			result["access_token"] = string(*v.AccessToken)
		}
	case oci_devops.GitlabAccessTokenConnectionSummary:
		result["connection_type"] = "GITLAB_ACCESS_TOKEN"

		if v.AccessToken != nil {
			result["access_token"] = string(*v.AccessToken)
		}
	case oci_devops.GitlabServerAccessTokenConnectionSummary:
		result["connection_type"] = "GITLAB_SERVER_ACCESS_TOKEN"

		if v.AccessToken != nil {
			result["access_token"] = string(*v.AccessToken)
		}

		if v.BaseUrl != nil {
			result["base_url"] = string(*v.BaseUrl)
		}

		if v.TlsVerifyConfig != nil {
			tlsVerifyConfigArray := []interface{}{}
			if tlsVerifyConfigMap := TlsVerifyConfigToMap(&v.TlsVerifyConfig); tlsVerifyConfigMap != nil {
				tlsVerifyConfigArray = append(tlsVerifyConfigArray, tlsVerifyConfigMap)
			}
			result["tls_verify_config"] = tlsVerifyConfigArray
		}
	case oci_devops.VbsAccessTokenConnectionSummary:
		result["connection_type"] = "VBS_ACCESS_TOKEN"

		if v.AccessToken != nil {
			result["access_token"] = string(*v.AccessToken)
		}

		if v.BaseUrl != nil {
			result["base_url"] = string(*v.BaseUrl)
		}
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func ConnectionValidationResultToMap(obj *oci_devops.ConnectionValidationResult) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	result["result"] = string(obj.Result)

	if obj.TimeValidated != nil {
		result["time_validated"] = obj.TimeValidated.String()
	}

	return result
}

func (s *DevopsConnectionResourceCrud) mapToTlsVerifyConfig(fieldKeyFormat string) (oci_devops.TlsVerifyConfig, error) {
	var baseObject oci_devops.TlsVerifyConfig
	//discriminator
	tlsVerifyModeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tls_verify_mode"))
	var tlsVerifyMode string
	if ok {
		tlsVerifyMode = tlsVerifyModeRaw.(string)
	} else {
		tlsVerifyMode = "" // default value
	}
	switch strings.ToLower(tlsVerifyMode) {
	case strings.ToLower("CA_CERTIFICATE_VERIFY"):
		details := oci_devops.CaCertVerify{}
		if caCertificateBundleId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ca_certificate_bundle_id")); ok {
			tmp := caCertificateBundleId.(string)
			details.CaCertificateBundleId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown tls_verify_mode '%v' was specified", tlsVerifyMode)
	}
	return baseObject, nil
}

func TlsVerifyConfigToMap(obj *oci_devops.TlsVerifyConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.CaCertVerify:
		result["tls_verify_mode"] = "CA_CERTIFICATE_VERIFY"

		if v.CaCertificateBundleId != nil {
			result["ca_certificate_bundle_id"] = string(*v.CaCertificateBundleId)
		}
	default:
		log.Printf("[WARN] Received 'tls_verify_mode' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DevopsConnectionResourceCrud) populateTopLevelPolymorphicCreateConnectionRequest(request *oci_devops.CreateConnectionRequest) error {
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists("connection_type")
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("BITBUCKET_CLOUD_APP_PASSWORD"):
		details := oci_devops.CreateBitbucketCloudAppPasswordConnectionDetails{}
		if appPassword, ok := s.D.GetOkExists("app_password"); ok {
			tmp := appPassword.(string)
			details.AppPassword = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		request.CreateConnectionDetails = details
	case strings.ToLower("BITBUCKET_SERVER_ACCESS_TOKEN"):
		details := oci_devops.CreateBitbucketServerAccessTokenConnectionDetails{}
		if accessToken, ok := s.D.GetOkExists("access_token"); ok {
			tmp := accessToken.(string)
			details.AccessToken = &tmp
		}
		if baseUrl, ok := s.D.GetOkExists("base_url"); ok {
			tmp := baseUrl.(string)
			details.BaseUrl = &tmp
		}
		if tlsVerifyConfig, ok := s.D.GetOkExists("tls_verify_config"); ok {
			if tmpList := tlsVerifyConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tls_verify_config", 0)
				tmp, err := s.mapToTlsVerifyConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TlsVerifyConfig = tmp
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
		request.CreateConnectionDetails = details
	case strings.ToLower("GITHUB_ACCESS_TOKEN"):
		details := oci_devops.CreateGithubAccessTokenConnectionDetails{}
		if accessToken, ok := s.D.GetOkExists("access_token"); ok {
			tmp := accessToken.(string)
			details.AccessToken = &tmp
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
		request.CreateConnectionDetails = details
	case strings.ToLower("GITLAB_ACCESS_TOKEN"):
		details := oci_devops.CreateGitlabAccessTokenConnectionDetails{}
		if accessToken, ok := s.D.GetOkExists("access_token"); ok {
			tmp := accessToken.(string)
			details.AccessToken = &tmp
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
		request.CreateConnectionDetails = details
	case strings.ToLower("GITLAB_SERVER_ACCESS_TOKEN"):
		details := oci_devops.CreateGitlabServerAccessTokenConnectionDetails{}
		if accessToken, ok := s.D.GetOkExists("access_token"); ok {
			tmp := accessToken.(string)
			details.AccessToken = &tmp
		}
		if baseUrl, ok := s.D.GetOkExists("base_url"); ok {
			tmp := baseUrl.(string)
			details.BaseUrl = &tmp
		}
		if tlsVerifyConfig, ok := s.D.GetOkExists("tls_verify_config"); ok {
			if tmpList := tlsVerifyConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tls_verify_config", 0)
				tmp, err := s.mapToTlsVerifyConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TlsVerifyConfig = tmp
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
		request.CreateConnectionDetails = details
	case strings.ToLower("VBS_ACCESS_TOKEN"):
		details := oci_devops.CreateVbsAccessTokenConnectionDetails{}
		if accessToken, ok := s.D.GetOkExists("access_token"); ok {
			tmp := accessToken.(string)
			details.AccessToken = &tmp
		}
		if baseUrl, ok := s.D.GetOkExists("base_url"); ok {
			tmp := baseUrl.(string)
			details.BaseUrl = &tmp
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
		request.CreateConnectionDetails = details
	default:
		return fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return nil
}

func (s *DevopsConnectionResourceCrud) populateTopLevelPolymorphicUpdateConnectionRequest(request *oci_devops.UpdateConnectionRequest) error {
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists("connection_type")
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("BITBUCKET_CLOUD_APP_PASSWORD"):
		details := oci_devops.UpdateBitbucketCloudAppPasswordConnectionDetails{}
		if appPassword, ok := s.D.GetOkExists("app_password"); ok {
			tmp := appPassword.(string)
			details.AppPassword = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
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
		request.UpdateConnectionDetails = details
	case strings.ToLower("BITBUCKET_SERVER_ACCESS_TOKEN"):
		details := oci_devops.UpdateBitbucketServerAccessTokenConnectionDetails{}
		if accessToken, ok := s.D.GetOkExists("access_token"); ok {
			tmp := accessToken.(string)
			details.AccessToken = &tmp
		}
		if baseUrl, ok := s.D.GetOkExists("base_url"); ok {
			tmp := baseUrl.(string)
			details.BaseUrl = &tmp
		}
		if tlsVerifyConfig, ok := s.D.GetOkExists("tls_verify_config"); ok {
			if tmpList := tlsVerifyConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tls_verify_config", 0)
				tmp, err := s.mapToTlsVerifyConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TlsVerifyConfig = tmp
			}
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
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
		request.UpdateConnectionDetails = details
	case strings.ToLower("GITHUB_ACCESS_TOKEN"):
		details := oci_devops.UpdateGithubAccessTokenConnectionDetails{}
		if accessToken, ok := s.D.GetOkExists("access_token"); ok {
			tmp := accessToken.(string)
			details.AccessToken = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
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
		request.UpdateConnectionDetails = details
	case strings.ToLower("GITLAB_ACCESS_TOKEN"):
		details := oci_devops.UpdateGitlabAccessTokenConnectionDetails{}
		if accessToken, ok := s.D.GetOkExists("access_token"); ok {
			tmp := accessToken.(string)
			details.AccessToken = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
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
		request.UpdateConnectionDetails = details
	case strings.ToLower("GITLAB_SERVER_ACCESS_TOKEN"):
		details := oci_devops.UpdateGitlabServerAccessTokenConnectionDetails{}
		if accessToken, ok := s.D.GetOkExists("access_token"); ok {
			tmp := accessToken.(string)
			details.AccessToken = &tmp
		}
		if baseUrl, ok := s.D.GetOkExists("base_url"); ok {
			tmp := baseUrl.(string)
			details.BaseUrl = &tmp
		}
		if tlsVerifyConfig, ok := s.D.GetOkExists("tls_verify_config"); ok {
			if tmpList := tlsVerifyConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tls_verify_config", 0)
				tmp, err := s.mapToTlsVerifyConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TlsVerifyConfig = tmp
			}
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
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
		request.UpdateConnectionDetails = details
	case strings.ToLower("VBS_ACCESS_TOKEN"):
		details := oci_devops.UpdateVbsAccessTokenConnectionDetails{}
		if accessToken, ok := s.D.GetOkExists("access_token"); ok {
			tmp := accessToken.(string)
			details.AccessToken = &tmp
		}
		if baseUrl, ok := s.D.GetOkExists("base_url"); ok {
			tmp := baseUrl.(string)
			details.BaseUrl = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
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
		request.UpdateConnectionDetails = details
	default:
		return fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return nil
}
