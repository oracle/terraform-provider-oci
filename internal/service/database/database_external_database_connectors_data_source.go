// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseExternalDatabaseConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseExternalDatabaseConnectors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_database_connectors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseExternalDatabaseConnectorResource()),
			},
		},
	}
}

func readDatabaseExternalDatabaseConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalDatabaseConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExternalDatabaseConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListExternalDatabaseConnectorsResponse
}

func (s *DatabaseExternalDatabaseConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExternalDatabaseConnectorsDataSourceCrud) Get() error {
	request := oci_database.ListExternalDatabaseConnectorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalDatabaseId, ok := s.D.GetOkExists("external_database_id"); ok {
		tmp := externalDatabaseId.(string)
		request.ExternalDatabaseId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ExternalDatabaseConnectorLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListExternalDatabaseConnectors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalDatabaseConnectors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseExternalDatabaseConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExternalDatabaseConnectorsDataSource-", DatabaseExternalDatabaseConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		result := map[string]interface{}{}
		switch v := (r).(type) {
		case oci_database.ExternalMacsConnectorSummary:
			result["connector_type"] = "MACS"

			if v.ConnectionCredentials != nil {
				connectionCredentialsArray := []interface{}{}
				if connectionCredentialsMap := s.DatabaseConnectionCredentialsToMap(&v.ConnectionCredentials); connectionCredentialsMap != nil {
					connectionCredentialsArray = append(connectionCredentialsArray, connectionCredentialsMap)
				}
				result["connection_credentials"] = connectionCredentialsArray
			}

			if v.ConnectionString != nil {
				result["connection_string"] = []interface{}{DatabaseConnectionStringToMap(v.ConnectionString)}
			}

			if v.ConnectorAgentId != nil {
				result["connector_agent_id"] = string(*v.ConnectorAgentId)
			}

			if v.CompartmentId != nil {
				result["compartment_id"] = string(*v.CompartmentId)
			}

			if v.ConnectionStatus != nil {
				result["connection_status"] = string(*v.ConnectionStatus)
			}

			if v.DefinedTags != nil {
				result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
			}

			if v.DisplayName != nil {
				result["display_name"] = string(*v.DisplayName)
			}

			if v.ExternalDatabaseId != nil {
				result["external_database_id"] = string(*v.ExternalDatabaseId)
			}

			result["freeform_tags"] = v.FreeformTags

			if v.Id != nil {
				result["id"] = string(*v.Id)
			}

			if v.LifecycleDetails != nil {
				result["lifecycle_details"] = string(*v.LifecycleDetails)
			}

			result["state"] = string(v.LifecycleState)

			if v.TimeConnectionStatusLastUpdated != nil {
				result["time_connection_status_last_updated"] = v.TimeConnectionStatusLastUpdated.String()
			}

			if v.TimeCreated != nil {
				result["time_created"] = v.TimeCreated.String()
			}
		default:
			log.Printf("[WARN] Received 'connector_type' of unknown type %v", r)
			return nil
		}

		resources = append(resources, result)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseExternalDatabaseConnectorsDataSource().Schema["external_database_connectors"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("external_database_connectors", resources); err != nil {
		return err
	}

	return nil
}

func (s *DatabaseExternalDatabaseConnectorsDataSourceCrud) DatabaseConnectionCredentialsToMap(obj *oci_database.DatabaseConnectionCredentials) map[string]interface{} {
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
	case oci_database.DatabaseConnectionCredentialsByName:
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
