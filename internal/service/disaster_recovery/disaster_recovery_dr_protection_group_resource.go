// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
								"COMPUTE_INSTANCE_MOVABLE",
								"COMPUTE_INSTANCE_NON_MOVABLE",
								"DATABASE",
								"FILE_SYSTEM",
								"LOAD_BALANCER",
								"NETWORK_LOAD_BALANCER",
								"VOLUME_GROUP",
							}, true),
						},

						// Optional
						"backend_set_mappings": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"destination_backend_set_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"is_backend_set_for_non_movable": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"source_backend_set_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"block_volume_operations": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"attachment_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"volume_attachment_reference_instance_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"block_volume_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mount_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"mount_point": {
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
						"destination_availability_domain": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destination_capacity_reservation_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
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
						"destination_load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destination_network_load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"export_mappings": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"destination_mount_target_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"export_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"file_system_operations": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"export_path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mount_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"mount_target_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"mount_point": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mount_target_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"unmount_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"mount_target_id": {
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
						"is_movable": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_retain_fault_domain": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_start_stop_enabled": {
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
						"vnic_mappings": {
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
									"destination_primary_private_ip_address": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"destination_primary_private_ip_hostname_label": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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
			"lifecycle_sub_state": {
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
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
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

	s.D.Set("lifecycle_sub_state", s.Res.LifecycleSubState)

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
	err = s.getDrProtectionGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery"), oci_disaster_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
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

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToComputeInstanceMovableVnicMappingDetails(fieldKeyFormat string) (oci_disaster_recovery.ComputeInstanceMovableVnicMappingDetails, error) {
	result := oci_disaster_recovery.ComputeInstanceMovableVnicMappingDetails{}

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

	if destinationPrimaryPrivateIpAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_primary_private_ip_address")); ok {
		tmp := destinationPrimaryPrivateIpAddress.(string)
		result.DestinationPrimaryPrivateIpAddress = &tmp
	}

	if destinationPrimaryPrivateIpHostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_primary_private_ip_hostname_label")); ok {
		tmp := destinationPrimaryPrivateIpHostnameLabel.(string)
		result.DestinationPrimaryPrivateIpHostnameLabel = &tmp
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

func ComputeInstanceMovableVnicMappingDetailsToMap(obj oci_disaster_recovery.ComputeInstanceMovableVnicMapping) map[string]interface{} {
	result := map[string]interface{}{}

	result["destination_nsg_id_list"] = obj.DestinationNsgIdList

	if obj.DestinationPrimaryPrivateIpAddress != nil {
		result["destination_primary_private_ip_address"] = string(*obj.DestinationPrimaryPrivateIpAddress)
	}

	if obj.DestinationPrimaryPrivateIpHostnameLabel != nil {
		result["destination_primary_private_ip_hostname_label"] = string(*obj.DestinationPrimaryPrivateIpHostnameLabel)
	}

	if obj.DestinationSubnetId != nil {
		result["destination_subnet_id"] = string(*obj.DestinationSubnetId)
	}

	if obj.SourceVnicId != nil {
		result["source_vnic_id"] = string(*obj.SourceVnicId)
	}

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

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToCreateBlockVolumeAttachmentDetails(fieldKeyFormat string) (oci_disaster_recovery.CreateBlockVolumeAttachmentDetails, error) {
	result := oci_disaster_recovery.CreateBlockVolumeAttachmentDetails{}

	if volumeAttachmentReferenceInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_attachment_reference_instance_id")); ok {
		tmp := volumeAttachmentReferenceInstanceId.(string)
		result.VolumeAttachmentReferenceInstanceId = &tmp
	}

	return result, nil
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToUpdateBlockVolumeAttachmentDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateBlockVolumeAttachmentDetails, error) {
	result := oci_disaster_recovery.UpdateBlockVolumeAttachmentDetails{}

	if volumeAttachmentReferenceInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_attachment_reference_instance_id")); ok {
		tmp := volumeAttachmentReferenceInstanceId.(string)
		result.VolumeAttachmentReferenceInstanceId = &tmp
	}

	return result, nil
}

func BlockVolumeAttachmentDetailsToMap(obj *oci_disaster_recovery.BlockVolumeAttachmentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.VolumeAttachmentReferenceInstanceId != nil {
		result["volume_attachment_reference_instance_id"] = string(*obj.VolumeAttachmentReferenceInstanceId)
	}

	return result
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToCreateBlockVolumeMountDetails(fieldKeyFormat string) (oci_disaster_recovery.CreateBlockVolumeMountDetails, error) {
	result := oci_disaster_recovery.CreateBlockVolumeMountDetails{}

	if mountPoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_point")); ok {
		tmp := mountPoint.(string)
		result.MountPoint = &tmp
	}

	return result, nil
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToUpdateBlockVolumeMountDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateBlockVolumeMountDetails, error) {
	result := oci_disaster_recovery.UpdateBlockVolumeMountDetails{}

	if mountPoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_point")); ok {
		tmp := mountPoint.(string)
		result.MountPoint = &tmp
	}

	return result, nil
}

func BlockVolumeMountDetailsToMap(obj *oci_disaster_recovery.BlockVolumeMountDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MountPoint != nil {
		result["mount_point"] = string(*obj.MountPoint)
	}

	return result
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToCreateComputeInstanceMovableFileSystemOperationDetails(fieldKeyFormat string) (oci_disaster_recovery.CreateComputeInstanceMovableFileSystemOperationDetails, error) {
	result := oci_disaster_recovery.CreateComputeInstanceMovableFileSystemOperationDetails{}

	if exportPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_path")); ok {
		tmp := exportPath.(string)
		result.ExportPath = &tmp
	}

	if mountDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_details")); ok {
		if tmpList := mountDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "mount_details"), 0)
			tmp, err := s.mapToCreateFileSystemMountDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert mount_details, encountered error: %v", err)
			}
			result.MountDetails = &tmp
		}
	}

	if mountPoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_point")); ok {
		tmp := mountPoint.(string)
		result.MountPoint = &tmp
	}

	if unmountDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unmount_details")); ok {
		if tmpList := unmountDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "unmount_details"), 0)
			tmp, err := s.mapToCreateFileSystemUnmountDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert unmount_details, encountered error: %v", err)
			}
			result.UnmountDetails = &tmp
		}
	}

	return result, nil
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToUpdateComputeInstanceMovableFileSystemOperationDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateComputeInstanceMovableFileSystemOperationDetails, error) {
	result := oci_disaster_recovery.UpdateComputeInstanceMovableFileSystemOperationDetails{}

	if exportPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_path")); ok {
		tmp := exportPath.(string)
		result.ExportPath = &tmp
	}

	if mountDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_details")); ok {
		if tmpList := mountDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "mount_details"), 0)
			tmp, err := s.mapToUpdateFileSystemMountDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert mount_details, encountered error: %v", err)
			}
			result.MountDetails = &tmp
		}
	}

	if mountPoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_point")); ok {
		tmp := mountPoint.(string)
		result.MountPoint = &tmp
	}

	if unmountDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unmount_details")); ok {
		if tmpList := unmountDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "unmount_details"), 0)
			tmp, err := s.mapToUpdateFileSystemUnmountDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert unmount_details, encountered error: %v", err)
			}
			result.UnmountDetails = &tmp
		}
	}

	return result, nil
}

func ComputeInstanceMovableFileSystemOperationToMap(obj oci_disaster_recovery.ComputeInstanceMovableFileSystemOperation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExportPath != nil {
		result["export_path"] = string(*obj.ExportPath)
	}

	if obj.MountDetails != nil {
		result["mount_details"] = []interface{}{FileSystemMountDetailsToMap(obj.MountDetails)}
	}

	if obj.MountPoint != nil {
		result["mount_point"] = string(*obj.MountPoint)
	}

	if obj.UnmountDetails != nil {
		result["unmount_details"] = []interface{}{FileSystemUnmountDetailsToMap(obj.UnmountDetails)}
	}

	return result
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToCreateComputeInstanceNonMovableBlockVolumeOperationDetails(fieldKeyFormat string) (oci_disaster_recovery.CreateComputeInstanceNonMovableBlockVolumeOperationDetails, error) {
	result := oci_disaster_recovery.CreateComputeInstanceNonMovableBlockVolumeOperationDetails{}

	if attachmentDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attachment_details")); ok {
		if tmpList := attachmentDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "attachment_details"), 0)
			tmp, err := s.mapToCreateBlockVolumeAttachmentDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert attachment_details, encountered error: %v", err)
			}
			result.AttachmentDetails = &tmp
		}
	}

	if blockVolumeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_id")); ok {
		tmp := blockVolumeId.(string)
		result.BlockVolumeId = &tmp
	}

	if mountDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_details")); ok {
		if tmpList := mountDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "mount_details"), 0)
			tmp, err := s.mapToCreateBlockVolumeMountDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert mount_details, encountered error: %v", err)
			}
			result.MountDetails = &tmp
		}
	}

	return result, nil
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToUpdateComputeInstanceNonMovableBlockVolumeOperationDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateComputeInstanceNonMovableBlockVolumeOperationDetails, error) {
	result := oci_disaster_recovery.UpdateComputeInstanceNonMovableBlockVolumeOperationDetails{}

	if attachmentDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attachment_details")); ok {
		if tmpList := attachmentDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "attachment_details"), 0)
			tmp, err := s.mapToUpdateBlockVolumeAttachmentDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert attachment_details, encountered error: %v", err)
			}
			result.AttachmentDetails = &tmp
		}
	}

	if blockVolumeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_id")); ok {
		tmp := blockVolumeId.(string)
		result.BlockVolumeId = &tmp
	}

	if mountDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_details")); ok {
		if tmpList := mountDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "mount_details"), 0)
			tmp, err := s.mapToUpdateBlockVolumeMountDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert mount_details, encountered error: %v", err)
			}
			result.MountDetails = &tmp
		}
	}

	return result, nil
}

func ComputeInstanceNonMovableBlockVolumeOperationToMap(obj oci_disaster_recovery.ComputeInstanceNonMovableBlockVolumeOperation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AttachmentDetails != nil {
		result["attachment_details"] = []interface{}{BlockVolumeAttachmentDetailsToMap(obj.AttachmentDetails)}
	}

	if obj.BlockVolumeId != nil {
		result["block_volume_id"] = string(*obj.BlockVolumeId)
	}

	if obj.MountDetails != nil {
		result["mount_details"] = []interface{}{BlockVolumeMountDetailsToMap(obj.MountDetails)}
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
	case strings.ToLower("COMPUTE_INSTANCE_MOVABLE"):
		details := oci_disaster_recovery.CreateDrProtectionGroupMemberComputeInstanceMovableDetails{}
		if destinationCapacityReservationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_capacity_reservation_id")); ok {
			tmp := destinationCapacityReservationId.(string)
			details.DestinationCapacityReservationId = &tmp
		}
		if destinationCompartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_compartment_id")); ok {
			tmp := destinationCompartmentId.(string)
			details.DestinationCompartmentId = &tmp
		}
		if destinationDedicatedVmHostId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_dedicated_vm_host_id")); ok {
			tmp := destinationDedicatedVmHostId.(string)
			details.DestinationDedicatedVmHostId = &tmp
		}
		if fileSystemOperations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_system_operations")); ok {
			interfaces := fileSystemOperations.([]interface{})
			tmp := make([]oci_disaster_recovery.CreateComputeInstanceMovableFileSystemOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_system_operations"), stateDataIndex)
				converted, err := s.mapToCreateComputeInstanceMovableFileSystemOperationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "file_system_operations")) {
				details.FileSystemOperations = tmp
			}
		}
		if isRetainFaultDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_retain_fault_domain")); ok {
			tmp := isRetainFaultDomain.(bool)
			details.IsRetainFaultDomain = &tmp
		}
		if vnicMappings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vnic_mappings")); ok {
			interfaces := vnicMappings.([]interface{})
			tmp := make([]oci_disaster_recovery.ComputeInstanceMovableVnicMappingDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "vnic_mappings"), stateDataIndex)
				converted, err := s.mapToComputeInstanceMovableVnicMappingDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "vnic_mappings")) {
				details.VnicMappings = tmp
			}
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("COMPUTE_INSTANCE_NON_MOVABLE"):
		details := oci_disaster_recovery.CreateDrProtectionGroupMemberComputeInstanceNonMovableDetails{}
		if blockVolumeOperations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_operations")); ok {
			interfaces := blockVolumeOperations.([]interface{})
			tmp := make([]oci_disaster_recovery.CreateComputeInstanceNonMovableBlockVolumeOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "block_volume_operations"), stateDataIndex)
				converted, err := s.mapToCreateComputeInstanceNonMovableBlockVolumeOperationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "block_volume_operations")) {
				details.BlockVolumeOperations = tmp
			}
		}
		if fileSystemOperations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_system_operations")); ok {
			interfaces := fileSystemOperations.([]interface{})
			tmp := make([]oci_disaster_recovery.CreateComputeInstanceNonMovableFileSystemOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_system_operations"), stateDataIndex)
				converted, err := s.mapToCreateComputeInstanceNonMovableFileSystemOperationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "file_system_operations")) {
				details.FileSystemOperations = tmp
			}
		}
		if isStartStopEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_start_stop_enabled")); ok {
			tmp := isStartStopEnabled.(bool)
			details.IsStartStopEnabled = &tmp
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
	case strings.ToLower("FILE_SYSTEM"):
		details := oci_disaster_recovery.CreateDrProtectionGroupMemberFileSystemDetails{}
		if destinationAvailabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_availability_domain")); ok {
			tmp := destinationAvailabilityDomain.(string)
			details.DestinationAvailabilityDomain = &tmp
		}
		if exportMappings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_mappings")); ok {
			interfaces := exportMappings.([]interface{})
			tmp := make([]oci_disaster_recovery.FileSystemExportMappingDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "export_mappings"), stateDataIndex)
				converted, err := s.mapToFileSystemExportMappingDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "export_mappings")) {
				details.ExportMappings = tmp
			}
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("LOAD_BALANCER"):
		details := oci_disaster_recovery.CreateDrProtectionGroupMemberLoadBalancerDetails{}
		if backendSetMappings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backend_set_mappings")); ok {
			interfaces := backendSetMappings.([]interface{})
			tmp := make([]oci_disaster_recovery.LoadBalancerBackendSetMappingDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backend_set_mappings"), stateDataIndex)
				converted, err := s.mapToLoadBalancerBackendSetMappingDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backend_set_mappings")) {
				details.BackendSetMappings = tmp
			}
		}
		if destinationLoadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_load_balancer_id")); ok {
			tmp := destinationLoadBalancerId.(string)
			details.DestinationLoadBalancerId = &tmp
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("NETWORK_LOAD_BALANCER"):
		details := oci_disaster_recovery.CreateDrProtectionGroupMemberNetworkLoadBalancerDetails{}
		if backendSetMappings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backend_set_mappings")); ok {
			interfaces := backendSetMappings.([]interface{})
			tmp := make([]oci_disaster_recovery.NetworkLoadBalancerBackendSetMappingDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backend_set_mappings"), stateDataIndex)
				converted, err := s.mapToNetworkLoadBalancerBackendSetMappingDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backend_set_mappings")) {
				details.BackendSetMappings = tmp
			}
		}
		if destinationNetworkLoadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_network_load_balancer_id")); ok {
			tmp := destinationNetworkLoadBalancerId.(string)
			details.DestinationNetworkLoadBalancerId = &tmp
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
			if tmp != "" {
				details.MemberId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("COMPUTE_INSTANCE"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberComputeInstanceDetails{}
		if destinationCompartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_compartment_id")); ok {
			tmp := destinationCompartmentId.(string)
			if tmp != "" {
				details.DestinationCompartmentId = &tmp
			}
		}
		if destinationDedicatedVmHostId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_dedicated_vm_host_id")); ok {
			tmp := destinationDedicatedVmHostId.(string)
			if tmp != "" {
				details.DestinationDedicatedVmHostId = &tmp
			}
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
			if tmp != "" {
				details.MemberId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("COMPUTE_INSTANCE_MOVABLE"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberComputeInstanceMovableDetails{}
		if destinationCapacityReservationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_capacity_reservation_id")); ok {
			tmp := destinationCapacityReservationId.(string)
			if tmp != "" {
				details.DestinationCapacityReservationId = &tmp
			}
		}
		if destinationCompartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_compartment_id")); ok {
			tmp := destinationCompartmentId.(string)
			if tmp != "" {
				details.DestinationCompartmentId = &tmp
			}
		}
		if destinationDedicatedVmHostId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_dedicated_vm_host_id")); ok {
			tmp := destinationDedicatedVmHostId.(string)
			if tmp != "" {
				details.DestinationDedicatedVmHostId = &tmp
			}
		}
		if fileSystemOperations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_system_operations")); ok {
			interfaces := fileSystemOperations.([]interface{})
			tmp := make([]oci_disaster_recovery.UpdateComputeInstanceMovableFileSystemOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_system_operations"), stateDataIndex)
				converted, err := s.mapToUpdateComputeInstanceMovableFileSystemOperationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "file_system_operations")) {
				details.FileSystemOperations = tmp
			}
		}
		if isRetainFaultDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_retain_fault_domain")); ok {
			tmp := isRetainFaultDomain.(bool)
			details.IsRetainFaultDomain = &tmp
		}
		if vnicMappings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vnic_mappings")); ok {
			interfaces := vnicMappings.([]interface{})
			tmp := make([]oci_disaster_recovery.ComputeInstanceMovableVnicMappingDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "vnic_mappings"), stateDataIndex)
				converted, err := s.mapToComputeInstanceMovableVnicMappingDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "vnic_mappings")) {
				details.VnicMappings = tmp
			}
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			if tmp != "" {
				details.MemberId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("COMPUTE_INSTANCE_NON_MOVABLE"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberComputeInstanceNonMovableDetails{}
		if blockVolumeOperations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_operations")); ok {
			interfaces := blockVolumeOperations.([]interface{})
			tmp := make([]oci_disaster_recovery.UpdateComputeInstanceNonMovableBlockVolumeOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "block_volume_operations"), stateDataIndex)
				converted, err := s.mapToUpdateComputeInstanceNonMovableBlockVolumeOperationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "block_volume_operations")) {
				details.BlockVolumeOperations = tmp
			}
		}
		if fileSystemOperations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_system_operations")); ok {
			interfaces := fileSystemOperations.([]interface{})
			tmp := make([]oci_disaster_recovery.UpdateComputeInstanceNonMovableFileSystemOperationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_system_operations"), stateDataIndex)
				converted, err := s.mapToUpdateComputeInstanceNonMovableFileSystemOperationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "file_system_operations")) {
				details.FileSystemOperations = tmp
			}
		}
		if isStartStopEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_start_stop_enabled")); ok {
			tmp := isStartStopEnabled.(bool)
			details.IsStartStopEnabled = &tmp
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			if tmp != "" {
				details.MemberId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("DATABASE"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberDatabaseDetails{}
		if passwordVaultSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_vault_secret_id")); ok {
			tmp := passwordVaultSecretId.(string)
			if tmp != "" {
				details.PasswordVaultSecretId = &tmp
			}
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			if tmp != "" {
				details.MemberId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("FILE_SYSTEM"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberFileSystemDetails{}
		if destinationAvailabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_availability_domain")); ok {
			tmp := destinationAvailabilityDomain.(string)
			if tmp != "" {
				details.DestinationAvailabilityDomain = &tmp
			}
		}
		if exportMappings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_mappings")); ok {
			interfaces := exportMappings.([]interface{})
			tmp := make([]oci_disaster_recovery.FileSystemExportMappingDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "export_mappings"), stateDataIndex)
				converted, err := s.mapToFileSystemExportMappingDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "export_mappings")) {
				details.ExportMappings = tmp
			}
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			if tmp != "" {
				details.MemberId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("LOAD_BALANCER"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberLoadBalancerDetails{}
		if backendSetMappings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backend_set_mappings")); ok {
			interfaces := backendSetMappings.([]interface{})
			tmp := make([]oci_disaster_recovery.LoadBalancerBackendSetMappingDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backend_set_mappings"), stateDataIndex)
				converted, err := s.mapToLoadBalancerBackendSetMappingDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backend_set_mappings")) {
				details.BackendSetMappings = tmp
			}
		}
		if destinationLoadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_load_balancer_id")); ok {
			tmp := destinationLoadBalancerId.(string)
			if tmp != "" {
				details.DestinationLoadBalancerId = &tmp
			}
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			if tmp != "" {
				details.MemberId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("NETWORK_LOAD_BALANCER"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberNetworkLoadBalancerDetails{}
		if backendSetMappings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backend_set_mappings")); ok {
			interfaces := backendSetMappings.([]interface{})
			tmp := make([]oci_disaster_recovery.NetworkLoadBalancerBackendSetMappingDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backend_set_mappings"), stateDataIndex)
				converted, err := s.mapToNetworkLoadBalancerBackendSetMappingDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backend_set_mappings")) {
				details.BackendSetMappings = tmp
			}
		}
		if destinationNetworkLoadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_network_load_balancer_id")); ok {
			tmp := destinationNetworkLoadBalancerId.(string)
			if tmp != "" {
				details.DestinationNetworkLoadBalancerId = &tmp
			}
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			if tmp != "" {
				details.MemberId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("VOLUME_GROUP"):
		details := oci_disaster_recovery.UpdateDrProtectionGroupMemberVolumeGroupDetails{}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			if tmp != "" {
				details.MemberId = &tmp
			}
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
	case oci_disaster_recovery.DrProtectionGroupMemberComputeInstanceMovable:
		result["member_type"] = "COMPUTE_INSTANCE_MOVABLE"

		if v.DestinationCapacityReservationId != nil {
			result["destination_capacity_reservation_id"] = string(*v.DestinationCapacityReservationId)
		}

		if v.DestinationCompartmentId != nil {
			result["destination_compartment_id"] = string(*v.DestinationCompartmentId)
		}

		if v.DestinationDedicatedVmHostId != nil {
			result["destination_dedicated_vm_host_id"] = string(*v.DestinationDedicatedVmHostId)
		}

		fileSystemOperations := []interface{}{}
		for _, item := range v.FileSystemOperations {
			fileSystemOperations = append(fileSystemOperations, ComputeInstanceMovableFileSystemOperationToMap(item))
		}
		result["file_system_operations"] = fileSystemOperations

		if v.IsRetainFaultDomain != nil {
			result["is_retain_fault_domain"] = bool(*v.IsRetainFaultDomain)
		}

		vnicMappings := []interface{}{}
		for _, item := range v.VnicMappings {
			vnicMappings = append(vnicMappings, ComputeInstanceMovableVnicMappingDetailsToMap(item))
		}
		result["vnic_mappings"] = vnicMappings

		if v.MemberId != nil {
			result["member_id"] = string(*v.MemberId)
		}
	case oci_disaster_recovery.DrProtectionGroupMemberComputeInstanceNonMovable:
		result["member_type"] = "COMPUTE_INSTANCE_NON_MOVABLE"

		blockVolumeOperations := []interface{}{}
		for _, item := range v.BlockVolumeOperations {
			blockVolumeOperations = append(blockVolumeOperations, ComputeInstanceNonMovableBlockVolumeOperationToMap(item))
		}
		result["block_volume_operations"] = blockVolumeOperations

		fileSystemOperations := []interface{}{}
		for _, item := range v.FileSystemOperations {
			fileSystemOperations = append(fileSystemOperations, ComputeInstanceNonMovableFileSystemOperationToMap(item))
		}
		result["file_system_operations"] = fileSystemOperations

		if v.IsStartStopEnabled != nil {
			result["is_start_stop_enabled"] = bool(*v.IsStartStopEnabled)
		}

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
	case oci_disaster_recovery.DrProtectionGroupMemberFileSystem:
		result["member_type"] = "FILE_SYSTEM"

		if v.DestinationAvailabilityDomain != nil {
			result["destination_availability_domain"] = string(*v.DestinationAvailabilityDomain)
		}

		exportMappings := []interface{}{}
		for _, item := range v.ExportMappings {
			exportMappings = append(exportMappings, FileSystemExportMappingDetailsToMap(item))
		}
		result["export_mappings"] = exportMappings

		if v.MemberId != nil {
			result["member_id"] = string(*v.MemberId)
		}
	case oci_disaster_recovery.DrProtectionGroupMemberLoadBalancer:
		result["member_type"] = "LOAD_BALANCER"

		backendSetMappings := []interface{}{}
		for _, item := range v.BackendSetMappings {
			backendSetMappings = append(backendSetMappings, LoadBalancerBackendSetMappingDetailsToMap(item))
		}
		result["backend_set_mappings"] = backendSetMappings

		if v.DestinationLoadBalancerId != nil {
			result["destination_load_balancer_id"] = string(*v.DestinationLoadBalancerId)
		}

		if v.MemberId != nil {
			result["member_id"] = string(*v.MemberId)
		}
	case oci_disaster_recovery.DrProtectionGroupMemberNetworkLoadBalancer:
		result["member_type"] = "NETWORK_LOAD_BALANCER"

		backendSetMappings := []interface{}{}
		for _, item := range v.BackendSetMappings {
			backendSetMappings = append(backendSetMappings, NetworkLoadBalancerBackendSetMappingDetailsToMap(item))
		}
		result["backend_set_mappings"] = backendSetMappings

		if v.DestinationNetworkLoadBalancerId != nil {
			result["destination_network_load_balancer_id"] = string(*v.DestinationNetworkLoadBalancerId)
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

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToCreateFileSystemMountDetails(fieldKeyFormat string) (oci_disaster_recovery.CreateFileSystemMountDetails, error) {
	result := oci_disaster_recovery.CreateFileSystemMountDetails{}

	if mountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_target_id")); ok {
		tmp := mountTargetId.(string)
		result.MountTargetId = &tmp
	}

	return result, nil
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToUpdateFileSystemMountDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateFileSystemMountDetails, error) {
	result := oci_disaster_recovery.UpdateFileSystemMountDetails{}

	if mountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_target_id")); ok {
		tmp := mountTargetId.(string)
		result.MountTargetId = &tmp
	}

	return result, nil
}

func FileSystemMountDetailsToMap(obj *oci_disaster_recovery.FileSystemMountDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MountTargetId != nil {
		result["mount_target_id"] = string(*obj.MountTargetId)
	}

	return result
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToCreateFileSystemUnmountDetails(fieldKeyFormat string) (oci_disaster_recovery.CreateFileSystemUnmountDetails, error) {
	result := oci_disaster_recovery.CreateFileSystemUnmountDetails{}

	if mountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_target_id")); ok {
		tmp := mountTargetId.(string)
		result.MountTargetId = &tmp
	}

	return result, nil
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToUpdateFileSystemUnmountDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateFileSystemUnmountDetails, error) {
	result := oci_disaster_recovery.UpdateFileSystemUnmountDetails{}

	if mountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_target_id")); ok {
		tmp := mountTargetId.(string)
		result.MountTargetId = &tmp
	}

	return result, nil
}

func FileSystemUnmountDetailsToMap(obj *oci_disaster_recovery.FileSystemUnmountDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MountTargetId != nil {
		result["mount_target_id"] = string(*obj.MountTargetId)
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

	result["lifecycle_sub_state"] = string(obj.LifecycleSubState)

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

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToFileSystemExportMappingDetails(fieldKeyFormat string) (oci_disaster_recovery.FileSystemExportMappingDetails, error) {
	result := oci_disaster_recovery.FileSystemExportMappingDetails{}

	if destinationMountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_mount_target_id")); ok {
		tmp := destinationMountTargetId.(string)
		result.DestinationMountTargetId = &tmp
	}

	if exportId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_id")); ok {
		tmp := exportId.(string)
		result.ExportId = &tmp
	}

	return result, nil
}

func FileSystemExportMappingDetailsToMap(obj oci_disaster_recovery.FileSystemExportMapping) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationMountTargetId != nil {
		result["destination_mount_target_id"] = string(*obj.DestinationMountTargetId)
	}

	if obj.ExportId != nil {
		result["export_id"] = string(*obj.ExportId)
	}

	return result
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToLoadBalancerBackendSetMappingDetails(fieldKeyFormat string) (oci_disaster_recovery.LoadBalancerBackendSetMappingDetails, error) {
	result := oci_disaster_recovery.LoadBalancerBackendSetMappingDetails{}

	if destinationBackendSetName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_backend_set_name")); ok {
		tmp := destinationBackendSetName.(string)
		result.DestinationBackendSetName = &tmp
	}

	if isBackendSetForNonMovable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_backend_set_for_non_movable")); ok {
		tmp := isBackendSetForNonMovable.(bool)
		result.IsBackendSetForNonMovable = &tmp
	}

	if sourceBackendSetName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_backend_set_name")); ok {
		tmp := sourceBackendSetName.(string)
		result.SourceBackendSetName = &tmp
	}

	return result, nil
}

func LoadBalancerBackendSetMappingDetailsToMap(obj oci_disaster_recovery.LoadBalancerBackendSetMapping) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationBackendSetName != nil {
		result["destination_backend_set_name"] = string(*obj.DestinationBackendSetName)
	}

	if obj.IsBackendSetForNonMovable != nil {
		result["is_backend_set_for_non_movable"] = bool(*obj.IsBackendSetForNonMovable)
	}

	if obj.SourceBackendSetName != nil {
		result["source_backend_set_name"] = string(*obj.SourceBackendSetName)
	}

	return result
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToNetworkLoadBalancerBackendSetMappingDetails(fieldKeyFormat string) (oci_disaster_recovery.NetworkLoadBalancerBackendSetMappingDetails, error) {
	result := oci_disaster_recovery.NetworkLoadBalancerBackendSetMappingDetails{}

	if destinationBackendSetName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_backend_set_name")); ok {
		tmp := destinationBackendSetName.(string)
		result.DestinationBackendSetName = &tmp
	}

	if isBackendSetForNonMovable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_backend_set_for_non_movable")); ok {
		tmp := isBackendSetForNonMovable.(bool)
		result.IsBackendSetForNonMovable = &tmp
	}

	if sourceBackendSetName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_backend_set_name")); ok {
		tmp := sourceBackendSetName.(string)
		result.SourceBackendSetName = &tmp
	}

	return result, nil
}

func NetworkLoadBalancerBackendSetMappingDetailsToMap(obj oci_disaster_recovery.NetworkLoadBalancerBackendSetMapping) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationBackendSetName != nil {
		result["destination_backend_set_name"] = string(*obj.DestinationBackendSetName)
	}

	if obj.IsBackendSetForNonMovable != nil {
		result["is_backend_set_for_non_movable"] = bool(*obj.IsBackendSetForNonMovable)
	}

	if obj.SourceBackendSetName != nil {
		result["source_backend_set_name"] = string(*obj.SourceBackendSetName)
	}

	return result
}

func UpdateBlockVolumeAttachmentDetailsToMap(obj *oci_disaster_recovery.UpdateBlockVolumeAttachmentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.VolumeAttachmentReferenceInstanceId != nil {
		result["volume_attachment_reference_instance_id"] = string(*obj.VolumeAttachmentReferenceInstanceId)
	}

	return result
}

func UpdateBlockVolumeMountDetailsToMap(obj *oci_disaster_recovery.UpdateBlockVolumeMountDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MountPoint != nil {
		result["mount_point"] = string(*obj.MountPoint)
	}

	return result
}

func UpdateComputeInstanceMovableFileSystemOperationDetailsToMap(obj oci_disaster_recovery.UpdateComputeInstanceMovableFileSystemOperationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExportPath != nil {
		result["export_path"] = string(*obj.ExportPath)
	}

	if obj.MountDetails != nil {
		result["mount_details"] = []interface{}{UpdateFileSystemMountDetailsToMap(obj.MountDetails)}
	}

	if obj.MountPoint != nil {
		result["mount_point"] = string(*obj.MountPoint)
	}

	if obj.UnmountDetails != nil {
		result["unmount_details"] = []interface{}{UpdateFileSystemUnmountDetailsToMap(obj.UnmountDetails)}
	}

	return result
}

func UpdateComputeInstanceNonMovableBlockVolumeOperationDetailsToMap(obj oci_disaster_recovery.UpdateComputeInstanceNonMovableBlockVolumeOperationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AttachmentDetails != nil {
		result["attachment_details"] = []interface{}{UpdateBlockVolumeAttachmentDetailsToMap(obj.AttachmentDetails)}
	}

	if obj.BlockVolumeId != nil {
		result["block_volume_id"] = string(*obj.BlockVolumeId)
	}

	if obj.MountDetails != nil {
		result["mount_details"] = []interface{}{UpdateBlockVolumeMountDetailsToMap(obj.MountDetails)}
	}

	return result
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToCreateComputeInstanceNonMovableFileSystemOperationDetails(fieldKeyFormat string) (oci_disaster_recovery.CreateComputeInstanceNonMovableFileSystemOperationDetails, error) {
	result := oci_disaster_recovery.CreateComputeInstanceNonMovableFileSystemOperationDetails{}

	if exportPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_path")); ok {
		tmp := exportPath.(string)
		result.ExportPath = &tmp
	}

	if mountPoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_point")); ok {
		tmp := mountPoint.(string)
		result.MountPoint = &tmp
	}

	if mountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_target_id")); ok {
		tmp := mountTargetId.(string)
		result.MountTargetId = &tmp
	}

	return result, nil
}

func (s *DisasterRecoveryDrProtectionGroupResourceCrud) mapToUpdateComputeInstanceNonMovableFileSystemOperationDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateComputeInstanceNonMovableFileSystemOperationDetails, error) {
	result := oci_disaster_recovery.UpdateComputeInstanceNonMovableFileSystemOperationDetails{}

	if exportPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_path")); ok {
		tmp := exportPath.(string)
		result.ExportPath = &tmp
	}

	if mountPoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_point")); ok {
		tmp := mountPoint.(string)
		result.MountPoint = &tmp
	}

	if mountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_target_id")); ok {
		tmp := mountTargetId.(string)
		result.MountTargetId = &tmp
	}

	return result, nil
}

func ComputeInstanceNonMovableFileSystemOperationToMap(obj oci_disaster_recovery.ComputeInstanceNonMovableFileSystemOperation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExportPath != nil {
		result["export_path"] = string(*obj.ExportPath)
	}

	if obj.MountPoint != nil {
		result["mount_point"] = string(*obj.MountPoint)
	}

	if obj.MountTargetId != nil {
		result["mount_target_id"] = string(*obj.MountTargetId)
	}

	return result
}

func UpdateFileSystemMountDetailsToMap(obj *oci_disaster_recovery.UpdateFileSystemMountDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MountTargetId != nil {
		result["mount_target_id"] = string(*obj.MountTargetId)
	}

	return result
}

func UpdateFileSystemUnmountDetailsToMap(obj *oci_disaster_recovery.UpdateFileSystemUnmountDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MountTargetId != nil {
		result["mount_target_id"] = string(*obj.MountTargetId)
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
