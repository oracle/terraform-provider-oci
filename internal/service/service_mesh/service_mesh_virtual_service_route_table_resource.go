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

func ServiceMeshVirtualServiceRouteTableResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createServiceMeshVirtualServiceRouteTable,
		Read:     readServiceMeshVirtualServiceRouteTable,
		Update:   updateServiceMeshVirtualServiceRouteTable,
		Delete:   deleteServiceMeshVirtualServiceRouteTable,
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
			"route_rules": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      routeRulesHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"destinations": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"virtual_deployment_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"weight": {
										Type:     schema.TypeInt,
										Required: true,
									},

									// Optional
									"port": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"HTTP",
								"TCP",
								"TLS_PASSTHROUGH",
							}, true),
						},

						// Optional
						"is_grpc": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"path_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"virtual_service_id": {
				Type:     schema.TypeString,
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
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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

func createServiceMeshVirtualServiceRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualServiceRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.CreateResource(d, sync)
}

func readServiceMeshVirtualServiceRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualServiceRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

func updateServiceMeshVirtualServiceRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualServiceRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteServiceMeshVirtualServiceRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualServiceRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ServiceMeshVirtualServiceRouteTableResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_service_mesh.ServiceMeshClient
	Res                    *oci_service_mesh.VirtualServiceRouteTable
	DisableNotFoundRetries bool
}

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_service_mesh.VirtualServiceRouteTableLifecycleStateCreating),
	}
}

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_service_mesh.VirtualServiceRouteTableLifecycleStateActive),
	}
}

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_service_mesh.VirtualServiceRouteTableLifecycleStateDeleting),
	}
}

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_service_mesh.VirtualServiceRouteTableLifecycleStateDeleted),
	}
}

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) Create() error {
	request := oci_service_mesh.CreateVirtualServiceRouteTableRequest{}

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

	if priority, ok := s.D.GetOkExists("priority"); ok {
		tmp := priority.(int)
		request.Priority = &tmp
	}

	if routeRules, ok := s.D.GetOkExists("route_rules"); ok {
		set := routeRules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_service_mesh.VirtualServiceTrafficRouteRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := routeRulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "route_rules", stateDataIndex)
			converted, err := s.mapToVirtualServiceTrafficRouteRuleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("route_rules") {
			request.RouteRules = tmp
		}
	}

	if virtualServiceId, ok := s.D.GetOkExists("virtual_service_id"); ok {
		tmp := virtualServiceId.(string)
		request.VirtualServiceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.CreateVirtualServiceRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getVirtualServiceRouteTableFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) getVirtualServiceRouteTableFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_service_mesh.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	virtualServiceRouteTableId, err := virtualServiceRouteTableWaitForWorkRequest(workId, "virtualserviceroutetable",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, virtualServiceRouteTableId)
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
	s.D.SetId(*virtualServiceRouteTableId)

	return s.Get()
}

func virtualServiceRouteTableWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func virtualServiceRouteTableWaitForWorkRequest(wId *string, entityType string, action oci_service_mesh.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_service_mesh.ServiceMeshClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "service_mesh")
	retryPolicy.ShouldRetryOperation = virtualServiceRouteTableWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromServiceMeshVirtualServiceRouteTableWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromServiceMeshVirtualServiceRouteTableWorkRequest(client *oci_service_mesh.ServiceMeshClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_service_mesh.ActionTypeEnum) error {
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

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) Get() error {
	request := oci_service_mesh.GetVirtualServiceRouteTableRequest{}

	tmp := s.D.Id()
	request.VirtualServiceRouteTableId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.GetVirtualServiceRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VirtualServiceRouteTable
	return nil
}

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_service_mesh.UpdateVirtualServiceRouteTableRequest{}

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

	if priority, ok := s.D.GetOkExists("priority"); ok {
		tmp := priority.(int)
		request.Priority = &tmp
	}

	if routeRules, ok := s.D.GetOkExists("route_rules"); ok {
		set := routeRules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_service_mesh.VirtualServiceTrafficRouteRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := routeRulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "route_rules", stateDataIndex)
			converted, err := s.mapToVirtualServiceTrafficRouteRuleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("route_rules") {
			request.RouteRules = tmp
		}
	}

	tmp := s.D.Id()
	request.VirtualServiceRouteTableId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.UpdateVirtualServiceRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVirtualServiceRouteTableFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) Delete() error {
	request := oci_service_mesh.DeleteVirtualServiceRouteTableRequest{}

	tmp := s.D.Id()
	request.VirtualServiceRouteTableId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.DeleteVirtualServiceRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := virtualServiceRouteTableWaitForWorkRequest(workId, "virtualserviceroutetable",
		oci_service_mesh.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) SetData() error {
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

	if s.Res.Priority != nil {
		s.D.Set("priority", *s.Res.Priority)
	}

	routeRules := []interface{}{}
	for _, item := range s.Res.RouteRules {
		routeRules = append(routeRules, VirtualServiceTrafficRouteRuleToMap(item))
	}
	s.D.Set("route_rules", schema.NewSet(routeRulesHashCodeForSets, routeRules))

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

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) mapToVirtualDeploymentTrafficRuleTargetDetails(fieldKeyFormat string) (oci_service_mesh.VirtualDeploymentTrafficRuleTargetDetails, error) {
	result := oci_service_mesh.VirtualDeploymentTrafficRuleTargetDetails{}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if virtualDeploymentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "virtual_deployment_id")); ok {
		tmp := virtualDeploymentId.(string)
		result.VirtualDeploymentId = &tmp
	}

	if weight, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weight")); ok {
		tmp := weight.(int)
		result.Weight = &tmp
	}

	return result, nil
}

func VirtualDeploymentTrafficRuleTargetToMap(obj oci_service_mesh.VirtualDeploymentTrafficRuleTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.VirtualDeploymentId != nil {
		result["virtual_deployment_id"] = string(*obj.VirtualDeploymentId)
	}

	if obj.Weight != nil {
		result["weight"] = int(*obj.Weight)
	}

	return result
}

func VirtualServiceRouteTableSummaryToMap(obj oci_service_mesh.VirtualServiceRouteTableSummary) map[string]interface{} {
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

	if obj.Priority != nil {
		result["priority"] = int(*obj.Priority)
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

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) mapToVirtualServiceTrafficRouteRuleDetails(fieldKeyFormat string) (oci_service_mesh.VirtualServiceTrafficRouteRuleDetails, error) {
	var baseObject oci_service_mesh.VirtualServiceTrafficRouteRuleDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("HTTP"):
		details := oci_service_mesh.HttpVirtualServiceTrafficRouteRuleDetails{}
		if isGrpc, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_grpc")); ok {
			tmp := isGrpc.(bool)
			details.IsGrpc = &tmp
		}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if pathType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path_type")); ok {
			details.PathType = oci_service_mesh.HttpVirtualServiceTrafficRouteRuleDetailsPathTypeEnum(pathType.(string))
		}
		if requestTimeoutInMs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_timeout_in_ms")); ok {
			tmp := requestTimeoutInMs.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert requestTimeoutInMs string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.RequestTimeoutInMs = &tmpInt64
		}
		if destinations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destinations")); ok {
			interfaces := destinations.([]interface{})
			tmp := make([]oci_service_mesh.VirtualDeploymentTrafficRuleTargetDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destinations"), stateDataIndex)
				converted, err := s.mapToVirtualDeploymentTrafficRuleTargetDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "destinations")) {
				details.Destinations = tmp
			}
		}
		baseObject = details
	case strings.ToLower("TCP"):
		details := oci_service_mesh.TcpVirtualServiceTrafficRouteRuleDetails{}
		if destinations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destinations")); ok {
			interfaces := destinations.([]interface{})
			tmp := make([]oci_service_mesh.VirtualDeploymentTrafficRuleTargetDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destinations"), stateDataIndex)
				converted, err := s.mapToVirtualDeploymentTrafficRuleTargetDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "destinations")) {
				details.Destinations = tmp
			}
		}
		baseObject = details
	case strings.ToLower("TLS_PASSTHROUGH"):
		details := oci_service_mesh.TlsPassthroughVirtualServiceTrafficRouteRuleDetails{}
		if destinations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destinations")); ok {
			interfaces := destinations.([]interface{})
			tmp := make([]oci_service_mesh.VirtualDeploymentTrafficRuleTargetDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destinations"), stateDataIndex)
				converted, err := s.mapToVirtualDeploymentTrafficRuleTargetDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "destinations")) {
				details.Destinations = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func VirtualServiceTrafficRouteRuleToMap(obj oci_service_mesh.VirtualServiceTrafficRouteRule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_service_mesh.HttpVirtualServiceTrafficRouteRule:
		result["type"] = "HTTP"

		if v.IsGrpc != nil {
			result["is_grpc"] = bool(*v.IsGrpc)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		result["path_type"] = string(v.PathType)

		if v.RequestTimeoutInMs != nil {
			result["request_timeout_in_ms"] = strconv.FormatInt(*v.RequestTimeoutInMs, 10)
		}

		destinations := []interface{}{}
		for _, item := range v.Destinations {
			destinations = append(destinations, VirtualDeploymentTrafficRuleTargetToMap(item))
		}
		result["destinations"] = destinations
	case oci_service_mesh.TcpVirtualServiceTrafficRouteRule:
		result["type"] = "TCP"

		destinations := []interface{}{}
		for _, item := range v.Destinations {
			destinations = append(destinations, VirtualDeploymentTrafficRuleTargetToMap(item))
		}
		result["destinations"] = destinations
	case oci_service_mesh.TlsPassthroughVirtualServiceTrafficRouteRule:
		result["type"] = "TLS_PASSTHROUGH"

		destinations := []interface{}{}
		for _, item := range v.Destinations {
			destinations = append(destinations, VirtualDeploymentTrafficRuleTargetToMap(item))
		}
		result["destinations"] = destinations
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *ServiceMeshVirtualServiceRouteTableResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_service_mesh.ChangeVirtualServiceRouteTableCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VirtualServiceRouteTableId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.ChangeVirtualServiceRouteTableCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVirtualServiceRouteTableFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
