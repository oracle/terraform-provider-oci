// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package disaster_recovery

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DisasterRecoveryDrProtectionGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDisasterRecoveryDrProtectionGroup,
		Read:     readDisasterRecoveryDrProtectionGroup,
		Update:   updateDisasterRecoveryDrProtectionGroup,
		Delete:   deleteDisasterRecoveryDrProtectionGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_location": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"bucket": {
							Type:     schema.TypeString,
							Required: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"object": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Optional
			"association": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"role": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"peer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"peer_region": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"members": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"member_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"member_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AUTONOMOUS_DATABASE",
								"COMPUTE_INSTANCE",
								"DATABASE",
								"VOLUME_GROUP",
							}, true),
						},

						// Optional
						"destination_compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destination_dedicated_vm_host_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_movable": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"password_vault_secret_id": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							Sensitive: true,
						},
						"vnic_mapping": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"destination_nsg_id_list": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"destination_subnet_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"source_vnic_id": {
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
			"disassociate_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"life_cycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": {
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

func createDisasterRecoveryDrProtectionGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrProtectionGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("disassociate_trigger"); ok {
		err := sync.DisassociateDrProtectionGroup()
		if err != nil {
			return err
		}
	}
	return nil

}

func readDisasterRecoveryDrProtectionGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrProtectionGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.ReadResource(sync)
}

func updateDisasterRecoveryDrProtectionGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrProtectionGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	if _, ok := sync.D.GetOkExists("disassociate_trigger"); ok && sync.D.HasChange("disassociate_trigger") {
		oldRaw, newRaw := sync.D.GetChange("disassociate_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.DisassociateDrProtectionGroup()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("disassociate_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteDisasterRecoveryDrProtectionGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrProtectionGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DisasterRecoveryDrProtectionGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_disaster_recovery.DisasterRecoveryClient
	Res                    *oci_disaster_recovery.DrProtectionGroup
	DisableNotFoundRetries bool
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_disaster_recovery.DrProtectionGroupLifecycleStateCreating),
	}
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_disaster_recovery.DrProtectionGroupLifecycleStateActive),
		string(oci_disaster_recovery.DrProtectionGroupLifecycleStateNeedsAttention),
	}
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_disaster_recovery.DrProtectionGroupLifecycleStateDeleting),
	}
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_disaster_recovery.DrProtectionGroupLifecycleStateDeleted),
	}
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) Create() error {
	request := oci_disaster_recovery.CreateDrProtectionGroupRequest{}

	if association, ok := s.D.GetOkExists("association"); ok {
		if tmpList := association.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "association", 0)
			tmp, err := s.mapToAssociateDrProtectionGroupDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Association = &tmp
		}
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

	if logLocation, ok := s.D.GetOkExists("log_location"); ok {
		if tmpList := logLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "log_location", 0)
			tmp, err := s.mapToCreateObjectStorageLogLocationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LogLocation = &tmp
		}
	}

	if members, ok := s.D.GetOkExists("members"); ok {
		interfaces := members.([]interface{})
		tmp := make([]oci_disaster_recovery.CreateDrProtectionGroupMemberDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "members", stateDataIndex)
			converted, err := s.mapToCreateDrProtectionGroupMemberDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("members") {
			request.Members = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.CreateDrProtectionGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDrProtectionGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery"), oci_disaster_recovery.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) getDrProtectionGroupFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_disaster_recovery.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	drProtectionGroupId, err := drProtectionGroupWaitForWorkRequest(workId, "drProtectionGroup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, drProtectionGroupId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_disaster_recovery.CancelWorkRequestRequest{
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
	s.D.SetId(*drProtectionGroupId)

	return s.Get()
}

func drProtectionGroupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "disaster_recovery", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_disaster_recovery.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func drProtectionGroupWaitForWorkRequest(wId *string, entityType string, action oci_disaster_recovery.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_disaster_recovery.DisasterRecoveryClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "disaster_recovery")
	retryPolicy.ShouldRetryOperation = drProtectionGroupWorkRequestShouldRetryFunc(timeout)

	response := oci_disaster_recovery.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_disaster_recovery.OperationStatusInProgress),
			string(oci_disaster_recovery.OperationStatusAccepted),
			string(oci_disaster_recovery.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_disaster_recovery.OperationStatusSucceeded),
			string(oci_disaster_recovery.OperationStatusFailed),
			string(oci_disaster_recovery.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_disaster_recovery.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_disaster_recovery.OperationStatusFailed || response.Status == oci_disaster_recovery.OperationStatusCanceled {
		return nil, getErrorFromDisasterRecoveryDrProtectionGroupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDisasterRecoveryDrProtectionGroupWorkRequest(client *oci_disaster_recovery.DisasterRecoveryClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_disaster_recovery.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_disaster_recovery.ListWorkRequestErrorsRequest{
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

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) Get() error {
	request := oci_disaster_recovery.GetDrProtectionGroupRequest{}

	tmp := s.D.Id()
	request.DrProtectionGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.GetDrProtectionGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrProtectionGroup
	return nil
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_disaster_recovery.UpdateDrProtectionGroupRequest{}

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

	tmp := s.D.Id()
	request.DrProtectionGroupId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if logLocation, ok := s.D.GetOkExists("log_location"); ok {
		if tmpList := logLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "log_location", 0)
			tmp, err := s.mapToUpdateObjectStorageLogLocationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LogLocation = &tmp
		}
	}

	if members, ok := s.D.GetOkExists("members"); ok {
		interfaces := members.([]interface{})
		tmp := make([]oci_disaster_recovery.UpdateDrProtectionGroupMemberDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "members", stateDataIndex)
			converted, err := s.mapToUpdateDrProtectionGroupMemberDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("members") {
			request.Members = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.UpdateDrProtectionGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDrProtectionGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery"), oci_disaster_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) Delete() error {
	request := oci_disaster_recovery.DeleteDrProtectionGroupRequest{}

	tmp := s.D.Id()
	request.DrProtectionGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.DeleteDrProtectionGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := drProtectionGroupWaitForWorkRequest(workId, "drProtectionGroup",
		oci_disaster_recovery.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) SetData() error {
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
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifeCycleDetails != nil {
		s.D.Set("life_cycle_details", *s.Res.LifeCycleDetails)
	}

	if s.Res.LogLocation != nil {
		s.D.Set("log_location", []interface{}{ObjectStorageLogLocationToMap(s.Res.LogLocation)})
	} else {
		s.D.Set("log_location", nil)
	}

	members := []interface{}{}
	for _, item := range s.Res.Members {
		members = append(members, DrProtectionGroupMemberToMap(item))
	}
	s.D.Set("members", members)

	if s.Res.PeerId != nil {
		s.D.Set("peer_id", *s.Res.PeerId)
	}

	if s.Res.PeerRegion != nil {
		s.D.Set("peer_region", *s.Res.PeerRegion)
	}

	s.D.Set("role", s.Res.Role)

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

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) DisassociateDrProtectionGroup() error {
	request := oci_disaster_recovery.DisassociateDrProtectionGroupRequest{}

	idTmp := s.D.Id()
	request.DrProtectionGroupId = &idTmp

	request.DisassociateDrProtectionGroupDetails = oci_disaster_recovery.DisassociateDrProtectionGroupDefaultDetails{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.DisassociateDrProtectionGroup(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	workId := response.OpcWorkRequestId
	s.getDrProtectionGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery"), oci_disaster_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))

	val := s.D.Get("disassociate_trigger")
	s.D.Set("disassociate_trigger", val)

	//s.Res = &response.DrProtectionGroup
	return nil
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToAssociateDrProtectionGroupDetails(fieldKeyFormat string) (oci_disaster_recovery.AssociateDrProtectionGroupDetails, error) {
	result := oci_disaster_recovery.AssociateDrProtectionGroupDetails{}

	if peerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_id")); ok {
		tmp := peerId.(string)
		result.PeerId = &tmp
	}

	if peerRegion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_region")); ok {
		tmp := peerRegion.(string)
		result.PeerRegion = &tmp
	}

	if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
		result.Role = oci_disaster_recovery.DrProtectionGroupRoleEnum(role.(string))
	}

	return result, nil
}

func AssociateDrProtectionGroupDetailsToMap(obj *oci_disaster_recovery.AssociateDrProtectionGroupDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PeerId != nil {
		result["peer_id"] = string(*obj.PeerId)
	}

	if obj.PeerRegion != nil {
		result["peer_region"] = string(*obj.PeerRegion)
	}

	result["role"] = string(obj.Role)

	return result
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToComputeInstanceVnicMappingDetails(fieldKeyFormat string) (oci_disaster_recovery.ComputeInstanceVnicMappingDetails, error) {
	result := oci_disaster_recovery.ComputeInstanceVnicMappingDetails{}

	if destinationNsgIdList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_nsg_id_list")); ok {
		interfaces := destinationNsgIdList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "destination_nsg_id_list")) {
			result.DestinationNsgIdList = tmp
		}
	}

	if destinationSubnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_subnet_id")); ok {
		tmp := destinationSubnetId.(string)
		result.DestinationSubnetId = &tmp
	}

	if sourceVnicId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_vnic_id")); ok {
		tmp := sourceVnicId.(string)
		result.SourceVnicId = &tmp
	}

	return result, nil
}

func ComputeInstanceVnicMappingToMap(obj oci_disaster_recovery.ComputeInstanceVnicMapping) map[string]interface{} {
	result := map[string]interface{}{}

	result["destination_nsg_id_list"] = obj.DestinationNsgIdList
	result["destination_nsg_id_list"] = obj.DestinationNsgIdList

	if obj.DestinationSubnetId != nil {
		result["destination_subnet_id"] = string(*obj.DestinationSubnetId)
	}

	if obj.SourceVnicId != nil {
		result["source_vnic_id"] = string(*obj.SourceVnicId)
	}

	return result
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToCreateDrProtectionGroupMemberDetails(fieldKeyFormat string) (oci_disaster_recovery.CreateDrProtectionGroupMemberDetails, error) {
	var baseObject oci_disaster_recovery.CreateDrProtectionGroupMemberDetails
	//discriminator
	memberTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_type"))
	var memberType string
	if ok {
		memberType = memberTypeRaw.(string)
	} else {
		memberType = "" // default value
	}
	switch strings.ToLower(memberType) {
	case strings.ToLower("AUTONOMOUS_DATABASE"):
		details := oci_disaster_recovery.CreateDrProtectionGroupMemberAutonomousDatabaseDetails{}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("COMPUTE_INSTANCE"):
		details := oci_disaster_recovery.CreateDrProtectionGroupMemberComputeInstanceDetails{}
		if destinationCompartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_compartment_id")); ok {
			tmp := destinationCompartmentId.(string)
			details.DestinationCompartmentId = &tmp
		}
		if destinationDedicatedVmHostId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_dedicated_vm_host_id")); ok {
			tmp := destinationDedicatedVmHostId.(string)
			details.DestinationDedicatedVmHostId = &tmp
		}
		if isMovable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_movable")); ok {
			tmp := isMovable.(bool)
			details.IsMovable = &tmp
		}
		if vnicMapping, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vnic_mapping")); ok {
			interfaces := vnicMapping.([]interface{})
			tmp := make([]oci_disaster_recovery.ComputeInstanceVnicMappingDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "vnic_mapping"), stateDataIndex)
				converted, err := s.mapToComputeInstanceVnicMappingDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "vnic_mapping")) {
				details.VnicMapping = tmp
			}
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("DATABASE"):
		details := oci_disaster_recovery.CreateDrProtectionGroupMemberDatabaseDetails{}
		if passwordVaultSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_vault_secret_id")); ok {
			tmp := passwordVaultSecretId.(string)
			details.PasswordVaultSecretId = &tmp
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("VOLUME_GROUP"):
		details := oci_disaster_recovery.CreateDrProtectionGroupMemberVolumeGroupDetails{}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown member_type '%v' was specified", memberType)
	}
	return baseObject, nil
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToUpdateDrProtectionGroupMemberDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateDrProtectionGroupMemberDetails, error) {
	var baseObject oci_disaster_recovery.UpdateDrProtectionGroupMemberDetails
	//discriminator
	memberTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_type"))
	var memberType string
	if ok {
		memberType = memberTypeRaw.(string)
	} else {
		memberType = "" // default value
	}
	switch strings.ToLower(memberType) {
	case strings.ToLower("AUTONOMOUS_DATABASE"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberAutonomousDatabaseDetails{}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("COMPUTE_INSTANCE"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberComputeInstanceDetails{}
		if destinationCompartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_compartment_id")); ok {
			tmp := destinationCompartmentId.(string)
			details.DestinationCompartmentId = &tmp
		}
		if destinationDedicatedVmHostId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_dedicated_vm_host_id")); ok {
			tmp := destinationDedicatedVmHostId.(string)
			details.DestinationDedicatedVmHostId = &tmp
		}
		if isMovable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_movable")); ok {
			tmp := isMovable.(bool)
			details.IsMovable = &tmp
		}
		if vnicMapping, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vnic_mapping")); ok {
			interfaces := vnicMapping.([]interface{})
			tmp := make([]oci_disaster_recovery.ComputeInstanceVnicMappingDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "vnic_mapping"), stateDataIndex)
				converted, err := s.mapToComputeInstanceVnicMappingDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "vnic_mapping")) {
				details.VnicMapping = tmp
			}
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("DATABASE"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberDatabaseDetails{}
		if passwordVaultSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_vault_secret_id")); ok {
			tmp := passwordVaultSecretId.(string)
			details.PasswordVaultSecretId = &tmp
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("VOLUME_GROUP"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberVolumeGroupDetails{}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown member_type '%v' was specified", memberType)
	}
	return baseObject, nil
}

func DrProtectionGroupMemberToMap(obj oci_disaster_recovery.DrProtectionGroupMember) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_disaster_recovery.DrProtectionGroupMemberAutonomousDatabase:
		result["member_type"] = "AUTONOMOUS_DATABASE"

		if v.MemberId != nil {
			result["member_id"] = string(*v.MemberId)
		}
	case oci_disaster_recovery.DrProtectionGroupMemberComputeInstance:
		result["member_type"] = "COMPUTE_INSTANCE"

		if v.DestinationCompartmentId != nil {
			result["destination_compartment_id"] = string(*v.DestinationCompartmentId)
		}

		if v.DestinationDedicatedVmHostId != nil {
			result["destination_dedicated_vm_host_id"] = string(*v.DestinationDedicatedVmHostId)
		}

		if v.IsMovable != nil {
			result["is_movable"] = bool(*v.IsMovable)
		}

		vnicMapping := []interface{}{}
		for _, item := range v.VnicMapping {
			vnicMapping = append(vnicMapping, ComputeInstanceVnicMappingToMap(item))
		}
		result["vnic_mapping"] = vnicMapping

		if v.MemberId != nil {
			result["member_id"] = string(*v.MemberId)
		}
	case oci_disaster_recovery.DrProtectionGroupMemberDatabase:
		result["member_type"] = "DATABASE"

		if v.PasswordVaultSecretId != nil {
			result["password_vault_secret_id"] = string(*v.PasswordVaultSecretId)
		}

		if v.MemberId != nil {
			result["member_id"] = string(*v.MemberId)
		}
	case oci_disaster_recovery.DrProtectionGroupMemberVolumeGroup:
		result["member_type"] = "VOLUME_GROUP"

		if v.MemberId != nil {
			result["member_id"] = string(*v.MemberId)
		}
	default:
		log.Printf("[WARN] Received 'member_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToCreateObjectStorageLogLocationDetails(fieldKeyFormat string) (oci_disaster_recovery.CreateObjectStorageLogLocationDetails, error) {
	result := oci_disaster_recovery.CreateObjectStorageLogLocationDetails{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.Bucket = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	return result, nil
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToUpdateObjectStorageLogLocationDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateObjectStorageLogLocationDetails, error) {
	result := oci_disaster_recovery.UpdateObjectStorageLogLocationDetails{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.Bucket = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	return result, nil
}

func ObjectStorageLogLocationToMap(obj *oci_disaster_recovery.ObjectStorageLogLocation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Bucket != nil {
		result["bucket"] = string(*obj.Bucket)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.Object != nil {
		result["object"] = string(*obj.Object)
	}

	return result
}

func DrProtectionGroupSummaryToMap(obj oci_disaster_recovery.DrProtectionGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifeCycleDetails != nil {
		result["life_cycle_details"] = string(*obj.LifeCycleDetails)
	}

	if obj.PeerId != nil {
		result["peer_id"] = string(*obj.PeerId)
	}

	if obj.PeerRegion != nil {
		result["peer_region"] = string(*obj.PeerRegion)
	}

	result["role"] = string(obj.Role)

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

/*func (s *DisasterRecoveryDrProtectionGroupResourceCrud) populateTopLevelPolymorphicUpdateDrProtectionGroupRequest(request *oci_disaster_recovery.UpdateDrProtectionGroupRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DEFAULT"):
		details := oci_disaster_recovery.DisassociateDrProtectionGroupDefaultDetails{}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		tmp := s.D.Id()
		request.DrProtectionGroupId = &tmp
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if logLocation, ok := s.D.GetOkExists("log_location"); ok {
			if tmpList := logLocation.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "log_location", 0)
				tmp, err := s.mapToUpdateObjectStorageLogLocationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.LogLocation = &tmp
			}
		}
		if members, ok := s.D.GetOkExists("members"); ok {
			interfaces := members.([]interface{})
			tmp := make([]oci_disaster_recovery.UpdateDrProtectionGroupMemberDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "members", stateDataIndex)
				converted, err := s.mapToUpdateDrProtectionGroupMemberDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("members") {
				details.Members = tmp
			}
		}
		request.DisassociateDrProtectionGroupDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}*/

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_disaster_recovery.ChangeDrProtectionGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DrProtectionGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.ChangeDrProtectionGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDrProtectionGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery"), oci_disaster_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
