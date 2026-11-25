// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"
	"fmt"
	"log"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedAutonomousDatabaseResource() *schema.Resource {
	resource := &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		CustomizeDiff: validateDistributedAutonomousDatabaseDiff,
		CreateContext: createDistributedDatabaseDistributedAutonomousDatabaseWithContext,
		ReadContext:   readDistributedDatabaseDistributedAutonomousDatabaseWithContext,
		UpdateContext: updateDistributedDatabaseDistributedAutonomousDatabaseWithContext,
		DeleteContext: deleteDistributedDatabaseDistributedAutonomousDatabaseWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"catalog_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"original_replica": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"database": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"cloud_autonomous_vm_cluster_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"compute_count": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"data_storage_size_in_gbs": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"ADBD_EXISTING_CLUSTER",
													}, true),
												},

												// Optional
												"admin_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"admin_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"admin_password_secret_version_number": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"fast_start_fail_over_lag_limit_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_auto_scaling_enabled": {
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
												"kms_key_version_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_end_point_group_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_key_store_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"protection_mode": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"standby_maintenance_buffer_in_days": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vault_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
												"container_database_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"metadata": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"map": {
																Type:     schema.TypeMap,
																Computed: true,
																Elem:     schema.TypeString,
															},
														},
													},
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"supporting_resource_id": {
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
										},
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

						// Optional
						"data_guard_replicas": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"database": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"cloud_autonomous_vm_cluster_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"compute_count": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"data_storage_size_in_gbs": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"ADBD_EXISTING_CLUSTER",
													}, true),
												},

												// Optional
												"admin_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"admin_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"admin_password_secret_version_number": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"fast_start_fail_over_lag_limit_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_auto_scaling_enabled": {
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
												"kms_key_version_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_end_point_group_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_key_store_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"protection_mode": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"standby_maintenance_buffer_in_days": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vault_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
												"container_database_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"metadata": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"map": {
																Type:     schema.TypeMap,
																Computed: true,
																Elem:     schema.TypeString,
															},
														},
													},
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"supporting_resource_id": {
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
										},
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
			"character_set": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_workload_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"listener_port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ncharacter_set": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ons_port_local": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ons_port_remote": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"prefix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"private_endpoint_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"composite_data_guard_shard_spaces": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"chunk_count": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"original_replica": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"databases": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"cloud_autonomous_vm_cluster_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"compute_count": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"data_storage_size_in_gbs": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"ADBD_EXISTING_CLUSTER",
													}, true),
												},

												// Optional
												"admin_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"admin_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"admin_password_secret_version_number": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"fast_start_fail_over_lag_limit_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_auto_scaling_enabled": {
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
												"kms_key_version_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_end_point_group_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_key_store_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"protection_mode": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"shard_space_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"standby_maintenance_buffer_in_days": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vault_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
												"container_database_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"metadata": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"map": {
																Type:     schema.TypeMap,
																Computed: true,
																Elem:     schema.TypeString,
															},
														},
													},
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"supporting_resource_id": {
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
										},
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

						// Optional
						"data_guard_replicas": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"databases": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"cloud_autonomous_vm_cluster_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"compute_count": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"data_storage_size_in_gbs": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"ADBD_EXISTING_CLUSTER",
													}, true),
												},

												// Optional
												"admin_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"admin_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"admin_password_secret_version_number": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"fast_start_fail_over_lag_limit_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_auto_scaling_enabled": {
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
												"kms_key_version_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_end_point_group_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_key_store_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"protection_mode": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"shard_space_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"standby_maintenance_buffer_in_days": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vault_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
												"container_database_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"metadata": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"map": {
																Type:     schema.TypeMap,
																Computed: true,
																Elem:     schema.TypeString,
															},
														},
													},
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"supporting_resource_id": {
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
										},
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
			"composite_raft_shard_spaces": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"chunk_count": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"raft_clusters": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"databases": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"cloud_autonomous_vm_cluster_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"compute_count": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"data_storage_size_in_gbs": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"ADBD_EXISTING_CLUSTER",
													}, true),
												},

												// Optional
												"admin_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"admin_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"admin_password_secret_version_number": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"fast_start_fail_over_lag_limit_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_auto_scaling_enabled": {
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
												"kms_key_version_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_end_point_group_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_key_store_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"protection_mode": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"shard_space_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"standby_maintenance_buffer_in_days": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vault_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
												"container_database_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"metadata": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"map": {
																Type:     schema.TypeMap,
																Computed: true,
																Elem:     schema.TypeString,
															},
														},
													},
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"supporting_resource_id": {
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
										},
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"replication_factor": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"ru_mode": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"witness_count": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"replication_unit_count": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"db_backup_config": {
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
						"backup_destination_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"dbrs_policy_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"internet_proxy": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_remote": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"remote_region": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"recovery_window_in_days": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
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
			"listener_port_tls": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"notification_topic_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"patch_operations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"operation": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"INSERT",
								"MERGE",
								"REMOVE",
							}, true),
						},
						"selection": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     validation.StringIsJSON,
							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
						},

						// Optional

						// Computed
					},
				},
			},
			"system_chunk_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"system_data_guard_databases": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"original_replica": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"databases": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"cloud_autonomous_vm_cluster_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"compute_count": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"data_storage_size_in_gbs": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"ADBD_EXISTING_CLUSTER",
													}, true),
												},

												// Optional
												"admin_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"admin_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"admin_password_secret_version_number": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"fast_start_fail_over_lag_limit_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_auto_scaling_enabled": {
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
												"kms_key_version_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_end_point_group_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_key_store_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"protection_mode": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"shard_space_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"standby_maintenance_buffer_in_days": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vault_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
												"container_database_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"metadata": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"map": {
																Type:     schema.TypeMap,
																Computed: true,
																Elem:     schema.TypeString,
															},
														},
													},
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"supporting_resource_id": {
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
										},
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

						// Optional
						"data_guard_replicas": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"databases": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"cloud_autonomous_vm_cluster_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"compute_count": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"data_storage_size_in_gbs": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"ADBD_EXISTING_CLUSTER",
													}, true),
												},

												// Optional
												"admin_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"admin_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"admin_password_secret_version_number": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"fast_start_fail_over_lag_limit_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_auto_scaling_enabled": {
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
												"kms_key_version_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_end_point_group_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_key_store_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"protection_mode": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"shard_space_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"standby_maintenance_buffer_in_days": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vault_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
												"container_database_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"metadata": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"map": {
																Type:     schema.TypeMap,
																Computed: true,
																Elem:     schema.TypeString,
															},
														},
													},
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"supporting_resource_id": {
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
										},
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
			"system_raft_clusters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"databases": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"cloud_autonomous_vm_cluster_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"compute_count": {
										Type:     schema.TypeFloat,
										Required: true,
										ForceNew: true,
									},
									"data_storage_size_in_gbs": {
										Type:     schema.TypeFloat,
										Required: true,
										ForceNew: true,
									},
									"source": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"ADBD_EXISTING_CLUSTER",
										}, true),
									},

									// Optional
									"admin_password": {
										Type:      schema.TypeString,
										Optional:  true,
										Computed:  true,
										ForceNew:  true,
										Sensitive: true,
									},
									"admin_password_secret_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"admin_password_secret_version_number": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"fast_start_fail_over_lag_limit_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_auto_scaling_enabled": {
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
									"kms_key_version_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"okv_end_point_group_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"okv_key_store_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"protection_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"shard_space_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"standby_maintenance_buffer_in_days": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"vault_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
									"container_database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"metadata": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"map": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
											},
										},
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"supporting_resource_id": {
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
							},
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"replication_factor": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"ru_mode": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"witness_count": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"system_raft_replication_unit_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"effective_system_raft_replication_unit_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"user_shard_spaces": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"original_replica": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"database": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"cloud_autonomous_vm_cluster_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"compute_count": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"data_storage_size_in_gbs": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"ADBD_EXISTING_CLUSTER",
													}, true),
												},

												// Optional
												"admin_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"admin_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"admin_password_secret_version_number": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"fast_start_fail_over_lag_limit_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_auto_scaling_enabled": {
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
												"kms_key_version_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_end_point_group_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_key_store_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"protection_mode": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"shard_space_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"standby_maintenance_buffer_in_days": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vault_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
												"container_database_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"metadata": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"map": {
																Type:     schema.TypeMap,
																Computed: true,
																Elem:     schema.TypeString,
															},
														},
													},
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"supporting_resource_id": {
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
										},
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

						// Optional
						"data_guard_replicas": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"database": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"cloud_autonomous_vm_cluster_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"compute_count": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"data_storage_size_in_gbs": {
													Type:     schema.TypeFloat,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"ADBD_EXISTING_CLUSTER",
													}, true),
												},

												// Optional
												"admin_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"admin_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"admin_password_secret_version_number": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"fast_start_fail_over_lag_limit_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_auto_scaling_enabled": {
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
												"kms_key_version_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_end_point_group_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"okv_key_store_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"protection_mode": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"shard_space_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"standby_maintenance_buffer_in_days": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vault_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
												"container_database_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"metadata": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"map": {
																Type:     schema.TypeMap,
																Computed: true,
																Elem:     schema.TypeString,
															},
														},
													},
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"supporting_resource_id": {
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
										},
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
			"vcn_nsg_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"nsg_ids": {
							Type:     schema.TypeSet,
							Required: true,
							ForceNew: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"vcn_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"add_gds_control_node_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"change_db_backup_config_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"configure_auto_resource_management_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"configure_gsm_wallet_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"configure_sharding_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"generate_wallet_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"move_replication_unit_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"recreate_failed_resource_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_gsms_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"operational_state": operationalStateSchema(),
			"start_database_trigger": {
				Deprecated: "Use operational_state to manage start/stop.",
				Type:       schema.TypeInt,
				Optional:   true,
			},
			"stop_database_trigger": {
				Deprecated: "Use operational_state to manage start/stop.",
				Type:       schema.TypeInt,
				Optional:   true,
			},
			"validate_ca_bundle_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"validate_network_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"auto_resource_management_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"action_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cool_off_period_in_minutes": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"max_move_ru_attempts": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_move_ru_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"notification_topic_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"stream_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"connection_strings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"all_connection_strings": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
			"db_deployment_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gds_control_node_details": {
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
						"gds_control_node_image_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"metadata": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"map": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supporting_resource_id": {
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
				},
			},
			"gsm_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compute_count": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"data_storage_size_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"gsm_image_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"metadata": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"map": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supporting_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_ssl_certificate_expires": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"latest_gsm_image": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"map": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
			"replication_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sharding_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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
	for observed, configured := range map[string]string{
		"observed_composite_data_guard_shard_spaces": "composite_data_guard_shard_spaces",
		"observed_composite_raft_shard_spaces":       "composite_raft_shard_spaces",
		"observed_system_raft_clusters":              "system_raft_clusters",
		"observed_system_data_guard_databases":       "system_data_guard_databases",
		"observed_user_shard_spaces":                 "user_shard_spaces",
	} {
		resource.Schema[observed] = computedOnlyTopologySchema(resource.Schema[configured])
	}
	configureInitialDatabaseTopology(resource)
	configureBackupConfigSchema(resource, true)
	configureMainDatabaseWorkRequestRecoverySchema(resource)
	resource.CustomizeDiff = customdiff.Sequence(validateBackupConfigDiff, validateOperationalStateDiff, resource.CustomizeDiff)
	return resource
}

func computedOnlyTopologySchema(configured *schema.Schema) *schema.Schema {
	observed := *configured
	observed.Required = false
	observed.Optional = false
	observed.Computed = true
	observed.ForceNew = false
	observed.MinItems = 0
	observed.MaxItems = 0
	observed.ConflictsWith = nil
	observed.ExactlyOneOf = nil
	observed.AtLeastOneOf = nil
	observed.RequiredWith = nil
	return &observed
}

func validateDistributedAutonomousDatabaseDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	if err := validateDistributedDatabaseWriteOnceReplicationUnitCountDiff(ctx, d, meta); err != nil {
		return err
	}
	return normalizeDistributedAutonomousDatabaseSystemRaftClustersDiff(ctx, d, meta)
}

func normalizeDistributedAutonomousDatabaseSystemRaftClustersDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	oldValue, newValue := d.GetChange("system_raft_clusters")
	if autonomousSystemRaftClustersEquivalent(oldValue, newValue) {
		return d.Clear("system_raft_clusters")
	}
	return nil
}

// Some GET responses omit computeCount while
// reporting the ECPU count in metadata. Preserve typed values when present and
// only use metadata with matching, known units; never infer an OCPU conversion.
func autonomousDatabaseComputeCount(computeCount *float32, metadata *oci_distributed_database.DistributedAutonomousDbMetadata) (float32, bool) {
	if computeCount != nil {
		return *computeCount, true
	}
	if metadata == nil || !strings.EqualFold(strings.TrimSpace(metadata.PropertiesMap["computeModel"]), "ECPU") {
		return 0, false
	}
	count, err := strconv.ParseFloat(strings.TrimSpace(metadata.PropertiesMap["cpu"]), 32)
	if err != nil || count <= 0 || math.IsNaN(count) || math.IsInf(count, 0) {
		return 0, false
	}
	return float32(count), true
}

func sortAutonomousSystemRaftClusters(clusters []interface{}) {
	sort.SliceStable(clusters, func(i, j int) bool {
		return autonomousSystemRaftClusterMapKey(clusters[i]) < autonomousSystemRaftClusterMapKey(clusters[j])
	})
}

func autonomousSystemRaftClusterMapKey(value interface{}) string {
	cluster, ok := value.(map[string]interface{})
	if !ok {
		return ""
	}
	name, nameOK := cluster["name"].(string)
	ruMode, ruModeOK := cluster["ru_mode"].(string)
	if !nameOK || !ruModeOK {
		return ""
	}
	return strings.ToLower(name) + "\x00" + strings.ToUpper(ruMode)
}

func autonomousSystemRaftClustersEquivalent(oldValue, newValue interface{}) bool {
	return reflect.DeepEqual(normalizeAutonomousSystemRaftClusters(oldValue), normalizeAutonomousSystemRaftClusters(newValue))
}

func normalizeAutonomousSystemRaftClusters(value interface{}) []map[string]interface{} {
	clusters, ok := value.([]interface{})
	if !ok || len(clusters) == 0 {
		return nil
	}
	result := make([]map[string]interface{}, 0, len(clusters))
	for _, raw := range clusters {
		cluster, ok := raw.(map[string]interface{})
		if !ok {
			return nil
		}
		copy := map[string]interface{}{}
		for key, item := range cluster {
			copy[key] = item
		}
		result = append(result, copy)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return autonomousSystemRaftClusterMapKey(result[i]) < autonomousSystemRaftClusterMapKey(result[j])
	})
	return result
}

func createDistributedDatabaseDistributedAutonomousDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if err := validateBackupConfigUpdate(d); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	if _, err := operationalStateTarget(d); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	if err := validateCreateActionTriggers(d, distributedAutonomousDatabaseActionTriggers); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	sync := &DistributedDatabaseDistributedAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	return nil

}

func readDistributedDatabaseDistributedAutonomousDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	if err := sync.resumePendingWorkRequest(ctx); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDistributedDatabaseDistributedAutonomousDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if err := validateBackupConfigUpdate(d); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	target, err := operationalStateTarget(d)
	if err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	ctx, cancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	sync := &DistributedDatabaseDistributedAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	if target == "AVAILABLE" {
		if err := sync.StartDistributedAutonomousDatabase(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("add_gds_control_node_trigger"); ok && sync.D.HasChange("add_gds_control_node_trigger") {
		oldRaw, newRaw := sync.D.GetChange("add_gds_control_node_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.AddDistributedAutonomousDatabaseGdsControlNode(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("add_gds_control_node_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if backupConfigUpdateRequested(sync.D) {
		if err := sync.ChangeDistributedAutonomousDbBackupConfig(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("configure_auto_resource_management_trigger"); ok && sync.D.HasChange("configure_auto_resource_management_trigger") {
		oldRaw, newRaw := sync.D.GetChange("configure_auto_resource_management_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ConfigureDistributedAutonomousDatabaseAutoResourceManagement(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("configure_auto_resource_management_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("configure_gsm_wallet_trigger"); ok && sync.D.HasChange("configure_gsm_wallet_trigger") {
		oldRaw, newRaw := sync.D.GetChange("configure_gsm_wallet_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ConfigureDistributedAutonomousDatabaseGsmWallet(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("configure_gsm_wallet_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("configure_sharding_trigger"); ok && sync.D.HasChange("configure_sharding_trigger") {
		oldRaw, newRaw := sync.D.GetChange("configure_sharding_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ConfigureDistributedAutonomousDatabaseSharding(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("configure_sharding_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("generate_wallet_trigger"); ok && sync.D.HasChange("generate_wallet_trigger") {
		oldRaw, newRaw := sync.D.GetChange("generate_wallet_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.GenerateDistributedAutonomousDatabaseWallet(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("generate_wallet_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("move_replication_unit_trigger"); ok && sync.D.HasChange("move_replication_unit_trigger") {
		oldRaw, newRaw := sync.D.GetChange("move_replication_unit_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.MoveDistributedAutonomousDatabaseReplicationUnit(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("move_replication_unit_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("recreate_failed_resource_trigger"); ok && sync.D.HasChange("recreate_failed_resource_trigger") {
		oldRaw, newRaw := sync.D.GetChange("recreate_failed_resource_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.RecreateFailedDistributedAutonomousDatabaseResource(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("recreate_failed_resource_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("scale_gsms_trigger"); ok && sync.D.HasChange("scale_gsms_trigger") {
		oldRaw, newRaw := sync.D.GetChange("scale_gsms_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ScaleDistributedAutonomousDatabaseGsms(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("scale_gsms_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("validate_ca_bundle_trigger"); ok && sync.D.HasChange("validate_ca_bundle_trigger") {
		oldRaw, newRaw := sync.D.GetChange("validate_ca_bundle_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ValidateDistributedAutonomousDatabaseCaBundle(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("validate_ca_bundle_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("validate_network_trigger"); ok && sync.D.HasChange("validate_network_trigger") {
		oldRaw, newRaw := sync.D.GetChange("validate_network_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ValidateDistributedAutonomousDatabaseNetwork(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("validate_network_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	// A start/stop-only request must not issue a configuration PUT.
	if target == "" || operationalStateConfigurationChanged(d) {
		if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}
	if target == "STOPPED" {
		if err := sync.StopDistributedAutonomousDatabase(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}
	if target != "" {
		return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
	}
	return nil
}

func deleteDistributedDatabaseDistributedAutonomousDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DistributedDatabaseDistributedAutonomousDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_distributed_database.DistributedAutonomousDbServiceClient
	Res                    *oci_distributed_database.DistributedAutonomousDatabase
	PatchResponse          *oci_distributed_database.DistributedAutonomousDatabase
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_distributed_database.DistributedDbWorkRequestServiceClient
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateCreating),
	}
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateInactive),
		string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive),
		string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateNeedsAttention),
	}
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateDeleting),
	}
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateDeleted),
	}
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_distributed_database.CreateDistributedAutonomousDatabaseRequest{}

	if catalogDetails, ok := s.D.GetOkExists("catalog_details"); ok {
		interfaces := catalogDetails.([]interface{})
		tmp := make([]oci_distributed_database.CreateAutonomousCatalogDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "catalog_details", stateDataIndex)
			converted, err := s.mapToCreateAutonomousCatalogDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("catalog_details") {
			request.CatalogDetails = tmp
		}
	}

	if characterSet, ok := s.D.GetOkExists("character_set"); ok {
		tmp := characterSet.(string)
		request.CharacterSet = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compositeDataGuardShardSpaces, ok := s.D.GetOkExists("composite_data_guard_shard_spaces"); ok {
		interfaces := compositeDataGuardShardSpaces.([]interface{})
		tmp := make([]oci_distributed_database.CreateAutonomousCompositeDataGuardShardSpaceDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "composite_data_guard_shard_spaces", stateDataIndex)
			converted, err := s.mapToCreateAutonomousCompositeDataGuardShardSpaceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("composite_data_guard_shard_spaces") {
			request.CompositeDataGuardShardSpaces = tmp
		}
	}

	if compositeRaftShardSpaces, ok := s.D.GetOkExists("composite_raft_shard_spaces"); ok {
		interfaces := compositeRaftShardSpaces.([]interface{})
		tmp := make([]oci_distributed_database.CreateAutonomousCompositeRaftShardSpaceDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "composite_raft_shard_spaces", stateDataIndex)
			converted, err := s.mapToCreateAutonomousCompositeRaftShardSpaceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("composite_raft_shard_spaces") {
			request.CompositeRaftShardSpaces = tmp
		}
	}

	if databaseVersion, ok := s.D.GetOkExists("database_version"); ok {
		tmp := databaseVersion.(string)
		request.DatabaseVersion = &tmp
	}

	if dbBackupConfig, ok := s.D.GetOkExists("db_backup_config"); ok {
		if tmpList := dbBackupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_backup_config", 0)
			tmp, err := s.mapToDistributedAutonomousDbBackupConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DbBackupConfig = &tmp
		}
	}

	if dbWorkloadType, ok := s.D.GetOkExists("db_workload_type"); ok {
		request.DbWorkloadType = oci_distributed_database.CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum(dbWorkloadType.(string))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if listenerPort, ok := s.D.GetOkExists("listener_port"); ok {
		tmp := listenerPort.(int)
		request.ListenerPort = &tmp
	}

	if listenerPortTls, ok := s.D.GetOkExists("listener_port_tls"); ok {
		tmp := listenerPortTls.(int)
		request.ListenerPortTls = &tmp
	}

	if ncharacterSet, ok := s.D.GetOkExists("ncharacter_set"); ok {
		tmp := ncharacterSet.(string)
		request.NcharacterSet = &tmp
	}

	if notificationTopicIds, ok := s.D.GetOkExists("notification_topic_ids"); ok {
		interfaces := notificationTopicIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("notification_topic_ids") {
			request.NotificationTopicIds = tmp
		}
	}

	if onsPortLocal, ok := s.D.GetOkExists("ons_port_local"); ok {
		tmp := onsPortLocal.(int)
		request.OnsPortLocal = &tmp
	}

	if onsPortRemote, ok := s.D.GetOkExists("ons_port_remote"); ok {
		tmp := onsPortRemote.(int)
		request.OnsPortRemote = &tmp
	}

	if prefix, ok := s.D.GetOkExists("prefix"); ok {
		tmp := prefix.(string)
		request.Prefix = &tmp
	}

	if privateEndpointIds, ok := s.D.GetOkExists("private_endpoint_ids"); ok {
		interfaces := privateEndpointIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("private_endpoint_ids") {
			request.PrivateEndpointIds = tmp
		}
	}

	if systemChunkCount, ok := s.D.GetOkExists("system_chunk_count"); ok {
		tmp := systemChunkCount.(int)
		request.SystemChunkCount = &tmp
	}

	if systemDataGuardDatabases, ok := s.D.GetOkExists("system_data_guard_databases"); ok {
		if tmpList := systemDataGuardDatabases.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "system_data_guard_databases", 0)
			tmp, err := s.mapToCreateAutonomousSystemDataGuardDatabaseDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SystemDataGuardDatabases = &tmp
		}
	}

	if systemRaftClusters, ok := s.D.GetOkExists("system_raft_clusters"); ok {
		interfaces := systemRaftClusters.([]interface{})
		tmp := make([]oci_distributed_database.CreateAutonomousSystemRaftClusterDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "system_raft_clusters", stateDataIndex)
			converted, err := s.mapToCreateAutonomousSystemRaftClusterDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("system_raft_clusters") {
			request.SystemRaftClusters = tmp
		}
	}

	if systemRaftReplicationUnitCount, ok := s.D.GetOkExists("system_raft_replication_unit_count"); ok {
		tmp := systemRaftReplicationUnitCount.(int)
		request.SystemRaftReplicationUnitCount = &tmp
	}

	if userShardSpaces, ok := s.D.GetOkExists("user_shard_spaces"); ok {
		interfaces := userShardSpaces.([]interface{})
		tmp := make([]oci_distributed_database.CreateAutonomousUserShardSpaceDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user_shard_spaces", stateDataIndex)
			converted, err := s.mapToCreateAutonomousUserShardSpaceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("user_shard_spaces") {
			request.UserShardSpaces = tmp
		}
	}

	if vcnNsgIds, ok := s.D.GetOkExists("vcn_nsg_ids"); ok {
		interfaces := vcnNsgIds.([]interface{})
		tmp := make([]oci_distributed_database.VcnNsgIdsDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vcn_nsg_ids", stateDataIndex)
			converted, err := s.mapToVcnNsgIdsDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("vcn_nsg_ids") {
			request.VcnNsgIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.CreateDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}

	if workId == nil {
		return fmt.Errorf("CreateDistributedAutonomousDatabase returned no work request ID")
	}
	if err := recordPendingMainDatabaseWorkRequest(s.D, *workId, "create", string(oci_distributed_database.ActionTypeCreated), ""); err != nil {
		return err
	}
	return s.resumePendingWorkRequest(ctx)
}
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) Patch(ctx context.Context) error {
	request := oci_distributed_database.PatchDistributedAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &tmp

	if patchOperations, ok := s.D.GetOkExists("patch_operations"); ok {
		interfaces := patchOperations.([]interface{})
		tmp := make([]oci_distributed_database.PatchInstruction, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patch_operations", stateDataIndex)
			converted, err := s.mapToPatchInstruction(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("patch_operations") {
			request.Items = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")
	response, err := s.Client.PatchDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDistributedAutonomousDatabaseFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) getDistributedAutonomousDatabaseFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_distributed_database.ActionTypeEnum, timeout time.Duration) error {

	if workId == nil {
		return fmt.Errorf("work request ID is required")
	}
	if err := recordPendingMainDatabaseWorkRequest(s.D, *workId, "update", string(actionTypeEnum), ""); err != nil {
		return err
	}
	// Wait until it finishes
	distributedAutonomousDatabaseId, err := distributedAutonomousDatabaseWaitForWorkRequest(ctx, workId, "distributedautonomousdatabase",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*distributedAutonomousDatabaseId)

	if err := s.GetWithContext(ctx); err != nil {
		return err
	}
	return clearPendingMainDatabaseWorkRequest(s.D)
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) resumePendingWorkRequest(ctx context.Context) error {
	pending, ok := pendingMainDatabaseWorkRequestFromState(s.D)
	if !ok {
		return nil
	}
	action := oci_distributed_database.ActionTypeEnum(pending.Action)
	if action == "" {
		return fmt.Errorf("pending work request %q has no action type", pending.ID)
	}
	timeout := s.D.Timeout(schema.TimeoutUpdate)
	if pending.Operation == "create" {
		timeout = s.D.Timeout(schema.TimeoutCreate)
	}
	workID := pending.ID
	id, err := distributedAutonomousDatabaseWaitForWorkRequest(ctx, &workID, "distributedautonomousdatabase", action, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)
	if err != nil {
		return fmt.Errorf("resume %s work request %q: %w", pending.Operation, pending.ID, err)
	}
	if pending.Operation == "create" {
		if id == nil {
			return fmt.Errorf("create work request %q completed without a database identifier", pending.ID)
		}
		s.D.SetId(*id)
		if err := s.GetWithContext(ctx); err != nil {
			return err
		}
	} else if pending.OperationalState != "" {
		if err := waitForOperationalState(ctx, timeout, pending.OperationalState, s.fetchOperationalState); err != nil {
			return err
		}
	} else if err := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); err != nil {
		return err
	}
	return clearPendingMainDatabaseWorkRequest(s.D)
}

func distributedAutonomousDatabaseWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_distributed_database.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_distributed_database.DistributedDbWorkRequestServiceClient) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if deadline, ok := ctx.Deadline(); ok {
		timeout = time.Until(deadline)
	}
	retryPolicy := distributedDatabaseWorkRequestRetryPolicy(timeout, disableFoundRetries)

	response := oci_distributed_database.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_distributed_database.OperationStatusInProgress),
			string(oci_distributed_database.OperationStatusAccepted),
			string(oci_distributed_database.OperationStatusWaiting),
			string(oci_distributed_database.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_distributed_database.OperationStatusSucceeded),
			string(oci_distributed_database.OperationStatusFailed),
			string(oci_distributed_database.OperationStatusCanceled),
			string(oci_distributed_database.OperationStatusNeedsAttention),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_distributed_database.GetWorkRequestRequest{
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
	if _, e := stateConf.WaitForStateContext(ctx); e != nil {
		return nil, e
	}

	// Action resources already know the target OCID and do not require discovery metadata.
	if response.Status == oci_distributed_database.OperationStatusSucceeded && entityType == "" {
		return nil, nil
	}
	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_distributed_database.OperationStatusFailed || response.Status == oci_distributed_database.OperationStatusCanceled || response.Status == oci_distributed_database.OperationStatusNeedsAttention {
		return nil, getErrorFromDistributedDatabaseDistributedAutonomousDatabaseWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDistributedDatabaseDistributedAutonomousDatabaseWorkRequest(ctx context.Context, client *oci_distributed_database.DistributedDbWorkRequestServiceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_distributed_database.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_distributed_database.ListWorkRequestErrorsRequest{
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
		message, code := "", ""
		if wrkErr.Message != nil {
			message = *wrkErr.Message
		}
		if wrkErr.Code != nil {
			code = *wrkErr.Code
		}
		allErrs = append(allErrs, strings.TrimSpace(code+": "+message))
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.GetDistributedAutonomousDatabaseRequest{Metadata: oci_common.String(operationalStateMetadataQuery)}

	tmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.GetDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DistributedAutonomousDatabase
	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_distributed_database.UpdateDistributedAutonomousDatabaseRequest{}

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

	tmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if notificationTopicIds, ok := s.D.GetOkExists("notification_topic_ids"); ok {
		interfaces := notificationTopicIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("notification_topic_ids") {
			request.NotificationTopicIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.UpdateDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DistributedAutonomousDatabase
	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_distributed_database.DeleteDistributedAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.DeleteDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := distributedAutonomousDatabaseWaitForWorkRequest(ctx, workId, "distributedautonomousdatabase",
		oci_distributed_database.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) SetData() error {
	initialTopology := initialDatabaseTopology(s.D)
	if s.Res.AutoResourceManagementConfig != nil {
		s.D.Set("auto_resource_management_config", []interface{}{AutoResourceManagementConfigurationDetailsToMap(s.Res.AutoResourceManagementConfig)})
	} else {
		s.D.Set("auto_resource_management_config", nil)
	}

	lifecycleDetails := ""
	if s.Res.LifecycleDetails != nil {
		lifecycleDetails = *s.Res.LifecycleDetails
	}
	if distributedDatabaseHasRecoverableInfrastructureFailure(string(s.Res.LifecycleState), lifecycleDetails) && s.D.Id() != "" {
		log.Printf("[DEBUG] Preserving configured catalog_details for distributed autonomous database %q during recoverable infrastructure failure", s.D.Id())
	} else {
		catalogDetails := []interface{}{}
		for _, item := range s.Res.CatalogDetails {
			catalogDetails = append(catalogDetails, DistributedAutonomousDatabaseCatalogToMap(item))
		}
		s.D.Set("catalog_details", catalogDetails)
	}

	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	compositeDataGuardShardSpaces := []interface{}{}
	for _, item := range s.Res.CompositeDataGuardShardSpaces {
		compositeDataGuardShardSpaces = append(compositeDataGuardShardSpaces, AutonomousCompositeDataGuardShardSpaceToMap(item))
	}
	s.D.Set("observed_composite_data_guard_shard_spaces", compositeDataGuardShardSpaces)

	compositeRaftShardSpaces := []interface{}{}
	for _, item := range s.Res.CompositeRaftShardSpaces {
		compositeRaftShardSpaces = append(compositeRaftShardSpaces, AutonomousCompositeRaftShardSpaceToMap(item))
	}
	s.D.Set("observed_composite_raft_shard_spaces", compositeRaftShardSpaces)

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{DistributedAutonomousDatabaseConnectionStringToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DbBackupConfig != nil {
		s.D.Set("db_backup_config", []interface{}{DistributedAutonomousDbBackupConfigToMap(s.Res.DbBackupConfig)})
	} else {
		s.D.Set("db_backup_config", nil)
	}

	s.D.Set("db_deployment_type", s.Res.DbDeploymentType)

	s.D.Set("db_workload_type", s.Res.DbWorkloadType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	gdsControlNodeDetails := []interface{}{}
	for _, item := range s.Res.GdsControlNodeDetails {
		gdsControlNodeDetails = append(gdsControlNodeDetails, DistributedAutonomousDatabaseGdsControlNodeToMap(item))
	}
	s.D.Set("gds_control_node_details", gdsControlNodeDetails)

	gsmDetails := []interface{}{}
	for _, item := range s.Res.GsmDetails {
		gsmDetails = append(gsmDetails, DistributedAutonomousDatabaseGsmToMap(item))
	}
	s.D.Set("gsm_details", gsmDetails)

	if s.Res.LatestGsmImage != nil {
		s.D.Set("latest_gsm_image", []interface{}{DistributedAutonomousDatabaseGsmImageToMap(s.Res.LatestGsmImage)})
	} else {
		s.D.Set("latest_gsm_image", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ListenerPort != nil {
		s.D.Set("listener_port", *s.Res.ListenerPort)
	}

	if s.Res.ListenerPortTls != nil {
		s.D.Set("listener_port_tls", *s.Res.ListenerPortTls)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{DistributedAutonomousDbMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

	s.D.Set("notification_topic_ids", s.Res.NotificationTopicIds)

	if s.Res.OnsPortLocal != nil {
		s.D.Set("ons_port_local", *s.Res.OnsPortLocal)
	}

	if s.Res.OnsPortRemote != nil {
		s.D.Set("ons_port_remote", *s.Res.OnsPortRemote)
	}

	if s.Res.Prefix != nil {
		s.D.Set("prefix", *s.Res.Prefix)
	}

	s.D.Set("private_endpoint_ids", s.Res.PrivateEndpointIds)

	s.D.Set("replication_method", s.Res.ReplicationMethod)

	s.D.Set("sharding_method", s.Res.ShardingMethod)

	s.D.Set("state", s.Res.LifecycleState)
	if err := preserveOperationalState(s.D); err != nil {
		return err
	}

	if s.Res.SystemChunkCount != nil {
		s.D.Set("system_chunk_count", *s.Res.SystemChunkCount)
	}

	if s.Res.SystemDataGuardDatabases != nil {
		s.D.Set("observed_system_data_guard_databases", []interface{}{AutonomousSystemDataGuardDatabaseToMap(s.Res.SystemDataGuardDatabases)})
	} else {
		s.D.Set("observed_system_data_guard_databases", nil)
	}

	systemRaftClusters := []interface{}{}
	for _, item := range s.Res.SystemRaftClusters {
		systemRaftClusters = append(systemRaftClusters, AutonomousSystemRaftClusterToMap(item))
	}
	sortAutonomousSystemRaftClusters(systemRaftClusters)
	s.D.Set("observed_system_raft_clusters", systemRaftClusters)

	if err := setDistributedDatabaseWriteOnceReplicationUnitCount(s.D, s.Res.SystemRaftReplicationUnitCount); err != nil {
		return err
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	userShardSpaces := []interface{}{}
	for _, item := range s.Res.UserShardSpaces {
		userShardSpaces = append(userShardSpaces, AutonomousUserShardSpaceToMap(item))
	}
	s.D.Set("observed_user_shard_spaces", userShardSpaces)

	vcnNsgIds := []interface{}{}
	for _, item := range s.Res.VcnNsgIds {
		vcnNsgIds = append(vcnNsgIds, VcnNsgIdsDetailsToMap(item, false))
	}
	s.D.Set("vcn_nsg_ids", vcnNsgIds)

	return preserveInitialDatabaseTopology(s.D, initialTopology, true)
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) AddDistributedAutonomousDatabaseGdsControlNode(ctx context.Context) error {
	request := oci_distributed_database.AddDistributedAutonomousDatabaseGdsControlNodeRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	if publicSshKey, ok := s.D.GetOkExists("public_ssh_key"); ok {
		tmp := publicSshKey.(string)
		request.PublicSshKey = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.AddDistributedAutonomousDatabaseGdsControlNode(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = distributedAutonomousDatabaseWaitForWorkRequest(ctx, workRequestId, "distributedautonomousdatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("add_gds_control_node_trigger")
	s.D.Set("add_gds_control_node_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ChangeDistributedAutonomousDbBackupConfig(ctx context.Context) error {
	request := oci_distributed_database.ChangeDistributedAutonomousDbBackupConfigRequest{}

	if err := decodeBackupConfig(s.D.GetRawConfig(), &request.DbBackupConfig); err != nil {
		return fmt.Errorf("ChangeDistributedAutonomousDbBackupConfig: %w", err)
	}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ChangeDistributedAutonomousDbBackupConfig(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		if err := recordPendingMainDatabaseWorkRequest(s.D, *workRequestId, "backup_config", string(oci_distributed_database.ActionTypeUpdated), ""); err != nil {
			return err
		}
		if err := s.resumePendingWorkRequest(ctx); err != nil {
			return err
		}
		val := s.D.Get("change_db_backup_config_trigger")
		s.D.Set("change_db_backup_config_trigger", val)
		return nil
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("change_db_backup_config_trigger")
	s.D.Set("change_db_backup_config_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ConfigureDistributedAutonomousDatabaseAutoResourceManagement(ctx context.Context) error {
	request := oci_distributed_database.ConfigureDistributedAutonomousDatabaseAutoResourceManagementRequest{}

	if actionType, ok := s.D.GetOkExists("action_type"); ok {
		request.ActionType = oci_distributed_database.ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum(actionType.(string))
	}

	if coolOffPeriodInMinutes, ok := s.D.GetOkExists("cool_off_period_in_minutes"); ok {
		tmp := coolOffPeriodInMinutes.(int)
		request.CoolOffPeriodInMinutes = &tmp
	}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if maxMoveRuAttempts, ok := s.D.GetOkExists("max_move_ru_attempts"); ok {
		tmp := maxMoveRuAttempts.(int)
		request.MaxMoveRuAttempts = &tmp
	}

	if maxMoveRuUnit, ok := s.D.GetOkExists("max_move_ru_unit"); ok {
		request.MaxMoveRuUnit = oci_distributed_database.ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum(maxMoveRuUnit.(string))
	}

	if mode, ok := s.D.GetOkExists("mode"); ok {
		request.Mode = oci_distributed_database.ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum(mode.(string))
	}

	if notificationTopicIds, ok := s.D.GetOkExists("notification_topic_ids"); ok {
		interfaces := notificationTopicIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("notification_topic_ids") {
			request.NotificationTopicIds = tmp
		}
	}

	if streamIds, ok := s.D.GetOkExists("stream_ids"); ok {
		interfaces := streamIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("stream_ids") {
			request.StreamIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ConfigureDistributedAutonomousDatabaseAutoResourceManagement(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = distributedAutonomousDatabaseWaitForWorkRequest(ctx, workRequestId, "distributedautonomousdatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("configure_auto_resource_management_trigger")
	s.D.Set("configure_auto_resource_management_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ConfigureDistributedAutonomousDatabaseGsmWallet(ctx context.Context) error {
	request := oci_distributed_database.ConfigureDistributedAutonomousDatabaseGsmWalletRequest{}

	if caBundleId, ok := s.D.GetOkExists("ca_bundle_id"); ok {
		tmp := caBundleId.(string)
		request.CaBundleId = &tmp
	}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
	}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ConfigureDistributedAutonomousDatabaseGsmWallet(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = distributedAutonomousDatabaseWaitForWorkRequest(ctx, workRequestId, "distributedautonomousdatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("configure_gsm_wallet_trigger")
	s.D.Set("configure_gsm_wallet_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ConfigureDistributedAutonomousDatabaseSharding(ctx context.Context) error {
	request := oci_distributed_database.ConfigureDistributedAutonomousDatabaseShardingRequest{}

	if caBundleId, ok := s.D.GetOkExists("ca_bundle_id"); ok {
		tmp := caBundleId.(string)
		request.CaBundleId = &tmp
	}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
	}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	if isRebalanceRequired, ok := s.D.GetOkExists("is_rebalance_required"); ok {
		tmp := isRebalanceRequired.(bool)
		request.IsRebalanceRequired = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ConfigureDistributedAutonomousDatabaseSharding(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = distributedAutonomousDatabaseWaitForWorkRequest(ctx, workRequestId, "distributedautonomousdatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("configure_sharding_trigger")
	s.D.Set("configure_sharding_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) GenerateDistributedAutonomousDatabaseWallet(ctx context.Context) error {
	request := oci_distributed_database.GenerateDistributedAutonomousDatabaseWalletRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.GenerateDistributedAutonomousDatabaseWallet(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = distributedAutonomousDatabaseWaitForWorkRequest(ctx, workRequestId, "distributedautonomousdatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("generate_wallet_trigger")
	s.D.Set("generate_wallet_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) MoveDistributedAutonomousDatabaseReplicationUnit(ctx context.Context) error {
	request := oci_distributed_database.MoveDistributedAutonomousDatabaseReplicationUnitRequest{}

	if destinationShardName, ok := s.D.GetOkExists("destination_shard_name"); ok {
		tmp := destinationShardName.(string)
		request.DestinationShardName = &tmp
	}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	if replicationUnits, ok := s.D.GetOkExists("replication_units"); ok {
		interfaces := replicationUnits.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("replication_units") {
			request.ReplicationUnits = tmp
		}
	}

	if sourceShardName, ok := s.D.GetOkExists("source_shard_name"); ok {
		tmp := sourceShardName.(string)
		request.SourceShardName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.MoveDistributedAutonomousDatabaseReplicationUnit(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = distributedAutonomousDatabaseWaitForWorkRequest(ctx, workRequestId, "distributedautonomousdatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("move_replication_unit_trigger")
	s.D.Set("move_replication_unit_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) RecreateFailedDistributedAutonomousDatabaseResource(ctx context.Context) error {
	request := oci_distributed_database.RecreateFailedDistributedAutonomousDatabaseResourceRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.RecreateFailedDistributedAutonomousDatabaseResource(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = distributedAutonomousDatabaseWaitForWorkRequest(ctx, workRequestId, "distributedautonomousdatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("recreate_failed_resource_trigger")
	s.D.Set("recreate_failed_resource_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ScaleDistributedAutonomousDatabaseGsms(ctx context.Context) error {
	request := oci_distributed_database.ScaleDistributedAutonomousDatabaseGsmsRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ScaleDistributedAutonomousDatabaseGsms(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = distributedAutonomousDatabaseWaitForWorkRequest(ctx, workRequestId, "distributedautonomousdatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("scale_gsms_trigger")
	s.D.Set("scale_gsms_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) StartDistributedAutonomousDatabase(ctx context.Context) error {
	request := oci_distributed_database.StartDistributedAutonomousDatabaseRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.StartDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		if err := recordPendingMainDatabaseWorkRequest(s.D, *workRequestId, "operational_state", string(oci_distributed_database.ActionTypeUpdated), "AVAILABLE"); err != nil {
			return err
		}
		if err := s.resumePendingWorkRequest(ctx); err != nil {
			return err
		}
		return s.D.Set("operational_state", "AVAILABLE")
	}

	if err := waitForOperationalState(ctx, s.D.Timeout(schema.TimeoutUpdate), "AVAILABLE", s.fetchOperationalState); err != nil {
		return err
	}
	return s.D.Set("operational_state", "AVAILABLE")
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) StopDistributedAutonomousDatabase(ctx context.Context) error {
	request := oci_distributed_database.StopDistributedAutonomousDatabaseRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.StopDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		if err := recordPendingMainDatabaseWorkRequest(s.D, *workRequestId, "operational_state", string(oci_distributed_database.ActionTypeUpdated), "STOPPED"); err != nil {
			return err
		}
		if err := s.resumePendingWorkRequest(ctx); err != nil {
			return err
		}
		return s.D.Set("operational_state", "STOPPED")
	}

	if err := waitForOperationalState(ctx, s.D.Timeout(schema.TimeoutUpdate), "STOPPED", s.fetchOperationalState); err != nil {
		return err
	}
	return s.D.Set("operational_state", "STOPPED")
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ValidateDistributedAutonomousDatabaseCaBundle(ctx context.Context) error {
	request := oci_distributed_database.ValidateDistributedAutonomousDatabaseCaBundleRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ValidateDistributedAutonomousDatabaseCaBundle(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = distributedAutonomousDatabaseWaitForWorkRequest(ctx, workRequestId, "distributedautonomousdatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("validate_ca_bundle_trigger")
	s.D.Set("validate_ca_bundle_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ValidateDistributedAutonomousDatabaseNetwork(ctx context.Context) error {
	request := oci_distributed_database.ValidateDistributedAutonomousDatabaseNetworkRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	if isSurrogate, ok := s.D.GetOkExists("is_surrogate"); ok {
		tmp := isSurrogate.(bool)
		request.IsSurrogate = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ValidateDistributedAutonomousDatabaseNetwork(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = distributedAutonomousDatabaseWaitForWorkRequest(ctx, workRequestId, "distributedautonomousdatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("validate_network_trigger")
	s.D.Set("validate_network_trigger", val)

	return nil
}

func AutoResourceManagementConfigurationDetailsToMap(obj *oci_distributed_database.AutoResourceManagementConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["action_type"] = string(obj.ActionType)

	if obj.CoolOffPeriodInMinutes != nil {
		result["cool_off_period_in_minutes"] = int(*obj.CoolOffPeriodInMinutes)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.MaxMoveRuAttempts != nil {
		result["max_move_ru_attempts"] = int(*obj.MaxMoveRuAttempts)
	}

	result["max_move_ru_unit"] = string(obj.MaxMoveRuUnit)

	result["mode"] = string(obj.Mode)

	result["notification_topic_ids"] = obj.NotificationTopicIds

	result["stream_ids"] = obj.StreamIds

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousCatalogDatabaseDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousCatalogDatabaseDetails, error) {
	var baseObject oci_distributed_database.CreateAutonomousCatalogDatabaseDetails
	//discriminator
	sourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source"))
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("ADBD_EXISTING_CLUSTER"):
		details := oci_distributed_database.CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails{}
		if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
		}
		if adminPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password_secret_id")); ok {
			tmp := adminPasswordSecretId.(string)
			details.AdminPasswordSecretId = &tmp
		}
		if adminPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password_secret_version_number")); ok {
			tmp := adminPasswordSecretVersionNumber.(int)
			details.AdminPasswordSecretVersionNumber = &tmp
		}
		if cloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_autonomous_vm_cluster_id")); ok {
			tmp := cloudAutonomousVmClusterId.(string)
			details.CloudAutonomousVmClusterId = &tmp
		}
		if computeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_count")); ok {
			tmp := float32(computeCount.(float64))
			details.ComputeCount = &tmp
		}
		if dataStorageSizeInGbs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_gbs")); ok {
			tmp := dataStorageSizeInGbs.(float64)
			details.DataStorageSizeInGbs = &tmp
		}
		if fastStartFailOverLagLimitInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fast_start_fail_over_lag_limit_in_seconds")); ok {
			tmp := fastStartFailOverLagLimitInSeconds.(int)
			details.FastStartFailOverLagLimitInSeconds = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_scaling_enabled")); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if okvEndPointGroupName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "okv_end_point_group_name")); ok {
			tmp := okvEndPointGroupName.(string)
			details.OkvEndPointGroupName = &tmp
		}
		if okvKeyStoreId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "okv_key_store_id")); ok {
			tmp := okvKeyStoreId.(string)
			details.OkvKeyStoreId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedAutonomousDbProtectionModeEnum(protectionMode.(string))
		}
		if standbyMaintenanceBufferInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "standby_maintenance_buffer_in_days")); ok {
			tmp := standbyMaintenanceBufferInDays.(int)
			details.StandbyMaintenanceBufferInDays = &tmp
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source '%v' was specified", source)
	}
	return baseObject, nil
}

func DistributedAutonomousDatabaseCatalogDatabaseToMap(obj *oci_distributed_database.DistributedAutonomousDatabaseCatalogDatabase) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_distributed_database.DistributedAutonomousDatabaseCatalogWithDedicatedInfra:
		// preserve response status and database IDs in all topology readbacks.
		result["source"] = "ADBD_EXISTING_CLUSTER"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

		if v.CloudAutonomousVmClusterId != nil {
			result["cloud_autonomous_vm_cluster_id"] = string(*v.CloudAutonomousVmClusterId)
		}

		if count, ok := autonomousDatabaseComputeCount(v.ComputeCount, v.Metadata); ok {
			result["compute_count"] = count
		}

		if v.DataStorageSizeInGbs != nil {
			result["data_storage_size_in_gbs"] = float64(*v.DataStorageSizeInGbs)
		}

		if v.IsAutoScalingEnabled != nil {
			result["is_auto_scaling_enabled"] = bool(*v.IsAutoScalingEnabled)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		if v.OkvEndPointGroupName != nil {
			result["okv_end_point_group_name"] = string(*v.OkvEndPointGroupName)
		}

		if v.OkvKeyStoreId != nil {
			result["okv_key_store_id"] = string(*v.OkvKeyStoreId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		if v.StandbyMaintenanceBufferInDays != nil {
			result["standby_maintenance_buffer_in_days"] = int(*v.StandbyMaintenanceBufferInDays)
		}

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedAutonomousDbMetadataToMap(v.Metadata)}
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}
	default:
		log.Printf("[WARN] Received 'source' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousCatalogDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousCatalogDetails, error) {
	result := oci_distributed_database.CreateAutonomousCatalogDetails{}

	if dataGuardReplicas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")); ok {
		interfaces := dataGuardReplicas.([]interface{})
		tmp := make([]oci_distributed_database.CreateAutonomousCatalogReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data_guard_replicas"), stateDataIndex)
			converted, err := s.mapToCreateAutonomousCatalogReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")) {
			result.DataGuardReplicas = tmp
		}
	}

	if originalReplica, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "original_replica")); ok {
		if tmpList := originalReplica.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "original_replica"), 0)
			tmp, err := s.mapToCreateAutonomousCatalogReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert original_replica, encountered error: %v", err)
			}
			result.OriginalReplica = &tmp
		}
	}

	return result, nil
}

func DistributedAutonomousDatabaseCatalogToMap(obj oci_distributed_database.DistributedAutonomousDatabaseCatalog) map[string]interface{} {
	result := map[string]interface{}{}

	dataGuardReplicas := []interface{}{}
	for _, item := range obj.DataGuardReplicas {
		dataGuardReplicas = append(dataGuardReplicas, DistributedAutonomousDatabaseCatalogReplicaToMap(item))
	}
	result["data_guard_replicas"] = dataGuardReplicas

	if obj.OriginalReplica != nil {
		result["original_replica"] = []interface{}{DistributedAutonomousDatabaseCatalogReplicaToMap(*obj.OriginalReplica)}
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousCatalogReplicaDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousCatalogReplicaDetails, error) {
	result := oci_distributed_database.CreateAutonomousCatalogReplicaDetails{}

	if database, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database")); ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database"), 0)
			tmp, err := s.mapToCreateAutonomousCatalogDatabaseDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert database, encountered error: %v", err)
			}
			result.Database = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func DistributedAutonomousDatabaseCatalogReplicaToMap(obj oci_distributed_database.DistributedAutonomousDatabaseCatalogReplica) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Database != nil {
		databaseArray := []interface{}{}
		if databaseMap := DistributedAutonomousDatabaseCatalogDatabaseToMap(&obj.Database); databaseMap != nil {
			databaseArray = append(databaseArray, databaseMap)
		}
		result["database"] = databaseArray
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousCompositeDataGuardShardSpaceDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousCompositeDataGuardShardSpaceDetails, error) {
	result := oci_distributed_database.CreateAutonomousCompositeDataGuardShardSpaceDetails{}

	if chunkCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "chunk_count")); ok {
		tmp := chunkCount.(int)
		result.ChunkCount = &tmp
	}

	if dataGuardReplicas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")); ok {
		interfaces := dataGuardReplicas.([]interface{})
		tmp := make([]oci_distributed_database.CreateAutonomousCompositeDataGuardShardSpaceReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data_guard_replicas"), stateDataIndex)
			converted, err := s.mapToCreateAutonomousCompositeDataGuardShardSpaceReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")) {
			result.DataGuardReplicas = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if originalReplica, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "original_replica")); ok {
		if tmpList := originalReplica.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "original_replica"), 0)
			tmp, err := s.mapToCreateAutonomousCompositeDataGuardShardSpaceReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert original_replica, encountered error: %v", err)
			}
			result.OriginalReplica = &tmp
		}
	}

	return result, nil
}

func AutonomousCompositeDataGuardShardSpaceToMap(obj oci_distributed_database.AutonomousCompositeDataGuardShardSpace) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ChunkCount != nil {
		result["chunk_count"] = int(*obj.ChunkCount)
	}

	dataGuardReplicas := []interface{}{}
	for _, item := range obj.DataGuardReplicas {
		dataGuardReplicas = append(dataGuardReplicas, AutonomousCompositeDataGuardShardSpaceReplicaToMap(item))
	}
	result["data_guard_replicas"] = dataGuardReplicas

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.OriginalReplica != nil {
		result["original_replica"] = []interface{}{AutonomousCompositeDataGuardShardSpaceReplicaToMap(*obj.OriginalReplica)}
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousCompositeDataGuardShardSpaceReplicaDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousCompositeDataGuardShardSpaceReplicaDetails, error) {
	result := oci_distributed_database.CreateAutonomousCompositeDataGuardShardSpaceReplicaDetails{}

	if databases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "databases")); ok {
		interfaces := databases.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedAutonomousDatabaseShardDatabaseDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "databases"), stateDataIndex)
			converted, err := s.mapToCreateDistributedAutonomousDatabaseShardDatabaseDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "databases")) {
			result.Databases = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func AutonomousCompositeDataGuardShardSpaceReplicaToMap(obj oci_distributed_database.AutonomousCompositeDataGuardShardSpaceReplica) map[string]interface{} {
	result := map[string]interface{}{}

	databases := []interface{}{}
	for _, item := range obj.Databases {
		databases = append(databases, DistributedAutonomousDatabaseShardDatabaseToMap(item))
	}
	result["databases"] = databases

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousCompositeRaftShardSpaceDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousCompositeRaftShardSpaceDetails, error) {
	result := oci_distributed_database.CreateAutonomousCompositeRaftShardSpaceDetails{}

	if chunkCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "chunk_count")); ok {
		tmp := chunkCount.(int)
		result.ChunkCount = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if raftClusters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "raft_clusters")); ok {
		interfaces := raftClusters.([]interface{})
		tmp := make([]oci_distributed_database.CreateAutonomousCompositeShardSpaceRaftClusterDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "raft_clusters"), stateDataIndex)
			converted, err := s.mapToCreateAutonomousCompositeShardSpaceRaftClusterDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "raft_clusters")) {
			result.RaftClusters = tmp
		}
	}

	if replicationUnitCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replication_unit_count")); ok {
		tmp := replicationUnitCount.(int)
		result.ReplicationUnitCount = &tmp
	}

	return result, nil
}

func AutonomousCompositeRaftShardSpaceToMap(obj oci_distributed_database.AutonomousCompositeRaftShardSpace) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ChunkCount != nil {
		result["chunk_count"] = int(*obj.ChunkCount)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	raftClusters := []interface{}{}
	for _, item := range obj.RaftClusters {
		raftClusters = append(raftClusters, AutonomousCompositeShardSpaceRaftClusterToMap(item))
	}
	result["raft_clusters"] = raftClusters

	if obj.ReplicationUnitCount != nil {
		result["replication_unit_count"] = int(*obj.ReplicationUnitCount)
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousCompositeShardSpaceRaftClusterDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousCompositeShardSpaceRaftClusterDetails, error) {
	result := oci_distributed_database.CreateAutonomousCompositeShardSpaceRaftClusterDetails{}

	if databases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "databases")); ok {
		interfaces := databases.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedAutonomousDatabaseShardDatabaseDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "databases"), stateDataIndex)
			converted, err := s.mapToCreateDistributedAutonomousDatabaseShardDatabaseDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "databases")) {
			result.Databases = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if replicationFactor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replication_factor")); ok {
		tmp := replicationFactor.(int)
		result.ReplicationFactor = &tmp
	}

	if ruMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ru_mode")); ok {
		result.RuMode = oci_distributed_database.CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum(ruMode.(string))
	}

	if witnessCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "witness_count")); ok {
		tmp := witnessCount.(int)
		result.WitnessCount = &tmp
	}

	return result, nil
}

func AutonomousCompositeShardSpaceRaftClusterToMap(obj oci_distributed_database.AutonomousCompositeShardSpaceRaftCluster) map[string]interface{} {
	result := map[string]interface{}{}

	databases := []interface{}{}
	for _, item := range obj.Databases {
		databases = append(databases, DistributedAutonomousDatabaseShardDatabaseToMap(item))
	}
	result["databases"] = databases

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ReplicationFactor != nil {
		result["replication_factor"] = int(*obj.ReplicationFactor)
	}

	result["ru_mode"] = string(obj.RuMode)

	if obj.WitnessCount != nil {
		result["witness_count"] = int(*obj.WitnessCount)
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousSystemDataGuardDatabaseDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousSystemDataGuardDatabaseDetails, error) {
	result := oci_distributed_database.CreateAutonomousSystemDataGuardDatabaseDetails{}

	if dataGuardReplicas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")); ok {
		interfaces := dataGuardReplicas.([]interface{})
		tmp := make([]oci_distributed_database.CreateAutonomousSystemDataGuardReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data_guard_replicas"), stateDataIndex)
			converted, err := s.mapToCreateAutonomousSystemDataGuardReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")) {
			result.DataGuardReplicas = tmp
		}
	}

	if originalReplica, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "original_replica")); ok {
		if tmpList := originalReplica.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "original_replica"), 0)
			tmp, err := s.mapToCreateAutonomousSystemDataGuardReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert original_replica, encountered error: %v", err)
			}
			result.OriginalReplica = &tmp
		}
	}

	return result, nil
}

func AutonomousSystemDataGuardDatabaseToMap(obj *oci_distributed_database.AutonomousSystemDataGuardDatabase) map[string]interface{} {
	result := map[string]interface{}{}

	dataGuardReplicas := []interface{}{}
	for _, item := range obj.DataGuardReplicas {
		dataGuardReplicas = append(dataGuardReplicas, AutonomousSystemDataGuardReplicaToMap(item))
	}
	result["data_guard_replicas"] = dataGuardReplicas

	if obj.OriginalReplica != nil {
		result["original_replica"] = []interface{}{AutonomousSystemDataGuardReplicaToMap(*obj.OriginalReplica)}
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousSystemDataGuardReplicaDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousSystemDataGuardReplicaDetails, error) {
	result := oci_distributed_database.CreateAutonomousSystemDataGuardReplicaDetails{}

	if databases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "databases")); ok {
		interfaces := databases.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedAutonomousDatabaseShardDatabaseDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "databases"), stateDataIndex)
			converted, err := s.mapToCreateDistributedAutonomousDatabaseShardDatabaseDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "databases")) {
			result.Databases = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func AutonomousSystemDataGuardReplicaToMap(obj oci_distributed_database.AutonomousSystemDataGuardReplica) map[string]interface{} {
	result := map[string]interface{}{}

	databases := []interface{}{}
	for _, item := range obj.Databases {
		databases = append(databases, DistributedAutonomousDatabaseShardDatabaseToMap(item))
	}
	result["databases"] = databases

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousSystemRaftClusterDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousSystemRaftClusterDetails, error) {
	result := oci_distributed_database.CreateAutonomousSystemRaftClusterDetails{}

	if databases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "databases")); ok {
		interfaces := databases.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedAutonomousDatabaseShardDatabaseDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "databases"), stateDataIndex)
			converted, err := s.mapToCreateDistributedAutonomousDatabaseShardDatabaseDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "databases")) {
			result.Databases = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if replicationFactor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replication_factor")); ok {
		tmp := replicationFactor.(int)
		result.ReplicationFactor = &tmp
	}

	if ruMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ru_mode")); ok {
		result.RuMode = oci_distributed_database.CreateAutonomousSystemRaftClusterDetailsRuModeEnum(ruMode.(string))
	}

	if witnessCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "witness_count")); ok {
		tmp := witnessCount.(int)
		result.WitnessCount = &tmp
	}

	return result, nil
}

func AutonomousSystemRaftClusterToMap(obj oci_distributed_database.AutonomousSystemRaftCluster) map[string]interface{} {
	result := map[string]interface{}{}

	databases := []interface{}{}
	for _, item := range obj.Databases {
		databases = append(databases, DistributedAutonomousDatabaseShardDatabaseToMap(item))
	}
	result["databases"] = databases

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ReplicationFactor != nil {
		result["replication_factor"] = int(*obj.ReplicationFactor)
	}

	result["ru_mode"] = string(obj.RuMode)

	if obj.WitnessCount != nil {
		result["witness_count"] = int(*obj.WitnessCount)
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousUserShardSpaceDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousUserShardSpaceDetails, error) {
	result := oci_distributed_database.CreateAutonomousUserShardSpaceDetails{}

	if dataGuardReplicas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")); ok {
		interfaces := dataGuardReplicas.([]interface{})
		tmp := make([]oci_distributed_database.CreateAutonomousUserShardSpaceReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data_guard_replicas"), stateDataIndex)
			converted, err := s.mapToCreateAutonomousUserShardSpaceReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")) {
			result.DataGuardReplicas = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if originalReplica, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "original_replica")); ok {
		if tmpList := originalReplica.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "original_replica"), 0)
			tmp, err := s.mapToCreateAutonomousUserShardSpaceReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert original_replica, encountered error: %v", err)
			}
			result.OriginalReplica = &tmp
		}
	}

	return result, nil
}

func AutonomousUserShardSpaceToMap(obj oci_distributed_database.AutonomousUserShardSpace) map[string]interface{} {
	result := map[string]interface{}{}

	dataGuardReplicas := []interface{}{}
	for _, item := range obj.DataGuardReplicas {
		dataGuardReplicas = append(dataGuardReplicas, AutonomousUserShardSpaceReplicaToMap(item))
	}
	result["data_guard_replicas"] = dataGuardReplicas

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.OriginalReplica != nil {
		result["original_replica"] = []interface{}{AutonomousUserShardSpaceReplicaToMap(*obj.OriginalReplica)}
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateAutonomousUserShardSpaceReplicaDetails(fieldKeyFormat string) (oci_distributed_database.CreateAutonomousUserShardSpaceReplicaDetails, error) {
	result := oci_distributed_database.CreateAutonomousUserShardSpaceReplicaDetails{}

	if database, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database")); ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database"), 0)
			tmp, err := s.mapToCreateDistributedAutonomousDatabaseShardDatabaseDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert database, encountered error: %v", err)
			}
			result.Database = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func AutonomousUserShardSpaceReplicaToMap(obj oci_distributed_database.AutonomousUserShardSpaceReplica) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Database != nil {
		databaseArray := []interface{}{}
		if databaseMap := DistributedAutonomousDatabaseShardDatabaseToMap(obj.Database); databaseMap != nil {
			databaseArray = append(databaseArray, databaseMap)
		}
		result["database"] = databaseArray
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateDistributedAutonomousDatabaseShardDatabaseDetails(fieldKeyFormat string) (oci_distributed_database.CreateDistributedAutonomousDatabaseShardDatabaseDetails, error) {
	var baseObject oci_distributed_database.CreateDistributedAutonomousDatabaseShardDatabaseDetails
	//discriminator
	sourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source"))
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("ADBD_EXISTING_CLUSTER"):
		details := oci_distributed_database.CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails{}
		if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
		}
		if adminPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password_secret_id")); ok {
			tmp := adminPasswordSecretId.(string)
			details.AdminPasswordSecretId = &tmp
		}
		if adminPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password_secret_version_number")); ok {
			tmp := adminPasswordSecretVersionNumber.(int)
			details.AdminPasswordSecretVersionNumber = &tmp
		}
		if cloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_autonomous_vm_cluster_id")); ok {
			tmp := cloudAutonomousVmClusterId.(string)
			details.CloudAutonomousVmClusterId = &tmp
		}
		if computeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_count")); ok {
			tmp := float32(computeCount.(float64))
			details.ComputeCount = &tmp
		}
		if dataStorageSizeInGbs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_gbs")); ok {
			tmp := dataStorageSizeInGbs.(float64)
			details.DataStorageSizeInGbs = &tmp
		}
		if fastStartFailOverLagLimitInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fast_start_fail_over_lag_limit_in_seconds")); ok {
			tmp := fastStartFailOverLagLimitInSeconds.(int)
			details.FastStartFailOverLagLimitInSeconds = &tmp
		}
		if isAutoScalingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_scaling_enabled")); ok {
			tmp := isAutoScalingEnabled.(bool)
			details.IsAutoScalingEnabled = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if okvEndPointGroupName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "okv_end_point_group_name")); ok {
			tmp := okvEndPointGroupName.(string)
			details.OkvEndPointGroupName = &tmp
		}
		if okvKeyStoreId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "okv_key_store_id")); ok {
			tmp := okvKeyStoreId.(string)
			details.OkvKeyStoreId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedAutonomousDbProtectionModeEnum(protectionMode.(string))
		}
		if shardSpaceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shard_space_name")); ok {
			tmp := shardSpaceName.(string)
			details.ShardSpaceName = &tmp
		}
		if standbyMaintenanceBufferInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "standby_maintenance_buffer_in_days")); ok {
			tmp := standbyMaintenanceBufferInDays.(int)
			details.StandbyMaintenanceBufferInDays = &tmp
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source '%v' was specified", source)
	}
	return baseObject, nil
}

func DistributedAutonomousDatabaseShardDatabaseToMap(obj oci_distributed_database.DistributedAutonomousDatabaseShardDatabase) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := obj.(type) {
	case oci_distributed_database.DistributedAutonomousDatabaseShardWithDedicatedInfra:
		// preserve response status and database IDs in all topology readbacks.
		result["source"] = "ADBD_EXISTING_CLUSTER"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

		if v.CloudAutonomousVmClusterId != nil {
			result["cloud_autonomous_vm_cluster_id"] = string(*v.CloudAutonomousVmClusterId)
		}

		if count, ok := autonomousDatabaseComputeCount(v.ComputeCount, v.Metadata); ok {
			result["compute_count"] = count
		}

		if v.DataStorageSizeInGbs != nil {
			result["data_storage_size_in_gbs"] = float64(*v.DataStorageSizeInGbs)
		}

		if v.FastStartFailOverLagLimitInSeconds != nil {
			result["fast_start_fail_over_lag_limit_in_seconds"] = int(*v.FastStartFailOverLagLimitInSeconds)
		}

		if v.IsAutoScalingEnabled != nil {
			result["is_auto_scaling_enabled"] = bool(*v.IsAutoScalingEnabled)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		if v.OkvEndPointGroupName != nil {
			result["okv_end_point_group_name"] = string(*v.OkvEndPointGroupName)
		}

		if v.OkvKeyStoreId != nil {
			result["okv_key_store_id"] = string(*v.OkvKeyStoreId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		if v.StandbyMaintenanceBufferInDays != nil {
			result["standby_maintenance_buffer_in_days"] = int(*v.StandbyMaintenanceBufferInDays)
		}

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedAutonomousDbMetadataToMap(v.Metadata)}
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}
	default:
		log.Printf("[WARN] Received 'source' of unknown type %v", obj)
		return nil
	}

	return result
}

func DistributedAutonomousDatabaseConnectionStringToMap(obj *oci_distributed_database.DistributedAutonomousDatabaseConnectionString) map[string]interface{} {
	result := map[string]interface{}{}

	result["all_connection_strings"] = obj.AllConnectionStrings

	return result
}

func DistributedAutonomousDatabaseGdsControlNodeToMap(obj oci_distributed_database.DistributedAutonomousDatabaseGdsControlNode) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.GdsControlNodeImageDetails != nil {
		result["gds_control_node_image_details"] = []interface{}{DistributedAutonomousDatabaseGdsControlNodeImageToMap(obj.GdsControlNodeImageDetails)}
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{DistributedAutonomousDbMetadataToMap(obj.Metadata)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.SupportingResourceId != nil {
		result["supporting_resource_id"] = string(*obj.SupportingResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func DistributedAutonomousDatabaseGdsControlNodeImageToMap(obj *oci_distributed_database.DistributedAutonomousDatabaseGdsControlNodeImage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func DistributedAutonomousDatabaseGsmToMap(obj oci_distributed_database.DistributedAutonomousDatabaseGsm) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComputeCount != nil {
		result["compute_count"] = float32(*obj.ComputeCount)
	}

	if obj.DataStorageSizeInGbs != nil {
		result["data_storage_size_in_gbs"] = float64(*obj.DataStorageSizeInGbs)
	}

	if obj.GsmImageDetails != nil {
		result["gsm_image_details"] = []interface{}{DistributedAutonomousDatabaseGsmImageToMap(obj.GsmImageDetails)}
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{DistributedAutonomousDbMetadataToMap(obj.Metadata)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["status"] = string(obj.Status)

	if obj.SupportingResourceId != nil {
		result["supporting_resource_id"] = string(*obj.SupportingResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeSslCertificateExpires != nil {
		result["time_ssl_certificate_expires"] = obj.TimeSslCertificateExpires.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func DistributedAutonomousDatabaseGsmImageToMap(obj *oci_distributed_database.DistributedAutonomousDatabaseGsmImage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func DistributedAutonomousDatabaseSummaryToMap(obj oci_distributed_database.DistributedAutonomousDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AutoResourceManagementConfig != nil {
		result["auto_resource_management_config"] = []interface{}{AutoResourceManagementConfigurationDetailsToMap(obj.AutoResourceManagementConfig)}
	}

	if obj.CharacterSet != nil {
		result["character_set"] = string(*obj.CharacterSet)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatabaseVersion != nil {
		result["database_version"] = string(*obj.DatabaseVersion)
	}

	result["db_deployment_type"] = string(obj.DbDeploymentType)

	result["db_workload_type"] = string(obj.DbWorkloadType)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ListenerPort != nil {
		result["listener_port"] = int(*obj.ListenerPort)
	}

	if obj.ListenerPortTls != nil {
		result["listener_port_tls"] = int(*obj.ListenerPortTls)
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{DistributedAutonomousDbMetadataToMap(obj.Metadata)}
	}

	if obj.NcharacterSet != nil {
		result["ncharacter_set"] = string(*obj.NcharacterSet)
	}

	result["notification_topic_ids"] = obj.NotificationTopicIds

	if obj.OnsPortLocal != nil {
		result["ons_port_local"] = int(*obj.OnsPortLocal)
	}

	if obj.OnsPortRemote != nil {
		result["ons_port_remote"] = int(*obj.OnsPortRemote)
	}

	if obj.Prefix != nil {
		result["prefix"] = string(*obj.Prefix)
	}

	result["private_endpoint_ids"] = obj.PrivateEndpointIds

	result["replication_method"] = string(obj.ReplicationMethod)

	result["sharding_method"] = string(obj.ShardingMethod)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemChunkCount != nil {
		result["system_chunk_count"] = int(*obj.SystemChunkCount)
	}

	if obj.SystemRaftReplicationUnitCount != nil {
		result["system_raft_replication_unit_count"] = int(*obj.SystemRaftReplicationUnitCount)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToDistributedAutonomousDbBackupConfig(fieldKeyFormat string) (oci_distributed_database.DistributedAutonomousDbBackupConfig, error) {
	result := oci_distributed_database.DistributedAutonomousDbBackupConfig{}

	if backupDestinationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")); ok {
		interfaces := backupDestinationDetails.([]interface{})
		tmp := make([]oci_distributed_database.DistributedAutonomousDbBackupDestination, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backup_destination_details"), stateDataIndex)
			converted, err := s.mapToDistributedAutonomousDbBackupDestination(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")) {
			result.BackupDestinationDetails = tmp
		}
	}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	return result, nil
}

func DistributedAutonomousDbBackupConfigToMap(obj *oci_distributed_database.DistributedAutonomousDbBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

	backupDestinationDetails := []interface{}{}
	for _, item := range obj.BackupDestinationDetails {
		backupDestinationDetails = append(backupDestinationDetails, DistributedAutonomousDbBackupDestinationToMap(item))
	}
	result["backup_destination_details"] = backupDestinationDetails

	if obj.RecoveryWindowInDays != nil {
		result["recovery_window_in_days"] = int(*obj.RecoveryWindowInDays)
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToDistributedAutonomousDbBackupDestination(fieldKeyFormat string) (oci_distributed_database.DistributedAutonomousDbBackupDestination, error) {
	result := oci_distributed_database.DistributedAutonomousDbBackupDestination{}

	if dbrsPolicyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dbrs_policy_id")); ok {
		tmp := dbrsPolicyId.(string)
		result.DbrsPolicyId = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if internetProxy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "internet_proxy")); ok {
		tmp := internetProxy.(string)
		result.InternetProxy = &tmp
	}

	if isRemote, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_remote")); ok {
		tmp := isRemote.(bool)
		result.IsRemote = &tmp
	}

	if remoteRegion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remote_region")); ok {
		tmp := remoteRegion.(string)
		result.RemoteRegion = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_distributed_database.DistributedAutonomousDbBackupDestinationTypeEnum(type_.(string))
	}

	return result, nil
}

func DistributedAutonomousDbBackupDestinationToMap(obj oci_distributed_database.DistributedAutonomousDbBackupDestination) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbrsPolicyId != nil {
		result["dbrs_policy_id"] = string(*obj.DbrsPolicyId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InternetProxy != nil {
		result["internet_proxy"] = string(*obj.InternetProxy)
	}

	if obj.IsRemote != nil {
		result["is_remote"] = bool(*obj.IsRemote)
	}

	if obj.RemoteRegion != nil {
		result["remote_region"] = string(*obj.RemoteRegion)
	}

	result["type"] = string(obj.Type)

	return result
}

func DistributedAutonomousDbMetadataToMap(obj *oci_distributed_database.DistributedAutonomousDbMetadata) map[string]interface{} {
	result := map[string]interface{}{}
	metadata := map[string]interface{}{}
	if obj != nil {
		for key, value := range obj.PropertiesMap {
			metadata[key] = value
		}
	}
	result["map"] = metadata

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToDistributedAutonomousDbMetadata(fieldKeyFormat string) (oci_distributed_database.DistributedAutonomousDbMetadata, error) {
	result := oci_distributed_database.DistributedAutonomousDbMetadata{}
	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "map")); ok {
		result.PropertiesMap = tfresource.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}
	return result, nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToPatchInstruction(fieldKeyFormat string) (oci_distributed_database.PatchInstruction, error) {
	var baseObject oci_distributed_database.PatchInstruction
	//discriminator
	operationRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation"))
	var operation string
	if ok {
		operation = operationRaw.(string)
	} else {
		operation = "" // default value
	}
	switch strings.ToLower(operation) {
	case strings.ToLower("INSERT"):
		details := oci_distributed_database.PatchInsertInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("MERGE"):
		details := oci_distributed_database.PatchMergeInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("REMOVE"):
		details := oci_distributed_database.PatchRemoveInstruction{}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown operation '%v' was specified", operation)
	}
	return baseObject, nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToVcnNsgIdsDetails(fieldKeyFormat string) (oci_distributed_database.VcnNsgIdsDetails, error) {
	result := oci_distributed_database.VcnNsgIdsDetails{}

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

	if vcnId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcn_id")); ok {
		tmp := vcnId.(string)
		result.VcnId = &tmp
	}

	return result, nil
}

func VcnNsgIdsDetailsToMap(obj oci_distributed_database.VcnNsgIdsDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_distributed_database.ChangeDistributedAutonomousDatabaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DistributedAutonomousDatabaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ChangeDistributedAutonomousDatabaseCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDistributedAutonomousDatabaseFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func observedAutonomousOperationalState(res *oci_distributed_database.DistributedAutonomousDatabase) string {
	return observedOperationalState(observedAutonomousTopology(res), func(metadata map[string]interface{}) string {
		return strings.ToUpper(actionString(metadata["supportingResourceLifecycleState"]))
	})
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) fetchOperationalState(ctx context.Context) (string, error) {
	if err := s.GetWithContext(ctx); err != nil {
		return "", err
	}
	state := observedAutonomousOperationalState(s.Res)
	return state, nil
}

func observedAutonomousTopology(res *oci_distributed_database.DistributedAutonomousDatabase) databaseTopologyValues {
	values := databaseTopologyValues{}
	for _, item := range res.CatalogDetails {
		values["catalog_details"] = append(values["catalog_details"], DistributedAutonomousDatabaseCatalogToMap(item))
	}
	for _, item := range res.UserShardSpaces {
		values["user_shard_spaces"] = append(values["user_shard_spaces"], AutonomousUserShardSpaceToMap(item))
	}
	for _, item := range res.SystemRaftClusters {
		values["system_raft_clusters"] = append(values["system_raft_clusters"], AutonomousSystemRaftClusterToMap(item))
	}
	for _, item := range res.CompositeRaftShardSpaces {
		values["composite_raft_shard_spaces"] = append(values["composite_raft_shard_spaces"], AutonomousCompositeRaftShardSpaceToMap(item))
	}
	if res.SystemDataGuardDatabases != nil {
		values["system_data_guard_databases"] = []interface{}{AutonomousSystemDataGuardDatabaseToMap(res.SystemDataGuardDatabases)}
	}
	for _, item := range res.CompositeDataGuardShardSpaces {
		values["composite_data_guard_shard_spaces"] = append(values["composite_data_guard_shard_spaces"], AutonomousCompositeDataGuardShardSpaceToMap(item))
	}
	return operationalStateTopologyStatuses(values, res)
}
