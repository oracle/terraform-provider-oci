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

func DbmulticloudOracleDbAwsKeyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oracle_db_aws_key_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DbmulticloudOracleDbAwsKeyResource(), fieldMap, readSingularDbmulticloudOracleDbAwsKeyWithContext)
}

func readSingularDbmulticloudOracleDbAwsKeyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsKeyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DbmulticloudOracleDbAwsKeyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.DbMulticloudAwsProviderClient
	Res    *oci_dbmulticloud.GetOracleDbAwsKeyResponse
}

func (s *DbmulticloudOracleDbAwsKeyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAwsKeyDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.GetOracleDbAwsKeyRequest{}

	if oracleDbAwsKeyId, ok := s.D.GetOkExists("oracle_db_aws_key_id"); ok {
		tmp := oracleDbAwsKeyId.(string)
		request.OracleDbAwsKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.GetOracleDbAwsKey(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbmulticloudOracleDbAwsKeyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AwsAccountId != nil {
		s.D.Set("aws_account_id", *s.Res.AwsAccountId)
	}

	if s.Res.AwsKeyArn != nil {
		s.D.Set("aws_key_arn", *s.Res.AwsKeyArn)
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

	if s.Res.IsAwsKeyEnabled != nil {
		s.D.Set("is_aws_key_enabled", *s.Res.IsAwsKeyEnabled)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.Location != nil {
		s.D.Set("location", *s.Res.Location)
	}

	if s.Res.OracleDbConnectorId != nil {
		s.D.Set("oracle_db_connector_id", *s.Res.OracleDbConnectorId)
	}

	s.D.Set("properties", s.Res.Properties)

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

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	return nil
}
