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

func DbmulticloudOracleDbAzureConnectorDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oracle_db_azure_connector_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DbmulticloudOracleDbAzureConnectorResource(), fieldMap, readSingularDbmulticloudOracleDbAzureConnector)
}

func readSingularDbmulticloudOracleDbAzureConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureConnectorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureConnectorClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureConnectorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDBAzureConnectorClient
	Res    *oci_dbmulticloud.GetOracleDbAzureConnectorResponse
}

func (s *DbmulticloudOracleDbAzureConnectorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureConnectorDataSourceCrud) Get() error {
	request := oci_dbmulticloud.GetOracleDbAzureConnectorRequest{}

	if oracleDbAzureConnectorId, ok := s.D.GetOkExists("oracle_db_azure_connector_id"); ok {
		tmp := oracleDbAzureConnectorId.(string)
		request.OracleDbAzureConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.GetOracleDbAzureConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbmulticloudOracleDbAzureConnectorDataSourceCrud) SetData() error {

	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AccessToken != nil {
		s.D.Set("access_token", *s.Res.AccessToken)
	}

	arcAgentNodes := []interface{}{}
	for _, item := range s.Res.ArcAgentNodes {
		arcAgentNodes = append(arcAgentNodes, ArcAgentNodesToMap(item))
	}
	s.D.Set("arc_agent_nodes", arcAgentNodes)

	s.D.Set("azure_identity_connectivity_status", s.Res.AzureIdentityConnectivityStatus)
	s.D.Set("azure_identity_mechanism", s.Res.AzureIdentityMechanism)

	if s.Res.AzureResourceGroup != nil {
		s.D.Set("azure_resource_group", *s.Res.AzureResourceGroup)
	}

	if s.Res.AzureSubscriptionId != nil {
		s.D.Set("azure_subscription_id", *s.Res.AzureSubscriptionId)
	}

	if s.Res.AzureTenantId != nil {
		s.D.Set("azure_tenant_id", *s.Res.AzureTenantId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbClusterResourceId != nil {
		s.D.Set("db_cluster_resource_id", *s.Res.DbClusterResourceId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.LastModification != nil {
		s.D.Set("last_modification", *s.Res.LastModification)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.PrivateEndpointDnsAlias != nil {
		s.D.Set("private_endpoint_dns_alias", *s.Res.PrivateEndpointDnsAlias)
	}

	if s.Res.PrivateEndpointIpAddress != nil {
		s.D.Set("private_endpoint_ip_address", *s.Res.PrivateEndpointIpAddress)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}
	return nil
}
