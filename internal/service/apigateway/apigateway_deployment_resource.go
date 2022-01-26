// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_apigateway "github.com/oracle/oci-go-sdk/v56/apigateway"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func ApigatewayDeploymentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
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

			// Optional
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
			"specification": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
												"type": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"CUSTOM_AUTHENTICATION",
														"JWT_AUTHENTICATION",
													}, true),
												},

												// Optional
												"audiences": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"function_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"is_anonymous_access_allowed": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"issuers": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"max_clock_skew_in_seconds": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"public_keys": {
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
																	"REMOTE_JWKS",
																	"STATIC_KEYS",
																}, true),
															},

															// Optional
															"is_ssl_verify_disabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"keys": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"format": {
																			Type:             schema.TypeString,
																			Required:         true,
																			DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																			ValidateFunc: validation.StringInSlice([]string{
																				"JSON_WEB_KEY",
																				"PEM",
																			}, true),
																		},

																		// Optional
																		"alg": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"e": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"key": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"key_ops": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Computed: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"kid": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"kty": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"n": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"use": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},

																		// Computed
																	},
																},
															},
															"max_cache_duration_in_hours": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"uri": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"token_auth_scheme": {
													Type:     schema.TypeString,
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
												"verify_claims": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"is_required": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"values": {
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
									"mutual_tls": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"allowed_sans": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"is_verified_certificate_required": {
													Type:     schema.TypeBool,
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
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
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
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
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
												"body_validation": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															// Optional
															"content": {
																Type:     schema.TypeSet,
																Optional: true,
																Computed: true,
																Set:      mediaTypeHashCodeForBodyValidationContentSets,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"media_type": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"validation_type": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},
															"required": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"validation_mode": {
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
												"header_transformations": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"filter_headers": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"items": {
																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},

																					// Optional

																					// Computed
																				},
																			},
																		},
																		"type": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		// Optional

																		// Computed
																	},
																},
															},
															"rename_headers": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"items": {
																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required
																					"from": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"to": {
																						Type:     schema.TypeString,
																						Required: true,
																					},

																					// Optional

																					// Computed
																				},
																			},
																		},

																		// Optional

																		// Computed
																	},
																},
															},
															"set_headers": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"items": {
																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"values": {
																						Type:     schema.TypeList,
																						Required: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					// Optional
																					"if_exists": {
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

															// Computed
														},
													},
												},
												"header_validations": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"headers": {
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
																		"required": {
																			Type:     schema.TypeBool,
																			Optional: true,
																			Computed: true,
																		},

																		// Computed
																	},
																},
															},
															"validation_mode": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"query_parameter_transformations": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"filter_query_parameters": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"items": {
																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},

																					// Optional

																					// Computed
																				},
																			},
																		},
																		"type": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		// Optional

																		// Computed
																	},
																},
															},
															"rename_query_parameters": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"items": {
																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required
																					"from": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"to": {
																						Type:     schema.TypeString,
																						Required: true,
																					},

																					// Optional

																					// Computed
																				},
																			},
																		},

																		// Optional

																		// Computed
																	},
																},
															},
															"set_query_parameters": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"items": {
																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"values": {
																						Type:     schema.TypeList,
																						Required: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					// Optional
																					"if_exists": {
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

															// Computed
														},
													},
												},
												"query_parameter_validations": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"parameters": {
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
																		"required": {
																			Type:     schema.TypeBool,
																			Optional: true,
																			Computed: true,
																		},

																		// Computed
																	},
																},
															},
															"validation_mode": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"response_cache_lookup": {
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
																	"SIMPLE_LOOKUP_POLICY",
																}, true),
															},

															// Optional
															"cache_key_additions": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"is_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"is_private_caching_enabled": {
																Type:     schema.TypeBool,
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
									"response_policies": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"header_transformations": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"filter_headers": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"items": {
																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},

																					// Optional

																					// Computed
																				},
																			},
																		},
																		"type": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		// Optional

																		// Computed
																	},
																},
															},
															"rename_headers": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"items": {
																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required
																					"from": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"to": {
																						Type:     schema.TypeString,
																						Required: true,
																					},

																					// Optional

																					// Computed
																				},
																			},
																		},

																		// Optional

																		// Computed
																	},
																},
															},
															"set_headers": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"items": {
																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"values": {
																						Type:     schema.TypeList,
																						Required: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					// Optional
																					"if_exists": {
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

															// Computed
														},
													},
												},
												"response_cache_store": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"time_to_live_in_seconds": {
																Type:     schema.TypeInt,
																Required: true,
															},
															"type": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"FIXED_TTL_STORE_POLICY",
																}, true),
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
								},
							},
						},

						// Computed
					},
				},
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
	sync.Client = m.(*client.OracleClients).DeploymentClient()
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.CreateResource(d, sync)
}

func readApigatewayDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DeploymentClient()

	return tfresource.ReadResource(sync)
}

func updateApigatewayDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DeploymentClient()
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApigatewayDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DeploymentClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.DeleteResource(d, sync)
}

type ApigatewayDeploymentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apigateway.DeploymentClient
	Res                    *oci_apigateway.Deployment
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_apigateway.WorkRequestsClient
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

func (s *ApigatewayDeploymentResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_apigateway.DeploymentLifecycleStateUpdating),
	}
}

func (s *ApigatewayDeploymentResourceCrud) UpdatedTarget() []string {
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.CreateDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeploymentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ApigatewayDeploymentResourceCrud) getDeploymentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_apigateway.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	deploymentId, err := deploymentWaitForWorkRequest(workId, "deployment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, deploymentId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(context.Background(),
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
		if tfresource.ShouldRetry(response, false, "apigateway", startTime) {
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
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "apigateway")
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

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_apigateway.WorkRequestStatusFailed || response.Status == oci_apigateway.WorkRequestStatusCanceled {
		return nil, getErrorFromApigatewayDeploymentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromApigatewayDeploymentWorkRequest(client *oci_apigateway.WorkRequestsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_apigateway.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_apigateway.ListWorkRequestErrorsRequest{
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

func (s *ApigatewayDeploymentResourceCrud) Get() error {
	request := oci_apigateway.GetDeploymentRequest{}

	tmp := s.D.Id()
	request.DeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

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
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.UpdateDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeploymentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ApigatewayDeploymentResourceCrud) Delete() error {
	request := oci_apigateway.DeleteDeploymentRequest{}

	tmp := s.D.Id()
	request.DeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.DeleteDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := deploymentWaitForWorkRequest(workId, "deployment",
		oci_apigateway.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *ApigatewayDeploymentResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
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
		s.D.Set("specification", []interface{}{ApiSpecificationToMap(s.Res.Specification, false)})
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

func ApiSpecificationToMap(obj *oci_apigateway.ApiSpecification, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LoggingPolicies != nil {
		result["logging_policies"] = []interface{}{ApiSpecificationLoggingPoliciesToMap(obj.LoggingPolicies)}
	}

	if obj.RequestPolicies != nil {
		result["request_policies"] = []interface{}{ApiSpecificationRequestPoliciesToMap(obj.RequestPolicies)}
	}

	routes := []interface{}{}
	for _, item := range obj.Routes {
		routes = append(routes, ApiSpecificationRouteToMap(item, datasource))
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

	if mutualTls, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mutual_tls")); ok {
		if tmpList := mutualTls.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "mutual_tls"), 0)
			tmp, err := s.mapToMutualTlsDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert mutual_tls, encountered error: %v", err)
			}
			result.MutualTls = &tmp
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

	if obj.MutualTls != nil {
		result["mutual_tls"] = []interface{}{MutualTlsDetailsToMap(obj.MutualTls)}
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

	if responsePolicies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_policies")); ok {
		if tmpList := responsePolicies.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "response_policies"), 0)
			tmp, err := s.mapToApiSpecificationRouteResponsePolicies(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert response_policies, encountered error: %v", err)
			}
			result.ResponsePolicies = &tmp
		}
	}

	return result, nil
}

func ApiSpecificationRouteToMap(obj oci_apigateway.ApiSpecificationRoute, datasource bool) map[string]interface{} {
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
		result["request_policies"] = []interface{}{ApiSpecificationRouteRequestPoliciesToMap(obj.RequestPolicies, datasource)}
	}

	if obj.ResponsePolicies != nil {
		result["response_policies"] = []interface{}{ApiSpecificationRouteResponsePoliciesToMap(obj.ResponsePolicies)}
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

	if bodyValidation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "body_validation")); ok {
		if tmpList := bodyValidation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "body_validation"), 0)
			tmp, err := s.mapToBodyValidationRequestPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert body_validation, encountered error: %v", err)
			}
			result.BodyValidation = &tmp
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

	if headerTransformations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header_transformations")); ok {
		if tmpList := headerTransformations.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "header_transformations"), 0)
			tmp, err := s.mapToHeaderTransformationPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert header_transformations, encountered error: %v", err)
			}
			result.HeaderTransformations = &tmp
		}
	}

	if headerValidations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header_validations")); ok {
		if tmpList := headerValidations.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "header_validations"), 0)
			tmp, err := s.mapToHeaderValidationRequestPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert header_validations, encountered error: %v", err)
			}
			result.HeaderValidations = &tmp
		}
	}

	if queryParameterTransformations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query_parameter_transformations")); ok {
		if tmpList := queryParameterTransformations.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "query_parameter_transformations"), 0)
			tmp, err := s.mapToQueryParameterTransformationPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert query_parameter_transformations, encountered error: %v", err)
			}
			result.QueryParameterTransformations = &tmp
		}
	}

	if queryParameterValidations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query_parameter_validations")); ok {
		if tmpList := queryParameterValidations.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "query_parameter_validations"), 0)
			tmp, err := s.mapToQueryParameterValidationRequestPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert query_parameter_validations, encountered error: %v", err)
			}
			result.QueryParameterValidations = &tmp
		}
	}

	if responseCacheLookup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_cache_lookup")); ok {
		if tmpList := responseCacheLookup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "response_cache_lookup"), 0)
			tmp, err := s.mapToResponseCacheLookupPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert response_cache_lookup, encountered error: %v", err)
			}
			result.ResponseCacheLookup = tmp
		}
	}

	return result, nil
}

func ApiSpecificationRouteRequestPoliciesToMap(obj *oci_apigateway.ApiSpecificationRouteRequestPolicies, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Authorization != nil {
		authorizationArray := []interface{}{}
		if authorizationMap := RouteAuthorizationPolicyToMap(&obj.Authorization); authorizationMap != nil {
			authorizationArray = append(authorizationArray, authorizationMap)
		}
		result["authorization"] = authorizationArray
	}

	if obj.BodyValidation != nil {
		result["body_validation"] = []interface{}{BodyValidationRequestPolicyToMap(obj.BodyValidation, datasource)}
	}

	if obj.Cors != nil {
		result["cors"] = []interface{}{CorsPolicyToMap(obj.Cors)}
	}

	if obj.HeaderTransformations != nil {
		result["header_transformations"] = []interface{}{HeaderTransformationPolicyToMap(obj.HeaderTransformations)}
	}

	if obj.HeaderValidations != nil {
		result["header_validations"] = []interface{}{HeaderValidationRequestPolicyToMap(obj.HeaderValidations)}
	}

	if obj.QueryParameterTransformations != nil {
		result["query_parameter_transformations"] = []interface{}{QueryParameterTransformationPolicyToMap(obj.QueryParameterTransformations)}
	}

	if obj.QueryParameterValidations != nil {
		result["query_parameter_validations"] = []interface{}{QueryParameterValidationRequestPolicyToMap(obj.QueryParameterValidations)}
	}

	if obj.ResponseCacheLookup != nil {
		responseCacheLookupArray := []interface{}{}
		if responseCacheLookupMap := ResponseCacheLookupPolicyToMap(&obj.ResponseCacheLookup); responseCacheLookupMap != nil {
			responseCacheLookupArray = append(responseCacheLookupArray, responseCacheLookupMap)
		}
		result["response_cache_lookup"] = responseCacheLookupArray
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToApiSpecificationRouteResponsePolicies(fieldKeyFormat string) (oci_apigateway.ApiSpecificationRouteResponsePolicies, error) {
	result := oci_apigateway.ApiSpecificationRouteResponsePolicies{}

	if headerTransformations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header_transformations")); ok {
		if tmpList := headerTransformations.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "header_transformations"), 0)
			tmp, err := s.mapToHeaderTransformationPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert header_transformations, encountered error: %v", err)
			}
			result.HeaderTransformations = &tmp
		}
	}

	if responseCacheStore, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_cache_store")); ok {
		if tmpList := responseCacheStore.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "response_cache_store"), 0)
			tmp, err := s.mapToResponseCacheStorePolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert response_cache_store, encountered error: %v", err)
			}
			result.ResponseCacheStore = tmp
		}
	}

	return result, nil
}

func ApiSpecificationRouteResponsePoliciesToMap(obj *oci_apigateway.ApiSpecificationRouteResponsePolicies) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HeaderTransformations != nil {
		result["header_transformations"] = []interface{}{HeaderTransformationPolicyToMap(obj.HeaderTransformations)}
	}

	if obj.ResponseCacheStore != nil {
		responseCacheStoreArray := []interface{}{}
		if responseCacheStoreMap := ResponseCacheStorePolicyToMap(&obj.ResponseCacheStore); responseCacheStoreMap != nil {
			responseCacheStoreArray = append(responseCacheStoreArray, responseCacheStoreMap)
		}
		result["response_cache_store"] = responseCacheStoreArray
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
	case strings.ToLower("JWT_AUTHENTICATION"):
		details := oci_apigateway.JwtAuthenticationPolicy{}
		if audiences, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "audiences")); ok {
			interfaces := audiences.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "audiences")) {
				details.Audiences = tmp
			}
		}
		if issuers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "issuers")); ok {
			interfaces := issuers.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "issuers")) {
				details.Issuers = tmp
			}
		}
		if maxClockSkewInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_clock_skew_in_seconds")); ok {
			tmp := float32(maxClockSkewInSeconds.(float64))
			details.MaxClockSkewInSeconds = &tmp
		}
		if publicKeys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "public_keys")); ok {
			if tmpList := publicKeys.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "public_keys"), 0)
				tmp, err := s.mapToPublicKeySet(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert public_keys, encountered error: %v", err)
				}
				details.PublicKeys = tmp
			}
		}
		if tokenAuthScheme, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "token_auth_scheme")); ok {
			tmp := tokenAuthScheme.(string)
			details.TokenAuthScheme = &tmp
		}
		if tokenHeader, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "token_header")); ok {
			tmp := tokenHeader.(string)
			details.TokenHeader = &tmp
		}
		if tokenQueryParam, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "token_query_param")); ok {
			tmp := tokenQueryParam.(string)
			if len(tmp) != 0 {
				details.TokenQueryParam = &tmp
			}
		}
		if verifyClaims, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_claims")); ok {
			interfaces := verifyClaims.([]interface{})
			tmp := make([]oci_apigateway.JsonWebTokenClaim, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "verify_claims"), stateDataIndex)
				converted, err := s.mapToJsonWebTokenClaim(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "verify_claims")) {
				details.VerifyClaims = tmp
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
	case oci_apigateway.JwtAuthenticationPolicy:
		result["type"] = "JWT_AUTHENTICATION"

		result["audiences"] = v.Audiences

		result["issuers"] = v.Issuers

		if v.MaxClockSkewInSeconds != nil {
			result["max_clock_skew_in_seconds"] = float32(*v.MaxClockSkewInSeconds)
		}

		if v.PublicKeys != nil {
			publicKeysArray := []interface{}{}
			if publicKeysMap := PublicKeySetToMap(&v.PublicKeys); publicKeysMap != nil {
				publicKeysArray = append(publicKeysArray, publicKeysMap)
			}
			result["public_keys"] = publicKeysArray
		}

		if v.TokenAuthScheme != nil {
			result["token_auth_scheme"] = string(*v.TokenAuthScheme)
		}

		if v.TokenHeader != nil {
			result["token_header"] = string(*v.TokenHeader)
		}

		if v.TokenQueryParam != nil {
			result["token_query_param"] = string(*v.TokenQueryParam)
		}

		verifyClaims := []interface{}{}
		for _, item := range v.VerifyClaims {
			verifyClaims = append(verifyClaims, JsonWebTokenClaimToMap(item))
		}
		result["verify_claims"] = verifyClaims

		if v.IsAnonymousAccessAllowed != nil {
			result["is_anonymous_access_allowed"] = bool(*v.IsAnonymousAccessAllowed)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToContentValidation(fieldKeyFormat string) (oci_apigateway.ContentValidation, error) {
	var baseObject oci_apigateway.ContentValidation

	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validation_type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}

	switch strings.ToLower(type_) {
	case strings.ToLower("NONE"):
		validation := oci_apigateway.NoContentValidation{}
		baseObject = validation
	default:
		return nil, fmt.Errorf("unknown validation type '%v' was specified", type_)
	}

	return baseObject, nil
}

func (s *ApigatewayDeploymentResourceCrud) mapToBodyValidationRequestPolicy(fieldKeyFormat string) (oci_apigateway.BodyValidationRequestPolicy, error) {
	result := oci_apigateway.BodyValidationRequestPolicy{}

	if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
		set := content.(*schema.Set)
		interfaces := set.List()

		content := map[string]oci_apigateway.ContentValidation{}
		for _, mediaTypeObject := range interfaces {
			hash := mediaTypeHashCodeForBodyValidationContentSets(mediaTypeObject)
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "content"), hash)

			mediaTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "media_type"))
			var mediaType string
			if ok {
				mediaType = mediaTypeRaw.(string)
			} else {
				return result, fmt.Errorf("unable to convert media_type")
			}

			converted, err := s.mapToContentValidation(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}

			if _, ok = content[mediaType]; ok {
				return result, fmt.Errorf("media_type shadows a previous content media_type declaration. A request body validation policy may only contain unique media_types")
			}

			content[mediaType] = converted
		}

		result.Content = content
	} else {
		content := map[string]oci_apigateway.ContentValidation{}
		result.Content = content
	}

	if required, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "required")); ok {
		tmp := required.(bool)
		result.Required = &tmp
	}

	if validationMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validation_mode")); ok {
		result.ValidationMode = oci_apigateway.BodyValidationRequestPolicyValidationModeEnum(validationMode.(string))
	}

	return result, nil
}

func ContentValidationToMap(obj *oci_apigateway.ContentValidation) map[string]interface{} {
	result := map[string]interface{}{}

	switch (*obj).(type) {
	case oci_apigateway.NoContentValidation:
		result["validation_type"] = "NONE"
	default:
		log.Printf("[WARN] Received 'validation_type' of unknown type %v", *obj)
	}

	return result
}

func BodyValidationRequestPolicyToMap(obj *oci_apigateway.BodyValidationRequestPolicy, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	contentArray := []interface{}{}
	for mediaType, mediaTypeObject := range obj.Content {
		contentMap := ContentValidationToMap(&mediaTypeObject)
		contentMap["media_type"] = mediaType
		contentArray = append(contentArray, contentMap)
	}

	if datasource {
		result["content"] = contentArray
	} else {
		result["content"] = schema.NewSet(mediaTypeHashCodeForBodyValidationContentSets, contentArray)
	}

	if obj.Required != nil {
		result["required"] = bool(*obj.Required)
	}

	result["validation_mode"] = string(obj.ValidationMode)

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
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
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

func (s *ApigatewayDeploymentResourceCrud) mapToFilterHeaderPolicy(fieldKeyFormat string) (oci_apigateway.FilterHeaderPolicy, error) {
	result := oci_apigateway.FilterHeaderPolicy{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_apigateway.FilterHeaderPolicyItem, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToFilterHeaderPolicyItem(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_apigateway.FilterHeaderPolicyTypeEnum(type_.(string))
	}

	return result, nil
}

func FilterHeaderPolicyToMap(obj *oci_apigateway.FilterHeaderPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, FilterHeaderPolicyItemToMap(item))
	}
	result["items"] = items

	result["type"] = string(obj.Type)

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToFilterHeaderPolicyItem(fieldKeyFormat string) (oci_apigateway.FilterHeaderPolicyItem, error) {
	result := oci_apigateway.FilterHeaderPolicyItem{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func FilterHeaderPolicyItemToMap(obj oci_apigateway.FilterHeaderPolicyItem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToFilterQueryParameterPolicy(fieldKeyFormat string) (oci_apigateway.FilterQueryParameterPolicy, error) {
	result := oci_apigateway.FilterQueryParameterPolicy{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_apigateway.FilterQueryParameterPolicyItem, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToFilterQueryParameterPolicyItem(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_apigateway.FilterQueryParameterPolicyTypeEnum(type_.(string))
	}

	return result, nil
}

func FilterQueryParameterPolicyToMap(obj *oci_apigateway.FilterQueryParameterPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, FilterQueryParameterPolicyItemToMap(item))
	}
	result["items"] = items

	result["type"] = string(obj.Type)

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToFilterQueryParameterPolicyItem(fieldKeyFormat string) (oci_apigateway.FilterQueryParameterPolicyItem, error) {
	result := oci_apigateway.FilterQueryParameterPolicyItem{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func FilterQueryParameterPolicyItemToMap(obj oci_apigateway.FilterQueryParameterPolicyItem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

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

func (s *ApigatewayDeploymentResourceCrud) mapToHeaderTransformationPolicy(fieldKeyFormat string) (oci_apigateway.HeaderTransformationPolicy, error) {
	result := oci_apigateway.HeaderTransformationPolicy{}

	if filterHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filter_headers")); ok {
		if tmpList := filterHeaders.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "filter_headers"), 0)
			tmp, err := s.mapToFilterHeaderPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert filter_headers, encountered error: %v", err)
			}
			result.FilterHeaders = &tmp
		}
	}

	if renameHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rename_headers")); ok {
		if tmpList := renameHeaders.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rename_headers"), 0)
			tmp, err := s.mapToRenameHeaderPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert rename_headers, encountered error: %v", err)
			}
			result.RenameHeaders = &tmp
		}
	}

	if setHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "set_headers")); ok {
		if tmpList := setHeaders.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "set_headers"), 0)
			tmp, err := s.mapToSetHeaderPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert set_headers, encountered error: %v", err)
			}
			result.SetHeaders = &tmp
		}
	}

	return result, nil
}

func HeaderTransformationPolicyToMap(obj *oci_apigateway.HeaderTransformationPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FilterHeaders != nil {
		result["filter_headers"] = []interface{}{FilterHeaderPolicyToMap(obj.FilterHeaders)}
	}

	if obj.RenameHeaders != nil {
		result["rename_headers"] = []interface{}{RenameHeaderPolicyToMap(obj.RenameHeaders)}
	}

	if obj.SetHeaders != nil {
		result["set_headers"] = []interface{}{SetHeaderPolicyToMap(obj.SetHeaders)}
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToHeaderValidationItem(fieldKeyFormat string) (oci_apigateway.HeaderValidationItem, error) {
	result := oci_apigateway.HeaderValidationItem{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if required, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "required")); ok {
		tmp := required.(bool)
		result.Required = &tmp
	}

	return result, nil
}

func HeaderValidationItemToMap(obj oci_apigateway.HeaderValidationItem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Required != nil {
		result["required"] = bool(*obj.Required)
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToHeaderValidationRequestPolicy(fieldKeyFormat string) (oci_apigateway.HeaderValidationRequestPolicy, error) {
	result := oci_apigateway.HeaderValidationRequestPolicy{}

	if headers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "headers")); ok {
		interfaces := headers.([]interface{})
		tmp := make([]oci_apigateway.HeaderValidationItem, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "headers"), stateDataIndex)
			converted, err := s.mapToHeaderValidationItem(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "headers")) {
			result.Headers = tmp
		}
	}

	if validationMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validation_mode")); ok {
		result.ValidationMode = oci_apigateway.HeaderValidationRequestPolicyValidationModeEnum(validationMode.(string))
	}

	return result, nil
}

func HeaderValidationRequestPolicyToMap(obj *oci_apigateway.HeaderValidationRequestPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	headers := []interface{}{}
	for _, item := range obj.Headers {
		headers = append(headers, HeaderValidationItemToMap(item))
	}
	result["headers"] = headers

	result["validation_mode"] = string(obj.ValidationMode)

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToJsonWebTokenClaim(fieldKeyFormat string) (oci_apigateway.JsonWebTokenClaim, error) {
	result := oci_apigateway.JsonWebTokenClaim{}

	if isRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_required")); ok {
		tmp := isRequired.(bool)
		result.IsRequired = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
		interfaces := values.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
			result.Values = tmp
		}
	}

	return result, nil
}

func JsonWebTokenClaimToMap(obj oci_apigateway.JsonWebTokenClaim) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsRequired != nil {
		result["is_required"] = bool(*obj.IsRequired)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["values"] = obj.Values

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToMutualTlsDetails(fieldKeyFormat string) (oci_apigateway.MutualTlsDetails, error) {
	result := oci_apigateway.MutualTlsDetails{}

	if allowedSans, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_sans")); ok {
		interfaces := allowedSans.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_sans")) {
			result.AllowedSans = tmp
		}
	}

	if isVerifiedCertificateRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_verified_certificate_required")); ok {
		tmp := isVerifiedCertificateRequired.(bool)
		result.IsVerifiedCertificateRequired = &tmp
	}

	return result, nil
}

func MutualTlsDetailsToMap(obj *oci_apigateway.MutualTlsDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_sans"] = obj.AllowedSans

	if obj.IsVerifiedCertificateRequired != nil {
		result["is_verified_certificate_required"] = bool(*obj.IsVerifiedCertificateRequired)
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToPublicKeySet(fieldKeyFormat string) (oci_apigateway.PublicKeySet, error) {
	var baseObject oci_apigateway.PublicKeySet
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("REMOTE_JWKS"):
		details := oci_apigateway.RemoteJsonWebKeySet{}
		if isSslVerifyDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ssl_verify_disabled")); ok {
			tmp := isSslVerifyDisabled.(bool)
			details.IsSslVerifyDisabled = &tmp
		}
		if maxCacheDurationInHours, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_cache_duration_in_hours")); ok {
			tmp := maxCacheDurationInHours.(int)
			details.MaxCacheDurationInHours = &tmp
		}
		if uri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	case strings.ToLower("STATIC_KEYS"):
		details := oci_apigateway.StaticPublicKeySet{}
		if keys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "keys")); ok {
			interfaces := keys.([]interface{})
			tmp := make([]oci_apigateway.StaticPublicKey, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "keys"), stateDataIndex)
				converted, err := s.mapToStaticPublicKey(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "keys")) {
				details.Keys = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func PublicKeySetToMap(obj *oci_apigateway.PublicKeySet) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_apigateway.RemoteJsonWebKeySet:
		result["type"] = "REMOTE_JWKS"

		if v.IsSslVerifyDisabled != nil {
			result["is_ssl_verify_disabled"] = bool(*v.IsSslVerifyDisabled)
		}

		if v.MaxCacheDurationInHours != nil {
			result["max_cache_duration_in_hours"] = int(*v.MaxCacheDurationInHours)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	case oci_apigateway.StaticPublicKeySet:
		result["type"] = "STATIC_KEYS"

		keys := []interface{}{}
		for _, item := range v.Keys {
			keys = append(keys, StaticPublicKeyToMap(item))
		}
		result["keys"] = keys
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToQueryParameterTransformationPolicy(fieldKeyFormat string) (oci_apigateway.QueryParameterTransformationPolicy, error) {
	result := oci_apigateway.QueryParameterTransformationPolicy{}

	if filterQueryParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filter_query_parameters")); ok {
		if tmpList := filterQueryParameters.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "filter_query_parameters"), 0)
			tmp, err := s.mapToFilterQueryParameterPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert filter_query_parameters, encountered error: %v", err)
			}
			result.FilterQueryParameters = &tmp
		}
	}

	if renameQueryParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rename_query_parameters")); ok {
		if tmpList := renameQueryParameters.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rename_query_parameters"), 0)
			tmp, err := s.mapToRenameQueryParameterPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert rename_query_parameters, encountered error: %v", err)
			}
			result.RenameQueryParameters = &tmp
		}
	}

	if setQueryParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "set_query_parameters")); ok {
		if tmpList := setQueryParameters.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "set_query_parameters"), 0)
			tmp, err := s.mapToSetQueryParameterPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert set_query_parameters, encountered error: %v", err)
			}
			result.SetQueryParameters = &tmp
		}
	}

	return result, nil
}

func QueryParameterTransformationPolicyToMap(obj *oci_apigateway.QueryParameterTransformationPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FilterQueryParameters != nil {
		result["filter_query_parameters"] = []interface{}{FilterQueryParameterPolicyToMap(obj.FilterQueryParameters)}
	}

	if obj.RenameQueryParameters != nil {
		result["rename_query_parameters"] = []interface{}{RenameQueryParameterPolicyToMap(obj.RenameQueryParameters)}
	}

	if obj.SetQueryParameters != nil {
		result["set_query_parameters"] = []interface{}{SetQueryParameterPolicyToMap(obj.SetQueryParameters)}
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToQueryParameterValidationItem(fieldKeyFormat string) (oci_apigateway.QueryParameterValidationItem, error) {
	result := oci_apigateway.QueryParameterValidationItem{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if required, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "required")); ok {
		tmp := required.(bool)
		result.Required = &tmp
	}

	return result, nil
}

func QueryParameterValidationItemToMap(obj oci_apigateway.QueryParameterValidationItem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Required != nil {
		result["required"] = bool(*obj.Required)
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToQueryParameterValidationRequestPolicy(fieldKeyFormat string) (oci_apigateway.QueryParameterValidationRequestPolicy, error) {
	result := oci_apigateway.QueryParameterValidationRequestPolicy{}

	if parameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parameters")); ok {
		interfaces := parameters.([]interface{})
		tmp := make([]oci_apigateway.QueryParameterValidationItem, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parameters"), stateDataIndex)
			converted, err := s.mapToQueryParameterValidationItem(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "parameters")) {
			result.Parameters = tmp
		}
	}

	if validationMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validation_mode")); ok {
		result.ValidationMode = oci_apigateway.QueryParameterValidationRequestPolicyValidationModeEnum(validationMode.(string))
	}

	return result, nil
}

func QueryParameterValidationRequestPolicyToMap(obj *oci_apigateway.QueryParameterValidationRequestPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	parameters := []interface{}{}
	for _, item := range obj.Parameters {
		parameters = append(parameters, QueryParameterValidationItemToMap(item))
	}
	result["parameters"] = parameters

	result["validation_mode"] = string(obj.ValidationMode)

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

func (s *ApigatewayDeploymentResourceCrud) mapToRenameHeaderPolicy(fieldKeyFormat string) (oci_apigateway.RenameHeaderPolicy, error) {
	result := oci_apigateway.RenameHeaderPolicy{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_apigateway.RenameHeaderPolicyItem, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToRenameHeaderPolicyItem(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func RenameHeaderPolicyToMap(obj *oci_apigateway.RenameHeaderPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, RenameHeaderPolicyItemToMap(item))
	}
	result["items"] = items

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToRenameHeaderPolicyItem(fieldKeyFormat string) (oci_apigateway.RenameHeaderPolicyItem, error) {
	result := oci_apigateway.RenameHeaderPolicyItem{}

	if from, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "from")); ok {
		tmp := from.(string)
		result.From = &tmp
	}

	if to, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "to")); ok {
		tmp := to.(string)
		result.To = &tmp
	}

	return result, nil
}

func RenameHeaderPolicyItemToMap(obj oci_apigateway.RenameHeaderPolicyItem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.From != nil {
		result["from"] = string(*obj.From)
	}

	if obj.To != nil {
		result["to"] = string(*obj.To)
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToRenameQueryParameterPolicy(fieldKeyFormat string) (oci_apigateway.RenameQueryParameterPolicy, error) {
	result := oci_apigateway.RenameQueryParameterPolicy{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_apigateway.RenameQueryParameterPolicyItem, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToRenameQueryParameterPolicyItem(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func RenameQueryParameterPolicyToMap(obj *oci_apigateway.RenameQueryParameterPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, RenameQueryParameterPolicyItemToMap(item))
	}
	result["items"] = items

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToRenameQueryParameterPolicyItem(fieldKeyFormat string) (oci_apigateway.RenameQueryParameterPolicyItem, error) {
	result := oci_apigateway.RenameQueryParameterPolicyItem{}

	if from, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "from")); ok {
		tmp := from.(string)
		result.From = &tmp
	}

	if to, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "to")); ok {
		tmp := to.(string)
		result.To = &tmp
	}

	return result, nil
}

func RenameQueryParameterPolicyItemToMap(obj oci_apigateway.RenameQueryParameterPolicyItem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.From != nil {
		result["from"] = string(*obj.From)
	}

	if obj.To != nil {
		result["to"] = string(*obj.To)
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToResponseCacheLookupPolicy(fieldKeyFormat string) (oci_apigateway.ResponseCacheLookupPolicy, error) {
	var baseObject oci_apigateway.ResponseCacheLookupPolicy
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("SIMPLE_LOOKUP_POLICY"):
		details := oci_apigateway.SimpleLookupPolicy{}
		if cacheKeyAdditions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cache_key_additions")); ok {
			interfaces := cacheKeyAdditions.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cache_key_additions")) {
				details.CacheKeyAdditions = tmp
			}
		}
		if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
			tmp := isEnabled.(bool)
			details.IsEnabled = &tmp
		}
		if isPrivateCachingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_private_caching_enabled")); ok {
			tmp := isPrivateCachingEnabled.(bool)
			details.IsPrivateCachingEnabled = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ResponseCacheLookupPolicyToMap(obj *oci_apigateway.ResponseCacheLookupPolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_apigateway.SimpleLookupPolicy:
		result["type"] = "SIMPLE_LOOKUP_POLICY"

		result["cache_key_additions"] = v.CacheKeyAdditions

		if v.IsEnabled != nil {
			result["is_enabled"] = bool(*v.IsEnabled)
		}

		if v.IsPrivateCachingEnabled != nil {
			result["is_private_caching_enabled"] = bool(*v.IsPrivateCachingEnabled)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToResponseCacheStorePolicy(fieldKeyFormat string) (oci_apigateway.ResponseCacheStorePolicy, error) {
	var baseObject oci_apigateway.ResponseCacheStorePolicy
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("FIXED_TTL_STORE_POLICY"):
		details := oci_apigateway.FixedTtlResponseCacheStorePolicy{}
		if timeToLiveInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_to_live_in_seconds")); ok {
			tmp := timeToLiveInSeconds.(int)
			details.TimeToLiveInSeconds = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ResponseCacheStorePolicyToMap(obj *oci_apigateway.ResponseCacheStorePolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_apigateway.FixedTtlResponseCacheStorePolicy:
		result["type"] = "FIXED_TTL_STORE_POLICY"

		if v.TimeToLiveInSeconds != nil {
			result["time_to_live_in_seconds"] = int(*v.TimeToLiveInSeconds)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

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

func (s *ApigatewayDeploymentResourceCrud) mapToSetHeaderPolicy(fieldKeyFormat string) (oci_apigateway.SetHeaderPolicy, error) {
	result := oci_apigateway.SetHeaderPolicy{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_apigateway.SetHeaderPolicyItem, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToSetHeaderPolicyItem(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func SetHeaderPolicyToMap(obj *oci_apigateway.SetHeaderPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, SetHeaderPolicyItemToMap(item))
	}
	result["items"] = items

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToSetHeaderPolicyItem(fieldKeyFormat string) (oci_apigateway.SetHeaderPolicyItem, error) {
	result := oci_apigateway.SetHeaderPolicyItem{}

	if ifExists, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "if_exists")); ok {
		result.IfExists = oci_apigateway.SetHeaderPolicyItemIfExistsEnum(ifExists.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
		interfaces := values.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
			result.Values = tmp
		}
	}

	return result, nil
}

func SetHeaderPolicyItemToMap(obj oci_apigateway.SetHeaderPolicyItem) map[string]interface{} {
	result := map[string]interface{}{}

	result["if_exists"] = string(obj.IfExists)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["values"] = obj.Values

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToSetQueryParameterPolicy(fieldKeyFormat string) (oci_apigateway.SetQueryParameterPolicy, error) {
	result := oci_apigateway.SetQueryParameterPolicy{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_apigateway.SetQueryParameterPolicyItem, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToSetQueryParameterPolicyItem(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func SetQueryParameterPolicyToMap(obj *oci_apigateway.SetQueryParameterPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, SetQueryParameterPolicyItemToMap(item))
	}
	result["items"] = items

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToSetQueryParameterPolicyItem(fieldKeyFormat string) (oci_apigateway.SetQueryParameterPolicyItem, error) {
	result := oci_apigateway.SetQueryParameterPolicyItem{}

	if ifExists, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "if_exists")); ok {
		result.IfExists = oci_apigateway.SetQueryParameterPolicyItemIfExistsEnum(ifExists.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
		interfaces := values.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
			result.Values = tmp
		}
	}

	return result, nil
}

func SetQueryParameterPolicyItemToMap(obj oci_apigateway.SetQueryParameterPolicyItem) map[string]interface{} {
	result := map[string]interface{}{}

	result["if_exists"] = string(obj.IfExists)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["values"] = obj.Values

	return result
}

func (s *ApigatewayDeploymentResourceCrud) mapToStaticPublicKey(fieldKeyFormat string) (oci_apigateway.StaticPublicKey, error) {
	var baseObject oci_apigateway.StaticPublicKey
	//discriminator
	formatRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format"))
	var format string
	if ok {
		format = formatRaw.(string)
	} else {
		format = "" // default value
	}
	switch strings.ToLower(format) {
	case strings.ToLower("JSON_WEB_KEY"):
		details := oci_apigateway.JsonWebKey{}
		if alg, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "alg")); ok {
			tmp := alg.(string)
			details.Alg = &tmp
		}
		if e, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "e")); ok {
			tmp := e.(string)
			details.E = &tmp
		}
		if keyOps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_ops")); ok {
			interfaces := keyOps.([]interface{})
			tmp := make([]oci_apigateway.JsonWebKeyKeyOpsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_apigateway.JsonWebKeyKeyOpsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "key_ops")) {
				details.KeyOps = tmp
			}
		}
		if kty, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kty")); ok {
			details.Kty = oci_apigateway.JsonWebKeyKtyEnum(kty.(string))
		}
		if n, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "n")); ok {
			tmp := n.(string)
			details.N = &tmp
		}
		if use, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "use")); ok {
			details.Use = oci_apigateway.JsonWebKeyUseEnum(use.(string))
		}
		if kid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kid")); ok {
			tmp := kid.(string)
			details.Kid = &tmp
		}
		baseObject = details
	case strings.ToLower("PEM"):
		details := oci_apigateway.PemEncodedPublicKey{}
		if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if kid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kid")); ok {
			tmp := kid.(string)
			details.Kid = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown format '%v' was specified", format)
	}
	return baseObject, nil
}

func StaticPublicKeyToMap(obj oci_apigateway.StaticPublicKey) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_apigateway.JsonWebKey:
		result["format"] = "JSON_WEB_KEY"

		if v.Alg != nil {
			result["alg"] = string(*v.Alg)
		}

		if v.E != nil {
			result["e"] = string(*v.E)
		}

		result["key_ops"] = v.KeyOps

		result["kty"] = string(v.Kty)

		if v.N != nil {
			result["n"] = string(*v.N)
		}

		if v.Kid != nil {
			result["kid"] = string(*v.Kid)
		}
		result["use"] = string(v.Use)
	case oci_apigateway.PemEncodedPublicKey:
		result["format"] = "PEM"

		if v.Key != nil {
			result["key"] = string(*v.Key)
		}
		if v.Kid != nil {
			result["kid"] = string(*v.Kid)
		}
	default:
		log.Printf("[WARN] Received 'format' of unknown type %v", obj)
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

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.ChangeDeploymentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeploymentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func mediaTypeHashCodeForBodyValidationContentSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if media_type, ok := m["media_type"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", media_type))
	}
	if validation_type, ok := m["validation_type"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", validation_type))
	}
	return hashcode.String(buf.String())
}
