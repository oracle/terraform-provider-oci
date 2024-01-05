// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineVirtualNodePoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createContainerengineVirtualNodePool,
		Read:     readContainerengineVirtualNodePool,
		Update:   updateContainerengineVirtualNodePool,
		Delete:   deleteContainerengineVirtualNodePool,
		Schema: map[string]*schema.Schema{
			// Required
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"placement_configurations": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"availability_domain": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},
						"fault_domain": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"initial_virtual_node_labels": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			// Required
			"pod_configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			// Optional
			"taints": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"effect": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"virtual_node_tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
					},
				},
			},

			// Computed
			"kubernetes_version": {
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

func createContainerengineVirtualNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineVirtualNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.CreateResource(d, sync)
}

func readContainerengineVirtualNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineVirtualNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

func updateContainerengineVirtualNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineVirtualNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteContainerengineVirtualNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineVirtualNodePoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ContainerengineVirtualNodePoolResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_containerengine.ContainerEngineClient
	Res                    *oci_containerengine.VirtualNodePool
	DisableNotFoundRetries bool
}

func (s *ContainerengineVirtualNodePoolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ContainerengineVirtualNodePoolResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_containerengine.VirtualNodePoolLifecycleStateCreating),
	}
}

func (s *ContainerengineVirtualNodePoolResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_containerengine.VirtualNodePoolLifecycleStateActive),
		string(oci_containerengine.VirtualNodePoolLifecycleStateNeedsAttention),
	}
}

func (s *ContainerengineVirtualNodePoolResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_containerengine.VirtualNodePoolLifecycleStateDeleting),
	}
}

func (s *ContainerengineVirtualNodePoolResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_containerengine.VirtualNodePoolLifecycleStateDeleted),
	}
}

func (s *ContainerengineVirtualNodePoolResourceCrud) Create() error {
	request := oci_containerengine.CreateVirtualNodePoolRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
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

	if initialVirtualNodeLabels, ok := s.D.GetOkExists("initial_virtual_node_labels"); ok {
		interfaces := initialVirtualNodeLabels.([]interface{})
		tmp := make([]oci_containerengine.InitialVirtualNodeLabel, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "initial_virtual_node_labels", stateDataIndex)
			converted, err := s.mapToInitialVirtualNodeLabel(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("initial_virtual_node_labels") {
			request.InitialVirtualNodeLabels = tmp
		}
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if placementConfigurations, ok := s.D.GetOkExists("placement_configurations"); ok {
		interfaces := placementConfigurations.([]interface{})
		tmp := make([]oci_containerengine.PlacementConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "placement_configurations", stateDataIndex)
			converted, err := s.mapToPlacementConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("placement_configurations") {
			request.PlacementConfigurations = tmp
		}
	}

	if podConfiguration, ok := s.D.GetOkExists("pod_configuration"); ok {
		if tmpList := podConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "pod_configuration", 0)
			tmp, err := s.mapToPodConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PodConfiguration = &tmp
		}
	}

	if size, ok := s.D.GetOkExists("size"); ok {
		tmp := size.(int)
		request.Size = &tmp
	}

	if taints, ok := s.D.GetOkExists("taints"); ok {
		interfaces := taints.([]interface{})
		tmp := make([]oci_containerengine.Taint, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "taints", stateDataIndex)
			converted, err := s.mapToTaint(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("taints") {
			request.Taints = tmp
		}
	}

	if virtualNodeTags, ok := s.D.GetOkExists("virtual_node_tags"); ok {
		if tmpList := virtualNodeTags.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "virtual_node_tags", 0)
			tmp, err := s.mapToVirtualNodeTags(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.VirtualNodeTags = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.CreateVirtualNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_containerengine.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_containerengine.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "virtualnodepool") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getVirtualNodePoolFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ContainerengineVirtualNodePoolResourceCrud) getVirtualNodePoolFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_containerengine.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	virtualNodePoolId, err := virtualNodePoolWaitForWorkRequest(workId, "virtualnodepool",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*virtualNodePoolId)

	return s.Get()
}

func virtualNodePoolWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "containerengine", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_containerengine.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func virtualNodePoolWaitForWorkRequest(wId *string, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_containerengine.ContainerEngineClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "containerengine")
	retryPolicy.ShouldRetryOperation = virtualNodePoolWorkRequestShouldRetryFunc(timeout)

	response := oci_containerengine.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_containerengine.WorkRequestStatusInProgress),
			string(oci_containerengine.WorkRequestStatusAccepted),
			string(oci_containerengine.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_containerengine.WorkRequestStatusSucceeded),
			string(oci_containerengine.WorkRequestStatusFailed),
			string(oci_containerengine.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_containerengine.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_containerengine.WorkRequestStatusFailed || response.Status == oci_containerengine.WorkRequestStatusCanceled {
		return nil, getErrorFromContainerengineVirtualNodePoolWorkRequest(client, wId, response.CompartmentId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromContainerengineVirtualNodePoolWorkRequest(client *oci_containerengine.ContainerEngineClient, workId *string, compartmentId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_containerengine.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			CompartmentId: compartmentId,
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

func (s *ContainerengineVirtualNodePoolResourceCrud) Get() error {
	request := oci_containerengine.GetVirtualNodePoolRequest{}

	tmp := s.D.Id()
	request.VirtualNodePoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.GetVirtualNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VirtualNodePool
	return nil
}

func (s *ContainerengineVirtualNodePoolResourceCrud) Update() error {
	request := oci_containerengine.UpdateVirtualNodePoolRequest{}

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

	if initialVirtualNodeLabels, ok := s.D.GetOkExists("initial_virtual_node_labels"); ok {
		interfaces := initialVirtualNodeLabels.([]interface{})
		tmp := make([]oci_containerengine.InitialVirtualNodeLabel, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "initial_virtual_node_labels", stateDataIndex)
			converted, err := s.mapToInitialVirtualNodeLabel(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("initial_virtual_node_labels") {
			request.InitialVirtualNodeLabels = tmp
		}
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if placementConfigurations, ok := s.D.GetOkExists("placement_configurations"); ok {
		interfaces := placementConfigurations.([]interface{})
		tmp := make([]oci_containerengine.PlacementConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "placement_configurations", stateDataIndex)
			converted, err := s.mapToPlacementConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("placement_configurations") {
			request.PlacementConfigurations = tmp
		}
	}

	if podConfiguration, ok := s.D.GetOkExists("pod_configuration"); ok {
		if tmpList := podConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "pod_configuration", 0)
			tmp, err := s.mapToPodConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PodConfiguration = &tmp
		}
	}

	if size, ok := s.D.GetOkExists("size"); ok {
		tmp := size.(int)
		request.Size = &tmp
	}

	if taints, ok := s.D.GetOkExists("taints"); ok {
		interfaces := taints.([]interface{})
		tmp := make([]oci_containerengine.Taint, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "taints", stateDataIndex)
			converted, err := s.mapToTaint(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("taints") {
			request.Taints = tmp
		}
	}

	tmp := s.D.Id()
	request.VirtualNodePoolId = &tmp

	if virtualNodeTags, ok := s.D.GetOkExists("virtual_node_tags"); ok {
		if tmpList := virtualNodeTags.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "virtual_node_tags", 0)
			tmp, err := s.mapToVirtualNodeTags(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.VirtualNodeTags = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.UpdateVirtualNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVirtualNodePoolFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ContainerengineVirtualNodePoolResourceCrud) Delete() error {
	request := oci_containerengine.DeleteVirtualNodePoolRequest{}

	if isForceDeletionAfterOverrideGraceDurationVnp, ok := s.D.GetOkExists("is_force_deletion_after_override_grace_duration_vnp"); ok {
		tmp := isForceDeletionAfterOverrideGraceDurationVnp.(bool)
		request.IsForceDeletionAfterOverrideGraceDurationVnp = &tmp
	}

	if overrideEvictionGraceDurationVnp, ok := s.D.GetOkExists("override_eviction_grace_duration_vnp"); ok {
		tmp := overrideEvictionGraceDurationVnp.(string)
		request.OverrideEvictionGraceDurationVnp = &tmp
	}

	tmp := s.D.Id()
	request.VirtualNodePoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.DeleteVirtualNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := virtualNodePoolWaitForWorkRequest(workId, "virtualnodepool",
		oci_containerengine.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ContainerengineVirtualNodePoolResourceCrud) SetData() error {
	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
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

	initialVirtualNodeLabels := []interface{}{}
	for _, item := range s.Res.InitialVirtualNodeLabels {
		initialVirtualNodeLabels = append(initialVirtualNodeLabels, InitialVirtualNodeLabelToMap(item))
	}
	s.D.Set("initial_virtual_node_labels", initialVirtualNodeLabels)

	if s.Res.KubernetesVersion != nil {
		s.D.Set("kubernetes_version", *s.Res.KubernetesVersion)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	placementConfigurations := []interface{}{}
	for _, item := range s.Res.PlacementConfigurations {
		placementConfigurations = append(placementConfigurations, PlacementConfigurationToMap(item))
	}
	s.D.Set("placement_configurations", placementConfigurations)

	if s.Res.PodConfiguration != nil {
		s.D.Set("pod_configuration", []interface{}{PodConfigurationToMap(s.Res.PodConfiguration, false)})
	} else {
		s.D.Set("pod_configuration", nil)
	}

	if s.Res.Size != nil {
		s.D.Set("size", *s.Res.Size)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	taints := []interface{}{}
	for _, item := range s.Res.Taints {
		taints = append(taints, TaintToMap(item))
	}
	s.D.Set("taints", taints)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VirtualNodeTags != nil {
		s.D.Set("virtual_node_tags", []interface{}{VirtualNodeTagsToMap(s.Res.VirtualNodeTags)})
	} else {
		s.D.Set("virtual_node_tags", nil)
	}

	return nil
}

func (s *ContainerengineVirtualNodePoolResourceCrud) mapToInitialVirtualNodeLabel(fieldKeyFormat string) (oci_containerengine.InitialVirtualNodeLabel, error) {
	result := oci_containerengine.InitialVirtualNodeLabel{}

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

func InitialVirtualNodeLabelToMap(obj oci_containerengine.InitialVirtualNodeLabel) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *ContainerengineVirtualNodePoolResourceCrud) mapToPlacementConfiguration(fieldKeyFormat string) (oci_containerengine.PlacementConfiguration, error) {
	result := oci_containerengine.PlacementConfiguration{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if faultDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_domain")); ok {
		interfaces := faultDomain.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "fault_domain")) {
			result.FaultDomain = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func PlacementConfigurationToMap(obj oci_containerengine.PlacementConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	result["fault_domain"] = obj.FaultDomain

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *ContainerengineVirtualNodePoolResourceCrud) mapToPodConfiguration(fieldKeyFormat string) (oci_containerengine.PodConfiguration, error) {
	result := oci_containerengine.PodConfiguration{}

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
			result.NsgIds = tmp
		}
	}

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func PodConfigurationToMap(obj *oci_containerengine.PodConfiguration, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *ContainerengineVirtualNodePoolResourceCrud) mapToTaint(fieldKeyFormat string) (oci_containerengine.Taint, error) {
	result := oci_containerengine.Taint{}

	if effect, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "effect")); ok {
		tmp := effect.(string)
		result.Effect = &tmp
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

func TaintToMap(obj oci_containerengine.Taint) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Effect != nil {
		result["effect"] = string(*obj.Effect)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *ContainerengineVirtualNodePoolResourceCrud) mapToVirtualNodeTags(fieldKeyFormat string) (oci_containerengine.VirtualNodeTags, error) {
	result := oci_containerengine.VirtualNodeTags{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	return result, nil
}

func VirtualNodeTagsToMap(obj *oci_containerengine.VirtualNodeTags) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	return result
}
