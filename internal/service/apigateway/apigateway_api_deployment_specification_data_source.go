// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/apigateway"
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
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
									"cache_key": {
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
									"parameters": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"public_keys": {
										Type:     schema.TypeList,
										Computed: true,
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
									"validation_failure_policy": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"client_details": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"client_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"client_secret_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"client_secret_version_number": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"type": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"fallback_redirect_path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"logout_path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"max_expiry_duration_in_hours": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"response_code": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"response_header_transformations": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"filter_headers": {
																Type:     schema.TypeList,
																Computed: true,
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
												"response_message": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"response_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"scopes": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_uri_details": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
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
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"use_cookies_for_intermediate_steps": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"use_cookies_for_session": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"use_pkce": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"validation_policy": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"additional_validation_policy": {
													Type:     schema.TypeList,
													Computed: true,
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
															"issuers": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
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
												"client_details": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"client_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"client_secret_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"client_secret_version_number": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"type": {
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
												"source_uri_details": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
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
						"dynamic_authentication": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"authentication_servers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"authentication_server_detail": {
													Type:     schema.TypeList,
													Computed: true,
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
															"validation_failure_policy": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"client_details": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"client_id": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"client_secret_id": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"client_secret_version_number": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"type": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"fallback_redirect_path": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"logout_path": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"max_expiry_duration_in_hours": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"response_code": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"response_header_transformations": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"filter_headers": {
																						Type:     schema.TypeList,
																						Computed: true,
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
																		"response_message": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"response_type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"scopes": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"source_uri_details": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
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
																		"type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"use_cookies_for_intermediate_steps": {
																			Type:     schema.TypeBool,
																			Computed: true,
																		},
																		"use_cookies_for_session": {
																			Type:     schema.TypeBool,
																			Computed: true,
																		},
																		"use_pkce": {
																			Type:     schema.TypeBool,
																			Computed: true,
																		},
																	},
																},
															},
															"validation_policy": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"additional_validation_policy": {
																			Type:     schema.TypeList,
																			Computed: true,
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
																					"issuers": {
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
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
																		"client_details": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"client_id": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"client_secret_id": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"client_secret_version_number": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"type": {
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
																		"source_uri_details": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
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
															"cache_key": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"parameters": {
																Type:     schema.TypeMap,
																Computed: true,
																Elem:     schema.TypeString,
															},
														},
													},
												},
												"key": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"expression": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"is_default": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"type": {
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
									"selection_source": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"selector": {
													Type:     schema.TypeString,
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
						"mutual_tls": {
							Type:     schema.TypeList,
							Computed: true,
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
						"usage_plans": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"token_locations": {
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"allowed_post_logout_uris": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
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
									"post_logout_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"read_timeout_in_seconds": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"routing_backends": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"expression": {
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
															"is_default": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"backend": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
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
											},
										},
									},
									"selection_source": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"selector": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"access_log": {
										Type:     schema.TypeList,
										Computed: true,
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"authorization": {
										Type:     schema.TypeList,
										Computed: true,
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
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"filter_headers": {
													Type:     schema.TypeList,
													Computed: true,
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
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"filter_query_parameters": {
													Type:     schema.TypeList,
													Computed: true,
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"header_transformations": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"filter_headers": {
													Type:     schema.TypeList,
													Computed: true,
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToAdditionalValidationPolicy(fieldKeyFormat string) (oci_apigateway.AdditionalValidationPolicy, error) {
	result := oci_apigateway.AdditionalValidationPolicy{}

	if audiences, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "audiences")); ok {
		interfaces := audiences.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "audiences")) {
			result.Audiences = tmp
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
			result.Issuers = tmp
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
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "verify_claims")) {
			result.VerifyClaims = tmp
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToSelectionSourcePolicy(fieldKeyFormat string) (oci_apigateway.SelectionSourcePolicy, error) {
	var baseObject oci_apigateway.SelectionSourcePolicy
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "SINGLE" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("SINGLE"):
		details := oci_apigateway.SingleSelectionSourcePolicy{}
		if selector, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selector")); ok {
			tmp := selector.(string)
			details.Selector = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}

	return baseObject, nil
}

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToClientAppDetails(fieldKeyFormat string) (oci_apigateway.ClientAppDetails, error) {
	var baseObject oci_apigateway.ClientAppDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("CUSTOM"):
		details := oci_apigateway.CustomClientAppDetails{}
		if clientId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_id")); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if clientSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_secret_id")); ok {
			tmp := clientSecretId.(string)
			details.ClientSecretId = &tmp
		}
		if clientSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_secret_version_number")); ok {
			tmp := clientSecretVersionNumber.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert clientSecretVersionNumber string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.ClientSecretVersionNumber = &tmpInt64
		}
		baseObject = details
	case strings.ToLower("VALIDATION_BLOCK"):
		details := oci_apigateway.ValidationBlockClientAppDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToSourceUriDetails(fieldKeyFormat string) (oci_apigateway.SourceUriDetails, error) {
	var baseObject oci_apigateway.SourceUriDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DISCOVERY_URI"):
		details := oci_apigateway.DiscoveryUriSourceUriDetails{}
		if uri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	case strings.ToLower("VALIDATION_BLOCK"):
		details := oci_apigateway.ValidationBlockSourceUriDetails{}
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToTokenAuthenticationValidationPolicy(fieldKeyFormat string) (oci_apigateway.TokenAuthenticationValidationPolicy, error) {
	var baseObject oci_apigateway.TokenAuthenticationValidationPolicy
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("REMOTE_DISCOVERY"):
		details := oci_apigateway.TokenAuthenticationRemoteDiscoveryValidationPolicy{}
		if clientDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_details")); ok {
			if tmpList := clientDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "client_details"), 0)
				tmp, err := s.mapToClientAppDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert client_details, encountered error: %v", err)
				}
				details.ClientDetails = tmp
			}
		}
		if isSslVerifyDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ssl_verify_disabled")); ok {
			tmp := isSslVerifyDisabled.(bool)
			details.IsSslVerifyDisabled = &tmp
		}
		if maxCacheDurationInHours, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_cache_duration_in_hours")); ok {
			tmp := maxCacheDurationInHours.(int)
			details.MaxCacheDurationInHours = &tmp
		}
		if sourceUriDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_uri_details")); ok {
			if tmpList := sourceUriDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_uri_details"), 0)
				tmp, err := s.mapToSourceUriDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert source_uri_details, encountered error: %v", err)
				}
				details.SourceUriDetails = tmp
			}
		}
		if additionalValidationPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_validation_policy")); ok {
			if tmpList := additionalValidationPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "additional_validation_policy"), 0)
				tmp, err := s.mapToAdditionalValidationPolicy(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert additional_validation_policy, encountered error: %v", err)
				}
				details.AdditionalValidationPolicy = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("REMOTE_JWKS"):
		details := oci_apigateway.TokenAuthenticationRemoteJwksValidationPolicy{}
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
		if additionalValidationPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_validation_policy")); ok {
			if tmpList := additionalValidationPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "additional_validation_policy"), 0)
				tmp, err := s.mapToAdditionalValidationPolicy(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert additional_validation_policy, encountered error: %v", err)
				}
				details.AdditionalValidationPolicy = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("STATIC_KEYS"):
		details := oci_apigateway.TokenAuthenticationStaticKeysValidationPolicy{}
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
		if additionalValidationPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_validation_policy")); ok {
			if tmpList := additionalValidationPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "additional_validation_policy"), 0)
				tmp, err := s.mapToAdditionalValidationPolicy(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert additional_validation_policy, encountered error: %v", err)
				}
				details.AdditionalValidationPolicy = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToValidationFailurePolicy(fieldKeyFormat string) (oci_apigateway.ValidationFailurePolicy, error) {
	var baseObject oci_apigateway.ValidationFailurePolicy
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("MODIFY_RESPONSE"):
		details := oci_apigateway.ModifyResponseValidationFailurePolicy{}
		if responseCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_code")); ok {
			tmp := responseCode.(string)
			details.ResponseCode = &tmp
		}
		if responseHeaderTransformations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_header_transformations")); ok {
			if tmpList := responseHeaderTransformations.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "response_header_transformations"), 0)
				tmp, err := s.mapToHeaderTransformationPolicy(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert response_header_transformations, encountered error: %v", err)
				}
				details.ResponseHeaderTransformations = &tmp
			}
		}
		if responseMessage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_message")); ok {
			tmp := responseMessage.(string)
			details.ResponseMessage = &tmp
		}
		baseObject = details
	case strings.ToLower("OAUTH2"):
		details := oci_apigateway.OAuth2ResponseValidationFailurePolicy{}
		if clientDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_details")); ok {
			if tmpList := clientDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "client_details"), 0)
				tmp, err := s.mapToClientAppDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert client_details, encountered error: %v", err)
				}
				details.ClientDetails = tmp
			}
		}
		if fallbackRedirectPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fallback_redirect_path")); ok {
			tmp := fallbackRedirectPath.(string)
			details.FallbackRedirectPath = &tmp
		}
		if logoutPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logout_path")); ok {
			tmp := logoutPath.(string)
			details.LogoutPath = &tmp
		}
		if maxExpiryDurationInHours, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_expiry_duration_in_hours")); ok {
			tmp := maxExpiryDurationInHours.(int)
			details.MaxExpiryDurationInHours = &tmp
		}
		if responseType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_type")); ok {
			details.ResponseType = oci_apigateway.OAuth2ResponseValidationFailurePolicyResponseTypeEnum(responseType.(string))
		}
		if scopes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scopes")); ok {
			interfaces := scopes.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "scopes")) {
				details.Scopes = tmp
			}
		}
		if sourceUriDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_uri_details")); ok {
			if tmpList := sourceUriDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_uri_details"), 0)
				tmp, err := s.mapToSourceUriDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert source_uri_details, encountered error: %v", err)
				}
				details.SourceUriDetails = tmp
			}
		}
		if useCookiesForIntermediateSteps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "use_cookies_for_intermediate_steps")); ok {
			tmp := useCookiesForIntermediateSteps.(bool)
			details.UseCookiesForIntermediateSteps = &tmp
		}
		if useCookiesForSession, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "use_cookies_for_session")); ok {
			tmp := useCookiesForSession.(bool)
			details.UseCookiesForSession = &tmp
		}
		if usePkce, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "use_pkce")); ok {
			tmp := usePkce.(bool)
			details.UsePkce = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToHeaderTransformationPolicy(fieldKeyFormat string) (oci_apigateway.HeaderTransformationPolicy, error) {
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToFilterHeaderPolicy(fieldKeyFormat string) (oci_apigateway.FilterHeaderPolicy, error) {
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToSetHeaderPolicy(fieldKeyFormat string) (oci_apigateway.SetHeaderPolicy, error) {
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToRenameHeaderPolicy(fieldKeyFormat string) (oci_apigateway.RenameHeaderPolicy, error) {
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToRenameHeaderPolicyItem(fieldKeyFormat string) (oci_apigateway.RenameHeaderPolicyItem, error) {
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

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToFilterHeaderPolicyItem(fieldKeyFormat string) (oci_apigateway.FilterHeaderPolicyItem, error) {
	result := oci_apigateway.FilterHeaderPolicyItem{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func (s *ApigatewayApiDeploymentSpecificationDataSourceCrud) mapToSetHeaderPolicyItem(fieldKeyFormat string) (oci_apigateway.SetHeaderPolicyItem, error) {
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
