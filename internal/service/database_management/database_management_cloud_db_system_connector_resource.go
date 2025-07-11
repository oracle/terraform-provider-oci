// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementCloudDbSystemConnectorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementCloudDbSystemConnector,
		Read:     readDatabaseManagementCloudDbSystemConnector,
		Update:   updateDatabaseManagementCloudDbSystemConnector,
		Delete:   deleteDatabaseManagementCloudDbSystemConnector,
		Schema: map[string]*schema.Schema{
			// Required
			"cloud_db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"connector_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"MACS",
				}, true),
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
				Computed: true,
				Optional: true,
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

func createDatabaseManagementCloudDbSystemConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementCloudDbSystemConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementCloudDbSystemConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementCloudDbSystemConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementCloudDbSystemConnectorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.CloudDbSystemConnector
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) ID() string {
	cloudDbSystemConnector := *s.Res
	return *cloudDbSystemConnector.GetId()
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.CloudDbSystemConnectorLifecycleStateCreating),
	}
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.CloudDbSystemConnectorLifecycleStateNotConnected),
		string(oci_database_management.CloudDbSystemConnectorLifecycleStateActive),
	}
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.CloudDbSystemConnectorLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.CloudDbSystemConnectorLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) Create() error {
	request := oci_database_management.CreateCloudDbSystemConnectorRequest{}
	err := s.populateTopLevelPolymorphicCreateCloudDbSystemConnectorRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.CreateCloudDbSystemConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudDbSystemConnector
	return nil
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) populateTopLevelPolymorphicCreateCloudDbSystemConnectorRequest(request *oci_database_management.CreateCloudDbSystemConnectorRequest) error {
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
		details := oci_database_management.CreateCloudDbSystemMacsConnectorDetails{}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}

		if cloudDbSystemId, ok := s.D.GetOkExists("cloud_db_system_id"); ok {
			tmp := cloudDbSystemId.(string)
			details.CloudDbSystemId = &tmp
		}

		if agentId, ok := s.D.GetOkExists("agent_id"); ok {
			tmp := agentId.(string)
			details.AgentId = &tmp
		}

		if connectionInfo, ok := s.D.GetOkExists("connection_info"); ok {
			if tmpList := connectionInfo.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_info", 0)
				tmp, err := s.mapToCloudDbSystemConnectionInfo(fieldKeyFormat)
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

		request.CreateCloudDbSystemConnectorDetails = details
	default:
		return fmt.Errorf("unknown connectorType '%v' was specified", connectorType)
	}
	return nil
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) getCloudDbSystemConnectorFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	cloudDbSystemConnectorId, err := cloudDbSystemConnectorWaitForWorkRequest(workId, "connector",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*cloudDbSystemConnectorId)

	return s.Get()
}

func cloudDbSystemConnectorWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func cloudDbSystemConnectorWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = cloudDbSystemConnectorWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDatabaseManagementCloudDbSystemConnectorWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementCloudDbSystemConnectorWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) Get() error {
	request := oci_database_management.GetCloudDbSystemConnectorRequest{}

	tmp := s.D.Id()
	request.CloudDbSystemConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetCloudDbSystemConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudDbSystemConnector
	return nil
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) Update() error {
	request := oci_database_management.UpdateCloudDbSystemConnectorRequest{}

	tmp := s.D.Id()
	request.CloudDbSystemConnectorId = &tmp

	err := s.populateTopLevelPolymorphicUpdateCloudDbSystemConnectorRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateCloudDbSystemConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getCloudDbSystemConnectorFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) populateTopLevelPolymorphicUpdateCloudDbSystemConnectorRequest(request *oci_database_management.UpdateCloudDbSystemConnectorRequest) error {
	connectorTypeRaw, ok := s.D.GetOkExists("connector_type")
	var connectorType string
	if ok {
		connectorType = connectorTypeRaw.(string)
	} else {
		connectorType = "" // default value
	}
	switch strings.ToLower(connectorType) {
	case strings.ToLower("MACS"):
		details := oci_database_management.UpdateCloudDbSystemMacsConnectorDetails{}

		if connectionInfo, ok := s.D.GetOkExists("connection_info"); ok {
			if tmpList := connectionInfo.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_info", 0)
				tmp, err := s.mapToCloudDbSystemConnectionInfo(fieldKeyFormat)
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

		request.UpdateCloudDbSystemConnectorDetails = details
	default:
		return fmt.Errorf("unknown connectorType '%v' was specified", connectorType)
	}
	return nil
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) Delete() error {
	request := oci_database_management.DeleteCloudDbSystemConnectorRequest{}

	tmp := s.D.Id()
	request.CloudDbSystemConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	_, err := s.Client.DeleteCloudDbSystemConnector(context.Background(), request)
	return err
}

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_database_management.CloudDbSystemMacsConnector:
		s.D.Set("connector_type", "MACS")

		if v.AgentId != nil {
			s.D.Set("agent_id", *v.AgentId)
		}

		if v.ConnectionInfo != nil {
			connectionInfoArray := []interface{}{}
			if connectionInfoMap := CloudDbSystemConnectionInfoToMap(&v.ConnectionInfo); connectionInfoMap != nil {
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

		if v.CloudDbSystemId != nil {
			s.D.Set("cloud_db_system_id", *v.CloudDbSystemId)
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

		/*	if v.Id != nil {
			s.D.Set("id", *v.Id)
		} */

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

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) mapToAsmConnectionString(fieldKeyFormat string) (oci_database_management.AsmConnectionString, error) {
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

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) mapToCloudAsmConnectionCredentials(fieldKeyFormat string) (oci_database_management.CloudAsmConnectionCredentials, error) {
	var baseObject oci_database_management.CloudAsmConnectionCredentials
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
		details := oci_database_management.CloudAsmConnectionCredentialsByDetails{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database_management.CloudAsmConnectionCredentialsByDetailsRoleEnum(role.(string))
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		baseObject = details
	case strings.ToLower("NAME_REFERENCE"):
		details := oci_database_management.CloudAsmConnectionCredentialsByName{}
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

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) mapToCloudDbSystemConnectionInfo(fieldKeyFormat string) (oci_database_management.CloudDbSystemConnectionInfo, error) {
	var baseObject oci_database_management.CloudDbSystemConnectionInfo
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
		details := oci_database_management.CloudAsmConnectionInfo{}
		if connectionCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_credentials")); ok {
			if tmpList := connectionCredentials.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_credentials"), 0)
				tmp, err := s.mapToCloudAsmConnectionCredentials(fieldKeyFormatNextLevel)
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
		details := oci_database_management.CloudDatabaseConnectionInfo{}
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
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown component_type '%v' was specified", componentType)
	}
	return baseObject, nil
}

func CloudDbSystemConnectorSummaryToMap(obj oci_database_management.CloudDbSystemConnectorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
	}

	if obj.CloudDbSystemId != nil {
		result["cloud_db_system_id"] = string(*obj.CloudDbSystemId)
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

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) mapToDatabaseConnectionCredentials(fieldKeyFormat string) (oci_database_management.DatabaseConnectionCredentials, error) {
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

func (s *DatabaseManagementCloudDbSystemConnectorResourceCrud) mapToDatabaseConnectionString(fieldKeyFormat string) (oci_database_management.DatabaseConnectionString, error) {
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
