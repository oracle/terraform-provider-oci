// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_waas "github.com/oracle/oci-go-sdk/v58/waas"
)

func WaasProtectionRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("2h"),
			Update: tfresource.GetTimeoutDuration("2h"),
			Delete: tfresource.GetTimeoutDuration("2h"),
		},
		Create: createWaasProtectionRule,
		Read:   readWaasProtectionRule,
		Update: updateWaasProtectionRule,
		Delete: deleteWaasProtectionRule,
		Schema: map[string]*schema.Schema{
			// Required
			"waas_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exclusions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"exclusions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"target": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			// Computed
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"mod_security_rule_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createWaasProtectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &WaasProtectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.CreateResource(d, sync)
}

func readWaasProtectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &WaasProtectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.ReadResource(sync)
}

func updateWaasProtectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &WaasProtectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteWaasProtectionRule(d *schema.ResourceData, m interface{}) error {
	return nil
}

type WaasProtectionRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_waas.WaasClient
	Res                    *oci_waas.ProtectionRule
	DisableNotFoundRetries bool
}

func (s *WaasProtectionRuleResourceCrud) ID() string {
	return getWaasProtectionRuleCompositeId(s.D.Get("waas_policy_id").(string), s.D.Get("key").(string))
}

func getWaasProtectionRuleCompositeId(waasPolicyId, key string) string {
	waasPolicyId = url.PathEscape(waasPolicyId)
	key = url.PathEscape(key)
	compositeId := "waasPolicyId/" + waasPolicyId + "/key/" + key

	return compositeId
}

func parseWaasProtectionRuleCompositeId(compositeId string) (waasPolicyId, key string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("waasPolicyId/.*/key/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	waasPolicyId, _ = url.PathUnescape(parts[1])
	key, _ = url.PathUnescape(parts[3])
	return
}

func (s *WaasProtectionRuleResourceCrud) Create() error {
	return s.Update()
}

func (s *WaasProtectionRuleResourceCrud) getProtectionRuleFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_waas.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	waasPolicyId, err := protectionRuleWaitForWorkRequest(workId, "waas",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, waasPolicyId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_waas.CancelWorkRequestRequest{
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

	return s.Get()
}

func protectionRuleWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "waas", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_waas.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func protectionRuleWaitForWorkRequest(wId *string, entityType string, action oci_waas.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_waas.WaasClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "waas")
	retryPolicy.ShouldRetryOperation = protectionRuleWorkRequestShouldRetryFunc(timeout)

	response := oci_waas.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_waas.WorkRequestStatusValuesInProgress),
			string(oci_waas.WorkRequestStatusValuesAccepted),
			string(oci_waas.WorkRequestStatusValuesCanceling),
		},
		Target: []string{
			string(oci_waas.WorkRequestStatusValuesSucceeded),
			string(oci_waas.WorkRequestStatusValuesFailed),
			string(oci_waas.WorkRequestStatusValuesCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_waas.GetWorkRequestRequest{
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

	// The workrequest didn't do all its intended tasks, if the errors is set; so we should check for it
	var workRequestErr error
	if len(response.Errors) > 0 {
		errorMessage := getErrorFromProtectionRuleWorkRequest(response)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)
	}

	return identifier, workRequestErr
}

func getErrorFromProtectionRuleWorkRequest(response oci_waas.GetWorkRequestResponse) string {
	allErrs := make([]string, 0)
	for _, wrkErr := range response.Errors {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")
	return errorMessage
}

func (s *WaasProtectionRuleResourceCrud) Get() error {

	waasPolicyId, key, err := parseWaasProtectionRuleCompositeId(s.D.Id())
	if err != nil {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		return err
	}

	request := oci_waas.GetProtectionRuleRequest{
		WaasPolicyId:      &waasPolicyId,
		ProtectionRuleKey: &key,
	}
	s.D.Set("waas_policy_id", waasPolicyId) // Response doesn't have this field

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.GetProtectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProtectionRule
	return nil
}

func (s *WaasProtectionRuleResourceCrud) Update() error {
	request := oci_waas.UpdateProtectionRulesRequest{}
	protectionRuleAction := oci_waas.ProtectionRuleAction{}

	if waasPolicyId, ok := s.D.GetOkExists("waas_policy_id"); ok {
		tmp := waasPolicyId.(string)
		request.WaasPolicyId = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		protectionRuleAction.Key = &tmp
	}
	compositeId := getWaasProtectionRuleCompositeId(*request.WaasPolicyId, *protectionRuleAction.Key)
	s.D.SetId(compositeId)

	if action, ok := s.D.GetOkExists("action"); ok {
		protectionRuleAction.Action = oci_waas.ProtectionRuleActionActionEnum(action.(string))
	}

	if exclusions, ok := s.D.GetOkExists("exclusions"); ok {
		interfaces := exclusions.([]interface{})
		tmp := make([]oci_waas.ProtectionRuleExclusion, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "exclusions", stateDataIndex)
			converted, err := s.mapToProtectionRulesExclusion(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("exclusions") {
			protectionRuleAction.Exclusions = tmp
		}
	}

	request.ProtectionRules = append(request.ProtectionRules, protectionRuleAction)

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.UpdateProtectionRules(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getProtectionRuleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas"), oci_waas.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *WaasProtectionRuleResourceCrud) SetData() error {
	s.D.Set("action", s.Res.Action)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	exclusions := []interface{}{}
	for _, item := range s.Res.Exclusions {
		exclusions = append(exclusions, ProtectionRuleExclusionToMap(item))
	}
	s.D.Set("exclusions", exclusions)

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("labels", s.Res.Labels)

	s.D.Set("mod_security_rule_ids", s.Res.ModSecurityRuleIds)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	return nil
}

func (s *WaasProtectionRuleResourceCrud) mapToProtectionRulesExclusion(fieldKeyFormat string) (oci_waas.ProtectionRuleExclusion, error) {
	result := oci_waas.ProtectionRuleExclusion{}
	if exclusions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclusions")); ok {
		interfaces := exclusions.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclusions")) {
			result.Exclusions = tmp
		}
	}
	if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
		result.Target = oci_waas.ProtectionRuleExclusionTargetEnum(target.(string))
	}
	return result, nil
}
