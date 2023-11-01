// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSqlFirewallPolicyManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSqlFirewallPolicyManagement,
		Read:     readDataSafeSqlFirewallPolicyManagement,
		Update:   updateDataSafeSqlFirewallPolicyManagement,
		Delete:   deleteDataSafeSqlFirewallPolicyManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"sql_firewall_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"db_user_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
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

func createDataSafeSqlFirewallPolicyManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallPolicyManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	err := sync.GetSqlFirewallPolicyWorkReq()
	err1 := sync.Get()
	if err != nil {
		return err
	}
	if err1 != nil {
		return err1
	}
	return updateDataSafeSqlFirewallPolicyManagement(d, m)
}

func readDataSafeSqlFirewallPolicyManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallPolicyManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeSqlFirewallPolicyManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallPolicyManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeSqlFirewallPolicyManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallPolicyManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true
	return tfresource.DeleteResource(d, sync)
}

type DataSafeSqlFirewallPolicyManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SqlFirewallPolicy
	DisableNotFoundRetries bool
}

func (s *DataSafeSqlFirewallPolicyManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeSqlFirewallPolicyManagementResourceCrud) GetSqlFirewallPolicyWorkReq() error {
	listWorkRequestsRequest := oci_data_safe.ListWorkRequestsRequest{SortBy: oci_data_safe.ListWorkRequestsSortByEnum("ACCEPTEDTIME"), SortOrder: oci_data_safe.ListWorkRequestsSortOrderEnum("DESC")}
	var workId *string
	var sqlfwpId *string
	var databaseUserNm *string
	tmp := "GENERATE_FIREWALL_POLICY"
	listWorkRequestsRequest.OperationType = &tmp
	compartmentIdInSubtree := true
	listWorkRequestsRequest.CompartmentIdInSubtree = &compartmentIdInSubtree
	listWorkRequestsRequest.AccessLevel = oci_data_safe.ListWorkRequestsAccessLevelAccessible

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		listWorkRequestsRequest.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		listWorkRequestsRequest.ResourceId = &tmp
	}

	if dbUserName, ok := s.D.GetOkExists("db_user_name"); ok {
		tmp := dbUserName.(string)
		databaseUserNm = &tmp
	}

	listWorkRequestsResponse, err := s.Client.ListWorkRequests(context.Background(), listWorkRequestsRequest)
	if listWorkRequestsResponse.Items != nil && len(listWorkRequestsResponse.Items) > 0 {
		// Get the latest work request
		var tmp1 = &listWorkRequestsResponse.Items[0]

		// Get the Sql firewall policyId from the workrequest resources.
		for _, res := range tmp1.Resources {
			if strings.Contains(strings.ToLower(*res.EntityType), "sqlfirewallpolicy") {
				if res.ActionType == oci_data_safe.WorkRequestResourceActionTypeInProgress {
					fmt.Println("IN_PROGRESS Work request found for the given targetId")
					sqlfwpId = res.Identifier
					break
				}
			}
		}

		if sqlfwpId != nil {
			response := oci_data_safe.GetSqlFirewallPolicyResponse{}
			response, _ = s.Client.GetSqlFirewallPolicy(context.Background(),
				oci_data_safe.GetSqlFirewallPolicyRequest{
					SqlFirewallPolicyId: sqlfwpId,
				})
			sqlfwp := &response.SqlFirewallPolicy
			if sqlfwp != nil && *sqlfwp.DbUserName == *databaseUserNm {
				fmt.Println("SQL Firewall policy matches with the given dbUserName")
				workId = tmp1.Id
			}
		}
	}

	if err != nil {
		return err
	}

	if workId != nil {
		return s.getSqlFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate))
	} else {
		return s.GetSqlFirewallPolicyList()
	}
}

func (s *DataSafeSqlFirewallPolicyManagementResourceCrud) GetSqlFirewallPolicyList() error {

	// Filter Policy Deployment for the given targetId to fetch securityPolicyId
	pdRequest := oci_data_safe.ListSecurityPolicyDeploymentsRequest{}
	var securityPolicyId *string
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		pdRequest.CompartmentId = &tmp
		compartmentIdInSubtree := true
		pdRequest.CompartmentIdInSubtree = &compartmentIdInSubtree
		pdRequest.AccessLevel = oci_data_safe.ListSecurityPolicyDeploymentsAccessLevelAccessible
	}
	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		pdRequest.TargetId = &tmp
	}
	pdResponse, err := s.Client.ListSecurityPolicyDeployments(context.Background(), pdRequest)
	if err != nil {
		return err
	}
	if pdResponse.SecurityPolicyDeploymentCollection.Items != nil && len(pdResponse.SecurityPolicyDeploymentCollection.Items) > 0 {
		securityPolicyId = pdResponse.SecurityPolicyDeploymentCollection.Items[0].SecurityPolicyId
	}

	request := oci_data_safe.ListSqlFirewallPoliciesRequest{}
	var sqlFirewallPolicy = new(oci_data_safe.SqlFirewallPolicy)
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
		compartmentIdInSubtree := true
		request.CompartmentIdInSubtree = &compartmentIdInSubtree
		request.AccessLevel = oci_data_safe.ListSqlFirewallPoliciesAccessLevelAccessible
		request.LifecycleState = oci_data_safe.ListSqlFirewallPoliciesLifecycleStateActive
	}

	if securityPolicyId != nil {
		request.SecurityPolicyId = securityPolicyId
	}

	if dbUserName, ok := s.D.GetOkExists("db_user_name"); ok {
		tmp := dbUserName.(string)
		request.DbUserName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListSqlFirewallPolicies(context.Background(), request)
	if err != nil {
		return err
	}
	if response.SqlFirewallPolicyCollection.Items != nil && len(response.SqlFirewallPolicyCollection.Items) > 0 {
		tmp1 := &response.SqlFirewallPolicyCollection.Items[0]
		sqlFirewallPolicy.Id = tmp1.Id
		fmt.Println("Active SQL Firewall policy found")
	} else {
		request.LifecycleState = oci_data_safe.ListSqlFirewallPoliciesLifecycleStateInactive
		response, err := s.Client.ListSqlFirewallPolicies(context.Background(), request)
		if err != nil {
			return err
		}
		if response.SqlFirewallPolicyCollection.Items != nil && len(response.SqlFirewallPolicyCollection.Items) > 0 {
			tmp1 := &response.SqlFirewallPolicyCollection.Items[0]
			sqlFirewallPolicy.Id = tmp1.Id
			fmt.Println("Inactive SQL Firewall policy found")
		} else {
			request.LifecycleState = oci_data_safe.ListSqlFirewallPoliciesLifecycleStateFailed
			response, err := s.Client.ListSqlFirewallPolicies(context.Background(), request)
			if err != nil {
				return err
			}
			if response.SqlFirewallPolicyCollection.Items != nil && len(response.SqlFirewallPolicyCollection.Items) > 0 {
				tmp1 := &response.SqlFirewallPolicyCollection.Items[0]
				sqlFirewallPolicy.Id = tmp1.Id
				fmt.Println("Failed SQL Firewall policy found")
			}
		}
	}

	if sqlFirewallPolicy.Id == nil {
		return nil
	}

	s.D.SetId(*sqlFirewallPolicy.Id)
	return nil
}

func (s *DataSafeSqlFirewallPolicyManagementResourceCrud) Get() error {
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

func (s *DataSafeSqlFirewallPolicyManagementResourceCrud) Update() error {
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

func (s *DataSafeSqlFirewallPolicyManagementResourceCrud) SetData() error {

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

func (s *DataSafeSqlFirewallPolicyManagementResourceCrud) getSqlFirewallPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	sqlFirewallPolicyId, err := sqlFirewallPolicyWaitForWorkRequest(workId, "sqlfirewallpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*sqlFirewallPolicyId)

	return s.Get()
}

func (s *DataSafeSqlFirewallPolicyManagementResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeSqlFirewallPolicyCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.SqlFirewallPolicyId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeSqlFirewallPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSqlFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSqlFirewallPolicyManagementResourceCrud) Delete() error {
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
