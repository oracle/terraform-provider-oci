// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalExadataStorageConnectorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalExadataStorageConnector,
		Read:     readDatabaseManagementExternalExadataStorageConnector,
		Update:   updateDatabaseManagementExternalExadataStorageConnector,
		Delete:   deleteDatabaseManagementExternalExadataStorageConnector,
		Schema: map[string]*schema.Schema{
			// Required
			"agent_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"connection_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"connector_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"credential_info": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
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
						"ssl_trust_store_location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_trust_store_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							Sensitive: true,
						},
						"ssl_trust_store_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"storage_server_id": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"additional_details": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"internal_id": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseManagementExternalExadataStorageConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalExadataStorageConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementExternalExadataStorageConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalExadataStorageConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementExternalExadataStorageConnectorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ExternalExadataStorageConnector
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalExadataStorageConnectorResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementExternalExadataStorageConnectorResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.LifecycleStatesCreating),
	}
}

func (s *DatabaseManagementExternalExadataStorageConnectorResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.LifecycleStatesActive),
	}
}

func (s *DatabaseManagementExternalExadataStorageConnectorResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.LifecycleStatesDeleting),
	}
}

func (s *DatabaseManagementExternalExadataStorageConnectorResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.LifecycleStatesDeleted),
	}
}

func (s *DatabaseManagementExternalExadataStorageConnectorResourceCrud) Create() error {
	request := oci_database_management.CreateExternalExadataStorageConnectorRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if connectionUri, ok := s.D.GetOkExists("connection_uri"); ok {
		tmp := connectionUri.(string)
		request.ConnectionUri = &tmp
	}

	if connectorName, ok := s.D.GetOkExists("connector_name"); ok {
		tmp := connectorName.(string)
		request.ConnectorName = &tmp
	}

	if credentialInfo, ok := s.D.GetOkExists("credential_info"); ok {
		if tmpList := credentialInfo.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credential_info", 0)
			tmp, err := s.mapToRestCredential(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CredentialInfo = &tmp
		}
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

	if storageServerId, ok := s.D.GetOkExists("storage_server_id"); ok {
		tmp := storageServerId.(string)
		request.StorageServerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.CreateExternalExadataStorageConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalExadataStorageConnector
	return nil
}

func (s *DatabaseManagementExternalExadataStorageConnectorResourceCrud) Get() error {
	request := oci_database_management.GetExternalExadataStorageConnectorRequest{}

	tmp := s.D.Id()
	request.ExternalExadataStorageConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetExternalExadataStorageConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalExadataStorageConnector
	return nil
}

func (s *DatabaseManagementExternalExadataStorageConnectorResourceCrud) Update() error {
	request := oci_database_management.UpdateExternalExadataStorageConnectorRequest{}

	if connectionUri, ok := s.D.GetOkExists("connection_uri"); ok {
		tmp := connectionUri.(string)
		request.ConnectionUri = &tmp
	}

	if connectorName, ok := s.D.GetOkExists("connector_name"); ok {
		tmp := connectorName.(string)
		request.ConnectorName = &tmp
	}

	if credentialInfo, ok := s.D.GetOkExists("credential_info"); ok {
		if tmpList := credentialInfo.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credential_info", 0)
			tmp, err := s.mapToRestCredential(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CredentialInfo = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.ExternalExadataStorageConnectorId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalExadataStorageConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalExadataStorageConnector
	return nil
}

func (s *DatabaseManagementExternalExadataStorageConnectorResourceCrud) Delete() error {
	request := oci_database_management.DeleteExternalExadataStorageConnectorRequest{}

	tmp := s.D.Id()
	request.ExternalExadataStorageConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	_, err := s.Client.DeleteExternalExadataStorageConnector(context.Background(), request)
	return err
}

func (s *DatabaseManagementExternalExadataStorageConnectorResourceCrud) SetData() error {
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.AgentId != nil {
		s.D.Set("agent_id", *s.Res.AgentId)
	}

	if s.Res.ConnectionUri != nil {
		s.D.Set("connection_uri", *s.Res.ConnectionUri)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExadataInfrastructureId != nil {
		s.D.Set("exadata_infrastructure_id", *s.Res.ExadataInfrastructureId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InternalId != nil {
		s.D.Set("internal_id", *s.Res.InternalId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Status != nil {
		s.D.Set("status", *s.Res.Status)
	}

	if s.Res.StorageServerId != nil {
		s.D.Set("storage_server_id", *s.Res.StorageServerId)
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

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func ExternalExadataStorageConnectorSummaryToMap(obj oci_database_management.ExternalExadataStorageConnectorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["additional_details"] = obj.AdditionalDetails

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
	}

	if obj.ConnectionUri != nil {
		result["connection_uri"] = string(*obj.ConnectionUri)
	}

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

	if obj.InternalId != nil {
		result["internal_id"] = string(*obj.InternalId)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.StorageServerId != nil {
		result["storage_server_id"] = string(*obj.StorageServerId)
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

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func (s *DatabaseManagementExternalExadataStorageConnectorResourceCrud) mapToRestCredential(fieldKeyFormat string) (oci_database_management.RestCredential, error) {
	result := oci_database_management.RestCredential{}

	if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
		tmp := password.(string)
		result.Password = &tmp
	}

	if sslTrustStoreLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_trust_store_location")); ok {
		tmp := sslTrustStoreLocation.(string)
		result.SslTrustStoreLocation = &tmp
	}

	if sslTrustStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_trust_store_password")); ok {
		tmp := sslTrustStorePassword.(string)
		result.SslTrustStorePassword = &tmp
	}

	if sslTrustStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_trust_store_type")); ok {
		result.SslTrustStoreType = oci_database_management.RestCredentialSslTrustStoreTypeEnum(sslTrustStoreType.(string))
	}

	if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
		tmp := username.(string)
		result.Username = &tmp
	}

	return result, nil
}

func RestCredentialToMap(obj *oci_database_management.RestCredential) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Password != nil {
		result["password"] = string(*obj.Password)
	}

	if obj.SslTrustStoreLocation != nil {
		result["ssl_trust_store_location"] = string(*obj.SslTrustStoreLocation)
	}

	if obj.SslTrustStorePassword != nil {
		result["ssl_trust_store_password"] = string(*obj.SslTrustStorePassword)
	}

	result["ssl_trust_store_type"] = string(obj.SslTrustStoreType)

	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}

	return result
}
