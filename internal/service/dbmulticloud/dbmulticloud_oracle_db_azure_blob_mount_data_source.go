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

func DbmulticloudOracleDbAzureBlobMountDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oracle_db_azure_blob_mount_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DbmulticloudOracleDbAzureBlobMountResource(), fieldMap, readSingularDbmulticloudOracleDbAzureBlobMount)
}

func readSingularDbmulticloudOracleDbAzureBlobMount(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureBlobMountDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureBlobMountClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureBlobMountDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDBAzureBlobMountClient
	Res    *oci_dbmulticloud.GetOracleDbAzureBlobMountResponse
}

func (s *DbmulticloudOracleDbAzureBlobMountDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureBlobMountDataSourceCrud) Get() error {
	request := oci_dbmulticloud.GetOracleDbAzureBlobMountRequest{}

	if oracleDbAzureBlobMountId, ok := s.D.GetOkExists("oracle_db_azure_blob_mount_id"); ok {
		tmp := oracleDbAzureBlobMountId.(string)
		request.OracleDbAzureBlobMountId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.GetOracleDbAzureBlobMount(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbmulticloudOracleDbAzureBlobMountDataSourceCrud) SetData() error {
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

	if s.Res.LastModification != nil {
		s.D.Set("last_modification", *s.Res.LastModification)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.MountPath != nil {
		s.D.Set("mount_path", *s.Res.MountPath)
	}

	if s.Res.OracleDbAzureBlobContainerId != nil {
		s.D.Set("oracle_db_azure_blob_container_id", *s.Res.OracleDbAzureBlobContainerId)
	}

	if s.Res.OracleDbAzureConnectorId != nil {
		s.D.Set("oracle_db_azure_connector_id", *s.Res.OracleDbAzureConnectorId)
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
