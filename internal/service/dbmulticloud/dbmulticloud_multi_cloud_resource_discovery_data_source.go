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

func DbmulticloudMultiCloudResourceDiscoveryDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["multi_cloud_resource_discovery_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DbmulticloudMultiCloudResourceDiscoveryResource(), fieldMap, readSingularDbmulticloudMultiCloudResourceDiscovery)
}

func readSingularDbmulticloudMultiCloudResourceDiscovery(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudMultiCloudResourceDiscoveryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MultiCloudResourceDiscoveryClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudMultiCloudResourceDiscoveryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.MultiCloudResourceDiscoveryClient
	Res    *oci_dbmulticloud.GetMultiCloudResourceDiscoveryResponse
}

func (s *DbmulticloudMultiCloudResourceDiscoveryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudMultiCloudResourceDiscoveryDataSourceCrud) Get() error {
	request := oci_dbmulticloud.GetMultiCloudResourceDiscoveryRequest{}

	if multiCloudResourceDiscoveryId, ok := s.D.GetOkExists("multi_cloud_resource_discovery_id"); ok {
		tmp := multiCloudResourceDiscoveryId.(string)
		request.MultiCloudResourceDiscoveryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.GetMultiCloudResourceDiscovery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbmulticloudMultiCloudResourceDiscoveryDataSourceCrud) SetData() error {
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

	if s.Res.OracleDbConnectorId != nil {
		s.D.Set("oracle_db_connector_id", *s.Res.OracleDbConnectorId)
	}

	s.D.Set("resource_type", s.Res.ResourceType)

	// s.D.Set("oracle_db_azure_vault_id", s.Res.OracleDbAzureVaultId)

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, ResourcesToMap(item))
	}
	s.D.Set("resources", resources)

	s.D.Set("resources_filter", s.Res.ResourcesFilter)

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
