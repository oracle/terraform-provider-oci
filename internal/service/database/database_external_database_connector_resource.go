// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabaseExternalDatabaseConnectorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExternalDatabaseConnector,
		Read:     readDatabaseExternalDatabaseConnector,
		Update:   updateDatabaseExternalDatabaseConnector,
		Delete:   deleteDatabaseExternalDatabaseConnector,
		Schema: map[string]*schema.Schema{
			// Required
			"connection_credentials": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"credential_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						// Optional
						"credential_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DETAILS",
								"NAME_REFERENCE",
								"SSL_DETAILS",
							}, true),
						},
						"password": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:  true,
							Sensitive: true,
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
						"username": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed: true,
						},

						// Computed
					},
				},
			},
			"connection_string": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"hostname": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"connector_agent_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"connector_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"MACS",
				}, true),
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
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_connection_status_last_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseExternalDatabaseConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalDatabaseConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseExternalDatabaseConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalDatabaseConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func updateDatabaseExternalDatabaseConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalDatabaseConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseExternalDatabaseConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalDatabaseConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExternalDatabaseConnectorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ExternalDatabaseConnector
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) ID() string {
	externalDatabaseConnector := *s.Res
	return *externalDatabaseConnector.GetId()
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExternalDatabaseConnectorLifecycleStateProvisioning),
	}
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExternalDatabaseConnectorLifecycleStateAvailable),
	}
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExternalDatabaseConnectorLifecycleStateTerminating),
	}
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExternalDatabaseConnectorLifecycleStateTerminated),
	}
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) Create() error {
	request := oci_database.CreateExternalDatabaseConnectorRequest{}
	err := s.populateTopLevelPolymorphicCreateExternalDatabaseConnectorRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateExternalDatabaseConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.ExternalDatabaseConnector

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.GetId()
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalDatabaseConnector", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) Get() error {
	request := oci_database.GetExternalDatabaseConnectorRequest{}

	tmp := s.D.Id()
	request.ExternalDatabaseConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExternalDatabaseConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalDatabaseConnector
	return nil
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) Update() error {
	request := oci_database.UpdateExternalDatabaseConnectorRequest{}
	err := s.populateTopLevelPolymorphicUpdateExternalDatabaseConnectorRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateExternalDatabaseConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalDatabaseConnector", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) Delete() error {
	request := oci_database.DeleteExternalDatabaseConnectorRequest{}

	tmp := s.D.Id()
	request.ExternalDatabaseConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteExternalDatabaseConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalDatabaseConnector", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_database.ExternalMacsConnector:
		s.D.Set("connector_type", "MACS")

		if v.ConnectionCredentials != nil {
			connectionCredentialsArray := []interface{}{}
			if connectionCredentialsMap := s.DatabaseConnectionCredentialsToMap(&v.ConnectionCredentials); connectionCredentialsMap != nil {
				connectionCredentialsArray = append(connectionCredentialsArray, connectionCredentialsMap)
			}
			s.D.Set("connection_credentials", connectionCredentialsArray)
		} else {
			s.D.Set("connection_credentials", nil)
		}

		if v.ConnectionString != nil {
			s.D.Set("connection_string", []interface{}{DatabaseConnectionStringToMap(v.ConnectionString)})
		} else {
			s.D.Set("connection_string", nil)
		}

		if v.ConnectorAgentId != nil {
			s.D.Set("connector_agent_id", *v.ConnectorAgentId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.ConnectionStatus != nil {
			s.D.Set("connection_status", *v.ConnectionStatus)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.ExternalDatabaseId != nil {
			s.D.Set("external_database_id", *v.ExternalDatabaseId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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
	default:
		log.Printf("[WARN] Received 'connector_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) mapToDatabaseConnectionCredentials(fieldKeyFormat string) (oci_database.DatabaseConnectionCredentials, error) {
	var baseObject oci_database.DatabaseConnectionCredentials
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
		details := oci_database.DatabaseConnectionCredentialsByDetails{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database.DatabaseConnectionCredentialsByDetailsRoleEnum(role.(string))
		}
		if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		baseObject = details
	case strings.ToLower("NAME_REFERENCE"):
		details := oci_database.DatabaseConnectionCredentailsByName{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		baseObject = details
	case strings.ToLower("SSL_DETAILS"):
		details := oci_database.DatabaseSslConnectionCredentials{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database.DatabaseSslConnectionCredentialsRoleEnum(role.(string))
		}
		if sslSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_secret_id")); ok {
			tmp := sslSecretId.(string)
			details.SslSecretId = &tmp
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

func (s *DatabaseExternalDatabaseConnectorResourceCrud) DatabaseConnectionCredentialsToMap(obj *oci_database.DatabaseConnectionCredentials) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database.DatabaseConnectionCredentialsByDetails:
		result["credential_type"] = "DETAILS"

		if v.CredentialName != nil {
			result["credential_name"] = string(*v.CredentialName)
		}

		if password, ok := s.D.GetOkExists("connection_credentials.0.password"); ok && password != nil {
			result["password"] = password.(string)
		}

		result["role"] = string(v.Role)

		if username, ok := s.D.GetOkExists("connection_credentials.0.username"); ok && username != nil {
			result["username"] = username.(string)
		}
	case oci_database.DatabaseConnectionCredentailsByName:
		result["credential_type"] = "NAME_REFERENCE"

		if v.CredentialName != nil {
			result["credential_name"] = string(*v.CredentialName)
		}
	case oci_database.DatabaseSslConnectionCredentials:
		result["credential_type"] = "SSL_DETAILS"

		if v.CredentialName != nil {
			result["credential_name"] = string(*v.CredentialName)
		}

		if password, ok := s.D.GetOkExists("connection_credentials.0.password"); ok && password != nil {
			result["password"] = password.(string)
		}

		result["role"] = string(v.Role)

		if v.SslSecretId != nil {
			result["ssl_secret_id"] = string(*v.SslSecretId)
		}

		if username, ok := s.D.GetOkExists("connection_credentials.0.username"); ok && username != nil {
			result["username"] = username.(string)
		}

	default:
		log.Printf("[WARN] Received 'credential_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) mapToDatabaseConnectionString(fieldKeyFormat string) (oci_database.DatabaseConnectionString, error) {
	result := oci_database.DatabaseConnectionString{}

	if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
		tmp := hostname.(string)
		result.Hostname = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_database.DatabaseConnectionStringProtocolEnum(protocol.(string))
	}

	if service, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service")); ok {
		tmp := service.(string)
		result.Service = &tmp
	}

	return result, nil
}

func DatabaseConnectionStringToMap(obj *oci_database.DatabaseConnectionString) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	result["protocol"] = string(obj.Protocol)

	if obj.Service != nil {
		result["service"] = string(*obj.Service)
	}

	return result
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) populateTopLevelPolymorphicCreateExternalDatabaseConnectorRequest(request *oci_database.CreateExternalDatabaseConnectorRequest) error {
	//discriminator
	connectorTypeRaw, ok := s.D.GetOkExists("connector_type")
	var connectorType string
	if ok {
		connectorType = connectorTypeRaw.(string)
	} else {
		connectorType = "MACS" // default value
	}
	switch strings.ToLower(connectorType) {
	case strings.ToLower("MACS"):
		details := oci_database.CreateExternalMacsConnectorDetails{}
		if connectionCredentials, ok := s.D.GetOkExists("connection_credentials"); ok {
			if tmpList := connectionCredentials.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_credentials", 0)
				tmp, err := s.mapToDatabaseConnectionCredentials(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionCredentials = tmp
			}
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			if tmpList := connectionString.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_string", 0)
				tmp, err := s.mapToDatabaseConnectionString(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionString = &tmp
			}
		}
		if connectorAgentId, ok := s.D.GetOkExists("connector_agent_id"); ok {
			tmp := connectorAgentId.(string)
			details.ConnectorAgentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if externalDatabaseId, ok := s.D.GetOkExists("external_database_id"); ok {
			tmp := externalDatabaseId.(string)
			details.ExternalDatabaseId = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateExternalDatabaseConnectorDetails = details
	default:
		return fmt.Errorf("unknown connector_type '%v' was specified", connectorType)
	}
	return nil
}

func (s *DatabaseExternalDatabaseConnectorResourceCrud) populateTopLevelPolymorphicUpdateExternalDatabaseConnectorRequest(request *oci_database.UpdateExternalDatabaseConnectorRequest) error {
	//discriminator
	connectorTypeRaw, ok := s.D.GetOkExists("connector_type")
	var connectorType string
	if ok {
		connectorType = connectorTypeRaw.(string)
	} else {
		connectorType = "MACS" // default value
	}
	switch strings.ToLower(connectorType) {
	case strings.ToLower("MACS"):
		details := oci_database.UpdateExternalMacsConnectorDetails{}
		if connectionCredentials, ok := s.D.GetOkExists("connection_credentials"); ok {
			if tmpList := connectionCredentials.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_credentials", 0)
				tmp, err := s.mapToDatabaseConnectionCredentials(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionCredentials = tmp
			}
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			if tmpList := connectionString.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_string", 0)
				tmp, err := s.mapToDatabaseConnectionString(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ConnectionString = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		tmp := s.D.Id()
		request.ExternalDatabaseConnectorId = &tmp
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateExternalDatabaseConnectorDetails = details
	default:
		return fmt.Errorf("unknown connector_type '%v' was specified", connectorType)
	}
	return nil
}
