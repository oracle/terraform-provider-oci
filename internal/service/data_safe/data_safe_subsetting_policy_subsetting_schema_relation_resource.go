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

func DataSafeSubsettingPolicySubsettingSchemaRelationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataSafeSubsettingPolicySubsettingSchemaRelationWithContext,
		ReadContext:   readDataSafeSubsettingPolicySubsettingSchemaRelationWithContext,
		DeleteContext: deleteDataSafeSubsettingPolicySubsettingSchemaRelationWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"child_columns": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"child_object_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"child_schema_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parent_columns": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"parent_object_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parent_schema_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subsetting_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"child_object_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"parent_object_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"relation_type": {
				Type:     schema.TypeString,
				Computed: true,
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

func createDataSafeSubsettingPolicySubsettingSchemaRelationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDataSafeSubsettingPolicySubsettingSchemaRelationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func deleteDataSafeSubsettingPolicySubsettingSchemaRelationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SubsettingSchemaRelation
	DisableNotFoundRetries bool
}

func (s *DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud) ID() string {
	policyID, policyOK := s.D.GetOk("subsetting_policy_id")
	relationKey, relationOK := s.D.GetOk("key")
	if !policyOK || !relationOK {
		return s.D.Id()
	}
	return GetSubsettingPolicySubsettingSchemaRelationCompositeId(policyID.(string), relationKey.(string))
}

func (s *DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud) CreateWithContext(ctx context.Context) error {
	request, err := s.createSubsettingSchemaRelationRequest()
	if err != nil {
		return err
	}

	response, err := s.Client.CreateSubsettingSchemaRelation(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSubsettingPolicySubsettingSchemaRelationFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud) createSubsettingSchemaRelationRequest() (oci_data_safe.CreateSubsettingSchemaRelationRequest, error) {
	request := oci_data_safe.CreateSubsettingSchemaRelationRequest{}

	if childColumns, ok := s.D.GetOkExists("child_columns"); ok {
		interfaces := childColumns.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("child_columns") {
			request.ChildColumns = tmp
		}
	}

	if childObjectName, ok := s.D.GetOkExists("child_object_name"); ok {
		tmp := childObjectName.(string)
		request.ChildObjectName = &tmp
	}

	if childSchemaName, ok := s.D.GetOkExists("child_schema_name"); ok {
		tmp := childSchemaName.(string)
		request.ChildSchemaName = &tmp
	}

	if childObjectKey, ok := s.D.GetOkExists("child_object_key"); ok {
		tmp := childObjectKey.(string)
		request.ChildObjectKey = &tmp
	}

	if parentColumns, ok := s.D.GetOkExists("parent_columns"); ok {
		interfaces := parentColumns.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("parent_columns") {
			request.ParentColumns = tmp
		}
	}

	if parentObjectName, ok := s.D.GetOkExists("parent_object_name"); ok {
		tmp := parentObjectName.(string)
		request.ParentObjectName = &tmp
	}

	if parentSchemaName, ok := s.D.GetOkExists("parent_schema_name"); ok {
		tmp := parentSchemaName.(string)
		request.ParentSchemaName = &tmp
	}

	if parentObjectKey, ok := s.D.GetOkExists("parent_object_key"); ok {
		tmp := parentObjectKey.(string)
		request.ParentObjectKey = &tmp
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")
	return request, nil
}

func (s *DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud) getSubsettingPolicySubsettingSchemaRelationFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	subsettingPolicySubsettingSchemaRelationId, err := subsettingPolicySubsettingSchemaRelationWaitForWorkRequest(ctx, workId, "subsettingpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, subsettingPolicySubsettingSchemaRelationId)
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
	// The work request identifies the parent policy rather than the newly
	// created schema relation. Resolve the relation key from the policy's
	// APP_DEFINED relations before performing the first read.
	subsettingSchemaRelationKey, err := s.findCreatedSubsettingSchemaRelationKey(ctx)
	if err != nil {
		return err
	}
	s.D.Set("key", subsettingSchemaRelationKey)
	s.D.SetId(GetSubsettingPolicySubsettingSchemaRelationCompositeId(s.D.Get("subsetting_policy_id").(string), subsettingSchemaRelationKey))

	return s.GetWithContext(ctx)
}

func (s *DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud) findCreatedSubsettingSchemaRelationKey(ctx context.Context) (string, error) {
	policyID := s.D.Get("subsetting_policy_id").(string)
	request := oci_data_safe.ListSubsettingSchemaRelationsRequest{
		SubsettingPolicyId: &policyID,
		RelationType:       oci_data_safe.ListSubsettingSchemaRelationsRelationTypeAppDefined,
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListSubsettingSchemaRelations(ctx, request)
	if err != nil {
		return "", err
	}
	relations := response.Items
	for response.OpcNextPage != nil {
		request.Page = response.OpcNextPage
		response, err = s.Client.ListSubsettingSchemaRelations(ctx, request)
		if err != nil {
			return "", err
		}
		relations = append(relations, response.Items...)
	}

	parentColumnValues := normalizeSubsettingIdentifierList(subsettingStringList(s.D.Get("parent_columns")))
	childColumnValues := normalizeSubsettingIdentifierList(subsettingStringList(s.D.Get("child_columns")))
	parentObjectName := normalizeSubsettingIdentifier(s.D.Get("parent_object_name").(string))
	parentSchemaName := normalizeSubsettingIdentifier(s.D.Get("parent_schema_name").(string))
	childObjectName := normalizeSubsettingIdentifier(s.D.Get("child_object_name").(string))
	childSchemaName := normalizeSubsettingIdentifier(s.D.Get("child_schema_name").(string))

	for _, relation := range relations {
		if relation.Key != nil &&
			relation.ParentSchemaName != nil && normalizeSubsettingIdentifier(*relation.ParentSchemaName) == parentSchemaName &&
			relation.ParentObjectName != nil && normalizeSubsettingIdentifier(*relation.ParentObjectName) == parentObjectName &&
			relation.ChildSchemaName != nil && normalizeSubsettingIdentifier(*relation.ChildSchemaName) == childSchemaName &&
			relation.ChildObjectName != nil && normalizeSubsettingIdentifier(*relation.ChildObjectName) == childObjectName &&
			reflect.DeepEqual(normalizeSubsettingIdentifierList(relation.ParentColumns), parentColumnValues) &&
			reflect.DeepEqual(normalizeSubsettingIdentifierList(relation.ChildColumns), childColumnValues) {
			if parentObjectKey, ok := s.D.GetOk("parent_object_key"); ok &&
				(relation.ParentObjectKey == nil || strings.TrimSpace(*relation.ParentObjectKey) != strings.TrimSpace(parentObjectKey.(string))) {
				continue
			}
			if childObjectKey, ok := s.D.GetOk("child_object_key"); ok &&
				(relation.ChildObjectKey == nil || strings.TrimSpace(*relation.ChildObjectKey) != strings.TrimSpace(childObjectKey.(string))) {
				continue
			}
			return *relation.Key, nil
		}
	}

	return "", fmt.Errorf("unable to determine the key of the newly-created APP_DEFINED schema relation in policy %s", policyID)
}

func subsettingPolicySubsettingSchemaRelationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func subsettingPolicySubsettingSchemaRelationWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = subsettingPolicySubsettingSchemaRelationWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeSubsettingPolicySubsettingSchemaRelationWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSubsettingPolicySubsettingSchemaRelationWorkRequest(ctx context.Context, client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetSubsettingSchemaRelationRequest{}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if subsettingSchemaRelationKey, ok := s.D.GetOkExists("key"); ok {
		tmp := subsettingSchemaRelationKey.(string)
		request.SubsettingSchemaRelationKey = &tmp
	}

	subsettingPolicyId, subsettingSchemaRelationKey, err := parseSubsettingPolicySubsettingSchemaRelationCompositeId(s.D.Id())
	if err == nil {
		request.SubsettingPolicyId = &subsettingPolicyId
		request.SubsettingSchemaRelationKey = &subsettingSchemaRelationKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSubsettingSchemaRelation(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.SubsettingSchemaRelation
	return nil
}

func (s *DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_data_safe.DeleteSubsettingSchemaRelationRequest{}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if subsettingSchemaRelationKey, ok := s.D.GetOkExists("key"); ok {
		tmp := subsettingSchemaRelationKey.(string)
		request.SubsettingSchemaRelationKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteSubsettingSchemaRelation(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := subsettingPolicySubsettingSchemaRelationWaitForWorkRequest(ctx, workId, "subsettingpolicy",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud) SetData() error {

	subsettingPolicyId, subsettingSchemaRelationKey, err := parseSubsettingPolicySubsettingSchemaRelationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("subsetting_policy_id", &subsettingPolicyId)
		s.D.Set("key", &subsettingSchemaRelationKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("child_columns", s.Res.ChildColumns)

	if s.Res.ChildObjectKey != nil {
		s.D.Set("child_object_key", *s.Res.ChildObjectKey)
	}

	if s.Res.ChildObjectName != nil {
		s.D.Set("child_object_name", *s.Res.ChildObjectName)
	}

	if s.Res.ChildSchemaName != nil {
		s.D.Set("child_schema_name", *s.Res.ChildSchemaName)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("parent_columns", s.Res.ParentColumns)

	if s.Res.ParentObjectKey != nil {
		s.D.Set("parent_object_key", *s.Res.ParentObjectKey)
	}

	if s.Res.ParentObjectName != nil {
		s.D.Set("parent_object_name", *s.Res.ParentObjectName)
	}

	if s.Res.ParentSchemaName != nil {
		s.D.Set("parent_schema_name", *s.Res.ParentSchemaName)
	}

	s.D.Set("relation_type", s.Res.RelationType)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetSubsettingPolicySubsettingSchemaRelationCompositeId(subsettingPolicyId string, subsettingSchemaRelationKey string) string {
	subsettingPolicyId = url.PathEscape(subsettingPolicyId)
	subsettingSchemaRelationKey = url.PathEscape(subsettingSchemaRelationKey)
	compositeId := "subsettingPolicies/" + subsettingPolicyId + "/subsettingSchemaRelations/" + subsettingSchemaRelationKey
	return compositeId
}

func parseSubsettingPolicySubsettingSchemaRelationCompositeId(compositeId string) (subsettingPolicyId string, subsettingSchemaRelationKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("subsettingPolicies/.*/subsettingSchemaRelations/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	subsettingPolicyId, _ = url.PathUnescape(parts[1])
	subsettingSchemaRelationKey, _ = url.PathUnescape(parts[3])

	return
}

func SubsettingSchemaRelationSummaryToMap(obj oci_data_safe.SubsettingSchemaRelationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["child_columns"] = obj.ChildColumns

	if obj.ChildObjectKey != nil {
		result["child_object_key"] = string(*obj.ChildObjectKey)
	}

	if obj.ChildObjectName != nil {
		result["child_object_name"] = string(*obj.ChildObjectName)
	}

	if obj.ChildSchemaName != nil {
		result["child_schema_name"] = string(*obj.ChildSchemaName)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["parent_columns"] = obj.ParentColumns

	if obj.ParentObjectKey != nil {
		result["parent_object_key"] = string(*obj.ParentObjectKey)
	}

	if obj.ParentObjectName != nil {
		result["parent_object_name"] = string(*obj.ParentObjectName)
	}

	if obj.ParentSchemaName != nil {
		result["parent_schema_name"] = string(*obj.ParentSchemaName)
	}

	result["relation_type"] = string(obj.RelationType)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
