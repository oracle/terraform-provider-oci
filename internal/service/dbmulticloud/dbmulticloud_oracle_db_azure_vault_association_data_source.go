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

func DbmulticloudOracleDbAzureVaultAssociationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oracle_db_azure_vault_association_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DbmulticloudOracleDbAzureVaultAssociationResource(), fieldMap, readSingularDbmulticloudOracleDbAzureVaultAssociation)
}

func readSingularDbmulticloudOracleDbAzureVaultAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureVaultAssociationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDbAzureVaultAssociationClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureVaultAssociationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDbAzureVaultAssociationClient
	Res    *oci_dbmulticloud.GetOracleDbAzureVaultAssociationResponse
}

func (s *DbmulticloudOracleDbAzureVaultAssociationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureVaultAssociationDataSourceCrud) Get() error {
	request := oci_dbmulticloud.GetOracleDbAzureVaultAssociationRequest{}

	if oracleDbAzureVaultAssociationId, ok := s.D.GetOkExists("oracle_db_azure_vault_association_id"); ok {
		tmp := oracleDbAzureVaultAssociationId.(string)
		request.OracleDbAzureVaultAssociationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.GetOracleDbAzureVaultAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbmulticloudOracleDbAzureVaultAssociationDataSourceCrud) SetData() error {
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

	if s.Res.IsResourceAccessible != nil {
		s.D.Set("is_resource_accessible", *s.Res.IsResourceAccessible)
	}

	if s.Res.LastModification != nil {
		s.D.Set("last_modification", *s.Res.LastModification)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.OracleDbAzureConnectorId != nil {
		s.D.Set("oracle_db_azure_connector_id", *s.Res.OracleDbAzureConnectorId)
	}

	if s.Res.OracleDbAzureVaultId != nil {
		s.D.Set("oracle_db_azure_vault_id", *s.Res.OracleDbAzureVaultId)
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
