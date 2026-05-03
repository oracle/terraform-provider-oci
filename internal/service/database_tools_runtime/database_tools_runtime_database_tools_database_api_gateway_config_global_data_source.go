// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools_runtime

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools_runtime "github.com/oracle/oci-go-sdk/v65/databasetoolsruntime"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_tools_database_api_gateway_config_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["global_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResource(), fieldMap, readSingularDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalWithContext)
}

func readSingularDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigGlobalResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigGlobalRequest{}

	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}

	if globalKey, ok := s.D.GetOkExists("global_key"); ok {
		request.GlobalKey = oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum(globalKey.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.GetDatabaseToolsDatabaseApiGatewayConfigGlobal(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalDataSource-", DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalDataSource(), s.D))

	if key := s.Res.DatabaseToolsDatabaseApiGatewayConfigGlobal.GetKey(); key != nil {
		s.D.Set("key", *key)
	}

	s.D.Set("metadata_source", s.Res.DatabaseToolsDatabaseApiGatewayConfigGlobal.GetMetadataSource())

	if timeCreated := s.Res.DatabaseToolsDatabaseApiGatewayConfigGlobal.GetTimeCreated(); timeCreated != nil {
		s.D.Set("time_created", timeCreated.String())
	}

	if timeUpdated := s.Res.DatabaseToolsDatabaseApiGatewayConfigGlobal.GetTimeUpdated(); timeUpdated != nil {
		s.D.Set("time_updated", timeUpdated.String())
	}

	switch v := (s.Res.DatabaseToolsDatabaseApiGatewayConfigGlobal).(type) {
	case oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigGlobalDefault:
		s.D.Set("type", "DEFAULT")

		s.D.Set("advanced_properties", v.AdvancedProperties)

		if v.CertificateBundle != nil {
			certificateBundleArray := []interface{}{}
			if certificateBundleMap := DatabaseApiGatewayConfigCertificateBundleToMap(&v.CertificateBundle); certificateBundleMap != nil {
				certificateBundleArray = append(certificateBundleArray, certificateBundleMap)
			}
			s.D.Set("certificate_bundle", certificateBundleArray)
		} else {
			s.D.Set("certificate_bundle", nil)
		}

		s.D.Set("database_api_status", v.DatabaseApiStatus)

		if v.DocumentRoot != nil {
			s.D.Set("document_root", *v.DocumentRoot)
		}

		if v.HttpPort != nil {
			s.D.Set("http_port", *v.HttpPort)
		}

		if v.HttpsPort != nil {
			s.D.Set("https_port", *v.HttpsPort)
		}

		s.D.Set("pool_route", v.PoolRoute)

		if v.PoolRoutingHeader != nil {
			s.D.Set("pool_routing_header", *v.PoolRoutingHeader)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.DatabaseToolsDatabaseApiGatewayConfigGlobal)
		return nil
	}

	return nil
}
