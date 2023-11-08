// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaDownloadsJavaDownloadReportResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createJmsJavaDownloadsJavaDownloadReport,
		Read:     readJmsJavaDownloadsJavaDownloadReport,
		Delete:   deleteJmsJavaDownloadsJavaDownloadReport,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"format": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"time_end": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"time_start": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},

			// Computed
			"checksum_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"checksum_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_size_in_bytes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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
		},
	}
}

func createJmsJavaDownloadsJavaDownloadReport(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.CreateResource(d, sync)
}

func readJmsJavaDownloadsJavaDownloadReport(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

func deleteJmsJavaDownloadsJavaDownloadReport(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type JmsJavaDownloadsJavaDownloadReportResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_jms_java_downloads.JavaDownloadClient
	Res                    *oci_jms_java_downloads.JavaDownloadReport
	DisableNotFoundRetries bool
}

func (s *JmsJavaDownloadsJavaDownloadReportResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *JmsJavaDownloadsJavaDownloadReportResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateCreating),
	}
}

func (s *JmsJavaDownloadsJavaDownloadReportResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateActive),
		string(oci_jms_java_downloads.LifecycleStateNeedsAttention),
	}
}

func (s *JmsJavaDownloadsJavaDownloadReportResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateDeleting),
	}
}

func (s *JmsJavaDownloadsJavaDownloadReportResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateDeleted),
	}
}

func (s *JmsJavaDownloadsJavaDownloadReportResourceCrud) Create() error {
	request := oci_jms_java_downloads.CreateJavaDownloadReportRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if format, ok := s.D.GetOkExists("format"); ok {
		request.Format = oci_jms_java_downloads.JavaDownloadReportFormatEnum(format.(string))
	}

	if timeEnd, ok := s.D.GetOkExists("time_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
		if err != nil {
			return err
		}
		request.TimeEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeStart, ok := s.D.GetOkExists("time_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return err
		}
		request.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads")

	response, err := s.Client.CreateJavaDownloadReport(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_jms_java_downloads.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_jms_java_downloads.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "jmsjavadownloadreport") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getJavaDownloadReportFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads"), oci_jms_java_downloads.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *JmsJavaDownloadsJavaDownloadReportResourceCrud) getJavaDownloadReportFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_jms_java_downloads.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	javaDownloadReportId, err := javaDownloadReportWaitForWorkRequest(workId, "jmsjavadownloadreport",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*javaDownloadReportId)

	return s.Get()
}

func javaDownloadReportWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "jms_java_downloads", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_jms_java_downloads.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func javaDownloadReportWaitForWorkRequest(wId *string, entityType string, action oci_jms_java_downloads.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_jms_java_downloads.JavaDownloadClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "jms_java_downloads")
	retryPolicy.ShouldRetryOperation = javaDownloadReportWorkRequestShouldRetryFunc(timeout)

	response := oci_jms_java_downloads.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_jms_java_downloads.ListWorkRequestsStatusInProgress),
			string(oci_jms_java_downloads.ListWorkRequestsStatusAccepted),
			string(oci_jms_java_downloads.ListWorkRequestsStatusCanceling),
		},
		Target: []string{
			string(oci_jms_java_downloads.ListWorkRequestsStatusSucceeded),
			string(oci_jms_java_downloads.ListWorkRequestsStatusFailed),
			string(oci_jms_java_downloads.ListWorkRequestsStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_jms_java_downloads.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_jms_java_downloads.OperationStatusEnum(oci_jms_java_downloads.ListWorkRequestsStatusFailed) || response.Status == oci_jms_java_downloads.OperationStatusEnum(oci_jms_java_downloads.ListWorkRequestsStatusCanceled) {
		return nil, getErrorFromJmsJavaDownloadsJavaDownloadReportWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromJmsJavaDownloadsJavaDownloadReportWorkRequest(client *oci_jms_java_downloads.JavaDownloadClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_jms_java_downloads.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_jms_java_downloads.ListWorkRequestErrorsRequest{
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

func (s *JmsJavaDownloadsJavaDownloadReportResourceCrud) Get() error {
	request := oci_jms_java_downloads.GetJavaDownloadReportRequest{}

	tmp := s.D.Id()
	request.JavaDownloadReportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads")

	response, err := s.Client.GetJavaDownloadReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.JavaDownloadReport
	return nil
}

func (s *JmsJavaDownloadsJavaDownloadReportResourceCrud) Delete() error {
	request := oci_jms_java_downloads.DeleteJavaDownloadReportRequest{}

	tmp := s.D.Id()
	request.JavaDownloadReportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads")

	response, err := s.Client.DeleteJavaDownloadReport(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := javaDownloadReportWaitForWorkRequest(workId, "jmsjavadownloadreport",
		oci_jms_java_downloads.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *JmsJavaDownloadsJavaDownloadReportResourceCrud) SetData() error {
	s.D.Set("checksum_type", s.Res.ChecksumType)

	if s.Res.ChecksumValue != nil {
		s.D.Set("checksum_value", *s.Res.ChecksumValue)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", []interface{}{PrincipalToMap(s.Res.CreatedBy)})
	} else {
		s.D.Set("created_by", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FileSizeInBytes != nil {
		s.D.Set("file_size_in_bytes", strconv.FormatInt(*s.Res.FileSizeInBytes, 10))
	}

	s.D.Set("format", s.Res.Format)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func JavaDownloadReportSummaryToMap(obj oci_jms_java_downloads.JavaDownloadReportSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["checksum_type"] = string(obj.ChecksumType)

	if obj.ChecksumValue != nil {
		result["checksum_value"] = string(*obj.ChecksumValue)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = []interface{}{PrincipalToMap(obj.CreatedBy)}
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FileSizeInBytes != nil {
		result["file_size_in_bytes"] = strconv.FormatInt(*obj.FileSizeInBytes, 10)
	}

	result["format"] = string(obj.Format)

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
