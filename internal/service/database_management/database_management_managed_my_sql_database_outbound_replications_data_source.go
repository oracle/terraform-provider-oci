// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedMySqlDatabaseOutboundReplicationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedMySqlDatabaseOutboundReplications,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_my_sql_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_my_sql_database_outbound_replication_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"replica_host": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"replica_port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"replica_server_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"replica_uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"outbound_replications_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementManagedMySqlDatabaseOutboundReplications(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedMySqlDatabaseOutboundReplicationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedMySqlDatabasesClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedMySqlDatabaseOutboundReplicationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.ManagedMySqlDatabasesClient
	Res    *oci_database_management.ListOutboundReplicationsResponse
}

func (s *DatabaseManagementManagedMySqlDatabaseOutboundReplicationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedMySqlDatabaseOutboundReplicationsDataSourceCrud) Get() error {
	request := oci_database_management.ListOutboundReplicationsRequest{}

	if managedMySqlDatabaseId, ok := s.D.GetOkExists("managed_my_sql_database_id"); ok {
		tmp := managedMySqlDatabaseId.(string)
		request.ManagedMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListOutboundReplications(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOutboundReplications(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedMySqlDatabaseOutboundReplicationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedMySqlDatabaseOutboundReplicationsDataSource-", DatabaseManagementManagedMySqlDatabaseOutboundReplicationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedMySqlDatabaseOutboundReplication := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedMySqlDatabaseOutboundReplicationSummaryToMap(item))
	}
	managedMySqlDatabaseOutboundReplication["items"] = items

	if s.Res.OutboundReplicationsCount != nil {
		managedMySqlDatabaseOutboundReplication["outbound_replications_count"] = *s.Res.OutboundReplicationsCount
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedMySqlDatabaseOutboundReplicationsDataSource().Schema["managed_my_sql_database_outbound_replication_collection"].Elem.(*schema.Resource).Schema)
		managedMySqlDatabaseOutboundReplication["items"] = items
	}

	resources = append(resources, managedMySqlDatabaseOutboundReplication)
	if err := s.D.Set("managed_my_sql_database_outbound_replication_collection", resources); err != nil {
		return err
	}

	return nil
}

func ManagedMySqlDatabaseOutboundReplicationSummaryToMap(obj oci_database_management.ManagedMySqlDatabaseOutboundReplicationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ReplicaHost != nil {
		result["replica_host"] = string(*obj.ReplicaHost)
	}

	if obj.ReplicaPort != nil {
		result["replica_port"] = int(*obj.ReplicaPort)
	}

	if obj.ReplicaServerId != nil {
		result["replica_server_id"] = strconv.FormatInt(*obj.ReplicaServerId, 10)
	}

	if obj.ReplicaUuid != nil {
		result["replica_uuid"] = string(*obj.ReplicaUuid)
	}

	return result
}
