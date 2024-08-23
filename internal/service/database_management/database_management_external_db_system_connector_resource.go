// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalDbSystemConnectorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalDbSystemConnector,
		Read:     readDatabaseManagementExternalDbSystemConnector,
		Update:   updateDatabaseManagementExternalDbSystemConnector,
		Delete:   deleteDatabaseManagementExternalDbSystemConnector,
		Schema: map[string]*schema.Schema{
			// Required
			"connector_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"MACS",
				}, true),
			},
			"external_db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"agent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_failure_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_info": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"component_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"connection_credentials": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"credential_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"credential_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"named_credential_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"password_secret_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"role": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssl_secret_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"user_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"connection_string": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"host_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"hosts": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"port": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"protocol": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"service": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"database_credential": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"credential_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"named_credential_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"password": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"password_secret_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"username": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"connection_status": {
				Type:     schema.TypeString,
				Computed: true,
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
			"lifecycle_details": {
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
			"time_connection_status_last_updated": {
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

func createDatabaseManagementExternalDbSystemConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbSystemConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalDbSystemConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbSystemConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementExternalDbSystemConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbSystemConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalDbSystemConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbSystemConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementExternalDbSystemConnectorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ExternalDbSystemConnector
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) ID() string {
	externalDbSystemConnector := *s.Res
	return *externalDbSystemConnector.GetId()
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.ExternalDbSystemConnectorLifecycleStateCreating),
	}
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.ExternalDbSystemConnectorLifecycleStateNotConnected),
		string(oci_database_management.ExternalDbSystemConnectorLifecycleStateActive),
	}
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.ExternalDbSystemConnectorLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.ExternalDbSystemConnectorLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) Create() error {
	request := oci_database_management.CreateExternalDbSystemConnectorRequest{}
	err := s.populateTopLevelPolymorphicCreateExternalDbSystemConnectorRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.CreateExternalDbSystemConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalDbSystemConnector
	return nil
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) getExternalDbSystemConnectorFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	externalDbSystemConnectorId, err := externalDbSystemConnectorWaitForWorkRequest(workId, "connector",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*externalDbSystemConnectorId)

	return s.Get()
}

func externalDbSystemConnectorWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func externalDbSystemConnectorWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = externalDbSystemConnectorWorkRequestShouldRetryFunc(timeout)

	response := oci_database_management.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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
		return nil, getErrorFromDatabaseManagementExternalDbSystemConnectorWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementExternalDbSystemConnectorWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) Get() error {
	request := oci_database_management.GetExternalDbSystemConnectorRequest{}

	tmp := s.D.Id()
	request.ExternalDbSystemConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetExternalDbSystemConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalDbSystemConnector
	return nil
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) Update() error {
	request := oci_database_management.UpdateExternalDbSystemConnectorRequest{}
	err := s.populateTopLevelPolymorphicUpdateExternalDbSystemConnectorRequest(&request)
	if err != nil {
		return err
	}

	tmp := s.D.Id()
	request.ExternalDbSystemConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalDbSystemConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getExternalDbSystemConnectorFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) Delete() error {
	request := oci_database_management.DeleteExternalDbSystemConnectorRequest{}

	tmp := s.D.Id()
	request.ExternalDbSystemConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	_, err := s.Client.DeleteExternalDbSystemConnector(context.Background(), request)
	return err
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_database_management.ExternalDbSystemMacsConnector:
		s.D.Set("connector_type", "MACS")

		if v.AgentId != nil {
			s.D.Set("agent_id", *v.AgentId)
		}

		if v.ConnectionInfo != nil {
			connectionInfoArray := []interface{}{}
			if connectionInfoMap := ExternalDbSystemConnectionInfoToMap(&v.ConnectionInfo); connectionInfoMap != nil {
				connectionInfoArray = append(connectionInfoArray, connectionInfoMap)
			}
			s.D.Set("connection_info", connectionInfoArray)
		} else {
			s.D.Set("connection_info", nil)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.ConnectionFailureMessage != nil {
			s.D.Set("connection_failure_message", *v.ConnectionFailureMessage)
		}

		if v.ConnectionStatus != nil {
			s.D.Set("connection_status", *v.ConnectionStatus)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.ExternalDbSystemId != nil {
			s.D.Set("external_db_system_id", *v.ExternalDbSystemId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeConnectionStatusLastUpdated != nil {
			s.D.Set("time_connection_status_last_updated", v.TimeConnectionStatusLastUpdated.String())
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'connector_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) mapToAsmConnectionCredentials(fieldKeyFormat string) (oci_database_management.AsmConnectionCredentials, error) {
	var baseObject oci_database_management.AsmConnectionCredentials
	//discriminator
	credentialTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type"))
	var credentialType string
	if ok {
		credentialType = credentialTypeRaw.(string)
	} else {
		credentialType = "DETAILS" // default value
	}
	switch strings.ToLower(credentialType) {
	case strings.ToLower("DETAILS"):
		details := oci_database_management.AsmConnectionCredentialsByDetails{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database_management.AsmConnectionCredentialsByDetailsRoleEnum(role.(string))
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		baseObject = details
	case strings.ToLower("NAME_REFERENCE"):
		details := oci_database_management.AsmConnectionCredentailsByName{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_type '%v' was specified", credentialType)
	}
	return baseObject, nil
}

func AsmConnectionCredentialsToMap(obj *oci_database_management.AsmConnectionCredentials) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_management.AsmConnectionCredentialsByDetails:
		result["credential_type"] = "DETAILS"

		if v.CredentialName != nil {
			result["credential_name"] = string(*v.CredentialName)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		result["role"] = string(v.Role)

		if v.UserName != nil {
			result["user_name"] = string(*v.UserName)
		}
	case oci_database_management.AsmConnectionCredentailsByName:
		result["credential_type"] = "NAME_REFERENCE"

		if v.CredentialName != nil {
			result["credential_name"] = string(*v.CredentialName)
		}
	default:
		log.Printf("[WARN] Received 'credential_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) mapToAsmConnectionString(fieldKeyFormat string) (oci_database_management.AsmConnectionString, error) {
	result := oci_database_management.AsmConnectionString{}

	if hosts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hosts")); ok {
		interfaces := hosts.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hosts")) {
			result.Hosts = tmp
		}
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_database_management.AsmConnectionStringProtocolEnum(protocol.(string))
	}

	if service, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service")); ok {
		tmp := service.(string)
		result.Service = &tmp
	}

	return result, nil
}

func AsmConnectionStringToMap(obj *oci_database_management.AsmConnectionString) map[string]interface{} {
	result := map[string]interface{}{}

	result["hosts"] = obj.Hosts
	result["hosts"] = obj.Hosts

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	result["protocol"] = string(obj.Protocol)

	if obj.Service != nil {
		result["service"] = string(*obj.Service)
	}

	return result
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) mapToDatabaseConnectionCredentials(fieldKeyFormat string) (oci_database_management.DatabaseConnectionCredentials, error) {
	var baseObject oci_database_management.DatabaseConnectionCredentials
	//discriminator
	credentialTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type"))
	var credentialType string
	if ok {
		credentialType = credentialTypeRaw.(string)
	} else {
		credentialType = "DETAILS" // default value
	}
	switch strings.ToLower(credentialType) {
	case strings.ToLower("DETAILS"):
		details := oci_database_management.DatabaseConnectionCredentialsByDetails{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database_management.DatabaseConnectionCredentialsByDetailsRoleEnum(role.(string))
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		baseObject = details
	case strings.ToLower("NAMED_CREDENTIAL"):
		details := oci_database_management.DatabaseNamedCredentialConnectionDetails{}
		if namedCredentialId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "named_credential_id")); ok {
			tmp := namedCredentialId.(string)
			details.NamedCredentialId = &tmp
		}
		baseObject = details
	case strings.ToLower("NAME_REFERENCE"):
		details := oci_database_management.DatabaseConnectionCredentailsByName{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		baseObject = details
	case strings.ToLower("SSL_DETAILS"):
		details := oci_database_management.DatabaseSslConnectionCredentials{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database_management.DatabaseSslConnectionCredentialsRoleEnum(role.(string))
		}
		if sslSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_secret_id")); ok {
			tmp := sslSecretId.(string)
			details.SslSecretId = &tmp
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_type '%v' was specified", credentialType)
	}
	return baseObject, nil
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) mapToDatabaseConnectionString(fieldKeyFormat string) (oci_database_management.DatabaseConnectionString, error) {
	result := oci_database_management.DatabaseConnectionString{}

	if hostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name")); ok {
		tmp := hostName.(string)
		result.HostName = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_database_management.DatabaseConnectionStringProtocolEnum(protocol.(string))
	}

	if service, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service")); ok {
		tmp := service.(string)
		result.Service = &tmp
	}

	return result, nil
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) mapToDatabaseCredentialDetails(fieldKeyFormat string) (oci_database_management.DatabaseCredentialDetails, error) {
	var baseObject oci_database_management.DatabaseCredentialDetails
	//discriminator
	credentialTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type"))
	var credentialType string
	if ok {
		credentialType = credentialTypeRaw.(string)
	} else {
		credentialType = "" // default value
	}
	switch strings.ToLower(credentialType) {
	case strings.ToLower("NAMED_CREDENTIAL"):
		details := oci_database_management.DatabaseNamedCredentialDetails{}
		if namedCredentialId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "named_credential_id")); ok {
			tmp := namedCredentialId.(string)
			details.NamedCredentialId = &tmp
		}
		baseObject = details
	case strings.ToLower("PASSWORD"):
		details := oci_database_management.DatabasePasswordCredentialDetails{}
		if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database_management.DatabasePasswordCredentialDetailsRoleEnum(role.(string))
		}
		if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		baseObject = details
	case strings.ToLower("SECRET"):
		details := oci_database_management.DatabaseSecretCredentialDetails{}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database_management.DatabaseSecretCredentialDetailsRoleEnum(role.(string))
		}
		if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_type '%v' was specified", credentialType)
	}
	return baseObject, nil
}

func DatabaseCredentialDetailsToMap(obj *oci_database_management.DatabaseCredentialDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_management.DatabaseNamedCredentialDetails:
		result["credential_type"] = "NAMED_CREDENTIAL"

		if v.NamedCredentialId != nil {
			result["named_credential_id"] = string(*v.NamedCredentialId)
		}
	case oci_database_management.DatabasePasswordCredentialDetails:
		result["credential_type"] = "PASSWORD"

		if v.Password != nil {
			result["password"] = string(*v.Password)
		}

		result["role"] = string(v.Role)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_database_management.DatabaseSecretCredentialDetails:
		result["credential_type"] = "SECRET"

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		result["role"] = string(v.Role)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	default:
		log.Printf("[WARN] Received 'credential_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) mapToExternalDbSystemConnectionInfo(fieldKeyFormat string) (oci_database_management.ExternalDbSystemConnectionInfo, error) {
	var baseObject oci_database_management.ExternalDbSystemConnectionInfo
	//discriminator
	componentTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "component_type"))
	var componentType string
	if ok {
		componentType = componentTypeRaw.(string)
	} else {
		componentType = "" // default value
	}
	switch strings.ToLower(componentType) {
	case strings.ToLower("ASM"):
		details := oci_database_management.ExternalAsmConnectionInfo{}
		if connectionCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_credentials")); ok {
			if tmpList := connectionCredentials.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_credentials"), 0)
				tmp, err := s.mapToAsmConnectionCredentials(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connection_credentials, encountered error: %v", err)
				}
				details.ConnectionCredentials = tmp
			}
		}
		if connectionString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_string")); ok {
			if tmpList := connectionString.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_string"), 0)
				tmp, err := s.mapToAsmConnectionString(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connection_string, encountered error: %v", err)
				}
				details.ConnectionString = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("DATABASE"):
		details := oci_database_management.ExternalDatabaseConnectionInfo{}
		if connectionCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_credentials")); ok {
			if tmpList := connectionCredentials.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_credentials"), 0)
				tmp, err := s.mapToDatabaseConnectionCredentials(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connection_credentials, encountered error: %v", err)
				}
				details.ConnectionCredentials = tmp
			}
		}
		if connectionString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_string")); ok {
			if tmpList := connectionString.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_string"), 0)
				tmp, err := s.mapToDatabaseConnectionString(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connection_string, encountered error: %v", err)
				}
				details.ConnectionString = &tmp
			}
		}
		if databaseCredential, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_credential")); ok {
			if tmpList := databaseCredential.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database_credential"), 0)
				tmp, err := s.mapToDatabaseCredentialDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert database_credential, encountered error: %v", err)
				}
				details.DatabaseCredential = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown component_type '%v' was specified", componentType)
	}
	return baseObject, nil
}

func ExternalDbSystemConnectionInfoToMap(obj *oci_database_management.ExternalDbSystemConnectionInfo) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_management.ExternalAsmConnectionInfo:
		result["component_type"] = "ASM"

		if v.ConnectionCredentials != nil {
			connectionCredentialsArray := []interface{}{}
			if connectionCredentialsMap := AsmConnectionCredentialsToMap(&v.ConnectionCredentials); connectionCredentialsMap != nil {
				connectionCredentialsArray = append(connectionCredentialsArray, connectionCredentialsMap)
			}
			result["connection_credentials"] = connectionCredentialsArray
		}

		if v.ConnectionString != nil {
			result["connection_string"] = []interface{}{AsmConnectionStringToMap(v.ConnectionString)}
		}
	case oci_database_management.ExternalDatabaseConnectionInfo:
		result["component_type"] = "DATABASE"

		if v.ConnectionCredentials != nil {
			connectionCredentialsArray := []interface{}{}
			if connectionCredentialsMap := DatabaseConnectionCredentialsToMap(&v.ConnectionCredentials); connectionCredentialsMap != nil {
				connectionCredentialsArray = append(connectionCredentialsArray, connectionCredentialsMap)
			}
			result["connection_credentials"] = connectionCredentialsArray
		}

		if v.ConnectionString != nil {
			result["connection_string"] = []interface{}{DatabaseConnectionStringToMap(v.ConnectionString)}
		}

		if v.DatabaseCredential != nil {
			databaseCredentialArray := []interface{}{}
			if databaseCredentialMap := DatabaseCredentialDetailsToMap(&v.DatabaseCredential); databaseCredentialMap != nil {
				databaseCredentialArray = append(databaseCredentialArray, databaseCredentialMap)
			}
			result["database_credential"] = databaseCredentialArray
		}
	default:
		log.Printf("[WARN] Received 'component_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) populateTopLevelPolymorphicCreateExternalDbSystemConnectorRequest(request *oci_database_management.CreateExternalDbSystemConnectorRequest) error {
	//discriminator
	connectorTypeRaw, ok := s.D.GetOkExists("connector_type")
	var connectorType string
	if ok {
		connectorType = connectorTypeRaw.(string)
	} else {
		connectorType = "" // default value
	}
	switch strings.ToLower(connectorType) {
	case strings.ToLower("MACS"):
		details := oci_database_management.CreateExternalDbSystemMacsConnectorDetails{}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}

		if externalDbSystemId, ok := s.D.GetOkExists("external_db_system_id"); ok {
			tmp := externalDbSystemId.(string)
			details.ExternalDbSystemId = &tmp
		}

		if agentId, ok := s.D.GetOkExists("agent_id"); ok {
			tmp := agentId.(string)
			details.AgentId = &tmp
		}

		if connectionInfo, ok := s.D.GetOkExists("connection_info"); ok {
			if tmpList := connectionInfo.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_info", 0)
				tmp, err := s.mapToExternalDbSystemConnectionInfo(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionInfo = tmp
			}
		}

		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}

		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}

		request.CreateExternalDbSystemConnectorDetails = details
	default:
		return fmt.Errorf("unknown connectorType '%v' was specified", connectorType)
	}
	return nil
}

func (s *DatabaseManagementExternalDbSystemConnectorResourceCrud) populateTopLevelPolymorphicUpdateExternalDbSystemConnectorRequest(request *oci_database_management.UpdateExternalDbSystemConnectorRequest) error {
	//discriminator
	connectorTypeRaw, ok := s.D.GetOkExists("connector_type")
	var connectorType string
	if ok {
		connectorType = connectorTypeRaw.(string)
	} else {
		connectorType = "" // default value
	}
	switch strings.ToLower(connectorType) {
	case strings.ToLower("MACS"):
		details := oci_database_management.UpdateExternalDbSystemMacsConnectorDetails{}

		if connectionInfo, ok := s.D.GetOkExists("connection_info"); ok {
			if tmpList := connectionInfo.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_info", 0)
				tmp, err := s.mapToExternalDbSystemConnectionInfo(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionInfo = tmp
			}
		}

		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}

		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}

		request.UpdateExternalDbSystemConnectorDetails = details
	default:
		return fmt.Errorf("unknown connectorType '%v' was specified", connectorType)
	}
	return nil
}

func ExternalDbSystemConnectorSummaryToMap(obj oci_database_management.ExternalDbSystemConnectorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["connector_type"] = string(obj.ConnectorType)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExternalDbSystemId != nil {
		result["external_db_system_id"] = string(*obj.ExternalDbSystemId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
