// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/oracle/oci-go-sdk/v58/apigateway"
	oci_apigateway "github.com/oracle/oci-go-sdk/v58/apigateway"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ApigatewayApiDeploymentSpecificationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularApigatewayApiDeploymentSpecification,
		Schema: map[string]*schema.Schema{
			"api_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"logging_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"access_log": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"execution_log": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"log_level": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"request_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"authentication": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"audiences": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"function_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_anonymous_access_allowed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"issuers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"max_clock_skew_in_seconds": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"public_keys": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"is_ssl_verify_disabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"keys": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"alg": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"e": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"format": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"key": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"key_ops": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"kid": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"kty": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"n": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"use": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"max_cache_duration_in_hours": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"uri": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"token_auth_scheme": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"token_header": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"token_query_param": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"verify_claims": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"is_required": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"key": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"cors": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"allowed_headers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"allowed_methods": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"allowed_origins": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"exposed_headers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_allow_credentials_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"max_age_in_seconds": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"mutual_tls": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"allowed_sans": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_verified_certificate_required": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"rate_limiting": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"rate_in_requests_per_second": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"rate_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"routes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"backend": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"body": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connect_timeout_in_seconds": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"function_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"headers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"is_ssl_verify_disabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"read_timeout_in_seconds": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"send_timeout_in_seconds": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"logging_policies": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"access_log": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"is_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"execution_log": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"is_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"log_level": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"methods": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"request_policies": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"authorization": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_scope": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"body_validation": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"content": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"media_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"validation_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"required": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"validation_mode": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"cors": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"allowed_headers": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_methods": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_origins": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"exposed_headers": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"is_allow_credentials_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"max_age_in_seconds": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"header_transformations": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"filter_headers": {
													Type:     schema.TypeList,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"items": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"type": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"rename_headers": {
													Type:     schema.TypeList,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"items": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"from": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"to": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"set_headers": {
													Type:     schema.TypeList,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"items": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"if_exists": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"values": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"header_validations": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"headers": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"required": {
																Type:     schema.TypeBool,
																Computed: true,
															},
														},
													},
												},
												"validation_mode": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"query_parameter_transformations": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"filter_query_parameters": {
													Type:     schema.TypeList,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"items": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"type": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"rename_query_parameters": {
													Type:     schema.TypeList,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"items": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"from": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"to": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"set_query_parameters": {
													Type:     schema.TypeList,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"items": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"if_exists": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"values": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"query_parameter_validations": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"parameters": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"required": {
																Type:     schema.TypeBool,
																Computed: true,
															},
														},
													},
												},
												"validation_mode": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"response_cache_lookup": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"cache_key_additions": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"is_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"is_private_caching_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"response_policies": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"header_transformations": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"filter_headers": {
													Type:     schema.TypeList,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"items": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"type": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"rename_headers": {
													Type:     schema.TypeList,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"items": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"from": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"to": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"set_headers": {
													Type:     schema.TypeList,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"items": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"if_exists": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"values": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"response_cache_store": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"time_to_live_in_seconds": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readSingularApigatewayApiDeploymentSpecification(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayApiDeploymentSpecificationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiGatewayClient()

	return tfresource.ReadResource(sync)
}

type ApigatewayApiDeploymentSpecificationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.ApiGatewayClient
	Res    *oci_apigateway.GetApiDeploymentSpecificationResponse
}

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) Get() error {
	request := oci_apigateway.GetApiDeploymentSpecificationRequest{}

	if apiId, ok := s.D.GetOkExists("api_id"); ok {
		tmp := apiId.(string)
		request.ApiId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apigateway")

	response, err := s.Client.GetApiDeploymentSpecification(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApigatewayApiDeploymentSpecificationDataSource-", ApigatewayApiDeploymentSpecificationDataSource(), s.D))

	if s.Res.LoggingPolicies != nil {
		s.D.Set("logging_policies", []interface{}{ApiSpecificationLoggingPoliciesToMap(s.Res.LoggingPolicies)})
	} else {
		s.D.Set("logging_policies", nil)
	}

	if s.Res.RequestPolicies != nil {
		s.D.Set("request_policies", []interface{}{ApiSpecificationRequestPoliciesToMap(s.Res.RequestPolicies)})
	} else {
		s.D.Set("request_policies", nil)
	}

	routes := []interface{}{}
	for _, item := range s.Res.Routes {
		routes = append(routes, ApiSpecificationRouteToMap(item, true))
	}
	s.D.Set("routes", routes)

	return nil
}

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToHeaderFieldSpecification(fieldKeyFormat string) (oci_apigateway.HeaderFieldSpecification, error) {
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToJsonWebTokenClaim(fieldKeyFormat string) (oci_apigateway.JsonWebTokenClaim, error) {
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToPublicKeySet(fieldKeyFormat string) (oci_apigateway.PublicKeySet, error) {
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToStaticPublicKey(fieldKeyFormat string) (oci_apigateway.StaticPublicKey, error) {
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
			tmp := make([]apigateway.JsonWebKeyKeyOpsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(apigateway.JsonWebKeyKeyOpsEnum)
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
