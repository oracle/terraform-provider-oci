// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_migrations

import (
	"context"
	"fmt"
	"log"
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

func CloudMigrationsMigrationPlanResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudMigrationsMigrationPlan,
		Read:     readCloudMigrationsMigrationPlan,
		Update:   updateCloudMigrationsMigrationPlan,
		Delete:   deleteCloudMigrationsMigrationPlan,
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
			"migration_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
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
			"source_migration_plan_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"strategies": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"resource_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"strategy_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AS_IS",
								"AVERAGE",
								"PEAK",
								"PERCENTILE",
							}, true),
						},

						// Optional
						"adjustment_multiplier": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"metric_time_window": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metric_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"percentile": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"target_environments": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subnet": {
							Type:     schema.TypeString,
							Required: true,
						},
						"target_environment_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"VM_TARGET_ENV",
							}, true),
						},
						"vcn": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"availability_domain": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},
						"dedicated_vm_host": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fault_domain": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ms_license": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preferred_shape_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"target_compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"calculated_limits": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"migration_plan_stats": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_estimated_cost": {
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
						"vm_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"reference_to_rms_stack": {
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
}

func createCloudMigrationsMigrationPlan(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudMigrationsMigrationPlan(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.ReadResource(sync)
}

func updateCloudMigrationsMigrationPlan(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudMigrationsMigrationPlan(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudMigrationsMigrationPlanResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_migrations.MigrationClient
	Res                    *oci_cloud_migrations.MigrationPlan
	DisableNotFoundRetries bool
}

func (s *CloudMigrationsMigrationPlanResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudMigrationsMigrationPlanResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_migrations.MigrationPlanLifecycleStateCreating),
	}
}

func (s *CloudMigrationsMigrationPlanResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_migrations.MigrationPlanLifecycleStateNeedsAttention),
		string(oci_cloud_migrations.MigrationPlanLifecycleStateActive),
	}
}

func (s *CloudMigrationsMigrationPlanResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_migrations.MigrationPlanLifecycleStateDeleting),
	}
}

func (s *CloudMigrationsMigrationPlanResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_migrations.MigrationPlanLifecycleStateDeleted),
	}
}

func (s *CloudMigrationsMigrationPlanResourceCrud) Create() error {
	request := oci_cloud_migrations.CreateMigrationPlanRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	if migrationId, ok := s.D.GetOkExists("migration_id"); ok {
		tmp := migrationId.(string)
		request.MigrationId = &tmp
	}

	if sourceMigrationPlanId, ok := s.D.GetOkExists("source_migration_plan_id"); ok {
		tmp := sourceMigrationPlanId.(string)
		request.SourceMigrationPlanId = &tmp
	}

	if strategies, ok := s.D.GetOkExists("strategies"); ok {
		interfaces := strategies.([]interface{})
		tmp := make([]oci_cloud_migrations.ResourceAssessmentStrategy, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "strategies", stateDataIndex)
			converted, err := s.mapToResourceAssessmentStrategy(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("strategies") {
			request.Strategies = tmp
		}
	}

	if targetEnvironments, ok := s.D.GetOkExists("target_environments"); ok {
		interfaces := targetEnvironments.([]interface{})
		tmp := make([]oci_cloud_migrations.TargetEnvironment, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_environments", stateDataIndex)
			converted, err := s.mapToTargetEnvironment(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("target_environments") {
			request.TargetEnvironments = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.CreateMigrationPlan(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getMigrationPlanFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations"), oci_cloud_migrations.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *CloudMigrationsMigrationPlanResourceCrud) getMigrationPlanFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_cloud_migrations.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	migrationPlanId, err := migrationPlanWaitForWorkRequest(workId, "migrationPlan",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, migrationPlanId)
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
	s.D.SetId(*migrationPlanId)

	return s.Get()
}

func migrationPlanWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func migrationPlanWaitForWorkRequest(wId *string, entityType string, action oci_cloud_migrations.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_cloud_migrations.MigrationClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "cloud_migrations")
	retryPolicy.ShouldRetryOperation = migrationPlanWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromCloudMigrationsMigrationPlanWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromCloudMigrationsMigrationPlanWorkRequest(client *oci_cloud_migrations.MigrationClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_cloud_migrations.ActionTypeEnum) error {
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

func (s *CloudMigrationsMigrationPlanResourceCrud) Get() error {
	request := oci_cloud_migrations.GetMigrationPlanRequest{}

	tmp := s.D.Id()
	request.MigrationPlanId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.GetMigrationPlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MigrationPlan
	return nil
}

func (s *CloudMigrationsMigrationPlanResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_migrations.UpdateMigrationPlanRequest{}

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
	request.MigrationPlanId = &tmp

	if strategies, ok := s.D.GetOkExists("strategies"); ok {
		interfaces := strategies.([]interface{})
		tmp := make([]oci_cloud_migrations.ResourceAssessmentStrategy, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "strategies", stateDataIndex)
			converted, err := s.mapToResourceAssessmentStrategy(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("strategies") {
			request.Strategies = tmp
		}
	}

	if targetEnvironments, ok := s.D.GetOkExists("target_environments"); ok {
		interfaces := targetEnvironments.([]interface{})
		tmp := make([]oci_cloud_migrations.TargetEnvironment, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_environments", stateDataIndex)
			converted, err := s.mapToTargetEnvironment(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("target_environments") {
			request.TargetEnvironments = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.UpdateMigrationPlan(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMigrationPlanFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations"), oci_cloud_migrations.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *CloudMigrationsMigrationPlanResourceCrud) Delete() error {
	request := oci_cloud_migrations.DeleteMigrationPlanRequest{}

	tmp := s.D.Id()
	request.MigrationPlanId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.DeleteMigrationPlan(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := migrationPlanWaitForWorkRequest(workId, "migrationPlan",
		oci_cloud_migrations.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *CloudMigrationsMigrationPlanResourceCrud) SetData() error {

	s.D.SetId(*s.Res.Id)
	s.D.Set("calculated_limits", s.Res.CalculatedLimits)
	s.D.Set("calculated_limits", s.Res.CalculatedLimits)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MigrationId != nil {
		s.D.Set("migration_id", *s.Res.MigrationId)
	}

	if s.Res.MigrationPlanStats != nil {
		s.D.Set("migration_plan_stats", []interface{}{MigrationPlanStatsToMap(s.Res.MigrationPlanStats)})
	} else {
		s.D.Set("migration_plan_stats", nil)
	}

	if s.Res.ReferenceToRmsStack != nil {
		s.D.Set("reference_to_rms_stack", *s.Res.ReferenceToRmsStack)
	}

	if s.Res.SourceMigrationPlanId != nil {
		s.D.Set("source_migration_plan_id", *s.Res.SourceMigrationPlanId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	strategies := []interface{}{}
	for _, item := range s.Res.Strategies {
		strategies = append(strategies, ResourceAssessmentStrategyToMap(item))
	}
	s.D.Set("strategies", strategies)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	targetEnvironments := []interface{}{}
	for _, item := range s.Res.TargetEnvironments {
		targetEnvironments = append(targetEnvironments, TargetEnvironmentToMap(item))
	}
	s.D.Set("target_environments", targetEnvironments)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ComputeCostEstimationToMap(obj *oci_cloud_migrations.ComputeCostEstimation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GpuCount != nil {
		result["gpu_count"] = float32(*obj.GpuCount)
	}

	if obj.GpuPerHour != nil {
		result["gpu_per_hour"] = float32(*obj.GpuPerHour)
	}

	if obj.GpuPerHourBySubscription != nil {
		result["gpu_per_hour_by_subscription"] = float32(*obj.GpuPerHourBySubscription)
	}

	if obj.MemoryAmountGb != nil {
		result["memory_amount_gb"] = float32(*obj.MemoryAmountGb)
	}

	if obj.MemoryGbPerHour != nil {
		result["memory_gb_per_hour"] = float32(*obj.MemoryGbPerHour)
	}

	if obj.MemoryGbPerHourBySubscription != nil {
		result["memory_gb_per_hour_by_subscription"] = float32(*obj.MemoryGbPerHourBySubscription)
	}

	if obj.OcpuCount != nil {
		result["ocpu_count"] = float32(*obj.OcpuCount)
	}

	if obj.OcpuPerHour != nil {
		result["ocpu_per_hour"] = float32(*obj.OcpuPerHour)
	}

	if obj.OcpuPerHourBySubscription != nil {
		result["ocpu_per_hour_by_subscription"] = float32(*obj.OcpuPerHourBySubscription)
	}

	if obj.TotalPerHour != nil {
		result["total_per_hour"] = float32(*obj.TotalPerHour)
	}

	if obj.TotalPerHourBySubscription != nil {
		result["total_per_hour_by_subscription"] = float32(*obj.TotalPerHourBySubscription)
	}

	return result
}

func CostEstimationToMap(obj *oci_cloud_migrations.CostEstimation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Compute != nil {
		result["compute"] = []interface{}{ComputeCostEstimationToMap(obj.Compute)}
	}

	if obj.CurrencyCode != nil {
		result["currency_code"] = string(*obj.CurrencyCode)
	}

	if obj.OsImage != nil {
		result["os_image"] = []interface{}{OsImageEstimationToMap(obj.OsImage)}
	}

	if obj.Storage != nil {
		result["storage"] = []interface{}{StorageCostEstimationToMap(obj.Storage)}
	}

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	if obj.TotalEstimationPerMonth != nil {
		result["total_estimation_per_month"] = float32(*obj.TotalEstimationPerMonth)
	}

	if obj.TotalEstimationPerMonthBySubscription != nil {
		result["total_estimation_per_month_by_subscription"] = float32(*obj.TotalEstimationPerMonthBySubscription)
	}

	return result
}

func MigrationPlanStatsToMap(obj *oci_cloud_migrations.MigrationPlanStats) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TotalEstimatedCost != nil {
		result["total_estimated_cost"] = []interface{}{CostEstimationToMap(obj.TotalEstimatedCost)}
	}

	if obj.VmCount != nil {
		result["vm_count"] = int(*obj.VmCount)
	}

	return result
}

func MigrationPlanSummaryToMap(obj oci_cloud_migrations.MigrationPlanSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["calculated_limits"] = obj.CalculatedLimits

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
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MigrationId != nil {
		result["migration_id"] = string(*obj.MigrationId)
	}

	if obj.MigrationPlanStats != nil {
		result["migration_plan_stats"] = []interface{}{MigrationPlanStatsToMap(obj.MigrationPlanStats)}
	}

	if obj.ReferenceToRmsStack != nil {
		result["reference_to_rms_stack"] = string(*obj.ReferenceToRmsStack)
	}

	if obj.SourceMigrationPlanId != nil {
		result["source_migration_plan_id"] = string(*obj.SourceMigrationPlanId)
	}

	result["state"] = string(obj.LifecycleState)

	strategies := []interface{}{}
	for _, item := range obj.Strategies {
		strategies = append(strategies, ResourceAssessmentStrategyToMap(item))
	}
	result["strategies"] = strategies

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	targetEnvironments := []interface{}{}
	for _, item := range obj.TargetEnvironments {
		targetEnvironments = append(targetEnvironments, TargetEnvironmentToMap(item))
	}
	result["target_environments"] = targetEnvironments

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func OsImageEstimationToMap(obj *oci_cloud_migrations.OsImageEstimation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TotalPerHour != nil {
		result["total_per_hour"] = float32(*obj.TotalPerHour)
	}

	if obj.TotalPerHourBySubscription != nil {
		result["total_per_hour_by_subscription"] = float32(*obj.TotalPerHourBySubscription)
	}

	return result
}

func (s *CloudMigrationsMigrationPlanResourceCrud) mapToResourceAssessmentStrategy(fieldKeyFormat string) (oci_cloud_migrations.ResourceAssessmentStrategy, error) {
	var baseObject oci_cloud_migrations.ResourceAssessmentStrategy
	//discriminator
	strategyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "strategy_type"))
	var strategyType string
	if ok {
		strategyType = strategyTypeRaw.(string)
	} else {
		strategyType = "" // default value
	}
	switch strings.ToLower(strategyType) {
	case strings.ToLower("AS_IS"):
		details := oci_cloud_migrations.AsIsResourceAssessmentStrategy{}
		if adjustmentMultiplier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "adjustment_multiplier")); ok {
			tmp, ok := adjustmentMultiplier.(float32)
			if !ok {
				tmp = float32(adjustmentMultiplier.(float64))
			}
			details.AdjustmentMultiplier = &tmp
		}
		if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
			details.ResourceType = oci_cloud_migrations.ResourceAssessmentStrategyResourceTypeEnum(resourceType.(string))
		}
		baseObject = details
	case strings.ToLower("AVERAGE"):
		details := oci_cloud_migrations.AverageResourceAssessmentStrategy{}
		if adjustmentMultiplier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "adjustment_multiplier")); ok {
			tmp, ok := adjustmentMultiplier.(float32)
			if !ok {
				tmp = float32(adjustmentMultiplier.(float64))
			}
			details.AdjustmentMultiplier = &tmp
		}
		if metricTimeWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_time_window")); ok {
			details.MetricTimeWindow = oci_cloud_migrations.MetricTimeWindowEnum(metricTimeWindow.(string))
		}
		if metricType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_type")); ok {
			details.MetricType = oci_cloud_migrations.MetricTypeEnum(metricType.(string))
		}
		if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
			details.ResourceType = oci_cloud_migrations.ResourceAssessmentStrategyResourceTypeEnum(resourceType.(string))
		}
		baseObject = details
	case strings.ToLower("PEAK"):
		details := oci_cloud_migrations.PeakResourceAssessmentStrategy{}
		if adjustmentMultiplier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "adjustment_multiplier")); ok {
			tmp, ok := adjustmentMultiplier.(float32)
			if !ok {
				tmp = float32(adjustmentMultiplier.(float64))
			}
			details.AdjustmentMultiplier = &tmp
		}
		if metricTimeWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_time_window")); ok {
			details.MetricTimeWindow = oci_cloud_migrations.MetricTimeWindowEnum(metricTimeWindow.(string))
		}
		if metricType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_type")); ok {
			details.MetricType = oci_cloud_migrations.MetricTypeEnum(metricType.(string))
		}
		if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
			details.ResourceType = oci_cloud_migrations.ResourceAssessmentStrategyResourceTypeEnum(resourceType.(string))
		}
		baseObject = details
	case strings.ToLower("PERCENTILE"):
		details := oci_cloud_migrations.PercentileResourceAssessmentStrategy{}
		if adjustmentMultiplier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "adjustment_multiplier")); ok {
			tmp, ok := adjustmentMultiplier.(float32)
			if !ok {
				tmp = float32(adjustmentMultiplier.(float64))
			}
			details.AdjustmentMultiplier = &tmp
		}
		if metricTimeWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_time_window")); ok {
			details.MetricTimeWindow = oci_cloud_migrations.MetricTimeWindowEnum(metricTimeWindow.(string))
		}
		if percentile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "percentile")); ok {
			details.Percentile = oci_cloud_migrations.PercentileResourceAssessmentStrategyPercentileEnum(percentile.(string))
		}
		if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
			details.ResourceType = oci_cloud_migrations.ResourceAssessmentStrategyResourceTypeEnum(resourceType.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown strategy_type '%v' was specified", strategyType)
	}
	return baseObject, nil
}

func ResourceAssessmentStrategyToMap(obj oci_cloud_migrations.ResourceAssessmentStrategy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_cloud_migrations.AsIsResourceAssessmentStrategy:
		result["strategy_type"] = "AS_IS"

		if v.AdjustmentMultiplier != nil {
			result["adjustment_multiplier"] = float32(*v.AdjustmentMultiplier)
		}

		result["resource_type"] = string(v.ResourceType)
	case oci_cloud_migrations.AverageResourceAssessmentStrategy:
		result["strategy_type"] = "AVERAGE"

		if v.AdjustmentMultiplier != nil {
			result["adjustment_multiplier"] = float32(*v.AdjustmentMultiplier)
		}

		result["metric_time_window"] = string(v.MetricTimeWindow)

		result["metric_type"] = string(v.MetricType)

		result["resource_type"] = string(v.ResourceType)
	case oci_cloud_migrations.PeakResourceAssessmentStrategy:
		result["strategy_type"] = "PEAK"

		if v.AdjustmentMultiplier != nil {
			result["adjustment_multiplier"] = float32(*v.AdjustmentMultiplier)
		}

		result["metric_time_window"] = string(v.MetricTimeWindow)

		result["metric_type"] = string(v.MetricType)

		result["resource_type"] = string(v.ResourceType)
	case oci_cloud_migrations.PercentileResourceAssessmentStrategy:
		result["strategy_type"] = "PERCENTILE"

		if v.AdjustmentMultiplier != nil {
			result["adjustment_multiplier"] = float32(*v.AdjustmentMultiplier)
		}

		result["metric_time_window"] = string(v.MetricTimeWindow)

		result["percentile"] = string(v.Percentile)

		result["resource_type"] = string(v.ResourceType)
	default:
		log.Printf("[WARN] Received 'strategy_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func StorageCostEstimationToMap(obj *oci_cloud_migrations.StorageCostEstimation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TotalGbPerMonth != nil {
		result["total_gb_per_month"] = float32(*obj.TotalGbPerMonth)
	}

	if obj.TotalGbPerMonthBySubscription != nil {
		result["total_gb_per_month_by_subscription"] = float32(*obj.TotalGbPerMonthBySubscription)
	}

	volumes := []interface{}{}
	for _, item := range obj.Volumes {
		volumes = append(volumes, VolumeCostEstimationToMap(item))
	}
	result["volumes"] = volumes

	return result
}

func (s *CloudMigrationsMigrationPlanResourceCrud) mapToTargetEnvironment(fieldKeyFormat string) (oci_cloud_migrations.TargetEnvironment, error) {
	var baseObject oci_cloud_migrations.TargetEnvironment
	//discriminator
	targetEnvironmentTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_environment_type"))
	var targetEnvironmentType string
	if ok {
		targetEnvironmentType = targetEnvironmentTypeRaw.(string)
	} else {
		targetEnvironmentType = "" // default value
	}
	switch strings.ToLower(targetEnvironmentType) {
	case strings.ToLower("VM_TARGET_ENV"):
		details := oci_cloud_migrations.VmTargetEnvironment{}
		if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		if dedicatedVmHost, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dedicated_vm_host")); ok {
			tmp := dedicatedVmHost.(string)
			if tmp != "" {
				details.DedicatedVmHost = &tmp
			}
		}
		if faultDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_domain")); ok {
			tmp := faultDomain.(string)
			if tmp != "" {
				details.FaultDomain = &tmp
			}
		}
		if msLicense, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ms_license")); ok {
			tmp := msLicense.(string)
			details.MsLicense = &tmp
		}
		if preferredShapeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preferred_shape_type")); ok {
			details.PreferredShapeType = oci_cloud_migrations.VmTargetAssetPreferredShapeTypeEnum(preferredShapeType.(string))
		}
		if subnet, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet")); ok {
			tmp := subnet.(string)
			details.Subnet = &tmp
		}
		if vcn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcn")); ok {
			tmp := vcn.(string)
			details.Vcn = &tmp
		}
		if targetCompartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_compartment_id")); ok {
			tmp := targetCompartmentId.(string)
			details.TargetCompartmentId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown target_environment_type '%v' was specified", targetEnvironmentType)
	}
	return baseObject, nil
}

func TargetEnvironmentToMap(obj oci_cloud_migrations.TargetEnvironment) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_cloud_migrations.VmTargetEnvironment:
		result["target_environment_type"] = "VM_TARGET_ENV"

		if v.AvailabilityDomain != nil {
			result["availability_domain"] = string(*v.AvailabilityDomain)
		}

		if v.DedicatedVmHost != nil {
			result["dedicated_vm_host"] = string(*v.DedicatedVmHost)
		}

		if v.FaultDomain != nil {
			result["fault_domain"] = string(*v.FaultDomain)
		}

		if v.MsLicense != nil {
			result["ms_license"] = string(*v.MsLicense)
		}

		result["preferred_shape_type"] = string(v.PreferredShapeType)

		if v.Subnet != nil {
			result["subnet"] = string(*v.Subnet)
		}

		if v.Vcn != nil {
			result["vcn"] = string(*v.Vcn)
		}

		if v.TargetCompartmentId != nil {
			result["target_compartment_id"] = string(*v.TargetCompartmentId)
		}
	default:
		log.Printf("[WARN] Received 'target_environment_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func VolumeCostEstimationToMap(obj oci_cloud_migrations.VolumeCostEstimation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CapacityGb != nil {
		result["capacity_gb"] = float32(*obj.CapacityGb)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.TotalGbPerMonth != nil {
		result["total_gb_per_month"] = float32(*obj.TotalGbPerMonth)
	}

	if obj.TotalGbPerMonthBySubscription != nil {
		result["total_gb_per_month_by_subscription"] = float32(*obj.TotalGbPerMonthBySubscription)
	}

	return result
}

func (s *CloudMigrationsMigrationPlanResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_migrations.ChangeMigrationPlanCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MigrationPlanId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.ChangeMigrationPlanCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMigrationPlanFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations"), oci_cloud_migrations.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
