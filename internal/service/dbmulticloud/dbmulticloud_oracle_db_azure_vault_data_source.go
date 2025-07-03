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

func DbmulticloudOracleDbAzureVaultDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oracle_db_azure_vault_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DbmulticloudOracleDbAzureVaultResource(), fieldMap, readSingularDbmulticloudOracleDbAzureVault)
}

func readSingularDbmulticloudOracleDbAzureVault(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureVaultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDbAzureVaultClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureVaultDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDbAzureVaultClient
	Res    *oci_dbmulticloud.GetOracleDbAzureVaultResponse
}

func (s *DbmulticloudOracleDbAzureVaultDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureVaultDataSourceCrud) Get() error {
	request := oci_dbmulticloud.GetOracleDbAzureVaultRequest{}

	if oracleDbAzureVaultId, ok := s.D.GetOkExists("oracle_db_azure_vault_id"); ok {
		tmp := oracleDbAzureVaultId.(string)
		request.OracleDbAzureVaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.GetOracleDbAzureVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbmulticloudOracleDbAzureVaultDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AzureVaultId != nil {
		s.D.Set("azure_vault_id", *s.Res.AzureVaultId)
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

	if s.Res.LastModification != nil {
		s.D.Set("last_modification", *s.Res.LastModification)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.Location != nil {
		s.D.Set("location", *s.Res.Location)
	}

	if s.Res.OracleDbAzureResourceGroup != nil {
		s.D.Set("oracle_db_azure_resource_group", *s.Res.OracleDbAzureResourceGroup)
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
