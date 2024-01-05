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

func ServiceMeshIngressGatewayResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createServiceMeshIngressGateway,
		Read:     readServiceMeshIngressGateway,
		Update:   updateServiceMeshIngressGateway,
		Delete:   deleteServiceMeshIngressGateway,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hosts": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"listeners": {
							Type:     schema.TypeList,
							Required: true,
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
									"tls": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"mode": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional
												"client_validation": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"subject_alternate_names": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"trusted_ca_bundle": {
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
																				"LOCAL_FILE",
																				"OCI_CERTIFICATES",
																			}, true),
																		},

																		// Optional
																		"ca_bundle_id": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"secret_name": {
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
												"server_certificate": {
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
																	"LOCAL_FILE",
																	"OCI_CERTIFICATES",
																}, true),
															},

															// Optional
															"certificate_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"secret_name": {
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

									// Computed
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
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

						// Computed
					},
				},
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
			"mtls": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"maximum_validity": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
						"certificate_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
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

func createServiceMeshIngressGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshIngressGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.CreateResource(d, sync)
}

func readServiceMeshIngressGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshIngressGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

func updateServiceMeshIngressGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshIngressGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteServiceMeshIngressGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshIngressGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ServiceMeshIngressGatewayResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_service_mesh.ServiceMeshClient
	Res                    *oci_service_mesh.IngressGateway
	DisableNotFoundRetries bool
}

func (s *ServiceMeshIngressGatewayResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ServiceMeshIngressGatewayResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_service_mesh.IngressGatewayLifecycleStateCreating),
	}
}

func (s *ServiceMeshIngressGatewayResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_service_mesh.IngressGatewayLifecycleStateActive),
	}
}

func (s *ServiceMeshIngressGatewayResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_service_mesh.IngressGatewayLifecycleStateDeleting),
	}
}

func (s *ServiceMeshIngressGatewayResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_service_mesh.IngressGatewayLifecycleStateDeleted),
	}
}

func (s *ServiceMeshIngressGatewayResourceCrud) Create() error {
	request := oci_service_mesh.CreateIngressGatewayRequest{}

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

	if hosts, ok := s.D.GetOkExists("hosts"); ok {
		interfaces := hosts.([]interface{})
		tmp := make([]oci_service_mesh.IngressGatewayHost, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "hosts", stateDataIndex)
			converted, err := s.mapToIngressGatewayHost(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("hosts") {
			request.Hosts = tmp
		}
	}

	if meshId, ok := s.D.GetOkExists("mesh_id"); ok {
		tmp := meshId.(string)
		request.MeshId = &tmp
	}

	if mtls, ok := s.D.GetOkExists("mtls"); ok {
		if tmpList := mtls.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mtls", 0)
			tmp, err := s.mapToCreateIngressGatewayMutualTransportLayerSecurityDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Mtls = &tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.CreateIngressGateway(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getIngressGatewayFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ServiceMeshIngressGatewayResourceCrud) getIngressGatewayFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_service_mesh.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	ingressGatewayId, err := ingressGatewayWaitForWorkRequest(workId, "meshingressgateway",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, ingressGatewayId)
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
	s.D.SetId(*ingressGatewayId)

	return s.Get()
}

func ingressGatewayWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func ingressGatewayWaitForWorkRequest(wId *string, entityType string, action oci_service_mesh.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_service_mesh.ServiceMeshClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "service_mesh")
	retryPolicy.ShouldRetryOperation = ingressGatewayWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromServiceMeshIngressGatewayWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromServiceMeshIngressGatewayWorkRequest(client *oci_service_mesh.ServiceMeshClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_service_mesh.ActionTypeEnum) error {
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

func (s *ServiceMeshIngressGatewayResourceCrud) Get() error {
	request := oci_service_mesh.GetIngressGatewayRequest{}

	tmp := s.D.Id()
	request.IngressGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.GetIngressGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IngressGateway
	return nil
}

func (s *ServiceMeshIngressGatewayResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_service_mesh.UpdateIngressGatewayRequest{}

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

	if hosts, ok := s.D.GetOkExists("hosts"); ok {
		interfaces := hosts.([]interface{})
		tmp := make([]oci_service_mesh.IngressGatewayHost, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "hosts", stateDataIndex)
			converted, err := s.mapToIngressGatewayHost(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("hosts") {
			request.Hosts = tmp
		}
	}

	tmp := s.D.Id()
	request.IngressGatewayId = &tmp

	if mtls, ok := s.D.GetOkExists("mtls"); ok {
		if tmpList := mtls.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mtls", 0)
			tmp, err := s.mapToCreateIngressGatewayMutualTransportLayerSecurityDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Mtls = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.UpdateIngressGateway(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIngressGatewayFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ServiceMeshIngressGatewayResourceCrud) Delete() error {
	request := oci_service_mesh.DeleteIngressGatewayRequest{}

	tmp := s.D.Id()
	request.IngressGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.DeleteIngressGateway(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := ingressGatewayWaitForWorkRequest(workId, "meshingressgateway",
		oci_service_mesh.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ServiceMeshIngressGatewayResourceCrud) SetData() error {
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

	hosts := []interface{}{}
	for _, item := range s.Res.Hosts {
		hosts = append(hosts, IngressGatewayHostToMap(item))
	}
	s.D.Set("hosts", hosts)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MeshId != nil {
		s.D.Set("mesh_id", *s.Res.MeshId)
	}

	if s.Res.Mtls != nil {
		s.D.Set("mtls", []interface{}{IngressGatewayMutualTransportLayerSecurityToMap(s.Res.Mtls)})
	} else {
		s.D.Set("mtls", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
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

func (s *ServiceMeshIngressGatewayResourceCrud) mapToAccessLoggingConfiguration(fieldKeyFormat string) (oci_service_mesh.AccessLoggingConfiguration, error) {
	result := oci_service_mesh.AccessLoggingConfiguration{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func (s *ServiceMeshIngressGatewayResourceCrud) mapToCaBundle(fieldKeyFormat string) (oci_service_mesh.CaBundle, error) {
	var baseObject oci_service_mesh.CaBundle
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("LOCAL_FILE"):
		details := oci_service_mesh.LocalFileCaBundle{}
		if secretName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_name")); ok {
			tmp := secretName.(string)
			details.SecretName = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_CERTIFICATES"):
		details := oci_service_mesh.OciCaBundle{}
		if caBundleId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ca_bundle_id")); ok {
			tmp := caBundleId.(string)
			details.CaBundleId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func CaBundleToMap(obj *oci_service_mesh.CaBundle) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_service_mesh.LocalFileCaBundle:
		result["type"] = "LOCAL_FILE"

		if v.SecretName != nil {
			result["secret_name"] = string(*v.SecretName)
		}
	case oci_service_mesh.OciCaBundle:
		result["type"] = "OCI_CERTIFICATES"

		if v.CaBundleId != nil {
			result["ca_bundle_id"] = string(*v.CaBundleId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ServiceMeshIngressGatewayResourceCrud) mapToCreateIngressGatewayMutualTransportLayerSecurityDetails(fieldKeyFormat string) (oci_service_mesh.IngressGatewayMutualTransportLayerSecurityDetails, error) {
	result := oci_service_mesh.IngressGatewayMutualTransportLayerSecurityDetails{}

	if maximumValidity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_validity")); ok {
		tmp := maximumValidity.(int)
		result.MaximumValidity = &tmp
	}

	return result, nil
}

func IngressGatewayMutualTransportLayerSecurityToMap(obj *oci_service_mesh.IngressGatewayMutualTransportLayerSecurity) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CertificateId != nil {
		result["certificate_id"] = string(*obj.CertificateId)
	}

	if obj.MaximumValidity != nil {
		result["maximum_validity"] = int(*obj.MaximumValidity)
	}

	return result
}

func (s *ServiceMeshIngressGatewayResourceCrud) mapToIngressGatewayHost(fieldKeyFormat string) (oci_service_mesh.IngressGatewayHost, error) {
	result := oci_service_mesh.IngressGatewayHost{}

	if hostnames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostnames")); ok {
		interfaces := hostnames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hostnames")) {
			result.Hostnames = tmp
		}
	}

	if listeners, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listeners")); ok {
		interfaces := listeners.([]interface{})
		tmp := make([]oci_service_mesh.IngressGatewayListener, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "listeners"), stateDataIndex)
			converted, err := s.mapToIngressGatewayListener(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "listeners")) {
			result.Listeners = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func IngressGatewayHostToMap(obj oci_service_mesh.IngressGatewayHost) map[string]interface{} {
	result := map[string]interface{}{}

	result["hostnames"] = obj.Hostnames

	listeners := []interface{}{}
	for _, item := range obj.Listeners {
		listeners = append(listeners, IngressGatewayListenerToMap(item))
	}
	result["listeners"] = listeners

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *ServiceMeshIngressGatewayResourceCrud) mapToIngressGatewayListener(fieldKeyFormat string) (oci_service_mesh.IngressGatewayListener, error) {
	result := oci_service_mesh.IngressGatewayListener{}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_service_mesh.IngressGatewayListenerProtocolEnum(protocol.(string))
	}

	if tls, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tls")); ok {
		if tmpList := tls.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tls"), 0)
			tmp, err := s.mapToIngressListenerTlsConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert tls, encountered error: %v", err)
			}
			result.Tls = &tmp
		}
	}

	return result, nil
}

func IngressGatewayListenerToMap(obj oci_service_mesh.IngressGatewayListener) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	result["protocol"] = string(obj.Protocol)

	if obj.Tls != nil {
		result["tls"] = []interface{}{IngressListenerTlsConfigToMap(obj.Tls)}
	}

	return result
}

func (s *ServiceMeshIngressGatewayResourceCrud) mapToIngressGatewayMutualTransportLayerSecurityDetails(fieldKeyFormat string) (oci_service_mesh.IngressGatewayMutualTransportLayerSecurityDetails, error) {
	result := oci_service_mesh.IngressGatewayMutualTransportLayerSecurityDetails{}

	if maximumValidity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_validity")); ok {
		tmp := maximumValidity.(int)
		result.MaximumValidity = &tmp
	}

	return result, nil
}

func IngressGatewaySummaryToMap(obj oci_service_mesh.IngressGatewaySummary) map[string]interface{} {
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

func (s *ServiceMeshIngressGatewayResourceCrud) mapToIngressListenerClientValidationConfig(fieldKeyFormat string) (oci_service_mesh.IngressListenerClientValidationConfig, error) {
	result := oci_service_mesh.IngressListenerClientValidationConfig{}

	if subjectAlternateNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subject_alternate_names")); ok {
		interfaces := subjectAlternateNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "subject_alternate_names")) {
			result.SubjectAlternateNames = tmp
		}
	}

	if trustedCaBundle, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trusted_ca_bundle")); ok {
		if tmpList := trustedCaBundle.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "trusted_ca_bundle"), 0)
			tmp, err := s.mapToCaBundle(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert trusted_ca_bundle, encountered error: %v", err)
			}
			result.TrustedCaBundle = tmp
		}
	}

	return result, nil
}

func IngressListenerClientValidationConfigToMap(obj *oci_service_mesh.IngressListenerClientValidationConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["subject_alternate_names"] = obj.SubjectAlternateNames

	if obj.TrustedCaBundle != nil {
		trustedCaBundleArray := []interface{}{}
		if trustedCaBundleMap := CaBundleToMap(&obj.TrustedCaBundle); trustedCaBundleMap != nil {
			trustedCaBundleArray = append(trustedCaBundleArray, trustedCaBundleMap)
		}
		result["trusted_ca_bundle"] = trustedCaBundleArray
	}

	return result
}

func (s *ServiceMeshIngressGatewayResourceCrud) mapToIngressListenerTlsConfig(fieldKeyFormat string) (oci_service_mesh.IngressListenerTlsConfig, error) {
	result := oci_service_mesh.IngressListenerTlsConfig{}

	if clientValidation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_validation")); ok {
		if tmpList := clientValidation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "client_validation"), 0)
			tmp, err := s.mapToIngressListenerClientValidationConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert client_validation, encountered error: %v", err)
			}
			result.ClientValidation = &tmp
		}
	}

	if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
		result.Mode = oci_service_mesh.IngressListenerTlsConfigModeEnum(mode.(string))
	}

	if serverCertificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "server_certificate")); ok {
		if tmpList := serverCertificate.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "server_certificate"), 0)
			tmp, err := s.mapToTlsCertificate(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert server_certificate, encountered error: %v", err)
			}
			result.ServerCertificate = tmp
		}
	}

	return result, nil
}

func IngressListenerTlsConfigToMap(obj *oci_service_mesh.IngressListenerTlsConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClientValidation != nil {
		result["client_validation"] = []interface{}{IngressListenerClientValidationConfigToMap(obj.ClientValidation)}
	}

	result["mode"] = string(obj.Mode)

	if obj.ServerCertificate != nil {
		serverCertificateArray := []interface{}{}
		if serverCertificateMap := TlsCertificateToMap(&obj.ServerCertificate); serverCertificateMap != nil {
			serverCertificateArray = append(serverCertificateArray, serverCertificateMap)
		}
		result["server_certificate"] = serverCertificateArray
	}

	return result
}

func (s *ServiceMeshIngressGatewayResourceCrud) mapToTlsCertificate(fieldKeyFormat string) (oci_service_mesh.TlsCertificate, error) {
	var baseObject oci_service_mesh.TlsCertificate
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("LOCAL_FILE"):
		details := oci_service_mesh.LocalFileTlsCertificate{}
		if secretName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_name")); ok {
			tmp := secretName.(string)
			details.SecretName = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_CERTIFICATES"):
		details := oci_service_mesh.OciTlsCertificate{}
		if certificateId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_id")); ok {
			tmp := certificateId.(string)
			details.CertificateId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func TlsCertificateToMap(obj *oci_service_mesh.TlsCertificate) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_service_mesh.LocalFileTlsCertificate:
		result["type"] = "LOCAL_FILE"

		if v.SecretName != nil {
			result["secret_name"] = string(*v.SecretName)
		}
	case oci_service_mesh.OciTlsCertificate:
		result["type"] = "OCI_CERTIFICATES"

		if v.CertificateId != nil {
			result["certificate_id"] = string(*v.CertificateId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ServiceMeshIngressGatewayResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_service_mesh.ChangeIngressGatewayCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.IngressGatewayId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh")

	response, err := s.Client.ChangeIngressGatewayCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIngressGatewayFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_mesh"), oci_service_mesh.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
