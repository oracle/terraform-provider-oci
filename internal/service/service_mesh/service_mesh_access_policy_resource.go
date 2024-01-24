// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_mesh

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
	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ServiceMeshAccessPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createServiceMeshAccessPolicy,
		Read:     readServiceMeshAccessPolicy,
		Update:   updateServiceMeshAccessPolicy,
		Delete:   deleteServiceMeshAccessPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mesh_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"destination": {
							Type:     schema.TypeList,
							Required: true,
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
											"ALL_VIRTUAL_SERVICES",
											"EXTERNAL_SERVICE",
											"INGRESS_GATEWAY",
											"VIRTUAL_SERVICE",
										}, true),
									},

									// Optional
									"hostnames": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ingress_gateway_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ip_addresses": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ports": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 65535,
										MinItems: 1,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"protocol": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"virtual_service_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"source": {
							Type:     schema.TypeList,
							Required: true,
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
											"ALL_VIRTUAL_SERVICES",
											"EXTERNAL_SERVICE",
											"INGRESS_GATEWAY",
											"VIRTUAL_SERVICE",
										}, true),
									},

									// Optional
									"hostnames": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ingress_gateway_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ip_addresses": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ports": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 65535,
										MinItems: 1,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"protocol": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"virtual_service_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
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

func createServiceMeshAccessPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshAccessPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.CreateResource(d, sync)
}

func readServiceMeshAccessPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshAccessPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

func updateServiceMeshAccessPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshAccessPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteServiceMeshAccessPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshAccessPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ServiceMeshAccessPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_service_mesh.ServiceMeshClient
	Res                    *oci_service_mesh.AccessPolicy
	DisableNotFoundRetries bool
}

func (s *ServiceMeshAccessPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ServiceMeshAccessPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_service_mesh.AccessPolicyLifecycleStateCreating),
	}
}

func (s *ServiceMeshAccessPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_service_mesh.AccessPolicyLifecycleStateActive),
	}
}

func (s *ServiceMeshAccessPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_service_mesh.AccessPolicyLifecycleStateDeleting),
	}
}

func (s *ServiceMeshAccessPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_service_mesh.AccessPolicyLifecycleStateDeleted),
	}
}

func (s *ServiceMeshAccessPolicyResourceCrud) Create() error {
	request := oci_service_mesh.CreateAccessPolicyRequest{}

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

	if meshId, ok := s.D.GetOkExists("mesh_id"); ok {
		tmp := meshId.(string)
		request.MeshId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if rules, ok := s.D.GetOkExists("rules"); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_service_mesh.AccessPolicyRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToAccessPolicyRuleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("rules") {
			request.Rules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.CreateAccessPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getAccessPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ServiceMeshAccessPolicyResourceCrud) getAccessPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_service_mesh.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	accessPolicyId, err := accessPolicyWaitForWorkRequest(workId, "accesspolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, accessPolicyId)
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
	s.D.SetId(*accessPolicyId)

	return s.Get()
}

func accessPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func accessPolicyWaitForWorkRequest(wId *string, entityType string, action oci_service_mesh.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_service_mesh.ServiceMeshClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "service_mesh")
	retryPolicy.ShouldRetryOperation = accessPolicyWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromServiceMeshAccessPolicyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromServiceMeshAccessPolicyWorkRequest(client *oci_service_mesh.ServiceMeshClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_service_mesh.ActionTypeEnum) error {
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

func (s *ServiceMeshAccessPolicyResourceCrud) Get() error {
	request := oci_service_mesh.GetAccessPolicyRequest{}

	tmp := s.D.Id()
	request.AccessPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.GetAccessPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AccessPolicy
	return nil
}

func (s *ServiceMeshAccessPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_service_mesh.UpdateAccessPolicyRequest{}

	tmp := s.D.Id()
	request.AccessPolicyId = &tmp

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

	if rules, ok := s.D.GetOkExists("rules"); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_service_mesh.AccessPolicyRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToAccessPolicyRuleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("rules") {
			request.Rules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.UpdateAccessPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAccessPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ServiceMeshAccessPolicyResourceCrud) Delete() error {
	request := oci_service_mesh.DeleteAccessPolicyRequest{}

	tmp := s.D.Id()
	request.AccessPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.DeleteAccessPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := accessPolicyWaitForWorkRequest(workId, "accesspolicy",
		oci_service_mesh.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ServiceMeshAccessPolicyResourceCrud) SetData() error {
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

	if s.Res.MeshId != nil {
		s.D.Set("mesh_id", *s.Res.MeshId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	rules := []interface{}{}
	for _, item := range s.Res.Rules {
		rules = append(rules, AccessPolicyRuleToMap(item))
	}
	s.D.Set("rules", rules)

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

func (s *ServiceMeshAccessPolicyResourceCrud) mapToAccessPolicyRuleDetails(fieldKeyFormat string) (oci_service_mesh.AccessPolicyRuleDetails, error) {
	result := oci_service_mesh.AccessPolicyRuleDetails{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_service_mesh.AccessPolicyRuleDetailsActionEnum(action.(string))
	}

	if destination, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination")); ok {
		if tmpList := destination.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destination"), 0)
			tmp, err := s.mapToAccessPolicyTargetDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert destination, encountered error: %v", err)
			}
			result.Destination = tmp
		}
	}

	if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
			tmp, err := s.mapToAccessPolicyTargetDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source, encountered error: %v", err)
			}
			result.Source = tmp
		}
	}

	return result, nil
}

func AccessPolicyRuleToMap(obj oci_service_mesh.AccessPolicyRule) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.Destination != nil {
		destinationArray := []interface{}{}
		if destinationMap := AccessPolicyTargetToMap(&obj.Destination); destinationMap != nil {
			destinationArray = append(destinationArray, destinationMap)
		}
		result["destination"] = destinationArray
	}

	if obj.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := AccessPolicyTargetToMap(&obj.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		result["source"] = sourceArray
	}

	return result
}

func AccessPolicySummaryToMap(obj oci_service_mesh.AccessPolicySummary) map[string]interface{} {
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

	if obj.MeshId != nil {
		result["mesh_id"] = string(*obj.MeshId)
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

func (s *ServiceMeshAccessPolicyResourceCrud) mapToAccessPolicyTargetDetails(fieldKeyFormat string) (oci_service_mesh.AccessPolicyTargetDetails, error) {
	var baseObject oci_service_mesh.AccessPolicyTargetDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ALL_VIRTUAL_SERVICES"):
		details := oci_service_mesh.AllVirtualServicesAccessPolicyTargetDetails{}
		baseObject = details
	case strings.ToLower("EXTERNAL_SERVICE"):
		details := oci_service_mesh.ExternalServiceAccessPolicyTargetDetails{}
		if hostnames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostnames")); ok {
			interfaces := hostnames.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hostnames")) {
				details.Hostnames = tmp
			}
		}
		if ipAddresses, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ip_addresses")); ok {
			interfaces := ipAddresses.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ip_addresses")) {
				details.IpAddresses = tmp
			}
		}
		if ports, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ports")); ok {
			interfaces := ports.([]interface{})
			tmp := make([]int, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(int)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ports")) {
				details.Ports = tmp
			}
		}
		if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
			details.Protocol = oci_service_mesh.ExternalServiceAccessPolicyTargetDetailsProtocolEnum(protocol.(string))
		}
		baseObject = details
	case strings.ToLower("INGRESS_GATEWAY"):
		details := oci_service_mesh.IngressGatewayAccessPolicyTargetDetails{}
		if ingressGatewayId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ingress_gateway_id")); ok {
			tmp := ingressGatewayId.(string)
			details.IngressGatewayId = &tmp
		}
		baseObject = details
	case strings.ToLower("VIRTUAL_SERVICE"):
		details := oci_service_mesh.VirtualServiceAccessPolicyTargetDetails{}
		if virtualServiceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "virtual_service_id")); ok {
			tmp := virtualServiceId.(string)
			details.VirtualServiceId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func AccessPolicyTargetToMap(obj *oci_service_mesh.AccessPolicyTarget) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_service_mesh.AllVirtualServicesAccessPolicyTarget:
		result["type"] = "ALL_VIRTUAL_SERVICES"
	case oci_service_mesh.ExternalServiceAccessPolicyTarget:
		result["type"] = "EXTERNAL_SERVICE"

		result["hostnames"] = v.Hostnames

		result["ip_addresses"] = v.IpAddresses

		result["ports"] = v.Ports

		result["protocol"] = string(v.Protocol)
	case oci_service_mesh.IngressGatewayAccessPolicyTarget:
		result["type"] = "INGRESS_GATEWAY"

		if v.IngressGatewayId != nil {
			result["ingress_gateway_id"] = string(*v.IngressGatewayId)
		}
	case oci_service_mesh.VirtualServiceAccessPolicyTarget:
		result["type"] = "VIRTUAL_SERVICE"

		if v.VirtualServiceId != nil {
			result["virtual_service_id"] = string(*v.VirtualServiceId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ServiceMeshAccessPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_service_mesh.ChangeAccessPolicyCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AccessPolicyId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.ChangeAccessPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAccessPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
