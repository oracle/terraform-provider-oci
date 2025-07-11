// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementCloudDbSystemDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cloud_db_system_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementCloudDbSystemResource(), fieldMap, readSingularDatabaseManagementCloudDbSystem)
}

func readSingularDatabaseManagementCloudDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudDbSystemDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetCloudDbSystemResponse
}

func (s *DatabaseManagementCloudDbSystemDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudDbSystemDataSourceCrud) Get() error {
	request := oci_database_management.GetCloudDbSystemRequest{}

	if cloudDbSystemId, ok := s.D.GetOkExists("cloud_db_system_id"); ok {
		tmp := cloudDbSystemId.(string)
		request.CloudDbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetCloudDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementCloudDbSystemDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseManagementConfig != nil {
		s.D.Set("database_management_config", []interface{}{CloudDbSystemDatabaseManagementConfigDetailsToMap(s.Res.DatabaseManagementConfig)})
	} else {
		s.D.Set("database_management_config", nil)
	}

	if s.Res.DbSystemDiscoveryId != nil {
		s.D.Set("db_system_discovery_id", *s.Res.DbSystemDiscoveryId)
	}

	if s.Res.DbaasParentInfrastructureId != nil {
		s.D.Set("dbaas_parent_infrastructure_id", *s.Res.DbaasParentInfrastructureId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("deployment_type", s.Res.DeploymentType)

	if s.Res.DiscoveryAgentId != nil {
		s.D.Set("discovery_agent_id", *s.Res.DiscoveryAgentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HomeDirectory != nil {
		s.D.Set("home_directory", *s.Res.HomeDirectory)
	}

	if s.Res.IsCluster != nil {
		s.D.Set("is_cluster", *s.Res.IsCluster)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.StackMonitoringConfig != nil {
		s.D.Set("stack_monitoring_config", []interface{}{CloudDbSystemStackMonitoringConfigDetailsToMap(s.Res.StackMonitoringConfig)})
	} else {
		s.D.Set("stack_monitoring_config", nil)
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
