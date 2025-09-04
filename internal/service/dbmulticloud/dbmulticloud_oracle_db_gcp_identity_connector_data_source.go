// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudOracleDbGcpIdentityConnectorDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oracle_db_gcp_identity_connector_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DbmulticloudOracleDbGcpIdentityConnectorResource(), fieldMap, readSingularDbmulticloudOracleDbGcpIdentityConnector)
}

func readSingularDbmulticloudOracleDbGcpIdentityConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbGcpIdentityConnectorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudGCPProviderClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbGcpIdentityConnectorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.DbMulticloudGCPProviderClient
	Res    *oci_dbmulticloud.GetOracleDbGcpIdentityConnectorResponse
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorDataSourceCrud) Get() error {
	request := oci_dbmulticloud.GetOracleDbGcpIdentityConnectorRequest{}

	if oracleDbGcpIdentityConnectorId, ok := s.D.GetOkExists("oracle_db_gcp_identity_connector_id"); ok {
		tmp := oracleDbGcpIdentityConnectorId.(string)
		request.OracleDbGcpIdentityConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.GetOracleDbGcpIdentityConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("gcp_identity_connectivity_status", s.Res.GcpIdentityConnectivityStatus)

	if s.Res.GcpLocation != nil {
		s.D.Set("gcp_location", *s.Res.GcpLocation)
	}

	gcpNodes := []interface{}{}
	for _, item := range s.Res.GcpNodes {
		gcpNodes = append(gcpNodes, GcpNodesToMap(item))
	}
	s.D.Set("gcp_nodes", gcpNodes)

	if s.Res.GcpResourceServiceAgentId != nil {
		s.D.Set("gcp_resource_service_agent_id", *s.Res.GcpResourceServiceAgentId)
	}

	if s.Res.GcpWorkloadIdentityPoolId != nil {
		s.D.Set("gcp_workload_identity_pool_id", *s.Res.GcpWorkloadIdentityPoolId)
	}

	if s.Res.GcpWorkloadIdentityProviderId != nil {
		s.D.Set("gcp_workload_identity_provider_id", *s.Res.GcpWorkloadIdentityProviderId)
	}

	if s.Res.IssuerUrl != nil {
		s.D.Set("issuer_url", *s.Res.IssuerUrl)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
