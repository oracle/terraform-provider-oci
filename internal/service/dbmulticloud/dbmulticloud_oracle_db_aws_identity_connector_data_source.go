// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudOracleDbAwsIdentityConnectorDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oracle_db_aws_identity_connector_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DbmulticloudOracleDbAwsIdentityConnectorResource(), fieldMap, readSingularDbmulticloudOracleDbAwsIdentityConnectorWithContext)
}

func readSingularDbmulticloudOracleDbAwsIdentityConnectorWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsIdentityConnectorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DbmulticloudOracleDbAwsIdentityConnectorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.DbMulticloudAwsProviderClient
	Res    *oci_dbmulticloud.GetOracleDbAwsIdentityConnectorResponse
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.GetOracleDbAwsIdentityConnectorRequest{}

	if oracleDbAwsIdentityConnectorId, ok := s.D.GetOkExists("oracle_db_aws_identity_connector_id"); ok {
		tmp := oracleDbAwsIdentityConnectorId.(string)
		request.OracleDbAwsIdentityConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.GetOracleDbAwsIdentityConnector(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AwsAccountId != nil {
		s.D.Set("aws_account_id", *s.Res.AwsAccountId)
	}

	if s.Res.AwsLocation != nil {
		s.D.Set("aws_location", *s.Res.AwsLocation)
	}

	if s.Res.AwsStsPrivateEndpoint != nil {
		s.D.Set("aws_sts_private_endpoint", *s.Res.AwsStsPrivateEndpoint)
	}

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

	if s.Res.IssuerUrl != nil {
		s.D.Set("issuer_url", *s.Res.IssuerUrl)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.OidcScope != nil {
		s.D.Set("oidc_scope", *s.Res.OidcScope)
	}

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	serviceRoleDetails := []interface{}{}
	for _, item := range s.Res.ServiceRoleDetails {
		serviceRoleDetails = append(serviceRoleDetails, ServiceRoleDetailsToMap(item))
	}
	s.D.Set("service_role_details", serviceRoleDetails)

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
