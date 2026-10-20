// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
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

func DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectWithContext,
		ReadContext:   readDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectWithContext,
		UpdateContext: updateDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectWithContext,
		DeleteContext: deleteDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"processing_chain_object_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subsetting_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subsetting_rule_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Required
			"is_enabled_for_processing": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"approximate_row_count_before_subsetting": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"child_columns": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"child_object_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"child_schema_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"estimated_row_count_after_subsetting": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_enabled_for_processing": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_columns": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"parent_object_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_schema_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"propagation_impact": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subsetting_schema_relation_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SubsettingRuleProcessingChainObjectsCollection
	StableRelationKey      string
	DisableNotFoundRetries bool
}

func (s *DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud) ID() string {
	return GetSubsettingPolicySubsettingRuleProcessingChainObjectCompositeId(s.D.Get("processing_chain_object_key").(string), s.D.Get("subsetting_policy_id").(string), s.D.Get("subsetting_rule_key").(string))
}

func (s *DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud) CreateWithContext(ctx context.Context) error {
	if err := s.captureStableRelationKey(ctx); err != nil {
		return err
	}

	request := oci_data_safe.UpdateProcessingChainObjectRequest{}

	isEnabledForProcessing := s.D.Get("is_enabled_for_processing").(bool)
	request.IsEnabledForProcessing = &isEnabledForProcessing

	if processingChainObjectKey, ok := s.D.GetOkExists("processing_chain_object_key"); ok {
		tmp := processingChainObjectKey.(string)
		request.ProcessingChainObjectKey = &tmp
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if subsettingRuleKey, ok := s.D.GetOkExists("subsetting_rule_key"); ok {
		tmp := subsettingRuleKey.(string)
		request.SubsettingRuleKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateProcessingChainObject(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSubsettingPolicySubsettingRuleProcessingChainObjectFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud) getSubsettingPolicySubsettingRuleProcessingChainObjectFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	subsettingPolicySubsettingRuleProcessingChainObjectId, err := subsettingPolicySubsettingRuleProcessingChainObjectWaitForWorkRequest(ctx, workId, "subsettingpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, subsettingPolicySubsettingRuleProcessingChainObjectId)
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
	// The work request identifies the parent subsetting policy, not the
	// processing-chain object. Keep the resource's composite ID so subsequent
	// reads continue to use the policy, rule, and processing-chain object keys.
	s.D.SetId(GetSubsettingPolicySubsettingRuleProcessingChainObjectCompositeId(
		s.D.Get("processing_chain_object_key").(string),
		s.D.Get("subsetting_policy_id").(string),
		s.D.Get("subsetting_rule_key").(string),
	))

	return s.GetWithContext(ctx)
}

func subsettingPolicySubsettingRuleProcessingChainObjectWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func subsettingPolicySubsettingRuleProcessingChainObjectWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = subsettingPolicySubsettingRuleProcessingChainObjectWorkRequestShouldRetryFunc(timeout)

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
		if dataSafeWorkRequestEntityTypeMatches(res.EntityType, entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectWorkRequest(ctx context.Context, client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingRuleProcessingChainObjectsRequest{}

	if isEnabledForProcessing, ok := s.D.GetOkExists("is_enabled_for_processing"); ok {
		tmp := isEnabledForProcessing.(bool)
		request.IsEnabledForProcessing = &tmp
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if subsettingRuleKey, ok := s.D.GetOkExists("subsetting_rule_key"); ok {
		tmp := subsettingRuleKey.(string)
		request.SubsettingRuleKey = &tmp
	}

	_, subsettingPolicyId, subsettingRuleKey, err := parseSubsettingPolicySubsettingRuleProcessingChainObjectCompositeId(s.D.Id())
	if err == nil {
		request.SubsettingPolicyId = &subsettingPolicyId
		request.SubsettingRuleKey = &subsettingRuleKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListSubsettingRuleProcessingChainObjects(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.SubsettingRuleProcessingChainObjectsCollection
	return nil
}

func (s *DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud) UpdateWithContext(ctx context.Context) error {
	if err := s.captureStableRelationKey(ctx); err != nil {
		return err
	}

	request := oci_data_safe.UpdateProcessingChainObjectRequest{}

	isEnabledForProcessing := s.D.Get("is_enabled_for_processing").(bool)
	request.IsEnabledForProcessing = &isEnabledForProcessing

	if processingChainObjectKey, ok := s.D.GetOkExists("processing_chain_object_key"); ok {
		tmp := processingChainObjectKey.(string)
		request.ProcessingChainObjectKey = &tmp
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if subsettingRuleKey, ok := s.D.GetOkExists("subsetting_rule_key"); ok {
		tmp := subsettingRuleKey.(string)
		request.SubsettingRuleKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateProcessingChainObject(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSubsettingPolicySubsettingRuleProcessingChainObjectFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud) DeleteWithContext(ctx context.Context) error {
	if err := s.captureStableRelationKey(ctx); err != nil {
		return err
	}

	isEnabledForProcessing := false
	request := oci_data_safe.UpdateProcessingChainObjectRequest{}
	request.IsEnabledForProcessing = &isEnabledForProcessing

	processingChainObjectKey := s.D.Get("processing_chain_object_key").(string)
	request.ProcessingChainObjectKey = &processingChainObjectKey
	subsettingPolicyId := s.D.Get("subsetting_policy_id").(string)
	request.SubsettingPolicyId = &subsettingPolicyId
	subsettingRuleKey := s.D.Get("subsetting_rule_key").(string)
	request.SubsettingRuleKey = &subsettingRuleKey
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateProcessingChainObject(ctx, request)
	if err != nil {
		return err
	}

	return s.getSubsettingPolicySubsettingRuleProcessingChainObjectFromWorkRequest(ctx, response.OpcWorkRequestId,
		tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated,
		s.D.Timeout(schema.TimeoutDelete))
}

func (s *DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud) SetData() error {

	processingChainObjectKey, subsettingPolicyId, subsettingRuleKey, err := parseSubsettingPolicySubsettingRuleProcessingChainObjectCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("processing_chain_object_key", &processingChainObjectKey)
		s.D.Set("subsetting_policy_id", &subsettingPolicyId)
		s.D.Set("subsetting_rule_key", &subsettingRuleKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		isResourceItem := item.Key != nil && *item.Key == processingChainObjectKey
		if s.StableRelationKey != "" && item.SubsettingSchemaRelationKey != nil && *item.SubsettingSchemaRelationKey == s.StableRelationKey && item.Key != nil {
			isResourceItem = true
			s.D.Set("processing_chain_object_key", *item.Key)
			s.D.SetId(GetSubsettingPolicySubsettingRuleProcessingChainObjectCompositeId(
				*item.Key,
				s.D.Get("subsetting_policy_id").(string),
				s.D.Get("subsetting_rule_key").(string),
			))
		}
		if isResourceItem && item.IsEnabledForProcessing != nil {
			s.D.Set("is_enabled_for_processing", *item.IsEnabledForProcessing)
		}
		items = append(items, SubsettingRuleProcessingChainObjectSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func (s *DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectResourceCrud) captureStableRelationKey(ctx context.Context) error {
	policyID := s.D.Get("subsetting_policy_id").(string)
	ruleKey := s.D.Get("subsetting_rule_key").(string)
	request := oci_data_safe.ListSubsettingRuleProcessingChainObjectsRequest{
		SubsettingPolicyId: &policyID,
		SubsettingRuleKey:  &ruleKey,
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")
	response, err := s.Client.ListSubsettingRuleProcessingChainObjects(ctx, request)
	if err != nil {
		return err
	}
	s.Res = &response.SubsettingRuleProcessingChainObjectsCollection

	processingChainObjectKey := s.D.Get("processing_chain_object_key").(string)
	for _, item := range s.Res.Items {
		if item.Key != nil && *item.Key == processingChainObjectKey && item.SubsettingSchemaRelationKey != nil {
			s.StableRelationKey = *item.SubsettingSchemaRelationKey
			return nil
		}
	}

	return fmt.Errorf("unable to determine the stable schema relation key for processing chain object %s", processingChainObjectKey)
}

func GetSubsettingPolicySubsettingRuleProcessingChainObjectCompositeId(processingChainObjectKey string, subsettingPolicyId string, subsettingRuleKey string) string {
	processingChainObjectKey = url.PathEscape(processingChainObjectKey)
	subsettingPolicyId = url.PathEscape(subsettingPolicyId)
	subsettingRuleKey = url.PathEscape(subsettingRuleKey)
	compositeId := "subsettingPolicies/" + subsettingPolicyId + "/subsettingRules/" + subsettingRuleKey + "/processingChainObjects/" + processingChainObjectKey
	return compositeId
}

func parseSubsettingPolicySubsettingRuleProcessingChainObjectCompositeId(compositeId string) (processingChainObjectKey string, subsettingPolicyId string, subsettingRuleKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("subsettingPolicies/.*/subsettingRules/.*/processingChainObjects/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	subsettingPolicyId, _ = url.PathUnescape(parts[1])
	subsettingRuleKey, _ = url.PathUnescape(parts[3])
	processingChainObjectKey, _ = url.PathUnescape(parts[5])

	return
}

func SubsettingRuleProcessingChainObjectSummaryToMap(obj oci_data_safe.SubsettingRuleProcessingChainObjectSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApproximateRowCountBeforeSubsetting != nil {
		result["approximate_row_count_before_subsetting"] = strconv.FormatInt(*obj.ApproximateRowCountBeforeSubsetting, 10)
	}

	result["child_columns"] = obj.ChildColumns

	if obj.ChildObjectName != nil {
		result["child_object_name"] = string(*obj.ChildObjectName)
	}

	if obj.ChildSchemaName != nil {
		result["child_schema_name"] = string(*obj.ChildSchemaName)
	}

	if obj.EstimatedRowCountAfterSubsetting != nil {
		result["estimated_row_count_after_subsetting"] = strconv.FormatInt(*obj.EstimatedRowCountAfterSubsetting, 10)
	}

	if obj.IsEnabledForProcessing != nil {
		result["is_enabled_for_processing"] = bool(*obj.IsEnabledForProcessing)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["parent_columns"] = obj.ParentColumns

	if obj.ParentObjectName != nil {
		result["parent_object_name"] = string(*obj.ParentObjectName)
	}

	if obj.ParentSchemaName != nil {
		result["parent_schema_name"] = string(*obj.ParentSchemaName)
	}

	result["propagation_impact"] = string(obj.PropagationImpact)

	if obj.SubsettingSchemaRelationKey != nil {
		result["subsetting_schema_relation_key"] = string(*obj.SubsettingSchemaRelationKey)
	}

	return result
}
