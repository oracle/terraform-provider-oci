// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceMlApplicationInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.MlApplicationInstanceTimeout,
		Create:   createDatascienceMlApplicationInstance,
		Read:     readDatascienceMlApplicationInstance,
		Update:   updateDatascienceMlApplicationInstance,
		Delete:   deleteDatascienceMlApplicationInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ml_application_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ml_application_implementation_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"auth_configuration": {
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
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"IAM",
								"IDCS",
							}, true),
						},
						"application_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"domain_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
					},
				},
			},
			"configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_substate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ml_application_implementation_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ml_application_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"prediction_endpoint_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"base_prediction_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"prediction_uris": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"use_case": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
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

func createDatascienceMlApplicationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.CreateResource(d, sync)
}

func readDatascienceMlApplicationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

func updateDatascienceMlApplicationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatascienceMlApplicationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatascienceMlApplicationInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.MlApplicationInstance
	DisableNotFoundRetries bool
}

func (s *DatascienceMlApplicationInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceMlApplicationInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datascience.MlApplicationInstanceLifecycleStateCreating),
	}
}

func (s *DatascienceMlApplicationInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.MlApplicationInstanceLifecycleStateActive),
		string(oci_datascience.MlApplicationInstanceLifecycleStateNeedsAttention),
	}
}

func (s *DatascienceMlApplicationInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datascience.MlApplicationInstanceLifecycleStateDeleting),
	}
}

func (s *DatascienceMlApplicationInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.MlApplicationInstanceLifecycleStateDeleted),
	}
}

func (s *DatascienceMlApplicationInstanceResourceCrud) Create() error {
	request := oci_datascience.CreateMlApplicationInstanceRequest{}

	if authConfiguration, ok := s.D.GetOkExists("auth_configuration"); ok {
		if tmpList := authConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "auth_configuration", 0)
			tmp, err := s.mapToCreateAuthConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AuthConfiguration = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		interfaces := configuration.([]interface{})
		tmp := make([]oci_datascience.ConfigurationProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration", stateDataIndex)
			converted, err := s.mapToConfigurationProperty(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configuration") {
			request.Configuration = tmp
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

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if mlApplicationId, ok := s.D.GetOkExists("ml_application_id"); ok {
		tmp := mlApplicationId.(string)
		request.MlApplicationId = &tmp
	}

	if mlApplicationImplementationId, ok := s.D.GetOkExists("ml_application_implementation_id"); ok {
		tmp := mlApplicationImplementationId.(string)
		request.MlApplicationImplementationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateMlApplicationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getMlApplicationInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatascienceMlApplicationInstanceResourceCrud) getMlApplicationInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datascience.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	mlApplicationInstanceId, err := mlApplicationInstanceWaitForWorkRequest(workId, "mlapplicationinstance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, mlApplicationInstanceId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_datascience.CancelWorkRequestRequest{
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
	s.D.SetId(*mlApplicationInstanceId)

	return s.Get()
}

func mlApplicationInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "datascience", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_datascience.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func mlApplicationInstanceWaitForWorkRequest(wId *string, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datascience.DataScienceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datascience")
	retryPolicy.ShouldRetryOperation = mlApplicationInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_datascience.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_datascience.WorkRequestStatusInProgress),
			string(oci_datascience.WorkRequestStatusAccepted),
			string(oci_datascience.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_datascience.WorkRequestStatusSucceeded),
			string(oci_datascience.WorkRequestStatusFailed),
			string(oci_datascience.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_datascience.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_datascience.WorkRequestStatusFailed || response.Status == oci_datascience.WorkRequestStatusCanceled {
		return nil, getErrorFromDatascienceMlApplicationInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatascienceMlApplicationInstanceWorkRequest(client *oci_datascience.DataScienceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_datascience.ListWorkRequestErrorsRequest{
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

func (s *DatascienceMlApplicationInstanceResourceCrud) Get() error {
	request := oci_datascience.GetMlApplicationInstanceRequest{}

	tmp := s.D.Id()
	request.MlApplicationInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetMlApplicationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MlApplicationInstance
	return nil
}

func (s *DatascienceMlApplicationInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateMlApplicationInstanceRequest{}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		interfaces := configuration.([]interface{})
		tmp := make([]oci_datascience.ConfigurationProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration", stateDataIndex)
			converted, err := s.mapToConfigurationProperty(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configuration") {
			request.Configuration = tmp
		}
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

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if mlApplicationImplementationId, ok := s.D.GetOkExists("ml_application_implementation_id"); ok {
		tmp := mlApplicationImplementationId.(string)
		request.MlApplicationImplementationId = &tmp
	}

	tmp := s.D.Id()
	request.MlApplicationInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateMlApplicationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMlApplicationInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceMlApplicationInstanceResourceCrud) Delete() error {
	request := oci_datascience.DeleteMlApplicationInstanceRequest{}

	tmp := s.D.Id()
	request.MlApplicationInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.DeleteMlApplicationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := mlApplicationInstanceWaitForWorkRequest(workId, "mlapplicationinstance",
		oci_datascience.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatascienceMlApplicationInstanceResourceCrud) SetData() error {
	if s.Res.AuthConfiguration != nil {
		authConfigurationArray := []interface{}{}
		if authConfigurationMap := AuthConfigurationToMap(&s.Res.AuthConfiguration); authConfigurationMap != nil {
			authConfigurationArray = append(authConfigurationArray, authConfigurationMap)
		}
		s.D.Set("auth_configuration", authConfigurationArray)
	} else {
		s.D.Set("auth_configuration", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	configuration := []interface{}{}
	for _, item := range s.Res.Configuration {
		configuration = append(configuration, ConfigurationPropertyToMap(item))
	}
	s.D.Set("configuration", configuration)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_substate", s.Res.LifecycleSubstate)

	if s.Res.MlApplicationId != nil {
		s.D.Set("ml_application_id", *s.Res.MlApplicationId)
	}

	if s.Res.MlApplicationImplementationId != nil {
		s.D.Set("ml_application_implementation_id", *s.Res.MlApplicationImplementationId)
	}

	if s.Res.MlApplicationImplementationName != nil {
		s.D.Set("ml_application_implementation_name", *s.Res.MlApplicationImplementationName)
	}

	if s.Res.MlApplicationName != nil {
		s.D.Set("ml_application_name", *s.Res.MlApplicationName)
	}

	if s.Res.PredictionEndpointDetails != nil {
		s.D.Set("prediction_endpoint_details", []interface{}{PredictionEndpointDetailsToMap(s.Res.PredictionEndpointDetails)})
	} else {
		s.D.Set("prediction_endpoint_details", nil)
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

func (s *DatascienceMlApplicationInstanceResourceCrud) mapToConfigurationProperty(fieldKeyFormat string) (oci_datascience.ConfigurationProperty, error) {
	result := oci_datascience.ConfigurationProperty{}

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

func ConfigurationPropertyToMap(obj oci_datascience.ConfigurationProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DatascienceMlApplicationInstanceResourceCrud) mapToCreateAuthConfigurationDetails(fieldKeyFormat string) (oci_datascience.CreateAuthConfigurationDetails, error) {
	var baseObject oci_datascience.CreateAuthConfigurationDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("IAM"):
		details := oci_datascience.CreateIamAuthConfigurationCreateDetails{}
		baseObject = details
	case strings.ToLower("IDCS"):
		details := oci_datascience.CreateIdcsAuthConfigurationDetails{}
		if domainId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_id")); ok {
			tmp := domainId.(string)
			details.DomainId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func AuthConfigurationToMap(obj *oci_datascience.AuthConfiguration) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.IamAuthConfiguration:
		result["type"] = "IAM"
	case oci_datascience.IdcsAuthConfiguration:
		result["type"] = "IDCS"

		if v.DomainId != nil {
			result["domain_id"] = string(*v.DomainId)
		}
		if v.ApplicationName != nil {
			result["application_name"] = string(*v.ApplicationName)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func MlApplicationInstanceSummaryToMap(obj oci_datascience.MlApplicationInstanceSummary) map[string]interface{} {
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

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["lifecycle_substate"] = string(obj.LifecycleSubstate)

	if obj.MlApplicationId != nil {
		result["ml_application_id"] = string(*obj.MlApplicationId)
	}

	if obj.MlApplicationImplementationId != nil {
		result["ml_application_implementation_id"] = string(*obj.MlApplicationImplementationId)
	}

	if obj.MlApplicationImplementationName != nil {
		result["ml_application_implementation_name"] = string(*obj.MlApplicationImplementationName)
	}

	if obj.MlApplicationName != nil {
		result["ml_application_name"] = string(*obj.MlApplicationName)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func PredictionEndpointDetailsToMap(obj *oci_datascience.PredictionEndpointDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BasePredictionUri != nil {
		result["base_prediction_uri"] = string(*obj.BasePredictionUri)
	}

	predictionUris := []interface{}{}
	for _, item := range obj.PredictionUris {
		predictionUris = append(predictionUris, PredictionUriToMap(item))
	}
	result["prediction_uris"] = predictionUris

	return result
}

func PredictionUriToMap(obj oci_datascience.PredictionUri) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Uri != nil {
		result["uri"] = string(*obj.Uri)
	}

	if obj.UseCase != nil {
		result["use_case"] = string(*obj.UseCase)
	}

	return result
}

func (s *DatascienceMlApplicationInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeMlApplicationInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MlApplicationInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.ChangeMlApplicationInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMlApplicationInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
