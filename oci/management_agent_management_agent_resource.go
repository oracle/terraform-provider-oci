// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v27/common"
	oci_management_agent "github.com/oracle/oci-go-sdk/v27/managementagent"
)

func init() {
	RegisterResource("oci_management_agent_management_agent", ManagementAgentManagementAgentResource())
}

func ManagementAgentManagementAgentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createManagementAgentManagementAgent,
		Read:     readManagementAgentManagementAgent,
		Update:   updateManagementAgentManagementAgent,
		Delete:   deleteManagementAgentManagementAgent,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_agent_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
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
			"is_agent_auto_upgradable": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"deploy_plugins_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"install_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"install_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plugin_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"plugin_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plugin_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plugin_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plugin_version": {
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_heartbeat": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createManagementAgentManagementAgent(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).managementAgentClient()

	return CreateResource(d, sync)
}

func readManagementAgentManagementAgent(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).managementAgentClient()

	return ReadResource(sync)
}

func updateManagementAgentManagementAgent(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).managementAgentClient()

	return UpdateResource(d, sync)
}

func deleteManagementAgentManagementAgent(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).managementAgentClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type ManagementAgentManagementAgentResourceCrud struct {
	BaseCrud
	Client                 *oci_management_agent.ManagementAgentClient
	Res                    *oci_management_agent.ManagementAgent
	DisableNotFoundRetries bool
}

func (s *ManagementAgentManagementAgentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ManagementAgentManagementAgentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesCreating),
	}
}

func (s *ManagementAgentManagementAgentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesActive),
	}
}

func (s *ManagementAgentManagementAgentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesDeleting),
	}
}

func (s *ManagementAgentManagementAgentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesTerminated),
		string(oci_management_agent.LifecycleStatesDeleted),
	}
}

func (s *ManagementAgentManagementAgentResourceCrud) getManagementAgentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_management_agent.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	managementAgentId, err := managementAgentWaitForWorkRequest(workId, "managementagent",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, managementAgentId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_management_agent.DeleteWorkRequestRequest{
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
	s.D.SetId(*managementAgentId)

	return s.Get()
}

func managementAgentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "management_agent", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_management_agent.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func managementAgentWaitForWorkRequest(wId *string, entityType string, action oci_management_agent.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_management_agent.ManagementAgentClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "management_agent")
	retryPolicy.ShouldRetryOperation = managementAgentWorkRequestShouldRetryFunc(timeout)

	response := oci_management_agent.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_management_agent.WorkDeliveryStatusInProgress),
			string(oci_management_agent.WorkDeliveryStatusAccepted),
			string(oci_management_agent.WorkDeliveryStatusCanceling),
		},
		Target: []string{
			string(oci_management_agent.WorkDeliveryStatusSucceeded),
			string(oci_management_agent.WorkDeliveryStatusFailed),
			string(oci_management_agent.WorkDeliveryStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_management_agent.GetWorkRequestRequest{
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

	// The workrequest didn't do all its intended tasks, if the errors is set; so we should check for it
	var workRequestErr error
	if identifier == nil {
		errorMessage := getErrorFromManagementAgentWorkRequest(client, wId, retryPolicy, entityType, action)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)
	}

	return identifier, workRequestErr
}

func getErrorFromManagementAgentWorkRequest(client *oci_management_agent.ManagementAgentClient, wId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_management_agent.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_management_agent.ListWorkRequestErrorsRequest{
			WorkRequestId: wId,
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

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *ManagementAgentManagementAgentResourceCrud) Get() error {
	request := oci_management_agent.GetManagementAgentRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_agent_id"); ok {
		tmp := managedInstanceId.(string)
		s.D.SetId(tmp)
	}

	managedInstanceId := s.D.Id() // For import case
	request.ManagementAgentId = &managedInstanceId
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.GetManagementAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementAgent
	return nil
}

func (s *ManagementAgentManagementAgentResourceCrud) Update() error {
	request := oci_management_agent.UpdateManagementAgentRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAgentAutoUpgradable, ok := s.D.GetOkExists("is_agent_auto_upgradable"); ok {
		tmp := isAgentAutoUpgradable.(bool)
		request.IsAgentAutoUpgradable = &tmp
	}

	agentId := s.D.Id()
	request.ManagementAgentId = &agentId

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.UpdateManagementAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementAgent

	if deployPluginIds, ok := s.D.GetOkExists("deploy_plugins_id"); ok {
		interfaces := deployPluginIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("deploy_plugins_id") {
			comparmtentId := response.CompartmentId
			if err = s.deployPlugin(tmp, agentId, *comparmtentId); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *ManagementAgentManagementAgentResourceCrud) Delete() error {
	request := oci_management_agent.DeleteManagementAgentRequest{}

	tmp := s.D.Id()
	request.ManagementAgentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	_, err := s.Client.DeleteManagementAgent(context.Background(), request)
	return err
}

func (s *ManagementAgentManagementAgentResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Host != nil {
		s.D.Set("host", *s.Res.Host)
	}

	if s.Res.InstallKeyId != nil {
		s.D.Set("install_key_id", *s.Res.InstallKeyId)
	}

	if s.Res.InstallPath != nil {
		s.D.Set("install_path", *s.Res.InstallPath)
	}

	if s.Res.IsAgentAutoUpgradable != nil {
		s.D.Set("is_agent_auto_upgradable", *s.Res.IsAgentAutoUpgradable)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PlatformName != nil {
		s.D.Set("platform_name", *s.Res.PlatformName)
	}

	s.D.Set("platform_type", s.Res.PlatformType)

	if s.Res.PlatformVersion != nil {
		s.D.Set("platform_version", *s.Res.PlatformVersion)
	}

	pluginList := []interface{}{}
	for _, item := range s.Res.PluginList {
		pluginList = append(pluginList, ManagementAgentPluginDetailsToMap(item))
	}
	s.D.Set("plugin_list", pluginList)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastHeartbeat != nil {
		s.D.Set("time_last_heartbeat", s.Res.TimeLastHeartbeat.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func ManagementAgentPluginDetailsToMap(obj oci_management_agent.ManagementAgentPluginDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PluginDisplayName != nil {
		result["plugin_display_name"] = string(*obj.PluginDisplayName)
	}

	if obj.PluginId != nil {
		result["plugin_id"] = string(*obj.PluginId)
	}

	if obj.PluginName != nil {
		result["plugin_name"] = string(*obj.PluginName)
	}

	if obj.PluginVersion != nil {
		result["plugin_version"] = string(*obj.PluginVersion)
	}

	return result
}

func (s *ManagementAgentManagementAgentResourceCrud) Create() error {
	e := s.Get()
	if e != nil {
		return e
	}

	return s.Update()
}

func (s *ManagementAgentManagementAgentResourceCrud) deployPlugin(pluginIds []string, agentId string, compartmentId string) error {
	request := oci_management_agent.DeployPluginsRequest{}
	request.AgentIds = append(request.AgentIds, agentId)
	request.AgentCompartmentId = &compartmentId
	request.PluginIds = pluginIds

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.DeployPlugins(context.Background(), request)
	if err != nil {
		return err
	}
	workRequestId := response.OpcWorkRequestId

	_, workRequestErr := managementAgentWaitForWorkRequest(workRequestId, "managementagent",
		oci_management_agent.ActionTypesUpdated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return workRequestErr
}
