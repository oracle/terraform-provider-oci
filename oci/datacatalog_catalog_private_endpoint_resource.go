// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v33/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v33/datacatalog"
)

func init() {
	RegisterResource("oci_datacatalog_catalog_private_endpoint", DatacatalogCatalogPrivateEndpointResource())
}

func DatacatalogCatalogPrivateEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatacatalogCatalogPrivateEndpoint,
		Read:     readDatacatalogCatalogPrivateEndpoint,
		Update:   updateDatacatalogCatalogPrivateEndpoint,
		Delete:   deleteDatacatalogCatalogPrivateEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dns_zones": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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
			"attached_catalogs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
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

func createDatacatalogCatalogPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataCatalogClient()

	return CreateResource(d, sync)
}

func readDatacatalogCatalogPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataCatalogClient()

	return ReadResource(sync)
}

func updateDatacatalogCatalogPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataCatalogClient()

	return UpdateResource(d, sync)
}

func deleteDatacatalogCatalogPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataCatalogClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatacatalogCatalogPrivateEndpointResourceCrud struct {
	BaseCrud
	Client                 *oci_datacatalog.DataCatalogClient
	Res                    *oci_datacatalog.CatalogPrivateEndpoint
	DisableNotFoundRetries bool
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateCreating),
	}
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateActive),
	}
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateDeleting),
	}
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datacatalog.LifecycleStateDeleted),
	}
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) Create() error {
	request := oci_datacatalog.CreateCatalogPrivateEndpointRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if dnsZones, ok := s.D.GetOkExists("dns_zones"); ok {
		interfaces := dnsZones.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns_zones") {
			request.DnsZones = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.CreateCatalogPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getCatalogPrivateEndpointFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "datacatalog"), oci_datacatalog.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) getCatalogPrivateEndpointFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datacatalog.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	catalogPrivateEndpointId, err := catalogPrivateEndpointWaitForWorkRequest(workId, "catalogPrivateEndpoint",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*catalogPrivateEndpointId)

	return s.Get()
}

func catalogPrivateEndpointWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "datacatalog", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_datacatalog.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func catalogPrivateEndpointWaitForWorkRequest(wId *string, entityType string, action oci_datacatalog.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datacatalog.DataCatalogClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "datacatalog")
	retryPolicy.ShouldRetryOperation = catalogPrivateEndpointWorkRequestShouldRetryFunc(timeout)

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
		if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	if identifier == nil {
		return nil, getErrorFromCatalogWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) Get() error {
	request := oci_datacatalog.GetCatalogPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.CatalogPrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.GetCatalogPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CatalogPrivateEndpoint
	return nil
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datacatalog.UpdateCatalogPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.CatalogPrivateEndpointId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if dnsZones, ok := s.D.GetOkExists("dns_zones"); ok {
		interfaces := dnsZones.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns_zones") {
			request.DnsZones = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.UpdateCatalogPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getCatalogPrivateEndpointFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "datacatalog"), oci_datacatalog.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) Delete() error {
	request := oci_datacatalog.DeleteCatalogPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.CatalogPrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.DeleteCatalogPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := catalogPrivateEndpointWaitForWorkRequest(workId, "catalogPrivateEndpoint",
		oci_datacatalog.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) SetData() error {
	s.D.Set("attached_catalogs", s.Res.AttachedCatalogs)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("dns_zones", s.Res.DnsZones)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datacatalog.ChangeCatalogPrivateEndpointCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CatalogPrivateEndpointId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

	response, err := s.Client.ChangeCatalogPrivateEndpointCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId

	// Wait until it finishes
	_, err = catalogPrivateEndpointWaitForWorkRequest(workId, "catalogPrivateEndpoint",
		oci_datacatalog.WorkRequestResourceActionTypeMoved, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatacatalogCatalogPrivateEndpointResourceCrud) detachCatalog(detachCatalogs []interface{}) error {
	for _, detachCatalog := range detachCatalogs {
		detachCatalogId := detachCatalog.(string)
		catalogPrivateEndpointId := s.D.Id()
		request := oci_datacatalog.DetachCatalogPrivateEndpointRequest{}
		request.CatalogPrivateEndpointId = &catalogPrivateEndpointId
		request.CatalogId = &detachCatalogId
		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datacatalog")

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
