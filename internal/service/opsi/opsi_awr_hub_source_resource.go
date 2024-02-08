// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiAwrHubSourceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiAwrHubSource,
		Read:     readOpsiAwrHubSource,
		Update:   updateOpsiAwrHubSource,
		Delete:   deleteOpsiAwrHubSource,
		Schema: map[string]*schema.Schema{
			// Required
			"awr_hub_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"associated_opsi_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"associated_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"awr_hub_opsi_source_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"awr_source_database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hours_since_last_import": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"is_registered_with_awr_hub": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"max_snapshot_identifier": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"min_snapshot_identifier": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"source_mail_box_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
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
			"time_first_snapshot_generated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_snapshot_generated": {
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

func createOpsiAwrHubSource(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.CreateResource(d, sync)
}

func readOpsiAwrHubSource(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

func updateOpsiAwrHubSource(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOpsiAwrHubSource(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpsiAwrHubSourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *oci_opsi.AwrHubSource
	DisableNotFoundRetries bool
}

func (s *OpsiAwrHubSourceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OpsiAwrHubSourceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opsi.AwrHubSourceLifecycleStateCreating),
	}
}

func (s *OpsiAwrHubSourceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opsi.AwrHubSourceLifecycleStateActive),
	}
}

func (s *OpsiAwrHubSourceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opsi.AwrHubSourceLifecycleStateDeleting),
	}
}

func (s *OpsiAwrHubSourceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opsi.AwrHubSourceLifecycleStateDeleted),
	}
}

func (s *OpsiAwrHubSourceResourceCrud) Create() error {
	request := oci_opsi.CreateAwrHubSourceRequest{}

	if associatedOpsiId, ok := s.D.GetOkExists("associated_opsi_id"); ok {
		tmp := associatedOpsiId.(string)
		request.AssociatedOpsiId = &tmp
	}

	if associatedResourceId, ok := s.D.GetOkExists("associated_resource_id"); ok {
		tmp := associatedResourceId.(string)
		request.AssociatedResourceId = &tmp
	}

	if awrHubId, ok := s.D.GetOkExists("awr_hub_id"); ok {
		tmp := awrHubId.(string)
		request.AwrHubId = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_opsi.AwrHubSourceTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.CreateAwrHubSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getAwrHubSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OpsiAwrHubSourceResourceCrud) getAwrHubSourceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	awrHubSourceId, err := awrHubSourceWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*awrHubSourceId)

	return s.Get()
}

func awrHubSourceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "opsi", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_opsi.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func awrHubSourceWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = awrHubSourceWorkRequestShouldRetryFunc(timeout)

	response := oci_opsi.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_opsi.OperationStatusInProgress),
			string(oci_opsi.OperationStatusAccepted),
			string(oci_opsi.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_opsi.OperationStatusSucceeded),
			string(oci_opsi.OperationStatusFailed),
			string(oci_opsi.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_opsi.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_opsi.OperationStatusFailed || response.Status == oci_opsi.OperationStatusCanceled {
		return nil, getErrorFromOpsiAwrHubSourceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiAwrHubSourceWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_opsi.ListWorkRequestErrorsRequest{
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

func (s *OpsiAwrHubSourceResourceCrud) Get() error {
	request := oci_opsi.GetAwrHubSourceRequest{}

	tmp := s.D.Id()
	request.AwrHubSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.GetAwrHubSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AwrHubSource
	return nil
}

func (s *OpsiAwrHubSourceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_opsi.UpdateAwrHubSourceRequest{}

	tmp := s.D.Id()
	request.AwrHubSourceId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_opsi.AwrHubSourceTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.UpdateAwrHubSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAwrHubSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpsiAwrHubSourceResourceCrud) Delete() error {
	request := oci_opsi.DeleteAwrHubSourceRequest{}

	tmp := s.D.Id()
	request.AwrHubSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DeleteAwrHubSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := awrHubSourceWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpsiAwrHubSourceResourceCrud) SetData() error {
	if s.Res.AssociatedOpsiId != nil {
		s.D.Set("associated_opsi_id", *s.Res.AssociatedOpsiId)
	}

	if s.Res.AssociatedResourceId != nil {
		s.D.Set("associated_resource_id", *s.Res.AssociatedResourceId)
	}

	if s.Res.AwrHubId != nil {
		s.D.Set("awr_hub_id", *s.Res.AwrHubId)
	}

	if s.Res.AwrHubOpsiSourceId != nil {
		s.D.Set("awr_hub_opsi_source_id", *s.Res.AwrHubOpsiSourceId)
	}

	if s.Res.AwrSourceDatabaseId != nil {
		s.D.Set("awr_source_database_id", *s.Res.AwrSourceDatabaseId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HoursSinceLastImport != nil {
		s.D.Set("hours_since_last_import", *s.Res.HoursSinceLastImport)
	}

	if s.Res.IsRegisteredWithAwrHub != nil {
		s.D.Set("is_registered_with_awr_hub", *s.Res.IsRegisteredWithAwrHub)
	}

	if s.Res.MaxSnapshotIdentifier != nil {
		s.D.Set("max_snapshot_identifier", *s.Res.MaxSnapshotIdentifier)
	}

	if s.Res.MinSnapshotIdentifier != nil {
		s.D.Set("min_snapshot_identifier", *s.Res.MinSnapshotIdentifier)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.SourceMailBoxUrl != nil {
		s.D.Set("source_mail_box_url", *s.Res.SourceMailBoxUrl)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeFirstSnapshotGenerated != nil {
		s.D.Set("time_first_snapshot_generated", s.Res.TimeFirstSnapshotGenerated.String())
	}

	if s.Res.TimeLastSnapshotGenerated != nil {
		s.D.Set("time_last_snapshot_generated", s.Res.TimeLastSnapshotGenerated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}

func AwrHubSourceSummaryToMap(obj oci_opsi.AwrHubSourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssociatedOpsiId != nil {
		result["associated_opsi_id"] = string(*obj.AssociatedOpsiId)
	}

	if obj.AssociatedResourceId != nil {
		result["associated_resource_id"] = string(*obj.AssociatedResourceId)
	}

	if obj.AwrHubId != nil {
		result["awr_hub_id"] = string(*obj.AwrHubId)
	}

	if obj.AwrHubOpsiSourceId != nil {
		result["awr_hub_opsi_source_id"] = string(*obj.AwrHubOpsiSourceId)
	}

	if obj.AwrSourceDatabaseId != nil {
		result["awr_source_database_id"] = string(*obj.AwrSourceDatabaseId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.HoursSinceLastImport != nil {
		result["hours_since_last_import"] = float64(*obj.HoursSinceLastImport)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsRegisteredWithAwrHub != nil {
		result["is_registered_with_awr_hub"] = bool(*obj.IsRegisteredWithAwrHub)
	}

	if obj.MaxSnapshotIdentifier != nil {
		result["max_snapshot_identifier"] = float32(*obj.MaxSnapshotIdentifier)
	}

	if obj.MinSnapshotIdentifier != nil {
		result["min_snapshot_identifier"] = float32(*obj.MinSnapshotIdentifier)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.SourceMailBoxUrl != nil {
		result["source_mail_box_url"] = string(*obj.SourceMailBoxUrl)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeFirstSnapshotGenerated != nil {
		result["time_first_snapshot_generated"] = obj.TimeFirstSnapshotGenerated.String()
	}

	if obj.TimeLastSnapshotGenerated != nil {
		result["time_last_snapshot_generated"] = obj.TimeLastSnapshotGenerated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *OpsiAwrHubSourceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_opsi.ChangeAwrHubSourceCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AwrHubSourceId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangeAwrHubSourceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAwrHubSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
