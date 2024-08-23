// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementManagedDatabase,
		Read:     readDatabaseManagementManagedDatabase,
		Update:   updateDatabaseManagementManagedDatabase,
		Delete:   deleteDatabaseManagementManagedDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_database_id": {
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
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_platform_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_sub_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dbmgmt_feature_configs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"connector_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"connector_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_connector_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"management_agent_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_end_point_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"database_connection_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"connection_credentials": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"credential_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"credential_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"named_credential_id": {
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
												"ssl_secret_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"user_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"connection_string": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"connection_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"port": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"protocol": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"service": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"feature": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"feature_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license_model": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"deployment_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_cluster": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"managed_database_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"management_option": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_container_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_system_id": {
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
			"workload_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseManagementManagedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementManagedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementManagedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementManagedDatabase(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementManagedDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ManagedDatabase
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementManagedDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementManagedDatabaseResourceCrud) Create() error {
	request := oci_database_management.UpdateManagedDatabaseRequest{}

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

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateManagedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedDatabase
	return nil
}

func (s *DatabaseManagementManagedDatabaseResourceCrud) Get() error {
	request := oci_database_management.GetManagedDatabaseRequest{}

	tmp := s.D.Id()
	request.ManagedDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetManagedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedDatabase
	return nil
}

func (s *DatabaseManagementManagedDatabaseResourceCrud) Update() error {
	request := oci_database_management.UpdateManagedDatabaseRequest{}

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

	tmp := s.D.Id()
	request.ManagedDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateManagedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedDatabase
	return nil
}

func (s *DatabaseManagementManagedDatabaseResourceCrud) SetData() error {
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabasePlatformName != nil {
		s.D.Set("database_platform_name", *s.Res.DatabasePlatformName)
	}

	s.D.Set("database_status", s.Res.DatabaseStatus)

	s.D.Set("database_sub_type", s.Res.DatabaseSubType)

	s.D.Set("database_type", s.Res.DatabaseType)

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	dbmgmtFeatureConfigs := []interface{}{}
	for _, item := range s.Res.DbmgmtFeatureConfigs {
		dbmgmtFeatureConfigs = append(dbmgmtFeatureConfigs, DatabaseFeatureConfigurationToMap(item))
	}
	s.D.Set("dbmgmt_feature_configs", dbmgmtFeatureConfigs)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("deployment_type", s.Res.DeploymentType)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCluster != nil {
		s.D.Set("is_cluster", *s.Res.IsCluster)
	}

	managedDatabaseGroups := []interface{}{}
	for _, item := range s.Res.ManagedDatabaseGroups {
		managedDatabaseGroups = append(managedDatabaseGroups, ParentGroupToMap(item))
	}
	s.D.Set("managed_database_groups", managedDatabaseGroups)

	s.D.Set("management_option", s.Res.ManagementOption)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentContainerId != nil {
		s.D.Set("parent_container_id", *s.Res.ParentContainerId)
	}

	if s.Res.StorageSystemId != nil {
		s.D.Set("storage_system_id", *s.Res.StorageSystemId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("workload_type", s.Res.WorkloadType)

	return nil
}

func ConnectorDetailsToMap(obj *oci_database_management.ConnectorDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_management.ExternalConnectorDetails:
		result["connector_type"] = "EXTERNAL"

		if v.DatabaseConnectorId != nil {
			result["database_connector_id"] = string(*v.DatabaseConnectorId)
		}
	case oci_database_management.MacsConnectorDetails:
		result["connector_type"] = "MACS"

		if v.ManagementAgentId != nil {
			result["management_agent_id"] = string(*v.ManagementAgentId)
		}
	case oci_database_management.PrivateEndPointConnectorDetails:
		result["connector_type"] = "PE"

		if v.PrivateEndPointId != nil {
			result["private_end_point_id"] = string(*v.PrivateEndPointId)
		}
	default:
		log.Printf("[WARN] Received 'connector_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DatabaseConnectionCredentialsToMap(obj *oci_database_management.DatabaseConnectionCredentials) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_management.DatabaseConnectionCredentialsByDetails:
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
	case oci_database_management.DatabaseNamedCredentialConnectionDetails:
		result["credential_type"] = "NAMED_CREDENTIAL"

		if v.NamedCredentialId != nil {
			result["named_credential_id"] = string(*v.NamedCredentialId)
		}
	case oci_database_management.DatabaseConnectionCredentailsByName:
		result["credential_type"] = "NAME_REFERENCE"

		if v.CredentialName != nil {
			result["credential_name"] = string(*v.CredentialName)
		}
	case oci_database_management.DatabaseSslConnectionCredentials:
		result["credential_type"] = "SSL_DETAILS"

		if v.CredentialName != nil {
			result["credential_name"] = string(*v.CredentialName)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		result["role"] = string(v.Role)

		if v.SslSecretId != nil {
			result["ssl_secret_id"] = string(*v.SslSecretId)
		}

		if v.UserName != nil {
			result["user_name"] = string(*v.UserName)
		}
	default:
		log.Printf("[WARN] Received 'credential_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DatabaseConnectionDetailsToMap(obj *oci_database_management.DatabaseConnectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectionCredentials != nil {
		connectionCredentialsArray := []interface{}{}
		if connectionCredentialsMap := DatabaseConnectionCredentialsToMap(&obj.ConnectionCredentials); connectionCredentialsMap != nil {
			connectionCredentialsArray = append(connectionCredentialsArray, connectionCredentialsMap)
		}
		result["connection_credentials"] = connectionCredentialsArray
	}

	if obj.ConnectionString != nil {
		connectionStringArray := []interface{}{}
		if connectionStringMap := DatabaseConnectionStringDetailsToMap(&obj.ConnectionString); connectionStringMap != nil {
			connectionStringArray = append(connectionStringArray, connectionStringMap)
		}
		result["connection_string"] = connectionStringArray
	}

	return result
}

func DatabaseConnectionStringDetailsToMap(obj *oci_database_management.DatabaseConnectionStringDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_management.BasicDatabaseConnectionStringDetails:
		result["connection_type"] = "BASIC"

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		result["protocol"] = string(v.Protocol)

		if v.Service != nil {
			result["service"] = string(*v.Service)
		}
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DatabaseFeatureConfigurationToMap(obj oci_database_management.DatabaseFeatureConfiguration) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_management.DatabaseLifecycleFeatureConfiguration:
		result["feature"] = "DB_LIFECYCLE_MANAGEMENT"

		result["license_model"] = string(v.LicenseModel)

		if v.ConnectorDetails != nil {
			connectorDetailsArray := []interface{}{}
			if connectorDetailsMap := ConnectorDetailsToMap(&v.ConnectorDetails); connectorDetailsMap != nil {
				connectorDetailsArray = append(connectorDetailsArray, connectorDetailsMap)
			}
			result["connector_details"] = connectorDetailsArray
		}

		if v.DatabaseConnectionDetails != nil {
			result["database_connection_details"] = []interface{}{DatabaseConnectionDetailsToMap(v.DatabaseConnectionDetails)}
		}

		result["feature_status"] = string(v.FeatureStatus)
	case oci_database_management.DatabaseDiagnosticsAndManagementFeatureConfiguration:
		result["feature"] = "DIAGNOSTICS_AND_MANAGEMENT"

		result["license_model"] = string(v.LicenseModel)

		if v.ConnectorDetails != nil {
			connectorDetailsArray := []interface{}{}
			if connectorDetailsMap := ConnectorDetailsToMap(&v.ConnectorDetails); connectorDetailsMap != nil {
				connectorDetailsArray = append(connectorDetailsArray, connectorDetailsMap)
			}
			result["connector_details"] = connectorDetailsArray
		}

		if v.DatabaseConnectionDetails != nil {
			result["database_connection_details"] = []interface{}{DatabaseConnectionDetailsToMap(v.DatabaseConnectionDetails)}
		}

		result["feature_status"] = string(v.FeatureStatus)
	case oci_database_management.DatabaseSqlWatchFeatureConfiguration:
		result["feature"] = "SQLWATCH"

		if v.ConnectorDetails != nil {
			connectorDetailsArray := []interface{}{}
			if connectorDetailsMap := ConnectorDetailsToMap(&v.ConnectorDetails); connectorDetailsMap != nil {
				connectorDetailsArray = append(connectorDetailsArray, connectorDetailsMap)
			}
			result["connector_details"] = connectorDetailsArray
		}

		if v.DatabaseConnectionDetails != nil {
			result["database_connection_details"] = []interface{}{DatabaseConnectionDetailsToMap(v.DatabaseConnectionDetails)}
		}

		result["feature_status"] = string(v.FeatureStatus)
	default:
		log.Printf("[WARN] Received 'feature' of unknown type %v", obj)
		return nil
	}

	return result
}

func ManagedDatabaseSummaryToMap(obj oci_database_management.ManagedDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatabasePlatformName != nil {
		result["database_platform_name"] = string(*obj.DatabasePlatformName)
	}

	result["database_sub_type"] = string(obj.DatabaseSubType)

	result["database_type"] = string(obj.DatabaseType)

	if obj.DatabaseVersion != nil {
		result["database_version"] = string(*obj.DatabaseVersion)
	}

	if obj.DbSystemId != nil {
		result["db_system_id"] = string(*obj.DbSystemId)
	}

	dbmgmtFeatureConfigs := []interface{}{}
	for _, item := range obj.DbmgmtFeatureConfigs {
		dbmgmtFeatureConfigs = append(dbmgmtFeatureConfigs, DatabaseFeatureConfigurationToMap(item))
	}
	result["dbmgmt_feature_configs"] = dbmgmtFeatureConfigs

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["deployment_type"] = string(obj.DeploymentType)

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsCluster != nil {
		result["is_cluster"] = bool(*obj.IsCluster)
	}

	result["management_option"] = string(obj.ManagementOption)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentContainerId != nil {
		result["parent_container_id"] = string(*obj.ParentContainerId)
	}

	if obj.StorageSystemId != nil {
		result["storage_system_id"] = string(*obj.StorageSystemId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	result["workload_type"] = string(obj.WorkloadType)

	return result
}

func ParentGroupToMap(obj oci_database_management.ParentGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
