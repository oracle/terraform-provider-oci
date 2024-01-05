// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_mesh

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ServiceMeshVirtualDeploymentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createServiceMeshVirtualDeployment,
		Read:     readServiceMeshVirtualDeployment,
		Update:   updateServiceMeshVirtualDeployment,
		Delete:   deleteServiceMeshVirtualDeployment,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"virtual_service_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"access_logging": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_enabled": {
							Type:     schema.TypeBool,
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
			"listeners": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"idle_timeout_in_ms": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"request_timeout_in_ms": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"service_discovery": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DISABLED",
								"DNS",
							}, true),
						},

						// Optional
						"hostname": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

func createServiceMeshVirtualDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.CreateResource(d, sync)
}

func readServiceMeshVirtualDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

func updateServiceMeshVirtualDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteServiceMeshVirtualDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ServiceMeshVirtualDeploymentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_service_mesh.ServiceMeshClient
	Res                    *oci_service_mesh.VirtualDeployment
	DisableNotFoundRetries bool
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_service_mesh.VirtualDeploymentLifecycleStateCreating),
	}
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_service_mesh.VirtualDeploymentLifecycleStateActive),
	}
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_service_mesh.VirtualDeploymentLifecycleStateDeleting),
	}
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_service_mesh.VirtualDeploymentLifecycleStateDeleted),
	}
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) Create() error {
	request := oci_service_mesh.CreateVirtualDeploymentRequest{}

	if accessLogging, ok := s.D.GetOkExists("access_logging"); ok {
		if tmpList := accessLogging.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "access_logging", 0)
			tmp, err := s.mapToAccessLoggingConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AccessLogging = &tmp
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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if listeners, ok := s.D.GetOkExists("listeners"); ok {
		interfaces := listeners.([]interface{})
		tmp := make([]oci_service_mesh.VirtualDeploymentListener, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "listeners", stateDataIndex)
			converted, err := s.mapToVirtualDeploymentListener(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("listeners") {
			request.Listeners = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if serviceDiscovery, ok := s.D.GetOkExists("service_discovery"); ok {
		if tmpList := serviceDiscovery.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "service_discovery", 0)
			tmp, err := s.mapToServiceDiscoveryConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ServiceDiscovery = tmp
		}
	}

	if virtualServiceId, ok := s.D.GetOkExists("virtual_service_id"); ok {
		tmp := virtualServiceId.(string)
		request.VirtualServiceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.CreateVirtualDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getVirtualDeploymentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) getVirtualDeploymentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_service_mesh.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	virtualDeploymentId, err := virtualDeploymentWaitForWorkRequest(workId, "virtualdeployment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, virtualDeploymentId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_service_mesh.CancelWorkRequestRequest{
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
	s.D.SetId(*virtualDeploymentId)

	return s.Get()
}

func virtualDeploymentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "service_mesh", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_service_mesh.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func virtualDeploymentWaitForWorkRequest(wId *string, entityType string, action oci_service_mesh.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_service_mesh.ServiceMeshClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "service_mesh")
	retryPolicy.ShouldRetryOperation = virtualDeploymentWorkRequestShouldRetryFunc(timeout)

	response := oci_service_mesh.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_service_mesh.OperationStatusInProgress),
			string(oci_service_mesh.OperationStatusAccepted),
			string(oci_service_mesh.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_service_mesh.OperationStatusSucceeded),
			string(oci_service_mesh.OperationStatusFailed),
			string(oci_service_mesh.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_service_mesh.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_service_mesh.OperationStatusFailed || response.Status == oci_service_mesh.OperationStatusCanceled {
		return nil, getErrorFromServiceMeshVirtualDeploymentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromServiceMeshVirtualDeploymentWorkRequest(client *oci_service_mesh.ServiceMeshClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_service_mesh.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_service_mesh.ListWorkRequestErrorsRequest{
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

func (s *ServiceMeshVirtualDeploymentResourceCrud) Get() error {
	request := oci_service_mesh.GetVirtualDeploymentRequest{}

	tmp := s.D.Id()
	request.VirtualDeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.GetVirtualDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VirtualDeployment
	return nil
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_service_mesh.UpdateVirtualDeploymentRequest{}

	if accessLogging, ok := s.D.GetOkExists("access_logging"); ok {
		if tmpList := accessLogging.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "access_logging", 0)
			tmp, err := s.mapToAccessLoggingConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AccessLogging = &tmp
		}
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

	if listeners, ok := s.D.GetOkExists("listeners"); ok {
		interfaces := listeners.([]interface{})
		tmp := make([]oci_service_mesh.VirtualDeploymentListener, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "listeners", stateDataIndex)
			converted, err := s.mapToVirtualDeploymentListener(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("listeners") {
			request.Listeners = tmp
		}
	}

	if serviceDiscovery, ok := s.D.GetOkExists("service_discovery"); ok {
		if tmpList := serviceDiscovery.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "service_discovery", 0)
			tmp, err := s.mapToServiceDiscoveryConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ServiceDiscovery = tmp
		}
	}

	tmp := s.D.Id()
	request.VirtualDeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.UpdateVirtualDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVirtualDeploymentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) Delete() error {
	request := oci_service_mesh.DeleteVirtualDeploymentRequest{}

	tmp := s.D.Id()
	request.VirtualDeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.DeleteVirtualDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := virtualDeploymentWaitForWorkRequest(workId, "virtualdeployment",
		oci_service_mesh.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) SetData() error {
	if s.Res.AccessLogging != nil {
		s.D.Set("access_logging", []interface{}{AccessLoggingConfigurationToMap(s.Res.AccessLogging)})
	} else {
		s.D.Set("access_logging", nil)
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	listeners := []interface{}{}
	for _, item := range s.Res.Listeners {
		listeners = append(listeners, VirtualDeploymentListenerToMap(item))
	}
	s.D.Set("listeners", listeners)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ServiceDiscovery != nil {
		serviceDiscoveryArray := []interface{}{}
		if serviceDiscoveryMap := ServiceDiscoveryConfigurationToMap(&s.Res.ServiceDiscovery); serviceDiscoveryMap != nil {
			serviceDiscoveryArray = append(serviceDiscoveryArray, serviceDiscoveryMap)
		}
		s.D.Set("service_discovery", serviceDiscoveryArray)
	} else {
		s.D.Set("service_discovery", nil)
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

	if s.Res.VirtualServiceId != nil {
		s.D.Set("virtual_service_id", *s.Res.VirtualServiceId)
	}

	return nil
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) mapToAccessLoggingConfiguration(fieldKeyFormat string) (oci_service_mesh.AccessLoggingConfiguration, error) {
	result := oci_service_mesh.AccessLoggingConfiguration{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) mapToServiceDiscoveryConfiguration(fieldKeyFormat string) (oci_service_mesh.ServiceDiscoveryConfiguration, error) {
	var baseObject oci_service_mesh.ServiceDiscoveryConfiguration
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DISABLED"):
		details := oci_service_mesh.DisabledServiceDiscoveryConfiguration{}
		baseObject = details
	case strings.ToLower("DNS"):
		details := oci_service_mesh.DnsServiceDiscoveryConfiguration{}
		if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ServiceDiscoveryConfigurationToMap(obj *oci_service_mesh.ServiceDiscoveryConfiguration) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_service_mesh.DisabledServiceDiscoveryConfiguration:
		result["type"] = "DISABLED"
	case oci_service_mesh.DnsServiceDiscoveryConfiguration:
		result["type"] = "DNS"

		if v.Hostname != nil {
			result["hostname"] = string(*v.Hostname)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) mapToVirtualDeploymentListener(fieldKeyFormat string) (oci_service_mesh.VirtualDeploymentListener, error) {
	result := oci_service_mesh.VirtualDeploymentListener{}

	if idleTimeoutInMs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idle_timeout_in_ms")); ok {
		tmp := idleTimeoutInMs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert idleTimeoutInMs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.IdleTimeoutInMs = &tmpInt64
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_service_mesh.VirtualDeploymentListenerProtocolEnum(protocol.(string))
	}

	if requestTimeoutInMs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_timeout_in_ms")); ok {
		tmp := requestTimeoutInMs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert requestTimeoutInMs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.RequestTimeoutInMs = &tmpInt64
	}

	return result, nil
}

func VirtualDeploymentListenerToMap(obj oci_service_mesh.VirtualDeploymentListener) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IdleTimeoutInMs != nil {
		result["idle_timeout_in_ms"] = strconv.FormatInt(*obj.IdleTimeoutInMs, 10)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	result["protocol"] = string(obj.Protocol)

	if obj.RequestTimeoutInMs != nil {
		result["request_timeout_in_ms"] = strconv.FormatInt(*obj.RequestTimeoutInMs, 10)
	}

	return result
}

func VirtualDeploymentSummaryToMap(obj oci_service_mesh.VirtualDeploymentSummary) map[string]interface{} {
	result := map[string]interface{}{}

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

	if obj.VirtualServiceId != nil {
		result["virtual_service_id"] = string(*obj.VirtualServiceId)
	}

	return result
}

func (s *ServiceMeshVirtualDeploymentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_service_mesh.ChangeVirtualDeploymentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VirtualDeploymentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.ChangeVirtualDeploymentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVirtualDeploymentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
