// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSqlFirewallPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSqlFirewallPolicy,
		Read:     readDataSafeSqlFirewallPolicy,
		Update:   updateDataSafeSqlFirewallPolicy,
		Delete:   deleteDataSafeSqlFirewallPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"sql_firewall_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"allowed_client_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allowed_client_os_usernames": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allowed_client_programs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
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
			"enforcement_scope": {
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
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"violation_action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"violation_audit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"db_user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sql_level": {
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

func createDataSafeSqlFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	compartment, ok := sync.D.GetOkExists("compartment_id")

	err := tfresource.CreateResource(d, sync)
	if err != nil {
		return err
	}

	if ok && compartment != *sync.Res.CompartmentId {
		err = sync.updateCompartment(compartment)
		if err != nil {
			return err
		}
		tmp := compartment.(string)
		sync.Res.CompartmentId = &tmp
		err := sync.Get()
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

func readDataSafeSqlFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeSqlFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeSqlFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true
	return tfresource.DeleteResource(d, sync)
}

type DataSafeSqlFirewallPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SqlFirewallPolicy
	DisableNotFoundRetries bool
}

func (s *DataSafeSqlFirewallPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeSqlFirewallPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.SqlFirewallPolicyLifecycleStateCreating),
	}
}

func (s *DataSafeSqlFirewallPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.SqlFirewallPolicyLifecycleStateActive),
		string(oci_data_safe.SqlFirewallPolicyLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeSqlFirewallPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.SqlFirewallPolicyLifecycleStateDeleting),
	}
}

func (s *DataSafeSqlFirewallPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.SqlFirewallPolicyLifecycleStateDeleted),
	}
}

func (s *DataSafeSqlFirewallPolicyResourceCrud) Create() error {
	request := oci_data_safe.UpdateSqlFirewallPolicyRequest{}

	if allowedClientIps, ok := s.D.GetOkExists("allowed_client_ips"); ok {
		interfaces := allowedClientIps.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_client_ips") {
			request.AllowedClientIps = tmp
		}
	}

	if allowedClientOsUsernames, ok := s.D.GetOkExists("allowed_client_os_usernames"); ok {
		interfaces := allowedClientOsUsernames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_client_os_usernames") {
			request.AllowedClientOsUsernames = tmp
		}
	}

	if allowedClientPrograms, ok := s.D.GetOkExists("allowed_client_programs"); ok {
		interfaces := allowedClientPrograms.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_client_programs") {
			request.AllowedClientPrograms = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if enforcementScope, ok := s.D.GetOkExists("enforcement_scope"); ok {
		request.EnforcementScope = oci_data_safe.UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum(enforcementScope.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if sqlFirewallPolicyId, ok := s.D.GetOkExists("sql_firewall_policy_id"); ok {
		tmp := sqlFirewallPolicyId.(string)
		request.SqlFirewallPolicyId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_data_safe.UpdateSqlFirewallPolicyDetailsStatusEnum(status.(string))
	}

	if violationAction, ok := s.D.GetOkExists("violation_action"); ok {
		request.ViolationAction = oci_data_safe.UpdateSqlFirewallPolicyDetailsViolationActionEnum(violationAction.(string))
	}

	if violationAudit, ok := s.D.GetOkExists("violation_audit"); ok {
		request.ViolationAudit = oci_data_safe.UpdateSqlFirewallPolicyDetailsViolationAuditEnum(violationAudit.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateSqlFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_data_safe.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_data_safe.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "sqlfirewallpolicy") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getSqlFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSqlFirewallPolicyResourceCrud) getSqlFirewallPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	sqlFirewallPolicyId, err := sqlFirewallPolicyWaitForWorkRequest(workId, "sqlfirewallpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		//Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, sqlFirewallPolicyId)
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
	s.D.SetId(*sqlFirewallPolicyId)

	return s.Get()
}

func sqlFirewallPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func sqlFirewallPolicyWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = sqlFirewallPolicyWorkRequestShouldRetryFunc(timeout)

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
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeSqlFirewallPolicyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSqlFirewallPolicyWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeSqlFirewallPolicyResourceCrud) Get() error {
	request := oci_data_safe.GetSqlFirewallPolicyRequest{}

	tmp := s.D.Id()
	request.SqlFirewallPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSqlFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SqlFirewallPolicy
	return nil
}

func (s *DataSafeSqlFirewallPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateSqlFirewallPolicyRequest{}

	if allowedClientIps, ok := s.D.GetOkExists("allowed_client_ips"); ok {
		interfaces := allowedClientIps.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_client_ips") {
			request.AllowedClientIps = tmp
		}
	}

	if allowedClientOsUsernames, ok := s.D.GetOkExists("allowed_client_os_usernames"); ok {
		interfaces := allowedClientOsUsernames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_client_os_usernames") {
			request.AllowedClientOsUsernames = tmp
		}
	}

	if allowedClientPrograms, ok := s.D.GetOkExists("allowed_client_programs"); ok {
		interfaces := allowedClientPrograms.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_client_programs") {
			request.AllowedClientPrograms = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if enforcementScope, ok := s.D.GetOkExists("enforcement_scope"); ok {
		request.EnforcementScope = oci_data_safe.UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum(enforcementScope.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.SqlFirewallPolicyId = &tmp

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_data_safe.UpdateSqlFirewallPolicyDetailsStatusEnum(status.(string))
	}

	if violationAction, ok := s.D.GetOkExists("violation_action"); ok {
		request.ViolationAction = oci_data_safe.UpdateSqlFirewallPolicyDetailsViolationActionEnum(violationAction.(string))
	}

	if violationAudit, ok := s.D.GetOkExists("violation_audit"); ok {
		request.ViolationAudit = oci_data_safe.UpdateSqlFirewallPolicyDetailsViolationAuditEnum(violationAudit.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateSqlFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSqlFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSqlFirewallPolicyResourceCrud) Delete() error {
	request := oci_data_safe.DeleteSqlFirewallPolicyRequest{}

	tmp := s.D.Id()
	request.SqlFirewallPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteSqlFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := sqlFirewallPolicyWaitForWorkRequest(workId, "sqlfirewallpolicy",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeSqlFirewallPolicyResourceCrud) SetData() error {
	s.D.Set("allowed_client_ips", s.Res.AllowedClientIps)
	s.D.Set("allowed_client_ips", s.Res.AllowedClientIps)

	s.D.Set("allowed_client_os_usernames", s.Res.AllowedClientOsUsernames)
	s.D.Set("allowed_client_os_usernames", s.Res.AllowedClientOsUsernames)

	s.D.Set("allowed_client_programs", s.Res.AllowedClientPrograms)
	s.D.Set("allowed_client_programs", s.Res.AllowedClientPrograms)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbUserName != nil {
		s.D.Set("db_user_name", *s.Res.DbUserName)
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

	s.D.Set("enforcement_scope", s.Res.EnforcementScope)

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SecurityPolicyId != nil {
		s.D.Set("security_policy_id", *s.Res.SecurityPolicyId)
	}

	s.D.Set("sql_level", s.Res.SqlLevel)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("violation_action", s.Res.ViolationAction)

	s.D.Set("violation_audit", s.Res.ViolationAudit)

	return nil
}

func SqlFirewallPolicySummaryToMap(obj oci_data_safe.SqlFirewallPolicySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DbUserName != nil {
		result["db_user_name"] = string(*obj.DbUserName)
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

	result["enforcement_scope"] = string(obj.EnforcementScope)

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.SecurityPolicyId != nil {
		result["security_policy_id"] = string(*obj.SecurityPolicyId)
	}

	result["sql_level"] = string(obj.SqlLevel)

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["violation_action"] = string(obj.ViolationAction)

	result["violation_audit"] = string(obj.ViolationAudit)

	return result
}

func (s *DataSafeSqlFirewallPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeSqlFirewallPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SqlFirewallPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeSqlFirewallPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSqlFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
