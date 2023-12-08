// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacatalog

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v65/datacatalog"
)

func DatacatalogMetastoreResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatacatalogMetastore,
		Read:     readDatacatalogMetastore,
		Update:   updateDatacatalogMetastore,
		Delete:   deleteDatacatalogMetastore,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_external_table_location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_managed_table_location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
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

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"locks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createDatacatalogMetastore(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogMetastoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.CreateResource(d, sync)
}

func readDatacatalogMetastore(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogMetastoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.ReadResource(sync)
}

func updateDatacatalogMetastore(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogMetastoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatacatalogMetastore(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogMetastoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatacatalogMetastoreResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datacatalog.DataCatalogClient
	Res                    *oci_datacatalog.Metastore
	DisableNotFoundRetries bool
}

func (s *DatacatalogMetastoreResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatacatalogMetastoreResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateCreating),
	}
}

func (s *DatacatalogMetastoreResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateActive),
	}
}

func (s *DatacatalogMetastoreResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateDeleting),
	}
}

func (s *DatacatalogMetastoreResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateDeleted),
	}
}

func (s *DatacatalogMetastoreResourceCrud) Create() error {
	request := oci_datacatalog.CreateMetastoreRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if defaultExternalTableLocation, ok := s.D.GetOkExists("default_external_table_location"); ok {
		tmp := defaultExternalTableLocation.(string)
		request.DefaultExternalTableLocation = &tmp
	}

	if defaultManagedTableLocation, ok := s.D.GetOkExists("default_managed_table_location"); ok {
		tmp := defaultManagedTableLocation.(string)
		request.DefaultManagedTableLocation = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.CreateMetastore(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_datacatalog.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_datacatalog.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "metastore") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	err = s.getMetastoreFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog"), oci_datacatalog.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	return nil
}

func (s *DatacatalogMetastoreResourceCrud) getMetastoreFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datacatalog.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	catalogId, err := catalogMetastoreWaitForWorkRequest(workId, "metastore",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*catalogId)

	return s.Get()
}

func catalogMetastoreWaitForWorkRequest(wId *string, entityType string, action oci_datacatalog.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datacatalog.DataCatalogClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datacatalog")
	retryPolicy.ShouldRetryOperation = catalogMetastoreWorkRequestShouldRetryFunc(timeout)

	response := oci_datacatalog.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_datacatalog.WorkRequestStatusInProgress),
			string(oci_datacatalog.WorkRequestStatusAccepted),
			string(oci_datacatalog.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_datacatalog.WorkRequestStatusSucceeded),
			string(oci_datacatalog.WorkRequestStatusFailed),
			string(oci_datacatalog.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_datacatalog.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_datacatalog.WorkRequestStatusFailed || response.Status == oci_datacatalog.WorkRequestStatusCanceled {
		return nil, getErrorFromDatacatalogCatalogMetastoreWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func catalogMetastoreWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "datacatalog", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_datacatalog.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func getErrorFromDatacatalogCatalogMetastoreWorkRequest(client *oci_datacatalog.DataCatalogClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datacatalog.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_datacatalog.ListWorkRequestErrorsRequest{
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

func (s *DatacatalogMetastoreResourceCrud) Get() error {
	request := oci_datacatalog.GetMetastoreRequest{}

	tmp := s.D.Id()
	request.MetastoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.GetMetastore(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Metastore
	return nil
}

func (s *DatacatalogMetastoreResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datacatalog.UpdateMetastoreRequest{}

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

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	tmp := s.D.Id()
	request.MetastoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.UpdateMetastore(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Metastore
	return nil
}

func (s *DatacatalogMetastoreResourceCrud) Delete() error {
	request := oci_datacatalog.DeleteMetastoreRequest{}

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	tmp := s.D.Id()
	request.MetastoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	_, err := s.Client.DeleteMetastore(context.Background(), request)
	return err
}

func (s *DatacatalogMetastoreResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefaultExternalTableLocation != nil {
		s.D.Set("default_external_table_location", *s.Res.DefaultExternalTableLocation)
	}

	if s.Res.DefaultManagedTableLocation != nil {
		s.D.Set("default_managed_table_location", *s.Res.DefaultManagedTableLocation)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, ResourceLockToMapMetastore(item))
	}
	s.D.Set("locks", locks)

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

func ResourceLockToMapMetastore(obj oci_datacatalog.ResourceLock) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	if obj.RelatedResourceId != nil {
		result["related_resource_id"] = string(*obj.RelatedResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *DatacatalogMetastoreResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datacatalog.ChangeMetastoreCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	idTmp := s.D.Id()
	changeCompartmentRequest.MetastoreId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	_, err := s.Client.ChangeMetastoreCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
