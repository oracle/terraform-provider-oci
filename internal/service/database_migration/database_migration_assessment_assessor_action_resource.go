// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationAssessmentAssessorActionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseMigrationAssessmentAssessorAction,
		Read:     readDatabaseMigrationAssessmentAssessorAction,
		Delete:   deleteDatabaseMigrationAssessmentAssessorAction,
		Schema: map[string]*schema.Schema{
			// Required
			"assessment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"assessor_action": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"assessor_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"items": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional

			// Computed
		},
	}
}

func createDatabaseMigrationAssessmentAssessorAction(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentAssessorActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseMigrationAssessmentAssessorAction(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseMigrationAssessmentAssessorAction(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseMigrationAssessmentAssessorActionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_migration.DatabaseMigrationClient
	DisableNotFoundRetries bool
}

func (s *DatabaseMigrationAssessmentAssessorActionResourceCrud) ID() string {
	return s.D.Id()
}

func (s *DatabaseMigrationAssessmentAssessorActionResourceCrud) Get() error {
	return nil
}

func (s *DatabaseMigrationAssessmentAssessorActionResourceCrud) Create() error {
	request := oci_database_migration.PerformAssessorActionRequest{}

	if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := assessmentId.(string)
		request.AssessmentId = &tmp
	}

	if assessorAction, ok := s.D.GetOkExists("assessor_action"); ok {
		tmp := assessorAction.(string)
		request.AssessorAction = &tmp
	}

	if assessorName, ok := s.D.GetOkExists("assessor_name"); ok {
		tmp := assessorName.(string)
		request.AssessorName = &tmp
	}

	if items, ok := s.D.GetOkExists("items"); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_database_migration.AssessorCheckNameValue, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "items", stateDataIndex)
			converted, err := s.mapToAssessorCheckNameValue(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("items") {
			request.Items = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.PerformAssessorAction(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Guard against nil work request ID to avoid marshalling a nil pointer in subsequent requests.
	if workId == nil {
		if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
			s.D.SetId(assessmentId.(string))
		}
		// No work request to poll; treat as successful since action was accepted without a WR.
		return nil
	}
	workRequestResponse := oci_database_migration.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_database_migration.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "assessment") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getAssessmentAssessorActionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration"), oci_database_migration.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseMigrationAssessmentAssessorActionResourceCrud) getAssessmentAssessorActionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_migration.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	assessmentAssessorActionId, err := assessmentAssessorActionWaitForWorkRequest(workId, "assessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*assessmentAssessorActionId)

	return s.Get()
}

func assessmentAssessorActionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "database_migration", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_database_migration.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func assessmentAssessorActionWaitForWorkRequest(wId *string, entityType string, action oci_database_migration.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_migration.DatabaseMigrationClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_migration")
	retryPolicy.ShouldRetryOperation = assessmentAssessorActionWorkRequestShouldRetryFunc(timeout)

	response := oci_database_migration.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_database_migration.OperationStatusInProgress),
			string(oci_database_migration.OperationStatusAccepted),
			string(oci_database_migration.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_database_migration.OperationStatusSucceeded),
			string(oci_database_migration.OperationStatusFailed),
			string(oci_database_migration.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_database_migration.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_database_migration.OperationStatusFailed || response.Status == oci_database_migration.OperationStatusCanceled {
		return nil, getErrorFromDatabaseMigrationAssessmentAssessorActionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseMigrationAssessmentAssessorActionWorkRequest(client *oci_database_migration.DatabaseMigrationClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_migration.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_database_migration.ListWorkRequestErrorsRequest{
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

func (s *DatabaseMigrationAssessmentAssessorActionResourceCrud) SetData() error {
	return nil
}

func (s *DatabaseMigrationAssessmentAssessorActionResourceCrud) mapToAssessorCheckNameValue(fieldKeyFormat string) (oci_database_migration.AssessorCheckNameValue, error) {
	result := oci_database_migration.AssessorCheckNameValue{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AssessorCheckNameValueToMap(obj oci_database_migration.AssessorCheckNameValue) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
