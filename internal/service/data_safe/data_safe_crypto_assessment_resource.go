// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

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
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeCryptoAssessmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataSafeCryptoAssessmentWithContext,
		ReadContext:   readDataSafeCryptoAssessmentWithContext,
		UpdateContext: updateDataSafeCryptoAssessmentWithContext,
		DeleteContext: deleteDataSafeCryptoAssessmentWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"crypto_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_assessment_scheduled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"schedule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"crypto_posture": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"backup_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"encrypted_backup_pieces_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"fips_mode_configured": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fips_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_encryption": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"nne": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"are_weak_options_allowed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"encryption_configured": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"fips_mode_configured": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"integrity": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"key_exchange": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"quantum_readiness": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"server_encryption": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"server_integrity": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"tde": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"db_credentials_encryption_observed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"encrypted_tablespaces_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"encryption_configured": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"fips_mode_configured": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"integrity_configured": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"key_cache_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_store_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"master_key_encryption_algorithm": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"master_key_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"quantum_readiness": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"redo_encryption_observed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_master_key_last_rotation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unencrypted_tablespaces_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"wallet_location": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"tls": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"are_weak_cipher_suites_allowed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cipher_suites_configured": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"fips_mode_configured": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_mtls_configured": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"quantum_readiness": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"revocation_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"versions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"wallet_location": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"unencrypted_backup_pieces_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"crypto_provider": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issue_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"posture_category": {
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
			"target_database_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"targets_with_issues_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_assessed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"triggered_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeCryptoAssessmentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	compartment, ok := sync.D.GetOkExists("compartment_id")

	err := tfresource.CreateResourceWithContext(ctx, d, sync)
	if err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	if ok && compartment != *sync.Res.CompartmentId {
		err = sync.updateCompartment(ctx, compartment)
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		tmp := compartment.(string)
		sync.Res.CompartmentId = &tmp
		err := sync.GetWithContext(ctx)
		if err != nil {
			log.Printf("error doing a Get() after compartment update: %v", err)
		}
		err = sync.SetData()
		if err != nil {
			log.Printf("error doing a SetData() after compartment update: %v", err)
		}
	}
	return nil
}

func readDataSafeCryptoAssessmentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDataSafeCryptoAssessmentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDataSafeCryptoAssessmentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataSafeCryptoAssessmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.CryptoAssessment
	DisableNotFoundRetries bool
}

func (s *DataSafeCryptoAssessmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeCryptoAssessmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.CryptoAssessmentLifecycleStateCreating),
	}
}

func (s *DataSafeCryptoAssessmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.CryptoAssessmentLifecycleStateActive),
	}
}

func (s *DataSafeCryptoAssessmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.CryptoAssessmentLifecycleStateDeleting),
	}
}

func (s *DataSafeCryptoAssessmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.CryptoAssessmentLifecycleStateDeleted),
	}
}

func (s *DataSafeCryptoAssessmentResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_data_safe.UpdateCryptoAssessmentRequest{}

	if cryptoAssessmentId, ok := s.D.GetOkExists("crypto_assessment_id"); ok {
		tmp := cryptoAssessmentId.(string)
		request.CryptoAssessmentId = &tmp
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

	if isAssessmentScheduled, ok := s.D.GetOkExists("is_assessment_scheduled"); ok {
		tmp := isAssessmentScheduled.(bool)
		request.IsAssessmentScheduled = &tmp
	}

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		tmp := schedule.(string)
		request.Schedule = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateCryptoAssessment(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_data_safe.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(ctx,
		oci_data_safe.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "cryptoassessment") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getCryptoAssessmentFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeCryptoAssessmentResourceCrud) getCryptoAssessmentFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	cryptoAssessmentId, err := cryptoAssessmentWaitForWorkRequest(ctx, workId, "cryptoassessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, cryptoAssessmentId)
		_, cancelErr := s.Client.CancelWorkRequest(ctx,
			oci_data_safe.CancelWorkRequestRequest{
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
	s.D.SetId(*cryptoAssessmentId)

	return s.GetWithContext(ctx)
}

func cryptoAssessmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func cryptoAssessmentWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = cryptoAssessmentWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
			string(oci_data_safe.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
			string(oci_data_safe.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_data_safe.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeCryptoAssessmentWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeCryptoAssessmentWorkRequest(ctx context.Context, client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_data_safe.ListWorkRequestErrorsRequest{
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

func (s *DataSafeCryptoAssessmentResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetCryptoAssessmentRequest{}

	tmp := s.D.Id()
	request.CryptoAssessmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetCryptoAssessment(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CryptoAssessment
	return nil
}

func (s *DataSafeCryptoAssessmentResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateCryptoAssessmentRequest{}

	tmp := s.D.Id()
	request.CryptoAssessmentId = &tmp

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

	if isAssessmentScheduled, ok := s.D.GetOkExists("is_assessment_scheduled"); ok {
		tmp := isAssessmentScheduled.(bool)
		request.IsAssessmentScheduled = &tmp
	}

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		tmp := schedule.(string)
		request.Schedule = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateCryptoAssessment(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getCryptoAssessmentFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeCryptoAssessmentResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_data_safe.DeleteCryptoAssessmentRequest{}

	tmp := s.D.Id()
	request.CryptoAssessmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteCryptoAssessment(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := cryptoAssessmentWaitForWorkRequest(ctx, workId, "cryptoassessment",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeCryptoAssessmentResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CryptoPosture != nil {
		s.D.Set("crypto_posture", []interface{}{CryptoPostureToMap(s.Res.CryptoPosture)})
	} else {
		s.D.Set("crypto_posture", nil)
	}

	if s.Res.CryptoProvider != nil {
		s.D.Set("crypto_provider", *s.Res.CryptoProvider)
	}

	if s.Res.DatabaseArchitecture != nil {
		s.D.Set("database_architecture", *s.Res.DatabaseArchitecture)
	}

	if s.Res.DatabaseName != nil {
		s.D.Set("database_name", *s.Res.DatabaseName)
	}

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAssessmentScheduled != nil {
		s.D.Set("is_assessment_scheduled", *s.Res.IsAssessmentScheduled)
	}

	if s.Res.IssueCount != nil {
		s.D.Set("issue_count", *s.Res.IssueCount)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("posture_category", s.Res.PostureCategory)

	if s.Res.Schedule != nil {
		s.D.Set("schedule", *s.Res.Schedule)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetDatabaseGroupId != nil {
		s.D.Set("target_database_group_id", *s.Res.TargetDatabaseGroupId)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	s.D.Set("target_type", s.Res.TargetType)

	if s.Res.TargetsWithIssuesCount != nil {
		s.D.Set("targets_with_issues_count", *s.Res.TargetsWithIssuesCount)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastAssessed != nil {
		s.D.Set("time_last_assessed", s.Res.TimeLastAssessed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("triggered_by", s.Res.TriggeredBy)

	s.D.Set("type", s.Res.Type)

	return nil
}

func CryptoAssessmentSummaryToMap(obj oci_data_safe.CryptoAssessmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatabaseVersion != nil {
		result["database_version"] = string(*obj.DatabaseVersion)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsAssessmentScheduled != nil {
		result["is_assessment_scheduled"] = bool(*obj.IsAssessmentScheduled)
	}

	if obj.IssueCount != nil {
		result["issue_count"] = int(*obj.IssueCount)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["posture_category"] = string(obj.PostureCategory)

	if obj.Schedule != nil {
		result["schedule"] = string(*obj.Schedule)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetDatabaseGroupId != nil {
		result["target_database_group_id"] = string(*obj.TargetDatabaseGroupId)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	result["target_type"] = string(obj.TargetType)

	if obj.TargetsWithIssuesCount != nil {
		result["targets_with_issues_count"] = int(*obj.TargetsWithIssuesCount)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastAssessed != nil {
		result["time_last_assessed"] = obj.TimeLastAssessed.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["triggered_by"] = string(obj.TriggeredBy)

	result["type"] = string(obj.Type)

	return result
}

func CryptoNnePostureToMap(obj *oci_data_safe.CryptoNnePosture) map[string]interface{} {
	result := map[string]interface{}{}

	result["are_weak_options_allowed"] = string(obj.AreWeakOptionsAllowed)

	result["encryption_configured"] = obj.EncryptionConfigured

	result["fips_mode_configured"] = string(obj.FipsModeConfigured)

	result["integrity"] = obj.Integrity

	if obj.KeyExchange != nil {
		result["key_exchange"] = string(*obj.KeyExchange)
	}

	result["quantum_readiness"] = string(obj.QuantumReadiness)

	result["server_encryption"] = string(obj.ServerEncryption)

	result["server_integrity"] = obj.ServerIntegrity

	result["status"] = string(obj.Status)

	return result
}

func CryptoPostureToMap(obj *oci_data_safe.CryptoPosture) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupStatus != nil {
		result["backup_status"] = string(*obj.BackupStatus)
	}

	if obj.EncryptedBackupPiecesCount != nil {
		result["encrypted_backup_pieces_count"] = int(*obj.EncryptedBackupPiecesCount)
	}

	result["fips_mode_configured"] = string(obj.FipsModeConfigured)

	result["fips_status"] = string(obj.FipsStatus)

	result["network_encryption"] = obj.NetworkEncryption

	if obj.Nne != nil {
		result["nne"] = []interface{}{CryptoNnePostureToMap(obj.Nne)}
	}

	if obj.Tde != nil {
		result["tde"] = []interface{}{CryptoTdePostureToMap(obj.Tde)}
	}

	if obj.Tls != nil {
		result["tls"] = []interface{}{CryptoTlsPostureToMap(obj.Tls)}
	}

	if obj.UnencryptedBackupPiecesCount != nil {
		result["unencrypted_backup_pieces_count"] = int(*obj.UnencryptedBackupPiecesCount)
	}

	return result
}

func CryptoTdePostureToMap(obj *oci_data_safe.CryptoTdePosture) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbCredentialsEncryptionObserved != nil {
		result["db_credentials_encryption_observed"] = string(*obj.DbCredentialsEncryptionObserved)
	}

	if obj.EncryptedTablespacesCount != nil {
		result["encrypted_tablespaces_count"] = int(*obj.EncryptedTablespacesCount)
	}

	result["encryption_configured"] = obj.EncryptionConfigured

	result["fips_mode_configured"] = string(obj.FipsModeConfigured)

	result["integrity_configured"] = obj.IntegrityConfigured

	result["key_cache_status"] = string(obj.KeyCacheStatus)

	result["key_store_type"] = string(obj.KeyStoreType)

	if obj.MasterKeyEncryptionAlgorithm != nil {
		result["master_key_encryption_algorithm"] = string(*obj.MasterKeyEncryptionAlgorithm)
	}

	if obj.MasterKeyId != nil {
		result["master_key_id"] = string(*obj.MasterKeyId)
	}

	result["quantum_readiness"] = string(obj.QuantumReadiness)

	if obj.RedoEncryptionObserved != nil {
		result["redo_encryption_observed"] = string(*obj.RedoEncryptionObserved)
	}

	result["status"] = string(obj.Status)

	if obj.TimeMasterKeyLastRotation != nil {
		result["time_master_key_last_rotation"] = obj.TimeMasterKeyLastRotation.String()
	}

	if obj.UnencryptedTablespacesCount != nil {
		result["unencrypted_tablespaces_count"] = int(*obj.UnencryptedTablespacesCount)
	}

	if obj.WalletLocation != nil {
		result["wallet_location"] = string(*obj.WalletLocation)
	}

	return result
}

func CryptoTlsPostureToMap(obj *oci_data_safe.CryptoTlsPosture) map[string]interface{} {
	result := map[string]interface{}{}

	result["are_weak_cipher_suites_allowed"] = string(obj.AreWeakCipherSuitesAllowed)

	result["cipher_suites_configured"] = obj.CipherSuitesConfigured

	result["fips_mode_configured"] = string(obj.FipsModeConfigured)

	result["is_mtls_configured"] = string(obj.IsMtlsConfigured)

	result["quantum_readiness"] = string(obj.QuantumReadiness)

	if obj.RevocationMode != nil {
		result["revocation_mode"] = string(*obj.RevocationMode)
	}

	result["status"] = string(obj.Status)

	result["versions"] = obj.Versions

	if obj.WalletLocation != nil {
		result["wallet_location"] = string(*obj.WalletLocation)
	}

	return result
}

func (s *DataSafeCryptoAssessmentResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeCryptoAssessmentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.CryptoAssessmentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeCryptoAssessmentCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getCryptoAssessmentFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
