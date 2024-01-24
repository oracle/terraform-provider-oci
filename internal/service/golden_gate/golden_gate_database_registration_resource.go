// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"
)

func GoldenGateDatabaseRegistrationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createGoldenGateDatabaseRegistration,
		Read:     readGoldenGateDatabaseRegistration,
		Update:   updateGoldenGateDatabaseRegistration,
		Delete:   deleteGoldenGateDatabaseRegistration,
		Schema: map[string]*schema.Schema{
			// Required
			"alias_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"connection_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database_id": {
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
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"secret_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"session_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vault_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"wallet": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rce_private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_id": {
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

func createGoldenGateDatabaseRegistration(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDatabaseRegistrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.CreateResource(d, sync)
}

func readGoldenGateDatabaseRegistration(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDatabaseRegistrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

func updateGoldenGateDatabaseRegistration(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDatabaseRegistrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteGoldenGateDatabaseRegistration(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDatabaseRegistrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GoldenGateDatabaseRegistrationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_golden_gate.GoldenGateClient
	Res                    *oci_golden_gate.DatabaseRegistration
	DisableNotFoundRetries bool
}

func (s *GoldenGateDatabaseRegistrationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GoldenGateDatabaseRegistrationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateCreating),
		string(oci_golden_gate.LifecycleStateInProgress),
	}
}

func (s *GoldenGateDatabaseRegistrationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateActive),
		string(oci_golden_gate.LifecycleStateNeedsAttention),
		string(oci_golden_gate.LifecycleStateSucceeded),
	}
}

func (s *GoldenGateDatabaseRegistrationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateDeleting),
	}
}

func (s *GoldenGateDatabaseRegistrationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateDeleted),
	}
}

func (s *GoldenGateDatabaseRegistrationResourceCrud) Create() error {
	request := oci_golden_gate.CreateDatabaseRegistrationRequest{}

	if aliasName, ok := s.D.GetOkExists("alias_name"); ok {
		tmp := aliasName.(string)
		request.AliasName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
		tmp := connectionString.(string)
		request.ConnectionString = &tmp
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
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

	if fqdn, ok := s.D.GetOkExists("fqdn"); ok {
		tmp := fqdn.(string)
		request.Fqdn = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	if secretCompartmentId, ok := s.D.GetOkExists("secret_compartment_id"); ok {
		tmp := secretCompartmentId.(string)
		request.SecretCompartmentId = &tmp
	}

	if sessionMode, ok := s.D.GetOkExists("session_mode"); ok {
		request.SessionMode = oci_golden_gate.CreateDatabaseRegistrationDetailsSessionModeEnum(sessionMode.(string))
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if username, ok := s.D.GetOkExists("username"); ok {
		tmp := username.(string)
		request.Username = &tmp
	}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	if wallet, ok := s.D.GetOkExists("wallet"); ok {
		tmp := wallet.(string)
		request.Wallet = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.CreateDatabaseRegistration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDatabaseRegistrationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GoldenGateDatabaseRegistrationResourceCrud) getDatabaseRegistrationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_golden_gate.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	databaseRegistrationId, err := databaseRegistrationWaitForWorkRequest(workId, "databaseregistration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] operation failed: %v for identifier: %v\n", workId, databaseRegistrationId)
		return err
	}
	s.D.SetId(*databaseRegistrationId)

	return s.Get()
}

func databaseRegistrationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "golden_gate", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_golden_gate.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func databaseRegistrationWaitForWorkRequest(wId *string, entityType string, action oci_golden_gate.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_golden_gate.GoldenGateClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "golden_gate")
	retryPolicy.ShouldRetryOperation = databaseRegistrationWorkRequestShouldRetryFunc(timeout)

	response := oci_golden_gate.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_golden_gate.OperationStatusInProgress),
			string(oci_golden_gate.OperationStatusAccepted),
		},
		Target: []string{
			string(oci_golden_gate.OperationStatusSucceeded),
			string(oci_golden_gate.OperationStatusFailed),
			string(oci_golden_gate.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_golden_gate.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_golden_gate.OperationStatusFailed || response.Status == oci_golden_gate.OperationStatusCanceled {
		return nil, getErrorFromGoldenGateDatabaseRegistrationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGoldenGateDatabaseRegistrationWorkRequest(client *oci_golden_gate.GoldenGateClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_golden_gate.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_golden_gate.ListWorkRequestErrorsRequest{
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

func (s *GoldenGateDatabaseRegistrationResourceCrud) Get() error {
	request := oci_golden_gate.GetDatabaseRegistrationRequest{}

	tmp := s.D.Id()
	request.DatabaseRegistrationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.GetDatabaseRegistration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseRegistration
	return nil
}

func (s *GoldenGateDatabaseRegistrationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_golden_gate.UpdateDatabaseRegistrationRequest{}

	if aliasName, ok := s.D.GetOkExists("alias_name"); ok {
		tmp := aliasName.(string)
		request.AliasName = &tmp
	}

	if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
		tmp := connectionString.(string)
		request.ConnectionString = &tmp
	}

	tmp := s.D.Id()
	request.DatabaseRegistrationId = &tmp

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

	if fqdn, ok := s.D.GetOkExists("fqdn"); ok {
		tmp := fqdn.(string)
		request.Fqdn = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	if sessionMode, ok := s.D.GetOkExists("session_mode"); ok {
		request.SessionMode = oci_golden_gate.UpdateDatabaseRegistrationDetailsSessionModeEnum(sessionMode.(string))
	}

	if username, ok := s.D.GetOkExists("username"); ok {
		tmp := username.(string)
		request.Username = &tmp
	}

	if wallet, ok := s.D.GetOkExists("wallet"); ok {
		tmp := wallet.(string)
		request.Wallet = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.UpdateDatabaseRegistration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseRegistrationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GoldenGateDatabaseRegistrationResourceCrud) Delete() error {
	request := oci_golden_gate.DeleteDatabaseRegistrationRequest{}

	tmp := s.D.Id()
	request.DatabaseRegistrationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.DeleteDatabaseRegistration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := databaseRegistrationWaitForWorkRequest(workId, "databaseregistration",
		oci_golden_gate.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GoldenGateDatabaseRegistrationResourceCrud) SetData() error {
	if s.Res.AliasName != nil {
		s.D.Set("alias_name", *s.Res.AliasName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionString != nil {
		s.D.Set("connection_string", *s.Res.ConnectionString)
	}

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
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

	if s.Res.Fqdn != nil {
		s.D.Set("fqdn", *s.Res.Fqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.KeyId != nil {
		s.D.Set("key_id", *s.Res.KeyId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RcePrivateIp != nil {
		s.D.Set("rce_private_ip", *s.Res.RcePrivateIp)
	}

	if s.Res.SecretCompartmentId != nil {
		s.D.Set("secret_compartment_id", *s.Res.SecretCompartmentId)
	}

	if s.Res.SecretId != nil {
		s.D.Set("secret_id", *s.Res.SecretId)
	}

	s.D.Set("session_mode", s.Res.SessionMode)

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

	if s.Res.Username != nil {
		s.D.Set("username", *s.Res.Username)
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	return nil
}

func DatabaseRegistrationSummaryToMap(obj oci_golden_gate.DatabaseRegistrationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AliasName != nil {
		result["alias_name"] = string(*obj.AliasName)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConnectionString != nil {
		result["connection_string"] = string(*obj.ConnectionString)
	}

	if obj.DatabaseId != nil {
		result["database_id"] = string(*obj.DatabaseId)
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

	if obj.Fqdn != nil {
		result["fqdn"] = string(*obj.Fqdn)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.SecretId != nil {
		result["secret_id"] = string(*obj.SecretId)
	}

	result["session_mode"] = string(obj.SessionMode)

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

	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}

	return result
}

func (s *GoldenGateDatabaseRegistrationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_golden_gate.ChangeDatabaseRegistrationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatabaseRegistrationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.ChangeDatabaseRegistrationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, changeWorkRequestErr := databaseRegistrationWaitForWorkRequest(workId, "databaseregistration",
		oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	return changeWorkRequestErr
}
