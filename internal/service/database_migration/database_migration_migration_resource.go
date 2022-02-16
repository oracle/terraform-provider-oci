// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v58/databasemigration"
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
			"agent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

						// Optional
						"database_link_details": {
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
									"wallet_bucket": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"bucket": {
													Type:     schema.TypeString,
													Required: true,
												},
												"namespace": {
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
						"object_storage_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"bucket": {
										Type:     schema.TypeString,
										Required: true,
									},
									"namespace": {
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
			"datapump_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
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
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"path": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
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
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"path": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"job_mode": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metadata_remaps": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      metadataRemapsHashCodeForSets,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"new_value": {
										Type:     schema.TypeString,
										Required: true,
									},
									"old_value": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
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
			"dump_transfer_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

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

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"exclude_objects": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      excludeObjectsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"object": {
							Type:     schema.TypeString,
							Required: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"golden_gate_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"hub": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
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
									"source_db_admin_credentials": {
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
									"source_microservices_deployment_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"target_db_admin_credentials": {
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
									"target_microservices_deployment_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"url": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"compute_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"source_container_db_admin_credentials": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
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

									// Computed
								},
							},
						},

						// Optional
						"settings": {
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
												"map_parallelism": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"max_apply_parallelism": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"min_apply_parallelism": {
													Type:     schema.TypeInt,
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

						// Computed
					},
				},
			},
			"include_objects": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"object": {
							Type:     schema.TypeString,
							Required: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"vault_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"key_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"vault_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"credentials_secret_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
	if tType, ok := m["type"]; ok && tType != "" {
		buf.WriteString(fmt.Sprintf("%v-", tType))
	}

	return hashcode.String(buf.String())
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
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.Migration

	DisableNotFoundRetries bool
}

func (s *DatabaseMigrationMigrationResourceCrud) ID() string {
	return *s.Res.Id
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

	if advisorSettings, ok := s.D.GetOkExists("advisor_settings"); ok {
		if tmpList := advisorSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "advisor_settings", 0)
			tmp, err := s.mapToCreateAdvisorSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AdvisorSettings = &tmp
		}
	}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataTransferMediumDetails, ok := s.D.GetOkExists("data_transfer_medium_details"); ok {
		if tmpList := dataTransferMediumDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_transfer_medium_details", 0)
			tmp, err := s.mapToCreateDataTransferMediumDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataTransferMediumDetails = &tmp
		}
	}

	if datapumpSettings, ok := s.D.GetOkExists("datapump_settings"); ok {
		if tmpList := datapumpSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "datapump_settings", 0)
			tmp, err := s.mapToCreateDataPumpSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatapumpSettings = &tmp
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

	if dumpTransferDetails, ok := s.D.GetOkExists("dump_transfer_details"); ok {
		if tmpList := dumpTransferDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dump_transfer_details", 0)
			tmp, err := s.mapToCreateDumpTransferDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DumpTransferDetails = &tmp
		}
	}

	if excludeObjects, ok := s.D.GetOkExists("exclude_objects"); ok {
		set := excludeObjects.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database_migration.DatabaseObject, len(interfaces))
		for i := range interfaces {
			stateDataIndex := excludeObjectsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "exclude_objects", stateDataIndex)
			converted, err := s.mapToDatabaseObject(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("exclude_objects") {
			request.ExcludeObjects = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if goldenGateDetails, ok := s.D.GetOkExists("golden_gate_details"); ok {
		if tmpList := goldenGateDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "golden_gate_details", 0)
			tmp, err := s.mapToCreateGoldenGateDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.GoldenGateDetails = &tmp
		}
	}

	if includeObjects, ok := s.D.GetOkExists("include_objects"); ok {
		interfaces := includeObjects.([]interface{})
		tmp := make([]oci_database_migration.DatabaseObject, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "include_objects", stateDataIndex)
			converted, err := s.mapToDatabaseObject(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("include_objects") {
			request.IncludeObjects = tmp
		}
	}

	if sourceContainerDatabaseConnectionId, ok := s.D.GetOkExists("source_container_database_connection_id"); ok {
		tmp := sourceContainerDatabaseConnectionId.(string)
		request.SourceContainerDatabaseConnectionId = &tmp
	}

	if sourceDatabaseConnectionId, ok := s.D.GetOkExists("source_database_connection_id"); ok {
		tmp := sourceDatabaseConnectionId.(string)
		request.SourceDatabaseConnectionId = &tmp
	}

	if targetDatabaseConnectionId, ok := s.D.GetOkExists("target_database_connection_id"); ok {
		tmp := targetDatabaseConnectionId.(string)
		request.TargetDatabaseConnectionId = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_database_migration.MigrationTypesEnum(type_.(string))
	}

	if vaultDetails, ok := s.D.GetOkExists("vault_details"); ok {
		if tmpList := vaultDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vault_details", 0)
			tmp, err := s.mapToCreateVaultDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.VaultDetails = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.CreateMigration(context.Background(), request)

	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
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

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_database_migration.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "migration")

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

	if advisorSettings, ok := s.D.GetOkExists("advisor_settings"); ok {
		if tmpList := advisorSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "advisor_settings", 0)
			tmp, err := s.mapToUpdateAdvisorSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AdvisorSettings = &tmp
		}
	}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if dataTransferMediumDetails, ok := s.D.GetOkExists("data_transfer_medium_details"); ok && s.D.HasChange("data_transfer_medium_details") {
		if tmpList := dataTransferMediumDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_transfer_medium_details", 0)
			tmp, err := s.mapToUpdateDataTransferMediumDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataTransferMediumDetails = &tmp
		}
	}

	if datapumpSettings, ok := s.D.GetOkExists("datapump_settings"); ok && s.D.HasChange("datapump_settings") {
		if tmpList := datapumpSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "datapump_settings", 0)
			tmp, err := s.mapToUpdateDataPumpSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatapumpSettings = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
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

	if dumpTransferDetails, ok := s.D.GetOkExists("dump_transfer_details"); ok {
		if tmpList := dumpTransferDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dump_transfer_details", 0)
			tmp, err := s.mapToUpdateDumpTransferDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DumpTransferDetails = &tmp
		}
	}

	if excludeObjects, ok := s.D.GetOkExists("exclude_objects"); ok {
		set := excludeObjects.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database_migration.DatabaseObject, len(interfaces))
		for i := range interfaces {
			stateDataIndex := excludeObjectsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "exclude_objects", stateDataIndex)
			converted, err := s.mapToDatabaseObject(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("exclude_objects") {
			request.ExcludeObjects = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if goldenGateDetails, ok := s.D.GetOkExists("golden_gate_details"); ok && s.D.HasChange("golden_gate_details") {
		if tmpList := goldenGateDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "golden_gate_details", 0)
			tmp, err := s.mapToUpdateGoldenGateDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.GoldenGateDetails = &tmp
		}
	}

	if includeObjects, ok := s.D.GetOkExists("include_objects"); ok {
		interfaces := includeObjects.([]interface{})
		tmp := make([]oci_database_migration.DatabaseObject, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "include_objects", stateDataIndex)
			converted, err := s.mapToDatabaseObject(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("include_objects") {
			request.IncludeObjects = tmp
		}
	}

	tmp := s.D.Id()
	request.MigrationId = &tmp

	if sourceContainerDatabaseConnectionId, ok := s.D.GetOkExists("source_container_database_connection_id"); ok && s.D.HasChange("source_container_database_connection_id") {
		tmp := sourceContainerDatabaseConnectionId.(string)
		request.SourceContainerDatabaseConnectionId = &tmp
	}

	if sourceDatabaseConnectionId, ok := s.D.GetOkExists("source_database_connection_id"); ok && s.D.HasChange("source_database_connection_id") {
		tmp := sourceDatabaseConnectionId.(string)
		request.SourceDatabaseConnectionId = &tmp
	}

	if targetDatabaseConnectionId, ok := s.D.GetOkExists("target_database_connection_id"); ok && s.D.HasChange("target_database_connection_id") {
		tmp := targetDatabaseConnectionId.(string)
		request.TargetDatabaseConnectionId = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok && s.D.HasChange("type") {
		request.Type = oci_database_migration.MigrationTypesEnum(type_.(string))
	}

	if vaultDetails, ok := s.D.GetOkExists("vault_details"); ok && s.D.HasChange("vault_details") {
		if tmpList := vaultDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vault_details", 0)
			tmp, err := s.mapToUpdateVaultDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.VaultDetails = &tmp
		}
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
	if s.Res.AdvisorSettings != nil {
		s.D.Set("advisor_settings", []interface{}{AdvisorSettingsToMap(s.Res.AdvisorSettings)})
	} else {
		s.D.Set("advisor_settings", nil)
	}

	if s.Res.AgentId != nil {
		s.D.Set("agent_id", *s.Res.AgentId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CredentialsSecretId != nil {
		s.D.Set("credentials_secret_id", *s.Res.CredentialsSecretId)
	}

	if s.Res.DataTransferMediumDetails != nil {
		s.D.Set("data_transfer_medium_details", []interface{}{DataTransferMediumDetailsToMap(s.Res.DataTransferMediumDetails)})
	} else {
		s.D.Set("data_transfer_medium_details", nil)
	}

	if s.Res.DatapumpSettings != nil {
		s.D.Set("datapump_settings", []interface{}{DataPumpSettingsToMap(s.Res.DatapumpSettings)})
	} else {
		s.D.Set("datapump_settings", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DumpTransferDetails != nil {
		s.D.Set("dump_transfer_details", []interface{}{DumpTransferDetailsToMap(s.Res.DumpTransferDetails)})
	} else {
		s.D.Set("dump_transfer_details", nil)
	}

	excludeObjects := []interface{}{}
	for _, item := range s.Res.ExcludeObjects {
		excludeObjects = append(excludeObjects, DatabaseObjectToMap(item))
	}
	s.D.Set("exclude_objects", excludeObjects)

	if s.Res.ExecutingJobId != nil {
		s.D.Set("executing_job_id", *s.Res.ExecutingJobId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GoldenGateDetails != nil {
		//s.D.Set("golden_gate_details", []interface{}{GoldenGateDetailsToMap(s.Res.GoldenGateDetails)})
		s.D.Set("golden_gate_details", []interface{}{GoldenGateDetailsToMapPass(s.Res.GoldenGateDetails, s.D)})

	} else {
		s.D.Set("golden_gate_details", nil)
	}

	includeObjects := []interface{}{}
	for _, item := range s.Res.IncludeObjects {
		includeObjects = append(includeObjects, DatabaseObjectToMap(item))
	}
	s.D.Set("include_objects", includeObjects)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.SourceContainerDatabaseConnectionId != nil {
		s.D.Set("source_container_database_connection_id", *s.Res.SourceContainerDatabaseConnectionId)
	}

	if s.Res.SourceDatabaseConnectionId != nil {
		s.D.Set("source_database_connection_id", *s.Res.SourceDatabaseConnectionId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetDatabaseConnectionId != nil {
		s.D.Set("target_database_connection_id", *s.Res.TargetDatabaseConnectionId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastMigration != nil {
		s.D.Set("time_last_migration", s.Res.TimeLastMigration.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.VaultDetails != nil {
		s.D.Set("vault_details", []interface{}{VaultDetailsToMap(s.Res.VaultDetails)})
	} else {
		s.D.Set("vault_details", nil)
	}

	s.D.Set("wait_after", s.Res.WaitAfter)

	return nil
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

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateAdvisorSettings(fieldKeyFormat string) (oci_database_migration.UpdateAdvisorSettings, error) {
	result := oci_database_migration.UpdateAdvisorSettings{}

	if isSkipAdvisor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_advisor")); ok {
		tmp := isSkipAdvisor.(bool)
		result.IsSkipAdvisor = &tmp
	}

	if isIgnoreErrors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ignore_errors")); ok {
		tmp := isIgnoreErrors.(bool)
		result.IsIgnoreErrors = &tmp
	}

	return result, nil
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateDumpTransferDetails(fieldKeyFormat string) (oci_database_migration.UpdateDumpTransferDetails, error) {
	result := oci_database_migration.UpdateDumpTransferDetails{}

	if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
			tmp, err := s.mapToUpdateHostDumpTransferDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source, encountered error: %v", err)
			}
			result.Source = tmp
		}
	}

	if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
		if tmpList := target.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "target"), 0)
			tmp, err := s.mapToUpdateHostDumpTransferDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert target, encountered error: %v", err)
			}
			result.Target = tmp
		}
	}

	return result, nil
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

/*func AdminCredentialsToMap(obj *oci_database_migration.AdminCredentials) map[string]interface{} {
	result := map[string]interface{}{}


		if obj.Password != nil {
			result["password"] = string(*obj.Password)
		}


	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}

	return result
}*/

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateAdvisorSettings(fieldKeyFormat string) (oci_database_migration.CreateAdvisorSettings, error) {
	result := oci_database_migration.CreateAdvisorSettings{}

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

func AdvisorSettingsToMap(obj *oci_database_migration.AdvisorSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsIgnoreErrors != nil {
		result["is_ignore_errors"] = bool(*obj.IsIgnoreErrors)
	}

	if obj.IsSkipAdvisor != nil {
		result["is_skip_advisor"] = bool(*obj.IsSkipAdvisor)
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
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude_parameters"), stateDataIndex)
			converted, err := s.mapToDataPumpExcludeParameters(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclude_parameters")) {
			result.ExcludeParameters = tmp
		}
	}

	if exportParallelismDegree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_parallelism_degree")); ok {
		tmp := exportParallelismDegree.(int)
		result.ExportParallelismDegree = &tmp
	}

	if importParallelismDegree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "import_parallelism_degree")); ok {
		tmp := importParallelismDegree.(int)
		result.ImportParallelismDegree = &tmp
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

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateDataPumpParameters(fieldKeyFormat string) (oci_database_migration.UpdateDataPumpParameters, error) {
	result := oci_database_migration.UpdateDataPumpParameters{}

	if estimate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "estimate")); ok {
		result.Estimate = oci_database_migration.DataPumpEstimateEnum(estimate.(string))
	}

	if excludeParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_parameters")); ok {
		interfaces := excludeParameters.([]interface{})
		tmp := make([]oci_database_migration.DataPumpExcludeParametersEnum, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude_parameters"), stateDataIndex)
			converted, err := s.mapToDataPumpExcludeParameters(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclude_parameters")) {
			result.ExcludeParameters = tmp
		}
	}

	if exportParallelismDegree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_parallelism_degree")); ok {
		tmp := exportParallelismDegree.(int)
		result.ExportParallelismDegree = &tmp
	}

	if importParallelismDegree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "import_parallelism_degree")); ok {
		tmp := importParallelismDegree.(int)
		result.ImportParallelismDegree = &tmp
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

	excludeParameters := []interface{}{}
	for _, item := range obj.ExcludeParameters {
		excludeParameters = append(excludeParameters, DataPumpExcludeParametersToMap(item))
	}
	result["exclude_parameters"] = excludeParameters

	if obj.ExportParallelismDegree != nil {
		result["export_parallelism_degree"] = int(*obj.ExportParallelismDegree)
	}

	if obj.ImportParallelismDegree != nil {
		result["import_parallelism_degree"] = int(*obj.ImportParallelismDegree)
	}

	if obj.IsCluster != nil {
		result["is_cluster"] = bool(*obj.IsCluster)
	}

	result["table_exists_action"] = string(obj.TableExistsAction)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateDataPumpSettings(fieldKeyFormat string) (oci_database_migration.CreateDataPumpSettings, error) {
	result := oci_database_migration.CreateDataPumpSettings{}

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
		result.JobMode = oci_database_migration.DataPumpJobModeEnum(jobMode.(string))
	}

	if metadataRemaps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata_remaps")); ok {
		set := metadataRemaps.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database_migration.MetadataRemap, len(interfaces))
		for i := range interfaces {
			stateDataIndex := metadataRemapsHashCodeForSets(interfaces[i])
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

	return result, nil
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateDataPumpSettings(fieldKeyFormat string) (oci_database_migration.UpdateDataPumpSettings, error) {
	result := oci_database_migration.UpdateDataPumpSettings{}

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
		result.JobMode = oci_database_migration.DataPumpJobModeEnum(jobMode.(string))
	}

	if metadataRemaps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata_remaps")); ok {
		set := metadataRemaps.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database_migration.MetadataRemap, len(interfaces))
		for i := range interfaces {
			stateDataIndex := metadataRemapsHashCodeForSets(interfaces[i])
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
	return result, nil
}

func DataPumpSettingsToMap(obj *oci_database_migration.DataPumpSettings) map[string]interface{} {
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
	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateDataTransferMediumDetails(fieldKeyFormat string) (oci_database_migration.CreateDataTransferMediumDetails, error) {
	result := oci_database_migration.CreateDataTransferMediumDetails{}

	if databaseLinkDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_link_details")); ok {
		if tmpList := databaseLinkDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database_link_details"), 0)
			tmp, err := s.mapToCreateDatabaseLinkDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert database_link_details, encountered error: %v", err)
			}
			result.DatabaseLinkDetails = &tmp
		}
	}

	if objectStorageDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_details")); ok {
		if tmpList := objectStorageDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_details"), 0)
			tmp, err := s.mapToCreateObjectStoreBucket(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert object_storage_details, encountered error: %v", err)
			}
			result.ObjectStorageDetails = &tmp
		}
	}

	return result, nil
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateDataTransferMediumDetails(fieldKeyFormat string) (oci_database_migration.UpdateDataTransferMediumDetails, error) {
	result := oci_database_migration.UpdateDataTransferMediumDetails{}

	if databaseLinkDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_link_details")); ok {
		if tmpList := databaseLinkDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database_link_details"), 0)
			tmp, err := s.mapToUpdateDatabaseLinkDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert database_link_details, encountered error: %v", err)
			}
			result.DatabaseLinkDetails = &tmp
		}
	}

	if objectStorageDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_details")); ok {
		if tmpList := objectStorageDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_details"), 0)
			tmp, err := s.mapToUpdateObjectStoreBucket(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert object_storage_details, encountered error: %v", err)
			}
			result.ObjectStorageDetails = &tmp
		}
	}

	return result, nil
}

func DataTransferMediumDetailsToMap(obj *oci_database_migration.DataTransferMediumDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DatabaseLinkDetails != nil {
		result["database_link_details"] = []interface{}{DatabaseLinkDetailsToMap(obj.DatabaseLinkDetails)}
	}

	if obj.ObjectStorageDetails != nil {
		result["object_storage_details"] = []interface{}{ObjectStoreBucketToMap(obj.ObjectStorageDetails)}
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateDatabaseLinkDetails(fieldKeyFormat string) (oci_database_migration.CreateDatabaseLinkDetails, error) {
	result := oci_database_migration.CreateDatabaseLinkDetails{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if walletBucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wallet_bucket")); ok {
		if tmpList := walletBucket.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "wallet_bucket"), 0)
			tmp, err := s.mapToCreateObjectStoreBucket(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert wallet_bucket, encountered error: %v", err)
			}
			result.WalletBucket = &tmp
		}
	}

	return result, nil
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateDatabaseLinkDetails(fieldKeyFormat string) (oci_database_migration.UpdateDatabaseLinkDetails, error) {
	result := oci_database_migration.UpdateDatabaseLinkDetails{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func DatabaseLinkDetailsToMap(obj *oci_database_migration.DatabaseLinkDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.WalletBucket != nil {
		result["wallet_bucket"] = []interface{}{ObjectStoreBucketToMap(obj.WalletBucket)}
	}

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

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateDirectoryObject(fieldKeyFormat string) (oci_database_migration.UpdateDirectoryObject, error) {
	result := oci_database_migration.UpdateDirectoryObject{}

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

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateDumpTransferDetails(fieldKeyFormat string) (oci_database_migration.CreateDumpTransferDetails, error) {
	result := oci_database_migration.CreateDumpTransferDetails{}

	if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
			tmp, err := s.mapToCreateHostDumpTransferDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source, encountered error: %v", err)
			}
			result.Source = tmp
		}
	}

	if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
		if tmpList := target.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "target"), 0)
			tmp, err := s.mapToCreateHostDumpTransferDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert target, encountered error: %v", err)
			}
			result.Target = tmp
		}
	}

	return result, nil
}

func DumpTransferDetailsToMap(obj *oci_database_migration.DumpTransferDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := HostDumpTransferDetailsToMap(&obj.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		result["source"] = sourceArray
	}

	if obj.Target != nil {
		targetArray := []interface{}{}
		if targetMap := HostDumpTransferDetailsToMap(&obj.Target); targetMap != nil {
			targetArray = append(targetArray, targetMap)
		}
		result["target"] = targetArray
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateExtract(fieldKeyFormat string) (oci_database_migration.CreateExtract, error) {
	result := oci_database_migration.CreateExtract{}

	if longTransDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "long_trans_duration")); ok {
		tmp := longTransDuration.(int)
		result.LongTransDuration = &tmp
	}

	if performanceProfile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "performance_profile")); ok {
		result.PerformanceProfile = oci_database_migration.ExtractPerformanceProfileEnum(performanceProfile.(string))
	}

	return result, nil
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateExtract(fieldKeyFormat string) (oci_database_migration.UpdateExtract, error) {
	result := oci_database_migration.UpdateExtract{}

	if longTransDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "long_trans_duration")); ok {
		tmp := longTransDuration.(int)
		result.LongTransDuration = &tmp
	}

	if performanceProfile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "performance_profile")); ok {
		result.PerformanceProfile = oci_database_migration.ExtractPerformanceProfileEnum(performanceProfile.(string))
	}

	return result, nil
}

func ExtractToMap(obj *oci_database_migration.Extract) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LongTransDuration != nil {
		result["long_trans_duration"] = int(*obj.LongTransDuration)
	}

	result["performance_profile"] = string(obj.PerformanceProfile)

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateGoldenGateDetails(fieldKeyFormat string) (oci_database_migration.CreateGoldenGateDetails, error) {
	result := oci_database_migration.CreateGoldenGateDetails{}

	if hub, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hub")); ok {
		if tmpList := hub.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "hub"), 0)
			tmp, err := s.mapToCreateGoldenGateHub(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert hub, encountered error: %v", err)
			}
			result.Hub = &tmp
		}
	}

	if settings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "settings")); ok {
		if tmpList := settings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "settings"), 0)
			tmp, err := s.mapToCreateGoldenGateSettings(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert settings, encountered error: %v", err)
			}
			result.Settings = &tmp
		}
	}

	return result, nil
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateGoldenGateDetails(fieldKeyFormat string) (oci_database_migration.UpdateGoldenGateDetails, error) {
	result := oci_database_migration.UpdateGoldenGateDetails{}

	if hub, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hub")); ok {
		if tmpList := hub.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "hub"), 0)
			tmp, err := s.mapToUpdateGoldenGateHub(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert hub, encountered error: %v", err)
			}
			result.Hub = &tmp
		}
	}

	if settings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "settings")); ok {
		if tmpList := settings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "settings"), 0)
			tmp, err := s.mapToUpdateGoldenGateSettings(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert settings, encountered error: %v", err)
			}
			result.Settings = &tmp
		}
	}

	return result, nil
}
func GoldenGateDetailsToMap(obj *oci_database_migration.GoldenGateDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hub != nil {
		result["hub"] = []interface{}{GoldenGateHubToMap(obj.Hub)}
	}

	if obj.Settings != nil {
		result["settings"] = []interface{}{GoldenGateSettingsToMap(obj.Settings)}
	}

	return result
}

func GoldenGateDetailsToMapPass(obj *oci_database_migration.GoldenGateDetails, resourceData *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}
	if obj.Hub != nil {
		result["hub"] = []interface{}{GoldenGateHubToMapPass(obj.Hub, resourceData)}
	}

	if obj.Settings != nil {
		result["settings"] = []interface{}{GoldenGateSettingsToMap(obj.Settings)}
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateGoldenGateHub(fieldKeyFormat string) (oci_database_migration.CreateGoldenGateHub, error) {
	result := oci_database_migration.CreateGoldenGateHub{}

	if computeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_id")); ok {
		tmp := computeId.(string)
		result.ComputeId = &tmp
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

	if sourceContainerDbAdminCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_container_db_admin_credentials")); ok {
		if tmpList := sourceContainerDbAdminCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_container_db_admin_credentials"), 0)
			tmp, err := s.mapToCreateAdminCredentials(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_container_db_admin_credentials, encountered error: %v", err)
			}
			result.SourceContainerDbAdminCredentials = &tmp
		}
	}

	if sourceDbAdminCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_db_admin_credentials")); ok {
		if tmpList := sourceDbAdminCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_db_admin_credentials"), 0)
			tmp, err := s.mapToCreateAdminCredentials(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_db_admin_credentials, encountered error: %v", err)
			}
			result.SourceDbAdminCredentials = &tmp
		}
	}

	if sourceMicroservicesDeploymentName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_microservices_deployment_name")); ok {
		tmp := sourceMicroservicesDeploymentName.(string)
		result.SourceMicroservicesDeploymentName = &tmp
	}

	if targetDbAdminCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_db_admin_credentials")); ok {
		if tmpList := targetDbAdminCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "target_db_admin_credentials"), 0)
			tmp, err := s.mapToCreateAdminCredentials(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert target_db_admin_credentials, encountered error: %v", err)
			}
			result.TargetDbAdminCredentials = &tmp
		}
	}

	if targetMicroservicesDeploymentName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_microservices_deployment_name")); ok {
		tmp := targetMicroservicesDeploymentName.(string)
		result.TargetMicroservicesDeploymentName = &tmp
	}

	if url, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url")); ok {
		tmp := url.(string)
		result.Url = &tmp
	}

	return result, nil
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateGoldenGateHub(fieldKeyFormat string) (oci_database_migration.UpdateGoldenGateHub, error) {
	result := oci_database_migration.UpdateGoldenGateHub{}

	if computeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_id")); ok {
		tmp := computeId.(string)
		result.ComputeId = &tmp
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

	if sourceContainerDbAdminCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_container_db_admin_credentials")); ok {
		if tmpList := sourceContainerDbAdminCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_container_db_admin_credentials"), 0)
			tmp, err := s.mapToUpdateAdminCredentials(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_container_db_admin_credentials, encountered error: %v", err)
			}
			result.SourceContainerDbAdminCredentials = &tmp
		}
	}

	if sourceDbAdminCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_db_admin_credentials")); ok {
		if tmpList := sourceDbAdminCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_db_admin_credentials"), 0)
			tmp, err := s.mapToUpdateAdminCredentials(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_db_admin_credentials, encountered error: %v", err)
			}
			result.SourceDbAdminCredentials = &tmp
		}
	}

	if sourceMicroservicesDeploymentName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_microservices_deployment_name")); ok {
		tmp := sourceMicroservicesDeploymentName.(string)
		result.SourceMicroservicesDeploymentName = &tmp
	}

	if targetDbAdminCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_db_admin_credentials")); ok {
		if tmpList := targetDbAdminCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "target_db_admin_credentials"), 0)
			tmp, err := s.mapToUpdateAdminCredentials(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert target_db_admin_credentials, encountered error: %v", err)
			}
			result.TargetDbAdminCredentials = &tmp
		}
	}

	if targetMicroservicesDeploymentName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_microservices_deployment_name")); ok {
		tmp := targetMicroservicesDeploymentName.(string)
		result.TargetMicroservicesDeploymentName = &tmp
	}

	if url, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url")); ok {
		tmp := url.(string)
		result.Url = &tmp
	}

	return result, nil
}

func GoldenGateHubToMap(obj *oci_database_migration.GoldenGateHub) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComputeId != nil {
		result["compute_id"] = string(*obj.ComputeId)
	}

	if obj.RestAdminCredentials != nil {
		result["rest_admin_credentials"] = []interface{}{AdminCredentialsToMap(obj.RestAdminCredentials)}
		//result["rest_admin_credentials"] = []interface{}{AdminCredentialsToMapPassword(obj.RestAdminCredentials, obj.)}

	}

	if obj.SourceContainerDbAdminCredentials != nil {
		result["source_container_db_admin_credentials"] = []interface{}{AdminCredentialsToMap(obj.SourceContainerDbAdminCredentials)}
	}

	if obj.SourceDbAdminCredentials != nil {
		result["source_db_admin_credentials"] = []interface{}{AdminCredentialsToMap(obj.SourceDbAdminCredentials)}
	}

	if obj.SourceMicroservicesDeploymentName != nil {
		result["source_microservices_deployment_name"] = string(*obj.SourceMicroservicesDeploymentName)
	}

	if obj.TargetDbAdminCredentials != nil {
		result["target_db_admin_credentials"] = []interface{}{AdminCredentialsToMap(obj.TargetDbAdminCredentials)}
	}

	if obj.TargetMicroservicesDeploymentName != nil {
		result["target_microservices_deployment_name"] = string(*obj.TargetMicroservicesDeploymentName)
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func GoldenGateHubToMapPass(obj *oci_database_migration.GoldenGateHub, resourceData *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComputeId != nil {
		result["compute_id"] = string(*obj.ComputeId)
	}

	if obj.RestAdminCredentials != nil {
		//result["rest_admin_credentials"] = []interface{}{AdminCredentialsToMap(obj.RestAdminCredentials)}
		result["rest_admin_credentials"] = []interface{}{AdminCredentialsToMapPasswordRest(obj, resourceData)}

	}

	if obj.SourceContainerDbAdminCredentials != nil {
		//result["source_container_db_admin_credentials"] = []interface{}{AdminCredentialsToMap(obj.SourceContainerDbAdminCredentials)}
		result["source_container_db_admin_credentials"] = []interface{}{AdminCredentialsToMapPasswordContainer(obj, resourceData)}

	}

	if obj.SourceDbAdminCredentials != nil {
		//result["source_db_admin_credentials"] = []interface{}{AdminCredentialsToMap(obj.SourceDbAdminCredentials)}
		result["source_db_admin_credentials"] = []interface{}{AdminCredentialsToMapPasswordSource(obj, resourceData)}

	}

	if obj.SourceMicroservicesDeploymentName != nil {
		result["source_microservices_deployment_name"] = string(*obj.SourceMicroservicesDeploymentName)
	}

	if obj.TargetDbAdminCredentials != nil {
		//result["target_db_admin_credentials"] = []interface{}{AdminCredentialsToMap(obj.TargetDbAdminCredentials)}
		result["target_db_admin_credentials"] = []interface{}{AdminCredentialsToMapPasswordTarget(obj, resourceData)}

	}

	if obj.TargetMicroservicesDeploymentName != nil {
		result["target_microservices_deployment_name"] = string(*obj.TargetMicroservicesDeploymentName)
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateGoldenGateSettings(fieldKeyFormat string) (oci_database_migration.CreateGoldenGateSettings, error) {
	result := oci_database_migration.CreateGoldenGateSettings{}

	if acceptableLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "acceptable_lag")); ok {
		tmp := acceptableLag.(int)
		result.AcceptableLag = &tmp
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

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateGoldenGateSettings(fieldKeyFormat string) (oci_database_migration.UpdateGoldenGateSettings, error) {
	result := oci_database_migration.UpdateGoldenGateSettings{}

	if acceptableLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "acceptable_lag")); ok {
		tmp := acceptableLag.(int)
		result.AcceptableLag = &tmp
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

func GoldenGateSettingsToMap(obj *oci_database_migration.GoldenGateSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptableLag != nil {
		result["acceptable_lag"] = int(*obj.AcceptableLag)
	}

	if obj.Extract != nil {
		result["extract"] = []interface{}{ExtractToMap(obj.Extract)}
	}

	if obj.Replicat != nil {
		result["replicat"] = []interface{}{ReplicatToMap(obj.Replicat)}
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateHostDumpTransferDetails(fieldKeyFormat string) (oci_database_migration.CreateHostDumpTransferDetails, error) {
	var baseObject oci_database_migration.CreateHostDumpTransferDetails
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "CURL" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("CURL"):
		details := oci_database_migration.UpdateCurlTransferDetails{}
		baseObject = details
	case strings.ToLower("OCI_CLI"):
		details := oci_database_migration.UpdateOciCliDumpTransferDetails{}
		if ociHome, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oci_home")); ok {
			tmp := ociHome.(string)
			details.OciHome = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateHostDumpTransferDetails(fieldKeyFormat string) (oci_database_migration.UpdateHostDumpTransferDetails, error) {
	var baseObject oci_database_migration.UpdateHostDumpTransferDetails
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "CURL" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("CURL"):
		details := oci_database_migration.UpdateCurlTransferDetails{}
		baseObject = details
	case strings.ToLower("OCI_CLI"):
		details := oci_database_migration.UpdateOciCliDumpTransferDetails{}
		if ociHome, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oci_home")); ok {
			tmp := ociHome.(string)
			details.OciHome = &tmp
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
	case oci_database_migration.UpdateCurlTransferDetails:
		result["kind"] = "CURL"
	case oci_database_migration.UpdateOciCliDumpTransferDetails:
		result["kind"] = "OCI_CLI"

		if v.OciHome != nil {
			result["oci_home"] = string(*v.OciHome)
		}
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", *obj)
		return nil
	}

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

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateReplicat(fieldKeyFormat string) (oci_database_migration.CreateReplicat, error) {
	result := oci_database_migration.CreateReplicat{}

	if mapParallelism, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "map_parallelism")); ok {
		tmp := mapParallelism.(int)
		result.MapParallelism = &tmp
	}

	if maxApplyParallelism, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_apply_parallelism")); ok {
		tmp := maxApplyParallelism.(int)
		result.MaxApplyParallelism = &tmp
	}

	if minApplyParallelism, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_apply_parallelism")); ok {
		tmp := minApplyParallelism.(int)
		result.MinApplyParallelism = &tmp
	}

	return result, nil
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateReplicat(fieldKeyFormat string) (oci_database_migration.UpdateReplicat, error) {
	result := oci_database_migration.UpdateReplicat{}

	if mapParallelism, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "map_parallelism")); ok {
		tmp := mapParallelism.(int)
		result.MapParallelism = &tmp
	}

	if maxApplyParallelism, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_apply_parallelism")); ok {
		tmp := maxApplyParallelism.(int)
		result.MaxApplyParallelism = &tmp
	}

	if minApplyParallelism, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_apply_parallelism")); ok {
		tmp := minApplyParallelism.(int)
		result.MinApplyParallelism = &tmp
	}

	return result, nil
}

func ReplicatToMap(obj *oci_database_migration.Replicat) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MapParallelism != nil {
		result["map_parallelism"] = int(*obj.MapParallelism)
	}

	if obj.MaxApplyParallelism != nil {
		result["max_apply_parallelism"] = int(*obj.MaxApplyParallelism)
	}

	if obj.MinApplyParallelism != nil {
		result["min_apply_parallelism"] = int(*obj.MinApplyParallelism)
	}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToCreateVaultDetails(fieldKeyFormat string) (oci_database_migration.CreateVaultDetails, error) {
	result := oci_database_migration.CreateVaultDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if keyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_id")); ok {
		tmp := keyId.(string)
		result.KeyId = &tmp
	}

	if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
		tmp := vaultId.(string)
		result.VaultId = &tmp
	}

	return result, nil
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToUpdateVaultDetails(fieldKeyFormat string) (oci_database_migration.UpdateVaultDetails, error) {
	result := oci_database_migration.UpdateVaultDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if keyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_id")); ok {
		tmp := keyId.(string)
		result.KeyId = &tmp
	}

	if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
		tmp := vaultId.(string)
		result.VaultId = &tmp
	}

	return result, nil
}

/*func VaultDetailsToMap(obj *oci_database_migration.VaultDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.KeyId != nil {
		result["key_id"] = string(*obj.KeyId)
	}

	if obj.VaultId != nil {
		result["vault_id"] = string(*obj.VaultId)
	}

	return result
}*/

func (s *DatabaseMigrationMigrationResourceCrud) mapToDataPumpExcludeParameters(fieldKeyFormat string) (oci_database_migration.DataPumpExcludeParametersEnum, error) {
	//result := make([]oci_database_migration.DataPumpExcludeParametersEnum, 3)
	result := oci_database_migration.DataPumpExcludeParametersIndex
	return result, nil
}

func DataPumpExcludeParametersToMap(obj oci_database_migration.DataPumpExcludeParametersEnum) map[string]interface{} {
	result := map[string]interface{}{}

	return result
}

func (s *DatabaseMigrationMigrationResourceCrud) mapToDatabaseObject(fieldKeyFormat string) (oci_database_migration.DatabaseObject, error) {
	result := oci_database_migration.DatabaseObject{}

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

func DatabaseObjectToMap(obj oci_database_migration.DatabaseObject) map[string]interface{} {
	result := map[string]interface{}{}

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

func MigrationSummaryToMap(obj oci_database_migration.MigrationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
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

	if obj.ExecutingJobId != nil {
		result["executing_job_id"] = string(*obj.ExecutingJobId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["lifecycle_details"] = string(obj.LifecycleDetails)

	if obj.SourceContainerDatabaseConnectionId != nil {
		result["source_container_database_connection_id"] = string(*obj.SourceContainerDatabaseConnectionId)
	}

	if obj.SourceDatabaseConnectionId != nil {
		result["source_database_connection_id"] = string(*obj.SourceDatabaseConnectionId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetDatabaseConnectionId != nil {
		result["target_database_connection_id"] = string(*obj.TargetDatabaseConnectionId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastMigration != nil {
		result["time_last_migration"] = obj.TimeLastMigration.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	if obj.VaultDetails != nil {
		result["vault_details"] = []interface{}{VaultDetailsToMap(obj.VaultDetails)}
	}

	return result
}

/*func VaultDetailsToMap(obj *oci_database_migration.VaultDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.KeyId != nil {
		result["key_id"] = string(*obj.KeyId)
	}

	if obj.VaultId != nil {
		result["vault_id"] = string(*obj.VaultId)
	}

	return result
}*/

func metadataRemapsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if newValue, ok := m["new_value"]; ok && newValue != "" {
		buf.WriteString(fmt.Sprintf("%v-", newValue))
	}
	if oldValue, ok := m["old_value"]; ok && oldValue != "" {
		buf.WriteString(fmt.Sprintf("%v-", oldValue))
	}
	if type_, ok := m["type"]; ok && type_ != "" {
		buf.WriteString(fmt.Sprintf("%v-", type_))
	}
	return hashcode.String(buf.String())
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
