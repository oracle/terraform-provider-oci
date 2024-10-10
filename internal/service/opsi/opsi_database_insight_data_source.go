// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"
)

func OpsiDatabaseInsightDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_insight_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpsiDatabaseInsightResource(), fieldMap, readSingularOpsiDatabaseInsight)
}

func readSingularOpsiDatabaseInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiDatabaseInsightDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiDatabaseInsightDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.GetDatabaseInsightResponse
}

func (s *OpsiDatabaseInsightDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiDatabaseInsightDataSourceCrud) Get() error {
	request := oci_opsi.GetDatabaseInsightRequest{}

	if databaseInsightId, ok := s.D.GetOkExists("database_insight_id"); ok {
		tmp := databaseInsightId.(string)
		request.DatabaseInsightId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetDatabaseInsight(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiDatabaseInsightDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiDatabaseInsightsSingularDataSource-", OpsiDatabaseInsightsDataSource(), s.D))
	switch v := (s.Res.DatabaseInsight).(type) {
	case oci_opsi.AutonomousDatabaseInsight:
		s.D.Set("entity_source", "AUTONOMOUS_DATABASE")

		if v.ConnectionDetails != nil {
			s.D.Set("connection_details", []interface{}{ConnectionDetailsToMap(v.ConnectionDetails)})
		} else {
			s.D.Set("connection_details", nil)
		}
		log.Printf("[DEBUG] in data source setData")
		if v.CredentialDetails != nil {
			credentialDetailsArray := []interface{}{}
			if credentialDetailsMap := CredentialDetailsToMap(&v.CredentialDetails); credentialDetailsMap != nil {
				credentialDetailsArray = append(credentialDetailsArray, credentialDetailsMap)
			}
			s.D.Set("credential_details", credentialDetailsArray)
		} else {
			s.D.Set("credential_details", nil)
		}

		if v.DatabaseDisplayName != nil {
			s.D.Set("database_display_name", *v.DatabaseDisplayName)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DatabaseResourceType != nil {
			s.D.Set("database_resource_type", *v.DatabaseResourceType)
		}

		if v.OpsiPrivateEndpointId != nil {
			s.D.Set("opsi_private_endpoint_id", *v.OpsiPrivateEndpointId)
		}

		if v.IsAdvancedFeaturesEnabled != nil {
			s.D.Set("is_advanced_features_enabled", *v.IsAdvancedFeaturesEnabled)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseConnectionStatusDetails != nil {
			s.D.Set("database_connection_status_details", *v.DatabaseConnectionStatusDetails)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			s.D.Set("database_version", *v.DatabaseVersion)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_opsi.EmManagedExternalDatabaseInsight:
		s.D.Set("entity_source", "EM_MANAGED_EXTERNAL_DATABASE")

		if v.EnterpriseManagerBridgeId != nil {
			s.D.Set("enterprise_manager_bridge_id", *v.EnterpriseManagerBridgeId)
		}

		if v.EnterpriseManagerEntityDisplayName != nil {
			s.D.Set("enterprise_manager_entity_display_name", *v.EnterpriseManagerEntityDisplayName)
		}

		if v.EnterpriseManagerEntityIdentifier != nil {
			s.D.Set("enterprise_manager_entity_identifier", *v.EnterpriseManagerEntityIdentifier)
		}

		if v.EnterpriseManagerEntityName != nil {
			s.D.Set("enterprise_manager_entity_name", *v.EnterpriseManagerEntityName)
		}

		if v.EnterpriseManagerEntityType != nil {
			s.D.Set("enterprise_manager_entity_type", *v.EnterpriseManagerEntityType)
		}

		if v.EnterpriseManagerIdentifier != nil {
			s.D.Set("enterprise_manager_identifier", *v.EnterpriseManagerIdentifier)
		}

		if v.ExadataInsightId != nil {
			s.D.Set("exadata_insight_id", *v.ExadataInsightId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			s.D.Set("database_version", *v.DatabaseVersion)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_opsi.MacsManagedCloudDatabaseInsight:
		s.D.Set("entity_source", "MACS_MANAGED_CLOUD_DATABASE")

		if v.ConnectionCredentialDetails != nil {
			connectionCredentialDetailsArray := []interface{}{}
			if connectionCredentialDetailsMap := CredentialDetailsToMap(&v.ConnectionCredentialDetails); connectionCredentialDetailsMap != nil {
				connectionCredentialDetailsArray = append(connectionCredentialDetailsArray, connectionCredentialDetailsMap)
			}
			s.D.Set("connection_credential_details", connectionCredentialDetailsArray)
		} else {
			s.D.Set("connection_credential_details", nil)
		}

		if v.ConnectionDetails != nil {
			s.D.Set("connection_details", []interface{}{ConnectionDetailsToMap(v.ConnectionDetails)})
		} else {
			s.D.Set("connection_details", nil)
		}

		if v.DatabaseDisplayName != nil {
			s.D.Set("database_display_name", *v.DatabaseDisplayName)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DatabaseResourceType != nil {
			s.D.Set("database_resource_type", *v.DatabaseResourceType)
		}

		if v.ManagementAgentId != nil {
			s.D.Set("management_agent_id", *v.ManagementAgentId)
		}

		if v.ParentId != nil {
			s.D.Set("parent_id", *v.ParentId)
		}

		if v.RootId != nil {
			s.D.Set("root_id", *v.RootId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseConnectionStatusDetails != nil {
			s.D.Set("database_connection_status_details", *v.DatabaseConnectionStatusDetails)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			s.D.Set("database_version", *v.DatabaseVersion)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_opsi.MdsMySqlDatabaseInsight:
		s.D.Set("entity_source", "MDS_MYSQL_DATABASE_SYSTEM")

		if v.DatabaseDisplayName != nil {
			s.D.Set("database_display_name", *v.DatabaseDisplayName)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DatabaseResourceType != nil {
			s.D.Set("database_resource_type", *v.DatabaseResourceType)
		}

		if v.IsHeatWaveClusterAttached != nil {
			s.D.Set("is_heat_wave_cluster_attached", *v.IsHeatWaveClusterAttached)
		}

		if v.IsHighlyAvailable != nil {
			s.D.Set("is_highly_available", *v.IsHighlyAvailable)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseConnectionStatusDetails != nil {
			s.D.Set("database_connection_status_details", *v.DatabaseConnectionStatusDetails)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_opsi.PeComanagedDatabaseInsight:
		s.D.Set("entity_source", "PE_COMANAGED_DATABASE")

		if v.CredentialDetails != nil {
			credentialDetailsArray := []interface{}{}
			if credentialDetailsMap := CredentialDetailsToMap(&v.CredentialDetails); credentialDetailsMap != nil {
				credentialDetailsArray = append(credentialDetailsArray, credentialDetailsMap)
			}
			s.D.Set("credential_details", credentialDetailsArray)
		} else {
			s.D.Set("credential_details", nil)
		}

		if v.DatabaseDisplayName != nil {
			s.D.Set("database_display_name", *v.DatabaseDisplayName)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DatabaseResourceType != nil {
			s.D.Set("database_resource_type", *v.DatabaseResourceType)
		}

		if v.OpsiPrivateEndpointId != nil {
			s.D.Set("opsi_private_endpoint_id", *v.OpsiPrivateEndpointId)
		}

		if v.ParentId != nil {
			s.D.Set("parent_id", *v.ParentId)
		}

		if v.RootId != nil {
			s.D.Set("root_id", *v.RootId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseConnectionStatusDetails != nil {
			s.D.Set("database_connection_status_details", *v.DatabaseConnectionStatusDetails)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.DatabaseType != nil {
			s.D.Set("database_type", *v.DatabaseType)
		}

		if v.DatabaseVersion != nil {
			s.D.Set("database_version", *v.DatabaseVersion)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.OpsiPrivateEndpointId != nil {
			s.D.Set("opsi_private_endpoint_id", *v.OpsiPrivateEndpointId)
		}

		if v.ProcessorCount != nil {
			s.D.Set("processor_count", *v.ProcessorCount)
		}

		s.D.Set("state", v.LifecycleState)

		s.D.Set("status", v.Status)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", s.Res.DatabaseInsight)
		return nil
	}

	return nil
}
