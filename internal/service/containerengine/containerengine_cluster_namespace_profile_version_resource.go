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

func ContainerengineClusterNamespaceProfileVersionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createContainerengineClusterNamespaceProfileVersion,
		Read:     readContainerengineClusterNamespaceProfileVersion,
		Update:   updateContainerengineClusterNamespaceProfileVersion,
		Delete:   deleteContainerengineClusterNamespaceProfileVersion,
		Schema: map[string]*schema.Schema{
			// Required
			"admin_cluster_role_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_namespace_profile_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"allowed_namespace_annotations": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Set:      allowedNamespaceAnnotationsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"value": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"allowed_namespace_labels": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Set:      allowedNamespaceLabelsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"value": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
			"fixed_namespace_annotations": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Set:      fixedNamespaceAnnotationsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"fixed_namespace_labels": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Set:      fixedNamespaceLabelsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_deprecated": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"required_namespace_annotations": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Set:      requiredNamespaceAnnotationsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"value": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"required_namespace_labels": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Set:      requiredNamespaceLabelsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"value": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},

			// Computed
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

func createContainerengineClusterNamespaceProfileVersion(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterNamespaceProfileVersionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.CreateResource(d, sync)
}

func readContainerengineClusterNamespaceProfileVersion(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterNamespaceProfileVersionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

func updateContainerengineClusterNamespaceProfileVersion(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterNamespaceProfileVersionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteContainerengineClusterNamespaceProfileVersion(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterNamespaceProfileVersionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ContainerengineClusterNamespaceProfileVersionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_containerengine.ContainerEngineClient
	Res                    *oci_containerengine.ClusterNamespaceProfileVersion
	DisableNotFoundRetries bool
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_containerengine.ClusterNamespaceProfileVersionLifecycleStateCreating),
	}
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_containerengine.ClusterNamespaceProfileVersionLifecycleStateActive),
	}
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_containerengine.ClusterNamespaceProfileVersionLifecycleStateDeleting),
	}
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_containerengine.ClusterNamespaceProfileVersionLifecycleStateDeleted),
	}
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) Create() error {
	request := oci_containerengine.CreateClusterNamespaceProfileVersionRequest{}

	if adminClusterRoleName, ok := s.D.GetOkExists("admin_cluster_role_name"); ok {
		tmp := adminClusterRoleName.(string)
		request.AdminClusterRoleName = &tmp
	}

	if allowedNamespaceAnnotations, ok := s.D.GetOkExists("allowed_namespace_annotations"); ok {
		set := allowedNamespaceAnnotations.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.AllowedNamespaceAnnotation, len(interfaces))
		for i := range interfaces {
			stateDataIndex := allowedNamespaceAnnotationsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "allowed_namespace_annotations", stateDataIndex)
			converted, err := s.mapToAllowedNamespaceAnnotation(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_namespace_annotations") {
			request.AllowedNamespaceAnnotations = tmp
		}
	}

	if allowedNamespaceLabels, ok := s.D.GetOkExists("allowed_namespace_labels"); ok {
		set := allowedNamespaceLabels.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.AllowedNamespaceLabel, len(interfaces))
		for i := range interfaces {
			stateDataIndex := allowedNamespaceLabelsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "allowed_namespace_labels", stateDataIndex)
			converted, err := s.mapToAllowedNamespaceLabel(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_namespace_labels") {
			request.AllowedNamespaceLabels = tmp
		}
	}

	if clusterNamespaceProfileId, ok := s.D.GetOkExists("cluster_namespace_profile_id"); ok {
		tmp := clusterNamespaceProfileId.(string)
		request.ClusterNamespaceProfileId = &tmp
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

	if fixedNamespaceAnnotations, ok := s.D.GetOkExists("fixed_namespace_annotations"); ok {
		set := fixedNamespaceAnnotations.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.NamespaceAnnotation, len(interfaces))
		for i := range interfaces {
			stateDataIndex := fixedNamespaceAnnotationsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "fixed_namespace_annotations", stateDataIndex)
			converted, err := s.mapToNamespaceAnnotation(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("fixed_namespace_annotations") {
			request.FixedNamespaceAnnotations = tmp
		}
	}

	if fixedNamespaceLabels, ok := s.D.GetOkExists("fixed_namespace_labels"); ok {
		set := fixedNamespaceLabels.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.NamespaceLabel, len(interfaces))
		for i := range interfaces {
			stateDataIndex := fixedNamespaceLabelsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "fixed_namespace_labels", stateDataIndex)
			converted, err := s.mapToNamespaceLabel(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("fixed_namespace_labels") {
			request.FixedNamespaceLabels = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isDeprecated, ok := s.D.GetOkExists("is_deprecated"); ok {
		tmp := isDeprecated.(bool)
		request.IsDeprecated = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if requiredNamespaceAnnotations, ok := s.D.GetOkExists("required_namespace_annotations"); ok {
		set := requiredNamespaceAnnotations.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.RequiredNamespaceAnnotation, len(interfaces))
		for i := range interfaces {
			stateDataIndex := requiredNamespaceAnnotationsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "required_namespace_annotations", stateDataIndex)
			converted, err := s.mapToRequiredNamespaceAnnotation(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("required_namespace_annotations") {
			request.RequiredNamespaceAnnotations = tmp
		}
	}

	if requiredNamespaceLabels, ok := s.D.GetOkExists("required_namespace_labels"); ok {
		set := requiredNamespaceLabels.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_containerengine.RequiredNamespaceLabel, len(interfaces))
		for i := range interfaces {
			stateDataIndex := requiredNamespaceLabelsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "required_namespace_labels", stateDataIndex)
			converted, err := s.mapToRequiredNamespaceLabel(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("required_namespace_labels") {
			request.RequiredNamespaceLabels = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.CreateClusterNamespaceProfileVersion(context.Background(), request)
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "clusternamespaceprofileversion") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getClusterNamespaceProfileVersionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) getClusterNamespaceProfileVersionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_containerengine.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	clusterNamespaceProfileVersionId, err := clusterNamespaceProfileVersionWaitForWorkRequest(workId, "clusternamespaceprofileversion",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*clusterNamespaceProfileVersionId)

	return s.Get()
}

func clusterNamespaceProfileVersionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func clusterNamespaceProfileVersionWaitForWorkRequest(wId *string, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_containerengine.ContainerEngineClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "containerengine")
	retryPolicy.ShouldRetryOperation = clusterNamespaceProfileVersionWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromContainerengineClusterNamespaceProfileVersionWorkRequest(client, wId, response.CompartmentId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromContainerengineClusterNamespaceProfileVersionWorkRequest(client *oci_containerengine.ContainerEngineClient, workId *string, compartmentId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_containerengine.WorkRequestResourceActionTypeEnum) error {
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

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) Get() error {
	request := oci_containerengine.GetClusterNamespaceProfileVersionRequest{}

	tmp := s.D.Id()
	request.ClusterNamespaceProfileVersionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.GetClusterNamespaceProfileVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ClusterNamespaceProfileVersion
	return nil
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_containerengine.UpdateClusterNamespaceProfileVersionRequest{}

	tmp := s.D.Id()
	request.ClusterNamespaceProfileVersionId = &tmp

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

	if isDeprecated, ok := s.D.GetOkExists("is_deprecated"); ok {
		tmp := isDeprecated.(bool)
		request.IsDeprecated = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.UpdateClusterNamespaceProfileVersion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getClusterNamespaceProfileVersionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) Delete() error {
	request := oci_containerengine.DeleteClusterNamespaceProfileVersionRequest{}

	tmp := s.D.Id()
	request.ClusterNamespaceProfileVersionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.DeleteClusterNamespaceProfileVersion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := clusterNamespaceProfileVersionWaitForWorkRequest(workId, "clusternamespaceprofileversion",
		oci_containerengine.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) SetData() error {
	if s.Res.AdminClusterRoleName != nil {
		s.D.Set("admin_cluster_role_name", *s.Res.AdminClusterRoleName)
	}

	allowedNamespaceAnnotations := []interface{}{}
	for _, item := range s.Res.AllowedNamespaceAnnotations {
		allowedNamespaceAnnotations = append(allowedNamespaceAnnotations, AllowedNamespaceAnnotationToMap(item, false))
	}
	s.D.Set("allowed_namespace_annotations", schema.NewSet(allowedNamespaceAnnotationsHashCodeForSets, allowedNamespaceAnnotations))

	allowedNamespaceLabels := []interface{}{}
	for _, item := range s.Res.AllowedNamespaceLabels {
		allowedNamespaceLabels = append(allowedNamespaceLabels, AllowedNamespaceLabelToMap(item, false))
	}
	s.D.Set("allowed_namespace_labels", schema.NewSet(allowedNamespaceLabelsHashCodeForSets, allowedNamespaceLabels))

	if s.Res.ClusterNamespaceProfileId != nil {
		s.D.Set("cluster_namespace_profile_id", *s.Res.ClusterNamespaceProfileId)
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

	fixedNamespaceAnnotations := []interface{}{}
	for _, item := range s.Res.FixedNamespaceAnnotations {
		fixedNamespaceAnnotations = append(fixedNamespaceAnnotations, NamespaceAnnotationToMap(item))
	}
	s.D.Set("fixed_namespace_annotations", schema.NewSet(fixedNamespaceAnnotationsHashCodeForSets, fixedNamespaceAnnotations))

	fixedNamespaceLabels := []interface{}{}
	for _, item := range s.Res.FixedNamespaceLabels {
		fixedNamespaceLabels = append(fixedNamespaceLabels, NamespaceLabelToMap(item))
	}
	s.D.Set("fixed_namespace_labels", schema.NewSet(fixedNamespaceLabelsHashCodeForSets, fixedNamespaceLabels))

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsDeprecated != nil {
		s.D.Set("is_deprecated", *s.Res.IsDeprecated)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	requiredNamespaceAnnotations := []interface{}{}
	for _, item := range s.Res.RequiredNamespaceAnnotations {
		requiredNamespaceAnnotations = append(requiredNamespaceAnnotations, RequiredNamespaceAnnotationToMap(item, false))
	}
	s.D.Set("required_namespace_annotations", schema.NewSet(requiredNamespaceAnnotationsHashCodeForSets, requiredNamespaceAnnotations))

	requiredNamespaceLabels := []interface{}{}
	for _, item := range s.Res.RequiredNamespaceLabels {
		requiredNamespaceLabels = append(requiredNamespaceLabels, RequiredNamespaceLabelToMap(item, false))
	}
	s.D.Set("required_namespace_labels", schema.NewSet(requiredNamespaceLabelsHashCodeForSets, requiredNamespaceLabels))

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

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) mapToAllowedNamespaceAnnotation(fieldKeyFormat string) (oci_containerengine.AllowedNamespaceAnnotation, error) {
	result := oci_containerengine.AllowedNamespaceAnnotation{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		set := value.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "value")) {
			result.Value = tmp
		}
	}

	return result, nil
}

func AllowedNamespaceAnnotationToMap(obj oci_containerengine.AllowedNamespaceAnnotation, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	value := []interface{}{}
	for _, item := range obj.Value {
		value = append(value, item)
	}
	if datasource {
		result["value"] = value
	} else {
		result["value"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, value)
	}

	return result
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) mapToAllowedNamespaceLabel(fieldKeyFormat string) (oci_containerengine.AllowedNamespaceLabel, error) {
	result := oci_containerengine.AllowedNamespaceLabel{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		set := value.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "value")) {
			result.Value = tmp
		}
	}

	return result, nil
}

func AllowedNamespaceLabelToMap(obj oci_containerengine.AllowedNamespaceLabel, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	value := []interface{}{}
	for _, item := range obj.Value {
		value = append(value, item)
	}
	if datasource {
		result["value"] = value
	} else {
		result["value"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, value)
	}

	return result
}

func ClusterNamespaceProfileVersionSummaryToMap(obj oci_containerengine.ClusterNamespaceProfileVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClusterNamespaceProfileId != nil {
		result["cluster_namespace_profile_id"] = string(*obj.ClusterNamespaceProfileId)
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

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) mapToNamespaceAnnotation(fieldKeyFormat string) (oci_containerengine.NamespaceAnnotation, error) {
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

func NamespaceAnnotationToMap(obj oci_containerengine.NamespaceAnnotation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) mapToNamespaceLabel(fieldKeyFormat string) (oci_containerengine.NamespaceLabel, error) {
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

func NamespaceLabelToMap(obj oci_containerengine.NamespaceLabel) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) mapToRequiredNamespaceAnnotation(fieldKeyFormat string) (oci_containerengine.RequiredNamespaceAnnotation, error) {
	result := oci_containerengine.RequiredNamespaceAnnotation{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		set := value.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "value")) {
			result.Value = tmp
		}
	}

	return result, nil
}

func RequiredNamespaceAnnotationToMap(obj oci_containerengine.RequiredNamespaceAnnotation, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	value := []interface{}{}
	for _, item := range obj.Value {
		value = append(value, item)
	}
	if datasource {
		result["value"] = value
	} else {
		result["value"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, value)
	}

	return result
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) mapToRequiredNamespaceLabel(fieldKeyFormat string) (oci_containerengine.RequiredNamespaceLabel, error) {
	result := oci_containerengine.RequiredNamespaceLabel{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		set := value.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "value")) {
			result.Value = tmp
		}
	}

	return result, nil
}

func RequiredNamespaceLabelToMap(obj oci_containerengine.RequiredNamespaceLabel, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	value := []interface{}{}
	for _, item := range obj.Value {
		value = append(value, item)
	}
	if datasource {
		result["value"] = value
	} else {
		result["value"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, value)
	}

	return result
}

func allowedNamespaceAnnotationsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if key, ok := m["key"]; ok && key != "" {
		buf.WriteString(fmt.Sprintf("%v-", key))
	}
	if value, ok := m["value"]; ok && value != "" {
	}
	return utils.GetStringHashcode(buf.String())
}

func allowedNamespaceLabelsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if key, ok := m["key"]; ok && key != "" {
		buf.WriteString(fmt.Sprintf("%v-", key))
	}
	if value, ok := m["value"]; ok && value != "" {
	}
	return utils.GetStringHashcode(buf.String())
}

func fixedNamespaceAnnotationsHashCodeForSets(v interface{}) int {
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

func fixedNamespaceLabelsHashCodeForSets(v interface{}) int {
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

func requiredNamespaceAnnotationsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if key, ok := m["key"]; ok && key != "" {
		buf.WriteString(fmt.Sprintf("%v-", key))
	}
	if value, ok := m["value"]; ok && value != "" {
	}
	return utils.GetStringHashcode(buf.String())
}

func requiredNamespaceLabelsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if key, ok := m["key"]; ok && key != "" {
		buf.WriteString(fmt.Sprintf("%v-", key))
	}
	if value, ok := m["value"]; ok && value != "" {
	}
	return utils.GetStringHashcode(buf.String())
}

func (s *ContainerengineClusterNamespaceProfileVersionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_containerengine.ChangeClusterNamespaceProfileVersionCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ClusterNamespaceProfileVersionId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.ChangeClusterNamespaceProfileVersionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getClusterNamespaceProfileVersionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"), oci_containerengine.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
