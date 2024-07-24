// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globally_distributed_database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_globally_distributed_database "github.com/oracle/oci-go-sdk/v65/globallydistributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GloballyDistributedDatabasePrivateEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createGloballyDistributedDatabasePrivateEndpoint,
		Read:   readGloballyDistributedDatabasePrivateEndpoint,
		Update: updateGloballyDistributedDatabasePrivateEndpoint,
		Delete: deleteGloballyDistributedDatabasePrivateEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
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
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
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
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sharded_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"vcn_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createGloballyDistributedDatabasePrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabasePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()

	return tfresource.CreateResource(d, sync)
}

func readGloballyDistributedDatabasePrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabasePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()

	return tfresource.ReadResource(sync)
}

func updateGloballyDistributedDatabasePrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabasePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteGloballyDistributedDatabasePrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabasePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GloballyDistributedDatabasePrivateEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_globally_distributed_database.ShardedDatabaseServiceClient
	Res                    *oci_globally_distributed_database.PrivateEndpoint
	DisableNotFoundRetries bool
}

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_globally_distributed_database.PrivateEndpointLifecycleStateCreating),
	}
}

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_globally_distributed_database.PrivateEndpointLifecycleStateActive),
	}
}

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_globally_distributed_database.PrivateEndpointLifecycleStateDeleting),
	}
}

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_globally_distributed_database.PrivateEndpointLifecycleStateDeleted),
	}
}

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) Create() error {
	request := oci_globally_distributed_database.CreatePrivateEndpointRequest{}

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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.CreatePrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getPrivateEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database"), oci_globally_distributed_database.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) getPrivateEndpointFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_globally_distributed_database.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	privateEndpointId, err := privateEndpointWaitForWorkRequest(workId, "privateendpoint",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*privateEndpointId)

	return s.Get()
}

func privateEndpointWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "globally_distributed_database", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_globally_distributed_database.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func privateEndpointWaitForWorkRequest(wId *string, entityType string, action oci_globally_distributed_database.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_globally_distributed_database.ShardedDatabaseServiceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "globally_distributed_database")
	retryPolicy.ShouldRetryOperation = privateEndpointWorkRequestShouldRetryFunc(timeout)

	response := oci_globally_distributed_database.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_globally_distributed_database.OperationStatusInProgress),
			string(oci_globally_distributed_database.OperationStatusAccepted),
			string(oci_globally_distributed_database.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_globally_distributed_database.OperationStatusSucceeded),
			string(oci_globally_distributed_database.OperationStatusFailed),
			string(oci_globally_distributed_database.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_globally_distributed_database.GetWorkRequestRequest{
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
	// Removing check on EntityType
	for _, res := range response.Resources {
		if res.ActionType == action {
			identifier = res.Identifier
			break
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_globally_distributed_database.OperationStatusFailed || response.Status == oci_globally_distributed_database.OperationStatusCanceled {
		return nil, getErrorFromGloballyDistributedDatabasePrivateEndpointWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGloballyDistributedDatabasePrivateEndpointWorkRequest(client *oci_globally_distributed_database.ShardedDatabaseServiceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_globally_distributed_database.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_globally_distributed_database.ListWorkRequestErrorsRequest{
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

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) Get() error {
	request := oci_globally_distributed_database.GetPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.PrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.GetPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateEndpoint
	return nil
}

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_globally_distributed_database.UpdatePrivateEndpointRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	tmp := s.D.Id()
	request.PrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.UpdatePrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateEndpoint
	return nil
}

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) Delete() error {
	request := oci_globally_distributed_database.DeletePrivateEndpointRequest{}

	tmp := s.D.Id()
	request.PrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.DeletePrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := privateEndpointWaitForWorkRequest(workId, "privateendpoint",
		oci_globally_distributed_database.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.PrivateIp != nil {
		s.D.Set("private_ip", *s.Res.PrivateIp)
	}

	s.D.Set("sharded_databases", s.Res.ShardedDatabases)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func PrivateEndpointSummaryToMap(obj oci_globally_distributed_database.PrivateEndpointSummary, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	return result
}

func (s *GloballyDistributedDatabasePrivateEndpointResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_globally_distributed_database.ChangePrivateEndpointCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.PrivateEndpointId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.ChangePrivateEndpointCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPrivateEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database"), oci_globally_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
