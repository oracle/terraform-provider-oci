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

func DatabaseManagementCloudExadataStorageServerIormPlanDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementCloudExadataStorageServerIormPlan,
		Schema: map[string]*schema.Schema{
			"cloud_exadata_storage_server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"db_plan": {
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
									"allocation": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"asm_cluster": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"flash_cache_limit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"flash_cache_min": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"flash_cache_size": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_flash_cache_on": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_flash_log_on": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_pmem_cache_on": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_pmem_log_on": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"level": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pmem_cache_limit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pmem_cache_min": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pmem_cache_size": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"share": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"plan_objective": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plan_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseManagementCloudExadataStorageServerIormPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudExadataStorageServerIormPlanDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudExadataStorageServerIormPlanDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetCloudIormPlanResponse
}

func (s *DatabaseManagementCloudExadataStorageServerIormPlanDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudExadataStorageServerIormPlanDataSourceCrud) Get() error {
	request := oci_database_management.GetCloudIormPlanRequest{}

	if cloudExadataStorageServerId, ok := s.D.GetOkExists("cloud_exadata_storage_server_id"); ok {
		tmp := cloudExadataStorageServerId.(string)
		request.CloudExadataStorageServerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetCloudIormPlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementCloudExadataStorageServerIormPlanDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementCloudExadataStorageServerIormPlanDataSource-", DatabaseManagementCloudExadataStorageServerIormPlanDataSource(), s.D))

	if s.Res.DbPlan != nil {
		s.D.Set("db_plan", []interface{}{DatabasePlanToMap(s.Res.DbPlan)})
	} else {
		s.D.Set("db_plan", nil)
	}

	s.D.Set("plan_objective", s.Res.PlanObjective)

	s.D.Set("plan_status", s.Res.PlanStatus)

	return nil
}
