// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_migrations

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudMigrationsTargetAssetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudMigrationsTargetAsset,
		Read:     readCloudMigrationsTargetAsset,
		Update:   updateCloudMigrationsTargetAsset,
		Delete:   deleteCloudMigrationsTargetAsset,
		Schema: map[string]*schema.Schema{
			// Required
			"is_excluded_from_execution": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"migration_plan_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"preferred_shape_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"INSTANCE",
				}, true),
			},
			"user_spec": {
				Type:     schema.TypeList,
				Required: true,
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
									},
									"is_management_disabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"is_monitoring_disabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"plugins_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"desired_state": {
													Type:     schema.TypeString,
													Required: true,
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
						"availability_domain": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},
						"capacity_reservation_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"assign_private_dns_record": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"assign_public_ip": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
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
									"hostname_label": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"nsg_ids": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Set:      tfresource.LiteralTypeHashCodeForSets,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"private_ip": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"skip_source_dest_check": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vlan_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"dedicated_vm_host_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"fault_domain": {
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
						"hostname_label": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"instance_options": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
									},

									// Computed
								},
							},
						},
						"ipxe_script": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_pv_encryption_in_transit_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"preemptible_instance_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"preemption_action": {
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
														"TERMINATE",
													}, true),
												},

												// Optional
												"preserve_boot_volume": {
													Type:     schema.TypeBool,
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
						"shape": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
									},
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"ocpus": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"source_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"source_type": {
										Type:             schema.TypeString,
										Required:         true,
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
									},
									"boot_volume_size_in_gbs": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ValidateFunc:     tfresource.ValidateInt64TypeString,
										DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
									},
									"boot_volume_vpus_per_gb": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ValidateFunc:     tfresource.ValidateInt64TypeString,
										DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
									},
									"image_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"kms_key_id": {
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
			"block_volumes_performance": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ms_license": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compatibility_messages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"severity": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"created_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"estimated_cost": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compute": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"gpu_count": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"gpu_per_hour": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"gpu_per_hour_by_subscription": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"memory_amount_gb": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"memory_gb_per_hour": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"memory_gb_per_hour_by_subscription": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"ocpu_count": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"ocpu_per_hour": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"ocpu_per_hour_by_subscription": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"total_per_hour": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"total_per_hour_by_subscription": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"currency_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_image": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"total_per_hour": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"total_per_hour_by_subscription": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"storage": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"total_gb_per_month": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"total_gb_per_month_by_subscription": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"volumes": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"capacity_gb": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"total_gb_per_month": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"total_gb_per_month_by_subscription": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"subscription_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_estimation_per_month": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"total_estimation_per_month_by_subscription": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"migration_asset": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"depended_on_by": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"depends_on": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"migration_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"notifications": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"parent_snapshot": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"replication_compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"replication_schedule_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"snap_shot_bucket_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"snapshots": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"source_asset_data": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"source_asset_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenancy_id": {
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
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"recommended_spec": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"agent_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"are_all_plugins_disabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_management_disabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_monitoring_disabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"plugins_config": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"desired_state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"capacity_reservation_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_vnic_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assign_private_dns_record": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"assign_public_ip": {
										Type:     schema.TypeBool,
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
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"hostname_label": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"nsg_ids": {
										Type:     schema.TypeSet,
										Computed: true,
										Set:      tfresource.LiteralTypeHashCodeForSets,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"private_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"skip_source_dest_check": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vlan_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"dedicated_vm_host_id": {
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
						"fault_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"hostname_label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_options": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"are_legacy_imds_endpoints_disabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"ipxe_script": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_pv_encryption_in_transit_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"preemptible_instance_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"preemption_action": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"preserve_boot_volume": {
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
						"shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"baseline_ocpu_utilization": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"ocpus": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"source_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"boot_volume_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"boot_volume_size_in_gbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"boot_volume_vpus_per_gb": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"image_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"kms_key_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"test_spec": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"agent_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"are_all_plugins_disabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_management_disabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_monitoring_disabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"plugins_config": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"desired_state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"capacity_reservation_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_vnic_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assign_private_dns_record": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"assign_public_ip": {
										Type:     schema.TypeBool,
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
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"hostname_label": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"nsg_ids": {
										Type:     schema.TypeSet,
										Computed: true,
										Set:      tfresource.LiteralTypeHashCodeForSets,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"private_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"skip_source_dest_check": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vlan_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"dedicated_vm_host_id": {
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
						"fault_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"hostname_label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_options": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"are_legacy_imds_endpoints_disabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"ipxe_script": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_pv_encryption_in_transit_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"preemptible_instance_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"preemption_action": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"preserve_boot_volume": {
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
						"shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"baseline_ocpu_utilization": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"ocpus": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"source_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"boot_volume_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"boot_volume_size_in_gbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"boot_volume_vpus_per_gb": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"image_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"kms_key_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"time_assessed": {
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

func createCloudMigrationsTargetAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsTargetAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudMigrationsTargetAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsTargetAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.ReadResource(sync)
}

func updateCloudMigrationsTargetAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsTargetAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudMigrationsTargetAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsTargetAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudMigrationsTargetAssetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_migrations.MigrationClient
	Res                    *oci_cloud_migrations.TargetAsset
	DisableNotFoundRetries bool
}

func (s *CloudMigrationsTargetAssetResourceCrud) ID() string {
	targetAsset := *s.Res
	return *targetAsset.GetId()
}

func (s *CloudMigrationsTargetAssetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_migrations.TargetAssetLifecycleStateCreating),
	}
}

func (s *CloudMigrationsTargetAssetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_migrations.TargetAssetLifecycleStateNeedsAttention),
		string(oci_cloud_migrations.TargetAssetLifecycleStateActive),
	}
}

func (s *CloudMigrationsTargetAssetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_migrations.TargetAssetLifecycleStateDeleting),
	}
}

func (s *CloudMigrationsTargetAssetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_migrations.TargetAssetLifecycleStateDeleted),
	}
}

func (s *CloudMigrationsTargetAssetResourceCrud) Create() error {
	request := oci_cloud_migrations.CreateTargetAssetRequest{}
	err := s.populateTopLevelPolymorphicCreateTargetAssetRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.CreateTargetAsset(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getTargetAssetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations"), oci_cloud_migrations.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *CloudMigrationsTargetAssetResourceCrud) getTargetAssetFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_cloud_migrations.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	targetAssetId, err := targetAssetWaitForWorkRequest(workId, "targetasset",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, targetAssetId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_cloud_migrations.CancelWorkRequestRequest{
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
	s.D.SetId(*targetAssetId)

	return s.Get()
}

func targetAssetWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "cloud_migrations", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_cloud_migrations.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func targetAssetWaitForWorkRequest(wId *string, entityType string, action oci_cloud_migrations.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_cloud_migrations.MigrationClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "cloud_migrations")
	retryPolicy.ShouldRetryOperation = targetAssetWorkRequestShouldRetryFunc(timeout)

	response := oci_cloud_migrations.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_cloud_migrations.OperationStatusInProgress),
			string(oci_cloud_migrations.OperationStatusAccepted),
			string(oci_cloud_migrations.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_cloud_migrations.OperationStatusSucceeded),
			string(oci_cloud_migrations.OperationStatusFailed),
			string(oci_cloud_migrations.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_cloud_migrations.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_cloud_migrations.OperationStatusFailed || response.Status == oci_cloud_migrations.OperationStatusCanceled {
		return nil, getErrorFromCloudMigrationsTargetAssetWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromCloudMigrationsTargetAssetWorkRequest(client *oci_cloud_migrations.MigrationClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_cloud_migrations.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_cloud_migrations.ListWorkRequestErrorsRequest{
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

func (s *CloudMigrationsTargetAssetResourceCrud) Get() error {
	request := oci_cloud_migrations.GetTargetAssetRequest{}

	tmp := s.D.Id()
	request.TargetAssetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.GetTargetAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TargetAsset
	return nil
}

func (s *CloudMigrationsTargetAssetResourceCrud) Update() error {
	request := oci_cloud_migrations.UpdateTargetAssetRequest{}
	err := s.populateTopLevelPolymorphicUpdateTargetAssetRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.UpdateTargetAsset(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getTargetAssetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations"), oci_cloud_migrations.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *CloudMigrationsTargetAssetResourceCrud) Delete() error {
	request := oci_cloud_migrations.DeleteTargetAssetRequest{}

	tmp := s.D.Id()
	request.TargetAssetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.DeleteTargetAsset(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := targetAssetWaitForWorkRequest(workId, "targetasset",
		oci_cloud_migrations.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *CloudMigrationsTargetAssetResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_cloud_migrations.VmTargetAsset:
		s.D.Set("type", "INSTANCE")

		if v.BlockVolumesPerformance != nil {
			s.D.Set("block_volumes_performance", *v.BlockVolumesPerformance)
		}

		if v.MsLicense != nil {
			s.D.Set("ms_license", *v.MsLicense)
		}

		s.D.Set("preferred_shape_type", v.PreferredShapeType)

		if v.RecommendedSpec != nil {
			s.D.Set("recommended_spec", []interface{}{LaunchInstanceDetailsToMap(v.RecommendedSpec, false)})
		} else {
			s.D.Set("recommended_spec", nil)
		}

		if v.TestSpec != nil {
			s.D.Set("test_spec", []interface{}{LaunchInstanceDetailsToMap(v.TestSpec, false)})
		} else {
			s.D.Set("test_spec", nil)
		}

		if v.UserSpec != nil {
			s.D.Set("user_spec", []interface{}{LaunchInstanceDetailsToMap(v.UserSpec, false)})
		} else {
			s.D.Set("user_spec", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		compatibilityMessages := []interface{}{}
		for _, item := range v.CompatibilityMessages {
			compatibilityMessages = append(compatibilityMessages, CompatibilityMessageToMap(item))
		}
		s.D.Set("compatibility_messages", compatibilityMessages)

		if v.CreatedResourceId != nil {
			s.D.Set("created_resource_id", *v.CreatedResourceId)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.EstimatedCost != nil {
			s.D.Set("estimated_cost", []interface{}{CostEstimationToMap(v.EstimatedCost)})
		} else {
			s.D.Set("estimated_cost", nil)
		}

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		if v.IsExcludedFromExecution != nil {
			s.D.Set("is_excluded_from_execution", *v.IsExcludedFromExecution)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.MigrationAsset != nil {
			s.D.Set("migration_asset", []interface{}{MigrationAssetToMap(v.MigrationAsset)})
		} else {
			s.D.Set("migration_asset", nil)
		}

		if v.MigrationPlanId != nil {
			s.D.Set("migration_plan_id", *v.MigrationPlanId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeAssessed != nil {
			s.D.Set("time_assessed", v.TimeAssessed.String())
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func CompatibilityMessageToMap(obj oci_cloud_migrations.CompatibilityMessage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	result["name"] = string(obj.Name)

	result["severity"] = string(obj.Severity)

	return result
}

func (s *CloudMigrationsTargetAssetResourceCrud) mapToCreateVnicDetails(fieldKeyFormat string) (oci_cloud_migrations.CreateVnicDetails, error) {
	result := oci_cloud_migrations.CreateVnicDetails{}

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
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	if vlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vlan_id")); ok {
		tmp := vlanId.(string)
		result.VlanId = &tmp
	}

	return result, nil
}

func CreateVnicDetailsToMap(obj *oci_cloud_migrations.CreateVnicDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssignPrivateDnsRecord != nil {
		result["assign_private_dns_record"] = bool(*obj.AssignPrivateDnsRecord)
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
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
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

	if obj.VlanId != nil {
		result["vlan_id"] = string(*obj.VlanId)
	}

	return result
}

func (s *CloudMigrationsTargetAssetResourceCrud) mapToInstanceAgentPluginConfigDetails(fieldKeyFormat string) (oci_cloud_migrations.InstanceAgentPluginConfigDetails, error) {
	result := oci_cloud_migrations.InstanceAgentPluginConfigDetails{}

	if desiredState, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "desired_state")); ok {
		result.DesiredState = oci_cloud_migrations.InstanceAgentPluginConfigDetailsDesiredStateEnum(desiredState.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func InstanceAgentPluginConfigDetailsToMap(obj oci_cloud_migrations.InstanceAgentPluginConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["desired_state"] = string(obj.DesiredState)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *CloudMigrationsTargetAssetResourceCrud) mapToInstanceOptions(fieldKeyFormat string) (oci_cloud_migrations.InstanceOptions, error) {
	result := oci_cloud_migrations.InstanceOptions{}

	if areLegacyImdsEndpointsDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_legacy_imds_endpoints_disabled")); ok {
		tmp := areLegacyImdsEndpointsDisabled.(bool)
		result.AreLegacyImdsEndpointsDisabled = &tmp
	}

	return result, nil
}

func InstanceOptionsToMap(obj *oci_cloud_migrations.InstanceOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AreLegacyImdsEndpointsDisabled != nil {
		result["are_legacy_imds_endpoints_disabled"] = bool(*obj.AreLegacyImdsEndpointsDisabled)
	}

	return result
}

func (s *CloudMigrationsTargetAssetResourceCrud) mapToInstanceSourceDetails(fieldKeyFormat string) (oci_cloud_migrations.InstanceSourceDetails, error) {
	var baseObject oci_cloud_migrations.InstanceSourceDetails
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
		details := oci_cloud_migrations.InstanceSourceViaBootVolumeDetails{}
		if bootVolumeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_id")); ok {
			tmp := bootVolumeId.(string)
			details.BootVolumeId = &tmp
		}
		baseObject = details
	case strings.ToLower("image"):
		details := oci_cloud_migrations.InstanceSourceViaImageDetails{}
		if bootVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_size_in_gbs")); ok {
			tmp := bootVolumeSizeInGBs.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert bootVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.BootVolumeSizeInGBs = &tmpInt64
		}
		if bootVolumeVpusPerGB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_vpus_per_gb")); ok {
			tmp := bootVolumeVpusPerGB.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert bootVolumeVpusPerGB string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.BootVolumeVpusPerGB = &tmpInt64
		}
		if imageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_id")); ok {
			tmp := imageId.(string)
			details.ImageId = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func InstanceSourceDetailsToMap(obj *oci_cloud_migrations.InstanceSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_cloud_migrations.InstanceSourceViaBootVolumeDetails:
		result["source_type"] = "bootVolume"

		if v.BootVolumeId != nil {
			result["boot_volume_id"] = string(*v.BootVolumeId)
		}
	case oci_cloud_migrations.InstanceSourceViaImageDetails:
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
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CloudMigrationsTargetAssetResourceCrud) mapToLaunchInstanceAgentConfigDetails(fieldKeyFormat string) (oci_cloud_migrations.LaunchInstanceAgentConfigDetails, error) {
	result := oci_cloud_migrations.LaunchInstanceAgentConfigDetails{}

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
		tmp := make([]oci_cloud_migrations.InstanceAgentPluginConfigDetails, len(interfaces))
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

func LaunchInstanceAgentConfigDetailsToMap(obj *oci_cloud_migrations.LaunchInstanceAgentConfigDetails) map[string]interface{} {
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

func (s *CloudMigrationsTargetAssetResourceCrud) mapToLaunchInstanceDetails(fieldKeyFormat string) (oci_cloud_migrations.LaunchInstanceDetails, error) {
	result := oci_cloud_migrations.LaunchInstanceDetails{}

	if agentConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "agent_config")); ok {
		if tmpList := agentConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "agent_config"), 0)
			tmp, err := s.mapToLaunchInstanceAgentConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert agent_config, encountered error: %v", err)
			}
			result.AgentConfig = &tmp
		}
	}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		if availabilityDomain != "" {
			result.AvailabilityDomain = &tmp
		}
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
			tmp, err := s.mapToCreateVnicDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert create_vnic_details, encountered error: %v", err)
			}
			result.CreateVnicDetails = &tmp
		}
	}

	if dedicatedVmHostId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dedicated_vm_host_id")); ok {
		tmp := dedicatedVmHostId.(string)
		if dedicatedVmHostId != "" {
			result.DedicatedVmHostId = &tmp
		}
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

	if faultDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_domain")); ok {
		tmp := faultDomain.(string)
		result.FaultDomain = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
	}

	if instanceOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_options")); ok {
		if tmpList := instanceOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "instance_options"), 0)
			tmp, err := s.mapToInstanceOptions(fieldKeyFormatNextLevel)
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

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
	}

	if shapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_config")); ok {
		if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "shape_config"), 0)
			tmp, err := s.mapToLaunchInstanceShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert shape_config, encountered error: %v", err)
			}
			result.ShapeConfig = &tmp
		}
	}

	if sourceDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_details")); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_details"), 0)
			tmp, err := s.mapToInstanceSourceDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_details, encountered error: %v", err)
			}
			result.SourceDetails = tmp
		}
	}

	return result, nil
}

func LaunchInstanceDetailsToMap(obj *oci_cloud_migrations.LaunchInstanceDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentConfig != nil {
		result["agent_config"] = []interface{}{LaunchInstanceAgentConfigDetailsToMap(obj.AgentConfig)}
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
		result["create_vnic_details"] = []interface{}{CreateVnicDetailsToMap(obj.CreateVnicDetails, datasource)}
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

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.HostnameLabel != nil {
		result["hostname_label"] = string(*obj.HostnameLabel)
	}

	if obj.InstanceOptions != nil {
		result["instance_options"] = []interface{}{InstanceOptionsToMap(obj.InstanceOptions)}
	}

	if obj.IpxeScript != nil {
		result["ipxe_script"] = string(*obj.IpxeScript)
	}

	if obj.IsPvEncryptionInTransitEnabled != nil {
		result["is_pv_encryption_in_transit_enabled"] = bool(*obj.IsPvEncryptionInTransitEnabled)
	}

	if obj.PreemptibleInstanceConfig != nil {
		result["preemptible_instance_config"] = []interface{}{PreemptibleInstanceConfigDetailsToMap(obj.PreemptibleInstanceConfig)}
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.ShapeConfig != nil {
		result["shape_config"] = []interface{}{LaunchInstanceShapeConfigDetailsToMap(obj.ShapeConfig)}
	}

	if obj.SourceDetails != nil {
		var sourceDetailsArray []interface{}
		if sourceDetailsMap := InstanceSourceDetailsToMap(&obj.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		result["source_details"] = sourceDetailsArray
	}

	return result
}

func (s *CloudMigrationsTargetAssetResourceCrud) mapToLaunchInstanceShapeConfigDetails(fieldKeyFormat string) (oci_cloud_migrations.LaunchInstanceShapeConfigDetails, error) {
	result := oci_cloud_migrations.LaunchInstanceShapeConfigDetails{}

	if baselineOcpuUtilization, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "baseline_ocpu_utilization")); ok {
		result.BaselineOcpuUtilization = oci_cloud_migrations.LaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum(baselineOcpuUtilization.(string))
	}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp, ok := memoryInGBs.(float32)
		if !ok {
			tmp = float32(memoryInGBs.(float64))
		}
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp, ok := ocpus.(float32)
		if !ok {
			tmp = float32(ocpus.(float64))
		}
		result.Ocpus = &tmp
	}

	return result, nil
}

func LaunchInstanceShapeConfigDetailsToMap(obj *oci_cloud_migrations.LaunchInstanceShapeConfigDetails) map[string]interface{} {
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

func MigrationAssetToMap(obj *oci_cloud_migrations.MigrationAsset) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["depended_on_by"] = obj.DependedOnBy
	result["depended_on_by"] = obj.DependedOnBy

	result["depends_on"] = obj.DependsOn
	result["depends_on"] = obj.DependsOn

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MigrationId != nil {
		result["migration_id"] = string(*obj.MigrationId)
	}

	result["notifications"] = obj.Notifications
	result["notifications"] = obj.Notifications

	if obj.ParentSnapshot != nil {
		result["parent_snapshot"] = string(*obj.ParentSnapshot)
	}

	if obj.ReplicationCompartmentId != nil {
		result["replication_compartment_id"] = string(*obj.ReplicationCompartmentId)
	}

	if obj.ReplicationScheduleId != nil {
		result["replication_schedule_id"] = string(*obj.ReplicationScheduleId)
	}

	if obj.SnapShotBucketName != nil {
		result["snap_shot_bucket_name"] = string(*obj.SnapShotBucketName)
	}

	result["snapshots"] = obj.Snapshots
	result["snapshots"] = obj.Snapshots

	if obj.SourceAssetId != nil {
		result["source_asset_id"] = string(*obj.SourceAssetId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func (s *CloudMigrationsTargetAssetResourceCrud) mapToPreemptibleInstanceConfigDetails(fieldKeyFormat string) (oci_cloud_migrations.PreemptibleInstanceConfigDetails, error) {
	result := oci_cloud_migrations.PreemptibleInstanceConfigDetails{}

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

func PreemptibleInstanceConfigDetailsToMap(obj *oci_cloud_migrations.PreemptibleInstanceConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PreemptionAction != nil {
		preemptionActionArray := []interface{}{}
		if preemptionActionMap := PreemptionActionToMap(&obj.PreemptionAction); preemptionActionMap != nil {
			preemptionActionArray = append(preemptionActionArray, preemptionActionMap)
		}
		result["preemption_action"] = preemptionActionArray
	}

	return result
}

func (s *CloudMigrationsTargetAssetResourceCrud) mapToPreemptionAction(fieldKeyFormat string) (oci_cloud_migrations.PreemptionAction, error) {
	var baseObject oci_cloud_migrations.PreemptionAction
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
		details := oci_cloud_migrations.TerminatePreemptionAction{}
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

func PreemptionActionToMap(obj *oci_cloud_migrations.PreemptionAction) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_cloud_migrations.TerminatePreemptionAction:
		result["type"] = "TERMINATE"

		if v.PreserveBootVolume != nil {
			result["preserve_boot_volume"] = bool(*v.PreserveBootVolume)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func TargetAssetSummaryToMap(obj oci_cloud_migrations.TargetAssetSummary, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_cloud_migrations.VmTargetAssetSummary:
		result["type"] = "INSTANCE"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
		result["state"] = string(v.LifecycleState)

		if v.BlockVolumesPerformance != nil {
			result["block_volumes_performance"] = int(*v.BlockVolumesPerformance)
		}

		if v.MsLicense != nil {
			result["ms_license"] = string(*v.MsLicense)
		}

		result["preferred_shape_type"] = string(v.PreferredShapeType)

		if v.RecommendedSpec != nil {
			result["recommended_spec"] = []interface{}{LaunchInstanceDetailsToMap(v.RecommendedSpec, datasource)}
		}

		if v.UserSpec != nil {
			result["user_spec"] = []interface{}{LaunchInstanceDetailsToMap(v.UserSpec, datasource)}
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *CloudMigrationsTargetAssetResourceCrud) populateTopLevelPolymorphicCreateTargetAssetRequest(request *oci_cloud_migrations.CreateTargetAssetRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("INSTANCE"):
		details := oci_cloud_migrations.CreateVmTargetAssetDetails{}
		if blockVolumesPerformance, ok := s.D.GetOkExists("block_volumes_performance"); ok {
			tmp := blockVolumesPerformance.(int)
			details.BlockVolumesPerformance = &tmp
		}
		if msLicense, ok := s.D.GetOkExists("ms_license"); ok {
			tmp := msLicense.(string)
			details.MsLicense = &tmp
		}
		if preferredShapeType, ok := s.D.GetOkExists("preferred_shape_type"); ok {
			details.PreferredShapeType = oci_cloud_migrations.VmTargetAssetPreferredShapeTypeEnum(preferredShapeType.(string))
		}
		if userSpec, ok := s.D.GetOkExists("user_spec"); ok {
			if tmpList := userSpec.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user_spec", 0)
				tmp, err := s.mapToLaunchInstanceDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.UserSpec = &tmp
			}
		}
		if isExcludedFromExecution, ok := s.D.GetOkExists("is_excluded_from_execution"); ok {
			tmp := isExcludedFromExecution.(bool)
			details.IsExcludedFromExecution = &tmp
		}
		if migrationPlanId, ok := s.D.GetOkExists("migration_plan_id"); ok {
			tmp := migrationPlanId.(string)
			details.MigrationPlanId = &tmp
		}
		request.CreateTargetAssetDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *CloudMigrationsTargetAssetResourceCrud) populateTopLevelPolymorphicUpdateTargetAssetRequest(request *oci_cloud_migrations.UpdateTargetAssetRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("INSTANCE"):
		details := oci_cloud_migrations.UpdateVmTargetAssetDetails{}
		if blockVolumesPerformance, ok := s.D.GetOkExists("block_volumes_performance"); ok {
			tmp := blockVolumesPerformance.(int)
			details.BlockVolumesPerformance = &tmp
		}
		if msLicense, ok := s.D.GetOkExists("ms_license"); ok {
			tmp := msLicense.(string)
			details.MsLicense = &tmp
		}
		if preferredShapeType, ok := s.D.GetOkExists("preferred_shape_type"); ok {
			details.PreferredShapeType = oci_cloud_migrations.VmTargetAssetPreferredShapeTypeEnum(preferredShapeType.(string))
		}
		if userSpec, ok := s.D.GetOkExists("user_spec"); ok {
			if tmpList := userSpec.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user_spec", 0)
				tmp, err := s.mapToLaunchInstanceDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.UserSpec = &tmp
			}
		}
		if isExcludedFromExecution, ok := s.D.GetOkExists("is_excluded_from_execution"); ok {
			tmp := isExcludedFromExecution.(bool)
			details.IsExcludedFromExecution = &tmp
		}
		tmp := s.D.Id()
		request.TargetAssetId = &tmp
		request.UpdateTargetAssetDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
