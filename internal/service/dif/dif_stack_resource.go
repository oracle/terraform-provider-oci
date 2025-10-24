// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dif

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dif "github.com/oracle/oci-go-sdk/v65/dif"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var (
	difServiceToTemplate = map[string]string{
		"OBJECTSTORAGE": "DATALAKE",
		"ADB":           "DATALAKE",
		"DATAFLOW":      "DATATRANSFORMATION",
		"GGCS":          "DATAPIPELINE",
		"GENAI":         "AISERVICES",
	}
	difBlockToService = map[string]string{
		"adb":           "ADB",
		"dataflow":      "DATAFLOW",
		"genai":         "GENAI",
		"ggcs":          "GGCS",
		"objectstorage": "OBJECTSTORAGE",
	}
)

func DifStackResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		CreateContext: createDifStackWithContext,
		ReadContext:   readDifStackWithContext,
		UpdateContext: updateDifStackWithContext,
		DeleteContext: deleteDifStackWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"services": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"stack_templates": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"adb": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"admin_password_id": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"data_storage_size_in_tbs": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"db_version": {
							Type:     schema.TypeString,
							Required: true,
						},
						"db_workload": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ecpu": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"is_mtls_connection_required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_public": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tools_public_access": {
							Type:     schema.TypeString,
							Optional: true,
						},
						// Artifact deployment fields
						"artifact_object_storage_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"db_credentials": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"secret_id": {
										Type:      schema.TypeString,
										Required:  true,
										Sensitive: true,
									},
									"user_type": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						// Computed
					},
				},
			},
			"dataflow": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"driver_shape": {
							Type:     schema.TypeString,
							Required: true,
						},
						"executor_shape": {
							Type:     schema.TypeString,
							Required: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"log_bucket_instance_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"num_executors": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"spark_version": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"connections": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"connection_details": {
										Type:     schema.TypeList,
										Required: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"dif_dependencies": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"service_instance_id": {
																Type:     schema.TypeString,
																Required: true,
															},
															"service_type": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},
												"domain_names": {
													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												// Computed
											},
										},
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"driver_shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"memory_in_gbs": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"ocpus": {
										Type:     schema.TypeInt,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"executor_shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"memory_in_gbs": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"ocpus": {
										Type:     schema.TypeInt,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"private_endpoint_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"warehouse_bucket_instance_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						// Artifact deployment fields
						"execute": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"archive_uri": {
							Type:     schema.TypeString,
							Optional: true,
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
			"genai": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"base_model": {
							Type:     schema.TypeString,
							Required: true,
						},
						"cluster_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"oci_region": {
							Type:     schema.TypeString,
							Required: true,
						},
						"unit_count": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"endpoints": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"endpoint_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"is_content_moderation_enabled": {
										Type:     schema.TypeBool,
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
			"ggcs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"instance_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ocpu": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"password_secret_id": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"connections": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"connection_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"connection_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"dif_dependencies": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"service_instance_id": {
													Type:     schema.TypeString,
													Required: true,
												},
												"service_type": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
									"gg_admin_secret_id": {
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
								},
							},
						},
						"ogg_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"public_subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						// Artifact deployment fields
						"artifact_object_storage_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"users": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"secret_id": {
										Type:      schema.TypeString,
										Required:  true,
										Sensitive: true,
									},
									"user_type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"action": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"sources": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"action": {
										Type:     schema.TypeString,
										Required: true,
									},
									"should_start_source_operations": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"target_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"target_connection_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"targets": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"action": {
										Type:     schema.TypeString,
										Required: true,
									},
									"should_start_target_operations": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"source_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"source_connection_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						// Computed
					},
				},
			},
			"notification_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"objectstorage": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"instance_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"object_versioning": {
							Type:     schema.TypeString,
							Required: true,
						},
						"storage_tier": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"auto_tiering": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"add_service_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"deploy_artifacts_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"additional_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assigned_connections": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"connection_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"connection_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"requested_by": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"endpoint_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"endpoint_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"endpoint_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"model_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"model_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oci_region": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_endpoint_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"current_artifact_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			//subnet id for artifacts deploy
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
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
}

func createDifStackWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DifStackResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackClient()

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	if _, ok := sync.D.GetOkExists("deploy_artifacts_trigger"); ok {
		err := sync.DeployArtifactsWithContext(ctx)
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}
	return nil

}

func readDifStackWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DifStackResourceCrud{InRefresh: true}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

/*
scanStackChanges inspects the current planned changes and classifies them into:
- hasAdditions: new services/templates or new instances in service blocks
- hasDeletions: removed services/templates or removed instances in service blocks
- hasUpdates: in-place edits to existing service blocks or top-level attributes (tags, notification_email)
*/
func scanStackChanges(d *schema.ResourceData) (hasAdditions, hasDeletions, hasUpdates bool) {
	// Services (TypeSet)
	if d.HasChange("services") {
		oldVal, newVal := d.GetChange("services")
		oldSet := oldVal.(*schema.Set)
		newSet := newVal.(*schema.Set)
		oldList := oldSet.List()
		newList := newSet.List()

		oldMap := make(map[string]bool)
		for _, v := range oldList {
			if v != nil {
				oldMap[v.(string)] = true
			}
		}
		newMap := make(map[string]bool)
		for _, v := range newList {
			if v != nil {
				newMap[v.(string)] = true
			}
		}
		for v := range newMap {
			if !oldMap[v] {
				hasAdditions = true
			}
		}
		for v := range oldMap {
			if !newMap[v] {
				hasDeletions = true
			}
		}
	}

	// Stack templates (TypeSet)
	if d.HasChange("stack_templates") {
		oldVal, newVal := d.GetChange("stack_templates")
		oldSet := oldVal.(*schema.Set)
		newSet := newVal.(*schema.Set)
		oldList := oldSet.List()
		newList := newSet.List()

		oldMap := make(map[string]bool)
		for _, v := range oldList {
			if v != nil {
				oldMap[v.(string)] = true
			}
		}
		newMap := make(map[string]bool)
		for _, v := range newList {
			if v != nil {
				newMap[v.(string)] = true
			}
		}
		for v := range newMap {
			if !oldMap[v] {
				hasAdditions = true
			}
		}
		for v := range oldMap {
			if !newMap[v] {
				hasDeletions = true
			}
		}
	}

	// Service blocks: adb, dataflow, genai, ggcs, objectstorage
	for _, block := range []string{"adb", "dataflow", "genai", "ggcs", "objectstorage"} {
		if d.HasChange(block) {
			oldRaw, newRaw := d.GetChange(block)
			oldList := oldRaw.([]interface{})
			newList := newRaw.([]interface{})
			if len(newList) > len(oldList) {
				hasAdditions = true
			} else if len(newList) < len(oldList) {
				hasDeletions = true
			} else {
				// In-place edits to the block (note: this includes artifact fields as they share the same schema)
				hasUpdates = true
			}
		}
	}

	// Top-level fields treated as standard updates
	for _, field := range []string{"defined_tags", "freeform_tags", "notification_email"} {
		if d.HasChange(field) {
			hasUpdates = true
		}
	}

	return
}

// didTriggerIncrease returns true if a numeric trigger field increased (new > old)
func didTriggerIncrease(d *schema.ResourceData, field string) bool {
	if _, ok := d.GetOkExists(field); ok && d.HasChange(field) {
		oldRaw, newRaw := d.GetChange(field)
		oldVal, okOld := oldRaw.(int)
		newVal, okNew := newRaw.(int)
		if okOld && okNew {
			return newVal > oldVal
		}
	}
	return false
}

func updateDifStackWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DifStackResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackClient()

	if _, ok := sync.D.GetOkExists("add_service_trigger"); ok && sync.D.HasChange("add_service_trigger") {
		oldRaw, newRaw := sync.D.GetChange("add_service_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			sync.SkipSetData = true
			err := sync.AddServiceWithContext(ctx)
			if err != nil {
				// leave SkipSetData true on error
				log.Printf("[DEBUG] Add Service Failed ")
				sync.D.Set("add_service_trigger", oldRaw)
				// roll back deploy_artifacts_trigger as well so its diff is preserved
				if sync.D.HasChange("deploy_artifacts_trigger") {
					if depOld, _ := sync.D.GetChange("deploy_artifacts_trigger"); depOld != nil {
						sync.D.Set("deploy_artifacts_trigger", depOld)
					}
				}
				return tfresource.HandleDiagError(m, err)
			}
			sync.SkipSetData = false
		} else {
			sync.D.Set("add_service_trigger", oldRaw)
			return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
		}
	}

	if _, ok := sync.D.GetOkExists("deploy_artifacts_trigger"); ok && sync.D.HasChange("deploy_artifacts_trigger") {
		oldRaw, newRaw := sync.D.GetChange("deploy_artifacts_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			sync.SkipSetData = true
			err := sync.DeployArtifactsWithContext(ctx)
			if err != nil {
				// leave SkipSetData true on error
				sync.D.Set("deploy_artifacts_trigger", oldRaw)
				return tfresource.HandleDiagError(m, err)
			}
			sync.SkipSetData = false
		} else {
			sync.D.Set("deploy_artifacts_trigger", oldRaw)
			return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
		}
	}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	return nil
}

func deleteDifStackWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DifStackResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DifStackResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dif.StackClient
	Res                    *oci_dif.Stack
	DisableNotFoundRetries bool
	SkipSetData            bool
	InRefresh              bool
}

func (s *DifStackResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DifStackResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dif.StackLifecycleStateCreating),
	}
}

func (s *DifStackResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dif.StackLifecycleStateActive),
		string(oci_dif.StackLifecycleStateNeedsAttention),
	}
}

func (s *DifStackResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dif.StackLifecycleStateDeleting),
	}
}

func (s *DifStackResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dif.StackLifecycleStateDeleted),
	}
}

func (s *DifStackResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_dif.CreateStackRequest{}

	if adb, ok := s.D.GetOkExists("adb"); ok {
		interfaces := adb.([]interface{})
		tmp := make([]oci_dif.AdbDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "adb", stateDataIndex)
			converted, err := s.mapToAdbDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("adb") {
			request.Adb = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataflow, ok := s.D.GetOkExists("dataflow"); ok {
		interfaces := dataflow.([]interface{})
		tmp := make([]oci_dif.DataflowDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dataflow", stateDataIndex)
			converted, err := s.mapToDataflowDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("dataflow") {
			request.Dataflow = tmp
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

	if genai, ok := s.D.GetOkExists("genai"); ok {
		interfaces := genai.([]interface{})
		tmp := make([]oci_dif.GenAiDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "genai", stateDataIndex)
			converted, err := s.mapToGenAiDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("genai") {
			request.Genai = tmp
		}
	}

	if ggcs, ok := s.D.GetOkExists("ggcs"); ok {
		interfaces := ggcs.([]interface{})
		tmp := make([]oci_dif.GgcsDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ggcs", stateDataIndex)
			converted, err := s.mapToGgcsDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("ggcs") {
			request.Ggcs = tmp
		}
	}

	if notificationEmail, ok := s.D.GetOkExists("notification_email"); ok {
		tmp := notificationEmail.(string)
		request.NotificationEmail = &tmp
	}

	if objectstorage, ok := s.D.GetOkExists("objectstorage"); ok {
		interfaces := objectstorage.([]interface{})
		tmp := make([]oci_dif.ObjectStorageDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "objectstorage", stateDataIndex)
			converted, err := s.mapToObjectStorageDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("objectstorage") {
			request.Objectstorage = tmp
		}
	}

	if services, ok := s.D.GetOkExists("services"); ok {
		set := services.(*schema.Set)
		list := set.List()
		tmp := make([]oci_dif.ServiceEnum, len(list))
		for i := range list {
			if list[i] != nil {
				tmp[i] = oci_dif.ServiceEnum(list[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("services") {
			request.Services = tmp
		}
	}

	if stackTemplates, ok := s.D.GetOkExists("stack_templates"); ok {
		set := stackTemplates.(*schema.Set)
		list := set.List()
		tmp := make([]oci_dif.StackTemplateEnum, len(list))
		for i := range list {
			if list[i] != nil {
				tmp[i] = oci_dif.StackTemplateEnum(list[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("stack_templates") {
			request.StackTemplates = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif")

	response, err := s.Client.CreateStack(ctx, request)
	if err != nil {
		log.Printf("[ERROR] CreateStack failed: error=%v", err)
		return err
	}
	var _id, _wrId, _opc string
	if response.Id != nil {
		_id = *response.Id
	}
	if response.OpcWorkRequestId != nil {
		_wrId = *response.OpcWorkRequestId
	}
	if response.OpcRequestId != nil {
		_opc = *response.OpcRequestId
	}
	log.Printf("[INFO] CreateStack initiated: id=%s workRequestId=%s opcRequestId=%s", _id, _wrId, _opc)

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getStackFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif"), "dataintelligenceentity", oci_dif.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DifStackResourceCrud) getStackFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	entityType string, actionTypeEnum oci_dif.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	stackId, err := stackWaitForWorkRequest(ctx, workId, entityType,
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		var _wrId, _sid string
		if workId != nil {
			_wrId = *workId
		}
		if stackId != nil {
			_sid = *stackId
		}
		log.Printf("[ERROR] work request failed: workRequestId=%s stackId=%s entity=%s action=%s error=%v", _wrId, _sid, entityType, actionTypeEnum, err)
		return err
	}
	var _wrSuc, _sidSuc string
	if workId != nil {
		_wrSuc = *workId
	}
	if stackId != nil {
		_sidSuc = *stackId
	}
	log.Printf("[INFO] work request succeeded: workRequestId=%s stackId=%s entity=%s action=%s", _wrSuc, _sidSuc, entityType, actionTypeEnum)
	s.D.SetId(*stackId)

	return s.GetWithContext(ctx)
}

func stackWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "dif", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_dif.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func stackWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_dif.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_dif.StackClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "dif")
	retryPolicy.ShouldRetryOperation = stackWorkRequestShouldRetryFunc(timeout)

	response := oci_dif.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_dif.OperationStatusInProgress),
			string(oci_dif.OperationStatusAccepted),
			string(oci_dif.OperationStatusCancelling),
		},
		Target: []string{
			string(oci_dif.OperationStatusSucceeded),
			string(oci_dif.OperationStatusFailed),
			string(oci_dif.OperationStatusCancelled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_dif.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			if err != nil {
				var _wr string
				if wId != nil {
					_wr = *wId
				}
				log.Printf("[ERROR] GetWorkRequest failed: workRequestId=%s error=%v", _wr, err)
			}
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	for _, res := range response.Resources {
		if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}
	// if a single resource is returned, use its identifier
	if identifier == nil && len(response.Resources) == 1 {
		identifier = response.Resources[0].Identifier
	}

	if identifier == nil || response.Status == oci_dif.OperationStatusFailed || response.Status == oci_dif.OperationStatusCancelled {
		return nil, getErrorFromDifStackWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDifStackWorkRequest(ctx context.Context, client *oci_dif.StackClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_dif.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_dif.ListWorkRequestErrorsRequest{
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

	var _wr string
	if workId != nil {
		_wr = *workId
	}
	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", _wr, entityType, action, errorMessage)

	return workRequestErr
}

func (s *DifStackResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_dif.GetStackRequest{}

	tmp := s.D.Id()
	request.StackId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif")

	response, err := s.Client.GetStack(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Stack
	return nil
}

func (s *DifStackResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			s.SkipSetData = true
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				// leave SkipSetData true on error
				s.D.Set("compartment_id", oldRaw)
				return err
			}
			s.SkipSetData = false
		}
	}

	hasAdditions, hasDeletions, hasUpdates := scanStackChanges(s.D)

	triggerChanged := s.D.HasChange("add_service_trigger")

	if hasDeletions {
		return fmt.Errorf("Deletions are not permitted for services, stack templates, or service blocks")
	}
	if hasAdditions && !triggerChanged {
		return fmt.Errorf("Additions require increasing add_service_trigger; increment this value to proceed")
	}
	if !hasUpdates {
		return nil
	}

	// Build UpdateStack request only with changed fields
	request := oci_dif.UpdateStackRequest{}
	var changedServices []string

	if adb, ok := s.D.GetOkExists("adb"); ok && s.D.HasChange("adb") {
		interfaces := adb.([]interface{})
		tmp := make([]oci_dif.AdbUpdateDetail, 0, len(interfaces))
		for i := range interfaces {
			if s.D.HasChange(fmt.Sprintf("adb.%d.data_storage_size_in_tbs", i)) ||
				s.D.HasChange(fmt.Sprintf("adb.%d.ecpu", i)) ||
				s.D.HasChange(fmt.Sprintf("adb.%d.is_mtls_connection_required", i)) {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "adb", i)
				converted, err := s.mapToAdbUpdateDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp = append(tmp, converted)
			}
		}
		if len(tmp) > 0 {
			request.Adb = tmp
			changedServices = append(changedServices, "ADB")
		}
	}

	if dataflow, ok := s.D.GetOkExists("dataflow"); ok && s.D.HasChange("dataflow") {
		interfaces := dataflow.([]interface{})
		tmp := make([]oci_dif.DataflowUpdateDetail, 0, len(interfaces))
		for i := range interfaces {
			if s.D.HasChange(fmt.Sprintf("dataflow.%d.num_executors", i)) ||
				s.D.HasChange(fmt.Sprintf("dataflow.%d.driver_shape", i)) ||
				s.D.HasChange(fmt.Sprintf("dataflow.%d.executor_shape", i)) ||
				s.D.HasChange(fmt.Sprintf("dataflow.%d.spark_version", i)) ||
				s.D.HasChange(fmt.Sprintf("dataflow.%d.driver_shape_config", i)) ||
				s.D.HasChange(fmt.Sprintf("dataflow.%d.executor_shape_config", i)) ||
				s.D.HasChange(fmt.Sprintf("dataflow.%d.private_endpoint_id", i)) ||
				s.D.HasChange(fmt.Sprintf("dataflow.%d.connections", i)) ||
				s.D.HasChange(fmt.Sprintf("dataflow.%d.log_bucket_instance_id", i)) ||
				s.D.HasChange(fmt.Sprintf("dataflow.%d.warehouse_bucket_instance_id", i)) {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dataflow", i)
				converted, err := s.mapToDataflowUpdateDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp = append(tmp, converted)
			}
		}
		if len(tmp) > 0 {
			request.Dataflow = tmp
			changedServices = append(changedServices, "DATAFLOW")
		}
	}

	// Tags
	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}
	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if genai, ok := s.D.GetOkExists("genai"); ok && s.D.HasChange("genai") {
		interfaces := genai.([]interface{})
		tmp := make([]oci_dif.GenAiUpdateDetail, 0, len(interfaces))
		for i := range interfaces {
			if s.D.HasChange(fmt.Sprintf("genai.%d.unit_count", i)) ||
				s.D.HasChange(fmt.Sprintf("genai.%d.endpoints", i)) {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "genai", i)
				converted, err := s.mapToGenAiUpdateDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp = append(tmp, converted)
			}
		}
		if len(tmp) > 0 {
			request.Genai = tmp
			changedServices = append(changedServices, "GENAI")
		}
	}

	if ggcs, ok := s.D.GetOkExists("ggcs"); ok && s.D.HasChange("ggcs") {
		interfaces := ggcs.([]interface{})
		tmp := make([]oci_dif.GgcsUpdateDetail, 0, len(interfaces))
		for i := range interfaces {
			if s.D.HasChange(fmt.Sprintf("ggcs.%d.ocpu", i)) ||
				s.D.HasChange(fmt.Sprintf("ggcs.%d.connections", i)) ||
				s.D.HasChange(fmt.Sprintf("ggcs.%d.public_subnet_id", i)) {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ggcs", i)
				converted, err := s.mapToGgcsUpdateDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp = append(tmp, converted)
			}
		}
		if len(tmp) > 0 {
			request.Ggcs = tmp
			changedServices = append(changedServices, "GGCS")
		}
	}

	if notificationEmail, ok := s.D.GetOkExists("notification_email"); ok && s.D.HasChange("notification_email") {
		tmp := notificationEmail.(string)
		request.NotificationEmail = &tmp
	}

	if objectstorage, ok := s.D.GetOkExists("objectstorage"); ok && s.D.HasChange("objectstorage") {
		interfaces := objectstorage.([]interface{})
		tmp := make([]oci_dif.ObjectStorageUpdateDetail, 0, len(interfaces))
		for i := range interfaces {
			if s.D.HasChange(fmt.Sprintf("objectstorage.%d.auto_tiering", i)) ||
				s.D.HasChange(fmt.Sprintf("objectstorage.%d.object_versioning", i)) {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "objectstorage", i)
				converted, err := s.mapToObjectStorageUpdateDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp = append(tmp, converted)
			}
		}
		if len(tmp) > 0 {
			request.Objectstorage = tmp
			changedServices = append(changedServices, "OBJECTSTORAGE")
		}
	}

	// Derive services and templates from changed blocks
	if len(changedServices) > 0 {
		var svc []oci_dif.ServiceEnum
		var tmpl []oci_dif.StackTemplateEnum
		for _, sName := range changedServices {
			svc = append(svc, oci_dif.ServiceEnum(sName))
			if tStr, ok := difServiceToTemplate[sName]; ok {
				tmpl = append(tmpl, oci_dif.StackTemplateEnum(tStr))
			}
		}
		if len(svc) > 0 {
			request.Services = deduplicateServices(svc)
		}
		if len(tmpl) > 0 {
			request.StackTemplates = deduplicateTemplates(tmpl)
		}
	}

	// Set stack id and call API
	stackId := s.D.Id()
	request.StackId = &stackId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif")

	// Skip if no actual updates
	if len(changedServices) == 0 &&
		request.DefinedTags == nil &&
		request.FreeformTags == nil &&
		request.NotificationEmail == nil {
		log.Printf("[DEBUG] No updates to apply for stack %s, skipping UpdateStack", stackId)
	} else {
		// prevent state overwrite until the work-request finishes successfully
		s.SkipSetData = true
		response, err := s.Client.UpdateStack(ctx, request)
		if err != nil {
			log.Printf("[ERROR] UpdateStack failed: stackId=%s error=%v", stackId, err)
			return err
		}
		var _updWrId, _updOpc string
		if response.OpcWorkRequestId != nil {
			_updWrId = *response.OpcWorkRequestId
		}
		if response.OpcRequestId != nil {
			_updOpc = *response.OpcRequestId
		}
		log.Printf("[INFO] UpdateStack initiated: stackId=%s workRequestId=%s opcRequestId=%s", stackId, _updWrId, _updOpc)

		workId := response.OpcWorkRequestId
		err = s.getStackFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif"),
			"dataintelligenceentity", oci_dif.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.SkipSetData = false
		return nil
	}
	return nil
}

func (s *DifStackResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_dif.DeleteStackRequest{}

	tmp := s.D.Id()
	request.StackId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif")

	response, err := s.Client.DeleteStack(ctx, request)
	if err != nil {
		log.Printf("[ERROR] DeleteStack failed: stackId=%s error=%v", *request.StackId, err)
		return err
	}
	var _delWrId, _delOpc string
	if response.OpcWorkRequestId != nil {
		_delWrId = *response.OpcWorkRequestId
	}
	if response.OpcRequestId != nil {
		_delOpc = *response.OpcRequestId
	}
	log.Printf("[INFO] DeleteStack initiated: stackId=%s workRequestId=%s opcRequestId=%s", *request.StackId, _delWrId, _delOpc)

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := stackWaitForWorkRequest(ctx, workId, "dataintelligenceentity",
		oci_dif.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DifStackResourceCrud) SetData() error {
	if s.SkipSetData {
		return nil // error path – keep diff
	}

	// When called from Read (pre-plan refresh) we must keep only already-existing
	// service instances; dropping not-yet-created ones preserves the diff after a
	// failed AddService / UpdateStack.
	if s.InRefresh {
		existing := map[string]bool{}
		if s.Res != nil && s.Res.ServiceDetails != nil {
			for _, sd := range s.Res.ServiceDetails {
				if sd.ServiceType != nil && sd.InstanceId != nil && *sd.InstanceId != "" {
					key := fmt.Sprintf("%s#%s", *sd.ServiceType, *sd.InstanceId)
					existing[key] = true
				}
			}
		}
		filter := func(block, svc string) {
			if cfg, ok := s.D.GetOk(block); ok {
				kept := make([]interface{}, 0)
				for _, raw := range cfg.([]interface{}) {
					if raw == nil {
						continue
					}
					if m, ok := raw.(map[string]interface{}); ok {
						if idRaw, ok := m["instance_id"]; ok && idRaw.(string) != "" {
							if existing[fmt.Sprintf("%s#%s", svc, idRaw.(string))] {
								kept = append(kept, raw)
							}
						}
					}
				}
				if len(kept) > 0 {
					s.D.Set(block, kept)
				} else {
					s.D.Set(block, nil)
				}
			} else {
				s.D.Set(block, nil)
			}
		}
		for blk, svc := range difBlockToService {
			filter(blk, svc)
		}
		// fall through to copy remote attributes below (tags, state, etc.)
	}

	if s.Res == nil {
		return nil
	}

	adb := []interface{}{}
	if s.Res.Adb != nil {
		for _, item := range s.Res.Adb {
			if item.InstanceId != nil && *item.InstanceId != "" {
				adb = append(adb, AdbDetailToMap(item))
			}
		}
	}
	if configAdb, ok := s.D.GetOk("adb"); ok {
		s.D.Set("adb", configAdb) // Preserve user's configuration
	} else {
		s.D.Set("adb", adb) // Use API data only when no user config
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	dataflow := []interface{}{}
	if s.Res.Dataflow != nil {
		for _, item := range s.Res.Dataflow {
			if item.InstanceId != nil && *item.InstanceId != "" {
				dataflow = append(dataflow, DataflowDetailToMap(item))
			}
		}
	}
	if cfg, ok := s.D.GetOk("dataflow"); ok {
		s.D.Set("dataflow", cfg)
	} else {
		s.D.Set("dataflow", dataflow)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if cfg, ok := s.D.GetOk("genai"); ok {
		s.D.Set("genai", cfg)
	} else {
		genai := []interface{}{}
		if s.Res.Genai != nil {
			for _, item := range s.Res.Genai {
				genai = append(genai, GenAiDetailToMap(item))
			}
		}
		s.D.Set("genai", genai)
	}

	if cfg, ok := s.D.GetOk("ggcs"); ok {
		s.D.Set("ggcs", cfg)
	} else {
		ggcs := []interface{}{}
		if s.Res.Ggcs != nil {
			for _, item := range s.Res.Ggcs {
				ggcs = append(ggcs, GgcsDetailToMap(item))
			}
		}
		s.D.Set("ggcs", ggcs)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NotificationEmail != nil {
		s.D.Set("notification_email", *s.Res.NotificationEmail)
	}

	if cfg, ok := s.D.GetOk("objectstorage"); ok {
		s.D.Set("objectstorage", cfg)
	} else {
		objectstorage := []interface{}{}
		if s.Res.Objectstorage != nil {
			for _, item := range s.Res.Objectstorage {
				objectstorage = append(objectstorage, ObjectStorageDetailToMap(item))
			}
		}
		s.D.Set("objectstorage", objectstorage)
	}

	serviceDetails := []interface{}{}
	if s.Res.ServiceDetails != nil {
		for _, item := range s.Res.ServiceDetails {
			serviceDetails = append(serviceDetails, ServiceDetailResponseToMap(item))
		}
	}
	s.D.Set("service_details", serviceDetails)

	if s.Res.Services != nil {
		vals := make([]interface{}, 0, len(s.Res.Services))
		for _, v := range s.Res.Services {
			vals = append(vals, string(v))
		}
		s.D.Set("services", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, vals))
	}
	if s.Res.StackTemplates != nil {
		vals := make([]interface{}, 0, len(s.Res.StackTemplates))
		for _, v := range s.Res.StackTemplates {
			vals = append(vals, string(v))
		}
		s.D.Set("stack_templates", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, vals))
	}

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

func (s *DifStackResourceCrud) AddServiceWithContext(ctx context.Context) error {
	request := oci_dif.AddServiceRequest{}

	// Fetch current stack to detect existing service instances
	stackId := s.D.Id()
	current, err := s.Client.GetStack(ctx, oci_dif.GetStackRequest{
		StackId: &stackId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif"),
		},
	})
	if err != nil {
		log.Printf("[ERROR] GetStack for AddService failed: stackId=%s error=%v", stackId, err)
		return err
	}
	var _opc string
	if current.OpcRequestId != nil {
		_opc = *current.OpcRequestId
	}
	log.Printf("[INFO] GetStack for AddService success: stackId=%s opcRequestId=%s", stackId, _opc)

	existingServiceInstances := make(map[string]bool)
	for _, sd := range current.Stack.ServiceDetails {
		if sd.InstanceId != nil && *sd.InstanceId != "" {
			existingServiceInstances[*sd.InstanceId] = true
		}
	}

	// Track which service blocks have new instances
	newServiceInstances := make(map[string]bool)
	for block, serviceName := range difBlockToService {
		if s.D.HasChange(block) {
			oldRaw, newRaw := s.D.GetChange(block)
			oldList := oldRaw.([]interface{})
			newList := newRaw.([]interface{})
			if len(newList) > len(oldList) {
				newServiceInstances[serviceName] = true
			}
		}
	}

	// Build per-service request only for new instances
	if newServiceInstances["ADB"] {
		if adb, ok := s.D.GetOkExists("adb"); ok {
			interfaces := adb.([]interface{})
			var filtered []oci_dif.AdbDetail
			for i := range interfaces {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "adb", i)
				converted, err := s.mapToAdbDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				if converted.InstanceId != nil && !existingServiceInstances[*converted.InstanceId] {
					filtered = append(filtered, converted)
				}
			}
			if len(filtered) > 0 {
				request.Adb = filtered
			}
		}
	}
	if newServiceInstances["DATAFLOW"] {
		if dataflow, ok := s.D.GetOkExists("dataflow"); ok {
			interfaces := dataflow.([]interface{})
			var filtered []oci_dif.DataflowDetail
			for i := range interfaces {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dataflow", i)
				converted, err := s.mapToDataflowDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				if converted.InstanceId != nil && !existingServiceInstances[*converted.InstanceId] {
					filtered = append(filtered, converted)
				}
			}
			if len(filtered) > 0 {
				request.Dataflow = filtered
			}
		}
	}
	if newServiceInstances["GENAI"] {
		if genai, ok := s.D.GetOkExists("genai"); ok {
			interfaces := genai.([]interface{})
			var filtered []oci_dif.GenAiDetail
			for i := range interfaces {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "genai", i)
				converted, err := s.mapToGenAiDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				if converted.InstanceId != nil && !existingServiceInstances[*converted.InstanceId] {
					filtered = append(filtered, converted)
				}
			}
			if len(filtered) > 0 {
				request.Genai = filtered
			}
		}
	}
	if newServiceInstances["GGCS"] {
		if ggcs, ok := s.D.GetOkExists("ggcs"); ok {
			interfaces := ggcs.([]interface{})
			var filtered []oci_dif.GgcsDetail
			for i := range interfaces {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ggcs", i)
				converted, err := s.mapToGgcsDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				if converted.InstanceId != nil && !existingServiceInstances[*converted.InstanceId] {
					filtered = append(filtered, converted)
				}
			}
			if len(filtered) > 0 {
				request.Ggcs = filtered
			}
		}
	}
	if newServiceInstances["OBJECTSTORAGE"] {
		if objectstorage, ok := s.D.GetOkExists("objectstorage"); ok {
			interfaces := objectstorage.([]interface{})
			var filtered []oci_dif.ObjectStorageDetail
			for i := range interfaces {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "objectstorage", i)
				converted, err := s.mapToObjectStorageDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				if converted.InstanceId != nil && !existingServiceInstances[*converted.InstanceId] {
					filtered = append(filtered, converted)
				}
			}
			if len(filtered) > 0 {
				request.Objectstorage = filtered
			}
		}
	}

	// Detect added services and templates from sets
	addedServices := make(map[string]bool)
	addedTemplates := make(map[string]bool)
	if s.D.HasChange("services") {
		oldSvc, newSvc := s.D.GetChange("services")
		oldSvcSet := oldSvc.(*schema.Set)
		newSvcSet := newSvc.(*schema.Set)
		for _, s := range newSvcSet.List() {
			svc := s.(string)
			if !oldSvcSet.Contains(svc) {
				addedServices[svc] = true
			}
		}
	}
	if s.D.HasChange("stack_templates") {
		oldTpl, newTpl := s.D.GetChange("stack_templates")
		oldTplSet := oldTpl.(*schema.Set)
		newTplSet := newTpl.(*schema.Set)
		for _, t := range newTplSet.List() {
			tpl := t.(string)
			if !oldTplSet.Contains(tpl) {
				addedTemplates[tpl] = true
			}
		}
	}

	// Get current services and templates
	currentServices := make(map[string]bool)
	if svcSet, ok := s.D.GetOk("services"); ok {
		for _, s := range svcSet.(*schema.Set).List() {
			currentServices[s.(string)] = true
		}
	}
	currentTemplates := make(map[string]bool)
	if tplSet, ok := s.D.GetOk("stack_templates"); ok {
		for _, t := range tplSet.(*schema.Set).List() {
			currentTemplates[t.(string)] = true
		}
	}

	// Add services from set additions
	var servicesToAdd []oci_dif.ServiceEnum
	var templatesToAdd []oci_dif.StackTemplateEnum
	for svc := range addedServices {
		servicesToAdd = append(servicesToAdd, oci_dif.ServiceEnum(svc))
		if tpl, ok := difServiceToTemplate[svc]; ok && currentTemplates[tpl] {
			templatesToAdd = append(templatesToAdd, oci_dif.StackTemplateEnum(tpl))
		}
	}

	// For new instances in blocks, add service and template if present in sets
	for svc := range newServiceInstances {
		if currentServices[svc] {
			servicesToAdd = append(servicesToAdd, oci_dif.ServiceEnum(svc))
			if tpl, ok := difServiceToTemplate[svc]; ok && currentTemplates[tpl] {
				templatesToAdd = append(templatesToAdd, oci_dif.StackTemplateEnum(tpl))
			}
		}
	}

	servicesToAdd = deduplicateServices(servicesToAdd)
	templatesToAdd = deduplicateTemplates(templatesToAdd)
	if len(servicesToAdd) > 0 {
		request.Services = servicesToAdd
	}
	if len(templatesToAdd) > 0 {
		request.StackTemplates = templatesToAdd
	}

	// If nothing net-new, reset trigger and exit
	if len(servicesToAdd) == 0 &&
		len(request.Adb) == 0 && len(request.Dataflow) == 0 && len(request.Genai) == 0 &&
		len(request.Ggcs) == 0 && len(request.Objectstorage) == 0 {
		val := s.D.Get("add_service_trigger")
		s.D.Set("add_service_trigger", val)
		log.Printf("[DEBUG] No new services to add")
		return nil
	}

	idTmp := s.D.Id()
	request.StackId = &idTmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif")

	resp, err := s.Client.AddService(ctx, request)
	if err != nil {
		log.Printf("[ERROR] AddService failed: stackId=%s error=%v", idTmp, err)
		return err
	}
	var _addWr, _addOpc string
	if resp.OpcWorkRequestId != nil {
		_addWr = *resp.OpcWorkRequestId
	}
	if resp.OpcRequestId != nil {
		_addOpc = *resp.OpcRequestId
	}
	log.Printf("[INFO] AddService initiated: stackId=%s workRequestId=%s opcRequestId=%s", idTmp, _addWr, _addOpc)
	workId := resp.OpcWorkRequestId
	err = s.getStackFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif"), "dataintelligenceentity", oci_dif.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	val := s.D.Get("add_service_trigger")
	s.D.Set("add_service_trigger", val)
	return nil
}

func (s *DifStackResourceCrud) DeployArtifactsWithContext(ctx context.Context) error {
	request := oci_dif.DeployArtifactsRequest{}

	// Only include instances with actual artifact payloads
	if adb, ok := s.D.GetOkExists("adb"); ok {
		interfaces := adb.([]interface{})
		var tmp []oci_dif.AdbArtifactsDetail
		for i := range interfaces {
			if interfaces[i] == nil {
				continue
			}
			if item, ok := interfaces[i].(map[string]interface{}); ok && hasAdbArtifactDeployment(item) {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "adb", i)
				converted, err := s.mapToAdbArtifactsDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp = append(tmp, converted)
			}
		}
		if len(tmp) > 0 {
			request.Adb = tmp
		}
	}
	if dataflow, ok := s.D.GetOkExists("dataflow"); ok {
		interfaces := dataflow.([]interface{})
		var tmp []oci_dif.DataflowArtifactsDetail
		for i := range interfaces {
			if interfaces[i] == nil {
				continue
			}
			if item, ok := interfaces[i].(map[string]interface{}); ok && hasDataflowArtifactDeployment(item) {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dataflow", i)
				converted, err := s.mapToDataflowArtifactsDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp = append(tmp, converted)
			}
		}
		if len(tmp) > 0 {
			request.Dataflow = tmp
		}
	}
	if ggcs, ok := s.D.GetOkExists("ggcs"); ok {
		interfaces := ggcs.([]interface{})
		var tmp []oci_dif.GgcsArtifactsDetail
		for i := range interfaces {
			if interfaces[i] == nil {
				continue
			}
			if item, ok := interfaces[i].(map[string]interface{}); ok && hasGgcsArtifactDeployment(item) {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ggcs", i)
				converted, err := s.mapToGgcsArtifactsDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp = append(tmp, converted)
			}
		}
		if len(tmp) > 0 {
			request.Ggcs = tmp
		}
	}

	// Build services and templates from those with artifacts
	servicesWithArtifacts := make(map[string]bool)
	templatesWithArtifacts := make(map[string]bool)
	if request.Adb != nil && len(request.Adb) > 0 {
		servicesWithArtifacts["ADB"] = true
		templatesWithArtifacts["DATALAKE"] = true
	}
	if request.Dataflow != nil && len(request.Dataflow) > 0 {
		servicesWithArtifacts["DATAFLOW"] = true
		templatesWithArtifacts["DATATRANSFORMATION"] = true
	}
	if request.Ggcs != nil && len(request.Ggcs) > 0 {
		servicesWithArtifacts["GGCS"] = true
		templatesWithArtifacts["DATAPIPELINE"] = true
	}

	// If nothing to deploy, reset trigger and exit
	if (request.Adb == nil || len(request.Adb) == 0) &&
		(request.Dataflow == nil || len(request.Dataflow) == 0) &&
		(request.Ggcs == nil || len(request.Ggcs) == 0) {
		val := s.D.Get("deploy_artifacts_trigger")
		s.D.Set("deploy_artifacts_trigger", val)
		return nil
	}

	// Assign services/templates derived from artifacts
	if len(servicesWithArtifacts) > 0 {
		var svc []oci_dif.ServiceEnum
		for k := range servicesWithArtifacts {
			svc = append(svc, oci_dif.ServiceEnum(k))
		}
		request.Services = svc
	}
	if len(templatesWithArtifacts) > 0 {
		var tmpl []oci_dif.StackTemplateEnum
		for k := range templatesWithArtifacts {
			tmpl = append(tmpl, oci_dif.StackTemplateEnum(k))
		}
		request.StackTemplates = tmpl
	}

	idTmp := s.D.Id()
	request.StackId = &idTmp

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif")

	response, err := s.Client.DeployArtifacts(ctx, request)
	if err != nil {
		log.Printf("[ERROR] DeployArtifacts failed: stackId=%s error=%v", idTmp, err)
		return err
	}
	var _wrDep, _opcDep string
	if response.OpcWorkRequestId != nil {
		_wrDep = *response.OpcWorkRequestId
	}
	if response.OpcRequestId != nil {
		_opcDep = *response.OpcRequestId
	}
	log.Printf("[INFO] DeployArtifacts initiated: stackId=%s workRequestId=%s opcRequestId=%s", idTmp, _wrDep, _opcDep)
	workId := response.OpcWorkRequestId
	err = s.getStackFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif"), "artifactdeployentity", oci_dif.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	val := s.D.Get("deploy_artifacts_trigger")
	s.D.Set("deploy_artifacts_trigger", val)
	return nil
}

func (s *DifStackResourceCrud) mapToAdbDetail(fieldKeyFormat string) (oci_dif.AdbDetail, error) {
	result := oci_dif.AdbDetail{}

	if adminPasswordId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password_id")); ok {
		tmp := adminPasswordId.(string)
		result.AdminPasswordId = &tmp
	}

	if dataStorageSizeInTBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_tbs")); ok {
		tmp := dataStorageSizeInTBs.(int)
		result.DataStorageSizeInTBs = &tmp
	}

	if dbVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_version")); ok {
		tmp := dbVersion.(string)
		result.DbVersion = &tmp
	}

	if dbWorkload, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_workload")); ok {
		result.DbWorkload = oci_dif.DbWorkloadEnum(dbWorkload.(string))
	}

	if ecpu, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ecpu")); ok {
		tmp := ecpu.(int)
		result.Ecpu = &tmp
	}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	if isMtlsConnectionRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_mtls_connection_required")); ok {
		tmp := isMtlsConnectionRequired.(bool)
		result.IsMtlsConnectionRequired = &tmp
	}

	if isPublic, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_public")); ok {
		tmp := isPublic.(bool)
		result.IsPublic = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	if toolsPublicAccess, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tools_public_access")); ok {
		tmp := toolsPublicAccess.(string)
		result.ToolsPublicAccess = &tmp
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToAdbUpdateDetail(fieldKeyFormat string) (oci_dif.AdbUpdateDetail, error) {
	result := oci_dif.AdbUpdateDetail{}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	// Only set fields that changed
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_tbs")) {
		if dataStorageSizeInTBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_tbs")); ok {
			tmp := dataStorageSizeInTBs.(int)
			result.DataStorageSizeInTBs = &tmp
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ecpu")) {
		if ecpu, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ecpu")); ok {
			tmp := ecpu.(int)
			result.Ecpu = &tmp
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "is_mtls_connection_required")) {
		if isMtlsConnectionRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_mtls_connection_required")); ok {
			tmp := isMtlsConnectionRequired.(bool)
			result.IsMtlsConnectionRequired = &tmp
		}
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToAdbArtifactsDetail(fieldKeyFormat string) (oci_dif.AdbArtifactsDetail, error) {
	result := oci_dif.AdbArtifactsDetail{}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	if objectStoragePath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "artifact_object_storage_path")); ok {
		tmp := objectStoragePath.(string)
		result.ArtifactObjectStoragePath = &tmp
	}
	if dbCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_credentials")); ok {
		interfaces := dbCredentials.([]interface{})
		tmp := make([]oci_dif.DbCredentialsDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "db_credentials"), stateDataIndex)
			converted, err := s.mapToDbCredentialsDetail(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "db_credentials")) {
			result.DbCredentials = tmp
		}
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToDbCredentialsDetail(fieldKeyFormat string) (oci_dif.DbCredentialsDetail, error) {
	result := oci_dif.DbCredentialsDetail{}

	if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
		tmp := userName.(string)
		result.UserName = &tmp
	}

	if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
		tmp := secretId.(string)
		result.SecretId = &tmp
	}

	if userType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_type")); ok {
		tmp := userType.(string)
		result.UserType = &tmp
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToDataflowArtifactsDetail(fieldKeyFormat string) (oci_dif.DataflowArtifactsDetail, error) {
	result := oci_dif.DataflowArtifactsDetail{}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	if execute, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "execute")); ok {
		tmp := execute.(string)
		result.Execute = &tmp
	}

	if archiveUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "archive_uri")); ok {
		tmp := archiveUri.(string)
		result.ArchiveUri = &tmp
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToGgcsArtifactsDetail(fieldKeyFormat string) (oci_dif.GgcsArtifactsDetail, error) {
	result := oci_dif.GgcsArtifactsDetail{}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	if objectStoragePath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "artifact_object_storage_path")); ok {
		tmp := objectStoragePath.(string)
		result.ArtifactObjectStoragePath = &tmp
	}
	if ggcsUsers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "users")); ok {
		interfaces := ggcsUsers.([]interface{})
		tmp := make([]oci_dif.GgcsUserDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "users"), stateDataIndex)
			converted, err := s.mapToGgcsUserDetail(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "users")) {
			result.Users = tmp
		}
	}

	if ggcsSources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sources")); ok {
		interfaces := ggcsSources.([]interface{})
		tmp := make([]oci_dif.GgcsSourceDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "sources"), stateDataIndex)
			converted, err := s.mapToGgcsSourceDetail(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sources")) {
			result.Sources = tmp
		}
	}

	if ggcsTargets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "targets")); ok {
		interfaces := ggcsTargets.([]interface{})
		tmp := make([]oci_dif.GgcsTargetDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "targets"), stateDataIndex)
			converted, err := s.mapToGgcsTargetDetail(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "targets")) {
			result.Targets = tmp
		}
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToGgcsUserDetail(fieldKeyFormat string) (oci_dif.GgcsUserDetail, error) {
	result := oci_dif.GgcsUserDetail{}

	if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
		tmp := userName.(string)
		result.UserName = &tmp
	}

	if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
		tmp := secretId.(string)
		result.SecretId = &tmp
	}

	if userType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_type")); ok {
		tmp := userType.(string)
		result.UserType = &tmp
	}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		tmp := oci_dif.WorkflowActionEnum(action.(string))
		result.Action = tmp
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToGgcsSourceDetail(fieldKeyFormat string) (oci_dif.GgcsSourceDetail, error) {
	result := oci_dif.GgcsSourceDetail{}

	if sourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_id")); ok {
		tmp := sourceId.(string)
		result.SourceId = &tmp
	}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		tmp := oci_dif.WorkflowActionEnum(action.(string))
		result.Action = tmp
	}

	if targetUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_uri")); ok {
		tmp := targetUri.(string)
		result.TargetUri = &tmp
	}

	if targetConnectionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_connection_name")); ok {
		tmp := targetConnectionName.(string)
		result.TargetConnectionName = &tmp
	}

	if shouldStartSourceOperations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_start_source_operations")); ok {
		tmp := shouldStartSourceOperations.(bool)
		result.ShouldStartSourceOperations = &tmp
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToGgcsTargetDetail(fieldKeyFormat string) (oci_dif.GgcsTargetDetail, error) {
	result := oci_dif.GgcsTargetDetail{}

	if targetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_id")); ok {
		tmp := targetId.(string)
		result.TargetId = &tmp
	}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		tmp := oci_dif.WorkflowActionEnum(action.(string))
		result.Action = tmp
	}

	if sourceUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_uri")); ok {
		tmp := sourceUri.(string)
		result.SourceUri = &tmp
	}

	if sourceConnectionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_connection_name")); ok {
		tmp := sourceConnectionName.(string)
		result.SourceConnectionName = &tmp
	}

	if shouldStartTargetOperations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_start_target_operations")); ok {
		tmp := shouldStartTargetOperations.(bool)
		result.ShouldStartTargetOperations = &tmp
	}

	return result, nil
}

func AdbDetailToMap(obj oci_dif.AdbDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminPasswordId != nil {
		result["admin_password_id"] = string(*obj.AdminPasswordId)
	}

	if obj.DataStorageSizeInTBs != nil {
		result["data_storage_size_in_tbs"] = int(*obj.DataStorageSizeInTBs)
	}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	result["db_workload"] = string(obj.DbWorkload)

	if obj.Ecpu != nil {
		result["ecpu"] = int(*obj.Ecpu)
	}

	if obj.InstanceId != nil {
		result["instance_id"] = string(*obj.InstanceId)
	}

	if obj.IsMtlsConnectionRequired != nil {
		result["is_mtls_connection_required"] = bool(*obj.IsMtlsConnectionRequired)
	}

	if obj.IsPublic != nil {
		result["is_public"] = bool(*obj.IsPublic)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.ToolsPublicAccess != nil {
		result["tools_public_access"] = string(*obj.ToolsPublicAccess)
	}

	return result
}

func AdditionalDetailsToMap(obj *oci_dif.AdditionalDetails) map[string]interface{} {
	result := map[string]interface{}{}

	assignedConnections := []interface{}{}
	for _, item := range obj.AssignedConnections {
		assignedConnections = append(assignedConnections, AssignedConnectionDetailsToMap(item))
	}
	result["assigned_connections"] = assignedConnections

	endpointDetails := []interface{}{}
	for _, item := range obj.EndpointDetails {
		endpointDetails = append(endpointDetails, EndpointAdditionalToMap(item))
	}
	result["endpoint_details"] = endpointDetails

	if obj.ModelId != nil {
		result["model_id"] = string(*obj.ModelId)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.OciRegion != nil {
		result["oci_region"] = string(*obj.OciRegion)
	}

	if obj.PrivateEndpointId != nil {
		result["private_endpoint_id"] = string(*obj.PrivateEndpointId)
	}

	return result
}

func AssignedConnectionDetailsToMap(obj oci_dif.AssignedConnectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectionId != nil {
		result["connection_id"] = string(*obj.ConnectionId)
	}

	if obj.ConnectionName != nil {
		result["connection_name"] = string(*obj.ConnectionName)
	}

	if obj.RequestedBy != nil {
		result["requested_by"] = string(*obj.RequestedBy)
	}

	return result
}

func (s *DifStackResourceCrud) mapToDataflowConnectionDetails(fieldKeyFormat string) (oci_dif.DataflowConnectionDetails, error) {
	result := oci_dif.DataflowConnectionDetails{}

	if difDependencies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dif_dependencies")); ok {
		interfaces := difDependencies.([]interface{})
		tmp := make([]oci_dif.DifDependencyDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "dif_dependencies"), stateDataIndex)
			converted, err := s.mapToDifDependencyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "dif_dependencies")) {
			result.DifDependencies = tmp
		}
	}

	if domainNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_names")); ok {
		set := domainNames.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "domain_names")) {
			result.DomainNames = tmp
		}
	}

	return result, nil
}

func DataflowConnectionDetailsToMap(obj *oci_dif.DataflowConnectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	difDependencies := []interface{}{}
	for _, item := range obj.DifDependencies {
		difDependencies = append(difDependencies, DifDependencyDetailsToMap(item))
	}
	result["dif_dependencies"] = difDependencies

	result["domain_names"] = obj.DomainNames

	return result
}

func (s *DifStackResourceCrud) mapToDataflowConnections(fieldKeyFormat string) (oci_dif.DataflowConnections, error) {
	result := oci_dif.DataflowConnections{}

	if connectionDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_details")); ok {
		if tmpList := connectionDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_details"), 0)
			tmp, err := s.mapToDataflowConnectionDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert connection_details, encountered error: %v", err)
			}
			result.ConnectionDetails = &tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToUpdateDataflowConnections(fieldKeyFormat string) (oci_dif.UpdateDataflowConnections, error) {
	result := oci_dif.UpdateDataflowConnections{}

	if connectionDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_details")); ok {
		if tmpList := connectionDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_details"), 0)
			tmp, err := s.mapToDataflowConnectionDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert connection_details, encountered error: %v", err)
			}
			result.ConnectionDetails = &tmp
		}
	}

	return result, nil
}

func DataflowConnectionsToMap(obj *oci_dif.DataflowConnections) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectionDetails != nil {
		result["connection_details"] = []interface{}{DataflowConnectionDetailsToMap(obj.ConnectionDetails)}
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *DifStackResourceCrud) mapToDataflowDetail(fieldKeyFormat string) (oci_dif.DataflowDetail, error) {
	result := oci_dif.DataflowDetail{}

	if connections, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connections")); ok {
		if tmpList := connections.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connections"), 0)
			tmp, err := s.mapToDataflowConnections(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert connections, encountered error: %v", err)
			}
			result.Connections = &tmp
		}
	}

	if driverShape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "driver_shape")); ok {
		tmp := driverShape.(string)
		result.DriverShape = &tmp
	}

	if driverShapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "driver_shape_config")); ok {
		if tmpList := driverShapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "driver_shape_config"), 0)
			tmp, err := s.mapToShapeConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert driver_shape_config, encountered error: %v", err)
			}
			result.DriverShapeConfig = &tmp
		}
	}

	if executorShape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "executor_shape")); ok {
		tmp := executorShape.(string)
		result.ExecutorShape = &tmp
	}

	if executorShapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "executor_shape_config")); ok {
		if tmpList := executorShapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "executor_shape_config"), 0)
			tmp, err := s.mapToShapeConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert executor_shape_config, encountered error: %v", err)
			}
			result.ExecutorShapeConfig = &tmp
		}
	}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	if logBucketInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_bucket_instance_id")); ok {
		tmp := logBucketInstanceId.(string)
		result.LogBucketInstanceId = &tmp
	}

	if numExecutors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "num_executors")); ok {
		tmp := numExecutors.(int)
		result.NumExecutors = &tmp
	}

	if privateEndpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_endpoint_id")); ok {
		tmp := privateEndpointId.(string)
		result.PrivateEndpointId = &tmp
	}

	if sparkVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "spark_version")); ok {
		tmp := sparkVersion.(string)
		result.SparkVersion = &tmp
	}

	if warehouseBucketInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "warehouse_bucket_instance_id")); ok {
		tmp := warehouseBucketInstanceId.(string)
		result.WarehouseBucketInstanceId = &tmp
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToDataflowUpdateDetail(fieldKeyFormat string) (oci_dif.DataflowUpdateDetail, error) {
	result := oci_dif.DataflowUpdateDetail{}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	// Only set changed fields
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "connections")) {
		if connections, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connections")); ok {
			if tmpList := connections.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connections"), 0)
				tmp, err := s.mapToUpdateDataflowConnections(fieldKeyFormatNextLevel)
				if err != nil {
					return result, fmt.Errorf("unable to convert connections, encountered error: %v", err)
				}
				result.Connections = &tmp
			}
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "driver_shape")) {
		if driverShape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "driver_shape")); ok {
			tmp := driverShape.(string)
			result.DriverShape = &tmp
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "driver_shape_config")) {
		if driverShapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "driver_shape_config")); ok {
			if tmpList := driverShapeConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNext := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "driver_shape_config"), 0)
				tmp, err := s.mapToShapeConfigUpdate(fieldKeyFormatNext)
				if err != nil {
					return result, fmt.Errorf("unable to convert driver_shape_config, encountered error: %v", err)
				}
				result.DriverShapeConfig = &tmp
			}
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "executor_shape")) {
		if executorShape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "executor_shape")); ok {
			tmp := executorShape.(string)
			result.ExecutorShape = &tmp
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "executor_shape_config")) {
		if executorShapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "executor_shape_config")); ok {
			if tmpList := executorShapeConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNext := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "executor_shape_config"), 0)
				tmp, err := s.mapToShapeConfigUpdate(fieldKeyFormatNext)
				if err != nil {
					return result, fmt.Errorf("unable to convert executor_shape_config, encountered error: %v", err)
				}
				result.ExecutorShapeConfig = &tmp
			}
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "log_bucket_instance_id")) {
		if logBucketInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_bucket_instance_id")); ok {
			tmp := logBucketInstanceId.(string)
			result.LogBucketInstanceId = &tmp
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "num_executors")) {
		if numExecutors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "num_executors")); ok {
			tmp := numExecutors.(int)
			result.NumExecutors = &tmp
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "private_endpoint_id")) {
		if privateEndpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_endpoint_id")); ok {
			tmp := privateEndpointId.(string)
			result.PrivateEndpointId = &tmp
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "spark_version")) {
		if sparkVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "spark_version")); ok {
			tmp := sparkVersion.(string)
			result.SparkVersion = &tmp
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "warehouse_bucket_instance_id")) {
		if warehouseBucketInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "warehouse_bucket_instance_id")); ok {
			tmp := warehouseBucketInstanceId.(string)
			result.WarehouseBucketInstanceId = &tmp
		}
	}
	return result, nil
}

func DataflowDetailToMap(obj oci_dif.DataflowDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Connections != nil {
		result["connections"] = []interface{}{DataflowConnectionsToMap(obj.Connections)}
	}

	if obj.DriverShape != nil {
		result["driver_shape"] = string(*obj.DriverShape)
	}

	if obj.DriverShapeConfig != nil {
		result["driver_shape_config"] = []interface{}{ShapeConfigToMap(obj.DriverShapeConfig)}
	}

	if obj.ExecutorShape != nil {
		result["executor_shape"] = string(*obj.ExecutorShape)
	}

	if obj.ExecutorShapeConfig != nil {
		result["executor_shape_config"] = []interface{}{ShapeConfigToMap(obj.ExecutorShapeConfig)}
	}

	if obj.InstanceId != nil {
		result["instance_id"] = string(*obj.InstanceId)
	}

	if obj.LogBucketInstanceId != nil {
		result["log_bucket_instance_id"] = string(*obj.LogBucketInstanceId)
	}

	if obj.NumExecutors != nil {
		result["num_executors"] = int(*obj.NumExecutors)
	}

	if obj.PrivateEndpointId != nil {
		result["private_endpoint_id"] = string(*obj.PrivateEndpointId)
	}

	if obj.SparkVersion != nil {
		result["spark_version"] = string(*obj.SparkVersion)
	}

	if obj.WarehouseBucketInstanceId != nil {
		result["warehouse_bucket_instance_id"] = string(*obj.WarehouseBucketInstanceId)
	}

	return result
}

func (s *DifStackResourceCrud) mapToDifDependencyDetails(fieldKeyFormat string) (oci_dif.DifDependencyDetails, error) {
	result := oci_dif.DifDependencyDetails{}

	if serviceInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_instance_id")); ok {
		tmp := serviceInstanceId.(string)
		result.ServiceInstanceId = &tmp
	}

	if serviceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_type")); ok {
		tmp := serviceType.(string)
		result.ServiceType = &tmp
	}

	return result, nil
}

func DifDependencyDetailsToMap(obj oci_dif.DifDependencyDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ServiceInstanceId != nil {
		result["service_instance_id"] = string(*obj.ServiceInstanceId)
	}

	if obj.ServiceType != nil {
		result["service_type"] = string(*obj.ServiceType)
	}

	return result
}

func EndpointAdditionalToMap(obj oci_dif.EndpointAdditional) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EndpointId != nil {
		result["endpoint_id"] = string(*obj.EndpointId)
	}

	if obj.EndpointName != nil {
		result["endpoint_name"] = string(*obj.EndpointName)
	}

	return result
}

func (s *DifStackResourceCrud) mapToEndpointDetails(fieldKeyFormat string) (oci_dif.EndpointDetails, error) {
	result := oci_dif.EndpointDetails{}

	if endpointName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "endpoint_name")); ok {
		tmp := endpointName.(string)
		result.EndpointName = &tmp
	}

	if isContentModerationEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_content_moderation_enabled")); ok {
		tmp := isContentModerationEnabled.(bool)
		result.IsContentModerationEnabled = &tmp
	}

	return result, nil
}

func EndpointDetailsToMap(obj oci_dif.EndpointDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EndpointName != nil {
		result["endpoint_name"] = string(*obj.EndpointName)
	}

	if obj.IsContentModerationEnabled != nil {
		result["is_content_moderation_enabled"] = bool(*obj.IsContentModerationEnabled)
	}

	return result
}

func (s *DifStackResourceCrud) mapToGenAiDetail(fieldKeyFormat string) (oci_dif.GenAiDetail, error) {
	result := oci_dif.GenAiDetail{}

	if baseModel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_model")); ok {
		tmp := baseModel.(string)
		result.BaseModel = &tmp
	}

	if clusterType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cluster_type")); ok {
		result.ClusterType = oci_dif.ClusterTypeEnum(clusterType.(string))
	}

	if endpoints, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "endpoints")); ok {
		interfaces := endpoints.([]interface{})
		tmp := make([]oci_dif.EndpointDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "endpoints"), stateDataIndex)
			converted, err := s.mapToEndpointDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "endpoints")) {
			result.Endpoints = tmp
		}
	}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	if ociRegion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oci_region")); ok {
		tmp := ociRegion.(string)
		result.OciRegion = &tmp
	}

	if unitCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit_count")); ok {
		tmp := unitCount.(int)
		result.UnitCount = &tmp
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToGenAiUpdateDetail(fieldKeyFormat string) (oci_dif.GenAiUpdateDetail, error) {
	result := oci_dif.GenAiUpdateDetail{}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	// Only set changed fields
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "unit_count")) {
		if unitCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit_count")); ok {
			tmp := unitCount.(int)
			result.UnitCount = &tmp
		}
	}

	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "endpoints")) {
		if endpoints, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "endpoints")); ok {
			interfaces := endpoints.([]interface{})
			tmp := make([]oci_dif.EndpointDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "endpoints"), stateDataIndex)
				converted, err := s.mapToEndpointDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return result, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 {
				result.Endpoints = tmp
			}
		}
	}

	return result, nil
}

func GenAiDetailToMap(obj oci_dif.GenAiDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseModel != nil {
		result["base_model"] = string(*obj.BaseModel)
	}

	result["cluster_type"] = string(obj.ClusterType)

	endpoints := []interface{}{}
	for _, item := range obj.Endpoints {
		endpoints = append(endpoints, EndpointDetailsToMap(item))
	}
	result["endpoints"] = endpoints

	if obj.InstanceId != nil {
		result["instance_id"] = string(*obj.InstanceId)
	}

	if obj.OciRegion != nil {
		result["oci_region"] = string(*obj.OciRegion)
	}

	if obj.UnitCount != nil {
		result["unit_count"] = int(*obj.UnitCount)
	}

	return result
}

func (s *DifStackResourceCrud) mapToGgcsConnectionDetails(fieldKeyFormat string) (oci_dif.GgcsConnectionDetails, error) {
	result := oci_dif.GgcsConnectionDetails{}

	if connectionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_id")); ok {
		tmp := connectionId.(string)
		result.ConnectionId = &tmp
	}

	if connectionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_name")); ok {
		tmp := connectionName.(string)
		result.ConnectionName = &tmp
	}

	if difDependencies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dif_dependencies")); ok {
		interfaces := difDependencies.([]interface{})
		tmp := make([]oci_dif.DifDependencyDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "dif_dependencies"), stateDataIndex)
			converted, err := s.mapToDifDependencyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "dif_dependencies")) {
			result.DifDependencies = tmp
		}
	}

	if ggAdminSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "gg_admin_secret_id")); ok {
		tmp := ggAdminSecretId.(string)
		result.GgAdminSecretId = &tmp
	}

	return result, nil
}

func GgcsConnectionDetailsToMap(obj oci_dif.GgcsConnectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectionId != nil {
		result["connection_id"] = string(*obj.ConnectionId)
	}

	if obj.ConnectionName != nil {
		result["connection_name"] = string(*obj.ConnectionName)
	}

	difDependencies := []interface{}{}
	for _, item := range obj.DifDependencies {
		difDependencies = append(difDependencies, DifDependencyDetailsToMap(item))
	}
	result["dif_dependencies"] = difDependencies

	if obj.GgAdminSecretId != nil {
		result["gg_admin_secret_id"] = string(*obj.GgAdminSecretId)
	}

	return result
}

func (s *DifStackResourceCrud) mapToGgcsDetail(fieldKeyFormat string) (oci_dif.GgcsDetail, error) {
	result := oci_dif.GgcsDetail{}

	if connections, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connections")); ok {
		interfaces := connections.([]interface{})
		tmp := make([]oci_dif.GgcsConnectionDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connections"), stateDataIndex)
			converted, err := s.mapToGgcsConnectionDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "connections")) {
			result.Connections = tmp
		}
	}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	if ocpu, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpu")); ok {
		tmp := ocpu.(int)
		result.Ocpu = &tmp
	}

	if oggVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ogg_version")); ok {
		tmp := oggVersion.(string)
		result.OggVersion = &tmp
	}

	if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
		tmp := passwordSecretId.(string)
		result.PasswordSecretId = &tmp
	}

	if publicSubnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "public_subnet_id")); ok {
		tmp := publicSubnetId.(string)
		result.PublicSubnetId = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToGgcsUpdateDetail(fieldKeyFormat string) (oci_dif.GgcsUpdateDetail, error) {
	result := oci_dif.GgcsUpdateDetail{}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	// Only changed fields
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "connections")) {
		if connections, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connections")); ok {
			interfaces := connections.([]interface{})
			tmp := make([]oci_dif.GgcsConnectionDetails, len(interfaces))
			for i := range interfaces {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connections"), i)
				converted, err := s.mapToGgcsConnectionDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return result, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 {
				result.Connections = tmp
			}
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ocpu")) {
		if ocpu, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpu")); ok {
			tmp := ocpu.(int)
			result.Ocpu = &tmp
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "public_subnet_id")) {
		if publicSubnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "public_subnet_id")); ok {
			tmp := publicSubnetId.(string)
			result.PublicSubnetId = &tmp
		}
	}

	return result, nil
}

func GgcsDetailToMap(obj oci_dif.GgcsDetail) map[string]interface{} {
	result := map[string]interface{}{}

	connections := []interface{}{}
	for _, item := range obj.Connections {
		connections = append(connections, GgcsConnectionDetailsToMap(item))
	}
	result["connections"] = connections

	if obj.InstanceId != nil {
		result["instance_id"] = string(*obj.InstanceId)
	}

	if obj.Ocpu != nil {
		result["ocpu"] = int(*obj.Ocpu)
	}

	if obj.OggVersion != nil {
		result["ogg_version"] = string(*obj.OggVersion)
	}

	if obj.PasswordSecretId != nil {
		result["password_secret_id"] = string(*obj.PasswordSecretId)
	}

	if obj.PublicSubnetId != nil {
		result["public_subnet_id"] = string(*obj.PublicSubnetId)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *DifStackResourceCrud) mapToObjectStorageDetail(fieldKeyFormat string) (oci_dif.ObjectStorageDetail, error) {
	result := oci_dif.ObjectStorageDetail{}

	if autoTiering, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_tiering")); ok {
		result.AutoTiering = oci_dif.AutoTieringEnum(autoTiering.(string))
	}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}

	if objectVersioning, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_versioning")); ok {
		result.ObjectVersioning = oci_dif.ObjectVersioningEnum(objectVersioning.(string))
	}

	if storageTier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_tier")); ok {
		result.StorageTier = oci_dif.StorageTierEnum(storageTier.(string))
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToObjectStorageUpdateDetail(fieldKeyFormat string) (oci_dif.ObjectStorageUpdateDetail, error) {
	result := oci_dif.ObjectStorageUpdateDetail{}

	if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
		tmp := instanceId.(string)
		result.InstanceId = &tmp
	}
	// Only changed fields
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "auto_tiering")) {
		if autoTiering, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_tiering")); ok {
			result.AutoTiering = oci_dif.AutoTieringEnum(autoTiering.(string))
		}
	}
	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "object_versioning")) {
		if objectVersioning, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_versioning")); ok {
			result.ObjectVersioning = oci_dif.ObjectVersioningEnum(objectVersioning.(string))
		}
	}
	return result, nil
}

func ObjectStorageDetailToMap(obj oci_dif.ObjectStorageDetail) map[string]interface{} {
	result := map[string]interface{}{}

	result["auto_tiering"] = string(obj.AutoTiering)

	if obj.InstanceId != nil {
		result["instance_id"] = string(*obj.InstanceId)
	}

	result["object_versioning"] = string(obj.ObjectVersioning)

	result["storage_tier"] = string(obj.StorageTier)

	return result
}

func ServiceDetailResponseToMap(obj oci_dif.ServiceDetailResponse) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdditionalDetails != nil {
		result["additional_details"] = []interface{}{AdditionalDetailsToMap(obj.AdditionalDetails)}
	}

	if obj.CurrentArtifactPath != nil {
		result["current_artifact_path"] = string(*obj.CurrentArtifactPath)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.InstanceId != nil {
		result["instance_id"] = string(*obj.InstanceId)
	}

	if obj.ServiceId != nil {
		result["service_id"] = string(*obj.ServiceId)
	}

	if obj.ServiceType != nil {
		result["service_type"] = string(*obj.ServiceType)
	}

	if obj.ServiceUrl != nil {
		result["service_url"] = string(*obj.ServiceUrl)
	}

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	return result
}

func (s *DifStackResourceCrud) mapToShapeConfig(fieldKeyFormat string) (oci_dif.ShapeConfig, error) {
	result := oci_dif.ShapeConfig{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := memoryInGBs.(int)
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := ocpus.(int)
		result.Ocpus = &tmp
	}

	return result, nil
}

func (s *DifStackResourceCrud) mapToShapeConfigUpdate(fieldKeyFormat string) (oci_dif.ShapeConfig, error) {
	result := oci_dif.ShapeConfig{}

	if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")) || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ocpus")) {
		if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
			tmp := memoryInGBs.(int)
			result.MemoryInGBs = &tmp
		}
		if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
			tmp := ocpus.(int)
			result.Ocpus = &tmp
		}
	}

	return result, nil
}

func ShapeConfigToMap(obj *oci_dif.ShapeConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = int(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = int(*obj.Ocpus)
	}

	return result
}

func StackSummaryToMap(obj oci_dif.StackSummary) map[string]interface{} {
	result := map[string]interface{}{}

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

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

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

func (s *DifStackResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_dif.ChangeStackCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.StackId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif")

	response, err := s.Client.ChangeStackCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		log.Printf("[ERROR] ChangeStackCompartment failed: stackId=%s error=%v", idTmp, err)
		return err
	}
	var _chgWr, _chgOpc string
	if response.OpcWorkRequestId != nil {
		_chgWr = *response.OpcWorkRequestId
	}
	if response.OpcRequestId != nil {
		_chgOpc = *response.OpcRequestId
	}
	log.Printf("[INFO] ChangeStackCompartment initiated: stackId=%s workRequestId=%s opcRequestId=%s", idTmp, _chgWr, _chgOpc)

	workId := response.OpcWorkRequestId
	return s.getStackFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dif"), "dataintelligenceentity", oci_dif.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
