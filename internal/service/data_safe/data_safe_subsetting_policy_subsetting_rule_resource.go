// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingPolicySubsettingRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataSafeSubsettingPolicySubsettingRuleWithContext,
		ReadContext:   readDataSafeSubsettingPolicySubsettingRuleWithContext,
		UpdateContext: updateDataSafeSubsettingPolicySubsettingRuleWithContext,
		DeleteContext: deleteDataSafeSubsettingPolicySubsettingRuleWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"scope": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"schema_name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringLenBetween(1, 128),
						},
						"scope_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ALL",
								"SPECIFIC",
							}, true),
						},

						// Optional
						"object": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringLenBetween(1, 128),
						},

						// Computed
					},
				},
			},
			"subset_rule_entry": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"rule_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CONDITION",
								"PARTITION",
								"PERCENT",
							}, true),
						},

						// Optional
						"condition": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"partitions_list": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"percent": {
							Type:         schema.TypeInt,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.IntBetween(0, 100),
						},
						"sub_partitions_list": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"subsetting_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
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
			"peer_tables_action": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"MINIMUM_ROWS", "SUBSET", "MAXIMUM_ROWS"}, true),
			},
			"related_tables_propagation": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"ANCESTORS_AND_DESCENDANTS", "ANCESTORS", "DESCENDANTS", "NONE"}, true),
			},
			"rule_combination_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"UNION", "SERIAL"}, true),
			},

			// Computed
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeSubsettingPolicySubsettingRuleWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDataSafeSubsettingPolicySubsettingRuleWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDataSafeSubsettingPolicySubsettingRuleWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDataSafeSubsettingPolicySubsettingRuleWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataSafeSubsettingPolicySubsettingRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SubsettingRule
	DisableNotFoundRetries bool
}

func (s *DataSafeSubsettingPolicySubsettingRuleResourceCrud) ID() string {
	return GetSubsettingPolicySubsettingRuleCompositeId(s.D.Get("subsetting_policy_id").(string), s.D.Get("key").(string))
}

func (s *DataSafeSubsettingPolicySubsettingRuleResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_data_safe.CreateSubsettingRuleRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if peerTablesAction, ok := s.D.GetOkExists("peer_tables_action"); ok {
		request.PeerTablesAction = oci_data_safe.CreateSubsettingRuleDetailsPeerTablesActionEnum(peerTablesAction.(string))
	}

	if relatedTablesPropagation, ok := s.D.GetOkExists("related_tables_propagation"); ok {
		request.RelatedTablesPropagation = oci_data_safe.CreateSubsettingRuleDetailsRelatedTablesPropagationEnum(relatedTablesPropagation.(string))
	}

	if ruleCombinationMode, ok := s.D.GetOkExists("rule_combination_mode"); ok {
		request.RuleCombinationMode = oci_data_safe.CreateSubsettingRuleDetailsRuleCombinationModeEnum(ruleCombinationMode.(string))
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		if tmpList := scope.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scope", 0)
			tmp, err := s.mapToSubsetScope(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Scope = tmp
		}
	}

	if subsetRuleEntry, ok := s.D.GetOkExists("subset_rule_entry"); ok {
		if tmpList := subsetRuleEntry.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "subset_rule_entry", 0)
			tmp, err := s.mapToSubsetRuleEntry(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SubsetRuleEntry = tmp
		}
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateSubsettingRule(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSubsettingPolicySubsettingRuleFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSubsettingPolicySubsettingRuleResourceCrud) getSubsettingPolicySubsettingRuleFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	subsettingPolicySubsettingRuleId, err := subsettingPolicySubsettingRuleWaitForWorkRequest(ctx, workId, "subsettingpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, subsettingPolicySubsettingRuleId)
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

	if actionTypeEnum == oci_data_safe.WorkRequestResourceActionTypeCreated {
		// The create work request identifies the parent subsetting policy, rather
		// than the newly-created rule. Resolve the rule key before performing the
		// first read so the composite resource ID is valid.
		subsettingRuleKey, err := s.findCreatedSubsettingRuleKey(ctx)
		if err != nil {
			return err
		}
		s.D.Set("key", subsettingRuleKey)
		s.D.SetId(GetSubsettingPolicySubsettingRuleCompositeId(s.D.Get("subsetting_policy_id").(string), subsettingRuleKey))
	}

	return s.GetWithContext(ctx)
}

func (s *DataSafeSubsettingPolicySubsettingRuleResourceCrud) findCreatedSubsettingRuleKey(ctx context.Context) (string, error) {
	policyID := s.D.Get("subsetting_policy_id").(string)
	request := oci_data_safe.ListSubsettingRulesRequest{SubsettingPolicyId: &policyID}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListSubsettingRules(ctx, request)
	if err != nil {
		return "", err
	}
	rules := response.Items
	for response.OpcNextPage != nil {
		request.Page = response.OpcNextPage
		response, err = s.Client.ListSubsettingRules(ctx, request)
		if err != nil {
			return "", err
		}
		rules = append(rules, response.Items...)
	}

	expectedScope, err := s.mapToSubsetScope("scope.0.%s")
	if err != nil {
		return "", err
	}
	expectedSubsetRuleEntry, err := s.mapToSubsetRuleEntry("subset_rule_entry.0.%s")
	if err != nil {
		return "", err
	}

	matchingKeys := make([]string, 0, 1)
	for _, rule := range rules {
		if rule.Key == nil || !subsettingRuleFieldMatches(expectedScope, rule.Scope) ||
			!subsettingRuleFieldMatches(expectedSubsetRuleEntry, rule.SubsetRuleEntry) {
			continue
		}

		if displayName, ok := s.D.GetOk("display_name"); ok {
			if rule.DisplayName == nil || *rule.DisplayName != displayName.(string) {
				continue
			}
		}
		if description, ok := s.D.GetOk("description"); ok {
			if rule.Description == nil || *rule.Description != description.(string) {
				continue
			}
		}

		matchingKeys = append(matchingKeys, *rule.Key)
	}

	if len(matchingKeys) == 1 {
		return matchingKeys[0], nil
	}
	if len(matchingKeys) > 1 {
		return "", fmt.Errorf("unable to uniquely determine the key of the newly-created subsetting rule in policy %s: %d matching rules found", policyID, len(matchingKeys))
	}
	return "", fmt.Errorf("unable to determine the key of the newly-created subsetting rule in policy %s", policyID)
}

func subsettingRuleFieldMatches(expected interface{}, actual interface{}) bool {
	return reflect.DeepEqual(canonicalizeSubsettingRuleField(expected), canonicalizeSubsettingRuleField(actual))
}

func canonicalizeSubsettingRuleField(field interface{}) interface{} {
	switch value := field.(type) {
	case oci_data_safe.SubsetScopeForAllObjects:
		return map[string]interface{}{
			"scope_type":  "ALL",
			"schema_name": normalizeSubsettingIdentifier(pointerString(value.SchemaName)),
		}
	case oci_data_safe.SubsetScopeForSpecificObjects:
		return map[string]interface{}{
			"scope_type":  "SPECIFIC",
			"object":      normalizeSubsettingIdentifier(pointerString(value.ObjectName)),
			"schema_name": normalizeSubsettingIdentifier(pointerString(value.SchemaName)),
		}
	case oci_data_safe.ConditionSubsetRuleEntry:
		return map[string]interface{}{
			"rule_type": "CONDITION",
			"condition": strings.TrimSpace(pointerString(value.Condition)),
		}
	case oci_data_safe.PartitionSubsetRuleEntry:
		return map[string]interface{}{
			"rule_type":           "PARTITION",
			"partitions_list":     normalizeSubsettingIdentifierSet(value.PartitionsList),
			"sub_partitions_list": normalizeSubsettingIdentifierSet(value.SubPartitionsList),
		}
	case oci_data_safe.PercentSubsetRuleEntry:
		result := map[string]interface{}{"rule_type": "PERCENT"}
		if value.Percent != nil {
			result["percent"] = *value.Percent
		}
		return result
	case map[string]interface{}:
		return canonicalizeSubsettingRuleMap(value)
	default:
		return field
	}
}

func canonicalizeSubsettingRuleMap(field map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(field))
	for key, value := range field {
		switch key {
		case "scope_type", "rule_type":
			if stringValue, ok := value.(string); ok {
				result[key] = strings.ToUpper(strings.TrimSpace(stringValue))
			} else {
				result[key] = value
			}
		case "schema_name", "object":
			if stringValue, ok := value.(string); ok {
				result[key] = normalizeSubsettingIdentifier(stringValue)
			} else {
				result[key] = value
			}
		case "condition":
			if stringValue, ok := value.(string); ok {
				result[key] = strings.TrimSpace(stringValue)
			} else {
				result[key] = value
			}
		case "partitions_list", "sub_partitions_list":
			result[key] = normalizeSubsettingIdentifierSet(subsettingStringList(value))
		default:
			result[key] = value
		}
	}
	return result
}

func pointerString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func normalizeSubsettingIdentifier(value string) string {
	return strings.ToUpper(strings.TrimSpace(value))
}

func normalizeSubsettingIdentifierList(values []string) []string {
	result := make([]string, len(values))
	for i, value := range values {
		result[i] = normalizeSubsettingIdentifier(value)
	}
	return result
}

func normalizeSubsettingIdentifierSet(values []string) []string {
	set := make(map[string]struct{}, len(values))
	for _, value := range normalizeSubsettingIdentifierList(values) {
		set[value] = struct{}{}
	}
	result := make([]string, 0, len(set))
	for value := range set {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func subsettingStringList(value interface{}) []string {
	switch values := value.(type) {
	case []string:
		return values
	case []interface{}:
		result := make([]string, 0, len(values))
		for _, item := range values {
			if stringValue, ok := item.(string); ok {
				result = append(result, stringValue)
			}
		}
		return result
	default:
		return nil
	}
}

func subsettingPolicySubsettingRuleWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func subsettingPolicySubsettingRuleWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = subsettingPolicySubsettingRuleWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeSubsettingPolicySubsettingRuleWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSubsettingPolicySubsettingRuleWorkRequest(ctx context.Context, client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeSubsettingPolicySubsettingRuleResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetSubsettingRuleRequest{}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if subsettingRuleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := subsettingRuleKey.(string)
		request.SubsettingRuleKey = &tmp
	}

	subsettingPolicyId, subsettingRuleKey, err := parseSubsettingPolicySubsettingRuleCompositeId(s.D.Id())
	if err == nil {
		request.SubsettingPolicyId = &subsettingPolicyId
		request.SubsettingRuleKey = &subsettingRuleKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSubsettingRule(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.SubsettingRule
	return nil
}

func (s *DataSafeSubsettingPolicySubsettingRuleResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_data_safe.UpdateSubsettingRuleRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if peerTablesAction, ok := s.D.GetOkExists("peer_tables_action"); ok {
		request.PeerTablesAction = oci_data_safe.UpdateSubsettingRuleDetailsPeerTablesActionEnum(peerTablesAction.(string))
	}

	if relatedTablesPropagation, ok := s.D.GetOkExists("related_tables_propagation"); ok {
		request.RelatedTablesPropagation = oci_data_safe.UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum(relatedTablesPropagation.(string))
	}

	if ruleCombinationMode, ok := s.D.GetOkExists("rule_combination_mode"); ok {
		request.RuleCombinationMode = oci_data_safe.UpdateSubsettingRuleDetailsRuleCombinationModeEnum(ruleCombinationMode.(string))
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		if tmpList := scope.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scope", 0)
			tmp, err := s.mapToSubsetScope(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Scope = tmp
		}
	}

	if subsetRuleEntry, ok := s.D.GetOkExists("subset_rule_entry"); ok {
		if tmpList := subsetRuleEntry.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "subset_rule_entry", 0)
			tmp, err := s.mapToSubsetRuleEntry(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SubsetRuleEntry = tmp
		}
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if subsettingRuleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := subsettingRuleKey.(string)
		request.SubsettingRuleKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateSubsettingRule(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSubsettingPolicySubsettingRuleFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSubsettingPolicySubsettingRuleResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_data_safe.DeleteSubsettingRuleRequest{}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if subsettingRuleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := subsettingRuleKey.(string)
		request.SubsettingRuleKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteSubsettingRule(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := subsettingPolicySubsettingRuleWaitForWorkRequest(ctx, workId, "subsettingpolicy",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeSubsettingPolicySubsettingRuleResourceCrud) SetData() error {

	subsettingPolicyId, subsettingRuleKey, err := parseSubsettingPolicySubsettingRuleCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("subsetting_policy_id", &subsettingPolicyId)
		s.D.Set("key", &subsettingRuleKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("peer_tables_action", s.Res.PeerTablesAction)

	s.D.Set("related_tables_propagation", s.Res.RelatedTablesPropagation)

	s.D.Set("rule_combination_mode", s.Res.RuleCombinationMode)

	if s.Res.Scope != nil {
		scopeArray := []interface{}{}
		if scopeMap := SubsetScopeToMap(&s.Res.Scope); scopeMap != nil {
			scopeArray = append(scopeArray, scopeMap)
		}
		s.D.Set("scope", scopeArray)
	} else {
		s.D.Set("scope", nil)
	}

	if s.Res.SubsetRuleEntry != nil {
		subsetRuleEntryArray := []interface{}{}
		if subsetRuleEntryMap := SubsetRuleEntryToMap(&s.Res.SubsetRuleEntry); subsetRuleEntryMap != nil {
			subsetRuleEntryArray = append(subsetRuleEntryArray, subsetRuleEntryMap)
		}
		s.D.Set("subset_rule_entry", subsetRuleEntryArray)
	} else {
		s.D.Set("subset_rule_entry", nil)
	}

	return nil
}

func GetSubsettingPolicySubsettingRuleCompositeId(subsettingPolicyId string, subsettingRuleKey string) string {
	subsettingPolicyId = url.PathEscape(subsettingPolicyId)
	subsettingRuleKey = url.PathEscape(subsettingRuleKey)
	compositeId := "subsettingPolicies/" + subsettingPolicyId + "/subsettingRules/" + subsettingRuleKey
	return compositeId
}

func parseSubsettingPolicySubsettingRuleCompositeId(compositeId string) (subsettingPolicyId string, subsettingRuleKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("subsettingPolicies/.*/subsettingRules/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	subsettingPolicyId, _ = url.PathUnescape(parts[1])
	subsettingRuleKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *DataSafeSubsettingPolicySubsettingRuleResourceCrud) mapToSubsetRuleEntry(fieldKeyFormat string) (oci_data_safe.SubsetRuleEntry, error) {
	var baseObject oci_data_safe.SubsetRuleEntry
	//discriminator
	ruleTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_type"))
	var ruleType string
	if ok {
		ruleType = ruleTypeRaw.(string)
	} else {
		ruleType = "" // default value
	}
	switch strings.ToLower(ruleType) {
	case strings.ToLower("CONDITION"):
		details := oci_data_safe.ConditionSubsetRuleEntry{}
		condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition"))
		if !ok || strings.TrimSpace(condition.(string)) == "" {
			return nil, fmt.Errorf("condition must be specified when rule_type is CONDITION")
		}
		tmp := condition.(string)
		details.Condition = &tmp
		baseObject = details
	case strings.ToLower("PARTITION"):
		details := oci_data_safe.PartitionSubsetRuleEntry{}
		if partitionsList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "partitions_list")); ok {
			interfaces := partitionsList.([]interface{})
			if len(interfaces) > 999 {
				return nil, fmt.Errorf("partitions_list cannot contain more than 999 items")
			}
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) == 0 && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "partitions_list")) {
				return nil, fmt.Errorf("partitions_list must contain at least one item when specified")
			}
			if len(tmp) != 0 {
				details.PartitionsList = tmp
			}
		}
		if subPartitionsList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sub_partitions_list")); ok {
			interfaces := subPartitionsList.([]interface{})
			if len(interfaces) > 999 {
				return nil, fmt.Errorf("sub_partitions_list cannot contain more than 999 items")
			}
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) == 0 && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sub_partitions_list")) {
				return nil, fmt.Errorf("sub_partitions_list must contain at least one item when specified")
			}
			if len(tmp) != 0 {
				details.SubPartitionsList = tmp
			}
		}
		baseObject = details
	case strings.ToLower("PERCENT"):
		details := oci_data_safe.PercentSubsetRuleEntry{}
		percent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "percent"))
		if !ok {
			return nil, fmt.Errorf("percent must be specified when rule_type is PERCENT")
		}
		tmp := percent.(int)
		details.Percent = &tmp
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown rule_type '%v' was specified", ruleType)
	}
	return baseObject, nil
}

func SubsetRuleEntryToMap(obj *oci_data_safe.SubsetRuleEntry) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_data_safe.ConditionSubsetRuleEntry:
		result["rule_type"] = "CONDITION"

		if v.Condition != nil {
			result["condition"] = string(*v.Condition)
		}
	case oci_data_safe.PartitionSubsetRuleEntry:
		result["rule_type"] = "PARTITION"

		result["partitions_list"] = v.PartitionsList

		result["sub_partitions_list"] = v.SubPartitionsList
	case oci_data_safe.PercentSubsetRuleEntry:
		result["rule_type"] = "PERCENT"

		if v.Percent != nil {
			result["percent"] = int(*v.Percent)
		}
	default:
		log.Printf("[WARN] Received 'rule_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DataSafeSubsettingPolicySubsettingRuleResourceCrud) mapToSubsetScope(fieldKeyFormat string) (oci_data_safe.SubsetScope, error) {
	var baseObject oci_data_safe.SubsetScope
	//discriminator
	scopeTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope_type"))
	var scopeType string
	if ok {
		scopeType = scopeTypeRaw.(string)
	} else {
		scopeType = "" // default value
	}
	switch strings.ToLower(scopeType) {
	case strings.ToLower("ALL"):
		details := oci_data_safe.SubsetScopeForAllObjects{}
		if schemaName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_name")); ok {
			tmp := schemaName.(string)
			details.SchemaName = &tmp
		}
		baseObject = details
	case strings.ToLower("SPECIFIC"):
		details := oci_data_safe.SubsetScopeForSpecificObjects{}
		object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object"))
		if !ok || strings.TrimSpace(object.(string)) == "" {
			return nil, fmt.Errorf("object must be specified when scope_type is SPECIFIC")
		}
		tmp := object.(string)
		details.ObjectName = &tmp
		if schemaName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_name")); ok {
			tmp := schemaName.(string)
			details.SchemaName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown scope_type '%v' was specified", scopeType)
	}
	return baseObject, nil
}

func SubsetScopeToMap(obj *oci_data_safe.SubsetScope) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_data_safe.SubsetScopeForAllObjects:
		result["scope_type"] = "ALL"

		if v.SchemaName != nil {
			result["schema_name"] = string(*v.SchemaName)
		}
	case oci_data_safe.SubsetScopeForSpecificObjects:
		result["scope_type"] = "SPECIFIC"

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}

		if v.SchemaName != nil {
			result["schema_name"] = string(*v.SchemaName)
		}
	default:
		log.Printf("[WARN] Received 'scope_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func SubsettingRuleSummaryToMap(obj oci_data_safe.SubsettingRuleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["peer_tables_action"] = string(obj.PeerTablesAction)

	result["related_tables_propagation"] = string(obj.RelatedTablesPropagation)

	result["rule_combination_mode"] = string(obj.RuleCombinationMode)

	if obj.Scope != nil {
		scopeArray := []interface{}{}
		if scopeMap := SubsetScopeToMap(&obj.Scope); scopeMap != nil {
			scopeArray = append(scopeArray, scopeMap)
		}
		result["scope"] = scopeArray
	}

	if obj.SubsetRuleEntry != nil {
		subsetRuleEntryArray := []interface{}{}
		if subsetRuleEntryMap := SubsetRuleEntryToMap(&obj.SubsetRuleEntry); subsetRuleEntryMap != nil {
			subsetRuleEntryArray = append(subsetRuleEntryArray, subsetRuleEntryMap)
		}
		result["subset_rule_entry"] = subsetRuleEntryArray
	}

	return result
}
