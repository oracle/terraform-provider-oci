// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSecurityAssessmentFindingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSecurityAssessmentFinding,
		Read:     readDataSafeSecurityAssessmentFinding,
		Update:   updateDataSafeSecurityAssessmentFinding,
		Delete:   deleteDataSafeSecurityAssessmentFinding,
		Schema: map[string]*schema.Schema{
			// Required
			"security_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
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
							Type:             schema.TypeMap,
							Required:         true,
							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"assessment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"has_target_db_risk_level_changed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_risk_modified": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_top_finding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"justification": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oneline": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oracle_defined_severity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"references": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cis": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gdpr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"obp": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"stig": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"remarks": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"summary": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_valid_until": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"title": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeSecurityAssessmentFinding(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentFindingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeSecurityAssessmentFinding(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentFindingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeSecurityAssessmentFinding(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentFindingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeSecurityAssessmentFinding(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeSecurityAssessmentFindingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.FindingSummary
	PatchResponse          *oci_data_safe.Finding
	DisableNotFoundRetries bool
}

func (s *DataSafeSecurityAssessmentFindingResourceCrud) ID() string {
	return GetSecurityAssessmentFindingCompositeId(s.D.Get("security_assessment_id").(string))
}

func (s *DataSafeSecurityAssessmentFindingResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DataSafeSecurityAssessmentFindingResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.FindingLifecycleStateActive),
		string(oci_data_safe.FindingLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeSecurityAssessmentFindingResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DataSafeSecurityAssessmentFindingResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DataSafeSecurityAssessmentFindingResourceCrud) Create() error {
	return nil
}
func (s *DataSafeSecurityAssessmentFindingResourceCrud) Patch() error {
	request := oci_data_safe.PatchFindingsRequest{}

	if patchOperations, ok := s.D.GetOkExists("patch_operations"); ok {
		interfaces := patchOperations.([]interface{})
		tmp := make([]oci_data_safe.PatchInstruction, len(interfaces))
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

	if securityAssessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := securityAssessmentId.(string)
		request.SecurityAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")
	response, err := s.Client.PatchFindings(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSecurityAssessmentFindingFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSecurityAssessmentFindingResourceCrud) getSecurityAssessmentFindingFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	securityAssessmentFindingId, err := securityAssessmentFindingWaitForWorkRequest(workId, "securityassessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, securityAssessmentFindingId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
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
	s.D.SetId(*securityAssessmentFindingId)

	return s.Get()
}

func securityAssessmentFindingWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func securityAssessmentFindingWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = securityAssessmentFindingWorkRequestShouldRetryFunc(timeout)

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
			response, err = client.GetWorkRequest(context.Background(),
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeSecurityAssessmentFindingWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSecurityAssessmentFindingWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
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

func (s *DataSafeSecurityAssessmentFindingResourceCrud) Get() error {
	request := oci_data_safe.ListFindingsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListFindingsAccessLevelEnum(accessLevel.(string))
	}

	if category, ok := s.D.GetOkExists("category"); ok {
		tmp := category.(string)
		request.Category = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if containsReferences, ok := s.D.GetOkExists("contains_references"); ok {
		interfaces := containsReferences.([]interface{})
		tmp := make([]oci_data_safe.SecurityAssessmentReferencesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.SecurityAssessmentReferencesEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("contains_references") {
			request.ContainsReferences = tmp
		}
	}

	if containsSeverity, ok := s.D.GetOkExists("contains_severity"); ok {
		interfaces := containsSeverity.([]interface{})
		tmp := make([]oci_data_safe.ListFindingsContainsSeverityEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.ListFindingsContainsSeverityEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("contains_severity") {
			request.ContainsSeverity = tmp
		}
	}

	if field, ok := s.D.GetOkExists("field"); ok {
		interfaces := field.([]interface{})
		tmp := make([]oci_data_safe.ListFindingsFieldEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.ListFindingsFieldEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("field") {
			request.Field = tmp
		}
	}

	if findingKey, ok := s.D.GetOkExists("finding_key"); ok {
		tmp := findingKey.(string)
		request.FindingKey = &tmp
	}

	if isTopFinding, ok := s.D.GetOkExists("is_top_finding"); ok {
		tmp := isTopFinding.(bool)
		request.IsTopFinding = &tmp
	}

	if references, ok := s.D.GetOkExists("references"); ok {
		if tmpList := references.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "references", 0)
			tmp, err := s.mapToReferences(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.References = tmp
		}
	}

	if scimQuery, ok := s.D.GetOkExists("scim_query"); ok {
		tmp := scimQuery.(string)
		request.ScimQuery = &tmp
	}

	if securityAssessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := securityAssessmentId.(string)
		request.SecurityAssessmentId = &tmp
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_data_safe.ListFindingsSeverityEnum(severity.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListFindingsLifecycleStateEnum(state.(string))
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if targetIds, ok := s.D.GetOkExists("target_ids"); ok {
		interfaces := targetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("target_ids") {
			request.TargetIds = tmp
		}
	}

	securityAssessmentId, err := parseSecurityAssessmentFindingCompositeId(s.D.Id())
	if err == nil {
		request.SecurityAssessmentId = &securityAssessmentId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListFindings(context.Background(), request)
	if err != nil {
		return err
	}
	if len(response.Items) > 0 {
		finding := response.Items[0]
		s.Res = &finding
	}
	return err
}

func (s *DataSafeSecurityAssessmentFindingResourceCrud) mapToReferences(fieldKeyFormat string) (oci_data_safe.ListFindingsReferencesEnum, error) {
	panic("unimplemented")
}

func (s *DataSafeSecurityAssessmentFindingResourceCrud) Update() error {
	err := s.Patch()
	if err != nil {
		log.Printf("[ERROR] Failed to execute Patch operation: %v", err)
		return err
	}
	return nil
}

func (s *DataSafeSecurityAssessmentFindingResourceCrud) SetData() error {

	securityAssessmentId, err := parseSecurityAssessmentFindingCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("security_assessment_id", &securityAssessmentId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AssessmentId != nil {
		s.D.Set("assessment_id", *s.Res.AssessmentId)
	}

	// if s.Res.CategoryName != nil {
	// 	s.D.Set("category", *s.Res.Category)
	// }

	if s.Res.Details != nil {
		s.D.Set("details", []interface{}{objectToMap((*s.Res.Details).(map[string]interface{}))})
	} else {
		s.D.Set("details", nil)
	}

	if s.Res.HasTargetDbRiskLevelChanged != nil {
		s.D.Set("has_target_db_risk_level_changed", *s.Res.HasTargetDbRiskLevelChanged)
	}

	if s.Res.IsRiskModified != nil {
		s.D.Set("is_risk_modified", *s.Res.IsRiskModified)
	}

	// if s.Res.IsTopFinding != nil {
	// 	s.D.Set("is_top_finding", *s.Res.IsTopFinding)
	// }

	if s.Res.Justification != nil {
		s.D.Set("justification", *s.Res.Justification)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	// if s.Res.Oneline != nil {
	// 	s.D.Set("oneline", *s.Res.Oneline)
	// }

	s.D.Set("oracle_defined_severity", s.Res.OracleDefinedSeverity)

	if s.Res.References != nil {
		s.D.Set("references", []interface{}{ReferencesToMap(s.Res.References)})
	} else {
		s.D.Set("references", nil)
	}

	if s.Res.Remarks != nil {
		s.D.Set("remarks", *s.Res.Remarks)
	}

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Summary != nil {
		s.D.Set("summary", *s.Res.Summary)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeValidUntil != nil {
		s.D.Set("time_valid_until", s.Res.TimeValidUntil.String())
	}

	if s.Res.Title != nil {
		s.D.Set("title", *s.Res.Title)
	}

	return nil
}

func GetSecurityAssessmentFindingCompositeId(securityAssessmentId string) string {
	securityAssessmentId = url.PathEscape(securityAssessmentId)
	compositeId := "securityAssessments/" + securityAssessmentId + "/findings"
	return compositeId
}

func parseSecurityAssessmentFindingCompositeId(compositeId string) (securityAssessmentId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("securityAssessments/.*/findings", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	securityAssessmentId, _ = url.PathUnescape(parts[1])

	return
}

func (s *DataSafeSecurityAssessmentFindingResourceCrud) mapToPatchInstruction(fieldKeyFormat string) (oci_data_safe.PatchInstruction, error) {
	var baseObject oci_data_safe.PatchInstruction
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
		details := oci_data_safe.PatchInsertInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("MERGE"):
		details := oci_data_safe.PatchMergeInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("REMOVE"):
		details := oci_data_safe.PatchRemoveInstruction{}
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

func PatchInstructionToMap1(obj oci_data_safe.PatchInstruction) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_data_safe.PatchInsertInstruction:
		result["operation"] = "INSERT"

		if v.Value != nil {
			result["value"] = []interface{}{
				objectToMap((*v.Value).(map[string]interface{})),
			}

		}
	case oci_data_safe.PatchMergeInstruction:
		result["operation"] = "MERGE"

		if v.Value != nil {
			result["value"] = []interface{}{
				objectToMap((*v.Value).(map[string]interface{})),
			}

		}
	case oci_data_safe.PatchRemoveInstruction:
		result["operation"] = "REMOVE"
	default:
		log.Printf("[WARN] Received 'operation' of unknown type %v", obj)
		return nil
	}

	return result
}
