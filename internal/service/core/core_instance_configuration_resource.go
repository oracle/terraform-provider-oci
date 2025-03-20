// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreInstanceConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreInstanceConfiguration,
		Read:     readCoreInstanceConfiguration,
		Update:   updateCoreInstanceConfiguration,
		Delete:   deleteCoreInstanceConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"instance_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"instance_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"compute",
								"instance_options",
							}, true),
						},

						// Optional
						"block_volumes": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"attach_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"type": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"iscsi",
														"paravirtualized",
													}, true),
												},

												// Optional
												"device": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_pv_encryption_in_transit_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_read_only": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_shareable": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"use_chap": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"create_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"autotune_policies": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"autotune_type": {
																Type:             schema.TypeString,
																Required:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"DETACHED_VOLUME",
																	"PERFORMANCE_BASED",
																}, true),
															},

															// Optional
															"max_vpus_per_gb": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																ForceNew:         true,
																ValidateFunc:     tfresource.ValidateInt64TypeString,
																DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
															},

															// Computed
														},
													},
												},
												"availability_domain": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
												},
												"backup_policy_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"block_volume_replicas": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MinItems: 1,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"availability_domain": {
																Type:             schema.TypeString,
																Required:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
															},

															// Optional
															"display_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"cluster_placement_group_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"compartment_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"defined_tags": {
													Type:             schema.TypeMap,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
													Elem:             schema.TypeString,
												},
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"freeform_tags": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"is_auto_tune_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"kms_key_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"size_in_gbs": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													ValidateFunc:     tfresource.ValidateInt64TypeString,
													DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
												},
												"source_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"type": {
																Type:             schema.TypeString,
																Required:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"volume",
																	"volumeBackup",
																}, true),
															},

															// Optional
															"id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"vpus_per_gb": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													ValidateFunc:     tfresource.ValidateInt64TypeString,
													DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
												},
												"xrc_kms_key_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"volume_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"launch_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"agent_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"are_all_plugins_disabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_management_disabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_monitoring_disabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"plugins_config": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"desired_state": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},
									"availability_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"is_live_migration_preferred": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"recovery_action": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"availability_domain": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
									},
									"capacity_reservation_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"cluster_placement_group_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"create_vnic_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"assign_ipv6ip": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"assign_private_dns_record": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
													ForceNew: true,
												},
												"assign_public_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"defined_tags": {
													Type:             schema.TypeMap,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
													Elem:             schema.TypeString,
												},
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"freeform_tags": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"hostname_label": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"ipv6address_ipv6subnet_cidr_pair_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"ipv6address": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"ipv6subnet_cidr": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"nsg_ids": {
													Type:     schema.TypeSet,
													Optional: true,
													ForceNew: true,
													Set:      tfresource.LiteralTypeHashCodeForSets,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"private_ip": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"security_attributes": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"skip_source_dest_check": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"subnet_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"dedicated_vm_host_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"defined_tags": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
										Elem:             schema.TypeString,
									},
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"extended_metadata": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
										Elem:             schema.TypeString,
									},
									"fault_domain": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem:     schema.TypeString,
									},
									"instance_options": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"are_legacy_imds_endpoints_disabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"ipxe_script": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_pv_encryption_in_transit_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"launch_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"launch_options": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"boot_volume_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"firmware": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_consistent_volume_naming_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_pv_encryption_in_transit_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"network_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"remote_data_volume_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"licensing_configs": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"type": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},

												// Optional
												"license_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"metadata": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem:     schema.TypeString,
									},
									"platform_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"type": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"AMD_MILAN_BM",
														"AMD_MILAN_BM_GPU",
														"AMD_ROME_BM",
														"AMD_ROME_BM_GPU",
														"AMD_VM",
														"GENERIC_BM",
														"INTEL_ICELAKE_BM",
														"INTEL_SKYLAKE_BM",
														"INTEL_VM",
													}, true),
												},

												// Optional
												"are_virtual_instructions_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"config_map": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"is_access_control_service_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_input_output_memory_management_unit_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_measured_boot_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_memory_encryption_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_secure_boot_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_symmetric_multi_threading_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_trusted_platform_module_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"numa_nodes_per_socket": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"percentage_of_cores_enabled": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"preemptible_instance_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"preemption_action": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"type": {
																Type:             schema.TypeString,
																Required:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"TERMINATE",
																}, true),
															},

															// Optional
															"preserve_boot_volume": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},
									"preferred_maintenance_action": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"security_attributes": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem:     schema.TypeString,
									},
									"shape": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"shape_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"baseline_ocpu_utilization": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"memory_in_gbs": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"nvmes": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"ocpus": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vcpus": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"source_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"source_type": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"bootVolume",
														"image",
													}, true),
												},

												// Optional
												"boot_volume_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"boot_volume_size_in_gbs": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													ValidateFunc:     tfresource.ValidateInt64TypeString,
													DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
												},
												"boot_volume_vpus_per_gb": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													ValidateFunc:     tfresource.ValidateInt64TypeString,
													DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
												},
												"image_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"instance_source_image_filter_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
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
															"defined_tags_filter": {
																Type:             schema.TypeMap,
																Optional:         true,
																Computed:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
																Elem:             schema.TypeString,
															},
															"operating_system": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"operating_system_version": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"kms_key_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},
						"options": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"block_volumes": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"attach_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"type": {
																Type:             schema.TypeString,
																Required:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"iscsi",
																	"paravirtualized",
																}, true),
															},

															// Optional
															"device": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"display_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_pv_encryption_in_transit_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_read_only": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_shareable": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"use_chap": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"create_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"autotune_policies": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"autotune_type": {
																			Type:             schema.TypeString,
																			Required:         true,
																			ForceNew:         true,
																			DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																			ValidateFunc: validation.StringInSlice([]string{
																				"DETACHED_VOLUME",
																				"PERFORMANCE_BASED",
																			}, true),
																		},

																		// Optional
																		"max_vpus_per_gb": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			Computed:         true,
																			ForceNew:         true,
																			ValidateFunc:     tfresource.ValidateInt64TypeString,
																			DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
																		},

																		// Computed
																	},
																},
															},
															"availability_domain": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
															},
															"backup_policy_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"block_volume_replicas": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
																MinItems: 1,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"availability_domain": {
																			Type:             schema.TypeString,
																			Required:         true,
																			ForceNew:         true,
																			DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																		},

																		// Optional
																		"display_name": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																			ForceNew: true,
																		},

																		// Computed
																	},
																},
															},
															"cluster_placement_group_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"defined_tags": {
																Type:             schema.TypeMap,
																Optional:         true,
																Computed:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
																Elem:             schema.TypeString,
															},
															"display_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"freeform_tags": {
																Type:     schema.TypeMap,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem:     schema.TypeString,
															},
															"is_auto_tune_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"kms_key_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"size_in_gbs": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																ForceNew:         true,
																ValidateFunc:     tfresource.ValidateInt64TypeString,
																DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
															},
															"source_details": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"type": {
																			Type:             schema.TypeString,
																			Required:         true,
																			ForceNew:         true,
																			DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																			ValidateFunc: validation.StringInSlice([]string{
																				"volume",
																				"volumeBackup",
																			}, true),
																		},

																		// Optional
																		"id": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																			ForceNew: true,
																		},

																		// Computed
																	},
																},
															},
															"vpus_per_gb": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																ForceNew:         true,
																ValidateFunc:     tfresource.ValidateInt64TypeString,
																DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
															},
															"xrc_kms_key_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"volume_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"launch_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"agent_config": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"are_all_plugins_disabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_management_disabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_monitoring_disabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"plugins_config": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"desired_state": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																			ForceNew: true,
																		},
																		"name": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																			ForceNew: true,
																		},

																		// Computed
																	},
																},
															},

															// Computed
														},
													},
												},
												"availability_config": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"recovery_action": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_live_migration_preferred": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"availability_domain": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
												},
												"capacity_reservation_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"cluster_placement_group_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"compartment_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"create_vnic_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"assign_ipv6ip": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"assign_private_dns_record": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"assign_public_ip": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"defined_tags": {
																Type:             schema.TypeMap,
																Optional:         true,
																Computed:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
																Elem:             schema.TypeString,
															},
															"display_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"freeform_tags": {
																Type:     schema.TypeMap,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem:     schema.TypeString,
															},
															"hostname_label": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"ipv6address_ipv6subnet_cidr_pair_details": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"ipv6address": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																			ForceNew: true,
																		},
																		"ipv6subnet_cidr": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																			ForceNew: true,
																		},

																		// Computed
																	},
																},
															},
															"nsg_ids": {
																Type:     schema.TypeSet,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Set:      tfresource.LiteralTypeHashCodeForSets,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"private_ip": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"security_attributes": {
																Type:     schema.TypeMap,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem:     schema.TypeString,
															},
															"skip_source_dest_check": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"subnet_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"dedicated_vm_host_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"defined_tags": {
													Type:             schema.TypeMap,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
													Elem:             schema.TypeString,
												},
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"extended_metadata": {
													Type:             schema.TypeMap,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
													Elem:             schema.TypeString,
												},
												"fault_domain": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"freeform_tags": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"instance_options": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"are_legacy_imds_endpoints_disabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"ipxe_script": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_pv_encryption_in_transit_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"launch_mode": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"launch_options": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"boot_volume_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"firmware": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_consistent_volume_naming_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_pv_encryption_in_transit_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"network_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"remote_data_volume_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"licensing_configs": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"type": {
																Type:     schema.TypeString,
																Required: true,
																ForceNew: true,
															},

															// Optional
															"license_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"metadata": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"platform_config": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"type": {
																Type:             schema.TypeString,
																Required:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"AMD_MILAN_BM",
																	"AMD_ROME_BM",
																	"AMD_ROME_BM_GPU",
																	"AMD_VM",
																	"INTEL_ICELAKE_BM",
																	"INTEL_SKYLAKE_BM",
																	"INTEL_VM",
																}, true),
															},

															// Optional
															"are_virtual_instructions_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_access_control_service_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_input_output_memory_management_unit_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_measured_boot_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_memory_encryption_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_secure_boot_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_symmetric_multi_threading_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_trusted_platform_module_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"numa_nodes_per_socket": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"percentage_of_cores_enabled": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"preemptible_instance_config": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"preemption_action": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"type": {
																			Type:             schema.TypeString,
																			Required:         true,
																			ForceNew:         true,
																			DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																			ValidateFunc: validation.StringInSlice([]string{
																				"TERMINATE",
																			}, true),
																		},

																		// Optional
																		"preserve_boot_volume": {
																			Type:     schema.TypeBool,
																			Optional: true,
																			Computed: true,
																			ForceNew: true,
																		},

																		// Computed
																	},
																},
															},

															// Computed
														},
													},
												},
												"preferred_maintenance_action": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"security_attributes": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"shape": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"shape_config": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"baseline_ocpu_utilization": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"memory_in_gbs": {
																Type:     schema.TypeFloat,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"nvmes": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"ocpus": {
																Type:     schema.TypeFloat,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"vcpus": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"source_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"source_type": {
																Type:             schema.TypeString,
																Required:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"bootVolume",
																	"image",
																}, true),
															},

															// Optional
															"boot_volume_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"boot_volume_size_in_gbs": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																ForceNew:         true,
																ValidateFunc:     tfresource.ValidateInt64TypeString,
																DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
															},
															"boot_volume_vpus_per_gb": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																ForceNew:         true,
																ValidateFunc:     tfresource.ValidateInt64TypeString,
																DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
															},
															"image_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"kms_key_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"instance_source_image_filter_details": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
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
																		"defined_tags_filter": {
																			Type:             schema.TypeMap,
																			Optional:         true,
																			Computed:         true,
																			ForceNew:         true,
																			DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
																			Elem:             schema.TypeString,
																		},
																		"operating_system": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																			ForceNew: true,
																		},
																		"operating_system_version": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																			ForceNew: true,
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
									"secondary_vnics": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"create_vnic_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"assign_ipv6ip": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"assign_private_dns_record": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"assign_public_ip": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"defined_tags": {
																Type:             schema.TypeMap,
																Optional:         true,
																Computed:         true,
																ForceNew:         true,
																DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
																Elem:             schema.TypeString,
															},
															"display_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"freeform_tags": {
																Type:     schema.TypeMap,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem:     schema.TypeString,
															},
															"hostname_label": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"ipv6address_ipv6subnet_cidr_pair_details": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"ipv6address": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																			ForceNew: true,
																		},
																		"ipv6subnet_cidr": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																			ForceNew: true,
																		},

																		// Computed
																	},
																},
															},
															"nsg_ids": {
																Type:     schema.TypeSet,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Set:      tfresource.LiteralTypeHashCodeForSets,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"private_ip": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"security_attributes": {
																Type:     schema.TypeMap,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem:     schema.TypeString,
															},
															"skip_source_dest_check": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"subnet_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"nic_index": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},
						"secondary_vnics": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"create_vnic_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"assign_ipv6ip": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"assign_private_dns_record": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
													ForceNew: true,
												},
												"assign_public_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"defined_tags": {
													Type:             schema.TypeMap,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
													Elem:             schema.TypeString,
												},
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"freeform_tags": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"hostname_label": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"ipv6address_ipv6subnet_cidr_pair_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"ipv6address": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"ipv6subnet_cidr": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"nsg_ids": {
													Type:     schema.TypeSet,
													Optional: true,
													ForceNew: true,
													Set:      tfresource.LiteralTypeHashCodeForSets,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"private_ip": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"security_attributes": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"skip_source_dest_check": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"subnet_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"nic_index": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"INSTANCE",
					"NONE",
				}, true),
			},

			// Computed
			"deferred_fields": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.ReadResource(sync)
}

func updateCoreInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreInstanceConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeManagementClient
	Res                    *oci_core.InstanceConfiguration
	DisableNotFoundRetries bool
}

func (s *CoreInstanceConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreInstanceConfigurationResourceCrud) Create() error {
	request := oci_core.CreateInstanceConfigurationRequest{}
	err := s.populateTopLevelPolymorphicCreateInstanceConfigurationRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateInstanceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConfiguration
	return nil
}

func (s *CoreInstanceConfigurationResourceCrud) Get() error {
	request := oci_core.GetInstanceConfigurationRequest{}

	tmp := s.D.Id()
	request.InstanceConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInstanceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConfiguration
	return nil
}

func (s *CoreInstanceConfigurationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateInstanceConfigurationRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.InstanceConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInstanceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConfiguration
	return nil
}

func (s *CoreInstanceConfigurationResourceCrud) Delete() error {
	request := oci_core.DeleteInstanceConfigurationRequest{}

	tmp := s.D.Id()
	request.InstanceConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteInstanceConfiguration(context.Background(), request)
	return err
}

func (s *CoreInstanceConfigurationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("deferred_fields", s.Res.DeferredFields)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceDetails != nil {
		instanceDetailsArray := []interface{}{}
		if instanceDetailsMap := InstanceConfigurationInstanceDetailsToMap(&s.Res.InstanceDetails, false); instanceDetailsMap != nil {
			instanceDetailsArray = append(instanceDetailsArray, instanceDetailsMap)
		}
		s.D.Set("instance_details", instanceDetailsArray)
	} else {
		s.D.Set("instance_details", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreInstanceConfigurationResourceCrud) mapToComputeInstanceDetails(fieldKeyFormat string) (oci_core.ComputeInstanceDetails, error) {
	result := oci_core.ComputeInstanceDetails{}

	if blockVolumes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volumes")); ok {
		interfaces := blockVolumes.([]interface{})
		tmp := make([]oci_core.InstanceConfigurationBlockVolumeDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "block_volumes"), stateDataIndex)
			converted, err := s.mapToInstanceConfigurationBlockVolumeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "block_volumes")) {
			result.BlockVolumes = tmp
		}
	}

	if launchDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "launch_details")); ok {
		if tmpList := launchDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "launch_details"), 0)
			tmp, err := s.mapToInstanceConfigurationLaunchInstanceDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert launch_details, encountered error: %v", err)
			}
			result.LaunchDetails = &tmp
		}
	}

	if secondaryVnics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secondary_vnics")); ok {
		interfaces := secondaryVnics.([]interface{})
		tmp := make([]oci_core.InstanceConfigurationAttachVnicDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "secondary_vnics"), stateDataIndex)
			converted, err := s.mapToInstanceConfigurationAttachVnicDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "secondary_vnics")) {
			result.SecondaryVnics = tmp
		}
	}

	return result, nil
}

func ComputeInstanceDetailsToMap(obj oci_core.ComputeInstanceDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	blockVolumes := []interface{}{}
	for _, item := range obj.BlockVolumes {
		blockVolumes = append(blockVolumes, InstanceConfigurationBlockVolumeDetailsToMap(item))
	}
	result["block_volumes"] = blockVolumes

	if obj.LaunchDetails != nil {
		result["launch_details"] = []interface{}{InstanceConfigurationLaunchInstanceDetailsToMap(obj.LaunchDetails, datasource)}
	}

	secondaryVnics := []interface{}{}
	for _, item := range obj.SecondaryVnics {
		secondaryVnics = append(secondaryVnics, InstanceConfigurationAttachVnicDetailsToMap(item, datasource))
	}
	result["secondary_vnics"] = secondaryVnics

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceAgentPluginConfigDetails(fieldKeyFormat string) (oci_core.InstanceAgentPluginConfigDetails, error) {
	result := oci_core.InstanceAgentPluginConfigDetails{}

	if desiredState, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "desired_state")); ok {
		result.DesiredState = oci_core.InstanceAgentPluginConfigDetailsDesiredStateEnum(desiredState.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func InstanceAgentPluginConfigDetailsToMap(obj oci_core.InstanceAgentPluginConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["desired_state"] = string(obj.DesiredState)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationAttachVnicDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationAttachVnicDetails, error) {
	result := oci_core.InstanceConfigurationAttachVnicDetails{}

	if createVnicDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "create_vnic_details")); ok {
		if tmpList := createVnicDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "create_vnic_details"), 0)
			tmp, err := s.mapToInstanceConfigurationCreateVnicDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert create_vnic_details, encountered error: %v", err)
			}
			result.CreateVnicDetails = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if nicIndex, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nic_index")); ok {
		tmp := nicIndex.(int)
		result.NicIndex = &tmp
	}

	return result, nil
}

func InstanceConfigurationAttachVnicDetailsToMap(obj oci_core.InstanceConfigurationAttachVnicDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CreateVnicDetails != nil {
		result["create_vnic_details"] = []interface{}{InstanceConfigurationCreateVnicDetailsToMap(obj.CreateVnicDetails, datasource)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.NicIndex != nil {
		result["nic_index"] = int(*obj.NicIndex)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationAttachVolumeDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationAttachVolumeDetails, error) {
	var baseObject oci_core.InstanceConfigurationAttachVolumeDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("iscsi"):
		details := oci_core.InstanceConfigurationIscsiAttachVolumeDetails{}
		if useChap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "use_chap")); ok {
			tmp := useChap.(bool)
			details.UseChap = &tmp
		}
		if device, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "device")); ok {
			tmp := device.(string)
			details.Device = &tmp
		}
		if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if isReadOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_read_only")); ok {
			tmp := isReadOnly.(bool)
			details.IsReadOnly = &tmp
		}
		if isShareable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_shareable")); ok {
			tmp := isShareable.(bool)
			details.IsShareable = &tmp
		}
		baseObject = details
	case strings.ToLower("paravirtualized"):
		details := oci_core.InstanceConfigurationParavirtualizedAttachVolumeDetails{}
		if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pv_encryption_in_transit_enabled")); ok {
			tmp := isPvEncryptionInTransitEnabled.(bool)
			details.IsPvEncryptionInTransitEnabled = &tmp
		}
		if device, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "device")); ok {
			tmp := device.(string)
			details.Device = &tmp
		}
		if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if isReadOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_read_only")); ok {
			tmp := isReadOnly.(bool)
			details.IsReadOnly = &tmp
		}
		if isShareable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_shareable")); ok {
			tmp := isShareable.(bool)
			details.IsShareable = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func InstanceConfigurationAttachVolumeDetailsToMap(obj *oci_core.InstanceConfigurationAttachVolumeDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.InstanceConfigurationIscsiAttachVolumeDetails:
		result["type"] = "iscsi"

		if v.DisplayName != nil {
			result["display_name"] = *v.DisplayName
		}

		if v.IsReadOnly != nil {
			result["is_read_only"] = bool(*v.IsReadOnly)
		}

		if v.UseChap != nil {
			result["use_chap"] = bool(*v.UseChap)
		}

		if v.Device != nil {
			result["device"] = *v.Device
		}

		if v.IsShareable != nil {
			result["is_shareable"] = bool(*v.IsShareable)
		}
	case oci_core.InstanceConfigurationParavirtualizedAttachVolumeDetails:
		result["type"] = "paravirtualized"

		if v.IsPvEncryptionInTransitEnabled != nil {
			result["is_pv_encryption_in_transit_enabled"] = bool(*v.IsPvEncryptionInTransitEnabled)
		}

		if v.DisplayName != nil {
			result["display_name"] = *v.DisplayName
		}

		if v.Device != nil {
			result["device"] = *v.Device
		}

		if v.IsReadOnly != nil {
			result["is_read_only"] = bool(*v.IsReadOnly)
		}

		if v.IsShareable != nil {
			result["is_shareable"] = bool(*v.IsShareable)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationAutotunePolicy(fieldKeyFormat string) (oci_core.InstanceConfigurationAutotunePolicy, error) {
	var baseObject oci_core.InstanceConfigurationAutotunePolicy
	//discriminator
	autotuneTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "autotune_type"))
	var autotuneType string
	if ok {
		autotuneType = autotuneTypeRaw.(string)
	} else {
		autotuneType = "" // default value
	}
	switch strings.ToLower(autotuneType) {
	case strings.ToLower("DETACHED_VOLUME"):
		details := oci_core.InstanceConfigurationDetachedVolumeAutotunePolicy{}
		baseObject = details
	case strings.ToLower("PERFORMANCE_BASED"):
		details := oci_core.InstanceConfigurationPerformanceBasedAutotunePolicy{}
		if maxVpusPerGB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_vpus_per_gb")); ok {
			tmp := maxVpusPerGB.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert maxVpusPerGB string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.MaxVpusPerGB = &tmpInt64
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown autotune_type '%v' was specified", autotuneType)
	}
	return baseObject, nil
}

func InstanceConfigurationAutotunePolicyToMap(obj oci_core.InstanceConfigurationAutotunePolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_core.InstanceConfigurationDetachedVolumeAutotunePolicy:
		result["autotune_type"] = "DETACHED_VOLUME"
	case oci_core.InstanceConfigurationPerformanceBasedAutotunePolicy:
		result["autotune_type"] = "PERFORMANCE_BASED"

		if v.MaxVpusPerGB != nil {
			result["max_vpus_per_gb"] = strconv.FormatInt(*v.MaxVpusPerGB, 10)
		}
	default:
		log.Printf("[WARN] Received 'autotune_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationAvailabilityConfig(fieldKeyFormat string) (oci_core.InstanceConfigurationAvailabilityConfig, error) {
	result := oci_core.InstanceConfigurationAvailabilityConfig{}

	if isLiveMigrationPreferred, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_live_migration_preferred")); ok {
		tmp := isLiveMigrationPreferred.(bool)
		result.IsLiveMigrationPreferred = &tmp
	}

	if recoveryAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_action")); ok {
		result.RecoveryAction = oci_core.InstanceConfigurationAvailabilityConfigRecoveryActionEnum(recoveryAction.(string))
	}

	return result, nil
}

func InstanceConfigurationAvailabilityConfigToMap(obj *oci_core.InstanceConfigurationAvailabilityConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsLiveMigrationPreferred != nil {
		result["is_live_migration_preferred"] = bool(*obj.IsLiveMigrationPreferred)
	}

	result["recovery_action"] = string(obj.RecoveryAction)

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationBlockVolumeDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationBlockVolumeDetails, error) {
	result := oci_core.InstanceConfigurationBlockVolumeDetails{}

	if attachDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attach_details")); ok {
		if tmpList := attachDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "attach_details"), 0)
			tmp, err := s.mapToInstanceConfigurationAttachVolumeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert attach_details, encountered error: %v", err)
			}
			result.AttachDetails = tmp
		}
	}

	if createDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "create_details")); ok {
		if tmpList := createDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "create_details"), 0)
			tmp, err := s.mapToInstanceConfigurationCreateVolumeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert create_details, encountered error: %v", err)
			}
			result.CreateDetails = &tmp
		}
	}

	if volumeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_id")); ok {
		tmp := volumeId.(string)
		result.VolumeId = &tmp
	}

	return result, nil
}

func InstanceConfigurationBlockVolumeDetailsToMap(obj oci_core.InstanceConfigurationBlockVolumeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AttachDetails != nil {
		attachDetailsArray := []interface{}{}
		if attachDetailsMap := InstanceConfigurationAttachVolumeDetailsToMap(&obj.AttachDetails); attachDetailsMap != nil {
			attachDetailsArray = append(attachDetailsArray, attachDetailsMap)
		}
		result["attach_details"] = attachDetailsArray
	}

	if obj.CreateDetails != nil {
		result["create_details"] = []interface{}{InstanceConfigurationCreateVolumeDetailsToMap(obj.CreateDetails)}
	}

	if obj.VolumeId != nil {
		result["volume_id"] = string(*obj.VolumeId)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationBlockVolumeReplicaDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationBlockVolumeReplicaDetails, error) {
	result := oci_core.InstanceConfigurationBlockVolumeReplicaDetails{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result, nil
}

func InstanceConfigurationBlockVolumeReplicaDetailsToMap(obj oci_core.InstanceConfigurationBlockVolumeReplicaDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationCreateVnicDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationCreateVnicDetails, error) {
	result := oci_core.InstanceConfigurationCreateVnicDetails{}

	if assignIpv6Ip, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_ipv6ip")); ok {
		tmp := assignIpv6Ip.(bool)
		result.AssignIpv6Ip = &tmp
	}

	if assignPrivateDnsRecord, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_private_dns_record")); ok {
		tmp := assignPrivateDnsRecord.(bool)
		result.AssignPrivateDnsRecord = &tmp
	}

	if assignPublicIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_public_ip")); ok {
		tmp := assignPublicIp.(bool)
		result.AssignPublicIp = &tmp
	} else {
		t := true
		result.AssignPublicIp = &t
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
	}

	if ipv6AddressIpv6SubnetCidrPairDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details")); ok {
		interfaces := ipv6AddressIpv6SubnetCidrPairDetails.([]interface{})
		tmp := make([]oci_core.InstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details"), stateDataIndex)
			converted, err := s.mapToInstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details")) {
			result.Ipv6AddressIpv6SubnetCidrPairDetails = tmp
		}
	}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
			result.NsgIds = tmp
		}
	}

	if privateIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_ip")); ok {
		tmp := privateIp.(string)
		result.PrivateIp = &tmp
	}

	if securityAttributes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "security_attributes")); ok {
		result.SecurityAttributes = securityAttributes.(map[string]map[string]interface{})
	}

	if skipSourceDestCheck, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "skip_source_dest_check")); ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func InstanceConfigurationCreateVnicDetailsToMap(obj *oci_core.InstanceConfigurationCreateVnicDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssignIpv6Ip != nil {
		result["assign_ipv6ip"] = bool(*obj.AssignIpv6Ip)
	}

	if obj.AssignPrivateDnsRecord != nil {
		result["assign_private_dns_record"] = bool(*obj.AssignPrivateDnsRecord)
	} else {
		result["assign_private_dns_record"] = true
	}

	if obj.AssignPublicIp != nil {
		result["assign_public_ip"] = bool(*obj.AssignPublicIp)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostnameLabel != nil {
		result["hostname_label"] = string(*obj.HostnameLabel)
	}

	ipv6AddressIpv6SubnetCidrPairDetails := []interface{}{}
	for _, item := range obj.Ipv6AddressIpv6SubnetCidrPairDetails {
		ipv6AddressIpv6SubnetCidrPairDetails = append(ipv6AddressIpv6SubnetCidrPairDetails, InstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetailsToMap(item))
	}
	result["ipv6address_ipv6subnet_cidr_pair_details"] = ipv6AddressIpv6SubnetCidrPairDetails

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

	result["security_attributes"] = obj.SecurityAttributes

	if obj.SkipSourceDestCheck != nil {
		result["skip_source_dest_check"] = bool(*obj.SkipSourceDestCheck)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationCreateVolumeDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationCreateVolumeDetails, error) {
	result := oci_core.InstanceConfigurationCreateVolumeDetails{}

	if autotunePolicies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "autotune_policies")); ok {
		interfaces := autotunePolicies.([]interface{})
		tmp := make([]oci_core.InstanceConfigurationAutotunePolicy, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "autotune_policies"), stateDataIndex)
			converted, err := s.mapToInstanceConfigurationAutotunePolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "autotune_policies")) {
			result.AutotunePolicies = tmp
		}
	}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if backupPolicyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_policy_id")); ok {
		tmp := backupPolicyId.(string)
		result.BackupPolicyId = &tmp
	}

	if blockVolumeReplicas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_replicas")); ok {
		interfaces := blockVolumeReplicas.([]interface{})
		tmp := make([]oci_core.InstanceConfigurationBlockVolumeReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "block_volume_replicas"), stateDataIndex)
			converted, err := s.mapToInstanceConfigurationBlockVolumeReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "block_volume_replicas")) {
			result.BlockVolumeReplicas = tmp
		}
	}

	if clusterPlacementGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cluster_placement_group_id")); ok {
		tmp := clusterPlacementGroupId.(string)
		result.ClusterPlacementGroupId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAutoTuneEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_tune_enabled")); ok {
		tmp := isAutoTuneEnabled.(bool)
		result.IsAutoTuneEnabled = &tmp
	}

	if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
		tmp := kmsKeyId.(string)
		result.KmsKeyId = &tmp
	}

	if sizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "size_in_gbs")); ok {
		tmp := sizeInGBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert sizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.SizeInGBs = &tmpInt64
	}

	if sourceDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_details")); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_details"), 0)
			tmp, err := s.mapToInstanceConfigurationVolumeSourceDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_details, encountered error: %v", err)
			}
			result.SourceDetails = &tmp
		}
	}

	if vpusPerGB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpus_per_gb")); ok {
		tmp := vpusPerGB.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert vpusPerGB string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.VpusPerGB = &tmpInt64
	}

	if xrcKmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "xrc_kms_key_id")); ok {
		tmp := xrcKmsKeyId.(string)
		result.XrcKmsKeyId = &tmp
	}

	return result, nil
}

func InstanceConfigurationCreateVolumeDetailsToMap(obj *oci_core.InstanceConfigurationCreateVolumeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	autotunePolicies := []interface{}{}
	for _, item := range obj.AutotunePolicies {
		autotunePolicies = append(autotunePolicies, InstanceConfigurationAutotunePolicyToMap(item))
	}
	result["autotune_policies"] = autotunePolicies

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.BackupPolicyId != nil {
		result["backup_policy_id"] = string(*obj.BackupPolicyId)
	}

	blockVolumeReplicas := []interface{}{}
	for _, item := range obj.BlockVolumeReplicas {
		blockVolumeReplicas = append(blockVolumeReplicas, InstanceConfigurationBlockVolumeReplicaDetailsToMap(item))
	}
	result["block_volume_replicas"] = blockVolumeReplicas

	if obj.ClusterPlacementGroupId != nil {
		result["cluster_placement_group_id"] = string(*obj.ClusterPlacementGroupId)
	}

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

	if obj.IsAutoTuneEnabled != nil {
		result["is_auto_tune_enabled"] = bool(*obj.IsAutoTuneEnabled)
	}

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}

	if obj.SizeInGBs != nil {
		result["size_in_gbs"] = strconv.FormatInt(*obj.SizeInGBs, 10)
	}

	if obj.SourceDetails != nil {
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := InstanceConfigurationVolumeSourceDetailsToMap(&obj.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		result["source_details"] = sourceDetailsArray
	}

	if obj.VpusPerGB != nil {
		result["vpus_per_gb"] = strconv.FormatInt(*obj.VpusPerGB, 10)
	}

	if obj.XrcKmsKeyId != nil {
		result["xrc_kms_key_id"] = string(*obj.XrcKmsKeyId)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationInstanceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationInstanceDetails, error) {
	var baseObject oci_core.InstanceConfigurationInstanceDetails
	//discriminator
	instanceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_type"))
	var instanceType string
	if ok {
		instanceType = instanceTypeRaw.(string)
	} else {
		instanceType = "compute" // default value
	}
	switch strings.ToLower(instanceType) {
	case strings.ToLower("compute"):
		details := oci_core.ComputeInstanceDetails{}
		if blockVolumes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volumes")); ok {
			interfaces := blockVolumes.([]interface{})
			tmp := make([]oci_core.InstanceConfigurationBlockVolumeDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "block_volumes"), stateDataIndex)
				converted, err := s.mapToInstanceConfigurationBlockVolumeDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "block_volumes")) {
				details.BlockVolumes = tmp
			}
		}
		if launchDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "launch_details")); ok {
			if tmpList := launchDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "launch_details"), 0)
				tmp, err := s.mapToInstanceConfigurationLaunchInstanceDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert launch_details, encountered error: %v", err)
				}
				details.LaunchDetails = &tmp
			}
		}
		if secondaryVnics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secondary_vnics")); ok {
			interfaces := secondaryVnics.([]interface{})
			tmp := make([]oci_core.InstanceConfigurationAttachVnicDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "secondary_vnics"), stateDataIndex)
				converted, err := s.mapToInstanceConfigurationAttachVnicDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "secondary_vnics")) {
				details.SecondaryVnics = tmp
			}
		}
		baseObject = details
	case strings.ToLower("instance_options"):
		details := oci_core.ComputeInstanceOptions{}
		if options, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "options")); ok {
			interfaces := options.([]interface{})
			tmp := make([]oci_core.ComputeInstanceDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "options"), stateDataIndex)
				converted, err := s.mapToComputeInstanceDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "options")) {
				details.Options = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown instance_type '%v' was specified", instanceType)
	}
	return baseObject, nil
}

func InstanceConfigurationInstanceDetailsToMap(obj *oci_core.InstanceConfigurationInstanceDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.ComputeInstanceDetails:
		result["instance_type"] = "compute"

		blockVolumes := []interface{}{}
		for _, item := range v.BlockVolumes {
			blockVolumes = append(blockVolumes, InstanceConfigurationBlockVolumeDetailsToMap(item))
		}
		result["block_volumes"] = blockVolumes

		if v.LaunchDetails != nil {
			result["launch_details"] = []interface{}{InstanceConfigurationLaunchInstanceDetailsToMap(v.LaunchDetails, datasource)}
		}

		secondaryVnics := []interface{}{}
		for _, item := range v.SecondaryVnics {
			secondaryVnics = append(secondaryVnics, InstanceConfigurationAttachVnicDetailsToMap(item, datasource))
		}
		result["secondary_vnics"] = secondaryVnics
	case oci_core.ComputeInstanceOptions:
		result["instance_type"] = "instance_options"

		options := []interface{}{}
		for _, item := range v.Options {
			options = append(options, ComputeInstanceDetailsToMap(item, datasource))
		}
		result["options"] = options
	default:
		log.Printf("[WARN] Received 'instance_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationInstanceOptions(fieldKeyFormat string) (oci_core.InstanceConfigurationInstanceOptions, error) {
	result := oci_core.InstanceConfigurationInstanceOptions{}

	if areLegacyImdsEndpointsDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_legacy_imds_endpoints_disabled")); ok {
		tmp := areLegacyImdsEndpointsDisabled.(bool)
		result.AreLegacyImdsEndpointsDisabled = &tmp
	}

	return result, nil
}

func InstanceConfigurationInstanceOptionsToMap(obj *oci_core.InstanceConfigurationInstanceOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AreLegacyImdsEndpointsDisabled != nil {
		result["are_legacy_imds_endpoints_disabled"] = bool(*obj.AreLegacyImdsEndpointsDisabled)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationInstanceSourceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationInstanceSourceDetails, error) {
	var baseObject oci_core.InstanceConfigurationInstanceSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	var defaultVpusPerGb int64 = 10
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("bootVolume"):
		details := oci_core.InstanceConfigurationInstanceSourceViaBootVolumeDetails{}
		if bootVolumeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_id")); ok {
			tmp := bootVolumeId.(string)
			details.BootVolumeId = &tmp
		}
		baseObject = details
	case strings.ToLower("image"):
		details := oci_core.InstanceConfigurationInstanceSourceViaImageDetails{}
		if bootVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_size_in_gbs")); ok {
			tmp := bootVolumeSizeInGBs.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert bootVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.BootVolumeSizeInGBs = &tmpInt64
		}

		bootVolumeVpusPerGB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_vpus_per_gb"))
		if ok {
			tmp := bootVolumeVpusPerGB.(string)
			if tmp != "" {
				tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
				if err != nil {
					return details, fmt.Errorf("unable to convert bootVolumeVpusPerGB string: %s to an int64 and encountered error: %v", tmp, err)
				}
				details.BootVolumeVpusPerGB = &tmpInt64
			} else {
				details.BootVolumeVpusPerGB = &defaultVpusPerGb
			}
		} else {
			details.BootVolumeVpusPerGB = &defaultVpusPerGb
		}
		if imageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_id")); ok {
			tmp := imageId.(string)
			details.ImageId = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if instanceSourceImageFilterDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_source_image_filter_details")); ok {
			if tmpList := instanceSourceImageFilterDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "instance_source_image_filter_details"), 0)
				tmp, err := s.mapToInstanceConfigurationInstanceSourceImageFilterDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert instance_source_image_filter_details, encountered error: %v", err)
				}
				details.InstanceSourceImageFilterDetails = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func InstanceConfigurationInstanceSourceDetailsToMap(obj *oci_core.InstanceConfigurationInstanceSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.InstanceConfigurationInstanceSourceViaBootVolumeDetails:
		result["source_type"] = "bootVolume"

		if v.BootVolumeId != nil {
			result["boot_volume_id"] = string(*v.BootVolumeId)
		}
	case oci_core.InstanceConfigurationInstanceSourceViaImageDetails:
		result["source_type"] = "image"

		if v.BootVolumeSizeInGBs != nil {
			result["boot_volume_size_in_gbs"] = strconv.FormatInt(*v.BootVolumeSizeInGBs, 10)
		}

		if v.BootVolumeVpusPerGB != nil {
			result["boot_volume_vpus_per_gb"] = strconv.FormatInt(*v.BootVolumeVpusPerGB, 10)
		}

		if v.ImageId != nil {
			result["image_id"] = string(*v.ImageId)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.InstanceSourceImageFilterDetails != nil {
			result["instance_source_image_filter_details"] = []interface{}{InstanceConfigurationInstanceSourceImageFilterDetailsToMap(v.InstanceSourceImageFilterDetails)}
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationInstanceSourceImageFilterDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationInstanceSourceImageFilterDetails, error) {
	result := oci_core.InstanceConfigurationInstanceSourceImageFilterDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if definedTagsFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags_filter")); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTagsFilter.(map[string]interface{}))
		if err != nil {
			return result, err
		}
		result.DefinedTagsFilter = convertedDefinedTags
	}

	if operatingSystem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operating_system")); ok {
		tmp := operatingSystem.(string)
		result.OperatingSystem = &tmp
	}

	if operatingSystemVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operating_system_version")); ok {
		tmp := operatingSystemVersion.(string)
		result.OperatingSystemVersion = &tmp
	}

	return result, nil
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetails, error) {
	result := oci_core.InstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetails{}

	if ipv6Address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipv6address")); ok {
		tmp := ipv6Address.(string)
		result.Ipv6Address = &tmp
	}

	if ipv6SubnetCidr, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipv6subnet_cidr")); ok {
		tmp := ipv6SubnetCidr.(string)
		result.Ipv6SubnetCidr = &tmp
	}

	return result, nil
}

func InstanceConfigurationInstanceSourceImageFilterDetailsToMap(obj *oci_core.InstanceConfigurationInstanceSourceImageFilterDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTagsFilter != nil {
		result["defined_tags_filter"] = tfresource.DefinedTagsToMap(obj.DefinedTagsFilter)
	}

	if obj.OperatingSystem != nil {
		result["operating_system"] = string(*obj.OperatingSystem)
	}

	if obj.OperatingSystemVersion != nil {
		result["operating_system_version"] = string(*obj.OperatingSystemVersion)
	}

	return result
}

func InstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetailsToMap(obj oci_core.InstanceConfigurationIpv6AddressIpv6SubnetCidrPairDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ipv6Address != nil {
		result["ipv6address"] = string(*obj.Ipv6Address)
	}

	if obj.Ipv6SubnetCidr != nil {
		result["ipv6subnet_cidr"] = string(*obj.Ipv6SubnetCidr)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationLaunchInstanceAgentConfigDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationLaunchInstanceAgentConfigDetails, error) {
	result := oci_core.InstanceConfigurationLaunchInstanceAgentConfigDetails{}

	if areAllPluginsDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_all_plugins_disabled")); ok {
		tmp := areAllPluginsDisabled.(bool)
		result.AreAllPluginsDisabled = &tmp
	}

	if isManagementDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_management_disabled")); ok {
		tmp := isManagementDisabled.(bool)
		result.IsManagementDisabled = &tmp
	}

	if isMonitoringDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_monitoring_disabled")); ok {
		tmp := isMonitoringDisabled.(bool)
		result.IsMonitoringDisabled = &tmp
	}

	if pluginsConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "plugins_config")); ok {
		interfaces := pluginsConfig.([]interface{})
		tmp := make([]oci_core.InstanceAgentPluginConfigDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "plugins_config"), stateDataIndex)
			converted, err := s.mapToInstanceAgentPluginConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "plugins_config")) {
			result.PluginsConfig = tmp
		}
	}

	return result, nil
}

func InstanceConfigurationLaunchInstanceAgentConfigDetailsToMap(obj *oci_core.InstanceConfigurationLaunchInstanceAgentConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AreAllPluginsDisabled != nil {
		result["are_all_plugins_disabled"] = bool(*obj.AreAllPluginsDisabled)
	}

	if obj.IsManagementDisabled != nil {
		result["is_management_disabled"] = bool(*obj.IsManagementDisabled)
	}

	if obj.IsMonitoringDisabled != nil {
		result["is_monitoring_disabled"] = bool(*obj.IsMonitoringDisabled)
	}

	pluginsConfig := []interface{}{}
	for _, item := range obj.PluginsConfig {
		pluginsConfig = append(pluginsConfig, InstanceAgentPluginConfigDetailsToMap(item))
	}
	result["plugins_config"] = pluginsConfig

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationLaunchInstanceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationLaunchInstanceDetails, error) {
	result := oci_core.InstanceConfigurationLaunchInstanceDetails{}

	if agentConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "agent_config")); ok {
		if tmpList := agentConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "agent_config"), 0)
			tmp, err := s.mapToInstanceConfigurationLaunchInstanceAgentConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert agent_config, encountered error: %v", err)
			}
			result.AgentConfig = &tmp
		}
	}

	if availabilityConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_config")); ok {
		if tmpList := availabilityConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "availability_config"), 0)
			tmp, err := s.mapToInstanceConfigurationAvailabilityConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert availability_config, encountered error: %v", err)
			}
			result.AvailabilityConfig = &tmp
		}
	}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if capacityReservationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "capacity_reservation_id")); ok {
		tmp := capacityReservationId.(string)
		result.CapacityReservationId = &tmp
	}

	if clusterPlacementGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cluster_placement_group_id")); ok {
		tmp := clusterPlacementGroupId.(string)
		result.ClusterPlacementGroupId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if createVnicDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "create_vnic_details")); ok {
		if tmpList := createVnicDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "create_vnic_details"), 0)
			tmp, err := s.mapToInstanceConfigurationCreateVnicDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert create_vnic_details, encountered error: %v", err)
			}
			result.CreateVnicDetails = &tmp
		}
	}

	if dedicatedVmHostId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dedicated_vm_host_id")); ok {
		tmp := dedicatedVmHostId.(string)
		result.DedicatedVmHostId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if extendedMetadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extended_metadata")); ok {
		extendedMetadata, err := mapToExtendedMetadata(extendedMetadata.(map[string]interface{}))
		if err != nil {
			return result, err
		}
		result.ExtendedMetadata = extendedMetadata
	}

	if faultDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_domain")); ok {
		tmp := faultDomain.(string)
		result.FaultDomain = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_options")); ok {
		if tmpList := instanceOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "instance_options"), 0)
			tmp, err := s.mapToInstanceConfigurationInstanceOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert instance_options, encountered error: %v", err)
			}
			result.InstanceOptions = &tmp
		}
	}

	if ipxeScript, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipxe_script")); ok {
		tmp := ipxeScript.(string)
		result.IpxeScript = &tmp
	}

	if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pv_encryption_in_transit_enabled")); ok {
		tmp := isPvEncryptionInTransitEnabled.(bool)
		result.IsPvEncryptionInTransitEnabled = &tmp
	}

	if launchMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "launch_mode")); ok {
		result.LaunchMode = oci_core.InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum(launchMode.(string))
	}

	if launchOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "launch_options")); ok {
		if tmpList := launchOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "launch_options"), 0)
			tmp, err := s.mapToInstanceConfigurationLaunchOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert launch_options, encountered error: %v", err)
			}
			result.LaunchOptions = &tmp
		}
	}

	if licensingConfigs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "licensing_configs")); ok {
		interfaces := licensingConfigs.([]interface{})
		tmp := make([]oci_core.LaunchInstanceLicensingConfig, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "licensing_configs"), stateDataIndex)
			converted, err := s.mapToLaunchInstanceLicensingConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "licensing_configs")) {
			result.LicensingConfigs = tmp
		}
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		result.Metadata = tfresource.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}

	if platformConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "platform_config")); ok {
		if tmpList := platformConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "platform_config"), 0)
			tmp, err := s.mapToInstanceConfigurationLaunchInstancePlatformConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert platform_config, encountered error: %v", err)
			}
			result.PlatformConfig = tmp
		}
	}

	if preemptibleInstanceConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preemptible_instance_config")); ok {
		if tmpList := preemptibleInstanceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "preemptible_instance_config"), 0)
			tmp, err := s.mapToPreemptibleInstanceConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert preemptible_instance_config, encountered error: %v", err)
			}
			result.PreemptibleInstanceConfig = &tmp
		}
	}

	if preferredMaintenanceAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preferred_maintenance_action")); ok {
		result.PreferredMaintenanceAction = oci_core.InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum(preferredMaintenanceAction.(string))
	}

	if securityAttributes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "security_attributes")); ok {
		result.SecurityAttributes = securityAttributes.(map[string]map[string]interface{})
	}

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
	}

	if shapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_config")); ok {
		if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "shape_config"), 0)
			tmp, err := s.mapToInstanceConfigurationLaunchInstanceShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert shape_config, encountered error: %v", err)
			}
			result.ShapeConfig = &tmp
		}
	}

	if sourceDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_details")); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_details"), 0)
			tmp, err := s.mapToInstanceConfigurationInstanceSourceDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_details, encountered error: %v", err)
			}
			result.SourceDetails = &tmp
		}
	}

	return result, nil
}

func InstanceConfigurationLaunchInstanceDetailsToMap(obj *oci_core.InstanceConfigurationLaunchInstanceDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentConfig != nil {
		result["agent_config"] = []interface{}{InstanceConfigurationLaunchInstanceAgentConfigDetailsToMap(obj.AgentConfig)}
	}

	if obj.AvailabilityConfig != nil {
		result["availability_config"] = []interface{}{InstanceConfigurationAvailabilityConfigToMap(obj.AvailabilityConfig)}
	}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CapacityReservationId != nil {
		result["capacity_reservation_id"] = string(*obj.CapacityReservationId)
	}

	if obj.ClusterPlacementGroupId != nil {
		result["cluster_placement_group_id"] = string(*obj.ClusterPlacementGroupId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CreateVnicDetails != nil {
		result["create_vnic_details"] = []interface{}{InstanceConfigurationCreateVnicDetailsToMap(obj.CreateVnicDetails, datasource)}
	}

	if obj.DedicatedVmHostId != nil {
		result["dedicated_vm_host_id"] = string(*obj.DedicatedVmHostId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["extended_metadata"] = tfresource.GenericMapToJsonMap(obj.ExtendedMetadata)

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.InstanceOptions != nil {
		result["instance_options"] = []interface{}{InstanceConfigurationInstanceOptionsToMap(obj.InstanceOptions)}
	}

	if obj.IpxeScript != nil {
		result["ipxe_script"] = string(*obj.IpxeScript)
	}

	if obj.IsPvEncryptionInTransitEnabled != nil {
		result["is_pv_encryption_in_transit_enabled"] = bool(*obj.IsPvEncryptionInTransitEnabled)
	}

	result["launch_mode"] = string(obj.LaunchMode)

	if obj.LaunchOptions != nil {
		result["launch_options"] = []interface{}{InstanceConfigurationLaunchOptionsToMap(obj.LaunchOptions)}
	}

	licensingConfigs := []interface{}{}
	for _, item := range obj.LicensingConfigs {
		licensingConfigs = append(licensingConfigs, LaunchInstanceLicensingConfigToMap(item))
	}
	result["licensing_configs"] = licensingConfigs

	result["metadata"] = obj.Metadata

	if obj.PlatformConfig != nil {
		platformConfigArray := []interface{}{}
		if platformConfigMap := InstanceConfigurationLaunchInstancePlatformConfigToMap(&obj.PlatformConfig); platformConfigMap != nil {
			platformConfigArray = append(platformConfigArray, platformConfigMap)
		}
		result["platform_config"] = platformConfigArray
	}

	if obj.PreemptibleInstanceConfig != nil {
		result["preemptible_instance_config"] = []interface{}{PreemptibleInstanceConfigDetailsToMap(obj.PreemptibleInstanceConfig)}
	}

	result["preferred_maintenance_action"] = string(obj.PreferredMaintenanceAction)

	result["security_attributes"] = obj.SecurityAttributes

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.ShapeConfig != nil {
		result["shape_config"] = []interface{}{InstanceConfigurationLaunchInstanceShapeConfigDetailsToMap(obj.ShapeConfig)}
	}

	if obj.SourceDetails != nil {
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := InstanceConfigurationInstanceSourceDetailsToMap(&obj.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		result["source_details"] = sourceDetailsArray
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationLaunchInstancePlatformConfig(fieldKeyFormat string) (oci_core.InstanceConfigurationLaunchInstancePlatformConfig, error) {
	var baseObject oci_core.InstanceConfigurationLaunchInstancePlatformConfig
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("AMD_MILAN_BM"):
		details := oci_core.InstanceConfigurationAmdMilanBmLaunchInstancePlatformConfig{}
		if areVirtualInstructionsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_virtual_instructions_enabled")); ok {
			tmp := areVirtualInstructionsEnabled.(bool)
			details.AreVirtualInstructionsEnabled = &tmp
		}
		if configMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_map")); ok {
			details.ConfigMap = tfresource.ObjectMapToStringMap(configMap.(map[string]interface{}))
		}
		if isAccessControlServiceEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_access_control_service_enabled")); ok {
			tmp := isAccessControlServiceEnabled.(bool)
			details.IsAccessControlServiceEnabled = &tmp
		}
		if isInputOutputMemoryManagementUnitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_input_output_memory_management_unit_enabled")); ok {
			tmp := isInputOutputMemoryManagementUnitEnabled.(bool)
			details.IsInputOutputMemoryManagementUnitEnabled = &tmp
		}
		if isSymmetricMultiThreadingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_symmetric_multi_threading_enabled")); ok {
			tmp := isSymmetricMultiThreadingEnabled.(bool)
			details.IsSymmetricMultiThreadingEnabled = &tmp
		}
		if numaNodesPerSocket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "numa_nodes_per_socket")); ok {
			details.NumaNodesPerSocket = oci_core.InstanceConfigurationAmdMilanBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum(numaNodesPerSocket.(string))
		}
		if percentageOfCoresEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "percentage_of_cores_enabled")); ok {
			tmp := percentageOfCoresEnabled.(int)
			details.PercentageOfCoresEnabled = &tmp
		}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isMemoryEncryptionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_memory_encryption_enabled")); ok {
			tmp := isMemoryEncryptionEnabled.(bool)
			details.IsMemoryEncryptionEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("AMD_MILAN_BM_GPU"):
		details := oci_core.InstanceConfigurationAmdMilanBmGpuLaunchInstancePlatformConfig{}
		if areVirtualInstructionsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_virtual_instructions_enabled")); ok {
			tmp := areVirtualInstructionsEnabled.(bool)
			details.AreVirtualInstructionsEnabled = &tmp
		}
		if configMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_map")); ok {
			details.ConfigMap = tfresource.ObjectMapToStringMap(configMap.(map[string]interface{}))
		}
		if isAccessControlServiceEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_access_control_service_enabled")); ok {
			tmp := isAccessControlServiceEnabled.(bool)
			details.IsAccessControlServiceEnabled = &tmp
		}
		if isInputOutputMemoryManagementUnitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_input_output_memory_management_unit_enabled")); ok {
			tmp := isInputOutputMemoryManagementUnitEnabled.(bool)
			details.IsInputOutputMemoryManagementUnitEnabled = &tmp
		}
		if isSymmetricMultiThreadingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_symmetric_multi_threading_enabled")); ok {
			tmp := isSymmetricMultiThreadingEnabled.(bool)
			details.IsSymmetricMultiThreadingEnabled = &tmp
		}
		if numaNodesPerSocket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "numa_nodes_per_socket")); ok {
			details.NumaNodesPerSocket = oci_core.InstanceConfigurationAmdMilanBmGpuLaunchInstancePlatformConfigNumaNodesPerSocketEnum(numaNodesPerSocket.(string))
		}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isMemoryEncryptionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_memory_encryption_enabled")); ok {
			tmp := isMemoryEncryptionEnabled.(bool)
			details.IsMemoryEncryptionEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("AMD_ROME_BM"):
		details := oci_core.InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig{}
		if areVirtualInstructionsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_virtual_instructions_enabled")); ok {
			tmp := areVirtualInstructionsEnabled.(bool)
			details.AreVirtualInstructionsEnabled = &tmp
		}
		if configMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_map")); ok {
			details.ConfigMap = tfresource.ObjectMapToStringMap(configMap.(map[string]interface{}))
		}
		if isAccessControlServiceEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_access_control_service_enabled")); ok {
			tmp := isAccessControlServiceEnabled.(bool)
			details.IsAccessControlServiceEnabled = &tmp
		}
		if isInputOutputMemoryManagementUnitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_input_output_memory_management_unit_enabled")); ok {
			tmp := isInputOutputMemoryManagementUnitEnabled.(bool)
			details.IsInputOutputMemoryManagementUnitEnabled = &tmp
		}
		if isSymmetricMultiThreadingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_symmetric_multi_threading_enabled")); ok {
			tmp := isSymmetricMultiThreadingEnabled.(bool)
			details.IsSymmetricMultiThreadingEnabled = &tmp
		}
		if numaNodesPerSocket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "numa_nodes_per_socket")); ok {
			details.NumaNodesPerSocket = oci_core.InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum(numaNodesPerSocket.(string))
		}
		if percentageOfCoresEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "percentage_of_cores_enabled")); ok {
			tmp := percentageOfCoresEnabled.(int)
			details.PercentageOfCoresEnabled = &tmp
		}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isMemoryEncryptionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_memory_encryption_enabled")); ok {
			tmp := isMemoryEncryptionEnabled.(bool)
			details.IsMemoryEncryptionEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("AMD_ROME_BM_GPU"):
		details := oci_core.InstanceConfigurationAmdRomeBmGpuLaunchInstancePlatformConfig{}
		if areVirtualInstructionsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_virtual_instructions_enabled")); ok {
			tmp := areVirtualInstructionsEnabled.(bool)
			details.AreVirtualInstructionsEnabled = &tmp
		}
		if configMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_map")); ok {
			details.ConfigMap = tfresource.ObjectMapToStringMap(configMap.(map[string]interface{}))
		}
		if isAccessControlServiceEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_access_control_service_enabled")); ok {
			tmp := isAccessControlServiceEnabled.(bool)
			details.IsAccessControlServiceEnabled = &tmp
		}
		if isInputOutputMemoryManagementUnitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_input_output_memory_management_unit_enabled")); ok {
			tmp := isInputOutputMemoryManagementUnitEnabled.(bool)
			details.IsInputOutputMemoryManagementUnitEnabled = &tmp
		}
		if isSymmetricMultiThreadingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_symmetric_multi_threading_enabled")); ok {
			tmp := isSymmetricMultiThreadingEnabled.(bool)
			details.IsSymmetricMultiThreadingEnabled = &tmp
		}
		if numaNodesPerSocket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "numa_nodes_per_socket")); ok {
			details.NumaNodesPerSocket = oci_core.InstanceConfigurationAmdRomeBmGpuLaunchInstancePlatformConfigNumaNodesPerSocketEnum(numaNodesPerSocket.(string))
		}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isMemoryEncryptionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_memory_encryption_enabled")); ok {
			tmp := isMemoryEncryptionEnabled.(bool)
			details.IsMemoryEncryptionEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("AMD_VM"):
		details := oci_core.InstanceConfigurationAmdVmLaunchInstancePlatformConfig{}
		if isSymmetricMultiThreadingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_symmetric_multi_threading_enabled")); ok {
			tmp := isSymmetricMultiThreadingEnabled.(bool)
			details.IsSymmetricMultiThreadingEnabled = &tmp
		}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isMemoryEncryptionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_memory_encryption_enabled")); ok {
			tmp := isMemoryEncryptionEnabled.(bool)
			details.IsMemoryEncryptionEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("GENERIC_BM"):
		details := oci_core.InstanceConfigurationGenericBmLaunchInstancePlatformConfig{}
		if areVirtualInstructionsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_virtual_instructions_enabled")); ok {
			tmp := areVirtualInstructionsEnabled.(bool)
			details.AreVirtualInstructionsEnabled = &tmp
		}
		if configMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_map")); ok {
			details.ConfigMap = tfresource.ObjectMapToStringMap(configMap.(map[string]interface{}))
		}
		if isAccessControlServiceEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_access_control_service_enabled")); ok {
			tmp := isAccessControlServiceEnabled.(bool)
			details.IsAccessControlServiceEnabled = &tmp
		}
		if isInputOutputMemoryManagementUnitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_input_output_memory_management_unit_enabled")); ok {
			tmp := isInputOutputMemoryManagementUnitEnabled.(bool)
			details.IsInputOutputMemoryManagementUnitEnabled = &tmp
		}
		if isSymmetricMultiThreadingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_symmetric_multi_threading_enabled")); ok {
			tmp := isSymmetricMultiThreadingEnabled.(bool)
			details.IsSymmetricMultiThreadingEnabled = &tmp
		}
		if numaNodesPerSocket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "numa_nodes_per_socket")); ok {
			details.NumaNodesPerSocket = oci_core.InstanceConfigurationGenericBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum(numaNodesPerSocket.(string))
		}
		if percentageOfCoresEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "percentage_of_cores_enabled")); ok {
			tmp := percentageOfCoresEnabled.(int)
			details.PercentageOfCoresEnabled = &tmp
		}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isMemoryEncryptionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_memory_encryption_enabled")); ok {
			tmp := isMemoryEncryptionEnabled.(bool)
			details.IsMemoryEncryptionEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("INTEL_ICELAKE_BM"):
		details := oci_core.InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig{}
		if configMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_map")); ok {
			details.ConfigMap = tfresource.ObjectMapToStringMap(configMap.(map[string]interface{}))
		}
		if isInputOutputMemoryManagementUnitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_input_output_memory_management_unit_enabled")); ok {
			tmp := isInputOutputMemoryManagementUnitEnabled.(bool)
			details.IsInputOutputMemoryManagementUnitEnabled = &tmp
		}
		if isSymmetricMultiThreadingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_symmetric_multi_threading_enabled")); ok {
			tmp := isSymmetricMultiThreadingEnabled.(bool)
			details.IsSymmetricMultiThreadingEnabled = &tmp
		}
		if numaNodesPerSocket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "numa_nodes_per_socket")); ok {
			details.NumaNodesPerSocket = oci_core.InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum(numaNodesPerSocket.(string))
		}
		if percentageOfCoresEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "percentage_of_cores_enabled")); ok {
			tmp := percentageOfCoresEnabled.(int)
			details.PercentageOfCoresEnabled = &tmp
		}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isMemoryEncryptionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_memory_encryption_enabled")); ok {
			tmp := isMemoryEncryptionEnabled.(bool)
			details.IsMemoryEncryptionEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("INTEL_SKYLAKE_BM"):
		details := oci_core.InstanceConfigurationIntelSkylakeBmLaunchInstancePlatformConfig{}
		if configMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_map")); ok {
			details.ConfigMap = tfresource.ObjectMapToStringMap(configMap.(map[string]interface{}))
		}
		if isInputOutputMemoryManagementUnitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_input_output_memory_management_unit_enabled")); ok {
			tmp := isInputOutputMemoryManagementUnitEnabled.(bool)
			details.IsInputOutputMemoryManagementUnitEnabled = &tmp
		}
		if isSymmetricMultiThreadingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_symmetric_multi_threading_enabled")); ok {
			tmp := isSymmetricMultiThreadingEnabled.(bool)
			details.IsSymmetricMultiThreadingEnabled = &tmp
		}
		if numaNodesPerSocket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "numa_nodes_per_socket")); ok {
			details.NumaNodesPerSocket = oci_core.InstanceConfigurationIntelSkylakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum(numaNodesPerSocket.(string))
		}
		if percentageOfCoresEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "percentage_of_cores_enabled")); ok {
			tmp := percentageOfCoresEnabled.(int)
			details.PercentageOfCoresEnabled = &tmp
		}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isMemoryEncryptionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_memory_encryption_enabled")); ok {
			tmp := isMemoryEncryptionEnabled.(bool)
			details.IsMemoryEncryptionEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("INTEL_VM"):
		details := oci_core.InstanceConfigurationIntelVmLaunchInstancePlatformConfig{}
		if isSymmetricMultiThreadingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_symmetric_multi_threading_enabled")); ok {
			tmp := isSymmetricMultiThreadingEnabled.(bool)
			details.IsSymmetricMultiThreadingEnabled = &tmp
		}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isMemoryEncryptionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_memory_encryption_enabled")); ok {
			tmp := isMemoryEncryptionEnabled.(bool)
			details.IsMemoryEncryptionEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func InstanceConfigurationLaunchInstancePlatformConfigToMap(obj *oci_core.InstanceConfigurationLaunchInstancePlatformConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.InstanceConfigurationAmdMilanBmLaunchInstancePlatformConfig:
		result["type"] = "AMD_MILAN_BM"

		if v.AreVirtualInstructionsEnabled != nil {
			result["are_virtual_instructions_enabled"] = bool(*v.AreVirtualInstructionsEnabled)
		}

		result["config_map"] = v.ConfigMap
		result["config_map"] = v.ConfigMap

		if v.IsAccessControlServiceEnabled != nil {
			result["is_access_control_service_enabled"] = bool(*v.IsAccessControlServiceEnabled)
		}

		if v.IsInputOutputMemoryManagementUnitEnabled != nil {
			result["is_input_output_memory_management_unit_enabled"] = bool(*v.IsInputOutputMemoryManagementUnitEnabled)
		}

		if v.IsSymmetricMultiThreadingEnabled != nil {
			result["is_symmetric_multi_threading_enabled"] = bool(*v.IsSymmetricMultiThreadingEnabled)
		}

		result["numa_nodes_per_socket"] = string(v.NumaNodesPerSocket)

		if v.PercentageOfCoresEnabled != nil {
			result["percentage_of_cores_enabled"] = int(*v.PercentageOfCoresEnabled)
		}

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsMemoryEncryptionEnabled != nil {
			result["is_memory_encryption_enabled"] = bool(*v.IsMemoryEncryptionEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.InstanceConfigurationAmdMilanBmGpuLaunchInstancePlatformConfig:
		result["type"] = "AMD_MILAN_BM_GPU"

		if v.AreVirtualInstructionsEnabled != nil {
			result["are_virtual_instructions_enabled"] = bool(*v.AreVirtualInstructionsEnabled)
		}

		result["config_map"] = v.ConfigMap
		result["config_map"] = v.ConfigMap

		if v.IsAccessControlServiceEnabled != nil {
			result["is_access_control_service_enabled"] = bool(*v.IsAccessControlServiceEnabled)
		}

		if v.IsInputOutputMemoryManagementUnitEnabled != nil {
			result["is_input_output_memory_management_unit_enabled"] = bool(*v.IsInputOutputMemoryManagementUnitEnabled)
		}

		if v.IsSymmetricMultiThreadingEnabled != nil {
			result["is_symmetric_multi_threading_enabled"] = bool(*v.IsSymmetricMultiThreadingEnabled)
		}

		result["numa_nodes_per_socket"] = string(v.NumaNodesPerSocket)

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsMemoryEncryptionEnabled != nil {
			result["is_memory_encryption_enabled"] = bool(*v.IsMemoryEncryptionEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig:
		result["type"] = "AMD_ROME_BM"

		if v.AreVirtualInstructionsEnabled != nil {
			result["are_virtual_instructions_enabled"] = bool(*v.AreVirtualInstructionsEnabled)
		}

		result["config_map"] = v.ConfigMap
		result["config_map"] = v.ConfigMap

		if v.IsAccessControlServiceEnabled != nil {
			result["is_access_control_service_enabled"] = bool(*v.IsAccessControlServiceEnabled)
		}

		if v.IsInputOutputMemoryManagementUnitEnabled != nil {
			result["is_input_output_memory_management_unit_enabled"] = bool(*v.IsInputOutputMemoryManagementUnitEnabled)
		}

		if v.IsSymmetricMultiThreadingEnabled != nil {
			result["is_symmetric_multi_threading_enabled"] = bool(*v.IsSymmetricMultiThreadingEnabled)
		}

		result["numa_nodes_per_socket"] = string(v.NumaNodesPerSocket)

		if v.PercentageOfCoresEnabled != nil {
			result["percentage_of_cores_enabled"] = int(*v.PercentageOfCoresEnabled)
		}

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsMemoryEncryptionEnabled != nil {
			result["is_memory_encryption_enabled"] = bool(*v.IsMemoryEncryptionEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.InstanceConfigurationAmdRomeBmGpuLaunchInstancePlatformConfig:
		result["type"] = "AMD_ROME_BM_GPU"

		if v.AreVirtualInstructionsEnabled != nil {
			result["are_virtual_instructions_enabled"] = bool(*v.AreVirtualInstructionsEnabled)
		}

		result["config_map"] = v.ConfigMap
		result["config_map"] = v.ConfigMap

		if v.IsAccessControlServiceEnabled != nil {
			result["is_access_control_service_enabled"] = bool(*v.IsAccessControlServiceEnabled)
		}

		if v.IsInputOutputMemoryManagementUnitEnabled != nil {
			result["is_input_output_memory_management_unit_enabled"] = bool(*v.IsInputOutputMemoryManagementUnitEnabled)
		}

		if v.IsSymmetricMultiThreadingEnabled != nil {
			result["is_symmetric_multi_threading_enabled"] = bool(*v.IsSymmetricMultiThreadingEnabled)
		}

		result["numa_nodes_per_socket"] = string(v.NumaNodesPerSocket)

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsMemoryEncryptionEnabled != nil {
			result["is_memory_encryption_enabled"] = bool(*v.IsMemoryEncryptionEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.InstanceConfigurationAmdVmLaunchInstancePlatformConfig:
		result["type"] = "AMD_VM"

		if v.IsSymmetricMultiThreadingEnabled != nil {
			result["is_symmetric_multi_threading_enabled"] = bool(*v.IsSymmetricMultiThreadingEnabled)
		}

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsMemoryEncryptionEnabled != nil {
			result["is_memory_encryption_enabled"] = bool(*v.IsMemoryEncryptionEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.InstanceConfigurationGenericBmLaunchInstancePlatformConfig:
		result["type"] = "GENERIC_BM"

		if v.AreVirtualInstructionsEnabled != nil {
			result["are_virtual_instructions_enabled"] = bool(*v.AreVirtualInstructionsEnabled)
		}

		result["config_map"] = v.ConfigMap
		result["config_map"] = v.ConfigMap

		if v.IsAccessControlServiceEnabled != nil {
			result["is_access_control_service_enabled"] = bool(*v.IsAccessControlServiceEnabled)
		}

		if v.IsInputOutputMemoryManagementUnitEnabled != nil {
			result["is_input_output_memory_management_unit_enabled"] = bool(*v.IsInputOutputMemoryManagementUnitEnabled)
		}

		if v.IsSymmetricMultiThreadingEnabled != nil {
			result["is_symmetric_multi_threading_enabled"] = bool(*v.IsSymmetricMultiThreadingEnabled)
		}

		result["numa_nodes_per_socket"] = string(v.NumaNodesPerSocket)

		if v.PercentageOfCoresEnabled != nil {
			result["percentage_of_cores_enabled"] = int(*v.PercentageOfCoresEnabled)
		}

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsMemoryEncryptionEnabled != nil {
			result["is_memory_encryption_enabled"] = bool(*v.IsMemoryEncryptionEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig:
		result["type"] = "INTEL_ICELAKE_BM"

		result["config_map"] = v.ConfigMap
		result["config_map"] = v.ConfigMap

		if v.IsInputOutputMemoryManagementUnitEnabled != nil {
			result["is_input_output_memory_management_unit_enabled"] = bool(*v.IsInputOutputMemoryManagementUnitEnabled)
		}

		if v.IsSymmetricMultiThreadingEnabled != nil {
			result["is_symmetric_multi_threading_enabled"] = bool(*v.IsSymmetricMultiThreadingEnabled)
		}

		result["numa_nodes_per_socket"] = string(v.NumaNodesPerSocket)

		if v.PercentageOfCoresEnabled != nil {
			result["percentage_of_cores_enabled"] = int(*v.PercentageOfCoresEnabled)
		}

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsMemoryEncryptionEnabled != nil {
			result["is_memory_encryption_enabled"] = bool(*v.IsMemoryEncryptionEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.InstanceConfigurationIntelSkylakeBmLaunchInstancePlatformConfig:
		result["type"] = "INTEL_SKYLAKE_BM"

		result["config_map"] = v.ConfigMap
		result["config_map"] = v.ConfigMap

		if v.IsInputOutputMemoryManagementUnitEnabled != nil {
			result["is_input_output_memory_management_unit_enabled"] = bool(*v.IsInputOutputMemoryManagementUnitEnabled)
		}

		if v.IsSymmetricMultiThreadingEnabled != nil {
			result["is_symmetric_multi_threading_enabled"] = bool(*v.IsSymmetricMultiThreadingEnabled)
		}

		result["numa_nodes_per_socket"] = string(v.NumaNodesPerSocket)

		if v.PercentageOfCoresEnabled != nil {
			result["percentage_of_cores_enabled"] = int(*v.PercentageOfCoresEnabled)
		}

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsMemoryEncryptionEnabled != nil {
			result["is_memory_encryption_enabled"] = bool(*v.IsMemoryEncryptionEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.InstanceConfigurationIntelVmLaunchInstancePlatformConfig:
		result["type"] = "INTEL_VM"

		if v.IsSymmetricMultiThreadingEnabled != nil {
			result["is_symmetric_multi_threading_enabled"] = bool(*v.IsSymmetricMultiThreadingEnabled)
		}

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsMemoryEncryptionEnabled != nil {
			result["is_memory_encryption_enabled"] = bool(*v.IsMemoryEncryptionEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationLaunchInstanceShapeConfigDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationLaunchInstanceShapeConfigDetails, error) {
	result := oci_core.InstanceConfigurationLaunchInstanceShapeConfigDetails{}

	if baselineOcpuUtilization, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "baseline_ocpu_utilization")); ok {
		result.BaselineOcpuUtilization = oci_core.InstanceConfigurationLaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum(baselineOcpuUtilization.(string))
	}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	}

	if nvmes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nvmes")); ok {
		tmp := nvmes.(int)
		result.Nvmes = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	if vcpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcpus")); ok {
		tmp := vcpus.(int)
		result.Vcpus = &tmp
	}

	return result, nil
}

func InstanceConfigurationLaunchInstanceShapeConfigDetailsToMap(obj *oci_core.InstanceConfigurationLaunchInstanceShapeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["baseline_ocpu_utilization"] = string(obj.BaselineOcpuUtilization)

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.Nvmes != nil {
		result["nvmes"] = int(*obj.Nvmes)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	if obj.Vcpus != nil {
		result["vcpus"] = int(*obj.Vcpus)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationLaunchOptions(fieldKeyFormat string) (oci_core.InstanceConfigurationLaunchOptions, error) {
	result := oci_core.InstanceConfigurationLaunchOptions{}

	if bootVolumeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_type")); ok {
		result.BootVolumeType = oci_core.InstanceConfigurationLaunchOptionsBootVolumeTypeEnum(bootVolumeType.(string))
	}

	if firmware, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "firmware")); ok {
		result.Firmware = oci_core.InstanceConfigurationLaunchOptionsFirmwareEnum(firmware.(string))
	}

	if isConsistentVolumeNamingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_consistent_volume_naming_enabled")); ok {
		tmp := isConsistentVolumeNamingEnabled.(bool)
		result.IsConsistentVolumeNamingEnabled = &tmp
	}

	if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pv_encryption_in_transit_enabled")); ok {
		tmp := isPvEncryptionInTransitEnabled.(bool)
		result.IsPvEncryptionInTransitEnabled = &tmp
	}

	if networkType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_type")); ok {
		result.NetworkType = oci_core.InstanceConfigurationLaunchOptionsNetworkTypeEnum(networkType.(string))
	}

	if remoteDataVolumeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remote_data_volume_type")); ok {
		result.RemoteDataVolumeType = oci_core.InstanceConfigurationLaunchOptionsRemoteDataVolumeTypeEnum(remoteDataVolumeType.(string))
	}

	return result, nil
}

func InstanceConfigurationLaunchOptionsToMap(obj *oci_core.InstanceConfigurationLaunchOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["boot_volume_type"] = string(obj.BootVolumeType)

	result["firmware"] = string(obj.Firmware)

	if obj.IsConsistentVolumeNamingEnabled != nil {
		result["is_consistent_volume_naming_enabled"] = bool(*obj.IsConsistentVolumeNamingEnabled)
	}

	if obj.IsPvEncryptionInTransitEnabled != nil {
		result["is_pv_encryption_in_transit_enabled"] = bool(*obj.IsPvEncryptionInTransitEnabled)
	}

	result["network_type"] = string(obj.NetworkType)

	result["remote_data_volume_type"] = string(obj.RemoteDataVolumeType)

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationVolumeSourceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationVolumeSourceDetails, error) {
	var baseObject oci_core.InstanceConfigurationVolumeSourceDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("volume"):
		details := oci_core.InstanceConfigurationVolumeSourceFromVolumeDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	case strings.ToLower("volumeBackup"):
		details := oci_core.InstanceConfigurationVolumeSourceFromVolumeBackupDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func InstanceConfigurationVolumeSourceDetailsToMap(obj *oci_core.InstanceConfigurationVolumeSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.InstanceConfigurationVolumeSourceFromVolumeDetails:
		result["type"] = "volume"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	case oci_core.InstanceConfigurationVolumeSourceFromVolumeBackupDetails:
		result["type"] = "volumeBackup"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToLaunchInstanceLicensingConfig(fieldKeyFormat string) (oci_core.LaunchInstanceLicensingConfig, error) {
	var baseObject oci_core.LaunchInstanceLicensingConfig
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("WINDOWS"):
		details := oci_core.LaunchInstanceWindowsLicensingConfig{}
		if licenseType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "license_type")); ok {
			tmp := licenseType.(string)
			details.LicenseType = oci_core.LaunchInstanceLicensingConfigLicenseTypeEnum(tmp)
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func LaunchInstanceLicensingConfigToMap(obj oci_core.LaunchInstanceLicensingConfig) map[string]interface{} {
	result := map[string]interface{}{}
	if obj, ok := obj.(oci_core.LaunchInstanceWindowsLicensingConfig); ok {
		result["type"] = "WINDOWS"
		result["license_type"] = string(obj.LicenseType)
	} else {
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}
	/*switch v := (obj).(type) {
	case oci_core.LaunchInstanceWindowsLicensingConfig:
		result["type"] = "WINDOWS"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}*/

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToPreemptibleInstanceConfigDetails(fieldKeyFormat string) (oci_core.PreemptibleInstanceConfigDetails, error) {
	result := oci_core.PreemptibleInstanceConfigDetails{}

	if preemptionAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preemption_action")); ok {
		if tmpList := preemptionAction.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "preemption_action"), 0)
			tmp, err := s.mapToPreemptionAction(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert preemption_action, encountered error: %v", err)
			}
			result.PreemptionAction = tmp
		}
	}

	return result, nil
}

func (s *CoreInstanceConfigurationResourceCrud) mapToPreemptionAction(fieldKeyFormat string) (oci_core.PreemptionAction, error) {
	var baseObject oci_core.PreemptionAction
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("TERMINATE"):
		details := oci_core.TerminatePreemptionAction{}
		if preserveBootVolume, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preserve_boot_volume")); ok {
			tmp := preserveBootVolume.(bool)
			details.PreserveBootVolume = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *CoreInstanceConfigurationResourceCrud) populateTopLevelPolymorphicCreateInstanceConfigurationRequest(request *oci_core.CreateInstanceConfigurationRequest) error {
	//discriminator
	sourceRaw, ok := s.D.GetOkExists("source")
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "NONE" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("INSTANCE"):
		details := oci_core.CreateInstanceConfigurationFromInstanceDetails{}
		if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
			tmp := instanceId.(string)
			details.InstanceId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateInstanceConfiguration = details
	case strings.ToLower("NONE"):
		details := oci_core.CreateInstanceConfigurationDetails{}
		if instanceDetails, ok := s.D.GetOkExists("instance_details"); ok {
			if tmpList := instanceDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "instance_details", 0)
				tmp, err := s.mapToInstanceConfigurationInstanceDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.InstanceDetails = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateInstanceConfiguration = details
	default:
		return fmt.Errorf("unknown source '%v' was specified", source)
	}
	return nil
}

func (s *CoreInstanceConfigurationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeInstanceConfigurationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.InstanceConfigurationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeInstanceConfigurationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
