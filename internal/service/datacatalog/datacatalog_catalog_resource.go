// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func DatacatalogCatalogResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatacatalogCatalog,
		Read:     readDatacatalogCatalog,
		Update:   updateDatacatalogCatalog,
		Delete:   deleteDatacatalogCatalog,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"attached_catalog_private_endpoints": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
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
			"number_of_objects": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"service_api_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_console_url": {
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

func createDatacatalogCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.CreateResource(d, sync)
}

func readDatacatalogCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.ReadResource(sync)
}

func updateDatacatalogCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatacatalogCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatacatalogCatalogResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datacatalog.DataCatalogClient
	Res                    *oci_datacatalog.Catalog
	DisableNotFoundRetries bool
}

func (s *DatacatalogCatalogResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatacatalogCatalogResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateCreating),
	}
}

func (s *DatacatalogCatalogResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateActive),
	}
}

func (s *DatacatalogCatalogResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateDeleting),
	}
}

func (s *DatacatalogCatalogResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateDeleted),
	}
}

func (s *DatacatalogCatalogResourceCrud) Create() error {
	request := oci_datacatalog.CreateCatalogRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.CreateCatalog(context.Background(), request)
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "catalog") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}

	err = s.getCatalogFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog"), oci_datacatalog.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	if attachedCatalogPrivateEndpoints, ok := s.D.GetOkExists("attached_catalog_private_endpoints"); ok {
		set := attachedCatalogPrivateEndpoints.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attached_catalog_private_endpoints") {
			err := s.attachCatalogPrivateEndpoints(tmp)
			if err != nil {
				// store the state when datacatalog is created but attach fail
				if setDataErr := s.SetData(); setDataErr != nil {
					return setDataErr
				}
				return err
			}
		}
	}

	return nil
}

func (s *DatacatalogCatalogResourceCrud) getCatalogFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datacatalog.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	catalogId, err := catalogWaitForWorkRequest(workId, "catalog",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*catalogId)

	return s.Get()
}

func catalogWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func catalogWaitForWorkRequest(wId *string, entityType string, action oci_datacatalog.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datacatalog.DataCatalogClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datacatalog")
	retryPolicy.ShouldRetryOperation = catalogWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDatacatalogCatalogWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatacatalogCatalogWorkRequest(client *oci_datacatalog.DataCatalogClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datacatalog.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatacatalogCatalogResourceCrud) Get() error {
	request := oci_datacatalog.GetCatalogRequest{}

	tmp := s.D.Id()
	request.CatalogId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.GetCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Catalog
	return nil
}

func (s *DatacatalogCatalogResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	if _, ok := s.D.GetOkExists("attached_catalog_private_endpoints"); ok && s.D.HasChange("attached_catalog_private_endpoints") {
		o, n := s.D.GetChange("attached_catalog_private_endpoints")
		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}

		os := o.(*schema.Set)
		ns := n.(*schema.Set)

		newPrivateEndpointsToAttach := ns.Difference(os).List()
		oldPrivateEndpointsToDetach := os.Difference(ns).List()

		if len(newPrivateEndpointsToAttach) > 0 {
			tmp := make([]string, len(newPrivateEndpointsToAttach))
			for i := range newPrivateEndpointsToAttach {
				if newPrivateEndpointsToAttach[i] != nil {
					tmp[i] = newPrivateEndpointsToAttach[i].(string)
				}
			}
			err := s.attachCatalogPrivateEndpoints(tmp)
			if err != nil {
				return err
			}
		}

		if len(oldPrivateEndpointsToDetach) > 0 {
			tmp := make([]string, len(oldPrivateEndpointsToDetach))
			for i := range oldPrivateEndpointsToDetach {
				if oldPrivateEndpointsToDetach[i] != nil {
					tmp[i] = oldPrivateEndpointsToDetach[i].(string)
				}
			}
			err := s.detachCatalogPrivateEndpoints(tmp)
			if err != nil {
				return err
			}
		}
	}

	request := oci_datacatalog.UpdateCatalogRequest{}

	tmp := s.D.Id()
	request.CatalogId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.UpdateCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Catalog
	return nil
}

func (s *DatacatalogCatalogResourceCrud) Delete() error {
	request := oci_datacatalog.DeleteCatalogRequest{}

	tmp := s.D.Id()
	request.CatalogId = &tmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	// Need to detach all Private Endpoints before delete Catalog
	if attachedCatalogPrivateEndpoints, ok := s.D.GetOkExists("attached_catalog_private_endpoints"); ok {
		set := attachedCatalogPrivateEndpoints.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 {
			err := s.detachCatalogPrivateEndpoints(tmp)
			if err != nil {
				return err
			}
		}
	}

	response, err := s.Client.DeleteCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := catalogWaitForWorkRequest(workId, "catalog",
		oci_datacatalog.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatacatalogCatalogResourceCrud) SetData() error {
	attachedCatalogPrivateEndpoints := []interface{}{}
	for _, item := range s.Res.AttachedCatalogPrivateEndpoints {
		attachedCatalogPrivateEndpoints = append(attachedCatalogPrivateEndpoints, item)
	}
	s.D.Set("attached_catalog_private_endpoints", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, attachedCatalogPrivateEndpoints))

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, ResourceLockToMapCatalog(item))
	}
	s.D.Set("locks", locks)

	if s.Res.NumberOfObjects != nil {
		s.D.Set("number_of_objects", *s.Res.NumberOfObjects)
	}

	if s.Res.ServiceApiUrl != nil {
		s.D.Set("service_api_url", *s.Res.ServiceApiUrl)
	}

	if s.Res.ServiceConsoleUrl != nil {
		s.D.Set("service_console_url", *s.Res.ServiceConsoleUrl)
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

func ResourceLockToMapCatalog(obj oci_datacatalog.ResourceLock) map[string]interface{} {
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

func (s *DatacatalogCatalogResourceCrud) attachCatalogPrivateEndpoints(attachCatalogPrivateEndpoints []string) error {
	for _, attachCatalogPrivateEndpoint := range attachCatalogPrivateEndpoints {
		catalogId := s.D.Id()
		request := oci_datacatalog.AttachCatalogPrivateEndpointRequest{}
		request.CatalogPrivateEndpointId = &attachCatalogPrivateEndpoint
		request.CatalogId = &catalogId
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

		response, err := s.Client.AttachCatalogPrivateEndpoint(context.Background(), request)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId

		// Wait until Update finishes
		_, err = catalogWaitForWorkRequest(workId, "catalog",
			oci_datacatalog.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	}
	return nil
}

func (s *DatacatalogCatalogResourceCrud) detachCatalogPrivateEndpoints(detachCatalogPrivateEndpoints []string) error {
	for _, detachCatalogPrivateEndpoint := range detachCatalogPrivateEndpoints {
		catalogId := s.D.Id()
		request := oci_datacatalog.DetachCatalogPrivateEndpointRequest{}
		request.CatalogPrivateEndpointId = &detachCatalogPrivateEndpoint
		request.CatalogId = &catalogId
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

		response, err := s.Client.DetachCatalogPrivateEndpoint(context.Background(), request)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId

		// Wait until it finishes
		_, err = catalogWaitForWorkRequest(workId, "catalog",
			oci_datacatalog.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	}
	return nil
}

func (s *DatacatalogCatalogResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datacatalog.ChangeCatalogCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CatalogId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.ChangeCatalogCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getCatalogFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "catalog"), oci_datacatalog.WorkRequestResourceActionTypeMoved, s.D.Timeout(schema.TimeoutUpdate))
}
