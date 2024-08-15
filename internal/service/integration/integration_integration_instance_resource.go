// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integration

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_integration "github.com/oracle/oci-go-sdk/v65/integration"
)

func IntegrationIntegrationInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createIntegrationIntegrationInstance,
		Read:   readIntegrationIntegrationInstance,
		Update: updateIntegrationIntegrationInstance,
		Delete: deleteIntegrationIntegrationInstance,
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
			"integration_instance_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_byol": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"message_packs": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Optional
			"alternate_custom_endpoints": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      alternateCustomEndpointsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"hostname": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"certificate_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"alias": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate_secret_version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"consumption_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"custom_endpoint": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"hostname": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"certificate_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"alias": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate_secret_version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
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
			"domain_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"idcs_at": {
				Type:      schema.TypeString,
				Optional:  true,
				StateFunc: tfresource.GetMd5Hash,
				Sensitive: true,
			},
			"is_file_server_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_visual_builder_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"network_endpoint_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"network_endpoint_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"PUBLIC",
							}, true),
						},

						// Optional
						"allowlisted_http_ips": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"allowlisted_http_vcns": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Set:      allowlistedHttpVcnsHashCodeForSets,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"allowlisted_ips": {
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
						"is_integration_vcn_allowlisted": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"enable_process_automation_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"extend_data_retention_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_implicit": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_instance_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_service_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"data_retention_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"idcs_app_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_app_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_app_location_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_app_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_primary_audience_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"instance_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_endpoint_outbound_connection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"nsg_ids": {
							Type:     schema.TypeSet,
							Computed: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"outbound_connection_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_integration.IntegrationInstanceLifecycleStateActive),
					string(oci_integration.IntegrationInstanceLifecycleStateInactive),
				}, true),
			},
			"state_message": {
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

func createIntegrationIntegrationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationIntegrationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IntegrationInstanceClient()

	if _, ok := sync.D.GetOkExists("enable_process_automation_trigger"); ok {
		err := sync.EnableProcessAutomation()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("extend_data_retention_trigger"); ok {
		err := sync.ExtendDataRetention()
		if err != nil {
			return err
		}
	}
	var powerOff = false
	if configState, ok := sync.D.GetOkExists("state"); ok {
		wantedState := oci_integration.IntegrationInstanceLifecycleStateEnum(strings.ToUpper(configState.(string)))
		if wantedState == oci_integration.IntegrationInstanceLifecycleStateInactive {
			powerOff = true
		}
	}

	if error := tfresource.CreateResource(d, sync); error != nil {
		return error
	}

	if powerOff {
		return powerOffIntegrationInstance(d, sync)
	}

	return nil
}

func powerOffIntegrationInstance(d *schema.ResourceData, sync *IntegrationIntegrationInstanceResourceCrud) error {
	if err := sync.StopIntegerationInstance(); err != nil {
		return err
	}
	return tfresource.ReadResource(sync)
}

func readIntegrationIntegrationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationIntegrationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IntegrationInstanceClient()

	return tfresource.ReadResource(sync)
}

func updateIntegrationIntegrationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationIntegrationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IntegrationInstanceClient()

	if _, ok := sync.D.GetOkExists("enable_process_automation_trigger"); ok && sync.D.HasChange("enable_process_automation_trigger") {
		oldRaw, newRaw := sync.D.GetChange("enable_process_automation_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.EnableProcessAutomation()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("enable_process_automation_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}
	// Start/Stop Integration instance
	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_integration.IntegrationInstanceLifecycleStateActive == oci_integration.IntegrationInstanceLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_integration.IntegrationInstanceLifecycleStateInactive == oci_integration.IntegrationInstanceLifecycleStateEnum(wantedState) {
			powerOff = true
		} else {
			return fmt.Errorf("[ERROR] Invalid state input for Update %v", wantedState)
		}
	}

	if powerOn {
		if err := sync.StartIntegerationInstance(); err != nil {
			return err
		}
		if err := sync.D.Set("state", oci_integration.IntegrationInstanceLifecycleStateActive); err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("extend_data_retention_trigger"); ok && sync.D.HasChange("extend_data_retention_trigger") {
		oldRaw, newRaw := sync.D.GetChange("extend_data_retention_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ExtendDataRetention()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("extend_data_retention_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopIntegerationInstance(); err != nil {
			return err
		}
		if err := sync.D.Set("state", oci_integration.IntegrationInstanceLifecycleStateInactive); err != nil {
			return err
		}
	}

	return nil
}

func deleteIntegrationIntegrationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationIntegrationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IntegrationInstanceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IntegrationIntegrationInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_integration.IntegrationInstanceClient
	Res                    *oci_integration.IntegrationInstance
	DisableNotFoundRetries bool
}

func (s *IntegrationIntegrationInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IntegrationIntegrationInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateCreating),
	}
}

func (s *IntegrationIntegrationInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateActive),
	}
}

func (s *IntegrationIntegrationInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateDeleting),
	}
}

func (s *IntegrationIntegrationInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_integration.IntegrationInstanceLifecycleStateDeleted),
	}
}

func (s *IntegrationIntegrationInstanceResourceCrud) Create() error {
	request := oci_integration.CreateIntegrationInstanceRequest{}

	if alternateCustomEndpoints, ok := s.D.GetOkExists("alternate_custom_endpoints"); ok {
		set := alternateCustomEndpoints.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_integration.CreateCustomEndpointDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := alternateCustomEndpointsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "alternate_custom_endpoints", stateDataIndex)
			converted, err := s.mapToCreateCustomEndpointDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("alternate_custom_endpoints") {
			request.AlternateCustomEndpoints = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if consumptionModel, ok := s.D.GetOkExists("consumption_model"); ok {
		request.ConsumptionModel = oci_integration.CreateIntegrationInstanceDetailsConsumptionModelEnum(consumptionModel.(string))
	}

	if customEndpoint, ok := s.D.GetOkExists("custom_endpoint"); ok {
		if tmpList := customEndpoint.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_endpoint", 0)
			tmp, err := s.mapToCreateCustomEndpointDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CustomEndpoint = &tmp
		}
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

	if domainId, ok := s.D.GetOkExists("domain_id"); ok {
		tmp := domainId.(string)
		request.DomainId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if idcsAt, ok := s.D.GetOkExists("idcs_at"); ok {
		tmp := idcsAt.(string)
		request.IdcsAt = &tmp
	}

	if integrationInstanceType, ok := s.D.GetOkExists("integration_instance_type"); ok {
		request.IntegrationInstanceType = oci_integration.CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum(integrationInstanceType.(string))
	}

	if isByol, ok := s.D.GetOkExists("is_byol"); ok {
		tmp := isByol.(bool)
		request.IsByol = &tmp
	}

	if isFileServerEnabled, ok := s.D.GetOkExists("is_file_server_enabled"); ok {
		tmp := isFileServerEnabled.(bool)
		request.IsFileServerEnabled = &tmp
	}

	if isVisualBuilderEnabled, ok := s.D.GetOkExists("is_visual_builder_enabled"); ok {
		tmp := isVisualBuilderEnabled.(bool)
		request.IsVisualBuilderEnabled = &tmp
	}

	if messagePacks, ok := s.D.GetOkExists("message_packs"); ok {
		tmp := messagePacks.(int)
		request.MessagePacks = &tmp
	}

	if networkEndpointDetails, ok := s.D.GetOkExists("network_endpoint_details"); ok {
		if tmpList := networkEndpointDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_endpoint_details", 0)
			tmp, err := s.mapToNetworkEndpointDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkEndpointDetails = tmp
		}
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		request.Shape = oci_integration.CreateIntegrationInstanceDetailsShapeEnum(shape.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration")

	response, err := s.Client.CreateIntegrationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_integration.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_integration.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "integration") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getIntegrationInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration"), oci_integration.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IntegrationIntegrationInstanceResourceCrud) getIntegrationInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_integration.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	integrationInstanceId, err := integrationInstanceWaitForWorkRequest(workId, "integration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*integrationInstanceId)

	return s.Get()
}

func integrationInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "integration", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_integration.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func integrationInstanceWaitForWorkRequest(wId *string, entityType string, action oci_integration.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_integration.IntegrationInstanceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "integration")
	retryPolicy.ShouldRetryOperation = integrationInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_integration.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_integration.WorkRequestStatusInProgress),
			string(oci_integration.WorkRequestStatusAccepted),
			string(oci_integration.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_integration.WorkRequestStatusSucceeded),
			string(oci_integration.WorkRequestStatusFailed),
			string(oci_integration.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_integration.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_integration.WorkRequestStatusFailed || response.Status == oci_integration.WorkRequestStatusCanceled {
		return nil, getErrorFromIntegrationIntegrationInstanceWorkRequest(client, wId, response.CompartmentId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromIntegrationIntegrationInstanceWorkRequest(client *oci_integration.IntegrationInstanceClient, compartmentId *string, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_integration.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_integration.ListWorkRequestErrorsRequest{
			CompartmentId: compartmentId,
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

func (s *IntegrationIntegrationInstanceResourceCrud) Get() error {
	request := oci_integration.GetIntegrationInstanceRequest{}

	tmp := s.D.Id()
	request.IntegrationInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration")

	response, err := s.Client.GetIntegrationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IntegrationInstance
	return nil
}

func (s *IntegrationIntegrationInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_integration.UpdateIntegrationInstanceRequest{}

	if alternateCustomEndpoints, ok := s.D.GetOkExists("alternate_custom_endpoints"); ok {
		set := alternateCustomEndpoints.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_integration.UpdateCustomEndpointDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := alternateCustomEndpointsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "alternate_custom_endpoints", stateDataIndex)
			converted, err := s.mapToUpdateCustomEndpointDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("alternate_custom_endpoints") {
			request.AlternateCustomEndpoints = tmp
		}
	}

	if customEndpoint, ok := s.D.GetOkExists("custom_endpoint"); ok {
		if tmpList := customEndpoint.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_endpoint", 0)
			tmp, err := s.mapToUpdateCustomEndpointDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CustomEndpoint = &tmp
		}
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

	tmp := s.D.Id()
	request.IntegrationInstanceId = &tmp

	if integrationInstanceType, ok := s.D.GetOkExists("integration_instance_type"); ok {
		request.IntegrationInstanceType = oci_integration.UpdateIntegrationInstanceDetailsIntegrationInstanceTypeEnum(integrationInstanceType.(string))
	}

	if isByol, ok := s.D.GetOkExists("is_byol"); ok {
		tmp := isByol.(bool)
		request.IsByol = &tmp
	}

	if isFileServerEnabled, ok := s.D.GetOkExists("is_file_server_enabled"); ok {
		tmp := isFileServerEnabled.(bool)
		request.IsFileServerEnabled = &tmp
	}

	if isVisualBuilderEnabled, ok := s.D.GetOkExists("is_visual_builder_enabled"); ok {
		tmp := isVisualBuilderEnabled.(bool)
		request.IsVisualBuilderEnabled = &tmp
	}

	if messagePacks, ok := s.D.GetOkExists("message_packs"); ok {
		tmp := messagePacks.(int)
		request.MessagePacks = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration")

	response, err := s.Client.UpdateIntegrationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIntegrationInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration"), oci_integration.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *IntegrationIntegrationInstanceResourceCrud) Delete() error {
	request := oci_integration.DeleteIntegrationInstanceRequest{}

	tmp := s.D.Id()
	request.IntegrationInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration")

	response, err := s.Client.DeleteIntegrationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	// _, delWorkRequestErr := integrationInstanceWaitForWorkRequest(workId, "integration",
	// 	oci_integration.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	// return delWorkRequestErr

	if _, err := integrationInstanceWaitForWorkRequest(workId, "integration", oci_integration.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client); err != nil {
		return err
	}
	return nil

}

func (s *IntegrationIntegrationInstanceResourceCrud) SetData() error {
	alternateCustomEndpoints := []interface{}{}
	for _, item := range s.Res.AlternateCustomEndpoints {
		alternateCustomEndpoints = append(alternateCustomEndpoints, CustomEndpointDetailsToMap(&item))
	}
	s.D.Set("alternate_custom_endpoints", schema.NewSet(alternateCustomEndpointsHashCodeForSets, alternateCustomEndpoints))

	attachments := []interface{}{}
	for _, item := range s.Res.Attachments {
		attachments = append(attachments, AttachmentDetailsToMap(item))
	}
	s.D.Set("attachments", attachments)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("consumption_model", s.Res.ConsumptionModel)

	if s.Res.CustomEndpoint != nil {
		s.D.Set("custom_endpoint", []interface{}{CustomEndpointDetailsToMap(s.Res.CustomEndpoint)})
	} else {
		s.D.Set("custom_endpoint", nil)
	}

	s.D.Set("data_retention_period", s.Res.DataRetentionPeriod)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdcsInfo != nil {
		s.D.Set("idcs_info", []interface{}{IdcsInfoDetailsToMap(s.Res.IdcsInfo)})
	} else {
		s.D.Set("idcs_info", nil)
	}

	if s.Res.InstanceUrl != nil {
		s.D.Set("instance_url", *s.Res.InstanceUrl)
	}

	s.D.Set("integration_instance_type", s.Res.IntegrationInstanceType)

	if s.Res.IsByol != nil {
		s.D.Set("is_byol", *s.Res.IsByol)
	}

	if s.Res.IsFileServerEnabled != nil {
		s.D.Set("is_file_server_enabled", *s.Res.IsFileServerEnabled)
	}

	if s.Res.IsVisualBuilderEnabled != nil {
		s.D.Set("is_visual_builder_enabled", *s.Res.IsVisualBuilderEnabled)
	}

	if s.Res.MessagePacks != nil {
		s.D.Set("message_packs", *s.Res.MessagePacks)
	}

	if s.Res.NetworkEndpointDetails != nil {
		networkEndpointDetailsArray := []interface{}{}
		if networkEndpointDetailsMap := IntegNetworkEndpointDetailsToMap(&s.Res.NetworkEndpointDetails, false); networkEndpointDetailsMap != nil {
			networkEndpointDetailsArray = append(networkEndpointDetailsArray, networkEndpointDetailsMap)
		}
		s.D.Set("network_endpoint_details", networkEndpointDetailsArray)
	} else {
		s.D.Set("network_endpoint_details", nil)
	}

	s.D.Set("shape", s.Res.Shape)

	if s.Res.PrivateEndpointOutboundConnection != nil {
		privateEndpointOutboundConnectionArray := []interface{}{}
		if privateEndpointOutboundConnectionMap := OutboundConnectionToMap(&s.Res.PrivateEndpointOutboundConnection, false); privateEndpointOutboundConnectionMap != nil {
			privateEndpointOutboundConnectionArray = append(privateEndpointOutboundConnectionArray, privateEndpointOutboundConnectionMap)
		}
		s.D.Set("private_endpoint_outbound_connection", privateEndpointOutboundConnectionArray)
	} else {
		s.D.Set("private_endpoint_outbound_connection", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

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

func (s *IntegrationIntegrationInstanceResourceCrud) EnableProcessAutomation() error {
	request := oci_integration.EnableProcessAutomationRequest{}

	idTmp := s.D.Id()
	request.IntegrationInstanceId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration")

	_, err := s.Client.EnableProcessAutomation(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("enable_process_automation_trigger")
	s.D.Set("enable_process_automation_trigger", val)

	return nil
}

func (s *IntegrationIntegrationInstanceResourceCrud) ExtendDataRetention() error {
	request := oci_integration.ExtendDataRetentionRequest{}

	if dataRetentionPeriod, ok := s.D.GetOkExists("data_retention_period"); ok {
		request.DataRetentionPeriod = oci_integration.ExtendDataRetentionDetailsDataRetentionPeriodEnum(dataRetentionPeriod.(string))
	}

	idTmp := s.D.Id()
	request.IntegrationInstanceId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration")

	_, err := s.Client.ExtendDataRetention(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("extend_data_retention_trigger")
	s.D.Set("extend_data_retention_trigger", val)

	return nil
}

func AttachmentDetailsToMap(obj oci_integration.AttachmentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsImplicit != nil {
		result["is_implicit"] = bool(*obj.IsImplicit)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TargetInstanceUrl != nil {
		result["target_instance_url"] = string(*obj.TargetInstanceUrl)
	}

	result["target_role"] = string(obj.TargetRole)

	if obj.TargetServiceType != nil {
		result["target_service_type"] = string(*obj.TargetServiceType)
	}

	return result
}

func (s *IntegrationIntegrationInstanceResourceCrud) mapToCreateCustomEndpointDetails(fieldKeyFormat string) (oci_integration.CreateCustomEndpointDetails, error) {
	result := oci_integration.CreateCustomEndpointDetails{}

	if certificateSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_secret_id")); ok {
		tmp := certificateSecretId.(string)
		result.CertificateSecretId = &tmp
	}

	if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
		tmp := hostname.(string)
		result.Hostname = &tmp
	}

	return result, nil
}

func (s *IntegrationIntegrationInstanceResourceCrud) mapToUpdateCustomEndpointDetails(fieldKeyFormat string) (oci_integration.UpdateCustomEndpointDetails, error) {
	result := oci_integration.UpdateCustomEndpointDetails{}

	if certificateSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_secret_id")); ok && certificateSecretId != "" {
		tmp := certificateSecretId.(string)
		result.CertificateSecretId = &tmp
	}

	if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
		tmp := hostname.(string)
		result.Hostname = &tmp
	}

	return result, nil
}

func CustomEndpointDetailsToMap(obj *oci_integration.CustomEndpointDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Alias != nil {
		result["alias"] = string(*obj.Alias)
	}

	if obj.CertificateSecretId != nil {
		result["certificate_secret_id"] = string(*obj.CertificateSecretId)
	}

	if obj.CertificateSecretVersion != nil {
		result["certificate_secret_version"] = int(*obj.CertificateSecretVersion)
	}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	return result
}

func IdcsInfoDetailsToMap(obj *oci_integration.IdcsInfoDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IdcsAppDisplayName != nil {
		result["idcs_app_display_name"] = string(*obj.IdcsAppDisplayName)
	}

	if obj.IdcsAppId != nil {
		result["idcs_app_id"] = string(*obj.IdcsAppId)
	}

	if obj.IdcsAppLocationUrl != nil {
		result["idcs_app_location_url"] = string(*obj.IdcsAppLocationUrl)
	}

	if obj.IdcsAppName != nil {
		result["idcs_app_name"] = string(*obj.IdcsAppName)
	}

	if obj.InstancePrimaryAudienceUrl != nil {
		result["instance_primary_audience_url"] = string(*obj.InstancePrimaryAudienceUrl)
	}

	return result
}

func (s *IntegrationIntegrationInstanceResourceCrud) mapToNetworkEndpointDetails(fieldKeyFormat string) (oci_integration.NetworkEndpointDetails, error) {
	var baseObject oci_integration.NetworkEndpointDetails
	//discriminator
	networkEndpointTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_endpoint_type"))
	var networkEndpointType string
	if ok {
		networkEndpointType = networkEndpointTypeRaw.(string)
	} else {
		networkEndpointType = "" // default value
	}
	switch strings.ToLower(networkEndpointType) {
	case strings.ToLower("PUBLIC"):
		details := oci_integration.PublicEndpointDetails{}
		if allowlistedHttpIps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowlisted_http_ips")); ok {
			set := allowlistedHttpIps.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowlisted_http_ips")) {
				details.AllowlistedHttpIps = tmp
			}
		}
		if allowlistedHttpVcns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowlisted_http_vcns")); ok {
			set := allowlistedHttpVcns.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_integration.VirtualCloudNetwork, len(interfaces))
			for i := range interfaces {
				stateDataIndex := allowlistedHttpVcnsHashCodeForSets(interfaces[i])
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "allowlisted_http_vcns"), stateDataIndex)
				converted, err := s.mapToVirtualCloudNetwork(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowlisted_http_vcns")) {
				details.AllowlistedHttpVcns = tmp
			}
		}
		if isIntegrationVcnAllowlisted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_integration_vcn_allowlisted")); ok {
			tmp := isIntegrationVcnAllowlisted.(bool)
			details.IsIntegrationVcnAllowlisted = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown network_endpoint_type '%v' was specified", networkEndpointType)
	}
	return baseObject, nil
}

func IntegNetworkEndpointDetailsToMap(obj *oci_integration.NetworkEndpointDetails, datasource bool) map[string]interface{} {

	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_integration.PublicEndpointDetails:
		result["network_endpoint_type"] = "PUBLIC"

		allowlistedHttpIps := []interface{}{}
		for _, item := range v.AllowlistedHttpIps {
			allowlistedHttpIps = append(allowlistedHttpIps, item)
		}
		if datasource {
			result["allowlisted_http_ips"] = allowlistedHttpIps
		} else {
			result["allowlisted_http_ips"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, allowlistedHttpIps)
		}

		allowlistedHttpVcns := []interface{}{}
		for _, item := range v.AllowlistedHttpVcns {
			allowlistedHttpVcns = append(allowlistedHttpVcns, IntegVirtualCloudNetworkToMap(item, datasource))
		}
		if datasource {
			result["allowlisted_http_vcns"] = allowlistedHttpVcns
		} else {
			result["allowlisted_http_vcns"] = schema.NewSet(allowlistedHttpVcnsHashCodeForSets, allowlistedHttpVcns)
		}

		if v.IsIntegrationVcnAllowlisted != nil {
			result["is_integration_vcn_allowlisted"] = bool(*v.IsIntegrationVcnAllowlisted)
		}
	default:
		log.Printf("[WARN] Received 'network_endpoint_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func OutboundConnectionToMap(obj *oci_integration.OutboundConnection, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_integration.NoneOutboundConnection:
		result["outbound_connection_type"] = "NONE"
	case oci_integration.PrivateEndpointOutboundConnection:
		result["outbound_connection_type"] = "PRIVATE_ENDPOINT"

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		if datasource {
			result["nsg_ids"] = nsgIds
		} else {
			result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
		}

		if v.SubnetId != nil {
			result["subnet_id"] = string(*v.SubnetId)
		}
	default:
		log.Printf("[WARN] Received 'outbound_connection_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *IntegrationIntegrationInstanceResourceCrud) mapToVirtualCloudNetwork(fieldKeyFormat string) (oci_integration.VirtualCloudNetwork, error) {
	result := oci_integration.VirtualCloudNetwork{}

	if allowlistedIps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowlisted_ips")); ok {
		set := allowlistedIps.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowlisted_ips")) {
			result.AllowlistedIps = tmp
		}
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func IntegVirtualCloudNetworkToMap(obj oci_integration.VirtualCloudNetwork, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	allowlistedIps := []interface{}{}
	for _, item := range obj.AllowlistedIps {
		allowlistedIps = append(allowlistedIps, item)
	}
	if datasource {
		result["allowlisted_ips"] = allowlistedIps
	} else {
		result["allowlisted_ips"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, allowlistedIps)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func allowlistedHttpVcnsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if allowlistedIps, ok := m["allowlisted_ips"]; ok && allowlistedIps != "" {
	}
	if id, ok := m["id"]; ok && id != "" {
		buf.WriteString(fmt.Sprintf("%v-", id))
	}
	return utils.GetStringHashcode(buf.String())
}

func alternateCustomEndpointsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if certificateSecretId, ok := m["certificate_secret_id"]; ok && certificateSecretId != "" {
		buf.WriteString(fmt.Sprintf("%v-", certificateSecretId))
	}
	if hostname, ok := m["hostname"]; ok && hostname != "" {
		buf.WriteString(fmt.Sprintf("%v-", hostname))
	}
	return utils.GetStringHashcode(buf.String())
}

func (s *IntegrationIntegrationInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_integration.ChangeIntegrationInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.IntegrationInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration")

	response, err := s.Client.ChangeIntegrationInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIntegrationInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration"), oci_integration.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *IntegrationIntegrationInstanceResourceCrud) StartIntegerationInstance() error {
	state := oci_integration.IntegrationInstanceLifecycleStateActive
	if err := s.Get(); err != nil {
		return err
	}
	if s.Res.LifecycleState == state {
		fmt.Printf("[WARN] The Integration instance already in the wanted state: %v", state)
		return nil
	}
	request := oci_integration.StartIntegrationInstanceRequest{}

	tmp := s.D.Id()
	request.IntegrationInstanceId = &tmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration")

	if _, err := s.Client.StartIntegrationInstance(context.Background(), request); err != nil {
		return err
	}
	resourceChangedFunc := func() bool { return s.Res.LifecycleState == state }

	return tfresource.WaitForResourceCondition(s, resourceChangedFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *IntegrationIntegrationInstanceResourceCrud) StopIntegerationInstance() error {
	state := oci_integration.IntegrationInstanceLifecycleStateInactive
	if err := s.Get(); err != nil {
		return err
	}
	if s.Res.LifecycleState == state {
		fmt.Printf("[WARN] The Integration instance already in the wanted state: %v", state)
		return nil
	}
	request := oci_integration.StopIntegrationInstanceRequest{}

	tmp := s.D.Id()
	request.IntegrationInstanceId = &tmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "integration")

	if _, err := s.Client.StopIntegrationInstance(context.Background(), request); err != nil {
		return err
	}
	resourceChangedFunc := func() bool { return s.Res.LifecycleState == state }

	return tfresource.WaitForResourceCondition(s, resourceChangedFunc, s.D.Timeout(schema.TimeoutUpdate))
}
