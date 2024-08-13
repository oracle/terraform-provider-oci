// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func DataSafeCalculateAuditVolumeAvailableResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeCalculateAuditVolumeAvailable,
		Read:     readDataSafeCalculateAuditVolumeAvailable,
		Delete:   deleteDataSafeCalculateAuditVolumeAvailable,
		Schema: map[string]*schema.Schema{
			// Required
			"audit_profile_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"audit_collection_start_time": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"database_unique_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trail_locations": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"available_audit_volumes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"audit_profile_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"month_in_consideration": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trail_location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createDataSafeCalculateAuditVolumeAvailable(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeCalculateAuditVolumeAvailableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeCalculateAuditVolumeAvailable(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDataSafeCalculateAuditVolumeAvailable(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeCalculateAuditVolumeAvailableResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.CalculateAuditVolumeAvailableResponse
	DisableNotFoundRetries bool
}

type DataSafeAuditProfileAvailableAuditVolumesResourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAvailableAuditVolumesResponse
}

func (s *DataSafeCalculateAuditVolumeAvailableResourceCrud) ID() string {
	return *s.Res.OpcRequestId
}

func (s *DataSafeCalculateAuditVolumeAvailableResourceCrud) Create() error {
	request := oci_data_safe.CalculateAuditVolumeAvailableRequest{}

	if auditCollectionStartTime, ok := s.D.GetOkExists("audit_collection_start_time"); ok {
		tmp, err := time.Parse(time.RFC3339, auditCollectionStartTime.(string))
		if err != nil {
			return err
		}
		request.AuditCollectionStartTime = &oci_common.SDKTime{Time: tmp}
	}

	if auditProfileId, ok := s.D.GetOkExists("audit_profile_id"); ok {
		tmp := auditProfileId.(string)
		request.AuditProfileId = &tmp
	}

	if databaseUniqueName, ok := s.D.GetOkExists("database_unique_name"); ok {
		tmp := databaseUniqueName.(string)
		request.DatabaseUniqueName = &tmp
	}

	if trailLocations, ok := s.D.GetOkExists("trail_locations"); ok {
		interfaces := trailLocations.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 {
			request.TrailLocations = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")
	response, err := s.Client.CalculateAuditVolumeAvailable(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response
	return s.getCalculateAuditVolumeAvailableFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeCalculateAuditVolumeAvailableResourceCrud) getCalculateAuditVolumeAvailableFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := calculateAuditVolumeAvailableWaitForWorkRequest(workId, "auditprofile",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] computation failed, attempting to cancel the workrequest: %v \n", workId)
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
	return s.GetAvailableAuditVolumes(workId)
}

func (s *DataSafeCalculateAuditVolumeAvailableResourceCrud) GetAvailableAuditVolumes(workId *string) error {
	request := oci_data_safe.ListAvailableAuditVolumesRequest{}

	if auditProfileId, ok := s.D.GetOkExists("audit_profile_id"); ok {
		tmp := auditProfileId.(string)
		request.AuditProfileId = &tmp
	}

	request.WorkRequestId = workId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")
	listResponse, err := s.Client.ListAvailableAuditVolumes(context.Background(), request)
	if err != nil {
		return err
	}

	items := []interface{}{}
	for _, item := range listResponse.Items {
		items = append(items, AvailableAuditVolumeSummaryToMap(item))
	}

	s.D.Set("available_audit_volumes", items)
	return nil
}

func calculateAuditVolumeAvailableWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func calculateAuditVolumeAvailableWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = calculateAuditVolumeAvailableWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeCalculateAuditVolumeAvailableWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeCalculateAuditVolumeAvailableWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeCalculateAuditVolumeAvailableResourceCrud) SetData() error {
	return nil
}
