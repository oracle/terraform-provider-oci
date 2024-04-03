// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
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
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
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
			"availability_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_source_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"allow_metrics": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connection_timeout": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_daemon_set": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric_dimensions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"proxy_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"read_data_limit": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"read_timeout": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"schedule_mins": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"data_source_summary_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_daemon_set": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"host": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_id": {
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
			"install_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_customer_deployed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_agent_properties": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"units": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"values": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
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
						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
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
						"plugin_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plugin_status_message": {
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
			"resource_artifact_version": {
				Type:     schema.TypeString,
				Computed: true,
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
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.CreateResource(d, sync)
}

func readManagementAgentManagementAgent(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

func updateManagementAgentManagementAgent(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteManagementAgentManagementAgent(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()
	sync.DisableNotFoundRetries = true

	//Destroy does not remove the item from the Management Agent backend.. only removes from state
	sync.VoidState()
	return nil
}

type ManagementAgentManagementAgentResourceCrud struct {
	tfresource.BaseCrud
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
		if tfresource.ShouldRetry(response, false, "management_agent", startTime) {
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
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "management_agent")
	retryPolicy.ShouldRetryOperation = managementAgentWorkRequestShouldRetryFunc(timeout)

	response := oci_management_agent.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_management_agent.OperationStatusInProgress),
			string(oci_management_agent.OperationStatusAccepted),
			string(oci_management_agent.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_management_agent.OperationStatusSucceeded),
			string(oci_management_agent.OperationStatusFailed),
			string(oci_management_agent.OperationStatusCanceled),
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

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_management_agent.OperationStatusFailed || response.Status == oci_management_agent.OperationStatusCanceled {
		return nil, getErrorFromManagementAgentManagementAgentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

// Function to wait for a compute instance "host_id" (or instance ocid) to appear in agent list
// When a compute is created with OCA plugin, it can take up to 10 minutes for the agent to appear in MACS
// This function waits for the agent to appear.  Used with data source for "host_id" and "wait_for_host_id"
func managementAgentWaitForInstanceAgent(hostId *string, compartmentId *string,
	timeout time.Duration, disableFoundRetries bool, client *oci_management_agent.ManagementAgentClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "management_agent")
	retryPolicy.ShouldRetryOperation = managementAgentWorkRequestShouldRetryFunc(timeout)

	response := oci_management_agent.ListManagementAgentsResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string("not found"),
		},
		Target: []string{
			string("found"),
		},
		Refresh: func() (interface{}, string, error) {
			request := oci_management_agent.ListManagementAgentsRequest{}
			request.CompartmentId = compartmentId
			request.HostId = hostId
			request.RequestMetadata = oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			}

			response, err := client.ListManagementAgents(context.Background(), request)

			var state string
			if len(response.Items) > 0 {
				state = "found"
			} else {
				state = "not found"
			}

			return response, state, err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	if len(response.Items) > 0 {
		return response.Items[0].Id, nil
	} else {
		return nil, nil
	}

}

func getErrorFromManagementAgentManagementAgentWorkRequest(client *oci_management_agent.ManagementAgentClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_management_agent.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_management_agent.ListWorkRequestErrorsRequest{
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

func (s *ManagementAgentManagementAgentResourceCrud) Get() error {
	request := oci_management_agent.GetManagementAgentRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_agent_id"); ok {
		tmp := managedInstanceId.(string)
		s.D.SetId(tmp)
	}

	managedInstanceId := s.D.Id() // For import case
	request.ManagementAgentId = &managedInstanceId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

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

	/*if isAgentAutoUpgradable, ok := s.D.GetOkExists("is_agent_auto_upgradable"); ok {
		tmp := isAgentAutoUpgradable.(bool)
		request.IsAgentAutoUpgradable = &tmp
	}*/

	agentId := s.D.Id()
	request.ManagementAgentId = &agentId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.UpdateManagementAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementAgent

	if deployPluginIds, ok := s.D.GetOkExists("deploy_plugins_id"); ok {
		interfaces := deployPluginIds.([]interface{})
		var deployPluginsId []string
		for i := range interfaces {
			if interfaces[i] != nil {
				if !contains(s.Res.PluginList, interfaces[i].(string)) {
					deployPluginsId = append(deployPluginsId, interfaces[i].(string))
				} else {
					log.Printf("[INFO] deploy plugin %v already deployed on agent %v", interfaces[i].(string), agentId)
				}
			}
		}

		if len(deployPluginsId) != 0 && s.D.HasChange("deploy_plugins_id") {
			comparmtentId := response.CompartmentId
			if err = s.deployPlugin(deployPluginsId, agentId, *comparmtentId); err != nil {
				return err
			}
		}
	}
	return nil
}

func contains(s []oci_management_agent.ManagementAgentPluginDetails, str string) bool {
	for _, v := range s {
		if *v.PluginId == str {
			return true
		}
	}
	return false
}
func (s *ManagementAgentManagementAgentResourceCrud) Delete() error {
	request := oci_management_agent.DeleteManagementAgentRequest{}

	tmp := s.D.Id()
	request.ManagementAgentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	_, err := s.Client.DeleteManagementAgent(context.Background(), request)
	return err
}

func (s *ManagementAgentManagementAgentResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)
	s.D.Set("availability_status", s.Res.AvailabilityStatus)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	dataSourceList := []interface{}{}
	for _, item := range s.Res.DataSourceList {
		dataSourceList = append(dataSourceList, DataSourceToMap(item))
	}
	s.D.Set("data_source_list", dataSourceList)

	s.D.Set("data_source_summary_list", []string{})

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Host != nil {
		s.D.Set("host", *s.Res.Host)
	}

	if s.Res.HostId != nil {
		s.D.Set("host_id", *s.Res.HostId)
	}

	if s.Res.InstallKeyId != nil {
		s.D.Set("install_key_id", *s.Res.InstallKeyId)
	}

	if s.Res.InstallPath != nil {
		s.D.Set("install_path", *s.Res.InstallPath)
	}

	s.D.Set("install_type", s.Res.InstallType)

	if s.Res.IsAgentAutoUpgradable != nil {
		s.D.Set("is_agent_auto_upgradable", *s.Res.IsAgentAutoUpgradable)
	}

	if s.Res.IsCustomerDeployed != nil {
		s.D.Set("is_customer_deployed", *s.Res.IsCustomerDeployed)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	managementAgentProperties := []interface{}{}
	for _, item := range s.Res.ManagementAgentProperties {
		managementAgentProperties = append(managementAgentProperties, ManagementAgentPropertyToMap(item))
	}
	s.D.Set("management_agent_properties", managementAgentProperties)

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

	if s.Res.ResourceArtifactVersion != nil {
		s.D.Set("resource_artifact_version", *s.Res.ResourceArtifactVersion)
	}

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

func DataSourceToMap(obj oci_management_agent.DataSource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_management_agent.KubernetesClusterDataSource:
		result["type"] = "KUBERNETES_CLUSTER"

		if v.IsDaemonSet != nil {
			result["is_daemon_set"] = bool(*v.IsDaemonSet)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.Key != nil {
			result["key"] = string(*v.Key)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		result["state"] = string(v.State)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}
	case oci_management_agent.PrometheusEmitterDataSource:
		result["type"] = "PROMETHEUS_EMITTER"

		if v.AllowMetrics != nil {
			result["allow_metrics"] = string(*v.AllowMetrics)
		}

		if v.ConnectionTimeout != nil {
			result["connection_timeout"] = int(*v.ConnectionTimeout)
		}

		metricDimensions := []interface{}{}
		for _, item := range v.MetricDimensions {
			metricDimensions = append(metricDimensions, MetricDimensionToMap(item))
		}
		result["metric_dimensions"] = metricDimensions

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.ProxyUrl != nil {
			result["proxy_url"] = string(*v.ProxyUrl)
		}

		if v.ReadDataLimit != nil {
			result["read_data_limit"] = int(*v.ReadDataLimit)
		}

		if v.ReadTimeout != nil {
			result["read_timeout"] = int(*v.ReadTimeout)
		}

		if v.ResourceGroup != nil {
			result["resource_group"] = string(*v.ResourceGroup)
		}

		if v.ScheduleMins != nil {
			result["schedule_mins"] = int(*v.ScheduleMins)
		}

		if v.Url != nil {
			result["url"] = string(*v.Url)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.Key != nil {
			result["key"] = string(*v.Key)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		result["state"] = string(v.State)

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func DataSourceSummaryItemToMap(obj oci_management_agent.DataSourceSummaryItem) map[string]interface{} {
	result := map[string]interface{}{}
	if obj.GetKey() != nil {
		result["key"] = string(*obj.GetKey())
	}
	if obj.GetName() != nil {
		result["name"] = string(*obj.GetName())
	}
	switch v := (obj).(type) {
	case oci_management_agent.KubernetesClusterDataSourceSummaryItem:
		result["type"] = "KUBERNETES_CLUSTER"

		if v.IsDaemonSet != nil {
			result["is_daemon_set"] = bool(*v.IsDaemonSet)
		}
	case oci_management_agent.PrometheusEmitterDataSourceSummaryItem:
		result["type"] = "PROMETHEUS_EMITTER"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func ManagementAgentPluginDetailsToMap(obj oci_management_agent.ManagementAgentPluginDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.PluginDisplayName != nil {
		result["plugin_display_name"] = string(*obj.PluginDisplayName)
	}

	if obj.PluginId != nil {
		result["plugin_id"] = string(*obj.PluginId)
	}

	if obj.PluginName != nil {
		result["plugin_name"] = string(*obj.PluginName)
	}

	result["plugin_status"] = string(obj.PluginStatus)

	if obj.PluginStatusMessage != nil {
		result["plugin_status_message"] = string(*obj.PluginStatusMessage)
	}

	if obj.PluginVersion != nil {
		result["plugin_version"] = string(*obj.PluginVersion)
	}

	return result
}

func ManagementAgentPropertyToMap(obj oci_management_agent.ManagementAgentProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["units"] = string(obj.Units)

	result["values"] = obj.Values

	return result
}

func (s *ManagementAgentManagementAgentResourceCrud) mapToMetricDimension(fieldKeyFormat string) (oci_management_agent.MetricDimension, error) {
	result := oci_management_agent.MetricDimension{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func (s *ManagementAgentManagementAgentResourceCrud) Create() error {
	e := s.Get()
	if e != nil {
		return e
	}

	if e := s.SetData(); e != nil {
		return e
	}

	return s.Update()
}

func (s *ManagementAgentManagementAgentResourceCrud) deployPlugin(pluginIds []string, agentId string, compartmentId string) error {
	request := oci_management_agent.DeployPluginsRequest{}
	request.AgentIds = append(request.AgentIds, agentId)
	request.AgentCompartmentId = &compartmentId
	request.PluginIds = pluginIds

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.DeployPlugins(context.Background(), request)
	if err != nil {
		return err
	}
	workRequestId := response.OpcWorkRequestId

	_, workRequestErr := managementAgentWaitForWorkRequest(workRequestId, "managementagent",
		oci_management_agent.ActionTypesUpdated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return workRequestErr
}
