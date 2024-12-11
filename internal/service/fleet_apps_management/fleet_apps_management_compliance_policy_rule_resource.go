// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementCompliancePolicyRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementCompliancePolicyRule,
		Read:     readFleetAppsManagementCompliancePolicyRule,
		Update:   updateFleetAppsManagementCompliancePolicyRule,
		Delete:   deleteFleetAppsManagementCompliancePolicyRule,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"patch_selection": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"selection_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"PATCH_LEVEL",
								"PATCH_NAME",
								"PATCH_RELEASE_DATE",
							}, true),
						},

						// Optional
						"days_since_release": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"patch_level": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"patch_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"patch_type": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"product_version": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"version": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"is_applicable_for_all_higher_versions": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Optional
			"compliance_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"grace_period": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFleetAppsManagementCompliancePolicyRule(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCompliancePolicyRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementCompliancePolicyRule(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCompliancePolicyRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementCompliancePolicyRule(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCompliancePolicyRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementCompliancePolicyRule(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCompliancePolicyRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementCompliancePolicyRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementAdminClient
	FleetClient            *oci_fleet_apps_management.FleetAppsManagementClient
	Res                    *oci_fleet_apps_management.CompliancePolicyRule
	DisableNotFoundRetries bool
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fleet_apps_management.CompliancePolicyRuleLifecycleStateCreating),
	}
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.CompliancePolicyRuleLifecycleStateActive),
	}
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.CompliancePolicyRuleLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.CompliancePolicyRuleLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateCompliancePolicyRuleRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compliancePolicyId, ok := s.D.GetOkExists("compliance_policy_id"); ok {
		tmp := compliancePolicyId.(string)
		request.CompliancePolicyId = &tmp
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

	if gracePeriod, ok := s.D.GetOkExists("grace_period"); ok {
		tmp := gracePeriod.(string)
		request.GracePeriod = &tmp
	}

	if patchSelection, ok := s.D.GetOkExists("patch_selection"); ok {
		if tmpList := patchSelection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patch_selection", 0)
			tmp, err := s.mapToPatchSelectionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PatchSelection = tmp
		}
	}

	if patchType, ok := s.D.GetOkExists("patch_type"); ok {
		interfaces := patchType.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("patch_type") {
			request.PatchType = tmp
		}
	}

	if productVersion, ok := s.D.GetOkExists("product_version"); ok {
		if tmpList := productVersion.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "product_version", 0)
			tmp, err := s.mapToProductVersionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ProductVersion = &tmp
		}
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		interfaces := severity.([]interface{})
		tmp := make([]oci_fleet_apps_management.ComplianceRuleSeverityEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				// Assert the interface{} to string, as Terraform stores schema values as strings
				strValue := interfaces[i].(string)

				// Convert the string to SeverityEnum and validate it
				switch oci_fleet_apps_management.ComplianceRuleSeverityEnum(strValue) {
				case oci_fleet_apps_management.ComplianceRuleSeverityCritical, oci_fleet_apps_management.ComplianceRuleSeverityHigh, oci_fleet_apps_management.ComplianceRuleSeverityMedium, oci_fleet_apps_management.ComplianceRuleSeverityLow:
					tmp[i] = oci_fleet_apps_management.ComplianceRuleSeverityEnum(strValue) // Assign the valid SeverityEnum
				default:
					return fmt.Errorf("invalid severity level: %s", strValue) // Return error for invalid values
				}
				// tmp[i] = interfaces[i].(oci_fleet_apps_management.ComplianceRuleSeverityEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("severity") {
			request.Severity = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateCompliancePolicyRule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getCompliancePolicyRuleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) getCompliancePolicyRuleFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	compliancePolicyRuleId, err := compliancePolicyRuleWaitForWorkRequest(workId, "compliancepolicyrule",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.FleetClient)

	if err != nil {
		return err
	}
	s.D.SetId(*compliancePolicyRuleId)

	return s.Get()
}

func compliancePolicyRuleWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fleet_apps_management", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fleet_apps_management.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func compliancePolicyRuleWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = compliancePolicyRuleWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_apps_management.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_fleet_apps_management.OperationStatusInProgress),
			string(oci_fleet_apps_management.OperationStatusAccepted),
			string(oci_fleet_apps_management.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_fleet_apps_management.OperationStatusSucceeded),
			string(oci_fleet_apps_management.OperationStatusFailed),
			string(oci_fleet_apps_management.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			// Proceed with GetWorkRequest call
			response, err = client.GetWorkRequest(context.Background(),
				oci_fleet_apps_management.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_fleet_apps_management.OperationStatusFailed || response.Status == oci_fleet_apps_management.OperationStatusCanceled {
		return nil, getErrorFromFleetAppsManagementCompliancePolicyRuleWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementCompliancePolicyRuleWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_fleet_apps_management.ListWorkRequestErrorsRequest{
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

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetCompliancePolicyRuleRequest{}

	tmp := s.D.Id()
	request.CompliancePolicyRuleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetCompliancePolicyRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CompliancePolicyRule
	return nil
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdateCompliancePolicyRuleRequest{}

	tmp := s.D.Id()
	request.CompliancePolicyRuleId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gracePeriod, ok := s.D.GetOkExists("grace_period"); ok {
		tmp := gracePeriod.(string)
		request.GracePeriod = &tmp
	}

	if patchSelection, ok := s.D.GetOkExists("patch_selection"); ok {
		if tmpList := patchSelection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patch_selection", 0)
			tmp, err := s.mapToPatchSelectionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PatchSelection = tmp
		}
	}

	if patchType, ok := s.D.GetOkExists("patch_type"); ok {
		interfaces := patchType.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("patch_type") {
			request.PatchType = tmp
		}
	}

	if productVersion, ok := s.D.GetOkExists("product_version"); ok {
		if tmpList := productVersion.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "product_version", 0)
			tmp, err := s.mapToProductVersionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ProductVersion = &tmp
		}
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		interfaces := severity.([]interface{})
		tmp := make([]oci_fleet_apps_management.ComplianceRuleSeverityEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				// Assert the interface{} to string, as Terraform stores schema values as strings
				strValue := interfaces[i].(string)

				// Convert the string to SeverityEnum and validate it
				switch oci_fleet_apps_management.ComplianceRuleSeverityEnum(strValue) {
				case oci_fleet_apps_management.ComplianceRuleSeverityCritical, oci_fleet_apps_management.ComplianceRuleSeverityHigh, oci_fleet_apps_management.ComplianceRuleSeverityMedium, oci_fleet_apps_management.ComplianceRuleSeverityLow:
					tmp[i] = oci_fleet_apps_management.ComplianceRuleSeverityEnum(strValue) // Assign the valid SeverityEnum
				default:
					return fmt.Errorf("invalid severity level: %s", strValue) // Return error for invalid values
				}
				// tmp[i] = interfaces[i].(oci_fleet_apps_management.ComplianceRuleSeverityEnum)
			}
		}

		if len(tmp) != 0 || s.D.HasChange("severity") {
			request.Severity = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateCompliancePolicyRule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getCompliancePolicyRuleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteCompliancePolicyRuleRequest{}

	tmp := s.D.Id()
	request.CompliancePolicyRuleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.DeleteCompliancePolicyRule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := compliancePolicyRuleWaitForWorkRequest(workId, "compliancepolicyrule",
		oci_fleet_apps_management.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.FleetClient)
	return delWorkRequestErr
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CompliancePolicyId != nil {
		s.D.Set("compliance_policy_id", *s.Res.CompliancePolicyId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GracePeriod != nil {
		s.D.Set("grace_period", *s.Res.GracePeriod)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PatchSelection != nil {
		patchSelectionArray := []interface{}{}
		if patchSelectionMap := PatchSelectionDetailsToMap(&s.Res.PatchSelection); patchSelectionMap != nil {
			patchSelectionArray = append(patchSelectionArray, patchSelectionMap)
		}
		s.D.Set("patch_selection", patchSelectionArray)
	} else {
		s.D.Set("patch_selection", nil)
	}

	s.D.Set("patch_type", s.Res.PatchType)

	if s.Res.ProductVersion != nil {
		s.D.Set("product_version", []interface{}{ProductVersionDetailsToMap(s.Res.ProductVersion)})
	} else {
		s.D.Set("product_version", nil)
	}

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func CompliancePolicyRuleSummaryToMap(obj oci_fleet_apps_management.CompliancePolicyRuleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CompliancePolicyId != nil {
		result["compliance_policy_id"] = string(*obj.CompliancePolicyId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.GracePeriod != nil {
		result["grace_period"] = string(*obj.GracePeriod)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.PatchSelection != nil {
		patchSelectionArray := []interface{}{}
		if patchSelectionMap := PatchSelectionDetailsToMap(&obj.PatchSelection); patchSelectionMap != nil {
			patchSelectionArray = append(patchSelectionArray, patchSelectionMap)
		}
		result["patch_selection"] = patchSelectionArray
	}

	result["patch_type"] = obj.PatchType

	if obj.ProductVersion != nil {
		result["product_version"] = []interface{}{ProductVersionDetailsToMap(obj.ProductVersion)}
	}

	result["severity"] = obj.Severity

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) mapToPatchSelectionDetails(fieldKeyFormat string) (oci_fleet_apps_management.PatchSelectionDetails, error) {
	var baseObject oci_fleet_apps_management.PatchSelectionDetails
	//discriminator
	selectionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection_type"))
	var selectionType string
	if ok {
		selectionType = selectionTypeRaw.(string)
	} else {
		selectionType = "" // default value
	}
	switch strings.ToLower(selectionType) {
	case strings.ToLower("PATCH_LEVEL"):
		details := oci_fleet_apps_management.PatchLevelSelectionDetails{}
		if patchLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patch_level")); ok {
			details.PatchLevel = oci_fleet_apps_management.PatchLevelSelectionDetailsPatchLevelEnum(patchLevel.(string))
		}
		baseObject = details
	case strings.ToLower("PATCH_NAME"):
		details := oci_fleet_apps_management.PatchNameSelectionDetails{}
		if patchName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patch_name")); ok {
			tmp := patchName.(string)
			details.PatchName = &tmp
		}
		baseObject = details
	case strings.ToLower("PATCH_RELEASE_DATE"):
		details := oci_fleet_apps_management.PatchReleaseDateSelectionDetails{}
		if daysSinceRelease, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days_since_release")); ok {
			tmp := daysSinceRelease.(int)
			details.DaysSinceRelease = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown selection_type '%v' was specified", selectionType)
	}
	return baseObject, nil
}

func PatchSelectionDetailsToMap(obj *oci_fleet_apps_management.PatchSelectionDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.PatchLevelSelectionDetails:
		result["selection_type"] = "PATCH_LEVEL"

		result["patch_level"] = string(v.PatchLevel)
	case oci_fleet_apps_management.PatchNameSelectionDetails:
		result["selection_type"] = "PATCH_NAME"

		if v.PatchName != nil {
			result["patch_name"] = string(*v.PatchName)
		}
	case oci_fleet_apps_management.PatchReleaseDateSelectionDetails:
		result["selection_type"] = "PATCH_RELEASE_DATE"

		if v.DaysSinceRelease != nil {
			result["days_since_release"] = int(*v.DaysSinceRelease)
		}
	default:
		log.Printf("[WARN] Received 'selection_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementCompliancePolicyRuleResourceCrud) mapToProductVersionDetails(fieldKeyFormat string) (oci_fleet_apps_management.ProductVersionDetails, error) {
	result := oci_fleet_apps_management.ProductVersionDetails{}

	if isApplicableForAllHigherVersions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_applicable_for_all_higher_versions")); ok {
		tmp := isApplicableForAllHigherVersions.(bool)
		result.IsApplicableForAllHigherVersions = &tmp
	}

	if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
		tmp := version.(string)
		result.Version = &tmp
	}

	return result, nil
}

func ProductVersionDetailsToMap(obj *oci_fleet_apps_management.ProductVersionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsApplicableForAllHigherVersions != nil {
		result["is_applicable_for_all_higher_versions"] = bool(*obj.IsApplicableForAllHigherVersions)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
