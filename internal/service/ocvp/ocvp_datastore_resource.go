// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpDatastoreResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOcvpDatastore,
		Read:     readOcvpDatastore,
		Update:   updateOcvpDatastore,
		Delete:   deleteOcvpDatastore,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"block_volume_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
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

			// Computed
			"block_volume_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"attachments": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"esxi_host_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"iqn": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"capacity_in_gbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sddc_id": {
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

func createOcvpDatastore(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpDatastoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatastoreClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOcvpDatastore(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpDatastoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatastoreClient()

	return tfresource.ReadResource(sync)
}

func updateOcvpDatastore(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpDatastoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatastoreClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOcvpDatastore(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpDatastoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatastoreClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type OcvpDatastoreResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ocvp.DatastoreClient
	Res                    *oci_ocvp.Datastore
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_ocvp.WorkRequestClient
}

func (s *OcvpDatastoreResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OcvpDatastoreResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesCreating),
	}
}

func (s *OcvpDatastoreResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesActive),
	}
}

func (s *OcvpDatastoreResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleting),
	}
}

func (s *OcvpDatastoreResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleted),
	}
}

func (s *OcvpDatastoreResourceCrud) Create() error {
	request := oci_ocvp.CreateDatastoreRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if blockVolumeIds, ok := s.D.GetOkExists("block_volume_ids"); ok {
		interfaces := blockVolumeIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("block_volume_ids") {
			request.BlockVolumeIds = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.CreateDatastore(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_ocvp.GetWorkRequestResponse{}
	workRequestResponse, err = s.WorkRequestClient.GetWorkRequest(context.Background(),
		oci_ocvp.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "sddc-datastore") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getDatastoreFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OcvpDatastoreResourceCrud) getDatastoreFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ocvp.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	datastoreId, err := datastoreWaitForWorkRequest(workId, "sddc-datastore",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*datastoreId)

	return s.Get()
}

func datastoreWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ocvp", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ocvp.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func datastoreWaitForWorkRequest(wId *string, entityType string, action oci_ocvp.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ocvp.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ocvp")
	retryPolicy.ShouldRetryOperation = datastoreWorkRequestShouldRetryFunc(timeout)

	response := oci_ocvp.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_ocvp.OperationStatusInProgress),
			string(oci_ocvp.OperationStatusAccepted),
			string(oci_ocvp.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ocvp.OperationStatusSucceeded),
			string(oci_ocvp.OperationStatusFailed),
			string(oci_ocvp.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ocvp.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ocvp.OperationStatusFailed || response.Status == oci_ocvp.OperationStatusCanceled {
		return nil, getErrorFromOcvpDatastoreWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOcvpDatastoreWorkRequest(client *oci_ocvp.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ocvp.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ocvp.ListWorkRequestErrorsRequest{
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

func (s *OcvpDatastoreResourceCrud) Get() error {
	request := oci_ocvp.GetDatastoreRequest{}

	tmp := s.D.Id()
	request.DatastoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.GetDatastore(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Datastore
	return nil
}

func (s *OcvpDatastoreResourceCrud) Update() error {
	var blockVolumeIdsToRemove []string
	var blockVolumeIdsToAdd []string
	if s.D.HasChange("block_volume_ids") {
		oldBvIds, newBvIds := s.D.GetChange("block_volume_ids")
		log.Printf("[DEBUG] block_volume_ids has been updated in config from %v to %v", oldBvIds, newBvIds)
		blockVolumeIdsToRemove, blockVolumeIdsToAdd = compareSlices(castToSliceOfStrings(oldBvIds), castToSliceOfStrings(newBvIds))
		log.Printf("[DEBUG] block volumes to remove: %v; block volumes to add: %v", blockVolumeIdsToRemove, blockVolumeIdsToAdd)
	}

	if len(blockVolumeIdsToRemove) > 0 {
		return fmt.Errorf("Block volumes cannot be removed from datastore. Please consider tainting/replacing the datastore if needed")
	}

	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ocvp.UpdateDatastoreRequest{}

	datastoreId := s.D.Id()
	request.DatastoreId = &datastoreId

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.UpdateDatastore(context.Background(), request)
	if err != nil {
		return err
	}

	_, err = datastoreWaitForWorkRequest(response.OpcWorkRequestId, "sddc-datastore", oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
	if err != nil {
		return err
	}

	for _, blockVolumeId := range blockVolumeIdsToAdd {
		log.Printf("[DEBUG] adding block volume %v", blockVolumeId)
		err = s.addBlockVolume(blockVolumeId)
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func compareSlices(old []string, new []string) (toRemove []string, toAdd []string) {
	oldSet := make(map[string]struct{})
	newSet := make(map[string]struct{})
	for _, item := range old {
		oldSet[item] = struct{}{}
	}
	for _, item := range new {
		newSet[item] = struct{}{}
	}
	for _, item := range old {
		if _, ok := newSet[item]; !ok {
			toRemove = append(toRemove, item)
		}
	}
	for _, item := range new {
		if _, ok := oldSet[item]; !ok {
			toAdd = append(toAdd, item)
		}
	}
	return
}

func (s *OcvpDatastoreResourceCrud) Delete() error {
	request := oci_ocvp.DeleteDatastoreRequest{}

	tmp := s.D.Id()
	request.DatastoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.DeleteDatastore(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := datastoreWaitForWorkRequest(workId, "sddc-datastore",
		oci_ocvp.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *OcvpDatastoreResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	blockVolumeDetails := []interface{}{}
	for _, item := range s.Res.BlockVolumeDetails {
		blockVolumeDetails = append(blockVolumeDetails, BlockVolumeDetailsToMap(item))
	}
	s.D.Set("block_volume_details", blockVolumeDetails)

	s.D.Set("block_volume_ids", s.Res.BlockVolumeIds)

	if s.Res.CapacityInGBs != nil {
		s.D.Set("capacity_in_gbs", *s.Res.CapacityInGBs)
	}

	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.SddcId != nil {
		s.D.Set("sddc_id", *s.Res.SddcId)
	}

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

func BlockVolumeAttachmentToMap(obj oci_ocvp.BlockVolumeAttachment) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EsxiHostId != nil {
		result["esxi_host_id"] = string(*obj.EsxiHostId)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	return result
}

func BlockVolumeDetailsToMap(obj oci_ocvp.BlockVolumeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	attachments := []interface{}{}
	for _, item := range obj.Attachments {
		attachments = append(attachments, BlockVolumeAttachmentToMap(item))
	}
	result["attachments"] = attachments

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Iqn != nil {
		result["iqn"] = string(*obj.Iqn)
	}

	return result
}

func DatastoreSummaryToMap(obj oci_ocvp.DatastoreSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	blockVolumeDetails := []interface{}{}
	for _, item := range obj.BlockVolumeDetails {
		blockVolumeDetails = append(blockVolumeDetails, BlockVolumeDetailsToMap(item))
	}
	result["block_volume_details"] = blockVolumeDetails

	result["block_volume_ids"] = obj.BlockVolumeIds

	if obj.CapacityInGBs != nil {
		result["capacity_in_gbs"] = float64(*obj.CapacityInGBs)
	}

	if obj.ClusterId != nil {
		result["cluster_id"] = string(*obj.ClusterId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.SddcId != nil {
		result["sddc_id"] = string(*obj.SddcId)
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

func (s *OcvpDatastoreResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ocvp.ChangeDatastoreCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatastoreId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	_, err := s.Client.ChangeDatastoreCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *OcvpDatastoreResourceCrud) addBlockVolume(blockVolumeId string) error {
	datastoreId := s.D.Id()
	request := oci_ocvp.AddBlockVolumeToDatastoreRequest{
		DatastoreId: &datastoreId,
		AddBlockVolumeToDatastoreDetails: oci_ocvp.AddBlockVolumeToDatastoreDetails{
			BlockVolumeId: &blockVolumeId,
		},
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")
	response, err := s.Client.AddBlockVolumeToDatastore(context.Background(), request)
	if err != nil {
		return err
	}
	_, err = datastoreWaitForWorkRequest(response.OpcWorkRequestId, "sddc-datastore", oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient)
	return err
}
