// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package desktops

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_desktops "github.com/oracle/oci-go-sdk/v65/desktops"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DesktopsDesktopPoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDesktopsDesktopPool,
		Read:     readDesktopsDesktopPool,
		Update:   updateDesktopsDesktopPool,
		Delete:   deleteDesktopsDesktopPool,
		Schema: map[string]*schema.Schema{
			// Required
			"are_privileged_users": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"availability_policy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"start_schedule": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"cron_expression": {
										Type:     schema.TypeString,
										Required: true,
									},
									"timezone": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"stop_schedule": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"cron_expression": {
										Type:     schema.TypeString,
										Required: true,
									},
									"timezone": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"contact_details": {
				Type:     schema.TypeString,
				Required: true,
			},
			"device_policy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"audio_mode": {
							Type:     schema.TypeString,
							Required: true,
						},
						"cdm_mode": {
							Type:     schema.TypeString,
							Required: true,
						},
						"clipboard_mode": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_display_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"is_keyboard_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"is_pointer_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"is_printing_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"image_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"image_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"operating_system": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"is_storage_enabled": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"maximum_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"network_configuration": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"vcn_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"shape_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"standby_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"storage_backup_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"storage_size_in_gbs": {
				Type:     schema.TypeInt,
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
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shape_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"baseline_ocpu_utilization": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"memory_in_gbs": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"ocpus": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"private_access_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							ForceNew: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"private_ip": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},

						// Computed
						"endpoint_fqdn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vcn_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"session_lifecycle_actions": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"disconnect": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"action": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"grace_period_in_minutes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"inactivity": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"action": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"grace_period_in_minutes": {
										Type:     schema.TypeInt,
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
			"time_start_scheduled": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"time_stop_scheduled": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"use_dedicated_vm_host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"active_desktops": {
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
		},
	}
}

func createDesktopsDesktopPool(d *schema.ResourceData, m interface{}) error {
	sync := &DesktopsDesktopPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DesktopServiceClient()

	return tfresource.CreateResource(d, sync)
}

func readDesktopsDesktopPool(d *schema.ResourceData, m interface{}) error {
	sync := &DesktopsDesktopPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DesktopServiceClient()

	return tfresource.ReadResource(sync)
}

func updateDesktopsDesktopPool(d *schema.ResourceData, m interface{}) error {
	sync := &DesktopsDesktopPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DesktopServiceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDesktopsDesktopPool(d *schema.ResourceData, m interface{}) error {
	sync := &DesktopsDesktopPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DesktopServiceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DesktopsDesktopPoolResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_desktops.DesktopServiceClient
	Res                    *oci_desktops.DesktopPool
	DisableNotFoundRetries bool
}

func (s *DesktopsDesktopPoolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DesktopsDesktopPoolResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_desktops.LifecycleStateCreating),
	}
}

func (s *DesktopsDesktopPoolResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_desktops.LifecycleStateActive),
	}
}

func (s *DesktopsDesktopPoolResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_desktops.LifecycleStateDeleting),
	}
}

func (s *DesktopsDesktopPoolResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_desktops.LifecycleStateDeleted),
	}
}

func (s *DesktopsDesktopPoolResourceCrud) Create() error {
	request := oci_desktops.CreateDesktopPoolRequest{}

	if arePrivilegedUsers, ok := s.D.GetOkExists("are_privileged_users"); ok {
		tmp := arePrivilegedUsers.(bool)
		request.ArePrivilegedUsers = &tmp
	}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if availabilityPolicy, ok := s.D.GetOkExists("availability_policy"); ok {
		if tmpList := availabilityPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "availability_policy", 0)
			tmp, err := s.mapToDesktopAvailabilityPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AvailabilityPolicy = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if contactDetails, ok := s.D.GetOkExists("contact_details"); ok {
		tmp := contactDetails.(string)
		request.ContactDetails = &tmp
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

	if devicePolicy, ok := s.D.GetOkExists("device_policy"); ok {
		if tmpList := devicePolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "device_policy", 0)
			tmp, err := s.mapToDesktopDevicePolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DevicePolicy = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if image, ok := s.D.GetOkExists("image"); ok {
		if tmpList := image.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "image", 0)
			tmp, err := s.mapToDesktopImage(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Image = &tmp
		}
	}

	if isStorageEnabled, ok := s.D.GetOkExists("is_storage_enabled"); ok {
		tmp := isStorageEnabled.(bool)
		request.IsStorageEnabled = &tmp
	}

	if maximumSize, ok := s.D.GetOkExists("maximum_size"); ok {
		tmp := maximumSize.(int)
		request.MaximumSize = &tmp
	}

	if networkConfiguration, ok := s.D.GetOkExists("network_configuration"); ok {
		if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_configuration", 0)
			tmp, err := s.mapToDesktopNetworkConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkConfiguration = &tmp
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

	if shapeConfig, ok := s.D.GetOkExists("shape_config"); ok {
		if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shape_config", 0)
			tmp, err := s.mapToCreateDesktopPoolShapeConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ShapeConfig = &tmp
		}
	}

	if privateAccessDetails, ok := s.D.GetOkExists("private_access_details"); ok {
		if tmpList := privateAccessDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "private_access_details", 0)
			tmp, err := s.mapToCreateDesktopPoolPrivateAccessDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PrivateAccessDetails = &tmp
		}
	}

	if sessionLifecycleActions, ok := s.D.GetOkExists("session_lifecycle_actions"); ok {
		if tmpList := sessionLifecycleActions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "session_lifecycle_actions", 0)
			tmp, err := s.mapToCreateDesktopPoolDesktopSessionLifecycleActions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SessionLifecycleActions = &tmp
		}
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	if standbySize, ok := s.D.GetOkExists("standby_size"); ok {
		tmp := standbySize.(int)
		request.StandbySize = &tmp
	}

	if storageBackupPolicyId, ok := s.D.GetOkExists("storage_backup_policy_id"); ok {
		tmp := storageBackupPolicyId.(string)
		request.StorageBackupPolicyId = &tmp
	}

	if storageSizeInGBs, ok := s.D.GetOkExists("storage_size_in_gbs"); ok {
		tmp := storageSizeInGBs.(int)
		request.StorageSizeInGBs = &tmp
	}

	if timeStartScheduled, ok := s.D.GetOkExists("time_start_scheduled"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStartScheduled.(string))
		if err != nil {
			return err
		}
		request.TimeStartScheduled = &oci_common.SDKTime{Time: tmp}
	}

	if timeStopScheduled, ok := s.D.GetOkExists("time_stop_scheduled"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStopScheduled.(string))
		if err != nil {
			return err
		}
		request.TimeStopScheduled = &oci_common.SDKTime{Time: tmp}
	}

	if useDedicatedVmHost, ok := s.D.GetOkExists("use_dedicated_vm_host"); ok {
		request.UseDedicatedVmHost = oci_desktops.CreateDesktopPoolDetailsUseDedicatedVmHostEnum(useDedicatedVmHost.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "desktops")

	response, err := s.Client.CreateDesktopPool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDesktopPoolFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "desktops"), oci_desktops.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DesktopsDesktopPoolResourceCrud) getDesktopPoolFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_desktops.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	desktopPoolId, err := desktopPoolWaitForWorkRequest(workId, "desktoppool",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, desktopPoolId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_desktops.CancelWorkRequestRequest{
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
	s.D.SetId(*desktopPoolId)

	return s.Get()
}

func desktopPoolWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "desktops", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_desktops.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func desktopPoolWaitForWorkRequest(wId *string, entityType string, action oci_desktops.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_desktops.DesktopServiceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "desktops")
	retryPolicy.ShouldRetryOperation = desktopPoolWorkRequestShouldRetryFunc(timeout)

	response := oci_desktops.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_desktops.OperationStatusInProgress),
			string(oci_desktops.OperationStatusAccepted),
			string(oci_desktops.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_desktops.OperationStatusSucceeded),
			string(oci_desktops.OperationStatusFailed),
			string(oci_desktops.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_desktops.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_desktops.OperationStatusFailed || response.Status == oci_desktops.OperationStatusCanceled {
		return nil, getErrorFromDesktopsDesktopPoolWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDesktopsDesktopPoolWorkRequest(client *oci_desktops.DesktopServiceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_desktops.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_desktops.ListWorkRequestErrorsRequest{
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

func (s *DesktopsDesktopPoolResourceCrud) Get() error {
	request := oci_desktops.GetDesktopPoolRequest{}

	tmp := s.D.Id()
	request.DesktopPoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "desktops")

	response, err := s.Client.GetDesktopPool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DesktopPool
	return nil
}

func (s *DesktopsDesktopPoolResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_desktops.UpdateDesktopPoolRequest{}

	if availabilityPolicy, ok := s.D.GetOkExists("availability_policy"); ok {
		if tmpList := availabilityPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "availability_policy", 0)
			tmp, err := s.mapToDesktopAvailabilityPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AvailabilityPolicy = &tmp
		}
	}

	if contactDetails, ok := s.D.GetOkExists("contact_details"); ok {
		tmp := contactDetails.(string)
		request.ContactDetails = &tmp
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

	tmp := s.D.Id()
	request.DesktopPoolId = &tmp

	if devicePolicy, ok := s.D.GetOkExists("device_policy"); ok {
		if tmpList := devicePolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "device_policy", 0)
			tmp, err := s.mapToDesktopDevicePolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DevicePolicy = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if maximumSize, ok := s.D.GetOkExists("maximum_size"); ok {
		tmp := maximumSize.(int)
		request.MaximumSize = &tmp
	}

	if standbySize, ok := s.D.GetOkExists("standby_size"); ok {
		tmp := standbySize.(int)
		request.StandbySize = &tmp
	}

	if timeStartScheduled, ok := s.D.GetOkExists("time_start_scheduled"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStartScheduled.(string))
		if err != nil {
			return err
		}
		request.TimeStartScheduled = &oci_common.SDKTime{Time: tmp}
	}

	if timeStopScheduled, ok := s.D.GetOkExists("time_stop_scheduled"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStopScheduled.(string))
		if err != nil {
			return err
		}
		request.TimeStopScheduled = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "desktops")

	response, err := s.Client.UpdateDesktopPool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDesktopPoolFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "desktops"), oci_desktops.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DesktopsDesktopPoolResourceCrud) Delete() error {
	request := oci_desktops.DeleteDesktopPoolRequest{}

	if areVolumesPreserved, ok := s.D.GetOkExists("are_volumes_preserved"); ok {
		tmp := areVolumesPreserved.(bool)
		request.AreVolumesPreserved = &tmp
	}

	tmp := s.D.Id()
	request.DesktopPoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "desktops")

	response, err := s.Client.DeleteDesktopPool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := desktopPoolWaitForWorkRequest(workId, "desktoppool",
		oci_desktops.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DesktopsDesktopPoolResourceCrud) SetData() error {

	if s.Res.ArePrivilegedUsers != nil {
		s.D.Set("are_privileged_users", *s.Res.ArePrivilegedUsers)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.AvailabilityPolicy != nil {
		s.D.Set("availability_policy", []interface{}{DesktopAvailabilityPolicyToMap(s.Res.AvailabilityPolicy)})
	} else {
		s.D.Set("availability_policy", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ContactDetails != nil {
		s.D.Set("contact_details", *s.Res.ContactDetails)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DevicePolicy != nil {
		s.D.Set("device_policy", []interface{}{DesktopDevicePolicyToMap(s.Res.DevicePolicy)})
	} else {
		s.D.Set("device_policy", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Image != nil {
		s.D.Set("image", []interface{}{DesktopImageToMap(s.Res.Image)})
	} else {
		s.D.Set("image", nil)
	}

	if s.Res.IsStorageEnabled != nil {
		s.D.Set("is_storage_enabled", *s.Res.IsStorageEnabled)
	}

	if s.Res.MaximumSize != nil {
		s.D.Set("maximum_size", *s.Res.MaximumSize)
	}

	if s.Res.NetworkConfiguration != nil {
		s.D.Set("network_configuration", []interface{}{DesktopNetworkConfigurationToMap(s.Res.NetworkConfiguration)})
	} else {
		s.D.Set("network_configuration", nil)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.ShapeConfig != nil {
		s.D.Set("shape_config", []interface{}{DesktopPoolShapeConfigToMap(s.Res.ShapeConfig)})
	} else {
		s.D.Set("shape_config", nil)
	}

	if s.Res.PrivateAccessDetails != nil {
		s.D.Set("private_access_details", []interface{}{DesktopPoolPrivateAccessDetailsToMap(s.Res.PrivateAccessDetails, false)})
	} else {
		s.D.Set("private_access_details", nil)
	}

	if s.Res.SessionLifecycleActions != nil {
		s.D.Set("session_lifecycle_actions", []interface{}{DesktopSessionLifecycleActionsToMap(s.Res.SessionLifecycleActions)})
	} else {
		s.D.Set("session_lifecycle_actions", nil)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	if s.Res.StandbySize != nil {
		s.D.Set("standby_size", *s.Res.StandbySize)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageBackupPolicyId != nil {
		s.D.Set("storage_backup_policy_id", *s.Res.StorageBackupPolicyId)
	}

	if s.Res.StorageSizeInGBs != nil {
		s.D.Set("storage_size_in_gbs", *s.Res.StorageSizeInGBs)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeStartScheduled != nil {
		s.D.Set("time_start_scheduled", s.Res.TimeStartScheduled.Format(time.RFC3339Nano))
	}

	if s.Res.TimeStopScheduled != nil {
		s.D.Set("time_stop_scheduled", s.Res.TimeStopScheduled.Format(time.RFC3339Nano))
	}

	s.D.Set("use_dedicated_vm_host", s.Res.UseDedicatedVmHost)

	return nil
}

func (s *DesktopsDesktopPoolResourceCrud) mapToCreateDesktopPoolShapeConfigDetails(fieldKeyFormat string) (oci_desktops.CreateDesktopPoolShapeConfigDetails, error) {
	result := oci_desktops.CreateDesktopPoolShapeConfigDetails{}

	if baselineOcpuUtilization, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "baseline_ocpu_utilization")); ok {
		result.BaselineOcpuUtilization = oci_desktops.CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum(baselineOcpuUtilization.(string))
	}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := memoryInGBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert memoryInGBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.MemoryInGBs = &tmpInt64
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := ocpus.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert ocpus string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.Ocpus = &tmpInt64
	}

	return result, nil
}

func DesktopPoolShapeConfigToMap(obj *oci_desktops.DesktopPoolShapeConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["baseline_ocpu_utilization"] = string(obj.BaselineOcpuUtilization)

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = strconv.FormatInt(*obj.MemoryInGBs, 10)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = strconv.FormatInt(*obj.Ocpus, 10)
	}

	return result
}

func (s *DesktopsDesktopPoolResourceCrud) mapToCreateDesktopPoolPrivateAccessDetails(fieldKeyFormat string) (oci_desktops.CreateDesktopPoolPrivateAccessDetails, error) {
	result := oci_desktops.CreateDesktopPoolPrivateAccessDetails{}

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

	if privateIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_ip")); ok {
		tmp := privateIp.(string)
		result.PrivateIp = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func (s *DesktopsDesktopPoolResourceCrud) mapToCreateDesktopPoolDesktopSessionLifecycleActions(fieldKeyFormat string) (oci_desktops.CreateDesktopPoolDesktopSessionLifecycleActions, error) {
	result := oci_desktops.CreateDesktopPoolDesktopSessionLifecycleActions{}

	if disconnect, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "disconnect")); ok {
		if tmpList := disconnect.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "disconnect"), 0)
			tmp, err := s.mapToDisconnectConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert disconnect, encountered error: %v", err)
			}
			result.Disconnect = &tmp
		}
	}

	if inactivity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "inactivity")); ok {
		if tmpList := inactivity.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "inactivity"), 0)
			tmp, err := s.mapToInactivityConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert inactivity, encountered error: %v", err)
			}
			result.Inactivity = &tmp
		}
	}

	return result, nil
}

func DesktopPoolPrivateAccessDetailsToMap(obj *oci_desktops.DesktopPoolPrivateAccessDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EndpointFqdn != nil {
		result["endpoint_fqdn"] = string(*obj.EndpointFqdn)
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	return result
}

func CreateDesktopPoolDesktopSessionLifecycleActionsToMap(obj *oci_desktops.CreateDesktopPoolDesktopSessionLifecycleActions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Disconnect != nil {
		result["disconnect"] = []interface{}{DisconnectConfigToMap(obj.Disconnect)}
	}

	if obj.Inactivity != nil {
		result["inactivity"] = []interface{}{InactivityConfigToMap(obj.Inactivity)}
	}

	return result
}

func DesktopSessionLifecycleActionsToMap(obj *oci_desktops.DesktopSessionLifecycleActions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Disconnect != nil {
		result["disconnect"] = []interface{}{DisconnectConfigToMap(obj.Disconnect)}
	}

	if obj.Inactivity != nil {
		result["inactivity"] = []interface{}{InactivityConfigToMap(obj.Inactivity)}
	}

	return result
}

func (s *DesktopsDesktopPoolResourceCrud) mapToDesktopAvailabilityPolicy(fieldKeyFormat string) (oci_desktops.DesktopAvailabilityPolicy, error) {
	result := oci_desktops.DesktopAvailabilityPolicy{}

	if startSchedule, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_schedule")); ok {
		if tmpList := startSchedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "start_schedule"), 0)
			tmp, err := s.mapToDesktopSchedule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert start_schedule, encountered error: %v", err)
			}
			result.StartSchedule = &tmp
		}
	}

	if stopSchedule, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stop_schedule")); ok {
		if tmpList := stopSchedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "stop_schedule"), 0)
			tmp, err := s.mapToDesktopSchedule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert stop_schedule, encountered error: %v", err)
			}
			result.StopSchedule = &tmp
		}
	}

	return result, nil
}

func DesktopAvailabilityPolicyToMap(obj *oci_desktops.DesktopAvailabilityPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.StartSchedule != nil && obj.StartSchedule.CronExpression != nil {
		result["start_schedule"] = []interface{}{DesktopScheduleToMap(obj.StartSchedule)}
	}

	if obj.StopSchedule != nil && obj.StopSchedule.CronExpression != nil {
		result["stop_schedule"] = []interface{}{DesktopScheduleToMap(obj.StopSchedule)}
	}

	return result
}

func (s *DesktopsDesktopPoolResourceCrud) mapToDesktopDevicePolicy(fieldKeyFormat string) (oci_desktops.DesktopDevicePolicy, error) {
	result := oci_desktops.DesktopDevicePolicy{}

	if audioMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "audio_mode")); ok {
		result.AudioMode = oci_desktops.DesktopDevicePolicyAudioModeEnum(audioMode.(string))
	}

	if cdmMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cdm_mode")); ok {
		result.CdmMode = oci_desktops.DesktopDevicePolicyCdmModeEnum(cdmMode.(string))
	}

	if clipboardMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "clipboard_mode")); ok {
		result.ClipboardMode = oci_desktops.DesktopDevicePolicyClipboardModeEnum(clipboardMode.(string))
	}

	if isDisplayEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_display_enabled")); ok {
		tmp := isDisplayEnabled.(bool)
		result.IsDisplayEnabled = &tmp
	}

	if isKeyboardEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keyboard_enabled")); ok {
		tmp := isKeyboardEnabled.(bool)
		result.IsKeyboardEnabled = &tmp
	}

	if isPointerEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pointer_enabled")); ok {
		tmp := isPointerEnabled.(bool)
		result.IsPointerEnabled = &tmp
	}

	if isPrintingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_printing_enabled")); ok {
		tmp := isPrintingEnabled.(bool)
		result.IsPrintingEnabled = &tmp
	}

	return result, nil
}

func DesktopDevicePolicyToMap(obj *oci_desktops.DesktopDevicePolicy) map[string]interface{} {
	result := map[string]interface{}{}

	result["audio_mode"] = string(obj.AudioMode)

	result["cdm_mode"] = string(obj.CdmMode)

	result["clipboard_mode"] = string(obj.ClipboardMode)

	if obj.IsDisplayEnabled != nil {
		result["is_display_enabled"] = bool(*obj.IsDisplayEnabled)
	}

	if obj.IsKeyboardEnabled != nil {
		result["is_keyboard_enabled"] = bool(*obj.IsKeyboardEnabled)
	}

	if obj.IsPointerEnabled != nil {
		result["is_pointer_enabled"] = bool(*obj.IsPointerEnabled)
	}

	if obj.IsPrintingEnabled != nil {
		result["is_printing_enabled"] = bool(*obj.IsPrintingEnabled)
	}

	return result
}

func (s *DesktopsDesktopPoolResourceCrud) mapToDesktopImage(fieldKeyFormat string) (oci_desktops.DesktopImage, error) {
	result := oci_desktops.DesktopImage{}

	if imageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_id")); ok {
		tmp := imageId.(string)
		result.ImageId = &tmp
	}

	if imageName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_name")); ok {
		tmp := imageName.(string)
		result.ImageName = &tmp
	}

	if operatingSystem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operating_system")); ok {
		tmp := operatingSystem.(string)
		result.OperatingSystem = &tmp
	}

	return result, nil
}

func DesktopImageToMap(obj *oci_desktops.DesktopImage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ImageId != nil {
		result["image_id"] = string(*obj.ImageId)
	}

	if obj.ImageName != nil {
		result["image_name"] = string(*obj.ImageName)
	}

	if obj.OperatingSystem != nil {
		result["operating_system"] = string(*obj.OperatingSystem)
	}

	return result
}

func (s *DesktopsDesktopPoolResourceCrud) mapToDesktopNetworkConfiguration(fieldKeyFormat string) (oci_desktops.DesktopNetworkConfiguration, error) {
	result := oci_desktops.DesktopNetworkConfiguration{}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcn_id")); ok {
		tmp := vcnId.(string)
		result.VcnId = &tmp
	}

	return result, nil
}

func DesktopNetworkConfigurationToMap(obj *oci_desktops.DesktopNetworkConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	return result
}

func DesktopPoolSummaryToMap(obj oci_desktops.DesktopPoolSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActiveDesktops != nil {
		result["active_desktops"] = int(*obj.ActiveDesktops)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ContactDetails != nil {
		result["contact_details"] = string(*obj.ContactDetails)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MaximumSize != nil {
		result["maximum_size"] = int(*obj.MaximumSize)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *DesktopsDesktopPoolResourceCrud) mapToDesktopSchedule(fieldKeyFormat string) (oci_desktops.DesktopSchedule, error) {
	result := oci_desktops.DesktopSchedule{}

	if cronExpression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cron_expression")); ok {
		tmp := cronExpression.(string)
		result.CronExpression = &tmp
	}

	if timezone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timezone")); ok {
		tmp := timezone.(string)
		result.Timezone = &tmp
	}

	return result, nil
}

func DesktopScheduleToMap(obj *oci_desktops.DesktopSchedule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CronExpression != nil {
		result["cron_expression"] = string(*obj.CronExpression)
	}

	if obj.Timezone != nil {
		result["timezone"] = string(*obj.Timezone)
	}

	return result
}

func (s *DesktopsDesktopPoolResourceCrud) mapToDisconnectConfig(fieldKeyFormat string) (oci_desktops.DisconnectConfig, error) {
	result := oci_desktops.DisconnectConfig{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_desktops.DisconnectConfigActionEnum(action.(string))
	}

	if gracePeriodInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grace_period_in_minutes")); ok {
		tmp := gracePeriodInMinutes.(int)
		result.GracePeriodInMinutes = &tmp
	}

	return result, nil
}

func DisconnectConfigToMap(obj *oci_desktops.DisconnectConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.GracePeriodInMinutes != nil {
		result["grace_period_in_minutes"] = int(*obj.GracePeriodInMinutes)
	}

	return result
}

func (s *DesktopsDesktopPoolResourceCrud) mapToInactivityConfig(fieldKeyFormat string) (oci_desktops.InactivityConfig, error) {
	result := oci_desktops.InactivityConfig{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_desktops.InactivityConfigActionEnum(action.(string))
	}

	if gracePeriodInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grace_period_in_minutes")); ok {
		tmp := gracePeriodInMinutes.(int)
		result.GracePeriodInMinutes = &tmp
	}

	return result, nil
}

func InactivityConfigToMap(obj *oci_desktops.InactivityConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.GracePeriodInMinutes != nil {
		result["grace_period_in_minutes"] = int(*obj.GracePeriodInMinutes)
	}

	return result
}

func (s *DesktopsDesktopPoolResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_desktops.ChangeDesktopPoolCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DesktopPoolId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "desktops")

	response, err := s.Client.ChangeDesktopPoolCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDesktopPoolFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "desktops"), oci_desktops.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
