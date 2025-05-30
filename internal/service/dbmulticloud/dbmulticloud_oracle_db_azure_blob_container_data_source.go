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

func DbmulticloudOracleDbAzureBlobContainerDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oracle_db_azure_blob_container_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DbmulticloudOracleDbAzureBlobContainerResource(), fieldMap, readSingularDbmulticloudOracleDbAzureBlobContainer)
}

func readSingularDbmulticloudOracleDbAzureBlobContainer(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureBlobContainerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureBlobContainerClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureBlobContainerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDBAzureBlobContainerClient
	Res    *oci_dbmulticloud.GetOracleDbAzureBlobContainerResponse
}

func (s *DbmulticloudOracleDbAzureBlobContainerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureBlobContainerDataSourceCrud) Get() error {
	request := oci_dbmulticloud.GetOracleDbAzureBlobContainerRequest{}

	if oracleDbAzureBlobContainerId, ok := s.D.GetOkExists("oracle_db_azure_blob_container_id"); ok {
		tmp := oracleDbAzureBlobContainerId.(string)
		request.OracleDbAzureBlobContainerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.GetOracleDbAzureBlobContainer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbmulticloudOracleDbAzureBlobContainerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AzureStorageAccountName != nil {
		s.D.Set("azure_storage_account_name", *s.Res.AzureStorageAccountName)
	}

	if s.Res.AzureStorageContainerName != nil {
		s.D.Set("azure_storage_container_name", *s.Res.AzureStorageContainerName)
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

	if s.Res.PrivateEndpointDnsAlias != nil {
		s.D.Set("private_endpoint_dns_alias", *s.Res.PrivateEndpointDnsAlias)
	}

	if s.Res.PrivateEndpointIpAddress != nil {
		s.D.Set("private_endpoint_ip_address", *s.Res.PrivateEndpointIpAddress)
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
