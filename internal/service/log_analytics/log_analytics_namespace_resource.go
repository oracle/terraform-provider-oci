// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"
)

func LogAnalyticsNamespaceResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("2m"),
			Update: tfresource.GetTimeoutDuration("2m"),
		},
		Create: createLogAnalyticsNamespace,
		Read:   readLogAnalyticsNamespace,
		Update: updateLogAnalyticsNamespace,
		Delete: deleteLogAnalyticsNamespace,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_onboarded": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createLogAnalyticsNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateLogAnalyticsNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLogAnalyticsNamespace(d *schema.ResourceData, m interface{}) error {
	return nil
}

type LogAnalyticsNamespaceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.NamespaceSummary
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsNamespaceResourceCrud) ID() string {
	return s.D.Get("namespace").(string)
}

func (s *LogAnalyticsNamespaceResourceCrud) Create() error {
	// This resource can't actually be created. So treat it as an Update instead.
	if err := s.Get(); err != nil {
		return err
	}
	return s.Update()
}

func (s *LogAnalyticsNamespaceResourceCrud) Update() error {
	var desiredState bool
	if isOnboarded, ok := s.D.GetOkExists("is_onboarded"); ok && s.D.HasChange("is_onboarded") {
		desiredState = isOnboarded.(bool)
		if desiredState == true {
			return s.OnboardNamespace()
		} else {
			return s.OffboardNamespace()
		}
	}
	return nil
}

func (s *LogAnalyticsNamespaceResourceCrud) OnboardNamespace() error {
	request := oci_log_analytics.OnboardNamespaceRequest{}
	var namespace *string
	if ns, ok := s.D.GetOkExists("namespace"); ok {
		tmp := ns.(string)
		request.NamespaceName = &tmp
		namespace = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")
	response, err := s.Client.OnboardNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNamespaceFromWorkRequest(workId, namespace, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics"), oci_log_analytics.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *LogAnalyticsNamespaceResourceCrud) OffboardNamespace() error {
	request := oci_log_analytics.OffboardNamespaceRequest{}
	var namespace *string
	if ns, ok := s.D.GetOkExists("namespace"); ok {
		tmp := ns.(string)
		request.NamespaceName = &tmp
		namespace = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")
	response, err := s.Client.OffboardNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNamespaceFromWorkRequest(workId, namespace, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics"), oci_log_analytics.ActionTypesDeleted, s.D.Timeout(schema.TimeoutCreate))
}

func (s *LogAnalyticsNamespaceResourceCrud) getNamespaceFromWorkRequest(workId *string, ns *string, retryPolicy *oci_common.RetryPolicy, actionTypeEnum oci_log_analytics.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	namespaceName, err := logAnalyticsWaitForWorkRequest(workId, ns, "loganalytics", actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*namespaceName)

	return s.Get()
}

// GET namespace returns 404 Not Found if tenancy not on-boarded.
// LIST namespace returns response irrespective of whether tenancy is on-boarded or off-boarded
// if tenancy is off-boarded during Update, the GET would throw 404 but LIST would work, hence using LIST instead of GET
func (s *LogAnalyticsNamespaceResourceCrud) Get() error {
	request := oci_log_analytics.ListNamespacesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListNamespaces(context.Background(), request)
	if err != nil {
		return err
	}

	var namespace *string
	if ns, ok := s.D.GetOkExists("namespace"); ok {
		tmp := ns.(string)
		namespace = &tmp
	}

	for _, item := range response.Items {
		if strings.ToLower(*item.NamespaceName) == strings.ToLower(*namespace) {
			s.Res = &item
		}
	}

	if s.Res == nil {
		return fmt.Errorf("[ERROR] log analytics namespace %v not found", namespace)
	}

	return nil
}

func (s *LogAnalyticsNamespaceResourceCrud) SetData() error {
	if s.Res.NamespaceName != nil {
		s.D.Set("namespace", *s.Res.NamespaceName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.IsOnboarded != nil {
		s.D.Set("is_onboarded", *s.Res.IsOnboarded)
	}

	return nil
}

func logAnalyticsWaitForWorkRequest(wId *string, ns *string, entityType string, action oci_log_analytics.ActionTypesEnum, timeout time.Duration, disableFoundRetries bool, client *oci_log_analytics.LogAnalyticsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "log_analytics")
	retryPolicy.ShouldRetryOperation = logAnalyticsWorkRequestShouldRetryFunc(timeout)

	response := oci_log_analytics.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_log_analytics.OperationStatusInProgress),
			string(oci_log_analytics.OperationStatusAccepted),
			string(oci_log_analytics.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_log_analytics.OperationStatusSucceeded),
			string(oci_log_analytics.OperationStatusFailed),
			string(oci_log_analytics.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_log_analytics.GetWorkRequestRequest{
					NamespaceName: ns,
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

	// The Log Analytics workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_log_analytics.OperationStatusFailed || response.Status == oci_log_analytics.OperationStatusCanceled {
		return nil, getErrorFromLogAnalyticsWorkRequest(client, wId, ns, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func logAnalyticsWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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
		if workRequestResponse, ok := response.Response.(oci_log_analytics.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func getErrorFromLogAnalyticsWorkRequest(client *oci_log_analytics.LogAnalyticsClient, wId *string, ns *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_log_analytics.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_log_analytics.ListWorkRequestErrorsRequest{
			WorkRequestId: wId,
			NamespaceName: ns,
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
	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)
	return workRequestErr
}
