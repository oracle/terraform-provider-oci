// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func OcvpDatastoreClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOcvpDatastoreCluster,
		Read:     readOcvpDatastoreCluster,
		Update:   updateOcvpDatastoreCluster,
		Delete:   deleteOcvpDatastoreCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datastore_cluster_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"datastore_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

			// Computed
			"capacity_in_gbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"esxi_host_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createOcvpDatastoreCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpDatastoreClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatastoreClusterClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOcvpDatastoreCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpDatastoreClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatastoreClusterClient()

	return tfresource.ReadResource(sync)
}

func updateOcvpDatastoreCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpDatastoreClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatastoreClusterClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOcvpDatastoreCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpDatastoreClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatastoreClusterClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type OcvpDatastoreClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ocvp.DatastoreClusterClient
	Res                    *oci_ocvp.DatastoreCluster
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_ocvp.WorkRequestClient
}

func (s *OcvpDatastoreClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OcvpDatastoreClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesCreating),
	}
}

func (s *OcvpDatastoreClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesActive),
	}
}

func (s *OcvpDatastoreClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleting),
	}
}

func (s *OcvpDatastoreClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleted),
	}
}

func (s *OcvpDatastoreClusterResourceCrud) Create() error {
	request := oci_ocvp.CreateDatastoreClusterRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if datastoreClusterType, ok := s.D.GetOkExists("datastore_cluster_type"); ok {
		request.DatastoreClusterType = oci_ocvp.DatastoreClusterTypesEnum(datastoreClusterType.(string))
	}

	if datastoreIds, ok := s.D.GetOkExists("datastore_ids"); ok {
		interfaces := datastoreIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("datastore_ids") {
			request.DatastoreIds = tmp
		}
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

	response, err := s.Client.CreateDatastoreCluster(context.Background(), request)
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "sddc-datastore-cluster") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getDatastoreClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OcvpDatastoreClusterResourceCrud) getDatastoreClusterFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ocvp.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	datastoreClusterId, err := datastoreClusterWaitForWorkRequest(workId, "sddc-datastore-cluster",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*datastoreClusterId)

	return s.Get()
}

func datastoreClusterWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func datastoreClusterWaitForWorkRequest(wId *string, entityType string, action oci_ocvp.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ocvp.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ocvp")
	retryPolicy.ShouldRetryOperation = datastoreClusterWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromOcvpDatastoreClusterWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOcvpDatastoreClusterWorkRequest(client *oci_ocvp.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ocvp.ActionTypesEnum) error {
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

func (s *OcvpDatastoreClusterResourceCrud) Get() error {
	request := oci_ocvp.GetDatastoreClusterRequest{}

	tmp := s.D.Id()
	request.DatastoreClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.GetDatastoreCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DatastoreCluster
	return nil
}

func (s *OcvpDatastoreClusterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ocvp.UpdateDatastoreClusterRequest{}

	tmp := s.D.Id()
	request.DatastoreClusterId = &tmp

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

	response, err := s.Client.UpdateDatastoreCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getDatastoreClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	var datastoreIdsToRemove []string
	var datastoreIdsToAdd []string
	if s.D.HasChange("datastore_ids") {
		oldDatastoreIds, newDatastoreIds := s.D.GetChange("datastore_ids")
		log.Printf("[DEBUG] datastore_ids have been updated in config from %v to %v", oldDatastoreIds, newDatastoreIds)
		datastoreIdsToRemove, datastoreIdsToAdd = compareSlices(castToSliceOfStrings(oldDatastoreIds), castToSliceOfStrings(newDatastoreIds))
		log.Printf("[DEBUG] datastore_ids to remove: %v; datastore_ids to add: %v", datastoreIdsToRemove, datastoreIdsToAdd)
	}

	if len(datastoreIdsToRemove) > 0 {
		log.Printf("[DEBUG] removing datastores %v", datastoreIdsToRemove)
		if err = s.removeDatastore(datastoreIdsToRemove); err != nil {
			return err
		}
	}

	if len(datastoreIdsToAdd) > 0 {
		log.Printf("[DEBUG] adding datastores %v", datastoreIdsToAdd)
		if err = s.addDatastore(datastoreIdsToAdd); err != nil {
			return err
		}
	}
	return s.Get()
}

func castToSliceOfStrings(untypedInput interface{}) []string {
	if untypedInput == nil {
		return []string{}
	}
	sliceOfUntyped := untypedInput.([]interface{})
	sliceOfStrings := make([]string, 0, len(sliceOfUntyped))
	for _, untypedElement := range sliceOfUntyped {
		sliceOfStrings = append(sliceOfStrings, untypedElement.(string))
	}
	return sliceOfStrings
}

func (s *OcvpDatastoreClusterResourceCrud) Delete() error {
	// Cannot delete datastore cluster with datastores. So remove them first.
	if datastoreIds, hasDatastores := s.D.GetOk("datastore_ids"); hasDatastores {
		if err := s.removeDatastore(castToSliceOfStrings(datastoreIds)); err != nil {
			return err
		}

	}

	tmp := s.D.Id()
	request := oci_ocvp.DeleteDatastoreClusterRequest{}

	request.DatastoreClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.DeleteDatastoreCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := datastoreClusterWaitForWorkRequest(workId, "sddc-datastore-cluster",
		oci_ocvp.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *OcvpDatastoreClusterResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CapacityInGBs != nil {
		s.D.Set("capacity_in_gbs", *s.Res.CapacityInGBs)
	}

	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("datastore_cluster_type", s.Res.DatastoreClusterType)

	s.D.Set("datastore_ids", s.Res.DatastoreIds)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("esxi_host_ids", s.Res.EsxiHostIds)

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

func DatastoreClusterSummaryToMap(obj oci_ocvp.DatastoreClusterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CapacityInGBs != nil {
		result["capacity_in_gbs"] = float64(*obj.CapacityInGBs)
	}

	if obj.ClusterId != nil {
		result["cluster_id"] = string(*obj.ClusterId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["datastore_cluster_type"] = string(obj.DatastoreClusterType)

	result["datastore_ids"] = obj.DatastoreIds

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["esxi_host_ids"] = obj.EsxiHostIds

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

func (s *OcvpDatastoreClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ocvp.ChangeDatastoreClusterCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatastoreClusterId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	_, err := s.Client.ChangeDatastoreClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *OcvpDatastoreClusterResourceCrud) addDatastore(datastoreId []string) error {
	datastoreClusterId := s.D.Id()
	request := oci_ocvp.AddDatastoreToDatastoreClusterRequest{
		DatastoreClusterId: &datastoreClusterId,
		AddDatastoreToDatastoreClusterDetails: oci_ocvp.AddDatastoreToDatastoreClusterDetails{
			DatastoreIds: datastoreId,
		},
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")
	response, err := s.Client.AddDatastoreToDatastoreCluster(context.Background(), request)
	if err != nil {
		return err
	}
	_, err = datastoreClusterWaitForWorkRequest(response.OpcWorkRequestId, "sddc-datastore-cluster", oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate),
		s.DisableNotFoundRetries, s.WorkRequestClient)
	return err
}

func (s *OcvpDatastoreClusterResourceCrud) removeDatastore(datastoreId []string) error {
	datastoreClusterId := s.D.Id()
	request := oci_ocvp.RemoveDatastoreFromDatastoreClusterRequest{
		DatastoreClusterId: &datastoreClusterId,
		RemoveDatastoreFromDatastoreClusterDetails: oci_ocvp.RemoveDatastoreFromDatastoreClusterDetails{
			DatastoreIds: datastoreId,
		},
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")
	response, err := s.Client.RemoveDatastoreFromDatastoreCluster(context.Background(), request)
	if err != nil {
		return err
	}
	_, err = datastoreClusterWaitForWorkRequest(response.OpcWorkRequestId, "sddc-datastore-cluster", oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate),
		s.DisableNotFoundRetries, s.WorkRequestClient)
	return err
}
