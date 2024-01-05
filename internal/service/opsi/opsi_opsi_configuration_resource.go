// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

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
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiOpsiConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiOpsiConfiguration,
		Read:     readOpsiOpsiConfiguration,
		Update:   updateOpsiOpsiConfiguration,
		Delete:   deleteOpsiOpsiConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"opsi_config_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config_item_custom_status": {
				Type:     schema.TypeList,
				Optional: true,
				//Computed: true,
				//ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"config_item_field": {
				Type:     schema.TypeList,
				Optional: true,
				//Computed: true,
				//ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"config_items": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"config_item_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"BASIC",
							}, true),
						},

						// Optional
						"name": {
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
						"applicable_contexts": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"default_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"config_item_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"data_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unit_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"unit": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"value_input_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_value_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"max_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"min_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"possible_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"config_items_applicable_context": {
				Type:     schema.TypeList,
				Optional: true,
				//Computed: true,
				//ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"opsi_config_field": {
				Type:     schema.TypeList,
				Optional: true,
				//Computed: true,
				//ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
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

func createOpsiOpsiConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOpsiConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.CreateResource(d, sync)
}

func readOpsiOpsiConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOpsiConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

func updateOpsiOpsiConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOpsiConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOpsiOpsiConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOpsiConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpsiOpsiConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *oci_opsi.OpsiConfiguration
	DisableNotFoundRetries bool
}

func (s *OpsiOpsiConfigurationResourceCrud) ID() string {
	opsiConfiguration := *s.Res
	return *opsiConfiguration.GetId()
}

func (s *OpsiOpsiConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opsi.OpsiConfigurationLifecycleStateCreating),
	}
}

func (s *OpsiOpsiConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opsi.OpsiConfigurationLifecycleStateActive),
	}
}

func (s *OpsiOpsiConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opsi.OpsiConfigurationLifecycleStateDeleting),
	}
}

func (s *OpsiOpsiConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opsi.OpsiConfigurationLifecycleStateDeleted),
	}
}

func (s *OpsiOpsiConfigurationResourceCrud) Create() error {
	request := oci_opsi.CreateOpsiConfigurationRequest{}
	err := s.populateTopLevelPolymorphicCreateOpsiConfigurationRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.CreateOpsiConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getOpsiConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OpsiOpsiConfigurationResourceCrud) getOpsiConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	opsiConfigurationId, err := opsiConfigurationWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*opsiConfigurationId)

	return s.Get()
}

func opsiConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "opsi", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_opsi.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func opsiConfigurationWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = opsiConfigurationWorkRequestShouldRetryFunc(timeout)

	response := oci_opsi.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_opsi.OperationStatusInProgress),
			string(oci_opsi.OperationStatusAccepted),
			string(oci_opsi.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_opsi.OperationStatusSucceeded),
			string(oci_opsi.OperationStatusFailed),
			string(oci_opsi.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_opsi.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_opsi.OperationStatusFailed || response.Status == oci_opsi.OperationStatusCanceled {
		return nil, getErrorFromOpsiOpsiConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiOpsiConfigurationWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_opsi.ListWorkRequestErrorsRequest{
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

func (s *OpsiOpsiConfigurationResourceCrud) Get() error {
	request := oci_opsi.GetOpsiConfigurationRequest{}

	if configItemCustomStatus, ok := s.D.GetOkExists("config_item_custom_status"); ok {
		interfaces := configItemCustomStatus.([]interface{})
		tmp := make([]string, len(interfaces))
		tmp2 := make([]oci_opsi.GetOpsiConfigurationConfigItemCustomStatusEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
				tmp2[i], _ = oci_opsi.GetMappingGetOpsiConfigurationConfigItemCustomStatusEnum(strings.ToLower(tmp[i]))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("config_item_custom_status") {
			request.ConfigItemCustomStatus = tmp2
		}
	}

	if configItemField, ok := s.D.GetOkExists("config_item_field"); ok {
		interfaces := configItemField.([]interface{})
		tmp := make([]string, len(interfaces))
		tmp2 := make([]oci_opsi.GetOpsiConfigurationConfigItemFieldEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
				tmp2[i], _ = oci_opsi.GetMappingGetOpsiConfigurationConfigItemFieldEnum(strings.ToLower(tmp[i]))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("config_item_field") {
			request.ConfigItemField = tmp2
		}
	}

	if configItemsApplicableContext, ok := s.D.GetOkExists("config_items_applicable_context"); ok {
		interfaces := configItemsApplicableContext.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("config_items_applicable_context") {
			request.ConfigItemsApplicableContext = tmp
		}
	}

	if opsiConfigField, ok := s.D.GetOkExists("opsi_config_field"); ok {
		interfaces := opsiConfigField.([]interface{})
		tmp := make([]string, len(interfaces))
		tmp2 := make([]oci_opsi.GetOpsiConfigurationOpsiConfigFieldEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
				tmp2[i], _ = oci_opsi.GetMappingGetOpsiConfigurationOpsiConfigFieldEnum(strings.ToLower(tmp[i]))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("opsi_config_field") {
			request.OpsiConfigField = tmp2
		}
	}

	tmp := s.D.Id()
	request.OpsiConfigurationId = &tmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
	response, err := s.Client.GetOpsiConfiguration(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response.OpsiConfiguration
	return nil
}

func (s *OpsiOpsiConfigurationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_opsi.UpdateOpsiConfigurationRequest{}
	err := s.populateTopLevelPolymorphicUpdateOpsiConfigurationRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.UpdateOpsiConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOpsiConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpsiOpsiConfigurationResourceCrud) Delete() error {
	request := oci_opsi.DeleteOpsiConfigurationRequest{}

	tmp := s.D.Id()
	request.OpsiConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DeleteOpsiConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := opsiConfigurationWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpsiOpsiConfigurationResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_opsi.OpsiUxConfiguration:
		s.D.Set("opsi_config_type", "UX_CONFIGURATION")

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.ConfigItems != nil && len(v.ConfigItems) != 0 {
			//if v.ConfigItems != nil {
			configItems := []interface{}{}
			for _, item := range v.ConfigItems {
				configItems = append(configItems, OpsiConfigurationConfigurationItemSummaryToMap(item))
			}
			s.D.Set("config_items", configItems)
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

		if v.FreeformTags != nil {
			s.D.Set("freeform_tags", v.FreeformTags)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}
	default:
		log.Printf("[WARN] Received 'opsi_config_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *OpsiOpsiConfigurationResourceCrud) mapToConfigurationItemAllowedValueDetails(fieldKeyFormat string) (oci_opsi.ConfigurationItemAllowedValueDetails, error) {
	var baseObject oci_opsi.ConfigurationItemAllowedValueDetails
	//discriminator
	allowedValueTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_value_type"))
	var allowedValueType string
	if ok {
		allowedValueType = allowedValueTypeRaw.(string)
	} else {
		allowedValueType = "" // default value
	}
	switch strings.ToLower(allowedValueType) {
	case strings.ToLower("FREE_TEXT"):
		details := oci_opsi.ConfigurationItemFreeTextAllowedValueDetails{}
		baseObject = details
	case strings.ToLower("LIMIT"):
		details := oci_opsi.ConfigurationItemLimitAllowedValueDetails{}
		if maxValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_value")); ok {
			tmp := maxValue.(string)
			details.MaxValue = &tmp
		}
		if minValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_value")); ok {
			tmp := minValue.(string)
			details.MinValue = &tmp
		}
		baseObject = details
	case strings.ToLower("PICK"):
		details := oci_opsi.ConfigurationItemPickAllowedValueDetails{}
		if possibleValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "possible_values")); ok {
			interfaces := possibleValues.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "possible_values")) {
				details.PossibleValues = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown allowed_value_type '%v' was specified", allowedValueType)
	}
	return baseObject, nil
}

func ConfigurationItemAllowedValueDetailsToMap(obj *oci_opsi.ConfigurationItemAllowedValueDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_opsi.ConfigurationItemFreeTextAllowedValueDetails:
		result["allowed_value_type"] = "FREE_TEXT"
	case oci_opsi.ConfigurationItemLimitAllowedValueDetails:
		result["allowed_value_type"] = "LIMIT"

		if v.MaxValue != nil {
			result["max_value"] = string(*v.MaxValue)
		}

		if v.MinValue != nil {
			result["min_value"] = string(*v.MinValue)
		}
	case oci_opsi.ConfigurationItemPickAllowedValueDetails:
		result["allowed_value_type"] = "PICK"

		result["possible_values"] = v.PossibleValues
		//result["possible_values"] = v.PossibleValues
	default:
		log.Printf("[WARN] Received 'allowed_value_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func ConfigurationItemMetadataToMap(obj *oci_opsi.ConfigurationItemMetadata) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_opsi.BasicConfigurationItemMetadata:
		result["config_item_type"] = "BASIC"

		if v.DataType != nil {
			result["data_type"] = string(*v.DataType)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.UnitDetails != nil {
			result["unit_details"] = []interface{}{ConfigurationItemUnitDetailsToMap(v.UnitDetails)}
		}

		if v.ValueInputDetails != nil {
			valueInputDetailsArray := []interface{}{}
			if valueInputDetailsMap := ConfigurationItemAllowedValueDetailsToMap(&v.ValueInputDetails); valueInputDetailsMap != nil {
				valueInputDetailsArray = append(valueInputDetailsArray, valueInputDetailsMap)
			}
			result["value_input_details"] = valueInputDetailsArray
		}
	default:
		log.Printf("[WARN] Received 'config_item_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *OpsiOpsiConfigurationResourceCrud) mapToConfigurationItemUnitDetails(fieldKeyFormat string) (oci_opsi.ConfigurationItemUnitDetails, error) {
	result := oci_opsi.ConfigurationItemUnitDetails{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if unit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit")); ok {
		tmp := unit.(string)
		result.Unit = &tmp
	}

	return result, nil
}

func ConfigurationItemUnitDetailsToMap(obj *oci_opsi.ConfigurationItemUnitDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Unit != nil {
		result["unit"] = string(*obj.Unit)
	}

	return result
}

func (s *OpsiOpsiConfigurationResourceCrud) mapToCreateConfigurationItemDetails(fieldKeyFormat string) (oci_opsi.CreateConfigurationItemDetails, error) {
	var baseObject oci_opsi.CreateConfigurationItemDetails
	//discriminator
	configItemTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_item_type"))
	var configItemType string
	if ok {
		configItemType = configItemTypeRaw.(string)
	} else {
		configItemType = "" // default value
	}
	switch strings.ToLower(configItemType) {
	case strings.ToLower("BASIC"):
		details := oci_opsi.CreateBasicConfigurationItemDetails{}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown config_item_type '%v' was specified", configItemType)
	}
	return baseObject, nil
}

func (s *OpsiOpsiConfigurationResourceCrud) mapToUpdateConfigurationItemDetails(fieldKeyFormat string) (oci_opsi.UpdateConfigurationItemDetails, error) {
	var baseObject oci_opsi.UpdateConfigurationItemDetails
	//discriminator
	configItemTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_item_type"))
	var configItemType string
	if ok {
		configItemType = configItemTypeRaw.(string)
	} else {
		configItemType = "" // default value
	}
	switch strings.ToLower(configItemType) {
	case strings.ToLower("BASIC"):
		details := oci_opsi.UpdateBasicConfigurationItemDetails{}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown config_item_type '%v' was specified", configItemType)
	}
	return baseObject, nil
}

func OpsiConfigurationConfigurationItemSummaryToMap(obj oci_opsi.OpsiConfigurationConfigurationItemSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_opsi.OpsiConfigurationBasicConfigurationItemSummary:
		result["config_item_type"] = "BASIC"

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}

		if v.DefaultValue != nil {
			result["default_value"] = string(*v.DefaultValue)
		}

		if v.ApplicableContexts != nil && len(v.ApplicableContexts) != 0 {
			result["applicable_contexts"] = v.ApplicableContexts
		}

		result["config_item_type"] = "BASIC"

		if v.Metadata != nil {
			metadataArray := []interface{}{}
			if metadataMap := ConfigurationItemMetadataToMap(&v.Metadata); metadataMap != nil {
				metadataArray = append(metadataArray, metadataMap)
			}
			result["metadata"] = metadataArray
		}

	default:
		log.Printf("[WARN] Received 'config_item_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func OpsiConfigurationSummaryToMap(obj oci_opsi.OpsiConfigurationSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch (obj).(type) {
	case oci_opsi.OpsiUxConfigurationSummary:

		result["opsi_config_type"] = "UX_CONFIGURATION"

		if obj.GetId() != nil {
			result["id"] = string(*obj.GetId())
		}

		if obj.GetCompartmentId() != nil {
			result["compartment_id"] = string(*obj.GetCompartmentId())
		}

		if obj.GetDisplayName() != nil {
			result["display_name"] = string(*obj.GetDisplayName())
		}

		if obj.GetDescription() != nil {
			result["description"] = string(*obj.GetDescription())
		}

		result["freeform_tags"] = obj.GetFreeformTags()

		if obj.GetDefinedTags() != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
		}

		if obj.GetSystemTags() != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(obj.GetSystemTags())
		}

		if obj.GetTimeCreated() != nil {
			result["time_created"] = obj.GetTimeCreated().String()
		}

		if obj.GetTimeUpdated() != nil {
			result["time_updated"] = obj.GetTimeUpdated().String()
		}

		if obj.GetLifecycleDetails() != nil {
			result["lifecycle_details"] = string(*obj.GetLifecycleDetails())
		}

		result["state"] = string(obj.GetLifecycleState())

	default:
		log.Printf("[WARN] Received 'opsi_config_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *OpsiOpsiConfigurationResourceCrud) populateTopLevelPolymorphicCreateOpsiConfigurationRequest(request *oci_opsi.CreateOpsiConfigurationRequest) error {
	//discriminator
	opsiConfigTypeRaw, ok := s.D.GetOkExists("opsi_config_type")
	var opsiConfigType string
	if ok {
		opsiConfigType = opsiConfigTypeRaw.(string)
	} else {
		opsiConfigType = "" // default value
	}
	switch strings.ToLower(opsiConfigType) {
	case strings.ToLower("UX_CONFIGURATION"):
		details := oci_opsi.CreateOpsiUxConfigurationDetails{}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if configItemCustomStatus, ok := s.D.GetOkExists("config_item_custom_status"); ok {
			interfaces := configItemCustomStatus.([]interface{})
			tmp := make([]string, len(interfaces))
			tmp2 := make([]oci_opsi.CreateOpsiConfigurationConfigItemCustomStatusEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
					tmp2[i], _ = oci_opsi.GetMappingCreateOpsiConfigurationConfigItemCustomStatusEnum(strings.ToLower(tmp[i]))
				}
			}
			if len(tmp2) != 0 || s.D.HasChange("config_item_custom_status") {
				request.ConfigItemCustomStatus = tmp2
			}
		}
		if configItemField, ok := s.D.GetOkExists("config_item_field"); ok {
			interfaces := configItemField.([]interface{})
			tmp := make([]string, len(interfaces))
			tmp2 := make([]oci_opsi.CreateOpsiConfigurationConfigItemFieldEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
					tmp2[i], _ = oci_opsi.GetMappingCreateOpsiConfigurationConfigItemFieldEnum(strings.ToLower(tmp[i]))
				}
			}
			if len(tmp2) != 0 || s.D.HasChange("config_item_field") {
				request.ConfigItemField = tmp2
			}
		}
		if configItems, ok := s.D.GetOkExists("config_items"); ok {
			interfaces := configItems.([]interface{})
			tmp := make([]oci_opsi.CreateConfigurationItemDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_items", stateDataIndex)
				converted, err := s.mapToCreateConfigurationItemDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("config_items") {
				details.ConfigItems = tmp
			}
		}
		if configItemsApplicableContext, ok := s.D.GetOkExists("config_items_applicable_context"); ok {
			interfaces := configItemsApplicableContext.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("config_items_applicable_context") {
				request.ConfigItemsApplicableContext = tmp
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
		if opsiConfigField, ok := s.D.GetOkExists("opsi_config_field"); ok {
			interfaces := opsiConfigField.([]interface{})
			tmp := make([]string, len(interfaces))
			tmp2 := make([]oci_opsi.CreateOpsiConfigurationOpsiConfigFieldEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
					tmp2[i], _ = oci_opsi.GetMappingCreateOpsiConfigurationOpsiConfigFieldEnum(strings.ToLower(tmp[i]))
				}
			}
			if len(tmp2) != 0 || s.D.HasChange("opsi_config_field") {
				request.OpsiConfigField = tmp2
			}
		}
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		request.CreateOpsiConfigurationDetails = details
	default:
		return fmt.Errorf("unknown opsi_config_type '%v' was specified", opsiConfigType)
	}
	return nil
}

func (s *OpsiOpsiConfigurationResourceCrud) populateTopLevelPolymorphicUpdateOpsiConfigurationRequest(request *oci_opsi.UpdateOpsiConfigurationRequest) error {
	//discriminator
	opsiConfigTypeRaw, ok := s.D.GetOkExists("opsi_config_type")
	var opsiConfigType string
	if ok {
		opsiConfigType = opsiConfigTypeRaw.(string)
	} else {
		opsiConfigType = "" // default value
	}
	switch strings.ToLower(opsiConfigType) {
	case strings.ToLower("UX_CONFIGURATION"):
		details := oci_opsi.UpdateOpsiUxConfigurationDetails{}
		if configItems, ok := s.D.GetOkExists("config_items"); ok {
			interfaces := configItems.([]interface{})
			tmp := make([]oci_opsi.UpdateConfigurationItemDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_items", stateDataIndex)
				converted, err := s.mapToUpdateConfigurationItemDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("config_items") {
				details.ConfigItems = tmp
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
		tmp := s.D.Id()
		request.OpsiConfigurationId = &tmp
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		request.UpdateOpsiConfigurationDetails = details
	default:
		return fmt.Errorf("unknown opsi_config_type '%v' was specified", opsiConfigType)
	}
	return nil
}

func (s *OpsiOpsiConfigurationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_opsi.ChangeOpsiConfigurationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OpsiConfigurationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangeOpsiConfigurationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOpsiConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
