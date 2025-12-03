// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAdvancedClusterFileSystemsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseAdvancedClusterFileSystemsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"advanced_cluster_file_system_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseAdvancedClusterFileSystemResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseAdvancedClusterFileSystemsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseAdvancedClusterFileSystemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseAdvancedClusterFileSystemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAdvancedClusterFileSystemsResponse
}

func (s *DatabaseAdvancedClusterFileSystemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAdvancedClusterFileSystemsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database.ListAdvancedClusterFileSystemsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AdvancedClusterFileSystemLifecycleStateEnum(state.(string))
	}

	if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
		tmp := vmClusterId.(string)
		request.VmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAdvancedClusterFileSystems(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAdvancedClusterFileSystems(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAdvancedClusterFileSystemsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAdvancedClusterFileSystemsDataSource-", DatabaseAdvancedClusterFileSystemsDataSource(), s.D))
	resources := []map[string]interface{}{}
	advancedClusterFileSystem := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AdvancedClusterFileSystemSummaryToMap(item))
	}
	advancedClusterFileSystem["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseAdvancedClusterFileSystemsDataSource().Schema["advanced_cluster_file_system_collection"].Elem.(*schema.Resource).Schema)
		advancedClusterFileSystem["items"] = items
	}

	resources = append(resources, advancedClusterFileSystem)
	if err := s.D.Set("advanced_cluster_file_system_collection", resources); err != nil {
		return err
	}

	return nil
}
