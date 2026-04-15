// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		// NOTE (WorkRequest long-running operation):
		// Distributed Autonomous Database create/update/delete are asynchronous and can take
		// significantly longer than the provider default timeouts.
		// With DefaultTimeout, the WorkRequest polling retry/backoff may hit Terraform's
		// context deadline and fail with:
		//   "now() + computed backoff duration exceeds request deadline".
		// Increase resource timeouts to allow WorkRequest completion.
		//Timeouts:      tfresource.DefaultTimeout,
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		CreateContext: createDistributedDatabaseDistributedDatabaseWithContext,
		ReadContext:   readDistributedDatabaseDistributedDatabaseWithContext,
		UpdateContext: updateDistributedDatabaseDistributedDatabaseWithContext,
		DeleteContext: deleteDistributedDatabaseDistributedDatabaseWithContext,
		CustomizeDiff: distributedDatabaseDistributedDatabaseCustomizeDiff,
		Schema: map[string]*schema.Schema{
			// Required
			"catalog_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"admin_password": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							Sensitive:        true,
							DiffSuppressFunc: suppressMaskedPasswordDiff,
						},
						"source": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"EXADB_XS",
								"NEW_VAULT_AND_CLUSTER",
							}, true),
						},

						// Optional
						"availability_domain": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},
						"db_storage_vault_details": {
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
									"additional_flash_cache_in_percent": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"high_capacity_database_storage": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
									"db_storage_vault_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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
						"peer_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
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
									"db_storage_vault_details": {
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
												"additional_flash_cache_in_percent": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"high_capacity_database_storage": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
												"db_storage_vault_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"protection_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"transport_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"vm_cluster_details": {
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
												"backup_network_nsg_ids": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Set:      tfresource.LiteralTypeHashCodeForSets,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"backup_subnet_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"domain": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"enabled_ecpu_count": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_diagnostics_events_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_health_monitoring_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_incident_logs_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"license_model": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
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
												"private_zone_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"ssh_public_keys": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"subnet_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"total_ecpu_count": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vm_file_system_storage_size": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"vm_cluster_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"vm_cluster_id": {
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
									"shard_group": {
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
						"peer_vm_cluster_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"shard_space": {
							Type:     schema.TypeString,
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
						"vm_cluster_details": {
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
									"backup_network_nsg_ids": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Set:      tfresource.LiteralTypeHashCodeForSets,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"backup_subnet_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"domain": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"enabled_ecpu_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_diagnostics_events_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_health_monitoring_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_incident_logs_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"license_model": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
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
									"private_zone_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"ssh_public_keys": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"total_ecpu_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"vm_file_system_storage_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vm_cluster_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"vm_cluster_id": {
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
						"db_home_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						/*"metadata": {
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
						},*/
						"metadata": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"map": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shard_group": {
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
			"db_deployment_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// NOTE (codegen bug): TOP-9471
			// The OCI CreateDistributedDatabase API does NOT accept a resource OCID as input.
			// The service generates the OCID and returns it in the response (and/or via work request).
			// The code generator incorrectly modeled the response `id` as a Required schema argument.
			// Terraform must manage the OCID via `d.SetId(...)` during Create and use `d.Id()` for
			// subsequent Get/Update/Delete calls.
			/*"distributed_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},*/
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
			"shard_details": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"admin_password": {
							Type:             schema.TypeString,
							Required:         true,
							Sensitive:        true,
							DiffSuppressFunc: suppressMaskedPasswordDiff,
						},
						"source": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"EXADB_XS",
								"NEW_VAULT_AND_CLUSTER",
							}, true),
						},

						// Optional
						"availability_domain": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},
						"db_storage_vault_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"additional_flash_cache_in_percent": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"high_capacity_database_storage": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
									"db_storage_vault_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"kms_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"kms_key_version_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"peer_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"availability_domain": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
									},
									"db_storage_vault_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"additional_flash_cache_in_percent": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"high_capacity_database_storage": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},

												// Computed
												"db_storage_vault_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"protection_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"transport_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vm_cluster_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"backup_network_nsg_ids": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													Set:      tfresource.LiteralTypeHashCodeForSets,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"backup_subnet_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"domain": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"enabled_ecpu_count": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"is_diagnostics_events_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_health_monitoring_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_incident_logs_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"license_model": {
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
												"private_zone_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ssh_public_keys": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"subnet_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"total_ecpu_count": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"vm_file_system_storage_size": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},

												// Computed
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"vm_cluster_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"vm_cluster_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"container_database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									/*"metadata": {
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
									},*/
									"metadata": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"map": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"shard_group": {
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
						"peer_vm_cluster_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"shard_space": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vault_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vm_cluster_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"backup_network_nsg_ids": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Set:      tfresource.LiteralTypeHashCodeForSets,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"backup_subnet_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"domain": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"enabled_ecpu_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"is_diagnostics_events_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"is_health_monitoring_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"is_incident_logs_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"license_model": {
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
									"private_zone_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssh_public_keys": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"total_ecpu_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"vm_file_system_storage_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vm_cluster_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"vm_cluster_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"container_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_home_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"map": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shard_group": {
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
			"sharding_method": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"chunks": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			//TOP-9494 - Remove Force New from db_backup_config to allow in-place updates
			"db_backup_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"auto_backup_window": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_full_backup_day": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_full_backup_window": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"backup_deletion_policy": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"backup_destination_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"dbrs_policy_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"internet_proxy": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"is_remote": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"is_zero_data_loss_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"remote_region": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vpc_password": {
										Type:      schema.TypeString,
										Optional:  true,
										Computed:  true,
										Sensitive: true,
										// API may omit write-only password in GET; suppress
										// post-create drift when config keeps a non-empty value.
										DiffSuppressFunc: suppressMaskedPasswordDiff,
									},
									"vpc_user": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"can_run_immediate_full_backup": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_auto_backup_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_remote_backup_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"recovery_window_in_days": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"remote_region": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"gsm_ssh_public_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"listener_port_tls": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			/*"patch_operations": {
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
							Type:             schema.TypeMap,
							Required:         true,
							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
						},

						// Optional

						// Computed
					},
				},
			},*/

			// NOTE:
			// The PATCH API expects the `value` field to be a JSON *string*, not a structured object.
			// Although the payload semantically represents an object, the service contract requires
			// it to be sent as a stringified JSON blob (see API examples).
			//
			// Using TypeString here allows callers to pass `jsonencode({...})` from HCL and ensures
			// the request body matches the API contract exactly.
			//
			// DO NOT change this to TypeMap / TypeList.
			// Doing so will cause Terraform to send an object instead of a JSON string and will
			// break PATCH operations at runtime.
			// TOP-9499
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
							Optional:         true,
							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
						},

						// Optional

						// Computed
					},
				},
			},
			"replication_factor": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"replication_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"replication_unit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"scan_listener_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_distributed_database.DistributedDatabaseLifecycleStateInactive),
					string(oci_distributed_database.DistributedDatabaseLifecycleStateActive),
				}, true),
			},
			"change_db_backup_config_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"configure_sharding_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// NOTE (CODEGEN GAP): expose IsRebalanceRequired for ConfigureSharding.
			// See JIRA: TOP-9490
			"configure_sharding_is_rebalance_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"download_gsm_certificate_signing_request_trigger": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: "This trigger/action API is deprecated.",
			},
			"generate_gsm_certificate_signing_request_trigger": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: "This trigger/action API is deprecated.",
			},
			"generate_wallet_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// NOTE (JIRA: TOP-9470):
			// The code generator exposes generate_wallet_trigger but omits the required
			// password parameter needed by the GenerateWallet API.
			// A separate schema attribute is added to allow users to provide the password
			// when triggering wallet generation.
			"generate_wallet_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			// GenerateDistributedAutonomousDatabaseWallet calls the GenerateWallet action API and stores the
			// returned wallet zip (binary body) in Terraform state as base64 so users can write it to disk
			// using local_file.content_base64.
			// TOP-9492
			"generate_wallet_downloaded_wallet_zip_base64": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"generate_wallet_downloaded_wallet_etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"generate_wallet_downloaded_wallet_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"generate_wallet_downloaded_wallet_content_length": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"move_replication_unit_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"recreate_failed_resource_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upload_signed_certificate_and_generate_wallet_trigger": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: "This trigger/action API is deprecated.",
			},
			"validate_network_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// WORKAROUND FOR CODEGEN LIMITATION:
			// The ValidateDistributedAutonomousDatabaseNetwork action API supports additional
			// optional request parameters (isSurrogate, resourceName, shardGroup) that are not
			// generated into the Terraform schema by default.
			//
			// These parameters are required for valid network validation scenarios and must be
			// explicitly exposed to users.
			// See JIRA: TOP-9477
			"validate_network_details": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_surrogate": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"resource_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shard_group": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			// CODEGEN FIX: TOP-9474
			// Start/Stop are imperative action APIs (not desired-state changes).
			// Terraform cannot infer when to run them unless a trigger changes.
			// Use monotonic int triggers to make the action explicit and idempotent.
			"start_database_trigger": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Increment this value to trigger StartDistributedDatabase action.",
			},
			"stop_database_trigger": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Increment this value to trigger StopDistributedDatabase action.",
			},

			// Computed
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
									"version_number": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						/*"metadata": {
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
						},*/
						"metadata": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"map": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
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
			"latest_gsm_image_details": {
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
						"version_number": {
							Type:     schema.TypeInt,
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
						"map": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
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
}

// IMPORTANT:
// This resource explicitly forbids combining CREATE with action triggers
// (e.g. start/stop, wallet generation, sharding, backup config changes).
//
// Reasoning:
// Terraform follows a strict lifecycle model where CREATE must be a pure
// provisioning operation. Action-style APIs (modeled here as *_trigger fields)
// are intended to be executed in separate apply steps, *after* the resource
// exists and its state is fully known.
//
// Mixing CREATE + action triggers in a single apply leads to:
//   - Undefined execution ordering
//   - Hidden behavior not visible in terraform plan
//   - Inconsistent state handling
//   - Hard-to-debug failures during apply
//
// This behavior was unintentionally enabled by code generation and is not
// a supported Terraform pattern.
//
// Enforcement:
// If any action trigger is set during CREATE, the provider will fail fast
// with a clear error *before* any resource creation occurs.
// Users must run a separate `terraform apply` after creation to execute
// action APIs.
// TOP-9510

// distributedDatabaseDistributedDatabaseCustomizeDiff validates PATCH request shape and
// decides when shard_details mutations should stay in-place vs require replacement.
func distributedDatabaseDistributedDatabaseCustomizeDiff(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	_ = ctx
	_ = meta

	if err := validateDistributedDatabasePatchOperations(diff); err != nil {
		return err
	}

	return configureDistributedDatabaseShardDetailsDiff(diff)
}

func validateDistributedDatabasePatchOperations(diff *schema.ResourceDiff) error {
	if diff.Id() == "" || !diff.HasChange("patch_operations") {
		return nil
	}

	_, newRaw := diff.GetChange("patch_operations")
	ops, ok := newRaw.([]interface{})
	if !ok || len(ops) == 0 {
		return nil
	}

	operationKinds := map[string]struct{}{}
	for i, raw := range ops {
		op, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}

		operation, _ := op["operation"].(string)
		normalizedOperation := strings.ToUpper(operation)
		operationKinds[normalizedOperation] = struct{}{}

		if normalizedOperation == "INSERT" {
			value, ok := op["value"].(string)
			if !ok || strings.TrimSpace(value) == "" {
				// TF_CODE_GEN: TERSI-4920-TOP-21 PATCH REMOVE instructions omit value in the API
				// contract and SDK, while INSERT still requires a JSON payload. Validate the
				// operation-specific requirement here instead of forcing a dummy value for REMOVE.
				return fmt.Errorf("patch_operations.%d.value must be set to a non-empty JSON string when operation is INSERT", i)
			}
		}
	}

	if len(operationKinds) > 1 {
		return fmt.Errorf("patch_operations must contain only one operation type per request; split INSERT, MERGE, and REMOVE instructions into separate applies")
	}

	return nil
}

func configureDistributedDatabaseShardDetailsDiff(diff *schema.ResourceDiff) error {
	if diff.Id() == "" || !diff.HasChange("shard_details") {
		return nil
	}

	if patchOperationsTargetShardDetails(diff) {
		return nil
	}

	// TF_CODE_GEN: TERSI-4920-TOP-22 shard_details participates in both the create payload
	// and observed topology state. Require replacement only for direct shard_details edits that
	// are not accompanied by a shardDetails PATCH so runtime insert/remove workflows stay in-place.
	return diff.ForceNew("shard_details")
}

func patchOperationsListTargetShardDetails(ops []interface{}) bool {
	if len(ops) == 0 {
		return false
	}

	for _, raw := range ops {
		op, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}

		selection, _ := op["selection"].(string)
		if strings.HasPrefix(selection, "shardDetails") {
			return true
		}
	}

	return false
}

func patchOperationsListHaveOperation(ops []interface{}, operation string) bool {
	if len(ops) == 0 {
		return false
	}

	normalizedOperation := strings.ToUpper(operation)
	matched := false

	for _, raw := range ops {
		op, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}

		currentOperation, _ := op["operation"].(string)
		if strings.ToUpper(currentOperation) != normalizedOperation {
			return false
		}

		matched = true
	}

	return matched
}

func patchOperationsTargetShardDetails(diff *schema.ResourceDiff) bool {
	ops, ok := diff.Get("patch_operations").([]interface{})
	if !ok || len(ops) == 0 {
		return false
	}

	return patchOperationsListTargetShardDetails(ops)
}

func configuredPatchOperationsTargetShardDetails(d *schema.ResourceData) bool {
	if ops, ok := d.Get("patch_operations").([]interface{}); ok && patchOperationsListTargetShardDetails(ops) {
		return true
	}

	value, ok := getConfiguredRawFieldValue(d, "patch_operations")
	if !ok || value.IsNull() || !value.IsKnown() || value.LengthInt() == 0 {
		return false
	}

	for _, op := range value.AsValueSlice() {
		if op.IsNull() || !op.IsKnown() {
			continue
		}

		selection := op.GetAttr("selection")
		if !selection.IsNull() && selection.IsKnown() && strings.HasPrefix(selection.AsString(), "shardDetails") {
			return true
		}
	}

	return false
}

// TF_CODE_GEN: TERSI-4920-TOP-25 admin_password exists only on create-time
// request shapes; GET/read models omit it because the field is write-only.
// Rehydrate configured passwords into flattened state to avoid perpetual drift
// and to keep subsequent create/action request builders from losing the secret.
func rehydrateConfiguredAdminPasswords(existingRaw interface{}, flattened []interface{}) []interface{} {
	existing, ok := existingRaw.([]interface{})
	if !ok || len(existing) == 0 || len(flattened) == 0 {
		return flattened
	}

	merged := make([]interface{}, len(flattened))
	for i, raw := range flattened {
		currentMap, ok := raw.(map[string]interface{})
		if !ok {
			merged[i] = raw
			continue
		}

		current := make(map[string]interface{}, len(currentMap)+1)
		for key, value := range currentMap {
			current[key] = value
		}

		if i < len(existing) {
			if existingMap, ok := existing[i].(map[string]interface{}); ok {
				if adminPassword, ok := existingMap["admin_password"].(string); ok && adminPassword != "" {
					current["admin_password"] = adminPassword
				}
			}
		}

		merged[i] = current
	}

	return merged
}

func createDistributedDatabaseDistributedDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	for _, attr := range []string{
		"change_db_backup_config_trigger",
		"configure_sharding_trigger",
		"download_gsm_certificate_signing_request_trigger",
		"generate_gsm_certificate_signing_request_trigger",
		"generate_wallet_trigger",
		"upload_signed_certificate_and_generate_wallet_trigger",
		"validate_network_trigger",
		"start_database_trigger",
		"stop_database_trigger",
	} {
		v, ok := d.GetOkExists(attr)
		if !ok {
			continue
		}

		// triggers are TypeInt: block only when non-zero
		if i, ok := v.(int); ok && i != 0 {
			return diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Invalid CREATE + action trigger combination",
				Detail: fmt.Sprintf(
					"Trigger %q cannot be used during resource creation.\n\n"+
						"CREATE operations must not be combined with action APIs.\n"+
						"Please run `terraform apply` again after the resource is created to execute this action.",
					attr,
				),
			}}
		}
	}

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	return nil
}

func readDistributedDatabaseDistributedDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

// TOP-9493: Refactored to first handle standard updates (PUT/PATCH/compartment move),
func updateDistributedDatabaseDistributedDatabaseWithContext(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	disabled := []string{
		"upload_signed_certificate_and_generate_wallet_trigger",
		"generate_gsm_certificate_signing_request_trigger",
		"download_gsm_certificate_signing_request_trigger",
	}

	if diags := rejectDisabledTriggers(d, disabled, "UPDATE"); diags != nil {
		return diags
	}

	// ---- 1) Detect standard update categories (handled in UpdateWithContext) ----
	needsCompartmentMove := d.HasChange("compartment_id")

	needsPutUpdate :=
		d.HasChange("display_name") ||
			d.HasChange("freeform_tags") ||
			d.HasChange("defined_tags")

	needsPatch := false
	if v, ok := d.GetOkExists("patch_operations"); ok && d.HasChange("patch_operations") {
		if ops, ok2 := v.([]interface{}); ok2 && len(ops) > 0 {
			needsPatch = true
		}
	}

	// If any standard update is needed, do it FIRST.
	// This avoids: action -> immediate PUT/PATCH while resource is UPDATING.
	needsStandardUpdate := needsPutUpdate || needsPatch || needsCompartmentMove
	if needsStandardUpdate {
		if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	// ---- 2) Action triggers (increment-only) ----
	actionInvoked := false

	// Helper: increment-only trigger check.
	// If new <= old, keep old value in state to avoid drift.
	triggerBumped := func(attr string) (bool, int, int) {
		if _, ok := d.GetOkExists(attr); !ok || !d.HasChange(attr) {
			return false, 0, 0
		}
		oldRaw, newRaw := d.GetChange(attr)
		oldV := oldRaw.(int)
		newV := newRaw.(int)

		if newV <= oldV {
			_ = d.Set(attr, oldV)
			return false, oldV, newV
		}
		return true, oldV, newV
	}

	// START
	if ok, _, newV := triggerBumped("start_database_trigger"); ok {
		if err := sync.StartDistributedDatabase(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("start_database_trigger", newV)
	}

	// STOP
	if ok, _, newV := triggerBumped("stop_database_trigger"); ok {
		if err := sync.StopDistributedDatabase(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("stop_database_trigger", newV)
	}

	// Generate GSM CSR
	if ok, _, newV := triggerBumped("generate_gsm_certificate_signing_request_trigger"); ok {
		if err := sync.GenerateDistributedDatabaseGsmCertificateSigningRequest(); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("generate_gsm_certificate_signing_request_trigger", newV)
	}

	// Change backup config
	if ok, _, newV := triggerBumped("change_db_backup_config_trigger"); ok {
		if err := sync.ChangeDistributedDbBackupConfig(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("change_db_backup_config_trigger", newV)
	}

	if ok, _, newV := triggerBumped("move_replication_unit_trigger"); ok {
		if err := sync.MoveDistributedDatabaseReplicationUnit(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("move_replication_unit_trigger", newV)
	}

	if ok, _, newV := triggerBumped("recreate_failed_resource_trigger"); ok {
		if err := sync.RecreateFailedDistributedDatabaseResource(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("recreate_failed_resource_trigger", newV)
	}

	// Configure sharding
	if ok, _, newV := triggerBumped("configure_sharding_trigger"); ok {
		if err := sync.ConfigureDistributedDatabaseSharding(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("configure_sharding_trigger", newV)
	}

	// Download GSM CSR
	if ok, _, newV := triggerBumped("download_gsm_certificate_signing_request_trigger"); ok {
		if err := sync.DownloadDistributedDatabaseGsmCertificateSigningRequest(); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("download_gsm_certificate_signing_request_trigger", newV)
	}

	// Generate wallet
	if ok, _, newV := triggerBumped("generate_wallet_trigger"); ok {
		if err := sync.GenerateDistributedDatabaseWallet(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("generate_wallet_trigger", newV)
	}

	// Upload signed cert + generate wallet
	if ok, _, newV := triggerBumped("upload_signed_certificate_and_generate_wallet_trigger"); ok {
		if err := sync.UploadDistributedDatabaseSignedCertificateAndGenerateWallet(); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("upload_signed_certificate_and_generate_wallet_trigger", newV)
	}

	// Validate network
	if ok, _, newV := triggerBumped("validate_network_trigger"); ok {
		if err := sync.ValidateDistributedDatabaseNetwork(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("validate_network_trigger", newV)
	}

	// If this was action-only, do NOT run any extra update logic.
	// (We already executed UpdateResourceWithContext earlier only when needed.)
	if actionInvoked && !needsStandardUpdate {
		return nil
	}

	return nil
}

/*
func updateDistributedDatabaseDistributedDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()
	disabled := []string{
		"upload_signed_certificate_and_generate_wallet_trigger",
		"generate_gsm_certificate_signing_request_trigger",
		"download_gsm_certificate_signing_request_trigger",
	}

	if diags := rejectDisabledTriggers(d, disabled, "UPDATE"); diags != nil {
		return diags
	}
	/* powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_distributed_database.DistributedDatabaseLifecycleStateActive == oci_distributed_database.DistributedDatabaseLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_distributed_database.DistributedDatabaseLifecycleStateInactive == oci_distributed_database.DistributedDatabaseLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	} */

/* if powerOn {

// NOTE:
// This function contains multiple workarounds for known Terraform code
// generator issues (incorrect return types, missing context propagation).
// These changes are required for compilation and correctness but should
// be removed once the generator is fixed.
// See JIRA: TOP-9394
*/ /*
	if err := sync.StartDistributedDatabase(); err != nil {
		return tfresource.HandleDiagError(m, err)
	}*/ /*

	if err := sync.StartDistributedDatabase(ctx); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	sync.D.Set("state", oci_distributed_database.DistributedDatabaseLifecycleStateActive)
} */

// CODEGEN FIX: TOP-9474
// Start/Stop are imperative action APIs (not desired-state changes).
// Terraform cannot infer when to run them unless a trigger changes.
// Use monotonic int triggers to make the action explicit and idempotent.
/*if d.HasChange("start_database_trigger") && d.HasChange("stop_database_trigger") {
	log.Printf("start_database_trigger and stop_database_trigger cannot be changed in the same apply")
}

// START trigger
if oldRaw, newRaw := d.GetChange("start_database_trigger"); d.HasChange("start_database_trigger") {
	oldV := oldRaw.(int)
	newV := newRaw.(int)
	if newV <= oldV {
		log.Printf("start_database_trigger must be incremented to retrigger start action (old=%d new=%d)", oldV, newV)
		return nil
	}
	if err := sync.StartDistributedDatabase(context.Background()); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	// Preserve trigger in state
	_ = d.Set("start_database_trigger", newV)
}

// STOP trigger
if oldRaw, newRaw := d.GetChange("stop_database_trigger"); d.HasChange("stop_database_trigger") {
	oldV := oldRaw.(int)
	newV := newRaw.(int)
	if newV <= oldV {
		log.Printf("stop_database_trigger must be incremented to retrigger stop action (old=%d new=%d)", oldV, newV)
		return nil
	}
	if err := sync.StopDistributedDatabase(context.Background()); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	// Preserve trigger in state
	_ = d.Set("stop_database_trigger", newV)
}

if _, ok := sync.D.GetOkExists("change_db_backup_config_trigger"); ok && sync.D.HasChange("change_db_backup_config_trigger") {
	oldRaw, newRaw := sync.D.GetChange("change_db_backup_config_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.ChangeDistributedDbBackupConfig(ctx)

		if err != nil {

			// NOTE (TOP-9389):
			// Context-based CRUD entrypoints return diag.Diagnostics; the generator incorrectly
			// emits `return err` / `return fmt.Errorf(...)` in trigger branches. Wrap as diagnostics.
			//return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("change_db_backup_config_trigger", oldRaw)
		// NOTE (TOP-9389): return diagnostics instead of a raw error in UpdateContext.
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

if _, ok := sync.D.GetOkExists("configure_sharding_trigger"); ok && sync.D.HasChange("configure_sharding_trigger") {
	oldRaw, newRaw := sync.D.GetChange("configure_sharding_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.ConfigureDistributedDatabaseSharding(ctx)

		if err != nil {
			// NOTE (TOP-9389):
			// Context-based CRUD entrypoints return diag.Diagnostics; the generator incorrectly
			// emits `return err` / `return fmt.Errorf(...)` in trigger branches. Wrap as diagnostics.
			//return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("configure_sharding_trigger", oldRaw)
		// NOTE (TOP-9389): return diagnostics instead of a raw error in UpdateContext.
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

if _, ok := sync.D.GetOkExists("download_gsm_certificate_signing_request_trigger"); ok && sync.D.HasChange("download_gsm_certificate_signing_request_trigger") {
	oldRaw, newRaw := sync.D.GetChange("download_gsm_certificate_signing_request_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.DownloadDistributedDatabaseGsmCertificateSigningRequest()

		if err != nil {
			// NOTE (TOP-9389):
			// Context-based CRUD entrypoints return diag.Diagnostics; the generator incorrectly
			// emits `return err` / `return fmt.Errorf(...)` in trigger branches. Wrap as diagnostics.
			//return err
		}
	} else {
		sync.D.Set("download_gsm_certificate_signing_request_trigger", oldRaw)
		// NOTE (TOP-9389): return diagnostics instead of a raw error in UpdateContext.
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

if _, ok := sync.D.GetOkExists("generate_gsm_certificate_signing_request_trigger"); ok && sync.D.HasChange("generate_gsm_certificate_signing_request_trigger") {
	oldRaw, newRaw := sync.D.GetChange("generate_gsm_certificate_signing_request_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.GenerateDistributedDatabaseGsmCertificateSigningRequest()

		if err != nil {
			// NOTE (TOP-9389):
			// Context-based CRUD entrypoints return diag.Diagnostics; the generator incorrectly
			// emits `return err` / `return fmt.Errorf(...)` in trigger branches. Wrap as diagnostics.
			//return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("generate_gsm_certificate_signing_request_trigger", oldRaw)
		// NOTE (TOP-9389): return diagnostics instead of a raw error in UpdateContext.
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

if _, ok := sync.D.GetOkExists("generate_wallet_trigger"); ok && sync.D.HasChange("generate_wallet_trigger") {
	oldRaw, newRaw := sync.D.GetChange("generate_wallet_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.GenerateDistributedDatabaseWallet()

		if err != nil {
			// NOTE (TOP-9389):
			// Context-based CRUD entrypoints return diag.Diagnostics; the generator incorrectly
			// emits `return err` / `return fmt.Errorf(...)` in trigger branches. Wrap as diagnostics.
			//return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("generate_wallet_trigger", oldRaw)
		// NOTE (TOP-9389): return diagnostics instead of a raw error in UpdateContext.
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

<<<<<<< ours
	if _, ok := sync.D.GetOkExists("move_replication_unit_trigger"); ok && sync.D.HasChange("move_replication_unit_trigger") {
		oldRaw, newRaw := sync.D.GetChange("move_replication_unit_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.MoveDistributedDatabaseReplicationUnit(ctx)

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
			err := sync.RecreateFailedDistributedDatabaseResource(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("recreate_failed_resource_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("upload_signed_certificate_and_generate_wallet_trigger"); ok && sync.D.HasChange("upload_signed_certificate_and_generate_wallet_trigger") {
		oldRaw, newRaw := sync.D.GetChange("upload_signed_certificate_and_generate_wallet_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.UploadDistributedDatabaseSignedCertificateAndGenerateWallet(ctx)
=======
if _, ok := sync.D.GetOkExists("upload_signed_certificate_and_generate_wallet_trigger"); ok && sync.D.HasChange("upload_signed_certificate_and_generate_wallet_trigger") {
	oldRaw, newRaw := sync.D.GetChange("upload_signed_certificate_and_generate_wallet_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.UploadDistributedDatabaseSignedCertificateAndGenerateWallet()
>>>>>>> theirs

		if err != nil {
			// NOTE (TOP-9389):
			// Context-based CRUD entrypoints return diag.Diagnostics; the generator incorrectly
			// emits `return err` / `return fmt.Errorf(...)` in trigger branches. Wrap as diagnostics.
			// return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("upload_signed_certificate_and_generate_wallet_trigger", oldRaw)
		// NOTE (TOP-9389): return diagnostics instead of a raw error in UpdateContext.
		// return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

if _, ok := sync.D.GetOkExists("validate_network_trigger"); ok && sync.D.HasChange("validate_network_trigger") {
	oldRaw, newRaw := sync.D.GetChange("validate_network_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.ValidateDistributedDatabaseNetwork(ctx)

		if err != nil {
			// NOTE (TOP-9389):
			// Context-based CRUD entrypoints return diag.Diagnostics; the generator incorrectly
			// emits `return err` / `return fmt.Errorf(...)` in trigger branches. Wrap as diagnostics.
			//return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("validate_network_trigger", oldRaw)
		// NOTE (TOP-9389): return diagnostics instead of a raw error in UpdateContext.
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

// WORKAROUND FOR GENERATED CODE ISSUE:
// The Terraform code generator invokes UpdateResourceWithContext using an
// outdated signature that does not include context.Context, causing a
// compile-time error.
//
// The current helper signature requires:
//   UpdateResourceWithContext(ctx context.Context, d schema.ResourceData, updater ResourceUpdaterWithContext)
//
// This call explicitly passes `ctx` as a temporary workaround.
// DO NOT remove unless the generator is fixed.
// See JIRA: TOP-9395
/*	if err := tfresource.UpdateResourceWithContext(d, sync); err != nil {
	return tfresource.HandleDiagError(m, err)
}*/
/*if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
	return tfresource.HandleDiagError(m, err)
}

/* if powerOff {
// NOTE:
// This function contains multiple workarounds for known Terraform code
// generator issues (incorrect return types, missing context propagation).
// These changes are required for compilation and correctness but should
// be removed once the generator is fixed.
// See JIRA: TOP-9394
*/ /*
			if err := sync.StopDistributedDatabase(); err != nil {
	return tfresource.HandleDiagError(m, err)
}*/ /*
	if err := sync.StopDistributedDatabase(ctx); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	sync.D.Set("state", oci_distributed_database.DistributedDatabaseLifecycleStateInactive)
} */

/*return nil
}*/

func deleteDistributedDatabaseDistributedDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DistributedDatabaseDistributedDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_distributed_database.DistributedDbServiceClient
	Res                    *oci_distributed_database.DistributedDatabase
	PatchResponse          *oci_distributed_database.DistributedDatabase
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_distributed_database.DistributedDbWorkRequestServiceClient
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_distributed_database.DistributedDatabaseLifecycleStateCreating),
	}
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_distributed_database.DistributedDatabaseLifecycleStateActive),
		string(oci_distributed_database.DistributedDatabaseLifecycleStateNeedsAttention),
		// NOTE:
		// Distributed Database creation is asynchronous.
		// After a successful Create operation, the service intentionally returns
		// the resource in lifecycle state INACTIVE.
		//
		// INACTIVE is a valid terminal post-create state. The resource transitions
		// to ACTIVE only after specific action APIs (e.g. Configure Sharding)
		// are executed successfully.
		//
		// The provider MUST treat INACTIVE as a successful Create completion state.
		// Waiting for ACTIVE during Create is incorrect and causes false failures.
		// JIRA: TOP-9466
		string(oci_distributed_database.DistributedDatabaseLifecycleStateInactive),
	}
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_distributed_database.DistributedDatabaseLifecycleStateDeleting),
	}
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_distributed_database.DistributedDatabaseLifecycleStateDeleted),
	}
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_distributed_database.CreateDistributedDatabaseRequest{}

	if catalogDetails, ok := s.D.GetOkExists("catalog_details"); ok {
		interfaces := catalogDetails.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedDatabaseCatalogDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "catalog_details", stateDataIndex)
			converted, err := s.mapToCreateDistributedDatabaseCatalogDetails(fieldKeyFormat)
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

	if chunks, ok := s.D.GetOkExists("chunks"); ok {
		tmp := chunks.(int)
		request.Chunks = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if databaseVersion, ok := s.D.GetOkExists("database_version"); ok {
		tmp := databaseVersion.(string)
		request.DatabaseVersion = &tmp
	}

	if dbBackupConfig, ok := s.D.GetOkExists("db_backup_config"); ok {
		if tmpList := dbBackupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_backup_config", 0)
			tmp, err := s.mapToDistributedDbBackupConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DbBackupConfig = &tmp
		}
	}

	if dbDeploymentType, ok := s.D.GetOkExists("db_deployment_type"); ok {
		request.DbDeploymentType = oci_distributed_database.CreateDistributedDatabaseDetailsDbDeploymentTypeEnum(dbDeploymentType.(string))
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

	if gsmSshPublicKey, ok := s.D.GetOkExists("gsm_ssh_public_key"); ok {
		tmp := gsmSshPublicKey.(string)
		request.GsmSshPublicKey = &tmp
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

	if replicationFactor, ok := s.D.GetOkExists("replication_factor"); ok {
		tmp := replicationFactor.(int)
		request.ReplicationFactor = &tmp
	}

	if replicationMethod, ok := s.D.GetOkExists("replication_method"); ok {
		request.ReplicationMethod = oci_distributed_database.CreateDistributedDatabaseDetailsReplicationMethodEnum(replicationMethod.(string))
	}

	if replicationUnit, ok := s.D.GetOkExists("replication_unit"); ok {
		tmp := replicationUnit.(int)
		request.ReplicationUnit = &tmp
	}

	if scanListenerPort, ok := s.D.GetOkExists("scan_listener_port"); ok {
		tmp := scanListenerPort.(int)
		request.ScanListenerPort = &tmp
	}

	if shardDetails, ok := s.D.GetOkExists("shard_details"); ok {
		interfaces := shardDetails.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedDatabaseShardDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shard_details", stateDataIndex)
			converted, err := s.mapToCreateDistributedDatabaseShardDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("shard_details") {
			request.ShardDetails = tmp
		}
	}

	if shardingMethod, ok := s.D.GetOkExists("sharding_method"); ok {
		request.ShardingMethod = oci_distributed_database.CreateDistributedDatabaseDetailsShardingMethodEnum(shardingMethod.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.CreateDistributedDatabase(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}

	err = s.getDistributedDatabaseFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The code generator omits the required context.Context argument when calling
	// getDistributedDatabaseFromWorkRequest, resulting in a compile-time error.
	// This call explicitly propagates `ctx`.
	// See JIRA: TOP-9396
	//err = s.Patch()
	// 	err = s.Patch(ctx)
	// NOTE (CODEGEN WORKAROUND):
	// The code generator may invoke Patch() when `patch_operations` changes to an empty list,
	// resulting in an empty PATCH payload (`items: []`). The service rejects this with:
	//   OSD-10162: Items provided in payload of patch operation cannot be empty.
	//
	// Additionally, OCI Go SDK models PatchInstruction as a polymorphic interface, not a struct.
	// Generated code must not assume fields like Operation/Path/Value/From exist on the interface.
	//
	// Workaround: only call Patch() when `patch_operations` is both changed and non-empty.
	// See  Jira: TOP-9473
	if v, ok := s.D.GetOkExists("patch_operations"); ok {
		ops := v.([]interface{})
		if len(ops) > 0 && s.D.HasChange("patch_operations") {
			if err := s.Patch(ctx); err != nil {
				return err
			}
		}
	}
	if err != nil {
		log.Printf("[ERROR] Failed to execute Patch operation: %v", err)
		return err
	}
	return nil
}

// WORKAROUND FOR GENERATED CODE ISSUE:
// The code generator omits the required context.Context argument when calling
// getDistributedDatabaseFromWorkRequest, resulting in a compile-time error.
// This call explicitly propagates `ctx`.
// See JIRA: TOP-9396
// func (s *DistributedDatabaseDistributedDatabaseResourceCrud) Patch() error {
//func (s *DistributedDatabaseDistributedDatabaseResourceCrud) Patch(ctx context.Context) error {
//	request := oci_distributed_database.PatchDistributedDatabaseRequest{}
//
//	//if distributedDatabaseId, ok := s.D.GetOkExists("id"); ok {
//	//	tmp := distributedDatabaseId.(string)
//	//	request.DistributedDatabaseId = &tmp
//	//}
//	request.DistributedDatabaseId = &id
//
//	if patchOperations, ok := s.D.GetOkExists("patch_operations"); ok {
//		interfaces := patchOperations.([]interface{})
//		tmp := make([]oci_distributed_database.PatchInstruction, len(interfaces))
//		for i := range interfaces {
//			stateDataIndex := i
//			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patch_operations", stateDataIndex)
//			converted, err := s.mapToPatchInstruction(fieldKeyFormat)
//			if err != nil {
//				return err
//			}
//			tmp[i] = converted
//		}
//		if len(tmp) != 0 || s.D.HasChange("patch_operations") {
//			request.Items = tmp
//		}
//	}
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")
//	response, err := s.Client.PatchDistributedDatabase(ctx, request)
//	if err != nil {
//		return err
//	}
//
//	workId := response.OpcWorkRequestId
//	// WORKAROUND FOR GENERATED CODE ISSUE:
//	// The code generator omits the required context.Context argument when calling
//	// getDistributedDatabaseFromWorkRequest, resulting in a compile-time error.
//	// This call explicitly propagates `ctx`.
//	// See JIRA: TOP-9396
//	//return s.getDistributedDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
//	return s.getDistributedDatabaseFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
//}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) Patch(ctx context.Context) error {
	request := oci_distributed_database.PatchDistributedDatabaseRequest{}

	id := s.D.Id()
	request.DistributedDatabaseId = &id

	if patchOperations, ok := s.D.GetOkExists("patch_operations"); ok {
		interfaces := patchOperations.([]interface{})
		tmp := make([]oci_distributed_database.PatchInstruction, len(interfaces))
		for i := range interfaces {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patch_operations", i)
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

	response, err := s.Client.PatchDistributedDatabase(ctx, request) // use ctx, not Background()
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId == nil || *workId == "" {
		return fmt.Errorf("missing opc-work-request-id for PatchDistributedDatabase")
	}

	// TF_CODE_GEN: TERSI-4920-TOP-19 PatchDistributedDatabase work requests can finish SUCCEEDED
	// while resources[] reports a non-terminal actionType (for example IN_PROGRESS); rely on WR
	// terminal status and then refresh the existing DDB instead of failing on metadata drift.
	if err := waitForDistributedDatabaseWorkRequestCompletion(
		ctx,
		workId,
		"distributeddatabase",
		oci_distributed_database.ActionTypeUpdated,
		s.D.Timeout(schema.TimeoutUpdate),
		s.DisableNotFoundRetries,
		s.WorkRequestClient,
	); err != nil {
		return fmt.Errorf("PatchDistributedDatabase failed for distributed database %s (work request %s): %w", s.D.Id(), *workId, err)
	}

	// PATCH can legitimately settle back into ACTIVE/INACTIVE/NEEDS_ATTENTION depending on the
	// previous state and any shard-level issue surfaced during reconciliation.
	stable := func() bool {
		if err := s.Get(); err != nil {
			log.Printf("[WARN] post-patch Get() failed: %v", err)
			return false
		}
		if s.Res == nil {
			return false
		}
		st := s.Res.LifecycleState
		return st == oci_distributed_database.DistributedDatabaseLifecycleStateActive ||
			st == oci_distributed_database.DistributedDatabaseLifecycleStateInactive ||
			st == oci_distributed_database.DistributedDatabaseLifecycleStateNeedsAttention
	}

	return tfresource.WaitForResourceCondition(s, stable, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) getDistributedDatabaseFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_distributed_database.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	distributedDatabaseId, err := distributedDatabaseWaitForWorkRequest(ctx, workId, "distributeddatabase",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*distributedDatabaseId)

	return s.GetWithContext(ctx)
}

func distributedDatabaseWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "distributed_database", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_distributed_database.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func distributedDatabaseWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_distributed_database.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_distributed_database.DistributedDbWorkRequestServiceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "distributed_database")
	retryPolicy.ShouldRetryOperation = distributedDatabaseWorkRequestShouldRetryFunc(timeout)

	response := oci_distributed_database.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_distributed_database.OperationStatusInProgress),
			string(oci_distributed_database.OperationStatusAccepted),
			string(oci_distributed_database.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_distributed_database.OperationStatusSucceeded),
			string(oci_distributed_database.OperationStatusFailed),
			string(oci_distributed_database.OperationStatusCanceled),
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
	if identifier == nil || response.Status == oci_distributed_database.OperationStatusFailed || response.Status == oci_distributed_database.OperationStatusCanceled {
		return nil, getErrorFromDistributedDatabaseDistributedDatabaseWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func waitForDistributedDatabaseWorkRequestCompletion(ctx context.Context, wId *string, entityType string, action oci_distributed_database.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_distributed_database.DistributedDbWorkRequestServiceClient) error {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "distributed_database")
	retryPolicy.ShouldRetryOperation = distributedDatabaseWorkRequestShouldRetryFunc(timeout)

	response := oci_distributed_database.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_distributed_database.OperationStatusInProgress),
			string(oci_distributed_database.OperationStatusAccepted),
			string(oci_distributed_database.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_distributed_database.OperationStatusSucceeded),
			string(oci_distributed_database.OperationStatusFailed),
			string(oci_distributed_database.OperationStatusCanceled),
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
	if _, err := stateConf.WaitForState(); err != nil {
		return err
	}

	if response.Status == oci_distributed_database.OperationStatusFailed || response.Status == oci_distributed_database.OperationStatusCanceled {
		return getErrorFromDistributedDatabaseDistributedDatabaseWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return nil
}

func getErrorFromDistributedDatabaseDistributedDatabaseWorkRequest(ctx context.Context, client *oci_distributed_database.DistributedDbWorkRequestServiceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_distributed_database.ActionTypeEnum) error {
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
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.GetDistributedDatabaseRequest{}

	tmp := s.D.Id()
	request.DistributedDatabaseId = &tmp

	// WORKAROUND FOR GENERATED CODE ISSUE:
	//
	// The Terraform provider generator attempts to populate
	// GetDistributedAutonomousDatabaseRequest.Metadata using
	// DistributedAutonomousDbMetadata.
	//
	// However, in the OCI Go SDK, the GetDistributedAutonomousDatabaseRequest
	// defines Metadata as a *string (not a metadata struct), and the
	// DistributedAutonomousDbMetadata type is returned only on the response
	// model, not accepted as request input.
	//
	// Passing metadata here results in a compile-time type mismatch and is
	// semantically incorrect. Metadata must be read from the response only
	// and must not be sent on the GET request.
	//
	// Remove once generator is fixed.
	// See JIRA: TOP-9397

	/*if metadata, ok := s.D.GetOkExists("metadata"); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", 0)
			tmp, err := s.mapToDistributedDbMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Metadata = &tmp
		}
	}*/

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.GetDistributedDatabase(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DistributedDatabase
	return nil
}

/*func (s *DistributedDatabaseDistributedDatabaseResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_distributed_database.UpdateDistributedDatabaseRequest{}

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
	request.DistributedDatabaseId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.UpdateDistributedDatabase(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DistributedDatabase
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The code generator omits the required context.Context argument when calling
	// getDistributedDatabaseFromWorkRequest, resulting in a compile-time error.
	// This call explicitly propagates `ctx`.
	// See JIRA: TOP-9396
	//err = s.Patch()
	//err = s.Patch(ctx)

	// NOTE (CODEGEN WORKAROUND):
	// The code generator may invoke Patch() when `patch_operations` changes to an empty list,
	// resulting in an empty PATCH payload (`items: []`). The service rejects this with:
	//   OSD-10162: Items provided in payload of patch operation cannot be empty.
	//
	// Additionally, OCI Go SDK models PatchInstruction as a polymorphic interface, not a struct.
	// Generated code must not assume fields like Operation/Path/Value/From exist on the interface.
	//
	// Workaround: only call Patch() when `patch_operations` is both changed and non-empty.
	// See  Jira: TOP-9473
	if v, ok := s.D.GetOkExists("patch_operations"); ok {
		ops := v.([]interface{})
		if len(ops) > 0 && s.D.HasChange("patch_operations") {
			if err := s.Patch(ctx); err != nil {
				return err
			}
		}
	}
	if err != nil {
		log.Printf("[ERROR] Failed to execute Patch operation: %v", err)
		return err
	}
	return nil
}*/

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) UpdateWithContext(ctx context.Context) error {
	// 1) Detect which categories of updates are needed
	needsCompartmentMove := s.D.HasChange("compartment_id")

	needsPutUpdate :=
		s.D.HasChange("display_name") ||
			s.D.HasChange("freeform_tags") ||
			s.D.HasChange("defined_tags")
	// IMPORTANT: compartment_id is handled via ChangeCompartment API (WR), not PUT.

	needsPatch := false
	if v, ok := s.D.GetOkExists("patch_operations"); ok && s.D.HasChange("patch_operations") {
		if ops, ok2 := v.([]interface{}); ok2 && len(ops) > 0 {
			needsPatch = true
		}
	}

	// 2) If there is nothing to do besides compartment move, do it and return.
	if needsCompartmentMove && !needsPutUpdate && !needsPatch {
		if compartment, ok := s.D.GetOkExists("compartment_id"); ok {
			oldRaw, newRaw := s.D.GetChange("compartment_id")
			oldStr, _ := oldRaw.(string)
			newStr, _ := newRaw.(string)
			// Keep your guard to avoid weird empty transitions
			if oldStr != "" && newStr != "" {
				return s.updateCompartment(ctx, compartment)
			}
		}
		return nil
	}

	// 3) Execute PUT update first (only if PUT-able fields changed)
	if needsPutUpdate {
		request := oci_distributed_database.UpdateDistributedDatabaseRequest{}

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

		id := s.D.Id()
		request.DistributedDatabaseId = &id
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

		response, err := s.Client.UpdateDistributedDatabase(ctx, request)
		if err != nil {
			return err
		}

		// Keep the model for downstream waits/reads
		s.Res = &response.DistributedDatabase
	}

	// 4) Execute PATCH second (only when patch_operations changed AND non-empty)
	// NOTE (TOP-9460): avoid sending empty PATCH payload (items: []) which service rejects.
	if needsPatch {
		if err := s.Patch(ctx); err != nil {
			return err
		}
	}

	// 5) Execute compartment move last (async WR) to avoid PUT/PATCH being rejected in UPDATING state
	if needsCompartmentMove {
		if compartment, ok := s.D.GetOkExists("compartment_id"); ok {
			oldRaw, newRaw := s.D.GetChange("compartment_id")
			oldStr, _ := oldRaw.(string)
			newStr, _ := newRaw.(string)
			if oldStr != "" && newStr != "" {
				if err := s.updateCompartment(ctx, compartment); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_distributed_database.DeleteDistributedDatabaseRequest{}

	tmp := s.D.Id()
	request.DistributedDatabaseId = &tmp

	if mustDeleteInfra, ok := s.D.GetOkExists("must_delete_infra"); ok {
		tmp := mustDeleteInfra.(bool)
		request.MustDeleteInfra = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.DeleteDistributedDatabase(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := distributedDatabaseWaitForWorkRequest(ctx, workId, "distributeddatabase",
		oci_distributed_database.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) SetData() error {
	catalogDetails := []interface{}{}
	for _, item := range s.Res.CatalogDetails {
		catalogDetails = append(catalogDetails, DistributedDatabaseCatalogToMap(item))
	}
	catalogDetails = rehydrateConfiguredAdminPasswords(s.D.Get("catalog_details"), catalogDetails)
	s.D.Set("catalog_details", catalogDetails)

	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.Chunks != nil {
		s.D.Set("chunks", *s.Res.Chunks)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{DistributedDbConnectionStringToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DbBackupConfig != nil {
		s.D.Set("db_backup_config", []interface{}{DistributedDbBackupConfigToMap(s.Res.DbBackupConfig)})
	} else {
		s.D.Set("db_backup_config", nil)
	}

	s.D.Set("db_deployment_type", s.Res.DbDeploymentType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	gsmDetails := []interface{}{}
	for _, item := range s.Res.GsmDetails {
		gsmDetails = append(gsmDetails, DistributedDatabaseGsmToMap(item))
	}
	s.D.Set("gsm_details", gsmDetails)

	if s.Res.GsmSshPublicKey != nil {
		s.D.Set("gsm_ssh_public_key", *s.Res.GsmSshPublicKey)
	}

	if s.Res.LatestGsmImageDetails != nil {
		s.D.Set("latest_gsm_image_details", []interface{}{DistributedDbGsmImageToMap(s.Res.LatestGsmImageDetails)})
	} else {
		s.D.Set("latest_gsm_image_details", nil)
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
		s.D.Set("metadata", []interface{}{DistributedDbMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", []interface{}{})
	}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

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

	if s.Res.ReplicationFactor != nil {
		s.D.Set("replication_factor", *s.Res.ReplicationFactor)
	}

	s.D.Set("replication_method", s.Res.ReplicationMethod)

	if s.Res.ReplicationUnit != nil {
		s.D.Set("replication_unit", *s.Res.ReplicationUnit)
	}

	if s.Res.ScanListenerPort != nil {
		s.D.Set("scan_listener_port", *s.Res.ScanListenerPort)
	}

	if configuredPatchOperationsTargetShardDetails(s.D) {
		// TF_CODE_GEN: TERSI-4920-TOP-23 shard_details is both a create-time input and
		// the service's observed topology. When shard topology is managed through configured
		// patch_operations, keep the configured shard_details in state so refresh does not
		// turn a successful INSERT/REMOVE PATCH into a perpetual recreate diff.
		s.D.Set("shard_details", s.D.Get("shard_details"))
	} else {
		shardDetails := []interface{}{}
		for _, item := range s.Res.ShardDetails {
			shardDetails = append(shardDetails, DistributedDatabaseShardToMap(item))
		}
		shardDetails = rehydrateConfiguredAdminPasswords(s.D.Get("shard_details"), shardDetails)
		s.D.Set("shard_details", shardDetails)
	}

	s.D.Set("sharding_method", s.Res.ShardingMethod)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) StartDistributedDatabase(ctx context.Context) error {
	request := oci_distributed_database.StartDistributedDatabaseRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	_, err := s.Client.StartDistributedDatabase(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return s.Res != nil &&
			(s.Res.LifecycleState == oci_distributed_database.DistributedDatabaseLifecycleStateActive ||
				s.Res.LifecycleState == oci_distributed_database.DistributedDatabaseLifecycleStateNeedsAttention ||
				s.Res.LifecycleState == oci_distributed_database.DistributedDatabaseLifecycleStateInactive)
	}
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) StopDistributedDatabase(ctx context.Context) error {
	request := oci_distributed_database.StopDistributedDatabaseRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	_, err := s.Client.StopDistributedDatabase(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return s.Res != nil &&
			(s.Res.LifecycleState == oci_distributed_database.DistributedDatabaseLifecycleStateActive ||
				s.Res.LifecycleState == oci_distributed_database.DistributedDatabaseLifecycleStateNeedsAttention ||
				s.Res.LifecycleState == oci_distributed_database.DistributedDatabaseLifecycleStateInactive)
	}
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

// NOTE (JIRA: TOP-9495):
// ChangeDistributedDbBackupConfig is asynchronous and returns opc-work-request-id.
// Provider must poll the work request to completion before returning, otherwise Terraform
// may proceed while the resource is still UPDATING / WR IN_PROGRESS.

// NOTE (JIRA: TOP-9YYY):
// After WR completion, refresh the resource and accept terminal lifecycle states:
// ACTIVE, INACTIVE, NEEDS_ATTENTION (final state depends on previous power state).
func (s *DistributedDatabaseDistributedDatabaseResourceCrud) ChangeDistributedDbBackupConfig(ctx context.Context) error {
	request := oci_distributed_database.ChangeDistributedDbBackupConfigRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	// NOTE (JIRA: TOP-9494):
	// Action API payload must come from action-specific schema field, not create-time db_backup_config.
	if v, ok := s.D.GetOkExists("db_backup_config"); ok {
		if tmpList := v.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_backup_config", 0)
			tmp, err := s.mapToDistributedDbBackupConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DbBackupConfig = &tmp
		}
	} else {
		// If the trigger was bumped but payload is missing, fail fast (prevents empty-body calls)
		return fmt.Errorf("db_backup_config must be set when change_db_backup_config_trigger is used")
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")
	log.Printf("------------------------- >>>>>>>>>>>>>> [DEBUG] ChangeDistributedDbBackupConfigRequest: %#v", request)
	resp, err := s.Client.ChangeDistributedDbBackupConfig(ctx, request)
	if err != nil {
		return err
	}

	// Prefer WR polling because service transitions through UPDATING and may not reflect immediately via GET.
	workId := resp.OpcWorkRequestId
	if workId != nil && *workId != "" {
		if err := s.getDistributedDatabaseFromWorkRequest(
			ctx,
			workId,
			tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"),
			oci_distributed_database.ActionTypeUpdated,
			s.D.Timeout(schema.TimeoutUpdate),
		); err != nil {
			return err
		}
	}

	// Refresh after WR completion
	if err := s.Get(); err != nil {
		return err
	}

	// Validate final lifecycle after WR is done (could be ACTIVE or INACTIVE depending on prior state).
	if s.Res == nil {
		return fmt.Errorf("resource read returned nil after backup config change")
	}
	switch s.Res.LifecycleState {
	case oci_distributed_database.DistributedDatabaseLifecycleStateActive,
		oci_distributed_database.DistributedDatabaseLifecycleStateInactive,
		oci_distributed_database.DistributedDatabaseLifecycleStateNeedsAttention:
		// ok
	default:
		return fmt.Errorf("unexpected lifecycle state after ChangeDistributedDbBackupConfig: %s", s.Res.LifecycleState)
	}

	// Preserve trigger semantics in Terraform state.
	val := s.D.Get("change_db_backup_config_trigger")
	_ = s.D.Set("change_db_backup_config_trigger", val)

	return nil
}

// NOTE (WORKREQUEST + STATE SETTLE):
// ConfigureSharding is asynchronous. We must wait for the Opc-Work-Request-Id to complete,
// then refresh the DADB and wait for a stable lifecycle state.
// For Distributed Autonomous DB, INACTIVE is a valid steady state after create and after some actions,
// so the acceptable post-action lifecycle set is {ACTIVE, INACTIVE, NEEDS_ATTENTION}.
// See JIRA: TOP-9491
func (s *DistributedDatabaseDistributedDatabaseResourceCrud) ConfigureDistributedDatabaseSharding(ctx context.Context) error {
	request := oci_distributed_database.ConfigureDistributedDatabaseShardingRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	// NOTE (CODEGEN GAP):
	// ConfigureSharding supports request param IsRebalanceRequired, but codegen did not
	// expose it in Terraform schema. We surface it as `configure_sharding_is_rebalance_required`.
	// See JIRA: TOP-9490
	if v, ok := s.D.GetOkExists("configure_sharding_is_rebalance_required"); ok {
		tmp := v.(bool)
		request.IsRebalanceRequired = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	resp, err := s.Client.ConfigureDistributedDatabaseSharding(ctx, request)
	if err != nil {
		return err
	}

	workId := resp.OpcWorkRequestId
	if workId == nil || *workId == "" {
		return fmt.Errorf("missing opc-work-request-id for ConfigureSharding action")
	}

	// Wait for WR completion (Updated action type is consistent for action APIs)
	// TF_CODE_GEN: TERSI-4920-TOP-17 ConfigureSharding work requests can finish SUCCEEDED without
	// a matching distributeddatabase resource/action entry in resources[]; trust terminal WR status
	// and refresh the existing OCID instead of treating metadata drift as a failed action.
	if err := waitForDistributedDatabaseWorkRequestCompletion(
		ctx,
		workId,
		"distributeddatabase",
		oci_distributed_database.ActionTypeCreated,
		s.D.Timeout(schema.TimeoutUpdate),
		s.DisableNotFoundRetries,
		s.WorkRequestClient,
	); err != nil {
		return fmt.Errorf("ConfigureDistributedDatabaseSharding failed for distributed database %s (work request %s): %w", s.D.Id(), *workId, err)
	}

	// After WR is completed, refresh resource and ensure it reaches a stable state.
	// ConfigureSharding can legitimately leave the DADB in ACTIVE, INACTIVE, or NEEDS_ATTENTION.
	stable := func() bool {
		if err := s.Get(); err != nil {
			log.Printf("[WARN] post-WR Get() failed: %v", err)
			return false
		}
		if s.Res == nil {
			return false
		}
		st := s.Res.LifecycleState
		return st == oci_distributed_database.DistributedDatabaseLifecycleStateActive ||
			st == oci_distributed_database.DistributedDatabaseLifecycleStateInactive ||
			st == oci_distributed_database.DistributedDatabaseLifecycleStateNeedsAttention
	}

	if err := tfresource.WaitForResourceCondition(s, stable, s.D.Timeout(schema.TimeoutUpdate)); err != nil {
		return err
	}

	// Preserve trigger semantics in state.
	val := s.D.Get("configure_sharding_trigger")
	_ = s.D.Set("configure_sharding_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) DownloadDistributedDatabaseGsmCertificateSigningRequest() error {
	request := oci_distributed_database.DownloadDistributedDatabaseGsmCertificateSigningRequestRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	/*response, err := s.Client.DownloadDistributedDatabaseGsmCertificateSigningRequest(context.Background(), request)
	if err != nil {
		return err
	}*/

	// WORKAROUND FOR GENERATED CODE ISSUE:
	// This is an action-style operation. The OCI Go SDK response does not include a
	// DistributedAutonomousDatabase model, so we cannot populate s.Res from response.
	// The correct behavior is to invoke the action and rely on WaitForUpdatedState + subsequent reads.
	// See JIRA: TOP-9400
	/*
		response, err := s.Client.DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest(context.Background(), request)
		if err != nil {
			return err
		}

		if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
			return waitErr
		}

		val := s.D.Get("download_gsm_certificate_signing_request_trigger")
		s.D.Set("download_gsm_certificate_signing_request_trigger", val)

		s.Res = &response.DistributedAutonomousDatabase
		return nil
	*/
	_, err := s.Client.DownloadDistributedDatabaseGsmCertificateSigningRequest(context.Background(), request)
	if err != nil {
		return err
	}

	// NOTE (TOP-9398):
	// The legacy WaitForUpdatedState helper requires the non-context
	// ResourceUpdater interface (Update()), which this CRUD intentionally
	// does not implement. Use the context-aware waiter instead to align
	// with UpdateWithContext-based CRUD implementations.
	/*if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}*/

	if waitErr := tfresource.WaitForUpdatedStateWithContext(context.Background(), s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("download_gsm_certificate_signing_request_trigger")
	s.D.Set("download_gsm_certificate_signing_request_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) GenerateDistributedDatabaseGsmCertificateSigningRequest() error {
	request := oci_distributed_database.GenerateDistributedDatabaseGsmCertificateSigningRequestRequest{}

	/*if caBundleId, ok := s.D.GetOkExists("ca_bundle_id"); ok {
		tmp := caBundleId.(string)
		request.CaBundleId = &tmp
	}*/

	// WORKAROUND FOR CODEGEN LIMITATION:
	// The generateGsmCertificateSigningRequest action supports an optional
	// caBundleId query parameter, which is not generated into the Terraform
	// schema by default.
	//
	// This parameter is required to support CA-specific CSR generation.
	// See JIRA: TOP-9478
	if caBundleId, ok := s.D.GetOkExists("ca_bundle_id"); ok {
		tmp := caBundleId.(string)
		request.CaBundleId = &tmp
	}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	// WORKAROUND FOR GENERATED CODE ISSUE:
	// This is an action-style operation. The OCI Go SDK response does not include a
	// DistributedAutonomousDatabase model, so we cannot populate s.Res from response.
	// Also, the generator emitted an empty error-handling block; return the error.
	// See JIRA: TOP-9400

	/*
		response, err := s.Client.GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest(context.Background(), request)
		if err != nil {

		}

		if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
			return waitErr
		}

		val := s.D.Get("generate_gsm_certificate_signing_request_trigger")
		s.D.Set("generate_gsm_certificate_signing_request_trigger", val)

		s.Res = &response.DistributedAutonomousDatabase
		return nil
	*/

	/*response, err := s.Client.GenerateDistributedDatabaseGsmCertificateSigningRequest(context.Background(), request)
	if err != nil {
		return err
	}*/

	_, err := s.Client.GenerateDistributedDatabaseGsmCertificateSigningRequest(context.Background(), request)
	if err != nil {
		return err
	}

	// NOTE (TOP-9398):
	// The legacy WaitForUpdatedState helper requires the non-context
	// ResourceUpdater interface (Update()), which this CRUD intentionally
	// does not implement. Use the context-aware waiter instead to align
	// with UpdateWithContext-based CRUD implementations.

	/*if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}*/

	if waitErr := tfresource.WaitForUpdatedStateWithContext(context.Background(), s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("generate_gsm_certificate_signing_request_trigger")
	s.D.Set("generate_gsm_certificate_signing_request_trigger", val)

	//s.Res = &response.DistributedDatabase
	return nil
}

/*func (s *DistributedDatabaseDistributedDatabaseResourceCrud) GenerateDistributedDatabaseWallet() error {
request := oci_distributed_database.GenerateDistributedDatabaseWalletRequest{}

idTmp := s.D.Id()
request.DistributedDatabaseId = &idTmp

if password, ok := s.D.GetOkExists("password"); ok {
	tmp := password.(string)
	request.Password = &tmp
}

request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

// WORKAROUND FOR GENERATED CODE ISSUE:
// This is an action-style operation. The OCI Go SDK response does not include a
// DistributedAutonomousDatabase model, so we cannot populate s.Res from response.
// Follow the established provider pattern: invoke action, wait for update, preserve trigger.
// See JIRA: TOP-9400

/*
	response, err := s.Client.GenerateDistributedAutonomousDatabaseWallet(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("generate_wallet_trigger")
	s.D.Set("generate_wallet_trigger", val)

	s.Res = &response.DistributedAutonomousDatabase
	return nil
*/

/*	response, err := s.Client.GenerateDistributedDatabaseWallet(context.Background(), request)
	if err != nil {
		return err
	}*/

/*_, err := s.Client.GenerateDistributedDatabaseWallet(context.Background(), request)
if err != nil {
	return err
}

// NOTE (TOP-9398):
// The legacy WaitForUpdatedState helper requires the non-context
// ResourceUpdater interface (Update()), which this CRUD intentionally
// does not implement. Use the context-aware waiter instead to align
// with UpdateWithContext-based CRUD implementations.
/*
	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}*/

/*if waitErr := tfresource.WaitForUpdatedStateWithContext(context.Background(), s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("generate_wallet_trigger")
	s.D.Set("generate_wallet_trigger", val)

	//s.Res = &response.DistributedDatabase
	return nil
}*/
// GenerateDistributedAutonomousDatabaseWallet calls the GenerateWallet action API and stores the
// returned wallet zip (binary body) in Terraform state as base64 so users can write it to disk
// using local_file.content_base64.
//
// NOTE (JIRA: TOP-9470):
// The code generator exposes generate_wallet_trigger but omits the required password parameter
// needed by the GenerateWallet API. We add generate_wallet_password to allow users to provide it.
//
// NOTE (WALLET DOWNLOAD SEMANTICS):
// This API returns the wallet as a binary stream (io.ReadCloser) and does NOT return a WorkRequestId.
// There is no lifecycle transition to poll, so we must not call WaitForUpdatedState*.
// We read/close the stream and persist it in state.
// TOP-9492
func (s *DistributedDatabaseDistributedDatabaseResourceCrud) GenerateDistributedDatabaseWallet(ctx context.Context) error {
	request := oci_distributed_database.GenerateDistributedDatabaseWalletRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	// Password is required by the action details; exposed via generate_wallet_password (Sensitive).
	if password, ok := s.D.GetOkExists("generate_wallet_password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	resp, err := s.Client.GenerateDistributedDatabaseWallet(ctx, request)
	if err != nil {
		return err
	}
	defer func() {
		if resp.Content != nil {
			_ = resp.Content.Close()
		}
	}()

	// Read binary zip content and store as base64 (Sensitive + Computed)
	if resp.Content != nil {
		b, readErr := io.ReadAll(resp.Content)
		if readErr != nil {
			return readErr
		}
		encoded := base64.StdEncoding.EncodeToString(b)
		if setErr := s.D.Set("generate_wallet_downloaded_wallet_zip_base64", encoded); setErr != nil {
			return setErr
		}
	}

	// Store optional response headers as computed attributes
	if resp.Etag != nil {
		_ = s.D.Set("generate_wallet_downloaded_wallet_etag", *resp.Etag)
	}
	if resp.LastModified != nil {
		_ = s.D.Set("generate_wallet_downloaded_wallet_last_modified", resp.LastModified.String())
	}
	if resp.ContentLength != nil {
		_ = s.D.Set("generate_wallet_downloaded_wallet_content_length", int(*resp.ContentLength))
	}

	// Preserve trigger semantics in Terraform state.
	val := s.D.Get("generate_wallet_trigger")
	_ = s.D.Set("generate_wallet_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) MoveDistributedDatabaseReplicationUnit(ctx context.Context) error {
	request := oci_distributed_database.MoveDistributedDatabaseReplicationUnitRequest{}

	if destinationShardName, ok := s.D.GetOkExists("destination_shard_name"); ok {
		tmp := destinationShardName.(string)
		request.DestinationShardName = &tmp
	}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

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

	// TF_CODE_GEN: TERSI-4920-TOP-18 generated action helpers treat metadata-only SDK responses as if they returned a DistributedDatabase payload; refresh with Get() after waiting instead.
	_, err := s.Client.MoveDistributedDatabaseReplicationUnit(ctx, request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("move_replication_unit_trigger")
	s.D.Set("move_replication_unit_trigger", val)

	return s.Get()
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) RecreateFailedDistributedDatabaseResource(ctx context.Context) error {
	request := oci_distributed_database.RecreateFailedDistributedDatabaseResourceRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	if shardGroup, ok := s.D.GetOkExists("shard_group"); ok {
		tmp := shardGroup.(string)
		request.ShardGroup = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	// TF_CODE_GEN: TERSI-4920-TOP-18 generated action helpers treat metadata-only SDK responses as if they returned a DistributedDatabase payload; refresh with Get() after waiting instead.
	_, err := s.Client.RecreateFailedDistributedDatabaseResource(ctx, request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("recreate_failed_resource_trigger")
	s.D.Set("recreate_failed_resource_trigger", val)

	return s.Get()
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) UploadDistributedDatabaseSignedCertificateAndGenerateWallet() error {
	request := oci_distributed_database.UploadDistributedDatabaseSignedCertificateAndGenerateWalletRequest{}

	if caSignedCertificate, ok := s.D.GetOkExists("ca_signed_certificate"); ok {
		tmp := caSignedCertificate.(string)
		request.CaSignedCertificate = &tmp
	}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	// WORKAROUND FOR GENERATED CODE ISSUE:
	// This is an action-style operation. The OCI Go SDK response does not include a
	// DistributedAutonomousDatabase model, so we cannot populate s.Res from response.
	// Follow the established provider pattern: invoke action, wait for update, preserve trigger.
	// See JIRA: TOP-9400
	/*
		response, err := s.Client.UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet(context.Background(), request)
		if err != nil {
			return err
		}

		if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
			return waitErr
		}

		val := s.D.Get("upload_signed_certificate_and_generate_wallet_trigger")
		s.D.Set("upload_signed_certificate_and_generate_wallet_trigger", val)

		s.Res = &response.DistributedAutonomousDatabase
		return nil*/

	/*response, err := s.Client.UploadDistributedDatabaseSignedCertificateAndGenerateWallet(context.Background(), request)
	if err != nil {
		return err
	}*/

	_, err := s.Client.UploadDistributedDatabaseSignedCertificateAndGenerateWallet(context.Background(), request)
	if err != nil {
		return err
	}

	// NOTE (TOP-9398):
	// The legacy WaitForUpdatedState helper requires the non-context
	// ResourceUpdater interface (Update()), which this CRUD intentionally
	// does not implement. Use the context-aware waiter instead to align
	// with UpdateWithContext-based CRUD implementations.
	/*
		if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
			return waitErr
		}*/

	if waitErr := tfresource.WaitForUpdatedStateWithContext(context.Background(), s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("upload_signed_certificate_and_generate_wallet_trigger")
	s.D.Set("upload_signed_certificate_and_generate_wallet_trigger", val)

	//s.Res = &response.DistributedDatabase
	return nil
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) ValidateDistributedDatabaseNetwork(ctx context.Context) error {
	request := oci_distributed_database.ValidateDistributedDatabaseNetworkRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp
	/*
		if isSurrogate, ok := s.D.GetOkExists("is_surrogate"); ok {
			tmp := isSurrogate.(bool)
			request.IsSurrogate = &tmp
		}

		if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
			tmp := resourceName.(string)
			request.ResourceName = &tmp
		}

		if shardGroup, ok := s.D.GetOkExists("shard_group"); ok {
			tmp := shardGroup.(string)
			request.ShardGroup = &tmp
		}
	*/
	if v, ok := s.D.GetOkExists("validate_network_details"); ok {
		l := v.([]interface{})
		if len(l) > 0 && l[0] != nil {
			m := l[0].(map[string]interface{})

			if vv, ok := m["is_surrogate"]; ok {
				b := vv.(bool)
				request.IsSurrogate = &b
			}
			if vv, ok := m["resource_name"]; ok && vv.(string) != "" {
				str := vv.(string)
				request.ResourceName = &str
			}
			if vv, ok := m["shard_group"]; ok && vv.(string) != "" {
				str := vv.(string)
				request.ShardGroup = &str
			}
		}
	}
	//request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	// WORKAROUND FOR GENERATED CODE ISSUE:
	// This is an action-style operation. The OCI Go SDK response does not include a
	// DistributedAutonomousDatabase model, so we cannot populate s.Res from response.
	// Follow the established provider pattern: invoke action, wait for lifecycle stabilization,
	// then WaitForUpdatedState and preserve trigger semantics.
	// See JIRA: TOP-9400
	/*
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

		response, err := s.Client.ValidateDistributedAutonomousDatabaseNetwork(context.Background(), request)
		if err != nil {
			return err
		}

		if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
			return waitErr
		}

		val := s.D.Get("validate_network_trigger")
		s.D.Set("validate_network_trigger", val)

		s.Res = &response.DistributedAutonomousDatabase
		return nil
	*/

	/*
			response, err := s.Client.ValidateDistributedDatabaseNetwork(context.Background(), request)
		if err != nil {
			return err
		}
	*/
	_, err := s.Client.ValidateDistributedDatabaseNetwork(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		if err := s.Get(); err != nil { // Refresh status
			log.Printf("[WARN] Failed to refresh resource during wait: %v", err)
			return false
		}
		return s.Res != nil &&
			(s.Res.LifecycleState == oci_distributed_database.DistributedDatabaseLifecycleStateActive ||
				s.Res.LifecycleState == oci_distributed_database.DistributedDatabaseLifecycleStateInactive ||
				s.Res.LifecycleState == oci_distributed_database.DistributedDatabaseLifecycleStateNeedsAttention)
	}

	// NOTE (TOP-9398):
	// The legacy WaitForUpdatedState helper requires the non-context
	// ResourceUpdater interface (Update()), which this CRUD intentionally
	// does not implement. Use the context-aware waiter instead to align
	// with UpdateWithContext-based CRUD implementations.
	/*
		if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
			return waitErr
		}*/

	if err := tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate)); err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("validate_network_trigger")
	s.D.Set("validate_network_trigger", val)

	//s.Res = &response.DistributedDatabase
	return nil
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateCatalogPeerWithExadbXsDetails(fieldKeyFormat string) (oci_distributed_database.CreateCatalogPeerWithExadbXsDetails, error) {
	result := oci_distributed_database.CreateCatalogPeerWithExadbXsDetails{}

	if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
		result.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
	}

	if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
		result.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
	}

	if vmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_id")); ok {
		tmp := vmClusterId.(string)
		result.VmClusterId = &tmp
	}

	return result, nil
}

func CreateCatalogPeerWithExadbXsDetailsToMap(obj oci_distributed_database.CreateCatalogPeerWithExadbXsDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["protection_mode"] = string(obj.ProtectionMode)

	result["transport_type"] = string(obj.TransportType)

	if obj.VmClusterId != nil {
		result["vm_cluster_id"] = string(*obj.VmClusterId)
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateCatalogPeerWithExadbXsNewVaultAndClusterDetails(fieldKeyFormat string) (oci_distributed_database.CreateCatalogPeerWithExadbXsNewVaultAndClusterDetails, error) {
	result := oci_distributed_database.CreateCatalogPeerWithExadbXsNewVaultAndClusterDetails{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if dbStorageVaultDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_storage_vault_details")); ok {
		if tmpList := dbStorageVaultDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "db_storage_vault_details"), 0)
			tmp, err := s.mapToDbStorageVaultDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert db_storage_vault_details, encountered error: %v", err)
			}
			result.DbStorageVaultDetails = &tmp
		}
	}

	if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
		result.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
	}

	if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
		result.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
	}

	if vmClusterDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_details")); ok {
		if tmpList := vmClusterDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "vm_cluster_details"), 0)
			tmp, err := s.mapToVmClusterDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert vm_cluster_details, encountered error: %v", err)
			}
			result.VmClusterDetails = &tmp
		}
	}

	return result, nil
}

func CreateCatalogPeerWithExadbXsNewVaultAndClusterDetailsToMap(obj oci_distributed_database.CreateCatalogPeerWithExadbXsNewVaultAndClusterDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.DbStorageVaultDetails != nil {
		result["db_storage_vault_details"] = []interface{}{DbStorageVaultDetailsToMap(obj.DbStorageVaultDetails)}
	}

	result["protection_mode"] = string(obj.ProtectionMode)

	result["transport_type"] = string(obj.TransportType)

	if obj.VmClusterDetails != nil {
		result["vm_cluster_details"] = []interface{}{VmClusterDetailsToMap(obj.VmClusterDetails, false)}
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateDistributedDatabaseCatalogDetails(fieldKeyFormat string) (oci_distributed_database.CreateDistributedDatabaseCatalogDetails, error) {
	var baseObject oci_distributed_database.CreateDistributedDatabaseCatalogDetails
	//discriminator
	sourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source"))
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("EXADB_XS"):
		details := oci_distributed_database.CreateDistributedDatabaseCatalogWithExadbXsDetails{}
		if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if peerDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_details")); ok {
			interfaces := peerDetails.([]interface{})
			tmp := make([]oci_distributed_database.CreateCatalogPeerWithExadbXsDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "peer_details"), stateDataIndex)
				converted, err := s.mapToCreateCatalogPeerWithExadbXsDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "peer_details")) {
				details.PeerDetails = tmp
			}
		}
		if peerVmClusterIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_vm_cluster_ids")); ok {
			interfaces := peerVmClusterIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "peer_vm_cluster_ids")) {
				details.PeerVmClusterIds = tmp
			}
		}
		if shardSpace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shard_space")); ok {
			tmp := shardSpace.(string)
			details.ShardSpace = &tmp
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if vmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_id")); ok {
			tmp := vmClusterId.(string)
			details.VmClusterId = &tmp
		}
		baseObject = details
	case strings.ToLower("NEW_VAULT_AND_CLUSTER"):
		details := oci_distributed_database.CreateDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterDetails{}
		if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
		}
		if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		if dbStorageVaultDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_storage_vault_details")); ok {
			if tmpList := dbStorageVaultDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "db_storage_vault_details"), 0)
				tmp, err := s.mapToDbStorageVaultDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert db_storage_vault_details, encountered error: %v", err)
				}
				details.DbStorageVaultDetails = &tmp
			}
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if peerDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_details")); ok {
			interfaces := peerDetails.([]interface{})
			tmp := make([]oci_distributed_database.CreateCatalogPeerWithExadbXsNewVaultAndClusterDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "peer_details"), stateDataIndex)
				converted, err := s.mapToCreateCatalogPeerWithExadbXsNewVaultAndClusterDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "peer_details")) {
				details.PeerDetails = tmp
			}
		}
		if shardSpace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shard_space")); ok {
			tmp := shardSpace.(string)
			details.ShardSpace = &tmp
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if vmClusterDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_details")); ok {
			if tmpList := vmClusterDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "vm_cluster_details"), 0)
				tmp, err := s.mapToVmClusterDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert vm_cluster_details, encountered error: %v", err)
				}
				details.VmClusterDetails = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source '%v' was specified", source)
	}
	return baseObject, nil
}

func DistributedDatabaseCatalogToMap(obj oci_distributed_database.DistributedDatabaseCatalog) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_distributed_database.DistributedDatabaseCatalogWithExadbXs:
		result["source"] = "EXADB_XS"

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		peerDetails := []interface{}{}
		for _, item := range v.PeerDetails {
			peerDetails = append(peerDetails, CatalogPeerWithExadbXsToMap(item))
		}
		result["peer_details"] = peerDetails

		if v.ShardGroup != nil {
			result["shard_group"] = string(*v.ShardGroup)
		}

		if v.DbHomeId != nil {
			result["db_home_id"] = string(*v.DbHomeId)
		}

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.VmClusterId != nil {
			result["vm_cluster_id"] = string(*v.VmClusterId)
		}

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = string(*v.SupportingResourceId)
		}

		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = string(*v.ContainerDatabaseId)
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
		result["status"] = string(v.Status)
	default:
		log.Printf("[WARN] Received 'source' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateDistributedDatabaseShardDetails(fieldKeyFormat string) (oci_distributed_database.CreateDistributedDatabaseShardDetails, error) {
	var baseObject oci_distributed_database.CreateDistributedDatabaseShardDetails
	//discriminator
	sourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source"))
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("EXADB_XS"):
		details := oci_distributed_database.CreateDistributedDatabaseShardWithExadbXsDetails{}
		if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if peerDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_details")); ok {
			interfaces := peerDetails.([]interface{})
			tmp := make([]oci_distributed_database.CreateShardPeerWithExadbXsDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "peer_details"), stateDataIndex)
				converted, err := s.mapToCreateShardPeerWithExadbXsDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "peer_details")) {
				details.PeerDetails = tmp
			}
		}
		if peerVmClusterIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_vm_cluster_ids")); ok {
			interfaces := peerVmClusterIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "peer_vm_cluster_ids")) {
				details.PeerVmClusterIds = tmp
			}
		}
		if shardSpace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shard_space")); ok {
			tmp := shardSpace.(string)
			details.ShardSpace = &tmp
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if vmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_id")); ok {
			tmp := vmClusterId.(string)
			details.VmClusterId = &tmp
		}
		baseObject = details
	case strings.ToLower("NEW_VAULT_AND_CLUSTER"):
		details := oci_distributed_database.CreateDistributedDatabaseShardWithExadbXsNewVaultAndClusterDetails{}
		if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
		}
		if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		if dbStorageVaultDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_storage_vault_details")); ok {
			if tmpList := dbStorageVaultDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "db_storage_vault_details"), 0)
				tmp, err := s.mapToDbStorageVaultDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert db_storage_vault_details, encountered error: %v", err)
				}
				details.DbStorageVaultDetails = &tmp
			}
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if peerDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_details")); ok {
			interfaces := peerDetails.([]interface{})
			tmp := make([]oci_distributed_database.CreateShardPeerWithExadbXsNewVaultAndClusterDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "peer_details"), stateDataIndex)
				converted, err := s.mapToCreateShardPeerWithExadbXsNewVaultAndClusterDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "peer_details")) {
				details.PeerDetails = tmp
			}
		}
		if shardSpace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shard_space")); ok {
			tmp := shardSpace.(string)
			details.ShardSpace = &tmp
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if vmClusterDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_details")); ok {
			if tmpList := vmClusterDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "vm_cluster_details"), 0)
				tmp, err := s.mapToVmClusterDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert vm_cluster_details, encountered error: %v", err)
				}
				details.VmClusterDetails = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source '%v' was specified", source)
	}
	return baseObject, nil
}

func DistributedDatabaseShardToMap(obj oci_distributed_database.DistributedDatabaseShard) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The generator incorrectly emits a type-switch case for the request-side
	// CreateDistributedDatabaseShardWithDedicatedInfraDetails, but the
	// interface here is the response-side DistributedDatabaseShard.
	// The Create*Details type can never be a dynamic type of this interface
	// (it doesn't implement GetName), causing an "impossible type switch case"
	// and follow-on missing field errors (Name/TimeCreated/TimeUpdated).
	//
	// Use the SDK model type returned by GET/List instead.
	// See JIRA: TOP-9407

	//case oci_distributed_database.CreateDistributedDatabaseShardWithExadbXsDetails:
	case oci_distributed_database.DistributedDatabaseShardWithExadbXs:
		result["source"] = "EXADB_XS"

		/*
			if v.AdminPassword != nil {
				result["admin_password"] = string(*v.AdminPassword)
			}*/

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		/*peerDetails := []interface{}{}
		for _, item := range v.PeerDetails {
			peerDetails = append(peerDetails, CreateShardPeerWithExadbXsDetailsToMap(item))
		}
		result["peer_details"] = peerDetails

		result["peer_vm_cluster_ids"] = v.PeerVmClusterIds*/

		// WORKAROUND FOR GENERATED CODE ISSUE:
		// Read/ToMap must operate on SDK model types, not Create*Details request types.
		// The generator incorrectly reused CreateShardPeerWithExadbXsDetailsToMap,
		// which expects a request-side struct, but v.PeerDetails contains the
		// response-side SDK model type ShardPeerWithExadbXs.
		// See JIRA: <JIRA-KEY>
		peerDetails := []interface{}{}
		for _, item := range v.PeerDetails {
			peerDetails = append(peerDetails, ShardPeerWithExadbXsToMap(item))
		}
		result["peer_details"] = peerDetails

		if v.ShardSpace != nil {
			result["shard_space"] = string(*v.ShardSpace)
		}

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.VmClusterId != nil {
			result["vm_cluster_id"] = string(*v.VmClusterId)
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

		if v.ShardGroup != nil {
			result["shard_group"] = string(*v.ShardGroup)
		}

		if v.DbHomeId != nil {
			result["db_home_id"] = string(*v.DbHomeId)
		}

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = string(*v.SupportingResourceId)
		}

		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = string(*v.ContainerDatabaseId)
		}

		result["status"] = string(v.Status)
	default:
		log.Printf("[WARN] Received 'source' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateShardPeerWithExadbXsDetails(fieldKeyFormat string) (oci_distributed_database.CreateShardPeerWithExadbXsDetails, error) {
	result := oci_distributed_database.CreateShardPeerWithExadbXsDetails{}

	if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
		result.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
	}

	if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
		result.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
	}

	if vmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_id")); ok {
		tmp := vmClusterId.(string)
		result.VmClusterId = &tmp
	}

	return result, nil
}

func CreateShardPeerWithExadbXsDetailsToMap(obj oci_distributed_database.CreateShardPeerWithExadbXsDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["protection_mode"] = string(obj.ProtectionMode)

	result["transport_type"] = string(obj.TransportType)

	if obj.VmClusterId != nil {
		result["vm_cluster_id"] = string(*obj.VmClusterId)
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateShardPeerWithExadbXsNewVaultAndClusterDetails(fieldKeyFormat string) (oci_distributed_database.CreateShardPeerWithExadbXsNewVaultAndClusterDetails, error) {
	result := oci_distributed_database.CreateShardPeerWithExadbXsNewVaultAndClusterDetails{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if dbStorageVaultDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_storage_vault_details")); ok {
		if tmpList := dbStorageVaultDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "db_storage_vault_details"), 0)
			tmp, err := s.mapToDbStorageVaultDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert db_storage_vault_details, encountered error: %v", err)
			}
			result.DbStorageVaultDetails = &tmp
		}
	}

	if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
		result.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
	}

	if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
		result.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
	}

	if vmClusterDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_details")); ok {
		if tmpList := vmClusterDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "vm_cluster_details"), 0)
			tmp, err := s.mapToVmClusterDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert vm_cluster_details, encountered error: %v", err)
			}
			result.VmClusterDetails = &tmp
		}
	}

	return result, nil
}

func CreateShardPeerWithExadbXsNewVaultAndClusterDetailsToMap(obj oci_distributed_database.CreateShardPeerWithExadbXsNewVaultAndClusterDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.DbStorageVaultDetails != nil {
		result["db_storage_vault_details"] = []interface{}{DbStorageVaultDetailsToMap(obj.DbStorageVaultDetails)}
	}

	result["protection_mode"] = string(obj.ProtectionMode)

	result["transport_type"] = string(obj.TransportType)

	if obj.VmClusterDetails != nil {
		result["vm_cluster_details"] = []interface{}{VmClusterDetailsToMap(obj.VmClusterDetails, false)}
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToDbStorageVaultDetails(fieldKeyFormat string) (oci_distributed_database.DbStorageVaultDetails, error) {
	result := oci_distributed_database.DbStorageVaultDetails{}

	if additionalFlashCacheInPercent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_flash_cache_in_percent")); ok {
		tmp := additionalFlashCacheInPercent.(int)
		result.AdditionalFlashCacheInPercent = &tmp
	}

	if highCapacityDatabaseStorage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "high_capacity_database_storage")); ok {
		tmp := highCapacityDatabaseStorage.(int)
		result.HighCapacityDatabaseStorage = &tmp
	}

	return result, nil
}

func DbStorageVaultDetailsToMap(obj *oci_distributed_database.DbStorageVaultDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdditionalFlashCacheInPercent != nil {
		result["additional_flash_cache_in_percent"] = int(*obj.AdditionalFlashCacheInPercent)
	}

	if obj.HighCapacityDatabaseStorage != nil {
		result["high_capacity_database_storage"] = int(*obj.HighCapacityDatabaseStorage)
	}

	return result
}

func DistributedDatabaseGsmToMap(obj oci_distributed_database.DistributedDatabaseGsm) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComputeCount != nil {
		result["compute_count"] = float32(*obj.ComputeCount)
	}

	if obj.DataStorageSizeInGbs != nil {
		result["data_storage_size_in_gbs"] = float64(*obj.DataStorageSizeInGbs)
	}

	if obj.GsmImageDetails != nil {
		result["gsm_image_details"] = []interface{}{DistributedDbGsmImageToMap(obj.GsmImageDetails)}
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{DistributedDbMetadataToMap(obj.Metadata)}
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

func DistributedDatabaseSummaryToMap(obj oci_distributed_database.DistributedDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CharacterSet != nil {
		result["character_set"] = string(*obj.CharacterSet)
	}

	if obj.Chunks != nil {
		result["chunks"] = int(*obj.Chunks)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConnectionStrings != nil {
		result["connection_strings"] = []interface{}{DistributedDbConnectionStringToMap(obj.ConnectionStrings)}
	}

	if obj.DatabaseVersion != nil {
		result["database_version"] = string(*obj.DatabaseVersion)
	}

	result["db_deployment_type"] = string(obj.DbDeploymentType)

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
		result["metadata"] = []interface{}{DistributedDbMetadataToMap(obj.Metadata)}
	}

	if obj.NcharacterSet != nil {
		result["ncharacter_set"] = string(*obj.NcharacterSet)
	}

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

	if obj.ReplicationFactor != nil {
		result["replication_factor"] = int(*obj.ReplicationFactor)
	}

	result["replication_method"] = string(obj.ReplicationMethod)

	if obj.ReplicationUnit != nil {
		result["replication_unit"] = int(*obj.ReplicationUnit)
	}

	result["sharding_method"] = string(obj.ShardingMethod)

	result["state"] = string(obj.LifecycleState)

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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToDistributedDbBackupConfig(fieldKeyFormat string) (oci_distributed_database.DistributedDbBackupConfig, error) {
	result := oci_distributed_database.DistributedDbBackupConfig{}

	if autoBackupWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_window")); ok {
		result.AutoBackupWindow = oci_distributed_database.DistributedDbBackupConfigAutoBackupWindowEnum(autoBackupWindow.(string))
	}

	if autoFullBackupDay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_full_backup_day")); ok {
		result.AutoFullBackupDay = oci_distributed_database.DistributedDbBackupConfigAutoFullBackupDayEnum(autoFullBackupDay.(string))
	}

	if autoFullBackupWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_full_backup_window")); ok {
		result.AutoFullBackupWindow = oci_distributed_database.DistributedDbBackupConfigAutoFullBackupWindowEnum(autoFullBackupWindow.(string))
	}

	if backupDeletionPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_deletion_policy")); ok {
		result.BackupDeletionPolicy = oci_distributed_database.DistributedDbBackupConfigBackupDeletionPolicyEnum(backupDeletionPolicy.(string))
	}

	if backupDestinationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")); ok {
		interfaces := backupDestinationDetails.([]interface{})
		tmp := make([]oci_distributed_database.DistributedDbBackupDestination, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backup_destination_details"), stateDataIndex)
			converted, err := s.mapToDistributedDbBackupDestination(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")) {
			result.BackupDestinationDetails = tmp
		}
	}

	if canRunImmediateFullBackup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "can_run_immediate_full_backup")); ok {
		tmp := canRunImmediateFullBackup.(bool)
		result.CanRunImmediateFullBackup = &tmp
	}

	if isAutoBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_backup_enabled")); ok {
		tmp := isAutoBackupEnabled.(bool)
		result.IsAutoBackupEnabled = &tmp
	}

	if isRemoteBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_remote_backup_enabled")); ok {
		tmp := isRemoteBackupEnabled.(bool)
		result.IsRemoteBackupEnabled = &tmp
	}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	if remoteRegion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remote_region")); ok {
		tmp := remoteRegion.(string)
		result.RemoteRegion = &tmp
	}

	return result, nil
}

/*
func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToDistributedDbBackupConfig(fieldKeyFormat string) (oci_distributed_database.DistributedDbBackupConfig, error) {
	result := oci_distributed_database.DistributedDbBackupConfig{}

	if backupDestinationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")); ok {
		interfaces := backupDestinationDetails.([]interface{})
		tmp := make([]oci_distributed_database.DistributedDbBackupDestination, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backup_destination_details"), stateDataIndex)
			converted, err := s.mapToDistributedDbBackupDestination(fieldKeyFormatNextLevel)
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
*/

func DistributedDbBackupConfigToMap(obj *oci_distributed_database.DistributedDbBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["auto_backup_window"] = string(obj.AutoBackupWindow)

	result["auto_full_backup_day"] = string(obj.AutoFullBackupDay)

	result["auto_full_backup_window"] = string(obj.AutoFullBackupWindow)

	result["backup_deletion_policy"] = string(obj.BackupDeletionPolicy)

	backupDestinationDetails := []interface{}{}
	for _, item := range obj.BackupDestinationDetails {
		backupDestinationDetails = append(backupDestinationDetails, DistributedDbBackupDestinationToMap(item))
	}
	result["backup_destination_details"] = backupDestinationDetails

	if obj.CanRunImmediateFullBackup != nil {
		result["can_run_immediate_full_backup"] = bool(*obj.CanRunImmediateFullBackup)
	}

	if obj.IsAutoBackupEnabled != nil {
		result["is_auto_backup_enabled"] = bool(*obj.IsAutoBackupEnabled)
	}

	if obj.IsRemoteBackupEnabled != nil {
		result["is_remote_backup_enabled"] = bool(*obj.IsRemoteBackupEnabled)
	}

	if obj.RecoveryWindowInDays != nil {
		result["recovery_window_in_days"] = int(*obj.RecoveryWindowInDays)
	}

	if obj.RemoteRegion != nil {
		result["remote_region"] = string(*obj.RemoteRegion)
	}

	return result
}

func fieldKeyToRawConfigPath(fieldKey string) (cty.Path, error) {
	parts := strings.Split(fieldKey, ".")
	path := make(cty.Path, 0, len(parts))
	for _, part := range parts {
		if index, err := strconv.Atoi(part); err == nil {
			path = append(path, cty.IndexStep{Key: cty.NumberIntVal(int64(index))})
			continue
		}
		path = append(path, cty.GetAttrStep{Name: part})
	}
	return path, nil
}

func getConfiguredRawFieldValue(d *schema.ResourceData, fieldKey string) (cty.Value, bool) {
	path, err := fieldKeyToRawConfigPath(fieldKey)
	if err != nil {
		return cty.DynamicVal, false
	}

	value, diags := d.GetRawConfigAt(path)
	if diags.HasError() || !value.IsKnown() || value.IsNull() {
		return cty.DynamicVal, false
	}

	return value, true
}

func getConfiguredStringPointer(d *schema.ResourceData, fieldKey string) *string {
	value, ok := getConfiguredRawFieldValue(d, fieldKey)
	if !ok || value.Type() != cty.String || value.AsString() == "" {
		return nil
	}

	configured := value.AsString()
	return &configured
}

func getConfiguredBoolPointer(d *schema.ResourceData, fieldKey string) *bool {
	value, ok := getConfiguredRawFieldValue(d, fieldKey)
	if !ok || value.Type() != cty.Bool {
		return nil
	}

	configured := d.Get(fieldKey).(bool)
	return &configured
}

/*func DistributedDbBackupConfigToMap(obj *oci_distributed_database.DistributedDbBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

	backupDestinationDetails := []interface{}{}
	for _, item := range obj.BackupDestinationDetails {
		backupDestinationDetails = append(backupDestinationDetails, DistributedDbBackupDestinationToMap(item))
	}
	result["backup_destination_details"] = backupDestinationDetails

	if obj.RecoveryWindowInDays != nil {
		result["recovery_window_in_days"] = int(*obj.RecoveryWindowInDays)
	}

	return result
}*/

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToDistributedDbBackupDestination(fieldKeyFormat string) (oci_distributed_database.DistributedDbBackupDestination, error) {
	result := oci_distributed_database.DistributedDbBackupDestination{}

	// TF_CODE_GEN: TERSI-4920-TOP-18 backup destination fields are Optional+Computed;
	// inspect raw config so ChangeDbBackupConfig only forwards user-supplied values
	// instead of empty computed strings/bools echoed from state.
	if tmp := getConfiguredStringPointer(s.D, fmt.Sprintf(fieldKeyFormat, "dbrs_policy_id")); tmp != nil {
		result.DbrsPolicyId = tmp
	}

	if tmp := getConfiguredStringPointer(s.D, fmt.Sprintf(fieldKeyFormat, "id")); tmp != nil {
		result.Id = tmp
	}

	if tmp := getConfiguredStringPointer(s.D, fmt.Sprintf(fieldKeyFormat, "internet_proxy")); tmp != nil {
		result.InternetProxy = tmp
	}

	if tmp := getConfiguredBoolPointer(s.D, fmt.Sprintf(fieldKeyFormat, "is_remote")); tmp != nil {
		result.IsRemote = tmp
	}

	if tmp := getConfiguredBoolPointer(s.D, fmt.Sprintf(fieldKeyFormat, "is_zero_data_loss_enabled")); tmp != nil {
		result.IsZeroDataLossEnabled = tmp
	}

	if tmp := getConfiguredStringPointer(s.D, fmt.Sprintf(fieldKeyFormat, "remote_region")); tmp != nil {
		result.RemoteRegion = tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_distributed_database.DistributedDbBackupDestinationTypeEnum(type_.(string))
	}

	if tmp := getConfiguredStringPointer(s.D, fmt.Sprintf(fieldKeyFormat, "vpc_password")); tmp != nil {
		result.VpcPassword = tmp
	}

	if tmp := getConfiguredStringPointer(s.D, fmt.Sprintf(fieldKeyFormat, "vpc_user")); tmp != nil {
		result.VpcUser = tmp
	}

	return result, nil
}

func DistributedDbBackupDestinationToMap(obj oci_distributed_database.DistributedDbBackupDestination) map[string]interface{} {
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

	if obj.IsZeroDataLossEnabled != nil {
		result["is_zero_data_loss_enabled"] = bool(*obj.IsZeroDataLossEnabled)
	}

	if obj.RemoteRegion != nil {
		result["remote_region"] = string(*obj.RemoteRegion)
	}

	result["type"] = string(obj.Type)

	if obj.VpcPassword != nil {
		result["vpc_password"] = string(*obj.VpcPassword)
	}

	if obj.VpcUser != nil {
		result["vpc_user"] = string(*obj.VpcUser)
	}

	return result
}

func DistributedDbConnectionStringToMap(obj *oci_distributed_database.DistributedDbConnectionString) map[string]interface{} {
	result := map[string]interface{}{}

	result["all_connection_strings"] = obj.AllConnectionStrings

	return result
}

func DistributedDbGsmImageToMap(obj *oci_distributed_database.DistributedDbGsmImage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.VersionNumber != nil {
		result["version_number"] = int(*obj.VersionNumber)
	}

	return result
}

//func DistributedDbMetadataToMap(obj *oci_distributed_database.DistributedDbMetadata) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	// WORKAROUND FOR GENERATED CODE ISSUE:
//	//
//	// The generator emits calls to mapToDistributedDbMetadata but either
//	// doesn't generate the function and/or assumes the SDK model has a `Map` field.
//	// In OCI Go SDK, DistributedDbMetadata uses `PropertiesMap` (json:"map").
//	//
//	// This mapper converts the Terraform schema field `metadata.0.map` into
//	// `DistributedDbMetadata.PropertiesMap`.
//	//
//	// Remove once generator is fixed.
//	// See JIRA: TOP-9397
//	//result["map"] = obj.Map
//	result["map"] = obj.PropertiesMap
//
//	return result
//}

func DistributedDbMetadataToMap(
	obj *oci_distributed_database.DistributedDbMetadata,
) map[string]interface{} {

	result := map[string]interface{}{}

	// Always return a map with "map" key to match schema.
	m := map[string]interface{}{}

	if obj != nil && obj.PropertiesMap != nil {
		for k, v := range obj.PropertiesMap {
			m[k] = v
		}
	}

	result["map"] = m
	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToPatchInstruction(fieldKeyFormat string) (oci_distributed_database.PatchInstruction, error) {
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

// NOTE: Duplicate generated helper function.
// This function is identical to PatchInstructionToMap defined in another
// generated file within the same package and causes a Go compile-time
// redeclaration error.
//
// The duplication is introduced by the Terraform code generator.
// See JIRA: TOP-9388
//
// DO NOT re-enable locally — generator must be fixed.
/*
/*
func PatchInstructionToMap(obj oci_distributed_database.PatchInstruction) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_distributed_database.PatchInsertInstruction:
		result["operation"] = "INSERT"

		if v.Value != nil {
			result["value"] = []interface{}{objectToMap(v.Value)}
		}
	case oci_distributed_database.PatchMergeInstruction:
		result["operation"] = "MERGE"

		if v.Value != nil {
			result["value"] = []interface{}{objectToMap(v.Value)}
		}
	case oci_distributed_database.PatchRemoveInstruction:
		result["operation"] = "REMOVE"
	default:
		log.Printf("[WARN] Received 'operation' of unknown type %v", obj)
		return nil
	}

	return result
}
*/

/*
func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToobject(fieldKeyFormat string) (oci_distributed_database.Object, error) {
	result := oci_distributed_database.Object{}

	return result, nil
}
*/

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_distributed_database.ChangeDistributedDatabaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DistributedDatabaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ChangeDistributedDatabaseCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDistributedDatabaseFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToVmClusterDetails(fieldKeyFormat string) (oci_distributed_database.VmClusterDetails, error) {
	result := oci_distributed_database.VmClusterDetails{}

	if backupNetworkNsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_network_nsg_ids")); ok {
		set := backupNetworkNsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backup_network_nsg_ids")) {
			result.BackupNetworkNsgIds = tmp
		}
	}

	if backupSubnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_subnet_id")); ok {
		tmp := backupSubnetId.(string)
		result.BackupSubnetId = &tmp
	}

	if domain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain")); ok {
		tmp := domain.(string)
		result.Domain = &tmp
	}

	if enabledECpuCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enabled_ecpu_count")); ok {
		tmp := enabledECpuCount.(int)
		result.EnabledECpuCount = &tmp
	}

	if isDiagnosticsEventsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_diagnostics_events_enabled")); ok {
		tmp := isDiagnosticsEventsEnabled.(bool)
		result.IsDiagnosticsEventsEnabled = &tmp
	}

	if isHealthMonitoringEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_health_monitoring_enabled")); ok {
		tmp := isHealthMonitoringEnabled.(bool)
		result.IsHealthMonitoringEnabled = &tmp
	}

	if isIncidentLogsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_incident_logs_enabled")); ok {
		tmp := isIncidentLogsEnabled.(bool)
		result.IsIncidentLogsEnabled = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "license_model")); ok {
		result.LicenseModel = oci_distributed_database.VmClusterDetailsLicenseModelEnum(licenseModel.(string))
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

	if privateZoneId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_zone_id")); ok {
		tmp := privateZoneId.(string)
		result.PrivateZoneId = &tmp
	}

	if sshPublicKeys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssh_public_keys")); ok {
		interfaces := sshPublicKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ssh_public_keys")) {
			result.SshPublicKeys = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	if totalECpuCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "total_ecpu_count")); ok {
		tmp := totalECpuCount.(int)
		result.TotalECpuCount = &tmp
	}

	if vmFileSystemStorageSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_file_system_storage_size")); ok {
		tmp := vmFileSystemStorageSize.(int)
		result.VmFileSystemStorageSize = &tmp
	}

	return result, nil
}

func VmClusterDetailsToMap(obj *oci_distributed_database.VmClusterDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	backupNetworkNsgIds := []interface{}{}
	for _, item := range obj.BackupNetworkNsgIds {
		backupNetworkNsgIds = append(backupNetworkNsgIds, item)
	}
	if datasource {
		result["backup_network_nsg_ids"] = backupNetworkNsgIds
	} else {
		result["backup_network_nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, backupNetworkNsgIds)
	}

	if obj.BackupSubnetId != nil {
		result["backup_subnet_id"] = string(*obj.BackupSubnetId)
	}

	if obj.Domain != nil {
		result["domain"] = string(*obj.Domain)
	}

	if obj.EnabledECpuCount != nil {
		result["enabled_ecpu_count"] = int(*obj.EnabledECpuCount)
	}

	if obj.IsDiagnosticsEventsEnabled != nil {
		result["is_diagnostics_events_enabled"] = bool(*obj.IsDiagnosticsEventsEnabled)
	}

	if obj.IsHealthMonitoringEnabled != nil {
		result["is_health_monitoring_enabled"] = bool(*obj.IsHealthMonitoringEnabled)
	}

	if obj.IsIncidentLogsEnabled != nil {
		result["is_incident_logs_enabled"] = bool(*obj.IsIncidentLogsEnabled)
	}

	result["license_model"] = string(obj.LicenseModel)

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.PrivateZoneId != nil {
		result["private_zone_id"] = string(*obj.PrivateZoneId)
	}

	result["ssh_public_keys"] = obj.SshPublicKeys

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.TotalECpuCount != nil {
		result["total_ecpu_count"] = int(*obj.TotalECpuCount)
	}

	if obj.VmFileSystemStorageSize != nil {
		result["vm_file_system_storage_size"] = int(*obj.VmFileSystemStorageSize)
	}

	return result
}

/*
func (s *DistributedDatabaseDistributedDatabaseResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_distributed_database.ChangeDistributedDatabaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DistributedDatabaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ChangeDistributedDatabaseCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDistributedDatabaseFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
*/

// WORKAROUND FOR GENERATED CODE ISSUE:
//
// The generator emits calls to mapToDistributedDbMetadata but either
// doesn't generate the function and/or assumes the SDK model has a `Map` field.
// In OCI Go SDK, DistributedDbMetadata uses `PropertiesMap` (json:"map").
//
// This mapper converts the Terraform schema field `metadata.0.map` into
// `DistributedDbMetadata.PropertiesMap`.
//
// Remove once generator is fixed.
// See JIRA: TOP-9397
// Updated method receiver from DistributedDatabaseDistributedAutonomousDatabaseResourceCrud
// to DistributedDatabaseDistributedDatabaseResourceCrud
func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToDistributedDbMetadata(
	fieldKeyFormat string,
) (oci_distributed_database.DistributedDbMetadata, error) {
	result := oci_distributed_database.DistributedDbMetadata{}

	if m, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "map")); ok {
		result.PropertiesMap = tfresource.ObjectMapToStringMap(m.(map[string]interface{}))
	}

	return result, nil
}

// WORKAROUND FOR GENERATED CODE ISSUE:
//
// Some tfresource helpers (and legacy wait patterns) expect a non-context Get().
// This resource CRUD implements GetWithContext(ctx) only.
// Add a thin adapter for compatibility.
//
// See JIRA: TOP-9399
func (s *DistributedDatabaseDistributedDatabaseResourceCrud) Get() error {
	return s.GetWithContext(context.Background())
}

// WORKAROUND FOR GENERATED CODE ISSUE:
// Read/ToMap must operate on SDK model types, not Create*Details request types.
// The generator reused CreateCatalogPeerWithExadbXsDetailsToMap in the Read path,
// but PeerDetails is returned as the SDK model type CatalogPeerWithExadbXs.
// See JIRA: TOP-9405
func CatalogPeerWithExadbXsToMap(obj oci_distributed_database.CatalogPeerWithExadbXs) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.VmClusterId != nil {
		result["vm_cluster_id"] = string(*obj.VmClusterId)
	}

	if obj.ShardGroup != nil {
		result["shard_group"] = string(*obj.ShardGroup)
	}

	if obj.Status != "" {
		result["status"] = string(obj.Status)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.SupportingResourceId != nil {
		result["supporting_resource_id"] = string(*obj.SupportingResourceId)
	}

	if obj.ContainerDatabaseId != nil {
		result["container_database_id"] = string(*obj.ContainerDatabaseId)
	}

	if obj.ProtectionMode != "" {
		result["protection_mode"] = string(obj.ProtectionMode)
	}

	if obj.TransportType != "" {
		result["transport_type"] = string(obj.TransportType)
	}

	// NOTE:
	// obj.Metadata is *DistributedDbMetadata. Only map it if the TF schema defines
	// a corresponding block/object. If not in schema, intentionally skip.
	// If schema expects it, implement DistributedDbMetadataToMap(obj.Metadata).

	return result
}

// WORKAROUND FOR GENERATED CODE ISSUE:
// Read/ToMap must operate on SDK model types, not Create*Details request types.
// The generator incorrectly reused CreateShardPeerWithExadbXsDetailsToMap in Read paths,
// but PeerDetails is returned as the SDK model type ShardPeerWithExadbXs.
// See JIRA: TOP-9407
func ShardPeerWithExadbXsToMap(obj oci_distributed_database.ShardPeerWithExadbXs) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.VmClusterId != nil {
		result["vm_cluster_id"] = string(*obj.VmClusterId)
	}

	if obj.ShardGroup != nil {
		result["shard_group"] = string(*obj.ShardGroup)
	}

	if obj.Status != "" {
		result["status"] = string(obj.Status)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.SupportingResourceId != nil {
		result["supporting_resource_id"] = string(*obj.SupportingResourceId)
	}

	if obj.ContainerDatabaseId != nil {
		result["container_database_id"] = string(*obj.ContainerDatabaseId)
	}

	if obj.ProtectionMode != "" {
		result["protection_mode"] = string(obj.ProtectionMode)
	}

	if obj.TransportType != "" {
		result["transport_type"] = string(obj.TransportType)
	}

	// NOTE:
	// Metadata is a response-side SDK model (*DistributedDbMetadata).
	// Only map this if the Terraform schema defines a corresponding block/object.
	// If not present in schema, intentionally skip to avoid state/schema drift.
	// See JIRA: TOP-9407
	/*
		if obj.Metadata != nil {
			result["metadata"] = DistributedDbMetadataToMap(obj.Metadata)
		}
	*/

	return result
}

func rejectDisabledTriggers(d *schema.ResourceData, triggers []string, context string) diag.Diagnostics {
	for _, attr := range triggers {
		v, ok := d.GetOkExists(attr)
		if !ok {
			continue
		}
		// triggers are int in your provider
		if i, ok2 := v.(int); ok2 && i != 0 {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Trigger disabled in provider",
				Detail: fmt.Sprintf(
					"%s: Trigger %q is disabled and cannot be used.\n\n"+
						"This action API is not supported from Terraform for this resource.\n"+
						"Set it to 0/null and run apply again.",
					context, attr,
				),
			}}
		}
	}
	return nil
}
