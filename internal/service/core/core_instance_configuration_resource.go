// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"
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
													ValidateFunc:     utils.ValidateInt64TypeString,
													DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
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
													ValidateFunc:     utils.ValidateInt64TypeString,
													DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
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
															"desired_state": {
																Type:     schema.TypeString,
																Required: true,
																ForceNew: true,
															},
															"name": {
																Type:     schema.TypeString,
																Required: true,
																ForceNew: true,
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
												"nsg_ids": {
													Type:     schema.TypeSet,
													Optional: true,
													ForceNew: true,
													Set:      utils.LiteralTypeHashCodeForSets,
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
										DiffSuppressFunc: utils.JsonStringDiffSuppressFunction,
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
														"AMD_VM",
														"INTEL_SKYLAKE_BM",
														"INTEL_VM",
													}, true),
												},

												// Optional
												"is_measured_boot_enabled": {
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
												"preemption_action": {
													Type:     schema.TypeList,
													Required: true,
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

												// Optional

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
												"ocpus": {
													Type:     schema.TypeFloat,
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
													ValidateFunc:     utils.ValidateInt64TypeString,
													DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
												},
												"image_id": {
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
												"nsg_ids": {
													Type:     schema.TypeSet,
													Optional: true,
													ForceNew: true,
													Set:      utils.LiteralTypeHashCodeForSets,
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationAvailabilityConfig(fieldKeyFormat string) (oci_core.InstanceConfigurationAvailabilityConfig, error) {
	result := oci_core.InstanceConfigurationAvailabilityConfig{}

	if recoveryAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_action")); ok {
		result.RecoveryAction = oci_core.InstanceConfigurationAvailabilityConfigRecoveryActionEnum(recoveryAction.(string))
	}

	return result, nil
}

func InstanceConfigurationAvailabilityConfigToMap(obj *oci_core.InstanceConfigurationAvailabilityConfig) map[string]interface{} {
	result := map[string]interface{}{}

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

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationCreateVnicDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationCreateVnicDetails, error) {
	result := oci_core.InstanceConfigurationCreateVnicDetails{}

	if assignPrivateDnsRecord, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_private_dns_record")); ok {
		tmp := assignPrivateDnsRecord.(bool)
		result.AssignPrivateDnsRecord = &tmp
	}

	if assignPublicIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_public_ip")); ok {
		tmp := assignPublicIp.(bool)
		result.AssignPublicIp = &tmp
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
		result.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
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

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(utils.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

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

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if backupPolicyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_policy_id")); ok {
		tmp := backupPolicyId.(string)
		result.BackupPolicyId = &tmp
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
		result.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	return result, nil
}

func InstanceConfigurationCreateVolumeDetailsToMap(obj *oci_core.InstanceConfigurationCreateVolumeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.BackupPolicyId != nil {
		result["backup_policy_id"] = string(*obj.BackupPolicyId)
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
		instanceType = "" // default value
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
		if imageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_id")); ok {
			tmp := imageId.(string)
			details.ImageId = &tmp
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

		if v.ImageId != nil {
			result["image_id"] = string(*v.ImageId)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
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
		result.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		result.Metadata = utils.ObjectMapToStringMap(metadata.(map[string]interface{}))
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
		if numaNodesPerSocket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "numa_nodes_per_socket")); ok {
			details.NumaNodesPerSocket = oci_core.InstanceConfigurationAmdMilanBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum(numaNodesPerSocket.(string))
		}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
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
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
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
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
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
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
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
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
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

		result["numa_nodes_per_socket"] = string(v.NumaNodesPerSocket)
	case oci_core.InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig:
		result["type"] = "AMD_ROME_BM"
	case oci_core.InstanceConfigurationAmdVmLaunchInstancePlatformConfig:
		result["type"] = "AMD_VM"
	case oci_core.InstanceConfigurationIntelSkylakeBmLaunchInstancePlatformConfig:
		result["type"] = "INTEL_SKYLAKE_BM"
	case oci_core.InstanceConfigurationIntelVmLaunchInstancePlatformConfig:
		result["type"] = "INTEL_VM"
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

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	return result, nil
}

func InstanceConfigurationLaunchInstanceShapeConfigDetailsToMap(obj *oci_core.InstanceConfigurationLaunchInstanceShapeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["baseline_ocpu_utilization"] = string(obj.BaselineOcpuUtilization)

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
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
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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
