// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	oci_waas "github.com/oracle/oci-go-sdk/v56/waas"
)

func WaasWaasPolicyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularWaasWaasPolicy,
		Schema: map[string]*schema.Schema{
			"waas_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"additional_domains": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"origin_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"origin_group": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"origin": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"weight": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"origins": {
				Type:     schema.TypeList,
				Computed: true,
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
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"certificate_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cipher_group": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_address_header": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"health_checks": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"expected_response_code_group": {
										Type:             schema.TypeList,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"expected_response_text": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"headers": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										Elem:             schema.TypeString,
									},
									"healthy_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"interval_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"is_response_text_check_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"method": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"timeout_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"unhealthy_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"is_behind_cdn": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_cache_control_respected": {
							Type:     schema.TypeBool,
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
						"is_origin_compression_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_response_buffering_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_sni_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"load_balancing_method": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"method": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"IP_HASH",
											"ROUND_ROBIN",
											"STICKY_COOKIE",
										}, true),
									},

									// Optional
									"domain": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"expiration_time_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"tls_protocols": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"websocket_path_prefixes": {
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
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"waf_config": {
				Type:     schema.TypeList,
				Computed: true,
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
												"is_case_sensitive": {
													Type:     schema.TypeBool,
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
									"bypass_challenges": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
									"redirect_response_code": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"redirect_url": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"response_header_manipulation": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"action": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"ADD_HTTP_RESPONSE_HEADER",
														"EXTEND_HTTP_RESPONSE_HEADER",
														"REMOVE_HTTP_RESPONSE_HEADER",
													}, true),
												},
												"header": {
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
						"caching_rules": {
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
									"caching_duration": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"client_caching_duration": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"is_client_caching_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
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
						"custom_protection_rules": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"action": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"exclusions": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"exclusions": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"target": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"id": {
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
									"is_nat_enabled": {
										Type:     schema.TypeBool,
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
									"are_redirects_challenged": {
										Type:     schema.TypeBool,
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
									"criteria": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
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
												"is_case_sensitive": {
													Type:     schema.TypeBool,
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
									"is_nat_enabled": {
										Type:     schema.TypeBool,
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
						"origin_groups": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"address_lists": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"addresses": {
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
		},
	}
}

func readSingularWaasWaasPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WaasWaasPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.ReadResource(sync)
}

type WaasWaasPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.WaasClient
	Res    *oci_waas.GetWaasPolicyResponse
}

func (s *WaasWaasPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasWaasPolicyDataSourceCrud) Get() error {
	request := oci_waas.GetWaasPolicyRequest{}

	if waasPolicyId, ok := s.D.GetOkExists("waas_policy_id"); ok {
		tmp := waasPolicyId.(string)
		request.WaasPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waas")

	response, err := s.Client.GetWaasPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WaasWaasPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_domains", s.Res.AdditionalDomains)

	if s.Res.Cname != nil {
		s.D.Set("cname", *s.Res.Cname)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.OriginGroups != nil {
		s.D.Set("origin_groups", OriginGroupMapToMap(s.Res.OriginGroups))
	} else {
		s.D.Set("origin_groups", nil)
	}

	if s.Res.Origins != nil {
		s.D.Set("origins", OriginMapToMap(s.Res.Origins))
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
