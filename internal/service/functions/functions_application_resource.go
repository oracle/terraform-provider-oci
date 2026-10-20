// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/httpreplay"

	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"
)

func FunctionsApplicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFunctionsApplication,
		Read:     readFunctionsApplication,
		Update:   updateFunctionsApplication,
		Delete:   deleteFunctionsApplication,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"config": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
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
			"image_policy_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"is_policy_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional
						"key_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"kms_key_id": {
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
			"logging": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"line_format": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"network_security_group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"security_attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"shape": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"syslog_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trace_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"domain_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
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

func createFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FunctionsWorkRequestManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FunctionsWorkRequestManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).FunctionsWorkRequestManagementClient()

	return tfresource.DeleteResource(d, sync)
}

type FunctionsApplicationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_functions.FunctionsManagementClient
	Res                    *oci_functions.Application
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_functions.WorkRequestManagementClient
}

func (s *FunctionsApplicationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FunctionsApplicationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateCreating),
	}
}

func (s *FunctionsApplicationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateActive),
	}
}

func (s *FunctionsApplicationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateDeleting),
	}
}

func (s *FunctionsApplicationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateDeleted),
	}
}

func (s *FunctionsApplicationResourceCrud) Create() error {
	request := oci_functions.CreateApplicationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if config, ok := s.D.GetOkExists("config"); ok {
		request.Config = tfresource.ObjectMapToStringMap(config.(map[string]interface{}))
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

	if imagePolicyConfig, ok := s.D.GetOkExists("image_policy_config"); ok {
		if tmpList := imagePolicyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "image_policy_config", 0)
			tmp, err := s.mapToImagePolicyConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ImagePolicyConfig = &tmp
		}
	}

	if logging, ok := s.D.GetOkExists("logging"); ok {
		if tmpList := logging.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "logging", 0)
			tmp, err := s.mapToApplicationLoggingConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Logging = &tmp
		}
	}

	if networkSecurityGroupIds, ok := s.D.GetOkExists("network_security_group_ids"); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("network_security_group_ids") {
			request.NetworkSecurityGroupIds = tmp
		}
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		request.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		request.Shape = oci_functions.CreateApplicationDetailsShapeEnum(shape.(string))
	}

	if subnetIds, ok := s.D.GetOkExists("subnet_ids"); ok {
		interfaces := subnetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("subnet_ids") {
			request.SubnetIds = tmp
		}
	}

	if syslogUrl, ok := s.D.GetOkExists("syslog_url"); ok {
		tmp := syslogUrl.(string)
		request.SyslogUrl = &tmp
	}

	if traceConfig, ok := s.D.GetOkExists("trace_config"); ok {
		if tmpList := traceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trace_config", 0)
			tmp, err := s.mapToApplicationTraceConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TraceConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.CreateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getApplicationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions"), oci_functions.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

// getApplicationFromWorkRequest waits for an async application operation and refreshes Terraform state.
func (s *FunctionsApplicationResourceCrud) getApplicationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_functions.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	applicationId, err := applicationWaitForWorkRequest(workId, "application",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, applicationId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(context.Background(),
			oci_functions.CancelWorkRequestRequest{
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
	s.D.SetId(*applicationId)

	return s.Get()
}

// applicationWorkRequestShouldRetryFunc keeps polling until the work request finishes or times out.
func applicationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "functions", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_functions.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

// applicationWaitForWorkRequest polls a work request and returns the affected application identifier.
func applicationWaitForWorkRequest(wId *string, entityType string, action oci_functions.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_functions.WorkRequestManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "functions")
	retryPolicy.ShouldRetryOperation = applicationWorkRequestShouldRetryFunc(timeout)

	response := oci_functions.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_functions.OperationStatusInProgress),
			string(oci_functions.OperationStatusAccepted),
			string(oci_functions.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_functions.OperationStatusSucceeded),
			string(oci_functions.OperationStatusFailed),
			string(oci_functions.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_functions.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_functions.OperationStatusFailed || response.Status == oci_functions.OperationStatusCanceled {
		return nil, getErrorFromFunctionsApplicationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

// getErrorFromFunctionsApplicationWorkRequest converts OCI work request errors into a Terraform error.
func getErrorFromFunctionsApplicationWorkRequest(client *oci_functions.WorkRequestManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_functions.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_functions.ListWorkRequestErrorsRequest{
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

func (s *FunctionsApplicationResourceCrud) Get() error {
	request := oci_functions.GetApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.GetApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *FunctionsApplicationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_functions.UpdateApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	if config, ok := s.D.GetOkExists("config"); ok {
		request.Config = tfresource.ObjectMapToStringMap(config.(map[string]interface{}))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if imagePolicyConfig, ok := s.D.GetOkExists("image_policy_config"); ok {
		if tmpList := imagePolicyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "image_policy_config", 0)
			tmp, err := s.mapToImagePolicyConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ImagePolicyConfig = &tmp
		}
	}

	if logging, ok := s.D.GetOkExists("logging"); ok {
		if tmpList := logging.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "logging", 0)
			tmp, err := s.mapToApplicationLoggingConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Logging = &tmp
		}
	}

	if networkSecurityGroupIds, ok := s.D.GetOkExists("network_security_group_ids"); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("network_security_group_ids") {
			request.NetworkSecurityGroupIds = tmp
		}
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		request.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
	}

	if syslogUrl, ok := s.D.GetOkExists("syslog_url"); ok {
		tmp := syslogUrl.(string)
		request.SyslogUrl = &tmp
	}

	if traceConfig, ok := s.D.GetOkExists("trace_config"); ok && s.D.HasChange("trace_config") {
		if tmpList := traceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trace_config", 0)
			tmp, err := s.mapToApplicationTraceConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TraceConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.UpdateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getApplicationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions"), oci_functions.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FunctionsApplicationResourceCrud) Delete() error {
	request := oci_functions.DeleteApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.DeleteApplication(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := applicationWaitForWorkRequest(workId, "application",
		oci_functions.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *FunctionsApplicationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("config", s.Res.Config)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImagePolicyConfig != nil {
		s.D.Set("image_policy_config", []interface{}{ImagePolicyConfigToMapFunctions(s.Res.ImagePolicyConfig)})
	} else {
		s.D.Set("image_policy_config", nil)
	}

	if s.Res.Logging != nil {
		s.D.Set("logging", []interface{}{ApplicationLoggingConfigToMap(s.Res.Logging)})
	} else {
		s.D.Set("logging", nil)
	}

	networkSecurityGroupIds := []interface{}{}
	for _, item := range s.Res.NetworkSecurityGroupIds {
		networkSecurityGroupIds = append(networkSecurityGroupIds, item)
	}
	s.D.Set("network_security_group_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, networkSecurityGroupIds))

	s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))

	s.D.Set("shape", s.Res.Shape)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subnet_ids", s.Res.SubnetIds)

	if s.Res.SyslogUrl != nil {
		s.D.Set("syslog_url", *s.Res.SyslogUrl)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TraceConfig != nil {
		s.D.Set("trace_config", []interface{}{ApplicationTraceConfigToMap(s.Res.TraceConfig)})
	} else {
		s.D.Set("trace_config", nil)
	}

	return nil
}

func (s *FunctionsApplicationResourceCrud) mapToApplicationLoggingConfig(fieldKeyFormat string) (oci_functions.ApplicationLoggingConfig, error) {
	result := oci_functions.ApplicationLoggingConfig{}

	if lineFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "line_format")); ok {
		result.LineFormat = oci_functions.ApplicationLoggingConfigLineFormatEnum(lineFormat.(string))
	}

	return result, nil
}

func ApplicationLoggingConfigToMap(obj *oci_functions.ApplicationLoggingConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["line_format"] = string(obj.LineFormat)

	return result
}

func (s *FunctionsApplicationResourceCrud) mapToApplicationTraceConfig(fieldKeyFormat string) (oci_functions.ApplicationTraceConfig, error) {
	result := oci_functions.ApplicationTraceConfig{}

	if domainId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_id")); ok {
		tmp := domainId.(string)
		result.DomainId = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func ApplicationTraceConfigToMap(obj *oci_functions.ApplicationTraceConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DomainId != nil {
		result["domain_id"] = string(*obj.DomainId)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func (s *FunctionsApplicationResourceCrud) mapToImagePolicyConfig(fieldKeyFormat string) (oci_functions.ImagePolicyConfig, error) {
	result := oci_functions.ImagePolicyConfig{}

	if isPolicyEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_policy_enabled")); ok {
		tmp := isPolicyEnabled.(bool)
		result.IsPolicyEnabled = &tmp
	}

	if keyDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_details")); ok {
		interfaces := keyDetails.([]interface{})
		tmp := make([]oci_functions.KeyDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_details"), stateDataIndex)
			converted, err := s.mapToKeyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "key_details")) {
			result.KeyDetails = tmp
		}
	}

	return result, nil
}

func ImagePolicyConfigToMapFunctions(obj *oci_functions.ImagePolicyConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsPolicyEnabled != nil {
		result["is_policy_enabled"] = bool(*obj.IsPolicyEnabled)
	}

	keyDetails := []interface{}{}
	for _, item := range obj.KeyDetails {
		keyDetails = append(keyDetails, KeyDetailsToMapFunctions(item))
	}
	result["key_details"] = keyDetails

	return result
}

func (s *FunctionsApplicationResourceCrud) mapToKeyDetails(fieldKeyFormat string) (oci_functions.KeyDetails, error) {
	result := oci_functions.KeyDetails{}

	if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
		tmp := kmsKeyId.(string)
		result.KmsKeyId = &tmp
	}

	return result, nil
}

func KeyDetailsToMapFunctions(obj oci_functions.KeyDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}

	return result
}

func (s *FunctionsApplicationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_functions.ChangeApplicationCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ApplicationId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.ChangeApplicationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getApplicationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions"), oci_functions.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FunctionsApplicationResourceCrud) ExtraWaitPostDelete() time.Duration {
	if httpreplay.ShouldRetryImmediately() {
		return time.Duration(1 * time.Second)
	}
	log.Printf("[DEBUG] Waiting for 5 minutes post destroy of application resource due to known service issue")
	return time.Duration(5 * time.Minute)
}
