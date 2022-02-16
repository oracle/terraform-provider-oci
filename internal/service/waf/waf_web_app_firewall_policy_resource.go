// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_waf "github.com/oracle/oci-go-sdk/v58/waf"
)

func WafWebAppFirewallPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createWafWebAppFirewallPolicy,
		Read:     readWafWebAppFirewallPolicy,
		Update:   updateWafWebAppFirewallPolicy,
		Delete:   deleteWafWebAppFirewallPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"actions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ALLOW",
								"CHECK",
								"RETURN_HTTP_RESPONSE",
							}, true),
						},

						// Optional
						"body": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"text": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"STATIC_TEXT",
										}, true),
									},

									// Optional

									// Computed
								},
							},
						},
						"code": {
							Type:     schema.TypeInt,
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
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"request_access_control": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"default_action_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"rules": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"action_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"condition_language": {
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
			"request_protection": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"rules": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"action_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"protection_capabilities": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"key": {
													Type:     schema.TypeString,
													Required: true,
												},
												"version": {
													Type:     schema.TypeInt,
													Required: true,
												},

												// Optional
												"action_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"collaborative_action_threshold": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"collaborative_weights": {
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
															"weight": {
																Type:     schema.TypeInt,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},
												"exclusions": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"args": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"request_cookies": {
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

												// Computed
											},
										},
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"condition_language": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"protection_capability_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"allowed_http_methods": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 255,
													MinItems: 1,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"max_http_request_header_length": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"max_http_request_headers": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"max_number_of_arguments": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"max_single_argument_length": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"max_total_argument_length": {
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
			"request_rate_limiting": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"rules": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"action_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"configurations": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"period_in_seconds": {
													Type:     schema.TypeInt,
													Required: true,
												},
												"requests_limit": {
													Type:     schema.TypeInt,
													Required: true,
												},

												// Optional
												"action_duration_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"condition_language": {
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
			"response_access_control": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"rules": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"action_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"condition_language": {
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
			"response_protection": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"rules": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"action_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"protection_capabilities": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"key": {
													Type:     schema.TypeString,
													Required: true,
												},
												"version": {
													Type:     schema.TypeInt,
													Required: true,
												},

												// Optional
												"action_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"collaborative_action_threshold": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"collaborative_weights": {
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
															"weight": {
																Type:     schema.TypeInt,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},
												"exclusions": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"args": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"request_cookies": {
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

												// Computed
											},
										},
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"condition_language": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"protection_capability_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"allowed_http_methods": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 255,
													MinItems: 1,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"max_http_request_header_length": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"max_http_request_headers": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"max_number_of_arguments": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"max_single_argument_length": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"max_total_argument_length": {
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

func createWafWebAppFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WafWebAppFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.CreateResource(d, sync)
}

func readWafWebAppFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WafWebAppFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.ReadResource(sync)
}

func updateWafWebAppFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WafWebAppFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteWafWebAppFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WafWebAppFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type WafWebAppFirewallPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_waf.WafClient
	Res                    *oci_waf.WebAppFirewallPolicy
	DisableNotFoundRetries bool
}

func (s *WafWebAppFirewallPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *WafWebAppFirewallPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_waf.WebAppFirewallPolicyLifecycleStateCreating),
	}
}

func (s *WafWebAppFirewallPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_waf.WebAppFirewallPolicyLifecycleStateActive),
	}
}

func (s *WafWebAppFirewallPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_waf.WebAppFirewallPolicyLifecycleStateDeleting),
	}
}

func (s *WafWebAppFirewallPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_waf.WebAppFirewallPolicyLifecycleStateDeleted),
	}
}

func (s *WafWebAppFirewallPolicyResourceCrud) Create() error {
	request := oci_waf.CreateWebAppFirewallPolicyRequest{}

	if actions, ok := s.D.GetOkExists("actions"); ok {
		interfaces := actions.([]interface{})
		tmp := make([]oci_waf.Action, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
			converted, err := s.mapToAction(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("actions") {
			request.Actions = tmp
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if requestAccessControl, ok := s.D.GetOkExists("request_access_control"); ok {
		if tmpList := requestAccessControl.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "request_access_control", 0)
			tmp, err := s.mapToRequestAccessControl(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RequestAccessControl = &tmp
		}
	}

	if requestProtection, ok := s.D.GetOkExists("request_protection"); ok {
		if tmpList := requestProtection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "request_protection", 0)
			tmp, err := s.mapToRequestProtection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RequestProtection = &tmp
		}
	}

	if requestRateLimiting, ok := s.D.GetOkExists("request_rate_limiting"); ok {
		if tmpList := requestRateLimiting.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "request_rate_limiting", 0)
			tmp, err := s.mapToRequestRateLimiting(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RequestRateLimiting = &tmp
		}
	}

	if responseAccessControl, ok := s.D.GetOkExists("response_access_control"); ok {
		if tmpList := responseAccessControl.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "response_access_control", 0)
			tmp, err := s.mapToResponseAccessControl(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResponseAccessControl = &tmp
		}
	}

	if responseProtection, ok := s.D.GetOkExists("response_protection"); ok {
		if tmpList := responseProtection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "response_protection", 0)
			tmp, err := s.mapToResponseProtection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResponseProtection = &tmp
		}
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf")

	response, err := s.Client.CreateWebAppFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getWebAppFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf"), oci_waf.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *WafWebAppFirewallPolicyResourceCrud) getWebAppFirewallPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_waf.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	webAppFirewallPolicyId, err := webAppFirewallPolicyWaitForWorkRequest(workId, "webAppFirewallPolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*webAppFirewallPolicyId)

	return s.Get()
}

func webAppFirewallPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "waf", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_waf.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func webAppFirewallPolicyWaitForWorkRequest(wId *string, entityType string, action oci_waf.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_waf.WafClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "waf")
	retryPolicy.ShouldRetryOperation = webAppFirewallPolicyWorkRequestShouldRetryFunc(timeout)

	response := oci_waf.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_waf.WorkRequestStatusInProgress),
			string(oci_waf.WorkRequestStatusAccepted),
			string(oci_waf.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_waf.WorkRequestStatusSucceeded),
			string(oci_waf.WorkRequestStatusFailed),
			string(oci_waf.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_waf.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_waf.WorkRequestStatusFailed || response.Status == oci_waf.WorkRequestStatusCanceled {
		return nil, getErrorFromWafWebAppFirewallPolicyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromWafWebAppFirewallPolicyWorkRequest(client *oci_waf.WafClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_waf.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_waf.ListWorkRequestErrorsRequest{
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

func (s *WafWebAppFirewallPolicyResourceCrud) Get() error {
	request := oci_waf.GetWebAppFirewallPolicyRequest{}

	tmp := s.D.Id()
	request.WebAppFirewallPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf")

	response, err := s.Client.GetWebAppFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.WebAppFirewallPolicy
	return nil
}

func (s *WafWebAppFirewallPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_waf.UpdateWebAppFirewallPolicyRequest{}

	if actions, ok := s.D.GetOkExists("actions"); ok {
		interfaces := actions.([]interface{})
		tmp := make([]oci_waf.Action, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
			converted, err := s.mapToAction(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("actions") {
			request.Actions = tmp
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if requestAccessControl, ok := s.D.GetOkExists("request_access_control"); ok {
		if tmpList := requestAccessControl.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "request_access_control", 0)
			tmp, err := s.mapToRequestAccessControl(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RequestAccessControl = &tmp
		}
	}

	if requestProtection, ok := s.D.GetOkExists("request_protection"); ok {
		if tmpList := requestProtection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "request_protection", 0)
			tmp, err := s.mapToRequestProtection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RequestProtection = &tmp
		}
	}

	if requestRateLimiting, ok := s.D.GetOkExists("request_rate_limiting"); ok {
		if tmpList := requestRateLimiting.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "request_rate_limiting", 0)
			tmp, err := s.mapToRequestRateLimiting(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RequestRateLimiting = &tmp
		}
	}

	if responseAccessControl, ok := s.D.GetOkExists("response_access_control"); ok {
		if tmpList := responseAccessControl.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "response_access_control", 0)
			tmp, err := s.mapToResponseAccessControl(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResponseAccessControl = &tmp
		}
	}

	if responseProtection, ok := s.D.GetOkExists("response_protection"); ok {
		if tmpList := responseProtection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "response_protection", 0)
			tmp, err := s.mapToResponseProtection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResponseProtection = &tmp
		}
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	tmp := s.D.Id()
	request.WebAppFirewallPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf")

	response, err := s.Client.UpdateWebAppFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getWebAppFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf"), oci_waf.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *WafWebAppFirewallPolicyResourceCrud) Delete() error {
	request := oci_waf.DeleteWebAppFirewallPolicyRequest{}

	tmp := s.D.Id()
	request.WebAppFirewallPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf")

	response, err := s.Client.DeleteWebAppFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := webAppFirewallPolicyWaitForWorkRequest(workId, "webAppFirewallPolicy",
		oci_waf.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *WafWebAppFirewallPolicyResourceCrud) SetData() error {
	actions := []interface{}{}
	for _, item := range s.Res.Actions {
		actions = append(actions, WafActionToMap(item))
	}
	s.D.Set("actions", actions)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RequestAccessControl != nil {
		s.D.Set("request_access_control", []interface{}{RequestAccessControlToMap(s.Res.RequestAccessControl)})
	} else {
		s.D.Set("request_access_control", nil)
	}

	if s.Res.RequestProtection != nil {
		s.D.Set("request_protection", []interface{}{RequestProtectionToMap(s.Res.RequestProtection)})
	} else {
		s.D.Set("request_protection", nil)
	}

	if s.Res.RequestRateLimiting != nil {
		s.D.Set("request_rate_limiting", []interface{}{RequestRateLimitingToMap(s.Res.RequestRateLimiting)})
	} else {
		s.D.Set("request_rate_limiting", nil)
	}

	if s.Res.ResponseAccessControl != nil {
		s.D.Set("response_access_control", []interface{}{ResponseAccessControlToMap(s.Res.ResponseAccessControl)})
	} else {
		s.D.Set("response_access_control", nil)
	}

	if s.Res.ResponseProtection != nil {
		s.D.Set("response_protection", []interface{}{ResponseProtectionToMap(s.Res.ResponseProtection)})
	} else {
		s.D.Set("response_protection", nil)
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

func (s *WafWebAppFirewallPolicyResourceCrud) mapToAccessControlRule(fieldKeyFormat string) (oci_waf.AccessControlRule, error) {
	result := oci_waf.AccessControlRule{}

	if actionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_name")); ok {
		tmp := actionName.(string)
		result.ActionName = &tmp
	}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		tmp := condition.(string)
		result.Condition = &tmp
	}

	if conditionLanguage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition_language")); ok {
		result.ConditionLanguage = oci_waf.WebAppFirewallPolicyRuleConditionLanguageEnum(conditionLanguage.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func AccessControlRuleToMap(obj oci_waf.AccessControlRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActionName != nil {
		result["action_name"] = string(*obj.ActionName)
	}

	if obj.Condition != nil {
		result["condition"] = string(*obj.Condition)
	}

	result["condition_language"] = string(obj.ConditionLanguage)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["type"] = string(oci_waf.WebAppFirewallPolicyRuleTypeAccessControl)

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToAction(fieldKeyFormat string) (oci_waf.Action, error) {
	var baseObject oci_waf.Action
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ALLOW"):
		details := oci_waf.AllowAction{}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	case strings.ToLower("CHECK"):
		details := oci_waf.CheckAction{}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	case strings.ToLower("RETURN_HTTP_RESPONSE"):
		details := oci_waf.ReturnHttpResponseAction{}
		if body, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "body")); ok {
			if tmpList := body.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "body"), 0)
				tmp, err := s.mapToHttpResponseBody(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert body, encountered error: %v", err)
				}
				details.Body = tmp
			}
		}
		if code, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "code")); ok {
			tmp := code.(int)
			details.Code = &tmp
		}
		if headers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "headers")); ok {
			interfaces := headers.([]interface{})
			tmp := make([]oci_waf.ResponseHeader, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "headers"), stateDataIndex)
				converted, err := s.mapToResponseHeader(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "headers")) {
				details.Headers = tmp
			}
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func WafActionToMap(obj oci_waf.Action) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_waf.AllowAction:
		result["type"] = "ALLOW"

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	case oci_waf.CheckAction:
		result["type"] = "CHECK"

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	case oci_waf.ReturnHttpResponseAction:
		result["type"] = "RETURN_HTTP_RESPONSE"

		if v.Body != nil {
			bodyArray := []interface{}{}
			if bodyMap := HttpResponseBodyToMap(&v.Body); bodyMap != nil {
				bodyArray = append(bodyArray, bodyMap)
			}
			result["body"] = bodyArray
		}

		if v.Code != nil {
			result["code"] = int(*v.Code)
		}

		headers := []interface{}{}
		for _, item := range v.Headers {
			headers = append(headers, ResponseHeaderToMap(item))
		}
		result["headers"] = headers

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToCollaborativeCapabilityWeightOverride(fieldKeyFormat string) (oci_waf.CollaborativeCapabilityWeightOverride, error) {
	result := oci_waf.CollaborativeCapabilityWeightOverride{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if weight, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weight")); ok {
		tmp := weight.(int)
		result.Weight = &tmp
	}

	return result, nil
}

func CollaborativeCapabilityWeightOverrideToMap(obj oci_waf.CollaborativeCapabilityWeightOverride) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Weight != nil {
		result["weight"] = int(*obj.Weight)
	}

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToHttpResponseBody(fieldKeyFormat string) (oci_waf.HttpResponseBody, error) {
	var baseObject oci_waf.HttpResponseBody
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("STATIC_TEXT"):
		details := oci_waf.StaticTextHttpResponseBody{}
		if text, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "text")); ok {
			tmp := text.(string)
			details.Text = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func HttpResponseBodyToMap(obj *oci_waf.HttpResponseBody) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_waf.StaticTextHttpResponseBody:
		result["type"] = "STATIC_TEXT"

		if v.Text != nil {
			result["text"] = string(*v.Text)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToProtectionCapability(fieldKeyFormat string) (oci_waf.ProtectionCapability, error) {
	result := oci_waf.ProtectionCapability{}

	if actionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_name")); ok {
		tmp := actionName.(string)
		result.ActionName = &tmp
	}

	if collaborativeActionThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "collaborative_action_threshold")); ok {
		tmp := collaborativeActionThreshold.(int)
		result.CollaborativeActionThreshold = &tmp
	}

	if collaborativeWeights, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "collaborative_weights")); ok {
		interfaces := collaborativeWeights.([]interface{})
		tmp := make([]oci_waf.CollaborativeCapabilityWeightOverride, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "collaborative_weights"), stateDataIndex)
			converted, err := s.mapToCollaborativeCapabilityWeightOverride(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "collaborative_weights")) {
			result.CollaborativeWeights = tmp
		}
	}

	if exclusions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclusions")); ok {
		if tmpList := exclusions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclusions"), 0)
			tmp, err := s.mapToProtectionCapabilityExclusions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert exclusions, encountered error: %v", err)
			}
			result.Exclusions = &tmp
		}
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
		tmp := version.(int)
		result.Version = &tmp
	}

	return result, nil
}

func ProtectionCapabilityToMap(obj oci_waf.ProtectionCapability) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActionName != nil {
		result["action_name"] = string(*obj.ActionName)
	}

	if obj.CollaborativeActionThreshold != nil {
		result["collaborative_action_threshold"] = int(*obj.CollaborativeActionThreshold)
	}

	collaborativeWeights := []interface{}{}
	for _, item := range obj.CollaborativeWeights {
		collaborativeWeights = append(collaborativeWeights, CollaborativeCapabilityWeightOverrideToMap(item))
	}
	result["collaborative_weights"] = collaborativeWeights

	if obj.Exclusions != nil {
		result["exclusions"] = []interface{}{ProtectionCapabilityExclusionsToMap(obj.Exclusions)}
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Version != nil {
		result["version"] = int(*obj.Version)
	}

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToProtectionCapabilityExclusions(fieldKeyFormat string) (oci_waf.ProtectionCapabilityExclusions, error) {
	result := oci_waf.ProtectionCapabilityExclusions{}

	if args, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "args")); ok {
		interfaces := args.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "args")) {
			result.Args = tmp
		}
	}

	if requestCookies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_cookies")); ok {
		interfaces := requestCookies.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "request_cookies")) {
			result.RequestCookies = tmp
		}
	}

	return result, nil
}

func ProtectionCapabilityExclusionsToMap(obj *oci_waf.ProtectionCapabilityExclusions) map[string]interface{} {
	result := map[string]interface{}{}

	result["args"] = obj.Args

	result["request_cookies"] = obj.RequestCookies

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToProtectionCapabilitySettings(fieldKeyFormat string) (oci_waf.ProtectionCapabilitySettings, error) {
	result := oci_waf.ProtectionCapabilitySettings{}

	if allowedHttpMethods, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_http_methods")); ok {
		interfaces := allowedHttpMethods.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_http_methods")) {
			result.AllowedHttpMethods = tmp
		}
	}

	if maxHttpRequestHeaderLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_http_request_header_length")); ok {
		tmp := maxHttpRequestHeaderLength.(int)
		result.MaxHttpRequestHeaderLength = &tmp
	}

	if maxHttpRequestHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_http_request_headers")); ok {
		tmp := maxHttpRequestHeaders.(int)
		result.MaxHttpRequestHeaders = &tmp
	}

	if maxNumberOfArguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_number_of_arguments")); ok {
		tmp := maxNumberOfArguments.(int)
		result.MaxNumberOfArguments = &tmp
	}

	if maxSingleArgumentLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_single_argument_length")); ok {
		tmp := maxSingleArgumentLength.(int)
		result.MaxSingleArgumentLength = &tmp
	}

	if maxTotalArgumentLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_total_argument_length")); ok {
		tmp := maxTotalArgumentLength.(int)
		result.MaxTotalArgumentLength = &tmp
	}

	return result, nil
}

func ProtectionCapabilitySettingsToMap(obj *oci_waf.ProtectionCapabilitySettings) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_http_methods"] = obj.AllowedHttpMethods

	if obj.MaxHttpRequestHeaderLength != nil {
		result["max_http_request_header_length"] = int(*obj.MaxHttpRequestHeaderLength)
	}

	if obj.MaxHttpRequestHeaders != nil {
		result["max_http_request_headers"] = int(*obj.MaxHttpRequestHeaders)
	}

	if obj.MaxNumberOfArguments != nil {
		result["max_number_of_arguments"] = int(*obj.MaxNumberOfArguments)
	}

	if obj.MaxSingleArgumentLength != nil {
		result["max_single_argument_length"] = int(*obj.MaxSingleArgumentLength)
	}

	if obj.MaxTotalArgumentLength != nil {
		result["max_total_argument_length"] = int(*obj.MaxTotalArgumentLength)
	}

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToProtectionRule(fieldKeyFormat string) (oci_waf.ProtectionRule, error) {
	result := oci_waf.ProtectionRule{}

	if actionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_name")); ok {
		tmp := actionName.(string)
		result.ActionName = &tmp
	}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		tmp := condition.(string)
		result.Condition = &tmp
	}

	if conditionLanguage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition_language")); ok {
		result.ConditionLanguage = oci_waf.WebAppFirewallPolicyRuleConditionLanguageEnum(conditionLanguage.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if protectionCapabilities, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_capabilities")); ok {
		interfaces := protectionCapabilities.([]interface{})
		tmp := make([]oci_waf.ProtectionCapability, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "protection_capabilities"), stateDataIndex)
			converted, err := s.mapToProtectionCapability(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "protection_capabilities")) {
			result.ProtectionCapabilities = tmp
		}
	}

	if protectionCapabilitySettings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_capability_settings")); ok {
		if tmpList := protectionCapabilitySettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "protection_capability_settings"), 0)
			tmp, err := s.mapToProtectionCapabilitySettings(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert protection_capability_settings, encountered error: %v", err)
			}
			result.ProtectionCapabilitySettings = &tmp
		}
	}

	return result, nil
}

func ProtectionRuleToMap(obj oci_waf.ProtectionRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActionName != nil {
		result["action_name"] = string(*obj.ActionName)
	}

	if obj.Condition != nil {
		result["condition"] = string(*obj.Condition)
	}

	result["condition_language"] = string(obj.ConditionLanguage)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	protectionCapabilities := []interface{}{}
	for _, item := range obj.ProtectionCapabilities {
		protectionCapabilities = append(protectionCapabilities, ProtectionCapabilityToMap(item))
	}
	result["protection_capabilities"] = protectionCapabilities

	if obj.ProtectionCapabilitySettings != nil {
		result["protection_capability_settings"] = []interface{}{ProtectionCapabilitySettingsToMap(obj.ProtectionCapabilitySettings)}
	}

	result["type"] = string(oci_waf.WebAppFirewallPolicyRuleTypeProtection)

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToRequestAccessControl(fieldKeyFormat string) (oci_waf.RequestAccessControl, error) {
	result := oci_waf.RequestAccessControl{}

	if defaultActionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_action_name")); ok {
		tmp := defaultActionName.(string)
		result.DefaultActionName = &tmp
	}

	if rules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rules")); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_waf.AccessControlRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rules"), stateDataIndex)
			converted, err := s.mapToAccessControlRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "rules")) {
			result.Rules = tmp
		}
	}

	return result, nil
}

func RequestAccessControlToMap(obj *oci_waf.RequestAccessControl) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultActionName != nil {
		result["default_action_name"] = string(*obj.DefaultActionName)
	}

	rules := []interface{}{}
	for _, item := range obj.Rules {
		rules = append(rules, AccessControlRuleToMap(item))
	}
	result["rules"] = rules

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToRequestProtection(fieldKeyFormat string) (oci_waf.RequestProtection, error) {
	result := oci_waf.RequestProtection{}

	if rules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rules")); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_waf.ProtectionRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rules"), stateDataIndex)
			converted, err := s.mapToProtectionRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "rules")) {
			result.Rules = tmp
		}
	}

	return result, nil
}

func RequestProtectionToMap(obj *oci_waf.RequestProtection) map[string]interface{} {
	result := map[string]interface{}{}

	rules := []interface{}{}
	for _, item := range obj.Rules {
		rules = append(rules, ProtectionRuleToMap(item))
	}
	result["rules"] = rules

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToRequestRateLimiting(fieldKeyFormat string) (oci_waf.RequestRateLimiting, error) {
	result := oci_waf.RequestRateLimiting{}

	if rules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rules")); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_waf.RequestRateLimitingRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rules"), stateDataIndex)
			converted, err := s.mapToRequestRateLimitingRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "rules")) {
			result.Rules = tmp
		}
	}

	return result, nil
}

func RequestRateLimitingToMap(obj *oci_waf.RequestRateLimiting) map[string]interface{} {
	result := map[string]interface{}{}

	rules := []interface{}{}
	for _, item := range obj.Rules {
		rules = append(rules, RequestRateLimitingRuleToMap(item))
	}
	result["rules"] = rules

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToRequestRateLimitingConfiguration(fieldKeyFormat string) (oci_waf.RequestRateLimitingConfiguration, error) {
	result := oci_waf.RequestRateLimitingConfiguration{}

	if actionDurationInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_duration_in_seconds")); ok {
		tmp := actionDurationInSeconds.(int)
		result.ActionDurationInSeconds = &tmp
	}

	if periodInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "period_in_seconds")); ok {
		tmp := periodInSeconds.(int)
		result.PeriodInSeconds = &tmp
	}

	if requestsLimit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "requests_limit")); ok {
		tmp := requestsLimit.(int)
		result.RequestsLimit = &tmp
	}

	return result, nil
}

func RequestRateLimitingConfigurationToMap(obj oci_waf.RequestRateLimitingConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActionDurationInSeconds != nil {
		result["action_duration_in_seconds"] = int(*obj.ActionDurationInSeconds)
	}

	if obj.PeriodInSeconds != nil {
		result["period_in_seconds"] = int(*obj.PeriodInSeconds)
	}

	if obj.RequestsLimit != nil {
		result["requests_limit"] = int(*obj.RequestsLimit)
	}

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToRequestRateLimitingRule(fieldKeyFormat string) (oci_waf.RequestRateLimitingRule, error) {
	result := oci_waf.RequestRateLimitingRule{}

	if actionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_name")); ok {
		tmp := actionName.(string)
		result.ActionName = &tmp
	}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		tmp := condition.(string)
		result.Condition = &tmp
	}

	if conditionLanguage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition_language")); ok {
		result.ConditionLanguage = oci_waf.WebAppFirewallPolicyRuleConditionLanguageEnum(conditionLanguage.(string))
	}

	if configurations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configurations")); ok {
		interfaces := configurations.([]interface{})
		tmp := make([]oci_waf.RequestRateLimitingConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configurations"), stateDataIndex)
			converted, err := s.mapToRequestRateLimitingConfiguration(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "configurations")) {
			result.Configurations = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func RequestRateLimitingRuleToMap(obj oci_waf.RequestRateLimitingRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActionName != nil {
		result["action_name"] = string(*obj.ActionName)
	}

	if obj.Condition != nil {
		result["condition"] = string(*obj.Condition)
	}

	result["condition_language"] = string(obj.ConditionLanguage)

	configurations := []interface{}{}
	for _, item := range obj.Configurations {
		configurations = append(configurations, RequestRateLimitingConfigurationToMap(item))
	}
	result["configurations"] = configurations

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["type"] = string(oci_waf.WebAppFirewallPolicyRuleTypeRequestRateLimiting)

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToResponseAccessControl(fieldKeyFormat string) (oci_waf.ResponseAccessControl, error) {
	result := oci_waf.ResponseAccessControl{}

	if rules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rules")); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_waf.AccessControlRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rules"), stateDataIndex)
			converted, err := s.mapToAccessControlRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "rules")) {
			result.Rules = tmp
		}
	}

	return result, nil
}

func ResponseAccessControlToMap(obj *oci_waf.ResponseAccessControl) map[string]interface{} {
	result := map[string]interface{}{}

	rules := []interface{}{}
	for _, item := range obj.Rules {
		rules = append(rules, AccessControlRuleToMap(item))
	}
	result["rules"] = rules

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToResponseHeader(fieldKeyFormat string) (oci_waf.ResponseHeader, error) {
	result := oci_waf.ResponseHeader{}

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

func ResponseHeaderToMap(obj oci_waf.ResponseHeader) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *WafWebAppFirewallPolicyResourceCrud) mapToResponseProtection(fieldKeyFormat string) (oci_waf.ResponseProtection, error) {
	result := oci_waf.ResponseProtection{}

	if rules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rules")); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_waf.ProtectionRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rules"), stateDataIndex)
			converted, err := s.mapToProtectionRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "rules")) {
			result.Rules = tmp
		}
	}

	return result, nil
}

func ResponseProtectionToMap(obj *oci_waf.ResponseProtection) map[string]interface{} {
	result := map[string]interface{}{}

	rules := []interface{}{}
	for _, item := range obj.Rules {
		rules = append(rules, ProtectionRuleToMap(item))
	}
	result["rules"] = rules

	return result
}

func WebAppFirewallPolicySummaryToMap(obj oci_waf.WebAppFirewallPolicySummary) map[string]interface{} {
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

func (s *WafWebAppFirewallPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_waf.ChangeWebAppFirewallPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.WebAppFirewallPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf")

	response, err := s.Client.ChangeWebAppFirewallPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getWebAppFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waf"), oci_waf.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
