// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalMySqlDatabaseConnectorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalMySqlDatabaseConnector,
		Read:     readDatabaseManagementExternalMySqlDatabaseConnector,
		Update:   updateDatabaseManagementExternalMySqlDatabaseConnector,
		Delete:   deleteDatabaseManagementExternalMySqlDatabaseConnector,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"connector_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"credential_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"external_database_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"macs_agent_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"network_protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"ssl_secret_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"is_test_connection_param": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"check_connection_status_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"associated_services": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connector_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"credential_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"macs_agent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_protocol": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"source_database": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_database_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_secret_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_secret_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_connection_status_updated": {
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

func createDatabaseManagementExternalMySqlDatabaseConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("check_connection_status_trigger"); ok {
		err := sync.CheckExternalMySqlDatabaseConnectorConnectionStatus()
		if err != nil {
			return err
		}
	}
	return nil

}

func readDatabaseManagementExternalMySqlDatabaseConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementExternalMySqlDatabaseConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	if _, ok := sync.D.GetOkExists("check_connection_status_trigger"); ok && sync.D.HasChange("check_connection_status_trigger") {
		oldRaw, newRaw := sync.D.GetChange("check_connection_status_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.CheckExternalMySqlDatabaseConnectorConnectionStatus()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("check_connection_status_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteDatabaseManagementExternalMySqlDatabaseConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ExternalMySqlDatabaseConnector
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.LifecycleStatesCreating),
	}
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.LifecycleStatesActive),
	}
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.LifecycleStatesDeleting),
	}
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.LifecycleStatesDeleted),
	}
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) Create() error {
	request := oci_database_management.CreateExternalMySqlDatabaseConnectorRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if connectorDetails, ok := s.D.GetOkExists("connector_details"); ok {
		if tmpList := connectorDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connector_details", 0)
			tmp, err := s.mapToCreateMySqlDatabaseConnectorDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConnectorDetails = &tmp
		}
	}

	if isTestConnectionParam, ok := s.D.GetOkExists("is_test_connection_param"); ok {
		tmp := isTestConnectionParam.(bool)
		request.IsTestConnectionParam = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.CreateExternalMySqlDatabaseConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalMySqlDatabaseConnector
	return nil
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) getExternalMySqlDatabaseConnectorFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	externalMySqlDatabaseConnectorId, err := externalMySqlDatabaseConnectorWaitForWorkRequest(workId, "connector",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*externalMySqlDatabaseConnectorId)

	return s.Get()
}

func externalMySqlDatabaseConnectorWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "database_management", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_database_management.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func externalMySqlDatabaseConnectorWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = externalMySqlDatabaseConnectorWorkRequestShouldRetryFunc(timeout)

	response := oci_database_management.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_database_management.WorkRequestStatusInProgress),
			string(oci_database_management.WorkRequestStatusAccepted),
			string(oci_database_management.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_database_management.WorkRequestStatusSucceeded),
			string(oci_database_management.WorkRequestStatusFailed),
			string(oci_database_management.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_database_management.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_database_management.WorkRequestStatusFailed || response.Status == oci_database_management.WorkRequestStatusCanceled {
		return nil, getErrorFromDatabaseManagementExternalMySqlDatabaseConnectorWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementExternalMySqlDatabaseConnectorWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_database_management.ListWorkRequestErrorsRequest{
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

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) Get() error {
	request := oci_database_management.GetExternalMySqlDatabaseConnectorRequest{}

	tmp := s.D.Id()
	request.ExternalMySqlDatabaseConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetExternalMySqlDatabaseConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalMySqlDatabaseConnector
	return nil
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) Update() error {
	request := oci_database_management.UpdateExternalMysqlDatabaseConnectorRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if connectorDetails, ok := s.D.GetOkExists("connector_details"); ok {
		if tmpList := connectorDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connector_details", 0)
			tmp, err := s.mapToUpdateMySqlDatabaseConnectorDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConnectorDetails = &tmp
		}
	}

	tmp := s.D.Id()
	request.ExternalMySqlDatabaseConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalMysqlDatabaseConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getExternalMySqlDatabaseConnectorFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) Delete() error {
	request := oci_database_management.DeleteExternalMySqlDatabaseConnectorRequest{}

	tmp := s.D.Id()
	request.ExternalMySqlDatabaseConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DeleteExternalMySqlDatabaseConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := externalMySqlDatabaseConnectorWaitForWorkRequest(workId, "connector",
		oci_database_management.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) SetData() error {
	if s.Res.AssociatedServices != nil {
		s.D.Set("associated_services", *s.Res.AssociatedServices)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStatus != nil {
		s.D.Set("connection_status", *s.Res.ConnectionStatus)
	}

	s.D.Set("connector_type", s.Res.ConnectorType)

	s.D.Set("credential_type", s.Res.CredentialType)

	if s.Res.ExternalDatabaseId != nil {
		s.D.Set("external_database_id", *s.Res.ExternalDatabaseId)
	}

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.MacsAgentId != nil {
		s.D.Set("macs_agent_id", *s.Res.MacsAgentId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("network_protocol", s.Res.NetworkProtocol)

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.SourceDatabase != nil {
		s.D.Set("source_database", *s.Res.SourceDatabase)
	}

	s.D.Set("source_database_type", s.Res.SourceDatabaseType)

	if s.Res.SslSecretId != nil {
		s.D.Set("ssl_secret_id", *s.Res.SslSecretId)
	}

	if s.Res.SslSecretName != nil {
		s.D.Set("ssl_secret_name", *s.Res.SslSecretName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeConnectionStatusUpdated != nil {
		s.D.Set("time_connection_status_updated", s.Res.TimeConnectionStatusUpdated.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) CheckExternalMySqlDatabaseConnectorConnectionStatus() error {
	request := oci_database_management.CheckExternalMySqlDatabaseConnectorConnectionStatusRequest{}

	idTmp := s.D.Id()
	request.ExternalMySqlDatabaseConnectorId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	_, err := s.Client.CheckExternalMySqlDatabaseConnectorConnectionStatus(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("check_connection_status_trigger")
	s.D.Set("check_connection_status_trigger", val)

	//s.Res = &response.ExternalMySqlDatabaseConnector
	return nil
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) mapToCreateMySqlDatabaseConnectorDetails(fieldKeyFormat string) (oci_database_management.CreateMySqlDatabaseConnectorDetails, error) {
	result := oci_database_management.CreateMySqlDatabaseConnectorDetails{}

	if credentialType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type")); ok {
		result.CredentialType = oci_database_management.MySqlCredTypeEnum(credentialType.(string))
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if externalDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "external_database_id")); ok {
		tmp := externalDatabaseId.(string)
		result.ExternalDatabaseId = &tmp
	}

	if hostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name")); ok {
		tmp := hostName.(string)
		result.HostName = &tmp
	}

	if macsAgentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "macs_agent_id")); ok {
		tmp := macsAgentId.(string)
		result.MacsAgentId = &tmp
	}

	if networkProtocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_protocol")); ok {
		result.NetworkProtocol = oci_database_management.MySqlNetworkProtocolTypeEnum(networkProtocol.(string))
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if sslSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_secret_id")); ok {
		tmp := sslSecretId.(string)
		result.SslSecretId = &tmp
	}

	return result, nil
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorResourceCrud) mapToUpdateMySqlDatabaseConnectorDetails(fieldKeyFormat string) (oci_database_management.UpdateMySqlDatabaseConnectorDetails, error) {
	result := oci_database_management.UpdateMySqlDatabaseConnectorDetails{}

	if credentialType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type")); ok {
		result.CredentialType = oci_database_management.MySqlCredTypeEnum(credentialType.(string))
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if externalDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "external_database_id")); ok {
		tmp := externalDatabaseId.(string)
		result.ExternalDatabaseId = &tmp
	}

	if hostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name")); ok {
		tmp := hostName.(string)
		result.HostName = &tmp
	}

	if macsAgentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "macs_agent_id")); ok {
		tmp := macsAgentId.(string)
		result.MacsAgentId = &tmp
	}

	if networkProtocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_protocol")); ok {
		result.NetworkProtocol = oci_database_management.MySqlNetworkProtocolTypeEnum(networkProtocol.(string))
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if sslSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_secret_id")); ok {
		tmp := sslSecretId.(string)
		result.SslSecretId = &tmp
	}

	return result, nil
}

func CreateMySqlDatabaseConnectorDetailsToMap(obj *oci_database_management.CreateMySqlDatabaseConnectorDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["credential_type"] = string(obj.CredentialType)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExternalDatabaseId != nil {
		result["external_database_id"] = string(*obj.ExternalDatabaseId)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.MacsAgentId != nil {
		result["macs_agent_id"] = string(*obj.MacsAgentId)
	}

	result["network_protocol"] = string(obj.NetworkProtocol)

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.SslSecretId != nil {
		result["ssl_secret_id"] = string(*obj.SslSecretId)
	}

	return result
}

func MySqlDatabaseConnectorSummaryToMap(obj oci_database_management.MySqlDatabaseConnectorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssociatedServices != nil {
		result["associated_services"] = string(*obj.AssociatedServices)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConnectionStatus != nil {
		result["connection_status"] = string(*obj.ConnectionStatus)
	}

	result["connector_type"] = string(obj.ConnectorType)

	result["credential_type"] = string(obj.CredentialType)

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MacsAgentId != nil {
		result["macs_agent_id"] = string(*obj.MacsAgentId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["network_protocol"] = string(obj.NetworkProtocol)

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.SourceDatabase != nil {
		result["source_database"] = string(*obj.SourceDatabase)
	}

	result["source_database_type"] = string(obj.SourceDatabaseType)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeConnectionStatusUpdated != nil {
		result["time_connection_status_updated"] = obj.TimeConnectionStatusUpdated.String()
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
