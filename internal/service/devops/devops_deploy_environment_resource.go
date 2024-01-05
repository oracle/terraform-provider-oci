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

func DevopsDeployEnvironmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsDeployEnvironment,
		Read:     readDevopsDeployEnvironment,
		Update:   updateDevopsDeployEnvironment,
		Delete:   deleteDevopsDeployEnvironment,
		Schema: map[string]*schema.Schema{
			// Required
			"deploy_environment_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"COMPUTE_INSTANCE_GROUP",
					"FUNCTION",
					"OKE_CLUSTER",
				}, true),
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compute_instance_group_selectors": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"selector_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"INSTANCE_IDS",
											"INSTANCE_QUERY",
										}, true),
									},

									// Optional
									"compute_instance_ids": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"query": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"region": {
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
			"function_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_channel": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"network_channel_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"PRIVATE_ENDPOINT_CHANNEL",
								"SERVICE_VNIC_CHANNEL",
							}, true),
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"nsg_ids": {
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

func createDevopsDeployEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsDeployEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsDeployEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsDeployEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DevopsDeployEnvironmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.DeployEnvironment
	DisableNotFoundRetries bool
}

func (s *DevopsDeployEnvironmentResourceCrud) ID() string {
	deployEnvironment := *s.Res
	return *deployEnvironment.GetId()
}

func (s *DevopsDeployEnvironmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_devops.DeployEnvironmentLifecycleStateCreating),
	}
}

func (s *DevopsDeployEnvironmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_devops.DeployEnvironmentLifecycleStateActive),
		string(oci_devops.DeployEnvironmentLifecycleStateNeedsAttention),
	}
}

func (s *DevopsDeployEnvironmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_devops.DeployEnvironmentLifecycleStateDeleting),
	}
}

func (s *DevopsDeployEnvironmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_devops.DeployEnvironmentLifecycleStateDeleted),
	}
}

func (s *DevopsDeployEnvironmentResourceCrud) Create() error {
	request := oci_devops.CreateDeployEnvironmentRequest{}
	err := s.populateTopLevelPolymorphicCreateDeployEnvironmentRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.CreateDeployEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDeployEnvironmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DevopsDeployEnvironmentResourceCrud) getDeployEnvironmentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_devops.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	deployEnvironmentId, err := deployEnvironmentWaitForWorkRequest(workId, "environment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*deployEnvironmentId)

	return s.Get()
}

func deployEnvironmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func deployEnvironmentWaitForWorkRequest(wId *string, entityType string, action oci_devops.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_devops.DevopsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "devops")
	retryPolicy.ShouldRetryOperation = deployEnvironmentWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDevopsDeployEnvironmentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDevopsDeployEnvironmentWorkRequest(client *oci_devops.DevopsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_devops.ActionTypeEnum) error {
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

func (s *DevopsDeployEnvironmentResourceCrud) Get() error {
	request := oci_devops.GetDeployEnvironmentRequest{}

	tmp := s.D.Id()
	request.DeployEnvironmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetDeployEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DeployEnvironment
	return nil
}

func (s *DevopsDeployEnvironmentResourceCrud) Update() error {
	request := oci_devops.UpdateDeployEnvironmentRequest{}
	err := s.populateTopLevelPolymorphicUpdateDeployEnvironmentRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateDeployEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeployEnvironmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DevopsDeployEnvironmentResourceCrud) Delete() error {
	request := oci_devops.DeleteDeployEnvironmentRequest{}

	tmp := s.D.Id()
	request.DeployEnvironmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.DeleteDeployEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := deployEnvironmentWaitForWorkRequest(workId, "environment",
		oci_devops.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DevopsDeployEnvironmentResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_devops.ComputeInstanceGroupDeployEnvironment:
		s.D.Set("deploy_environment_type", "COMPUTE_INSTANCE_GROUP")

		if v.ComputeInstanceGroupSelectors != nil {
			s.D.Set("compute_instance_group_selectors", []interface{}{ComputeInstanceGroupSelectorCollectionToMap(v.ComputeInstanceGroupSelectors)})
		} else {
			s.D.Set("compute_instance_group_selectors", nil)
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
	case oci_devops.FunctionDeployEnvironment:
		s.D.Set("deploy_environment_type", "FUNCTION")

		if v.FunctionId != nil {
			s.D.Set("function_id", *v.FunctionId)
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
	case oci_devops.OkeClusterDeployEnvironment:
		s.D.Set("deploy_environment_type", "OKE_CLUSTER")

		if v.ClusterId != nil {
			s.D.Set("cluster_id", *v.ClusterId)
		}

		if v.NetworkChannel != nil {
			networkChannelArray := []interface{}{}
			if networkChannelMap := NetworkChannelToMap(&v.NetworkChannel); networkChannelMap != nil {
				networkChannelArray = append(networkChannelArray, networkChannelMap)
			}
			s.D.Set("network_channel", networkChannelArray)
		} else {
			s.D.Set("network_channel", nil)
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
		log.Printf("[WARN] Received 'deploy_environment_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *DevopsDeployEnvironmentResourceCrud) mapToComputeInstanceGroupSelector(fieldKeyFormat string) (oci_devops.ComputeInstanceGroupSelector, error) {
	var baseObject oci_devops.ComputeInstanceGroupSelector
	//discriminator
	selectorTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selector_type"))
	var selectorType string
	if ok {
		selectorType = selectorTypeRaw.(string)
	} else {
		selectorType = "" // default value
	}
	switch strings.ToLower(selectorType) {
	case strings.ToLower("INSTANCE_IDS"):
		details := oci_devops.ComputeInstanceGroupByIdsSelector{}
		if computeInstanceIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_instance_ids")); ok {
			interfaces := computeInstanceIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "compute_instance_ids")) {
				details.ComputeInstanceIds = tmp
			}
		}
		baseObject = details
	case strings.ToLower("INSTANCE_QUERY"):
		details := oci_devops.ComputeInstanceGroupByQuerySelector{}
		if query, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query")); ok {
			tmp := query.(string)
			details.Query = &tmp
		}
		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown selector_type '%v' was specified", selectorType)
	}
	return baseObject, nil
}

func ComputeInstanceGroupSelectorToMap(obj oci_devops.ComputeInstanceGroupSelector) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_devops.ComputeInstanceGroupByIdsSelector:
		result["selector_type"] = "INSTANCE_IDS"

		result["compute_instance_ids"] = v.ComputeInstanceIds
	case oci_devops.ComputeInstanceGroupByQuerySelector:
		result["selector_type"] = "INSTANCE_QUERY"

		if v.Query != nil {
			result["query"] = string(*v.Query)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}
	default:
		log.Printf("[WARN] Received 'selector_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DevopsDeployEnvironmentResourceCrud) mapToComputeInstanceGroupSelectorCollection(fieldKeyFormat string) (oci_devops.ComputeInstanceGroupSelectorCollection, error) {
	result := oci_devops.ComputeInstanceGroupSelectorCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.ComputeInstanceGroupSelector, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToComputeInstanceGroupSelector(fieldKeyFormatNextLevel)
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

func ComputeInstanceGroupSelectorCollectionToMap(obj *oci_devops.ComputeInstanceGroupSelectorCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, ComputeInstanceGroupSelectorToMap(item))
	}
	result["items"] = items

	return result
}

func DeployEnvironmentSummaryToMap(obj oci_devops.DeployEnvironmentSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_devops.ComputeInstanceGroupDeployEnvironmentSummary:
		result["deploy_environment_type"] = "COMPUTE_INSTANCE_GROUP"

		if v.ComputeInstanceGroupSelectors != nil {
			result["compute_instance_group_selectors"] = []interface{}{ComputeInstanceGroupSelectorCollectionToMap(v.ComputeInstanceGroupSelectors)}
		}
	case oci_devops.FunctionDeployEnvironmentSummary:
		result["deploy_environment_type"] = "FUNCTION"

		if v.FunctionId != nil {
			result["function_id"] = string(*v.FunctionId)
		}
	case oci_devops.OkeClusterDeployEnvironmentSummary:
		result["deploy_environment_type"] = "OKE_CLUSTER"

		if v.ClusterId != nil {
			result["cluster_id"] = string(*v.ClusterId)
		}

		if v.NetworkChannel != nil {
			networkChannelArray := []interface{}{}
			if networkChannelMap := NetworkChannelToMap(&v.NetworkChannel); networkChannelMap != nil {
				networkChannelArray = append(networkChannelArray, networkChannelMap)
			}
			result["network_channel"] = networkChannelArray
		}
	default:
		log.Printf("[WARN] Received 'deploy_environment_type' of unknown type %v", obj)
		return nil
	}

	if obj.GetId() != nil {
		result["id"] = obj.GetId()
	}

	if obj.GetDefinedTags() != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
	}

	if obj.GetCompartmentId() != nil {
		result["compartment_id"] = obj.GetCompartmentId()
	}

	if obj.GetDescription() != nil {
		result["description"] = obj.GetDescription()
	}

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}

	if obj.GetTimeUpdated() != nil {
		result["time_updated"] = obj.GetTimeUpdated().String()
	}

	result["state"] = obj.GetLifecycleState()

	if obj.GetLifecycleDetails() != nil {
		result["lifecycle_details"] = obj.GetLifecycleDetails()
	}

	result["freeform_tags"] = obj.GetFreeformTags()

	if obj.GetProjectId() != nil {
		result["project_id"] = obj.GetProjectId()
	}

	if obj.GetSystemTags() != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.GetSystemTags())
	}

	return result
}

func (s *DevopsDeployEnvironmentResourceCrud) mapToNetworkChannel(fieldKeyFormat string) (oci_devops.NetworkChannel, error) {
	var baseObject oci_devops.NetworkChannel
	//discriminator
	networkChannelTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_channel_type"))
	var networkChannelType string
	if ok {
		networkChannelType = networkChannelTypeRaw.(string)
	} else {
		networkChannelType = "" // default value
	}
	switch strings.ToLower(networkChannelType) {
	case strings.ToLower("PRIVATE_ENDPOINT_CHANNEL"):
		details := oci_devops.PrivateEndpointChannel{}
		if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
			interfaces := nsgIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
				details.NsgIds = tmp
			}
		}
		if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		baseObject = details
	case strings.ToLower("SERVICE_VNIC_CHANNEL"):
		details := oci_devops.ServiceVnicChannel{}
		if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
				details.NsgIds = tmp
			}
		}
		if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown network_channel_type '%v' was specified", networkChannelType)
	}
	return baseObject, nil
}

func NetworkChannelToMap(obj *oci_devops.NetworkChannel) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.PrivateEndpointChannel:
		result["network_channel_type"] = "PRIVATE_ENDPOINT_CHANNEL"

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		result["nsg_ids"] = nsgIds

		if v.SubnetId != nil {
			result["subnet_id"] = string(*v.SubnetId)
		}
	case oci_devops.ServiceVnicChannel:
		result["network_channel_type"] = "SERVICE_VNIC_CHANNEL"

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		result["nsg_ids"] = nsgIds

		if v.SubnetId != nil {
			result["subnet_id"] = string(*v.SubnetId)
		}
	default:
		log.Printf("[WARN] Received 'network_channel_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DevopsDeployEnvironmentResourceCrud) populateTopLevelPolymorphicCreateDeployEnvironmentRequest(request *oci_devops.CreateDeployEnvironmentRequest) error {
	//discriminator
	deployEnvironmentTypeRaw, ok := s.D.GetOkExists("deploy_environment_type")
	var deployEnvironmentType string
	if ok {
		deployEnvironmentType = deployEnvironmentTypeRaw.(string)
	} else {
		deployEnvironmentType = "" // default value
	}
	switch strings.ToLower(deployEnvironmentType) {
	case strings.ToLower("COMPUTE_INSTANCE_GROUP"):
		details := oci_devops.CreateComputeInstanceGroupDeployEnvironmentDetails{}
		if computeInstanceGroupSelectors, ok := s.D.GetOkExists("compute_instance_group_selectors"); ok {
			if tmpList := computeInstanceGroupSelectors.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "compute_instance_group_selectors", 0)
				tmp, err := s.mapToComputeInstanceGroupSelectorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ComputeInstanceGroupSelectors = &tmp
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
		request.CreateDeployEnvironmentDetails = details
	case strings.ToLower("FUNCTION"):
		details := oci_devops.CreateFunctionDeployEnvironmentDetails{}
		if functionId, ok := s.D.GetOkExists("function_id"); ok {
			tmp := functionId.(string)
			details.FunctionId = &tmp
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
		request.CreateDeployEnvironmentDetails = details
	case strings.ToLower("OKE_CLUSTER"):
		details := oci_devops.CreateOkeClusterDeployEnvironmentDetails{}
		if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
			tmp := clusterId.(string)
			details.ClusterId = &tmp
		}
		if networkChannel, ok := s.D.GetOkExists("network_channel"); ok {
			if tmpList := networkChannel.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_channel", 0)
				tmp, err := s.mapToNetworkChannel(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.NetworkChannel = tmp
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
		request.CreateDeployEnvironmentDetails = details
	default:
		return fmt.Errorf("unknown deploy_environment_type '%v' was specified", deployEnvironmentType)
	}
	return nil
}

func (s *DevopsDeployEnvironmentResourceCrud) populateTopLevelPolymorphicUpdateDeployEnvironmentRequest(request *oci_devops.UpdateDeployEnvironmentRequest) error {
	//discriminator
	deployEnvironmentTypeRaw, ok := s.D.GetOkExists("deploy_environment_type")
	var deployEnvironmentType string
	if ok {
		deployEnvironmentType = deployEnvironmentTypeRaw.(string)
	} else {
		deployEnvironmentType = "" // default value
	}
	switch strings.ToLower(deployEnvironmentType) {
	case strings.ToLower("COMPUTE_INSTANCE_GROUP"):
		details := oci_devops.UpdateComputeInstanceGroupDeployEnvironmentDetails{}
		if computeInstanceGroupSelectors, ok := s.D.GetOkExists("compute_instance_group_selectors"); ok {
			if tmpList := computeInstanceGroupSelectors.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "compute_instance_group_selectors", 0)
				tmp, err := s.mapToComputeInstanceGroupSelectorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ComputeInstanceGroupSelectors = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployEnvironmentId = &tmp
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
		request.UpdateDeployEnvironmentDetails = details
	case strings.ToLower("FUNCTION"):
		details := oci_devops.UpdateFunctionDeployEnvironmentDetails{}
		if functionId, ok := s.D.GetOkExists("function_id"); ok {
			tmp := functionId.(string)
			details.FunctionId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployEnvironmentId = &tmp
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
		request.UpdateDeployEnvironmentDetails = details
	case strings.ToLower("OKE_CLUSTER"):
		details := oci_devops.UpdateOkeClusterDeployEnvironmentDetails{}
		if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
			tmp := clusterId.(string)
			details.ClusterId = &tmp
		}
		if networkChannel, ok := s.D.GetOkExists("network_channel"); ok {
			if tmpList := networkChannel.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_channel", 0)
				tmp, err := s.mapToNetworkChannel(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.NetworkChannel = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployEnvironmentId = &tmp
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
		request.UpdateDeployEnvironmentDetails = details
	default:
		return fmt.Errorf("unknown deploy_environment_type '%v' was specified", deployEnvironmentType)
	}
	return nil
}
