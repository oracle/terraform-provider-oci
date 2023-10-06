// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"bytes"
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
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func ContainerengineClusterNamespaceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createContainerengineClusterNamespace,
		Read:     readContainerengineClusterNamespace,
		Update:   updateContainerengineClusterNamespace,
		Delete:   deleteContainerengineClusterNamespace,
		Schema: map[string]*schema.Schema{
			// Required
			"cluster_namespace_profile_version_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"namespace_annotations": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      namespaceAnnotationsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"namespace_labels": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      namespaceLabelsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"cluster_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace": {
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

func createContainerengineClusterNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.CreateResource(d, sync)
}

func readContainerengineClusterNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

func updateContainerengineClusterNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteContainerengineClusterNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ContainerengineClusterNamespaceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_containerengine.ContainerEngineClient
	Res                    *oci_containerengine.ClusterNamespace
	DisableNotFoundRetries bool
}

func (s *ContainerengineClusterNamespaceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ContainerengineClusterNamespaceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_containerengine.ClusterNamespaceLifecycleStateCreating),
	}
}

func (s *ContainerengineClusterNamespaceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_containerengine.ClusterNamespaceLifecycleStateActive),
	}
}

func (s *ContainerengineClusterNamespaceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_containerengine.ClusterNamespaceLifecycleStateDeleting),
	}
}

func (s *ContainerengineClusterNamespaceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_containerengine.ClusterNamespaceLifecycleStateDeleted),
	}
}

func (s *ContainerengineClusterNamespaceResourceCrud) Create() error {
	request := oci_containerengine.CreateClusterNamespaceRequest{}

	if clusterNamespaceProfileVersionId, ok := s.D.GetOkExists("cluster_namespace_profile_version_id"); ok {
		tmp := clusterNamespaceProfileVersionId.(string)
		request.ClusterNamespaceProfileVersionId = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespaceAnnotations, ok := s.D.GetOkExists("namespace_annotations"); ok {
		set := namespaceAnnotations.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.NamespaceAnnotation, len(interfaces))
		for i := range interfaces {
			stateDataIndex := namespaceAnnotationsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "namespace_annotations", stateDataIndex)
			converted, err := s.mapToNamespaceAnnotation(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("namespace_annotations") {
			request.NamespaceAnnotations = tmp
		}
	}

	if namespaceLabels, ok := s.D.GetOkExists("namespace_labels"); ok {
		set := namespaceLabels.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.NamespaceLabel, len(interfaces))
		for i := range interfaces {
			stateDataIndex := namespaceLabelsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "namespace_labels", stateDataIndex)
			converted, err := s.mapToNamespaceLabel(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("namespace_labels") {
			request.NamespaceLabels = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.CreateClusterNamespace(context.Background(), request)
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "clusternamespace") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getClusterNamespaceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ContainerengineClusterNamespaceResourceCrud) getClusterNamespaceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_containerengine.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	clusterNamespaceId, err := clusterNamespaceWaitForWorkRequest(workId, "clusternamespace",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*clusterNamespaceId)

	return s.Get()
}

func clusterNamespaceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func clusterNamespaceWaitForWorkRequest(wId *string, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_containerengine.ContainerEngineClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "containerengine")
	retryPolicy.ShouldRetryOperation = clusterNamespaceWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromContainerengineClusterNamespaceWorkRequest(client, wId, response.CompartmentId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromContainerengineClusterNamespaceWorkRequest(client *oci_containerengine.ContainerEngineClient, workId *string, compartmentId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum) error {
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

func (s *ContainerengineClusterNamespaceResourceCrud) Get() error {
	request := oci_containerengine.GetClusterNamespaceRequest{}

	tmp := s.D.Id()
	request.ClusterNamespaceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.GetClusterNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ClusterNamespace
	return nil
}

func (s *ContainerengineClusterNamespaceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_containerengine.UpdateClusterNamespaceRequest{}

	tmp := s.D.Id()
	request.ClusterNamespaceId = &tmp

	if clusterNamespaceProfileVersionId, ok := s.D.GetOkExists("cluster_namespace_profile_version_id"); ok {
		tmp := clusterNamespaceProfileVersionId.(string)
		request.ClusterNamespaceProfileVersionId = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if namespaceAnnotations, ok := s.D.GetOkExists("namespace_annotations"); ok {
		set := namespaceAnnotations.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.NamespaceAnnotation, len(interfaces))
		for i := range interfaces {
			stateDataIndex := namespaceAnnotationsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "namespace_annotations", stateDataIndex)
			converted, err := s.mapToNamespaceAnnotation(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("namespace_annotations") {
			request.NamespaceAnnotations = tmp
		}
	}

	if namespaceLabels, ok := s.D.GetOkExists("namespace_labels"); ok {
		set := namespaceLabels.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.NamespaceLabel, len(interfaces))
		for i := range interfaces {
			stateDataIndex := namespaceLabelsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "namespace_labels", stateDataIndex)
			converted, err := s.mapToNamespaceLabel(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("namespace_labels") {
			request.NamespaceLabels = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.UpdateClusterNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getClusterNamespaceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ContainerengineClusterNamespaceResourceCrud) Delete() error {
	request := oci_containerengine.DeleteClusterNamespaceRequest{}

	tmp := s.D.Id()
	request.ClusterNamespaceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.DeleteClusterNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := clusterNamespaceWaitForWorkRequest(workId, "clusternamespace",
		oci_containerengine.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ContainerengineClusterNamespaceResourceCrud) SetData() error {
	s.D.Set("cluster_ids", s.Res.ClusterIds)
	s.D.Set("cluster_ids", s.Res.ClusterIds)

	if s.Res.ClusterNamespaceProfileVersionId != nil {
		s.D.Set("cluster_namespace_profile_version_id", *s.Res.ClusterNamespaceProfileVersionId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.NamespaceName != nil {
		s.D.Set("namespace", *s.Res.NamespaceName)
	}

	namespaceAnnotations := []interface{}{}
	for _, item := range s.Res.NamespaceAnnotations {
		namespaceAnnotations = append(namespaceAnnotations, NamespaceAnnotationToMap(item))
	}
	s.D.Set("namespace_annotations", schema.NewSet(namespaceAnnotationsHashCodeForSets, namespaceAnnotations))

	namespaceLabels := []interface{}{}
	for _, item := range s.Res.NamespaceLabels {
		namespaceLabels = append(namespaceLabels, NamespaceLabelToMap(item))
	}
	s.D.Set("namespace_labels", schema.NewSet(namespaceLabelsHashCodeForSets, namespaceLabels))

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

func ClusterNamespaceSummaryToMap(obj oci_containerengine.ClusterNamespaceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClusterNamespaceProfileVersionId != nil {
		result["cluster_namespace_profile_version_id"] = string(*obj.ClusterNamespaceProfileVersionId)
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

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
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

func (s *ContainerengineClusterNamespaceResourceCrud) mapToNamespaceAnnotation(fieldKeyFormat string) (oci_containerengine.NamespaceAnnotation, error) {
	result := oci_containerengine.NamespaceAnnotation{}

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

func (s *ContainerengineClusterNamespaceResourceCrud) mapToNamespaceLabel(fieldKeyFormat string) (oci_containerengine.NamespaceLabel, error) {
	result := oci_containerengine.NamespaceLabel{}

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

func namespaceAnnotationsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if key, ok := m["key"]; ok && key != "" {
		buf.WriteString(fmt.Sprintf("%v-", key))
	}
	if value, ok := m["value"]; ok && value != "" {
		buf.WriteString(fmt.Sprintf("%v-", value))
	}
	return utils.GetStringHashcode(buf.String())
}

func namespaceLabelsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if key, ok := m["key"]; ok && key != "" {
		buf.WriteString(fmt.Sprintf("%v-", key))
	}
	if value, ok := m["value"]; ok && value != "" {
		buf.WriteString(fmt.Sprintf("%v-", value))
	}
	return utils.GetStringHashcode(buf.String())
}

func (s *ContainerengineClusterNamespaceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_containerengine.ChangeClusterNamespaceCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ClusterNamespaceId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.ChangeClusterNamespaceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getClusterNamespaceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
