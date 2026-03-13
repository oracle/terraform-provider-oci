// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedAutonomousDatabaseResource() *schema.Resource {
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
						// NOTE: Codegen issue
						//
						// admin_password is a write-only field in the OCI API.
						// It is required on create but never returned in GET/LIST responses.
						// The generated schema expects round-trip state consistency, which
						// results in perpetual Terraform drift and forced recreation.
						//
						// DiffSuppressFunc is applied here as a workaround to suppress
						// false-positive diffs after the resource has been created.
						// JIRA: TOP-9459
						// Required
						"admin_password": {
							Type:      schema.TypeString,
							Required:  true,
							ForceNew:  true,
							Sensitive: true,
							/*Optional:         true,
							Sensitive:        true,
							Computed:         true,*/
							DiffSuppressFunc: suppressMaskedPasswordDiff,
						},
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
						"is_auto_scaling_enabled": {
							Type:     schema.TypeBool,
							Required: true,
							ForceNew: true,
						},
						"source": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ADB_D",
							}, true),
						},

						// Optional
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
						"peer_cloud_autonomous_vm_cluster_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"peer_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"cloud_autonomous_vm_cluster_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"fast_start_fail_over_lag_limit_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_automatic_failover_enabled": {
										Type:     schema.TypeBool,
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
													//Elem: &schema.Schema{Type: schema.TypeString},
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
										//Elem: &schema.Schema{Type: schema.TypeString},
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
			"db_workload": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// NOTE (codegen bug): TOP-9449
			// The OCI CreateDistributedAutonomousDatabase API does NOT accept a resource OCID as input.
			// The service generates the OCID and returns it in the response (and/or via work request).
			// The code generator incorrectly modeled the response `id` as a Required schema argument.
			// Terraform must manage the OCID via `d.SetId(...)` during Create and use `d.Id()` for
			// subsequent Get/Update/Delete calls.
			/*"distributed_autonomous_database_id": {
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
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						// NOTE: Codegen issue
						//
						// admin_password is a write-only field in the OCI API.
						// It is required on create but never returned in GET/LIST responses.
						// The generated schema expects round-trip state consistency, which
						// results in perpetual Terraform drift and forced recreation.
						//
						// DiffSuppressFunc is applied here as a workaround to suppress
						// false-positive diffs after the resource has been created.
						// JIRA: TOP-9459
						"admin_password": {
							Type:      schema.TypeString,
							Required:  true,
							ForceNew:  true,
							Sensitive: true,
							/*Optional:         true,
							Sensitive:        true,
							Computed:         true,*/
							DiffSuppressFunc: suppressMaskedPasswordDiff,
						},
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
						"is_auto_scaling_enabled": {
							Type:     schema.TypeBool,
							Required: true,
							ForceNew: true,
						},
						"source": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ADB_D",
							}, true),
						},

						// Optional
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
						"peer_cloud_autonomous_vm_cluster_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"peer_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"cloud_autonomous_vm_cluster_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"fast_start_fail_over_lag_limit_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_automatic_failover_enabled": {
										Type:     schema.TypeBool,
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

									// Computed
									"container_database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									/*
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
														//Elem: &schema.Schema{Type: schema.TypeString},
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

						// Computed
						"container_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						/*
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
											//Elem: &schema.Schema{Type: schema.TypeString},
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
				//ForceNew: true,
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
							//ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
										//ForceNew: true,
									},

									// Optional
									"dbrs_policy_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										//ForceNew: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										//ForceNew: true,
									},
									"internet_proxy": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										//ForceNew: true,
									},
									"is_remote": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										//ForceNew: true,
									},
									"remote_region": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										//ForceNew: true,
									},
									"vpc_password": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										//ForceNew:  true,
										Sensitive: true,
										// API may omit write-only password in GET; suppress
										// post-create drift when config keeps a non-empty value.
										DiffSuppressFunc: suppressMaskedPasswordDiff,
									},
									"vpc_user": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										//ForceNew: true,
									},

									// Computed
								},
							},
						},
						"recovery_window_in_days": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							//ForceNew: true,
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
							Required:         true,
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
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateInactive),
					string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive),
				}, true),
			},
			"change_db_backup_config_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// NOTE (JIRA: TOP-9494):
			// Codegen gap: ChangeDistributedAutonomousDbBackupConfig action requires DbBackupConfig in the payload
			// (ChangeDistributedAutonomousDbBackupConfigDetails.dbBackupConfig), but schema did not expose an
			// action-specific input. Added `change_db_backup_config` to make the action usable from Terraform.
			/*"change_db_backup_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// use the same schema fields you already use under db_backup_config
						// (whatever DistributedAutonomousDbBackupConfig supports)
					},
				},
			},*/
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
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed output populated by the downloadGsmCertificateSigningRequest action.
			// Holds the full PEM CSR returned by the service.
			// See JIRA: TOP-9481
			"downloaded_gsm_csr_pem": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"generate_gsm_certificate_signing_request_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// The generateGsmCertificateSigningRequest action supports an optional
			// caBundleId query parameter, which is not generated into the Terraform
			// schema by default.
			//
			// This parameter is required to support CA-specific CSR generation.
			// See JIRA: TOP-9478
			"generate_gsm_certificate_signing_request_trigger_ca_bundle_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			"upload_signed_certificate_and_generate_wallet_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// NOTE (CODEGEN GAP):
			// The action API uploadSignedCertificateAndGenerateWallet requires a CA-signed certificate payload,
			// but the generator did not include an input attribute to pass it from Terraform.
			// See JIRA: TOP-9482
			"upload_ca_signed_certificate": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
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
				Description: "Increment this value to trigger StartDistributedAutonomousDatabase action.",
			},
			"stop_database_trigger": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Increment this value to trigger StopDistributedAutonomousDatabase action.",
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
						/*
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
											//Elem: &schema.Schema{Type: schema.TypeString},
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
							//Elem:     schema.TypeString,
							Elem: &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},*/
			/*
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
								//Elem: &schema.Schema{Type: schema.TypeString},
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

/*
func createDistributedDatabaseDistributedAutonomousDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	return nil

}*/
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
func createDistributedDatabaseDistributedAutonomousDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	// Block CREATE + action triggers (triggers are TypeInt; allow only 0 / unset)
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

func readDistributedDatabaseDistributedAutonomousDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

/*
func updateDistributedDatabaseDistributedAutonomousDatabaseWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	actionInvoked := false
	needsUpdate :=
    d.HasChange("display_name") ||
    d.HasChange("freeform_tags") ||
    d.HasChange("defined_tags") ||
    d.HasChange("compartment_id") // if you treat it via updateCompartment

	if needsUpdate {
    if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
        return tfresource.HandleDiagError(m, err)
    }
	} else if actionInvoked {
    // No PUT-able changes; action already performed + waited inside action handler
    return nil
	}

	/*powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateInactive == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}*/

//if powerOn {
// NOTE:
// This function contains multiple workarounds for known Terraform code
// generator issues (incorrect return types, missing context propagation).
// These changes are required for compilation and correctness but should
// be removed once the generator is fixed.
// See JIRA: TOP-9394
/*
	if err := sync.StartDistributedAutonomousDatabase(); err != nil {
		return tfresource.HandleDiagError(m, err)
	}*/
/*if err := sync.StartDistributedAutonomousDatabase(ctx); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	sync.D.Set("state", oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive)
}*/

// CODEGEN FIX: TOP-9474
// Start/Stop are imperative action APIs (not desired-state changes).
// Terraform cannot infer when to run them unless a trigger changes.
// Use monotonic int triggers to make the action explicit and idempotent.
/*if d.HasChange("start_database_trigger") && d.HasChange("stop_database_trigger") {
	log.Printf("start_database_trigger and stop_database_trigger cannot be changed in the same apply")
}*/

// START trigger
/*if oldRaw, newRaw := d.GetChange("start_database_trigger"); d.HasChange("start_database_trigger") {
	oldV := oldRaw.(int)
	newV := newRaw.(int)
	if newV <= oldV {
		log.Printf("start_database_trigger must be incremented to retrigger start action (old=%d new=%d)", oldV, newV)
		return nil
	}
	if err := sync.StartDistributedAutonomousDatabase(ctx); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	actionInvoked = true
	// Preserve trigger in state
	_ = d.Set("start_database_trigger", newV)
}*/

// STOP trigger
/*if oldRaw, newRaw := d.GetChange("stop_database_trigger"); d.HasChange("stop_database_trigger") {
	oldV := oldRaw.(int)
	newV := newRaw.(int)
	if newV <= oldV {
		log.Printf("stop_database_trigger must be incremented to retrigger stop action (old=%d new=%d)", oldV, newV)
		return nil
	}
	if err := sync.StopDistributedAutonomousDatabase(ctx); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	actionInvoked = true
	// Preserve trigger in state
	_ = d.Set("stop_database_trigger", newV)
}

if _, ok := sync.D.GetOkExists("change_db_backup_config_trigger"); ok && sync.D.HasChange("change_db_backup_config_trigger") {
	oldRaw, newRaw := sync.D.GetChange("change_db_backup_config_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.ChangeDistributedAutonomousDbBackupConfig()
		actionInvoked = true
		if err != nil {
			// WORKAROUND:
			// This code is generated and incorrectly returns `error` from Context-based
			// CRUD functions that must return `diag.Diagnostics`.
			//
			// The generator emits `return err` in multiple places, which causes a
			// compile-time type mismatch with Terraform Plugin SDK v2.
			//
			// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
			// DO NOT remove unless the generator is updated.
			// See JIRA: TOP-9389
			// return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("change_db_backup_config_trigger", oldRaw)
		// WORKAROUND:
		// This code is generated and incorrectly returns `error` from Context-based
		// CRUD functions that must return `diag.Diagnostics`.
		//
		// The generator emits `return err` in multiple places, which causes a
		// compile-time type mismatch with Terraform Plugin SDK v2.
		//
		// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
		// DO NOT remove unless the generator is updated.
		// See JIRA: TOP-9389
		// return err
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

if _, ok := sync.D.GetOkExists("configure_sharding_trigger"); ok && sync.D.HasChange("configure_sharding_trigger") {
	oldRaw, newRaw := sync.D.GetChange("configure_sharding_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.ConfigureDistributedAutonomousDatabaseSharding()
		actionInvoked = true
		if err != nil {
			// WORKAROUND:
			// This code is generated and incorrectly returns `error` from Context-based
			// CRUD functions that must return `diag.Diagnostics`.
			//
			// The generator emits `return err` in multiple places, which causes a
			// compile-time type mismatch with Terraform Plugin SDK v2.
			//
			// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
			// DO NOT remove unless the generator is updated.
			// See JIRA: TOP-9389
			// return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("configure_sharding_trigger", oldRaw)
		// WORKAROUND:
		// This code is generated and incorrectly returns `error` from Context-based
		// CRUD functions that must return `diag.Diagnostics`.
		//
		// The generator emits `return err` in multiple places, which causes a
		// compile-time type mismatch with Terraform Plugin SDK v2.
		//
		// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
		// DO NOT remove unless the generator is updated.
		// See JIRA: TOP-9389
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

if _, ok := sync.D.GetOkExists("download_gsm_certificate_signing_request_trigger"); ok && sync.D.HasChange("download_gsm_certificate_signing_request_trigger") {
	oldRaw, newRaw := sync.D.GetChange("download_gsm_certificate_signing_request_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest()
		actionInvoked = true
		if err != nil {
			// WORKAROUND:
			// This code is generated and incorrectly returns `error` from Context-based
			// CRUD functions that must return `diag.Diagnostics`.
			//
			// The generator emits `return err` in multiple places, which causes a
			// compile-time type mismatch with Terraform Plugin SDK v2.
			//
			// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
			// DO NOT remove unless the generator is updated.
			// See JIRA: TOP-9389
			// return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("download_gsm_certificate_signing_request_trigger", oldRaw)
		// WORKAROUND:
		// This code is generated and incorrectly returns `error` from Context-based
		// CRUD functions that must return `diag.Diagnostics`.
		//
		// The generator emits `return err` in multiple places, which causes a
		// compile-time type mismatch with Terraform Plugin SDK v2.
		//
		// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
		// DO NOT remove unless the generator is updated.
		// See JIRA: TOP-9389
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

if _, ok := sync.D.GetOkExists("generate_gsm_certificate_signing_request_trigger"); ok && sync.D.HasChange("generate_gsm_certificate_signing_request_trigger") {
	oldRaw, newRaw := sync.D.GetChange("generate_gsm_certificate_signing_request_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest()
		actionInvoked = true
		if err != nil {
			// WORKAROUND:
			// This code is generated and incorrectly returns `error` from Context-based
			// CRUD functions that must return `diag.Diagnostics`.
			//
			// The generator emits `return err` in multiple places, which causes a
			// compile-time type mismatch with Terraform Plugin SDK v2.
			//
			// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
			// DO NOT remove unless the generator is updated.
			// See JIRA: TOP-9389
			// return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("generate_gsm_certificate_signing_request_trigger", oldRaw)
		// WORKAROUND:
		// This code is generated and incorrectly returns `error` from Context-based
		// CRUD functions that must return `diag.Diagnostics`.
		//
		// The generator emits `return err` in multiple places, which causes a
		// compile-time type mismatch with Terraform Plugin SDK v2.
		//
		// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
		// DO NOT remove unless the generator is updated.
		// See JIRA: TOP-9389
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

if _, ok := sync.D.GetOkExists("generate_wallet_trigger"); ok && sync.D.HasChange("generate_wallet_trigger") {
	oldRaw, newRaw := sync.D.GetChange("generate_wallet_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.GenerateDistributedAutonomousDatabaseWallet()
		actionInvoked = true
		if err != nil {
			// WORKAROUND:
			// This code is generated and incorrectly returns `error` from Context-based
			// CRUD functions that must return `diag.Diagnostics`.
			//
			// The generator emits `return err` in multiple places, which causes a
			// compile-time type mismatch with Terraform Plugin SDK v2.
			//
			// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
			// DO NOT remove unless the generator is updated.
			// See JIRA: TOP-9389
			// return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("generate_wallet_trigger", oldRaw)
		// WORKAROUND:
		// This code is generated and incorrectly returns `error` from Context-based
		// CRUD functions that must return `diag.Diagnostics`.
		//
		// The generator emits `return err` in multiple places, which causes a
		// compile-time type mismatch with Terraform Plugin SDK v2.
		//
		// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
		// DO NOT remove unless the generator is updated.
		// See JIRA: TOP-9389
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

if _, ok := sync.D.GetOkExists("upload_signed_certificate_and_generate_wallet_trigger"); ok && sync.D.HasChange("upload_signed_certificate_and_generate_wallet_trigger") {
	oldRaw, newRaw := sync.D.GetChange("upload_signed_certificate_and_generate_wallet_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet()
		actionInvoked = true
		if err != nil {
			// WORKAROUND:
			// This code is generated and incorrectly returns `error` from Context-based
			// CRUD functions that must return `diag.Diagnostics`.
			//
			// The generator emits `return err` in multiple places, which causes a
			// compile-time type mismatch with Terraform Plugin SDK v2.
			//
			// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
			// DO NOT remove unless the generator is updated.
			// See JIRA: TOP-9389
			// return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("upload_signed_certificate_and_generate_wallet_trigger", oldRaw)
		// WORKAROUND:
		// This code is generated and incorrectly returns `error` from Context-based
		// CRUD functions that must return `diag.Diagnostics`.
		//
		// The generator emits `return err` in multiple places, which causes a
		// compile-time type mismatch with Terraform Plugin SDK v2.
		//
		// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
		// DO NOT remove unless the generator is updated.
		// See JIRA: TOP-9389
		//return fmt.Errorf("new value of trigger should be greater than the old value")
		return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
	}
}

if _, ok := sync.D.GetOkExists("validate_network_trigger"); ok && sync.D.HasChange("validate_network_trigger") {
	oldRaw, newRaw := sync.D.GetChange("validate_network_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if oldValue < newValue {
		err := sync.ValidateDistributedAutonomousDatabaseNetwork()
		actionInvoked = true
		if err != nil {
			// WORKAROUND:
			// This code is generated and incorrectly returns `error` from Context-based
			// CRUD functions that must return `diag.Diagnostics`.
			//
			// The generator emits `return err` in multiple places, which causes a
			// compile-time type mismatch with Terraform Plugin SDK v2.
			//
			// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
			// DO NOT remove unless the generator is updated.
			// See JIRA: TOP-9389
			// return err
			return tfresource.HandleDiagError(m, err)
		}
	} else {
		sync.D.Set("validate_network_trigger", oldRaw)
		// WORKAROUND:
		// This code is generated and incorrectly returns `error` from Context-based
		// CRUD functions that must return `diag.Diagnostics`.
		//
		// The generator emits `return err` in multiple places, which causes a
		// compile-time type mismatch with Terraform Plugin SDK v2.
		//
		// Errors are wrapped here using tfresource.HandleDiagError as a temporary fix.
		// DO NOT remove unless the generator is updated.
		// See JIRA: TOP-9389
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
/*if err := tfresource.UpdateResourceWithContext(d, sync); err != nil {
	return tfresource.HandleDiagError(m, err)
}*/
/*if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
	return tfresource.HandleDiagError(m, err)
}*/

/*if powerOff {
// NOTE:
// This function contains multiple workarounds for known Terraform code
// generator issues (incorrect return types, missing context propagation).
// These changes are required for compilation and correctness but should
// be removed once the generator is fixed.
// See JIRA: TOP-9394
/*
	if err := sync.StopDistributedAutonomousDatabase(); err != nil {
		return tfresource.HandleDiagError(m, err)
	}*/
/*if err := sync.StopDistributedAutonomousDatabase(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		sync.D.Set("state", oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateInactive)
	}

	return nil
}*/

// Action triggers are increment-only and intentionally bypass
// UpdateResourceWithContext to avoid invalid PUT calls
// while the resource is in UPDATING state.
// See JIRA: TOP-9479

/*func updateDistributedDatabaseDistributedAutonomousDatabaseWithContext(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	// 1) Compute whether a PUT update is needed (fields handled by UpdateWithContext)
	/*needsUpdate :=
		d.HasChange("display_name") ||
			d.HasChange("freeform_tags") ||
			d.HasChange("defined_tags") ||
			d.HasChange("compartment_id")*/

// 1) Detect which categories of updates are needed
/*needsCompartmentMove := s.D.HasChange("compartment_id")

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

	// 2) Compute whether any action trigger is going to run (increment-only triggers)
	actionInvoked := false
	hasActionTrigger := false

	// Helper: increment-only trigger check
	triggerBumped := func(attr string) (bool, int, int) {
		if _, ok := d.GetOkExists(attr); !ok || !d.HasChange(attr) {
			return false, 0, 0
		}
		oldRaw, newRaw := d.GetChange(attr)
		oldV := oldRaw.(int)
		newV := newRaw.(int)
		if newV <= oldV {
			// Keep old value to avoid drift
			_ = d.Set(attr, oldV)
			return false, oldV, newV
		}
		return true, oldV, newV
	}

	// Pre-scan all triggers (so we can decide whether to skip PUT)
	triggerAttrs := []string{
		"start_database_trigger",
		"stop_database_trigger",
		"change_db_backup_config_trigger",
		"configure_sharding_trigger",
		"download_gsm_certificate_signing_request_trigger",
		"generate_gsm_certificate_signing_request_trigger",
		"generate_wallet_trigger",
		"upload_signed_certificate_and_generate_wallet_trigger",
		"validate_network_trigger",
	}
	for _, a := range triggerAttrs {
		if ok, _, _ := triggerBumped(a); ok {
			hasActionTrigger = true
			break
		}
	}

	// 3) If there are PUT-able changes, do them FIRST.
	// This avoids action -> immediate PUT while resource is UPDATING.
	if needsPutUpdate {
		if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	// 4) Run action triggers (action handlers should wait on WR + refresh state).
	// START
	if ok, _, newV := triggerBumped("start_database_trigger"); ok {
		if err := sync.StartDistributedAutonomousDatabase(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("start_database_trigger", newV)
	}

	// STOP
	if ok, _, newV := triggerBumped("stop_database_trigger"); ok {
		if err := sync.StopDistributedAutonomousDatabase(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("stop_database_trigger", newV)
	}

	// generate_gsm_certificate_signing_request_trigger
	if ok, _, newV := triggerBumped("generate_gsm_certificate_signing_request_trigger"); ok {
		if err := sync.GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("generate_gsm_certificate_signing_request_trigger", newV)
	}

	// change_db_backup_config_trigger
	if ok, _, newV := triggerBumped("change_db_backup_config_trigger"); ok {
		if err := sync.ChangeDistributedAutonomousDbBackupConfig(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("change_db_backup_config_trigger", newV)
	}

	//configure_sharding_trigger
	if ok, _, newV := triggerBumped("configure_sharding_trigger"); ok {
		if err := sync.ConfigureDistributedAutonomousDatabaseSharding(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("configure_sharding_trigger", newV)
	}

	//download_gsm_certificate_signing_request_trigger
	if ok, _, newV := triggerBumped("download_gsm_certificate_signing_request_trigger"); ok {
		if err := sync.DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("download_gsm_certificate_signing_request_trigger", newV)
	}

	//generate_wallet_trigger
	if ok, _, newV := triggerBumped("generate_wallet_trigger"); ok {
		if err := sync.GenerateDistributedAutonomousDatabaseWallet(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("generate_wallet_trigger", newV)
	}

	//upload_signed_certificate_and_generate_wallet_trigger
	if ok, _, newV := triggerBumped("upload_signed_certificate_and_generate_wallet_trigger"); ok {
		if err := sync.UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("upload_signed_certificate_and_generate_wallet_trigger", newV)
	}

	// validate_network_trigger
	if ok, _, newV := triggerBumped("validate_network_trigger"); ok {
		if err := sync.ValidateDistributedAutonomousDatabaseNetwork(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("validate_network_trigger", newV)
	}

	// 5) If this was an action-only update, we must NOT fall through into any generic update call.
	// (In this pattern we already ran PUT earlier only if needsUpdate==true.)
	if hasActionTrigger && !needsPutUpdate && actionInvoked {
		return nil
	}

	return nil
}*/
// NOTE (JIRA: TOP-9479):
// Action triggers are increment-only and intentionally bypass additional PUT updates
// to avoid invalid PUT calls while the resource is in UPDATING state (service rejects).
//
// NOTE (JIRA: TOP-9460):
// patch_operations must be applied only when changed AND non-empty to avoid sending
// an empty PATCH payload (items: []) which the service rejects.

// TOP-9493: Refactored to first handle standard updates (PUT/PATCH/compartment move),
func updateDistributedDatabaseDistributedAutonomousDatabaseWithContext(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

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
		if err := sync.StartDistributedAutonomousDatabase(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("start_database_trigger", newV)
	}

	// STOP
	if ok, _, newV := triggerBumped("stop_database_trigger"); ok {
		if err := sync.StopDistributedAutonomousDatabase(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("stop_database_trigger", newV)
	}

	// Generate GSM CSR
	if ok, _, newV := triggerBumped("generate_gsm_certificate_signing_request_trigger"); ok {
		if err := sync.GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("generate_gsm_certificate_signing_request_trigger", newV)
	}

	// Change backup config
	if ok, _, newV := triggerBumped("change_db_backup_config_trigger"); ok {
		if err := sync.ChangeDistributedAutonomousDbBackupConfig(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("change_db_backup_config_trigger", newV)
	}

	// Configure sharding
	if ok, _, newV := triggerBumped("configure_sharding_trigger"); ok {
		if err := sync.ConfigureDistributedAutonomousDatabaseSharding(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("configure_sharding_trigger", newV)
	}

	// Download GSM CSR
	if ok, _, newV := triggerBumped("download_gsm_certificate_signing_request_trigger"); ok {
		if err := sync.DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("download_gsm_certificate_signing_request_trigger", newV)
	}

	// Generate wallet
	if ok, _, newV := triggerBumped("generate_wallet_trigger"); ok {
		if err := sync.GenerateDistributedAutonomousDatabaseWallet(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("generate_wallet_trigger", newV)
	}

	// Upload signed cert + generate wallet
	if ok, _, newV := triggerBumped("upload_signed_certificate_and_generate_wallet_trigger"); ok {
		if err := sync.UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		actionInvoked = true
		_ = d.Set("upload_signed_certificate_and_generate_wallet_trigger", newV)
	}

	// Validate network
	if ok, _, newV := triggerBumped("validate_network_trigger"); ok {
		if err := sync.ValidateDistributedAutonomousDatabaseNetwork(ctx); err != nil {
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
		string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive),
		string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateNeedsAttention),
		// NOTE:
		// Distributed Autonomous Database creation is asynchronous.
		// After a successful Create operation, the service intentionally returns
		// the resource in lifecycle state INACTIVE.
		//
		// INACTIVE is a valid terminal post-create state. The resource transitions
		// to ACTIVE only after specific action APIs (e.g. activation/enablement)
		// are executed successfully.
		//
		// The provider MUST treat INACTIVE as a successful Create completion state.
		// Waiting for ACTIVE during Create is incorrect and causes false failures.
		// JIRA 9465
		string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateInactive),
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
		tmp := make([]oci_distributed_database.CreateDistributedAutonomousDatabaseCatalogDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "catalog_details", stateDataIndex)
			converted, err := s.mapToCreateDistributedAutonomousDatabaseCatalogDetails(fieldKeyFormat)
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
			tmp, err := s.mapToDistributedAutonomousDbBackupConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DbBackupConfig = &tmp
		}
	}

	if dbDeploymentType, ok := s.D.GetOkExists("db_deployment_type"); ok {
		request.DbDeploymentType = oci_distributed_database.CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum(dbDeploymentType.(string))
	}

	if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
		request.DbWorkload = oci_distributed_database.CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum(dbWorkload.(string))
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
		request.ReplicationMethod = oci_distributed_database.CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum(replicationMethod.(string))
	}

	if replicationUnit, ok := s.D.GetOkExists("replication_unit"); ok {
		tmp := replicationUnit.(int)
		request.ReplicationUnit = &tmp
	}

	if shardDetails, ok := s.D.GetOkExists("shard_details"); ok {
		interfaces := shardDetails.([]interface{})
		tmp := make([]oci_distributed_database.CreateDistributedAutonomousDatabaseShardDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shard_details", stateDataIndex)
			converted, err := s.mapToCreateDistributedAutonomousDatabaseShardDetails(fieldKeyFormat)
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
		request.ShardingMethod = oci_distributed_database.CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum(shardingMethod.(string))
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

	err = s.getDistributedAutonomousDatabaseFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The code generator omits the required context.Context argument when calling
	// getDistributedAutonomousDatabaseFromWorkRequest, resulting in a compile-time error.
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
	// See  Jira: TOP-9460
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
// getDistributedAutonomousDatabaseFromWorkRequest, resulting in a compile-time error.
// This call explicitly propagates `ctx`.
// See JIRA: TOP-9396
/*
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) Patch() error {
request := oci_distributed_database.PatchDistributedAutonomousDatabaseRequest{}
*/

/*func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) Patch(ctx context.Context) error {
request := oci_distributed_database.PatchDistributedAutonomousDatabaseRequest{}*/

/*if distributedAutonomousDatabaseId, ok := s.D.GetOkExists("id"); ok {
	tmp := distributedAutonomousDatabaseId.(string)
	request.DistributedAutonomousDatabaseId = &tmp
}*/

/*id := s.D.Id()
	request.DistributedAutonomousDatabaseId = &id

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
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The code generator omits the required context.Context argument when calling
	// getDistributedAutonomousDatabaseFromWorkRequest, resulting in a compile-time error.
	// This call explicitly propagates `ctx`.
	// See JIRA: TOP-9396
	// return s.getDistributedAutonomousDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
	return s.getDistributedAutonomousDatabaseFromWorkRequest(
		ctx,
		workId,
		tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"),
		oci_distributed_database.ActionTypeUpdated,
		s.D.Timeout(schema.TimeoutUpdate),
	)
}*/
// NOTE (JIRA: TOP-9501):
// shard_details/catalog_details are mutable post-create only via PATCH in the service,
// not via standard PUT Update APIs. The code generator currently models these fields
// as regular schema blocks (and/or ForceNew/Update semantics), which causes drift after
// PATCH operations: apply succeeds, but terraform plan shows diffs until config is
// manually aligned with the patched state.
// TODO: detect diffs on shard_details/catalog_details after create and translate them
// into PatchDistributedAutonomousDatabase instructions (not PUT / not ForceNew).
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) Patch(ctx context.Context) error {
	request := oci_distributed_database.PatchDistributedAutonomousDatabaseRequest{}

	id := s.D.Id()
	request.DistributedAutonomousDatabaseId = &id

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

	response, err := s.Client.PatchDistributedAutonomousDatabase(ctx, request) // use ctx, not Background()
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDistributedAutonomousDatabaseFromWorkRequest(
		ctx,
		workId,
		tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"),
		oci_distributed_database.ActionTypeUpdated,
		s.D.Timeout(schema.TimeoutUpdate),
	)
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) getDistributedAutonomousDatabaseFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_distributed_database.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	distributedAutonomousDatabaseId, err := distributedAutonomousDatabaseWaitForWorkRequest(ctx, workId, "distributedautonomousdatabase",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*distributedAutonomousDatabaseId)

	return s.GetWithContext(ctx)
}

func distributedAutonomousDatabaseWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func distributedAutonomousDatabaseWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_distributed_database.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_distributed_database.DistributedDbWorkRequestServiceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "distributed_database")
	retryPolicy.ShouldRetryOperation = distributedAutonomousDatabaseWorkRequestShouldRetryFunc(timeout)

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
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.GetDistributedAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &tmp
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
	/*
		if metadata, ok := s.D.GetOkExists("metadata"); ok {
			if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", 0)
				tmp, err := s.mapToDistributedAutonomousDbMetadata(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.Metadata = &tmp
			}
		}*/

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.GetDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DistributedAutonomousDatabase
	return nil
}

/*
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.UpdateDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DistributedAutonomousDatabase
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The code generator omits the required context.Context argument when calling
	// getDistributedAutonomousDatabaseFromWorkRequest, resulting in a compile-time error.
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
	// See  Jira: TOP-9460
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

// NOTE (JIRA: TOP-9493):
// UpdateWithContext must be diff-aware and carefully ordered.
//
// The Distributed Autonomous Database supports multiple update mechanisms:
//   - PUT (UpdateDistributedAutonomousDatabase) for mutable fields
//   - PATCH (patch_operations) for targeted updates
//   - ACTION (ChangeDistributedAutonomousDatabaseCompartment)
//
// The compartment change API is an action-style operation that places the
// resource in UPDATING state. While in this state, the service rejects
// subsequent PUT/PATCH calls with OSD-10184.
//
// To avoid invalid update sequences, this method:
//  1. Executes PUT only when PUT-able fields have changed
//  2. Executes PATCH only when patch_operations changed and are non-empty
//  3. Executes ChangeCompartment LAST, after all PUT/PATCH operations
//  4. Skips unnecessary PUT/PATCH calls entirely when not required
//
// This ensures correct lifecycle transitions, avoids spurious failures,
// and aligns provider behavior with OCI API constraints.
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) UpdateWithContext(ctx context.Context) error {
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

		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}

		id := s.D.Id()
		request.DistributedAutonomousDatabaseId = &id
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

		response, err := s.Client.UpdateDistributedAutonomousDatabase(ctx, request)
		if err != nil {
			return err
		}

		// Keep the model for downstream waits/reads
		s.Res = &response.DistributedAutonomousDatabase
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
	catalogDetails := []interface{}{}
	for _, item := range s.Res.CatalogDetails {
		catalogDetails = append(catalogDetails, DistributedAutonomousDatabaseCatalogToMap(item))
	}
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

	s.D.Set("db_workload", s.Res.DbWorkload)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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

	/*if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{DistributedAutonomousDbMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}*/

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{DistributedAutonomousDbMetadataToMap(s.Res.Metadata)})
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

	shardDetails := []interface{}{}
	for _, item := range s.Res.ShardDetails {
		shardDetails = append(shardDetails, DistributedAutonomousDatabaseShardToMap(item))
	}
	s.D.Set("shard_details", shardDetails)

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

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) StartDistributedAutonomousDatabase(ctx context.Context) error {
	request := oci_distributed_database.StartDistributedAutonomousDatabaseRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	_, err := s.Client.StartDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return s.Res != nil &&
			(s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive ||
				s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateNeedsAttention ||
				s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateInactive)
	}

	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) StopDistributedAutonomousDatabase(ctx context.Context) error {
	request := oci_distributed_database.StopDistributedAutonomousDatabaseRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	_, err := s.Client.StopDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return s.Res != nil &&
			(s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateInactive ||
				s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateNeedsAttention ||
				s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive)
	}
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

/*
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ChangeDistributedAutonomousDbBackupConfig(ctx context.Context) error {
	request := oci_distributed_database.ChangeDistributedAutonomousDbBackupConfigRequest{}

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

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	// WORKAROUND FOR GENERATED CODE ISSUE:
	// This is an action-style operation. The OCI Go SDK response does not include a
	// DistributedAutonomousDatabase model, so we cannot populate s.Res from response.
	// Refresh via GET and wait for a stable lifecycle state instead.
	// See JIRA: TOP-9400

	/*response, err := s.Client.ChangeDistributedAutonomousDbBackupConfig(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("change_db_backup_config_trigger")
	s.D.Set("change_db_backup_config_trigger", val)

	s.Res = &response.DistributedAutonomousDatabase
	return nil*/

/*retentionPolicyFunc := func() bool {
	if err := s.Get(); err != nil { // Refresh status
		log.Printf("[WARN] Failed to refresh resource during wait: %v", err)
		return false
	}
	return s.Res != nil &&
		(s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive ||
			s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateNeedsAttention)
}

_, err := s.Client.ChangeDistributedAutonomousDbBackupConfig(ctx, request)
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

/*if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	// Ensure the resource settles into a stable state (handles eventual consistency).
	if err := tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate)); err != nil {
		return err
	}

	// Preserve trigger semantics in Terraform state.
	val := s.D.Get("change_db_backup_config_trigger")
	s.D.Set("change_db_backup_config_trigger", val)

	return nil

}
*/
// NOTE (JIRA: TOP-9495):
// ChangeDistributedAutonomousDbBackupConfig is asynchronous and returns opc-work-request-id.
// Provider must poll the work request to completion before returning, otherwise Terraform
// may proceed while the resource is still UPDATING / WR IN_PROGRESS.

// NOTE (JIRA: TOP-9YYY):
// After WR completion, refresh the resource and accept terminal lifecycle states:
// ACTIVE, INACTIVE, NEEDS_ATTENTION (final state depends on previous power state).
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ChangeDistributedAutonomousDbBackupConfig(ctx context.Context) error {
	request := oci_distributed_database.ChangeDistributedAutonomousDbBackupConfigRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	// NOTE (JIRA: TOP-9494):
	// Action API payload must come from action-specific schema field, not create-time db_backup_config.
	if v, ok := s.D.GetOkExists("db_backup_config"); ok {
		if tmpList := v.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_backup_config", 0)
			tmp, err := s.mapToDistributedAutonomousDbBackupConfig(fieldKeyFormat)
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

	resp, err := s.Client.ChangeDistributedAutonomousDbBackupConfig(ctx, request)
	if err != nil {
		return err
	}

	// Prefer WR polling because service transitions through UPDATING and may not reflect immediately via GET.
	workId := resp.OpcWorkRequestId
	if workId != nil && *workId != "" {
		if err := s.getDistributedAutonomousDatabaseFromWorkRequest(
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
	case oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive,
		oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateInactive,
		oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateNeedsAttention:
		// ok
	default:
		return fmt.Errorf("unexpected lifecycle state after ChangeDistributedAutonomousDbBackupConfig: %s", s.Res.LifecycleState)
	}

	// Preserve trigger semantics in Terraform state.
	val := s.D.Get("change_db_backup_config_trigger")
	_ = s.D.Set("change_db_backup_config_trigger", val)

	return nil
}

// WORKAROUND FOR GENERATED CODE ISSUE:
//
// tfresource.WaitForUpdatedState expects a tfresource.ResourceUpdater (Update()),
// but this resource CRUD implements UpdateWithContext(ctx) only.
// Add a thin adapter to satisfy the interface.
//
// See JIRA: TOP-9398

/*func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) Update() error {
	return s.UpdateWithContext(context.Background())
}*/
/*
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ConfigureDistributedAutonomousDatabaseSharding(ctx context.Context) error {
	request := oci_distributed_database.ConfigureDistributedAutonomousDatabaseShardingRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	if isRebalanceRequired, ok := s.D.GetOkExists("is_rebalance_required"); ok {
		tmp := isRebalanceRequired.(bool)
		request.IsRebalanceRequired = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// This is an action-style operation. The OCI Go SDK response does not include a
	// DistributedAutonomousDatabase model, so we cannot populate s.Res from response.
	// We must refresh via GET after the operation and wait for a stable lifecycle state.
	// See JIRA: TOP-9400
	/*
		response, err := s.Client.ConfigureDistributedAutonomousDatabaseSharding(context.Background(), request)
		if err != nil {
			return err
		}

		if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
			return waitErr
		}

		val := s.D.Get("configure_sharding_trigger")
		s.D.Set("configure_sharding_trigger", val)

		s.Res = &response.DistributedAutonomousDatabase
		return nil* */

/*_, err := s.Client.ConfigureDistributedAutonomousDatabaseSharding(ctx, request)
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
/*if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	// Preserve trigger semantics in state.
	val := s.D.Get("configure_sharding_trigger")
	s.D.Set("configure_sharding_trigger", val)

	retentionPolicyFunc := func() bool {
		// Refresh status
		if err := s.Get(); err != nil {
			log.Printf("[WARN] Failed to refresh resource during wait: %v", err)
			return false
		}
		return s.Res != nil &&
			(s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive ||
				s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateNeedsAttention)
	}

	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))

}*/

// NOTE (WORKREQUEST + STATE SETTLE):
// ConfigureSharding is asynchronous. We must wait for the Opc-Work-Request-Id to complete,
// then refresh the DADB and wait for a stable lifecycle state.
// For Distributed Autonomous DB, INACTIVE is a valid steady state after create and after some actions,
// so the acceptable post-action lifecycle set is {ACTIVE, INACTIVE, NEEDS_ATTENTION}.
// See JIRA: TOP-9491

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ConfigureDistributedAutonomousDatabaseSharding(ctx context.Context) error {
	request := oci_distributed_database.ConfigureDistributedAutonomousDatabaseShardingRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	// NOTE (CODEGEN GAP):
	// ConfigureSharding supports request param IsRebalanceRequired, but codegen did not
	// expose it in Terraform schema. We surface it as `configure_sharding_is_rebalance_required`.
	// See JIRA: TOP-9490
	if v, ok := s.D.GetOkExists("configure_sharding_is_rebalance_required"); ok {
		tmp := v.(bool)
		request.IsRebalanceRequired = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	// IMPORTANT (ACTION WORK REQUEST):
	// This action returns Opc-Work-Request-Id; lifecycle polling is not reliable.
	// Always wait on the Work Request until it completes.
	// See JIRA: TOP-9491
	resp, err := s.Client.ConfigureDistributedAutonomousDatabaseSharding(ctx, request)
	if err != nil {
		return err
	}

	workId := resp.OpcWorkRequestId
	if workId == nil || *workId == "" {
		return fmt.Errorf("missing opc-work-request-id for ConfigureSharding action")
	}

	// Wait for WR completion (Updated action type is consistent for action APIs)
	if err := s.getDistributedAutonomousDatabaseFromWorkRequest(
		ctx,
		workId,
		tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"),
		oci_distributed_database.ActionTypeUpdated,
		s.D.Timeout(schema.TimeoutUpdate),
	); err != nil {
		return err
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
		return st == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive ||
			st == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateInactive ||
			st == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateNeedsAttention
	}

	if err := tfresource.WaitForResourceCondition(s, stable, s.D.Timeout(schema.TimeoutUpdate)); err != nil {
		return err
	}

	// Preserve trigger semantics in state.
	val := s.D.Get("configure_sharding_trigger")
	_ = s.D.Set("configure_sharding_trigger", val)

	return nil
}

// WORKAROUND FOR GENERATED CODE ISSUE:
//
// Some tfresource helpers (and legacy wait patterns) expect a non-context Get().
// This resource CRUD implements GetWithContext(ctx) only.
// Add a thin adapter for compatibility.
//
// See JIRA: TOP-9399
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) Get() error {
	return s.GetWithContext(context.Background())
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx context.Context) error {
	request := oci_distributed_database.DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequestRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

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
	/*_, err := s.Client.DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx, request)
	if err != nil {
		return err
	}*/

	// NOTE (CODEGEN GAP):
	// The DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest action API
	// returns a PEM-formatted CSR in the response body, but the Terraform code generator
	// does not expose this payload on the resource.
	//
	// To support standard Terraform workflows (e.g. writing CSR to disk),
	// we manually persist the response content into a Computed, Sensitive attribute:
	//
	//   downloaded_gsm_csr_pem
	//
	// This is a temporary workaround until the generator is enhanced to
	// automatically map action API response bodies.
	// See JIRA: TOP-9481

	resp, err := s.Client.DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx, request)
	if err != nil {
		return err
	}
	defer resp.RawResponse.Body.Close()

	b, err := io.ReadAll(resp.RawResponse.Body)
	if err != nil {
		return err
	}

	csr := strings.TrimSpace(string(b))

	// Store in state as sensitive computed output
	if err := s.D.Set("downloaded_gsm_csr_pem", csr); err != nil {
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
	/*
		if waitErr := tfresource.WaitForUpdatedStateWithContext(s.D, s); waitErr != nil {
			return waitErr
		}*/
	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	// Preserve trigger semantics in Terraform state.
	val := s.D.Get("download_gsm_certificate_signing_request_trigger")
	s.D.Set("download_gsm_certificate_signing_request_trigger", val)

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx context.Context) error {
	request := oci_distributed_database.GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequestRequest{}

	// WORKAROUND FOR CODEGEN LIMITATION:
	// The generateGsmCertificateSigningRequest action supports an optional
	// caBundleId query parameter, which is not generated into the Terraform
	// schema by default.
	//
	// This parameter is required to support CA-specific CSR generation.
	// See JIRA: TOP-9478
	if caBundleId, ok := s.D.GetOkExists("generate_gsm_certificate_signing_request_trigger_ca_bundle_id"); ok {
		tmp := caBundleId.(string)
		request.CaBundleId = &tmp
	}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

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

	_, err := s.Client.GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx, request)
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

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	// Preserve trigger semantics in Terraform state.
	val := s.D.Get("generate_gsm_certificate_signing_request_trigger")
	s.D.Set("generate_gsm_certificate_signing_request_trigger", val)

	return nil
}

/*func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) GenerateDistributedAutonomousDatabaseWallet(ctx context.Context) error {
request := oci_distributed_database.GenerateDistributedAutonomousDatabaseWalletRequest{}

idTmp := s.D.Id()
request.DistributedAutonomousDatabaseId = &idTmp*/

// NOTE (JIRA: TOP-9470):
// The code generator exposes generate_wallet_trigger but omits the required
// password parameter needed by the GenerateWallet API.
// A separate schema attribute is added to allow users to provide the password
// when triggering wallet generation.
//if password, ok := s.D.GetOkExists("password"); ok {
/*if password, ok := s.D.GetOkExists("generate_wallet_password"); ok {
	tmp := password.(string)
	request.Password = &tmp
}

request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")*/

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

/*_, err := s.Client.GenerateDistributedAutonomousDatabaseWallet(ctx, request)
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

/*if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	// Preserve trigger semantics in Terraform state.
	val := s.D.Get("generate_wallet_trigger")
	s.D.Set("generate_wallet_trigger", val)

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
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) GenerateDistributedAutonomousDatabaseWallet(ctx context.Context) error {
	request := oci_distributed_database.GenerateDistributedAutonomousDatabaseWalletRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	// Password is required by the action details; exposed via generate_wallet_password (Sensitive).
	if password, ok := s.D.GetOkExists("generate_wallet_password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	resp, err := s.Client.GenerateDistributedAutonomousDatabaseWallet(ctx, request)
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

/*func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet(ctx context.Context) error {
request := oci_distributed_database.UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWalletRequest{}
// NOTE (CODEGEN GAP):
// The generator used/missed a schema field for CaSignedCertificate. We expose it as
// "upload_ca_signed_certificate" to make the action usable from Terraform.
// See JIRA: TOP-9XXX
if caSignedCertificate, ok := s.D.GetOkExists("upload_ca_signed_certificate"); ok {
	tmp := caSignedCertificate.(string)
	request.CaSignedCertificate = &tmp
}

idTmp := s.D.Id()
request.DistributedAutonomousDatabaseId = &idTmp

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
/*
	_, err := s.Client.UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet(ctx, request)
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

/*if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	// Preserve trigger semantics in Terraform state.
	val := s.D.Get("upload_signed_certificate_and_generate_wallet_trigger")
	s.D.Set("upload_signed_certificate_and_generate_wallet_trigger", val)

	return nil
}*/

// NOTE (CODEGEN GAP):
// The UploadSignedCertificateAndGenerateWallet action returns an Opc-Work-Request-Id,
// but the generated Terraform code does NOT wait on the associated Work Request.
// Instead, it incorrectly relies on resource lifecycle polling.
//
// This is incorrect because:
// - The Distributed Autonomous Database lifecycle state may NOT change
//   while the action Work Request is still IN_PROGRESS.
// - Terraform may return success even though the backend operation is still running.
//
// Correct behavior for this action is to explicitly poll the Work Request until it
// reaches a terminal state.
//
// See JIRA: TOP-9483

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet(ctx context.Context) error {
	request := oci_distributed_database.UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWalletRequest{}

	// NOTE (CODEGEN GAP):
	// The generator missed schema wiring for CaSignedCertificate. We expose it as
	// "upload_ca_signed_certificate" to make the action usable from Terraform.
	// See JIRA: TOP-9482
	if caSignedCertificate, ok := s.D.GetOkExists("upload_ca_signed_certificate"); ok {
		tmp := caSignedCertificate.(string)
		request.CaSignedCertificate = &tmp
	}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	// IMPORTANT:
	// This is an async action that returns an OPC work request id.
	// The resource lifecycle may not change while the work request is IN_PROGRESS,
	// so waiting on resource status (WaitForUpdatedState*) can return early.
	// Always wait on the work request.
	// See JIRA: TOP-9483
	resp, err := s.Client.UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet(ctx, request)
	if err != nil {
		return err
	}

	workId := resp.OpcWorkRequestId
	if workId == nil || *workId == "" {
		return fmt.Errorf("missing opc-work-request-id for uploadSignedCertificateAndGenerateWallet action")
	}

	// Wait for WR completion; use ActionTypeUpdated (or the exact enum if service uses a dedicated one).
	if err := s.getDistributedAutonomousDatabaseFromWorkRequest(
		ctx,
		workId,
		tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"),
		oci_distributed_database.ActionTypeUpdated,
		s.D.Timeout(schema.TimeoutUpdate),
	); err != nil {
		return err
	}

	// Preserve trigger semantics in Terraform state.
	val := s.D.Get("upload_signed_certificate_and_generate_wallet_trigger")
	_ = s.D.Set("upload_signed_certificate_and_generate_wallet_trigger", val)

	// Optional: refresh resource after WR completes.
	return s.Get()
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) ValidateDistributedAutonomousDatabaseNetwork(ctx context.Context) error {
	request := oci_distributed_database.ValidateDistributedAutonomousDatabaseNetworkRequest{}

	idTmp := s.D.Id()
	request.DistributedAutonomousDatabaseId = &idTmp

	// WORKAROUND FOR CODEGEN LIMITATION:
	// The ValidateDistributedAutonomousDatabaseNetwork action API supports additional
	// optional request parameters (isSurrogate, resourceName, shardGroup) that are not
	// generated into the Terraform schema by default.
	//
	// These parameters are required for valid network validation scenarios and must be
	// explicitly exposed to users.
	// See JIRA: TOP-9477

	/*if isSurrogate, ok := s.D.GetOkExists("is_surrogate"); ok {
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
	}*/

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
	_, err := s.Client.ValidateDistributedAutonomousDatabaseNetwork(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		if err := s.Get(); err != nil { // Refresh status
			log.Printf("[WARN] Failed to refresh resource during wait: %v", err)
			return false
		}
		return s.Res != nil &&
			(s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive ||
				s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateInactive ||
				s.Res.LifecycleState == oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateNeedsAttention)
	}

	if err := tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate)); err != nil {
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
	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	// Preserve trigger semantics in Terraform state.
	val := s.D.Get("validate_network_trigger")
	s.D.Set("validate_network_trigger", val)

	return nil

}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateCatalogPeerWithDedicatedInfraDetails(fieldKeyFormat string) (oci_distributed_database.CreateCatalogPeerWithDedicatedInfraDetails, error) {
	result := oci_distributed_database.CreateCatalogPeerWithDedicatedInfraDetails{}

	if cloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_autonomous_vm_cluster_id")); ok {
		tmp := cloudAutonomousVmClusterId.(string)
		result.CloudAutonomousVmClusterId = &tmp
	}

	if fastStartFailOverLagLimitInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fast_start_fail_over_lag_limit_in_seconds")); ok {
		tmp := fastStartFailOverLagLimitInSeconds.(int)
		result.FastStartFailOverLagLimitInSeconds = &tmp
	}

	if isAutomaticFailoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_automatic_failover_enabled")); ok {
		tmp := isAutomaticFailoverEnabled.(bool)
		result.IsAutomaticFailoverEnabled = &tmp
	}

	if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
		result.ProtectionMode = oci_distributed_database.DistributedAutonomousDbProtectionModeEnum(protectionMode.(string))
	}

	if standbyMaintenanceBufferInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "standby_maintenance_buffer_in_days")); ok {
		tmp := standbyMaintenanceBufferInDays.(int)
		result.StandbyMaintenanceBufferInDays = &tmp
	}

	return result, nil
}

func CreateCatalogPeerWithDedicatedInfraDetailsToMap(obj oci_distributed_database.CreateCatalogPeerWithDedicatedInfraDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudAutonomousVmClusterId != nil {
		result["cloud_autonomous_vm_cluster_id"] = string(*obj.CloudAutonomousVmClusterId)
	}

	if obj.FastStartFailOverLagLimitInSeconds != nil {
		result["fast_start_fail_over_lag_limit_in_seconds"] = int(*obj.FastStartFailOverLagLimitInSeconds)
	}

	if obj.IsAutomaticFailoverEnabled != nil {
		result["is_automatic_failover_enabled"] = bool(*obj.IsAutomaticFailoverEnabled)
	}

	result["protection_mode"] = string(obj.ProtectionMode)

	if obj.StandbyMaintenanceBufferInDays != nil {
		result["standby_maintenance_buffer_in_days"] = int(*obj.StandbyMaintenanceBufferInDays)
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateDistributedAutonomousDatabaseCatalogDetails(fieldKeyFormat string) (oci_distributed_database.CreateDistributedAutonomousDatabaseCatalogDetails, error) {
	var baseObject oci_distributed_database.CreateDistributedAutonomousDatabaseCatalogDetails
	//discriminator
	sourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source"))
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("ADB_D"):
		details := oci_distributed_database.CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails{}
		if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
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
		if peerCloudAutonomousVmClusterIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_cloud_autonomous_vm_cluster_ids")); ok {
			interfaces := peerCloudAutonomousVmClusterIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "peer_cloud_autonomous_vm_cluster_ids")) {
				details.PeerCloudAutonomousVmClusterIds = tmp
			}
		}
		if peerDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_details")); ok {
			interfaces := peerDetails.([]interface{})
			tmp := make([]oci_distributed_database.CreateCatalogPeerWithDedicatedInfraDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "peer_details"), stateDataIndex)
				converted, err := s.mapToCreateCatalogPeerWithDedicatedInfraDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "peer_details")) {
				details.PeerDetails = tmp
			}
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

func DistributedAutonomousDatabaseCatalogToMap(obj oci_distributed_database.DistributedAutonomousDatabaseCatalog) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// This ToMap helper switches on the OCI Go SDK *model interface* type
	// (e.g. distributeddatabase.DistributedAutonomousDatabaseCatalog).
	// The code generator incorrectly emitted Create*Details discriminator types in the
	// type-switch cases. Those Create* types do not implement the model interface
	// (e.g. missing GetName()), causing an "impossible type switch case" compile error.
	//
	// Fix: use the corresponding SDK model/response concrete type that actually implements
	// the interface (as provided by the vendored oci-go-sdk version).
	// See JIRA: TOP-9402
	//case oci_distributed_database.CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails:
	case oci_distributed_database.DistributedAutonomousDatabaseCatalogWithDedicatedInfra:
		result["source"] = "ADB_D"
		// WORKAROUND FOR GENERATED CODE ISSUE / API DESIGN:
		//
		// AdminPassword is an input-only field available only on
		// CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails.
		// It is not returned by the service and is intentionally absent from
		// the OCI Go SDK response model
		// (DistributedAutonomousDatabaseCatalogWithDedicatedInfra).
		//
		// The generator incorrectly attempted to read AdminPassword from the
		// response model, which is not possible and causes a compile-time error.
		// Do not attempt to populate this field during Read/SetData.
		// See JIRA: TOP-9403

		/*if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}*/

		if v.CloudAutonomousVmClusterId != nil {
			result["cloud_autonomous_vm_cluster_id"] = string(*v.CloudAutonomousVmClusterId)
		}

		if v.ComputeCount != nil {
			result["compute_count"] = float32(*v.ComputeCount)
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

		result["peer_cloud_autonomous_vm_cluster_ids"] = v.PeerCloudAutonomousVmClusterIds

		peerDetails := []interface{}{}
		for _, item := range v.PeerDetails {

			// WORKAROUND FOR GENERATED CODE ISSUE:
			// Read/ToMap must operate on SDK model types, not Create*Details types.
			// The generator incorrectly reused CreateCatalogPeerWithDedicatedInfraDetailsToMap,
			// which expects a request-side struct.
			// See JIRA: TOP-9405
			//peerDetails = append(peerDetails, CreateCatalogPeerWithDedicatedInfraDetailsToMap(item))

			peerDetails = append(peerDetails, CatalogPeerWithDedicatedInfraToMap(item))
		}

		result["peer_details"] = peerDetails

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
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

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateDistributedAutonomousDatabaseShardDetails(fieldKeyFormat string) (oci_distributed_database.CreateDistributedAutonomousDatabaseShardDetails, error) {
	var baseObject oci_distributed_database.CreateDistributedAutonomousDatabaseShardDetails
	//discriminator
	sourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source"))
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("ADB_D"):
		details := oci_distributed_database.CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails{}
		if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
			tmp := adminPassword.(string)
			details.AdminPassword = &tmp
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
		if peerCloudAutonomousVmClusterIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_cloud_autonomous_vm_cluster_ids")); ok {
			interfaces := peerCloudAutonomousVmClusterIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "peer_cloud_autonomous_vm_cluster_ids")) {
				details.PeerCloudAutonomousVmClusterIds = tmp
			}
		}
		if peerDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_details")); ok {
			interfaces := peerDetails.([]interface{})
			tmp := make([]oci_distributed_database.CreateShardPeerWithDedicatedInfraDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "peer_details"), stateDataIndex)
				converted, err := s.mapToCreateShardPeerWithDedicatedInfraDetails(fieldKeyFormatNextLevel)
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
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source '%v' was specified", source)
	}
	return baseObject, nil
}

// WORKAROUND FOR GENERATED CODE ISSUE:
// Read/ToMap must operate on SDK model types, not Create*Details types.
// The generator incorrectly reused CreateCatalogPeerWithDedicatedInfraDetailsToMap,
// which expects a request-side struct.
// See JIRA: TOP-9405

func CatalogPeerWithDedicatedInfraToMap(obj oci_distributed_database.CatalogPeerWithDedicatedInfra) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudAutonomousVmClusterId != nil {
		result["cloud_autonomous_vm_cluster_id"] = *obj.CloudAutonomousVmClusterId
	}

	if obj.FastStartFailOverLagLimitInSeconds != nil {
		result["fast_start_fail_over_lag_limit_in_seconds"] = int(*obj.FastStartFailOverLagLimitInSeconds)
	}

	if obj.IsAutomaticFailoverEnabled != nil {
		result["is_automatic_failover_enabled"] = *obj.IsAutomaticFailoverEnabled
	}

	if obj.ProtectionMode != "" {
		result["protection_mode"] = string(obj.ProtectionMode)
	}

	if obj.StandbyMaintenanceBufferInDays != nil {
		result["standby_maintenance_buffer_in_days"] = int(*obj.StandbyMaintenanceBufferInDays)
	}

	return result
}

// TOP-9407
/*
func DistributedAutonomousDatabaseShardToMap(obj oci_distributed_database.DistributedAutonomousDatabaseShard) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_distributed_database.CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails:
		result["source"] = "ADB_D"

		if v.AdminPassword != nil {
			result["admin_password"] = string(*v.AdminPassword)
		}

		if v.CloudAutonomousVmClusterId != nil {
			result["cloud_autonomous_vm_cluster_id"] = string(*v.CloudAutonomousVmClusterId)
		}

		if v.ComputeCount != nil {
			result["compute_count"] = float32(*v.ComputeCount)
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

		result["peer_cloud_autonomous_vm_cluster_ids"] = v.PeerCloudAutonomousVmClusterIds

		peerDetails := []interface{}{}
		for _, item := range v.PeerDetails {
			peerDetails = append(peerDetails, CreateShardPeerWithDedicatedInfraDetailsToMap(item))
		}
		result["peer_details"] = peerDetails

		if v.ShardSpace != nil {
			result["shard_space"] = string(*v.ShardSpace)
		}

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
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
*/

func DistributedAutonomousDatabaseShardToMap(obj oci_distributed_database.DistributedAutonomousDatabaseShard) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The generator incorrectly emits a type-switch case for the request-side
	// CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails, but the
	// interface here is the response-side DistributedAutonomousDatabaseShard.
	// The Create*Details type can never be a dynamic type of this interface
	// (it doesn't implement GetName), causing an "impossible type switch case"
	// and follow-on missing field errors (Name/TimeCreated/TimeUpdated).
	//
	// Use the SDK model type returned by GET/List instead.
	// See JIRA: TOP-9407
	//case oci_distributed_database.CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails:
	case oci_distributed_database.DistributedAutonomousDatabaseShardWithDedicatedInfra:
		// discriminator used by the SDK for this model
		result["source"] = "ADB_D"

		if v.Name != nil {
			result["name"] = *v.Name
		}
		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}
		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.ComputeCount != nil {
			result["compute_count"] = float64(*v.ComputeCount)
		}
		if v.DataStorageSizeInGbs != nil {
			result["data_storage_size_in_gbs"] = *v.DataStorageSizeInGbs
		}
		if v.IsAutoScalingEnabled != nil {
			result["is_auto_scaling_enabled"] = *v.IsAutoScalingEnabled
		}
		if v.ShardGroup != nil {
			result["shard_group"] = *v.ShardGroup
		}
		if v.CloudAutonomousVmClusterId != nil {
			result["cloud_autonomous_vm_cluster_id"] = *v.CloudAutonomousVmClusterId
		}
		if v.ShardSpace != nil {
			result["shard_space"] = *v.ShardSpace
		}
		if v.VaultId != nil {
			result["vault_id"] = *v.VaultId
		}
		if v.KmsKeyId != nil {
			result["kms_key_id"] = *v.KmsKeyId
		}
		if v.KmsKeyVersionId != nil {
			result["kms_key_version_id"] = *v.KmsKeyVersionId
		}
		if v.SupportingResourceId != nil {
			result["supporting_resource_id"] = *v.SupportingResourceId
		}
		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = *v.ContainerDatabaseId
		}

		// TOP-9408
		peerDetails := []interface{}{}
		for _, item := range v.PeerDetails {
			peerDetails = append(peerDetails, ShardPeerWithDedicatedInfraToMap(item))
		}
		result["peer_details"] = peerDetails

		result["status"] = string(v.Status)

	default:
		log.Printf("[WARN] Unsupported shard type for ToMap: %T", obj)
		return nil
	}

	return result
}

// WORKAROUND FOR GENERATED CODE ISSUE:
// Read/ToMap must operate on SDK *model* types, not Create*Details (request) types.
// The generator incorrectly reused CreateShardPeerWithDedicatedInfraDetailsToMap,
// which expects a request-side struct, but v.PeerDetails contains the response-side
// ShardPeerWithDedicatedInfra model.
// See JIRA: TOP-9407
func ShardPeerWithDedicatedInfraToMap(obj oci_distributed_database.ShardPeerWithDedicatedInfra) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudAutonomousVmClusterId != nil {
		result["cloud_autonomous_vm_cluster_id"] = *obj.CloudAutonomousVmClusterId
	}
	if obj.FastStartFailOverLagLimitInSeconds != nil {
		result["fast_start_fail_over_lag_limit_in_seconds"] = int(*obj.FastStartFailOverLagLimitInSeconds)
	}
	if obj.IsAutomaticFailoverEnabled != nil {
		result["is_automatic_failover_enabled"] = *obj.IsAutomaticFailoverEnabled
	}
	if obj.ProtectionMode != "" {
		result["protection_mode"] = string(obj.ProtectionMode)
	}
	if obj.StandbyMaintenanceBufferInDays != nil {
		result["standby_maintenance_buffer_in_days"] = int(*obj.StandbyMaintenanceBufferInDays)
	}

	return result
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToCreateShardPeerWithDedicatedInfraDetails(fieldKeyFormat string) (oci_distributed_database.CreateShardPeerWithDedicatedInfraDetails, error) {
	result := oci_distributed_database.CreateShardPeerWithDedicatedInfraDetails{}

	if cloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_autonomous_vm_cluster_id")); ok {
		tmp := cloudAutonomousVmClusterId.(string)
		result.CloudAutonomousVmClusterId = &tmp
	}

	if fastStartFailOverLagLimitInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fast_start_fail_over_lag_limit_in_seconds")); ok {
		tmp := fastStartFailOverLagLimitInSeconds.(int)
		result.FastStartFailOverLagLimitInSeconds = &tmp
	}

	if isAutomaticFailoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_automatic_failover_enabled")); ok {
		tmp := isAutomaticFailoverEnabled.(bool)
		result.IsAutomaticFailoverEnabled = &tmp
	}

	if protectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protection_mode")); ok {
		result.ProtectionMode = oci_distributed_database.DistributedAutonomousDbProtectionModeEnum(protectionMode.(string))
	}

	if standbyMaintenanceBufferInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "standby_maintenance_buffer_in_days")); ok {
		tmp := standbyMaintenanceBufferInDays.(int)
		result.StandbyMaintenanceBufferInDays = &tmp
	}

	return result, nil
}

func CreateShardPeerWithDedicatedInfraDetailsToMap(obj oci_distributed_database.CreateShardPeerWithDedicatedInfraDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudAutonomousVmClusterId != nil {
		result["cloud_autonomous_vm_cluster_id"] = string(*obj.CloudAutonomousVmClusterId)
	}

	if obj.FastStartFailOverLagLimitInSeconds != nil {
		result["fast_start_fail_over_lag_limit_in_seconds"] = int(*obj.FastStartFailOverLagLimitInSeconds)
	}

	if obj.IsAutomaticFailoverEnabled != nil {
		result["is_automatic_failover_enabled"] = bool(*obj.IsAutomaticFailoverEnabled)
	}

	result["protection_mode"] = string(obj.ProtectionMode)

	if obj.StandbyMaintenanceBufferInDays != nil {
		result["standby_maintenance_buffer_in_days"] = int(*obj.StandbyMaintenanceBufferInDays)
	}

	return result
}

func DistributedAutonomousDatabaseConnectionStringToMap(obj *oci_distributed_database.DistributedAutonomousDatabaseConnectionString) map[string]interface{} {
	result := map[string]interface{}{}

	result["all_connection_strings"] = obj.AllConnectionStrings

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

	if obj.VersionNumber != nil {
		result["version_number"] = int(*obj.VersionNumber)
	}

	return result
}

func DistributedAutonomousDatabaseSummaryToMap(obj oci_distributed_database.DistributedAutonomousDatabaseSummary) map[string]interface{} {
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
		result["connection_strings"] = []interface{}{DistributedAutonomousDatabaseConnectionStringToMap(obj.ConnectionStrings)}
	}

	if obj.DatabaseVersion != nil {
		result["database_version"] = string(*obj.DatabaseVersion)
	}

	result["db_deployment_type"] = string(obj.DbDeploymentType)

	result["db_workload"] = string(obj.DbWorkload)

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

	if vpcPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpc_password")); ok {
		tmp := vpcPassword.(string)
		result.VpcPassword = &tmp
	}

	if vpcUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpc_user")); ok {
		tmp := vpcUser.(string)
		result.VpcUser = &tmp
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

	if obj.VpcPassword != nil {
		result["vpc_password"] = string(*obj.VpcPassword)
	}

	if obj.VpcUser != nil {
		result["vpc_user"] = string(*obj.VpcUser)
	}

	return result
}

/*func DistributedAutonomousDbMetadataToMap(obj *oci_distributed_database.DistributedAutonomousDbMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	// WORKAROUND FOR GENERATED CODE ISSUE:
	//
	// The generator emits calls to mapToDistributedAutonomousDbMetadata but either
	// doesn't generate the function and/or assumes the SDK model has a `Map` field.
	// In OCI Go SDK, DistributedAutonomousDbMetadata uses `PropertiesMap` (json:"map").
	//
	// This mapper converts the Terraform schema field `metadata.0.map` into
	// `DistributedAutonomousDbMetadata.PropertiesMap`.
	//
	// Remove once generator is fixed.
	// See JIRA: TOP-9397
	//result["map"] = obj.Map
	//result["map"] = obj.PropertiesMap

	// If the input is nil or PropertiesMap is nil, return an empty map
	if obj == nil || obj.PropertiesMap == nil {
		result["map"] = map[string]string{}
		return result
	}

	// Copy values directly (no need to dereference)
	props := make(map[string]string)
	for k, v := range obj.PropertiesMap {
		props[k] = v
	}
	result["map"] = props

	return result
}*/

func DistributedAutonomousDbMetadataToMap(
	obj *oci_distributed_database.DistributedAutonomousDbMetadata,
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

/*func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToPatchInstruction(fieldKeyFormat string) (oci_distributed_database.PatchInstruction, error) {
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
}*/

// mapToPatchInstruction converts Terraform patch_operations blocks into
// OCI PatchInstruction models.
//
// IMPORTANT:
// The PATCH API requires `value` to be a JSON string. Terraform, however,
// may provide the value either as:
//   - a JSON string (via jsonencode in HCL), or
//   - a map/object (depending on how the config is authored)
//
// To ensure API compatibility, the value is normalized before being set
// on the PATCH instruction model.
// See JIRA: TOP-9499
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToPatchInstruction(fieldKeyFormat string) (oci_distributed_database.PatchInstruction, error) {
	var baseObject oci_distributed_database.PatchInstruction

	operationRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation"))
	operation := ""
	if ok {
		operation = operationRaw.(string)
	}

	switch strings.ToLower(operation) {
	case strings.ToLower("INSERT"):
		details := oci_distributed_database.PatchInsertInstruction{}

		if valueRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			normalized, err := normalizePatchValue(valueRaw)
			if err != nil {
				return nil, err
			}
			details.Value = &normalized
		}

		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details

	case strings.ToLower("MERGE"):
		details := oci_distributed_database.PatchMergeInstruction{}

		if valueRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			normalized, err := normalizePatchValue(valueRaw)
			if err != nil {
				return nil, err
			}
			details.Value = &normalized
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
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToobject(fieldKeyFormat string) (oci_distributed_database.Object, error) {
	result := oci_distributed_database.Object{}

	return result, nil
}*/

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

// WORKAROUND FOR GENERATED CODE ISSUE:
//
// The generator emits calls to mapToDistributedAutonomousDbMetadata but either
// doesn't generate the function and/or assumes the SDK model has a `Map` field.
// In OCI Go SDK, DistributedAutonomousDbMetadata uses `PropertiesMap` (json:"map").
//
// This mapper converts the Terraform schema field `metadata.0.map` into
// `DistributedAutonomousDbMetadata.PropertiesMap`.
//
// Remove once generator is fixed.
// See JIRA: TOP-9397
func (s *DistributedDatabaseDistributedAutonomousDatabaseResourceCrud) mapToDistributedAutonomousDbMetadata(
	fieldKeyFormat string,
) (oci_distributed_database.DistributedAutonomousDbMetadata, error) {
	result := oci_distributed_database.DistributedAutonomousDbMetadata{}

	if m, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "map")); ok {
		result.PropertiesMap = tfresource.ObjectMapToStringMap(m.(map[string]interface{}))
	}

	return result, nil
}

// NOTE: Codegen issue
//
// admin_password is a write-only field in the OCI API.
// It is required on create but never returned in GET/LIST responses.
// The generated schema expects round-trip state consistency, which
// results in perpetual Terraform drift and forced recreation.
//
// DiffSuppressFunc is applied here as a workaround to suppress
// false-positive diffs after the resource has been created.
// JIRA: TOP-9459

func suppressMaskedPasswordDiff(k, old, new string, d *schema.ResourceData) bool {
	if d.Id() != "" && old == "" && new != "" {
		log.Printf("[DEBUG]   Suppressing masked admin_password diff for %s", k)
		return true
	}
	return false
}

// normalizePatchValue ensures the PATCH instruction `value` is always
// sent as a JSON string, as required by the Distributed Database PATCH API.
//
// Supported inputs:
//   - string: assumed to already be valid JSON (passed through)
//   - map/object: marshaled to JSON string
//   - any other type: best-effort JSON marshaling
//
// This abstraction prevents subtle bugs where Terraform would otherwise
// send an object instead of a string and cause PATCH requests to fail.
// See JIRA: TOP-9499
func normalizePatchValue(raw interface{}) (interface{}, error) {
	switch v := raw.(type) {
	case string:
		// If user already gave JSON string, pass through.
		return v, nil
	case map[string]interface{}:
		// Convert map -> JSON string (what API examples show)
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		return string(b), nil
	default:
		// If TF ever gives something else, still try to marshal
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		return string(b), nil
	}
}
