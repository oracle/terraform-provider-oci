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

func DatabaseManagementCloudDbSystemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementCloudDbSystems,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dbaas_parent_infrastructure_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deployment_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_db_system_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementCloudDbSystemResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementCloudDbSystems(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudDbSystemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListCloudDbSystemsResponse
}

func (s *DatabaseManagementCloudDbSystemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudDbSystemsDataSourceCrud) Get() error {
	request := oci_database_management.ListCloudDbSystemsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbaasParentInfrastructureId, ok := s.D.GetOkExists("dbaas_parent_infrastructure_id"); ok {
		tmp := dbaasParentInfrastructureId.(string)
		request.DbaasParentInfrastructureId = &tmp
	}

	if deploymentType, ok := s.D.GetOkExists("deployment_type"); ok {
		request.DeploymentType = oci_database_management.ListCloudDbSystemsDeploymentTypeEnum(deploymentType.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database_management.ListCloudDbSystemsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListCloudDbSystems(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudDbSystems(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementCloudDbSystemsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementCloudDbSystemsDataSource-", DatabaseManagementCloudDbSystemsDataSource(), s.D))
	resources := []map[string]interface{}{}
	cloudDbSystem := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CloudDbSystemSummaryToMap(item))
	}
	cloudDbSystem["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementCloudDbSystemsDataSource().Schema["cloud_db_system_collection"].Elem.(*schema.Resource).Schema)
		cloudDbSystem["items"] = items
	}

	resources = append(resources, cloudDbSystem)
	if err := s.D.Set("cloud_db_system_collection", resources); err != nil {
		return err
	}

	return nil
}
