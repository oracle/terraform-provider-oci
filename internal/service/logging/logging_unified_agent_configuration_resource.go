// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package logging

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
	oci_logging "github.com/oracle/oci-go-sdk/v65/logging"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LoggingUnifiedAgentConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLoggingUnifiedAgentConfiguration,
		Read:     readLoggingUnifiedAgentConfiguration,
		Update:   updateLoggingUnifiedAgentConfiguration,
		Delete:   deleteLoggingUnifiedAgentConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"service_configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"configuration_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"LOGGING",
								"MONITORING",
							}, true),
						},

						// Optional
						"application_configurations": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"destination": {
										Type:     schema.TypeList,
										Required: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"compartment_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"metrics_namespace": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"source_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"KUBERNETES",
											"TAIL",
											"URL",
										}, true),
									},

									// Optional
									"source": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"scrape_targets": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"k8s_namespace": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"resource_group": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"resource_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"service_name": {
																Type:     schema.TypeString,
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

												// Computed
											},
										},
									},
									"sources": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"advanced_options": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"is_read_from_head": {
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
													Optional: true,
													Computed: true,
												},
												"parser": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"parser_type": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"APACHE2",
																	"APACHE_ERROR",
																	"AUDITD",
																	"CRI",
																	"CSV",
																	"GROK",
																	"JSON",
																	"MSGPACK",
																	"MULTILINE",
																	"MULTILINE_GROK",
																	"NONE",
																	"OPENMETRICS",
																	"REGEXP",
																	"SYSLOG",
																	"TSV",
																}, true),
															},

															// Optional
															"delimiter": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"expression": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"field_time_key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"format": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"format_firstline": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"grok_failure_key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"grok_name_key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"is_estimate_current_event": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"is_keep_time_key": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"is_merge_cri_fields": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"is_null_empty_string": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"is_support_colonless_ident": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"is_with_priority": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"keys": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"message_format": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"message_key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"multi_line_start_regexp": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"nested_parser": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"parse_nested": {
																			Type:     schema.TypeBool,
																			Optional: true,
																			Computed: true,
																		},
																		"separator": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"field_time_key": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"time_format": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"time_type": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"is_keep_time_key": {
																			Type:     schema.TypeBool,
																			Optional: true,
																			Computed: true,
																		}, // Computed
																	},
																},
															},
															"null_value_pattern": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"parse_nested": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"patterns": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"field_time_format": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"field_time_key": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"field_time_zone": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"name": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"pattern": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},

																		// Computed
																	},
																},
															},
															"record_input": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"dimensions": {
																			Type:     schema.TypeMap,
																			Optional: true,
																			Computed: true,
																			Elem:     schema.TypeString,
																		},
																		"namespace": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"resource_group": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},

																		// Computed
																	},
																},
															},
															"rfc5424time_format": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"separator": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"syslog_parser_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"time_format": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"time_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"timeout_in_milliseconds": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"types": {
																Type:     schema.TypeMap,
																Optional: true,
																Computed: true,
																Elem:     schema.TypeString,
															},

															// Computed
														},
													},
												},
												"paths": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"unified_agent_configuration_filter": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"allow_list": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"deny_list": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"filter_type": {
													Type:     schema.TypeString,
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

									// Computed
								},
							},
						},
						"destination": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"log_object_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"operational_metrics_configuration": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"destination": {
													Type:     schema.TypeList,
													Required: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"compartment_id": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},
												"source": {
													Type:     schema.TypeList,
													Required: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"type": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional
															"metrics": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															// Required
															"record_input": {
																Type:     schema.TypeList,
																Required: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"namespace": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		// Optional
																		"resource_group": {
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

												// Optional

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},
						"sources": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"source_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"CUSTOM_PLUGIN",
											"LOG_TAIL",
											"WINDOWS_EVENT_LOG",
										}, true),
									},

									// Optional
									"advanced_options": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"is_read_from_head": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"channels": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"custom_plugin": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"parser": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"parser_type": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"APACHE2",
														"APACHE_ERROR",
														"AUDITD",
														"CRI",
														"CSV",
														"GROK",
														"JSON",
														"MSGPACK",
														"MULTILINE",
														"MULTILINE_GROK",
														"NONE",
														"OPENMETRICS",
														"REGEXP",
														"SYSLOG",
														"TSV",
													}, true),
												},

												// Optional
												"delimiter": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"expression": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"field_time_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"format": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"format_firstline": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"grok_failure_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"grok_name_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"is_estimate_current_event": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_keep_time_key": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_merge_cri_fields": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_null_empty_string": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_support_colonless_ident": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_with_priority": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"keys": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"message_format": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"message_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"multi_line_start_regexp": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"nested_parser": {

													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Optional
															"parse_nested": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"separator": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"field_time_key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"time_format": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"time_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"is_keep_time_key": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															}, // Computed
														},
													},
												},
												"null_value_pattern": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"parse_nested": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"patterns": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"field_time_format": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"field_time_key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"field_time_zone": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"pattern": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"record_input": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"dimensions": {
																Type:     schema.TypeMap,
																Optional: true,
																Computed: true,
																Elem:     schema.TypeString,
															},
															"namespace": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"resource_group": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"rfc5424time_format": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"separator": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"syslog_parser_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"time_format": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"time_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"timeout_in_milliseconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"types": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													Elem:     schema.TypeString,
												},

												// Computed
											},
										},
									},
									"paths": {
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
						"unified_agent_configuration_filter": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"filter_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"CUSTOM_FILTER",
											"GREP_FILTER",
											"PARSER_FILTER",
											"RECORD_TRANSFORMER_FILTER",
										}, true),
									},

									// Optional
									"allow_list": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"pattern": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"custom_filter_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"custom_sections": {
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
												"params": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													Elem:     schema.TypeString,
												},

												// Computed
											},
										},
									},
									"deny_list": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"pattern": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"emit_invalid_record_to_error": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"hash_value_field": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"inject_key_prefix": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"is_auto_typecast_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"is_renew_record_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"is_ruby_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"keep_keys": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"key_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"params": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"parser": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"parser_type": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"APACHE2",
														"APACHE_ERROR",
														"AUDITD",
														"CRI",
														"CSV",
														"GROK",
														"JSON",
														"MSGPACK",
														"MULTILINE",
														"MULTILINE_GROK",
														"NONE",
														"OPENMETRICS",
														"REGEXP",
														"SYSLOG",
														"TSV",
													}, true),
												},

												// Optional
												"delimiter": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"expression": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"field_time_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"format": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"format_firstline": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"grok_failure_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"grok_name_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"is_estimate_current_event": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_keep_time_key": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_merge_cri_fields": {
													Type:             schema.TypeBool,
													Optional:         true,
													Computed:         true,
													DiffSuppressFunc: criDiffSuppressfunc,
												},
												"is_null_empty_string": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_support_colonless_ident": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_with_priority": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"keys": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"message_format": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"message_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"multi_line_start_regexp": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"nested_parser": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															// Optional
															"parse_nested": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"separator": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"field_time_key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"time_format": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"time_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"is_keep_time_key": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
														},
													},
													DiffSuppressFunc: criDiffSuppressfunc,
												},
												"null_value_pattern": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"parse_nested": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"patterns": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"field_time_format": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"field_time_key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"field_time_zone": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"pattern": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"record_input": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"dimensions": {
																Type:     schema.TypeMap,
																Optional: true,
																Computed: true,
																Elem:     schema.TypeString,
															},
															"namespace": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"resource_group": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"rfc5424time_format": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"separator": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"syslog_parser_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"time_format": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"time_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"timeout_in_milliseconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"types": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													Elem:     schema.TypeString,
												},

												// Computed
											},
										},
									},
									"record_list": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"key": {
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
									"remove_key_name_field": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"remove_keys": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"renew_time_key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"replace_invalid_sequence": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"reserve_data": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"reserve_time": {
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

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"group_association": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"group_list": {
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
			"configuration_state": {
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
			"time_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoggingUnifiedAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingUnifiedAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readLoggingUnifiedAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingUnifiedAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.ReadResource(sync)
}

func updateLoggingUnifiedAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingUnifiedAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLoggingUnifiedAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingUnifiedAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LoggingUnifiedAgentConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_logging.LoggingManagementClient
	Res                    *oci_logging.UnifiedAgentConfiguration
	DisableNotFoundRetries bool
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_logging.LogLifecycleStateCreating),
	}
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_logging.LogLifecycleStateActive),
	}
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_logging.LogLifecycleStateDeleting),
	}
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) Create() error {
	request := oci_logging.CreateUnifiedAgentConfigurationRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if groupAssociation, ok := s.D.GetOkExists("group_association"); ok {
		if tmpList := groupAssociation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "group_association", 0)
			tmp, err := s.mapToGroupAssociationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.GroupAssociation = &tmp
		}
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if serviceConfiguration, ok := s.D.GetOkExists("service_configuration"); ok {
		if tmpList := serviceConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "service_configuration", 0)
			tmp, err := s.mapToUnifiedAgentServiceConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ServiceConfiguration = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.CreateUnifiedAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_logging.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_logging.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "unifiedagentconfiguration") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getUnifiedAgentConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) getUnifiedAgentConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_logging.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	unifiedAgentConfigurationId, err := unifiedAgentConfigurationWaitForWorkRequest(workId, "unifiedagentconfiguration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*unifiedAgentConfigurationId)

	return s.Get()
}

func unifiedAgentConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "logging", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_logging.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func unifiedAgentConfigurationWaitForWorkRequest(wId *string, entityType string, action oci_logging.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_logging.LoggingManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "logging")
	retryPolicy.ShouldRetryOperation = unifiedAgentConfigurationWorkRequestShouldRetryFunc(timeout)

	response := oci_logging.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_logging.OperationStatusInProgress),
			string(oci_logging.OperationStatusAccepted),
			string(oci_logging.OperationStatusCancelling),
		},
		Target: []string{
			string(oci_logging.OperationStatusSucceeded),
			string(oci_logging.OperationStatusFailed),
			string(oci_logging.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_logging.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_logging.OperationStatusFailed || response.Status == oci_logging.OperationStatusCanceled {
		return nil, getErrorFromLoggingUnifiedAgentConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromLoggingUnifiedAgentConfigurationWorkRequest(client *oci_logging.LoggingManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_logging.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_logging.ListWorkRequestErrorsRequest{
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

func (s *LoggingUnifiedAgentConfigurationResourceCrud) Get() error {
	request := oci_logging.GetUnifiedAgentConfigurationRequest{}

	tmp := s.D.Id()
	request.UnifiedAgentConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.GetUnifiedAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UnifiedAgentConfiguration
	return nil
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_logging.UpdateUnifiedAgentConfigurationRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if groupAssociation, ok := s.D.GetOkExists("group_association"); ok {
		if tmpList := groupAssociation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "group_association", 0)
			tmp, err := s.mapToGroupAssociationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.GroupAssociation = &tmp
		}
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if serviceConfiguration, ok := s.D.GetOkExists("service_configuration"); ok {
		if tmpList := serviceConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "service_configuration", 0)
			tmp, err := s.mapToUnifiedAgentServiceConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ServiceConfiguration = tmp
		}
	}

	tmp := s.D.Id()
	request.UnifiedAgentConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.UpdateUnifiedAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getUnifiedAgentConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) Delete() error {
	request := oci_logging.DeleteUnifiedAgentConfigurationRequest{}

	tmp := s.D.Id()
	request.UnifiedAgentConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.DeleteUnifiedAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	// Wait until it finishes
	_, delWorkRequestErr := unifiedAgentConfigurationWaitForWorkRequest(workId, "unifiedagentconfiguration",
		oci_logging.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("configuration_state", s.Res.ConfigurationState)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GroupAssociation != nil {
		s.D.Set("group_association", []interface{}{GroupAssociationDetailsToMap(s.Res.GroupAssociation)})
	} else {
		s.D.Set("group_association", nil)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.ServiceConfiguration != nil {
		serviceConfigurationArray := []interface{}{}
		if serviceConfigurationMap := UnifiedAgentServiceConfigurationDetailsToMap(&s.Res.ServiceConfiguration); serviceConfigurationMap != nil {
			serviceConfigurationArray = append(serviceConfigurationArray, serviceConfigurationMap)
		}
		s.D.Set("service_configuration", serviceConfigurationArray)
	} else {
		s.D.Set("service_configuration", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToGrepFilterAllowRule(fieldKeyFormat string) (oci_logging.GrepFilterAllowRule, error) {
	result := oci_logging.GrepFilterAllowRule{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if pattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pattern")); ok {
		tmp := pattern.(string)
		result.Pattern = &tmp
	}

	return result, nil
}

func GrepFilterAllowRuleToMap(obj oci_logging.GrepFilterAllowRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Pattern != nil {
		result["pattern"] = string(*obj.Pattern)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToGrepFilterDenyRule(fieldKeyFormat string) (oci_logging.GrepFilterDenyRule, error) {
	result := oci_logging.GrepFilterDenyRule{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if pattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pattern")); ok {
		tmp := pattern.(string)
		result.Pattern = &tmp
	}

	return result, nil
}

func GrepFilterDenyRuleToMap(obj oci_logging.GrepFilterDenyRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Pattern != nil {
		result["pattern"] = string(*obj.Pattern)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToGrokPattern(fieldKeyFormat string) (oci_logging.GrokPattern, error) {
	result := oci_logging.GrokPattern{}

	if fieldTimeFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_format")); ok {
		tmp := fieldTimeFormat.(string)
		result.FieldTimeFormat = &tmp
	}

	if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
		tmp := fieldTimeKey.(string)
		result.FieldTimeKey = &tmp
	}

	if fieldTimeZone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_zone")); ok {
		tmp := fieldTimeZone.(string)
		result.FieldTimeZone = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if pattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pattern")); ok {
		tmp := pattern.(string)
		result.Pattern = &tmp
	}

	return result, nil
}

func GrokPatternToMap(obj oci_logging.GrokPattern) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FieldTimeFormat != nil {
		result["field_time_format"] = string(*obj.FieldTimeFormat)
	}

	if obj.FieldTimeKey != nil {
		result["field_time_key"] = string(*obj.FieldTimeKey)
	}

	if obj.FieldTimeZone != nil {
		result["field_time_zone"] = string(*obj.FieldTimeZone)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Pattern != nil {
		result["pattern"] = string(*obj.Pattern)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToGroupAssociationDetails(fieldKeyFormat string) (oci_logging.GroupAssociationDetails, error) {
	result := oci_logging.GroupAssociationDetails{}

	if groupList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_list")); ok {
		interfaces := groupList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "group_list")) {
			result.GroupList = tmp
		}
	}

	return result, nil
}

func GroupAssociationDetailsToMap(obj *oci_logging.GroupAssociationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["group_list"] = obj.GroupList

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToOperationalMetricsConfiguration(fieldKeyFormat string) (oci_logging.OperationalMetricsConfiguration, error) {
	result := oci_logging.OperationalMetricsConfiguration{}

	if destination, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination")); ok {
		if tmpList := destination.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destination"), 0)
			tmp, err := s.mapToOperationalMetricsDestination(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert destination, encountered error: %v", err)
			}
			result.Destination = &tmp
		}
	}

	if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
			tmp, err := s.mapToOperationalMetricsSource(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source, encountered error: %v", err)
			}
			result.Source = &tmp
		}
	}

	return result, nil
}

func OperationalMetricsConfigurationToMap(obj *oci_logging.OperationalMetricsConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Destination != nil {
		result["destination"] = []interface{}{OperationalMetricsDestinationToMap(obj.Destination)}
	}

	if obj.Source != nil {
		result["source"] = []interface{}{OperationalMetricsSourceToMap(obj.Source)}
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToOperationalMetricsDestination(fieldKeyFormat string) (oci_logging.OperationalMetricsDestination, error) {
	result := oci_logging.OperationalMetricsDestination{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	return result, nil
}

func OperationalMetricsDestinationToMap(obj *oci_logging.OperationalMetricsDestination) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToOperationalMetricsRecordInput(fieldKeyFormat string) (oci_logging.OperationalMetricsRecordInput, error) {
	result := oci_logging.OperationalMetricsRecordInput{}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if resourceGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_group")); ok {
		tmp := resourceGroup.(string)
		result.ResourceGroup = &tmp
	}

	return result, nil
}

func OperationalMetricsRecordInputToMap(obj *oci_logging.OperationalMetricsRecordInput) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.ResourceGroup != nil {
		result["resource_group"] = string(*obj.ResourceGroup)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToOperationalMetricsSource(fieldKeyFormat string) (oci_logging.OperationalMetricsSource, error) {
	result := oci_logging.OperationalMetricsSource{}

	if metrics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metrics")); ok {
		interfaces := metrics.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "metrics")) {
			result.Metrics = tmp
		}
	}

	if recordInput, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "record_input")); ok {
		if tmpList := recordInput.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "record_input"), 0)
			tmp, err := s.mapToOperationalMetricsRecordInput(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert record_input, encountered error: %v", err)
			}
			result.RecordInput = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_logging.OperationalMetricsSourceTypeEnum(type_.(string))
	}

	return result, nil
}

func OperationalMetricsSourceToMap(obj *oci_logging.OperationalMetricsSource) map[string]interface{} {
	result := map[string]interface{}{}

	result["metrics"] = obj.Metrics

	if obj.RecordInput != nil {
		result["record_input"] = []interface{}{OperationalMetricsRecordInputToMap(obj.RecordInput)}
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToRecordTransformerPair(fieldKeyFormat string) (oci_logging.RecordTransformerPair, error) {
	result := oci_logging.RecordTransformerPair{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func RecordTransformerPairToMap(obj oci_logging.RecordTransformerPair) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func UnifiedAgentConfigurationSummaryToMap(obj oci_logging.UnifiedAgentConfigurationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["configuration_state"] = string(obj.ConfigurationState)

	result["configuration_type"] = string(obj.ConfigurationType)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastModified != nil {
		result["time_last_modified"] = obj.TimeLastModified.String()
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentCustomSection(fieldKeyFormat string) (oci_logging.UnifiedAgentCustomSection, error) {
	result := oci_logging.UnifiedAgentCustomSection{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if params, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "params")); ok {
		result.Params = tfresource.ObjectMapToStringMap(params.(map[string]interface{}))
	}

	return result, nil
}

func UnifiedAgentCustomSectionToMap(obj oci_logging.UnifiedAgentCustomSection) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["params"] = obj.Params

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentKubernetesFilter(fieldKeyFormat string) (oci_logging.UnifiedAgentKubernetesFilter, error) {
	result := oci_logging.UnifiedAgentKubernetesFilter{}

	if allowList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allow_list")); ok {
		interfaces := allowList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allow_list")) {
			result.AllowList = tmp
		}
	}

	if denyList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deny_list")); ok {
		interfaces := denyList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "deny_list")) {
			result.DenyList = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func UnifiedAgentKubernetesFilterToMap(obj *oci_logging.UnifiedAgentKubernetesFilter) map[string]interface{} {
	result := map[string]interface{}{}

	result["allow_list"] = obj.AllowList

	result["deny_list"] = obj.DenyList

	result["filter_type"] = string(oci_logging.UnifiedAgentMonitoringFilterFilterTypeKubernetesFilter)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentKubernetesScrapeTarget(fieldKeyFormat string) (oci_logging.UnifiedAgentKubernetesScrapeTarget, error) {
	result := oci_logging.UnifiedAgentKubernetesScrapeTarget{}

	if k8sNamespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "k8s_namespace")); ok {
		tmp := k8sNamespace.(string)
		result.K8sNamespace = &tmp
	}

	if resourceGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_group")); ok {
		tmp := resourceGroup.(string)
		result.ResourceGroup = &tmp
	}

	if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
		result.ResourceType = oci_logging.UnifiedAgentKubernetesScrapeTargetResourceTypeEnum(resourceType.(string))
	}

	if serviceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_name")); ok {
		tmp := serviceName.(string)
		result.ServiceName = &tmp
	}

	return result, nil
}

func UnifiedAgentKubernetesScrapeTargetToMap(obj oci_logging.UnifiedAgentKubernetesScrapeTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.K8sNamespace != nil {
		result["k8s_namespace"] = string(*obj.K8sNamespace)
	}

	if obj.ResourceGroup != nil {
		result["resource_group"] = string(*obj.ResourceGroup)
	}

	result["resource_type"] = string(obj.ResourceType)

	if obj.ServiceName != nil {
		result["service_name"] = string(*obj.ServiceName)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentKubernetesSource(fieldKeyFormat string) (oci_logging.UnifiedAgentKubernetesSource, error) {
	result := oci_logging.UnifiedAgentKubernetesSource{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if scrapeTargets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scrape_targets")); ok {
		interfaces := scrapeTargets.([]interface{})
		tmp := make([]oci_logging.UnifiedAgentKubernetesScrapeTarget, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scrape_targets"), stateDataIndex)
			converted, err := s.mapToUnifiedAgentKubernetesScrapeTarget(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "scrape_targets")) {
			result.ScrapeTargets = tmp
		}
	}

	return result, nil
}

func UnifiedAgentKubernetesSourceToMap(obj *oci_logging.UnifiedAgentKubernetesSource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	scrapeTargets := []interface{}{}
	for _, item := range obj.ScrapeTargets {
		scrapeTargets = append(scrapeTargets, UnifiedAgentKubernetesScrapeTargetToMap(item))
	}
	result["scrape_targets"] = scrapeTargets

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentLoggingDestination(fieldKeyFormat string) (oci_logging.UnifiedAgentLoggingDestination, error) {
	result := oci_logging.UnifiedAgentLoggingDestination{}

	if logObjectId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_object_id")); ok {
		tmp := logObjectId.(string)
		result.LogObjectId = &tmp
	}

	if operationalMetricsConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operational_metrics_configuration")); ok {
		if tmpList := operationalMetricsConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "operational_metrics_configuration"), 0)
			tmp, err := s.mapToOperationalMetricsConfiguration(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert operational_metrics_configuration, encountered error: %v", err)
			}
			result.OperationalMetricsConfiguration = &tmp
		}
	}

	return result, nil
}

func UnifiedAgentLoggingDestinationToMap(obj *oci_logging.UnifiedAgentLoggingDestination) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogObjectId != nil {
		result["log_object_id"] = string(*obj.LogObjectId)
	}

	if obj.OperationalMetricsConfiguration != nil {
		result["operational_metrics_configuration"] = []interface{}{OperationalMetricsConfigurationToMap(obj.OperationalMetricsConfiguration)}
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentLoggingFilter(fieldKeyFormat string) (oci_logging.UnifiedAgentLoggingFilter, error) {
	var baseObject oci_logging.UnifiedAgentLoggingFilter
	//discriminator
	filterTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filter_type"))
	var filterType string
	if ok {
		filterType = filterTypeRaw.(string)
	} else {
		filterType = "" // default value
	}
	switch strings.ToLower(filterType) {
	case strings.ToLower("CUSTOM_FILTER"):
		details := oci_logging.UnifiedAgentCustomFilter{}
		if customFilterType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_filter_type")); ok {
			tmp := customFilterType.(string)
			details.CustomFilterType = &tmp
		}
		if customSections, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_sections")); ok {
			interfaces := customSections.([]interface{})
			tmp := make([]oci_logging.UnifiedAgentCustomSection, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "custom_sections"), stateDataIndex)
				converted, err := s.mapToUnifiedAgentCustomSection(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "custom_sections")) {
				details.CustomSections = tmp
			}
		}
		if params, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "params")); ok {
			details.Params = tfresource.ObjectMapToStringMap(params.(map[string]interface{}))
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	case strings.ToLower("GREP_FILTER"):
		details := oci_logging.UnifiedAgentLoggingGrepFilter{}
		if allowList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allow_list")); ok {
			interfaces := allowList.([]interface{})
			tmp := make([]oci_logging.GrepFilterAllowRule, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "allow_list"), stateDataIndex)
				converted, err := s.mapToGrepFilterAllowRule(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allow_list")) {
				details.AllowList = tmp
			}
		}
		if denyList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deny_list")); ok {
			interfaces := denyList.([]interface{})
			tmp := make([]oci_logging.GrepFilterDenyRule, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "deny_list"), stateDataIndex)
				converted, err := s.mapToGrepFilterDenyRule(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "deny_list")) {
				details.DenyList = tmp
			}
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	case strings.ToLower("PARSER_FILTER"):
		details := oci_logging.UnifiedAgentParserFilter{}
		if emitInvalidRecordToError, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "emit_invalid_record_to_error")); ok {
			tmp := emitInvalidRecordToError.(bool)
			details.EmitInvalidRecordToError = &tmp
		}
		if hashValueField, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hash_value_field")); ok {
			tmp := hashValueField.(string)
			details.HashValueField = &tmp
		}
		if injectKeyPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "inject_key_prefix")); ok {
			tmp := injectKeyPrefix.(string)
			details.InjectKeyPrefix = &tmp
		}
		if keyName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_name")); ok {
			tmp := keyName.(string)
			details.KeyName = &tmp
		}
		if parser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parser")); ok {
			if tmpList := parser.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parser"), 0)
				tmp, err := s.mapToUnifiedAgentParser(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert parser, encountered error: %v", err)
				}
				details.Parser = tmp
			}
		}
		if removeKeyNameField, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remove_key_name_field")); ok {
			tmp := removeKeyNameField.(bool)
			details.RemoveKeyNameField = &tmp
		}
		if replaceInvalidSequence, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replace_invalid_sequence")); ok {
			tmp := replaceInvalidSequence.(bool)
			details.ReplaceInvalidSequence = &tmp
		}
		if reserveData, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reserve_data")); ok {
			tmp := reserveData.(bool)
			details.ReserveData = &tmp
		}
		if reserveTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reserve_time")); ok {
			tmp := reserveTime.(bool)
			details.ReserveTime = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	case strings.ToLower("RECORD_TRANSFORMER_FILTER"):
		details := oci_logging.UnifiedAgentLoggingRecordTransformerFilter{}
		if isAutoTypecastEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_typecast_enabled")); ok {
			tmp := isAutoTypecastEnabled.(bool)
			details.IsAutoTypecastEnabled = &tmp
		}
		if isRenewRecordEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_renew_record_enabled")); ok {
			tmp := isRenewRecordEnabled.(bool)
			details.IsRenewRecordEnabled = &tmp
		}
		if isRubyEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ruby_enabled")); ok {
			tmp := isRubyEnabled.(bool)
			details.IsRubyEnabled = &tmp
		}
		if keepKeys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "keep_keys")); ok {
			interfaces := keepKeys.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "keep_keys")) {
				details.KeepKeys = tmp
			}
		}
		if recordList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "record_list")); ok {
			interfaces := recordList.([]interface{})
			tmp := make([]oci_logging.RecordTransformerPair, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "record_list"), stateDataIndex)
				converted, err := s.mapToRecordTransformerPair(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "record_list")) {
				details.RecordList = tmp
			}
		}
		if removeKeys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remove_keys")); ok {
			interfaces := removeKeys.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "remove_keys")) {
				details.RemoveKeys = tmp
			}
		}
		if renewTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "renew_time_key")); ok {
			tmp := renewTimeKey.(string)
			details.RenewTimeKey = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown filter_type '%v' was specified", filterType)
	}
	return baseObject, nil
}

func UnifiedAgentLoggingFilterToMap(obj oci_logging.UnifiedAgentLoggingFilter) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_logging.UnifiedAgentCustomFilter:
		result["filter_type"] = "CUSTOM_FILTER"
		result["name"] = v.Name

		if v.CustomFilterType != nil {
			result["custom_filter_type"] = string(*v.CustomFilterType)
		}

		customSections := []interface{}{}
		for _, item := range v.CustomSections {
			customSections = append(customSections, UnifiedAgentCustomSectionToMap(item))
		}
		result["custom_sections"] = customSections

		result["params"] = v.Params
	case oci_logging.UnifiedAgentLoggingGrepFilter:
		result["filter_type"] = "GREP_FILTER"
		result["name"] = v.Name

		allowList := []interface{}{}
		for _, item := range v.AllowList {
			allowList = append(allowList, GrepFilterAllowRuleToMap(item))
		}
		result["allow_list"] = allowList

		denyList := []interface{}{}
		for _, item := range v.DenyList {
			denyList = append(denyList, GrepFilterDenyRuleToMap(item))
		}
		result["deny_list"] = denyList
	case oci_logging.UnifiedAgentParserFilter:
		result["filter_type"] = "PARSER_FILTER"
		result["name"] = v.Name

		if v.EmitInvalidRecordToError != nil {
			result["emit_invalid_record_to_error"] = bool(*v.EmitInvalidRecordToError)
		}

		if v.HashValueField != nil {
			result["hash_value_field"] = string(*v.HashValueField)
		}

		if v.InjectKeyPrefix != nil {
			result["inject_key_prefix"] = string(*v.InjectKeyPrefix)
		}

		if v.KeyName != nil {
			result["key_name"] = string(*v.KeyName)
		}

		if v.Parser != nil {
			parserArray := []interface{}{}
			if parserMap := UnifiedAgentParserToMap(&v.Parser); parserMap != nil {
				parserArray = append(parserArray, parserMap)
			}
			result["parser"] = parserArray
		}

		if v.RemoveKeyNameField != nil {
			result["remove_key_name_field"] = bool(*v.RemoveKeyNameField)
		}

		if v.ReplaceInvalidSequence != nil {
			result["replace_invalid_sequence"] = bool(*v.ReplaceInvalidSequence)
		}

		if v.ReserveData != nil {
			result["reserve_data"] = bool(*v.ReserveData)
		}

		if v.ReserveTime != nil {
			result["reserve_time"] = bool(*v.ReserveTime)
		}
	case oci_logging.UnifiedAgentLoggingRecordTransformerFilter:
		result["filter_type"] = "RECORD_TRANSFORMER_FILTER"
		result["name"] = v.Name

		if v.IsAutoTypecastEnabled != nil {
			result["is_auto_typecast_enabled"] = bool(*v.IsAutoTypecastEnabled)
		}

		if v.IsRenewRecordEnabled != nil {
			result["is_renew_record_enabled"] = bool(*v.IsRenewRecordEnabled)
		}

		if v.IsRubyEnabled != nil {
			result["is_ruby_enabled"] = bool(*v.IsRubyEnabled)
		}

		result["keep_keys"] = v.KeepKeys

		recordList := []interface{}{}
		for _, item := range v.RecordList {
			recordList = append(recordList, RecordTransformerPairToMap(item))
		}
		result["record_list"] = recordList

		result["remove_keys"] = v.RemoveKeys

		if v.RenewTimeKey != nil {
			result["renew_time_key"] = string(*v.RenewTimeKey)
		}
	default:
		log.Printf("[WARN] Received 'filter_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentLoggingSource(fieldKeyFormat string) (oci_logging.UnifiedAgentLoggingSource, error) {
	var baseObject oci_logging.UnifiedAgentLoggingSource
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("CUSTOM_PLUGIN"):
		details := oci_logging.UnifiedAgentCustomPluginLogSource{}
		if customPlugin, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_plugin")); ok {
			tmp := customPlugin.(string)
			details.CustomPlugin = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	case strings.ToLower("LOG_TAIL"):
		details := oci_logging.UnifiedAgentTailLogSource{}
		if advancedOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "advanced_options")); ok {
			if tmpList := advancedOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "advanced_options"), 0)
				tmp, err := s.mapToUnifiedAgentTailSourceAdvancedOptions(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert advanced_options, encountered error: %v", err)
				}
				details.AdvancedOptions = &tmp
			}
		}
		if parser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parser")); ok {
			if tmpList := parser.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parser"), 0)
				tmp, err := s.mapToUnifiedAgentParser(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert parser, encountered error: %v", err)
				}
				details.Parser = tmp
			}
		}
		if paths, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "paths")); ok {
			interfaces := paths.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "paths")) {
				details.Paths = tmp
			}
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	case strings.ToLower("WINDOWS_EVENT_LOG"):
		details := oci_logging.UnifiedAgentWindowsEventSource{}
		if channels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "channels")); ok {
			interfaces := channels.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "channels")) {
				details.Channels = tmp
			}
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if parser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parser")); ok {
			if tmpList := parser.([]interface{}); len(tmpList) > 0 {
				return nil, fmt.Errorf("parser is not supported for windows_event_log source type")
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func UnifiedAgentLoggingSourceToMap(obj oci_logging.UnifiedAgentLoggingSource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_logging.UnifiedAgentCustomPluginLogSource:
		result["source_type"] = "CUSTOM_PLUGIN"

		if v.CustomPlugin != nil {
			result["custom_plugin"] = string(*v.CustomPlugin)
		}
	case oci_logging.UnifiedAgentTailLogSource:
		result["source_type"] = "LOG_TAIL"

		if v.AdvancedOptions != nil {
			result["advanced_options"] = []interface{}{UnifiedAgentTailSourceAdvancedOptionsToMap(v.AdvancedOptions)}
		}

		if v.Parser != nil {
			parserArray := []interface{}{}
			if parserMap := UnifiedAgentParserToMap(&v.Parser); parserMap != nil {
				parserArray = append(parserArray, parserMap)
			}
			result["parser"] = parserArray
		}
		result["name"] = v.Name
		result["paths"] = v.Paths
	case oci_logging.UnifiedAgentWindowsEventSource:
		result["source_type"] = "WINDOWS_EVENT_LOG"
		result["name"] = v.Name
		result["channels"] = v.Channels
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentMonitoringApplicationConfigurationDetails(fieldKeyFormat string) (oci_logging.UnifiedAgentMonitoringApplicationConfigurationDetails, error) {
	var baseObject oci_logging.UnifiedAgentMonitoringApplicationConfigurationDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("KUBERNETES"):
		details := oci_logging.UnifiedAgentKubernetesConfigurationDetails{}
		if destination, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination")); ok {
			if tmpList := destination.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destination"), 0)
				tmp, err := s.mapToUnifiedAgentMonitoringDestination(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert destination, encountered error: %v", err)
				}
				details.Destination = &tmp
			}
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			if tmpList := source.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
				tmp, err := s.mapToUnifiedAgentKubernetesSource(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert source, encountered error: %v", err)
				}
				details.Source = &tmp
			}
		}
		if unifiedAgentConfigurationFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unified_agent_configuration_filter")); ok {
			if tmpList := unifiedAgentConfigurationFilter.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "unified_agent_configuration_filter"), 0)
				tmp, err := s.mapToUnifiedAgentKubernetesFilter(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert unified_agent_configuration_filter, encountered error: %v", err)
				}
				details.Filter = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("TAIL"):
		details := oci_logging.UnifiedAgentOpenmetricsTailConfigurationDetails{}
		if destination, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination")); ok {
			if tmpList := destination.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destination"), 0)
				tmp, err := s.mapToUnifiedAgentMonitoringDestination(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert destination, encountered error: %v", err)
				}
				details.Destination = &tmp
			}
		}
		if sources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sources")); ok {
			interfaces := sources.([]interface{})
			tmp := make([]oci_logging.UnifiedAgentTailLogSource, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "sources"), stateDataIndex)
				converted, err := s.mapToUnifiedAgentTailLogSource(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sources")) {
				details.Sources = tmp
			}
		}
		baseObject = details
	case strings.ToLower("URL"):
		details := oci_logging.UnifiedAgentUrlConfigurationDetails{}
		if destination, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination")); ok {
			if tmpList := destination.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destination"), 0)
				tmp, err := s.mapToUnifiedAgentMonitoringDestination(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert destination, encountered error: %v", err)
				}
				details.Destination = &tmp
			}
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			if tmpList := source.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
				tmp, err := s.mapToUnifiedAgentMonitoringUrlSource(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert source, encountered error: %v", err)
				}
				details.Source = &tmp
			}
		}
		if unifiedAgentConfigurationFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unified_agent_configuration_filter")); ok {
			if tmpList := unifiedAgentConfigurationFilter.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "unified_agent_configuration_filter"), 0)
				tmp, err := s.mapToUnifiedAgentUrlFilter(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert unified_agent_configuration_filter, encountered error: %v", err)
				}
				details.Filter = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func UnifiedAgentMonitoringApplicationConfigurationDetailsToMap(obj oci_logging.UnifiedAgentMonitoringApplicationConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_logging.UnifiedAgentKubernetesConfigurationDetails:
		result["source_type"] = "KUBERNETES"

		if v.Destination != nil {
			result["destination"] = []interface{}{UnifiedAgentMonitoringDestinationToMap(v.Destination)}
		}

		if v.Source != nil {
			result["source"] = []interface{}{UnifiedAgentKubernetesSourceToMap(v.Source)}
		}

		if v.Filter != nil {
			result["unified_agent_configuration_filter"] = []interface{}{UnifiedAgentKubernetesFilterToMap(v.Filter)}
		}
	case oci_logging.UnifiedAgentOpenmetricsTailConfigurationDetails:
		result["source_type"] = "TAIL"

		if v.Destination != nil {
			result["destination"] = []interface{}{UnifiedAgentMonitoringDestinationToMap(v.Destination)}
		}

		sources := []interface{}{}
		for _, item := range v.Sources {
			sources = append(sources, UnifiedAgentTailLogSourceToMap(item))
		}
		result["sources"] = sources
	case oci_logging.UnifiedAgentUrlConfigurationDetails:
		result["source_type"] = "URL"

		if v.Destination != nil {
			result["destination"] = []interface{}{UnifiedAgentMonitoringDestinationToMap(v.Destination)}
		}

		if v.Source != nil {
			result["source"] = []interface{}{UnifiedAgentMonitoringUrlSourceToMap(v.Source)}
		}

		if v.Filter != nil {
			result["unified_agent_configuration_filter"] = []interface{}{UnifiedAgentUrlFilterToMap(v.Filter)}
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentMonitoringDestination(fieldKeyFormat string) (oci_logging.UnifiedAgentMonitoringDestination, error) {
	result := oci_logging.UnifiedAgentMonitoringDestination{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if metricsNamespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metrics_namespace")); ok {
		tmp := metricsNamespace.(string)
		result.MetricsNamespace = &tmp
	}

	return result, nil
}

func UnifiedAgentMonitoringDestinationToMap(obj *oci_logging.UnifiedAgentMonitoringDestination) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.MetricsNamespace != nil {
		result["metrics_namespace"] = string(*obj.MetricsNamespace)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentMonitoringUrlSource(fieldKeyFormat string) (oci_logging.UnifiedAgentMonitoringUrlSource, error) {
	result := oci_logging.UnifiedAgentMonitoringUrlSource{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if scrapeTargets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scrape_targets")); ok {
		interfaces := scrapeTargets.([]interface{})
		tmp := make([]oci_logging.UnifiedAgentUrlScrapeTarget, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scrape_targets"), stateDataIndex)
			converted, err := s.mapToUnifiedAgentUrlScrapeTarget(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "scrape_targets")) {
			result.ScrapeTargets = tmp
		}
	}

	return result, nil
}

func UnifiedAgentMonitoringUrlSourceToMap(obj *oci_logging.UnifiedAgentMonitoringUrlSource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	scrapeTargets := []interface{}{}
	for _, item := range obj.ScrapeTargets {
		scrapeTargets = append(scrapeTargets, UnifiedAgentUrlScrapeTargetToMap(item))
	}
	result["scrape_targets"] = scrapeTargets

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentOpenmetricsParserRecord(fieldKeyFormat string) (oci_logging.UnifiedAgentOpenmetricsParserRecord, error) {
	result := oci_logging.UnifiedAgentOpenmetricsParserRecord{}

	if dimensions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dimensions")); ok {
		result.Dimensions = tfresource.ObjectMapToStringMap(dimensions.(map[string]interface{}))
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if resourceGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_group")); ok {
		tmp := resourceGroup.(string)
		result.ResourceGroup = &tmp
	}

	return result, nil
}

func UnifiedAgentOpenmetricsParserRecordToMap(obj *oci_logging.UnifiedAgentOpenmetricsParserRecord) map[string]interface{} {
	result := map[string]interface{}{}
	if obj.Dimensions != nil {
		result["dimensions"] = obj.Dimensions
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.ResourceGroup != nil {
		result["resource_group"] = string(*obj.ResourceGroup)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentParser(fieldKeyFormat string) (oci_logging.UnifiedAgentParser, error) {
	var baseObject oci_logging.UnifiedAgentParser
	//discriminator
	parserTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parser_type"))
	var parserType string
	if ok {
		parserType = parserTypeRaw.(string)
	} else {
		parserType = "" // default value
	}

	switch strings.ToLower(parserType) {
	case strings.ToLower("APACHE2"):
		details := oci_logging.UnifiedAgentApache2Parser{}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("APACHE_ERROR"):
		details := oci_logging.UnifiedAgentApacheErrorParser{}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("AUDITD"):
		details := oci_logging.UnifiedAgentAuditdParser{}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("CRI"):
		details := oci_logging.UnifiedAgentCriParser{}
		if isMergeCriFields, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_merge_cri_fields")); ok {
			tmp := isMergeCriFields.(bool)
			details.IsMergeCriFields = &tmp
		}
		if nestedParser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nested_parser")); ok {
			if tmpList := nestedParser.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "nested_parser"), 0)
				tmp, err := s.mapToUnifiedJsonParser(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert nested_parser, encountered error: %v", err)
				}
				details.NestedParser = &tmp
			}
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("CSV"):
		details := oci_logging.UnifiedAgentCsvParser{}
		if delimiter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "delimiter")); ok {
			tmp := delimiter.(string)
			details.Delimiter = &tmp
		}
		if keys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "keys")); ok {
			interfaces := keys.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "keys")) {
				details.Keys = tmp
			}
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("GROK"):
		details := oci_logging.UnifiedAgentGrokParser{}
		if grokFailureKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grok_failure_key")); ok {
			tmp := grokFailureKey.(string)
			details.GrokFailureKey = &tmp
		}
		if grokNameKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grok_name_key")); ok {
			tmp := grokNameKey.(string)
			details.GrokNameKey = &tmp
		}
		if patterns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patterns")); ok {
			interfaces := patterns.([]interface{})
			tmp := make([]oci_logging.GrokPattern, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "patterns"), stateDataIndex)
				converted, err := s.mapToGrokPattern(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "patterns")) {
				details.Patterns = tmp
			}
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("JSON"):
		details := oci_logging.UnifiedJsonParser{}
		if parseNested, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parse_nested")); ok {
			tmp := parseNested.(bool)
			details.ParseNested = &tmp
		}
		if separator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "separator")); ok {
			tmp := separator.(string)
			details.Separator = &tmp
		}
		if timeFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_format")); ok {
			tmp := timeFormat.(string)
			details.TimeFormat = &tmp
		}
		if timeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_type")); ok {
			details.TimeType = oci_logging.UnifiedJsonParserTimeTypeEnum(timeType.(string))
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("MSGPACK"):
		details := oci_logging.UnifiedAgentMsgpackParser{}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("MULTILINE"):
		details := oci_logging.UnifiedAgentMultilineParser{}
		if format, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format")); ok {
			interfaces := format.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "format")) {
				details.Format = tmp
			}
		}
		if formatFirstline, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format_firstline")); ok {
			tmp := formatFirstline.(string)
			details.FormatFirstline = &tmp
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("MULTILINE_GROK"):
		details := oci_logging.UnifiedAgentMultilineGrokParser{}
		if grokFailureKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grok_failure_key")); ok {
			tmp := grokFailureKey.(string)
			details.GrokFailureKey = &tmp
		}
		if grokNameKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grok_name_key")); ok {
			tmp := grokNameKey.(string)
			details.GrokNameKey = &tmp
		}
		if multiLineStartRegexp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "multi_line_start_regexp")); ok {
			tmp := multiLineStartRegexp.(string)
			details.MultiLineStartRegexp = &tmp
		}
		if patterns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patterns")); ok {
			interfaces := patterns.([]interface{})
			tmp := make([]oci_logging.GrokPattern, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "patterns"), stateDataIndex)
				converted, err := s.mapToGrokPattern(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "patterns")) {
				details.Patterns = tmp
			}
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_logging.UnifiedAgentNoneParser{}
		if messageKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message_key")); ok {
			tmp := messageKey.(string)
			details.MessageKey = &tmp
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("OPENMETRICS"):
		details := oci_logging.UnifiedAgentOpenmetricsParser{}
		if recordInput, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "record_input")); ok {
			if tmpList := recordInput.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "record_input"), 0)
				tmp, err := s.mapToUnifiedAgentOpenmetricsParserRecord(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert record_input, encountered error: %v", err)
				}
				details.RecordInput = &tmp
			}
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("REGEXP"):
		details := oci_logging.UnifiedAgentRegexParser{}
		if expression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "expression")); ok {
			tmp := expression.(string)
			details.Expression = &tmp
		}
		if timeFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_format")); ok {
			tmp := timeFormat.(string)
			details.TimeFormat = &tmp
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("SYSLOG"):
		details := oci_logging.UnifiedAgentSyslogParser{}
		if isSupportColonlessIdent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_support_colonless_ident")); ok {
			tmp := isSupportColonlessIdent.(bool)
			details.IsSupportColonlessIdent = &tmp
		}
		if isWithPriority, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_with_priority")); ok {
			tmp := isWithPriority.(bool)
			details.IsWithPriority = &tmp
		}
		if messageFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message_format")); ok {
			details.MessageFormat = oci_logging.UnifiedAgentSyslogParserMessageFormatEnum(messageFormat.(string))
		}
		if rfc5424TimeFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rfc5424time_format")); ok {
			tmp := rfc5424TimeFormat.(string)
			details.Rfc5424TimeFormat = &tmp
		}
		if syslogParserType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "syslog_parser_type")); ok {
			details.SyslogParserType = oci_logging.UnifiedAgentSyslogParserSyslogParserTypeEnum(syslogParserType.(string))
		}
		if timeFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_format")); ok {
			tmp := timeFormat.(string)
			details.TimeFormat = &tmp
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("TSV"):
		details := oci_logging.UnifiedAgentTsvParser{}
		if delimiter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "delimiter")); ok {
			tmp := delimiter.(string)
			details.Delimiter = &tmp
		}
		if keys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "keys")); ok {
			interfaces := keys.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "keys")) {
				details.Keys = tmp
			}
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = tfresource.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown parser_type '%v' was specified", parserType)
	}
	return baseObject, nil
}

func UnifiedAgentParserToMap(obj *oci_logging.UnifiedAgentParser) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_logging.UnifiedAgentApache2Parser:
		result["parser_type"] = "APACHE2"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}
	case oci_logging.UnifiedAgentApacheErrorParser:
		result["parser_type"] = "APACHE_ERROR"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}
	case oci_logging.UnifiedAgentAuditdParser:
		result["parser_type"] = "AUDITD"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}
	case oci_logging.UnifiedAgentCriParser:
		result["parser_type"] = "CRI"

		if v.IsMergeCriFields != nil {
			result["is_merge_cri_fields"] = bool(*v.IsMergeCriFields)
		}

		if v.NestedParser != nil {
			result["nested_parser"] = []interface{}{UnifiedJsonParserToMap(v.NestedParser)}
		}

	case oci_logging.UnifiedAgentCsvParser:
		result["parser_type"] = "CSV"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}
		if v.Delimiter != nil {
			result["delimiter"] = string(*v.Delimiter)
		}

		result["keys"] = v.Keys
	case oci_logging.UnifiedAgentGrokParser:
		result["parser_type"] = "GROK"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}

		if v.GrokFailureKey != nil {
			result["grok_failure_key"] = string(*v.GrokFailureKey)
		}

		if v.GrokNameKey != nil {
			result["grok_name_key"] = string(*v.GrokNameKey)
		}

		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}

		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}

		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}

		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}

		patterns := []interface{}{}
		for _, item := range v.Patterns {
			patterns = append(patterns, GrokPatternToMap(item))
		}
		result["patterns"] = patterns
	case oci_logging.UnifiedJsonParser:
		result["parser_type"] = "JSON"
		if v.ParseNested != nil {
			result["parse_nested"] = bool(*v.ParseNested)
		}

		if v.Separator != nil {
			result["separator"] = string(*v.Separator)
		}
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}
		if v.TimeFormat != nil {
			result["time_format"] = string(*v.TimeFormat)
		}

		result["time_type"] = string(v.TimeType)
	case oci_logging.UnifiedAgentMsgpackParser:
		result["parser_type"] = "MSGPACK"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}
	case oci_logging.UnifiedAgentMultilineParser:
		result["parser_type"] = "MULTILINE"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}
		result["format"] = v.Format

		if v.FormatFirstline != nil {
			result["format_firstline"] = string(*v.FormatFirstline)
		}
	case oci_logging.UnifiedAgentMultilineGrokParser:
		result["parser_type"] = "MULTILINE_GROK"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}
		if v.GrokFailureKey != nil {
			result["grok_failure_key"] = string(*v.GrokFailureKey)
		}

		if v.GrokNameKey != nil {
			result["grok_name_key"] = string(*v.GrokNameKey)
		}

		if v.MultiLineStartRegexp != nil {
			result["multi_line_start_regexp"] = string(*v.MultiLineStartRegexp)
		}

		patterns := []interface{}{}
		for _, item := range v.Patterns {
			patterns = append(patterns, GrokPatternToMap(item))
		}
		result["patterns"] = patterns
	case oci_logging.UnifiedAgentNoneParser:
		result["parser_type"] = "NONE"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}
		if v.MessageKey != nil {
			result["message_key"] = string(*v.MessageKey)
		}
	case oci_logging.UnifiedAgentOpenmetricsParser:
		result["parser_type"] = "OPENMETRICS"

		if v.RecordInput != nil {
			result["record_input"] = []interface{}{UnifiedAgentOpenmetricsParserRecordToMap(v.RecordInput)}
		}
	case oci_logging.UnifiedAgentRegexParser:
		result["parser_type"] = "REGEXP"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}

		if v.Expression != nil {
			result["expression"] = string(*v.Expression)
		}

		if v.TimeFormat != nil {
			result["time_format"] = string(*v.TimeFormat)
		}
	case oci_logging.UnifiedAgentSyslogParser:
		result["parser_type"] = "SYSLOG"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}

		if v.IsSupportColonlessIdent != nil {
			result["is_support_colonless_ident"] = bool(*v.IsSupportColonlessIdent)
		}

		if v.IsWithPriority != nil {
			result["is_with_priority"] = bool(*v.IsWithPriority)
		}

		result["message_format"] = string(v.MessageFormat)

		if v.Rfc5424TimeFormat != nil {
			result["rfc5424time_format"] = string(*v.Rfc5424TimeFormat)
		}

		result["syslog_parser_type"] = string(v.SyslogParserType)

		if v.TimeFormat != nil {
			result["time_format"] = string(*v.TimeFormat)
		}
	case oci_logging.UnifiedAgentTsvParser:
		result["parser_type"] = "TSV"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		} else {
			result["is_estimate_current_event"] = true
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = tfresource.StringMapToObjectMap(v.Types)
		}
		if v.Delimiter != nil {
			result["delimiter"] = string(*v.Delimiter)
		}

		result["keys"] = v.Keys
	default:
		log.Printf("[WARN] Received 'parser_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentServiceConfigurationDetails(fieldKeyFormat string) (oci_logging.UnifiedAgentServiceConfigurationDetails, error) {
	var baseObject oci_logging.UnifiedAgentServiceConfigurationDetails
	//discriminator
	configurationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration_type"))
	var configurationType string
	if ok {
		configurationType = configurationTypeRaw.(string)
	} else {
		configurationType = "" // default value
	}
	switch strings.ToLower(configurationType) {
	case strings.ToLower("LOGGING"):
		details := oci_logging.UnifiedAgentLoggingConfiguration{}
		if destination, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination")); ok {
			if tmpList := destination.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destination"), 0)
				tmp, err := s.mapToUnifiedAgentLoggingDestination(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert destination, encountered error: %v", err)
				}
				details.Destination = &tmp
			}
		}
		if sources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sources")); ok {
			interfaces := sources.([]interface{})
			tmp := make([]oci_logging.UnifiedAgentLoggingSource, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "sources"), stateDataIndex)
				converted, err := s.mapToUnifiedAgentLoggingSource(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sources")) {
				details.Sources = tmp
			}
		}
		if unifiedAgentConfigurationFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unified_agent_configuration_filter")); ok {
			interfaces := unifiedAgentConfigurationFilter.([]interface{})
			tmp := make([]oci_logging.UnifiedAgentLoggingFilter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "unified_agent_configuration_filter"), stateDataIndex)
				converted, err := s.mapToUnifiedAgentLoggingFilter(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "unified_agent_configuration_filter")) {
				details.Filter = tmp
			}
		}
		baseObject = details
	case strings.ToLower("MONITORING"):
		details := oci_logging.UnifiedAgentMonitoringConfigurationDetails{}
		if applicationConfigurations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "application_configurations")); ok {
			interfaces := applicationConfigurations.([]interface{})
			tmp := make([]oci_logging.UnifiedAgentMonitoringApplicationConfigurationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "application_configurations"), stateDataIndex)
				converted, err := s.mapToUnifiedAgentMonitoringApplicationConfigurationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "application_configurations")) {
				details.ApplicationConfigurations = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown configuration_type '%v' was specified", configurationType)
	}
	return baseObject, nil
}

func UnifiedAgentServiceConfigurationDetailsToMap(obj *oci_logging.UnifiedAgentServiceConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_logging.UnifiedAgentLoggingConfiguration:
		result["configuration_type"] = "LOGGING"

		if v.Destination != nil {
			result["destination"] = []interface{}{UnifiedAgentLoggingDestinationToMap(v.Destination)}
		}

		sources := []interface{}{}
		for _, item := range v.Sources {
			sources = append(sources, UnifiedAgentLoggingSourceToMap(item))
		}
		result["sources"] = sources

		unifiedAgentConfigurationFilter := []interface{}{}
		for _, item := range v.Filter {
			unifiedAgentConfigurationFilter = append(unifiedAgentConfigurationFilter, UnifiedAgentLoggingFilterToMap(item))
		}
		result["unified_agent_configuration_filter"] = unifiedAgentConfigurationFilter
	case oci_logging.UnifiedAgentMonitoringConfigurationDetails:
		result["configuration_type"] = "MONITORING"

		applicationConfigurations := []interface{}{}
		for _, item := range v.ApplicationConfigurations {
			applicationConfigurations = append(applicationConfigurations, UnifiedAgentMonitoringApplicationConfigurationDetailsToMap(item))
		}
		result["application_configurations"] = applicationConfigurations
	default:
		log.Printf("[WARN] Received 'configuration_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentTailLogSource(fieldKeyFormat string) (oci_logging.UnifiedAgentTailLogSource, error) {
	result := oci_logging.UnifiedAgentTailLogSource{}

	if advancedOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "advanced_options")); ok {
		if tmpList := advancedOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "advanced_options"), 0)
			tmp, err := s.mapToUnifiedAgentTailSourceAdvancedOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert advanced_options, encountered error: %v", err)
			}
			result.AdvancedOptions = &tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if parser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parser")); ok {
		if tmpList := parser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parser"), 0)
			tmp, err := s.mapToUnifiedAgentParser(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parser, encountered error: %v", err)
			}
			result.Parser = tmp
		}
	}

	if paths, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "paths")); ok {
		interfaces := paths.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "paths")) {
			result.Paths = tmp
		}
	}

	return result, nil
}

func UnifiedAgentTailLogSourceToMap(obj oci_logging.UnifiedAgentTailLogSource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdvancedOptions != nil {
		result["advanced_options"] = []interface{}{UnifiedAgentTailSourceAdvancedOptionsToMap(obj.AdvancedOptions)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Parser != nil {
		parserArray := []interface{}{}
		if parserMap := UnifiedAgentParserToMap(&obj.Parser); parserMap != nil {
			parserArray = append(parserArray, parserMap)
		}
		result["parser"] = parserArray
	}

	result["paths"] = obj.Paths

	result["source_type"] = string(oci_logging.UnifiedAgentLoggingSourceSourceTypeLogTail)

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentTailSourceAdvancedOptions(fieldKeyFormat string) (oci_logging.UnifiedAgentTailSourceAdvancedOptions, error) {
	result := oci_logging.UnifiedAgentTailSourceAdvancedOptions{}

	if isReadFromHead, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_read_from_head")); ok {
		tmp := isReadFromHead.(bool)
		result.IsReadFromHead = &tmp
	}

	return result, nil
}

func UnifiedAgentTailSourceAdvancedOptionsToMap(obj *oci_logging.UnifiedAgentTailSourceAdvancedOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsReadFromHead != nil {
		result["is_read_from_head"] = bool(*obj.IsReadFromHead)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentUrlFilter(fieldKeyFormat string) (oci_logging.UnifiedAgentUrlFilter, error) {
	result := oci_logging.UnifiedAgentUrlFilter{}

	if allowList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allow_list")); ok {
		interfaces := allowList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allow_list")) {
			result.AllowList = tmp
		}
	}

	if denyList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deny_list")); ok {
		interfaces := denyList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "deny_list")) {
			result.DenyList = tmp
		}
	}
	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func UnifiedAgentUrlFilterToMap(obj *oci_logging.UnifiedAgentUrlFilter) map[string]interface{} {
	result := map[string]interface{}{}

	result["allow_list"] = obj.AllowList

	result["deny_list"] = obj.DenyList

	result["filter_type"] = string(oci_logging.UnifiedAgentMonitoringFilterFilterTypeUrlFilter)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentUrlScrapeTarget(fieldKeyFormat string) (oci_logging.UnifiedAgentUrlScrapeTarget, error) {
	result := oci_logging.UnifiedAgentUrlScrapeTarget{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if url, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url")); ok {
		tmp := url.(string)
		result.Url = &tmp
	}

	return result, nil
}

func UnifiedAgentUrlScrapeTargetToMap(obj oci_logging.UnifiedAgentUrlScrapeTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedJsonParser(fieldKeyFormat string) (oci_logging.UnifiedJsonParser, error) {
	result := oci_logging.UnifiedJsonParser{}

	if parseNested, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parse_nested")); ok {
		tmp := parseNested.(bool)
		result.ParseNested = &tmp
	}

	if separator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "separator")); ok {
		tmp := separator.(string)
		result.Separator = &tmp
	}

	if timeFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_format")); ok {
		tmp := timeFormat.(string)
		result.TimeFormat = &tmp
	}

	if timeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_type")); ok {
		result.TimeType = oci_logging.UnifiedJsonParserTimeTypeEnum(timeType.(string))
	}

	if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
		tmp := fieldTimeKey.(string)
		result.FieldTimeKey = &tmp
	}

	if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
		tmp := isKeepTimeKey.(bool)
		result.IsKeepTimeKey = &tmp
	}

	return result, nil
}

func UnifiedJsonParserToMap(obj *oci_logging.UnifiedJsonParser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ParseNested != nil {
		result["parse_nested"] = bool(*obj.ParseNested)
	}

	if obj.Separator != nil {
		result["separator"] = string(*obj.Separator)
	}

	if obj.TimeFormat != nil {
		result["time_format"] = string(*obj.TimeFormat)
	}

	result["time_type"] = string(obj.TimeType)

	if obj.FieldTimeKey != nil {
		result["field_time_key"] = *obj.FieldTimeKey
	}

	if obj.IsKeepTimeKey != nil {
		result["is_keep_time_key"] = *obj.IsKeepTimeKey
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_logging.ChangeUnifiedAgentConfigurationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.UnifiedAgentConfigurationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.ChangeUnifiedAgentConfigurationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getUnifiedAgentConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesRelated, s.D.Timeout(schema.TimeoutUpdate))
}

func criDiffSuppressfunc(k string, old string, new string, d *schema.ResourceData) bool {
	// k = "service_configuration.0.sources.0.parser.0.xxx"
	var parserTypeStr string
	dotIndex := findNthDotIndex(k, 6)
	if dotIndex != -1 {
		parserTypePath := k[:dotIndex+1] + "parser_type"
		if parserType, ok := d.GetOkExists(parserTypePath); ok {
			parserTypeStr = parserType.(string)
			if strings.ToLower(parserTypeStr) == "cri" {
				return false
			}
		}
	}
	log.Printf("Diff suppress for parser_type: %v. k: %v, old: %v, new: %v", parserTypeStr, k, old, new)
	return true
}

func findNthDotIndex(s string, n int) int {
	count := 0
	for i, char := range s {
		if char == '.' {
			count++
			if count == n {
				return i
			}
		}
	}
	return -1 // Return -1 if the ith dot is not found
}
