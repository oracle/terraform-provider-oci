// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDistributedDatabaseDistributedDatabasesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"db_deployment_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_auto_resource_management_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"metadata": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"distributed_database_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DistributedDatabaseDistributedDatabaseResource()),
						},
					},
				},
			},
		},
	}
}

func readDistributedDatabaseDistributedDatabasesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DistributedDatabaseDistributedDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_distributed_database.DistributedDbServiceClient
	Res    *oci_distributed_database.ListDistributedDatabasesResponse
}

func (s *DistributedDatabaseDistributedDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DistributedDatabaseDistributedDatabasesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.ListDistributedDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if dbDeploymentType, ok := s.D.GetOkExists("db_deployment_type"); ok {
		tmp := dbDeploymentType.(string)
		request.DbDeploymentType = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isAutoResourceManagementEnabled, ok := s.D.GetOkExists("is_auto_resource_management_enabled"); ok {
		tmp := isAutoResourceManagementEnabled.(bool)
		request.IsAutoResourceManagementEnabled = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		tmp := metadata.(string)
		request.Metadata = &tmp
	}

	if privateEndpointId, ok := s.D.GetOkExists("private_endpoint_id"); ok {
		tmp := privateEndpointId.(string)
		request.PrivateEndpointId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_distributed_database.DistributedDatabaseLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "distributed_database")

	response, err := s.Client.ListDistributedDatabases(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDistributedDatabases(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DistributedDatabaseDistributedDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DistributedDatabaseDistributedDatabasesDataSource-", DistributedDatabaseDistributedDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	distributedDatabase := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DistributedDatabaseSummaryToMap(item))
	}
	distributedDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DistributedDatabaseDistributedDatabasesDataSource().Schema["distributed_database_collection"].Elem.(*schema.Resource).Schema)
		distributedDatabase["items"] = items
	}

	resources = append(resources, distributedDatabase)
	if err := s.D.Set("distributed_database_collection", resources); err != nil {
		return err
	}

	return nil
}
