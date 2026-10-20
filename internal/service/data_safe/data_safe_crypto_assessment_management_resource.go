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

func DataSafeCryptoAssessmentManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataSafeCryptoAssessmentManagementWithContext,
		ReadContext:   readDataSafeCryptoAssessmentManagementWithContext,
		UpdateContext: updateDataSafeCryptoAssessmentManagementWithContext,
		DeleteContext: deleteDataSafeCryptoAssessmentManagementWithContext,
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
				Optional: true,
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

func createDataSafeCryptoAssessmentManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	cryptoAssessmentId, ok := d.GetOk("crypto_assessment_id")
	if !ok || cryptoAssessmentId.(string) == "" {
		return tfresource.HandleDiagError(m, fmt.Errorf("crypto_assessment_id must be set"))
	}
	d.SetId(cryptoAssessmentId.(string))

	if err := sync.GetWithContext(ctx); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return updateDataSafeCryptoAssessmentManagementWithContext(ctx, d, m)
}

func readDataSafeCryptoAssessmentManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDataSafeCryptoAssessmentManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDataSafeCryptoAssessmentManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type DataSafeCryptoAssessmentManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.CryptoAssessment
	DisableNotFoundRetries bool
}

func (s *DataSafeCryptoAssessmentManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeCryptoAssessmentManagementResourceCrud) GetWithContext(ctx context.Context) error {
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

func (s *DataSafeCryptoAssessmentManagementResourceCrud) UpdateWithContext(ctx context.Context) error {
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

func (s *DataSafeCryptoAssessmentManagementResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_data_safe.DeleteCryptoAssessmentRequest{}

	tmp := s.D.Id()
	request.CryptoAssessmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteCryptoAssessment(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, delWorkRequestErr := cryptoAssessmentManagementWaitForWorkRequest(ctx, workId, "cryptoassessment",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeCryptoAssessmentManagementResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
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

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastAssessed != nil {
		s.D.Set("time_last_assessed", s.Res.TimeLastAssessed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Type != "" {
		s.D.Set("type", s.Res.Type)
	}

	if s.Res.CryptoProvider != nil {
		s.D.Set("crypto_provider", *s.Res.CryptoProvider)
	}

	if s.Res.CryptoPosture != nil {
		s.D.Set("crypto_posture", []interface{}{CryptoPostureToMap(s.Res.CryptoPosture)})
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

	if s.Res.IssueCount != nil {
		s.D.Set("issue_count", *s.Res.IssueCount)
	}

	if string(s.Res.PostureCategory) != "" {
		s.D.Set("posture_category", s.Res.PostureCategory)
	}

	if s.Res.TargetType != "" {
		s.D.Set("target_type", s.Res.TargetType)
	}

	if s.Res.TargetsWithIssuesCount != nil {
		s.D.Set("targets_with_issues_count", *s.Res.TargetsWithIssuesCount)
	}

	if s.Res.TimeLastAssessed != nil {
		s.D.Set("time_last_assessed", s.Res.TimeLastAssessed.String())
	}

	if s.Res.TriggeredBy != "" {
		s.D.Set("triggered_by", s.Res.TriggeredBy)
	}

	return nil
}

func (s *DataSafeCryptoAssessmentManagementResourceCrud) getCryptoAssessmentFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {
	cryptoAssessmentId, err := cryptoAssessmentManagementWaitForWorkRequest(ctx, workId, "cryptoassessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
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

func (s *DataSafeCryptoAssessmentManagementResourceCrud) GetCryptoAssessmentWorkReq(ctx context.Context) error {
	var tmpTrue = true
	listWorkRequestsRequest := oci_data_safe.ListWorkRequestsRequest{
		SortBy:                 oci_data_safe.ListWorkRequestsSortByEnum("ACCEPTEDTIME"),
		SortOrder:              oci_data_safe.ListWorkRequestsSortOrderEnum("ASC"),
		AccessLevel:            oci_data_safe.ListWorkRequestsAccessLevelEnum("ACCESSIBLE"),
		CompartmentIdInSubtree: &tmpTrue,
	}
	tmp := "CREATE_CRYPTO_ASSESSMENT"
	listWorkRequestsRequest.OperationType = &tmp

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		listWorkRequestsRequest.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		listWorkRequestsRequest.TargetDatabaseId = &tmp
	}

	listWorkRequestsResponse, err := s.Client.ListWorkRequests(ctx, listWorkRequestsRequest)
	if err != nil {
		return err
	}

	if listWorkRequestsResponse.Items != nil && len(listWorkRequestsResponse.Items) > 0 {
		workId := listWorkRequestsResponse.Items[0].Id
		if workId != nil {
			return s.getCryptoAssessmentFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
		}
	}

	return fmt.Errorf("no CREATE_CRYPTO_ASSESSMENT work request found")
}

func cryptoAssessmentManagementWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	retryPolicy.ShouldRetryOperation = func(response oci_common.OCIOperationResponse) bool {
		if time.Now().After(stopTime) {
			return false
		}

		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}

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
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, fmt.Errorf("work request did not succeed for crypto assessment, workId: %s", *wId)
	}

	return identifier, nil
}
