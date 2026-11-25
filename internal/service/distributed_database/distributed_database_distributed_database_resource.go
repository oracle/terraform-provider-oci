// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/go-cty/cty"
	ctyjson "github.com/hashicorp/go-cty/cty/json"
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

var distributedDatabaseActionTriggers = []string{
	"add_gds_control_node_trigger",
	"change_db_backup_config_trigger",
	"configure_auto_resource_management_trigger",
	"configure_sharding_trigger",
	"generate_wallet_trigger",
	"move_replication_unit_trigger",
	"recreate_failed_resource_trigger",
	"scale_gsms_trigger",
	"start_database_trigger",
	"stop_database_trigger",
	"validate_network_trigger",
}

var distributedAutonomousDatabaseActionTriggers = []string{
	"add_gds_control_node_trigger",
	"change_db_backup_config_trigger",
	"configure_auto_resource_management_trigger",
	"configure_gsm_wallet_trigger",
	"configure_sharding_trigger",
	"generate_wallet_trigger",
	"move_replication_unit_trigger",
	"recreate_failed_resource_trigger",
	"scale_gsms_trigger",
	"start_database_trigger",
	"stop_database_trigger",
	"validate_ca_bundle_trigger",
	"validate_network_trigger",
}

func validateCreateActionTriggers(d *schema.ResourceData, triggers []string) error {
	for _, trigger := range triggers {
		if value, ok := d.GetOkExists(trigger); ok && value.(int) != 0 {
			return fmt.Errorf("%s must be 0 during resource creation; increment it in a subsequent update to run the action", trigger)
		}
	}

	return nil
}

func hasNonEmptyChangedPatchOperations(d *schema.ResourceData) bool {
	if !d.HasChange("patch_operations") {
		return false
	}

	operations, ok := d.Get("patch_operations").([]interface{})
	return ok && len(operations) > 0
}

func validateDistributedDatabaseWriteOnceReplicationUnitCountDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	if d.Id() == "" || !d.HasChange("system_raft_replication_unit_count") {
		return nil
	}

	return fmt.Errorf("`system_raft_replication_unit_count` is write-once after create; use `effective_system_raft_replication_unit_count` to inspect the live value")
}

func setDistributedDatabaseWriteOnceReplicationUnitCount(d *schema.ResourceData, observed *int) error {
	if observed != nil {
		if err := d.Set("effective_system_raft_replication_unit_count", *observed); err != nil {
			return err
		}
	}

	if d.HasChange("system_raft_replication_unit_count") {
		old, _ := d.GetChange("system_raft_replication_unit_count")
		if configured, ok := old.(int); ok && configured != 0 {
			return d.Set("system_raft_replication_unit_count", configured)
		}
	}

	if configured, ok := d.GetOkExists("system_raft_replication_unit_count"); ok {
		if value, ok := configured.(int); ok && value != 0 {
			return d.Set("system_raft_replication_unit_count", value)
		}
	}

	if observed != nil {
		return d.Set("system_raft_replication_unit_count", *observed)
	}

	return nil
}

func distributedDatabaseHasRecoverableInfrastructureFailure(lifecycleState, lifecycleDetails string) bool {
	if strings.EqualFold(lifecycleState, "FAILED") || strings.EqualFold(lifecycleState, "NEEDS_ATTENTION") {
		return true
	}

	normalizedDetails := strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(lifecycleDetails), "_", " "))
	return strings.Contains(normalizedDetails, "MISSING INFRASTRUCTURE")
}

// Secret OCIDs are ordinary references. Plaintext
// admin passwords are optional sensitive inputs retained with creation topology.
func DistributedDatabaseDistributedDatabaseResource() *schema.Resource {
	resource := &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		CustomizeDiff: validateDistributedDatabaseWriteOnceReplicationUnitCountDiff,
		CreateContext: createDistributedDatabaseDistributedDatabaseWithContext,
		ReadContext:   readDistributedDatabaseDistributedDatabaseWithContext,
		UpdateContext: updateDistributedDatabaseDistributedDatabaseWithContext,
		DeleteContext: deleteDistributedDatabaseDistributedDatabaseWithContext,
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
												"db_home_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"EXISTING_DB_HOME",
														"XD_EXISTING_CLUSTER",
														"XS_EXISTING_CLUSTER",
														"XS_NEW_CLUSTER",
														"XS_NEW_VAULT_AND_CLUSTER",
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
															"autoscale_limit_in_gbs": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"high_capacity_database_storage": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_autoscale_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"subscription_id": {
																Type:     schema.TypeString,
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
												"exascale_db_storage_vault_id": {
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
												"kms_key_version_id": {
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
												"tde_wallet_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"tde_wallet_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"tde_wallet_password_secret_version_number": {
													Type:     schema.TypeInt,
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
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
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
															"subscription_id": {
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
												"db_home_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"EXISTING_DB_HOME",
														"XD_EXISTING_CLUSTER",
														"XS_EXISTING_CLUSTER",
														"XS_NEW_CLUSTER",
														"XS_NEW_VAULT_AND_CLUSTER",
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
															"autoscale_limit_in_gbs": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"high_capacity_database_storage": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_autoscale_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"subscription_id": {
																Type:     schema.TypeString,
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
												"exascale_db_storage_vault_id": {
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
												"kms_key_version_id": {
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
												"tde_wallet_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"tde_wallet_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"tde_wallet_password_secret_version_number": {
													Type:     schema.TypeInt,
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
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
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
															"subscription_id": {
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
												"db_home_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"EXISTING_DB_HOME",
														"XD_EXISTING_CLUSTER",
														"XS_EXISTING_CLUSTER",
														"XS_NEW_CLUSTER",
														"XS_NEW_VAULT_AND_CLUSTER",
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
															"autoscale_limit_in_gbs": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"high_capacity_database_storage": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_autoscale_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"subscription_id": {
																Type:     schema.TypeString,
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
												"exascale_db_storage_vault_id": {
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
												"kms_key_version_id": {
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
												"tde_wallet_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"tde_wallet_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"tde_wallet_password_secret_version_number": {
													Type:     schema.TypeInt,
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
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
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
															"subscription_id": {
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
												"db_home_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"EXISTING_DB_HOME",
														"XD_EXISTING_CLUSTER",
														"XS_EXISTING_CLUSTER",
														"XS_NEW_CLUSTER",
														"XS_NEW_VAULT_AND_CLUSTER",
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
															"autoscale_limit_in_gbs": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"high_capacity_database_storage": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_autoscale_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"subscription_id": {
																Type:     schema.TypeString,
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
												"exascale_db_storage_vault_id": {
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
												"kms_key_version_id": {
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
												"tde_wallet_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"tde_wallet_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"tde_wallet_password_secret_version_number": {
													Type:     schema.TypeInt,
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
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
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
															"subscription_id": {
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
												"db_home_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"EXISTING_DB_HOME",
														"XD_EXISTING_CLUSTER",
														"XS_EXISTING_CLUSTER",
														"XS_NEW_CLUSTER",
														"XS_NEW_VAULT_AND_CLUSTER",
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
															"autoscale_limit_in_gbs": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"high_capacity_database_storage": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_autoscale_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"subscription_id": {
																Type:     schema.TypeString,
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
												"exascale_db_storage_vault_id": {
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
												"kms_key_version_id": {
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
												"tde_wallet_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"tde_wallet_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"tde_wallet_password_secret_version_number": {
													Type:     schema.TypeInt,
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
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
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
															"subscription_id": {
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
						"auto_backup_window": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"auto_full_backup_day": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"auto_full_backup_window": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"backup_deletion_policy": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
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
						"can_run_immediate_full_backup": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_auto_backup_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_remote_backup_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"recovery_window_in_days": {
							Type:     schema.TypeInt,
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
			"scan_listener_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
												"db_home_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"EXISTING_DB_HOME",
														"XD_EXISTING_CLUSTER",
														"XS_EXISTING_CLUSTER",
														"XS_NEW_CLUSTER",
														"XS_NEW_VAULT_AND_CLUSTER",
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
															"autoscale_limit_in_gbs": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"high_capacity_database_storage": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_autoscale_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"subscription_id": {
																Type:     schema.TypeString,
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
												"exascale_db_storage_vault_id": {
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
												"kms_key_version_id": {
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
												"tde_wallet_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"tde_wallet_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"tde_wallet_password_secret_version_number": {
													Type:     schema.TypeInt,
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
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
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
															"subscription_id": {
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
												"db_home_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"EXISTING_DB_HOME",
														"XD_EXISTING_CLUSTER",
														"XS_EXISTING_CLUSTER",
														"XS_NEW_CLUSTER",
														"XS_NEW_VAULT_AND_CLUSTER",
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
															"autoscale_limit_in_gbs": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"high_capacity_database_storage": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_autoscale_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"subscription_id": {
																Type:     schema.TypeString,
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
												"exascale_db_storage_vault_id": {
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
												"kms_key_version_id": {
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
												"tde_wallet_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"tde_wallet_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"tde_wallet_password_secret_version_number": {
													Type:     schema.TypeInt,
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
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
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
															"subscription_id": {
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
									"db_home_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"source": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"EXISTING_DB_HOME",
											"XD_EXISTING_CLUSTER",
											"XS_EXISTING_CLUSTER",
											"XS_NEW_CLUSTER",
											"XS_NEW_VAULT_AND_CLUSTER",
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
												"autoscale_limit_in_gbs": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"compartment_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"high_capacity_database_storage": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_autoscale_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"subscription_id": {
													Type:     schema.TypeString,
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
									"exascale_db_storage_vault_id": {
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
									"kms_key_version_id": {
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
									"tde_wallet_password": {
										Type:      schema.TypeString,
										Optional:  true,
										Computed:  true,
										ForceNew:  true,
										Sensitive: true,
									},
									"tde_wallet_password_secret_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"tde_wallet_password_secret_version_number": {
										Type:     schema.TypeInt,
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
												"compartment_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
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
												"subscription_id": {
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
												"db_home_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"EXISTING_DB_HOME",
														"XD_EXISTING_CLUSTER",
														"XS_EXISTING_CLUSTER",
														"XS_NEW_CLUSTER",
														"XS_NEW_VAULT_AND_CLUSTER",
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
															"autoscale_limit_in_gbs": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"high_capacity_database_storage": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_autoscale_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"subscription_id": {
																Type:     schema.TypeString,
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
												"exascale_db_storage_vault_id": {
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
												"kms_key_version_id": {
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
												"tde_wallet_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"tde_wallet_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"tde_wallet_password_secret_version_number": {
													Type:     schema.TypeInt,
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
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
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
															"subscription_id": {
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
												"db_home_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"source": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"EXISTING_DB_HOME",
														"XD_EXISTING_CLUSTER",
														"XS_EXISTING_CLUSTER",
														"XS_NEW_CLUSTER",
														"XS_NEW_VAULT_AND_CLUSTER",
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
															"autoscale_limit_in_gbs": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"high_capacity_database_storage": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_autoscale_enabled": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"subscription_id": {
																Type:     schema.TypeString,
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
												"exascale_db_storage_vault_id": {
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
												"kms_key_version_id": {
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
												"tde_wallet_password": {
													Type:      schema.TypeString,
													Optional:  true,
													Computed:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"tde_wallet_password_secret_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"tde_wallet_password_secret_version_number": {
													Type:     schema.TypeInt,
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
															"compartment_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
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
															"subscription_id": {
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
	configureDatabaseHomeSchema(resource)
	configureInitialDatabaseTopology(resource)
	configureBackupConfigSchema(resource, false)
	configureMainDatabaseWorkRequestRecoverySchema(resource)
	resource.CustomizeDiff = customdiff.Sequence(validateBackupConfigDiff, validateOperationalStateDiff, resource.CustomizeDiff)
	return resource
}

func createDistributedDatabaseDistributedDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if err := validateBackupConfigUpdate(d); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	if _, err := operationalStateTarget(d); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	if err := validateCreateActionTriggers(d, distributedDatabaseActionTriggers); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	sync := &DistributedDatabaseDistributedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	return nil

}

func readDistributedDatabaseDistributedDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	if err := sync.resumePendingWorkRequest(ctx); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDistributedDatabaseDistributedDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if err := validateBackupConfigUpdate(d); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	target, err := operationalStateTarget(d)
	if err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	ctx, cancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	sync := &DistributedDatabaseDistributedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	if target == "AVAILABLE" {
		if err := sync.StartDistributedDatabase(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("add_gds_control_node_trigger"); ok && sync.D.HasChange("add_gds_control_node_trigger") {
		oldRaw, newRaw := sync.D.GetChange("add_gds_control_node_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.AddDistributedDatabaseGdsControlNode(ctx)

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
		if err := sync.ChangeDistributedDbBackupConfig(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("configure_auto_resource_management_trigger"); ok && sync.D.HasChange("configure_auto_resource_management_trigger") {
		oldRaw, newRaw := sync.D.GetChange("configure_auto_resource_management_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ConfigureDistributedDatabaseAutoResourceManagement(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("configure_auto_resource_management_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("configure_sharding_trigger"); ok && sync.D.HasChange("configure_sharding_trigger") {
		oldRaw, newRaw := sync.D.GetChange("configure_sharding_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ConfigureDistributedDatabaseSharding(ctx)

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
			err := sync.GenerateDistributedDatabaseWallet(ctx)

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

	if _, ok := sync.D.GetOkExists("scale_gsms_trigger"); ok && sync.D.HasChange("scale_gsms_trigger") {
		oldRaw, newRaw := sync.D.GetChange("scale_gsms_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ScaleDistributedDatabaseGsms(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("scale_gsms_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("validate_network_trigger"); ok && sync.D.HasChange("validate_network_trigger") {
		oldRaw, newRaw := sync.D.GetChange("validate_network_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ValidateDistributedDatabaseNetwork(ctx)

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
		if err := sync.StopDistributedDatabase(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}
	if target != "" {
		return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
	}
	return nil
}

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
		string(oci_distributed_database.DistributedDatabaseLifecycleStateInactive),
		string(oci_distributed_database.DistributedDatabaseLifecycleStateActive),
		string(oci_distributed_database.DistributedDatabaseLifecycleStateNeedsAttention),
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

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compositeDataGuardShardSpaces, ok := s.D.GetOkExists("composite_data_guard_shard_spaces"); ok {
		interfaces := compositeDataGuardShardSpaces.([]interface{})
		tmp := make([]oci_distributed_database.CreateCompositeDataGuardShardSpaceDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "composite_data_guard_shard_spaces", stateDataIndex)
			converted, err := s.mapToCreateCompositeDataGuardShardSpaceDetails(fieldKeyFormat)
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
		tmp := make([]oci_distributed_database.CreateCompositeRaftShardSpaceDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "composite_raft_shard_spaces", stateDataIndex)
			converted, err := s.mapToCreateCompositeRaftShardSpaceDetails(fieldKeyFormat)
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
			tmp, err := s.mapToDistributedDbBackupConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DbBackupConfig = &tmp
		}
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

	if scanListenerPort, ok := s.D.GetOkExists("scan_listener_port"); ok {
		tmp := scanListenerPort.(int)
		request.ScanListenerPort = &tmp
	}

	if systemChunkCount, ok := s.D.GetOkExists("system_chunk_count"); ok {
		tmp := systemChunkCount.(int)
		request.SystemChunkCount = &tmp
	}

	if systemDataGuardDatabases, ok := s.D.GetOkExists("system_data_guard_databases"); ok {
		if tmpList := systemDataGuardDatabases.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "system_data_guard_databases", 0)
			tmp, err := s.mapToCreateSystemDataGuardDatabaseDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SystemDataGuardDatabases = &tmp
		}
	}

	if systemRaftClusters, ok := s.D.GetOkExists("system_raft_clusters"); ok {
		interfaces := systemRaftClusters.([]interface{})
		tmp := make([]oci_distributed_database.CreateSystemRaftClusterDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "system_raft_clusters", stateDataIndex)
			converted, err := s.mapToCreateSystemRaftClusterDetails(fieldKeyFormat)
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
		tmp := make([]oci_distributed_database.CreateUserShardSpaceDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user_shard_spaces", stateDataIndex)
			converted, err := s.mapToCreateUserShardSpaceDetails(fieldKeyFormat)
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

	if workId == nil {
		return fmt.Errorf("CreateDistributedDatabase returned no work request ID")
	}
	if err := recordPendingMainDatabaseWorkRequest(s.D, *workId, "create", string(oci_distributed_database.ActionTypeCreated), ""); err != nil {
		return err
	}
	return s.resumePendingWorkRequest(ctx)
}
func (s *DistributedDatabaseDistributedDatabaseResourceCrud) Patch(ctx context.Context) error {
	request := oci_distributed_database.PatchDistributedDatabaseRequest{}

	tmp := s.D.Id()
	request.DistributedDatabaseId = &tmp

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
	response, err := s.Client.PatchDistributedDatabase(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDistributedDatabaseFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) getDistributedDatabaseFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_distributed_database.ActionTypeEnum, timeout time.Duration) error {

	if workId == nil {
		return fmt.Errorf("work request ID is required")
	}
	if err := recordPendingMainDatabaseWorkRequest(s.D, *workId, "update", string(actionTypeEnum), ""); err != nil {
		return err
	}
	// Wait until it finishes
	distributedDatabaseId, err := distributedDatabaseWaitForWorkRequest(ctx, workId, "distributeddatabase",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*distributedDatabaseId)

	if err := s.GetWithContext(ctx); err != nil {
		return err
	}
	return clearPendingMainDatabaseWorkRequest(s.D)
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) resumePendingWorkRequest(ctx context.Context) error {
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
	id, err := distributedDatabaseWaitForWorkRequest(ctx, &workID, "distributeddatabase", action, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)
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

func distributedDatabaseWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_distributed_database.ActionTypeEnum,
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
		return nil, getErrorFromDistributedDatabaseDistributedDatabaseWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.GetDistributedDatabaseRequest{Metadata: oci_common.String(operationalStateMetadataQuery)}

	tmp := s.D.Id()
	request.DistributedDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.GetDistributedDatabase(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DistributedDatabase
	return nil
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) UpdateWithContext(ctx context.Context) error {
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

	response, err := s.Client.UpdateDistributedDatabase(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DistributedDatabase
	return nil
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_distributed_database.DeleteDistributedDatabaseRequest{}

	tmp := s.D.Id()
	request.DistributedDatabaseId = &tmp

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
		log.Printf("[DEBUG] Preserving configured catalog_details for distributed database %q during recoverable infrastructure failure", s.D.Id())
	} else {
		catalogDetails := []interface{}{}
		for _, item := range s.Res.CatalogDetails {
			catalogDetails = append(catalogDetails, DistributedDatabaseCatalogToMap(item))
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
		compositeDataGuardShardSpaces = append(compositeDataGuardShardSpaces, CompositeDataGuardShardSpaceToMap(item))
	}
	s.D.Set("composite_data_guard_shard_spaces", compositeDataGuardShardSpaces)

	compositeRaftShardSpaces := []interface{}{}
	for _, item := range s.Res.CompositeRaftShardSpaces {
		compositeRaftShardSpaces = append(compositeRaftShardSpaces, CompositeRaftShardSpaceToMap(item))
	}
	s.D.Set("composite_raft_shard_spaces", compositeRaftShardSpaces)

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

	gdsControlNodeDetails := []interface{}{}
	for _, item := range s.Res.GdsControlNodeDetails {
		gdsControlNodeDetails = append(gdsControlNodeDetails, DistributedDatabaseGdsControlNodeToMap(item))
	}
	s.D.Set("gds_control_node_details", gdsControlNodeDetails)

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

	if s.Res.ScanListenerPort != nil {
		s.D.Set("scan_listener_port", *s.Res.ScanListenerPort)
	}

	s.D.Set("sharding_method", s.Res.ShardingMethod)

	s.D.Set("state", s.Res.LifecycleState)
	if err := preserveOperationalState(s.D); err != nil {
		return err
	}

	if s.Res.SystemChunkCount != nil {
		s.D.Set("system_chunk_count", *s.Res.SystemChunkCount)
	}

	if s.Res.SystemDataGuardDatabases != nil {
		s.D.Set("system_data_guard_databases", []interface{}{SystemDataGuardDatabaseToMap(s.Res.SystemDataGuardDatabases)})
	} else {
		s.D.Set("system_data_guard_databases", nil)
	}

	systemRaftClusters := []interface{}{}
	for _, item := range s.Res.SystemRaftClusters {
		systemRaftClusters = append(systemRaftClusters, SystemRaftClusterToMap(item))
	}
	s.D.Set("system_raft_clusters", systemRaftClusters)

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
		userShardSpaces = append(userShardSpaces, UserShardSpaceToMap(item))
	}
	s.D.Set("user_shard_spaces", userShardSpaces)

	vcnNsgIds := []interface{}{}
	for _, item := range s.Res.VcnNsgIds {
		vcnNsgIds = append(vcnNsgIds, VcnNsgIdsDetailsToMap(item, false))
	}
	s.D.Set("vcn_nsg_ids", vcnNsgIds)

	return preserveInitialDatabaseTopology(s.D, initialTopology, false)
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) AddDistributedDatabaseGdsControlNode(ctx context.Context) error {
	request := oci_distributed_database.AddDistributedDatabaseGdsControlNodeRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	if publicSshKey, ok := s.D.GetOkExists("public_ssh_key"); ok {
		tmp := publicSshKey.(string)
		request.PublicSshKey = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.AddDistributedDatabaseGdsControlNode(ctx, request)
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
		_, err = distributedDatabaseWaitForWorkRequest(ctx, workRequestId, "distributeddatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) ChangeDistributedDbBackupConfig(ctx context.Context) error {
	request := oci_distributed_database.ChangeDistributedDbBackupConfigRequest{}

	if err := decodeBackupConfig(s.D.GetRawConfig(), &request.DbBackupConfig); err != nil {
		return fmt.Errorf("ChangeDistributedDbBackupConfig: %w", err)
	}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ChangeDistributedDbBackupConfig(ctx, request)
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) ConfigureDistributedDatabaseAutoResourceManagement(ctx context.Context) error {
	request := oci_distributed_database.ConfigureDistributedDatabaseAutoResourceManagementRequest{}

	if actionType, ok := s.D.GetOkExists("action_type"); ok {
		request.ActionType = oci_distributed_database.ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum(actionType.(string))
	}

	if coolOffPeriodInMinutes, ok := s.D.GetOkExists("cool_off_period_in_minutes"); ok {
		tmp := coolOffPeriodInMinutes.(int)
		request.CoolOffPeriodInMinutes = &tmp
	}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if maxMoveRuAttempts, ok := s.D.GetOkExists("max_move_ru_attempts"); ok {
		tmp := maxMoveRuAttempts.(int)
		request.MaxMoveRuAttempts = &tmp
	}

	if maxMoveRuUnit, ok := s.D.GetOkExists("max_move_ru_unit"); ok {
		request.MaxMoveRuUnit = oci_distributed_database.ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum(maxMoveRuUnit.(string))
	}

	if mode, ok := s.D.GetOkExists("mode"); ok {
		request.Mode = oci_distributed_database.ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum(mode.(string))
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

	response, err := s.Client.ConfigureDistributedDatabaseAutoResourceManagement(ctx, request)
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
		_, err = distributedDatabaseWaitForWorkRequest(ctx, workRequestId, "distributeddatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) ConfigureDistributedDatabaseSharding(ctx context.Context) error {
	request := oci_distributed_database.ConfigureDistributedDatabaseShardingRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	if isRebalanceRequired, ok := s.D.GetOkExists("is_rebalance_required"); ok {
		tmp := isRebalanceRequired.(bool)
		request.IsRebalanceRequired = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ConfigureDistributedDatabaseSharding(ctx, request)
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
		_, err = distributedDatabaseWaitForWorkRequest(ctx, workRequestId, "distributeddatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) GenerateDistributedDatabaseWallet(ctx context.Context) error {
	request := oci_distributed_database.GenerateDistributedDatabaseWalletRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.GenerateDistributedDatabaseWallet(ctx, request)
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
		_, err = distributedDatabaseWaitForWorkRequest(ctx, workRequestId, "distributeddatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
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

	response, err := s.Client.MoveDistributedDatabaseReplicationUnit(ctx, request)
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
		_, err = distributedDatabaseWaitForWorkRequest(ctx, workRequestId, "distributeddatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) RecreateFailedDistributedDatabaseResource(ctx context.Context) error {
	request := oci_distributed_database.RecreateFailedDistributedDatabaseResourceRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.RecreateFailedDistributedDatabaseResource(ctx, request)
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
		_, err = distributedDatabaseWaitForWorkRequest(ctx, workRequestId, "distributeddatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) ScaleDistributedDatabaseGsms(ctx context.Context) error {
	request := oci_distributed_database.ScaleDistributedDatabaseGsmsRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ScaleDistributedDatabaseGsms(ctx, request)
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
		_, err = distributedDatabaseWaitForWorkRequest(ctx, workRequestId, "distributeddatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) StartDistributedDatabase(ctx context.Context) error {
	request := oci_distributed_database.StartDistributedDatabaseRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.StartDistributedDatabase(ctx, request)
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) StopDistributedDatabase(ctx context.Context) error {
	request := oci_distributed_database.StopDistributedDatabaseRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.StopDistributedDatabase(ctx, request)
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) ValidateDistributedDatabaseNetwork(ctx context.Context) error {
	request := oci_distributed_database.ValidateDistributedDatabaseNetworkRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabaseId = &idTmp

	if isSurrogate, ok := s.D.GetOkExists("is_surrogate"); ok {
		tmp := isSurrogate.(bool)
		request.IsSurrogate = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ValidateDistributedDatabaseNetwork(ctx, request)
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
		_, err = distributedDatabaseWaitForWorkRequest(ctx, workRequestId, "distributeddatabase", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateCompositeDataGuardShardSpaceDetails(fieldKeyFormat string) (oci_distributed_database.CreateCompositeDataGuardShardSpaceDetails, error) {
	result := oci_distributed_database.CreateCompositeDataGuardShardSpaceDetails{}

	if chunkCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "chunk_count")); ok {
		tmp := chunkCount.(int)
		result.ChunkCount = &tmp
	}

	if dataGuardReplicas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")); ok {
		interfaces := dataGuardReplicas.([]interface{})
		tmp := make([]oci_distributed_database.CreateCompositeDataGuardShardSpaceReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data_guard_replicas"), stateDataIndex)
			converted, err := s.mapToCreateCompositeDataGuardShardSpaceReplicaDetails(fieldKeyFormatNextLevel)
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
			tmp, err := s.mapToCreateCompositeDataGuardShardSpaceReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert original_replica, encountered error: %v", err)
			}
			result.OriginalReplica = &tmp
		}
	}

	return result, nil
}

func CompositeDataGuardShardSpaceToMap(obj oci_distributed_database.CompositeDataGuardShardSpace) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ChunkCount != nil {
		result["chunk_count"] = int(*obj.ChunkCount)
	}

	dataGuardReplicas := []interface{}{}
	for _, item := range obj.DataGuardReplicas {
		dataGuardReplicas = append(dataGuardReplicas, CompositeDataGuardShardSpaceReplicaToMap(item))
	}
	result["data_guard_replicas"] = dataGuardReplicas

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.OriginalReplica != nil {
		result["original_replica"] = []interface{}{CompositeDataGuardShardSpaceReplicaToMap(*obj.OriginalReplica)}
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateCompositeDataGuardShardSpaceReplicaDetails(fieldKeyFormat string) (oci_distributed_database.CreateCompositeDataGuardShardSpaceReplicaDetails, error) {
	result := oci_distributed_database.CreateCompositeDataGuardShardSpaceReplicaDetails{}

	if databases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "databases")); ok {
		interfaces := databases.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedDatabaseShardDatabaseDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "databases"), stateDataIndex)
			converted, err := s.mapToCreateDistributedDatabaseShardDatabaseDetails(fieldKeyFormatNextLevel)
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

func CompositeDataGuardShardSpaceReplicaToMap(obj oci_distributed_database.CompositeDataGuardShardSpaceReplica) map[string]interface{} {
	result := map[string]interface{}{}

	databases := []interface{}{}
	for _, item := range obj.Databases {
		databases = append(databases, DistributedDatabaseShardDatabaseToMap(item))
	}
	result["databases"] = databases

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateCompositeRaftShardSpaceDetails(fieldKeyFormat string) (oci_distributed_database.CreateCompositeRaftShardSpaceDetails, error) {
	result := oci_distributed_database.CreateCompositeRaftShardSpaceDetails{}

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
		tmp := make([]oci_distributed_database.CreateCompositeShardSpaceRaftClusterDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "raft_clusters"), stateDataIndex)
			converted, err := s.mapToCreateCompositeShardSpaceRaftClusterDetails(fieldKeyFormatNextLevel)
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

func CompositeRaftShardSpaceToMap(obj oci_distributed_database.CompositeRaftShardSpace) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ChunkCount != nil {
		result["chunk_count"] = int(*obj.ChunkCount)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	raftClusters := []interface{}{}
	for _, item := range obj.RaftClusters {
		raftClusters = append(raftClusters, CompositeShardSpaceRaftClusterToMap(item))
	}
	result["raft_clusters"] = raftClusters

	if obj.ReplicationUnitCount != nil {
		result["replication_unit_count"] = int(*obj.ReplicationUnitCount)
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateCompositeShardSpaceRaftClusterDetails(fieldKeyFormat string) (oci_distributed_database.CreateCompositeShardSpaceRaftClusterDetails, error) {
	result := oci_distributed_database.CreateCompositeShardSpaceRaftClusterDetails{}

	if databases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "databases")); ok {
		interfaces := databases.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedDatabaseShardDatabaseDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "databases"), stateDataIndex)
			converted, err := s.mapToCreateDistributedDatabaseShardDatabaseDetails(fieldKeyFormatNextLevel)
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
		result.RuMode = oci_distributed_database.CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum(ruMode.(string))
	}

	if witnessCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "witness_count")); ok {
		tmp := witnessCount.(int)
		result.WitnessCount = &tmp
	}

	return result, nil
}

func CompositeShardSpaceRaftClusterToMap(obj oci_distributed_database.CompositeShardSpaceRaftCluster) map[string]interface{} {
	result := map[string]interface{}{}

	databases := []interface{}{}
	for _, item := range obj.Databases {
		databases = append(databases, DistributedDatabaseShardDatabaseToMap(item))
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateDistributedDatabaseCatalogDatabaseDetails(fieldKeyFormat string) (oci_distributed_database.CreateDistributedDatabaseCatalogDatabaseDetails, error) {
	var baseObject oci_distributed_database.CreateDistributedDatabaseCatalogDatabaseDetails
	//discriminator
	sourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source"))
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("EXISTING_DB_HOME"):
		details := oci_distributed_database.CreateDistributedDatabaseCatalogDatabaseWithDbHomeDetails{}
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
		if dbHomeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_home_id")); ok {
			tmp := dbHomeId.(string)
			details.DbHomeId = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
			details.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
			tmp := tdeWalletPassword.(string)
			details.TdeWalletPassword = &tmp
		}
		if tdeWalletPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_id")); ok {
			tmp := tdeWalletPasswordSecretId.(string)
			details.TdeWalletPasswordSecretId = &tmp
		}
		if tdeWalletPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_version_number")); ok {
			tmp := tdeWalletPasswordSecretVersionNumber.(int)
			details.TdeWalletPasswordSecretVersionNumber = &tmp
		}
		baseObject = details
	case strings.ToLower("XD_EXISTING_CLUSTER"):
		details := oci_distributed_database.CreateDistributedDatabaseCatalogDatabaseWithExadbXdDetails{}
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
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
			details.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if vmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_id")); ok {
			tmp := vmClusterId.(string)
			details.VmClusterId = &tmp
		}
		if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
			tmp := tdeWalletPassword.(string)
			details.TdeWalletPassword = &tmp
		}
		if tdeWalletPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_id")); ok {
			tmp := tdeWalletPasswordSecretId.(string)
			details.TdeWalletPasswordSecretId = &tmp
		}
		if tdeWalletPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_version_number")); ok {
			tmp := tdeWalletPasswordSecretVersionNumber.(int)
			details.TdeWalletPasswordSecretVersionNumber = &tmp
		}
		baseObject = details
	case strings.ToLower("XS_EXISTING_CLUSTER"):
		details := oci_distributed_database.CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails{}
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
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
			details.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if vmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_id")); ok {
			tmp := vmClusterId.(string)
			details.VmClusterId = &tmp
		}
		if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
			tmp := tdeWalletPassword.(string)
			details.TdeWalletPassword = &tmp
		}
		if tdeWalletPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_id")); ok {
			tmp := tdeWalletPasswordSecretId.(string)
			details.TdeWalletPasswordSecretId = &tmp
		}
		if tdeWalletPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_version_number")); ok {
			tmp := tdeWalletPasswordSecretVersionNumber.(int)
			details.TdeWalletPasswordSecretVersionNumber = &tmp
		}
		baseObject = details
	case strings.ToLower("XS_NEW_CLUSTER"):
		details := oci_distributed_database.CreateDistributedDatabaseCatalogWithExadbXsNewClusterDetails{}
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
		if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		if exascaleDbStorageVaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exascale_db_storage_vault_id")); ok {
			tmp := exascaleDbStorageVaultId.(string)
			details.ExascaleDbStorageVaultId = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
			details.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
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
		if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
			tmp := tdeWalletPassword.(string)
			details.TdeWalletPassword = &tmp
		}
		if tdeWalletPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_id")); ok {
			tmp := tdeWalletPasswordSecretId.(string)
			details.TdeWalletPasswordSecretId = &tmp
		}
		if tdeWalletPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_version_number")); ok {
			tmp := tdeWalletPasswordSecretVersionNumber.(int)
			details.TdeWalletPasswordSecretVersionNumber = &tmp
		}
		baseObject = details
	case strings.ToLower("XS_NEW_VAULT_AND_CLUSTER"):
		details := oci_distributed_database.CreateDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterDetails{}
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
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
			details.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
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
		if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
			tmp := tdeWalletPassword.(string)
			details.TdeWalletPassword = &tmp
		}
		if tdeWalletPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_id")); ok {
			tmp := tdeWalletPasswordSecretId.(string)
			details.TdeWalletPasswordSecretId = &tmp
		}
		if tdeWalletPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_version_number")); ok {
			tmp := tdeWalletPasswordSecretVersionNumber.(int)
			details.TdeWalletPasswordSecretVersionNumber = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source '%v' was specified", source)
	}
	return baseObject, nil
}

func DistributedDatabaseCatalogDatabaseToMap(obj oci_distributed_database.DistributedDatabaseCatalogDatabase) map[string]interface{} {
	if obj == nil {
		return nil
	}

	result := map[string]interface{}{}
	// preserve response status and database IDs for every SDK variant.
	switch v := obj.(type) {
	case oci_distributed_database.DistributedDatabaseCatalogDatabaseWithDbHome:
		result["source"] = "EXISTING_DB_HOME"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

	case oci_distributed_database.DistributedDatabaseCatalogDatabaseWithExadbXd:
		result["source"] = "XD_EXISTING_CLUSTER"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

	case oci_distributed_database.DistributedDatabaseCatalogDatabaseWithExadbXs:
		result["source"] = "XS_EXISTING_CLUSTER"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

	case oci_distributed_database.DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster:
		result["source"] = "XS_NEW_CLUSTER"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

	case oci_distributed_database.DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster:
		result["source"] = "XS_NEW_VAULT_AND_CLUSTER"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

	}
	if metadata := obj.GetMetadata(); metadata != nil {
		result["metadata"] = []interface{}{DistributedDbMetadataToMap(metadata)}
	}
	if name := obj.GetName(); name != nil {
		result["name"] = *name
	}
	if timeCreated := obj.GetTimeCreated(); timeCreated != nil {
		result["time_created"] = timeCreated.String()
	}
	if timeUpdated := obj.GetTimeUpdated(); timeUpdated != nil {
		result["time_updated"] = timeUpdated.String()
	}
	return result
}

/*
func generatedDistributedDatabaseCatalogDatabaseToMap(obj *oci_distributed_database.DistributedDatabaseCatalogDatabase) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_distributed_database.CreateDistributedDatabaseCatalogDatabaseWithDbHomeDetails:
		result["source"] = "EXISTING_DB_HOME"

		if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}

		if v.AdminPasswordSecretId != nil {
			result["admin_password_secret_id"] = string(*v.AdminPasswordSecretId)
		}

		if v.AdminPasswordSecretVersionNumber != nil {
			result["admin_password_secret_version_number"] = int(*v.AdminPasswordSecretVersionNumber)
		}

		if v.DbHomeId != nil {
			result["db_home_id"] = string(*v.DbHomeId)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		result["transport_type"] = string(v.TransportType)

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedDbMetadataToMap(v.Metadata)}
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
	case oci_distributed_database.CreateDistributedDatabaseCatalogDatabaseWithExadbXdDetails:
		result["source"] = "XD_EXISTING_CLUSTER"

		if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}

		if v.AdminPasswordSecretId != nil {
			result["admin_password_secret_id"] = string(*v.AdminPasswordSecretId)
		}

		if v.AdminPasswordSecretVersionNumber != nil {
			result["admin_password_secret_version_number"] = int(*v.AdminPasswordSecretVersionNumber)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		result["transport_type"] = string(v.TransportType)

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.VmClusterId != nil {
			result["vm_cluster_id"] = string(*v.VmClusterId)
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedDbMetadataToMap(v.Metadata)}
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
	case oci_distributed_database.CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails:
		result["source"] = "XS_EXISTING_CLUSTER"

		if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}

		if v.AdminPasswordSecretId != nil {
			result["admin_password_secret_id"] = string(*v.AdminPasswordSecretId)
		}

		if v.AdminPasswordSecretVersionNumber != nil {
			result["admin_password_secret_version_number"] = int(*v.AdminPasswordSecretVersionNumber)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		result["transport_type"] = string(v.TransportType)

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.VmClusterId != nil {
			result["vm_cluster_id"] = string(*v.VmClusterId)
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedDbMetadataToMap(v.Metadata)}
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
	case oci_distributed_database.CreateDistributedDatabaseCatalogWithExadbXsNewClusterDetails:
		result["source"] = "XS_NEW_CLUSTER"

		if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}

		if v.AdminPasswordSecretId != nil {
			result["admin_password_secret_id"] = string(*v.AdminPasswordSecretId)
		}

		if v.AdminPasswordSecretVersionNumber != nil {
			result["admin_password_secret_version_number"] = int(*v.AdminPasswordSecretVersionNumber)
		}

		if v.AvailabilityDomain != nil {
			result["availability_domain"] = string(*v.AvailabilityDomain)
		}

		if v.ExascaleDbStorageVaultId != nil {
			result["exascale_db_storage_vault_id"] = string(*v.ExascaleDbStorageVaultId)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		result["transport_type"] = string(v.TransportType)

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.VmClusterDetails != nil {
			result["vm_cluster_details"] = []interface{}{VmClusterDetailsToMap(v.VmClusterDetails, false)}
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedDbMetadataToMap(v.Metadata)}
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
	case oci_distributed_database.CreateDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterDetails:
		result["source"] = "XS_NEW_VAULT_AND_CLUSTER"

		if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}

		if v.AdminPasswordSecretId != nil {
			result["admin_password_secret_id"] = string(*v.AdminPasswordSecretId)
		}

		if v.AdminPasswordSecretVersionNumber != nil {
			result["admin_password_secret_version_number"] = int(*v.AdminPasswordSecretVersionNumber)
		}

		if v.AvailabilityDomain != nil {
			result["availability_domain"] = string(*v.AvailabilityDomain)
		}

		if v.DbStorageVaultDetails != nil {
			result["db_storage_vault_details"] = []interface{}{DbStorageVaultDetailsToMap(v.DbStorageVaultDetails)}
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		result["transport_type"] = string(v.TransportType)

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.VmClusterDetails != nil {
			result["vm_cluster_details"] = []interface{}{VmClusterDetailsToMap(v.VmClusterDetails, false)}
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedDbMetadataToMap(v.Metadata)}
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
*/

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateDistributedDatabaseCatalogDetails(fieldKeyFormat string) (oci_distributed_database.CreateDistributedDatabaseCatalogDetails, error) {
	result := oci_distributed_database.CreateDistributedDatabaseCatalogDetails{}

	if dataGuardReplicas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")); ok {
		interfaces := dataGuardReplicas.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedDatabaseCatalogReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data_guard_replicas"), stateDataIndex)
			converted, err := s.mapToCreateDistributedDatabaseCatalogReplicaDetails(fieldKeyFormatNextLevel)
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
			tmp, err := s.mapToCreateDistributedDatabaseCatalogReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert original_replica, encountered error: %v", err)
			}
			result.OriginalReplica = &tmp
		}
	}

	return result, nil
}

func DistributedDatabaseCatalogToMap(obj oci_distributed_database.DistributedDatabaseCatalog) map[string]interface{} {
	result := map[string]interface{}{}

	dataGuardReplicas := []interface{}{}
	for _, item := range obj.DataGuardReplicas {
		dataGuardReplicas = append(dataGuardReplicas, DistributedDatabaseCatalogReplicaDetailsToMap(item))
	}
	result["data_guard_replicas"] = dataGuardReplicas

	if obj.OriginalReplica != nil {
		result["original_replica"] = []interface{}{DistributedDatabaseCatalogReplicaDetailsToMap(*obj.OriginalReplica)}
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateDistributedDatabaseCatalogReplicaDetails(fieldKeyFormat string) (oci_distributed_database.CreateDistributedDatabaseCatalogReplicaDetails, error) {
	result := oci_distributed_database.CreateDistributedDatabaseCatalogReplicaDetails{}

	if database, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database")); ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database"), 0)
			tmp, err := s.mapToCreateDistributedDatabaseCatalogDatabaseDetails(fieldKeyFormatNextLevel)
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

func DistributedDatabaseCatalogReplicaDetailsToMap(obj oci_distributed_database.DistributedDatabaseCatalogReplicaDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Database != nil {
		databaseArray := []interface{}{}
		if databaseMap := DistributedDatabaseCatalogDatabaseToMap(obj.Database); databaseMap != nil {
			databaseArray = append(databaseArray, databaseMap)
		}
		result["database"] = databaseArray
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateDistributedDatabaseShardDatabaseDetails(fieldKeyFormat string) (oci_distributed_database.CreateDistributedDatabaseShardDatabaseDetails, error) {
	var baseObject oci_distributed_database.CreateDistributedDatabaseShardDatabaseDetails
	//discriminator
	sourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source"))
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("EXISTING_DB_HOME"):
		details := oci_distributed_database.CreateDistributedDatabaseShardDatabaseWithDbHomeDetails{}
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
		if dbHomeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_home_id")); ok {
			tmp := dbHomeId.(string)
			details.DbHomeId = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
			details.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
			tmp := tdeWalletPassword.(string)
			details.TdeWalletPassword = &tmp
		}
		if tdeWalletPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_id")); ok {
			tmp := tdeWalletPasswordSecretId.(string)
			details.TdeWalletPasswordSecretId = &tmp
		}
		if tdeWalletPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_version_number")); ok {
			tmp := tdeWalletPasswordSecretVersionNumber.(int)
			details.TdeWalletPasswordSecretVersionNumber = &tmp
		}
		baseObject = details
	case strings.ToLower("XD_EXISTING_CLUSTER"):
		details := oci_distributed_database.CreateDistributedDatabaseShardDatabaseWithExadbXdDetails{}
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
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
			details.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if vmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_id")); ok {
			tmp := vmClusterId.(string)
			details.VmClusterId = &tmp
		}
		if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
			tmp := tdeWalletPassword.(string)
			details.TdeWalletPassword = &tmp
		}
		if tdeWalletPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_id")); ok {
			tmp := tdeWalletPasswordSecretId.(string)
			details.TdeWalletPasswordSecretId = &tmp
		}
		if tdeWalletPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_version_number")); ok {
			tmp := tdeWalletPasswordSecretVersionNumber.(int)
			details.TdeWalletPasswordSecretVersionNumber = &tmp
		}
		baseObject = details
	case strings.ToLower("XS_EXISTING_CLUSTER"):
		details := oci_distributed_database.CreateDistributedDatabaseShardDatabaseWithExadbXsDetails{}
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
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
			details.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		if vmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_id")); ok {
			tmp := vmClusterId.(string)
			details.VmClusterId = &tmp
		}
		if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
			tmp := tdeWalletPassword.(string)
			details.TdeWalletPassword = &tmp
		}
		if tdeWalletPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_id")); ok {
			tmp := tdeWalletPasswordSecretId.(string)
			details.TdeWalletPasswordSecretId = &tmp
		}
		if tdeWalletPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_version_number")); ok {
			tmp := tdeWalletPasswordSecretVersionNumber.(int)
			details.TdeWalletPasswordSecretVersionNumber = &tmp
		}
		baseObject = details
	case strings.ToLower("XS_NEW_CLUSTER"):
		details := oci_distributed_database.CreateDistributedDatabaseShardWithExadbXsNewClusterDetails{}
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
		if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		if exascaleDbStorageVaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exascale_db_storage_vault_id")); ok {
			tmp := exascaleDbStorageVaultId.(string)
			details.ExascaleDbStorageVaultId = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
			details.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
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
		if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
			tmp := tdeWalletPassword.(string)
			details.TdeWalletPassword = &tmp
		}
		if tdeWalletPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_id")); ok {
			tmp := tdeWalletPasswordSecretId.(string)
			details.TdeWalletPasswordSecretId = &tmp
		}
		if tdeWalletPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_version_number")); ok {
			tmp := tdeWalletPasswordSecretVersionNumber.(int)
			details.TdeWalletPasswordSecretVersionNumber = &tmp
		}
		baseObject = details
	case strings.ToLower("XS_NEW_VAULT_AND_CLUSTER"):
		details := oci_distributed_database.CreateDistributedDatabaseShardWithExadbXsNewVaultAndClusterDetails{}
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
		if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
			details.ProtectionMode = oci_distributed_database.DistributedDbProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_type")); ok {
			details.TransportType = oci_distributed_database.DistributedDbTransportTypeEnum(transportType.(string))
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
		if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
			tmp := tdeWalletPassword.(string)
			details.TdeWalletPassword = &tmp
		}
		if tdeWalletPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_id")); ok {
			tmp := tdeWalletPasswordSecretId.(string)
			details.TdeWalletPasswordSecretId = &tmp
		}
		if tdeWalletPasswordSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password_secret_version_number")); ok {
			tmp := tdeWalletPasswordSecretVersionNumber.(int)
			details.TdeWalletPasswordSecretVersionNumber = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source '%v' was specified", source)
	}
	return baseObject, nil
}

func DistributedDatabaseShardDatabaseToMap(obj oci_distributed_database.DistributedDatabaseShardDatabase) map[string]interface{} {
	if obj == nil {
		return nil
	}

	result := map[string]interface{}{}
	// preserve response status and database IDs for every SDK variant.
	switch v := obj.(type) {
	case oci_distributed_database.DistributedDatabaseShardDatabaseWithDbHome:
		result["source"] = "EXISTING_DB_HOME"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

	case oci_distributed_database.DistributedDatabaseShardDatabaseWithExadbXd:
		result["source"] = "XD_EXISTING_CLUSTER"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

	case oci_distributed_database.DistributedDatabaseShardDatabaseWithExadbXs:
		result["source"] = "XS_EXISTING_CLUSTER"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

	case oci_distributed_database.DistributedDatabaseShardDatabaseWithExadbXsNewCluster:
		result["source"] = "XS_NEW_CLUSTER"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

	case oci_distributed_database.DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster:
		result["source"] = "XS_NEW_VAULT_AND_CLUSTER"
		result["status"] = string(v.Status)

		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

	}
	if metadata := obj.GetMetadata(); metadata != nil {
		result["metadata"] = []interface{}{DistributedDbMetadataToMap(metadata)}
	}
	if name := obj.GetName(); name != nil {
		result["name"] = *name
	}
	if timeCreated := obj.GetTimeCreated(); timeCreated != nil {
		result["time_created"] = timeCreated.String()
	}
	if timeUpdated := obj.GetTimeUpdated(); timeUpdated != nil {
		result["time_updated"] = timeUpdated.String()
	}
	return result
}

/*
func generatedDistributedDatabaseShardDatabaseToMap(obj *oci_distributed_database.DistributedDatabaseShardDatabase) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_distributed_database.CreateDistributedDatabaseShardDatabaseWithDbHomeDetails:
		result["source"] = "EXISTING_DB_HOME"

		if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}

		if v.AdminPasswordSecretId != nil {
			result["admin_password_secret_id"] = string(*v.AdminPasswordSecretId)
		}

		if v.AdminPasswordSecretVersionNumber != nil {
			result["admin_password_secret_version_number"] = int(*v.AdminPasswordSecretVersionNumber)
		}

		if v.DbHomeId != nil {
			result["db_home_id"] = string(*v.DbHomeId)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		result["transport_type"] = string(v.TransportType)

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedDbMetadataToMap(v.Metadata)}
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
	case oci_distributed_database.CreateDistributedDatabaseShardDatabaseWithExadbXdDetails:
		result["source"] = "XD_EXISTING_CLUSTER"

		if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}

		if v.AdminPasswordSecretId != nil {
			result["admin_password_secret_id"] = string(*v.AdminPasswordSecretId)
		}

		if v.AdminPasswordSecretVersionNumber != nil {
			result["admin_password_secret_version_number"] = int(*v.AdminPasswordSecretVersionNumber)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		result["transport_type"] = string(v.TransportType)

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.VmClusterId != nil {
			result["vm_cluster_id"] = string(*v.VmClusterId)
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedDbMetadataToMap(v.Metadata)}
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
	case oci_distributed_database.CreateDistributedDatabaseShardDatabaseWithExadbXsDetails:
		result["source"] = "XS_EXISTING_CLUSTER"

		if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}

		if v.AdminPasswordSecretId != nil {
			result["admin_password_secret_id"] = string(*v.AdminPasswordSecretId)
		}

		if v.AdminPasswordSecretVersionNumber != nil {
			result["admin_password_secret_version_number"] = int(*v.AdminPasswordSecretVersionNumber)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		result["transport_type"] = string(v.TransportType)

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.VmClusterId != nil {
			result["vm_cluster_id"] = string(*v.VmClusterId)
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedDbMetadataToMap(v.Metadata)}
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
	case oci_distributed_database.CreateDistributedDatabaseShardWithExadbXsNewClusterDetails:
		result["source"] = "XS_NEW_CLUSTER"

		if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}

		if v.AdminPasswordSecretId != nil {
			result["admin_password_secret_id"] = string(*v.AdminPasswordSecretId)
		}

		if v.AdminPasswordSecretVersionNumber != nil {
			result["admin_password_secret_version_number"] = int(*v.AdminPasswordSecretVersionNumber)
		}

		if v.AvailabilityDomain != nil {
			result["availability_domain"] = string(*v.AvailabilityDomain)
		}

		if v.ExascaleDbStorageVaultId != nil {
			result["exascale_db_storage_vault_id"] = string(*v.ExascaleDbStorageVaultId)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		result["transport_type"] = string(v.TransportType)

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.VmClusterDetails != nil {
			result["vm_cluster_details"] = []interface{}{VmClusterDetailsToMap(v.VmClusterDetails, false)}
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedDbMetadataToMap(v.Metadata)}
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
	case oci_distributed_database.CreateDistributedDatabaseShardWithExadbXsNewVaultAndClusterDetails:
		result["source"] = "XS_NEW_VAULT_AND_CLUSTER"

		if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}

		if v.AdminPasswordSecretId != nil {
			result["admin_password_secret_id"] = string(*v.AdminPasswordSecretId)
		}

		if v.AdminPasswordSecretVersionNumber != nil {
			result["admin_password_secret_version_number"] = int(*v.AdminPasswordSecretVersionNumber)
		}

		if v.AvailabilityDomain != nil {
			result["availability_domain"] = string(*v.AvailabilityDomain)
		}

		if v.DbStorageVaultDetails != nil {
			result["db_storage_vault_details"] = []interface{}{DbStorageVaultDetailsToMap(v.DbStorageVaultDetails)}
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = string(*v.KmsKeyVersionId)
		}

		result["protection_mode"] = string(v.ProtectionMode)

		result["transport_type"] = string(v.TransportType)

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}

		if v.VmClusterDetails != nil {
			result["vm_cluster_details"] = []interface{}{VmClusterDetailsToMap(v.VmClusterDetails, false)}
		}

		if v.Metadata != nil {
			result["metadata"] = []interface{}{DistributedDbMetadataToMap(v.Metadata)}
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
*/

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateSystemDataGuardDatabaseDetails(fieldKeyFormat string) (oci_distributed_database.CreateSystemDataGuardDatabaseDetails, error) {
	result := oci_distributed_database.CreateSystemDataGuardDatabaseDetails{}

	if dataGuardReplicas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")); ok {
		interfaces := dataGuardReplicas.([]interface{})
		tmp := make([]oci_distributed_database.CreateSystemDataGuardReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data_guard_replicas"), stateDataIndex)
			converted, err := s.mapToCreateSystemDataGuardReplicaDetails(fieldKeyFormatNextLevel)
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
			tmp, err := s.mapToCreateSystemDataGuardReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert original_replica, encountered error: %v", err)
			}
			result.OriginalReplica = &tmp
		}
	}

	return result, nil
}

func SystemDataGuardDatabaseToMap(obj *oci_distributed_database.SystemDataGuardDatabase) map[string]interface{} {
	result := map[string]interface{}{}

	dataGuardReplicas := []interface{}{}
	for _, item := range obj.DataGuardReplicas {
		dataGuardReplicas = append(dataGuardReplicas, SystemDataGuardReplicaToMap(item))
	}
	result["data_guard_replicas"] = dataGuardReplicas

	if obj.OriginalReplica != nil {
		result["original_replica"] = []interface{}{SystemDataGuardReplicaToMap(*obj.OriginalReplica)}
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateSystemDataGuardReplicaDetails(fieldKeyFormat string) (oci_distributed_database.CreateSystemDataGuardReplicaDetails, error) {
	result := oci_distributed_database.CreateSystemDataGuardReplicaDetails{}

	if databases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "databases")); ok {
		interfaces := databases.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedDatabaseShardDatabaseDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "databases"), stateDataIndex)
			converted, err := s.mapToCreateDistributedDatabaseShardDatabaseDetails(fieldKeyFormatNextLevel)
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

func SystemDataGuardReplicaToMap(obj oci_distributed_database.SystemDataGuardReplica) map[string]interface{} {
	result := map[string]interface{}{}

	databases := []interface{}{}
	for _, item := range obj.Databases {
		databases = append(databases, DistributedDatabaseShardDatabaseToMap(item))
	}
	result["databases"] = databases

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateSystemRaftClusterDetails(fieldKeyFormat string) (oci_distributed_database.CreateSystemRaftClusterDetails, error) {
	result := oci_distributed_database.CreateSystemRaftClusterDetails{}

	if databases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "databases")); ok {
		interfaces := databases.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedDatabaseShardDatabaseDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "databases"), stateDataIndex)
			converted, err := s.mapToCreateDistributedDatabaseShardDatabaseDetails(fieldKeyFormatNextLevel)
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
		result.RuMode = oci_distributed_database.CreateSystemRaftClusterDetailsRuModeEnum(ruMode.(string))
	}

	if witnessCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "witness_count")); ok {
		tmp := witnessCount.(int)
		result.WitnessCount = &tmp
	}

	return result, nil
}

func SystemRaftClusterToMap(obj oci_distributed_database.SystemRaftCluster) map[string]interface{} {
	result := map[string]interface{}{}

	databases := []interface{}{}
	for _, item := range obj.Databases {
		databases = append(databases, DistributedDatabaseShardDatabaseToMap(item))
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateUserShardSpaceDetails(fieldKeyFormat string) (oci_distributed_database.CreateUserShardSpaceDetails, error) {
	result := oci_distributed_database.CreateUserShardSpaceDetails{}

	if dataGuardReplicas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_guard_replicas")); ok {
		interfaces := dataGuardReplicas.([]interface{})
		tmp := make([]oci_distributed_database.CreateUserShardSpaceReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data_guard_replicas"), stateDataIndex)
			converted, err := s.mapToCreateUserShardSpaceReplicaDetails(fieldKeyFormatNextLevel)
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
			tmp, err := s.mapToCreateUserShardSpaceReplicaDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert original_replica, encountered error: %v", err)
			}
			result.OriginalReplica = &tmp
		}
	}

	return result, nil
}

func UserShardSpaceToMap(obj oci_distributed_database.UserShardSpace) map[string]interface{} {
	result := map[string]interface{}{}

	dataGuardReplicas := []interface{}{}
	for _, item := range obj.DataGuardReplicas {
		dataGuardReplicas = append(dataGuardReplicas, UserShardSpaceReplicaToMap(item))
	}
	result["data_guard_replicas"] = dataGuardReplicas

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.OriginalReplica != nil {
		result["original_replica"] = []interface{}{UserShardSpaceReplicaToMap(*obj.OriginalReplica)}
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToCreateUserShardSpaceReplicaDetails(fieldKeyFormat string) (oci_distributed_database.CreateUserShardSpaceReplicaDetails, error) {
	result := oci_distributed_database.CreateUserShardSpaceReplicaDetails{}

	if database, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database")); ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database"), 0)
			tmp, err := s.mapToCreateDistributedDatabaseShardDatabaseDetails(fieldKeyFormatNextLevel)
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

func UserShardSpaceReplicaToMap(obj oci_distributed_database.UserShardSpaceReplica) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Database != nil {
		databaseArray := []interface{}{}
		if databaseMap := DistributedDatabaseShardDatabaseToMap(obj.Database); databaseMap != nil {
			databaseArray = append(databaseArray, databaseMap)
		}
		result["database"] = databaseArray
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToDbStorageVaultDetails(fieldKeyFormat string) (oci_distributed_database.DbStorageVaultDetails, error) {
	result := oci_distributed_database.DbStorageVaultDetails{}

	if additionalFlashCacheInPercent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_flash_cache_in_percent")); ok {
		tmp := additionalFlashCacheInPercent.(int)
		result.AdditionalFlashCacheInPercent = &tmp
	}

	if autoscaleLimitInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "autoscale_limit_in_gbs")); ok {
		tmp := autoscaleLimitInGBs.(int)
		result.AutoscaleLimitInGBs = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if highCapacityDatabaseStorage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "high_capacity_database_storage")); ok {
		tmp := highCapacityDatabaseStorage.(int)
		result.HighCapacityDatabaseStorage = &tmp
	}

	if isAutoscaleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_autoscale_enabled")); ok {
		tmp := isAutoscaleEnabled.(bool)
		result.IsAutoscaleEnabled = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subscription_id")); ok {
		tmp := subscriptionId.(string)
		result.SubscriptionId = &tmp
	}

	return result, nil
}

func DbStorageVaultDetailsToMap(obj *oci_distributed_database.DbStorageVaultDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdditionalFlashCacheInPercent != nil {
		result["additional_flash_cache_in_percent"] = int(*obj.AdditionalFlashCacheInPercent)
	}

	if obj.AutoscaleLimitInGBs != nil {
		result["autoscale_limit_in_gbs"] = int(*obj.AutoscaleLimitInGBs)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.HighCapacityDatabaseStorage != nil {
		result["high_capacity_database_storage"] = int(*obj.HighCapacityDatabaseStorage)
	}

	if obj.IsAutoscaleEnabled != nil {
		result["is_autoscale_enabled"] = bool(*obj.IsAutoscaleEnabled)
	}

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	return result
}

func DistributedDatabaseGdsControlNodeToMap(obj oci_distributed_database.DistributedDatabaseGdsControlNode) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.GdsControlNodeImageDetails != nil {
		result["gds_control_node_image_details"] = []interface{}{DistributedDbGdsControlNodeImageToMap(obj.GdsControlNodeImageDetails)}
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{DistributedDbMetadataToMap(obj.Metadata)}
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToDistributedDbBackupDestination(fieldKeyFormat string) (oci_distributed_database.DistributedDbBackupDestination, error) {
	result := oci_distributed_database.DistributedDbBackupDestination{}

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
		result.Type = oci_distributed_database.DistributedDbBackupDestinationTypeEnum(type_.(string))
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

	if obj.RemoteRegion != nil {
		result["remote_region"] = string(*obj.RemoteRegion)
	}

	result["type"] = string(obj.Type)

	return result
}

func DistributedDbConnectionStringToMap(obj *oci_distributed_database.DistributedDbConnectionString) map[string]interface{} {
	result := map[string]interface{}{}

	result["all_connection_strings"] = obj.AllConnectionStrings

	return result
}

func DistributedDbGdsControlNodeImageToMap(obj *oci_distributed_database.DistributedDbGdsControlNodeImage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func DistributedDbGsmImageToMap(obj *oci_distributed_database.DistributedDbGsmImage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func DistributedDbMetadataToMap(obj *oci_distributed_database.DistributedDbMetadata) map[string]interface{} {
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToDistributedDbMetadata(fieldKeyFormat string) (oci_distributed_database.DistributedDbMetadata, error) {
	result := oci_distributed_database.DistributedDbMetadata{}
	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "map")); ok {
		result.PropertiesMap = tfresource.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}
	return result, nil
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

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) mapToVcnNsgIdsDetails(fieldKeyFormat string) (oci_distributed_database.VcnNsgIdsDetails, error) {
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

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
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

	if subscriptionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subscription_id")); ok {
		tmp := subscriptionId.(string)
		result.SubscriptionId = &tmp
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

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
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

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	if obj.TotalECpuCount != nil {
		result["total_ecpu_count"] = int(*obj.TotalECpuCount)
	}

	if obj.VmFileSystemStorageSize != nil {
		result["vm_file_system_storage_size"] = int(*obj.VmFileSystemStorageSize)
	}

	return result
}

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

func observedExascaleOperationalState(res *oci_distributed_database.DistributedDatabase) string {
	return observedOperationalState(observedExascaleTopology(res), func(metadata map[string]interface{}) string {
		switch strings.ToUpper(actionString(metadata["openMode"])) {
		case "READ_WRITE":
			return "AVAILABLE"
		case "MOUNTED":
			return "STOPPED"
		default:
			return "UNKNOWN"
		}
	})
}

func (s *DistributedDatabaseDistributedDatabaseResourceCrud) fetchOperationalState(ctx context.Context) (string, error) {
	if err := s.GetWithContext(ctx); err != nil {
		return "", err
	}
	state := observedExascaleOperationalState(s.Res)
	return state, nil
}

// Operational-state helpers shared by the Exascale and Autonomous resources.
// operational_state is action-owned; GET metadata verifies Start/Stop completion only.
// Include live node/catalog metadata for Start/Stop polling.
const operationalStateMetadataQuery = "VM_CLUSTER_INFO,ADDITIONAL_RESOURCE_INFO,RESOURCE_HIERARCHY"

func operationalStateSchema() *schema.Schema {
	return &schema.Schema{
		Type: schema.TypeString, Optional: true, Computed: true,
		Description:      "Desired operational state of an existing database: AVAILABLE to start or STOPPED to stop. Omit during creation and to leave start/stop unmanaged.",
		ValidateFunc:     validation.StringInSlice([]string{"AVAILABLE", "STOPPED"}, true),
		DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
	}
}

// preserveOperationalState keeps the action-owned state present across normal
// Terraform refreshes. Omitting a Computed field from Read clears it from
// state, which is different from preserving its prior value.
func preserveOperationalState(d *schema.ResourceData) error {
	if state, ok := d.GetOkExists("operational_state"); ok {
		return d.Set("operational_state", state)
	}
	return nil
}

type operationalStateData interface {
	GetRawConfig() cty.Value
	HasChange(string) bool
	GetChange(string) (interface{}, interface{})
	Id() string
}

// Read explicit configuration so computed values and removing the argument do
// not request an action. Deprecated counters still work and can be removed.
func operationalStateTarget(d operationalStateData) (string, error) {
	target := ""
	configured := false
	config := d.GetRawConfig()
	if !config.IsNull() && config.IsKnown() && config.Type().HasAttribute("operational_state") {
		value := config.GetAttr("operational_state")
		configured = !value.IsNull()
		if configured {
			if d.Id() == "" {
				return "", fmt.Errorf("omit operational_state during creation; set it after infrastructure provisioning")
			}
			if value.IsKnown() && d.HasChange("operational_state") {
				target = strings.ToUpper(value.AsString())
			}
		}
	}
	triggerTarget := ""
	for _, action := range []struct{ field, target string }{
		{"start_database_trigger", "AVAILABLE"}, {"stop_database_trigger", "STOPPED"},
	} {
		if !d.HasChange(action.field) {
			continue
		}
		if !config.IsNull() && config.IsKnown() && config.Type().HasAttribute(action.field) && !config.GetAttr(action.field).IsKnown() {
			continue
		}
		old, next := d.GetChange(action.field)
		if next.(int) == 0 {
			continue
		}
		if next.(int) <= old.(int) {
			return "", fmt.Errorf("%s must increase to run the action", action.field)
		}
		if configured || triggerTarget != "" {
			return "", fmt.Errorf("request start/stop using operational_state or one incremented counter, not both")
		}
		triggerTarget = action.target
	}
	if triggerTarget != "" {
		return triggerTarget, nil
	}
	return target, nil
}

func validateOperationalStateDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	// Replacement planning can evaluate a diff without prior resource state.
	// The Create handlers enforce the creation-only operational-state guard.
	if d.Id() == "" {
		return nil
	}
	_, err := operationalStateTarget(d)
	return err
}

func operationalStateConfigurationChanged(d *schema.ResourceData) bool {
	for _, field := range []string{"defined_tags", "freeform_tags", "display_name", "notification_topic_ids", "compartment_id"} {
		if d.HasChange(field) {
			return true
		}
	}
	for _, root := range databaseTopologyRoots {
		if d.HasChange(root) {
			return true
		}
	}
	return false
}

// Reuse the existing SDK-to-topology mappings. Only database nodes and catalog
// replicas participate; parent lifecycle, GSM and GDSCTL state are not checked.
func observedOperationalState(values databaseTopologyValues, nodeState func(map[string]interface{}) string) string {
	available, stopped, unknown, nodes, catalogs := 0, 0, false, 0, 0
	var walk func([]interface{}, bool)
	walk = func(items []interface{}, catalog bool) {
		for _, item := range items {
			value := actionMap(item)
			for _, field := range []string{"database", "databases"} {
				if raw, exists := value[field]; exists {
					databases := actionList(raw)
					if len(databases) == 0 {
						unknown = true
					}
					for _, database := range databases {
						// GET retains deleted nodes with empty metadata.
						if strings.EqualFold(actionString(actionMap(database)["status"]), "DELETED") {
							continue
						}
						if catalog {
							catalogs++
						} else {
							nodes++
						}
						metadata := actionMap(actionFirst(actionMap(database)["metadata"])["map"])
						state := nodeState(metadata)
						switch state {
						case "AVAILABLE":
							available++
						case "STOPPED":
							stopped++
						default:
							unknown = true
						}
					}
				}
			}
			for _, field := range []string{"raft_clusters", "original_replica", "data_guard_replicas"} {
				walk(actionList(value[field]), catalog)
			}
		}
	}
	for _, root := range databaseTopologyRoots {
		walk(values[root], root == "catalog_details")
	}
	if unknown || nodes == 0 || catalogs == 0 {
		return "UNKNOWN"
	}
	if available > 0 && stopped > 0 {
		return "MIXED"
	}
	if stopped > 0 {
		return "STOPPED"
	}
	return "AVAILABLE"
}

// The API determines whether Start/Stop is allowed. After it succeeds, poll DDB
// metadata until all nodes and catalogs reach the requested state/mode.
func waitForOperationalState(ctx context.Context, timeout time.Duration, target string, fetch func(context.Context) (string, error)) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	observed := ""
	for {
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("waiting for operational_state %s (last observed %s): %w", target, observed, err)
		}
		var err error
		observed, err = fetch(ctx)
		if err != nil {
			return err
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		if observed == target {
			return nil
		}
		timer := time.NewTimer(time.Second)
		select {
		case <-ctx.Done():
			timer.Stop()
			return fmt.Errorf("waiting for operational_state %s (last observed %s): %w", target, observed, ctx.Err())
		case <-timer.C:
		}
	}
}

// Preserve SDK status for all variants in the internal topology used for
// Start/Stop reads, including variants whose public flatteners omit status.
func operationalStateTopologyStatuses(values databaseTopologyValues, resource interface{}) databaseTopologyValues {
	encoded, err := json.Marshal(resource)
	if err != nil {
		return nil
	}
	var raw map[string]interface{}
	if err := json.Unmarshal(encoded, &raw); err != nil {
		return nil
	}
	fields := map[string]string{
		"catalog_details": "catalogDetails", "user_shard_spaces": "userShardSpaces",
		"system_raft_clusters": "systemRaftClusters", "composite_raft_shard_spaces": "compositeRaftShardSpaces",
		"system_data_guard_databases": "systemDataGuardDatabases", "composite_data_guard_shard_spaces": "compositeDataGuardShardSpaces",
		"raft_clusters": "raftClusters", "original_replica": "originalReplica", "data_guard_replicas": "dataGuardReplicas",
		"database": "database", "databases": "databases",
	}
	var walk func([]interface{}, interface{})
	walk = func(mapped []interface{}, source interface{}) {
		items := actionList(source)
		if object, ok := source.(map[string]interface{}); ok {
			items = []interface{}{object}
		}
		for index, item := range mapped {
			if index >= len(items) {
				continue
			}
			value, original := actionMap(item), actionMap(items[index])
			if value == nil || original == nil {
				continue
			}
			if status, ok := original["status"]; ok {
				value["status"] = status
			}
			for field, sdkField := range fields {
				if children, ok := value[field]; ok {
					walk(actionList(children), original[sdkField])
				}
			}
		}
	}
	for _, root := range databaseTopologyRoots {
		walk(values[root], raw[fields[root]])
	}
	return values
}

func observedExascaleTopology(res *oci_distributed_database.DistributedDatabase) databaseTopologyValues {
	values := databaseTopologyValues{}
	for _, item := range res.CatalogDetails {
		values["catalog_details"] = append(values["catalog_details"], DistributedDatabaseCatalogToMap(item))
	}
	for _, item := range res.UserShardSpaces {
		values["user_shard_spaces"] = append(values["user_shard_spaces"], UserShardSpaceToMap(item))
	}
	for _, item := range res.SystemRaftClusters {
		values["system_raft_clusters"] = append(values["system_raft_clusters"], SystemRaftClusterToMap(item))
	}
	for _, item := range res.CompositeRaftShardSpaces {
		values["composite_raft_shard_spaces"] = append(values["composite_raft_shard_spaces"], CompositeRaftShardSpaceToMap(item))
	}
	if res.SystemDataGuardDatabases != nil {
		values["system_data_guard_databases"] = []interface{}{SystemDataGuardDatabaseToMap(res.SystemDataGuardDatabases)}
	}
	for _, item := range res.CompositeDataGuardShardSpaces {
		values["composite_data_guard_shard_spaces"] = append(values["composite_data_guard_shard_spaces"], CompositeDataGuardShardSpaceToMap(item))
	}
	return operationalStateTopologyStatuses(values, res)
}

// Backup configuration helpers shared by the Exascale and Autonomous resources.
const backupConfigTriggerField = "change_db_backup_config_trigger"

// Backup settings are mutable through the backup-change APIs.
// Apply configured backup diffs in Update, never in Create.
func configureBackupConfigSchema(resource *schema.Resource, autonomous bool) {
	config := resource.Schema["db_backup_config"]
	clearBackupConfigForceNew(config)
	fields := config.Elem.(*schema.Resource).Schema
	minimum, maximum := 1, 60
	if autonomous {
		minimum, maximum = 0, 95
	}
	fields["recovery_window_in_days"].ValidateFunc = validation.IntBetween(minimum, maximum)
	destinations := fields["backup_destination_details"]
	destinations.MaxItems = 2
	destinations.Elem.(*schema.Resource).Schema["type"].ValidateFunc = validation.StringInSlice([]string{"NFS", "OBJECT_STORE", "LOCAL", "DBRS"}, false)
	resource.Schema[backupConfigTriggerField].Deprecated = "Backup configuration changes are applied automatically. Remove this trigger and edit db_backup_config instead."
}

func clearBackupConfigForceNew(field *schema.Schema) {
	field.ForceNew = false
	if nested, ok := field.Elem.(*schema.Resource); ok {
		for _, child := range nested.Schema {
			clearBackupConfigForceNew(child)
		}
	}
}

type backupConfigData interface {
	Id() string
	GetRawConfig() cty.Value
	HasChange(string) bool
	GetChange(string) (interface{}, interface{})
}

func backupConfigConfigured(config cty.Value) bool {
	if config.IsNull() || !config.IsKnown() || !config.Type().HasAttribute("db_backup_config") {
		return false
	}
	backup := config.GetAttr("db_backup_config")
	return backup.IsKnown() && !backup.IsNull() && backup.LengthInt() == 1
}

func backupConfigLegacyTriggerConfigured(config cty.Value) bool {
	if config.IsNull() || !config.IsKnown() || !config.Type().HasAttribute(backupConfigTriggerField) {
		return false
	}
	trigger := config.GetAttr(backupConfigTriggerField)
	return trigger.IsKnown() && !trigger.IsNull()
}

// Omitted Optional/Computed backup settings are readback, not an update request.
// Legacy counter increments can still repeat an action with configured settings.
func backupConfigUpdateRequested(d backupConfigData) bool {
	config := d.GetRawConfig()
	if d.Id() == "" || !backupConfigConfigured(config) {
		return false
	}
	if d.HasChange("db_backup_config") {
		return true
	}
	if backupConfigLegacyTriggerConfigured(config) && d.HasChange(backupConfigTriggerField) {
		old, next := d.GetChange(backupConfigTriggerField)
		return next.(int) > old.(int)
	}
	return false
}

func validateBackupConfigUpdate(d backupConfigData) error {
	config := d.GetRawConfig()
	if d.Id() == "" || !backupConfigLegacyTriggerConfigured(config) || !d.HasChange(backupConfigTriggerField) {
		return nil
	}
	oldRaw, nextRaw := d.GetChange(backupConfigTriggerField)
	old, next := oldRaw.(int), nextRaw.(int)
	if next == 0 {
		return nil // Removing the deprecated counter must not run an action.
	}
	if next <= old {
		return fmt.Errorf("%s must increase to repeat the backup action", backupConfigTriggerField)
	}
	if !backupConfigConfigured(config) {
		if config.Type().HasAttribute("db_backup_config") && !config.GetAttr("db_backup_config").IsKnown() {
			return nil
		}
		return fmt.Errorf("db_backup_config is required when the deprecated %s increases", backupConfigTriggerField)
	}
	return nil
}

func validateBackupConfigDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	return validateBackupConfigUpdate(d)
}

// Project raw configuration into the special-update request so omitted values
// are not echoed from computed state or filled with Terraform zero values.
func decodeBackupConfig(config cty.Value, target interface{}) error {
	if !backupConfigConfigured(config) {
		return fmt.Errorf("db_backup_config is required for Change DB Backup Config")
	}
	backup := config.GetAttr("db_backup_config")
	if !backup.IsWhollyKnown() {
		return fmt.Errorf("db_backup_config must be known before Change DB Backup Config")
	}
	value := backup.Index(cty.NumberIntVal(0))
	encoded, err := ctyjson.Marshal(value, value.Type())
	if err != nil {
		return fmt.Errorf("encoding backup configuration: %w", err)
	}
	var configured interface{}
	if err := json.Unmarshal(encoded, &configured); err != nil {
		return fmt.Errorf("reading backup configuration: %w", err)
	}
	encoded, err = json.Marshal(backupConfigJSONValue(configured))
	if err != nil {
		return fmt.Errorf("encoding backup update request: %w", err)
	}
	if err := json.Unmarshal(encoded, target); err != nil {
		return fmt.Errorf("mapping backup update request: %w", err)
	}
	return nil
}

func backupConfigJSONValue(value interface{}) interface{} {
	switch value := value.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{}, len(value))
		for key, item := range value {
			if item == nil {
				continue
			}
			parts := strings.Split(key, "_")
			for i := 1; i < len(parts); i++ {
				parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
			}
			result[strings.Join(parts, "")] = backupConfigJSONValue(item)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(value))
		for i, item := range value {
			result[i] = backupConfigJSONValue(item)
		}
		return result
	default:
		return value
	}
}
