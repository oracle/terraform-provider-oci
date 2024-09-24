// // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // Licensed under the Mozilla Public License v2.0
package database_migration

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationMigrationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseMigrationMigration,
		Read:     readDatabaseMigrationMigration,
		Update:   updateDatabaseMigrationMigration,
		Delete:   deleteDatabaseMigrationMigration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_combination": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"MYSQL",
					"ORACLE",
				}, true),
			},
			"source_database_connection_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_database_connection_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"advanced_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"data_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
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
			"advisor_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_ignore_errors": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_skip_advisor": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"bulk_include_exclude_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"data_transfer_medium_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
								"AWS_S3",
								"DBLINK",
								"NFS",
								"OBJECT_STORAGE",
							}, true),
						},

						// Optional
						"access_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"object_storage_bucket": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secret_access_key": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							Sensitive: true,
						},
						"shared_storage_mount_target_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"kind": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"CURL",
											"OCI_CLI",
										}, true),
									},

									// Optional
									"oci_home": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"wallet_location": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"target": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"kind": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"CURL",
											"OCI_CLI",
										}, true),
									},

									// Optional
									"oci_home": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"wallet_location": {
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
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exclude_objects": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Set:      excludeObjectsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"object": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"is_omit_excluded_table_from_replication": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"schema": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"ggs_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"acceptable_lag": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"extract": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"long_trans_duration": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"performance_profile": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"replicat": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"performance_profile": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
						"ggs_deployment": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"deployment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ggs_admin_credentials_secret_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"hub_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"rest_admin_credentials": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"password": {
										Type:      schema.TypeString,
										Required:  true,
										Sensitive: true,
									},
									"username": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
						"vault_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"acceptable_lag": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"compute_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"extract": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"long_trans_duration": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"performance_profile": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"replicat": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"performance_profile": {
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
			"include_objects": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"object": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"is_omit_excluded_table_from_replication": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"schema": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"initial_load_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"job_mode": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"compatibility": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"data_pump_parameters": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"estimate": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"exclude_parameters": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"export_parallelism_degree": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"import_parallelism_degree": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"is_cluster": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"table_exists_action": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"export_directory_object": {
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
									"path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"handle_grant_errors": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"import_directory_object": {
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
									"path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"is_consistent": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_ignore_existing_objects": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_tz_utc": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"metadata_remaps": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"new_value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"old_value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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
						"primary_key_compatibility": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tablespace_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"target_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"ADB_D_AUTOCREATE",
											"ADB_D_REMAP",
											"ADB_S_REMAP",
											"NON_ADB_AUTOCREATE",
											"NON_ADB_REMAP",
											"TARGET_DEFAULTS_AUTOCREATE",
											"TARGET_DEFAULTS_REMAP",
										}, true),
									},

									// Optional
									"block_size_in_kbs": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"extend_size_in_mbs": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"is_auto_create": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"is_big_file": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"remap_target": {
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
			"source_container_database_connection_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"executing_job_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
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
			"time_last_migration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wait_after": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func excludeObjectsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if object, ok := m["object"]; ok && object != "" {
		buf.WriteString(fmt.Sprintf("%v-", object))
	}
	if owner, ok := m["owner"]; ok && owner != "" {
		buf.WriteString(fmt.Sprintf("%v-", owner))
	}

	return utils.GetStringHashcode(buf.String())
}

func createDatabaseMigrationMigration(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationMigrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseMigrationMigration(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationMigrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseMigrationMigration(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationMigrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseMigrationMigration(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationMigrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseMigrationMigrationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_migration.DatabaseMigrationClient
	Res                    *oci_database_migration.Migration
	DisableNotFoundRetries bool
}

func (s *DatabaseMigrationMigrationResourceCrud) ID() string {
	migration := *s.Res
	return *migration.GetId()
}

func (s *DatabaseMigrationMigrationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_migration.MigrationLifecycleStatesCreating),
		string(oci_database_migration.MigrationLifecycleStatesInProgress),
	}
}

func (s *DatabaseMigrationMigrationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_migration.MigrationLifecycleStatesActive),
		string(oci_database_migration.MigrationLifecycleStatesSucceeded),
		string(oci_database_migration.MigrationLifecycleStatesNeedsAttention),
	}
}

func (s *DatabaseMigrationMigrationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_migration.MigrationLifecycleStatesDeleting),
	}
}

func (s *DatabaseMigrationMigrationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_migration.MigrationLifecycleStatesDeleted),
	}
}

func (s *DatabaseMigrationMigrationResourceCrud) Create() error {
	request := oci_database_migration.CreateMigrationRequest{}
	err := s.populateTopLevelPolymorphicCreateMigrationRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.CreateMigration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getMigrationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration"), oci_database_migration.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseMigrationMigrationResourceCrud) getMigrationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_migration.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	migrationId, err := migrationWaitForWorkRequest(workId, "migration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*migrationId)

	return s.Get()
}

func migrationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "database_migration", startTime) {
			return true
		}

		// Only stop if status of work request response is succeeded
		if workRequestResponse, ok := response.Response.(oci_database_migration.GetWorkRequestResponse); ok {
			return workRequestResponse.Status != oci_database_migration.OperationStatusSucceeded
		}

		return false
	}
}

func migrationWaitForWorkRequest(wId *string, entityType string, action oci_database_migration.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_migration.DatabaseMigrationClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_migration")
	retryPolicy.ShouldRetryOperation = migrationWorkRequestShouldRetryFunc(timeout)

	response := oci_database_migration.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_database_migration.OperationStatusInProgress),
			string(oci_database_migration.OperationStatusAccepted),
			string(oci_database_migration.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_database_migration.OperationStatusSucceeded),
			string(oci_database_migration.OperationStatusFailed),
			string(oci_database_migration.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_database_migration.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_database_migration.OperationStatusFailed || response.Status == oci_database_migration.OperationStatusCanceled {
		return nil, getErrorFromDatabaseMigrationMigrationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseMigrationMigrationWorkRequest(client *oci_database_migration.DatabaseMigrationClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_migration.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_database_migration.ListWorkRequestErrorsRequest{
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

func (s *DatabaseMigrationMigrationResourceCrud) Get() error {
	request := oci_database_migration.GetMigrationRequest{}

	tmp := s.D.Id()
	request.MigrationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.GetMigration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Migration
	return nil
}

func (s *DatabaseMigrationMigrationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_migration.UpdateMigrationRequest{}
	err := s.populateTopLevelPolymorphicUpdateMigrationRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.UpdateMigration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMigrationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration"), oci_database_migration.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseMigrationMigrationResourceCrud) Delete() error {
	request := oci_database_migration.DeleteMigrationRequest{}

	tmp := s.D.Id()
	request.MigrationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.DeleteMigration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := migrationWaitForWorkRequest(workId, "migration",
		oci_database_migration.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatabaseMigrationMigrationResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_database_migration.MySqlMigration:
		s.D.Set("database_combination", "MYSQL")

		if v.AdvisorSettings != nil {
			s.D.Set("advisor_settings", []interface{}{MySqlAdvisorSettingsToMap(v.AdvisorSettings)})
		} else {
			s.D.Set("advisor_settings", nil)
		}

		if v.DataTransferMediumDetails != nil {
			dataTransferMediumDetailsArray := []interface{}{}
			if dataTransferMediumDetailsMap := MySqlDataTransferMediumDetailsToMap(&v.DataTransferMediumDetails); dataTransferMediumDetailsMap != nil {
				dataTransferMediumDetailsArray = append(dataTransferMediumDetailsArray, dataTransferMediumDetailsMap)
			}
			s.D.Set("data_transfer_medium_details", dataTransferMediumDetailsArray)
		} else {
			s.D.Set("data_transfer_medium_details", nil)
		}

		if v.GgsDetails != nil {
			s.D.Set("ggs_details", []interface{}{MySqlGgsDeploymentDetailsToMap(v.GgsDetails)})
		} else {
			s.D.Set("ggs_details", nil)
		}

		if v.HubDetails != nil {
			s.D.Set("hub_details", []interface{}{GoldenGateHubDetailsToMap(v.HubDetails)})
		} else {
			s.D.Set("hub_details", nil)
		}

		if v.InitialLoadSettings != nil {
			s.D.Set("initial_load_settings", []interface{}{MySqlInitialLoadSettingsToMap(v.InitialLoadSettings)})
		} else {
			s.D.Set("initial_load_settings", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.ExecutingJobId != nil {
			s.D.Set("executing_job_id", *v.ExecutingJobId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		s.D.Set("lifecycle_details", v.LifecycleDetails)

		if v.SourceDatabaseConnectionId != nil {
			s.D.Set("source_database_connection_id", *v.SourceDatabaseConnectionId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetDatabaseConnectionId != nil {
			s.D.Set("target_database_connection_id", *v.TargetDatabaseConnectionId)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeLastMigration != nil {
			s.D.Set("time_last_migration", v.TimeLastMigration.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		s.D.Set("type", v.Type)

		s.D.Set("wait_after", v.WaitAfter)
	case oci_database_migration.OracleMigration:
		s.D.Set("database_combination", "ORACLE")

		if v.AdvancedParameters != nil {
			advancedParameters := []interface{}{}
			for _, item := range v.AdvancedParameters {
				advancedParameters = append(advancedParameters, migrationParameterDetailsToMap(item))
			}
			s.D.Set("advanced_parameters", advancedParameters)
		}

		if v.AdvisorSettings != nil {
			s.D.Set("advisor_settings", []interface{}{OracleAdvisorSettingsToMap(v.AdvisorSettings)})
		} else {
			s.D.Set("advisor_settings", nil)
		}

		if v.DataTransferMediumDetails != nil {
			dataTransferMediumDetailsArray := []interface{}{}
			if dataTransferMediumDetailsMap := OracleDataTransferMediumDetailsToMap(&v.DataTransferMediumDetails); dataTransferMediumDetailsMap != nil {
				dataTransferMediumDetailsArray = append(dataTransferMediumDetailsArray, dataTransferMediumDetailsMap)
			}
			s.D.Set("data_transfer_medium_details", dataTransferMediumDetailsArray)
		} else {
			s.D.Set("data_transfer_medium_details", nil)
		}

		if v.GgsDetails != nil {
			s.D.Set("ggs_details", []interface{}{OracleGgsDeploymentDetailsToMap(v.GgsDetails)})
		} else {
			s.D.Set("ggs_details", nil)
		}

		if v.HubDetails != nil {
			s.D.Set("hub_details", []interface{}{GoldenGateHubDetailsToMap(v.HubDetails)})
		} else {
			s.D.Set("hub_details", nil)
		}

		if v.InitialLoadSettings != nil {
			s.D.Set("initial_load_settings", []interface{}{OracleInitialLoadSettingsToMap(v.InitialLoadSettings)})
		} else {
			s.D.Set("initial_load_settings", nil)
		}

		if v.SourceContainerDatabaseConnectionId != nil {
			s.D.Set("source_container_database_connection_id", *v.SourceContainerDatabaseConnectionId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.ExecutingJobId != nil {
			s.D.Set("executing_job_id", *v.ExecutingJobId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		s.D.Set("lifecycle_details", v.LifecycleDetails)

		if v.SourceDatabaseConnectionId != nil {
			s.D.Set("source_database_connection_id", *v.SourceDatabaseConnectionId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetDatabaseConnectionId != nil {
			s.D.Set("target_database_connection_id", *v.TargetDatabaseConnectionId)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeLastMigration != nil {
			s.D.Set("time_last_migration", v.TimeLastMigration.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		s.D.Set("type", v.Type)

		s.D.Set("wait_after", v.WaitAfter)
	default:
		log.Printf("[WARN] Received 'database_combination' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func migrationParameterDetailsToMap(obj oci_database_migration.MigrationParameterDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	result["data_type"] = obj.DataType

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToAdminCredentials(fieldKeyFormat string) (oci_database_migration.AdminCredentials, error) {
	result := oci_database_migration.AdminCredentials{}

	if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
		tmp := username.(string)
		result.Username = &tmp
	}

	return result, nil
}

func AdminCredentialsToMap(obj *oci_database_migration.AdminCredentials) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateAdminCredentials(fieldKeyFormat string) (oci_database_migration.CreateAdminCredentials, error) {
	result := oci_database_migration.CreateAdminCredentials{}

	if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
		tmp := password.(string)
		result.Password = &tmp
	}

	if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
		tmp := username.(string)
		result.Username = &tmp
	}

	return result, nil
}

func CreateAdminCredentialsToMap(obj *oci_database_migration.CreateAdminCredentials) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Password != nil {
		result["password"] = string(*obj.Password)
	}

	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateDataPumpParameters(fieldKeyFormat string) (oci_database_migration.CreateDataPumpParameters, error) {
	result := oci_database_migration.CreateDataPumpParameters{}

	if estimate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "estimate")); ok {
		result.Estimate = oci_database_migration.DataPumpEstimateEnum(estimate.(string))
	}

	if excludeParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_parameters")); ok {
		interfaces := excludeParameters.([]interface{})
		tmp := make([]oci_database_migration.DataPumpExcludeParametersEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_database_migration.DataPumpExcludeParametersEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclude_parameters")) {
			result.ExcludeParameters = tmp
		}
	}

	exportParallelismDegree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_parallelism_degree"))
	if ok && exportParallelismDegree.(int) != 0 {
		tmp := exportParallelismDegree.(int)
		result.ExportParallelismDegree = &tmp
	} else {
		result.ExportParallelismDegree = nil
	}

	importParallelismDegree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "import_parallelism_degree"))
	if ok && importParallelismDegree.(int) != 0 {
		tmp := importParallelismDegree.(int)
		result.ImportParallelismDegree = &tmp
	} else {
		result.ImportParallelismDegree = nil
	}

	if isCluster, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_cluster")); ok {
		tmp := isCluster.(bool)
		result.IsCluster = &tmp
	}

	if tableExistsAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "table_exists_action")); ok {
		result.TableExistsAction = oci_database_migration.DataPumpTableExistsActionEnum(tableExistsAction.(string))
	}

	return result, nil
}

func CreateDataPumpParametersToMap(obj *oci_database_migration.CreateDataPumpParameters) map[string]interface{} {
	result := map[string]interface{}{}

	result["estimate"] = string(obj.Estimate)

	result["exclude_parameters"] = obj.ExcludeParameters

	if obj.ExportParallelismDegree != nil && int(*obj.ExportParallelismDegree) != 0 {
		result["export_parallelism_degree"] = int(*obj.ExportParallelismDegree)
	} else {
		result["export_parallelism_degree"] = nil
	}

	if obj.ImportParallelismDegree != nil && int(*obj.ImportParallelismDegree) != 0 {
		result["import_parallelism_degree"] = int(*obj.ImportParallelismDegree)
	} else {
		result["import_parallelism_degree"] = nil
	}

	if obj.IsCluster != nil {
		result["is_cluster"] = bool(*obj.IsCluster)
	}

	result["table_exists_action"] = string(obj.TableExistsAction)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateDirectoryObject(fieldKeyFormat string) (oci_database_migration.CreateDirectoryObject, error) {
	result := oci_database_migration.CreateDirectoryObject{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
		tmp := path.(string)
		result.Path = &tmp
	}

	return result, nil
}

func CreateDirectoryObjectToMap(obj *oci_database_migration.CreateDirectoryObject) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateExtract(fieldKeyFormat string) (oci_database_migration.CreateExtract, error) {
	result := oci_database_migration.CreateExtract{}

	longTransDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "long_trans_duration"))
	if ok && longTransDuration.(int) != 0 {
		tmp := longTransDuration.(int)
		result.LongTransDuration = &tmp
	} else {
		result.LongTransDuration = nil
	}

	if performanceProfile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "performance_profile")); ok {
		result.PerformanceProfile = oci_database_migration.ExtractPerformanceProfileEnum(performanceProfile.(string))
	}

	return result, nil
}

func CreateExtractToMap(obj *oci_database_migration.CreateExtract) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LongTransDuration != nil && int(*obj.LongTransDuration) != 0 {
		result["long_trans_duration"] = int(*obj.LongTransDuration)
	} else {
		result["long_trans_duration"] = nil
	}

	result["performance_profile"] = string(obj.PerformanceProfile)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateGoldenGateHubDetails(fieldKeyFormat string) (oci_database_migration.CreateGoldenGateHubDetails, error) {
	result := oci_database_migration.CreateGoldenGateHubDetails{}

	acceptableLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "acceptable_lag"))
	if ok && acceptableLag.(int) != 0 {
		tmp := acceptableLag.(int)
		result.AcceptableLag = &tmp
	} else {
		result.AcceptableLag = nil
	}

	if computeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_id")); ok {
		tmp := computeId.(string)
		result.ComputeId = &tmp
	}

	if extract, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extract")); ok {
		if tmpList := extract.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "extract"), 0)
			tmp, err := s.mapToCreateExtract(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert extract, encountered error: %v", err)
			}
			result.Extract = &tmp
		}
	}

	if keyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_id")); ok {
		tmp := keyId.(string)
		result.KeyId = &tmp
	}

	if replicat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replicat")); ok {
		if tmpList := replicat.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "replicat"), 0)
			tmp, err := s.mapToCreateReplicat(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert replicat, encountered error: %v", err)
			}
			result.Replicat = &tmp
		}
	}

	if restAdminCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rest_admin_credentials")); ok {
		if tmpList := restAdminCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rest_admin_credentials"), 0)
			tmp, err := s.mapToCreateAdminCredentials(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert rest_admin_credentials, encountered error: %v", err)
			}
			result.RestAdminCredentials = &tmp
		}
	}

	if url, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url")); ok {
		tmp := url.(string)
		result.Url = &tmp
	}

	if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
		tmp := vaultId.(string)
		result.VaultId = &tmp
	}

	return result, nil
}

func CreateGoldenGateHubDetailsToMap(obj *oci_database_migration.CreateGoldenGateHubDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptableLag != nil {
		result["acceptable_lag"] = int(*obj.AcceptableLag)
	} else {
		result["acceptable_lag"] = nil
	}

	if obj.ComputeId != nil {
		result["compute_id"] = string(*obj.ComputeId)
	}

	if obj.Extract != nil {
		result["extract"] = []interface{}{CreateExtractToMap(obj.Extract)}
	}

	if obj.KeyId != nil {
		result["key_id"] = string(*obj.KeyId)
	}

	if obj.Replicat != nil {
		result["replicat"] = []interface{}{CreateReplicatToMap(obj.Replicat)}
	}

	if obj.RestAdminCredentials != nil {
		result["rest_admin_credentials"] = []interface{}{CreateAdminCredentialsToMap(obj.RestAdminCredentials)}
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	if obj.VaultId != nil {
		result["vault_id"] = string(*obj.VaultId)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateMySqlAdvisorSettings(fieldKeyFormat string) (oci_database_migration.CreateMySqlAdvisorSettings, error) {
	result := oci_database_migration.CreateMySqlAdvisorSettings{}

	if isIgnoreErrors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ignore_errors")); ok {
		tmp := isIgnoreErrors.(bool)
		result.IsIgnoreErrors = &tmp
	}

	if isSkipAdvisor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_advisor")); ok {
		tmp := isSkipAdvisor.(bool)
		result.IsSkipAdvisor = &tmp
	}

	return result, nil
}

func CreateMySqlAdvisorSettingsToMap(obj *oci_database_migration.CreateMySqlAdvisorSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsIgnoreErrors != nil {
		result["is_ignore_errors"] = bool(*obj.IsIgnoreErrors)
	}

	if obj.IsSkipAdvisor != nil {
		result["is_skip_advisor"] = bool(*obj.IsSkipAdvisor)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateMySqlDataTransferMediumDetails(fieldKeyFormat string) (oci_database_migration.CreateMySqlDataTransferMediumDetails, error) {
	var baseObject oci_database_migration.CreateMySqlDataTransferMediumDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_database_migration.CreateMySqlObjectStorageDataTransferMediumDetails{}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToCreateObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func CreateMySqlDataTransferMediumDetailsToMap(obj *oci_database_migration.CreateMySqlDataTransferMediumDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_migration.CreateMySqlObjectStorageDataTransferMediumDetails:
		result["type"] = "OBJECT_STORAGE"

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{CreateObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateMySqlGgsDeploymentDetails(fieldKeyFormat string) (oci_database_migration.CreateMySqlGgsDeploymentDetails, error) {
	result := oci_database_migration.CreateMySqlGgsDeploymentDetails{}

	acceptableLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "acceptable_lag"))
	if ok && acceptableLag.(int) != 0 {
		tmp := acceptableLag.(int)
		result.AcceptableLag = &tmp
	} else {
		result.AcceptableLag = nil
	}

	if replicat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replicat")); ok {
		if tmpList := replicat.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "replicat"), 0)
			tmp, err := s.mapToCreateReplicat(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert replicat, encountered error: %v", err)
			}
			result.Replicat = &tmp
		}
	}

	return result, nil
}

func CreateMySqlGgsDeploymentDetailsToMap(obj *oci_database_migration.CreateMySqlGgsDeploymentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptableLag != nil {
		result["acceptable_lag"] = int(*obj.AcceptableLag)
	} else {
		result["acceptable_lag"] = nil
	}

	if obj.Replicat != nil {
		result["replicat"] = []interface{}{CreateReplicatToMap(obj.Replicat)}
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateMySqlInitialLoadSettings(fieldKeyFormat string) (oci_database_migration.CreateMySqlInitialLoadSettings, error) {
	result := oci_database_migration.CreateMySqlInitialLoadSettings{}

	if compatibility, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compatibility")); ok {
		interfaces := compatibility.([]interface{})
		tmp := make([]oci_database_migration.CompatibilityOptionEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_database_migration.CompatibilityOptionEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "compatibility")) {
			result.Compatibility = tmp
		}
	}

	if handleGrantErrors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "handle_grant_errors")); ok {
		result.HandleGrantErrors = oci_database_migration.HandleGrantErrorsEnum(handleGrantErrors.(string))
	}

	if isConsistent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_consistent")); ok {
		tmp := isConsistent.(bool)
		result.IsConsistent = &tmp
	}

	if isIgnoreExistingObjects, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ignore_existing_objects")); ok {
		tmp := isIgnoreExistingObjects.(bool)
		result.IsIgnoreExistingObjects = &tmp
	}

	if isTzUtc, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_tz_utc")); ok {
		tmp := isTzUtc.(bool)
		result.IsTzUtc = &tmp
	}

	if jobMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_mode")); ok {
		result.JobMode = oci_database_migration.JobModeMySqlEnum(jobMode.(string))
	}

	if primaryKeyCompatibility, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_key_compatibility")); ok {
		result.PrimaryKeyCompatibility = oci_database_migration.PrimaryKeyCompatibilityEnum(primaryKeyCompatibility.(string))
	}

	return result, nil
}

func CreateMySqlInitialLoadSettingsToMap(obj *oci_database_migration.CreateMySqlInitialLoadSettings) map[string]interface{} {
	result := map[string]interface{}{}

	result["compatibility"] = obj.Compatibility

	result["handle_grant_errors"] = string(obj.HandleGrantErrors)

	if obj.IsConsistent != nil {
		result["is_consistent"] = bool(*obj.IsConsistent)
	}

	if obj.IsIgnoreExistingObjects != nil {
		result["is_ignore_existing_objects"] = bool(*obj.IsIgnoreExistingObjects)
	}

	if obj.IsTzUtc != nil {
		result["is_tz_utc"] = bool(*obj.IsTzUtc)
	}

	result["job_mode"] = string(obj.JobMode)

	result["primary_key_compatibility"] = string(obj.PrimaryKeyCompatibility)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateObjectStoreBucket(fieldKeyFormat string) (oci_database_migration.CreateObjectStoreBucket, error) {
	result := oci_database_migration.CreateObjectStoreBucket{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.NamespaceName = &tmp
	}

	return result, nil
}

func CreateObjectStoreBucketToMap(obj *oci_database_migration.CreateObjectStoreBucket) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateOracleAdvisorSettings(fieldKeyFormat string) (oci_database_migration.CreateOracleAdvisorSettings, error) {
	result := oci_database_migration.CreateOracleAdvisorSettings{}

	if isIgnoreErrors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ignore_errors")); ok {
		tmp := isIgnoreErrors.(bool)
		result.IsIgnoreErrors = &tmp
	}

	if isSkipAdvisor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_advisor")); ok {
		tmp := isSkipAdvisor.(bool)
		result.IsSkipAdvisor = &tmp
	}

	return result, nil
}

func CreateOracleAdvisorSettingsToMap(obj *oci_database_migration.CreateOracleAdvisorSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsIgnoreErrors != nil {
		result["is_ignore_errors"] = bool(*obj.IsIgnoreErrors)
	}

	if obj.IsSkipAdvisor != nil {
		result["is_skip_advisor"] = bool(*obj.IsSkipAdvisor)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateOracleDataTransferMediumDetails(fieldKeyFormat string) (oci_database_migration.CreateOracleDataTransferMediumDetails, error) {
	var baseObject oci_database_migration.CreateOracleDataTransferMediumDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("AWS_S3"):
		details := oci_database_migration.CreateOracleAwsS3DataTransferMediumDetails{}
		if accessKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if secretAccessKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_access_key")); ok {
			tmp := secretAccessKey.(string)
			details.SecretAccessKey = &tmp
		}
		baseObject = details
	case strings.ToLower("DBLINK"):
		details := oci_database_migration.CreateOracleDbLinkDataTransferMediumDetails{}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToCreateObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("NFS"):
		details := oci_database_migration.CreateOracleNfsDataTransferMediumDetails{}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToCreateObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		if sharedStorageMountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shared_storage_mount_target_id")); ok {
			tmp := sharedStorageMountTargetId.(string)
			details.SharedStorageMountTargetId = &tmp
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			if tmpList := source.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert source, encountered error: %v", err)
				}
				details.Source = tmp
			}
		}
		if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
			if tmpList := target.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "target"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert target, encountered error: %v", err)
				}
				details.Target = tmp
			}
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_database_migration.CreateOracleObjectStorageDataTransferMediumDetails{}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToCreateObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			if tmpList := source.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert source, encountered error: %v", err)
				}
				details.Source = tmp
			}
		}
		if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
			if tmpList := target.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "target"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert target, encountered error: %v", err)
				}
				details.Target = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func CreateOracleDataTransferMediumDetailsToMap(obj *oci_database_migration.CreateOracleDataTransferMediumDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_migration.CreateOracleAwsS3DataTransferMediumDetails:
		result["type"] = "AWS_S3"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{ObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.SecretAccessKey != nil {
			result["secret_access_key"] = string(*v.SecretAccessKey)
		}
	case oci_database_migration.CreateOracleDbLinkDataTransferMediumDetails:
		result["type"] = "DBLINK"

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{CreateObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}
	case oci_database_migration.CreateOracleNfsDataTransferMediumDetails:
		result["type"] = "NFS"

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{CreateObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}

		if v.SharedStorageMountTargetId != nil {
			result["shared_storage_mount_target_id"] = string(*v.SharedStorageMountTargetId)
		}

		if v.Source != nil {
			sourceArray := []interface{}{}
			if sourceMap := HostDumpTransferDetailsToMap(&v.Source); sourceMap != nil {
				sourceArray = append(sourceArray, sourceMap)
			}
			result["source"] = sourceArray
		}

		if v.Target != nil {
			targetArray := []interface{}{}
			if targetMap := HostDumpTransferDetailsToMap(&v.Target); targetMap != nil {
				targetArray = append(targetArray, targetMap)
			}
			result["target"] = targetArray
		}
	case oci_database_migration.CreateOracleObjectStorageDataTransferMediumDetails:
		result["type"] = "OBJECT_STORAGE"

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{CreateObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}

		if v.Source != nil {
			sourceArray := []interface{}{}
			if sourceMap := HostDumpTransferDetailsToMap(&v.Source); sourceMap != nil {
				sourceArray = append(sourceArray, sourceMap)
			}
			result["source"] = sourceArray
		}

		if v.Target != nil {
			targetArray := []interface{}{}
			if targetMap := HostDumpTransferDetailsToMap(&v.Target); targetMap != nil {
				targetArray = append(targetArray, targetMap)
			}
			result["target"] = targetArray
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateOracleGgsDeploymentDetails(fieldKeyFormat string) (oci_database_migration.CreateOracleGgsDeploymentDetails, error) {
	result := oci_database_migration.CreateOracleGgsDeploymentDetails{}

	acceptableLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "acceptable_lag"))
	if ok && acceptableLag.(int) != 0 {
		tmp := acceptableLag.(int)
		result.AcceptableLag = &tmp
	} else {
		result.AcceptableLag = nil
	}

	if extract, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extract")); ok {
		if tmpList := extract.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "extract"), 0)
			tmp, err := s.mapToCreateExtract(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert extract, encountered error: %v", err)
			}
			result.Extract = &tmp
		}
	}

	if replicat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replicat")); ok {
		if tmpList := replicat.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "replicat"), 0)
			tmp, err := s.mapToCreateReplicat(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert replicat, encountered error: %v", err)
			}
			result.Replicat = &tmp
		}
	}

	return result, nil
}

func CreateOracleGgsDeploymentDetailsToMap(obj *oci_database_migration.CreateOracleGgsDeploymentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptableLag != nil {
		result["acceptable_lag"] = int(*obj.AcceptableLag)
	} else {
		result["acceptable_lag"] = nil
	}

	if obj.Extract != nil {
		result["extract"] = []interface{}{CreateExtractToMap(obj.Extract)}
	}

	if obj.Replicat != nil {
		result["replicat"] = []interface{}{CreateReplicatToMap(obj.Replicat)}
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateOracleInitialLoadSettings(fieldKeyFormat string) (oci_database_migration.CreateOracleInitialLoadSettings, error) {
	result := oci_database_migration.CreateOracleInitialLoadSettings{}

	if dataPumpParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_pump_parameters")); ok {
		if tmpList := dataPumpParameters.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data_pump_parameters"), 0)
			tmp, err := s.mapToCreateDataPumpParameters(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert data_pump_parameters, encountered error: %v", err)
			}
			result.DataPumpParameters = &tmp
		}
	}

	if exportDirectoryObject, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_directory_object")); ok {
		if tmpList := exportDirectoryObject.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "export_directory_object"), 0)
			tmp, err := s.mapToCreateDirectoryObject(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert export_directory_object, encountered error: %v", err)
			}
			result.ExportDirectoryObject = &tmp
		}
	}

	if importDirectoryObject, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "import_directory_object")); ok {
		if tmpList := importDirectoryObject.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "import_directory_object"), 0)
			tmp, err := s.mapToCreateDirectoryObject(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert import_directory_object, encountered error: %v", err)
			}
			result.ImportDirectoryObject = &tmp
		}
	}

	if jobMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_mode")); ok {
		result.JobMode = oci_database_migration.JobModeOracleEnum(jobMode.(string))
	}

	if metadataRemaps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata_remaps")); ok {
		interfaces := metadataRemaps.([]interface{})
		tmp := make([]oci_database_migration.MetadataRemap, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "metadata_remaps"), stateDataIndex)
			converted, err := s.mapToMetadataRemap(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "metadata_remaps")) {
			result.MetadataRemaps = tmp
		}
	}

	if tablespaceDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tablespace_details")); ok {
		if tmpList := tablespaceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tablespace_details"), 0)
			tmp, err := s.mapToCreateTargetTypeTablespaceDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert tablespace_details, encountered error: %v", err)
			}
			result.TablespaceDetails = tmp
		}
	}

	return result, nil
}

func CreateOracleInitialLoadSettingsToMap(obj *oci_database_migration.CreateOracleInitialLoadSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DataPumpParameters != nil {
		result["data_pump_parameters"] = []interface{}{CreateDataPumpParametersToMap(obj.DataPumpParameters)}
	}

	if obj.ExportDirectoryObject != nil {
		result["export_directory_object"] = []interface{}{CreateDirectoryObjectToMap(obj.ExportDirectoryObject)}
	}

	if obj.ImportDirectoryObject != nil {
		result["import_directory_object"] = []interface{}{CreateDirectoryObjectToMap(obj.ImportDirectoryObject)}
	}

	result["job_mode"] = string(obj.JobMode)

	metadataRemaps := []interface{}{}
	for _, item := range obj.MetadataRemaps {
		metadataRemaps = append(metadataRemaps, MetadataRemapToMap(item))
	}
	result["metadata_remaps"] = metadataRemaps

	if obj.TablespaceDetails != nil {
		tablespaceDetailsArray := []interface{}{}
		if tablespaceDetailsMap := CreateTargetTypeTablespaceDetailsToMap(&obj.TablespaceDetails); tablespaceDetailsMap != nil {
			tablespaceDetailsArray = append(tablespaceDetailsArray, tablespaceDetailsMap)
		}
		result["tablespace_details"] = tablespaceDetailsArray
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateReplicat(fieldKeyFormat string) (oci_database_migration.CreateReplicat, error) {
	result := oci_database_migration.CreateReplicat{}

	if performanceProfile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "performance_profile")); ok {
		result.PerformanceProfile = oci_database_migration.ReplicatPerformanceProfileEnum(performanceProfile.(string))
	}

	return result, nil
}

func CreateReplicatToMap(obj *oci_database_migration.CreateReplicat) map[string]interface{} {
	result := map[string]interface{}{}

	result["performance_profile"] = string(obj.PerformanceProfile)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateTargetTypeTablespaceDetails(fieldKeyFormat string) (oci_database_migration.CreateTargetTypeTablespaceDetails, error) {
	var baseObject oci_database_migration.CreateTargetTypeTablespaceDetails
	//discriminator
	targetTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_type"))
	var targetType string
	if ok {
		targetType = targetTypeRaw.(string)
	} else {
		targetType = "" // default value
	}
	switch strings.ToLower(targetType) {
	case strings.ToLower("ADB_D_AUTOCREATE"):
		details := oci_database_migration.CreateAdbDedicatedAutoCreateTablespaceDetails{}
		if blockSizeInKBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_size_in_kbs")); ok {
			details.BlockSizeInKBs = oci_database_migration.DataPumpTablespaceBlockSizesInKbEnum(blockSizeInKBs.(string))
		}
		if extendSizeInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extend_size_in_mbs")); ok {
			tmp := extendSizeInMBs.(int)
			details.ExtendSizeInMBs = &tmp
		}
		if isAutoCreate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_create")); ok {
			tmp := isAutoCreate.(bool)
			details.IsAutoCreate = &tmp
		}
		if isBigFile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_big_file")); ok {
			tmp := isBigFile.(bool)
			details.IsBigFile = &tmp
		}
		baseObject = details
	case strings.ToLower("ADB_D_REMAP"):
		details := oci_database_migration.CreateAdbDedicatedRemapTargetTablespaceDetails{}
		if remapTarget, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remap_target")); ok {
			tmp := remapTarget.(string)
			details.RemapTarget = &tmp
		}
		baseObject = details
	case strings.ToLower("ADB_S_REMAP"):
		details := oci_database_migration.CreateAdbServerlesTablespaceDetails{}
		baseObject = details
	case strings.ToLower("NON_ADB_AUTOCREATE"):
		details := oci_database_migration.CreateNonAdbAutoCreateTablespaceDetails{}
		if blockSizeInKBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_size_in_kbs")); ok {
			details.BlockSizeInKBs = oci_database_migration.DataPumpTablespaceBlockSizesInKbEnum(blockSizeInKBs.(string))
		}
		if extendSizeInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extend_size_in_mbs")); ok {
			tmp := extendSizeInMBs.(int)
			details.ExtendSizeInMBs = &tmp
		}
		if isAutoCreate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_create")); ok {
			tmp := isAutoCreate.(bool)
			details.IsAutoCreate = &tmp
		}
		if isBigFile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_big_file")); ok {
			tmp := isBigFile.(bool)
			details.IsBigFile = &tmp
		}
		baseObject = details
	case strings.ToLower("NON_ADB_REMAP"):
		details := oci_database_migration.CreateNonAdbRemapTablespaceDetails{}
		if remapTarget, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remap_target")); ok {
			tmp := remapTarget.(string)
			details.RemapTarget = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown target_type '%v' was specified", targetType)
	}
	return baseObject, nil
}

func CreateTargetTypeTablespaceDetailsToMap(obj *oci_database_migration.CreateTargetTypeTablespaceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_migration.CreateAdbDedicatedAutoCreateTablespaceDetails: //CreateADBDedicatedAutoCreateTablespaceDetails
		result["target_type"] = "ADB_D_AUTOCREATE"

		result["block_size_in_kbs"] = string(v.BlockSizeInKBs)

		if v.ExtendSizeInMBs != nil {
			result["extend_size_in_mbs"] = int(*v.ExtendSizeInMBs)
		}

		result["is_auto_create"] = bool(*v.IsAutoCreate)

		if v.IsBigFile != nil {
			result["is_big_file"] = bool(*v.IsBigFile)
		}
	case oci_database_migration.CreateAdbDedicatedRemapTargetTablespaceDetails:
		result["target_type"] = "ADB_D_REMAP"

		if v.RemapTarget != nil {
			result["remap_target"] = string(*v.RemapTarget)
		}
	case oci_database_migration.CreateAdbServerlesTablespaceDetails:
		result["target_type"] = "ADB_S_REMAP"
	case oci_database_migration.CreateNonAdbAutoCreateTablespaceDetails:
		result["target_type"] = "NON_ADB_AUTOCREATE"

		result["block_size_in_kbs"] = string(v.BlockSizeInKBs)

		if v.ExtendSizeInMBs != nil {
			result["extend_size_in_mbs"] = int(*v.ExtendSizeInMBs)
		}

		result["is_auto_create"] = bool(*v.IsAutoCreate)

		if v.IsBigFile != nil {
			result["is_big_file"] = bool(*v.IsBigFile)
		}
	case oci_database_migration.CreateNonAdbRemapTablespaceDetails:
		result["target_type"] = "NON_ADB_REMAP"

		if v.RemapTarget != nil {
			result["remap_target"] = string(*v.RemapTarget)
		}
	default:
		log.Printf("[WARN] Received 'target_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToDataPumpParameters(fieldKeyFormat string) (oci_database_migration.DataPumpParameters, error) {
	result := oci_database_migration.DataPumpParameters{}

	if estimate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "estimate")); ok {
		result.Estimate = oci_database_migration.DataPumpEstimateEnum(estimate.(string))
	}

	if excludeParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_parameters")); ok {
		interfaces := excludeParameters.([]interface{})
		tmp := make([]oci_database_migration.DataPumpExcludeParametersEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_database_migration.DataPumpExcludeParametersEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclude_parameters")) {
			result.ExcludeParameters = tmp
		}
	}

	exportParallelismDegree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_parallelism_degree"))
	if ok && exportParallelismDegree.(int) != 0 {
		tmp := exportParallelismDegree.(int)
		result.ExportParallelismDegree = &tmp
	} else {
		result.ExportParallelismDegree = nil
	}

	importParallelismDegree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "import_parallelism_degree"))
	if ok && importParallelismDegree.(int) != 0 {
		tmp := importParallelismDegree.(int)
		result.ImportParallelismDegree = &tmp
	} else {
		result.ImportParallelismDegree = nil
	}

	if isCluster, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_cluster")); ok {
		tmp := isCluster.(bool)
		result.IsCluster = &tmp
	}

	if tableExistsAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "table_exists_action")); ok {
		result.TableExistsAction = oci_database_migration.DataPumpTableExistsActionEnum(tableExistsAction.(string))
	}

	return result, nil
}

func DataPumpParametersToMap(obj *oci_database_migration.DataPumpParameters) map[string]interface{} {
	result := map[string]interface{}{}

	result["estimate"] = string(obj.Estimate)

	result["exclude_parameters"] = obj.ExcludeParameters

	if obj.ExportParallelismDegree != nil && int(*obj.ExportParallelismDegree) != 0 {
		result["export_parallelism_degree"] = int(*obj.ExportParallelismDegree)
	} else {
		result["export_parallelism_degree"] = nil
	}

	if obj.ImportParallelismDegree != nil && int(*obj.ImportParallelismDegree) != 0 {
		result["import_parallelism_degree"] = int(*obj.ImportParallelismDegree)
	} else {
		result["import_parallelism_degree"] = nil
	}

	if obj.IsCluster != nil {
		result["is_cluster"] = bool(*obj.IsCluster)
	}

	result["table_exists_action"] = string(obj.TableExistsAction)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToDirectoryObject(fieldKeyFormat string) (oci_database_migration.DirectoryObject, error) {
	result := oci_database_migration.DirectoryObject{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
		tmp := path.(string)
		result.Path = &tmp
	}

	return result, nil
}

func DirectoryObjectToMap(obj *oci_database_migration.DirectoryObject) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToExtract(fieldKeyFormat string) (oci_database_migration.Extract, error) {
	result := oci_database_migration.Extract{}

	longTransDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "long_trans_duration"))
	if ok && longTransDuration.(int) != 0 {
		tmp := longTransDuration.(int)
		result.LongTransDuration = &tmp
	} else {
		result.LongTransDuration = nil
	}

	if performanceProfile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "performance_profile")); ok {
		result.PerformanceProfile = oci_database_migration.ExtractPerformanceProfileEnum(performanceProfile.(string))
	}

	return result, nil
}

func ExtractToMap(obj *oci_database_migration.Extract) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LongTransDuration != nil && int(*obj.LongTransDuration) != 0 {
		result["long_trans_duration"] = int(*obj.LongTransDuration)
	} else {
		result["long_trans_duration"] = nil
	}

	result["performance_profile"] = string(obj.PerformanceProfile)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToGgsDeployment(fieldKeyFormat string) (oci_database_migration.GgsDeployment, error) {
	result := oci_database_migration.GgsDeployment{}

	if deploymentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deployment_id")); ok {
		tmp := deploymentId.(string)
		result.DeploymentId = &tmp
	}

	if ggsAdminCredentialsSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ggs_admin_credentials_secret_id")); ok {
		tmp := ggsAdminCredentialsSecretId.(string)
		result.GgsAdminCredentialsSecretId = &tmp
	}

	return result, nil
}

func GgsDeploymentToMap(obj *oci_database_migration.GgsDeployment) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DeploymentId != nil {
		result["deployment_id"] = string(*obj.DeploymentId)
	}

	if obj.GgsAdminCredentialsSecretId != nil {
		result["ggs_admin_credentials_secret_id"] = string(*obj.GgsAdminCredentialsSecretId)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToGoldenGateHubDetails(fieldKeyFormat string) (oci_database_migration.GoldenGateHubDetails, error) {
	result := oci_database_migration.GoldenGateHubDetails{}

	acceptableLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "acceptable_lag"))
	if ok && acceptableLag.(int) != 0 {
		tmp := acceptableLag.(int)
		result.AcceptableLag = &tmp
	} else {
		result.AcceptableLag = nil
	}

	if computeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_id")); ok {
		tmp := computeId.(string)
		result.ComputeId = &tmp
	}

	if extract, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extract")); ok {
		if tmpList := extract.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "extract"), 0)
			tmp, err := s.mapToExtract(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert extract, encountered error: %v", err)
			}
			result.Extract = &tmp
		}
	}

	if keyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_id")); ok {
		tmp := keyId.(string)
		result.KeyId = &tmp
	}

	if replicat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replicat")); ok {
		if tmpList := replicat.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "replicat"), 0)
			tmp, err := s.mapToReplicat(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert replicat, encountered error: %v", err)
			}
			result.Replicat = &tmp
		}
	}

	if restAdminCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rest_admin_credentials")); ok {
		if tmpList := restAdminCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rest_admin_credentials"), 0)
			tmp, err := s.mapToAdminCredentials(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert rest_admin_credentials, encountered error: %v", err)
			}
			result.RestAdminCredentials = &tmp
		}
	}

	if url, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url")); ok {
		tmp := url.(string)
		result.Url = &tmp
	}

	if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
		tmp := vaultId.(string)
		result.VaultId = &tmp
	}

	return result, nil
}

func GoldenGateHubDetailsToMap(obj *oci_database_migration.GoldenGateHubDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptableLag != nil {
		result["acceptable_lag"] = int(*obj.AcceptableLag)
	} else {
		result["acceptable_lag"] = nil
	}

	if obj.ComputeId != nil {
		result["compute_id"] = string(*obj.ComputeId)
	}

	if obj.Extract != nil {
		result["extract"] = []interface{}{ExtractToMap(obj.Extract)}
	}

	if obj.KeyId != nil {
		result["key_id"] = string(*obj.KeyId)
	}

	if obj.Replicat != nil {
		result["replicat"] = []interface{}{ReplicatToMap(obj.Replicat)}
	}

	if obj.RestAdminCredentials != nil {
		result["rest_admin_credentials"] = []interface{}{AdminCredentialsToMap(obj.RestAdminCredentials)}
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	if obj.VaultId != nil {
		result["vault_id"] = string(*obj.VaultId)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToHostDumpTransferDetails(fieldKeyFormat string) (oci_database_migration.HostDumpTransferDetails, error) {
	var baseObject oci_database_migration.HostDumpTransferDetails
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("CURL"):
		details := oci_database_migration.CurlTransferDetails{}
		if walletLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wallet_location")); ok {
			tmp := walletLocation.(string)
			details.WalletLocation = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_CLI"):
		details := oci_database_migration.OciCliDumpTransferDetails{}
		if ociHome, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oci_home")); ok {
			tmp := ociHome.(string)
			details.OciHome = &tmp
		}
		if walletLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wallet_location")); ok {
			tmp := walletLocation.(string)
			details.WalletLocation = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func HostDumpTransferDetailsToMap(obj *oci_database_migration.HostDumpTransferDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_migration.CurlTransferDetails:
		result["kind"] = "CURL"
		if v.WalletLocation != nil {
			result["wallet_location"] = string(*v.WalletLocation)
		}
	case oci_database_migration.OciCliDumpTransferDetails:
		result["kind"] = "OCI_CLI"

		if v.OciHome != nil {
			result["oci_home"] = string(*v.OciHome)
		}
		if v.WalletLocation != nil {
			result["wallet_location"] = string(*v.WalletLocation)
		}
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToMetadataRemap(fieldKeyFormat string) (oci_database_migration.MetadataRemap, error) {
	result := oci_database_migration.MetadataRemap{}

	if newValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "new_value")); ok {
		tmp := newValue.(string)
		result.NewValue = &tmp
	}

	if oldValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "old_value")); ok {
		tmp := oldValue.(string)
		result.OldValue = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_database_migration.MetadataRemapTypeEnum(type_.(string))
	}

	return result, nil
}

func MetadataRemapToMap(obj oci_database_migration.MetadataRemap) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NewValue != nil {
		result["new_value"] = string(*obj.NewValue)
	}

	if obj.OldValue != nil {
		result["old_value"] = string(*obj.OldValue)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToMigrationParameterDetails(fieldKeyFormat string) (oci_database_migration.MigrationParameterDetails, error) {
	result := oci_database_migration.MigrationParameterDetails{}

	if dataType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_type")); ok {
		result.DataType = oci_database_migration.AdvancedParameterDataTypesEnum(dataType.(string))
	}

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

func MigrationSummaryToMap(obj oci_database_migration.MigrationSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_migration.MySqlMigrationSummary:
		result["database_combination"] = "MYSQL"
	case oci_database_migration.OracleMigrationSummary:
		result["database_combination"] = "ORACLE"

		if v.SourceContainerDatabaseConnectionId != nil {
			result["source_container_database_connection_id"] = string(*v.SourceContainerDatabaseConnectionId)
		}
	default:
		log.Printf("[WARN] Received 'database_combination' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToMySqlAdvisorSettings(fieldKeyFormat string) (oci_database_migration.MySqlAdvisorSettings, error) {
	result := oci_database_migration.MySqlAdvisorSettings{}

	if isIgnoreErrors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ignore_errors")); ok {
		tmp := isIgnoreErrors.(bool)
		result.IsIgnoreErrors = &tmp
	}

	if isSkipAdvisor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_advisor")); ok {
		tmp := isSkipAdvisor.(bool)
		result.IsSkipAdvisor = &tmp
	}

	return result, nil
}

func MySqlAdvisorSettingsToMap(obj *oci_database_migration.MySqlAdvisorSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsIgnoreErrors != nil {
		result["is_ignore_errors"] = bool(*obj.IsIgnoreErrors)
	} else {
		result["is_ignore_errors"] = false
	}
	if obj.IsSkipAdvisor != nil {
		result["is_skip_advisor"] = bool(*obj.IsSkipAdvisor)
	} else {
		result["is_skip_advisor"] = false
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToMySqlDataTransferMediumDetails(fieldKeyFormat string) (oci_database_migration.MySqlDataTransferMediumDetails, error) {
	var baseObject oci_database_migration.MySqlDataTransferMediumDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_database_migration.MySqlObjectStorageDataTransferMediumDetails{}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func MySqlDataTransferMediumDetailsToMap(obj *oci_database_migration.MySqlDataTransferMediumDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_migration.MySqlObjectStorageDataTransferMediumDetails:
		result["type"] = "OBJECT_STORAGE"

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{ObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToMySqlDatabaseObject(fieldKeyFormat string) (oci_database_migration.MySqlDatabaseObject, error) {
	result := oci_database_migration.MySqlDatabaseObject{}

	if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
		tmp := object.(string)
		result.ObjectName = &tmp
	}

	if schema, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema")); ok {
		tmp := schema.(string)
		result.Schema = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func MySqlDatabaseObjectToMap(obj oci_database_migration.MySqlDatabaseObject) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	if obj.Schema != nil {
		result["schema"] = string(*obj.Schema)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToMySqlGgsDeploymentDetails(fieldKeyFormat string) (oci_database_migration.MySqlGgsDeploymentDetails, error) {
	result := oci_database_migration.MySqlGgsDeploymentDetails{}

	acceptableLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "acceptable_lag"))
	if ok && acceptableLag.(int) != 0 {
		tmp := acceptableLag.(int)
		result.AcceptableLag = &tmp
	} else {
		result.AcceptableLag = nil
	}

	if ggsDeployment, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ggs_deployment")); ok {
		if tmpList := ggsDeployment.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "ggs_deployment"), 0)
			tmp, err := s.mapToGgsDeployment(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert ggs_deployment, encountered error: %v", err)
			}
			result.GgsDeployment = &tmp
		}
	}

	if replicat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replicat")); ok {
		if tmpList := replicat.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "replicat"), 0)
			tmp, err := s.mapToReplicat(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert replicat, encountered error: %v", err)
			}
			result.Replicat = &tmp
		}
	}

	return result, nil
}

func MySqlGgsDeploymentDetailsToMap(obj *oci_database_migration.MySqlGgsDeploymentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptableLag != nil {
		result["acceptable_lag"] = int(*obj.AcceptableLag)
	} else {
		result["acceptable_lag"] = nil
	}

	if obj.GgsDeployment != nil {
		result["ggs_deployment"] = []interface{}{GgsDeploymentToMap(obj.GgsDeployment)}
	}

	if obj.Replicat != nil {
		result["replicat"] = []interface{}{ReplicatToMap(obj.Replicat)}
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToMySqlInitialLoadSettings(fieldKeyFormat string) (oci_database_migration.MySqlInitialLoadSettings, error) {
	result := oci_database_migration.MySqlInitialLoadSettings{}

	if compatibility, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compatibility")); ok {
		interfaces := compatibility.([]interface{})
		tmp := make([]oci_database_migration.CompatibilityOptionEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_database_migration.CompatibilityOptionEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "compatibility")) {
			result.Compatibility = tmp
		}
	}

	if handleGrantErrors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "handle_grant_errors")); ok {
		result.HandleGrantErrors = oci_database_migration.HandleGrantErrorsEnum(handleGrantErrors.(string))
	}

	if isConsistent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_consistent")); ok {
		tmp := isConsistent.(bool)
		result.IsConsistent = &tmp
	}

	if isIgnoreExistingObjects, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ignore_existing_objects")); ok {
		tmp := isIgnoreExistingObjects.(bool)
		result.IsIgnoreExistingObjects = &tmp
	}

	if isTzUtc, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_tz_utc")); ok {
		tmp := isTzUtc.(bool)
		result.IsTzUtc = &tmp
	}

	if jobMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_mode")); ok {
		result.JobMode = oci_database_migration.JobModeMySqlEnum(jobMode.(string))
	}

	if primaryKeyCompatibility, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_key_compatibility")); ok {
		result.PrimaryKeyCompatibility = oci_database_migration.PrimaryKeyCompatibilityEnum(primaryKeyCompatibility.(string))
	}

	return result, nil
}

func MySqlInitialLoadSettingsToMap(obj *oci_database_migration.MySqlInitialLoadSettings) map[string]interface{} {
	result := map[string]interface{}{}

	result["compatibility"] = obj.Compatibility

	result["handle_grant_errors"] = string(obj.HandleGrantErrors)

	if obj.IsConsistent != nil {
		result["is_consistent"] = bool(*obj.IsConsistent)
	}

	if obj.IsIgnoreExistingObjects != nil {
		result["is_ignore_existing_objects"] = bool(*obj.IsIgnoreExistingObjects)
	}

	if obj.IsTzUtc != nil {
		result["is_tz_utc"] = bool(*obj.IsTzUtc)
	}

	result["job_mode"] = string(obj.JobMode)

	result["primary_key_compatibility"] = string(obj.PrimaryKeyCompatibility)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToObjectStoreBucket(fieldKeyFormat string) (oci_database_migration.ObjectStoreBucket, error) {
	result := oci_database_migration.ObjectStoreBucket{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.NamespaceName = &tmp
	}

	return result, nil
}

func ObjectStoreBucketToMap(obj *oci_database_migration.ObjectStoreBucket) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToOracleAdvisorSettings(fieldKeyFormat string) (oci_database_migration.OracleAdvisorSettings, error) {
	result := oci_database_migration.OracleAdvisorSettings{}

	if isIgnoreErrors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ignore_errors")); ok {
		tmp := isIgnoreErrors.(bool)
		result.IsIgnoreErrors = &tmp
	}

	if isSkipAdvisor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_advisor")); ok {
		tmp := isSkipAdvisor.(bool)
		result.IsSkipAdvisor = &tmp
	}

	return result, nil
}

func OracleAdvisorSettingsToMap(obj *oci_database_migration.OracleAdvisorSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsIgnoreErrors != nil {
		result["is_ignore_errors"] = bool(*obj.IsIgnoreErrors)
	} else {
		result["is_ignore_errors"] = false
	}

	if obj.IsSkipAdvisor != nil {
		result["is_skip_advisor"] = bool(*obj.IsSkipAdvisor)
	} else {
		result["is_skip_advisor"] = false
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToOracleDataTransferMediumDetails(fieldKeyFormat string) (oci_database_migration.OracleDataTransferMediumDetails, error) {
	var baseObject oci_database_migration.OracleDataTransferMediumDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("AWS_S3"):
		details := oci_database_migration.OracleAwsS3DataTransferMediumDetails{}
		if accessKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if secretAccessKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_access_key")); ok {
			tmp := secretAccessKey.(string)
			details.SecretAccessKey = &tmp
		}
		baseObject = details
	case strings.ToLower("DBLINK"):
		details := oci_database_migration.OracleDbLinkDataTransferMediumDetails{}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("NFS"):
		details := oci_database_migration.OracleNfsDataTransferMediumDetails{}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		if sharedStorageMountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shared_storage_mount_target_id")); ok {
			tmp := sharedStorageMountTargetId.(string)
			details.SharedStorageMountTargetId = &tmp
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			if tmpList := source.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert source, encountered error: %v", err)
				}
				details.Source = tmp
			}
		}
		if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
			if tmpList := target.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "target"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert target, encountered error: %v", err)
				}
				details.Target = tmp
			}
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_database_migration.OracleObjectStorageDataTransferMediumDetails{}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			if tmpList := source.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert source, encountered error: %v", err)
				}
				details.Source = tmp
			}
		}
		if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
			if tmpList := target.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "target"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert target, encountered error: %v", err)
				}
				details.Target = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func OracleDataTransferMediumDetailsToMap(obj *oci_database_migration.OracleDataTransferMediumDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_migration.OracleAwsS3DataTransferMediumDetails:
		result["type"] = "AWS_S3"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{ObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.SecretAccessKey != nil {
			result["secret_access_key"] = string(*v.SecretAccessKey)
		}
	case oci_database_migration.OracleDbLinkDataTransferMediumDetails:
		result["type"] = "DBLINK"

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{ObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}
	case oci_database_migration.OracleNfsDataTransferMediumDetails:
		result["type"] = "NFS"

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{ObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}

		if v.SharedStorageMountTargetId != nil {
			result["shared_storage_mount_target_id"] = string(*v.SharedStorageMountTargetId)
		}

		if v.Source != nil {
			sourceArray := []interface{}{}
			if sourceMap := HostDumpTransferDetailsToMap(&v.Source); sourceMap != nil {
				sourceArray = append(sourceArray, sourceMap)
			}
			result["source"] = sourceArray
		}

		if v.Target != nil {
			targetArray := []interface{}{}
			if targetMap := HostDumpTransferDetailsToMap(&v.Target); targetMap != nil {
				targetArray = append(targetArray, targetMap)
			}
			result["target"] = targetArray
		}
	case oci_database_migration.OracleObjectStorageDataTransferMediumDetails:
		result["type"] = "OBJECT_STORAGE"

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{ObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}

		if v.Source != nil {
			sourceArray := []interface{}{}
			if sourceMap := HostDumpTransferDetailsToMap(&v.Source); sourceMap != nil {
				sourceArray = append(sourceArray, sourceMap)
			}
			result["source"] = sourceArray
		}

		if v.Target != nil {
			targetArray := []interface{}{}
			if targetMap := HostDumpTransferDetailsToMap(&v.Target); targetMap != nil {
				targetArray = append(targetArray, targetMap)
			}
			result["target"] = targetArray
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToOracleDatabaseObject(fieldKeyFormat string) (oci_database_migration.OracleDatabaseObject, error) {
	result := oci_database_migration.OracleDatabaseObject{}

	if isOmitExcludedTableFromReplication, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_omit_excluded_table_from_replication")); ok {
		tmp := isOmitExcludedTableFromReplication.(bool)
		result.IsOmitExcludedTableFromReplication = &tmp
	}

	if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
		tmp := object.(string)
		result.ObjectName = &tmp
	}

	if owner, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "owner")); ok {
		tmp := owner.(string)
		result.Owner = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func OracleDatabaseObjectToMap(obj oci_database_migration.OracleDatabaseObject) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsOmitExcludedTableFromReplication != nil {
		result["is_omit_excluded_table_from_replication"] = bool(*obj.IsOmitExcludedTableFromReplication)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	if obj.Owner != nil {
		result["owner"] = string(*obj.Owner)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToOracleGgsDeploymentDetails(fieldKeyFormat string) (oci_database_migration.OracleGgsDeploymentDetails, error) {
	result := oci_database_migration.OracleGgsDeploymentDetails{}

	acceptableLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "acceptable_lag"))
	if ok && acceptableLag.(int) != 0 {
		tmp := acceptableLag.(int)
		result.AcceptableLag = &tmp
	} else {
		result.AcceptableLag = nil
	}

	if extract, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extract")); ok {
		if tmpList := extract.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "extract"), 0)
			tmp, err := s.mapToExtract(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert extract, encountered error: %v", err)
			}
			result.Extract = &tmp
		}
	}

	if ggsDeployment, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ggs_deployment")); ok {
		if tmpList := ggsDeployment.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "ggs_deployment"), 0)
			tmp, err := s.mapToGgsDeployment(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert ggs_deployment, encountered error: %v", err)
			}
			result.GgsDeployment = &tmp
		}
	}

	if replicat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replicat")); ok {
		if tmpList := replicat.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "replicat"), 0)
			tmp, err := s.mapToReplicat(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert replicat, encountered error: %v", err)
			}
			result.Replicat = &tmp
		}
	}

	return result, nil
}

func OracleGgsDeploymentDetailsToMap(obj *oci_database_migration.OracleGgsDeploymentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptableLag != nil {
		result["acceptable_lag"] = int(*obj.AcceptableLag)
	} else {
		result["acceptable_lag"] = nil
	}

	if obj.Extract != nil {
		result["extract"] = []interface{}{ExtractToMap(obj.Extract)}
	}

	if obj.GgsDeployment != nil {
		result["ggs_deployment"] = []interface{}{GgsDeploymentToMap(obj.GgsDeployment)}
	}

	if obj.Replicat != nil {
		result["replicat"] = []interface{}{ReplicatToMap(obj.Replicat)}
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToOracleInitialLoadSettings(fieldKeyFormat string) (oci_database_migration.OracleInitialLoadSettings, error) {
	result := oci_database_migration.OracleInitialLoadSettings{}

	if dataPumpParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_pump_parameters")); ok {
		if tmpList := dataPumpParameters.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data_pump_parameters"), 0)
			tmp, err := s.mapToDataPumpParameters(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert data_pump_parameters, encountered error: %v", err)
			}
			result.DataPumpParameters = &tmp
		}
	}

	if exportDirectoryObject, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_directory_object")); ok {
		if tmpList := exportDirectoryObject.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "export_directory_object"), 0)
			tmp, err := s.mapToDirectoryObject(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert export_directory_object, encountered error: %v", err)
			}
			result.ExportDirectoryObject = &tmp
		}
	}

	if importDirectoryObject, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "import_directory_object")); ok {
		if tmpList := importDirectoryObject.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "import_directory_object"), 0)
			tmp, err := s.mapToDirectoryObject(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert import_directory_object, encountered error: %v", err)
			}
			result.ImportDirectoryObject = &tmp
		}
	}

	if jobMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_mode")); ok {
		result.JobMode = oci_database_migration.JobModeOracleEnum(jobMode.(string))
	}

	if metadataRemaps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata_remaps")); ok {
		interfaces := metadataRemaps.([]interface{})
		tmp := make([]oci_database_migration.MetadataRemap, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "metadata_remaps"), stateDataIndex)
			converted, err := s.mapToMetadataRemap(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "metadata_remaps")) {
			result.MetadataRemaps = tmp
		}
	}

	if tablespaceDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tablespace_details")); ok {
		if tmpList := tablespaceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tablespace_details"), 0)
			tmp, err := s.mapToTargetTypeTablespaceDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert tablespace_details, encountered error: %v", err)
			}
			result.TablespaceDetails = tmp
		}
	}

	return result, nil
}

func OracleInitialLoadSettingsToMap(obj *oci_database_migration.OracleInitialLoadSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DataPumpParameters != nil {
		result["data_pump_parameters"] = []interface{}{DataPumpParametersToMap(obj.DataPumpParameters)}
	}

	if obj.ExportDirectoryObject != nil {
		result["export_directory_object"] = []interface{}{DirectoryObjectToMap(obj.ExportDirectoryObject)}
	}

	if obj.ImportDirectoryObject != nil {
		result["import_directory_object"] = []interface{}{DirectoryObjectToMap(obj.ImportDirectoryObject)}
	}

	result["job_mode"] = string(obj.JobMode)

	metadataRemaps := []interface{}{}
	for _, item := range obj.MetadataRemaps {
		metadataRemaps = append(metadataRemaps, MetadataRemapToMap(item))
	}
	result["metadata_remaps"] = metadataRemaps

	if obj.TablespaceDetails != nil {
		tablespaceDetailsArray := []interface{}{}
		if tablespaceDetailsMap := TargetTypeTablespaceDetailsToMap(&obj.TablespaceDetails); tablespaceDetailsMap != nil {
			tablespaceDetailsArray = append(tablespaceDetailsArray, tablespaceDetailsMap)
		}
		result["tablespace_details"] = tablespaceDetailsArray
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToReplicat(fieldKeyFormat string) (oci_database_migration.Replicat, error) {
	result := oci_database_migration.Replicat{}

	if performanceProfile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "performance_profile")); ok {
		result.PerformanceProfile = oci_database_migration.ReplicatPerformanceProfileEnum(performanceProfile.(string))
	}

	return result, nil
}

func ReplicatToMap(obj *oci_database_migration.Replicat) map[string]interface{} {
	result := map[string]interface{}{}

	result["performance_profile"] = string(obj.PerformanceProfile)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToTargetTypeTablespaceDetails(fieldKeyFormat string) (oci_database_migration.TargetTypeTablespaceDetails, error) {
	var baseObject oci_database_migration.TargetTypeTablespaceDetails
	//discriminator
	targetTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_type"))
	var targetType string
	if ok {
		targetType = targetTypeRaw.(string)
	} else {
		targetType = "" // default value
	}
	switch strings.ToLower(targetType) {
	case strings.ToLower("ADB_D_AUTOCREATE"):
		details := oci_database_migration.AdbDedicatedAutoCreateTablespaceDetails{}
		if blockSizeInKBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_size_in_kbs")); ok {
			details.BlockSizeInKBs = oci_database_migration.DataPumpTablespaceBlockSizesInKbEnum(blockSizeInKBs.(string))
		}
		if extendSizeInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extend_size_in_mbs")); ok {
			tmp := extendSizeInMBs.(int)
			details.ExtendSizeInMBs = &tmp
		}
		if isAutoCreate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_create")); ok {
			tmp := isAutoCreate.(bool)
			details.IsAutoCreate = &tmp
		}
		if isBigFile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_big_file")); ok {
			tmp := isBigFile.(bool)
			details.IsBigFile = &tmp
		}
		baseObject = details
	case strings.ToLower("ADB_D_REMAP"):
		details := oci_database_migration.AdbDedicatedRemapTargetTablespaceDetails{}
		if remapTarget, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remap_target")); ok {
			tmp := remapTarget.(string)
			details.RemapTarget = &tmp
		}
		baseObject = details
	case strings.ToLower("ADB_S_REMAP"):
		details := oci_database_migration.AdbServerlesTablespaceDetails{}
		if remapTarget, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remap_target")); ok {
			details.RemapTarget = oci_database_migration.AdbServerlesTablespaceDetailsRemapTargetEnum(remapTarget.(string))
		}
		baseObject = details
	case strings.ToLower("NON_ADB_AUTOCREATE"):
		details := oci_database_migration.NonAdbAutoCreateTablespaceDetails{}
		if blockSizeInKBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_size_in_kbs")); ok {
			details.BlockSizeInKBs = oci_database_migration.DataPumpTablespaceBlockSizesInKbEnum(blockSizeInKBs.(string))
		}
		if extendSizeInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extend_size_in_mbs")); ok {
			tmp := extendSizeInMBs.(int)
			details.ExtendSizeInMBs = &tmp
		}
		if isAutoCreate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_create")); ok {
			tmp := isAutoCreate.(bool)
			details.IsAutoCreate = &tmp
		}
		if isBigFile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_big_file")); ok {
			tmp := isBigFile.(bool)
			details.IsBigFile = &tmp
		}
		baseObject = details
	case strings.ToLower("NON_ADB_REMAP"):
		details := oci_database_migration.NonAdbRemapTablespaceDetails{}
		if remapTarget, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remap_target")); ok {
			tmp := remapTarget.(string)
			details.RemapTarget = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown target_type '%v' was specified", targetType)
	}
	return baseObject, nil
}

func TargetTypeTablespaceDetailsToMap(obj *oci_database_migration.TargetTypeTablespaceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_migration.AdbDedicatedAutoCreateTablespaceDetails:
		result["target_type"] = "ADB_D_AUTOCREATE"

		result["block_size_in_kbs"] = string(v.BlockSizeInKBs)

		if v.ExtendSizeInMBs != nil {
			result["extend_size_in_mbs"] = int(*v.ExtendSizeInMBs)
		}

		result["is_auto_create"] = bool(*v.IsAutoCreate)

		if v.IsBigFile != nil {
			result["is_big_file"] = bool(*v.IsBigFile)
		}
	case oci_database_migration.AdbDedicatedRemapTargetTablespaceDetails:
		result["target_type"] = "ADB_D_REMAP"

		if v.RemapTarget != nil {
			result["remap_target"] = string(*v.RemapTarget)
		}
	case oci_database_migration.AdbServerlesTablespaceDetails:
		result["target_type"] = "ADB_S_REMAP"

		result["remap_target"] = string(v.RemapTarget)
	case oci_database_migration.NonAdbAutoCreateTablespaceDetails:
		result["target_type"] = "NON_ADB_AUTOCREATE"

		result["block_size_in_kbs"] = string(v.BlockSizeInKBs)

		if v.ExtendSizeInMBs != nil {
			result["extend_size_in_mbs"] = int(*v.ExtendSizeInMBs)
		}

		result["is_auto_create"] = bool(*v.IsAutoCreate)

		if v.IsBigFile != nil {
			result["is_big_file"] = bool(*v.IsBigFile)
		}
	case oci_database_migration.NonAdbRemapTablespaceDetails:
		result["target_type"] = "NON_ADB_REMAP"

		if v.RemapTarget != nil {
			result["remap_target"] = string(*v.RemapTarget)
		}
	default:
		log.Printf("[WARN] Received 'target_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateAdminCredentials(fieldKeyFormat string) (oci_database_migration.UpdateAdminCredentials, error) {
	result := oci_database_migration.UpdateAdminCredentials{}

	if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
		tmp := password.(string)
		result.Password = &tmp
	}

	if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
		tmp := username.(string)
		result.Username = &tmp
	}

	return result, nil
}

func UpdateAdminCredentialsToMap(obj *oci_database_migration.UpdateAdminCredentials) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Password != nil {
		result["password"] = string(*obj.Password)
	}

	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateDataPumpParameters(fieldKeyFormat string) (oci_database_migration.UpdateDataPumpParameters, error) {
	result := oci_database_migration.UpdateDataPumpParameters{}

	if estimate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "estimate")); ok {
		result.Estimate = oci_database_migration.DataPumpEstimateEnum(estimate.(string))
	}

	if excludeParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_parameters")); ok {
		interfaces := excludeParameters.([]interface{})
		tmp := make([]oci_database_migration.DataPumpExcludeParametersEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_database_migration.DataPumpExcludeParametersEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclude_parameters")) {
			result.ExcludeParameters = tmp
		}
	}

	exportParallelismDegree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_parallelism_degree"))
	if ok && exportParallelismDegree.(int) != 0 {
		tmp := exportParallelismDegree.(int)
		result.ExportParallelismDegree = &tmp
	} else {
		result.ExportParallelismDegree = nil
	}

	importParallelismDegree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "import_parallelism_degree"))
	if ok && importParallelismDegree.(int) != 0 {
		tmp := importParallelismDegree.(int)
		result.ImportParallelismDegree = &tmp
	} else {
		result.ImportParallelismDegree = nil
	}

	if isCluster, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_cluster")); ok {
		tmp := isCluster.(bool)
		result.IsCluster = &tmp
	}

	if tableExistsAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "table_exists_action")); ok {
		result.TableExistsAction = oci_database_migration.DataPumpTableExistsActionEnum(tableExistsAction.(string))
	}

	return result, nil
}

func UpdateDataPumpParametersToMap(obj *oci_database_migration.UpdateDataPumpParameters) map[string]interface{} {
	result := map[string]interface{}{}

	result["estimate"] = string(obj.Estimate)

	result["exclude_parameters"] = obj.ExcludeParameters

	if obj.ExportParallelismDegree != nil && int(*obj.ExportParallelismDegree) != 0 {
		result["export_parallelism_degree"] = int(*obj.ExportParallelismDegree)
	} else {
		result["export_parallelism_degree"] = nil
	}

	if obj.ImportParallelismDegree != nil && int(*obj.ImportParallelismDegree) != 0 {
		result["import_parallelism_degree"] = int(*obj.ImportParallelismDegree)
	} else {
		result["import_parallelism_degree"] = nil
	}

	if obj.IsCluster != nil {
		result["is_cluster"] = bool(*obj.IsCluster)
	}

	result["table_exists_action"] = string(obj.TableExistsAction)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateDirectoryObject(fieldKeyFormat string) (oci_database_migration.UpdateDirectoryObject, error) {
	result := oci_database_migration.UpdateDirectoryObject{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path"))
	if ok && path.(string) != "" {
		tmp := path.(string)
		result.Path = &tmp
	}

	return result, nil
}

func UpdateDirectoryObjectToMap(obj *oci_database_migration.UpdateDirectoryObject) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateExtract(fieldKeyFormat string) (oci_database_migration.UpdateExtract, error) {
	result := oci_database_migration.UpdateExtract{}

	longTransDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "long_trans_duration"))
	if ok && longTransDuration.(int) != 0 {
		tmp := longTransDuration.(int)
		result.LongTransDuration = &tmp
	} else {
		result.LongTransDuration = nil
	}

	if performanceProfile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "performance_profile")); ok {
		result.PerformanceProfile = oci_database_migration.ExtractPerformanceProfileEnum(performanceProfile.(string))
	}

	return result, nil
}

func UpdateExtractToMap(obj *oci_database_migration.UpdateExtract) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LongTransDuration != nil && int(*obj.LongTransDuration) != 0 {
		result["long_trans_duration"] = int(*obj.LongTransDuration)
	} else {
		result["long_trans_duration"] = nil
	}

	result["performance_profile"] = string(obj.PerformanceProfile)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateGoldenGateHubDetails(fieldKeyFormat string) (oci_database_migration.UpdateGoldenGateHubDetails, error) {
	result := oci_database_migration.UpdateGoldenGateHubDetails{}

	acceptableLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "acceptable_lag"))
	if ok && acceptableLag.(int) != 0 {
		tmp := acceptableLag.(int)
		result.AcceptableLag = &tmp
	} else {
		result.AcceptableLag = nil
	}

	if computeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_id")); ok {
		tmp := computeId.(string)
		result.ComputeId = &tmp
	}

	if extract, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extract")); ok {
		if tmpList := extract.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "extract"), 0)
			tmp, err := s.mapToUpdateExtract(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert extract, encountered error: %v", err)
			}
			result.Extract = &tmp
		}
	}

	if keyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_id")); ok {
		tmp := keyId.(string)
		result.KeyId = &tmp
	}

	if replicat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replicat")); ok {
		if tmpList := replicat.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "replicat"), 0)
			tmp, err := s.mapToUpdateReplicat(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert replicat, encountered error: %v", err)
			}
			result.Replicat = &tmp
		}
	}

	if restAdminCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rest_admin_credentials")); ok {
		if tmpList := restAdminCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rest_admin_credentials"), 0)
			tmp, err := s.mapToUpdateAdminCredentials(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert rest_admin_credentials, encountered error: %v", err)
			}
			result.RestAdminCredentials = &tmp
		}
	}

	if url, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url")); ok {
		tmp := url.(string)
		result.Url = &tmp
	}

	if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
		tmp := vaultId.(string)
		result.VaultId = &tmp
	}

	return result, nil
}

func UpdateGoldenGateHubDetailsToMap(obj *oci_database_migration.UpdateGoldenGateHubDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptableLag != nil {
		result["acceptable_lag"] = int(*obj.AcceptableLag)
	} else {
		result["acceptable_lag"] = nil
	}

	if obj.ComputeId != nil {
		result["compute_id"] = string(*obj.ComputeId)
	}

	if obj.Extract != nil {
		result["extract"] = []interface{}{UpdateExtractToMap(obj.Extract)}
	}

	if obj.KeyId != nil {
		result["key_id"] = string(*obj.KeyId)
	}

	if obj.Replicat != nil {
		result["replicat"] = []interface{}{UpdateReplicatToMap(obj.Replicat)}
	}

	if obj.RestAdminCredentials != nil {
		result["rest_admin_credentials"] = []interface{}{UpdateAdminCredentialsToMap(obj.RestAdminCredentials)}
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	if obj.VaultId != nil {
		result["vault_id"] = string(*obj.VaultId)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateMySqlAdvisorSettings(fieldKeyFormat string) (oci_database_migration.UpdateMySqlAdvisorSettings, error) {
	result := oci_database_migration.UpdateMySqlAdvisorSettings{}

	if isIgnoreErrors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ignore_errors")); ok {
		tmp := isIgnoreErrors.(bool)
		result.IsIgnoreErrors = &tmp
	}

	if isSkipAdvisor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_advisor")); ok {
		tmp := isSkipAdvisor.(bool)
		result.IsSkipAdvisor = &tmp
	}

	return result, nil
}

func UpdateMySqlAdvisorSettingsToMap(obj *oci_database_migration.UpdateMySqlAdvisorSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsIgnoreErrors != nil {
		result["is_ignore_errors"] = bool(*obj.IsIgnoreErrors)
	}

	if obj.IsSkipAdvisor != nil {
		result["is_skip_advisor"] = bool(*obj.IsSkipAdvisor)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateMySqlDataTransferMediumDetails(fieldKeyFormat string) (oci_database_migration.UpdateMySqlDataTransferMediumDetails, error) {
	var baseObject oci_database_migration.UpdateMySqlDataTransferMediumDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_database_migration.UpdateMySqlObjectStorageDataTransferMediumDetails{}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToUpdateObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func UpdateMySqlDataTransferMediumDetailsToMap(obj *oci_database_migration.UpdateMySqlDataTransferMediumDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_migration.UpdateMySqlObjectStorageDataTransferMediumDetails:
		result["type"] = "OBJECT_STORAGE"

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{UpdateObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateMySqlGgsDeploymentDetails(fieldKeyFormat string) (oci_database_migration.UpdateMySqlGgsDeploymentDetails, error) {
	result := oci_database_migration.UpdateMySqlGgsDeploymentDetails{}

	acceptableLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "acceptable_lag"))
	if ok && acceptableLag.(int) != 0 {
		tmp := acceptableLag.(int)
		result.AcceptableLag = &tmp
	} else {
		result.AcceptableLag = nil
	}

	if replicat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replicat")); ok {
		if tmpList := replicat.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "replicat"), 0)
			tmp, err := s.mapToUpdateReplicat(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert replicat, encountered error: %v", err)
			}
			result.Replicat = &tmp
		}
	}

	return result, nil
}

func UpdateMySqlGgsDeploymentDetailsToMap(obj *oci_database_migration.UpdateMySqlGgsDeploymentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptableLag != nil {
		result["acceptable_lag"] = int(*obj.AcceptableLag)
	} else {
		result["acceptable_lag"] = nil
	}

	if obj.Replicat != nil {
		result["replicat"] = []interface{}{UpdateReplicatToMap(obj.Replicat)}
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateMySqlInitialLoadSettings(fieldKeyFormat string) (oci_database_migration.UpdateMySqlInitialLoadSettings, error) {
	result := oci_database_migration.UpdateMySqlInitialLoadSettings{}

	if compatibility, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compatibility")); ok {
		interfaces := compatibility.([]interface{})
		tmp := make([]oci_database_migration.CompatibilityOptionEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_database_migration.CompatibilityOptionEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "compatibility")) {
			result.Compatibility = tmp
		}
	}

	if handleGrantErrors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "handle_grant_errors")); ok {
		result.HandleGrantErrors = oci_database_migration.HandleGrantErrorsEnum(handleGrantErrors.(string))
	}

	if isConsistent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_consistent")); ok {
		tmp := isConsistent.(bool)
		result.IsConsistent = &tmp
	}

	if isIgnoreExistingObjects, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ignore_existing_objects")); ok {
		tmp := isIgnoreExistingObjects.(bool)
		result.IsIgnoreExistingObjects = &tmp
	}

	if isTzUtc, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_tz_utc")); ok {
		tmp := isTzUtc.(bool)
		result.IsTzUtc = &tmp
	}

	if jobMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_mode")); ok {
		result.JobMode = oci_database_migration.JobModeMySqlEnum(jobMode.(string))
	}

	if primaryKeyCompatibility, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_key_compatibility")); ok {
		result.PrimaryKeyCompatibility = oci_database_migration.PrimaryKeyCompatibilityEnum(primaryKeyCompatibility.(string))
	}

	return result, nil
}

func UpdateMySqlInitialLoadSettingsToMap(obj *oci_database_migration.UpdateMySqlInitialLoadSettings) map[string]interface{} {
	result := map[string]interface{}{}

	result["compatibility"] = obj.Compatibility

	result["handle_grant_errors"] = string(obj.HandleGrantErrors)

	if obj.IsConsistent != nil {
		result["is_consistent"] = bool(*obj.IsConsistent)
	}

	if obj.IsIgnoreExistingObjects != nil {
		result["is_ignore_existing_objects"] = bool(*obj.IsIgnoreExistingObjects)
	}

	if obj.IsTzUtc != nil {
		result["is_tz_utc"] = bool(*obj.IsTzUtc)
	}

	result["job_mode"] = string(obj.JobMode)

	result["primary_key_compatibility"] = string(obj.PrimaryKeyCompatibility)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateObjectStoreBucket(fieldKeyFormat string) (oci_database_migration.UpdateObjectStoreBucket, error) {
	result := oci_database_migration.UpdateObjectStoreBucket{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.NamespaceName = &tmp
	}

	return result, nil
}

func UpdateObjectStoreBucketToMap(obj *oci_database_migration.UpdateObjectStoreBucket) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateOracleAdvisorSettings(fieldKeyFormat string) (oci_database_migration.UpdateOracleAdvisorSettings, error) {
	result := oci_database_migration.UpdateOracleAdvisorSettings{}

	if isIgnoreErrors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ignore_errors")); ok {
		tmp := isIgnoreErrors.(bool)
		result.IsIgnoreErrors = &tmp
	}

	if isSkipAdvisor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_advisor")); ok {
		tmp := isSkipAdvisor.(bool)
		result.IsSkipAdvisor = &tmp
	}

	return result, nil
}

func UpdateOracleAdvisorSettingsToMap(obj *oci_database_migration.UpdateOracleAdvisorSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsIgnoreErrors != nil {
		result["is_ignore_errors"] = bool(*obj.IsIgnoreErrors)
	}

	if obj.IsSkipAdvisor != nil {
		result["is_skip_advisor"] = bool(*obj.IsSkipAdvisor)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateOracleDataTransferMediumDetails(fieldKeyFormat string) (oci_database_migration.UpdateOracleDataTransferMediumDetails, error) {
	var baseObject oci_database_migration.UpdateOracleDataTransferMediumDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("AWS_S3"):
		details := oci_database_migration.UpdateOracleAwsS3DataTransferMediumDetails{}
		if accessKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if secretAccessKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_access_key")); ok {
			tmp := secretAccessKey.(string)
			details.SecretAccessKey = &tmp
		}
		baseObject = details
	case strings.ToLower("DBLINK"):
		details := oci_database_migration.UpdateOracleDbLinkDataTransferMediumDetails{}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToUpdateObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("NFS"):
		details := oci_database_migration.UpdateOracleNfsDataTransferMediumDetails{}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToUpdateObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}

		sharedStorageMountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shared_storage_mount_target_id"))
		if ok && sharedStorageMountTargetId.(string) != "" {
			tmp := sharedStorageMountTargetId.(string)
			details.SharedStorageMountTargetId = &tmp
		}

		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			if tmpList := source.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert source, encountered error: %v", err)
				}
				details.Source = tmp
			}
		}
		if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
			if tmpList := target.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "target"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert target, encountered error: %v", err)
				}
				details.Target = tmp
			}
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_database_migration.UpdateOracleObjectStorageDataTransferMediumDetails{}
		if objectStorageBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket")); ok {
			if tmpList := objectStorageBucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_bucket"), 0)
				tmp, err := s.mapToUpdateObjectStoreBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_bucket, encountered error: %v", err)
				}
				details.ObjectStorageBucket = &tmp
			}
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			if tmpList := source.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert source, encountered error: %v", err)
				}
				details.Source = tmp
			}
		}
		if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
			if tmpList := target.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "target"), 0)
				tmp, err := s.mapToHostDumpTransferDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert target, encountered error: %v", err)
				}
				details.Target = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func UpdateOracleDataTransferMediumDetailsToMap(obj *oci_database_migration.UpdateOracleDataTransferMediumDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_migration.UpdateOracleAwsS3DataTransferMediumDetails:
		result["type"] = "AWS_S3"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{ObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.SecretAccessKey != nil {
			result["secret_access_key"] = string(*v.SecretAccessKey)
		}
	case oci_database_migration.UpdateOracleDbLinkDataTransferMediumDetails:
		result["type"] = "DBLINK"

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{UpdateObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}
	case oci_database_migration.UpdateOracleNfsDataTransferMediumDetails:
		result["type"] = "NFS"

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{UpdateObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}

		if v.SharedStorageMountTargetId != nil {
			result["shared_storage_mount_target_id"] = string(*v.SharedStorageMountTargetId)
		}

		if v.Source != nil {
			sourceArray := []interface{}{}
			if sourceMap := HostDumpTransferDetailsToMap(&v.Source); sourceMap != nil {
				sourceArray = append(sourceArray, sourceMap)
			}
			result["source"] = sourceArray
		}

		if v.Target != nil {
			targetArray := []interface{}{}
			if targetMap := HostDumpTransferDetailsToMap(&v.Target); targetMap != nil {
				targetArray = append(targetArray, targetMap)
			}
			result["target"] = targetArray
		}
	case oci_database_migration.UpdateOracleObjectStorageDataTransferMediumDetails:
		result["type"] = "OBJECT_STORAGE"

		if v.ObjectStorageBucket != nil {
			result["object_storage_bucket"] = []interface{}{UpdateObjectStoreBucketToMap(v.ObjectStorageBucket)}
		}

		if v.Source != nil {
			sourceArray := []interface{}{}
			if sourceMap := HostDumpTransferDetailsToMap(&v.Source); sourceMap != nil {
				sourceArray = append(sourceArray, sourceMap)
			}
			result["source"] = sourceArray
		}

		if v.Target != nil {
			targetArray := []interface{}{}
			if targetMap := HostDumpTransferDetailsToMap(&v.Target); targetMap != nil {
				targetArray = append(targetArray, targetMap)
			}
			result["target"] = targetArray
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateOracleGgsDeploymentDetails(fieldKeyFormat string) (oci_database_migration.UpdateOracleGgsDeploymentDetails, error) {
	result := oci_database_migration.UpdateOracleGgsDeploymentDetails{}

	acceptableLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "acceptable_lag"))
	if ok && acceptableLag.(int) != 0 {
		tmp := acceptableLag.(int)
		result.AcceptableLag = &tmp
	} else {
		result.AcceptableLag = nil
	}

	if extract, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extract")); ok {
		if tmpList := extract.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "extract"), 0)
			tmp, err := s.mapToUpdateExtract(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert extract, encountered error: %v", err)
			}
			result.Extract = &tmp
		}
	}

	if replicat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replicat")); ok {
		if tmpList := replicat.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "replicat"), 0)
			tmp, err := s.mapToUpdateReplicat(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert replicat, encountered error: %v", err)
			}
			result.Replicat = &tmp
		}
	}

	return result, nil
}

func UpdateOracleGgsDeploymentDetailsToMap(obj *oci_database_migration.UpdateOracleGgsDeploymentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptableLag != nil {
		result["acceptable_lag"] = int(*obj.AcceptableLag)
	} else {
		result["acceptable_lag"] = nil
	}

	if obj.Extract != nil {
		result["extract"] = []interface{}{UpdateExtractToMap(obj.Extract)}
	}

	if obj.Replicat != nil {
		result["replicat"] = []interface{}{UpdateReplicatToMap(obj.Replicat)}
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateOracleInitialLoadSettings(fieldKeyFormat string) (oci_database_migration.UpdateOracleInitialLoadSettings, error) {
	result := oci_database_migration.UpdateOracleInitialLoadSettings{}

	if dataPumpParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_pump_parameters")); ok {
		if tmpList := dataPumpParameters.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data_pump_parameters"), 0)
			tmp, err := s.mapToUpdateDataPumpParameters(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert data_pump_parameters, encountered error: %v", err)
			}
			result.DataPumpParameters = &tmp
		}
	}

	if exportDirectoryObject, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_directory_object")); ok {
		if tmpList := exportDirectoryObject.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "export_directory_object"), 0)
			tmp, err := s.mapToUpdateDirectoryObject(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert export_directory_object, encountered error: %v", err)
			}
			result.ExportDirectoryObject = &tmp
		}
	}

	if importDirectoryObject, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "import_directory_object")); ok {
		if tmpList := importDirectoryObject.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "import_directory_object"), 0)
			tmp, err := s.mapToUpdateDirectoryObject(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert import_directory_object, encountered error: %v", err)
			}
			result.ImportDirectoryObject = &tmp
		}
	}

	if jobMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_mode")); ok {
		result.JobMode = oci_database_migration.JobModeOracleEnum(jobMode.(string))
	}

	if metadataRemaps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata_remaps")); ok {
		interfaces := metadataRemaps.([]interface{})
		tmp := make([]oci_database_migration.MetadataRemap, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "metadata_remaps"), stateDataIndex)
			converted, err := s.mapToMetadataRemap(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "metadata_remaps")) {
			result.MetadataRemaps = tmp
		}
	}

	if tablespaceDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tablespace_details")); ok {
		if tmpList := tablespaceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tablespace_details"), 0)
			tmp, err := s.mapToUpdateTargetTypeTablespaceDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert tablespace_details, encountered error: %v", err)
			}
			result.TablespaceDetails = tmp
		}
	}

	return result, nil
}

func UpdateOracleInitialLoadSettingsToMap(obj *oci_database_migration.UpdateOracleInitialLoadSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DataPumpParameters != nil {
		result["data_pump_parameters"] = []interface{}{UpdateDataPumpParametersToMap(obj.DataPumpParameters)}
	}

	if obj.ExportDirectoryObject != nil {
		result["export_directory_object"] = []interface{}{UpdateDirectoryObjectToMap(obj.ExportDirectoryObject)}
	}

	if obj.ImportDirectoryObject != nil {
		result["import_directory_object"] = []interface{}{UpdateDirectoryObjectToMap(obj.ImportDirectoryObject)}
	}

	result["job_mode"] = string(obj.JobMode)

	metadataRemaps := []interface{}{}
	for _, item := range obj.MetadataRemaps {
		metadataRemaps = append(metadataRemaps, MetadataRemapToMap(item))
	}
	result["metadata_remaps"] = metadataRemaps

	if obj.TablespaceDetails != nil {
		tablespaceDetailsArray := []interface{}{}
		if tablespaceDetailsMap := UpdateTargetTypeTablespaceDetailsToMap(&obj.TablespaceDetails); tablespaceDetailsMap != nil {
			tablespaceDetailsArray = append(tablespaceDetailsArray, tablespaceDetailsMap)
		}
		result["tablespace_details"] = tablespaceDetailsArray
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateReplicat(fieldKeyFormat string) (oci_database_migration.UpdateReplicat, error) {
	result := oci_database_migration.UpdateReplicat{}

	if performanceProfile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "performance_profile")); ok {
		result.PerformanceProfile = oci_database_migration.ReplicatPerformanceProfileEnum(performanceProfile.(string))
	}

	return result, nil
}

func UpdateReplicatToMap(obj *oci_database_migration.UpdateReplicat) map[string]interface{} {
	result := map[string]interface{}{}

	result["performance_profile"] = string(obj.PerformanceProfile)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateTargetTypeTablespaceDetails(fieldKeyFormat string) (oci_database_migration.UpdateTargetTypeTablespaceDetails, error) {
	var baseObject oci_database_migration.UpdateTargetTypeTablespaceDetails
	//discriminator
	targetTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_type"))
	var targetType string
	if ok {
		targetType = targetTypeRaw.(string)
	} else {
		targetType = "" // default value
	}
	switch strings.ToLower(targetType) {
	case strings.ToLower("ADB_D_AUTOCREATE"):
		details := oci_database_migration.UpdateAdbDedicatedAutoCreateTablespaceDetails{}
		if blockSizeInKBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_size_in_kbs")); ok {
			details.BlockSizeInKBs = oci_database_migration.DataPumpTablespaceBlockSizesInKbEnum(blockSizeInKBs.(string))
		}
		if extendSizeInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extend_size_in_mbs")); ok {
			tmp := extendSizeInMBs.(int)
			details.ExtendSizeInMBs = &tmp
		}
		if isAutoCreate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_create")); ok {
			tmp := isAutoCreate.(bool)
			details.IsAutoCreate = &tmp
		}
		if isBigFile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_big_file")); ok {
			tmp := isBigFile.(bool)
			details.IsBigFile = &tmp
		}
		baseObject = details
	case strings.ToLower("ADB_D_REMAP"):
		details := oci_database_migration.UpdateAdbDedicatedRemapTargetTablespaceDetails{}
		remapTarget, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remap_target"))
		if ok && remapTarget.(string) != "" {
			tmp := remapTarget.(string)
			details.RemapTarget = &tmp
		}
		baseObject = details
	case strings.ToLower("ADB_S_REMAP"):
		details := oci_database_migration.UpdateAdbServerlesTablespaceDetails{}
		baseObject = details
	case strings.ToLower("NON_ADB_AUTOCREATE"):
		details := oci_database_migration.UpdateNonAdbAutoCreateTablespaceDetails{}
		if blockSizeInKBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_size_in_kbs")); ok {
			details.BlockSizeInKBs = oci_database_migration.DataPumpTablespaceBlockSizesInKbEnum(blockSizeInKBs.(string))
		}
		if extendSizeInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extend_size_in_mbs")); ok {
			tmp := extendSizeInMBs.(int)
			details.ExtendSizeInMBs = &tmp
		}
		if isAutoCreate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_create")); ok {
			tmp := isAutoCreate.(bool)
			details.IsAutoCreate = &tmp
		}
		if isBigFile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_big_file")); ok {
			tmp := isBigFile.(bool)
			details.IsBigFile = &tmp
		}
		baseObject = details
	case strings.ToLower("NON_ADB_REMAP"):
		details := oci_database_migration.UpdateNonAdbRemapTablespaceDetails{}
		remapTarget, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remap_target"))
		if ok && remapTarget.(string) != "" {
			tmp := remapTarget.(string)
			details.RemapTarget = &tmp
		}
		baseObject = details
	case strings.ToLower("TARGET_DEFAULTS_AUTOCREATE"):
		details := oci_database_migration.UpdateTargetDefaultsAutoCreateTablespaceDetails{}
		baseObject = details
	case strings.ToLower("TARGET_DEFAULTS_REMAP"):
		details := oci_database_migration.UpdateTargetDefaultsRemapTablespaceDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown target_type '%v' was specified", targetType)
	}
	return baseObject, nil
}

func UpdateTargetTypeTablespaceDetailsToMap(obj *oci_database_migration.UpdateTargetTypeTablespaceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_migration.UpdateAdbDedicatedAutoCreateTablespaceDetails:
		result["target_type"] = "ADB_D_AUTOCREATE"

		result["block_size_in_kbs"] = string(v.BlockSizeInKBs)

		if v.ExtendSizeInMBs != nil {
			result["extend_size_in_mbs"] = int(*v.ExtendSizeInMBs)
		}

		result["is_auto_create"] = bool(*v.IsAutoCreate)

		if v.IsBigFile != nil {
			result["is_big_file"] = bool(*v.IsBigFile)
		}
	case oci_database_migration.UpdateAdbDedicatedRemapTargetTablespaceDetails:
		result["target_type"] = "ADB_D_REMAP"

		if v.RemapTarget != nil {
			result["remap_target"] = string(*v.RemapTarget)
		}
	case oci_database_migration.UpdateAdbServerlesTablespaceDetails:
		result["target_type"] = "ADB_S_REMAP"
	case oci_database_migration.UpdateNonAdbAutoCreateTablespaceDetails:
		result["target_type"] = "NON_ADB_AUTOCREATE"

		result["block_size_in_kbs"] = string(v.BlockSizeInKBs)

		if v.ExtendSizeInMBs != nil {
			result["extend_size_in_mbs"] = int(*v.ExtendSizeInMBs)
		}

		result["is_auto_create"] = bool(*v.IsAutoCreate)

		if v.IsBigFile != nil {
			result["is_big_file"] = bool(*v.IsBigFile)
		}
	case oci_database_migration.UpdateNonAdbRemapTablespaceDetails:
		result["target_type"] = "NON_ADB_REMAP"

		if v.RemapTarget != nil {
			result["remap_target"] = string(*v.RemapTarget)
		}
	case oci_database_migration.UpdateTargetDefaultsAutoCreateTablespaceDetails:
		result["target_type"] = "TARGET_DEFAULTS_AUTOCREATE"
	case oci_database_migration.UpdateTargetDefaultsRemapTablespaceDetails:
		result["target_type"] = "TARGET_DEFAULTS_REMAP"
	default:
		log.Printf("[WARN] Received 'target_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) populateTopLevelPolymorphicCreateMigrationRequest(request *oci_database_migration.CreateMigrationRequest) error {
	//discriminator
	databaseCombinationRaw, ok := s.D.GetOkExists("database_combination")
	var databaseCombination string
	if ok {
		databaseCombination = databaseCombinationRaw.(string)
	} else {
		databaseCombination = "" // default value
	}
	switch strings.ToLower(databaseCombination) {
	case strings.ToLower("MYSQL"):
		details := oci_database_migration.CreateMySqlMigrationDetails{}
		if advisorSettings, ok := s.D.GetOkExists("advisor_settings"); ok {
			if tmpList := advisorSettings.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "advisor_settings", 0)
				tmp, err := s.mapToCreateMySqlAdvisorSettings(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AdvisorSettings = &tmp
			}
		}
		if bulkIncludeExcludeData, ok := s.D.GetOkExists("bulk_include_exclude_data"); ok {
			tmp := bulkIncludeExcludeData.(string)
			details.BulkIncludeExcludeData = &tmp
		}
		if dataTransferMediumDetails, ok := s.D.GetOkExists("data_transfer_medium_details"); ok {
			if tmpList := dataTransferMediumDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_transfer_medium_details", 0)
				tmp, err := s.mapToCreateMySqlDataTransferMediumDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataTransferMediumDetails = tmp
			}
		}
		if excludeObjects, ok := s.D.GetOkExists("exclude_objects"); ok {
			set := excludeObjects.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database_migration.MySqlDatabaseObject, len(interfaces))
			for i := range interfaces {
				stateDataIndex := excludeObjectsHashCodeForSets(interfaces[i])
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "exclude_objects", stateDataIndex)
				converted, err := s.mapToMySqlDatabaseObject(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("exclude_objects") {
				details.ExcludeObjects = tmp
			}
		}
		if ggsDetails, ok := s.D.GetOkExists("ggs_details"); ok {
			if tmpList := ggsDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ggs_details", 0)
				tmp, err := s.mapToCreateMySqlGgsDeploymentDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.GgsDetails = &tmp
			}
		}
		if hubDetails, ok := s.D.GetOkExists("hub_details"); ok {
			if tmpList := hubDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "hub_details", 0)
				tmp, err := s.mapToCreateGoldenGateHubDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.HubDetails = &tmp
			}
		}
		if includeObjects, ok := s.D.GetOkExists("include_objects"); ok {
			interfaces := includeObjects.([]interface{})
			tmp := make([]oci_database_migration.MySqlDatabaseObject, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "include_objects", stateDataIndex)
				converted, err := s.mapToMySqlDatabaseObject(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("include_objects") {
				details.IncludeObjects = tmp
			}
		}
		if initialLoadSettings, ok := s.D.GetOkExists("initial_load_settings"); ok {
			if tmpList := initialLoadSettings.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "initial_load_settings", 0)
				tmp, err := s.mapToCreateMySqlInitialLoadSettings(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.InitialLoadSettings = &tmp
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
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if sourceDatabaseConnectionId, ok := s.D.GetOkExists("source_database_connection_id"); ok {
			tmp := sourceDatabaseConnectionId.(string)
			details.SourceDatabaseConnectionId = &tmp
		}
		if targetDatabaseConnectionId, ok := s.D.GetOkExists("target_database_connection_id"); ok {
			tmp := targetDatabaseConnectionId.(string)
			details.TargetDatabaseConnectionId = &tmp
		}
		if type_, ok := s.D.GetOkExists("type"); ok {
			details.Type = oci_database_migration.MigrationTypesEnum(type_.(string))
		}
		request.CreateMigrationDetails = details
	case strings.ToLower("ORACLE"):
		details := oci_database_migration.CreateOracleMigrationDetails{}
		if advancedParameters, ok := s.D.GetOkExists("advanced_parameters"); ok {
			interfaces := advancedParameters.([]interface{})
			tmp := make([]oci_database_migration.MigrationParameterDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "advanced_parameters", stateDataIndex)
				converted, err := s.mapToMigrationParameterDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("advanced_parameters") {
				details.AdvancedParameters = tmp
			}
		}
		if advisorSettings, ok := s.D.GetOkExists("advisor_settings"); ok {
			if tmpList := advisorSettings.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "advisor_settings", 0)
				tmp, err := s.mapToCreateOracleAdvisorSettings(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AdvisorSettings = &tmp
			}
		}
		if bulkIncludeExcludeData, ok := s.D.GetOkExists("bulk_include_exclude_data"); ok {
			tmp := bulkIncludeExcludeData.(string)
			details.BulkIncludeExcludeData = &tmp
		}
		if dataTransferMediumDetails, ok := s.D.GetOkExists("data_transfer_medium_details"); ok {
			if tmpList := dataTransferMediumDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_transfer_medium_details", 0)
				tmp, err := s.mapToCreateOracleDataTransferMediumDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataTransferMediumDetails = tmp
			}
		}
		if excludeObjects, ok := s.D.GetOkExists("exclude_objects"); ok {
			set := excludeObjects.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_database_migration.OracleDatabaseObject, len(interfaces))
			for i := range interfaces {
				stateDataIndex := excludeObjectsHashCodeForSets(interfaces[i])
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "exclude_objects", stateDataIndex)
				converted, err := s.mapToOracleDatabaseObject(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("exclude_objects") {
				details.ExcludeObjects = tmp
			}
		}
		if ggsDetails, ok := s.D.GetOkExists("ggs_details"); ok {
			if tmpList := ggsDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ggs_details", 0)
				tmp, err := s.mapToCreateOracleGgsDeploymentDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.GgsDetails = &tmp
			}
		}
		if hubDetails, ok := s.D.GetOkExists("hub_details"); ok {
			if tmpList := hubDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "hub_details", 0)
				tmp, err := s.mapToCreateGoldenGateHubDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.HubDetails = &tmp
			}
		}
		if includeObjects, ok := s.D.GetOkExists("include_objects"); ok {
			interfaces := includeObjects.([]interface{})
			tmp := make([]oci_database_migration.OracleDatabaseObject, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "include_objects", stateDataIndex)
				converted, err := s.mapToOracleDatabaseObject(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("include_objects") {
				details.IncludeObjects = tmp
			}
		}
		if initialLoadSettings, ok := s.D.GetOkExists("initial_load_settings"); ok {
			if tmpList := initialLoadSettings.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "initial_load_settings", 0)
				tmp, err := s.mapToCreateOracleInitialLoadSettings(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.InitialLoadSettings = &tmp
			}
		}
		if sourceContainerDatabaseConnectionId, ok := s.D.GetOkExists("source_container_database_connection_id"); ok {
			tmp := sourceContainerDatabaseConnectionId.(string)
			details.SourceContainerDatabaseConnectionId = &tmp
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
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if sourceDatabaseConnectionId, ok := s.D.GetOkExists("source_database_connection_id"); ok {
			tmp := sourceDatabaseConnectionId.(string)
			details.SourceDatabaseConnectionId = &tmp
		}
		if targetDatabaseConnectionId, ok := s.D.GetOkExists("target_database_connection_id"); ok {
			tmp := targetDatabaseConnectionId.(string)
			details.TargetDatabaseConnectionId = &tmp
		}
		if type_, ok := s.D.GetOkExists("type"); ok {
			details.Type = oci_database_migration.MigrationTypesEnum(type_.(string))
		}
		request.CreateMigrationDetails = details
	default:
		return fmt.Errorf("unknown database_combination '%v' was specified", databaseCombination)
	}
	return nil
}

func (s *DatabaseMigrationMigrationResourceCrud) populateTopLevelPolymorphicUpdateMigrationRequest(request *oci_database_migration.UpdateMigrationRequest) error {
	//discriminator
	databaseCombinationRaw, ok := s.D.GetOkExists("database_combination")
	var databaseCombination string
	if ok {
		databaseCombination = databaseCombinationRaw.(string)
	} else {
		databaseCombination = "" // default value
	}
	switch strings.ToLower(databaseCombination) {
	case strings.ToLower("MYSQL"):
		details := oci_database_migration.UpdateMySqlMigrationDetails{}
		if advisorSettings, ok := s.D.GetOkExists("advisor_settings"); ok {
			if tmpList := advisorSettings.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "advisor_settings", 0)
				tmp, err := s.mapToUpdateMySqlAdvisorSettings(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AdvisorSettings = &tmp
			}
		}
		if dataTransferMediumDetails, ok := s.D.GetOkExists("data_transfer_medium_details"); ok {
			if tmpList := dataTransferMediumDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_transfer_medium_details", 0)
				tmp, err := s.mapToUpdateMySqlDataTransferMediumDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataTransferMediumDetails = tmp
			}
		}
		if ggsDetails, ok := s.D.GetOkExists("ggs_details"); ok {
			if tmpList := ggsDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ggs_details", 0)
				tmp, err := s.mapToUpdateMySqlGgsDeploymentDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.GgsDetails = &tmp
			}
		}
		if hubDetails, ok := s.D.GetOkExists("hub_details"); ok {
			if tmpList := hubDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "hub_details", 0)
				tmp, err := s.mapToUpdateGoldenGateHubDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.HubDetails = &tmp
			}
		}
		if initialLoadSettings, ok := s.D.GetOkExists("initial_load_settings"); ok {
			if tmpList := initialLoadSettings.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "initial_load_settings", 0)
				tmp, err := s.mapToUpdateMySqlInitialLoadSettings(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.InitialLoadSettings = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.MigrationId = &tmp
		if sourceDatabaseConnectionId, ok := s.D.GetOkExists("source_database_connection_id"); ok {
			tmp := sourceDatabaseConnectionId.(string)
			details.SourceDatabaseConnectionId = &tmp
		}
		if targetDatabaseConnectionId, ok := s.D.GetOkExists("target_database_connection_id"); ok {
			tmp := targetDatabaseConnectionId.(string)
			details.TargetDatabaseConnectionId = &tmp
		}
		if type_, ok := s.D.GetOkExists("type"); ok {
			details.Type = oci_database_migration.MigrationTypesEnum(type_.(string))
		}
		request.UpdateMigrationDetails = details
	case strings.ToLower("ORACLE"):
		details := oci_database_migration.UpdateOracleMigrationDetails{}
		if advancedParameters, ok := s.D.GetOkExists("advanced_parameters"); ok {
			interfaces := advancedParameters.([]interface{})
			tmp := make([]oci_database_migration.MigrationParameterDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "advanced_parameters", stateDataIndex)
				converted, err := s.mapToMigrationParameterDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("advanced_parameters") {
				details.AdvancedParameters = tmp
			}
		}
		if advisorSettings, ok := s.D.GetOkExists("advisor_settings"); ok {
			if tmpList := advisorSettings.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "advisor_settings", 0)
				tmp, err := s.mapToUpdateOracleAdvisorSettings(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AdvisorSettings = &tmp
			}
		}
		if dataTransferMediumDetails, ok := s.D.GetOkExists("data_transfer_medium_details"); ok {
			if tmpList := dataTransferMediumDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_transfer_medium_details", 0)
				tmp, err := s.mapToUpdateOracleDataTransferMediumDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataTransferMediumDetails = tmp
			}
		}
		if ggsDetails, ok := s.D.GetOkExists("ggs_details"); ok {
			if tmpList := ggsDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ggs_details", 0)
				tmp, err := s.mapToUpdateOracleGgsDeploymentDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.GgsDetails = &tmp
			}
		}
		if hubDetails, ok := s.D.GetOkExists("hub_details"); ok {
			if tmpList := hubDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "hub_details", 0)
				tmp, err := s.mapToUpdateGoldenGateHubDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.HubDetails = &tmp
			}
		}
		if initialLoadSettings, ok := s.D.GetOkExists("initial_load_settings"); ok {
			if tmpList := initialLoadSettings.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "initial_load_settings", 0)
				tmp, err := s.mapToUpdateOracleInitialLoadSettings(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.InitialLoadSettings = &tmp
			}
		}
		if sourceContainerDatabaseConnectionId, ok := s.D.GetOkExists("source_container_database_connection_id"); ok {
			tmp := sourceContainerDatabaseConnectionId.(string)
			details.SourceContainerDatabaseConnectionId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.MigrationId = &tmp
		if sourceDatabaseConnectionId, ok := s.D.GetOkExists("source_database_connection_id"); ok {
			tmp := sourceDatabaseConnectionId.(string)
			details.SourceDatabaseConnectionId = &tmp
		}
		if targetDatabaseConnectionId, ok := s.D.GetOkExists("target_database_connection_id"); ok {
			tmp := targetDatabaseConnectionId.(string)
			details.TargetDatabaseConnectionId = &tmp
		}
		if type_, ok := s.D.GetOkExists("type"); ok {
			details.Type = oci_database_migration.MigrationTypesEnum(type_.(string))
		}
		request.UpdateMigrationDetails = details
	default:
		return fmt.Errorf("unknown database_combination '%v' was specified", databaseCombination)
	}
	return nil
}

func (s *DatabaseMigrationMigrationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database_migration.ChangeMigrationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MigrationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	_, err := s.Client.ChangeMigrationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
