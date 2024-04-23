// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cluster_placement_groups

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_cluster_placement_groups "github.com/oracle/oci-go-sdk/v65/clusterplacementgroups"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ClusterPlacementGroupsClusterPlacementGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createClusterPlacementGroupsClusterPlacementGroup,
		Read:     readClusterPlacementGroupsClusterPlacementGroup,
		Update:   updateClusterPlacementGroupsClusterPlacementGroup,
		Delete:   deleteClusterPlacementGroupsClusterPlacementGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"cluster_placement_group_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"capabilities": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"items": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"service": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional

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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"opc_dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"placement_instruction": {
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
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateInactive),
					string(oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateActive),
				}, true),
			},

			// Computed
			"lifecycle_details": {
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

func createClusterPlacementGroupsClusterPlacementGroup(d *schema.ResourceData, m interface{}) error {
	sync := &ClusterPlacementGroupsClusterPlacementGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ClusterPlacementGroupsCPClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopClusterPlacementGroup(); err != nil {
			return err
		}
		sync.D.Set("state", oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateInactive)
	}
	return nil

}

func readClusterPlacementGroupsClusterPlacementGroup(d *schema.ResourceData, m interface{}) error {
	sync := &ClusterPlacementGroupsClusterPlacementGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ClusterPlacementGroupsCPClient()

	return tfresource.ReadResource(sync)
}

func updateClusterPlacementGroupsClusterPlacementGroup(d *schema.ResourceData, m interface{}) error {
	sync := &ClusterPlacementGroupsClusterPlacementGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ClusterPlacementGroupsCPClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateActive == oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateInactive == oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartClusterPlacementGroup(); err != nil {
			return err
		}
		sync.D.Set("state", oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopClusterPlacementGroup(); err != nil {
			return err
		}
		sync.D.Set("state", oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateInactive)
	}

	return nil
}

func deleteClusterPlacementGroupsClusterPlacementGroup(d *schema.ResourceData, m interface{}) error {
	sync := &ClusterPlacementGroupsClusterPlacementGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ClusterPlacementGroupsCPClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ClusterPlacementGroupsClusterPlacementGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cluster_placement_groups.ClusterPlacementGroupsCPClient
	Res                    *oci_cluster_placement_groups.ClusterPlacementGroup
	DisableNotFoundRetries bool
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateCreating),
	}
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateActive),
	}
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateDeleting),
	}
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateDeleted),
	}
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) Create() error {
	request := oci_cluster_placement_groups.CreateClusterPlacementGroupRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if capabilities, ok := s.D.GetOkExists("capabilities"); ok {
		if tmpList := capabilities.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "capabilities", 0)
			tmp, err := s.mapToCapabilitiesCollection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Capabilities = &tmp
		}
	}

	if clusterPlacementGroupType, ok := s.D.GetOkExists("cluster_placement_group_type"); ok {
		request.ClusterPlacementGroupType = oci_cluster_placement_groups.ClusterPlacementGroupTypeEnum(clusterPlacementGroupType.(string))
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

	if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
		tmp := opcDryRun.(bool)
		request.OpcDryRun = &tmp
	}

	if placementInstruction, ok := s.D.GetOkExists("placement_instruction"); ok {
		if tmpList := placementInstruction.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "placement_instruction", 0)
			tmp, err := s.mapToPlacementInstructionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PlacementInstruction = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_placement_groups")

	response, err := s.Client.CreateClusterPlacementGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getClusterPlacementGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_placement_groups"), oci_cluster_placement_groups.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) getClusterPlacementGroupFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_cluster_placement_groups.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	clusterPlacementGroupId, err := clusterPlacementGroupWaitForWorkRequest(workId, "clusterplacementgroup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, clusterPlacementGroupId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_cluster_placement_groups.CancelWorkRequestRequest{
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
	s.D.SetId(*clusterPlacementGroupId)

	return s.Get()
}

func clusterPlacementGroupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "cluster_placement_groups", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_cluster_placement_groups.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func clusterPlacementGroupWaitForWorkRequest(wId *string, entityType string, action oci_cluster_placement_groups.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_cluster_placement_groups.ClusterPlacementGroupsCPClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "cluster_placement_groups")
	retryPolicy.ShouldRetryOperation = clusterPlacementGroupWorkRequestShouldRetryFunc(timeout)

	response := oci_cluster_placement_groups.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_cluster_placement_groups.OperationStatusInProgress),
			string(oci_cluster_placement_groups.OperationStatusAccepted),
			string(oci_cluster_placement_groups.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_cluster_placement_groups.OperationStatusSucceeded),
			string(oci_cluster_placement_groups.OperationStatusFailed),
			string(oci_cluster_placement_groups.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_cluster_placement_groups.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_cluster_placement_groups.OperationStatusFailed || response.Status == oci_cluster_placement_groups.OperationStatusCanceled {
		return nil, getErrorFromClusterPlacementGroupsClusterPlacementGroupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromClusterPlacementGroupsClusterPlacementGroupWorkRequest(client *oci_cluster_placement_groups.ClusterPlacementGroupsCPClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_cluster_placement_groups.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_cluster_placement_groups.ListWorkRequestErrorsRequest{
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

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) Get() error {
	request := oci_cluster_placement_groups.GetClusterPlacementGroupRequest{}

	tmp := s.D.Id()
	request.ClusterPlacementGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_placement_groups")

	response, err := s.Client.GetClusterPlacementGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ClusterPlacementGroup
	return nil
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cluster_placement_groups.UpdateClusterPlacementGroupRequest{}

	tmp := s.D.Id()
	request.ClusterPlacementGroupId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_placement_groups")

	response, err := s.Client.UpdateClusterPlacementGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getClusterPlacementGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_placement_groups"), oci_cluster_placement_groups.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) Delete() error {
	request := oci_cluster_placement_groups.DeleteClusterPlacementGroupRequest{}

	tmp := s.D.Id()
	request.ClusterPlacementGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_placement_groups")

	response, err := s.Client.DeleteClusterPlacementGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := clusterPlacementGroupWaitForWorkRequest(workId, "clusterplacementgroup",
		oci_cluster_placement_groups.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.Capabilities != nil {
		s.D.Set("capabilities", []interface{}{CapabilitiesCollectionToMap(s.Res.Capabilities)})
	} else {
		s.D.Set("capabilities", nil)
	}

	s.D.Set("cluster_placement_group_type", s.Res.ClusterPlacementGroupType)

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PlacementInstruction != nil {
		s.D.Set("placement_instruction", []interface{}{PlacementInstructionDetailsToMap(s.Res.PlacementInstruction)})
	} else {
		s.D.Set("placement_instruction", nil)
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

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) StartClusterPlacementGroup() error {
	request := oci_cluster_placement_groups.ActivateClusterPlacementGroupRequest{}

	idTmp := s.D.Id()
	request.ClusterPlacementGroupId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_placement_groups")

	_, err := s.Client.ActivateClusterPlacementGroup(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return s.Res.LifecycleState == oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateActive
	}
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) StopClusterPlacementGroup() error {
	request := oci_cluster_placement_groups.DeactivateClusterPlacementGroupRequest{}

	idTmp := s.D.Id()
	request.ClusterPlacementGroupId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_placement_groups")

	_, err := s.Client.DeactivateClusterPlacementGroup(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return s.Res.LifecycleState == oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateInactive
	}
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) mapToCapabilitiesCollection(fieldKeyFormat string) (oci_cluster_placement_groups.CapabilitiesCollection, error) {
	result := oci_cluster_placement_groups.CapabilitiesCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_cluster_placement_groups.CapabilityDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToCapabilityDetails(fieldKeyFormatNextLevel)
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

func CapabilitiesCollectionToMap(obj *oci_cluster_placement_groups.CapabilitiesCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, CapabilityDetailsToMap(item))
	}
	result["items"] = items

	return result
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) mapToCapabilityDetails(fieldKeyFormat string) (oci_cluster_placement_groups.CapabilityDetails, error) {
	result := oci_cluster_placement_groups.CapabilityDetails{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if service, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service")); ok {
		tmp := service.(string)
		result.Service = &tmp
	}

	return result, nil
}

func CapabilityDetailsToMap(obj oci_cluster_placement_groups.CapabilityDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Service != nil {
		result["service"] = string(*obj.Service)
	}

	return result
}

func ClusterPlacementGroupSummaryToMap(obj oci_cluster_placement_groups.ClusterPlacementGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	result["cluster_placement_group_type"] = string(obj.ClusterPlacementGroupType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

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

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) mapToPlacementInstructionDetails(fieldKeyFormat string) (oci_cluster_placement_groups.PlacementInstructionDetails, error) {
	result := oci_cluster_placement_groups.PlacementInstructionDetails{}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_cluster_placement_groups.PlacementInstructionDetailsTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func PlacementInstructionDetailsToMap(obj *oci_cluster_placement_groups.PlacementInstructionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *ClusterPlacementGroupsClusterPlacementGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cluster_placement_groups.ChangeClusterPlacementGroupCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ClusterPlacementGroupId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_placement_groups")

	_, err := s.Client.ChangeClusterPlacementGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
