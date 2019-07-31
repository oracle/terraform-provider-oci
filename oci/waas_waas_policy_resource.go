// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_waas "github.com/oracle/oci-go-sdk/waas"
)

func WaasWaasPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createWaasWaasPolicy,
		Read:     readWaasWaasPolicy,
		Update:   updateWaasWaasPolicy,
		Delete:   deleteWaasWaasPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"additional_domains": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
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
			"origins": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      originsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"label": {
							Type:     schema.TypeString,
							Required: true,
						},

						"uri": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"http_port": {
							Type:     schema.TypeInt,
							Default:  80,
							Optional: true,
						},
						"https_port": {
							Type:     schema.TypeInt,
							Default:  443,
							Optional: true,
						},
						"custom_headers": {
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
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						// Computed
					},
				},
			},
			"policy_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"certificate_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_https_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_https_forced": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"waf_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"access_rules": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"action": {
										Type:     schema.TypeString,
										Required: true,
									},
									"criteria": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"condition": {
													Type:     schema.TypeString,
													Required: true,
												},
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional

												// Computed
											},
										},
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"block_action": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"block_error_page_code": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"block_error_page_description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"block_error_page_message": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"block_response_code": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"address_rate_limiting": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"is_enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},

									// Optional
									"allowed_rate_per_address": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"block_response_code": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_delayed_count_per_address": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"captchas": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"failure_message": {
										Type:     schema.TypeString,
										Required: true,
									},
									"session_expiration_in_seconds": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"submit_label": {
										Type:     schema.TypeString,
										Required: true,
									},
									"title": {
										Type:     schema.TypeString,
										Required: true,
									},
									"url": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"footer_text": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"header_text": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"device_fingerprint_challenge": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"is_enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},

									// Optional
									"action": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action_expiration_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"challenge_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"block_action": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_error_page_code": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_error_page_description": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_error_page_message": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_response_code": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"captcha_footer": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"captcha_header": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"captcha_submit_label": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"captcha_title": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"failure_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"failure_threshold_expiration_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_address_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_address_count_expiration_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"human_interaction_challenge": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"is_enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},

									// Optional
									"action": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action_expiration_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"challenge_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"block_action": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_error_page_code": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_error_page_description": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_error_page_message": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_response_code": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"captcha_footer": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"captcha_header": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"captcha_submit_label": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"captcha_title": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"failure_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"failure_threshold_expiration_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"interaction_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"recording_period_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"set_http_header": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"value": {
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
						"js_challenge": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"is_enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},

									// Optional
									"action": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action_expiration_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"challenge_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"block_action": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_error_page_code": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_error_page_description": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_error_page_message": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"block_response_code": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"captcha_footer": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"captcha_header": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"captcha_submit_label": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"captcha_title": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"failure_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"set_http_header": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"value": {
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
						"origin": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protection_settings": {
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
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"block_action": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"block_error_page_code": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"block_error_page_description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"block_error_page_message": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"block_response_code": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"is_response_inspected": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"max_argument_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_name_length_per_argument": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_response_size_in_ki_b": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_total_name_length_of_arguments": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"media_types": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"recommendations_period_in_days": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"whitelists": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"addresses": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"name": {
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

			// Computed
			"cname": {
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
		},
	}
}

func createWaasWaasPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WaasWaasPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient

	return CreateResource(d, sync)
}

func readWaasWaasPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WaasWaasPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient

	return ReadResource(sync)
}

func updateWaasWaasPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WaasWaasPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient

	return UpdateResource(d, sync)
}

func deleteWaasWaasPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WaasWaasPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type WaasWaasPolicyResourceCrud struct {
	BaseCrud
	Client                 *oci_waas.WaasClient
	Res                    *oci_waas.WaasPolicy
	DisableNotFoundRetries bool
}

func (s *WaasWaasPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *WaasWaasPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_waas.WaasPolicyLifecycleStateCreating),
	}
}

func (s *WaasWaasPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_waas.WaasPolicyLifecycleStateActive),
		string(oci_waas.WaasPolicyLifecycleStateFailed),
	}
}

func (s *WaasWaasPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_waas.WaasPolicyLifecycleStateDeleting),
	}
}

func (s *WaasWaasPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_waas.WaasPolicyLifecycleStateDeleted),
	}
}

func (s *WaasWaasPolicyResourceCrud) Create() error {
	request := oci_waas.CreateWaasPolicyRequest{}

	request.AdditionalDomains = []string{}
	if additionalDomains, ok := s.D.GetOkExists("additional_domains"); ok {
		interfaces := additionalDomains.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.AdditionalDomains = tmp
	}

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

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if origins, ok := s.D.GetOkExists("origins"); ok {
		resultMap, err := s.objectMapToOriginMap(origins)
		if err != nil {
			return err
		}
		if len(resultMap) > 0 {
			request.Origins = resultMap
		}
	}

	if policyConfig, ok := s.D.GetOkExists("policy_config"); ok {
		if tmpList := policyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "policy_config", 0)
			tmp, err := s.mapToPolicyConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PolicyConfig = &tmp
		}
	}

	if wafConfig, ok := s.D.GetOkExists("waf_config"); ok {
		if tmpList := wafConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "waf_config", 0)
			tmp, err := s.mapToWafConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.WafConfig = &tmp
		}
	}

	retryPolicy := getRetryPolicy(s.DisableNotFoundRetries, "waas")
	request.RequestMetadata.RetryPolicy = retryPolicy

	response, err := s.Client.CreateWaasPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getWaasPolicyFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "waas"), oci_waas.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *WaasWaasPolicyResourceCrud) getWaasPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_waas.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	waasPolicyId, err := waasPolicyWaitForWorkRequest(workId, "waas",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, waasPolicyId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_waas.CancelWorkRequestRequest{
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
	s.D.SetId(*waasPolicyId)

	return s.Get()
}

func waasPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "waas", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_waas.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func waasPolicyWaitForWorkRequest(wId *string, entityType string, action oci_waas.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_waas.WaasClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "waas")
	retryPolicy.ShouldRetryOperation = waasPolicyWorkRequestShouldRetryFunc(timeout)

	response := oci_waas.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_waas.WorkRequestStatusInProgress),
			string(oci_waas.WorkRequestStatusAccepted),
			string(oci_waas.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_waas.WorkRequestStatusSucceeded),
			string(oci_waas.WorkRequestStatusFailed),
			string(oci_waas.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_waas.GetWorkRequestRequest{
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

	// The workrequest didn't do all its intended tasks, if the errors is set; so we should check for it
	var workRequestErr error
	if len(response.Errors) > 0 {
		errorMessage := getErrorFromWaasPolicyWorkRequest(response)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)
	}

	return identifier, workRequestErr
}

func getErrorFromWaasPolicyWorkRequest(response oci_waas.GetWorkRequestResponse) string {
	allErrs := make([]string, 0)
	for _, wrkErr := range response.Errors {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")
	return errorMessage
}

func (s *WaasWaasPolicyResourceCrud) mapToOrigin(fieldKeyFormat string) (oci_waas.Origin, error) {
	result := oci_waas.Origin{}

	result.CustomHeaders = []oci_waas.Header{}
	if customHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_headers")); ok {
		interfaces := customHeaders.([]interface{})
		tmp := make([]oci_waas.Header, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "custom_headers"), stateDataIndex)
			converted, err := s.mapToHeader(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		result.CustomHeaders = tmp
	}

	if httpPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "http_port")); ok {
		tmp := httpPort.(int)
		result.HttpPort = &tmp
	}

	if httpsPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "https_port")); ok {
		tmp := httpsPort.(int)
		result.HttpsPort = &tmp
	}

	if uri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
		tmp := uri.(string)
		result.Uri = &tmp
	}

	return result, nil
}

func (s *WaasWaasPolicyResourceCrud) Get() error {
	request := oci_waas.GetWaasPolicyRequest{}

	tmp := s.D.Id()
	request.WaasPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.GetWaasPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.WaasPolicy
	return nil
}

func (s *WaasWaasPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_waas.UpdateWaasPolicyRequest{}

	request.AdditionalDomains = []string{}
	if additionalDomains, ok := s.D.GetOkExists("additional_domains"); ok {
		interfaces := additionalDomains.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.AdditionalDomains = tmp
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

	if origins, ok := s.D.GetOkExists("origins"); ok {
		resultMap, err := s.objectMapToOriginMap(origins)
		if err != nil {
			return err
		}
		request.Origins = resultMap
	}

	if policyConfig, ok := s.D.GetOkExists("policy_config"); ok {
		if tmpList := policyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "policy_config", 0)
			tmp, err := s.mapToPolicyConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PolicyConfig = &tmp
		}
	}

	tmp := s.D.Id()
	request.WaasPolicyId = &tmp

	if wafConfig, ok := s.D.GetOkExists("waf_config"); ok {
		if tmpList := wafConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "waf_config", 0)
			tmp, err := s.mapToWafConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.WafConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.UpdateWaasPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getWaasPolicyFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "waas"), oci_waas.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *WaasWaasPolicyResourceCrud) objectMapToOriginMap(origins interface{}) (map[string]oci_waas.Origin, error) {

	resultMap := map[string]oci_waas.Origin{}
	set := origins.(*schema.Set)
	tmpList := set.List()
	for _, ifc := range tmpList {
		stateDataIndex := originsHashCodeForSets(ifc)
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "origins", stateDataIndex)
		converted, err := s.mapToOrigin(fieldKeyFormat)
		if err != nil {
			return nil, err
		}

		if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
			tmp := label.(string)
			resultMap[tmp] = converted
		}
	}

	return resultMap, nil
}

func (s *WaasWaasPolicyResourceCrud) Delete() error {
	request := oci_waas.DeleteWaasPolicyRequest{}

	tmp := s.D.Id()
	request.WaasPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.DeleteWaasPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := waasPolicyWaitForWorkRequest(workId, "waas",
		oci_waas.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *WaasWaasPolicyResourceCrud) SetData() error {
	s.D.Set("additional_domains", s.Res.AdditionalDomains)

	if s.Res.Cname != nil {
		s.D.Set("cname", *s.Res.Cname)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Origins != nil {
		s.D.Set("origins", schema.NewSet(originsHashCodeForSets, OriginMapToMap(s.Res.Origins)))
	} else {
		s.D.Set("origins", nil)
	}

	if s.Res.PolicyConfig != nil {
		s.D.Set("policy_config", []interface{}{PolicyConfigToMap(s.Res.PolicyConfig)})
	} else {
		s.D.Set("policy_config", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.WafConfig != nil {
		s.D.Set("waf_config", []interface{}{WafConfigToMap(s.Res.WafConfig)})
	} else {
		s.D.Set("waf_config", nil)
	}

	return nil
}

func (s *WaasWaasPolicyResourceCrud) mapToAccessRule(fieldKeyFormat string) (oci_waas.AccessRule, error) {
	result := oci_waas.AccessRule{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_waas.AccessRuleActionEnum(action.(string))
	}

	if blockAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_action")); ok {
		result.BlockAction = oci_waas.AccessRuleBlockActionEnum(blockAction.(string))
	}

	if blockErrorPageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_error_page_code")); ok {
		tmp := blockErrorPageCode.(string)
		result.BlockErrorPageCode = &tmp
	}

	if blockErrorPageDescription, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_error_page_description")); ok {
		tmp := blockErrorPageDescription.(string)
		result.BlockErrorPageDescription = &tmp
	}

	if blockErrorPageMessage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_error_page_message")); ok {
		tmp := blockErrorPageMessage.(string)
		result.BlockErrorPageMessage = &tmp
	}

	if blockResponseCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_response_code")); ok {
		tmp := blockResponseCode.(int)
		result.BlockResponseCode = &tmp
	}

	result.Criteria = []oci_waas.AccessRuleCriteria{}
	if criteria, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "criteria")); ok {
		interfaces := criteria.([]interface{})
		tmp := make([]oci_waas.AccessRuleCriteria, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "criteria"), stateDataIndex)
			converted, err := s.mapToAccessRuleCriteria(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		result.Criteria = tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func AccessRuleToMap(obj oci_waas.AccessRule) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	result["block_action"] = string(obj.BlockAction)

	if obj.BlockErrorPageCode != nil {
		result["block_error_page_code"] = string(*obj.BlockErrorPageCode)
	}

	if obj.BlockErrorPageDescription != nil {
		result["block_error_page_description"] = string(*obj.BlockErrorPageDescription)
	}

	if obj.BlockErrorPageMessage != nil {
		result["block_error_page_message"] = string(*obj.BlockErrorPageMessage)
	}

	if obj.BlockResponseCode != nil {
		result["block_response_code"] = int(*obj.BlockResponseCode)
	}

	criteria := []interface{}{}
	for _, item := range obj.Criteria {
		criteria = append(criteria, AccessRuleCriteriaToMap(item))
	}
	result["criteria"] = criteria

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *WaasWaasPolicyResourceCrud) mapToAccessRuleCriteria(fieldKeyFormat string) (oci_waas.AccessRuleCriteria, error) {
	result := oci_waas.AccessRuleCriteria{}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		result.Condition = oci_waas.AccessRuleCriteriaConditionEnum(condition.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AccessRuleCriteriaToMap(obj oci_waas.AccessRuleCriteria) map[string]interface{} {
	result := map[string]interface{}{}

	result["condition"] = string(obj.Condition)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *WaasWaasPolicyResourceCrud) mapToAddressRateLimiting(fieldKeyFormat string) (oci_waas.AddressRateLimiting, error) {
	result := oci_waas.AddressRateLimiting{}

	if allowedRatePerAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_rate_per_address")); ok {
		tmp := allowedRatePerAddress.(int)
		result.AllowedRatePerAddress = &tmp
	}

	if blockResponseCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_response_code")); ok {
		tmp := blockResponseCode.(int)
		result.BlockResponseCode = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if maxDelayedCountPerAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_delayed_count_per_address")); ok {
		tmp := maxDelayedCountPerAddress.(int)
		result.MaxDelayedCountPerAddress = &tmp
	}

	return result, nil
}

func AddressRateLimitingToMap(obj *oci_waas.AddressRateLimiting) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllowedRatePerAddress != nil {
		result["allowed_rate_per_address"] = int(*obj.AllowedRatePerAddress)
	}

	if obj.BlockResponseCode != nil {
		result["block_response_code"] = int(*obj.BlockResponseCode)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.MaxDelayedCountPerAddress != nil {
		result["max_delayed_count_per_address"] = int(*obj.MaxDelayedCountPerAddress)
	}

	return result
}

func (s *WaasWaasPolicyResourceCrud) mapToBlockChallengeSettings(fieldKeyFormat string) (oci_waas.BlockChallengeSettings, error) {
	result := oci_waas.BlockChallengeSettings{}

	if blockAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_action")); ok {
		result.BlockAction = oci_waas.BlockChallengeSettingsBlockActionEnum(blockAction.(string))
	}

	if blockErrorPageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_error_page_code")); ok {
		tmp := blockErrorPageCode.(string)
		result.BlockErrorPageCode = &tmp
	}

	if blockErrorPageDescription, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_error_page_description")); ok {
		tmp := blockErrorPageDescription.(string)
		result.BlockErrorPageDescription = &tmp
	}

	if blockErrorPageMessage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_error_page_message")); ok {
		tmp := blockErrorPageMessage.(string)
		result.BlockErrorPageMessage = &tmp
	}

	if blockResponseCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_response_code")); ok {
		tmp := blockResponseCode.(int)
		result.BlockResponseCode = &tmp
	}

	if captchaFooter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "captcha_footer")); ok {
		tmp := captchaFooter.(string)
		result.CaptchaFooter = &tmp
	}

	if captchaHeader, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "captcha_header")); ok {
		tmp := captchaHeader.(string)
		result.CaptchaHeader = &tmp
	}

	if captchaSubmitLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "captcha_submit_label")); ok {
		tmp := captchaSubmitLabel.(string)
		result.CaptchaSubmitLabel = &tmp
	}

	if captchaTitle, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "captcha_title")); ok {
		tmp := captchaTitle.(string)
		result.CaptchaTitle = &tmp
	}

	return result, nil
}

func BlockChallengeSettingsToMap(obj *oci_waas.BlockChallengeSettings) map[string]interface{} {
	result := map[string]interface{}{}

	result["block_action"] = string(obj.BlockAction)

	if obj.BlockErrorPageCode != nil {
		result["block_error_page_code"] = string(*obj.BlockErrorPageCode)
	}

	if obj.BlockErrorPageDescription != nil {
		result["block_error_page_description"] = string(*obj.BlockErrorPageDescription)
	}

	if obj.BlockErrorPageMessage != nil {
		result["block_error_page_message"] = string(*obj.BlockErrorPageMessage)
	}

	if obj.BlockResponseCode != nil {
		result["block_response_code"] = int(*obj.BlockResponseCode)
	}

	if obj.CaptchaFooter != nil {
		result["captcha_footer"] = string(*obj.CaptchaFooter)
	}

	if obj.CaptchaHeader != nil {
		result["captcha_header"] = string(*obj.CaptchaHeader)
	}

	if obj.CaptchaSubmitLabel != nil {
		result["captcha_submit_label"] = string(*obj.CaptchaSubmitLabel)
	}

	if obj.CaptchaTitle != nil {
		result["captcha_title"] = string(*obj.CaptchaTitle)
	}

	return result
}

func (s *WaasWaasPolicyResourceCrud) mapToCaptcha(fieldKeyFormat string) (oci_waas.Captcha, error) {
	result := oci_waas.Captcha{}

	if failureMessage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_message")); ok {
		tmp := failureMessage.(string)
		result.FailureMessage = &tmp
	}

	if footerText, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "footer_text")); ok {
		tmp := footerText.(string)
		result.FooterText = &tmp
	}

	if headerText, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header_text")); ok {
		tmp := headerText.(string)
		result.HeaderText = &tmp
	}

	if sessionExpirationInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "session_expiration_in_seconds")); ok {
		tmp := sessionExpirationInSeconds.(int)
		result.SessionExpirationInSeconds = &tmp
	}

	if submitLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "submit_label")); ok {
		tmp := submitLabel.(string)
		result.SubmitLabel = &tmp
	}

	if title, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "title")); ok {
		tmp := title.(string)
		result.Title = &tmp
	}

	if url, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url")); ok {
		tmp := url.(string)
		result.Url = &tmp
	}

	return result, nil
}

func CaptchaToMap(obj oci_waas.Captcha) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FailureMessage != nil {
		result["failure_message"] = string(*obj.FailureMessage)
	}

	if obj.FooterText != nil {
		result["footer_text"] = string(*obj.FooterText)
	}

	if obj.HeaderText != nil {
		result["header_text"] = string(*obj.HeaderText)
	}

	if obj.SessionExpirationInSeconds != nil {
		result["session_expiration_in_seconds"] = int(*obj.SessionExpirationInSeconds)
	}

	if obj.SubmitLabel != nil {
		result["submit_label"] = string(*obj.SubmitLabel)
	}

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func (s *WaasWaasPolicyResourceCrud) mapToDeviceFingerprintChallenge(fieldKeyFormat string) (oci_waas.DeviceFingerprintChallenge, error) {
	result := oci_waas.DeviceFingerprintChallenge{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_waas.DeviceFingerprintChallengeActionEnum(action.(string))
	}

	if actionExpirationInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_expiration_in_seconds")); ok {
		tmp := actionExpirationInSeconds.(int)
		result.ActionExpirationInSeconds = &tmp
	}

	if challengeSettings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "challenge_settings")); ok {
		if tmpList := challengeSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "challenge_settings"), 0)
			tmp, err := s.mapToBlockChallengeSettings(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert challenge_settings, encountered error: %v", err)
			}
			result.ChallengeSettings = &tmp
		}
	}

	if failureThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_threshold")); ok {
		tmp := failureThreshold.(int)
		result.FailureThreshold = &tmp
	}

	if failureThresholdExpirationInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_threshold_expiration_in_seconds")); ok {
		tmp := failureThresholdExpirationInSeconds.(int)
		result.FailureThresholdExpirationInSeconds = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if maxAddressCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_address_count")); ok {
		tmp := maxAddressCount.(int)
		result.MaxAddressCount = &tmp
	}

	if maxAddressCountExpirationInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_address_count_expiration_in_seconds")); ok {
		tmp := maxAddressCountExpirationInSeconds.(int)
		result.MaxAddressCountExpirationInSeconds = &tmp
	}

	return result, nil
}

func DeviceFingerprintChallengeToMap(obj *oci_waas.DeviceFingerprintChallenge) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.ActionExpirationInSeconds != nil {
		result["action_expiration_in_seconds"] = int(*obj.ActionExpirationInSeconds)
	}

	if obj.ChallengeSettings != nil {
		result["challenge_settings"] = []interface{}{BlockChallengeSettingsToMap(obj.ChallengeSettings)}
	}

	if obj.FailureThreshold != nil {
		result["failure_threshold"] = int(*obj.FailureThreshold)
	}

	if obj.FailureThresholdExpirationInSeconds != nil {
		result["failure_threshold_expiration_in_seconds"] = int(*obj.FailureThresholdExpirationInSeconds)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.MaxAddressCount != nil {
		result["max_address_count"] = int(*obj.MaxAddressCount)
	}

	if obj.MaxAddressCountExpirationInSeconds != nil {
		result["max_address_count_expiration_in_seconds"] = int(*obj.MaxAddressCountExpirationInSeconds)
	}

	return result
}

func (s *WaasWaasPolicyResourceCrud) mapToHeader(fieldKeyFormat string) (oci_waas.Header, error) {
	result := oci_waas.Header{}

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

func HeaderToMap(obj *oci_waas.Header) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *WaasWaasPolicyResourceCrud) mapToHumanInteractionChallenge(fieldKeyFormat string) (oci_waas.HumanInteractionChallenge, error) {
	result := oci_waas.HumanInteractionChallenge{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_waas.HumanInteractionChallengeActionEnum(action.(string))
	}

	if actionExpirationInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_expiration_in_seconds")); ok {
		tmp := actionExpirationInSeconds.(int)
		result.ActionExpirationInSeconds = &tmp
	}

	if challengeSettings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "challenge_settings")); ok {
		if tmpList := challengeSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "challenge_settings"), 0)
			tmp, err := s.mapToBlockChallengeSettings(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert challenge_settings, encountered error: %v", err)
			}
			result.ChallengeSettings = &tmp
		}
	}

	if failureThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_threshold")); ok {
		tmp := failureThreshold.(int)
		result.FailureThreshold = &tmp
	}

	if failureThresholdExpirationInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_threshold_expiration_in_seconds")); ok {
		tmp := failureThresholdExpirationInSeconds.(int)
		result.FailureThresholdExpirationInSeconds = &tmp
	}

	if interactionThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interaction_threshold")); ok {
		tmp := interactionThreshold.(int)
		result.InteractionThreshold = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if recordingPeriodInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recording_period_in_seconds")); ok {
		tmp := recordingPeriodInSeconds.(int)
		result.RecordingPeriodInSeconds = &tmp
	}

	if setHttpHeader, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "set_http_header")); ok {
		if tmpList := setHttpHeader.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "set_http_header"), 0)
			tmp, err := s.mapToHeader(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert set_http_header, encountered error: %v", err)
			}
			result.SetHttpHeader = &tmp
		}
	}

	return result, nil
}

func HumanInteractionChallengeToMap(obj *oci_waas.HumanInteractionChallenge) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.ActionExpirationInSeconds != nil {
		result["action_expiration_in_seconds"] = int(*obj.ActionExpirationInSeconds)
	}

	if obj.ChallengeSettings != nil {
		result["challenge_settings"] = []interface{}{BlockChallengeSettingsToMap(obj.ChallengeSettings)}
	}

	if obj.FailureThreshold != nil {
		result["failure_threshold"] = int(*obj.FailureThreshold)
	}

	if obj.FailureThresholdExpirationInSeconds != nil {
		result["failure_threshold_expiration_in_seconds"] = int(*obj.FailureThresholdExpirationInSeconds)
	}

	if obj.InteractionThreshold != nil {
		result["interaction_threshold"] = int(*obj.InteractionThreshold)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.RecordingPeriodInSeconds != nil {
		result["recording_period_in_seconds"] = int(*obj.RecordingPeriodInSeconds)
	}

	if obj.SetHttpHeader != nil {
		result["set_http_header"] = []interface{}{HeaderToMap(obj.SetHttpHeader)}
	}

	return result
}

func (s *WaasWaasPolicyResourceCrud) mapToJsChallenge(fieldKeyFormat string) (oci_waas.JsChallenge, error) {
	result := oci_waas.JsChallenge{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_waas.JsChallengeActionEnum(action.(string))
	}

	if actionExpirationInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_expiration_in_seconds")); ok {
		tmp := actionExpirationInSeconds.(int)
		result.ActionExpirationInSeconds = &tmp
	}

	if challengeSettings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "challenge_settings")); ok {
		if tmpList := challengeSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "challenge_settings"), 0)
			tmp, err := s.mapToBlockChallengeSettings(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert challenge_settings, encountered error: %v", err)
			}
			result.ChallengeSettings = &tmp
		}
	}

	if failureThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_threshold")); ok {
		tmp := failureThreshold.(int)
		result.FailureThreshold = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if setHttpHeader, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "set_http_header")); ok {
		if tmpList := setHttpHeader.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "set_http_header"), 0)
			tmp, err := s.mapToHeader(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert set_http_header, encountered error: %v", err)
			}
			result.SetHttpHeader = &tmp
		}
	}

	return result, nil
}

func JsChallengeToMap(obj *oci_waas.JsChallenge) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.ActionExpirationInSeconds != nil {
		result["action_expiration_in_seconds"] = int(*obj.ActionExpirationInSeconds)
	}

	if obj.ChallengeSettings != nil {
		result["challenge_settings"] = []interface{}{BlockChallengeSettingsToMap(obj.ChallengeSettings)}
	}

	if obj.FailureThreshold != nil {
		result["failure_threshold"] = int(*obj.FailureThreshold)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.SetHttpHeader != nil {
		result["set_http_header"] = []interface{}{HeaderToMap(obj.SetHttpHeader)}
	}

	return result
}

func (s *WaasWaasPolicyResourceCrud) mapToPolicyConfig(fieldKeyFormat string) (oci_waas.PolicyConfig, error) {
	result := oci_waas.PolicyConfig{}

	if certificateId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_id")); ok {
		tmp := certificateId.(string)
		result.CertificateId = &tmp
	}

	if isHttpsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_https_enabled")); ok {
		tmp := isHttpsEnabled.(bool)
		result.IsHttpsEnabled = &tmp
	}

	if isHttpsForced, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_https_forced")); ok {
		tmp := isHttpsForced.(bool)
		result.IsHttpsForced = &tmp
	}

	return result, nil
}

func PolicyConfigToMap(obj *oci_waas.PolicyConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CertificateId != nil {
		result["certificate_id"] = string(*obj.CertificateId)
	}

	if obj.IsHttpsEnabled != nil {
		result["is_https_enabled"] = bool(*obj.IsHttpsEnabled)
	}

	if obj.IsHttpsForced != nil {
		result["is_https_forced"] = bool(*obj.IsHttpsForced)
	}

	return result
}

func CustomHeaderToMap(obj oci_waas.Header) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func OriginToMap(obj oci_waas.Origin) map[string]interface{} {
	result := map[string]interface{}{}

	customHeaders := []interface{}{}
	for _, item := range obj.CustomHeaders {
		customHeaders = append(customHeaders, CustomHeaderToMap(item))
	}
	result["custom_headers"] = customHeaders

	if obj.Uri != nil {
		result["uri"] = string(*obj.Uri)
	}

	if obj.HttpPort != nil {
		result["http_port"] = int(*obj.HttpPort)
	}

	if obj.HttpsPort != nil {
		result["https_port"] = int(*obj.HttpsPort)
	}

	return result
}

func OriginMapToMap(originMap map[string]oci_waas.Origin) []interface{} {
	origins := []interface{}{}

	// This is because we model the API's map as a List for Terraform convenience
	for label, origin := range originMap {
		originResultMap := OriginToMap(origin)
		originResultMap["label"] = label
		origins = append(origins, originResultMap)
	}

	return origins
}

func (s *WaasWaasPolicyResourceCrud) mapToProtectionSettings(fieldKeyFormat string) (oci_waas.ProtectionSettings, error) {
	result := oci_waas.ProtectionSettings{}

	result.AllowedHttpMethods = []oci_waas.ProtectionSettingsAllowedHttpMethodsEnum{}
	if allowedHttpMethods, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_http_methods")); ok {
		interfaces := allowedHttpMethods.([]interface{})
		tmp := make([]oci_waas.ProtectionSettingsAllowedHttpMethodsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_waas.ProtectionSettingsAllowedHttpMethodsEnum(interfaces[i].(string))
			}
		}
		result.AllowedHttpMethods = tmp
	}

	if blockAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_action")); ok {
		result.BlockAction = oci_waas.ProtectionSettingsBlockActionEnum(blockAction.(string))
	}

	if blockErrorPageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_error_page_code")); ok {
		tmp := blockErrorPageCode.(string)
		result.BlockErrorPageCode = &tmp
	}

	if blockErrorPageDescription, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_error_page_description")); ok {
		tmp := blockErrorPageDescription.(string)
		result.BlockErrorPageDescription = &tmp
	}

	if blockErrorPageMessage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_error_page_message")); ok {
		tmp := blockErrorPageMessage.(string)
		result.BlockErrorPageMessage = &tmp
	}

	if blockResponseCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_response_code")); ok {
		tmp := blockResponseCode.(int)
		result.BlockResponseCode = &tmp
	}

	if isResponseInspected, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_response_inspected")); ok {
		tmp := isResponseInspected.(bool)
		result.IsResponseInspected = &tmp
	}

	if maxArgumentCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_argument_count")); ok {
		tmp := maxArgumentCount.(int)
		result.MaxArgumentCount = &tmp
	}

	if maxNameLengthPerArgument, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_name_length_per_argument")); ok {
		tmp := maxNameLengthPerArgument.(int)
		result.MaxNameLengthPerArgument = &tmp
	}

	if maxResponseSizeInKiB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_response_size_in_ki_b")); ok {
		tmp := maxResponseSizeInKiB.(int)
		result.MaxResponseSizeInKiB = &tmp
	}

	if maxTotalNameLengthOfArguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_total_name_length_of_arguments")); ok {
		tmp := maxTotalNameLengthOfArguments.(int)
		result.MaxTotalNameLengthOfArguments = &tmp
	}

	result.MediaTypes = []string{}
	if mediaTypes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "media_types")); ok {
		interfaces := mediaTypes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		result.MediaTypes = tmp
	}

	if recommendationsPeriodInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recommendations_period_in_days")); ok {
		tmp := recommendationsPeriodInDays.(int)
		result.RecommendationsPeriodInDays = &tmp
	}

	return result, nil
}

func ProtectionSettingsToMap(obj *oci_waas.ProtectionSettings) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_http_methods"] = obj.AllowedHttpMethods

	result["block_action"] = string(obj.BlockAction)

	if obj.BlockErrorPageCode != nil {
		result["block_error_page_code"] = string(*obj.BlockErrorPageCode)
	}

	if obj.BlockErrorPageDescription != nil {
		result["block_error_page_description"] = string(*obj.BlockErrorPageDescription)
	}

	if obj.BlockErrorPageMessage != nil {
		result["block_error_page_message"] = string(*obj.BlockErrorPageMessage)
	}

	if obj.BlockResponseCode != nil {
		result["block_response_code"] = int(*obj.BlockResponseCode)
	}

	if obj.IsResponseInspected != nil {
		result["is_response_inspected"] = bool(*obj.IsResponseInspected)
	}

	if obj.MaxArgumentCount != nil {
		result["max_argument_count"] = int(*obj.MaxArgumentCount)
	}

	if obj.MaxNameLengthPerArgument != nil {
		result["max_name_length_per_argument"] = int(*obj.MaxNameLengthPerArgument)
	}

	if obj.MaxResponseSizeInKiB != nil {
		result["max_response_size_in_ki_b"] = int(*obj.MaxResponseSizeInKiB)
	}

	if obj.MaxTotalNameLengthOfArguments != nil {
		result["max_total_name_length_of_arguments"] = int(*obj.MaxTotalNameLengthOfArguments)
	}

	result["media_types"] = obj.MediaTypes

	if obj.RecommendationsPeriodInDays != nil {
		result["recommendations_period_in_days"] = int(*obj.RecommendationsPeriodInDays)
	}

	return result
}

func (s *WaasWaasPolicyResourceCrud) mapToWafConfigDetails(fieldKeyFormat string) (oci_waas.WafConfigDetails, error) {
	result := oci_waas.WafConfigDetails{}

	result.AccessRules = []oci_waas.AccessRule{}
	if accessRules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_rules")); ok {
		interfaces := accessRules.([]interface{})
		tmp := make([]oci_waas.AccessRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "access_rules"), stateDataIndex)
			converted, err := s.mapToAccessRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		result.AccessRules = tmp
	}

	if addressRateLimiting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address_rate_limiting")); ok {
		if tmpList := addressRateLimiting.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "address_rate_limiting"), 0)
			tmp, err := s.mapToAddressRateLimiting(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert address_rate_limiting, encountered error: %v", err)
			}
			result.AddressRateLimiting = &tmp
		}
	}

	result.Captchas = []oci_waas.Captcha{}
	if captchas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "captchas")); ok {
		interfaces := captchas.([]interface{})
		tmp := make([]oci_waas.Captcha, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "captchas"), stateDataIndex)
			converted, err := s.mapToCaptcha(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		result.Captchas = tmp
	}

	if deviceFingerprintChallenge, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "device_fingerprint_challenge")); ok {
		if tmpList := deviceFingerprintChallenge.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "device_fingerprint_challenge"), 0)
			tmp, err := s.mapToDeviceFingerprintChallenge(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert device_fingerprint_challenge, encountered error: %v", err)
			}
			result.DeviceFingerprintChallenge = &tmp
		}
	}

	if humanInteractionChallenge, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "human_interaction_challenge")); ok {
		if tmpList := humanInteractionChallenge.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "human_interaction_challenge"), 0)
			tmp, err := s.mapToHumanInteractionChallenge(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert human_interaction_challenge, encountered error: %v", err)
			}
			result.HumanInteractionChallenge = &tmp
		}
	}

	if jsChallenge, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "js_challenge")); ok {
		if tmpList := jsChallenge.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "js_challenge"), 0)
			tmp, err := s.mapToJsChallenge(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert js_challenge, encountered error: %v", err)
			}
			result.JsChallenge = &tmp
		}
	}

	if origin, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "origin")); ok {
		tmp := origin.(string)
		result.Origin = &tmp
	}

	if protectionSettings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_settings")); ok {
		if tmpList := protectionSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "protection_settings"), 0)
			tmp, err := s.mapToProtectionSettings(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert protection_settings, encountered error: %v", err)
			}
			result.ProtectionSettings = &tmp
		}
	}

	result.Whitelists = []oci_waas.Whitelist{}
	if whitelists, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "whitelists")); ok {
		interfaces := whitelists.([]interface{})
		tmp := make([]oci_waas.Whitelist, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "whitelists"), stateDataIndex)
			converted, err := s.mapToWhitelist(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		result.Whitelists = tmp
	}

	return result, nil
}

func (s *WaasWaasPolicyResourceCrud) mapToWafConfig(fieldKeyFormat string) (oci_waas.WafConfig, error) {
	result := oci_waas.WafConfig{}

	result.AccessRules = []oci_waas.AccessRule{}
	if accessRules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_rules")); ok {
		interfaces := accessRules.([]interface{})
		tmp := make([]oci_waas.AccessRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "access_rules"), stateDataIndex)
			converted, err := s.mapToAccessRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		result.AccessRules = tmp
	}

	if addressRateLimiting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address_rate_limiting")); ok {
		if tmpList := addressRateLimiting.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "address_rate_limiting"), 0)
			tmp, err := s.mapToAddressRateLimiting(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert address_rate_limiting, encountered error: %v", err)
			}
			result.AddressRateLimiting = &tmp
		}
	}

	result.Captchas = []oci_waas.Captcha{}
	if captchas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "captchas")); ok {
		interfaces := captchas.([]interface{})
		tmp := make([]oci_waas.Captcha, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "captchas"), stateDataIndex)
			converted, err := s.mapToCaptcha(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		result.Captchas = tmp
	}

	if deviceFingerprintChallenge, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "device_fingerprint_challenge")); ok {
		if tmpList := deviceFingerprintChallenge.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "device_fingerprint_challenge"), 0)
			tmp, err := s.mapToDeviceFingerprintChallenge(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert device_fingerprint_challenge, encountered error: %v", err)
			}
			result.DeviceFingerprintChallenge = &tmp
		}
	}

	if humanInteractionChallenge, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "human_interaction_challenge")); ok {
		if tmpList := humanInteractionChallenge.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "human_interaction_challenge"), 0)
			tmp, err := s.mapToHumanInteractionChallenge(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert human_interaction_challenge, encountered error: %v", err)
			}
			result.HumanInteractionChallenge = &tmp
		}
	}

	if jsChallenge, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "js_challenge")); ok {
		if tmpList := jsChallenge.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "js_challenge"), 0)
			tmp, err := s.mapToJsChallenge(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert js_challenge, encountered error: %v", err)
			}
			result.JsChallenge = &tmp
		}
	}

	if origin, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "origin")); ok {
		tmp := origin.(string)
		result.Origin = &tmp
	}

	if protectionSettings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_settings")); ok {
		if tmpList := protectionSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "protection_settings"), 0)
			tmp, err := s.mapToProtectionSettings(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert protection_settings, encountered error: %v", err)
			}
			result.ProtectionSettings = &tmp
		}
	}

	result.Whitelists = []oci_waas.Whitelist{}
	if whitelists, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "whitelists")); ok {
		interfaces := whitelists.([]interface{})
		tmp := make([]oci_waas.Whitelist, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "whitelists"), stateDataIndex)
			converted, err := s.mapToWhitelist(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		result.Whitelists = tmp
	}

	return result, nil
}

func WafConfigToMap(obj *oci_waas.WafConfig) map[string]interface{} {
	result := map[string]interface{}{}

	accessRules := []interface{}{}
	for _, item := range obj.AccessRules {
		accessRules = append(accessRules, AccessRuleToMap(item))
	}
	result["access_rules"] = accessRules

	if obj.AddressRateLimiting != nil {
		result["address_rate_limiting"] = []interface{}{AddressRateLimitingToMap(obj.AddressRateLimiting)}
	}

	captchas := []interface{}{}
	for _, item := range obj.Captchas {
		captchas = append(captchas, CaptchaToMap(item))
	}
	result["captchas"] = captchas

	if obj.DeviceFingerprintChallenge != nil {
		result["device_fingerprint_challenge"] = []interface{}{DeviceFingerprintChallengeToMap(obj.DeviceFingerprintChallenge)}
	}

	if obj.HumanInteractionChallenge != nil {
		result["human_interaction_challenge"] = []interface{}{HumanInteractionChallengeToMap(obj.HumanInteractionChallenge)}
	}

	if obj.JsChallenge != nil {
		result["js_challenge"] = []interface{}{JsChallengeToMap(obj.JsChallenge)}
	}

	if obj.Origin != nil {
		result["origin"] = string(*obj.Origin)
	}

	if obj.ProtectionSettings != nil {
		result["protection_settings"] = []interface{}{ProtectionSettingsToMap(obj.ProtectionSettings)}
	}

	whitelists := []interface{}{}
	for _, item := range obj.Whitelists {
		whitelists = append(whitelists, WhitelistToMap(item))
	}
	result["whitelists"] = whitelists

	return result
}

func (s *WaasWaasPolicyResourceCrud) mapToWhitelist(fieldKeyFormat string) (oci_waas.Whitelist, error) {
	result := oci_waas.Whitelist{}

	result.Addresses = []string{}
	if addresses, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "addresses")); ok {
		interfaces := addresses.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		result.Addresses = tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func WhitelistToMap(obj oci_waas.Whitelist) map[string]interface{} {
	result := map[string]interface{}{}

	result["addresses"] = obj.Addresses

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *WaasWaasPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_waas.ChangeWaasPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.WaasPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	_, err := s.Client.ChangeWaasPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func originsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if label, ok := m["label"]; ok && label != "" {
		buf.WriteString(fmt.Sprintf("%v-", label))
	}
	if uri, ok := m["uri"]; ok && uri != "" {
		buf.WriteString(fmt.Sprintf("%v-", uri))
	}
	if httpPort, ok := m["http_port"]; ok && httpPort != "" {
		buf.WriteString(fmt.Sprintf("%v-", httpPort))
	}
	if httpsPort, ok := m["https_port"]; ok && httpsPort != "" {
		buf.WriteString(fmt.Sprintf("%v-", httpsPort))
	}
	if customHeaders, ok := m["custom_headers"]; ok {
		if tmpList := customHeaders.([]interface{}); len(tmpList) > 0 {
			buf.WriteString("custom_headers-")
			for _, customHeadersRaw := range tmpList {
				tmpMap := customHeadersRaw.(map[string]interface{})
				if name, ok := tmpMap["name"]; ok {
					buf.WriteString(fmt.Sprintf("%v-", name))
				}
				if value, ok := tmpMap["value"]; ok {
					buf.WriteString(fmt.Sprintf("%v-", value))
				}
			}
		}
	}
	return hashcode.String(buf.String())
}
