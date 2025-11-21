// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globally_distributed_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_globally_distributed_database "github.com/oracle/oci-go-sdk/v65/globallydistributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GloballyDistributedDatabaseShardedDatabaseSummaryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_deployment_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"character_set": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"chunks": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cluster_certificate_common_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				ForceNew: true,
			},
			"db_workload": {
				Type:     schema.TypeString,
				ForceNew: true,
			},
			"listener_port": {
				Type:     schema.TypeInt,
				ForceNew: true,
			},
			"listener_port_tls": {
				Type:     schema.TypeInt,
				ForceNew: true,
			},
			"ncharacter_set": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ons_port_local": {
				Type:     schema.TypeInt,
				ForceNew: true,
			},
			"ons_port_remote": {
				Type:     schema.TypeInt,
				ForceNew: true,
			},
			"prefix": {
				Type:     schema.TypeString,
				ForceNew: true,
			},
			"sharding_method": {
				Type:     schema.TypeString,
				ForceNew: true,
			},
			"total_cpu_count": {
				Type:     schema.TypeFloat,
				ForceNew: true,
			},
			"total_data_storage_size_in_gbs": {
				Type:     schema.TypeFloat,
				ForceNew: true,
			},
			"replication_factor": {
				Type:     schema.TypeInt,
				ForceNew: true,
			},
			"replication_method": {
				Type:     schema.TypeString,
				ForceNew: true,
			},
			"replication_unit": {
				Type:     schema.TypeInt,
				ForceNew: true,
			},
		},
	}
}

func GloballyDistributedDatabaseShardedDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGloballyDistributedDatabaseShardedDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			/*"lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},*/
			"sharded_database_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(GloballyDistributedDatabaseShardedDatabaseSummaryResource()),
							//Elem: tfresource.GetDataSourceItemSchema(GloballyDistributedDatabaseShardedDatabaseResource()),
						},
					},
				},
			},
		},
	}
}

func readGloballyDistributedDatabaseShardedDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabaseShardedDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()

	return tfresource.ReadResource(sync)
}

type GloballyDistributedDatabaseShardedDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_globally_distributed_database.ShardedDatabaseServiceClient
	Res    *oci_globally_distributed_database.ListShardedDatabasesResponse
}

func (s *GloballyDistributedDatabaseShardedDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GloballyDistributedDatabaseShardedDatabasesDataSourceCrud) Get() error {
	request := oci_globally_distributed_database.ListShardedDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_globally_distributed_database.ShardedDatabaseLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "globally_distributed_database")

	response, err := s.Client.ListShardedDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListShardedDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GloballyDistributedDatabaseShardedDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GloballyDistributedDatabaseShardedDatabasesDataSource-", GloballyDistributedDatabaseShardedDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	shardedDatabase := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ShardedDatabaseSummaryToMap(item))
	}
	shardedDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GloballyDistributedDatabaseShardedDatabasesDataSource().Schema["sharded_database_collection"].Elem.(*schema.Resource).Schema)
		shardedDatabase["items"] = items
	}

	resources = append(resources, shardedDatabase)
	if err := s.D.Set("sharded_database_collection", resources); err != nil {
		return err
	}

	return nil
}
