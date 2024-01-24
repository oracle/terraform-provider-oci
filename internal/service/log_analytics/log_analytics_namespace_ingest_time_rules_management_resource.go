// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceIngestTimeRulesManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsNamespaceIngestTimeRulesManagement,
		Read:     readLogAnalyticsNamespaceIngestTimeRulesManagement,
		Update:   updateLogAnalyticsNamespaceIngestTimeRulesManagement,
		Delete:   deleteLogAnalyticsNamespaceIngestTimeRulesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"ingest_time_rule_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_ingest_time_rule": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional

			// Computed
		},
	}
}

func createLogAnalyticsNamespaceIngestTimeRulesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceIngestTimeRulesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.Res = &LogAnalyticsNamespaceIngestTimeRulesManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsNamespaceIngestTimeRulesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateLogAnalyticsNamespaceIngestTimeRulesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceIngestTimeRulesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.Res = &LogAnalyticsNamespaceIngestTimeRulesManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteLogAnalyticsNamespaceIngestTimeRulesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceIngestTimeRulesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.Res = &LogAnalyticsNamespaceIngestTimeRulesManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsNamespaceIngestTimeRulesManagementResponse struct {
	enableResponse  *oci_log_analytics.EnableIngestTimeRuleResponse
	disableResponse *oci_log_analytics.DisableIngestTimeRuleResponse
}

type LogAnalyticsNamespaceIngestTimeRulesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *LogAnalyticsNamespaceIngestTimeRulesManagementResponse
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsNamespaceIngestTimeRulesManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceIngestTimeRulesManagementResource-", LogAnalyticsNamespaceIngestTimeRulesManagementResource(), s.D)
}

func (s *LogAnalyticsNamespaceIngestTimeRulesManagementResourceCrud) Create() error {
	var operation bool
	var namespaceName string
	if enableOperation, ok := s.D.GetOkExists("enable_ingest_time_rule"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_log_analytics.EnableIngestTimeRuleRequest{}

		if ingestTimeRuleId, ok := s.D.GetOkExists("ingest_time_rule_id"); ok {
			tmp := ingestTimeRuleId.(string)
			request.IngestTimeRuleId = &tmp
		}

		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			namespaceName = namespace.(string)
			request.NamespaceName = &namespaceName
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

		response, err := s.Client.EnableIngestTimeRule(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getNamespaceIngestTimeRulesManagementFromWorkRequest(&namespaceName, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics"), oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeEnableIngestTimeRule, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_log_analytics.DisableIngestTimeRuleRequest{}

	if ingestTimeRuleId, ok := s.D.GetOkExists("ingest_time_rule_id"); ok {
		tmp := ingestTimeRuleId.(string)
		request.IngestTimeRuleId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		namespaceName = namespace.(string)
		request.NamespaceName = &namespaceName
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.DisableIngestTimeRule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getNamespaceIngestTimeRulesManagementFromWorkRequest(&namespaceName, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics"), oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeDisableIngestTimeRule, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *LogAnalyticsNamespaceIngestTimeRulesManagementResourceCrud) getNamespaceIngestTimeRulesManagementFromWorkRequest(namespaceName *string, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := namespaceIngestTimeRulesManagementWaitForWorkRequest(namespaceName, workId, "log_analytics",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func namespaceIngestTimeRulesManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "log_analytics", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_log_analytics.GetConfigWorkRequestResponse); ok {
			return workRequestResponse.LogAnalyticsConfigWorkRequest.TimeFinished == nil
		}
		return false
	}
}

func namespaceIngestTimeRulesManagementWaitForWorkRequest(namespaceName *string, wId *string, entityType string, action oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_log_analytics.LogAnalyticsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "log_analytics")
	retryPolicy.ShouldRetryOperation = namespaceIngestTimeRulesManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_log_analytics.GetConfigWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateInProgress),
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateAccepted),
		},
		Target: []string{
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateSucceeded),
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateFailed),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetConfigWorkRequest(context.Background(),
				oci_log_analytics.GetConfigWorkRequestRequest{
					NamespaceName: namespaceName,
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.LogAnalyticsConfigWorkRequest
			return wr, string(wr.LifecycleState), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if response.LogAnalyticsConfigWorkRequest.LifecycleState == oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateFailed {
		return nil, getErrorFromLogAnalyticsNamespaceIngestTimeRulesManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return nil, nil
}

func getErrorFromLogAnalyticsNamespaceIngestTimeRulesManagementWorkRequest(client *oci_log_analytics.LogAnalyticsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeEnum) error {
	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s", *workId, entityType, action)
	return workRequestErr
}

func (s *LogAnalyticsNamespaceIngestTimeRulesManagementResourceCrud) Update() error {
	var operation bool
	var namespaceName string

	if enableOperation, ok := s.D.GetOkExists("enable_ingest_time_rule"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_log_analytics.EnableIngestTimeRuleRequest{}

		if ingestTimeRuleId, ok := s.D.GetOkExists("ingest_time_rule_id"); ok {
			tmp := ingestTimeRuleId.(string)
			request.IngestTimeRuleId = &tmp
		}

		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			namespaceName = namespace.(string)
			request.NamespaceName = &namespaceName
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

		response, err := s.Client.EnableIngestTimeRule(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getNamespaceIngestTimeRulesManagementFromWorkRequest(&namespaceName, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics"), oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeEnableIngestTimeRule, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_log_analytics.DisableIngestTimeRuleRequest{}

	if ingestTimeRuleId, ok := s.D.GetOkExists("ingest_time_rule_id"); ok {
		tmp := ingestTimeRuleId.(string)
		request.IngestTimeRuleId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		namespaceName = namespace.(string)
		request.NamespaceName = &namespaceName
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.DisableIngestTimeRule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getNamespaceIngestTimeRulesManagementFromWorkRequest(&namespaceName, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics"), oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeDisableIngestTimeRule, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *LogAnalyticsNamespaceIngestTimeRulesManagementResourceCrud) Delete() error {
	var operation bool
	var namespaceName string

	if enableOperation, ok := s.D.GetOkExists("enable_ingest_time_rule"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_log_analytics.DisableIngestTimeRuleRequest{}

	if ingestTimeRuleId, ok := s.D.GetOkExists("ingest_time_rule_id"); ok {
		tmp := ingestTimeRuleId.(string)
		request.IngestTimeRuleId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		namespaceName = namespace.(string)
		request.NamespaceName = &namespaceName
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.DisableIngestTimeRule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getNamespaceIngestTimeRulesManagementFromWorkRequest(&namespaceName, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics"), oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeDisableIngestTimeRule, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *LogAnalyticsNamespaceIngestTimeRulesManagementResourceCrud) SetData() error {
	return nil
}
