// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceTaskResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataintegrationWorkspaceTask,
		Read:     readDataintegrationWorkspaceTask,
		Update:   updateDataintegrationWorkspaceTask,
		Delete:   deleteDataintegrationWorkspaceTask,
		Schema: map[string]*schema.Schema{
			// Required
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"REST_TASK",
				}, true),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"registry_metadata": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"aggregator_key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_favorite": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"registry_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"api_call_mode": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"SYNCHRONOUS",
					"ASYNC_GENERIC",
				}, false),
			},
			"auth_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"model_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"OCI_RESOURCE_AUTH_CONFIG",
								"NO_AUTH_CONFIG",
							}, true),
						},
						"model_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"parent_ref": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"parent": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"root_doc_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"resource_principal_source": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"WORKSPACE",
								"APPLICATION",
							}, false),
						},

						// Computed
					},
				},
			},
			"cancel_rest_call_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						// Optional
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"model_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CANCEL_REST_CALL_CONFIG",
							}, false),
						},
						"config_values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"config_param_values": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Optional
												"request_url": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"string_value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"request_payload": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ref_value": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"model_type": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			Computed:         true,
																			DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																			ValidateFunc: validation.StringInSlice([]string{
																				"JSON_TEXT",
																			}, false),
																		},
																		"config_values": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Computed: true,
																			MaxItems: 1,
																			MinItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"config_param_values": {
																						Type:     schema.TypeList,
																						Optional: true,
																						Computed: true,
																						MaxItems: 1,
																						MinItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								// Required
																								"data_param": {
																									Type:     schema.TypeList,
																									Optional: true,
																									Computed: true,
																									MaxItems: 1,
																									MinItems: 1,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"string_value": {
																												Type:     schema.TypeString,
																												Optional: true,
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
																},
															},
														},
													},
												},
											},
										},
									},
									"parent_ref": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"parent": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"root_doc_id": {
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
						"method_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"POST",
								"PUT",
								"DELETE",
								"PATCH",
							}, false),
						},
						"request_headers": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},

			// This Model is not used as part of Rest Task. whenever we will be adding new task model this can be uncommented

			//"conditional_composite_field_map": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Computed: true,
			//	MaxItems: 1,
			//	MinItems: 1,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			// Required
			//
			//			// Optional
			//			"config_values": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"config_param_values": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"int_value": {
			//										Type:     schema.TypeInt,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"object_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"parameter_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"ref_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"root_object_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"string_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"parent_ref": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"parent": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"root_doc_id": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"description": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"field_map_scope": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//						"model_type": {
			//							Type:             schema.TypeString,
			//							Required:         true,
			//							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			//							ValidateFunc: validation.StringInSlice([]string{
			//								"GROUPED_NAME_PATTERN_RULE",
			//								"NAME_LIST_RULE",
			//								"NAME_PATTERN_RULE",
			//								"RENAME_RULE",
			//								"TYPED_NAME_PATTERN_RULE",
			//								"TYPE_LIST_RULE",
			//							}, true),
			//						},
			//
			//						// Optional
			//						"config_values": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"config_param_values": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"int_value": {
			//													Type:     schema.TypeInt,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"object_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"parameter_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"ref_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"root_object_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"string_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//									"parent_ref": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"parent": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"root_doc_id": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"description": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"from_name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"is_cascade": {
			//							Type:     schema.TypeBool,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"is_case_sensitive": {
			//							Type:     schema.TypeBool,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"is_java_regex_syntax": {
			//							Type:     schema.TypeBool,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"is_skip_remaining_rules_on_match": {
			//							Type:     schema.TypeBool,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"key": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"matching_strategy": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"model_version": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"names": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							Elem: &schema.Schema{
			//								Type: schema.TypeString,
			//							},
			//						},
			//						"object_status": {
			//							Type:     schema.TypeInt,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"parent_ref": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"parent": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"root_doc_id": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"pattern": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"rule_type": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"scope": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//
			//							/*MaxItems:         1,
			//							  MinItems:         1,
			//							  DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//						},
			//						"to_name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"types": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//
			//							Elem: &schema.Schema{
			//								Type: schema.TypeString,
			//							},
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"field_maps": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				Elem: &schema.Schema{
			//					Type: schema.TypeString,
			//				},
			//			},
			//			"key": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"model_type": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"model_version": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"object_status": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"parent_ref": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"parent": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"root_doc_id": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//
			//			// Computed
			//		},
			//	},
			//},
			"config_provider_delegate": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"bindings": {

							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"parameter_values": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"root_object_value": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"model_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"model_version": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"object_status": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"simple_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												// Computed
											},
										},
									},
								},
							},
						},

						// Computed
					},
				},
			},

			// This Model is not used as part of Rest Task. whenever we will be adding new task model this can be uncommented

			//"data_flow": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Computed: true,
			//	MaxItems: 1,
			//	MinItems: 1,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			// Required
			//
			//			// Optional
			//			"description": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"flow_config_values": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"config_param_values": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"int_value": {
			//										Type:     schema.TypeInt,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"object_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"parameter_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"ref_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"root_object_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"string_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"parent_ref": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"parent": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"root_doc_id": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"identifier": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"key": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"key_map": {
			//				Type:     schema.TypeMap,
			//				Optional: true,
			//				Computed: true,
			//				Elem:     schema.TypeString,
			//			},
			//			"metadata": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"aggregator": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"description": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"identifier": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"key": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"name": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"type": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"aggregator_key": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"count_statistics": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"object_type_count_list": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"object_count": {
			//													Type:             schema.TypeString,
			//													Optional:         true,
			//													Computed:         true,
			//													ValidateFunc:     tfresource.ValidateInt64TypeString,
			//													DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			//												},
			//												"object_type": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"created_by": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"created_by_name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"identifier_path": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"info_fields": {
			//							Type:     schema.TypeMap,
			//							Optional: true,
			//							Computed: true,
			//							Elem:     schema.TypeString,
			//						},
			//						"is_favorite": {
			//							Type:     schema.TypeBool,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"labels": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							Elem: &schema.Schema{
			//								Type: schema.TypeString,
			//							},
			//						},
			//						"registry_version": {
			//							Type:     schema.TypeInt,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"time_created": {
			//							Type:             schema.TypeString,
			//							Optional:         true,
			//							Computed:         true,
			//							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			//						},
			//						"time_updated": {
			//							Type:             schema.TypeString,
			//							Optional:         true,
			//							Computed:         true,
			//							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			//						},
			//						"updated_by": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"updated_by_name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"model_type": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"model_version": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"name": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"nodes": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"config_provider_delegate": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"description": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"input_links": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"description": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"field_map": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"from_link": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"key": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"model_type": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"model_version": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"object_status": {
			//										Type:     schema.TypeInt,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"parent_ref": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"parent": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"root_doc_id": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//									"port": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"key": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"model_type": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"model_version": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"object_status": {
			//							Type:     schema.TypeInt,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"operator": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"output_links": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"description": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"key": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"model_type": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"model_version": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"object_status": {
			//										Type:     schema.TypeInt,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"parent_ref": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"parent": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"root_doc_id": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//									"port": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"to_links": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										Elem: &schema.Schema{
			//											Type: schema.TypeString,
			//										},
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"parent_ref": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"parent": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"root_doc_id": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"ui_properties": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"coordinate_x": {
			//										Type:     schema.TypeFloat,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"coordinate_y": {
			//										Type:     schema.TypeFloat,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"object_status": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"object_version": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"parameters": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"config_values": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"config_param_values": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"int_value": {
			//													Type:     schema.TypeInt,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"object_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"parameter_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"ref_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"root_object_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"string_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//									"parent_ref": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"parent": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"root_doc_id": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"default_value": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//
			//							/*MaxItems:         1,
			//							MinItems:         1,
			//							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//						},
			//						"description": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"is_input": {
			//							Type:     schema.TypeBool,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"is_output": {
			//							Type:     schema.TypeBool,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"key": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"model_type": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"model_version": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"object_status": {
			//							Type:     schema.TypeInt,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"output_aggregation_type": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"parent_ref": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"parent": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"root_doc_id": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"root_object_default_value": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//
			//							/*MaxItems:         1,
			//							MinItems:         1,
			//							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//						},
			//						"type": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//
			//							/*MaxItems:         1,
			//							  MinItems:         1,
			//							  DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//						},
			//						"type_name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"used_for": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"parent_ref": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"parent": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"root_doc_id": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"target_field_map_summary": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"field_map": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"typed_object_map": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"typed_object": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//
			//			// Computed
			//		},
			//	},
			//},
			//"dataflow_application": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Computed: true,
			//	MaxItems: 1,
			//	MinItems: 1,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			// Required
			//
			//			// Optional
			//			"application_id": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"compartment_id": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"config_values": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"config_param_values": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"int_value": {
			//										Type:     schema.TypeInt,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"object_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"parameter_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"ref_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"root_object_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"string_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"parent_ref": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"parent": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"root_doc_id": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//
			//			// Computed
			//		},
			//	},
			//},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"execute_rest_call_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						// Optional
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"model_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"REST_CALL_CONFIG",
							}, false),
						},
						"config_values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"config_param_values": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Optional
												"request_url": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"string_value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"request_payload": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ref_value": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"model_type": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			Computed:         true,
																			DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																			ValidateFunc: validation.StringInSlice([]string{
																				"JSON_TEXT",
																			}, false),
																		},
																		"config_values": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Computed: true,
																			MaxItems: 1,
																			MinItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"config_param_values": {
																						Type:     schema.TypeList,
																						Optional: true,
																						Computed: true,
																						MaxItems: 1,
																						MinItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								// Required
																								"data_param": {
																									Type:     schema.TypeList,
																									Optional: true,
																									Computed: true,
																									MaxItems: 1,
																									MinItems: 1,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"string_value": {
																												Type:     schema.TypeString,
																												Optional: true,
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
																},
															},
														},
													},
												},
											},
										},
									},
									"parent_ref": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"parent": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"root_doc_id": {
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
						"method_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
								"PUT",
								"DELETE",
								"PATCH",
							}, false),
						},
						"request_headers": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},

						// Computed
					},
				},
			},
			"input_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"model_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"config_values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"config_param_values": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"int_value": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"object_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"parameter_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ref_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"root_object_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"string_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"parent_ref": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"parent": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"root_doc_id": {
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
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fields": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"model_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"object_status": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"parent_ref": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"parent": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"root_doc_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"port_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"is_single_load": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"model_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"op_config_values": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"config_param_values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"config_param_value": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Optional
												"int_value": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"object_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"parameter_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ref_value": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"model_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"model_version": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"object_status": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"root_object_value": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"model_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"model_version": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"object_status": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"string_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},

									// Computed
								},
							},
						},
						"parent_ref": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"parent": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"root_doc_id": {
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
			"operation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"output_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"model_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"config_values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"config_param_values": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"int_value": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"object_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"parameter_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ref_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"root_object_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"string_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"parent_ref": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"parent": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"root_doc_id": {
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
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fields": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"model_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"object_status": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"parent_ref": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"parent": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"root_doc_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"port_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"parallel_load_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"model_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"PARAMETER",
							}, false),
						},

						// Optional
						"config_values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"config_param_values": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"config_param_value": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Optional
															"int_value": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"object_value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"parameter_value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ref_value": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																ForceNew: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"name": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"model_type": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"model_version": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"object_status": {
																			Type:     schema.TypeInt,
																			Optional: true,
																			Computed: true,
																		},
																	},
																},
															},
															"root_object_value": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																ForceNew: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"model_type": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"model_version": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"object_status": {
																			Type:     schema.TypeInt,
																			Optional: true,
																			Computed: true,
																		},
																	},
																},
															},
															"string_value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												// Computed
											},
										},
									},
									"parent_ref": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"parent": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"root_doc_id": {
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
						"default_value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_input": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_output": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"model_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"object_status": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"output_aggregation_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"parent_ref": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"parent": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"root_doc_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"root_object_default_value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"used_for": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"parent_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"parent": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"root_doc_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// This Model is not used as part of Rest Task. whenever we will be adding new task model this can be uncommented

			//"pipeline": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Computed: true,
			//	MaxItems: 1,
			//	MinItems: 1,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			// Required
			//
			//			// Optional
			//			"description": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"flow_config_values": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"config_param_values": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"int_value": {
			//										Type:     schema.TypeInt,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"object_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"parameter_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"ref_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"root_object_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//
			//										/*MaxItems:         1,
			//										MinItems:         1,
			//										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//									},
			//									"string_value": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"parent_ref": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"parent": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"root_doc_id": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"identifier": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"key": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"metadata": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"aggregator": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"description": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"identifier": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"key": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"name": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"type": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"aggregator_key": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"count_statistics": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"object_type_count_list": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"object_count": {
			//													Type:             schema.TypeString,
			//													Optional:         true,
			//													Computed:         true,
			//													ValidateFunc:     tfresource.ValidateInt64TypeString,
			//													DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			//												},
			//												"object_type": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"created_by": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"created_by_name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"identifier_path": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"info_fields": {
			//							Type:     schema.TypeMap,
			//							Optional: true,
			//							Computed: true,
			//							Elem:     schema.TypeString,
			//						},
			//						"is_favorite": {
			//							Type:     schema.TypeBool,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"labels": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							Elem: &schema.Schema{
			//								Type: schema.TypeString,
			//							},
			//						},
			//						"registry_version": {
			//							Type:     schema.TypeInt,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"time_created": {
			//							Type:             schema.TypeString,
			//							Optional:         true,
			//							Computed:         true,
			//							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			//						},
			//						"time_updated": {
			//							Type:             schema.TypeString,
			//							Optional:         true,
			//							Computed:         true,
			//							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			//						},
			//						"updated_by": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"updated_by_name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"model_type": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"model_version": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"name": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"nodes": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"config_provider_delegate": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"description": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"input_links": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"description": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"field_map": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"from_link": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"key": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"model_type": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"model_version": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"object_status": {
			//										Type:     schema.TypeInt,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"parent_ref": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"parent": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"root_doc_id": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//									"port": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"key": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"model_type": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"model_version": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"object_status": {
			//							Type:     schema.TypeInt,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"operator": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"output_links": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"description": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"key": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"model_type": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"model_version": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"object_status": {
			//										Type:     schema.TypeInt,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"parent_ref": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"parent": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"root_doc_id": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//									"port": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"to_links": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										Elem: &schema.Schema{
			//											Type: schema.TypeString,
			//										},
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"parent_ref": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"parent": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"root_doc_id": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"ui_properties": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"coordinate_x": {
			//										Type:     schema.TypeFloat,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"coordinate_y": {
			//										Type:     schema.TypeFloat,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"object_status": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"object_version": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"parameters": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"config_values": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"config_param_values": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"int_value": {
			//													Type:     schema.TypeInt,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"object_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"parameter_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"ref_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"root_object_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"string_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//									"parent_ref": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"parent": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"root_doc_id": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"default_value": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//
			//							/*MaxItems:         1,
			//							MinItems:         1,
			//							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//						},
			//						"description": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"is_input": {
			//							Type:     schema.TypeBool,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"is_output": {
			//							Type:     schema.TypeBool,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"key": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"model_type": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"model_version": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"object_status": {
			//							Type:     schema.TypeInt,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"output_aggregation_type": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"parent_ref": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"parent": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"root_doc_id": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"root_object_default_value": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//
			//							/*MaxItems:         1,
			//							MinItems:         1,
			//							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//						},
			//						"type": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//
			//							/*MaxItems:         1,
			//							  MinItems:         1,
			//							  DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//						},
			//						"type_name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"used_for": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"parent_ref": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"parent": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"root_doc_id": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//			"variables": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"config_values": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"config_param_values": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"int_value": {
			//													Type:     schema.TypeInt,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"object_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"parameter_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"ref_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"root_object_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//
			//													/*MaxItems:         1,
			//													MinItems:         1,
			//													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//												},
			//												"string_value": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//									"parent_ref": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"parent": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"root_doc_id": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"default_value": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//
			//							/*MaxItems:         1,
			//							MinItems:         1,
			//							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,*/
			//						},
			//						"description": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"identifier": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"key": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"model_type": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"model_version": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"name": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"object_status": {
			//							Type:     schema.TypeInt,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"object_version": {
			//							Type:     schema.TypeInt,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"parent_ref": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"parent": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"root_doc_id": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"root_object_default_value": {
			//							Type:     schema.TypeList,
			//							Optional: true,
			//							Computed: true,
			//							MaxItems: 1,
			//							MinItems: 1,
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									// Required
			//
			//									// Optional
			//									"key": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"model_type": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"model_version": {
			//										Type:     schema.TypeString,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"object_status": {
			//										Type:     schema.TypeInt,
			//										Optional: true,
			//										Computed: true,
			//									},
			//									"parent_ref": {
			//										Type:     schema.TypeList,
			//										Optional: true,
			//										Computed: true,
			//										MaxItems: 1,
			//										MinItems: 1,
			//										Elem: &schema.Resource{
			//											Schema: map[string]*schema.Schema{
			//												// Required
			//
			//												// Optional
			//												"parent": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//												"root_doc_id": {
			//													Type:     schema.TypeString,
			//													Optional: true,
			//													Computed: true,
			//												},
			//
			//												// Computed
			//											},
			//										},
			//									},
			//
			//									// Computed
			//								},
			//							},
			//						},
			//						"type": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//
			//			// Computed
			//		},
			//	},
			//},
			"poll_rest_call_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"model_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"POLL_REST_CALL_CONFIG",
							}, false),
						},
						"config_values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"config_param_values": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"request_url": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"string_value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"request_payload": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ref_value": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"name": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"model_type": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			Computed:         true,
																			DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																			ValidateFunc: validation.StringInSlice([]string{
																				"JSON_TEXT",
																			}, false),
																		},
																		"config_values": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Computed: true,
																			MaxItems: 1,
																			MinItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"config_param_values": {
																						Type:     schema.TypeList,
																						Optional: true,
																						Computed: true,
																						MaxItems: 1,
																						MinItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								// Required
																								"data_param": {
																									Type:     schema.TypeList,
																									Optional: true,
																									Computed: true,
																									MaxItems: 1,
																									MinItems: 1,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"string_value": {
																												Type:     schema.TypeString,
																												Optional: true,
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
																},
															},
														},
													},
												},
												"poll_max_duration": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"object_value": {
																Type:     schema.TypeFloat,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"poll_max_duration_unit": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"string_value": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"SECONDS",
																	"MINUTES",
																	"HOURS",
																	"DAYS",
																}, false),
															},
														},
													},
												},
												"poll_interval": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"object_value": {
																Type:     schema.TypeFloat,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"poll_interval_unit": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"string_value": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"SECONDS",
																	"MINUTES",
																	"HOURS",
																	"DAYS",
																}, false),
															},
														},
													},
												},
												"poll_condition": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ref_value": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"name": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"model_type": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			Computed:         true,
																			DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																			ValidateFunc: validation.StringInSlice([]string{
																				"EXPRESSION",
																			}, false),
																		},
																		"expr_string": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},

												// Computed
											},
										},
									},
									"parent_ref": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"parent": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"root_doc_id": {
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
						"method_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
								"PUT",
							}, false),
						},
						"request_headers": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},

						// Computed
					},
				},
			},
			// This Model is not used as part of Rest Task. whenever we will be adding new task model this can be uncommented

			//"script": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Computed: true,
			//	MaxItems: 1,
			//	MinItems: 1,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			// Required
			//
			//			// Optional
			//			"key": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"model_type": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"model_version": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"object_status": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//			},
			//			"parent_ref": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				Computed: true,
			//				MaxItems: 1,
			//				MinItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						// Required
			//
			//						// Optional
			//						"parent": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//						"root_doc_id": {
			//							Type:     schema.TypeString,
			//							Optional: true,
			//							Computed: true,
			//						},
			//
			//						// Computed
			//					},
			//				},
			//			},
			//
			//			// Computed
			//		},
			//	},
			//},
			//"sql_script_type": {
			//	Type:     schema.TypeString,
			//	Optional: true,
			//	Computed: true,
			//},
			"typed_expressions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"config_values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"config_param_values": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"length": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"int_value": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"scale": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"int_value": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
														},
													},
												},

												// Computed
											},
										},
									},
									"parent_ref": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"parent": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"root_doc_id": {
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
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expression": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"model_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"TYPED_EXPRESSION",
							}, false),
						},
						"model_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"object_status": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"parent_ref": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"parent": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"root_doc_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"key_map": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"aggregator": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
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
								},
							},
						},
						"aggregator_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"count_statistics": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"object_type_count_list": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"object_count": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"object_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"created_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"info_fields": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"is_favorite": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"registry_version": {
							Type:     schema.TypeInt,
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
						"updated_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_by_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"object_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createDataintegrationWorkspaceTask(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceTaskResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()
	err := tfresource.CreateResource(d, sync)

	return err
}

func readDataintegrationWorkspaceTask(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceTaskResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

func updateDataintegrationWorkspaceTask(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceTaskResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataintegrationWorkspaceTask(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceTaskResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataintegrationWorkspaceTaskResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataintegration.DataIntegrationClient
	Res                    *oci_dataintegration.Task
	DisableNotFoundRetries bool
}

func (s *DataintegrationWorkspaceTaskResourceCrud) ID() string {
	workspaceTask := *s.Res
	return GetWorkspaceTaskCompositeId(*(workspaceTask.GetKey()), s.D.Get("workspace_id").(string))
}

func (s *DataintegrationWorkspaceTaskResourceCrud) Create() error {
	request := oci_dataintegration.CreateTaskRequest{}
	err := s.populateTopLevelPolymorphicCreateTaskRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.CreateTask(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Task
	return nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) Get() error {
	request := oci_dataintegration.GetTaskRequest{}

	if expandReferences, ok := s.D.GetOkExists("expand_references"); ok {
		tmp := expandReferences.(string)
		request.ExpandReferences = &tmp
	}

	if taskKey, ok := s.D.GetOkExists("key"); ok {
		tmp := taskKey.(string)
		request.TaskKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	taskKey, workspaceId, err := parseWorkspaceTaskCompositeId(s.D.Id())
	if err == nil {
		request.TaskKey = &taskKey
		request.WorkspaceId = &workspaceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.GetTask(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Task
	log.Printf(" Get() Task Response %v", response.Task)
	return nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) Update() error {
	request := oci_dataintegration.UpdateTaskRequest{}
	err := s.populateTopLevelPolymorphicUpdateTaskRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.UpdateTask(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Task
	return nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) Delete() error {
	request := oci_dataintegration.DeleteTaskRequest{}

	if taskKey, ok := s.D.GetOkExists("key"); ok {
		tmp := taskKey.(string)
		request.TaskKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	_, err := s.Client.DeleteTask(context.Background(), request)
	return err
}

func (s *DataintegrationWorkspaceTaskResourceCrud) SetData() error {

	taskKey, workspaceId, err := parseWorkspaceTaskCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key", &taskKey)
		s.D.Set("workspace_id", &workspaceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_dataintegration.TaskFromDataLoaderTaskDetails:
		s.D.Set("model_type", "DATA_LOADER_TASK")

		if v.ConditionalCompositeFieldMap != nil {
			s.D.Set("conditional_composite_field_map", []interface{}{ConditionalCompositeFieldMapToMap(v.ConditionalCompositeFieldMap)})
		} else {
			s.D.Set("conditional_composite_field_map", nil)
		}

		if v.DataFlow != nil {
			s.D.Set("data_flow", []interface{}{DataFlowToMap(v.DataFlow)})
		} else {
			s.D.Set("data_flow", nil)
		}

		if v.IsSingleLoad != nil {
			s.D.Set("is_single_load", *v.IsSingleLoad)
		}

		if v.ParallelLoadLimit != nil {
			s.D.Set("parallel_load_limit", *v.ParallelLoadLimit)
		}

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		inputPorts := []interface{}{}
		for _, item := range v.InputPorts {
			inputPorts = append(inputPorts, InputPortToMap(item))
		}
		s.D.Set("input_ports", inputPorts)

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{DataIntegration_Task_ObjectMetadataToMap(v.Metadata)})
			s.D.Set("registry_metadata", []interface{}{DataintegrationTaskObjectMetadataToRegistryMetadataMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
			s.D.Set("registry_metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		outputPorts := []interface{}{}
		for _, item := range v.OutputPorts {
			outputPorts = append(outputPorts, OutputPortToMap(item))
		}
		s.D.Set("output_ports", outputPorts)

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{DataIntegration_Task_ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}

	case oci_dataintegration.TaskFromIntegrationTaskDetails:
		s.D.Set("model_type", "INTEGRATION_TASK")

		if v.DataFlow != nil {
			s.D.Set("data_flow", []interface{}{DataFlowToMap(v.DataFlow)})
		} else {
			s.D.Set("data_flow", nil)
		}

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		inputPorts := []interface{}{}
		for _, item := range v.InputPorts {
			inputPorts = append(inputPorts, InputPortToMap(item))
		}
		s.D.Set("input_ports", inputPorts)

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{DataIntegration_Task_ObjectMetadataToMap(v.Metadata)})
			s.D.Set("registry_metadata", []interface{}{DataintegrationTaskObjectMetadataToRegistryMetadataMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
			s.D.Set("registry_metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		outputPorts := []interface{}{}
		for _, item := range v.OutputPorts {
			outputPorts = append(outputPorts, OutputPortToMap(item))
		}
		s.D.Set("output_ports", outputPorts)

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{DataIntegration_Task_ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}

	case oci_dataintegration.TaskFromOciDataflowTaskDetails:
		s.D.Set("model_type", "OCI_DATAFLOW_TASK")

		if v.DataflowApplication != nil {
			s.D.Set("dataflow_application", []interface{}{DataflowApplicationToMap(v.DataflowApplication)})
		} else {
			s.D.Set("dataflow_application", nil)
		}

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		inputPorts := []interface{}{}
		for _, item := range v.InputPorts {
			inputPorts = append(inputPorts, InputPortToMap(item))
		}
		s.D.Set("input_ports", inputPorts)

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{DataIntegration_Task_ObjectMetadataToMap(v.Metadata)})
			s.D.Set("registry_metadata", []interface{}{DataintegrationTaskObjectMetadataToRegistryMetadataMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
			s.D.Set("registry_metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		outputPorts := []interface{}{}
		for _, item := range v.OutputPorts {
			outputPorts = append(outputPorts, OutputPortToMap(item))
		}
		s.D.Set("output_ports", outputPorts)

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{DataIntegration_Task_ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}

	case oci_dataintegration.TaskFromPipelineTaskDetails:
		s.D.Set("model_type", "PIPELINE_TASK")

		if v.Pipeline != nil {
			s.D.Set("pipeline", []interface{}{PipelineToMap(v.Pipeline)})
		} else {
			s.D.Set("pipeline", nil)
		}

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		inputPorts := []interface{}{}
		for _, item := range v.InputPorts {
			inputPorts = append(inputPorts, InputPortToMap(item))
		}
		s.D.Set("input_ports", inputPorts)

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{DataIntegration_Task_ObjectMetadataToMap(v.Metadata)})
			s.D.Set("registry_metadata", []interface{}{DataintegrationTaskObjectMetadataToRegistryMetadataMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
			s.D.Set("registry_metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		outputPorts := []interface{}{}
		for _, item := range v.OutputPorts {
			outputPorts = append(outputPorts, OutputPortToMap(item))
		}
		s.D.Set("output_ports", outputPorts)

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{DataIntegration_Task_ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}

	case oci_dataintegration.TaskFromRestTaskDetails:
		s.D.Set("model_type", "REST_TASK")

		s.D.Set("api_call_mode", v.ApiCallMode)

		if v.AuthConfig != nil {
			authConfigArray := []interface{}{}
			if authConfigMap := AuthConfigToMap(&v.AuthConfig); authConfigMap != nil {
				authConfigArray = append(authConfigArray, authConfigMap)
			}
			s.D.Set("auth_config", authConfigArray)
		} else {
			s.D.Set("auth_config", nil)
		}
		if v.CancelRestCallConfig != nil {
			s.D.Set("cancel_rest_call_config", []interface{}{CancelRestCallConfigToMap(v.CancelRestCallConfig)})
		} else {
			s.D.Set("cancel_rest_call_config", nil)
		}

		if v.ExecuteRestCallConfig != nil {
			s.D.Set("execute_rest_call_config", []interface{}{ExecuteRestCallConfigToMap(v.ExecuteRestCallConfig)})
		} else {
			s.D.Set("execute_rest_call_config", nil)
		}

		if v.PollRestCallConfig != nil {
			s.D.Set("poll_rest_call_config", []interface{}{PollRestCallConfigToMap(v.PollRestCallConfig)})
		} else {
			s.D.Set("poll_rest_call_config", nil)
		}

		typedExpressions := []interface{}{}
		for _, item := range v.TypedExpressions {
			typedExpressions = append(typedExpressions, TypedExpressionToMap(item))
		}
		s.D.Set("typed_expressions", typedExpressions)

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		inputPorts := []interface{}{}
		for _, item := range v.InputPorts {
			inputPorts = append(inputPorts, InputPortToMap(item))
		}
		s.D.Set("input_ports", inputPorts)

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{DataIntegration_Task_ObjectMetadataToMap(v.Metadata)})
			s.D.Set("registry_metadata", []interface{}{DataintegrationTaskObjectMetadataToRegistryMetadataMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
			s.D.Set("registry_metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		outputPorts := []interface{}{}
		for _, item := range v.OutputPorts {
			outputPorts = append(outputPorts, OutputPortToMap(item))
		}
		s.D.Set("output_ports", outputPorts)

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{DataIntegration_Task_ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}
	case oci_dataintegration.TaskFromSqlTaskDetails:
		s.D.Set("model_type", "SQL_TASK")

		if v.Operation != nil {
			s.D.Set("operation", v.Operation)

		} else {
			s.D.Set("operation", nil)
		}

		if v.Script != nil {
			s.D.Set("script", []interface{}{ScriptToMap(v.Script)})
		} else {
			s.D.Set("script", nil)
		}

		s.D.Set("sql_script_type", v.SqlScriptType)

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		inputPorts := []interface{}{}
		for _, item := range v.InputPorts {
			inputPorts = append(inputPorts, InputPortToMap(item))
		}
		s.D.Set("input_ports", inputPorts)

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{DataIntegration_Task_ObjectMetadataToMap(v.Metadata)})
			s.D.Set("registry_metadata", []interface{}{DataintegrationTaskObjectMetadataToRegistryMetadataMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
			s.D.Set("registry_metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		outputPorts := []interface{}{}
		for _, item := range v.OutputPorts {
			outputPorts = append(outputPorts, OutputPortToMap(item))
		}
		s.D.Set("output_ports", outputPorts)

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{DataIntegration_Task_ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}

	default:
		log.Printf("[WARN] Received 'model_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func GetWorkspaceTaskCompositeId(taskKey string, workspaceId string) string {
	taskKey = url.PathEscape(taskKey)
	workspaceId = url.PathEscape(workspaceId)
	compositeId := "workspaces/" + workspaceId + "/tasks/" + taskKey
	return compositeId
}

func parseWorkspaceTaskCompositeId(compositeId string) (taskKey string, workspaceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("workspaces/.*/tasks/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	workspaceId, _ = url.PathUnescape(parts[1])
	taskKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToAggregatorSummary(fieldKeyFormat string) (oci_dataintegration.AggregatorSummary, error) {
	result := oci_dataintegration.AggregatorSummary{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func DataIntegration_Task_AggregatorSummaryToMap(obj *oci_dataintegration.AggregatorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToAuthConfig(fieldKeyFormat string) (oci_dataintegration.AuthConfig, error) {
	var baseObject oci_dataintegration.AuthConfig
	//discriminator
	modelTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type"))
	var modelType string
	if ok {
		modelType = modelTypeRaw.(string)
	} else {
		modelType = "" // default value
	}
	switch strings.ToLower(modelType) {
	case strings.ToLower("OCI_RESOURCE_AUTH_CONFIG"):
		details := oci_dataintegration.ResourcePrincipalAuthConfig{}
		if resourcePrincipalSource, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_principal_source")); ok {
			details.ResourcePrincipalSource = oci_dataintegration.ResourcePrincipalAuthConfigResourcePrincipalSourceEnum(resourcePrincipalSource.(string))
		}
		if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
				tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
				}
				details.ParentRef = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown model_type '%v' was specified", modelType)
	}
	return baseObject, nil
}

func AuthConfigToMap(obj *oci_dataintegration.AuthConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_dataintegration.ResourcePrincipalAuthConfig:
		result["model_type"] = "OCI_RESOURCE_AUTH_CONFIG"
		result["model_version"] = *v.ModelVersion

		result["resource_principal_source"] = string(v.ResourcePrincipalSource)
	default:
		log.Printf("[WARN] Received 'model_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToAuthDetails(fieldKeyFormat string) (oci_dataintegration.AuthDetails, error) {
	result := oci_dataintegration.AuthDetails{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		result.ModelType = oci_dataintegration.AuthDetailsModelTypeEnum(modelType.(string))
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	return result, nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToCancelRestCallConfig(fieldKeyFormat string) (oci_dataintegration.CancelRestCallConfig, error) {
	result := oci_dataintegration.CancelRestCallConfig{}

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp, err := s.mapToRestCallConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
			}
			result.ConfigValues = &tmp
		}
	}
	if methodType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "method_type")); ok {
		result.MethodType = oci_dataintegration.CancelRestCallConfigMethodTypeEnum(methodType.(string))
	}

	if requestHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_headers")); ok {
		result.RequestHeaders = tfresource.ObjectMapToStringMap(requestHeaders.(map[string]interface{}))
	}

	return result, nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToRestCallConfigValues(fieldKeyFormat string) (oci_dataintegration.ConfigValues, error) {
	result := oci_dataintegration.ConfigValues{}

	if configParamValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_param_values")); ok {
		if tmpList := configParamValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_param_values"), 0)
			tmp := map[string]oci_dataintegration.ConfigParameterValue{}
			var err error
			if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "request_url")); ok {
				if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
					result := oci_dataintegration.ConfigParameterValue{}
					result.StringValue = ObjectMapToStringValue(tmpList[0].(map[string]interface{}))
					tmp["requestURL"] = result
				}
			}
			if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "request_payload")); ok {
				if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
					tmp["requestPayload"], err = s.mapToRequestPayload(fieldKeyFormatNextLevel, tmpList)
				}
			}
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ConfigParamValues = tmp
		}
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	return result, nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToRequestPayload(fieldKeyFormat string, tmpList []interface{}) (oci_dataintegration.ConfigParameterValue, error) {
	result := oci_dataintegration.ConfigParameterValue{}

	configParamValue := tmpList[0].(map[string]interface{})
	log.Printf("configParamValue Map for request_payload %v", configParamValue)
	val, ok := configParamValue["parameter_value"]
	if ok {
		temp := val.(string)
		log.Printf("parameter_value from test %v", temp)
		if len(temp) > 0 {
			result.ParameterValue = &temp
			log.Printf("Parameter_value for  request_payalod is %v", *result.ParameterValue)
		}
	}
	fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "request_payload"), 0)
	if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "ref_value")); ok {
		if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextToNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormatNextLevel, "ref_value"), 0)

			tmp := s.ObjectMapToPayloadRefValue(fieldKeyFormatNextToNextLevel)
			log.Printf("Ref_value going in the request %v", tmp)
			result.RefValue = &tmp
		}
	}
	return result, nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) ObjectMapToPayloadRefValue(fieldKeyFormat string) interface{} {
	result := make(map[string]interface{})

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp := s.ObjectMapToPayloadRefValueConfigValues(fieldKeyFormatNextLevel)
			log.Printf("Ref_value config_values going in the request %v", tmp)
			result["configValues"] = &tmp
		}
	}
	if model_type, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		result["modelType"] = model_type.(string)
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) ObjectMapToPayloadRefValueConfigValues(fieldKeyFormat string) oci_dataintegration.ConfigValues {
	result := oci_dataintegration.ConfigValues{}

	if configParamValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_param_values")); ok {
		if tmpList := configParamValues.([]interface{}); len(tmpList) > 0 {
			tmp := map[string]oci_dataintegration.ConfigParameterValue{}
			configParameterValue := oci_dataintegration.ConfigParameterValue{}
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_param_values"), 0)
			if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "data_param")); ok {
				if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
					log.Printf("Ref_value config_values before dataParam %v", tmp)
					configParameterValue.StringValue = ObjectMapToStringValue(tmpList[0].(map[string]interface{}))
					tmp["dataParam"] = configParameterValue
				}
			}
			log.Printf("Ref_value config_values going dataParam in the request %v", tmp)
			result.ConfigParamValues = tmp
		}
	}
	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToPayloadRefDataParam(fieldKeyFormat string) oci_dataintegration.ConfigParameterValue {
	result := oci_dataintegration.ConfigParameterValue{}

	if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_param")); ok {
		if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
			result.StringValue = ObjectMapToStringValue(tmpList[0].(map[string]interface{}))
		}
	}
	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToPollCondition(tmpList []interface{}) (oci_dataintegration.ConfigParameterValue, error) {
	result := oci_dataintegration.ConfigParameterValue{}
	temp := tmpList[0].(map[string]interface{})
	parameterValue, ok := temp["parameter_value"]
	if ok {
		pnt := parameterValue.(string)
		if len(pnt) > 0 && pnt != "" {
			log.Printf("Should never come here mapToPollCondition ParameterValue %v", pnt)
			result.ParameterValue = &pnt
		}
	}
	val, ok := temp["ref_value"]
	log.Printf("Ref_value from mapToPollCondition %v", val)
	temp1 := val.([]interface{})
	if len(temp1) > 0 {
		tmp := ObjectMapToRefValue(temp1[0].(map[string]interface{}))
		log.Printf("Ref_value After conversion %v", tmp)
		result.RefValue = &tmp
	}
	return result, nil
}

func ObjectMapToStringValue(configParamValue map[string]interface{}) *string {
	val, ok := configParamValue["string_value"]
	if ok {
		tmp := val.(string)
		log.Printf("string_value %v", tmp)
		return &tmp
	}
	return nil
}

func ObjectMapToObjectvalue(configParamValue map[string]interface{}) *interface{} {
	val, ok := configParamValue["object_value"]
	if ok {
		return &val
	}
	return nil
}

func ObjectMapToRefValue(configParamValue map[string]interface{}) interface{} {
	result := make(map[string]interface{})
	modelType, ok := configParamValue["model_type"]
	if ok {
		log.Printf("Ref_value modelType in ObjectMapToRefValue %v", modelType)
		result["modelType"] = modelType.(string)
	}
	name, ok := configParamValue["name"]
	if ok {
		log.Printf("Ref_value name in ObjectMapToRefValue %v", name)
		if len(name.(string)) > 0 {
			result["name"] = name.(string)
		}
	}
	exprString, ok := configParamValue["expr_string"]
	if ok {
		log.Printf("Ref_value exprString in ObjectMapToRefValue %v", exprString)
		result["exprString"] = exprString.(string)
	}
	return result
}

func CancelRestCallConfigToMap(obj *oci_dataintegration.CancelRestCallConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigValues != nil {
		result["config_values"] = []interface{}{ConfigValuesToMapForRest(obj.ConfigValues)}
	}

	result["method_type"] = string(obj.MethodType)
	result["model_type"] = "CANCEL_REST_CALL_CONFIG"
	result["request_headers"] = obj.RequestHeaders
	result["request_headers"] = obj.RequestHeaders

	return result
}

func ConfigValuesToMapForRest(obj *oci_dataintegration.ConfigValues) map[string]interface{} {
	result := map[string]interface{}{}

	result["config_param_values"] = ConfigValuesMapToObjectArrayRest(obj.ConfigParamValues)

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	return result
}

func ConfigValuesMapToObjectPayloadRef(configParamValues map[string]interface{}) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	dataParam, ok := configParamValues["dataParam"]
	if ok {
		result[0]["data_param"] = ConfigValueMapToRequestPayloadString(dataParam.(map[string]interface{}))
	}
	return result
}

func ConfigValueMapToRequestPayloadString(configParameterValue map[string]interface{}) []map[string]string {
	var result = make([]map[string]string, 1)
	result[0] = make(map[string]string)
	configParamValues, ok := configParameterValue["stringValue"]
	if ok {
		result[0]["string_value"] = configParamValues.(string)
	}
	return result
}

func ConfigValuesMapToObjectArrayRest(configParamValues map[string]oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	//i := 0
	result[0] = make(map[string]interface{})
	requestUrl, ok := configParamValues["requestURL"]
	if ok {
		result[0]["request_url"] = ConfigValueMapToRequestUrlRest(requestUrl)
	}

	requestPayload, ok := configParamValues["requestPayload"]
	if ok {
		result[0]["request_payload"] = ConfigValueMapToRequestPayloadRest(requestPayload)
	}
	return result
}

func ConfigValueMapToRequestUrlRest(configParameterValue oci_dataintegration.ConfigParameterValue) []map[string]string {
	var result = make([]map[string]string, 1)
	result[0] = make(map[string]string)
	if configParameterValue.StringValue != nil {
		result[0]["string_value"] = *configParameterValue.StringValue
	}
	return result
}

func ConfigValueMapToRequestPayloadRest(configParameterValue oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	if configParameterValue.ParameterValue != nil {
		tmp := *configParameterValue.ParameterValue
		if len(tmp) > 0 && tmp != "" {
			log.Printf("Should nevger be here. ParameterValue :=  %v", tmp)
			result[0]["parameter_value"] = tmp
		}

	}
	if configParameterValue.RefValue != nil {
		result[0]["ref_value"] = ReqPayloadRefToMap(*configParameterValue.RefValue)
	}
	return result
}

func ReqPayloadRefToMap(refValue interface{}) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	refValueObj := refValue.(map[string]interface{})
	configVal, ok := refValueObj["configValues"]
	if ok {
		result[0]["config_values"] = ConfigValuesToMapForReqPayloadRef(configVal.(map[string]interface{}))
	}

	modelType, ok := refValueObj["modelType"]
	if ok {
		result[0]["model_type"] = modelType.(string)
	}

	return result
}

func ConfigValuesToMapForReqPayloadRef(obj map[string]interface{}) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	configParamValues, ok := obj["configParamValues"]
	if ok {
		result[0]["config_param_values"] = ConfigValuesMapToObjectPayloadRef(configParamValues.(map[string]interface{}))
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToConditionalCompositeFieldMap(fieldKeyFormat string) (oci_dataintegration.ConditionalCompositeFieldMap, error) {
	result := oci_dataintegration.ConditionalCompositeFieldMap{}

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
			}
			result.ConfigValues = &tmp
		}
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if fieldMapScope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_map_scope")); ok {
		interfaces := fieldMapScope.([]interface{})
		tmp := make([]oci_dataintegration.ProjectionRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "field_map_scope"), stateDataIndex)
			converted, err := s.mapToProjectionRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "field_map_scope")) {
			result.FieldMapScope = tmp
		}
	}

	if fieldMaps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_maps")); ok {
		tmp := strings.Join(fieldMaps.([]string), " ")
		var fieldMapsObj []oci_dataintegration.FieldMap
		err := json.Unmarshal([]byte(tmp), &fieldMapsObj)
		if err != nil {
			return result, err
		}
		result.FieldMaps = fieldMapsObj
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	return result, nil
}

func ConditionalCompositeFieldMapToMap(obj *oci_dataintegration.ConditionalCompositeFieldMap) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigValues != nil {
		result["config_values"] = []interface{}{ConfigValuesToMap(obj.ConfigValues)}
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	fieldMapScope := []interface{}{}
	for _, item := range obj.FieldMapScope {
		fieldMapScope = append(fieldMapScope, ProjectionRuleToMap(item))
	}
	result["field_map_scope"] = fieldMapScope

	if obj.FieldMaps != nil {
		tmp, _ := json.Marshal(obj.FieldMaps)
		result["field_maps"] = string(tmp)
	}
	result["field_maps"] = obj.FieldMaps

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToConfigValues(fieldKeyFormat string) (oci_dataintegration.ConfigValues, error) {
	result := oci_dataintegration.ConfigValues{}

	if configParamValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_param_values")); ok {
		tmp := configParamValues.([]interface{})
		tempResult := map[string]oci_dataintegration.ConfigParameterValue{}
		for _, k := range tmp {
			tmp1 := k.(map[string]interface{}) //key ParameterValue object
			if tmp1["key"] == nil || tmp1["config_param_value"] == nil {
				return result, nil
			}
			tempResult[tmp1["key"].(string)] = ObjectMapToConfigParamValueMap(tmp1["config_param_value"])
		}
		log.Printf("ConfigValues Before %v", tempResult)
		result.ConfigParamValues = tempResult
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	return result, nil
}

func ConfigValuesToMap(obj *oci_dataintegration.ConfigValues) map[string]interface{} {
	result := map[string]interface{}{}

	result["config_param_values"] = ConfigValuesMapToObjectArray(obj.ConfigParamValues)

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToCountStatistic(fieldKeyFormat string) (oci_dataintegration.CountStatistic, error) {
	result := oci_dataintegration.CountStatistic{}

	if objectTypeCountList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_type_count_list")); ok {
		interfaces := objectTypeCountList.([]interface{})
		tmp := make([]oci_dataintegration.CountStatisticSummary, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_type_count_list"), stateDataIndex)
			converted, err := s.mapToCountStatisticSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "object_type_count_list")) {
			result.ObjectTypeCountList = tmp
		}
	}

	return result, nil
}

func DataIntegration_Task_CountStatisticToMap(obj *oci_dataintegration.CountStatistic) map[string]interface{} {
	result := map[string]interface{}{}

	objectTypeCountList := []interface{}{}
	for _, item := range obj.ObjectTypeCountList {
		objectTypeCountList = append(objectTypeCountList, DataIntegration_Task_CountStatisticSummaryToMap(item))
	}
	result["object_type_count_list"] = objectTypeCountList

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToCountStatisticSummary(fieldKeyFormat string) (oci_dataintegration.CountStatisticSummary, error) {
	result := oci_dataintegration.CountStatisticSummary{}

	if objectCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_count")); ok {
		tmp := objectCount.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert objectCount string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.ObjectCount = &tmpInt64
	}

	if objectType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_type")); ok {
		result.ObjectType = oci_dataintegration.CountStatisticSummaryObjectTypeEnum(objectType.(string))
	}

	return result, nil
}

func DataIntegration_Task_CountStatisticSummaryToMap(obj oci_dataintegration.CountStatisticSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectCount != nil {
		result["object_count"] = strconv.FormatInt(*obj.ObjectCount, 10)
	}

	result["object_type"] = string(obj.ObjectType)

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToCreateConfigProvider(fieldKeyFormat string) (oci_dataintegration.CreateConfigProvider, error) {
	result := oci_dataintegration.CreateConfigProvider{}

	if bindings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bindings")); ok {
		tmp := bindings.([]interface{})
		tempresult := map[string]oci_dataintegration.ParameterValue{}
		for _, k := range tmp {
			tmp1 := k.(map[string]interface{}) //key ParameterValue object
			if tmp1["key"] == nil || tmp1["parameter_values"] == nil {
				return result, nil
			}
			key := tmp1["key"].(string)
			val := ObjectMapToParamValueMap(tmp1["parameter_values"])
			log.Printf("Updating Bindings with key %v and value %v", key, val)
			tempresult[key] = val
		}
		result.Bindings = tempresult
	}

	return result, nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToConfigProvider(fieldKeyFormat string) (oci_dataintegration.ConfigProvider, error) {
	result := oci_dataintegration.ConfigProvider{}

	if bindings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bindings")); ok {
		tmp := bindings.([]interface{})
		tempresult := map[string]oci_dataintegration.ParameterValue{}
		for _, k := range tmp {
			tmp1 := k.(map[string]interface{}) //key ParameterValue object
			if tmp1["key"] == nil || tmp1["parameter_values"] == nil {
				return result, nil
			}
			key := tmp1["key"].(string)
			val := ObjectMapToParamValueMap(tmp1["parameter_values"])
			log.Printf("Updating Bindings with key %v and value %v", key, val)
			tempresult[key] = val
		}
		result.Bindings = tempresult
	}

	return result, nil
}

func ConfigProviderToMap(obj *oci_dataintegration.ConfigProvider) map[string]interface{} {
	result := map[string]interface{}{}
	result["bindings"] = ParameterMapToObjectArray(obj.Bindings)

	return result
}

func ParameterMapToObjectArray(bindings map[string]oci_dataintegration.ParameterValue) []map[string]interface{} {
	var result = make([]map[string]interface{}, len(bindings))
	i := 0
	for k, v := range bindings {
		result[i] = make(map[string]interface{})
		result[i]["key"] = k
		result[i]["parameter_values"] = ParamValueMapToObjectMap(v)
		i++
	}
	return result
}

func ConfigValuesMapToObjectArray(bindings map[string]oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	dupCount := 0
	for k, _ := range bindings {
		if k == "usedInParam" { //"usedInParam" configValue is added by create API for each parameter
			dupCount++
		}
	}

	var result = make([]map[string]interface{}, len(bindings)-dupCount)
	i := 0
	for k, v := range bindings {
		if k == "usedInParam" { //"usedInParam" configValue is added by create API for each parameter
			continue
		}
		result[i] = make(map[string]interface{})
		result[i]["key"] = k
		result[i]["config_param_value"] = ConfigParamValueMapToObjectMap(v)
		i++
	}
	log.Printf("ConfigValues after %v", result)
	return result
}

func ConfigParamValueMapToObjectMap(configParamValue oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	result := make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	if configParamValue.StringValue != nil && len(*configParamValue.StringValue) != 0 {
		result[0]["string_value"] = *configParamValue.StringValue
	}
	if configParamValue.RootObjectValue != nil {
		result[0]["root_object_value"] = *configParamValue.RootObjectValue
	}
	if configParamValue.ParameterValue != nil {
		tmp := *configParamValue.ParameterValue
		if len(tmp) > 0 {
			log.Printf("Should never Enter here ConfigParamValueMapToObjectMap for ParameterValue %v", tmp)
			result[0]["parameter_value"] = tmp
		}
	}
	if configParamValue.ObjectValue != nil && len((*configParamValue.ObjectValue).(string)) != 0 {
		result[0]["object_value"] = (*configParamValue.ObjectValue).(string)
	}
	if configParamValue.IntValue != nil && *configParamValue.IntValue != 0 {
		log.Printf("Entering Wrong if statement with %v", *configParamValue.IntValue)
		result[0]["int_value"] = *configParamValue.IntValue
	}
	if configParamValue.RefValue != nil {
		result[0]["ref_value"] = *configParamValue.RefValue
	}

	return result
}

func ObjectMapToParamValueMap(paramValue interface{}) oci_dataintegration.ParameterValue {
	result := oci_dataintegration.ParameterValue{}
	tmp1 := paramValue.([]interface{})
	tmp := tmp1[0].(map[string]interface{})
	log.Printf("ParamValue Object %v", tmp)
	for k, v := range tmp {
		switch k {
		case "simple_value":
			log.Printf("Updating SimpleValue with %v", v)
			temp := v
			result.SimpleValue = &temp
		}
	}
	return result
}

func ParamValueMapToObjectMap(paramValue oci_dataintegration.ParameterValue) []map[string]interface{} {
	result := make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	if paramValue.SimpleValue != nil {
		result[0]["simple_value"] = (*paramValue.SimpleValue).(string)
	}
	if paramValue.RootObjectValue != nil {
		result[0]["root_object_value"] = *paramValue.RootObjectValue
	}

	return result
}

func ObjectMapToConfigParamValueMap(configParamValue interface{}) oci_dataintegration.ConfigParameterValue {
	result := oci_dataintegration.ConfigParameterValue{}
	tmp1 := configParamValue.([]interface{})
	tmp := tmp1[0].(map[string]interface{})
	log.Printf("configParamValue Object %v", tmp)
	for k, v := range tmp {
		switch k {
		case "int_value":
			temp := v.(int)
			if temp != 0 {
				//emp := v.(int)
				result.IntValue = &temp
			}
		case "string_value":
			temp := v.(string)
			if temp != "" {
				result.StringValue = &temp
			}
		case "parameter_value":
			temp := v.(string)
			if len(temp) > 0 {
				result.ParameterValue = &temp
			}
		case "object_value":
			temp := v.(string)
			if temp != "" {
				result.ObjectValue = &v
			}

		}
	}
	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToDataFlow(fieldKeyFormat string) (oci_dataintegration.DataFlow, error) {
	result := oci_dataintegration.DataFlow{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if flowConfigValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "flow_config_values")); ok {
		if tmpList := flowConfigValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "flow_config_values"), 0)
			tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert flow_config_values, encountered error: %v", err)
			}
			result.FlowConfigValues = &tmp
		}
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if keyMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_map")); ok {
		result.KeyMap = tfresource.ObjectMapToStringMap(keyMap.(map[string]interface{}))
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "metadata"), 0)
			tmp, err := s.mapToObjectMetadata(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert metadata, encountered error: %v", err)
			}
			result.Metadata = &tmp
		}
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if nodes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nodes")); ok {
		interfaces := nodes.([]interface{})
		tmp := make([]oci_dataintegration.FlowNode, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "nodes"), stateDataIndex)
			converted, err := s.mapToFlowNode(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nodes")) {
			result.Nodes = tmp
		}
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_version")); ok {
		tmp := objectVersion.(int)
		result.ObjectVersion = &tmp
	}

	if parameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parameters")); ok {
		interfaces := parameters.([]interface{})
		tmp := make([]oci_dataintegration.Parameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parameters"), stateDataIndex)
			converted, err := s.mapToParameter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "parameters")) {
			result.Parameters = tmp
		}
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if targetFieldMapSummary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_field_map_summary")); ok {
		result.TargetFieldMapSummary = ObjectMapToTargetFieldMapSummary(targetFieldMapSummary.(map[string]interface{}))
	}

	if typedObjectMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "typed_object_map")); ok {
		result.TypedObjectMap = ObjectMapToTypedObjectMap(typedObjectMap.(map[string]interface{}))
	}

	return result, nil
}

func ObjectMapToTypedObjectMap(m map[string]interface{}) map[string]oci_dataintegration.TypedObjectWrapper {
	result := map[string]oci_dataintegration.TypedObjectWrapper{}
	return result
}

func ObjectMapToTargetFieldMapSummary(m map[string]interface{}) map[string]oci_dataintegration.FieldMapWrapper {
	result := map[string]oci_dataintegration.FieldMapWrapper{}
	return result
}

func DataFlowToMap(obj *oci_dataintegration.DataFlow) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.FlowConfigValues != nil {
		result["flow_config_values"] = []interface{}{ConfigValuesToMap(obj.FlowConfigValues)}
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["key_map"] = obj.KeyMap

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{DataIntegration_Task_ObjectMetadataToMap(obj.Metadata)}
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	nodes := []interface{}{}
	for _, item := range obj.Nodes {
		nodes = append(nodes, FlowNodeToMap(item))
	}
	result["nodes"] = nodes

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	parameters := []interface{}{}
	for _, item := range obj.Parameters {
		parameters = append(parameters, ParameterToMap(item))
	}
	result["parameters"] = parameters

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	result["target_field_map_summary"] = obj.TargetFieldMapSummary
	result["target_field_map_summary"] = obj.TargetFieldMapSummary

	result["typed_object_map"] = obj.TypedObjectMap
	result["typed_object_map"] = obj.TypedObjectMap

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToDataflowApplication(fieldKeyFormat string) (oci_dataintegration.DataflowApplication, error) {
	result := oci_dataintegration.DataflowApplication{}

	if applicationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "application_id")); ok {
		tmp := applicationId.(string)
		result.ApplicationId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
			}
			result.ConfigValues = &tmp
		}
	}

	return result, nil
}

func DataflowApplicationToMap(obj *oci_dataintegration.DataflowApplication) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicationId != nil {
		result["application_id"] = string(*obj.ApplicationId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConfigValues != nil {
		result["config_values"] = []interface{}{ConfigValuesToMap(obj.ConfigValues)}
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToExecuteRestCallConfig(fieldKeyFormat string) (oci_dataintegration.ExecuteRestCallConfig, error) {
	result := oci_dataintegration.ExecuteRestCallConfig{}

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp, err := s.mapToRestCallConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
			}
			result.ConfigValues = &tmp
		}
	}

	if methodType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "method_type")); ok {
		result.MethodType = oci_dataintegration.ExecuteRestCallConfigMethodTypeEnum(methodType.(string))
	}

	if requestHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_headers")); ok {
		result.RequestHeaders = tfresource.ObjectMapToStringMap(requestHeaders.(map[string]interface{}))
	}

	return result, nil
}

func ExecuteRestCallConfigToMap(obj *oci_dataintegration.ExecuteRestCallConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigValues != nil {
		result["config_values"] = []interface{}{ConfigValuesToMapForRest(obj.ConfigValues)}
	}

	result["method_type"] = string(obj.MethodType)
	result["model_type"] = "REST_CALL_CONFIG"

	if obj.RequestHeaders != nil {
		result["request_headers"] = obj.RequestHeaders
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToExpression(fieldKeyFormat string) (oci_dataintegration.Expression, error) {
	result := oci_dataintegration.Expression{}

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
			}
			result.ConfigValues = &tmp
		}
	}

	if exprString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "expr_string")); ok {
		tmp := exprString.(string)
		result.ExprString = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	return result, nil
}

func ExpressionToMap(obj *oci_dataintegration.Expression) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigValues != nil {
		result["config_values"] = []interface{}{ConfigValuesToMap(obj.ConfigValues)}
	}

	if obj.ExprString != nil {
		result["expr_string"] = string(*obj.ExprString)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToFlowNode(fieldKeyFormat string) (oci_dataintegration.FlowNode, error) {
	result := oci_dataintegration.FlowNode{}

	if configProviderDelegate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_provider_delegate")); ok {
		tmp := configProviderDelegate.(string)
		var configProviderDelegateObj oci_dataintegration.ConfigProvider
		err := json.Unmarshal([]byte(tmp), &configProviderDelegateObj)
		if err != nil {
			return result, err
		}
		result.ConfigProviderDelegate = &configProviderDelegateObj
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if inputLinks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "input_links")); ok {
		interfaces := inputLinks.([]interface{})
		tmp := make([]oci_dataintegration.InputLink, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "input_links"), stateDataIndex)
			converted, err := s.mapToInputLink(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "input_links")) {
			result.InputLinks = tmp
		}
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
		tmp := operator.(string)
		var operatorObj oci_dataintegration.Operator
		err := json.Unmarshal([]byte(tmp), &operatorObj)
		if err != nil {
			return result, err
		}
		result.Operator = operatorObj
	}

	if outputLinks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_links")); ok {
		interfaces := outputLinks.([]interface{})
		tmp := make([]oci_dataintegration.OutputLink, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "output_links"), stateDataIndex)
			converted, err := s.mapToOutputLink(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "output_links")) {
			result.OutputLinks = tmp
		}
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if uiProperties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ui_properties")); ok {
		if tmpList := uiProperties.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "ui_properties"), 0)
			tmp, err := s.mapToUIProperties(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert ui_properties, encountered error: %v", err)
			}
			result.UiProperties = &tmp
		}
	}

	return result, nil
}

func FlowNodeToMap(obj oci_dataintegration.FlowNode) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigProviderDelegate != nil {
		tmp, _ := json.Marshal(obj.ConfigProviderDelegate)
		result["config_provider_delegate"] = string(tmp)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	inputLinks := []interface{}{}
	for _, item := range obj.InputLinks {
		inputLinks = append(inputLinks, InputLinkToMap(item))
	}
	result["input_links"] = inputLinks

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.Operator != nil {
		tmp, _ := json.Marshal(obj.Operator)
		result["operator"] = string(tmp)
	}

	outputLinks := []interface{}{}
	for _, item := range obj.OutputLinks {
		outputLinks = append(outputLinks, OutputLinkToMap(item))
	}
	result["output_links"] = outputLinks

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	if obj.UiProperties != nil {
		result["ui_properties"] = []interface{}{UIPropertiesToMap(obj.UiProperties)}
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToInputLink(fieldKeyFormat string) (oci_dataintegration.InputLink, error) {
	result := oci_dataintegration.InputLink{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if fieldMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_map")); ok {
		tmp := fieldMap.(string)
		var fieldMapObj oci_dataintegration.FieldMap
		err := json.Unmarshal([]byte(tmp), &fieldMapObj)
		if err != nil {
			return result, err
		}
		result.FieldMap = fieldMapObj
	}

	if fromLink, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "from_link")); ok {
		tmp := fromLink.(string)
		result.FromLink = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}
	// obj.ModelType does not exist. we need to change it when we will implement new task type
	/*if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		result.ModelType = oci_dataintegration.FlowPortLinkModelTypeEnum(modelType.(string))
	}*/

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(string)
		result.Port = &tmp
	}

	return result, nil
}

func InputLinkToMap(obj oci_dataintegration.InputLink) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.FieldMap != nil {
		tmp, _ := json.Marshal(obj.FieldMap)
		result["field_map"] = string(tmp)
	}

	if obj.FromLink != nil {
		result["from_link"] = string(*obj.FromLink)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	if obj.Port != nil {
		result["port"] = string(*obj.Port)
	}

	return result
}

func InputPortToMap(obj oci_dataintegration.InputPort) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigValues != nil {
		result["config_values"] = []interface{}{ConfigValuesToMap(obj.ConfigValues)}
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Fields != nil {
		tmp, _ := json.Marshal(obj.Fields)
		result["fields"] = string(tmp)
	}
	result["fields"] = obj.Fields

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	result["port_type"] = string(obj.PortType)

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToObjectMetadata(fieldKeyFormat string) (oci_dataintegration.ObjectMetadata, error) {
	result := oci_dataintegration.ObjectMetadata{}

	if aggregator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator")); ok {
		if tmpList := aggregator.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "aggregator"), 0)
			tmp, err := s.mapToAggregatorSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert aggregator, encountered error: %v", err)
			}
			result.Aggregator = &tmp
		}
	}

	if aggregatorKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator_key")); ok {
		tmp := aggregatorKey.(string)
		result.AggregatorKey = &tmp
	}

	if countStatistics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "count_statistics")); ok {
		if tmpList := countStatistics.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "count_statistics"), 0)
			tmp, err := s.mapToCountStatistic(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert count_statistics, encountered error: %v", err)
			}
			result.CountStatistics = &tmp
		}
	}

	if createdBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "created_by")); ok {
		tmp := createdBy.(string)
		result.CreatedBy = &tmp
	}

	if createdByName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "created_by_name")); ok {
		tmp := createdByName.(string)
		result.CreatedByName = &tmp
	}

	if identifierPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier_path")); ok {
		tmp := identifierPath.(string)
		result.IdentifierPath = &tmp
	}

	if infoFields, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "info_fields")); ok {
		result.InfoFields = tfresource.ObjectMapToStringMap(infoFields.(map[string]interface{}))
	}

	if isFavorite, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_favorite")); ok {
		tmp := isFavorite.(bool)
		result.IsFavorite = &tmp
	}

	if labels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "labels")); ok {
		interfaces := labels.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "labels")) {
			result.Labels = tmp
		}
	}

	if registryVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "registry_version")); ok {
		tmp := registryVersion.(int)
		result.RegistryVersion = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_updated")); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdated.(string))
		if err != nil {
			return result, err
		}
		result.TimeUpdated = &oci_common.SDKTime{Time: tmp}
	}

	if updatedBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "updated_by")); ok {
		tmp := updatedBy.(string)
		result.UpdatedBy = &tmp
	}

	if updatedByName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "updated_by_name")); ok {
		tmp := updatedByName.(string)
		result.UpdatedByName = &tmp
	}

	return result, nil
}

func DataIntegration_Task_ObjectMetadataToMap(obj *oci_dataintegration.ObjectMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Aggregator != nil {
		result["aggregator"] = []interface{}{DataIntegration_Task_AggregatorSummaryToMap(obj.Aggregator)}
	}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.CountStatistics != nil {
		result["count_statistics"] = []interface{}{DataIntegration_Task_CountStatisticToMap(obj.CountStatistics)}
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.CreatedByName != nil {
		result["created_by_name"] = string(*obj.CreatedByName)
	}

	if obj.IdentifierPath != nil {
		result["identifier_path"] = string(*obj.IdentifierPath)
	}

	result["info_fields"] = obj.InfoFields
	result["info_fields"] = obj.InfoFields

	if obj.IsFavorite != nil {
		result["is_favorite"] = bool(*obj.IsFavorite)
	}

	result["labels"] = obj.Labels
	result["labels"] = obj.Labels

	if obj.RegistryVersion != nil {
		result["registry_version"] = int(*obj.RegistryVersion)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.Format(time.RFC3339Nano)
	}

	if obj.UpdatedBy != nil {
		result["updated_by"] = string(*obj.UpdatedBy)
	}

	if obj.UpdatedByName != nil {
		result["updated_by_name"] = string(*obj.UpdatedByName)
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToOutputLink(fieldKeyFormat string) (oci_dataintegration.OutputLink, error) {
	result := oci_dataintegration.OutputLink{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(string)
		result.Port = &tmp
	}

	if toLinks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "to_links")); ok {
		interfaces := toLinks.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "to_links")) {
			result.ToLinks = tmp
		}
	}

	return result, nil
}

func OutputLinkToMap(obj oci_dataintegration.OutputLink) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	if obj.Port != nil {
		result["port"] = string(*obj.Port)
	}

	result["to_links"] = obj.ToLinks
	result["to_links"] = obj.ToLinks

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToOutputPort(fieldKeyFormat string) (oci_dataintegration.OutputPort, error) {
	result := oci_dataintegration.OutputPort{}

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
			}
			result.ConfigValues = &tmp
		}
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if fields, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fields")); ok {
		tmp := strings.Join(fields.([]string), " ")
		var fieldsObj []oci_dataintegration.TypedObject
		err := json.Unmarshal([]byte(tmp), &fieldsObj)
		if err != nil {
			return result, err
		}
		result.Fields = fieldsObj
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if portType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port_type")); ok {
		result.PortType = oci_dataintegration.OutputPortPortTypeEnum(portType.(string))
	}

	return result, nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToInputPort(fieldKeyFormat string) (oci_dataintegration.InputPort, error) {
	result := oci_dataintegration.InputPort{}

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
			}
			result.ConfigValues = &tmp
		}
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if fields, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fields")); ok {
		tmp := strings.Join(fields.([]string), " ")
		var fieldsObj []oci_dataintegration.TypedObject
		err := json.Unmarshal([]byte(tmp), &fieldsObj)
		if err != nil {
			return result, err
		}
		result.Fields = fieldsObj
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	// obj.ModelType does not exist. we need to change it when we will implement new task type
	/* if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		result.ModelType = oci_dataintegration.TypedObjectModelTypeEnum(modelType.(string))
	}*/

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if portType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port_type")); ok {
		result.PortType = oci_dataintegration.InputPortPortTypeEnum(portType.(string))
	}

	return result, nil
}

func OutputPortToMap(obj oci_dataintegration.OutputPort) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigValues != nil {
		result["config_values"] = []interface{}{ConfigValuesToMap(obj.ConfigValues)}
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Fields != nil {
		tmp, _ := json.Marshal(obj.Fields)
		result["fields"] = string(tmp)
	}
	result["fields"] = obj.Fields

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	// obj.ModelType does not exist. we need to change it when we will implement new task type
	//result["model_type"] = string(obj.ModelType)

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	result["port_type"] = string(obj.PortType)

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToParameter(fieldKeyFormat string) (oci_dataintegration.Parameter, error) {
	result := oci_dataintegration.Parameter{}

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
			}
			result.ConfigValues = &tmp
		}
	}

	if defaultValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_value")); ok {
		result.DefaultValue = &defaultValue
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if isInput, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_input")); ok {
		tmp := isInput.(bool)
		result.IsInput = &tmp
	}

	if isOutput, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_output")); ok {
		tmp := isOutput.(bool)
		result.IsOutput = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		temp := key.(string)
		if temp != "" {
			result.Key = &temp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if outputAggregationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_aggregation_type")); ok {
		result.OutputAggregationType = oci_dataintegration.ParameterOutputAggregationTypeEnum(outputAggregationType.(string))
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = &type_
	}

	if typeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type_name")); ok {
		tmp := typeName.(string)
		result.TypeName = &tmp
	}

	if usedFor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "used_for")); ok {
		tmp := usedFor.(string)
		result.UsedFor = &tmp
	}

	return result, nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToDataType(fieldKeyFormat string) (interface{}, error) {
	result := map[string]string{}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result["model_type"] = tmp
	}

	if dtType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dt_type")); ok {
		tmp := dtType.(string)
		result["dt_type"] = tmp
	}

	if typeSystemName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type_system_name")); ok {
		tmp := typeSystemName.(string)
		result["type_system_name"] = tmp
	}

	if typeSystemName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "java_type_name")); ok {
		tmp := typeSystemName.(string)
		result["java_type_name"] = tmp
	}
	return result, nil
}

func ParameterToMap(obj oci_dataintegration.Parameter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigValues != nil {
		result["config_values"] = []interface{}{ConfigValuesToMap(obj.ConfigValues)}
	}

	if obj.DefaultValue != nil {
		result["default_value"] = obj.DefaultValue
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsInput != nil {
		result["is_input"] = bool(*obj.IsInput)
	}

	if obj.IsOutput != nil {
		result["is_output"] = bool(*obj.IsOutput)
	}

	if obj.Key != nil {
		tmp := *obj.Key
		log.Printf("Inside Parameter with key  %v", tmp)
		result["key"] = string(tmp)
	}

	result["model_type"] = "PARAMETER"

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	result["output_aggregation_type"] = string(obj.OutputAggregationType)

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	if obj.RootObjectDefaultValue != nil {
		result["root_object_default_value"] = obj.RootObjectDefaultValue
	}

	if obj.Type != nil {
		result["type"] = obj.Type
	}

	if obj.TypeName != nil {
		result["type_name"] = string(*obj.TypeName)
	}

	if obj.UsedFor != nil {
		result["used_for"] = string(*obj.UsedFor)
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToParentReference(fieldKeyFormat string) (oci_dataintegration.ParentReference, error) {
	result := oci_dataintegration.ParentReference{}

	if parent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent")); ok {
		tmp := parent.(string)
		result.Parent = &tmp
	}

	return result, nil
}

func DataIntegration_Task_ParentReferenceToMap(obj *oci_dataintegration.ParentReference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Parent != nil {
		result["parent"] = string(*obj.Parent)
	}

	if obj.RootDocId != nil {
		result["root_doc_id"] = string(*obj.RootDocId)
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToPipeline(fieldKeyFormat string) (oci_dataintegration.Pipeline, error) {
	result := oci_dataintegration.Pipeline{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if flowConfigValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "flow_config_values")); ok {
		if tmpList := flowConfigValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "flow_config_values"), 0)
			tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert flow_config_values, encountered error: %v", err)
			}
			result.FlowConfigValues = &tmp
		}
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "metadata"), 0)
			tmp, err := s.mapToObjectMetadata(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert metadata, encountered error: %v", err)
			}
			result.Metadata = &tmp
		}
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if nodes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nodes")); ok {
		interfaces := nodes.([]interface{})
		tmp := make([]oci_dataintegration.FlowNode, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "nodes"), stateDataIndex)
			converted, err := s.mapToFlowNode(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nodes")) {
			result.Nodes = tmp
		}
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_version")); ok {
		tmp := objectVersion.(int)
		result.ObjectVersion = &tmp
	}

	if parameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parameters")); ok {
		interfaces := parameters.([]interface{})
		tmp := make([]oci_dataintegration.Parameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parameters"), stateDataIndex)
			converted, err := s.mapToParameter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "parameters")) {
			result.Parameters = tmp
		}
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if variables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "variables")); ok {
		interfaces := variables.([]interface{})
		tmp := make([]oci_dataintegration.Variable, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "variables"), stateDataIndex)
			converted, err := s.mapToVariable(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "variables")) {
			result.Variables = tmp
		}
	}

	return result, nil
}

func PipelineToMap(obj *oci_dataintegration.Pipeline) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.FlowConfigValues != nil {
		result["flow_config_values"] = []interface{}{ConfigValuesToMap(obj.FlowConfigValues)}
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{DataIntegration_Task_ObjectMetadataToMap(obj.Metadata)}
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	nodes := []interface{}{}
	for _, item := range obj.Nodes {
		nodes = append(nodes, FlowNodeToMap(item))
	}
	result["nodes"] = nodes

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	parameters := []interface{}{}
	for _, item := range obj.Parameters {
		parameters = append(parameters, ParameterToMap(item))
	}
	result["parameters"] = parameters

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	variables := []interface{}{}
	for _, item := range obj.Variables {
		variables = append(variables, VariableToMap(item))
	}
	result["variables"] = variables

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToPollRestCallConfig(fieldKeyFormat string) (oci_dataintegration.PollRestCallConfig, error) {
	result := oci_dataintegration.PollRestCallConfig{}

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp, err := s.mapToPollRestCallConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
			}
			result.ConfigValues = &tmp
		}
	}

	if methodType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "method_type")); ok {
		result.MethodType = oci_dataintegration.PollRestCallConfigMethodTypeEnum(methodType.(string))
	}

	if requestHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_headers")); ok {
		result.RequestHeaders = tfresource.ObjectMapToStringMap(requestHeaders.(map[string]interface{}))
	}

	return result, nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToTypedExpressionConfigValues(fieldKeyFormat string) (oci_dataintegration.ConfigValues, error) {
	result := oci_dataintegration.ConfigValues{}

	if configParamValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_param_values")); ok {
		if tmpList := configParamValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_param_values"), 0)
			tmp := map[string]oci_dataintegration.ConfigParameterValue{}
			var err error

			if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "length")); ok {
				if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
					result := oci_dataintegration.ConfigParameterValue{}
					temp := tmpList[0].(map[string]interface{})["int_value"]
					val := temp.(int)
					result.IntValue = &val
					tmp["length"] = result
				}
			}
			if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "scale")); ok {
				if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
					result := oci_dataintegration.ConfigParameterValue{}
					temp := tmpList[0].(map[string]interface{})["int_value"]
					val := temp.(int)
					result.IntValue = &val
					tmp["scale"] = result
				}
			}
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ConfigParamValues = tmp
		}
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	return result, nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToPollRestCallConfigValues(fieldKeyFormat string) (oci_dataintegration.ConfigValues, error) {
	result := oci_dataintegration.ConfigValues{}

	if configParamValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_param_values")); ok {
		if tmpList := configParamValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_param_values"), 0)
			tmp := map[string]oci_dataintegration.ConfigParameterValue{}
			var err error
			if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "request_url")); ok {
				if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
					result := oci_dataintegration.ConfigParameterValue{}
					result.StringValue = ObjectMapToStringValue(tmpList[0].(map[string]interface{}))
					tmp["requestURL"] = result
				}
			}
			if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "poll_max_duration")); ok {
				if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
					result := oci_dataintegration.ConfigParameterValue{}
					result.ObjectValue = ObjectMapToObjectvalue(tmpList[0].(map[string]interface{}))
					tmp["pollMaxDuration"] = result
				}
			}
			if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "poll_max_duration_unit")); ok {
				if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
					result := oci_dataintegration.ConfigParameterValue{}
					result.StringValue = ObjectMapToStringValue(tmpList[0].(map[string]interface{}))
					tmp["pollMaxDurationUnit"] = result
				}
			}
			if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "poll_interval")); ok {
				if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
					result := oci_dataintegration.ConfigParameterValue{}
					result.ObjectValue = ObjectMapToObjectvalue(tmpList[0].(map[string]interface{}))
					tmp["pollInterval"] = result
				}
			}
			if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "poll_interval_unit")); ok {
				if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
					result := oci_dataintegration.ConfigParameterValue{}
					result.StringValue = ObjectMapToStringValue(tmpList[0].(map[string]interface{}))
					tmp["pollIntervalUnit"] = result
				}
			}
			if configParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormatNextLevel, "poll_condition")); ok {
				if tmpList := configParamValue.([]interface{}); len(tmpList) > 0 {
					tmp["pollCondition"], err = s.mapToPollCondition(tmpList)
				}
			}
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ConfigParamValues = tmp
		}
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	return result, nil
}

func PollRestCallConfigToMap(obj *oci_dataintegration.PollRestCallConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigValues != nil {
		result["config_values"] = []interface{}{ConfigValuesToMapForPollRest(obj.ConfigValues)}
	}

	result["method_type"] = string(obj.MethodType)
	result["model_type"] = "POLL_REST_CALL_CONFIG"
	if obj.RequestHeaders != nil {
		result["request_headers"] = obj.RequestHeaders
	}

	return result
}

func ConfigValuesToMapForPollRest(obj *oci_dataintegration.ConfigValues) map[string]interface{} {
	result := map[string]interface{}{}

	result["config_param_values"] = ConfigValuesMapToObjectArrayPollRest(obj.ConfigParamValues)

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	return result
}

func ConfigValuesMapToObjectArrayPollRest(configParamValues map[string]oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	requestUrl, ok := configParamValues["requestURL"]
	if ok {
		result[0]["request_url"] = ConfigValueMapToRequestUrlRest(requestUrl)
	}
	tmp, ok := configParamValues["pollMaxDuration"]
	if ok {
		result[0]["poll_max_duration"] = ConfigValueMapToPollMaxDurationRest(tmp)
	}
	tmp, ok = configParamValues["pollMaxDurationUnit"]
	if ok {
		result[0]["poll_max_duration_unit"] = ConfigValueMapToPollMaxDurationUnitRest(tmp)
	}
	tmp, ok = configParamValues["pollInterval"]
	if ok {
		result[0]["poll_interval"] = ConfigValueMapToPollIntervalRest(tmp)
	}
	tmp, ok = configParamValues["pollIntervalUnit"]
	if ok {
		result[0]["poll_interval_unit"] = ConfigValueMapToPollIntervalUnitRest(tmp)
	}
	tmp, ok = configParamValues["pollCondition"]
	if ok {
		result[0]["poll_condition"] = ConfigValueMapToPollConditionRest(tmp)
	}

	return result
}

func ConfigValueMapToPollMaxDurationRest(configParameterValue oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	if configParameterValue.ObjectValue != nil {
		result[0]["object_value"] = *configParameterValue.ObjectValue
	}
	return result
}

func ConfigValueMapToPollMaxDurationUnitRest(configParameterValue oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	if configParameterValue.StringValue != nil {
		result[0]["string_value"] = *configParameterValue.StringValue
	}
	return result
}

func ConfigValueMapToPollIntervalRest(configParameterValue oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	if configParameterValue.ObjectValue != nil {
		result[0]["object_value"] = *configParameterValue.ObjectValue
	}
	return result
}

func ConfigValueMapToPollIntervalUnitRest(configParameterValue oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	if configParameterValue.StringValue != nil {
		result[0]["string_value"] = *configParameterValue.StringValue
	}
	return result
}

func ConfigValueMapToPollConditionRest(configParameterValue oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	if configParameterValue.ParameterValue != nil {
		tmp := *configParameterValue.ParameterValue
		if len(tmp) > 0 && tmp != "" {
			log.Printf("Should never come here for ParameterValue := %v", tmp)
			result[0]["parameter_value"] = tmp
		}
	}
	if configParameterValue.RefValue != nil {
		result[0]["ref_value"] = RefObjectToMap(*configParameterValue.RefValue)
	}
	return result
}

func RefObjectToMap(refValue interface{}) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})
	temp := refValue.(map[string]interface{})
	tmp, ok := temp["modelType"]
	if ok {
		result[0]["model_type"] = tmp
	}
	tmp, ok = temp["exprString"]
	if ok {
		result[0]["expr_string"] = tmp
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToProjectionRule(fieldKeyFormat string) (oci_dataintegration.ProjectionRule, error) {
	var baseObject oci_dataintegration.ProjectionRule
	//discriminator
	modelTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type"))
	var modelType string
	if ok {
		modelType = modelTypeRaw.(string)
	} else {
		modelType = "" // default value
	}
	switch strings.ToLower(modelType) {
	case strings.ToLower("GROUPED_NAME_PATTERN_RULE"):
		details := oci_dataintegration.GroupedNamePatternRule{}
		if isCascade, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_cascade")); ok {
			tmp := isCascade.(bool)
			details.IsCascade = &tmp
		}
		if isCaseSensitive, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_case_sensitive")); ok {
			tmp := isCaseSensitive.(bool)
			details.IsCaseSensitive = &tmp
		}
		if isSkipRemainingRulesOnMatch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_remaining_rules_on_match")); ok {
			tmp := isSkipRemainingRulesOnMatch.(bool)
			details.IsSkipRemainingRulesOnMatch = &tmp
		}
		if matchingStrategy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "matching_strategy")); ok {
			details.MatchingStrategy = oci_dataintegration.GroupedNamePatternRuleMatchingStrategyEnum(matchingStrategy.(string))
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if pattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pattern")); ok {
			tmp := pattern.(string)
			details.Pattern = &tmp
		}
		if ruleType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_type")); ok {
			details.RuleType = oci_dataintegration.GroupedNamePatternRuleRuleTypeEnum(ruleType.(string))
		}
		if scope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope")); ok {
			//mapToobject method does not exist. we need to change it when we will implement new task type
			/*if tmpList := scope.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scope"), 0)
				tmp, err := s.mapToobject(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scope, encountered error: %v", err)
				}/

				details.Scope = &tmpList[0]
			}*/
			details.Scope = &scope
		}
		if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
			if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
				}
				details.ConfigValues = &tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if isJavaRegexSyntax, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_java_regex_syntax")); ok {
			tmp := isJavaRegexSyntax.(bool)
			details.IsJavaRegexSyntax = &tmp
		}
		if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
				tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
				}
				details.ParentRef = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("NAME_LIST_RULE"):
		details := oci_dataintegration.NameListRule{}
		if isCascade, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_cascade")); ok {
			tmp := isCascade.(bool)
			details.IsCascade = &tmp
		}
		if isCaseSensitive, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_case_sensitive")); ok {
			tmp := isCaseSensitive.(bool)
			details.IsCaseSensitive = &tmp
		}
		if isSkipRemainingRulesOnMatch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_remaining_rules_on_match")); ok {
			tmp := isSkipRemainingRulesOnMatch.(bool)
			details.IsSkipRemainingRulesOnMatch = &tmp
		}
		if matchingStrategy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "matching_strategy")); ok {
			details.MatchingStrategy = oci_dataintegration.NameListRuleMatchingStrategyEnum(matchingStrategy.(string))
		}
		if names, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "names")); ok {
			interfaces := names.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "names")) {
				details.Names = tmp
			}
		}
		if ruleType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_type")); ok {
			details.RuleType = oci_dataintegration.NameListRuleRuleTypeEnum(ruleType.(string))
		}
		if scope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope")); ok {
			//mapToobject method does not exist. we need to change it when we will implement new task type
			/*if tmpList := scope.([]interface{}); len(tmpList) > 0 {
				/*fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scope"), 0)
				tmp, err := s.mapToobject(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scope, encountered error: %v", err)
				}

				tmp := scope.(interface{})

			}*/
			details.Scope = &scope
		}
		if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
			if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
				}
				details.ConfigValues = &tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if isJavaRegexSyntax, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_java_regex_syntax")); ok {
			tmp := isJavaRegexSyntax.(bool)
			details.IsJavaRegexSyntax = &tmp
		}
		if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
				tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
				}
				details.ParentRef = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("NAME_PATTERN_RULE"):
		details := oci_dataintegration.NamePatternRule{}
		if isCascade, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_cascade")); ok {
			tmp := isCascade.(bool)
			details.IsCascade = &tmp
		}
		if isCaseSensitive, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_case_sensitive")); ok {
			tmp := isCaseSensitive.(bool)
			details.IsCaseSensitive = &tmp
		}
		if isSkipRemainingRulesOnMatch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_remaining_rules_on_match")); ok {
			tmp := isSkipRemainingRulesOnMatch.(bool)
			details.IsSkipRemainingRulesOnMatch = &tmp
		}
		if matchingStrategy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "matching_strategy")); ok {
			details.MatchingStrategy = oci_dataintegration.NamePatternRuleMatchingStrategyEnum(matchingStrategy.(string))
		}
		if pattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pattern")); ok {
			tmp := pattern.(string)
			details.Pattern = &tmp
		}
		if ruleType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_type")); ok {
			details.RuleType = oci_dataintegration.NamePatternRuleRuleTypeEnum(ruleType.(string))
		}
		if scope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope")); ok {
			//mapToobject method does not exist. we need to change it when we will implement new task type
			/*if tmpList := scope.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scope"), 0)
				tmp, err := s.mapToobject(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scope, encountered error: %v", err)
				}

				tmp := scope.(interface{})

			}*/
			details.Scope = &scope
		}
		if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
			if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
				}
				details.ConfigValues = &tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if isJavaRegexSyntax, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_java_regex_syntax")); ok {
			tmp := isJavaRegexSyntax.(bool)
			details.IsJavaRegexSyntax = &tmp
		}
		if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
				tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
				}
				details.ParentRef = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("RENAME_RULE"):
		details := oci_dataintegration.RenameRule{}
		if fromName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "from_name")); ok {
			tmp := fromName.(string)
			details.FromName = &tmp
		}
		if isSkipRemainingRulesOnMatch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_remaining_rules_on_match")); ok {
			tmp := isSkipRemainingRulesOnMatch.(bool)
			details.IsSkipRemainingRulesOnMatch = &tmp
		}
		if toName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "to_name")); ok {
			tmp := toName.(string)
			details.ToName = &tmp
		}
		if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
			if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
				}
				details.ConfigValues = &tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if isJavaRegexSyntax, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_java_regex_syntax")); ok {
			tmp := isJavaRegexSyntax.(bool)
			details.IsJavaRegexSyntax = &tmp
		}
		if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
				tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
				}
				details.ParentRef = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("TYPED_NAME_PATTERN_RULE"):
		details := oci_dataintegration.TypedNamePatternRule{}
		if isCascade, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_cascade")); ok {
			tmp := isCascade.(bool)
			details.IsCascade = &tmp
		}
		if isCaseSensitive, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_case_sensitive")); ok {
			tmp := isCaseSensitive.(bool)
			details.IsCaseSensitive = &tmp
		}
		if isSkipRemainingRulesOnMatch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_remaining_rules_on_match")); ok {
			tmp := isSkipRemainingRulesOnMatch.(bool)
			details.IsSkipRemainingRulesOnMatch = &tmp
		}
		if matchingStrategy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "matching_strategy")); ok {
			details.MatchingStrategy = oci_dataintegration.TypedNamePatternRuleMatchingStrategyEnum(matchingStrategy.(string))
		}
		if names, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "names")); ok {
			interfaces := names.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "names")) {
				details.Names = tmp
			}
		}
		if pattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pattern")); ok {
			tmp := pattern.(string)
			details.Pattern = &tmp
		}
		if ruleType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_type")); ok {
			details.RuleType = oci_dataintegration.TypedNamePatternRuleRuleTypeEnum(ruleType.(string))
		}
		if scope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope")); ok {
			//mapToobject method does not exist. we need to change it when we will implement new task type
			/*if tmpList := scope.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scope"), 0)
				tmp, err := s.mapToobject(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scope, encountered error: %v", err)
				}

				tmp := scope.(interface{})

			}*/
			details.Scope = &scope
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			interfaces := types.([]interface{})
			tmp := make([]interface{}, len(interfaces))
			for i := range interfaces {
				tmp[i] = interfaces[i]
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "types")) {
				details.Types = tmp
			}
		}
		if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
			if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
				}
				details.ConfigValues = &tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if isJavaRegexSyntax, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_java_regex_syntax")); ok {
			tmp := isJavaRegexSyntax.(bool)
			details.IsJavaRegexSyntax = &tmp
		}
		if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
				tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
				}
				details.ParentRef = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("TYPE_LIST_RULE"):
		details := oci_dataintegration.TypeListRule{}
		if isCascade, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_cascade")); ok {
			tmp := isCascade.(bool)
			details.IsCascade = &tmp
		}
		if isCaseSensitive, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_case_sensitive")); ok {
			tmp := isCaseSensitive.(bool)
			details.IsCaseSensitive = &tmp
		}
		if isSkipRemainingRulesOnMatch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_remaining_rules_on_match")); ok {
			tmp := isSkipRemainingRulesOnMatch.(bool)
			details.IsSkipRemainingRulesOnMatch = &tmp
		}
		if matchingStrategy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "matching_strategy")); ok {
			details.MatchingStrategy = oci_dataintegration.TypeListRuleMatchingStrategyEnum(matchingStrategy.(string))
		}
		if ruleType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_type")); ok {
			details.RuleType = oci_dataintegration.TypeListRuleRuleTypeEnum(ruleType.(string))
		}
		if scope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope")); ok {
			//mapToobject method does not exist. we need to change it when we will implement TYPE_LIST_RULE
			/*if tmpList := scope.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scope"), 0)
				tmp, err := s.mapToobject(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scope, encountered error: %v", err)
				}

				tmp := scope
				details.Scope = &tmp
			}*/
			details.Scope = &scope
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			interfaces := types.([]interface{})
			tmp := make([]interface{}, len(interfaces))
			for i := range interfaces {
				tmp[i] = interfaces[i]
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "types")) {
				details.Types = tmp
			}
		}
		if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
			if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
				}
				details.ConfigValues = &tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if isJavaRegexSyntax, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_java_regex_syntax")); ok {
			tmp := isJavaRegexSyntax.(bool)
			details.IsJavaRegexSyntax = &tmp
		}
		if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
				tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
				}
				details.ParentRef = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown model_type '%v' was specified", modelType)
	}
	return baseObject, nil
}

func ProjectionRuleToMap(obj oci_dataintegration.ProjectionRule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_dataintegration.GroupedNamePatternRule:
		result["model_type"] = "GROUPED_NAME_PATTERN_RULE"

		if v.IsCascade != nil {
			result["is_cascade"] = bool(*v.IsCascade)
		}

		if v.IsCaseSensitive != nil {
			result["is_case_sensitive"] = bool(*v.IsCaseSensitive)
		}

		if v.IsSkipRemainingRulesOnMatch != nil {
			result["is_skip_remaining_rules_on_match"] = bool(*v.IsSkipRemainingRulesOnMatch)
		}

		result["matching_strategy"] = string(v.MatchingStrategy)

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.Pattern != nil {
			result["pattern"] = string(*v.Pattern)
		}

		result["rule_type"] = string(v.RuleType)

		if v.Scope != nil {
			result["scope"] = []interface{}{v.Scope}
		}
	case oci_dataintegration.NameListRule:
		result["model_type"] = "NAME_LIST_RULE"

		if v.IsCascade != nil {
			result["is_cascade"] = bool(*v.IsCascade)
		}

		if v.IsCaseSensitive != nil {
			result["is_case_sensitive"] = bool(*v.IsCaseSensitive)
		}

		if v.IsSkipRemainingRulesOnMatch != nil {
			result["is_skip_remaining_rules_on_match"] = bool(*v.IsSkipRemainingRulesOnMatch)
		}

		result["matching_strategy"] = string(v.MatchingStrategy)

		result["names"] = v.Names
		result["names"] = v.Names

		result["rule_type"] = string(v.RuleType)

		if v.Scope != nil {
			result["scope"] = v.Scope
		}
	case oci_dataintegration.NamePatternRule:
		result["model_type"] = "NAME_PATTERN_RULE"

		if v.IsCascade != nil {
			result["is_cascade"] = bool(*v.IsCascade)
		}

		if v.IsCaseSensitive != nil {
			result["is_case_sensitive"] = bool(*v.IsCaseSensitive)
		}

		if v.IsSkipRemainingRulesOnMatch != nil {
			result["is_skip_remaining_rules_on_match"] = bool(*v.IsSkipRemainingRulesOnMatch)
		}

		result["matching_strategy"] = string(v.MatchingStrategy)

		if v.Pattern != nil {
			result["pattern"] = string(*v.Pattern)
		}

		result["rule_type"] = string(v.RuleType)

		if v.Scope != nil {
			result["scope"] = v.Scope
		}
	case oci_dataintegration.RenameRule:
		result["model_type"] = "RENAME_RULE"

		if v.FromName != nil {
			result["from_name"] = string(*v.FromName)
		}

		if v.IsSkipRemainingRulesOnMatch != nil {
			result["is_skip_remaining_rules_on_match"] = bool(*v.IsSkipRemainingRulesOnMatch)
		}

		if v.ToName != nil {
			result["to_name"] = string(*v.ToName)
		}
	case oci_dataintegration.TypedNamePatternRule:
		result["model_type"] = "TYPED_NAME_PATTERN_RULE"

		if v.IsCascade != nil {
			result["is_cascade"] = bool(*v.IsCascade)
		}

		if v.IsCaseSensitive != nil {
			result["is_case_sensitive"] = bool(*v.IsCaseSensitive)
		}

		if v.IsSkipRemainingRulesOnMatch != nil {
			result["is_skip_remaining_rules_on_match"] = bool(*v.IsSkipRemainingRulesOnMatch)
		}

		result["matching_strategy"] = string(v.MatchingStrategy)

		result["names"] = v.Names
		result["names"] = v.Names

		if v.Pattern != nil {
			result["pattern"] = string(*v.Pattern)
		}

		result["rule_type"] = string(v.RuleType)

		if v.Scope != nil {
			result["scope"] = v.Scope
		}

		types := []interface{}{}
		for _, item := range v.Types {
			types = append(types, item)
		}
		result["types"] = types
	case oci_dataintegration.TypeListRule:
		result["model_type"] = "TYPE_LIST_RULE"

		if v.IsCascade != nil {
			result["is_cascade"] = bool(*v.IsCascade)
		}

		if v.IsCaseSensitive != nil {
			result["is_case_sensitive"] = bool(*v.IsCaseSensitive)
		}

		if v.IsSkipRemainingRulesOnMatch != nil {
			result["is_skip_remaining_rules_on_match"] = bool(*v.IsSkipRemainingRulesOnMatch)
		}

		result["matching_strategy"] = string(v.MatchingStrategy)

		result["rule_type"] = string(v.RuleType)

		if v.Scope != nil {
			result["scope"] = v.Scope
		}

		types := []interface{}{}
		for _, item := range v.Types {
			types = append(types, item)
		}
		result["types"] = types
	default:
		log.Printf("[WARN] Received 'model_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToRegistryMetadata(fieldKeyFormat string) (oci_dataintegration.RegistryMetadata, error) {
	result := oci_dataintegration.RegistryMetadata{}

	if aggregatorKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator_key")); ok {
		tmp := aggregatorKey.(string)
		result.AggregatorKey = &tmp
	}

	if isFavorite, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_favorite")); ok {
		tmp := isFavorite.(bool)
		result.IsFavorite = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if labels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "labels")); ok {
		interfaces := labels.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "labels")) {
			result.Labels = tmp
		}
	}

	if registryVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "registry_version")); ok {
		tmp := registryVersion.(int)
		result.RegistryVersion = &tmp
	}

	return result, nil
}

func DataIntegration_Task_RegistryMetadataToMap(obj *oci_dataintegration.RegistryMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.IsFavorite != nil {
		result["is_favorite"] = bool(*obj.IsFavorite)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["labels"] = obj.Labels
	result["labels"] = obj.Labels

	if obj.RegistryVersion != nil {
		result["registry_version"] = int(*obj.RegistryVersion)
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToRootObject(fieldKeyFormat string) (oci_dataintegration.RootObject, error) {
	result := oci_dataintegration.RootObject{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	return result, nil
}

func RootObjectToMap(obj *oci_dataintegration.RootObject) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToScript(fieldKeyFormat string) (oci_dataintegration.Script, error) {
	result := oci_dataintegration.Script{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	return result, nil
}

func ScriptToMap(obj *oci_dataintegration.Script) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	return result
}

func TaskSummaryToMap(obj oci_dataintegration.TaskSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_dataintegration.TaskSummaryFromDataLoaderTask:
		result["model_type"] = "DATA_LOADER_TASK"

		if v.ConditionalCompositeFieldMap != nil {
			result["conditional_composite_field_map"] = []interface{}{ConditionalCompositeFieldMapToMap(v.ConditionalCompositeFieldMap)}
		}

		if v.DataFlow != nil {
			result["data_flow"] = []interface{}{DataFlowToMap(v.DataFlow)}
		}

		if v.IsSingleLoad != nil {
			result["is_single_load"] = bool(*v.IsSingleLoad)
		}

		if v.ParallelLoadLimit != nil {
			result["parallel_load_limit"] = int(*v.ParallelLoadLimit)
		}
	case oci_dataintegration.TaskSummaryFromIntegrationTask:
		result["model_type"] = "INTEGRATION_TASK"

		if v.DataFlow != nil {
			result["data_flow"] = []interface{}{DataFlowToMap(v.DataFlow)}
		}
	case oci_dataintegration.TaskSummaryFromOciDataflowTask:
		result["model_type"] = "OCI_DATAFLOW_TASK"

		if v.DataflowApplication != nil {
			result["dataflow_application"] = []interface{}{DataflowApplicationToMap(v.DataflowApplication)}
		}
	case oci_dataintegration.TaskSummaryFromPipelineTask:
		result["model_type"] = "PIPELINE_TASK"

		if v.Pipeline != nil {
			result["pipeline"] = []interface{}{PipelineToMap(v.Pipeline)}
		}
	case oci_dataintegration.TaskSummaryFromRestTask:
		result["model_type"] = "REST_TASK"

		result["api_call_mode"] = string(v.ApiCallMode)

		result["name"] = string(*v.Name)

		result["key"] = string(*v.Key)

		if v.AuthConfig != nil {
			authConfigArray := []interface{}{}
			if authConfigMap := AuthConfigToMap(&v.AuthConfig); authConfigMap != nil {
				authConfigArray = append(authConfigArray, authConfigMap)
			}
			result["auth_config"] = authConfigArray
		}
		if v.CancelRestCallConfig != nil {
			result["cancel_rest_call_config"] = []interface{}{CancelRestCallConfigToMap(v.CancelRestCallConfig)}
		}

		if v.ExecuteRestCallConfig != nil {
			result["execute_rest_call_config"] = []interface{}{ExecuteRestCallConfigToMap(v.ExecuteRestCallConfig)}
		}

		if v.PollRestCallConfig != nil {
			result["poll_rest_call_config"] = []interface{}{PollRestCallConfigToMap(v.PollRestCallConfig)}
		}
	case oci_dataintegration.TaskSummaryFromSqlTask:
		result["model_type"] = "SQL_TASK"

		if v.Operation != nil {
			result["operation"] = v.Operation
		}

		if v.Script != nil {
			result["script"] = []interface{}{ScriptToMap(v.Script)}
		}

		result["sql_script_type"] = string(v.SqlScriptType)
	default:
		log.Printf("[WARN] Received 'model_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToTypedExpression(fieldKeyFormat string) (oci_dataintegration.TypedExpression, error) {
	result := oci_dataintegration.TypedExpression{}

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp, err := s.mapToTypedExpressionConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
			}
			result.ConfigValues = &tmp
		}
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if expression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "expression")); ok {
		tmp := expression.(string)
		result.Expression = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func TypedExpressionToMap(obj oci_dataintegration.TypedExpression) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigValues != nil {
		result["config_values"] = []interface{}{ConfigValuesToMapTypedExpr(obj.ConfigValues)}
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Expression != nil {
		result["expression"] = string(*obj.Expression)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["model_type"] = "TYPED_EXPRESSION"

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func ConfigValuesToMapTypedExpr(obj *oci_dataintegration.ConfigValues) map[string]interface{} {
	result := map[string]interface{}{}

	result["config_param_values"] = ConfigValuesMapToObjectArrayTypedExpr(obj.ConfigParamValues)

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	return result
}

func ConfigValuesMapToObjectArrayTypedExpr(configParamValues map[string]oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = map[string]interface{}{}
	length, ok := configParamValues["length"]
	if ok {
		result[0]["length"] = ConfigValueMapToIntValue(length)
	}
	scale, ok := configParamValues["scale"]
	if ok {
		result[0]["scale"] = ConfigValueMapToIntValue(scale)
	}
	return result
}

func ConfigValueMapToIntValue(configParameterValue oci_dataintegration.ConfigParameterValue) []map[string]interface{} {
	var result = make([]map[string]interface{}, 1)
	result[0] = map[string]interface{}{}
	if configParameterValue.IntValue != nil {
		result[0]["int_value"] = *configParameterValue.IntValue
	}
	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToUIProperties(fieldKeyFormat string) (oci_dataintegration.UiProperties, error) {
	result := oci_dataintegration.UiProperties{}

	if coordinateX, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "coordinate_x")); ok {
		tmp := coordinateX.(float32)
		result.CoordinateX = &tmp
	}

	if coordinateY, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "coordinate_y")); ok {
		tmp := coordinateY.(float32)
		result.CoordinateY = &tmp
	}

	return result, nil
}

func UIPropertiesToMap(obj *oci_dataintegration.UiProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CoordinateX != nil {
		result["coordinate_x"] = float32(*obj.CoordinateX)
	}

	if obj.CoordinateY != nil {
		result["coordinate_y"] = float32(*obj.CoordinateY)
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) mapToVariable(fieldKeyFormat string) (oci_dataintegration.Variable, error) {
	result := oci_dataintegration.Variable{}

	if configValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_values")); ok {
		if tmpList := configValues.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_values"), 0)
			tmp, err := s.mapToConfigValues(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_values, encountered error: %v", err)
			}
			result.ConfigValues = &tmp
		}
	}

	if defaultValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_value")); ok {
		result.DefaultValue = &defaultValue
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_version")); ok {
		tmp := objectVersion.(int)
		result.ObjectVersion = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if rootObjectDefaultValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "root_object_default_value")); ok {
		if tmpList := rootObjectDefaultValue.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "root_object_default_value"), 0)
			tmp, err := s.mapToRootObject(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert root_object_default_value, encountered error: %v", err)
			}
			result.RootObjectDefaultValue = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		var typeObj oci_dataintegration.BaseType
		err := json.Unmarshal([]byte(tmp), &typeObj)
		if err != nil {
			return result, err
		}
		result.Type = typeObj
	}

	return result, nil
}

func VariableToMap(obj oci_dataintegration.Variable) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigValues != nil {
		result["config_values"] = []interface{}{ConfigValuesToMap(obj.ConfigValues)}
	}

	if obj.DefaultValue != nil {
		result["default_value"] = obj.DefaultValue
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataIntegration_Task_ParentReferenceToMap(obj.ParentRef)}
	}

	if obj.RootObjectDefaultValue != nil {
		result["root_object_default_value"] = []interface{}{RootObjectToMap(obj.RootObjectDefaultValue)}
	}

	if obj.Type != nil {
		tmp, _ := json.Marshal(obj.Type)
		result["type"] = string(tmp)
	}

	return result
}

func (s *DataintegrationWorkspaceTaskResourceCrud) populateTopLevelPolymorphicCreateTaskRequest(request *oci_dataintegration.CreateTaskRequest) error {
	//discriminator
	modelTypeRaw, ok := s.D.GetOkExists("model_type")
	var modelType string
	if ok {
		modelType = modelTypeRaw.(string)
	} else {
		modelType = "" // default value
	}
	switch strings.ToLower(modelType) {
	case strings.ToLower("DATA_LOADER_TASK"):
		details := oci_dataintegration.CreateTaskFromDataLoaderTask{}
		if conditionalCompositeFieldMap, ok := s.D.GetOkExists("conditional_composite_field_map"); ok {
			if tmpList := conditionalCompositeFieldMap.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "conditional_composite_field_map", 0)
				tmp, err := s.mapToConditionalCompositeFieldMap(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConditionalCompositeFieldMap = &tmp
			}
		}
		if dataFlow, ok := s.D.GetOkExists("data_flow"); ok {
			if tmpList := dataFlow.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_flow", 0)
				tmp, err := s.mapToDataFlow(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataFlow = &tmp
			}
		}
		if isSingleLoad, ok := s.D.GetOkExists("is_single_load"); ok {
			tmp := isSingleLoad.(bool)
			details.IsSingleLoad = &tmp
		}
		if parallelLoadLimit, ok := s.D.GetOkExists("parallel_load_limit"); ok {
			tmp := parallelLoadLimit.(int)
			details.ParallelLoadLimit = &tmp
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToCreateConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}
		request.CreateTaskDetails = details
	case strings.ToLower("INTEGRATION_TASK"):
		details := oci_dataintegration.CreateTaskFromIntegrationTask{}
		if dataFlow, ok := s.D.GetOkExists("data_flow"); ok {
			if tmpList := dataFlow.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_flow", 0)
				tmp, err := s.mapToDataFlow(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataFlow = &tmp
			}
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToCreateConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}
		request.CreateTaskDetails = details
	case strings.ToLower("OCI_DATAFLOW_TASK"):
		details := oci_dataintegration.CreateTaskFromOciDataflowTask{}
		if dataflowApplication, ok := s.D.GetOkExists("dataflow_application"); ok {
			if tmpList := dataflowApplication.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dataflow_application", 0)
				tmp, err := s.mapToDataflowApplication(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataflowApplication = &tmp
			}
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToCreateConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}
		request.CreateTaskDetails = details
	case strings.ToLower("PIPELINE_TASK"):
		details := oci_dataintegration.CreateTaskFromPipelineTask{}
		if pipeline, ok := s.D.GetOkExists("pipeline"); ok {
			if tmpList := pipeline.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "pipeline", 0)
				tmp, err := s.mapToPipeline(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Pipeline = &tmp
			}
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToCreateConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}
		request.CreateTaskDetails = details
	case strings.ToLower("REST_TASK"):
		details := oci_dataintegration.CreateTaskFromRestTask{}
		if apiCallMode, ok := s.D.GetOkExists("api_call_mode"); ok {
			details.ApiCallMode = oci_dataintegration.CreateTaskFromRestTaskApiCallModeEnum(apiCallMode.(string))
		}
		if authConfig, ok := s.D.GetOkExists("auth_config"); ok {
			if tmpList := authConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "auth_config", 0)
				tmp, err := s.mapToAuthConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AuthConfig = tmp
			}
		}

		if cancelRestCallConfig, ok := s.D.GetOkExists("cancel_rest_call_config"); ok {
			if tmpList := cancelRestCallConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cancel_rest_call_config", 0)
				tmp, err := s.mapToCancelRestCallConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.CancelRestCallConfig = &tmp
			}
		}
		if executeRestCallConfig, ok := s.D.GetOkExists("execute_rest_call_config"); ok {
			if tmpList := executeRestCallConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "execute_rest_call_config", 0)
				tmp, err := s.mapToExecuteRestCallConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ExecuteRestCallConfig = &tmp
			}
		}
		if pollRestCallConfig, ok := s.D.GetOkExists("poll_rest_call_config"); ok {
			if tmpList := pollRestCallConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "poll_rest_call_config", 0)
				tmp, err := s.mapToPollRestCallConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.PollRestCallConfig = &tmp
			}
		}
		if typedExpressions, ok := s.D.GetOkExists("typed_expressions"); ok {
			interfaces := typedExpressions.([]interface{})
			tmp := make([]oci_dataintegration.TypedExpression, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "typed_expressions", stateDataIndex)
				converted, err := s.mapToTypedExpression(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("typed_expressions") {
				details.TypedExpressions = tmp
			}
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToCreateConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}
		request.CreateTaskDetails = details
	case strings.ToLower("SQL_TASK"):
		details := oci_dataintegration.CreateTaskFromSqlTask{}
		if operation, ok := s.D.GetOkExists("operation"); ok {
			//mapToobject method does not exist. we need to change it when we will implement SqlTask
			/*if tmpList := operation.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "operation", 0)
				tmp, err := s.mapToobject(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp := operation
				details.Operation = &tmp
			}*/
			details.Operation = &operation
		}
		if script, ok := s.D.GetOkExists("script"); ok {
			if tmpList := script.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "script", 0)
				tmp, err := s.mapToScript(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Script = &tmp
			}
		}
		if sqlScriptType, ok := s.D.GetOkExists("sql_script_type"); ok {
			details.SqlScriptType = oci_dataintegration.CreateTaskFromSqlTaskSqlScriptTypeEnum(sqlScriptType.(string))
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToCreateConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}
		request.CreateTaskDetails = details
	default:
		return fmt.Errorf("unknown model_type '%v' was specified", modelType)
	}
	return nil
}

func (s *DataintegrationWorkspaceTaskResourceCrud) populateTopLevelPolymorphicUpdateTaskRequest(request *oci_dataintegration.UpdateTaskRequest) error {
	//discriminator
	modelTypeRaw, ok := s.D.GetOkExists("model_type")
	var modelType string
	if ok {
		modelType = modelTypeRaw.(string)
	} else {
		modelType = "" // default value
	}
	switch strings.ToLower(modelType) {
	case strings.ToLower("DATA_LOADER_TASK"):
		details := oci_dataintegration.UpdateTaskFromDataLoaderTask{}
		if conditionalCompositeFieldMap, ok := s.D.GetOkExists("conditional_composite_field_map"); ok {
			if tmpList := conditionalCompositeFieldMap.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "conditional_composite_field_map", 0)
				tmp, err := s.mapToConditionalCompositeFieldMap(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConditionalCompositeFieldMap = &tmp
			}
		}
		if dataFlow, ok := s.D.GetOkExists("data_flow"); ok {
			if tmpList := dataFlow.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_flow", 0)
				tmp, err := s.mapToDataFlow(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataFlow = &tmp
			}
		}
		if isSingleLoad, ok := s.D.GetOkExists("is_single_load"); ok {
			tmp := isSingleLoad.(bool)
			details.IsSingleLoad = &tmp
		}
		if parallelLoadLimit, ok := s.D.GetOkExists("parallel_load_limit"); ok {
			tmp := parallelLoadLimit.(int)
			details.ParallelLoadLimit = &tmp
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
			tmp := objectVersion.(int)
			details.ObjectVersion = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if taskKey, ok := s.D.GetOkExists("task_key"); ok {
			tmp := taskKey.(string)
			request.TaskKey = &tmp
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}
		request.UpdateTaskDetails = details
	case strings.ToLower("INTEGRATION_TASK"):
		details := oci_dataintegration.UpdateTaskFromIntegrationTask{}
		if dataFlow, ok := s.D.GetOkExists("data_flow"); ok {
			if tmpList := dataFlow.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_flow", 0)
				tmp, err := s.mapToDataFlow(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataFlow = &tmp
			}
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
			tmp := objectVersion.(int)
			details.ObjectVersion = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if taskKey, ok := s.D.GetOkExists("task_key"); ok {
			tmp := taskKey.(string)
			request.TaskKey = &tmp
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}
		request.UpdateTaskDetails = details
	case strings.ToLower("OCI_DATAFLOW_TASK"):
		details := oci_dataintegration.UpdateTaskFromOciDataflowTask{}
		if dataflowApplication, ok := s.D.GetOkExists("dataflow_application"); ok {
			if tmpList := dataflowApplication.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dataflow_application", 0)
				tmp, err := s.mapToDataflowApplication(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataflowApplication = &tmp
			}
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
			tmp := objectVersion.(int)
			details.ObjectVersion = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if taskKey, ok := s.D.GetOkExists("task_key"); ok {
			tmp := taskKey.(string)
			request.TaskKey = &tmp
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}
		request.UpdateTaskDetails = details
	case strings.ToLower("PIPELINE_TASK"):
		details := oci_dataintegration.UpdateTaskFromPipelineTask{}
		if pipeline, ok := s.D.GetOkExists("pipeline"); ok {
			if tmpList := pipeline.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "pipeline", 0)
				tmp, err := s.mapToPipeline(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Pipeline = &tmp
			}
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
			tmp := objectVersion.(int)
			details.ObjectVersion = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if taskKey, ok := s.D.GetOkExists("task_key"); ok {
			tmp := taskKey.(string)
			request.TaskKey = &tmp
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}

		request.UpdateTaskDetails = details
	case strings.ToLower("REST_TASK"):
		details := oci_dataintegration.UpdateTaskFromRestTask{}
		if additionalProperties, ok := s.D.GetOkExists("additional_properties"); ok {
			tmp := additionalProperties.(string)
			details.AdditionalProperties = &tmp
		}
		if apiCallMode, ok := s.D.GetOkExists("api_call_mode"); ok {
			details.ApiCallMode = oci_dataintegration.UpdateTaskFromRestTaskApiCallModeEnum(apiCallMode.(string))
		}
		if authConfig, ok := s.D.GetOkExists("auth_config"); ok {
			if tmpList := authConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "auth_config", 0)
				tmp, err := s.mapToAuthConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AuthConfig = tmp
			}
		}
		if cancelRestCallConfig, ok := s.D.GetOkExists("cancel_rest_call_config"); ok {
			if tmpList := cancelRestCallConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cancel_rest_call_config", 0)
				tmp, err := s.mapToCancelRestCallConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.CancelRestCallConfig = &tmp
			}
		}
		if executeRestCallConfig, ok := s.D.GetOkExists("execute_rest_call_config"); ok {
			if tmpList := executeRestCallConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "execute_rest_call_config", 0)
				tmp, err := s.mapToExecuteRestCallConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ExecuteRestCallConfig = &tmp
			}
		}
		if pollRestCallConfig, ok := s.D.GetOkExists("poll_rest_call_config"); ok {
			if tmpList := pollRestCallConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "poll_rest_call_config", 0)
				tmp, err := s.mapToPollRestCallConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.PollRestCallConfig = &tmp
			}
		}
		if typedExpressions, ok := s.D.GetOkExists("typed_expressions"); ok {
			interfaces := typedExpressions.([]interface{})
			tmp := make([]oci_dataintegration.TypedExpression, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "typed_expressions", stateDataIndex)
				converted, err := s.mapToTypedExpression(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("typed_expressions") {
				details.TypedExpressions = tmp
			}
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
			request.TaskKey = &tmp

		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
			tmp := objectVersion.(int)
			details.ObjectVersion = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}
		request.UpdateTaskDetails = details
	case strings.ToLower("SQL_TASK"):
		details := oci_dataintegration.UpdateTaskFromSqlTask{}
		if operation, ok := s.D.GetOkExists("operation"); ok {
			//mapToobject method does not exist. we need to change it when we will implement SqlTask

			/*if tmpList := operation.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "operation", 0)
				tmp, err := s.mapToobject(fieldKeyFormat)
				if err != nil {
					return err
				}

			}*/
			details.Operation = &operation
		}
		if script, ok := s.D.GetOkExists("script"); ok {
			if tmpList := script.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "script", 0)
				tmp, err := s.mapToScript(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Script = &tmp
			}
		}
		if sqlScriptType, ok := s.D.GetOkExists("sql_script_type"); ok {
			details.SqlScriptType = oci_dataintegration.UpdateTaskFromSqlTaskSqlScriptTypeEnum(sqlScriptType.(string))
		}
		if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
			if tmpList := configProviderDelegate.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_provider_delegate", 0)
				tmp, err := s.mapToConfigProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConfigProviderDelegate = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if identifier, ok := s.D.GetOkExists("identifier"); ok {
			tmp := identifier.(string)
			details.Identifier = &tmp
		}
		if inputPorts, ok := s.D.GetOkExists("input_ports"); ok {
			interfaces := inputPorts.([]interface{})
			tmp := make([]oci_dataintegration.InputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_ports", stateDataIndex)
				converted, err := s.mapToInputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("input_ports") {
				details.InputPorts = tmp
			}
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
			tmp := modelVersion.(string)
			details.ModelVersion = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
			tmp := objectStatus.(int)
			details.ObjectStatus = &tmp
		}
		if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
			tmp := objectVersion.(int)
			details.ObjectVersion = &tmp
		}
		if opConfigValues, ok := s.D.GetOkExists("op_config_values"); ok {
			if tmpList := opConfigValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "op_config_values", 0)
				tmp, err := s.mapToConfigValues(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.OpConfigValues = &tmp
			}
		}
		if outputPorts, ok := s.D.GetOkExists("output_ports"); ok {
			interfaces := outputPorts.([]interface{})
			tmp := make([]oci_dataintegration.OutputPort, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_ports", stateDataIndex)
				converted, err := s.mapToOutputPort(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("output_ports") {
				details.OutputPorts = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			interfaces := parameters.([]interface{})
			tmp := make([]oci_dataintegration.Parameter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
				converted, err := s.mapToParameter(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("parameters") {
				details.Parameters = tmp
			}
		}
		if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
			if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
				tmp, err := s.mapToParentReference(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ParentRef = &tmp
			}
		}
		if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
			if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
				tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RegistryMetadata = &tmp
			}
		}
		if taskKey, ok := s.D.GetOkExists("task_key"); ok {
			tmp := taskKey.(string)
			request.TaskKey = &tmp
		}
		if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
			tmp := workspaceId.(string)
			request.WorkspaceId = &tmp
		}
		request.UpdateTaskDetails = details
	default:
		return fmt.Errorf("unknown model_type '%v' was specified", modelType)
	}
	return nil
}

func DataintegrationTaskObjectMetadataToRegistryMetadataMap(obj *oci_dataintegration.ObjectMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.RegistryVersion != nil {
		result["registry_version"] = int(*obj.RegistryVersion)
	}

	if obj.IsFavorite != nil {
		result["is_favorite"] = bool(*obj.IsFavorite)
	}

	result["labels"] = obj.Labels

	return result
}
