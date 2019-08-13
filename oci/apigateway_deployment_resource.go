// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_apigateway "github.com/oracle/oci-go-sdk/apigateway"
	oci_common "github.com/oracle/oci-go-sdk/common"
)

func ApigatewayDeploymentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createApigatewayDeployment,
		Read:     readApigatewayDeployment,
		Update:   updateApigatewayDeployment,
		Delete:   deleteApigatewayDeployment,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"gateway_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"path_prefix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"specification": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"logging_policies": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"access_log": {
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
									"execution_log": {
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
												"log_level": {
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
						"request_policies": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"authentication": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"function_id": {
													Type:     schema.TypeString,
													Required: true,
												},
												"type": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"CUSTOM_AUTHENTICATION",
													}, true),
												},

												// Optional
												"is_anonymous_access_allowed": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"token_header": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"token_query_param": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"cors": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"allowed_origins": {
													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												// Optional
												"allowed_headers": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_methods": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"exposed_headers": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"is_allow_credentials_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"max_age_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"rate_limiting": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"rate_in_requests_per_second": {
													Type:     schema.TypeInt,
													Required: true,
												},
												"rate_key": {
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
						"routes": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"backend": {
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
													DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"HTTP_BACKEND",
														"ORACLE_FUNCTIONS_BACKEND",
														"STOCK_RESPONSE_BACKEND",
													}, true),
												},

												// Optional
												"body": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"connect_timeout_in_seconds": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"function_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"headers": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

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
														},
													},
												},
												"is_ssl_verify_disabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"read_timeout_in_seconds": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"send_timeout_in_seconds": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"status": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"url": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"path": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"logging_policies": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"access_log": {
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
												"execution_log": {
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
															"log_level": {
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
									"methods": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"request_policies": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"authorization": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"allowed_scope": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"type": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"ANONYMOUS",
																	"ANY_OF",
																	"AUTHENTICATION_ONLY",
																}, true),
															},

															// Computed
														},
													},
												},
												"cors": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"allowed_origins": {
																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															// Optional
															"allowed_headers": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"allowed_methods": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"exposed_headers": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"is_allow_credentials_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"max_age_in_seconds": {
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

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
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

			// Computed
			"endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
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

func createApigatewayDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).deploymentClient
	sync.WorkRequestsClient = m.(*OracleClients).gatewayWorkRequestsClient

	return CreateResource(d, sync)
}

func readApigatewayDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).deploymentClient
	sync.WorkRequestsClient = m.(*OracleClients).gatewayWorkRequestsClient

	return ReadResource(sync)
}

func updateApigatewayDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).deploymentClient
	sync.WorkRequestsClient = m.(*OracleClients).gatewayWorkRequestsClient

	return UpdateResource(d, sync)
}

func deleteApigatewayDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).deploymentClient
	sync.WorkRequestsClient = m.(*OracleClients).gatewayWorkRequestsClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type ApigatewayDeploymentResourceCrud struct {
	BaseCrud
	Client                 *oci_apigateway.DeploymentClient
	WorkRequestsClient     *oci_apigateway.WorkRequestsClient
	Res                    *oci_apigateway.Deployment
	DisableNotFoundRetries bool
}

func (s *ApigatewayDeploymentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ApigatewayDeploymentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_apigateway.DeploymentLifecycleStateCreating),
	}
}

func (s *ApigatewayDeploymentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_apigateway.DeploymentLifecycleStateActive),
	}
}

func (s *ApigatewayDeploymentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_apigateway.DeploymentLifecycleStateDeleting),
	}
}

func (s *ApigatewayDeploymentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_apigateway.DeploymentLifecycleStateDeleted),
	}
}

func (s *ApigatewayDeploymentResourceCrud) Create() error {
	request := oci_apigateway.CreateDeploymentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gatewayId, ok := s.D.GetOkExists("gateway_id"); ok {
		tmp := gatewayId.(string)
		request.GatewayId = &tmp
	}

	if pathPrefix, ok := s.D.GetOkExists("path_prefix"); ok {
		tmp := pathPrefix.(string)
		request.PathPrefix = &tmp
	}

	if specification, ok := s.D.GetOkExists("specification"); ok {
		if tmpList := specification.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "specification", 0)
			tmp, err := s.mapToApiSpecification(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Specification = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.CreateDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeploymentFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ApigatewayDeploymentResourceCrud) getDeploymentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_apigateway.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	deploymentId, err := deploymentWaitForWorkRequest(workId, "deployment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestsClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, deploymentId)
		_, cancelErr := s.WorkRequestsClient.CancelWorkRequest(context.Background(),
			oci_apigateway.CancelWorkRequestRequest{
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
	s.D.SetId(*deploymentId)

	return s.Get()
}

func deploymentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "apigateway", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_apigateway.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func deploymentWaitForWorkRequest(wId *string, entityType string, action oci_apigateway.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_apigateway.WorkRequestsClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "apigateway")
	retryPolicy.ShouldRetryOperation = deploymentWorkRequestShouldRetryFunc(timeout)

	response := oci_apigateway.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_apigateway.WorkRequestStatusInProgress),
			string(oci_apigateway.WorkRequestStatusAccepted),
			string(oci_apigateway.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_apigateway.WorkRequestStatusSucceeded),
			string(oci_apigateway.WorkRequestStatusFailed),
			string(oci_apigateway.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_apigateway.GetWorkRequestRequest{
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

	// The API Gateway workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_apigateway.WorkRequestStatusFailed || response.Status == oci_apigateway.WorkRequestStatusCanceled {
		return nil, getErrorFromGatewayWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func (s *ApigatewayDeploymentResourceCrud) Get() error {
	request := oci_apigateway.GetDeploymentRequest{}

	tmp := s.D.Id()
	request.DeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.GetDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Deployment
	return nil
}

func (s *ApigatewayDeploymentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_apigateway.UpdateDeploymentRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.DeploymentId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if specification, ok := s.D.GetOkExists("specification"); ok {
		if tmpList := specification.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "specification", 0)
			tmp, err := s.mapToApiSpecification(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Specification = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.UpdateDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeploymentFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ApigatewayDeploymentResourceCrud) Delete() error {
	request := oci_apigateway.DeleteDeploymentRequest{}

	tmp := s.D.Id()
	request.DeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.DeleteDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := deploymentWaitForWorkRequest(workId, "deployment",
		oci_apigateway.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestsClient)
	return delWorkRequestErr
}

func (s *ApigatewayDeploymentResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Endpoint != nil {
		s.D.Set("endpoint", *s.Res.Endpoint)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GatewayId != nil {
		s.D.Set("gateway_id", *s.Res.GatewayId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PathPrefix != nil {
		s.D.Set("path_prefix", *s.Res.PathPrefix)
	}

	if s.Res.Specification != nil {
		s.D.Set("specification", []interface{}{ApiSpecificationToMap(s.Res.Specification)})
	} else {
		s.D.Set("specification", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *ApigatewayDeploymentResourceCrud) mapToAccessLogPolicy(fieldKeyFormat string) (oci_apigateway.AccessLogPolicy, error) {
	result := oci_apigateway.AccessLogPolicy{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func AccessLogPolicyToMap(obj *oci_apigateway.AccessLogPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToApiSpecification(fieldKeyFormat string) (oci_apigateway.ApiSpecification, error) {
	result := oci_apigateway.ApiSpecification{}

	if loggingPolicies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logging_policies")); ok {
		if tmpList := loggingPolicies.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "logging_policies"), 0)
			tmp, err := s.mapToApiSpecificationLoggingPolicies(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert logging_policies, encountered error: %v", err)
			}
			result.LoggingPolicies = &tmp
		}
	}

	if requestPolicies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_policies")); ok {
		if tmpList := requestPolicies.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "request_policies"), 0)
			tmp, err := s.mapToApiSpecificationRequestPolicies(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert request_policies, encountered error: %v", err)
			}
			result.RequestPolicies = &tmp
		}
	}

	if routes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "routes")); ok {
		interfaces := routes.([]interface{})
		tmp := make([]oci_apigateway.ApiSpecificationRoute, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "routes"), stateDataIndex)
			converted, err := s.mapToApiSpecificationRoute(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "routes")) {
			result.Routes = tmp
		}
	}

	return result, nil
}

func ApiSpecificationToMap(obj *oci_apigateway.ApiSpecification) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LoggingPolicies != nil {
		result["logging_policies"] = []interface{}{ApiSpecificationLoggingPoliciesToMap(obj.LoggingPolicies)}
	}

	if obj.RequestPolicies != nil {
		result["request_policies"] = []interface{}{ApiSpecificationRequestPoliciesToMap(obj.RequestPolicies)}
	}

	routes := []interface{}{}
	for _, item := range obj.Routes {
		routes = append(routes, ApiSpecificationRouteToMap(item))
	}
	result["routes"] = routes

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToApiSpecificationLoggingPolicies(fieldKeyFormat string) (oci_apigateway.ApiSpecificationLoggingPolicies, error) {
	result := oci_apigateway.ApiSpecificationLoggingPolicies{}

	if accessLog, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_log")); ok {
		if tmpList := accessLog.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "access_log"), 0)
			tmp, err := s.mapToAccessLogPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert access_log, encountered error: %v", err)
			}
			result.AccessLog = &tmp
		}
	}

	if executionLog, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "execution_log")); ok {
		if tmpList := executionLog.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "execution_log"), 0)
			tmp, err := s.mapToExecutionLogPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert execution_log, encountered error: %v", err)
			}
			result.ExecutionLog = &tmp
		}
	}

	return result, nil
}

func ApiSpecificationLoggingPoliciesToMap(obj *oci_apigateway.ApiSpecificationLoggingPolicies) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccessLog != nil {
		result["access_log"] = []interface{}{AccessLogPolicyToMap(obj.AccessLog)}
	}

	if obj.ExecutionLog != nil {
		result["execution_log"] = []interface{}{ExecutionLogPolicyToMap(obj.ExecutionLog)}
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToApiSpecificationRequestPolicies(fieldKeyFormat string) (oci_apigateway.ApiSpecificationRequestPolicies, error) {
	result := oci_apigateway.ApiSpecificationRequestPolicies{}

	if authentication, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authentication")); ok {
		if tmpList := authentication.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "authentication"), 0)
			tmp, err := s.mapToAuthenticationPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert authentication, encountered error: %v", err)
			}
			result.Authentication = tmp
		}
	}

	if cors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cors")); ok {
		if tmpList := cors.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "cors"), 0)
			tmp, err := s.mapToCorsPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert cors, encountered error: %v", err)
			}
			result.Cors = &tmp
		}
	}

	if rateLimiting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rate_limiting")); ok {
		if tmpList := rateLimiting.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rate_limiting"), 0)
			tmp, err := s.mapToRateLimitingPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert rate_limiting, encountered error: %v", err)
			}
			result.RateLimiting = &tmp
		}
	}

	return result, nil
}

func ApiSpecificationRequestPoliciesToMap(obj *oci_apigateway.ApiSpecificationRequestPolicies) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Authentication != nil {
		authenticationArray := []interface{}{}
		if authenticationMap := AuthenticationPolicyToMap(&obj.Authentication); authenticationMap != nil {
			authenticationArray = append(authenticationArray, authenticationMap)
		}
		result["authentication"] = authenticationArray
	}

	if obj.Cors != nil {
		result["cors"] = []interface{}{CorsPolicyToMap(obj.Cors)}
	}

	if obj.RateLimiting != nil {
		result["rate_limiting"] = []interface{}{RateLimitingPolicyToMap(obj.RateLimiting)}
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToApiSpecificationRoute(fieldKeyFormat string) (oci_apigateway.ApiSpecificationRoute, error) {
	result := oci_apigateway.ApiSpecificationRoute{}

	if backend, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backend")); ok {
		if tmpList := backend.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backend"), 0)
			tmp, err := s.mapToApiSpecificationRouteBackend(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert backend, encountered error: %v", err)
			}
			result.Backend = tmp
		}
	}

	if loggingPolicies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logging_policies")); ok {
		if tmpList := loggingPolicies.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "logging_policies"), 0)
			tmp, err := s.mapToApiSpecificationLoggingPolicies(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert logging_policies, encountered error: %v", err)
			}
			result.LoggingPolicies = &tmp
		}
	}

	if methods, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "methods")); ok {
		interfaces := methods.([]interface{})
		tmp := make([]oci_apigateway.ApiSpecificationRouteMethodsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_apigateway.ApiSpecificationRouteMethodsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "methods")) {
			result.Methods = tmp
		}
	}

	if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
		tmp := path.(string)
		result.Path = &tmp
	}

	if requestPolicies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_policies")); ok {
		if tmpList := requestPolicies.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "request_policies"), 0)
			tmp, err := s.mapToApiSpecificationRouteRequestPolicies(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert request_policies, encountered error: %v", err)
			}
			result.RequestPolicies = &tmp
		}
	}

	return result, nil
}

func ApiSpecificationRouteToMap(obj oci_apigateway.ApiSpecificationRoute) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Backend != nil {
		backendArray := []interface{}{}
		if backendMap := ApiSpecificationRouteBackendToMap(&obj.Backend); backendMap != nil {
			backendArray = append(backendArray, backendMap)
		}
		result["backend"] = backendArray
	}

	if obj.LoggingPolicies != nil {
		result["logging_policies"] = []interface{}{ApiSpecificationLoggingPoliciesToMap(obj.LoggingPolicies)}
	}

	result["methods"] = obj.Methods

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	if obj.RequestPolicies != nil {
		result["request_policies"] = []interface{}{ApiSpecificationRouteRequestPoliciesToMap(obj.RequestPolicies)}
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToApiSpecificationRouteBackend(fieldKeyFormat string) (oci_apigateway.ApiSpecificationRouteBackend, error) {
	var baseObject oci_apigateway.ApiSpecificationRouteBackend
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("HTTP_BACKEND"):
		details := oci_apigateway.HttpBackend{}
		if connectTimeoutInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connect_timeout_in_seconds")); ok {
			tmp := float32(connectTimeoutInSeconds.(float64))
			details.ConnectTimeoutInSeconds = &tmp
		}
		if isSslVerifyDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ssl_verify_disabled")); ok {
			tmp := isSslVerifyDisabled.(bool)
			details.IsSslVerifyDisabled = &tmp
		}
		if readTimeoutInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "read_timeout_in_seconds")); ok {
			tmp := float32(readTimeoutInSeconds.(float64))
			details.ReadTimeoutInSeconds = &tmp
		}
		if sendTimeoutInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "send_timeout_in_seconds")); ok {
			tmp := float32(sendTimeoutInSeconds.(float64))
			details.SendTimeoutInSeconds = &tmp
		}
		if url, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url")); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		baseObject = details
	case strings.ToLower("ORACLE_FUNCTIONS_BACKEND"):
		details := oci_apigateway.OracleFunctionBackend{}
		if functionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "function_id")); ok {
			tmp := functionId.(string)
			details.FunctionId = &tmp
		}
		baseObject = details
	case strings.ToLower("STOCK_RESPONSE_BACKEND"):
		details := oci_apigateway.StockResponseBackend{}
		if body, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "body")); ok {
			tmp := body.(string)
			details.Body = &tmp
		}
		if headers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "headers")); ok {
			interfaces := headers.([]interface{})
			tmp := make([]oci_apigateway.HeaderFieldSpecification, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "headers"), stateDataIndex)
				converted, err := s.mapToHeaderFieldSpecification(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "headers")) {
				details.Headers = tmp
			}
		}
		if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
			tmp := status.(int)
			details.Status = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ApiSpecificationRouteBackendToMap(obj *oci_apigateway.ApiSpecificationRouteBackend) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_apigateway.HttpBackend:
		result["type"] = "HTTP_BACKEND"

		if v.ConnectTimeoutInSeconds != nil {
			result["connect_timeout_in_seconds"] = float32(*v.ConnectTimeoutInSeconds)
		}

		if v.IsSslVerifyDisabled != nil {
			result["is_ssl_verify_disabled"] = bool(*v.IsSslVerifyDisabled)
		}

		if v.ReadTimeoutInSeconds != nil {
			result["read_timeout_in_seconds"] = float32(*v.ReadTimeoutInSeconds)
		}

		if v.SendTimeoutInSeconds != nil {
			result["send_timeout_in_seconds"] = float32(*v.SendTimeoutInSeconds)
		}

		if v.Url != nil {
			result["url"] = string(*v.Url)
		}
	case oci_apigateway.OracleFunctionBackend:
		result["type"] = "ORACLE_FUNCTIONS_BACKEND"

		if v.FunctionId != nil {
			result["function_id"] = string(*v.FunctionId)
		}
	case oci_apigateway.StockResponseBackend:
		result["type"] = "STOCK_RESPONSE_BACKEND"

		if v.Body != nil {
			result["body"] = string(*v.Body)
		}

		headers := []interface{}{}
		for _, item := range v.Headers {
			headers = append(headers, HeaderFieldSpecificationToMap(item))
		}
		result["headers"] = headers

		if v.Status != nil {
			result["status"] = int(*v.Status)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToApiSpecificationRouteRequestPolicies(fieldKeyFormat string) (oci_apigateway.ApiSpecificationRouteRequestPolicies, error) {
	result := oci_apigateway.ApiSpecificationRouteRequestPolicies{}

	if authorization, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authorization")); ok {
		if tmpList := authorization.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "authorization"), 0)
			tmp, err := s.mapToRouteAuthorizationPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert authorization, encountered error: %v", err)
			}
			result.Authorization = tmp
		}
	}

	if cors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cors")); ok {
		if tmpList := cors.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "cors"), 0)
			tmp, err := s.mapToCorsPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert cors, encountered error: %v", err)
			}
			result.Cors = &tmp
		}
	}

	return result, nil
}

func ApiSpecificationRouteRequestPoliciesToMap(obj *oci_apigateway.ApiSpecificationRouteRequestPolicies) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Authorization != nil {
		authorizationArray := []interface{}{}
		if authorizationMap := RouteAuthorizationPolicyToMap(&obj.Authorization); authorizationMap != nil {
			authorizationArray = append(authorizationArray, authorizationMap)
		}
		result["authorization"] = authorizationArray
	}

	if obj.Cors != nil {
		result["cors"] = []interface{}{CorsPolicyToMap(obj.Cors)}
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToAuthenticationPolicy(fieldKeyFormat string) (oci_apigateway.AuthenticationPolicy, error) {
	var baseObject oci_apigateway.AuthenticationPolicy
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("CUSTOM_AUTHENTICATION"):
		details := oci_apigateway.CustomAuthenticationPolicy{}
		if functionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "function_id")); ok {
			tmp := functionId.(string)
			details.FunctionId = &tmp
		}
		if tokenHeader, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "token_header")); ok {
			tmp := tokenHeader.(string)
			if len(tmp) > 0 {
				details.TokenHeader = &tmp
			}
		}
		if tokenQueryParam, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "token_query_param")); ok {
			tmp := tokenQueryParam.(string)
			if len(tmp) > 0 {
				details.TokenQueryParam = &tmp
			}
		}
		if isAnonymousAccessAllowed, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_anonymous_access_allowed")); ok {
			tmp := isAnonymousAccessAllowed.(bool)
			details.IsAnonymousAccessAllowed = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func AuthenticationPolicyToMap(obj *oci_apigateway.AuthenticationPolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_apigateway.CustomAuthenticationPolicy:
		result["type"] = "CUSTOM_AUTHENTICATION"

		if v.FunctionId != nil {
			result["function_id"] = string(*v.FunctionId)
		}

		if v.TokenHeader != nil {
			result["token_header"] = string(*v.TokenHeader)
		}

		if v.TokenQueryParam != nil {
			result["token_query_param"] = string(*v.TokenQueryParam)
		}

		if v.IsAnonymousAccessAllowed != nil {
			result["is_anonymous_access_allowed"] = bool(*v.IsAnonymousAccessAllowed)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToCorsPolicy(fieldKeyFormat string) (oci_apigateway.CorsPolicy, error) {
	result := oci_apigateway.CorsPolicy{}

	if allowedHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_headers")); ok {
		interfaces := allowedHeaders.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_headers")) {
			result.AllowedHeaders = tmp
		}
	}

	if allowedMethods, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_methods")); ok {
		interfaces := allowedMethods.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_methods")) {
			result.AllowedMethods = tmp
		}
	}

	if allowedOrigins, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_origins")); ok {
		interfaces := allowedOrigins.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_origins")) {
			result.AllowedOrigins = tmp
		}
	}

	if exposedHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exposed_headers")); ok {
		interfaces := exposedHeaders.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exposed_headers")) {
			result.ExposedHeaders = tmp
		}
	}

	if isAllowCredentialsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_allow_credentials_enabled")); ok {
		tmp := isAllowCredentialsEnabled.(bool)
		result.IsAllowCredentialsEnabled = &tmp
	}

	if maxAgeInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_age_in_seconds")); ok {
		tmp := maxAgeInSeconds.(int)
		result.MaxAgeInSeconds = &tmp
	}

	return result, nil
}

func CorsPolicyToMap(obj *oci_apigateway.CorsPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_headers"] = obj.AllowedHeaders

	result["allowed_methods"] = obj.AllowedMethods

	result["allowed_origins"] = obj.AllowedOrigins

	result["exposed_headers"] = obj.ExposedHeaders

	if obj.IsAllowCredentialsEnabled != nil {
		result["is_allow_credentials_enabled"] = bool(*obj.IsAllowCredentialsEnabled)
	}

	if obj.MaxAgeInSeconds != nil {
		result["max_age_in_seconds"] = int(*obj.MaxAgeInSeconds)
	}

	return result
}

func DeploymentSummaryToMap(obj oci_apigateway.DeploymentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Endpoint != nil {
		result["endpoint"] = string(*obj.Endpoint)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.GatewayId != nil {
		result["gateway_id"] = string(*obj.GatewayId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.PathPrefix != nil {
		result["path_prefix"] = string(*obj.PathPrefix)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToExecutionLogPolicy(fieldKeyFormat string) (oci_apigateway.ExecutionLogPolicy, error) {
	result := oci_apigateway.ExecutionLogPolicy{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if logLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_level")); ok {
		result.LogLevel = oci_apigateway.ExecutionLogPolicyLogLevelEnum(logLevel.(string))
	}

	return result, nil
}

func ExecutionLogPolicyToMap(obj *oci_apigateway.ExecutionLogPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	result["log_level"] = string(obj.LogLevel)

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToHeaderFieldSpecification(fieldKeyFormat string) (oci_apigateway.HeaderFieldSpecification, error) {
	result := oci_apigateway.HeaderFieldSpecification{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func HeaderFieldSpecificationToMap(obj oci_apigateway.HeaderFieldSpecification) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToRateLimitingPolicy(fieldKeyFormat string) (oci_apigateway.RateLimitingPolicy, error) {
	result := oci_apigateway.RateLimitingPolicy{}

	if rateInRequestsPerSecond, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rate_in_requests_per_second")); ok {
		tmp := rateInRequestsPerSecond.(int)
		result.RateInRequestsPerSecond = &tmp
	}

	if rateKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rate_key")); ok {
		result.RateKey = oci_apigateway.RateLimitingPolicyRateKeyEnum(rateKey.(string))
	}

	return result, nil
}

func RateLimitingPolicyToMap(obj *oci_apigateway.RateLimitingPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RateInRequestsPerSecond != nil {
		result["rate_in_requests_per_second"] = int(*obj.RateInRequestsPerSecond)
	}

	result["rate_key"] = string(obj.RateKey)

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToRouteAuthorizationPolicy(fieldKeyFormat string) (oci_apigateway.RouteAuthorizationPolicy, error) {
	var baseObject oci_apigateway.RouteAuthorizationPolicy
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "AUTHENTICATION_ONLY" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ANONYMOUS"):
		details := oci_apigateway.AnonymousRouteAuthorizationPolicy{}
		baseObject = details
	case strings.ToLower("ANY_OF"):
		details := oci_apigateway.AnyOfRouteAuthorizationPolicy{}
		if allowedScope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_scope")); ok {
			interfaces := allowedScope.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_scope")) {
				details.AllowedScope = tmp
			}
		}
		baseObject = details
	case strings.ToLower("AUTHENTICATION_ONLY"):
		details := oci_apigateway.AuthenticationOnlyRouteAuthorizationPolicy{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func RouteAuthorizationPolicyToMap(obj *oci_apigateway.RouteAuthorizationPolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_apigateway.AnonymousRouteAuthorizationPolicy:
		result["type"] = "ANONYMOUS"
	case oci_apigateway.AnyOfRouteAuthorizationPolicy:
		result["type"] = "ANY_OF"

		result["allowed_scope"] = v.AllowedScope
	case oci_apigateway.AuthenticationOnlyRouteAuthorizationPolicy:
		result["type"] = "AUTHENTICATION_ONLY"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_apigateway.ChangeDeploymentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DeploymentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	_, err := s.Client.ChangeDeploymentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
