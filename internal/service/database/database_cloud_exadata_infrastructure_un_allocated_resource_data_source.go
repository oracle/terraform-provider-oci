// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseCloudExadataInfrastructureUnAllocatedResourceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseCloudExadataInfrastructureUnAllocatedResource,
		Schema: map[string]*schema.Schema{
			"cloud_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			// Computed
			"cloud_autonomous_vm_clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"un_allocated_adb_storage_in_tbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
			"cloud_exadata_infrastructure_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exadata_storage_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"local_storage_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ocpus": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseCloudExadataInfrastructureUnAllocatedResource(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudExadataInfrastructureUnAllocatedResourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseCloudExadataInfrastructureUnAllocatedResourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetCloudExadataInfrastructureUnallocatedResourcesResponse
}

func (s *DatabaseCloudExadataInfrastructureUnAllocatedResourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseCloudExadataInfrastructureUnAllocatedResourceDataSourceCrud) Get() error {
	request := oci_database.GetCloudExadataInfrastructureUnallocatedResourcesRequest{}

	if cloudExadataInfrastructureId, ok := s.D.GetOkExists("cloud_exadata_infrastructure_id"); ok {
		tmp := cloudExadataInfrastructureId.(string)
		request.CloudExadataInfrastructureId = &tmp
	}

	if dbServers, ok := s.D.GetOkExists("db_servers"); ok {
		interfaces := dbServers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("db_servers") {
			request.DbServers = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetCloudExadataInfrastructureUnallocatedResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseCloudExadataInfrastructureUnAllocatedResourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseCloudExadataInfrastructureUnAllocatedResourceDataSource-", DatabaseCloudExadataInfrastructureUnAllocatedResourceDataSource(), s.D))

	cloudAutonomousVmClusters := []interface{}{}
	for _, item := range s.Res.CloudAutonomousVmClusters {
		cloudAutonomousVmClusters = append(cloudAutonomousVmClusters, CloudAutonomousVmClusterResourceDetailsToMap(item))
	}
	s.D.Set("cloud_autonomous_vm_clusters", cloudAutonomousVmClusters)

	if s.Res.CloudExadataInfrastructureDisplayName != nil {
		s.D.Set("cloud_exadata_infrastructure_display_name", *s.Res.CloudExadataInfrastructureDisplayName)
	}

	if s.Res.ExadataStorageInTBs != nil {
		s.D.Set("exadata_storage_in_tbs", *s.Res.ExadataStorageInTBs)
	}

	if s.Res.LocalStorageInGbs != nil {
		s.D.Set("local_storage_in_gbs", *s.Res.LocalStorageInGbs)
	}

	if s.Res.MemoryInGBs != nil {
		s.D.Set("memory_in_gbs", *s.Res.MemoryInGBs)
	}

	if s.Res.Ocpus != nil {
		s.D.Set("ocpus", *s.Res.Ocpus)
	}

	return nil
}

func CloudAutonomousVmClusterResourceDetailsToMap(obj oci_database.CloudAutonomousVmClusterResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.UnAllocatedAdbStorageInTBs != nil {
		result["un_allocated_adb_storage_in_tbs"] = float64(*obj.UnAllocatedAdbStorageInTBs)
	}

	return result
}
