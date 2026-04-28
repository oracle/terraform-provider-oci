// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeAuditPolicyManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataSafeAuditPolicyManagementWithContext,
		ReadContext:   readDataSafeAuditPolicyManagementWithContext,
		UpdateContext: updateDataSafeAuditPolicyManagementWithContext,
		DeleteContext: deleteDataSafeAuditPolicyManagementWithContext,
		Schema: map[string]*schema.Schema{
			// Required

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"provision_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"retrieve_from_target_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"audit_conditions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"audit_policy_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enable_conditions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"entity_names": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"entity_selection": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"entity_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"operation_status": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"is_data_safe_service_account_audited": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_priv_users_managed_by_data_safe": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"is_data_safe_service_account_excluded": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed

			"audit_specifications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"audit_policy_category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"audit_policy_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_policy_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"enable_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled_entities": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_created": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_enabled_for_all_users": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_seeded_in_data_safe": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_seeded_in_target": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_view_only": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"partially_enabled_msg": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"lifecycle_details": {
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
			"time_last_provisioned": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_retrieved": {
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

func createDataSafeAuditPolicyManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeAuditPolicyManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	err := sync.GetAuditPolicyWorkReq(ctx)
	err1 := sync.GetWithContext(ctx)
	if err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	if err1 != nil {
		return tfresource.HandleDiagError(m, err1)
	}
	return updateDataSafeAuditPolicyManagementWithContext(ctx, d, m)
}

func readDataSafeAuditPolicyManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeAuditPolicyManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDataSafeAuditPolicyManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeAuditPolicyManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	if _, ok := sync.D.GetOkExists("provision_trigger"); ok {
		err := sync.ProvisionAuditPolicy(ctx)
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("retrieve_from_target_trigger"); ok {
		err := sync.RetrieveAuditPolicies(ctx)
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if err := tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync)); err != nil {
		return err
	}

	return nil
}

func deleteDataSafeAuditPolicyManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type DataSafeAuditPolicyManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.AuditPolicy
	DisableNotFoundRetries bool
}

func (s *DataSafeAuditPolicyManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeAuditPolicyManagementResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_data_safe.UpdateAuditPolicyRequest{}
	response, err := s.Client.UpdateAuditPolicy(ctx, request)
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "auditpolicy") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getAuditPolicyFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeAuditPolicyManagementResourceCrud) getAuditPolicyFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	auditPolicyId, err := auditPolicyWaitForWorkRequest(ctx, workId, "auditpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*auditPolicyId)

	return s.GetWithContext(ctx)
}

func (s *DataSafeAuditPolicyManagementResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetAuditPolicyRequest{}

	tmp := s.D.Id()
	request.AuditPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetAuditPolicy(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.AuditPolicy
	return nil
}

func (s *DataSafeAuditPolicyManagementResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateAuditPolicyRequest{}

	tmp := s.D.Id()
	request.AuditPolicyId = &tmp

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAuditPolicy(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditPolicyFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditPolicyManagementResourceCrud) SetData() error {
	auditConditions := []interface{}{}
	for _, item1 := range s.Res.AuditConditions {
		for _, item2 := range s.Res.AuditSpecifications {
			if item1.AuditPolicyName != nil && item2.AuditPolicyName != nil &&
				*item1.AuditPolicyName == *item2.AuditPolicyName &&
				item2.IsViewOnly != nil && *item2.IsViewOnly == false {
				auditConditions = append(auditConditions, AuditConditionsToMap(item1))
			}
		}
	}
	s.D.Set("audit_conditions", auditConditions)

	auditSpecifications := []interface{}{}
	for _, item := range s.Res.AuditSpecifications {
		auditSpecifications = append(auditSpecifications, AuditSpecificationToMap(item))
	}
	s.D.Set("audit_specifications", auditSpecifications)

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

	if s.Res.IsDataSafeServiceAccountExcluded != nil {
		s.D.Set("is_data_safe_service_account_excluded", *s.Res.IsDataSafeServiceAccountExcluded)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastProvisioned != nil {
		s.D.Set("time_last_provisioned", s.Res.TimeLastProvisioned.String())
	}

	if s.Res.TimeLastRetrieved != nil {
		s.D.Set("time_last_retrieved", s.Res.TimeLastRetrieved.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DataSafeAuditPolicyManagementResourceCrud) GetAuditPolicyWorkReq(ctx context.Context) error {
	var tmpTrue = true
	listWorkRequestsRequest := oci_data_safe.ListWorkRequestsRequest{SortBy: oci_data_safe.ListWorkRequestsSortByEnum("ACCEPTEDTIME"), SortOrder: oci_data_safe.ListWorkRequestsSortOrderEnum("ASC"),
		AccessLevel: oci_data_safe.ListWorkRequestsAccessLevelEnum("ACCESSIBLE"), CompartmentIdInSubtree: &tmpTrue}
	var workId *string
	tmp := "RETRIEVE_POLICY"
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
	if listWorkRequestsResponse.Items != nil && len(listWorkRequestsResponse.Items) > 0 {
		var tmp1 = &listWorkRequestsResponse.Items[0]
		workId = tmp1.Id
	}

	if err != nil {
		return err
	}

	if workId != nil {
		return s.getAuditPolicyFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate))
	} else {
		return s.GetAuditPolicyList(ctx)
	}
}

func (s *DataSafeAuditPolicyManagementResourceCrud) GetAuditPolicyList(ctx context.Context) error {
	request := oci_data_safe.ListAuditPoliciesRequest{}
	var auditPolicy = new(oci_data_safe.AuditPolicy)
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListAuditPolicies(ctx, request)
	if err != nil {
		return err
	}
	if response.AuditPolicyCollection.Items != nil && len(response.AuditPolicyCollection.Items) > 0 {
		tmp1 := &response.AuditPolicyCollection.Items[0]
		auditPolicy.Id = tmp1.Id
	}

	if auditPolicy.Id == nil {
		return nil
	}

	s.D.SetId(*auditPolicy.Id)
	return nil
}

func (s *DataSafeAuditPolicyManagementResourceCrud) ProvisionAuditPolicy(ctx context.Context) error {
	request := oci_data_safe.ProvisionAuditPolicyRequest{}

	idTmp := s.D.Id()
	request.AuditPolicyId = &idTmp

	if isDataSafeServiceAccountExcluded, ok := s.D.GetOkExists("is_data_safe_service_account_excluded"); ok {
		tmp := isDataSafeServiceAccountExcluded.(bool)
		request.IsDataSafeServiceAccountExcluded = &tmp
	}
	if provisionAuditConditions, ok := s.D.GetOkExists("audit_conditions"); ok {
		interfaces := provisionAuditConditions.([]interface{})
		tmp := make([]oci_data_safe.ProvisionAuditConditions, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "audit_conditions", stateDataIndex)
			converted, err := s.mapToProvisionAuditConditions(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("audit_conditions") {
			request.ProvisionAuditConditions = tmp
		}
	}
	fmt.Printf(" request = %+v\n", request)
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ProvisionAuditPolicy(ctx, request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("provision_trigger")
	s.D.Set("provision_trigger", val)

	return nil
}

func (s *DataSafeAuditPolicyManagementResourceCrud) RetrieveAuditPolicies(ctx context.Context) error {
	request := oci_data_safe.RetrieveAuditPoliciesRequest{}

	idTmp := s.D.Id()
	request.AuditPolicyId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.RetrieveAuditPolicies(ctx, request)

	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("retrieve_from_target_trigger")
	s.D.Set("retrieve_from_target_trigger", val)

	//s.Res = &response.AuditPolicy
	return nil
}

func (s *DataSafeAuditPolicyManagementResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeAuditPolicyCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AuditPolicyId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeAuditPolicyCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditPolicyFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditPolicyManagementResourceCrud) mapToProvisionAuditConditions(fieldKeyFormat string) (oci_data_safe.ProvisionAuditConditions, error) {
	result := oci_data_safe.ProvisionAuditConditions{}

	if auditPolicyName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "audit_policy_name")); ok {
		tmp := auditPolicyName.(string)
		result.AuditPolicyName = &tmp
	}

	if enableConditions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_conditions")); ok {
		interfaces := enableConditions.([]interface{})
		tmp := make([]oci_data_safe.EnableConditions, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "enable_conditions"), stateDataIndex)
			converted, err := s.mapToEnableConditions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "enable_conditions")) {
			result.EnableConditions = tmp
		}
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if isPrivUsersManagedByDataSafe, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_priv_users_managed_by_data_safe")); ok {
		tmp := isPrivUsersManagedByDataSafe.(bool)
		result.IsPrivUsersManagedByDataSafe = &tmp
	}

	return result, nil
}
func (s *DataSafeAuditPolicyManagementResourceCrud) mapToEnableConditions(fieldKeyFormat string) (oci_data_safe.EnableConditions, error) {
	result := oci_data_safe.EnableConditions{}

	if entityNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_names")); ok {
		interfaces := entityNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "entity_names")) {
			result.EntityNames = tmp
		}
	}

	if entitySelection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_selection")); ok {
		result.EntitySelection = oci_data_safe.EnableConditionsEntitySelectionEnum(entitySelection.(string))
	}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_data_safe.EnableConditionsEntityTypeEnum(entityType.(string))
	}

	if operationStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_status")); ok {
		result.OperationStatus = oci_data_safe.EnableConditionsOperationStatusEnum(operationStatus.(string))
	}

	return result, nil
}
