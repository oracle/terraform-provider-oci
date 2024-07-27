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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAlertPolicyRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeAlertPolicyRule,
		Read:     readDataSafeAlertPolicyRule,
		Update:   updateDataSafeAlertPolicyRule,
		Delete:   deleteDataSafeAlertPolicyRule,
		Schema: map[string]*schema.Schema{
			// Required
			"alert_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"expression": {
				Type:     schema.TypeString,
				Required: true,
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

			// Computed
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeAlertPolicyRule(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAlertPolicyRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeAlertPolicyRule(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAlertPolicyRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeAlertPolicyRule(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAlertPolicyRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeAlertPolicyRule(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAlertPolicyRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeAlertPolicyRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.AlertPolicyRule
	DisableNotFoundRetries bool
}

func (s *DataSafeAlertPolicyRuleResourceCrud) ID() string {
	rule := *s.Res
	return GetAlertPolicyRuleCompositeId(s.D.Get("alert_policy_id").(string), *rule.Key)
}

func (s *DataSafeAlertPolicyRuleResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.AlertPolicyRuleLifecycleStateCreating),
	}
}

func (s *DataSafeAlertPolicyRuleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.AlertPolicyRuleLifecycleStateActive),
	}
}

func (s *DataSafeAlertPolicyRuleResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.AlertPolicyRuleLifecycleStateDeleting),
	}
}

func (s *DataSafeAlertPolicyRuleResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DataSafeAlertPolicyRuleResourceCrud) Create() error {
	request := oci_data_safe.CreateAlertPolicyRuleRequest{}

	if alertPolicyId, ok := s.D.GetOkExists("alert_policy_id"); ok {
		tmp := alertPolicyId.(string)
		request.AlertPolicyId = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if expression, ok := s.D.GetOkExists("expression"); ok {
		tmp := expression.(string)
		request.Expression = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateAlertPolicyRule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAlertPolicyRuleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeAlertPolicyRuleResourceCrud) getAlertPolicyRuleFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	alertPolicyRuleId, err := alertPolicyRuleWaitForWorkRequest(workId, "alertpolicyrule",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, alertPolicyRuleId)
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
	s.D.SetId(*alertPolicyRuleId)

	return s.Get()
}

func alertPolicyRuleWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func alertPolicyRuleWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = alertPolicyRuleWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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
				identifier = res.EntityUri
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeAlertPolicyRuleWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeAlertPolicyRuleWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeAlertPolicyRuleResourceCrud) Get() error {
	request := oci_data_safe.GetAlertPolicyRuleRequest{}

	if alertPolicyId, ok := s.D.GetOkExists("alert_policy_id"); ok {
		tmp := alertPolicyId.(string)
		request.AlertPolicyId = &tmp
	}

	if ruleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := ruleKey.(string)
		request.RuleKey = &tmp
	}

	alertPolicyId, ruleKey, err := parseAlertPolicyRuleCompositeId(s.D.Id())
	if err == nil {
		request.AlertPolicyId = &alertPolicyId
		request.RuleKey = &ruleKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetAlertPolicyRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AlertPolicyRule
	return nil
}

func (s *DataSafeAlertPolicyRuleResourceCrud) Update() error {
	request := oci_data_safe.UpdateAlertPolicyRuleRequest{}

	if alertPolicyId, ok := s.D.GetOkExists("alert_policy_id"); ok {
		tmp := alertPolicyId.(string)
		request.AlertPolicyId = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if expression, ok := s.D.GetOkExists("expression"); ok {
		tmp := expression.(string)
		request.Expression = &tmp
	}

	if ruleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := ruleKey.(string)
		request.RuleKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAlertPolicyRule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAlertPolicyRuleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAlertPolicyRuleResourceCrud) Delete() error {
	request := oci_data_safe.DeleteAlertPolicyRuleRequest{}

	if alertPolicyId, ok := s.D.GetOkExists("alert_policy_id"); ok {
		tmp := alertPolicyId.(string)
		request.AlertPolicyId = &tmp
	}

	if ruleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := ruleKey.(string)
		request.RuleKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteAlertPolicyRule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := alertPolicyRuleWaitForWorkRequest(workId, "alertpolicyrule",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeAlertPolicyRuleResourceCrud) SetData() error {

	alertPolicyId, ruleKey, err := parseAlertPolicyRuleCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("alert_policy_id", &alertPolicyId)
		s.D.Set("key", &ruleKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Expression != nil {
		s.D.Set("expression", *s.Res.Expression)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func GetAlertPolicyRuleCompositeId(alertPolicyId string, ruleKey string) string {
	alertPolicyId = url.PathEscape(alertPolicyId)
	ruleKey = url.PathEscape(ruleKey)
	compositeId := "alertPolicies/" + alertPolicyId + "/rules/" + ruleKey
	return compositeId
}

func parseAlertPolicyRuleCompositeId(compositeId string) (alertPolicyId string, ruleKey string, err error) {
	firstChar := compositeId[0:1]
	var compositeIdStr string
	if firstChar == "/" {
		compositeIdStr = trimLeftChar(compositeId)
	} else {
		compositeIdStr = compositeId
	}
	parts := strings.Split(compositeIdStr, "/")
	match, _ := regexp.MatchString("alertPolicies/.*/rules/.*", compositeIdStr)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeIdStr)
		return
	}
	alertPolicyId, _ = url.PathUnescape(parts[1])
	ruleKey, _ = url.PathUnescape(parts[3])

	return
}
